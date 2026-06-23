package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PayoutMethodsDBFieldNameType string

type payoutMethodsDBFieldName struct {
	Id                   PayoutMethodsDBFieldNameType
	MerchantPartyId      PayoutMethodsDBFieldNameType
	MethodType           PayoutMethodsDBFieldNameType
	ProviderCode         PayoutMethodsDBFieldNameType
	CountryCode          PayoutMethodsDBFieldNameType
	CurrencyCode         PayoutMethodsDBFieldNameType
	AccountHolderName    PayoutMethodsDBFieldNameType
	AccountNoLast4       PayoutMethodsDBFieldNameType
	BankCode             PayoutMethodsDBFieldNameType
	BankName             PayoutMethodsDBFieldNameType
	AccountTokenRef      PayoutMethodsDBFieldNameType
	AccountEncryptedBlob PayoutMethodsDBFieldNameType
	VerificationStatus   PayoutMethodsDBFieldNameType
	IsDefault            PayoutMethodsDBFieldNameType
	MethodStatus         PayoutMethodsDBFieldNameType
	Metadata             PayoutMethodsDBFieldNameType
	MetaCreatedAt        PayoutMethodsDBFieldNameType
	MetaCreatedBy        PayoutMethodsDBFieldNameType
	MetaUpdatedAt        PayoutMethodsDBFieldNameType
	MetaUpdatedBy        PayoutMethodsDBFieldNameType
	MetaDeletedAt        PayoutMethodsDBFieldNameType
	MetaDeletedBy        PayoutMethodsDBFieldNameType
}

var PayoutMethodsDBFieldName = payoutMethodsDBFieldName{
	Id:                   "id",
	MerchantPartyId:      "merchant_party_id",
	MethodType:           "method_type",
	ProviderCode:         "provider_code",
	CountryCode:          "country_code",
	CurrencyCode:         "currency_code",
	AccountHolderName:    "account_holder_name",
	AccountNoLast4:       "account_no_last4",
	BankCode:             "bank_code",
	BankName:             "bank_name",
	AccountTokenRef:      "account_token_ref",
	AccountEncryptedBlob: "account_encrypted_blob",
	VerificationStatus:   "verification_status",
	IsDefault:            "is_default",
	MethodStatus:         "method_status",
	Metadata:             "metadata",
	MetaCreatedAt:        "meta_created_at",
	MetaCreatedBy:        "meta_created_by",
	MetaUpdatedAt:        "meta_updated_at",
	MetaUpdatedBy:        "meta_updated_by",
	MetaDeletedAt:        "meta_deleted_at",
	MetaDeletedBy:        "meta_deleted_by",
}

