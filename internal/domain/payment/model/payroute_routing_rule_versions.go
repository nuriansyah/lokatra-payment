package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

// ---------------------------------------------------------------------------
// Model
// ---------------------------------------------------------------------------

type PayrouteRoutingRuleVersions struct {
	Id            uuid.UUID       `db:"id"`
	RuleId        uuid.UUID       `db:"rule_id"`
	Version       int             `db:"version"`
	Snapshot      json.RawMessage `db:"snapshot"`
	ChangedBy     uuid.UUID       `db:"changed_by"`
	ChangeReason  null.String     `db:"change_reason"`

	shared.MetaSignature
}

type PayrouteRoutingRuleVersionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayrouteRoutingRuleVersions) ToPayrouteRoutingRuleVersionsPrimaryID() PayrouteRoutingRuleVersionsPrimaryID {
	return PayrouteRoutingRuleVersionsPrimaryID{Id: d.Id}
}

// ---------------------------------------------------------------------------
// DB Field Names
// ---------------------------------------------------------------------------

type PayrouteRoutingRuleVersionsDBFieldNameType string

type payrouteRoutingRuleVersionsDBFieldName struct {
	Id            PayrouteRoutingRuleVersionsDBFieldNameType
	RuleId        PayrouteRoutingRuleVersionsDBFieldNameType
	Version       PayrouteRoutingRuleVersionsDBFieldNameType
	ChangedBy     PayrouteRoutingRuleVersionsDBFieldNameType
	ChangeReason  PayrouteRoutingRuleVersionsDBFieldNameType
	MetaCreatedAt PayrouteRoutingRuleVersionsDBFieldNameType
	MetaCreatedBy PayrouteRoutingRuleVersionsDBFieldNameType
	MetaUpdatedAt PayrouteRoutingRuleVersionsDBFieldNameType
	MetaUpdatedBy PayrouteRoutingRuleVersionsDBFieldNameType
	MetaDeletedAt PayrouteRoutingRuleVersionsDBFieldNameType
	MetaDeletedBy PayrouteRoutingRuleVersionsDBFieldNameType
}

var PayrouteRoutingRuleVersionsDBFieldName = payrouteRoutingRuleVersionsDBFieldName{
	Id:            "id",
	RuleId:        "rule_id",
	Version:       "version",
	ChangedBy:     "changed_by",
	ChangeReason:  "change_reason",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewPayrouteRoutingRuleVersionsDBFieldNameFromStr(field string) (dbField PayrouteRoutingRuleVersionsDBFieldNameType, found bool) {
	switch field {
	case string(PayrouteRoutingRuleVersionsDBFieldName.Id):
		return PayrouteRoutingRuleVersionsDBFieldName.Id, true
	case string(PayrouteRoutingRuleVersionsDBFieldName.RuleId):
		return PayrouteRoutingRuleVersionsDBFieldName.RuleId, true
	case string(PayrouteRoutingRuleVersionsDBFieldName.Version):
		return PayrouteRoutingRuleVersionsDBFieldName.Version, true
	case string(PayrouteRoutingRuleVersionsDBFieldName.MetaCreatedAt):
		return PayrouteRoutingRuleVersionsDBFieldName.MetaCreatedAt, true
	}
	return "unknown", false
}

// ---------------------------------------------------------------------------
// Filter & Query Support
// ---------------------------------------------------------------------------

var PayrouteRoutingRuleVersionsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath: "id", DefaultOutputPath: "id", Column: "id",
		SQLAlias: "id", Selectable: true, Filterable: true, Sortable: true,
	},
	"rule_id": {
		SourcePath: "rule_id", DefaultOutputPath: "ruleId", Column: "rule_id",
		SQLAlias: "rule_id", Selectable: true, Filterable: true, Sortable: true,
	},
	"version": {
		SourcePath: "version", DefaultOutputPath: "version", Column: "version",
		SQLAlias: "version", Selectable: true, Filterable: true, Sortable: true,
	},
	"meta_created_at": {
		SourcePath: "meta_created_at", DefaultOutputPath: "metaCreatedAt", Column: "meta_created_at",
		SQLAlias: "meta_created_at", Selectable: true, Filterable: true, Sortable: true,
	},
}

func NewPayrouteRoutingRuleVersionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayrouteRoutingRuleVersionsFilterFields[field]
	return
}

func ValidatePayrouteRoutingRuleVersionsFieldNameFilter(filter Filter) error {
	for _, field := range filter.FilterFields {
		spec, exist := NewPayrouteRoutingRuleVersionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			return failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// List types
// ---------------------------------------------------------------------------

type PayrouteRoutingRuleVersionsList []*PayrouteRoutingRuleVersions

type PayrouteRoutingRuleVersionsFilterResult struct {
	PayrouteRoutingRuleVersions
	FilterCount int `db:"count"`
}

type PayrouteRoutingRuleVersionsFilterResultList []*PayrouteRoutingRuleVersionsFilterResult
