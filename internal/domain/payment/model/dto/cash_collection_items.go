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

type CashCollectionItemsDTOFieldNameType string

type cashCollectionItemsDTOFieldName struct {
	Id                      CashCollectionItemsDTOFieldNameType
	CashCollectionSessionId CashCollectionItemsDTOFieldNameType
	PaymentIntentId         CashCollectionItemsDTOFieldNameType
	PaymentAttemptId        CashCollectionItemsDTOFieldNameType
	CollectionType          CashCollectionItemsDTOFieldNameType
	Amount                  CashCollectionItemsDTOFieldNameType
	Currency                CashCollectionItemsDTOFieldNameType
	Status                  CashCollectionItemsDTOFieldNameType
	CollectedAt             CashCollectionItemsDTOFieldNameType
	VoidedAt                CashCollectionItemsDTOFieldNameType
	VoidReason              CashCollectionItemsDTOFieldNameType
	Notes                   CashCollectionItemsDTOFieldNameType
	Metadata                CashCollectionItemsDTOFieldNameType
	MetaCreatedAt           CashCollectionItemsDTOFieldNameType
	MetaCreatedBy           CashCollectionItemsDTOFieldNameType
	MetaUpdatedAt           CashCollectionItemsDTOFieldNameType
	MetaUpdatedBy           CashCollectionItemsDTOFieldNameType
	MetaDeletedAt           CashCollectionItemsDTOFieldNameType
	MetaDeletedBy           CashCollectionItemsDTOFieldNameType
}

var CashCollectionItemsDTOFieldName = cashCollectionItemsDTOFieldName{
	Id:                      "id",
	CashCollectionSessionId: "cashCollectionSessionId",
	PaymentIntentId:         "paymentIntentId",
	PaymentAttemptId:        "paymentAttemptId",
	CollectionType:          "collectionType",
	Amount:                  "amount",
	Currency:                "currency",
	Status:                  "status",
	CollectedAt:             "collectedAt",
	VoidedAt:                "voidedAt",
	VoidReason:              "voidReason",
	Notes:                   "notes",
	Metadata:                "metadata",
	MetaCreatedAt:           "metaCreatedAt",
	MetaCreatedBy:           "metaCreatedBy",
	MetaUpdatedAt:           "metaUpdatedAt",
	MetaUpdatedBy:           "metaUpdatedBy",
	MetaDeletedAt:           "metaDeletedAt",
	MetaDeletedBy:           "metaDeletedBy",
}

func transformCashCollectionItemsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(CashCollectionItemsDTOFieldName.Id):
		return string(model.CashCollectionItemsDBFieldName.Id), true

	case string(CashCollectionItemsDTOFieldName.CashCollectionSessionId):
		return string(model.CashCollectionItemsDBFieldName.CashCollectionSessionId), true

	case string(CashCollectionItemsDTOFieldName.PaymentIntentId):
		return string(model.CashCollectionItemsDBFieldName.PaymentIntentId), true

	case string(CashCollectionItemsDTOFieldName.PaymentAttemptId):
		return string(model.CashCollectionItemsDBFieldName.PaymentAttemptId), true

	case string(CashCollectionItemsDTOFieldName.CollectionType):
		return string(model.CashCollectionItemsDBFieldName.CollectionType), true

	case string(CashCollectionItemsDTOFieldName.Amount):
		return string(model.CashCollectionItemsDBFieldName.Amount), true

	case string(CashCollectionItemsDTOFieldName.Currency):
		return string(model.CashCollectionItemsDBFieldName.Currency), true

	case string(CashCollectionItemsDTOFieldName.Status):
		return string(model.CashCollectionItemsDBFieldName.Status), true

	case string(CashCollectionItemsDTOFieldName.CollectedAt):
		return string(model.CashCollectionItemsDBFieldName.CollectedAt), true

	case string(CashCollectionItemsDTOFieldName.VoidedAt):
		return string(model.CashCollectionItemsDBFieldName.VoidedAt), true

	case string(CashCollectionItemsDTOFieldName.VoidReason):
		return string(model.CashCollectionItemsDBFieldName.VoidReason), true

	case string(CashCollectionItemsDTOFieldName.Notes):
		return string(model.CashCollectionItemsDBFieldName.Notes), true

	case string(CashCollectionItemsDTOFieldName.Metadata):
		return string(model.CashCollectionItemsDBFieldName.Metadata), true

	case string(CashCollectionItemsDTOFieldName.MetaCreatedAt):
		return string(model.CashCollectionItemsDBFieldName.MetaCreatedAt), true

	case string(CashCollectionItemsDTOFieldName.MetaCreatedBy):
		return string(model.CashCollectionItemsDBFieldName.MetaCreatedBy), true

	case string(CashCollectionItemsDTOFieldName.MetaUpdatedAt):
		return string(model.CashCollectionItemsDBFieldName.MetaUpdatedAt), true

	case string(CashCollectionItemsDTOFieldName.MetaUpdatedBy):
		return string(model.CashCollectionItemsDBFieldName.MetaUpdatedBy), true

	case string(CashCollectionItemsDTOFieldName.MetaDeletedAt):
		return string(model.CashCollectionItemsDBFieldName.MetaDeletedAt), true

	case string(CashCollectionItemsDTOFieldName.MetaDeletedBy):
		return string(model.CashCollectionItemsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewCashCollectionItemsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isCashCollectionItemsBaseFilterField(field string) bool {
	spec, found := model.NewCashCollectionItemsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeCashCollectionItemsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateCashCollectionItemsProjectionOutputPath(path string) error {
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

func transformCashCollectionItemsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformCashCollectionItemsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformCashCollectionItemsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformCashCollectionItemsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformCashCollectionItemsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isCashCollectionItemsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateCashCollectionItemsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeCashCollectionItemsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformCashCollectionItemsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformCashCollectionItemsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformCashCollectionItemsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultCashCollectionItemsFilter(filter *model.Filter) {
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
			Field: string(CashCollectionItemsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type CashCollectionItemsSelectableResponse map[string]interface{}
type CashCollectionItemsSelectableListResponse []*CashCollectionItemsSelectableResponse

func assignCashCollectionItemsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setCashCollectionItemsSelectableValue(out CashCollectionItemsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignCashCollectionItemsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewCashCollectionItemsSelectableResponse(cashCollectionItems model.CashCollectionItems, filter model.Filter) CashCollectionItemsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.CashCollectionItemsDBFieldName.Id),
			string(model.CashCollectionItemsDBFieldName.CashCollectionSessionId),
			string(model.CashCollectionItemsDBFieldName.PaymentIntentId),
			string(model.CashCollectionItemsDBFieldName.PaymentAttemptId),
			string(model.CashCollectionItemsDBFieldName.CollectionType),
			string(model.CashCollectionItemsDBFieldName.Amount),
			string(model.CashCollectionItemsDBFieldName.Currency),
			string(model.CashCollectionItemsDBFieldName.Status),
			string(model.CashCollectionItemsDBFieldName.CollectedAt),
			string(model.CashCollectionItemsDBFieldName.VoidedAt),
			string(model.CashCollectionItemsDBFieldName.VoidReason),
			string(model.CashCollectionItemsDBFieldName.Notes),
			string(model.CashCollectionItemsDBFieldName.Metadata),
			string(model.CashCollectionItemsDBFieldName.MetaCreatedAt),
			string(model.CashCollectionItemsDBFieldName.MetaCreatedBy),
			string(model.CashCollectionItemsDBFieldName.MetaUpdatedAt),
			string(model.CashCollectionItemsDBFieldName.MetaUpdatedBy),
			string(model.CashCollectionItemsDBFieldName.MetaDeletedAt),
			string(model.CashCollectionItemsDBFieldName.MetaDeletedBy),
		)
	}
	cashCollectionItemsSelectableResponse := CashCollectionItemsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.CashCollectionItemsDBFieldName.Id):
			key := string(CashCollectionItemsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.Id, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.CashCollectionSessionId):
			key := string(CashCollectionItemsDTOFieldName.CashCollectionSessionId)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.CashCollectionSessionId, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.PaymentIntentId):
			key := string(CashCollectionItemsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.PaymentIntentId, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.PaymentAttemptId):
			key := string(CashCollectionItemsDTOFieldName.PaymentAttemptId)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.PaymentAttemptId.UUID, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.CollectionType):
			key := string(CashCollectionItemsDTOFieldName.CollectionType)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.CollectionType, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.Amount):
			key := string(CashCollectionItemsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.Amount, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.Currency):
			key := string(CashCollectionItemsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.Currency, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.Status):
			key := string(CashCollectionItemsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, model.CashItemStatus(cashCollectionItems.Status), explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.CollectedAt):
			key := string(CashCollectionItemsDTOFieldName.CollectedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.CollectedAt, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.VoidedAt):
			key := string(CashCollectionItemsDTOFieldName.VoidedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.VoidedAt.Time, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.VoidReason):
			key := string(CashCollectionItemsDTOFieldName.VoidReason)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.VoidReason.String, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.Notes):
			key := string(CashCollectionItemsDTOFieldName.Notes)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.Notes.String, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.Metadata):
			key := string(CashCollectionItemsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.Metadata, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.MetaCreatedAt):
			key := string(CashCollectionItemsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.MetaCreatedAt, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.MetaCreatedBy):
			key := string(CashCollectionItemsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.MetaCreatedBy, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.MetaUpdatedAt):
			key := string(CashCollectionItemsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.MetaUpdatedAt, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.MetaUpdatedBy):
			key := string(CashCollectionItemsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.MetaUpdatedBy, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.MetaDeletedAt):
			key := string(CashCollectionItemsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.MetaDeletedAt.Time, explicitAlias)

		case string(model.CashCollectionItemsDBFieldName.MetaDeletedBy):
			key := string(CashCollectionItemsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setCashCollectionItemsSelectableValue(cashCollectionItemsSelectableResponse, key, cashCollectionItems.MetaDeletedBy, explicitAlias)

		}
	}
	return cashCollectionItemsSelectableResponse
}

