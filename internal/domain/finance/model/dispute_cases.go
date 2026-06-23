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

type DisputeCasesDBFieldNameType string

type disputeCasesDBFieldName struct {
	Id                 DisputeCasesDBFieldNameType
	DisputeCode        DisputeCasesDBFieldNameType
	BookId             DisputeCasesDBFieldNameType
	SourceType         DisputeCasesDBFieldNameType
	SourceId           DisputeCasesDBFieldNameType
	MerchantPartyId    DisputeCasesDBFieldNameType
	CustomerPartyId    DisputeCasesDBFieldNameType
	ProviderCode       DisputeCasesDBFieldNameType
	ProviderDisputeRef DisputeCasesDBFieldNameType
	IdempotencyKey     DisputeCasesDBFieldNameType
	CurrencyCode       DisputeCasesDBFieldNameType
	DisputedAmount     DisputeCasesDBFieldNameType
	DisputeReasonCode  DisputeCasesDBFieldNameType
	DisputeStatus      DisputeCasesDBFieldNameType
	OpenedAt           DisputeCasesDBFieldNameType
	DueAt              DisputeCasesDBFieldNameType
	ClosedAt           DisputeCasesDBFieldNameType
	Metadata           DisputeCasesDBFieldNameType
	MetaCreatedAt      DisputeCasesDBFieldNameType
	MetaCreatedBy      DisputeCasesDBFieldNameType
	MetaUpdatedAt      DisputeCasesDBFieldNameType
	MetaUpdatedBy      DisputeCasesDBFieldNameType
	MetaDeletedAt      DisputeCasesDBFieldNameType
	MetaDeletedBy      DisputeCasesDBFieldNameType
}

var DisputeCasesDBFieldName = disputeCasesDBFieldName{
	Id:                 "id",
	DisputeCode:        "dispute_code",
	BookId:             "book_id",
	SourceType:         "source_type",
	SourceId:           "source_id",
	MerchantPartyId:    "merchant_party_id",
	CustomerPartyId:    "customer_party_id",
	ProviderCode:       "provider_code",
	ProviderDisputeRef: "provider_dispute_ref",
	IdempotencyKey:     "idempotency_key",
	CurrencyCode:       "currency_code",
	DisputedAmount:     "disputed_amount",
	DisputeReasonCode:  "dispute_reason_code",
	DisputeStatus:      "dispute_status",
	OpenedAt:           "opened_at",
	DueAt:              "due_at",
	ClosedAt:           "closed_at",
	Metadata:           "metadata",
	MetaCreatedAt:      "meta_created_at",
	MetaCreatedBy:      "meta_created_by",
	MetaUpdatedAt:      "meta_updated_at",
	MetaUpdatedBy:      "meta_updated_by",
	MetaDeletedAt:      "meta_deleted_at",
	MetaDeletedBy:      "meta_deleted_by",
}

