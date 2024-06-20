package service

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
	"github.com/cxptek/vdex/models/pg"
	"github.com/cxptek/vdex/util"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func AddOrder(ctx context.Context, order *models.Order) error {
	return pg.SharedStore().AddOrder(ctx, order)
}
func AddOrderOnConflictTxHash(ctx context.Context, order *models.Order) error {
	return pg.SharedStore().AddOrderOnConflictTxHash(ctx, order)
}

func UpdateOrder(ctx context.Context, orderID int64, oldStatus, newStatus models.OrderStatus) (*models.Order, error) {
	return pg.SharedStore().UpdateOrderStatus(ctx, orderID, oldStatus, newStatus)
}

func GetOrderByID(ctx context.Context, orderID int64) (*models.Order, error) {
	return pg.SharedStore().GetOrderByID(ctx, orderID)
}

func GetOrdersByIDs(ctx context.Context, orderIDs []int64) ([]models.Order, error) {
	return pg.SharedStore().GetOrdersByIDs(ctx, orderIDs)
}

func ExecuteFill(ctx context.Context, orderID int64) error {
	// tx
	db, err := pg.SharedStore().BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	order, err := db.GetOrderByIDForUpdate(ctx, orderID)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	if order == nil {
		return fmt.Errorf("order not found: %v", orderID)
	}
	if order.Status == models.OrderStatusFilled || order.Status == models.OrderStatusCancelled {
		return fmt.Errorf("order status invalid: %v %v", orderID, order.Status)
	}
	// if _, err := GetProductByID(ctx, order.ProductID); err != nil {
	// 	logrus.Errorln(err)
	// 	return err
	// }

	fills, err := pg.SharedStore().GetUnsettledFillsByOrderID(ctx, orderID)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	if len(fills) == 0 {
		return nil
	}

	for _, fill := range fills {
		fill.Settled = true

		if fill.DoneReason == models.DoneReasonCancelled {
			order.Status = models.OrderStatusCancelled
			order.Settled = true
			settlement := &models.Settlement{
				RefID:   fill.OrderID,
				Settled: false,
				Type:    models.SettlementOrder,
				Time:    time.Now(),
			}
			if err := db.AddSettlement(ctx, settlement); err != nil {
				logrus.Errorln(err)
				return err
			}
		} else if fill.DoneReason == models.DoneReasonFilled {
			executedValue := fill.Size.Mul(fill.Price)
			order.ExecutedValue = order.ExecutedValue.Add(executedValue)
			order.FilledSize = order.FilledSize.Add(fill.Size)
			if order.Type == models.OrderTypeMarket {
				order.Status = models.OrderStatusFilled
				order.Settled = true
				continue
			}
			// order type limit
			if order.FilledSize.Equal(order.Size) {
				order.Status = models.OrderStatusFilled
				order.Settled = true
			}
			if order.Status == models.OrderStatusNew {
				order.Status = models.OrderStatusOpen
			}
		} else {
			logrus.Fatalf("unknown done reason: %v", fill.DoneReason)
		}
	}

	err = db.UpdateOrder(ctx, order)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	for _, fill := range fills {
		if err := db.UpdateFill(ctx, fill); err != nil {
			logrus.Errorln(err)
			return err
		}
	}

	return db.CommitTx()
}

func GetOrdersByUserID(ctx context.Context, userID int64, productID *string, statuses []models.OrderStatus, side *models.Side, afterID, limit int64) ([]*models.Order, error) {
	return pg.SharedStore().GetOrdersByUserID(ctx, userID, statuses, side, productID, afterID, limit)
}

func AddGaslessOrderDirectly(ctx context.Context, pkIndex int, order *models.Order, permit []byte, placeOrderSig []byte) (*models.Order, error) {
	db, err := pg.SharedStore().BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = db.Rollback() }()

	user, err := db.GetUserByID(ctx, order.UserID)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	currencies := strings.Split(order.ProductID, "-")
	ethClient, err := ethclient.DialContext(ctx, config.GetConfig().JsonRpc)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	// call tx
	vdexContract := common.HexToAddress(config.GetConfig().Contracts.Vdex)
	vdex, err := contract.NewVdex(vdexContract, ethClient)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	contractOrder := contract.Order{
		Trader:           common.HexToAddress(user.Address),
		BaseAsset:        common.HexToAddress(config.GetAddressBySymbol(currencies[0])),
		QuoteAsset:       common.HexToAddress(config.GetAddressBySymbol(currencies[1])),
		BaseAssetAmount:  util.ToWei(order.Size, 18),
		QuoteAssetAmount: order.Funds.BigInt(),
		OrderType:        util.OrderTypeToNumber(order.Type),
		Side:             util.OrderSideToNumber(order.Side),
		Expiration:       big.NewInt(order.Expiration),
		Nonce:            big.NewInt(order.Nonce),
	}
	txOps, err := getTxOpsForCreateOrderWithPermit(pkIndex, ethClient, vdexContract, contractOrder, placeOrderSig, permit)
	if err != nil {
		logrus.Errorln(err)
		spew.Dump(contractOrder)
		return nil, err
	}
	tx, err := vdex.CreateOrderWithPermit(txOps, contractOrder, placeOrderSig, permit)
	if err != nil {
		return nil, err
	}
	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		logrus.Errorln("gasless tx failed ", tx.Hash().Hex())
		return nil, err
	}

	order.CreatedTxHash = strings.ToLower(tx.Hash().Hex())
	if err := db.AddOrder(ctx, order); err != nil {
		return nil, err
	}
	if err := db.CommitTx(); err != nil {
		return nil, err
	}

	return order, nil
}

func getTxOpsForCreateOrderWithPermit(idx int, ethClient *ethclient.Client, vdexContract common.Address, order contract.Order, signature []byte, permit []byte) (*bind.TransactOpts, error) {
	ctx := context.Background()
	privateKey, err := crypto.HexToECDSA(config.GetGaslessDispatcherByID(idx))
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

	data, err := vdexABI.Pack("createOrderWithPermit", order, signature, permit)
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
