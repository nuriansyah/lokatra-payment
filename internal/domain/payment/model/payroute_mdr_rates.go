package model

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type PayrouteMdrRates struct {
	Id                uuid.UUID           `db:"id"`
	ProviderAccountId uuid.UUID           `db:"provider_account_id"`
	PaymentMethodId   uuid.UUID           `db:"payment_method_id"`
	PaymentChannelId  *uuid.UUID          `db:"payment_channel_id"`
	Percentage        decimal.Decimal     `db:"percentage"`
	FixedFee          decimal.Decimal     `db:"fixed_fee"`
	MinAmount         decimal.NullDecimal `db:"min_amount"`
	MaxAmount         decimal.NullDecimal `db:"max_amount"`
	Currency          string              `db:"currency"`
	EffectiveFrom     null.Time           `db:"effective_from"`
	EffectiveTo       null.Time           `db:"effective_to"`

	shared.MetaSignature
}

type PayrouteMdrRatesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayrouteMdrRates) ToPayrouteMdrRatesPrimaryID() PayrouteMdrRatesPrimaryID {
	return PayrouteMdrRatesPrimaryID{Id: d.Id}
}

func (d PayrouteMdrRates) CalculateFee(amount decimal.Decimal) decimal.Decimal {
	pctFee := amount.Mul(d.Percentage).Div(decimal.NewFromInt(100))
	return pctFee.Add(d.FixedFee)
}

type PayrouteMdrRatesDBFieldNameType string

type payrouteMdrRatesDBFieldName struct {
	Id                PayrouteMdrRatesDBFieldNameType
	ProviderAccountId PayrouteMdrRatesDBFieldNameType
	PaymentMethodId   PayrouteMdrRatesDBFieldNameType
	PaymentChannelId  PayrouteMdrRatesDBFieldNameType
	Percentage        PayrouteMdrRatesDBFieldNameType
	FixedFee          PayrouteMdrRatesDBFieldNameType
	Currency          PayrouteMdrRatesDBFieldNameType
	EffectiveFrom     PayrouteMdrRatesDBFieldNameType
	EffectiveTo       PayrouteMdrRatesDBFieldNameType
	MetaCreatedAt     PayrouteMdrRatesDBFieldNameType
	MetaCreatedBy     PayrouteMdrRatesDBFieldNameType
	MetaUpdatedAt     PayrouteMdrRatesDBFieldNameType
	MetaUpdatedBy     PayrouteMdrRatesDBFieldNameType
	MetaDeletedAt     PayrouteMdrRatesDBFieldNameType
	MetaDeletedBy     PayrouteMdrRatesDBFieldNameType
}

var PayrouteMdrRatesDBFieldName = payrouteMdrRatesDBFieldName{
	Id:                "id",
	ProviderAccountId: "provider_account_id",
	PaymentMethodId:   "payment_method_id",
	PaymentChannelId:  "payment_channel_id",
	Percentage:        "percentage",
	FixedFee:          "fixed_fee",
	Currency:          "currency",
	EffectiveFrom:     "effective_from",
	EffectiveTo:       "effective_to",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewPayrouteMdrRatesDBFieldNameFromStr(field string) (dbField PayrouteMdrRatesDBFieldNameType, found bool) {
	switch field {
	case string(PayrouteMdrRatesDBFieldName.Id):
		return PayrouteMdrRatesDBFieldName.Id, true
	case string(PayrouteMdrRatesDBFieldName.ProviderAccountId):
		return PayrouteMdrRatesDBFieldName.ProviderAccountId, true
	case string(PayrouteMdrRatesDBFieldName.Percentage):
		return PayrouteMdrRatesDBFieldName.Percentage, true
	case string(PayrouteMdrRatesDBFieldName.MetaCreatedAt):
		return PayrouteMdrRatesDBFieldName.MetaCreatedAt, true
	}
	return "unknown", false
}

var PayrouteMdrRatesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath: "id", DefaultOutputPath: "id", Column: "id",
		SQLAlias: "id", Selectable: true, Filterable: true, Sortable: true,
	},
	"provider_account_id": {
		SourcePath: "provider_account_id", DefaultOutputPath: "providerAccountId", Column: "provider_account_id",
		SQLAlias: "provider_account_id", Selectable: true, Filterable: true, Sortable: true,
	},
	"payment_method_id": {
		SourcePath: "payment_method_id", DefaultOutputPath: "paymentMethodId", Column: "payment_method_id",
		SQLAlias: "payment_method_id", Selectable: true, Filterable: true, Sortable: true,
	},
	"percentage": {
		SourcePath: "percentage", DefaultOutputPath: "percentage", Column: "percentage",
		SQLAlias: "percentage", Selectable: true, Filterable: true, Sortable: true,
	},
	"currency": {
		SourcePath: "currency", DefaultOutputPath: "currency", Column: "currency",
		SQLAlias: "currency", Selectable: true, Filterable: true, Sortable: true,
	},
	"effective_from": {
		SourcePath: "effective_from", DefaultOutputPath: "effectiveFrom", Column: "effective_from",
		SQLAlias: "effective_from", Selectable: true, Filterable: true, Sortable: true,
	},
	"effective_to": {
		SourcePath: "effective_to", DefaultOutputPath: "effectiveTo", Column: "effective_to",
		SQLAlias: "effective_to", Selectable: true, Filterable: true, Sortable: true,
	},
	"meta_created_at": {
		SourcePath: "meta_created_at", DefaultOutputPath: "metaCreatedAt", Column: "meta_created_at",
		SQLAlias: "meta_created_at", Selectable: true, Filterable: true, Sortable: true,
	},
}

func NewPayrouteMdrRatesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayrouteMdrRatesFilterFields[field]
	return
}

func ValidatePayrouteMdrRatesFieldNameFilter(filter Filter) error {
	for _, field := range filter.FilterFields {
		spec, exist := NewPayrouteMdrRatesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			return failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
		}
	}
	return nil
}

type PayrouteMdrRatesList []*PayrouteMdrRates

type PayrouteMdrRatesFilterResult struct {
	PayrouteMdrRates
	FilterCount int `db:"count"`
}

type PayrouteMdrRatesFilterResultList []*PayrouteMdrRatesFilterResult
