package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

type CashCollectionSessionsDBFieldNameType string

type cashCollectionSessionsDBFieldName struct {
	Id                 CashCollectionSessionsDBFieldNameType
	SessionCode        CashCollectionSessionsDBFieldNameType
	MerchantId         CashCollectionSessionsDBFieldNameType
	CollectorId        CashCollectionSessionsDBFieldNameType
	LocationId         CashCollectionSessionsDBFieldNameType
	OpenedAt           CashCollectionSessionsDBFieldNameType
	ClosedAt           CashCollectionSessionsDBFieldNameType
	Status             CashCollectionSessionsDBFieldNameType
	OpeningFloatAmount CashCollectionSessionsDBFieldNameType
	ExpectedAmount     CashCollectionSessionsDBFieldNameType
	CountedAmount      CashCollectionSessionsDBFieldNameType
	VarianceAmount     CashCollectionSessionsDBFieldNameType
	Currency           CashCollectionSessionsDBFieldNameType
	Notes              CashCollectionSessionsDBFieldNameType
	Metadata           CashCollectionSessionsDBFieldNameType
	MetaCreatedAt      CashCollectionSessionsDBFieldNameType
	MetaCreatedBy      CashCollectionSessionsDBFieldNameType
	MetaUpdatedAt      CashCollectionSessionsDBFieldNameType
	MetaUpdatedBy      CashCollectionSessionsDBFieldNameType
	MetaDeletedAt      CashCollectionSessionsDBFieldNameType
	MetaDeletedBy      CashCollectionSessionsDBFieldNameType
}

