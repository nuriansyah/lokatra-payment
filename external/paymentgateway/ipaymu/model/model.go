package model

type RedirectPaymentRequest struct {
	Product        []string `json:"product"`
	Qty            []int64  `json:"qty"`
	Price          []string `json:"price"`
	ReturnURL      string   `json:"returnUrl,omitempty"`
	CancelURL      string   `json:"cancelUrl,omitempty"`
	NotifyURL      string   `json:"notifyUrl,omitempty"`
	ReferenceID    string   `json:"referenceId"`
	BuyerName      string   `json:"buyerName,omitempty"`
	BuyerEmail     string   `json:"buyerEmail,omitempty"`
	BuyerPhone     string   `json:"buyerPhone,omitempty"`
	PaymentMethod  string   `json:"paymentMethod,omitempty"`
	PaymentChannel string   `json:"paymentChannel,omitempty"`
	Expired        string   `json:"expired,omitempty"`
}

type PaymentResponse struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    PaymentData `json:"Data"`
}

type PaymentData struct {
	SessionID     string `json:"SessionID"`
	TransactionID string `json:"TransactionId"`
	ReferenceID   string `json:"ReferenceId"`
	URL           string `json:"Url"`
	PaymentNo     string `json:"PaymentNo"`
	QRString      string `json:"QrString"`
	Total         string `json:"Total"`
	Status        string `json:"Status"`
}

type StatusRequest struct {
	TransactionID string `json:"transactionId,omitempty"`
	ReferenceID   string `json:"referenceId,omitempty"`
	Account       string `json:"account,omitempty"`
}
type RefundRequest struct {
	TransactionID string `json:"transactionId"`
	ReferenceID   string `json:"referenceId"`
	Amount        string `json:"amount"`
	Reason        string `json:"reason,omitempty"`
}
type RefundResponse = PaymentResponse

type Webhook struct {
	TrxID         string `json:"trx_id"`
	Sid           string `json:"sid"`
	ReferenceID   string `json:"reference_id"`
	Status        string `json:"status"`
	StatusCode    string `json:"status_code"`
	Amount        string `json:"amount"`
	Fee           string `json:"fee"`
	Channel       string `json:"channel"`
	PaymentMethod string `json:"payment_method"`
	Signature     string `json:"signature"`
}
