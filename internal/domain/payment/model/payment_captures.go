package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type PaymentCapturesDBFieldNameType string

type paymentCapturesDBFieldName struct {
	Id                     PaymentCapturesDBFieldNameType
	PaymentAuthorizationId PaymentCapturesDBFieldNameType
	PaymentIntentId        PaymentCapturesDBFieldNameType
	Amount                 PaymentCapturesDBFieldNameType
	Currency               PaymentCapturesDBFieldNameType
	Status                 PaymentCapturesDBFieldNameType
	ProviderCaptureId      PaymentCapturesDBFieldNameType
	CapturedAt             PaymentCapturesDBFieldNameType
	FailureCode            PaymentCapturesDBFieldNameType
	FailureMessage         PaymentCapturesDBFieldNameType
	RawRequest             PaymentCapturesDBFieldNameType
	RawResponse            PaymentCapturesDBFieldNameType
	Metadata               PaymentCapturesDBFieldNameType
	MetaCreatedAt          PaymentCapturesDBFieldNameType
	MetaCreatedBy          PaymentCapturesDBFieldNameType
	MetaUpdatedAt          PaymentCapturesDBFieldNameType
	MetaUpdatedBy          PaymentCapturesDBFieldNameType
	MetaDeletedAt          PaymentCapturesDBFieldNameType
	MetaDeletedBy          PaymentCapturesDBFieldNameType
}

var PaymentCapturesDBFieldName = paymentCapturesDBFieldName{
	Id:                     "id",
	PaymentAuthorizationId: "payment_authorization_id",
	PaymentIntentId:        "payment_intent_id",
	Amount:                 "amount",
	Currency:               "currency",
	Status:                 "status",
	ProviderCaptureId:      "provider_capture_id",
	CapturedAt:             "captured_at",
	FailureCode:            "failure_code",
	FailureMessage:         "failure_message",
	RawRequest:             "raw_request",
	RawResponse:            "raw_response",
	Metadata:               "metadata",
	MetaCreatedAt:          "meta_created_at",
	MetaCreatedBy:          "meta_created_by",
	MetaUpdatedAt:          "meta_updated_at",
	MetaUpdatedBy:          "meta_updated_by",
	MetaDeletedAt:          "meta_deleted_at",
	MetaDeletedBy:          "meta_deleted_by",
}

