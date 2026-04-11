package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/idempotency/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type IdempotencyKeysDTOFieldNameType string

type idempotencyKeysDTOFieldName struct {
	Id              IdempotencyKeysDTOFieldNameType
	IdempotencyKey  IdempotencyKeysDTOFieldNameType
	MerchantId      IdempotencyKeysDTOFieldNameType
	RequestPath     IdempotencyKeysDTOFieldNameType
	RequestBodyHash IdempotencyKeysDTOFieldNameType
	ResponseStatus  IdempotencyKeysDTOFieldNameType
	ResponseBody    IdempotencyKeysDTOFieldNameType
	LockedAt        IdempotencyKeysDTOFieldNameType
	LockedUntil     IdempotencyKeysDTOFieldNameType
	CompletedAt     IdempotencyKeysDTOFieldNameType
	ExpiresAt       IdempotencyKeysDTOFieldNameType
	MetaCreatedAt   IdempotencyKeysDTOFieldNameType
	MetaCreatedBy   IdempotencyKeysDTOFieldNameType
	MetaUpdatedAt   IdempotencyKeysDTOFieldNameType
	MetaUpdatedBy   IdempotencyKeysDTOFieldNameType
	MetaDeletedAt   IdempotencyKeysDTOFieldNameType
	MetaDeletedBy   IdempotencyKeysDTOFieldNameType
}

var IdempotencyKeysDTOFieldName = idempotencyKeysDTOFieldName{
	Id:              "id",
	IdempotencyKey:  "idempotencyKey",
	MerchantId:      "merchantId",
	RequestPath:     "requestPath",
	RequestBodyHash: "requestBodyHash",
	ResponseStatus:  "responseStatus",
	ResponseBody:    "responseBody",
	LockedAt:        "lockedAt",
	LockedUntil:     "lockedUntil",
	CompletedAt:     "completedAt",
	ExpiresAt:       "expiresAt",
	MetaCreatedAt:   "metaCreatedAt",
	MetaCreatedBy:   "metaCreatedBy",
	MetaUpdatedAt:   "metaUpdatedAt",
	MetaUpdatedBy:   "metaUpdatedBy",
	MetaDeletedAt:   "metaDeletedAt",
	MetaDeletedBy:   "metaDeletedBy",
}

func NewIdempotencyKeysListResponseFromFilterResult(result []model.IdempotencyKeysFilterResult, filter model.Filter) IdempotencyKeysSelectableListResponse {
	dtoIdempotencyKeysListResponse := IdempotencyKeysSelectableListResponse{}
	for _, idempotencyKeys := range result {
		dtoIdempotencyKeysResponse := NewIdempotencyKeysSelectableResponse(idempotencyKeys.IdempotencyKeys, filter)
		dtoIdempotencyKeysListResponse = append(dtoIdempotencyKeysListResponse, &dtoIdempotencyKeysResponse)
	}
	return dtoIdempotencyKeysListResponse
}

