package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderWebhookEndpointsDBFieldNameType string

type providerWebhookEndpointsDBFieldName struct {
	Id                 ProviderWebhookEndpointsDBFieldNameType
	ProviderAccountId  ProviderWebhookEndpointsDBFieldNameType
	ProviderCode       ProviderWebhookEndpointsDBFieldNameType
	EndpointKey        ProviderWebhookEndpointsDBFieldNameType
	Environment        ProviderWebhookEndpointsDBFieldNameType
	SecretRef          ProviderWebhookEndpointsDBFieldNameType
	SignatureAlgorithm ProviderWebhookEndpointsDBFieldNameType
	IsActive           ProviderWebhookEndpointsDBFieldNameType
	Metadata           ProviderWebhookEndpointsDBFieldNameType
	MetaCreatedAt      ProviderWebhookEndpointsDBFieldNameType
	MetaCreatedBy      ProviderWebhookEndpointsDBFieldNameType
	MetaUpdatedAt      ProviderWebhookEndpointsDBFieldNameType
	MetaUpdatedBy      ProviderWebhookEndpointsDBFieldNameType
	MetaDeletedAt      ProviderWebhookEndpointsDBFieldNameType
	MetaDeletedBy      ProviderWebhookEndpointsDBFieldNameType
}

var ProviderWebhookEndpointsDBFieldName = providerWebhookEndpointsDBFieldName{
	Id:                 "id",
	ProviderAccountId:  "provider_account_id",
	ProviderCode:       "provider_code",
	EndpointKey:        "endpoint_key",
	Environment:        "environment",
	SecretRef:          "secret_ref",
	SignatureAlgorithm: "signature_algorithm",
	IsActive:           "is_active",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewProviderWebhookEndpointsDBFieldNameFromStr(field string) (dbField ProviderWebhookEndpointsDBFieldNameType, found bool) {
	switch field {

	case string(ProviderWebhookEndpointsDBFieldName.Id):
		return ProviderWebhookEndpointsDBFieldName.Id, true

	case string(ProviderWebhookEndpointsDBFieldName.ProviderAccountId):
		return ProviderWebhookEndpointsDBFieldName.ProviderAccountId, true

	case string(ProviderWebhookEndpointsDBFieldName.ProviderCode):
		return ProviderWebhookEndpointsDBFieldName.ProviderCode, true

	case string(ProviderWebhookEndpointsDBFieldName.EndpointKey):
		return ProviderWebhookEndpointsDBFieldName.EndpointKey, true

	case string(ProviderWebhookEndpointsDBFieldName.Environment):
		return ProviderWebhookEndpointsDBFieldName.Environment, true

	case string(ProviderWebhookEndpointsDBFieldName.SecretRef):
		return ProviderWebhookEndpointsDBFieldName.SecretRef, true

	case string(ProviderWebhookEndpointsDBFieldName.SignatureAlgorithm):
		return ProviderWebhookEndpointsDBFieldName.SignatureAlgorithm, true

	case string(ProviderWebhookEndpointsDBFieldName.IsActive):
		return ProviderWebhookEndpointsDBFieldName.IsActive, true

	case string(ProviderWebhookEndpointsDBFieldName.Metadata):
		return ProviderWebhookEndpointsDBFieldName.Metadata, true

	case string(ProviderWebhookEndpointsDBFieldName.MetaCreatedAt):
		return ProviderWebhookEndpointsDBFieldName.MetaCreatedAt, true

	case string(ProviderWebhookEndpointsDBFieldName.MetaCreatedBy):
		return ProviderWebhookEndpointsDBFieldName.MetaCreatedBy, true

	case string(ProviderWebhookEndpointsDBFieldName.MetaUpdatedAt):
		return ProviderWebhookEndpointsDBFieldName.MetaUpdatedAt, true

	case string(ProviderWebhookEndpointsDBFieldName.MetaUpdatedBy):
		return ProviderWebhookEndpointsDBFieldName.MetaUpdatedBy, true

	case string(ProviderWebhookEndpointsDBFieldName.MetaDeletedAt):
		return ProviderWebhookEndpointsDBFieldName.MetaDeletedAt, true

	case string(ProviderWebhookEndpointsDBFieldName.MetaDeletedBy):
		return ProviderWebhookEndpointsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderWebhookEndpointsFilterJoins = map[string]JoinSpec{}

var ProviderWebhookEndpointsFilterFields = map[string]FilterFieldSpec{
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
	"provider_code": {
		SourcePath:        "provider_code",
		DefaultOutputPath: "providerCode",
		Column:            "provider_code",
		SQLAlias:          "provider_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"endpoint_key": {
		SourcePath:        "endpoint_key",
		DefaultOutputPath: "endpointKey",
		Column:            "endpoint_key",
		SQLAlias:          "endpoint_key",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"environment": {
		SourcePath:        "environment",
		DefaultOutputPath: "environment",
		Column:            "environment",
		SQLAlias:          "environment",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"secret_ref": {
		SourcePath:        "secret_ref",
		DefaultOutputPath: "secretRef",
		Column:            "secret_ref",
		SQLAlias:          "secret_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"signature_algorithm": {
		SourcePath:        "signature_algorithm",
		DefaultOutputPath: "signatureAlgorithm",
		Column:            "signature_algorithm",
		SQLAlias:          "signature_algorithm",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_active": {
		SourcePath:        "is_active",
		DefaultOutputPath: "isActive",
		Column:            "is_active",
		SQLAlias:          "is_active",
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

func NewProviderWebhookEndpointsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderWebhookEndpointsFilterFields[field]
	return
}

type ProviderWebhookEndpointsFilterResult struct {
	ProviderWebhookEndpoints
	FilterCount int `db:"count"`
}

func ValidateProviderWebhookEndpointsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderWebhookEndpointsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderWebhookEndpointsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderWebhookEndpointsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderWebhookEndpointsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderWebhookEndpointsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderWebhookEndpointsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderWebhookEndpointsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ProviderWebhookEndpoints struct {
	Id                 uuid.UUID       `db:"id"`
	ProviderAccountId  uuid.UUID       `db:"provider_account_id"`
	ProviderCode       string          `db:"provider_code"`
	EndpointKey        string          `db:"endpoint_key"`
	Environment        string          `db:"environment"`
	SecretRef          string          `db:"secret_ref"`
	SignatureAlgorithm string          `db:"signature_algorithm"`
	IsActive           bool            `db:"is_active"`
	Metadata           json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ProviderWebhookEndpointsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderWebhookEndpoints) ToProviderWebhookEndpointsPrimaryID() ProviderWebhookEndpointsPrimaryID {
	return ProviderWebhookEndpointsPrimaryID{
		Id: d.Id,
	}
}

type ProviderWebhookEndpointsList []*ProviderWebhookEndpoints

type ProviderWebhookEndpointsFilterResultList []*ProviderWebhookEndpointsFilterResult
