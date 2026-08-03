package service

import (
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
)

// ServiceImpl implements pg.PaymentGateway for Xendit.
// Each domain operation lives in its own file:
//   - payment.go:  CreatePayment, GetPaymentStatus, CancelPayment
//   - refund.go:   RefundPayment
//   - payout.go:   CreatePayout, GetPayoutStatus
//   - webhook.go:  VerifyWebhook, NormalizeWebhook
//   - config.go:   Capabilities
type ServiceImpl struct {
	client *resty.Client
	cfg    pg.ProviderConfig
	logger zerolog.Logger
}

// ProvideService constructs a Xendit gateway adapter.
func ProvideService(cfg pg.ProviderConfig) *ServiceImpl {
	return &ServiceImpl{
		client: pg.NewRestyClient(cfg),
		cfg:    cfg,
		logger: zerolog.Nop(),
	}
}

// ProvideServiceWithLogger constructs a Xendit gateway adapter with structured logging.
func ProvideServiceWithLogger(cfg pg.ProviderConfig, logger zerolog.Logger) *ServiceImpl {
	return &ServiceImpl{
		client: pg.NewRestyClient(cfg),
		cfg:    cfg,
		logger: logger.With().Str("provider", "xendit").Logger(),
	}
}

// ProviderCode returns the canonical provider identifier.
func (s *ServiceImpl) ProviderCode() pg.ProviderCode { return pg.ProviderXendit }

// auth sets the Xendit Basic Auth header on the request.
func (s *ServiceImpl) auth(req *resty.Request) *resty.Request {
	return req.SetHeader("Authorization", pg.BasicAuthValue(s.cfg.APIKey, ""))
}

// log returns the service-scoped logger.
func (s *ServiceImpl) log() *zerolog.Logger {
	return &s.logger
}