func NewPaymentCapturesDBFieldNameFromStr(field string) (dbField PaymentCapturesDBFieldNameType, found bool) {
	switch field {

	case string(PaymentCapturesDBFieldName.Id):
		return PaymentCapturesDBFieldName.Id, true

	case string(PaymentCapturesDBFieldName.PaymentAuthorizationId):
		return PaymentCapturesDBFieldName.PaymentAuthorizationId, true

	case string(PaymentCapturesDBFieldName.PaymentIntentId):
		return PaymentCapturesDBFieldName.PaymentIntentId, true

	case string(PaymentCapturesDBFieldName.Amount):
		return PaymentCapturesDBFieldName.Amount, true

	case string(PaymentCapturesDBFieldName.Currency):
		return PaymentCapturesDBFieldName.Currency, true

	case string(PaymentCapturesDBFieldName.Status):
		return PaymentCapturesDBFieldName.Status, true

	case string(PaymentCapturesDBFieldName.ProviderCaptureId):
		return PaymentCapturesDBFieldName.ProviderCaptureId, true

	case string(PaymentCapturesDBFieldName.CapturedAt):
		return PaymentCapturesDBFieldName.CapturedAt, true

	case string(PaymentCapturesDBFieldName.FailureCode):
		return PaymentCapturesDBFieldName.FailureCode, true

	case string(PaymentCapturesDBFieldName.FailureMessage):
		return PaymentCapturesDBFieldName.FailureMessage, true

	case string(PaymentCapturesDBFieldName.RawRequest):
		return PaymentCapturesDBFieldName.RawRequest, true

	case string(PaymentCapturesDBFieldName.RawResponse):
		return PaymentCapturesDBFieldName.RawResponse, true

	case string(PaymentCapturesDBFieldName.Metadata):
		return PaymentCapturesDBFieldName.Metadata, true

	case string(PaymentCapturesDBFieldName.MetaCreatedAt):
		return PaymentCapturesDBFieldName.MetaCreatedAt, true

	case string(PaymentCapturesDBFieldName.MetaCreatedBy):
		return PaymentCapturesDBFieldName.MetaCreatedBy, true

	case string(PaymentCapturesDBFieldName.MetaUpdatedAt):
		return PaymentCapturesDBFieldName.MetaUpdatedAt, true

	case string(PaymentCapturesDBFieldName.MetaUpdatedBy):
		return PaymentCapturesDBFieldName.MetaUpdatedBy, true

	case string(PaymentCapturesDBFieldName.MetaDeletedAt):
		return PaymentCapturesDBFieldName.MetaDeletedAt, true

	case string(PaymentCapturesDBFieldName.MetaDeletedBy):
		return PaymentCapturesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentCapturesFilterJoins = map[string]JoinSpec{}

var PaymentCapturesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_authorization_id": {
		SourcePath:        "payment_authorization_id",
		DefaultOutputPath: "paymentAuthorizationId",
		Column:            "payment_authorization_id",
		SQLAlias:          "payment_authorization_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_intent_id": {
		SourcePath:        "payment_intent_id",
		DefaultOutputPath: "paymentIntentId",
		Column:            "payment_intent_id",
		SQLAlias:          "payment_intent_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency": {
		SourcePath:        "currency",
		DefaultOutputPath: "currency",
		Column:            "currency",
		SQLAlias:          "currency",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"status": {
		SourcePath:        "status",
		DefaultOutputPath: "status",
		Column:            "status",
		SQLAlias:          "status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_capture_id": {
		SourcePath:        "provider_capture_id",
		DefaultOutputPath: "providerCaptureId",
		Column:            "provider_capture_id",
		SQLAlias:          "provider_capture_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"captured_at": {
		SourcePath:        "captured_at",
		DefaultOutputPath: "capturedAt",
		Column:            "captured_at",
		SQLAlias:          "captured_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"failure_code": {
		SourcePath:        "failure_code",
		DefaultOutputPath: "failureCode",
		Column:            "failure_code",
		SQLAlias:          "failure_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"failure_message": {
		SourcePath:        "failure_message",
		DefaultOutputPath: "failureMessage",
		Column:            "failure_message",
		SQLAlias:          "failure_message",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"raw_request": {
		SourcePath:        "raw_request",
		DefaultOutputPath: "rawRequest",
		Column:            "raw_request",
		SQLAlias:          "raw_request",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"raw_response": {
		SourcePath:        "raw_response",
		DefaultOutputPath: "rawResponse",
		Column:            "raw_response",
		SQLAlias:          "raw_response",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"metadata": {
		SourcePath:        "metadata",
		DefaultOutputPath: "metadata",
		Column:            "metadata",
		SQLAlias:          "metadata",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_at": {
		SourcePath:        "meta_created_at",
		DefaultOutputPath: "metaCreatedAt",
		Column:            "meta_created_at",
		SQLAlias:          "meta_created_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_by": {
		SourcePath:        "meta_created_by",
		DefaultOutputPath: "metaCreatedBy",
		Column:            "meta_created_by",
		SQLAlias:          "meta_created_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_at": {
		SourcePath:        "meta_updated_at",
		DefaultOutputPath: "metaUpdatedAt",
		Column:            "meta_updated_at",
		SQLAlias:          "meta_updated_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_by": {
		SourcePath:        "meta_updated_by",
		DefaultOutputPath: "metaUpdatedBy",
		Column:            "meta_updated_by",
		SQLAlias:          "meta_updated_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_at": {
		SourcePath:        "meta_deleted_at",
		DefaultOutputPath: "metaDeletedAt",
		Column:            "meta_deleted_at",
		SQLAlias:          "meta_deleted_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_by": {
		SourcePath:        "meta_deleted_by",
		DefaultOutputPath: "metaDeletedBy",
		Column:            "meta_deleted_by",
		SQLAlias:          "meta_deleted_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
}

func NewPaymentCapturesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentCapturesFilterFields[field]
	return
}

type PaymentCapturesFilterResult struct {
	PaymentCaptures
	FilterCount int `db:"count"`
}

func ValidatePaymentCapturesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentCapturesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentCapturesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentCapturesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentCapturesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentCapturesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentCapturesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentCapturesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentCaptureStatus string

const (
	PaymentCaptureStatusRequested PaymentCaptureStatus = "requested"
	PaymentCaptureStatusCaptured  PaymentCaptureStatus = "captured"
	PaymentCaptureStatusFailed    PaymentCaptureStatus = "failed"
)

type PaymentCaptures struct {
	Id                     uuid.UUID            `db:"id"`
	PaymentAuthorizationId uuid.UUID            `db:"payment_authorization_id"`
	PaymentIntentId        uuid.UUID            `db:"payment_intent_id"`
	Amount                 decimal.Decimal      `db:"amount"`
	Currency               string               `db:"currency"`
	Status                 PaymentCaptureStatus `db:"status"`
	ProviderCaptureId      null.String          `db:"provider_capture_id"`
	CapturedAt             null.Time            `db:"captured_at"`
	FailureCode            null.String          `db:"failure_code"`
	FailureMessage         null.String          `db:"failure_message"`
	RawRequest             json.RawMessage      `db:"raw_request"`
	RawResponse            json.RawMessage      `db:"raw_response"`
	Metadata               json.RawMessage      `db:"metadata"`

	shared.MetaSignature
}
type PaymentCapturesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentCaptures) ToPaymentCapturesPrimaryID() PaymentCapturesPrimaryID {
	return PaymentCapturesPrimaryID{
		Id: d.Id,
	}
}

type PaymentCapturesList []*PaymentCaptures

type PaymentCapturesFilterResultList []*PaymentCapturesFilterResult
