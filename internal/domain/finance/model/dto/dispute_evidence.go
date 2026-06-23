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

type DisputeEvidenceDTOFieldNameType string

type disputeEvidenceDTOFieldName struct {
	Id             DisputeEvidenceDTOFieldNameType
	DisputeId      DisputeEvidenceDTOFieldNameType
	EvidenceType   DisputeEvidenceDTOFieldNameType
	StorageUri     DisputeEvidenceDTOFieldNameType
	ContentHash    DisputeEvidenceDTOFieldNameType
	IdempotencyKey DisputeEvidenceDTOFieldNameType
	SubmittedBy    DisputeEvidenceDTOFieldNameType
	SubmittedAt    DisputeEvidenceDTOFieldNameType
	Metadata       DisputeEvidenceDTOFieldNameType
	MetaCreatedAt  DisputeEvidenceDTOFieldNameType
	MetaCreatedBy  DisputeEvidenceDTOFieldNameType
	MetaUpdatedAt  DisputeEvidenceDTOFieldNameType
	MetaUpdatedBy  DisputeEvidenceDTOFieldNameType
	MetaDeletedAt  DisputeEvidenceDTOFieldNameType
	MetaDeletedBy  DisputeEvidenceDTOFieldNameType
}

var DisputeEvidenceDTOFieldName = disputeEvidenceDTOFieldName{
	Id:             "id",
	DisputeId:      "disputeId",
	EvidenceType:   "evidenceType",
	StorageUri:     "storageUri",
	ContentHash:    "contentHash",
	IdempotencyKey: "idempotencyKey",
	SubmittedBy:    "submittedBy",
	SubmittedAt:    "submittedAt",
	Metadata:       "metadata",
	MetaCreatedAt:  "metaCreatedAt",
	MetaCreatedBy:  "metaCreatedBy",
	MetaUpdatedAt:  "metaUpdatedAt",
	MetaUpdatedBy:  "metaUpdatedBy",
	MetaDeletedAt:  "metaDeletedAt",
	MetaDeletedBy:  "metaDeletedBy",
}

func transformDisputeEvidenceDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(DisputeEvidenceDTOFieldName.Id):
		return string(model.DisputeEvidenceDBFieldName.Id), true

	case string(DisputeEvidenceDTOFieldName.DisputeId):
		return string(model.DisputeEvidenceDBFieldName.DisputeId), true

	case string(DisputeEvidenceDTOFieldName.EvidenceType):
		return string(model.DisputeEvidenceDBFieldName.EvidenceType), true

	case string(DisputeEvidenceDTOFieldName.StorageUri):
		return string(model.DisputeEvidenceDBFieldName.StorageUri), true

	case string(DisputeEvidenceDTOFieldName.ContentHash):
		return string(model.DisputeEvidenceDBFieldName.ContentHash), true

	case string(DisputeEvidenceDTOFieldName.IdempotencyKey):
		return string(model.DisputeEvidenceDBFieldName.IdempotencyKey), true

	case string(DisputeEvidenceDTOFieldName.SubmittedBy):
		return string(model.DisputeEvidenceDBFieldName.SubmittedBy), true

	case string(DisputeEvidenceDTOFieldName.SubmittedAt):
		return string(model.DisputeEvidenceDBFieldName.SubmittedAt), true

	case string(DisputeEvidenceDTOFieldName.Metadata):
		return string(model.DisputeEvidenceDBFieldName.Metadata), true

	case string(DisputeEvidenceDTOFieldName.MetaCreatedAt):
		return string(model.DisputeEvidenceDBFieldName.MetaCreatedAt), true

	case string(DisputeEvidenceDTOFieldName.MetaCreatedBy):
		return string(model.DisputeEvidenceDBFieldName.MetaCreatedBy), true

	case string(DisputeEvidenceDTOFieldName.MetaUpdatedAt):
		return string(model.DisputeEvidenceDBFieldName.MetaUpdatedAt), true

	case string(DisputeEvidenceDTOFieldName.MetaUpdatedBy):
		return string(model.DisputeEvidenceDBFieldName.MetaUpdatedBy), true

	case string(DisputeEvidenceDTOFieldName.MetaDeletedAt):
		return string(model.DisputeEvidenceDBFieldName.MetaDeletedAt), true

	case string(DisputeEvidenceDTOFieldName.MetaDeletedBy):
		return string(model.DisputeEvidenceDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewDisputeEvidenceFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isDisputeEvidenceBaseFilterField(field string) bool {
	spec, found := model.NewDisputeEvidenceFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeDisputeEvidenceProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateDisputeEvidenceProjectionOutputPath(path string) error {
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

func transformDisputeEvidenceFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformDisputeEvidenceDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformDisputeEvidenceFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformDisputeEvidenceFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformDisputeEvidenceDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isDisputeEvidenceBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateDisputeEvidenceProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeDisputeEvidenceProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformDisputeEvidenceDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformDisputeEvidenceDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformDisputeEvidenceFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultDisputeEvidenceFilter(filter *model.Filter) {
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
			Field: string(DisputeEvidenceDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type DisputeEvidenceSelectableResponse map[string]interface{}
type DisputeEvidenceSelectableListResponse []*DisputeEvidenceSelectableResponse

func assignDisputeEvidenceNestedValue(out map[string]interface{}, path string, value interface{}) {
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

func setDisputeEvidenceSelectableValue(out DisputeEvidenceSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignDisputeEvidenceNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewDisputeEvidenceSelectableResponse(disputeEvidence model.DisputeEvidence, filter model.Filter) DisputeEvidenceSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.DisputeEvidenceDBFieldName.Id),
			string(model.DisputeEvidenceDBFieldName.DisputeId),
			string(model.DisputeEvidenceDBFieldName.EvidenceType),
			string(model.DisputeEvidenceDBFieldName.StorageUri),
			string(model.DisputeEvidenceDBFieldName.ContentHash),
			string(model.DisputeEvidenceDBFieldName.IdempotencyKey),
			string(model.DisputeEvidenceDBFieldName.SubmittedBy),
			string(model.DisputeEvidenceDBFieldName.SubmittedAt),
			string(model.DisputeEvidenceDBFieldName.Metadata),
			string(model.DisputeEvidenceDBFieldName.MetaCreatedAt),
			string(model.DisputeEvidenceDBFieldName.MetaCreatedBy),
			string(model.DisputeEvidenceDBFieldName.MetaUpdatedAt),
			string(model.DisputeEvidenceDBFieldName.MetaUpdatedBy),
			string(model.DisputeEvidenceDBFieldName.MetaDeletedAt),
			string(model.DisputeEvidenceDBFieldName.MetaDeletedBy),
		)
	}
	disputeEvidenceSelectableResponse := DisputeEvidenceSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.DisputeEvidenceDBFieldName.Id):
			key := string(DisputeEvidenceDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.Id, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.DisputeId):
			key := string(DisputeEvidenceDTOFieldName.DisputeId)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.DisputeId, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.EvidenceType):
			key := string(DisputeEvidenceDTOFieldName.EvidenceType)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.EvidenceType, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.StorageUri):
			key := string(DisputeEvidenceDTOFieldName.StorageUri)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.StorageUri, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.ContentHash):
			key := string(DisputeEvidenceDTOFieldName.ContentHash)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.ContentHash.String, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.IdempotencyKey):
			key := string(DisputeEvidenceDTOFieldName.IdempotencyKey)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.IdempotencyKey.String, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.SubmittedBy):
			key := string(DisputeEvidenceDTOFieldName.SubmittedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.SubmittedBy, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.SubmittedAt):
			key := string(DisputeEvidenceDTOFieldName.SubmittedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.SubmittedAt, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.Metadata):
			key := string(DisputeEvidenceDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.Metadata, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.MetaCreatedAt):
			key := string(DisputeEvidenceDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.MetaCreatedAt, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.MetaCreatedBy):
			key := string(DisputeEvidenceDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.MetaCreatedBy, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.MetaUpdatedAt):
			key := string(DisputeEvidenceDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.MetaUpdatedAt, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.MetaUpdatedBy):
			key := string(DisputeEvidenceDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.MetaUpdatedBy, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.MetaDeletedAt):
			key := string(DisputeEvidenceDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.MetaDeletedAt.Time, explicitAlias)

		case string(model.DisputeEvidenceDBFieldName.MetaDeletedBy):
			key := string(DisputeEvidenceDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setDisputeEvidenceSelectableValue(disputeEvidenceSelectableResponse, key, disputeEvidence.MetaDeletedBy, explicitAlias)

		}
	}
	return disputeEvidenceSelectableResponse
}

