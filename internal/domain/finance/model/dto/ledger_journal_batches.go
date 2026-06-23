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

type LedgerJournalBatchesDTOFieldNameType string

type ledgerJournalBatchesDTOFieldName struct {
	Id            LedgerJournalBatchesDTOFieldNameType
	BatchCode     LedgerJournalBatchesDTOFieldNameType
	BatchType     LedgerJournalBatchesDTOFieldNameType
	SourceRef     LedgerJournalBatchesDTOFieldNameType
	BatchStatus   LedgerJournalBatchesDTOFieldNameType
	BookedAt      LedgerJournalBatchesDTOFieldNameType
	Description   LedgerJournalBatchesDTOFieldNameType
	Metadata      LedgerJournalBatchesDTOFieldNameType
	MetaCreatedAt LedgerJournalBatchesDTOFieldNameType
	MetaCreatedBy LedgerJournalBatchesDTOFieldNameType
	MetaUpdatedAt LedgerJournalBatchesDTOFieldNameType
	MetaUpdatedBy LedgerJournalBatchesDTOFieldNameType
	MetaDeletedAt LedgerJournalBatchesDTOFieldNameType
	MetaDeletedBy LedgerJournalBatchesDTOFieldNameType
}

var LedgerJournalBatchesDTOFieldName = ledgerJournalBatchesDTOFieldName{
	Id:            "id",
	BatchCode:     "batchCode",
	BatchType:     "batchType",
	SourceRef:     "sourceRef",
	BatchStatus:   "batchStatus",
	BookedAt:      "bookedAt",
	Description:   "description",
	Metadata:      "metadata",
	MetaCreatedAt: "metaCreatedAt",
	MetaCreatedBy: "metaCreatedBy",
	MetaUpdatedAt: "metaUpdatedAt",
	MetaUpdatedBy: "metaUpdatedBy",
	MetaDeletedAt: "metaDeletedAt",
	MetaDeletedBy: "metaDeletedBy",
}

func transformLedgerJournalBatchesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(LedgerJournalBatchesDTOFieldName.Id):
		return string(model.LedgerJournalBatchesDBFieldName.Id), true

	case string(LedgerJournalBatchesDTOFieldName.BatchCode):
		return string(model.LedgerJournalBatchesDBFieldName.BatchCode), true

	case string(LedgerJournalBatchesDTOFieldName.BatchType):
		return string(model.LedgerJournalBatchesDBFieldName.BatchType), true

	case string(LedgerJournalBatchesDTOFieldName.SourceRef):
		return string(model.LedgerJournalBatchesDBFieldName.SourceRef), true

	case string(LedgerJournalBatchesDTOFieldName.BatchStatus):
		return string(model.LedgerJournalBatchesDBFieldName.BatchStatus), true

	case string(LedgerJournalBatchesDTOFieldName.BookedAt):
		return string(model.LedgerJournalBatchesDBFieldName.BookedAt), true

	case string(LedgerJournalBatchesDTOFieldName.Description):
		return string(model.LedgerJournalBatchesDBFieldName.Description), true

	case string(LedgerJournalBatchesDTOFieldName.Metadata):
		return string(model.LedgerJournalBatchesDBFieldName.Metadata), true

	case string(LedgerJournalBatchesDTOFieldName.MetaCreatedAt):
		return string(model.LedgerJournalBatchesDBFieldName.MetaCreatedAt), true

	case string(LedgerJournalBatchesDTOFieldName.MetaCreatedBy):
		return string(model.LedgerJournalBatchesDBFieldName.MetaCreatedBy), true

	case string(LedgerJournalBatchesDTOFieldName.MetaUpdatedAt):
		return string(model.LedgerJournalBatchesDBFieldName.MetaUpdatedAt), true

	case string(LedgerJournalBatchesDTOFieldName.MetaUpdatedBy):
		return string(model.LedgerJournalBatchesDBFieldName.MetaUpdatedBy), true

	case string(LedgerJournalBatchesDTOFieldName.MetaDeletedAt):
		return string(model.LedgerJournalBatchesDBFieldName.MetaDeletedAt), true

	case string(LedgerJournalBatchesDTOFieldName.MetaDeletedBy):
		return string(model.LedgerJournalBatchesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewLedgerJournalBatchesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isLedgerJournalBatchesBaseFilterField(field string) bool {
	spec, found := model.NewLedgerJournalBatchesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeLedgerJournalBatchesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateLedgerJournalBatchesProjectionOutputPath(path string) error {
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

func transformLedgerJournalBatchesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformLedgerJournalBatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformLedgerJournalBatchesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformLedgerJournalBatchesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformLedgerJournalBatchesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isLedgerJournalBatchesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateLedgerJournalBatchesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeLedgerJournalBatchesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformLedgerJournalBatchesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformLedgerJournalBatchesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformLedgerJournalBatchesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultLedgerJournalBatchesFilter(filter *model.Filter) {
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
			Field: string(LedgerJournalBatchesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type LedgerJournalBatchesSelectableResponse map[string]interface{}
type LedgerJournalBatchesSelectableListResponse []*LedgerJournalBatchesSelectableResponse

func assignLedgerJournalBatchesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setLedgerJournalBatchesSelectableValue(out LedgerJournalBatchesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignLedgerJournalBatchesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewLedgerJournalBatchesSelectableResponse(ledgerJournalBatches model.LedgerJournalBatches, filter model.Filter) LedgerJournalBatchesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.LedgerJournalBatchesDBFieldName.Id),
			string(model.LedgerJournalBatchesDBFieldName.BatchCode),
			string(model.LedgerJournalBatchesDBFieldName.BatchType),
			string(model.LedgerJournalBatchesDBFieldName.SourceRef),
			string(model.LedgerJournalBatchesDBFieldName.BatchStatus),
			string(model.LedgerJournalBatchesDBFieldName.BookedAt),
			string(model.LedgerJournalBatchesDBFieldName.Description),
			string(model.LedgerJournalBatchesDBFieldName.Metadata),
			string(model.LedgerJournalBatchesDBFieldName.MetaCreatedAt),
			string(model.LedgerJournalBatchesDBFieldName.MetaCreatedBy),
			string(model.LedgerJournalBatchesDBFieldName.MetaUpdatedAt),
			string(model.LedgerJournalBatchesDBFieldName.MetaUpdatedBy),
			string(model.LedgerJournalBatchesDBFieldName.MetaDeletedAt),
			string(model.LedgerJournalBatchesDBFieldName.MetaDeletedBy),
		)
	}
	ledgerJournalBatchesSelectableResponse := LedgerJournalBatchesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.LedgerJournalBatchesDBFieldName.Id):
			key := string(LedgerJournalBatchesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.Id, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.BatchCode):
			key := string(LedgerJournalBatchesDTOFieldName.BatchCode)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.BatchCode, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.BatchType):
			key := string(LedgerJournalBatchesDTOFieldName.BatchType)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, model.BatchType(ledgerJournalBatches.BatchType), explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.SourceRef):
			key := string(LedgerJournalBatchesDTOFieldName.SourceRef)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.SourceRef.String, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.BatchStatus):
			key := string(LedgerJournalBatchesDTOFieldName.BatchStatus)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, model.LedgerJournalBatchesBatchStatus(ledgerJournalBatches.BatchStatus), explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.BookedAt):
			key := string(LedgerJournalBatchesDTOFieldName.BookedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.BookedAt, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.Description):
			key := string(LedgerJournalBatchesDTOFieldName.Description)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.Description.String, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.Metadata):
			key := string(LedgerJournalBatchesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.Metadata, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.MetaCreatedAt):
			key := string(LedgerJournalBatchesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.MetaCreatedAt, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.MetaCreatedBy):
			key := string(LedgerJournalBatchesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.MetaCreatedBy, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.MetaUpdatedAt):
			key := string(LedgerJournalBatchesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.MetaUpdatedAt, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.MetaUpdatedBy):
			key := string(LedgerJournalBatchesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.MetaUpdatedBy, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.MetaDeletedAt):
			key := string(LedgerJournalBatchesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.MetaDeletedAt.Time, explicitAlias)

		case string(model.LedgerJournalBatchesDBFieldName.MetaDeletedBy):
			key := string(LedgerJournalBatchesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setLedgerJournalBatchesSelectableValue(ledgerJournalBatchesSelectableResponse, key, ledgerJournalBatches.MetaDeletedBy, explicitAlias)

		}
	}
	return ledgerJournalBatchesSelectableResponse
}

