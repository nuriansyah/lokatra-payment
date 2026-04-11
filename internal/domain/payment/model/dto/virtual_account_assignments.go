package dto

import (
	"fmt"
	"math"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type VirtualAccountAssignmentsDTOFieldNameType string

type virtualAccountAssignmentsDTOFieldName struct {
	Id       VirtualAccountAssignmentsDTOFieldNameType
	IntentId VirtualAccountAssignmentsDTOFieldNameType

	BankCode         VirtualAccountAssignmentsDTOFieldNameType
	VaNumber         VirtualAccountAssignmentsDTOFieldNameType
	VaNumberMasked   VirtualAccountAssignmentsDTOFieldNameType
	ExpiresAt        VirtualAccountAssignmentsDTOFieldNameType
	IsReusable       VirtualAccountAssignmentsDTOFieldNameType
	PaidAt           VirtualAccountAssignmentsDTOFieldNameType
	PspTransactionId VirtualAccountAssignmentsDTOFieldNameType
	MetaCreatedAt    VirtualAccountAssignmentsDTOFieldNameType
	MetaCreatedBy    VirtualAccountAssignmentsDTOFieldNameType
	MetaUpdatedAt    VirtualAccountAssignmentsDTOFieldNameType
	MetaUpdatedBy    VirtualAccountAssignmentsDTOFieldNameType
}

var VirtualAccountAssignmentsDTOFieldName = virtualAccountAssignmentsDTOFieldName{
	Id:       "id",
	IntentId: "intentId",

	BankCode:         "bankCode",
	VaNumber:         "vaNumber",
	VaNumberMasked:   "vaNumberMasked",
	ExpiresAt:        "expiresAt",
	IsReusable:       "isReusable",
	PaidAt:           "paidAt",
	PspTransactionId: "pspTransactionId",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
}

func NewVirtualAccountAssignmentsListResponseFromFilterResult(result []model.VirtualAccountAssignmentsFilterResult, filter model.Filter) VirtualAccountAssignmentsSelectableListResponse {
	dtoVirtualAccountAssignmentsListResponse := VirtualAccountAssignmentsSelectableListResponse{}
	for _, virtualAccountAssignments := range result {
		dtoVirtualAccountAssignmentsResponse := NewVirtualAccountAssignmentsSelectableResponse(virtualAccountAssignments.VirtualAccountAssignments, filter)
		dtoVirtualAccountAssignmentsListResponse = append(dtoVirtualAccountAssignmentsListResponse, &dtoVirtualAccountAssignmentsResponse)
	}
	return dtoVirtualAccountAssignmentsListResponse
}

func transformVirtualAccountAssignmentsDTOFieldNameFromStr(field string) (dbField model.VirtualAccountAssignmentsDBFieldNameType, found bool) {
	switch field {

	case string(VirtualAccountAssignmentsDTOFieldName.Id):
		return model.VirtualAccountAssignmentsDBFieldName.Id, true

	case string(VirtualAccountAssignmentsDTOFieldName.IntentId):
		return model.VirtualAccountAssignmentsDBFieldName.IntentId, true

	case string(VirtualAccountAssignmentsDTOFieldName.BankCode):
		return model.VirtualAccountAssignmentsDBFieldName.BankCode, true

	case string(VirtualAccountAssignmentsDTOFieldName.VaNumber):
		return model.VirtualAccountAssignmentsDBFieldName.VaNumber, true

	case string(VirtualAccountAssignmentsDTOFieldName.VaNumberMasked):
		return model.VirtualAccountAssignmentsDBFieldName.VaNumberMasked, true

	case string(VirtualAccountAssignmentsDTOFieldName.ExpiresAt):
		return model.VirtualAccountAssignmentsDBFieldName.ExpiresAt, true

	case string(VirtualAccountAssignmentsDTOFieldName.IsReusable):
		return model.VirtualAccountAssignmentsDBFieldName.IsReusable, true

	case string(VirtualAccountAssignmentsDTOFieldName.PaidAt):
		return model.VirtualAccountAssignmentsDBFieldName.PaidAt, true

	case string(VirtualAccountAssignmentsDTOFieldName.PspTransactionId):
		return model.VirtualAccountAssignmentsDBFieldName.PspTransactionId, true

	case string(VirtualAccountAssignmentsDTOFieldName.MetaCreatedAt):
		return model.VirtualAccountAssignmentsDBFieldName.MetaCreatedAt, true

	case string(VirtualAccountAssignmentsDTOFieldName.MetaCreatedBy):
		return model.VirtualAccountAssignmentsDBFieldName.MetaCreatedBy, true

	case string(VirtualAccountAssignmentsDTOFieldName.MetaUpdatedAt):
		return model.VirtualAccountAssignmentsDBFieldName.MetaUpdatedAt, true

	case string(VirtualAccountAssignmentsDTOFieldName.MetaUpdatedBy):
		return model.VirtualAccountAssignmentsDBFieldName.MetaUpdatedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformVirtualAccountAssignmentsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformVirtualAccountAssignmentsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformVirtualAccountAssignmentsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformVirtualAccountAssignmentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultVirtualAccountAssignmentsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(VirtualAccountAssignmentsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type VirtualAccountAssignmentsSelectableResponse map[string]interface{}
type VirtualAccountAssignmentsSelectableListResponse []*VirtualAccountAssignmentsSelectableResponse

func NewVirtualAccountAssignmentsSelectableResponse(virtualAccountAssignments model.VirtualAccountAssignments, filter model.Filter) VirtualAccountAssignmentsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.VirtualAccountAssignmentsDBFieldName.Id),
			string(model.VirtualAccountAssignmentsDBFieldName.IntentId),
			string(model.VirtualAccountAssignmentsDBFieldName.BankCode),
			string(model.VirtualAccountAssignmentsDBFieldName.VaNumber),
			string(model.VirtualAccountAssignmentsDBFieldName.VaNumberMasked),
			string(model.VirtualAccountAssignmentsDBFieldName.ExpiresAt),
			string(model.VirtualAccountAssignmentsDBFieldName.IsReusable),
			string(model.VirtualAccountAssignmentsDBFieldName.PaidAt),
			string(model.VirtualAccountAssignmentsDBFieldName.PspTransactionId),
			string(model.VirtualAccountAssignmentsDBFieldName.MetaCreatedAt),
			string(model.VirtualAccountAssignmentsDBFieldName.MetaCreatedBy),
			string(model.VirtualAccountAssignmentsDBFieldName.MetaUpdatedAt),
			string(model.VirtualAccountAssignmentsDBFieldName.MetaUpdatedBy),
		)
	}
	virtualAccountAssignmentsSelectableResponse := VirtualAccountAssignmentsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.VirtualAccountAssignmentsDBFieldName.Id):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.Id)] = virtualAccountAssignments.Id

		case string(model.VirtualAccountAssignmentsDBFieldName.IntentId):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.IntentId)] = virtualAccountAssignments.IntentId

		case string(model.VirtualAccountAssignmentsDBFieldName.BankCode):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.BankCode)] = virtualAccountAssignments.BankCode

		case string(model.VirtualAccountAssignmentsDBFieldName.VaNumber):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.VaNumber)] = virtualAccountAssignments.VaNumber

		case string(model.VirtualAccountAssignmentsDBFieldName.VaNumberMasked):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.VaNumberMasked)] = virtualAccountAssignments.VaNumberMasked

		case string(model.VirtualAccountAssignmentsDBFieldName.ExpiresAt):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.ExpiresAt)] = virtualAccountAssignments.ExpiresAt

		case string(model.VirtualAccountAssignmentsDBFieldName.IsReusable):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.IsReusable)] = virtualAccountAssignments.IsReusable

		case string(model.VirtualAccountAssignmentsDBFieldName.PaidAt):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.PaidAt)] = virtualAccountAssignments.PaidAt

		case string(model.VirtualAccountAssignmentsDBFieldName.PspTransactionId):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.PspTransactionId)] = virtualAccountAssignments.PspTransactionId

		case string(model.VirtualAccountAssignmentsDBFieldName.MetaCreatedAt):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.MetaCreatedAt)] = virtualAccountAssignments.MetaCreatedAt

		case string(model.VirtualAccountAssignmentsDBFieldName.MetaCreatedBy):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.MetaCreatedBy)] = virtualAccountAssignments.MetaCreatedBy

		case string(model.VirtualAccountAssignmentsDBFieldName.MetaUpdatedAt):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.MetaUpdatedAt)] = virtualAccountAssignments.MetaUpdatedAt

		case string(model.VirtualAccountAssignmentsDBFieldName.MetaUpdatedBy):
			virtualAccountAssignmentsSelectableResponse[string(VirtualAccountAssignmentsDTOFieldName.MetaUpdatedBy)] = virtualAccountAssignments.MetaUpdatedBy

		}
	}
	return virtualAccountAssignmentsSelectableResponse
}

