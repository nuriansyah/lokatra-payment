package service

import (
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/internal/domain/routing/repository"
)

type Service interface {
}

// ServiceImpl is the service implementation for Credential entities.
type ServiceImpl struct {
	routingRepo repository.Repository
	cfg         *configs.Config
}

// ProvideRoutingService is the provider for this service.
func ProvideRoutingService(routingRepo repository.Repository, cfg *configs.Config) *ServiceImpl {
	s := new(ServiceImpl)
	s.routingRepo = routingRepo
	s.cfg = cfg
	return s
}
