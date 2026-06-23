package model

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
	"time"
)

type MerchantBalanceSummaryDBFieldNameType string

type merchantBalanceSummaryDBFieldName struct {
	MerchantPartyId  MerchantBalanceSummaryDBFieldNameType
	CurrencyCode     MerchantBalanceSummaryDBFieldNameType
	PendingAmount    MerchantBalanceSummaryDBFieldNameType
	AvailableAmount  MerchantBalanceSummaryDBFieldNameType
	ReservedAmount   MerchantBalanceSummaryDBFieldNameType
	PayableAmount    MerchantBalanceSummaryDBFieldNameType
	PaidOutAmount    MerchantBalanceSummaryDBFieldNameType
	NegativeAmount   MerchantBalanceSummaryDBFieldNameType
	RefundableAmount MerchantBalanceSummaryDBFieldNameType
	RefreshedAt      MerchantBalanceSummaryDBFieldNameType
}

var MerchantBalanceSummaryDBFieldName = merchantBalanceSummaryDBFieldName{
	MerchantPartyId:  "merchant_party_id",
	CurrencyCode:     "currency_code",
	PendingAmount:    "pending_amount",
	AvailableAmount:  "available_amount",
	ReservedAmount:   "reserved_amount",
	PayableAmount:    "payable_amount",
	PaidOutAmount:    "paid_out_amount",
	NegativeAmount:   "negative_amount",
	RefundableAmount: "refundable_amount",
	RefreshedAt:      "refreshed_at",
}

func NewMerchantBalanceSummaryDBFieldNameFromStr(field string) (dbField MerchantBalanceSummaryDBFieldNameType, found bool) {
	switch field {

	case string(MerchantBalanceSummaryDBFieldName.MerchantPartyId):
		return MerchantBalanceSummaryDBFieldName.MerchantPartyId, true

	case string(MerchantBalanceSummaryDBFieldName.CurrencyCode):
		return MerchantBalanceSummaryDBFieldName.CurrencyCode, true

	case string(MerchantBalanceSummaryDBFieldName.PendingAmount):
		return MerchantBalanceSummaryDBFieldName.PendingAmount, true

	case string(MerchantBalanceSummaryDBFieldName.AvailableAmount):
		return MerchantBalanceSummaryDBFieldName.AvailableAmount, true

	case string(MerchantBalanceSummaryDBFieldName.ReservedAmount):
		return MerchantBalanceSummaryDBFieldName.ReservedAmount, true

	case string(MerchantBalanceSummaryDBFieldName.PayableAmount):
		return MerchantBalanceSummaryDBFieldName.PayableAmount, true

	case string(MerchantBalanceSummaryDBFieldName.PaidOutAmount):
		return MerchantBalanceSummaryDBFieldName.PaidOutAmount, true

	case string(MerchantBalanceSummaryDBFieldName.NegativeAmount):
		return MerchantBalanceSummaryDBFieldName.NegativeAmount, true

	case string(MerchantBalanceSummaryDBFieldName.RefundableAmount):
		return MerchantBalanceSummaryDBFieldName.RefundableAmount, true

	case string(MerchantBalanceSummaryDBFieldName.RefreshedAt):
		return MerchantBalanceSummaryDBFieldName.RefreshedAt, true

	}
	return "unknown", false
}

var MerchantBalanceSummaryFilterJoins = map[string]JoinSpec{}

var MerchantBalanceSummaryFilterFields = map[string]FilterFieldSpec{
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
	"pending_amount": {
		SourcePath:        "pending_amount",
		DefaultOutputPath: "pendingAmount",
		Column:            "pending_amount",
		SQLAlias:          "pending_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"available_amount": {
		SourcePath:        "available_amount",
		DefaultOutputPath: "availableAmount",
		Column:            "available_amount",
		SQLAlias:          "available_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reserved_amount": {
		SourcePath:        "reserved_amount",
		DefaultOutputPath: "reservedAmount",
		Column:            "reserved_amount",
		SQLAlias:          "reserved_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payable_amount": {
		SourcePath:        "payable_amount",
		DefaultOutputPath: "payableAmount",
		Column:            "payable_amount",
		SQLAlias:          "payable_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"paid_out_amount": {
		SourcePath:        "paid_out_amount",
		DefaultOutputPath: "paidOutAmount",
		Column:            "paid_out_amount",
		SQLAlias:          "paid_out_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"negative_amount": {
		SourcePath:        "negative_amount",
		DefaultOutputPath: "negativeAmount",
		Column:            "negative_amount",
		SQLAlias:          "negative_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refundable_amount": {
		SourcePath:        "refundable_amount",
		DefaultOutputPath: "refundableAmount",
		Column:            "refundable_amount",
		SQLAlias:          "refundable_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refreshed_at": {
		SourcePath:        "refreshed_at",
		DefaultOutputPath: "refreshedAt",
		Column:            "refreshed_at",
		SQLAlias:          "refreshed_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
}

func NewMerchantBalanceSummaryFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = MerchantBalanceSummaryFilterFields[field]
	return
}

type MerchantBalanceSummaryFilterResult struct {
	MerchantBalanceSummary
	FilterCount int `db:"count"`
}

func ValidateMerchantBalanceSummaryFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewMerchantBalanceSummaryFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewMerchantBalanceSummaryFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewMerchantBalanceSummaryFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateMerchantBalanceSummaryFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateMerchantBalanceSummaryFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewMerchantBalanceSummaryFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateMerchantBalanceSummaryFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type MerchantBalanceSummary struct {
	MerchantPartyId  uuid.UUID       `db:"merchant_party_id"`
	CurrencyCode     string          `db:"currency_code"`
	PendingAmount    decimal.Decimal `db:"pending_amount"`
	AvailableAmount  decimal.Decimal `db:"available_amount"`
	ReservedAmount   decimal.Decimal `db:"reserved_amount"`
	PayableAmount    decimal.Decimal `db:"payable_amount"`
	PaidOutAmount    decimal.Decimal `db:"paid_out_amount"`
	NegativeAmount   decimal.Decimal `db:"negative_amount"`
	RefundableAmount decimal.Decimal `db:"refundable_amount"`
	RefreshedAt      time.Time       `db:"refreshed_at"`
}
type MerchantBalanceSummaryPrimaryID struct {
	MerchantPartyId uuid.UUID `db:"merchant_party_id"`
	CurrencyCode    string    `db:"currency_code"`
}

func (d MerchantBalanceSummary) ToMerchantBalanceSummaryPrimaryID() MerchantBalanceSummaryPrimaryID {
	return MerchantBalanceSummaryPrimaryID{
		MerchantPartyId: d.MerchantPartyId,
		CurrencyCode:    d.CurrencyCode,
	}
}

type MerchantBalanceSummaryList []*MerchantBalanceSummary

type MerchantBalanceSummaryFilterResultList []*MerchantBalanceSummaryFilterResult
