package infras

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/nuriansyah/lokatra-payment/configs"
)

type DistributedMutex interface {
	Lock() error
	Unlock() (bool, error)
}

// RedisNewClient create new instance of redis
func RedisNewClient(config *configs.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Cache.Redis.Primary.Host, config.Cache.Redis.Primary.Port),
		Password: config.Cache.Redis.Primary.Password,
		DB:       config.Cache.Redis.Primary.DB,
	})

	return client
}

// ProvideRedisMutex create new instance of redis distributed lock
func ProvideRedisMutex(redis *redis.Client) *redsync.Redsync {
	return redsync.New(goredis.NewPool(redis))
}
