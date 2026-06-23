package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerJournalEntriesDTOFieldNameType string

type ledgerJournalEntriesDTOFieldName struct {
	Id                  LedgerJournalEntriesDTOFieldNameType
	BatchId             LedgerJournalEntriesDTOFieldNameType
	JournalCode         LedgerJournalEntriesDTOFieldNameType
	BookId              LedgerJournalEntriesDTOFieldNameType
	JournalType         LedgerJournalEntriesDTOFieldNameType
	SourceType          LedgerJournalEntriesDTOFieldNameType
	SourceId            LedgerJournalEntriesDTOFieldNameType
	IdempotencyKey      LedgerJournalEntriesDTOFieldNameType
	JournalStatus       LedgerJournalEntriesDTOFieldNameType
	EffectiveAt         LedgerJournalEntriesDTOFieldNameType
	BookedAt            LedgerJournalEntriesDTOFieldNameType
	Description         LedgerJournalEntriesDTOFieldNameType
	ReversalOfJournalId LedgerJournalEntriesDTOFieldNameType
	Metadata            LedgerJournalEntriesDTOFieldNameType
	MetaCreatedAt       LedgerJournalEntriesDTOFieldNameType
	MetaCreatedBy       LedgerJournalEntriesDTOFieldNameType
	MetaUpdatedAt       LedgerJournalEntriesDTOFieldNameType
	MetaUpdatedBy       LedgerJournalEntriesDTOFieldNameType
	MetaDeletedAt       LedgerJournalEntriesDTOFieldNameType
	MetaDeletedBy       LedgerJournalEntriesDTOFieldNameType
}

