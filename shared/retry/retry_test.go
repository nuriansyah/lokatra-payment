package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/nuriansyah/lokatra-payment/shared/retry"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("RetryDefaultInstantiate", func(t *testing.T) {
		actualRetry := retry.New()
		expectedRetry := retry.Retry{
			Config: retry.Config{
				MaxAttempt: 3,
				MaxDelay:   1 * time.Second,
			},
		}

		assert.Equal(t, expectedRetry, actualRetry)
	})

	t.Run("RetryConfigurableInstantiate", func(t *testing.T) {
		actualRetry := retry.New(retry.SetMaxAttempt(5), retry.SetMaxDelay(1*time.Second))
		expectedRetry := retry.Retry{
			Config: retry.Config{
				MaxAttempt: 5,
				MaxDelay:   1 * time.Second,
			},
		}
		assert.Equal(t, expectedRetry, actualRetry)
	})
}

func TestRetryable(t *testing.T) {
	t.Run("RetryNoError", func(t *testing.T) {
		r := retry.New(retry.SetMaxAttempt(0))
		actualErr := r.Do(func() error {
			return nil
		})
		var expectedErr error = nil

		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("RetryError", func(t *testing.T) {
		r := retry.New()
		actualErr := r.Do(func() error {
			return errors.New("error")
		})
		var expectedErr error = errors.New("error")
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("RetryWithExponential", func(t *testing.T) {
		r := retry.New(retry.SetMaxAttempt(0))
		actualErr := r.WithExponential(func() error {
			return nil
		})
		var expectedErr error = nil

		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("RetryErrorWithExponential", func(t *testing.T) {
		r := retry.New()
		actualErr := r.WithExponential(func() error {
			return errors.New("error")
		})
		var expectedErr error = errors.New("error")
		assert.Equal(t, expectedErr, actualErr)
	})
}
