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

type ProviderStatementFilesDTOFieldNameType string

type providerStatementFilesDTOFieldName struct {
	Id                ProviderStatementFilesDTOFieldNameType
	ProviderAccountId ProviderStatementFilesDTOFieldNameType
	StatementType     ProviderStatementFilesDTOFieldNameType
	StatementDate     ProviderStatementFilesDTOFieldNameType
	FileName          ProviderStatementFilesDTOFieldNameType
	StorageUrl        ProviderStatementFilesDTOFieldNameType
	FileHash          ProviderStatementFilesDTOFieldNameType
	ParseStatus       ProviderStatementFilesDTOFieldNameType
	Metadata          ProviderStatementFilesDTOFieldNameType
	MetaCreatedAt     ProviderStatementFilesDTOFieldNameType
	MetaCreatedBy     ProviderStatementFilesDTOFieldNameType
	MetaUpdatedAt     ProviderStatementFilesDTOFieldNameType
	MetaUpdatedBy     ProviderStatementFilesDTOFieldNameType
	MetaDeletedAt     ProviderStatementFilesDTOFieldNameType
	MetaDeletedBy     ProviderStatementFilesDTOFieldNameType
}

var ProviderStatementFilesDTOFieldName = providerStatementFilesDTOFieldName{
	Id:                "id",
	ProviderAccountId: "providerAccountId",
	StatementType:     "statementType",
	StatementDate:     "statementDate",
	FileName:          "fileName",
	StorageUrl:        "storageUrl",
	FileHash:          "fileHash",
	ParseStatus:       "parseStatus",
	Metadata:          "metadata",
	MetaCreatedAt:     "metaCreatedAt",
	MetaCreatedBy:     "metaCreatedBy",
	MetaUpdatedAt:     "metaUpdatedAt",
	MetaUpdatedBy:     "metaUpdatedBy",
	MetaDeletedAt:     "metaDeletedAt",
	MetaDeletedBy:     "metaDeletedBy",
}

func transformProviderStatementFilesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderStatementFilesDTOFieldName.Id):
		return string(model.ProviderStatementFilesDBFieldName.Id), true

	case string(ProviderStatementFilesDTOFieldName.ProviderAccountId):
		return string(model.ProviderStatementFilesDBFieldName.ProviderAccountId), true

	case string(ProviderStatementFilesDTOFieldName.StatementType):
		return string(model.ProviderStatementFilesDBFieldName.StatementType), true

	case string(ProviderStatementFilesDTOFieldName.StatementDate):
		return string(model.ProviderStatementFilesDBFieldName.StatementDate), true

	case string(ProviderStatementFilesDTOFieldName.FileName):
		return string(model.ProviderStatementFilesDBFieldName.FileName), true

	case string(ProviderStatementFilesDTOFieldName.StorageUrl):
		return string(model.ProviderStatementFilesDBFieldName.StorageUrl), true

	case string(ProviderStatementFilesDTOFieldName.FileHash):
		return string(model.ProviderStatementFilesDBFieldName.FileHash), true

	case string(ProviderStatementFilesDTOFieldName.ParseStatus):
		return string(model.ProviderStatementFilesDBFieldName.ParseStatus), true

	case string(ProviderStatementFilesDTOFieldName.Metadata):
		return string(model.ProviderStatementFilesDBFieldName.Metadata), true

	case string(ProviderStatementFilesDTOFieldName.MetaCreatedAt):
		return string(model.ProviderStatementFilesDBFieldName.MetaCreatedAt), true

	case string(ProviderStatementFilesDTOFieldName.MetaCreatedBy):
		return string(model.ProviderStatementFilesDBFieldName.MetaCreatedBy), true

	case string(ProviderStatementFilesDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderStatementFilesDBFieldName.MetaUpdatedAt), true

	case string(ProviderStatementFilesDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderStatementFilesDBFieldName.MetaUpdatedBy), true

	case string(ProviderStatementFilesDTOFieldName.MetaDeletedAt):
		return string(model.ProviderStatementFilesDBFieldName.MetaDeletedAt), true

	case string(ProviderStatementFilesDTOFieldName.MetaDeletedBy):
		return string(model.ProviderStatementFilesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderStatementFilesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderStatementFilesBaseFilterField(field string) bool {
	spec, found := model.NewProviderStatementFilesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderStatementFilesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderStatementFilesProjectionOutputPath(path string) error {
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

func transformProviderStatementFilesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderStatementFilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderStatementFilesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderStatementFilesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderStatementFilesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderStatementFilesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderStatementFilesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderStatementFilesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderStatementFilesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderStatementFilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderStatementFilesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderStatementFilesFilter(filter *model.Filter) {
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
			Field: string(ProviderStatementFilesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderStatementFilesSelectableResponse map[string]interface{}
type ProviderStatementFilesSelectableListResponse []*ProviderStatementFilesSelectableResponse

func assignProviderStatementFilesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setProviderStatementFilesSelectableValue(out ProviderStatementFilesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderStatementFilesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderStatementFilesSelectableResponse(providerStatementFiles model.ProviderStatementFiles, filter model.Filter) ProviderStatementFilesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderStatementFilesDBFieldName.Id),
			string(model.ProviderStatementFilesDBFieldName.ProviderAccountId),
			string(model.ProviderStatementFilesDBFieldName.StatementType),
			string(model.ProviderStatementFilesDBFieldName.StatementDate),
			string(model.ProviderStatementFilesDBFieldName.FileName),
			string(model.ProviderStatementFilesDBFieldName.StorageUrl),
			string(model.ProviderStatementFilesDBFieldName.FileHash),
			string(model.ProviderStatementFilesDBFieldName.ParseStatus),
			string(model.ProviderStatementFilesDBFieldName.Metadata),
			string(model.ProviderStatementFilesDBFieldName.MetaCreatedAt),
			string(model.ProviderStatementFilesDBFieldName.MetaCreatedBy),
			string(model.ProviderStatementFilesDBFieldName.MetaUpdatedAt),
			string(model.ProviderStatementFilesDBFieldName.MetaUpdatedBy),
			string(model.ProviderStatementFilesDBFieldName.MetaDeletedAt),
			string(model.ProviderStatementFilesDBFieldName.MetaDeletedBy),
		)
	}
	providerStatementFilesSelectableResponse := ProviderStatementFilesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderStatementFilesDBFieldName.Id):
			key := string(ProviderStatementFilesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.Id, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.ProviderAccountId):
			key := string(ProviderStatementFilesDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.ProviderAccountId, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.StatementType):
			key := string(ProviderStatementFilesDTOFieldName.StatementType)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, model.StatementType(providerStatementFiles.StatementType), explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.StatementDate):
			key := string(ProviderStatementFilesDTOFieldName.StatementDate)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.StatementDate, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.FileName):
			key := string(ProviderStatementFilesDTOFieldName.FileName)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.FileName, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.StorageUrl):
			key := string(ProviderStatementFilesDTOFieldName.StorageUrl)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.StorageUrl.String, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.FileHash):
			key := string(ProviderStatementFilesDTOFieldName.FileHash)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.FileHash, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.ParseStatus):
			key := string(ProviderStatementFilesDTOFieldName.ParseStatus)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, model.ParseStatus(providerStatementFiles.ParseStatus), explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.Metadata):
			key := string(ProviderStatementFilesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.Metadata, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.MetaCreatedAt):
			key := string(ProviderStatementFilesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.MetaCreatedAt, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.MetaCreatedBy):
			key := string(ProviderStatementFilesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.MetaCreatedBy, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.MetaUpdatedAt):
			key := string(ProviderStatementFilesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.MetaUpdatedBy):
			key := string(ProviderStatementFilesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.MetaUpdatedBy, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.MetaDeletedAt):
			key := string(ProviderStatementFilesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderStatementFilesDBFieldName.MetaDeletedBy):
			key := string(ProviderStatementFilesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderStatementFilesSelectableValue(providerStatementFilesSelectableResponse, key, providerStatementFiles.MetaDeletedBy, explicitAlias)

		}
	}
	return providerStatementFilesSelectableResponse
}

