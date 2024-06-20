package worker

import (
	"context"
	"crypto/ecdsa"
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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	lru "github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/sirupsen/logrus"
)

type SettlementExecutor struct {
	workerChs []chan *models.Settlement
	ethClient *ethclient.Client
	numWorker int64
}

func NewSettlementExecutor() *SettlementExecutor {
	ctx := context.Background()
	ethClient, err := ethclient.DialContext(ctx, config.GetConfig().JsonRpc)
	if err != nil {
		logrus.Panic(err)
	}
	numWorker := len(config.GetConfig().DispatcherWallets)
	t := &SettlementExecutor{
		workerChs: make([]chan *models.Settlement, numWorker),
		ethClient: ethClient,
		numWorker: int64(numWorker),
	}

	for i := 0; i < numWorker; i++ {
		t.workerChs[i] = make(chan *models.Settlement, 512)
		go func(idx int) {
			settledCache := lru.NewLRU[string, string](1000, nil, time.Minute*2)

			for {
				select {
				case settlement := <-t.workerChs[idx]:
					cacheKey := fmt.Sprintf("%v-%v", settlement.RefID, settlement.Type)
					if settledCache.Contains(cacheKey) {
						continue
					}
					settledCache.Add(cacheKey, "")

					if err := service.UpdateSettlementFee(ctx, settlement); err != nil {
						logrus.Errorln(err)
						continue
					}

					var tx *types.Transaction
					var err error
					if settlement.Type == models.SettlementOrder {
						tx, err = t.submitCancelOrder(idx, settlement)
					} else if settlement.Type == models.SettlementTrade {
						tx, err = t.submitExecuteOrderBookTrade(idx, settlement)
					}
					if err != nil || tx == nil {
						logrus.Errorln(err)
						continue
					}

					receipt, err := bind.WaitMined(ctx, ethClient, tx)
					if err != nil {
						logrus.Errorln(err)
						continue
					}
					if receipt.Status == types.ReceiptStatusFailed {
						logrus.Errorln("tx failed in settlement", settlement.ID)
						continue
					}

					settlement.Settled = true
					settlement.TxHash = strings.ToLower(tx.Hash().Hex())
					if err := service.UpdateSettlement(ctx, settlement); err != nil {
						logrus.Errorln("tx failed in trade", settlement.ID)
					}
					logrus.Infoln("TxHash", tx.Hash().Hex())
				}
			}
		}(i)
	}

	return t
}

func (s *SettlementExecutor) Start() {
	go s.runInspector()
}

func (s *SettlementExecutor) runInspector() {
	ctx := context.Background()
	for {
		select {
		case <-time.After(3 * time.Second):
			settlements, err := service.GetUnsettledSettlements(ctx, 1000)
			if err != nil {
				logrus.Error(err)
				continue
			}

			for _, settlement := range settlements {
				s.workerChs[settlement.ID%int64(s.numWorker)] <- settlement
			}
		}
	}
}

