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

type ReconciliationCandidatesDBFieldNameType string

type reconciliationCandidatesDBFieldName struct {
	Id                  ReconciliationCandidatesDBFieldNameType
	ReconciliationRunId ReconciliationCandidatesDBFieldNameType
	SourceSystem        ReconciliationCandidatesDBFieldNameType
	SourceRefId         ReconciliationCandidatesDBFieldNameType
	CandidateKey        ReconciliationCandidatesDBFieldNameType
	Amount              ReconciliationCandidatesDBFieldNameType
	OccurredAt          ReconciliationCandidatesDBFieldNameType
	NormalizedPayload   ReconciliationCandidatesDBFieldNameType
	Metadata            ReconciliationCandidatesDBFieldNameType
	MetaCreatedAt       ReconciliationCandidatesDBFieldNameType
	MetaCreatedBy       ReconciliationCandidatesDBFieldNameType
	MetaUpdatedAt       ReconciliationCandidatesDBFieldNameType
	MetaUpdatedBy       ReconciliationCandidatesDBFieldNameType
	MetaDeletedAt       ReconciliationCandidatesDBFieldNameType
	MetaDeletedBy       ReconciliationCandidatesDBFieldNameType
}

var ReconciliationCandidatesDBFieldName = reconciliationCandidatesDBFieldName{
	Id:                  "id",
	ReconciliationRunId: "reconciliation_run_id",
	SourceSystem:        "source_system",
	SourceRefId:         "source_ref_id",
	CandidateKey:        "candidate_key",
	Amount:              "amount",
	OccurredAt:          "occurred_at",
	NormalizedPayload:   "normalized_payload",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewReconciliationCandidatesDBFieldNameFromStr(field string) (dbField ReconciliationCandidatesDBFieldNameType, found bool) {
	switch field {

	case string(ReconciliationCandidatesDBFieldName.Id):
		return ReconciliationCandidatesDBFieldName.Id, true

	case string(ReconciliationCandidatesDBFieldName.ReconciliationRunId):
		return ReconciliationCandidatesDBFieldName.ReconciliationRunId, true

	case string(ReconciliationCandidatesDBFieldName.SourceSystem):
		return ReconciliationCandidatesDBFieldName.SourceSystem, true

	case string(ReconciliationCandidatesDBFieldName.SourceRefId):
		return ReconciliationCandidatesDBFieldName.SourceRefId, true

	case string(ReconciliationCandidatesDBFieldName.CandidateKey):
		return ReconciliationCandidatesDBFieldName.CandidateKey, true

	case string(ReconciliationCandidatesDBFieldName.Amount):
		return ReconciliationCandidatesDBFieldName.Amount, true

	case string(ReconciliationCandidatesDBFieldName.OccurredAt):
		return ReconciliationCandidatesDBFieldName.OccurredAt, true

	case string(ReconciliationCandidatesDBFieldName.NormalizedPayload):
		return ReconciliationCandidatesDBFieldName.NormalizedPayload, true

	case string(ReconciliationCandidatesDBFieldName.Metadata):
		return ReconciliationCandidatesDBFieldName.Metadata, true

	case string(ReconciliationCandidatesDBFieldName.MetaCreatedAt):
		return ReconciliationCandidatesDBFieldName.MetaCreatedAt, true

	case string(ReconciliationCandidatesDBFieldName.MetaCreatedBy):
		return ReconciliationCandidatesDBFieldName.MetaCreatedBy, true

	case string(ReconciliationCandidatesDBFieldName.MetaUpdatedAt):
		return ReconciliationCandidatesDBFieldName.MetaUpdatedAt, true

	case string(ReconciliationCandidatesDBFieldName.MetaUpdatedBy):
		return ReconciliationCandidatesDBFieldName.MetaUpdatedBy, true

	case string(ReconciliationCandidatesDBFieldName.MetaDeletedAt):
		return ReconciliationCandidatesDBFieldName.MetaDeletedAt, true

	case string(ReconciliationCandidatesDBFieldName.MetaDeletedBy):
		return ReconciliationCandidatesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ReconciliationCandidatesFilterJoins = map[string]JoinSpec{}

var ReconciliationCandidatesFilterFields = map[string]FilterFieldSpec{
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
	"source_system": {
		SourcePath:        "source_system",
		DefaultOutputPath: "sourceSystem",
		Column:            "source_system",
		SQLAlias:          "source_system",
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
	"candidate_key": {
		SourcePath:        "candidate_key",
		DefaultOutputPath: "candidateKey",
		Column:            "candidate_key",
		SQLAlias:          "candidate_key",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"occurred_at": {
		SourcePath:        "occurred_at",
		DefaultOutputPath: "occurredAt",
		Column:            "occurred_at",
		SQLAlias:          "occurred_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"normalized_payload": {
		SourcePath:        "normalized_payload",
		DefaultOutputPath: "normalizedPayload",
		Column:            "normalized_payload",
		SQLAlias:          "normalized_payload",
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

func NewReconciliationCandidatesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ReconciliationCandidatesFilterFields[field]
	return
}

type ReconciliationCandidatesFilterResult struct {
	ReconciliationCandidates
	FilterCount int `db:"count"`
}

func ValidateReconciliationCandidatesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewReconciliationCandidatesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewReconciliationCandidatesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewReconciliationCandidatesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateReconciliationCandidatesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateReconciliationCandidatesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewReconciliationCandidatesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateReconciliationCandidatesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type SourceSystem string

const (
	SourceSystemLedger            SourceSystem = "ledger"
	SourceSystemProviderStatement SourceSystem = "provider_statement"
	SourceSystemBankStatement     SourceSystem = "bank_statement"
)

type ReconciliationCandidates struct {
	Id                  uuid.UUID       `db:"id"`
	ReconciliationRunId uuid.UUID       `db:"reconciliation_run_id"`
	SourceSystem        SourceSystem    `db:"source_system"`
	SourceRefId         uuid.UUID       `db:"source_ref_id"`
	CandidateKey        string          `db:"candidate_key"`
	Amount              decimal.Decimal `db:"amount"`
	OccurredAt          null.Time       `db:"occurred_at"`
	NormalizedPayload   json.RawMessage `db:"normalized_payload"`
	Metadata            json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ReconciliationCandidatesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ReconciliationCandidates) ToReconciliationCandidatesPrimaryID() ReconciliationCandidatesPrimaryID {
	return ReconciliationCandidatesPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationCandidatesList []*ReconciliationCandidates

type ReconciliationCandidatesFilterResultList []*ReconciliationCandidatesFilterResult
