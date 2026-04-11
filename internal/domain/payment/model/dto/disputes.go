package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type DisputesDTOFieldNameType string

type disputesDTOFieldName struct {
	Id                  DisputesDTOFieldNameType
	DisputeCode         DisputesDTOFieldNameType
	PaymentId           DisputesDTOFieldNameType
	PspDisputeId        DisputesDTOFieldNameType
	DisputeType         DisputesDTOFieldNameType
	ReasonCode          DisputesDTOFieldNameType
	ReasonDescription   DisputesDTOFieldNameType
	Amount              DisputesDTOFieldNameType
	Currency            DisputesDTOFieldNameType
	Status              DisputesDTOFieldNameType
	OpenedAt            DisputesDTOFieldNameType
	RespondBy           DisputesDTOFieldNameType
	ResolvedAt          DisputesDTOFieldNameType
	Outcome             DisputesDTOFieldNameType
	EvidenceDueAt       DisputesDTOFieldNameType
	EvidenceSubmittedAt DisputesDTOFieldNameType
	EvidenceFiles       DisputesDTOFieldNameType
	Notes               DisputesDTOFieldNameType
	HandledBy           DisputesDTOFieldNameType
	MetaCreatedAt       DisputesDTOFieldNameType
	MetaCreatedBy       DisputesDTOFieldNameType
	MetaUpdatedAt       DisputesDTOFieldNameType
	MetaUpdatedBy       DisputesDTOFieldNameType
	MetaDeletedAt       DisputesDTOFieldNameType
	MetaDeletedBy       DisputesDTOFieldNameType
}

var DisputesDTOFieldName = disputesDTOFieldName{
	Id:                  "id",
	DisputeCode:         "disputeCode",
	PaymentId:           "paymentId",
	PspDisputeId:        "pspDisputeId",
	DisputeType:         "disputeType",
	ReasonCode:          "reasonCode",
	ReasonDescription:   "reasonDescription",
	Amount:              "amount",
	Currency:            "currency",
	Status:              "status",
	OpenedAt:            "openedAt",
	RespondBy:           "respondBy",
	ResolvedAt:          "resolvedAt",
	Outcome:             "outcome",
	EvidenceDueAt:       "evidenceDueAt",
	EvidenceSubmittedAt: "evidenceSubmittedAt",
	EvidenceFiles:       "evidenceFiles",
	Notes:               "notes",
	HandledBy:           "handledBy",
	MetaCreatedAt:       "metaCreatedAt",
	MetaCreatedBy:       "metaCreatedBy",
	MetaUpdatedAt:       "metaUpdatedAt",
	MetaUpdatedBy:       "metaUpdatedBy",
	MetaDeletedAt:       "metaDeletedAt",
	MetaDeletedBy:       "metaDeletedBy",
}

func NewDisputesListResponseFromFilterResult(result []model.DisputesFilterResult, filter model.Filter) DisputesSelectableListResponse {
	dtoDisputesListResponse := DisputesSelectableListResponse{}
	for _, disputes := range result {
		dtoDisputesResponse := NewDisputesSelectableResponse(disputes.Disputes, filter)
		dtoDisputesListResponse = append(dtoDisputesListResponse, &dtoDisputesResponse)
	}
	return dtoDisputesListResponse
}

