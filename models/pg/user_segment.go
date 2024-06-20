package pg

import (
	"context"
	"time"

	"github.com/cxptek/vdex/models"
)

func (s *Store) GetLastUserSegmentByUserIDForUpdate(ctx context.Context, userID int64) (*models.UserSegment, error) {
	segment := &models.UserSegment{}
	if err := s.db.NewSelect().Model(segment).Where("user_id = ?", userID).
		Order("time DESC").Limit(1).For("UPDATE").Scan(ctx); err != nil {
		return nil, err
	}

	return segment, nil
}

func (s *Store) GetLastUserSegment(ctx context.Context) (*models.UserSegment, error) {
	segment := &models.UserSegment{}
	if err := s.db.NewSelect().Model(segment).
		Order("time DESC").Limit(1).Scan(ctx); err != nil {
		return nil, err
	}

	return segment, nil
}

func (s *Store) GetUserSegmentByUserIDAndTime(ctx context.Context, userID int64, time int64) (*models.UserSegment, error) {
	segment := &models.UserSegment{}
	if err := s.db.NewSelect().Model(segment).Where("user_id = ? AND time = ?", userID, time).
		Limit(1).Scan(ctx); err != nil {
		return nil, err
	}

	return segment, nil
}

func (s *Store) AddUserSegments(ctx context.Context, segments []*models.UserSegment) error {
	if len(segments) == 0 {
		return nil
	}

	for i := range segments {
		now := time.Now()
		segments[i].CreatedAt = now
		segments[i].UpdatedAt = now
	}

	_, err := s.db.NewInsert().
		Model(&segments).
		On("CONFLICT (user_id,time) DO UPDATE").
		Set("updated_at = EXCLUDED.updated_at,volume = EXCLUDED.volume,type = EXCLUDED.type,log_offset = EXCLUDED.log_offset").
		Exec(ctx)

	return err
}

func (s *Store) AddUserSegment(ctx context.Context, segment *models.UserSegment) (*models.UserSegment, error) {
	now := time.Now()
	segment.CreatedAt = now
	segment.UpdatedAt = now

	_, err := s.db.NewInsert().
		Model(segment).
		On("CONFLICT (user_id,time) DO UPDATE").
		Set("updated_at = EXCLUDED.updated_at,volume = EXCLUDED.volume,type = EXCLUDED.type,log_offset = EXCLUDED.log_offset").
		Exec(ctx)

	return segment, err
}
