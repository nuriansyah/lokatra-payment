package model

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type VirtualAccountAssignmentsDBFieldNameType string

type virtualAccountAssignmentsDBFieldName struct {
	Id               VirtualAccountAssignmentsDBFieldNameType
	IntentId         VirtualAccountAssignmentsDBFieldNameType
	BankCode         VirtualAccountAssignmentsDBFieldNameType
	VaNumber         VirtualAccountAssignmentsDBFieldNameType
	VaNumberMasked   VirtualAccountAssignmentsDBFieldNameType
	ExpiresAt        VirtualAccountAssignmentsDBFieldNameType
	IsReusable       VirtualAccountAssignmentsDBFieldNameType
	PaidAt           VirtualAccountAssignmentsDBFieldNameType
	PspTransactionId VirtualAccountAssignmentsDBFieldNameType
	MetaCreatedAt    VirtualAccountAssignmentsDBFieldNameType
	MetaCreatedBy    VirtualAccountAssignmentsDBFieldNameType
	MetaUpdatedAt    VirtualAccountAssignmentsDBFieldNameType
	MetaUpdatedBy    VirtualAccountAssignmentsDBFieldNameType
}

var VirtualAccountAssignmentsDBFieldName = virtualAccountAssignmentsDBFieldName{
	Id:               "id",
	IntentId:         "intent_id",
	BankCode:         "bank_code",
	VaNumber:         "va_number",
	VaNumberMasked:   "va_number_masked",
	ExpiresAt:        "expires_at",
	IsReusable:       "is_reusable",
	PaidAt:           "paid_at",
	PspTransactionId: "psp_transaction_id",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
}

func NewVirtualAccountAssignmentsDBFieldNameFromStr(field string) (dbField VirtualAccountAssignmentsDBFieldNameType, found bool) {
	switch field {

	case string(VirtualAccountAssignmentsDBFieldName.Id):
		return VirtualAccountAssignmentsDBFieldName.Id, true

	case string(VirtualAccountAssignmentsDBFieldName.IntentId):
		return VirtualAccountAssignmentsDBFieldName.IntentId, true

	case string(VirtualAccountAssignmentsDBFieldName.BankCode):
		return VirtualAccountAssignmentsDBFieldName.BankCode, true

	case string(VirtualAccountAssignmentsDBFieldName.VaNumber):
		return VirtualAccountAssignmentsDBFieldName.VaNumber, true

	case string(VirtualAccountAssignmentsDBFieldName.VaNumberMasked):
		return VirtualAccountAssignmentsDBFieldName.VaNumberMasked, true

	case string(VirtualAccountAssignmentsDBFieldName.ExpiresAt):
		return VirtualAccountAssignmentsDBFieldName.ExpiresAt, true

	case string(VirtualAccountAssignmentsDBFieldName.IsReusable):
		return VirtualAccountAssignmentsDBFieldName.IsReusable, true

	case string(VirtualAccountAssignmentsDBFieldName.PaidAt):
		return VirtualAccountAssignmentsDBFieldName.PaidAt, true

	case string(VirtualAccountAssignmentsDBFieldName.PspTransactionId):
		return VirtualAccountAssignmentsDBFieldName.PspTransactionId, true

	case string(VirtualAccountAssignmentsDBFieldName.MetaCreatedAt):
		return VirtualAccountAssignmentsDBFieldName.MetaCreatedAt, true

	case string(VirtualAccountAssignmentsDBFieldName.MetaCreatedBy):
		return VirtualAccountAssignmentsDBFieldName.MetaCreatedBy, true

	case string(VirtualAccountAssignmentsDBFieldName.MetaUpdatedAt):
		return VirtualAccountAssignmentsDBFieldName.MetaUpdatedAt, true

	case string(VirtualAccountAssignmentsDBFieldName.MetaUpdatedBy):
		return VirtualAccountAssignmentsDBFieldName.MetaUpdatedBy, true

	}
	return "unknown", false
}

type VirtualAccountAssignmentsFilterResult struct {
	VirtualAccountAssignments
	FilterCount int `db:"count"`
}

func ValidateVirtualAccountAssignmentsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewVirtualAccountAssignmentsDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewVirtualAccountAssignmentsDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewVirtualAccountAssignmentsDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type VirtualAccountAssignments struct {
	Id               uuid.UUID   `db:"id"`
	IntentId         uuid.UUID   `db:"intent_id"`
	BankCode         string      `db:"bank_code"`
	VaNumber         string      `db:"va_number"`
	VaNumberMasked   null.String `db:"va_number_masked"`
	ExpiresAt        time.Time   `db:"expires_at"`
	IsReusable       bool        `db:"is_reusable"`
	PaidAt           null.Time   `db:"paid_at"`
	PspTransactionId null.String `db:"psp_transaction_id"`
	MetaCreatedAt    time.Time   `db:"meta_created_at"`
	MetaCreatedBy    uuid.UUID   `db:"meta_created_by"`
	MetaUpdatedAt    time.Time   `db:"meta_updated_at"`
	MetaUpdatedBy    nuuid.NUUID `db:"meta_updated_by"`
}
type VirtualAccountAssignmentsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d VirtualAccountAssignments) ToVirtualAccountAssignmentsPrimaryID() VirtualAccountAssignmentsPrimaryID {
	return VirtualAccountAssignmentsPrimaryID{
		Id: d.Id,
	}
}

type VirtualAccountAssignmentsList []*VirtualAccountAssignments