func transformDisputesDTOFieldNameFromStr(field string) (dbField model.DisputesDBFieldNameType, found bool) {
	switch field {

	case string(DisputesDTOFieldName.Id):
		return model.DisputesDBFieldName.Id, true

	case string(DisputesDTOFieldName.DisputeCode):
		return model.DisputesDBFieldName.DisputeCode, true

	case string(DisputesDTOFieldName.PaymentId):
		return model.DisputesDBFieldName.PaymentId, true

	case string(DisputesDTOFieldName.PspDisputeId):
		return model.DisputesDBFieldName.PspDisputeId, true

	case string(DisputesDTOFieldName.DisputeType):
		return model.DisputesDBFieldName.DisputeType, true

	case string(DisputesDTOFieldName.ReasonCode):
		return model.DisputesDBFieldName.ReasonCode, true

	case string(DisputesDTOFieldName.ReasonDescription):
		return model.DisputesDBFieldName.ReasonDescription, true

	case string(DisputesDTOFieldName.Amount):
		return model.DisputesDBFieldName.Amount, true

	case string(DisputesDTOFieldName.Currency):
		return model.DisputesDBFieldName.Currency, true

	case string(DisputesDTOFieldName.Status):
		return model.DisputesDBFieldName.Status, true

	case string(DisputesDTOFieldName.OpenedAt):
		return model.DisputesDBFieldName.OpenedAt, true

	case string(DisputesDTOFieldName.RespondBy):
		return model.DisputesDBFieldName.RespondBy, true

	case string(DisputesDTOFieldName.ResolvedAt):
		return model.DisputesDBFieldName.ResolvedAt, true

	case string(DisputesDTOFieldName.Outcome):
		return model.DisputesDBFieldName.Outcome, true

	case string(DisputesDTOFieldName.EvidenceDueAt):
		return model.DisputesDBFieldName.EvidenceDueAt, true

	case string(DisputesDTOFieldName.EvidenceSubmittedAt):
		return model.DisputesDBFieldName.EvidenceSubmittedAt, true

	case string(DisputesDTOFieldName.EvidenceFiles):
		return model.DisputesDBFieldName.EvidenceFiles, true

	case string(DisputesDTOFieldName.Notes):
		return model.DisputesDBFieldName.Notes, true

	case string(DisputesDTOFieldName.HandledBy):
		return model.DisputesDBFieldName.HandledBy, true

	case string(DisputesDTOFieldName.MetaCreatedAt):
		return model.DisputesDBFieldName.MetaCreatedAt, true

	case string(DisputesDTOFieldName.MetaCreatedBy):
		return model.DisputesDBFieldName.MetaCreatedBy, true

	case string(DisputesDTOFieldName.MetaUpdatedAt):
		return model.DisputesDBFieldName.MetaUpdatedAt, true

	case string(DisputesDTOFieldName.MetaUpdatedBy):
		return model.DisputesDBFieldName.MetaUpdatedBy, true

	case string(DisputesDTOFieldName.MetaDeletedAt):
		return model.DisputesDBFieldName.MetaDeletedAt, true

	case string(DisputesDTOFieldName.MetaDeletedBy):
		return model.DisputesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformDisputesFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformDisputesDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformDisputesDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformDisputesDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultDisputesFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(DisputesDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type DisputesSelectableResponse map[string]interface{}
type DisputesSelectableListResponse []*DisputesSelectableResponse

func NewDisputesSelectableResponse(disputes model.Disputes, filter model.Filter) DisputesSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.DisputesDBFieldName.Id),
			string(model.DisputesDBFieldName.DisputeCode),
			string(model.DisputesDBFieldName.PaymentId),
			string(model.DisputesDBFieldName.PspDisputeId),
			string(model.DisputesDBFieldName.DisputeType),
			string(model.DisputesDBFieldName.ReasonCode),
			string(model.DisputesDBFieldName.ReasonDescription),
			string(model.DisputesDBFieldName.Amount),
			string(model.DisputesDBFieldName.Currency),
			string(model.DisputesDBFieldName.Status),
			string(model.DisputesDBFieldName.OpenedAt),
			string(model.DisputesDBFieldName.RespondBy),
			string(model.DisputesDBFieldName.ResolvedAt),
			string(model.DisputesDBFieldName.Outcome),
			string(model.DisputesDBFieldName.EvidenceDueAt),
			string(model.DisputesDBFieldName.EvidenceSubmittedAt),
			string(model.DisputesDBFieldName.EvidenceFiles),
			string(model.DisputesDBFieldName.Notes),
			string(model.DisputesDBFieldName.HandledBy),
			string(model.DisputesDBFieldName.MetaCreatedAt),
			string(model.DisputesDBFieldName.MetaCreatedBy),
			string(model.DisputesDBFieldName.MetaUpdatedAt),
			string(model.DisputesDBFieldName.MetaUpdatedBy),
			string(model.DisputesDBFieldName.MetaDeletedAt),
			string(model.DisputesDBFieldName.MetaDeletedBy),
		)
	}
	disputesSelectableResponse := DisputesSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.DisputesDBFieldName.Id):
			disputesSelectableResponse[string(DisputesDTOFieldName.Id)] = disputes.Id

		case string(model.DisputesDBFieldName.DisputeCode):
			disputesSelectableResponse[string(DisputesDTOFieldName.DisputeCode)] = disputes.DisputeCode

		case string(model.DisputesDBFieldName.PaymentId):
			disputesSelectableResponse[string(DisputesDTOFieldName.PaymentId)] = disputes.PaymentId

		case string(model.DisputesDBFieldName.PspDisputeId):
			disputesSelectableResponse[string(DisputesDTOFieldName.PspDisputeId)] = disputes.PspDisputeId

		case string(model.DisputesDBFieldName.DisputeType):
			disputesSelectableResponse[string(DisputesDTOFieldName.DisputeType)] = disputes.DisputeType

		case string(model.DisputesDBFieldName.ReasonCode):
			disputesSelectableResponse[string(DisputesDTOFieldName.ReasonCode)] = disputes.ReasonCode

		case string(model.DisputesDBFieldName.ReasonDescription):
			disputesSelectableResponse[string(DisputesDTOFieldName.ReasonDescription)] = disputes.ReasonDescription

		case string(model.DisputesDBFieldName.Amount):
			disputesSelectableResponse[string(DisputesDTOFieldName.Amount)] = disputes.Amount

		case string(model.DisputesDBFieldName.Currency):
			disputesSelectableResponse[string(DisputesDTOFieldName.Currency)] = disputes.Currency

		case string(model.DisputesDBFieldName.Status):
			disputesSelectableResponse[string(DisputesDTOFieldName.Status)] = disputes.Status

		case string(model.DisputesDBFieldName.OpenedAt):
			disputesSelectableResponse[string(DisputesDTOFieldName.OpenedAt)] = disputes.OpenedAt

		case string(model.DisputesDBFieldName.RespondBy):
			disputesSelectableResponse[string(DisputesDTOFieldName.RespondBy)] = disputes.RespondBy

		case string(model.DisputesDBFieldName.ResolvedAt):
			disputesSelectableResponse[string(DisputesDTOFieldName.ResolvedAt)] = disputes.ResolvedAt

		case string(model.DisputesDBFieldName.Outcome):
			disputesSelectableResponse[string(DisputesDTOFieldName.Outcome)] = disputes.Outcome

		case string(model.DisputesDBFieldName.EvidenceDueAt):
			disputesSelectableResponse[string(DisputesDTOFieldName.EvidenceDueAt)] = disputes.EvidenceDueAt

		case string(model.DisputesDBFieldName.EvidenceSubmittedAt):
			disputesSelectableResponse[string(DisputesDTOFieldName.EvidenceSubmittedAt)] = disputes.EvidenceSubmittedAt

		case string(model.DisputesDBFieldName.EvidenceFiles):
			disputesSelectableResponse[string(DisputesDTOFieldName.EvidenceFiles)] = disputes.EvidenceFiles

		case string(model.DisputesDBFieldName.Notes):
			disputesSelectableResponse[string(DisputesDTOFieldName.Notes)] = disputes.Notes

		case string(model.DisputesDBFieldName.HandledBy):
			disputesSelectableResponse[string(DisputesDTOFieldName.HandledBy)] = disputes.HandledBy

		case string(model.DisputesDBFieldName.MetaCreatedAt):
			disputesSelectableResponse[string(DisputesDTOFieldName.MetaCreatedAt)] = disputes.MetaCreatedAt

		case string(model.DisputesDBFieldName.MetaCreatedBy):
			disputesSelectableResponse[string(DisputesDTOFieldName.MetaCreatedBy)] = disputes.MetaCreatedBy

		case string(model.DisputesDBFieldName.MetaUpdatedAt):
			disputesSelectableResponse[string(DisputesDTOFieldName.MetaUpdatedAt)] = disputes.MetaUpdatedAt

		case string(model.DisputesDBFieldName.MetaUpdatedBy):
			disputesSelectableResponse[string(DisputesDTOFieldName.MetaUpdatedBy)] = disputes.MetaUpdatedBy

		case string(model.DisputesDBFieldName.MetaDeletedAt):
			disputesSelectableResponse[string(DisputesDTOFieldName.MetaDeletedAt)] = disputes.MetaDeletedAt

		case string(model.DisputesDBFieldName.MetaDeletedBy):
			disputesSelectableResponse[string(DisputesDTOFieldName.MetaDeletedBy)] = disputes.MetaDeletedBy

		}
	}
	return disputesSelectableResponse
}

