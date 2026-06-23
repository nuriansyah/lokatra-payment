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

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ReconciliationExceptionsDTOFieldNameType string

type reconciliationExceptionsDTOFieldName struct {
	Id                  ReconciliationExceptionsDTOFieldNameType
	ReconciliationRunId ReconciliationExceptionsDTOFieldNameType
	ExceptionCode       ReconciliationExceptionsDTOFieldNameType
	ExceptionType       ReconciliationExceptionsDTOFieldNameType
	SourceRefType       ReconciliationExceptionsDTOFieldNameType
	SourceRefId         ReconciliationExceptionsDTOFieldNameType
	Severity            ReconciliationExceptionsDTOFieldNameType
	AmountDifference    ReconciliationExceptionsDTOFieldNameType
	ExceptionStatus     ReconciliationExceptionsDTOFieldNameType
	Resolution          ReconciliationExceptionsDTOFieldNameType
	ResolvedAt          ReconciliationExceptionsDTOFieldNameType
	Metadata            ReconciliationExceptionsDTOFieldNameType
	MetaCreatedAt       ReconciliationExceptionsDTOFieldNameType
	MetaCreatedBy       ReconciliationExceptionsDTOFieldNameType
	MetaUpdatedAt       ReconciliationExceptionsDTOFieldNameType
	MetaUpdatedBy       ReconciliationExceptionsDTOFieldNameType
	MetaDeletedAt       ReconciliationExceptionsDTOFieldNameType
	MetaDeletedBy       ReconciliationExceptionsDTOFieldNameType
}

