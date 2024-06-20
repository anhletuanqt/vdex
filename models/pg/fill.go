package pg

import (
	"context"
	"time"

	"github.com/cxptek/vdex/models"
)

func (s *Store) AddFills(ctx context.Context, fills []*models.Fill) error {
	if len(fills) == 0 {
		return nil
	}
	for i := range fills {
		now := time.Now()
		fills[i].CreatedAt = now
		fills[i].UpdatedAt = now
	}
	_, err := s.db.NewInsert().Model(&fills).Exec(ctx)

	return err
}

func (s *Store) GetLastFillByProductID(ctx context.Context, productID string) (*models.Fill, error) {
	fill := &models.Fill{}
	err := s.db.NewSelect().Model(fill).Where("product_id = ?", productID).
		Order("id DESC").Limit(1).Scan(ctx)

	return fill, err
}

func (s *Store) GetUnsettledFills(ctx context.Context, count int32) ([]*models.Fill, error) {
	fills := []*models.Fill{}
	if err := s.db.NewSelect().Model(&fills).Where("settled = ?", false).
		Order("id ASC").Scan(ctx); err != nil {
		return nil, err
	}

	return fills, nil
}

func (s *Store) GetUnsettledFillsByOrderID(ctx context.Context, orderID int64) ([]*models.Fill, error) {
	fills := []*models.Fill{}

	if err := s.db.NewSelect().Model(&fills).Where("settled = ? AND order_id = ?", false, orderID).
		Order("id ASC").Scan(ctx); err != nil {
		return nil, err
	}

	return fills, nil
}

func (s *Store) UpdateFill(ctx context.Context, fill *models.Fill) error {
	_, err := s.db.NewUpdate().Model(fill).OmitZero().WherePK().Exec(ctx)
	return err
}

func (s *Store) GetFilledByUserID(ctx context.Context, userID int64, limit int64) ([]*models.Fill, error) {
	fills := []*models.Fill{}

	if err := s.db.NewSelect().Model(&fills).Where("user_id = ? AND done_reason = ?", userID, models.DoneReasonFilled).
		Order("id DESC").Limit(int(limit)).Scan(ctx); err != nil {
		return nil, err
	}

	return fills, nil
}
