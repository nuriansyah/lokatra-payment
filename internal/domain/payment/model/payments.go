package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type PaymentsDBFieldNameType string

type paymentsDBFieldName struct {
	Id                         PaymentsDBFieldNameType
	PaymentCode                PaymentsDBFieldNameType
	IntentId                   PaymentsDBFieldNameType
	AttemptNumber              PaymentsDBFieldNameType
	Psp                        PaymentsDBFieldNameType
	PspTransactionId           PaymentsDBFieldNameType
	PspReference               PaymentsDBFieldNameType
	PspRawRequest              PaymentsDBFieldNameType
	PspRawResponse             PaymentsDBFieldNameType
	Amount                     PaymentsDBFieldNameType
	Currency                   PaymentsDBFieldNameType
	AmountInSettlementCurrency PaymentsDBFieldNameType
	SettlementCurrency         PaymentsDBFieldNameType
	FxRate                     PaymentsDBFieldNameType
	FxRateSnapshotId           PaymentsDBFieldNameType
	PaymentMethodId            PaymentsDBFieldNameType
	PaymentMethodType          PaymentsDBFieldNameType
	Status                     PaymentsDBFieldNameType
	FailureCode                PaymentsDBFieldNameType
	FailureMessage             PaymentsDBFieldNameType
	FailureCategory            PaymentsDBFieldNameType
	AuthorisedAt               PaymentsDBFieldNameType
	AuthorisedAmount           PaymentsDBFieldNameType
	CapturedAt                 PaymentsDBFieldNameType
	CapturedAmount             PaymentsDBFieldNameType
	ProcessingFee              PaymentsDBFieldNameType
	ProcessingFeeCurrency      PaymentsDBFieldNameType
	RiskScoreId                PaymentsDBFieldNameType
	Description                PaymentsDBFieldNameType
	Metadata                   PaymentsDBFieldNameType
	CompletedAt                PaymentsDBFieldNameType
	CancelledAt                PaymentsDBFieldNameType
	ExpiredAt                  PaymentsDBFieldNameType
	MetaCreatedAt              PaymentsDBFieldNameType
	MetaCreatedBy              PaymentsDBFieldNameType
	MetaUpdatedAt              PaymentsDBFieldNameType
	MetaUpdatedBy              PaymentsDBFieldNameType
	MetaDeletedAt              PaymentsDBFieldNameType
	MetaDeletedBy              PaymentsDBFieldNameType
}

var PaymentsDBFieldName = paymentsDBFieldName{
	Id:                         "id",
	PaymentCode:                "payment_code",
	IntentId:                   "intent_id",
	AttemptNumber:              "attempt_number",
	Psp:                        "psp",
	PspTransactionId:           "psp_transaction_id",
	PspReference:               "psp_reference",
	PspRawRequest:              "psp_raw_request",
	PspRawResponse:             "psp_raw_response",
	Amount:                     "amount",
	Currency:                   "currency",
	AmountInSettlementCurrency: "amount_in_settlement_currency",
	SettlementCurrency:         "settlement_currency",
	FxRate:                     "fx_rate",
	FxRateSnapshotId:           "fx_rate_snapshot_id",
	PaymentMethodId:            "payment_method_id",
	PaymentMethodType:          "payment_method_type",
	Status:                     "status",
	FailureCode:                "failure_code",
	FailureMessage:             "failure_message",
	FailureCategory:            "failure_category",
	AuthorisedAt:               "authorised_at",
	AuthorisedAmount:           "authorised_amount",
	CapturedAt:                 "captured_at",
	CapturedAmount:             "captured_amount",
	ProcessingFee:              "processing_fee",
	ProcessingFeeCurrency:      "processing_fee_currency",
	RiskScoreId:                "risk_score_id",
	Description:                "description",
	Metadata:                   "metadata",
	CompletedAt:                "completed_at",
	CancelledAt:                "cancelled_at",
	ExpiredAt:                  "expired_at",
	MetaCreatedAt:              "meta_created_at",
	MetaCreatedBy:              "meta_created_by",
	MetaUpdatedAt:              "meta_updated_at",
	MetaUpdatedBy:              "meta_updated_by",
	MetaDeletedAt:              "meta_deleted_at",
	MetaDeletedBy:              "meta_deleted_by",
}

