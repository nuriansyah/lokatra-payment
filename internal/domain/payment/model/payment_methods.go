package model

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type PaymentMethodsDBFieldNameType string

type paymentMethodsDBFieldName struct {
	Id                     PaymentMethodsDBFieldNameType
	UserId                 PaymentMethodsDBFieldNameType
	MerchantId             PaymentMethodsDBFieldNameType
	MethodType             PaymentMethodsDBFieldNameType
	Psp                    PaymentMethodsDBFieldNameType
	TokenRef               PaymentMethodsDBFieldNameType
	TokenType              PaymentMethodsDBFieldNameType
	TokenExpiresAt         PaymentMethodsDBFieldNameType
	CardBrand              PaymentMethodsDBFieldNameType
	CardLastFour           PaymentMethodsDBFieldNameType
	CardExpMonth           PaymentMethodsDBFieldNameType
	CardExpYear            PaymentMethodsDBFieldNameType
	CardCountry            PaymentMethodsDBFieldNameType
	CardFundingType        PaymentMethodsDBFieldNameType
	CardBin                PaymentMethodsDBFieldNameType
	WalletAccountRef       PaymentMethodsDBFieldNameType
	VaBankCode             PaymentMethodsDBFieldNameType
	DisplayLabel           PaymentMethodsDBFieldNameType
	IsDefault              PaymentMethodsDBFieldNameType
	IsActive               PaymentMethodsDBFieldNameType
	VerifiedAt             PaymentMethodsDBFieldNameType
	Fingerprint            PaymentMethodsDBFieldNameType
	GdprErasureRequestedAt PaymentMethodsDBFieldNameType
	GdprErasedAt           PaymentMethodsDBFieldNameType
	MetaCreatedAt          PaymentMethodsDBFieldNameType
	MetaCreatedBy          PaymentMethodsDBFieldNameType
	MetaUpdatedAt          PaymentMethodsDBFieldNameType
	MetaUpdatedBy          PaymentMethodsDBFieldNameType
	MetaDeletedAt          PaymentMethodsDBFieldNameType
	MetaDeletedBy          PaymentMethodsDBFieldNameType
}

var PaymentMethodsDBFieldName = paymentMethodsDBFieldName{
	Id:                     "id",
	UserId:                 "user_id",
	MerchantId:             "merchant_id",
	MethodType:             "method_type",
	Psp:                    "psp",
	TokenRef:               "token_ref",
	TokenType:              "token_type",
	TokenExpiresAt:         "token_expires_at",
	CardBrand:              "card_brand",
	CardLastFour:           "card_last_four",
	CardExpMonth:           "card_exp_month",
	CardExpYear:            "card_exp_year",
	CardCountry:            "card_country",
	CardFundingType:        "card_funding_type",
	CardBin:                "card_bin",
	WalletAccountRef:       "wallet_account_ref",
	VaBankCode:             "va_bank_code",
	DisplayLabel:           "display_label",
	IsDefault:              "is_default",
	IsActive:               "is_active",
	VerifiedAt:             "verified_at",
	Fingerprint:            "fingerprint",
	GdprErasureRequestedAt: "gdpr_erasure_requested_at",
	GdprErasedAt:           "gdpr_erased_at",
	MetaCreatedAt:          "meta_created_at",
	MetaCreatedBy:          "meta_created_by",
	MetaUpdatedAt:          "meta_updated_at",
	MetaUpdatedBy:          "meta_updated_by",
	MetaDeletedAt:          "meta_deleted_at",
	MetaDeletedBy:          "meta_deleted_by",
}

