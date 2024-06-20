package event_listener

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/models"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var productId2Writer sync.Map

func getWriter(productID string) *kafka.Writer {
	writer, found := productId2Writer.Load(productID)
	if found {
		return writer.(*kafka.Writer)
	}

	newWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      config.GetConfig().Kafka.Brokers,
		Topic:        matching.TopicOrderPrefix + productID,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})
	productId2Writer.Store(productID, newWriter)
	return newWriter
}

func submitOrder(order *models.Order) {
	buf, err := json.Marshal(order)
	if err != nil {
		logrus.Error(err)
		return
	}

	err = getWriter(order.ProductID).WriteMessages(context.Background(), kafka.Message{Value: buf})
	if err != nil {
		logrus.Error(err)
	}
}
