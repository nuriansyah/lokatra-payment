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

type RefundAllocationsDBFieldNameType string

type refundAllocationsDBFieldName struct {
	Id                 RefundAllocationsDBFieldNameType
	RefundId           RefundAllocationsDBFieldNameType
	AllocationType     RefundAllocationsDBFieldNameType
	ResponsiblePartyId RefundAllocationsDBFieldNameType
	Amount             RefundAllocationsDBFieldNameType
	CurrencyCode       RefundAllocationsDBFieldNameType
	AllocationStatus   RefundAllocationsDBFieldNameType
	Metadata           RefundAllocationsDBFieldNameType
	MetaCreatedAt      RefundAllocationsDBFieldNameType
	MetaCreatedBy      RefundAllocationsDBFieldNameType
	MetaUpdatedAt      RefundAllocationsDBFieldNameType
	MetaUpdatedBy      RefundAllocationsDBFieldNameType
	MetaDeletedAt      RefundAllocationsDBFieldNameType
	MetaDeletedBy      RefundAllocationsDBFieldNameType
}

var RefundAllocationsDBFieldName = refundAllocationsDBFieldName{
	Id:                 "id",
	RefundId:           "refund_id",
	AllocationType:     "allocation_type",
	ResponsiblePartyId: "responsible_party_id",
	Amount:             "amount",
	CurrencyCode:       "currency_code",
	AllocationStatus:   "allocation_status",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewRefundAllocationsDBFieldNameFromStr(field string) (dbField RefundAllocationsDBFieldNameType, found bool) {
	switch field {

	case string(RefundAllocationsDBFieldName.Id):
		return RefundAllocationsDBFieldName.Id, true

	case string(RefundAllocationsDBFieldName.RefundId):
		return RefundAllocationsDBFieldName.RefundId, true

	case string(RefundAllocationsDBFieldName.AllocationType):
		return RefundAllocationsDBFieldName.AllocationType, true

	case string(RefundAllocationsDBFieldName.ResponsiblePartyId):
		return RefundAllocationsDBFieldName.ResponsiblePartyId, true

	case string(RefundAllocationsDBFieldName.Amount):
		return RefundAllocationsDBFieldName.Amount, true

	case string(RefundAllocationsDBFieldName.CurrencyCode):
		return RefundAllocationsDBFieldName.CurrencyCode, true

	case string(RefundAllocationsDBFieldName.AllocationStatus):
		return RefundAllocationsDBFieldName.AllocationStatus, true

	case string(RefundAllocationsDBFieldName.Metadata):
		return RefundAllocationsDBFieldName.Metadata, true

	case string(RefundAllocationsDBFieldName.MetaCreatedAt):
		return RefundAllocationsDBFieldName.MetaCreatedAt, true

	case string(RefundAllocationsDBFieldName.MetaCreatedBy):
		return RefundAllocationsDBFieldName.MetaCreatedBy, true

	case string(RefundAllocationsDBFieldName.MetaUpdatedAt):
		return RefundAllocationsDBFieldName.MetaUpdatedAt, true

	case string(RefundAllocationsDBFieldName.MetaUpdatedBy):
		return RefundAllocationsDBFieldName.MetaUpdatedBy, true

	case string(RefundAllocationsDBFieldName.MetaDeletedAt):
		return RefundAllocationsDBFieldName.MetaDeletedAt, true

	case string(RefundAllocationsDBFieldName.MetaDeletedBy):
		return RefundAllocationsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var RefundAllocationsFilterJoins = map[string]JoinSpec{}

var RefundAllocationsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refund_id": {
		SourcePath:        "refund_id",
		DefaultOutputPath: "refundId",
		Column:            "refund_id",
		SQLAlias:          "refund_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"allocation_type": {
		SourcePath:        "allocation_type",
		DefaultOutputPath: "allocationType",
		Column:            "allocation_type",
		SQLAlias:          "allocation_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"responsible_party_id": {
		SourcePath:        "responsible_party_id",
		DefaultOutputPath: "responsiblePartyId",
		Column:            "responsible_party_id",
		SQLAlias:          "responsible_party_id",
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
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"allocation_status": {
		SourcePath:        "allocation_status",
		DefaultOutputPath: "allocationStatus",
		Column:            "allocation_status",
		SQLAlias:          "allocation_status",
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

func NewRefundAllocationsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = RefundAllocationsFilterFields[field]
	return
}

type RefundAllocationsFilterResult struct {
	RefundAllocations
	FilterCount int `db:"count"`
}

func ValidateRefundAllocationsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewRefundAllocationsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewRefundAllocationsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewRefundAllocationsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateRefundAllocationsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateRefundAllocationsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewRefundAllocationsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateRefundAllocationsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type AllocationStatus string

const (
	AllocationStatusComputed AllocationStatus = "computed"
	AllocationStatusPosted   AllocationStatus = "posted"
	AllocationStatusReversed AllocationStatus = "reversed"
)

type AllocationType string

const (
	AllocationTypeMerchantRecovery    AllocationType = "merchant_recovery"
	AllocationTypePlatformFeeReversal AllocationType = "platform_fee_reversal"
	AllocationTypeTaxReversal         AllocationType = "tax_reversal"
	AllocationTypeProviderFeeLoss     AllocationType = "provider_fee_loss"
	AllocationTypeCustomerCredit      AllocationType = "customer_credit"
)

type RefundAllocations struct {
	Id                 uuid.UUID        `db:"id"`
	RefundId           uuid.UUID        `db:"refund_id"`
	AllocationType     AllocationType   `db:"allocation_type"`
	ResponsiblePartyId nuuid.NUUID      `db:"responsible_party_id"`
	Amount             decimal.Decimal  `db:"amount"`
	CurrencyCode       string           `db:"currency_code"`
	AllocationStatus   AllocationStatus `db:"allocation_status"`
	Metadata           json.RawMessage  `db:"metadata"`

	shared.MetaSignature
}
type RefundAllocationsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RefundAllocations) ToRefundAllocationsPrimaryID() RefundAllocationsPrimaryID {
	return RefundAllocationsPrimaryID{
		Id: d.Id,
	}
}

type RefundAllocationsList []*RefundAllocations

type RefundAllocationsFilterResultList []*RefundAllocationsFilterResult
