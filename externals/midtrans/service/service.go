package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/configs"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
)

// Gateway is the Midtrans adapter for the normalized PSP bridge.
type Gateway struct {
	enabled      bool
	accountID    uuid.UUID
	accountLabel string
	serverKey    string
	environment  midtrans.EnvironmentType
	coreClient   midtransCoreClient
}

type midtransCoreClient interface {
	ChargeTransaction(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error)
	CheckTransaction(param string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
}

type coreAPIChargeClient struct {
	client *coreapi.Client
}

func (c *coreAPIChargeClient) ChargeTransaction(req *coreapi.ChargeReq) (*coreapi.ChargeResponse, *midtrans.Error) {
	return c.client.ChargeTransaction(req)
}

func (c *coreAPIChargeClient) CheckTransaction(param string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {
	return c.client.CheckTransaction(param)
}

func ProvideGateway(cfg *configs.Config) *Gateway {
	gateway := &Gateway{
		environment: midtrans.Production,
	}
	if cfg == nil {
		gateway.coreClient = buildCoreAPIClient(gateway.serverKey, gateway.environment, 15*time.Second)
		return gateway
	}

	baseURL := fallbackString(cfg.Externals.Providers.Midtrans.BaseURL, "https://api.midtrans.com")
	timeout := 15 * time.Second
	if cfg.Externals.Providers.Midtrans.TimeoutSeconds > 0 {
		timeout = time.Duration(cfg.Externals.Providers.Midtrans.TimeoutSeconds) * time.Second
	}

	gateway.enabled = cfg.Externals.Providers.Midtrans.Enabled
	gateway.accountID = parseUUID(cfg.Externals.Providers.Midtrans.AccountID)
	gateway.accountLabel = fallbackString(cfg.Externals.Providers.Midtrans.AccountLabel, "Midtrans")
	gateway.serverKey = cfg.Externals.Providers.Midtrans.ServerKey
	gateway.environment = resolveMidtransEnvironment(baseURL)
	gateway.coreClient = buildCoreAPIClient(gateway.serverKey, gateway.environment, timeout)
	return gateway
}

func (g *Gateway) PSP() paymentmodel.Psp {
	return paymentmodel.PspMidtrans
}

func (g *Gateway) Descriptor() paymentmodel.GatewayDescriptor {
	return paymentmodel.GatewayDescriptor{
		PSP:                 g.PSP(),
		AccountID:           g.accountID,
		AccountLabel:        g.accountLabel,
		Enabled:             g.enabled,
		SupportedMethods:    []paymentmodel.PaymentMethodType{paymentmodel.PaymentMethodTypeCard, paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeBankTransfer, paymentmodel.PaymentMethodTypePaylater},
		SupportedCurrencies: []paymentmodel.PaymentCurrency{paymentmodel.PaymentCurrencyIdr, paymentmodel.PaymentCurrencyUsd, paymentmodel.PaymentCurrencySgd},
	}
}

func (g *Gateway) Supports(method paymentmodel.PaymentMethodType, currency paymentmodel.PaymentCurrency) bool {
	if !g.enabled {
		return false
	}
	supportedMethods := map[paymentmodel.PaymentMethodType]struct{}{
		paymentmodel.PaymentMethodTypeCard:           {},
		paymentmodel.PaymentMethodTypeVirtualAccount: {},
		paymentmodel.PaymentMethodTypeBankTransfer:   {},
		paymentmodel.PaymentMethodTypePaylater:       {},
	}
	if _, ok := supportedMethods[method]; !ok {
		return false
	}
	supportedCurrencies := map[paymentmodel.PaymentCurrency]struct{}{
		paymentmodel.PaymentCurrencyIdr: {},
		paymentmodel.PaymentCurrencyUsd: {},
		paymentmodel.PaymentCurrencySgd: {},
	}
	_, ok := supportedCurrencies[currency]
	return ok
}

func (g *Gateway) Charge(ctx context.Context, request paymentmodel.ChargeRequest) (paymentmodel.ChargeResult, error) {
	if !g.enabled {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans gateway is disabled")
	}
	if g.serverKey == "" {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans server key is not configured")
	}
	if !g.Supports(request.PaymentMethodType, request.Currency) {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans does not support method %s with currency %s", request.PaymentMethodType, request.Currency)
	}
	if g.coreClient == nil {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans charge client is not initialized")
	}

	providerRequest, err := buildCoreAPIChargeRequest(request)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}
	providerRequestJSON, err := json.Marshal(providerRequest)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}

	providerResponse, sdkErr := g.coreClient.ChargeTransaction(providerRequest)
	if sdkErr != nil {
		failedResult := chargeResultFromSDKError(request, providerRequestJSON, sdkErr)
		if sdkErr.StatusCode >= http.StatusBadRequest {
			return failedResult, nil
		}
		return failedResult, fmt.Errorf("midtrans charge request failed: %w", sdkErr)
	}
	if providerResponse == nil {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans charge request returned empty response")
	}

	providerResponseJSON, err := json.Marshal(providerResponse)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}

	status, requiresAction, transactionID, reference, failureCode, failureMessage := interpretMidtransChargeResponse(providerResponse)
	nextStatus := status
	if nextStatus == "" {
		nextStatus = paymentmodel.PaymentStatusPending
	}
	if transactionID == "" {
		transactionID = buildTransactionID("MT")
	}
	if reference == "" {
		reference = request.PaymentCode
	}

	result := paymentmodel.ChargeResult{
		PSPTransactionID: transactionID,
		PSPReference:     reference,
		RawRequest:       providerRequestJSON,
		RawResponse:      providerResponseJSON,
		NextStatus:       nextStatus,
		RequiresAction:   requiresAction,
		FailureCode:      failureCode,
		FailureMessage:   failureMessage,
	}
	if nextStatus == paymentmodel.PaymentStatusCaptured || nextStatus == paymentmodel.PaymentStatusCompleted || nextStatus == paymentmodel.PaymentStatusAuthorised {
		result.AuthorisedAmount = request.Amount
		if nextStatus == paymentmodel.PaymentStatusCaptured || nextStatus == paymentmodel.PaymentStatusCompleted {
			result.CapturedAmount = request.Amount
		}
	}
	return result, nil
}

