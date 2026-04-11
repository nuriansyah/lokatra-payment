package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/psp/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PspWebhooksDTOFieldNameType string

type pspWebhooksDTOFieldName struct {
	Id                 PspWebhooksDTOFieldNameType
	PspAccountId       PspWebhooksDTOFieldNameType
	Psp                PspWebhooksDTOFieldNameType
	PspEventId         PspWebhooksDTOFieldNameType
	PspEventType       PspWebhooksDTOFieldNameType
	ReceivedAt         PspWebhooksDTOFieldNameType
	Headers            PspWebhooksDTOFieldNameType
	RawPayload         PspWebhooksDTOFieldNameType
	HmacValid          PspWebhooksDTOFieldNameType
	Status             PspWebhooksDTOFieldNameType
	ProcessingAttempts PspWebhooksDTOFieldNameType
	LastError          PspWebhooksDTOFieldNameType
	ProcessedAt        PspWebhooksDTOFieldNameType
	ResolvedPaymentId  PspWebhooksDTOFieldNameType
	ResolvedIntentId   PspWebhooksDTOFieldNameType
	MetaCreatedAt      PspWebhooksDTOFieldNameType
	MetaCreatedBy      PspWebhooksDTOFieldNameType
	MetaUpdatedAt      PspWebhooksDTOFieldNameType
	MetaUpdatedBy      PspWebhooksDTOFieldNameType
}

var PspWebhooksDTOFieldName = pspWebhooksDTOFieldName{
	Id:                 "id",
	PspAccountId:       "pspAccountId",
	Psp:                "psp",
	PspEventId:         "pspEventId",
	PspEventType:       "pspEventType",
	ReceivedAt:         "receivedAt",
	Headers:            "headers",
	RawPayload:         "rawPayload",
	HmacValid:          "hmacValid",
	Status:             "status",
	ProcessingAttempts: "processingAttempts",
	LastError:          "lastError",
	ProcessedAt:        "processedAt",
	ResolvedPaymentId:  "resolvedPaymentId",
	ResolvedIntentId:   "resolvedIntentId",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
}

func NewPspWebhooksListResponseFromFilterResult(result []model.PspWebhooksFilterResult, filter model.Filter) PspWebhooksSelectableListResponse {
	dtoPspWebhooksListResponse := PspWebhooksSelectableListResponse{}
	for _, pspWebhooks := range result {
		dtoPspWebhooksResponse := NewPspWebhooksSelectableResponse(pspWebhooks.PspWebhooks, filter)
		dtoPspWebhooksListResponse = append(dtoPspWebhooksListResponse, &dtoPspWebhooksResponse)
	}
	return dtoPspWebhooksListResponse
}

