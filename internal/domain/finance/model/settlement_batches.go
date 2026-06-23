package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type SettlementBatchesDBFieldNameType string

type settlementBatchesDBFieldName struct {
	Id               SettlementBatchesDBFieldNameType
	BatchCode        SettlementBatchesDBFieldNameType
	MerchantPartyId  SettlementBatchesDBFieldNameType
	CurrencyCode     SettlementBatchesDBFieldNameType
	PeriodStart      SettlementBatchesDBFieldNameType
	PeriodEnd        SettlementBatchesDBFieldNameType
	GrossAmount      SettlementBatchesDBFieldNameType
	FeeAmount        SettlementBatchesDBFieldNameType
	TaxAmount        SettlementBatchesDBFieldNameType
	ReserveAmount    SettlementBatchesDBFieldNameType
	AdjustmentAmount SettlementBatchesDBFieldNameType
	NetAmount        SettlementBatchesDBFieldNameType
	BatchStatus      SettlementBatchesDBFieldNameType
	ApprovedAt       SettlementBatchesDBFieldNameType
	LockedAt         SettlementBatchesDBFieldNameType
	Metadata         SettlementBatchesDBFieldNameType
	MetaCreatedAt    SettlementBatchesDBFieldNameType
	MetaCreatedBy    SettlementBatchesDBFieldNameType
	MetaUpdatedAt    SettlementBatchesDBFieldNameType
	MetaUpdatedBy    SettlementBatchesDBFieldNameType
	MetaDeletedAt    SettlementBatchesDBFieldNameType
	MetaDeletedBy    SettlementBatchesDBFieldNameType
}

