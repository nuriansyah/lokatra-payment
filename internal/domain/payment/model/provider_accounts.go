package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type ProviderAccountsDBFieldNameType string

type providerAccountsDBFieldName struct {
	Id                  ProviderAccountsDBFieldNameType
	ProviderId          ProviderAccountsDBFieldNameType
	AccountName         ProviderAccountsDBFieldNameType
	Environment         ProviderAccountsDBFieldNameType
	OwnerType           ProviderAccountsDBFieldNameType
	OwnerId             ProviderAccountsDBFieldNameType
	MerchantRef         ProviderAccountsDBFieldNameType
	CredentialSecretRef ProviderAccountsDBFieldNameType
	WebhookSecretRef    ProviderAccountsDBFieldNameType
	PublicKeyRef        ProviderAccountsDBFieldNameType
	Status              ProviderAccountsDBFieldNameType
	Config              ProviderAccountsDBFieldNameType
	Metadata            ProviderAccountsDBFieldNameType
	MetaCreatedAt       ProviderAccountsDBFieldNameType
	MetaCreatedBy       ProviderAccountsDBFieldNameType
	MetaUpdatedAt       ProviderAccountsDBFieldNameType
	MetaUpdatedBy       ProviderAccountsDBFieldNameType
	MetaDeletedAt       ProviderAccountsDBFieldNameType
	MetaDeletedBy       ProviderAccountsDBFieldNameType
}

