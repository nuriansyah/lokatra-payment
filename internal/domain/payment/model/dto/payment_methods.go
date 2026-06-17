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

type PaymentMethodsDTOFieldNameType string

type paymentMethodsDTOFieldName struct {
	Id            PaymentMethodsDTOFieldNameType
	Code          PaymentMethodsDTOFieldNameType
	MethodType    PaymentMethodsDTOFieldNameType
	Name          PaymentMethodsDTOFieldNameType
	Status        PaymentMethodsDTOFieldNameType
	Metadata      PaymentMethodsDTOFieldNameType
	MetaCreatedAt PaymentMethodsDTOFieldNameType
	MetaCreatedBy PaymentMethodsDTOFieldNameType
	MetaUpdatedAt PaymentMethodsDTOFieldNameType
	MetaUpdatedBy PaymentMethodsDTOFieldNameType
	MetaDeletedAt PaymentMethodsDTOFieldNameType
	MetaDeletedBy PaymentMethodsDTOFieldNameType
}

var PaymentMethodsDTOFieldName = paymentMethodsDTOFieldName{
	Id:            "id",
	Code:          "code",
	MethodType:    "methodType",
	Name:          "name",
	Status:        "status",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformPaymentMethodsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentMethodsDTOFieldName.Id):
		return string(model.PaymentMethodsDBFieldName.Id), true

	case string(PaymentMethodsDTOFieldName.Code):
		return string(model.PaymentMethodsDBFieldName.Code), true

	case string(PaymentMethodsDTOFieldName.MethodType):
		return string(model.PaymentMethodsDBFieldName.MethodType), true

	case string(PaymentMethodsDTOFieldName.Name):
		return string(model.PaymentMethodsDBFieldName.Name), true

	case string(PaymentMethodsDTOFieldName.Status):
		return string(model.PaymentMethodsDBFieldName.Status), true

	case string(PaymentMethodsDTOFieldName.Metadata):
		return string(model.PaymentMethodsDBFieldName.Metadata), true

	case string(PaymentMethodsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentMethodsDBFieldName.MetaCreatedAt), true

	case string(PaymentMethodsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentMethodsDBFieldName.MetaCreatedBy), true

	case string(PaymentMethodsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentMethodsDBFieldName.MetaUpdatedAt), true

	case string(PaymentMethodsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentMethodsDBFieldName.MetaUpdatedBy), true

	case string(PaymentMethodsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentMethodsDBFieldName.MetaDeletedAt), true

	case string(PaymentMethodsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentMethodsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentMethodsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentMethodsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentMethodsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentMethodsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentMethodsProjectionOutputPath(path string) error {
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

func transformPaymentMethodsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentMethodsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentMethodsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentMethodsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentMethodsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentMethodsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentMethodsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentMethodsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentMethodsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentMethodsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentMethodsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentMethodsFilter(filter *model.Filter) {
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
			Field: string(PaymentMethodsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentMethodsSelectableResponse map[string]interface{}
type PaymentMethodsSelectableListResponse []*PaymentMethodsSelectableResponse

func assignPaymentMethodsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentMethodsSelectableValue(out PaymentMethodsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentMethodsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentMethodsSelectableResponse(paymentMethods model.PaymentMethods, filter model.Filter) PaymentMethodsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentMethodsDBFieldName.Id),
			string(model.PaymentMethodsDBFieldName.Code),
			string(model.PaymentMethodsDBFieldName.MethodType),
			string(model.PaymentMethodsDBFieldName.Name),
			string(model.PaymentMethodsDBFieldName.Status),
			string(model.PaymentMethodsDBFieldName.Metadata),
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
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentMethodsDBFieldName.Id):
			key := string(PaymentMethodsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.Id, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.Code):
			key := string(PaymentMethodsDTOFieldName.Code)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.Code, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.MethodType):
			key := string(PaymentMethodsDTOFieldName.MethodType)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, model.PaymentMethodType(paymentMethods.MethodType), explicitAlias)

		case string(model.PaymentMethodsDBFieldName.Name):
			key := string(PaymentMethodsDTOFieldName.Name)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.Name, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.Status):
			key := string(PaymentMethodsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, model.PaymentMethodStatus(paymentMethods.Status), explicitAlias)

		case string(model.PaymentMethodsDBFieldName.Metadata):
			key := string(PaymentMethodsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.Metadata, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.MetaCreatedAt):
			key := string(PaymentMethodsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.MetaCreatedAt, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.MetaCreatedBy):
			key := string(PaymentMethodsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.MetaCreatedBy, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.MetaUpdatedAt):
			key := string(PaymentMethodsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.MetaUpdatedBy):
			key := string(PaymentMethodsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.MetaDeletedAt):
			key := string(PaymentMethodsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentMethodsDBFieldName.MetaDeletedBy):
			key := string(PaymentMethodsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentMethodsSelectableValue(paymentMethodsSelectableResponse, key, paymentMethods.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentMethodsSelectableResponse
}

func NewPaymentMethodsListResponseFromFilterResult(result []model.PaymentMethodsFilterResult, filter model.Filter) PaymentMethodsSelectableListResponse {
	dtoPaymentMethodsListResponse := PaymentMethodsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentMethodsResponse := NewPaymentMethodsSelectableResponse(row.PaymentMethods, filter)
		dtoPaymentMethodsListResponse = append(dtoPaymentMethodsListResponse, &dtoPaymentMethodsResponse)
	}
	return dtoPaymentMethodsListResponse
}

