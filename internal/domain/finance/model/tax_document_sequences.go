package model

import (
	"encoding/json"
	"fmt"

	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type TaxDocumentSequencesDBFieldNameType string

type taxDocumentSequencesDBFieldName struct {
	SequenceCode  TaxDocumentSequencesDBFieldNameType
	CurrentValue  TaxDocumentSequencesDBFieldNameType
	Metadata      TaxDocumentSequencesDBFieldNameType
	MetaCreatedAt TaxDocumentSequencesDBFieldNameType
	MetaCreatedBy TaxDocumentSequencesDBFieldNameType
	MetaUpdatedAt TaxDocumentSequencesDBFieldNameType
	MetaUpdatedBy TaxDocumentSequencesDBFieldNameType
	MetaDeletedAt TaxDocumentSequencesDBFieldNameType
	MetaDeletedBy TaxDocumentSequencesDBFieldNameType
}

var TaxDocumentSequencesDBFieldName = taxDocumentSequencesDBFieldName{
	SequenceCode:  "sequence_code",
	CurrentValue:  "current_value",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewTaxDocumentSequencesDBFieldNameFromStr(field string) (dbField TaxDocumentSequencesDBFieldNameType, found bool) {
	switch field {

	case string(TaxDocumentSequencesDBFieldName.SequenceCode):
		return TaxDocumentSequencesDBFieldName.SequenceCode, true

	case string(TaxDocumentSequencesDBFieldName.CurrentValue):
		return TaxDocumentSequencesDBFieldName.CurrentValue, true

	case string(TaxDocumentSequencesDBFieldName.Metadata):
		return TaxDocumentSequencesDBFieldName.Metadata, true

	case string(TaxDocumentSequencesDBFieldName.MetaCreatedAt):
		return TaxDocumentSequencesDBFieldName.MetaCreatedAt, true

	case string(TaxDocumentSequencesDBFieldName.MetaCreatedBy):
		return TaxDocumentSequencesDBFieldName.MetaCreatedBy, true

	case string(TaxDocumentSequencesDBFieldName.MetaUpdatedAt):
		return TaxDocumentSequencesDBFieldName.MetaUpdatedAt, true

	case string(TaxDocumentSequencesDBFieldName.MetaUpdatedBy):
		return TaxDocumentSequencesDBFieldName.MetaUpdatedBy, true

	case string(TaxDocumentSequencesDBFieldName.MetaDeletedAt):
		return TaxDocumentSequencesDBFieldName.MetaDeletedAt, true

	case string(TaxDocumentSequencesDBFieldName.MetaDeletedBy):
		return TaxDocumentSequencesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var TaxDocumentSequencesFilterJoins = map[string]JoinSpec{}

var TaxDocumentSequencesFilterFields = map[string]FilterFieldSpec{
	"sequence_code": {
		SourcePath:        "sequence_code",
		DefaultOutputPath: "sequenceCode",
		Column:            "sequence_code",
		SQLAlias:          "sequence_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"current_value": {
		SourcePath:        "current_value",
		DefaultOutputPath: "currentValue",
		Column:            "current_value",
		SQLAlias:          "current_value",
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

func NewTaxDocumentSequencesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = TaxDocumentSequencesFilterFields[field]
	return
}

type TaxDocumentSequencesFilterResult struct {
	TaxDocumentSequences
	FilterCount int `db:"count"`
}

func ValidateTaxDocumentSequencesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewTaxDocumentSequencesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewTaxDocumentSequencesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewTaxDocumentSequencesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateTaxDocumentSequencesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateTaxDocumentSequencesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewTaxDocumentSequencesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateTaxDocumentSequencesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type TaxDocumentSequences struct {
	SequenceCode string          `db:"sequence_code"`
	CurrentValue int64           `db:"current_value"`
	Metadata     json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type TaxDocumentSequencesPrimaryID struct {
	SequenceCode string `db:"sequence_code"`
}

func (d TaxDocumentSequences) ToTaxDocumentSequencesPrimaryID() TaxDocumentSequencesPrimaryID {
	return TaxDocumentSequencesPrimaryID{
		SequenceCode: d.SequenceCode,
	}
}

type TaxDocumentSequencesList []*TaxDocumentSequences

type TaxDocumentSequencesFilterResultList []*TaxDocumentSequencesFilterResult
