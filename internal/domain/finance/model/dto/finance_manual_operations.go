package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceManualOperationsDTOFieldNameType string

type financeManualOperationsDTOFieldName struct {
	Id              FinanceManualOperationsDTOFieldNameType
	OperationCode   FinanceManualOperationsDTOFieldNameType
	OperationType   FinanceManualOperationsDTOFieldNameType
	TargetRefType   FinanceManualOperationsDTOFieldNameType
	TargetRefId     FinanceManualOperationsDTOFieldNameType
	RequestedBy     FinanceManualOperationsDTOFieldNameType
	OperationStatus FinanceManualOperationsDTOFieldNameType
	ReasonCode      FinanceManualOperationsDTOFieldNameType
	ReasonDetail    FinanceManualOperationsDTOFieldNameType
	Payload         FinanceManualOperationsDTOFieldNameType
	ExecutedAt      FinanceManualOperationsDTOFieldNameType
	Metadata        FinanceManualOperationsDTOFieldNameType
	MetaCreatedAt   FinanceManualOperationsDTOFieldNameType
	MetaCreatedBy   FinanceManualOperationsDTOFieldNameType
	MetaUpdatedAt   FinanceManualOperationsDTOFieldNameType
	MetaUpdatedBy   FinanceManualOperationsDTOFieldNameType
	MetaDeletedAt   FinanceManualOperationsDTOFieldNameType
	MetaDeletedBy   FinanceManualOperationsDTOFieldNameType
}

