package pushing

import (
	"fmt"

	"github.com/cxptek/vdex/config"
	"github.com/go-redis/redis"
)

const (
	topicSnapshotPrefix = "pushing_snapshot_"
)

type PushingSnapshotStore struct {
	redisClient *redis.Client
}

type Snapshot struct {
	OrderOffset int64 `json:"orderOffset"`
	TradeOffset int64 `json:"tradeOffset"`
}

var redisPushingSnapshotStore *PushingSnapshotStore = nil

func NewRedisPushingSnapshotStore() *PushingSnapshotStore {
	urlStr := config.GetRedisURL()
	opt, err := redis.ParseURL(urlStr)
	if err != nil {
		panic(err)
	}
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     gbeConfig.Redis.Addr,
	// 	Password: gbeConfig.Redis.Password,
	// 	DB:       0,
	// })
	redisClient := redis.NewClient(opt)
	if err := redisClient.Ping().Err(); err != nil {
		panic(fmt.Sprintf("connect redis err: %v", err.Error()))
	}
	redisPushingSnapshotStore = &PushingSnapshotStore{
		redisClient: redisClient,
	}

	return redisPushingSnapshotStore
}

func (s *PushingSnapshotStore) StorePushingOffset(key string, val int64) error {
	return s.redisClient.Set(topicSnapshotPrefix+key, val, 0).Err()
}

func (s *PushingSnapshotStore) GetPushingOffset(key string) (int64, error) {
	ret, err := s.redisClient.Get(topicSnapshotPrefix + key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	return ret, nil
}