func buildCoreAPIChargeRequest(request paymentmodel.ChargeRequest) (*coreapi.ChargeReq, error) {
	amountInt, err := decimalAmountToInt64(request.Amount)
	if err != nil {
		return nil, err
	}

	paymentType := midtransPaymentType(request.PaymentMethodType)
	providerRequest := &coreapi.ChargeReq{
		PaymentType: paymentType,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.PaymentCode,
			GrossAmt: amountInt,
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: request.CustomerName,
			Email: request.CustomerEmail,
			Phone: request.CustomerPhone,
		},
		Items: &[]midtrans.ItemDetails{{
			ID:    request.IntentCode,
			Price: amountInt,
			Qty:   1,
			Name:  fallbackString(request.Description, request.PaymentCode),
		}},
		Metadata: request.Metadata,
	}

	switch request.PaymentMethodType {
	case paymentmodel.PaymentMethodTypeCard:
		providerRequest.CreditCard = &coreapi.CreditCardDetails{Authentication: request.Requires3DS}
	case paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeBankTransfer:
		providerRequest.BankTransfer = &coreapi.BankTransferDetails{Bank: midtrans.BankBca}
	case paymentmodel.PaymentMethodTypePaylater:
		providerRequest.ConvStore = &coreapi.ConvStoreDetails{Store: "indomaret", Message: fallbackString(request.Description, request.PaymentCode)}
	}

	return providerRequest, nil
}

func buildCoreAPIClient(serverKey string, env midtrans.EnvironmentType, timeout time.Duration) midtransCoreClient {
	client := &coreapi.Client{}
	client.New(serverKey, env)
	client.HttpClient = &midtrans.HttpClientImplementation{
		HttpClient: &http.Client{Timeout: timeout},
		Logger:     midtrans.GetDefaultLogger(env),
	}
	return &coreAPIChargeClient{client: client}
}

