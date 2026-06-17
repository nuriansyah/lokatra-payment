package model

type CreatePaymentRequest struct {
	ReferenceID   string               `json:"reference_id"`
	Currency      string               `json:"currency"`
	Amount        float64              `json:"amount"`
	Country       string               `json:"country,omitempty"`
	PaymentMethod PaymentMethodRequest `json:"payment_method"`
	Customer      *Customer            `json:"customer,omitempty"`
	Description   string               `json:"description,omitempty"`
	Metadata      map[string]string    `json:"metadata,omitempty"`
	CaptureMethod string               `json:"capture_method,omitempty"`
	Initiator     string               `json:"initiator,omitempty"`
	ChannelProps  map[string]any       `json:"channel_properties,omitempty"`
}

type PaymentMethodRequest struct {
	Type           string          `json:"type"`
	Reusability    string          `json:"reusability,omitempty"`
	ReferenceID    string          `json:"reference_id,omitempty"`
	EWallet        *EWallet        `json:"ewallet,omitempty"`
	VirtualAccount *VirtualAccount `json:"virtual_account,omitempty"`
	QR             *QR             `json:"qr,omitempty"`
	RetailOutlet   *RetailOutlet   `json:"retail_outlet,omitempty"`
}

type EWallet struct {
	ChannelCode string `json:"channel_code"`
}
type VirtualAccount struct {
	ChannelCode       string         `json:"channel_code"`
	ChannelProperties map[string]any `json:"channel_properties,omitempty"`
}
type QR struct {
	ChannelCode string `json:"channel_code"`
}
type RetailOutlet struct {
	ChannelCode string `json:"channel_code"`
}

type Customer struct {
	ReferenceID      string            `json:"reference_id,omitempty"`
	Type             string            `json:"type,omitempty"`
	Email            string            `json:"email,omitempty"`
	MobileNumber     string            `json:"mobile_number,omitempty"`
	IndividualDetail *IndividualDetail `json:"individual_detail,omitempty"`
}

type IndividualDetail struct {
	GivenNames string `json:"given_names,omitempty"`
	Surname    string `json:"surname,omitempty"`
}

type PaymentRequestResponse struct {
	ID            string                `json:"id"`
	ReferenceID   string                `json:"reference_id"`
	Status        string                `json:"status"`
	Currency      string                `json:"currency"`
	Amount        float64               `json:"amount"`
	Country       string                `json:"country"`
	PaymentMethod PaymentMethodResponse `json:"payment_method"`
	Actions       []Action              `json:"actions,omitempty"`
	Metadata      map[string]string     `json:"metadata,omitempty"`
	Created       string                `json:"created"`
	Updated       string                `json:"updated"`
}

type PaymentMethodResponse struct {
	ID             string                  `json:"id"`
	Type           string                  `json:"type"`
	Status         string                  `json:"status"`
	VirtualAccount *VirtualAccountResponse `json:"virtual_account,omitempty"`
	QR             *QRResponse             `json:"qr,omitempty"`
	EWallet        *EWalletResponse        `json:"ewallet,omitempty"`
	RetailOutlet   *RetailOutletResponse   `json:"retail_outlet,omitempty"`
}

type VirtualAccountResponse struct {
	ChannelCode       string       `json:"channel_code"`
	ChannelProperties VAProperties `json:"channel_properties"`
}
type VAProperties struct {
	CustomerName         string `json:"customer_name,omitempty"`
	VirtualAccountNumber string `json:"virtual_account_number,omitempty"`
	ExpiresAt            string `json:"expires_at,omitempty"`
}
type QRResponse struct {
	ChannelCode       string       `json:"channel_code"`
	ChannelProperties QRProperties `json:"channel_properties"`
}
type QRProperties struct {
	QRString  string `json:"qr_string,omitempty"`
	ExpiresAt string `json:"expires_at,omitempty"`
}
type EWalletResponse struct {
	ChannelCode       string         `json:"channel_code"`
	ChannelProperties map[string]any `json:"channel_properties,omitempty"`
}
type RetailOutletResponse struct {
	ChannelCode       string         `json:"channel_code"`
	ChannelProperties map[string]any `json:"channel_properties,omitempty"`
}

type Action struct {
	Action  string `json:"action"`
	Method  string `json:"method"`
	URL     string `json:"url"`
	URLType string `json:"url_type"`
}

type RefundRequest struct {
	ReferenceID      string  `json:"reference_id"`
	PaymentRequestID string  `json:"payment_request_id,omitempty"`
	PaymentID        string  `json:"payment_id,omitempty"`
	Currency         string  `json:"currency"`
	Amount           float64 `json:"amount"`
	Reason           string  `json:"reason,omitempty"`
}

type RefundResponse struct {
	ID          string  `json:"id"`
	ReferenceID string  `json:"reference_id"`
	Status      string  `json:"status"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

type PayoutRequest struct {
	ReferenceID       string            `json:"reference_id"`
	ChannelCode       string            `json:"channel_code"`
	ChannelProperties map[string]any    `json:"channel_properties"`
	Amount            float64           `json:"amount"`
	Currency          string            `json:"currency"`
	Description       string            `json:"description,omitempty"`
	Metadata          map[string]string `json:"metadata,omitempty"`
}
type PayoutResponse struct {
	ID          string  `json:"id"`
	ReferenceID string  `json:"reference_id"`
	Status      string  `json:"status"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

type WebhookPayment struct {
	ID         string                 `json:"id"`
	Event      string                 `json:"event"`
	BusinessID string                 `json:"business_id"`
	Created    string                 `json:"created"`
	Data       PaymentRequestResponse `json:"data"`
}
