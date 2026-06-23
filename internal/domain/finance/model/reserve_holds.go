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

type ReserveHoldsDBFieldNameType string

type reserveHoldsDBFieldName struct {
	Id              ReserveHoldsDBFieldNameType
	MerchantPartyId ReserveHoldsDBFieldNameType
	CurrencyCode    ReserveHoldsDBFieldNameType
	SourceType      ReserveHoldsDBFieldNameType
	SourceId        ReserveHoldsDBFieldNameType
	HoldType        ReserveHoldsDBFieldNameType
	HoldStatus      ReserveHoldsDBFieldNameType
	HoldAmount      ReserveHoldsDBFieldNameType
	ReleasedAmount  ReserveHoldsDBFieldNameType
	EffectiveAt     ReserveHoldsDBFieldNameType
	ReleaseAt       ReserveHoldsDBFieldNameType
	ReleasedAt      ReserveHoldsDBFieldNameType
	ReasonCode      ReserveHoldsDBFieldNameType
	ReasonDetail    ReserveHoldsDBFieldNameType
	Metadata        ReserveHoldsDBFieldNameType
	MetaCreatedAt   ReserveHoldsDBFieldNameType
	MetaCreatedBy   ReserveHoldsDBFieldNameType
	MetaUpdatedAt   ReserveHoldsDBFieldNameType
	MetaUpdatedBy   ReserveHoldsDBFieldNameType
	MetaDeletedAt   ReserveHoldsDBFieldNameType
	MetaDeletedBy   ReserveHoldsDBFieldNameType
}

