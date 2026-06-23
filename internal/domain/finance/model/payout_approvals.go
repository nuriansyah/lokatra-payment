package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type PayoutApprovalsDBFieldNameType string

type payoutApprovalsDBFieldName struct {
	Id                     PayoutApprovalsDBFieldNameType
	PayoutId               PayoutApprovalsDBFieldNameType
	ApprovalStatus         PayoutApprovalsDBFieldNameType
	ReasonCode             PayoutApprovalsDBFieldNameType
	ReasonDetail           PayoutApprovalsDBFieldNameType
	ApprovedBy             PayoutApprovalsDBFieldNameType
	ApprovedAt             PayoutApprovalsDBFieldNameType
	Metadata               PayoutApprovalsDBFieldNameType
	ApprovedAmountSnapshot PayoutApprovalsDBFieldNameType
	CurrencyCodeSnapshot   PayoutApprovalsDBFieldNameType
	PayoutRevisionHash     PayoutApprovalsDBFieldNameType
	MetaCreatedAt          PayoutApprovalsDBFieldNameType
	MetaCreatedBy          PayoutApprovalsDBFieldNameType
	MetaUpdatedAt          PayoutApprovalsDBFieldNameType
	MetaUpdatedBy          PayoutApprovalsDBFieldNameType
	MetaDeletedAt          PayoutApprovalsDBFieldNameType
	MetaDeletedBy          PayoutApprovalsDBFieldNameType
}

var PayoutApprovalsDBFieldName = payoutApprovalsDBFieldName{
	Id:                     "id",
	PayoutId:               "payout_id",
	ApprovalStatus:         "approval_status",
	ReasonCode:             "reason_code",
	ReasonDetail:           "reason_detail",
	ApprovedBy:             "approved_by",
	ApprovedAt:             "approved_at",
	Metadata:               "metadata",
	ApprovedAmountSnapshot: "approved_amount_snapshot",
	CurrencyCodeSnapshot:   "currency_code_snapshot",
	PayoutRevisionHash:     "payout_revision_hash",
	MetaCreatedAt:          "meta_created_at",
	MetaCreatedBy:          "meta_created_by",
	MetaUpdatedAt:          "meta_updated_at",
	MetaUpdatedBy:          "meta_updated_by",
	MetaDeletedAt:          "meta_deleted_at",
	MetaDeletedBy:          "meta_deleted_by",
}