func (t *SettlementExecutor) submitCancelOrder(idx int, settlement *models.Settlement) (*types.Transaction, error) {
	ctx := context.Background()
	order, err := service.GetOrderByID(ctx, settlement.RefID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	trader, err := service.GetUserByID(ctx, order.UserID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	product, err := service.GetProductByID(ctx, order.ProductID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	vdexContract := common.HexToAddress(config.GetConfig().Contracts.Vdex)
	vdex, err := contract.NewVdex(vdexContract, t.ethClient)
	if err != nil {
		return nil, err
	}
	var orderType uint8 = 0 // market
	if order.Type == models.OrderTypeLimit {
		orderType = 1
	}

	// 0: buy, 1: sell
	var orderSide uint8 = 0
	if order.Side == models.SideSell {
		orderSide = 1
	}
	swapOrder := contract.BaseOrderUtilsCancelOrderParams{
		WalletSignature: []byte("0x0"),
		Order: contract.Order{
			Trader:           common.HexToAddress(trader.Address),
			BaseAsset:        common.HexToAddress(config.GetAddressBySymbol(product.BaseCurrency)),
			QuoteAsset:       common.HexToAddress(config.GetAddressBySymbol(product.QuoteCurrency)),
			BaseAssetAmount:  util.ToWei(order.Size, 18),
			QuoteAssetAmount: util.ToWei(order.Funds, 18),
			OrderType:        orderType,
			Side:             orderSide,
			Expiration:       big.NewInt(order.Expiration),
			Nonce:            big.NewInt(order.Nonce),
		},
		CancellationFee: util.ToWei(settlement.MakerFee, 18),
	}
	swapOrders := []contract.BaseOrderUtilsCancelOrderParams{swapOrder}
	txOps, err := getTxOpsForCancelOrder(idx, t.ethClient, vdexContract, swapOrders)
	if err != nil {
		spew.Dump(swapOrder)
		logrus.Errorln("error in settlement:", settlement.ID)
		return nil, err
	}
	tx, err := vdex.CancelMultiple(txOps, swapOrders)
	if err != nil {
		logrus.Error(txOps)
		return nil, err
	}
	return tx, nil
}

func (t *SettlementExecutor) submitExecuteOrderBookTrade(idx int, settlement *models.Settlement) (*types.Transaction, error) {
	ctx := context.Background()
	trade, err := service.GetTradeByID(ctx, settlement.RefID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	orders, err := service.GetOrdersByIDs(ctx, []int64{trade.MakerOrderID, trade.TakerOrderID})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	traders, err := service.GetUsersByIDs(ctx, []int64{orders[0].UserID, orders[1].UserID})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	product, err := service.GetProductByID(ctx, trade.ProductID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	vdexContract := common.HexToAddress(config.GetConfig().Contracts.Vdex)
	vdex, err := contract.NewVdex(vdexContract, t.ethClient)
	if err != nil {
		return nil, err
	}

	buyOrder := models.Order{}
	sellOrder := models.Order{}
	if orders[0].Side == models.SideBuy {
		buyOrder = orders[0]
		sellOrder = orders[1]

	} else {
		buyOrder = orders[1]
		sellOrder = orders[0]

	}
	makerOrder := orders[0]
	if trade.MakerOrderID == orders[1].ID {
		makerOrder = orders[1]
	}
	buy := contract.TypesOrderParam{
		Trader:           common.HexToAddress(util.GetUserByID(buyOrder.UserID, traders).Address),
		BaseAssetAmount:  util.ToWei(buyOrder.Size, 18),
		QuoteAssetAmount: util.ToWei(buyOrder.Funds, 18),
		OrderType:        util.OrderTypeToNumber(buyOrder.Type),
		Side:             util.OrderSideToNumber(buyOrder.Side),
		Expiration:       big.NewInt(buyOrder.Expiration),
		Nonce:            big.NewInt(buyOrder.Nonce),
		WalletSignature:  []byte("0x00"),
	}
	sell := contract.TypesOrderParam{
		Trader:           common.HexToAddress(util.GetUserByID(sellOrder.UserID, traders).Address),
		BaseAssetAmount:  util.ToWei(sellOrder.Size, 18),
		QuoteAssetAmount: util.ToWei(sellOrder.Funds, 18),
		OrderType:        util.OrderTypeToNumber(sellOrder.Type),
		Side:             util.OrderSideToNumber(sellOrder.Side),
		Expiration:       big.NewInt(sellOrder.Expiration),
		Nonce:            big.NewInt(sellOrder.Nonce),
		WalletSignature:  []byte("0x00"),
	}

	makerFeeAssetAddress := common.HexToAddress(config.GetAddressBySymbol(product.BaseCurrency))
	takerFeeAssetAddress := common.HexToAddress(config.GetAddressBySymbol(product.QuoteCurrency))
	// maker fee
	makerFeeQuantity := util.ToWei(settlement.MakerFee, 18)
	takerFeeQuantity := util.ToWei(settlement.TakerFee, 18)
	// grossQuantity
	grossBaseQuantity := util.ToWei(trade.Size, 18)
	grossQuoteQuantity := util.ToWei(trade.Price.Mul(trade.Size), 18)
	// netQuantity
	netBaseQuantity := big.NewInt(0).Sub(grossBaseQuantity, makerFeeQuantity)
	netQuoteQuantity := big.NewInt(0).Sub(grossQuoteQuantity, takerFeeQuantity)
	if makerOrder.Side == models.SideSell {
		makerFeeAssetAddress = common.HexToAddress(config.GetAddressBySymbol(product.QuoteCurrency))
		takerFeeAssetAddress = common.HexToAddress(config.GetAddressBySymbol(product.BaseCurrency))

		netBaseQuantity = big.NewInt(0).Sub(grossBaseQuantity, takerFeeQuantity)
		netQuoteQuantity = big.NewInt(0).Sub(grossQuoteQuantity, makerFeeQuantity)
	}
	orderbookTrade := contract.OrderBookTrade{
		BaseAssetAddress:     common.HexToAddress(config.GetAddressBySymbol(product.BaseCurrency)),
		QuoteAssetAddress:    common.HexToAddress(config.GetAddressBySymbol(product.QuoteCurrency)),
		GrossBaseQuantity:    grossBaseQuantity,
		GrossQuoteQuantity:   grossQuoteQuantity,
		NetBaseQuantity:      netBaseQuantity,
		NetQuoteQuantity:     netQuoteQuantity,
		MakerFeeAssetAddress: makerFeeAssetAddress,
		TakerFeeAssetAddress: takerFeeAssetAddress,
		MakerFeeQuantity:     makerFeeQuantity,
		TakerFeeQuantity:     takerFeeQuantity,
		Price:                util.ToWei(trade.Price, 18),
		MakerSide:            util.OrderSideToNumber(makerOrder.Side),
	}
	txOps, err := getTxOpsForExecuteOrderBookTrade(idx, t.ethClient, vdexContract, buy, sell, orderbookTrade)
	if err != nil {
		spew.Dump(buy, sell, orderbookTrade)
		logrus.Errorln("error in settlement:", settlement.ID)
		return nil, err
	}
	tx, err := vdex.ExecuteOrderBookTrade(txOps, buy, sell, orderbookTrade)
	if err != nil {
		logrus.Error(txOps)
		return nil, err
	}
	return tx, nil
}

func getTxOpsForCancelOrder(idx int, ethClient *ethclient.Client, vdexContract common.Address, swapOrders []contract.BaseOrderUtilsCancelOrderParams) (*bind.TransactOpts, error) {
	ctx := context.Background()
	privateKey, err := crypto.HexToECDSA(config.GetDispatcherByID(idx))
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Errorln(err)
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	vdexABI, err := abi.JSON(strings.NewReader(contract.VdexABI))
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	data, err := vdexABI.Pack("cancelMultiple", swapOrders)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	gasPrice, err := ethClient.SuggestGasPrice(ctx)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	callMsg := ethereum.CallMsg{
		From:     fromAddress,
		To:       &vdexContract,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		Data:     data,
	}
	gasLimit, err := ethClient.EstimateGas(ctx, callMsg)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	txOps, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.GetConfig().ChainID))
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	txOps.Nonce = big.NewInt(int64(nonce))
	txOps.Value = big.NewInt(0) // in wei
	txOps.GasLimit = gasLimit
	txOps.GasPrice = gasPrice

	return txOps, nil
}

func getTxOpsForExecuteOrderBookTrade(idx int, ethClient *ethclient.Client, vdexContract common.Address, buy contract.TypesOrderParam, sell contract.TypesOrderParam, orderbookTrade contract.OrderBookTrade) (*bind.TransactOpts, error) {
	ctx := context.Background()
	privateKey, err := crypto.HexToECDSA(config.GetDispatcherByID(idx))
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Errorln(err)
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	vdexABI, err := abi.JSON(strings.NewReader(contract.VdexABI))
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	data, err := vdexABI.Pack("executeOrderBookTrade", buy, sell, orderbookTrade)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	gasPrice, err := ethClient.SuggestGasPrice(ctx)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	callMsg := ethereum.CallMsg{
		From:     fromAddress,
		To:       &vdexContract,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		Data:     data,
	}
	gasLimit, err := ethClient.EstimateGas(ctx, callMsg)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	txOps, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.GetConfig().ChainID))
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	txOps.Nonce = big.NewInt(int64(nonce))
	txOps.Value = big.NewInt(0) // in wei
	txOps.GasLimit = gasLimit + gasLimit*20/100
	txOps.GasPrice = gasPrice

	return txOps, nil
}
