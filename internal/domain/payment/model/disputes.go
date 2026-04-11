package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type DisputesDBFieldNameType string

type disputesDBFieldName struct {
	Id                  DisputesDBFieldNameType
	DisputeCode         DisputesDBFieldNameType
	PaymentId           DisputesDBFieldNameType
	PspDisputeId        DisputesDBFieldNameType
	DisputeType         DisputesDBFieldNameType
	ReasonCode          DisputesDBFieldNameType
	ReasonDescription   DisputesDBFieldNameType
	Amount              DisputesDBFieldNameType
	Currency            DisputesDBFieldNameType
	Status              DisputesDBFieldNameType
	OpenedAt            DisputesDBFieldNameType
	RespondBy           DisputesDBFieldNameType
	ResolvedAt          DisputesDBFieldNameType
	Outcome             DisputesDBFieldNameType
	EvidenceDueAt       DisputesDBFieldNameType
	EvidenceSubmittedAt DisputesDBFieldNameType
	EvidenceFiles       DisputesDBFieldNameType
	Notes               DisputesDBFieldNameType
	HandledBy           DisputesDBFieldNameType
	MetaCreatedAt       DisputesDBFieldNameType
	MetaCreatedBy       DisputesDBFieldNameType
	MetaUpdatedAt       DisputesDBFieldNameType
	MetaUpdatedBy       DisputesDBFieldNameType
	MetaDeletedAt       DisputesDBFieldNameType
	MetaDeletedBy       DisputesDBFieldNameType
}

