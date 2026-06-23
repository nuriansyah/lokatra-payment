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

type ChargebackEventsDTOFieldNameType string

type chargebackEventsDTOFieldName struct {
	Id               ChargebackEventsDTOFieldNameType
	DisputeId        ChargebackEventsDTOFieldNameType
	EventType        ChargebackEventsDTOFieldNameType
	JournalEntryId   ChargebackEventsDTOFieldNameType
	Amount           ChargebackEventsDTOFieldNameType
	CurrencyCode     ChargebackEventsDTOFieldNameType
	ProviderEventRef ChargebackEventsDTOFieldNameType
	EventPayload     ChargebackEventsDTOFieldNameType
	OccurredAt       ChargebackEventsDTOFieldNameType
	Metadata         ChargebackEventsDTOFieldNameType
	MetaCreatedAt    ChargebackEventsDTOFieldNameType
	MetaCreatedBy    ChargebackEventsDTOFieldNameType
	MetaUpdatedAt    ChargebackEventsDTOFieldNameType
	MetaUpdatedBy    ChargebackEventsDTOFieldNameType
	MetaDeletedAt    ChargebackEventsDTOFieldNameType
	MetaDeletedBy    ChargebackEventsDTOFieldNameType
}

var ChargebackEventsDTOFieldName = chargebackEventsDTOFieldName{
	Id:               "id",
	DisputeId:        "disputeId",
	EventType:        "eventType",
	JournalEntryId:   "journalEntryId",
	Amount:           "amount",
	CurrencyCode:     "currencyCode",
	ProviderEventRef: "providerEventRef",
	EventPayload:     "eventPayload",
	OccurredAt:       "occurredAt",
	Metadata:         "metadata",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
	MetaDeletedAt:    "metaDeletedAt",
	MetaDeletedBy:    "metaDeletedBy",
}

func transformChargebackEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ChargebackEventsDTOFieldName.Id):
		return string(model.ChargebackEventsDBFieldName.Id), true

	case string(ChargebackEventsDTOFieldName.DisputeId):
		return string(model.ChargebackEventsDBFieldName.DisputeId), true

	case string(ChargebackEventsDTOFieldName.EventType):
		return string(model.ChargebackEventsDBFieldName.EventType), true

	case string(ChargebackEventsDTOFieldName.JournalEntryId):
		return string(model.ChargebackEventsDBFieldName.JournalEntryId), true

	case string(ChargebackEventsDTOFieldName.Amount):
		return string(model.ChargebackEventsDBFieldName.Amount), true

	case string(ChargebackEventsDTOFieldName.CurrencyCode):
		return string(model.ChargebackEventsDBFieldName.CurrencyCode), true

	case string(ChargebackEventsDTOFieldName.ProviderEventRef):
		return string(model.ChargebackEventsDBFieldName.ProviderEventRef), true

	case string(ChargebackEventsDTOFieldName.EventPayload):
		return string(model.ChargebackEventsDBFieldName.EventPayload), true

	case string(ChargebackEventsDTOFieldName.OccurredAt):
		return string(model.ChargebackEventsDBFieldName.OccurredAt), true

	case string(ChargebackEventsDTOFieldName.Metadata):
		return string(model.ChargebackEventsDBFieldName.Metadata), true

	case string(ChargebackEventsDTOFieldName.MetaCreatedAt):
		return string(model.ChargebackEventsDBFieldName.MetaCreatedAt), true

	case string(ChargebackEventsDTOFieldName.MetaCreatedBy):
		return string(model.ChargebackEventsDBFieldName.MetaCreatedBy), true

	case string(ChargebackEventsDTOFieldName.MetaUpdatedAt):
		return string(model.ChargebackEventsDBFieldName.MetaUpdatedAt), true

	case string(ChargebackEventsDTOFieldName.MetaUpdatedBy):
		return string(model.ChargebackEventsDBFieldName.MetaUpdatedBy), true

	case string(ChargebackEventsDTOFieldName.MetaDeletedAt):
		return string(model.ChargebackEventsDBFieldName.MetaDeletedAt), true

	case string(ChargebackEventsDTOFieldName.MetaDeletedBy):
		return string(model.ChargebackEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewChargebackEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isChargebackEventsBaseFilterField(field string) bool {
	spec, found := model.NewChargebackEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeChargebackEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateChargebackEventsProjectionOutputPath(path string) error {
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

func transformChargebackEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformChargebackEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformChargebackEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformChargebackEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformChargebackEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isChargebackEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateChargebackEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeChargebackEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformChargebackEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformChargebackEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformChargebackEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultChargebackEventsFilter(filter *model.Filter) {
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
			Field: string(ChargebackEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ChargebackEventsSelectableResponse map[string]interface{}
type ChargebackEventsSelectableListResponse []*ChargebackEventsSelectableResponse

func assignChargebackEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setChargebackEventsSelectableValue(out ChargebackEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignChargebackEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewChargebackEventsSelectableResponse(chargebackEvents model.ChargebackEvents, filter model.Filter) ChargebackEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ChargebackEventsDBFieldName.Id),
			string(model.ChargebackEventsDBFieldName.DisputeId),
			string(model.ChargebackEventsDBFieldName.EventType),
			string(model.ChargebackEventsDBFieldName.JournalEntryId),
			string(model.ChargebackEventsDBFieldName.Amount),
			string(model.ChargebackEventsDBFieldName.CurrencyCode),
			string(model.ChargebackEventsDBFieldName.ProviderEventRef),
			string(model.ChargebackEventsDBFieldName.EventPayload),
			string(model.ChargebackEventsDBFieldName.OccurredAt),
			string(model.ChargebackEventsDBFieldName.Metadata),
			string(model.ChargebackEventsDBFieldName.MetaCreatedAt),
			string(model.ChargebackEventsDBFieldName.MetaCreatedBy),
			string(model.ChargebackEventsDBFieldName.MetaUpdatedAt),
			string(model.ChargebackEventsDBFieldName.MetaUpdatedBy),
			string(model.ChargebackEventsDBFieldName.MetaDeletedAt),
			string(model.ChargebackEventsDBFieldName.MetaDeletedBy),
		)
	}
	chargebackEventsSelectableResponse := ChargebackEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ChargebackEventsDBFieldName.Id):
			key := string(ChargebackEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.Id, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.DisputeId):
			key := string(ChargebackEventsDTOFieldName.DisputeId)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.DisputeId, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.EventType):
			key := string(ChargebackEventsDTOFieldName.EventType)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, model.EventType(chargebackEvents.EventType), explicitAlias)

		case string(model.ChargebackEventsDBFieldName.JournalEntryId):
			key := string(ChargebackEventsDTOFieldName.JournalEntryId)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.JournalEntryId.UUID, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.Amount):
			key := string(ChargebackEventsDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.Amount, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.CurrencyCode):
			key := string(ChargebackEventsDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.CurrencyCode, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.ProviderEventRef):
			key := string(ChargebackEventsDTOFieldName.ProviderEventRef)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.ProviderEventRef.String, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.EventPayload):
			key := string(ChargebackEventsDTOFieldName.EventPayload)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.EventPayload, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.OccurredAt):
			key := string(ChargebackEventsDTOFieldName.OccurredAt)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.OccurredAt, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.Metadata):
			key := string(ChargebackEventsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.Metadata, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.MetaCreatedAt):
			key := string(ChargebackEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.MetaCreatedAt, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.MetaCreatedBy):
			key := string(ChargebackEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.MetaCreatedBy, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.MetaUpdatedAt):
			key := string(ChargebackEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.MetaUpdatedAt, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.MetaUpdatedBy):
			key := string(ChargebackEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.MetaUpdatedBy, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.MetaDeletedAt):
			key := string(ChargebackEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.ChargebackEventsDBFieldName.MetaDeletedBy):
			key := string(ChargebackEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setChargebackEventsSelectableValue(chargebackEventsSelectableResponse, key, chargebackEvents.MetaDeletedBy, explicitAlias)

		}
	}
	return chargebackEventsSelectableResponse
}

