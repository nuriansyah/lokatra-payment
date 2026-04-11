package model

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type QrisAssignmentsDBFieldNameType string

type qrisAssignmentsDBFieldName struct {
	Id               QrisAssignmentsDBFieldNameType
	IntentId         QrisAssignmentsDBFieldNameType
	QrString         QrisAssignmentsDBFieldNameType
	QrUrl            QrisAssignmentsDBFieldNameType
	ExpiresAt        QrisAssignmentsDBFieldNameType
	PaidAt           QrisAssignmentsDBFieldNameType
	PspTransactionId QrisAssignmentsDBFieldNameType
	MetaCreatedAt    QrisAssignmentsDBFieldNameType
	MetaCreatedBy    QrisAssignmentsDBFieldNameType
	MetaUpdatedAt    QrisAssignmentsDBFieldNameType
	MetaUpdatedBy    QrisAssignmentsDBFieldNameType
}

var QrisAssignmentsDBFieldName = qrisAssignmentsDBFieldName{
	Id:               "id",
	IntentId:         "intent_id",
	QrString:         "qr_string",
	QrUrl:            "qr_url",
	ExpiresAt:        "expires_at",
	PaidAt:           "paid_at",
	PspTransactionId: "psp_transaction_id",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
}

func NewQrisAssignmentsDBFieldNameFromStr(field string) (dbField QrisAssignmentsDBFieldNameType, found bool) {
	switch field {

	case string(QrisAssignmentsDBFieldName.Id):
		return QrisAssignmentsDBFieldName.Id, true

	case string(QrisAssignmentsDBFieldName.IntentId):
		return QrisAssignmentsDBFieldName.IntentId, true

	case string(QrisAssignmentsDBFieldName.QrString):
		return QrisAssignmentsDBFieldName.QrString, true

	case string(QrisAssignmentsDBFieldName.QrUrl):
		return QrisAssignmentsDBFieldName.QrUrl, true

	case string(QrisAssignmentsDBFieldName.ExpiresAt):
		return QrisAssignmentsDBFieldName.ExpiresAt, true

	case string(QrisAssignmentsDBFieldName.PaidAt):
		return QrisAssignmentsDBFieldName.PaidAt, true

	case string(QrisAssignmentsDBFieldName.PspTransactionId):
		return QrisAssignmentsDBFieldName.PspTransactionId, true

	case string(QrisAssignmentsDBFieldName.MetaCreatedAt):
		return QrisAssignmentsDBFieldName.MetaCreatedAt, true

	case string(QrisAssignmentsDBFieldName.MetaCreatedBy):
		return QrisAssignmentsDBFieldName.MetaCreatedBy, true

	case string(QrisAssignmentsDBFieldName.MetaUpdatedAt):
		return QrisAssignmentsDBFieldName.MetaUpdatedAt, true

	case string(QrisAssignmentsDBFieldName.MetaUpdatedBy):
		return QrisAssignmentsDBFieldName.MetaUpdatedBy, true

	}
	return "unknown", false
}

type QrisAssignmentsFilterResult struct {
	QrisAssignments
	FilterCount int `db:"count"`
}

func ValidateQrisAssignmentsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewQrisAssignmentsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewQrisAssignmentsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewQrisAssignmentsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type QrisAssignments struct {
	Id               uuid.UUID   `db:"id"`
	IntentId         uuid.UUID   `db:"intent_id"`
	QrString         string      `db:"qr_string"`
	QrUrl            null.String `db:"qr_url"`
	ExpiresAt        time.Time   `db:"expires_at"`
	PaidAt           null.Time   `db:"paid_at"`
	PspTransactionId null.String `db:"psp_transaction_id"`
	MetaCreatedAt    time.Time   `db:"meta_created_at"`
	MetaCreatedBy    uuid.UUID   `db:"meta_created_by"`
	MetaUpdatedAt    time.Time   `db:"meta_updated_at"`
	MetaUpdatedBy    nuuid.NUUID `db:"meta_updated_by"`
}
type QrisAssignmentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d QrisAssignments) ToQrisAssignmentsPrimaryID() QrisAssignmentsPrimaryID {
	return QrisAssignmentsPrimaryID{
		Id: d.Id,
	}
}

type QrisAssignmentsList []*QrisAssignments
