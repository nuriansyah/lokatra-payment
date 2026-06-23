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

type PaymentProvidersDTOFieldNameType string

type paymentProvidersDTOFieldName struct {
	Id                    PaymentProvidersDTOFieldNameType
	Code                  PaymentProvidersDTOFieldNameType
	Name                  PaymentProvidersDTOFieldNameType
	ProviderType          PaymentProvidersDTOFieldNameType
	Status                PaymentProvidersDTOFieldNameType
	SupportsRefund        PaymentProvidersDTOFieldNameType
	SupportsPartialRefund PaymentProvidersDTOFieldNameType
	SupportsAuthorization PaymentProvidersDTOFieldNameType
	SupportsCapture       PaymentProvidersDTOFieldNameType
	SupportsVoid          PaymentProvidersDTOFieldNameType
	SupportsWebhook       PaymentProvidersDTOFieldNameType
	Metadata              PaymentProvidersDTOFieldNameType
	MetaCreatedAt         PaymentProvidersDTOFieldNameType
	MetaCreatedBy         PaymentProvidersDTOFieldNameType
	MetaUpdatedAt         PaymentProvidersDTOFieldNameType
	MetaUpdatedBy         PaymentProvidersDTOFieldNameType
	MetaDeletedAt         PaymentProvidersDTOFieldNameType
	MetaDeletedBy         PaymentProvidersDTOFieldNameType
}

var PaymentProvidersDTOFieldName = paymentProvidersDTOFieldName{
	Id:                    "id",
	Code:                  "code",
	Name:                  "name",
	ProviderType:          "providerType",
	Status:                "status",
	SupportsRefund:        "supportsRefund",
	SupportsPartialRefund: "supportsPartialRefund",
	SupportsAuthorization: "supportsAuthorization",
	SupportsCapture:       "supportsCapture",
	SupportsVoid:          "supportsVoid",
	SupportsWebhook:       "supportsWebhook",
	Metadata:              "metadata",
	MetaCreatedAt:         "metaCreatedAt",
	MetaCreatedBy:         "metaCreatedBy",
	MetaUpdatedAt:         "metaUpdatedAt",
	MetaUpdatedBy:         "metaUpdatedBy",
	MetaDeletedAt:         "metaDeletedAt",
	MetaDeletedBy:         "metaDeletedBy",
}

func transformPaymentProvidersDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentProvidersDTOFieldName.Id):
		return string(model.PaymentProvidersDBFieldName.Id), true

	case string(PaymentProvidersDTOFieldName.Code):
		return string(model.PaymentProvidersDBFieldName.Code), true

	case string(PaymentProvidersDTOFieldName.Name):
		return string(model.PaymentProvidersDBFieldName.Name), true

	case string(PaymentProvidersDTOFieldName.ProviderType):
		return string(model.PaymentProvidersDBFieldName.ProviderType), true

	case string(PaymentProvidersDTOFieldName.Status):
		return string(model.PaymentProvidersDBFieldName.Status), true

	case string(PaymentProvidersDTOFieldName.SupportsRefund):
		return string(model.PaymentProvidersDBFieldName.SupportsRefund), true

	case string(PaymentProvidersDTOFieldName.SupportsPartialRefund):
		return string(model.PaymentProvidersDBFieldName.SupportsPartialRefund), true

	case string(PaymentProvidersDTOFieldName.SupportsAuthorization):
		return string(model.PaymentProvidersDBFieldName.SupportsAuthorization), true

	case string(PaymentProvidersDTOFieldName.SupportsCapture):
		return string(model.PaymentProvidersDBFieldName.SupportsCapture), true

	case string(PaymentProvidersDTOFieldName.SupportsVoid):
		return string(model.PaymentProvidersDBFieldName.SupportsVoid), true

	case string(PaymentProvidersDTOFieldName.SupportsWebhook):
		return string(model.PaymentProvidersDBFieldName.SupportsWebhook), true

	case string(PaymentProvidersDTOFieldName.Metadata):
		return string(model.PaymentProvidersDBFieldName.Metadata), true

	case string(PaymentProvidersDTOFieldName.MetaCreatedAt):
		return string(model.PaymentProvidersDBFieldName.MetaCreatedAt), true

	case string(PaymentProvidersDTOFieldName.MetaCreatedBy):
		return string(model.PaymentProvidersDBFieldName.MetaCreatedBy), true

	case string(PaymentProvidersDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentProvidersDBFieldName.MetaUpdatedAt), true

	case string(PaymentProvidersDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentProvidersDBFieldName.MetaUpdatedBy), true

	case string(PaymentProvidersDTOFieldName.MetaDeletedAt):
		return string(model.PaymentProvidersDBFieldName.MetaDeletedAt), true

	case string(PaymentProvidersDTOFieldName.MetaDeletedBy):
		return string(model.PaymentProvidersDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentProvidersFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentProvidersBaseFilterField(field string) bool {
	spec, found := model.NewPaymentProvidersFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentProvidersProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentProvidersProjectionOutputPath(path string) error {
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

func transformPaymentProvidersFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentProvidersDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentProvidersFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentProvidersFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentProvidersDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentProvidersBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentProvidersProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentProvidersProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentProvidersDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentProvidersDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentProvidersFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentProvidersFilter(filter *model.Filter) {
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
			Field: string(PaymentProvidersDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentProvidersSelectableResponse map[string]interface{}
type PaymentProvidersSelectableListResponse []*PaymentProvidersSelectableResponse

func assignPaymentProvidersNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentProvidersSelectableValue(out PaymentProvidersSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentProvidersNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentProvidersSelectableResponse(paymentProviders model.PaymentProviders, filter model.Filter) PaymentProvidersSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentProvidersDBFieldName.Id),
			string(model.PaymentProvidersDBFieldName.Code),
			string(model.PaymentProvidersDBFieldName.Name),
			string(model.PaymentProvidersDBFieldName.ProviderType),
			string(model.PaymentProvidersDBFieldName.Status),
			string(model.PaymentProvidersDBFieldName.SupportsRefund),
			string(model.PaymentProvidersDBFieldName.SupportsPartialRefund),
			string(model.PaymentProvidersDBFieldName.SupportsAuthorization),
			string(model.PaymentProvidersDBFieldName.SupportsCapture),
			string(model.PaymentProvidersDBFieldName.SupportsVoid),
			string(model.PaymentProvidersDBFieldName.SupportsWebhook),
			string(model.PaymentProvidersDBFieldName.Metadata),
			string(model.PaymentProvidersDBFieldName.MetaCreatedAt),
			string(model.PaymentProvidersDBFieldName.MetaCreatedBy),
			string(model.PaymentProvidersDBFieldName.MetaUpdatedAt),
			string(model.PaymentProvidersDBFieldName.MetaUpdatedBy),
			string(model.PaymentProvidersDBFieldName.MetaDeletedAt),
			string(model.PaymentProvidersDBFieldName.MetaDeletedBy),
		)
	}
	paymentProvidersSelectableResponse := PaymentProvidersSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentProvidersDBFieldName.Id):
			key := string(PaymentProvidersDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.Id, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.Code):
			key := string(PaymentProvidersDTOFieldName.Code)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.Code, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.Name):
			key := string(PaymentProvidersDTOFieldName.Name)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.Name, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.ProviderType):
			key := string(PaymentProvidersDTOFieldName.ProviderType)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, model.ProviderType(paymentProviders.ProviderType), explicitAlias)

		case string(model.PaymentProvidersDBFieldName.Status):
			key := string(PaymentProvidersDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, model.ProviderStatus(paymentProviders.Status), explicitAlias)

		case string(model.PaymentProvidersDBFieldName.SupportsRefund):
			key := string(PaymentProvidersDTOFieldName.SupportsRefund)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.SupportsRefund, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.SupportsPartialRefund):
			key := string(PaymentProvidersDTOFieldName.SupportsPartialRefund)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.SupportsPartialRefund, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.SupportsAuthorization):
			key := string(PaymentProvidersDTOFieldName.SupportsAuthorization)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.SupportsAuthorization, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.SupportsCapture):
			key := string(PaymentProvidersDTOFieldName.SupportsCapture)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.SupportsCapture, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.SupportsVoid):
			key := string(PaymentProvidersDTOFieldName.SupportsVoid)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.SupportsVoid, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.SupportsWebhook):
			key := string(PaymentProvidersDTOFieldName.SupportsWebhook)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.SupportsWebhook, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.Metadata):
			key := string(PaymentProvidersDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.Metadata, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.MetaCreatedAt):
			key := string(PaymentProvidersDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.MetaCreatedAt, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.MetaCreatedBy):
			key := string(PaymentProvidersDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.MetaCreatedBy, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.MetaUpdatedAt):
			key := string(PaymentProvidersDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.MetaUpdatedBy):
			key := string(PaymentProvidersDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.MetaUpdatedBy, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.MetaDeletedAt):
			key := string(PaymentProvidersDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentProvidersDBFieldName.MetaDeletedBy):
			key := string(PaymentProvidersDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentProvidersSelectableValue(paymentProvidersSelectableResponse, key, paymentProviders.MetaDeletedBy, explicitAlias)

		}
	}
	return paymentProvidersSelectableResponse
}

