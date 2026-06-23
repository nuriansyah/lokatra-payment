package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

type SettlementEligibilitySnapshotsDBFieldNameType string

type settlementEligibilitySnapshotsDBFieldName struct {
	Id                        SettlementEligibilitySnapshotsDBFieldNameType
	SourceType                SettlementEligibilitySnapshotsDBFieldNameType
	SourceId                  SettlementEligibilitySnapshotsDBFieldNameType
	MerchantPartyId           SettlementEligibilitySnapshotsDBFieldNameType
	SettlementPolicyVersionId SettlementEligibilitySnapshotsDBFieldNameType
	CurrencyCode              SettlementEligibilitySnapshotsDBFieldNameType
	GrossAmount               SettlementEligibilitySnapshotsDBFieldNameType
	FeeAmount                 SettlementEligibilitySnapshotsDBFieldNameType
	TaxAmount                 SettlementEligibilitySnapshotsDBFieldNameType
	ReserveAmount             SettlementEligibilitySnapshotsDBFieldNameType
	NetSettleableAmount       SettlementEligibilitySnapshotsDBFieldNameType
	EligibilityStatus         SettlementEligibilitySnapshotsDBFieldNameType
	EligibleAt                SettlementEligibilitySnapshotsDBFieldNameType
	SnapshotPayload           SettlementEligibilitySnapshotsDBFieldNameType
	MetaCreatedAt             SettlementEligibilitySnapshotsDBFieldNameType
	MetaCreatedBy             SettlementEligibilitySnapshotsDBFieldNameType
	MetaUpdatedAt             SettlementEligibilitySnapshotsDBFieldNameType
	MetaUpdatedBy             SettlementEligibilitySnapshotsDBFieldNameType
	MetaDeletedAt             SettlementEligibilitySnapshotsDBFieldNameType
	MetaDeletedBy             SettlementEligibilitySnapshotsDBFieldNameType
}

var SettlementEligibilitySnapshotsDBFieldName = settlementEligibilitySnapshotsDBFieldName{
	Id:                        "id",
	SourceType:                "source_type",
	SourceId:                  "source_id",
	MerchantPartyId:           "merchant_party_id",
	SettlementPolicyVersionId: "settlement_policy_version_id",
	CurrencyCode:              "currency_code",
	GrossAmount:               "gross_amount",
	FeeAmount:                 "fee_amount",
	TaxAmount:                 "tax_amount",
	ReserveAmount:             "reserve_amount",
	NetSettleableAmount:       "net_settleable_amount",
	EligibilityStatus:         "eligibility_status",
	EligibleAt:                "eligible_at",
	SnapshotPayload:           "snapshot_payload",
	MetaCreatedAt:             "meta_created_at",
	MetaCreatedBy:             "meta_created_by",
	MetaUpdatedAt:             "meta_updated_at",
	MetaUpdatedBy:             "meta_updated_by",
	MetaDeletedAt:             "meta_deleted_at",
	MetaDeletedBy:             "meta_deleted_by",
}