func NewLedgerJournalBatchesListResponseFromFilterResult(result []model.LedgerJournalBatchesFilterResult, filter model.Filter) LedgerJournalBatchesSelectableListResponse {
	dtoLedgerJournalBatchesListResponse := LedgerJournalBatchesSelectableListResponse{}
	for _, row := range result {
		dtoLedgerJournalBatchesResponse := NewLedgerJournalBatchesSelectableResponse(row.LedgerJournalBatches, filter)
		dtoLedgerJournalBatchesListResponse = append(dtoLedgerJournalBatchesListResponse, &dtoLedgerJournalBatchesResponse)
	}
	return dtoLedgerJournalBatchesListResponse
}

type LedgerJournalBatchesFilterResponse struct {
	Metadata Metadata                                   `json:"metadata"`
	Data     LedgerJournalBatchesSelectableListResponse `json:"data"`
}

func reverseLedgerJournalBatchesFilterResults(result []model.LedgerJournalBatchesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewLedgerJournalBatchesFilterResponse(result []model.LedgerJournalBatchesFilterResult, filter model.Filter) (resp LedgerJournalBatchesFilterResponse) {
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
			reverseLedgerJournalBatchesFilterResults(dataResult)
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

	resp.Data = NewLedgerJournalBatchesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type LedgerJournalBatchesCreateRequest struct {
	BatchCode   string                                `json:"batchCode"`
	BatchType   model.BatchType                       `json:"batchType" example:"payment" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation"`
	SourceRef   string                                `json:"sourceRef"`
	BatchStatus model.LedgerJournalBatchesBatchStatus `json:"batchStatus" example:"open" enums:"open,posting,posted,failed,closed"`
	BookedAt    time.Time                             `json:"bookedAt"`
	Description string                                `json:"description"`
	Metadata    json.RawMessage                       `json:"metadata"`
}

func (d *LedgerJournalBatchesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *LedgerJournalBatchesCreateRequest) ToModel() model.LedgerJournalBatches {
	id, _ := uuid.NewV4()
	return model.LedgerJournalBatches{
		Id:          id,
		BatchCode:   d.BatchCode,
		BatchType:   d.BatchType,
		SourceRef:   null.StringFrom(d.SourceRef),
		BatchStatus: d.BatchStatus,
		BookedAt:    d.BookedAt,
		Description: null.StringFrom(d.Description),
		Metadata:    d.Metadata,
	}
}

type LedgerJournalBatchesListCreateRequest []*LedgerJournalBatchesCreateRequest

func (d LedgerJournalBatchesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalBatches := range d {
		err = validator.Struct(ledgerJournalBatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerJournalBatchesListCreateRequest) ToModelList() []model.LedgerJournalBatches {
	out := make([]model.LedgerJournalBatches, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type LedgerJournalBatchesUpdateRequest struct {
	BatchCode   string                                `json:"batchCode"`
	BatchType   model.BatchType                       `json:"batchType" example:"payment" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation"`
	SourceRef   string                                `json:"sourceRef"`
	BatchStatus model.LedgerJournalBatchesBatchStatus `json:"batchStatus" example:"open" enums:"open,posting,posted,failed,closed"`
	BookedAt    time.Time                             `json:"bookedAt"`
	Description string                                `json:"description"`
	Metadata    json.RawMessage                       `json:"metadata"`
}

func (d *LedgerJournalBatchesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d LedgerJournalBatchesUpdateRequest) ToModel() model.LedgerJournalBatches {
	return model.LedgerJournalBatches{
		BatchCode:   d.BatchCode,
		BatchType:   d.BatchType,
		SourceRef:   null.StringFrom(d.SourceRef),
		BatchStatus: d.BatchStatus,
		BookedAt:    d.BookedAt,
		Description: null.StringFrom(d.Description),
		Metadata:    d.Metadata,
	}
}

type LedgerJournalBatchesBulkUpdateRequest struct {
	Id          uuid.UUID                             `json:"id"`
	BatchCode   string                                `json:"batchCode"`
	BatchType   model.BatchType                       `json:"batchType" example:"payment" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation"`
	SourceRef   string                                `json:"sourceRef"`
	BatchStatus model.LedgerJournalBatchesBatchStatus `json:"batchStatus" example:"open" enums:"open,posting,posted,failed,closed"`
	BookedAt    time.Time                             `json:"bookedAt"`
	Description string                                `json:"description"`
	Metadata    json.RawMessage                       `json:"metadata"`
}

func (d LedgerJournalBatchesBulkUpdateRequest) PrimaryID() LedgerJournalBatchesPrimaryID {
	return LedgerJournalBatchesPrimaryID{
		Id: d.Id,
	}
}

type LedgerJournalBatchesListBulkUpdateRequest []*LedgerJournalBatchesBulkUpdateRequest

func (d LedgerJournalBatchesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalBatches := range d {
		err = validator.Struct(ledgerJournalBatches)
		if err != nil {
			return
		}
	}
	return nil
}

func (d LedgerJournalBatchesBulkUpdateRequest) ToModel() model.LedgerJournalBatches {
	return model.LedgerJournalBatches{
		Id:          d.Id,
		BatchCode:   d.BatchCode,
		BatchType:   d.BatchType,
		SourceRef:   null.StringFrom(d.SourceRef),
		BatchStatus: d.BatchStatus,
		BookedAt:    d.BookedAt,
		Description: null.StringFrom(d.Description),
		Metadata:    d.Metadata,
	}
}

type LedgerJournalBatchesResponse struct {
	Id          uuid.UUID                             `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BatchCode   string                                `json:"batchCode" validate:"required"`
	BatchType   model.BatchType                       `json:"batchType" validate:"required,oneof=payment refund settlement payout chargeback adjustment close reconciliation" enums:"payment,refund,settlement,payout,chargeback,adjustment,close,reconciliation"`
	SourceRef   string                                `json:"sourceRef"`
	BatchStatus model.LedgerJournalBatchesBatchStatus `json:"batchStatus" validate:"oneof=open posting posted failed closed" enums:"open,posting,posted,failed,closed"`
	BookedAt    time.Time                             `json:"bookedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Description string                                `json:"description"`
	Metadata    json.RawMessage                       `json:"metadata" swaggertype:"object"`
}

func NewLedgerJournalBatchesResponse(ledgerJournalBatches model.LedgerJournalBatches) LedgerJournalBatchesResponse {
	return LedgerJournalBatchesResponse{
		Id:          ledgerJournalBatches.Id,
		BatchCode:   ledgerJournalBatches.BatchCode,
		BatchType:   model.BatchType(ledgerJournalBatches.BatchType),
		SourceRef:   ledgerJournalBatches.SourceRef.String,
		BatchStatus: model.LedgerJournalBatchesBatchStatus(ledgerJournalBatches.BatchStatus),
		BookedAt:    ledgerJournalBatches.BookedAt,
		Description: ledgerJournalBatches.Description.String,
		Metadata:    ledgerJournalBatches.Metadata,
	}
}

type LedgerJournalBatchesListResponse []*LedgerJournalBatchesResponse

func NewLedgerJournalBatchesListResponse(ledgerJournalBatchesList model.LedgerJournalBatchesList) LedgerJournalBatchesListResponse {
	dtoLedgerJournalBatchesListResponse := LedgerJournalBatchesListResponse{}
	for _, ledgerJournalBatches := range ledgerJournalBatchesList {
		dtoLedgerJournalBatchesResponse := NewLedgerJournalBatchesResponse(*ledgerJournalBatches)
		dtoLedgerJournalBatchesListResponse = append(dtoLedgerJournalBatchesListResponse, &dtoLedgerJournalBatchesResponse)
	}
	return dtoLedgerJournalBatchesListResponse
}

type LedgerJournalBatchesPrimaryIDList []LedgerJournalBatchesPrimaryID

func (d LedgerJournalBatchesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, ledgerJournalBatches := range d {
		err = validator.Struct(ledgerJournalBatches)
		if err != nil {
			return
		}
	}
	return nil
}

type LedgerJournalBatchesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *LedgerJournalBatchesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d LedgerJournalBatchesPrimaryID) ToModel() model.LedgerJournalBatchesPrimaryID {
	return model.LedgerJournalBatchesPrimaryID{
		Id: d.Id,
	}
}
