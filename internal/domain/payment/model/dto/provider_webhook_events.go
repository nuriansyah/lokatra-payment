package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ProviderWebhookEventsDTOFieldNameType string

type providerWebhookEventsDTOFieldName struct {
	Id                 ProviderWebhookEventsDTOFieldNameType
	WebhookEndpointId  ProviderWebhookEventsDTOFieldNameType
	EndpointKey        ProviderWebhookEventsDTOFieldNameType
	ProviderAccountId  ProviderWebhookEventsDTOFieldNameType
	ProviderCode       ProviderWebhookEventsDTOFieldNameType
	EventId            ProviderWebhookEventsDTOFieldNameType
	EventType          ProviderWebhookEventsDTOFieldNameType
	ProviderReference  ProviderWebhookEventsDTOFieldNameType
	ProviderStatus     ProviderWebhookEventsDTOFieldNameType
	SignatureValid     ProviderWebhookEventsDTOFieldNameType
	SignatureAlgorithm ProviderWebhookEventsDTOFieldNameType
	Headers            ProviderWebhookEventsDTOFieldNameType
	RawBody            ProviderWebhookEventsDTOFieldNameType
	RawBodySha256      ProviderWebhookEventsDTOFieldNameType
	ParsedBody         ProviderWebhookEventsDTOFieldNameType
	ProcessingStatus   ProviderWebhookEventsDTOFieldNameType
	RetryCount         ProviderWebhookEventsDTOFieldNameType
	NextRetryAt        ProviderWebhookEventsDTOFieldNameType
	LockedUntil        ProviderWebhookEventsDTOFieldNameType
	ReceivedAt         ProviderWebhookEventsDTOFieldNameType
	ProcessedAt        ProviderWebhookEventsDTOFieldNameType
	ErrorCode          ProviderWebhookEventsDTOFieldNameType
	ErrorMessage       ProviderWebhookEventsDTOFieldNameType
	Metadata           ProviderWebhookEventsDTOFieldNameType
	MetaCreatedAt      ProviderWebhookEventsDTOFieldNameType
	MetaCreatedBy      ProviderWebhookEventsDTOFieldNameType
	MetaUpdatedAt      ProviderWebhookEventsDTOFieldNameType
	MetaUpdatedBy      ProviderWebhookEventsDTOFieldNameType
	MetaDeletedAt      ProviderWebhookEventsDTOFieldNameType
	MetaDeletedBy      ProviderWebhookEventsDTOFieldNameType
}

var ProviderWebhookEventsDTOFieldName = providerWebhookEventsDTOFieldName{
	Id:                 "id",
	WebhookEndpointId:  "webhookEndpointId",
	EndpointKey:        "endpointKey",
	ProviderAccountId:  "providerAccountId",
	ProviderCode:       "providerCode",
	EventId:            "eventId",
	EventType:          "eventType",
	ProviderReference:  "providerReference",
	ProviderStatus:     "providerStatus",
	SignatureValid:     "signatureValid",
	SignatureAlgorithm: "signatureAlgorithm",
	Headers:            "headers",
	RawBody:            "rawBody",
	RawBodySha256:      "rawBodySha256",
	ParsedBody:         "parsedBody",
	ProcessingStatus:   "processingStatus",
	RetryCount:         "retryCount",
	NextRetryAt:        "nextRetryAt",
	LockedUntil:        "lockedUntil",
	ReceivedAt:         "receivedAt",
	ProcessedAt:        "processedAt",
	ErrorCode:          "errorCode",
	ErrorMessage:       "errorMessage",
	Metadata:           "metadata",
	MetaCreatedAt:      "metaCreatedAt",
	MetaCreatedBy:      "metaCreatedBy",
	MetaUpdatedAt:      "metaUpdatedAt",
	MetaUpdatedBy:      "metaUpdatedBy",
	MetaDeletedAt:      "metaDeletedAt",
	MetaDeletedBy:      "metaDeletedBy",
}

func transformProviderWebhookEventsDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ProviderWebhookEventsDTOFieldName.Id):
		return string(model.ProviderWebhookEventsDBFieldName.Id), true

	case string(ProviderWebhookEventsDTOFieldName.WebhookEndpointId):
		return string(model.ProviderWebhookEventsDBFieldName.WebhookEndpointId), true

	case string(ProviderWebhookEventsDTOFieldName.EndpointKey):
		return string(model.ProviderWebhookEventsDBFieldName.EndpointKey), true

	case string(ProviderWebhookEventsDTOFieldName.ProviderAccountId):
		return string(model.ProviderWebhookEventsDBFieldName.ProviderAccountId), true

	case string(ProviderWebhookEventsDTOFieldName.ProviderCode):
		return string(model.ProviderWebhookEventsDBFieldName.ProviderCode), true

	case string(ProviderWebhookEventsDTOFieldName.EventId):
		return string(model.ProviderWebhookEventsDBFieldName.EventId), true

	case string(ProviderWebhookEventsDTOFieldName.EventType):
		return string(model.ProviderWebhookEventsDBFieldName.EventType), true

	case string(ProviderWebhookEventsDTOFieldName.ProviderReference):
		return string(model.ProviderWebhookEventsDBFieldName.ProviderReference), true

	case string(ProviderWebhookEventsDTOFieldName.ProviderStatus):
		return string(model.ProviderWebhookEventsDBFieldName.ProviderStatus), true

	case string(ProviderWebhookEventsDTOFieldName.SignatureValid):
		return string(model.ProviderWebhookEventsDBFieldName.SignatureValid), true

	case string(ProviderWebhookEventsDTOFieldName.SignatureAlgorithm):
		return string(model.ProviderWebhookEventsDBFieldName.SignatureAlgorithm), true

	case string(ProviderWebhookEventsDTOFieldName.Headers):
		return string(model.ProviderWebhookEventsDBFieldName.Headers), true

	case string(ProviderWebhookEventsDTOFieldName.RawBody):
		return string(model.ProviderWebhookEventsDBFieldName.RawBody), true

	case string(ProviderWebhookEventsDTOFieldName.RawBodySha256):
		return string(model.ProviderWebhookEventsDBFieldName.RawBodySha256), true

	case string(ProviderWebhookEventsDTOFieldName.ParsedBody):
		return string(model.ProviderWebhookEventsDBFieldName.ParsedBody), true

	case string(ProviderWebhookEventsDTOFieldName.ProcessingStatus):
		return string(model.ProviderWebhookEventsDBFieldName.ProcessingStatus), true

	case string(ProviderWebhookEventsDTOFieldName.RetryCount):
		return string(model.ProviderWebhookEventsDBFieldName.RetryCount), true

	case string(ProviderWebhookEventsDTOFieldName.NextRetryAt):
		return string(model.ProviderWebhookEventsDBFieldName.NextRetryAt), true

	case string(ProviderWebhookEventsDTOFieldName.LockedUntil):
		return string(model.ProviderWebhookEventsDBFieldName.LockedUntil), true

	case string(ProviderWebhookEventsDTOFieldName.ReceivedAt):
		return string(model.ProviderWebhookEventsDBFieldName.ReceivedAt), true

	case string(ProviderWebhookEventsDTOFieldName.ProcessedAt):
		return string(model.ProviderWebhookEventsDBFieldName.ProcessedAt), true

	case string(ProviderWebhookEventsDTOFieldName.ErrorCode):
		return string(model.ProviderWebhookEventsDBFieldName.ErrorCode), true

	case string(ProviderWebhookEventsDTOFieldName.ErrorMessage):
		return string(model.ProviderWebhookEventsDBFieldName.ErrorMessage), true

	case string(ProviderWebhookEventsDTOFieldName.Metadata):
		return string(model.ProviderWebhookEventsDBFieldName.Metadata), true

	case string(ProviderWebhookEventsDTOFieldName.MetaCreatedAt):
		return string(model.ProviderWebhookEventsDBFieldName.MetaCreatedAt), true

	case string(ProviderWebhookEventsDTOFieldName.MetaCreatedBy):
		return string(model.ProviderWebhookEventsDBFieldName.MetaCreatedBy), true

	case string(ProviderWebhookEventsDTOFieldName.MetaUpdatedAt):
		return string(model.ProviderWebhookEventsDBFieldName.MetaUpdatedAt), true

	case string(ProviderWebhookEventsDTOFieldName.MetaUpdatedBy):
		return string(model.ProviderWebhookEventsDBFieldName.MetaUpdatedBy), true

	case string(ProviderWebhookEventsDTOFieldName.MetaDeletedAt):
		return string(model.ProviderWebhookEventsDBFieldName.MetaDeletedAt), true

	case string(ProviderWebhookEventsDTOFieldName.MetaDeletedBy):
		return string(model.ProviderWebhookEventsDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewProviderWebhookEventsFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isProviderWebhookEventsBaseFilterField(field string) bool {
	spec, found := model.NewProviderWebhookEventsFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeProviderWebhookEventsProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateProviderWebhookEventsProjectionOutputPath(path string) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return failure.BadRequest(fmt.Errorf("field alias cannot be empty"))
	}
	if !strings.Contains(path, ".") {
		return nil
	}
	for _, part := range strings.Split(path, ".") {
		if strings.TrimSpace(part) == "" {
			return failure.BadRequest(fmt.Errorf("field alias %s is invalid", path))
		}
	}
	return nil
}

func transformProviderWebhookEventsFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformProviderWebhookEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformProviderWebhookEventsFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformProviderWebhookEventsFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformProviderWebhookEventsDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isProviderWebhookEventsBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateProviderWebhookEventsProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeProviderWebhookEventsProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformProviderWebhookEventsDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformProviderWebhookEventsDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformProviderWebhookEventsFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultProviderWebhookEventsFilter(filter *model.Filter) {
	if filter.Pagination.Strategy == "" {
		if filter.Pagination.Page > 0 {
			filter.Pagination.Strategy = model.PaginationStrategyOffset
		} else {
			filter.Pagination.Strategy = model.DefaultPaginationStrategy
		}
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = model.DefaultPageSize
	}

	if filter.Pagination.PageSize > model.MaxPageSize {
		filter.Pagination.PageSize = model.MaxPageSize
	}

	if filter.Pagination.Strategy == model.PaginationStrategyOffset && filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.Direction == "" {
		filter.Pagination.Direction = model.CursorDirectionNext
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(ProviderWebhookEventsDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ProviderWebhookEventsSelectableResponse map[string]interface{}
type ProviderWebhookEventsSelectableListResponse []*ProviderWebhookEventsSelectableResponse

func assignProviderWebhookEventsNestedValue(out map[string]interface{}, path string, value interface{}) {
	parts := strings.Split(path, ".")
	current := out
	for _, part := range parts[:len(parts)-1] {
		next, ok := current[part].(map[string]interface{})
		if !ok {
			next = map[string]interface{}{}
			current[part] = next
		}
		current = next
	}
	current[parts[len(parts)-1]] = value
}

func setProviderWebhookEventsSelectableValue(out ProviderWebhookEventsSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignProviderWebhookEventsNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewProviderWebhookEventsSelectableResponse(providerWebhookEvents model.ProviderWebhookEvents, filter model.Filter) ProviderWebhookEventsSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ProviderWebhookEventsDBFieldName.Id),
			string(model.ProviderWebhookEventsDBFieldName.WebhookEndpointId),
			string(model.ProviderWebhookEventsDBFieldName.EndpointKey),
			string(model.ProviderWebhookEventsDBFieldName.ProviderAccountId),
			string(model.ProviderWebhookEventsDBFieldName.ProviderCode),
			string(model.ProviderWebhookEventsDBFieldName.EventId),
			string(model.ProviderWebhookEventsDBFieldName.EventType),
			string(model.ProviderWebhookEventsDBFieldName.ProviderReference),
			string(model.ProviderWebhookEventsDBFieldName.ProviderStatus),
			string(model.ProviderWebhookEventsDBFieldName.SignatureValid),
			string(model.ProviderWebhookEventsDBFieldName.SignatureAlgorithm),
			string(model.ProviderWebhookEventsDBFieldName.Headers),
			string(model.ProviderWebhookEventsDBFieldName.RawBody),
			string(model.ProviderWebhookEventsDBFieldName.RawBodySha256),
			string(model.ProviderWebhookEventsDBFieldName.ParsedBody),
			string(model.ProviderWebhookEventsDBFieldName.ProcessingStatus),
			string(model.ProviderWebhookEventsDBFieldName.RetryCount),
			string(model.ProviderWebhookEventsDBFieldName.NextRetryAt),
			string(model.ProviderWebhookEventsDBFieldName.LockedUntil),
			string(model.ProviderWebhookEventsDBFieldName.ReceivedAt),
			string(model.ProviderWebhookEventsDBFieldName.ProcessedAt),
			string(model.ProviderWebhookEventsDBFieldName.ErrorCode),
			string(model.ProviderWebhookEventsDBFieldName.ErrorMessage),
			string(model.ProviderWebhookEventsDBFieldName.Metadata),
			string(model.ProviderWebhookEventsDBFieldName.MetaCreatedAt),
			string(model.ProviderWebhookEventsDBFieldName.MetaCreatedBy),
			string(model.ProviderWebhookEventsDBFieldName.MetaUpdatedAt),
			string(model.ProviderWebhookEventsDBFieldName.MetaUpdatedBy),
			string(model.ProviderWebhookEventsDBFieldName.MetaDeletedAt),
			string(model.ProviderWebhookEventsDBFieldName.MetaDeletedBy),
		)
	}
	providerWebhookEventsSelectableResponse := ProviderWebhookEventsSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ProviderWebhookEventsDBFieldName.Id):
			key := string(ProviderWebhookEventsDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.Id, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.WebhookEndpointId):
			key := string(ProviderWebhookEventsDTOFieldName.WebhookEndpointId)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.WebhookEndpointId.UUID, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.EndpointKey):
			key := string(ProviderWebhookEventsDTOFieldName.EndpointKey)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.EndpointKey.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ProviderAccountId):
			key := string(ProviderWebhookEventsDTOFieldName.ProviderAccountId)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ProviderAccountId, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ProviderCode):
			key := string(ProviderWebhookEventsDTOFieldName.ProviderCode)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ProviderCode, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.EventId):
			key := string(ProviderWebhookEventsDTOFieldName.EventId)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.EventId.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.EventType):
			key := string(ProviderWebhookEventsDTOFieldName.EventType)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.EventType.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ProviderReference):
			key := string(ProviderWebhookEventsDTOFieldName.ProviderReference)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ProviderReference.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ProviderStatus):
			key := string(ProviderWebhookEventsDTOFieldName.ProviderStatus)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ProviderStatus.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.SignatureValid):
			key := string(ProviderWebhookEventsDTOFieldName.SignatureValid)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.SignatureValid, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.SignatureAlgorithm):
			key := string(ProviderWebhookEventsDTOFieldName.SignatureAlgorithm)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.SignatureAlgorithm.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.Headers):
			key := string(ProviderWebhookEventsDTOFieldName.Headers)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.Headers, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.RawBody):
			key := string(ProviderWebhookEventsDTOFieldName.RawBody)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.RawBody, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.RawBodySha256):
			key := string(ProviderWebhookEventsDTOFieldName.RawBodySha256)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.RawBodySha256, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ParsedBody):
			key := string(ProviderWebhookEventsDTOFieldName.ParsedBody)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ParsedBody, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ProcessingStatus):
			key := string(ProviderWebhookEventsDTOFieldName.ProcessingStatus)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, model.WebhookProcessingStatus(providerWebhookEvents.ProcessingStatus), explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.RetryCount):
			key := string(ProviderWebhookEventsDTOFieldName.RetryCount)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.RetryCount, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.NextRetryAt):
			key := string(ProviderWebhookEventsDTOFieldName.NextRetryAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.NextRetryAt.Time, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.LockedUntil):
			key := string(ProviderWebhookEventsDTOFieldName.LockedUntil)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.LockedUntil.Time, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ReceivedAt):
			key := string(ProviderWebhookEventsDTOFieldName.ReceivedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ReceivedAt, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ProcessedAt):
			key := string(ProviderWebhookEventsDTOFieldName.ProcessedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ProcessedAt.Time, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ErrorCode):
			key := string(ProviderWebhookEventsDTOFieldName.ErrorCode)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ErrorCode.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.ErrorMessage):
			key := string(ProviderWebhookEventsDTOFieldName.ErrorMessage)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.ErrorMessage.String, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.Metadata):
			key := string(ProviderWebhookEventsDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.Metadata, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.MetaCreatedAt):
			key := string(ProviderWebhookEventsDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.MetaCreatedAt, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.MetaCreatedBy):
			key := string(ProviderWebhookEventsDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.MetaCreatedBy, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.MetaUpdatedAt):
			key := string(ProviderWebhookEventsDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.MetaUpdatedAt, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.MetaUpdatedBy):
			key := string(ProviderWebhookEventsDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.MetaUpdatedBy.UUID, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.MetaDeletedAt):
			key := string(ProviderWebhookEventsDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.MetaDeletedAt.Time, explicitAlias)

		case string(model.ProviderWebhookEventsDBFieldName.MetaDeletedBy):
			key := string(ProviderWebhookEventsDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setProviderWebhookEventsSelectableValue(providerWebhookEventsSelectableResponse, key, providerWebhookEvents.MetaDeletedBy.UUID, explicitAlias)

		}
	}
	return providerWebhookEventsSelectableResponse
}

