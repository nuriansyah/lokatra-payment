package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type PaymentIntentItemsDBFieldNameType string

type paymentIntentItemsDBFieldName struct {
	Id               PaymentIntentItemsDBFieldNameType
	IntentId         PaymentIntentItemsDBFieldNameType
	ProductId        PaymentIntentItemsDBFieldNameType
	ProductType      PaymentIntentItemsDBFieldNameType
	ProductName      PaymentIntentItemsDBFieldNameType
	Quantity         PaymentIntentItemsDBFieldNameType
	UnitPrice        PaymentIntentItemsDBFieldNameType
	DiscountAmount   PaymentIntentItemsDBFieldNameType
	TotalPrice       PaymentIntentItemsDBFieldNameType
	SellerMerchantId PaymentIntentItemsDBFieldNameType
	Metadata         PaymentIntentItemsDBFieldNameType
	MetaCreatedAt    PaymentIntentItemsDBFieldNameType
	MetaCreatedBy    PaymentIntentItemsDBFieldNameType
	MetaUpdatedAt    PaymentIntentItemsDBFieldNameType
	MetaUpdatedBy    PaymentIntentItemsDBFieldNameType
	MetaDeletedAt    PaymentIntentItemsDBFieldNameType
	MetaDeletedBy    PaymentIntentItemsDBFieldNameType
}

var PaymentIntentItemsDBFieldName = paymentIntentItemsDBFieldName{
	Id:               "id",
	IntentId:         "intent_id",
	ProductId:        "product_id",
	ProductType:      "product_type",
	ProductName:      "product_name",
	Quantity:         "quantity",
	UnitPrice:        "unit_price",
	DiscountAmount:   "discount_amount",
	TotalPrice:       "total_price",
	SellerMerchantId: "seller_merchant_id",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewPaymentIntentItemsDBFieldNameFromStr(field string) (dbField PaymentIntentItemsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentIntentItemsDBFieldName.Id):
		return PaymentIntentItemsDBFieldName.Id, true

	case string(PaymentIntentItemsDBFieldName.IntentId):
		return PaymentIntentItemsDBFieldName.IntentId, true

	case string(PaymentIntentItemsDBFieldName.ProductId):
		return PaymentIntentItemsDBFieldName.ProductId, true

	case string(PaymentIntentItemsDBFieldName.ProductType):
		return PaymentIntentItemsDBFieldName.ProductType, true

	case string(PaymentIntentItemsDBFieldName.ProductName):
		return PaymentIntentItemsDBFieldName.ProductName, true

	case string(PaymentIntentItemsDBFieldName.Quantity):
		return PaymentIntentItemsDBFieldName.Quantity, true

	case string(PaymentIntentItemsDBFieldName.UnitPrice):
		return PaymentIntentItemsDBFieldName.UnitPrice, true

	case string(PaymentIntentItemsDBFieldName.DiscountAmount):
		return PaymentIntentItemsDBFieldName.DiscountAmount, true

	case string(PaymentIntentItemsDBFieldName.TotalPrice):
		return PaymentIntentItemsDBFieldName.TotalPrice, true

	case string(PaymentIntentItemsDBFieldName.SellerMerchantId):
		return PaymentIntentItemsDBFieldName.SellerMerchantId, true

	case string(PaymentIntentItemsDBFieldName.Metadata):
		return PaymentIntentItemsDBFieldName.Metadata, true

	case string(PaymentIntentItemsDBFieldName.MetaCreatedAt):
		return PaymentIntentItemsDBFieldName.MetaCreatedAt, true

	case string(PaymentIntentItemsDBFieldName.MetaCreatedBy):
		return PaymentIntentItemsDBFieldName.MetaCreatedBy, true

	case string(PaymentIntentItemsDBFieldName.MetaUpdatedAt):
		return PaymentIntentItemsDBFieldName.MetaUpdatedAt, true

	case string(PaymentIntentItemsDBFieldName.MetaUpdatedBy):
		return PaymentIntentItemsDBFieldName.MetaUpdatedBy, true

	case string(PaymentIntentItemsDBFieldName.MetaDeletedAt):
		return PaymentIntentItemsDBFieldName.MetaDeletedAt, true

	case string(PaymentIntentItemsDBFieldName.MetaDeletedBy):
		return PaymentIntentItemsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type PaymentIntentItemsFilterResult struct {
	PaymentIntentItems
	FilterCount int `db:"count"`
}

func ValidatePaymentIntentItemsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewPaymentIntentItemsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewPaymentIntentItemsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewPaymentIntentItemsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type PaymentIntentItems struct {
	Id               uuid.UUID       `db:"id"`
	IntentId         uuid.UUID       `db:"intent_id"`
	ProductId        uuid.UUID       `db:"product_id"`
	ProductType      string          `db:"product_type"`
	ProductName      string          `db:"product_name"`
	Quantity         int             `db:"quantity"`
	UnitPrice        decimal.Decimal `db:"unit_price"`
	DiscountAmount   decimal.Decimal `db:"discount_amount"`
	TotalPrice       decimal.Decimal `db:"total_price"`
	SellerMerchantId nuuid.NUUID     `db:"seller_merchant_id"`
	Metadata         json.RawMessage `db:"metadata"`
	MetaCreatedAt    time.Time       `db:"meta_created_at"`
	MetaCreatedBy    uuid.UUID       `db:"meta_created_by"`
	MetaUpdatedAt    time.Time       `db:"meta_updated_at"`
	MetaUpdatedBy    nuuid.NUUID     `db:"meta_updated_by"`
	MetaDeletedAt    null.Time       `db:"meta_deleted_at"`
	MetaDeletedBy    nuuid.NUUID     `db:"meta_deleted_by"`
}
type PaymentIntentItemsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentIntentItems) ToPaymentIntentItemsPrimaryID() PaymentIntentItemsPrimaryID {
	return PaymentIntentItemsPrimaryID{
		Id: d.Id,
	}
}

type PaymentIntentItemsList []*PaymentIntentItems
