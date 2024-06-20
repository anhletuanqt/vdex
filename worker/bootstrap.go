package worker

import (
	"context"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/service"
)

func Start() {
	ctx := context.Background()
	products, err := service.GetProducts(ctx)
	if err != nil {
		panic(err)
	}
	pushingStore := NewKafkaPushingStore(config.GetConfig().Kafka.Brokers)

	for _, product := range products {
		NewTickMaker(matching.NewKafkaLogReader("tickMaker", product.ID, config.GetConfig().Kafka.Brokers), pushingStore).Start()
		NewTradeMaker(matching.NewKafkaLogReader("tradeMaker", product.ID, config.GetConfig().Kafka.Brokers), pushingStore).Start()
		NewFillMaker(matching.NewKafkaLogReader("fillMaker", product.ID, config.GetConfig().Kafka.Brokers), pushingStore).Start()
		NewUserSegmentMaker(matching.NewKafkaLogReader("userSegmentMaker", product.ID, config.GetConfig().Kafka.Brokers)).Start()
	}
}
