package pushing

import (
	"context"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/service"
	"github.com/cxptek/vdex/worker"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	ctx := context.Background()
	gbeConfig := config.GetConfig()

	sub := newSubscription()
	NewRedisPushingSnapshotStore()
	products, err := service.GetProducts(ctx)
	if err != nil {
		panic(err)
	}

	newTradeStream(sub, worker.NewKafkaPushingReader("tradeStream", gbeConfig.Kafka.Brokers)).Start()
	newOrderStream(sub, worker.NewKafkaPushingReader("orderStream", gbeConfig.Kafka.Brokers)).Start()
	for _, product := range products {
		NewOrderbookStream(sub, matching.NewKafkaLogReader("orderBookStream", product.ID, config.GetConfig().Kafka.Brokers)).Start()
	}

	NewServer(gbeConfig.PushServer.Port, gbeConfig.PushServer.Path, sub).Run()
	logrus.Infoln("websocket server ok")

}