type VirtualAccountAssignmentsFilterResponse struct {
	Metadata Metadata                                        `json:"metadata"`
	Data     VirtualAccountAssignmentsSelectableListResponse `json:"data"`
}

func NewVirtualAccountAssignmentsFilterResponse(result []model.VirtualAccountAssignmentsFilterResult, filter model.Filter) (resp VirtualAccountAssignmentsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewVirtualAccountAssignmentsListResponseFromFilterResult(result, filter)
	return resp
}

type VirtualAccountAssignmentsCreateRequest struct {
	IntentId         uuid.UUID `json:"intentId"`
	BankCode         string    `json:"bankCode"`
	VaNumber         string    `json:"vaNumber"`
	VaNumberMasked   string    `json:"vaNumberMasked"`
	ExpiresAt        time.Time `json:"expiresAt"`
	IsReusable       bool      `json:"isReusable"`
	PaidAt           time.Time `json:"paidAt"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy"`
}

func (d *VirtualAccountAssignmentsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *VirtualAccountAssignmentsCreateRequest) ToModel() model.VirtualAccountAssignments {
	id, _ := uuid.NewV7()
	return model.VirtualAccountAssignments{
		Id:               id,
		IntentId:         d.IntentId,
		BankCode:         d.BankCode,
		VaNumber:         d.VaNumber,
		VaNumberMasked:   null.StringFrom(d.VaNumberMasked),
		ExpiresAt:        d.ExpiresAt,
		IsReusable:       d.IsReusable,
		PaidAt:           null.TimeFrom(d.PaidAt),
		PspTransactionId: null.StringFrom(d.PspTransactionId),
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
	}
}

