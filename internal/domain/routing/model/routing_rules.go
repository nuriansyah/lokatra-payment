package model

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type RoutingRulesDBFieldNameType string

type routingRulesDBFieldName struct {
	Id                 RoutingRulesDBFieldNameType
	ProfileId          RoutingRulesDBFieldNameType
	Priority           RoutingRulesDBFieldNameType
	Name               RoutingRulesDBFieldNameType
	IsActive           RoutingRulesDBFieldNameType
	MatchPaymentMethod RoutingRulesDBFieldNameType
	MatchCurrency      RoutingRulesDBFieldNameType
	MatchAmountMin     RoutingRulesDBFieldNameType
	MatchAmountMax     RoutingRulesDBFieldNameType
	MatchUserCountry   RoutingRulesDBFieldNameType
	MatchCardBin       RoutingRulesDBFieldNameType
	MatchProductType   RoutingRulesDBFieldNameType
	CostWeight         RoutingRulesDBFieldNameType
	Notes              RoutingRulesDBFieldNameType
	MetaCreatedAt      RoutingRulesDBFieldNameType
	MetaCreatedBy      RoutingRulesDBFieldNameType
	MetaUpdatedAt      RoutingRulesDBFieldNameType
	MetaUpdatedBy      RoutingRulesDBFieldNameType
	MetaDeletedAt      RoutingRulesDBFieldNameType
	MetaDeletedBy      RoutingRulesDBFieldNameType
}

