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

type PaymentAttemptsDBFieldNameType string

type paymentAttemptsDBFieldName struct {
	Id                    PaymentAttemptsDBFieldNameType
	PaymentIntentId       PaymentAttemptsDBFieldNameType
	AttemptNo             PaymentAttemptsDBFieldNameType
	ProviderAccountId     PaymentAttemptsDBFieldNameType
	RouteDecisionId       PaymentAttemptsDBFieldNameType
	ProviderCode          PaymentAttemptsDBFieldNameType
	MethodCode            PaymentAttemptsDBFieldNameType
	ChannelCode           PaymentAttemptsDBFieldNameType
	Amount                PaymentAttemptsDBFieldNameType
	Currency              PaymentAttemptsDBFieldNameType
	Status                PaymentAttemptsDBFieldNameType
	ProviderReference     PaymentAttemptsDBFieldNameType
	ProviderTransactionId PaymentAttemptsDBFieldNameType
	ProviderOrderId       PaymentAttemptsDBFieldNameType
	ProviderPaymentId     PaymentAttemptsDBFieldNameType
	FailureCode           PaymentAttemptsDBFieldNameType
	FailureMessage        PaymentAttemptsDBFieldNameType
	ExpiresAt             PaymentAttemptsDBFieldNameType
	AuthorizedAt          PaymentAttemptsDBFieldNameType
	CapturedAt            PaymentAttemptsDBFieldNameType
	PaidAt                PaymentAttemptsDBFieldNameType
	FailedAt              PaymentAttemptsDBFieldNameType
	CanceledAt            PaymentAttemptsDBFieldNameType
	StatusSyncRequiredAt  PaymentAttemptsDBFieldNameType
	LastStatusSyncAt      PaymentAttemptsDBFieldNameType
	RawRequest            PaymentAttemptsDBFieldNameType
	RawResponse           PaymentAttemptsDBFieldNameType
	Metadata              PaymentAttemptsDBFieldNameType
	MetaCreatedAt         PaymentAttemptsDBFieldNameType
	MetaCreatedBy         PaymentAttemptsDBFieldNameType
	MetaUpdatedAt         PaymentAttemptsDBFieldNameType
	MetaUpdatedBy         PaymentAttemptsDBFieldNameType
	MetaDeletedAt         PaymentAttemptsDBFieldNameType
	MetaDeletedBy         PaymentAttemptsDBFieldNameType
}

var PaymentAttemptsDBFieldName = paymentAttemptsDBFieldName{
	Id:                    "id",
	PaymentIntentId:       "payment_intent_id",
	AttemptNo:             "attempt_no",
	ProviderAccountId:     "provider_account_id",
	RouteDecisionId:       "route_decision_id",
	ProviderCode:          "provider_code",
	MethodCode:            "method_code",
	ChannelCode:           "channel_code",
	Amount:                "amount",
	Currency:              "currency",
	Status:                "status",
	ProviderReference:     "provider_reference",
	ProviderTransactionId: "provider_transaction_id",
	ProviderOrderId:       "provider_order_id",
	ProviderPaymentId:     "provider_payment_id",
	FailureCode:           "failure_code",
	FailureMessage:        "failure_message",
	ExpiresAt:             "expires_at",
	AuthorizedAt:          "authorized_at",
	CapturedAt:            "captured_at",
	PaidAt:                "paid_at",
	FailedAt:              "failed_at",
	CanceledAt:            "canceled_at",
	StatusSyncRequiredAt:  "status_sync_required_at",
	LastStatusSyncAt:      "last_status_sync_at",
	RawRequest:            "raw_request",
	RawResponse:           "raw_response",
	Metadata:              "metadata",
	MetaCreatedAt:         "meta_created_at",
	MetaCreatedBy:         "meta_created_by",
	MetaUpdatedAt:         "meta_updated_at",
	MetaUpdatedBy:         "meta_updated_by",
	MetaDeletedAt:         "meta_deleted_at",
	MetaDeletedBy:         "meta_deleted_by",
}

