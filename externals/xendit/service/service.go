package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/configs"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
)

// Gateway is the Xendit adapter for the normalized PSP bridge.
type Gateway struct {
	enabled      bool
	accountID    uuid.UUID
	accountLabel string
	baseURL      string
	chargePath   string
	secretKey    string
	httpClient   *http.Client
}

func ProvideGateway(cfg *configs.Config) *Gateway {
	gateway := &Gateway{
		baseURL:    "https://api.xendit.co",
		chargePath: "/charges",
		httpClient: &http.Client{Timeout: 15 * time.Second},
	}
	if cfg == nil {
		return gateway
	}

	gateway.enabled = cfg.Externals.Providers.Xendit.Enabled
	gateway.accountID = parseUUID(cfg.Externals.Providers.Xendit.AccountID)
	gateway.accountLabel = fallbackString(cfg.Externals.Providers.Xendit.AccountLabel, "Xendit")
	gateway.baseURL = fallbackString(cfg.Externals.Providers.Xendit.BaseURL, gateway.baseURL)
	gateway.chargePath = fallbackString(cfg.Externals.Providers.Xendit.ChargePath, gateway.chargePath)
	gateway.secretKey = cfg.Externals.Providers.Xendit.SecretKey
	if cfg.Externals.Providers.Xendit.TimeoutSeconds > 0 {
		gateway.httpClient.Timeout = time.Duration(cfg.Externals.Providers.Xendit.TimeoutSeconds) * time.Second
	}
	return gateway
}

func (g *Gateway) PSP() paymentmodel.Psp {
	return paymentmodel.PspXendit
}

func (g *Gateway) Descriptor() paymentmodel.GatewayDescriptor {
	return paymentmodel.GatewayDescriptor{
		PSP:                 g.PSP(),
		AccountID:           g.accountID,
		AccountLabel:        g.accountLabel,
		Enabled:             g.enabled,
		SupportedMethods:    []paymentmodel.PaymentMethodType{paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeQris, paymentmodel.PaymentMethodTypeEwallet, paymentmodel.PaymentMethodTypeBankTransfer},
		SupportedCurrencies: []paymentmodel.PaymentCurrency{paymentmodel.PaymentCurrencyIdr, paymentmodel.PaymentCurrencyUsd, paymentmodel.PaymentCurrencySgd, paymentmodel.PaymentCurrencyMyr, paymentmodel.PaymentCurrencyPhp},
	}
}

func (g *Gateway) Supports(method paymentmodel.PaymentMethodType, currency paymentmodel.PaymentCurrency) bool {
	if !g.enabled {
		return false
	}
	supportedMethods := map[paymentmodel.PaymentMethodType]struct{}{
		paymentmodel.PaymentMethodTypeVirtualAccount: {},
		paymentmodel.PaymentMethodTypeQris:           {},
		paymentmodel.PaymentMethodTypeEwallet:        {},
		paymentmodel.PaymentMethodTypeBankTransfer:   {},
	}
	if _, ok := supportedMethods[method]; !ok {
		return false
	}
	supportedCurrencies := map[paymentmodel.PaymentCurrency]struct{}{
		paymentmodel.PaymentCurrencyIdr: {},
		paymentmodel.PaymentCurrencyUsd: {},
		paymentmodel.PaymentCurrencySgd: {},
		paymentmodel.PaymentCurrencyMyr: {},
		paymentmodel.PaymentCurrencyPhp: {},
	}
	_, ok := supportedCurrencies[currency]
	return ok
}

