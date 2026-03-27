package middleware

import (
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/infras"
)

type Authentication struct {
	db  *infras.PostgresConn
	cfg *configs.Config
}
type ContextKey string

const (
	HeaderAuthorization       = "Authorization"
	HeaderAuthorizationPrefix = "Bearer"

	UserInfo ContextKey = "UserInfo"
)

func ProvideAuthentication(cfg *configs.Config, db *infras.PostgresConn) *Authentication {
	return &Authentication{
		cfg: cfg,
		db:  db,
	}
}
