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

type ProviderStatementFilesDBFieldNameType string

type providerStatementFilesDBFieldName struct {
	Id                ProviderStatementFilesDBFieldNameType
	ProviderAccountId ProviderStatementFilesDBFieldNameType
	StatementType     ProviderStatementFilesDBFieldNameType
	StatementDate     ProviderStatementFilesDBFieldNameType
	FileName          ProviderStatementFilesDBFieldNameType
	StorageUrl        ProviderStatementFilesDBFieldNameType
	FileHash          ProviderStatementFilesDBFieldNameType
	ParseStatus       ProviderStatementFilesDBFieldNameType
	Metadata          ProviderStatementFilesDBFieldNameType
	MetaCreatedAt     ProviderStatementFilesDBFieldNameType
	MetaCreatedBy     ProviderStatementFilesDBFieldNameType
	MetaUpdatedAt     ProviderStatementFilesDBFieldNameType
	MetaUpdatedBy     ProviderStatementFilesDBFieldNameType
	MetaDeletedAt     ProviderStatementFilesDBFieldNameType
	MetaDeletedBy     ProviderStatementFilesDBFieldNameType
}

var ProviderStatementFilesDBFieldName = providerStatementFilesDBFieldName{
	Id:                "id",
	ProviderAccountId: "provider_account_id",
	StatementType:     "statement_type",
	StatementDate:     "statement_date",
	FileName:          "file_name",
	StorageUrl:        "storage_url",
	FileHash:          "file_hash",
	ParseStatus:       "parse_status",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewProviderStatementFilesDBFieldNameFromStr(field string) (dbField ProviderStatementFilesDBFieldNameType, found bool) {
	switch field {

	case string(ProviderStatementFilesDBFieldName.Id):
		return ProviderStatementFilesDBFieldName.Id, true

	case string(ProviderStatementFilesDBFieldName.ProviderAccountId):
		return ProviderStatementFilesDBFieldName.ProviderAccountId, true

	case string(ProviderStatementFilesDBFieldName.StatementType):
		return ProviderStatementFilesDBFieldName.StatementType, true

	case string(ProviderStatementFilesDBFieldName.StatementDate):
		return ProviderStatementFilesDBFieldName.StatementDate, true

	case string(ProviderStatementFilesDBFieldName.FileName):
		return ProviderStatementFilesDBFieldName.FileName, true

	case string(ProviderStatementFilesDBFieldName.StorageUrl):
		return ProviderStatementFilesDBFieldName.StorageUrl, true

	case string(ProviderStatementFilesDBFieldName.FileHash):
		return ProviderStatementFilesDBFieldName.FileHash, true

	case string(ProviderStatementFilesDBFieldName.ParseStatus):
		return ProviderStatementFilesDBFieldName.ParseStatus, true

	case string(ProviderStatementFilesDBFieldName.Metadata):
		return ProviderStatementFilesDBFieldName.Metadata, true

	case string(ProviderStatementFilesDBFieldName.MetaCreatedAt):
		return ProviderStatementFilesDBFieldName.MetaCreatedAt, true

	case string(ProviderStatementFilesDBFieldName.MetaCreatedBy):
		return ProviderStatementFilesDBFieldName.MetaCreatedBy, true

	case string(ProviderStatementFilesDBFieldName.MetaUpdatedAt):
		return ProviderStatementFilesDBFieldName.MetaUpdatedAt, true

	case string(ProviderStatementFilesDBFieldName.MetaUpdatedBy):
		return ProviderStatementFilesDBFieldName.MetaUpdatedBy, true

	case string(ProviderStatementFilesDBFieldName.MetaDeletedAt):
		return ProviderStatementFilesDBFieldName.MetaDeletedAt, true

	case string(ProviderStatementFilesDBFieldName.MetaDeletedBy):
		return ProviderStatementFilesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderStatementFilesFilterJoins = map[string]JoinSpec{}

var ProviderStatementFilesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_account_id": {
		SourcePath:        "provider_account_id",
		DefaultOutputPath: "providerAccountId",
		Column:            "provider_account_id",
		SQLAlias:          "provider_account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"statement_type": {
		SourcePath:        "statement_type",
		DefaultOutputPath: "statementType",
		Column:            "statement_type",
		SQLAlias:          "statement_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"statement_date": {
		SourcePath:        "statement_date",
		DefaultOutputPath: "statementDate",
		Column:            "statement_date",
		SQLAlias:          "statement_date",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"file_name": {
		SourcePath:        "file_name",
		DefaultOutputPath: "fileName",
		Column:            "file_name",
		SQLAlias:          "file_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"storage_url": {
		SourcePath:        "storage_url",
		DefaultOutputPath: "storageUrl",
		Column:            "storage_url",
		SQLAlias:          "storage_url",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"file_hash": {
		SourcePath:        "file_hash",
		DefaultOutputPath: "fileHash",
		Column:            "file_hash",
		SQLAlias:          "file_hash",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"parse_status": {
		SourcePath:        "parse_status",
		DefaultOutputPath: "parseStatus",
		Column:            "parse_status",
		SQLAlias:          "parse_status",
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

func NewProviderStatementFilesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderStatementFilesFilterFields[field]
	return
}

type ProviderStatementFilesFilterResult struct {
	ProviderStatementFiles
	FilterCount int `db:"count"`
}

func ValidateProviderStatementFilesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderStatementFilesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderStatementFilesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderStatementFilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderStatementFilesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderStatementFilesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderStatementFilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderStatementFilesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ParseStatus string

const (
	ParseStatusUploaded    ParseStatus = "uploaded"
	ParseStatusParsed      ParseStatus = "parsed"
	ParseStatusFailed      ParseStatus = "failed"
	ParseStatusReprocessed ParseStatus = "reprocessed"
)

type StatementType string

const (
	StatementTypeCollection StatementType = "collection"
	StatementTypeRefund     StatementType = "refund"
	StatementTypePayout     StatementType = "payout"
	StatementTypeBalance    StatementType = "balance"
	StatementTypeChargeback StatementType = "chargeback"
)

type ProviderStatementFiles struct {
	Id                uuid.UUID       `db:"id"`
	ProviderAccountId uuid.UUID       `db:"provider_account_id"`
	StatementType     StatementType   `db:"statement_type"`
	StatementDate     time.Time       `db:"statement_date"`
	FileName          string          `db:"file_name"`
	StorageUrl        null.String     `db:"storage_url"`
	FileHash          string          `db:"file_hash"`
	ParseStatus       ParseStatus     `db:"parse_status"`
	Metadata          json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ProviderStatementFilesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderStatementFiles) ToProviderStatementFilesPrimaryID() ProviderStatementFilesPrimaryID {
	return ProviderStatementFilesPrimaryID{
		Id: d.Id,
	}
}

type ProviderStatementFilesList []*ProviderStatementFiles

type ProviderStatementFilesFilterResultList []*ProviderStatementFilesFilterResult