func NewPayoutApprovalsDBFieldNameFromStr(field string) (dbField PayoutApprovalsDBFieldNameType, found bool) {
	switch field {

	case string(PayoutApprovalsDBFieldName.Id):
		return PayoutApprovalsDBFieldName.Id, true

	case string(PayoutApprovalsDBFieldName.PayoutId):
		return PayoutApprovalsDBFieldName.PayoutId, true

	case string(PayoutApprovalsDBFieldName.ApprovalStatus):
		return PayoutApprovalsDBFieldName.ApprovalStatus, true

	case string(PayoutApprovalsDBFieldName.ReasonCode):
		return PayoutApprovalsDBFieldName.ReasonCode, true

	case string(PayoutApprovalsDBFieldName.ReasonDetail):
		return PayoutApprovalsDBFieldName.ReasonDetail, true

	case string(PayoutApprovalsDBFieldName.ApprovedBy):
		return PayoutApprovalsDBFieldName.ApprovedBy, true

	case string(PayoutApprovalsDBFieldName.ApprovedAt):
		return PayoutApprovalsDBFieldName.ApprovedAt, true

	case string(PayoutApprovalsDBFieldName.Metadata):
		return PayoutApprovalsDBFieldName.Metadata, true

	case string(PayoutApprovalsDBFieldName.ApprovedAmountSnapshot):
		return PayoutApprovalsDBFieldName.ApprovedAmountSnapshot, true

	case string(PayoutApprovalsDBFieldName.CurrencyCodeSnapshot):
		return PayoutApprovalsDBFieldName.CurrencyCodeSnapshot, true

	case string(PayoutApprovalsDBFieldName.PayoutRevisionHash):
		return PayoutApprovalsDBFieldName.PayoutRevisionHash, true

	case string(PayoutApprovalsDBFieldName.MetaCreatedAt):
		return PayoutApprovalsDBFieldName.MetaCreatedAt, true

	case string(PayoutApprovalsDBFieldName.MetaCreatedBy):
		return PayoutApprovalsDBFieldName.MetaCreatedBy, true

	case string(PayoutApprovalsDBFieldName.MetaUpdatedAt):
		return PayoutApprovalsDBFieldName.MetaUpdatedAt, true

	case string(PayoutApprovalsDBFieldName.MetaUpdatedBy):
		return PayoutApprovalsDBFieldName.MetaUpdatedBy, true

	case string(PayoutApprovalsDBFieldName.MetaDeletedAt):
		return PayoutApprovalsDBFieldName.MetaDeletedAt, true

	case string(PayoutApprovalsDBFieldName.MetaDeletedBy):
		return PayoutApprovalsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PayoutApprovalsFilterJoins = map[string]JoinSpec{}

var PayoutApprovalsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_id": {
		SourcePath:        "payout_id",
		DefaultOutputPath: "payoutId",
		Column:            "payout_id",
		SQLAlias:          "payout_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"approval_status": {
		SourcePath:        "approval_status",
		DefaultOutputPath: "approvalStatus",
		Column:            "approval_status",
		SQLAlias:          "approval_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reason_code": {
		SourcePath:        "reason_code",
		DefaultOutputPath: "reasonCode",
		Column:            "reason_code",
		SQLAlias:          "reason_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reason_detail": {
		SourcePath:        "reason_detail",
		DefaultOutputPath: "reasonDetail",
		Column:            "reason_detail",
		SQLAlias:          "reason_detail",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"approved_by": {
		SourcePath:        "approved_by",
		DefaultOutputPath: "approvedBy",
		Column:            "approved_by",
		SQLAlias:          "approved_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"approved_at": {
		SourcePath:        "approved_at",
		DefaultOutputPath: "approvedAt",
		Column:            "approved_at",
		SQLAlias:          "approved_at",
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
	"approved_amount_snapshot": {
		SourcePath:        "approved_amount_snapshot",
		DefaultOutputPath: "approvedAmountSnapshot",
		Column:            "approved_amount_snapshot",
		SQLAlias:          "approved_amount_snapshot",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency_code_snapshot": {
		SourcePath:        "currency_code_snapshot",
		DefaultOutputPath: "currencyCodeSnapshot",
		Column:            "currency_code_snapshot",
		SQLAlias:          "currency_code_snapshot",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_revision_hash": {
		SourcePath:        "payout_revision_hash",
		DefaultOutputPath: "payoutRevisionHash",
		Column:            "payout_revision_hash",
		SQLAlias:          "payout_revision_hash",
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

func NewPayoutApprovalsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayoutApprovalsFilterFields[field]
	return
}

type PayoutApprovalsFilterResult struct {
	PayoutApprovals
	FilterCount int `db:"count"`
}

func ValidatePayoutApprovalsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPayoutApprovalsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayoutApprovalsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPayoutApprovalsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePayoutApprovalsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePayoutApprovalsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPayoutApprovalsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePayoutApprovalsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PayoutApprovalsApprovalStatus string

const (
	PayoutApprovalsApprovalStatusApproved PayoutApprovalsApprovalStatus = "approved"
	PayoutApprovalsApprovalStatusRejected PayoutApprovalsApprovalStatus = "rejected"
)

type PayoutApprovals struct {
	Id                     uuid.UUID                     `db:"id"`
	PayoutId               uuid.UUID                     `db:"payout_id"`
	ApprovalStatus         PayoutApprovalsApprovalStatus `db:"approval_status"`
	ReasonCode             string                        `db:"reason_code"`
	ReasonDetail           null.String                   `db:"reason_detail"`
	ApprovedBy             uuid.UUID                     `db:"approved_by"`
	ApprovedAt             time.Time                     `db:"approved_at"`
	Metadata               json.RawMessage               `db:"metadata"`
	ApprovedAmountSnapshot decimal.NullDecimal           `db:"approved_amount_snapshot"`
	CurrencyCodeSnapshot   null.String                   `db:"currency_code_snapshot"`
	PayoutRevisionHash     null.String                   `db:"payout_revision_hash"`

	shared.MetaSignature
}
type PayoutApprovalsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayoutApprovals) ToPayoutApprovalsPrimaryID() PayoutApprovalsPrimaryID {
	return PayoutApprovalsPrimaryID{
		Id: d.Id,
	}
}

type PayoutApprovalsList []*PayoutApprovals

type PayoutApprovalsFilterResultList []*PayoutApprovalsFilterResult
