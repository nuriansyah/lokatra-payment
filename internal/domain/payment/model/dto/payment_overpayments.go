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

type PaymentOverpaymentsDTOFieldNameType string

type paymentOverpaymentsDTOFieldName struct {
	Id                PaymentOverpaymentsDTOFieldNameType
	PaymentIntentId   PaymentOverpaymentsDTOFieldNameType
	PaidAttemptId     PaymentOverpaymentsDTOFieldNameType
	OverpaidAttemptId PaymentOverpaymentsDTOFieldNameType
	ExpectedAmount    PaymentOverpaymentsDTOFieldNameType
	ReceivedAmount    PaymentOverpaymentsDTOFieldNameType
	OverpaidAmount    PaymentOverpaymentsDTOFieldNameType
	Currency          PaymentOverpaymentsDTOFieldNameType
	Status            PaymentOverpaymentsDTOFieldNameType
	ResolutionAction  PaymentOverpaymentsDTOFieldNameType
	ResolutionNotes   PaymentOverpaymentsDTOFieldNameType
	ResolvedAt        PaymentOverpaymentsDTOFieldNameType
	ResolvedBy        PaymentOverpaymentsDTOFieldNameType
	Metadata          PaymentOverpaymentsDTOFieldNameType
	MetaCreatedAt     PaymentOverpaymentsDTOFieldNameType
	MetaCreatedBy     PaymentOverpaymentsDTOFieldNameType
	MetaUpdatedAt     PaymentOverpaymentsDTOFieldNameType
	MetaUpdatedBy     PaymentOverpaymentsDTOFieldNameType
	MetaDeletedAt     PaymentOverpaymentsDTOFieldNameType
	MetaDeletedBy     PaymentOverpaymentsDTOFieldNameType
}

