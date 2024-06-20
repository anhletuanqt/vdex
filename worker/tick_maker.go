package worker

import (
	"context"
	"database/sql"
	"time"

	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/service"
	"github.com/jinzhu/now"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var minutes = []int64{1, 5, 15, 30, 60, 120, 240, 360, 480, 720, 1440, 10080, 43200}

type TickMaker struct {
	ticks        map[int64]*models.Tick
	tickCh       chan models.Tick
	logReader    matching.LogReader
	logOffset    int64
	pushingStore *KafkaPushingStore
}

func NewTickMaker(logReader matching.LogReader, pushingStore *KafkaPushingStore) *TickMaker {
	ctx := context.Background()

	t := &TickMaker{
		ticks:        map[int64]*models.Tick{},
		tickCh:       make(chan models.Tick, 1000),
		logReader:    logReader,
		pushingStore: pushingStore,
	}
	for _, granularity := range minutes {
		tick, err := service.GetLastTickByProductID(ctx, logReader.GetProductId(), granularity)
		if err != nil && err != sql.ErrNoRows {
			panic(err)
		}

		if tick != nil {
			t.ticks[granularity] = tick
			t.logOffset = tick.LogOffset
		}
	}
	t.logReader.RegisterObserver(t)

	return t
}

func (t *TickMaker) Start() {
	if t.logOffset > 0 {
		t.logOffset++
	}
	go t.logReader.Run(t.logOffset)
	go t.flusher()
}

func (t *TickMaker) OnOpenLog(log *matching.LogOrder, offset int64) {
	// do nothing
}

func (t *TickMaker) OnCancelLog(log *matching.LogOrder, offset int64) {
	// do nothing
}

func (t *TickMaker) OnDoneLog(log *matching.LogOrder, offset int64) {
	for _, granularity := range minutes {
		tickTime := log.Time.UTC().Truncate(time.Duration(granularity) * time.Minute).Unix()
		if granularity == 43200 { // 1 month
			tickTime = now.New(log.Time.UTC()).BeginningOfMonth().Unix()
		}
		tick, found := t.ticks[granularity]
		if !found || tick.Time != tickTime {
			tick = &models.Tick{
				Open:        log.Price,
				Close:       log.Price,
				Low:         log.Price,
				High:        log.Price,
				Volume:      log.Quantity,
				ProductID:   log.ProductID,
				Granularity: granularity,
				Time:        tickTime,
				LogOffset:   offset,
			}
			t.ticks[granularity] = tick
		} else {
			tick.Close = log.Price
			tick.Low = decimal.Min(tick.Low, log.Price)
			tick.High = decimal.Max(tick.High, log.Price)
			tick.Volume = tick.Volume.Add(log.Quantity)
			tick.LogOffset = offset
		}

		t.tickCh <- *tick
	}
}

func (t *TickMaker) flusher() {
	var ticks []*models.Tick
	ctx := context.Background()

	for {
		select {
		case tick := <-t.tickCh:
			_, index, _ := lo.FindIndexOf[*models.Tick](ticks, func(_t *models.Tick) bool {
				if _t.ProductID == tick.ProductID && _t.Time == tick.Time && _t.Granularity == tick.Granularity {
					return true
				}
				return false
			})
			if index != -1 {
				ticks[index] = &tick
			} else {
				ticks = append(ticks, &tick)
			}

			if len(t.tickCh) > 0 && len(ticks) < 1000 {
				continue
			}

			if err := service.AddTicks(ctx, ticks); err != nil {
				time.Sleep(time.Second)
				logrus.Error(err)
				continue
			}
			// push to kafka
			if err := t.pushingStore.StoreTick(ticks); err != nil {
				logrus.Error(err)
			}
			ticks = nil
		}
	}
}