var ReconciliationExceptionsDTOFieldName = reconciliationExceptionsDTOFieldName{
	Id:                  "id",
	ReconciliationRunId: "reconciliationRunId",
	ExceptionCode:       "exceptionCode",
	ExceptionType:       "exceptionType",
	SourceRefType:       "sourceRefType",
	SourceRefId:         "sourceRefId",
	Severity:            "severity",
	AmountDifference:    "amountDifference",
	ExceptionStatus:     "exceptionStatus",
	Resolution:          "resolution",
	ResolvedAt:          "resolvedAt",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformReconciliationExceptionsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ReconciliationExceptionsDTOFieldName.Id):
		return string(model.ReconciliationExceptionsDBFieldName.Id), true

	case string(ReconciliationExceptionsDTOFieldName.ReconciliationRunId):
		return string(model.ReconciliationExceptionsDBFieldName.ReconciliationRunId), true

	case string(ReconciliationExceptionsDTOFieldName.ExceptionCode):
		return string(model.ReconciliationExceptionsDBFieldName.ExceptionCode), true

	case string(ReconciliationExceptionsDTOFieldName.ExceptionType):
		return string(model.ReconciliationExceptionsDBFieldName.ExceptionType), true

	case string(ReconciliationExceptionsDTOFieldName.SourceRefType):
		return string(model.ReconciliationExceptionsDBFieldName.SourceRefType), true

	case string(ReconciliationExceptionsDTOFieldName.SourceRefId):
		return string(model.ReconciliationExceptionsDBFieldName.SourceRefId), true

	case string(ReconciliationExceptionsDTOFieldName.Severity):
		return string(model.ReconciliationExceptionsDBFieldName.Severity), true

	case string(ReconciliationExceptionsDTOFieldName.AmountDifference):
		return string(model.ReconciliationExceptionsDBFieldName.AmountDifference), true

	case string(ReconciliationExceptionsDTOFieldName.ExceptionStatus):
		return string(model.ReconciliationExceptionsDBFieldName.ExceptionStatus), true

	case string(ReconciliationExceptionsDTOFieldName.Resolution):
		return string(model.ReconciliationExceptionsDBFieldName.Resolution), true

	case string(ReconciliationExceptionsDTOFieldName.ResolvedAt):
		return string(model.ReconciliationExceptionsDBFieldName.ResolvedAt), true

	case string(ReconciliationExceptionsDTOFieldName.Metadata):
		return string(model.ReconciliationExceptionsDBFieldName.Metadata), true

	case string(ReconciliationExceptionsDTOFieldName.MetaCreatedAt):
		return string(model.ReconciliationExceptionsDBFieldName.MetaCreatedAt), true

	case string(ReconciliationExceptionsDTOFieldName.MetaCreatedBy):
		return string(model.ReconciliationExceptionsDBFieldName.MetaCreatedBy), true

	case string(ReconciliationExceptionsDTOFieldName.MetaUpdatedAt):
		return string(model.ReconciliationExceptionsDBFieldName.MetaUpdatedAt), true

	case string(ReconciliationExceptionsDTOFieldName.MetaUpdatedBy):
		return string(model.ReconciliationExceptionsDBFieldName.MetaUpdatedBy), true

	case string(ReconciliationExceptionsDTOFieldName.MetaDeletedAt):
		return string(model.ReconciliationExceptionsDBFieldName.MetaDeletedAt), true

	case string(ReconciliationExceptionsDTOFieldName.MetaDeletedBy):
		return string(model.ReconciliationExceptionsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewReconciliationExceptionsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isReconciliationExceptionsBaseFilterField(field string) bool {
	spec, found := model.NewReconciliationExceptionsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeReconciliationExceptionsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateReconciliationExceptionsProjectionOutputPath(path string) error {
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

func transformReconciliationExceptionsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformReconciliationExceptionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformReconciliationExceptionsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformReconciliationExceptionsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformReconciliationExceptionsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isReconciliationExceptionsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateReconciliationExceptionsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeReconciliationExceptionsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformReconciliationExceptionsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformReconciliationExceptionsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformReconciliationExceptionsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultReconciliationExceptionsFilter(filter *model.Filter) {
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
			Field: string(ReconciliationExceptionsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ReconciliationExceptionsSelectableResponse map[string]interface{}
type ReconciliationExceptionsSelectableListResponse []*ReconciliationExceptionsSelectableResponse

func assignReconciliationExceptionsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setReconciliationExceptionsSelectableValue(out ReconciliationExceptionsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignReconciliationExceptionsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewReconciliationExceptionsSelectableResponse(reconciliationExceptions model.ReconciliationExceptions, filter model.Filter) ReconciliationExceptionsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ReconciliationExceptionsDBFieldName.Id),
			string(model.ReconciliationExceptionsDBFieldName.ReconciliationRunId),
			string(model.ReconciliationExceptionsDBFieldName.ExceptionCode),
			string(model.ReconciliationExceptionsDBFieldName.ExceptionType),
			string(model.ReconciliationExceptionsDBFieldName.SourceRefType),
			string(model.ReconciliationExceptionsDBFieldName.SourceRefId),
			string(model.ReconciliationExceptionsDBFieldName.Severity),
			string(model.ReconciliationExceptionsDBFieldName.AmountDifference),
			string(model.ReconciliationExceptionsDBFieldName.ExceptionStatus),
			string(model.ReconciliationExceptionsDBFieldName.Resolution),
			string(model.ReconciliationExceptionsDBFieldName.ResolvedAt),
			string(model.ReconciliationExceptionsDBFieldName.Metadata),
			string(model.ReconciliationExceptionsDBFieldName.MetaCreatedAt),
			string(model.ReconciliationExceptionsDBFieldName.MetaCreatedBy),
			string(model.ReconciliationExceptionsDBFieldName.MetaUpdatedAt),
			string(model.ReconciliationExceptionsDBFieldName.MetaUpdatedBy),
			string(model.ReconciliationExceptionsDBFieldName.MetaDeletedAt),
			string(model.ReconciliationExceptionsDBFieldName.MetaDeletedBy),
		)
	}
	reconciliationExceptionsSelectableResponse := ReconciliationExceptionsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ReconciliationExceptionsDBFieldName.Id):
			key := string(ReconciliationExceptionsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.Id, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.ReconciliationRunId):
			key := string(ReconciliationExceptionsDTOFieldName.ReconciliationRunId)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.ReconciliationRunId, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.ExceptionCode):
			key := string(ReconciliationExceptionsDTOFieldName.ExceptionCode)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.ExceptionCode, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.ExceptionType):
			key := string(ReconciliationExceptionsDTOFieldName.ExceptionType)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, model.ExceptionType(reconciliationExceptions.ExceptionType), explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.SourceRefType):
			key := string(ReconciliationExceptionsDTOFieldName.SourceRefType)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.SourceRefType, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.SourceRefId):
			key := string(ReconciliationExceptionsDTOFieldName.SourceRefId)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.SourceRefId, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.Severity):
			key := string(ReconciliationExceptionsDTOFieldName.Severity)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, model.Severity(reconciliationExceptions.Severity), explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.AmountDifference):
			key := string(ReconciliationExceptionsDTOFieldName.AmountDifference)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.AmountDifference.Decimal, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.ExceptionStatus):
			key := string(ReconciliationExceptionsDTOFieldName.ExceptionStatus)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, model.ExceptionStatus(reconciliationExceptions.ExceptionStatus), explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.Resolution):
			key := string(ReconciliationExceptionsDTOFieldName.Resolution)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.Resolution.String, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.ResolvedAt):
			key := string(ReconciliationExceptionsDTOFieldName.ResolvedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.ResolvedAt.Time, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.Metadata):
			key := string(ReconciliationExceptionsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.Metadata, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.MetaCreatedAt):
			key := string(ReconciliationExceptionsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.MetaCreatedAt, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.MetaCreatedBy):
			key := string(ReconciliationExceptionsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.MetaCreatedBy, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.MetaUpdatedAt):
			key := string(ReconciliationExceptionsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.MetaUpdatedAt, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.MetaUpdatedBy):
			key := string(ReconciliationExceptionsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.MetaUpdatedBy, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.MetaDeletedAt):
			key := string(ReconciliationExceptionsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.MetaDeletedAt.Time, explicitAlias)

		case string(model.ReconciliationExceptionsDBFieldName.MetaDeletedBy):
			key := string(ReconciliationExceptionsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationExceptionsSelectableValue(reconciliationExceptionsSelectableResponse, key, reconciliationExceptions.MetaDeletedBy, explicitAlias)

		}
	}
	return reconciliationExceptionsSelectableResponse
}

