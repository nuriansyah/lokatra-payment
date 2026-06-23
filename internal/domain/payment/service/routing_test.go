package service

import (
	"context"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/configs"
	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/stretchr/testify/require"
)

func TestNewRoutingConfigParsesMethodAndChannelRules(t *testing.T) {
	configSource := &configs.Config{}
	configSource.Internal.Payment.Routing.DefaultProviders = "durianpay,xendit"
	configSource.Internal.Payment.Routing.MaxAttempts = 4
	configSource.Internal.Payment.Routing.RulesJSON = `[{"method":"qris","channel":"qris","providers":["xendit","durianpay"],"maxAttempts":3}]`

	config := NewRoutingConfig(configSource)
	require.Equal(t, []pg.ProviderCode{pg.ProviderDurianpay, pg.ProviderXendit}, config.DefaultProviders)
	require.Equal(t, 4, config.MaxAttempts)
	require.Len(t, config.Rules, 1)
	require.Equal(t, pg.PaymentMethodQRIS, config.Rules[0].Method)
	require.Equal(t, []pg.ProviderCode{pg.ProviderXendit, pg.ProviderDurianpay}, config.Rules[0].Providers)
}

func TestMemoryCircuitBreakerAllowsSingleHalfOpenProbe(t *testing.T) {
	breaker := NewMemoryCircuitBreaker(3, time.Minute)
	now := time.Now().UTC()
	for range 3 {
		breaker.RecordFailure("xendit:qris:qris", now)
	}
	require.False(t, breaker.Allow("xendit:qris:qris", now.Add(30*time.Second)))
	require.True(t, breaker.Allow("xendit:qris:qris", now.Add(time.Minute)))
	require.False(t, breaker.Allow("xendit:qris:qris", now.Add(time.Minute)), "only one half-open probe is allowed")
	breaker.RecordSuccess("xendit:qris:qris")
	require.True(t, breaker.Allow("xendit:qris:qris", now.Add(time.Minute)))
}

func TestRoutingRuleSpecificity(t *testing.T) {
	engine := &RoutingEngine{config: RoutingConfig{
		Rules: []RoutingRule{
			{Method: pg.PaymentMethodVirtualAccount, Channel: "*", Providers: []pg.ProviderCode{pg.ProviderXendit, pg.ProviderDurianpay}, MaxAttempts: 3},
			{Method: pg.PaymentMethodVirtualAccount, Channel: "bca_va", Providers: []pg.ProviderCode{pg.ProviderMidtrans, pg.ProviderXendit}, MaxAttempts: 2},
		},
		DefaultProviders: []pg.ProviderCode{pg.ProviderDurianpay}, MaxAttempts: 1,
	}}

	providers, attempts, reason := engine.matchRule(pg.PaymentMethodVirtualAccount, "bca_va")
	require.Equal(t, []pg.ProviderCode{pg.ProviderMidtrans, pg.ProviderXendit}, providers)
	require.Equal(t, 2, attempts)
	require.Equal(t, "method_channel_rule", reason)

	providers, attempts, reason = engine.matchRule(pg.PaymentMethodVirtualAccount, "bni_va")
	require.Equal(t, []pg.ProviderCode{pg.ProviderXendit, pg.ProviderDurianpay}, providers)
	require.Equal(t, 3, attempts)
	require.Equal(t, "method_rule", reason)
}

type routingTestGateway struct {
	pg.PaymentGateway
	code         pg.ProviderCode
	capabilities []pg.Capability
	create       func() (pg.CreatePaymentResponse, error)
}

func (g *routingTestGateway) ProviderCode() pg.ProviderCode { return g.code }
func (g *routingTestGateway) Capabilities(context.Context, pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{ProviderCode: g.code, Items: g.capabilities}, nil
}
func (g *routingTestGateway) CreatePayment(context.Context, pg.CreatePaymentRequest) (pg.CreatePaymentResponse, error) {
	return g.create()
}