func NewDisputeEvidenceListResponseFromFilterResult(result []model.DisputeEvidenceFilterResult, filter model.Filter) DisputeEvidenceSelectableListResponse {
	dtoDisputeEvidenceListResponse := DisputeEvidenceSelectableListResponse{}
	for _, row := range result {
		dtoDisputeEvidenceResponse := NewDisputeEvidenceSelectableResponse(row.DisputeEvidence, filter)
		dtoDisputeEvidenceListResponse = append(dtoDisputeEvidenceListResponse, &dtoDisputeEvidenceResponse)
	}
	return dtoDisputeEvidenceListResponse
}

type DisputeEvidenceFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     DisputeEvidenceSelectableListResponse `json:"data"`
}

func reverseDisputeEvidenceFilterResults(result []model.DisputeEvidenceFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewDisputeEvidenceFilterResponse(result []model.DisputeEvidenceFilterResult, filter model.Filter) (resp DisputeEvidenceFilterResponse) {
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
			reverseDisputeEvidenceFilterResults(dataResult)
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

	resp.Data = NewDisputeEvidenceListResponseFromFilterResult(dataResult, filter)
	return resp
}

type DisputeEvidenceCreateRequest struct {
	DisputeId      uuid.UUID       `json:"disputeId"`
	EvidenceType   string          `json:"evidenceType"`
	StorageUri     string          `json:"storageUri"`
	ContentHash    string          `json:"contentHash"`
	IdempotencyKey string          `json:"idempotencyKey"`
	SubmittedBy    uuid.UUID       `json:"submittedBy"`
	SubmittedAt    time.Time       `json:"submittedAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d *DisputeEvidenceCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *DisputeEvidenceCreateRequest) ToModel() model.DisputeEvidence {
	id, _ := uuid.NewV4()
	return model.DisputeEvidence{
		Id:             id,
		DisputeId:      d.DisputeId,
		EvidenceType:   d.EvidenceType,
		StorageUri:     d.StorageUri,
		ContentHash:    null.StringFrom(d.ContentHash),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		SubmittedBy:    d.SubmittedBy,
		SubmittedAt:    d.SubmittedAt,
		Metadata:       d.Metadata,
	}
}

type DisputeEvidenceListCreateRequest []*DisputeEvidenceCreateRequest

func (d DisputeEvidenceListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeEvidence := range d {
		err = validator.Struct(disputeEvidence)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputeEvidenceListCreateRequest) ToModelList() []model.DisputeEvidence {
	out := make([]model.DisputeEvidence, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type DisputeEvidenceUpdateRequest struct {
	DisputeId      uuid.UUID       `json:"disputeId"`
	EvidenceType   string          `json:"evidenceType"`
	StorageUri     string          `json:"storageUri"`
	ContentHash    string          `json:"contentHash"`
	IdempotencyKey string          `json:"idempotencyKey"`
	SubmittedBy    uuid.UUID       `json:"submittedBy"`
	SubmittedAt    time.Time       `json:"submittedAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d *DisputeEvidenceUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d DisputeEvidenceUpdateRequest) ToModel() model.DisputeEvidence {
	return model.DisputeEvidence{
		DisputeId:      d.DisputeId,
		EvidenceType:   d.EvidenceType,
		StorageUri:     d.StorageUri,
		ContentHash:    null.StringFrom(d.ContentHash),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		SubmittedBy:    d.SubmittedBy,
		SubmittedAt:    d.SubmittedAt,
		Metadata:       d.Metadata,
	}
}

type DisputeEvidenceBulkUpdateRequest struct {
	Id             uuid.UUID       `json:"id"`
	DisputeId      uuid.UUID       `json:"disputeId"`
	EvidenceType   string          `json:"evidenceType"`
	StorageUri     string          `json:"storageUri"`
	ContentHash    string          `json:"contentHash"`
	IdempotencyKey string          `json:"idempotencyKey"`
	SubmittedBy    uuid.UUID       `json:"submittedBy"`
	SubmittedAt    time.Time       `json:"submittedAt"`
	Metadata       json.RawMessage `json:"metadata"`
}

func (d DisputeEvidenceBulkUpdateRequest) PrimaryID() DisputeEvidencePrimaryID {
	return DisputeEvidencePrimaryID{
		Id: d.Id,
	}
}

type DisputeEvidenceListBulkUpdateRequest []*DisputeEvidenceBulkUpdateRequest

func (d DisputeEvidenceListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeEvidence := range d {
		err = validator.Struct(disputeEvidence)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputeEvidenceBulkUpdateRequest) ToModel() model.DisputeEvidence {
	return model.DisputeEvidence{
		Id:             d.Id,
		DisputeId:      d.DisputeId,
		EvidenceType:   d.EvidenceType,
		StorageUri:     d.StorageUri,
		ContentHash:    null.StringFrom(d.ContentHash),
		IdempotencyKey: null.StringFrom(d.IdempotencyKey),
		SubmittedBy:    d.SubmittedBy,
		SubmittedAt:    d.SubmittedAt,
		Metadata:       d.Metadata,
	}
}

type DisputeEvidenceResponse struct {
	Id             uuid.UUID       `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	DisputeId      uuid.UUID       `json:"disputeId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	EvidenceType   string          `json:"evidenceType" validate:"required"`
	StorageUri     string          `json:"storageUri" validate:"required"`
	ContentHash    string          `json:"contentHash"`
	IdempotencyKey string          `json:"idempotencyKey"`
	SubmittedBy    uuid.UUID       `json:"submittedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SubmittedAt    time.Time       `json:"submittedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Metadata       json.RawMessage `json:"metadata" swaggertype:"object"`
}

func NewDisputeEvidenceResponse(disputeEvidence model.DisputeEvidence) DisputeEvidenceResponse {
	return DisputeEvidenceResponse{
		Id:             disputeEvidence.Id,
		DisputeId:      disputeEvidence.DisputeId,
		EvidenceType:   disputeEvidence.EvidenceType,
		StorageUri:     disputeEvidence.StorageUri,
		ContentHash:    disputeEvidence.ContentHash.String,
		IdempotencyKey: disputeEvidence.IdempotencyKey.String,
		SubmittedBy:    disputeEvidence.SubmittedBy,
		SubmittedAt:    disputeEvidence.SubmittedAt,
		Metadata:       disputeEvidence.Metadata,
	}
}

type DisputeEvidenceListResponse []*DisputeEvidenceResponse

func NewDisputeEvidenceListResponse(disputeEvidenceList model.DisputeEvidenceList) DisputeEvidenceListResponse {
	dtoDisputeEvidenceListResponse := DisputeEvidenceListResponse{}
	for _, disputeEvidence := range disputeEvidenceList {
		dtoDisputeEvidenceResponse := NewDisputeEvidenceResponse(*disputeEvidence)
		dtoDisputeEvidenceListResponse = append(dtoDisputeEvidenceListResponse, &dtoDisputeEvidenceResponse)
	}
	return dtoDisputeEvidenceListResponse
}

type DisputeEvidencePrimaryIDList []DisputeEvidencePrimaryID

func (d DisputeEvidencePrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputeEvidence := range d {
		err = validator.Struct(disputeEvidence)
		if err != nil {
			return
		}
	}
	return nil
}

type DisputeEvidencePrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *DisputeEvidencePrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d DisputeEvidencePrimaryID) ToModel() model.DisputeEvidencePrimaryID {
	return model.DisputeEvidencePrimaryID{
		Id: d.Id,
	}
}
