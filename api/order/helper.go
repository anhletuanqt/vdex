package order

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/umbracle/ethgo/abi"
)

type DecodedPermit struct {
	Token    common.Address
	Owner    common.Address
	Spender  common.Address
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        []byte
	S        []byte
}

type DecodedDispatch struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Deadline  *big.Int
	V         uint8
	R         []byte
	S         []byte
}

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

func decodePermit(encoded string, encodedType string, result interface{}) error {
	permitType := abi.MustNewType("tuple(address token,address owner,address spender,uint256 value,uint256 deadline,uint8 v,bytes32 r,bytes32 s)")
	dispatchType := abi.MustNewType("tuple(address sender,address recipient,uint256 amount,uint256 deadline,uint8 v,bytes32 r,bytes32 s)")

	var err error = nil
	if encodedType == "permit" {
		err = permitType.DecodeStruct(hexutil.MustDecode(encoded), result)
	} else if encodedType == "dispatch" {
		err = dispatchType.DecodeStruct(hexutil.MustDecode(encoded), result)
	}

	return err
}

func (d *DecodedDispatch) VerifyOrder(amount *big.Int) error {
	if d.Amount.Cmp(amount) != 0 {
		return errors.New("ERR_INVALID_ORDER_AMOUNT")
	}

	return nil
}
func (d *DecodedPermit) VerifyOrder(amount *big.Int) error {
	if d.Value.Cmp(amount) != 0 {
		return errors.New("ERR_INVALID_ORDER_AMOUNT")
	}

	return nil
}

func verifyOrder(decodedPermit interface{}, order models.Order) error {
	ctx := context.Background()
	product, err := service.GetProductByID(ctx, order.ProductID)
	if err != nil {
		return err
	}

	if order.Expiration < time.Now().Unix() {
		return errors.New("ERR_GASLESS_EXPIRED")
	}

	switch v := decodedPermit.(type) {
	case *DecodedDispatch:
		if v.Amount.Cmp(order.Funds.BigInt()) != 0 {
			err = errors.New("INVALID_GASLESS_AMOUNT")
		}
	case *DecodedPermit:
		if v.Value.Cmp(order.Size.BigInt()) != 0 {
			err = errors.New("INVALID_GASLESS_AMOUNT")
		}
		if v.Token.Cmp(common.HexToAddress(config.GetAddressBySymbol(product.BaseCurrency))) != 0 {
			err = errors.New("INVALID_GASLESS_TOKEN")
		}
	default:
		err = errors.New("ERR_INVALID_TYPE")
	}

	return nil
}
