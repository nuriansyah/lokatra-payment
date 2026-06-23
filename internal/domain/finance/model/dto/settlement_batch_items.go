package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type SettlementBatchItemsDTOFieldNameType string

type settlementBatchItemsDTOFieldName struct {
	Id                   SettlementBatchItemsDTOFieldNameType
	SettlementBatchId    SettlementBatchItemsDTOFieldNameType
	SourceType           SettlementBatchItemsDTOFieldNameType
	SourceId             SettlementBatchItemsDTOFieldNameType
	MerchantPartyId      SettlementBatchItemsDTOFieldNameType
	CurrencyCode         SettlementBatchItemsDTOFieldNameType
	GrossAmount          SettlementBatchItemsDTOFieldNameType
	FeeAmount            SettlementBatchItemsDTOFieldNameType
	TaxAmount            SettlementBatchItemsDTOFieldNameType
	ReserveAmount        SettlementBatchItemsDTOFieldNameType
	NetAmount            SettlementBatchItemsDTOFieldNameType
	LinkedJournalEntryId SettlementBatchItemsDTOFieldNameType
	ItemStatus           SettlementBatchItemsDTOFieldNameType
	Metadata             SettlementBatchItemsDTOFieldNameType
	MetaCreatedAt        SettlementBatchItemsDTOFieldNameType
	MetaCreatedBy        SettlementBatchItemsDTOFieldNameType
	MetaUpdatedAt        SettlementBatchItemsDTOFieldNameType
	MetaUpdatedBy        SettlementBatchItemsDTOFieldNameType
	MetaDeletedAt        SettlementBatchItemsDTOFieldNameType
	MetaDeletedBy        SettlementBatchItemsDTOFieldNameType
}

var SettlementBatchItemsDTOFieldName = settlementBatchItemsDTOFieldName{
	Id:                   "id",
	SettlementBatchId:    "settlementBatchId",
	SourceType:           "sourceType",
	SourceId:             "sourceId",
	MerchantPartyId:      "merchantPartyId",
	CurrencyCode:         "currencyCode",
	GrossAmount:          "grossAmount",
	FeeAmount:            "feeAmount",
	TaxAmount:            "taxAmount",
	ReserveAmount:        "reserveAmount",
	NetAmount:            "netAmount",
	LinkedJournalEntryId: "linkedJournalEntryId",
	ItemStatus:           "itemStatus",
	Metadata:             "metadata",
	MetaCreatedAt:        "metaCreatedAt",
	MetaCreatedBy:        "metaCreatedBy",
	MetaUpdatedAt:        "metaUpdatedAt",
	MetaUpdatedBy:        "metaUpdatedBy",
	MetaDeletedAt:        "metaDeletedAt",
	MetaDeletedBy:        "metaDeletedBy",
}

func transformSettlementBatchItemsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(SettlementBatchItemsDTOFieldName.Id):
		return string(model.SettlementBatchItemsDBFieldName.Id), true

	case string(SettlementBatchItemsDTOFieldName.SettlementBatchId):
		return string(model.SettlementBatchItemsDBFieldName.SettlementBatchId), true

	case string(SettlementBatchItemsDTOFieldName.SourceType):
		return string(model.SettlementBatchItemsDBFieldName.SourceType), true

	case string(SettlementBatchItemsDTOFieldName.SourceId):
		return string(model.SettlementBatchItemsDBFieldName.SourceId), true

	case string(SettlementBatchItemsDTOFieldName.MerchantPartyId):
		return string(model.SettlementBatchItemsDBFieldName.MerchantPartyId), true

	case string(SettlementBatchItemsDTOFieldName.CurrencyCode):
		return string(model.SettlementBatchItemsDBFieldName.CurrencyCode), true

	case string(SettlementBatchItemsDTOFieldName.GrossAmount):
		return string(model.SettlementBatchItemsDBFieldName.GrossAmount), true

	case string(SettlementBatchItemsDTOFieldName.FeeAmount):
		return string(model.SettlementBatchItemsDBFieldName.FeeAmount), true

	case string(SettlementBatchItemsDTOFieldName.TaxAmount):
		return string(model.SettlementBatchItemsDBFieldName.TaxAmount), true

	case string(SettlementBatchItemsDTOFieldName.ReserveAmount):
		return string(model.SettlementBatchItemsDBFieldName.ReserveAmount), true

	case string(SettlementBatchItemsDTOFieldName.NetAmount):
		return string(model.SettlementBatchItemsDBFieldName.NetAmount), true

	case string(SettlementBatchItemsDTOFieldName.LinkedJournalEntryId):
		return string(model.SettlementBatchItemsDBFieldName.LinkedJournalEntryId), true

	case string(SettlementBatchItemsDTOFieldName.ItemStatus):
		return string(model.SettlementBatchItemsDBFieldName.ItemStatus), true

	case string(SettlementBatchItemsDTOFieldName.Metadata):
		return string(model.SettlementBatchItemsDBFieldName.Metadata), true

	case string(SettlementBatchItemsDTOFieldName.MetaCreatedAt):
		return string(model.SettlementBatchItemsDBFieldName.MetaCreatedAt), true

	case string(SettlementBatchItemsDTOFieldName.MetaCreatedBy):
		return string(model.SettlementBatchItemsDBFieldName.MetaCreatedBy), true

	case string(SettlementBatchItemsDTOFieldName.MetaUpdatedAt):
		return string(model.SettlementBatchItemsDBFieldName.MetaUpdatedAt), true

	case string(SettlementBatchItemsDTOFieldName.MetaUpdatedBy):
		return string(model.SettlementBatchItemsDBFieldName.MetaUpdatedBy), true

	case string(SettlementBatchItemsDTOFieldName.MetaDeletedAt):
		return string(model.SettlementBatchItemsDBFieldName.MetaDeletedAt), true

	case string(SettlementBatchItemsDTOFieldName.MetaDeletedBy):
		return string(model.SettlementBatchItemsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewSettlementBatchItemsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isSettlementBatchItemsBaseFilterField(field string) bool {
	spec, found := model.NewSettlementBatchItemsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeSettlementBatchItemsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateSettlementBatchItemsProjectionOutputPath(path string) error {
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

func transformSettlementBatchItemsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformSettlementBatchItemsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformSettlementBatchItemsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformSettlementBatchItemsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformSettlementBatchItemsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isSettlementBatchItemsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateSettlementBatchItemsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeSettlementBatchItemsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformSettlementBatchItemsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformSettlementBatchItemsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformSettlementBatchItemsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultSettlementBatchItemsFilter(filter *model.Filter) {
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
			Field: string(SettlementBatchItemsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type SettlementBatchItemsSelectableResponse map[string]interface{}
type SettlementBatchItemsSelectableListResponse []*SettlementBatchItemsSelectableResponse

func assignSettlementBatchItemsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setSettlementBatchItemsSelectableValue(out SettlementBatchItemsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignSettlementBatchItemsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewSettlementBatchItemsSelectableResponse(settlementBatchItems model.SettlementBatchItems, filter model.Filter) SettlementBatchItemsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.SettlementBatchItemsDBFieldName.Id),
			string(model.SettlementBatchItemsDBFieldName.SettlementBatchId),
			string(model.SettlementBatchItemsDBFieldName.SourceType),
			string(model.SettlementBatchItemsDBFieldName.SourceId),
			string(model.SettlementBatchItemsDBFieldName.MerchantPartyId),
			string(model.SettlementBatchItemsDBFieldName.CurrencyCode),
			string(model.SettlementBatchItemsDBFieldName.GrossAmount),
			string(model.SettlementBatchItemsDBFieldName.FeeAmount),
			string(model.SettlementBatchItemsDBFieldName.TaxAmount),
			string(model.SettlementBatchItemsDBFieldName.ReserveAmount),
			string(model.SettlementBatchItemsDBFieldName.NetAmount),
			string(model.SettlementBatchItemsDBFieldName.LinkedJournalEntryId),
			string(model.SettlementBatchItemsDBFieldName.ItemStatus),
			string(model.SettlementBatchItemsDBFieldName.Metadata),
			string(model.SettlementBatchItemsDBFieldName.MetaCreatedAt),
			string(model.SettlementBatchItemsDBFieldName.MetaCreatedBy),
			string(model.SettlementBatchItemsDBFieldName.MetaUpdatedAt),
			string(model.SettlementBatchItemsDBFieldName.MetaUpdatedBy),
			string(model.SettlementBatchItemsDBFieldName.MetaDeletedAt),
			string(model.SettlementBatchItemsDBFieldName.MetaDeletedBy),
		)
	}
	settlementBatchItemsSelectableResponse := SettlementBatchItemsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.SettlementBatchItemsDBFieldName.Id):
			key := string(SettlementBatchItemsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.Id, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.SettlementBatchId):
			key := string(SettlementBatchItemsDTOFieldName.SettlementBatchId)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.SettlementBatchId, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.SourceType):
			key := string(SettlementBatchItemsDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.SourceType, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.SourceId):
			key := string(SettlementBatchItemsDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.SourceId, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.MerchantPartyId):
			key := string(SettlementBatchItemsDTOFieldName.MerchantPartyId)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.MerchantPartyId, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.CurrencyCode):
			key := string(SettlementBatchItemsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.CurrencyCode, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.GrossAmount):
			key := string(SettlementBatchItemsDTOFieldName.GrossAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.GrossAmount, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.FeeAmount):
			key := string(SettlementBatchItemsDTOFieldName.FeeAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.FeeAmount, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.TaxAmount):
			key := string(SettlementBatchItemsDTOFieldName.TaxAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.TaxAmount, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.ReserveAmount):
			key := string(SettlementBatchItemsDTOFieldName.ReserveAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.ReserveAmount, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.NetAmount):
			key := string(SettlementBatchItemsDTOFieldName.NetAmount)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.NetAmount, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.LinkedJournalEntryId):
			key := string(SettlementBatchItemsDTOFieldName.LinkedJournalEntryId)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.LinkedJournalEntryId.UUID, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.ItemStatus):
			key := string(SettlementBatchItemsDTOFieldName.ItemStatus)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, model.ItemStatus(settlementBatchItems.ItemStatus), explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.Metadata):
			key := string(SettlementBatchItemsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.Metadata, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.MetaCreatedAt):
			key := string(SettlementBatchItemsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.MetaCreatedAt, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.MetaCreatedBy):
			key := string(SettlementBatchItemsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.MetaCreatedBy, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.MetaUpdatedAt):
			key := string(SettlementBatchItemsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.MetaUpdatedAt, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.MetaUpdatedBy):
			key := string(SettlementBatchItemsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.MetaUpdatedBy, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.MetaDeletedAt):
			key := string(SettlementBatchItemsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.MetaDeletedAt.Time, explicitAlias)

		case string(model.SettlementBatchItemsDBFieldName.MetaDeletedBy):
			key := string(SettlementBatchItemsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setSettlementBatchItemsSelectableValue(settlementBatchItemsSelectableResponse, key, settlementBatchItems.MetaDeletedBy, explicitAlias)

		}
	}
	return settlementBatchItemsSelectableResponse
}

