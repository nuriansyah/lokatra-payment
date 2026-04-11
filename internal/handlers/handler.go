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
}

func ProvideHandler(
	config *configs.Config,
	paymentService *service.ServiceImpl,
) *Handler {
	return &Handler{
		config:         config,
		PaymentService: paymentService,
	}
}

func (h *Handler) Router(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Route("/payments", func(pr chi.Router) {
		r.Group(func(r chi.Router) {
			pr.Post("/flows/quote", h.quotePaymentFlow)
			pr.Post("/flows/execute", h.executePaymentFlow)
		})
		// Webhook routes (typically not authenticated for PSP callbacks)
		pr.Route("/webhooks", func(wr chi.Router) {
			wr.Post("/midtrans", h.HandleMidtransWebhook)
			wr.Post("/xendit", h.HandleXenditWebhook)
			wr.Get("/health", h.HandleWebhookTest)
		})
	})
}