type VirtualAccountAssignmentsListCreateRequest []*VirtualAccountAssignmentsCreateRequest

func (d VirtualAccountAssignmentsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, virtualAccountAssignments := range d {
		err = validator.Struct(virtualAccountAssignments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d VirtualAccountAssignmentsListCreateRequest) ToModelList() []model.VirtualAccountAssignments {
	out := make([]model.VirtualAccountAssignments, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type VirtualAccountAssignmentsUpdateRequest struct {
	IntentId         uuid.UUID `json:"intentId"`
	BankCode         string    `json:"bankCode"`
	VaNumber         string    `json:"vaNumber"`
	VaNumberMasked   string    `json:"vaNumberMasked"`
	ExpiresAt        time.Time `json:"expiresAt"`
	IsReusable       bool      `json:"isReusable"`
	PaidAt           time.Time `json:"paidAt"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy"`
}

func (d *VirtualAccountAssignmentsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d VirtualAccountAssignmentsUpdateRequest) ToModel() model.VirtualAccountAssignments {
	return model.VirtualAccountAssignments{
		IntentId:         d.IntentId,
		BankCode:         d.BankCode,
		VaNumber:         d.VaNumber,
		VaNumberMasked:   null.StringFrom(d.VaNumberMasked),
		ExpiresAt:        d.ExpiresAt,
		IsReusable:       d.IsReusable,
		PaidAt:           null.TimeFrom(d.PaidAt),
		PspTransactionId: null.StringFrom(d.PspTransactionId),
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
	}
}

type VirtualAccountAssignmentsBulkUpdateRequest struct {
	Id               uuid.UUID `json:"id"`
	IntentId         uuid.UUID `json:"intentId"`
	BankCode         string    `json:"bankCode"`
	VaNumber         string    `json:"vaNumber"`
	VaNumberMasked   string    `json:"vaNumberMasked"`
	ExpiresAt        time.Time `json:"expiresAt"`
	IsReusable       bool      `json:"isReusable"`
	PaidAt           time.Time `json:"paidAt"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy"`
}

func (d VirtualAccountAssignmentsBulkUpdateRequest) PrimaryID() VirtualAccountAssignmentsPrimaryID {
	return VirtualAccountAssignmentsPrimaryID{
		Id: d.Id,
	}
}

type VirtualAccountAssignmentsListBulkUpdateRequest []*VirtualAccountAssignmentsBulkUpdateRequest

func (d VirtualAccountAssignmentsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, virtualAccountAssignments := range d {
		err = validator.Struct(virtualAccountAssignments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d VirtualAccountAssignmentsBulkUpdateRequest) ToModel() model.VirtualAccountAssignments {
	return model.VirtualAccountAssignments{
		Id:               d.Id,
		IntentId:         d.IntentId,
		BankCode:         d.BankCode,
		VaNumber:         d.VaNumber,
		VaNumberMasked:   null.StringFrom(d.VaNumberMasked),
		ExpiresAt:        d.ExpiresAt,
		IsReusable:       d.IsReusable,
		PaidAt:           null.TimeFrom(d.PaidAt),
		PspTransactionId: null.StringFrom(d.PspTransactionId),
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
	}
}

type VirtualAccountAssignmentsResponse struct {
	Id               uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IntentId         uuid.UUID `json:"intentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	BankCode         string    `json:"bankCode" validate:"required"`
	VaNumber         string    `json:"vaNumber" validate:"required"`
	VaNumberMasked   string    `json:"vaNumberMasked"`
	ExpiresAt        time.Time `json:"expiresAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	IsReusable       bool      `json:"isReusable" example:"true"`
	PaidAt           time.Time `json:"paidAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewVirtualAccountAssignmentsResponse(virtualAccountAssignments model.VirtualAccountAssignments) VirtualAccountAssignmentsResponse {
	return VirtualAccountAssignmentsResponse{
		Id:               virtualAccountAssignments.Id,
		IntentId:         virtualAccountAssignments.IntentId,
		BankCode:         virtualAccountAssignments.BankCode,
		VaNumber:         virtualAccountAssignments.VaNumber,
		VaNumberMasked:   virtualAccountAssignments.VaNumberMasked.String,
		ExpiresAt:        virtualAccountAssignments.ExpiresAt,
		IsReusable:       virtualAccountAssignments.IsReusable,
		PaidAt:           virtualAccountAssignments.PaidAt.Time,
		PspTransactionId: virtualAccountAssignments.PspTransactionId.String,
		MetaCreatedAt:    virtualAccountAssignments.MetaCreatedAt,
		MetaCreatedBy:    virtualAccountAssignments.MetaCreatedBy,
		MetaUpdatedAt:    virtualAccountAssignments.MetaUpdatedAt,
		MetaUpdatedBy:    virtualAccountAssignments.MetaUpdatedBy.UUID,
	}
}

type VirtualAccountAssignmentsListResponse []*VirtualAccountAssignmentsResponse

func NewVirtualAccountAssignmentsListResponse(virtualAccountAssignmentsList model.VirtualAccountAssignmentsList) VirtualAccountAssignmentsListResponse {
	dtoVirtualAccountAssignmentsListResponse := VirtualAccountAssignmentsListResponse{}
	for _, virtualAccountAssignments := range virtualAccountAssignmentsList {
		dtoVirtualAccountAssignmentsResponse := NewVirtualAccountAssignmentsResponse(*virtualAccountAssignments)
		dtoVirtualAccountAssignmentsListResponse = append(dtoVirtualAccountAssignmentsListResponse, &dtoVirtualAccountAssignmentsResponse)
	}
	return dtoVirtualAccountAssignmentsListResponse
}

type VirtualAccountAssignmentsPrimaryIDList []VirtualAccountAssignmentsPrimaryID

func (d VirtualAccountAssignmentsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, virtualAccountAssignments := range d {
		err = validator.Struct(virtualAccountAssignments)
		if err != nil {
			return
		}
	}
	return nil
}

type VirtualAccountAssignmentsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *VirtualAccountAssignmentsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d VirtualAccountAssignmentsPrimaryID) ToModel() model.VirtualAccountAssignmentsPrimaryID {
	return model.VirtualAccountAssignmentsPrimaryID{
		Id: d.Id,
	}
}
