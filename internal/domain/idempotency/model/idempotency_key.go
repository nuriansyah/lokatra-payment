package model

// IdempotencyKeyStatus represents the status of an idempotency key.
type IdempotencyKeyStatus string

const (
	// PENDING operation is in progress
	IdempotencyKeyStatusPending IdempotencyKeyStatus = "PENDING"

	// SUCCESS operation completed successfully
	IdempotencyKeyStatusSuccess IdempotencyKeyStatus = "SUCCESS"

	// FAILED operation completed with error
	IdempotencyKeyStatusFailed IdempotencyKeyStatus = "FAILED"

	// EXPIRED key has expired and no longer valid
	IdempotencyKeyStatusExpired IdempotencyKeyStatus = "EXPIRED"
)

// String returns the string representation of the status.
func (s IdempotencyKeyStatus) String() string {
	return string(s)
}
