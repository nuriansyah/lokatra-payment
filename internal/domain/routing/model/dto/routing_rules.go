package dto

import (
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/internal/domain/routing/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type RoutingRulesDTOFieldNameType string

type routingRulesDTOFieldName struct {
	Id                 RoutingRulesDTOFieldNameType
	ProfileId          RoutingRulesDTOFieldNameType
	Priority           RoutingRulesDTOFieldNameType
	Name               RoutingRulesDTOFieldNameType
	IsActive           RoutingRulesDTOFieldNameType
	MatchPaymentMethod RoutingRulesDTOFieldNameType
	MatchCurrency      RoutingRulesDTOFieldNameType
	MatchAmountMin     RoutingRulesDTOFieldNameType
	MatchAmountMax     RoutingRulesDTOFieldNameType
	MatchUserCountry   RoutingRulesDTOFieldNameType
	MatchCardBin       RoutingRulesDTOFieldNameType
	MatchProductType   RoutingRulesDTOFieldNameType
	CostWeight         RoutingRulesDTOFieldNameType
	Notes              RoutingRulesDTOFieldNameType
	MetaCreatedAt      RoutingRulesDTOFieldNameType
	MetaCreatedBy      RoutingRulesDTOFieldNameType
	MetaUpdatedAt      RoutingRulesDTOFieldNameType
	MetaUpdatedBy      RoutingRulesDTOFieldNameType
	MetaDeletedAt      RoutingRulesDTOFieldNameType
	MetaDeletedBy      RoutingRulesDTOFieldNameType
}

var RoutingRulesDTOFieldName = routingRulesDTOFieldName{
	Id:                 "id",
	ProfileId:          "profileId",
	Priority:           "priority",
	Name:               "name",
	IsActive:           "isActive",
	MatchPaymentMethod: "matchPaymentMethod",
	MatchCurrency:      "matchCurrency",
	MatchAmountMin:     "matchAmountMin",
	MatchAmountMax:     "matchAmountMax",
	MatchUserCountry:   "matchUserCountry",
	MatchCardBin:       "matchCardBin",
	MatchProductType:   "matchProductType",
	CostWeight:         "costWeight",
	Notes:              "notes",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func NewRoutingRulesListResponseFromFilterResult(result []model.RoutingRulesFilterResult, filter model.Filter) RoutingRulesSelectableListResponse {
	dtoRoutingRulesListResponse := RoutingRulesSelectableListResponse{}
	for _, routingRules := range result {
		dtoRoutingRulesResponse := NewRoutingRulesSelectableResponse(routingRules.RoutingRules, filter)
		dtoRoutingRulesListResponse = append(dtoRoutingRulesListResponse, &dtoRoutingRulesResponse)
	}
	return dtoRoutingRulesListResponse
}

func transformRoutingRulesDTOFieldNameFromStr(field string) (dbField model.RoutingRulesDBFieldNameType, found bool) {
	switch field {

	case string(RoutingRulesDTOFieldName.Id):
		return model.RoutingRulesDBFieldName.Id, true

	case string(RoutingRulesDTOFieldName.ProfileId):
		return model.RoutingRulesDBFieldName.ProfileId, true

	case string(RoutingRulesDTOFieldName.Priority):
		return model.RoutingRulesDBFieldName.Priority, true

	case string(RoutingRulesDTOFieldName.Name):
		return model.RoutingRulesDBFieldName.Name, true

	case string(RoutingRulesDTOFieldName.IsActive):
		return model.RoutingRulesDBFieldName.IsActive, true

	case string(RoutingRulesDTOFieldName.MatchPaymentMethod):
		return model.RoutingRulesDBFieldName.MatchPaymentMethod, true

	case string(RoutingRulesDTOFieldName.MatchCurrency):
		return model.RoutingRulesDBFieldName.MatchCurrency, true

	case string(RoutingRulesDTOFieldName.MatchAmountMin):
		return model.RoutingRulesDBFieldName.MatchAmountMin, true

	case string(RoutingRulesDTOFieldName.MatchAmountMax):
		return model.RoutingRulesDBFieldName.MatchAmountMax, true

	case string(RoutingRulesDTOFieldName.MatchUserCountry):
		return model.RoutingRulesDBFieldName.MatchUserCountry, true

	case string(RoutingRulesDTOFieldName.MatchCardBin):
		return model.RoutingRulesDBFieldName.MatchCardBin, true

	case string(RoutingRulesDTOFieldName.MatchProductType):
		return model.RoutingRulesDBFieldName.MatchProductType, true

	case string(RoutingRulesDTOFieldName.CostWeight):
		return model.RoutingRulesDBFieldName.CostWeight, true

	case string(RoutingRulesDTOFieldName.Notes):
		return model.RoutingRulesDBFieldName.Notes, true

	case string(RoutingRulesDTOFieldName.MetaCreatedAt):
		return model.RoutingRulesDBFieldName.MetaCreatedAt, true

	case string(RoutingRulesDTOFieldName.MetaCreatedBy):
		return model.RoutingRulesDBFieldName.MetaCreatedBy, true

	case string(RoutingRulesDTOFieldName.MetaUpdatedAt):
		return model.RoutingRulesDBFieldName.MetaUpdatedAt, true

	case string(RoutingRulesDTOFieldName.MetaUpdatedBy):
		return model.RoutingRulesDBFieldName.MetaUpdatedBy, true

	case string(RoutingRulesDTOFieldName.MetaDeletedAt):
		return model.RoutingRulesDBFieldName.MetaDeletedAt, true

	case string(RoutingRulesDTOFieldName.MetaDeletedBy):
		return model.RoutingRulesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformRoutingRulesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformRoutingRulesDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformRoutingRulesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformRoutingRulesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultRoutingRulesFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(RoutingRulesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type RoutingRulesSelectableResponse map[string]interface{}
type RoutingRulesSelectableListResponse []*RoutingRulesSelectableResponse

func NewRoutingRulesSelectableResponse(routingRules model.RoutingRules, filter model.Filter) RoutingRulesSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.RoutingRulesDBFieldName.Id),
			string(model.RoutingRulesDBFieldName.ProfileId),
			string(model.RoutingRulesDBFieldName.Priority),
			string(model.RoutingRulesDBFieldName.Name),
			string(model.RoutingRulesDBFieldName.IsActive),
			string(model.RoutingRulesDBFieldName.MatchPaymentMethod),
			string(model.RoutingRulesDBFieldName.MatchCurrency),
			string(model.RoutingRulesDBFieldName.MatchAmountMin),
			string(model.RoutingRulesDBFieldName.MatchAmountMax),
			string(model.RoutingRulesDBFieldName.MatchUserCountry),
			string(model.RoutingRulesDBFieldName.MatchCardBin),
			string(model.RoutingRulesDBFieldName.MatchProductType),
			string(model.RoutingRulesDBFieldName.CostWeight),
			string(model.RoutingRulesDBFieldName.Notes),
			string(model.RoutingRulesDBFieldName.MetaCreatedAt),
			string(model.RoutingRulesDBFieldName.MetaCreatedBy),
			string(model.RoutingRulesDBFieldName.MetaUpdatedAt),
			string(model.RoutingRulesDBFieldName.MetaUpdatedBy),
			string(model.RoutingRulesDBFieldName.MetaDeletedAt),
			string(model.RoutingRulesDBFieldName.MetaDeletedBy),
		)
	}
	routingRulesSelectableResponse := RoutingRulesSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.RoutingRulesDBFieldName.Id):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.Id)] = routingRules.Id

		case string(model.RoutingRulesDBFieldName.ProfileId):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.ProfileId)] = routingRules.ProfileId

		case string(model.RoutingRulesDBFieldName.Priority):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.Priority)] = routingRules.Priority

		case string(model.RoutingRulesDBFieldName.Name):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.Name)] = routingRules.Name

		case string(model.RoutingRulesDBFieldName.IsActive):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.IsActive)] = routingRules.IsActive

		case string(model.RoutingRulesDBFieldName.MatchPaymentMethod):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MatchPaymentMethod)] = routingRules.MatchPaymentMethod

		case string(model.RoutingRulesDBFieldName.MatchCurrency):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MatchCurrency)] = routingRules.MatchCurrency

		case string(model.RoutingRulesDBFieldName.MatchAmountMin):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MatchAmountMin)] = routingRules.MatchAmountMin

		case string(model.RoutingRulesDBFieldName.MatchAmountMax):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MatchAmountMax)] = routingRules.MatchAmountMax

		case string(model.RoutingRulesDBFieldName.MatchUserCountry):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MatchUserCountry)] = routingRules.MatchUserCountry

		case string(model.RoutingRulesDBFieldName.MatchCardBin):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MatchCardBin)] = routingRules.MatchCardBin

		case string(model.RoutingRulesDBFieldName.MatchProductType):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MatchProductType)] = routingRules.MatchProductType

		case string(model.RoutingRulesDBFieldName.CostWeight):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.CostWeight)] = routingRules.CostWeight

		case string(model.RoutingRulesDBFieldName.Notes):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.Notes)] = routingRules.Notes

		case string(model.RoutingRulesDBFieldName.MetaCreatedAt):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MetaCreatedAt)] = routingRules.MetaCreatedAt

		case string(model.RoutingRulesDBFieldName.MetaCreatedBy):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MetaCreatedBy)] = routingRules.MetaCreatedBy

		case string(model.RoutingRulesDBFieldName.MetaUpdatedAt):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MetaUpdatedAt)] = routingRules.MetaUpdatedAt

		case string(model.RoutingRulesDBFieldName.MetaUpdatedBy):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MetaUpdatedBy)] = routingRules.MetaUpdatedBy

		case string(model.RoutingRulesDBFieldName.MetaDeletedAt):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MetaDeletedAt)] = routingRules.MetaDeletedAt

		case string(model.RoutingRulesDBFieldName.MetaDeletedBy):
			routingRulesSelectableResponse[string(RoutingRulesDTOFieldName.MetaDeletedBy)] = routingRules.MetaDeletedBy

		}
	}
	return routingRulesSelectableResponse
}

