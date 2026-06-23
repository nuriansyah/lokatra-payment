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

type ReconciliationCandidatesDTOFieldNameType string

type reconciliationCandidatesDTOFieldName struct {
	Id                  ReconciliationCandidatesDTOFieldNameType
	ReconciliationRunId ReconciliationCandidatesDTOFieldNameType
	SourceSystem        ReconciliationCandidatesDTOFieldNameType
	SourceRefId         ReconciliationCandidatesDTOFieldNameType
	CandidateKey        ReconciliationCandidatesDTOFieldNameType
	Amount              ReconciliationCandidatesDTOFieldNameType
	OccurredAt          ReconciliationCandidatesDTOFieldNameType
	NormalizedPayload   ReconciliationCandidatesDTOFieldNameType
	Metadata            ReconciliationCandidatesDTOFieldNameType
	MetaCreatedAt       ReconciliationCandidatesDTOFieldNameType
	MetaCreatedBy       ReconciliationCandidatesDTOFieldNameType
	MetaUpdatedAt       ReconciliationCandidatesDTOFieldNameType
	MetaUpdatedBy       ReconciliationCandidatesDTOFieldNameType
	MetaDeletedAt       ReconciliationCandidatesDTOFieldNameType
	MetaDeletedBy       ReconciliationCandidatesDTOFieldNameType
}