func NewCashCollectionItemsListResponseFromFilterResult(result []model.CashCollectionItemsFilterResult, filter model.Filter) CashCollectionItemsSelectableListResponse {
	dtoCashCollectionItemsListResponse := CashCollectionItemsSelectableListResponse{}
	for _, row := range result {
		dtoCashCollectionItemsResponse := NewCashCollectionItemsSelectableResponse(row.CashCollectionItems, filter)
		dtoCashCollectionItemsListResponse = append(dtoCashCollectionItemsListResponse, &dtoCashCollectionItemsResponse)
	}
	return dtoCashCollectionItemsListResponse
}

type CashCollectionItemsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     CashCollectionItemsSelectableListResponse `json:"data"`
}

func reverseCashCollectionItemsFilterResults(result []model.CashCollectionItemsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewCashCollectionItemsFilterResponse(result []model.CashCollectionItemsFilterResult, filter model.Filter) (resp CashCollectionItemsFilterResponse) {
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
			reverseCashCollectionItemsFilterResults(dataResult)
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

	resp.Data = NewCashCollectionItemsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type CashCollectionItemsCreateRequest struct {
	CashCollectionSessionId uuid.UUID            `json:"cashCollectionSessionId"`
	PaymentIntentId         uuid.UUID            `json:"paymentIntentId"`
	PaymentAttemptId        uuid.UUID            `json:"paymentAttemptId"`
	CollectionType          string               `json:"collectionType"`
	Amount                  decimal.Decimal      `json:"amount"`
	Currency                string               `json:"currency"`
	Status                  model.CashItemStatus `json:"status" example:"collected" enums:"collected,voided"`
	CollectedAt             time.Time            `json:"collectedAt"`
	VoidedAt                time.Time            `json:"voidedAt"`
	VoidReason              string               `json:"voidReason"`
	Notes                   string               `json:"notes"`
	Metadata                json.RawMessage      `json:"metadata"`
}

func (d *CashCollectionItemsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *CashCollectionItemsCreateRequest) ToModel() model.CashCollectionItems {
	id, _ := uuid.NewV4()
	return model.CashCollectionItems{
		Id:                      id,
		CashCollectionSessionId: d.CashCollectionSessionId,
		PaymentIntentId:         d.PaymentIntentId,
		PaymentAttemptId:        nuuid.From(d.PaymentAttemptId),
		CollectionType:          d.CollectionType,
		Amount:                  d.Amount,
		Currency:                d.Currency,
		Status:                  d.Status,
		CollectedAt:             d.CollectedAt,
		VoidedAt:                null.TimeFrom(d.VoidedAt),
		VoidReason:              null.StringFrom(d.VoidReason),
		Notes:                   null.StringFrom(d.Notes),
		Metadata:                d.Metadata,
	}
}

type CashCollectionItemsListCreateRequest []*CashCollectionItemsCreateRequest

func (d CashCollectionItemsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, cashCollectionItems := range d {
		err = validator.Struct(cashCollectionItems)
		if err != nil {
			return
		}
	}
	return nil
}

func (d CashCollectionItemsListCreateRequest) ToModelList() []model.CashCollectionItems {
	out := make([]model.CashCollectionItems, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type CashCollectionItemsUpdateRequest struct {
	CashCollectionSessionId uuid.UUID            `json:"cashCollectionSessionId"`
	PaymentIntentId         uuid.UUID            `json:"paymentIntentId"`
	PaymentAttemptId        uuid.UUID            `json:"paymentAttemptId"`
	CollectionType          string               `json:"collectionType"`
	Amount                  decimal.Decimal      `json:"amount"`
	Currency                string               `json:"currency"`
	Status                  model.CashItemStatus `json:"status" example:"collected" enums:"collected,voided"`
	CollectedAt             time.Time            `json:"collectedAt"`
	VoidedAt                time.Time            `json:"voidedAt"`
	VoidReason              string               `json:"voidReason"`
	Notes                   string               `json:"notes"`
	Metadata                json.RawMessage      `json:"metadata"`
}

func (d *CashCollectionItemsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d CashCollectionItemsUpdateRequest) ToModel() model.CashCollectionItems {
	return model.CashCollectionItems{
		CashCollectionSessionId: d.CashCollectionSessionId,
		PaymentIntentId:         d.PaymentIntentId,
		PaymentAttemptId:        nuuid.From(d.PaymentAttemptId),
		CollectionType:          d.CollectionType,
		Amount:                  d.Amount,
		Currency:                d.Currency,
		Status:                  d.Status,
		CollectedAt:             d.CollectedAt,
		VoidedAt:                null.TimeFrom(d.VoidedAt),
		VoidReason:              null.StringFrom(d.VoidReason),
		Notes:                   null.StringFrom(d.Notes),
		Metadata:                d.Metadata,
	}
}

type CashCollectionItemsBulkUpdateRequest struct {
	Id                      uuid.UUID            `json:"id"`
	CashCollectionSessionId uuid.UUID            `json:"cashCollectionSessionId"`
	PaymentIntentId         uuid.UUID            `json:"paymentIntentId"`
	PaymentAttemptId        uuid.UUID            `json:"paymentAttemptId"`
	CollectionType          string               `json:"collectionType"`
	Amount                  decimal.Decimal      `json:"amount"`
	Currency                string               `json:"currency"`
	Status                  model.CashItemStatus `json:"status" example:"collected" enums:"collected,voided"`
	CollectedAt             time.Time            `json:"collectedAt"`
	VoidedAt                time.Time            `json:"voidedAt"`
	VoidReason              string               `json:"voidReason"`
	Notes                   string               `json:"notes"`
	Metadata                json.RawMessage      `json:"metadata"`
}

func (d CashCollectionItemsBulkUpdateRequest) PrimaryID() CashCollectionItemsPrimaryID {
	return CashCollectionItemsPrimaryID{
		Id: d.Id,
	}
}

type CashCollectionItemsListBulkUpdateRequest []*CashCollectionItemsBulkUpdateRequest

func (d CashCollectionItemsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, cashCollectionItems := range d {
		err = validator.Struct(cashCollectionItems)
		if err != nil {
			return
		}
	}
	return nil
}

func (d CashCollectionItemsBulkUpdateRequest) ToModel() model.CashCollectionItems {
	return model.CashCollectionItems{
		Id:                      d.Id,
		CashCollectionSessionId: d.CashCollectionSessionId,
		PaymentIntentId:         d.PaymentIntentId,
		PaymentAttemptId:        nuuid.From(d.PaymentAttemptId),
		CollectionType:          d.CollectionType,
		Amount:                  d.Amount,
		Currency:                d.Currency,
		Status:                  d.Status,
		CollectedAt:             d.CollectedAt,
		VoidedAt:                null.TimeFrom(d.VoidedAt),
		VoidReason:              null.StringFrom(d.VoidReason),
		Notes:                   null.StringFrom(d.Notes),
		Metadata:                d.Metadata,
	}
}

type CashCollectionItemsResponse struct {
	Id                      uuid.UUID            `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CashCollectionSessionId uuid.UUID            `json:"cashCollectionSessionId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId         uuid.UUID            `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAttemptId        uuid.UUID            `json:"paymentAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CollectionType          string               `json:"collectionType" validate:"required"`
	Amount                  decimal.Decimal      `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency                string               `json:"currency"`
	Status                  model.CashItemStatus `json:"status" validate:"oneof=collected voided" enums:"collected,voided"`
	CollectedAt             time.Time            `json:"collectedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	VoidedAt                time.Time            `json:"voidedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	VoidReason              string               `json:"voidReason"`
	Notes                   string               `json:"notes"`
	Metadata                json.RawMessage      `json:"metadata" swaggertype:"object"`
}

func NewCashCollectionItemsResponse(cashCollectionItems model.CashCollectionItems) CashCollectionItemsResponse {
	return CashCollectionItemsResponse{
		Id:                      cashCollectionItems.Id,
		CashCollectionSessionId: cashCollectionItems.CashCollectionSessionId,
		PaymentIntentId:         cashCollectionItems.PaymentIntentId,
		PaymentAttemptId:        cashCollectionItems.PaymentAttemptId.UUID,
		CollectionType:          cashCollectionItems.CollectionType,
		Amount:                  cashCollectionItems.Amount,
		Currency:                cashCollectionItems.Currency,
		Status:                  model.CashItemStatus(cashCollectionItems.Status),
		CollectedAt:             cashCollectionItems.CollectedAt,
		VoidedAt:                cashCollectionItems.VoidedAt.Time,
		VoidReason:              cashCollectionItems.VoidReason.String,
		Notes:                   cashCollectionItems.Notes.String,
		Metadata:                cashCollectionItems.Metadata,
	}
}

type CashCollectionItemsListResponse []*CashCollectionItemsResponse

func NewCashCollectionItemsListResponse(cashCollectionItemsList model.CashCollectionItemsList) CashCollectionItemsListResponse {
	dtoCashCollectionItemsListResponse := CashCollectionItemsListResponse{}
	for _, cashCollectionItems := range cashCollectionItemsList {
		dtoCashCollectionItemsResponse := NewCashCollectionItemsResponse(*cashCollectionItems)
		dtoCashCollectionItemsListResponse = append(dtoCashCollectionItemsListResponse, &dtoCashCollectionItemsResponse)
	}
	return dtoCashCollectionItemsListResponse
}

type CashCollectionItemsPrimaryIDList []CashCollectionItemsPrimaryID

func (d CashCollectionItemsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, cashCollectionItems := range d {
		err = validator.Struct(cashCollectionItems)
		if err != nil {
			return
		}
	}
	return nil
}

type CashCollectionItemsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *CashCollectionItemsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d CashCollectionItemsPrimaryID) ToModel() model.CashCollectionItemsPrimaryID {
	return model.CashCollectionItemsPrimaryID{
		Id: d.Id,
	}
}
