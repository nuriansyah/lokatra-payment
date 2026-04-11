package service

import (
	"strings"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/configs"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	routingmodel "github.com/nuriansyah/lokatra-payment/internal/domain/routing/model"
)

func NewPaymentServiceConfig(cfg *configs.Config) PaymentServiceConfig {
	serviceConfig := PaymentServiceConfig{
		UseDatabaseFallback: true,
		DefaultUseCase:      PaymentUseCaseGenericTransaction,
		DefaultStrategy:     routingmodel.RoutingStrategyHighestSuccessRate,
	}

	if cfg == nil {
		return serviceConfig
	}

	serviceConfig.UseDatabaseFallback = cfg.Internal.Payment.Routing.UseDatabaseFallback
	serviceConfig.DefaultUseCase = parseUseCase(cfg.Internal.Payment.Routing.DefaultUseCase)
	serviceConfig.DefaultStrategy = parseStrategy(cfg.Internal.Payment.Routing.DefaultStrategy)

	serviceConfig.Providers = []ProviderConfig{
		{
			PSP:                 paymentmodel.PspMidtrans,
			AccountID:           parseUUID(cfg.Externals.Providers.Midtrans.AccountID),
			AccountLabel:        fallbackString(cfg.Externals.Providers.Midtrans.AccountLabel, "Midtrans"),
			Enabled:             cfg.Externals.Providers.Midtrans.Enabled,
			SupportedMethods:    []paymentmodel.PaymentMethodType{paymentmodel.PaymentMethodTypeCard, paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeBankTransfer, paymentmodel.PaymentMethodTypePaylater},
			SupportedCurrencies: []paymentmodel.PaymentCurrency{paymentmodel.PaymentCurrencyIdr, paymentmodel.PaymentCurrencyUsd, paymentmodel.PaymentCurrencySgd},
		},
		{
			PSP:                 paymentmodel.PspXendit,
			AccountID:           parseUUID(cfg.Externals.Providers.Xendit.AccountID),
			AccountLabel:        fallbackString(cfg.Externals.Providers.Xendit.AccountLabel, "Xendit"),
			Enabled:             cfg.Externals.Providers.Xendit.Enabled,
			SupportedMethods:    []paymentmodel.PaymentMethodType{paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeQris, paymentmodel.PaymentMethodTypeEwallet, paymentmodel.PaymentMethodTypeBankTransfer},
			SupportedCurrencies: []paymentmodel.PaymentCurrency{paymentmodel.PaymentCurrencyIdr, paymentmodel.PaymentCurrencyUsd, paymentmodel.PaymentCurrencySgd, paymentmodel.PaymentCurrencyMyr, paymentmodel.PaymentCurrencyPhp},
		},
	}

	serviceConfig.Policies = defaultRoutingPolicies(serviceConfig.DefaultStrategy)
	return serviceConfig
}

func defaultRoutingPolicies(defaultStrategy routingmodel.RoutingStrategy) []RoutingPolicy {
	maxEventAmount := decimal.RequireFromString("10000000")
	maxTripAmount := decimal.RequireFromString("25000000")
	return []RoutingPolicy{
		{
			Name:             "event-payments",
			Enabled:          true,
			Priority:         10,
			UseCases:         []PaymentUseCase{PaymentUseCaseEventPayment},
			Currencies:       []paymentmodel.PaymentCurrency{paymentmodel.PaymentCurrencyIdr},
			PaymentMethods:   []paymentmodel.PaymentMethodType{paymentmodel.PaymentMethodTypeQris, paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeCard},
			MaxAmount:        &maxEventAmount,
			PreferredPSP:     paymentmodel.PspXendit,
			Strategy:         routingmodel.RoutingStrategyWaterfall,
			DatabaseFallback: true,
		},
		{
			Name:             "tours-trips",
			Enabled:          true,
			Priority:         20,
			UseCases:         []PaymentUseCase{PaymentUseCaseTourOrTrip},
			Currencies:       []paymentmodel.PaymentCurrency{paymentmodel.PaymentCurrencyIdr, paymentmodel.PaymentCurrencyUsd, paymentmodel.PaymentCurrencySgd},
			PaymentMethods:   []paymentmodel.PaymentMethodType{paymentmodel.PaymentMethodTypeCard, paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeEwallet, paymentmodel.PaymentMethodTypeBankTransfer},
			MaxAmount:        &maxTripAmount,
			PreferredPSP:     paymentmodel.PspMidtrans,
			Strategy:         routingmodel.RoutingStrategyHighestSuccessRate,
			DatabaseFallback: true,
		},
		{
			Name:             "generic-transaction",
			Enabled:          true,
			Priority:         100,
			UseCases:         []PaymentUseCase{PaymentUseCaseGenericTransaction},
			PaymentMethods:   []paymentmodel.PaymentMethodType{paymentmodel.PaymentMethodTypeCard, paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeQris, paymentmodel.PaymentMethodTypeEwallet, paymentmodel.PaymentMethodTypeBankTransfer, paymentmodel.PaymentMethodTypePaylater},
			PreferredPSP:     paymentmodel.PspXendit,
			Strategy:         defaultStrategy,
			DatabaseFallback: true,
		},
	}
}

func parseUseCase(value string) PaymentUseCase {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case string(PaymentUseCaseEventPayment):
		return PaymentUseCaseEventPayment
	case string(PaymentUseCaseTourOrTrip):
		return PaymentUseCaseTourOrTrip
	default:
		return PaymentUseCaseGenericTransaction
	}
}

func parseStrategy(value string) routingmodel.RoutingStrategy {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case string(routingmodel.RoutingStrategyLowestCost):
		return routingmodel.RoutingStrategyLowestCost
	case string(routingmodel.RoutingStrategyRoundRobin):
		return routingmodel.RoutingStrategyRoundRobin
	case string(routingmodel.RoutingStrategyGeoPreferred):
		return routingmodel.RoutingStrategyGeoPreferred
	case string(routingmodel.RoutingStrategyManual):
		return routingmodel.RoutingStrategyManual
	case string(routingmodel.RoutingStrategyWaterfall):
		return routingmodel.RoutingStrategyWaterfall
	default:
		return routingmodel.RoutingStrategyHighestSuccessRate
	}
}

func parseUUID(value string) uuid.UUID {
	if strings.TrimSpace(value) == "" {
		return uuid.Nil
	}
	id, err := uuid.FromString(value)
	if err != nil {
		return uuid.Nil
	}
	return id
}

func fallbackString(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