func NewPaymentProvidersListResponseFromFilterResult(result []model.PaymentProvidersFilterResult, filter model.Filter) PaymentProvidersSelectableListResponse {
	dtoPaymentProvidersListResponse := PaymentProvidersSelectableListResponse{}
	for _, row := range result {
		dtoPaymentProvidersResponse := NewPaymentProvidersSelectableResponse(row.PaymentProviders, filter)
		dtoPaymentProvidersListResponse = append(dtoPaymentProvidersListResponse, &dtoPaymentProvidersResponse)
	}
	return dtoPaymentProvidersListResponse
}

type PaymentProvidersFilterResponse struct {
	Metadata Metadata                               `json:"metadata"`
	Data     PaymentProvidersSelectableListResponse `json:"data"`
}

func reversePaymentProvidersFilterResults(result []model.PaymentProvidersFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentProvidersFilterResponse(result []model.PaymentProvidersFilterResult, filter model.Filter) (resp PaymentProvidersFilterResponse) {
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
			reversePaymentProvidersFilterResults(dataResult)
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

	resp.Data = NewPaymentProvidersListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentProvidersCreateRequest struct {
	Code                  string               `json:"code"`
	Name                  string               `json:"name"`
	ProviderType          model.ProviderType   `json:"providerType" example:"gateway" enums:"gateway,wallet,bank_transfer,cash,cod"`
	Status                model.ProviderStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	SupportsRefund        bool                 `json:"supportsRefund"`
	SupportsPartialRefund bool                 `json:"supportsPartialRefund"`
	SupportsAuthorization bool                 `json:"supportsAuthorization"`
	SupportsCapture       bool                 `json:"supportsCapture"`
	SupportsVoid          bool                 `json:"supportsVoid"`
	SupportsWebhook       bool                 `json:"supportsWebhook"`
	Metadata              json.RawMessage      `json:"metadata"`
}

func (d *PaymentProvidersCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentProvidersCreateRequest) ToModel() model.PaymentProviders {
	id, _ := uuid.NewV4()
	return model.PaymentProviders{
		Id:                    id,
		Code:                  d.Code,
		Name:                  d.Name,
		ProviderType:          d.ProviderType,
		Status:                d.Status,
		SupportsRefund:        d.SupportsRefund,
		SupportsPartialRefund: d.SupportsPartialRefund,
		SupportsAuthorization: d.SupportsAuthorization,
		SupportsCapture:       d.SupportsCapture,
		SupportsVoid:          d.SupportsVoid,
		SupportsWebhook:       d.SupportsWebhook,
		Metadata:              d.Metadata,
	}
}

type PaymentProvidersListCreateRequest []*PaymentProvidersCreateRequest

func (d PaymentProvidersListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentProviders := range d {
		err = validator.Struct(paymentProviders)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentProvidersListCreateRequest) ToModelList() []model.PaymentProviders {
	out := make([]model.PaymentProviders, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentProvidersUpdateRequest struct {
	Code                  string               `json:"code"`
	Name                  string               `json:"name"`
	ProviderType          model.ProviderType   `json:"providerType" example:"gateway" enums:"gateway,wallet,bank_transfer,cash,cod"`
	Status                model.ProviderStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	SupportsRefund        bool                 `json:"supportsRefund"`
	SupportsPartialRefund bool                 `json:"supportsPartialRefund"`
	SupportsAuthorization bool                 `json:"supportsAuthorization"`
	SupportsCapture       bool                 `json:"supportsCapture"`
	SupportsVoid          bool                 `json:"supportsVoid"`
	SupportsWebhook       bool                 `json:"supportsWebhook"`
	Metadata              json.RawMessage      `json:"metadata"`
}

func (d *PaymentProvidersUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentProvidersUpdateRequest) ToModel() model.PaymentProviders {
	return model.PaymentProviders{
		Code:                  d.Code,
		Name:                  d.Name,
		ProviderType:          d.ProviderType,
		Status:                d.Status,
		SupportsRefund:        d.SupportsRefund,
		SupportsPartialRefund: d.SupportsPartialRefund,
		SupportsAuthorization: d.SupportsAuthorization,
		SupportsCapture:       d.SupportsCapture,
		SupportsVoid:          d.SupportsVoid,
		SupportsWebhook:       d.SupportsWebhook,
		Metadata:              d.Metadata,
	}
}

type PaymentProvidersBulkUpdateRequest struct {
	Id                    uuid.UUID            `json:"id"`
	Code                  string               `json:"code"`
	Name                  string               `json:"name"`
	ProviderType          model.ProviderType   `json:"providerType" example:"gateway" enums:"gateway,wallet,bank_transfer,cash,cod"`
	Status                model.ProviderStatus `json:"status" example:"active" enums:"active,inactive,deprecated"`
	SupportsRefund        bool                 `json:"supportsRefund"`
	SupportsPartialRefund bool                 `json:"supportsPartialRefund"`
	SupportsAuthorization bool                 `json:"supportsAuthorization"`
	SupportsCapture       bool                 `json:"supportsCapture"`
	SupportsVoid          bool                 `json:"supportsVoid"`
	SupportsWebhook       bool                 `json:"supportsWebhook"`
	Metadata              json.RawMessage      `json:"metadata"`
}

func (d PaymentProvidersBulkUpdateRequest) PrimaryID() PaymentProvidersPrimaryID {
	return PaymentProvidersPrimaryID{
		Id: d.Id,
	}
}

type PaymentProvidersListBulkUpdateRequest []*PaymentProvidersBulkUpdateRequest

func (d PaymentProvidersListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentProviders := range d {
		err = validator.Struct(paymentProviders)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentProvidersBulkUpdateRequest) ToModel() model.PaymentProviders {
	return model.PaymentProviders{
		Id:                    d.Id,
		Code:                  d.Code,
		Name:                  d.Name,
		ProviderType:          d.ProviderType,
		Status:                d.Status,
		SupportsRefund:        d.SupportsRefund,
		SupportsPartialRefund: d.SupportsPartialRefund,
		SupportsAuthorization: d.SupportsAuthorization,
		SupportsCapture:       d.SupportsCapture,
		SupportsVoid:          d.SupportsVoid,
		SupportsWebhook:       d.SupportsWebhook,
		Metadata:              d.Metadata,
	}
}

type PaymentProvidersResponse struct {
	Id                    uuid.UUID            `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Code                  string               `json:"code" validate:"required"`
	Name                  string               `json:"name" validate:"required"`
	ProviderType          model.ProviderType   `json:"providerType" validate:"oneof=gateway wallet bank_transfer cash cod" enums:"gateway,wallet,bank_transfer,cash,cod"`
	Status                model.ProviderStatus `json:"status" validate:"oneof=active inactive deprecated" enums:"active,inactive,deprecated"`
	SupportsRefund        bool                 `json:"supportsRefund" example:"true"`
	SupportsPartialRefund bool                 `json:"supportsPartialRefund" example:"true"`
	SupportsAuthorization bool                 `json:"supportsAuthorization" example:"true"`
	SupportsCapture       bool                 `json:"supportsCapture" example:"true"`
	SupportsVoid          bool                 `json:"supportsVoid" example:"true"`
	SupportsWebhook       bool                 `json:"supportsWebhook" example:"true"`
	Metadata              json.RawMessage      `json:"metadata" swaggertype:"object"`
}

func NewPaymentProvidersResponse(paymentProviders model.PaymentProviders) PaymentProvidersResponse {
	return PaymentProvidersResponse{
		Id:                    paymentProviders.Id,
		Code:                  paymentProviders.Code,
		Name:                  paymentProviders.Name,
		ProviderType:          model.ProviderType(paymentProviders.ProviderType),
		Status:                model.ProviderStatus(paymentProviders.Status),
		SupportsRefund:        paymentProviders.SupportsRefund,
		SupportsPartialRefund: paymentProviders.SupportsPartialRefund,
		SupportsAuthorization: paymentProviders.SupportsAuthorization,
		SupportsCapture:       paymentProviders.SupportsCapture,
		SupportsVoid:          paymentProviders.SupportsVoid,
		SupportsWebhook:       paymentProviders.SupportsWebhook,
		Metadata:              paymentProviders.Metadata,
	}
}

type PaymentProvidersListResponse []*PaymentProvidersResponse

func NewPaymentProvidersListResponse(paymentProvidersList model.PaymentProvidersList) PaymentProvidersListResponse {
	dtoPaymentProvidersListResponse := PaymentProvidersListResponse{}
	for _, paymentProviders := range paymentProvidersList {
		dtoPaymentProvidersResponse := NewPaymentProvidersResponse(*paymentProviders)
		dtoPaymentProvidersListResponse = append(dtoPaymentProvidersListResponse, &dtoPaymentProvidersResponse)
	}
	return dtoPaymentProvidersListResponse
}

type PaymentProvidersPrimaryIDList []PaymentProvidersPrimaryID

func (d PaymentProvidersPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentProviders := range d {
		err = validator.Struct(paymentProviders)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentProvidersPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentProvidersPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentProvidersPrimaryID) ToModel() model.PaymentProvidersPrimaryID {
	return model.PaymentProvidersPrimaryID{
		Id: d.Id,
	}
}
