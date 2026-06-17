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

type CashCollectionItemsDBFieldNameType string

type cashCollectionItemsDBFieldName struct {
	Id                      CashCollectionItemsDBFieldNameType
	CashCollectionSessionId CashCollectionItemsDBFieldNameType
	PaymentIntentId         CashCollectionItemsDBFieldNameType
	PaymentAttemptId        CashCollectionItemsDBFieldNameType
	CollectionType          CashCollectionItemsDBFieldNameType
	Amount                  CashCollectionItemsDBFieldNameType
	Currency                CashCollectionItemsDBFieldNameType
	Status                  CashCollectionItemsDBFieldNameType
	CollectedAt             CashCollectionItemsDBFieldNameType
	VoidedAt                CashCollectionItemsDBFieldNameType
	VoidReason              CashCollectionItemsDBFieldNameType
	Notes                   CashCollectionItemsDBFieldNameType
	Metadata                CashCollectionItemsDBFieldNameType
	MetaCreatedAt           CashCollectionItemsDBFieldNameType
	MetaCreatedBy           CashCollectionItemsDBFieldNameType
	MetaUpdatedAt           CashCollectionItemsDBFieldNameType
	MetaUpdatedBy           CashCollectionItemsDBFieldNameType
	MetaDeletedAt           CashCollectionItemsDBFieldNameType
	MetaDeletedBy           CashCollectionItemsDBFieldNameType
}

var CashCollectionItemsDBFieldName = cashCollectionItemsDBFieldName{
	Id:                      "id",
	CashCollectionSessionId: "cash_collection_session_id",
	PaymentIntentId:         "payment_intent_id",
	PaymentAttemptId:        "payment_attempt_id",
	CollectionType:          "collection_type",
	Amount:                  "amount",
	Currency:                "currency",
	Status:                  "status",
	CollectedAt:             "collected_at",
	VoidedAt:                "voided_at",
	VoidReason:              "void_reason",
	Notes:                   "notes",
	Metadata:                "metadata",
	MetaCreatedAt:           "meta_created_at",
	MetaCreatedBy:           "meta_created_by",
	MetaUpdatedAt:           "meta_updated_at",
	MetaUpdatedBy:           "meta_updated_by",
	MetaDeletedAt:           "meta_deleted_at",
	MetaDeletedBy:           "meta_deleted_by",
}

