package pg

import (
	"context"
	"time"

	"github.com/cxptek/vdex/models"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

func (s *Store) AddOrder(ctx context.Context, order *models.Order) error {
	_, err := s.db.NewInsert().Model(order).Exec(ctx)
	return err
}
func (s *Store) AddOrderOnConflictTxHash(ctx context.Context, order *models.Order) error {
	// _, err := s.db.NewInsert().Model(order).Exec(ctx)
	_, err := s.db.NewInsert().
		Model(order).
		On("CONFLICT (created_tx_hash) DO UPDATE").
		Set("updated_at = EXCLUDED.updated_at,nonce = EXCLUDED.nonce").
		Exec(ctx)
	return err
}

func (s *Store) UpdateOrder(ctx context.Context, order *models.Order) error {
	order.UpdatedAt = time.Now()
	_, err := s.db.NewUpdate().Model(order).OmitZero().WherePK().Exec(ctx)
	return err
}

// TODO: apply bulk create
func (s *Store) UpdateOrderFees(ctx context.Context, orders []models.Order) error {
	for i := range orders {
		orders[i].UpdatedAt = time.Now()
	}
	// _, err := s.db.NewUpdate().Model(&orders).OmitZero().Exec(ctx)
	_, err := s.db.NewUpdate().
		Model(&orders).
		Column("fill_fees", "updated_at").
		Bulk().
		Exec(ctx)
	return err
}

func (s *Store) UpdateOrderStatus(ctx context.Context, orderID int64, oldStatus, newStatus models.OrderStatus) (*models.Order, error) {
	order := &models.Order{
		Status:    newStatus,
		UpdatedAt: time.Now(),
	}
	_, err := s.db.NewUpdate().Model(order).Column("status").Where("id = ? AND status = ?", orderID, oldStatus).Returning("*").Exec(ctx)

	return order, err
}
func (s *Store) UpdateOrderFeeByID(ctx context.Context, id int64, fee decimal.Decimal) error {
	_, err := s.db.NewUpdate().Model((*models.Order)(nil)).Set("fill_fees = fill_fees + ?", fee).Where("id = ?", id).Exec(ctx)

	return err
}

func (s *Store) GetOrderByID(ctx context.Context, orderID int64) (*models.Order, error) {
	order := &models.Order{}
	if err := s.db.NewSelect().Model(order).Where("id = ?", orderID).Scan(ctx); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Store) GetOrdersByIDs(ctx context.Context, orderIDs []int64) ([]models.Order, error) {
	orders := []models.Order{}
	if err := s.db.NewSelect().Model(&orders).Where("id IN (?)", bun.In(orderIDs)).Scan(ctx); err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Store) GetOrderByIDForUpdate(ctx context.Context, orderID int64) (*models.Order, error) {
	order := &models.Order{}
	if err := s.db.NewSelect().Model(order).Where("id = ?", orderID).For("UPDATE").Scan(ctx); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Store) GetOrdersByUserID(ctx context.Context, userID int64, statuses []models.OrderStatus, side *models.Side, productID *string,
	afterID, limit int64) ([]*models.Order, error) {
	orders := []*models.Order{}
	db := s.db.NewSelect().Model(&orders).Where("user_id = ?", userID)
	if len(statuses) != 0 {
		db = db.Where("status IN (?)", bun.In(statuses))
	}
	if productID != nil {
		db = db.Where("product_id = ?", productID)
	}
	if side != nil {
		db = db.Where("side = ?", side)
	}
	if afterID > 0 {
		db = db.Where("id < ?", afterID)
	}
	if limit > 0 {
		db.Limit(int(limit))
	}

	db = db.Order("id DESC")
	err := db.Scan(ctx)

	return orders, err
}