type RoutingRulesFilterResponse struct {
	Metadata Metadata                           `json:"metadata"`
	Data     RoutingRulesSelectableListResponse `json:"data"`
}

func NewRoutingRulesFilterResponse(result []model.RoutingRulesFilterResult, filter model.Filter) (resp RoutingRulesFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewRoutingRulesListResponseFromFilterResult(result, filter)
	return resp
}

type RoutingRulesCreateRequest struct {
	ProfileId          uuid.UUID               `json:"profileId"`
	Priority           int16                   `json:"priority"`
	Name               string                  `json:"name"`
	IsActive           bool                    `json:"isActive"`
	MatchPaymentMethod model.PaymentMethodType `json:"matchPaymentMethod" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	MatchCurrency      model.PaymentCurrency   `json:"matchCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	MatchAmountMin     decimal.Decimal         `json:"matchAmountMin"`
	MatchAmountMax     decimal.Decimal         `json:"matchAmountMax"`
	MatchUserCountry   string                  `json:"matchUserCountry"`
	MatchCardBin       string                  `json:"matchCardBin"`
	MatchProductType   string                  `json:"matchProductType"`
	CostWeight         decimal.Decimal         `json:"costWeight"`
	Notes              string                  `json:"notes"`
	MetaCreatedAt      time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy      uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt      time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy      uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt      time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy      uuid.UUID               `json:"metaDeletedBy"`
}

