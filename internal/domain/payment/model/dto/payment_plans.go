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

type PaymentPlansDTOFieldNameType string

type paymentPlansDTOFieldName struct {
	Id                        PaymentPlansDTOFieldNameType
	PaymentIntentId           PaymentPlansDTOFieldNameType
	PlanType                  PaymentPlansDTOFieldNameType
	Status                    PaymentPlansDTOFieldNameType
	TotalAmount               PaymentPlansDTOFieldNameType
	Currency                  PaymentPlansDTOFieldNameType
	InstallmentCount          PaymentPlansDTOFieldNameType
	DepositAmount             PaymentPlansDTOFieldNameType
	AutoCancelOnDefault       PaymentPlansDTOFieldNameType
	DefaultGracePeriodSeconds PaymentPlansDTOFieldNameType
	CompletedAt               PaymentPlansDTOFieldNameType
	CanceledAt                PaymentPlansDTOFieldNameType
	Metadata                  PaymentPlansDTOFieldNameType
	MetaCreatedAt             PaymentPlansDTOFieldNameType
	MetaCreatedBy             PaymentPlansDTOFieldNameType
	MetaUpdatedAt             PaymentPlansDTOFieldNameType
	MetaUpdatedBy             PaymentPlansDTOFieldNameType
	MetaDeletedAt             PaymentPlansDTOFieldNameType
	MetaDeletedBy             PaymentPlansDTOFieldNameType
}

var PaymentPlansDTOFieldName = paymentPlansDTOFieldName{
	Id:                        "id",
	PaymentIntentId:           "paymentIntentId",
	PlanType:                  "planType",
	Status:                    "status",
	TotalAmount:               "totalAmount",
	Currency:                  "currency",
	InstallmentCount:          "installmentCount",
	DepositAmount:             "depositAmount",
	AutoCancelOnDefault:       "autoCancelOnDefault",
	DefaultGracePeriodSeconds: "defaultGracePeriodSeconds",
	CompletedAt:               "completedAt",
	CanceledAt:                "canceledAt",
	Metadata:                  "metadata",
	MetaCreatedAt:             "metaCreatedAt",
	MetaCreatedBy:             "metaCreatedBy",
	MetaUpdatedAt:             "metaUpdatedAt",
	MetaUpdatedBy:             "metaUpdatedBy",
	MetaDeletedAt:             "metaDeletedAt",
	MetaDeletedBy:             "metaDeletedBy",
}

func transformPaymentPlansDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentPlansDTOFieldName.Id):
		return string(model.PaymentPlansDBFieldName.Id), true

	case string(PaymentPlansDTOFieldName.PaymentIntentId):
		return string(model.PaymentPlansDBFieldName.PaymentIntentId), true

	case string(PaymentPlansDTOFieldName.PlanType):
		return string(model.PaymentPlansDBFieldName.PlanType), true

	case string(PaymentPlansDTOFieldName.Status):
		return string(model.PaymentPlansDBFieldName.Status), true

	case string(PaymentPlansDTOFieldName.TotalAmount):
		return string(model.PaymentPlansDBFieldName.TotalAmount), true

	case string(PaymentPlansDTOFieldName.Currency):
		return string(model.PaymentPlansDBFieldName.Currency), true

	case string(PaymentPlansDTOFieldName.InstallmentCount):
		return string(model.PaymentPlansDBFieldName.InstallmentCount), true

	case string(PaymentPlansDTOFieldName.DepositAmount):
		return string(model.PaymentPlansDBFieldName.DepositAmount), true

	case string(PaymentPlansDTOFieldName.AutoCancelOnDefault):
		return string(model.PaymentPlansDBFieldName.AutoCancelOnDefault), true

	case string(PaymentPlansDTOFieldName.DefaultGracePeriodSeconds):
		return string(model.PaymentPlansDBFieldName.DefaultGracePeriodSeconds), true

	case string(PaymentPlansDTOFieldName.CompletedAt):
		return string(model.PaymentPlansDBFieldName.CompletedAt), true

	case string(PaymentPlansDTOFieldName.CanceledAt):
		return string(model.PaymentPlansDBFieldName.CanceledAt), true

	case string(PaymentPlansDTOFieldName.Metadata):
		return string(model.PaymentPlansDBFieldName.Metadata), true

	case string(PaymentPlansDTOFieldName.MetaCreatedAt):
		return string(model.PaymentPlansDBFieldName.MetaCreatedAt), true

	case string(PaymentPlansDTOFieldName.MetaCreatedBy):
		return string(model.PaymentPlansDBFieldName.MetaCreatedBy), true

	case string(PaymentPlansDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentPlansDBFieldName.MetaUpdatedAt), true

	case string(PaymentPlansDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentPlansDBFieldName.MetaUpdatedBy), true

	case string(PaymentPlansDTOFieldName.MetaDeletedAt):
		return string(model.PaymentPlansDBFieldName.MetaDeletedAt), true

	case string(PaymentPlansDTOFieldName.MetaDeletedBy):
		return string(model.PaymentPlansDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentPlansFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentPlansBaseFilterField(field string) bool {
	spec, found := model.NewPaymentPlansFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentPlansProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentPlansProjectionOutputPath(path string) error {
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

func transformPaymentPlansFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentPlansDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentPlansFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentPlansFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentPlansDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentPlansBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentPlansProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentPlansProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentPlansDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentPlansDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentPlansFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentPlansFilter(filter *model.Filter) {
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
			Field: string(PaymentPlansDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentPlansSelectableResponse map[string]interface{}
type PaymentPlansSelectableListResponse []*PaymentPlansSelectableResponse

func assignPaymentPlansNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentPlansSelectableValue(out PaymentPlansSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentPlansNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentPlansSelectableResponse(paymentPlans model.PaymentPlans, filter model.Filter) PaymentPlansSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentPlansDBFieldName.Id),
			string(model.PaymentPlansDBFieldName.PaymentIntentId),
			string(model.PaymentPlansDBFieldName.PlanType),
			string(model.PaymentPlansDBFieldName.Status),
			string(model.PaymentPlansDBFieldName.TotalAmount),
			string(model.PaymentPlansDBFieldName.Currency),
			string(model.PaymentPlansDBFieldName.InstallmentCount),
			string(model.PaymentPlansDBFieldName.DepositAmount),
			string(model.PaymentPlansDBFieldName.AutoCancelOnDefault),
			string(model.PaymentPlansDBFieldName.DefaultGracePeriodSeconds),
			string(model.PaymentPlansDBFieldName.CompletedAt),
			string(model.PaymentPlansDBFieldName.CanceledAt),
			string(model.PaymentPlansDBFieldName.Metadata),
			string(model.PaymentPlansDBFieldName.MetaCreatedAt),
			string(model.PaymentPlansDBFieldName.MetaCreatedBy),
			string(model.PaymentPlansDBFieldName.MetaUpdatedAt),
			string(model.PaymentPlansDBFieldName.MetaUpdatedBy),
			string(model.PaymentPlansDBFieldName.MetaDeletedAt),
			string(model.PaymentPlansDBFieldName.MetaDeletedBy),
		)
	}
	paymentPlansSelectableResponse := PaymentPlansSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentPlansDBFieldName.Id):
			key := string(PaymentPlansDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.Id, explicitAlias)

		case string(model.PaymentPlansDBFieldName.PaymentIntentId):
			key := string(PaymentPlansDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.PaymentIntentId, explicitAlias)

		case string(model.PaymentPlansDBFieldName.PlanType):
			key := string(PaymentPlansDTOFieldName.PlanType)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, model.PaymentPlanType(paymentPlans.PlanType), explicitAlias)

		case string(model.PaymentPlansDBFieldName.Status):
			key := string(PaymentPlansDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, model.PaymentPlanStatus(paymentPlans.Status), explicitAlias)

		case string(model.PaymentPlansDBFieldName.TotalAmount):
			key := string(PaymentPlansDTOFieldName.TotalAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.TotalAmount, explicitAlias)

		case string(model.PaymentPlansDBFieldName.Currency):
			key := string(PaymentPlansDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.Currency, explicitAlias)

		case string(model.PaymentPlansDBFieldName.InstallmentCount):
			key := string(PaymentPlansDTOFieldName.InstallmentCount)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.InstallmentCount, explicitAlias)

		case string(model.PaymentPlansDBFieldName.DepositAmount):
			key := string(PaymentPlansDTOFieldName.DepositAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.DepositAmount.Decimal, explicitAlias)

		case string(model.PaymentPlansDBFieldName.AutoCancelOnDefault):
			key := string(PaymentPlansDTOFieldName.AutoCancelOnDefault)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.AutoCancelOnDefault, explicitAlias)

		case string(model.PaymentPlansDBFieldName.DefaultGracePeriodSeconds):
			key := string(PaymentPlansDTOFieldName.DefaultGracePeriodSeconds)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.DefaultGracePeriodSeconds, explicitAlias)

		case string(model.PaymentPlansDBFieldName.CompletedAt):
			key := string(PaymentPlansDTOFieldName.CompletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.CompletedAt.Time, explicitAlias)

		case string(model.PaymentPlansDBFieldName.CanceledAt):
			key := string(PaymentPlansDTOFieldName.CanceledAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.CanceledAt.Time, explicitAlias)

		case string(model.PaymentPlansDBFieldName.Metadata):
			key := string(PaymentPlansDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.Metadata, explicitAlias)

		case string(model.PaymentPlansDBFieldName.MetaCreatedAt):
			key := string(PaymentPlansDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.MetaCreatedAt, explicitAlias)

		case string(model.PaymentPlansDBFieldName.MetaCreatedBy):
			key := string(PaymentPlansDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.MetaCreatedBy, explicitAlias)

		case string(model.PaymentPlansDBFieldName.MetaUpdatedAt):
			key := string(PaymentPlansDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentPlansDBFieldName.MetaUpdatedBy):
			key := string(PaymentPlansDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.MetaUpdatedBy, explicitAlias)

		case string(model.PaymentPlansDBFieldName.MetaDeletedAt):
			key := string(PaymentPlansDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentPlansDBFieldName.MetaDeletedBy):
			key := string(PaymentPlansDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentPlansSelectableValue(paymentPlansSelectableResponse, key, paymentPlans.MetaDeletedBy, explicitAlias)

		}
	}
	return paymentPlansSelectableResponse
}

