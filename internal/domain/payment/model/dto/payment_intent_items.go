package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentIntentItemsDTOFieldNameType string

type paymentIntentItemsDTOFieldName struct {
	Id               PaymentIntentItemsDTOFieldNameType
	IntentId         PaymentIntentItemsDTOFieldNameType
	ProductId        PaymentIntentItemsDTOFieldNameType
	ProductType      PaymentIntentItemsDTOFieldNameType
	ProductName      PaymentIntentItemsDTOFieldNameType
	Quantity         PaymentIntentItemsDTOFieldNameType
	UnitPrice        PaymentIntentItemsDTOFieldNameType
	DiscountAmount   PaymentIntentItemsDTOFieldNameType
	TotalPrice       PaymentIntentItemsDTOFieldNameType
	SellerMerchantId PaymentIntentItemsDTOFieldNameType
	Metadata         PaymentIntentItemsDTOFieldNameType
	MetaCreatedAt    PaymentIntentItemsDTOFieldNameType
	MetaCreatedBy    PaymentIntentItemsDTOFieldNameType
	MetaUpdatedAt    PaymentIntentItemsDTOFieldNameType
	MetaUpdatedBy    PaymentIntentItemsDTOFieldNameType
	MetaDeletedAt    PaymentIntentItemsDTOFieldNameType
	MetaDeletedBy    PaymentIntentItemsDTOFieldNameType
}

