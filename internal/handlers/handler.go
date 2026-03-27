package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/nuriansyah/lokatra-payment/configs"
	"github.com/nuriansyah/lokatra-payment/transport/http/middleware"
)

type Handler struct {
	config *configs.Config
}

func ProvideHandler(
	config *configs.Config,
) *Handler {
	return &Handler{
		config: config,
	}
}

func (h *Handler) Router(r chi.Router) {
	r.Use(middleware.RequestID)

}