func NewPaymentPlansListResponseFromFilterResult(result []model.PaymentPlansFilterResult, filter model.Filter) PaymentPlansSelectableListResponse {
	dtoPaymentPlansListResponse := PaymentPlansSelectableListResponse{}
	for _, row := range result {
		dtoPaymentPlansResponse := NewPaymentPlansSelectableResponse(row.PaymentPlans, filter)
		dtoPaymentPlansListResponse = append(dtoPaymentPlansListResponse, &dtoPaymentPlansResponse)
	}
	return dtoPaymentPlansListResponse
}

type PaymentPlansFilterResponse struct {
	Metadata Metadata                           `json:"metadata"`
	Data     PaymentPlansSelectableListResponse `json:"data"`
}

func reversePaymentPlansFilterResults(result []model.PaymentPlansFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentPlansFilterResponse(result []model.PaymentPlansFilterResult, filter model.Filter) (resp PaymentPlansFilterResponse) {
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
			reversePaymentPlansFilterResults(dataResult)
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

	resp.Data = NewPaymentPlansListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentPlansCreateRequest struct {
	PaymentIntentId           uuid.UUID               `json:"paymentIntentId"`
	PlanType                  model.PaymentPlanType   `json:"planType" example:"installment" enums:"installment,subscription"`
	Status                    model.PaymentPlanStatus `json:"status" example:"active" enums:"active,completed,canceled"`
	TotalAmount               decimal.Decimal         `json:"totalAmount"`
	Currency                  string                  `json:"currency"`
	InstallmentCount          int                     `json:"installmentCount"`
	DepositAmount             decimal.Decimal         `json:"depositAmount"`
	AutoCancelOnDefault       bool                    `json:"autoCancelOnDefault"`
	DefaultGracePeriodSeconds int                     `json:"defaultGracePeriodSeconds"`
	CompletedAt               time.Time               `json:"completedAt"`
	CanceledAt                time.Time               `json:"canceledAt"`
	Metadata                  json.RawMessage         `json:"metadata"`
}

func (d *PaymentPlansCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentPlansCreateRequest) ToModel() model.PaymentPlans {
	id, _ := uuid.NewV4()
	return model.PaymentPlans{
		Id:                        id,
		PaymentIntentId:           d.PaymentIntentId,
		PlanType:                  d.PlanType,
		Status:                    d.Status,
		TotalAmount:               d.TotalAmount,
		Currency:                  d.Currency,
		InstallmentCount:          d.InstallmentCount,
		DepositAmount:             decimal.NewNullDecimal(d.DepositAmount),
		AutoCancelOnDefault:       d.AutoCancelOnDefault,
		DefaultGracePeriodSeconds: d.DefaultGracePeriodSeconds,
		CompletedAt:               null.TimeFrom(d.CompletedAt),
		CanceledAt:                null.TimeFrom(d.CanceledAt),
		Metadata:                  d.Metadata,
	}
}

type PaymentPlansListCreateRequest []*PaymentPlansCreateRequest

func (d PaymentPlansListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentPlans := range d {
		err = validator.Struct(paymentPlans)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentPlansListCreateRequest) ToModelList() []model.PaymentPlans {
	out := make([]model.PaymentPlans, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentPlansUpdateRequest struct {
	PaymentIntentId           uuid.UUID               `json:"paymentIntentId"`
	PlanType                  model.PaymentPlanType   `json:"planType" example:"installment" enums:"installment,subscription"`
	Status                    model.PaymentPlanStatus `json:"status" example:"active" enums:"active,completed,canceled"`
	TotalAmount               decimal.Decimal         `json:"totalAmount"`
	Currency                  string                  `json:"currency"`
	InstallmentCount          int                     `json:"installmentCount"`
	DepositAmount             decimal.Decimal         `json:"depositAmount"`
	AutoCancelOnDefault       bool                    `json:"autoCancelOnDefault"`
	DefaultGracePeriodSeconds int                     `json:"defaultGracePeriodSeconds"`
	CompletedAt               time.Time               `json:"completedAt"`
	CanceledAt                time.Time               `json:"canceledAt"`
	Metadata                  json.RawMessage         `json:"metadata"`
}

func (d *PaymentPlansUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentPlansUpdateRequest) ToModel() model.PaymentPlans {
	return model.PaymentPlans{
		PaymentIntentId:           d.PaymentIntentId,
		PlanType:                  d.PlanType,
		Status:                    d.Status,
		TotalAmount:               d.TotalAmount,
		Currency:                  d.Currency,
		InstallmentCount:          d.InstallmentCount,
		DepositAmount:             decimal.NewNullDecimal(d.DepositAmount),
		AutoCancelOnDefault:       d.AutoCancelOnDefault,
		DefaultGracePeriodSeconds: d.DefaultGracePeriodSeconds,
		CompletedAt:               null.TimeFrom(d.CompletedAt),
		CanceledAt:                null.TimeFrom(d.CanceledAt),
		Metadata:                  d.Metadata,
	}
}

type PaymentPlansBulkUpdateRequest struct {
	Id                        uuid.UUID               `json:"id"`
	PaymentIntentId           uuid.UUID               `json:"paymentIntentId"`
	PlanType                  model.PaymentPlanType   `json:"planType" example:"installment" enums:"installment,subscription"`
	Status                    model.PaymentPlanStatus `json:"status" example:"active" enums:"active,completed,canceled"`
	TotalAmount               decimal.Decimal         `json:"totalAmount"`
	Currency                  string                  `json:"currency"`
	InstallmentCount          int                     `json:"installmentCount"`
	DepositAmount             decimal.Decimal         `json:"depositAmount"`
	AutoCancelOnDefault       bool                    `json:"autoCancelOnDefault"`
	DefaultGracePeriodSeconds int                     `json:"defaultGracePeriodSeconds"`
	CompletedAt               time.Time               `json:"completedAt"`
	CanceledAt                time.Time               `json:"canceledAt"`
	Metadata                  json.RawMessage         `json:"metadata"`
}

func (d PaymentPlansBulkUpdateRequest) PrimaryID() PaymentPlansPrimaryID {
	return PaymentPlansPrimaryID{
		Id: d.Id,
	}
}

type PaymentPlansListBulkUpdateRequest []*PaymentPlansBulkUpdateRequest

func (d PaymentPlansListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentPlans := range d {
		err = validator.Struct(paymentPlans)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentPlansBulkUpdateRequest) ToModel() model.PaymentPlans {
	return model.PaymentPlans{
		Id:                        d.Id,
		PaymentIntentId:           d.PaymentIntentId,
		PlanType:                  d.PlanType,
		Status:                    d.Status,
		TotalAmount:               d.TotalAmount,
		Currency:                  d.Currency,
		InstallmentCount:          d.InstallmentCount,
		DepositAmount:             decimal.NewNullDecimal(d.DepositAmount),
		AutoCancelOnDefault:       d.AutoCancelOnDefault,
		DefaultGracePeriodSeconds: d.DefaultGracePeriodSeconds,
		CompletedAt:               null.TimeFrom(d.CompletedAt),
		CanceledAt:                null.TimeFrom(d.CanceledAt),
		Metadata:                  d.Metadata,
	}
}

type PaymentPlansResponse struct {
	Id                        uuid.UUID               `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId           uuid.UUID               `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PlanType                  model.PaymentPlanType   `json:"planType" validate:"required,oneof=installment subscription" enums:"installment,subscription"`
	Status                    model.PaymentPlanStatus `json:"status" validate:"oneof=active completed canceled" enums:"active,completed,canceled"`
	TotalAmount               decimal.Decimal         `json:"totalAmount" validate:"required" format:"decimal" example:"100.50"`
	Currency                  string                  `json:"currency"`
	InstallmentCount          int                     `json:"installmentCount" example:"1"`
	DepositAmount             decimal.Decimal         `json:"depositAmount" format:"decimal" example:"100.50"`
	AutoCancelOnDefault       bool                    `json:"autoCancelOnDefault" example:"true"`
	DefaultGracePeriodSeconds int                     `json:"defaultGracePeriodSeconds" example:"1"`
	CompletedAt               time.Time               `json:"completedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CanceledAt                time.Time               `json:"canceledAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata                  json.RawMessage         `json:"metadata" swaggertype:"object"`
}

func NewPaymentPlansResponse(paymentPlans model.PaymentPlans) PaymentPlansResponse {
	return PaymentPlansResponse{
		Id:                        paymentPlans.Id,
		PaymentIntentId:           paymentPlans.PaymentIntentId,
		PlanType:                  model.PaymentPlanType(paymentPlans.PlanType),
		Status:                    model.PaymentPlanStatus(paymentPlans.Status),
		TotalAmount:               paymentPlans.TotalAmount,
		Currency:                  paymentPlans.Currency,
		InstallmentCount:          paymentPlans.InstallmentCount,
		DepositAmount:             paymentPlans.DepositAmount.Decimal,
		AutoCancelOnDefault:       paymentPlans.AutoCancelOnDefault,
		DefaultGracePeriodSeconds: paymentPlans.DefaultGracePeriodSeconds,
		CompletedAt:               paymentPlans.CompletedAt.Time,
		CanceledAt:                paymentPlans.CanceledAt.Time,
		Metadata:                  paymentPlans.Metadata,
	}
}

type PaymentPlansListResponse []*PaymentPlansResponse

func NewPaymentPlansListResponse(paymentPlansList model.PaymentPlansList) PaymentPlansListResponse {
	dtoPaymentPlansListResponse := PaymentPlansListResponse{}
	for _, paymentPlans := range paymentPlansList {
		dtoPaymentPlansResponse := NewPaymentPlansResponse(*paymentPlans)
		dtoPaymentPlansListResponse = append(dtoPaymentPlansListResponse, &dtoPaymentPlansResponse)
	}
	return dtoPaymentPlansListResponse
}

type PaymentPlansPrimaryIDList []PaymentPlansPrimaryID

func (d PaymentPlansPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentPlans := range d {
		err = validator.Struct(paymentPlans)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentPlansPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentPlansPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentPlansPrimaryID) ToModel() model.PaymentPlansPrimaryID {
	return model.PaymentPlansPrimaryID{
		Id: d.Id,
	}
}