var PaymentOverpaymentsDTOFieldName = paymentOverpaymentsDTOFieldName{
	Id:                "id",
	PaymentIntentId:   "paymentIntentId",
	PaidAttemptId:     "paidAttemptId",
	OverpaidAttemptId: "overpaidAttemptId",
	ExpectedAmount:    "expectedAmount",
	ReceivedAmount:    "receivedAmount",
	OverpaidAmount:    "overpaidAmount",
	Currency:          "currency",
	Status:            "status",
	ResolutionAction:  "resolutionAction",
	ResolutionNotes:   "resolutionNotes",
	ResolvedAt:        "resolvedAt",
	ResolvedBy:        "resolvedBy",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformPaymentOverpaymentsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentOverpaymentsDTOFieldName.Id):
		return string(model.PaymentOverpaymentsDBFieldName.Id), true

	case string(PaymentOverpaymentsDTOFieldName.PaymentIntentId):
		return string(model.PaymentOverpaymentsDBFieldName.PaymentIntentId), true

	case string(PaymentOverpaymentsDTOFieldName.PaidAttemptId):
		return string(model.PaymentOverpaymentsDBFieldName.PaidAttemptId), true

	case string(PaymentOverpaymentsDTOFieldName.OverpaidAttemptId):
		return string(model.PaymentOverpaymentsDBFieldName.OverpaidAttemptId), true

	case string(PaymentOverpaymentsDTOFieldName.ExpectedAmount):
		return string(model.PaymentOverpaymentsDBFieldName.ExpectedAmount), true

	case string(PaymentOverpaymentsDTOFieldName.ReceivedAmount):
		return string(model.PaymentOverpaymentsDBFieldName.ReceivedAmount), true

	case string(PaymentOverpaymentsDTOFieldName.OverpaidAmount):
		return string(model.PaymentOverpaymentsDBFieldName.OverpaidAmount), true

	case string(PaymentOverpaymentsDTOFieldName.Currency):
		return string(model.PaymentOverpaymentsDBFieldName.Currency), true

	case string(PaymentOverpaymentsDTOFieldName.Status):
		return string(model.PaymentOverpaymentsDBFieldName.Status), true

	case string(PaymentOverpaymentsDTOFieldName.ResolutionAction):
		return string(model.PaymentOverpaymentsDBFieldName.ResolutionAction), true

	case string(PaymentOverpaymentsDTOFieldName.ResolutionNotes):
		return string(model.PaymentOverpaymentsDBFieldName.ResolutionNotes), true

	case string(PaymentOverpaymentsDTOFieldName.ResolvedAt):
		return string(model.PaymentOverpaymentsDBFieldName.ResolvedAt), true

	case string(PaymentOverpaymentsDTOFieldName.ResolvedBy):
		return string(model.PaymentOverpaymentsDBFieldName.ResolvedBy), true

	case string(PaymentOverpaymentsDTOFieldName.Metadata):
		return string(model.PaymentOverpaymentsDBFieldName.Metadata), true

	case string(PaymentOverpaymentsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentOverpaymentsDBFieldName.MetaCreatedAt), true

	case string(PaymentOverpaymentsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentOverpaymentsDBFieldName.MetaCreatedBy), true

	case string(PaymentOverpaymentsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentOverpaymentsDBFieldName.MetaUpdatedAt), true

	case string(PaymentOverpaymentsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentOverpaymentsDBFieldName.MetaUpdatedBy), true

	case string(PaymentOverpaymentsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentOverpaymentsDBFieldName.MetaDeletedAt), true

	case string(PaymentOverpaymentsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentOverpaymentsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentOverpaymentsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentOverpaymentsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentOverpaymentsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentOverpaymentsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentOverpaymentsProjectionOutputPath(path string) error {
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

func transformPaymentOverpaymentsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentOverpaymentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentOverpaymentsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentOverpaymentsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentOverpaymentsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentOverpaymentsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentOverpaymentsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentOverpaymentsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentOverpaymentsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentOverpaymentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentOverpaymentsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentOverpaymentsFilter(filter *model.Filter) {
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
			Field: string(PaymentOverpaymentsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentOverpaymentsSelectableResponse map[string]interface{}
type PaymentOverpaymentsSelectableListResponse []*PaymentOverpaymentsSelectableResponse

func assignPaymentOverpaymentsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentOverpaymentsSelectableValue(out PaymentOverpaymentsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentOverpaymentsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentOverpaymentsSelectableResponse(paymentOverpayments model.PaymentOverpayments, filter model.Filter) PaymentOverpaymentsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentOverpaymentsDBFieldName.Id),
			string(model.PaymentOverpaymentsDBFieldName.PaymentIntentId),
			string(model.PaymentOverpaymentsDBFieldName.PaidAttemptId),
			string(model.PaymentOverpaymentsDBFieldName.OverpaidAttemptId),
			string(model.PaymentOverpaymentsDBFieldName.ExpectedAmount),
			string(model.PaymentOverpaymentsDBFieldName.ReceivedAmount),
			string(model.PaymentOverpaymentsDBFieldName.OverpaidAmount),
			string(model.PaymentOverpaymentsDBFieldName.Currency),
			string(model.PaymentOverpaymentsDBFieldName.Status),
			string(model.PaymentOverpaymentsDBFieldName.ResolutionAction),
			string(model.PaymentOverpaymentsDBFieldName.ResolutionNotes),
			string(model.PaymentOverpaymentsDBFieldName.ResolvedAt),
			string(model.PaymentOverpaymentsDBFieldName.ResolvedBy),
			string(model.PaymentOverpaymentsDBFieldName.Metadata),
			string(model.PaymentOverpaymentsDBFieldName.MetaCreatedAt),
			string(model.PaymentOverpaymentsDBFieldName.MetaCreatedBy),
			string(model.PaymentOverpaymentsDBFieldName.MetaUpdatedAt),
			string(model.PaymentOverpaymentsDBFieldName.MetaUpdatedBy),
			string(model.PaymentOverpaymentsDBFieldName.MetaDeletedAt),
			string(model.PaymentOverpaymentsDBFieldName.MetaDeletedBy),
		)
	}
	paymentOverpaymentsSelectableResponse := PaymentOverpaymentsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentOverpaymentsDBFieldName.Id):
			key := string(PaymentOverpaymentsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.Id, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.PaymentIntentId):
			key := string(PaymentOverpaymentsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.PaymentIntentId, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.PaidAttemptId):
			key := string(PaymentOverpaymentsDTOFieldName.PaidAttemptId)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.PaidAttemptId.UUID, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.OverpaidAttemptId):
			key := string(PaymentOverpaymentsDTOFieldName.OverpaidAttemptId)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.OverpaidAttemptId.UUID, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.ExpectedAmount):
			key := string(PaymentOverpaymentsDTOFieldName.ExpectedAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.ExpectedAmount, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.ReceivedAmount):
			key := string(PaymentOverpaymentsDTOFieldName.ReceivedAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.ReceivedAmount, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.OverpaidAmount):
			key := string(PaymentOverpaymentsDTOFieldName.OverpaidAmount)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.OverpaidAmount, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.Currency):
			key := string(PaymentOverpaymentsDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.Currency, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.Status):
			key := string(PaymentOverpaymentsDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.Status, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.ResolutionAction):
			key := string(PaymentOverpaymentsDTOFieldName.ResolutionAction)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.ResolutionAction.String, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.ResolutionNotes):
			key := string(PaymentOverpaymentsDTOFieldName.ResolutionNotes)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.ResolutionNotes.String, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.ResolvedAt):
			key := string(PaymentOverpaymentsDTOFieldName.ResolvedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.ResolvedAt.Time, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.ResolvedBy):
			key := string(PaymentOverpaymentsDTOFieldName.ResolvedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.ResolvedBy.UUID, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.Metadata):
			key := string(PaymentOverpaymentsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.Metadata, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.MetaCreatedAt):
			key := string(PaymentOverpaymentsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.MetaCreatedAt, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.MetaCreatedBy):
			key := string(PaymentOverpaymentsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.MetaCreatedBy, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.MetaUpdatedAt):
			key := string(PaymentOverpaymentsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.MetaUpdatedBy):
			key := string(PaymentOverpaymentsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.MetaDeletedAt):
			key := string(PaymentOverpaymentsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentOverpaymentsDBFieldName.MetaDeletedBy):
			key := string(PaymentOverpaymentsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentOverpaymentsSelectableValue(paymentOverpaymentsSelectableResponse, key, paymentOverpayments.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentOverpaymentsSelectableResponse
}

func NewPaymentOverpaymentsListResponseFromFilterResult(result []model.PaymentOverpaymentsFilterResult, filter model.Filter) PaymentOverpaymentsSelectableListResponse {
	dtoPaymentOverpaymentsListResponse := PaymentOverpaymentsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentOverpaymentsResponse := NewPaymentOverpaymentsSelectableResponse(row.PaymentOverpayments, filter)
		dtoPaymentOverpaymentsListResponse = append(dtoPaymentOverpaymentsListResponse, &dtoPaymentOverpaymentsResponse)
	}
	return dtoPaymentOverpaymentsListResponse
}

type PaymentOverpaymentsFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     PaymentOverpaymentsSelectableListResponse `json:"data"`
}

func reversePaymentOverpaymentsFilterResults(result []model.PaymentOverpaymentsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentOverpaymentsFilterResponse(result []model.PaymentOverpaymentsFilterResult, filter model.Filter) (resp PaymentOverpaymentsFilterResponse) {
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
			reversePaymentOverpaymentsFilterResults(dataResult)
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

	resp.Data = NewPaymentOverpaymentsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentOverpaymentsCreateRequest struct {
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId"`
	PaidAttemptId     uuid.UUID       `json:"paidAttemptId"`
	OverpaidAttemptId uuid.UUID       `json:"overpaidAttemptId"`
	ExpectedAmount    decimal.Decimal `json:"expectedAmount"`
	ReceivedAmount    decimal.Decimal `json:"receivedAmount"`
	OverpaidAmount    decimal.Decimal `json:"overpaidAmount"`
	Currency          string          `json:"currency"`
	Status            string          `json:"status"`
	ResolutionAction  string          `json:"resolutionAction"`
	ResolutionNotes   string          `json:"resolutionNotes"`
	ResolvedAt        time.Time       `json:"resolvedAt"`
	ResolvedBy        uuid.UUID       `json:"resolvedBy"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *PaymentOverpaymentsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentOverpaymentsCreateRequest) ToModel() model.PaymentOverpayments {
	id, _ := uuid.NewV4()
	return model.PaymentOverpayments{
		Id:                id,
		PaymentIntentId:   d.PaymentIntentId,
		PaidAttemptId:     nuuid.From(d.PaidAttemptId),
		OverpaidAttemptId: nuuid.From(d.OverpaidAttemptId),
		ExpectedAmount:    d.ExpectedAmount,
		ReceivedAmount:    d.ReceivedAmount,
		OverpaidAmount:    d.OverpaidAmount,
		Currency:          d.Currency,
		Status:            d.Status,
		ResolutionAction:  null.StringFrom(d.ResolutionAction),
		ResolutionNotes:   null.StringFrom(d.ResolutionNotes),
		ResolvedAt:        null.TimeFrom(d.ResolvedAt),
		ResolvedBy:        nuuid.From(d.ResolvedBy),
		Metadata:          d.Metadata,
	}
}

type PaymentOverpaymentsListCreateRequest []*PaymentOverpaymentsCreateRequest

func (d PaymentOverpaymentsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentOverpayments := range d {
		err = validator.Struct(paymentOverpayments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentOverpaymentsListCreateRequest) ToModelList() []model.PaymentOverpayments {
	out := make([]model.PaymentOverpayments, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentOverpaymentsUpdateRequest struct {
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId"`
	PaidAttemptId     uuid.UUID       `json:"paidAttemptId"`
	OverpaidAttemptId uuid.UUID       `json:"overpaidAttemptId"`
	ExpectedAmount    decimal.Decimal `json:"expectedAmount"`
	ReceivedAmount    decimal.Decimal `json:"receivedAmount"`
	OverpaidAmount    decimal.Decimal `json:"overpaidAmount"`
	Currency          string          `json:"currency"`
	Status            string          `json:"status"`
	ResolutionAction  string          `json:"resolutionAction"`
	ResolutionNotes   string          `json:"resolutionNotes"`
	ResolvedAt        time.Time       `json:"resolvedAt"`
	ResolvedBy        uuid.UUID       `json:"resolvedBy"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *PaymentOverpaymentsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentOverpaymentsUpdateRequest) ToModel() model.PaymentOverpayments {
	return model.PaymentOverpayments{
		PaymentIntentId:   d.PaymentIntentId,
		PaidAttemptId:     nuuid.From(d.PaidAttemptId),
		OverpaidAttemptId: nuuid.From(d.OverpaidAttemptId),
		ExpectedAmount:    d.ExpectedAmount,
		ReceivedAmount:    d.ReceivedAmount,
		OverpaidAmount:    d.OverpaidAmount,
		Currency:          d.Currency,
		Status:            d.Status,
		ResolutionAction:  null.StringFrom(d.ResolutionAction),
		ResolutionNotes:   null.StringFrom(d.ResolutionNotes),
		ResolvedAt:        null.TimeFrom(d.ResolvedAt),
		ResolvedBy:        nuuid.From(d.ResolvedBy),
		Metadata:          d.Metadata,
	}
}

type PaymentOverpaymentsBulkUpdateRequest struct {
	Id                uuid.UUID       `json:"id"`
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId"`
	PaidAttemptId     uuid.UUID       `json:"paidAttemptId"`
	OverpaidAttemptId uuid.UUID       `json:"overpaidAttemptId"`
	ExpectedAmount    decimal.Decimal `json:"expectedAmount"`
	ReceivedAmount    decimal.Decimal `json:"receivedAmount"`
	OverpaidAmount    decimal.Decimal `json:"overpaidAmount"`
	Currency          string          `json:"currency"`
	Status            string          `json:"status"`
	ResolutionAction  string          `json:"resolutionAction"`
	ResolutionNotes   string          `json:"resolutionNotes"`
	ResolvedAt        time.Time       `json:"resolvedAt"`
	ResolvedBy        uuid.UUID       `json:"resolvedBy"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d PaymentOverpaymentsBulkUpdateRequest) PrimaryID() PaymentOverpaymentsPrimaryID {
	return PaymentOverpaymentsPrimaryID{
		Id: d.Id,
	}
}

type PaymentOverpaymentsListBulkUpdateRequest []*PaymentOverpaymentsBulkUpdateRequest

func (d PaymentOverpaymentsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentOverpayments := range d {
		err = validator.Struct(paymentOverpayments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentOverpaymentsBulkUpdateRequest) ToModel() model.PaymentOverpayments {
	return model.PaymentOverpayments{
		Id:                d.Id,
		PaymentIntentId:   d.PaymentIntentId,
		PaidAttemptId:     nuuid.From(d.PaidAttemptId),
		OverpaidAttemptId: nuuid.From(d.OverpaidAttemptId),
		ExpectedAmount:    d.ExpectedAmount,
		ReceivedAmount:    d.ReceivedAmount,
		OverpaidAmount:    d.OverpaidAmount,
		Currency:          d.Currency,
		Status:            d.Status,
		ResolutionAction:  null.StringFrom(d.ResolutionAction),
		ResolutionNotes:   null.StringFrom(d.ResolutionNotes),
		ResolvedAt:        null.TimeFrom(d.ResolvedAt),
		ResolvedBy:        nuuid.From(d.ResolvedBy),
		Metadata:          d.Metadata,
	}
}

type PaymentOverpaymentsResponse struct {
	Id                uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId   uuid.UUID       `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaidAttemptId     uuid.UUID       `json:"paidAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OverpaidAttemptId uuid.UUID       `json:"overpaidAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ExpectedAmount    decimal.Decimal `json:"expectedAmount" validate:"required" format:"decimal" example:"100.50"`
	ReceivedAmount    decimal.Decimal `json:"receivedAmount" validate:"required" format:"decimal" example:"100.50"`
	OverpaidAmount    decimal.Decimal `json:"overpaidAmount" validate:"required" format:"decimal" example:"100.50"`
	Currency          string          `json:"currency"`
	Status            string          `json:"status"`
	ResolutionAction  string          `json:"resolutionAction"`
	ResolutionNotes   string          `json:"resolutionNotes"`
	ResolvedAt        time.Time       `json:"resolvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ResolvedBy        uuid.UUID       `json:"resolvedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Metadata          json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewPaymentOverpaymentsResponse(paymentOverpayments model.PaymentOverpayments) PaymentOverpaymentsResponse {
	return PaymentOverpaymentsResponse{
		Id:                paymentOverpayments.Id,
		PaymentIntentId:   paymentOverpayments.PaymentIntentId,
		PaidAttemptId:     paymentOverpayments.PaidAttemptId.UUID,
		OverpaidAttemptId: paymentOverpayments.OverpaidAttemptId.UUID,
		ExpectedAmount:    paymentOverpayments.ExpectedAmount,
		ReceivedAmount:    paymentOverpayments.ReceivedAmount,
		OverpaidAmount:    paymentOverpayments.OverpaidAmount,
		Currency:          paymentOverpayments.Currency,
		Status:            paymentOverpayments.Status,
		ResolutionAction:  paymentOverpayments.ResolutionAction.String,
		ResolutionNotes:   paymentOverpayments.ResolutionNotes.String,
		ResolvedAt:        paymentOverpayments.ResolvedAt.Time,
		ResolvedBy:        paymentOverpayments.ResolvedBy.UUID,
		Metadata:          paymentOverpayments.Metadata,
	}
}

type PaymentOverpaymentsListResponse []*PaymentOverpaymentsResponse

func NewPaymentOverpaymentsListResponse(paymentOverpaymentsList model.PaymentOverpaymentsList) PaymentOverpaymentsListResponse {
	dtoPaymentOverpaymentsListResponse := PaymentOverpaymentsListResponse{}
	for _, paymentOverpayments := range paymentOverpaymentsList {
		dtoPaymentOverpaymentsResponse := NewPaymentOverpaymentsResponse(*paymentOverpayments)
		dtoPaymentOverpaymentsListResponse = append(dtoPaymentOverpaymentsListResponse, &dtoPaymentOverpaymentsResponse)
	}
	return dtoPaymentOverpaymentsListResponse
}

type PaymentOverpaymentsPrimaryIDList []PaymentOverpaymentsPrimaryID

func (d PaymentOverpaymentsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentOverpayments := range d {
		err = validator.Struct(paymentOverpayments)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentOverpaymentsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentOverpaymentsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentOverpaymentsPrimaryID) ToModel() model.PaymentOverpaymentsPrimaryID {
	return model.PaymentOverpaymentsPrimaryID{
		Id: d.Id,
	}
}
