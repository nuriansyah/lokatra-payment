package service

import (
	"strings"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"

	midtransservice "github.com/nuriansyah/lokatra-payment/externals/midtrans/service"
	xenditservice "github.com/nuriansyah/lokatra-payment/externals/xendit/service"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
)

// GatewayRegistry is the bridge registry used by the facade.
type GatewayRegistry struct {
	ByPSP     map[paymentmodel.Psp]paymentmodel.PaymentGateway
	ByAccount map[uuid.UUID]paymentmodel.PaymentGateway
}

func NewGatewayRegistry(gateways ...paymentmodel.PaymentGateway) GatewayRegistry {
	registry := GatewayRegistry{
		ByPSP:     map[paymentmodel.Psp]paymentmodel.PaymentGateway{},
		ByAccount: map[uuid.UUID]paymentmodel.PaymentGateway{},
	}
	for _, gateway := range gateways {
		if gateway == nil {
			continue
		}
		descriptor := gateway.Descriptor()
		registry.ByPSP[gateway.PSP()] = gateway
		if descriptor.AccountID != uuid.Nil {
			registry.ByAccount[descriptor.AccountID] = gateway
		}
	}
	return registry
}

func (r GatewayRegistry) LookupByPSP(psp paymentmodel.Psp) (paymentmodel.PaymentGateway, bool) {
	gateway, ok := r.ByPSP[psp]
	return gateway, ok
}

func (r GatewayRegistry) LookupByAccountID(accountID uuid.UUID) (paymentmodel.PaymentGateway, bool) {
	if accountID == uuid.Nil {
		return nil, false
	}
	gateway, ok := r.ByAccount[accountID]
	return gateway, ok
}

func ProvideGatewayRegistry(midtransGateway *midtransservice.Gateway, xenditGateway *xenditservice.Gateway) GatewayRegistry {
	gateways := make([]paymentmodel.PaymentGateway, 0, 2)
	if midtransGateway != nil && strings.EqualFold(string(midtransGateway.PSP()), string(paymentmodel.PspMidtrans)) {
		gateways = append(gateways, midtransGateway)
	}
	if xenditGateway != nil && strings.EqualFold(string(xenditGateway.PSP()), string(paymentmodel.PspXendit)) {
		gateways = append(gateways, xenditGateway)
	}
	registry := NewGatewayRegistry(gateways...)
	log.Info().Int("gateway_count", len(registry.ByPSP)).Msg("payment gateway registry initialized")
	return registry
}