var ProviderAccountsDBFieldName = providerAccountsDBFieldName{
	Id:                  "id",
	ProviderId:          "provider_id",
	AccountName:         "account_name",
	Environment:         "environment",
	OwnerType:           "owner_type",
	OwnerId:             "owner_id",
	MerchantRef:         "merchant_ref",
	CredentialSecretRef: "credential_secret_ref",
	WebhookSecretRef:    "webhook_secret_ref",
	PublicKeyRef:        "public_key_ref",
	Status:              "status",
	Config:              "config",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewProviderAccountsDBFieldNameFromStr(field string) (dbField ProviderAccountsDBFieldNameType, found bool) {
	switch field {

	case string(ProviderAccountsDBFieldName.Id):
		return ProviderAccountsDBFieldName.Id, true

	case string(ProviderAccountsDBFieldName.ProviderId):
		return ProviderAccountsDBFieldName.ProviderId, true

	case string(ProviderAccountsDBFieldName.AccountName):
		return ProviderAccountsDBFieldName.AccountName, true

	case string(ProviderAccountsDBFieldName.Environment):
		return ProviderAccountsDBFieldName.Environment, true

	case string(ProviderAccountsDBFieldName.OwnerType):
		return ProviderAccountsDBFieldName.OwnerType, true

	case string(ProviderAccountsDBFieldName.OwnerId):
		return ProviderAccountsDBFieldName.OwnerId, true

	case string(ProviderAccountsDBFieldName.MerchantRef):
		return ProviderAccountsDBFieldName.MerchantRef, true

	case string(ProviderAccountsDBFieldName.CredentialSecretRef):
		return ProviderAccountsDBFieldName.CredentialSecretRef, true

	case string(ProviderAccountsDBFieldName.WebhookSecretRef):
		return ProviderAccountsDBFieldName.WebhookSecretRef, true

	case string(ProviderAccountsDBFieldName.PublicKeyRef):
		return ProviderAccountsDBFieldName.PublicKeyRef, true

	case string(ProviderAccountsDBFieldName.Status):
		return ProviderAccountsDBFieldName.Status, true

	case string(ProviderAccountsDBFieldName.Config):
		return ProviderAccountsDBFieldName.Config, true

	case string(ProviderAccountsDBFieldName.Metadata):
		return ProviderAccountsDBFieldName.Metadata, true

	case string(ProviderAccountsDBFieldName.MetaCreatedAt):
		return ProviderAccountsDBFieldName.MetaCreatedAt, true

	case string(ProviderAccountsDBFieldName.MetaCreatedBy):
		return ProviderAccountsDBFieldName.MetaCreatedBy, true

	case string(ProviderAccountsDBFieldName.MetaUpdatedAt):
		return ProviderAccountsDBFieldName.MetaUpdatedAt, true

	case string(ProviderAccountsDBFieldName.MetaUpdatedBy):
		return ProviderAccountsDBFieldName.MetaUpdatedBy, true

	case string(ProviderAccountsDBFieldName.MetaDeletedAt):
		return ProviderAccountsDBFieldName.MetaDeletedAt, true

	case string(ProviderAccountsDBFieldName.MetaDeletedBy):
		return ProviderAccountsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderAccountsFilterJoins = map[string]JoinSpec{}

var ProviderAccountsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_id": {
		SourcePath:        "provider_id",
		DefaultOutputPath: "providerId",
		Column:            "provider_id",
		SQLAlias:          "provider_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_name": {
		SourcePath:        "account_name",
		DefaultOutputPath: "accountName",
		Column:            "account_name",
		SQLAlias:          "account_name",
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
	"owner_type": {
		SourcePath:        "owner_type",
		DefaultOutputPath: "ownerType",
		Column:            "owner_type",
		SQLAlias:          "owner_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"owner_id": {
		SourcePath:        "owner_id",
		DefaultOutputPath: "ownerId",
		Column:            "owner_id",
		SQLAlias:          "owner_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_ref": {
		SourcePath:        "merchant_ref",
		DefaultOutputPath: "merchantRef",
		Column:            "merchant_ref",
		SQLAlias:          "merchant_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"credential_secret_ref": {
		SourcePath:        "credential_secret_ref",
		DefaultOutputPath: "credentialSecretRef",
		Column:            "credential_secret_ref",
		SQLAlias:          "credential_secret_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"webhook_secret_ref": {
		SourcePath:        "webhook_secret_ref",
		DefaultOutputPath: "webhookSecretRef",
		Column:            "webhook_secret_ref",
		SQLAlias:          "webhook_secret_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"public_key_ref": {
		SourcePath:        "public_key_ref",
		DefaultOutputPath: "publicKeyRef",
		Column:            "public_key_ref",
		SQLAlias:          "public_key_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"status": {
		SourcePath:        "status",
		DefaultOutputPath: "status",
		Column:            "status",
		SQLAlias:          "status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"config": {
		SourcePath:        "config",
		DefaultOutputPath: "config",
		Column:            "config",
		SQLAlias:          "config",
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

func NewProviderAccountsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderAccountsFilterFields[field]
	return
}

type ProviderAccountsFilterResult struct {
	ProviderAccounts
	FilterCount int `db:"count"`
}

func ValidateProviderAccountsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderAccountsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderAccountsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderAccountsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderAccountsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderAccountsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ProviderAccountStatus string

const (
	ProviderAccountStatusActive     ProviderAccountStatus = "active"
	ProviderAccountStatusInactive   ProviderAccountStatus = "inactive"
	ProviderAccountStatusDeprecated ProviderAccountStatus = "deprecated"
)

type ProviderAccounts struct {
	Id                  uuid.UUID             `db:"id"`
	ProviderId          uuid.UUID             `db:"provider_id"`
	AccountName         string                `db:"account_name"`
	Environment         string                `db:"environment"`
	OwnerType           string                `db:"owner_type"`
	OwnerId             nuuid.NUUID           `db:"owner_id"`
	MerchantRef         null.String           `db:"merchant_ref"`
	CredentialSecretRef string                `db:"credential_secret_ref"`
	WebhookSecretRef    null.String           `db:"webhook_secret_ref"`
	PublicKeyRef        null.String           `db:"public_key_ref"`
	Status              ProviderAccountStatus `db:"status"`
	Config              json.RawMessage       `db:"config"`
	Metadata            json.RawMessage       `db:"metadata"`

	shared.MetaSignature
}
type ProviderAccountsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderAccounts) ToProviderAccountsPrimaryID() ProviderAccountsPrimaryID {
	return ProviderAccountsPrimaryID{
		Id: d.Id,
	}
}

type ProviderAccountsList []*ProviderAccounts

type ProviderAccountsFilterResultList []*ProviderAccountsFilterResult