var RoutingRulesDBFieldName = routingRulesDBFieldName{
	Id:                 "id",
	ProfileId:          "profile_id",
	Priority:           "priority",
	Name:               "name",
	IsActive:           "is_active",
	MatchPaymentMethod: "match_payment_method",
	MatchCurrency:      "match_currency",
	MatchAmountMin:     "match_amount_min",
	MatchAmountMax:     "match_amount_max",
	MatchUserCountry:   "match_user_country",
	MatchCardBin:       "match_card_bin",
	MatchProductType:   "match_product_type",
	CostWeight:         "cost_weight",
	Notes:              "notes",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewRoutingRulesDBFieldNameFromStr(field string) (dbField RoutingRulesDBFieldNameType, found bool) {
	switch field {

	case string(RoutingRulesDBFieldName.Id):
		return RoutingRulesDBFieldName.Id, true

	case string(RoutingRulesDBFieldName.ProfileId):
		return RoutingRulesDBFieldName.ProfileId, true

	case string(RoutingRulesDBFieldName.Priority):
		return RoutingRulesDBFieldName.Priority, true

	case string(RoutingRulesDBFieldName.Name):
		return RoutingRulesDBFieldName.Name, true

	case string(RoutingRulesDBFieldName.IsActive):
		return RoutingRulesDBFieldName.IsActive, true

	case string(RoutingRulesDBFieldName.MatchPaymentMethod):
		return RoutingRulesDBFieldName.MatchPaymentMethod, true

	case string(RoutingRulesDBFieldName.MatchCurrency):
		return RoutingRulesDBFieldName.MatchCurrency, true

	case string(RoutingRulesDBFieldName.MatchAmountMin):
		return RoutingRulesDBFieldName.MatchAmountMin, true

	case string(RoutingRulesDBFieldName.MatchAmountMax):
		return RoutingRulesDBFieldName.MatchAmountMax, true

	case string(RoutingRulesDBFieldName.MatchUserCountry):
		return RoutingRulesDBFieldName.MatchUserCountry, true

	case string(RoutingRulesDBFieldName.MatchCardBin):
		return RoutingRulesDBFieldName.MatchCardBin, true

	case string(RoutingRulesDBFieldName.MatchProductType):
		return RoutingRulesDBFieldName.MatchProductType, true

	case string(RoutingRulesDBFieldName.CostWeight):
		return RoutingRulesDBFieldName.CostWeight, true

	case string(RoutingRulesDBFieldName.Notes):
		return RoutingRulesDBFieldName.Notes, true

	case string(RoutingRulesDBFieldName.MetaCreatedAt):
		return RoutingRulesDBFieldName.MetaCreatedAt, true

	case string(RoutingRulesDBFieldName.MetaCreatedBy):
		return RoutingRulesDBFieldName.MetaCreatedBy, true

	case string(RoutingRulesDBFieldName.MetaUpdatedAt):
		return RoutingRulesDBFieldName.MetaUpdatedAt, true

	case string(RoutingRulesDBFieldName.MetaUpdatedBy):
		return RoutingRulesDBFieldName.MetaUpdatedBy, true

	case string(RoutingRulesDBFieldName.MetaDeletedAt):
		return RoutingRulesDBFieldName.MetaDeletedAt, true

	case string(RoutingRulesDBFieldName.MetaDeletedBy):
		return RoutingRulesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type RoutingRulesFilterResult struct {
	RoutingRules
	FilterCount int `db:"count"`
}

func ValidateRoutingRulesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewRoutingRulesDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewRoutingRulesDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewRoutingRulesDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type PaymentCurrency string

const (
	PaymentCurrencyIdr PaymentCurrency = "IDR"
	PaymentCurrencyUsd PaymentCurrency = "USD"
	PaymentCurrencySgd PaymentCurrency = "SGD"
	PaymentCurrencyMyr PaymentCurrency = "MYR"
	PaymentCurrencyPhp PaymentCurrency = "PHP"
	PaymentCurrencyThb PaymentCurrency = "THB"
	PaymentCurrencyAed PaymentCurrency = "AED"
	PaymentCurrencyEur PaymentCurrency = "EUR"
	PaymentCurrencyGbp PaymentCurrency = "GBP"
	PaymentCurrencyJpy PaymentCurrency = "JPY"
)

type PaymentMethodType string

const (
	PaymentMethodTypeCard           PaymentMethodType = "CARD"
	PaymentMethodTypeVirtualAccount PaymentMethodType = "VIRTUAL_ACCOUNT"
	PaymentMethodTypeQris           PaymentMethodType = "QRIS"
	PaymentMethodTypeEwallet        PaymentMethodType = "EWALLET"
	PaymentMethodTypeDirectDebit    PaymentMethodType = "DIRECT_DEBIT"
	PaymentMethodTypeBankTransfer   PaymentMethodType = "BANK_TRANSFER"
	PaymentMethodTypePaylater       PaymentMethodType = "PAYLATER"
	PaymentMethodTypeVoucher        PaymentMethodType = "VOUCHER"
	PaymentMethodTypePoints         PaymentMethodType = "POINTS"
	PaymentMethodTypeCashOnDelivery PaymentMethodType = "CASH_ON_DELIVERY"
)

type RoutingRules struct {
	Id                 uuid.UUID           `db:"id"`
	ProfileId          uuid.UUID           `db:"profile_id"`
	Priority           int16               `db:"priority"`
	Name               string              `db:"name"`
	IsActive           bool                `db:"is_active"`
	MatchPaymentMethod PaymentMethodType   `db:"match_payment_method"`
	MatchCurrency      PaymentCurrency     `db:"match_currency"`
	MatchAmountMin     decimal.NullDecimal `db:"match_amount_min"`
	MatchAmountMax     decimal.NullDecimal `db:"match_amount_max"`
	MatchUserCountry   null.String         `db:"match_user_country"`
	MatchCardBin       null.String         `db:"match_card_bin"`
	MatchProductType   null.String         `db:"match_product_type"`
	CostWeight         decimal.NullDecimal `db:"cost_weight"`
	Notes              null.String         `db:"notes"`
	MetaCreatedAt      time.Time           `db:"meta_created_at"`
	MetaCreatedBy      uuid.UUID           `db:"meta_created_by"`
	MetaUpdatedAt      time.Time           `db:"meta_updated_at"`
	MetaUpdatedBy      nuuid.NUUID         `db:"meta_updated_by"`
	MetaDeletedAt      null.Time           `db:"meta_deleted_at"`
	MetaDeletedBy      nuuid.NUUID         `db:"meta_deleted_by"`
}
type RoutingRulesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RoutingRules) ToRoutingRulesPrimaryID() RoutingRulesPrimaryID {
	return RoutingRulesPrimaryID{
		Id: d.Id,
	}
}

type RoutingRulesList []*RoutingRules