type DisputesFilterResponse struct {
	Metadata Metadata                       `json:"metadata"`
	Data     DisputesSelectableListResponse `json:"data"`
}

func NewDisputesFilterResponse(result []model.DisputesFilterResult, filter model.Filter) (resp DisputesFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewDisputesListResponseFromFilterResult(result, filter)
	return resp
}

type DisputesCreateRequest struct {
	DisputeCode         string                `json:"disputeCode"`
	PaymentId           uuid.UUID             `json:"paymentId"`
	PspDisputeId        string                `json:"pspDisputeId"`
	DisputeType         string                `json:"disputeType"`
	ReasonCode          string                `json:"reasonCode"`
	ReasonDescription   string                `json:"reasonDescription"`
	Amount              decimal.Decimal       `json:"amount"`
	Currency            model.PaymentCurrency `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Status              model.DisputeStatus   `json:"status" example:"OPEN" enums:"OPEN,EVIDENCE_SUBMITTED,WON,LOST,CLOSED_BY_ISSUER"`
	OpenedAt            time.Time             `json:"openedAt"`
	RespondBy           time.Time             `json:"respondBy"`
	ResolvedAt          time.Time             `json:"resolvedAt"`
	Outcome             string                `json:"outcome"`
	EvidenceDueAt       time.Time             `json:"evidenceDueAt"`
	EvidenceSubmittedAt time.Time             `json:"evidenceSubmittedAt"`
	EvidenceFiles       json.RawMessage       `json:"evidenceFiles"`
	Notes               string                `json:"notes"`
	HandledBy           uuid.UUID             `json:"handledBy"`
	MetaCreatedAt       time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy       uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt       time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy       uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt       time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy       uuid.UUID             `json:"metaDeletedBy"`
}

func (d *DisputesCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *DisputesCreateRequest) ToModel() model.Disputes {
	id, _ := uuid.NewV7()
	return model.Disputes{
		Id:                  id,
		DisputeCode:         d.DisputeCode,
		PaymentId:           d.PaymentId,
		PspDisputeId:        d.PspDisputeId,
		DisputeType:         d.DisputeType,
		ReasonCode:          null.StringFrom(d.ReasonCode),
		ReasonDescription:   null.StringFrom(d.ReasonDescription),
		Amount:              d.Amount,
		Currency:            d.Currency,
		Status:              d.Status,
		OpenedAt:            d.OpenedAt,
		RespondBy:           null.TimeFrom(d.RespondBy),
		ResolvedAt:          null.TimeFrom(d.ResolvedAt),
		Outcome:             null.StringFrom(d.Outcome),
		EvidenceDueAt:       null.TimeFrom(d.EvidenceDueAt),
		EvidenceSubmittedAt: null.TimeFrom(d.EvidenceSubmittedAt),
		EvidenceFiles:       d.EvidenceFiles,
		Notes:               null.StringFrom(d.Notes),
		HandledBy:           nuuid.From(d.HandledBy),
		MetaCreatedAt:       d.MetaCreatedAt,
		MetaCreatedBy:       d.MetaCreatedBy,
		MetaUpdatedAt:       d.MetaUpdatedAt,
		MetaUpdatedBy:       nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:       null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:       nuuid.From(d.MetaDeletedBy),
	}
}

type DisputesListCreateRequest []*DisputesCreateRequest

func (d DisputesListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputes := range d {
		err = validator.Struct(disputes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputesListCreateRequest) ToModelList() []model.Disputes {
	out := make([]model.Disputes, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type DisputesUpdateRequest struct {
	DisputeCode         string                `json:"disputeCode"`
	PaymentId           uuid.UUID             `json:"paymentId"`
	PspDisputeId        string                `json:"pspDisputeId"`
	DisputeType         string                `json:"disputeType"`
	ReasonCode          string                `json:"reasonCode"`
	ReasonDescription   string                `json:"reasonDescription"`
	Amount              decimal.Decimal       `json:"amount"`
	Currency            model.PaymentCurrency `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Status              model.DisputeStatus   `json:"status" example:"OPEN" enums:"OPEN,EVIDENCE_SUBMITTED,WON,LOST,CLOSED_BY_ISSUER"`
	OpenedAt            time.Time             `json:"openedAt"`
	RespondBy           time.Time             `json:"respondBy"`
	ResolvedAt          time.Time             `json:"resolvedAt"`
	Outcome             string                `json:"outcome"`
	EvidenceDueAt       time.Time             `json:"evidenceDueAt"`
	EvidenceSubmittedAt time.Time             `json:"evidenceSubmittedAt"`
	EvidenceFiles       json.RawMessage       `json:"evidenceFiles"`
	Notes               string                `json:"notes"`
	HandledBy           uuid.UUID             `json:"handledBy"`
	MetaCreatedAt       time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy       uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt       time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy       uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt       time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy       uuid.UUID             `json:"metaDeletedBy"`
}

