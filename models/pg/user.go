package pg

import (
	"context"
	"time"

	"github.com/cxptek/vdex/models"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (s *Store) AddUser(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := s.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (s *Store) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	user := &models.User{}
	if err := s.db.NewSelect().Model(user).Where("id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByPublicID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	if err := s.db.NewSelect().Model(user).Where("public_id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByAddress(ctx context.Context, address string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.NewSelect().Model(user).Where("address = ?", address).Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUsersByIDs(ctx context.Context, userIDs []int64) ([]models.User, error) {
	user := []models.User{}
	if err := s.db.NewSelect().Model(&user).Where("id IN (?)", bun.In(userIDs)).Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}
