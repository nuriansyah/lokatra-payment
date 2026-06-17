package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type PaymentRefundsDBFieldNameType string

type paymentRefundsDBFieldName struct {
	Id                PaymentRefundsDBFieldNameType
	PaymentIntentId   PaymentRefundsDBFieldNameType
	PaymentAttemptId  PaymentRefundsDBFieldNameType
	RefundCode        PaymentRefundsDBFieldNameType
	Amount            PaymentRefundsDBFieldNameType
	Currency          PaymentRefundsDBFieldNameType
	Reason            PaymentRefundsDBFieldNameType
	Status            PaymentRefundsDBFieldNameType
	ProviderRefundId  PaymentRefundsDBFieldNameType
	ProviderReference PaymentRefundsDBFieldNameType
	RequestedBy       PaymentRefundsDBFieldNameType
	RequestedAt       PaymentRefundsDBFieldNameType
	ApprovedBy        PaymentRefundsDBFieldNameType
	ApprovedAt        PaymentRefundsDBFieldNameType
	RejectedBy        PaymentRefundsDBFieldNameType
	RejectedAt        PaymentRefundsDBFieldNameType
	RejectionReason   PaymentRefundsDBFieldNameType
	ProcessingAt      PaymentRefundsDBFieldNameType
	SucceededAt       PaymentRefundsDBFieldNameType
	FailedAt          PaymentRefundsDBFieldNameType
	FailureCode       PaymentRefundsDBFieldNameType
	FailureMessage    PaymentRefundsDBFieldNameType
	RawRequest        PaymentRefundsDBFieldNameType
	RawResponse       PaymentRefundsDBFieldNameType
	Metadata          PaymentRefundsDBFieldNameType
	MetaCreatedAt     PaymentRefundsDBFieldNameType
	MetaCreatedBy     PaymentRefundsDBFieldNameType
	MetaUpdatedAt     PaymentRefundsDBFieldNameType
	MetaUpdatedBy     PaymentRefundsDBFieldNameType
	MetaDeletedAt     PaymentRefundsDBFieldNameType
	MetaDeletedBy     PaymentRefundsDBFieldNameType
}

