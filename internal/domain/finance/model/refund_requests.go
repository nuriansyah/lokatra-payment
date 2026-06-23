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
	"time"
)

type RefundRequestsDBFieldNameType string

type refundRequestsDBFieldName struct {
	Id                   RefundRequestsDBFieldNameType
	RefundCode           RefundRequestsDBFieldNameType
	PaymentRefId         RefundRequestsDBFieldNameType
	MerchantPartyId      RefundRequestsDBFieldNameType
	CustomerPartyId      RefundRequestsDBFieldNameType
	RefundPolicyId       RefundRequestsDBFieldNameType
	CurrencyCode         RefundRequestsDBFieldNameType
	RequestedAmount      RefundRequestsDBFieldNameType
	ApprovedAmount       RefundRequestsDBFieldNameType
	RefundReasonCode     RefundRequestsDBFieldNameType
	RefundStatus         RefundRequestsDBFieldNameType
	RequestedAt          RefundRequestsDBFieldNameType
	ApprovedAt           RefundRequestsDBFieldNameType
	SettledFinanciallyAt RefundRequestsDBFieldNameType
	Metadata             RefundRequestsDBFieldNameType
	MetaCreatedAt        RefundRequestsDBFieldNameType
	MetaCreatedBy        RefundRequestsDBFieldNameType
	MetaUpdatedAt        RefundRequestsDBFieldNameType
	MetaUpdatedBy        RefundRequestsDBFieldNameType
	MetaDeletedAt        RefundRequestsDBFieldNameType
	MetaDeletedBy        RefundRequestsDBFieldNameType
}

var RefundRequestsDBFieldName = refundRequestsDBFieldName{
	Id:                   "id",
	RefundCode:           "refund_code",
	PaymentRefId:         "payment_ref_id",
	MerchantPartyId:      "merchant_party_id",
	CustomerPartyId:      "customer_party_id",
	RefundPolicyId:       "refund_policy_id",
	CurrencyCode:         "currency_code",
	RequestedAmount:      "requested_amount",
	ApprovedAmount:       "approved_amount",
	RefundReasonCode:     "refund_reason_code",
	RefundStatus:         "refund_status",
	RequestedAt:          "requested_at",
	ApprovedAt:           "approved_at",
	SettledFinanciallyAt: "settled_financially_at",
	Metadata:             "metadata",
	MetaCreatedAt:        "meta_created_at",
	MetaCreatedBy:        "meta_created_by",
	MetaUpdatedAt:        "meta_updated_at",
	MetaUpdatedBy:        "meta_updated_by",
	MetaDeletedAt:        "meta_deleted_at",
	MetaDeletedBy:        "meta_deleted_by",
}