func NewCashCollectionItemsDBFieldNameFromStr(field string) (dbField CashCollectionItemsDBFieldNameType, found bool) {
	switch field {

	case string(CashCollectionItemsDBFieldName.Id):
		return CashCollectionItemsDBFieldName.Id, true

	case string(CashCollectionItemsDBFieldName.CashCollectionSessionId):
		return CashCollectionItemsDBFieldName.CashCollectionSessionId, true

	case string(CashCollectionItemsDBFieldName.PaymentIntentId):
		return CashCollectionItemsDBFieldName.PaymentIntentId, true

	case string(CashCollectionItemsDBFieldName.PaymentAttemptId):
		return CashCollectionItemsDBFieldName.PaymentAttemptId, true

	case string(CashCollectionItemsDBFieldName.CollectionType):
		return CashCollectionItemsDBFieldName.CollectionType, true

	case string(CashCollectionItemsDBFieldName.Amount):
		return CashCollectionItemsDBFieldName.Amount, true

	case string(CashCollectionItemsDBFieldName.Currency):
		return CashCollectionItemsDBFieldName.Currency, true

	case string(CashCollectionItemsDBFieldName.Status):
		return CashCollectionItemsDBFieldName.Status, true

	case string(CashCollectionItemsDBFieldName.CollectedAt):
		return CashCollectionItemsDBFieldName.CollectedAt, true

	case string(CashCollectionItemsDBFieldName.VoidedAt):
		return CashCollectionItemsDBFieldName.VoidedAt, true

	case string(CashCollectionItemsDBFieldName.VoidReason):
		return CashCollectionItemsDBFieldName.VoidReason, true

	case string(CashCollectionItemsDBFieldName.Notes):
		return CashCollectionItemsDBFieldName.Notes, true

	case string(CashCollectionItemsDBFieldName.Metadata):
		return CashCollectionItemsDBFieldName.Metadata, true

	case string(CashCollectionItemsDBFieldName.MetaCreatedAt):
		return CashCollectionItemsDBFieldName.MetaCreatedAt, true

	case string(CashCollectionItemsDBFieldName.MetaCreatedBy):
		return CashCollectionItemsDBFieldName.MetaCreatedBy, true

	case string(CashCollectionItemsDBFieldName.MetaUpdatedAt):
		return CashCollectionItemsDBFieldName.MetaUpdatedAt, true

	case string(CashCollectionItemsDBFieldName.MetaUpdatedBy):
		return CashCollectionItemsDBFieldName.MetaUpdatedBy, true

	case string(CashCollectionItemsDBFieldName.MetaDeletedAt):
		return CashCollectionItemsDBFieldName.MetaDeletedAt, true

	case string(CashCollectionItemsDBFieldName.MetaDeletedBy):
		return CashCollectionItemsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var CashCollectionItemsFilterJoins = map[string]JoinSpec{}

var CashCollectionItemsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"cash_collection_session_id": {
		SourcePath:        "cash_collection_session_id",
		DefaultOutputPath: "cashCollectionSessionId",
		Column:            "cash_collection_session_id",
		SQLAlias:          "cash_collection_session_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_intent_id": {
		SourcePath:        "payment_intent_id",
		DefaultOutputPath: "paymentIntentId",
		Column:            "payment_intent_id",
		SQLAlias:          "payment_intent_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_attempt_id": {
		SourcePath:        "payment_attempt_id",
		DefaultOutputPath: "paymentAttemptId",
		Column:            "payment_attempt_id",
		SQLAlias:          "payment_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"collection_type": {
		SourcePath:        "collection_type",
		DefaultOutputPath: "collectionType",
		Column:            "collection_type",
		SQLAlias:          "collection_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"amount": {
		SourcePath:        "amount",
		DefaultOutputPath: "amount",
		Column:            "amount",
		SQLAlias:          "amount",
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
	"status": {
		SourcePath:        "status",
		DefaultOutputPath: "status",
		Column:            "status",
		SQLAlias:          "status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"collected_at": {
		SourcePath:        "collected_at",
		DefaultOutputPath: "collectedAt",
		Column:            "collected_at",
		SQLAlias:          "collected_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"voided_at": {
		SourcePath:        "voided_at",
		DefaultOutputPath: "voidedAt",
		Column:            "voided_at",
		SQLAlias:          "voided_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"void_reason": {
		SourcePath:        "void_reason",
		DefaultOutputPath: "voidReason",
		Column:            "void_reason",
		SQLAlias:          "void_reason",
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

func NewCashCollectionItemsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = CashCollectionItemsFilterFields[field]
	return
}

type CashCollectionItemsFilterResult struct {
	CashCollectionItems
	FilterCount int `db:"count"`
}

func ValidateCashCollectionItemsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewCashCollectionItemsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewCashCollectionItemsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewCashCollectionItemsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateCashCollectionItemsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateCashCollectionItemsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewCashCollectionItemsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateCashCollectionItemsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type CashItemStatus string

const (
	CashItemStatusCollected CashItemStatus = "collected"
	CashItemStatusVoided    CashItemStatus = "voided"
)

type CashCollectionItems struct {
	Id                      uuid.UUID       `db:"id"`
	CashCollectionSessionId uuid.UUID       `db:"cash_collection_session_id"`
	PaymentIntentId         uuid.UUID       `db:"payment_intent_id"`
	PaymentAttemptId        nuuid.NUUID     `db:"payment_attempt_id"`
	CollectionType          string          `db:"collection_type"`
	Amount                  decimal.Decimal `db:"amount"`
	Currency                string          `db:"currency"`
	Status                  CashItemStatus  `db:"status"`
	CollectedAt             time.Time       `db:"collected_at"`
	VoidedAt                null.Time       `db:"voided_at"`
	VoidReason              null.String     `db:"void_reason"`
	Notes                   null.String     `db:"notes"`
	Metadata                json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type CashCollectionItemsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d CashCollectionItems) ToCashCollectionItemsPrimaryID() CashCollectionItemsPrimaryID {
	return CashCollectionItemsPrimaryID{
		Id: d.Id,
	}
}

type CashCollectionItemsList []*CashCollectionItems

type CashCollectionItemsFilterResultList []*CashCollectionItemsFilterResult
