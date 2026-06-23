package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type MerchantBalanceAccountsDBFieldNameType string

type merchantBalanceAccountsDBFieldName struct {
	Id                    MerchantBalanceAccountsDBFieldNameType
	MerchantPartyId       MerchantBalanceAccountsDBFieldNameType
	BalanceType           MerchantBalanceAccountsDBFieldNameType
	CurrencyCode          MerchantBalanceAccountsDBFieldNameType
	LinkedLedgerAccountId MerchantBalanceAccountsDBFieldNameType
	AccountStatus         MerchantBalanceAccountsDBFieldNameType
	Metadata              MerchantBalanceAccountsDBFieldNameType
	MetaCreatedAt         MerchantBalanceAccountsDBFieldNameType
	MetaCreatedBy         MerchantBalanceAccountsDBFieldNameType
	MetaUpdatedAt         MerchantBalanceAccountsDBFieldNameType
	MetaUpdatedBy         MerchantBalanceAccountsDBFieldNameType
	MetaDeletedAt         MerchantBalanceAccountsDBFieldNameType
	MetaDeletedBy         MerchantBalanceAccountsDBFieldNameType
}

var MerchantBalanceAccountsDBFieldName = merchantBalanceAccountsDBFieldName{
	Id:                    "id",
	MerchantPartyId:       "merchant_party_id",
	BalanceType:           "balance_type",
	CurrencyCode:          "currency_code",
	LinkedLedgerAccountId: "linked_ledger_account_id",
	AccountStatus:         "account_status",
	Metadata:              "metadata",
	MetaCreatedAt:         "meta_created_at",
	MetaCreatedBy:         "meta_created_by",
	MetaUpdatedAt:         "meta_updated_at",
	MetaUpdatedBy:         "meta_updated_by",
	MetaDeletedAt:         "meta_deleted_at",
	MetaDeletedBy:         "meta_deleted_by",
}

