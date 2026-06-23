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

type RefundAttemptsDBFieldNameType string

type refundAttemptsDBFieldName struct {
	Id                RefundAttemptsDBFieldNameType
	RefundId          RefundAttemptsDBFieldNameType
	AttemptNo         RefundAttemptsDBFieldNameType
	AttemptType       RefundAttemptsDBFieldNameType
	ProviderAccountId RefundAttemptsDBFieldNameType
	Amount            RefundAttemptsDBFieldNameType
	CurrencyCode      RefundAttemptsDBFieldNameType
	AttemptStatus     RefundAttemptsDBFieldNameType
	ProviderRefundRef RefundAttemptsDBFieldNameType
	FailureCode       RefundAttemptsDBFieldNameType
	FailureReason     RefundAttemptsDBFieldNameType
	RawRequest        RefundAttemptsDBFieldNameType
	RawResponse       RefundAttemptsDBFieldNameType
	Metadata          RefundAttemptsDBFieldNameType
	MetaCreatedAt     RefundAttemptsDBFieldNameType
	MetaCreatedBy     RefundAttemptsDBFieldNameType
	MetaUpdatedAt     RefundAttemptsDBFieldNameType
	MetaUpdatedBy     RefundAttemptsDBFieldNameType
	MetaDeletedAt     RefundAttemptsDBFieldNameType
	MetaDeletedBy     RefundAttemptsDBFieldNameType
}