func (g *Gateway) CheckTransactionStatus(ctx context.Context, reference string) (paymentmodel.ChargeResult, error) {
	if !g.enabled {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans gateway is disabled")
	}
	if g.serverKey == "" {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans server key is not configured")
	}
	if strings.TrimSpace(reference) == "" {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans reference is required")
	}
	if g.coreClient == nil {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans charge client is not initialized")
	}

	providerResponse, sdkErr := g.coreClient.CheckTransaction(reference)
	if sdkErr != nil {
		failedResult := chargeResultFromSDKError(paymentmodel.ChargeRequest{PaymentCode: reference}, nil, sdkErr)
		if sdkErr.StatusCode >= http.StatusBadRequest {
			return failedResult, nil
		}
		return failedResult, fmt.Errorf("midtrans status check failed: %w", sdkErr)
	}
	if providerResponse == nil {
		return paymentmodel.ChargeResult{}, fmt.Errorf("midtrans status check returned empty response")
	}

	providerResponseJSON, err := json.Marshal(providerResponse)
	if err != nil {
		return paymentmodel.ChargeResult{}, err
	}

	status, requiresAction, transactionID, resolvedReference, failureCode, failureMessage := interpretMidtransTransactionStatusResponse(providerResponse)
	if status == "" {
		status = paymentmodel.PaymentStatusPending
	}
	if transactionID == "" {
		transactionID = buildTransactionID("MT")
	}
	if resolvedReference == "" {
		resolvedReference = reference
	}

	result := paymentmodel.ChargeResult{
		PSPTransactionID: transactionID,
		PSPReference:     resolvedReference,
		RawResponse:      providerResponseJSON,
		NextStatus:       status,
		RequiresAction:   requiresAction,
		FailureCode:      failureCode,
		FailureMessage:   failureMessage,
	}
	return result, nil
}

func interpretMidtransTransactionStatusResponse(response *coreapi.TransactionStatusResponse) (paymentmodel.PaymentStatus, bool, string, string, string, string) {
	if response == nil {
		return paymentmodel.PaymentStatusFailed, false, "", "", "empty_response", "midtrans returned empty response"
	}

	transactionStatus := strings.ToLower(strings.TrimSpace(response.TransactionStatus))
	fraudStatus := strings.ToLower(strings.TrimSpace(response.FraudStatus))

	switch transactionStatus {
	case "capture":
		if fraudStatus == "challenge" {
			return paymentmodel.PaymentStatusPending, true, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
		}
		if fraudStatus == "deny" {
			return paymentmodel.PaymentStatusFailed, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
		}
		return paymentmodel.PaymentStatusCaptured, false, response.TransactionID, response.OrderID, "", ""
	case "settlement", "completed":
		return paymentmodel.PaymentStatusCaptured, false, response.TransactionID, response.OrderID, "", ""
	case "authorize", "authorized":
		return paymentmodel.PaymentStatusAuthorised, true, response.TransactionID, response.OrderID, "", ""
	case "pending":
		return paymentmodel.PaymentStatusPending, true, response.TransactionID, response.OrderID, "", ""
	case "cancel", "cancelled", "canceled":
		return paymentmodel.PaymentStatusCancelled, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
	case "expire", "expired":
		return paymentmodel.PaymentStatusExpired, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
	case "deny", "failed", "failure":
		return paymentmodel.PaymentStatusFailed, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
	default:
		return paymentmodel.PaymentStatusPending, true, response.TransactionID, response.OrderID, "", ""
	}
}

func resolveMidtransEnvironment(baseURL string) midtrans.EnvironmentType {
	normalized := strings.ToLower(strings.TrimSpace(baseURL))
	if strings.Contains(normalized, "sandbox") {
		return midtrans.Sandbox
	}
	return midtrans.Production
}