var PaymentIntentItemsDTOFieldName = paymentIntentItemsDTOFieldName{
	Id:               "id",
	IntentId:         "intentId",
	ProductId:        "productId",
	ProductType:      "productType",
	ProductName:      "productName",
	Quantity:         "quantity",
	UnitPrice:        "unitPrice",
	DiscountAmount:   "discountAmount",
	TotalPrice:       "totalPrice",
	SellerMerchantId: "sellerMerchantId",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func NewPaymentIntentItemsListResponseFromFilterResult(result []model.PaymentIntentItemsFilterResult, filter model.Filter) PaymentIntentItemsSelectableListResponse {
	dtoPaymentIntentItemsListResponse := PaymentIntentItemsSelectableListResponse{}
	for _, paymentIntentItems := range result {
		dtoPaymentIntentItemsResponse := NewPaymentIntentItemsSelectableResponse(paymentIntentItems.PaymentIntentItems, filter)
		dtoPaymentIntentItemsListResponse = append(dtoPaymentIntentItemsListResponse, &dtoPaymentIntentItemsResponse)
	}
	return dtoPaymentIntentItemsListResponse
}

func transformPaymentIntentItemsDTOFieldNameFromStr(field string) (dbField model.PaymentIntentItemsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentIntentItemsDTOFieldName.Id):
		return model.PaymentIntentItemsDBFieldName.Id, true

	case string(PaymentIntentItemsDTOFieldName.IntentId):
		return model.PaymentIntentItemsDBFieldName.IntentId, true

	case string(PaymentIntentItemsDTOFieldName.ProductId):
		return model.PaymentIntentItemsDBFieldName.ProductId, true

	case string(PaymentIntentItemsDTOFieldName.ProductType):
		return model.PaymentIntentItemsDBFieldName.ProductType, true

	case string(PaymentIntentItemsDTOFieldName.ProductName):
		return model.PaymentIntentItemsDBFieldName.ProductName, true

	case string(PaymentIntentItemsDTOFieldName.Quantity):
		return model.PaymentIntentItemsDBFieldName.Quantity, true

	case string(PaymentIntentItemsDTOFieldName.UnitPrice):
		return model.PaymentIntentItemsDBFieldName.UnitPrice, true

	case string(PaymentIntentItemsDTOFieldName.DiscountAmount):
		return model.PaymentIntentItemsDBFieldName.DiscountAmount, true

	case string(PaymentIntentItemsDTOFieldName.TotalPrice):
		return model.PaymentIntentItemsDBFieldName.TotalPrice, true

	case string(PaymentIntentItemsDTOFieldName.SellerMerchantId):
		return model.PaymentIntentItemsDBFieldName.SellerMerchantId, true

	case string(PaymentIntentItemsDTOFieldName.Metadata):
		return model.PaymentIntentItemsDBFieldName.Metadata, true

	case string(PaymentIntentItemsDTOFieldName.MetaCreatedAt):
		return model.PaymentIntentItemsDBFieldName.MetaCreatedAt, true

	case string(PaymentIntentItemsDTOFieldName.MetaCreatedBy):
		return model.PaymentIntentItemsDBFieldName.MetaCreatedBy, true

	case string(PaymentIntentItemsDTOFieldName.MetaUpdatedAt):
		return model.PaymentIntentItemsDBFieldName.MetaUpdatedAt, true

	case string(PaymentIntentItemsDTOFieldName.MetaUpdatedBy):
		return model.PaymentIntentItemsDBFieldName.MetaUpdatedBy, true

	case string(PaymentIntentItemsDTOFieldName.MetaDeletedAt):
		return model.PaymentIntentItemsDBFieldName.MetaDeletedAt, true

	case string(PaymentIntentItemsDTOFieldName.MetaDeletedBy):
		return model.PaymentIntentItemsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformPaymentIntentItemsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformPaymentIntentItemsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentIntentItemsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentIntentItemsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultPaymentIntentItemsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(PaymentIntentItemsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentIntentItemsSelectableResponse map[string]interface{}
type PaymentIntentItemsSelectableListResponse []*PaymentIntentItemsSelectableResponse

func NewPaymentIntentItemsSelectableResponse(paymentIntentItems model.PaymentIntentItems, filter model.Filter) PaymentIntentItemsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentIntentItemsDBFieldName.Id),
			string(model.PaymentIntentItemsDBFieldName.IntentId),
			string(model.PaymentIntentItemsDBFieldName.ProductId),
			string(model.PaymentIntentItemsDBFieldName.ProductType),
			string(model.PaymentIntentItemsDBFieldName.ProductName),
			string(model.PaymentIntentItemsDBFieldName.Quantity),
			string(model.PaymentIntentItemsDBFieldName.UnitPrice),
			string(model.PaymentIntentItemsDBFieldName.DiscountAmount),
			string(model.PaymentIntentItemsDBFieldName.TotalPrice),
			string(model.PaymentIntentItemsDBFieldName.SellerMerchantId),
			string(model.PaymentIntentItemsDBFieldName.Metadata),
			string(model.PaymentIntentItemsDBFieldName.MetaCreatedAt),
			string(model.PaymentIntentItemsDBFieldName.MetaCreatedBy),
			string(model.PaymentIntentItemsDBFieldName.MetaUpdatedAt),
			string(model.PaymentIntentItemsDBFieldName.MetaUpdatedBy),
			string(model.PaymentIntentItemsDBFieldName.MetaDeletedAt),
			string(model.PaymentIntentItemsDBFieldName.MetaDeletedBy),
		)
	}
	paymentIntentItemsSelectableResponse := PaymentIntentItemsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.PaymentIntentItemsDBFieldName.Id):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.Id)] = paymentIntentItems.Id

		case string(model.PaymentIntentItemsDBFieldName.IntentId):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.IntentId)] = paymentIntentItems.IntentId

		case string(model.PaymentIntentItemsDBFieldName.ProductId):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.ProductId)] = paymentIntentItems.ProductId

		case string(model.PaymentIntentItemsDBFieldName.ProductType):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.ProductType)] = paymentIntentItems.ProductType

		case string(model.PaymentIntentItemsDBFieldName.ProductName):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.ProductName)] = paymentIntentItems.ProductName

		case string(model.PaymentIntentItemsDBFieldName.Quantity):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.Quantity)] = paymentIntentItems.Quantity

		case string(model.PaymentIntentItemsDBFieldName.UnitPrice):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.UnitPrice)] = paymentIntentItems.UnitPrice

		case string(model.PaymentIntentItemsDBFieldName.DiscountAmount):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.DiscountAmount)] = paymentIntentItems.DiscountAmount

		case string(model.PaymentIntentItemsDBFieldName.TotalPrice):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.TotalPrice)] = paymentIntentItems.TotalPrice

		case string(model.PaymentIntentItemsDBFieldName.SellerMerchantId):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.SellerMerchantId)] = paymentIntentItems.SellerMerchantId

		case string(model.PaymentIntentItemsDBFieldName.Metadata):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.Metadata)] = paymentIntentItems.Metadata

		case string(model.PaymentIntentItemsDBFieldName.MetaCreatedAt):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.MetaCreatedAt)] = paymentIntentItems.MetaCreatedAt

		case string(model.PaymentIntentItemsDBFieldName.MetaCreatedBy):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.MetaCreatedBy)] = paymentIntentItems.MetaCreatedBy

		case string(model.PaymentIntentItemsDBFieldName.MetaUpdatedAt):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.MetaUpdatedAt)] = paymentIntentItems.MetaUpdatedAt

		case string(model.PaymentIntentItemsDBFieldName.MetaUpdatedBy):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.MetaUpdatedBy)] = paymentIntentItems.MetaUpdatedBy

		case string(model.PaymentIntentItemsDBFieldName.MetaDeletedAt):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.MetaDeletedAt)] = paymentIntentItems.MetaDeletedAt

		case string(model.PaymentIntentItemsDBFieldName.MetaDeletedBy):
			paymentIntentItemsSelectableResponse[string(PaymentIntentItemsDTOFieldName.MetaDeletedBy)] = paymentIntentItems.MetaDeletedBy

		}
	}
	return paymentIntentItemsSelectableResponse
}

type PaymentIntentItemsFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     PaymentIntentItemsSelectableListResponse `json:"data"`
}

func NewPaymentIntentItemsFilterResponse(result []model.PaymentIntentItemsFilterResult, filter model.Filter) (resp PaymentIntentItemsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewPaymentIntentItemsListResponseFromFilterResult(result, filter)
	return resp
}

type PaymentIntentItemsCreateRequest struct {
	IntentId         uuid.UUID       `json:"intentId"`
	ProductId        uuid.UUID       `json:"productId"`
	ProductType      string          `json:"productType"`
	ProductName      string          `json:"productName"`
	Quantity         int             `json:"quantity"`
	UnitPrice        decimal.Decimal `json:"unitPrice"`
	DiscountAmount   decimal.Decimal `json:"discountAmount"`
	TotalPrice       decimal.Decimal `json:"totalPrice"`
	SellerMerchantId uuid.UUID       `json:"sellerMerchantId"`
	Metadata         json.RawMessage `json:"metadata"`
	MetaCreatedAt    time.Time       `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID       `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time       `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID       `json:"metaUpdatedBy"`
	MetaDeletedAt    time.Time       `json:"metaDeletedAt"`
	MetaDeletedBy    uuid.UUID       `json:"metaDeletedBy"`
}

func (d *PaymentIntentItemsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentIntentItemsCreateRequest) ToModel() model.PaymentIntentItems {
	id, _ := uuid.NewV7()
	return model.PaymentIntentItems{
		Id:               id,
		IntentId:         d.IntentId,
		ProductId:        d.ProductId,
		ProductType:      d.ProductType,
		ProductName:      d.ProductName,
		Quantity:         d.Quantity,
		UnitPrice:        d.UnitPrice,
		DiscountAmount:   d.DiscountAmount,
		TotalPrice:       d.TotalPrice,
		SellerMerchantId: nuuid.From(d.SellerMerchantId),
		Metadata:         d.Metadata,
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:    null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:    nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentIntentItemsListCreateRequest []*PaymentIntentItemsCreateRequest

func (d PaymentIntentItemsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntentItems := range d {
		err = validator.Struct(paymentIntentItems)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentIntentItemsListCreateRequest) ToModelList() []model.PaymentIntentItems {
	out := make([]model.PaymentIntentItems, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentIntentItemsUpdateRequest struct {
	IntentId         uuid.UUID       `json:"intentId"`
	ProductId        uuid.UUID       `json:"productId"`
	ProductType      string          `json:"productType"`
	ProductName      string          `json:"productName"`
	Quantity         int             `json:"quantity"`
	UnitPrice        decimal.Decimal `json:"unitPrice"`
	DiscountAmount   decimal.Decimal `json:"discountAmount"`
	TotalPrice       decimal.Decimal `json:"totalPrice"`
	SellerMerchantId uuid.UUID       `json:"sellerMerchantId"`
	Metadata         json.RawMessage `json:"metadata"`
	MetaCreatedAt    time.Time       `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID       `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time       `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID       `json:"metaUpdatedBy"`
	MetaDeletedAt    time.Time       `json:"metaDeletedAt"`
	MetaDeletedBy    uuid.UUID       `json:"metaDeletedBy"`
}

