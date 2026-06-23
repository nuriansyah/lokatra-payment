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

type TaxDocumentArtifactsDBFieldNameType string

type taxDocumentArtifactsDBFieldName struct {
	Id             TaxDocumentArtifactsDBFieldNameType
	DocumentType   TaxDocumentArtifactsDBFieldNameType
	DocumentId     TaxDocumentArtifactsDBFieldNameType
	ArtifactType   TaxDocumentArtifactsDBFieldNameType
	StorageUri     TaxDocumentArtifactsDBFieldNameType
	ContentHash    TaxDocumentArtifactsDBFieldNameType
	IdempotencyKey TaxDocumentArtifactsDBFieldNameType
	GeneratedAt    TaxDocumentArtifactsDBFieldNameType
	Metadata       TaxDocumentArtifactsDBFieldNameType
	MetaCreatedAt  TaxDocumentArtifactsDBFieldNameType
	MetaCreatedBy  TaxDocumentArtifactsDBFieldNameType
	MetaUpdatedAt  TaxDocumentArtifactsDBFieldNameType
	MetaUpdatedBy  TaxDocumentArtifactsDBFieldNameType
	MetaDeletedAt  TaxDocumentArtifactsDBFieldNameType
	MetaDeletedBy  TaxDocumentArtifactsDBFieldNameType
}

var TaxDocumentArtifactsDBFieldName = taxDocumentArtifactsDBFieldName{
	Id:             "id",
	DocumentType:   "document_type",
	DocumentId:     "document_id",
	ArtifactType:   "artifact_type",
	StorageUri:     "storage_uri",
	ContentHash:    "content_hash",
	IdempotencyKey: "idempotency_key",
	GeneratedAt:    "generated_at",
	Metadata:       "metadata",
	MetaCreatedAt:  "meta_created_at",
	MetaCreatedBy:  "meta_created_by",
	MetaUpdatedAt:  "meta_updated_at",
	MetaUpdatedBy:  "meta_updated_by",
	MetaDeletedAt:  "meta_deleted_at",
	MetaDeletedBy:  "meta_deleted_by",
}

func NewTaxDocumentArtifactsDBFieldNameFromStr(field string) (dbField TaxDocumentArtifactsDBFieldNameType, found bool) {
	switch field {

	case string(TaxDocumentArtifactsDBFieldName.Id):
		return TaxDocumentArtifactsDBFieldName.Id, true

	case string(TaxDocumentArtifactsDBFieldName.DocumentType):
		return TaxDocumentArtifactsDBFieldName.DocumentType, true

	case string(TaxDocumentArtifactsDBFieldName.DocumentId):
		return TaxDocumentArtifactsDBFieldName.DocumentId, true

	case string(TaxDocumentArtifactsDBFieldName.ArtifactType):
		return TaxDocumentArtifactsDBFieldName.ArtifactType, true

	case string(TaxDocumentArtifactsDBFieldName.StorageUri):
		return TaxDocumentArtifactsDBFieldName.StorageUri, true

	case string(TaxDocumentArtifactsDBFieldName.ContentHash):
		return TaxDocumentArtifactsDBFieldName.ContentHash, true

	case string(TaxDocumentArtifactsDBFieldName.IdempotencyKey):
		return TaxDocumentArtifactsDBFieldName.IdempotencyKey, true

	case string(TaxDocumentArtifactsDBFieldName.GeneratedAt):
		return TaxDocumentArtifactsDBFieldName.GeneratedAt, true

	case string(TaxDocumentArtifactsDBFieldName.Metadata):
		return TaxDocumentArtifactsDBFieldName.Metadata, true

	case string(TaxDocumentArtifactsDBFieldName.MetaCreatedAt):
		return TaxDocumentArtifactsDBFieldName.MetaCreatedAt, true

	case string(TaxDocumentArtifactsDBFieldName.MetaCreatedBy):
		return TaxDocumentArtifactsDBFieldName.MetaCreatedBy, true

	case string(TaxDocumentArtifactsDBFieldName.MetaUpdatedAt):
		return TaxDocumentArtifactsDBFieldName.MetaUpdatedAt, true

	case string(TaxDocumentArtifactsDBFieldName.MetaUpdatedBy):
		return TaxDocumentArtifactsDBFieldName.MetaUpdatedBy, true

	case string(TaxDocumentArtifactsDBFieldName.MetaDeletedAt):
		return TaxDocumentArtifactsDBFieldName.MetaDeletedAt, true

	case string(TaxDocumentArtifactsDBFieldName.MetaDeletedBy):
		return TaxDocumentArtifactsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var TaxDocumentArtifactsFilterJoins = map[string]JoinSpec{}

var TaxDocumentArtifactsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"document_type": {
		SourcePath:        "document_type",
		DefaultOutputPath: "documentType",
		Column:            "document_type",
		SQLAlias:          "document_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"document_id": {
		SourcePath:        "document_id",
		DefaultOutputPath: "documentId",
		Column:            "document_id",
		SQLAlias:          "document_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"artifact_type": {
		SourcePath:        "artifact_type",
		DefaultOutputPath: "artifactType",
		Column:            "artifact_type",
		SQLAlias:          "artifact_type",
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
	"generated_at": {
		SourcePath:        "generated_at",
		DefaultOutputPath: "generatedAt",
		Column:            "generated_at",
		SQLAlias:          "generated_at",
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

func NewTaxDocumentArtifactsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = TaxDocumentArtifactsFilterFields[field]
	return
}

type TaxDocumentArtifactsFilterResult struct {
	TaxDocumentArtifacts
	FilterCount int `db:"count"`
}

func ValidateTaxDocumentArtifactsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewTaxDocumentArtifactsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewTaxDocumentArtifactsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewTaxDocumentArtifactsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateTaxDocumentArtifactsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateTaxDocumentArtifactsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewTaxDocumentArtifactsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateTaxDocumentArtifactsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type DocumentType string

const (
	DocumentTypeTaxInvoice DocumentType = "tax_invoice"
	DocumentTypeCreditNote DocumentType = "credit_note"
)

type TaxDocumentArtifacts struct {
	Id             uuid.UUID       `db:"id"`
	DocumentType   DocumentType    `db:"document_type"`
	DocumentId     uuid.UUID       `db:"document_id"`
	ArtifactType   string          `db:"artifact_type"`
	StorageUri     string          `db:"storage_uri"`
	ContentHash    null.String     `db:"content_hash"`
	IdempotencyKey null.String     `db:"idempotency_key"`
	GeneratedAt    time.Time       `db:"generated_at"`
	Metadata       json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type TaxDocumentArtifactsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d TaxDocumentArtifacts) ToTaxDocumentArtifactsPrimaryID() TaxDocumentArtifactsPrimaryID {
	return TaxDocumentArtifactsPrimaryID{
		Id: d.Id,
	}
}

type TaxDocumentArtifactsList []*TaxDocumentArtifacts

type TaxDocumentArtifactsFilterResultList []*TaxDocumentArtifactsFilterResult