func NewPaymentsDBFieldNameFromStr(field string) (dbField PaymentsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentsDBFieldName.Id):
		return PaymentsDBFieldName.Id, true

	case string(PaymentsDBFieldName.PaymentCode):
		return PaymentsDBFieldName.PaymentCode, true

	case string(PaymentsDBFieldName.IntentId):
		return PaymentsDBFieldName.IntentId, true

	case string(PaymentsDBFieldName.AttemptNumber):
		return PaymentsDBFieldName.AttemptNumber, true

	case string(PaymentsDBFieldName.Psp):
		return PaymentsDBFieldName.Psp, true

	case string(PaymentsDBFieldName.PspTransactionId):
		return PaymentsDBFieldName.PspTransactionId, true

	case string(PaymentsDBFieldName.PspReference):
		return PaymentsDBFieldName.PspReference, true

	case string(PaymentsDBFieldName.PspRawRequest):
		return PaymentsDBFieldName.PspRawRequest, true

	case string(PaymentsDBFieldName.PspRawResponse):
		return PaymentsDBFieldName.PspRawResponse, true

	case string(PaymentsDBFieldName.Amount):
		return PaymentsDBFieldName.Amount, true

	case string(PaymentsDBFieldName.Currency):
		return PaymentsDBFieldName.Currency, true

	case string(PaymentsDBFieldName.AmountInSettlementCurrency):
		return PaymentsDBFieldName.AmountInSettlementCurrency, true

	case string(PaymentsDBFieldName.SettlementCurrency):
		return PaymentsDBFieldName.SettlementCurrency, true

	case string(PaymentsDBFieldName.FxRate):
		return PaymentsDBFieldName.FxRate, true

	case string(PaymentsDBFieldName.FxRateSnapshotId):
		return PaymentsDBFieldName.FxRateSnapshotId, true

	case string(PaymentsDBFieldName.PaymentMethodId):
		return PaymentsDBFieldName.PaymentMethodId, true

	case string(PaymentsDBFieldName.PaymentMethodType):
		return PaymentsDBFieldName.PaymentMethodType, true

	case string(PaymentsDBFieldName.Status):
		return PaymentsDBFieldName.Status, true

	case string(PaymentsDBFieldName.FailureCode):
		return PaymentsDBFieldName.FailureCode, true

	case string(PaymentsDBFieldName.FailureMessage):
		return PaymentsDBFieldName.FailureMessage, true

	case string(PaymentsDBFieldName.FailureCategory):
		return PaymentsDBFieldName.FailureCategory, true

	case string(PaymentsDBFieldName.AuthorisedAt):
		return PaymentsDBFieldName.AuthorisedAt, true

	case string(PaymentsDBFieldName.AuthorisedAmount):
		return PaymentsDBFieldName.AuthorisedAmount, true

	case string(PaymentsDBFieldName.CapturedAt):
		return PaymentsDBFieldName.CapturedAt, true

	case string(PaymentsDBFieldName.CapturedAmount):
		return PaymentsDBFieldName.CapturedAmount, true

	case string(PaymentsDBFieldName.ProcessingFee):
		return PaymentsDBFieldName.ProcessingFee, true

	case string(PaymentsDBFieldName.ProcessingFeeCurrency):
		return PaymentsDBFieldName.ProcessingFeeCurrency, true

	case string(PaymentsDBFieldName.RiskScoreId):
		return PaymentsDBFieldName.RiskScoreId, true

	case string(PaymentsDBFieldName.Description):
		return PaymentsDBFieldName.Description, true

	case string(PaymentsDBFieldName.Metadata):
		return PaymentsDBFieldName.Metadata, true

	case string(PaymentsDBFieldName.CompletedAt):
		return PaymentsDBFieldName.CompletedAt, true

	case string(PaymentsDBFieldName.CancelledAt):
		return PaymentsDBFieldName.CancelledAt, true

	case string(PaymentsDBFieldName.ExpiredAt):
		return PaymentsDBFieldName.ExpiredAt, true

	case string(PaymentsDBFieldName.MetaCreatedAt):
		return PaymentsDBFieldName.MetaCreatedAt, true

	case string(PaymentsDBFieldName.MetaCreatedBy):
		return PaymentsDBFieldName.MetaCreatedBy, true

	case string(PaymentsDBFieldName.MetaUpdatedAt):
		return PaymentsDBFieldName.MetaUpdatedAt, true

	case string(PaymentsDBFieldName.MetaUpdatedBy):
		return PaymentsDBFieldName.MetaUpdatedBy, true

	case string(PaymentsDBFieldName.MetaDeletedAt):
		return PaymentsDBFieldName.MetaDeletedAt, true

	case string(PaymentsDBFieldName.MetaDeletedBy):
		return PaymentsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type PaymentsFilterResult struct {
	Payments
	FilterCount int `db:"count"`
}

func ValidatePaymentsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewPaymentsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewPaymentsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewPaymentsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type PaymentCurrency string

const (
	PaymentCurrencyIdr PaymentCurrency = "IDR"
	PaymentCurrencyUsd PaymentCurrency = "USD"
	PaymentCurrencySgd PaymentCurrency = "SGD"
	PaymentCurrencyMyr PaymentCurrency = "MYR"
	PaymentCurrencyPhp PaymentCurrency = "PHP"
	PaymentCurrencyThb PaymentCurrency = "THB"
	PaymentCurrencyAed PaymentCurrency = "AED"
	PaymentCurrencyEur PaymentCurrency = "EUR"
	PaymentCurrencyGbp PaymentCurrency = "GBP"
	PaymentCurrencyJpy PaymentCurrency = "JPY"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCard           PaymentMethodType = "CARD"
	PaymentMethodTypeVirtualAccount PaymentMethodType = "VIRTUAL_ACCOUNT"
	PaymentMethodTypeQris           PaymentMethodType = "QRIS"
	PaymentMethodTypeEwallet        PaymentMethodType = "EWALLET"
	PaymentMethodTypeDirectDebit    PaymentMethodType = "DIRECT_DEBIT"
	PaymentMethodTypeBankTransfer   PaymentMethodType = "BANK_TRANSFER"
	PaymentMethodTypePaylater       PaymentMethodType = "PAYLATER"
	PaymentMethodTypeVoucher        PaymentMethodType = "VOUCHER"
	PaymentMethodTypePoints         PaymentMethodType = "POINTS"
	PaymentMethodTypeCashOnDelivery PaymentMethodType = "CASH_ON_DELIVERY"
)

type PaymentStatus string

const (
	PaymentStatusInitiated         PaymentStatus = "INITIATED"
	PaymentStatusPending           PaymentStatus = "PENDING"
	PaymentStatusAuthorised        PaymentStatus = "AUTHORISED"
	PaymentStatusCaptured          PaymentStatus = "CAPTURED"
	PaymentStatusPartiallyCaptured PaymentStatus = "PARTIALLY_CAPTURED"
	PaymentStatusCompleted         PaymentStatus = "COMPLETED"
	PaymentStatusFailed            PaymentStatus = "FAILED"
	PaymentStatusCancelled         PaymentStatus = "CANCELLED"
	PaymentStatusExpired           PaymentStatus = "EXPIRED"
	PaymentStatusRefunding         PaymentStatus = "REFUNDING"
	PaymentStatusRefunded          PaymentStatus = "REFUNDED"
	PaymentStatusPartiallyRefunded PaymentStatus = "PARTIALLY_REFUNDED"
	PaymentStatusDisputed          PaymentStatus = "DISPUTED"
	PaymentStatusChargebackWon     PaymentStatus = "CHARGEBACK_WON"
	PaymentStatusChargebackLost    PaymentStatus = "CHARGEBACK_LOST"
)

