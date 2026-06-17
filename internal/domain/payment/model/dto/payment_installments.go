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

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentInstallmentsDTOFieldNameType string

type paymentInstallmentsDTOFieldName struct {
	Id              PaymentInstallmentsDTOFieldNameType
	PaymentPlanId   PaymentInstallmentsDTOFieldNameType
	PaymentIntentId PaymentInstallmentsDTOFieldNameType
	InstallmentNo   PaymentInstallmentsDTOFieldNameType
	DueAmount       PaymentInstallmentsDTOFieldNameType
	PaidAmount      PaymentInstallmentsDTOFieldNameType
	Currency        PaymentInstallmentsDTOFieldNameType
	DueAt           PaymentInstallmentsDTOFieldNameType
	Status          PaymentInstallmentsDTOFieldNameType
	PaidAt          PaymentInstallmentsDTOFieldNameType
	OverdueAt       PaymentInstallmentsDTOFieldNameType
	Metadata        PaymentInstallmentsDTOFieldNameType
	MetaCreatedAt   PaymentInstallmentsDTOFieldNameType
	MetaCreatedBy   PaymentInstallmentsDTOFieldNameType
	MetaUpdatedAt   PaymentInstallmentsDTOFieldNameType
	MetaUpdatedBy   PaymentInstallmentsDTOFieldNameType
	MetaDeletedAt   PaymentInstallmentsDTOFieldNameType
	MetaDeletedBy   PaymentInstallmentsDTOFieldNameType
}

