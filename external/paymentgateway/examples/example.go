package examples

import (
	"context"
	"time"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	"github.com/nuriansyah/lokatra-payment/external/paymentgateway/factory"
)

func ExampleCreatePayment(ctx context.Context) (pg.CreatePaymentResponse, error) {
	providerFactory := factory.NewProviderFactory(pg.Config{
		HTTP: pg.HTTPClientConfig{Timeout: 15 * time.Second, RetryCount: 2, RetryWaitTime: 200 * time.Millisecond},
		Providers: map[pg.ProviderCode]pg.ProviderConfig{
			pg.ProviderMidtrans: {
				Code:            pg.ProviderMidtrans,
				BaseURL:         "https://api.sandbox.midtrans.com",
				ServerKey:       "MIDTRANS_SERVER_KEY_FROM_VAULT",
				DefaultCurrency: "IDR",
			},
		},
	})
	gateway, err := providerFactory.Create(pg.ProviderMidtrans)
	if err != nil {
		return pg.CreatePaymentResponse{}, err
	}
	return gateway.CreatePayment(ctx, pg.CreatePaymentRequest{
		PaymentIntentID: "pi_123",
		AttemptID:       "pa_123",
		OrderID:         "LKT-20260616-000001",
		Amount:          pg.Money{Amount: "250000.00", Currency: "IDR"},
		Method:          pg.PaymentMethodVirtualAccount,
		ChannelCode:     "bca_va",
		Description:     "Ciwidey ticket booking",
		Customer:        pg.Customer{Name: "Boyan", Email: "boyan@example.com", Phone: "+628123456789"},
		IdempotencyKey:  "pi_123:pa_123:create",
	})
}