func NewMerchantBalanceAccountsDBFieldNameFromStr(field string) (dbField MerchantBalanceAccountsDBFieldNameType, found bool) {
	switch field {

	case string(MerchantBalanceAccountsDBFieldName.Id):
		return MerchantBalanceAccountsDBFieldName.Id, true

	case string(MerchantBalanceAccountsDBFieldName.MerchantPartyId):
		return MerchantBalanceAccountsDBFieldName.MerchantPartyId, true

	case string(MerchantBalanceAccountsDBFieldName.BalanceType):
		return MerchantBalanceAccountsDBFieldName.BalanceType, true

	case string(MerchantBalanceAccountsDBFieldName.CurrencyCode):
		return MerchantBalanceAccountsDBFieldName.CurrencyCode, true

	case string(MerchantBalanceAccountsDBFieldName.LinkedLedgerAccountId):
		return MerchantBalanceAccountsDBFieldName.LinkedLedgerAccountId, true

	case string(MerchantBalanceAccountsDBFieldName.AccountStatus):
		return MerchantBalanceAccountsDBFieldName.AccountStatus, true

	case string(MerchantBalanceAccountsDBFieldName.Metadata):
		return MerchantBalanceAccountsDBFieldName.Metadata, true

	case string(MerchantBalanceAccountsDBFieldName.MetaCreatedAt):
		return MerchantBalanceAccountsDBFieldName.MetaCreatedAt, true

	case string(MerchantBalanceAccountsDBFieldName.MetaCreatedBy):
		return MerchantBalanceAccountsDBFieldName.MetaCreatedBy, true

	case string(MerchantBalanceAccountsDBFieldName.MetaUpdatedAt):
		return MerchantBalanceAccountsDBFieldName.MetaUpdatedAt, true

	case string(MerchantBalanceAccountsDBFieldName.MetaUpdatedBy):
		return MerchantBalanceAccountsDBFieldName.MetaUpdatedBy, true

	case string(MerchantBalanceAccountsDBFieldName.MetaDeletedAt):
		return MerchantBalanceAccountsDBFieldName.MetaDeletedAt, true

	case string(MerchantBalanceAccountsDBFieldName.MetaDeletedBy):
		return MerchantBalanceAccountsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var MerchantBalanceAccountsFilterJoins = map[string]JoinSpec{}

var MerchantBalanceAccountsFilterFields = map[string]FilterFieldSpec{
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
	"balance_type": {
		SourcePath:        "balance_type",
		DefaultOutputPath: "balanceType",
		Column:            "balance_type",
		SQLAlias:          "balance_type",
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
	"linked_ledger_account_id": {
		SourcePath:        "linked_ledger_account_id",
		DefaultOutputPath: "linkedLedgerAccountId",
		Column:            "linked_ledger_account_id",
		SQLAlias:          "linked_ledger_account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_status": {
		SourcePath:        "account_status",
		DefaultOutputPath: "accountStatus",
		Column:            "account_status",
		SQLAlias:          "account_status",
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

func NewMerchantBalanceAccountsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = MerchantBalanceAccountsFilterFields[field]
	return
}

type MerchantBalanceAccountsFilterResult struct {
	MerchantBalanceAccounts
	FilterCount int `db:"count"`
}

func ValidateMerchantBalanceAccountsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewMerchantBalanceAccountsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewMerchantBalanceAccountsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewMerchantBalanceAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateMerchantBalanceAccountsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateMerchantBalanceAccountsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewMerchantBalanceAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateMerchantBalanceAccountsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type MerchantBalanceAccountsAccountStatus string

const (
	MerchantBalanceAccountsAccountStatusActive   MerchantBalanceAccountsAccountStatus = "active"
	MerchantBalanceAccountsAccountStatusInactive MerchantBalanceAccountsAccountStatus = "inactive"
)

type MerchantBalanceAccountsBalanceType string

const (
	MerchantBalanceAccountsBalanceTypePending    MerchantBalanceAccountsBalanceType = "pending"
	MerchantBalanceAccountsBalanceTypeAvailable  MerchantBalanceAccountsBalanceType = "available"
	MerchantBalanceAccountsBalanceTypeReserved   MerchantBalanceAccountsBalanceType = "reserved"
	MerchantBalanceAccountsBalanceTypePayable    MerchantBalanceAccountsBalanceType = "payable"
	MerchantBalanceAccountsBalanceTypePaidOut    MerchantBalanceAccountsBalanceType = "paid_out"
	MerchantBalanceAccountsBalanceTypeNegative   MerchantBalanceAccountsBalanceType = "negative"
	MerchantBalanceAccountsBalanceTypeDisputed   MerchantBalanceAccountsBalanceType = "disputed"
	MerchantBalanceAccountsBalanceTypeRefundable MerchantBalanceAccountsBalanceType = "refundable"
)

type MerchantBalanceAccounts struct {
	Id                    uuid.UUID                            `db:"id"`
	MerchantPartyId       uuid.UUID                            `db:"merchant_party_id"`
	BalanceType           MerchantBalanceAccountsBalanceType   `db:"balance_type"`
	CurrencyCode          string                               `db:"currency_code"`
	LinkedLedgerAccountId uuid.UUID                            `db:"linked_ledger_account_id"`
	AccountStatus         MerchantBalanceAccountsAccountStatus `db:"account_status"`
	Metadata              json.RawMessage                      `db:"metadata"`

	shared.MetaSignature
}
type MerchantBalanceAccountsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d MerchantBalanceAccounts) ToMerchantBalanceAccountsPrimaryID() MerchantBalanceAccountsPrimaryID {
	return MerchantBalanceAccountsPrimaryID{
		Id: d.Id,
	}
}

type MerchantBalanceAccountsList []*MerchantBalanceAccounts

type MerchantBalanceAccountsFilterResultList []*MerchantBalanceAccountsFilterResult
