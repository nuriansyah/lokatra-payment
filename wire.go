//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/infras"
	cacherepo "github.com/nuriansyah/lokatra-payment/internal/domain/cache/repository"
	"github.com/nuriansyah/lokatra-payment/internal/handlers"
	"github.com/nuriansyah/lokatra-payment/transport/http"
	"github.com/nuriansyah/lokatra-payment/transport/http/middleware"
	"github.com/nuriansyah/lokatra-payment/transport/http/router"
)

// Wiring for configurations.
var configurationsService = wire.NewSet(
	configs.Get,
)

// Wiring for cache repository.
var cacheRepo = wire.NewSet(cacherepo.ProvideRepository, wire.Bind(
	new(cacherepo.RepositoryCache),
	new(*cacherepo.RepositoryCacheImpl),
))

// Wiring for persistences.
var persistencesService = wire.NewSet(
	infras.RedisNewClient,
	infras.ProvideRedisMutex,
	infras.ProvidePostgresConn,
)
var middlewares = wire.NewSet(
	middleware.ProvideAuthentication,
)

// Wiring for HTTP routing.
var routingService = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	handlers.ProvideHandler,
	router.ProvideRouter,
)

// Wiring for everything.
func InitializeServiceService() *http.HTTP {
	wire.Build(
		// configurations
		configurationsService,
		// persistences
		persistencesService,
		// routing
		routingService,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
