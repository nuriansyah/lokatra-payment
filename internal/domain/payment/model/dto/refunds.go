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

type RefundsDTOFieldNameType string

type refundsDTOFieldName struct {
	Id               RefundsDTOFieldNameType
	RefundCode       RefundsDTOFieldNameType
	PaymentId        RefundsDTOFieldNameType
	IntentId         RefundsDTOFieldNameType
	Amount           RefundsDTOFieldNameType
	Currency         RefundsDTOFieldNameType
	Reason           RefundsDTOFieldNameType
	ReasonDetail     RefundsDTOFieldNameType
	Status           RefundsDTOFieldNameType
	PspRefundId      RefundsDTOFieldNameType
	PspRawResponse   RefundsDTOFieldNameType
	RequestedBy      RefundsDTOFieldNameType
	ReviewedBy       RefundsDTOFieldNameType
	ReviewedAt       RefundsDTOFieldNameType
	ReviewNotes      RefundsDTOFieldNameType
	RefundedAt       RefundsDTOFieldNameType
	EstimatedArrival RefundsDTOFieldNameType
	FailureReason    RefundsDTOFieldNameType
	IdempotencyKeyId RefundsDTOFieldNameType
	Metadata         RefundsDTOFieldNameType
	MetaCreatedAt    RefundsDTOFieldNameType
	MetaCreatedBy    RefundsDTOFieldNameType
	MetaUpdatedAt    RefundsDTOFieldNameType
	MetaUpdatedBy    RefundsDTOFieldNameType
	MetaDeletedAt    RefundsDTOFieldNameType
}

var RefundsDTOFieldName = refundsDTOFieldName{
	Id:               "id",
	RefundCode:       "refundCode",
	PaymentId:        "paymentId",
	IntentId:         "intentId",
	Amount:           "amount",
	Currency:         "currency",
	Reason:           "reason",
	ReasonDetail:     "reasonDetail",
	Status:           "status",
	PspRefundId:      "pspRefundId",
	PspRawResponse:   "pspRawResponse",
	RequestedBy:      "requestedBy",
	ReviewedBy:       "reviewedBy",
	ReviewedAt:       "reviewedAt",
	ReviewNotes:      "reviewNotes",
	RefundedAt:       "refundedAt",
	EstimatedArrival: "estimatedArrival",
	FailureReason:    "failureReason",
	IdempotencyKeyId: "idempotencyKeyId",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
}

func NewRefundsListResponseFromFilterResult(result []model.RefundsFilterResult, filter model.Filter) RefundsSelectableListResponse {
	dtoRefundsListResponse := RefundsSelectableListResponse{}
	for _, refunds := range result {
		dtoRefundsResponse := NewRefundsSelectableResponse(refunds.Refunds, filter)
		dtoRefundsListResponse = append(dtoRefundsListResponse, &dtoRefundsResponse)
	}
	return dtoRefundsListResponse
}

