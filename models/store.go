package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Store interface {
	BeginTx(ctx context.Context) (Store, error)
	Rollback() error
	CommitTx() error

	// Product
	GetProductByID(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context) ([]Product, error)
	GetProductsByIDs(ctx context.Context, ids []string) ([]Product, error)

	// Tick
	GetLastTickByProductID(ctx context.Context, productID string, granularity int64) (*Tick, error)
	AddTicks(ctx context.Context, ticks []*Tick) error
	GetTicksByProductID(ctx context.Context, productID string, granularity int64, limit int64) ([]*Tick, error)

	// Order
	AddOrder(ctx context.Context, order *Order) error
	AddOrderOnConflictTxHash(ctx context.Context, order *Order) error
	UpdateOrder(ctx context.Context, order *Order) error
	UpdateOrderFeeByID(ctx context.Context, id int64, fee decimal.Decimal) error
	UpdateOrderStatus(ctx context.Context, orderID int64, oldStatus, newStatus OrderStatus) (*Order, error)
	GetOrderByID(ctx context.Context, orderID int64) (*Order, error)
	GetOrderByIDForUpdate(ctx context.Context, orderID int64) (*Order, error)
	GetOrdersByUserID(ctx context.Context, userID int64, statuses []OrderStatus, side *Side, productID *string,
		afterID, limit int64) ([]*Order, error)
	GetOrdersByIDs(ctx context.Context, orderIDs []int64) ([]Order, error)
	UpdateOrderFees(ctx context.Context, orders []Order) error

	// Trade
	GetLastTradeByProductID(ctx context.Context, productID string) (*Trade, error)
	AddTrades(ctx context.Context, trades []*Trade) error
	GetTradesByProductID(ctx context.Context, productID string, limit int64) ([]Trade, error)
	GetTradeByID(ctx context.Context, tradeID int64) (*Trade, error)
	GetTradeByIDForUpdate(ctx context.Context, tradeID int64) (*Trade, error)

	// TradeTransaction
	AddSettlements(ctx context.Context, settlements []Settlement) error
	AddSettlement(ctx context.Context, settlement *Settlement) error
	GetUnsettledSettlements(ctx context.Context, count int32) ([]*Settlement, error)
	UpdateSettlement(ctx context.Context, tradeTransaction *Settlement) error

	// Fill
	GetLastFillByProductID(ctx context.Context, productID string) (*Fill, error)
	AddFills(ctx context.Context, fills []*Fill) error
	GetUnsettledFills(ctx context.Context, count int32) ([]*Fill, error)
	GetUnsettledFillsByOrderID(ctx context.Context, orderID int64) ([]*Fill, error)
	UpdateFill(ctx context.Context, fill *Fill) error
	GetFilledByUserID(ctx context.Context, userID int64, limit int64) ([]*Fill, error)

	// User
	AddUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, userID int64) (*User, error)
	GetUserByPublicID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUsersByIDs(ctx context.Context, userID []int64) ([]User, error)

	// WorkerConfig
	AddWorkerConfig(ctx context.Context, wk *WorkerConfig) error
	UpdateWorkerConfig(ctx context.Context, wk *WorkerConfig) error
	GetWorkerConfig(ctx context.Context) (*WorkerConfig, error)

	// UserSegment
	GetLastUserSegmentByUserIDForUpdate(ctx context.Context, userID int64) (*UserSegment, error)
	GetLastUserSegment(ctx context.Context) (*UserSegment, error)
	AddUserSegments(ctx context.Context, segments []*UserSegment) error
	AddUserSegment(ctx context.Context, segments *UserSegment) (*UserSegment, error)
	GetUserSegmentByUserIDAndTime(ctx context.Context, userID int64, time int64) (*UserSegment, error)
}
