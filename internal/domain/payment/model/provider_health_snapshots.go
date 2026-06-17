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

type ProviderHealthSnapshotsDBFieldNameType string

type providerHealthSnapshotsDBFieldName struct {
	Id                ProviderHealthSnapshotsDBFieldNameType
	ProviderAccountId ProviderHealthSnapshotsDBFieldNameType
	MethodCode        ProviderHealthSnapshotsDBFieldNameType
	ChannelCode       ProviderHealthSnapshotsDBFieldNameType
	HealthScore       ProviderHealthSnapshotsDBFieldNameType
	SuccessRate       ProviderHealthSnapshotsDBFieldNameType
	TimeoutRate       ProviderHealthSnapshotsDBFieldNameType
	ErrorRate         ProviderHealthSnapshotsDBFieldNameType
	P95LatencyMs      ProviderHealthSnapshotsDBFieldNameType
	SampleSize        ProviderHealthSnapshotsDBFieldNameType
	WindowStartedAt   ProviderHealthSnapshotsDBFieldNameType
	WindowEndedAt     ProviderHealthSnapshotsDBFieldNameType
	Metadata          ProviderHealthSnapshotsDBFieldNameType
	MetaCreatedAt     ProviderHealthSnapshotsDBFieldNameType
	MetaCreatedBy     ProviderHealthSnapshotsDBFieldNameType
	MetaUpdatedAt     ProviderHealthSnapshotsDBFieldNameType
	MetaUpdatedBy     ProviderHealthSnapshotsDBFieldNameType
	MetaDeletedAt     ProviderHealthSnapshotsDBFieldNameType
	MetaDeletedBy     ProviderHealthSnapshotsDBFieldNameType
}