func (d *DisputesUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d DisputesUpdateRequest) ToModel() model.Disputes {
	return model.Disputes{
		DisputeCode:         d.DisputeCode,
		PaymentId:           d.PaymentId,
		PspDisputeId:        d.PspDisputeId,
		DisputeType:         d.DisputeType,
		ReasonCode:          null.StringFrom(d.ReasonCode),
		ReasonDescription:   null.StringFrom(d.ReasonDescription),
		Amount:              d.Amount,
		Currency:            d.Currency,
		Status:              d.Status,
		OpenedAt:            d.OpenedAt,
		RespondBy:           null.TimeFrom(d.RespondBy),
		ResolvedAt:          null.TimeFrom(d.ResolvedAt),
		Outcome:             null.StringFrom(d.Outcome),
		EvidenceDueAt:       null.TimeFrom(d.EvidenceDueAt),
		EvidenceSubmittedAt: null.TimeFrom(d.EvidenceSubmittedAt),
		EvidenceFiles:       d.EvidenceFiles,
		Notes:               null.StringFrom(d.Notes),
		HandledBy:           nuuid.From(d.HandledBy),
		MetaCreatedAt:       d.MetaCreatedAt,
		MetaCreatedBy:       d.MetaCreatedBy,
		MetaUpdatedAt:       d.MetaUpdatedAt,
		MetaUpdatedBy:       nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:       null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:       nuuid.From(d.MetaDeletedBy),
	}
}

