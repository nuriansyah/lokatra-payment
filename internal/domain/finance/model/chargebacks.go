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

type ChargebacksDBFieldNameType string

type chargebacksDBFieldName struct {
	Id                ChargebacksDBFieldNameType
	ChargebackCode    ChargebacksDBFieldNameType
	PaymentRefId      ChargebacksDBFieldNameType
	MerchantPartyId   ChargebacksDBFieldNameType
	ProviderAccountId ChargebacksDBFieldNameType
	CurrencyCode      ChargebacksDBFieldNameType
	DisputedAmount    ChargebacksDBFieldNameType
	ChargebackStatus  ChargebacksDBFieldNameType
	ReasonCode        ChargebacksDBFieldNameType
	OpenedAt          ChargebacksDBFieldNameType
	ClosedAt          ChargebacksDBFieldNameType
	DueAt             ChargebacksDBFieldNameType
	Metadata          ChargebacksDBFieldNameType
	MetaCreatedAt     ChargebacksDBFieldNameType
	MetaCreatedBy     ChargebacksDBFieldNameType
	MetaUpdatedAt     ChargebacksDBFieldNameType
	MetaUpdatedBy     ChargebacksDBFieldNameType
	MetaDeletedAt     ChargebacksDBFieldNameType
	MetaDeletedBy     ChargebacksDBFieldNameType
}

