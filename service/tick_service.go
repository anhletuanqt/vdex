package service

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
)

func GetLastTickByProductID(ctx context.Context, productId string, granularity int64) (*models.Tick, error) {
	return pg.SharedStore().GetLastTickByProductID(ctx, productId, granularity)
}

func AddTicks(ctx context.Context, ticks []*models.Tick) error {
	if len(ticks) == 0 {
		return nil
	}
	return pg.SharedStore().AddTicks(ctx, ticks)
}

func GetTicksByProductID(ctx context.Context, productID string, granularity int64, limit int64) ([]*models.Tick, error) {
	return pg.SharedStore().GetTicksByProductID(ctx, productID, granularity, limit)
}
