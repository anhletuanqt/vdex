package worker

import (
	"context"
	"database/sql"

	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
	"github.com/cxptek/vdex/service"
	"github.com/sirupsen/logrus"
)

type FillMaker struct {
	fillCh       chan *models.Fill
	logReader    matching.LogReader
	logOffset    int64
	pushingStore *KafkaPushingStore
}

func NewFillMaker(logReader matching.LogReader, pushingStore *KafkaPushingStore) *FillMaker {
	ctx := context.Background()

	t := &FillMaker{
		fillCh:       make(chan *models.Fill, 1000),
		logReader:    logReader,
		pushingStore: pushingStore,
	}

	lastFill, err := pg.SharedStore().GetLastFillByProductID(ctx, logReader.GetProductId())
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if lastFill != nil {
		t.logOffset = lastFill.LogOffset
	}
	t.logReader.RegisterObserver(t)

	return t
}

func (t *FillMaker) Start() {
	if t.logOffset > 0 {
		t.logOffset++
	}
	go t.logReader.Run(t.logOffset)
	go t.flusher()
}

func (t *FillMaker) OnOpenLog(log *matching.LogOrder, offset int64) {
	order, err := service.UpdateOrder(context.Background(), log.OrderID, models.OrderStatusNew, models.OrderStatusOpen)
	if err != nil {
		logrus.Error(err)
		return
	}

	t.pushingStore.StoreOrder([]*models.Order{order})
}

func (t *FillMaker) OnCancelLog(log *matching.LogOrder, offset int64) {
	t.fillCh <- &models.Fill{
		OrderID:    log.OrderID,
		ProductID:  log.ProductID,
		Size:       log.Quantity,
		Price:      log.Price,
		Side:       models.Side(log.Side),
		DoneReason: models.DoneReasonCancelled,
		LogOffset:  offset,
	}
}

func (t *FillMaker) OnDoneLog(log *matching.LogOrder, offset int64) {
	t.fillCh <- &models.Fill{
		OrderID:    log.TakerOrderID,
		ProductID:  log.ProductID,
		Size:       log.Quantity,
		Price:      log.Price,
		Side:       models.Side(log.Side).Opposite(),
		DoneReason: models.DoneReasonFilled,
		LogOffset:  offset,
		Liquidity:  "T",
	}
	t.fillCh <- &models.Fill{
		OrderID:    log.MakerOrderID,
		ProductID:  log.ProductID,
		Size:       log.Quantity,
		Price:      log.Price,
		Side:       models.Side(log.Side),
		DoneReason: models.DoneReasonFilled,
		LogOffset:  offset,
		Liquidity:  "M",
	}
}

func (t *FillMaker) flusher() {
	ctx := context.Background()

	for {
		select {
		case fill := <-t.fillCh:
			order, err := service.GetOrderByID(ctx, fill.OrderID)
			if err != nil {
				logrus.Error(err)
				continue
			}
			fill.UserID = order.UserID

			if err := service.AddFills(ctx, []*models.Fill{fill}); err != nil {
				logrus.Error(err)
			}
		}
	}
}