func NewDisputeCasesDBFieldNameFromStr(field string) (dbField DisputeCasesDBFieldNameType, found bool) {
	switch field {

	case string(DisputeCasesDBFieldName.Id):
		return DisputeCasesDBFieldName.Id, true

	case string(DisputeCasesDBFieldName.DisputeCode):
		return DisputeCasesDBFieldName.DisputeCode, true

	case string(DisputeCasesDBFieldName.BookId):
		return DisputeCasesDBFieldName.BookId, true

	case string(DisputeCasesDBFieldName.SourceType):
		return DisputeCasesDBFieldName.SourceType, true

	case string(DisputeCasesDBFieldName.SourceId):
		return DisputeCasesDBFieldName.SourceId, true

	case string(DisputeCasesDBFieldName.MerchantPartyId):
		return DisputeCasesDBFieldName.MerchantPartyId, true

	case string(DisputeCasesDBFieldName.CustomerPartyId):
		return DisputeCasesDBFieldName.CustomerPartyId, true

	case string(DisputeCasesDBFieldName.ProviderCode):
		return DisputeCasesDBFieldName.ProviderCode, true

	case string(DisputeCasesDBFieldName.ProviderDisputeRef):
		return DisputeCasesDBFieldName.ProviderDisputeRef, true

	case string(DisputeCasesDBFieldName.IdempotencyKey):
		return DisputeCasesDBFieldName.IdempotencyKey, true

	case string(DisputeCasesDBFieldName.CurrencyCode):
		return DisputeCasesDBFieldName.CurrencyCode, true

	case string(DisputeCasesDBFieldName.DisputedAmount):
		return DisputeCasesDBFieldName.DisputedAmount, true

	case string(DisputeCasesDBFieldName.DisputeReasonCode):
		return DisputeCasesDBFieldName.DisputeReasonCode, true

	case string(DisputeCasesDBFieldName.DisputeStatus):
		return DisputeCasesDBFieldName.DisputeStatus, true

	case string(DisputeCasesDBFieldName.OpenedAt):
		return DisputeCasesDBFieldName.OpenedAt, true

	case string(DisputeCasesDBFieldName.DueAt):
		return DisputeCasesDBFieldName.DueAt, true

	case string(DisputeCasesDBFieldName.ClosedAt):
		return DisputeCasesDBFieldName.ClosedAt, true

	case string(DisputeCasesDBFieldName.Metadata):
		return DisputeCasesDBFieldName.Metadata, true

	case string(DisputeCasesDBFieldName.MetaCreatedAt):
		return DisputeCasesDBFieldName.MetaCreatedAt, true

	case string(DisputeCasesDBFieldName.MetaCreatedBy):
		return DisputeCasesDBFieldName.MetaCreatedBy, true

	case string(DisputeCasesDBFieldName.MetaUpdatedAt):
		return DisputeCasesDBFieldName.MetaUpdatedAt, true

	case string(DisputeCasesDBFieldName.MetaUpdatedBy):
		return DisputeCasesDBFieldName.MetaUpdatedBy, true

	case string(DisputeCasesDBFieldName.MetaDeletedAt):
		return DisputeCasesDBFieldName.MetaDeletedAt, true

	case string(DisputeCasesDBFieldName.MetaDeletedBy):
		return DisputeCasesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var DisputeCasesFilterJoins = map[string]JoinSpec{}

var DisputeCasesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"dispute_code": {
		SourcePath:        "dispute_code",
		DefaultOutputPath: "disputeCode",
		Column:            "dispute_code",
		SQLAlias:          "dispute_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"book_id": {
		SourcePath:        "book_id",
		DefaultOutputPath: "bookId",
		Column:            "book_id",
		SQLAlias:          "book_id",
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
	"customer_party_id": {
		SourcePath:        "customer_party_id",
		DefaultOutputPath: "customerPartyId",
		Column:            "customer_party_id",
		SQLAlias:          "customer_party_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_code": {
		SourcePath:        "provider_code",
		DefaultOutputPath: "providerCode",
		Column:            "provider_code",
		SQLAlias:          "provider_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"provider_dispute_ref": {
		SourcePath:        "provider_dispute_ref",
		DefaultOutputPath: "providerDisputeRef",
		Column:            "provider_dispute_ref",
		SQLAlias:          "provider_dispute_ref",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"idempotency_key": {
		SourcePath:        "idempotency_key",
		DefaultOutputPath: "idempotencyKey",
		Column:            "idempotency_key",
		SQLAlias:          "idempotency_key",
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
	"dispute_reason_code": {
		SourcePath:        "dispute_reason_code",
		DefaultOutputPath: "disputeReasonCode",
		Column:            "dispute_reason_code",
		SQLAlias:          "dispute_reason_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"dispute_status": {
		SourcePath:        "dispute_status",
		DefaultOutputPath: "disputeStatus",
		Column:            "dispute_status",
		SQLAlias:          "dispute_status",
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
	"due_at": {
		SourcePath:        "due_at",
		DefaultOutputPath: "dueAt",
		Column:            "due_at",
		SQLAlias:          "due_at",
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

func NewDisputeCasesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = DisputeCasesFilterFields[field]
	return
}

type DisputeCasesFilterResult struct {
	DisputeCases
	FilterCount int `db:"count"`
}

func ValidateDisputeCasesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewDisputeCasesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewDisputeCasesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewDisputeCasesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateDisputeCasesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateDisputeCasesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewDisputeCasesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateDisputeCasesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type DisputeStatus string

const (
	DisputeStatusOpened            DisputeStatus = "opened"
	DisputeStatusEvidenceRequired  DisputeStatus = "evidence_required"
	DisputeStatusEvidenceSubmitted DisputeStatus = "evidence_submitted"
	DisputeStatusUnderReview       DisputeStatus = "under_review"
	DisputeStatusWon               DisputeStatus = "won"
	DisputeStatusLost              DisputeStatus = "lost"
	DisputeStatusAccepted          DisputeStatus = "accepted"
	DisputeStatusClosed            DisputeStatus = "closed"
	DisputeStatusCanceled          DisputeStatus = "canceled"
)

type DisputeCases struct {
	Id                 uuid.UUID       `db:"id"`
	DisputeCode        string          `db:"dispute_code"`
	BookId             uuid.UUID       `db:"book_id"`
	SourceType         string          `db:"source_type"`
	SourceId           uuid.UUID       `db:"source_id"`
	MerchantPartyId    nuuid.NUUID     `db:"merchant_party_id"`
	CustomerPartyId    nuuid.NUUID     `db:"customer_party_id"`
	ProviderCode       null.String     `db:"provider_code"`
	ProviderDisputeRef null.String     `db:"provider_dispute_ref"`
	IdempotencyKey     null.String     `db:"idempotency_key"`
	CurrencyCode       string          `db:"currency_code"`
	DisputedAmount     decimal.Decimal `db:"disputed_amount"`
	DisputeReasonCode  string          `db:"dispute_reason_code"`
	DisputeStatus      DisputeStatus   `db:"dispute_status"`
	OpenedAt           time.Time       `db:"opened_at"`
	DueAt              null.Time       `db:"due_at"`
	ClosedAt           null.Time       `db:"closed_at"`
	Metadata           json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type DisputeCasesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d DisputeCases) ToDisputeCasesPrimaryID() DisputeCasesPrimaryID {
	return DisputeCasesPrimaryID{
		Id: d.Id,
	}
}

type DisputeCasesList []*DisputeCases

type DisputeCasesFilterResultList []*DisputeCasesFilterResult
