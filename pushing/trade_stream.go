package pushing

import (
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/worker"
)

type TradeStream struct {
	// productID     string
	sub           *subscription
	pushingReader *worker.KafkaPushingReader
	logOffset     int64
}

func newTradeStream(sub *subscription, pushingReader *worker.KafkaPushingReader) *TradeStream {
	s := &TradeStream{
		// productID:     productID,
		sub:           sub,
		pushingReader: pushingReader,
	}
	// offset, err := redisPushingSnapshotStore.GetPushingOffset(string(ChannelTrade))
	// if err != nil {
	// 	panic(err)
	// }
	s.logOffset = -1
	s.pushingReader.RegisterObserver(s)

	return s
}

func (s *TradeStream) Start() {
	if s.logOffset > 0 {
		s.logOffset++
	}
	go s.pushingReader.RunTradeReader(s.logOffset)
}

func (s *TradeStream) PushOrder(log *models.Order, logOffset int64) {
	// do nothing
}

func (s *TradeStream) PushTrade(trade *models.Trade, logOffset int64) {
	channel := string(ChannelTrade)
	data := Response{
		Type: channel,
		Data: trade,
	}

	s.sub.publish(ChannelTrade.FormatWithProductID(trade.ProductID), data)
	// if err := redisPushingSnapshotStore.StorePushingOffset(channel, logOffset); err != nil {
	// 	logrus.Errorln(err)
	// }
}
