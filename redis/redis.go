package redisClient

import (
	"context"
	"fmt"
	"time"

	"github.com/bsm/redislock"
	"github.com/cxptek/vdex/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var redisLock *redislock.Client
var (
	Vdex_Total_Gasless_Tx_Key = "vdex_total_gasless_tx_key"
	Vdex_Gasless_Executor     = "vdex_gasless_executor"
)

func NewRedis() {
	urlStr := config.GetRedisURL()
	opt, err := redis.ParseURL(urlStr)
	if err != nil {
		panic(err)
	}
	redisClient = redis.NewClient(opt)
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("connect redis err: %v", err.Error()))
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetRedisLockWithDeadline(ctx context.Context, key string) (*redislock.Lock, error) {
	if redisLock == nil {
		redisLock = redislock.New(redisClient)
	}

	backoff := redislock.LinearBackoff(200 * time.Millisecond)

	// Obtain lock with retry + custom deadline
	lock, err := redisLock.Obtain(ctx, key, time.Second, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err != nil {
		return nil, err
	}

	return lock, nil
}