func NewReconciliationExceptionsListResponseFromFilterResult(result []model.ReconciliationExceptionsFilterResult, filter model.Filter) ReconciliationExceptionsSelectableListResponse {
	dtoReconciliationExceptionsListResponse := ReconciliationExceptionsSelectableListResponse{}
	for _, row := range result {
		dtoReconciliationExceptionsResponse := NewReconciliationExceptionsSelectableResponse(row.ReconciliationExceptions, filter)
		dtoReconciliationExceptionsListResponse = append(dtoReconciliationExceptionsListResponse, &dtoReconciliationExceptionsResponse)
	}
	return dtoReconciliationExceptionsListResponse
}

type ReconciliationExceptionsFilterResponse struct {
	Metadata Metadata                                       `json:"metadata"`
	Data     ReconciliationExceptionsSelectableListResponse `json:"data"`
}

func reverseReconciliationExceptionsFilterResults(result []model.ReconciliationExceptionsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewReconciliationExceptionsFilterResponse(result []model.ReconciliationExceptionsFilterResult, filter model.Filter) (resp ReconciliationExceptionsFilterResponse) {
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
			reverseReconciliationExceptionsFilterResults(dataResult)
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

	resp.Data = NewReconciliationExceptionsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ReconciliationExceptionsCreateRequest struct {
	ReconciliationRunId uuid.UUID             `json:"reconciliationRunId"`
	ExceptionCode       string                `json:"exceptionCode"`
	ExceptionType       model.ExceptionType   `json:"exceptionType" example:"missing_provider" enums:"missing_provider,missing_bank,missing_ledger,amount_mismatch,duplicate,late_settlement,late_webhook,manual_review"`
	SourceRefType       string                `json:"sourceRefType"`
	SourceRefId         uuid.UUID             `json:"sourceRefId"`
	Severity            model.Severity        `json:"severity" example:"low" enums:"low,medium,high,critical"`
	AmountDifference    decimal.Decimal       `json:"amountDifference"`
	ExceptionStatus     model.ExceptionStatus `json:"exceptionStatus" example:"open" enums:"open,assigned,resolved,waived"`
	Resolution          string                `json:"resolution"`
	ResolvedAt          time.Time             `json:"resolvedAt"`
	Metadata            json.RawMessage       `json:"metadata"`
}

func (d *ReconciliationExceptionsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ReconciliationExceptionsCreateRequest) ToModel() model.ReconciliationExceptions {
	id, _ := uuid.NewV4()
	return model.ReconciliationExceptions{
		Id:                  id,
		ReconciliationRunId: d.ReconciliationRunId,
		ExceptionCode:       d.ExceptionCode,
		ExceptionType:       d.ExceptionType,
		SourceRefType:       d.SourceRefType,
		SourceRefId:         d.SourceRefId,
		Severity:            d.Severity,
		AmountDifference:    decimal.NewNullDecimal(d.AmountDifference),
		ExceptionStatus:     d.ExceptionStatus,
		Resolution:          null.StringFrom(d.Resolution),
		ResolvedAt:          null.TimeFrom(d.ResolvedAt),
		Metadata:            d.Metadata,
	}
}

type ReconciliationExceptionsListCreateRequest []*ReconciliationExceptionsCreateRequest

func (d ReconciliationExceptionsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationExceptions := range d {
		err = validator.Struct(reconciliationExceptions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationExceptionsListCreateRequest) ToModelList() []model.ReconciliationExceptions {
	out := make([]model.ReconciliationExceptions, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ReconciliationExceptionsUpdateRequest struct {
	ReconciliationRunId uuid.UUID             `json:"reconciliationRunId"`
	ExceptionCode       string                `json:"exceptionCode"`
	ExceptionType       model.ExceptionType   `json:"exceptionType" example:"missing_provider" enums:"missing_provider,missing_bank,missing_ledger,amount_mismatch,duplicate,late_settlement,late_webhook,manual_review"`
	SourceRefType       string                `json:"sourceRefType"`
	SourceRefId         uuid.UUID             `json:"sourceRefId"`
	Severity            model.Severity        `json:"severity" example:"low" enums:"low,medium,high,critical"`
	AmountDifference    decimal.Decimal       `json:"amountDifference"`
	ExceptionStatus     model.ExceptionStatus `json:"exceptionStatus" example:"open" enums:"open,assigned,resolved,waived"`
	Resolution          string                `json:"resolution"`
	ResolvedAt          time.Time             `json:"resolvedAt"`
	Metadata            json.RawMessage       `json:"metadata"`
}

func (d *ReconciliationExceptionsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ReconciliationExceptionsUpdateRequest) ToModel() model.ReconciliationExceptions {
	return model.ReconciliationExceptions{
		ReconciliationRunId: d.ReconciliationRunId,
		ExceptionCode:       d.ExceptionCode,
		ExceptionType:       d.ExceptionType,
		SourceRefType:       d.SourceRefType,
		SourceRefId:         d.SourceRefId,
		Severity:            d.Severity,
		AmountDifference:    decimal.NewNullDecimal(d.AmountDifference),
		ExceptionStatus:     d.ExceptionStatus,
		Resolution:          null.StringFrom(d.Resolution),
		ResolvedAt:          null.TimeFrom(d.ResolvedAt),
		Metadata:            d.Metadata,
	}
}

type ReconciliationExceptionsBulkUpdateRequest struct {
	Id                  uuid.UUID             `json:"id"`
	ReconciliationRunId uuid.UUID             `json:"reconciliationRunId"`
	ExceptionCode       string                `json:"exceptionCode"`
	ExceptionType       model.ExceptionType   `json:"exceptionType" example:"missing_provider" enums:"missing_provider,missing_bank,missing_ledger,amount_mismatch,duplicate,late_settlement,late_webhook,manual_review"`
	SourceRefType       string                `json:"sourceRefType"`
	SourceRefId         uuid.UUID             `json:"sourceRefId"`
	Severity            model.Severity        `json:"severity" example:"low" enums:"low,medium,high,critical"`
	AmountDifference    decimal.Decimal       `json:"amountDifference"`
	ExceptionStatus     model.ExceptionStatus `json:"exceptionStatus" example:"open" enums:"open,assigned,resolved,waived"`
	Resolution          string                `json:"resolution"`
	ResolvedAt          time.Time             `json:"resolvedAt"`
	Metadata            json.RawMessage       `json:"metadata"`
}

func (d ReconciliationExceptionsBulkUpdateRequest) PrimaryID() ReconciliationExceptionsPrimaryID {
	return ReconciliationExceptionsPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationExceptionsListBulkUpdateRequest []*ReconciliationExceptionsBulkUpdateRequest

func (d ReconciliationExceptionsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationExceptions := range d {
		err = validator.Struct(reconciliationExceptions)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationExceptionsBulkUpdateRequest) ToModel() model.ReconciliationExceptions {
	return model.ReconciliationExceptions{
		Id:                  d.Id,
		ReconciliationRunId: d.ReconciliationRunId,
		ExceptionCode:       d.ExceptionCode,
		ExceptionType:       d.ExceptionType,
		SourceRefType:       d.SourceRefType,
		SourceRefId:         d.SourceRefId,
		Severity:            d.Severity,
		AmountDifference:    decimal.NewNullDecimal(d.AmountDifference),
		ExceptionStatus:     d.ExceptionStatus,
		Resolution:          null.StringFrom(d.Resolution),
		ResolvedAt:          null.TimeFrom(d.ResolvedAt),
		Metadata:            d.Metadata,
	}
}

type ReconciliationExceptionsResponse struct {
	Id                  uuid.UUID             `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReconciliationRunId uuid.UUID             `json:"reconciliationRunId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ExceptionCode       string                `json:"exceptionCode" validate:"required"`
	ExceptionType       model.ExceptionType   `json:"exceptionType" validate:"required,oneof=missing_provider missing_bank missing_ledger amount_mismatch duplicate late_settlement late_webhook manual_review" enums:"missing_provider,missing_bank,missing_ledger,amount_mismatch,duplicate,late_settlement,late_webhook,manual_review"`
	SourceRefType       string                `json:"sourceRefType" validate:"required"`
	SourceRefId         uuid.UUID             `json:"sourceRefId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Severity            model.Severity        `json:"severity" validate:"required,oneof=low medium high critical" enums:"low,medium,high,critical"`
	AmountDifference    decimal.Decimal       `json:"amountDifference" format:"decimal" example:"100.50"`
	ExceptionStatus     model.ExceptionStatus `json:"exceptionStatus" validate:"oneof=open assigned resolved waived" enums:"open,assigned,resolved,waived"`
	Resolution          string                `json:"resolution"`
	ResolvedAt          time.Time             `json:"resolvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata            json.RawMessage       `json:"metadata" swaggertype:"object"`
}

func NewReconciliationExceptionsResponse(reconciliationExceptions model.ReconciliationExceptions) ReconciliationExceptionsResponse {
	return ReconciliationExceptionsResponse{
		Id:                  reconciliationExceptions.Id,
		ReconciliationRunId: reconciliationExceptions.ReconciliationRunId,
		ExceptionCode:       reconciliationExceptions.ExceptionCode,
		ExceptionType:       model.ExceptionType(reconciliationExceptions.ExceptionType),
		SourceRefType:       reconciliationExceptions.SourceRefType,
		SourceRefId:         reconciliationExceptions.SourceRefId,
		Severity:            model.Severity(reconciliationExceptions.Severity),
		AmountDifference:    reconciliationExceptions.AmountDifference.Decimal,
		ExceptionStatus:     model.ExceptionStatus(reconciliationExceptions.ExceptionStatus),
		Resolution:          reconciliationExceptions.Resolution.String,
		ResolvedAt:          reconciliationExceptions.ResolvedAt.Time,
		Metadata:            reconciliationExceptions.Metadata,
	}
}

type ReconciliationExceptionsListResponse []*ReconciliationExceptionsResponse

func NewReconciliationExceptionsListResponse(reconciliationExceptionsList model.ReconciliationExceptionsList) ReconciliationExceptionsListResponse {
	dtoReconciliationExceptionsListResponse := ReconciliationExceptionsListResponse{}
	for _, reconciliationExceptions := range reconciliationExceptionsList {
		dtoReconciliationExceptionsResponse := NewReconciliationExceptionsResponse(*reconciliationExceptions)
		dtoReconciliationExceptionsListResponse = append(dtoReconciliationExceptionsListResponse, &dtoReconciliationExceptionsResponse)
	}
	return dtoReconciliationExceptionsListResponse
}

type ReconciliationExceptionsPrimaryIDList []ReconciliationExceptionsPrimaryID

func (d ReconciliationExceptionsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationExceptions := range d {
		err = validator.Struct(reconciliationExceptions)
		if err != nil {
			return
		}
	}
	return nil
}

type ReconciliationExceptionsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ReconciliationExceptionsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ReconciliationExceptionsPrimaryID) ToModel() model.ReconciliationExceptionsPrimaryID {
	return model.ReconciliationExceptionsPrimaryID{
		Id: d.Id,
	}
}
