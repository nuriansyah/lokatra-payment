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

type PaymentVoidsDBFieldNameType string

type paymentVoidsDBFieldName struct {
	Id                     PaymentVoidsDBFieldNameType
	PaymentAuthorizationId PaymentVoidsDBFieldNameType
	PaymentIntentId        PaymentVoidsDBFieldNameType
	Amount                 PaymentVoidsDBFieldNameType
	Currency               PaymentVoidsDBFieldNameType
	Status                 PaymentVoidsDBFieldNameType
	ProviderVoidId         PaymentVoidsDBFieldNameType
	VoidedAt               PaymentVoidsDBFieldNameType
	FailureCode            PaymentVoidsDBFieldNameType
	FailureMessage         PaymentVoidsDBFieldNameType
	RawRequest             PaymentVoidsDBFieldNameType
	RawResponse            PaymentVoidsDBFieldNameType
	Metadata               PaymentVoidsDBFieldNameType
	MetaCreatedAt          PaymentVoidsDBFieldNameType
	MetaCreatedBy          PaymentVoidsDBFieldNameType
	MetaUpdatedAt          PaymentVoidsDBFieldNameType
	MetaUpdatedBy          PaymentVoidsDBFieldNameType
	MetaDeletedAt          PaymentVoidsDBFieldNameType
	MetaDeletedBy          PaymentVoidsDBFieldNameType
}

var PaymentVoidsDBFieldName = paymentVoidsDBFieldName{
	Id:                     "id",
	PaymentAuthorizationId: "payment_authorization_id",
	PaymentIntentId:        "payment_intent_id",
	Amount:                 "amount",
	Currency:               "currency",
	Status:                 "status",
	ProviderVoidId:         "provider_void_id",
	VoidedAt:               "voided_at",
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

func NewPaymentVoidsDBFieldNameFromStr(field string) (dbField PaymentVoidsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentVoidsDBFieldName.Id):
		return PaymentVoidsDBFieldName.Id, true

	case string(PaymentVoidsDBFieldName.PaymentAuthorizationId):
		return PaymentVoidsDBFieldName.PaymentAuthorizationId, true

	case string(PaymentVoidsDBFieldName.PaymentIntentId):
		return PaymentVoidsDBFieldName.PaymentIntentId, true

	case string(PaymentVoidsDBFieldName.Amount):
		return PaymentVoidsDBFieldName.Amount, true

	case string(PaymentVoidsDBFieldName.Currency):
		return PaymentVoidsDBFieldName.Currency, true

	case string(PaymentVoidsDBFieldName.Status):
		return PaymentVoidsDBFieldName.Status, true

	case string(PaymentVoidsDBFieldName.ProviderVoidId):
		return PaymentVoidsDBFieldName.ProviderVoidId, true

	case string(PaymentVoidsDBFieldName.VoidedAt):
		return PaymentVoidsDBFieldName.VoidedAt, true

	case string(PaymentVoidsDBFieldName.FailureCode):
		return PaymentVoidsDBFieldName.FailureCode, true

	case string(PaymentVoidsDBFieldName.FailureMessage):
		return PaymentVoidsDBFieldName.FailureMessage, true

	case string(PaymentVoidsDBFieldName.RawRequest):
		return PaymentVoidsDBFieldName.RawRequest, true

	case string(PaymentVoidsDBFieldName.RawResponse):
		return PaymentVoidsDBFieldName.RawResponse, true

	case string(PaymentVoidsDBFieldName.Metadata):
		return PaymentVoidsDBFieldName.Metadata, true

	case string(PaymentVoidsDBFieldName.MetaCreatedAt):
		return PaymentVoidsDBFieldName.MetaCreatedAt, true

	case string(PaymentVoidsDBFieldName.MetaCreatedBy):
		return PaymentVoidsDBFieldName.MetaCreatedBy, true

	case string(PaymentVoidsDBFieldName.MetaUpdatedAt):
		return PaymentVoidsDBFieldName.MetaUpdatedAt, true

	case string(PaymentVoidsDBFieldName.MetaUpdatedBy):
		return PaymentVoidsDBFieldName.MetaUpdatedBy, true

	case string(PaymentVoidsDBFieldName.MetaDeletedAt):
		return PaymentVoidsDBFieldName.MetaDeletedAt, true

	case string(PaymentVoidsDBFieldName.MetaDeletedBy):
		return PaymentVoidsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentVoidsFilterJoins = map[string]JoinSpec{}

var PaymentVoidsFilterFields = map[string]FilterFieldSpec{
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
	"provider_void_id": {
		SourcePath:        "provider_void_id",
		DefaultOutputPath: "providerVoidId",
		Column:            "provider_void_id",
		SQLAlias:          "provider_void_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"voided_at": {
		SourcePath:        "voided_at",
		DefaultOutputPath: "voidedAt",
		Column:            "voided_at",
		SQLAlias:          "voided_at",
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

func NewPaymentVoidsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentVoidsFilterFields[field]
	return
}

type PaymentVoidsFilterResult struct {
	PaymentVoids
	FilterCount int `db:"count"`
}

func ValidatePaymentVoidsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentVoidsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentVoidsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentVoidsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentVoidsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentVoidsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentVoidsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentVoidsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentVoidStatus string

const (
	PaymentVoidStatusRequested PaymentVoidStatus = "requested"
	PaymentVoidStatusVoided    PaymentVoidStatus = "voided"
	PaymentVoidStatusFailed    PaymentVoidStatus = "failed"
)

type PaymentVoids struct {
	Id                     uuid.UUID         `db:"id"`
	PaymentAuthorizationId uuid.UUID         `db:"payment_authorization_id"`
	PaymentIntentId        uuid.UUID         `db:"payment_intent_id"`
	Amount                 decimal.Decimal   `db:"amount"`
	Currency               string            `db:"currency"`
	Status                 PaymentVoidStatus `db:"status"`
	ProviderVoidId         null.String       `db:"provider_void_id"`
	VoidedAt               null.Time         `db:"voided_at"`
	FailureCode            null.String       `db:"failure_code"`
	FailureMessage         null.String       `db:"failure_message"`
	RawRequest             json.RawMessage   `db:"raw_request"`
	RawResponse            json.RawMessage   `db:"raw_response"`
	Metadata               json.RawMessage   `db:"metadata"`

	shared.MetaSignature
}
type PaymentVoidsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentVoids) ToPaymentVoidsPrimaryID() PaymentVoidsPrimaryID {
	return PaymentVoidsPrimaryID{
		Id: d.Id,
	}
}

type PaymentVoidsList []*PaymentVoids

type PaymentVoidsFilterResultList []*PaymentVoidsFilterResult