func transformRefundsDTOFieldNameFromStr(field string) (dbField model.RefundsDBFieldNameType, found bool) {
	switch field {

	case string(RefundsDTOFieldName.Id):
		return model.RefundsDBFieldName.Id, true

	case string(RefundsDTOFieldName.RefundCode):
		return model.RefundsDBFieldName.RefundCode, true

	case string(RefundsDTOFieldName.PaymentId):
		return model.RefundsDBFieldName.PaymentId, true

	case string(RefundsDTOFieldName.IntentId):
		return model.RefundsDBFieldName.IntentId, true

	case string(RefundsDTOFieldName.Amount):
		return model.RefundsDBFieldName.Amount, true

	case string(RefundsDTOFieldName.Currency):
		return model.RefundsDBFieldName.Currency, true

	case string(RefundsDTOFieldName.Reason):
		return model.RefundsDBFieldName.Reason, true

	case string(RefundsDTOFieldName.ReasonDetail):
		return model.RefundsDBFieldName.ReasonDetail, true

	case string(RefundsDTOFieldName.Status):
		return model.RefundsDBFieldName.Status, true

	case string(RefundsDTOFieldName.PspRefundId):
		return model.RefundsDBFieldName.PspRefundId, true

	case string(RefundsDTOFieldName.PspRawResponse):
		return model.RefundsDBFieldName.PspRawResponse, true

	case string(RefundsDTOFieldName.RequestedBy):
		return model.RefundsDBFieldName.RequestedBy, true

	case string(RefundsDTOFieldName.ReviewedBy):
		return model.RefundsDBFieldName.ReviewedBy, true

	case string(RefundsDTOFieldName.ReviewedAt):
		return model.RefundsDBFieldName.ReviewedAt, true

	case string(RefundsDTOFieldName.ReviewNotes):
		return model.RefundsDBFieldName.ReviewNotes, true

	case string(RefundsDTOFieldName.RefundedAt):
		return model.RefundsDBFieldName.RefundedAt, true

	case string(RefundsDTOFieldName.EstimatedArrival):
		return model.RefundsDBFieldName.EstimatedArrival, true

	case string(RefundsDTOFieldName.FailureReason):
		return model.RefundsDBFieldName.FailureReason, true

	case string(RefundsDTOFieldName.IdempotencyKeyId):
		return model.RefundsDBFieldName.IdempotencyKeyId, true

	case string(RefundsDTOFieldName.Metadata):
		return model.RefundsDBFieldName.Metadata, true

	case string(RefundsDTOFieldName.MetaCreatedAt):
		return model.RefundsDBFieldName.MetaCreatedAt, true

	case string(RefundsDTOFieldName.MetaCreatedBy):
		return model.RefundsDBFieldName.MetaCreatedBy, true

	case string(RefundsDTOFieldName.MetaUpdatedAt):
		return model.RefundsDBFieldName.MetaUpdatedAt, true

	case string(RefundsDTOFieldName.MetaUpdatedBy):
		return model.RefundsDBFieldName.MetaUpdatedBy, true

	case string(RefundsDTOFieldName.MetaDeletedAt):
		return model.RefundsDBFieldName.MetaDeletedAt, true

	}
	return "unknown", false
}

func ValidateAndTransformRefundsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformRefundsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRefundsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRefundsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultRefundsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(RefundsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RefundsSelectableResponse map[string]interface{}
type RefundsSelectableListResponse []*RefundsSelectableResponse

func NewRefundsSelectableResponse(refunds model.Refunds, filter model.Filter) RefundsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RefundsDBFieldName.Id),
			string(model.RefundsDBFieldName.RefundCode),
			string(model.RefundsDBFieldName.PaymentId),
			string(model.RefundsDBFieldName.IntentId),
			string(model.RefundsDBFieldName.Amount),
			string(model.RefundsDBFieldName.Currency),
			string(model.RefundsDBFieldName.Reason),
			string(model.RefundsDBFieldName.ReasonDetail),
			string(model.RefundsDBFieldName.Status),
			string(model.RefundsDBFieldName.PspRefundId),
			string(model.RefundsDBFieldName.PspRawResponse),
			string(model.RefundsDBFieldName.RequestedBy),
			string(model.RefundsDBFieldName.ReviewedBy),
			string(model.RefundsDBFieldName.ReviewedAt),
			string(model.RefundsDBFieldName.ReviewNotes),
			string(model.RefundsDBFieldName.RefundedAt),
			string(model.RefundsDBFieldName.EstimatedArrival),
			string(model.RefundsDBFieldName.FailureReason),
			string(model.RefundsDBFieldName.IdempotencyKeyId),
			string(model.RefundsDBFieldName.Metadata),
			string(model.RefundsDBFieldName.MetaCreatedAt),
			string(model.RefundsDBFieldName.MetaCreatedBy),
			string(model.RefundsDBFieldName.MetaUpdatedAt),
			string(model.RefundsDBFieldName.MetaUpdatedBy),
			string(model.RefundsDBFieldName.MetaDeletedAt),
		)
	}
	refundsSelectableResponse := RefundsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.RefundsDBFieldName.Id):
			refundsSelectableResponse[string(RefundsDTOFieldName.Id)] = refunds.Id

		case string(model.RefundsDBFieldName.RefundCode):
			refundsSelectableResponse[string(RefundsDTOFieldName.RefundCode)] = refunds.RefundCode

		case string(model.RefundsDBFieldName.PaymentId):
			refundsSelectableResponse[string(RefundsDTOFieldName.PaymentId)] = refunds.PaymentId

		case string(model.RefundsDBFieldName.IntentId):
			refundsSelectableResponse[string(RefundsDTOFieldName.IntentId)] = refunds.IntentId

		case string(model.RefundsDBFieldName.Amount):
			refundsSelectableResponse[string(RefundsDTOFieldName.Amount)] = refunds.Amount

		case string(model.RefundsDBFieldName.Currency):
			refundsSelectableResponse[string(RefundsDTOFieldName.Currency)] = refunds.Currency

		case string(model.RefundsDBFieldName.Reason):
			refundsSelectableResponse[string(RefundsDTOFieldName.Reason)] = refunds.Reason

		case string(model.RefundsDBFieldName.ReasonDetail):
			refundsSelectableResponse[string(RefundsDTOFieldName.ReasonDetail)] = refunds.ReasonDetail

		case string(model.RefundsDBFieldName.Status):
			refundsSelectableResponse[string(RefundsDTOFieldName.Status)] = refunds.Status

		case string(model.RefundsDBFieldName.PspRefundId):
			refundsSelectableResponse[string(RefundsDTOFieldName.PspRefundId)] = refunds.PspRefundId

		case string(model.RefundsDBFieldName.PspRawResponse):
			refundsSelectableResponse[string(RefundsDTOFieldName.PspRawResponse)] = refunds.PspRawResponse

		case string(model.RefundsDBFieldName.RequestedBy):
			refundsSelectableResponse[string(RefundsDTOFieldName.RequestedBy)] = refunds.RequestedBy

		case string(model.RefundsDBFieldName.ReviewedBy):
			refundsSelectableResponse[string(RefundsDTOFieldName.ReviewedBy)] = refunds.ReviewedBy

		case string(model.RefundsDBFieldName.ReviewedAt):
			refundsSelectableResponse[string(RefundsDTOFieldName.ReviewedAt)] = refunds.ReviewedAt

		case string(model.RefundsDBFieldName.ReviewNotes):
			refundsSelectableResponse[string(RefundsDTOFieldName.ReviewNotes)] = refunds.ReviewNotes

		case string(model.RefundsDBFieldName.RefundedAt):
			refundsSelectableResponse[string(RefundsDTOFieldName.RefundedAt)] = refunds.RefundedAt

		case string(model.RefundsDBFieldName.EstimatedArrival):
			refundsSelectableResponse[string(RefundsDTOFieldName.EstimatedArrival)] = refunds.EstimatedArrival

		case string(model.RefundsDBFieldName.FailureReason):
			refundsSelectableResponse[string(RefundsDTOFieldName.FailureReason)] = refunds.FailureReason

		case string(model.RefundsDBFieldName.IdempotencyKeyId):
			refundsSelectableResponse[string(RefundsDTOFieldName.IdempotencyKeyId)] = refunds.IdempotencyKeyId

		case string(model.RefundsDBFieldName.Metadata):
			refundsSelectableResponse[string(RefundsDTOFieldName.Metadata)] = refunds.Metadata

		case string(model.RefundsDBFieldName.MetaCreatedAt):
			refundsSelectableResponse[string(RefundsDTOFieldName.MetaCreatedAt)] = refunds.MetaCreatedAt

		case string(model.RefundsDBFieldName.MetaCreatedBy):
			refundsSelectableResponse[string(RefundsDTOFieldName.MetaCreatedBy)] = refunds.MetaCreatedBy

		case string(model.RefundsDBFieldName.MetaUpdatedAt):
			refundsSelectableResponse[string(RefundsDTOFieldName.MetaUpdatedAt)] = refunds.MetaUpdatedAt

		case string(model.RefundsDBFieldName.MetaUpdatedBy):
			refundsSelectableResponse[string(RefundsDTOFieldName.MetaUpdatedBy)] = refunds.MetaUpdatedBy

		case string(model.RefundsDBFieldName.MetaDeletedAt):
			refundsSelectableResponse[string(RefundsDTOFieldName.MetaDeletedAt)] = refunds.MetaDeletedAt

		}
	}
	return refundsSelectableResponse
}

type RefundsFilterResponse struct {
	Metadata Metadata                      `json:"metadata"`
	Data     RefundsSelectableListResponse `json:"data"`
}

func NewRefundsFilterResponse(result []model.RefundsFilterResult, filter model.Filter) (resp RefundsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewRefundsListResponseFromFilterResult(result, filter)
	return resp
}