var PaymentRefundsDBFieldName = paymentRefundsDBFieldName{
	Id:                "id",
	PaymentIntentId:   "payment_intent_id",
	PaymentAttemptId:  "payment_attempt_id",
	RefundCode:        "refund_code",
	Amount:            "amount",
	Currency:          "currency",
	Reason:            "reason",
	Status:            "status",
	ProviderRefundId:  "provider_refund_id",
	ProviderReference: "provider_reference",
	RequestedBy:       "requested_by",
	RequestedAt:       "requested_at",
	ApprovedBy:        "approved_by",
	ApprovedAt:        "approved_at",
	RejectedBy:        "rejected_by",
	RejectedAt:        "rejected_at",
	RejectionReason:   "rejection_reason",
	ProcessingAt:      "processing_at",
	SucceededAt:       "succeeded_at",
	FailedAt:          "failed_at",
	FailureCode:       "failure_code",
	FailureMessage:    "failure_message",
	RawRequest:        "raw_request",
	RawResponse:       "raw_response",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewPaymentRefundsDBFieldNameFromStr(field string) (dbField PaymentRefundsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentRefundsDBFieldName.Id):
		return PaymentRefundsDBFieldName.Id, true

	case string(PaymentRefundsDBFieldName.PaymentIntentId):
		return PaymentRefundsDBFieldName.PaymentIntentId, true

	case string(PaymentRefundsDBFieldName.PaymentAttemptId):
		return PaymentRefundsDBFieldName.PaymentAttemptId, true

	case string(PaymentRefundsDBFieldName.RefundCode):
		return PaymentRefundsDBFieldName.RefundCode, true

	case string(PaymentRefundsDBFieldName.Amount):
		return PaymentRefundsDBFieldName.Amount, true

	case string(PaymentRefundsDBFieldName.Currency):
		return PaymentRefundsDBFieldName.Currency, true

	case string(PaymentRefundsDBFieldName.Reason):
		return PaymentRefundsDBFieldName.Reason, true

	case string(PaymentRefundsDBFieldName.Status):
		return PaymentRefundsDBFieldName.Status, true

	case string(PaymentRefundsDBFieldName.ProviderRefundId):
		return PaymentRefundsDBFieldName.ProviderRefundId, true

	case string(PaymentRefundsDBFieldName.ProviderReference):
		return PaymentRefundsDBFieldName.ProviderReference, true

	case string(PaymentRefundsDBFieldName.RequestedBy):
		return PaymentRefundsDBFieldName.RequestedBy, true

	case string(PaymentRefundsDBFieldName.RequestedAt):
		return PaymentRefundsDBFieldName.RequestedAt, true

	case string(PaymentRefundsDBFieldName.ApprovedBy):
		return PaymentRefundsDBFieldName.ApprovedBy, true

	case string(PaymentRefundsDBFieldName.ApprovedAt):
		return PaymentRefundsDBFieldName.ApprovedAt, true

	case string(PaymentRefundsDBFieldName.RejectedBy):
		return PaymentRefundsDBFieldName.RejectedBy, true

	case string(PaymentRefundsDBFieldName.RejectedAt):
		return PaymentRefundsDBFieldName.RejectedAt, true

	case string(PaymentRefundsDBFieldName.RejectionReason):
		return PaymentRefundsDBFieldName.RejectionReason, true

	case string(PaymentRefundsDBFieldName.ProcessingAt):
		return PaymentRefundsDBFieldName.ProcessingAt, true

	case string(PaymentRefundsDBFieldName.SucceededAt):
		return PaymentRefundsDBFieldName.SucceededAt, true

	case string(PaymentRefundsDBFieldName.FailedAt):
		return PaymentRefundsDBFieldName.FailedAt, true

	case string(PaymentRefundsDBFieldName.FailureCode):
		return PaymentRefundsDBFieldName.FailureCode, true

	case string(PaymentRefundsDBFieldName.FailureMessage):
		return PaymentRefundsDBFieldName.FailureMessage, true

	case string(PaymentRefundsDBFieldName.RawRequest):
		return PaymentRefundsDBFieldName.RawRequest, true

	case string(PaymentRefundsDBFieldName.RawResponse):
		return PaymentRefundsDBFieldName.RawResponse, true

	case string(PaymentRefundsDBFieldName.Metadata):
		return PaymentRefundsDBFieldName.Metadata, true

	case string(PaymentRefundsDBFieldName.MetaCreatedAt):
		return PaymentRefundsDBFieldName.MetaCreatedAt, true

	case string(PaymentRefundsDBFieldName.MetaCreatedBy):
		return PaymentRefundsDBFieldName.MetaCreatedBy, true

	case string(PaymentRefundsDBFieldName.MetaUpdatedAt):
		return PaymentRefundsDBFieldName.MetaUpdatedAt, true

	case string(PaymentRefundsDBFieldName.MetaUpdatedBy):
		return PaymentRefundsDBFieldName.MetaUpdatedBy, true

	case string(PaymentRefundsDBFieldName.MetaDeletedAt):
		return PaymentRefundsDBFieldName.MetaDeletedAt, true

	case string(PaymentRefundsDBFieldName.MetaDeletedBy):
		return PaymentRefundsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentRefundsFilterJoins = map[string]JoinSpec{}

var PaymentRefundsFilterFields = map[string]FilterFieldSpec{
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
	"refund_code": {
		SourcePath:        "refund_code",
		DefaultOutputPath: "refundCode",
		Column:            "refund_code",
		SQLAlias:          "refund_code",
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
	"reason": {
		SourcePath:        "reason",
		DefaultOutputPath: "reason",
		Column:            "reason",
		SQLAlias:          "reason",
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
	"provider_refund_id": {
		SourcePath:        "provider_refund_id",
		DefaultOutputPath: "providerRefundId",
		Column:            "provider_refund_id",
		SQLAlias:          "provider_refund_id",
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
	"requested_by": {
		SourcePath:        "requested_by",
		DefaultOutputPath: "requestedBy",
		Column:            "requested_by",
		SQLAlias:          "requested_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"requested_at": {
		SourcePath:        "requested_at",
		DefaultOutputPath: "requestedAt",
		Column:            "requested_at",
		SQLAlias:          "requested_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"approved_by": {
		SourcePath:        "approved_by",
		DefaultOutputPath: "approvedBy",
		Column:            "approved_by",
		SQLAlias:          "approved_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"approved_at": {
		SourcePath:        "approved_at",
		DefaultOutputPath: "approvedAt",
		Column:            "approved_at",
		SQLAlias:          "approved_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rejected_by": {
		SourcePath:        "rejected_by",
		DefaultOutputPath: "rejectedBy",
		Column:            "rejected_by",
		SQLAlias:          "rejected_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rejected_at": {
		SourcePath:        "rejected_at",
		DefaultOutputPath: "rejectedAt",
		Column:            "rejected_at",
		SQLAlias:          "rejected_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rejection_reason": {
		SourcePath:        "rejection_reason",
		DefaultOutputPath: "rejectionReason",
		Column:            "rejection_reason",
		SQLAlias:          "rejection_reason",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"processing_at": {
		SourcePath:        "processing_at",
		DefaultOutputPath: "processingAt",
		Column:            "processing_at",
		SQLAlias:          "processing_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"succeeded_at": {
		SourcePath:        "succeeded_at",
		DefaultOutputPath: "succeededAt",
		Column:            "succeeded_at",
		SQLAlias:          "succeeded_at",
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

func NewPaymentRefundsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentRefundsFilterFields[field]
	return
}

type PaymentRefundsFilterResult struct {
	PaymentRefunds
	FilterCount int `db:"count"`
}

func ValidatePaymentRefundsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentRefundsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentRefundsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentRefundsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentRefundsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentRefundsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentRefundsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentRefundsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentRefundStatus string

const (
	PaymentRefundStatusRequested  PaymentRefundStatus = "requested"
	PaymentRefundStatusApproved   PaymentRefundStatus = "approved"
	PaymentRefundStatusRejected   PaymentRefundStatus = "rejected"
	PaymentRefundStatusProcessing PaymentRefundStatus = "processing"
	PaymentRefundStatusSucceeded  PaymentRefundStatus = "succeeded"
	PaymentRefundStatusFailed     PaymentRefundStatus = "failed"
)

type PaymentRefunds struct {
	Id                uuid.UUID           `db:"id"`
	PaymentIntentId   uuid.UUID           `db:"payment_intent_id"`
	PaymentAttemptId  nuuid.NUUID         `db:"payment_attempt_id"`
	RefundCode        string              `db:"refund_code"`
	Amount            decimal.Decimal     `db:"amount"`
	Currency          string              `db:"currency"`
	Reason            null.String         `db:"reason"`
	Status            PaymentRefundStatus `db:"status"`
	ProviderRefundId  null.String         `db:"provider_refund_id"`
	ProviderReference null.String         `db:"provider_reference"`
	RequestedBy       uuid.UUID           `db:"requested_by"`
	RequestedAt       time.Time           `db:"requested_at"`
	ApprovedBy        nuuid.NUUID         `db:"approved_by"`
	ApprovedAt        null.Time           `db:"approved_at"`
	RejectedBy        nuuid.NUUID         `db:"rejected_by"`
	RejectedAt        null.Time           `db:"rejected_at"`
	RejectionReason   null.String         `db:"rejection_reason"`
	ProcessingAt      null.Time           `db:"processing_at"`
	SucceededAt       null.Time           `db:"succeeded_at"`
	FailedAt          null.Time           `db:"failed_at"`
	FailureCode       null.String         `db:"failure_code"`
	FailureMessage    null.String         `db:"failure_message"`
	RawRequest        json.RawMessage     `db:"raw_request"`
	RawResponse       json.RawMessage     `db:"raw_response"`
	Metadata          json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type PaymentRefundsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentRefunds) ToPaymentRefundsPrimaryID() PaymentRefundsPrimaryID {
	return PaymentRefundsPrimaryID{
		Id: d.Id,
	}
}

type PaymentRefundsList []*PaymentRefunds

type PaymentRefundsFilterResultList []*PaymentRefundsFilterResult
