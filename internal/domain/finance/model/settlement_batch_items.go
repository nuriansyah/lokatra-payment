package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type SettlementBatchItemsDBFieldNameType string

type settlementBatchItemsDBFieldName struct {
	Id                   SettlementBatchItemsDBFieldNameType
	SettlementBatchId    SettlementBatchItemsDBFieldNameType
	SourceType           SettlementBatchItemsDBFieldNameType
	SourceId             SettlementBatchItemsDBFieldNameType
	MerchantPartyId      SettlementBatchItemsDBFieldNameType
	CurrencyCode         SettlementBatchItemsDBFieldNameType
	GrossAmount          SettlementBatchItemsDBFieldNameType
	FeeAmount            SettlementBatchItemsDBFieldNameType
	TaxAmount            SettlementBatchItemsDBFieldNameType
	ReserveAmount        SettlementBatchItemsDBFieldNameType
	NetAmount            SettlementBatchItemsDBFieldNameType
	LinkedJournalEntryId SettlementBatchItemsDBFieldNameType
	ItemStatus           SettlementBatchItemsDBFieldNameType
	Metadata             SettlementBatchItemsDBFieldNameType
	MetaCreatedAt        SettlementBatchItemsDBFieldNameType
	MetaCreatedBy        SettlementBatchItemsDBFieldNameType
	MetaUpdatedAt        SettlementBatchItemsDBFieldNameType
	MetaUpdatedBy        SettlementBatchItemsDBFieldNameType
	MetaDeletedAt        SettlementBatchItemsDBFieldNameType
	MetaDeletedBy        SettlementBatchItemsDBFieldNameType
}

var SettlementBatchItemsDBFieldName = settlementBatchItemsDBFieldName{
	Id:                   "id",
	SettlementBatchId:    "settlement_batch_id",
	SourceType:           "source_type",
	SourceId:             "source_id",
	MerchantPartyId:      "merchant_party_id",
	CurrencyCode:         "currency_code",
	GrossAmount:          "gross_amount",
	FeeAmount:            "fee_amount",
	TaxAmount:            "tax_amount",
	ReserveAmount:        "reserve_amount",
	NetAmount:            "net_amount",
	LinkedJournalEntryId: "linked_journal_entry_id",
	ItemStatus:           "item_status",
	Metadata:             "metadata",
	MetaCreatedAt:        "meta_created_at",
	MetaCreatedBy:        "meta_created_by",
	MetaUpdatedAt:        "meta_updated_at",
	MetaUpdatedBy:        "meta_updated_by",
	MetaDeletedAt:        "meta_deleted_at",
	MetaDeletedBy:        "meta_deleted_by",
}

