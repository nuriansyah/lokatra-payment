package paymentgateway

import (
	"context"
	"fmt"
	"sync"
)

type Registry struct {
	mu       sync.RWMutex
	gateways map[ProviderCode]PaymentGateway
}

func NewRegistry() *Registry {
	return &Registry{gateways: make(map[ProviderCode]PaymentGateway)}
}

func (r *Registry) Register(g PaymentGateway) error {
	if g == nil {
		return fmt.Errorf("register nil gateway: %w", ErrInvalidRequest)
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.gateways[g.ProviderCode()] = g
	return nil
}

func (r *Registry) Get(code ProviderCode) (PaymentGateway, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	g, ok := r.gateways[code]
	if !ok {
		return nil, fmt.Errorf("%s: %w", code, ErrProviderNotConfigured)
	}
	return g, nil
}

func (r *Registry) Capabilities(ctx context.Context) (map[ProviderCode]CapabilitiesResponse, error) {
	r.mu.RLock()
	items := make([]PaymentGateway, 0, len(r.gateways))
	for _, g := range r.gateways {
		items = append(items, g)
	}
	r.mu.RUnlock()

	out := make(map[ProviderCode]CapabilitiesResponse, len(items))
	for _, g := range items {
		capResp, err := g.Capabilities(ctx, CapabilitiesRequest{})
		if err != nil {
			return nil, err
		}
		out[g.ProviderCode()] = capResp
	}
	return out, nil
}