var FinanceManualOperationsDTOFieldName = financeManualOperationsDTOFieldName{
	Id:              "id",
	OperationCode:   "operationCode",
	OperationType:   "operationType",
	TargetRefType:   "targetRefType",
	TargetRefId:     "targetRefId",
	RequestedBy:     "requestedBy",
	OperationStatus: "operationStatus",
	ReasonCode:      "reasonCode",
	ReasonDetail:    "reasonDetail",
	Payload:         "payload",
	ExecutedAt:      "executedAt",
	Metadata:        "metadata",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func transformFinanceManualOperationsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(FinanceManualOperationsDTOFieldName.Id):
		return string(model.FinanceManualOperationsDBFieldName.Id), true

	case string(FinanceManualOperationsDTOFieldName.OperationCode):
		return string(model.FinanceManualOperationsDBFieldName.OperationCode), true

	case string(FinanceManualOperationsDTOFieldName.OperationType):
		return string(model.FinanceManualOperationsDBFieldName.OperationType), true

	case string(FinanceManualOperationsDTOFieldName.TargetRefType):
		return string(model.FinanceManualOperationsDBFieldName.TargetRefType), true

	case string(FinanceManualOperationsDTOFieldName.TargetRefId):
		return string(model.FinanceManualOperationsDBFieldName.TargetRefId), true

	case string(FinanceManualOperationsDTOFieldName.RequestedBy):
		return string(model.FinanceManualOperationsDBFieldName.RequestedBy), true

	case string(FinanceManualOperationsDTOFieldName.OperationStatus):
		return string(model.FinanceManualOperationsDBFieldName.OperationStatus), true

	case string(FinanceManualOperationsDTOFieldName.ReasonCode):
		return string(model.FinanceManualOperationsDBFieldName.ReasonCode), true

	case string(FinanceManualOperationsDTOFieldName.ReasonDetail):
		return string(model.FinanceManualOperationsDBFieldName.ReasonDetail), true

	case string(FinanceManualOperationsDTOFieldName.Payload):
		return string(model.FinanceManualOperationsDBFieldName.Payload), true

	case string(FinanceManualOperationsDTOFieldName.ExecutedAt):
		return string(model.FinanceManualOperationsDBFieldName.ExecutedAt), true

	case string(FinanceManualOperationsDTOFieldName.Metadata):
		return string(model.FinanceManualOperationsDBFieldName.Metadata), true

	case string(FinanceManualOperationsDTOFieldName.MetaCreatedAt):
		return string(model.FinanceManualOperationsDBFieldName.MetaCreatedAt), true

	case string(FinanceManualOperationsDTOFieldName.MetaCreatedBy):
		return string(model.FinanceManualOperationsDBFieldName.MetaCreatedBy), true

	case string(FinanceManualOperationsDTOFieldName.MetaUpdatedAt):
		return string(model.FinanceManualOperationsDBFieldName.MetaUpdatedAt), true

	case string(FinanceManualOperationsDTOFieldName.MetaUpdatedBy):
		return string(model.FinanceManualOperationsDBFieldName.MetaUpdatedBy), true

	case string(FinanceManualOperationsDTOFieldName.MetaDeletedAt):
		return string(model.FinanceManualOperationsDBFieldName.MetaDeletedAt), true

	case string(FinanceManualOperationsDTOFieldName.MetaDeletedBy):
		return string(model.FinanceManualOperationsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewFinanceManualOperationsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isFinanceManualOperationsBaseFilterField(field string) bool {
	spec, found := model.NewFinanceManualOperationsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeFinanceManualOperationsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateFinanceManualOperationsProjectionOutputPath(path string) error {
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

func transformFinanceManualOperationsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformFinanceManualOperationsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformFinanceManualOperationsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformFinanceManualOperationsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformFinanceManualOperationsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isFinanceManualOperationsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateFinanceManualOperationsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeFinanceManualOperationsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformFinanceManualOperationsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformFinanceManualOperationsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformFinanceManualOperationsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultFinanceManualOperationsFilter(filter *model.Filter) {
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
			Field: string(FinanceManualOperationsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type FinanceManualOperationsSelectableResponse map[string]interface{}
type FinanceManualOperationsSelectableListResponse []*FinanceManualOperationsSelectableResponse

func assignFinanceManualOperationsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setFinanceManualOperationsSelectableValue(out FinanceManualOperationsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignFinanceManualOperationsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewFinanceManualOperationsSelectableResponse(financeManualOperations model.FinanceManualOperations, filter model.Filter) FinanceManualOperationsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.FinanceManualOperationsDBFieldName.Id),
			string(model.FinanceManualOperationsDBFieldName.OperationCode),
			string(model.FinanceManualOperationsDBFieldName.OperationType),
			string(model.FinanceManualOperationsDBFieldName.TargetRefType),
			string(model.FinanceManualOperationsDBFieldName.TargetRefId),
			string(model.FinanceManualOperationsDBFieldName.RequestedBy),
			string(model.FinanceManualOperationsDBFieldName.OperationStatus),
			string(model.FinanceManualOperationsDBFieldName.ReasonCode),
			string(model.FinanceManualOperationsDBFieldName.ReasonDetail),
			string(model.FinanceManualOperationsDBFieldName.Payload),
			string(model.FinanceManualOperationsDBFieldName.ExecutedAt),
			string(model.FinanceManualOperationsDBFieldName.Metadata),
			string(model.FinanceManualOperationsDBFieldName.MetaCreatedAt),
			string(model.FinanceManualOperationsDBFieldName.MetaCreatedBy),
			string(model.FinanceManualOperationsDBFieldName.MetaUpdatedAt),
			string(model.FinanceManualOperationsDBFieldName.MetaUpdatedBy),
			string(model.FinanceManualOperationsDBFieldName.MetaDeletedAt),
			string(model.FinanceManualOperationsDBFieldName.MetaDeletedBy),
		)
	}
	financeManualOperationsSelectableResponse := FinanceManualOperationsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.FinanceManualOperationsDBFieldName.Id):
			key := string(FinanceManualOperationsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.Id, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.OperationCode):
			key := string(FinanceManualOperationsDTOFieldName.OperationCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.OperationCode, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.OperationType):
			key := string(FinanceManualOperationsDTOFieldName.OperationType)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, model.OperationType(financeManualOperations.OperationType), explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.TargetRefType):
			key := string(FinanceManualOperationsDTOFieldName.TargetRefType)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.TargetRefType, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.TargetRefId):
			key := string(FinanceManualOperationsDTOFieldName.TargetRefId)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.TargetRefId, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.RequestedBy):
			key := string(FinanceManualOperationsDTOFieldName.RequestedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.RequestedBy, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.OperationStatus):
			key := string(FinanceManualOperationsDTOFieldName.OperationStatus)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, model.OperationStatus(financeManualOperations.OperationStatus), explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.ReasonCode):
			key := string(FinanceManualOperationsDTOFieldName.ReasonCode)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.ReasonCode, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.ReasonDetail):
			key := string(FinanceManualOperationsDTOFieldName.ReasonDetail)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.ReasonDetail.String, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.Payload):
			key := string(FinanceManualOperationsDTOFieldName.Payload)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.Payload, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.ExecutedAt):
			key := string(FinanceManualOperationsDTOFieldName.ExecutedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.ExecutedAt.Time, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.Metadata):
			key := string(FinanceManualOperationsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.Metadata, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.MetaCreatedAt):
			key := string(FinanceManualOperationsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.MetaCreatedAt, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.MetaCreatedBy):
			key := string(FinanceManualOperationsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.MetaCreatedBy, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.MetaUpdatedAt):
			key := string(FinanceManualOperationsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.MetaUpdatedAt, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.MetaUpdatedBy):
			key := string(FinanceManualOperationsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.MetaUpdatedBy, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.MetaDeletedAt):
			key := string(FinanceManualOperationsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.MetaDeletedAt.Time, explicitAlias)

		case string(model.FinanceManualOperationsDBFieldName.MetaDeletedBy):
			key := string(FinanceManualOperationsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setFinanceManualOperationsSelectableValue(financeManualOperationsSelectableResponse, key, financeManualOperations.MetaDeletedBy, explicitAlias)

		}
	}
	return financeManualOperationsSelectableResponse
}