var RefundAttemptsDBFieldName = refundAttemptsDBFieldName{
	Id:                "id",
	RefundId:          "refund_id",
	AttemptNo:         "attempt_no",
	AttemptType:       "attempt_type",
	ProviderAccountId: "provider_account_id",
	Amount:            "amount",
	CurrencyCode:      "currency_code",
	AttemptStatus:     "attempt_status",
	ProviderRefundRef: "provider_refund_ref",
	FailureCode:       "failure_code",
	FailureReason:     "failure_reason",
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

func NewRefundAttemptsDBFieldNameFromStr(field string) (dbField RefundAttemptsDBFieldNameType, found bool) {
	switch field {

	case string(RefundAttemptsDBFieldName.Id):
		return RefundAttemptsDBFieldName.Id, true

	case string(RefundAttemptsDBFieldName.RefundId):
		return RefundAttemptsDBFieldName.RefundId, true

	case string(RefundAttemptsDBFieldName.AttemptNo):
		return RefundAttemptsDBFieldName.AttemptNo, true

	case string(RefundAttemptsDBFieldName.AttemptType):
		return RefundAttemptsDBFieldName.AttemptType, true

	case string(RefundAttemptsDBFieldName.ProviderAccountId):
		return RefundAttemptsDBFieldName.ProviderAccountId, true

	case string(RefundAttemptsDBFieldName.Amount):
		return RefundAttemptsDBFieldName.Amount, true

	case string(RefundAttemptsDBFieldName.CurrencyCode):
		return RefundAttemptsDBFieldName.CurrencyCode, true

	case string(RefundAttemptsDBFieldName.AttemptStatus):
		return RefundAttemptsDBFieldName.AttemptStatus, true

	case string(RefundAttemptsDBFieldName.ProviderRefundRef):
		return RefundAttemptsDBFieldName.ProviderRefundRef, true

	case string(RefundAttemptsDBFieldName.FailureCode):
		return RefundAttemptsDBFieldName.FailureCode, true

	case string(RefundAttemptsDBFieldName.FailureReason):
		return RefundAttemptsDBFieldName.FailureReason, true

	case string(RefundAttemptsDBFieldName.RawRequest):
		return RefundAttemptsDBFieldName.RawRequest, true

	case string(RefundAttemptsDBFieldName.RawResponse):
		return RefundAttemptsDBFieldName.RawResponse, true

	case string(RefundAttemptsDBFieldName.Metadata):
		return RefundAttemptsDBFieldName.Metadata, true

	case string(RefundAttemptsDBFieldName.MetaCreatedAt):
		return RefundAttemptsDBFieldName.MetaCreatedAt, true

	case string(RefundAttemptsDBFieldName.MetaCreatedBy):
		return RefundAttemptsDBFieldName.MetaCreatedBy, true

	case string(RefundAttemptsDBFieldName.MetaUpdatedAt):
		return RefundAttemptsDBFieldName.MetaUpdatedAt, true

	case string(RefundAttemptsDBFieldName.MetaUpdatedBy):
		return RefundAttemptsDBFieldName.MetaUpdatedBy, true

	case string(RefundAttemptsDBFieldName.MetaDeletedAt):
		return RefundAttemptsDBFieldName.MetaDeletedAt, true

	case string(RefundAttemptsDBFieldName.MetaDeletedBy):
		return RefundAttemptsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var RefundAttemptsFilterJoins = map[string]JoinSpec{}

var RefundAttemptsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refund_id": {
		SourcePath:        "refund_id",
		DefaultOutputPath: "refundId",
		Column:            "refund_id",
		SQLAlias:          "refund_id",
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
	"attempt_type": {
		SourcePath:        "attempt_type",
		DefaultOutputPath: "attemptType",
		Column:            "attempt_type",
		SQLAlias:          "attempt_type",
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
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"attempt_status": {
		SourcePath:        "attempt_status",
		DefaultOutputPath: "attemptStatus",
		Column:            "attempt_status",
		SQLAlias:          "attempt_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_refund_ref": {
		SourcePath:        "provider_refund_ref",
		DefaultOutputPath: "providerRefundRef",
		Column:            "provider_refund_ref",
		SQLAlias:          "provider_refund_ref",
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
	"failure_reason": {
		SourcePath:        "failure_reason",
		DefaultOutputPath: "failureReason",
		Column:            "failure_reason",
		SQLAlias:          "failure_reason",
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

func NewRefundAttemptsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = RefundAttemptsFilterFields[field]
	return
}

type RefundAttemptsFilterResult struct {
	RefundAttempts
	FilterCount int `db:"count"`
}

func ValidateRefundAttemptsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewRefundAttemptsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewRefundAttemptsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewRefundAttemptsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateRefundAttemptsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateRefundAttemptsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewRefundAttemptsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateRefundAttemptsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type RefundAttemptsAttemptStatus string

const (
	RefundAttemptsAttemptStatusCreated    RefundAttemptsAttemptStatus = "created"
	RefundAttemptsAttemptStatusProcessing RefundAttemptsAttemptStatus = "processing"
	RefundAttemptsAttemptStatusSucceeded  RefundAttemptsAttemptStatus = "succeeded"
	RefundAttemptsAttemptStatusFailed     RefundAttemptsAttemptStatus = "failed"
	RefundAttemptsAttemptStatusCancelled  RefundAttemptsAttemptStatus = "cancelled"
)

type RefundAttemptsAttemptType string

const (
	RefundAttemptsAttemptTypeProvider     RefundAttemptsAttemptType = "provider"
	RefundAttemptsAttemptTypeManual       RefundAttemptsAttemptType = "manual"
	RefundAttemptsAttemptTypeWalletCredit RefundAttemptsAttemptType = "wallet_credit"
	RefundAttemptsAttemptTypeBankTransfer RefundAttemptsAttemptType = "bank_transfer"
)

type RefundAttempts struct {
	Id                uuid.UUID                   `db:"id"`
	RefundId          uuid.UUID                   `db:"refund_id"`
	AttemptNo         int                         `db:"attempt_no"`
	AttemptType       RefundAttemptsAttemptType   `db:"attempt_type"`
	ProviderAccountId nuuid.NUUID                 `db:"provider_account_id"`
	Amount            decimal.Decimal             `db:"amount"`
	CurrencyCode      string                      `db:"currency_code"`
	AttemptStatus     RefundAttemptsAttemptStatus `db:"attempt_status"`
	ProviderRefundRef null.String                 `db:"provider_refund_ref"`
	FailureCode       null.String                 `db:"failure_code"`
	FailureReason     null.String                 `db:"failure_reason"`
	RawRequest        json.RawMessage             `db:"raw_request"`
	RawResponse       json.RawMessage             `db:"raw_response"`
	Metadata          json.RawMessage             `db:"metadata"`

	shared.MetaSignature
}
type RefundAttemptsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RefundAttempts) ToRefundAttemptsPrimaryID() RefundAttemptsPrimaryID {
	return RefundAttemptsPrimaryID{
		Id: d.Id,
	}
}

type RefundAttemptsList []*RefundAttempts

type RefundAttemptsFilterResultList []*RefundAttemptsFilterResult