type RefundsCreateRequest struct {
	RefundCode       string                `json:"refundCode"`
	PaymentId        uuid.UUID             `json:"paymentId"`
	IntentId         uuid.UUID             `json:"intentId"`
	Amount           decimal.Decimal       `json:"amount"`
	Currency         model.PaymentCurrency `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Reason           string                `json:"reason"`
	ReasonDetail     string                `json:"reasonDetail"`
	Status           model.RefundStatus    `json:"status" example:"PENDING" enums:"PENDING,PROCESSING,SUCCEEDED,FAILED,CANCELLED"`
	PspRefundId      string                `json:"pspRefundId"`
	PspRawResponse   json.RawMessage       `json:"pspRawResponse"`
	RequestedBy      uuid.UUID             `json:"requestedBy"`
	ReviewedBy       uuid.UUID             `json:"reviewedBy"`
	ReviewedAt       time.Time             `json:"reviewedAt"`
	ReviewNotes      string                `json:"reviewNotes"`
	RefundedAt       time.Time             `json:"refundedAt"`
	EstimatedArrival time.Time             `json:"estimatedArrival"`
	FailureReason    string                `json:"failureReason"`
	IdempotencyKeyId uuid.UUID             `json:"idempotencyKeyId"`
	Metadata         json.RawMessage       `json:"metadata"`
	MetaCreatedAt    time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt    time.Time             `json:"metaDeletedAt"`
}

func (d *RefundsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RefundsCreateRequest) ToModel() model.Refunds {
	id, _ := uuid.NewV7()
	return model.Refunds{
		Id:               id,
		RefundCode:       d.RefundCode,
		PaymentId:        d.PaymentId,
		IntentId:         d.IntentId,
		Amount:           d.Amount,
		Currency:         d.Currency,
		Reason:           d.Reason,
		ReasonDetail:     null.StringFrom(d.ReasonDetail),
		Status:           d.Status,
		PspRefundId:      null.StringFrom(d.PspRefundId),
		PspRawResponse:   d.PspRawResponse,
		RequestedBy:      d.RequestedBy,
		ReviewedBy:       nuuid.From(d.ReviewedBy),
		ReviewedAt:       null.TimeFrom(d.ReviewedAt),
		ReviewNotes:      null.StringFrom(d.ReviewNotes),
		RefundedAt:       null.TimeFrom(d.RefundedAt),
		EstimatedArrival: null.TimeFrom(d.EstimatedArrival),
		FailureReason:    null.StringFrom(d.FailureReason),
		IdempotencyKeyId: nuuid.From(d.IdempotencyKeyId),
		Metadata:         d.Metadata,
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:    null.TimeFrom(d.MetaDeletedAt),
	}
}

type RefundsListCreateRequest []*RefundsCreateRequest

func (d RefundsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refunds := range d {
		err = validator.Struct(refunds)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundsListCreateRequest) ToModelList() []model.Refunds {
	out := make([]model.Refunds, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RefundsUpdateRequest struct {
	RefundCode       string                `json:"refundCode"`
	PaymentId        uuid.UUID             `json:"paymentId"`
	IntentId         uuid.UUID             `json:"intentId"`
	Amount           decimal.Decimal       `json:"amount"`
	Currency         model.PaymentCurrency `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Reason           string                `json:"reason"`
	ReasonDetail     string                `json:"reasonDetail"`
	Status           model.RefundStatus    `json:"status" example:"PENDING" enums:"PENDING,PROCESSING,SUCCEEDED,FAILED,CANCELLED"`
	PspRefundId      string                `json:"pspRefundId"`
	PspRawResponse   json.RawMessage       `json:"pspRawResponse"`
	RequestedBy      uuid.UUID             `json:"requestedBy"`
	ReviewedBy       uuid.UUID             `json:"reviewedBy"`
	ReviewedAt       time.Time             `json:"reviewedAt"`
	ReviewNotes      string                `json:"reviewNotes"`
	RefundedAt       time.Time             `json:"refundedAt"`
	EstimatedArrival time.Time             `json:"estimatedArrival"`
	FailureReason    string                `json:"failureReason"`
	IdempotencyKeyId uuid.UUID             `json:"idempotencyKeyId"`
	Metadata         json.RawMessage       `json:"metadata"`
	MetaCreatedAt    time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt    time.Time             `json:"metaDeletedAt"`
}

