package dto

import (
	"time"
)

// PaymentFlowQuoteResponseDTO represents the HTTP response for payment quote requests.
// @Description Payment flow quote containing routing decision and payment details.
type PaymentFlowQuoteResponseDTO struct {
	// Payment intent details including amount, currency, and customer information.
	Intent *PaymentIntentResponseDTO `json:"intent,omitempty"`

	// Payment attempt details including PSP selection and transaction status.
	Payment *PaymentResponseDTO `json:"payment,omitempty"`

	// Routing decision showing which PSP was selected and why.
	Decision RoutingDecisionResponseDTO `json:"decision"`

	// Flow steps for observability and debugging.
	Steps []FlowStepResponseDTO `json:"steps,omitempty"`

	// ISO timestamp when the quote was generated.
	CreatedAt time.Time `json:"createdAt" example:"2026-04-11T13:21:20Z"`
}

// PaymentExecutionResponseDTO represents the HTTP response for executed payments.
// @Description Complete payment execution result including charge attempt and final status.
type PaymentExecutionResponseDTO struct {
	// Quote information similar to quote response.
	Quote PaymentFlowQuoteResponseDTO `json:"quote"`

	// Charge attempt result from the PSP provider.
	Charge ChargeResultResponseDTO `json:"charge"`

	// Final payment status after execution.
	FinalStatus string `json:"finalStatus" example:"FAILED"`

	// State transition path for the payment (e.g., PENDING → FAILED).
	StatePath []string `json:"statePath" example:"PENDING,FAILED"`
}

// PaymentIntentResponseDTO represents a payment intent in the response.
// @Description Payment intent containing order and customer information.
type PaymentIntentResponseDTO struct {
	// Unique identifier for the payment intent.
	ID string `json:"id" example:"1918b8bc-efc2-4e89-8ce9-6f6c11d30279"`

	// Merchant-facing transaction code for easy identification.
	IntentCode string `json:"intentCode" example:"PAY-20260411-12DA96"`

	// Total amount in smallest currency unit.
	Amount string `json:"amount" example:"350000"`

	// ISO 4217 currency code.
	Currency string `json:"currency" example:"IDR"`

	// Payment method type selected.
	PaymentMethodType string `json:"paymentMethodType" example:"CARD"`

	// Expiration timestamp for this intent.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`

	// Whether 3D Secure is required.
	Requires3DS bool `json:"requires3ds" example:"true"`

	// Current status (e.g., PENDING, AUTHORIZED, CAPTURED, FAILED).
	Status string `json:"status" example:"PENDING"`

	// Description provided by merchant.
	Description string `json:"description,omitempty" example:"2x VIP event ticket"`

	// Additional metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// ISO timestamp when intent was created.
	CreatedAt time.Time `json:"createdAt" example:"2026-04-11T13:21:20Z"`
}

// PaymentResponseDTO represents a payment attempt in the response.
// @Description Payment attempt with PSP details and transaction information.
type PaymentResponseDTO struct {
	// Unique identifier for this payment attempt.
	ID string `json:"id" example:"618e188d-6998-4a60-9f9d-9506aab3269d"`

	// Human-readable transaction code.
	PaymentCode string `json:"paymentCode" example:"TXN-20260411-09D1AB"`

	// Selected PSP provider (e.g., MIDTRANS, XENDIT).
	PSP string `json:"psp" example:"MIDTRANS"`

	// Transaction ID from the PSP provider.
	PSPTransactionID *string `json:"pspTransactionId,omitempty" example:"txn-12345"`

	// Amount processed by the PSP.
	Amount string `json:"amount" example:"350000"`

	// Currency of the transaction.
	Currency string `json:"currency" example:"IDR"`

	// Current payment status.
	Status string `json:"status" example:"FAILED"`

	// Failure code if applicable (e.g., "INVALID_KEY", "DECLINED").
	FailureCode *string `json:"failureCode,omitempty" example:"INVALID_KEY"`

	// Human-readable failure message.
	FailureMessage *string `json:"failureMessage,omitempty" example:"Unknown Merchant server_key/id"`

	// Raw request sent to PSP (for debugging).
	PSPRawRequest map[string]interface{} `json:"pspRawRequest,omitempty"`

	// Raw response received from PSP (for debugging).
	PSPRawResponse map[string]interface{} `json:"pspRawResponse,omitempty"`

	// When the payment was authorized (if applicable).
	AuthorisedAt *time.Time `json:"authorisedAt,omitempty"`

	// When the payment was captured/settled (if applicable).
	CapturedAt *time.Time `json:"capturedAt,omitempty"`

	// When the payment record was created.
	CreatedAt time.Time `json:"createdAt" example:"2026-04-11T13:21:20Z"`
}

// RoutingDecisionResponseDTO represents the routing decision in the response.
// @Description Routing decision showing PSP selection rationale.
type RoutingDecisionResponseDTO struct {
	// Routing strategy used (WATERFALL, ROUND_ROBIN, etc.).
	Strategy string `json:"strategy" example:"WATERFALL"`

	// Name of the policy that applied.
	PolicyName string `json:"policyName,omitempty" example:"event-payments"`

	// Selected PSP.
	PSP string `json:"psp" example:"MIDTRANS"`

	// PSP account ID used.
	PSPAccountID string `json:"pspAccountId,omitempty" example:"11111111-1111-1111-1111-111111111111"`

	// PSP account label for human reference.
	PSPAccountLabel string `json:"pspAccountLabel,omitempty" example:"Midtrans Primary"`

	// Reason for selection or fallback.
	Reason string `json:"reason,omitempty" example:"matched event-payments policy"`

	// Whether database fallback was used.
	FallbackUsed bool `json:"fallbackUsed" example:"false"`

	// List of candidate PSPs considered (for observability).
	Candidates []RoutingCandidateResponseDTO `json:"candidates,omitempty"`
}