func NewPaymentAttemptsDBFieldNameFromStr(field string) (dbField PaymentAttemptsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentAttemptsDBFieldName.Id):
		return PaymentAttemptsDBFieldName.Id, true

	case string(PaymentAttemptsDBFieldName.PaymentIntentId):
		return PaymentAttemptsDBFieldName.PaymentIntentId, true

	case string(PaymentAttemptsDBFieldName.AttemptNo):
		return PaymentAttemptsDBFieldName.AttemptNo, true

	case string(PaymentAttemptsDBFieldName.ProviderAccountId):
		return PaymentAttemptsDBFieldName.ProviderAccountId, true

	case string(PaymentAttemptsDBFieldName.RouteDecisionId):
		return PaymentAttemptsDBFieldName.RouteDecisionId, true

	case string(PaymentAttemptsDBFieldName.ProviderCode):
		return PaymentAttemptsDBFieldName.ProviderCode, true

	case string(PaymentAttemptsDBFieldName.MethodCode):
		return PaymentAttemptsDBFieldName.MethodCode, true

	case string(PaymentAttemptsDBFieldName.ChannelCode):
		return PaymentAttemptsDBFieldName.ChannelCode, true

	case string(PaymentAttemptsDBFieldName.Amount):
		return PaymentAttemptsDBFieldName.Amount, true

	case string(PaymentAttemptsDBFieldName.Currency):
		return PaymentAttemptsDBFieldName.Currency, true

	case string(PaymentAttemptsDBFieldName.Status):
		return PaymentAttemptsDBFieldName.Status, true

	case string(PaymentAttemptsDBFieldName.ProviderReference):
		return PaymentAttemptsDBFieldName.ProviderReference, true

	case string(PaymentAttemptsDBFieldName.ProviderTransactionId):
		return PaymentAttemptsDBFieldName.ProviderTransactionId, true

	case string(PaymentAttemptsDBFieldName.ProviderOrderId):
		return PaymentAttemptsDBFieldName.ProviderOrderId, true

	case string(PaymentAttemptsDBFieldName.ProviderPaymentId):
		return PaymentAttemptsDBFieldName.ProviderPaymentId, true

	case string(PaymentAttemptsDBFieldName.FailureCode):
		return PaymentAttemptsDBFieldName.FailureCode, true

	case string(PaymentAttemptsDBFieldName.FailureMessage):
		return PaymentAttemptsDBFieldName.FailureMessage, true

	case string(PaymentAttemptsDBFieldName.ExpiresAt):
		return PaymentAttemptsDBFieldName.ExpiresAt, true

	case string(PaymentAttemptsDBFieldName.AuthorizedAt):
		return PaymentAttemptsDBFieldName.AuthorizedAt, true

	case string(PaymentAttemptsDBFieldName.CapturedAt):
		return PaymentAttemptsDBFieldName.CapturedAt, true

	case string(PaymentAttemptsDBFieldName.PaidAt):
		return PaymentAttemptsDBFieldName.PaidAt, true

	case string(PaymentAttemptsDBFieldName.FailedAt):
		return PaymentAttemptsDBFieldName.FailedAt, true

	case string(PaymentAttemptsDBFieldName.CanceledAt):
		return PaymentAttemptsDBFieldName.CanceledAt, true

	case string(PaymentAttemptsDBFieldName.StatusSyncRequiredAt):
		return PaymentAttemptsDBFieldName.StatusSyncRequiredAt, true

	case string(PaymentAttemptsDBFieldName.LastStatusSyncAt):
		return PaymentAttemptsDBFieldName.LastStatusSyncAt, true

	case string(PaymentAttemptsDBFieldName.RawRequest):
		return PaymentAttemptsDBFieldName.RawRequest, true

	case string(PaymentAttemptsDBFieldName.RawResponse):
		return PaymentAttemptsDBFieldName.RawResponse, true

	case string(PaymentAttemptsDBFieldName.Metadata):
		return PaymentAttemptsDBFieldName.Metadata, true

	case string(PaymentAttemptsDBFieldName.MetaCreatedAt):
		return PaymentAttemptsDBFieldName.MetaCreatedAt, true

	case string(PaymentAttemptsDBFieldName.MetaCreatedBy):
		return PaymentAttemptsDBFieldName.MetaCreatedBy, true

	case string(PaymentAttemptsDBFieldName.MetaUpdatedAt):
		return PaymentAttemptsDBFieldName.MetaUpdatedAt, true

	case string(PaymentAttemptsDBFieldName.MetaUpdatedBy):
		return PaymentAttemptsDBFieldName.MetaUpdatedBy, true

	case string(PaymentAttemptsDBFieldName.MetaDeletedAt):
		return PaymentAttemptsDBFieldName.MetaDeletedAt, true

	case string(PaymentAttemptsDBFieldName.MetaDeletedBy):
		return PaymentAttemptsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentAttemptsFilterJoins = map[string]JoinSpec{}