func (d *RefundsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RefundsUpdateRequest) ToModel() model.Refunds {
	return model.Refunds{
		RefundCode:       d.RefundCode,
		PaymentId:        d.PaymentId,
		IntentId:         d.IntentId,
		Amount:           d.Amount,
		Currency:         d.Currency,
		Reason:           d.Reason,
		ReasonDetail:     null.StringFrom(d.ReasonDetail),
		Status:           d.Status,
		PspRefundId:      null.StringFrom(d.PspRefundId),
		PspRawResponse:   d.PspRawResponse,
		RequestedBy:      d.RequestedBy,
		ReviewedBy:       nuuid.From(d.ReviewedBy),
		ReviewedAt:       null.TimeFrom(d.ReviewedAt),
		ReviewNotes:      null.StringFrom(d.ReviewNotes),
		RefundedAt:       null.TimeFrom(d.RefundedAt),
		EstimatedArrival: null.TimeFrom(d.EstimatedArrival),
		FailureReason:    null.StringFrom(d.FailureReason),
		IdempotencyKeyId: nuuid.From(d.IdempotencyKeyId),
		Metadata:         d.Metadata,
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:    null.TimeFrom(d.MetaDeletedAt),
	}
}

type RefundsBulkUpdateRequest struct {
	Id               uuid.UUID             `json:"id"`
	RefundCode       string                `json:"refundCode"`
	PaymentId        uuid.UUID             `json:"paymentId"`
	IntentId         uuid.UUID             `json:"intentId"`
	Amount           decimal.Decimal       `json:"amount"`
	Currency         model.PaymentCurrency `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Reason           string                `json:"reason"`
	ReasonDetail     string                `json:"reasonDetail"`
	Status           model.RefundStatus    `json:"status" example:"PENDING" enums:"PENDING,PROCESSING,SUCCEEDED,FAILED,CANCELLED"`
	PspRefundId      string                `json:"pspRefundId"`
	PspRawResponse   json.RawMessage       `json:"pspRawResponse"`
	RequestedBy      uuid.UUID             `json:"requestedBy"`
	ReviewedBy       uuid.UUID             `json:"reviewedBy"`
	ReviewedAt       time.Time             `json:"reviewedAt"`
	ReviewNotes      string                `json:"reviewNotes"`
	RefundedAt       time.Time             `json:"refundedAt"`
	EstimatedArrival time.Time             `json:"estimatedArrival"`
	FailureReason    string                `json:"failureReason"`
	IdempotencyKeyId uuid.UUID             `json:"idempotencyKeyId"`
	Metadata         json.RawMessage       `json:"metadata"`
	MetaCreatedAt    time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt    time.Time             `json:"metaDeletedAt"`
}

func (d RefundsBulkUpdateRequest) PrimaryID() RefundsPrimaryID {
	return RefundsPrimaryID{
		Id: d.Id,
	}
}

type RefundsListBulkUpdateRequest []*RefundsBulkUpdateRequest

func (d RefundsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refunds := range d {
		err = validator.Struct(refunds)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RefundsBulkUpdateRequest) ToModel() model.Refunds {
	return model.Refunds{
		Id:               d.Id,
		RefundCode:       d.RefundCode,
		PaymentId:        d.PaymentId,
		IntentId:         d.IntentId,
		Amount:           d.Amount,
		Currency:         d.Currency,
		Reason:           d.Reason,
		ReasonDetail:     null.StringFrom(d.ReasonDetail),
		Status:           d.Status,
		PspRefundId:      null.StringFrom(d.PspRefundId),
		PspRawResponse:   d.PspRawResponse,
		RequestedBy:      d.RequestedBy,
		ReviewedBy:       nuuid.From(d.ReviewedBy),
		ReviewedAt:       null.TimeFrom(d.ReviewedAt),
		ReviewNotes:      null.StringFrom(d.ReviewNotes),
		RefundedAt:       null.TimeFrom(d.RefundedAt),
		EstimatedArrival: null.TimeFrom(d.EstimatedArrival),
		FailureReason:    null.StringFrom(d.FailureReason),
		IdempotencyKeyId: nuuid.From(d.IdempotencyKeyId),
		Metadata:         d.Metadata,
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:    null.TimeFrom(d.MetaDeletedAt),
	}
}

type RefundsResponse struct {
	Id               uuid.UUID             `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RefundCode       string                `json:"refundCode" validate:"required"`
	PaymentId        uuid.UUID             `json:"paymentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IntentId         uuid.UUID             `json:"intentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount           decimal.Decimal       `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency         model.PaymentCurrency `json:"currency" validate:"required,oneof=IDR USD SGD MYR PHP THB AED EUR GBP JPY" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Reason           string                `json:"reason" validate:"required"`
	ReasonDetail     string                `json:"reasonDetail"`
	Status           model.RefundStatus    `json:"status" validate:"oneof=PENDING PROCESSING SUCCEEDED FAILED CANCELLED" enums:"PENDING,PROCESSING,SUCCEEDED,FAILED,CANCELLED"`
	PspRefundId      string                `json:"pspRefundId"`
	PspRawResponse   json.RawMessage       `json:"pspRawResponse" validate:"required" swaggertype:"object"`
	RequestedBy      uuid.UUID             `json:"requestedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReviewedBy       uuid.UUID             `json:"reviewedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReviewedAt       time.Time             `json:"reviewedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ReviewNotes      string                `json:"reviewNotes"`
	RefundedAt       time.Time             `json:"refundedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	EstimatedArrival time.Time             `json:"estimatedArrival" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FailureReason    string                `json:"failureReason"`
	IdempotencyKeyId uuid.UUID             `json:"idempotencyKeyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Metadata         json.RawMessage       `json:"metadata" swaggertype:"object"`
	MetaCreatedAt    time.Time             `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy    uuid.UUID             `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt    time.Time             `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy    uuid.UUID             `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt    time.Time             `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
}