// RoutingCandidateResponseDTO represents a candidate PSP in routing decisions.
// @Description Candidate PSP with score and evaluation reason.
type RoutingCandidateResponseDTO struct {
	// PSP identifier.
	PSP string `json:"psp" example:"MIDTRANS"`

	// Account ID for this PSP.
	AccountID string `json:"accountId" example:"11111111-1111-1111-1111-111111111111"`

	// Account label.
	AccountLabel string `json:"accountLabel" example:"Midtrans Primary"`

	// Evaluation score (higher is better).
	Score int `json:"score" example:"100"`

	// Reason for this score.
	Reason string `json:"reason,omitempty" example:"policy match + success rate"`
}

// FlowStepResponseDTO represents a step in the payment flow for observability.
// @Description Flow step for tracing payment processing.
type FlowStepResponseDTO struct {
	// Step name for identification.
	Name string `json:"name" example:"routing"`

	// Description of what happened in this step.
	Description string `json:"description,omitempty" example:"Evaluated event-payments policy, selected MIDTRANS"`
}

// ChargeResultResponseDTO represents the result of a charge attempt.
// @Description Charge attempt result from PSP gateway.
type ChargeResultResponseDTO struct {
	// Transaction ID from the PSP.
	TransactionID string `json:"transactionId,omitempty" example:"txn-12345"`

	// Reference ID from the PSP (for customer support).
	ReferenceID string `json:"referenceId,omitempty" example:"ref-12345"`

	// Standard payment status.
	Status string `json:"status" example:"FAILED"`

	// Authorization code (if authorized).
	AuthCode *string `json:"authCode,omitempty" example:"AUTH123"`

	// Processing fee charged by PSP.
	Fee *string `json:"fee,omitempty" example:"1500"`

	// Currency of the fee.
	FeeCurrency *string `json:"feeCurrency,omitempty" example:"IDR"`

	// Error code if charge failed.
	ErrorCode *string `json:"errorCode,omitempty" example:"INVALID_KEY"`

	// Error message if charge failed.
	ErrorMessage *string `json:"errorMessage,omitempty" example:"Unknown Merchant server_key/id"`

	// Raw PSP response for detailed debugging.
	RawResponse map[string]interface{} `json:"rawResponse,omitempty"`

	// ISO timestamp of charge attempt.
	Timestamp time.Time `json:"timestamp" example:"2026-04-11T13:21:20Z"`
}

// WebhookEventResponseDTO represents the response to a webhook event.
// @Description Acknowledgment of webhook receipt and processing outcome.
type WebhookEventResponseDTO struct {
	// Provider that sent the webhook.
	Provider string `json:"provider" example:"MIDTRANS"`

	// Whether the webhook was successfully processed.
	Success bool `json:"success" example:"true"`

	// Unique identifier for this webhook processing.
	EventID string `json:"eventId" example:"evt-12345"`

	// Transaction ID from the PSP provider.
	TransactionID string `json:"transactionId,omitempty" example:"txn-12345"`

	// Payment identifier updated by the webhook.
	PaymentID string `json:"paymentId,omitempty" example:"618e188d-6998-4a60-9f9d-9506aab3269d"`

	// Payment intent identifier updated by the webhook.
	IntentID string `json:"intentId,omitempty" example:"1918b8bc-efc2-4e89-8ce9-6f6c11d30279"`

	// Final normalized payment status after processing.
	FinalStatus string `json:"finalStatus,omitempty" example:"PENDING"`

	// Message describing the processing result.
	Message string `json:"message" example:"Webhook processed successfully"`

	// ISO timestamp when the webhook was processed.
	ProcessedAt time.Time `json:"processedAt" example:"2026-04-11T13:21:20Z"`

	// If an error occurred, the error details.
	Error *WebhookErrorDTO `json:"error,omitempty"`
}

// WebhookErrorDTO represents an error in webhook processing.
// @Description Webhook processing error details.
type WebhookErrorDTO struct {
	// Error code for programmatic handling.
	Code string `json:"code" example:"SIGNATURE_VERIFICATION_FAILED"`

	// Human-readable error message.
	Message string `json:"message" example:"Webhook signature verification failed"`

	// Request ID for support reference.
	RequestID string `json:"requestId,omitempty" example:"req-12345"`
}

// ErrorResponseDTO represents an error response.
// @Description Standard error response format.
type ErrorResponseDTO struct {
	// HTTP status code.
	StatusCode int `json:"statusCode" example:"400"`

	// Error code for programmatic handling.
	ErrorCode string `json:"errorCode" example:"INVALID_REQUEST"`

	// Human-readable error message.
	Message string `json:"message" example:"Invalid payment amount"`

	// Timestamp of the error.
	Timestamp time.Time `json:"timestamp" example:"2026-04-11T13:21:20Z"`

	// Request ID for support reference.
	RequestID string `json:"requestId,omitempty" example:"req-12345"`

	// Additional error details.
	Details map[string]interface{} `json:"details,omitempty"`
}
