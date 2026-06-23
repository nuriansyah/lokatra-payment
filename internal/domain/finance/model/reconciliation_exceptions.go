package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type ReconciliationExceptionsDBFieldNameType string

type reconciliationExceptionsDBFieldName struct {
	Id                  ReconciliationExceptionsDBFieldNameType
	ReconciliationRunId ReconciliationExceptionsDBFieldNameType
	ExceptionCode       ReconciliationExceptionsDBFieldNameType
	ExceptionType       ReconciliationExceptionsDBFieldNameType
	SourceRefType       ReconciliationExceptionsDBFieldNameType
	SourceRefId         ReconciliationExceptionsDBFieldNameType
	Severity            ReconciliationExceptionsDBFieldNameType
	AmountDifference    ReconciliationExceptionsDBFieldNameType
	ExceptionStatus     ReconciliationExceptionsDBFieldNameType
	Resolution          ReconciliationExceptionsDBFieldNameType
	ResolvedAt          ReconciliationExceptionsDBFieldNameType
	Metadata            ReconciliationExceptionsDBFieldNameType
	MetaCreatedAt       ReconciliationExceptionsDBFieldNameType
	MetaCreatedBy       ReconciliationExceptionsDBFieldNameType
	MetaUpdatedAt       ReconciliationExceptionsDBFieldNameType
	MetaUpdatedBy       ReconciliationExceptionsDBFieldNameType
	MetaDeletedAt       ReconciliationExceptionsDBFieldNameType
	MetaDeletedBy       ReconciliationExceptionsDBFieldNameType
}

