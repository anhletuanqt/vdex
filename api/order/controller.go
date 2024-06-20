package order

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/models"
	redisClient "github.com/cxptek/vdex/redis"
	"github.com/cxptek/vdex/service"
	"github.com/cxptek/vdex/util"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func CreateOrder(c *fiber.Ctx) error {
	// user := service.GetUserFromContext(c)
	req := &CreateOrderReq{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	order := &models.Order{
		UserID:        req.UserID,
		ProductID:     req.ProductID,
		Side:          models.Side(req.Side),
		Price:         decimal.NewFromFloat(req.Price),
		Size:          decimal.NewFromFloat(req.Quantity),
		Funds:         decimal.NewFromFloat(req.Price * req.Quantity),
		Status:        models.OrderStatusNew,
		Type:          models.OrderType(req.Type),
		CreatedTxHash: fmt.Sprintf(fmt.Sprintf("tx-hash-%v", time.Now().Nanosecond())),
	}

	if err := service.AddOrder(c.Context(), order); err != nil {
		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}
	buf, err := json.Marshal(order)
	if err != nil {
		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	if err := getWriter(req.ProductID).WriteMessages(context.Background(), kafka.Message{Value: buf}); err != nil {
		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(order))
}

func CreateOrder1(c *fiber.Ctx) error {
	type TestCase struct {
		Price decimal.Decimal
		Size  decimal.Decimal
		User  int64
		Side  models.Side
	}
	testCases := []TestCase{
		{
			Price: decimal.NewFromFloat(1000),
			Size:  decimal.NewFromFloat(0.5),
			User:  2,
			Side:  models.SideSell,
		},
		{
			Price: decimal.NewFromFloat(1000),
			Size:  decimal.NewFromFloat(0.8),
			User:  2,
			Side:  models.SideSell,
		},
		{
			Price: decimal.NewFromFloat(930),
			Size:  decimal.NewFromFloat(0.3),
			User:  2,
			Side:  models.SideSell,
		},
		{
			Price: decimal.NewFromFloat(930),
			Size:  decimal.NewFromFloat(0.2),
			User:  2,
			Side:  models.SideSell,
		},
		{
			Price: decimal.NewFromFloat(1100),
			Size:  decimal.NewFromFloat(1.5),
			User:  3,
			Side:  models.SideBuy,
		},
	}
	for _, v := range testCases {
		pID := "BNB-VIC"
		order := &models.Order{
			UserID:        v.User,
			ProductID:     pID,
			Side:          v.Side,
			Price:         v.Price,
			Size:          v.Size,
			Funds:         v.Price.Mul(v.Size),
			Status:        models.OrderStatusNew,
			Type:          models.OrderTypeLimit,
			CreatedTxHash: fmt.Sprintf(fmt.Sprintf("tx-hash-%v", time.Now().Nanosecond())),
		}
		if err := service.AddOrder(c.Context(), order); err != nil {
			return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
		}
		buf, err := json.Marshal(order)
		if err != nil {
			return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
		}
		if err := getWriter(pID).WriteMessages(context.Background(), kafka.Message{Value: buf}); err != nil {
			return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
		}
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(""))
}

func GetOrders(c *fiber.Ctx) error {
	user, err := service.GetUserFromContext(c)
	if err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	query := &GetOrderQuery{}
	if err := c.QueryParser(query); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	statues := []models.OrderStatus{}
	for _, v := range strings.Split(query.Statuses, ",") {
		statues = append(statues, models.OrderStatus(v))
	}

	orders, err := service.GetOrdersByUserID(c.Context(), user.ID, query.ProductID, statues, query.Side, query.AfterID, query.Limit)
	if err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(orders))
}

func CancelOrder(c *fiber.Ctx) error {
	user, err := service.GetUserFromContext(c)
	if err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	params := &IDParams{}
	if err := c.ParamsParser(params); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}

	order, err := service.GetOrderByID(c.Context(), params.ID)
	if err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	if order.UserID != user.ID {
		return c.Status(200).JSON(util.ErrorResponse(400, "order not found"))
	}
	if order.Status != models.OrderStatusOpen {
		return c.Status(200).JSON(util.ErrorResponse(400, "invalid order status"))
	}

	order.Status = models.OrderStatusCancelling
	submitOrder(order)

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(order))
}

func CreateGaslessOrder(c *fiber.Ctx) error {
	user, err := service.GetUserFromContext(c)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	req := &CreateGaslessOrderReq{}
	if err := c.BodyParser(req); err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	size, err := decimal.NewFromString(req.Quantity)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	// verify permit and place order sig
	var result interface{}
	order := &models.Order{
		UserID:        user.ID,
		ProductID:     req.ProductID,
		Side:          models.Side(req.Side),
		Price:         price,
		Size:          size,
		Funds:         size.Mul(price),
		Status:        models.OrderStatusNew,
		Type:          models.OrderType(req.OrderType),
		CreatedTxHash: "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		TimeInForce:   "gtc",
		Expiration:    req.Expiration,
		Gasless:       true,
		Nonce:         0,
	}
	if req.Side == models.SideBuy.String() { // permit vic
		result = &DecodedDispatch{}
		err = decodePermit(req.EncodedPermit, "dispatch", result)
	} else {
		result = &DecodedPermit{}
		err = decodePermit(req.EncodedPermit, "permit", result)
	}
	if err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	if err := verifyOrder(result, *order); err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}

	// redis lock
	totalTx, err := redisClient.GetRedisClient().Incr(c.Context(), redisClient.Vdex_Total_Gasless_Tx_Key).Uint64()
	if err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	pkIndex := int(totalTx) % len(config.GetConfig().GaslessWallets)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	ctx, cancel := context.WithDeadline(c.Context(), time.Now().Add(time.Minute))
	defer cancel()
	lock, err := redisClient.GetRedisLockWithDeadline(ctx, fmt.Sprintf("%v_%v", redisClient.Vdex_Gasless_Executor, pkIndex))
	if err != nil {
		logrus.Errorln(err)
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	defer lock.Release(c.Context())
	order, err = service.AddGaslessOrderDirectly(c.Context(), pkIndex, order,
		hexutil.MustDecode(req.EncodedPermit),
		hexutil.MustDecode(req.PlaceOrderSig))
	if err != nil {
		logrus.Errorln(err)
		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(order))
}
