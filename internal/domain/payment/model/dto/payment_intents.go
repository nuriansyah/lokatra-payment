package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"net/netip"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentIntentsDTOFieldNameType string

type paymentIntentsDTOFieldName struct {
	Id                  PaymentIntentsDTOFieldNameType
	IntentCode          PaymentIntentsDTOFieldNameType
	MerchantId          PaymentIntentsDTOFieldNameType
	OrderId             PaymentIntentsDTOFieldNameType
	OrderType           PaymentIntentsDTOFieldNameType
	Amount              PaymentIntentsDTOFieldNameType
	Currency            PaymentIntentsDTOFieldNameType
	TaxAmount           PaymentIntentsDTOFieldNameType
	DiscountAmount      PaymentIntentsDTOFieldNameType
	TipAmount           PaymentIntentsDTOFieldNameType
	UserId              PaymentIntentsDTOFieldNameType
	CustomerName        PaymentIntentsDTOFieldNameType
	CustomerEmail       PaymentIntentsDTOFieldNameType
	CustomerPhone       PaymentIntentsDTOFieldNameType
	CustomerIp          PaymentIntentsDTOFieldNameType
	CustomerCountry     PaymentIntentsDTOFieldNameType
	PaymentMethodId     PaymentIntentsDTOFieldNameType
	PaymentMethodType   PaymentIntentsDTOFieldNameType
	Status              PaymentIntentsDTOFieldNameType
	RoutingProfileId    PaymentIntentsDTOFieldNameType
	ExpiresAt           PaymentIntentsDTOFieldNameType
	Requires3ds         PaymentIntentsDTOFieldNameType
	ThreeDsVersion      PaymentIntentsDTOFieldNameType
	Description         PaymentIntentsDTOFieldNameType
	StatementDescriptor PaymentIntentsDTOFieldNameType
	Metadata            PaymentIntentsDTOFieldNameType
	PromoCode           PaymentIntentsDTOFieldNameType
	PromoDiscountAmount PaymentIntentsDTOFieldNameType
	IdempotencyKeyId    PaymentIntentsDTOFieldNameType
	ConfirmedAt         PaymentIntentsDTOFieldNameType
	CancelledAt         PaymentIntentsDTOFieldNameType
	CancellationReason  PaymentIntentsDTOFieldNameType
	MetaCreatedAt       PaymentIntentsDTOFieldNameType
	MetaCreatedBy       PaymentIntentsDTOFieldNameType
	MetaUpdatedAt       PaymentIntentsDTOFieldNameType
	MetaUpdatedBy       PaymentIntentsDTOFieldNameType
	MetaDeletedAt       PaymentIntentsDTOFieldNameType
	MetaDeletedBy       PaymentIntentsDTOFieldNameType
}

