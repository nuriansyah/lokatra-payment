package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ReconciliationMatchesDTOFieldNameType string

type reconciliationMatchesDTOFieldName struct {
	Id                  ReconciliationMatchesDTOFieldNameType
	ReconciliationRunId ReconciliationMatchesDTOFieldNameType
	LedgerJournalId     ReconciliationMatchesDTOFieldNameType
	StatementLineId     ReconciliationMatchesDTOFieldNameType
	MatchType           ReconciliationMatchesDTOFieldNameType
	MatchStatus         ReconciliationMatchesDTOFieldNameType
	AmountDifference    ReconciliationMatchesDTOFieldNameType
	MatchedAt           ReconciliationMatchesDTOFieldNameType
	Metadata            ReconciliationMatchesDTOFieldNameType
	MetaCreatedAt       ReconciliationMatchesDTOFieldNameType
	MetaCreatedBy       ReconciliationMatchesDTOFieldNameType
	MetaUpdatedAt       ReconciliationMatchesDTOFieldNameType
	MetaUpdatedBy       ReconciliationMatchesDTOFieldNameType
	MetaDeletedAt       ReconciliationMatchesDTOFieldNameType
	MetaDeletedBy       ReconciliationMatchesDTOFieldNameType
}

var ReconciliationMatchesDTOFieldName = reconciliationMatchesDTOFieldName{
	Id:                  "id",
	ReconciliationRunId: "reconciliationRunId",
	LedgerJournalId:     "ledgerJournalId",
	StatementLineId:     "statementLineId",
	MatchType:           "matchType",
	MatchStatus:         "matchStatus",
	AmountDifference:    "amountDifference",
	MatchedAt:           "matchedAt",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformReconciliationMatchesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ReconciliationMatchesDTOFieldName.Id):
		return string(model.ReconciliationMatchesDBFieldName.Id), true

	case string(ReconciliationMatchesDTOFieldName.ReconciliationRunId):
		return string(model.ReconciliationMatchesDBFieldName.ReconciliationRunId), true

	case string(ReconciliationMatchesDTOFieldName.LedgerJournalId):
		return string(model.ReconciliationMatchesDBFieldName.LedgerJournalId), true

	case string(ReconciliationMatchesDTOFieldName.StatementLineId):
		return string(model.ReconciliationMatchesDBFieldName.StatementLineId), true

	case string(ReconciliationMatchesDTOFieldName.MatchType):
		return string(model.ReconciliationMatchesDBFieldName.MatchType), true

	case string(ReconciliationMatchesDTOFieldName.MatchStatus):
		return string(model.ReconciliationMatchesDBFieldName.MatchStatus), true

	case string(ReconciliationMatchesDTOFieldName.AmountDifference):
		return string(model.ReconciliationMatchesDBFieldName.AmountDifference), true

	case string(ReconciliationMatchesDTOFieldName.MatchedAt):
		return string(model.ReconciliationMatchesDBFieldName.MatchedAt), true

	case string(ReconciliationMatchesDTOFieldName.Metadata):
		return string(model.ReconciliationMatchesDBFieldName.Metadata), true

	case string(ReconciliationMatchesDTOFieldName.MetaCreatedAt):
		return string(model.ReconciliationMatchesDBFieldName.MetaCreatedAt), true

	case string(ReconciliationMatchesDTOFieldName.MetaCreatedBy):
		return string(model.ReconciliationMatchesDBFieldName.MetaCreatedBy), true

	case string(ReconciliationMatchesDTOFieldName.MetaUpdatedAt):
		return string(model.ReconciliationMatchesDBFieldName.MetaUpdatedAt), true

	case string(ReconciliationMatchesDTOFieldName.MetaUpdatedBy):
		return string(model.ReconciliationMatchesDBFieldName.MetaUpdatedBy), true

	case string(ReconciliationMatchesDTOFieldName.MetaDeletedAt):
		return string(model.ReconciliationMatchesDBFieldName.MetaDeletedAt), true

	case string(ReconciliationMatchesDTOFieldName.MetaDeletedBy):
		return string(model.ReconciliationMatchesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewReconciliationMatchesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isReconciliationMatchesBaseFilterField(field string) bool {
	spec, found := model.NewReconciliationMatchesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeReconciliationMatchesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateReconciliationMatchesProjectionOutputPath(path string) error {
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

func transformReconciliationMatchesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformReconciliationMatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformReconciliationMatchesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformReconciliationMatchesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformReconciliationMatchesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isReconciliationMatchesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateReconciliationMatchesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeReconciliationMatchesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformReconciliationMatchesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformReconciliationMatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformReconciliationMatchesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultReconciliationMatchesFilter(filter *model.Filter) {
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
			Field: string(ReconciliationMatchesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ReconciliationMatchesSelectableResponse map[string]interface{}
type ReconciliationMatchesSelectableListResponse []*ReconciliationMatchesSelectableResponse

func assignReconciliationMatchesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setReconciliationMatchesSelectableValue(out ReconciliationMatchesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignReconciliationMatchesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewReconciliationMatchesSelectableResponse(reconciliationMatches model.ReconciliationMatches, filter model.Filter) ReconciliationMatchesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ReconciliationMatchesDBFieldName.Id),
			string(model.ReconciliationMatchesDBFieldName.ReconciliationRunId),
			string(model.ReconciliationMatchesDBFieldName.LedgerJournalId),
			string(model.ReconciliationMatchesDBFieldName.StatementLineId),
			string(model.ReconciliationMatchesDBFieldName.MatchType),
			string(model.ReconciliationMatchesDBFieldName.MatchStatus),
			string(model.ReconciliationMatchesDBFieldName.AmountDifference),
			string(model.ReconciliationMatchesDBFieldName.MatchedAt),
			string(model.ReconciliationMatchesDBFieldName.Metadata),
			string(model.ReconciliationMatchesDBFieldName.MetaCreatedAt),
			string(model.ReconciliationMatchesDBFieldName.MetaCreatedBy),
			string(model.ReconciliationMatchesDBFieldName.MetaUpdatedAt),
			string(model.ReconciliationMatchesDBFieldName.MetaUpdatedBy),
			string(model.ReconciliationMatchesDBFieldName.MetaDeletedAt),
			string(model.ReconciliationMatchesDBFieldName.MetaDeletedBy),
		)
	}
	reconciliationMatchesSelectableResponse := ReconciliationMatchesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ReconciliationMatchesDBFieldName.Id):
			key := string(ReconciliationMatchesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.Id, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.ReconciliationRunId):
			key := string(ReconciliationMatchesDTOFieldName.ReconciliationRunId)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.ReconciliationRunId, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.LedgerJournalId):
			key := string(ReconciliationMatchesDTOFieldName.LedgerJournalId)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.LedgerJournalId.UUID, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.StatementLineId):
			key := string(ReconciliationMatchesDTOFieldName.StatementLineId)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.StatementLineId.UUID, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MatchType):
			key := string(ReconciliationMatchesDTOFieldName.MatchType)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, model.MatchType(reconciliationMatches.MatchType), explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MatchStatus):
			key := string(ReconciliationMatchesDTOFieldName.MatchStatus)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, model.MatchStatus(reconciliationMatches.MatchStatus), explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.AmountDifference):
			key := string(ReconciliationMatchesDTOFieldName.AmountDifference)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.AmountDifference, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MatchedAt):
			key := string(ReconciliationMatchesDTOFieldName.MatchedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.MatchedAt, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.Metadata):
			key := string(ReconciliationMatchesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.Metadata, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MetaCreatedAt):
			key := string(ReconciliationMatchesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.MetaCreatedAt, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MetaCreatedBy):
			key := string(ReconciliationMatchesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.MetaCreatedBy, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MetaUpdatedAt):
			key := string(ReconciliationMatchesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.MetaUpdatedAt, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MetaUpdatedBy):
			key := string(ReconciliationMatchesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.MetaUpdatedBy, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MetaDeletedAt):
			key := string(ReconciliationMatchesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.MetaDeletedAt.Time, explicitAlias)

		case string(model.ReconciliationMatchesDBFieldName.MetaDeletedBy):
			key := string(ReconciliationMatchesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setReconciliationMatchesSelectableValue(reconciliationMatchesSelectableResponse, key, reconciliationMatches.MetaDeletedBy, explicitAlias)

		}
	}
	return reconciliationMatchesSelectableResponse
}

