package model

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------------------------------------
// Model
// ---------------------------------------------------------------------------

type PayroutePspHealth struct {
	Id                  uuid.UUID       `db:"id"`
	ProviderAccountId   uuid.UUID       `db:"provider_account_id"`
	MethodCode          string          `db:"method_code"`
	ChannelCode         string          `db:"channel_code"`
	Currency            string          `db:"currency"`
	WindowStart         null.Time       `db:"window_start"`
	WindowEnd           null.Time       `db:"window_end"`
	TotalAttempts       int             `db:"total_attempts"`
	Successful          int             `db:"successful"`
	FailedSystem        int             `db:"failed_system"`
	FailedDecline       int             `db:"failed_decline"`
	SuccessRate         decimal.Decimal `db:"success_rate"`
	AvgLatencyMs        null.Int        `db:"avg_latency_ms"`
	P95LatencyMs        null.Int        `db:"p95_latency_ms"`
	SampleSize          int             `db:"sample_size"`

	MetaCreatedAt null.Time `db:"meta_created_at"`
}

type PayroutePspHealthPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayroutePspHealth) ToPayroutePspHealthPrimaryID() PayroutePspHealthPrimaryID {
	return PayroutePspHealthPrimaryID{Id: d.Id}
}

// IsDataSufficient returns true if the sample size meets the minimum threshold.
func (d PayroutePspHealth) IsDataSufficient(minSampleSize int) bool {
	return d.SampleSize >= minSampleSize
}

// ---------------------------------------------------------------------------
// DB Field Names
// ---------------------------------------------------------------------------

type PayroutePspHealthDBFieldNameType string

type payroutePspHealthDBFieldName struct {
	Id                PayroutePspHealthDBFieldNameType
	ProviderAccountId PayroutePspHealthDBFieldNameType
	MethodCode        PayroutePspHealthDBFieldNameType
	ChannelCode       PayroutePspHealthDBFieldNameType
	WindowStart       PayroutePspHealthDBFieldNameType
	SuccessRate       PayroutePspHealthDBFieldNameType
	SampleSize        PayroutePspHealthDBFieldNameType
	MetaCreatedAt     PayroutePspHealthDBFieldNameType
	MetaCreatedBy     PayroutePspHealthDBFieldNameType
	MetaUpdatedAt     PayroutePspHealthDBFieldNameType
	MetaUpdatedBy     PayroutePspHealthDBFieldNameType
	MetaDeletedAt     PayroutePspHealthDBFieldNameType
	MetaDeletedBy     PayroutePspHealthDBFieldNameType
}

var PayroutePspHealthDBFieldName = payroutePspHealthDBFieldName{
	Id:                "id",
	ProviderAccountId: "provider_account_id",
	MethodCode:        "method_code",
	ChannelCode:       "channel_code",
	WindowStart:       "window_start",
	SuccessRate:       "success_rate",
	SampleSize:        "sample_size",
	MetaCreatedAt:     "meta_created_at",
	MetaCreatedBy:     "meta_created_by",
	MetaUpdatedAt:     "meta_updated_at",
	MetaUpdatedBy:     "meta_updated_by",
	MetaDeletedAt:     "meta_deleted_at",
	MetaDeletedBy:     "meta_deleted_by",
}

func NewPayroutePspHealthDBFieldNameFromStr(field string) (dbField PayroutePspHealthDBFieldNameType, found bool) {
	switch field {
	case string(PayroutePspHealthDBFieldName.Id):
		return PayroutePspHealthDBFieldName.Id, true
	case string(PayroutePspHealthDBFieldName.ProviderAccountId):
		return PayroutePspHealthDBFieldName.ProviderAccountId, true
	case string(PayroutePspHealthDBFieldName.MethodCode):
		return PayroutePspHealthDBFieldName.MethodCode, true
	case string(PayroutePspHealthDBFieldName.WindowStart):
		return PayroutePspHealthDBFieldName.WindowStart, true
	case string(PayroutePspHealthDBFieldName.MetaCreatedAt):
		return PayroutePspHealthDBFieldName.MetaCreatedAt, true
	}
	return "unknown", false
}

// ---------------------------------------------------------------------------
// Filter & Query Support
// ---------------------------------------------------------------------------

var PayroutePspHealthFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath: "id", DefaultOutputPath: "id", Column: "id",
		SQLAlias: "id", Selectable: true, Filterable: true, Sortable: true,
	},
	"provider_account_id": {
		SourcePath: "provider_account_id", DefaultOutputPath: "providerAccountId", Column: "provider_account_id",
		SQLAlias: "provider_account_id", Selectable: true, Filterable: true, Sortable: true,
	},
	"method_code": {
		SourcePath: "method_code", DefaultOutputPath: "methodCode", Column: "method_code",
		SQLAlias: "method_code", Selectable: true, Filterable: true, Sortable: true,
	},
	"channel_code": {
		SourcePath: "channel_code", DefaultOutputPath: "channelCode", Column: "channel_code",
		SQLAlias: "channel_code", Selectable: true, Filterable: true, Sortable: true,
	},
	"success_rate": {
		SourcePath: "success_rate", DefaultOutputPath: "successRate", Column: "success_rate",
		SQLAlias: "success_rate", Selectable: true, Filterable: true, Sortable: true,
	},
	"sample_size": {
		SourcePath: "sample_size", DefaultOutputPath: "sampleSize", Column: "sample_size",
		SQLAlias: "sample_size", Selectable: true, Filterable: true, Sortable: true,
	},
	"window_start": {
		SourcePath: "window_start", DefaultOutputPath: "windowStart", Column: "window_start",
		SQLAlias: "window_start", Selectable: true, Filterable: true, Sortable: true,
	},
	"meta_created_at": {
		SourcePath: "meta_created_at", DefaultOutputPath: "metaCreatedAt", Column: "meta_created_at",
		SQLAlias: "meta_created_at", Selectable: true, Filterable: true, Sortable: true,
	},
}

func NewPayroutePspHealthFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayroutePspHealthFilterFields[field]
	return
}

func ValidatePayroutePspHealthFieldNameFilter(filter Filter) error {
	for _, field := range filter.FilterFields {
		spec, exist := NewPayroutePspHealthFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			return failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// List types
// ---------------------------------------------------------------------------

type PayroutePspHealthList []*PayroutePspHealth

type PayroutePspHealthFilterResult struct {
	PayroutePspHealth
	FilterCount int `db:"count"`
}

type PayroutePspHealthFilterResultList []*PayroutePspHealthFilterResult
