package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentsDTOFieldNameType string

type paymentsDTOFieldName struct {
	Id                         PaymentsDTOFieldNameType
	PaymentCode                PaymentsDTOFieldNameType
	IntentId                   PaymentsDTOFieldNameType
	AttemptNumber              PaymentsDTOFieldNameType
	Psp                        PaymentsDTOFieldNameType
	PspTransactionId           PaymentsDTOFieldNameType
	PspReference               PaymentsDTOFieldNameType
	PspRawRequest              PaymentsDTOFieldNameType
	PspRawResponse             PaymentsDTOFieldNameType
	Amount                     PaymentsDTOFieldNameType
	Currency                   PaymentsDTOFieldNameType
	AmountInSettlementCurrency PaymentsDTOFieldNameType
	SettlementCurrency         PaymentsDTOFieldNameType
	FxRate                     PaymentsDTOFieldNameType
	FxRateSnapshotId           PaymentsDTOFieldNameType
	PaymentMethodId            PaymentsDTOFieldNameType
	PaymentMethodType          PaymentsDTOFieldNameType
	Status                     PaymentsDTOFieldNameType
	FailureCode                PaymentsDTOFieldNameType
	FailureMessage             PaymentsDTOFieldNameType
	FailureCategory            PaymentsDTOFieldNameType
	AuthorisedAt               PaymentsDTOFieldNameType
	AuthorisedAmount           PaymentsDTOFieldNameType
	CapturedAt                 PaymentsDTOFieldNameType
	CapturedAmount             PaymentsDTOFieldNameType
	ProcessingFee              PaymentsDTOFieldNameType
	ProcessingFeeCurrency      PaymentsDTOFieldNameType
	RiskScoreId                PaymentsDTOFieldNameType
	Description                PaymentsDTOFieldNameType
	Metadata                   PaymentsDTOFieldNameType
	CompletedAt                PaymentsDTOFieldNameType
	CancelledAt                PaymentsDTOFieldNameType
	ExpiredAt                  PaymentsDTOFieldNameType
	MetaCreatedAt              PaymentsDTOFieldNameType
	MetaCreatedBy              PaymentsDTOFieldNameType
	MetaUpdatedAt              PaymentsDTOFieldNameType
	MetaUpdatedBy              PaymentsDTOFieldNameType
	MetaDeletedAt              PaymentsDTOFieldNameType
	MetaDeletedBy              PaymentsDTOFieldNameType
}

var PaymentsDTOFieldName = paymentsDTOFieldName{
	Id:                         "id",
	PaymentCode:                "paymentCode",
	IntentId:                   "intentId",
	AttemptNumber:              "attemptNumber",
	Psp:                        "psp",
	PspTransactionId:           "pspTransactionId",
	PspReference:               "pspReference",
	PspRawRequest:              "pspRawRequest",
	PspRawResponse:             "pspRawResponse",
	Amount:                     "amount",
	Currency:                   "currency",
	AmountInSettlementCurrency: "amountInSettlementCurrency",
	SettlementCurrency:         "settlementCurrency",
	FxRate:                     "fxRate",
	FxRateSnapshotId:           "fxRateSnapshotId",
	PaymentMethodId:            "paymentMethodId",
	PaymentMethodType:          "paymentMethodType",
	Status:                     "status",
	FailureCode:                "failureCode",
	FailureMessage:             "failureMessage",
	FailureCategory:            "failureCategory",
	AuthorisedAt:               "authorisedAt",
	AuthorisedAmount:           "authorisedAmount",
	CapturedAt:                 "capturedAt",
	CapturedAmount:             "capturedAmount",
	ProcessingFee:              "processingFee",
	ProcessingFeeCurrency:      "processingFeeCurrency",
	RiskScoreId:                "riskScoreId",
	Description:                "description",
	Metadata:                   "metadata",
	CompletedAt:                "completedAt",
	CancelledAt:                "cancelledAt",
	ExpiredAt:                  "expiredAt",
	MetaCreatedAt:              "metaCreatedAt",
	MetaCreatedBy:              "metaCreatedBy",
	MetaUpdatedAt:              "metaUpdatedAt",
	MetaUpdatedBy:              "metaUpdatedBy",
	MetaDeletedAt:              "metaDeletedAt",
	MetaDeletedBy:              "metaDeletedBy",
}

func NewPaymentsListResponseFromFilterResult(result []model.PaymentsFilterResult, filter model.Filter) PaymentsSelectableListResponse {
	dtoPaymentsListResponse := PaymentsSelectableListResponse{}
	for _, payments := range result {
		dtoPaymentsResponse := NewPaymentsSelectableResponse(payments.Payments, filter)
		dtoPaymentsListResponse = append(dtoPaymentsListResponse, &dtoPaymentsResponse)
	}
	return dtoPaymentsListResponse
}

