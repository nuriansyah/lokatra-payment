package repository

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/infras"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type Repository interface {
	FinanceCurrenciesRepository
}

// RepositoryImpl is the Postgres-backed implementation of Repository.
type RepositoryImpl struct {
	db *infras.PostgresConn
}

// ProvideRepository is the provider for this repository.
func ProvideRepository(db *infras.PostgresConn) *RepositoryImpl {
	s := new(RepositoryImpl)
	s.db = db
	return s
}

func (repo *RepositoryImpl) exec(ctx context.Context, command string, args []interface{}) (sql.Result, error) {
	result, err := repo.db.Write.ExecContext(ctx, command, args...)
	if err != nil {
		log.Error().Err(err).Msg("[exec] failed exec query")
		return nil, failure.InternalError(err)
	}

	return result, nil
}