var DisputesDBFieldName = disputesDBFieldName{
	Id:                  "id",
	DisputeCode:         "dispute_code",
	PaymentId:           "payment_id",
	PspDisputeId:        "psp_dispute_id",
	DisputeType:         "dispute_type",
	ReasonCode:          "reason_code",
	ReasonDescription:   "reason_description",
	Amount:              "amount",
	Currency:            "currency",
	Status:              "status",
	OpenedAt:            "opened_at",
	RespondBy:           "respond_by",
	ResolvedAt:          "resolved_at",
	Outcome:             "outcome",
	EvidenceDueAt:       "evidence_due_at",
	EvidenceSubmittedAt: "evidence_submitted_at",
	EvidenceFiles:       "evidence_files",
	Notes:               "notes",
	HandledBy:           "handled_by",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewDisputesDBFieldNameFromStr(field string) (dbField DisputesDBFieldNameType, found bool) {
	switch field {

	case string(DisputesDBFieldName.Id):
		return DisputesDBFieldName.Id, true

	case string(DisputesDBFieldName.DisputeCode):
		return DisputesDBFieldName.DisputeCode, true

	case string(DisputesDBFieldName.PaymentId):
		return DisputesDBFieldName.PaymentId, true

	case string(DisputesDBFieldName.PspDisputeId):
		return DisputesDBFieldName.PspDisputeId, true

	case string(DisputesDBFieldName.DisputeType):
		return DisputesDBFieldName.DisputeType, true

	case string(DisputesDBFieldName.ReasonCode):
		return DisputesDBFieldName.ReasonCode, true

	case string(DisputesDBFieldName.ReasonDescription):
		return DisputesDBFieldName.ReasonDescription, true

	case string(DisputesDBFieldName.Amount):
		return DisputesDBFieldName.Amount, true

	case string(DisputesDBFieldName.Currency):
		return DisputesDBFieldName.Currency, true

	case string(DisputesDBFieldName.Status):
		return DisputesDBFieldName.Status, true

	case string(DisputesDBFieldName.OpenedAt):
		return DisputesDBFieldName.OpenedAt, true

	case string(DisputesDBFieldName.RespondBy):
		return DisputesDBFieldName.RespondBy, true

	case string(DisputesDBFieldName.ResolvedAt):
		return DisputesDBFieldName.ResolvedAt, true

	case string(DisputesDBFieldName.Outcome):
		return DisputesDBFieldName.Outcome, true

	case string(DisputesDBFieldName.EvidenceDueAt):
		return DisputesDBFieldName.EvidenceDueAt, true

	case string(DisputesDBFieldName.EvidenceSubmittedAt):
		return DisputesDBFieldName.EvidenceSubmittedAt, true

	case string(DisputesDBFieldName.EvidenceFiles):
		return DisputesDBFieldName.EvidenceFiles, true

	case string(DisputesDBFieldName.Notes):
		return DisputesDBFieldName.Notes, true

	case string(DisputesDBFieldName.HandledBy):
		return DisputesDBFieldName.HandledBy, true

	case string(DisputesDBFieldName.MetaCreatedAt):
		return DisputesDBFieldName.MetaCreatedAt, true

	case string(DisputesDBFieldName.MetaCreatedBy):
		return DisputesDBFieldName.MetaCreatedBy, true

	case string(DisputesDBFieldName.MetaUpdatedAt):
		return DisputesDBFieldName.MetaUpdatedAt, true

	case string(DisputesDBFieldName.MetaUpdatedBy):
		return DisputesDBFieldName.MetaUpdatedBy, true

	case string(DisputesDBFieldName.MetaDeletedAt):
		return DisputesDBFieldName.MetaDeletedAt, true

	case string(DisputesDBFieldName.MetaDeletedBy):
		return DisputesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type DisputesFilterResult struct {
	Disputes
	FilterCount int `db:"count"`
}

func ValidateDisputesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewDisputesDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewDisputesDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewDisputesDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type DisputeStatus string

const (
	DisputeStatusOpen              DisputeStatus = "OPEN"
	DisputeStatusEvidenceSubmitted DisputeStatus = "EVIDENCE_SUBMITTED"
	DisputeStatusWon               DisputeStatus = "WON"
	DisputeStatusLost              DisputeStatus = "LOST"
	DisputeStatusClosedByIssuer    DisputeStatus = "CLOSED_BY_ISSUER"
)

type Disputes struct {
	Id                  uuid.UUID       `db:"id"`
	DisputeCode         string          `db:"dispute_code"`
	PaymentId           uuid.UUID       `db:"payment_id"`
	PspDisputeId        string          `db:"psp_dispute_id"`
	DisputeType         string          `db:"dispute_type"`
	ReasonCode          null.String     `db:"reason_code"`
	ReasonDescription   null.String     `db:"reason_description"`
	Amount              decimal.Decimal `db:"amount"`
	Currency            PaymentCurrency `db:"currency"`
	Status              DisputeStatus   `db:"status"`
	OpenedAt            time.Time       `db:"opened_at"`
	RespondBy           null.Time       `db:"respond_by"`
	ResolvedAt          null.Time       `db:"resolved_at"`
	Outcome             null.String     `db:"outcome"`
	EvidenceDueAt       null.Time       `db:"evidence_due_at"`
	EvidenceSubmittedAt null.Time       `db:"evidence_submitted_at"`
	EvidenceFiles       json.RawMessage `db:"evidence_files"`
	Notes               null.String     `db:"notes"`
	HandledBy           nuuid.NUUID     `db:"handled_by"`
	MetaCreatedAt       time.Time       `db:"meta_created_at"`
	MetaCreatedBy       uuid.UUID       `db:"meta_created_by"`
	MetaUpdatedAt       time.Time       `db:"meta_updated_at"`
	MetaUpdatedBy       nuuid.NUUID     `db:"meta_updated_by"`
	MetaDeletedAt       null.Time       `db:"meta_deleted_at"`
	MetaDeletedBy       nuuid.NUUID     `db:"meta_deleted_by"`
}
type DisputesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d Disputes) ToDisputesPrimaryID() DisputesPrimaryID {
	return DisputesPrimaryID{
		Id: d.Id,
	}
}

type DisputesList []*Disputes