func transformPaymentsDTOFieldNameFromStr(field string) (dbField model.PaymentsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentsDTOFieldName.Id):
		return model.PaymentsDBFieldName.Id, true

	case string(PaymentsDTOFieldName.PaymentCode):
		return model.PaymentsDBFieldName.PaymentCode, true

	case string(PaymentsDTOFieldName.IntentId):
		return model.PaymentsDBFieldName.IntentId, true

	case string(PaymentsDTOFieldName.AttemptNumber):
		return model.PaymentsDBFieldName.AttemptNumber, true

	case string(PaymentsDTOFieldName.Psp):
		return model.PaymentsDBFieldName.Psp, true

	case string(PaymentsDTOFieldName.PspTransactionId):
		return model.PaymentsDBFieldName.PspTransactionId, true

	case string(PaymentsDTOFieldName.PspReference):
		return model.PaymentsDBFieldName.PspReference, true

	case string(PaymentsDTOFieldName.PspRawRequest):
		return model.PaymentsDBFieldName.PspRawRequest, true

	case string(PaymentsDTOFieldName.PspRawResponse):
		return model.PaymentsDBFieldName.PspRawResponse, true

	case string(PaymentsDTOFieldName.Amount):
		return model.PaymentsDBFieldName.Amount, true

	case string(PaymentsDTOFieldName.Currency):
		return model.PaymentsDBFieldName.Currency, true

	case string(PaymentsDTOFieldName.AmountInSettlementCurrency):
		return model.PaymentsDBFieldName.AmountInSettlementCurrency, true

	case string(PaymentsDTOFieldName.SettlementCurrency):
		return model.PaymentsDBFieldName.SettlementCurrency, true

	case string(PaymentsDTOFieldName.FxRate):
		return model.PaymentsDBFieldName.FxRate, true

	case string(PaymentsDTOFieldName.FxRateSnapshotId):
		return model.PaymentsDBFieldName.FxRateSnapshotId, true

	case string(PaymentsDTOFieldName.PaymentMethodId):
		return model.PaymentsDBFieldName.PaymentMethodId, true

	case string(PaymentsDTOFieldName.PaymentMethodType):
		return model.PaymentsDBFieldName.PaymentMethodType, true

	case string(PaymentsDTOFieldName.Status):
		return model.PaymentsDBFieldName.Status, true

	case string(PaymentsDTOFieldName.FailureCode):
		return model.PaymentsDBFieldName.FailureCode, true

	case string(PaymentsDTOFieldName.FailureMessage):
		return model.PaymentsDBFieldName.FailureMessage, true

	case string(PaymentsDTOFieldName.FailureCategory):
		return model.PaymentsDBFieldName.FailureCategory, true

	case string(PaymentsDTOFieldName.AuthorisedAt):
		return model.PaymentsDBFieldName.AuthorisedAt, true

	case string(PaymentsDTOFieldName.AuthorisedAmount):
		return model.PaymentsDBFieldName.AuthorisedAmount, true

	case string(PaymentsDTOFieldName.CapturedAt):
		return model.PaymentsDBFieldName.CapturedAt, true

	case string(PaymentsDTOFieldName.CapturedAmount):
		return model.PaymentsDBFieldName.CapturedAmount, true

	case string(PaymentsDTOFieldName.ProcessingFee):
		return model.PaymentsDBFieldName.ProcessingFee, true

	case string(PaymentsDTOFieldName.ProcessingFeeCurrency):
		return model.PaymentsDBFieldName.ProcessingFeeCurrency, true

	case string(PaymentsDTOFieldName.RiskScoreId):
		return model.PaymentsDBFieldName.RiskScoreId, true

	case string(PaymentsDTOFieldName.Description):
		return model.PaymentsDBFieldName.Description, true

	case string(PaymentsDTOFieldName.Metadata):
		return model.PaymentsDBFieldName.Metadata, true

	case string(PaymentsDTOFieldName.CompletedAt):
		return model.PaymentsDBFieldName.CompletedAt, true

	case string(PaymentsDTOFieldName.CancelledAt):
		return model.PaymentsDBFieldName.CancelledAt, true

	case string(PaymentsDTOFieldName.ExpiredAt):
		return model.PaymentsDBFieldName.ExpiredAt, true

	case string(PaymentsDTOFieldName.MetaCreatedAt):
		return model.PaymentsDBFieldName.MetaCreatedAt, true

	case string(PaymentsDTOFieldName.MetaCreatedBy):
		return model.PaymentsDBFieldName.MetaCreatedBy, true

	case string(PaymentsDTOFieldName.MetaUpdatedAt):
		return model.PaymentsDBFieldName.MetaUpdatedAt, true

	case string(PaymentsDTOFieldName.MetaUpdatedBy):
		return model.PaymentsDBFieldName.MetaUpdatedBy, true

	case string(PaymentsDTOFieldName.MetaDeletedAt):
		return model.PaymentsDBFieldName.MetaDeletedAt, true

	case string(PaymentsDTOFieldName.MetaDeletedBy):
		return model.PaymentsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformPaymentsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformPaymentsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultPaymentsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(PaymentsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentsSelectableResponse map[string]interface{}
type PaymentsSelectableListResponse []*PaymentsSelectableResponse

func NewPaymentsSelectableResponse(payments model.Payments, filter model.Filter) PaymentsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentsDBFieldName.Id),
			string(model.PaymentsDBFieldName.PaymentCode),
			string(model.PaymentsDBFieldName.IntentId),
			string(model.PaymentsDBFieldName.AttemptNumber),
			string(model.PaymentsDBFieldName.Psp),
			string(model.PaymentsDBFieldName.PspTransactionId),
			string(model.PaymentsDBFieldName.PspReference),
			string(model.PaymentsDBFieldName.PspRawRequest),
			string(model.PaymentsDBFieldName.PspRawResponse),
			string(model.PaymentsDBFieldName.Amount),
			string(model.PaymentsDBFieldName.Currency),
			string(model.PaymentsDBFieldName.AmountInSettlementCurrency),
			string(model.PaymentsDBFieldName.SettlementCurrency),
			string(model.PaymentsDBFieldName.FxRate),
			string(model.PaymentsDBFieldName.FxRateSnapshotId),
			string(model.PaymentsDBFieldName.PaymentMethodId),
			string(model.PaymentsDBFieldName.PaymentMethodType),
			string(model.PaymentsDBFieldName.Status),
			string(model.PaymentsDBFieldName.FailureCode),
			string(model.PaymentsDBFieldName.FailureMessage),
			string(model.PaymentsDBFieldName.FailureCategory),
			string(model.PaymentsDBFieldName.AuthorisedAt),
			string(model.PaymentsDBFieldName.AuthorisedAmount),
			string(model.PaymentsDBFieldName.CapturedAt),
			string(model.PaymentsDBFieldName.CapturedAmount),
			string(model.PaymentsDBFieldName.ProcessingFee),
			string(model.PaymentsDBFieldName.ProcessingFeeCurrency),
			string(model.PaymentsDBFieldName.RiskScoreId),
			string(model.PaymentsDBFieldName.Description),
			string(model.PaymentsDBFieldName.Metadata),
			string(model.PaymentsDBFieldName.CompletedAt),
			string(model.PaymentsDBFieldName.CancelledAt),
			string(model.PaymentsDBFieldName.ExpiredAt),
			string(model.PaymentsDBFieldName.MetaCreatedAt),
			string(model.PaymentsDBFieldName.MetaCreatedBy),
			string(model.PaymentsDBFieldName.MetaUpdatedAt),
			string(model.PaymentsDBFieldName.MetaUpdatedBy),
			string(model.PaymentsDBFieldName.MetaDeletedAt),
			string(model.PaymentsDBFieldName.MetaDeletedBy),
		)
	}
	paymentsSelectableResponse := PaymentsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.PaymentsDBFieldName.Id):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.Id)] = payments.Id

		case string(model.PaymentsDBFieldName.PaymentCode):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.PaymentCode)] = payments.PaymentCode

		case string(model.PaymentsDBFieldName.IntentId):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.IntentId)] = payments.IntentId

		case string(model.PaymentsDBFieldName.AttemptNumber):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.AttemptNumber)] = payments.AttemptNumber

		case string(model.PaymentsDBFieldName.Psp):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.Psp)] = payments.Psp

		case string(model.PaymentsDBFieldName.PspTransactionId):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.PspTransactionId)] = payments.PspTransactionId

		case string(model.PaymentsDBFieldName.PspReference):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.PspReference)] = payments.PspReference

		case string(model.PaymentsDBFieldName.PspRawRequest):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.PspRawRequest)] = payments.PspRawRequest

		case string(model.PaymentsDBFieldName.PspRawResponse):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.PspRawResponse)] = payments.PspRawResponse

		case string(model.PaymentsDBFieldName.Amount):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.Amount)] = payments.Amount

		case string(model.PaymentsDBFieldName.Currency):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.Currency)] = payments.Currency

		case string(model.PaymentsDBFieldName.AmountInSettlementCurrency):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.AmountInSettlementCurrency)] = payments.AmountInSettlementCurrency

		case string(model.PaymentsDBFieldName.SettlementCurrency):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.SettlementCurrency)] = payments.SettlementCurrency

		case string(model.PaymentsDBFieldName.FxRate):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.FxRate)] = payments.FxRate

		case string(model.PaymentsDBFieldName.FxRateSnapshotId):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.FxRateSnapshotId)] = payments.FxRateSnapshotId

		case string(model.PaymentsDBFieldName.PaymentMethodId):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.PaymentMethodId)] = payments.PaymentMethodId

		case string(model.PaymentsDBFieldName.PaymentMethodType):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.PaymentMethodType)] = payments.PaymentMethodType

		case string(model.PaymentsDBFieldName.Status):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.Status)] = payments.Status

		case string(model.PaymentsDBFieldName.FailureCode):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.FailureCode)] = payments.FailureCode

		case string(model.PaymentsDBFieldName.FailureMessage):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.FailureMessage)] = payments.FailureMessage

		case string(model.PaymentsDBFieldName.FailureCategory):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.FailureCategory)] = payments.FailureCategory

		case string(model.PaymentsDBFieldName.AuthorisedAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.AuthorisedAt)] = payments.AuthorisedAt

		case string(model.PaymentsDBFieldName.AuthorisedAmount):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.AuthorisedAmount)] = payments.AuthorisedAmount

		case string(model.PaymentsDBFieldName.CapturedAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.CapturedAt)] = payments.CapturedAt

		case string(model.PaymentsDBFieldName.CapturedAmount):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.CapturedAmount)] = payments.CapturedAmount

		case string(model.PaymentsDBFieldName.ProcessingFee):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.ProcessingFee)] = payments.ProcessingFee

		case string(model.PaymentsDBFieldName.ProcessingFeeCurrency):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.ProcessingFeeCurrency)] = payments.ProcessingFeeCurrency

		case string(model.PaymentsDBFieldName.RiskScoreId):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.RiskScoreId)] = payments.RiskScoreId

		case string(model.PaymentsDBFieldName.Description):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.Description)] = payments.Description

		case string(model.PaymentsDBFieldName.Metadata):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.Metadata)] = payments.Metadata

		case string(model.PaymentsDBFieldName.CompletedAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.CompletedAt)] = payments.CompletedAt

		case string(model.PaymentsDBFieldName.CancelledAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.CancelledAt)] = payments.CancelledAt

		case string(model.PaymentsDBFieldName.ExpiredAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.ExpiredAt)] = payments.ExpiredAt

		case string(model.PaymentsDBFieldName.MetaCreatedAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.MetaCreatedAt)] = payments.MetaCreatedAt

		case string(model.PaymentsDBFieldName.MetaCreatedBy):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.MetaCreatedBy)] = payments.MetaCreatedBy

		case string(model.PaymentsDBFieldName.MetaUpdatedAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.MetaUpdatedAt)] = payments.MetaUpdatedAt

		case string(model.PaymentsDBFieldName.MetaUpdatedBy):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.MetaUpdatedBy)] = payments.MetaUpdatedBy

		case string(model.PaymentsDBFieldName.MetaDeletedAt):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.MetaDeletedAt)] = payments.MetaDeletedAt

		case string(model.PaymentsDBFieldName.MetaDeletedBy):
			paymentsSelectableResponse[string(PaymentsDTOFieldName.MetaDeletedBy)] = payments.MetaDeletedBy

		}
	}
	return paymentsSelectableResponse
}