var PaymentInstallmentsDTOFieldName = paymentInstallmentsDTOFieldName{
	Id:              "id",
	PaymentPlanId:   "paymentPlanId",
	PaymentIntentId: "paymentIntentId",
	InstallmentNo:   "installmentNo",
	DueAmount:       "dueAmount",
	PaidAmount:      "paidAmount",
	Currency:        "currency",
	DueAt:           "dueAt",
	Status:          "status",
	PaidAt:          "paidAt",
	OverdueAt:       "overdueAt",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformPaymentInstallmentsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentInstallmentsDTOFieldName.Id):
		return string(model.PaymentInstallmentsDBFieldName.Id), true

	case string(PaymentInstallmentsDTOFieldName.PaymentPlanId):
		return string(model.PaymentInstallmentsDBFieldName.PaymentPlanId), true

	case string(PaymentInstallmentsDTOFieldName.PaymentIntentId):
		return string(model.PaymentInstallmentsDBFieldName.PaymentIntentId), true

	case string(PaymentInstallmentsDTOFieldName.InstallmentNo):
		return string(model.PaymentInstallmentsDBFieldName.InstallmentNo), true

	case string(PaymentInstallmentsDTOFieldName.DueAmount):
		return string(model.PaymentInstallmentsDBFieldName.DueAmount), true

	case string(PaymentInstallmentsDTOFieldName.PaidAmount):
		return string(model.PaymentInstallmentsDBFieldName.PaidAmount), true

	case string(PaymentInstallmentsDTOFieldName.Currency):
		return string(model.PaymentInstallmentsDBFieldName.Currency), true

	case string(PaymentInstallmentsDTOFieldName.DueAt):
		return string(model.PaymentInstallmentsDBFieldName.DueAt), true

	case string(PaymentInstallmentsDTOFieldName.Status):
		return string(model.PaymentInstallmentsDBFieldName.Status), true

	case string(PaymentInstallmentsDTOFieldName.PaidAt):
		return string(model.PaymentInstallmentsDBFieldName.PaidAt), true

	case string(PaymentInstallmentsDTOFieldName.OverdueAt):
		return string(model.PaymentInstallmentsDBFieldName.OverdueAt), true

	case string(PaymentInstallmentsDTOFieldName.Metadata):
		return string(model.PaymentInstallmentsDBFieldName.Metadata), true

	case string(PaymentInstallmentsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentInstallmentsDBFieldName.MetaCreatedAt), true

	case string(PaymentInstallmentsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentInstallmentsDBFieldName.MetaCreatedBy), true

	case string(PaymentInstallmentsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentInstallmentsDBFieldName.MetaUpdatedAt), true

	case string(PaymentInstallmentsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentInstallmentsDBFieldName.MetaUpdatedBy), true

	case string(PaymentInstallmentsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentInstallmentsDBFieldName.MetaDeletedAt), true

	case string(PaymentInstallmentsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentInstallmentsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentInstallmentsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentInstallmentsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentInstallmentsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentInstallmentsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentInstallmentsProjectionOutputPath(path string) error {
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

func transformPaymentInstallmentsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentInstallmentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentInstallmentsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentInstallmentsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentInstallmentsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentInstallmentsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentInstallmentsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentInstallmentsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentInstallmentsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentInstallmentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentInstallmentsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentInstallmentsFilter(filter *model.Filter) {
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
			Field: string(PaymentInstallmentsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentInstallmentsSelectableResponse map[string]interface{}
type PaymentInstallmentsSelectableListResponse []*PaymentInstallmentsSelectableResponse

func assignPaymentInstallmentsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentInstallmentsSelectableValue(out PaymentInstallmentsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentInstallmentsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentInstallmentsSelectableResponse(paymentInstallments model.PaymentInstallments, filter model.Filter) PaymentInstallmentsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentInstallmentsDBFieldName.Id),
			string(model.PaymentInstallmentsDBFieldName.PaymentPlanId),
			string(model.PaymentInstallmentsDBFieldName.PaymentIntentId),
			string(model.PaymentInstallmentsDBFieldName.InstallmentNo),
			string(model.PaymentInstallmentsDBFieldName.DueAmount),
			string(model.PaymentInstallmentsDBFieldName.PaidAmount),
			string(model.PaymentInstallmentsDBFieldName.Currency),
			string(model.PaymentInstallmentsDBFieldName.DueAt),
			string(model.PaymentInstallmentsDBFieldName.Status),
			string(model.PaymentInstallmentsDBFieldName.PaidAt),
			string(model.PaymentInstallmentsDBFieldName.OverdueAt),
			string(model.PaymentInstallmentsDBFieldName.Metadata),
			string(model.PaymentInstallmentsDBFieldName.MetaCreatedAt),
			string(model.PaymentInstallmentsDBFieldName.MetaCreatedBy),
			string(model.PaymentInstallmentsDBFieldName.MetaUpdatedAt),
			string(model.PaymentInstallmentsDBFieldName.MetaUpdatedBy),
			string(model.PaymentInstallmentsDBFieldName.MetaDeletedAt),
			string(model.PaymentInstallmentsDBFieldName.MetaDeletedBy),
		)
	}
	paymentInstallmentsSelectableResponse := PaymentInstallmentsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentInstallmentsDBFieldName.Id):
			key := string(PaymentInstallmentsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.Id, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.PaymentPlanId):
			key := string(PaymentInstallmentsDTOFieldName.PaymentPlanId)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.PaymentPlanId, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.PaymentIntentId):
			key := string(PaymentInstallmentsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.PaymentIntentId, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.InstallmentNo):
			key := string(PaymentInstallmentsDTOFieldName.InstallmentNo)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.InstallmentNo, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.DueAmount):
			key := string(PaymentInstallmentsDTOFieldName.DueAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.DueAmount, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.PaidAmount):
			key := string(PaymentInstallmentsDTOFieldName.PaidAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.PaidAmount, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.Currency):
			key := string(PaymentInstallmentsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.Currency, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.DueAt):
			key := string(PaymentInstallmentsDTOFieldName.DueAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.DueAt, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.Status):
			key := string(PaymentInstallmentsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, model.PaymentInstallmentStatus(paymentInstallments.Status), explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.PaidAt):
			key := string(PaymentInstallmentsDTOFieldName.PaidAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.PaidAt.Time, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.OverdueAt):
			key := string(PaymentInstallmentsDTOFieldName.OverdueAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.OverdueAt.Time, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.Metadata):
			key := string(PaymentInstallmentsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.Metadata, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.MetaCreatedAt):
			key := string(PaymentInstallmentsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.MetaCreatedAt, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.MetaCreatedBy):
			key := string(PaymentInstallmentsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.MetaCreatedBy, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.MetaUpdatedAt):
			key := string(PaymentInstallmentsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.MetaUpdatedBy):
			key := string(PaymentInstallmentsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.MetaDeletedAt):
			key := string(PaymentInstallmentsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentInstallmentsDBFieldName.MetaDeletedBy):
			key := string(PaymentInstallmentsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentInstallmentsSelectableValue(paymentInstallmentsSelectableResponse, key, paymentInstallments.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentInstallmentsSelectableResponse
}

func NewPaymentInstallmentsListResponseFromFilterResult(result []model.PaymentInstallmentsFilterResult, filter model.Filter) PaymentInstallmentsSelectableListResponse {
	dtoPaymentInstallmentsListResponse := PaymentInstallmentsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentInstallmentsResponse := NewPaymentInstallmentsSelectableResponse(row.PaymentInstallments, filter)
		dtoPaymentInstallmentsListResponse = append(dtoPaymentInstallmentsListResponse, &dtoPaymentInstallmentsResponse)
	}
	return dtoPaymentInstallmentsListResponse
}

type PaymentInstallmentsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     PaymentInstallmentsSelectableListResponse `json:"data"`
}

func reversePaymentInstallmentsFilterResults(result []model.PaymentInstallmentsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentInstallmentsFilterResponse(result []model.PaymentInstallmentsFilterResult, filter model.Filter) (resp PaymentInstallmentsFilterResponse) {
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
			reversePaymentInstallmentsFilterResults(dataResult)
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

	resp.Data = NewPaymentInstallmentsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentInstallmentsCreateRequest struct {
	PaymentPlanId   uuid.UUID                      `json:"paymentPlanId"`
	PaymentIntentId uuid.UUID                      `json:"paymentIntentId"`
	InstallmentNo   int                            `json:"installmentNo"`
	DueAmount       decimal.Decimal                `json:"dueAmount"`
	PaidAmount      decimal.Decimal                `json:"paidAmount"`
	Currency        string                         `json:"currency"`
	DueAt           time.Time                      `json:"dueAt"`
	Status          model.PaymentInstallmentStatus `json:"status" example:"pending" enums:"pending,paid,overdue,canceled"`
	PaidAt          time.Time                      `json:"paidAt"`
	OverdueAt       time.Time                      `json:"overdueAt"`
	Metadata        json.RawMessage                `json:"metadata"`
}

func (d *PaymentInstallmentsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentInstallmentsCreateRequest) ToModel() model.PaymentInstallments {
	id, _ := uuid.NewV4()
	return model.PaymentInstallments{
		Id:              id,
		PaymentPlanId:   d.PaymentPlanId,
		PaymentIntentId: d.PaymentIntentId,
		InstallmentNo:   d.InstallmentNo,
		DueAmount:       d.DueAmount,
		PaidAmount:      d.PaidAmount,
		Currency:        d.Currency,
		DueAt:           d.DueAt,
		Status:          d.Status,
		PaidAt:          null.TimeFrom(d.PaidAt),
		OverdueAt:       null.TimeFrom(d.OverdueAt),
		Metadata:        d.Metadata,
	}
}

type PaymentInstallmentsListCreateRequest []*PaymentInstallmentsCreateRequest

func (d PaymentInstallmentsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentInstallments := range d {
		err = validator.Struct(paymentInstallments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentInstallmentsListCreateRequest) ToModelList() []model.PaymentInstallments {
	out := make([]model.PaymentInstallments, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentInstallmentsUpdateRequest struct {
	PaymentPlanId   uuid.UUID                      `json:"paymentPlanId"`
	PaymentIntentId uuid.UUID                      `json:"paymentIntentId"`
	InstallmentNo   int                            `json:"installmentNo"`
	DueAmount       decimal.Decimal                `json:"dueAmount"`
	PaidAmount      decimal.Decimal                `json:"paidAmount"`
	Currency        string                         `json:"currency"`
	DueAt           time.Time                      `json:"dueAt"`
	Status          model.PaymentInstallmentStatus `json:"status" example:"pending" enums:"pending,paid,overdue,canceled"`
	PaidAt          time.Time                      `json:"paidAt"`
	OverdueAt       time.Time                      `json:"overdueAt"`
	Metadata        json.RawMessage                `json:"metadata"`
}

func (d *PaymentInstallmentsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentInstallmentsUpdateRequest) ToModel() model.PaymentInstallments {
	return model.PaymentInstallments{
		PaymentPlanId:   d.PaymentPlanId,
		PaymentIntentId: d.PaymentIntentId,
		InstallmentNo:   d.InstallmentNo,
		DueAmount:       d.DueAmount,
		PaidAmount:      d.PaidAmount,
		Currency:        d.Currency,
		DueAt:           d.DueAt,
		Status:          d.Status,
		PaidAt:          null.TimeFrom(d.PaidAt),
		OverdueAt:       null.TimeFrom(d.OverdueAt),
		Metadata:        d.Metadata,
	}
}

type PaymentInstallmentsBulkUpdateRequest struct {
	Id              uuid.UUID                      `json:"id"`
	PaymentPlanId   uuid.UUID                      `json:"paymentPlanId"`
	PaymentIntentId uuid.UUID                      `json:"paymentIntentId"`
	InstallmentNo   int                            `json:"installmentNo"`
	DueAmount       decimal.Decimal                `json:"dueAmount"`
	PaidAmount      decimal.Decimal                `json:"paidAmount"`
	Currency        string                         `json:"currency"`
	DueAt           time.Time                      `json:"dueAt"`
	Status          model.PaymentInstallmentStatus `json:"status" example:"pending" enums:"pending,paid,overdue,canceled"`
	PaidAt          time.Time                      `json:"paidAt"`
	OverdueAt       time.Time                      `json:"overdueAt"`
	Metadata        json.RawMessage                `json:"metadata"`
}

func (d PaymentInstallmentsBulkUpdateRequest) PrimaryID() PaymentInstallmentsPrimaryID {
	return PaymentInstallmentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentInstallmentsListBulkUpdateRequest []*PaymentInstallmentsBulkUpdateRequest

func (d PaymentInstallmentsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentInstallments := range d {
		err = validator.Struct(paymentInstallments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentInstallmentsBulkUpdateRequest) ToModel() model.PaymentInstallments {
	return model.PaymentInstallments{
		Id:              d.Id,
		PaymentPlanId:   d.PaymentPlanId,
		PaymentIntentId: d.PaymentIntentId,
		InstallmentNo:   d.InstallmentNo,
		DueAmount:       d.DueAmount,
		PaidAmount:      d.PaidAmount,
		Currency:        d.Currency,
		DueAt:           d.DueAt,
		Status:          d.Status,
		PaidAt:          null.TimeFrom(d.PaidAt),
		OverdueAt:       null.TimeFrom(d.OverdueAt),
		Metadata:        d.Metadata,
	}
}

type PaymentInstallmentsResponse struct {
	Id              uuid.UUID                      `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentPlanId   uuid.UUID                      `json:"paymentPlanId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId uuid.UUID                      `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	InstallmentNo   int                            `json:"installmentNo" validate:"required" example:"1"`
	DueAmount       decimal.Decimal                `json:"dueAmount" validate:"required" format:"decimal" example:"100.50"`
	PaidAmount      decimal.Decimal                `json:"paidAmount" format:"decimal" example:"100.50"`
	Currency        string                         `json:"currency"`
	DueAt           time.Time                      `json:"dueAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Status          model.PaymentInstallmentStatus `json:"status" validate:"oneof=pending paid overdue canceled" enums:"pending,paid,overdue,canceled"`
	PaidAt          time.Time                      `json:"paidAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	OverdueAt       time.Time                      `json:"overdueAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata        json.RawMessage                `json:"metadata" swaggertype:"object"`
}

func NewPaymentInstallmentsResponse(paymentInstallments model.PaymentInstallments) PaymentInstallmentsResponse {
	return PaymentInstallmentsResponse{
		Id:              paymentInstallments.Id,
		PaymentPlanId:   paymentInstallments.PaymentPlanId,
		PaymentIntentId: paymentInstallments.PaymentIntentId,
		InstallmentNo:   paymentInstallments.InstallmentNo,
		DueAmount:       paymentInstallments.DueAmount,
		PaidAmount:      paymentInstallments.PaidAmount,
		Currency:        paymentInstallments.Currency,
		DueAt:           paymentInstallments.DueAt,
		Status:          model.PaymentInstallmentStatus(paymentInstallments.Status),
		PaidAt:          paymentInstallments.PaidAt.Time,
		OverdueAt:       paymentInstallments.OverdueAt.Time,
		Metadata:        paymentInstallments.Metadata,
	}
}

type PaymentInstallmentsListResponse []*PaymentInstallmentsResponse

func NewPaymentInstallmentsListResponse(paymentInstallmentsList model.PaymentInstallmentsList) PaymentInstallmentsListResponse {
	dtoPaymentInstallmentsListResponse := PaymentInstallmentsListResponse{}
	for _, paymentInstallments := range paymentInstallmentsList {
		dtoPaymentInstallmentsResponse := NewPaymentInstallmentsResponse(*paymentInstallments)
		dtoPaymentInstallmentsListResponse = append(dtoPaymentInstallmentsListResponse, &dtoPaymentInstallmentsResponse)
	}
	return dtoPaymentInstallmentsListResponse
}

type PaymentInstallmentsPrimaryIDList []PaymentInstallmentsPrimaryID

func (d PaymentInstallmentsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentInstallments := range d {
		err = validator.Struct(paymentInstallments)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentInstallmentsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentInstallmentsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentInstallmentsPrimaryID) ToModel() model.PaymentInstallmentsPrimaryID {
	return model.PaymentInstallmentsPrimaryID{
		Id: d.Id,
	}
}
