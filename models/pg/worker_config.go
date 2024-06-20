package pg

import (
	"context"

	"github.com/cxptek/vdex/models"
)

func (s *Store) AddWorkerConfig(ctx context.Context, wk *models.WorkerConfig) error {
	_, err := s.db.NewInsert().Model(wk).Exec(ctx)
	return err
}

func (s *Store) UpdateWorkerConfig(ctx context.Context, wk *models.WorkerConfig) error {
	_, err := s.db.NewUpdate().Model(wk).OmitZero().WherePK().Exec(ctx)
	return err
}

func (s *Store) GetWorkerConfig(ctx context.Context) (*models.WorkerConfig, error) {
	wk := &models.WorkerConfig{}
	err := s.db.NewSelect().Model(wk).Limit(1).Scan(ctx)
	return wk, err
}
