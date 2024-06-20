package pg

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/uptrace/bun"
)

func (s *Store) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	product := &models.Product{}
	if err := s.db.NewSelect().Model(product).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) GetProductsByIDs(ctx context.Context, ids []string) ([]models.Product, error) {
	product := []models.Product{}
	if err := s.db.NewSelect().Model(&product).Where("id IN (?)", bun.In(ids)).Scan(ctx); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) GetProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	if err := s.db.NewSelect().Model(&products).Scan(ctx); err != nil {
		return nil, err
	}

	return products, nil
}
