package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/infras"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type Repository interface {
	PaymentMethodsRepository
	PaymentsRepository
	PaymentIntentItemsRepository
	PaymentIntentsRepository
	QrisAssignmentsRepository
	VirtualAccountAssignmentsRepository
	DisputesRepository

	BeginTx(ctx context.Context) (Repository, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

// RepositoryImpl is the Postgres-backed implementation of Repository.
type RepositoryImpl struct {
	db   *infras.PostgresConn
	dbTx *sqlx.Tx
	cfg  *configs.Config
}

// ProvideRepository is the provider for this repository.
func ProvideRepository(db *infras.PostgresConn, cfg *configs.Config) *RepositoryImpl {
	s := new(RepositoryImpl)
	s.db = db
	s.cfg = cfg
	return s
}

func (r *RepositoryImpl) BeginTx(ctx context.Context) (Repository, error) {
	tx, err := r.db.Write.Beginx()
	if err != nil {
		log.Error().Err(err).Msg("[BeginTx] begin transaction failed")
		return nil, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
	}

	return &RepositoryImpl{
		db:   r.db,
		dbTx: tx,
	}, nil
}

func (r *RepositoryImpl) Rollback(ctx context.Context) error {
	if r.dbTx == nil {
		log.Error().Msg("[Rollback] not transaction")
		return failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
	}
	return r.dbTx.Rollback()
}

func (r *RepositoryImpl) Commit(ctx context.Context) error {
	if r.dbTx == nil {
		log.Error().Msg("[Commit] not transaction")
		return failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
	}
	return r.dbTx.Commit()
}

func (repo *RepositoryImpl) exec(ctx context.Context, command string, args []interface{}) (sql.Result, error) {
	var (
		stmt *sqlx.Stmt
		err  error
	)

	if repo.dbTx != nil {
		stmt, err = repo.dbTx.PreparexContext(ctx, command)
	} else {
		stmt, err = repo.db.Write.PreparexContext(ctx, command)
	}
	if err != nil {
		log.Error().Err(err).Msg("[exec] failed prepare query")
		return nil, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Error().Err(err).Msg("[exec] failed exec query")

		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return nil, failure.BadRequest(fmt.Errorf(string(shared.ErrorPQConstrainViolated)))
		}

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, failure.BadRequest(fmt.Errorf(string(shared.ErrorPQDuplicateConstrainViolated)))
		}

		return nil, failure.InternalError(fmt.Errorf(string(shared.ErrorInternalSystem)))
	}

	return result, nil
}