var PaymentAttemptsFilterFields = map[string]FilterFieldSpec{
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
	"attempt_no": {
		SourcePath:        "attempt_no",
		DefaultOutputPath: "attemptNo",
		Column:            "attempt_no",
		SQLAlias:          "attempt_no",
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
	"route_decision_id": {
		SourcePath:        "route_decision_id",
		DefaultOutputPath: "routeDecisionId",
		Column:            "route_decision_id",
		SQLAlias:          "route_decision_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_code": {
		SourcePath:        "provider_code",
		DefaultOutputPath: "providerCode",
		Column:            "provider_code",
		SQLAlias:          "provider_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"method_code": {
		SourcePath:        "method_code",
		DefaultOutputPath: "methodCode",
		Column:            "method_code",
		SQLAlias:          "method_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"channel_code": {
		SourcePath:        "channel_code",
		DefaultOutputPath: "channelCode",
		Column:            "channel_code",
		SQLAlias:          "channel_code",
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
	"provider_reference": {
		SourcePath:        "provider_reference",
		DefaultOutputPath: "providerReference",
		Column:            "provider_reference",
		SQLAlias:          "provider_reference",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_transaction_id": {
		SourcePath:        "provider_transaction_id",
		DefaultOutputPath: "providerTransactionId",
		Column:            "provider_transaction_id",
		SQLAlias:          "provider_transaction_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_order_id": {
		SourcePath:        "provider_order_id",
		DefaultOutputPath: "providerOrderId",
		Column:            "provider_order_id",
		SQLAlias:          "provider_order_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_payment_id": {
		SourcePath:        "provider_payment_id",
		DefaultOutputPath: "providerPaymentId",
		Column:            "provider_payment_id",
		SQLAlias:          "provider_payment_id",
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
	"expires_at": {
		SourcePath:        "expires_at",
		DefaultOutputPath: "expiresAt",
		Column:            "expires_at",
		SQLAlias:          "expires_at",
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
	"captured_at": {
		SourcePath:        "captured_at",
		DefaultOutputPath: "capturedAt",
		Column:            "captured_at",
		SQLAlias:          "captured_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"paid_at": {
		SourcePath:        "paid_at",
		DefaultOutputPath: "paidAt",
		Column:            "paid_at",
		SQLAlias:          "paid_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"failed_at": {
		SourcePath:        "failed_at",
		DefaultOutputPath: "failedAt",
		Column:            "failed_at",
		SQLAlias:          "failed_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"canceled_at": {
		SourcePath:        "canceled_at",
		DefaultOutputPath: "canceledAt",
		Column:            "canceled_at",
		SQLAlias:          "canceled_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"status_sync_required_at": {
		SourcePath:        "status_sync_required_at",
		DefaultOutputPath: "statusSyncRequiredAt",
		Column:            "status_sync_required_at",
		SQLAlias:          "status_sync_required_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"last_status_sync_at": {
		SourcePath:        "last_status_sync_at",
		DefaultOutputPath: "lastStatusSyncAt",
		Column:            "last_status_sync_at",
		SQLAlias:          "last_status_sync_at",
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

func NewPaymentAttemptsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentAttemptsFilterFields[field]
	return
}

type PaymentAttemptsFilterResult struct {
	PaymentAttempts
	FilterCount int `db:"count"`
}

func ValidatePaymentAttemptsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentAttemptsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentAttemptsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentAttemptsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentAttemptsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentAttemptsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentAttemptsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentAttemptsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentAttemptStatus string

const (
	PaymentAttemptStatusCreated    PaymentAttemptStatus = "created"
	PaymentAttemptStatusPending    PaymentAttemptStatus = "pending"
	PaymentAttemptStatusAuthorized PaymentAttemptStatus = "authorized"
	PaymentAttemptStatusCaptured   PaymentAttemptStatus = "captured"
	PaymentAttemptStatusPaid       PaymentAttemptStatus = "paid"
	PaymentAttemptStatusFailed     PaymentAttemptStatus = "failed"
	PaymentAttemptStatusCanceled   PaymentAttemptStatus = "canceled"
)

type PaymentAttempts struct {
	Id                    uuid.UUID            `db:"id"`
	PaymentIntentId       uuid.UUID            `db:"payment_intent_id"`
	AttemptNo             int                  `db:"attempt_no"`
	ProviderAccountId     nuuid.NUUID          `db:"provider_account_id"`
	RouteDecisionId       nuuid.NUUID          `db:"route_decision_id"`
	ProviderCode          null.String          `db:"provider_code"`
	MethodCode            string               `db:"method_code"`
	ChannelCode           null.String          `db:"channel_code"`
	Amount                decimal.Decimal      `db:"amount"`
	Currency              string               `db:"currency"`
	Status                PaymentAttemptStatus `db:"status"`
	ProviderReference     null.String          `db:"provider_reference"`
	ProviderTransactionId null.String          `db:"provider_transaction_id"`
	ProviderOrderId       null.String          `db:"provider_order_id"`
	ProviderPaymentId     null.String          `db:"provider_payment_id"`
	FailureCode           null.String          `db:"failure_code"`
	FailureMessage        null.String          `db:"failure_message"`
	ExpiresAt             null.Time            `db:"expires_at"`
	AuthorizedAt          null.Time            `db:"authorized_at"`
	CapturedAt            null.Time            `db:"captured_at"`
	PaidAt                null.Time            `db:"paid_at"`
	FailedAt              null.Time            `db:"failed_at"`
	CanceledAt            null.Time            `db:"canceled_at"`
	StatusSyncRequiredAt  null.Time            `db:"status_sync_required_at"`
	LastStatusSyncAt      null.Time            `db:"last_status_sync_at"`
	RawRequest            json.RawMessage      `db:"raw_request"`
	RawResponse           json.RawMessage      `db:"raw_response"`
	Metadata              json.RawMessage      `db:"metadata"`

	shared.MetaSignature
}
type PaymentAttemptsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentAttempts) ToPaymentAttemptsPrimaryID() PaymentAttemptsPrimaryID {
	return PaymentAttemptsPrimaryID{
		Id: d.Id,
	}
}

type PaymentAttemptsList []*PaymentAttempts

type PaymentAttemptsFilterResultList []*PaymentAttemptsFilterResult
