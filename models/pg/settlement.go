package pg

import (
	"context"
	"time"

	"github.com/cxptek/vdex/models"
)

func (s *Store) AddSettlements(ctx context.Context, settlements []models.Settlement) error {
	if len(settlements) == 0 {
		return nil
	}
	for i := range settlements {
		now := time.Now()
		settlements[i].CreatedAt = now
		settlements[i].UpdatedAt = now
	}
	_, err := s.db.NewInsert().Model(&settlements).Exec(ctx)

	return err
}

func (s *Store) AddSettlement(ctx context.Context, settlement *models.Settlement) error {
	settlement.CreatedAt = time.Now()
	settlement.UpdatedAt = time.Now()
	_, err := s.db.NewInsert().Model(settlement).Exec(ctx)

	return err
}

func (s *Store) GetUnsettledSettlements(ctx context.Context, count int32) ([]*models.Settlement, error) {
	settlements := []*models.Settlement{}
	if err := s.db.NewSelect().Model(&settlements).Where("settled = ?", false).
		Order("id ASC").Scan(ctx); err != nil {
		return nil, err
	}

	return settlements, nil
}

func (s *Store) UpdateSettlement(ctx context.Context, Settlement *models.Settlement) error {
	_, err := s.db.NewUpdate().Model(Settlement).OmitZero().WherePK().Exec(ctx)
	return err
}
