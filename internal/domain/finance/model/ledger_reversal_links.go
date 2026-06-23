package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerReversalLinksDBFieldNameType string

type ledgerReversalLinksDBFieldName struct {
	Id                LedgerReversalLinksDBFieldNameType
	OriginalJournalId LedgerReversalLinksDBFieldNameType
	ReversalJournalId LedgerReversalLinksDBFieldNameType
	ReversalReason    LedgerReversalLinksDBFieldNameType
	ReversedAt        LedgerReversalLinksDBFieldNameType
	Metadata          LedgerReversalLinksDBFieldNameType
	MetaCreatedAt     LedgerReversalLinksDBFieldNameType
	MetaCreatedBy     LedgerReversalLinksDBFieldNameType
	MetaUpdatedAt     LedgerReversalLinksDBFieldNameType
	MetaUpdatedBy     LedgerReversalLinksDBFieldNameType
	MetaDeletedAt     LedgerReversalLinksDBFieldNameType
	MetaDeletedBy     LedgerReversalLinksDBFieldNameType
}

var LedgerReversalLinksDBFieldName = ledgerReversalLinksDBFieldName{
	Id:                "id",
	OriginalJournalId: "original_journal_id",
	ReversalJournalId: "reversal_journal_id",
	ReversalReason:    "reversal_reason",
	ReversedAt:        "reversed_at",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewLedgerReversalLinksDBFieldNameFromStr(field string) (dbField LedgerReversalLinksDBFieldNameType, found bool) {
	switch field {

	case string(LedgerReversalLinksDBFieldName.Id):
		return LedgerReversalLinksDBFieldName.Id, true

	case string(LedgerReversalLinksDBFieldName.OriginalJournalId):
		return LedgerReversalLinksDBFieldName.OriginalJournalId, true

	case string(LedgerReversalLinksDBFieldName.ReversalJournalId):
		return LedgerReversalLinksDBFieldName.ReversalJournalId, true

	case string(LedgerReversalLinksDBFieldName.ReversalReason):
		return LedgerReversalLinksDBFieldName.ReversalReason, true

	case string(LedgerReversalLinksDBFieldName.ReversedAt):
		return LedgerReversalLinksDBFieldName.ReversedAt, true

	case string(LedgerReversalLinksDBFieldName.Metadata):
		return LedgerReversalLinksDBFieldName.Metadata, true

	case string(LedgerReversalLinksDBFieldName.MetaCreatedAt):
		return LedgerReversalLinksDBFieldName.MetaCreatedAt, true

	case string(LedgerReversalLinksDBFieldName.MetaCreatedBy):
		return LedgerReversalLinksDBFieldName.MetaCreatedBy, true

	case string(LedgerReversalLinksDBFieldName.MetaUpdatedAt):
		return LedgerReversalLinksDBFieldName.MetaUpdatedAt, true

	case string(LedgerReversalLinksDBFieldName.MetaUpdatedBy):
		return LedgerReversalLinksDBFieldName.MetaUpdatedBy, true

	case string(LedgerReversalLinksDBFieldName.MetaDeletedAt):
		return LedgerReversalLinksDBFieldName.MetaDeletedAt, true

	case string(LedgerReversalLinksDBFieldName.MetaDeletedBy):
		return LedgerReversalLinksDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerReversalLinksFilterJoins = map[string]JoinSpec{}

var LedgerReversalLinksFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"original_journal_id": {
		SourcePath:        "original_journal_id",
		DefaultOutputPath: "originalJournalId",
		Column:            "original_journal_id",
		SQLAlias:          "original_journal_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reversal_journal_id": {
		SourcePath:        "reversal_journal_id",
		DefaultOutputPath: "reversalJournalId",
		Column:            "reversal_journal_id",
		SQLAlias:          "reversal_journal_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reversal_reason": {
		SourcePath:        "reversal_reason",
		DefaultOutputPath: "reversalReason",
		Column:            "reversal_reason",
		SQLAlias:          "reversal_reason",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reversed_at": {
		SourcePath:        "reversed_at",
		DefaultOutputPath: "reversedAt",
		Column:            "reversed_at",
		SQLAlias:          "reversed_at",
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

func NewLedgerReversalLinksFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerReversalLinksFilterFields[field]
	return
}

type LedgerReversalLinksFilterResult struct {
	LedgerReversalLinks
	FilterCount int `db:"count"`
}

func ValidateLedgerReversalLinksFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerReversalLinksFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerReversalLinksFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerReversalLinksFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerReversalLinksFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerReversalLinksFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerReversalLinksFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerReversalLinksFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type LedgerReversalLinks struct {
	Id                uuid.UUID       `db:"id"`
	OriginalJournalId uuid.UUID       `db:"original_journal_id"`
	ReversalJournalId uuid.UUID       `db:"reversal_journal_id"`
	ReversalReason    string          `db:"reversal_reason"`
	ReversedAt        time.Time       `db:"reversed_at"`
	Metadata          json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type LedgerReversalLinksPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerReversalLinks) ToLedgerReversalLinksPrimaryID() LedgerReversalLinksPrimaryID {
	return LedgerReversalLinksPrimaryID{
		Id: d.Id,
	}
}

type LedgerReversalLinksList []*LedgerReversalLinks

type LedgerReversalLinksFilterResultList []*LedgerReversalLinksFilterResult
