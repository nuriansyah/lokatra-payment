package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type MerchantBalanceSnapshotsDBFieldNameType string

type merchantBalanceSnapshotsDBFieldName struct {
	Id               MerchantBalanceSnapshotsDBFieldNameType
	BalanceAccountId MerchantBalanceSnapshotsDBFieldNameType
	SnapshotAt       MerchantBalanceSnapshotsDBFieldNameType
	AvailableAmount  MerchantBalanceSnapshotsDBFieldNameType
	PendingAmount    MerchantBalanceSnapshotsDBFieldNameType
	ReservedAmount   MerchantBalanceSnapshotsDBFieldNameType
	DisputedAmount   MerchantBalanceSnapshotsDBFieldNameType
	NegativeAmount   MerchantBalanceSnapshotsDBFieldNameType
	Metadata         MerchantBalanceSnapshotsDBFieldNameType
	MetaCreatedAt    MerchantBalanceSnapshotsDBFieldNameType
	MetaCreatedBy    MerchantBalanceSnapshotsDBFieldNameType
	MetaUpdatedAt    MerchantBalanceSnapshotsDBFieldNameType
	MetaUpdatedBy    MerchantBalanceSnapshotsDBFieldNameType
	MetaDeletedAt    MerchantBalanceSnapshotsDBFieldNameType
	MetaDeletedBy    MerchantBalanceSnapshotsDBFieldNameType
}

var MerchantBalanceSnapshotsDBFieldName = merchantBalanceSnapshotsDBFieldName{
	Id:               "id",
	BalanceAccountId: "balance_account_id",
	SnapshotAt:       "snapshot_at",
	AvailableAmount:  "available_amount",
	PendingAmount:    "pending_amount",
	ReservedAmount:   "reserved_amount",
	DisputedAmount:   "disputed_amount",
	NegativeAmount:   "negative_amount",
	Metadata:         "metadata",
	MetaCreatedAt:    "meta_created_at",
	MetaCreatedBy:    "meta_created_by",
	MetaUpdatedAt:    "meta_updated_at",
	MetaUpdatedBy:    "meta_updated_by",
	MetaDeletedAt:    "meta_deleted_at",
	MetaDeletedBy:    "meta_deleted_by",
}

