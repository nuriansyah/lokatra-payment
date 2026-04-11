package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/internal/domain/idempotency/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/rs/zerolog/log"
)

// ClaimExecuteIdempotency tries to lock a key for execute flow and returns current row state.
// claimed=true means caller is allowed to proceed with business execution.
func (repo *RepositoryImpl) ClaimExecuteIdempotency(
	ctx context.Context,
	idempotencyKey string,
	merchantID uuid.UUID,
	requestPath string,
	requestBodyHash string,
	actorID uuid.UUID,
	lockUntil time.Time,
) (record model.IdempotencyKeys, claimed bool, err error) {
	tx, err := repo.db.Write.BeginTxx(ctx, nil)
	if err != nil {
		return model.IdempotencyKeys{}, false, failure.InternalError(err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = tx.GetContext(ctx, &record, `
		SELECT id, idempotency_key, merchant_id, request_path, request_body_hash,
		       response_status, response_body, locked_at, locked_until, completed_at,
		       expires_at, meta_created_at, meta_created_by, meta_updated_at,
		       meta_updated_by, meta_deleted_at, meta_deleted_by
		FROM idempotency_keys
		WHERE idempotency_key = $1
		FOR UPDATE
	`, idempotencyKey)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return model.IdempotencyKeys{}, false, failure.InternalError(err)
	}

	now := time.Now().UTC()
	if errors.Is(err, sql.ErrNoRows) {
		record = model.IdempotencyKeys{
			Id:              uuid.Must(uuid.NewV7()),
			IdempotencyKey:  idempotencyKey,
			MerchantId:      merchantID,
			RequestPath:     requestPath,
			RequestBodyHash: requestBodyHash,
			LockedAt:        null.TimeFrom(now),
			LockedUntil:     null.TimeFrom(lockUntil),
			ExpiresAt:       now.Add(24 * time.Hour),
			MetaCreatedAt:   now,
			MetaCreatedBy:   actorID,
			MetaUpdatedAt:   now,
			MetaUpdatedBy:   nuuid.From(actorID),
		}

		_, err = tx.ExecContext(ctx, `
			INSERT INTO idempotency_keys (
				id, idempotency_key, merchant_id, request_path, request_body_hash,
				locked_at, locked_until, expires_at, meta_created_at, meta_created_by,
				meta_updated_at, meta_updated_by
			) VALUES (
				$1, $2, $3, $4, $5,
				$6, $7, $8, $9, $10,
				$11, $12
			)
		`,
			record.Id,
			record.IdempotencyKey,
			record.MerchantId,
			record.RequestPath,
			record.RequestBodyHash,
			record.LockedAt,
			record.LockedUntil,
			record.ExpiresAt,
			record.MetaCreatedAt,
			record.MetaCreatedBy,
			record.MetaUpdatedAt,
			record.MetaUpdatedBy,
		)
		if err != nil {
			return model.IdempotencyKeys{}, false, failure.InternalError(err)
		}

		err = tx.Commit()
		if err != nil {
			return model.IdempotencyKeys{}, false, failure.InternalError(err)
		}
		return record, true, nil
	}

	if record.MerchantId != merchantID {
		return model.IdempotencyKeys{}, false, failure.Conflict("execute", "payment_flow", "idempotency key belongs to another merchant")
	}
	if record.RequestPath != requestPath {
		return model.IdempotencyKeys{}, false, failure.Conflict("execute", "payment_flow", "idempotency key was used on another endpoint")
	}
	if record.RequestBodyHash != requestBodyHash {
		return model.IdempotencyKeys{}, false, failure.Conflict("execute", "payment_flow", "idempotency key reused with different request payload")
	}
	if record.CompletedAt.Valid {
		err = tx.Commit()
		if err != nil {
			return model.IdempotencyKeys{}, false, failure.InternalError(err)
		}
		return record, false, nil
	}
	if record.LockedUntil.Valid && record.LockedUntil.Time.After(now) {
		err = tx.Commit()
		if err != nil {
			return model.IdempotencyKeys{}, false, failure.InternalError(err)
		}
		return record, false, nil
	}

	res, err := tx.ExecContext(ctx, `
		UPDATE idempotency_keys
		SET locked_at = $2,
			locked_until = $3,
			meta_updated_at = $4,
			meta_updated_by = $5
		WHERE id = $1
		  AND completed_at IS NULL
		  AND (locked_until IS NULL OR locked_until <= NOW())
	`, record.Id, now, lockUntil, now, nuuid.From(actorID))
	if err != nil {
		return model.IdempotencyKeys{}, false, failure.InternalError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return model.IdempotencyKeys{}, false, failure.InternalError(err)
	}
	if rowsAffected == 0 {
		err = tx.Commit()
		if err != nil {
			return model.IdempotencyKeys{}, false, failure.InternalError(err)
		}
		return record, false, nil
	}

	record.LockedAt = null.TimeFrom(now)
	record.LockedUntil = null.TimeFrom(lockUntil)
	record.MetaUpdatedAt = now
	record.MetaUpdatedBy = nuuid.From(actorID)

	err = tx.Commit()
	if err != nil {
		return model.IdempotencyKeys{}, false, failure.InternalError(err)
	}

	return record, true, nil
}

// CompleteExecuteIdempotency stores terminal response once.
// updated=false means a concurrent path completed it first.
func (repo *RepositoryImpl) CompleteExecuteIdempotency(
	ctx context.Context,
	id uuid.UUID,
	responseStatus int,
	responseBody json.RawMessage,
	actorID uuid.UUID,
) (updated bool, err error) {
	now := time.Now().UTC()
	res, err := repo.exec(ctx, `
		UPDATE idempotency_keys
		SET response_status = $2,
			response_body = $3,
			completed_at = $4,
			locked_until = NULL,
			meta_updated_at = $4,
			meta_updated_by = $5
		WHERE id = $1
		  AND completed_at IS NULL
	`, []interface{}{id, responseStatus, responseBody, now, nuuid.From(actorID)})
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("[CompleteExecuteIdempotency] failed to update idempotency row")
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, failure.InternalError(err)
	}
	return rowsAffected > 0, nil
}
