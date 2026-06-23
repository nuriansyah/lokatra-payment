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

type BankStatementFilesDTOFieldNameType string

type bankStatementFilesDTOFieldName struct {
	Id                   BankStatementFilesDTOFieldNameType
	BankCode             BankStatementFilesDTOFieldNameType
	AccountNumberMasked  BankStatementFilesDTOFieldNameType
	StatementPeriodStart BankStatementFilesDTOFieldNameType
	StatementPeriodEnd   BankStatementFilesDTOFieldNameType
	CurrencyCode         BankStatementFilesDTOFieldNameType
	StorageUri           BankStatementFilesDTOFieldNameType
	FileHash             BankStatementFilesDTOFieldNameType
	ImportStatus         BankStatementFilesDTOFieldNameType
	ImportedAt           BankStatementFilesDTOFieldNameType
	RowCount             BankStatementFilesDTOFieldNameType
	Metadata             BankStatementFilesDTOFieldNameType
	MetaCreatedAt        BankStatementFilesDTOFieldNameType
	MetaCreatedBy        BankStatementFilesDTOFieldNameType
	MetaUpdatedAt        BankStatementFilesDTOFieldNameType
	MetaUpdatedBy        BankStatementFilesDTOFieldNameType
	MetaDeletedAt        BankStatementFilesDTOFieldNameType
	MetaDeletedBy        BankStatementFilesDTOFieldNameType
}

var BankStatementFilesDTOFieldName = bankStatementFilesDTOFieldName{
	Id:                   "id",
	BankCode:             "bankCode",
	AccountNumberMasked:  "accountNumberMasked",
	StatementPeriodStart: "statementPeriodStart",
	StatementPeriodEnd:   "statementPeriodEnd",
	CurrencyCode:         "currencyCode",
	StorageUri:           "storageUri",
	FileHash:             "fileHash",
	ImportStatus:         "importStatus",
	ImportedAt:           "importedAt",
	RowCount:             "rowCount",
	Metadata:             "metadata",
	MetaCreatedAt:        "metaCreatedAt",
	MetaCreatedBy:        "metaCreatedBy",
	MetaUpdatedAt:        "metaUpdatedAt",
	MetaUpdatedBy:        "metaUpdatedBy",
	MetaDeletedAt:        "metaDeletedAt",
	MetaDeletedBy:        "metaDeletedBy",
}

func transformBankStatementFilesDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(BankStatementFilesDTOFieldName.Id):
		return string(model.BankStatementFilesDBFieldName.Id), true

	case string(BankStatementFilesDTOFieldName.BankCode):
		return string(model.BankStatementFilesDBFieldName.BankCode), true

	case string(BankStatementFilesDTOFieldName.AccountNumberMasked):
		return string(model.BankStatementFilesDBFieldName.AccountNumberMasked), true

	case string(BankStatementFilesDTOFieldName.StatementPeriodStart):
		return string(model.BankStatementFilesDBFieldName.StatementPeriodStart), true

	case string(BankStatementFilesDTOFieldName.StatementPeriodEnd):
		return string(model.BankStatementFilesDBFieldName.StatementPeriodEnd), true

	case string(BankStatementFilesDTOFieldName.CurrencyCode):
		return string(model.BankStatementFilesDBFieldName.CurrencyCode), true

	case string(BankStatementFilesDTOFieldName.StorageUri):
		return string(model.BankStatementFilesDBFieldName.StorageUri), true

	case string(BankStatementFilesDTOFieldName.FileHash):
		return string(model.BankStatementFilesDBFieldName.FileHash), true

	case string(BankStatementFilesDTOFieldName.ImportStatus):
		return string(model.BankStatementFilesDBFieldName.ImportStatus), true

	case string(BankStatementFilesDTOFieldName.ImportedAt):
		return string(model.BankStatementFilesDBFieldName.ImportedAt), true

	case string(BankStatementFilesDTOFieldName.RowCount):
		return string(model.BankStatementFilesDBFieldName.RowCount), true

	case string(BankStatementFilesDTOFieldName.Metadata):
		return string(model.BankStatementFilesDBFieldName.Metadata), true

	case string(BankStatementFilesDTOFieldName.MetaCreatedAt):
		return string(model.BankStatementFilesDBFieldName.MetaCreatedAt), true

	case string(BankStatementFilesDTOFieldName.MetaCreatedBy):
		return string(model.BankStatementFilesDBFieldName.MetaCreatedBy), true

	case string(BankStatementFilesDTOFieldName.MetaUpdatedAt):
		return string(model.BankStatementFilesDBFieldName.MetaUpdatedAt), true

	case string(BankStatementFilesDTOFieldName.MetaUpdatedBy):
		return string(model.BankStatementFilesDBFieldName.MetaUpdatedBy), true

	case string(BankStatementFilesDTOFieldName.MetaDeletedAt):
		return string(model.BankStatementFilesDBFieldName.MetaDeletedAt), true

	case string(BankStatementFilesDTOFieldName.MetaDeletedBy):
		return string(model.BankStatementFilesDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewBankStatementFilesFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isBankStatementFilesBaseFilterField(field string) bool {
	spec, found := model.NewBankStatementFilesFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeBankStatementFilesProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateBankStatementFilesProjectionOutputPath(path string) error {
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

func transformBankStatementFilesFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformBankStatementFilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformBankStatementFilesFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformBankStatementFilesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformBankStatementFilesDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isBankStatementFilesBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateBankStatementFilesProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeBankStatementFilesProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformBankStatementFilesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformBankStatementFilesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformBankStatementFilesFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultBankStatementFilesFilter(filter *model.Filter) {
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
			Field: string(BankStatementFilesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type BankStatementFilesSelectableResponse map[string]interface{}
type BankStatementFilesSelectableListResponse []*BankStatementFilesSelectableResponse

func assignBankStatementFilesNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setBankStatementFilesSelectableValue(out BankStatementFilesSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignBankStatementFilesNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewBankStatementFilesSelectableResponse(bankStatementFiles model.BankStatementFiles, filter model.Filter) BankStatementFilesSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.BankStatementFilesDBFieldName.Id),
			string(model.BankStatementFilesDBFieldName.BankCode),
			string(model.BankStatementFilesDBFieldName.AccountNumberMasked),
			string(model.BankStatementFilesDBFieldName.StatementPeriodStart),
			string(model.BankStatementFilesDBFieldName.StatementPeriodEnd),
			string(model.BankStatementFilesDBFieldName.CurrencyCode),
			string(model.BankStatementFilesDBFieldName.StorageUri),
			string(model.BankStatementFilesDBFieldName.FileHash),
			string(model.BankStatementFilesDBFieldName.ImportStatus),
			string(model.BankStatementFilesDBFieldName.ImportedAt),
			string(model.BankStatementFilesDBFieldName.RowCount),
			string(model.BankStatementFilesDBFieldName.Metadata),
			string(model.BankStatementFilesDBFieldName.MetaCreatedAt),
			string(model.BankStatementFilesDBFieldName.MetaCreatedBy),
			string(model.BankStatementFilesDBFieldName.MetaUpdatedAt),
			string(model.BankStatementFilesDBFieldName.MetaUpdatedBy),
			string(model.BankStatementFilesDBFieldName.MetaDeletedAt),
			string(model.BankStatementFilesDBFieldName.MetaDeletedBy),
		)
	}
	bankStatementFilesSelectableResponse := BankStatementFilesSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.BankStatementFilesDBFieldName.Id):
			key := string(BankStatementFilesDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.Id, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.BankCode):
			key := string(BankStatementFilesDTOFieldName.BankCode)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.BankCode, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.AccountNumberMasked):
			key := string(BankStatementFilesDTOFieldName.AccountNumberMasked)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.AccountNumberMasked.String, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.StatementPeriodStart):
			key := string(BankStatementFilesDTOFieldName.StatementPeriodStart)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.StatementPeriodStart, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.StatementPeriodEnd):
			key := string(BankStatementFilesDTOFieldName.StatementPeriodEnd)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.StatementPeriodEnd, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.CurrencyCode):
			key := string(BankStatementFilesDTOFieldName.CurrencyCode)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.CurrencyCode, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.StorageUri):
			key := string(BankStatementFilesDTOFieldName.StorageUri)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.StorageUri.String, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.FileHash):
			key := string(BankStatementFilesDTOFieldName.FileHash)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.FileHash, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.ImportStatus):
			key := string(BankStatementFilesDTOFieldName.ImportStatus)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, model.ImportStatus(bankStatementFiles.ImportStatus), explicitAlias)

		case string(model.BankStatementFilesDBFieldName.ImportedAt):
			key := string(BankStatementFilesDTOFieldName.ImportedAt)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.ImportedAt.Time, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.RowCount):
			key := string(BankStatementFilesDTOFieldName.RowCount)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.RowCount, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.Metadata):
			key := string(BankStatementFilesDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.Metadata, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.MetaCreatedAt):
			key := string(BankStatementFilesDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.MetaCreatedAt, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.MetaCreatedBy):
			key := string(BankStatementFilesDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.MetaCreatedBy, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.MetaUpdatedAt):
			key := string(BankStatementFilesDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.MetaUpdatedAt, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.MetaUpdatedBy):
			key := string(BankStatementFilesDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.MetaUpdatedBy, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.MetaDeletedAt):
			key := string(BankStatementFilesDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.MetaDeletedAt.Time, explicitAlias)

		case string(model.BankStatementFilesDBFieldName.MetaDeletedBy):
			key := string(BankStatementFilesDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setBankStatementFilesSelectableValue(bankStatementFilesSelectableResponse, key, bankStatementFiles.MetaDeletedBy, explicitAlias)

		}
	}
	return bankStatementFilesSelectableResponse
}

