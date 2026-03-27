package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/infras"
)

// RepositoryQuery represent repository interfaces
// which consist of repository query operations

type RepositoryCache interface {
	SetBulkCache(ctx context.Context, req map[string]interface{}, expiration time.Duration) error
	GetBulkCache(ctx context.Context, key ...string) ([]interface{}, error)
	SetExCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	GetStringCache(ctx context.Context, key string) (string, error)
	GetStringAndUnmarshalCache(ctx context.Context, key string, dest interface{}) error
	DeleteCache(ctx context.Context, key ...string) error
	NewDistMutex(key string, opts ...Option) infras.DistributedMutex
}
type RepositoryCacheImpl struct {
	cache   *redis.Client
	config  *configs.Config
	redSync *redsync.Redsync
}

// ProvideRepository repository provider
func ProvideRepository(redSync *redsync.Redsync, cache *redis.Client, config *configs.Config) *RepositoryCacheImpl {
	return &RepositoryCacheImpl{
		redSync: redSync,
		cache:   cache,
		config:  config,
	}
}
