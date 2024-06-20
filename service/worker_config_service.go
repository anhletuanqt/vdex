package service

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
)

func AddWorkerConfig(ctx context.Context, wk *models.WorkerConfig) error {
	return pg.SharedStore().AddWorkerConfig(ctx, wk)
}

func UpdateWorkerConfig(ctx context.Context, wk *models.WorkerConfig) error {
	return pg.SharedStore().UpdateWorkerConfig(ctx, wk)
}

func GetWorkerConfig(ctx context.Context) (*models.WorkerConfig, error) {
	return pg.SharedStore().GetWorkerConfig(ctx)
}