func NewSettlementBatchItemsDBFieldNameFromStr(field string) (dbField SettlementBatchItemsDBFieldNameType, found bool) {
	switch field {

	case string(SettlementBatchItemsDBFieldName.Id):
		return SettlementBatchItemsDBFieldName.Id, true

	case string(SettlementBatchItemsDBFieldName.SettlementBatchId):
		return SettlementBatchItemsDBFieldName.SettlementBatchId, true

	case string(SettlementBatchItemsDBFieldName.SourceType):
		return SettlementBatchItemsDBFieldName.SourceType, true

	case string(SettlementBatchItemsDBFieldName.SourceId):
		return SettlementBatchItemsDBFieldName.SourceId, true

	case string(SettlementBatchItemsDBFieldName.MerchantPartyId):
		return SettlementBatchItemsDBFieldName.MerchantPartyId, true

	case string(SettlementBatchItemsDBFieldName.CurrencyCode):
		return SettlementBatchItemsDBFieldName.CurrencyCode, true

	case string(SettlementBatchItemsDBFieldName.GrossAmount):
		return SettlementBatchItemsDBFieldName.GrossAmount, true

	case string(SettlementBatchItemsDBFieldName.FeeAmount):
		return SettlementBatchItemsDBFieldName.FeeAmount, true

	case string(SettlementBatchItemsDBFieldName.TaxAmount):
		return SettlementBatchItemsDBFieldName.TaxAmount, true

	case string(SettlementBatchItemsDBFieldName.ReserveAmount):
		return SettlementBatchItemsDBFieldName.ReserveAmount, true

	case string(SettlementBatchItemsDBFieldName.NetAmount):
		return SettlementBatchItemsDBFieldName.NetAmount, true

	case string(SettlementBatchItemsDBFieldName.LinkedJournalEntryId):
		return SettlementBatchItemsDBFieldName.LinkedJournalEntryId, true

	case string(SettlementBatchItemsDBFieldName.ItemStatus):
		return SettlementBatchItemsDBFieldName.ItemStatus, true

	case string(SettlementBatchItemsDBFieldName.Metadata):
		return SettlementBatchItemsDBFieldName.Metadata, true

	case string(SettlementBatchItemsDBFieldName.MetaCreatedAt):
		return SettlementBatchItemsDBFieldName.MetaCreatedAt, true

	case string(SettlementBatchItemsDBFieldName.MetaCreatedBy):
		return SettlementBatchItemsDBFieldName.MetaCreatedBy, true

	case string(SettlementBatchItemsDBFieldName.MetaUpdatedAt):
		return SettlementBatchItemsDBFieldName.MetaUpdatedAt, true

	case string(SettlementBatchItemsDBFieldName.MetaUpdatedBy):
		return SettlementBatchItemsDBFieldName.MetaUpdatedBy, true

	case string(SettlementBatchItemsDBFieldName.MetaDeletedAt):
		return SettlementBatchItemsDBFieldName.MetaDeletedAt, true

	case string(SettlementBatchItemsDBFieldName.MetaDeletedBy):
		return SettlementBatchItemsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var SettlementBatchItemsFilterJoins = map[string]JoinSpec{}

var SettlementBatchItemsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"settlement_batch_id": {
		SourcePath:        "settlement_batch_id",
		DefaultOutputPath: "settlementBatchId",
		Column:            "settlement_batch_id",
		SQLAlias:          "settlement_batch_id",
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
	"merchant_party_id": {
		SourcePath:        "merchant_party_id",
		DefaultOutputPath: "merchantPartyId",
		Column:            "merchant_party_id",
		SQLAlias:          "merchant_party_id",
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
	"gross_amount": {
		SourcePath:        "gross_amount",
		DefaultOutputPath: "grossAmount",
		Column:            "gross_amount",
		SQLAlias:          "gross_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_amount": {
		SourcePath:        "fee_amount",
		DefaultOutputPath: "feeAmount",
		Column:            "fee_amount",
		SQLAlias:          "fee_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_amount": {
		SourcePath:        "tax_amount",
		DefaultOutputPath: "taxAmount",
		Column:            "tax_amount",
		SQLAlias:          "tax_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reserve_amount": {
		SourcePath:        "reserve_amount",
		DefaultOutputPath: "reserveAmount",
		Column:            "reserve_amount",
		SQLAlias:          "reserve_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"net_amount": {
		SourcePath:        "net_amount",
		DefaultOutputPath: "netAmount",
		Column:            "net_amount",
		SQLAlias:          "net_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"linked_journal_entry_id": {
		SourcePath:        "linked_journal_entry_id",
		DefaultOutputPath: "linkedJournalEntryId",
		Column:            "linked_journal_entry_id",
		SQLAlias:          "linked_journal_entry_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"item_status": {
		SourcePath:        "item_status",
		DefaultOutputPath: "itemStatus",
		Column:            "item_status",
		SQLAlias:          "item_status",
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

func NewSettlementBatchItemsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = SettlementBatchItemsFilterFields[field]
	return
}

type SettlementBatchItemsFilterResult struct {
	SettlementBatchItems
	FilterCount int `db:"count"`
}

func ValidateSettlementBatchItemsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewSettlementBatchItemsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewSettlementBatchItemsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewSettlementBatchItemsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateSettlementBatchItemsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateSettlementBatchItemsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewSettlementBatchItemsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateSettlementBatchItemsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ItemStatus string

const (
	ItemStatusIncluded ItemStatus = "included"
	ItemStatusReleased ItemStatus = "released"
	ItemStatusReversed ItemStatus = "reversed"
)

type SettlementBatchItems struct {
	Id                   uuid.UUID       `db:"id"`
	SettlementBatchId    uuid.UUID       `db:"settlement_batch_id"`
	SourceType           string          `db:"source_type"`
	SourceId             uuid.UUID       `db:"source_id"`
	MerchantPartyId      uuid.UUID       `db:"merchant_party_id"`
	CurrencyCode         string          `db:"currency_code"`
	GrossAmount          decimal.Decimal `db:"gross_amount"`
	FeeAmount            decimal.Decimal `db:"fee_amount"`
	TaxAmount            decimal.Decimal `db:"tax_amount"`
	ReserveAmount        decimal.Decimal `db:"reserve_amount"`
	NetAmount            decimal.Decimal `db:"net_amount"`
	LinkedJournalEntryId nuuid.NUUID     `db:"linked_journal_entry_id"`
	ItemStatus           ItemStatus      `db:"item_status"`
	Metadata             json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type SettlementBatchItemsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d SettlementBatchItems) ToSettlementBatchItemsPrimaryID() SettlementBatchItemsPrimaryID {
	return SettlementBatchItemsPrimaryID{
		Id: d.Id,
	}
}

type SettlementBatchItemsList []*SettlementBatchItems

type SettlementBatchItemsFilterResultList []*SettlementBatchItemsFilterResult