var ReconciliationCandidatesDTOFieldName = reconciliationCandidatesDTOFieldName{
	Id:                  "id",
	ReconciliationRunId: "reconciliationRunId",
	SourceSystem:        "sourceSystem",
	SourceRefId:         "sourceRefId",
	CandidateKey:        "candidateKey",
	Amount:              "amount",
	OccurredAt:          "occurredAt",
	NormalizedPayload:   "normalizedPayload",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformReconciliationCandidatesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ReconciliationCandidatesDTOFieldName.Id):
		return string(model.ReconciliationCandidatesDBFieldName.Id), true

	case string(ReconciliationCandidatesDTOFieldName.ReconciliationRunId):
		return string(model.ReconciliationCandidatesDBFieldName.ReconciliationRunId), true

	case string(ReconciliationCandidatesDTOFieldName.SourceSystem):
		return string(model.ReconciliationCandidatesDBFieldName.SourceSystem), true

	case string(ReconciliationCandidatesDTOFieldName.SourceRefId):
		return string(model.ReconciliationCandidatesDBFieldName.SourceRefId), true

	case string(ReconciliationCandidatesDTOFieldName.CandidateKey):
		return string(model.ReconciliationCandidatesDBFieldName.CandidateKey), true

	case string(ReconciliationCandidatesDTOFieldName.Amount):
		return string(model.ReconciliationCandidatesDBFieldName.Amount), true

	case string(ReconciliationCandidatesDTOFieldName.OccurredAt):
		return string(model.ReconciliationCandidatesDBFieldName.OccurredAt), true

	case string(ReconciliationCandidatesDTOFieldName.NormalizedPayload):
		return string(model.ReconciliationCandidatesDBFieldName.NormalizedPayload), true

	case string(ReconciliationCandidatesDTOFieldName.Metadata):
		return string(model.ReconciliationCandidatesDBFieldName.Metadata), true

	case string(ReconciliationCandidatesDTOFieldName.MetaCreatedAt):
		return string(model.ReconciliationCandidatesDBFieldName.MetaCreatedAt), true

	case string(ReconciliationCandidatesDTOFieldName.MetaCreatedBy):
		return string(model.ReconciliationCandidatesDBFieldName.MetaCreatedBy), true

	case string(ReconciliationCandidatesDTOFieldName.MetaUpdatedAt):
		return string(model.ReconciliationCandidatesDBFieldName.MetaUpdatedAt), true

	case string(ReconciliationCandidatesDTOFieldName.MetaUpdatedBy):
		return string(model.ReconciliationCandidatesDBFieldName.MetaUpdatedBy), true

	case string(ReconciliationCandidatesDTOFieldName.MetaDeletedAt):
		return string(model.ReconciliationCandidatesDBFieldName.MetaDeletedAt), true

	case string(ReconciliationCandidatesDTOFieldName.MetaDeletedBy):
		return string(model.ReconciliationCandidatesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewReconciliationCandidatesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isReconciliationCandidatesBaseFilterField(field string) bool {
	spec, found := model.NewReconciliationCandidatesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeReconciliationCandidatesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateReconciliationCandidatesProjectionOutputPath(path string) error {
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

func transformReconciliationCandidatesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformReconciliationCandidatesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformReconciliationCandidatesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformReconciliationCandidatesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformReconciliationCandidatesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isReconciliationCandidatesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateReconciliationCandidatesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeReconciliationCandidatesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformReconciliationCandidatesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformReconciliationCandidatesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformReconciliationCandidatesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultReconciliationCandidatesFilter(filter *model.Filter) {
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
			Field: string(ReconciliationCandidatesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ReconciliationCandidatesSelectableResponse map[string]interface{}
type ReconciliationCandidatesSelectableListResponse []*ReconciliationCandidatesSelectableResponse

func assignReconciliationCandidatesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setReconciliationCandidatesSelectableValue(out ReconciliationCandidatesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignReconciliationCandidatesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewReconciliationCandidatesSelectableResponse(reconciliationCandidates model.ReconciliationCandidates, filter model.Filter) ReconciliationCandidatesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ReconciliationCandidatesDBFieldName.Id),
			string(model.ReconciliationCandidatesDBFieldName.ReconciliationRunId),
			string(model.ReconciliationCandidatesDBFieldName.SourceSystem),
			string(model.ReconciliationCandidatesDBFieldName.SourceRefId),
			string(model.ReconciliationCandidatesDBFieldName.CandidateKey),
			string(model.ReconciliationCandidatesDBFieldName.Amount),
			string(model.ReconciliationCandidatesDBFieldName.OccurredAt),
			string(model.ReconciliationCandidatesDBFieldName.NormalizedPayload),
			string(model.ReconciliationCandidatesDBFieldName.Metadata),
			string(model.ReconciliationCandidatesDBFieldName.MetaCreatedAt),
			string(model.ReconciliationCandidatesDBFieldName.MetaCreatedBy),
			string(model.ReconciliationCandidatesDBFieldName.MetaUpdatedAt),
			string(model.ReconciliationCandidatesDBFieldName.MetaUpdatedBy),
			string(model.ReconciliationCandidatesDBFieldName.MetaDeletedAt),
			string(model.ReconciliationCandidatesDBFieldName.MetaDeletedBy),
		)
	}
	reconciliationCandidatesSelectableResponse := ReconciliationCandidatesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ReconciliationCandidatesDBFieldName.Id):
			key := string(ReconciliationCandidatesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.Id, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.ReconciliationRunId):
			key := string(ReconciliationCandidatesDTOFieldName.ReconciliationRunId)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.ReconciliationRunId, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.SourceSystem):
			key := string(ReconciliationCandidatesDTOFieldName.SourceSystem)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, model.SourceSystem(reconciliationCandidates.SourceSystem), explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.SourceRefId):
			key := string(ReconciliationCandidatesDTOFieldName.SourceRefId)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.SourceRefId, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.CandidateKey):
			key := string(ReconciliationCandidatesDTOFieldName.CandidateKey)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.CandidateKey, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.Amount):
			key := string(ReconciliationCandidatesDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.Amount, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.OccurredAt):
			key := string(ReconciliationCandidatesDTOFieldName.OccurredAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.OccurredAt.Time, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.NormalizedPayload):
			key := string(ReconciliationCandidatesDTOFieldName.NormalizedPayload)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.NormalizedPayload, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.Metadata):
			key := string(ReconciliationCandidatesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.Metadata, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.MetaCreatedAt):
			key := string(ReconciliationCandidatesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.MetaCreatedAt, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.MetaCreatedBy):
			key := string(ReconciliationCandidatesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.MetaCreatedBy, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.MetaUpdatedAt):
			key := string(ReconciliationCandidatesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.MetaUpdatedAt, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.MetaUpdatedBy):
			key := string(ReconciliationCandidatesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.MetaUpdatedBy, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.MetaDeletedAt):
			key := string(ReconciliationCandidatesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.MetaDeletedAt.Time, explicitAlias)

		case string(model.ReconciliationCandidatesDBFieldName.MetaDeletedBy):
			key := string(ReconciliationCandidatesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationCandidatesSelectableValue(reconciliationCandidatesSelectableResponse, key, reconciliationCandidates.MetaDeletedBy, explicitAlias)

		}
	}
	return reconciliationCandidatesSelectableResponse
}

