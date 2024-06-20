package matching

import (
	"time"

	"github.com/cxptek/vdex/orderbook"
	"github.com/shopspring/decimal"
)

type LogType string

const (
	LogTypePartial = LogType("partial")
	LogTypeDone    = LogType("done")
	LogTypeOpen    = LogType("open")
	LogTypeCancel  = LogType("cancel")
)

type MatchedOrder struct {
	open                     *orderbook.Order
	cancel                   *orderbook.Order
	done                     []*orderbook.MatchedLog
	partial                  *orderbook.Order
	partialQuantityProcessed decimal.Decimal
	quantityLeft             decimal.Decimal
}

type LogOrder struct {
	LogType      LogType
	OrderID      int64
	ProductID    string
	Quantity     decimal.Decimal
	Price        decimal.Decimal
	Side         orderbook.Side
	Time         time.Time
	TakerOrderID int64
	MakerOrderID int64
}
