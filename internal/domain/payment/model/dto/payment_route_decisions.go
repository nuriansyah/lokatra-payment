package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentRouteDecisionsDTOFieldNameType string

type paymentRouteDecisionsDTOFieldName struct {
	Id                        PaymentRouteDecisionsDTOFieldNameType
	PaymentIntentId           PaymentRouteDecisionsDTOFieldNameType
	SelectedProviderAccountId PaymentRouteDecisionsDTOFieldNameType
	SelectedProviderCode      PaymentRouteDecisionsDTOFieldNameType
	MethodCode                PaymentRouteDecisionsDTOFieldNameType
	ChannelCode               PaymentRouteDecisionsDTOFieldNameType
	Reason                    PaymentRouteDecisionsDTOFieldNameType
	EvaluatedContext          PaymentRouteDecisionsDTOFieldNameType
	Candidates                PaymentRouteDecisionsDTOFieldNameType
	Metadata                  PaymentRouteDecisionsDTOFieldNameType
	MetaCreatedAt             PaymentRouteDecisionsDTOFieldNameType
	MetaCreatedBy             PaymentRouteDecisionsDTOFieldNameType
	MetaUpdatedAt             PaymentRouteDecisionsDTOFieldNameType
	MetaUpdatedBy             PaymentRouteDecisionsDTOFieldNameType
	MetaDeletedAt             PaymentRouteDecisionsDTOFieldNameType
	MetaDeletedBy             PaymentRouteDecisionsDTOFieldNameType
}

var PaymentRouteDecisionsDTOFieldName = paymentRouteDecisionsDTOFieldName{
	Id:                        "id",
	PaymentIntentId:           "paymentIntentId",
	SelectedProviderAccountId: "selectedProviderAccountId",
	SelectedProviderCode:      "selectedProviderCode",
	MethodCode:                "methodCode",
	ChannelCode:               "channelCode",
	Reason:                    "reason",
	EvaluatedContext:          "evaluatedContext",
	Candidates:                "candidates",
	Metadata:                  "metadata",
	MetaCreatedAt:             "metaCreatedAt",
	MetaCreatedBy:             "metaCreatedBy",
	MetaUpdatedAt:             "metaUpdatedAt",
	MetaUpdatedBy:             "metaUpdatedBy",
	MetaDeletedAt:             "metaDeletedAt",
	MetaDeletedBy:             "metaDeletedBy",
}

func transformPaymentRouteDecisionsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(PaymentRouteDecisionsDTOFieldName.Id):
		return string(model.PaymentRouteDecisionsDBFieldName.Id), true

	case string(PaymentRouteDecisionsDTOFieldName.PaymentIntentId):
		return string(model.PaymentRouteDecisionsDBFieldName.PaymentIntentId), true

	case string(PaymentRouteDecisionsDTOFieldName.SelectedProviderAccountId):
		return string(model.PaymentRouteDecisionsDBFieldName.SelectedProviderAccountId), true

	case string(PaymentRouteDecisionsDTOFieldName.SelectedProviderCode):
		return string(model.PaymentRouteDecisionsDBFieldName.SelectedProviderCode), true

	case string(PaymentRouteDecisionsDTOFieldName.MethodCode):
		return string(model.PaymentRouteDecisionsDBFieldName.MethodCode), true

	case string(PaymentRouteDecisionsDTOFieldName.ChannelCode):
		return string(model.PaymentRouteDecisionsDBFieldName.ChannelCode), true

	case string(PaymentRouteDecisionsDTOFieldName.Reason):
		return string(model.PaymentRouteDecisionsDBFieldName.Reason), true

	case string(PaymentRouteDecisionsDTOFieldName.EvaluatedContext):
		return string(model.PaymentRouteDecisionsDBFieldName.EvaluatedContext), true

	case string(PaymentRouteDecisionsDTOFieldName.Candidates):
		return string(model.PaymentRouteDecisionsDBFieldName.Candidates), true

	case string(PaymentRouteDecisionsDTOFieldName.Metadata):
		return string(model.PaymentRouteDecisionsDBFieldName.Metadata), true

	case string(PaymentRouteDecisionsDTOFieldName.MetaCreatedAt):
		return string(model.PaymentRouteDecisionsDBFieldName.MetaCreatedAt), true

	case string(PaymentRouteDecisionsDTOFieldName.MetaCreatedBy):
		return string(model.PaymentRouteDecisionsDBFieldName.MetaCreatedBy), true

	case string(PaymentRouteDecisionsDTOFieldName.MetaUpdatedAt):
		return string(model.PaymentRouteDecisionsDBFieldName.MetaUpdatedAt), true

	case string(PaymentRouteDecisionsDTOFieldName.MetaUpdatedBy):
		return string(model.PaymentRouteDecisionsDBFieldName.MetaUpdatedBy), true

	case string(PaymentRouteDecisionsDTOFieldName.MetaDeletedAt):
		return string(model.PaymentRouteDecisionsDBFieldName.MetaDeletedAt), true

	case string(PaymentRouteDecisionsDTOFieldName.MetaDeletedBy):
		return string(model.PaymentRouteDecisionsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewPaymentRouteDecisionsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isPaymentRouteDecisionsBaseFilterField(field string) bool {
	spec, found := model.NewPaymentRouteDecisionsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composePaymentRouteDecisionsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validatePaymentRouteDecisionsProjectionOutputPath(path string) error {
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

func transformPaymentRouteDecisionsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformPaymentRouteDecisionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformPaymentRouteDecisionsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformPaymentRouteDecisionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformPaymentRouteDecisionsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isPaymentRouteDecisionsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validatePaymentRouteDecisionsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composePaymentRouteDecisionsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPaymentRouteDecisionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPaymentRouteDecisionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformPaymentRouteDecisionsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultPaymentRouteDecisionsFilter(filter *model.Filter) {
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
			Field: string(PaymentRouteDecisionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PaymentRouteDecisionsSelectableResponse map[string]interface{}
type PaymentRouteDecisionsSelectableListResponse []*PaymentRouteDecisionsSelectableResponse

func assignPaymentRouteDecisionsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setPaymentRouteDecisionsSelectableValue(out PaymentRouteDecisionsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignPaymentRouteDecisionsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewPaymentRouteDecisionsSelectableResponse(paymentRouteDecisions model.PaymentRouteDecisions, filter model.Filter) PaymentRouteDecisionsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PaymentRouteDecisionsDBFieldName.Id),
			string(model.PaymentRouteDecisionsDBFieldName.PaymentIntentId),
			string(model.PaymentRouteDecisionsDBFieldName.SelectedProviderAccountId),
			string(model.PaymentRouteDecisionsDBFieldName.SelectedProviderCode),
			string(model.PaymentRouteDecisionsDBFieldName.MethodCode),
			string(model.PaymentRouteDecisionsDBFieldName.ChannelCode),
			string(model.PaymentRouteDecisionsDBFieldName.Reason),
			string(model.PaymentRouteDecisionsDBFieldName.EvaluatedContext),
			string(model.PaymentRouteDecisionsDBFieldName.Candidates),
			string(model.PaymentRouteDecisionsDBFieldName.Metadata),
			string(model.PaymentRouteDecisionsDBFieldName.MetaCreatedAt),
			string(model.PaymentRouteDecisionsDBFieldName.MetaCreatedBy),
			string(model.PaymentRouteDecisionsDBFieldName.MetaUpdatedAt),
			string(model.PaymentRouteDecisionsDBFieldName.MetaUpdatedBy),
			string(model.PaymentRouteDecisionsDBFieldName.MetaDeletedAt),
			string(model.PaymentRouteDecisionsDBFieldName.MetaDeletedBy),
		)
	}
	paymentRouteDecisionsSelectableResponse := PaymentRouteDecisionsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.PaymentRouteDecisionsDBFieldName.Id):
			key := string(PaymentRouteDecisionsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.Id, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.PaymentIntentId):
			key := string(PaymentRouteDecisionsDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.PaymentIntentId, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.SelectedProviderAccountId):
			key := string(PaymentRouteDecisionsDTOFieldName.SelectedProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.SelectedProviderAccountId.UUID, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.SelectedProviderCode):
			key := string(PaymentRouteDecisionsDTOFieldName.SelectedProviderCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.SelectedProviderCode.String, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.MethodCode):
			key := string(PaymentRouteDecisionsDTOFieldName.MethodCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.MethodCode, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.ChannelCode):
			key := string(PaymentRouteDecisionsDTOFieldName.ChannelCode)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.ChannelCode.String, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.Reason):
			key := string(PaymentRouteDecisionsDTOFieldName.Reason)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.Reason, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.EvaluatedContext):
			key := string(PaymentRouteDecisionsDTOFieldName.EvaluatedContext)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.EvaluatedContext, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.Candidates):
			key := string(PaymentRouteDecisionsDTOFieldName.Candidates)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.Candidates, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.Metadata):
			key := string(PaymentRouteDecisionsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.Metadata, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.MetaCreatedAt):
			key := string(PaymentRouteDecisionsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.MetaCreatedAt, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.MetaCreatedBy):
			key := string(PaymentRouteDecisionsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.MetaCreatedBy, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.MetaUpdatedAt):
			key := string(PaymentRouteDecisionsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.MetaUpdatedAt, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.MetaUpdatedBy):
			key := string(PaymentRouteDecisionsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.MetaDeletedAt):
			key := string(PaymentRouteDecisionsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.MetaDeletedAt.Time, explicitAlias)

		case string(model.PaymentRouteDecisionsDBFieldName.MetaDeletedBy):
			key := string(PaymentRouteDecisionsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setPaymentRouteDecisionsSelectableValue(paymentRouteDecisionsSelectableResponse, key, paymentRouteDecisions.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return paymentRouteDecisionsSelectableResponse
}

func NewPaymentRouteDecisionsListResponseFromFilterResult(result []model.PaymentRouteDecisionsFilterResult, filter model.Filter) PaymentRouteDecisionsSelectableListResponse {
	dtoPaymentRouteDecisionsListResponse := PaymentRouteDecisionsSelectableListResponse{}
	for _, row := range result {
		dtoPaymentRouteDecisionsResponse := NewPaymentRouteDecisionsSelectableResponse(row.PaymentRouteDecisions, filter)
		dtoPaymentRouteDecisionsListResponse = append(dtoPaymentRouteDecisionsListResponse, &dtoPaymentRouteDecisionsResponse)
	}
	return dtoPaymentRouteDecisionsListResponse
}

type PaymentRouteDecisionsFilterResponse struct {
	Metadata Metadata                                    `json:"metadata"`
	Data     PaymentRouteDecisionsSelectableListResponse `json:"data"`
}

func reversePaymentRouteDecisionsFilterResults(result []model.PaymentRouteDecisionsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewPaymentRouteDecisionsFilterResponse(result []model.PaymentRouteDecisionsFilterResult, filter model.Filter) (resp PaymentRouteDecisionsFilterResponse) {
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
			reversePaymentRouteDecisionsFilterResults(dataResult)
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

	resp.Data = NewPaymentRouteDecisionsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type PaymentRouteDecisionsCreateRequest struct {
	PaymentIntentId           uuid.UUID       `json:"paymentIntentId"`
	SelectedProviderAccountId uuid.UUID       `json:"selectedProviderAccountId"`
	SelectedProviderCode      string          `json:"selectedProviderCode"`
	MethodCode                string          `json:"methodCode"`
	ChannelCode               string          `json:"channelCode"`
	Reason                    string          `json:"reason"`
	EvaluatedContext          json.RawMessage `json:"evaluatedContext"`
	Candidates                json.RawMessage `json:"candidates"`
	Metadata                  json.RawMessage `json:"metadata"`
}

func (d *PaymentRouteDecisionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PaymentRouteDecisionsCreateRequest) ToModel() model.PaymentRouteDecisions {
	id, _ := uuid.NewV4()
	return model.PaymentRouteDecisions{
		Id:                        id,
		PaymentIntentId:           d.PaymentIntentId,
		SelectedProviderAccountId: nuuid.From(d.SelectedProviderAccountId),
		SelectedProviderCode:      null.StringFrom(d.SelectedProviderCode),
		MethodCode:                d.MethodCode,
		ChannelCode:               null.StringFrom(d.ChannelCode),
		Reason:                    d.Reason,
		EvaluatedContext:          d.EvaluatedContext,
		Candidates:                d.Candidates,
		Metadata:                  d.Metadata,
	}
}

type PaymentRouteDecisionsListCreateRequest []*PaymentRouteDecisionsCreateRequest

func (d PaymentRouteDecisionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRouteDecisions := range d {
		err = validator.Struct(paymentRouteDecisions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentRouteDecisionsListCreateRequest) ToModelList() []model.PaymentRouteDecisions {
	out := make([]model.PaymentRouteDecisions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PaymentRouteDecisionsUpdateRequest struct {
	PaymentIntentId           uuid.UUID       `json:"paymentIntentId"`
	SelectedProviderAccountId uuid.UUID       `json:"selectedProviderAccountId"`
	SelectedProviderCode      string          `json:"selectedProviderCode"`
	MethodCode                string          `json:"methodCode"`
	ChannelCode               string          `json:"channelCode"`
	Reason                    string          `json:"reason"`
	EvaluatedContext          json.RawMessage `json:"evaluatedContext"`
	Candidates                json.RawMessage `json:"candidates"`
	Metadata                  json.RawMessage `json:"metadata"`
}

func (d *PaymentRouteDecisionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PaymentRouteDecisionsUpdateRequest) ToModel() model.PaymentRouteDecisions {
	return model.PaymentRouteDecisions{
		PaymentIntentId:           d.PaymentIntentId,
		SelectedProviderAccountId: nuuid.From(d.SelectedProviderAccountId),
		SelectedProviderCode:      null.StringFrom(d.SelectedProviderCode),
		MethodCode:                d.MethodCode,
		ChannelCode:               null.StringFrom(d.ChannelCode),
		Reason:                    d.Reason,
		EvaluatedContext:          d.EvaluatedContext,
		Candidates:                d.Candidates,
		Metadata:                  d.Metadata,
	}
}

type PaymentRouteDecisionsBulkUpdateRequest struct {
	Id                        uuid.UUID       `json:"id"`
	PaymentIntentId           uuid.UUID       `json:"paymentIntentId"`
	SelectedProviderAccountId uuid.UUID       `json:"selectedProviderAccountId"`
	SelectedProviderCode      string          `json:"selectedProviderCode"`
	MethodCode                string          `json:"methodCode"`
	ChannelCode               string          `json:"channelCode"`
	Reason                    string          `json:"reason"`
	EvaluatedContext          json.RawMessage `json:"evaluatedContext"`
	Candidates                json.RawMessage `json:"candidates"`
	Metadata                  json.RawMessage `json:"metadata"`
}

func (d PaymentRouteDecisionsBulkUpdateRequest) PrimaryID() PaymentRouteDecisionsPrimaryID {
	return PaymentRouteDecisionsPrimaryID{
		Id: d.Id,
	}
}

type PaymentRouteDecisionsListBulkUpdateRequest []*PaymentRouteDecisionsBulkUpdateRequest

func (d PaymentRouteDecisionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRouteDecisions := range d {
		err = validator.Struct(paymentRouteDecisions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PaymentRouteDecisionsBulkUpdateRequest) ToModel() model.PaymentRouteDecisions {
	return model.PaymentRouteDecisions{
		Id:                        d.Id,
		PaymentIntentId:           d.PaymentIntentId,
		SelectedProviderAccountId: nuuid.From(d.SelectedProviderAccountId),
		SelectedProviderCode:      null.StringFrom(d.SelectedProviderCode),
		MethodCode:                d.MethodCode,
		ChannelCode:               null.StringFrom(d.ChannelCode),
		Reason:                    d.Reason,
		EvaluatedContext:          d.EvaluatedContext,
		Candidates:                d.Candidates,
		Metadata:                  d.Metadata,
	}
}

type PaymentRouteDecisionsResponse struct {
	Id                        uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId           uuid.UUID       `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SelectedProviderAccountId uuid.UUID       `json:"selectedProviderAccountId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SelectedProviderCode      string          `json:"selectedProviderCode"`
	MethodCode                string          `json:"methodCode" validate:"required"`
	ChannelCode               string          `json:"channelCode"`
	Reason                    string          `json:"reason" validate:"required"`
	EvaluatedContext          json.RawMessage `json:"evaluatedContext" swaggertype:"object"`
	Candidates                json.RawMessage `json:"candidates" swaggertype:"object"`
	Metadata                  json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewPaymentRouteDecisionsResponse(paymentRouteDecisions model.PaymentRouteDecisions) PaymentRouteDecisionsResponse {
	return PaymentRouteDecisionsResponse{
		Id:                        paymentRouteDecisions.Id,
		PaymentIntentId:           paymentRouteDecisions.PaymentIntentId,
		SelectedProviderAccountId: paymentRouteDecisions.SelectedProviderAccountId.UUID,
		SelectedProviderCode:      paymentRouteDecisions.SelectedProviderCode.String,
		MethodCode:                paymentRouteDecisions.MethodCode,
		ChannelCode:               paymentRouteDecisions.ChannelCode.String,
		Reason:                    paymentRouteDecisions.Reason,
		EvaluatedContext:          paymentRouteDecisions.EvaluatedContext,
		Candidates:                paymentRouteDecisions.Candidates,
		Metadata:                  paymentRouteDecisions.Metadata,
	}
}

type PaymentRouteDecisionsListResponse []*PaymentRouteDecisionsResponse

func NewPaymentRouteDecisionsListResponse(paymentRouteDecisionsList model.PaymentRouteDecisionsList) PaymentRouteDecisionsListResponse {
	dtoPaymentRouteDecisionsListResponse := PaymentRouteDecisionsListResponse{}
	for _, paymentRouteDecisions := range paymentRouteDecisionsList {
		dtoPaymentRouteDecisionsResponse := NewPaymentRouteDecisionsResponse(*paymentRouteDecisions)
		dtoPaymentRouteDecisionsListResponse = append(dtoPaymentRouteDecisionsListResponse, &dtoPaymentRouteDecisionsResponse)
	}
	return dtoPaymentRouteDecisionsListResponse
}

type PaymentRouteDecisionsPrimaryIDList []PaymentRouteDecisionsPrimaryID

func (d PaymentRouteDecisionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, paymentRouteDecisions := range d {
		err = validator.Struct(paymentRouteDecisions)
		if err != nil {
			return
		}
	}
	return nil
}

type PaymentRouteDecisionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PaymentRouteDecisionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PaymentRouteDecisionsPrimaryID) ToModel() model.PaymentRouteDecisionsPrimaryID {
	return model.PaymentRouteDecisionsPrimaryID{
		Id: d.Id,
	}
}