var ChargebacksDBFieldName = chargebacksDBFieldName{
	Id:                "id",
	ChargebackCode:    "chargeback_code",
	PaymentRefId:      "payment_ref_id",
	MerchantPartyId:   "merchant_party_id",
	ProviderAccountId: "provider_account_id",
	CurrencyCode:      "currency_code",
	DisputedAmount:    "disputed_amount",
	ChargebackStatus:  "chargeback_status",
	ReasonCode:        "reason_code",
	OpenedAt:          "opened_at",
	ClosedAt:          "closed_at",
	DueAt:             "due_at",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewChargebacksDBFieldNameFromStr(field string) (dbField ChargebacksDBFieldNameType, found bool) {
	switch field {

	case string(ChargebacksDBFieldName.Id):
		return ChargebacksDBFieldName.Id, true

	case string(ChargebacksDBFieldName.ChargebackCode):
		return ChargebacksDBFieldName.ChargebackCode, true

	case string(ChargebacksDBFieldName.PaymentRefId):
		return ChargebacksDBFieldName.PaymentRefId, true

	case string(ChargebacksDBFieldName.MerchantPartyId):
		return ChargebacksDBFieldName.MerchantPartyId, true

	case string(ChargebacksDBFieldName.ProviderAccountId):
		return ChargebacksDBFieldName.ProviderAccountId, true

	case string(ChargebacksDBFieldName.CurrencyCode):
		return ChargebacksDBFieldName.CurrencyCode, true

	case string(ChargebacksDBFieldName.DisputedAmount):
		return ChargebacksDBFieldName.DisputedAmount, true

	case string(ChargebacksDBFieldName.ChargebackStatus):
		return ChargebacksDBFieldName.ChargebackStatus, true

	case string(ChargebacksDBFieldName.ReasonCode):
		return ChargebacksDBFieldName.ReasonCode, true

	case string(ChargebacksDBFieldName.OpenedAt):
		return ChargebacksDBFieldName.OpenedAt, true

	case string(ChargebacksDBFieldName.ClosedAt):
		return ChargebacksDBFieldName.ClosedAt, true

	case string(ChargebacksDBFieldName.DueAt):
		return ChargebacksDBFieldName.DueAt, true

	case string(ChargebacksDBFieldName.Metadata):
		return ChargebacksDBFieldName.Metadata, true

	case string(ChargebacksDBFieldName.MetaCreatedAt):
		return ChargebacksDBFieldName.MetaCreatedAt, true

	case string(ChargebacksDBFieldName.MetaCreatedBy):
		return ChargebacksDBFieldName.MetaCreatedBy, true

	case string(ChargebacksDBFieldName.MetaUpdatedAt):
		return ChargebacksDBFieldName.MetaUpdatedAt, true

	case string(ChargebacksDBFieldName.MetaUpdatedBy):
		return ChargebacksDBFieldName.MetaUpdatedBy, true

	case string(ChargebacksDBFieldName.MetaDeletedAt):
		return ChargebacksDBFieldName.MetaDeletedAt, true

	case string(ChargebacksDBFieldName.MetaDeletedBy):
		return ChargebacksDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ChargebacksFilterJoins = map[string]JoinSpec{}

var ChargebacksFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"chargeback_code": {
		SourcePath:        "chargeback_code",
		DefaultOutputPath: "chargebackCode",
		Column:            "chargeback_code",
		SQLAlias:          "chargeback_code",
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
	"provider_account_id": {
		SourcePath:        "provider_account_id",
		DefaultOutputPath: "providerAccountId",
		Column:            "provider_account_id",
		SQLAlias:          "provider_account_id",
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
	"disputed_amount": {
		SourcePath:        "disputed_amount",
		DefaultOutputPath: "disputedAmount",
		Column:            "disputed_amount",
		SQLAlias:          "disputed_amount",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"chargeback_status": {
		SourcePath:        "chargeback_status",
		DefaultOutputPath: "chargebackStatus",
		Column:            "chargeback_status",
		SQLAlias:          "chargeback_status",
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
	"opened_at": {
		SourcePath:        "opened_at",
		DefaultOutputPath: "openedAt",
		Column:            "opened_at",
		SQLAlias:          "opened_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"closed_at": {
		SourcePath:        "closed_at",
		DefaultOutputPath: "closedAt",
		Column:            "closed_at",
		SQLAlias:          "closed_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"due_at": {
		SourcePath:        "due_at",
		DefaultOutputPath: "dueAt",
		Column:            "due_at",
		SQLAlias:          "due_at",
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

func NewChargebacksFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ChargebacksFilterFields[field]
	return
}

type ChargebacksFilterResult struct {
	Chargebacks
	FilterCount int `db:"count"`
}

func ValidateChargebacksFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewChargebacksFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewChargebacksFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewChargebacksFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateChargebacksFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateChargebacksFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewChargebacksFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateChargebacksFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ChargebackStatus string

const (
	ChargebackStatusOpened       ChargebackStatus = "opened"
	ChargebackStatusUnderReview  ChargebackStatus = "under_review"
	ChargebackStatusWon          ChargebackStatus = "won"
	ChargebackStatusLost         ChargebackStatus = "lost"
	ChargebackStatusPartiallyWon ChargebackStatus = "partially_won"
	ChargebackStatusReversed     ChargebackStatus = "reversed"
	ChargebackStatusClosed       ChargebackStatus = "closed"
)

type Chargebacks struct {
	Id                uuid.UUID        `db:"id"`
	ChargebackCode    string           `db:"chargeback_code"`
	PaymentRefId      uuid.UUID        `db:"payment_ref_id"`
	MerchantPartyId   nuuid.NUUID      `db:"merchant_party_id"`
	ProviderAccountId nuuid.NUUID      `db:"provider_account_id"`
	CurrencyCode      string           `db:"currency_code"`
	DisputedAmount    decimal.Decimal  `db:"disputed_amount"`
	ChargebackStatus  ChargebackStatus `db:"chargeback_status"`
	ReasonCode        null.String      `db:"reason_code"`
	OpenedAt          time.Time        `db:"opened_at"`
	ClosedAt          null.Time        `db:"closed_at"`
	DueAt             null.Time        `db:"due_at"`
	Metadata          json.RawMessage  `db:"metadata"`

	shared.MetaSignature
}
type ChargebacksPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d Chargebacks) ToChargebacksPrimaryID() ChargebacksPrimaryID {
	return ChargebacksPrimaryID{
		Id: d.Id,
	}
}

type ChargebacksList []*Chargebacks

type ChargebacksFilterResultList []*ChargebacksFilterResult