func NewProviderStatementFilesListResponseFromFilterResult(result []model.ProviderStatementFilesFilterResult, filter model.Filter) ProviderStatementFilesSelectableListResponse {
	dtoProviderStatementFilesListResponse := ProviderStatementFilesSelectableListResponse{}
	for _, row := range result {
		dtoProviderStatementFilesResponse := NewProviderStatementFilesSelectableResponse(row.ProviderStatementFiles, filter)
		dtoProviderStatementFilesListResponse = append(dtoProviderStatementFilesListResponse, &dtoProviderStatementFilesResponse)
	}
	return dtoProviderStatementFilesListResponse
}

type ProviderStatementFilesFilterResponse struct {
	Metadata Metadata                                     `json:"metadata"`
	Data     ProviderStatementFilesSelectableListResponse `json:"data"`
}

func reverseProviderStatementFilesFilterResults(result []model.ProviderStatementFilesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderStatementFilesFilterResponse(result []model.ProviderStatementFilesFilterResult, filter model.Filter) (resp ProviderStatementFilesFilterResponse) {
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
			reverseProviderStatementFilesFilterResults(dataResult)
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

	resp.Data = NewProviderStatementFilesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderStatementFilesCreateRequest struct {
	ProviderAccountId uuid.UUID           `json:"providerAccountId"`
	StatementType     model.StatementType `json:"statementType" example:"collection" enums:"collection,refund,payout,balance,chargeback"`
	StatementDate     time.Time           `json:"statementDate"`
	FileName          string              `json:"fileName"`
	StorageUrl        string              `json:"storageUrl"`
	FileHash          string              `json:"fileHash"`
	ParseStatus       model.ParseStatus   `json:"parseStatus" example:"uploaded" enums:"uploaded,parsed,failed,reprocessed"`
	Metadata          json.RawMessage     `json:"metadata"`
}

func (d *ProviderStatementFilesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderStatementFilesCreateRequest) ToModel() model.ProviderStatementFiles {
	id, _ := uuid.NewV4()
	return model.ProviderStatementFiles{
		Id:                id,
		ProviderAccountId: d.ProviderAccountId,
		StatementType:     d.StatementType,
		StatementDate:     d.StatementDate,
		FileName:          d.FileName,
		StorageUrl:        null.StringFrom(d.StorageUrl),
		FileHash:          d.FileHash,
		ParseStatus:       d.ParseStatus,
		Metadata:          d.Metadata,
	}
}

type ProviderStatementFilesListCreateRequest []*ProviderStatementFilesCreateRequest

func (d ProviderStatementFilesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerStatementFiles := range d {
		err = validator.Struct(providerStatementFiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderStatementFilesListCreateRequest) ToModelList() []model.ProviderStatementFiles {
	out := make([]model.ProviderStatementFiles, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderStatementFilesUpdateRequest struct {
	ProviderAccountId uuid.UUID           `json:"providerAccountId"`
	StatementType     model.StatementType `json:"statementType" example:"collection" enums:"collection,refund,payout,balance,chargeback"`
	StatementDate     time.Time           `json:"statementDate"`
	FileName          string              `json:"fileName"`
	StorageUrl        string              `json:"storageUrl"`
	FileHash          string              `json:"fileHash"`
	ParseStatus       model.ParseStatus   `json:"parseStatus" example:"uploaded" enums:"uploaded,parsed,failed,reprocessed"`
	Metadata          json.RawMessage     `json:"metadata"`
}

func (d *ProviderStatementFilesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderStatementFilesUpdateRequest) ToModel() model.ProviderStatementFiles {
	return model.ProviderStatementFiles{
		ProviderAccountId: d.ProviderAccountId,
		StatementType:     d.StatementType,
		StatementDate:     d.StatementDate,
		FileName:          d.FileName,
		StorageUrl:        null.StringFrom(d.StorageUrl),
		FileHash:          d.FileHash,
		ParseStatus:       d.ParseStatus,
		Metadata:          d.Metadata,
	}
}

type ProviderStatementFilesBulkUpdateRequest struct {
	Id                uuid.UUID           `json:"id"`
	ProviderAccountId uuid.UUID           `json:"providerAccountId"`
	StatementType     model.StatementType `json:"statementType" example:"collection" enums:"collection,refund,payout,balance,chargeback"`
	StatementDate     time.Time           `json:"statementDate"`
	FileName          string              `json:"fileName"`
	StorageUrl        string              `json:"storageUrl"`
	FileHash          string              `json:"fileHash"`
	ParseStatus       model.ParseStatus   `json:"parseStatus" example:"uploaded" enums:"uploaded,parsed,failed,reprocessed"`
	Metadata          json.RawMessage     `json:"metadata"`
}

func (d ProviderStatementFilesBulkUpdateRequest) PrimaryID() ProviderStatementFilesPrimaryID {
	return ProviderStatementFilesPrimaryID{
		Id: d.Id,
	}
}

type ProviderStatementFilesListBulkUpdateRequest []*ProviderStatementFilesBulkUpdateRequest

func (d ProviderStatementFilesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerStatementFiles := range d {
		err = validator.Struct(providerStatementFiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderStatementFilesBulkUpdateRequest) ToModel() model.ProviderStatementFiles {
	return model.ProviderStatementFiles{
		Id:                d.Id,
		ProviderAccountId: d.ProviderAccountId,
		StatementType:     d.StatementType,
		StatementDate:     d.StatementDate,
		FileName:          d.FileName,
		StorageUrl:        null.StringFrom(d.StorageUrl),
		FileHash:          d.FileHash,
		ParseStatus:       d.ParseStatus,
		Metadata:          d.Metadata,
	}
}

type ProviderStatementFilesResponse struct {
	Id                uuid.UUID           `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderAccountId uuid.UUID           `json:"providerAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	StatementType     model.StatementType `json:"statementType" validate:"required,oneof=collection refund payout balance chargeback" enums:"collection,refund,payout,balance,chargeback"`
	StatementDate     time.Time           `json:"statementDate" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	FileName          string              `json:"fileName" validate:"required"`
	StorageUrl        string              `json:"storageUrl" validate:"url"`
	FileHash          string              `json:"fileHash" validate:"required"`
	ParseStatus       model.ParseStatus   `json:"parseStatus" validate:"oneof=uploaded parsed failed reprocessed" enums:"uploaded,parsed,failed,reprocessed"`
	Metadata          json.RawMessage     `json:"metadata" swaggertype:"object"`
}

func NewProviderStatementFilesResponse(providerStatementFiles model.ProviderStatementFiles) ProviderStatementFilesResponse {
	return ProviderStatementFilesResponse{
		Id:                providerStatementFiles.Id,
		ProviderAccountId: providerStatementFiles.ProviderAccountId,
		StatementType:     model.StatementType(providerStatementFiles.StatementType),
		StatementDate:     providerStatementFiles.StatementDate,
		FileName:          providerStatementFiles.FileName,
		StorageUrl:        providerStatementFiles.StorageUrl.String,
		FileHash:          providerStatementFiles.FileHash,
		ParseStatus:       model.ParseStatus(providerStatementFiles.ParseStatus),
		Metadata:          providerStatementFiles.Metadata,
	}
}

type ProviderStatementFilesListResponse []*ProviderStatementFilesResponse

func NewProviderStatementFilesListResponse(providerStatementFilesList model.ProviderStatementFilesList) ProviderStatementFilesListResponse {
	dtoProviderStatementFilesListResponse := ProviderStatementFilesListResponse{}
	for _, providerStatementFiles := range providerStatementFilesList {
		dtoProviderStatementFilesResponse := NewProviderStatementFilesResponse(*providerStatementFiles)
		dtoProviderStatementFilesListResponse = append(dtoProviderStatementFilesListResponse, &dtoProviderStatementFilesResponse)
	}
	return dtoProviderStatementFilesListResponse
}

type ProviderStatementFilesPrimaryIDList []ProviderStatementFilesPrimaryID

func (d ProviderStatementFilesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerStatementFiles := range d {
		err = validator.Struct(providerStatementFiles)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderStatementFilesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderStatementFilesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderStatementFilesPrimaryID) ToModel() model.ProviderStatementFilesPrimaryID {
	return model.ProviderStatementFilesPrimaryID{
		Id: d.Id,
	}
}
