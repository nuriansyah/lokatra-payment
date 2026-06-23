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

type PayoutsDBFieldNameType string

type payoutsDBFieldName struct {
	Id                PayoutsDBFieldNameType
	PayoutCode        PayoutsDBFieldNameType
	PayoutBatchId     PayoutsDBFieldNameType
	SettlementBatchId PayoutsDBFieldNameType
	MerchantPartyId   PayoutsDBFieldNameType
	PayoutMethodId    PayoutsDBFieldNameType
	ProviderAccountId PayoutsDBFieldNameType
	CurrencyCode      PayoutsDBFieldNameType
	Amount            PayoutsDBFieldNameType
	FeeAmount         PayoutsDBFieldNameType
	NetSentAmount     PayoutsDBFieldNameType
	IdempotencyKey    PayoutsDBFieldNameType
	ProviderPayoutRef PayoutsDBFieldNameType
	PayoutStatus      PayoutsDBFieldNameType
	HoldReasonCode    PayoutsDBFieldNameType
	InitiatedAt       PayoutsDBFieldNameType
	CompletedAt       PayoutsDBFieldNameType
	FailedAt          PayoutsDBFieldNameType
	FailureCode       PayoutsDBFieldNameType
	FailureReason     PayoutsDBFieldNameType
	Metadata          PayoutsDBFieldNameType
	MetaCreatedAt     PayoutsDBFieldNameType
	MetaCreatedBy     PayoutsDBFieldNameType
	MetaUpdatedAt     PayoutsDBFieldNameType
	MetaUpdatedBy     PayoutsDBFieldNameType
	MetaDeletedAt     PayoutsDBFieldNameType
	MetaDeletedBy     PayoutsDBFieldNameType
}

