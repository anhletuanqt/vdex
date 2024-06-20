package pg

import (
	"context"
	"time"

	"github.com/cxptek/vdex/models"
)

func (s *Store) AddTrades(ctx context.Context, trades []*models.Trade) error {
	if len(trades) == 0 {
		return nil
	}
	for i := range trades {
		now := time.Now()
		trades[i].CreatedAt = now
		trades[i].UpdatedAt = now
	}
	_, err := s.db.NewInsert().Model(&trades).Returning("*").Exec(ctx)

	return err
}

func (s *Store) GetLastTradeByProductID(ctx context.Context, productID string) (*models.Trade, error) {
	trade := &models.Trade{}
	err := s.db.NewSelect().Model(trade).Where("product_id = ?", productID).
		Order("id DESC").Limit(1).Scan(ctx)

	return trade, err
}

func (s *Store) GetTradesByProductID(ctx context.Context, productID string, limit int64) ([]models.Trade, error) {
	trades := []models.Trade{}
	if err := s.db.NewSelect().Model(&trades).Where("product_id = ?", productID).
		Order("id DESC").Limit(int(limit)).Scan(ctx); err != nil {
		return nil, err
	}

	return trades, nil
}

func (s *Store) GetTradeByID(ctx context.Context, tradeID int64) (*models.Trade, error) {
	trade := &models.Trade{}
	if err := s.db.NewSelect().Model(trade).Where("id = ?", tradeID).Scan(ctx); err != nil {
		return nil, err
	}

	return trade, nil
}

func (s *Store) GetTradeByIDForUpdate(ctx context.Context, tradeID int64) (*models.Trade, error) {
	trade := &models.Trade{}
	if err := s.db.NewSelect().Model(trade).Where("id = ?", tradeID).For("UPDATE").Scan(ctx); err != nil {
		return nil, err
	}

	return trade, nil
}
