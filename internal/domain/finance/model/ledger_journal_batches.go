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

type LedgerJournalBatchesDBFieldNameType string

type ledgerJournalBatchesDBFieldName struct {
	Id            LedgerJournalBatchesDBFieldNameType
	BatchCode     LedgerJournalBatchesDBFieldNameType
	BatchType     LedgerJournalBatchesDBFieldNameType
	SourceRef     LedgerJournalBatchesDBFieldNameType
	BatchStatus   LedgerJournalBatchesDBFieldNameType
	BookedAt      LedgerJournalBatchesDBFieldNameType
	Description   LedgerJournalBatchesDBFieldNameType
	Metadata      LedgerJournalBatchesDBFieldNameType
	MetaCreatedAt LedgerJournalBatchesDBFieldNameType
	MetaCreatedBy LedgerJournalBatchesDBFieldNameType
	MetaUpdatedAt LedgerJournalBatchesDBFieldNameType
	MetaUpdatedBy LedgerJournalBatchesDBFieldNameType
	MetaDeletedAt LedgerJournalBatchesDBFieldNameType
	MetaDeletedBy LedgerJournalBatchesDBFieldNameType
}

var LedgerJournalBatchesDBFieldName = ledgerJournalBatchesDBFieldName{
	Id:            "id",
	BatchCode:     "batch_code",
	BatchType:     "batch_type",
	SourceRef:     "source_ref",
	BatchStatus:   "batch_status",
	BookedAt:      "booked_at",
	Description:   "description",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewLedgerJournalBatchesDBFieldNameFromStr(field string) (dbField LedgerJournalBatchesDBFieldNameType, found bool) {
	switch field {

	case string(LedgerJournalBatchesDBFieldName.Id):
		return LedgerJournalBatchesDBFieldName.Id, true

	case string(LedgerJournalBatchesDBFieldName.BatchCode):
		return LedgerJournalBatchesDBFieldName.BatchCode, true

	case string(LedgerJournalBatchesDBFieldName.BatchType):
		return LedgerJournalBatchesDBFieldName.BatchType, true

	case string(LedgerJournalBatchesDBFieldName.SourceRef):
		return LedgerJournalBatchesDBFieldName.SourceRef, true

	case string(LedgerJournalBatchesDBFieldName.BatchStatus):
		return LedgerJournalBatchesDBFieldName.BatchStatus, true

	case string(LedgerJournalBatchesDBFieldName.BookedAt):
		return LedgerJournalBatchesDBFieldName.BookedAt, true

	case string(LedgerJournalBatchesDBFieldName.Description):
		return LedgerJournalBatchesDBFieldName.Description, true

	case string(LedgerJournalBatchesDBFieldName.Metadata):
		return LedgerJournalBatchesDBFieldName.Metadata, true

	case string(LedgerJournalBatchesDBFieldName.MetaCreatedAt):
		return LedgerJournalBatchesDBFieldName.MetaCreatedAt, true

	case string(LedgerJournalBatchesDBFieldName.MetaCreatedBy):
		return LedgerJournalBatchesDBFieldName.MetaCreatedBy, true

	case string(LedgerJournalBatchesDBFieldName.MetaUpdatedAt):
		return LedgerJournalBatchesDBFieldName.MetaUpdatedAt, true

	case string(LedgerJournalBatchesDBFieldName.MetaUpdatedBy):
		return LedgerJournalBatchesDBFieldName.MetaUpdatedBy, true

	case string(LedgerJournalBatchesDBFieldName.MetaDeletedAt):
		return LedgerJournalBatchesDBFieldName.MetaDeletedAt, true

	case string(LedgerJournalBatchesDBFieldName.MetaDeletedBy):
		return LedgerJournalBatchesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerJournalBatchesFilterJoins = map[string]JoinSpec{}

var LedgerJournalBatchesFilterFields = map[string]FilterFieldSpec{
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
	"batch_type": {
		SourcePath:        "batch_type",
		DefaultOutputPath: "batchType",
		Column:            "batch_type",
		SQLAlias:          "batch_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_ref": {
		SourcePath:        "source_ref",
		DefaultOutputPath: "sourceRef",
		Column:            "source_ref",
		SQLAlias:          "source_ref",
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

func NewLedgerJournalBatchesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerJournalBatchesFilterFields[field]
	return
}

type LedgerJournalBatchesFilterResult struct {
	LedgerJournalBatches
	FilterCount int `db:"count"`
}

func ValidateLedgerJournalBatchesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerJournalBatchesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerJournalBatchesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerJournalBatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerJournalBatchesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerJournalBatchesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerJournalBatchesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerJournalBatchesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type BatchType string

const (
	BatchTypePayment        BatchType = "payment"
	BatchTypeRefund         BatchType = "refund"
	BatchTypeSettlement     BatchType = "settlement"
	BatchTypePayout         BatchType = "payout"
	BatchTypeChargeback     BatchType = "chargeback"
	BatchTypeAdjustment     BatchType = "adjustment"
	BatchTypeClose          BatchType = "close"
	BatchTypeReconciliation BatchType = "reconciliation"
)

type LedgerJournalBatchesBatchStatus string

const (
	LedgerJournalBatchesBatchStatusOpen    LedgerJournalBatchesBatchStatus = "open"
	LedgerJournalBatchesBatchStatusPosting LedgerJournalBatchesBatchStatus = "posting"
	LedgerJournalBatchesBatchStatusPosted  LedgerJournalBatchesBatchStatus = "posted"
	LedgerJournalBatchesBatchStatusFailed  LedgerJournalBatchesBatchStatus = "failed"
	LedgerJournalBatchesBatchStatusClosed  LedgerJournalBatchesBatchStatus = "closed"
)

type LedgerJournalBatches struct {
	Id          uuid.UUID                       `db:"id"`
	BatchCode   string                          `db:"batch_code"`
	BatchType   BatchType                       `db:"batch_type"`
	SourceRef   null.String                     `db:"source_ref"`
	BatchStatus LedgerJournalBatchesBatchStatus `db:"batch_status"`
	BookedAt    time.Time                       `db:"booked_at"`
	Description null.String                     `db:"description"`
	Metadata    json.RawMessage                 `db:"metadata"`

	shared.MetaSignature
}
type LedgerJournalBatchesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerJournalBatches) ToLedgerJournalBatchesPrimaryID() LedgerJournalBatchesPrimaryID {
	return LedgerJournalBatchesPrimaryID{
		Id: d.Id,
	}
}

type LedgerJournalBatchesList []*LedgerJournalBatches

type LedgerJournalBatchesFilterResultList []*LedgerJournalBatchesFilterResult
