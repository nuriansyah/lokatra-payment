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

type LedgerAccountResolutionRulesDBFieldNameType string

type ledgerAccountResolutionRulesDBFieldName struct {
	Id              LedgerAccountResolutionRulesDBFieldNameType
	BookId          LedgerAccountResolutionRulesDBFieldNameType
	Purpose         LedgerAccountResolutionRulesDBFieldNameType
	MerchantPartyId LedgerAccountResolutionRulesDBFieldNameType
	ProviderCode    LedgerAccountResolutionRulesDBFieldNameType
	CurrencyCode    LedgerAccountResolutionRulesDBFieldNameType
	SourceType      LedgerAccountResolutionRulesDBFieldNameType
	SourceSubtype   LedgerAccountResolutionRulesDBFieldNameType
	AccountId       LedgerAccountResolutionRulesDBFieldNameType
	Priority        LedgerAccountResolutionRulesDBFieldNameType
	IsActive        LedgerAccountResolutionRulesDBFieldNameType
	Conditions      LedgerAccountResolutionRulesDBFieldNameType
	Metadata        LedgerAccountResolutionRulesDBFieldNameType
	MetaCreatedAt   LedgerAccountResolutionRulesDBFieldNameType
	MetaCreatedBy   LedgerAccountResolutionRulesDBFieldNameType
	MetaUpdatedAt   LedgerAccountResolutionRulesDBFieldNameType
	MetaUpdatedBy   LedgerAccountResolutionRulesDBFieldNameType
	MetaDeletedAt   LedgerAccountResolutionRulesDBFieldNameType
	MetaDeletedBy   LedgerAccountResolutionRulesDBFieldNameType
}

