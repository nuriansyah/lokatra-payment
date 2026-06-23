package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type PayoutBatchesDBFieldNameType string

type payoutBatchesDBFieldName struct {
	Id                PayoutBatchesDBFieldNameType
	PayoutBatchCode   PayoutBatchesDBFieldNameType
	ProviderAccountId PayoutBatchesDBFieldNameType
	ScheduledFor      PayoutBatchesDBFieldNameType
	CurrencyCode      PayoutBatchesDBFieldNameType
	BatchStatus       PayoutBatchesDBFieldNameType
	TotalCount        PayoutBatchesDBFieldNameType
	TotalAmount       PayoutBatchesDBFieldNameType
	Metadata          PayoutBatchesDBFieldNameType
	MetaCreatedAt     PayoutBatchesDBFieldNameType
	MetaCreatedBy     PayoutBatchesDBFieldNameType
	MetaUpdatedAt     PayoutBatchesDBFieldNameType
	MetaUpdatedBy     PayoutBatchesDBFieldNameType
	MetaDeletedAt     PayoutBatchesDBFieldNameType
	MetaDeletedBy     PayoutBatchesDBFieldNameType
}

var PayoutBatchesDBFieldName = payoutBatchesDBFieldName{
	Id:                "id",
	PayoutBatchCode:   "payout_batch_code",
	ProviderAccountId: "provider_account_id",
	ScheduledFor:      "scheduled_for",
	CurrencyCode:      "currency_code",
	BatchStatus:       "batch_status",
	TotalCount:        "total_count",
	TotalAmount:       "total_amount",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewPayoutBatchesDBFieldNameFromStr(field string) (dbField PayoutBatchesDBFieldNameType, found bool) {
	switch field {

	case string(PayoutBatchesDBFieldName.Id):
		return PayoutBatchesDBFieldName.Id, true

	case string(PayoutBatchesDBFieldName.PayoutBatchCode):
		return PayoutBatchesDBFieldName.PayoutBatchCode, true

	case string(PayoutBatchesDBFieldName.ProviderAccountId):
		return PayoutBatchesDBFieldName.ProviderAccountId, true

	case string(PayoutBatchesDBFieldName.ScheduledFor):
		return PayoutBatchesDBFieldName.ScheduledFor, true

	case string(PayoutBatchesDBFieldName.CurrencyCode):
		return PayoutBatchesDBFieldName.CurrencyCode, true

	case string(PayoutBatchesDBFieldName.BatchStatus):
		return PayoutBatchesDBFieldName.BatchStatus, true

	case string(PayoutBatchesDBFieldName.TotalCount):
		return PayoutBatchesDBFieldName.TotalCount, true

	case string(PayoutBatchesDBFieldName.TotalAmount):
		return PayoutBatchesDBFieldName.TotalAmount, true

	case string(PayoutBatchesDBFieldName.Metadata):
		return PayoutBatchesDBFieldName.Metadata, true

	case string(PayoutBatchesDBFieldName.MetaCreatedAt):
		return PayoutBatchesDBFieldName.MetaCreatedAt, true

	case string(PayoutBatchesDBFieldName.MetaCreatedBy):
		return PayoutBatchesDBFieldName.MetaCreatedBy, true

	case string(PayoutBatchesDBFieldName.MetaUpdatedAt):
		return PayoutBatchesDBFieldName.MetaUpdatedAt, true

	case string(PayoutBatchesDBFieldName.MetaUpdatedBy):
		return PayoutBatchesDBFieldName.MetaUpdatedBy, true

	case string(PayoutBatchesDBFieldName.MetaDeletedAt):
		return PayoutBatchesDBFieldName.MetaDeletedAt, true

	case string(PayoutBatchesDBFieldName.MetaDeletedBy):
		return PayoutBatchesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PayoutBatchesFilterJoins = map[string]JoinSpec{}

var PayoutBatchesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_batch_code": {
		SourcePath:        "payout_batch_code",
		DefaultOutputPath: "payoutBatchCode",
		Column:            "payout_batch_code",
		SQLAlias:          "payout_batch_code",
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
	"scheduled_for": {
		SourcePath:        "scheduled_for",
		DefaultOutputPath: "scheduledFor",
		Column:            "scheduled_for",
		SQLAlias:          "scheduled_for",
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
	"batch_status": {
		SourcePath:        "batch_status",
		DefaultOutputPath: "batchStatus",
		Column:            "batch_status",
		SQLAlias:          "batch_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"total_count": {
		SourcePath:        "total_count",
		DefaultOutputPath: "totalCount",
		Column:            "total_count",
		SQLAlias:          "total_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"total_amount": {
		SourcePath:        "total_amount",
		DefaultOutputPath: "totalAmount",
		Column:            "total_amount",
		SQLAlias:          "total_amount",
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

func NewPayoutBatchesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayoutBatchesFilterFields[field]
	return
}

type PayoutBatchesFilterResult struct {
	PayoutBatches
	FilterCount int `db:"count"`
}

func ValidatePayoutBatchesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPayoutBatchesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayoutBatchesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPayoutBatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePayoutBatchesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePayoutBatchesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPayoutBatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePayoutBatchesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PayoutBatchesBatchStatus string

const (
	PayoutBatchesBatchStatusDraft              PayoutBatchesBatchStatus = "draft"
	PayoutBatchesBatchStatusQueued             PayoutBatchesBatchStatus = "queued"
	PayoutBatchesBatchStatusProcessing         PayoutBatchesBatchStatus = "processing"
	PayoutBatchesBatchStatusPartiallySucceeded PayoutBatchesBatchStatus = "partially_succeeded"
	PayoutBatchesBatchStatusSucceeded          PayoutBatchesBatchStatus = "succeeded"
	PayoutBatchesBatchStatusFailed             PayoutBatchesBatchStatus = "failed"
	PayoutBatchesBatchStatusCancelled          PayoutBatchesBatchStatus = "cancelled"
)

type PayoutBatches struct {
	Id                uuid.UUID                `db:"id"`
	PayoutBatchCode   string                   `db:"payout_batch_code"`
	ProviderAccountId nuuid.NUUID              `db:"provider_account_id"`
	ScheduledFor      time.Time                `db:"scheduled_for"`
	CurrencyCode      string                   `db:"currency_code"`
	BatchStatus       PayoutBatchesBatchStatus `db:"batch_status"`
	TotalCount        int                      `db:"total_count"`
	TotalAmount       decimal.Decimal          `db:"total_amount"`
	Metadata          json.RawMessage          `db:"metadata"`

	shared.MetaSignature
}
type PayoutBatchesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayoutBatches) ToPayoutBatchesPrimaryID() PayoutBatchesPrimaryID {
	return PayoutBatchesPrimaryID{
		Id: d.Id,
	}
}

type PayoutBatchesList []*PayoutBatches

type PayoutBatchesFilterResultList []*PayoutBatchesFilterResult
