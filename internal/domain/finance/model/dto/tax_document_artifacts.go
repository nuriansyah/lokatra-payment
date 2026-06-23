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

type TaxDocumentArtifactsDTOFieldNameType string

type taxDocumentArtifactsDTOFieldName struct {
	Id             TaxDocumentArtifactsDTOFieldNameType
	DocumentType   TaxDocumentArtifactsDTOFieldNameType
	DocumentId     TaxDocumentArtifactsDTOFieldNameType
	ArtifactType   TaxDocumentArtifactsDTOFieldNameType
	StorageUri     TaxDocumentArtifactsDTOFieldNameType
	ContentHash    TaxDocumentArtifactsDTOFieldNameType
	IdempotencyKey TaxDocumentArtifactsDTOFieldNameType
	GeneratedAt    TaxDocumentArtifactsDTOFieldNameType
	Metadata       TaxDocumentArtifactsDTOFieldNameType
	MetaCreatedAt  TaxDocumentArtifactsDTOFieldNameType
	MetaCreatedBy  TaxDocumentArtifactsDTOFieldNameType
	MetaUpdatedAt  TaxDocumentArtifactsDTOFieldNameType
	MetaUpdatedBy  TaxDocumentArtifactsDTOFieldNameType
	MetaDeletedAt  TaxDocumentArtifactsDTOFieldNameType
	MetaDeletedBy  TaxDocumentArtifactsDTOFieldNameType
}

var TaxDocumentArtifactsDTOFieldName = taxDocumentArtifactsDTOFieldName{
	Id:             "id",
	DocumentType:   "documentType",
	DocumentId:     "documentId",
	ArtifactType:   "artifactType",
	StorageUri:     "storageUri",
	ContentHash:    "contentHash",
	IdempotencyKey: "idempotencyKey",
	GeneratedAt:    "generatedAt",
	Metadata:       "metadata",
	MetaCreatedAt:  "metaCreatedAt",
	MetaCreatedBy:  "metaCreatedBy",
	MetaUpdatedAt:  "metaUpdatedAt",
	MetaUpdatedBy:  "metaUpdatedBy",
	MetaDeletedAt:  "metaDeletedAt",
	MetaDeletedBy:  "metaDeletedBy",
}

func transformTaxDocumentArtifactsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(TaxDocumentArtifactsDTOFieldName.Id):
		return string(model.TaxDocumentArtifactsDBFieldName.Id), true

	case string(TaxDocumentArtifactsDTOFieldName.DocumentType):
		return string(model.TaxDocumentArtifactsDBFieldName.DocumentType), true

	case string(TaxDocumentArtifactsDTOFieldName.DocumentId):
		return string(model.TaxDocumentArtifactsDBFieldName.DocumentId), true

	case string(TaxDocumentArtifactsDTOFieldName.ArtifactType):
		return string(model.TaxDocumentArtifactsDBFieldName.ArtifactType), true

	case string(TaxDocumentArtifactsDTOFieldName.StorageUri):
		return string(model.TaxDocumentArtifactsDBFieldName.StorageUri), true

	case string(TaxDocumentArtifactsDTOFieldName.ContentHash):
		return string(model.TaxDocumentArtifactsDBFieldName.ContentHash), true

	case string(TaxDocumentArtifactsDTOFieldName.IdempotencyKey):
		return string(model.TaxDocumentArtifactsDBFieldName.IdempotencyKey), true

	case string(TaxDocumentArtifactsDTOFieldName.GeneratedAt):
		return string(model.TaxDocumentArtifactsDBFieldName.GeneratedAt), true

	case string(TaxDocumentArtifactsDTOFieldName.Metadata):
		return string(model.TaxDocumentArtifactsDBFieldName.Metadata), true

	case string(TaxDocumentArtifactsDTOFieldName.MetaCreatedAt):
		return string(model.TaxDocumentArtifactsDBFieldName.MetaCreatedAt), true

	case string(TaxDocumentArtifactsDTOFieldName.MetaCreatedBy):
		return string(model.TaxDocumentArtifactsDBFieldName.MetaCreatedBy), true

	case string(TaxDocumentArtifactsDTOFieldName.MetaUpdatedAt):
		return string(model.TaxDocumentArtifactsDBFieldName.MetaUpdatedAt), true

	case string(TaxDocumentArtifactsDTOFieldName.MetaUpdatedBy):
		return string(model.TaxDocumentArtifactsDBFieldName.MetaUpdatedBy), true

	case string(TaxDocumentArtifactsDTOFieldName.MetaDeletedAt):
		return string(model.TaxDocumentArtifactsDBFieldName.MetaDeletedAt), true

	case string(TaxDocumentArtifactsDTOFieldName.MetaDeletedBy):
		return string(model.TaxDocumentArtifactsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewTaxDocumentArtifactsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isTaxDocumentArtifactsBaseFilterField(field string) bool {
	spec, found := model.NewTaxDocumentArtifactsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeTaxDocumentArtifactsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateTaxDocumentArtifactsProjectionOutputPath(path string) error {
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

func transformTaxDocumentArtifactsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformTaxDocumentArtifactsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformTaxDocumentArtifactsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformTaxDocumentArtifactsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformTaxDocumentArtifactsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isTaxDocumentArtifactsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateTaxDocumentArtifactsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeTaxDocumentArtifactsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformTaxDocumentArtifactsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformTaxDocumentArtifactsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformTaxDocumentArtifactsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultTaxDocumentArtifactsFilter(filter *model.Filter) {
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
			Field: string(TaxDocumentArtifactsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type TaxDocumentArtifactsSelectableResponse map[string]interface{}
type TaxDocumentArtifactsSelectableListResponse []*TaxDocumentArtifactsSelectableResponse

func assignTaxDocumentArtifactsNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setTaxDocumentArtifactsSelectableValue(out TaxDocumentArtifactsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignTaxDocumentArtifactsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewTaxDocumentArtifactsSelectableResponse(taxDocumentArtifacts model.TaxDocumentArtifacts, filter model.Filter) TaxDocumentArtifactsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.TaxDocumentArtifactsDBFieldName.Id),
			string(model.TaxDocumentArtifactsDBFieldName.DocumentType),
			string(model.TaxDocumentArtifactsDBFieldName.DocumentId),
			string(model.TaxDocumentArtifactsDBFieldName.ArtifactType),
			string(model.TaxDocumentArtifactsDBFieldName.StorageUri),
			string(model.TaxDocumentArtifactsDBFieldName.ContentHash),
			string(model.TaxDocumentArtifactsDBFieldName.IdempotencyKey),
			string(model.TaxDocumentArtifactsDBFieldName.GeneratedAt),
			string(model.TaxDocumentArtifactsDBFieldName.Metadata),
			string(model.TaxDocumentArtifactsDBFieldName.MetaCreatedAt),
			string(model.TaxDocumentArtifactsDBFieldName.MetaCreatedBy),
			string(model.TaxDocumentArtifactsDBFieldName.MetaUpdatedAt),
			string(model.TaxDocumentArtifactsDBFieldName.MetaUpdatedBy),
			string(model.TaxDocumentArtifactsDBFieldName.MetaDeletedAt),
			string(model.TaxDocumentArtifactsDBFieldName.MetaDeletedBy),
		)
	}
	taxDocumentArtifactsSelectableResponse := TaxDocumentArtifactsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.TaxDocumentArtifactsDBFieldName.Id):
			key := string(TaxDocumentArtifactsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.Id, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.DocumentType):
			key := string(TaxDocumentArtifactsDTOFieldName.DocumentType)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, model.DocumentType(taxDocumentArtifacts.DocumentType), explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.DocumentId):
			key := string(TaxDocumentArtifactsDTOFieldName.DocumentId)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.DocumentId, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.ArtifactType):
			key := string(TaxDocumentArtifactsDTOFieldName.ArtifactType)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.ArtifactType, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.StorageUri):
			key := string(TaxDocumentArtifactsDTOFieldName.StorageUri)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.StorageUri, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.ContentHash):
			key := string(TaxDocumentArtifactsDTOFieldName.ContentHash)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.ContentHash.String, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.IdempotencyKey):
			key := string(TaxDocumentArtifactsDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.IdempotencyKey.String, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.GeneratedAt):
			key := string(TaxDocumentArtifactsDTOFieldName.GeneratedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.GeneratedAt, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.Metadata):
			key := string(TaxDocumentArtifactsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.Metadata, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.MetaCreatedAt):
			key := string(TaxDocumentArtifactsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.MetaCreatedAt, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.MetaCreatedBy):
			key := string(TaxDocumentArtifactsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.MetaCreatedBy, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.MetaUpdatedAt):
			key := string(TaxDocumentArtifactsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.MetaUpdatedAt, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.MetaUpdatedBy):
			key := string(TaxDocumentArtifactsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.MetaUpdatedBy, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.MetaDeletedAt):
			key := string(TaxDocumentArtifactsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.MetaDeletedAt.Time, explicitAlias)

		case string(model.TaxDocumentArtifactsDBFieldName.MetaDeletedBy):
			key := string(TaxDocumentArtifactsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setTaxDocumentArtifactsSelectableValue(taxDocumentArtifactsSelectableResponse, key, taxDocumentArtifacts.MetaDeletedBy, explicitAlias)

		}
	}
	return taxDocumentArtifactsSelectableResponse
}

