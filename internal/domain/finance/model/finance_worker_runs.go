package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceWorkerRunsDBFieldNameType string

type financeWorkerRunsDBFieldName struct {
	Id             FinanceWorkerRunsDBFieldNameType
	WorkerName     FinanceWorkerRunsDBFieldNameType
	RunKey         FinanceWorkerRunsDBFieldNameType
	RunStatus      FinanceWorkerRunsDBFieldNameType
	StartedAt      FinanceWorkerRunsDBFieldNameType
	FinishedAt     FinanceWorkerRunsDBFieldNameType
	ProcessedCount FinanceWorkerRunsDBFieldNameType
	FailedCount    FinanceWorkerRunsDBFieldNameType
	ErrorCode      FinanceWorkerRunsDBFieldNameType
	ErrorDetail    FinanceWorkerRunsDBFieldNameType
	Metadata       FinanceWorkerRunsDBFieldNameType
	MetaCreatedAt  FinanceWorkerRunsDBFieldNameType
	MetaCreatedBy  FinanceWorkerRunsDBFieldNameType
	MetaUpdatedAt  FinanceWorkerRunsDBFieldNameType
	MetaUpdatedBy  FinanceWorkerRunsDBFieldNameType
	MetaDeletedAt  FinanceWorkerRunsDBFieldNameType
	MetaDeletedBy  FinanceWorkerRunsDBFieldNameType
}

var FinanceWorkerRunsDBFieldName = financeWorkerRunsDBFieldName{
	Id:             "id",
	WorkerName:     "worker_name",
	RunKey:         "run_key",
	RunStatus:      "run_status",
	StartedAt:      "started_at",
	FinishedAt:     "finished_at",
	ProcessedCount: "processed_count",
	FailedCount:    "failed_count",
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

func NewFinanceWorkerRunsDBFieldNameFromStr(field string) (dbField FinanceWorkerRunsDBFieldNameType, found bool) {
	switch field {

	case string(FinanceWorkerRunsDBFieldName.Id):
		return FinanceWorkerRunsDBFieldName.Id, true

	case string(FinanceWorkerRunsDBFieldName.WorkerName):
		return FinanceWorkerRunsDBFieldName.WorkerName, true

	case string(FinanceWorkerRunsDBFieldName.RunKey):
		return FinanceWorkerRunsDBFieldName.RunKey, true

	case string(FinanceWorkerRunsDBFieldName.RunStatus):
		return FinanceWorkerRunsDBFieldName.RunStatus, true

	case string(FinanceWorkerRunsDBFieldName.StartedAt):
		return FinanceWorkerRunsDBFieldName.StartedAt, true

	case string(FinanceWorkerRunsDBFieldName.FinishedAt):
		return FinanceWorkerRunsDBFieldName.FinishedAt, true

	case string(FinanceWorkerRunsDBFieldName.ProcessedCount):
		return FinanceWorkerRunsDBFieldName.ProcessedCount, true

	case string(FinanceWorkerRunsDBFieldName.FailedCount):
		return FinanceWorkerRunsDBFieldName.FailedCount, true

	case string(FinanceWorkerRunsDBFieldName.ErrorCode):
		return FinanceWorkerRunsDBFieldName.ErrorCode, true

	case string(FinanceWorkerRunsDBFieldName.ErrorDetail):
		return FinanceWorkerRunsDBFieldName.ErrorDetail, true

	case string(FinanceWorkerRunsDBFieldName.Metadata):
		return FinanceWorkerRunsDBFieldName.Metadata, true

	case string(FinanceWorkerRunsDBFieldName.MetaCreatedAt):
		return FinanceWorkerRunsDBFieldName.MetaCreatedAt, true

	case string(FinanceWorkerRunsDBFieldName.MetaCreatedBy):
		return FinanceWorkerRunsDBFieldName.MetaCreatedBy, true

	case string(FinanceWorkerRunsDBFieldName.MetaUpdatedAt):
		return FinanceWorkerRunsDBFieldName.MetaUpdatedAt, true

	case string(FinanceWorkerRunsDBFieldName.MetaUpdatedBy):
		return FinanceWorkerRunsDBFieldName.MetaUpdatedBy, true

	case string(FinanceWorkerRunsDBFieldName.MetaDeletedAt):
		return FinanceWorkerRunsDBFieldName.MetaDeletedAt, true

	case string(FinanceWorkerRunsDBFieldName.MetaDeletedBy):
		return FinanceWorkerRunsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceWorkerRunsFilterJoins = map[string]JoinSpec{}

var FinanceWorkerRunsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"worker_name": {
		SourcePath:        "worker_name",
		DefaultOutputPath: "workerName",
		Column:            "worker_name",
		SQLAlias:          "worker_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"run_key": {
		SourcePath:        "run_key",
		DefaultOutputPath: "runKey",
		Column:            "run_key",
		SQLAlias:          "run_key",
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
	"processed_count": {
		SourcePath:        "processed_count",
		DefaultOutputPath: "processedCount",
		Column:            "processed_count",
		SQLAlias:          "processed_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"failed_count": {
		SourcePath:        "failed_count",
		DefaultOutputPath: "failedCount",
		Column:            "failed_count",
		SQLAlias:          "failed_count",
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

func NewFinanceWorkerRunsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceWorkerRunsFilterFields[field]
	return
}

type FinanceWorkerRunsFilterResult struct {
	FinanceWorkerRuns
	FilterCount int `db:"count"`
}

func ValidateFinanceWorkerRunsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceWorkerRunsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceWorkerRunsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceWorkerRunsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceWorkerRunsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceWorkerRunsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceWorkerRunsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceWorkerRunsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FinanceWorkerRunsRunStatus string

const (
	FinanceWorkerRunsRunStatusRunning   FinanceWorkerRunsRunStatus = "running"
	FinanceWorkerRunsRunStatusSucceeded FinanceWorkerRunsRunStatus = "succeeded"
	FinanceWorkerRunsRunStatusFailed    FinanceWorkerRunsRunStatus = "failed"
)

type FinanceWorkerRuns struct {
	Id             uuid.UUID                  `db:"id"`
	WorkerName     string                     `db:"worker_name"`
	RunKey         string                     `db:"run_key"`
	RunStatus      FinanceWorkerRunsRunStatus `db:"run_status"`
	StartedAt      time.Time                  `db:"started_at"`
	FinishedAt     null.Time                  `db:"finished_at"`
	ProcessedCount int                        `db:"processed_count"`
	FailedCount    int                        `db:"failed_count"`
	ErrorCode      null.String                `db:"error_code"`
	ErrorDetail    null.String                `db:"error_detail"`
	Metadata       json.RawMessage            `db:"metadata"`

	shared.MetaSignature
}
type FinanceWorkerRunsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceWorkerRuns) ToFinanceWorkerRunsPrimaryID() FinanceWorkerRunsPrimaryID {
	return FinanceWorkerRunsPrimaryID{
		Id: d.Id,
	}
}

type FinanceWorkerRunsList []*FinanceWorkerRuns

type FinanceWorkerRunsFilterResultList []*FinanceWorkerRunsFilterResult
