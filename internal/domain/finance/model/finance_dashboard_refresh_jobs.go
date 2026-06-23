package model

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"time"
)

type FinanceDashboardRefreshJobsDBFieldNameType string

type financeDashboardRefreshJobsDBFieldName struct {
	Id             FinanceDashboardRefreshJobsDBFieldNameType
	JobKey         FinanceDashboardRefreshJobsDBFieldNameType
	RefreshScope   FinanceDashboardRefreshJobsDBFieldNameType
	ScopeRef       FinanceDashboardRefreshJobsDBFieldNameType
	CurrencyCode   FinanceDashboardRefreshJobsDBFieldNameType
	IdempotencyKey FinanceDashboardRefreshJobsDBFieldNameType
	RefreshStatus  FinanceDashboardRefreshJobsDBFieldNameType
	RequestedAt    FinanceDashboardRefreshJobsDBFieldNameType
	StartedAt      FinanceDashboardRefreshJobsDBFieldNameType
	FinishedAt     FinanceDashboardRefreshJobsDBFieldNameType
	ErrorCode      FinanceDashboardRefreshJobsDBFieldNameType
	ErrorDetail    FinanceDashboardRefreshJobsDBFieldNameType
	Metadata       FinanceDashboardRefreshJobsDBFieldNameType
	MetaCreatedAt  FinanceDashboardRefreshJobsDBFieldNameType
	MetaCreatedBy  FinanceDashboardRefreshJobsDBFieldNameType
	MetaUpdatedAt  FinanceDashboardRefreshJobsDBFieldNameType
	MetaUpdatedBy  FinanceDashboardRefreshJobsDBFieldNameType
	MetaDeletedAt  FinanceDashboardRefreshJobsDBFieldNameType
	MetaDeletedBy  FinanceDashboardRefreshJobsDBFieldNameType
}

var FinanceDashboardRefreshJobsDBFieldName = financeDashboardRefreshJobsDBFieldName{
	Id:             "id",
	JobKey:         "job_key",
	RefreshScope:   "refresh_scope",
	ScopeRef:       "scope_ref",
	CurrencyCode:   "currency_code",
	IdempotencyKey: "idempotency_key",
	RefreshStatus:  "refresh_status",
	RequestedAt:    "requested_at",
	StartedAt:      "started_at",
	FinishedAt:     "finished_at",
	ErrorCode:      "error_code",
	ErrorDetail:    "error_detail",
	Metadata:       "metadata",
	MetaCreatedAt:  "meta_created_at",
	MetaCreatedBy:  "meta_created_by",
	MetaUpdatedAt:  "meta_updated_at",
	MetaUpdatedBy:  "meta_updated_by",
	MetaDeletedAt:  "meta_deleted_at",
	MetaDeletedBy:  "meta_deleted_by",
}

