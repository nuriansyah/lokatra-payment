package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type MerchantBalanceTransactionsDBFieldNameType string

type merchantBalanceTransactionsDBFieldName struct {
	Id               MerchantBalanceTransactionsDBFieldNameType
	BalanceAccountId MerchantBalanceTransactionsDBFieldNameType
	MerchantPartyId  MerchantBalanceTransactionsDBFieldNameType
	BalanceType      MerchantBalanceTransactionsDBFieldNameType
	CurrencyCode     MerchantBalanceTransactionsDBFieldNameType
	SourceType       MerchantBalanceTransactionsDBFieldNameType
	SourceId         MerchantBalanceTransactionsDBFieldNameType
	Direction        MerchantBalanceTransactionsDBFieldNameType
	Amount           MerchantBalanceTransactionsDBFieldNameType
	BalanceBefore    MerchantBalanceTransactionsDBFieldNameType
	BalanceAfter     MerchantBalanceTransactionsDBFieldNameType
	ReasonCode       MerchantBalanceTransactionsDBFieldNameType
	Metadata         MerchantBalanceTransactionsDBFieldNameType
	MetaCreatedAt    MerchantBalanceTransactionsDBFieldNameType
	MetaCreatedBy    MerchantBalanceTransactionsDBFieldNameType
	MetaUpdatedAt    MerchantBalanceTransactionsDBFieldNameType
	MetaUpdatedBy    MerchantBalanceTransactionsDBFieldNameType
	MetaDeletedAt    MerchantBalanceTransactionsDBFieldNameType
	MetaDeletedBy    MerchantBalanceTransactionsDBFieldNameType
}

