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

type ReconciliationMatchesDBFieldNameType string

type reconciliationMatchesDBFieldName struct {
	Id                  ReconciliationMatchesDBFieldNameType
	ReconciliationRunId ReconciliationMatchesDBFieldNameType
	LedgerJournalId     ReconciliationMatchesDBFieldNameType
	StatementLineId     ReconciliationMatchesDBFieldNameType
	MatchType           ReconciliationMatchesDBFieldNameType
	MatchStatus         ReconciliationMatchesDBFieldNameType
	AmountDifference    ReconciliationMatchesDBFieldNameType
	MatchedAt           ReconciliationMatchesDBFieldNameType
	Metadata            ReconciliationMatchesDBFieldNameType
	MetaCreatedAt       ReconciliationMatchesDBFieldNameType
	MetaCreatedBy       ReconciliationMatchesDBFieldNameType
	MetaUpdatedAt       ReconciliationMatchesDBFieldNameType
	MetaUpdatedBy       ReconciliationMatchesDBFieldNameType
	MetaDeletedAt       ReconciliationMatchesDBFieldNameType
	MetaDeletedBy       ReconciliationMatchesDBFieldNameType
}

var ReconciliationMatchesDBFieldName = reconciliationMatchesDBFieldName{
	Id:                  "id",
	ReconciliationRunId: "reconciliation_run_id",
	LedgerJournalId:     "ledger_journal_id",
	StatementLineId:     "statement_line_id",
	MatchType:           "match_type",
	MatchStatus:         "match_status",
	AmountDifference:    "amount_difference",
	MatchedAt:           "matched_at",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewReconciliationMatchesDBFieldNameFromStr(field string) (dbField ReconciliationMatchesDBFieldNameType, found bool) {
	switch field {

	case string(ReconciliationMatchesDBFieldName.Id):
		return ReconciliationMatchesDBFieldName.Id, true

	case string(ReconciliationMatchesDBFieldName.ReconciliationRunId):
		return ReconciliationMatchesDBFieldName.ReconciliationRunId, true

	case string(ReconciliationMatchesDBFieldName.LedgerJournalId):
		return ReconciliationMatchesDBFieldName.LedgerJournalId, true

	case string(ReconciliationMatchesDBFieldName.StatementLineId):
		return ReconciliationMatchesDBFieldName.StatementLineId, true

	case string(ReconciliationMatchesDBFieldName.MatchType):
		return ReconciliationMatchesDBFieldName.MatchType, true

	case string(ReconciliationMatchesDBFieldName.MatchStatus):
		return ReconciliationMatchesDBFieldName.MatchStatus, true

	case string(ReconciliationMatchesDBFieldName.AmountDifference):
		return ReconciliationMatchesDBFieldName.AmountDifference, true

	case string(ReconciliationMatchesDBFieldName.MatchedAt):
		return ReconciliationMatchesDBFieldName.MatchedAt, true

	case string(ReconciliationMatchesDBFieldName.Metadata):
		return ReconciliationMatchesDBFieldName.Metadata, true

	case string(ReconciliationMatchesDBFieldName.MetaCreatedAt):
		return ReconciliationMatchesDBFieldName.MetaCreatedAt, true

	case string(ReconciliationMatchesDBFieldName.MetaCreatedBy):
		return ReconciliationMatchesDBFieldName.MetaCreatedBy, true

	case string(ReconciliationMatchesDBFieldName.MetaUpdatedAt):
		return ReconciliationMatchesDBFieldName.MetaUpdatedAt, true

	case string(ReconciliationMatchesDBFieldName.MetaUpdatedBy):
		return ReconciliationMatchesDBFieldName.MetaUpdatedBy, true

	case string(ReconciliationMatchesDBFieldName.MetaDeletedAt):
		return ReconciliationMatchesDBFieldName.MetaDeletedAt, true

	case string(ReconciliationMatchesDBFieldName.MetaDeletedBy):
		return ReconciliationMatchesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ReconciliationMatchesFilterJoins = map[string]JoinSpec{}

var ReconciliationMatchesFilterFields = map[string]FilterFieldSpec{
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
	"ledger_journal_id": {
		SourcePath:        "ledger_journal_id",
		DefaultOutputPath: "ledgerJournalId",
		Column:            "ledger_journal_id",
		SQLAlias:          "ledger_journal_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"statement_line_id": {
		SourcePath:        "statement_line_id",
		DefaultOutputPath: "statementLineId",
		Column:            "statement_line_id",
		SQLAlias:          "statement_line_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"match_type": {
		SourcePath:        "match_type",
		DefaultOutputPath: "matchType",
		Column:            "match_type",
		SQLAlias:          "match_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"match_status": {
		SourcePath:        "match_status",
		DefaultOutputPath: "matchStatus",
		Column:            "match_status",
		SQLAlias:          "match_status",
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
	"matched_at": {
		SourcePath:        "matched_at",
		DefaultOutputPath: "matchedAt",
		Column:            "matched_at",
		SQLAlias:          "matched_at",
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

func NewReconciliationMatchesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ReconciliationMatchesFilterFields[field]
	return
}

type ReconciliationMatchesFilterResult struct {
	ReconciliationMatches
	FilterCount int `db:"count"`
}

func ValidateReconciliationMatchesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewReconciliationMatchesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewReconciliationMatchesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewReconciliationMatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateReconciliationMatchesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateReconciliationMatchesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewReconciliationMatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateReconciliationMatchesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type MatchStatus string

const (
	MatchStatusMatched  MatchStatus = "matched"
	MatchStatusReview   MatchStatus = "review"
	MatchStatusReversed MatchStatus = "reversed"
)

type MatchType string

const (
	MatchTypeExact       MatchType = "exact"
	MatchTypeAmountDate  MatchType = "amount_date"
	MatchTypeProviderRef MatchType = "provider_ref"
	MatchTypeManual      MatchType = "manual"
)

type ReconciliationMatches struct {
	Id                  uuid.UUID       `db:"id"`
	ReconciliationRunId uuid.UUID       `db:"reconciliation_run_id"`
	LedgerJournalId     nuuid.NUUID     `db:"ledger_journal_id"`
	StatementLineId     nuuid.NUUID     `db:"statement_line_id"`
	MatchType           MatchType       `db:"match_type"`
	MatchStatus         MatchStatus     `db:"match_status"`
	AmountDifference    decimal.Decimal `db:"amount_difference"`
	MatchedAt           time.Time       `db:"matched_at"`
	Metadata            json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ReconciliationMatchesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ReconciliationMatches) ToReconciliationMatchesPrimaryID() ReconciliationMatchesPrimaryID {
	return ReconciliationMatchesPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationMatchesList []*ReconciliationMatches

type ReconciliationMatchesFilterResultList []*ReconciliationMatchesFilterResult
