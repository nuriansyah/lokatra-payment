package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PayoutBatchesDTOFieldNameType string

type payoutBatchesDTOFieldName struct {
	Id                PayoutBatchesDTOFieldNameType
	PayoutBatchCode   PayoutBatchesDTOFieldNameType
	ProviderAccountId PayoutBatchesDTOFieldNameType
	ScheduledFor      PayoutBatchesDTOFieldNameType
	CurrencyCode      PayoutBatchesDTOFieldNameType
	BatchStatus       PayoutBatchesDTOFieldNameType
	TotalCount        PayoutBatchesDTOFieldNameType
	TotalAmount       PayoutBatchesDTOFieldNameType
	Metadata          PayoutBatchesDTOFieldNameType
	MetaCreatedAt     PayoutBatchesDTOFieldNameType
	MetaCreatedBy     PayoutBatchesDTOFieldNameType
	MetaUpdatedAt     PayoutBatchesDTOFieldNameType
	MetaUpdatedBy     PayoutBatchesDTOFieldNameType
	MetaDeletedAt     PayoutBatchesDTOFieldNameType
	MetaDeletedBy     PayoutBatchesDTOFieldNameType
}

var PayoutBatchesDTOFieldName = payoutBatchesDTOFieldName{
	Id:                "id",
	PayoutBatchCode:   "payoutBatchCode",
	ProviderAccountId: "providerAccountId",
	ScheduledFor:      "scheduledFor",
	CurrencyCode:      "currencyCode",
	BatchStatus:       "batchStatus",
	TotalCount:        "totalCount",
	TotalAmount:       "totalAmount",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformPayoutBatchesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PayoutBatchesDTOFieldName.Id):
		return string(model.PayoutBatchesDBFieldName.Id), true

	case string(PayoutBatchesDTOFieldName.PayoutBatchCode):
		return string(model.PayoutBatchesDBFieldName.PayoutBatchCode), true

	case string(PayoutBatchesDTOFieldName.ProviderAccountId):
		return string(model.PayoutBatchesDBFieldName.ProviderAccountId), true

	case string(PayoutBatchesDTOFieldName.ScheduledFor):
		return string(model.PayoutBatchesDBFieldName.ScheduledFor), true

	case string(PayoutBatchesDTOFieldName.CurrencyCode):
		return string(model.PayoutBatchesDBFieldName.CurrencyCode), true

	case string(PayoutBatchesDTOFieldName.BatchStatus):
		return string(model.PayoutBatchesDBFieldName.BatchStatus), true

	case string(PayoutBatchesDTOFieldName.TotalCount):
		return string(model.PayoutBatchesDBFieldName.TotalCount), true

	case string(PayoutBatchesDTOFieldName.TotalAmount):
		return string(model.PayoutBatchesDBFieldName.TotalAmount), true

	case string(PayoutBatchesDTOFieldName.Metadata):
		return string(model.PayoutBatchesDBFieldName.Metadata), true

	case string(PayoutBatchesDTOFieldName.MetaCreatedAt):
		return string(model.PayoutBatchesDBFieldName.MetaCreatedAt), true

	case string(PayoutBatchesDTOFieldName.MetaCreatedBy):
		return string(model.PayoutBatchesDBFieldName.MetaCreatedBy), true

	case string(PayoutBatchesDTOFieldName.MetaUpdatedAt):
		return string(model.PayoutBatchesDBFieldName.MetaUpdatedAt), true

	case string(PayoutBatchesDTOFieldName.MetaUpdatedBy):
		return string(model.PayoutBatchesDBFieldName.MetaUpdatedBy), true

	case string(PayoutBatchesDTOFieldName.MetaDeletedAt):
		return string(model.PayoutBatchesDBFieldName.MetaDeletedAt), true

	case string(PayoutBatchesDTOFieldName.MetaDeletedBy):
		return string(model.PayoutBatchesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPayoutBatchesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPayoutBatchesBaseFilterField(field string) bool {
	spec, found := model.NewPayoutBatchesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePayoutBatchesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePayoutBatchesProjectionOutputPath(path string) error {
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

func transformPayoutBatchesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPayoutBatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPayoutBatchesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPayoutBatchesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPayoutBatchesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPayoutBatchesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePayoutBatchesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePayoutBatchesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPayoutBatchesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPayoutBatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPayoutBatchesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPayoutBatchesFilter(filter *model.Filter) {
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
			Field: string(PayoutBatchesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PayoutBatchesSelectableResponse map[string]interface{}
type PayoutBatchesSelectableListResponse []*PayoutBatchesSelectableResponse

func assignPayoutBatchesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPayoutBatchesSelectableValue(out PayoutBatchesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPayoutBatchesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPayoutBatchesSelectableResponse(payoutBatches model.PayoutBatches, filter model.Filter) PayoutBatchesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PayoutBatchesDBFieldName.Id),
			string(model.PayoutBatchesDBFieldName.PayoutBatchCode),
			string(model.PayoutBatchesDBFieldName.ProviderAccountId),
			string(model.PayoutBatchesDBFieldName.ScheduledFor),
			string(model.PayoutBatchesDBFieldName.CurrencyCode),
			string(model.PayoutBatchesDBFieldName.BatchStatus),
			string(model.PayoutBatchesDBFieldName.TotalCount),
			string(model.PayoutBatchesDBFieldName.TotalAmount),
			string(model.PayoutBatchesDBFieldName.Metadata),
			string(model.PayoutBatchesDBFieldName.MetaCreatedAt),
			string(model.PayoutBatchesDBFieldName.MetaCreatedBy),
			string(model.PayoutBatchesDBFieldName.MetaUpdatedAt),
			string(model.PayoutBatchesDBFieldName.MetaUpdatedBy),
			string(model.PayoutBatchesDBFieldName.MetaDeletedAt),
			string(model.PayoutBatchesDBFieldName.MetaDeletedBy),
		)
	}
	payoutBatchesSelectableResponse := PayoutBatchesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PayoutBatchesDBFieldName.Id):
			key := string(PayoutBatchesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.Id, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.PayoutBatchCode):
			key := string(PayoutBatchesDTOFieldName.PayoutBatchCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.PayoutBatchCode, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.ProviderAccountId):
			key := string(PayoutBatchesDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.ProviderAccountId.UUID, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.ScheduledFor):
			key := string(PayoutBatchesDTOFieldName.ScheduledFor)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.ScheduledFor, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.CurrencyCode):
			key := string(PayoutBatchesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.CurrencyCode, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.BatchStatus):
			key := string(PayoutBatchesDTOFieldName.BatchStatus)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, model.PayoutBatchesBatchStatus(payoutBatches.BatchStatus), explicitAlias)

		case string(model.PayoutBatchesDBFieldName.TotalCount):
			key := string(PayoutBatchesDTOFieldName.TotalCount)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.TotalCount, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.TotalAmount):
			key := string(PayoutBatchesDTOFieldName.TotalAmount)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.TotalAmount, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.Metadata):
			key := string(PayoutBatchesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.Metadata, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.MetaCreatedAt):
			key := string(PayoutBatchesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.MetaCreatedAt, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.MetaCreatedBy):
			key := string(PayoutBatchesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.MetaCreatedBy, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.MetaUpdatedAt):
			key := string(PayoutBatchesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.MetaUpdatedAt, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.MetaUpdatedBy):
			key := string(PayoutBatchesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.MetaUpdatedBy, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.MetaDeletedAt):
			key := string(PayoutBatchesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.MetaDeletedAt.Time, explicitAlias)

		case string(model.PayoutBatchesDBFieldName.MetaDeletedBy):
			key := string(PayoutBatchesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPayoutBatchesSelectableValue(payoutBatchesSelectableResponse, key, payoutBatches.MetaDeletedBy, explicitAlias)

		}
	}
	return payoutBatchesSelectableResponse
}