var MerchantBalanceTransactionsDBFieldName = merchantBalanceTransactionsDBFieldName{
	Id:               "id",
	BalanceAccountId: "balance_account_id",
	MerchantPartyId:  "merchant_party_id",
	BalanceType:      "balance_type",
	CurrencyCode:     "currency_code",
	SourceType:       "source_type",
	SourceId:         "source_id",
	Direction:        "direction",
	Amount:           "amount",
	BalanceBefore:    "balance_before",
	BalanceAfter:     "balance_after",
	ReasonCode:       "reason_code",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewMerchantBalanceTransactionsDBFieldNameFromStr(field string) (dbField MerchantBalanceTransactionsDBFieldNameType, found bool) {
	switch field {

	case string(MerchantBalanceTransactionsDBFieldName.Id):
		return MerchantBalanceTransactionsDBFieldName.Id, true

	case string(MerchantBalanceTransactionsDBFieldName.BalanceAccountId):
		return MerchantBalanceTransactionsDBFieldName.BalanceAccountId, true

	case string(MerchantBalanceTransactionsDBFieldName.MerchantPartyId):
		return MerchantBalanceTransactionsDBFieldName.MerchantPartyId, true

	case string(MerchantBalanceTransactionsDBFieldName.BalanceType):
		return MerchantBalanceTransactionsDBFieldName.BalanceType, true

	case string(MerchantBalanceTransactionsDBFieldName.CurrencyCode):
		return MerchantBalanceTransactionsDBFieldName.CurrencyCode, true

	case string(MerchantBalanceTransactionsDBFieldName.SourceType):
		return MerchantBalanceTransactionsDBFieldName.SourceType, true

	case string(MerchantBalanceTransactionsDBFieldName.SourceId):
		return MerchantBalanceTransactionsDBFieldName.SourceId, true

	case string(MerchantBalanceTransactionsDBFieldName.Direction):
		return MerchantBalanceTransactionsDBFieldName.Direction, true

	case string(MerchantBalanceTransactionsDBFieldName.Amount):
		return MerchantBalanceTransactionsDBFieldName.Amount, true

	case string(MerchantBalanceTransactionsDBFieldName.BalanceBefore):
		return MerchantBalanceTransactionsDBFieldName.BalanceBefore, true

	case string(MerchantBalanceTransactionsDBFieldName.BalanceAfter):
		return MerchantBalanceTransactionsDBFieldName.BalanceAfter, true

	case string(MerchantBalanceTransactionsDBFieldName.ReasonCode):
		return MerchantBalanceTransactionsDBFieldName.ReasonCode, true

	case string(MerchantBalanceTransactionsDBFieldName.Metadata):
		return MerchantBalanceTransactionsDBFieldName.Metadata, true

	case string(MerchantBalanceTransactionsDBFieldName.MetaCreatedAt):
		return MerchantBalanceTransactionsDBFieldName.MetaCreatedAt, true

	case string(MerchantBalanceTransactionsDBFieldName.MetaCreatedBy):
		return MerchantBalanceTransactionsDBFieldName.MetaCreatedBy, true

	case string(MerchantBalanceTransactionsDBFieldName.MetaUpdatedAt):
		return MerchantBalanceTransactionsDBFieldName.MetaUpdatedAt, true

	case string(MerchantBalanceTransactionsDBFieldName.MetaUpdatedBy):
		return MerchantBalanceTransactionsDBFieldName.MetaUpdatedBy, true

	case string(MerchantBalanceTransactionsDBFieldName.MetaDeletedAt):
		return MerchantBalanceTransactionsDBFieldName.MetaDeletedAt, true

	case string(MerchantBalanceTransactionsDBFieldName.MetaDeletedBy):
		return MerchantBalanceTransactionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var MerchantBalanceTransactionsFilterJoins = map[string]JoinSpec{}

var MerchantBalanceTransactionsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"balance_account_id": {
		SourcePath:        "balance_account_id",
		DefaultOutputPath: "balanceAccountId",
		Column:            "balance_account_id",
		SQLAlias:          "balance_account_id",
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
	"source_type": {
		SourcePath:        "source_type",
		DefaultOutputPath: "sourceType",
		Column:            "source_type",
		SQLAlias:          "source_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_id": {
		SourcePath:        "source_id",
		DefaultOutputPath: "sourceId",
		Column:            "source_id",
		SQLAlias:          "source_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"direction": {
		SourcePath:        "direction",
		DefaultOutputPath: "direction",
		Column:            "direction",
		SQLAlias:          "direction",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"balance_before": {
		SourcePath:        "balance_before",
		DefaultOutputPath: "balanceBefore",
		Column:            "balance_before",
		SQLAlias:          "balance_before",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"balance_after": {
		SourcePath:        "balance_after",
		DefaultOutputPath: "balanceAfter",
		Column:            "balance_after",
		SQLAlias:          "balance_after",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reason_code": {
		SourcePath:        "reason_code",
		DefaultOutputPath: "reasonCode",
		Column:            "reason_code",
		SQLAlias:          "reason_code",
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

func NewMerchantBalanceTransactionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = MerchantBalanceTransactionsFilterFields[field]
	return
}

type MerchantBalanceTransactionsFilterResult struct {
	MerchantBalanceTransactions
	FilterCount int `db:"count"`
}

func ValidateMerchantBalanceTransactionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewMerchantBalanceTransactionsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewMerchantBalanceTransactionsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewMerchantBalanceTransactionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateMerchantBalanceTransactionsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateMerchantBalanceTransactionsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewMerchantBalanceTransactionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateMerchantBalanceTransactionsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type Direction string

const (
	DirectionDebit  Direction = "debit"
	DirectionCredit Direction = "credit"
)

type MerchantBalanceTransactionsBalanceType string

const (
	MerchantBalanceTransactionsBalanceTypePending    MerchantBalanceTransactionsBalanceType = "pending"
	MerchantBalanceTransactionsBalanceTypeAvailable  MerchantBalanceTransactionsBalanceType = "available"
	MerchantBalanceTransactionsBalanceTypeReserved   MerchantBalanceTransactionsBalanceType = "reserved"
	MerchantBalanceTransactionsBalanceTypePayable    MerchantBalanceTransactionsBalanceType = "payable"
	MerchantBalanceTransactionsBalanceTypePaidOut    MerchantBalanceTransactionsBalanceType = "paid_out"
	MerchantBalanceTransactionsBalanceTypeNegative   MerchantBalanceTransactionsBalanceType = "negative"
	MerchantBalanceTransactionsBalanceTypeDisputed   MerchantBalanceTransactionsBalanceType = "disputed"
	MerchantBalanceTransactionsBalanceTypeRefundable MerchantBalanceTransactionsBalanceType = "refundable"
)

type MerchantBalanceTransactions struct {
	Id               uuid.UUID                              `db:"id"`
	BalanceAccountId uuid.UUID                              `db:"balance_account_id"`
	MerchantPartyId  uuid.UUID                              `db:"merchant_party_id"`
	BalanceType      MerchantBalanceTransactionsBalanceType `db:"balance_type"`
	CurrencyCode     string                                 `db:"currency_code"`
	SourceType       string                                 `db:"source_type"`
	SourceId         uuid.UUID                              `db:"source_id"`
	Direction        Direction                              `db:"direction"`
	Amount           decimal.Decimal                        `db:"amount"`
	BalanceBefore    decimal.Decimal                        `db:"balance_before"`
	BalanceAfter     decimal.Decimal                        `db:"balance_after"`
	ReasonCode       string                                 `db:"reason_code"`
	Metadata         json.RawMessage                        `db:"metadata"`

	shared.MetaSignature
}
type MerchantBalanceTransactionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d MerchantBalanceTransactions) ToMerchantBalanceTransactionsPrimaryID() MerchantBalanceTransactionsPrimaryID {
	return MerchantBalanceTransactionsPrimaryID{
		Id: d.Id,
	}
}

type MerchantBalanceTransactionsList []*MerchantBalanceTransactions

type MerchantBalanceTransactionsFilterResultList []*MerchantBalanceTransactionsFilterResult