var PaymentIntentsDTOFieldName = paymentIntentsDTOFieldName{
	Id:                  "id",
	IntentCode:          "intentCode",
	MerchantId:          "merchantId",
	OrderId:             "orderId",
	OrderType:           "orderType",
	Amount:              "amount",
	Currency:            "currency",
	TaxAmount:           "taxAmount",
	DiscountAmount:      "discountAmount",
	TipAmount:           "tipAmount",
	UserId:              "userId",
	CustomerName:        "customerName",
	CustomerEmail:       "customerEmail",
	CustomerPhone:       "customerPhone",
	CustomerIp:          "customerIp",
	CustomerCountry:     "customerCountry",
	PaymentMethodId:     "paymentMethodId",
	PaymentMethodType:   "paymentMethodType",
	Status:              "status",
	RoutingProfileId:    "routingProfileId",
	ExpiresAt:           "expiresAt",
	Requires3ds:         "requires3ds",
	ThreeDsVersion:      "threeDsVersion",
	Description:         "description",
	StatementDescriptor: "statementDescriptor",
	Metadata:            "metadata",
	PromoCode:           "promoCode",
	PromoDiscountAmount: "promoDiscountAmount",
	IdempotencyKeyId:    "idempotencyKeyId",
	ConfirmedAt:         "confirmedAt",
	CancelledAt:         "cancelledAt",
	CancellationReason:  "cancellationReason",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func NewPaymentIntentsListResponseFromFilterResult(result []model.PaymentIntentsFilterResult, filter model.Filter) PaymentIntentsSelectableListResponse {
	dtoPaymentIntentsListResponse := PaymentIntentsSelectableListResponse{}
	for _, paymentIntents := range result {
		dtoPaymentIntentsResponse := NewPaymentIntentsSelectableResponse(paymentIntents.PaymentIntents, filter)
		dtoPaymentIntentsListResponse = append(dtoPaymentIntentsListResponse, &dtoPaymentIntentsResponse)
	}
	return dtoPaymentIntentsListResponse
}

func transformPaymentIntentsDTOFieldNameFromStr(field string) (dbField model.PaymentIntentsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentIntentsDTOFieldName.Id):
		return model.PaymentIntentsDBFieldName.Id, true

	case string(PaymentIntentsDTOFieldName.IntentCode):
		return model.PaymentIntentsDBFieldName.IntentCode, true

	case string(PaymentIntentsDTOFieldName.MerchantId):
		return model.PaymentIntentsDBFieldName.MerchantId, true

	case string(PaymentIntentsDTOFieldName.OrderId):
		return model.PaymentIntentsDBFieldName.OrderId, true

	case string(PaymentIntentsDTOFieldName.OrderType):
		return model.PaymentIntentsDBFieldName.OrderType, true

	case string(PaymentIntentsDTOFieldName.Amount):
		return model.PaymentIntentsDBFieldName.Amount, true

	case string(PaymentIntentsDTOFieldName.Currency):
		return model.PaymentIntentsDBFieldName.Currency, true

	case string(PaymentIntentsDTOFieldName.TaxAmount):
		return model.PaymentIntentsDBFieldName.TaxAmount, true

	case string(PaymentIntentsDTOFieldName.DiscountAmount):
		return model.PaymentIntentsDBFieldName.DiscountAmount, true

	case string(PaymentIntentsDTOFieldName.TipAmount):
		return model.PaymentIntentsDBFieldName.TipAmount, true

	case string(PaymentIntentsDTOFieldName.UserId):
		return model.PaymentIntentsDBFieldName.UserId, true

	case string(PaymentIntentsDTOFieldName.CustomerName):
		return model.PaymentIntentsDBFieldName.CustomerName, true

	case string(PaymentIntentsDTOFieldName.CustomerEmail):
		return model.PaymentIntentsDBFieldName.CustomerEmail, true

	case string(PaymentIntentsDTOFieldName.CustomerPhone):
		return model.PaymentIntentsDBFieldName.CustomerPhone, true

	case string(PaymentIntentsDTOFieldName.CustomerIp):
		return model.PaymentIntentsDBFieldName.CustomerIp, true

	case string(PaymentIntentsDTOFieldName.CustomerCountry):
		return model.PaymentIntentsDBFieldName.CustomerCountry, true

	case string(PaymentIntentsDTOFieldName.PaymentMethodId):
		return model.PaymentIntentsDBFieldName.PaymentMethodId, true

	case string(PaymentIntentsDTOFieldName.PaymentMethodType):
		return model.PaymentIntentsDBFieldName.PaymentMethodType, true

	case string(PaymentIntentsDTOFieldName.Status):
		return model.PaymentIntentsDBFieldName.Status, true

	case string(PaymentIntentsDTOFieldName.RoutingProfileId):
		return model.PaymentIntentsDBFieldName.RoutingProfileId, true

	case string(PaymentIntentsDTOFieldName.ExpiresAt):
		return model.PaymentIntentsDBFieldName.ExpiresAt, true

	case string(PaymentIntentsDTOFieldName.Requires3ds):
		return model.PaymentIntentsDBFieldName.Requires3ds, true

	case string(PaymentIntentsDTOFieldName.ThreeDsVersion):
		return model.PaymentIntentsDBFieldName.ThreeDsVersion, true

	case string(PaymentIntentsDTOFieldName.Description):
		return model.PaymentIntentsDBFieldName.Description, true

	case string(PaymentIntentsDTOFieldName.StatementDescriptor):
		return model.PaymentIntentsDBFieldName.StatementDescriptor, true

	case string(PaymentIntentsDTOFieldName.Metadata):
		return model.PaymentIntentsDBFieldName.Metadata, true

	case string(PaymentIntentsDTOFieldName.PromoCode):
		return model.PaymentIntentsDBFieldName.PromoCode, true

	case string(PaymentIntentsDTOFieldName.PromoDiscountAmount):
		return model.PaymentIntentsDBFieldName.PromoDiscountAmount, true

	case string(PaymentIntentsDTOFieldName.IdempotencyKeyId):
		return model.PaymentIntentsDBFieldName.IdempotencyKeyId, true

	case string(PaymentIntentsDTOFieldName.ConfirmedAt):
		return model.PaymentIntentsDBFieldName.ConfirmedAt, true

	case string(PaymentIntentsDTOFieldName.CancelledAt):
		return model.PaymentIntentsDBFieldName.CancelledAt, true

	case string(PaymentIntentsDTOFieldName.CancellationReason):
		return model.PaymentIntentsDBFieldName.CancellationReason, true

	case string(PaymentIntentsDTOFieldName.MetaCreatedAt):
		return model.PaymentIntentsDBFieldName.MetaCreatedAt, true

	case string(PaymentIntentsDTOFieldName.MetaCreatedBy):
		return model.PaymentIntentsDBFieldName.MetaCreatedBy, true

	case string(PaymentIntentsDTOFieldName.MetaUpdatedAt):
		return model.PaymentIntentsDBFieldName.MetaUpdatedAt, true

	case string(PaymentIntentsDTOFieldName.MetaUpdatedBy):
		return model.PaymentIntentsDBFieldName.MetaUpdatedBy, true

	case string(PaymentIntentsDTOFieldName.MetaDeletedAt):
		return model.PaymentIntentsDBFieldName.MetaDeletedAt, true

	case string(PaymentIntentsDTOFieldName.MetaDeletedBy):
		return model.PaymentIntentsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformPaymentIntentsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformPaymentIntentsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentIntentsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentIntentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultPaymentIntentsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(PaymentIntentsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentIntentsSelectableResponse map[string]interface{}
type PaymentIntentsSelectableListResponse []*PaymentIntentsSelectableResponse

func NewPaymentIntentsSelectableResponse(paymentIntents model.PaymentIntents, filter model.Filter) PaymentIntentsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentIntentsDBFieldName.Id),
			string(model.PaymentIntentsDBFieldName.IntentCode),
			string(model.PaymentIntentsDBFieldName.MerchantId),
			string(model.PaymentIntentsDBFieldName.OrderId),
			string(model.PaymentIntentsDBFieldName.OrderType),
			string(model.PaymentIntentsDBFieldName.Amount),
			string(model.PaymentIntentsDBFieldName.Currency),
			string(model.PaymentIntentsDBFieldName.TaxAmount),
			string(model.PaymentIntentsDBFieldName.DiscountAmount),
			string(model.PaymentIntentsDBFieldName.TipAmount),
			string(model.PaymentIntentsDBFieldName.UserId),
			string(model.PaymentIntentsDBFieldName.CustomerName),
			string(model.PaymentIntentsDBFieldName.CustomerEmail),
			string(model.PaymentIntentsDBFieldName.CustomerPhone),
			string(model.PaymentIntentsDBFieldName.CustomerIp),
			string(model.PaymentIntentsDBFieldName.CustomerCountry),
			string(model.PaymentIntentsDBFieldName.PaymentMethodId),
			string(model.PaymentIntentsDBFieldName.PaymentMethodType),
			string(model.PaymentIntentsDBFieldName.Status),
			string(model.PaymentIntentsDBFieldName.RoutingProfileId),
			string(model.PaymentIntentsDBFieldName.ExpiresAt),
			string(model.PaymentIntentsDBFieldName.Requires3ds),
			string(model.PaymentIntentsDBFieldName.ThreeDsVersion),
			string(model.PaymentIntentsDBFieldName.Description),
			string(model.PaymentIntentsDBFieldName.StatementDescriptor),
			string(model.PaymentIntentsDBFieldName.Metadata),
			string(model.PaymentIntentsDBFieldName.PromoCode),
			string(model.PaymentIntentsDBFieldName.PromoDiscountAmount),
			string(model.PaymentIntentsDBFieldName.IdempotencyKeyId),
			string(model.PaymentIntentsDBFieldName.ConfirmedAt),
			string(model.PaymentIntentsDBFieldName.CancelledAt),
			string(model.PaymentIntentsDBFieldName.CancellationReason),
			string(model.PaymentIntentsDBFieldName.MetaCreatedAt),
			string(model.PaymentIntentsDBFieldName.MetaCreatedBy),
			string(model.PaymentIntentsDBFieldName.MetaUpdatedAt),
			string(model.PaymentIntentsDBFieldName.MetaUpdatedBy),
			string(model.PaymentIntentsDBFieldName.MetaDeletedAt),
			string(model.PaymentIntentsDBFieldName.MetaDeletedBy),
		)
	}
	paymentIntentsSelectableResponse := PaymentIntentsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.PaymentIntentsDBFieldName.Id):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.Id)] = paymentIntents.Id

		case string(model.PaymentIntentsDBFieldName.IntentCode):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.IntentCode)] = paymentIntents.IntentCode

		case string(model.PaymentIntentsDBFieldName.MerchantId):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.MerchantId)] = paymentIntents.MerchantId

		case string(model.PaymentIntentsDBFieldName.OrderId):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.OrderId)] = paymentIntents.OrderId

		case string(model.PaymentIntentsDBFieldName.OrderType):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.OrderType)] = paymentIntents.OrderType

		case string(model.PaymentIntentsDBFieldName.Amount):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.Amount)] = paymentIntents.Amount

		case string(model.PaymentIntentsDBFieldName.Currency):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.Currency)] = paymentIntents.Currency

		case string(model.PaymentIntentsDBFieldName.TaxAmount):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.TaxAmount)] = paymentIntents.TaxAmount

		case string(model.PaymentIntentsDBFieldName.DiscountAmount):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.DiscountAmount)] = paymentIntents.DiscountAmount

		case string(model.PaymentIntentsDBFieldName.TipAmount):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.TipAmount)] = paymentIntents.TipAmount

		case string(model.PaymentIntentsDBFieldName.UserId):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.UserId)] = paymentIntents.UserId

		case string(model.PaymentIntentsDBFieldName.CustomerName):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.CustomerName)] = paymentIntents.CustomerName

		case string(model.PaymentIntentsDBFieldName.CustomerEmail):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.CustomerEmail)] = paymentIntents.CustomerEmail

		case string(model.PaymentIntentsDBFieldName.CustomerPhone):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.CustomerPhone)] = paymentIntents.CustomerPhone

		case string(model.PaymentIntentsDBFieldName.CustomerIp):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.CustomerIp)] = paymentIntents.CustomerIp

		case string(model.PaymentIntentsDBFieldName.CustomerCountry):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.CustomerCountry)] = paymentIntents.CustomerCountry

		case string(model.PaymentIntentsDBFieldName.PaymentMethodId):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.PaymentMethodId)] = paymentIntents.PaymentMethodId

		case string(model.PaymentIntentsDBFieldName.PaymentMethodType):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.PaymentMethodType)] = paymentIntents.PaymentMethodType

		case string(model.PaymentIntentsDBFieldName.Status):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.Status)] = paymentIntents.Status

		case string(model.PaymentIntentsDBFieldName.RoutingProfileId):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.RoutingProfileId)] = paymentIntents.RoutingProfileId

		case string(model.PaymentIntentsDBFieldName.ExpiresAt):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.ExpiresAt)] = paymentIntents.ExpiresAt

		case string(model.PaymentIntentsDBFieldName.Requires3ds):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.Requires3ds)] = paymentIntents.Requires3ds

		case string(model.PaymentIntentsDBFieldName.ThreeDsVersion):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.ThreeDsVersion)] = paymentIntents.ThreeDsVersion

		case string(model.PaymentIntentsDBFieldName.Description):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.Description)] = paymentIntents.Description

		case string(model.PaymentIntentsDBFieldName.StatementDescriptor):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.StatementDescriptor)] = paymentIntents.StatementDescriptor

		case string(model.PaymentIntentsDBFieldName.Metadata):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.Metadata)] = paymentIntents.Metadata

		case string(model.PaymentIntentsDBFieldName.PromoCode):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.PromoCode)] = paymentIntents.PromoCode

		case string(model.PaymentIntentsDBFieldName.PromoDiscountAmount):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.PromoDiscountAmount)] = paymentIntents.PromoDiscountAmount

		case string(model.PaymentIntentsDBFieldName.IdempotencyKeyId):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.IdempotencyKeyId)] = paymentIntents.IdempotencyKeyId

		case string(model.PaymentIntentsDBFieldName.ConfirmedAt):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.ConfirmedAt)] = paymentIntents.ConfirmedAt

		case string(model.PaymentIntentsDBFieldName.CancelledAt):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.CancelledAt)] = paymentIntents.CancelledAt

		case string(model.PaymentIntentsDBFieldName.CancellationReason):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.CancellationReason)] = paymentIntents.CancellationReason

		case string(model.PaymentIntentsDBFieldName.MetaCreatedAt):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.MetaCreatedAt)] = paymentIntents.MetaCreatedAt

		case string(model.PaymentIntentsDBFieldName.MetaCreatedBy):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.MetaCreatedBy)] = paymentIntents.MetaCreatedBy

		case string(model.PaymentIntentsDBFieldName.MetaUpdatedAt):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.MetaUpdatedAt)] = paymentIntents.MetaUpdatedAt

		case string(model.PaymentIntentsDBFieldName.MetaUpdatedBy):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.MetaUpdatedBy)] = paymentIntents.MetaUpdatedBy

		case string(model.PaymentIntentsDBFieldName.MetaDeletedAt):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.MetaDeletedAt)] = paymentIntents.MetaDeletedAt

		case string(model.PaymentIntentsDBFieldName.MetaDeletedBy):
			paymentIntentsSelectableResponse[string(PaymentIntentsDTOFieldName.MetaDeletedBy)] = paymentIntents.MetaDeletedBy

		}
	}
	return paymentIntentsSelectableResponse
}

type PaymentIntentsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     PaymentIntentsSelectableListResponse `json:"data"`
}

func NewPaymentIntentsFilterResponse(result []model.PaymentIntentsFilterResult, filter model.Filter) (resp PaymentIntentsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewPaymentIntentsListResponseFromFilterResult(result, filter)
	return resp
}

type PaymentIntentsCreateRequest struct {
	IntentCode          string                  `json:"intentCode"`
	MerchantId          uuid.UUID               `json:"merchantId"`
	OrderId             uuid.UUID               `json:"orderId"`
	OrderType           string                  `json:"orderType"`
	Amount              decimal.Decimal         `json:"amount"`
	Currency            model.PaymentCurrency   `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	TaxAmount           decimal.Decimal         `json:"taxAmount"`
	DiscountAmount      decimal.Decimal         `json:"discountAmount"`
	TipAmount           decimal.Decimal         `json:"tipAmount"`
	UserId              uuid.UUID               `json:"userId"`
	CustomerName        string                  `json:"customerName"`
	CustomerEmail       string                  `json:"customerEmail"`
	CustomerPhone       string                  `json:"customerPhone"`
	CustomerIp          string                  `json:"customerIp"`
	CustomerCountry     string                  `json:"customerCountry"`
	PaymentMethodId     uuid.UUID               `json:"paymentMethodId"`
	PaymentMethodType   model.PaymentMethodType `json:"paymentMethodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status              model.PaymentStatus     `json:"status" example:"INITIATED" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	RoutingProfileId    uuid.UUID               `json:"routingProfileId"`
	ExpiresAt           time.Time               `json:"expiresAt"`
	Requires3ds         bool                    `json:"requires3ds"`
	ThreeDsVersion      string                  `json:"threeDsVersion"`
	Description         string                  `json:"description"`
	StatementDescriptor string                  `json:"statementDescriptor"`
	Metadata            json.RawMessage         `json:"metadata"`
	PromoCode           string                  `json:"promoCode"`
	PromoDiscountAmount decimal.Decimal         `json:"promoDiscountAmount"`
	IdempotencyKeyId    uuid.UUID               `json:"idempotencyKeyId"`
	ConfirmedAt         time.Time               `json:"confirmedAt"`
	CancelledAt         time.Time               `json:"cancelledAt"`
	CancellationReason  string                  `json:"cancellationReason"`
	MetaCreatedAt       time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy       uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt       time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy       uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt       time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy       uuid.UUID               `json:"metaDeletedBy"`
}