func NewPayoutBatchesListResponseFromFilterResult(result []model.PayoutBatchesFilterResult, filter model.Filter) PayoutBatchesSelectableListResponse {
	dtoPayoutBatchesListResponse := PayoutBatchesSelectableListResponse{}
	for _, row := range result {
		dtoPayoutBatchesResponse := NewPayoutBatchesSelectableResponse(row.PayoutBatches, filter)
		dtoPayoutBatchesListResponse = append(dtoPayoutBatchesListResponse, &dtoPayoutBatchesResponse)
	}
	return dtoPayoutBatchesListResponse
}

type PayoutBatchesFilterResponse struct {
	Metadata Metadata                            `json:"metadata"`
	Data     PayoutBatchesSelectableListResponse `json:"data"`
}

func reversePayoutBatchesFilterResults(result []model.PayoutBatchesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPayoutBatchesFilterResponse(result []model.PayoutBatchesFilterResult, filter model.Filter) (resp PayoutBatchesFilterResponse) {
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
			reversePayoutBatchesFilterResults(dataResult)
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

	resp.Data = NewPayoutBatchesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PayoutBatchesCreateRequest struct {
	PayoutBatchCode   string                         `json:"payoutBatchCode"`
	ProviderAccountId uuid.UUID                      `json:"providerAccountId"`
	ScheduledFor      time.Time                      `json:"scheduledFor"`
	CurrencyCode      string                         `json:"currencyCode"`
	BatchStatus       model.PayoutBatchesBatchStatus `json:"batchStatus" example:"draft" enums:"draft,queued,processing,partially_succeeded,succeeded,failed,cancelled"`
	TotalCount        int                            `json:"totalCount"`
	TotalAmount       decimal.Decimal                `json:"totalAmount"`
	Metadata          json.RawMessage                `json:"metadata"`
}

func (d *PayoutBatchesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PayoutBatchesCreateRequest) ToModel() model.PayoutBatches {
	id, _ := uuid.NewV4()
	return model.PayoutBatches{
		Id:                id,
		PayoutBatchCode:   d.PayoutBatchCode,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		ScheduledFor:      d.ScheduledFor,
		CurrencyCode:      d.CurrencyCode,
		BatchStatus:       d.BatchStatus,
		TotalCount:        d.TotalCount,
		TotalAmount:       d.TotalAmount,
		Metadata:          d.Metadata,
	}
}

type PayoutBatchesListCreateRequest []*PayoutBatchesCreateRequest

func (d PayoutBatchesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutBatches := range d {
		err = validator.Struct(payoutBatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutBatchesListCreateRequest) ToModelList() []model.PayoutBatches {
	out := make([]model.PayoutBatches, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PayoutBatchesUpdateRequest struct {
	PayoutBatchCode   string                         `json:"payoutBatchCode"`
	ProviderAccountId uuid.UUID                      `json:"providerAccountId"`
	ScheduledFor      time.Time                      `json:"scheduledFor"`
	CurrencyCode      string                         `json:"currencyCode"`
	BatchStatus       model.PayoutBatchesBatchStatus `json:"batchStatus" example:"draft" enums:"draft,queued,processing,partially_succeeded,succeeded,failed,cancelled"`
	TotalCount        int                            `json:"totalCount"`
	TotalAmount       decimal.Decimal                `json:"totalAmount"`
	Metadata          json.RawMessage                `json:"metadata"`
}

func (d *PayoutBatchesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PayoutBatchesUpdateRequest) ToModel() model.PayoutBatches {
	return model.PayoutBatches{
		PayoutBatchCode:   d.PayoutBatchCode,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		ScheduledFor:      d.ScheduledFor,
		CurrencyCode:      d.CurrencyCode,
		BatchStatus:       d.BatchStatus,
		TotalCount:        d.TotalCount,
		TotalAmount:       d.TotalAmount,
		Metadata:          d.Metadata,
	}
}

type PayoutBatchesBulkUpdateRequest struct {
	Id                uuid.UUID                      `json:"id"`
	PayoutBatchCode   string                         `json:"payoutBatchCode"`
	ProviderAccountId uuid.UUID                      `json:"providerAccountId"`
	ScheduledFor      time.Time                      `json:"scheduledFor"`
	CurrencyCode      string                         `json:"currencyCode"`
	BatchStatus       model.PayoutBatchesBatchStatus `json:"batchStatus" example:"draft" enums:"draft,queued,processing,partially_succeeded,succeeded,failed,cancelled"`
	TotalCount        int                            `json:"totalCount"`
	TotalAmount       decimal.Decimal                `json:"totalAmount"`
	Metadata          json.RawMessage                `json:"metadata"`
}

func (d PayoutBatchesBulkUpdateRequest) PrimaryID() PayoutBatchesPrimaryID {
	return PayoutBatchesPrimaryID{
		Id: d.Id,
	}
}

type PayoutBatchesListBulkUpdateRequest []*PayoutBatchesBulkUpdateRequest

func (d PayoutBatchesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutBatches := range d {
		err = validator.Struct(payoutBatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PayoutBatchesBulkUpdateRequest) ToModel() model.PayoutBatches {
	return model.PayoutBatches{
		Id:                d.Id,
		PayoutBatchCode:   d.PayoutBatchCode,
		ProviderAccountId: nuuid.From(d.ProviderAccountId),
		ScheduledFor:      d.ScheduledFor,
		CurrencyCode:      d.CurrencyCode,
		BatchStatus:       d.BatchStatus,
		TotalCount:        d.TotalCount,
		TotalAmount:       d.TotalAmount,
		Metadata:          d.Metadata,
	}
}

type PayoutBatchesResponse struct {
	Id                uuid.UUID                      `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PayoutBatchCode   string                         `json:"payoutBatchCode" validate:"required"`
	ProviderAccountId uuid.UUID                      `json:"providerAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ScheduledFor      time.Time                      `json:"scheduledFor" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CurrencyCode      string                         `json:"currencyCode" validate:"required"`
	BatchStatus       model.PayoutBatchesBatchStatus `json:"batchStatus" validate:"oneof=draft queued processing partially_succeeded succeeded failed cancelled" enums:"draft,queued,processing,partially_succeeded,succeeded,failed,cancelled"`
	TotalCount        int                            `json:"totalCount" example:"1"`
	TotalAmount       decimal.Decimal                `json:"totalAmount" format:"decimal" example:"100.50"`
	Metadata          json.RawMessage                `json:"metadata" swaggertype:"object"`
}

func NewPayoutBatchesResponse(payoutBatches model.PayoutBatches) PayoutBatchesResponse {
	return PayoutBatchesResponse{
		Id:                payoutBatches.Id,
		PayoutBatchCode:   payoutBatches.PayoutBatchCode,
		ProviderAccountId: payoutBatches.ProviderAccountId.UUID,
		ScheduledFor:      payoutBatches.ScheduledFor,
		CurrencyCode:      payoutBatches.CurrencyCode,
		BatchStatus:       model.PayoutBatchesBatchStatus(payoutBatches.BatchStatus),
		TotalCount:        payoutBatches.TotalCount,
		TotalAmount:       payoutBatches.TotalAmount,
		Metadata:          payoutBatches.Metadata,
	}
}

type PayoutBatchesListResponse []*PayoutBatchesResponse

func NewPayoutBatchesListResponse(payoutBatchesList model.PayoutBatchesList) PayoutBatchesListResponse {
	dtoPayoutBatchesListResponse := PayoutBatchesListResponse{}
	for _, payoutBatches := range payoutBatchesList {
		dtoPayoutBatchesResponse := NewPayoutBatchesResponse(*payoutBatches)
		dtoPayoutBatchesListResponse = append(dtoPayoutBatchesListResponse, &dtoPayoutBatchesResponse)
	}
	return dtoPayoutBatchesListResponse
}

type PayoutBatchesPrimaryIDList []PayoutBatchesPrimaryID

func (d PayoutBatchesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, payoutBatches := range d {
		err = validator.Struct(payoutBatches)
		if err != nil {
			return
		}
	}
	return nil
}

type PayoutBatchesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PayoutBatchesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PayoutBatchesPrimaryID) ToModel() model.PayoutBatchesPrimaryID {
	return model.PayoutBatchesPrimaryID{
		Id: d.Id,
	}
}