func NewFinanceDashboardRefreshJobsDBFieldNameFromStr(field string) (dbField FinanceDashboardRefreshJobsDBFieldNameType, found bool) {
	switch field {

	case string(FinanceDashboardRefreshJobsDBFieldName.Id):
		return FinanceDashboardRefreshJobsDBFieldName.Id, true

	case string(FinanceDashboardRefreshJobsDBFieldName.JobKey):
		return FinanceDashboardRefreshJobsDBFieldName.JobKey, true

	case string(FinanceDashboardRefreshJobsDBFieldName.RefreshScope):
		return FinanceDashboardRefreshJobsDBFieldName.RefreshScope, true

	case string(FinanceDashboardRefreshJobsDBFieldName.ScopeRef):
		return FinanceDashboardRefreshJobsDBFieldName.ScopeRef, true

	case string(FinanceDashboardRefreshJobsDBFieldName.CurrencyCode):
		return FinanceDashboardRefreshJobsDBFieldName.CurrencyCode, true

	case string(FinanceDashboardRefreshJobsDBFieldName.IdempotencyKey):
		return FinanceDashboardRefreshJobsDBFieldName.IdempotencyKey, true

	case string(FinanceDashboardRefreshJobsDBFieldName.RefreshStatus):
		return FinanceDashboardRefreshJobsDBFieldName.RefreshStatus, true

	case string(FinanceDashboardRefreshJobsDBFieldName.RequestedAt):
		return FinanceDashboardRefreshJobsDBFieldName.RequestedAt, true

	case string(FinanceDashboardRefreshJobsDBFieldName.StartedAt):
		return FinanceDashboardRefreshJobsDBFieldName.StartedAt, true

	case string(FinanceDashboardRefreshJobsDBFieldName.FinishedAt):
		return FinanceDashboardRefreshJobsDBFieldName.FinishedAt, true

	case string(FinanceDashboardRefreshJobsDBFieldName.ErrorCode):
		return FinanceDashboardRefreshJobsDBFieldName.ErrorCode, true

	case string(FinanceDashboardRefreshJobsDBFieldName.ErrorDetail):
		return FinanceDashboardRefreshJobsDBFieldName.ErrorDetail, true

	case string(FinanceDashboardRefreshJobsDBFieldName.Metadata):
		return FinanceDashboardRefreshJobsDBFieldName.Metadata, true

	case string(FinanceDashboardRefreshJobsDBFieldName.MetaCreatedAt):
		return FinanceDashboardRefreshJobsDBFieldName.MetaCreatedAt, true

	case string(FinanceDashboardRefreshJobsDBFieldName.MetaCreatedBy):
		return FinanceDashboardRefreshJobsDBFieldName.MetaCreatedBy, true

	case string(FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedAt):
		return FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedAt, true

	case string(FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedBy):
		return FinanceDashboardRefreshJobsDBFieldName.MetaUpdatedBy, true

	case string(FinanceDashboardRefreshJobsDBFieldName.MetaDeletedAt):
		return FinanceDashboardRefreshJobsDBFieldName.MetaDeletedAt, true

	case string(FinanceDashboardRefreshJobsDBFieldName.MetaDeletedBy):
		return FinanceDashboardRefreshJobsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceDashboardRefreshJobsFilterJoins = map[string]JoinSpec{}

var FinanceDashboardRefreshJobsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"job_key": {
		SourcePath:        "job_key",
		DefaultOutputPath: "jobKey",
		Column:            "job_key",
		SQLAlias:          "job_key",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refresh_scope": {
		SourcePath:        "refresh_scope",
		DefaultOutputPath: "refreshScope",
		Column:            "refresh_scope",
		SQLAlias:          "refresh_scope",
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
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
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
	"refresh_status": {
		SourcePath:        "refresh_status",
		DefaultOutputPath: "refreshStatus",
		Column:            "refresh_status",
		SQLAlias:          "refresh_status",
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
	"started_at": {
		SourcePath:        "started_at",
		DefaultOutputPath: "startedAt",
		Column:            "started_at",
		SQLAlias:          "started_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"finished_at": {
		SourcePath:        "finished_at",
		DefaultOutputPath: "finishedAt",
		Column:            "finished_at",
		SQLAlias:          "finished_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"error_code": {
		SourcePath:        "error_code",
		DefaultOutputPath: "errorCode",
		Column:            "error_code",
		SQLAlias:          "error_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"error_detail": {
		SourcePath:        "error_detail",
		DefaultOutputPath: "errorDetail",
		Column:            "error_detail",
		SQLAlias:          "error_detail",
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

func NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceDashboardRefreshJobsFilterFields[field]
	return
}

type FinanceDashboardRefreshJobsFilterResult struct {
	FinanceDashboardRefreshJobs
	FilterCount int `db:"count"`
}

func ValidateFinanceDashboardRefreshJobsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceDashboardRefreshJobsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceDashboardRefreshJobsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceDashboardRefreshJobsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type RefreshStatus string

const (
	RefreshStatusPending   RefreshStatus = "pending"
	RefreshStatusRunning   RefreshStatus = "running"
	RefreshStatusSucceeded RefreshStatus = "succeeded"
	RefreshStatusFailed    RefreshStatus = "failed"
)

type FinanceDashboardRefreshJobs struct {
	Id             uuid.UUID       `db:"id"`
	JobKey         string          `db:"job_key"`
	RefreshScope   string          `db:"refresh_scope"`
	ScopeRef       nuuid.NUUID     `db:"scope_ref"`
	CurrencyCode   null.String     `db:"currency_code"`
	IdempotencyKey null.String     `db:"idempotency_key"`
	RefreshStatus  RefreshStatus   `db:"refresh_status"`
	RequestedAt    time.Time       `db:"requested_at"`
	StartedAt      null.Time       `db:"started_at"`
	FinishedAt     null.Time       `db:"finished_at"`
	ErrorCode      null.String     `db:"error_code"`
	ErrorDetail    null.String     `db:"error_detail"`
	Metadata       json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceDashboardRefreshJobsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceDashboardRefreshJobs) ToFinanceDashboardRefreshJobsPrimaryID() FinanceDashboardRefreshJobsPrimaryID {
	return FinanceDashboardRefreshJobsPrimaryID{
		Id: d.Id,
	}
}

type FinanceDashboardRefreshJobsList []*FinanceDashboardRefreshJobs

type FinanceDashboardRefreshJobsFilterResultList []*FinanceDashboardRefreshJobsFilterResult
