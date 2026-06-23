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

type FinanceMetricSnapshotsDBFieldNameType string

type financeMetricSnapshotsDBFieldName struct {
	Id            FinanceMetricSnapshotsDBFieldNameType
	MetricName    FinanceMetricSnapshotsDBFieldNameType
	MetricScope   FinanceMetricSnapshotsDBFieldNameType
	ScopeRef      FinanceMetricSnapshotsDBFieldNameType
	PeriodStart   FinanceMetricSnapshotsDBFieldNameType
	PeriodEnd     FinanceMetricSnapshotsDBFieldNameType
	MetricValue   FinanceMetricSnapshotsDBFieldNameType
	MetricUnit    FinanceMetricSnapshotsDBFieldNameType
	Dimensions    FinanceMetricSnapshotsDBFieldNameType
	Metadata      FinanceMetricSnapshotsDBFieldNameType
	MetaCreatedAt FinanceMetricSnapshotsDBFieldNameType
	MetaCreatedBy FinanceMetricSnapshotsDBFieldNameType
	MetaUpdatedAt FinanceMetricSnapshotsDBFieldNameType
	MetaUpdatedBy FinanceMetricSnapshotsDBFieldNameType
	MetaDeletedAt FinanceMetricSnapshotsDBFieldNameType
	MetaDeletedBy FinanceMetricSnapshotsDBFieldNameType
}

var FinanceMetricSnapshotsDBFieldName = financeMetricSnapshotsDBFieldName{
	Id:            "id",
	MetricName:    "metric_name",
	MetricScope:   "metric_scope",
	ScopeRef:      "scope_ref",
	PeriodStart:   "period_start",
	PeriodEnd:     "period_end",
	MetricValue:   "metric_value",
	MetricUnit:    "metric_unit",
	Dimensions:    "dimensions",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewFinanceMetricSnapshotsDBFieldNameFromStr(field string) (dbField FinanceMetricSnapshotsDBFieldNameType, found bool) {
	switch field {

	case string(FinanceMetricSnapshotsDBFieldName.Id):
		return FinanceMetricSnapshotsDBFieldName.Id, true

	case string(FinanceMetricSnapshotsDBFieldName.MetricName):
		return FinanceMetricSnapshotsDBFieldName.MetricName, true

	case string(FinanceMetricSnapshotsDBFieldName.MetricScope):
		return FinanceMetricSnapshotsDBFieldName.MetricScope, true

	case string(FinanceMetricSnapshotsDBFieldName.ScopeRef):
		return FinanceMetricSnapshotsDBFieldName.ScopeRef, true

	case string(FinanceMetricSnapshotsDBFieldName.PeriodStart):
		return FinanceMetricSnapshotsDBFieldName.PeriodStart, true

	case string(FinanceMetricSnapshotsDBFieldName.PeriodEnd):
		return FinanceMetricSnapshotsDBFieldName.PeriodEnd, true

	case string(FinanceMetricSnapshotsDBFieldName.MetricValue):
		return FinanceMetricSnapshotsDBFieldName.MetricValue, true

	case string(FinanceMetricSnapshotsDBFieldName.MetricUnit):
		return FinanceMetricSnapshotsDBFieldName.MetricUnit, true

	case string(FinanceMetricSnapshotsDBFieldName.Dimensions):
		return FinanceMetricSnapshotsDBFieldName.Dimensions, true

	case string(FinanceMetricSnapshotsDBFieldName.Metadata):
		return FinanceMetricSnapshotsDBFieldName.Metadata, true

	case string(FinanceMetricSnapshotsDBFieldName.MetaCreatedAt):
		return FinanceMetricSnapshotsDBFieldName.MetaCreatedAt, true

	case string(FinanceMetricSnapshotsDBFieldName.MetaCreatedBy):
		return FinanceMetricSnapshotsDBFieldName.MetaCreatedBy, true

	case string(FinanceMetricSnapshotsDBFieldName.MetaUpdatedAt):
		return FinanceMetricSnapshotsDBFieldName.MetaUpdatedAt, true

	case string(FinanceMetricSnapshotsDBFieldName.MetaUpdatedBy):
		return FinanceMetricSnapshotsDBFieldName.MetaUpdatedBy, true

	case string(FinanceMetricSnapshotsDBFieldName.MetaDeletedAt):
		return FinanceMetricSnapshotsDBFieldName.MetaDeletedAt, true

	case string(FinanceMetricSnapshotsDBFieldName.MetaDeletedBy):
		return FinanceMetricSnapshotsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceMetricSnapshotsFilterJoins = map[string]JoinSpec{}

var FinanceMetricSnapshotsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"metric_name": {
		SourcePath:        "metric_name",
		DefaultOutputPath: "metricName",
		Column:            "metric_name",
		SQLAlias:          "metric_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"metric_scope": {
		SourcePath:        "metric_scope",
		DefaultOutputPath: "metricScope",
		Column:            "metric_scope",
		SQLAlias:          "metric_scope",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"scope_ref": {
		SourcePath:        "scope_ref",
		DefaultOutputPath: "scopeRef",
		Column:            "scope_ref",
		SQLAlias:          "scope_ref",
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
	"metric_value": {
		SourcePath:        "metric_value",
		DefaultOutputPath: "metricValue",
		Column:            "metric_value",
		SQLAlias:          "metric_value",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"metric_unit": {
		SourcePath:        "metric_unit",
		DefaultOutputPath: "metricUnit",
		Column:            "metric_unit",
		SQLAlias:          "metric_unit",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"dimensions": {
		SourcePath:        "dimensions",
		DefaultOutputPath: "dimensions",
		Column:            "dimensions",
		SQLAlias:          "dimensions",
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

func NewFinanceMetricSnapshotsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceMetricSnapshotsFilterFields[field]
	return
}

type FinanceMetricSnapshotsFilterResult struct {
	FinanceMetricSnapshots
	FilterCount int `db:"count"`
}

func ValidateFinanceMetricSnapshotsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceMetricSnapshotsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceMetricSnapshotsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceMetricSnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceMetricSnapshotsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceMetricSnapshotsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceMetricSnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceMetricSnapshotsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FinanceMetricSnapshots struct {
	Id          uuid.UUID       `db:"id"`
	MetricName  string          `db:"metric_name"`
	MetricScope string          `db:"metric_scope"`
	ScopeRef    null.String     `db:"scope_ref"`
	PeriodStart time.Time       `db:"period_start"`
	PeriodEnd   time.Time       `db:"period_end"`
	MetricValue decimal.Decimal `db:"metric_value"`
	MetricUnit  string          `db:"metric_unit"`
	Dimensions  json.RawMessage `db:"dimensions"`
	Metadata    json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceMetricSnapshotsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceMetricSnapshots) ToFinanceMetricSnapshotsPrimaryID() FinanceMetricSnapshotsPrimaryID {
	return FinanceMetricSnapshotsPrimaryID{
		Id: d.Id,
	}
}

type FinanceMetricSnapshotsList []*FinanceMetricSnapshots

type FinanceMetricSnapshotsFilterResultList []*FinanceMetricSnapshotsFilterResult
