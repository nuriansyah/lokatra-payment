package service

import (
	"context"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
)

// Capabilities returns the static list of payment methods this Xendit adapter supports.
// This is read-only configuration — no network call is made.
func (s *ServiceImpl) Capabilities(ctx context.Context, _ pg.CapabilitiesRequest) (pg.CapabilitiesResponse, error) {
	return pg.CapabilitiesResponse{
		ProviderCode: s.ProviderCode(),
		Items: []pg.Capability{
			{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bca_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
			{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "mandiri_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
			{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bni_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
			{Method: pg.PaymentMethodVirtualAccount, ChannelCode: "bri_va", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
			{Method: pg.PaymentMethodQRIS, ChannelCode: "qris", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
			{Method: pg.PaymentMethodEWallet, ChannelCode: "ovo", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
			{Method: pg.PaymentMethodEWallet, ChannelCode: "dana", Currency: s.cfg.Currency(), SupportsRefund: true, SupportsPartialRefund: true, SupportsExpiry: true},
		},
	}, nil
}
