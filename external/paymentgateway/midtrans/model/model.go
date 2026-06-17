package model

// ChargeRequest is a minimal Midtrans Core API charge request.
type ChargeRequest struct {
	PaymentType       string            `json:"payment_type"`
	TransactionDetail TransactionDetail `json:"transaction_details"`
	CustomerDetails   *CustomerDetails  `json:"customer_details,omitempty"`
	ItemDetails       []ItemDetail      `json:"item_details,omitempty"`
	BankTransfer      *BankTransfer     `json:"bank_transfer,omitempty"`
	Gopay             *Gopay            `json:"gopay,omitempty"`
	Qris              *Qris             `json:"qris,omitempty"`
	Expiry            *Expiry           `json:"custom_expiry,omitempty"`
	CustomField1      string            `json:"custom_field1,omitempty"`
	CustomField2      string            `json:"custom_field2,omitempty"`
	CustomField3      string            `json:"custom_field3,omitempty"`
}

type TransactionDetail struct {
	OrderID     string  `json:"order_id"`
	GrossAmount float64 `json:"gross_amount"`
}

type CustomerDetails struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type ItemDetail struct {
	ID       string  `json:"id,omitempty"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
	Name     string  `json:"name"`
	Category string  `json:"category,omitempty"`
}

type BankTransfer struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number,omitempty"`
}

type Gopay struct {
	EnableCallback bool   `json:"enable_callback,omitempty"`
	CallbackURL    string `json:"callback_url,omitempty"`
}

type Qris struct {
	Acquirer string `json:"acquirer,omitempty"`
}

type Expiry struct {
	OrderTime      string `json:"order_time,omitempty"`
	ExpiryDuration int64  `json:"expiry_duration,omitempty"`
	Unit           string `json:"unit,omitempty"`
}

type ChargeResponse struct {
	StatusCode        string          `json:"status_code"`
	StatusMessage     string          `json:"status_message"`
	TransactionID     string          `json:"transaction_id"`
	OrderID           string          `json:"order_id"`
	MerchantID        string          `json:"merchant_id"`
	GrossAmount       string          `json:"gross_amount"`
	Currency          string          `json:"currency"`
	PaymentType       string          `json:"payment_type"`
	TransactionTime   string          `json:"transaction_time"`
	TransactionStatus string          `json:"transaction_status"`
	FraudStatus       string          `json:"fraud_status"`
	VAStatus          string          `json:"va_status,omitempty"`
	PermataVANumber   string          `json:"permata_va_number,omitempty"`
	VANumbers         []VANumber      `json:"va_numbers,omitempty"`
	Actions           []PaymentAction `json:"actions,omitempty"`
	QRString          string          `json:"qr_string,omitempty"`
}

type VANumber struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

type PaymentAction struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

type StatusResponse = ChargeResponse

type RefundRequest struct {
	RefundKey string  `json:"refund_key"`
	Amount    float64 `json:"amount"`
	Reason    string  `json:"reason,omitempty"`
}

type RefundResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	OrderID           string `json:"order_id"`
	TransactionID     string `json:"transaction_id"`
	RefundKey         string `json:"refund_key"`
	RefundAmount      string `json:"refund_amount"`
	TransactionStatus string `json:"transaction_status"`
}

type Notification struct {
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	StatusMessage     string `json:"status_message"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	Currency          string `json:"currency"`
}
