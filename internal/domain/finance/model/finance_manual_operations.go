package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type FinanceManualOperationsDBFieldNameType string

type financeManualOperationsDBFieldName struct {
	Id              FinanceManualOperationsDBFieldNameType
	OperationCode   FinanceManualOperationsDBFieldNameType
	OperationType   FinanceManualOperationsDBFieldNameType
	TargetRefType   FinanceManualOperationsDBFieldNameType
	TargetRefId     FinanceManualOperationsDBFieldNameType
	RequestedBy     FinanceManualOperationsDBFieldNameType
	OperationStatus FinanceManualOperationsDBFieldNameType
	ReasonCode      FinanceManualOperationsDBFieldNameType
	ReasonDetail    FinanceManualOperationsDBFieldNameType
	Payload         FinanceManualOperationsDBFieldNameType
	ExecutedAt      FinanceManualOperationsDBFieldNameType
	Metadata        FinanceManualOperationsDBFieldNameType
	MetaCreatedAt   FinanceManualOperationsDBFieldNameType
	MetaCreatedBy   FinanceManualOperationsDBFieldNameType
	MetaUpdatedAt   FinanceManualOperationsDBFieldNameType
	MetaUpdatedBy   FinanceManualOperationsDBFieldNameType
	MetaDeletedAt   FinanceManualOperationsDBFieldNameType
	MetaDeletedBy   FinanceManualOperationsDBFieldNameType
}