func transformIdempotencyKeysDTOFieldNameFromStr(field string) (dbField model.IdempotencyKeysDBFieldNameType, found bool) {
	switch field {

	case string(IdempotencyKeysDTOFieldName.Id):
		return model.IdempotencyKeysDBFieldName.Id, true

	case string(IdempotencyKeysDTOFieldName.IdempotencyKey):
		return model.IdempotencyKeysDBFieldName.IdempotencyKey, true

	case string(IdempotencyKeysDTOFieldName.MerchantId):
		return model.IdempotencyKeysDBFieldName.MerchantId, true

	case string(IdempotencyKeysDTOFieldName.RequestPath):
		return model.IdempotencyKeysDBFieldName.RequestPath, true

	case string(IdempotencyKeysDTOFieldName.RequestBodyHash):
		return model.IdempotencyKeysDBFieldName.RequestBodyHash, true

	case string(IdempotencyKeysDTOFieldName.ResponseStatus):
		return model.IdempotencyKeysDBFieldName.ResponseStatus, true

	case string(IdempotencyKeysDTOFieldName.ResponseBody):
		return model.IdempotencyKeysDBFieldName.ResponseBody, true

	case string(IdempotencyKeysDTOFieldName.LockedAt):
		return model.IdempotencyKeysDBFieldName.LockedAt, true

	case string(IdempotencyKeysDTOFieldName.LockedUntil):
		return model.IdempotencyKeysDBFieldName.LockedUntil, true

	case string(IdempotencyKeysDTOFieldName.CompletedAt):
		return model.IdempotencyKeysDBFieldName.CompletedAt, true

	case string(IdempotencyKeysDTOFieldName.ExpiresAt):
		return model.IdempotencyKeysDBFieldName.ExpiresAt, true

	case string(IdempotencyKeysDTOFieldName.MetaCreatedAt):
		return model.IdempotencyKeysDBFieldName.MetaCreatedAt, true

	case string(IdempotencyKeysDTOFieldName.MetaCreatedBy):
		return model.IdempotencyKeysDBFieldName.MetaCreatedBy, true

	case string(IdempotencyKeysDTOFieldName.MetaUpdatedAt):
		return model.IdempotencyKeysDBFieldName.MetaUpdatedAt, true

	case string(IdempotencyKeysDTOFieldName.MetaUpdatedBy):
		return model.IdempotencyKeysDBFieldName.MetaUpdatedBy, true

	case string(IdempotencyKeysDTOFieldName.MetaDeletedAt):
		return model.IdempotencyKeysDBFieldName.MetaDeletedAt, true

	case string(IdempotencyKeysDTOFieldName.MetaDeletedBy):
		return model.IdempotencyKeysDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformIdempotencyKeysFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformIdempotencyKeysDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformIdempotencyKeysDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformIdempotencyKeysDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultIdempotencyKeysFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(IdempotencyKeysDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type IdempotencyKeysSelectableResponse map[string]interface{}
type IdempotencyKeysSelectableListResponse []*IdempotencyKeysSelectableResponse

func NewIdempotencyKeysSelectableResponse(idempotencyKeys model.IdempotencyKeys, filter model.Filter) IdempotencyKeysSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.IdempotencyKeysDBFieldName.Id),
			string(model.IdempotencyKeysDBFieldName.IdempotencyKey),
			string(model.IdempotencyKeysDBFieldName.MerchantId),
			string(model.IdempotencyKeysDBFieldName.RequestPath),
			string(model.IdempotencyKeysDBFieldName.RequestBodyHash),
			string(model.IdempotencyKeysDBFieldName.ResponseStatus),
			string(model.IdempotencyKeysDBFieldName.ResponseBody),
			string(model.IdempotencyKeysDBFieldName.LockedAt),
			string(model.IdempotencyKeysDBFieldName.LockedUntil),
			string(model.IdempotencyKeysDBFieldName.CompletedAt),
			string(model.IdempotencyKeysDBFieldName.ExpiresAt),
			string(model.IdempotencyKeysDBFieldName.MetaCreatedAt),
			string(model.IdempotencyKeysDBFieldName.MetaCreatedBy),
			string(model.IdempotencyKeysDBFieldName.MetaUpdatedAt),
			string(model.IdempotencyKeysDBFieldName.MetaUpdatedBy),
			string(model.IdempotencyKeysDBFieldName.MetaDeletedAt),
			string(model.IdempotencyKeysDBFieldName.MetaDeletedBy),
		)
	}
	idempotencyKeysSelectableResponse := IdempotencyKeysSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.IdempotencyKeysDBFieldName.Id):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.Id)] = idempotencyKeys.Id

		case string(model.IdempotencyKeysDBFieldName.IdempotencyKey):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.IdempotencyKey)] = idempotencyKeys.IdempotencyKey

		case string(model.IdempotencyKeysDBFieldName.MerchantId):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.MerchantId)] = idempotencyKeys.MerchantId

		case string(model.IdempotencyKeysDBFieldName.RequestPath):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.RequestPath)] = idempotencyKeys.RequestPath

		case string(model.IdempotencyKeysDBFieldName.RequestBodyHash):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.RequestBodyHash)] = idempotencyKeys.RequestBodyHash

		case string(model.IdempotencyKeysDBFieldName.ResponseStatus):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.ResponseStatus)] = idempotencyKeys.ResponseStatus

		case string(model.IdempotencyKeysDBFieldName.ResponseBody):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.ResponseBody)] = idempotencyKeys.ResponseBody

		case string(model.IdempotencyKeysDBFieldName.LockedAt):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.LockedAt)] = idempotencyKeys.LockedAt

		case string(model.IdempotencyKeysDBFieldName.LockedUntil):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.LockedUntil)] = idempotencyKeys.LockedUntil

		case string(model.IdempotencyKeysDBFieldName.CompletedAt):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.CompletedAt)] = idempotencyKeys.CompletedAt

		case string(model.IdempotencyKeysDBFieldName.ExpiresAt):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.ExpiresAt)] = idempotencyKeys.ExpiresAt

		case string(model.IdempotencyKeysDBFieldName.MetaCreatedAt):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.MetaCreatedAt)] = idempotencyKeys.MetaCreatedAt

		case string(model.IdempotencyKeysDBFieldName.MetaCreatedBy):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.MetaCreatedBy)] = idempotencyKeys.MetaCreatedBy

		case string(model.IdempotencyKeysDBFieldName.MetaUpdatedAt):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.MetaUpdatedAt)] = idempotencyKeys.MetaUpdatedAt

		case string(model.IdempotencyKeysDBFieldName.MetaUpdatedBy):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.MetaUpdatedBy)] = idempotencyKeys.MetaUpdatedBy

		case string(model.IdempotencyKeysDBFieldName.MetaDeletedAt):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.MetaDeletedAt)] = idempotencyKeys.MetaDeletedAt

		case string(model.IdempotencyKeysDBFieldName.MetaDeletedBy):
			idempotencyKeysSelectableResponse[string(IdempotencyKeysDTOFieldName.MetaDeletedBy)] = idempotencyKeys.MetaDeletedBy

		}
	}
	return idempotencyKeysSelectableResponse
}

type IdempotencyKeysFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     IdempotencyKeysSelectableListResponse `json:"data"`
}

func NewIdempotencyKeysFilterResponse(result []model.IdempotencyKeysFilterResult, filter model.Filter) (resp IdempotencyKeysFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewIdempotencyKeysListResponseFromFilterResult(result, filter)
	return resp
}

type IdempotencyKeysCreateRequest struct {
	IdempotencyKey  string          `json:"idempotencyKey"`
	MerchantId      uuid.UUID       `json:"merchantId"`
	RequestPath     string          `json:"requestPath"`
	RequestBodyHash string          `json:"requestBodyHash"`
	ResponseStatus  int16           `json:"responseStatus"`
	ResponseBody    json.RawMessage `json:"responseBody"`
	LockedAt        time.Time       `json:"lockedAt"`
	LockedUntil     time.Time       `json:"lockedUntil"`
	CompletedAt     time.Time       `json:"completedAt"`
	ExpiresAt       time.Time       `json:"expiresAt"`
	MetaCreatedAt   time.Time       `json:"metaCreatedAt"`
	MetaCreatedBy   uuid.UUID       `json:"metaCreatedBy"`
	MetaUpdatedAt   time.Time       `json:"metaUpdatedAt"`
	MetaUpdatedBy   uuid.UUID       `json:"metaUpdatedBy"`
	MetaDeletedAt   time.Time       `json:"metaDeletedAt"`
	MetaDeletedBy   uuid.UUID       `json:"metaDeletedBy"`
}

func (d *IdempotencyKeysCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *IdempotencyKeysCreateRequest) ToModel() model.IdempotencyKeys {
	id, _ := uuid.NewV4()
	return model.IdempotencyKeys{
		Id:              id,
		IdempotencyKey:  d.IdempotencyKey,
		MerchantId:      d.MerchantId,
		RequestPath:     d.RequestPath,
		RequestBodyHash: d.RequestBodyHash,
		ResponseStatus:  null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:    d.ResponseBody,
		LockedAt:        null.TimeFrom(d.LockedAt),
		LockedUntil:     null.TimeFrom(d.LockedUntil),
		CompletedAt:     null.TimeFrom(d.CompletedAt),
		ExpiresAt:       d.ExpiresAt,
		MetaCreatedAt:   d.MetaCreatedAt,
		MetaCreatedBy:   d.MetaCreatedBy,
		MetaUpdatedAt:   d.MetaUpdatedAt,
		MetaUpdatedBy:   nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:   null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:   nuuid.From(d.MetaDeletedBy),
	}
}

type IdempotencyKeysListCreateRequest []*IdempotencyKeysCreateRequest