var ProviderHealthSnapshotsDBFieldName = providerHealthSnapshotsDBFieldName{
	Id:                "id",
	ProviderAccountId: "provider_account_id",
	MethodCode:        "method_code",
	ChannelCode:       "channel_code",
	HealthScore:       "health_score",
	SuccessRate:       "success_rate",
	TimeoutRate:       "timeout_rate",
	ErrorRate:         "error_rate",
	P95LatencyMs:      "p95_latency_ms",
	SampleSize:        "sample_size",
	WindowStartedAt:   "window_started_at",
	WindowEndedAt:     "window_ended_at",
	Metadata:          "metadata",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewProviderHealthSnapshotsDBFieldNameFromStr(field string) (dbField ProviderHealthSnapshotsDBFieldNameType, found bool) {
	switch field {

	case string(ProviderHealthSnapshotsDBFieldName.Id):
		return ProviderHealthSnapshotsDBFieldName.Id, true

	case string(ProviderHealthSnapshotsDBFieldName.ProviderAccountId):
		return ProviderHealthSnapshotsDBFieldName.ProviderAccountId, true

	case string(ProviderHealthSnapshotsDBFieldName.MethodCode):
		return ProviderHealthSnapshotsDBFieldName.MethodCode, true

	case string(ProviderHealthSnapshotsDBFieldName.ChannelCode):
		return ProviderHealthSnapshotsDBFieldName.ChannelCode, true

	case string(ProviderHealthSnapshotsDBFieldName.HealthScore):
		return ProviderHealthSnapshotsDBFieldName.HealthScore, true

	case string(ProviderHealthSnapshotsDBFieldName.SuccessRate):
		return ProviderHealthSnapshotsDBFieldName.SuccessRate, true

	case string(ProviderHealthSnapshotsDBFieldName.TimeoutRate):
		return ProviderHealthSnapshotsDBFieldName.TimeoutRate, true

	case string(ProviderHealthSnapshotsDBFieldName.ErrorRate):
		return ProviderHealthSnapshotsDBFieldName.ErrorRate, true

	case string(ProviderHealthSnapshotsDBFieldName.P95LatencyMs):
		return ProviderHealthSnapshotsDBFieldName.P95LatencyMs, true

	case string(ProviderHealthSnapshotsDBFieldName.SampleSize):
		return ProviderHealthSnapshotsDBFieldName.SampleSize, true

	case string(ProviderHealthSnapshotsDBFieldName.WindowStartedAt):
		return ProviderHealthSnapshotsDBFieldName.WindowStartedAt, true

	case string(ProviderHealthSnapshotsDBFieldName.WindowEndedAt):
		return ProviderHealthSnapshotsDBFieldName.WindowEndedAt, true

	case string(ProviderHealthSnapshotsDBFieldName.Metadata):
		return ProviderHealthSnapshotsDBFieldName.Metadata, true

	case string(ProviderHealthSnapshotsDBFieldName.MetaCreatedAt):
		return ProviderHealthSnapshotsDBFieldName.MetaCreatedAt, true

	case string(ProviderHealthSnapshotsDBFieldName.MetaCreatedBy):
		return ProviderHealthSnapshotsDBFieldName.MetaCreatedBy, true

	case string(ProviderHealthSnapshotsDBFieldName.MetaUpdatedAt):
		return ProviderHealthSnapshotsDBFieldName.MetaUpdatedAt, true

	case string(ProviderHealthSnapshotsDBFieldName.MetaUpdatedBy):
		return ProviderHealthSnapshotsDBFieldName.MetaUpdatedBy, true

	case string(ProviderHealthSnapshotsDBFieldName.MetaDeletedAt):
		return ProviderHealthSnapshotsDBFieldName.MetaDeletedAt, true

	case string(ProviderHealthSnapshotsDBFieldName.MetaDeletedBy):
		return ProviderHealthSnapshotsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var ProviderHealthSnapshotsFilterJoins = map[string]JoinSpec{}

var ProviderHealthSnapshotsFilterFields = map[string]FilterFieldSpec{
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
	"health_score": {
		SourcePath:        "health_score",
		DefaultOutputPath: "healthScore",
		Column:            "health_score",
		SQLAlias:          "health_score",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"success_rate": {
		SourcePath:        "success_rate",
		DefaultOutputPath: "successRate",
		Column:            "success_rate",
		SQLAlias:          "success_rate",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"timeout_rate": {
		SourcePath:        "timeout_rate",
		DefaultOutputPath: "timeoutRate",
		Column:            "timeout_rate",
		SQLAlias:          "timeout_rate",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"error_rate": {
		SourcePath:        "error_rate",
		DefaultOutputPath: "errorRate",
		Column:            "error_rate",
		SQLAlias:          "error_rate",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"p95_latency_ms": {
		SourcePath:        "p95_latency_ms",
		DefaultOutputPath: "p95LatencyMs",
		Column:            "p95_latency_ms",
		SQLAlias:          "p95_latency_ms",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"sample_size": {
		SourcePath:        "sample_size",
		DefaultOutputPath: "sampleSize",
		Column:            "sample_size",
		SQLAlias:          "sample_size",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"window_started_at": {
		SourcePath:        "window_started_at",
		DefaultOutputPath: "windowStartedAt",
		Column:            "window_started_at",
		SQLAlias:          "window_started_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"window_ended_at": {
		SourcePath:        "window_ended_at",
		DefaultOutputPath: "windowEndedAt",
		Column:            "window_ended_at",
		SQLAlias:          "window_ended_at",
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

func NewProviderHealthSnapshotsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = ProviderHealthSnapshotsFilterFields[field]
	return
}

type ProviderHealthSnapshotsFilterResult struct {
	ProviderHealthSnapshots
	FilterCount int `db:"count"`
}

func ValidateProviderHealthSnapshotsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewProviderHealthSnapshotsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewProviderHealthSnapshotsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewProviderHealthSnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateProviderHealthSnapshotsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateProviderHealthSnapshotsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewProviderHealthSnapshotsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateProviderHealthSnapshotsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ProviderHealthSnapshots struct {
	Id                uuid.UUID           `db:"id"`
	ProviderAccountId uuid.UUID           `db:"provider_account_id"`
	MethodCode        null.String         `db:"method_code"`
	ChannelCode       null.String         `db:"channel_code"`
	HealthScore       int                 `db:"health_score"`
	SuccessRate       decimal.NullDecimal `db:"success_rate"`
	TimeoutRate       decimal.NullDecimal `db:"timeout_rate"`
	ErrorRate         decimal.NullDecimal `db:"error_rate"`
	P95LatencyMs      null.Int            `db:"p95_latency_ms"`
	SampleSize        int                 `db:"sample_size"`
	WindowStartedAt   time.Time           `db:"window_started_at"`
	WindowEndedAt     time.Time           `db:"window_ended_at"`
	Metadata          json.RawMessage     `db:"metadata"`

	shared.MetaSignature
}
type ProviderHealthSnapshotsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d ProviderHealthSnapshots) ToProviderHealthSnapshotsPrimaryID() ProviderHealthSnapshotsPrimaryID {
	return ProviderHealthSnapshotsPrimaryID{
		Id: d.Id,
	}
}

type ProviderHealthSnapshotsList []*ProviderHealthSnapshots

type ProviderHealthSnapshotsFilterResultList []*ProviderHealthSnapshotsFilterResult
