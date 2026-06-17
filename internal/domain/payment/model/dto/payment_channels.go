package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentChannelsDTOFieldNameType string

type paymentChannelsDTOFieldName struct {
	Id            PaymentChannelsDTOFieldNameType
	MethodId      PaymentChannelsDTOFieldNameType
	Code          PaymentChannelsDTOFieldNameType
	Name          PaymentChannelsDTOFieldNameType
	CountryCode   PaymentChannelsDTOFieldNameType
	Currency      PaymentChannelsDTOFieldNameType
	Status        PaymentChannelsDTOFieldNameType
	Metadata      PaymentChannelsDTOFieldNameType
	MetaCreatedAt PaymentChannelsDTOFieldNameType
	MetaCreatedBy PaymentChannelsDTOFieldNameType
	MetaUpdatedAt PaymentChannelsDTOFieldNameType
	MetaUpdatedBy PaymentChannelsDTOFieldNameType
	MetaDeletedAt PaymentChannelsDTOFieldNameType
	MetaDeletedBy PaymentChannelsDTOFieldNameType
}

var PaymentChannelsDTOFieldName = paymentChannelsDTOFieldName{
	Id:            "id",
	MethodId:      "methodId",
	Code:          "code",
	Name:          "name",
	CountryCode:   "countryCode",
	Currency:      "currency",
	Status:        "status",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformPaymentChannelsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentChannelsDTOFieldName.Id):
		return string(model.PaymentChannelsDBFieldName.Id), true

	case string(PaymentChannelsDTOFieldName.MethodId):
		return string(model.PaymentChannelsDBFieldName.MethodId), true

	case string(PaymentChannelsDTOFieldName.Code):
		return string(model.PaymentChannelsDBFieldName.Code), true

	case string(PaymentChannelsDTOFieldName.Name):
		return string(model.PaymentChannelsDBFieldName.Name), true

	case string(PaymentChannelsDTOFieldName.CountryCode):
		return string(model.PaymentChannelsDBFieldName.CountryCode), true

	case string(PaymentChannelsDTOFieldName.Currency):
		return string(model.PaymentChannelsDBFieldName.Currency), true

	case string(PaymentChannelsDTOFieldName.Status):
		return string(model.PaymentChannelsDBFieldName.Status), true

	case string(PaymentChannelsDTOFieldName.Metadata):
		return string(model.PaymentChannelsDBFieldName.Metadata), true

	case string(PaymentChannelsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentChannelsDBFieldName.MetaCreatedAt), true

	case string(PaymentChannelsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentChannelsDBFieldName.MetaCreatedBy), true

	case string(PaymentChannelsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentChannelsDBFieldName.MetaUpdatedAt), true

	case string(PaymentChannelsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentChannelsDBFieldName.MetaUpdatedBy), true

	case string(PaymentChannelsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentChannelsDBFieldName.MetaDeletedAt), true

	case string(PaymentChannelsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentChannelsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentChannelsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentChannelsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentChannelsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentChannelsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentChannelsProjectionOutputPath(path string) error {
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

func transformPaymentChannelsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentChannelsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentChannelsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentChannelsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentChannelsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentChannelsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentChannelsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentChannelsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentChannelsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentChannelsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentChannelsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentChannelsFilter(filter *model.Filter) {
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
			Field: string(PaymentChannelsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentChannelsSelectableResponse map[string]interface{}
type PaymentChannelsSelectableListResponse []*PaymentChannelsSelectableResponse

func assignPaymentChannelsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentChannelsSelectableValue(out PaymentChannelsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentChannelsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentChannelsSelectableResponse(paymentChannels model.PaymentChannels, filter model.Filter) PaymentChannelsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentChannelsDBFieldName.Id),
			string(model.PaymentChannelsDBFieldName.MethodId),
			string(model.PaymentChannelsDBFieldName.Code),
			string(model.PaymentChannelsDBFieldName.Name),
			string(model.PaymentChannelsDBFieldName.CountryCode),
			string(model.PaymentChannelsDBFieldName.Currency),
			string(model.PaymentChannelsDBFieldName.Status),
			string(model.PaymentChannelsDBFieldName.Metadata),
			string(model.PaymentChannelsDBFieldName.MetaCreatedAt),
			string(model.PaymentChannelsDBFieldName.MetaCreatedBy),
			string(model.PaymentChannelsDBFieldName.MetaUpdatedAt),
			string(model.PaymentChannelsDBFieldName.MetaUpdatedBy),
			string(model.PaymentChannelsDBFieldName.MetaDeletedAt),
			string(model.PaymentChannelsDBFieldName.MetaDeletedBy),
		)
	}
	paymentChannelsSelectableResponse := PaymentChannelsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentChannelsDBFieldName.Id):
			key := string(PaymentChannelsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.Id, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.MethodId):
			key := string(PaymentChannelsDTOFieldName.MethodId)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.MethodId, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.Code):
			key := string(PaymentChannelsDTOFieldName.Code)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.Code, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.Name):
			key := string(PaymentChannelsDTOFieldName.Name)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.Name, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.CountryCode):
			key := string(PaymentChannelsDTOFieldName.CountryCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.CountryCode, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.Currency):
			key := string(PaymentChannelsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.Currency, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.Status):
			key := string(PaymentChannelsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, model.PaymentChannelStatus(paymentChannels.Status), explicitAlias)

		case string(model.PaymentChannelsDBFieldName.Metadata):
			key := string(PaymentChannelsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.Metadata, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.MetaCreatedAt):
			key := string(PaymentChannelsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.MetaCreatedAt, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.MetaCreatedBy):
			key := string(PaymentChannelsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.MetaCreatedBy, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.MetaUpdatedAt):
			key := string(PaymentChannelsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.MetaUpdatedBy):
			key := string(PaymentChannelsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.MetaDeletedAt):
			key := string(PaymentChannelsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentChannelsDBFieldName.MetaDeletedBy):
			key := string(PaymentChannelsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentChannelsSelectableValue(paymentChannelsSelectableResponse, key, paymentChannels.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentChannelsSelectableResponse
}

func NewPaymentChannelsListResponseFromFilterResult(result []model.PaymentChannelsFilterResult, filter model.Filter) PaymentChannelsSelectableListResponse {
	dtoPaymentChannelsListResponse := PaymentChannelsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentChannelsResponse := NewPaymentChannelsSelectableResponse(row.PaymentChannels, filter)
		dtoPaymentChannelsListResponse = append(dtoPaymentChannelsListResponse, &dtoPaymentChannelsResponse)
	}
	return dtoPaymentChannelsListResponse
}

type PaymentChannelsFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     PaymentChannelsSelectableListResponse `json:"data"`
}

func reversePaymentChannelsFilterResults(result []model.PaymentChannelsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentChannelsFilterResponse(result []model.PaymentChannelsFilterResult, filter model.Filter) (resp PaymentChannelsFilterResponse) {
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
			reversePaymentChannelsFilterResults(dataResult)
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

	resp.Data = NewPaymentChannelsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentChannelsCreateRequest struct {
	MethodId    uuid.UUID                  `json:"methodId"`
	Code        string                     `json:"code"`
	Name        string                     `json:"name"`
	CountryCode string                     `json:"countryCode"`
	Currency    string                     `json:"currency"`
	Status      model.PaymentChannelStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Metadata    json.RawMessage            `json:"metadata"`
}

func (d *PaymentChannelsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentChannelsCreateRequest) ToModel() model.PaymentChannels {
	id, _ := uuid.NewV4()
	return model.PaymentChannels{
		Id:          id,
		MethodId:    d.MethodId,
		Code:        d.Code,
		Name:        d.Name,
		CountryCode: d.CountryCode,
		Currency:    d.Currency,
		Status:      d.Status,
		Metadata:    d.Metadata,
	}
}

type PaymentChannelsListCreateRequest []*PaymentChannelsCreateRequest

func (d PaymentChannelsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentChannels := range d {
		err = validator.Struct(paymentChannels)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentChannelsListCreateRequest) ToModelList() []model.PaymentChannels {
	out := make([]model.PaymentChannels, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentChannelsUpdateRequest struct {
	MethodId    uuid.UUID                  `json:"methodId"`
	Code        string                     `json:"code"`
	Name        string                     `json:"name"`
	CountryCode string                     `json:"countryCode"`
	Currency    string                     `json:"currency"`
	Status      model.PaymentChannelStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Metadata    json.RawMessage            `json:"metadata"`
}

func (d *PaymentChannelsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentChannelsUpdateRequest) ToModel() model.PaymentChannels {
	return model.PaymentChannels{
		MethodId:    d.MethodId,
		Code:        d.Code,
		Name:        d.Name,
		CountryCode: d.CountryCode,
		Currency:    d.Currency,
		Status:      d.Status,
		Metadata:    d.Metadata,
	}
}

type PaymentChannelsBulkUpdateRequest struct {
	Id          uuid.UUID                  `json:"id"`
	MethodId    uuid.UUID                  `json:"methodId"`
	Code        string                     `json:"code"`
	Name        string                     `json:"name"`
	CountryCode string                     `json:"countryCode"`
	Currency    string                     `json:"currency"`
	Status      model.PaymentChannelStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Metadata    json.RawMessage            `json:"metadata"`
}

func (d PaymentChannelsBulkUpdateRequest) PrimaryID() PaymentChannelsPrimaryID {
	return PaymentChannelsPrimaryID{
		Id: d.Id,
	}
}

type PaymentChannelsListBulkUpdateRequest []*PaymentChannelsBulkUpdateRequest

func (d PaymentChannelsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentChannels := range d {
		err = validator.Struct(paymentChannels)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentChannelsBulkUpdateRequest) ToModel() model.PaymentChannels {
	return model.PaymentChannels{
		Id:          d.Id,
		MethodId:    d.MethodId,
		Code:        d.Code,
		Name:        d.Name,
		CountryCode: d.CountryCode,
		Currency:    d.Currency,
		Status:      d.Status,
		Metadata:    d.Metadata,
	}
}

type PaymentChannelsResponse struct {
	Id          uuid.UUID                  `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MethodId    uuid.UUID                  `json:"methodId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Code        string                     `json:"code" validate:"required"`
	Name        string                     `json:"name" validate:"required"`
	CountryCode string                     `json:"countryCode"`
	Currency    string                     `json:"currency"`
	Status      model.PaymentChannelStatus `json:"status" validate:"oneof=active inactive deprecated" enums:"active,inactive,deprecated"`
	Metadata    json.RawMessage            `json:"metadata" swaggertype:"object"`
}

func NewPaymentChannelsResponse(paymentChannels model.PaymentChannels) PaymentChannelsResponse {
	return PaymentChannelsResponse{
		Id:          paymentChannels.Id,
		MethodId:    paymentChannels.MethodId,
		Code:        paymentChannels.Code,
		Name:        paymentChannels.Name,
		CountryCode: paymentChannels.CountryCode,
		Currency:    paymentChannels.Currency,
		Status:      model.PaymentChannelStatus(paymentChannels.Status),
		Metadata:    paymentChannels.Metadata,
	}
}

type PaymentChannelsListResponse []*PaymentChannelsResponse

func NewPaymentChannelsListResponse(paymentChannelsList model.PaymentChannelsList) PaymentChannelsListResponse {
	dtoPaymentChannelsListResponse := PaymentChannelsListResponse{}
	for _, paymentChannels := range paymentChannelsList {
		dtoPaymentChannelsResponse := NewPaymentChannelsResponse(*paymentChannels)
		dtoPaymentChannelsListResponse = append(dtoPaymentChannelsListResponse, &dtoPaymentChannelsResponse)
	}
	return dtoPaymentChannelsListResponse
}

type PaymentChannelsPrimaryIDList []PaymentChannelsPrimaryID

func (d PaymentChannelsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentChannels := range d {
		err = validator.Struct(paymentChannels)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentChannelsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentChannelsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentChannelsPrimaryID) ToModel() model.PaymentChannelsPrimaryID {
	return model.PaymentChannelsPrimaryID{
		Id: d.Id,
	}
}
