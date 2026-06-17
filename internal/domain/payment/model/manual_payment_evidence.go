package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type ManualPaymentEvidenceDBFieldNameType string

type manualPaymentEvidenceDBFieldName struct {
	Id                        ManualPaymentEvidenceDBFieldNameType
	PaymentIntentId           ManualPaymentEvidenceDBFieldNameType
	PaymentAttemptId          ManualPaymentEvidenceDBFieldNameType
	SubmittedBy               ManualPaymentEvidenceDBFieldNameType
	EvidenceType              ManualPaymentEvidenceDBFieldNameType
	EvidenceUrl               ManualPaymentEvidenceDBFieldNameType
	Amount                    ManualPaymentEvidenceDBFieldNameType
	ExpectedAmount            ManualPaymentEvidenceDBFieldNameType
	VarianceAmount            ManualPaymentEvidenceDBFieldNameType
	VarianceStatus            ManualPaymentEvidenceDBFieldNameType
	Currency                  ManualPaymentEvidenceDBFieldNameType
	BankCode                  ManualPaymentEvidenceDBFieldNameType
	BankName                  ManualPaymentEvidenceDBFieldNameType
	SenderAccountName         ManualPaymentEvidenceDBFieldNameType
	SenderAccountNumberMasked ManualPaymentEvidenceDBFieldNameType
	Notes                     ManualPaymentEvidenceDBFieldNameType
	Status                    ManualPaymentEvidenceDBFieldNameType
	ReviewedBy                ManualPaymentEvidenceDBFieldNameType
	ReviewedAt                ManualPaymentEvidenceDBFieldNameType
	RejectionReason           ManualPaymentEvidenceDBFieldNameType
	PolicyDecision            ManualPaymentEvidenceDBFieldNameType
	Metadata                  ManualPaymentEvidenceDBFieldNameType
	MetaCreatedAt             ManualPaymentEvidenceDBFieldNameType
	MetaCreatedBy             ManualPaymentEvidenceDBFieldNameType
	MetaUpdatedAt             ManualPaymentEvidenceDBFieldNameType
	MetaUpdatedBy             ManualPaymentEvidenceDBFieldNameType
	MetaDeletedAt             ManualPaymentEvidenceDBFieldNameType
	MetaDeletedBy             ManualPaymentEvidenceDBFieldNameType
}

var ManualPaymentEvidenceDBFieldName = manualPaymentEvidenceDBFieldName{
	Id:                        "id",
	PaymentIntentId:           "payment_intent_id",
	PaymentAttemptId:          "payment_attempt_id",
	SubmittedBy:               "submitted_by",
	EvidenceType:              "evidence_type",
	EvidenceUrl:               "evidence_url",
	Amount:                    "amount",
	ExpectedAmount:            "expected_amount",
	VarianceAmount:            "variance_amount",
	VarianceStatus:            "variance_status",
	Currency:                  "currency",
	BankCode:                  "bank_code",
	BankName:                  "bank_name",
	SenderAccountName:         "sender_account_name",
	SenderAccountNumberMasked: "sender_account_number_masked",
	Notes:                     "notes",
	Status:                    "status",
	ReviewedBy:                "reviewed_by",
	ReviewedAt:                "reviewed_at",
	RejectionReason:           "rejection_reason",
	PolicyDecision:            "policy_decision",
	Metadata:                  "metadata",
	MetaCreatedAt:             "meta_created_at",
	MetaCreatedBy:             "meta_created_by",
	MetaUpdatedAt:             "meta_updated_at",
	MetaUpdatedBy:             "meta_updated_by",
	MetaDeletedAt:             "meta_deleted_at",
	MetaDeletedBy:             "meta_deleted_by",
}

