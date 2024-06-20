package pushing

import (
	"time"

	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/orderbook"
	"github.com/shopspring/decimal"
)

type OrderbookStream struct {
	logCh     chan *matching.LogOrder
	logReader matching.LogReader
	sub       *subscription
	logOffset int64
}

func NewOrderbookStream(sub *subscription, logReader matching.LogReader) *OrderbookStream {
	t := &OrderbookStream{
		logCh:     make(chan *matching.LogOrder, 1000),
		logReader: logReader,
		sub:       sub,
	}
	// offset, err := redisPushingSnapshotStore.GetPushingOffset(string(ChannelOrderbook))
	// if err != nil {
	// 	panic(err)
	// }
	t.logOffset = -1

	t.logReader.RegisterObserver(t)

	return t
}

func (t *OrderbookStream) Start() {
	if t.logOffset > 0 {
		t.logOffset++
	}
	go t.logReader.Run(t.logOffset)
	go t.runFlusher()
}

func (t *OrderbookStream) OnOpenLog(log *matching.LogOrder, offset int64) {
	t.logCh <- log
	t.logOffset = offset
}

func (t *OrderbookStream) OnCancelLog(log *matching.LogOrder, offset int64) {
	t.logCh <- log
	t.logOffset = offset
}

func (t *OrderbookStream) OnDoneLog(log *matching.LogOrder, offset int64) {
	t.logCh <- log
	t.logOffset = offset
}

func (t *OrderbookStream) runFlusher() {
	channel := string(ChannelOrderbook)
	for {
		select {
		case log := <-t.logCh:
			volume := orderbook.GetOrderBookByID(log.ProductID).GetOrderSide(log.Side).VolumeByPrice(log.Price)
			oppositeSideVolume := orderbook.GetOrderBookByID(log.ProductID).GetOrderSide(log.Side.Opposite()).VolumeByPrice(log.Price)

			asks := [][2]decimal.Decimal{
				{log.Price, volume},
			}
			bids := [][2]decimal.Decimal{
				{log.Price, oppositeSideVolume},
			}
			if oppositeSideVolume.IsZero() {
				bids = [][2]decimal.Decimal{}
			}

			if log.Side == orderbook.Buy {
				asks = [][2]decimal.Decimal{
					{log.Price, oppositeSideVolume},
				}
				if oppositeSideVolume.IsZero() {
					asks = [][2]decimal.Decimal{}
				}
				bids = [][2]decimal.Decimal{
					{log.Price, volume},
				}
			}
			data := Response{
				Type: channel,
				Data: map[string]interface{}{
					"asks":      asks,
					"bids":      bids,
					"productId": log.ProductID,
				},
			}
			t.sub.publish(ChannelOrderbook.FormatWithProductID(log.ProductID), data)
		case <-time.After(3 * time.Second):
			redisPushingSnapshotStore.StorePushingOffset(channel, t.logOffset)
		}
	}
}