func NewRefundRequestsDBFieldNameFromStr(field string) (dbField RefundRequestsDBFieldNameType, found bool) {
	switch field {

	case string(RefundRequestsDBFieldName.Id):
		return RefundRequestsDBFieldName.Id, true

	case string(RefundRequestsDBFieldName.RefundCode):
		return RefundRequestsDBFieldName.RefundCode, true

	case string(RefundRequestsDBFieldName.PaymentRefId):
		return RefundRequestsDBFieldName.PaymentRefId, true

	case string(RefundRequestsDBFieldName.MerchantPartyId):
		return RefundRequestsDBFieldName.MerchantPartyId, true

	case string(RefundRequestsDBFieldName.CustomerPartyId):
		return RefundRequestsDBFieldName.CustomerPartyId, true

	case string(RefundRequestsDBFieldName.RefundPolicyId):
		return RefundRequestsDBFieldName.RefundPolicyId, true

	case string(RefundRequestsDBFieldName.CurrencyCode):
		return RefundRequestsDBFieldName.CurrencyCode, true

	case string(RefundRequestsDBFieldName.RequestedAmount):
		return RefundRequestsDBFieldName.RequestedAmount, true

	case string(RefundRequestsDBFieldName.ApprovedAmount):
		return RefundRequestsDBFieldName.ApprovedAmount, true

	case string(RefundRequestsDBFieldName.RefundReasonCode):
		return RefundRequestsDBFieldName.RefundReasonCode, true

	case string(RefundRequestsDBFieldName.RefundStatus):
		return RefundRequestsDBFieldName.RefundStatus, true

	case string(RefundRequestsDBFieldName.RequestedAt):
		return RefundRequestsDBFieldName.RequestedAt, true

	case string(RefundRequestsDBFieldName.ApprovedAt):
		return RefundRequestsDBFieldName.ApprovedAt, true

	case string(RefundRequestsDBFieldName.SettledFinanciallyAt):
		return RefundRequestsDBFieldName.SettledFinanciallyAt, true

	case string(RefundRequestsDBFieldName.Metadata):
		return RefundRequestsDBFieldName.Metadata, true

	case string(RefundRequestsDBFieldName.MetaCreatedAt):
		return RefundRequestsDBFieldName.MetaCreatedAt, true

	case string(RefundRequestsDBFieldName.MetaCreatedBy):
		return RefundRequestsDBFieldName.MetaCreatedBy, true

	case string(RefundRequestsDBFieldName.MetaUpdatedAt):
		return RefundRequestsDBFieldName.MetaUpdatedAt, true

	case string(RefundRequestsDBFieldName.MetaUpdatedBy):
		return RefundRequestsDBFieldName.MetaUpdatedBy, true

	case string(RefundRequestsDBFieldName.MetaDeletedAt):
		return RefundRequestsDBFieldName.MetaDeletedAt, true

	case string(RefundRequestsDBFieldName.MetaDeletedBy):
		return RefundRequestsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var RefundRequestsFilterJoins = map[string]JoinSpec{}

var RefundRequestsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refund_code": {
		SourcePath:        "refund_code",
		DefaultOutputPath: "refundCode",
		Column:            "refund_code",
		SQLAlias:          "refund_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_ref_id": {
		SourcePath:        "payment_ref_id",
		DefaultOutputPath: "paymentRefId",
		Column:            "payment_ref_id",
		SQLAlias:          "payment_ref_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_party_id": {
		SourcePath:        "merchant_party_id",
		DefaultOutputPath: "merchantPartyId",
		Column:            "merchant_party_id",
		SQLAlias:          "merchant_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"customer_party_id": {
		SourcePath:        "customer_party_id",
		DefaultOutputPath: "customerPartyId",
		Column:            "customer_party_id",
		SQLAlias:          "customer_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refund_policy_id": {
		SourcePath:        "refund_policy_id",
		DefaultOutputPath: "refundPolicyId",
		Column:            "refund_policy_id",
		SQLAlias:          "refund_policy_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"requested_amount": {
		SourcePath:        "requested_amount",
		DefaultOutputPath: "requestedAmount",
		Column:            "requested_amount",
		SQLAlias:          "requested_amount",
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
	"refund_reason_code": {
		SourcePath:        "refund_reason_code",
		DefaultOutputPath: "refundReasonCode",
		Column:            "refund_reason_code",
		SQLAlias:          "refund_reason_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"refund_status": {
		SourcePath:        "refund_status",
		DefaultOutputPath: "refundStatus",
		Column:            "refund_status",
		SQLAlias:          "refund_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"requested_at": {
		SourcePath:        "requested_at",
		DefaultOutputPath: "requestedAt",
		Column:            "requested_at",
		SQLAlias:          "requested_at",
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
	"settled_financially_at": {
		SourcePath:        "settled_financially_at",
		DefaultOutputPath: "settledFinanciallyAt",
		Column:            "settled_financially_at",
		SQLAlias:          "settled_financially_at",
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

func NewRefundRequestsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = RefundRequestsFilterFields[field]
	return
}

type RefundRequestsFilterResult struct {
	RefundRequests
	FilterCount int `db:"count"`
}

func ValidateRefundRequestsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewRefundRequestsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewRefundRequestsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewRefundRequestsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateRefundRequestsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateRefundRequestsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewRefundRequestsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateRefundRequestsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type RefundStatus string

const (
	RefundStatusRequested  RefundStatus = "requested"
	RefundStatusApproved   RefundStatus = "approved"
	RefundStatusProcessing RefundStatus = "processing"
	RefundStatusSucceeded  RefundStatus = "succeeded"
	RefundStatusFailed     RefundStatus = "failed"
	RefundStatusRejected   RefundStatus = "rejected"
	RefundStatusCancelled  RefundStatus = "cancelled"
)

type RefundRequests struct {
	Id                   uuid.UUID           `db:"id"`
	RefundCode           string              `db:"refund_code"`
	PaymentRefId         uuid.UUID           `db:"payment_ref_id"`
	MerchantPartyId      nuuid.NUUID         `db:"merchant_party_id"`
	CustomerPartyId      nuuid.NUUID         `db:"customer_party_id"`
	RefundPolicyId       nuuid.NUUID         `db:"refund_policy_id"`
	CurrencyCode         string              `db:"currency_code"`
	RequestedAmount      decimal.Decimal     `db:"requested_amount"`
	ApprovedAmount       decimal.NullDecimal `db:"approved_amount"`
	RefundReasonCode     string              `db:"refund_reason_code"`
	RefundStatus         RefundStatus        `db:"refund_status"`
	RequestedAt          time.Time           `db:"requested_at"`
	ApprovedAt           null.Time           `db:"approved_at"`
	SettledFinanciallyAt null.Time           `db:"settled_financially_at"`
	Metadata             json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type RefundRequestsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d RefundRequests) ToRefundRequestsPrimaryID() RefundRequestsPrimaryID {
	return RefundRequestsPrimaryID{
		Id: d.Id,
	}
}

type RefundRequestsList []*RefundRequests

type RefundRequestsFilterResultList []*RefundRequestsFilterResult