func NewReconciliationMatchesListResponseFromFilterResult(result []model.ReconciliationMatchesFilterResult, filter model.Filter) ReconciliationMatchesSelectableListResponse {
	dtoReconciliationMatchesListResponse := ReconciliationMatchesSelectableListResponse{}
	for _, row := range result {
		dtoReconciliationMatchesResponse := NewReconciliationMatchesSelectableResponse(row.ReconciliationMatches, filter)
		dtoReconciliationMatchesListResponse = append(dtoReconciliationMatchesListResponse, &dtoReconciliationMatchesResponse)
	}
	return dtoReconciliationMatchesListResponse
}

type ReconciliationMatchesFilterResponse struct {
	Metadata Metadata                                    `json:"metadata"`
	Data     ReconciliationMatchesSelectableListResponse `json:"data"`
}

func reverseReconciliationMatchesFilterResults(result []model.ReconciliationMatchesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewReconciliationMatchesFilterResponse(result []model.ReconciliationMatchesFilterResult, filter model.Filter) (resp ReconciliationMatchesFilterResponse) {
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
			reverseReconciliationMatchesFilterResults(dataResult)
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

	resp.Data = NewReconciliationMatchesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ReconciliationMatchesCreateRequest struct {
	ReconciliationRunId uuid.UUID         `json:"reconciliationRunId"`
	LedgerJournalId     uuid.UUID         `json:"ledgerJournalId"`
	StatementLineId     uuid.UUID         `json:"statementLineId"`
	MatchType           model.MatchType   `json:"matchType" example:"exact" enums:"exact,amount_date,provider_ref,manual"`
	MatchStatus         model.MatchStatus `json:"matchStatus" example:"matched" enums:"matched,review,reversed"`
	AmountDifference    decimal.Decimal   `json:"amountDifference"`
	MatchedAt           time.Time         `json:"matchedAt"`
	Metadata            json.RawMessage   `json:"metadata"`
}

func (d *ReconciliationMatchesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ReconciliationMatchesCreateRequest) ToModel() model.ReconciliationMatches {
	id, _ := uuid.NewV4()
	return model.ReconciliationMatches{
		Id:                  id,
		ReconciliationRunId: d.ReconciliationRunId,
		LedgerJournalId:     nuuid.From(d.LedgerJournalId),
		StatementLineId:     nuuid.From(d.StatementLineId),
		MatchType:           d.MatchType,
		MatchStatus:         d.MatchStatus,
		AmountDifference:    d.AmountDifference,
		MatchedAt:           d.MatchedAt,
		Metadata:            d.Metadata,
	}
}

type ReconciliationMatchesListCreateRequest []*ReconciliationMatchesCreateRequest

func (d ReconciliationMatchesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationMatches := range d {
		err = validator.Struct(reconciliationMatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationMatchesListCreateRequest) ToModelList() []model.ReconciliationMatches {
	out := make([]model.ReconciliationMatches, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ReconciliationMatchesUpdateRequest struct {
	ReconciliationRunId uuid.UUID         `json:"reconciliationRunId"`
	LedgerJournalId     uuid.UUID         `json:"ledgerJournalId"`
	StatementLineId     uuid.UUID         `json:"statementLineId"`
	MatchType           model.MatchType   `json:"matchType" example:"exact" enums:"exact,amount_date,provider_ref,manual"`
	MatchStatus         model.MatchStatus `json:"matchStatus" example:"matched" enums:"matched,review,reversed"`
	AmountDifference    decimal.Decimal   `json:"amountDifference"`
	MatchedAt           time.Time         `json:"matchedAt"`
	Metadata            json.RawMessage   `json:"metadata"`
}

func (d *ReconciliationMatchesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ReconciliationMatchesUpdateRequest) ToModel() model.ReconciliationMatches {
	return model.ReconciliationMatches{
		ReconciliationRunId: d.ReconciliationRunId,
		LedgerJournalId:     nuuid.From(d.LedgerJournalId),
		StatementLineId:     nuuid.From(d.StatementLineId),
		MatchType:           d.MatchType,
		MatchStatus:         d.MatchStatus,
		AmountDifference:    d.AmountDifference,
		MatchedAt:           d.MatchedAt,
		Metadata:            d.Metadata,
	}
}

type ReconciliationMatchesBulkUpdateRequest struct {
	Id                  uuid.UUID         `json:"id"`
	ReconciliationRunId uuid.UUID         `json:"reconciliationRunId"`
	LedgerJournalId     uuid.UUID         `json:"ledgerJournalId"`
	StatementLineId     uuid.UUID         `json:"statementLineId"`
	MatchType           model.MatchType   `json:"matchType" example:"exact" enums:"exact,amount_date,provider_ref,manual"`
	MatchStatus         model.MatchStatus `json:"matchStatus" example:"matched" enums:"matched,review,reversed"`
	AmountDifference    decimal.Decimal   `json:"amountDifference"`
	MatchedAt           time.Time         `json:"matchedAt"`
	Metadata            json.RawMessage   `json:"metadata"`
}

func (d ReconciliationMatchesBulkUpdateRequest) PrimaryID() ReconciliationMatchesPrimaryID {
	return ReconciliationMatchesPrimaryID{
		Id: d.Id,
	}
}

type ReconciliationMatchesListBulkUpdateRequest []*ReconciliationMatchesBulkUpdateRequest

func (d ReconciliationMatchesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationMatches := range d {
		err = validator.Struct(reconciliationMatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ReconciliationMatchesBulkUpdateRequest) ToModel() model.ReconciliationMatches {
	return model.ReconciliationMatches{
		Id:                  d.Id,
		ReconciliationRunId: d.ReconciliationRunId,
		LedgerJournalId:     nuuid.From(d.LedgerJournalId),
		StatementLineId:     nuuid.From(d.StatementLineId),
		MatchType:           d.MatchType,
		MatchStatus:         d.MatchStatus,
		AmountDifference:    d.AmountDifference,
		MatchedAt:           d.MatchedAt,
		Metadata:            d.Metadata,
	}
}

type ReconciliationMatchesResponse struct {
	Id                  uuid.UUID         `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReconciliationRunId uuid.UUID         `json:"reconciliationRunId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	LedgerJournalId     uuid.UUID         `json:"ledgerJournalId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	StatementLineId     uuid.UUID         `json:"statementLineId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MatchType           model.MatchType   `json:"matchType" validate:"required,oneof=exact amount_date provider_ref manual" enums:"exact,amount_date,provider_ref,manual"`
	MatchStatus         model.MatchStatus `json:"matchStatus" validate:"oneof=matched review reversed" enums:"matched,review,reversed"`
	AmountDifference    decimal.Decimal   `json:"amountDifference" format:"decimal" example:"100.50"`
	MatchedAt           time.Time         `json:"matchedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata            json.RawMessage   `json:"metadata" swaggertype:"object"`
}

func NewReconciliationMatchesResponse(reconciliationMatches model.ReconciliationMatches) ReconciliationMatchesResponse {
	return ReconciliationMatchesResponse{
		Id:                  reconciliationMatches.Id,
		ReconciliationRunId: reconciliationMatches.ReconciliationRunId,
		LedgerJournalId:     reconciliationMatches.LedgerJournalId.UUID,
		StatementLineId:     reconciliationMatches.StatementLineId.UUID,
		MatchType:           model.MatchType(reconciliationMatches.MatchType),
		MatchStatus:         model.MatchStatus(reconciliationMatches.MatchStatus),
		AmountDifference:    reconciliationMatches.AmountDifference,
		MatchedAt:           reconciliationMatches.MatchedAt,
		Metadata:            reconciliationMatches.Metadata,
	}
}

type ReconciliationMatchesListResponse []*ReconciliationMatchesResponse

func NewReconciliationMatchesListResponse(reconciliationMatchesList model.ReconciliationMatchesList) ReconciliationMatchesListResponse {
	dtoReconciliationMatchesListResponse := ReconciliationMatchesListResponse{}
	for _, reconciliationMatches := range reconciliationMatchesList {
		dtoReconciliationMatchesResponse := NewReconciliationMatchesResponse(*reconciliationMatches)
		dtoReconciliationMatchesListResponse = append(dtoReconciliationMatchesListResponse, &dtoReconciliationMatchesResponse)
	}
	return dtoReconciliationMatchesListResponse
}

type ReconciliationMatchesPrimaryIDList []ReconciliationMatchesPrimaryID

func (d ReconciliationMatchesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, reconciliationMatches := range d {
		err = validator.Struct(reconciliationMatches)
		if err != nil {
			return
		}
	}
	return nil
}

type ReconciliationMatchesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ReconciliationMatchesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ReconciliationMatchesPrimaryID) ToModel() model.ReconciliationMatchesPrimaryID {
	return model.ReconciliationMatchesPrimaryID{
		Id: d.Id,
	}
}
