package retry

import (
	"time"
)

type Retry struct {
	Config Config
}

const (
	defaultMaxAttempt = 3
	defaultMaxDelay   = 1 * time.Second
)

func New(opts ...Option) Retry {
	cfg := Config{
		MaxAttempt: defaultMaxAttempt,
		MaxDelay:   defaultMaxDelay,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return Retry{
		Config: cfg,
	}
}

func (r *Retry) Do(retryable RetryableFunc) error {
	var err error

	if r.Config.MaxAttempt == 0 {
		return retryable()
	}

	counter := 0
	for counter < r.Config.MaxAttempt {
		err = retryable()
		if err == nil {
			return nil
		}
		counter++
		time.Sleep(r.Config.MaxDelay)
	}

	return err
}

func (r *Retry) WithExponential(retryable RetryableFunc) error {
	var err error

	if r.Config.MaxAttempt == 0 {
		return retryable()
	}

	counter := 0
	for counter < r.Config.MaxAttempt {
		err = retryable()
		if err == nil {
			return nil
		}
		counter++
		sleep := time.Duration(counter/r.Config.MaxAttempt) * r.Config.MaxDelay
		time.Sleep(sleep)
	}

	return err
}

type Config struct {
	MaxAttempt int
	MaxDelay   time.Duration
}

type Option func(*Config)

func SetMaxAttempt(maxAttempt int) Option {
	return func(c *Config) {
		c.MaxAttempt = maxAttempt
	}
}

func SetMaxDelay(maxDelay time.Duration) Option {
	return func(c *Config) {
		c.MaxDelay = maxDelay
	}
}

type RetryableFunc func() error
