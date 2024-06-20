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

type TradeMaker struct {
	tradeCh      chan *models.Trade
	logReader    matching.LogReader
	logOffset    int64
	pushingStore *KafkaPushingStore
}

func NewTradeMaker(logReader matching.LogReader, pushingStore *KafkaPushingStore) *TradeMaker {
	ctx := context.Background()
	t := &TradeMaker{
		tradeCh:      make(chan *models.Trade, 1000),
		logReader:    logReader,
		pushingStore: pushingStore,
	}
	lastTrade, err := pg.SharedStore().GetLastTradeByProductID(ctx, logReader.GetProductId())
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if lastTrade != nil {
		t.logOffset = lastTrade.LogOffset
	}
	t.logReader.RegisterObserver(t)

	return t
}

func (t *TradeMaker) Start() {
	if t.logOffset > 0 {
		t.logOffset++
	}
	go t.logReader.Run(t.logOffset)
	go t.runFlusher()
}
func (t *TradeMaker) OnOpenLog(log *matching.LogOrder, offset int64) {
	// do nothing
}

func (t *TradeMaker) OnCancelLog(log *matching.LogOrder, offset int64) {
	// do nothing
}

func (t *TradeMaker) OnDoneLog(log *matching.LogOrder, offset int64) {
	t.tradeCh <- &models.Trade{
		ProductID:    log.ProductID,
		TakerOrderID: log.TakerOrderID,
		MakerOrderID: log.MakerOrderID,
		Price:        log.Price,
		Size:         log.Quantity,
		Side:         models.Side(log.Side),
		Time:         &log.Time,
		LogOffset:    offset,
	}
}

func (t *TradeMaker) runFlusher() {
	ctx := context.Background()
	var trades []*models.Trade
	for {
		select {
		case trade := <-t.tradeCh:
			trades = append(trades, trade)

			if len(t.tradeCh) > 0 && len(trades) < 1000 {
				continue
			}

			if err := service.ExecuteTrades(ctx, trades); err != nil {
				logrus.Error(err)
				continue
			}

			if err := t.pushingStore.StoreTrade(trades); err != nil {
				logrus.Error(err)
			}
			trades = nil
		}
	}
}
