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

type LedgerJournalEntriesDBFieldNameType string

type ledgerJournalEntriesDBFieldName struct {
	Id                  LedgerJournalEntriesDBFieldNameType
	BatchId             LedgerJournalEntriesDBFieldNameType
	JournalCode         LedgerJournalEntriesDBFieldNameType
	BookId              LedgerJournalEntriesDBFieldNameType
	JournalType         LedgerJournalEntriesDBFieldNameType
	SourceType          LedgerJournalEntriesDBFieldNameType
	SourceId            LedgerJournalEntriesDBFieldNameType
	IdempotencyKey      LedgerJournalEntriesDBFieldNameType
	JournalStatus       LedgerJournalEntriesDBFieldNameType
	EffectiveAt         LedgerJournalEntriesDBFieldNameType
	BookedAt            LedgerJournalEntriesDBFieldNameType
	Description         LedgerJournalEntriesDBFieldNameType
	ReversalOfJournalId LedgerJournalEntriesDBFieldNameType
	Metadata            LedgerJournalEntriesDBFieldNameType
	MetaCreatedAt       LedgerJournalEntriesDBFieldNameType
	MetaCreatedBy       LedgerJournalEntriesDBFieldNameType
	MetaUpdatedAt       LedgerJournalEntriesDBFieldNameType
	MetaUpdatedBy       LedgerJournalEntriesDBFieldNameType
	MetaDeletedAt       LedgerJournalEntriesDBFieldNameType
	MetaDeletedBy       LedgerJournalEntriesDBFieldNameType
}

