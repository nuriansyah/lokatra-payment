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

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type CashCollectionSessionsDTOFieldNameType string

type cashCollectionSessionsDTOFieldName struct {
	Id                 CashCollectionSessionsDTOFieldNameType
	SessionCode        CashCollectionSessionsDTOFieldNameType
	MerchantId         CashCollectionSessionsDTOFieldNameType
	CollectorId        CashCollectionSessionsDTOFieldNameType
	LocationId         CashCollectionSessionsDTOFieldNameType
	OpenedAt           CashCollectionSessionsDTOFieldNameType
	ClosedAt           CashCollectionSessionsDTOFieldNameType
	Status             CashCollectionSessionsDTOFieldNameType
	OpeningFloatAmount CashCollectionSessionsDTOFieldNameType
	ExpectedAmount     CashCollectionSessionsDTOFieldNameType
	CountedAmount      CashCollectionSessionsDTOFieldNameType
	VarianceAmount     CashCollectionSessionsDTOFieldNameType
	Currency           CashCollectionSessionsDTOFieldNameType
	Notes              CashCollectionSessionsDTOFieldNameType
	Metadata           CashCollectionSessionsDTOFieldNameType
	MetaCreatedAt      CashCollectionSessionsDTOFieldNameType
	MetaCreatedBy      CashCollectionSessionsDTOFieldNameType
	MetaUpdatedAt      CashCollectionSessionsDTOFieldNameType
	MetaUpdatedBy      CashCollectionSessionsDTOFieldNameType
	MetaDeletedAt      CashCollectionSessionsDTOFieldNameType
	MetaDeletedBy      CashCollectionSessionsDTOFieldNameType
}