func NewProviderWebhookEventsListResponseFromFilterResult(result []model.ProviderWebhookEventsFilterResult, filter model.Filter) ProviderWebhookEventsSelectableListResponse {
	dtoProviderWebhookEventsListResponse := ProviderWebhookEventsSelectableListResponse{}
	for _, row := range result {
		dtoProviderWebhookEventsResponse := NewProviderWebhookEventsSelectableResponse(row.ProviderWebhookEvents, filter)
		dtoProviderWebhookEventsListResponse = append(dtoProviderWebhookEventsListResponse, &dtoProviderWebhookEventsResponse)
	}
	return dtoProviderWebhookEventsListResponse
}

type ProviderWebhookEventsFilterResponse struct {
	Metadata Metadata                                    `json:"metadata"`
	Data     ProviderWebhookEventsSelectableListResponse `json:"data"`
}

func reverseProviderWebhookEventsFilterResults(result []model.ProviderWebhookEventsFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewProviderWebhookEventsFilterResponse(result []model.ProviderWebhookEventsFilterResult, filter model.Filter) (resp ProviderWebhookEventsFilterResponse) {
	resp.Metadata.Strategy = filter.Pagination.Strategy
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	dataResult := result

	if filter.Pagination.IsCursorMode() && filter.Pagination.PageSize > 0 {
		if len(dataResult) > filter.Pagination.PageSize {
			resp.Metadata.HasMore = true
			if filter.Pagination.Direction == model.CursorDirectionPrev {
				resp.Metadata.HasPrev = true
			} else {
				resp.Metadata.HasNext = true
			}
			dataResult = dataResult[:filter.Pagination.PageSize]
		}
		if filter.Pagination.Direction == model.CursorDirectionPrev {
			reverseProviderWebhookEventsFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].Id
			resp.Metadata.PrevCursor = dataResult[0].Id
		}
		if resp.Metadata.Page <= 0 {
			resp.Metadata.Page = 1
		}
	} else {
		if len(dataResult) > 0 {
			resp.Metadata.TotalData = dataResult[0].FilterCount
			resp.Metadata.TotalPage = int(math.Ceil(float64(resp.Metadata.TotalData) / float64(filter.Pagination.PageSize)))
			resp.Metadata.HasPrev = filter.Pagination.Page > 1
			resp.Metadata.HasNext = filter.Pagination.Page < resp.Metadata.TotalPage
			resp.Metadata.HasMore = resp.Metadata.HasNext
		}
	}

	resp.Data = NewProviderWebhookEventsListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ProviderWebhookEventsCreateRequest struct {
	WebhookEndpointId  uuid.UUID                     `json:"webhookEndpointId"`
	EndpointKey        string                        `json:"endpointKey"`
	ProviderAccountId  uuid.UUID                     `json:"providerAccountId"`
	ProviderCode       string                        `json:"providerCode"`
	EventId            string                        `json:"eventId"`
	EventType          string                        `json:"eventType"`
	ProviderReference  string                        `json:"providerReference"`
	ProviderStatus     string                        `json:"providerStatus"`
	SignatureValid     bool                          `json:"signatureValid"`
	SignatureAlgorithm string                        `json:"signatureAlgorithm"`
	Headers            json.RawMessage               `json:"headers"`
	RawBody            []byte                        `json:"rawBody"`
	RawBodySha256      string                        `json:"rawBodySha256"`
	ParsedBody         json.RawMessage               `json:"parsedBody"`
	ProcessingStatus   model.WebhookProcessingStatus `json:"processingStatus" example:"received" enums:"received,processing,processed,failed"`
	RetryCount         int                           `json:"retryCount"`
	NextRetryAt        time.Time                     `json:"nextRetryAt"`
	LockedUntil        time.Time                     `json:"lockedUntil"`
	ReceivedAt         time.Time                     `json:"receivedAt"`
	ProcessedAt        time.Time                     `json:"processedAt"`
	ErrorCode          string                        `json:"errorCode"`
	ErrorMessage       string                        `json:"errorMessage"`
	Metadata           json.RawMessage               `json:"metadata"`
}

func (d *ProviderWebhookEventsCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ProviderWebhookEventsCreateRequest) ToModel() model.ProviderWebhookEvents {
	id, _ := uuid.NewV4()
	return model.ProviderWebhookEvents{
		Id:                 id,
		WebhookEndpointId:  nuuid.From(d.WebhookEndpointId),
		EndpointKey:        null.StringFrom(d.EndpointKey),
		ProviderAccountId:  d.ProviderAccountId,
		ProviderCode:       d.ProviderCode,
		EventId:            null.StringFrom(d.EventId),
		EventType:          null.StringFrom(d.EventType),
		ProviderReference:  null.StringFrom(d.ProviderReference),
		ProviderStatus:     null.StringFrom(d.ProviderStatus),
		SignatureValid:     d.SignatureValid,
		SignatureAlgorithm: null.StringFrom(d.SignatureAlgorithm),
		Headers:            d.Headers,
		RawBody:            d.RawBody,
		RawBodySha256:      d.RawBodySha256,
		ParsedBody:         d.ParsedBody,
		ProcessingStatus:   d.ProcessingStatus,
		RetryCount:         d.RetryCount,
		NextRetryAt:        null.TimeFrom(d.NextRetryAt),
		LockedUntil:        null.TimeFrom(d.LockedUntil),
		ReceivedAt:         d.ReceivedAt,
		ProcessedAt:        null.TimeFrom(d.ProcessedAt),
		ErrorCode:          null.StringFrom(d.ErrorCode),
		ErrorMessage:       null.StringFrom(d.ErrorMessage),
		Metadata:           d.Metadata,
	}
}

type ProviderWebhookEventsListCreateRequest []*ProviderWebhookEventsCreateRequest

func (d ProviderWebhookEventsListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerWebhookEvents := range d {
		err = validator.Struct(providerWebhookEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderWebhookEventsListCreateRequest) ToModelList() []model.ProviderWebhookEvents {
	out := make([]model.ProviderWebhookEvents, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ProviderWebhookEventsUpdateRequest struct {
	WebhookEndpointId  uuid.UUID                     `json:"webhookEndpointId"`
	EndpointKey        string                        `json:"endpointKey"`
	ProviderAccountId  uuid.UUID                     `json:"providerAccountId"`
	ProviderCode       string                        `json:"providerCode"`
	EventId            string                        `json:"eventId"`
	EventType          string                        `json:"eventType"`
	ProviderReference  string                        `json:"providerReference"`
	ProviderStatus     string                        `json:"providerStatus"`
	SignatureValid     bool                          `json:"signatureValid"`
	SignatureAlgorithm string                        `json:"signatureAlgorithm"`
	Headers            json.RawMessage               `json:"headers"`
	RawBody            []byte                        `json:"rawBody"`
	RawBodySha256      string                        `json:"rawBodySha256"`
	ParsedBody         json.RawMessage               `json:"parsedBody"`
	ProcessingStatus   model.WebhookProcessingStatus `json:"processingStatus" example:"received" enums:"received,processing,processed,failed"`
	RetryCount         int                           `json:"retryCount"`
	NextRetryAt        time.Time                     `json:"nextRetryAt"`
	LockedUntil        time.Time                     `json:"lockedUntil"`
	ReceivedAt         time.Time                     `json:"receivedAt"`
	ProcessedAt        time.Time                     `json:"processedAt"`
	ErrorCode          string                        `json:"errorCode"`
	ErrorMessage       string                        `json:"errorMessage"`
	Metadata           json.RawMessage               `json:"metadata"`
}

func (d *ProviderWebhookEventsUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ProviderWebhookEventsUpdateRequest) ToModel() model.ProviderWebhookEvents {
	return model.ProviderWebhookEvents{
		WebhookEndpointId:  nuuid.From(d.WebhookEndpointId),
		EndpointKey:        null.StringFrom(d.EndpointKey),
		ProviderAccountId:  d.ProviderAccountId,
		ProviderCode:       d.ProviderCode,
		EventId:            null.StringFrom(d.EventId),
		EventType:          null.StringFrom(d.EventType),
		ProviderReference:  null.StringFrom(d.ProviderReference),
		ProviderStatus:     null.StringFrom(d.ProviderStatus),
		SignatureValid:     d.SignatureValid,
		SignatureAlgorithm: null.StringFrom(d.SignatureAlgorithm),
		Headers:            d.Headers,
		RawBody:            d.RawBody,
		RawBodySha256:      d.RawBodySha256,
		ParsedBody:         d.ParsedBody,
		ProcessingStatus:   d.ProcessingStatus,
		RetryCount:         d.RetryCount,
		NextRetryAt:        null.TimeFrom(d.NextRetryAt),
		LockedUntil:        null.TimeFrom(d.LockedUntil),
		ReceivedAt:         d.ReceivedAt,
		ProcessedAt:        null.TimeFrom(d.ProcessedAt),
		ErrorCode:          null.StringFrom(d.ErrorCode),
		ErrorMessage:       null.StringFrom(d.ErrorMessage),
		Metadata:           d.Metadata,
	}
}

type ProviderWebhookEventsBulkUpdateRequest struct {
	Id                 uuid.UUID                     `json:"id"`
	WebhookEndpointId  uuid.UUID                     `json:"webhookEndpointId"`
	EndpointKey        string                        `json:"endpointKey"`
	ProviderAccountId  uuid.UUID                     `json:"providerAccountId"`
	ProviderCode       string                        `json:"providerCode"`
	EventId            string                        `json:"eventId"`
	EventType          string                        `json:"eventType"`
	ProviderReference  string                        `json:"providerReference"`
	ProviderStatus     string                        `json:"providerStatus"`
	SignatureValid     bool                          `json:"signatureValid"`
	SignatureAlgorithm string                        `json:"signatureAlgorithm"`
	Headers            json.RawMessage               `json:"headers"`
	RawBody            []byte                        `json:"rawBody"`
	RawBodySha256      string                        `json:"rawBodySha256"`
	ParsedBody         json.RawMessage               `json:"parsedBody"`
	ProcessingStatus   model.WebhookProcessingStatus `json:"processingStatus" example:"received" enums:"received,processing,processed,failed"`
	RetryCount         int                           `json:"retryCount"`
	NextRetryAt        time.Time                     `json:"nextRetryAt"`
	LockedUntil        time.Time                     `json:"lockedUntil"`
	ReceivedAt         time.Time                     `json:"receivedAt"`
	ProcessedAt        time.Time                     `json:"processedAt"`
	ErrorCode          string                        `json:"errorCode"`
	ErrorMessage       string                        `json:"errorMessage"`
	Metadata           json.RawMessage               `json:"metadata"`
}

func (d ProviderWebhookEventsBulkUpdateRequest) PrimaryID() ProviderWebhookEventsPrimaryID {
	return ProviderWebhookEventsPrimaryID{
		Id: d.Id,
	}
}

type ProviderWebhookEventsListBulkUpdateRequest []*ProviderWebhookEventsBulkUpdateRequest

func (d ProviderWebhookEventsListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerWebhookEvents := range d {
		err = validator.Struct(providerWebhookEvents)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ProviderWebhookEventsBulkUpdateRequest) ToModel() model.ProviderWebhookEvents {
	return model.ProviderWebhookEvents{
		Id:                 d.Id,
		WebhookEndpointId:  nuuid.From(d.WebhookEndpointId),
		EndpointKey:        null.StringFrom(d.EndpointKey),
		ProviderAccountId:  d.ProviderAccountId,
		ProviderCode:       d.ProviderCode,
		EventId:            null.StringFrom(d.EventId),
		EventType:          null.StringFrom(d.EventType),
		ProviderReference:  null.StringFrom(d.ProviderReference),
		ProviderStatus:     null.StringFrom(d.ProviderStatus),
		SignatureValid:     d.SignatureValid,
		SignatureAlgorithm: null.StringFrom(d.SignatureAlgorithm),
		Headers:            d.Headers,
		RawBody:            d.RawBody,
		RawBodySha256:      d.RawBodySha256,
		ParsedBody:         d.ParsedBody,
		ProcessingStatus:   d.ProcessingStatus,
		RetryCount:         d.RetryCount,
		NextRetryAt:        null.TimeFrom(d.NextRetryAt),
		LockedUntil:        null.TimeFrom(d.LockedUntil),
		ReceivedAt:         d.ReceivedAt,
		ProcessedAt:        null.TimeFrom(d.ProcessedAt),
		ErrorCode:          null.StringFrom(d.ErrorCode),
		ErrorMessage:       null.StringFrom(d.ErrorMessage),
		Metadata:           d.Metadata,
	}
}

type ProviderWebhookEventsResponse struct {
	Id                 uuid.UUID                     `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	WebhookEndpointId  uuid.UUID                     `json:"webhookEndpointId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	EndpointKey        string                        `json:"endpointKey"`
	ProviderAccountId  uuid.UUID                     `json:"providerAccountId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ProviderCode       string                        `json:"providerCode" validate:"required"`
	EventId            string                        `json:"eventId"`
	EventType          string                        `json:"eventType"`
	ProviderReference  string                        `json:"providerReference"`
	ProviderStatus     string                        `json:"providerStatus"`
	SignatureValid     bool                          `json:"signatureValid" example:"true"`
	SignatureAlgorithm string                        `json:"signatureAlgorithm"`
	Headers            json.RawMessage               `json:"headers" swaggertype:"object"`
	RawBody            []byte                        `json:"rawBody" validate:"required"`
	RawBodySha256      string                        `json:"rawBodySha256" validate:"required"`
	ParsedBody         json.RawMessage               `json:"parsedBody" validate:"required" swaggertype:"object"`
	ProcessingStatus   model.WebhookProcessingStatus `json:"processingStatus" validate:"oneof=received processing processed failed" enums:"received,processing,processed,failed"`
	RetryCount         int                           `json:"retryCount" example:"1"`
	NextRetryAt        time.Time                     `json:"nextRetryAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	LockedUntil        time.Time                     `json:"lockedUntil" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ReceivedAt         time.Time                     `json:"receivedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ProcessedAt        time.Time                     `json:"processedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	ErrorCode          string                        `json:"errorCode"`
	ErrorMessage       string                        `json:"errorMessage"`
	Metadata           json.RawMessage               `json:"metadata" swaggertype:"object"`
}

func NewProviderWebhookEventsResponse(providerWebhookEvents model.ProviderWebhookEvents) ProviderWebhookEventsResponse {
	return ProviderWebhookEventsResponse{
		Id:                 providerWebhookEvents.Id,
		WebhookEndpointId:  providerWebhookEvents.WebhookEndpointId.UUID,
		EndpointKey:        providerWebhookEvents.EndpointKey.String,
		ProviderAccountId:  providerWebhookEvents.ProviderAccountId,
		ProviderCode:       providerWebhookEvents.ProviderCode,
		EventId:            providerWebhookEvents.EventId.String,
		EventType:          providerWebhookEvents.EventType.String,
		ProviderReference:  providerWebhookEvents.ProviderReference.String,
		ProviderStatus:     providerWebhookEvents.ProviderStatus.String,
		SignatureValid:     providerWebhookEvents.SignatureValid,
		SignatureAlgorithm: providerWebhookEvents.SignatureAlgorithm.String,
		Headers:            providerWebhookEvents.Headers,
		RawBody:            providerWebhookEvents.RawBody,
		RawBodySha256:      providerWebhookEvents.RawBodySha256,
		ParsedBody:         providerWebhookEvents.ParsedBody,
		ProcessingStatus:   model.WebhookProcessingStatus(providerWebhookEvents.ProcessingStatus),
		RetryCount:         providerWebhookEvents.RetryCount,
		NextRetryAt:        providerWebhookEvents.NextRetryAt.Time,
		LockedUntil:        providerWebhookEvents.LockedUntil.Time,
		ReceivedAt:         providerWebhookEvents.ReceivedAt,
		ProcessedAt:        providerWebhookEvents.ProcessedAt.Time,
		ErrorCode:          providerWebhookEvents.ErrorCode.String,
		ErrorMessage:       providerWebhookEvents.ErrorMessage.String,
		Metadata:           providerWebhookEvents.Metadata,
	}
}

type ProviderWebhookEventsListResponse []*ProviderWebhookEventsResponse

func NewProviderWebhookEventsListResponse(providerWebhookEventsList model.ProviderWebhookEventsList) ProviderWebhookEventsListResponse {
	dtoProviderWebhookEventsListResponse := ProviderWebhookEventsListResponse{}
	for _, providerWebhookEvents := range providerWebhookEventsList {
		dtoProviderWebhookEventsResponse := NewProviderWebhookEventsResponse(*providerWebhookEvents)
		dtoProviderWebhookEventsListResponse = append(dtoProviderWebhookEventsListResponse, &dtoProviderWebhookEventsResponse)
	}
	return dtoProviderWebhookEventsListResponse
}

type ProviderWebhookEventsPrimaryIDList []ProviderWebhookEventsPrimaryID

func (d ProviderWebhookEventsPrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, providerWebhookEvents := range d {
		err = validator.Struct(providerWebhookEvents)
		if err != nil {
			return
		}
	}
	return nil
}

type ProviderWebhookEventsPrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ProviderWebhookEventsPrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ProviderWebhookEventsPrimaryID) ToModel() model.ProviderWebhookEventsPrimaryID {
	return model.ProviderWebhookEventsPrimaryID{
		Id: d.Id,
	}
}
