package service

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nuriansyah/lokatra-payment/configs"
)

type ExecutionLocker interface {
	TryLock(ctx context.Context, key string, ttl time.Duration) (unlock func(), acquired bool, err error)
}

type MemoryExecutionLocker struct {
	locks sync.Map
}

func NewMemoryExecutionLocker() *MemoryExecutionLocker { return &MemoryExecutionLocker{} }

func (l *MemoryExecutionLocker) TryLock(_ context.Context, key string, _ time.Duration) (func(), bool, error) {
	if _, loaded := l.locks.LoadOrStore(key, struct{}{}); loaded {
		return func() {}, false, nil
	}
	return func() { l.locks.Delete(key) }, true, nil
}

type RedisExecutionLocker struct {
	client redis.UniversalClient
	prefix string
}

var releaseExecutionLock = redis.NewScript(`
if redis.call('GET', KEYS[1]) == ARGV[1] then
  return redis.call('DEL', KEYS[1])
end
return 0
`)

func ProvideExecutionLocker(cfg *configs.Config, client *redis.Client) ExecutionLocker {
	if cfg == nil || client == nil || strings.TrimSpace(cfg.Cache.Redis.Primary.Host) == "" || strings.TrimSpace(cfg.Cache.Redis.Primary.Port) == "" {
		return NewMemoryExecutionLocker()
	}
	return &RedisExecutionLocker{client: client, prefix: "payment:execution:lock"}
}

func (l *RedisExecutionLocker) TryLock(ctx context.Context, key string, ttl time.Duration) (func(), bool, error) {
	token := mustUUID().String()
	redisKey := l.prefix + ":" + key
	acquired, err := l.client.SetNX(ctx, redisKey, token, ttl).Result()
	if err != nil || !acquired {
		return func() {}, acquired, err
	}
	return func() {
		releaseCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, _ = releaseExecutionLock.Run(releaseCtx, l.client, []string{redisKey}, token).Result()
	}, true, nil
}
