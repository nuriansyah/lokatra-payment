package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type PaymentAuthorizationsDBFieldNameType string

type paymentAuthorizationsDBFieldName struct {
	Id                      PaymentAuthorizationsDBFieldNameType
	PaymentIntentId         PaymentAuthorizationsDBFieldNameType
	PaymentAttemptId        PaymentAuthorizationsDBFieldNameType
	ProviderAccountId       PaymentAuthorizationsDBFieldNameType
	ProviderAuthorizationId PaymentAuthorizationsDBFieldNameType
	Amount                  PaymentAuthorizationsDBFieldNameType
	Currency                PaymentAuthorizationsDBFieldNameType
	Status                  PaymentAuthorizationsDBFieldNameType
	AuthorizedAt            PaymentAuthorizationsDBFieldNameType
	ExpiresAt               PaymentAuthorizationsDBFieldNameType
	CapturedAmount          PaymentAuthorizationsDBFieldNameType
	FailureCode             PaymentAuthorizationsDBFieldNameType
	FailureMessage          PaymentAuthorizationsDBFieldNameType
	RawRequest              PaymentAuthorizationsDBFieldNameType
	RawResponse             PaymentAuthorizationsDBFieldNameType
	Metadata                PaymentAuthorizationsDBFieldNameType
	MetaCreatedAt           PaymentAuthorizationsDBFieldNameType
	MetaCreatedBy           PaymentAuthorizationsDBFieldNameType
	MetaUpdatedAt           PaymentAuthorizationsDBFieldNameType
	MetaUpdatedBy           PaymentAuthorizationsDBFieldNameType
	MetaDeletedAt           PaymentAuthorizationsDBFieldNameType
	MetaDeletedBy           PaymentAuthorizationsDBFieldNameType
}

var PaymentAuthorizationsDBFieldName = paymentAuthorizationsDBFieldName{
	Id:                      "id",
	PaymentIntentId:         "payment_intent_id",
	PaymentAttemptId:        "payment_attempt_id",
	ProviderAccountId:       "provider_account_id",
	ProviderAuthorizationId: "provider_authorization_id",
	Amount:                  "amount",
	Currency:                "currency",
	Status:                  "status",
	AuthorizedAt:            "authorized_at",
	ExpiresAt:               "expires_at",
	CapturedAmount:          "captured_amount",
	FailureCode:             "failure_code",
	FailureMessage:          "failure_message",
	RawRequest:              "raw_request",
	RawResponse:             "raw_response",
	Metadata:                "metadata",
	MetaCreatedAt:           "meta_created_at",
	MetaCreatedBy:           "meta_created_by",
	MetaUpdatedAt:           "meta_updated_at",
	MetaUpdatedBy:           "meta_updated_by",
	MetaDeletedAt:           "meta_deleted_at",
	MetaDeletedBy:           "meta_deleted_by",
}

func NewPaymentAuthorizationsDBFieldNameFromStr(field string) (dbField PaymentAuthorizationsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentAuthorizationsDBFieldName.Id):
		return PaymentAuthorizationsDBFieldName.Id, true

	case string(PaymentAuthorizationsDBFieldName.PaymentIntentId):
		return PaymentAuthorizationsDBFieldName.PaymentIntentId, true

	case string(PaymentAuthorizationsDBFieldName.PaymentAttemptId):
		return PaymentAuthorizationsDBFieldName.PaymentAttemptId, true

	case string(PaymentAuthorizationsDBFieldName.ProviderAccountId):
		return PaymentAuthorizationsDBFieldName.ProviderAccountId, true

	case string(PaymentAuthorizationsDBFieldName.ProviderAuthorizationId):
		return PaymentAuthorizationsDBFieldName.ProviderAuthorizationId, true

	case string(PaymentAuthorizationsDBFieldName.Amount):
		return PaymentAuthorizationsDBFieldName.Amount, true

	case string(PaymentAuthorizationsDBFieldName.Currency):
		return PaymentAuthorizationsDBFieldName.Currency, true

	case string(PaymentAuthorizationsDBFieldName.Status):
		return PaymentAuthorizationsDBFieldName.Status, true

	case string(PaymentAuthorizationsDBFieldName.AuthorizedAt):
		return PaymentAuthorizationsDBFieldName.AuthorizedAt, true

	case string(PaymentAuthorizationsDBFieldName.ExpiresAt):
		return PaymentAuthorizationsDBFieldName.ExpiresAt, true

	case string(PaymentAuthorizationsDBFieldName.CapturedAmount):
		return PaymentAuthorizationsDBFieldName.CapturedAmount, true

	case string(PaymentAuthorizationsDBFieldName.FailureCode):
		return PaymentAuthorizationsDBFieldName.FailureCode, true

	case string(PaymentAuthorizationsDBFieldName.FailureMessage):
		return PaymentAuthorizationsDBFieldName.FailureMessage, true

	case string(PaymentAuthorizationsDBFieldName.RawRequest):
		return PaymentAuthorizationsDBFieldName.RawRequest, true

	case string(PaymentAuthorizationsDBFieldName.RawResponse):
		return PaymentAuthorizationsDBFieldName.RawResponse, true

	case string(PaymentAuthorizationsDBFieldName.Metadata):
		return PaymentAuthorizationsDBFieldName.Metadata, true

	case string(PaymentAuthorizationsDBFieldName.MetaCreatedAt):
		return PaymentAuthorizationsDBFieldName.MetaCreatedAt, true

	case string(PaymentAuthorizationsDBFieldName.MetaCreatedBy):
		return PaymentAuthorizationsDBFieldName.MetaCreatedBy, true

	case string(PaymentAuthorizationsDBFieldName.MetaUpdatedAt):
		return PaymentAuthorizationsDBFieldName.MetaUpdatedAt, true

	case string(PaymentAuthorizationsDBFieldName.MetaUpdatedBy):
		return PaymentAuthorizationsDBFieldName.MetaUpdatedBy, true

	case string(PaymentAuthorizationsDBFieldName.MetaDeletedAt):
		return PaymentAuthorizationsDBFieldName.MetaDeletedAt, true

	case string(PaymentAuthorizationsDBFieldName.MetaDeletedBy):
		return PaymentAuthorizationsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentAuthorizationsFilterJoins = map[string]JoinSpec{}

var PaymentAuthorizationsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
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
	"payment_attempt_id": {
		SourcePath:        "payment_attempt_id",
		DefaultOutputPath: "paymentAttemptId",
		Column:            "payment_attempt_id",
		SQLAlias:          "payment_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_account_id": {
		SourcePath:        "provider_account_id",
		DefaultOutputPath: "providerAccountId",
		Column:            "provider_account_id",
		SQLAlias:          "provider_account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_authorization_id": {
		SourcePath:        "provider_authorization_id",
		DefaultOutputPath: "providerAuthorizationId",
		Column:            "provider_authorization_id",
		SQLAlias:          "provider_authorization_id",
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
	"authorized_at": {
		SourcePath:        "authorized_at",
		DefaultOutputPath: "authorizedAt",
		Column:            "authorized_at",
		SQLAlias:          "authorized_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"expires_at": {
		SourcePath:        "expires_at",
		DefaultOutputPath: "expiresAt",
		Column:            "expires_at",
		SQLAlias:          "expires_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"captured_amount": {
		SourcePath:        "captured_amount",
		DefaultOutputPath: "capturedAmount",
		Column:            "captured_amount",
		SQLAlias:          "captured_amount",
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

func NewPaymentAuthorizationsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentAuthorizationsFilterFields[field]
	return
}

type PaymentAuthorizationsFilterResult struct {
	PaymentAuthorizations
	FilterCount int `db:"count"`
}

func ValidatePaymentAuthorizationsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentAuthorizationsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentAuthorizationsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentAuthorizationsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentAuthorizationsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentAuthorizationsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentAuthorizationsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentAuthorizationsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentAuthorizationStatus string

const (
	PaymentAuthorizationStatusRequested  PaymentAuthorizationStatus = "requested"
	PaymentAuthorizationStatusAuthorized PaymentAuthorizationStatus = "authorized"
	PaymentAuthorizationStatusCaptured   PaymentAuthorizationStatus = "captured"
	PaymentAuthorizationStatusVoided     PaymentAuthorizationStatus = "voided"
	PaymentAuthorizationStatusFailed     PaymentAuthorizationStatus = "failed"
)

type PaymentAuthorizations struct {
	Id                      uuid.UUID                  `db:"id"`
	PaymentIntentId         uuid.UUID                  `db:"payment_intent_id"`
	PaymentAttemptId        nuuid.NUUID                `db:"payment_attempt_id"`
	ProviderAccountId       nuuid.NUUID                `db:"provider_account_id"`
	ProviderAuthorizationId null.String                `db:"provider_authorization_id"`
	Amount                  decimal.Decimal            `db:"amount"`
	Currency                string                     `db:"currency"`
	Status                  PaymentAuthorizationStatus `db:"status"`
	AuthorizedAt            null.Time                  `db:"authorized_at"`
	ExpiresAt               null.Time                  `db:"expires_at"`
	CapturedAmount          decimal.Decimal            `db:"captured_amount"`
	FailureCode             null.String                `db:"failure_code"`
	FailureMessage          null.String                `db:"failure_message"`
	RawRequest              json.RawMessage            `db:"raw_request"`
	RawResponse             json.RawMessage            `db:"raw_response"`
	Metadata                json.RawMessage            `db:"metadata"`

	shared.MetaSignature
}
type PaymentAuthorizationsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentAuthorizations) ToPaymentAuthorizationsPrimaryID() PaymentAuthorizationsPrimaryID {
	return PaymentAuthorizationsPrimaryID{
		Id: d.Id,
	}
}

type PaymentAuthorizationsList []*PaymentAuthorizations

type PaymentAuthorizationsFilterResultList []*PaymentAuthorizationsFilterResult