func transformPspWebhooksDTOFieldNameFromStr(field string) (dbField model.PspWebhooksDBFieldNameType, found bool) {
	switch field {

	case string(PspWebhooksDTOFieldName.Id):
		return model.PspWebhooksDBFieldName.Id, true

	case string(PspWebhooksDTOFieldName.PspAccountId):
		return model.PspWebhooksDBFieldName.PspAccountId, true

	case string(PspWebhooksDTOFieldName.Psp):
		return model.PspWebhooksDBFieldName.Psp, true

	case string(PspWebhooksDTOFieldName.PspEventId):
		return model.PspWebhooksDBFieldName.PspEventId, true

	case string(PspWebhooksDTOFieldName.PspEventType):
		return model.PspWebhooksDBFieldName.PspEventType, true

	case string(PspWebhooksDTOFieldName.ReceivedAt):
		return model.PspWebhooksDBFieldName.ReceivedAt, true

	case string(PspWebhooksDTOFieldName.Headers):
		return model.PspWebhooksDBFieldName.Headers, true

	case string(PspWebhooksDTOFieldName.RawPayload):
		return model.PspWebhooksDBFieldName.RawPayload, true

	case string(PspWebhooksDTOFieldName.HmacValid):
		return model.PspWebhooksDBFieldName.HmacValid, true

	case string(PspWebhooksDTOFieldName.Status):
		return model.PspWebhooksDBFieldName.Status, true

	case string(PspWebhooksDTOFieldName.ProcessingAttempts):
		return model.PspWebhooksDBFieldName.ProcessingAttempts, true

	case string(PspWebhooksDTOFieldName.LastError):
		return model.PspWebhooksDBFieldName.LastError, true

	case string(PspWebhooksDTOFieldName.ProcessedAt):
		return model.PspWebhooksDBFieldName.ProcessedAt, true

	case string(PspWebhooksDTOFieldName.ResolvedPaymentId):
		return model.PspWebhooksDBFieldName.ResolvedPaymentId, true

	case string(PspWebhooksDTOFieldName.ResolvedIntentId):
		return model.PspWebhooksDBFieldName.ResolvedIntentId, true

	case string(PspWebhooksDTOFieldName.MetaCreatedAt):
		return model.PspWebhooksDBFieldName.MetaCreatedAt, true

	case string(PspWebhooksDTOFieldName.MetaCreatedBy):
		return model.PspWebhooksDBFieldName.MetaCreatedBy, true

	case string(PspWebhooksDTOFieldName.MetaUpdatedAt):
		return model.PspWebhooksDBFieldName.MetaUpdatedAt, true

	case string(PspWebhooksDTOFieldName.MetaUpdatedBy):
		return model.PspWebhooksDBFieldName.MetaUpdatedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformPspWebhooksFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformPspWebhooksDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformPspWebhooksDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformPspWebhooksDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultPspWebhooksFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(PspWebhooksDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type PspWebhooksSelectableResponse map[string]interface{}
type PspWebhooksSelectableListResponse []*PspWebhooksSelectableResponse

func NewPspWebhooksSelectableResponse(pspWebhooks model.PspWebhooks, filter model.Filter) PspWebhooksSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.PspWebhooksDBFieldName.Id),
			string(model.PspWebhooksDBFieldName.PspAccountId),
			string(model.PspWebhooksDBFieldName.Psp),
			string(model.PspWebhooksDBFieldName.PspEventId),
			string(model.PspWebhooksDBFieldName.PspEventType),
			string(model.PspWebhooksDBFieldName.ReceivedAt),
			string(model.PspWebhooksDBFieldName.Headers),
			string(model.PspWebhooksDBFieldName.RawPayload),
			string(model.PspWebhooksDBFieldName.HmacValid),
			string(model.PspWebhooksDBFieldName.Status),
			string(model.PspWebhooksDBFieldName.ProcessingAttempts),
			string(model.PspWebhooksDBFieldName.LastError),
			string(model.PspWebhooksDBFieldName.ProcessedAt),
			string(model.PspWebhooksDBFieldName.ResolvedPaymentId),
			string(model.PspWebhooksDBFieldName.ResolvedIntentId),
			string(model.PspWebhooksDBFieldName.MetaCreatedAt),
			string(model.PspWebhooksDBFieldName.MetaCreatedBy),
			string(model.PspWebhooksDBFieldName.MetaUpdatedAt),
			string(model.PspWebhooksDBFieldName.MetaUpdatedBy),
		)
	}
	pspWebhooksSelectableResponse := PspWebhooksSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.PspWebhooksDBFieldName.Id):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.Id)] = pspWebhooks.Id

		case string(model.PspWebhooksDBFieldName.PspAccountId):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.PspAccountId)] = pspWebhooks.PspAccountId

		case string(model.PspWebhooksDBFieldName.Psp):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.Psp)] = pspWebhooks.Psp

		case string(model.PspWebhooksDBFieldName.PspEventId):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.PspEventId)] = pspWebhooks.PspEventId

		case string(model.PspWebhooksDBFieldName.PspEventType):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.PspEventType)] = pspWebhooks.PspEventType

		case string(model.PspWebhooksDBFieldName.ReceivedAt):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.ReceivedAt)] = pspWebhooks.ReceivedAt

		case string(model.PspWebhooksDBFieldName.Headers):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.Headers)] = pspWebhooks.Headers

		case string(model.PspWebhooksDBFieldName.RawPayload):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.RawPayload)] = pspWebhooks.RawPayload

		case string(model.PspWebhooksDBFieldName.HmacValid):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.HmacValid)] = pspWebhooks.HmacValid

		case string(model.PspWebhooksDBFieldName.Status):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.Status)] = pspWebhooks.Status

		case string(model.PspWebhooksDBFieldName.ProcessingAttempts):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.ProcessingAttempts)] = pspWebhooks.ProcessingAttempts

		case string(model.PspWebhooksDBFieldName.LastError):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.LastError)] = pspWebhooks.LastError

		case string(model.PspWebhooksDBFieldName.ProcessedAt):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.ProcessedAt)] = pspWebhooks.ProcessedAt

		case string(model.PspWebhooksDBFieldName.ResolvedPaymentId):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.ResolvedPaymentId)] = pspWebhooks.ResolvedPaymentId

		case string(model.PspWebhooksDBFieldName.ResolvedIntentId):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.ResolvedIntentId)] = pspWebhooks.ResolvedIntentId

		case string(model.PspWebhooksDBFieldName.MetaCreatedAt):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.MetaCreatedAt)] = pspWebhooks.MetaCreatedAt

		case string(model.PspWebhooksDBFieldName.MetaCreatedBy):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.MetaCreatedBy)] = pspWebhooks.MetaCreatedBy

		case string(model.PspWebhooksDBFieldName.MetaUpdatedAt):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.MetaUpdatedAt)] = pspWebhooks.MetaUpdatedAt

		case string(model.PspWebhooksDBFieldName.MetaUpdatedBy):
			pspWebhooksSelectableResponse[string(PspWebhooksDTOFieldName.MetaUpdatedBy)] = pspWebhooks.MetaUpdatedBy

		}
	}
	return pspWebhooksSelectableResponse
}

