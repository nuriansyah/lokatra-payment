package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/nuriansyah/lokatra-payment/infras"
	"github.com/rs/zerolog/log"
)

func (r *RepositoryCacheImpl) GetBulkCache(ctx context.Context, key ...string) ([]interface{}, error) {
	cmd := r.cache.MGet(ctx, key...)
	if cmd.Err() != nil {
		return []interface{}{}, cmd.Err()
	}

	return cmd.Val(), nil
}

func (r *RepositoryCacheImpl) SetBulkCache(ctx context.Context, req map[string]interface{}, expiration time.Duration) error {
	var keyValues []interface{}
	for key, val := range req {
		keyValues = append(keyValues, key, val)
	}

	pipe := r.cache.TxPipeline()

	stat := pipe.MSet(ctx, keyValues...)
	if stat.Err() != nil {
		return stat.Err()
	}

	for key, _ := range req {
		cmd := pipe.Expire(ctx, key, expiration)
		if cmd.Err() != nil {
			log.Warn().Err(cmd.Err()).Str("Key", key).Msg("Failed set redis Expiration")
		}
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryCacheImpl) SetExCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	byt, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.cache.SetEX(ctx, key, string(byt), expiration).Err()
}

func (r *RepositoryCacheImpl) GetStringCache(ctx context.Context, key string) (string, error) {
	cmd := r.cache.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return "", err
	}

	return cmd.Val(), nil
}

func (r *RepositoryCacheImpl) GetStringAndUnmarshalCache(ctx context.Context, key string, dest interface{}) error {
	result, err := r.GetStringCache(ctx, key)
	if err != nil {
		return nil
	}

	err = json.Unmarshal([]byte(result), dest)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryCacheImpl) DeleteCache(ctx context.Context, keys ...string) error {
	cmd := r.cache.Del(ctx, keys...)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}

type DistMutexConfig struct {
	options []redsync.Option
}
type Option func(*DistMutexConfig)

func WithRetries(number int) Option {
	return func(dmc *DistMutexConfig) {
		dmc.options = append(dmc.options, redsync.WithTries(number))
	}
}
func WithExpiry(expiry time.Duration) Option {
	return func(dmc *DistMutexConfig) {
		dmc.options = append(dmc.options, redsync.WithExpiry(expiry))
	}
}
func defaultDistMutexConfig() *DistMutexConfig {
	return &DistMutexConfig{
		options: []redsync.Option{},
	}
}
func (r *RepositoryCacheImpl) NewDistMutex(key string, opts ...Option) infras.DistributedMutex {
	defaultConfig := defaultDistMutexConfig()
	for _, opt := range opts {
		opt(defaultConfig)
	}
	key = fmt.Sprintf("%s.%s", key, r.config.Server.Env)
	log.Info().Msgf("create new distributed mutex with key : %s", key)
	return r.redSync.NewMutex(key, defaultConfig.options...)
}