var ReconciliationExceptionsDBFieldName = reconciliationExceptionsDBFieldName{
	Id:                  "id",
	ReconciliationRunId: "reconciliation_run_id",
	ExceptionCode:       "exception_code",
	ExceptionType:       "exception_type",
	SourceRefType:       "source_ref_type",
	SourceRefId:         "source_ref_id",
	Severity:            "severity",
	AmountDifference:    "amount_difference",
	ExceptionStatus:     "exception_status",
	Resolution:          "resolution",
	ResolvedAt:          "resolved_at",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewReconciliationExceptionsDBFieldNameFromStr(field string) (dbField ReconciliationExceptionsDBFieldNameType, found bool) {
	switch field {

	case string(ReconciliationExceptionsDBFieldName.Id):
		return ReconciliationExceptionsDBFieldName.Id, true

	case string(ReconciliationExceptionsDBFieldName.ReconciliationRunId):
		return ReconciliationExceptionsDBFieldName.ReconciliationRunId, true

	case string(ReconciliationExceptionsDBFieldName.ExceptionCode):
		return ReconciliationExceptionsDBFieldName.ExceptionCode, true

	case string(ReconciliationExceptionsDBFieldName.ExceptionType):
		return ReconciliationExceptionsDBFieldName.ExceptionType, true

	case string(ReconciliationExceptionsDBFieldName.SourceRefType):
		return ReconciliationExceptionsDBFieldName.SourceRefType, true

	case string(ReconciliationExceptionsDBFieldName.SourceRefId):
		return ReconciliationExceptionsDBFieldName.SourceRefId, true

	case string(ReconciliationExceptionsDBFieldName.Severity):
		return ReconciliationExceptionsDBFieldName.Severity, true

	case string(ReconciliationExceptionsDBFieldName.AmountDifference):
		return ReconciliationExceptionsDBFieldName.AmountDifference, true

	case string(ReconciliationExceptionsDBFieldName.ExceptionStatus):
		return ReconciliationExceptionsDBFieldName.ExceptionStatus, true

	case string(ReconciliationExceptionsDBFieldName.Resolution):
		return ReconciliationExceptionsDBFieldName.Resolution, true

	case string(ReconciliationExceptionsDBFieldName.ResolvedAt):
		return ReconciliationExceptionsDBFieldName.ResolvedAt, true

	case string(ReconciliationExceptionsDBFieldName.Metadata):
		return ReconciliationExceptionsDBFieldName.Metadata, true

	case string(ReconciliationExceptionsDBFieldName.MetaCreatedAt):
		return ReconciliationExceptionsDBFieldName.MetaCreatedAt, true

	case string(ReconciliationExceptionsDBFieldName.MetaCreatedBy):
		return ReconciliationExceptionsDBFieldName.MetaCreatedBy, true

	case string(ReconciliationExceptionsDBFieldName.MetaUpdatedAt):
		return ReconciliationExceptionsDBFieldName.MetaUpdatedAt, true

	case string(ReconciliationExceptionsDBFieldName.MetaUpdatedBy):
		return ReconciliationExceptionsDBFieldName.MetaUpdatedBy, true

	case string(ReconciliationExceptionsDBFieldName.MetaDeletedAt):
		return ReconciliationExceptionsDBFieldName.MetaDeletedAt, true

	case string(ReconciliationExceptionsDBFieldName.MetaDeletedBy):
		return ReconciliationExceptionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ReconciliationExceptionsFilterJoins = map[string]JoinSpec{}

var ReconciliationExceptionsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reconciliation_run_id": {
		SourcePath:        "reconciliation_run_id",
		DefaultOutputPath: "reconciliationRunId",
		Column:            "reconciliation_run_id",
		SQLAlias:          "reconciliation_run_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"exception_code": {
		SourcePath:        "exception_code",
		DefaultOutputPath: "exceptionCode",
		Column:            "exception_code",
		SQLAlias:          "exception_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"exception_type": {
		SourcePath:        "exception_type",
		DefaultOutputPath: "exceptionType",
		Column:            "exception_type",
		SQLAlias:          "exception_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_ref_type": {
		SourcePath:        "source_ref_type",
		DefaultOutputPath: "sourceRefType",
		Column:            "source_ref_type",
		SQLAlias:          "source_ref_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_ref_id": {
		SourcePath:        "source_ref_id",
		DefaultOutputPath: "sourceRefId",
		Column:            "source_ref_id",
		SQLAlias:          "source_ref_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"severity": {
		SourcePath:        "severity",
		DefaultOutputPath: "severity",
		Column:            "severity",
		SQLAlias:          "severity",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount_difference": {
		SourcePath:        "amount_difference",
		DefaultOutputPath: "amountDifference",
		Column:            "amount_difference",
		SQLAlias:          "amount_difference",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"exception_status": {
		SourcePath:        "exception_status",
		DefaultOutputPath: "exceptionStatus",
		Column:            "exception_status",
		SQLAlias:          "exception_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"resolution": {
		SourcePath:        "resolution",
		DefaultOutputPath: "resolution",
		Column:            "resolution",
		SQLAlias:          "resolution",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"resolved_at": {
		SourcePath:        "resolved_at",
		DefaultOutputPath: "resolvedAt",
		Column:            "resolved_at",
		SQLAlias:          "resolved_at",
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

func NewReconciliationExceptionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ReconciliationExceptionsFilterFields[field]
	return
}

type ReconciliationExceptionsFilterResult struct {
	ReconciliationExceptions
	FilterCount int `db:"count"`
}

func ValidateReconciliationExceptionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewReconciliationExceptionsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewReconciliationExceptionsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewReconciliationExceptionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateReconciliationExceptionsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateReconciliationExceptionsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewReconciliationExceptionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateReconciliationExceptionsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ExceptionStatus string

const (
	ExceptionStatusOpen     ExceptionStatus = "open"
	ExceptionStatusAssigned ExceptionStatus = "assigned"
	ExceptionStatusResolved ExceptionStatus = "resolved"
	ExceptionStatusWaived   ExceptionStatus = "waived"
)

type ExceptionType string

const (
	ExceptionTypeMissingProvider ExceptionType = "missing_provider"
	ExceptionTypeMissingBank     ExceptionType = "missing_bank"
	ExceptionTypeMissingLedger   ExceptionType = "missing_ledger"
	ExceptionTypeAmountMismatch  ExceptionType = "amount_mismatch"
	ExceptionTypeDuplicate       ExceptionType = "duplicate"
	ExceptionTypeLateSettlement  ExceptionType = "late_settlement"
	ExceptionTypeLateWebhook     ExceptionType = "late_webhook"
	ExceptionTypeManualReview    ExceptionType = "manual_review"
)

type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

type ReconciliationExceptions struct {
	Id                  uuid.UUID           `db:"id"`
	ReconciliationRunId uuid.UUID           `db:"reconciliation_run_id"`
	ExceptionCode       string              `db:"exception_code"`
	ExceptionType       ExceptionType       `db:"exception_type"`
	SourceRefType       string              `db:"source_ref_type"`
	SourceRefId         uuid.UUID           `db:"source_ref_id"`
	Severity            Severity            `db:"severity"`
	AmountDifference    decimal.NullDecimal `db:"amount_difference"`
	ExceptionStatus     ExceptionStatus     `db:"exception_status"`
	Resolution          null.String         `db:"resolution"`
	ResolvedAt          null.Time           `db:"resolved_at"`
	Metadata            json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type ReconciliationExceptionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ReconciliationExceptions) ToReconciliationExceptionsPrimaryID() ReconciliationExceptionsPrimaryID {
	return ReconciliationExceptionsPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationExceptionsList []*ReconciliationExceptions

type ReconciliationExceptionsFilterResultList []*ReconciliationExceptionsFilterResult
