package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type TaxProfilesDBFieldNameType string

type taxProfilesDBFieldName struct {
	Id                   TaxProfilesDBFieldNameType
	OwnerPartyId         TaxProfilesDBFieldNameType
	CountryCode          TaxProfilesDBFieldNameType
	TaxResidencyCountry  TaxProfilesDBFieldNameType
	TaxIdMasked          TaxProfilesDBFieldNameType
	TaxEntityType        TaxProfilesDBFieldNameType
	IsVatRegistered      TaxProfilesDBFieldNameType
	IsWithholdingSubject TaxProfilesDBFieldNameType
	ProfileStatus        TaxProfilesDBFieldNameType
	Metadata             TaxProfilesDBFieldNameType
	MetaCreatedAt        TaxProfilesDBFieldNameType
	MetaCreatedBy        TaxProfilesDBFieldNameType
	MetaUpdatedAt        TaxProfilesDBFieldNameType
	MetaUpdatedBy        TaxProfilesDBFieldNameType
	MetaDeletedAt        TaxProfilesDBFieldNameType
	MetaDeletedBy        TaxProfilesDBFieldNameType
}

var TaxProfilesDBFieldName = taxProfilesDBFieldName{
	Id:                   "id",
	OwnerPartyId:         "owner_party_id",
	CountryCode:          "country_code",
	TaxResidencyCountry:  "tax_residency_country",
	TaxIdMasked:          "tax_id_masked",
	TaxEntityType:        "tax_entity_type",
	IsVatRegistered:      "is_vat_registered",
	IsWithholdingSubject: "is_withholding_subject",
	ProfileStatus:        "profile_status",
	Metadata:             "metadata",
	MetaCreatedAt:        "meta_created_at",
	MetaCreatedBy:        "meta_created_by",
	MetaUpdatedAt:        "meta_updated_at",
	MetaUpdatedBy:        "meta_updated_by",
	MetaDeletedAt:        "meta_deleted_at",
	MetaDeletedBy:        "meta_deleted_by",
}

func NewTaxProfilesDBFieldNameFromStr(field string) (dbField TaxProfilesDBFieldNameType, found bool) {
	switch field {

	case string(TaxProfilesDBFieldName.Id):
		return TaxProfilesDBFieldName.Id, true

	case string(TaxProfilesDBFieldName.OwnerPartyId):
		return TaxProfilesDBFieldName.OwnerPartyId, true

	case string(TaxProfilesDBFieldName.CountryCode):
		return TaxProfilesDBFieldName.CountryCode, true

	case string(TaxProfilesDBFieldName.TaxResidencyCountry):
		return TaxProfilesDBFieldName.TaxResidencyCountry, true

	case string(TaxProfilesDBFieldName.TaxIdMasked):
		return TaxProfilesDBFieldName.TaxIdMasked, true

	case string(TaxProfilesDBFieldName.TaxEntityType):
		return TaxProfilesDBFieldName.TaxEntityType, true

	case string(TaxProfilesDBFieldName.IsVatRegistered):
		return TaxProfilesDBFieldName.IsVatRegistered, true

	case string(TaxProfilesDBFieldName.IsWithholdingSubject):
		return TaxProfilesDBFieldName.IsWithholdingSubject, true

	case string(TaxProfilesDBFieldName.ProfileStatus):
		return TaxProfilesDBFieldName.ProfileStatus, true

	case string(TaxProfilesDBFieldName.Metadata):
		return TaxProfilesDBFieldName.Metadata, true

	case string(TaxProfilesDBFieldName.MetaCreatedAt):
		return TaxProfilesDBFieldName.MetaCreatedAt, true

	case string(TaxProfilesDBFieldName.MetaCreatedBy):
		return TaxProfilesDBFieldName.MetaCreatedBy, true

	case string(TaxProfilesDBFieldName.MetaUpdatedAt):
		return TaxProfilesDBFieldName.MetaUpdatedAt, true

	case string(TaxProfilesDBFieldName.MetaUpdatedBy):
		return TaxProfilesDBFieldName.MetaUpdatedBy, true

	case string(TaxProfilesDBFieldName.MetaDeletedAt):
		return TaxProfilesDBFieldName.MetaDeletedAt, true

	case string(TaxProfilesDBFieldName.MetaDeletedBy):
		return TaxProfilesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var TaxProfilesFilterJoins = map[string]JoinSpec{}

var TaxProfilesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"owner_party_id": {
		SourcePath:        "owner_party_id",
		DefaultOutputPath: "ownerPartyId",
		Column:            "owner_party_id",
		SQLAlias:          "owner_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"country_code": {
		SourcePath:        "country_code",
		DefaultOutputPath: "countryCode",
		Column:            "country_code",
		SQLAlias:          "country_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_residency_country": {
		SourcePath:        "tax_residency_country",
		DefaultOutputPath: "taxResidencyCountry",
		Column:            "tax_residency_country",
		SQLAlias:          "tax_residency_country",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_id_masked": {
		SourcePath:        "tax_id_masked",
		DefaultOutputPath: "taxIdMasked",
		Column:            "tax_id_masked",
		SQLAlias:          "tax_id_masked",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_entity_type": {
		SourcePath:        "tax_entity_type",
		DefaultOutputPath: "taxEntityType",
		Column:            "tax_entity_type",
		SQLAlias:          "tax_entity_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_vat_registered": {
		SourcePath:        "is_vat_registered",
		DefaultOutputPath: "isVatRegistered",
		Column:            "is_vat_registered",
		SQLAlias:          "is_vat_registered",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_withholding_subject": {
		SourcePath:        "is_withholding_subject",
		DefaultOutputPath: "isWithholdingSubject",
		Column:            "is_withholding_subject",
		SQLAlias:          "is_withholding_subject",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"profile_status": {
		SourcePath:        "profile_status",
		DefaultOutputPath: "profileStatus",
		Column:            "profile_status",
		SQLAlias:          "profile_status",
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

func NewTaxProfilesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = TaxProfilesFilterFields[field]
	return
}

type TaxProfilesFilterResult struct {
	TaxProfiles
	FilterCount int `db:"count"`
}

func ValidateTaxProfilesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewTaxProfilesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewTaxProfilesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewTaxProfilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateTaxProfilesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateTaxProfilesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewTaxProfilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateTaxProfilesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ProfileStatus string

const (
	ProfileStatusActive   ProfileStatus = "active"
	ProfileStatusInactive ProfileStatus = "inactive"
)

type TaxProfiles struct {
	Id                   uuid.UUID       `db:"id"`
	OwnerPartyId         uuid.UUID       `db:"owner_party_id"`
	CountryCode          string          `db:"country_code"`
	TaxResidencyCountry  null.String     `db:"tax_residency_country"`
	TaxIdMasked          null.String     `db:"tax_id_masked"`
	TaxEntityType        null.String     `db:"tax_entity_type"`
	IsVatRegistered      bool            `db:"is_vat_registered"`
	IsWithholdingSubject bool            `db:"is_withholding_subject"`
	ProfileStatus        ProfileStatus   `db:"profile_status"`
	Metadata             json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type TaxProfilesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d TaxProfiles) ToTaxProfilesPrimaryID() TaxProfilesPrimaryID {
	return TaxProfilesPrimaryID{
		Id: d.Id,
	}
}

type TaxProfilesList []*TaxProfiles

type TaxProfilesFilterResultList []*TaxProfilesFilterResult
