package model

import (
	"encoding/json"
	"fmt"
	"net/netip"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type PaymentIntentsDBFieldNameType string

type paymentIntentsDBFieldName struct {
	Id                  PaymentIntentsDBFieldNameType
	IntentCode          PaymentIntentsDBFieldNameType
	MerchantId          PaymentIntentsDBFieldNameType
	OrderId             PaymentIntentsDBFieldNameType
	OrderType           PaymentIntentsDBFieldNameType
	Amount              PaymentIntentsDBFieldNameType
	Currency            PaymentIntentsDBFieldNameType
	TaxAmount           PaymentIntentsDBFieldNameType
	DiscountAmount      PaymentIntentsDBFieldNameType
	TipAmount           PaymentIntentsDBFieldNameType
	UserId              PaymentIntentsDBFieldNameType
	CustomerName        PaymentIntentsDBFieldNameType
	CustomerEmail       PaymentIntentsDBFieldNameType
	CustomerPhone       PaymentIntentsDBFieldNameType
	CustomerIp          PaymentIntentsDBFieldNameType
	CustomerCountry     PaymentIntentsDBFieldNameType
	PaymentMethodId     PaymentIntentsDBFieldNameType
	PaymentMethodType   PaymentIntentsDBFieldNameType
	Status              PaymentIntentsDBFieldNameType
	RoutingProfileId    PaymentIntentsDBFieldNameType
	ExpiresAt           PaymentIntentsDBFieldNameType
	Requires3ds         PaymentIntentsDBFieldNameType
	ThreeDsVersion      PaymentIntentsDBFieldNameType
	Description         PaymentIntentsDBFieldNameType
	StatementDescriptor PaymentIntentsDBFieldNameType
	Metadata            PaymentIntentsDBFieldNameType
	PromoCode           PaymentIntentsDBFieldNameType
	PromoDiscountAmount PaymentIntentsDBFieldNameType
	IdempotencyKeyId    PaymentIntentsDBFieldNameType
	ConfirmedAt         PaymentIntentsDBFieldNameType
	CancelledAt         PaymentIntentsDBFieldNameType
	CancellationReason  PaymentIntentsDBFieldNameType
	MetaCreatedAt       PaymentIntentsDBFieldNameType
	MetaCreatedBy       PaymentIntentsDBFieldNameType
	MetaUpdatedAt       PaymentIntentsDBFieldNameType
	MetaUpdatedBy       PaymentIntentsDBFieldNameType
	MetaDeletedAt       PaymentIntentsDBFieldNameType
	MetaDeletedBy       PaymentIntentsDBFieldNameType
}

var PaymentIntentsDBFieldName = paymentIntentsDBFieldName{
	Id:                  "id",
	IntentCode:          "intent_code",
	MerchantId:          "merchant_id",
	OrderId:             "order_id",
	OrderType:           "order_type",
	Amount:              "amount",
	Currency:            "currency",
	TaxAmount:           "tax_amount",
	DiscountAmount:      "discount_amount",
	TipAmount:           "tip_amount",
	UserId:              "user_id",
	CustomerName:        "customer_name",
	CustomerEmail:       "customer_email",
	CustomerPhone:       "customer_phone",
	CustomerIp:          "customer_ip",
	CustomerCountry:     "customer_country",
	PaymentMethodId:     "payment_method_id",
	PaymentMethodType:   "payment_method_type",
	Status:              "status",
	RoutingProfileId:    "routing_profile_id",
	ExpiresAt:           "expires_at",
	Requires3ds:         "requires_3ds",
	ThreeDsVersion:      "three_ds_version",
	Description:         "description",
	StatementDescriptor: "statement_descriptor",
	Metadata:            "metadata",
	PromoCode:           "promo_code",
	PromoDiscountAmount: "promo_discount_amount",
	IdempotencyKeyId:    "idempotency_key_id",
	ConfirmedAt:         "confirmed_at",
	CancelledAt:         "cancelled_at",
	CancellationReason:  "cancellation_reason",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewPaymentIntentsDBFieldNameFromStr(field string) (dbField PaymentIntentsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentIntentsDBFieldName.Id):
		return PaymentIntentsDBFieldName.Id, true

	case string(PaymentIntentsDBFieldName.IntentCode):
		return PaymentIntentsDBFieldName.IntentCode, true

	case string(PaymentIntentsDBFieldName.MerchantId):
		return PaymentIntentsDBFieldName.MerchantId, true

	case string(PaymentIntentsDBFieldName.OrderId):
		return PaymentIntentsDBFieldName.OrderId, true

	case string(PaymentIntentsDBFieldName.OrderType):
		return PaymentIntentsDBFieldName.OrderType, true

	case string(PaymentIntentsDBFieldName.Amount):
		return PaymentIntentsDBFieldName.Amount, true

	case string(PaymentIntentsDBFieldName.Currency):
		return PaymentIntentsDBFieldName.Currency, true

	case string(PaymentIntentsDBFieldName.TaxAmount):
		return PaymentIntentsDBFieldName.TaxAmount, true

	case string(PaymentIntentsDBFieldName.DiscountAmount):
		return PaymentIntentsDBFieldName.DiscountAmount, true

	case string(PaymentIntentsDBFieldName.TipAmount):
		return PaymentIntentsDBFieldName.TipAmount, true

	case string(PaymentIntentsDBFieldName.UserId):
		return PaymentIntentsDBFieldName.UserId, true

	case string(PaymentIntentsDBFieldName.CustomerName):
		return PaymentIntentsDBFieldName.CustomerName, true

	case string(PaymentIntentsDBFieldName.CustomerEmail):
		return PaymentIntentsDBFieldName.CustomerEmail, true

	case string(PaymentIntentsDBFieldName.CustomerPhone):
		return PaymentIntentsDBFieldName.CustomerPhone, true

	case string(PaymentIntentsDBFieldName.CustomerIp):
		return PaymentIntentsDBFieldName.CustomerIp, true

	case string(PaymentIntentsDBFieldName.CustomerCountry):
		return PaymentIntentsDBFieldName.CustomerCountry, true

	case string(PaymentIntentsDBFieldName.PaymentMethodId):
		return PaymentIntentsDBFieldName.PaymentMethodId, true

	case string(PaymentIntentsDBFieldName.PaymentMethodType):
		return PaymentIntentsDBFieldName.PaymentMethodType, true

	case string(PaymentIntentsDBFieldName.Status):
		return PaymentIntentsDBFieldName.Status, true

	case string(PaymentIntentsDBFieldName.RoutingProfileId):
		return PaymentIntentsDBFieldName.RoutingProfileId, true

	case string(PaymentIntentsDBFieldName.ExpiresAt):
		return PaymentIntentsDBFieldName.ExpiresAt, true

	case string(PaymentIntentsDBFieldName.Requires3ds):
		return PaymentIntentsDBFieldName.Requires3ds, true

	case string(PaymentIntentsDBFieldName.ThreeDsVersion):
		return PaymentIntentsDBFieldName.ThreeDsVersion, true

	case string(PaymentIntentsDBFieldName.Description):
		return PaymentIntentsDBFieldName.Description, true

	case string(PaymentIntentsDBFieldName.StatementDescriptor):
		return PaymentIntentsDBFieldName.StatementDescriptor, true

	case string(PaymentIntentsDBFieldName.Metadata):
		return PaymentIntentsDBFieldName.Metadata, true

	case string(PaymentIntentsDBFieldName.PromoCode):
		return PaymentIntentsDBFieldName.PromoCode, true

	case string(PaymentIntentsDBFieldName.PromoDiscountAmount):
		return PaymentIntentsDBFieldName.PromoDiscountAmount, true

	case string(PaymentIntentsDBFieldName.IdempotencyKeyId):
		return PaymentIntentsDBFieldName.IdempotencyKeyId, true

	case string(PaymentIntentsDBFieldName.ConfirmedAt):
		return PaymentIntentsDBFieldName.ConfirmedAt, true

	case string(PaymentIntentsDBFieldName.CancelledAt):
		return PaymentIntentsDBFieldName.CancelledAt, true

	case string(PaymentIntentsDBFieldName.CancellationReason):
		return PaymentIntentsDBFieldName.CancellationReason, true

	case string(PaymentIntentsDBFieldName.MetaCreatedAt):
		return PaymentIntentsDBFieldName.MetaCreatedAt, true

	case string(PaymentIntentsDBFieldName.MetaCreatedBy):
		return PaymentIntentsDBFieldName.MetaCreatedBy, true

	case string(PaymentIntentsDBFieldName.MetaUpdatedAt):
		return PaymentIntentsDBFieldName.MetaUpdatedAt, true

	case string(PaymentIntentsDBFieldName.MetaUpdatedBy):
		return PaymentIntentsDBFieldName.MetaUpdatedBy, true

	case string(PaymentIntentsDBFieldName.MetaDeletedAt):
		return PaymentIntentsDBFieldName.MetaDeletedAt, true

	case string(PaymentIntentsDBFieldName.MetaDeletedBy):
		return PaymentIntentsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type PaymentIntentsFilterResult struct {
	PaymentIntents
	FilterCount int `db:"count"`
}

func ValidatePaymentIntentsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewPaymentIntentsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewPaymentIntentsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewPaymentIntentsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type PaymentIntents struct {
	Id                  uuid.UUID         `db:"id"`
	IntentCode          string            `db:"intent_code"`
	MerchantId          uuid.UUID         `db:"merchant_id"`
	OrderId             nuuid.NUUID       `db:"order_id"`
	OrderType           null.String       `db:"order_type"`
	Amount              decimal.Decimal   `db:"amount"`
	Currency            PaymentCurrency   `db:"currency"`
	TaxAmount           decimal.Decimal   `db:"tax_amount"`
	DiscountAmount      decimal.Decimal   `db:"discount_amount"`
	TipAmount           decimal.Decimal   `db:"tip_amount"`
	UserId              nuuid.NUUID       `db:"user_id"`
	CustomerName        null.String       `db:"customer_name"`
	CustomerEmail       null.String       `db:"customer_email"`
	CustomerPhone       null.String       `db:"customer_phone"`
	CustomerIp          *netip.Addr       `db:"customer_ip"`
	CustomerCountry     null.String       `db:"customer_country"`
	PaymentMethodId     nuuid.NUUID       `db:"payment_method_id"`
	PaymentMethodType   PaymentMethodType `db:"payment_method_type"`
	Status              PaymentStatus     `db:"status"`
	RoutingProfileId    nuuid.NUUID       `db:"routing_profile_id"`
	ExpiresAt           time.Time         `db:"expires_at"`
	Requires3ds         bool              `db:"requires_3ds"`
	ThreeDsVersion      null.String       `db:"three_ds_version"`
	Description         null.String       `db:"description"`
	StatementDescriptor null.String       `db:"statement_descriptor"`
	Metadata            json.RawMessage   `db:"metadata"`
	PromoCode           null.String       `db:"promo_code"`
	PromoDiscountAmount decimal.Decimal   `db:"promo_discount_amount"`
	IdempotencyKeyId    nuuid.NUUID       `db:"idempotency_key_id"`
	ConfirmedAt         null.Time         `db:"confirmed_at"`
	CancelledAt         null.Time         `db:"cancelled_at"`
	CancellationReason  null.String       `db:"cancellation_reason"`
	MetaCreatedAt       time.Time         `db:"meta_created_at"`
	MetaCreatedBy       uuid.UUID         `db:"meta_created_by"`
	MetaUpdatedAt       time.Time         `db:"meta_updated_at"`
	MetaUpdatedBy       nuuid.NUUID       `db:"meta_updated_by"`
	MetaDeletedAt       null.Time         `db:"meta_deleted_at"`
	MetaDeletedBy       nuuid.NUUID       `db:"meta_deleted_by"`
}
type PaymentIntentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentIntents) ToPaymentIntentsPrimaryID() PaymentIntentsPrimaryID {
	return PaymentIntentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentIntentsList []*PaymentIntents
