package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ChargebacksDTOFieldNameType string

type chargebacksDTOFieldName struct {
	Id                ChargebacksDTOFieldNameType
	ChargebackCode    ChargebacksDTOFieldNameType
	PaymentRefId      ChargebacksDTOFieldNameType
	MerchantPartyId   ChargebacksDTOFieldNameType
	ProviderAccountId ChargebacksDTOFieldNameType
	CurrencyCode      ChargebacksDTOFieldNameType
	DisputedAmount    ChargebacksDTOFieldNameType
	ChargebackStatus  ChargebacksDTOFieldNameType
	ReasonCode        ChargebacksDTOFieldNameType
	OpenedAt          ChargebacksDTOFieldNameType
	ClosedAt          ChargebacksDTOFieldNameType
	DueAt             ChargebacksDTOFieldNameType
	Metadata          ChargebacksDTOFieldNameType
	MetaCreatedAt     ChargebacksDTOFieldNameType
	MetaCreatedBy     ChargebacksDTOFieldNameType
	MetaUpdatedAt     ChargebacksDTOFieldNameType
	MetaUpdatedBy     ChargebacksDTOFieldNameType
	MetaDeletedAt     ChargebacksDTOFieldNameType
	MetaDeletedBy     ChargebacksDTOFieldNameType
}

var ChargebacksDTOFieldName = chargebacksDTOFieldName{
	Id:                "id",
	ChargebackCode:    "chargebackCode",
	PaymentRefId:      "paymentRefId",
	MerchantPartyId:   "merchantPartyId",
	ProviderAccountId: "providerAccountId",
	CurrencyCode:      "currencyCode",
	DisputedAmount:    "disputedAmount",
	ChargebackStatus:  "chargebackStatus",
	ReasonCode:        "reasonCode",
	OpenedAt:          "openedAt",
	ClosedAt:          "closedAt",
	DueAt:             "dueAt",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformChargebacksDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ChargebacksDTOFieldName.Id):
		return string(model.ChargebacksDBFieldName.Id), true

	case string(ChargebacksDTOFieldName.ChargebackCode):
		return string(model.ChargebacksDBFieldName.ChargebackCode), true

	case string(ChargebacksDTOFieldName.PaymentRefId):
		return string(model.ChargebacksDBFieldName.PaymentRefId), true

	case string(ChargebacksDTOFieldName.MerchantPartyId):
		return string(model.ChargebacksDBFieldName.MerchantPartyId), true

	case string(ChargebacksDTOFieldName.ProviderAccountId):
		return string(model.ChargebacksDBFieldName.ProviderAccountId), true

	case string(ChargebacksDTOFieldName.CurrencyCode):
		return string(model.ChargebacksDBFieldName.CurrencyCode), true

	case string(ChargebacksDTOFieldName.DisputedAmount):
		return string(model.ChargebacksDBFieldName.DisputedAmount), true

	case string(ChargebacksDTOFieldName.ChargebackStatus):
		return string(model.ChargebacksDBFieldName.ChargebackStatus), true

	case string(ChargebacksDTOFieldName.ReasonCode):
		return string(model.ChargebacksDBFieldName.ReasonCode), true

	case string(ChargebacksDTOFieldName.OpenedAt):
		return string(model.ChargebacksDBFieldName.OpenedAt), true

	case string(ChargebacksDTOFieldName.ClosedAt):
		return string(model.ChargebacksDBFieldName.ClosedAt), true

	case string(ChargebacksDTOFieldName.DueAt):
		return string(model.ChargebacksDBFieldName.DueAt), true

	case string(ChargebacksDTOFieldName.Metadata):
		return string(model.ChargebacksDBFieldName.Metadata), true

	case string(ChargebacksDTOFieldName.MetaCreatedAt):
		return string(model.ChargebacksDBFieldName.MetaCreatedAt), true

	case string(ChargebacksDTOFieldName.MetaCreatedBy):
		return string(model.ChargebacksDBFieldName.MetaCreatedBy), true

	case string(ChargebacksDTOFieldName.MetaUpdatedAt):
		return string(model.ChargebacksDBFieldName.MetaUpdatedAt), true

	case string(ChargebacksDTOFieldName.MetaUpdatedBy):
		return string(model.ChargebacksDBFieldName.MetaUpdatedBy), true

	case string(ChargebacksDTOFieldName.MetaDeletedAt):
		return string(model.ChargebacksDBFieldName.MetaDeletedAt), true

	case string(ChargebacksDTOFieldName.MetaDeletedBy):
		return string(model.ChargebacksDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewChargebacksFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isChargebacksBaseFilterField(field string) bool {
	spec, found := model.NewChargebacksFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeChargebacksProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateChargebacksProjectionOutputPath(path string) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return failure.BadRequest(fmt.Errorf("field alias cannot be empty"))
	}
	if !strings.Contains(path, ".") {
		return nil
	}
	for _, part := range strings.Split(path, ".") {
		if strings.TrimSpace(part) == "" {
			return failure.BadRequest(fmt.Errorf("field alias %s is invalid", path))
		}
	}
	return nil
}

func transformChargebacksFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformChargebacksDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformChargebacksFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformChargebacksFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformChargebacksDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isChargebacksBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateChargebacksProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeChargebacksProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformChargebacksDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformChargebacksDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformChargebacksFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultChargebacksFilter(filter *model.Filter) {
	if filter.Pagination.Strategy == "" {
		if filter.Pagination.Page > 0 {
			filter.Pagination.Strategy = model.PaginationStrategyOffset
		} else {
			filter.Pagination.Strategy = model.DefaultPaginationStrategy
		}
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = model.DefaultPageSize
	}

	if filter.Pagination.PageSize > model.MaxPageSize {
		filter.Pagination.PageSize = model.MaxPageSize
	}

	if filter.Pagination.Strategy == model.PaginationStrategyOffset && filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.Direction == "" {
		filter.Pagination.Direction = model.CursorDirectionNext
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(ChargebacksDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ChargebacksSelectableResponse map[string]interface{}
type ChargebacksSelectableListResponse []*ChargebacksSelectableResponse

func assignChargebacksNestedValue(out map[string]interface{}, path string, value interface{}) {
	parts := strings.Split(path, ".")
	current := out
	for _, part := range parts[:len(parts)-1] {
		next, ok := current[part].(map[string]interface{})
		if !ok {
			next = map[string]interface{}{}
			current[part] = next
		}
		current = next
	}
	current[parts[len(parts)-1]] = value
}

func setChargebacksSelectableValue(out ChargebacksSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignChargebacksNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewChargebacksSelectableResponse(chargebacks model.Chargebacks, filter model.Filter) ChargebacksSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ChargebacksDBFieldName.Id),
			string(model.ChargebacksDBFieldName.ChargebackCode),
			string(model.ChargebacksDBFieldName.PaymentRefId),
			string(model.ChargebacksDBFieldName.MerchantPartyId),
			string(model.ChargebacksDBFieldName.ProviderAccountId),
			string(model.ChargebacksDBFieldName.CurrencyCode),
			string(model.ChargebacksDBFieldName.DisputedAmount),
			string(model.ChargebacksDBFieldName.ChargebackStatus),
			string(model.ChargebacksDBFieldName.ReasonCode),
			string(model.ChargebacksDBFieldName.OpenedAt),
			string(model.ChargebacksDBFieldName.ClosedAt),
			string(model.ChargebacksDBFieldName.DueAt),
			string(model.ChargebacksDBFieldName.Metadata),
			string(model.ChargebacksDBFieldName.MetaCreatedAt),
			string(model.ChargebacksDBFieldName.MetaCreatedBy),
			string(model.ChargebacksDBFieldName.MetaUpdatedAt),
			string(model.ChargebacksDBFieldName.MetaUpdatedBy),
			string(model.ChargebacksDBFieldName.MetaDeletedAt),
			string(model.ChargebacksDBFieldName.MetaDeletedBy),
		)
	}
	chargebacksSelectableResponse := ChargebacksSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ChargebacksDBFieldName.Id):
			key := string(ChargebacksDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.Id, explicitAlias)

		case string(model.ChargebacksDBFieldName.ChargebackCode):
			key := string(ChargebacksDTOFieldName.ChargebackCode)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.ChargebackCode, explicitAlias)

		case string(model.ChargebacksDBFieldName.PaymentRefId):
			key := string(ChargebacksDTOFieldName.PaymentRefId)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.PaymentRefId, explicitAlias)

		case string(model.ChargebacksDBFieldName.MerchantPartyId):
			key := string(ChargebacksDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.MerchantPartyId.UUID, explicitAlias)

		case string(model.ChargebacksDBFieldName.ProviderAccountId):
			key := string(ChargebacksDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.ProviderAccountId.UUID, explicitAlias)

		case string(model.ChargebacksDBFieldName.CurrencyCode):
			key := string(ChargebacksDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.CurrencyCode, explicitAlias)

		case string(model.ChargebacksDBFieldName.DisputedAmount):
			key := string(ChargebacksDTOFieldName.DisputedAmount)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.DisputedAmount, explicitAlias)

		case string(model.ChargebacksDBFieldName.ChargebackStatus):
			key := string(ChargebacksDTOFieldName.ChargebackStatus)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, model.ChargebackStatus(chargebacks.ChargebackStatus), explicitAlias)

		case string(model.ChargebacksDBFieldName.ReasonCode):
			key := string(ChargebacksDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.ReasonCode.String, explicitAlias)

		case string(model.ChargebacksDBFieldName.OpenedAt):
			key := string(ChargebacksDTOFieldName.OpenedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.OpenedAt, explicitAlias)

		case string(model.ChargebacksDBFieldName.ClosedAt):
			key := string(ChargebacksDTOFieldName.ClosedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.ClosedAt.Time, explicitAlias)

		case string(model.ChargebacksDBFieldName.DueAt):
			key := string(ChargebacksDTOFieldName.DueAt)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.DueAt.Time, explicitAlias)

		case string(model.ChargebacksDBFieldName.Metadata):
			key := string(ChargebacksDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.Metadata, explicitAlias)

		case string(model.ChargebacksDBFieldName.MetaCreatedAt):
			key := string(ChargebacksDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.MetaCreatedAt, explicitAlias)

		case string(model.ChargebacksDBFieldName.MetaCreatedBy):
			key := string(ChargebacksDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.MetaCreatedBy, explicitAlias)

		case string(model.ChargebacksDBFieldName.MetaUpdatedAt):
			key := string(ChargebacksDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.MetaUpdatedAt, explicitAlias)

		case string(model.ChargebacksDBFieldName.MetaUpdatedBy):
			key := string(ChargebacksDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.MetaUpdatedBy, explicitAlias)

		case string(model.ChargebacksDBFieldName.MetaDeletedAt):
			key := string(ChargebacksDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.MetaDeletedAt.Time, explicitAlias)

		case string(model.ChargebacksDBFieldName.MetaDeletedBy):
			key := string(ChargebacksDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setChargebacksSelectableValue(chargebacksSelectableResponse, key, chargebacks.MetaDeletedBy, explicitAlias)

		}
	}
	return chargebacksSelectableResponse
}

