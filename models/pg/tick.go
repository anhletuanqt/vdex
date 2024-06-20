package pg

import (
	"context"
	"time"

	"github.com/cxptek/vdex/models"
)

func (s *Store) GetLastTickByProductID(ctx context.Context, productID string, granularity int64) (*models.Tick, error) {
	tick := &models.Tick{}
	if err := s.db.NewSelect().Model(tick).Where("product_id = ? AND granularity = ?", productID, granularity).
		Order("time DESC").Limit(1).Scan(ctx); err != nil {
		return nil, err
	}

	return tick, nil
}

func (s *Store) AddTicks(ctx context.Context, ticks []*models.Tick) error {
	if len(ticks) == 0 {
		return nil
	}

	for i := range ticks {
		now := time.Now()
		ticks[i].CreatedAt = now
		ticks[i].UpdatedAt = now
	}

	_, err := s.db.NewInsert().
		Model(&ticks).
		On("CONFLICT (product_id,granularity,time) DO UPDATE").
		Set("updated_at = EXCLUDED.updated_at,close = EXCLUDED.close,low = EXCLUDED.low" + "," +
			"high = EXCLUDED.high,volume = EXCLUDED.volume,log_offset = EXCLUDED.log_offset").
		Exec(ctx)

	return err
}

func (s *Store) GetTicksByProductID(ctx context.Context, productID string, granularity int64, limit int64) ([]*models.Tick, error) {
	ticks := []*models.Tick{}
	if err := s.db.NewSelect().Model(&ticks).Where("product_id = ? AND granularity = ?", productID, granularity).
		Order("time DESC").Limit(int(limit)).Scan(ctx); err != nil {
		return nil, err
	}

	return ticks, nil
}