type PaymentMethodsFilterResponse struct {
	Metadata Metadata                             `json:"metadata"`
	Data     PaymentMethodsSelectableListResponse `json:"data"`
}

func reversePaymentMethodsFilterResults(result []model.PaymentMethodsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentMethodsFilterResponse(result []model.PaymentMethodsFilterResult, filter model.Filter) (resp PaymentMethodsFilterResponse) {
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
			reversePaymentMethodsFilterResults(dataResult)
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

	resp.Data = NewPaymentMethodsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentMethodsCreateRequest struct {
	Code       string                    `json:"code"`
	MethodType model.PaymentMethodType   `json:"methodType" example:"card" enums:"card,bank_transfer,ewallet,cash,cod"`
	Name       string                    `json:"name"`
	Status     model.PaymentMethodStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Metadata   json.RawMessage           `json:"metadata"`
}

func (d *PaymentMethodsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentMethodsCreateRequest) ToModel() model.PaymentMethods {
	id, _ := uuid.NewV4()
	return model.PaymentMethods{
		Id:         id,
		Code:       d.Code,
		MethodType: d.MethodType,
		Name:       d.Name,
		Status:     d.Status,
		Metadata:   d.Metadata,
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
	Code       string                    `json:"code"`
	MethodType model.PaymentMethodType   `json:"methodType" example:"card" enums:"card,bank_transfer,ewallet,cash,cod"`
	Name       string                    `json:"name"`
	Status     model.PaymentMethodStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Metadata   json.RawMessage           `json:"metadata"`
}

func (d *PaymentMethodsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentMethodsUpdateRequest) ToModel() model.PaymentMethods {
	return model.PaymentMethods{
		Code:       d.Code,
		MethodType: d.MethodType,
		Name:       d.Name,
		Status:     d.Status,
		Metadata:   d.Metadata,
	}
}

type PaymentMethodsBulkUpdateRequest struct {
	Id         uuid.UUID                 `json:"id"`
	Code       string                    `json:"code"`
	MethodType model.PaymentMethodType   `json:"methodType" example:"card" enums:"card,bank_transfer,ewallet,cash,cod"`
	Name       string                    `json:"name"`
	Status     model.PaymentMethodStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	Metadata   json.RawMessage           `json:"metadata"`
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
		Id:         d.Id,
		Code:       d.Code,
		MethodType: d.MethodType,
		Name:       d.Name,
		Status:     d.Status,
		Metadata:   d.Metadata,
	}
}

type PaymentMethodsResponse struct {
	Id         uuid.UUID                 `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Code       string                    `json:"code" validate:"required"`
	MethodType model.PaymentMethodType   `json:"methodType" validate:"required,oneof=card bank_transfer ewallet cash cod" enums:"card,bank_transfer,ewallet,cash,cod"`
	Name       string                    `json:"name" validate:"required"`
	Status     model.PaymentMethodStatus `json:"status" validate:"oneof=active inactive deprecated" enums:"active,inactive,deprecated"`
	Metadata   json.RawMessage           `json:"metadata" swaggertype:"object"`
}

func NewPaymentMethodsResponse(paymentMethods model.PaymentMethods) PaymentMethodsResponse {
	return PaymentMethodsResponse{
		Id:         paymentMethods.Id,
		Code:       paymentMethods.Code,
		MethodType: model.PaymentMethodType(paymentMethods.MethodType),
		Name:       paymentMethods.Name,
		Status:     model.PaymentMethodStatus(paymentMethods.Status),
		Metadata:   paymentMethods.Metadata,
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
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
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