var CashCollectionSessionsDTOFieldName = cashCollectionSessionsDTOFieldName{
	Id:                 "id",
	SessionCode:        "sessionCode",
	MerchantId:         "merchantId",
	CollectorId:        "collectorId",
	LocationId:         "locationId",
	OpenedAt:           "openedAt",
	ClosedAt:           "closedAt",
	Status:             "status",
	OpeningFloatAmount: "openingFloatAmount",
	ExpectedAmount:     "expectedAmount",
	CountedAmount:      "countedAmount",
	VarianceAmount:     "varianceAmount",
	Currency:           "currency",
	Notes:              "notes",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformCashCollectionSessionsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(CashCollectionSessionsDTOFieldName.Id):
		return string(model.CashCollectionSessionsDBFieldName.Id), true

	case string(CashCollectionSessionsDTOFieldName.SessionCode):
		return string(model.CashCollectionSessionsDBFieldName.SessionCode), true

	case string(CashCollectionSessionsDTOFieldName.MerchantId):
		return string(model.CashCollectionSessionsDBFieldName.MerchantId), true

	case string(CashCollectionSessionsDTOFieldName.CollectorId):
		return string(model.CashCollectionSessionsDBFieldName.CollectorId), true

	case string(CashCollectionSessionsDTOFieldName.LocationId):
		return string(model.CashCollectionSessionsDBFieldName.LocationId), true

	case string(CashCollectionSessionsDTOFieldName.OpenedAt):
		return string(model.CashCollectionSessionsDBFieldName.OpenedAt), true

	case string(CashCollectionSessionsDTOFieldName.ClosedAt):
		return string(model.CashCollectionSessionsDBFieldName.ClosedAt), true

	case string(CashCollectionSessionsDTOFieldName.Status):
		return string(model.CashCollectionSessionsDBFieldName.Status), true

	case string(CashCollectionSessionsDTOFieldName.OpeningFloatAmount):
		return string(model.CashCollectionSessionsDBFieldName.OpeningFloatAmount), true

	case string(CashCollectionSessionsDTOFieldName.ExpectedAmount):
		return string(model.CashCollectionSessionsDBFieldName.ExpectedAmount), true

	case string(CashCollectionSessionsDTOFieldName.CountedAmount):
		return string(model.CashCollectionSessionsDBFieldName.CountedAmount), true

	case string(CashCollectionSessionsDTOFieldName.VarianceAmount):
		return string(model.CashCollectionSessionsDBFieldName.VarianceAmount), true

	case string(CashCollectionSessionsDTOFieldName.Currency):
		return string(model.CashCollectionSessionsDBFieldName.Currency), true

	case string(CashCollectionSessionsDTOFieldName.Notes):
		return string(model.CashCollectionSessionsDBFieldName.Notes), true

	case string(CashCollectionSessionsDTOFieldName.Metadata):
		return string(model.CashCollectionSessionsDBFieldName.Metadata), true

	case string(CashCollectionSessionsDTOFieldName.MetaCreatedAt):
		return string(model.CashCollectionSessionsDBFieldName.MetaCreatedAt), true

	case string(CashCollectionSessionsDTOFieldName.MetaCreatedBy):
		return string(model.CashCollectionSessionsDBFieldName.MetaCreatedBy), true

	case string(CashCollectionSessionsDTOFieldName.MetaUpdatedAt):
		return string(model.CashCollectionSessionsDBFieldName.MetaUpdatedAt), true

	case string(CashCollectionSessionsDTOFieldName.MetaUpdatedBy):
		return string(model.CashCollectionSessionsDBFieldName.MetaUpdatedBy), true

	case string(CashCollectionSessionsDTOFieldName.MetaDeletedAt):
		return string(model.CashCollectionSessionsDBFieldName.MetaDeletedAt), true

	case string(CashCollectionSessionsDTOFieldName.MetaDeletedBy):
		return string(model.CashCollectionSessionsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewCashCollectionSessionsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isCashCollectionSessionsBaseFilterField(field string) bool {
	spec, found := model.NewCashCollectionSessionsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeCashCollectionSessionsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateCashCollectionSessionsProjectionOutputPath(path string) error {
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

func transformCashCollectionSessionsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformCashCollectionSessionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformCashCollectionSessionsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformCashCollectionSessionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformCashCollectionSessionsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isCashCollectionSessionsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateCashCollectionSessionsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeCashCollectionSessionsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformCashCollectionSessionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformCashCollectionSessionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformCashCollectionSessionsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultCashCollectionSessionsFilter(filter *model.Filter) {
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
			Field: string(CashCollectionSessionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type CashCollectionSessionsSelectableResponse map[string]interface{}
type CashCollectionSessionsSelectableListResponse []*CashCollectionSessionsSelectableResponse

func assignCashCollectionSessionsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setCashCollectionSessionsSelectableValue(out CashCollectionSessionsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignCashCollectionSessionsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewCashCollectionSessionsSelectableResponse(cashCollectionSessions model.CashCollectionSessions, filter model.Filter) CashCollectionSessionsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.CashCollectionSessionsDBFieldName.Id),
			string(model.CashCollectionSessionsDBFieldName.SessionCode),
			string(model.CashCollectionSessionsDBFieldName.MerchantId),
			string(model.CashCollectionSessionsDBFieldName.CollectorId),
			string(model.CashCollectionSessionsDBFieldName.LocationId),
			string(model.CashCollectionSessionsDBFieldName.OpenedAt),
			string(model.CashCollectionSessionsDBFieldName.ClosedAt),
			string(model.CashCollectionSessionsDBFieldName.Status),
			string(model.CashCollectionSessionsDBFieldName.OpeningFloatAmount),
			string(model.CashCollectionSessionsDBFieldName.ExpectedAmount),
			string(model.CashCollectionSessionsDBFieldName.CountedAmount),
			string(model.CashCollectionSessionsDBFieldName.VarianceAmount),
			string(model.CashCollectionSessionsDBFieldName.Currency),
			string(model.CashCollectionSessionsDBFieldName.Notes),
			string(model.CashCollectionSessionsDBFieldName.Metadata),
			string(model.CashCollectionSessionsDBFieldName.MetaCreatedAt),
			string(model.CashCollectionSessionsDBFieldName.MetaCreatedBy),
			string(model.CashCollectionSessionsDBFieldName.MetaUpdatedAt),
			string(model.CashCollectionSessionsDBFieldName.MetaUpdatedBy),
			string(model.CashCollectionSessionsDBFieldName.MetaDeletedAt),
			string(model.CashCollectionSessionsDBFieldName.MetaDeletedBy),
		)
	}
	cashCollectionSessionsSelectableResponse := CashCollectionSessionsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.CashCollectionSessionsDBFieldName.Id):
			key := string(CashCollectionSessionsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.Id, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.SessionCode):
			key := string(CashCollectionSessionsDTOFieldName.SessionCode)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.SessionCode, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.MerchantId):
			key := string(CashCollectionSessionsDTOFieldName.MerchantId)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.MerchantId, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.CollectorId):
			key := string(CashCollectionSessionsDTOFieldName.CollectorId)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.CollectorId, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.LocationId):
			key := string(CashCollectionSessionsDTOFieldName.LocationId)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.LocationId.UUID, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.OpenedAt):
			key := string(CashCollectionSessionsDTOFieldName.OpenedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.OpenedAt, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.ClosedAt):
			key := string(CashCollectionSessionsDTOFieldName.ClosedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.ClosedAt.Time, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.Status):
			key := string(CashCollectionSessionsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, model.CashSessionStatus(cashCollectionSessions.Status), explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.OpeningFloatAmount):
			key := string(CashCollectionSessionsDTOFieldName.OpeningFloatAmount)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.OpeningFloatAmount, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.ExpectedAmount):
			key := string(CashCollectionSessionsDTOFieldName.ExpectedAmount)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.ExpectedAmount, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.CountedAmount):
			key := string(CashCollectionSessionsDTOFieldName.CountedAmount)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.CountedAmount, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.VarianceAmount):
			key := string(CashCollectionSessionsDTOFieldName.VarianceAmount)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.VarianceAmount, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.Currency):
			key := string(CashCollectionSessionsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.Currency, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.Notes):
			key := string(CashCollectionSessionsDTOFieldName.Notes)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.Notes.String, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.Metadata):
			key := string(CashCollectionSessionsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.Metadata, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.MetaCreatedAt):
			key := string(CashCollectionSessionsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.MetaCreatedAt, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.MetaCreatedBy):
			key := string(CashCollectionSessionsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.MetaCreatedBy, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.MetaUpdatedAt):
			key := string(CashCollectionSessionsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.MetaUpdatedAt, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.MetaUpdatedBy):
			key := string(CashCollectionSessionsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.MetaDeletedAt):
			key := string(CashCollectionSessionsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.MetaDeletedAt.Time, explicitAlias)

		case string(model.CashCollectionSessionsDBFieldName.MetaDeletedBy):
			key := string(CashCollectionSessionsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionSessionsSelectableValue(cashCollectionSessionsSelectableResponse, key, cashCollectionSessions.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return cashCollectionSessionsSelectableResponse
}

func NewCashCollectionSessionsListResponseFromFilterResult(result []model.CashCollectionSessionsFilterResult, filter model.Filter) CashCollectionSessionsSelectableListResponse {
	dtoCashCollectionSessionsListResponse := CashCollectionSessionsSelectableListResponse{}
	for _, row := range result {
		dtoCashCollectionSessionsResponse := NewCashCollectionSessionsSelectableResponse(row.CashCollectionSessions, filter)
		dtoCashCollectionSessionsListResponse = append(dtoCashCollectionSessionsListResponse, &dtoCashCollectionSessionsResponse)
	}
	return dtoCashCollectionSessionsListResponse
}

type CashCollectionSessionsFilterResponse struct {
	Metadata Metadata                                     `json:"metadata"`
	Data     CashCollectionSessionsSelectableListResponse `json:"data"`
}

func reverseCashCollectionSessionsFilterResults(result []model.CashCollectionSessionsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewCashCollectionSessionsFilterResponse(result []model.CashCollectionSessionsFilterResult, filter model.Filter) (resp CashCollectionSessionsFilterResponse) {
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
			reverseCashCollectionSessionsFilterResults(dataResult)
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

	resp.Data = NewCashCollectionSessionsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type CashCollectionSessionsCreateRequest struct {
	SessionCode        string                  `json:"sessionCode"`
	MerchantId         uuid.UUID               `json:"merchantId"`
	CollectorId        uuid.UUID               `json:"collectorId"`
	LocationId         uuid.UUID               `json:"locationId"`
	OpenedAt           time.Time               `json:"openedAt"`
	ClosedAt           time.Time               `json:"closedAt"`
	Status             model.CashSessionStatus `json:"status" example:"open" enums:"open,closed,canceled"`
	OpeningFloatAmount decimal.Decimal         `json:"openingFloatAmount"`
	ExpectedAmount     decimal.Decimal         `json:"expectedAmount"`
	CountedAmount      decimal.Decimal         `json:"countedAmount"`
	VarianceAmount     decimal.Decimal         `json:"varianceAmount"`
	Currency           string                  `json:"currency"`
	Notes              string                  `json:"notes"`
	Metadata           json.RawMessage         `json:"metadata"`
}

func (d *CashCollectionSessionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *CashCollectionSessionsCreateRequest) ToModel() model.CashCollectionSessions {
	id, _ := uuid.NewV4()
	return model.CashCollectionSessions{
		Id:                 id,
		SessionCode:        d.SessionCode,
		MerchantId:         d.MerchantId,
		CollectorId:        d.CollectorId,
		LocationId:         nuuid.From(d.LocationId),
		OpenedAt:           d.OpenedAt,
		ClosedAt:           null.TimeFrom(d.ClosedAt),
		Status:             d.Status,
		OpeningFloatAmount: d.OpeningFloatAmount,
		ExpectedAmount:     d.ExpectedAmount,
		CountedAmount:      d.CountedAmount,
		VarianceAmount:     d.VarianceAmount,
		Currency:           d.Currency,
		Notes:              null.StringFrom(d.Notes),
		Metadata:           d.Metadata,
	}
}

type CashCollectionSessionsListCreateRequest []*CashCollectionSessionsCreateRequest

func (d CashCollectionSessionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, cashCollectionSessions := range d {
		err = validator.Struct(cashCollectionSessions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d CashCollectionSessionsListCreateRequest) ToModelList() []model.CashCollectionSessions {
	out := make([]model.CashCollectionSessions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type CashCollectionSessionsUpdateRequest struct {
	SessionCode        string                  `json:"sessionCode"`
	MerchantId         uuid.UUID               `json:"merchantId"`
	CollectorId        uuid.UUID               `json:"collectorId"`
	LocationId         uuid.UUID               `json:"locationId"`
	OpenedAt           time.Time               `json:"openedAt"`
	ClosedAt           time.Time               `json:"closedAt"`
	Status             model.CashSessionStatus `json:"status" example:"open" enums:"open,closed,canceled"`
	OpeningFloatAmount decimal.Decimal         `json:"openingFloatAmount"`
	ExpectedAmount     decimal.Decimal         `json:"expectedAmount"`
	CountedAmount      decimal.Decimal         `json:"countedAmount"`
	VarianceAmount     decimal.Decimal         `json:"varianceAmount"`
	Currency           string                  `json:"currency"`
	Notes              string                  `json:"notes"`
	Metadata           json.RawMessage         `json:"metadata"`
}

func (d *CashCollectionSessionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d CashCollectionSessionsUpdateRequest) ToModel() model.CashCollectionSessions {
	return model.CashCollectionSessions{
		SessionCode:        d.SessionCode,
		MerchantId:         d.MerchantId,
		CollectorId:        d.CollectorId,
		LocationId:         nuuid.From(d.LocationId),
		OpenedAt:           d.OpenedAt,
		ClosedAt:           null.TimeFrom(d.ClosedAt),
		Status:             d.Status,
		OpeningFloatAmount: d.OpeningFloatAmount,
		ExpectedAmount:     d.ExpectedAmount,
		CountedAmount:      d.CountedAmount,
		VarianceAmount:     d.VarianceAmount,
		Currency:           d.Currency,
		Notes:              null.StringFrom(d.Notes),
		Metadata:           d.Metadata,
	}
}

type CashCollectionSessionsBulkUpdateRequest struct {
	Id                 uuid.UUID               `json:"id"`
	SessionCode        string                  `json:"sessionCode"`
	MerchantId         uuid.UUID               `json:"merchantId"`
	CollectorId        uuid.UUID               `json:"collectorId"`
	LocationId         uuid.UUID               `json:"locationId"`
	OpenedAt           time.Time               `json:"openedAt"`
	ClosedAt           time.Time               `json:"closedAt"`
	Status             model.CashSessionStatus `json:"status" example:"open" enums:"open,closed,canceled"`
	OpeningFloatAmount decimal.Decimal         `json:"openingFloatAmount"`
	ExpectedAmount     decimal.Decimal         `json:"expectedAmount"`
	CountedAmount      decimal.Decimal         `json:"countedAmount"`
	VarianceAmount     decimal.Decimal         `json:"varianceAmount"`
	Currency           string                  `json:"currency"`
	Notes              string                  `json:"notes"`
	Metadata           json.RawMessage         `json:"metadata"`
}

func (d CashCollectionSessionsBulkUpdateRequest) PrimaryID() CashCollectionSessionsPrimaryID {
	return CashCollectionSessionsPrimaryID{
		Id: d.Id,
	}
}

type CashCollectionSessionsListBulkUpdateRequest []*CashCollectionSessionsBulkUpdateRequest

func (d CashCollectionSessionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, cashCollectionSessions := range d {
		err = validator.Struct(cashCollectionSessions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d CashCollectionSessionsBulkUpdateRequest) ToModel() model.CashCollectionSessions {
	return model.CashCollectionSessions{
		Id:                 d.Id,
		SessionCode:        d.SessionCode,
		MerchantId:         d.MerchantId,
		CollectorId:        d.CollectorId,
		LocationId:         nuuid.From(d.LocationId),
		OpenedAt:           d.OpenedAt,
		ClosedAt:           null.TimeFrom(d.ClosedAt),
		Status:             d.Status,
		OpeningFloatAmount: d.OpeningFloatAmount,
		ExpectedAmount:     d.ExpectedAmount,
		CountedAmount:      d.CountedAmount,
		VarianceAmount:     d.VarianceAmount,
		Currency:           d.Currency,
		Notes:              null.StringFrom(d.Notes),
		Metadata:           d.Metadata,
	}
}

type CashCollectionSessionsResponse struct {
	Id                 uuid.UUID               `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SessionCode        string                  `json:"sessionCode" validate:"required"`
	MerchantId         uuid.UUID               `json:"merchantId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CollectorId        uuid.UUID               `json:"collectorId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	LocationId         uuid.UUID               `json:"locationId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OpenedAt           time.Time               `json:"openedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ClosedAt           time.Time               `json:"closedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Status             model.CashSessionStatus `json:"status" validate:"oneof=open closed canceled" enums:"open,closed,canceled"`
	OpeningFloatAmount decimal.Decimal         `json:"openingFloatAmount" format:"decimal" example:"100.50"`
	ExpectedAmount     decimal.Decimal         `json:"expectedAmount" format:"decimal" example:"100.50"`
	CountedAmount      decimal.Decimal         `json:"countedAmount" format:"decimal" example:"100.50"`
	VarianceAmount     decimal.Decimal         `json:"varianceAmount" format:"decimal" example:"100.50"`
	Currency           string                  `json:"currency"`
	Notes              string                  `json:"notes"`
	Metadata           json.RawMessage         `json:"metadata" swaggertype:"object"`
}

func NewCashCollectionSessionsResponse(cashCollectionSessions model.CashCollectionSessions) CashCollectionSessionsResponse {
	return CashCollectionSessionsResponse{
		Id:                 cashCollectionSessions.Id,
		SessionCode:        cashCollectionSessions.SessionCode,
		MerchantId:         cashCollectionSessions.MerchantId,
		CollectorId:        cashCollectionSessions.CollectorId,
		LocationId:         cashCollectionSessions.LocationId.UUID,
		OpenedAt:           cashCollectionSessions.OpenedAt,
		ClosedAt:           cashCollectionSessions.ClosedAt.Time,
		Status:             model.CashSessionStatus(cashCollectionSessions.Status),
		OpeningFloatAmount: cashCollectionSessions.OpeningFloatAmount,
		ExpectedAmount:     cashCollectionSessions.ExpectedAmount,
		CountedAmount:      cashCollectionSessions.CountedAmount,
		VarianceAmount:     cashCollectionSessions.VarianceAmount,
		Currency:           cashCollectionSessions.Currency,
		Notes:              cashCollectionSessions.Notes.String,
		Metadata:           cashCollectionSessions.Metadata,
	}
}

type CashCollectionSessionsListResponse []*CashCollectionSessionsResponse

func NewCashCollectionSessionsListResponse(cashCollectionSessionsList model.CashCollectionSessionsList) CashCollectionSessionsListResponse {
	dtoCashCollectionSessionsListResponse := CashCollectionSessionsListResponse{}
	for _, cashCollectionSessions := range cashCollectionSessionsList {
		dtoCashCollectionSessionsResponse := NewCashCollectionSessionsResponse(*cashCollectionSessions)
		dtoCashCollectionSessionsListResponse = append(dtoCashCollectionSessionsListResponse, &dtoCashCollectionSessionsResponse)
	}
	return dtoCashCollectionSessionsListResponse
}

type CashCollectionSessionsPrimaryIDList []CashCollectionSessionsPrimaryID

func (d CashCollectionSessionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, cashCollectionSessions := range d {
		err = validator.Struct(cashCollectionSessions)
		if err != nil {
			return
		}
	}
	return nil
}

type CashCollectionSessionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *CashCollectionSessionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d CashCollectionSessionsPrimaryID) ToModel() model.CashCollectionSessionsPrimaryID {
	return model.CashCollectionSessionsPrimaryID{
		Id: d.Id,
	}
}
