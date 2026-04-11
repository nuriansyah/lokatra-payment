package dto

import (
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentMethodsDTOFieldNameType string

type paymentMethodsDTOFieldName struct {
	Id                     PaymentMethodsDTOFieldNameType
	UserId                 PaymentMethodsDTOFieldNameType
	MerchantId             PaymentMethodsDTOFieldNameType
	MethodType             PaymentMethodsDTOFieldNameType
	Psp                    PaymentMethodsDTOFieldNameType
	TokenRef               PaymentMethodsDTOFieldNameType
	TokenType              PaymentMethodsDTOFieldNameType
	TokenExpiresAt         PaymentMethodsDTOFieldNameType
	CardBrand              PaymentMethodsDTOFieldNameType
	CardLastFour           PaymentMethodsDTOFieldNameType
	CardExpMonth           PaymentMethodsDTOFieldNameType
	CardExpYear            PaymentMethodsDTOFieldNameType
	CardCountry            PaymentMethodsDTOFieldNameType
	CardFundingType        PaymentMethodsDTOFieldNameType
	CardBin                PaymentMethodsDTOFieldNameType
	WalletAccountRef       PaymentMethodsDTOFieldNameType
	VaBankCode             PaymentMethodsDTOFieldNameType
	DisplayLabel           PaymentMethodsDTOFieldNameType
	IsDefault              PaymentMethodsDTOFieldNameType
	IsActive               PaymentMethodsDTOFieldNameType
	VerifiedAt             PaymentMethodsDTOFieldNameType
	Fingerprint            PaymentMethodsDTOFieldNameType
	GdprErasureRequestedAt PaymentMethodsDTOFieldNameType
	GdprErasedAt           PaymentMethodsDTOFieldNameType
	MetaCreatedAt          PaymentMethodsDTOFieldNameType
	MetaCreatedBy          PaymentMethodsDTOFieldNameType
	MetaUpdatedAt          PaymentMethodsDTOFieldNameType
	MetaUpdatedBy          PaymentMethodsDTOFieldNameType
	MetaDeletedAt          PaymentMethodsDTOFieldNameType
	MetaDeletedBy          PaymentMethodsDTOFieldNameType
}

var PaymentMethodsDTOFieldName = paymentMethodsDTOFieldName{
	Id:                     "id",
	UserId:                 "userId",
	MerchantId:             "merchantId",
	MethodType:             "methodType",
	Psp:                    "psp",
	TokenRef:               "tokenRef",
	TokenType:              "tokenType",
	TokenExpiresAt:         "tokenExpiresAt",
	CardBrand:              "cardBrand",
	CardLastFour:           "cardLastFour",
	CardExpMonth:           "cardExpMonth",
	CardExpYear:            "cardExpYear",
	CardCountry:            "cardCountry",
	CardFundingType:        "cardFundingType",
	CardBin:                "cardBin",
	WalletAccountRef:       "walletAccountRef",
	VaBankCode:             "vaBankCode",
	DisplayLabel:           "displayLabel",
	IsDefault:              "isDefault",
	IsActive:               "isActive",
	VerifiedAt:             "verifiedAt",
	Fingerprint:            "fingerprint",
	GdprErasureRequestedAt: "gdprErasureRequestedAt",
	GdprErasedAt:           "gdprErasedAt",
	MetaCreatedAt:          "metaCreatedAt",
	MetaCreatedBy:          "metaCreatedBy",
	MetaUpdatedAt:          "metaUpdatedAt",
	MetaUpdatedBy:          "metaUpdatedBy",
	MetaDeletedAt:          "metaDeletedAt",
	MetaDeletedBy:          "metaDeletedBy",
}

func NewPaymentMethodsListResponseFromFilterResult(result []model.PaymentMethodsFilterResult, filter model.Filter) PaymentMethodsSelectableListResponse {
	dtoPaymentMethodsListResponse := PaymentMethodsSelectableListResponse{}
	for _, paymentMethods := range result {
		dtoPaymentMethodsResponse := NewPaymentMethodsSelectableResponse(paymentMethods.PaymentMethods, filter)
		dtoPaymentMethodsListResponse = append(dtoPaymentMethodsListResponse, &dtoPaymentMethodsResponse)
	}
	return dtoPaymentMethodsListResponse
}

func transformPaymentMethodsDTOFieldNameFromStr(field string) (dbField model.PaymentMethodsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentMethodsDTOFieldName.Id):
		return model.PaymentMethodsDBFieldName.Id, true

	case string(PaymentMethodsDTOFieldName.UserId):
		return model.PaymentMethodsDBFieldName.UserId, true

	case string(PaymentMethodsDTOFieldName.MerchantId):
		return model.PaymentMethodsDBFieldName.MerchantId, true

	case string(PaymentMethodsDTOFieldName.MethodType):
		return model.PaymentMethodsDBFieldName.MethodType, true

	case string(PaymentMethodsDTOFieldName.Psp):
		return model.PaymentMethodsDBFieldName.Psp, true

	case string(PaymentMethodsDTOFieldName.TokenRef):
		return model.PaymentMethodsDBFieldName.TokenRef, true

	case string(PaymentMethodsDTOFieldName.TokenType):
		return model.PaymentMethodsDBFieldName.TokenType, true

	case string(PaymentMethodsDTOFieldName.TokenExpiresAt):
		return model.PaymentMethodsDBFieldName.TokenExpiresAt, true

	case string(PaymentMethodsDTOFieldName.CardBrand):
		return model.PaymentMethodsDBFieldName.CardBrand, true

	case string(PaymentMethodsDTOFieldName.CardLastFour):
		return model.PaymentMethodsDBFieldName.CardLastFour, true

	case string(PaymentMethodsDTOFieldName.CardExpMonth):
		return model.PaymentMethodsDBFieldName.CardExpMonth, true

	case string(PaymentMethodsDTOFieldName.CardExpYear):
		return model.PaymentMethodsDBFieldName.CardExpYear, true

	case string(PaymentMethodsDTOFieldName.CardCountry):
		return model.PaymentMethodsDBFieldName.CardCountry, true

	case string(PaymentMethodsDTOFieldName.CardFundingType):
		return model.PaymentMethodsDBFieldName.CardFundingType, true

	case string(PaymentMethodsDTOFieldName.CardBin):
		return model.PaymentMethodsDBFieldName.CardBin, true

	case string(PaymentMethodsDTOFieldName.WalletAccountRef):
		return model.PaymentMethodsDBFieldName.WalletAccountRef, true

	case string(PaymentMethodsDTOFieldName.VaBankCode):
		return model.PaymentMethodsDBFieldName.VaBankCode, true

	case string(PaymentMethodsDTOFieldName.DisplayLabel):
		return model.PaymentMethodsDBFieldName.DisplayLabel, true

	case string(PaymentMethodsDTOFieldName.IsDefault):
		return model.PaymentMethodsDBFieldName.IsDefault, true

	case string(PaymentMethodsDTOFieldName.IsActive):
		return model.PaymentMethodsDBFieldName.IsActive, true

	case string(PaymentMethodsDTOFieldName.VerifiedAt):
		return model.PaymentMethodsDBFieldName.VerifiedAt, true

	case string(PaymentMethodsDTOFieldName.Fingerprint):
		return model.PaymentMethodsDBFieldName.Fingerprint, true

	case string(PaymentMethodsDTOFieldName.GdprErasureRequestedAt):
		return model.PaymentMethodsDBFieldName.GdprErasureRequestedAt, true

	case string(PaymentMethodsDTOFieldName.GdprErasedAt):
		return model.PaymentMethodsDBFieldName.GdprErasedAt, true

	case string(PaymentMethodsDTOFieldName.MetaCreatedAt):
		return model.PaymentMethodsDBFieldName.MetaCreatedAt, true

	case string(PaymentMethodsDTOFieldName.MetaCreatedBy):
		return model.PaymentMethodsDBFieldName.MetaCreatedBy, true

	case string(PaymentMethodsDTOFieldName.MetaUpdatedAt):
		return model.PaymentMethodsDBFieldName.MetaUpdatedAt, true

	case string(PaymentMethodsDTOFieldName.MetaUpdatedBy):
		return model.PaymentMethodsDBFieldName.MetaUpdatedBy, true

	case string(PaymentMethodsDTOFieldName.MetaDeletedAt):
		return model.PaymentMethodsDBFieldName.MetaDeletedAt, true

	case string(PaymentMethodsDTOFieldName.MetaDeletedBy):
		return model.PaymentMethodsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformPaymentMethodsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformPaymentMethodsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentMethodsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentMethodsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultPaymentMethodsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(PaymentMethodsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentMethodsSelectableResponse map[string]interface{}
type PaymentMethodsSelectableListResponse []*PaymentMethodsSelectableResponse

func NewPaymentMethodsSelectableResponse(paymentMethods model.PaymentMethods, filter model.Filter) PaymentMethodsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentMethodsDBFieldName.Id),
			string(model.PaymentMethodsDBFieldName.UserId),
			string(model.PaymentMethodsDBFieldName.MerchantId),
			string(model.PaymentMethodsDBFieldName.MethodType),
			string(model.PaymentMethodsDBFieldName.Psp),
			string(model.PaymentMethodsDBFieldName.TokenRef),
			string(model.PaymentMethodsDBFieldName.TokenType),
			string(model.PaymentMethodsDBFieldName.TokenExpiresAt),
			string(model.PaymentMethodsDBFieldName.CardBrand),
			string(model.PaymentMethodsDBFieldName.CardLastFour),
			string(model.PaymentMethodsDBFieldName.CardExpMonth),
			string(model.PaymentMethodsDBFieldName.CardExpYear),
			string(model.PaymentMethodsDBFieldName.CardCountry),
			string(model.PaymentMethodsDBFieldName.CardFundingType),
			string(model.PaymentMethodsDBFieldName.CardBin),
			string(model.PaymentMethodsDBFieldName.WalletAccountRef),
			string(model.PaymentMethodsDBFieldName.VaBankCode),
			string(model.PaymentMethodsDBFieldName.DisplayLabel),
			string(model.PaymentMethodsDBFieldName.IsDefault),
			string(model.PaymentMethodsDBFieldName.IsActive),
			string(model.PaymentMethodsDBFieldName.VerifiedAt),
			string(model.PaymentMethodsDBFieldName.Fingerprint),
			string(model.PaymentMethodsDBFieldName.GdprErasureRequestedAt),
			string(model.PaymentMethodsDBFieldName.GdprErasedAt),
			string(model.PaymentMethodsDBFieldName.MetaCreatedAt),
			string(model.PaymentMethodsDBFieldName.MetaCreatedBy),
			string(model.PaymentMethodsDBFieldName.MetaUpdatedAt),
			string(model.PaymentMethodsDBFieldName.MetaUpdatedBy),
			string(model.PaymentMethodsDBFieldName.MetaDeletedAt),
			string(model.PaymentMethodsDBFieldName.MetaDeletedBy),
		)
	}
	paymentMethodsSelectableResponse := PaymentMethodsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.PaymentMethodsDBFieldName.Id):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.Id)] = paymentMethods.Id

		case string(model.PaymentMethodsDBFieldName.UserId):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.UserId)] = paymentMethods.UserId

		case string(model.PaymentMethodsDBFieldName.MerchantId):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MerchantId)] = paymentMethods.MerchantId

		case string(model.PaymentMethodsDBFieldName.MethodType):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MethodType)] = paymentMethods.MethodType

		case string(model.PaymentMethodsDBFieldName.Psp):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.Psp)] = paymentMethods.Psp

		case string(model.PaymentMethodsDBFieldName.TokenRef):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.TokenRef)] = paymentMethods.TokenRef

		case string(model.PaymentMethodsDBFieldName.TokenType):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.TokenType)] = paymentMethods.TokenType

		case string(model.PaymentMethodsDBFieldName.TokenExpiresAt):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.TokenExpiresAt)] = paymentMethods.TokenExpiresAt

		case string(model.PaymentMethodsDBFieldName.CardBrand):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.CardBrand)] = paymentMethods.CardBrand

		case string(model.PaymentMethodsDBFieldName.CardLastFour):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.CardLastFour)] = paymentMethods.CardLastFour

		case string(model.PaymentMethodsDBFieldName.CardExpMonth):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.CardExpMonth)] = paymentMethods.CardExpMonth

		case string(model.PaymentMethodsDBFieldName.CardExpYear):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.CardExpYear)] = paymentMethods.CardExpYear

		case string(model.PaymentMethodsDBFieldName.CardCountry):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.CardCountry)] = paymentMethods.CardCountry

		case string(model.PaymentMethodsDBFieldName.CardFundingType):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.CardFundingType)] = paymentMethods.CardFundingType

		case string(model.PaymentMethodsDBFieldName.CardBin):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.CardBin)] = paymentMethods.CardBin

		case string(model.PaymentMethodsDBFieldName.WalletAccountRef):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.WalletAccountRef)] = paymentMethods.WalletAccountRef

		case string(model.PaymentMethodsDBFieldName.VaBankCode):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.VaBankCode)] = paymentMethods.VaBankCode

		case string(model.PaymentMethodsDBFieldName.DisplayLabel):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.DisplayLabel)] = paymentMethods.DisplayLabel

		case string(model.PaymentMethodsDBFieldName.IsDefault):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.IsDefault)] = paymentMethods.IsDefault

		case string(model.PaymentMethodsDBFieldName.IsActive):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.IsActive)] = paymentMethods.IsActive

		case string(model.PaymentMethodsDBFieldName.VerifiedAt):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.VerifiedAt)] = paymentMethods.VerifiedAt

		case string(model.PaymentMethodsDBFieldName.Fingerprint):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.Fingerprint)] = paymentMethods.Fingerprint

		case string(model.PaymentMethodsDBFieldName.GdprErasureRequestedAt):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.GdprErasureRequestedAt)] = paymentMethods.GdprErasureRequestedAt

		case string(model.PaymentMethodsDBFieldName.GdprErasedAt):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.GdprErasedAt)] = paymentMethods.GdprErasedAt

		case string(model.PaymentMethodsDBFieldName.MetaCreatedAt):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MetaCreatedAt)] = paymentMethods.MetaCreatedAt

		case string(model.PaymentMethodsDBFieldName.MetaCreatedBy):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MetaCreatedBy)] = paymentMethods.MetaCreatedBy

		case string(model.PaymentMethodsDBFieldName.MetaUpdatedAt):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MetaUpdatedAt)] = paymentMethods.MetaUpdatedAt

		case string(model.PaymentMethodsDBFieldName.MetaUpdatedBy):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MetaUpdatedBy)] = paymentMethods.MetaUpdatedBy

		case string(model.PaymentMethodsDBFieldName.MetaDeletedAt):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MetaDeletedAt)] = paymentMethods.MetaDeletedAt

		case string(model.PaymentMethodsDBFieldName.MetaDeletedBy):
			paymentMethodsSelectableResponse[string(PaymentMethodsDTOFieldName.MetaDeletedBy)] = paymentMethods.MetaDeletedBy

		}
	}
	return paymentMethodsSelectableResponse
}

type PaymentMethodsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     PaymentMethodsSelectableListResponse `json:"data"`
}

func NewPaymentMethodsFilterResponse(result []model.PaymentMethodsFilterResult, filter model.Filter) (resp PaymentMethodsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewPaymentMethodsListResponseFromFilterResult(result, filter)
	return resp
}

type PaymentMethodsCreateRequest struct {
	UserId                 uuid.UUID               `json:"userId"`
	MerchantId             uuid.UUID               `json:"merchantId"`
	MethodType             model.PaymentMethodType `json:"methodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Psp                    model.Psp               `json:"psp" example:"MIDTRANS" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	TokenRef               string                  `json:"tokenRef"`
	TokenType              string                  `json:"tokenType"`
	TokenExpiresAt         time.Time               `json:"tokenExpiresAt"`
	CardBrand              string                  `json:"cardBrand"`
	CardLastFour           string                  `json:"cardLastFour"`
	CardExpMonth           int16                   `json:"cardExpMonth"`
	CardExpYear            int16                   `json:"cardExpYear"`
	CardCountry            string                  `json:"cardCountry"`
	CardFundingType        string                  `json:"cardFundingType"`
	CardBin                string                  `json:"cardBin"`
	WalletAccountRef       string                  `json:"walletAccountRef"`
	VaBankCode             string                  `json:"vaBankCode"`
	DisplayLabel           string                  `json:"displayLabel"`
	IsDefault              bool                    `json:"isDefault"`
	IsActive               bool                    `json:"isActive"`
	VerifiedAt             time.Time               `json:"verifiedAt"`
	Fingerprint            string                  `json:"fingerprint"`
	GdprErasureRequestedAt time.Time               `json:"gdprErasureRequestedAt"`
	GdprErasedAt           time.Time               `json:"gdprErasedAt"`
	MetaCreatedAt          time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy          uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt          time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy          uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt          time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy          uuid.UUID               `json:"metaDeletedBy"`
}