type Psp string

const (
	PspMidtrans  Psp = "MIDTRANS"
	PspXendit    Psp = "XENDIT"
	PspStripe    Psp = "STRIPE"
	PspDoku      Psp = "DOKU"
	PspDana      Psp = "DANA"
	PspOvo       Psp = "OVO"
	PspGopay     Psp = "GOPAY"
	PspShopeePay Psp = "SHOPEE_PAY"
	PspLinkAja   Psp = "LINK_AJA"
	PspFlip      Psp = "FLIP"
	PspInternal  Psp = "INTERNAL"
)

type Payments struct {
	Id                         uuid.UUID           `db:"id"`
	PaymentCode                string              `db:"payment_code"`
	IntentId                   uuid.UUID           `db:"intent_id"`
	AttemptNumber              int16               `db:"attempt_number"`
	Psp                        Psp                 `db:"psp"`
	PspTransactionId           null.String         `db:"psp_transaction_id"`
	PspReference               null.String         `db:"psp_reference"`
	PspRawRequest              json.RawMessage     `db:"psp_raw_request"`
	PspRawResponse             json.RawMessage     `db:"psp_raw_response"`
	Amount                     decimal.Decimal     `db:"amount"`
	Currency                   PaymentCurrency     `db:"currency"`
	AmountInSettlementCurrency decimal.NullDecimal `db:"amount_in_settlement_currency"`
	SettlementCurrency         PaymentCurrency     `db:"settlement_currency"`
	FxRate                     decimal.NullDecimal `db:"fx_rate"`
	FxRateSnapshotId           nuuid.NUUID         `db:"fx_rate_snapshot_id"`
	PaymentMethodId            nuuid.NUUID         `db:"payment_method_id"`
	PaymentMethodType          PaymentMethodType   `db:"payment_method_type"`
	Status                     PaymentStatus       `db:"status"`
	FailureCode                null.String         `db:"failure_code"`
	FailureMessage             null.String         `db:"failure_message"`
	FailureCategory            null.String         `db:"failure_category"`
	AuthorisedAt               null.Time           `db:"authorised_at"`
	AuthorisedAmount           decimal.NullDecimal `db:"authorised_amount"`
	CapturedAt                 null.Time           `db:"captured_at"`
	CapturedAmount             decimal.NullDecimal `db:"captured_amount"`
	ProcessingFee              decimal.NullDecimal `db:"processing_fee"`
	ProcessingFeeCurrency      PaymentCurrency     `db:"processing_fee_currency"`
	RiskScoreId                nuuid.NUUID         `db:"risk_score_id"`
	Description                null.String         `db:"description"`
	Metadata                   json.RawMessage     `db:"metadata"`
	CompletedAt                null.Time           `db:"completed_at"`
	CancelledAt                null.Time           `db:"cancelled_at"`
	ExpiredAt                  null.Time           `db:"expired_at"`
	MetaCreatedAt              time.Time           `db:"meta_created_at"`
	MetaCreatedBy              uuid.UUID           `db:"meta_created_by"`
	MetaUpdatedAt              time.Time           `db:"meta_updated_at"`
	MetaUpdatedBy              nuuid.NUUID         `db:"meta_updated_by"`
	MetaDeletedAt              null.Time           `db:"meta_deleted_at"`
	MetaDeletedBy              nuuid.NUUID         `db:"meta_deleted_by"`
}
type PaymentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d Payments) ToPaymentsPrimaryID() PaymentsPrimaryID {
	return PaymentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentsList []*Payments
