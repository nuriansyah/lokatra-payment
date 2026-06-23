package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceProviderAccountsDBFieldNameType string

type financeProviderAccountsDBFieldName struct {
	Id                 FinanceProviderAccountsDBFieldNameType
	ProviderCode       FinanceProviderAccountsDBFieldNameType
	ProviderName       FinanceProviderAccountsDBFieldNameType
	OwnerPartyId       FinanceProviderAccountsDBFieldNameType
	Environment        FinanceProviderAccountsDBFieldNameType
	ApiBaseUrl         FinanceProviderAccountsDBFieldNameType
	MerchantRef        FinanceProviderAccountsDBFieldNameType
	SettlementCurrency FinanceProviderAccountsDBFieldNameType
	VaultSecretRef     FinanceProviderAccountsDBFieldNameType
	WebhookSecretRef   FinanceProviderAccountsDBFieldNameType
	ProviderStatus     FinanceProviderAccountsDBFieldNameType
	Capabilities       FinanceProviderAccountsDBFieldNameType
	Metadata           FinanceProviderAccountsDBFieldNameType
	MetaCreatedAt      FinanceProviderAccountsDBFieldNameType
	MetaCreatedBy      FinanceProviderAccountsDBFieldNameType
	MetaUpdatedAt      FinanceProviderAccountsDBFieldNameType
	MetaUpdatedBy      FinanceProviderAccountsDBFieldNameType
	MetaDeletedAt      FinanceProviderAccountsDBFieldNameType
	MetaDeletedBy      FinanceProviderAccountsDBFieldNameType
}

