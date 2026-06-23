package service

import (
	"strings"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/configs"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	durianpay "github.com/nuriansyah/lokatra-payment/external/paymentgateway/durianpay/service"
	midtrans "github.com/nuriansyah/lokatra-payment/external/paymentgateway/midtrans/service"
	xendit "github.com/nuriansyah/lokatra-payment/external/paymentgateway/xendit/service"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
)

type ServiceImpl struct {
	paymentRepo        repository.Repository
	gatewayRegistry    *pg.Registry
	providerAccountIDs map[pg.ProviderCode]uuid.UUID
	webhookConfigured  map[pg.ProviderCode]bool
	routingEngine      *RoutingEngine
	executionLocker    ExecutionLocker
}

func ProvidePaymentService(paymentRepo repository.Repository, cfg *configs.Config, breaker CircuitBreaker, locker ExecutionLocker) *ServiceImpl {
	registry := pg.NewRegistry()
	accountIDs := make(map[pg.ProviderCode]uuid.UUID)
	webhookConfigured := make(map[pg.ProviderCode]bool)
	if cfg != nil && cfg.Externals.Providers.Midtrans.Enabled {
		providerConfig := pg.ProviderConfig{
			Code:            pg.ProviderMidtrans,
			BaseURL:         cfg.Externals.Providers.Midtrans.BaseURL,
			ServerKey:       cfg.Externals.Providers.Midtrans.ServerKey,
			DefaultCurrency: "IDR",
		}
		_ = registry.Register(midtrans.ProvideService(providerConfig))
		accountIDs[pg.ProviderMidtrans] = parseAccountID(cfg.Externals.Providers.Midtrans.AccountID)
		webhookConfigured[pg.ProviderMidtrans] = strings.TrimSpace(providerConfig.ServerKey) != ""
	}
	if cfg != nil && cfg.Externals.Providers.Xendit.Enabled {
		providerConfig := pg.ProviderConfig{
			Code:            pg.ProviderXendit,
			BaseURL:         cfg.Externals.Providers.Xendit.BaseURL,
			APIKey:          cfg.Externals.Providers.Xendit.SecretKey,
			WebhookToken:    cfg.Externals.Providers.Xendit.WebhookToken,
			WebhookSecret:   cfg.Externals.Providers.Xendit.WebhookSecret,
			DefaultCurrency: "IDR",
		}
		_ = registry.Register(xendit.ProvideService(providerConfig))
		accountIDs[pg.ProviderXendit] = parseAccountID(cfg.Externals.Providers.Xendit.AccountID)
		webhookConfigured[pg.ProviderXendit] = strings.TrimSpace(providerConfig.WebhookToken) != "" || strings.TrimSpace(providerConfig.WebhookSecret) != ""
	}
	if cfg != nil && cfg.Externals.Providers.Durianpay.Enabled {
		providerConfig := pg.ProviderConfig{
			Code:            pg.ProviderDurianpay,
			BaseURL:         cfg.Externals.Providers.Durianpay.BaseURL,
			APIKey:          cfg.Externals.Providers.Durianpay.APIKey,
			WebhookSecret:   cfg.Externals.Providers.Durianpay.WebhookSecret,
			DefaultCurrency: "IDR",
		}
		_ = registry.Register(durianpay.ProvideService(providerConfig))
		accountIDs[pg.ProviderDurianpay] = parseAccountID(cfg.Externals.Providers.Durianpay.AccountID)
		webhookConfigured[pg.ProviderDurianpay] = strings.TrimSpace(providerConfig.WebhookSecret) != ""
	}
	routingConfig := NewRoutingConfig(cfg)
	if breaker == nil {
		breaker = NewMemoryCircuitBreaker(routingConfig.FailureThreshold, routingConfig.Cooldown)
	}
	if locker == nil {
		locker = NewMemoryExecutionLocker()
	}
	return &ServiceImpl{
		paymentRepo:        paymentRepo,
		gatewayRegistry:    registry,
		providerAccountIDs: accountIDs,
		webhookConfigured:  webhookConfigured,
		routingEngine:      NewRoutingEngine(registry, accountIDs, breaker, routingConfig),
		executionLocker:    locker,
	}
}

func parseAccountID(value string) uuid.UUID {
	id, err := uuid.FromString(strings.TrimSpace(value))
	if err != nil {
		return uuid.Nil
	}
	return id
}
