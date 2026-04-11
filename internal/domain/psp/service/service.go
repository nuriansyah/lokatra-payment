package service

import (
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/internal/domain/psp/repository"
)

type Service interface {
}

// ServiceImpl is the service implementation for Credential entities.
type ServiceImpl struct {
	pspRepo repository.Repository
	cfg     *configs.Config
}

// ProvidePspService is the provider for this service.
func ProvidePspService(pspRepo repository.Repository, cfg *configs.Config) *ServiceImpl {
	s := new(ServiceImpl)
	s.pspRepo = pspRepo
	s.cfg = cfg
	return s
}
