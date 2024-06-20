package event_listener

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/contract"
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/service"
	"github.com/cxptek/vdex/util"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type OrderCreated struct {
	event  *contract.VdexOrderCreated
	txHash string
}

type EventListener struct {
	orderCreatedCh chan *OrderCreated
	ethClient      *ethclient.Client
	workerConfig   *models.WorkerConfig
}

func NewOnchainEvent() *EventListener {
	ctx := context.Background()
	ethClient, err := ethclient.DialContext(ctx, config.GetConfig().JsonRpc)
	if err != nil {
		logrus.Panic(err)
	}
	block, err := ethClient.BlockNumber(ctx)
	if err != nil {
		logrus.Panic(err)
	}
	workerConfig, err := service.GetWorkerConfig(ctx)
	if err != nil {
		if err != sql.ErrNoRows {
			logrus.Panic(err)
		}

		workerConfig = &models.WorkerConfig{
			BlockNumber: block - 3,
		}
		if err := service.AddWorkerConfig(ctx, workerConfig); err != nil {
			logrus.Panic(err)
		}
	}

	return &EventListener{
		orderCreatedCh: make(chan *OrderCreated, 1000),
		ethClient:      ethClient,
		workerConfig:   workerConfig,
	}
}

func (e *EventListener) Start() {
	go e.eventLogs()
	go e.flusher()
}

func (e *EventListener) eventLogs() {
	ctx := context.Background()
	workerConfig := e.workerConfig

	for {
		time.Sleep(1 * time.Second)
		// get latest block
		block, err := e.ethClient.BlockNumber(ctx)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		if workerConfig.BlockNumber == 0 {
			workerConfig.BlockNumber = block - 3
		}

		latestBlock := block - 6

		fromBlock := workerConfig.BlockNumber + 1
		toBlock := fromBlock + 1000
		if fromBlock > latestBlock {
			time.Sleep(1 * time.Second)
			continue
		}
		if toBlock > latestBlock {
			toBlock = latestBlock
		}

		//** HANDLE EVENTS **//
		if err := e.vdexEvents(ctx, fromBlock, toBlock); err != nil {
			logrus.Errorln(err)
			continue
		}
		// update WorkerConfig
		workerConfig.BlockNumber = toBlock
		update := &models.WorkerConfig{
			ID:          workerConfig.ID,
			BlockNumber: toBlock,
		}
		if err := service.UpdateWorkerConfig(ctx, update); err != nil {
			logrus.Errorln(err)
		}
	}
}

func (e *EventListener) vdexEvents(ctx context.Context, fromBlock uint64, toBlock uint64) error {
	contractAddress := common.HexToAddress(config.GetConfig().Contracts.Vdex)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := e.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	contractAbi, err := abi.JSON(strings.NewReader(contract.VdexABI))
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	orderCreatedLock := crypto.Keccak256Hash([]byte("OrderCreated(bytes32,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256)"))
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case orderCreatedLock.Hex():
			var event contract.VdexOrderCreated
			if err := contractAbi.UnpackIntoInterface(&event, "OrderCreated", vLog.Data); err != nil {
				logrus.Errorln(err)
				continue
			}

			event.OrderHash = vLog.Topics[1]
			txHash := strings.ToLower(vLog.TxHash.Hex())
			orderCreated := &OrderCreated{
				event:  &event,
				txHash: txHash,
			}
			e.orderCreatedCh <- orderCreated
		}
	}

	return nil
}

func (e *EventListener) flusher() {
	ctx := context.Background()
	for {
		select {
		case orderCreated := <-e.orderCreatedCh:
			event := orderCreated.event

			// get user
			user, err := service.GetUserByAddress(ctx, strings.ToLower(event.MakerAccountOwner.Hex()))
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			baseToken := config.GetConfig().Tokens[strings.ToLower(event.BaseAsset.Hex())]
			quoteToken := config.GetConfig().Tokens[strings.ToLower(event.QuoteAsset.Hex())]
			productID := fmt.Sprintf("%v-%v", baseToken, quoteToken)
			product, err := service.GetProductByID(ctx, productID)
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			side := models.SideSell
			if event.OrderSide.Uint64() == 0 { // 0: buy, 1: sell
				side = models.SideBuy
			}
			size := util.ToDecimal(event.BaseAssetAmount, 18)
			funds := util.ToDecimal(event.QuoteAssetAmount, 18)
			price := funds.Div(size)
			order := &models.Order{
				CreatedTxHash: orderCreated.txHash,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
				UserID:        user.ID,
				ProductID:     product.ID,
				Side:          side,
				Price:         price,
				Size:          size,
				Funds:         funds,
				Status:        models.OrderStatusNew,
				Type:          models.NewOrderTypeFromOnChain(event.OrderType.Int64()),
				Nonce:         event.Nonce.Int64(),
				Expiration:    event.Expiration.Int64(),
				TimeInForce:   "gtc",
			}
			if err := service.AddOrderOnConflictTxHash(ctx, order); err != nil {
				logrus.Error(err)
				continue
			}
			spew.Dump(order)
			buf, err := json.Marshal(order)
			if err != nil {
				logrus.Error(err)
				continue
			}
			if err := getWriter(order.ProductID).WriteMessages(context.Background(), kafka.Message{Value: buf}); err != nil {
				logrus.Error(err)
			}
		}
	}
}