func NewTaxDocumentArtifactsListResponseFromFilterResult(result []model.TaxDocumentArtifactsFilterResult, filter model.Filter) TaxDocumentArtifactsSelectableListResponse {
	dtoTaxDocumentArtifactsListResponse := TaxDocumentArtifactsSelectableListResponse{}
	for _, row := range result {
		dtoTaxDocumentArtifactsResponse := NewTaxDocumentArtifactsSelectableResponse(row.TaxDocumentArtifacts, filter)
		dtoTaxDocumentArtifactsListResponse = append(dtoTaxDocumentArtifactsListResponse, &dtoTaxDocumentArtifactsResponse)
	}
	return dtoTaxDocumentArtifactsListResponse
}

type TaxDocumentArtifactsFilterResponse struct {
	Metadata Metadata                                   `json:"metadata"`
	Data     TaxDocumentArtifactsSelectableListResponse `json:"data"`
}

func reverseTaxDocumentArtifactsFilterResults(result []model.TaxDocumentArtifactsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewTaxDocumentArtifactsFilterResponse(result []model.TaxDocumentArtifactsFilterResult, filter model.Filter) (resp TaxDocumentArtifactsFilterResponse) {
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
			reverseTaxDocumentArtifactsFilterResults(dataResult)
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

	resp.Data = NewTaxDocumentArtifactsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type TaxDocumentArtifactsCreateRequest struct {
	DocumentType   model.DocumentType `json:"documentType" example:"tax_invoice" enums:"tax_invoice,credit_note"`
	DocumentId     uuid.UUID          `json:"documentId"`
	ArtifactType   string             `json:"artifactType"`
	StorageUri     string             `json:"storageUri"`
	ContentHash    string             `json:"contentHash"`
	IdempotencyKey string             `json:"idempotencyKey"`
	GeneratedAt    time.Time          `json:"generatedAt"`
	Metadata       json.RawMessage    `json:"metadata"`
}

func (d *TaxDocumentArtifactsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *TaxDocumentArtifactsCreateRequest) ToModel() model.TaxDocumentArtifacts {
	id, _ := uuid.NewV4()
	return model.TaxDocumentArtifacts{
		Id:             id,
		DocumentType:   d.DocumentType,
		DocumentId:     d.DocumentId,
		ArtifactType:   d.ArtifactType,
		StorageUri:     d.StorageUri,
		ContentHash:    null.StringFrom(d.ContentHash),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		GeneratedAt:    d.GeneratedAt,
		Metadata:       d.Metadata,
	}
}

type TaxDocumentArtifactsListCreateRequest []*TaxDocumentArtifactsCreateRequest

func (d TaxDocumentArtifactsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxDocumentArtifacts := range d {
		err = validator.Struct(taxDocumentArtifacts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxDocumentArtifactsListCreateRequest) ToModelList() []model.TaxDocumentArtifacts {
	out := make([]model.TaxDocumentArtifacts, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type TaxDocumentArtifactsUpdateRequest struct {
	DocumentType   model.DocumentType `json:"documentType" example:"tax_invoice" enums:"tax_invoice,credit_note"`
	DocumentId     uuid.UUID          `json:"documentId"`
	ArtifactType   string             `json:"artifactType"`
	StorageUri     string             `json:"storageUri"`
	ContentHash    string             `json:"contentHash"`
	IdempotencyKey string             `json:"idempotencyKey"`
	GeneratedAt    time.Time          `json:"generatedAt"`
	Metadata       json.RawMessage    `json:"metadata"`
}

func (d *TaxDocumentArtifactsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d TaxDocumentArtifactsUpdateRequest) ToModel() model.TaxDocumentArtifacts {
	return model.TaxDocumentArtifacts{
		DocumentType:   d.DocumentType,
		DocumentId:     d.DocumentId,
		ArtifactType:   d.ArtifactType,
		StorageUri:     d.StorageUri,
		ContentHash:    null.StringFrom(d.ContentHash),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		GeneratedAt:    d.GeneratedAt,
		Metadata:       d.Metadata,
	}
}

type TaxDocumentArtifactsBulkUpdateRequest struct {
	Id             uuid.UUID          `json:"id"`
	DocumentType   model.DocumentType `json:"documentType" example:"tax_invoice" enums:"tax_invoice,credit_note"`
	DocumentId     uuid.UUID          `json:"documentId"`
	ArtifactType   string             `json:"artifactType"`
	StorageUri     string             `json:"storageUri"`
	ContentHash    string             `json:"contentHash"`
	IdempotencyKey string             `json:"idempotencyKey"`
	GeneratedAt    time.Time          `json:"generatedAt"`
	Metadata       json.RawMessage    `json:"metadata"`
}

func (d TaxDocumentArtifactsBulkUpdateRequest) PrimaryID() TaxDocumentArtifactsPrimaryID {
	return TaxDocumentArtifactsPrimaryID{
		Id: d.Id,
	}
}

type TaxDocumentArtifactsListBulkUpdateRequest []*TaxDocumentArtifactsBulkUpdateRequest

func (d TaxDocumentArtifactsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxDocumentArtifacts := range d {
		err = validator.Struct(taxDocumentArtifacts)
		if err != nil {
			return
		}
	}
	return nil
}

func (d TaxDocumentArtifactsBulkUpdateRequest) ToModel() model.TaxDocumentArtifacts {
	return model.TaxDocumentArtifacts{
		Id:             d.Id,
		DocumentType:   d.DocumentType,
		DocumentId:     d.DocumentId,
		ArtifactType:   d.ArtifactType,
		StorageUri:     d.StorageUri,
		ContentHash:    null.StringFrom(d.ContentHash),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		GeneratedAt:    d.GeneratedAt,
		Metadata:       d.Metadata,
	}
}

type TaxDocumentArtifactsResponse struct {
	Id             uuid.UUID          `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	DocumentType   model.DocumentType `json:"documentType" validate:"required,oneof=tax_invoice credit_note" enums:"tax_invoice,credit_note"`
	DocumentId     uuid.UUID          `json:"documentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ArtifactType   string             `json:"artifactType"`
	StorageUri     string             `json:"storageUri" validate:"required"`
	ContentHash    string             `json:"contentHash"`
	IdempotencyKey string             `json:"idempotencyKey"`
	GeneratedAt    time.Time          `json:"generatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata       json.RawMessage    `json:"metadata" swaggertype:"object"`
}

func NewTaxDocumentArtifactsResponse(taxDocumentArtifacts model.TaxDocumentArtifacts) TaxDocumentArtifactsResponse {
	return TaxDocumentArtifactsResponse{
		Id:             taxDocumentArtifacts.Id,
		DocumentType:   model.DocumentType(taxDocumentArtifacts.DocumentType),
		DocumentId:     taxDocumentArtifacts.DocumentId,
		ArtifactType:   taxDocumentArtifacts.ArtifactType,
		StorageUri:     taxDocumentArtifacts.StorageUri,
		ContentHash:    taxDocumentArtifacts.ContentHash.String,
		IdempotencyKey: taxDocumentArtifacts.IdempotencyKey.String,
		GeneratedAt:    taxDocumentArtifacts.GeneratedAt,
		Metadata:       taxDocumentArtifacts.Metadata,
	}
}

type TaxDocumentArtifactsListResponse []*TaxDocumentArtifactsResponse

func NewTaxDocumentArtifactsListResponse(taxDocumentArtifactsList model.TaxDocumentArtifactsList) TaxDocumentArtifactsListResponse {
	dtoTaxDocumentArtifactsListResponse := TaxDocumentArtifactsListResponse{}
	for _, taxDocumentArtifacts := range taxDocumentArtifactsList {
		dtoTaxDocumentArtifactsResponse := NewTaxDocumentArtifactsResponse(*taxDocumentArtifacts)
		dtoTaxDocumentArtifactsListResponse = append(dtoTaxDocumentArtifactsListResponse, &dtoTaxDocumentArtifactsResponse)
	}
	return dtoTaxDocumentArtifactsListResponse
}

type TaxDocumentArtifactsPrimaryIDList []TaxDocumentArtifactsPrimaryID

func (d TaxDocumentArtifactsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, taxDocumentArtifacts := range d {
		err = validator.Struct(taxDocumentArtifacts)
		if err != nil {
			return
		}
	}
	return nil
}

type TaxDocumentArtifactsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *TaxDocumentArtifactsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d TaxDocumentArtifactsPrimaryID) ToModel() model.TaxDocumentArtifactsPrimaryID {
	return model.TaxDocumentArtifactsPrimaryID{
		Id: d.Id,
	}
}
