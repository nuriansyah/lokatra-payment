//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nuriansyah/lokatra-payment/configs"
	midtransservice "github.com/nuriansyah/lokatra-payment/externals/midtrans/service"
	xenditservice "github.com/nuriansyah/lokatra-payment/externals/xendit/service"
	"github.com/nuriansyah/lokatra-payment/infras"
	cacherepo "github.com/nuriansyah/lokatra-payment/internal/domain/cache/repository"
	idempotencyrepo "github.com/nuriansyah/lokatra-payment/internal/domain/idempotency/repository"
	paymentrepo "github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	paymentservice "github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	routingrepo "github.com/nuriansyah/lokatra-payment/internal/domain/routing/repository"
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

var repositoriesService = wire.NewSet(
	idempotencyrepo.ProvideRepository,
	wire.Bind(new(idempotencyrepo.Repository), new(*idempotencyrepo.RepositoryImpl)),
	paymentrepo.ProvideRepository,
	wire.Bind(new(paymentrepo.Repository), new(*paymentrepo.RepositoryImpl)),
	routingrepo.ProvideRepository,
	wire.Bind(new(routingrepo.Repository), new(*routingrepo.RepositoryImpl)),
)

var paymentServiceSet = wire.NewSet(
	midtransservice.ProvideGateway,
	xenditservice.ProvideGateway,
	paymentservice.ProvideGatewayRegistry,
	paymentservice.ProvidePaymentService,
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
		// repositories
		repositoriesService,
		// payment service
		paymentServiceSet,
		// routing
		routingService,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