type DisputesBulkUpdateRequest struct {
	Id                  uuid.UUID             `json:"id"`
	DisputeCode         string                `json:"disputeCode"`
	PaymentId           uuid.UUID             `json:"paymentId"`
	PspDisputeId        string                `json:"pspDisputeId"`
	DisputeType         string                `json:"disputeType"`
	ReasonCode          string                `json:"reasonCode"`
	ReasonDescription   string                `json:"reasonDescription"`
	Amount              decimal.Decimal       `json:"amount"`
	Currency            model.PaymentCurrency `json:"currency" example:"IDR" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Status              model.DisputeStatus   `json:"status" example:"OPEN" enums:"OPEN,EVIDENCE_SUBMITTED,WON,LOST,CLOSED_BY_ISSUER"`
	OpenedAt            time.Time             `json:"openedAt"`
	RespondBy           time.Time             `json:"respondBy"`
	ResolvedAt          time.Time             `json:"resolvedAt"`
	Outcome             string                `json:"outcome"`
	EvidenceDueAt       time.Time             `json:"evidenceDueAt"`
	EvidenceSubmittedAt time.Time             `json:"evidenceSubmittedAt"`
	EvidenceFiles       json.RawMessage       `json:"evidenceFiles"`
	Notes               string                `json:"notes"`
	HandledBy           uuid.UUID             `json:"handledBy"`
	MetaCreatedAt       time.Time             `json:"metaCreatedAt"`
	MetaCreatedBy       uuid.UUID             `json:"metaCreatedBy"`
	MetaUpdatedAt       time.Time             `json:"metaUpdatedAt"`
	MetaUpdatedBy       uuid.UUID             `json:"metaUpdatedBy"`
	MetaDeletedAt       time.Time             `json:"metaDeletedAt"`
	MetaDeletedBy       uuid.UUID             `json:"metaDeletedBy"`
}

func (d DisputesBulkUpdateRequest) PrimaryID() DisputesPrimaryID {
	return DisputesPrimaryID{
		Id: d.Id,
	}
}

type DisputesListBulkUpdateRequest []*DisputesBulkUpdateRequest

func (d DisputesListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputes := range d {
		err = validator.Struct(disputes)
		if err != nil {
			return
		}
	}
	return nil
}

func (d DisputesBulkUpdateRequest) ToModel() model.Disputes {
	return model.Disputes{
		Id:                  d.Id,
		DisputeCode:         d.DisputeCode,
		PaymentId:           d.PaymentId,
		PspDisputeId:        d.PspDisputeId,
		DisputeType:         d.DisputeType,
		ReasonCode:          null.StringFrom(d.ReasonCode),
		ReasonDescription:   null.StringFrom(d.ReasonDescription),
		Amount:              d.Amount,
		Currency:            d.Currency,
		Status:              d.Status,
		OpenedAt:            d.OpenedAt,
		RespondBy:           null.TimeFrom(d.RespondBy),
		ResolvedAt:          null.TimeFrom(d.ResolvedAt),
		Outcome:             null.StringFrom(d.Outcome),
		EvidenceDueAt:       null.TimeFrom(d.EvidenceDueAt),
		EvidenceSubmittedAt: null.TimeFrom(d.EvidenceSubmittedAt),
		EvidenceFiles:       d.EvidenceFiles,
		Notes:               null.StringFrom(d.Notes),
		HandledBy:           nuuid.From(d.HandledBy),
		MetaCreatedAt:       d.MetaCreatedAt,
		MetaCreatedBy:       d.MetaCreatedBy,
		MetaUpdatedAt:       d.MetaUpdatedAt,
		MetaUpdatedBy:       nuuid.From(d.MetaUpdatedBy),
		MetaDeletedAt:       null.TimeFrom(d.MetaDeletedAt),
		MetaDeletedBy:       nuuid.From(d.MetaDeletedBy),
	}
}

type DisputesResponse struct {
	Id                  uuid.UUID             `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	DisputeCode         string                `json:"disputeCode" validate:"required"`
	PaymentId           uuid.UUID             `json:"paymentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PspDisputeId        string                `json:"pspDisputeId" validate:"required"`
	DisputeType         string                `json:"disputeType" validate:"required"`
	ReasonCode          string                `json:"reasonCode"`
	ReasonDescription   string                `json:"reasonDescription"`
	Amount              decimal.Decimal       `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	Currency            model.PaymentCurrency `json:"currency" validate:"required,oneof=IDR USD SGD MYR PHP THB AED EUR GBP JPY" enums:"IDR,USD,SGD,MYR,PHP,THB,AED,EUR,GBP,JPY"`
	Status              model.DisputeStatus   `json:"status" validate:"oneof=OPEN EVIDENCE_SUBMITTED WON LOST CLOSED_BY_ISSUER" enums:"OPEN,EVIDENCE_SUBMITTED,WON,LOST,CLOSED_BY_ISSUER"`
	OpenedAt            time.Time             `json:"openedAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	RespondBy           time.Time             `json:"respondBy" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ResolvedAt          time.Time             `json:"resolvedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	Outcome             string                `json:"outcome"`
	EvidenceDueAt       time.Time             `json:"evidenceDueAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	EvidenceSubmittedAt time.Time             `json:"evidenceSubmittedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	EvidenceFiles       json.RawMessage       `json:"evidenceFiles" swaggertype:"object"`
	Notes               string                `json:"notes"`
	HandledBy           uuid.UUID             `json:"handledBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaCreatedAt       time.Time             `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy       uuid.UUID             `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt       time.Time             `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy       uuid.UUID             `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaDeletedAt       time.Time             `json:"metaDeletedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaDeletedBy       uuid.UUID             `json:"metaDeletedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewDisputesResponse(disputes model.Disputes) DisputesResponse {
	return DisputesResponse{
		Id:                  disputes.Id,
		DisputeCode:         disputes.DisputeCode,
		PaymentId:           disputes.PaymentId,
		PspDisputeId:        disputes.PspDisputeId,
		DisputeType:         disputes.DisputeType,
		ReasonCode:          disputes.ReasonCode.String,
		ReasonDescription:   disputes.ReasonDescription.String,
		Amount:              disputes.Amount,
		Currency:            model.PaymentCurrency(disputes.Currency),
		Status:              model.DisputeStatus(disputes.Status),
		OpenedAt:            disputes.OpenedAt,
		RespondBy:           disputes.RespondBy.Time,
		ResolvedAt:          disputes.ResolvedAt.Time,
		Outcome:             disputes.Outcome.String,
		EvidenceDueAt:       disputes.EvidenceDueAt.Time,
		EvidenceSubmittedAt: disputes.EvidenceSubmittedAt.Time,
		EvidenceFiles:       disputes.EvidenceFiles,
		Notes:               disputes.Notes.String,
		HandledBy:           disputes.HandledBy.UUID,
		MetaCreatedAt:       disputes.MetaCreatedAt,
		MetaCreatedBy:       disputes.MetaCreatedBy,
		MetaUpdatedAt:       disputes.MetaUpdatedAt,
		MetaUpdatedBy:       disputes.MetaUpdatedBy.UUID,
		MetaDeletedAt:       disputes.MetaDeletedAt.Time,
		MetaDeletedBy:       disputes.MetaDeletedBy.UUID,
	}
}

type DisputesListResponse []*DisputesResponse

func NewDisputesListResponse(disputesList model.DisputesList) DisputesListResponse {
	dtoDisputesListResponse := DisputesListResponse{}
	for _, disputes := range disputesList {
		dtoDisputesResponse := NewDisputesResponse(*disputes)
		dtoDisputesListResponse = append(dtoDisputesListResponse, &dtoDisputesResponse)
	}
	return dtoDisputesListResponse
}

type DisputesPrimaryIDList []DisputesPrimaryID

func (d DisputesPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, disputes := range d {
		err = validator.Struct(disputes)
		if err != nil {
			return
		}
	}
	return nil
}

type DisputesPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *DisputesPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d DisputesPrimaryID) ToModel() model.DisputesPrimaryID {
	return model.DisputesPrimaryID{
		Id: d.Id,
	}
}