func (d *PaymentMethodsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentMethodsCreateRequest) ToModel() model.PaymentMethods {
	id, _ := uuid.NewV7()
	return model.PaymentMethods{
		Id:                     id,
		UserId:                 nuuid.From(d.UserId),
		MerchantId:             nuuid.From(d.MerchantId),
		MethodType:             d.MethodType,
		Psp:                    d.Psp,
		TokenRef:               null.StringFrom(d.TokenRef),
		TokenType:              null.StringFrom(d.TokenType),
		TokenExpiresAt:         null.TimeFrom(d.TokenExpiresAt),
		CardBrand:              null.StringFrom(d.CardBrand),
		CardLastFour:           null.StringFrom(d.CardLastFour),
		CardExpMonth:           null.IntFrom(int64(d.CardExpMonth)),
		CardExpYear:            null.IntFrom(int64(d.CardExpYear)),
		CardCountry:            null.StringFrom(d.CardCountry),
		CardFundingType:        null.StringFrom(d.CardFundingType),
		CardBin:                null.StringFrom(d.CardBin),
		WalletAccountRef:       null.StringFrom(d.WalletAccountRef),
		VaBankCode:             null.StringFrom(d.VaBankCode),
		DisplayLabel:           null.StringFrom(d.DisplayLabel),
		IsDefault:              d.IsDefault,
		IsActive:               d.IsActive,
		VerifiedAt:             null.TimeFrom(d.VerifiedAt),
		Fingerprint:            null.StringFrom(d.Fingerprint),
		GdprErasureRequestedAt: null.TimeFrom(d.GdprErasureRequestedAt),
		GdprErasedAt:           null.TimeFrom(d.GdprErasedAt),
		MetaCreatedAt:          d.MetaCreatedAt,
		MetaCreatedBy:          d.MetaCreatedBy,
		MetaUpdatedAt:          d.MetaUpdatedAt,
		MetaUpdatedBy:          nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:          null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:          nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentMethodsListCreateRequest []*PaymentMethodsCreateRequest

func (d PaymentMethodsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentMethods := range d {
		err = validator.Struct(paymentMethods)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentMethodsListCreateRequest) ToModelList() []model.PaymentMethods {
	out := make([]model.PaymentMethods, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentMethodsUpdateRequest struct {
	UserId                 uuid.UUID               `json:"userId"`
	MerchantId             uuid.UUID               `json:"merchantId"`
	MethodType             model.PaymentMethodType `json:"methodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Psp                    model.Psp               `json:"psp" example:"MIDTRANS" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	TokenRef               string                  `json:"tokenRef"`
	TokenType              string                  `json:"tokenType"`
	TokenExpiresAt         time.Time               `json:"tokenExpiresAt"`
	CardBrand              string                  `json:"cardBrand"`
	CardLastFour           string                  `json:"cardLastFour"`
	CardExpMonth           int16                   `json:"cardExpMonth"`
	CardExpYear            int16                   `json:"cardExpYear"`
	CardCountry            string                  `json:"cardCountry"`
	CardFundingType        string                  `json:"cardFundingType"`
	CardBin                string                  `json:"cardBin"`
	WalletAccountRef       string                  `json:"walletAccountRef"`
	VaBankCode             string                  `json:"vaBankCode"`
	DisplayLabel           string                  `json:"displayLabel"`
	IsDefault              bool                    `json:"isDefault"`
	IsActive               bool                    `json:"isActive"`
	VerifiedAt             time.Time               `json:"verifiedAt"`
	Fingerprint            string                  `json:"fingerprint"`
	GdprErasureRequestedAt time.Time               `json:"gdprErasureRequestedAt"`
	GdprErasedAt           time.Time               `json:"gdprErasedAt"`
	MetaCreatedAt          time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy          uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt          time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy          uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt          time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy          uuid.UUID               `json:"metaDeletedBy"`
}

func (d *PaymentMethodsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentMethodsUpdateRequest) ToModel() model.PaymentMethods {
	return model.PaymentMethods{
		UserId:                 nuuid.From(d.UserId),
		MerchantId:             nuuid.From(d.MerchantId),
		MethodType:             d.MethodType,
		Psp:                    d.Psp,
		TokenRef:               null.StringFrom(d.TokenRef),
		TokenType:              null.StringFrom(d.TokenType),
		TokenExpiresAt:         null.TimeFrom(d.TokenExpiresAt),
		CardBrand:              null.StringFrom(d.CardBrand),
		CardLastFour:           null.StringFrom(d.CardLastFour),
		CardExpMonth:           null.IntFrom(int64(d.CardExpMonth)),
		CardExpYear:            null.IntFrom(int64(d.CardExpYear)),
		CardCountry:            null.StringFrom(d.CardCountry),
		CardFundingType:        null.StringFrom(d.CardFundingType),
		CardBin:                null.StringFrom(d.CardBin),
		WalletAccountRef:       null.StringFrom(d.WalletAccountRef),
		VaBankCode:             null.StringFrom(d.VaBankCode),
		DisplayLabel:           null.StringFrom(d.DisplayLabel),
		IsDefault:              d.IsDefault,
		IsActive:               d.IsActive,
		VerifiedAt:             null.TimeFrom(d.VerifiedAt),
		Fingerprint:            null.StringFrom(d.Fingerprint),
		GdprErasureRequestedAt: null.TimeFrom(d.GdprErasureRequestedAt),
		GdprErasedAt:           null.TimeFrom(d.GdprErasedAt),
		MetaCreatedAt:          d.MetaCreatedAt,
		MetaCreatedBy:          d.MetaCreatedBy,
		MetaUpdatedAt:          d.MetaUpdatedAt,
		MetaUpdatedBy:          nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:          null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:          nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentMethodsBulkUpdateRequest struct {
	Id                     uuid.UUID               `json:"id"`
	UserId                 uuid.UUID               `json:"userId"`
	MerchantId             uuid.UUID               `json:"merchantId"`
	MethodType             model.PaymentMethodType `json:"methodType" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Psp                    model.Psp               `json:"psp" example:"MIDTRANS" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	TokenRef               string                  `json:"tokenRef"`
	TokenType              string                  `json:"tokenType"`
	TokenExpiresAt         time.Time               `json:"tokenExpiresAt"`
	CardBrand              string                  `json:"cardBrand"`
	CardLastFour           string                  `json:"cardLastFour"`
	CardExpMonth           int16                   `json:"cardExpMonth"`
	CardExpYear            int16                   `json:"cardExpYear"`
	CardCountry            string                  `json:"cardCountry"`
	CardFundingType        string                  `json:"cardFundingType"`
	CardBin                string                  `json:"cardBin"`
	WalletAccountRef       string                  `json:"walletAccountRef"`
	VaBankCode             string                  `json:"vaBankCode"`
	DisplayLabel           string                  `json:"displayLabel"`
	IsDefault              bool                    `json:"isDefault"`
	IsActive               bool                    `json:"isActive"`
	VerifiedAt             time.Time               `json:"verifiedAt"`
	Fingerprint            string                  `json:"fingerprint"`
	GdprErasureRequestedAt time.Time               `json:"gdprErasureRequestedAt"`
	GdprErasedAt           time.Time               `json:"gdprErasedAt"`
	MetaCreatedAt          time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy          uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt          time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy          uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt          time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy          uuid.UUID               `json:"metaDeletedBy"`
}

func (d PaymentMethodsBulkUpdateRequest) PrimaryID() PaymentMethodsPrimaryID {
	return PaymentMethodsPrimaryID{
		Id: d.Id,
	}
}

type PaymentMethodsListBulkUpdateRequest []*PaymentMethodsBulkUpdateRequest

func (d PaymentMethodsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentMethods := range d {
		err = validator.Struct(paymentMethods)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentMethodsBulkUpdateRequest) ToModel() model.PaymentMethods {
	return model.PaymentMethods{
		Id:                     d.Id,
		UserId:                 nuuid.From(d.UserId),
		MerchantId:             nuuid.From(d.MerchantId),
		MethodType:             d.MethodType,
		Psp:                    d.Psp,
		TokenRef:               null.StringFrom(d.TokenRef),
		TokenType:              null.StringFrom(d.TokenType),
		TokenExpiresAt:         null.TimeFrom(d.TokenExpiresAt),
		CardBrand:              null.StringFrom(d.CardBrand),
		CardLastFour:           null.StringFrom(d.CardLastFour),
		CardExpMonth:           null.IntFrom(int64(d.CardExpMonth)),
		CardExpYear:            null.IntFrom(int64(d.CardExpYear)),
		CardCountry:            null.StringFrom(d.CardCountry),
		CardFundingType:        null.StringFrom(d.CardFundingType),
		CardBin:                null.StringFrom(d.CardBin),
		WalletAccountRef:       null.StringFrom(d.WalletAccountRef),
		VaBankCode:             null.StringFrom(d.VaBankCode),
		DisplayLabel:           null.StringFrom(d.DisplayLabel),
		IsDefault:              d.IsDefault,
		IsActive:               d.IsActive,
		VerifiedAt:             null.TimeFrom(d.VerifiedAt),
		Fingerprint:            null.StringFrom(d.Fingerprint),
		GdprErasureRequestedAt: null.TimeFrom(d.GdprErasureRequestedAt),
		GdprErasedAt:           null.TimeFrom(d.GdprErasedAt),
		MetaCreatedAt:          d.MetaCreatedAt,
		MetaCreatedBy:          d.MetaCreatedBy,
		MetaUpdatedAt:          d.MetaUpdatedAt,
		MetaUpdatedBy:          nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:          null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:          nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentMethodsResponse struct {
	Id                     uuid.UUID               `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserId                 uuid.UUID               `json:"userId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantId             uuid.UUID               `json:"merchantId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MethodType             model.PaymentMethodType `json:"methodType" validate:"required,oneof=CARD VIRTUAL_ACCOUNT QRIS EWALLET DIRECT_DEBIT BANK_TRANSFER PAYLATER VOUCHER POINTS CASH_ON_DELIVERY" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	Psp                    model.Psp               `json:"psp" validate:"required,oneof=MIDTRANS XENDIT STRIPE DOKU DANA OVO GOPAY SHOPEE_PAY LINK_AJA FLIP INTERNAL" enums:"MIDTRANS,XENDIT,STRIPE,DOKU,DANA,OVO,GOPAY,SHOPEE_PAY,LINK_AJA,FLIP,INTERNAL"`
	TokenRef               string                  `json:"tokenRef"`
	TokenType              string                  `json:"tokenType"`
	TokenExpiresAt         time.Time               `json:"tokenExpiresAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CardBrand              string                  `json:"cardBrand"`
	CardLastFour           string                  `json:"cardLastFour"`
	CardExpMonth           int16                   `json:"cardExpMonth"`
	CardExpYear            int16                   `json:"cardExpYear"`
	CardCountry            string                  `json:"cardCountry"`
	CardFundingType        string                  `json:"cardFundingType"`
	CardBin                string                  `json:"cardBin"`
	WalletAccountRef       string                  `json:"walletAccountRef"`
	VaBankCode             string                  `json:"vaBankCode"`
	DisplayLabel           string                  `json:"displayLabel"`
	IsDefault              bool                    `json:"isDefault" example:"true"`
	IsActive               bool                    `json:"isActive" example:"true"`
	VerifiedAt             time.Time               `json:"verifiedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Fingerprint            string                  `json:"fingerprint"`
	GdprErasureRequestedAt time.Time               `json:"gdprErasureRequestedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	GdprErasedAt           time.Time               `json:"gdprErasedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedAt          time.Time               `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy          uuid.UUID               `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt          time.Time               `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy          uuid.UUID               `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt          time.Time               `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy          uuid.UUID               `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewPaymentMethodsResponse(paymentMethods model.PaymentMethods) PaymentMethodsResponse {
	return PaymentMethodsResponse{
		Id:                     paymentMethods.Id,
		UserId:                 paymentMethods.UserId.UUID,
		MerchantId:             paymentMethods.MerchantId.UUID,
		MethodType:             model.PaymentMethodType(paymentMethods.MethodType),
		Psp:                    model.Psp(paymentMethods.Psp),
		TokenRef:               paymentMethods.TokenRef.String,
		TokenType:              paymentMethods.TokenType.String,
		TokenExpiresAt:         paymentMethods.TokenExpiresAt.Time,
		CardBrand:              paymentMethods.CardBrand.String,
		CardLastFour:           paymentMethods.CardLastFour.String,
		CardExpMonth:           int16(paymentMethods.CardExpMonth.ValueOrZero()),
		CardExpYear:            int16(paymentMethods.CardExpYear.ValueOrZero()),
		CardCountry:            paymentMethods.CardCountry.String,
		CardFundingType:        paymentMethods.CardFundingType.String,
		CardBin:                paymentMethods.CardBin.String,
		WalletAccountRef:       paymentMethods.WalletAccountRef.String,
		VaBankCode:             paymentMethods.VaBankCode.String,
		DisplayLabel:           paymentMethods.DisplayLabel.String,
		IsDefault:              paymentMethods.IsDefault,
		IsActive:               paymentMethods.IsActive,
		VerifiedAt:             paymentMethods.VerifiedAt.Time,
		Fingerprint:            paymentMethods.Fingerprint.String,
		GdprErasureRequestedAt: paymentMethods.GdprErasureRequestedAt.Time,
		GdprErasedAt:           paymentMethods.GdprErasedAt.Time,
		MetaCreatedAt:          paymentMethods.MetaCreatedAt,
		MetaCreatedBy:          paymentMethods.MetaCreatedBy,
		MetaUpdatedAt:          paymentMethods.MetaUpdatedAt,
		MetaUpdatedBy:          paymentMethods.MetaUpdatedBy.UUID,
		MetaDeletedAt:          paymentMethods.MetaDeletedAt.Time,
		MetaDeletedBy:          paymentMethods.MetaDeletedBy.UUID,
	}
}

type PaymentMethodsListResponse []*PaymentMethodsResponse

func NewPaymentMethodsListResponse(paymentMethodsList model.PaymentMethodsList) PaymentMethodsListResponse {
	dtoPaymentMethodsListResponse := PaymentMethodsListResponse{}
	for _, paymentMethods := range paymentMethodsList {
		dtoPaymentMethodsResponse := NewPaymentMethodsResponse(*paymentMethods)
		dtoPaymentMethodsListResponse = append(dtoPaymentMethodsListResponse, &dtoPaymentMethodsResponse)
	}
	return dtoPaymentMethodsListResponse
}

type PaymentMethodsPrimaryIDList []PaymentMethodsPrimaryID

func (d PaymentMethodsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentMethods := range d {
		err = validator.Struct(paymentMethods)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentMethodsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentMethodsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentMethodsPrimaryID) ToModel() model.PaymentMethodsPrimaryID {
	return model.PaymentMethodsPrimaryID{
		Id: d.Id,
	}
}