var LedgerJournalEntriesDBFieldName = ledgerJournalEntriesDBFieldName{
	Id:                  "id",
	BatchId:             "batch_id",
	JournalCode:         "journal_code",
	BookId:              "book_id",
	JournalType:         "journal_type",
	SourceType:          "source_type",
	SourceId:            "source_id",
	IdempotencyKey:      "idempotency_key",
	JournalStatus:       "journal_status",
	EffectiveAt:         "effective_at",
	BookedAt:            "booked_at",
	Description:         "description",
	ReversalOfJournalId: "reversal_of_journal_id",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewLedgerJournalEntriesDBFieldNameFromStr(field string) (dbField LedgerJournalEntriesDBFieldNameType, found bool) {
	switch field {

	case string(LedgerJournalEntriesDBFieldName.Id):
		return LedgerJournalEntriesDBFieldName.Id, true

	case string(LedgerJournalEntriesDBFieldName.BatchId):
		return LedgerJournalEntriesDBFieldName.BatchId, true

	case string(LedgerJournalEntriesDBFieldName.JournalCode):
		return LedgerJournalEntriesDBFieldName.JournalCode, true

	case string(LedgerJournalEntriesDBFieldName.BookId):
		return LedgerJournalEntriesDBFieldName.BookId, true

	case string(LedgerJournalEntriesDBFieldName.JournalType):
		return LedgerJournalEntriesDBFieldName.JournalType, true

	case string(LedgerJournalEntriesDBFieldName.SourceType):
		return LedgerJournalEntriesDBFieldName.SourceType, true

	case string(LedgerJournalEntriesDBFieldName.SourceId):
		return LedgerJournalEntriesDBFieldName.SourceId, true

	case string(LedgerJournalEntriesDBFieldName.IdempotencyKey):
		return LedgerJournalEntriesDBFieldName.IdempotencyKey, true

	case string(LedgerJournalEntriesDBFieldName.JournalStatus):
		return LedgerJournalEntriesDBFieldName.JournalStatus, true

	case string(LedgerJournalEntriesDBFieldName.EffectiveAt):
		return LedgerJournalEntriesDBFieldName.EffectiveAt, true

	case string(LedgerJournalEntriesDBFieldName.BookedAt):
		return LedgerJournalEntriesDBFieldName.BookedAt, true

	case string(LedgerJournalEntriesDBFieldName.Description):
		return LedgerJournalEntriesDBFieldName.Description, true

	case string(LedgerJournalEntriesDBFieldName.ReversalOfJournalId):
		return LedgerJournalEntriesDBFieldName.ReversalOfJournalId, true

	case string(LedgerJournalEntriesDBFieldName.Metadata):
		return LedgerJournalEntriesDBFieldName.Metadata, true

	case string(LedgerJournalEntriesDBFieldName.MetaCreatedAt):
		return LedgerJournalEntriesDBFieldName.MetaCreatedAt, true

	case string(LedgerJournalEntriesDBFieldName.MetaCreatedBy):
		return LedgerJournalEntriesDBFieldName.MetaCreatedBy, true

	case string(LedgerJournalEntriesDBFieldName.MetaUpdatedAt):
		return LedgerJournalEntriesDBFieldName.MetaUpdatedAt, true

	case string(LedgerJournalEntriesDBFieldName.MetaUpdatedBy):
		return LedgerJournalEntriesDBFieldName.MetaUpdatedBy, true

	case string(LedgerJournalEntriesDBFieldName.MetaDeletedAt):
		return LedgerJournalEntriesDBFieldName.MetaDeletedAt, true

	case string(LedgerJournalEntriesDBFieldName.MetaDeletedBy):
		return LedgerJournalEntriesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerJournalEntriesFilterJoins = map[string]JoinSpec{}

var LedgerJournalEntriesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"batch_id": {
		SourcePath:        "batch_id",
		DefaultOutputPath: "batchId",
		Column:            "batch_id",
		SQLAlias:          "batch_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"journal_code": {
		SourcePath:        "journal_code",
		DefaultOutputPath: "journalCode",
		Column:            "journal_code",
		SQLAlias:          "journal_code",
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
	"journal_type": {
		SourcePath:        "journal_type",
		DefaultOutputPath: "journalType",
		Column:            "journal_type",
		SQLAlias:          "journal_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_type": {
		SourcePath:        "source_type",
		DefaultOutputPath: "sourceType",
		Column:            "source_type",
		SQLAlias:          "source_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_id": {
		SourcePath:        "source_id",
		DefaultOutputPath: "sourceId",
		Column:            "source_id",
		SQLAlias:          "source_id",
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
	"journal_status": {
		SourcePath:        "journal_status",
		DefaultOutputPath: "journalStatus",
		Column:            "journal_status",
		SQLAlias:          "journal_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"effective_at": {
		SourcePath:        "effective_at",
		DefaultOutputPath: "effectiveAt",
		Column:            "effective_at",
		SQLAlias:          "effective_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"booked_at": {
		SourcePath:        "booked_at",
		DefaultOutputPath: "bookedAt",
		Column:            "booked_at",
		SQLAlias:          "booked_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"description": {
		SourcePath:        "description",
		DefaultOutputPath: "description",
		Column:            "description",
		SQLAlias:          "description",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reversal_of_journal_id": {
		SourcePath:        "reversal_of_journal_id",
		DefaultOutputPath: "reversalOfJournalId",
		Column:            "reversal_of_journal_id",
		SQLAlias:          "reversal_of_journal_id",
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

func NewLedgerJournalEntriesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerJournalEntriesFilterFields[field]
	return
}

type LedgerJournalEntriesFilterResult struct {
	LedgerJournalEntries
	FilterCount int `db:"count"`
}

func ValidateLedgerJournalEntriesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerJournalEntriesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerJournalEntriesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerJournalEntriesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerJournalEntriesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerJournalEntriesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerJournalEntriesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerJournalEntriesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type JournalStatus string

const (
	JournalStatusDraft     JournalStatus = "draft"
	JournalStatusValidated JournalStatus = "validated"
	JournalStatusPosted    JournalStatus = "posted"
	JournalStatusReversed  JournalStatus = "reversed"
	JournalStatusFailed    JournalStatus = "failed"
)

type JournalType string

const (
	JournalTypePayment        JournalType = "payment"
	JournalTypeRefund         JournalType = "refund"
	JournalTypeSettlement     JournalType = "settlement"
	JournalTypePayout         JournalType = "payout"
	JournalTypeChargeback     JournalType = "chargeback"
	JournalTypeAdjustment     JournalType = "adjustment"
	JournalTypeClose          JournalType = "close"
	JournalTypeReconciliation JournalType = "reconciliation"
	JournalTypeRevaluation    JournalType = "revaluation"
)

type LedgerJournalEntries struct {
	Id                  uuid.UUID       `db:"id"`
	BatchId             nuuid.NUUID     `db:"batch_id"`
	JournalCode         string          `db:"journal_code"`
	BookId              uuid.UUID       `db:"book_id"`
	JournalType         JournalType     `db:"journal_type"`
	SourceType          string          `db:"source_type"`
	SourceId            uuid.UUID       `db:"source_id"`
	IdempotencyKey      string          `db:"idempotency_key"`
	JournalStatus       JournalStatus   `db:"journal_status"`
	EffectiveAt         time.Time       `db:"effective_at"`
	BookedAt            time.Time       `db:"booked_at"`
	Description         null.String     `db:"description"`
	ReversalOfJournalId nuuid.NUUID     `db:"reversal_of_journal_id"`
	Metadata            json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type LedgerJournalEntriesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerJournalEntries) ToLedgerJournalEntriesPrimaryID() LedgerJournalEntriesPrimaryID {
	return LedgerJournalEntriesPrimaryID{
		Id: d.Id,
	}
}

type LedgerJournalEntriesList []*LedgerJournalEntries

type LedgerJournalEntriesFilterResultList []*LedgerJournalEntriesFilterResult