type PaymentsFilterResponse struct {
	Metadata Metadata                       `json:"metadata"`
	Data     PaymentsSelectableListResponse `json:"data"`
}

func NewPaymentsFilterResponse(result []model.PaymentsFilterResult, filter model.Filter) (resp PaymentsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewPaymentsListResponseFromFilterResult(result, filter)
	return resp
}

type PaymentsCreateRequest struct {
	PaymentCode                string                  `json:"paymentCode"`
	IntentId                   uuid.UUID               `json:"intentId"`
	AttemptNumber              int16                   `json:"attemptNumber"`
	Psp                        model.Psp               `json:"psp" example:"MIDTRANS" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	PspTransactionId           string                  `json:"pspTransactionId"`
	PspReference               string                  `json:"pspReference"`
	PspRawRequest              json.RawMessage         `json:"pspRawRequest"`
	PspRawResponse             json.RawMessage         `json:"pspRawResponse"`
	Amount                     decimal.Decimal         `json:"amount"`
	Currency                   model.PaymentCurrency   `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	AmountInSettlementCurrency decimal.Decimal         `json:"amountInSettlementCurrency"`
	SettlementCurrency         model.PaymentCurrency   `json:"settlementCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	FxRate                     decimal.Decimal         `json:"fxRate"`
	FxRateSnapshotId           uuid.UUID               `json:"fxRateSnapshotId"`
	PaymentMethodId            uuid.UUID               `json:"paymentMethodId"`
	PaymentMethodType          model.PaymentMethodType `json:"paymentMethodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status                     model.PaymentStatus     `json:"status" example:"INITIATED" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	FailureCode                string                  `json:"failureCode"`
	FailureMessage             string                  `json:"failureMessage"`
	FailureCategory            string                  `json:"failureCategory"`
	AuthorisedAt               time.Time               `json:"authorisedAt"`
	AuthorisedAmount           decimal.Decimal         `json:"authorisedAmount"`
	CapturedAt                 time.Time               `json:"capturedAt"`
	CapturedAmount             decimal.Decimal         `json:"capturedAmount"`
	ProcessingFee              decimal.Decimal         `json:"processingFee"`
	ProcessingFeeCurrency      model.PaymentCurrency   `json:"processingFeeCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	RiskScoreId                uuid.UUID               `json:"riskScoreId"`
	Description                string                  `json:"description"`
	Metadata                   json.RawMessage         `json:"metadata"`
	CompletedAt                time.Time               `json:"completedAt"`
	CancelledAt                time.Time               `json:"cancelledAt"`
	ExpiredAt                  time.Time               `json:"expiredAt"`
	MetaCreatedAt              time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy              uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt              time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy              uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt              time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy              uuid.UUID               `json:"metaDeletedBy"`
}

func (d *PaymentsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentsCreateRequest) ToModel() model.Payments {
	id, _ := uuid.NewV7()
	return model.Payments{
		Id:                         id,
		PaymentCode:                d.PaymentCode,
		IntentId:                   d.IntentId,
		AttemptNumber:              d.AttemptNumber,
		Psp:                        d.Psp,
		PspTransactionId:           null.StringFrom(d.PspTransactionId),
		PspReference:               null.StringFrom(d.PspReference),
		PspRawRequest:              d.PspRawRequest,
		PspRawResponse:             d.PspRawResponse,
		Amount:                     d.Amount,
		Currency:                   d.Currency,
		AmountInSettlementCurrency: decimal.NewNullDecimal(d.AmountInSettlementCurrency),
		SettlementCurrency:         d.SettlementCurrency,
		FxRate:                     decimal.NewNullDecimal(d.FxRate),
		FxRateSnapshotId:           nuuid.From(d.FxRateSnapshotId),
		PaymentMethodId:            nuuid.From(d.PaymentMethodId),
		PaymentMethodType:          d.PaymentMethodType,
		Status:                     d.Status,
		FailureCode:                null.StringFrom(d.FailureCode),
		FailureMessage:             null.StringFrom(d.FailureMessage),
		FailureCategory:            null.StringFrom(d.FailureCategory),
		AuthorisedAt:               null.TimeFrom(d.AuthorisedAt),
		AuthorisedAmount:           decimal.NewNullDecimal(d.AuthorisedAmount),
		CapturedAt:                 null.TimeFrom(d.CapturedAt),
		CapturedAmount:             decimal.NewNullDecimal(d.CapturedAmount),
		ProcessingFee:              decimal.NewNullDecimal(d.ProcessingFee),
		ProcessingFeeCurrency:      d.ProcessingFeeCurrency,
		RiskScoreId:                nuuid.From(d.RiskScoreId),
		Description:                null.StringFrom(d.Description),
		Metadata:                   d.Metadata,
		CompletedAt:                null.TimeFrom(d.CompletedAt),
		CancelledAt:                null.TimeFrom(d.CancelledAt),
		ExpiredAt:                  null.TimeFrom(d.ExpiredAt),
		MetaCreatedAt:              d.MetaCreatedAt,
		MetaCreatedBy:              d.MetaCreatedBy,
		MetaUpdatedAt:              d.MetaUpdatedAt,
		MetaUpdatedBy:              nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:              null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:              nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentsListCreateRequest []*PaymentsCreateRequest

func (d PaymentsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payments := range d {
		err = validator.Struct(payments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentsListCreateRequest) ToModelList() []model.Payments {
	out := make([]model.Payments, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentsUpdateRequest struct {
	PaymentCode                string                  `json:"paymentCode"`
	IntentId                   uuid.UUID               `json:"intentId"`
	AttemptNumber              int16                   `json:"attemptNumber"`
	Psp                        model.Psp               `json:"psp" example:"MIDTRANS" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	PspTransactionId           string                  `json:"pspTransactionId"`
	PspReference               string                  `json:"pspReference"`
	PspRawRequest              json.RawMessage         `json:"pspRawRequest"`
	PspRawResponse             json.RawMessage         `json:"pspRawResponse"`
	Amount                     decimal.Decimal         `json:"amount"`
	Currency                   model.PaymentCurrency   `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	AmountInSettlementCurrency decimal.Decimal         `json:"amountInSettlementCurrency"`
	SettlementCurrency         model.PaymentCurrency   `json:"settlementCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	FxRate                     decimal.Decimal         `json:"fxRate"`
	FxRateSnapshotId           uuid.UUID               `json:"fxRateSnapshotId"`
	PaymentMethodId            uuid.UUID               `json:"paymentMethodId"`
	PaymentMethodType          model.PaymentMethodType `json:"paymentMethodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status                     model.PaymentStatus     `json:"status" example:"INITIATED" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	FailureCode                string                  `json:"failureCode"`
	FailureMessage             string                  `json:"failureMessage"`
	FailureCategory            string                  `json:"failureCategory"`
	AuthorisedAt               time.Time               `json:"authorisedAt"`
	AuthorisedAmount           decimal.Decimal         `json:"authorisedAmount"`
	CapturedAt                 time.Time               `json:"capturedAt"`
	CapturedAmount             decimal.Decimal         `json:"capturedAmount"`
	ProcessingFee              decimal.Decimal         `json:"processingFee"`
	ProcessingFeeCurrency      model.PaymentCurrency   `json:"processingFeeCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	RiskScoreId                uuid.UUID               `json:"riskScoreId"`
	Description                string                  `json:"description"`
	Metadata                   json.RawMessage         `json:"metadata"`
	CompletedAt                time.Time               `json:"completedAt"`
	CancelledAt                time.Time               `json:"cancelledAt"`
	ExpiredAt                  time.Time               `json:"expiredAt"`
	MetaCreatedAt              time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy              uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt              time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy              uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt              time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy              uuid.UUID               `json:"metaDeletedBy"`
}

func (d *PaymentsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentsUpdateRequest) ToModel() model.Payments {
	return model.Payments{
		PaymentCode:                d.PaymentCode,
		IntentId:                   d.IntentId,
		AttemptNumber:              d.AttemptNumber,
		Psp:                        d.Psp,
		PspTransactionId:           null.StringFrom(d.PspTransactionId),
		PspReference:               null.StringFrom(d.PspReference),
		PspRawRequest:              d.PspRawRequest,
		PspRawResponse:             d.PspRawResponse,
		Amount:                     d.Amount,
		Currency:                   d.Currency,
		AmountInSettlementCurrency: decimal.NewNullDecimal(d.AmountInSettlementCurrency),
		SettlementCurrency:         d.SettlementCurrency,
		FxRate:                     decimal.NewNullDecimal(d.FxRate),
		FxRateSnapshotId:           nuuid.From(d.FxRateSnapshotId),
		PaymentMethodId:            nuuid.From(d.PaymentMethodId),
		PaymentMethodType:          d.PaymentMethodType,
		Status:                     d.Status,
		FailureCode:                null.StringFrom(d.FailureCode),
		FailureMessage:             null.StringFrom(d.FailureMessage),
		FailureCategory:            null.StringFrom(d.FailureCategory),
		AuthorisedAt:               null.TimeFrom(d.AuthorisedAt),
		AuthorisedAmount:           decimal.NewNullDecimal(d.AuthorisedAmount),
		CapturedAt:                 null.TimeFrom(d.CapturedAt),
		CapturedAmount:             decimal.NewNullDecimal(d.CapturedAmount),
		ProcessingFee:              decimal.NewNullDecimal(d.ProcessingFee),
		ProcessingFeeCurrency:      d.ProcessingFeeCurrency,
		RiskScoreId:                nuuid.From(d.RiskScoreId),
		Description:                null.StringFrom(d.Description),
		Metadata:                   d.Metadata,
		CompletedAt:                null.TimeFrom(d.CompletedAt),
		CancelledAt:                null.TimeFrom(d.CancelledAt),
		ExpiredAt:                  null.TimeFrom(d.ExpiredAt),
		MetaCreatedAt:              d.MetaCreatedAt,
		MetaCreatedBy:              d.MetaCreatedBy,
		MetaUpdatedAt:              d.MetaUpdatedAt,
		MetaUpdatedBy:              nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:              null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:              nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentsBulkUpdateRequest struct {
	Id                         uuid.UUID               `json:"id"`
	PaymentCode                string                  `json:"paymentCode"`
	IntentId                   uuid.UUID               `json:"intentId"`
	AttemptNumber              int16                   `json:"attemptNumber"`
	Psp                        model.Psp               `json:"psp" example:"MIDTRANS" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	PspTransactionId           string                  `json:"pspTransactionId"`
	PspReference               string                  `json:"pspReference"`
	PspRawRequest              json.RawMessage         `json:"pspRawRequest"`
	PspRawResponse             json.RawMessage         `json:"pspRawResponse"`
	Amount                     decimal.Decimal         `json:"amount"`
	Currency                   model.PaymentCurrency   `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	AmountInSettlementCurrency decimal.Decimal         `json:"amountInSettlementCurrency"`
	SettlementCurrency         model.PaymentCurrency   `json:"settlementCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	FxRate                     decimal.Decimal         `json:"fxRate"`
	FxRateSnapshotId           uuid.UUID               `json:"fxRateSnapshotId"`
	PaymentMethodId            uuid.UUID               `json:"paymentMethodId"`
	PaymentMethodType          model.PaymentMethodType `json:"paymentMethodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status                     model.PaymentStatus     `json:"status" example:"INITIATED" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	FailureCode                string                  `json:"failureCode"`
	FailureMessage             string                  `json:"failureMessage"`
	FailureCategory            string                  `json:"failureCategory"`
	AuthorisedAt               time.Time               `json:"authorisedAt"`
	AuthorisedAmount           decimal.Decimal         `json:"authorisedAmount"`
	CapturedAt                 time.Time               `json:"capturedAt"`
	CapturedAmount             decimal.Decimal         `json:"capturedAmount"`
	ProcessingFee              decimal.Decimal         `json:"processingFee"`
	ProcessingFeeCurrency      model.PaymentCurrency   `json:"processingFeeCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	RiskScoreId                uuid.UUID               `json:"riskScoreId"`
	Description                string                  `json:"description"`
	Metadata                   json.RawMessage         `json:"metadata"`
	CompletedAt                time.Time               `json:"completedAt"`
	CancelledAt                time.Time               `json:"cancelledAt"`
	ExpiredAt                  time.Time               `json:"expiredAt"`
	MetaCreatedAt              time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy              uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt              time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy              uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt              time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy              uuid.UUID               `json:"metaDeletedBy"`
}

func (d PaymentsBulkUpdateRequest) PrimaryID() PaymentsPrimaryID {
	return PaymentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentsListBulkUpdateRequest []*PaymentsBulkUpdateRequest

func (d PaymentsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payments := range d {
		err = validator.Struct(payments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentsBulkUpdateRequest) ToModel() model.Payments {
	return model.Payments{
		Id:                         d.Id,
		PaymentCode:                d.PaymentCode,
		IntentId:                   d.IntentId,
		AttemptNumber:              d.AttemptNumber,
		Psp:                        d.Psp,
		PspTransactionId:           null.StringFrom(d.PspTransactionId),
		PspReference:               null.StringFrom(d.PspReference),
		PspRawRequest:              d.PspRawRequest,
		PspRawResponse:             d.PspRawResponse,
		Amount:                     d.Amount,
		Currency:                   d.Currency,
		AmountInSettlementCurrency: decimal.NewNullDecimal(d.AmountInSettlementCurrency),
		SettlementCurrency:         d.SettlementCurrency,
		FxRate:                     decimal.NewNullDecimal(d.FxRate),
		FxRateSnapshotId:           nuuid.From(d.FxRateSnapshotId),
		PaymentMethodId:            nuuid.From(d.PaymentMethodId),
		PaymentMethodType:          d.PaymentMethodType,
		Status:                     d.Status,
		FailureCode:                null.StringFrom(d.FailureCode),
		FailureMessage:             null.StringFrom(d.FailureMessage),
		FailureCategory:            null.StringFrom(d.FailureCategory),
		AuthorisedAt:               null.TimeFrom(d.AuthorisedAt),
		AuthorisedAmount:           decimal.NewNullDecimal(d.AuthorisedAmount),
		CapturedAt:                 null.TimeFrom(d.CapturedAt),
		CapturedAmount:             decimal.NewNullDecimal(d.CapturedAmount),
		ProcessingFee:              decimal.NewNullDecimal(d.ProcessingFee),
		ProcessingFeeCurrency:      d.ProcessingFeeCurrency,
		RiskScoreId:                nuuid.From(d.RiskScoreId),
		Description:                null.StringFrom(d.Description),
		Metadata:                   d.Metadata,
		CompletedAt:                null.TimeFrom(d.CompletedAt),
		CancelledAt:                null.TimeFrom(d.CancelledAt),
		ExpiredAt:                  null.TimeFrom(d.ExpiredAt),
		MetaCreatedAt:              d.MetaCreatedAt,
		MetaCreatedBy:              d.MetaCreatedBy,
		MetaUpdatedAt:              d.MetaUpdatedAt,
		MetaUpdatedBy:              nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:              null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:              nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentsResponse struct {
	Id                         uuid.UUID               `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentCode                string                  `json:"paymentCode" validate:"required"`
	IntentId                   uuid.UUID               `json:"intentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	AttemptNumber              int16                   `json:"attemptNumber"`
	Psp                        model.Psp               `json:"psp" validate:"required,oneof=MIDTRANS XENDIT STRIPE DOKU DANA OVO GOPAY SHOPEE_PAY LINK_AJA FLIP INTERNAL" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	PspTransactionId           string                  `json:"pspTransactionId"`
	PspReference               string                  `json:"pspReference"`
	PspRawRequest              json.RawMessage         `json:"pspRawRequest" validate:"required" swaggertype:"object"`
	PspRawResponse             json.RawMessage         `json:"pspRawResponse" validate:"required" swaggertype:"object"`
	Amount                     decimal.Decimal         `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency                   model.PaymentCurrency   `json:"currency" validate:"required,oneof=IDR USD SGD MYR PHP THB AED EUR GBP JPY" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	AmountInSettlementCurrency decimal.Decimal         `json:"amountInSettlementCurrency" format:"decimal" example:"100.50"`
	SettlementCurrency         model.PaymentCurrency   `json:"settlementCurrency" validate:"required,oneof=IDR USD SGD MYR PHP THB AED EUR GBP JPY" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	FxRate                     decimal.Decimal         `json:"fxRate" format:"decimal" example:"100.50"`
	FxRateSnapshotId           uuid.UUID               `json:"fxRateSnapshotId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentMethodId            uuid.UUID               `json:"paymentMethodId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentMethodType          model.PaymentMethodType `json:"paymentMethodType" validate:"required,oneof=CARD VIRTUAL_ACCOUNT QRIS EWALLET DIRECT_DEBIT BANK_TRANSFER PAYLATER VOUCHER POINTS CASH_ON_DELIVERY" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status                     model.PaymentStatus     `json:"status" validate:"oneof=INITIATED PENDING AUTHORISED CAPTURED PARTIALLY_CAPTURED COMPLETED FAILED CANCELLED EXPIRED REFUNDING REFUNDED PARTIALLY_REFUNDED DISPUTED CHARGEBACK_WON CHARGEBACK_LOST" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	FailureCode                string                  `json:"failureCode"`
	FailureMessage             string                  `json:"failureMessage"`
	FailureCategory            string                  `json:"failureCategory"`
	AuthorisedAt               time.Time               `json:"authorisedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	AuthorisedAmount           decimal.Decimal         `json:"authorisedAmount" format:"decimal" example:"100.50"`
	CapturedAt                 time.Time               `json:"capturedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CapturedAmount             decimal.Decimal         `json:"capturedAmount" format:"decimal" example:"100.50"`
	ProcessingFee              decimal.Decimal         `json:"processingFee" format:"decimal" example:"100.50"`
	ProcessingFeeCurrency      model.PaymentCurrency   `json:"processingFeeCurrency" validate:"required,oneof=IDR USD SGD MYR PHP THB AED EUR GBP JPY" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	RiskScoreId                uuid.UUID               `json:"riskScoreId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Description                string                  `json:"description"`
	Metadata                   json.RawMessage         `json:"metadata" swaggertype:"object"`
	CompletedAt                time.Time               `json:"completedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CancelledAt                time.Time               `json:"cancelledAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ExpiredAt                  time.Time               `json:"expiredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedAt              time.Time               `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy              uuid.UUID               `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt              time.Time               `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy              uuid.UUID               `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt              time.Time               `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy              uuid.UUID               `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewPaymentsResponse(payments model.Payments) PaymentsResponse {
	return PaymentsResponse{
		Id:                         payments.Id,
		PaymentCode:                payments.PaymentCode,
		IntentId:                   payments.IntentId,
		AttemptNumber:              payments.AttemptNumber,
		Psp:                        model.Psp(payments.Psp),
		PspTransactionId:           payments.PspTransactionId.String,
		PspReference:               payments.PspReference.String,
		PspRawRequest:              payments.PspRawRequest,
		PspRawResponse:             payments.PspRawResponse,
		Amount:                     payments.Amount,
		Currency:                   model.PaymentCurrency(payments.Currency),
		AmountInSettlementCurrency: payments.AmountInSettlementCurrency.Decimal,
		SettlementCurrency:         model.PaymentCurrency(payments.SettlementCurrency),
		FxRate:                     payments.FxRate.Decimal,
		FxRateSnapshotId:           payments.FxRateSnapshotId.UUID,
		PaymentMethodId:            payments.PaymentMethodId.UUID,
		PaymentMethodType:          model.PaymentMethodType(payments.PaymentMethodType),
		Status:                     model.PaymentStatus(payments.Status),
		FailureCode:                payments.FailureCode.String,
		FailureMessage:             payments.FailureMessage.String,
		FailureCategory:            payments.FailureCategory.String,
		AuthorisedAt:               payments.AuthorisedAt.Time,
		AuthorisedAmount:           payments.AuthorisedAmount.Decimal,
		CapturedAt:                 payments.CapturedAt.Time,
		CapturedAmount:             payments.CapturedAmount.Decimal,
		ProcessingFee:              payments.ProcessingFee.Decimal,
		ProcessingFeeCurrency:      model.PaymentCurrency(payments.ProcessingFeeCurrency),
		RiskScoreId:                payments.RiskScoreId.UUID,
		Description:                payments.Description.String,
		Metadata:                   payments.Metadata,
		CompletedAt:                payments.CompletedAt.Time,
		CancelledAt:                payments.CancelledAt.Time,
		ExpiredAt:                  payments.ExpiredAt.Time,
		MetaCreatedAt:              payments.MetaCreatedAt,
		MetaCreatedBy:              payments.MetaCreatedBy,
		MetaUpdatedAt:              payments.MetaUpdatedAt,
		MetaUpdatedBy:              payments.MetaUpdatedBy.UUID,
		MetaDeletedAt:              payments.MetaDeletedAt.Time,
		MetaDeletedBy:              payments.MetaDeletedBy.UUID,
	}
}

type PaymentsListResponse []*PaymentsResponse

func NewPaymentsListResponse(paymentsList model.PaymentsList) PaymentsListResponse {
	dtoPaymentsListResponse := PaymentsListResponse{}
	for _, payments := range paymentsList {
		dtoPaymentsResponse := NewPaymentsResponse(*payments)
		dtoPaymentsListResponse = append(dtoPaymentsListResponse, &dtoPaymentsResponse)
	}
	return dtoPaymentsListResponse
}

type PaymentsPrimaryIDList []PaymentsPrimaryID

func (d PaymentsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payments := range d {
		err = validator.Struct(payments)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentsPrimaryID) ToModel() model.PaymentsPrimaryID {
	return model.PaymentsPrimaryID{
		Id: d.Id,
	}
}