func NewManualPaymentEvidenceDBFieldNameFromStr(field string) (dbField ManualPaymentEvidenceDBFieldNameType, found bool) {
	switch field {

	case string(ManualPaymentEvidenceDBFieldName.Id):
		return ManualPaymentEvidenceDBFieldName.Id, true

	case string(ManualPaymentEvidenceDBFieldName.PaymentIntentId):
		return ManualPaymentEvidenceDBFieldName.PaymentIntentId, true

	case string(ManualPaymentEvidenceDBFieldName.PaymentAttemptId):
		return ManualPaymentEvidenceDBFieldName.PaymentAttemptId, true

	case string(ManualPaymentEvidenceDBFieldName.SubmittedBy):
		return ManualPaymentEvidenceDBFieldName.SubmittedBy, true

	case string(ManualPaymentEvidenceDBFieldName.EvidenceType):
		return ManualPaymentEvidenceDBFieldName.EvidenceType, true

	case string(ManualPaymentEvidenceDBFieldName.EvidenceUrl):
		return ManualPaymentEvidenceDBFieldName.EvidenceUrl, true

	case string(ManualPaymentEvidenceDBFieldName.Amount):
		return ManualPaymentEvidenceDBFieldName.Amount, true

	case string(ManualPaymentEvidenceDBFieldName.ExpectedAmount):
		return ManualPaymentEvidenceDBFieldName.ExpectedAmount, true

	case string(ManualPaymentEvidenceDBFieldName.VarianceAmount):
		return ManualPaymentEvidenceDBFieldName.VarianceAmount, true

	case string(ManualPaymentEvidenceDBFieldName.VarianceStatus):
		return ManualPaymentEvidenceDBFieldName.VarianceStatus, true

	case string(ManualPaymentEvidenceDBFieldName.Currency):
		return ManualPaymentEvidenceDBFieldName.Currency, true

	case string(ManualPaymentEvidenceDBFieldName.BankCode):
		return ManualPaymentEvidenceDBFieldName.BankCode, true

	case string(ManualPaymentEvidenceDBFieldName.BankName):
		return ManualPaymentEvidenceDBFieldName.BankName, true

	case string(ManualPaymentEvidenceDBFieldName.SenderAccountName):
		return ManualPaymentEvidenceDBFieldName.SenderAccountName, true

	case string(ManualPaymentEvidenceDBFieldName.SenderAccountNumberMasked):
		return ManualPaymentEvidenceDBFieldName.SenderAccountNumberMasked, true

	case string(ManualPaymentEvidenceDBFieldName.Notes):
		return ManualPaymentEvidenceDBFieldName.Notes, true

	case string(ManualPaymentEvidenceDBFieldName.Status):
		return ManualPaymentEvidenceDBFieldName.Status, true

	case string(ManualPaymentEvidenceDBFieldName.ReviewedBy):
		return ManualPaymentEvidenceDBFieldName.ReviewedBy, true

	case string(ManualPaymentEvidenceDBFieldName.ReviewedAt):
		return ManualPaymentEvidenceDBFieldName.ReviewedAt, true

	case string(ManualPaymentEvidenceDBFieldName.RejectionReason):
		return ManualPaymentEvidenceDBFieldName.RejectionReason, true

	case string(ManualPaymentEvidenceDBFieldName.PolicyDecision):
		return ManualPaymentEvidenceDBFieldName.PolicyDecision, true

	case string(ManualPaymentEvidenceDBFieldName.Metadata):
		return ManualPaymentEvidenceDBFieldName.Metadata, true

	case string(ManualPaymentEvidenceDBFieldName.MetaCreatedAt):
		return ManualPaymentEvidenceDBFieldName.MetaCreatedAt, true

	case string(ManualPaymentEvidenceDBFieldName.MetaCreatedBy):
		return ManualPaymentEvidenceDBFieldName.MetaCreatedBy, true

	case string(ManualPaymentEvidenceDBFieldName.MetaUpdatedAt):
		return ManualPaymentEvidenceDBFieldName.MetaUpdatedAt, true

	case string(ManualPaymentEvidenceDBFieldName.MetaUpdatedBy):
		return ManualPaymentEvidenceDBFieldName.MetaUpdatedBy, true

	case string(ManualPaymentEvidenceDBFieldName.MetaDeletedAt):
		return ManualPaymentEvidenceDBFieldName.MetaDeletedAt, true

	case string(ManualPaymentEvidenceDBFieldName.MetaDeletedBy):
		return ManualPaymentEvidenceDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ManualPaymentEvidenceFilterJoins = map[string]JoinSpec{}

var ManualPaymentEvidenceFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_intent_id": {
		SourcePath:        "payment_intent_id",
		DefaultOutputPath: "paymentIntentId",
		Column:            "payment_intent_id",
		SQLAlias:          "payment_intent_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_attempt_id": {
		SourcePath:        "payment_attempt_id",
		DefaultOutputPath: "paymentAttemptId",
		Column:            "payment_attempt_id",
		SQLAlias:          "payment_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"submitted_by": {
		SourcePath:        "submitted_by",
		DefaultOutputPath: "submittedBy",
		Column:            "submitted_by",
		SQLAlias:          "submitted_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"evidence_type": {
		SourcePath:        "evidence_type",
		DefaultOutputPath: "evidenceType",
		Column:            "evidence_type",
		SQLAlias:          "evidence_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"evidence_url": {
		SourcePath:        "evidence_url",
		DefaultOutputPath: "evidenceUrl",
		Column:            "evidence_url",
		SQLAlias:          "evidence_url",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"expected_amount": {
		SourcePath:        "expected_amount",
		DefaultOutputPath: "expectedAmount",
		Column:            "expected_amount",
		SQLAlias:          "expected_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"variance_amount": {
		SourcePath:        "variance_amount",
		DefaultOutputPath: "varianceAmount",
		Column:            "variance_amount",
		SQLAlias:          "variance_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"variance_status": {
		SourcePath:        "variance_status",
		DefaultOutputPath: "varianceStatus",
		Column:            "variance_status",
		SQLAlias:          "variance_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency": {
		SourcePath:        "currency",
		DefaultOutputPath: "currency",
		Column:            "currency",
		SQLAlias:          "currency",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"bank_code": {
		SourcePath:        "bank_code",
		DefaultOutputPath: "bankCode",
		Column:            "bank_code",
		SQLAlias:          "bank_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"bank_name": {
		SourcePath:        "bank_name",
		DefaultOutputPath: "bankName",
		Column:            "bank_name",
		SQLAlias:          "bank_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"sender_account_name": {
		SourcePath:        "sender_account_name",
		DefaultOutputPath: "senderAccountName",
		Column:            "sender_account_name",
		SQLAlias:          "sender_account_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"sender_account_number_masked": {
		SourcePath:        "sender_account_number_masked",
		DefaultOutputPath: "senderAccountNumberMasked",
		Column:            "sender_account_number_masked",
		SQLAlias:          "sender_account_number_masked",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"notes": {
		SourcePath:        "notes",
		DefaultOutputPath: "notes",
		Column:            "notes",
		SQLAlias:          "notes",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"status": {
		SourcePath:        "status",
		DefaultOutputPath: "status",
		Column:            "status",
		SQLAlias:          "status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reviewed_by": {
		SourcePath:        "reviewed_by",
		DefaultOutputPath: "reviewedBy",
		Column:            "reviewed_by",
		SQLAlias:          "reviewed_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reviewed_at": {
		SourcePath:        "reviewed_at",
		DefaultOutputPath: "reviewedAt",
		Column:            "reviewed_at",
		SQLAlias:          "reviewed_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"rejection_reason": {
		SourcePath:        "rejection_reason",
		DefaultOutputPath: "rejectionReason",
		Column:            "rejection_reason",
		SQLAlias:          "rejection_reason",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"policy_decision": {
		SourcePath:        "policy_decision",
		DefaultOutputPath: "policyDecision",
		Column:            "policy_decision",
		SQLAlias:          "policy_decision",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"metadata": {
		SourcePath:        "metadata",
		DefaultOutputPath: "metadata",
		Column:            "metadata",
		SQLAlias:          "metadata",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_at": {
		SourcePath:        "meta_created_at",
		DefaultOutputPath: "metaCreatedAt",
		Column:            "meta_created_at",
		SQLAlias:          "meta_created_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_by": {
		SourcePath:        "meta_created_by",
		DefaultOutputPath: "metaCreatedBy",
		Column:            "meta_created_by",
		SQLAlias:          "meta_created_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_at": {
		SourcePath:        "meta_updated_at",
		DefaultOutputPath: "metaUpdatedAt",
		Column:            "meta_updated_at",
		SQLAlias:          "meta_updated_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_by": {
		SourcePath:        "meta_updated_by",
		DefaultOutputPath: "metaUpdatedBy",
		Column:            "meta_updated_by",
		SQLAlias:          "meta_updated_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_at": {
		SourcePath:        "meta_deleted_at",
		DefaultOutputPath: "metaDeletedAt",
		Column:            "meta_deleted_at",
		SQLAlias:          "meta_deleted_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_by": {
		SourcePath:        "meta_deleted_by",
		DefaultOutputPath: "metaDeletedBy",
		Column:            "meta_deleted_by",
		SQLAlias:          "meta_deleted_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
}

func NewManualPaymentEvidenceFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ManualPaymentEvidenceFilterFields[field]
	return
}

type ManualPaymentEvidenceFilterResult struct {
	ManualPaymentEvidence
	FilterCount int `db:"count"`
}

func ValidateManualPaymentEvidenceFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewManualPaymentEvidenceFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewManualPaymentEvidenceFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewManualPaymentEvidenceFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateManualPaymentEvidenceFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateManualPaymentEvidenceFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewManualPaymentEvidenceFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateManualPaymentEvidenceFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ManualEvidenceStatus string

const (
	ManualEvidenceStatusSubmitted   ManualEvidenceStatus = "submitted"
	ManualEvidenceStatusUnderReview ManualEvidenceStatus = "under_review"
	ManualEvidenceStatusApproved    ManualEvidenceStatus = "approved"
	ManualEvidenceStatusRejected    ManualEvidenceStatus = "rejected"
)

type ManualPaymentEvidence struct {
	Id                        uuid.UUID            `db:"id"`
	PaymentIntentId           uuid.UUID            `db:"payment_intent_id"`
	PaymentAttemptId          nuuid.NUUID          `db:"payment_attempt_id"`
	SubmittedBy               uuid.UUID            `db:"submitted_by"`
	EvidenceType              string               `db:"evidence_type"`
	EvidenceUrl               null.String          `db:"evidence_url"`
	Amount                    decimal.Decimal      `db:"amount"`
	ExpectedAmount            decimal.Decimal      `db:"expected_amount"`
	VarianceAmount            decimal.Decimal      `db:"variance_amount"`
	VarianceStatus            string               `db:"variance_status"`
	Currency                  string               `db:"currency"`
	BankCode                  null.String          `db:"bank_code"`
	BankName                  null.String          `db:"bank_name"`
	SenderAccountName         null.String          `db:"sender_account_name"`
	SenderAccountNumberMasked null.String          `db:"sender_account_number_masked"`
	Notes                     null.String          `db:"notes"`
	Status                    ManualEvidenceStatus `db:"status"`
	ReviewedBy                nuuid.NUUID          `db:"reviewed_by"`
	ReviewedAt                null.Time            `db:"reviewed_at"`
	RejectionReason           null.String          `db:"rejection_reason"`
	PolicyDecision            null.String          `db:"policy_decision"`
	Metadata                  json.RawMessage      `db:"metadata"`

	shared.MetaSignature
}
type ManualPaymentEvidencePrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ManualPaymentEvidence) ToManualPaymentEvidencePrimaryID() ManualPaymentEvidencePrimaryID {
	return ManualPaymentEvidencePrimaryID{
		Id: d.Id,
	}
}

type ManualPaymentEvidenceList []*ManualPaymentEvidence

type ManualPaymentEvidenceFilterResultList []*ManualPaymentEvidenceFilterResult