var FinanceProviderAccountsDBFieldName = financeProviderAccountsDBFieldName{
	Id:                 "id",
	ProviderCode:       "provider_code",
	ProviderName:       "provider_name",
	OwnerPartyId:       "owner_party_id",
	Environment:        "environment",
	ApiBaseUrl:         "api_base_url",
	MerchantRef:        "merchant_ref",
	SettlementCurrency: "settlement_currency",
	VaultSecretRef:     "vault_secret_ref",
	WebhookSecretRef:   "webhook_secret_ref",
	ProviderStatus:     "provider_status",
	Capabilities:       "capabilities",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewFinanceProviderAccountsDBFieldNameFromStr(field string) (dbField FinanceProviderAccountsDBFieldNameType, found bool) {
	switch field {

	case string(FinanceProviderAccountsDBFieldName.Id):
		return FinanceProviderAccountsDBFieldName.Id, true

	case string(FinanceProviderAccountsDBFieldName.ProviderCode):
		return FinanceProviderAccountsDBFieldName.ProviderCode, true

	case string(FinanceProviderAccountsDBFieldName.ProviderName):
		return FinanceProviderAccountsDBFieldName.ProviderName, true

	case string(FinanceProviderAccountsDBFieldName.OwnerPartyId):
		return FinanceProviderAccountsDBFieldName.OwnerPartyId, true

	case string(FinanceProviderAccountsDBFieldName.Environment):
		return FinanceProviderAccountsDBFieldName.Environment, true

	case string(FinanceProviderAccountsDBFieldName.ApiBaseUrl):
		return FinanceProviderAccountsDBFieldName.ApiBaseUrl, true

	case string(FinanceProviderAccountsDBFieldName.MerchantRef):
		return FinanceProviderAccountsDBFieldName.MerchantRef, true

	case string(FinanceProviderAccountsDBFieldName.SettlementCurrency):
		return FinanceProviderAccountsDBFieldName.SettlementCurrency, true

	case string(FinanceProviderAccountsDBFieldName.VaultSecretRef):
		return FinanceProviderAccountsDBFieldName.VaultSecretRef, true

	case string(FinanceProviderAccountsDBFieldName.WebhookSecretRef):
		return FinanceProviderAccountsDBFieldName.WebhookSecretRef, true

	case string(FinanceProviderAccountsDBFieldName.ProviderStatus):
		return FinanceProviderAccountsDBFieldName.ProviderStatus, true

	case string(FinanceProviderAccountsDBFieldName.Capabilities):
		return FinanceProviderAccountsDBFieldName.Capabilities, true

	case string(FinanceProviderAccountsDBFieldName.Metadata):
		return FinanceProviderAccountsDBFieldName.Metadata, true

	case string(FinanceProviderAccountsDBFieldName.MetaCreatedAt):
		return FinanceProviderAccountsDBFieldName.MetaCreatedAt, true

	case string(FinanceProviderAccountsDBFieldName.MetaCreatedBy):
		return FinanceProviderAccountsDBFieldName.MetaCreatedBy, true

	case string(FinanceProviderAccountsDBFieldName.MetaUpdatedAt):
		return FinanceProviderAccountsDBFieldName.MetaUpdatedAt, true

	case string(FinanceProviderAccountsDBFieldName.MetaUpdatedBy):
		return FinanceProviderAccountsDBFieldName.MetaUpdatedBy, true

	case string(FinanceProviderAccountsDBFieldName.MetaDeletedAt):
		return FinanceProviderAccountsDBFieldName.MetaDeletedAt, true

	case string(FinanceProviderAccountsDBFieldName.MetaDeletedBy):
		return FinanceProviderAccountsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceProviderAccountsFilterJoins = map[string]JoinSpec{}

var FinanceProviderAccountsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
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
	"provider_name": {
		SourcePath:        "provider_name",
		DefaultOutputPath: "providerName",
		Column:            "provider_name",
		SQLAlias:          "provider_name",
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
	"environment": {
		SourcePath:        "environment",
		DefaultOutputPath: "environment",
		Column:            "environment",
		SQLAlias:          "environment",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"api_base_url": {
		SourcePath:        "api_base_url",
		DefaultOutputPath: "apiBaseUrl",
		Column:            "api_base_url",
		SQLAlias:          "api_base_url",
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
	"settlement_currency": {
		SourcePath:        "settlement_currency",
		DefaultOutputPath: "settlementCurrency",
		Column:            "settlement_currency",
		SQLAlias:          "settlement_currency",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"vault_secret_ref": {
		SourcePath:        "vault_secret_ref",
		DefaultOutputPath: "vaultSecretRef",
		Column:            "vault_secret_ref",
		SQLAlias:          "vault_secret_ref",
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
	"provider_status": {
		SourcePath:        "provider_status",
		DefaultOutputPath: "providerStatus",
		Column:            "provider_status",
		SQLAlias:          "provider_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"capabilities": {
		SourcePath:        "capabilities",
		DefaultOutputPath: "capabilities",
		Column:            "capabilities",
		SQLAlias:          "capabilities",
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

func NewFinanceProviderAccountsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceProviderAccountsFilterFields[field]
	return
}

type FinanceProviderAccountsFilterResult struct {
	FinanceProviderAccounts
	FilterCount int `db:"count"`
}

func ValidateFinanceProviderAccountsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceProviderAccountsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceProviderAccountsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceProviderAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceProviderAccountsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceProviderAccountsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceProviderAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceProviderAccountsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type Environment string

const (
	EnvironmentSandbox    Environment = "sandbox"
	EnvironmentProduction Environment = "production"
)

type ProviderStatus string

const (
	ProviderStatusActive   ProviderStatus = "active"
	ProviderStatusInactive ProviderStatus = "inactive"
)

type FinanceProviderAccounts struct {
	Id                 uuid.UUID       `db:"id"`
	ProviderCode       string          `db:"provider_code"`
	ProviderName       string          `db:"provider_name"`
	OwnerPartyId       uuid.UUID       `db:"owner_party_id"`
	Environment        Environment     `db:"environment"`
	ApiBaseUrl         null.String     `db:"api_base_url"`
	MerchantRef        null.String     `db:"merchant_ref"`
	SettlementCurrency string          `db:"settlement_currency"`
	VaultSecretRef     string          `db:"vault_secret_ref"`
	WebhookSecretRef   null.String     `db:"webhook_secret_ref"`
	ProviderStatus     ProviderStatus  `db:"provider_status"`
	Capabilities       json.RawMessage `db:"capabilities"`
	Metadata           json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceProviderAccountsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceProviderAccounts) ToFinanceProviderAccountsPrimaryID() FinanceProviderAccountsPrimaryID {
	return FinanceProviderAccountsPrimaryID{
		Id: d.Id,
	}
}

type FinanceProviderAccountsList []*FinanceProviderAccounts

type FinanceProviderAccountsFilterResultList []*FinanceProviderAccountsFilterResult
