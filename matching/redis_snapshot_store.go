package matching

import (
	"context"
	"encoding/json"
	"time"

	redisClient "github.com/cxptek/vdex/redis"
	"github.com/redis/go-redis/v9"
)

const (
	topicSnapshotPrefix = "matching_snapshot_"
)

type RedisSnapshotStore struct {
	productId   string
	redisClient *redis.Client
}

func NewRedisSnapshotStore(productId string) SnapshotStore {
	// urlStr := config.GetRedisURL()
	// opt, err := redis.ParseURL(urlStr)
	// if err != nil {
	// 	panic(err)
	// }
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     gbeConfig.Redis.Addr,
	// 	Password: gbeConfig.Redis.Password,
	// 	DB:       0,
	// })
	return &RedisSnapshotStore{
		productId:   productId,
		redisClient: redisClient.GetRedisClient(),
	}
}

func (s *RedisSnapshotStore) Store(snapshot *Snapshot) error {
	ctx := context.Background()
	buf, err := json.Marshal(snapshot)
	if err != nil {
		return err
	}

	return s.redisClient.Set(ctx, topicSnapshotPrefix+s.productId, buf, 7*24*time.Hour).Err()
}

func (s *RedisSnapshotStore) GetLatest() (*Snapshot, error) {
	ctx := context.Background()
	ret, err := s.redisClient.Get(ctx, topicSnapshotPrefix+s.productId).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var snapshot Snapshot
	err = json.Unmarshal(ret, &snapshot)
	return &snapshot, err
}
