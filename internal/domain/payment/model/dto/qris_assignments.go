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

type QrisAssignmentsDTOFieldNameType string

type qrisAssignmentsDTOFieldName struct {
	Id               QrisAssignmentsDTOFieldNameType
	IntentId         QrisAssignmentsDTOFieldNameType
	QrString         QrisAssignmentsDTOFieldNameType
	QrUrl            QrisAssignmentsDTOFieldNameType
	ExpiresAt        QrisAssignmentsDTOFieldNameType
	PaidAt           QrisAssignmentsDTOFieldNameType
	PspTransactionId QrisAssignmentsDTOFieldNameType
	MetaCreatedAt    QrisAssignmentsDTOFieldNameType
	MetaCreatedBy    QrisAssignmentsDTOFieldNameType
	MetaUpdatedAt    QrisAssignmentsDTOFieldNameType
	MetaUpdatedBy    QrisAssignmentsDTOFieldNameType
}

var QrisAssignmentsDTOFieldName = qrisAssignmentsDTOFieldName{
	Id:               "id",
	IntentId:         "intentId",
	QrString:         "qrString",
	QrUrl:            "qrUrl",
	ExpiresAt:        "expiresAt",
	PaidAt:           "paidAt",
	PspTransactionId: "pspTransactionId",
	MetaCreatedAt:    "metaCreatedAt",
	MetaCreatedBy:    "metaCreatedBy",
	MetaUpdatedAt:    "metaUpdatedAt",
	MetaUpdatedBy:    "metaUpdatedBy",
}

func NewQrisAssignmentsListResponseFromFilterResult(result []model.QrisAssignmentsFilterResult, filter model.Filter) QrisAssignmentsSelectableListResponse {
	dtoQrisAssignmentsListResponse := QrisAssignmentsSelectableListResponse{}
	for _, qrisAssignments := range result {
		dtoQrisAssignmentsResponse := NewQrisAssignmentsSelectableResponse(qrisAssignments.QrisAssignments, filter)
		dtoQrisAssignmentsListResponse = append(dtoQrisAssignmentsListResponse, &dtoQrisAssignmentsResponse)
	}
	return dtoQrisAssignmentsListResponse
}

func transformQrisAssignmentsDTOFieldNameFromStr(field string) (dbField model.QrisAssignmentsDBFieldNameType, found bool) {
	switch field {

	case string(QrisAssignmentsDTOFieldName.Id):
		return model.QrisAssignmentsDBFieldName.Id, true

	case string(QrisAssignmentsDTOFieldName.IntentId):
		return model.QrisAssignmentsDBFieldName.IntentId, true

	case string(QrisAssignmentsDTOFieldName.QrString):
		return model.QrisAssignmentsDBFieldName.QrString, true

	case string(QrisAssignmentsDTOFieldName.QrUrl):
		return model.QrisAssignmentsDBFieldName.QrUrl, true

	case string(QrisAssignmentsDTOFieldName.ExpiresAt):
		return model.QrisAssignmentsDBFieldName.ExpiresAt, true

	case string(QrisAssignmentsDTOFieldName.PaidAt):
		return model.QrisAssignmentsDBFieldName.PaidAt, true

	case string(QrisAssignmentsDTOFieldName.PspTransactionId):
		return model.QrisAssignmentsDBFieldName.PspTransactionId, true

	case string(QrisAssignmentsDTOFieldName.MetaCreatedAt):
		return model.QrisAssignmentsDBFieldName.MetaCreatedAt, true

	case string(QrisAssignmentsDTOFieldName.MetaCreatedBy):
		return model.QrisAssignmentsDBFieldName.MetaCreatedBy, true

	case string(QrisAssignmentsDTOFieldName.MetaUpdatedAt):
		return model.QrisAssignmentsDBFieldName.MetaUpdatedAt, true

	case string(QrisAssignmentsDTOFieldName.MetaUpdatedBy):
		return model.QrisAssignmentsDBFieldName.MetaUpdatedBy, true

	}
	return "unknown", false
}

func ValidateAndTransformQrisAssignmentsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		dbField, exist := transformQrisAssignmentsDTOFieldNameFromStr(selectField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", selectField))
			return
		}
		filter.SelectFields[index] = string(dbField)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformQrisAssignmentsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = string(dbField)
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformQrisAssignmentsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = string(dbField)
	}
	return
}

