package service

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
)

func GetUnsettledSettlements(ctx context.Context, count int32) ([]*models.Settlement, error) {
	return pg.SharedStore().GetUnsettledSettlements(ctx, count)
}

func UpdateSettlement(ctx context.Context, tradeTransaction *models.Settlement) error {
	return pg.SharedStore().UpdateSettlement(ctx, tradeTransaction)
}
