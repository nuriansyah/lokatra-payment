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

type AccountingPeriodsDBFieldNameType string

type accountingPeriodsDBFieldName struct {
	Id            AccountingPeriodsDBFieldNameType
	BookId        AccountingPeriodsDBFieldNameType
	PeriodCode    AccountingPeriodsDBFieldNameType
	PeriodType    AccountingPeriodsDBFieldNameType
	PeriodStart   AccountingPeriodsDBFieldNameType
	PeriodEnd     AccountingPeriodsDBFieldNameType
	PeriodStatus  AccountingPeriodsDBFieldNameType
	ClosedAt      AccountingPeriodsDBFieldNameType
	ClosedBy      AccountingPeriodsDBFieldNameType
	LockReason    AccountingPeriodsDBFieldNameType
	ClosingHash   AccountingPeriodsDBFieldNameType
	Metadata      AccountingPeriodsDBFieldNameType
	MetaCreatedAt AccountingPeriodsDBFieldNameType
	MetaCreatedBy AccountingPeriodsDBFieldNameType
	MetaUpdatedAt AccountingPeriodsDBFieldNameType
	MetaUpdatedBy AccountingPeriodsDBFieldNameType
	MetaDeletedAt AccountingPeriodsDBFieldNameType
	MetaDeletedBy AccountingPeriodsDBFieldNameType
}

var AccountingPeriodsDBFieldName = accountingPeriodsDBFieldName{
	Id:            "id",
	BookId:        "book_id",
	PeriodCode:    "period_code",
	PeriodType:    "period_type",
	PeriodStart:   "period_start",
	PeriodEnd:     "period_end",
	PeriodStatus:  "period_status",
	ClosedAt:      "closed_at",
	ClosedBy:      "closed_by",
	LockReason:    "lock_reason",
	ClosingHash:   "closing_hash",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewAccountingPeriodsDBFieldNameFromStr(field string) (dbField AccountingPeriodsDBFieldNameType, found bool) {
	switch field {

	case string(AccountingPeriodsDBFieldName.Id):
		return AccountingPeriodsDBFieldName.Id, true

	case string(AccountingPeriodsDBFieldName.BookId):
		return AccountingPeriodsDBFieldName.BookId, true

	case string(AccountingPeriodsDBFieldName.PeriodCode):
		return AccountingPeriodsDBFieldName.PeriodCode, true

	case string(AccountingPeriodsDBFieldName.PeriodType):
		return AccountingPeriodsDBFieldName.PeriodType, true

	case string(AccountingPeriodsDBFieldName.PeriodStart):
		return AccountingPeriodsDBFieldName.PeriodStart, true

	case string(AccountingPeriodsDBFieldName.PeriodEnd):
		return AccountingPeriodsDBFieldName.PeriodEnd, true

	case string(AccountingPeriodsDBFieldName.PeriodStatus):
		return AccountingPeriodsDBFieldName.PeriodStatus, true

	case string(AccountingPeriodsDBFieldName.ClosedAt):
		return AccountingPeriodsDBFieldName.ClosedAt, true

	case string(AccountingPeriodsDBFieldName.ClosedBy):
		return AccountingPeriodsDBFieldName.ClosedBy, true

	case string(AccountingPeriodsDBFieldName.LockReason):
		return AccountingPeriodsDBFieldName.LockReason, true

	case string(AccountingPeriodsDBFieldName.ClosingHash):
		return AccountingPeriodsDBFieldName.ClosingHash, true

	case string(AccountingPeriodsDBFieldName.Metadata):
		return AccountingPeriodsDBFieldName.Metadata, true

	case string(AccountingPeriodsDBFieldName.MetaCreatedAt):
		return AccountingPeriodsDBFieldName.MetaCreatedAt, true

	case string(AccountingPeriodsDBFieldName.MetaCreatedBy):
		return AccountingPeriodsDBFieldName.MetaCreatedBy, true

	case string(AccountingPeriodsDBFieldName.MetaUpdatedAt):
		return AccountingPeriodsDBFieldName.MetaUpdatedAt, true

	case string(AccountingPeriodsDBFieldName.MetaUpdatedBy):
		return AccountingPeriodsDBFieldName.MetaUpdatedBy, true

	case string(AccountingPeriodsDBFieldName.MetaDeletedAt):
		return AccountingPeriodsDBFieldName.MetaDeletedAt, true

	case string(AccountingPeriodsDBFieldName.MetaDeletedBy):
		return AccountingPeriodsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var AccountingPeriodsFilterJoins = map[string]JoinSpec{}

var AccountingPeriodsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"book_id": {
		SourcePath:        "book_id",
		DefaultOutputPath: "bookId",
		Column:            "book_id",
		SQLAlias:          "book_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"period_code": {
		SourcePath:        "period_code",
		DefaultOutputPath: "periodCode",
		Column:            "period_code",
		SQLAlias:          "period_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"period_type": {
		SourcePath:        "period_type",
		DefaultOutputPath: "periodType",
		Column:            "period_type",
		SQLAlias:          "period_type",
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
	"period_status": {
		SourcePath:        "period_status",
		DefaultOutputPath: "periodStatus",
		Column:            "period_status",
		SQLAlias:          "period_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"closed_at": {
		SourcePath:        "closed_at",
		DefaultOutputPath: "closedAt",
		Column:            "closed_at",
		SQLAlias:          "closed_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"closed_by": {
		SourcePath:        "closed_by",
		DefaultOutputPath: "closedBy",
		Column:            "closed_by",
		SQLAlias:          "closed_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"lock_reason": {
		SourcePath:        "lock_reason",
		DefaultOutputPath: "lockReason",
		Column:            "lock_reason",
		SQLAlias:          "lock_reason",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"closing_hash": {
		SourcePath:        "closing_hash",
		DefaultOutputPath: "closingHash",
		Column:            "closing_hash",
		SQLAlias:          "closing_hash",
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

func NewAccountingPeriodsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = AccountingPeriodsFilterFields[field]
	return
}

type AccountingPeriodsFilterResult struct {
	AccountingPeriods
	FilterCount int `db:"count"`
}

func ValidateAccountingPeriodsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewAccountingPeriodsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewAccountingPeriodsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewAccountingPeriodsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateAccountingPeriodsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateAccountingPeriodsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewAccountingPeriodsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateAccountingPeriodsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PeriodStatus string

const (
	PeriodStatusOpen   PeriodStatus = "open"
	PeriodStatusClosed PeriodStatus = "closed"
	PeriodStatusLocked PeriodStatus = "locked"
)

type AccountingPeriods struct {
	Id           uuid.UUID       `db:"id"`
	BookId       uuid.UUID       `db:"book_id"`
	PeriodCode   string          `db:"period_code"`
	PeriodType   string          `db:"period_type"`
	PeriodStart  time.Time       `db:"period_start"`
	PeriodEnd    time.Time       `db:"period_end"`
	PeriodStatus PeriodStatus    `db:"period_status"`
	ClosedAt     null.Time       `db:"closed_at"`
	ClosedBy     nuuid.NUUID     `db:"closed_by"`
	LockReason   null.String     `db:"lock_reason"`
	ClosingHash  null.String     `db:"closing_hash"`
	Metadata     json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type AccountingPeriodsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d AccountingPeriods) ToAccountingPeriodsPrimaryID() AccountingPeriodsPrimaryID {
	return AccountingPeriodsPrimaryID{
		Id: d.Id,
	}
}

type AccountingPeriodsList []*AccountingPeriods

type AccountingPeriodsFilterResultList []*AccountingPeriodsFilterResult
