package service

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
	"github.com/cxptek/vdex/orderbook"
	"github.com/shopspring/decimal"
)

func GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	return pg.SharedStore().GetProductByID(ctx, id)
}

func GetProductsByIDs(ctx context.Context, ids []string) ([]models.Product, error) {
	return pg.SharedStore().GetProductsByIDs(ctx, ids)
}

func GetProducts(ctx context.Context) ([]models.Product, error) {
	return pg.SharedStore().GetProducts(ctx)
}

func GetProductDepth(productID string, limit int64) ([][]decimal.Decimal, [][]decimal.Decimal) {
	return orderbook.GetOrderBookByID(productID).DepthWithLimit(int(limit))
}