func (d *PaymentIntentsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentIntentsCreateRequest) ToModel() model.PaymentIntents {
	id, _ := uuid.NewV7()
	return model.PaymentIntents{
		Id:                  id,
		IntentCode:          d.IntentCode,
		MerchantId:          d.MerchantId,
		OrderId:             nuuid.From(d.OrderId),
		OrderType:           null.StringFrom(d.OrderType),
		Amount:              d.Amount,
		Currency:            d.Currency,
		TaxAmount:           d.TaxAmount,
		DiscountAmount:      d.DiscountAmount,
		TipAmount:           d.TipAmount,
		UserId:              nuuid.From(d.UserId),
		CustomerName:        null.StringFrom(d.CustomerName),
		CustomerEmail:       null.StringFrom(d.CustomerEmail),
		CustomerPhone:       null.StringFrom(d.CustomerPhone),
		CustomerIp:          parseCustomerIP(d.CustomerIp),
		CustomerCountry:     null.StringFrom(d.CustomerCountry),
		PaymentMethodId:     nuuid.From(d.PaymentMethodId),
		PaymentMethodType:   d.PaymentMethodType,
		Status:              d.Status,
		RoutingProfileId:    nuuid.From(d.RoutingProfileId),
		ExpiresAt:           d.ExpiresAt,
		Requires3ds:         d.Requires3ds,
		ThreeDsVersion:      null.StringFrom(d.ThreeDsVersion),
		Description:         null.StringFrom(d.Description),
		StatementDescriptor: null.StringFrom(d.StatementDescriptor),
		Metadata:            d.Metadata,
		PromoCode:           null.StringFrom(d.PromoCode),
		PromoDiscountAmount: d.PromoDiscountAmount,
		IdempotencyKeyId:    nuuid.From(d.IdempotencyKeyId),
		ConfirmedAt:         null.TimeFrom(d.ConfirmedAt),
		CancelledAt:         null.TimeFrom(d.CancelledAt),
		CancellationReason:  null.StringFrom(d.CancellationReason),
		MetaCreatedAt:       d.MetaCreatedAt,
		MetaCreatedBy:       d.MetaCreatedBy,
		MetaUpdatedAt:       d.MetaUpdatedAt,
		MetaUpdatedBy:       nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:       null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:       nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentIntentsListCreateRequest []*PaymentIntentsCreateRequest

func (d PaymentIntentsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntents := range d {
		err = validator.Struct(paymentIntents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentIntentsListCreateRequest) ToModelList() []model.PaymentIntents {
	out := make([]model.PaymentIntents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentIntentsUpdateRequest struct {
	IntentCode          string                  `json:"intentCode"`
	MerchantId          uuid.UUID               `json:"merchantId"`
	OrderId             uuid.UUID               `json:"orderId"`
	OrderType           string                  `json:"orderType"`
	Amount              decimal.Decimal         `json:"amount"`
	Currency            model.PaymentCurrency   `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	TaxAmount           decimal.Decimal         `json:"taxAmount"`
	DiscountAmount      decimal.Decimal         `json:"discountAmount"`
	TipAmount           decimal.Decimal         `json:"tipAmount"`
	UserId              uuid.UUID               `json:"userId"`
	CustomerName        string                  `json:"customerName"`
	CustomerEmail       string                  `json:"customerEmail"`
	CustomerPhone       string                  `json:"customerPhone"`
	CustomerIp          string                  `json:"customerIp"`
	CustomerCountry     string                  `json:"customerCountry"`
	PaymentMethodId     uuid.UUID               `json:"paymentMethodId"`
	PaymentMethodType   model.PaymentMethodType `json:"paymentMethodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status              model.PaymentStatus     `json:"status" example:"INITIATED" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	RoutingProfileId    uuid.UUID               `json:"routingProfileId"`
	ExpiresAt           time.Time               `json:"expiresAt"`
	Requires3ds         bool                    `json:"requires3ds"`
	ThreeDsVersion      string                  `json:"threeDsVersion"`
	Description         string                  `json:"description"`
	StatementDescriptor string                  `json:"statementDescriptor"`
	Metadata            json.RawMessage         `json:"metadata"`
	PromoCode           string                  `json:"promoCode"`
	PromoDiscountAmount decimal.Decimal         `json:"promoDiscountAmount"`
	IdempotencyKeyId    uuid.UUID               `json:"idempotencyKeyId"`
	ConfirmedAt         time.Time               `json:"confirmedAt"`
	CancelledAt         time.Time               `json:"cancelledAt"`
	CancellationReason  string                  `json:"cancellationReason"`
	MetaCreatedAt       time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy       uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt       time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy       uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt       time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy       uuid.UUID               `json:"metaDeletedBy"`
}

func (d *PaymentIntentsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentIntentsUpdateRequest) ToModel() model.PaymentIntents {
	return model.PaymentIntents{
		IntentCode:          d.IntentCode,
		MerchantId:          d.MerchantId,
		OrderId:             nuuid.From(d.OrderId),
		OrderType:           null.StringFrom(d.OrderType),
		Amount:              d.Amount,
		Currency:            d.Currency,
		TaxAmount:           d.TaxAmount,
		DiscountAmount:      d.DiscountAmount,
		TipAmount:           d.TipAmount,
		UserId:              nuuid.From(d.UserId),
		CustomerName:        null.StringFrom(d.CustomerName),
		CustomerEmail:       null.StringFrom(d.CustomerEmail),
		CustomerPhone:       null.StringFrom(d.CustomerPhone),
		CustomerIp:          parseCustomerIP(d.CustomerIp),
		CustomerCountry:     null.StringFrom(d.CustomerCountry),
		PaymentMethodId:     nuuid.From(d.PaymentMethodId),
		PaymentMethodType:   d.PaymentMethodType,
		Status:              d.Status,
		RoutingProfileId:    nuuid.From(d.RoutingProfileId),
		ExpiresAt:           d.ExpiresAt,
		Requires3ds:         d.Requires3ds,
		ThreeDsVersion:      null.StringFrom(d.ThreeDsVersion),
		Description:         null.StringFrom(d.Description),
		StatementDescriptor: null.StringFrom(d.StatementDescriptor),
		Metadata:            d.Metadata,
		PromoCode:           null.StringFrom(d.PromoCode),
		PromoDiscountAmount: d.PromoDiscountAmount,
		IdempotencyKeyId:    nuuid.From(d.IdempotencyKeyId),
		ConfirmedAt:         null.TimeFrom(d.ConfirmedAt),
		CancelledAt:         null.TimeFrom(d.CancelledAt),
		CancellationReason:  null.StringFrom(d.CancellationReason),
		MetaCreatedAt:       d.MetaCreatedAt,
		MetaCreatedBy:       d.MetaCreatedBy,
		MetaUpdatedAt:       d.MetaUpdatedAt,
		MetaUpdatedBy:       nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:       null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:       nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentIntentsBulkUpdateRequest struct {
	Id                  uuid.UUID               `json:"id"`
	IntentCode          string                  `json:"intentCode"`
	MerchantId          uuid.UUID               `json:"merchantId"`
	OrderId             uuid.UUID               `json:"orderId"`
	OrderType           string                  `json:"orderType"`
	Amount              decimal.Decimal         `json:"amount"`
	Currency            model.PaymentCurrency   `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	TaxAmount           decimal.Decimal         `json:"taxAmount"`
	DiscountAmount      decimal.Decimal         `json:"discountAmount"`
	TipAmount           decimal.Decimal         `json:"tipAmount"`
	UserId              uuid.UUID               `json:"userId"`
	CustomerName        string                  `json:"customerName"`
	CustomerEmail       string                  `json:"customerEmail"`
	CustomerPhone       string                  `json:"customerPhone"`
	CustomerIp          string                  `json:"customerIp"`
	CustomerCountry     string                  `json:"customerCountry"`
	PaymentMethodId     uuid.UUID               `json:"paymentMethodId"`
	PaymentMethodType   model.PaymentMethodType `json:"paymentMethodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status              model.PaymentStatus     `json:"status" example:"INITIATED" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	RoutingProfileId    uuid.UUID               `json:"routingProfileId"`
	ExpiresAt           time.Time               `json:"expiresAt"`
	Requires3ds         bool                    `json:"requires3ds"`
	ThreeDsVersion      string                  `json:"threeDsVersion"`
	Description         string                  `json:"description"`
	StatementDescriptor string                  `json:"statementDescriptor"`
	Metadata            json.RawMessage         `json:"metadata"`
	PromoCode           string                  `json:"promoCode"`
	PromoDiscountAmount decimal.Decimal         `json:"promoDiscountAmount"`
	IdempotencyKeyId    uuid.UUID               `json:"idempotencyKeyId"`
	ConfirmedAt         time.Time               `json:"confirmedAt"`
	CancelledAt         time.Time               `json:"cancelledAt"`
	CancellationReason  string                  `json:"cancellationReason"`
	MetaCreatedAt       time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy       uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt       time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy       uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt       time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy       uuid.UUID               `json:"metaDeletedBy"`
}

func (d PaymentIntentsBulkUpdateRequest) PrimaryID() PaymentIntentsPrimaryID {
	return PaymentIntentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentIntentsListBulkUpdateRequest []*PaymentIntentsBulkUpdateRequest

func (d PaymentIntentsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntents := range d {
		err = validator.Struct(paymentIntents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentIntentsBulkUpdateRequest) ToModel() model.PaymentIntents {
	return model.PaymentIntents{
		Id:                  d.Id,
		IntentCode:          d.IntentCode,
		MerchantId:          d.MerchantId,
		OrderId:             nuuid.From(d.OrderId),
		OrderType:           null.StringFrom(d.OrderType),
		Amount:              d.Amount,
		Currency:            d.Currency,
		TaxAmount:           d.TaxAmount,
		DiscountAmount:      d.DiscountAmount,
		TipAmount:           d.TipAmount,
		UserId:              nuuid.From(d.UserId),
		CustomerName:        null.StringFrom(d.CustomerName),
		CustomerEmail:       null.StringFrom(d.CustomerEmail),
		CustomerPhone:       null.StringFrom(d.CustomerPhone),
		CustomerIp:          parseCustomerIP(d.CustomerIp),
		CustomerCountry:     null.StringFrom(d.CustomerCountry),
		PaymentMethodId:     nuuid.From(d.PaymentMethodId),
		PaymentMethodType:   d.PaymentMethodType,
		Status:              d.Status,
		RoutingProfileId:    nuuid.From(d.RoutingProfileId),
		ExpiresAt:           d.ExpiresAt,
		Requires3ds:         d.Requires3ds,
		ThreeDsVersion:      null.StringFrom(d.ThreeDsVersion),
		Description:         null.StringFrom(d.Description),
		StatementDescriptor: null.StringFrom(d.StatementDescriptor),
		Metadata:            d.Metadata,
		PromoCode:           null.StringFrom(d.PromoCode),
		PromoDiscountAmount: d.PromoDiscountAmount,
		IdempotencyKeyId:    nuuid.From(d.IdempotencyKeyId),
		ConfirmedAt:         null.TimeFrom(d.ConfirmedAt),
		CancelledAt:         null.TimeFrom(d.CancelledAt),
		CancellationReason:  null.StringFrom(d.CancellationReason),
		MetaCreatedAt:       d.MetaCreatedAt,
		MetaCreatedBy:       d.MetaCreatedBy,
		MetaUpdatedAt:       d.MetaUpdatedAt,
		MetaUpdatedBy:       nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:       null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:       nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentIntentsResponse struct {
	Id                  uuid.UUID               `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IntentCode          string                  `json:"intentCode" validate:"required"`
	MerchantId          uuid.UUID               `json:"merchantId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OrderId             uuid.UUID               `json:"orderId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OrderType           string                  `json:"orderType"`
	Amount              decimal.Decimal         `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency            model.PaymentCurrency   `json:"currency" validate:"required,oneof=IDR USD SGD MYR PHP THB AED EUR GBP JPY" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	TaxAmount           decimal.Decimal         `json:"taxAmount" format:"decimal" example:"100.50"`
	DiscountAmount      decimal.Decimal         `json:"discountAmount" format:"decimal" example:"100.50"`
	TipAmount           decimal.Decimal         `json:"tipAmount" format:"decimal" example:"100.50"`
	UserId              uuid.UUID               `json:"userId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CustomerName        string                  `json:"customerName"`
	CustomerEmail       string                  `json:"customerEmail" validate:"email"`
	CustomerPhone       string                  `json:"customerPhone"`
	CustomerIp          string                  `json:"customerIp"`
	CustomerCountry     string                  `json:"customerCountry"`
	PaymentMethodId     uuid.UUID               `json:"paymentMethodId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentMethodType   model.PaymentMethodType `json:"paymentMethodType" validate:"required,oneof=CARD VIRTUAL_ACCOUNT QRIS EWALLET DIRECT_DEBIT BANK_TRANSFER PAYLATER VOUCHER POINTS CASH_ON_DELIVERY" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Status              model.PaymentStatus     `json:"status" validate:"oneof=INITIATED PENDING AUTHORISED CAPTURED PARTIALLY_CAPTURED COMPLETED FAILED CANCELLED EXPIRED REFUNDING REFUNDED PARTIALLY_REFUNDED DISPUTED CHARGEBACK_WON CHARGEBACK_LOST" enums:"INITIATED,PENDING,AUTHORISED,CAPTURED,PARTIALLY_CAPTURED,COMPLETED,FAILED,CANCELLED,EXPIRED,REFUNDING,REFUNDED,PARTIALLY_REFUNDED,DISPUTED,CHARGEBACK_WON,CHARGEBACK_LOST"`
	RoutingProfileId    uuid.UUID               `json:"routingProfileId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ExpiresAt           time.Time               `json:"expiresAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Requires3ds         bool                    `json:"requires3ds" example:"true"`
	ThreeDsVersion      string                  `json:"threeDsVersion"`
	Description         string                  `json:"description"`
	StatementDescriptor string                  `json:"statementDescriptor"`
	Metadata            json.RawMessage         `json:"metadata" swaggertype:"object"`
	PromoCode           string                  `json:"promoCode"`
	PromoDiscountAmount decimal.Decimal         `json:"promoDiscountAmount" format:"decimal" example:"100.50"`
	IdempotencyKeyId    uuid.UUID               `json:"idempotencyKeyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ConfirmedAt         time.Time               `json:"confirmedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CancelledAt         time.Time               `json:"cancelledAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CancellationReason  string                  `json:"cancellationReason"`
	MetaCreatedAt       time.Time               `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy       uuid.UUID               `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt       time.Time               `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy       uuid.UUID               `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt       time.Time               `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy       uuid.UUID               `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewPaymentIntentsResponse(paymentIntents model.PaymentIntents) PaymentIntentsResponse {
	return PaymentIntentsResponse{
		Id:                  paymentIntents.Id,
		IntentCode:          paymentIntents.IntentCode,
		MerchantId:          paymentIntents.MerchantId,
		OrderId:             paymentIntents.OrderId.UUID,
		OrderType:           paymentIntents.OrderType.String,
		Amount:              paymentIntents.Amount,
		Currency:            model.PaymentCurrency(paymentIntents.Currency),
		TaxAmount:           paymentIntents.TaxAmount,
		DiscountAmount:      paymentIntents.DiscountAmount,
		TipAmount:           paymentIntents.TipAmount,
		UserId:              paymentIntents.UserId.UUID,
		CustomerName:        paymentIntents.CustomerName.String,
		CustomerEmail:       paymentIntents.CustomerEmail.String,
		CustomerPhone:       paymentIntents.CustomerPhone.String,
		CustomerIp:          ipAddrToString(paymentIntents.CustomerIp),
		CustomerCountry:     paymentIntents.CustomerCountry.String,
		PaymentMethodId:     paymentIntents.PaymentMethodId.UUID,
		PaymentMethodType:   model.PaymentMethodType(paymentIntents.PaymentMethodType),
		Status:              model.PaymentStatus(paymentIntents.Status),
		RoutingProfileId:    paymentIntents.RoutingProfileId.UUID,
		ExpiresAt:           paymentIntents.ExpiresAt,
		Requires3ds:         paymentIntents.Requires3ds,
		ThreeDsVersion:      paymentIntents.ThreeDsVersion.String,
		Description:         paymentIntents.Description.String,
		StatementDescriptor: paymentIntents.StatementDescriptor.String,
		Metadata:            paymentIntents.Metadata,
		PromoCode:           paymentIntents.PromoCode.String,
		PromoDiscountAmount: paymentIntents.PromoDiscountAmount,
		IdempotencyKeyId:    paymentIntents.IdempotencyKeyId.UUID,
		ConfirmedAt:         paymentIntents.ConfirmedAt.Time,
		CancelledAt:         paymentIntents.CancelledAt.Time,
		CancellationReason:  paymentIntents.CancellationReason.String,
		MetaCreatedAt:       paymentIntents.MetaCreatedAt,
		MetaCreatedBy:       paymentIntents.MetaCreatedBy,
		MetaUpdatedAt:       paymentIntents.MetaUpdatedAt,
		MetaUpdatedBy:       paymentIntents.MetaUpdatedBy.UUID,
		MetaDeletedAt:       paymentIntents.MetaDeletedAt.Time,
		MetaDeletedBy:       paymentIntents.MetaDeletedBy.UUID,
	}
}

type PaymentIntentsListResponse []*PaymentIntentsResponse

func NewPaymentIntentsListResponse(paymentIntentsList model.PaymentIntentsList) PaymentIntentsListResponse {
	dtoPaymentIntentsListResponse := PaymentIntentsListResponse{}
	for _, paymentIntents := range paymentIntentsList {
		dtoPaymentIntentsResponse := NewPaymentIntentsResponse(*paymentIntents)
		dtoPaymentIntentsListResponse = append(dtoPaymentIntentsListResponse, &dtoPaymentIntentsResponse)
	}
	return dtoPaymentIntentsListResponse
}

type PaymentIntentsPrimaryIDList []PaymentIntentsPrimaryID

func (d PaymentIntentsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntents := range d {
		err = validator.Struct(paymentIntents)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentIntentsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentIntentsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentIntentsPrimaryID) ToModel() model.PaymentIntentsPrimaryID {
	return model.PaymentIntentsPrimaryID{
		Id: d.Id,
	}
}

// parseCustomerIP converts a string IP address to *netip.Addr for storing in PaymentIntents
func parseCustomerIP(ipStr string) *netip.Addr {
	if ipStr == "" {
		return nil
	}
	addr, err := netip.ParseAddr(ipStr)
	if err != nil {
		return nil
	}
	return &addr
}

// ipAddrToString converts *netip.Addr to string for DTO transfer
func ipAddrToString(addr *netip.Addr) string {
	if addr == nil {
		return ""
	}
	return addr.String()
}