func NewFinanceManualOperationsListResponseFromFilterResult(result []model.FinanceManualOperationsFilterResult, filter model.Filter) FinanceManualOperationsSelectableListResponse {
	dtoFinanceManualOperationsListResponse := FinanceManualOperationsSelectableListResponse{}
	for _, row := range result {
		dtoFinanceManualOperationsResponse := NewFinanceManualOperationsSelectableResponse(row.FinanceManualOperations, filter)
		dtoFinanceManualOperationsListResponse = append(dtoFinanceManualOperationsListResponse, &dtoFinanceManualOperationsResponse)
	}
	return dtoFinanceManualOperationsListResponse
}

type FinanceManualOperationsFilterResponse struct {
	Metadata Metadata                                      `json:"metadata"`
	Data     FinanceManualOperationsSelectableListResponse `json:"data"`
}

func reverseFinanceManualOperationsFilterResults(result []model.FinanceManualOperationsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewFinanceManualOperationsFilterResponse(result []model.FinanceManualOperationsFilterResult, filter model.Filter) (resp FinanceManualOperationsFilterResponse) {
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
			reverseFinanceManualOperationsFilterResults(dataResult)
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

	resp.Data = NewFinanceManualOperationsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type FinanceManualOperationsCreateRequest struct {
	OperationCode   string                `json:"operationCode"`
	OperationType   model.OperationType   `json:"operationType" example:"manual_adjustment" enums:"manual_adjustment,manual_refund,manual_release_hold,manual_payout_retry,manual_writeoff,manual_reconciliation_resolution"`
	TargetRefType   string                `json:"targetRefType"`
	TargetRefId     uuid.UUID             `json:"targetRefId"`
	RequestedBy     uuid.UUID             `json:"requestedBy"`
	OperationStatus model.OperationStatus `json:"operationStatus" example:"requested" enums:"requested,approved,rejected,executed,cancelled"`
	ReasonCode      string                `json:"reasonCode"`
	ReasonDetail    string                `json:"reasonDetail"`
	Payload         json.RawMessage       `json:"payload"`
	ExecutedAt      time.Time             `json:"executedAt"`
	Metadata        json.RawMessage       `json:"metadata"`
}

func (d *FinanceManualOperationsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *FinanceManualOperationsCreateRequest) ToModel() model.FinanceManualOperations {
	id, _ := uuid.NewV4()
	return model.FinanceManualOperations{
		Id:              id,
		OperationCode:   d.OperationCode,
		OperationType:   d.OperationType,
		TargetRefType:   d.TargetRefType,
		TargetRefId:     d.TargetRefId,
		RequestedBy:     d.RequestedBy,
		OperationStatus: d.OperationStatus,
		ReasonCode:      d.ReasonCode,
		ReasonDetail:    null.StringFrom(d.ReasonDetail),
		Payload:         d.Payload,
		ExecutedAt:      null.TimeFrom(d.ExecutedAt),
		Metadata:        d.Metadata,
	}
}

type FinanceManualOperationsListCreateRequest []*FinanceManualOperationsCreateRequest

func (d FinanceManualOperationsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeManualOperations := range d {
		err = validator.Struct(financeManualOperations)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceManualOperationsListCreateRequest) ToModelList() []model.FinanceManualOperations {
	out := make([]model.FinanceManualOperations, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type FinanceManualOperationsUpdateRequest struct {
	OperationCode   string                `json:"operationCode"`
	OperationType   model.OperationType   `json:"operationType" example:"manual_adjustment" enums:"manual_adjustment,manual_refund,manual_release_hold,manual_payout_retry,manual_writeoff,manual_reconciliation_resolution"`
	TargetRefType   string                `json:"targetRefType"`
	TargetRefId     uuid.UUID             `json:"targetRefId"`
	RequestedBy     uuid.UUID             `json:"requestedBy"`
	OperationStatus model.OperationStatus `json:"operationStatus" example:"requested" enums:"requested,approved,rejected,executed,cancelled"`
	ReasonCode      string                `json:"reasonCode"`
	ReasonDetail    string                `json:"reasonDetail"`
	Payload         json.RawMessage       `json:"payload"`
	ExecutedAt      time.Time             `json:"executedAt"`
	Metadata        json.RawMessage       `json:"metadata"`
}

func (d *FinanceManualOperationsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d FinanceManualOperationsUpdateRequest) ToModel() model.FinanceManualOperations {
	return model.FinanceManualOperations{
		OperationCode:   d.OperationCode,
		OperationType:   d.OperationType,
		TargetRefType:   d.TargetRefType,
		TargetRefId:     d.TargetRefId,
		RequestedBy:     d.RequestedBy,
		OperationStatus: d.OperationStatus,
		ReasonCode:      d.ReasonCode,
		ReasonDetail:    null.StringFrom(d.ReasonDetail),
		Payload:         d.Payload,
		ExecutedAt:      null.TimeFrom(d.ExecutedAt),
		Metadata:        d.Metadata,
	}
}

type FinanceManualOperationsBulkUpdateRequest struct {
	Id              uuid.UUID             `json:"id"`
	OperationCode   string                `json:"operationCode"`
	OperationType   model.OperationType   `json:"operationType" example:"manual_adjustment" enums:"manual_adjustment,manual_refund,manual_release_hold,manual_payout_retry,manual_writeoff,manual_reconciliation_resolution"`
	TargetRefType   string                `json:"targetRefType"`
	TargetRefId     uuid.UUID             `json:"targetRefId"`
	RequestedBy     uuid.UUID             `json:"requestedBy"`
	OperationStatus model.OperationStatus `json:"operationStatus" example:"requested" enums:"requested,approved,rejected,executed,cancelled"`
	ReasonCode      string                `json:"reasonCode"`
	ReasonDetail    string                `json:"reasonDetail"`
	Payload         json.RawMessage       `json:"payload"`
	ExecutedAt      time.Time             `json:"executedAt"`
	Metadata        json.RawMessage       `json:"metadata"`
}

func (d FinanceManualOperationsBulkUpdateRequest) PrimaryID() FinanceManualOperationsPrimaryID {
	return FinanceManualOperationsPrimaryID{
		Id: d.Id,
	}
}

type FinanceManualOperationsListBulkUpdateRequest []*FinanceManualOperationsBulkUpdateRequest

func (d FinanceManualOperationsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeManualOperations := range d {
		err = validator.Struct(financeManualOperations)
		if err != nil {
			return
		}
	}
	return nil
}

func (d FinanceManualOperationsBulkUpdateRequest) ToModel() model.FinanceManualOperations {
	return model.FinanceManualOperations{
		Id:              d.Id,
		OperationCode:   d.OperationCode,
		OperationType:   d.OperationType,
		TargetRefType:   d.TargetRefType,
		TargetRefId:     d.TargetRefId,
		RequestedBy:     d.RequestedBy,
		OperationStatus: d.OperationStatus,
		ReasonCode:      d.ReasonCode,
		ReasonDetail:    null.StringFrom(d.ReasonDetail),
		Payload:         d.Payload,
		ExecutedAt:      null.TimeFrom(d.ExecutedAt),
		Metadata:        d.Metadata,
	}
}

type FinanceManualOperationsResponse struct {
	Id              uuid.UUID             `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OperationCode   string                `json:"operationCode" validate:"required"`
	OperationType   model.OperationType   `json:"operationType" validate:"required,oneof=manual_adjustment manual_refund manual_release_hold manual_payout_retry manual_writeoff manual_reconciliation_resolution" enums:"manual_adjustment,manual_refund,manual_release_hold,manual_payout_retry,manual_writeoff,manual_reconciliation_resolution"`
	TargetRefType   string                `json:"targetRefType" validate:"required"`
	TargetRefId     uuid.UUID             `json:"targetRefId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RequestedBy     uuid.UUID             `json:"requestedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OperationStatus model.OperationStatus `json:"operationStatus" validate:"oneof=requested approved rejected executed cancelled" enums:"requested,approved,rejected,executed,cancelled"`
	ReasonCode      string                `json:"reasonCode" validate:"required"`
	ReasonDetail    string                `json:"reasonDetail"`
	Payload         json.RawMessage       `json:"payload" swaggertype:"object"`
	ExecutedAt      time.Time             `json:"executedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata        json.RawMessage       `json:"metadata" swaggertype:"object"`
}

func NewFinanceManualOperationsResponse(financeManualOperations model.FinanceManualOperations) FinanceManualOperationsResponse {
	return FinanceManualOperationsResponse{
		Id:              financeManualOperations.Id,
		OperationCode:   financeManualOperations.OperationCode,
		OperationType:   model.OperationType(financeManualOperations.OperationType),
		TargetRefType:   financeManualOperations.TargetRefType,
		TargetRefId:     financeManualOperations.TargetRefId,
		RequestedBy:     financeManualOperations.RequestedBy,
		OperationStatus: model.OperationStatus(financeManualOperations.OperationStatus),
		ReasonCode:      financeManualOperations.ReasonCode,
		ReasonDetail:    financeManualOperations.ReasonDetail.String,
		Payload:         financeManualOperations.Payload,
		ExecutedAt:      financeManualOperations.ExecutedAt.Time,
		Metadata:        financeManualOperations.Metadata,
	}
}

type FinanceManualOperationsListResponse []*FinanceManualOperationsResponse

func NewFinanceManualOperationsListResponse(financeManualOperationsList model.FinanceManualOperationsList) FinanceManualOperationsListResponse {
	dtoFinanceManualOperationsListResponse := FinanceManualOperationsListResponse{}
	for _, financeManualOperations := range financeManualOperationsList {
		dtoFinanceManualOperationsResponse := NewFinanceManualOperationsResponse(*financeManualOperations)
		dtoFinanceManualOperationsListResponse = append(dtoFinanceManualOperationsListResponse, &dtoFinanceManualOperationsResponse)
	}
	return dtoFinanceManualOperationsListResponse
}

type FinanceManualOperationsPrimaryIDList []FinanceManualOperationsPrimaryID

func (d FinanceManualOperationsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, financeManualOperations := range d {
		err = validator.Struct(financeManualOperations)
		if err != nil {
			return
		}
	}
	return nil
}

type FinanceManualOperationsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *FinanceManualOperationsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d FinanceManualOperationsPrimaryID) ToModel() model.FinanceManualOperationsPrimaryID {
	return model.FinanceManualOperationsPrimaryID{
		Id: d.Id,
	}
}
