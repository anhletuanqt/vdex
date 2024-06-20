package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Side string
type DoneReason string
type SettlementType string

const (
	OrderTypeLimit  = OrderType("limit")
	OrderTypeMarket = OrderType("market")

	SideBuy  = Side("buy")
	SideSell = Side("sell")

	OrderStatusNew        = OrderStatus("new")
	OrderStatusOpen       = OrderStatus("open")
	OrderStatusCancelling = OrderStatus("cancelling")
	OrderStatusCancelled  = OrderStatus("cancelled")
	OrderStatusFilled     = OrderStatus("filled")

	DoneReasonFilled    = DoneReason("filled")
	DoneReasonCancelled = DoneReason("cancelled")
	SettlementTrade     = SettlementType("trade")
	SettlementOrder     = SettlementType("order")
)

func NewSideFromString(s string) (*Side, error) {
	side := Side(s)
	switch side {
	case SideBuy:
	case SideSell:
	default:
		return nil, fmt.Errorf("invalid side: %v", s)
	}
	return &side, nil
}

func (s Side) Opposite() Side {
	if s == SideBuy {
		return SideSell
	}
	return SideBuy
}

func (s Side) String() string {
	return string(s)
}

type OrderType string

func (t OrderType) String() string {
	return string(t)
}

func NewOrderTypeFromOnChain(v int64) OrderType {
	if v == 0 {
		return OrderTypeMarket
	}

	return OrderTypeLimit
}

type OrderStatus string

func NewOrderStatusFromString(s string) (*OrderStatus, error) {
	status := OrderStatus(s)
	switch status {
	case OrderStatusNew:
	case OrderStatusOpen:
	case OrderStatusCancelling:
	case OrderStatusCancelled:
	case OrderStatusFilled:
	default:
		return nil, fmt.Errorf("invalid status: %v", s)
	}
	return &status, nil
}

func (t OrderStatus) String() string {
	return string(t)
}

// ************************************
// ************** TABLES **************
// ************************************
type User struct {
	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	PublicID  uuid.UUID `json:"publicId" bun:"default:gen_random_uuid()"`
}
type Product struct {
	ID             string          `json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	BaseCurrency   string          `json:"baseCurrency"`
	QuoteCurrency  string          `json:"quoteCurrency"`
	BaseMinSize    decimal.Decimal `json:"baseMinSize" sql:"type:decimal(32,16);"`
	BaseMaxSize    decimal.Decimal `json:"baseMaxSize" sql:"type:decimal(32,16);"`
	QuoteMinSize   decimal.Decimal `json:"quoteMinSize" sql:"type:decimal(32,16);"`
	QuoteMaxSize   decimal.Decimal `json:"quoteMaxSize" sql:"type:decimal(32,16);"`
	BaseScale      int32           `json:"baseScale"`
	QuoteScale     int32           `json:"quoteScale"`
	QuoteIncrement float64         `json:"quoteIncrement"`
}

type Order struct {
	ID            int64           `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	ProductID     string          `json:"quoteIncrement"`
	UserID        int64           `json:"userId"`
	Size          decimal.Decimal `json:"size" sql:"type:decimal(32,16);"`
	Funds         decimal.Decimal `json:"funds" sql:"type:decimal(32,16);"`
	FilledSize    decimal.Decimal `json:"filledSize" sql:"type:decimal(32,16);"`
	ExecutedValue decimal.Decimal `json:"executedValue" sql:"type:decimal(32,16);"`
	Price         decimal.Decimal `json:"price" sql:"type:decimal(32,16);"`
	FillFees      decimal.Decimal `json:"fillFees" sql:"type:decimal(32,16);"`
	Type          OrderType       `json:"type"`
	Side          Side            `json:"side"`
	TimeInForce   string          `json:"timeInForce"`
	Status        OrderStatus     `json:"status"`
	Settled       bool            `json:"settled"`
	Nonce         int64           `json:"nonce"`
	Expiration    int64           `json:"expiration"`
	CreatedTxHash string          `json:"createdTxHash"`
	PublicID      uuid.UUID       `json:"publicId" bun:"default:gen_random_uuid()"`
	Gasless       bool            `json:"gasless"`
}

type WorkerConfig struct {
	ID          uint64 `json:"id" bun:"id,pk,autoincrement"`
	BlockNumber uint64 `json:"blockNumber"`
}
type Trade struct {
	ID           int64           `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	Time         *time.Time      `json:"time"`
	ProductID    string          `json:"productId"`
	TakerOrderID int64           `json:"takerOrderId"`
	MakerOrderID int64           `json:"makerOrderId"`
	Price        decimal.Decimal `json:"price" sql:"type:decimal(32,16);"`
	Size         decimal.Decimal `json:"size" sql:"type:decimal(32,16);"`
	Side         Side            `json:"side"`
	LogOffset    int64           `json:"logOffset"`
	PublicID     uuid.UUID       `json:"publicId" bun:"default:gen_random_uuid()"`
}

type Tick struct {
	ID          int64           `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	ProductID   string          `json:"productId"`
	Granularity int64           `json:"granularity"`
	Time        int64           `json:"time"`
	Open        decimal.Decimal `json:"open" sql:"type:decimal(32,16);"`
	High        decimal.Decimal `json:"high" sql:"type:decimal(32,16);"`
	Low         decimal.Decimal `json:"low" sql:"type:decimal(32,16);"`
	Close       decimal.Decimal `json:"close" sql:"type:decimal(32,16);"`
	Volume      decimal.Decimal `json:"volume" sql:"type:decimal(32,16);"`
	LogOffset   int64           `json:"logOffset"`
}

type Fill struct {
	ID         int64           `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	UserID     int64           `json:"userId"`
	OrderID    int64           `json:"orderId"`
	ProductID  string          `json:"productId"`
	Size       decimal.Decimal `json:"size" sql:"type:decimal(32,16);"`
	Price      decimal.Decimal `json:"price" sql:"type:decimal(32,16);"`
	Funds      decimal.Decimal `json:"funds" sql:"type:decimal(32,16);"`
	Fee        decimal.Decimal `json:"fee" sql:"type:decimal(32,16);"`
	Settled    bool            `json:"settled"`
	Side       Side            `json:"side"`
	DoneReason DoneReason      `json:"doneReason"`
	LogOffset  int64           `json:"logOffset"`
	PublicID   uuid.UUID       `json:"publicId" bun:"default:gen_random_uuid()"`
	Liquidity  string          `json:"liquidity"`
}

type Settlement struct {
	ID        int64           `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	RefID     int64           `json:"refId"`
	Type      SettlementType  `json:"type"`
	TxHash    string          `json:"txHash"`
	Settled   bool            `json:"settled"`
	TakerFee  decimal.Decimal `json:"takerFee" sql:"type:decimal(32,16);"`
	MakerFee  decimal.Decimal `json:"makerFee" sql:"type:decimal(32,16);"`
	Time      time.Time       `json:"time"`
}

type UserSegment struct {
	ID        int64           `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	UserID    int64           `json:"userId"`
	Time      int64           `json:"time"`
	Type      int64           `json:"type"`
	Volume    decimal.Decimal `json:"volume" sql:"type:decimal(32,16);"`
	LogOffset int64           `json:"logOffset"`
}