func (d *RoutingRulesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *RoutingRulesCreateRequest) ToModel() model.RoutingRules {
	id, _ := uuid.NewV4()
	return model.RoutingRules{
		Id:                 id,
		ProfileId:          d.ProfileId,
		Priority:           d.Priority,
		Name:               d.Name,
		IsActive:           d.IsActive,
		MatchPaymentMethod: d.MatchPaymentMethod,
		MatchCurrency:      d.MatchCurrency,
		MatchAmountMin:     decimal.NewNullDecimal(d.MatchAmountMin),
		MatchAmountMax:     decimal.NewNullDecimal(d.MatchAmountMax),
		MatchUserCountry:   null.StringFrom(d.MatchUserCountry),
		MatchCardBin:       null.StringFrom(d.MatchCardBin),
		MatchProductType:   null.StringFrom(d.MatchProductType),
		CostWeight:         decimal.NewNullDecimal(d.CostWeight),
		Notes:              null.StringFrom(d.Notes),
		MetaCreatedAt:      d.MetaCreatedAt,
		MetaCreatedBy:      d.MetaCreatedBy,
		MetaUpdatedAt:      d.MetaUpdatedAt,
		MetaUpdatedBy:      nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:      null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:      nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingRulesListCreateRequest []*RoutingRulesCreateRequest

func (d RoutingRulesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingRules := range d {
		err = validator.Struct(routingRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RoutingRulesListCreateRequest) ToModelList() []model.RoutingRules {
	out := make([]model.RoutingRules, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type RoutingRulesUpdateRequest struct {
	ProfileId          uuid.UUID               `json:"profileId"`
	Priority           int16                   `json:"priority"`
	Name               string                  `json:"name"`
	IsActive           bool                    `json:"isActive"`
	MatchPaymentMethod model.PaymentMethodType `json:"matchPaymentMethod" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	MatchCurrency      model.PaymentCurrency   `json:"matchCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	MatchAmountMin     decimal.Decimal         `json:"matchAmountMin"`
	MatchAmountMax     decimal.Decimal         `json:"matchAmountMax"`
	MatchUserCountry   string                  `json:"matchUserCountry"`
	MatchCardBin       string                  `json:"matchCardBin"`
	MatchProductType   string                  `json:"matchProductType"`
	CostWeight         decimal.Decimal         `json:"costWeight"`
	Notes              string                  `json:"notes"`
	MetaCreatedAt      time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy      uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt      time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy      uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt      time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy      uuid.UUID               `json:"metaDeletedBy"`
}

func (d *RoutingRulesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d RoutingRulesUpdateRequest) ToModel() model.RoutingRules {
	return model.RoutingRules{
		ProfileId:          d.ProfileId,
		Priority:           d.Priority,
		Name:               d.Name,
		IsActive:           d.IsActive,
		MatchPaymentMethod: d.MatchPaymentMethod,
		MatchCurrency:      d.MatchCurrency,
		MatchAmountMin:     decimal.NewNullDecimal(d.MatchAmountMin),
		MatchAmountMax:     decimal.NewNullDecimal(d.MatchAmountMax),
		MatchUserCountry:   null.StringFrom(d.MatchUserCountry),
		MatchCardBin:       null.StringFrom(d.MatchCardBin),
		MatchProductType:   null.StringFrom(d.MatchProductType),
		CostWeight:         decimal.NewNullDecimal(d.CostWeight),
		Notes:              null.StringFrom(d.Notes),
		MetaCreatedAt:      d.MetaCreatedAt,
		MetaCreatedBy:      d.MetaCreatedBy,
		MetaUpdatedAt:      d.MetaUpdatedAt,
		MetaUpdatedBy:      nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:      null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:      nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingRulesBulkUpdateRequest struct {
	Id                 uuid.UUID               `json:"id"`
	ProfileId          uuid.UUID               `json:"profileId"`
	Priority           int16                   `json:"priority"`
	Name               string                  `json:"name"`
	IsActive           bool                    `json:"isActive"`
	MatchPaymentMethod model.PaymentMethodType `json:"matchPaymentMethod" example:"CARD" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	MatchCurrency      model.PaymentCurrency   `json:"matchCurrency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	MatchAmountMin     decimal.Decimal         `json:"matchAmountMin"`
	MatchAmountMax     decimal.Decimal         `json:"matchAmountMax"`
	MatchUserCountry   string                  `json:"matchUserCountry"`
	MatchCardBin       string                  `json:"matchCardBin"`
	MatchProductType   string                  `json:"matchProductType"`
	CostWeight         decimal.Decimal         `json:"costWeight"`
	Notes              string                  `json:"notes"`
	MetaCreatedAt      time.Time               `json:"metaCreatedAt"`
	MetaCreatedBy      uuid.UUID               `json:"metaCreatedBy"`
	MetaUpdatedAt      time.Time               `json:"metaUpdatedAt"`
	MetaUpdatedBy      uuid.UUID               `json:"metaUpdatedBy"`
	MetaDeletedAt      time.Time               `json:"metaDeletedAt"`
	MetaDeletedBy      uuid.UUID               `json:"metaDeletedBy"`
}

func (d RoutingRulesBulkUpdateRequest) PrimaryID() RoutingRulesPrimaryID {
	return RoutingRulesPrimaryID{
		Id: d.Id,
	}
}

type RoutingRulesListBulkUpdateRequest []*RoutingRulesBulkUpdateRequest

func (d RoutingRulesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingRules := range d {
		err = validator.Struct(routingRules)
		if err != nil {
			return
		}
	}
	return nil
}

func (d RoutingRulesBulkUpdateRequest) ToModel() model.RoutingRules {
	return model.RoutingRules{
		Id:                 d.Id,
		ProfileId:          d.ProfileId,
		Priority:           d.Priority,
		Name:               d.Name,
		IsActive:           d.IsActive,
		MatchPaymentMethod: d.MatchPaymentMethod,
		MatchCurrency:      d.MatchCurrency,
		MatchAmountMin:     decimal.NewNullDecimal(d.MatchAmountMin),
		MatchAmountMax:     decimal.NewNullDecimal(d.MatchAmountMax),
		MatchUserCountry:   null.StringFrom(d.MatchUserCountry),
		MatchCardBin:       null.StringFrom(d.MatchCardBin),
		MatchProductType:   null.StringFrom(d.MatchProductType),
		CostWeight:         decimal.NewNullDecimal(d.CostWeight),
		Notes:              null.StringFrom(d.Notes),
		MetaCreatedAt:      d.MetaCreatedAt,
		MetaCreatedBy:      d.MetaCreatedBy,
		MetaUpdatedAt:      d.MetaUpdatedAt,
		MetaUpdatedBy:      nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:      null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:      nuuid.From(d.MetaDeletedBy),
	}
}

type RoutingRulesResponse struct {
	Id                 uuid.UUID               `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProfileId          uuid.UUID               `json:"profileId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Priority           int16                   `json:"priority"`
	Name               string                  `json:"name" validate:"required"`
	IsActive           bool                    `json:"isActive" example:"true"`
	MatchPaymentMethod model.PaymentMethodType `json:"matchPaymentMethod" validate:"required,oneof=CARD VIRTUAL_ACCOUNT QRIS EWALLET DIRECT_DEBIT BANK_TRANSFER PAYLATER VOUCHER POINTS CASH_ON_DELIVERY" enums:"CARD,VIRTUAL_ACCOUNT,QRIS,EWALLET,DIRECT_DEBIT,BANK_TRANSFER,PAYLATER,VOUCHER,POINTS,CASH_ON_DELIVERY"`
	MatchCurrency      model.PaymentCurrency   `json:"matchCurrency" validate:"required,oneof=IDR USD SGD MYR PHP THB AED EUR GBP JPY" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	MatchAmountMin     decimal.Decimal         `json:"matchAmountMin" format:"decimal" example:"100.50"`
	MatchAmountMax     decimal.Decimal         `json:"matchAmountMax" format:"decimal" example:"100.50"`
	MatchUserCountry   string                  `json:"matchUserCountry"`
	MatchCardBin       string                  `json:"matchCardBin"`
	MatchProductType   string                  `json:"matchProductType"`
	CostWeight         decimal.Decimal         `json:"costWeight" format:"decimal" example:"100.50"`
	Notes              string                  `json:"notes"`
	MetaCreatedAt      time.Time               `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy      uuid.UUID               `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt      time.Time               `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy      uuid.UUID               `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt      time.Time               `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy      uuid.UUID               `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewRoutingRulesResponse(routingRules model.RoutingRules) RoutingRulesResponse {
	return RoutingRulesResponse{
		Id:                 routingRules.Id,
		ProfileId:          routingRules.ProfileId,
		Priority:           routingRules.Priority,
		Name:               routingRules.Name,
		IsActive:           routingRules.IsActive,
		MatchPaymentMethod: model.PaymentMethodType(routingRules.MatchPaymentMethod),
		MatchCurrency:      model.PaymentCurrency(routingRules.MatchCurrency),
		MatchAmountMin:     routingRules.MatchAmountMin.Decimal,
		MatchAmountMax:     routingRules.MatchAmountMax.Decimal,
		MatchUserCountry:   routingRules.MatchUserCountry.String,
		MatchCardBin:       routingRules.MatchCardBin.String,
		MatchProductType:   routingRules.MatchProductType.String,
		CostWeight:         routingRules.CostWeight.Decimal,
		Notes:              routingRules.Notes.String,
		MetaCreatedAt:      routingRules.MetaCreatedAt,
		MetaCreatedBy:      routingRules.MetaCreatedBy,
		MetaUpdatedAt:      routingRules.MetaUpdatedAt,
		MetaUpdatedBy:      routingRules.MetaUpdatedBy.UUID,
		MetaDeletedAt:      routingRules.MetaDeletedAt.Time,
		MetaDeletedBy:      routingRules.MetaDeletedBy.UUID,
	}
}

type RoutingRulesListResponse []*RoutingRulesResponse

func NewRoutingRulesListResponse(routingRulesList model.RoutingRulesList) RoutingRulesListResponse {
	dtoRoutingRulesListResponse := RoutingRulesListResponse{}
	for _, routingRules := range routingRulesList {
		dtoRoutingRulesResponse := NewRoutingRulesResponse(*routingRules)
		dtoRoutingRulesListResponse = append(dtoRoutingRulesListResponse, &dtoRoutingRulesResponse)
	}
	return dtoRoutingRulesListResponse
}

type RoutingRulesPrimaryIDList []RoutingRulesPrimaryID

func (d RoutingRulesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, routingRules := range d {
		err = validator.Struct(routingRules)
		if err != nil {
			return
		}
	}
	return nil
}

type RoutingRulesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *RoutingRulesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d RoutingRulesPrimaryID) ToModel() model.RoutingRulesPrimaryID {
	return model.RoutingRulesPrimaryID{
		Id: d.Id,
	}
}