var ReserveHoldsDBFieldName = reserveHoldsDBFieldName{
	Id:              "id",
	MerchantPartyId: "merchant_party_id",
	CurrencyCode:    "currency_code",
	SourceType:      "source_type",
	SourceId:        "source_id",
	HoldType:        "hold_type",
	HoldStatus:      "hold_status",
	HoldAmount:      "hold_amount",
	ReleasedAmount:  "released_amount",
	EffectiveAt:     "effective_at",
	ReleaseAt:       "release_at",
	ReleasedAt:      "released_at",
	ReasonCode:      "reason_code",
	ReasonDetail:    "reason_detail",
	Metadata:        "metadata",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewReserveHoldsDBFieldNameFromStr(field string) (dbField ReserveHoldsDBFieldNameType, found bool) {
	switch field {

	case string(ReserveHoldsDBFieldName.Id):
		return ReserveHoldsDBFieldName.Id, true

	case string(ReserveHoldsDBFieldName.MerchantPartyId):
		return ReserveHoldsDBFieldName.MerchantPartyId, true

	case string(ReserveHoldsDBFieldName.CurrencyCode):
		return ReserveHoldsDBFieldName.CurrencyCode, true

	case string(ReserveHoldsDBFieldName.SourceType):
		return ReserveHoldsDBFieldName.SourceType, true

	case string(ReserveHoldsDBFieldName.SourceId):
		return ReserveHoldsDBFieldName.SourceId, true

	case string(ReserveHoldsDBFieldName.HoldType):
		return ReserveHoldsDBFieldName.HoldType, true

	case string(ReserveHoldsDBFieldName.HoldStatus):
		return ReserveHoldsDBFieldName.HoldStatus, true

	case string(ReserveHoldsDBFieldName.HoldAmount):
		return ReserveHoldsDBFieldName.HoldAmount, true

	case string(ReserveHoldsDBFieldName.ReleasedAmount):
		return ReserveHoldsDBFieldName.ReleasedAmount, true

	case string(ReserveHoldsDBFieldName.EffectiveAt):
		return ReserveHoldsDBFieldName.EffectiveAt, true

	case string(ReserveHoldsDBFieldName.ReleaseAt):
		return ReserveHoldsDBFieldName.ReleaseAt, true

	case string(ReserveHoldsDBFieldName.ReleasedAt):
		return ReserveHoldsDBFieldName.ReleasedAt, true

	case string(ReserveHoldsDBFieldName.ReasonCode):
		return ReserveHoldsDBFieldName.ReasonCode, true

	case string(ReserveHoldsDBFieldName.ReasonDetail):
		return ReserveHoldsDBFieldName.ReasonDetail, true

	case string(ReserveHoldsDBFieldName.Metadata):
		return ReserveHoldsDBFieldName.Metadata, true

	case string(ReserveHoldsDBFieldName.MetaCreatedAt):
		return ReserveHoldsDBFieldName.MetaCreatedAt, true

	case string(ReserveHoldsDBFieldName.MetaCreatedBy):
		return ReserveHoldsDBFieldName.MetaCreatedBy, true

	case string(ReserveHoldsDBFieldName.MetaUpdatedAt):
		return ReserveHoldsDBFieldName.MetaUpdatedAt, true

	case string(ReserveHoldsDBFieldName.MetaUpdatedBy):
		return ReserveHoldsDBFieldName.MetaUpdatedBy, true

	case string(ReserveHoldsDBFieldName.MetaDeletedAt):
		return ReserveHoldsDBFieldName.MetaDeletedAt, true

	case string(ReserveHoldsDBFieldName.MetaDeletedBy):
		return ReserveHoldsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ReserveHoldsFilterJoins = map[string]JoinSpec{}

var ReserveHoldsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
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
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
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
	"hold_type": {
		SourcePath:        "hold_type",
		DefaultOutputPath: "holdType",
		Column:            "hold_type",
		SQLAlias:          "hold_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"hold_status": {
		SourcePath:        "hold_status",
		DefaultOutputPath: "holdStatus",
		Column:            "hold_status",
		SQLAlias:          "hold_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"hold_amount": {
		SourcePath:        "hold_amount",
		DefaultOutputPath: "holdAmount",
		Column:            "hold_amount",
		SQLAlias:          "hold_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"released_amount": {
		SourcePath:        "released_amount",
		DefaultOutputPath: "releasedAmount",
		Column:            "released_amount",
		SQLAlias:          "released_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"effective_at": {
		SourcePath:        "effective_at",
		DefaultOutputPath: "effectiveAt",
		Column:            "effective_at",
		SQLAlias:          "effective_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"release_at": {
		SourcePath:        "release_at",
		DefaultOutputPath: "releaseAt",
		Column:            "release_at",
		SQLAlias:          "release_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"released_at": {
		SourcePath:        "released_at",
		DefaultOutputPath: "releasedAt",
		Column:            "released_at",
		SQLAlias:          "released_at",
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

func NewReserveHoldsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ReserveHoldsFilterFields[field]
	return
}

type ReserveHoldsFilterResult struct {
	ReserveHolds
	FilterCount int `db:"count"`
}

func ValidateReserveHoldsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewReserveHoldsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewReserveHoldsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewReserveHoldsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateReserveHoldsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateReserveHoldsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewReserveHoldsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateReserveHoldsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type HoldStatus string

const (
	HoldStatusActive   HoldStatus = "active"
	HoldStatusReleased HoldStatus = "released"
	HoldStatusConsumed HoldStatus = "consumed"
	HoldStatusCanceled HoldStatus = "canceled"
)

type ReserveHolds struct {
	Id              uuid.UUID       `db:"id"`
	MerchantPartyId uuid.UUID       `db:"merchant_party_id"`
	CurrencyCode    string          `db:"currency_code"`
	SourceType      string          `db:"source_type"`
	SourceId        uuid.UUID       `db:"source_id"`
	HoldType        string          `db:"hold_type"`
	HoldStatus      HoldStatus      `db:"hold_status"`
	HoldAmount      decimal.Decimal `db:"hold_amount"`
	ReleasedAmount  decimal.Decimal `db:"released_amount"`
	EffectiveAt     time.Time       `db:"effective_at"`
	ReleaseAt       null.Time       `db:"release_at"`
	ReleasedAt      null.Time       `db:"released_at"`
	ReasonCode      string          `db:"reason_code"`
	ReasonDetail    null.String     `db:"reason_detail"`
	Metadata        json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ReserveHoldsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ReserveHolds) ToReserveHoldsPrimaryID() ReserveHoldsPrimaryID {
	return ReserveHoldsPrimaryID{
		Id: d.Id,
	}
}

type ReserveHoldsList []*ReserveHolds

type ReserveHoldsFilterResultList []*ReserveHoldsFilterResult
