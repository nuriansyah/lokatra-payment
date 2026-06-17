package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderCircuitBreakersDBFieldNameType string

type providerCircuitBreakersDBFieldName struct {
	Id                ProviderCircuitBreakersDBFieldNameType
	ProviderAccountId ProviderCircuitBreakersDBFieldNameType
	MethodCode        ProviderCircuitBreakersDBFieldNameType
	ChannelCode       ProviderCircuitBreakersDBFieldNameType
	Status            ProviderCircuitBreakersDBFieldNameType
	FailureCount      ProviderCircuitBreakersDBFieldNameType
	SuccessCount      ProviderCircuitBreakersDBFieldNameType
	LastFailureAt     ProviderCircuitBreakersDBFieldNameType
	LastSuccessAt     ProviderCircuitBreakersDBFieldNameType
	OpenedAt          ProviderCircuitBreakersDBFieldNameType
	OpenUntil         ProviderCircuitBreakersDBFieldNameType
	HalfOpenAt        ProviderCircuitBreakersDBFieldNameType
	Reason            ProviderCircuitBreakersDBFieldNameType
	Metadata          ProviderCircuitBreakersDBFieldNameType
	MetaCreatedAt     ProviderCircuitBreakersDBFieldNameType
	MetaCreatedBy     ProviderCircuitBreakersDBFieldNameType
	MetaUpdatedAt     ProviderCircuitBreakersDBFieldNameType
	MetaUpdatedBy     ProviderCircuitBreakersDBFieldNameType
	MetaDeletedAt     ProviderCircuitBreakersDBFieldNameType
	MetaDeletedBy     ProviderCircuitBreakersDBFieldNameType
}