func NewPaymentMethodsDBFieldNameFromStr(field string) (dbField PaymentMethodsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentMethodsDBFieldName.Id):
		return PaymentMethodsDBFieldName.Id, true

	case string(PaymentMethodsDBFieldName.UserId):
		return PaymentMethodsDBFieldName.UserId, true

	case string(PaymentMethodsDBFieldName.MerchantId):
		return PaymentMethodsDBFieldName.MerchantId, true

	case string(PaymentMethodsDBFieldName.MethodType):
		return PaymentMethodsDBFieldName.MethodType, true

	case string(PaymentMethodsDBFieldName.Psp):
		return PaymentMethodsDBFieldName.Psp, true

	case string(PaymentMethodsDBFieldName.TokenRef):
		return PaymentMethodsDBFieldName.TokenRef, true

	case string(PaymentMethodsDBFieldName.TokenType):
		return PaymentMethodsDBFieldName.TokenType, true

	case string(PaymentMethodsDBFieldName.TokenExpiresAt):
		return PaymentMethodsDBFieldName.TokenExpiresAt, true

	case string(PaymentMethodsDBFieldName.CardBrand):
		return PaymentMethodsDBFieldName.CardBrand, true

	case string(PaymentMethodsDBFieldName.CardLastFour):
		return PaymentMethodsDBFieldName.CardLastFour, true

	case string(PaymentMethodsDBFieldName.CardExpMonth):
		return PaymentMethodsDBFieldName.CardExpMonth, true

	case string(PaymentMethodsDBFieldName.CardExpYear):
		return PaymentMethodsDBFieldName.CardExpYear, true

	case string(PaymentMethodsDBFieldName.CardCountry):
		return PaymentMethodsDBFieldName.CardCountry, true

	case string(PaymentMethodsDBFieldName.CardFundingType):
		return PaymentMethodsDBFieldName.CardFundingType, true

	case string(PaymentMethodsDBFieldName.CardBin):
		return PaymentMethodsDBFieldName.CardBin, true

	case string(PaymentMethodsDBFieldName.WalletAccountRef):
		return PaymentMethodsDBFieldName.WalletAccountRef, true

	case string(PaymentMethodsDBFieldName.VaBankCode):
		return PaymentMethodsDBFieldName.VaBankCode, true

	case string(PaymentMethodsDBFieldName.DisplayLabel):
		return PaymentMethodsDBFieldName.DisplayLabel, true

	case string(PaymentMethodsDBFieldName.IsDefault):
		return PaymentMethodsDBFieldName.IsDefault, true

	case string(PaymentMethodsDBFieldName.IsActive):
		return PaymentMethodsDBFieldName.IsActive, true

	case string(PaymentMethodsDBFieldName.VerifiedAt):
		return PaymentMethodsDBFieldName.VerifiedAt, true

	case string(PaymentMethodsDBFieldName.Fingerprint):
		return PaymentMethodsDBFieldName.Fingerprint, true

	case string(PaymentMethodsDBFieldName.GdprErasureRequestedAt):
		return PaymentMethodsDBFieldName.GdprErasureRequestedAt, true

	case string(PaymentMethodsDBFieldName.GdprErasedAt):
		return PaymentMethodsDBFieldName.GdprErasedAt, true

	case string(PaymentMethodsDBFieldName.MetaCreatedAt):
		return PaymentMethodsDBFieldName.MetaCreatedAt, true

	case string(PaymentMethodsDBFieldName.MetaCreatedBy):
		return PaymentMethodsDBFieldName.MetaCreatedBy, true

	case string(PaymentMethodsDBFieldName.MetaUpdatedAt):
		return PaymentMethodsDBFieldName.MetaUpdatedAt, true

	case string(PaymentMethodsDBFieldName.MetaUpdatedBy):
		return PaymentMethodsDBFieldName.MetaUpdatedBy, true

	case string(PaymentMethodsDBFieldName.MetaDeletedAt):
		return PaymentMethodsDBFieldName.MetaDeletedAt, true

	case string(PaymentMethodsDBFieldName.MetaDeletedBy):
		return PaymentMethodsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type PaymentMethodsFilterResult struct {
	PaymentMethods
	FilterCount int `db:"count"`
}

func ValidatePaymentMethodsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewPaymentMethodsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewPaymentMethodsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewPaymentMethodsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type PaymentMethods struct {
	Id                     uuid.UUID         `db:"id"`
	UserId                 nuuid.NUUID       `db:"user_id"`
	MerchantId             nuuid.NUUID       `db:"merchant_id"`
	MethodType             PaymentMethodType `db:"method_type"`
	Psp                    Psp               `db:"psp"`
	TokenRef               null.String       `db:"token_ref"`
	TokenType              null.String       `db:"token_type"`
	TokenExpiresAt         null.Time         `db:"token_expires_at"`
	CardBrand              null.String       `db:"card_brand"`
	CardLastFour           null.String       `db:"card_last_four"`
	CardExpMonth           null.Int          `db:"card_exp_month"`
	CardExpYear            null.Int          `db:"card_exp_year"`
	CardCountry            null.String       `db:"card_country"`
	CardFundingType        null.String       `db:"card_funding_type"`
	CardBin                null.String       `db:"card_bin"`
	WalletAccountRef       null.String       `db:"wallet_account_ref"`
	VaBankCode             null.String       `db:"va_bank_code"`
	DisplayLabel           null.String       `db:"display_label"`
	IsDefault              bool              `db:"is_default"`
	IsActive               bool              `db:"is_active"`
	VerifiedAt             null.Time         `db:"verified_at"`
	Fingerprint            null.String       `db:"fingerprint"`
	GdprErasureRequestedAt null.Time         `db:"gdpr_erasure_requested_at"`
	GdprErasedAt           null.Time         `db:"gdpr_erased_at"`
	MetaCreatedAt          time.Time         `db:"meta_created_at"`
	MetaCreatedBy          uuid.UUID         `db:"meta_created_by"`
	MetaUpdatedAt          time.Time         `db:"meta_updated_at"`
	MetaUpdatedBy          nuuid.NUUID       `db:"meta_updated_by"`
	MetaDeletedAt          null.Time         `db:"meta_deleted_at"`
	MetaDeletedBy          nuuid.NUUID       `db:"meta_deleted_by"`
}
type PaymentMethodsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentMethods) ToPaymentMethodsPrimaryID() PaymentMethodsPrimaryID {
	return PaymentMethodsPrimaryID{
		Id: d.Id,
	}
}

type PaymentMethodsList []*PaymentMethods
