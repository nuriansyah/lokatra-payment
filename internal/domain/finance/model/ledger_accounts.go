package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type LedgerAccountsDBFieldNameType string

type ledgerAccountsDBFieldName struct {
	Id                 LedgerAccountsDBFieldNameType
	BookId             LedgerAccountsDBFieldNameType
	AccountCode        LedgerAccountsDBFieldNameType
	AccountName        LedgerAccountsDBFieldNameType
	AccountTypeCode    LedgerAccountsDBFieldNameType
	OwnerPartyId       LedgerAccountsDBFieldNameType
	ParentAccountId    LedgerAccountsDBFieldNameType
	CurrencyCode       LedgerAccountsDBFieldNameType
	AllowManualPosting LedgerAccountsDBFieldNameType
	AccountStatus      LedgerAccountsDBFieldNameType
	Metadata           LedgerAccountsDBFieldNameType
	MetaCreatedAt      LedgerAccountsDBFieldNameType
	MetaCreatedBy      LedgerAccountsDBFieldNameType
	MetaUpdatedAt      LedgerAccountsDBFieldNameType
	MetaUpdatedBy      LedgerAccountsDBFieldNameType
	MetaDeletedAt      LedgerAccountsDBFieldNameType
	MetaDeletedBy      LedgerAccountsDBFieldNameType
}

var LedgerAccountsDBFieldName = ledgerAccountsDBFieldName{
	Id:                 "id",
	BookId:             "book_id",
	AccountCode:        "account_code",
	AccountName:        "account_name",
	AccountTypeCode:    "account_type_code",
	OwnerPartyId:       "owner_party_id",
	ParentAccountId:    "parent_account_id",
	CurrencyCode:       "currency_code",
	AllowManualPosting: "allow_manual_posting",
	AccountStatus:      "account_status",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewLedgerAccountsDBFieldNameFromStr(field string) (dbField LedgerAccountsDBFieldNameType, found bool) {
	switch field {

	case string(LedgerAccountsDBFieldName.Id):
		return LedgerAccountsDBFieldName.Id, true

	case string(LedgerAccountsDBFieldName.BookId):
		return LedgerAccountsDBFieldName.BookId, true

	case string(LedgerAccountsDBFieldName.AccountCode):
		return LedgerAccountsDBFieldName.AccountCode, true

	case string(LedgerAccountsDBFieldName.AccountName):
		return LedgerAccountsDBFieldName.AccountName, true

	case string(LedgerAccountsDBFieldName.AccountTypeCode):
		return LedgerAccountsDBFieldName.AccountTypeCode, true

	case string(LedgerAccountsDBFieldName.OwnerPartyId):
		return LedgerAccountsDBFieldName.OwnerPartyId, true

	case string(LedgerAccountsDBFieldName.ParentAccountId):
		return LedgerAccountsDBFieldName.ParentAccountId, true

	case string(LedgerAccountsDBFieldName.CurrencyCode):
		return LedgerAccountsDBFieldName.CurrencyCode, true

	case string(LedgerAccountsDBFieldName.AllowManualPosting):
		return LedgerAccountsDBFieldName.AllowManualPosting, true

	case string(LedgerAccountsDBFieldName.AccountStatus):
		return LedgerAccountsDBFieldName.AccountStatus, true

	case string(LedgerAccountsDBFieldName.Metadata):
		return LedgerAccountsDBFieldName.Metadata, true

	case string(LedgerAccountsDBFieldName.MetaCreatedAt):
		return LedgerAccountsDBFieldName.MetaCreatedAt, true

	case string(LedgerAccountsDBFieldName.MetaCreatedBy):
		return LedgerAccountsDBFieldName.MetaCreatedBy, true

	case string(LedgerAccountsDBFieldName.MetaUpdatedAt):
		return LedgerAccountsDBFieldName.MetaUpdatedAt, true

	case string(LedgerAccountsDBFieldName.MetaUpdatedBy):
		return LedgerAccountsDBFieldName.MetaUpdatedBy, true

	case string(LedgerAccountsDBFieldName.MetaDeletedAt):
		return LedgerAccountsDBFieldName.MetaDeletedAt, true

	case string(LedgerAccountsDBFieldName.MetaDeletedBy):
		return LedgerAccountsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var LedgerAccountsFilterJoins = map[string]JoinSpec{}

var LedgerAccountsFilterFields = map[string]FilterFieldSpec{
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
	"account_code": {
		SourcePath:        "account_code",
		DefaultOutputPath: "accountCode",
		Column:            "account_code",
		SQLAlias:          "account_code",
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
	"account_type_code": {
		SourcePath:        "account_type_code",
		DefaultOutputPath: "accountTypeCode",
		Column:            "account_type_code",
		SQLAlias:          "account_type_code",
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
	"parent_account_id": {
		SourcePath:        "parent_account_id",
		DefaultOutputPath: "parentAccountId",
		Column:            "parent_account_id",
		SQLAlias:          "parent_account_id",
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
	"allow_manual_posting": {
		SourcePath:        "allow_manual_posting",
		DefaultOutputPath: "allowManualPosting",
		Column:            "allow_manual_posting",
		SQLAlias:          "allow_manual_posting",
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

func NewLedgerAccountsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = LedgerAccountsFilterFields[field]
	return
}

type LedgerAccountsFilterResult struct {
	LedgerAccounts
	FilterCount int `db:"count"`
}

func ValidateLedgerAccountsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewLedgerAccountsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewLedgerAccountsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewLedgerAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateLedgerAccountsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateLedgerAccountsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewLedgerAccountsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateLedgerAccountsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type LedgerAccountsAccountStatus string

const (
	LedgerAccountsAccountStatusActive   LedgerAccountsAccountStatus = "active"
	LedgerAccountsAccountStatusInactive LedgerAccountsAccountStatus = "inactive"
	LedgerAccountsAccountStatusClosed   LedgerAccountsAccountStatus = "closed"
)

type LedgerAccounts struct {
	Id                 uuid.UUID                   `db:"id"`
	BookId             uuid.UUID                   `db:"book_id"`
	AccountCode        string                      `db:"account_code"`
	AccountName        string                      `db:"account_name"`
	AccountTypeCode    string                      `db:"account_type_code"`
	OwnerPartyId       nuuid.NUUID                 `db:"owner_party_id"`
	ParentAccountId    nuuid.NUUID                 `db:"parent_account_id"`
	CurrencyCode       string                      `db:"currency_code"`
	AllowManualPosting bool                        `db:"allow_manual_posting"`
	AccountStatus      LedgerAccountsAccountStatus `db:"account_status"`
	Metadata           json.RawMessage             `db:"metadata"`

	shared.MetaSignature
}
type LedgerAccountsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d LedgerAccounts) ToLedgerAccountsPrimaryID() LedgerAccountsPrimaryID {
	return LedgerAccountsPrimaryID{
		Id: d.Id,
	}
}

type LedgerAccountsList []*LedgerAccounts

type LedgerAccountsFilterResultList []*LedgerAccountsFilterResult