func SetDefaultQrisAssignmentsFilter(filter *model.Filter) {
	if filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = 10
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(QrisAssignmentsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type QrisAssignmentsSelectableResponse map[string]interface{}
type QrisAssignmentsSelectableListResponse []*QrisAssignmentsSelectableResponse

func NewQrisAssignmentsSelectableResponse(qrisAssignments model.QrisAssignments, filter model.Filter) QrisAssignmentsSelectableResponse {
	// selected fields has been transformed to be db field name
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.QrisAssignmentsDBFieldName.Id),
			string(model.QrisAssignmentsDBFieldName.IntentId),
			string(model.QrisAssignmentsDBFieldName.QrString),
			string(model.QrisAssignmentsDBFieldName.QrUrl),
			string(model.QrisAssignmentsDBFieldName.ExpiresAt),
			string(model.QrisAssignmentsDBFieldName.PaidAt),
			string(model.QrisAssignmentsDBFieldName.PspTransactionId),
			string(model.QrisAssignmentsDBFieldName.MetaCreatedAt),
			string(model.QrisAssignmentsDBFieldName.MetaCreatedBy),
			string(model.QrisAssignmentsDBFieldName.MetaUpdatedAt),
			string(model.QrisAssignmentsDBFieldName.MetaUpdatedBy),
		)
	}
	qrisAssignmentsSelectableResponse := QrisAssignmentsSelectableResponse{}
	for _, selectField := range selectFields {
		switch selectField {

		case string(model.QrisAssignmentsDBFieldName.Id):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.Id)] = qrisAssignments.Id

		case string(model.QrisAssignmentsDBFieldName.IntentId):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.IntentId)] = qrisAssignments.IntentId

		case string(model.QrisAssignmentsDBFieldName.QrString):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.QrString)] = qrisAssignments.QrString

		case string(model.QrisAssignmentsDBFieldName.QrUrl):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.QrUrl)] = qrisAssignments.QrUrl

		case string(model.QrisAssignmentsDBFieldName.ExpiresAt):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.ExpiresAt)] = qrisAssignments.ExpiresAt

		case string(model.QrisAssignmentsDBFieldName.PaidAt):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.PaidAt)] = qrisAssignments.PaidAt

		case string(model.QrisAssignmentsDBFieldName.PspTransactionId):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.PspTransactionId)] = qrisAssignments.PspTransactionId

		case string(model.QrisAssignmentsDBFieldName.MetaCreatedAt):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.MetaCreatedAt)] = qrisAssignments.MetaCreatedAt

		case string(model.QrisAssignmentsDBFieldName.MetaCreatedBy):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.MetaCreatedBy)] = qrisAssignments.MetaCreatedBy

		case string(model.QrisAssignmentsDBFieldName.MetaUpdatedAt):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.MetaUpdatedAt)] = qrisAssignments.MetaUpdatedAt

		case string(model.QrisAssignmentsDBFieldName.MetaUpdatedBy):
			qrisAssignmentsSelectableResponse[string(QrisAssignmentsDTOFieldName.MetaUpdatedBy)] = qrisAssignments.MetaUpdatedBy

		}
	}
	return qrisAssignmentsSelectableResponse
}

type QrisAssignmentsFilterResponse struct {
	Metadata Metadata                              `json:"metadata"`
	Data     QrisAssignmentsSelectableListResponse `json:"data"`
}

func NewQrisAssignmentsFilterResponse(result []model.QrisAssignmentsFilterResult, filter model.Filter) (resp QrisAssignmentsFilterResponse) {
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	if len(result) > 0 {
		resp.Metadata.TotalData = result[0].FilterCount
		resp.Metadata.TotalPage = int(math.Ceil(float64(result[0].FilterCount) / float64(filter.Pagination.PageSize)))
	}
	resp.Data = NewQrisAssignmentsListResponseFromFilterResult(result, filter)
	return resp
}

type QrisAssignmentsCreateRequest struct {
	IntentId         uuid.UUID `json:"intentId"`
	QrString         string    `json:"qrString"`
	QrUrl            string    `json:"qrUrl"`
	ExpiresAt        time.Time `json:"expiresAt"`
	PaidAt           time.Time `json:"paidAt"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy"`
}

func (d *QrisAssignmentsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *QrisAssignmentsCreateRequest) ToModel() model.QrisAssignments {
	id, _ := uuid.NewV7()
	return model.QrisAssignments{
		Id:               id,
		IntentId:         d.IntentId,
		QrString:         d.QrString,
		QrUrl:            null.StringFrom(d.QrUrl),
		ExpiresAt:        d.ExpiresAt,
		PaidAt:           null.TimeFrom(d.PaidAt),
		PspTransactionId: null.StringFrom(d.PspTransactionId),
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
	}
}

type QrisAssignmentsListCreateRequest []*QrisAssignmentsCreateRequest