func NewRefundsResponse(refunds model.Refunds) RefundsResponse {
	return RefundsResponse{
		Id:               refunds.Id,
		RefundCode:       refunds.RefundCode,
		PaymentId:        refunds.PaymentId,
		IntentId:         refunds.IntentId,
		Amount:           refunds.Amount,
		Currency:         model.PaymentCurrency(refunds.Currency),
		Reason:           refunds.Reason,
		ReasonDetail:     refunds.ReasonDetail.String,
		Status:           model.RefundStatus(refunds.Status),
		PspRefundId:      refunds.PspRefundId.String,
		PspRawResponse:   refunds.PspRawResponse,
		RequestedBy:      refunds.RequestedBy,
		ReviewedBy:       refunds.ReviewedBy.UUID,
		ReviewedAt:       refunds.ReviewedAt.Time,
		ReviewNotes:      refunds.ReviewNotes.String,
		RefundedAt:       refunds.RefundedAt.Time,
		EstimatedArrival: refunds.EstimatedArrival.Time,
		FailureReason:    refunds.FailureReason.String,
		IdempotencyKeyId: refunds.IdempotencyKeyId.UUID,
		Metadata:         refunds.Metadata,
		MetaCreatedAt:    refunds.MetaCreatedAt,
		MetaCreatedBy:    refunds.MetaCreatedBy,
		MetaUpdatedAt:    refunds.MetaUpdatedAt,
		MetaUpdatedBy:    refunds.MetaUpdatedBy.UUID,
		MetaDeletedAt:    refunds.MetaDeletedAt.Time,
	}
}

type RefundsListResponse []*RefundsResponse

func NewRefundsListResponse(refundsList model.RefundsList) RefundsListResponse {
	dtoRefundsListResponse := RefundsListResponse{}
	for _, refunds := range refundsList {
		dtoRefundsResponse := NewRefundsResponse(*refunds)
		dtoRefundsListResponse = append(dtoRefundsListResponse, &dtoRefundsResponse)
	}
	return dtoRefundsListResponse
}

type RefundsPrimaryIDList []RefundsPrimaryID

func (d RefundsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, refunds := range d {
		err = validator.Struct(refunds)
		if err != nil {
			return
		}
	}
	return nil
}

type RefundsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RefundsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RefundsPrimaryID) ToModel() model.RefundsPrimaryID {
	return model.RefundsPrimaryID{
		Id: d.Id,
	}
}