func decimalAmountToInt64(amount decimal.Decimal) (int64, error) {
	if amount.LessThan(decimal.Zero) {
		return 0, fmt.Errorf("amount must be greater than or equal to zero")
	}
	if !amount.Equal(amount.Truncate(0)) {
		return 0, fmt.Errorf("midtrans requires integer amount without fractional digits")
	}
	return amount.IntPart(), nil
}

func midtransPaymentType(method paymentmodel.PaymentMethodType) coreapi.CoreapiPaymentType {
	switch method {
	case paymentmodel.PaymentMethodTypeCard:
		return coreapi.PaymentTypeCreditCard
	case paymentmodel.PaymentMethodTypeVirtualAccount, paymentmodel.PaymentMethodTypeBankTransfer:
		return coreapi.PaymentTypeBankTransfer
	case paymentmodel.PaymentMethodTypePaylater:
		return coreapi.PaymentTypeConvenienceStore
	default:
		return coreapi.CoreapiPaymentType(strings.ToLower(string(method)))
	}
}

func interpretMidtransChargeResponse(response *coreapi.ChargeResponse) (paymentmodel.PaymentStatus, bool, string, string, string, string) {
	if response == nil {
		return paymentmodel.PaymentStatusFailed, false, "", "", "empty_response", "midtrans returned empty response"
	}

	transactionStatus := strings.ToLower(strings.TrimSpace(response.TransactionStatus))
	fraudStatus := strings.ToLower(strings.TrimSpace(response.FraudStatus))

	switch transactionStatus {
	case "capture":
		if fraudStatus == "challenge" {
			return paymentmodel.PaymentStatusPending, true, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
		}
		if fraudStatus == "deny" {
			return paymentmodel.PaymentStatusFailed, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
		}
		return paymentmodel.PaymentStatusCaptured, false, response.TransactionID, response.OrderID, "", ""
	case "settlement", "completed":
		return paymentmodel.PaymentStatusCaptured, false, response.TransactionID, response.OrderID, "", ""
	case "authorize", "authorized":
		return paymentmodel.PaymentStatusAuthorised, true, response.TransactionID, response.OrderID, "", ""
	case "pending":
		return paymentmodel.PaymentStatusPending, true, response.TransactionID, response.OrderID, "", ""
	case "cancel", "cancelled", "canceled":
		return paymentmodel.PaymentStatusCancelled, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
	case "expire", "expired":
		return paymentmodel.PaymentStatusExpired, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
	case "deny", "failed", "failure":
		return paymentmodel.PaymentStatusFailed, false, response.TransactionID, response.OrderID, response.StatusCode, response.StatusMessage
	default:
		return paymentmodel.PaymentStatusPending, true, response.TransactionID, response.OrderID, "", ""
	}
}

func chargeResultFromSDKError(request paymentmodel.ChargeRequest, rawRequest []byte, sdkErr *midtrans.Error) paymentmodel.ChargeResult {
	responseBody := []byte(nil)
	if sdkErr != nil && sdkErr.RawApiResponse != nil {
		responseBody = sdkErr.RawApiResponse.RawBody
	}

	statusCode := 0
	message := "midtrans charge request failed"
	if sdkErr != nil {
		statusCode = sdkErr.StatusCode
		if strings.TrimSpace(sdkErr.Message) != "" {
			message = sdkErr.Message
		}
	}

	failureCode := "midtrans_error"
	if statusCode > 0 {
		failureCode = fmt.Sprintf("http_%d", statusCode)
	}

	return paymentmodel.ChargeResult{
		PSPTransactionID: buildTransactionID("MT"),
		PSPReference:     request.PaymentCode,
		RawRequest:       rawRequest,
		RawResponse:      responseBody,
		NextStatus:       paymentmodel.PaymentStatusFailed,
		RequiresAction:   false,
		FailureCode:      failureCode,
		FailureMessage:   message,
	}
}

func buildTransactionID(prefix string) string {
	id, _ := uuid.NewV4()
	compact := strings.ReplaceAll(id.String(), "-", "")
	if len(compact) > 10 {
		compact = compact[:10]
	}
	return fmt.Sprintf("%s-%s", prefix, strings.ToUpper(compact))
}
