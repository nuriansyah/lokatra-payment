package model

type PaymentRequest struct {
	MerchantID      string            `json:"merchantId"`
	MerchantTransID string            `json:"merchantTransId"`
	Amount          string            `json:"amount"`
	Currency        string            `json:"currency"`
	PaymentChannel  string            `json:"paymentChannel"`
	Description     string            `json:"description,omitempty"`
	CustomerName    string            `json:"customerName,omitempty"`
	CustomerEmail   string            `json:"customerEmail,omitempty"`
	CustomerPhone   string            `json:"customerPhone,omitempty"`
	CallbackURL     string            `json:"callbackUrl,omitempty"`
	ReturnURL       string            `json:"returnUrl,omitempty"`
	ExpiryTime      string            `json:"expiryTime,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
}

type PaymentResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	TransactionID   string `json:"transactionId"`
	MerchantTransID string `json:"merchantTransId"`
	PaymentChannel  string `json:"paymentChannel"`
	PaymentCode     string `json:"paymentCode"`
	PaymentURL      string `json:"paymentUrl"`
	QRString        string `json:"qrString"`
	Status          string `json:"status"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
}

type StatusRequest struct {
	MerchantID      string `json:"merchantId"`
	MerchantTransID string `json:"merchantTransId"`
	TransactionID   string `json:"transactionId,omitempty"`
}
type RefundRequest struct {
	MerchantID      string `json:"merchantId"`
	MerchantTransID string `json:"merchantTransId"`
	RefundID        string `json:"refundId"`
	Amount          string `json:"amount"`
	Reason          string `json:"reason,omitempty"`
}
type RefundResponse = PaymentResponse

type Notification struct {
	TransactionID   string `json:"transactionId"`
	MerchantTransID string `json:"merchantTransId"`
	PaymentChannel  string `json:"paymentChannel"`
	Status          string `json:"status"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	Signature       string `json:"signature"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}