func TestRoutingEngineRetriesPrimaryThenFallsBackAndOpensCircuit(t *testing.T) {
	xenditCalls, durianCalls := 0, 0
	registry := pg.NewRegistry()
	require.NoError(t, registry.Register(&routingTestGateway{
		code:         pg.ProviderXendit,
		capabilities: []pg.Capability{{Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: "IDR"}},
		create: func() (pg.CreatePaymentResponse, error) {
			xenditCalls++
			return pg.CreatePaymentResponse{}, pg.NewGatewayError(pg.ProviderXendit, pg.ErrorCodeProviderTimeout, 504, "timeout", true, nil)
		},
	}))
	require.NoError(t, registry.Register(&routingTestGateway{
		code:         pg.ProviderDurianpay,
		capabilities: []pg.Capability{{Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: "IDR"}},
		create: func() (pg.CreatePaymentResponse, error) {
			durianCalls++
			return pg.CreatePaymentResponse{ProviderCode: pg.ProviderDurianpay, Status: pg.PaymentStatusPending}, nil
		},
	}))

	config := RoutingConfig{
		Rules:            []RoutingRule{{Method: pg.PaymentMethodQRIS, Channel: "qris", Providers: []pg.ProviderCode{pg.ProviderXendit, pg.ProviderDurianpay}, MaxAttempts: 3}},
		DefaultProviders: []pg.ProviderCode{pg.ProviderXendit, pg.ProviderDurianpay},
		MaxAttempts:      3, FailureThreshold: 3, Cooldown: time.Minute,
	}
	engine := NewRoutingEngine(registry, map[pg.ProviderCode]uuid.UUID{
		pg.ProviderXendit: mustUUID(), pg.ProviderDurianpay: mustUUID(),
	}, NewMemoryCircuitBreaker(3, time.Minute), config)
	request := RoutingRequest{Method: pg.PaymentMethodQRIS, Channel: "qris", Currency: "IDR", GatewayCall: pg.CreatePaymentRequest{Method: pg.PaymentMethodQRIS}}

	result, err := engine.Execute(context.Background(), request)
	require.NoError(t, err)
	require.Equal(t, pg.ProviderDurianpay, result.Selected.ProviderCode)
	require.Equal(t, 3, xenditCalls)
	require.Equal(t, 1, durianCalls)
	require.Len(t, result.Attempts, 4)

	result, err = engine.Execute(context.Background(), request)
	require.NoError(t, err)
	require.Equal(t, pg.ProviderDurianpay, result.Selected.ProviderCode)
	require.Equal(t, 3, xenditCalls, "open circuit must skip the primary provider")
	require.Equal(t, 2, durianCalls)
	require.True(t, result.Candidates[0].Skipped)
	require.Equal(t, "circuit_open", result.Candidates[0].SkipReason)
}

func TestRoutingEngineDoesNotFallbackOnInvalidRequest(t *testing.T) {
	fallbackCalls := 0
	registry := pg.NewRegistry()
	capability := []pg.Capability{{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "va", Currency: "IDR"}}
	require.NoError(t, registry.Register(&routingTestGateway{
		code: pg.ProviderXendit, capabilities: capability,
		create: func() (pg.CreatePaymentResponse, error) {
			return pg.CreatePaymentResponse{}, pg.NewGatewayError(pg.ProviderXendit, pg.ErrorCodeInvalidRequest, 400, "invalid", false, nil)
		},
	}))
	require.NoError(t, registry.Register(&routingTestGateway{
		code: pg.ProviderDurianpay, capabilities: capability,
		create: func() (pg.CreatePaymentResponse, error) {
			fallbackCalls++
			return pg.CreatePaymentResponse{ProviderCode: pg.ProviderDurianpay}, nil
		},
	}))
	config := RoutingConfig{DefaultProviders: []pg.ProviderCode{pg.ProviderXendit, pg.ProviderDurianpay}, MaxAttempts: 3, FailureThreshold: 3, Cooldown: time.Minute}
	engine := NewRoutingEngine(registry, map[pg.ProviderCode]uuid.UUID{
		pg.ProviderXendit: mustUUID(), pg.ProviderDurianpay: mustUUID(),
	}, NewMemoryCircuitBreaker(3, time.Minute), config)

	_, err := engine.Execute(context.Background(), RoutingRequest{Method: pg.PaymentMethodVirtualAccount, Channel: "bca_va", Currency: "IDR"})
	require.Error(t, err)
	require.Zero(t, fallbackCalls)
}