type PspWebhooksFilterResponse struct {
	Metadata Metadata                          `json:"metadata"`
	Data     PspWebhooksSelectableListResponse `json:"data"`
}

func NewPspWebhooksFilterResponse(result []model.PspWebhooksFilterResult, filter model.Filter) (resp PspWebhooksFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewPspWebhooksListResponseFromFilterResult(result, filter)
	return resp
}

type PspWebhooksCreateRequest struct {
	PspAccountId       uuid.UUID           `json:"pspAccountId"`
	Psp                string              `json:"psp"`
	PspEventId         string              `json:"pspEventId"`
	PspEventType       string              `json:"pspEventType"`
	ReceivedAt         time.Time           `json:"receivedAt"`
	Headers            json.RawMessage     `json:"headers"`
	RawPayload         json.RawMessage     `json:"rawPayload"`
	HmacValid          bool                `json:"hmacValid"`
	Status             model.WebhookStatus `json:"status" example:"RECEIVED" enums:"RECEIVED,PROCESSING,PROCESSED,FAILED,IGNORED"`
	ProcessingAttempts int                 `json:"processingAttempts"`
	LastError          string              `json:"lastError"`
	ProcessedAt        time.Time           `json:"processedAt"`
	ResolvedPaymentId  uuid.UUID           `json:"resolvedPaymentId"`
	ResolvedIntentId   uuid.UUID           `json:"resolvedIntentId"`
	MetaCreatedAt      time.Time           `json:"metaCreatedAt"`
	MetaCreatedBy      uuid.UUID           `json:"metaCreatedBy"`
	MetaUpdatedAt      time.Time           `json:"metaUpdatedAt"`
	MetaUpdatedBy      uuid.UUID           `json:"metaUpdatedBy"`
}

func (d *PspWebhooksCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *PspWebhooksCreateRequest) ToModel() model.PspWebhooks {
	id, _ := uuid.NewV4()
	return model.PspWebhooks{
		Id:                 id,
		PspAccountId:       d.PspAccountId,
		Psp:                d.Psp,
		PspEventId:         null.StringFrom(d.PspEventId),
		PspEventType:       d.PspEventType,
		ReceivedAt:         d.ReceivedAt,
		Headers:            d.Headers,
		RawPayload:         d.RawPayload,
		HmacValid:          null.BoolFrom(d.HmacValid),
		Status:             d.Status,
		ProcessingAttempts: d.ProcessingAttempts,
		LastError:          null.StringFrom(d.LastError),
		ProcessedAt:        null.TimeFrom(d.ProcessedAt),
		ResolvedPaymentId:  nuuid.From(d.ResolvedPaymentId),
		ResolvedIntentId:   nuuid.From(d.ResolvedIntentId),
		MetaCreatedAt:      d.MetaCreatedAt,
		MetaCreatedBy:      d.MetaCreatedBy,
		MetaUpdatedAt:      d.MetaUpdatedAt,
		MetaUpdatedBy:      nuuid.From(d.MetaUpdatedBy),
	}
}

type PspWebhooksListCreateRequest []*PspWebhooksCreateRequest

