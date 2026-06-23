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

type RefundApprovalsDBFieldNameType string

type refundApprovalsDBFieldName struct {
	Id             RefundApprovalsDBFieldNameType
	RefundId       RefundApprovalsDBFieldNameType
	ApprovalStatus RefundApprovalsDBFieldNameType
	ApprovedAmount RefundApprovalsDBFieldNameType
	ReasonCode     RefundApprovalsDBFieldNameType
	ReasonDetail   RefundApprovalsDBFieldNameType
	ApprovedBy     RefundApprovalsDBFieldNameType
	ApprovedAt     RefundApprovalsDBFieldNameType
	Metadata       RefundApprovalsDBFieldNameType
	MetaCreatedAt  RefundApprovalsDBFieldNameType
	MetaCreatedBy  RefundApprovalsDBFieldNameType
	MetaUpdatedAt  RefundApprovalsDBFieldNameType
	MetaUpdatedBy  RefundApprovalsDBFieldNameType
	MetaDeletedAt  RefundApprovalsDBFieldNameType
	MetaDeletedBy  RefundApprovalsDBFieldNameType
}

var RefundApprovalsDBFieldName = refundApprovalsDBFieldName{
	Id:             "id",
	RefundId:       "refund_id",
	ApprovalStatus: "approval_status",
	ApprovedAmount: "approved_amount",
	ReasonCode:     "reason_code",
	ReasonDetail:   "reason_detail",
	ApprovedBy:     "approved_by",
	ApprovedAt:     "approved_at",
	Metadata:       "metadata",
	MetaCreatedAt:  "meta_created_at",
	MetaCreatedBy:  "meta_created_by",
	MetaUpdatedAt:  "meta_updated_at",
	MetaUpdatedBy:  "meta_updated_by",
	MetaDeletedAt:  "meta_deleted_at",
	MetaDeletedBy:  "meta_deleted_by",
}

func NewRefundApprovalsDBFieldNameFromStr(field string) (dbField RefundApprovalsDBFieldNameType, found bool) {
	switch field {

	case string(RefundApprovalsDBFieldName.Id):
		return RefundApprovalsDBFieldName.Id, true

	case string(RefundApprovalsDBFieldName.RefundId):
		return RefundApprovalsDBFieldName.RefundId, true

	case string(RefundApprovalsDBFieldName.ApprovalStatus):
		return RefundApprovalsDBFieldName.ApprovalStatus, true

	case string(RefundApprovalsDBFieldName.ApprovedAmount):
		return RefundApprovalsDBFieldName.ApprovedAmount, true

	case string(RefundApprovalsDBFieldName.ReasonCode):
		return RefundApprovalsDBFieldName.ReasonCode, true

	case string(RefundApprovalsDBFieldName.ReasonDetail):
		return RefundApprovalsDBFieldName.ReasonDetail, true

	case string(RefundApprovalsDBFieldName.ApprovedBy):
		return RefundApprovalsDBFieldName.ApprovedBy, true

	case string(RefundApprovalsDBFieldName.ApprovedAt):
		return RefundApprovalsDBFieldName.ApprovedAt, true

	case string(RefundApprovalsDBFieldName.Metadata):
		return RefundApprovalsDBFieldName.Metadata, true

	case string(RefundApprovalsDBFieldName.MetaCreatedAt):
		return RefundApprovalsDBFieldName.MetaCreatedAt, true

	case string(RefundApprovalsDBFieldName.MetaCreatedBy):
		return RefundApprovalsDBFieldName.MetaCreatedBy, true

	case string(RefundApprovalsDBFieldName.MetaUpdatedAt):
		return RefundApprovalsDBFieldName.MetaUpdatedAt, true

	case string(RefundApprovalsDBFieldName.MetaUpdatedBy):
		return RefundApprovalsDBFieldName.MetaUpdatedBy, true

	case string(RefundApprovalsDBFieldName.MetaDeletedAt):
		return RefundApprovalsDBFieldName.MetaDeletedAt, true

	case string(RefundApprovalsDBFieldName.MetaDeletedBy):
		return RefundApprovalsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var RefundApprovalsFilterJoins = map[string]JoinSpec{}

var RefundApprovalsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refund_id": {
		SourcePath:        "refund_id",
		DefaultOutputPath: "refundId",
		Column:            "refund_id",
		SQLAlias:          "refund_id",
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
	"approved_amount": {
		SourcePath:        "approved_amount",
		DefaultOutputPath: "approvedAmount",
		Column:            "approved_amount",
		SQLAlias:          "approved_amount",
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

func NewRefundApprovalsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = RefundApprovalsFilterFields[field]
	return
}

type RefundApprovalsFilterResult struct {
	RefundApprovals
	FilterCount int `db:"count"`
}

func ValidateRefundApprovalsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewRefundApprovalsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewRefundApprovalsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewRefundApprovalsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateRefundApprovalsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateRefundApprovalsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewRefundApprovalsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateRefundApprovalsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type RefundApprovalsApprovalStatus string

const (
	RefundApprovalsApprovalStatusApproved RefundApprovalsApprovalStatus = "approved"
	RefundApprovalsApprovalStatusRejected RefundApprovalsApprovalStatus = "rejected"
)

type RefundApprovals struct {
	Id             uuid.UUID                     `db:"id"`
	RefundId       uuid.UUID                     `db:"refund_id"`
	ApprovalStatus RefundApprovalsApprovalStatus `db:"approval_status"`
	ApprovedAmount decimal.Decimal               `db:"approved_amount"`
	ReasonCode     string                        `db:"reason_code"`
	ReasonDetail   null.String                   `db:"reason_detail"`
	ApprovedBy     uuid.UUID                     `db:"approved_by"`
	ApprovedAt     time.Time                     `db:"approved_at"`
	Metadata       json.RawMessage               `db:"metadata"`

	shared.MetaSignature
}
type RefundApprovalsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RefundApprovals) ToRefundApprovalsPrimaryID() RefundApprovalsPrimaryID {
	return RefundApprovalsPrimaryID{
		Id: d.Id,
	}
}

type RefundApprovalsList []*RefundApprovals

type RefundApprovalsFilterResultList []*RefundApprovalsFilterResult