func NewSettlementEligibilitySnapshotsDBFieldNameFromStr(field string) (dbField SettlementEligibilitySnapshotsDBFieldNameType, found bool) {
	switch field {

	case string(SettlementEligibilitySnapshotsDBFieldName.Id):
		return SettlementEligibilitySnapshotsDBFieldName.Id, true

	case string(SettlementEligibilitySnapshotsDBFieldName.SourceType):
		return SettlementEligibilitySnapshotsDBFieldName.SourceType, true

	case string(SettlementEligibilitySnapshotsDBFieldName.SourceId):
		return SettlementEligibilitySnapshotsDBFieldName.SourceId, true

	case string(SettlementEligibilitySnapshotsDBFieldName.MerchantPartyId):
		return SettlementEligibilitySnapshotsDBFieldName.MerchantPartyId, true

	case string(SettlementEligibilitySnapshotsDBFieldName.SettlementPolicyVersionId):
		return SettlementEligibilitySnapshotsDBFieldName.SettlementPolicyVersionId, true

	case string(SettlementEligibilitySnapshotsDBFieldName.CurrencyCode):
		return SettlementEligibilitySnapshotsDBFieldName.CurrencyCode, true

	case string(SettlementEligibilitySnapshotsDBFieldName.GrossAmount):
		return SettlementEligibilitySnapshotsDBFieldName.GrossAmount, true

	case string(SettlementEligibilitySnapshotsDBFieldName.FeeAmount):
		return SettlementEligibilitySnapshotsDBFieldName.FeeAmount, true

	case string(SettlementEligibilitySnapshotsDBFieldName.TaxAmount):
		return SettlementEligibilitySnapshotsDBFieldName.TaxAmount, true

	case string(SettlementEligibilitySnapshotsDBFieldName.ReserveAmount):
		return SettlementEligibilitySnapshotsDBFieldName.ReserveAmount, true

	case string(SettlementEligibilitySnapshotsDBFieldName.NetSettleableAmount):
		return SettlementEligibilitySnapshotsDBFieldName.NetSettleableAmount, true

	case string(SettlementEligibilitySnapshotsDBFieldName.EligibilityStatus):
		return SettlementEligibilitySnapshotsDBFieldName.EligibilityStatus, true

	case string(SettlementEligibilitySnapshotsDBFieldName.EligibleAt):
		return SettlementEligibilitySnapshotsDBFieldName.EligibleAt, true

	case string(SettlementEligibilitySnapshotsDBFieldName.SnapshotPayload):
		return SettlementEligibilitySnapshotsDBFieldName.SnapshotPayload, true

	case string(SettlementEligibilitySnapshotsDBFieldName.MetaCreatedAt):
		return SettlementEligibilitySnapshotsDBFieldName.MetaCreatedAt, true

	case string(SettlementEligibilitySnapshotsDBFieldName.MetaCreatedBy):
		return SettlementEligibilitySnapshotsDBFieldName.MetaCreatedBy, true

	case string(SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedAt):
		return SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedAt, true

	case string(SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedBy):
		return SettlementEligibilitySnapshotsDBFieldName.MetaUpdatedBy, true

	case string(SettlementEligibilitySnapshotsDBFieldName.MetaDeletedAt):
		return SettlementEligibilitySnapshotsDBFieldName.MetaDeletedAt, true

	case string(SettlementEligibilitySnapshotsDBFieldName.MetaDeletedBy):
		return SettlementEligibilitySnapshotsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var SettlementEligibilitySnapshotsFilterJoins = map[string]JoinSpec{}

var SettlementEligibilitySnapshotsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_type": {
		SourcePath:        "source_type",
		DefaultOutputPath: "sourceType",
		Column:            "source_type",
		SQLAlias:          "source_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"source_id": {
		SourcePath:        "source_id",
		DefaultOutputPath: "sourceId",
		Column:            "source_id",
		SQLAlias:          "source_id",
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
	"settlement_policy_version_id": {
		SourcePath:        "settlement_policy_version_id",
		DefaultOutputPath: "settlementPolicyVersionId",
		Column:            "settlement_policy_version_id",
		SQLAlias:          "settlement_policy_version_id",
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
	"gross_amount": {
		SourcePath:        "gross_amount",
		DefaultOutputPath: "grossAmount",
		Column:            "gross_amount",
		SQLAlias:          "gross_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"fee_amount": {
		SourcePath:        "fee_amount",
		DefaultOutputPath: "feeAmount",
		Column:            "fee_amount",
		SQLAlias:          "fee_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"tax_amount": {
		SourcePath:        "tax_amount",
		DefaultOutputPath: "taxAmount",
		Column:            "tax_amount",
		SQLAlias:          "tax_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reserve_amount": {
		SourcePath:        "reserve_amount",
		DefaultOutputPath: "reserveAmount",
		Column:            "reserve_amount",
		SQLAlias:          "reserve_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"net_settleable_amount": {
		SourcePath:        "net_settleable_amount",
		DefaultOutputPath: "netSettleableAmount",
		Column:            "net_settleable_amount",
		SQLAlias:          "net_settleable_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"eligibility_status": {
		SourcePath:        "eligibility_status",
		DefaultOutputPath: "eligibilityStatus",
		Column:            "eligibility_status",
		SQLAlias:          "eligibility_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"eligible_at": {
		SourcePath:        "eligible_at",
		DefaultOutputPath: "eligibleAt",
		Column:            "eligible_at",
		SQLAlias:          "eligible_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"snapshot_payload": {
		SourcePath:        "snapshot_payload",
		DefaultOutputPath: "snapshotPayload",
		Column:            "snapshot_payload",
		SQLAlias:          "snapshot_payload",
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

func NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = SettlementEligibilitySnapshotsFilterFields[field]
	return
}

type SettlementEligibilitySnapshotsFilterResult struct {
	SettlementEligibilitySnapshots
	FilterCount int `db:"count"`
}

func ValidateSettlementEligibilitySnapshotsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateSettlementEligibilitySnapshotsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateSettlementEligibilitySnapshotsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateSettlementEligibilitySnapshotsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type EligibilityStatus string

const (
	EligibilityStatusPending  EligibilityStatus = "pending"
	EligibilityStatusEligible EligibilityStatus = "eligible"
	EligibilityStatusHeld     EligibilityStatus = "held"
	EligibilityStatusBlocked  EligibilityStatus = "blocked"
	EligibilityStatusSettled  EligibilityStatus = "settled"
)

type SettlementEligibilitySnapshots struct {
	Id                        uuid.UUID         `db:"id"`
	SourceType                string            `db:"source_type"`
	SourceId                  uuid.UUID         `db:"source_id"`
	MerchantPartyId           uuid.UUID         `db:"merchant_party_id"`
	SettlementPolicyVersionId uuid.UUID         `db:"settlement_policy_version_id"`
	CurrencyCode              string            `db:"currency_code"`
	GrossAmount               decimal.Decimal   `db:"gross_amount"`
	FeeAmount                 decimal.Decimal   `db:"fee_amount"`
	TaxAmount                 decimal.Decimal   `db:"tax_amount"`
	ReserveAmount             decimal.Decimal   `db:"reserve_amount"`
	NetSettleableAmount       decimal.Decimal   `db:"net_settleable_amount"`
	EligibilityStatus         EligibilityStatus `db:"eligibility_status"`
	EligibleAt                null.Time         `db:"eligible_at"`
	SnapshotPayload           json.RawMessage   `db:"snapshot_payload"`

	shared.MetaSignature
}
type SettlementEligibilitySnapshotsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d SettlementEligibilitySnapshots) ToSettlementEligibilitySnapshotsPrimaryID() SettlementEligibilitySnapshotsPrimaryID {
	return SettlementEligibilitySnapshotsPrimaryID{
		Id: d.Id,
	}
}

type SettlementEligibilitySnapshotsList []*SettlementEligibilitySnapshots

type SettlementEligibilitySnapshotsFilterResultList []*SettlementEligibilitySnapshotsFilterResult
