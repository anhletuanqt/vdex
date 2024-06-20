package worker

import (
	"context"
	"fmt"
	"time"

	lru "github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/sirupsen/logrus"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/service"
)

const fillWorkerNum = 10

type FillExecutor struct {
	workerChs    [fillWorkerNum]chan *models.Fill
	pushingStore *KafkaPushingStore
}

func NewFillExecutor() *FillExecutor {
	ctx := context.Background()

	f := &FillExecutor{
		workerChs:    [fillWorkerNum]chan *models.Fill{},
		pushingStore: NewKafkaPushingStore(config.GetConfig().Kafka.Brokers),
	}

	for i := 0; i < fillWorkerNum; i++ {
		f.workerChs[i] = make(chan *models.Fill, 512)
		go func(idx int) {
			settledOrderCache := lru.NewLRU[string, string](1000, nil, time.Second*10)

			for {
				select {
				case fill := <-f.workerChs[idx]:
					cachedKey := fmt.Sprintf("%v", fill.OrderID)
					if settledOrderCache.Contains(cachedKey) {
						continue
					}
					settledOrderCache.Add(cachedKey, "")
					order, err := service.GetOrderByID(ctx, fill.OrderID)
					if err != nil {
						logrus.Warnf("order not found: %v", fill.OrderID)
						continue
					}

					if order.Status == models.OrderStatusCancelled || order.Status == models.OrderStatusFilled {
						continue
					}
					if err := service.ExecuteFill(ctx, fill.OrderID); err != nil {
						logrus.Error(err, fill.OrderID)
					}
					// push to kafka
					order, err = service.GetOrderByID(ctx, fill.OrderID)
					if err != nil {
						logrus.Warnf("order not found: %v", fill.OrderID)
						continue
					}
					if err := f.pushingStore.StoreOrder([]*models.Order{order}); err != nil {
						logrus.Error(err)
					}
				}
			}
		}(i)
	}

	return f
}

func (s *FillExecutor) Start() {
	go s.runInspector()
}

func (s *FillExecutor) runInspector() {
	ctx := context.Background()
	for {
		select {
		case <-time.After(2 * time.Second):
			fills, err := service.GetUnsettledFills(ctx, 1000)
			if err != nil {
				logrus.Error(err)
				continue
			}

			for _, fill := range fills {
				s.workerChs[fill.OrderID%fillWorkerNum] <- fill
			}
		}
	}
}