var ProviderCircuitBreakersDBFieldName = providerCircuitBreakersDBFieldName{
	Id:                "id",
	ProviderAccountId: "provider_account_id",
	MethodCode:        "method_code",
	ChannelCode:       "channel_code",
	Status:            "status",
	FailureCount:      "failure_count",
	SuccessCount:      "success_count",
	LastFailureAt:     "last_failure_at",
	LastSuccessAt:     "last_success_at",
	OpenedAt:          "opened_at",
	OpenUntil:         "open_until",
	HalfOpenAt:        "half_open_at",
	Reason:            "reason",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewProviderCircuitBreakersDBFieldNameFromStr(field string) (dbField ProviderCircuitBreakersDBFieldNameType, found bool) {
	switch field {

	case string(ProviderCircuitBreakersDBFieldName.Id):
		return ProviderCircuitBreakersDBFieldName.Id, true

	case string(ProviderCircuitBreakersDBFieldName.ProviderAccountId):
		return ProviderCircuitBreakersDBFieldName.ProviderAccountId, true

	case string(ProviderCircuitBreakersDBFieldName.MethodCode):
		return ProviderCircuitBreakersDBFieldName.MethodCode, true

	case string(ProviderCircuitBreakersDBFieldName.ChannelCode):
		return ProviderCircuitBreakersDBFieldName.ChannelCode, true

	case string(ProviderCircuitBreakersDBFieldName.Status):
		return ProviderCircuitBreakersDBFieldName.Status, true

	case string(ProviderCircuitBreakersDBFieldName.FailureCount):
		return ProviderCircuitBreakersDBFieldName.FailureCount, true

	case string(ProviderCircuitBreakersDBFieldName.SuccessCount):
		return ProviderCircuitBreakersDBFieldName.SuccessCount, true

	case string(ProviderCircuitBreakersDBFieldName.LastFailureAt):
		return ProviderCircuitBreakersDBFieldName.LastFailureAt, true

	case string(ProviderCircuitBreakersDBFieldName.LastSuccessAt):
		return ProviderCircuitBreakersDBFieldName.LastSuccessAt, true

	case string(ProviderCircuitBreakersDBFieldName.OpenedAt):
		return ProviderCircuitBreakersDBFieldName.OpenedAt, true

	case string(ProviderCircuitBreakersDBFieldName.OpenUntil):
		return ProviderCircuitBreakersDBFieldName.OpenUntil, true

	case string(ProviderCircuitBreakersDBFieldName.HalfOpenAt):
		return ProviderCircuitBreakersDBFieldName.HalfOpenAt, true

	case string(ProviderCircuitBreakersDBFieldName.Reason):
		return ProviderCircuitBreakersDBFieldName.Reason, true

	case string(ProviderCircuitBreakersDBFieldName.Metadata):
		return ProviderCircuitBreakersDBFieldName.Metadata, true

	case string(ProviderCircuitBreakersDBFieldName.MetaCreatedAt):
		return ProviderCircuitBreakersDBFieldName.MetaCreatedAt, true

	case string(ProviderCircuitBreakersDBFieldName.MetaCreatedBy):
		return ProviderCircuitBreakersDBFieldName.MetaCreatedBy, true

	case string(ProviderCircuitBreakersDBFieldName.MetaUpdatedAt):
		return ProviderCircuitBreakersDBFieldName.MetaUpdatedAt, true

	case string(ProviderCircuitBreakersDBFieldName.MetaUpdatedBy):
		return ProviderCircuitBreakersDBFieldName.MetaUpdatedBy, true

	case string(ProviderCircuitBreakersDBFieldName.MetaDeletedAt):
		return ProviderCircuitBreakersDBFieldName.MetaDeletedAt, true

	case string(ProviderCircuitBreakersDBFieldName.MetaDeletedBy):
		return ProviderCircuitBreakersDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderCircuitBreakersFilterJoins = map[string]JoinSpec{}

var ProviderCircuitBreakersFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
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
	"method_code": {
		SourcePath:        "method_code",
		DefaultOutputPath: "methodCode",
		Column:            "method_code",
		SQLAlias:          "method_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"channel_code": {
		SourcePath:        "channel_code",
		DefaultOutputPath: "channelCode",
		Column:            "channel_code",
		SQLAlias:          "channel_code",
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
	"failure_count": {
		SourcePath:        "failure_count",
		DefaultOutputPath: "failureCount",
		Column:            "failure_count",
		SQLAlias:          "failure_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"success_count": {
		SourcePath:        "success_count",
		DefaultOutputPath: "successCount",
		Column:            "success_count",
		SQLAlias:          "success_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"last_failure_at": {
		SourcePath:        "last_failure_at",
		DefaultOutputPath: "lastFailureAt",
		Column:            "last_failure_at",
		SQLAlias:          "last_failure_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"last_success_at": {
		SourcePath:        "last_success_at",
		DefaultOutputPath: "lastSuccessAt",
		Column:            "last_success_at",
		SQLAlias:          "last_success_at",
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
	"open_until": {
		SourcePath:        "open_until",
		DefaultOutputPath: "openUntil",
		Column:            "open_until",
		SQLAlias:          "open_until",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"half_open_at": {
		SourcePath:        "half_open_at",
		DefaultOutputPath: "halfOpenAt",
		Column:            "half_open_at",
		SQLAlias:          "half_open_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"reason": {
		SourcePath:        "reason",
		DefaultOutputPath: "reason",
		Column:            "reason",
		SQLAlias:          "reason",
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

func NewProviderCircuitBreakersFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderCircuitBreakersFilterFields[field]
	return
}

type ProviderCircuitBreakersFilterResult struct {
	ProviderCircuitBreakers
	FilterCount int `db:"count"`
}

func ValidateProviderCircuitBreakersFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderCircuitBreakersFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderCircuitBreakersFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderCircuitBreakersFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderCircuitBreakersFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderCircuitBreakersFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderCircuitBreakersFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderCircuitBreakersFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type CircuitStatus string

const (
	CircuitStatusClosed   CircuitStatus = "closed"
	CircuitStatusOpen     CircuitStatus = "open"
	CircuitStatusHalfOpen CircuitStatus = "half_open"
)

type ProviderCircuitBreakers struct {
	Id                uuid.UUID       `db:"id"`
	ProviderAccountId uuid.UUID       `db:"provider_account_id"`
	MethodCode        null.String     `db:"method_code"`
	ChannelCode       null.String     `db:"channel_code"`
	Status            CircuitStatus   `db:"status"`
	FailureCount      int             `db:"failure_count"`
	SuccessCount      int             `db:"success_count"`
	LastFailureAt     null.Time       `db:"last_failure_at"`
	LastSuccessAt     null.Time       `db:"last_success_at"`
	OpenedAt          null.Time       `db:"opened_at"`
	OpenUntil         null.Time       `db:"open_until"`
	HalfOpenAt        null.Time       `db:"half_open_at"`
	Reason            null.String     `db:"reason"`
	Metadata          json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type ProviderCircuitBreakersPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderCircuitBreakers) ToProviderCircuitBreakersPrimaryID() ProviderCircuitBreakersPrimaryID {
	return ProviderCircuitBreakersPrimaryID{
		Id: d.Id,
	}
}

type ProviderCircuitBreakersList []*ProviderCircuitBreakers

type ProviderCircuitBreakersFilterResultList []*ProviderCircuitBreakersFilterResult
