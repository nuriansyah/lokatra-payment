package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

// ---------------------------------------------------------------------------
// Model
// ---------------------------------------------------------------------------

type PayrouteConfig struct {
	Key           string          `db:"key"`
	Value         json.RawMessage `db:"value"`
	Description   null.String     `db:"description"`
	UpdatedBy     *uuid.UUID      `db:"updated_by"`
	MetaUpdatedAt null.Time       `db:"meta_updated_at"`
}

// GetString returns the value as a string, or empty string if not found.
func (d PayrouteConfig) GetString() string {
	var s string
	if err := json.Unmarshal(d.Value, &s); err != nil {
		return ""
	}
	return s
}

// GetInt returns the value as an int, or 0 if not found.
func (d PayrouteConfig) GetInt() int {
	var n int
	if err := json.Unmarshal(d.Value, &n); err != nil {
		return 0
	}
	return n
}

// GetFloat64 returns the value as a float64, or 0 if not found.
func (d PayrouteConfig) GetFloat64() float64 {
	var f float64
	if err := json.Unmarshal(d.Value, &f); err != nil {
		return 0
	}
	return f
}

// GetBool returns the value as a bool, or false if not found.
func (d PayrouteConfig) GetBool() bool {
	var b bool
	if err := json.Unmarshal(d.Value, &b); err != nil {
		return false
	}
	return b
}

// GetJSONSlice returns the value as a JSON array of strings.
func (d PayrouteConfig) GetJSONSlice() []string {
	var items []string
	if err := json.Unmarshal(d.Value, &items); err != nil {
		return nil
	}
	return items
}

// GetUUIDSlice returns the value as a JSON array of UUID strings.
func (d PayrouteConfig) GetUUIDSlice() []string {
	return d.GetJSONSlice()
}

// ---------------------------------------------------------------------------
// DB Field Names
// ---------------------------------------------------------------------------

type PayrouteConfigDBFieldNameType string

type payrouteConfigDBFieldName struct {
	Key           PayrouteConfigDBFieldNameType
	MetaCreatedAt PayrouteConfigDBFieldNameType
	MetaCreatedBy PayrouteConfigDBFieldNameType
	MetaUpdatedAt PayrouteConfigDBFieldNameType
	MetaUpdatedBy PayrouteConfigDBFieldNameType
	MetaDeletedAt PayrouteConfigDBFieldNameType
	MetaDeletedBy PayrouteConfigDBFieldNameType
}

var PayrouteConfigDBFieldName = payrouteConfigDBFieldName{
	Key:           "key",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewPayrouteConfigDBFieldNameFromStr(field string) (dbField PayrouteConfigDBFieldNameType, found bool) {
	switch field {
	case string(PayrouteConfigDBFieldName.Key):
		return PayrouteConfigDBFieldName.Key, true
	case string(PayrouteConfigDBFieldName.MetaUpdatedAt):
		return PayrouteConfigDBFieldName.MetaUpdatedAt, true
	}
	return "unknown", false
}

// ---------------------------------------------------------------------------
// Filter & Query Support
// ---------------------------------------------------------------------------

var PayrouteConfigFilterFields = map[string]FilterFieldSpec{
	"key": {
		SourcePath: "key", DefaultOutputPath: "key", Column: "key",
		SQLAlias: "key", Selectable: true, Filterable: true, Sortable: true,
	},
	"meta_updated_at": {
		SourcePath: "meta_updated_at", DefaultOutputPath: "metaUpdatedAt", Column: "meta_updated_at",
		SQLAlias: "meta_updated_at", Selectable: true, Filterable: true, Sortable: true,
	},
}

func NewPayrouteConfigFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayrouteConfigFilterFields[field]
	return
}

func ValidatePayrouteConfigFieldNameFilter(filter Filter) error {
	for _, field := range filter.FilterFields {
		spec, exist := NewPayrouteConfigFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			return failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// List types
// ---------------------------------------------------------------------------

type PayrouteConfigList []*PayrouteConfig

type PayrouteConfigFilterResult struct {
	PayrouteConfig
	FilterCount int `db:"count"`
}

type PayrouteConfigFilterResultList []*PayrouteConfigFilterResult