func NewMerchantBalanceSnapshotsDBFieldNameFromStr(field string) (dbField MerchantBalanceSnapshotsDBFieldNameType, found bool) {
	switch field {

	case string(MerchantBalanceSnapshotsDBFieldName.Id):
		return MerchantBalanceSnapshotsDBFieldName.Id, true

	case string(MerchantBalanceSnapshotsDBFieldName.BalanceAccountId):
		return MerchantBalanceSnapshotsDBFieldName.BalanceAccountId, true

	case string(MerchantBalanceSnapshotsDBFieldName.SnapshotAt):
		return MerchantBalanceSnapshotsDBFieldName.SnapshotAt, true

	case string(MerchantBalanceSnapshotsDBFieldName.AvailableAmount):
		return MerchantBalanceSnapshotsDBFieldName.AvailableAmount, true

	case string(MerchantBalanceSnapshotsDBFieldName.PendingAmount):
		return MerchantBalanceSnapshotsDBFieldName.PendingAmount, true

	case string(MerchantBalanceSnapshotsDBFieldName.ReservedAmount):
		return MerchantBalanceSnapshotsDBFieldName.ReservedAmount, true

	case string(MerchantBalanceSnapshotsDBFieldName.DisputedAmount):
		return MerchantBalanceSnapshotsDBFieldName.DisputedAmount, true

	case string(MerchantBalanceSnapshotsDBFieldName.NegativeAmount):
		return MerchantBalanceSnapshotsDBFieldName.NegativeAmount, true

	case string(MerchantBalanceSnapshotsDBFieldName.Metadata):
		return MerchantBalanceSnapshotsDBFieldName.Metadata, true

	case string(MerchantBalanceSnapshotsDBFieldName.MetaCreatedAt):
		return MerchantBalanceSnapshotsDBFieldName.MetaCreatedAt, true

	case string(MerchantBalanceSnapshotsDBFieldName.MetaCreatedBy):
		return MerchantBalanceSnapshotsDBFieldName.MetaCreatedBy, true

	case string(MerchantBalanceSnapshotsDBFieldName.MetaUpdatedAt):
		return MerchantBalanceSnapshotsDBFieldName.MetaUpdatedAt, true

	case string(MerchantBalanceSnapshotsDBFieldName.MetaUpdatedBy):
		return MerchantBalanceSnapshotsDBFieldName.MetaUpdatedBy, true

	case string(MerchantBalanceSnapshotsDBFieldName.MetaDeletedAt):
		return MerchantBalanceSnapshotsDBFieldName.MetaDeletedAt, true

	case string(MerchantBalanceSnapshotsDBFieldName.MetaDeletedBy):
		return MerchantBalanceSnapshotsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var MerchantBalanceSnapshotsFilterJoins = map[string]JoinSpec{}

var MerchantBalanceSnapshotsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"balance_account_id": {
		SourcePath:        "balance_account_id",
		DefaultOutputPath: "balanceAccountId",
		Column:            "balance_account_id",
		SQLAlias:          "balance_account_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"snapshot_at": {
		SourcePath:        "snapshot_at",
		DefaultOutputPath: "snapshotAt",
		Column:            "snapshot_at",
		SQLAlias:          "snapshot_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"available_amount": {
		SourcePath:        "available_amount",
		DefaultOutputPath: "availableAmount",
		Column:            "available_amount",
		SQLAlias:          "available_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"pending_amount": {
		SourcePath:        "pending_amount",
		DefaultOutputPath: "pendingAmount",
		Column:            "pending_amount",
		SQLAlias:          "pending_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reserved_amount": {
		SourcePath:        "reserved_amount",
		DefaultOutputPath: "reservedAmount",
		Column:            "reserved_amount",
		SQLAlias:          "reserved_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"disputed_amount": {
		SourcePath:        "disputed_amount",
		DefaultOutputPath: "disputedAmount",
		Column:            "disputed_amount",
		SQLAlias:          "disputed_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"negative_amount": {
		SourcePath:        "negative_amount",
		DefaultOutputPath: "negativeAmount",
		Column:            "negative_amount",
		SQLAlias:          "negative_amount",
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

func NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = MerchantBalanceSnapshotsFilterFields[field]
	return
}

type MerchantBalanceSnapshotsFilterResult struct {
	MerchantBalanceSnapshots
	FilterCount int `db:"count"`
}

func ValidateMerchantBalanceSnapshotsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateMerchantBalanceSnapshotsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateMerchantBalanceSnapshotsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateMerchantBalanceSnapshotsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type MerchantBalanceSnapshots struct {
	Id               uuid.UUID       `db:"id"`
	BalanceAccountId uuid.UUID       `db:"balance_account_id"`
	SnapshotAt       time.Time       `db:"snapshot_at"`
	AvailableAmount  decimal.Decimal `db:"available_amount"`
	PendingAmount    decimal.Decimal `db:"pending_amount"`
	ReservedAmount   decimal.Decimal `db:"reserved_amount"`
	DisputedAmount   decimal.Decimal `db:"disputed_amount"`
	NegativeAmount   decimal.Decimal `db:"negative_amount"`
	Metadata         json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type MerchantBalanceSnapshotsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d MerchantBalanceSnapshots) ToMerchantBalanceSnapshotsPrimaryID() MerchantBalanceSnapshotsPrimaryID {
	return MerchantBalanceSnapshotsPrimaryID{
		Id: d.Id,
	}
}

type MerchantBalanceSnapshotsList []*MerchantBalanceSnapshots

type MerchantBalanceSnapshotsFilterResultList []*MerchantBalanceSnapshotsFilterResult