var LedgerAccountResolutionRulesDBFieldName = ledgerAccountResolutionRulesDBFieldName{
	Id:              "id",
	BookId:          "book_id",
	Purpose:         "purpose",
	MerchantPartyId: "merchant_party_id",
	ProviderCode:    "provider_code",
	CurrencyCode:    "currency_code",
	SourceType:      "source_type",
	SourceSubtype:   "source_subtype",
	AccountId:       "account_id",
	Priority:        "priority",
	IsActive:        "is_active",
	Conditions:      "conditions",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewLedgerAccountResolutionRulesDBFieldNameFromStr(field string) (dbField LedgerAccountResolutionRulesDBFieldNameType, found bool) {
	switch field {

	case string(LedgerAccountResolutionRulesDBFieldName.Id):
		return LedgerAccountResolutionRulesDBFieldName.Id, true

	case string(LedgerAccountResolutionRulesDBFieldName.BookId):
		return LedgerAccountResolutionRulesDBFieldName.BookId, true

	case string(LedgerAccountResolutionRulesDBFieldName.Purpose):
		return LedgerAccountResolutionRulesDBFieldName.Purpose, true

	case string(LedgerAccountResolutionRulesDBFieldName.MerchantPartyId):
		return LedgerAccountResolutionRulesDBFieldName.MerchantPartyId, true

	case string(LedgerAccountResolutionRulesDBFieldName.ProviderCode):
		return LedgerAccountResolutionRulesDBFieldName.ProviderCode, true

	case string(LedgerAccountResolutionRulesDBFieldName.CurrencyCode):
		return LedgerAccountResolutionRulesDBFieldName.CurrencyCode, true

	case string(LedgerAccountResolutionRulesDBFieldName.SourceType):
		return LedgerAccountResolutionRulesDBFieldName.SourceType, true

	case string(LedgerAccountResolutionRulesDBFieldName.SourceSubtype):
		return LedgerAccountResolutionRulesDBFieldName.SourceSubtype, true

	case string(LedgerAccountResolutionRulesDBFieldName.AccountId):
		return LedgerAccountResolutionRulesDBFieldName.AccountId, true

	case string(LedgerAccountResolutionRulesDBFieldName.Priority):
		return LedgerAccountResolutionRulesDBFieldName.Priority, true

	case string(LedgerAccountResolutionRulesDBFieldName.IsActive):
		return LedgerAccountResolutionRulesDBFieldName.IsActive, true

	case string(LedgerAccountResolutionRulesDBFieldName.Conditions):
		return LedgerAccountResolutionRulesDBFieldName.Conditions, true

	case string(LedgerAccountResolutionRulesDBFieldName.Metadata):
		return LedgerAccountResolutionRulesDBFieldName.Metadata, true

	case string(LedgerAccountResolutionRulesDBFieldName.MetaCreatedAt):
		return LedgerAccountResolutionRulesDBFieldName.MetaCreatedAt, true

	case string(LedgerAccountResolutionRulesDBFieldName.MetaCreatedBy):
		return LedgerAccountResolutionRulesDBFieldName.MetaCreatedBy, true

	case string(LedgerAccountResolutionRulesDBFieldName.MetaUpdatedAt):
		return LedgerAccountResolutionRulesDBFieldName.MetaUpdatedAt, true

	case string(LedgerAccountResolutionRulesDBFieldName.MetaUpdatedBy):
		return LedgerAccountResolutionRulesDBFieldName.MetaUpdatedBy, true

	case string(LedgerAccountResolutionRulesDBFieldName.MetaDeletedAt):
		return LedgerAccountResolutionRulesDBFieldName.MetaDeletedAt, true

	case string(LedgerAccountResolutionRulesDBFieldName.MetaDeletedBy):
		return LedgerAccountResolutionRulesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerAccountResolutionRulesFilterJoins = map[string]JoinSpec{}

var LedgerAccountResolutionRulesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"book_id": {
		SourcePath:        "book_id",
		DefaultOutputPath: "bookId",
		Column:            "book_id",
		SQLAlias:          "book_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"purpose": {
		SourcePath:        "purpose",
		DefaultOutputPath: "purpose",
		Column:            "purpose",
		SQLAlias:          "purpose",
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
	"provider_code": {
		SourcePath:        "provider_code",
		DefaultOutputPath: "providerCode",
		Column:            "provider_code",
		SQLAlias:          "provider_code",
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
	"source_type": {
		SourcePath:        "source_type",
		DefaultOutputPath: "sourceType",
		Column:            "source_type",
		SQLAlias:          "source_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_subtype": {
		SourcePath:        "source_subtype",
		DefaultOutputPath: "sourceSubtype",
		Column:            "source_subtype",
		SQLAlias:          "source_subtype",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_id": {
		SourcePath:        "account_id",
		DefaultOutputPath: "accountId",
		Column:            "account_id",
		SQLAlias:          "account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"priority": {
		SourcePath:        "priority",
		DefaultOutputPath: "priority",
		Column:            "priority",
		SQLAlias:          "priority",
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
	"conditions": {
		SourcePath:        "conditions",
		DefaultOutputPath: "conditions",
		Column:            "conditions",
		SQLAlias:          "conditions",
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

func NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerAccountResolutionRulesFilterFields[field]
	return
}

type LedgerAccountResolutionRulesFilterResult struct {
	LedgerAccountResolutionRules
	FilterCount int `db:"count"`
}

func ValidateLedgerAccountResolutionRulesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerAccountResolutionRulesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerAccountResolutionRulesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerAccountResolutionRulesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type Purpose string

const (
	PurposeProviderClearing      Purpose = "provider_clearing"
	PurposeMerchantPending       Purpose = "merchant_pending"
	PurposeMerchantPayable       Purpose = "merchant_payable"
	PurposeMerchantAvailable     Purpose = "merchant_available"
	PurposeMerchantReserved      Purpose = "merchant_reserved"
	PurposeMerchantNegative      Purpose = "merchant_negative"
	PurposeCommissionRevenue     Purpose = "commission_revenue"
	PurposeTaxPayable            Purpose = "tax_payable"
	PurposeCustomerRefundPayable Purpose = "customer_refund_payable"
	PurposePayoutClearing        Purpose = "payout_clearing"
	PurposeProviderFeeExpense    Purpose = "provider_fee_expense"
)

type LedgerAccountResolutionRules struct {
	Id              uuid.UUID       `db:"id"`
	BookId          uuid.UUID       `db:"book_id"`
	Purpose         Purpose         `db:"purpose"`
	MerchantPartyId nuuid.NUUID     `db:"merchant_party_id"`
	ProviderCode    null.String     `db:"provider_code"`
	CurrencyCode    string          `db:"currency_code"`
	SourceType      null.String     `db:"source_type"`
	SourceSubtype   null.String     `db:"source_subtype"`
	AccountId       uuid.UUID       `db:"account_id"`
	Priority        int             `db:"priority"`
	IsActive        bool            `db:"is_active"`
	Conditions      json.RawMessage `db:"conditions"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type LedgerAccountResolutionRulesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerAccountResolutionRules) ToLedgerAccountResolutionRulesPrimaryID() LedgerAccountResolutionRulesPrimaryID {
	return LedgerAccountResolutionRulesPrimaryID{
		Id: d.Id,
	}
}

type LedgerAccountResolutionRulesList []*LedgerAccountResolutionRules

type LedgerAccountResolutionRulesFilterResultList []*LedgerAccountResolutionRulesFilterResult