func NewChargebackEventsListResponseFromFilterResult(result []model.ChargebackEventsFilterResult, filter model.Filter) ChargebackEventsSelectableListResponse {
	dtoChargebackEventsListResponse := ChargebackEventsSelectableListResponse{}
	for _, row := range result {
		dtoChargebackEventsResponse := NewChargebackEventsSelectableResponse(row.ChargebackEvents, filter)
		dtoChargebackEventsListResponse = append(dtoChargebackEventsListResponse, &dtoChargebackEventsResponse)
	}
	return dtoChargebackEventsListResponse
}

type ChargebackEventsFilterResponse struct {
	Metadata Metadata                               `json:"metadata"`
	Data     ChargebackEventsSelectableListResponse `json:"data"`
}

func reverseChargebackEventsFilterResults(result []model.ChargebackEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewChargebackEventsFilterResponse(result []model.ChargebackEventsFilterResult, filter model.Filter) (resp ChargebackEventsFilterResponse) {
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
			reverseChargebackEventsFilterResults(dataResult)
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

	resp.Data = NewChargebackEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ChargebackEventsCreateRequest struct {
	DisputeId        uuid.UUID       `json:"disputeId"`
	EventType        model.EventType `json:"eventType" example:"chargeback_debited" enums:"chargeback_debited,chargeback_reversed,representment_submitted,won,lost"`
	JournalEntryId   uuid.UUID       `json:"journalEntryId"`
	Amount           decimal.Decimal `json:"amount"`
	CurrencyCode     string          `json:"currencyCode"`
	ProviderEventRef string          `json:"providerEventRef"`
	EventPayload     json.RawMessage `json:"eventPayload"`
	OccurredAt       time.Time       `json:"occurredAt"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *ChargebackEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ChargebackEventsCreateRequest) ToModel() model.ChargebackEvents {
	id, _ := uuid.NewV4()
	return model.ChargebackEvents{
		Id:               id,
		DisputeId:        d.DisputeId,
		EventType:        d.EventType,
		JournalEntryId:   nuuid.From(d.JournalEntryId),
		Amount:           d.Amount,
		CurrencyCode:     d.CurrencyCode,
		ProviderEventRef: null.StringFrom(d.ProviderEventRef),
		EventPayload:     d.EventPayload,
		OccurredAt:       d.OccurredAt,
		Metadata:         d.Metadata,
	}
}

type ChargebackEventsListCreateRequest []*ChargebackEventsCreateRequest

func (d ChargebackEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, chargebackEvents := range d {
		err = validator.Struct(chargebackEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ChargebackEventsListCreateRequest) ToModelList() []model.ChargebackEvents {
	out := make([]model.ChargebackEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ChargebackEventsUpdateRequest struct {
	DisputeId        uuid.UUID       `json:"disputeId"`
	EventType        model.EventType `json:"eventType" example:"chargeback_debited" enums:"chargeback_debited,chargeback_reversed,representment_submitted,won,lost"`
	JournalEntryId   uuid.UUID       `json:"journalEntryId"`
	Amount           decimal.Decimal `json:"amount"`
	CurrencyCode     string          `json:"currencyCode"`
	ProviderEventRef string          `json:"providerEventRef"`
	EventPayload     json.RawMessage `json:"eventPayload"`
	OccurredAt       time.Time       `json:"occurredAt"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d *ChargebackEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ChargebackEventsUpdateRequest) ToModel() model.ChargebackEvents {
	return model.ChargebackEvents{
		DisputeId:        d.DisputeId,
		EventType:        d.EventType,
		JournalEntryId:   nuuid.From(d.JournalEntryId),
		Amount:           d.Amount,
		CurrencyCode:     d.CurrencyCode,
		ProviderEventRef: null.StringFrom(d.ProviderEventRef),
		EventPayload:     d.EventPayload,
		OccurredAt:       d.OccurredAt,
		Metadata:         d.Metadata,
	}
}

type ChargebackEventsBulkUpdateRequest struct {
	Id               uuid.UUID       `json:"id"`
	DisputeId        uuid.UUID       `json:"disputeId"`
	EventType        model.EventType `json:"eventType" example:"chargeback_debited" enums:"chargeback_debited,chargeback_reversed,representment_submitted,won,lost"`
	JournalEntryId   uuid.UUID       `json:"journalEntryId"`
	Amount           decimal.Decimal `json:"amount"`
	CurrencyCode     string          `json:"currencyCode"`
	ProviderEventRef string          `json:"providerEventRef"`
	EventPayload     json.RawMessage `json:"eventPayload"`
	OccurredAt       time.Time       `json:"occurredAt"`
	Metadata         json.RawMessage `json:"metadata"`
}

func (d ChargebackEventsBulkUpdateRequest) PrimaryID() ChargebackEventsPrimaryID {
	return ChargebackEventsPrimaryID{
		Id: d.Id,
	}
}

type ChargebackEventsListBulkUpdateRequest []*ChargebackEventsBulkUpdateRequest

func (d ChargebackEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, chargebackEvents := range d {
		err = validator.Struct(chargebackEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ChargebackEventsBulkUpdateRequest) ToModel() model.ChargebackEvents {
	return model.ChargebackEvents{
		Id:               d.Id,
		DisputeId:        d.DisputeId,
		EventType:        d.EventType,
		JournalEntryId:   nuuid.From(d.JournalEntryId),
		Amount:           d.Amount,
		CurrencyCode:     d.CurrencyCode,
		ProviderEventRef: null.StringFrom(d.ProviderEventRef),
		EventPayload:     d.EventPayload,
		OccurredAt:       d.OccurredAt,
		Metadata:         d.Metadata,
	}
}

type ChargebackEventsResponse struct {
	Id               uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	DisputeId        uuid.UUID       `json:"disputeId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	EventType        model.EventType `json:"eventType" validate:"required,oneof=chargeback_debited chargeback_reversed representment_submitted won lost" enums:"chargeback_debited,chargeback_reversed,representment_submitted,won,lost"`
	JournalEntryId   uuid.UUID       `json:"journalEntryId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount           decimal.Decimal `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	CurrencyCode     string          `json:"currencyCode" validate:"required"`
	ProviderEventRef string          `json:"providerEventRef"`
	EventPayload     json.RawMessage `json:"eventPayload" swaggertype:"object"`
	OccurredAt       time.Time       `json:"occurredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata         json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewChargebackEventsResponse(chargebackEvents model.ChargebackEvents) ChargebackEventsResponse {
	return ChargebackEventsResponse{
		Id:               chargebackEvents.Id,
		DisputeId:        chargebackEvents.DisputeId,
		EventType:        model.EventType(chargebackEvents.EventType),
		JournalEntryId:   chargebackEvents.JournalEntryId.UUID,
		Amount:           chargebackEvents.Amount,
		CurrencyCode:     chargebackEvents.CurrencyCode,
		ProviderEventRef: chargebackEvents.ProviderEventRef.String,
		EventPayload:     chargebackEvents.EventPayload,
		OccurredAt:       chargebackEvents.OccurredAt,
		Metadata:         chargebackEvents.Metadata,
	}
}

type ChargebackEventsListResponse []*ChargebackEventsResponse

func NewChargebackEventsListResponse(chargebackEventsList model.ChargebackEventsList) ChargebackEventsListResponse {
	dtoChargebackEventsListResponse := ChargebackEventsListResponse{}
	for _, chargebackEvents := range chargebackEventsList {
		dtoChargebackEventsResponse := NewChargebackEventsResponse(*chargebackEvents)
		dtoChargebackEventsListResponse = append(dtoChargebackEventsListResponse, &dtoChargebackEventsResponse)
	}
	return dtoChargebackEventsListResponse
}

type ChargebackEventsPrimaryIDList []ChargebackEventsPrimaryID

func (d ChargebackEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, chargebackEvents := range d {
		err = validator.Struct(chargebackEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type ChargebackEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ChargebackEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ChargebackEventsPrimaryID) ToModel() model.ChargebackEventsPrimaryID {
	return model.ChargebackEventsPrimaryID{
		Id: d.Id,
	}
}
