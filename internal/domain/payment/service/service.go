package service

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/nuriansyah/lokatra-payment/configs"
	idempotencyrepository "github.com/nuriansyah/lokatra-payment/internal/domain/idempotency/repository"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	routingrepository "github.com/nuriansyah/lokatra-payment/internal/domain/routing/repository"
)

type Service interface {
	PaymentService
}

// ServiceImpl is the facade implementation for payment orchestration.
type ServiceImpl struct {
	paymentRepo     repository.Repository
	idempotencyRepo idempotencyrepository.Repository
	routingRepo     routingrepository.Repository
	cfg             *configs.Config
	config          PaymentServiceConfig
	mutex           *redsync.Redsync
	gatewayRegistry GatewayRegistry
	routingEngine   *RoutingEngine
	stateMachine    *PaymentStateMachine
}

// ProvidePaymentService is the provider for this service.
func ProvidePaymentService(paymentRepo repository.Repository, idempotencyRepo idempotencyrepository.Repository, routingRepo routingrepository.Repository, cfg *configs.Config, registry GatewayRegistry, mutex *redsync.Redsync) *ServiceImpl {
	serviceConfig := NewPaymentServiceConfig(cfg)
	engine := NewRoutingEngine(serviceConfig, paymentRepo, routingRepo, registry)
	return &ServiceImpl{
		paymentRepo:     paymentRepo,
		idempotencyRepo: idempotencyRepo,
		routingRepo:     routingRepo,
		cfg:             cfg,
		config:          serviceConfig,
		mutex:           mutex,
		gatewayRegistry: registry,
		routingEngine:   engine,
		stateMachine:    NewPaymentStateMachine(),
	}
}
