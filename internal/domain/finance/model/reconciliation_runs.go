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

type ReconciliationRunsDBFieldNameType string

type reconciliationRunsDBFieldName struct {
	Id              ReconciliationRunsDBFieldNameType
	ReconCode       ReconciliationRunsDBFieldNameType
	ReconType       ReconciliationRunsDBFieldNameType
	PeriodStart     ReconciliationRunsDBFieldNameType
	PeriodEnd       ReconciliationRunsDBFieldNameType
	CurrencyCode    ReconciliationRunsDBFieldNameType
	RunStatus       ReconciliationRunsDBFieldNameType
	ToleranceAmount ReconciliationRunsDBFieldNameType
	ToleranceDays   ReconciliationRunsDBFieldNameType
	StartedAt       ReconciliationRunsDBFieldNameType
	CompletedAt     ReconciliationRunsDBFieldNameType
	Summary         ReconciliationRunsDBFieldNameType
	MetaCreatedAt   ReconciliationRunsDBFieldNameType
	MetaCreatedBy   ReconciliationRunsDBFieldNameType
	MetaUpdatedAt   ReconciliationRunsDBFieldNameType
	MetaUpdatedBy   ReconciliationRunsDBFieldNameType
	MetaDeletedAt   ReconciliationRunsDBFieldNameType
	MetaDeletedBy   ReconciliationRunsDBFieldNameType
}

