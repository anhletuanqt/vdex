package product

import (
	"time"

	"github.com/cxptek/vdex/models"
	"github.com/shopspring/decimal"
)

type IDParams struct {
	ID string `params:"id"`
}

type SubmitOrderBody struct {
	ID        int64   `json:"id"`
	Side      string  `json:"side"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Type      string  `json:"type"`
	ProductID string  `json:"productId"`
	IsCancel  bool    `json:"isCancel"`
}

type GetTradesQuery struct {
	Limit int64 `query:"limit"`
}

type GetDepthQuery struct {
	Limit int64 `query:"limit"`
}

type GetCandlesQuery struct {
	Granularity int64 `query:"granularity"`
	Limit       int64 `query:"limit"`
}

type ProductRes struct {
	ID             string          `json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	BaseCurrency   string          `json:"baseCurrency"`
	QuoteCurrency  string          `json:"quoteCurrency"`
	BaseMinSize    decimal.Decimal `json:"baseMinSize"`
	BaseMaxSize    decimal.Decimal `json:"baseMaxSize"`
	QuoteMinSize   decimal.Decimal `json:"quoteMinSize"`
	QuoteMaxSize   decimal.Decimal `json:"quoteMaxSize"`
	BaseScale      int32           `json:"baseScale"`
	QuoteScale     int32           `json:"quoteScale"`
	QuoteIncrement float64         `json:"quoteIncrement"`
}

func NewProductsRes(products []models.Product) []ProductRes {
	res := make([]ProductRes, len(products))
	for i := range products {
		res[i] = ProductRes(products[i])
	}

	return res
}
