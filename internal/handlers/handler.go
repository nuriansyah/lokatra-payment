package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/service"
	"github.com/nuriansyah/lokatra-payment/transport/http/middleware"
)

type Handler struct {
	config         *configs.Config
	PaymentService *service.ServiceImpl
	PayRouteAdmin  *PayRouteAdminHandler
	Invoice        *InvoiceHandler
}

func ProvideHandler(
	config *configs.Config,
	paymentService *service.ServiceImpl,
	payRouteAdmin *PayRouteAdminHandler,
	invoice *InvoiceHandler,
) *Handler {
	return &Handler{
		config:         config,
		PaymentService: paymentService,
		PayRouteAdmin:  payRouteAdmin,
		Invoice:        invoice,
	}
}

func (h *Handler) Router(r chi.Router) {
	r.Use(middleware.RequestID)

	r.Route("/payment-intents", func(r chi.Router) {
		r.Use(middleware.IdempotencyMiddleware)
		r.Post("/", h.CreatePaymentIntent)
		r.Get("/{paymentIntentID}", h.GetPaymentIntent)
		r.Post("/{paymentIntentID}/{action}", h.PaymentIntentAction)
	})
	r.Post("/webhooks/{provider}", h.HandleWebhook)
	r.With(middleware.IdempotencyMiddleware).Post("/refunds", h.CreateRefund)
	r.Route("/invoices", func(r chi.Router) {
		r.Use(middleware.IdempotencyMiddleware)
		h.Invoice.RegisterRoutes(r)
	})
	r.Route("/admin", func(r chi.Router) {
		r.Use(middleware.RequireAdminToken(h.config))
		r.Post("/refunds/{refundID}/{action}", h.RefundAction)
		r.Post("/webhooks/{webhookID}/{action}", h.AdminWebhookAction)
		r.Post("/manual-payment-evidence/{evidenceID}/{action}", h.ManualPaymentEvidenceAction)
		r.Post("/overpayments/{overpaymentID}/{action}", h.OverpaymentAction)
		r.Post("/cash-sessions", h.OpenCashCollectionSession)
		r.Post("/cash-sessions/{cashSessionID}/{action}", h.CashCollectionSessionAction)
		r.Post("/payment-installments/{installmentID}/{action}", h.PaymentInstallmentAction)
		r.Post("/payment-authorizations/{authorizationID}/{action}", h.PaymentAuthorizationAction)

		// PayRoute Admin Dashboard API
		if h.PayRouteAdmin != nil {
			h.PayRouteAdmin.Router(r)
		}
	})
}