var ReconciliationRunsDBFieldName = reconciliationRunsDBFieldName{
	Id:              "id",
	ReconCode:       "recon_code",
	ReconType:       "recon_type",
	PeriodStart:     "period_start",
	PeriodEnd:       "period_end",
	CurrencyCode:    "currency_code",
	RunStatus:       "run_status",
	ToleranceAmount: "tolerance_amount",
	ToleranceDays:   "tolerance_days",
	StartedAt:       "started_at",
	CompletedAt:     "completed_at",
	Summary:         "summary",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewReconciliationRunsDBFieldNameFromStr(field string) (dbField ReconciliationRunsDBFieldNameType, found bool) {
	switch field {

	case string(ReconciliationRunsDBFieldName.Id):
		return ReconciliationRunsDBFieldName.Id, true

	case string(ReconciliationRunsDBFieldName.ReconCode):
		return ReconciliationRunsDBFieldName.ReconCode, true

	case string(ReconciliationRunsDBFieldName.ReconType):
		return ReconciliationRunsDBFieldName.ReconType, true

	case string(ReconciliationRunsDBFieldName.PeriodStart):
		return ReconciliationRunsDBFieldName.PeriodStart, true

	case string(ReconciliationRunsDBFieldName.PeriodEnd):
		return ReconciliationRunsDBFieldName.PeriodEnd, true

	case string(ReconciliationRunsDBFieldName.CurrencyCode):
		return ReconciliationRunsDBFieldName.CurrencyCode, true

	case string(ReconciliationRunsDBFieldName.RunStatus):
		return ReconciliationRunsDBFieldName.RunStatus, true

	case string(ReconciliationRunsDBFieldName.ToleranceAmount):
		return ReconciliationRunsDBFieldName.ToleranceAmount, true

	case string(ReconciliationRunsDBFieldName.ToleranceDays):
		return ReconciliationRunsDBFieldName.ToleranceDays, true

	case string(ReconciliationRunsDBFieldName.StartedAt):
		return ReconciliationRunsDBFieldName.StartedAt, true

	case string(ReconciliationRunsDBFieldName.CompletedAt):
		return ReconciliationRunsDBFieldName.CompletedAt, true

	case string(ReconciliationRunsDBFieldName.Summary):
		return ReconciliationRunsDBFieldName.Summary, true

	case string(ReconciliationRunsDBFieldName.MetaCreatedAt):
		return ReconciliationRunsDBFieldName.MetaCreatedAt, true

	case string(ReconciliationRunsDBFieldName.MetaCreatedBy):
		return ReconciliationRunsDBFieldName.MetaCreatedBy, true

	case string(ReconciliationRunsDBFieldName.MetaUpdatedAt):
		return ReconciliationRunsDBFieldName.MetaUpdatedAt, true

	case string(ReconciliationRunsDBFieldName.MetaUpdatedBy):
		return ReconciliationRunsDBFieldName.MetaUpdatedBy, true

	case string(ReconciliationRunsDBFieldName.MetaDeletedAt):
		return ReconciliationRunsDBFieldName.MetaDeletedAt, true

	case string(ReconciliationRunsDBFieldName.MetaDeletedBy):
		return ReconciliationRunsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ReconciliationRunsFilterJoins = map[string]JoinSpec{}

var ReconciliationRunsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"recon_code": {
		SourcePath:        "recon_code",
		DefaultOutputPath: "reconCode",
		Column:            "recon_code",
		SQLAlias:          "recon_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"recon_type": {
		SourcePath:        "recon_type",
		DefaultOutputPath: "reconType",
		Column:            "recon_type",
		SQLAlias:          "recon_type",
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
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"run_status": {
		SourcePath:        "run_status",
		DefaultOutputPath: "runStatus",
		Column:            "run_status",
		SQLAlias:          "run_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tolerance_amount": {
		SourcePath:        "tolerance_amount",
		DefaultOutputPath: "toleranceAmount",
		Column:            "tolerance_amount",
		SQLAlias:          "tolerance_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tolerance_days": {
		SourcePath:        "tolerance_days",
		DefaultOutputPath: "toleranceDays",
		Column:            "tolerance_days",
		SQLAlias:          "tolerance_days",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"started_at": {
		SourcePath:        "started_at",
		DefaultOutputPath: "startedAt",
		Column:            "started_at",
		SQLAlias:          "started_at",
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
	"summary": {
		SourcePath:        "summary",
		DefaultOutputPath: "summary",
		Column:            "summary",
		SQLAlias:          "summary",
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

func NewReconciliationRunsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ReconciliationRunsFilterFields[field]
	return
}

type ReconciliationRunsFilterResult struct {
	ReconciliationRuns
	FilterCount int `db:"count"`
}

func ValidateReconciliationRunsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewReconciliationRunsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewReconciliationRunsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewReconciliationRunsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateReconciliationRunsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateReconciliationRunsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewReconciliationRunsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateReconciliationRunsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ReconType string

const (
	ReconTypeProviderVsLedger ReconType = "provider_vs_ledger"
	ReconTypeBankVsLedger     ReconType = "bank_vs_ledger"
	ReconTypeProviderVsBank   ReconType = "provider_vs_bank"
	ReconTypeThreeWay         ReconType = "three_way"
)

type ReconciliationRunsRunStatus string

const (
	ReconciliationRunsRunStatusRunning   ReconciliationRunsRunStatus = "running"
	ReconciliationRunsRunStatusCompleted ReconciliationRunsRunStatus = "completed"
	ReconciliationRunsRunStatusFailed    ReconciliationRunsRunStatus = "failed"
	ReconciliationRunsRunStatusCancelled ReconciliationRunsRunStatus = "cancelled"
)

type ReconciliationRuns struct {
	Id              uuid.UUID                   `db:"id"`
	ReconCode       string                      `db:"recon_code"`
	ReconType       ReconType                   `db:"recon_type"`
	PeriodStart     time.Time                   `db:"period_start"`
	PeriodEnd       time.Time                   `db:"period_end"`
	CurrencyCode    null.String                 `db:"currency_code"`
	RunStatus       ReconciliationRunsRunStatus `db:"run_status"`
	ToleranceAmount decimal.Decimal             `db:"tolerance_amount"`
	ToleranceDays   int                         `db:"tolerance_days"`
	StartedAt       time.Time                   `db:"started_at"`
	CompletedAt     null.Time                   `db:"completed_at"`
	Summary         json.RawMessage             `db:"summary"`

	shared.MetaSignature
}
type ReconciliationRunsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ReconciliationRuns) ToReconciliationRunsPrimaryID() ReconciliationRunsPrimaryID {
	return ReconciliationRunsPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationRunsList []*ReconciliationRuns

type ReconciliationRunsFilterResultList []*ReconciliationRunsFilterResult