func NewBankStatementFilesListResponseFromFilterResult(result []model.BankStatementFilesFilterResult, filter model.Filter) BankStatementFilesSelectableListResponse {
	dtoBankStatementFilesListResponse := BankStatementFilesSelectableListResponse{}
	for _, row := range result {
		dtoBankStatementFilesResponse := NewBankStatementFilesSelectableResponse(row.BankStatementFiles, filter)
		dtoBankStatementFilesListResponse = append(dtoBankStatementFilesListResponse, &dtoBankStatementFilesResponse)
	}
	return dtoBankStatementFilesListResponse
}

type BankStatementFilesFilterResponse struct {
	Metadata Metadata                                 `json:"metadata"`
	Data     BankStatementFilesSelectableListResponse `json:"data"`
}

func reverseBankStatementFilesFilterResults(result []model.BankStatementFilesFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewBankStatementFilesFilterResponse(result []model.BankStatementFilesFilterResult, filter model.Filter) (resp BankStatementFilesFilterResponse) {
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
			reverseBankStatementFilesFilterResults(dataResult)
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

	resp.Data = NewBankStatementFilesListResponseFromFilterResult(dataResult, filter)
	return resp
}

type BankStatementFilesCreateRequest struct {
	BankCode             string             `json:"bankCode"`
	AccountNumberMasked  string             `json:"accountNumberMasked"`
	StatementPeriodStart time.Time          `json:"statementPeriodStart"`
	StatementPeriodEnd   time.Time          `json:"statementPeriodEnd"`
	CurrencyCode         string             `json:"currencyCode"`
	StorageUri           string             `json:"storageUri"`
	FileHash             string             `json:"fileHash"`
	ImportStatus         model.ImportStatus `json:"importStatus" example:"uploaded" enums:"uploaded,processing,completed,failed,cancelled"`
	ImportedAt           time.Time          `json:"importedAt"`
	RowCount             int                `json:"rowCount"`
	Metadata             json.RawMessage    `json:"metadata"`
}

func (d *BankStatementFilesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *BankStatementFilesCreateRequest) ToModel() model.BankStatementFiles {
	id, _ := uuid.NewV4()
	return model.BankStatementFiles{
		Id:                   id,
		BankCode:             d.BankCode,
		AccountNumberMasked:  null.StringFrom(d.AccountNumberMasked),
		StatementPeriodStart: d.StatementPeriodStart,
		StatementPeriodEnd:   d.StatementPeriodEnd,
		CurrencyCode:         d.CurrencyCode,
		StorageUri:           null.StringFrom(d.StorageUri),
		FileHash:             d.FileHash,
		ImportStatus:         d.ImportStatus,
		ImportedAt:           null.TimeFrom(d.ImportedAt),
		RowCount:             d.RowCount,
		Metadata:             d.Metadata,
	}
}

type BankStatementFilesListCreateRequest []*BankStatementFilesCreateRequest

func (d BankStatementFilesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, bankStatementFiles := range d {
		err = validator.Struct(bankStatementFiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d BankStatementFilesListCreateRequest) ToModelList() []model.BankStatementFiles {
	out := make([]model.BankStatementFiles, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type BankStatementFilesUpdateRequest struct {
	BankCode             string             `json:"bankCode"`
	AccountNumberMasked  string             `json:"accountNumberMasked"`
	StatementPeriodStart time.Time          `json:"statementPeriodStart"`
	StatementPeriodEnd   time.Time          `json:"statementPeriodEnd"`
	CurrencyCode         string             `json:"currencyCode"`
	StorageUri           string             `json:"storageUri"`
	FileHash             string             `json:"fileHash"`
	ImportStatus         model.ImportStatus `json:"importStatus" example:"uploaded" enums:"uploaded,processing,completed,failed,cancelled"`
	ImportedAt           time.Time          `json:"importedAt"`
	RowCount             int                `json:"rowCount"`
	Metadata             json.RawMessage    `json:"metadata"`
}

func (d *BankStatementFilesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d BankStatementFilesUpdateRequest) ToModel() model.BankStatementFiles {
	return model.BankStatementFiles{
		BankCode:             d.BankCode,
		AccountNumberMasked:  null.StringFrom(d.AccountNumberMasked),
		StatementPeriodStart: d.StatementPeriodStart,
		StatementPeriodEnd:   d.StatementPeriodEnd,
		CurrencyCode:         d.CurrencyCode,
		StorageUri:           null.StringFrom(d.StorageUri),
		FileHash:             d.FileHash,
		ImportStatus:         d.ImportStatus,
		ImportedAt:           null.TimeFrom(d.ImportedAt),
		RowCount:             d.RowCount,
		Metadata:             d.Metadata,
	}
}

type BankStatementFilesBulkUpdateRequest struct {
	Id                   uuid.UUID          `json:"id"`
	BankCode             string             `json:"bankCode"`
	AccountNumberMasked  string             `json:"accountNumberMasked"`
	StatementPeriodStart time.Time          `json:"statementPeriodStart"`
	StatementPeriodEnd   time.Time          `json:"statementPeriodEnd"`
	CurrencyCode         string             `json:"currencyCode"`
	StorageUri           string             `json:"storageUri"`
	FileHash             string             `json:"fileHash"`
	ImportStatus         model.ImportStatus `json:"importStatus" example:"uploaded" enums:"uploaded,processing,completed,failed,cancelled"`
	ImportedAt           time.Time          `json:"importedAt"`
	RowCount             int                `json:"rowCount"`
	Metadata             json.RawMessage    `json:"metadata"`
}

func (d BankStatementFilesBulkUpdateRequest) PrimaryID() BankStatementFilesPrimaryID {
	return BankStatementFilesPrimaryID{
		Id: d.Id,
	}
}

type BankStatementFilesListBulkUpdateRequest []*BankStatementFilesBulkUpdateRequest

func (d BankStatementFilesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, bankStatementFiles := range d {
		err = validator.Struct(bankStatementFiles)
		if err != nil {
			return
		}
	}
	return nil
}

func (d BankStatementFilesBulkUpdateRequest) ToModel() model.BankStatementFiles {
	return model.BankStatementFiles{
		Id:                   d.Id,
		BankCode:             d.BankCode,
		AccountNumberMasked:  null.StringFrom(d.AccountNumberMasked),
		StatementPeriodStart: d.StatementPeriodStart,
		StatementPeriodEnd:   d.StatementPeriodEnd,
		CurrencyCode:         d.CurrencyCode,
		StorageUri:           null.StringFrom(d.StorageUri),
		FileHash:             d.FileHash,
		ImportStatus:         d.ImportStatus,
		ImportedAt:           null.TimeFrom(d.ImportedAt),
		RowCount:             d.RowCount,
		Metadata:             d.Metadata,
	}
}

type BankStatementFilesResponse struct {
	Id                   uuid.UUID          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BankCode             string             `json:"bankCode" validate:"required"`
	AccountNumberMasked  string             `json:"accountNumberMasked"`
	StatementPeriodStart time.Time          `json:"statementPeriodStart" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	StatementPeriodEnd   time.Time          `json:"statementPeriodEnd" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CurrencyCode         string             `json:"currencyCode" validate:"required"`
	StorageUri           string             `json:"storageUri"`
	FileHash             string             `json:"fileHash" validate:"required"`
	ImportStatus         model.ImportStatus `json:"importStatus" validate:"oneof=uploaded processing completed failed cancelled" enums:"uploaded,processing,completed,failed,cancelled"`
	ImportedAt           time.Time          `json:"importedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	RowCount             int                `json:"rowCount" example:"1"`
	Metadata             json.RawMessage    `json:"metadata" swaggertype:"object"`
}

func NewBankStatementFilesResponse(bankStatementFiles model.BankStatementFiles) BankStatementFilesResponse {
	return BankStatementFilesResponse{
		Id:                   bankStatementFiles.Id,
		BankCode:             bankStatementFiles.BankCode,
		AccountNumberMasked:  bankStatementFiles.AccountNumberMasked.String,
		StatementPeriodStart: bankStatementFiles.StatementPeriodStart,
		StatementPeriodEnd:   bankStatementFiles.StatementPeriodEnd,
		CurrencyCode:         bankStatementFiles.CurrencyCode,
		StorageUri:           bankStatementFiles.StorageUri.String,
		FileHash:             bankStatementFiles.FileHash,
		ImportStatus:         model.ImportStatus(bankStatementFiles.ImportStatus),
		ImportedAt:           bankStatementFiles.ImportedAt.Time,
		RowCount:             bankStatementFiles.RowCount,
		Metadata:             bankStatementFiles.Metadata,
	}
}

type BankStatementFilesListResponse []*BankStatementFilesResponse

func NewBankStatementFilesListResponse(bankStatementFilesList model.BankStatementFilesList) BankStatementFilesListResponse {
	dtoBankStatementFilesListResponse := BankStatementFilesListResponse{}
	for _, bankStatementFiles := range bankStatementFilesList {
		dtoBankStatementFilesResponse := NewBankStatementFilesResponse(*bankStatementFiles)
		dtoBankStatementFilesListResponse = append(dtoBankStatementFilesListResponse, &dtoBankStatementFilesResponse)
	}
	return dtoBankStatementFilesListResponse
}

type BankStatementFilesPrimaryIDList []BankStatementFilesPrimaryID

func (d BankStatementFilesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, bankStatementFiles := range d {
		err = validator.Struct(bankStatementFiles)
		if err != nil {
			return
		}
	}
	return nil
}

type BankStatementFilesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *BankStatementFilesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d BankStatementFilesPrimaryID) ToModel() model.BankStatementFilesPrimaryID {
	return model.BankStatementFilesPrimaryID{
		Id: d.Id,
	}
}
