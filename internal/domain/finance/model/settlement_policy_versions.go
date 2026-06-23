package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type SettlementPolicyVersionsDBFieldNameType string

type settlementPolicyVersionsDBFieldName struct {
	Id                 SettlementPolicyVersionsDBFieldNameType
	SettlementPolicyId SettlementPolicyVersionsDBFieldNameType
	VersionNo          SettlementPolicyVersionsDBFieldNameType
	TriggerType        SettlementPolicyVersionsDBFieldNameType
	DelayDays          SettlementPolicyVersionsDBFieldNameType
	MinPayoutAmount    SettlementPolicyVersionsDBFieldNameType
	PayoutFrequency    SettlementPolicyVersionsDBFieldNameType
	ReservePolicyId    SettlementPolicyVersionsDBFieldNameType
	AutoRelease        SettlementPolicyVersionsDBFieldNameType
	Conditions         SettlementPolicyVersionsDBFieldNameType
	IsCurrent          SettlementPolicyVersionsDBFieldNameType
	MetaCreatedAt      SettlementPolicyVersionsDBFieldNameType
	MetaCreatedBy      SettlementPolicyVersionsDBFieldNameType
	MetaUpdatedAt      SettlementPolicyVersionsDBFieldNameType
	MetaUpdatedBy      SettlementPolicyVersionsDBFieldNameType
	MetaDeletedAt      SettlementPolicyVersionsDBFieldNameType
	MetaDeletedBy      SettlementPolicyVersionsDBFieldNameType
}

