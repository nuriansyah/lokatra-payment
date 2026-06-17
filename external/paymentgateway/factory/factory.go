package factory

import (
	"fmt"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	durianpay "github.com/nuriansyah/lokatra-payment/external/paymentgateway/durianpay/service"
	finpay "github.com/nuriansyah/lokatra-payment/external/paymentgateway/finpay/service"
	ipaymu "github.com/nuriansyah/lokatra-payment/external/paymentgateway/ipaymu/service"
	midtrans "github.com/nuriansyah/lokatra-payment/external/paymentgateway/midtrans/service"
	xendit "github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/service"
)

// ProviderFactory implements Factory Method and Abstract Factory patterns.
// The orchestration layer depends on PaymentGateway interface, not concrete provider clients.
type ProviderFactory struct {
	cfg pg.Config
}

func NewProviderFactory(cfg pg.Config) *ProviderFactory { return &ProviderFactory{cfg: cfg} }

func (f *ProviderFactory) Create(code pg.ProviderCode) (pg.PaymentGateway, error) {
	cfg, ok := f.cfg.Provider(code)
	if !ok {
		return nil, fmt.Errorf("%s: %w", code, pg.ErrProviderNotConfigured)
	}
	switch code {
	case pg.ProviderMidtrans:
		return midtrans.ProvideService(cfg), nil
	case pg.ProviderXendit:
		return xendit.ProvideService(cfg), nil
	case pg.ProviderFinpay:
		return finpay.ProvideService(cfg), nil
	case pg.ProviderDurianpay:
		return durianpay.ProvideService(cfg), nil
	case pg.ProviderIpaymu:
		return ipaymu.ProvideService(cfg), nil
	default:
		return nil, fmt.Errorf("unknown payment provider %s: %w", code, pg.ErrProviderNotConfigured)
	}
}

func (f *ProviderFactory) Registry(codes ...pg.ProviderCode) (*pg.Registry, error) {
	reg := pg.NewRegistry()
	if len(codes) == 0 {
		for code := range f.cfg.Providers {
			codes = append(codes, code)
		}
	}
	for _, code := range codes {
		g, err := f.Create(code)
		if err != nil {
			return nil, err
		}
		if err := reg.Register(g); err != nil {
			return nil, err
		}
	}
	return reg, nil
}