var LedgerJournalEntriesDTOFieldName = ledgerJournalEntriesDTOFieldName{
	Id:                  "id",
	BatchId:             "batchId",
	JournalCode:         "journalCode",
	BookId:              "bookId",
	JournalType:         "journalType",
	SourceType:          "sourceType",
	SourceId:            "sourceId",
	IdempotencyKey:      "idempotencyKey",
	JournalStatus:       "journalStatus",
	EffectiveAt:         "effectiveAt",
	BookedAt:            "bookedAt",
	Description:         "description",
	ReversalOfJournalId: "reversalOfJournalId",
	Metadata:            "metadata",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func transformLedgerJournalEntriesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerJournalEntriesDTOFieldName.Id):
		return string(model.LedgerJournalEntriesDBFieldName.Id), true

	case string(LedgerJournalEntriesDTOFieldName.BatchId):
		return string(model.LedgerJournalEntriesDBFieldName.BatchId), true

	case string(LedgerJournalEntriesDTOFieldName.JournalCode):
		return string(model.LedgerJournalEntriesDBFieldName.JournalCode), true

	case string(LedgerJournalEntriesDTOFieldName.BookId):
		return string(model.LedgerJournalEntriesDBFieldName.BookId), true

	case string(LedgerJournalEntriesDTOFieldName.JournalType):
		return string(model.LedgerJournalEntriesDBFieldName.JournalType), true

	case string(LedgerJournalEntriesDTOFieldName.SourceType):
		return string(model.LedgerJournalEntriesDBFieldName.SourceType), true

	case string(LedgerJournalEntriesDTOFieldName.SourceId):
		return string(model.LedgerJournalEntriesDBFieldName.SourceId), true

	case string(LedgerJournalEntriesDTOFieldName.IdempotencyKey):
		return string(model.LedgerJournalEntriesDBFieldName.IdempotencyKey), true

	case string(LedgerJournalEntriesDTOFieldName.JournalStatus):
		return string(model.LedgerJournalEntriesDBFieldName.JournalStatus), true

	case string(LedgerJournalEntriesDTOFieldName.EffectiveAt):
		return string(model.LedgerJournalEntriesDBFieldName.EffectiveAt), true

	case string(LedgerJournalEntriesDTOFieldName.BookedAt):
		return string(model.LedgerJournalEntriesDBFieldName.BookedAt), true

	case string(LedgerJournalEntriesDTOFieldName.Description):
		return string(model.LedgerJournalEntriesDBFieldName.Description), true

	case string(LedgerJournalEntriesDTOFieldName.ReversalOfJournalId):
		return string(model.LedgerJournalEntriesDBFieldName.ReversalOfJournalId), true

	case string(LedgerJournalEntriesDTOFieldName.Metadata):
		return string(model.LedgerJournalEntriesDBFieldName.Metadata), true

	case string(LedgerJournalEntriesDTOFieldName.MetaCreatedAt):
		return string(model.LedgerJournalEntriesDBFieldName.MetaCreatedAt), true

	case string(LedgerJournalEntriesDTOFieldName.MetaCreatedBy):
		return string(model.LedgerJournalEntriesDBFieldName.MetaCreatedBy), true

	case string(LedgerJournalEntriesDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerJournalEntriesDBFieldName.MetaUpdatedAt), true

	case string(LedgerJournalEntriesDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerJournalEntriesDBFieldName.MetaUpdatedBy), true

	case string(LedgerJournalEntriesDTOFieldName.MetaDeletedAt):
		return string(model.LedgerJournalEntriesDBFieldName.MetaDeletedAt), true

	case string(LedgerJournalEntriesDTOFieldName.MetaDeletedBy):
		return string(model.LedgerJournalEntriesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerJournalEntriesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerJournalEntriesBaseFilterField(field string) bool {
	spec, found := model.NewLedgerJournalEntriesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerJournalEntriesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerJournalEntriesProjectionOutputPath(path string) error {
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

func transformLedgerJournalEntriesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerJournalEntriesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerJournalEntriesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerJournalEntriesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerJournalEntriesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerJournalEntriesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerJournalEntriesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerJournalEntriesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerJournalEntriesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerJournalEntriesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerJournalEntriesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerJournalEntriesFilter(filter *model.Filter) {
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
			Field: string(LedgerJournalEntriesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerJournalEntriesSelectableResponse map[string]interface{}
type LedgerJournalEntriesSelectableListResponse []*LedgerJournalEntriesSelectableResponse

func assignLedgerJournalEntriesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerJournalEntriesSelectableValue(out LedgerJournalEntriesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerJournalEntriesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerJournalEntriesSelectableResponse(ledgerJournalEntries model.LedgerJournalEntries, filter model.Filter) LedgerJournalEntriesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerJournalEntriesDBFieldName.Id),
			string(model.LedgerJournalEntriesDBFieldName.BatchId),
			string(model.LedgerJournalEntriesDBFieldName.JournalCode),
			string(model.LedgerJournalEntriesDBFieldName.BookId),
			string(model.LedgerJournalEntriesDBFieldName.JournalType),
			string(model.LedgerJournalEntriesDBFieldName.SourceType),
			string(model.LedgerJournalEntriesDBFieldName.SourceId),
			string(model.LedgerJournalEntriesDBFieldName.IdempotencyKey),
			string(model.LedgerJournalEntriesDBFieldName.JournalStatus),
			string(model.LedgerJournalEntriesDBFieldName.EffectiveAt),
			string(model.LedgerJournalEntriesDBFieldName.BookedAt),
			string(model.LedgerJournalEntriesDBFieldName.Description),
			string(model.LedgerJournalEntriesDBFieldName.ReversalOfJournalId),
			string(model.LedgerJournalEntriesDBFieldName.Metadata),
			string(model.LedgerJournalEntriesDBFieldName.MetaCreatedAt),
			string(model.LedgerJournalEntriesDBFieldName.MetaCreatedBy),
			string(model.LedgerJournalEntriesDBFieldName.MetaUpdatedAt),
			string(model.LedgerJournalEntriesDBFieldName.MetaUpdatedBy),
			string(model.LedgerJournalEntriesDBFieldName.MetaDeletedAt),
			string(model.LedgerJournalEntriesDBFieldName.MetaDeletedBy),
		)
	}
	ledgerJournalEntriesSelectableResponse := LedgerJournalEntriesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerJournalEntriesDBFieldName.Id):
			key := string(LedgerJournalEntriesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.Id, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.BatchId):
			key := string(LedgerJournalEntriesDTOFieldName.BatchId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.BatchId.UUID, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.JournalCode):
			key := string(LedgerJournalEntriesDTOFieldName.JournalCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.JournalCode, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.BookId):
			key := string(LedgerJournalEntriesDTOFieldName.BookId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.BookId, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.JournalType):
			key := string(LedgerJournalEntriesDTOFieldName.JournalType)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, model.JournalType(ledgerJournalEntries.JournalType), explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.SourceType):
			key := string(LedgerJournalEntriesDTOFieldName.SourceType)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.SourceType, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.SourceId):
			key := string(LedgerJournalEntriesDTOFieldName.SourceId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.SourceId, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.IdempotencyKey):
			key := string(LedgerJournalEntriesDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.IdempotencyKey, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.JournalStatus):
			key := string(LedgerJournalEntriesDTOFieldName.JournalStatus)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, model.JournalStatus(ledgerJournalEntries.JournalStatus), explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.EffectiveAt):
			key := string(LedgerJournalEntriesDTOFieldName.EffectiveAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.EffectiveAt, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.BookedAt):
			key := string(LedgerJournalEntriesDTOFieldName.BookedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.BookedAt, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.Description):
			key := string(LedgerJournalEntriesDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.Description.String, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.ReversalOfJournalId):
			key := string(LedgerJournalEntriesDTOFieldName.ReversalOfJournalId)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.ReversalOfJournalId.UUID, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.Metadata):
			key := string(LedgerJournalEntriesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.Metadata, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.MetaCreatedAt):
			key := string(LedgerJournalEntriesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.MetaCreatedAt, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.MetaCreatedBy):
			key := string(LedgerJournalEntriesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.MetaCreatedBy, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.MetaUpdatedAt):
			key := string(LedgerJournalEntriesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.MetaUpdatedBy):
			key := string(LedgerJournalEntriesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.MetaDeletedAt):
			key := string(LedgerJournalEntriesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerJournalEntriesDBFieldName.MetaDeletedBy):
			key := string(LedgerJournalEntriesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalEntriesSelectableValue(ledgerJournalEntriesSelectableResponse, key, ledgerJournalEntries.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerJournalEntriesSelectableResponse
}

func NewLedgerJournalEntriesListResponseFromFilterResult(result []model.LedgerJournalEntriesFilterResult, filter model.Filter) LedgerJournalEntriesSelectableListResponse {
	dtoLedgerJournalEntriesListResponse := LedgerJournalEntriesSelectableListResponse{}
	for _, row := range result {
		dtoLedgerJournalEntriesResponse := NewLedgerJournalEntriesSelectableResponse(row.LedgerJournalEntries, filter)
		dtoLedgerJournalEntriesListResponse = append(dtoLedgerJournalEntriesListResponse, &dtoLedgerJournalEntriesResponse)
	}
	return dtoLedgerJournalEntriesListResponse
}

type LedgerJournalEntriesFilterResponse struct {
	Metadata Metadata                                   `json:"metadata"`
	Data     LedgerJournalEntriesSelectableListResponse `json:"data"`
}

func reverseLedgerJournalEntriesFilterResults(result []model.LedgerJournalEntriesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerJournalEntriesFilterResponse(result []model.LedgerJournalEntriesFilterResult, filter model.Filter) (resp LedgerJournalEntriesFilterResponse) {
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
			reverseLedgerJournalEntriesFilterResults(dataResult)
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

	resp.Data = NewLedgerJournalEntriesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerJournalEntriesCreateRequest struct {
	BatchId             uuid.UUID           `json:"batchId"`
	JournalCode         string              `json:"journalCode"`
	BookId              uuid.UUID           `json:"bookId"`
	JournalType         model.JournalType   `json:"journalType" example:"payment" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation,revaluation"`
	SourceType          string              `json:"sourceType"`
	SourceId            uuid.UUID           `json:"sourceId"`
	IdempotencyKey      string              `json:"idempotencyKey"`
	JournalStatus       model.JournalStatus `json:"journalStatus" example:"draft" enums:"draft,validated,posted,reversed,failed"`
	EffectiveAt         time.Time           `json:"effectiveAt"`
	BookedAt            time.Time           `json:"bookedAt"`
	Description         string              `json:"description"`
	ReversalOfJournalId uuid.UUID           `json:"reversalOfJournalId"`
	Metadata            json.RawMessage     `json:"metadata"`
}

func (d *LedgerJournalEntriesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerJournalEntriesCreateRequest) ToModel() model.LedgerJournalEntries {
	id, _ := uuid.NewV4()
	return model.LedgerJournalEntries{
		Id:                  id,
		BatchId:             nuuid.From(d.BatchId),
		JournalCode:         d.JournalCode,
		BookId:              d.BookId,
		JournalType:         d.JournalType,
		SourceType:          d.SourceType,
		SourceId:            d.SourceId,
		IdempotencyKey:      d.IdempotencyKey,
		JournalStatus:       d.JournalStatus,
		EffectiveAt:         d.EffectiveAt,
		BookedAt:            d.BookedAt,
		Description:         null.StringFrom(d.Description),
		ReversalOfJournalId: nuuid.From(d.ReversalOfJournalId),
		Metadata:            d.Metadata,
	}
}

type LedgerJournalEntriesListCreateRequest []*LedgerJournalEntriesCreateRequest

func (d LedgerJournalEntriesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalEntries := range d {
		err = validator.Struct(ledgerJournalEntries)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerJournalEntriesListCreateRequest) ToModelList() []model.LedgerJournalEntries {
	out := make([]model.LedgerJournalEntries, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerJournalEntriesUpdateRequest struct {
	BatchId             uuid.UUID           `json:"batchId"`
	JournalCode         string              `json:"journalCode"`
	BookId              uuid.UUID           `json:"bookId"`
	JournalType         model.JournalType   `json:"journalType" example:"payment" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation,revaluation"`
	SourceType          string              `json:"sourceType"`
	SourceId            uuid.UUID           `json:"sourceId"`
	IdempotencyKey      string              `json:"idempotencyKey"`
	JournalStatus       model.JournalStatus `json:"journalStatus" example:"draft" enums:"draft,validated,posted,reversed,failed"`
	EffectiveAt         time.Time           `json:"effectiveAt"`
	BookedAt            time.Time           `json:"bookedAt"`
	Description         string              `json:"description"`
	ReversalOfJournalId uuid.UUID           `json:"reversalOfJournalId"`
	Metadata            json.RawMessage     `json:"metadata"`
}

func (d *LedgerJournalEntriesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerJournalEntriesUpdateRequest) ToModel() model.LedgerJournalEntries {
	return model.LedgerJournalEntries{
		BatchId:             nuuid.From(d.BatchId),
		JournalCode:         d.JournalCode,
		BookId:              d.BookId,
		JournalType:         d.JournalType,
		SourceType:          d.SourceType,
		SourceId:            d.SourceId,
		IdempotencyKey:      d.IdempotencyKey,
		JournalStatus:       d.JournalStatus,
		EffectiveAt:         d.EffectiveAt,
		BookedAt:            d.BookedAt,
		Description:         null.StringFrom(d.Description),
		ReversalOfJournalId: nuuid.From(d.ReversalOfJournalId),
		Metadata:            d.Metadata,
	}
}

type LedgerJournalEntriesBulkUpdateRequest struct {
	Id                  uuid.UUID           `json:"id"`
	BatchId             uuid.UUID           `json:"batchId"`
	JournalCode         string              `json:"journalCode"`
	BookId              uuid.UUID           `json:"bookId"`
	JournalType         model.JournalType   `json:"journalType" example:"payment" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation,revaluation"`
	SourceType          string              `json:"sourceType"`
	SourceId            uuid.UUID           `json:"sourceId"`
	IdempotencyKey      string              `json:"idempotencyKey"`
	JournalStatus       model.JournalStatus `json:"journalStatus" example:"draft" enums:"draft,validated,posted,reversed,failed"`
	EffectiveAt         time.Time           `json:"effectiveAt"`
	BookedAt            time.Time           `json:"bookedAt"`
	Description         string              `json:"description"`
	ReversalOfJournalId uuid.UUID           `json:"reversalOfJournalId"`
	Metadata            json.RawMessage     `json:"metadata"`
}

func (d LedgerJournalEntriesBulkUpdateRequest) PrimaryID() LedgerJournalEntriesPrimaryID {
	return LedgerJournalEntriesPrimaryID{
		Id: d.Id,
	}
}

type LedgerJournalEntriesListBulkUpdateRequest []*LedgerJournalEntriesBulkUpdateRequest

func (d LedgerJournalEntriesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalEntries := range d {
		err = validator.Struct(ledgerJournalEntries)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerJournalEntriesBulkUpdateRequest) ToModel() model.LedgerJournalEntries {
	return model.LedgerJournalEntries{
		Id:                  d.Id,
		BatchId:             nuuid.From(d.BatchId),
		JournalCode:         d.JournalCode,
		BookId:              d.BookId,
		JournalType:         d.JournalType,
		SourceType:          d.SourceType,
		SourceId:            d.SourceId,
		IdempotencyKey:      d.IdempotencyKey,
		JournalStatus:       d.JournalStatus,
		EffectiveAt:         d.EffectiveAt,
		BookedAt:            d.BookedAt,
		Description:         null.StringFrom(d.Description),
		ReversalOfJournalId: nuuid.From(d.ReversalOfJournalId),
		Metadata:            d.Metadata,
	}
}

type LedgerJournalEntriesResponse struct {
	Id                  uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BatchId             uuid.UUID           `json:"batchId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	JournalCode         string              `json:"journalCode" validate:"required"`
	BookId              uuid.UUID           `json:"bookId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	JournalType         model.JournalType   `json:"journalType" validate:"required,oneof=payment refund settlement payout chargeback adjustment close reconciliation revaluation" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation,revaluation"`
	SourceType          string              `json:"sourceType" validate:"required"`
	SourceId            uuid.UUID           `json:"sourceId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IdempotencyKey      string              `json:"idempotencyKey" validate:"required"`
	JournalStatus       model.JournalStatus `json:"journalStatus" validate:"oneof=draft validated posted reversed failed" enums:"draft,validated,posted,reversed,failed"`
	EffectiveAt         time.Time           `json:"effectiveAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	BookedAt            time.Time           `json:"bookedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Description         string              `json:"description"`
	ReversalOfJournalId uuid.UUID           `json:"reversalOfJournalId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Metadata            json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewLedgerJournalEntriesResponse(ledgerJournalEntries model.LedgerJournalEntries) LedgerJournalEntriesResponse {
	return LedgerJournalEntriesResponse{
		Id:                  ledgerJournalEntries.Id,
		BatchId:             ledgerJournalEntries.BatchId.UUID,
		JournalCode:         ledgerJournalEntries.JournalCode,
		BookId:              ledgerJournalEntries.BookId,
		JournalType:         model.JournalType(ledgerJournalEntries.JournalType),
		SourceType:          ledgerJournalEntries.SourceType,
		SourceId:            ledgerJournalEntries.SourceId,
		IdempotencyKey:      ledgerJournalEntries.IdempotencyKey,
		JournalStatus:       model.JournalStatus(ledgerJournalEntries.JournalStatus),
		EffectiveAt:         ledgerJournalEntries.EffectiveAt,
		BookedAt:            ledgerJournalEntries.BookedAt,
		Description:         ledgerJournalEntries.Description.String,
		ReversalOfJournalId: ledgerJournalEntries.ReversalOfJournalId.UUID,
		Metadata:            ledgerJournalEntries.Metadata,
	}
}

type LedgerJournalEntriesListResponse []*LedgerJournalEntriesResponse

func NewLedgerJournalEntriesListResponse(ledgerJournalEntriesList model.LedgerJournalEntriesList) LedgerJournalEntriesListResponse {
	dtoLedgerJournalEntriesListResponse := LedgerJournalEntriesListResponse{}
	for _, ledgerJournalEntries := range ledgerJournalEntriesList {
		dtoLedgerJournalEntriesResponse := NewLedgerJournalEntriesResponse(*ledgerJournalEntries)
		dtoLedgerJournalEntriesListResponse = append(dtoLedgerJournalEntriesListResponse, &dtoLedgerJournalEntriesResponse)
	}
	return dtoLedgerJournalEntriesListResponse
}

type LedgerJournalEntriesPrimaryIDList []LedgerJournalEntriesPrimaryID

func (d LedgerJournalEntriesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalEntries := range d {
		err = validator.Struct(ledgerJournalEntries)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerJournalEntriesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerJournalEntriesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerJournalEntriesPrimaryID) ToModel() model.LedgerJournalEntriesPrimaryID {
	return model.LedgerJournalEntriesPrimaryID{
		Id: d.Id,
	}
}
