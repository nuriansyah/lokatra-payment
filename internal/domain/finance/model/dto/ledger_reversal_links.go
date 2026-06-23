package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type LedgerReversalLinksDTOFieldNameType string

type ledgerReversalLinksDTOFieldName struct {
	Id                LedgerReversalLinksDTOFieldNameType
	OriginalJournalId LedgerReversalLinksDTOFieldNameType
	ReversalJournalId LedgerReversalLinksDTOFieldNameType
	ReversalReason    LedgerReversalLinksDTOFieldNameType
	ReversedAt        LedgerReversalLinksDTOFieldNameType
	Metadata          LedgerReversalLinksDTOFieldNameType
	MetaCreatedAt     LedgerReversalLinksDTOFieldNameType
	MetaCreatedBy     LedgerReversalLinksDTOFieldNameType
	MetaUpdatedAt     LedgerReversalLinksDTOFieldNameType
	MetaUpdatedBy     LedgerReversalLinksDTOFieldNameType
	MetaDeletedAt     LedgerReversalLinksDTOFieldNameType
	MetaDeletedBy     LedgerReversalLinksDTOFieldNameType
}

var LedgerReversalLinksDTOFieldName = ledgerReversalLinksDTOFieldName{
	Id:                "id",
	OriginalJournalId: "originalJournalId",
	ReversalJournalId: "reversalJournalId",
	ReversalReason:    "reversalReason",
	ReversedAt:        "reversedAt",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformLedgerReversalLinksDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerReversalLinksDTOFieldName.Id):
		return string(model.LedgerReversalLinksDBFieldName.Id), true

	case string(LedgerReversalLinksDTOFieldName.OriginalJournalId):
		return string(model.LedgerReversalLinksDBFieldName.OriginalJournalId), true

	case string(LedgerReversalLinksDTOFieldName.ReversalJournalId):
		return string(model.LedgerReversalLinksDBFieldName.ReversalJournalId), true

	case string(LedgerReversalLinksDTOFieldName.ReversalReason):
		return string(model.LedgerReversalLinksDBFieldName.ReversalReason), true

	case string(LedgerReversalLinksDTOFieldName.ReversedAt):
		return string(model.LedgerReversalLinksDBFieldName.ReversedAt), true

	case string(LedgerReversalLinksDTOFieldName.Metadata):
		return string(model.LedgerReversalLinksDBFieldName.Metadata), true

	case string(LedgerReversalLinksDTOFieldName.MetaCreatedAt):
		return string(model.LedgerReversalLinksDBFieldName.MetaCreatedAt), true

	case string(LedgerReversalLinksDTOFieldName.MetaCreatedBy):
		return string(model.LedgerReversalLinksDBFieldName.MetaCreatedBy), true

	case string(LedgerReversalLinksDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerReversalLinksDBFieldName.MetaUpdatedAt), true

	case string(LedgerReversalLinksDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerReversalLinksDBFieldName.MetaUpdatedBy), true

	case string(LedgerReversalLinksDTOFieldName.MetaDeletedAt):
		return string(model.LedgerReversalLinksDBFieldName.MetaDeletedAt), true

	case string(LedgerReversalLinksDTOFieldName.MetaDeletedBy):
		return string(model.LedgerReversalLinksDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerReversalLinksFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerReversalLinksBaseFilterField(field string) bool {
	spec, found := model.NewLedgerReversalLinksFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerReversalLinksProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerReversalLinksProjectionOutputPath(path string) error {
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

func transformLedgerReversalLinksFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerReversalLinksDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerReversalLinksFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerReversalLinksFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerReversalLinksDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerReversalLinksBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerReversalLinksProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerReversalLinksProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerReversalLinksDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerReversalLinksDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerReversalLinksFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerReversalLinksFilter(filter *model.Filter) {
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
			Field: string(LedgerReversalLinksDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerReversalLinksSelectableResponse map[string]interface{}
type LedgerReversalLinksSelectableListResponse []*LedgerReversalLinksSelectableResponse

func assignLedgerReversalLinksNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerReversalLinksSelectableValue(out LedgerReversalLinksSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerReversalLinksNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerReversalLinksSelectableResponse(ledgerReversalLinks model.LedgerReversalLinks, filter model.Filter) LedgerReversalLinksSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerReversalLinksDBFieldName.Id),
			string(model.LedgerReversalLinksDBFieldName.OriginalJournalId),
			string(model.LedgerReversalLinksDBFieldName.ReversalJournalId),
			string(model.LedgerReversalLinksDBFieldName.ReversalReason),
			string(model.LedgerReversalLinksDBFieldName.ReversedAt),
			string(model.LedgerReversalLinksDBFieldName.Metadata),
			string(model.LedgerReversalLinksDBFieldName.MetaCreatedAt),
			string(model.LedgerReversalLinksDBFieldName.MetaCreatedBy),
			string(model.LedgerReversalLinksDBFieldName.MetaUpdatedAt),
			string(model.LedgerReversalLinksDBFieldName.MetaUpdatedBy),
			string(model.LedgerReversalLinksDBFieldName.MetaDeletedAt),
			string(model.LedgerReversalLinksDBFieldName.MetaDeletedBy),
		)
	}
	ledgerReversalLinksSelectableResponse := LedgerReversalLinksSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerReversalLinksDBFieldName.Id):
			key := string(LedgerReversalLinksDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.Id, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.OriginalJournalId):
			key := string(LedgerReversalLinksDTOFieldName.OriginalJournalId)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.OriginalJournalId, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.ReversalJournalId):
			key := string(LedgerReversalLinksDTOFieldName.ReversalJournalId)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.ReversalJournalId, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.ReversalReason):
			key := string(LedgerReversalLinksDTOFieldName.ReversalReason)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.ReversalReason, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.ReversedAt):
			key := string(LedgerReversalLinksDTOFieldName.ReversedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.ReversedAt, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.Metadata):
			key := string(LedgerReversalLinksDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.Metadata, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.MetaCreatedAt):
			key := string(LedgerReversalLinksDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.MetaCreatedAt, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.MetaCreatedBy):
			key := string(LedgerReversalLinksDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.MetaCreatedBy, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.MetaUpdatedAt):
			key := string(LedgerReversalLinksDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.MetaUpdatedBy):
			key := string(LedgerReversalLinksDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.MetaDeletedAt):
			key := string(LedgerReversalLinksDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerReversalLinksDBFieldName.MetaDeletedBy):
			key := string(LedgerReversalLinksDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerReversalLinksSelectableValue(ledgerReversalLinksSelectableResponse, key, ledgerReversalLinks.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerReversalLinksSelectableResponse
}

func NewLedgerReversalLinksListResponseFromFilterResult(result []model.LedgerReversalLinksFilterResult, filter model.Filter) LedgerReversalLinksSelectableListResponse {
	dtoLedgerReversalLinksListResponse := LedgerReversalLinksSelectableListResponse{}
	for _, row := range result {
		dtoLedgerReversalLinksResponse := NewLedgerReversalLinksSelectableResponse(row.LedgerReversalLinks, filter)
		dtoLedgerReversalLinksListResponse = append(dtoLedgerReversalLinksListResponse, &dtoLedgerReversalLinksResponse)
	}
	return dtoLedgerReversalLinksListResponse
}

type LedgerReversalLinksFilterResponse struct {
	Metadata Metadata                                  `json:"metadata"`
	Data     LedgerReversalLinksSelectableListResponse `json:"data"`
}

func reverseLedgerReversalLinksFilterResults(result []model.LedgerReversalLinksFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerReversalLinksFilterResponse(result []model.LedgerReversalLinksFilterResult, filter model.Filter) (resp LedgerReversalLinksFilterResponse) {
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
			reverseLedgerReversalLinksFilterResults(dataResult)
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

	resp.Data = NewLedgerReversalLinksListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerReversalLinksCreateRequest struct {
	OriginalJournalId uuid.UUID       `json:"originalJournalId"`
	ReversalJournalId uuid.UUID       `json:"reversalJournalId"`
	ReversalReason    string          `json:"reversalReason"`
	ReversedAt        time.Time       `json:"reversedAt"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *LedgerReversalLinksCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerReversalLinksCreateRequest) ToModel() model.LedgerReversalLinks {
	id, _ := uuid.NewV4()
	return model.LedgerReversalLinks{
		Id:                id,
		OriginalJournalId: d.OriginalJournalId,
		ReversalJournalId: d.ReversalJournalId,
		ReversalReason:    d.ReversalReason,
		ReversedAt:        d.ReversedAt,
		Metadata:          d.Metadata,
	}
}

type LedgerReversalLinksListCreateRequest []*LedgerReversalLinksCreateRequest

func (d LedgerReversalLinksListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerReversalLinks := range d {
		err = validator.Struct(ledgerReversalLinks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerReversalLinksListCreateRequest) ToModelList() []model.LedgerReversalLinks {
	out := make([]model.LedgerReversalLinks, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerReversalLinksUpdateRequest struct {
	OriginalJournalId uuid.UUID       `json:"originalJournalId"`
	ReversalJournalId uuid.UUID       `json:"reversalJournalId"`
	ReversalReason    string          `json:"reversalReason"`
	ReversedAt        time.Time       `json:"reversedAt"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d *LedgerReversalLinksUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerReversalLinksUpdateRequest) ToModel() model.LedgerReversalLinks {
	return model.LedgerReversalLinks{
		OriginalJournalId: d.OriginalJournalId,
		ReversalJournalId: d.ReversalJournalId,
		ReversalReason:    d.ReversalReason,
		ReversedAt:        d.ReversedAt,
		Metadata:          d.Metadata,
	}
}

type LedgerReversalLinksBulkUpdateRequest struct {
	Id                uuid.UUID       `json:"id"`
	OriginalJournalId uuid.UUID       `json:"originalJournalId"`
	ReversalJournalId uuid.UUID       `json:"reversalJournalId"`
	ReversalReason    string          `json:"reversalReason"`
	ReversedAt        time.Time       `json:"reversedAt"`
	Metadata          json.RawMessage `json:"metadata"`
}

func (d LedgerReversalLinksBulkUpdateRequest) PrimaryID() LedgerReversalLinksPrimaryID {
	return LedgerReversalLinksPrimaryID{
		Id: d.Id,
	}
}

type LedgerReversalLinksListBulkUpdateRequest []*LedgerReversalLinksBulkUpdateRequest

func (d LedgerReversalLinksListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerReversalLinks := range d {
		err = validator.Struct(ledgerReversalLinks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerReversalLinksBulkUpdateRequest) ToModel() model.LedgerReversalLinks {
	return model.LedgerReversalLinks{
		Id:                d.Id,
		OriginalJournalId: d.OriginalJournalId,
		ReversalJournalId: d.ReversalJournalId,
		ReversalReason:    d.ReversalReason,
		ReversedAt:        d.ReversedAt,
		Metadata:          d.Metadata,
	}
}

type LedgerReversalLinksResponse struct {
	Id                uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	OriginalJournalId uuid.UUID       `json:"originalJournalId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReversalJournalId uuid.UUID       `json:"reversalJournalId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReversalReason    string          `json:"reversalReason" validate:"required"`
	ReversedAt        time.Time       `json:"reversedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata          json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewLedgerReversalLinksResponse(ledgerReversalLinks model.LedgerReversalLinks) LedgerReversalLinksResponse {
	return LedgerReversalLinksResponse{
		Id:                ledgerReversalLinks.Id,
		OriginalJournalId: ledgerReversalLinks.OriginalJournalId,
		ReversalJournalId: ledgerReversalLinks.ReversalJournalId,
		ReversalReason:    ledgerReversalLinks.ReversalReason,
		ReversedAt:        ledgerReversalLinks.ReversedAt,
		Metadata:          ledgerReversalLinks.Metadata,
	}
}

type LedgerReversalLinksListResponse []*LedgerReversalLinksResponse

func NewLedgerReversalLinksListResponse(ledgerReversalLinksList model.LedgerReversalLinksList) LedgerReversalLinksListResponse {
	dtoLedgerReversalLinksListResponse := LedgerReversalLinksListResponse{}
	for _, ledgerReversalLinks := range ledgerReversalLinksList {
		dtoLedgerReversalLinksResponse := NewLedgerReversalLinksResponse(*ledgerReversalLinks)
		dtoLedgerReversalLinksListResponse = append(dtoLedgerReversalLinksListResponse, &dtoLedgerReversalLinksResponse)
	}
	return dtoLedgerReversalLinksListResponse
}

type LedgerReversalLinksPrimaryIDList []LedgerReversalLinksPrimaryID

func (d LedgerReversalLinksPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerReversalLinks := range d {
		err = validator.Struct(ledgerReversalLinks)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerReversalLinksPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerReversalLinksPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerReversalLinksPrimaryID) ToModel() model.LedgerReversalLinksPrimaryID {
	return model.LedgerReversalLinksPrimaryID{
		Id: d.Id,
	}
}
