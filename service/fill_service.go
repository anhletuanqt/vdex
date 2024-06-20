package service

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
)

func AddFills(ctx context.Context, fills []*models.Fill) error {
	if len(fills) == 0 {
		return nil
	}
	return pg.SharedStore().AddFills(ctx, fills)
}

func GetUnsettledFills(ctx context.Context, count int32) ([]*models.Fill, error) {
	return pg.SharedStore().GetUnsettledFills(ctx, count)
}

func GetFilledByUserID(ctx context.Context, userID int64, limit int64) ([]*models.Fill, error) {
	return pg.SharedStore().GetFilledByUserID(ctx, userID, limit)
}
