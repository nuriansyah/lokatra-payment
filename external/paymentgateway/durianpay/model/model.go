package model

type CreatePaymentRequest struct {
	Amount         string            `json:"amount"`
	Currency       string            `json:"currency"`
	OrderRefID     string            `json:"order_ref_id"`
	Customer       *Customer         `json:"customer,omitempty"`
	Items          []Item            `json:"items,omitempty"`
	PaymentMethod  string            `json:"payment_method,omitempty"`
	PaymentChannel string            `json:"payment_channel,omitempty"`
	Description    string            `json:"description,omitempty"`
	CallbackURL    string            `json:"callback_url,omitempty"`
	ReturnURL      string            `json:"return_url,omitempty"`
	ExpiryAt       string            `json:"expiry_at,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
}

type Customer struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}
type Item struct {
	Name  string `json:"name"`
	Qty   int64  `json:"qty"`
	Price string `json:"price"`
	SKU   string `json:"sku,omitempty"`
}

type PaymentResponse struct {
	ID             string            `json:"id"`
	OrderRefID     string            `json:"order_ref_id"`
	Status         string            `json:"status"`
	Amount         string            `json:"amount"`
	Currency       string            `json:"currency"`
	PaymentURL     string            `json:"payment_url"`
	CheckoutURL    string            `json:"checkout_url"`
	VA             string            `json:"va_number"`
	PaymentCode    string            `json:"payment_code"`
	QRString       string            `json:"qr_string"`
	PaymentMethod  string            `json:"payment_method"`
	PaymentChannel string            `json:"payment_channel"`
	Metadata       map[string]string `json:"metadata,omitempty"`
}

type RefundRequest struct {
	PaymentID   string `json:"payment_id"`
	RefundRefID string `json:"refund_ref_id"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Reason      string `json:"reason,omitempty"`
}
type RefundResponse struct {
	ID          string `json:"id"`
	RefundRefID string `json:"refund_ref_id"`
	PaymentID   string `json:"payment_id"`
	Status      string `json:"status"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
}
type PayoutRequest struct {
	ReferenceID   string            `json:"reference_id"`
	Amount        string            `json:"amount"`
	Currency      string            `json:"currency"`
	BankCode      string            `json:"bank_code"`
	AccountNumber string            `json:"account_number"`
	AccountName   string            `json:"account_name"`
	Description   string            `json:"description,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}
type PayoutResponse struct {
	ID          string `json:"id"`
	ReferenceID string `json:"reference_id"`
	Status      string `json:"status"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
}

type Webhook struct {
	ID        string          `json:"id"`
	Event     string          `json:"event"`
	Data      PaymentResponse `json:"data"`
	CreatedAt string          `json:"created_at"`
}
