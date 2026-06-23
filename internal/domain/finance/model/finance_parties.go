package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinancePartiesDBFieldNameType string

type financePartiesDBFieldName struct {
	Id            FinancePartiesDBFieldNameType
	PartyType     FinancePartiesDBFieldNameType
	ExternalRef   FinancePartiesDBFieldNameType
	LegalName     FinancePartiesDBFieldNameType
	DisplayName   FinancePartiesDBFieldNameType
	CountryCode   FinancePartiesDBFieldNameType
	Email         FinancePartiesDBFieldNameType
	Phone         FinancePartiesDBFieldNameType
	PartyStatus   FinancePartiesDBFieldNameType
	Metadata      FinancePartiesDBFieldNameType
	MetaCreatedAt FinancePartiesDBFieldNameType
	MetaCreatedBy FinancePartiesDBFieldNameType
	MetaUpdatedAt FinancePartiesDBFieldNameType
	MetaUpdatedBy FinancePartiesDBFieldNameType
	MetaDeletedAt FinancePartiesDBFieldNameType
	MetaDeletedBy FinancePartiesDBFieldNameType
}

var FinancePartiesDBFieldName = financePartiesDBFieldName{
	Id:            "id",
	PartyType:     "party_type",
	ExternalRef:   "external_ref",
	LegalName:     "legal_name",
	DisplayName:   "display_name",
	CountryCode:   "country_code",
	Email:         "email",
	Phone:         "phone",
	PartyStatus:   "party_status",
	Metadata:      "metadata",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewFinancePartiesDBFieldNameFromStr(field string) (dbField FinancePartiesDBFieldNameType, found bool) {
	switch field {

	case string(FinancePartiesDBFieldName.Id):
		return FinancePartiesDBFieldName.Id, true

	case string(FinancePartiesDBFieldName.PartyType):
		return FinancePartiesDBFieldName.PartyType, true

	case string(FinancePartiesDBFieldName.ExternalRef):
		return FinancePartiesDBFieldName.ExternalRef, true

	case string(FinancePartiesDBFieldName.LegalName):
		return FinancePartiesDBFieldName.LegalName, true

	case string(FinancePartiesDBFieldName.DisplayName):
		return FinancePartiesDBFieldName.DisplayName, true

	case string(FinancePartiesDBFieldName.CountryCode):
		return FinancePartiesDBFieldName.CountryCode, true

	case string(FinancePartiesDBFieldName.Email):
		return FinancePartiesDBFieldName.Email, true

	case string(FinancePartiesDBFieldName.Phone):
		return FinancePartiesDBFieldName.Phone, true

	case string(FinancePartiesDBFieldName.PartyStatus):
		return FinancePartiesDBFieldName.PartyStatus, true

	case string(FinancePartiesDBFieldName.Metadata):
		return FinancePartiesDBFieldName.Metadata, true

	case string(FinancePartiesDBFieldName.MetaCreatedAt):
		return FinancePartiesDBFieldName.MetaCreatedAt, true

	case string(FinancePartiesDBFieldName.MetaCreatedBy):
		return FinancePartiesDBFieldName.MetaCreatedBy, true

	case string(FinancePartiesDBFieldName.MetaUpdatedAt):
		return FinancePartiesDBFieldName.MetaUpdatedAt, true

	case string(FinancePartiesDBFieldName.MetaUpdatedBy):
		return FinancePartiesDBFieldName.MetaUpdatedBy, true

	case string(FinancePartiesDBFieldName.MetaDeletedAt):
		return FinancePartiesDBFieldName.MetaDeletedAt, true

	case string(FinancePartiesDBFieldName.MetaDeletedBy):
		return FinancePartiesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinancePartiesFilterJoins = map[string]JoinSpec{}

var FinancePartiesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"party_type": {
		SourcePath:        "party_type",
		DefaultOutputPath: "partyType",
		Column:            "party_type",
		SQLAlias:          "party_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"external_ref": {
		SourcePath:        "external_ref",
		DefaultOutputPath: "externalRef",
		Column:            "external_ref",
		SQLAlias:          "external_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"legal_name": {
		SourcePath:        "legal_name",
		DefaultOutputPath: "legalName",
		Column:            "legal_name",
		SQLAlias:          "legal_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"display_name": {
		SourcePath:        "display_name",
		DefaultOutputPath: "displayName",
		Column:            "display_name",
		SQLAlias:          "display_name",
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
	"email": {
		SourcePath:        "email",
		DefaultOutputPath: "email",
		Column:            "email",
		SQLAlias:          "email",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"phone": {
		SourcePath:        "phone",
		DefaultOutputPath: "phone",
		Column:            "phone",
		SQLAlias:          "phone",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"party_status": {
		SourcePath:        "party_status",
		DefaultOutputPath: "partyStatus",
		Column:            "party_status",
		SQLAlias:          "party_status",
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

func NewFinancePartiesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinancePartiesFilterFields[field]
	return
}

type FinancePartiesFilterResult struct {
	FinanceParties
	FilterCount int `db:"count"`
}

func ValidateFinancePartiesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinancePartiesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinancePartiesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinancePartiesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinancePartiesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinancePartiesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinancePartiesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinancePartiesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PartyStatus string

const (
	PartyStatusActive   PartyStatus = "active"
	PartyStatusInactive PartyStatus = "inactive"
	PartyStatusBlocked  PartyStatus = "blocked"
)

type PartyType string

const (
	PartyTypePlatform     PartyType = "platform"
	PartyTypeMerchant     PartyType = "merchant"
	PartyTypeCustomer     PartyType = "customer"
	PartyTypeProvider     PartyType = "provider"
	PartyTypeBank         PartyType = "bank"
	PartyTypeTaxAuthority PartyType = "tax_authority"
	PartyTypeAffiliate    PartyType = "affiliate"
	PartyTypeGuide        PartyType = "guide"
	PartyTypeDriver       PartyType = "driver"
	PartyTypeManualOps    PartyType = "manual_ops"
)

type FinanceParties struct {
	Id          uuid.UUID       `db:"id"`
	PartyType   PartyType       `db:"party_type"`
	ExternalRef null.String     `db:"external_ref"`
	LegalName   string          `db:"legal_name"`
	DisplayName null.String     `db:"display_name"`
	CountryCode string          `db:"country_code"`
	Email       null.String     `db:"email"`
	Phone       null.String     `db:"phone"`
	PartyStatus PartyStatus     `db:"party_status"`
	Metadata    json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinancePartiesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceParties) ToFinancePartiesPrimaryID() FinancePartiesPrimaryID {
	return FinancePartiesPrimaryID{
		Id: d.Id,
	}
}

type FinancePartiesList []*FinanceParties

type FinancePartiesFilterResultList []*FinancePartiesFilterResult