func NewReconciliationCandidatesListResponseFromFilterResult(result []model.ReconciliationCandidatesFilterResult, filter model.Filter) ReconciliationCandidatesSelectableListResponse {
	dtoReconciliationCandidatesListResponse := ReconciliationCandidatesSelectableListResponse{}
	for _, row := range result {
		dtoReconciliationCandidatesResponse := NewReconciliationCandidatesSelectableResponse(row.ReconciliationCandidates, filter)
		dtoReconciliationCandidatesListResponse = append(dtoReconciliationCandidatesListResponse, &dtoReconciliationCandidatesResponse)
	}
	return dtoReconciliationCandidatesListResponse
}

type ReconciliationCandidatesFilterResponse struct {
	Metadata Metadata                                       `json:"metadata"`
	Data     ReconciliationCandidatesSelectableListResponse `json:"data"`
}

func reverseReconciliationCandidatesFilterResults(result []model.ReconciliationCandidatesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewReconciliationCandidatesFilterResponse(result []model.ReconciliationCandidatesFilterResult, filter model.Filter) (resp ReconciliationCandidatesFilterResponse) {
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
			reverseReconciliationCandidatesFilterResults(dataResult)
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

	resp.Data = NewReconciliationCandidatesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ReconciliationCandidatesCreateRequest struct {
	ReconciliationRunId uuid.UUID          `json:"reconciliationRunId"`
	SourceSystem        model.SourceSystem `json:"sourceSystem" example:"ledger" enums:"ledger,provider_statement,bank_statement"`
	SourceRefId         uuid.UUID          `json:"sourceRefId"`
	CandidateKey        string             `json:"candidateKey"`
	Amount              decimal.Decimal    `json:"amount"`
	OccurredAt          time.Time          `json:"occurredAt"`
	NormalizedPayload   json.RawMessage    `json:"normalizedPayload"`
	Metadata            json.RawMessage    `json:"metadata"`
}

func (d *ReconciliationCandidatesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ReconciliationCandidatesCreateRequest) ToModel() model.ReconciliationCandidates {
	id, _ := uuid.NewV4()
	return model.ReconciliationCandidates{
		Id:                  id,
		ReconciliationRunId: d.ReconciliationRunId,
		SourceSystem:        d.SourceSystem,
		SourceRefId:         d.SourceRefId,
		CandidateKey:        d.CandidateKey,
		Amount:              d.Amount,
		OccurredAt:          null.TimeFrom(d.OccurredAt),
		NormalizedPayload:   d.NormalizedPayload,
		Metadata:            d.Metadata,
	}
}

type ReconciliationCandidatesListCreateRequest []*ReconciliationCandidatesCreateRequest

func (d ReconciliationCandidatesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationCandidates := range d {
		err = validator.Struct(reconciliationCandidates)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationCandidatesListCreateRequest) ToModelList() []model.ReconciliationCandidates {
	out := make([]model.ReconciliationCandidates, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ReconciliationCandidatesUpdateRequest struct {
	ReconciliationRunId uuid.UUID          `json:"reconciliationRunId"`
	SourceSystem        model.SourceSystem `json:"sourceSystem" example:"ledger" enums:"ledger,provider_statement,bank_statement"`
	SourceRefId         uuid.UUID          `json:"sourceRefId"`
	CandidateKey        string             `json:"candidateKey"`
	Amount              decimal.Decimal    `json:"amount"`
	OccurredAt          time.Time          `json:"occurredAt"`
	NormalizedPayload   json.RawMessage    `json:"normalizedPayload"`
	Metadata            json.RawMessage    `json:"metadata"`
}

func (d *ReconciliationCandidatesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ReconciliationCandidatesUpdateRequest) ToModel() model.ReconciliationCandidates {
	return model.ReconciliationCandidates{
		ReconciliationRunId: d.ReconciliationRunId,
		SourceSystem:        d.SourceSystem,
		SourceRefId:         d.SourceRefId,
		CandidateKey:        d.CandidateKey,
		Amount:              d.Amount,
		OccurredAt:          null.TimeFrom(d.OccurredAt),
		NormalizedPayload:   d.NormalizedPayload,
		Metadata:            d.Metadata,
	}
}

type ReconciliationCandidatesBulkUpdateRequest struct {
	Id                  uuid.UUID          `json:"id"`
	ReconciliationRunId uuid.UUID          `json:"reconciliationRunId"`
	SourceSystem        model.SourceSystem `json:"sourceSystem" example:"ledger" enums:"ledger,provider_statement,bank_statement"`
	SourceRefId         uuid.UUID          `json:"sourceRefId"`
	CandidateKey        string             `json:"candidateKey"`
	Amount              decimal.Decimal    `json:"amount"`
	OccurredAt          time.Time          `json:"occurredAt"`
	NormalizedPayload   json.RawMessage    `json:"normalizedPayload"`
	Metadata            json.RawMessage    `json:"metadata"`
}

func (d ReconciliationCandidatesBulkUpdateRequest) PrimaryID() ReconciliationCandidatesPrimaryID {
	return ReconciliationCandidatesPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationCandidatesListBulkUpdateRequest []*ReconciliationCandidatesBulkUpdateRequest

func (d ReconciliationCandidatesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationCandidates := range d {
		err = validator.Struct(reconciliationCandidates)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationCandidatesBulkUpdateRequest) ToModel() model.ReconciliationCandidates {
	return model.ReconciliationCandidates{
		Id:                  d.Id,
		ReconciliationRunId: d.ReconciliationRunId,
		SourceSystem:        d.SourceSystem,
		SourceRefId:         d.SourceRefId,
		CandidateKey:        d.CandidateKey,
		Amount:              d.Amount,
		OccurredAt:          null.TimeFrom(d.OccurredAt),
		NormalizedPayload:   d.NormalizedPayload,
		Metadata:            d.Metadata,
	}
}

type ReconciliationCandidatesResponse struct {
	Id                  uuid.UUID          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReconciliationRunId uuid.UUID          `json:"reconciliationRunId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SourceSystem        model.SourceSystem `json:"sourceSystem" validate:"required,oneof=ledger provider_statement bank_statement" enums:"ledger,provider_statement,bank_statement"`
	SourceRefId         uuid.UUID          `json:"sourceRefId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	CandidateKey        string             `json:"candidateKey" validate:"required"`
	Amount              decimal.Decimal    `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	OccurredAt          time.Time          `json:"occurredAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	NormalizedPayload   json.RawMessage    `json:"normalizedPayload" swaggertype:"object"`
	Metadata            json.RawMessage    `json:"metadata" swaggertype:"object"`
}

func NewReconciliationCandidatesResponse(reconciliationCandidates model.ReconciliationCandidates) ReconciliationCandidatesResponse {
	return ReconciliationCandidatesResponse{
		Id:                  reconciliationCandidates.Id,
		ReconciliationRunId: reconciliationCandidates.ReconciliationRunId,
		SourceSystem:        model.SourceSystem(reconciliationCandidates.SourceSystem),
		SourceRefId:         reconciliationCandidates.SourceRefId,
		CandidateKey:        reconciliationCandidates.CandidateKey,
		Amount:              reconciliationCandidates.Amount,
		OccurredAt:          reconciliationCandidates.OccurredAt.Time,
		NormalizedPayload:   reconciliationCandidates.NormalizedPayload,
		Metadata:            reconciliationCandidates.Metadata,
	}
}

type ReconciliationCandidatesListResponse []*ReconciliationCandidatesResponse

func NewReconciliationCandidatesListResponse(reconciliationCandidatesList model.ReconciliationCandidatesList) ReconciliationCandidatesListResponse {
	dtoReconciliationCandidatesListResponse := ReconciliationCandidatesListResponse{}
	for _, reconciliationCandidates := range reconciliationCandidatesList {
		dtoReconciliationCandidatesResponse := NewReconciliationCandidatesResponse(*reconciliationCandidates)
		dtoReconciliationCandidatesListResponse = append(dtoReconciliationCandidatesListResponse, &dtoReconciliationCandidatesResponse)
	}
	return dtoReconciliationCandidatesListResponse
}

type ReconciliationCandidatesPrimaryIDList []ReconciliationCandidatesPrimaryID

func (d ReconciliationCandidatesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationCandidates := range d {
		err = validator.Struct(reconciliationCandidates)
		if err != nil {
			return
		}
	}
	return nil
}

type ReconciliationCandidatesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ReconciliationCandidatesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ReconciliationCandidatesPrimaryID) ToModel() model.ReconciliationCandidatesPrimaryID {
	return model.ReconciliationCandidatesPrimaryID{
		Id: d.Id,
	}
}
