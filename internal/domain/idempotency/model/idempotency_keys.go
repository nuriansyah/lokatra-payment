package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
)

type IdempotencyKeysDBFieldNameType string

type idempotencyKeysDBFieldName struct {
	Id              IdempotencyKeysDBFieldNameType
	IdempotencyKey  IdempotencyKeysDBFieldNameType
	MerchantId      IdempotencyKeysDBFieldNameType
	RequestPath     IdempotencyKeysDBFieldNameType
	RequestBodyHash IdempotencyKeysDBFieldNameType
	ResponseStatus  IdempotencyKeysDBFieldNameType
	ResponseBody    IdempotencyKeysDBFieldNameType
	LockedAt        IdempotencyKeysDBFieldNameType
	LockedUntil     IdempotencyKeysDBFieldNameType
	CompletedAt     IdempotencyKeysDBFieldNameType
	ExpiresAt       IdempotencyKeysDBFieldNameType
	MetaCreatedAt   IdempotencyKeysDBFieldNameType
	MetaCreatedBy   IdempotencyKeysDBFieldNameType
	MetaUpdatedAt   IdempotencyKeysDBFieldNameType
	MetaUpdatedBy   IdempotencyKeysDBFieldNameType
	MetaDeletedAt   IdempotencyKeysDBFieldNameType
	MetaDeletedBy   IdempotencyKeysDBFieldNameType
}

var IdempotencyKeysDBFieldName = idempotencyKeysDBFieldName{
	Id:              "id",
	IdempotencyKey:  "idempotency_key",
	MerchantId:      "merchant_id",
	RequestPath:     "request_path",
	RequestBodyHash: "request_body_hash",
	ResponseStatus:  "response_status",
	ResponseBody:    "response_body",
	LockedAt:        "locked_at",
	LockedUntil:     "locked_until",
	CompletedAt:     "completed_at",
	ExpiresAt:       "expires_at",
	MetaCreatedAt:   "meta_created_at",
	MetaCreatedBy:   "meta_created_by",
	MetaUpdatedAt:   "meta_updated_at",
	MetaUpdatedBy:   "meta_updated_by",
	MetaDeletedAt:   "meta_deleted_at",
	MetaDeletedBy:   "meta_deleted_by",
}

func NewIdempotencyKeysDBFieldNameFromStr(field string) (dbField IdempotencyKeysDBFieldNameType, found bool) {
	switch field {

	case string(IdempotencyKeysDBFieldName.Id):
		return IdempotencyKeysDBFieldName.Id, true

	case string(IdempotencyKeysDBFieldName.IdempotencyKey):
		return IdempotencyKeysDBFieldName.IdempotencyKey, true

	case string(IdempotencyKeysDBFieldName.MerchantId):
		return IdempotencyKeysDBFieldName.MerchantId, true

	case string(IdempotencyKeysDBFieldName.RequestPath):
		return IdempotencyKeysDBFieldName.RequestPath, true

	case string(IdempotencyKeysDBFieldName.RequestBodyHash):
		return IdempotencyKeysDBFieldName.RequestBodyHash, true

	case string(IdempotencyKeysDBFieldName.ResponseStatus):
		return IdempotencyKeysDBFieldName.ResponseStatus, true

	case string(IdempotencyKeysDBFieldName.ResponseBody):
		return IdempotencyKeysDBFieldName.ResponseBody, true

	case string(IdempotencyKeysDBFieldName.LockedAt):
		return IdempotencyKeysDBFieldName.LockedAt, true

	case string(IdempotencyKeysDBFieldName.LockedUntil):
		return IdempotencyKeysDBFieldName.LockedUntil, true

	case string(IdempotencyKeysDBFieldName.CompletedAt):
		return IdempotencyKeysDBFieldName.CompletedAt, true

	case string(IdempotencyKeysDBFieldName.ExpiresAt):
		return IdempotencyKeysDBFieldName.ExpiresAt, true

	case string(IdempotencyKeysDBFieldName.MetaCreatedAt):
		return IdempotencyKeysDBFieldName.MetaCreatedAt, true

	case string(IdempotencyKeysDBFieldName.MetaCreatedBy):
		return IdempotencyKeysDBFieldName.MetaCreatedBy, true

	case string(IdempotencyKeysDBFieldName.MetaUpdatedAt):
		return IdempotencyKeysDBFieldName.MetaUpdatedAt, true

	case string(IdempotencyKeysDBFieldName.MetaUpdatedBy):
		return IdempotencyKeysDBFieldName.MetaUpdatedBy, true

	case string(IdempotencyKeysDBFieldName.MetaDeletedAt):
		return IdempotencyKeysDBFieldName.MetaDeletedAt, true

	case string(IdempotencyKeysDBFieldName.MetaDeletedBy):
		return IdempotencyKeysDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

type IdempotencyKeysFilterResult struct {
	IdempotencyKeys
	FilterCount int `db:"count"`
}

func ValidateIdempotencyKeysFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		_, exist := NewIdempotencyKeysDBFieldNameFromStr(selectField)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", selectField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		_, exist := NewIdempotencyKeysDBFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		_, exist := NewIdempotencyKeysDBFieldNameFromStr(field.Field)
		if !exist {
			err = failure.InternalError(fmt.Errorf("field %s is not found", field.Field))
			return
		}
	}
	return
}

type IdempotencyKeys struct {
	Id              uuid.UUID       `db:"id"`
	IdempotencyKey  string          `db:"idempotency_key"`
	MerchantId      uuid.UUID       `db:"merchant_id"`
	RequestPath     string          `db:"request_path"`
	RequestBodyHash string          `db:"request_body_hash"`
	ResponseStatus  null.Int        `db:"response_status"`
	ResponseBody    json.RawMessage `db:"response_body"`
	LockedAt        null.Time       `db:"locked_at"`
	LockedUntil     null.Time       `db:"locked_until"`
	CompletedAt     null.Time       `db:"completed_at"`
	ExpiresAt       time.Time       `db:"expires_at"`
	MetaCreatedAt   time.Time       `db:"meta_created_at"`
	MetaCreatedBy   uuid.UUID       `db:"meta_created_by"`
	MetaUpdatedAt   time.Time       `db:"meta_updated_at"`
	MetaUpdatedBy   nuuid.NUUID     `db:"meta_updated_by"`
	MetaDeletedAt   null.Time       `db:"meta_deleted_at"`
	MetaDeletedBy   nuuid.NUUID     `db:"meta_deleted_by"`
}
type IdempotencyKeysPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d IdempotencyKeys) ToIdempotencyKeysPrimaryID() IdempotencyKeysPrimaryID {
	return IdempotencyKeysPrimaryID{
		Id: d.Id,
	}
}

type IdempotencyKeysList []*IdempotencyKeys