func (d QrisAssignmentsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, qrisAssignments := range d {
		err = validator.Struct(qrisAssignments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d QrisAssignmentsListCreateRequest) ToModelList() []model.QrisAssignments {
	out := make([]model.QrisAssignments, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type QrisAssignmentsUpdateRequest struct {
	IntentId         uuid.UUID `json:"intentId"`
	QrString         string    `json:"qrString"`
	QrUrl            string    `json:"qrUrl"`
	ExpiresAt        time.Time `json:"expiresAt"`
	PaidAt           time.Time `json:"paidAt"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy"`
}

func (d *QrisAssignmentsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d QrisAssignmentsUpdateRequest) ToModel() model.QrisAssignments {
	return model.QrisAssignments{
		IntentId:         d.IntentId,
		QrString:         d.QrString,
		QrUrl:            null.StringFrom(d.QrUrl),
		ExpiresAt:        d.ExpiresAt,
		PaidAt:           null.TimeFrom(d.PaidAt),
		PspTransactionId: null.StringFrom(d.PspTransactionId),
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
	}
}

type QrisAssignmentsBulkUpdateRequest struct {
	Id               uuid.UUID `json:"id"`
	IntentId         uuid.UUID `json:"intentId"`
	QrString         string    `json:"qrString"`
	QrUrl            string    `json:"qrUrl"`
	ExpiresAt        time.Time `json:"expiresAt"`
	PaidAt           time.Time `json:"paidAt"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy"`
}

func (d QrisAssignmentsBulkUpdateRequest) PrimaryID() QrisAssignmentsPrimaryID {
	return QrisAssignmentsPrimaryID{
		Id: d.Id,
	}
}

type QrisAssignmentsListBulkUpdateRequest []*QrisAssignmentsBulkUpdateRequest

func (d QrisAssignmentsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, qrisAssignments := range d {
		err = validator.Struct(qrisAssignments)
		if err != nil {
			return
		}
	}
	return nil
}

func (d QrisAssignmentsBulkUpdateRequest) ToModel() model.QrisAssignments {
	return model.QrisAssignments{
		Id:               d.Id,
		IntentId:         d.IntentId,
		QrString:         d.QrString,
		QrUrl:            null.StringFrom(d.QrUrl),
		ExpiresAt:        d.ExpiresAt,
		PaidAt:           null.TimeFrom(d.PaidAt),
		PspTransactionId: null.StringFrom(d.PspTransactionId),
		MetaCreatedAt:    d.MetaCreatedAt,
		MetaCreatedBy:    d.MetaCreatedBy,
		MetaUpdatedAt:    d.MetaUpdatedAt,
		MetaUpdatedBy:    nuuid.From(d.MetaUpdatedBy),
	}
}

type QrisAssignmentsResponse struct {
	Id               uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	IntentId         uuid.UUID `json:"intentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	QrString         string    `json:"qrString" validate:"required"`
	QrUrl            string    `json:"qrUrl" validate:"url"`
	ExpiresAt        time.Time `json:"expiresAt" validate:"required" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PaidAt           time.Time `json:"paidAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	PspTransactionId string    `json:"pspTransactionId"`
	MetaCreatedAt    time.Time `json:"metaCreatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaCreatedBy    uuid.UUID `json:"metaCreatedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	MetaUpdatedAt    time.Time `json:"metaUpdatedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	MetaUpdatedBy    uuid.UUID `json:"metaUpdatedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func NewQrisAssignmentsResponse(qrisAssignments model.QrisAssignments) QrisAssignmentsResponse {
	return QrisAssignmentsResponse{
		Id:               qrisAssignments.Id,
		IntentId:         qrisAssignments.IntentId,
		QrString:         qrisAssignments.QrString,
		QrUrl:            qrisAssignments.QrUrl.String,
		ExpiresAt:        qrisAssignments.ExpiresAt,
		PaidAt:           qrisAssignments.PaidAt.Time,
		PspTransactionId: qrisAssignments.PspTransactionId.String,
		MetaCreatedAt:    qrisAssignments.MetaCreatedAt,
		MetaCreatedBy:    qrisAssignments.MetaCreatedBy,
		MetaUpdatedAt:    qrisAssignments.MetaUpdatedAt,
		MetaUpdatedBy:    qrisAssignments.MetaUpdatedBy.UUID,
	}
}

type QrisAssignmentsListResponse []*QrisAssignmentsResponse

func NewQrisAssignmentsListResponse(qrisAssignmentsList model.QrisAssignmentsList) QrisAssignmentsListResponse {
	dtoQrisAssignmentsListResponse := QrisAssignmentsListResponse{}
	for _, qrisAssignments := range qrisAssignmentsList {
		dtoQrisAssignmentsResponse := NewQrisAssignmentsResponse(*qrisAssignments)
		dtoQrisAssignmentsListResponse = append(dtoQrisAssignmentsListResponse, &dtoQrisAssignmentsResponse)
	}
	return dtoQrisAssignmentsListResponse
}

type QrisAssignmentsPrimaryIDList []QrisAssignmentsPrimaryID

func (d QrisAssignmentsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, qrisAssignments := range d {
		err = validator.Struct(qrisAssignments)
		if err != nil {
			return
		}
	}
	return nil
}

type QrisAssignmentsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *QrisAssignmentsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d QrisAssignmentsPrimaryID) ToModel() model.QrisAssignmentsPrimaryID {
	return model.QrisAssignmentsPrimaryID{
		Id: d.Id,
	}
}