func (g *Gateway) Charge(ctx context.Context, request paymentmodel.ChargeRequest) (paymentmodel.ChargeResult, error) {
	if !g.enabled {
		return paymentmodel.ChargeResult{}, fmt.Errorf("xendit gateway is disabled")
	}
	if g.secretKey == "" {
		return paymentmodel.ChargeResult{}, fmt.Errorf("xendit secret key is not configured")
	}
	if !g.Supports(request.PaymentMethodType, request.Currency) {
		return paymentmodel.ChargeResult{}, fmt.Errorf("xendit does not support method %s with currency %s", request.PaymentMethodType, request.Currency)
	}

	providerRequest := map[string]any{
		"external_id":         request.PaymentCode,
		"amount":              request.Amount.String(),
		"currency":            string(request.Currency),
		"payment_method_type": strings.ToLower(string(request.PaymentMethodType)),
		"description":         request.Description,
		"metadata":            request.Metadata,
		"customer": map[string]any{
			"given_names":   request.CustomerName,
			"email":         request.CustomerEmail,
			"mobile_number": request.CustomerPhone,
			"country":       request.CustomerCountry,
		},
		"options": map[string]any{
			"capture":              strings.EqualFold(request.CaptureMode, "AUTOMATIC"),
			"three_ds":             request.Requires3DS,
			"statement_descriptor": request.StatementDescriptor,
		},
	}

	providerRequestJSON, err := json.Marshal(providerRequest)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}

	endpoint, err := url.JoinPath(g.baseURL, g.chargePath)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(providerRequestJSON))
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("Authorization", "Basic "+basicAuth(g.secretKey, ""))

	response, err := g.httpClient.Do(httpRequest)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}

	status, requiresAction, transactionID, reference, failureCode, failureMessage := interpretXenditResponse(response.StatusCode, responseBody)
	if transactionID == "" {
		transactionID = buildTransactionID("XD")
	}
	if reference == "" {
		reference = request.PaymentCode
	}

	result := paymentmodel.ChargeResult{
		PSPTransactionID: transactionID,
		PSPReference:     reference,
		RawRequest:       providerRequestJSON,
		RawResponse:      responseBody,
		NextStatus:       status,
		RequiresAction:   requiresAction,
		FailureCode:      failureCode,
		FailureMessage:   failureMessage,
	}
	if response.StatusCode < 400 && (status == paymentmodel.PaymentStatusCaptured || status == paymentmodel.PaymentStatusCompleted || status == paymentmodel.PaymentStatusAuthorised) {
		result.AuthorisedAmount = request.Amount
		if status == paymentmodel.PaymentStatusCaptured || status == paymentmodel.PaymentStatusCompleted {
			result.CapturedAmount = request.Amount
		}
	}
	return result, nil
}

func interpretXenditResponse(statusCode int, body []byte) (paymentmodel.PaymentStatus, bool, string, string, string, string) {
	if statusCode >= 400 {
		return paymentmodel.PaymentStatusFailed, false, "", "", fmt.Sprintf("http_%d", statusCode), string(body)
	}

	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		return paymentmodel.PaymentStatusPending, true, "", "", "", ""
	}

	state := strings.ToLower(stringValue(payload, "status"))
	if state == "" {
		state = strings.ToLower(stringValue(payload, "state"))
	}
	if state == "" {
		state = strings.ToLower(stringValue(payload, "transaction_status"))
	}
	reference := stringValue(payload, "external_id")
	if reference == "" {
		reference = stringValue(payload, "reference_id")
	}
	transactionID := stringValue(payload, "id")
	if transactionID == "" {
		transactionID = stringValue(payload, "transaction_id")
	}

	switch state {
	case "completed", "settlement", "captured", "paid", "succeeded", "success":
		return paymentmodel.PaymentStatusCaptured, false, transactionID, reference, "", ""
	case "authorized", "authorize":
		return paymentmodel.PaymentStatusAuthorised, true, transactionID, reference, "", ""
	case "pending", "waiting", "created":
		return paymentmodel.PaymentStatusPending, true, transactionID, reference, "", ""
	case "cancelled", "canceled":
		return paymentmodel.PaymentStatusCancelled, false, transactionID, reference, "", ""
	case "expired":
		return paymentmodel.PaymentStatusExpired, false, transactionID, reference, "", ""
	case "failed", "rejected", "denied":
		return paymentmodel.PaymentStatusFailed, false, transactionID, reference, stringValue(payload, "status_code"), stringValue(payload, "failure_message")
	default:
		return paymentmodel.PaymentStatusPending, true, transactionID, reference, "", ""
	}
}

func stringValue(payload map[string]any, key string) string {
	if value, ok := payload[key]; ok {
		switch typed := value.(type) {
		case string:
			return typed
		case fmt.Stringer:
			return typed.String()
		case float64:
			return decimal.NewFromFloat(typed).String()
		}
	}
	return ""
}

func basicAuth(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func buildTransactionID(prefix string) string {
	id, _ := uuid.NewV4()
	compact := strings.ReplaceAll(id.String(), "-", "")
	if len(compact) > 10 {
		compact = compact[:10]
	}
	return fmt.Sprintf("%s-%s", prefix, strings.ToUpper(compact))
}
