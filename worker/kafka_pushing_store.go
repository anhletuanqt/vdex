package worker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cxptek/vdex/models"
	"github.com/segmentio/kafka-go"
)

const (
	topicPushingOrder = "pushing_order"
	topicPushingTrade = "pushing_trade"
	topicPushingTick  = "pushing_tick"
)

type KafkaPushingStore struct {
	orderLogWriter *kafka.Writer
	tradeLogWriter *kafka.Writer
	tickLogWriter  *kafka.Writer
}

func NewKafkaPushingStore(brokers []string) *KafkaPushingStore {
	s := &KafkaPushingStore{}

	s.orderLogWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topicPushingOrder,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})
	s.tradeLogWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topicPushingTrade,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})

	s.tickLogWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topicPushingTick,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})
	return s
}

func (s *KafkaPushingStore) StoreOrder(logs []*models.Order) error {
	var messages []kafka.Message
	for _, log := range logs {
		val, err := json.Marshal(log)
		if err != nil {
			return err
		}
		messages = append(messages, kafka.Message{Value: val})
	}

	return s.orderLogWriter.WriteMessages(context.Background(), messages...)
}

func (s *KafkaPushingStore) StoreTrade(logs []*models.Trade) error {
	var messages []kafka.Message
	for _, log := range logs {
		val, err := json.Marshal(log)
		if err != nil {
			return err
		}
		messages = append(messages, kafka.Message{Value: val})
	}

	return s.tradeLogWriter.WriteMessages(context.Background(), messages...)
}

func (s *KafkaPushingStore) StoreTick(logs []*models.Tick) error {
	var messages []kafka.Message
	for _, log := range logs {
		val, err := json.Marshal(log)
		if err != nil {
			return err
		}
		messages = append(messages, kafka.Message{Value: val})
	}

	return s.tickLogWriter.WriteMessages(context.Background(), messages...)
}
