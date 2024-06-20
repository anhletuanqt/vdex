package worker

import (
	"context"
	"encoding/json"

	"github.com/cxptek/vdex/models"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type PushingObserver interface {
	PushOrder(order *models.Order, logOffset int64)
	PushTrade(trade *models.Trade, logOffset int64)
}

type KafkaPushingReader struct {
	readerId    string
	productId   string
	orderReader *kafka.Reader
	tradeReader *kafka.Reader
	tickReader  *kafka.Reader
	observer    PushingObserver
}

func NewKafkaPushingReader(readerId string, brokers []string) *KafkaPushingReader {
	tradeReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topicPushingTrade,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	orderReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topicPushingOrder,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	tickReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topicPushingTick,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})

	tradeReader.SetOffset(0)
	orderReader.SetOffset(0)
	tickReader.SetOffset(0)

	return &KafkaPushingReader{
		readerId:    readerId,
		tradeReader: tradeReader,
		orderReader: orderReader,
		tickReader:  tickReader,
	}
}

func (r *KafkaPushingReader) RegisterObserver(observer PushingObserver) {
	r.observer = observer
}

func (r *KafkaPushingReader) RunOrderReader(offset int64) {
	err := r.orderReader.SetOffset(offset)
	if err != nil {
		panic(err)
	}

	for {
		kMessage, err := r.orderReader.FetchMessage(context.Background())
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		var order models.Order
		err = json.Unmarshal(kMessage.Value, &order)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		r.observer.PushOrder(&order, kMessage.Offset)
	}
}
func (r *KafkaPushingReader) RunTradeReader(offset int64) {
	err := r.tradeReader.SetOffset(offset)
	if err != nil {
		panic(err)
	}

	for {
		kMessage, err := r.tradeReader.FetchMessage(context.Background())
		if err != nil {
			continue
		}
		var trade models.Trade
		err = json.Unmarshal(kMessage.Value, &trade)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		r.observer.PushTrade(&trade, kMessage.Offset)
	}
}