var SettlementBatchesDBFieldName = settlementBatchesDBFieldName{
	Id:               "id",
	BatchCode:        "batch_code",
	MerchantPartyId:  "merchant_party_id",
	CurrencyCode:     "currency_code",
	PeriodStart:      "period_start",
	PeriodEnd:        "period_end",
	GrossAmount:      "gross_amount",
	FeeAmount:        "fee_amount",
	TaxAmount:        "tax_amount",
	ReserveAmount:    "reserve_amount",
	AdjustmentAmount: "adjustment_amount",
	NetAmount:        "net_amount",
	BatchStatus:      "batch_status",
	ApprovedAt:       "approved_at",
	LockedAt:         "locked_at",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewSettlementBatchesDBFieldNameFromStr(field string) (dbField SettlementBatchesDBFieldNameType, found bool) {
	switch field {

	case string(SettlementBatchesDBFieldName.Id):
		return SettlementBatchesDBFieldName.Id, true

	case string(SettlementBatchesDBFieldName.BatchCode):
		return SettlementBatchesDBFieldName.BatchCode, true

	case string(SettlementBatchesDBFieldName.MerchantPartyId):
		return SettlementBatchesDBFieldName.MerchantPartyId, true

	case string(SettlementBatchesDBFieldName.CurrencyCode):
		return SettlementBatchesDBFieldName.CurrencyCode, true

	case string(SettlementBatchesDBFieldName.PeriodStart):
		return SettlementBatchesDBFieldName.PeriodStart, true

	case string(SettlementBatchesDBFieldName.PeriodEnd):
		return SettlementBatchesDBFieldName.PeriodEnd, true

	case string(SettlementBatchesDBFieldName.GrossAmount):
		return SettlementBatchesDBFieldName.GrossAmount, true

	case string(SettlementBatchesDBFieldName.FeeAmount):
		return SettlementBatchesDBFieldName.FeeAmount, true

	case string(SettlementBatchesDBFieldName.TaxAmount):
		return SettlementBatchesDBFieldName.TaxAmount, true

	case string(SettlementBatchesDBFieldName.ReserveAmount):
		return SettlementBatchesDBFieldName.ReserveAmount, true

	case string(SettlementBatchesDBFieldName.AdjustmentAmount):
		return SettlementBatchesDBFieldName.AdjustmentAmount, true

	case string(SettlementBatchesDBFieldName.NetAmount):
		return SettlementBatchesDBFieldName.NetAmount, true

	case string(SettlementBatchesDBFieldName.BatchStatus):
		return SettlementBatchesDBFieldName.BatchStatus, true

	case string(SettlementBatchesDBFieldName.ApprovedAt):
		return SettlementBatchesDBFieldName.ApprovedAt, true

	case string(SettlementBatchesDBFieldName.LockedAt):
		return SettlementBatchesDBFieldName.LockedAt, true

	case string(SettlementBatchesDBFieldName.Metadata):
		return SettlementBatchesDBFieldName.Metadata, true

	case string(SettlementBatchesDBFieldName.MetaCreatedAt):
		return SettlementBatchesDBFieldName.MetaCreatedAt, true

	case string(SettlementBatchesDBFieldName.MetaCreatedBy):
		return SettlementBatchesDBFieldName.MetaCreatedBy, true

	case string(SettlementBatchesDBFieldName.MetaUpdatedAt):
		return SettlementBatchesDBFieldName.MetaUpdatedAt, true

	case string(SettlementBatchesDBFieldName.MetaUpdatedBy):
		return SettlementBatchesDBFieldName.MetaUpdatedBy, true

	case string(SettlementBatchesDBFieldName.MetaDeletedAt):
		return SettlementBatchesDBFieldName.MetaDeletedAt, true

	case string(SettlementBatchesDBFieldName.MetaDeletedBy):
		return SettlementBatchesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var SettlementBatchesFilterJoins = map[string]JoinSpec{}

var SettlementBatchesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"batch_code": {
		SourcePath:        "batch_code",
		DefaultOutputPath: "batchCode",
		Column:            "batch_code",
		SQLAlias:          "batch_code",
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
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"period_start": {
		SourcePath:        "period_start",
		DefaultOutputPath: "periodStart",
		Column:            "period_start",
		SQLAlias:          "period_start",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"period_end": {
		SourcePath:        "period_end",
		DefaultOutputPath: "periodEnd",
		Column:            "period_end",
		SQLAlias:          "period_end",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"gross_amount": {
		SourcePath:        "gross_amount",
		DefaultOutputPath: "grossAmount",
		Column:            "gross_amount",
		SQLAlias:          "gross_amount",
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
	"tax_amount": {
		SourcePath:        "tax_amount",
		DefaultOutputPath: "taxAmount",
		Column:            "tax_amount",
		SQLAlias:          "tax_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reserve_amount": {
		SourcePath:        "reserve_amount",
		DefaultOutputPath: "reserveAmount",
		Column:            "reserve_amount",
		SQLAlias:          "reserve_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"adjustment_amount": {
		SourcePath:        "adjustment_amount",
		DefaultOutputPath: "adjustmentAmount",
		Column:            "adjustment_amount",
		SQLAlias:          "adjustment_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"net_amount": {
		SourcePath:        "net_amount",
		DefaultOutputPath: "netAmount",
		Column:            "net_amount",
		SQLAlias:          "net_amount",
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
	"approved_at": {
		SourcePath:        "approved_at",
		DefaultOutputPath: "approvedAt",
		Column:            "approved_at",
		SQLAlias:          "approved_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"locked_at": {
		SourcePath:        "locked_at",
		DefaultOutputPath: "lockedAt",
		Column:            "locked_at",
		SQLAlias:          "locked_at",
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

func NewSettlementBatchesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = SettlementBatchesFilterFields[field]
	return
}

type SettlementBatchesFilterResult struct {
	SettlementBatches
	FilterCount int `db:"count"`
}

func ValidateSettlementBatchesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewSettlementBatchesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewSettlementBatchesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewSettlementBatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateSettlementBatchesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateSettlementBatchesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewSettlementBatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateSettlementBatchesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type SettlementBatchesBatchStatus string

const (
	SettlementBatchesBatchStatusDraft         SettlementBatchesBatchStatus = "draft"
	SettlementBatchesBatchStatusReady         SettlementBatchesBatchStatus = "ready"
	SettlementBatchesBatchStatusFunding       SettlementBatchesBatchStatus = "funding"
	SettlementBatchesBatchStatusFunded        SettlementBatchesBatchStatus = "funded"
	SettlementBatchesBatchStatusPayouting     SettlementBatchesBatchStatus = "payouting"
	SettlementBatchesBatchStatusPaid          SettlementBatchesBatchStatus = "paid"
	SettlementBatchesBatchStatusPartiallyPaid SettlementBatchesBatchStatus = "partially_paid"
	SettlementBatchesBatchStatusFailed        SettlementBatchesBatchStatus = "failed"
	SettlementBatchesBatchStatusCancelled     SettlementBatchesBatchStatus = "cancelled"
)

type SettlementBatches struct {
	Id               uuid.UUID                    `db:"id"`
	BatchCode        string                       `db:"batch_code"`
	MerchantPartyId  uuid.UUID                    `db:"merchant_party_id"`
	CurrencyCode     string                       `db:"currency_code"`
	PeriodStart      time.Time                    `db:"period_start"`
	PeriodEnd        time.Time                    `db:"period_end"`
	GrossAmount      decimal.Decimal              `db:"gross_amount"`
	FeeAmount        decimal.Decimal              `db:"fee_amount"`
	TaxAmount        decimal.Decimal              `db:"tax_amount"`
	ReserveAmount    decimal.Decimal              `db:"reserve_amount"`
	AdjustmentAmount decimal.Decimal              `db:"adjustment_amount"`
	NetAmount        decimal.Decimal              `db:"net_amount"`
	BatchStatus      SettlementBatchesBatchStatus `db:"batch_status"`
	ApprovedAt       null.Time                    `db:"approved_at"`
	LockedAt         null.Time                    `db:"locked_at"`
	Metadata         json.RawMessage              `db:"metadata"`

	shared.MetaSignature
}
type SettlementBatchesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d SettlementBatches) ToSettlementBatchesPrimaryID() SettlementBatchesPrimaryID {
	return SettlementBatchesPrimaryID{
		Id: d.Id,
	}
}

type SettlementBatchesList []*SettlementBatches

type SettlementBatchesFilterResultList []*SettlementBatchesFilterResult
