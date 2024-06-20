package matching

import (
	"context"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/service"
	"github.com/sirupsen/logrus"
)

func StartEngine() {
	ctx := context.Background()
	products, err := service.GetProducts(ctx)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		orderReader := NewKafkaOrderReader(product.ID, config.GetConfig().Kafka.Brokers)
		logStore := NewKafkaLogStore(product.ID, config.GetConfig().Kafka.Brokers)
		snapshotStore := NewRedisSnapshotStore(product.ID)
		matchEngine := NewEngine(&product, orderReader, logStore, snapshotStore)
		matchEngine.Start()
	}
	logrus.Infoln("matching is running")
}