var PayoutsDBFieldName = payoutsDBFieldName{
	Id:                "id",
	PayoutCode:        "payout_code",
	PayoutBatchId:     "payout_batch_id",
	SettlementBatchId: "settlement_batch_id",
	MerchantPartyId:   "merchant_party_id",
	PayoutMethodId:    "payout_method_id",
	ProviderAccountId: "provider_account_id",
	CurrencyCode:      "currency_code",
	Amount:            "amount",
	FeeAmount:         "fee_amount",
	NetSentAmount:     "net_sent_amount",
	IdempotencyKey:    "idempotency_key",
	ProviderPayoutRef: "provider_payout_ref",
	PayoutStatus:      "payout_status",
	HoldReasonCode:    "hold_reason_code",
	InitiatedAt:       "initiated_at",
	CompletedAt:       "completed_at",
	FailedAt:          "failed_at",
	FailureCode:       "failure_code",
	FailureReason:     "failure_reason",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewPayoutsDBFieldNameFromStr(field string) (dbField PayoutsDBFieldNameType, found bool) {
	switch field {

	case string(PayoutsDBFieldName.Id):
		return PayoutsDBFieldName.Id, true

	case string(PayoutsDBFieldName.PayoutCode):
		return PayoutsDBFieldName.PayoutCode, true

	case string(PayoutsDBFieldName.PayoutBatchId):
		return PayoutsDBFieldName.PayoutBatchId, true

	case string(PayoutsDBFieldName.SettlementBatchId):
		return PayoutsDBFieldName.SettlementBatchId, true

	case string(PayoutsDBFieldName.MerchantPartyId):
		return PayoutsDBFieldName.MerchantPartyId, true

	case string(PayoutsDBFieldName.PayoutMethodId):
		return PayoutsDBFieldName.PayoutMethodId, true

	case string(PayoutsDBFieldName.ProviderAccountId):
		return PayoutsDBFieldName.ProviderAccountId, true

	case string(PayoutsDBFieldName.CurrencyCode):
		return PayoutsDBFieldName.CurrencyCode, true

	case string(PayoutsDBFieldName.Amount):
		return PayoutsDBFieldName.Amount, true

	case string(PayoutsDBFieldName.FeeAmount):
		return PayoutsDBFieldName.FeeAmount, true

	case string(PayoutsDBFieldName.NetSentAmount):
		return PayoutsDBFieldName.NetSentAmount, true

	case string(PayoutsDBFieldName.IdempotencyKey):
		return PayoutsDBFieldName.IdempotencyKey, true

	case string(PayoutsDBFieldName.ProviderPayoutRef):
		return PayoutsDBFieldName.ProviderPayoutRef, true

	case string(PayoutsDBFieldName.PayoutStatus):
		return PayoutsDBFieldName.PayoutStatus, true

	case string(PayoutsDBFieldName.HoldReasonCode):
		return PayoutsDBFieldName.HoldReasonCode, true

	case string(PayoutsDBFieldName.InitiatedAt):
		return PayoutsDBFieldName.InitiatedAt, true

	case string(PayoutsDBFieldName.CompletedAt):
		return PayoutsDBFieldName.CompletedAt, true

	case string(PayoutsDBFieldName.FailedAt):
		return PayoutsDBFieldName.FailedAt, true

	case string(PayoutsDBFieldName.FailureCode):
		return PayoutsDBFieldName.FailureCode, true

	case string(PayoutsDBFieldName.FailureReason):
		return PayoutsDBFieldName.FailureReason, true

	case string(PayoutsDBFieldName.Metadata):
		return PayoutsDBFieldName.Metadata, true

	case string(PayoutsDBFieldName.MetaCreatedAt):
		return PayoutsDBFieldName.MetaCreatedAt, true

	case string(PayoutsDBFieldName.MetaCreatedBy):
		return PayoutsDBFieldName.MetaCreatedBy, true

	case string(PayoutsDBFieldName.MetaUpdatedAt):
		return PayoutsDBFieldName.MetaUpdatedAt, true

	case string(PayoutsDBFieldName.MetaUpdatedBy):
		return PayoutsDBFieldName.MetaUpdatedBy, true

	case string(PayoutsDBFieldName.MetaDeletedAt):
		return PayoutsDBFieldName.MetaDeletedAt, true

	case string(PayoutsDBFieldName.MetaDeletedBy):
		return PayoutsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PayoutsFilterJoins = map[string]JoinSpec{}

var PayoutsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_code": {
		SourcePath:        "payout_code",
		DefaultOutputPath: "payoutCode",
		Column:            "payout_code",
		SQLAlias:          "payout_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_batch_id": {
		SourcePath:        "payout_batch_id",
		DefaultOutputPath: "payoutBatchId",
		Column:            "payout_batch_id",
		SQLAlias:          "payout_batch_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"settlement_batch_id": {
		SourcePath:        "settlement_batch_id",
		DefaultOutputPath: "settlementBatchId",
		Column:            "settlement_batch_id",
		SQLAlias:          "settlement_batch_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_party_id": {
		SourcePath:        "merchant_party_id",
		DefaultOutputPath: "merchantPartyId",
		Column:            "merchant_party_id",
		SQLAlias:          "merchant_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_method_id": {
		SourcePath:        "payout_method_id",
		DefaultOutputPath: "payoutMethodId",
		Column:            "payout_method_id",
		SQLAlias:          "payout_method_id",
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
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
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
	"fee_amount": {
		SourcePath:        "fee_amount",
		DefaultOutputPath: "feeAmount",
		Column:            "fee_amount",
		SQLAlias:          "fee_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"net_sent_amount": {
		SourcePath:        "net_sent_amount",
		DefaultOutputPath: "netSentAmount",
		Column:            "net_sent_amount",
		SQLAlias:          "net_sent_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"idempotency_key": {
		SourcePath:        "idempotency_key",
		DefaultOutputPath: "idempotencyKey",
		Column:            "idempotency_key",
		SQLAlias:          "idempotency_key",
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
	"payout_status": {
		SourcePath:        "payout_status",
		DefaultOutputPath: "payoutStatus",
		Column:            "payout_status",
		SQLAlias:          "payout_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"hold_reason_code": {
		SourcePath:        "hold_reason_code",
		DefaultOutputPath: "holdReasonCode",
		Column:            "hold_reason_code",
		SQLAlias:          "hold_reason_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"initiated_at": {
		SourcePath:        "initiated_at",
		DefaultOutputPath: "initiatedAt",
		Column:            "initiated_at",
		SQLAlias:          "initiated_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"completed_at": {
		SourcePath:        "completed_at",
		DefaultOutputPath: "completedAt",
		Column:            "completed_at",
		SQLAlias:          "completed_at",
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
	"failure_reason": {
		SourcePath:        "failure_reason",
		DefaultOutputPath: "failureReason",
		Column:            "failure_reason",
		SQLAlias:          "failure_reason",
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

func NewPayoutsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayoutsFilterFields[field]
	return
}

type PayoutsFilterResult struct {
	Payouts
	FilterCount int `db:"count"`
}

func ValidatePayoutsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPayoutsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayoutsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPayoutsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePayoutsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePayoutsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPayoutsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePayoutsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PayoutStatus string

const (
	PayoutStatusCreated    PayoutStatus = "created"
	PayoutStatusQueued     PayoutStatus = "queued"
	PayoutStatusProcessing PayoutStatus = "processing"
	PayoutStatusSucceeded  PayoutStatus = "succeeded"
	PayoutStatusFailed     PayoutStatus = "failed"
	PayoutStatusReversed   PayoutStatus = "reversed"
	PayoutStatusCancelled  PayoutStatus = "cancelled"
	PayoutStatusHeld       PayoutStatus = "held"
)

type Payouts struct {
	Id                uuid.UUID       `db:"id"`
	PayoutCode        string          `db:"payout_code"`
	PayoutBatchId     nuuid.NUUID     `db:"payout_batch_id"`
	SettlementBatchId nuuid.NUUID     `db:"settlement_batch_id"`
	MerchantPartyId   uuid.UUID       `db:"merchant_party_id"`
	PayoutMethodId    uuid.UUID       `db:"payout_method_id"`
	ProviderAccountId nuuid.NUUID     `db:"provider_account_id"`
	CurrencyCode      string          `db:"currency_code"`
	Amount            decimal.Decimal `db:"amount"`
	FeeAmount         decimal.Decimal `db:"fee_amount"`
	NetSentAmount     decimal.Decimal `db:"net_sent_amount"`
	IdempotencyKey    string          `db:"idempotency_key"`
	ProviderPayoutRef null.String     `db:"provider_payout_ref"`
	PayoutStatus      PayoutStatus    `db:"payout_status"`
	HoldReasonCode    null.String     `db:"hold_reason_code"`
	InitiatedAt       null.Time       `db:"initiated_at"`
	CompletedAt       null.Time       `db:"completed_at"`
	FailedAt          null.Time       `db:"failed_at"`
	FailureCode       null.String     `db:"failure_code"`
	FailureReason     null.String     `db:"failure_reason"`
	Metadata          json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type PayoutsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d Payouts) ToPayoutsPrimaryID() PayoutsPrimaryID {
	return PayoutsPrimaryID{
		Id: d.Id,
	}
}

type PayoutsList []*Payouts

type PayoutsFilterResultList []*PayoutsFilterResult
