package order

import "github.com/cxptek/vdex/models"

type IDParams struct {
	ID int64 `params:"id"`
}

type CreateOrderReq struct {
	Side      string  `json:"side"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Type      string  `json:"type"`
	ProductID string  `json:"productId"`
	UserID    int64   `json:"userId"`
}

type CreateGaslessOrderReq struct {
	EncodedPermit string `json:"encodedPermit"`
	PlaceOrderSig string `json:"placeOrderSig"`
	ProductID     string `json:"productId"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	OrderType     string `json:"orderType"`
	Side          string `json:"side"`
	Expiration    int64  `json:"expiration"`
}

type GetOrderQuery struct {
	ProductID *string      `query:"productId"`
	Limit     int64        `query:"limit"`
	Statuses  string       `query:"statuses"`
	AfterID   int64        `query:"afterId"`
	Side      *models.Side `query:"side"`
}

type CancelOrdersQuery struct {
	ProductID *string      `query:"productId"`
	Side      *models.Side `query:"side"`
}

type FilledHistoryQuery struct {
	Limit int64 `query:"limit"`
}
