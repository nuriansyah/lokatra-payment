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

type PayoutAttemptsDBFieldNameType string

type payoutAttemptsDBFieldName struct {
	Id                PayoutAttemptsDBFieldNameType
	PayoutId          PayoutAttemptsDBFieldNameType
	AttemptNo         PayoutAttemptsDBFieldNameType
	AttemptType       PayoutAttemptsDBFieldNameType
	ProviderAccountId PayoutAttemptsDBFieldNameType
	Amount            PayoutAttemptsDBFieldNameType
	CurrencyCode      PayoutAttemptsDBFieldNameType
	AttemptStatus     PayoutAttemptsDBFieldNameType
	ProviderPayoutRef PayoutAttemptsDBFieldNameType
	FailureCode       PayoutAttemptsDBFieldNameType
	FailureReason     PayoutAttemptsDBFieldNameType
	RawRequest        PayoutAttemptsDBFieldNameType
	RawResponse       PayoutAttemptsDBFieldNameType
	Metadata          PayoutAttemptsDBFieldNameType
	MetaCreatedAt     PayoutAttemptsDBFieldNameType
	MetaCreatedBy     PayoutAttemptsDBFieldNameType
	MetaUpdatedAt     PayoutAttemptsDBFieldNameType
	MetaUpdatedBy     PayoutAttemptsDBFieldNameType
	MetaDeletedAt     PayoutAttemptsDBFieldNameType
	MetaDeletedBy     PayoutAttemptsDBFieldNameType
}