func (d PspWebhooksListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, pspWebhooks := range d {
		err = validator.Struct(pspWebhooks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PspWebhooksListCreateRequest) ToModelList() []model.PspWebhooks {
	out := make([]model.PspWebhooks, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type PspWebhooksUpdateRequest struct {
	PspAccountId       uuid.UUID           `json:"pspAccountId"`
	Psp                string              `json:"psp"`
	PspEventId         string              `json:"pspEventId"`
	PspEventType       string              `json:"pspEventType"`
	ReceivedAt         time.Time           `json:"receivedAt"`
	Headers            json.RawMessage     `json:"headers"`
	RawPayload         json.RawMessage     `json:"rawPayload"`
	HmacValid          bool                `json:"hmacValid"`
	Status             model.WebhookStatus `json:"status" example:"RECEIVED" enums:"RECEIVED,PROCESSING,PROCESSED,FAILED,IGNORED"`
	ProcessingAttempts int                 `json:"processingAttempts"`
	LastError          string              `json:"lastError"`
	ProcessedAt        time.Time           `json:"processedAt"`
	ResolvedPaymentId  uuid.UUID           `json:"resolvedPaymentId"`
	ResolvedIntentId   uuid.UUID           `json:"resolvedIntentId"`
	MetaCreatedAt      time.Time           `json:"metaCreatedAt"`
	MetaCreatedBy      uuid.UUID           `json:"metaCreatedBy"`
	MetaUpdatedAt      time.Time           `json:"metaUpdatedAt"`
	MetaUpdatedBy      uuid.UUID           `json:"metaUpdatedBy"`
}

func (d *PspWebhooksUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d PspWebhooksUpdateRequest) ToModel() model.PspWebhooks {
	return model.PspWebhooks{
		PspAccountId:       d.PspAccountId,
		Psp:                d.Psp,
		PspEventId:         null.StringFrom(d.PspEventId),
		PspEventType:       d.PspEventType,
		ReceivedAt:         d.ReceivedAt,
		Headers:            d.Headers,
		RawPayload:         d.RawPayload,
		HmacValid:          null.BoolFrom(d.HmacValid),
		Status:             d.Status,
		ProcessingAttempts: d.ProcessingAttempts,
		LastError:          null.StringFrom(d.LastError),
		ProcessedAt:        null.TimeFrom(d.ProcessedAt),
		ResolvedPaymentId:  nuuid.From(d.ResolvedPaymentId),
		ResolvedIntentId:   nuuid.From(d.ResolvedIntentId),
		MetaCreatedAt:      d.MetaCreatedAt,
		MetaCreatedBy:      d.MetaCreatedBy,
		MetaUpdatedAt:      d.MetaUpdatedAt,
		MetaUpdatedBy:      nuuid.From(d.MetaUpdatedBy),
	}
}

type PspWebhooksBulkUpdateRequest struct {
	Id                 uuid.UUID           `json:"id"`
	PspAccountId       uuid.UUID           `json:"pspAccountId"`
	Psp                string              `json:"psp"`
	PspEventId         string              `json:"pspEventId"`
	PspEventType       string              `json:"pspEventType"`
	ReceivedAt         time.Time           `json:"receivedAt"`
	Headers            json.RawMessage     `json:"headers"`
	RawPayload         json.RawMessage     `json:"rawPayload"`
	HmacValid          bool                `json:"hmacValid"`
	Status             model.WebhookStatus `json:"status" example:"RECEIVED" enums:"RECEIVED,PROCESSING,PROCESSED,FAILED,IGNORED"`
	ProcessingAttempts int                 `json:"processingAttempts"`
	LastError          string              `json:"lastError"`
	ProcessedAt        time.Time           `json:"processedAt"`
	ResolvedPaymentId  uuid.UUID           `json:"resolvedPaymentId"`
	ResolvedIntentId   uuid.UUID           `json:"resolvedIntentId"`
	MetaCreatedAt      time.Time           `json:"metaCreatedAt"`
	MetaCreatedBy      uuid.UUID           `json:"metaCreatedBy"`
	MetaUpdatedAt      time.Time           `json:"metaUpdatedAt"`
	MetaUpdatedBy      uuid.UUID           `json:"metaUpdatedBy"`
}

func (d PspWebhooksBulkUpdateRequest) PrimaryID() PspWebhooksPrimaryID {
	return PspWebhooksPrimaryID{
		Id: d.Id,
	}
}

type PspWebhooksListBulkUpdateRequest []*PspWebhooksBulkUpdateRequest

func (d PspWebhooksListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, pspWebhooks := range d {
		err = validator.Struct(pspWebhooks)
		if err != nil {
			return
		}
	}
	return nil
}

func (d PspWebhooksBulkUpdateRequest) ToModel() model.PspWebhooks {
	return model.PspWebhooks{
		Id:                 d.Id,
		PspAccountId:       d.PspAccountId,
		Psp:                d.Psp,
		PspEventId:         null.StringFrom(d.PspEventId),
		PspEventType:       d.PspEventType,
		ReceivedAt:         d.ReceivedAt,
		Headers:            d.Headers,
		RawPayload:         d.RawPayload,
		HmacValid:          null.BoolFrom(d.HmacValid),
		Status:             d.Status,
		ProcessingAttempts: d.ProcessingAttempts,
		LastError:          null.StringFrom(d.LastError),
		ProcessedAt:        null.TimeFrom(d.ProcessedAt),
		ResolvedPaymentId:  nuuid.From(d.ResolvedPaymentId),
		ResolvedIntentId:   nuuid.From(d.ResolvedIntentId),
		MetaCreatedAt:      d.MetaCreatedAt,
		MetaCreatedBy:      d.MetaCreatedBy,
		MetaUpdatedAt:      d.MetaUpdatedAt,
		MetaUpdatedBy:      nuuid.From(d.MetaUpdatedBy),
	}
}

type PspWebhooksResponse struct {
	Id                 uuid.UUID           `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PspAccountId       uuid.UUID           `json:"pspAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Psp                string              `json:"psp" validate:"required"`
	PspEventId         string              `json:"pspEventId"`
	PspEventType       string              `json:"pspEventType" validate:"required"`
	ReceivedAt         time.Time           `json:"receivedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Headers            json.RawMessage     `json:"headers" swaggertype:"object"`
	RawPayload         json.RawMessage     `json:"rawPayload" validate:"required" swaggertype:"object"`
	HmacValid          bool                `json:"hmacValid" example:"true"`
	Status             model.WebhookStatus `json:"status" validate:"oneof=RECEIVED PROCESSING PROCESSED FAILED IGNORED" enums:"RECEIVED,PROCESSING,PROCESSED,FAILED,IGNORED"`
	ProcessingAttempts int                 `json:"processingAttempts" example:"1"`
	LastError          string              `json:"lastError"`
	ProcessedAt        time.Time           `json:"processedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ResolvedPaymentId  uuid.UUID           `json:"resolvedPaymentId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ResolvedIntentId   uuid.UUID           `json:"resolvedIntentId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaCreatedAt      time.Time           `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy      uuid.UUID           `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt      time.Time           `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy      uuid.UUID           `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewPspWebhooksResponse(pspWebhooks model.PspWebhooks) PspWebhooksResponse {
	return PspWebhooksResponse{
		Id:                 pspWebhooks.Id,
		PspAccountId:       pspWebhooks.PspAccountId,
		Psp:                pspWebhooks.Psp,
		PspEventId:         pspWebhooks.PspEventId.String,
		PspEventType:       pspWebhooks.PspEventType,
		ReceivedAt:         pspWebhooks.ReceivedAt,
		Headers:            pspWebhooks.Headers,
		RawPayload:         pspWebhooks.RawPayload,
		HmacValid:          pspWebhooks.HmacValid.Bool,
		Status:             model.WebhookStatus(pspWebhooks.Status),
		ProcessingAttempts: pspWebhooks.ProcessingAttempts,
		LastError:          pspWebhooks.LastError.String,
		ProcessedAt:        pspWebhooks.ProcessedAt.Time,
		ResolvedPaymentId:  pspWebhooks.ResolvedPaymentId.UUID,
		ResolvedIntentId:   pspWebhooks.ResolvedIntentId.UUID,
		MetaCreatedAt:      pspWebhooks.MetaCreatedAt,
		MetaCreatedBy:      pspWebhooks.MetaCreatedBy,
		MetaUpdatedAt:      pspWebhooks.MetaUpdatedAt,
		MetaUpdatedBy:      pspWebhooks.MetaUpdatedBy.UUID,
	}
}

type PspWebhooksListResponse []*PspWebhooksResponse

func NewPspWebhooksListResponse(pspWebhooksList model.PspWebhooksList) PspWebhooksListResponse {
	dtoPspWebhooksListResponse := PspWebhooksListResponse{}
	for _, pspWebhooks := range pspWebhooksList {
		dtoPspWebhooksResponse := NewPspWebhooksResponse(*pspWebhooks)
		dtoPspWebhooksListResponse = append(dtoPspWebhooksListResponse, &dtoPspWebhooksResponse)
	}
	return dtoPspWebhooksListResponse
}

type PspWebhooksPrimaryIDList []PspWebhooksPrimaryID

func (d PspWebhooksPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, pspWebhooks := range d {
		err = validator.Struct(pspWebhooks)
		if err != nil {
			return
		}
	}
	return nil
}

type PspWebhooksPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *PspWebhooksPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d PspWebhooksPrimaryID) ToModel() model.PspWebhooksPrimaryID {
	return model.PspWebhooksPrimaryID{
		Id: d.Id,
	}
}