var SettlementPolicyVersionsDBFieldName = settlementPolicyVersionsDBFieldName{
	Id:                 "id",
	SettlementPolicyId: "settlement_policy_id",
	VersionNo:          "version_no",
	TriggerType:        "trigger_type",
	DelayDays:          "delay_days",
	MinPayoutAmount:    "min_payout_amount",
	PayoutFrequency:    "payout_frequency",
	ReservePolicyId:    "reserve_policy_id",
	AutoRelease:        "auto_release",
	Conditions:         "conditions",
	IsCurrent:          "is_current",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewSettlementPolicyVersionsDBFieldNameFromStr(field string) (dbField SettlementPolicyVersionsDBFieldNameType, found bool) {
	switch field {

	case string(SettlementPolicyVersionsDBFieldName.Id):
		return SettlementPolicyVersionsDBFieldName.Id, true

	case string(SettlementPolicyVersionsDBFieldName.SettlementPolicyId):
		return SettlementPolicyVersionsDBFieldName.SettlementPolicyId, true

	case string(SettlementPolicyVersionsDBFieldName.VersionNo):
		return SettlementPolicyVersionsDBFieldName.VersionNo, true

	case string(SettlementPolicyVersionsDBFieldName.TriggerType):
		return SettlementPolicyVersionsDBFieldName.TriggerType, true

	case string(SettlementPolicyVersionsDBFieldName.DelayDays):
		return SettlementPolicyVersionsDBFieldName.DelayDays, true

	case string(SettlementPolicyVersionsDBFieldName.MinPayoutAmount):
		return SettlementPolicyVersionsDBFieldName.MinPayoutAmount, true

	case string(SettlementPolicyVersionsDBFieldName.PayoutFrequency):
		return SettlementPolicyVersionsDBFieldName.PayoutFrequency, true

	case string(SettlementPolicyVersionsDBFieldName.ReservePolicyId):
		return SettlementPolicyVersionsDBFieldName.ReservePolicyId, true

	case string(SettlementPolicyVersionsDBFieldName.AutoRelease):
		return SettlementPolicyVersionsDBFieldName.AutoRelease, true

	case string(SettlementPolicyVersionsDBFieldName.Conditions):
		return SettlementPolicyVersionsDBFieldName.Conditions, true

	case string(SettlementPolicyVersionsDBFieldName.IsCurrent):
		return SettlementPolicyVersionsDBFieldName.IsCurrent, true

	case string(SettlementPolicyVersionsDBFieldName.MetaCreatedAt):
		return SettlementPolicyVersionsDBFieldName.MetaCreatedAt, true

	case string(SettlementPolicyVersionsDBFieldName.MetaCreatedBy):
		return SettlementPolicyVersionsDBFieldName.MetaCreatedBy, true

	case string(SettlementPolicyVersionsDBFieldName.MetaUpdatedAt):
		return SettlementPolicyVersionsDBFieldName.MetaUpdatedAt, true

	case string(SettlementPolicyVersionsDBFieldName.MetaUpdatedBy):
		return SettlementPolicyVersionsDBFieldName.MetaUpdatedBy, true

	case string(SettlementPolicyVersionsDBFieldName.MetaDeletedAt):
		return SettlementPolicyVersionsDBFieldName.MetaDeletedAt, true

	case string(SettlementPolicyVersionsDBFieldName.MetaDeletedBy):
		return SettlementPolicyVersionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var SettlementPolicyVersionsFilterJoins = map[string]JoinSpec{}

var SettlementPolicyVersionsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"settlement_policy_id": {
		SourcePath:        "settlement_policy_id",
		DefaultOutputPath: "settlementPolicyId",
		Column:            "settlement_policy_id",
		SQLAlias:          "settlement_policy_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"version_no": {
		SourcePath:        "version_no",
		DefaultOutputPath: "versionNo",
		Column:            "version_no",
		SQLAlias:          "version_no",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"trigger_type": {
		SourcePath:        "trigger_type",
		DefaultOutputPath: "triggerType",
		Column:            "trigger_type",
		SQLAlias:          "trigger_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"delay_days": {
		SourcePath:        "delay_days",
		DefaultOutputPath: "delayDays",
		Column:            "delay_days",
		SQLAlias:          "delay_days",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"min_payout_amount": {
		SourcePath:        "min_payout_amount",
		DefaultOutputPath: "minPayoutAmount",
		Column:            "min_payout_amount",
		SQLAlias:          "min_payout_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payout_frequency": {
		SourcePath:        "payout_frequency",
		DefaultOutputPath: "payoutFrequency",
		Column:            "payout_frequency",
		SQLAlias:          "payout_frequency",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reserve_policy_id": {
		SourcePath:        "reserve_policy_id",
		DefaultOutputPath: "reservePolicyId",
		Column:            "reserve_policy_id",
		SQLAlias:          "reserve_policy_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"auto_release": {
		SourcePath:        "auto_release",
		DefaultOutputPath: "autoRelease",
		Column:            "auto_release",
		SQLAlias:          "auto_release",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"conditions": {
		SourcePath:        "conditions",
		DefaultOutputPath: "conditions",
		Column:            "conditions",
		SQLAlias:          "conditions",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_current": {
		SourcePath:        "is_current",
		DefaultOutputPath: "isCurrent",
		Column:            "is_current",
		SQLAlias:          "is_current",
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

func NewSettlementPolicyVersionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = SettlementPolicyVersionsFilterFields[field]
	return
}

type SettlementPolicyVersionsFilterResult struct {
	SettlementPolicyVersions
	FilterCount int `db:"count"`
}

func ValidateSettlementPolicyVersionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewSettlementPolicyVersionsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewSettlementPolicyVersionsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewSettlementPolicyVersionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateSettlementPolicyVersionsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateSettlementPolicyVersionsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewSettlementPolicyVersionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateSettlementPolicyVersionsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PayoutFrequency string

const (
	PayoutFrequencyDaily    PayoutFrequency = "daily"
	PayoutFrequencyWeekly   PayoutFrequency = "weekly"
	PayoutFrequencyBiweekly PayoutFrequency = "biweekly"
	PayoutFrequencyMonthly  PayoutFrequency = "monthly"
	PayoutFrequencyManual   PayoutFrequency = "manual"
)

type TriggerType string

const (
	TriggerTypePaymentSettled   TriggerType = "payment_settled"
	TriggerTypeVisitCompleted   TriggerType = "visit_completed"
	TriggerTypeServiceStarted   TriggerType = "service_started"
	TriggerTypeServiceCompleted TriggerType = "service_completed"
	TriggerTypeManual           TriggerType = "manual"
)

type SettlementPolicyVersions struct {
	Id                 uuid.UUID       `db:"id"`
	SettlementPolicyId uuid.UUID       `db:"settlement_policy_id"`
	VersionNo          int             `db:"version_no"`
	TriggerType        TriggerType     `db:"trigger_type"`
	DelayDays          int             `db:"delay_days"`
	MinPayoutAmount    decimal.Decimal `db:"min_payout_amount"`
	PayoutFrequency    PayoutFrequency `db:"payout_frequency"`
	ReservePolicyId    nuuid.NUUID     `db:"reserve_policy_id"`
	AutoRelease        bool            `db:"auto_release"`
	Conditions         json.RawMessage `db:"conditions"`
	IsCurrent          bool            `db:"is_current"`

	shared.MetaSignature
}
type SettlementPolicyVersionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d SettlementPolicyVersions) ToSettlementPolicyVersionsPrimaryID() SettlementPolicyVersionsPrimaryID {
	return SettlementPolicyVersionsPrimaryID{
		Id: d.Id,
	}
}

type SettlementPolicyVersionsList []*SettlementPolicyVersions

type SettlementPolicyVersionsFilterResultList []*SettlementPolicyVersionsFilterResult