var FinanceManualOperationsDBFieldName = financeManualOperationsDBFieldName{
	Id:              "id",
	OperationCode:   "operation_code",
	OperationType:   "operation_type",
	TargetRefType:   "target_ref_type",
	TargetRefId:     "target_ref_id",
	RequestedBy:     "requested_by",
	OperationStatus: "operation_status",
	ReasonCode:      "reason_code",
	ReasonDetail:    "reason_detail",
	Payload:         "payload",
	ExecutedAt:      "executed_at",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewFinanceManualOperationsDBFieldNameFromStr(field string) (dbField FinanceManualOperationsDBFieldNameType, found bool) {
	switch field {

	case string(FinanceManualOperationsDBFieldName.Id):
		return FinanceManualOperationsDBFieldName.Id, true

	case string(FinanceManualOperationsDBFieldName.OperationCode):
		return FinanceManualOperationsDBFieldName.OperationCode, true

	case string(FinanceManualOperationsDBFieldName.OperationType):
		return FinanceManualOperationsDBFieldName.OperationType, true

	case string(FinanceManualOperationsDBFieldName.TargetRefType):
		return FinanceManualOperationsDBFieldName.TargetRefType, true

	case string(FinanceManualOperationsDBFieldName.TargetRefId):
		return FinanceManualOperationsDBFieldName.TargetRefId, true

	case string(FinanceManualOperationsDBFieldName.RequestedBy):
		return FinanceManualOperationsDBFieldName.RequestedBy, true

	case string(FinanceManualOperationsDBFieldName.OperationStatus):
		return FinanceManualOperationsDBFieldName.OperationStatus, true

	case string(FinanceManualOperationsDBFieldName.ReasonCode):
		return FinanceManualOperationsDBFieldName.ReasonCode, true

	case string(FinanceManualOperationsDBFieldName.ReasonDetail):
		return FinanceManualOperationsDBFieldName.ReasonDetail, true

	case string(FinanceManualOperationsDBFieldName.Payload):
		return FinanceManualOperationsDBFieldName.Payload, true

	case string(FinanceManualOperationsDBFieldName.ExecutedAt):
		return FinanceManualOperationsDBFieldName.ExecutedAt, true

	case string(FinanceManualOperationsDBFieldName.Metadata):
		return FinanceManualOperationsDBFieldName.Metadata, true

	case string(FinanceManualOperationsDBFieldName.MetaCreatedAt):
		return FinanceManualOperationsDBFieldName.MetaCreatedAt, true

	case string(FinanceManualOperationsDBFieldName.MetaCreatedBy):
		return FinanceManualOperationsDBFieldName.MetaCreatedBy, true

	case string(FinanceManualOperationsDBFieldName.MetaUpdatedAt):
		return FinanceManualOperationsDBFieldName.MetaUpdatedAt, true

	case string(FinanceManualOperationsDBFieldName.MetaUpdatedBy):
		return FinanceManualOperationsDBFieldName.MetaUpdatedBy, true

	case string(FinanceManualOperationsDBFieldName.MetaDeletedAt):
		return FinanceManualOperationsDBFieldName.MetaDeletedAt, true

	case string(FinanceManualOperationsDBFieldName.MetaDeletedBy):
		return FinanceManualOperationsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var FinanceManualOperationsFilterJoins = map[string]JoinSpec{}

var FinanceManualOperationsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"operation_code": {
		SourcePath:        "operation_code",
		DefaultOutputPath: "operationCode",
		Column:            "operation_code",
		SQLAlias:          "operation_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"operation_type": {
		SourcePath:        "operation_type",
		DefaultOutputPath: "operationType",
		Column:            "operation_type",
		SQLAlias:          "operation_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"target_ref_type": {
		SourcePath:        "target_ref_type",
		DefaultOutputPath: "targetRefType",
		Column:            "target_ref_type",
		SQLAlias:          "target_ref_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"target_ref_id": {
		SourcePath:        "target_ref_id",
		DefaultOutputPath: "targetRefId",
		Column:            "target_ref_id",
		SQLAlias:          "target_ref_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"requested_by": {
		SourcePath:        "requested_by",
		DefaultOutputPath: "requestedBy",
		Column:            "requested_by",
		SQLAlias:          "requested_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"operation_status": {
		SourcePath:        "operation_status",
		DefaultOutputPath: "operationStatus",
		Column:            "operation_status",
		SQLAlias:          "operation_status",
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
	"payload": {
		SourcePath:        "payload",
		DefaultOutputPath: "payload",
		Column:            "payload",
		SQLAlias:          "payload",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"executed_at": {
		SourcePath:        "executed_at",
		DefaultOutputPath: "executedAt",
		Column:            "executed_at",
		SQLAlias:          "executed_at",
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

func NewFinanceManualOperationsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceManualOperationsFilterFields[field]
	return
}

type FinanceManualOperationsFilterResult struct {
	FinanceManualOperations
	FilterCount int `db:"count"`
}

func ValidateFinanceManualOperationsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceManualOperationsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceManualOperationsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceManualOperationsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceManualOperationsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceManualOperationsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceManualOperationsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceManualOperationsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type OperationStatus string

const (
	OperationStatusRequested OperationStatus = "requested"
	OperationStatusApproved  OperationStatus = "approved"
	OperationStatusRejected  OperationStatus = "rejected"
	OperationStatusExecuted  OperationStatus = "executed"
	OperationStatusCancelled OperationStatus = "cancelled"
)

type OperationType string

const (
	OperationTypeManualAdjustment               OperationType = "manual_adjustment"
	OperationTypeManualRefund                   OperationType = "manual_refund"
	OperationTypeManualReleaseHold              OperationType = "manual_release_hold"
	OperationTypeManualPayoutRetry              OperationType = "manual_payout_retry"
	OperationTypeManualWriteoff                 OperationType = "manual_writeoff"
	OperationTypeManualReconciliationResolution OperationType = "manual_reconciliation_resolution"
)

type FinanceManualOperations struct {
	Id              uuid.UUID       `db:"id"`
	OperationCode   string          `db:"operation_code"`
	OperationType   OperationType   `db:"operation_type"`
	TargetRefType   string          `db:"target_ref_type"`
	TargetRefId     uuid.UUID       `db:"target_ref_id"`
	RequestedBy     uuid.UUID       `db:"requested_by"`
	OperationStatus OperationStatus `db:"operation_status"`
	ReasonCode      string          `db:"reason_code"`
	ReasonDetail    null.String     `db:"reason_detail"`
	Payload         json.RawMessage `db:"payload"`
	ExecutedAt      null.Time       `db:"executed_at"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type FinanceManualOperationsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d FinanceManualOperations) ToFinanceManualOperationsPrimaryID() FinanceManualOperationsPrimaryID {
	return FinanceManualOperationsPrimaryID{
		Id: d.Id,
	}
}

type FinanceManualOperationsList []*FinanceManualOperations

type FinanceManualOperationsFilterResultList []*FinanceManualOperationsFilterResult