func (d *PaymentIntentItemsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentIntentItemsUpdateRequest) ToModel() model.PaymentIntentItems {
	return model.PaymentIntentItems{
		IntentId:         d.IntentId,
		ProductId:        d.ProductId,
		ProductType:      d.ProductType,
		ProductName:      d.ProductName,
		Quantity:         d.Quantity,
		UnitPrice:        d.UnitPrice,
		DiscountAmount:   d.DiscountAmount,
		TotalPrice:       d.TotalPrice,
		SellerMerchantId: nuuid.From(d.SellerMerchantId),
		Metadata:         d.Metadata,
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:    null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:    nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentIntentItemsBulkUpdateRequest struct {
	Id               uuid.UUID       `json:"id"`
	IntentId         uuid.UUID       `json:"intentId"`
	ProductId        uuid.UUID       `json:"productId"`
	ProductType      string          `json:"productType"`
	ProductName      string          `json:"productName"`
	Quantity         int             `json:"quantity"`
	UnitPrice        decimal.Decimal `json:"unitPrice"`
	DiscountAmount   decimal.Decimal `json:"discountAmount"`
	TotalPrice       decimal.Decimal `json:"totalPrice"`
	SellerMerchantId uuid.UUID       `json:"sellerMerchantId"`
	Metadata         json.RawMessage `json:"metadata"`
	MetaCreatedAt    time.Time       `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID       `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time       `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID       `json:"metaUpdatedBy"`
	MetaDeletedAt    time.Time       `json:"metaDeletedAt"`
	MetaDeletedBy    uuid.UUID       `json:"metaDeletedBy"`
}

func (d PaymentIntentItemsBulkUpdateRequest) PrimaryID() PaymentIntentItemsPrimaryID {
	return PaymentIntentItemsPrimaryID{
		Id: d.Id,
	}
}

type PaymentIntentItemsListBulkUpdateRequest []*PaymentIntentItemsBulkUpdateRequest

func (d PaymentIntentItemsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntentItems := range d {
		err = validator.Struct(paymentIntentItems)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentIntentItemsBulkUpdateRequest) ToModel() model.PaymentIntentItems {
	return model.PaymentIntentItems{
		Id:               d.Id,
		IntentId:         d.IntentId,
		ProductId:        d.ProductId,
		ProductType:      d.ProductType,
		ProductName:      d.ProductName,
		Quantity:         d.Quantity,
		UnitPrice:        d.UnitPrice,
		DiscountAmount:   d.DiscountAmount,
		TotalPrice:       d.TotalPrice,
		SellerMerchantId: nuuid.From(d.SellerMerchantId),
		Metadata:         d.Metadata,
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:    null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:    nuuid.From(d.MetaDeletedBy),
	}
}

type PaymentIntentItemsResponse struct {
	Id               uuid.UUID       `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IntentId         uuid.UUID       `json:"intentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProductId        uuid.UUID       `json:"productId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProductType      string          `json:"productType" validate:"required"`
	ProductName      string          `json:"productName" validate:"required"`
	Quantity         int             `json:"quantity" example:"1"`
	UnitPrice        decimal.Decimal `json:"unitPrice" validate:"required" format:"decimal" example:"100.50"`
	DiscountAmount   decimal.Decimal `json:"discountAmount" format:"decimal" example:"100.50"`
	TotalPrice       decimal.Decimal `json:"totalPrice" validate:"required" format:"decimal" example:"100.50"`
	SellerMerchantId uuid.UUID       `json:"sellerMerchantId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Metadata         json.RawMessage `json:"metadata" swaggertype:"object"`
	MetaCreatedAt    time.Time       `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy    uuid.UUID       `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt    time.Time       `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy    uuid.UUID       `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt    time.Time       `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy    uuid.UUID       `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewPaymentIntentItemsResponse(paymentIntentItems model.PaymentIntentItems) PaymentIntentItemsResponse {
	return PaymentIntentItemsResponse{
		Id:               paymentIntentItems.Id,
		IntentId:         paymentIntentItems.IntentId,
		ProductId:        paymentIntentItems.ProductId,
		ProductType:      paymentIntentItems.ProductType,
		ProductName:      paymentIntentItems.ProductName,
		Quantity:         paymentIntentItems.Quantity,
		UnitPrice:        paymentIntentItems.UnitPrice,
		DiscountAmount:   paymentIntentItems.DiscountAmount,
		TotalPrice:       paymentIntentItems.TotalPrice,
		SellerMerchantId: paymentIntentItems.SellerMerchantId.UUID,
		Metadata:         paymentIntentItems.Metadata,
		MetaCreatedAt:    paymentIntentItems.MetaCreatedAt,
		MetaCreatedBy:    paymentIntentItems.MetaCreatedBy,
		MetaUpdatedAt:    paymentIntentItems.MetaUpdatedAt,
		MetaUpdatedBy:    paymentIntentItems.MetaUpdatedBy.UUID,
		MetaDeletedAt:    paymentIntentItems.MetaDeletedAt.Time,
		MetaDeletedBy:    paymentIntentItems.MetaDeletedBy.UUID,
	}
}

type PaymentIntentItemsListResponse []*PaymentIntentItemsResponse

func NewPaymentIntentItemsListResponse(paymentIntentItemsList model.PaymentIntentItemsList) PaymentIntentItemsListResponse {
	dtoPaymentIntentItemsListResponse := PaymentIntentItemsListResponse{}
	for _, paymentIntentItems := range paymentIntentItemsList {
		dtoPaymentIntentItemsResponse := NewPaymentIntentItemsResponse(*paymentIntentItems)
		dtoPaymentIntentItemsListResponse = append(dtoPaymentIntentItemsListResponse, &dtoPaymentIntentItemsResponse)
	}
	return dtoPaymentIntentItemsListResponse
}

type PaymentIntentItemsPrimaryIDList []PaymentIntentItemsPrimaryID

func (d PaymentIntentItemsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentIntentItems := range d {
		err = validator.Struct(paymentIntentItems)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentIntentItemsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentIntentItemsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentIntentItemsPrimaryID) ToModel() model.PaymentIntentItemsPrimaryID {
	return model.PaymentIntentItemsPrimaryID{
		Id: d.Id,
	}
}