func NewPayoutMethodsDBFieldNameFromStr(field string) (dbField PayoutMethodsDBFieldNameType, found bool) {
	switch field {

	case string(PayoutMethodsDBFieldName.Id):
		return PayoutMethodsDBFieldName.Id, true

	case string(PayoutMethodsDBFieldName.MerchantPartyId):
		return PayoutMethodsDBFieldName.MerchantPartyId, true

	case string(PayoutMethodsDBFieldName.MethodType):
		return PayoutMethodsDBFieldName.MethodType, true

	case string(PayoutMethodsDBFieldName.ProviderCode):
		return PayoutMethodsDBFieldName.ProviderCode, true

	case string(PayoutMethodsDBFieldName.CountryCode):
		return PayoutMethodsDBFieldName.CountryCode, true

	case string(PayoutMethodsDBFieldName.CurrencyCode):
		return PayoutMethodsDBFieldName.CurrencyCode, true

	case string(PayoutMethodsDBFieldName.AccountHolderName):
		return PayoutMethodsDBFieldName.AccountHolderName, true

	case string(PayoutMethodsDBFieldName.AccountNoLast4):
		return PayoutMethodsDBFieldName.AccountNoLast4, true

	case string(PayoutMethodsDBFieldName.BankCode):
		return PayoutMethodsDBFieldName.BankCode, true

	case string(PayoutMethodsDBFieldName.BankName):
		return PayoutMethodsDBFieldName.BankName, true

	case string(PayoutMethodsDBFieldName.AccountTokenRef):
		return PayoutMethodsDBFieldName.AccountTokenRef, true

	case string(PayoutMethodsDBFieldName.AccountEncryptedBlob):
		return PayoutMethodsDBFieldName.AccountEncryptedBlob, true

	case string(PayoutMethodsDBFieldName.VerificationStatus):
		return PayoutMethodsDBFieldName.VerificationStatus, true

	case string(PayoutMethodsDBFieldName.IsDefault):
		return PayoutMethodsDBFieldName.IsDefault, true

	case string(PayoutMethodsDBFieldName.MethodStatus):
		return PayoutMethodsDBFieldName.MethodStatus, true

	case string(PayoutMethodsDBFieldName.Metadata):
		return PayoutMethodsDBFieldName.Metadata, true

	case string(PayoutMethodsDBFieldName.MetaCreatedAt):
		return PayoutMethodsDBFieldName.MetaCreatedAt, true

	case string(PayoutMethodsDBFieldName.MetaCreatedBy):
		return PayoutMethodsDBFieldName.MetaCreatedBy, true

	case string(PayoutMethodsDBFieldName.MetaUpdatedAt):
		return PayoutMethodsDBFieldName.MetaUpdatedAt, true

	case string(PayoutMethodsDBFieldName.MetaUpdatedBy):
		return PayoutMethodsDBFieldName.MetaUpdatedBy, true

	case string(PayoutMethodsDBFieldName.MetaDeletedAt):
		return PayoutMethodsDBFieldName.MetaDeletedAt, true

	case string(PayoutMethodsDBFieldName.MetaDeletedBy):
		return PayoutMethodsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PayoutMethodsFilterJoins = map[string]JoinSpec{}

var PayoutMethodsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_party_id": {
		SourcePath:        "merchant_party_id",
		DefaultOutputPath: "merchantPartyId",
		Column:            "merchant_party_id",
		SQLAlias:          "merchant_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"method_type": {
		SourcePath:        "method_type",
		DefaultOutputPath: "methodType",
		Column:            "method_type",
		SQLAlias:          "method_type",
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
	"country_code": {
		SourcePath:        "country_code",
		DefaultOutputPath: "countryCode",
		Column:            "country_code",
		SQLAlias:          "country_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_holder_name": {
		SourcePath:        "account_holder_name",
		DefaultOutputPath: "accountHolderName",
		Column:            "account_holder_name",
		SQLAlias:          "account_holder_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_no_last4": {
		SourcePath:        "account_no_last4",
		DefaultOutputPath: "accountNoLast4",
		Column:            "account_no_last4",
		SQLAlias:          "account_no_last4",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"bank_code": {
		SourcePath:        "bank_code",
		DefaultOutputPath: "bankCode",
		Column:            "bank_code",
		SQLAlias:          "bank_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"bank_name": {
		SourcePath:        "bank_name",
		DefaultOutputPath: "bankName",
		Column:            "bank_name",
		SQLAlias:          "bank_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_token_ref": {
		SourcePath:        "account_token_ref",
		DefaultOutputPath: "accountTokenRef",
		Column:            "account_token_ref",
		SQLAlias:          "account_token_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_encrypted_blob": {
		SourcePath:        "account_encrypted_blob",
		DefaultOutputPath: "accountEncryptedBlob",
		Column:            "account_encrypted_blob",
		SQLAlias:          "account_encrypted_blob",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"verification_status": {
		SourcePath:        "verification_status",
		DefaultOutputPath: "verificationStatus",
		Column:            "verification_status",
		SQLAlias:          "verification_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_default": {
		SourcePath:        "is_default",
		DefaultOutputPath: "isDefault",
		Column:            "is_default",
		SQLAlias:          "is_default",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"method_status": {
		SourcePath:        "method_status",
		DefaultOutputPath: "methodStatus",
		Column:            "method_status",
		SQLAlias:          "method_status",
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

func NewPayoutMethodsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayoutMethodsFilterFields[field]
	return
}

type PayoutMethodsFilterResult struct {
	PayoutMethods
	FilterCount int `db:"count"`
}

func ValidatePayoutMethodsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPayoutMethodsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayoutMethodsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPayoutMethodsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePayoutMethodsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePayoutMethodsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPayoutMethodsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePayoutMethodsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type MethodStatus string

const (
	MethodStatusActive   MethodStatus = "active"
	MethodStatusInactive MethodStatus = "inactive"
	MethodStatusBlocked  MethodStatus = "blocked"
)

type MethodType string

const (
	MethodTypeBankAccount MethodType = "bank_account"
	MethodTypeEwallet     MethodType = "ewallet"
	MethodTypeManual      MethodType = "manual"
	MethodTypeVirtualCard MethodType = "virtual_card"
)

type VerificationStatus string

const (
	VerificationStatusPending  VerificationStatus = "pending"
	VerificationStatusVerified VerificationStatus = "verified"
	VerificationStatusFailed   VerificationStatus = "failed"
	VerificationStatusExpired  VerificationStatus = "expired"
)

type PayoutMethods struct {
	Id                   uuid.UUID          `db:"id"`
	MerchantPartyId      uuid.UUID          `db:"merchant_party_id"`
	MethodType           MethodType         `db:"method_type"`
	ProviderCode         null.String        `db:"provider_code"`
	CountryCode          string             `db:"country_code"`
	CurrencyCode         string             `db:"currency_code"`
	AccountHolderName    null.String        `db:"account_holder_name"`
	AccountNoLast4       null.String        `db:"account_no_last4"`
	BankCode             null.String        `db:"bank_code"`
	BankName             null.String        `db:"bank_name"`
	AccountTokenRef      null.String        `db:"account_token_ref"`
	AccountEncryptedBlob null.String        `db:"account_encrypted_blob"`
	VerificationStatus   VerificationStatus `db:"verification_status"`
	IsDefault            bool               `db:"is_default"`
	MethodStatus         MethodStatus       `db:"method_status"`
	Metadata             json.RawMessage    `db:"metadata"`

	shared.MetaSignature
}
type PayoutMethodsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayoutMethods) ToPayoutMethodsPrimaryID() PayoutMethodsPrimaryID {
	return PayoutMethodsPrimaryID{
		Id: d.Id,
	}
}

type PayoutMethodsList []*PayoutMethods

type PayoutMethodsFilterResultList []*PayoutMethodsFilterResult
