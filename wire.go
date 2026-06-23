//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/infras"
	paymentrepo "github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	paymentservice "github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	"github.com/nuriansyah/lokatra-payment/internal/handlers"
	"github.com/nuriansyah/lokatra-payment/transport/http"
	"github.com/nuriansyah/lokatra-payment/transport/http/router"
)

// Wiring for configurations.
var configurationsService = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistencesService = wire.NewSet(
	infras.ProvidePostgresConn,
	infras.RedisNewClient,
)

var repositoriesService = wire.NewSet(
	paymentrepo.ProvideRepository,
	wire.Bind(new(paymentrepo.Repository), new(*paymentrepo.RepositoryImpl)),
)

var paymentServiceSet = wire.NewSet(
	paymentservice.ProvideCircuitBreaker,
	paymentservice.ProvideExecutionLocker,
	paymentservice.ProvidePaymentService,
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