func NewChargebacksListResponseFromFilterResult(result []model.ChargebacksFilterResult, filter model.Filter) ChargebacksSelectableListResponse {
	dtoChargebacksListResponse := ChargebacksSelectableListResponse{}
	for _, row := range result {
		dtoChargebacksResponse := NewChargebacksSelectableResponse(row.Chargebacks, filter)
		dtoChargebacksListResponse = append(dtoChargebacksListResponse, &dtoChargebacksResponse)
	}
	return dtoChargebacksListResponse
}

type ChargebacksFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     ChargebacksSelectableListResponse `json:"data"`
}

func reverseChargebacksFilterResults(result []model.ChargebacksFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewChargebacksFilterResponse(result []model.ChargebacksFilterResult, filter model.Filter) (resp ChargebacksFilterResponse) {
	resp.Metadata.Strategy = filter.Pagination.Strategy
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	dataResult := result

	if filter.Pagination.IsCursorMode() && filter.Pagination.PageSize > 0 {
		if len(dataResult) > filter.Pagination.PageSize {
			resp.Metadata.HasMore = true
			if filter.Pagination.Direction == model.CursorDirectionPrev {
				resp.Metadata.HasPrev = true
			} else {
				resp.Metadata.HasNext = true
			}
			dataResult = dataResult[:filter.Pagination.PageSize]
		}
		if filter.Pagination.Direction == model.CursorDirectionPrev {
			reverseChargebacksFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].Id
			resp.Metadata.PrevCursor = dataResult[0].Id
		}
		if resp.Metadata.Page <= 0 {
			resp.Metadata.Page = 1
		}
	} else {
		if len(dataResult) > 0 {
			resp.Metadata.TotalData = dataResult[0].FilterCount
			resp.Metadata.TotalPage = int(math.Ceil(float64(resp.Metadata.TotalData) / float64(filter.Pagination.PageSize)))
			resp.Metadata.HasPrev = filter.Pagination.Page > 1
			resp.Metadata.HasNext = filter.Pagination.Page < resp.Metadata.TotalPage
			resp.Metadata.HasMore = resp.Metadata.HasNext
		}
	}

	resp.Data = NewChargebacksListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ChargebacksCreateRequest struct {
	ChargebackCode    string                 `json:"chargebackCode"`
	PaymentRefId      uuid.UUID              `json:"paymentRefId"`
	MerchantPartyId   uuid.UUID              `json:"merchantPartyId"`
	ProviderAccountId uuid.UUID              `json:"providerAccountId"`
	CurrencyCode      string                 `json:"currencyCode"`
	DisputedAmount    decimal.Decimal        `json:"disputedAmount"`
	ChargebackStatus  model.ChargebackStatus `json:"chargebackStatus" example:"opened" enums:"opened,under_review,won,lost,partially_won,reversed,closed"`
	ReasonCode        string                 `json:"reasonCode"`
	OpenedAt          time.Time              `json:"openedAt"`
	ClosedAt          time.Time              `json:"closedAt"`
	DueAt             time.Time              `json:"dueAt"`
	Metadata          json.RawMessage        `json:"metadata"`
}

func (d *ChargebacksCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ChargebacksCreateRequest) ToModel() model.Chargebacks {
	id, _ := uuid.NewV4()
	return model.Chargebacks{
		Id:                id,
		ChargebackCode:    d.ChargebackCode,
		PaymentRefId:      d.PaymentRefId,
		MerchantPartyId:   nuuid.From(d.MerchantPartyId),
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		CurrencyCode:      d.CurrencyCode,
		DisputedAmount:    d.DisputedAmount,
		ChargebackStatus:  d.ChargebackStatus,
		ReasonCode:        null.StringFrom(d.ReasonCode),
		OpenedAt:          d.OpenedAt,
		ClosedAt:          null.TimeFrom(d.ClosedAt),
		DueAt:             null.TimeFrom(d.DueAt),
		Metadata:          d.Metadata,
	}
}

type ChargebacksListCreateRequest []*ChargebacksCreateRequest

func (d ChargebacksListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, chargebacks := range d {
		err = validator.Struct(chargebacks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ChargebacksListCreateRequest) ToModelList() []model.Chargebacks {
	out := make([]model.Chargebacks, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ChargebacksUpdateRequest struct {
	ChargebackCode    string                 `json:"chargebackCode"`
	PaymentRefId      uuid.UUID              `json:"paymentRefId"`
	MerchantPartyId   uuid.UUID              `json:"merchantPartyId"`
	ProviderAccountId uuid.UUID              `json:"providerAccountId"`
	CurrencyCode      string                 `json:"currencyCode"`
	DisputedAmount    decimal.Decimal        `json:"disputedAmount"`
	ChargebackStatus  model.ChargebackStatus `json:"chargebackStatus" example:"opened" enums:"opened,under_review,won,lost,partially_won,reversed,closed"`
	ReasonCode        string                 `json:"reasonCode"`
	OpenedAt          time.Time              `json:"openedAt"`
	ClosedAt          time.Time              `json:"closedAt"`
	DueAt             time.Time              `json:"dueAt"`
	Metadata          json.RawMessage        `json:"metadata"`
}

func (d *ChargebacksUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ChargebacksUpdateRequest) ToModel() model.Chargebacks {
	return model.Chargebacks{
		ChargebackCode:    d.ChargebackCode,
		PaymentRefId:      d.PaymentRefId,
		MerchantPartyId:   nuuid.From(d.MerchantPartyId),
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		CurrencyCode:      d.CurrencyCode,
		DisputedAmount:    d.DisputedAmount,
		ChargebackStatus:  d.ChargebackStatus,
		ReasonCode:        null.StringFrom(d.ReasonCode),
		OpenedAt:          d.OpenedAt,
		ClosedAt:          null.TimeFrom(d.ClosedAt),
		DueAt:             null.TimeFrom(d.DueAt),
		Metadata:          d.Metadata,
	}
}

type ChargebacksBulkUpdateRequest struct {
	Id                uuid.UUID              `json:"id"`
	ChargebackCode    string                 `json:"chargebackCode"`
	PaymentRefId      uuid.UUID              `json:"paymentRefId"`
	MerchantPartyId   uuid.UUID              `json:"merchantPartyId"`
	ProviderAccountId uuid.UUID              `json:"providerAccountId"`
	CurrencyCode      string                 `json:"currencyCode"`
	DisputedAmount    decimal.Decimal        `json:"disputedAmount"`
	ChargebackStatus  model.ChargebackStatus `json:"chargebackStatus" example:"opened" enums:"opened,under_review,won,lost,partially_won,reversed,closed"`
	ReasonCode        string                 `json:"reasonCode"`
	OpenedAt          time.Time              `json:"openedAt"`
	ClosedAt          time.Time              `json:"closedAt"`
	DueAt             time.Time              `json:"dueAt"`
	Metadata          json.RawMessage        `json:"metadata"`
}

func (d ChargebacksBulkUpdateRequest) PrimaryID() ChargebacksPrimaryID {
	return ChargebacksPrimaryID{
		Id: d.Id,
	}
}

type ChargebacksListBulkUpdateRequest []*ChargebacksBulkUpdateRequest

func (d ChargebacksListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, chargebacks := range d {
		err = validator.Struct(chargebacks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ChargebacksBulkUpdateRequest) ToModel() model.Chargebacks {
	return model.Chargebacks{
		Id:                d.Id,
		ChargebackCode:    d.ChargebackCode,
		PaymentRefId:      d.PaymentRefId,
		MerchantPartyId:   nuuid.From(d.MerchantPartyId),
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		CurrencyCode:      d.CurrencyCode,
		DisputedAmount:    d.DisputedAmount,
		ChargebackStatus:  d.ChargebackStatus,
		ReasonCode:        null.StringFrom(d.ReasonCode),
		OpenedAt:          d.OpenedAt,
		ClosedAt:          null.TimeFrom(d.ClosedAt),
		DueAt:             null.TimeFrom(d.DueAt),
		Metadata:          d.Metadata,
	}
}

type ChargebacksResponse struct {
	Id                uuid.UUID              `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ChargebackCode    string                 `json:"chargebackCode" validate:"required"`
	PaymentRefId      uuid.UUID              `json:"paymentRefId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId   uuid.UUID              `json:"merchantPartyId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId uuid.UUID              `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode      string                 `json:"currencyCode" validate:"required"`
	DisputedAmount    decimal.Decimal        `json:"disputedAmount" validate:"required" format:"decimal" example:"100.50"`
	ChargebackStatus  model.ChargebackStatus `json:"chargebackStatus" validate:"required,oneof=opened under_review won lost partially_won reversed closed" enums:"opened,under_review,won,lost,partially_won,reversed,closed"`
	ReasonCode        string                 `json:"reasonCode"`
	OpenedAt          time.Time              `json:"openedAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ClosedAt          time.Time              `json:"closedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	DueAt             time.Time              `json:"dueAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata          json.RawMessage        `json:"metadata" swaggertype:"object"`
}

func NewChargebacksResponse(chargebacks model.Chargebacks) ChargebacksResponse {
	return ChargebacksResponse{
		Id:                chargebacks.Id,
		ChargebackCode:    chargebacks.ChargebackCode,
		PaymentRefId:      chargebacks.PaymentRefId,
		MerchantPartyId:   chargebacks.MerchantPartyId.UUID,
		ProviderAccountId: chargebacks.ProviderAccountId.UUID,
		CurrencyCode:      chargebacks.CurrencyCode,
		DisputedAmount:    chargebacks.DisputedAmount,
		ChargebackStatus:  model.ChargebackStatus(chargebacks.ChargebackStatus),
		ReasonCode:        chargebacks.ReasonCode.String,
		OpenedAt:          chargebacks.OpenedAt,
		ClosedAt:          chargebacks.ClosedAt.Time,
		DueAt:             chargebacks.DueAt.Time,
		Metadata:          chargebacks.Metadata,
	}
}

type ChargebacksListResponse []*ChargebacksResponse

func NewChargebacksListResponse(chargebacksList model.ChargebacksList) ChargebacksListResponse {
	dtoChargebacksListResponse := ChargebacksListResponse{}
	for _, chargebacks := range chargebacksList {
		dtoChargebacksResponse := NewChargebacksResponse(*chargebacks)
		dtoChargebacksListResponse = append(dtoChargebacksListResponse, &dtoChargebacksResponse)
	}
	return dtoChargebacksListResponse
}

type ChargebacksPrimaryIDList []ChargebacksPrimaryID

func (d ChargebacksPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, chargebacks := range d {
		err = validator.Struct(chargebacks)
		if err != nil {
			return
		}
	}
	return nil
}

type ChargebacksPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ChargebacksPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ChargebacksPrimaryID) ToModel() model.ChargebacksPrimaryID {
	return model.ChargebacksPrimaryID{
		Id: d.Id,
	}
}