func (d IdempotencyKeysListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, idempotencyKeys := range d {
		err = validator.Struct(idempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

func (d IdempotencyKeysListCreateRequest) ToModelList() []model.IdempotencyKeys {
	out := make([]model.IdempotencyKeys, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type IdempotencyKeysUpdateRequest struct {
	IdempotencyKey  string          `json:"idempotencyKey"`
	MerchantId      uuid.UUID       `json:"merchantId"`
	RequestPath     string          `json:"requestPath"`
	RequestBodyHash string          `json:"requestBodyHash"`
	ResponseStatus  int16           `json:"responseStatus"`
	ResponseBody    json.RawMessage `json:"responseBody"`
	LockedAt        time.Time       `json:"lockedAt"`
	LockedUntil     time.Time       `json:"lockedUntil"`
	CompletedAt     time.Time       `json:"completedAt"`
	ExpiresAt       time.Time       `json:"expiresAt"`
	MetaCreatedAt   time.Time       `json:"metaCreatedAt"`
	MetaCreatedBy   uuid.UUID       `json:"metaCreatedBy"`
	MetaUpdatedAt   time.Time       `json:"metaUpdatedAt"`
	MetaUpdatedBy   uuid.UUID       `json:"metaUpdatedBy"`
	MetaDeletedAt   time.Time       `json:"metaDeletedAt"`
	MetaDeletedBy   uuid.UUID       `json:"metaDeletedBy"`
}

func (d *IdempotencyKeysUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d IdempotencyKeysUpdateRequest) ToModel() model.IdempotencyKeys {
	return model.IdempotencyKeys{
		IdempotencyKey:  d.IdempotencyKey,
		MerchantId:      d.MerchantId,
		RequestPath:     d.RequestPath,
		RequestBodyHash: d.RequestBodyHash,
		ResponseStatus:  null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:    d.ResponseBody,
		LockedAt:        null.TimeFrom(d.LockedAt),
		LockedUntil:     null.TimeFrom(d.LockedUntil),
		CompletedAt:     null.TimeFrom(d.CompletedAt),
		ExpiresAt:       d.ExpiresAt,
		MetaCreatedAt:   d.MetaCreatedAt,
		MetaCreatedBy:   d.MetaCreatedBy,
		MetaUpdatedAt:   d.MetaUpdatedAt,
		MetaUpdatedBy:   nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:   null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:   nuuid.From(d.MetaDeletedBy),
	}
}

type IdempotencyKeysBulkUpdateRequest struct {
	Id              uuid.UUID       `json:"id"`
	IdempotencyKey  string          `json:"idempotencyKey"`
	MerchantId      uuid.UUID       `json:"merchantId"`
	RequestPath     string          `json:"requestPath"`
	RequestBodyHash string          `json:"requestBodyHash"`
	ResponseStatus  int16           `json:"responseStatus"`
	ResponseBody    json.RawMessage `json:"responseBody"`
	LockedAt        time.Time       `json:"lockedAt"`
	LockedUntil     time.Time       `json:"lockedUntil"`
	CompletedAt     time.Time       `json:"completedAt"`
	ExpiresAt       time.Time       `json:"expiresAt"`
	MetaCreatedAt   time.Time       `json:"metaCreatedAt"`
	MetaCreatedBy   uuid.UUID       `json:"metaCreatedBy"`
	MetaUpdatedAt   time.Time       `json:"metaUpdatedAt"`
	MetaUpdatedBy   uuid.UUID       `json:"metaUpdatedBy"`
	MetaDeletedAt   time.Time       `json:"metaDeletedAt"`
	MetaDeletedBy   uuid.UUID       `json:"metaDeletedBy"`
}

func (d IdempotencyKeysBulkUpdateRequest) PrimaryID() IdempotencyKeysPrimaryID {
	return IdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

type IdempotencyKeysListBulkUpdateRequest []*IdempotencyKeysBulkUpdateRequest

func (d IdempotencyKeysListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, idempotencyKeys := range d {
		err = validator.Struct(idempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

func (d IdempotencyKeysBulkUpdateRequest) ToModel() model.IdempotencyKeys {
	return model.IdempotencyKeys{
		Id:              d.Id,
		IdempotencyKey:  d.IdempotencyKey,
		MerchantId:      d.MerchantId,
		RequestPath:     d.RequestPath,
		RequestBodyHash: d.RequestBodyHash,
		ResponseStatus:  null.IntFrom(int64(d.ResponseStatus)),
		ResponseBody:    d.ResponseBody,
		LockedAt:        null.TimeFrom(d.LockedAt),
		LockedUntil:     null.TimeFrom(d.LockedUntil),
		CompletedAt:     null.TimeFrom(d.CompletedAt),
		ExpiresAt:       d.ExpiresAt,
		MetaCreatedAt:   d.MetaCreatedAt,
		MetaCreatedBy:   d.MetaCreatedBy,
		MetaUpdatedAt:   d.MetaUpdatedAt,
		MetaUpdatedBy:   nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:   null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:   nuuid.From(d.MetaDeletedBy),
	}
}

type IdempotencyKeysResponse struct {
	Id              uuid.UUID       `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IdempotencyKey  string          `json:"idempotencyKey" validate:"required"`
	MerchantId      uuid.UUID       `json:"merchantId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	RequestPath     string          `json:"requestPath" validate:"required"`
	RequestBodyHash string          `json:"requestBodyHash" validate:"required"`
	ResponseStatus  int16           `json:"responseStatus"`
	ResponseBody    json.RawMessage `json:"responseBody" validate:"required" swaggertype:"object"`
	LockedAt        time.Time       `json:"lockedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	LockedUntil     time.Time       `json:"lockedUntil" format:"date-time" example:"2024-01-01T00:00:00Z"`
	CompletedAt     time.Time       `json:"completedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ExpiresAt       time.Time       `json:"expiresAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedAt   time.Time       `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy   uuid.UUID       `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt   time.Time       `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy   uuid.UUID       `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt   time.Time       `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy   uuid.UUID       `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewIdempotencyKeysResponse(idempotencyKeys model.IdempotencyKeys) IdempotencyKeysResponse {
	return IdempotencyKeysResponse{
		Id:              idempotencyKeys.Id,
		IdempotencyKey:  idempotencyKeys.IdempotencyKey,
		MerchantId:      idempotencyKeys.MerchantId,
		RequestPath:     idempotencyKeys.RequestPath,
		RequestBodyHash: idempotencyKeys.RequestBodyHash,
		ResponseStatus:  int16(idempotencyKeys.ResponseStatus.ValueOrZero()),
		ResponseBody:    idempotencyKeys.ResponseBody,
		LockedAt:        idempotencyKeys.LockedAt.Time,
		LockedUntil:     idempotencyKeys.LockedUntil.Time,
		CompletedAt:     idempotencyKeys.CompletedAt.Time,
		ExpiresAt:       idempotencyKeys.ExpiresAt,
		MetaCreatedAt:   idempotencyKeys.MetaCreatedAt,
		MetaCreatedBy:   idempotencyKeys.MetaCreatedBy,
		MetaUpdatedAt:   idempotencyKeys.MetaUpdatedAt,
		MetaUpdatedBy:   idempotencyKeys.MetaUpdatedBy.UUID,
		MetaDeletedAt:   idempotencyKeys.MetaDeletedAt.Time,
		MetaDeletedBy:   idempotencyKeys.MetaDeletedBy.UUID,
	}
}

type IdempotencyKeysListResponse []*IdempotencyKeysResponse

func NewIdempotencyKeysListResponse(idempotencyKeysList model.IdempotencyKeysList) IdempotencyKeysListResponse {
	dtoIdempotencyKeysListResponse := IdempotencyKeysListResponse{}
	for _, idempotencyKeys := range idempotencyKeysList {
		dtoIdempotencyKeysResponse := NewIdempotencyKeysResponse(*idempotencyKeys)
		dtoIdempotencyKeysListResponse = append(dtoIdempotencyKeysListResponse, &dtoIdempotencyKeysResponse)
	}
	return dtoIdempotencyKeysListResponse
}

type IdempotencyKeysPrimaryIDList []IdempotencyKeysPrimaryID

func (d IdempotencyKeysPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, idempotencyKeys := range d {
		err = validator.Struct(idempotencyKeys)
		if err != nil {
			return
		}
	}
	return nil
}

type IdempotencyKeysPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *IdempotencyKeysPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d IdempotencyKeysPrimaryID) ToModel() model.IdempotencyKeysPrimaryID {
	return model.IdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}
