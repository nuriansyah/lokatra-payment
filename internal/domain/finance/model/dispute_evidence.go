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

type DisputeEvidenceDBFieldNameType string

type disputeEvidenceDBFieldName struct {
	Id             DisputeEvidenceDBFieldNameType
	DisputeId      DisputeEvidenceDBFieldNameType
	EvidenceType   DisputeEvidenceDBFieldNameType
	StorageUri     DisputeEvidenceDBFieldNameType
	ContentHash    DisputeEvidenceDBFieldNameType
	IdempotencyKey DisputeEvidenceDBFieldNameType
	SubmittedBy    DisputeEvidenceDBFieldNameType
	SubmittedAt    DisputeEvidenceDBFieldNameType
	Metadata       DisputeEvidenceDBFieldNameType
	MetaCreatedAt  DisputeEvidenceDBFieldNameType
	MetaCreatedBy  DisputeEvidenceDBFieldNameType
	MetaUpdatedAt  DisputeEvidenceDBFieldNameType
	MetaUpdatedBy  DisputeEvidenceDBFieldNameType
	MetaDeletedAt  DisputeEvidenceDBFieldNameType
	MetaDeletedBy  DisputeEvidenceDBFieldNameType
}

var DisputeEvidenceDBFieldName = disputeEvidenceDBFieldName{
	Id:             "id",
	DisputeId:      "dispute_id",
	EvidenceType:   "evidence_type",
	StorageUri:     "storage_uri",
	ContentHash:    "content_hash",
	IdempotencyKey: "idempotency_key",
	SubmittedBy:    "submitted_by",
	SubmittedAt:    "submitted_at",
	Metadata:       "metadata",
	MetaCreatedAt:  "meta_created_at",
	MetaCreatedBy:  "meta_created_by",
	MetaUpdatedAt:  "meta_updated_at",
	MetaUpdatedBy:  "meta_updated_by",
	MetaDeletedAt:  "meta_deleted_at",
	MetaDeletedBy:  "meta_deleted_by",
}

func NewDisputeEvidenceDBFieldNameFromStr(field string) (dbField DisputeEvidenceDBFieldNameType, found bool) {
	switch field {

	case string(DisputeEvidenceDBFieldName.Id):
		return DisputeEvidenceDBFieldName.Id, true

	case string(DisputeEvidenceDBFieldName.DisputeId):
		return DisputeEvidenceDBFieldName.DisputeId, true

	case string(DisputeEvidenceDBFieldName.EvidenceType):
		return DisputeEvidenceDBFieldName.EvidenceType, true

	case string(DisputeEvidenceDBFieldName.StorageUri):
		return DisputeEvidenceDBFieldName.StorageUri, true

	case string(DisputeEvidenceDBFieldName.ContentHash):
		return DisputeEvidenceDBFieldName.ContentHash, true

	case string(DisputeEvidenceDBFieldName.IdempotencyKey):
		return DisputeEvidenceDBFieldName.IdempotencyKey, true

	case string(DisputeEvidenceDBFieldName.SubmittedBy):
		return DisputeEvidenceDBFieldName.SubmittedBy, true

	case string(DisputeEvidenceDBFieldName.SubmittedAt):
		return DisputeEvidenceDBFieldName.SubmittedAt, true

	case string(DisputeEvidenceDBFieldName.Metadata):
		return DisputeEvidenceDBFieldName.Metadata, true

	case string(DisputeEvidenceDBFieldName.MetaCreatedAt):
		return DisputeEvidenceDBFieldName.MetaCreatedAt, true

	case string(DisputeEvidenceDBFieldName.MetaCreatedBy):
		return DisputeEvidenceDBFieldName.MetaCreatedBy, true

	case string(DisputeEvidenceDBFieldName.MetaUpdatedAt):
		return DisputeEvidenceDBFieldName.MetaUpdatedAt, true

	case string(DisputeEvidenceDBFieldName.MetaUpdatedBy):
		return DisputeEvidenceDBFieldName.MetaUpdatedBy, true

	case string(DisputeEvidenceDBFieldName.MetaDeletedAt):
		return DisputeEvidenceDBFieldName.MetaDeletedAt, true

	case string(DisputeEvidenceDBFieldName.MetaDeletedBy):
		return DisputeEvidenceDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var DisputeEvidenceFilterJoins = map[string]JoinSpec{}

var DisputeEvidenceFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"dispute_id": {
		SourcePath:        "dispute_id",
		DefaultOutputPath: "disputeId",
		Column:            "dispute_id",
		SQLAlias:          "dispute_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"evidence_type": {
		SourcePath:        "evidence_type",
		DefaultOutputPath: "evidenceType",
		Column:            "evidence_type",
		SQLAlias:          "evidence_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"storage_uri": {
		SourcePath:        "storage_uri",
		DefaultOutputPath: "storageUri",
		Column:            "storage_uri",
		SQLAlias:          "storage_uri",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"content_hash": {
		SourcePath:        "content_hash",
		DefaultOutputPath: "contentHash",
		Column:            "content_hash",
		SQLAlias:          "content_hash",
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
	"submitted_by": {
		SourcePath:        "submitted_by",
		DefaultOutputPath: "submittedBy",
		Column:            "submitted_by",
		SQLAlias:          "submitted_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"submitted_at": {
		SourcePath:        "submitted_at",
		DefaultOutputPath: "submittedAt",
		Column:            "submitted_at",
		SQLAlias:          "submitted_at",
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

func NewDisputeEvidenceFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = DisputeEvidenceFilterFields[field]
	return
}

type DisputeEvidenceFilterResult struct {
	DisputeEvidence
	FilterCount int `db:"count"`
}

func ValidateDisputeEvidenceFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewDisputeEvidenceFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewDisputeEvidenceFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewDisputeEvidenceFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateDisputeEvidenceFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateDisputeEvidenceFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewDisputeEvidenceFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateDisputeEvidenceFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type DisputeEvidence struct {
	Id             uuid.UUID       `db:"id"`
	DisputeId      uuid.UUID       `db:"dispute_id"`
	EvidenceType   string          `db:"evidence_type"`
	StorageUri     string          `db:"storage_uri"`
	ContentHash    null.String     `db:"content_hash"`
	IdempotencyKey null.String     `db:"idempotency_key"`
	SubmittedBy    uuid.UUID       `db:"submitted_by"`
	SubmittedAt    time.Time       `db:"submitted_at"`
	Metadata       json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type DisputeEvidencePrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d DisputeEvidence) ToDisputeEvidencePrimaryID() DisputeEvidencePrimaryID {
	return DisputeEvidencePrimaryID{
		Id: d.Id,
	}
}

type DisputeEvidenceList []*DisputeEvidence

type DisputeEvidenceFilterResultList []*DisputeEvidenceFilterResult