var CashCollectionSessionsDBFieldName = cashCollectionSessionsDBFieldName{
	Id:                 "id",
	SessionCode:        "session_code",
	MerchantId:         "merchant_id",
	CollectorId:        "collector_id",
	LocationId:         "location_id",
	OpenedAt:           "opened_at",
	ClosedAt:           "closed_at",
	Status:             "status",
	OpeningFloatAmount: "opening_float_amount",
	ExpectedAmount:     "expected_amount",
	CountedAmount:      "counted_amount",
	VarianceAmount:     "variance_amount",
	Currency:           "currency",
	Notes:              "notes",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewCashCollectionSessionsDBFieldNameFromStr(field string) (dbField CashCollectionSessionsDBFieldNameType, found bool) {
	switch field {

	case string(CashCollectionSessionsDBFieldName.Id):
		return CashCollectionSessionsDBFieldName.Id, true

	case string(CashCollectionSessionsDBFieldName.SessionCode):
		return CashCollectionSessionsDBFieldName.SessionCode, true

	case string(CashCollectionSessionsDBFieldName.MerchantId):
		return CashCollectionSessionsDBFieldName.MerchantId, true

	case string(CashCollectionSessionsDBFieldName.CollectorId):
		return CashCollectionSessionsDBFieldName.CollectorId, true

	case string(CashCollectionSessionsDBFieldName.LocationId):
		return CashCollectionSessionsDBFieldName.LocationId, true

	case string(CashCollectionSessionsDBFieldName.OpenedAt):
		return CashCollectionSessionsDBFieldName.OpenedAt, true

	case string(CashCollectionSessionsDBFieldName.ClosedAt):
		return CashCollectionSessionsDBFieldName.ClosedAt, true

	case string(CashCollectionSessionsDBFieldName.Status):
		return CashCollectionSessionsDBFieldName.Status, true

	case string(CashCollectionSessionsDBFieldName.OpeningFloatAmount):
		return CashCollectionSessionsDBFieldName.OpeningFloatAmount, true

	case string(CashCollectionSessionsDBFieldName.ExpectedAmount):
		return CashCollectionSessionsDBFieldName.ExpectedAmount, true

	case string(CashCollectionSessionsDBFieldName.CountedAmount):
		return CashCollectionSessionsDBFieldName.CountedAmount, true

	case string(CashCollectionSessionsDBFieldName.VarianceAmount):
		return CashCollectionSessionsDBFieldName.VarianceAmount, true

	case string(CashCollectionSessionsDBFieldName.Currency):
		return CashCollectionSessionsDBFieldName.Currency, true

	case string(CashCollectionSessionsDBFieldName.Notes):
		return CashCollectionSessionsDBFieldName.Notes, true

	case string(CashCollectionSessionsDBFieldName.Metadata):
		return CashCollectionSessionsDBFieldName.Metadata, true

	case string(CashCollectionSessionsDBFieldName.MetaCreatedAt):
		return CashCollectionSessionsDBFieldName.MetaCreatedAt, true

	case string(CashCollectionSessionsDBFieldName.MetaCreatedBy):
		return CashCollectionSessionsDBFieldName.MetaCreatedBy, true

	case string(CashCollectionSessionsDBFieldName.MetaUpdatedAt):
		return CashCollectionSessionsDBFieldName.MetaUpdatedAt, true

	case string(CashCollectionSessionsDBFieldName.MetaUpdatedBy):
		return CashCollectionSessionsDBFieldName.MetaUpdatedBy, true

	case string(CashCollectionSessionsDBFieldName.MetaDeletedAt):
		return CashCollectionSessionsDBFieldName.MetaDeletedAt, true

	case string(CashCollectionSessionsDBFieldName.MetaDeletedBy):
		return CashCollectionSessionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var CashCollectionSessionsFilterJoins = map[string]JoinSpec{}

var CashCollectionSessionsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"session_code": {
		SourcePath:        "session_code",
		DefaultOutputPath: "sessionCode",
		Column:            "session_code",
		SQLAlias:          "session_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"merchant_id": {
		SourcePath:        "merchant_id",
		DefaultOutputPath: "merchantId",
		Column:            "merchant_id",
		SQLAlias:          "merchant_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"collector_id": {
		SourcePath:        "collector_id",
		DefaultOutputPath: "collectorId",
		Column:            "collector_id",
		SQLAlias:          "collector_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"location_id": {
		SourcePath:        "location_id",
		DefaultOutputPath: "locationId",
		Column:            "location_id",
		SQLAlias:          "location_id",
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
	"status": {
		SourcePath:        "status",
		DefaultOutputPath: "status",
		Column:            "status",
		SQLAlias:          "status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"opening_float_amount": {
		SourcePath:        "opening_float_amount",
		DefaultOutputPath: "openingFloatAmount",
		Column:            "opening_float_amount",
		SQLAlias:          "opening_float_amount",
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
	"counted_amount": {
		SourcePath:        "counted_amount",
		DefaultOutputPath: "countedAmount",
		Column:            "counted_amount",
		SQLAlias:          "counted_amount",
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
	"currency": {
		SourcePath:        "currency",
		DefaultOutputPath: "currency",
		Column:            "currency",
		SQLAlias:          "currency",
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

func NewCashCollectionSessionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = CashCollectionSessionsFilterFields[field]
	return
}

type CashCollectionSessionsFilterResult struct {
	CashCollectionSessions
	FilterCount int `db:"count"`
}

func ValidateCashCollectionSessionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewCashCollectionSessionsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewCashCollectionSessionsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewCashCollectionSessionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateCashCollectionSessionsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateCashCollectionSessionsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewCashCollectionSessionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateCashCollectionSessionsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type CashSessionStatus string

const (
	CashSessionStatusOpen     CashSessionStatus = "open"
	CashSessionStatusClosed   CashSessionStatus = "closed"
	CashSessionStatusCanceled CashSessionStatus = "canceled"
)

type CashCollectionSessions struct {
	Id                 uuid.UUID         `db:"id"`
	SessionCode        string            `db:"session_code"`
	MerchantId         uuid.UUID         `db:"merchant_id"`
	CollectorId        uuid.UUID         `db:"collector_id"`
	LocationId         nuuid.NUUID       `db:"location_id"`
	OpenedAt           time.Time         `db:"opened_at"`
	ClosedAt           null.Time         `db:"closed_at"`
	Status             CashSessionStatus `db:"status"`
	OpeningFloatAmount decimal.Decimal   `db:"opening_float_amount"`
	ExpectedAmount     decimal.Decimal   `db:"expected_amount"`
	CountedAmount      decimal.Decimal   `db:"counted_amount"`
	VarianceAmount     decimal.Decimal   `db:"variance_amount"`
	Currency           string            `db:"currency"`
	Notes              null.String       `db:"notes"`
	Metadata           json.RawMessage   `db:"metadata"`

	shared.MetaSignature
}
type CashCollectionSessionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d CashCollectionSessions) ToCashCollectionSessionsPrimaryID() CashCollectionSessionsPrimaryID {
	return CashCollectionSessionsPrimaryID{
		Id: d.Id,
	}
}

type CashCollectionSessionsList []*CashCollectionSessions

type CashCollectionSessionsFilterResultList []*CashCollectionSessionsFilterResult