func NewSettlementBatchItemsListResponseFromFilterResult(result []model.SettlementBatchItemsFilterResult, filter model.Filter) SettlementBatchItemsSelectableListResponse {
	dtoSettlementBatchItemsListResponse := SettlementBatchItemsSelectableListResponse{}
	for _, row := range result {
		dtoSettlementBatchItemsResponse := NewSettlementBatchItemsSelectableResponse(row.SettlementBatchItems, filter)
		dtoSettlementBatchItemsListResponse = append(dtoSettlementBatchItemsListResponse, &dtoSettlementBatchItemsResponse)
	}
	return dtoSettlementBatchItemsListResponse
}

type SettlementBatchItemsFilterResponse struct {
	Metadata Metadata                                   `json:"metadata"`
	Data     SettlementBatchItemsSelectableListResponse `json:"data"`
}

func reverseSettlementBatchItemsFilterResults(result []model.SettlementBatchItemsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewSettlementBatchItemsFilterResponse(result []model.SettlementBatchItemsFilterResult, filter model.Filter) (resp SettlementBatchItemsFilterResponse) {
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
			reverseSettlementBatchItemsFilterResults(dataResult)
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

	resp.Data = NewSettlementBatchItemsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type SettlementBatchItemsCreateRequest struct {
	SettlementBatchId    uuid.UUID        `json:"settlementBatchId"`
	SourceType           string           `json:"sourceType"`
	SourceId             uuid.UUID        `json:"sourceId"`
	MerchantPartyId      uuid.UUID        `json:"merchantPartyId"`
	CurrencyCode         string           `json:"currencyCode"`
	GrossAmount          decimal.Decimal  `json:"grossAmount"`
	FeeAmount            decimal.Decimal  `json:"feeAmount"`
	TaxAmount            decimal.Decimal  `json:"taxAmount"`
	ReserveAmount        decimal.Decimal  `json:"reserveAmount"`
	NetAmount            decimal.Decimal  `json:"netAmount"`
	LinkedJournalEntryId uuid.UUID        `json:"linkedJournalEntryId"`
	ItemStatus           model.ItemStatus `json:"itemStatus" example:"included" enums:"included,released,reversed"`
	Metadata             json.RawMessage  `json:"metadata"`
}

func (d *SettlementBatchItemsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *SettlementBatchItemsCreateRequest) ToModel() model.SettlementBatchItems {
	id, _ := uuid.NewV4()
	return model.SettlementBatchItems{
		Id:                   id,
		SettlementBatchId:    d.SettlementBatchId,
		SourceType:           d.SourceType,
		SourceId:             d.SourceId,
		MerchantPartyId:      d.MerchantPartyId,
		CurrencyCode:         d.CurrencyCode,
		GrossAmount:          d.GrossAmount,
		FeeAmount:            d.FeeAmount,
		TaxAmount:            d.TaxAmount,
		ReserveAmount:        d.ReserveAmount,
		NetAmount:            d.NetAmount,
		LinkedJournalEntryId: nuuid.From(d.LinkedJournalEntryId),
		ItemStatus:           d.ItemStatus,
		Metadata:             d.Metadata,
	}
}

type SettlementBatchItemsListCreateRequest []*SettlementBatchItemsCreateRequest

func (d SettlementBatchItemsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementBatchItems := range d {
		err = validator.Struct(settlementBatchItems)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementBatchItemsListCreateRequest) ToModelList() []model.SettlementBatchItems {
	out := make([]model.SettlementBatchItems, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type SettlementBatchItemsUpdateRequest struct {
	SettlementBatchId    uuid.UUID        `json:"settlementBatchId"`
	SourceType           string           `json:"sourceType"`
	SourceId             uuid.UUID        `json:"sourceId"`
	MerchantPartyId      uuid.UUID        `json:"merchantPartyId"`
	CurrencyCode         string           `json:"currencyCode"`
	GrossAmount          decimal.Decimal  `json:"grossAmount"`
	FeeAmount            decimal.Decimal  `json:"feeAmount"`
	TaxAmount            decimal.Decimal  `json:"taxAmount"`
	ReserveAmount        decimal.Decimal  `json:"reserveAmount"`
	NetAmount            decimal.Decimal  `json:"netAmount"`
	LinkedJournalEntryId uuid.UUID        `json:"linkedJournalEntryId"`
	ItemStatus           model.ItemStatus `json:"itemStatus" example:"included" enums:"included,released,reversed"`
	Metadata             json.RawMessage  `json:"metadata"`
}

func (d *SettlementBatchItemsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d SettlementBatchItemsUpdateRequest) ToModel() model.SettlementBatchItems {
	return model.SettlementBatchItems{
		SettlementBatchId:    d.SettlementBatchId,
		SourceType:           d.SourceType,
		SourceId:             d.SourceId,
		MerchantPartyId:      d.MerchantPartyId,
		CurrencyCode:         d.CurrencyCode,
		GrossAmount:          d.GrossAmount,
		FeeAmount:            d.FeeAmount,
		TaxAmount:            d.TaxAmount,
		ReserveAmount:        d.ReserveAmount,
		NetAmount:            d.NetAmount,
		LinkedJournalEntryId: nuuid.From(d.LinkedJournalEntryId),
		ItemStatus:           d.ItemStatus,
		Metadata:             d.Metadata,
	}
}

type SettlementBatchItemsBulkUpdateRequest struct {
	Id                   uuid.UUID        `json:"id"`
	SettlementBatchId    uuid.UUID        `json:"settlementBatchId"`
	SourceType           string           `json:"sourceType"`
	SourceId             uuid.UUID        `json:"sourceId"`
	MerchantPartyId      uuid.UUID        `json:"merchantPartyId"`
	CurrencyCode         string           `json:"currencyCode"`
	GrossAmount          decimal.Decimal  `json:"grossAmount"`
	FeeAmount            decimal.Decimal  `json:"feeAmount"`
	TaxAmount            decimal.Decimal  `json:"taxAmount"`
	ReserveAmount        decimal.Decimal  `json:"reserveAmount"`
	NetAmount            decimal.Decimal  `json:"netAmount"`
	LinkedJournalEntryId uuid.UUID        `json:"linkedJournalEntryId"`
	ItemStatus           model.ItemStatus `json:"itemStatus" example:"included" enums:"included,released,reversed"`
	Metadata             json.RawMessage  `json:"metadata"`
}

func (d SettlementBatchItemsBulkUpdateRequest) PrimaryID() SettlementBatchItemsPrimaryID {
	return SettlementBatchItemsPrimaryID{
		Id: d.Id,
	}
}

type SettlementBatchItemsListBulkUpdateRequest []*SettlementBatchItemsBulkUpdateRequest

func (d SettlementBatchItemsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementBatchItems := range d {
		err = validator.Struct(settlementBatchItems)
		if err != nil {
			return
		}
	}
	return nil
}

func (d SettlementBatchItemsBulkUpdateRequest) ToModel() model.SettlementBatchItems {
	return model.SettlementBatchItems{
		Id:                   d.Id,
		SettlementBatchId:    d.SettlementBatchId,
		SourceType:           d.SourceType,
		SourceId:             d.SourceId,
		MerchantPartyId:      d.MerchantPartyId,
		CurrencyCode:         d.CurrencyCode,
		GrossAmount:          d.GrossAmount,
		FeeAmount:            d.FeeAmount,
		TaxAmount:            d.TaxAmount,
		ReserveAmount:        d.ReserveAmount,
		NetAmount:            d.NetAmount,
		LinkedJournalEntryId: nuuid.From(d.LinkedJournalEntryId),
		ItemStatus:           d.ItemStatus,
		Metadata:             d.Metadata,
	}
}

type SettlementBatchItemsResponse struct {
	Id                   uuid.UUID        `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SettlementBatchId    uuid.UUID        `json:"settlementBatchId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceType           string           `json:"sourceType" validate:"required"`
	SourceId             uuid.UUID        `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MerchantPartyId      uuid.UUID        `json:"merchantPartyId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CurrencyCode         string           `json:"currencyCode" validate:"required"`
	GrossAmount          decimal.Decimal  `json:"grossAmount" validate:"required" format:"decimal" example:"100.50"`
	FeeAmount            decimal.Decimal  `json:"feeAmount" format:"decimal" example:"100.50"`
	TaxAmount            decimal.Decimal  `json:"taxAmount" format:"decimal" example:"100.50"`
	ReserveAmount        decimal.Decimal  `json:"reserveAmount" format:"decimal" example:"100.50"`
	NetAmount            decimal.Decimal  `json:"netAmount" validate:"required" format:"decimal" example:"100.50"`
	LinkedJournalEntryId uuid.UUID        `json:"linkedJournalEntryId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ItemStatus           model.ItemStatus `json:"itemStatus" validate:"oneof=included released reversed" enums:"included,released,reversed"`
	Metadata             json.RawMessage  `json:"metadata" swaggertype:"object"`
}

func NewSettlementBatchItemsResponse(settlementBatchItems model.SettlementBatchItems) SettlementBatchItemsResponse {
	return SettlementBatchItemsResponse{
		Id:                   settlementBatchItems.Id,
		SettlementBatchId:    settlementBatchItems.SettlementBatchId,
		SourceType:           settlementBatchItems.SourceType,
		SourceId:             settlementBatchItems.SourceId,
		MerchantPartyId:      settlementBatchItems.MerchantPartyId,
		CurrencyCode:         settlementBatchItems.CurrencyCode,
		GrossAmount:          settlementBatchItems.GrossAmount,
		FeeAmount:            settlementBatchItems.FeeAmount,
		TaxAmount:            settlementBatchItems.TaxAmount,
		ReserveAmount:        settlementBatchItems.ReserveAmount,
		NetAmount:            settlementBatchItems.NetAmount,
		LinkedJournalEntryId: settlementBatchItems.LinkedJournalEntryId.UUID,
		ItemStatus:           model.ItemStatus(settlementBatchItems.ItemStatus),
		Metadata:             settlementBatchItems.Metadata,
	}
}

type SettlementBatchItemsListResponse []*SettlementBatchItemsResponse

func NewSettlementBatchItemsListResponse(settlementBatchItemsList model.SettlementBatchItemsList) SettlementBatchItemsListResponse {
	dtoSettlementBatchItemsListResponse := SettlementBatchItemsListResponse{}
	for _, settlementBatchItems := range settlementBatchItemsList {
		dtoSettlementBatchItemsResponse := NewSettlementBatchItemsResponse(*settlementBatchItems)
		dtoSettlementBatchItemsListResponse = append(dtoSettlementBatchItemsListResponse, &dtoSettlementBatchItemsResponse)
	}
	return dtoSettlementBatchItemsListResponse
}

type SettlementBatchItemsPrimaryIDList []SettlementBatchItemsPrimaryID

func (d SettlementBatchItemsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, settlementBatchItems := range d {
		err = validator.Struct(settlementBatchItems)
		if err != nil {
			return
		}
	}
	return nil
}

type SettlementBatchItemsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *SettlementBatchItemsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d SettlementBatchItemsPrimaryID) ToModel() model.SettlementBatchItemsPrimaryID {
	return model.SettlementBatchItemsPrimaryID{
		Id: d.Id,
	}
}