var PayoutAttemptsDBFieldName = payoutAttemptsDBFieldName{
	Id:                "id",
	PayoutId:          "payout_id",
	AttemptNo:         "attempt_no",
	AttemptType:       "attempt_type",
	ProviderAccountId: "provider_account_id",
	Amount:            "amount",
	CurrencyCode:      "currency_code",
	AttemptStatus:     "attempt_status",
	ProviderPayoutRef: "provider_payout_ref",
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

func NewPayoutAttemptsDBFieldNameFromStr(field string) (dbField PayoutAttemptsDBFieldNameType, found bool) {
	switch field {

	case string(PayoutAttemptsDBFieldName.Id):
		return PayoutAttemptsDBFieldName.Id, true

	case string(PayoutAttemptsDBFieldName.PayoutId):
		return PayoutAttemptsDBFieldName.PayoutId, true

	case string(PayoutAttemptsDBFieldName.AttemptNo):
		return PayoutAttemptsDBFieldName.AttemptNo, true

	case string(PayoutAttemptsDBFieldName.AttemptType):
		return PayoutAttemptsDBFieldName.AttemptType, true

	case string(PayoutAttemptsDBFieldName.ProviderAccountId):
		return PayoutAttemptsDBFieldName.ProviderAccountId, true

	case string(PayoutAttemptsDBFieldName.Amount):
		return PayoutAttemptsDBFieldName.Amount, true

	case string(PayoutAttemptsDBFieldName.CurrencyCode):
		return PayoutAttemptsDBFieldName.CurrencyCode, true

	case string(PayoutAttemptsDBFieldName.AttemptStatus):
		return PayoutAttemptsDBFieldName.AttemptStatus, true

	case string(PayoutAttemptsDBFieldName.ProviderPayoutRef):
		return PayoutAttemptsDBFieldName.ProviderPayoutRef, true

	case string(PayoutAttemptsDBFieldName.FailureCode):
		return PayoutAttemptsDBFieldName.FailureCode, true

	case string(PayoutAttemptsDBFieldName.FailureReason):
		return PayoutAttemptsDBFieldName.FailureReason, true

	case string(PayoutAttemptsDBFieldName.RawRequest):
		return PayoutAttemptsDBFieldName.RawRequest, true

	case string(PayoutAttemptsDBFieldName.RawResponse):
		return PayoutAttemptsDBFieldName.RawResponse, true

	case string(PayoutAttemptsDBFieldName.Metadata):
		return PayoutAttemptsDBFieldName.Metadata, true

	case string(PayoutAttemptsDBFieldName.MetaCreatedAt):
		return PayoutAttemptsDBFieldName.MetaCreatedAt, true

	case string(PayoutAttemptsDBFieldName.MetaCreatedBy):
		return PayoutAttemptsDBFieldName.MetaCreatedBy, true

	case string(PayoutAttemptsDBFieldName.MetaUpdatedAt):
		return PayoutAttemptsDBFieldName.MetaUpdatedAt, true

	case string(PayoutAttemptsDBFieldName.MetaUpdatedBy):
		return PayoutAttemptsDBFieldName.MetaUpdatedBy, true

	case string(PayoutAttemptsDBFieldName.MetaDeletedAt):
		return PayoutAttemptsDBFieldName.MetaDeletedAt, true

	case string(PayoutAttemptsDBFieldName.MetaDeletedBy):
		return PayoutAttemptsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PayoutAttemptsFilterJoins = map[string]JoinSpec{}

var PayoutAttemptsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_id": {
		SourcePath:        "payout_id",
		DefaultOutputPath: "payoutId",
		Column:            "payout_id",
		SQLAlias:          "payout_id",
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
	"provider_payout_ref": {
		SourcePath:        "provider_payout_ref",
		DefaultOutputPath: "providerPayoutRef",
		Column:            "provider_payout_ref",
		SQLAlias:          "provider_payout_ref",
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

func NewPayoutAttemptsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayoutAttemptsFilterFields[field]
	return
}

type PayoutAttemptsFilterResult struct {
	PayoutAttempts
	FilterCount int `db:"count"`
}

func ValidatePayoutAttemptsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPayoutAttemptsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayoutAttemptsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPayoutAttemptsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePayoutAttemptsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePayoutAttemptsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPayoutAttemptsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePayoutAttemptsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PayoutAttemptsAttemptStatus string

const (
	PayoutAttemptsAttemptStatusCreated    PayoutAttemptsAttemptStatus = "created"
	PayoutAttemptsAttemptStatusProcessing PayoutAttemptsAttemptStatus = "processing"
	PayoutAttemptsAttemptStatusSucceeded  PayoutAttemptsAttemptStatus = "succeeded"
	PayoutAttemptsAttemptStatusFailed     PayoutAttemptsAttemptStatus = "failed"
	PayoutAttemptsAttemptStatusCancelled  PayoutAttemptsAttemptStatus = "cancelled"
)

type PayoutAttemptsAttemptType string

const (
	PayoutAttemptsAttemptTypeProvider     PayoutAttemptsAttemptType = "provider"
	PayoutAttemptsAttemptTypeManual       PayoutAttemptsAttemptType = "manual"
	PayoutAttemptsAttemptTypeBankTransfer PayoutAttemptsAttemptType = "bank_transfer"
)

type PayoutAttempts struct {
	Id                uuid.UUID                   `db:"id"`
	PayoutId          uuid.UUID                   `db:"payout_id"`
	AttemptNo         int                         `db:"attempt_no"`
	AttemptType       PayoutAttemptsAttemptType   `db:"attempt_type"`
	ProviderAccountId nuuid.NUUID                 `db:"provider_account_id"`
	Amount            decimal.Decimal             `db:"amount"`
	CurrencyCode      string                      `db:"currency_code"`
	AttemptStatus     PayoutAttemptsAttemptStatus `db:"attempt_status"`
	ProviderPayoutRef null.String                 `db:"provider_payout_ref"`
	FailureCode       null.String                 `db:"failure_code"`
	FailureReason     null.String                 `db:"failure_reason"`
	RawRequest        json.RawMessage             `db:"raw_request"`
	RawResponse       json.RawMessage             `db:"raw_response"`
	Metadata          json.RawMessage             `db:"metadata"`

	shared.MetaSignature
}
type PayoutAttemptsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayoutAttempts) ToPayoutAttemptsPrimaryID() PayoutAttemptsPrimaryID {
	return PayoutAttemptsPrimaryID{
		Id: d.Id,
	}
}

type PayoutAttemptsList []*PayoutAttempts

type PayoutAttemptsFilterResultList []*PayoutAttemptsFilterResult
