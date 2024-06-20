package pushing

import (
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/worker"
)

type OrderStream struct {
	// productID     string
	sub           *subscription
	pushingReader *worker.KafkaPushingReader
	logOffset     int64
}

func newOrderStream(sub *subscription, pushingReader *worker.KafkaPushingReader) *OrderStream {
	s := &OrderStream{
		// productID:     productID,
		sub:           sub,
		pushingReader: pushingReader,
	}
	// if err := redisPushingSnapshotStore.StorePushingOffset(string(ChannelOrder), 0); err != nil {
	// 	logrus.Errorln(err)
	// }
	// offset, err := redisPushingSnapshotStore.GetPushingOffset(string(ChannelOrder))
	// if err != nil {
	// 	panic(err)
	// }
	s.logOffset = -1

	s.pushingReader.RegisterObserver(s)
	return s
}

func (s *OrderStream) Start() {
	if s.logOffset > 0 {
		s.logOffset++
	}

	go s.pushingReader.RunOrderReader(s.logOffset)
}

func (s *OrderStream) PushOrder(order *models.Order, logOffset int64) {
	channel := string(ChannelOrder)
	data := Response{
		Type: channel,
		Data: order,
	}
	// if order.ID != 0 {
	// logrus.Infof("PushOrder %v \n", *order)
	// }
	s.sub.publish(ChannelOrder.Format(order.ProductID, order.UserID), data)
	// if err := redisPushingSnapshotStore.StorePushingOffset(channel, logOffset); err != nil {
	// 	logrus.Errorln(err)
	// }
}

func (s *OrderStream) PushTrade(log *models.Trade, logOffset int64) {
	// do nothing
}
