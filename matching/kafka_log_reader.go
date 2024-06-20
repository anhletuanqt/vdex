package matching

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaLogReader struct {
	readerId  string
	productId string
	reader    *kafka.Reader
	observer  LogObserver
}

func NewKafkaLogReader(readerId, productId string, brokers []string) LogReader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topicBookMessagePrefix + productId,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	return &KafkaLogReader{readerId: readerId, productId: productId, reader: reader}
}

func (r *KafkaLogReader) GetProductId() string {
	return r.productId
}

func (r *KafkaLogReader) RegisterObserver(observer LogObserver) {
	r.observer = observer
}

func (r *KafkaLogReader) Run(offset int64) {
	err := r.reader.SetOffset(offset)
	if err != nil {
		panic(err)
	}

	for {
		kMessage, err := r.reader.FetchMessage(context.Background())
		if err != nil {
			continue
		}
		var log LogOrder
		err = json.Unmarshal(kMessage.Value, &log)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		switch log.LogType {
		case LogTypeOpen:
			r.observer.OnOpenLog(&log, kMessage.Offset)
		case LogTypeDone:
			r.observer.OnDoneLog(&log, kMessage.Offset)
		case LogTypeCancel:
			r.observer.OnCancelLog(&log, kMessage.Offset)
		}
	}
}
