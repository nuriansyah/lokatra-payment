package model

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

// ---------------------------------------------------------------------------
// Model
// ---------------------------------------------------------------------------

type PayrouteAuditLog struct {
	Id            uuid.UUID       `db:"id"`
	ActorUserId   uuid.UUID       `db:"actor_user_id"`
	ActorEmail    null.String     `db:"actor_email"`
	ActorRole     null.String     `db:"actor_role"`
	Action        string          `db:"action"`
	TargetEntity  string          `db:"target_entity"`
	TargetId      *uuid.UUID      `db:"target_id"`
	BeforeValue   json.RawMessage `db:"before_value"`
	AfterValue    json.RawMessage `db:"after_value"`
	Reason        null.String     `db:"reason"`
	IpAddress     *net.IP         `db:"ip_address"`
	UserAgent     null.String     `db:"user_agent"`

	MetaCreatedAt null.Time `db:"meta_created_at"`
}

type PayrouteAuditLogPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PayrouteAuditLog) ToPayrouteAuditLogPrimaryID() PayrouteAuditLogPrimaryID {
	return PayrouteAuditLogPrimaryID{Id: d.Id}
}

// ---------------------------------------------------------------------------
// DB Field Names
// ---------------------------------------------------------------------------

type PayrouteAuditLogDBFieldNameType string

type payrouteAuditLogDBFieldName struct {
	Id            PayrouteAuditLogDBFieldNameType
	ActorUserId   PayrouteAuditLogDBFieldNameType
	Action        PayrouteAuditLogDBFieldNameType
	TargetEntity  PayrouteAuditLogDBFieldNameType
	TargetId      PayrouteAuditLogDBFieldNameType
	MetaCreatedAt PayrouteAuditLogDBFieldNameType
	MetaCreatedBy PayrouteAuditLogDBFieldNameType
	MetaUpdatedAt PayrouteAuditLogDBFieldNameType
	MetaUpdatedBy PayrouteAuditLogDBFieldNameType
	MetaDeletedAt PayrouteAuditLogDBFieldNameType
	MetaDeletedBy PayrouteAuditLogDBFieldNameType
}

var PayrouteAuditLogDBFieldName = payrouteAuditLogDBFieldName{
	Id:            "id",
	ActorUserId:   "actor_user_id",
	Action:        "action",
	TargetEntity:  "target_entity",
	TargetId:      "target_id",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

func NewPayrouteAuditLogDBFieldNameFromStr(field string) (dbField PayrouteAuditLogDBFieldNameType, found bool) {
	switch field {
	case string(PayrouteAuditLogDBFieldName.Id):
		return PayrouteAuditLogDBFieldName.Id, true
	case string(PayrouteAuditLogDBFieldName.ActorUserId):
		return PayrouteAuditLogDBFieldName.ActorUserId, true
	case string(PayrouteAuditLogDBFieldName.Action):
		return PayrouteAuditLogDBFieldName.Action, true
	case string(PayrouteAuditLogDBFieldName.TargetEntity):
		return PayrouteAuditLogDBFieldName.TargetEntity, true
	case string(PayrouteAuditLogDBFieldName.MetaCreatedAt):
		return PayrouteAuditLogDBFieldName.MetaCreatedAt, true
	}
	return "unknown", false
}

// ---------------------------------------------------------------------------
// Filter & Query Support
// ---------------------------------------------------------------------------

var PayrouteAuditLogFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath: "id", DefaultOutputPath: "id", Column: "id",
		SQLAlias: "id", Selectable: true, Filterable: true, Sortable: true,
	},
	"actor_user_id": {
		SourcePath: "actor_user_id", DefaultOutputPath: "actorUserId", Column: "actor_user_id",
		SQLAlias: "actor_user_id", Selectable: true, Filterable: true, Sortable: true,
	},
	"action": {
		SourcePath: "action", DefaultOutputPath: "action", Column: "action",
		SQLAlias: "action", Selectable: true, Filterable: true, Sortable: true,
	},
	"target_entity": {
		SourcePath: "target_entity", DefaultOutputPath: "targetEntity", Column: "target_entity",
		SQLAlias: "target_entity", Selectable: true, Filterable: true, Sortable: true,
	},
	"target_id": {
		SourcePath: "target_id", DefaultOutputPath: "targetId", Column: "target_id",
		SQLAlias: "target_id", Selectable: true, Filterable: true, Sortable: true,
	},
	"meta_created_at": {
		SourcePath: "meta_created_at", DefaultOutputPath: "metaCreatedAt", Column: "meta_created_at",
		SQLAlias: "meta_created_at", Selectable: true, Filterable: true, Sortable: true,
	},
}

func NewPayrouteAuditLogFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PayrouteAuditLogFilterFields[field]
	return
}

func ValidatePayrouteAuditLogFieldNameFilter(filter Filter) error {
	for _, field := range filter.FilterFields {
		spec, exist := NewPayrouteAuditLogFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			return failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPayrouteAuditLogFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			return failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// List types
// ---------------------------------------------------------------------------

type PayrouteAuditLogList []*PayrouteAuditLog

type PayrouteAuditLogFilterResult struct {
	PayrouteAuditLog
	FilterCount int `db:"count"`
}

type PayrouteAuditLogFilterResultList []*PayrouteAuditLogFilterResult
