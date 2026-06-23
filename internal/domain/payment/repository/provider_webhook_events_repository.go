package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsProviderWebhookEvents(providerWebhookEventsList []model.ProviderWebhookEvents, fieldsInsert ...ProviderWebhookEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewProviderWebhookEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, providerWebhookEvents := range providerWebhookEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, providerWebhookEvents.Id)
			case selectField.WebhookEndpointId():
				args = append(args, providerWebhookEvents.WebhookEndpointId)
			case selectField.EndpointKey():
				args = append(args, providerWebhookEvents.EndpointKey)
			case selectField.ProviderAccountId():
				args = append(args, providerWebhookEvents.ProviderAccountId)
			case selectField.ProviderCode():
				args = append(args, providerWebhookEvents.ProviderCode)
			case selectField.EventId():
				args = append(args, providerWebhookEvents.EventId)
			case selectField.EventType():
				args = append(args, providerWebhookEvents.EventType)
			case selectField.ProviderReference():
				args = append(args, providerWebhookEvents.ProviderReference)
			case selectField.ProviderStatus():
				args = append(args, providerWebhookEvents.ProviderStatus)
			case selectField.SignatureValid():
				args = append(args, providerWebhookEvents.SignatureValid)
			case selectField.SignatureAlgorithm():
				args = append(args, providerWebhookEvents.SignatureAlgorithm)
			case selectField.Headers():
				args = append(args, providerWebhookEvents.Headers)
			case selectField.RawBody():
				args = append(args, providerWebhookEvents.RawBody)
			case selectField.RawBodySha256():
				args = append(args, providerWebhookEvents.RawBodySha256)
			case selectField.ParsedBody():
				args = append(args, providerWebhookEvents.ParsedBody)
			case selectField.ProcessingStatus():
				args = append(args, providerWebhookEvents.ProcessingStatus)
			case selectField.RetryCount():
				args = append(args, providerWebhookEvents.RetryCount)
			case selectField.NextRetryAt():
				args = append(args, providerWebhookEvents.NextRetryAt)
			case selectField.LockedUntil():
				args = append(args, providerWebhookEvents.LockedUntil)
			case selectField.ReceivedAt():
				args = append(args, providerWebhookEvents.ReceivedAt)
			case selectField.ProcessedAt():
				args = append(args, providerWebhookEvents.ProcessedAt)
			case selectField.ErrorCode():
				args = append(args, providerWebhookEvents.ErrorCode)
			case selectField.ErrorMessage():
				args = append(args, providerWebhookEvents.ErrorMessage)
			case selectField.Metadata():
				args = append(args, providerWebhookEvents.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, providerWebhookEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, providerWebhookEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, providerWebhookEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, providerWebhookEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, providerWebhookEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, providerWebhookEvents.MetaDeletedBy)

			}
		}
		for i := 1; i <= len(fields); i++ {
			values = append(values, fmt.Sprintf("$%d", index+i))
		}
		index += len(fields)

		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}

	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))

	return
}

func composeProviderWebhookEventsCompositePrimaryKeyWhere(primaryIDs []model.ProviderWebhookEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"provider_webhook_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultProviderWebhookEventsSelectFields() string {
	fields := NewProviderWebhookEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeProviderWebhookEventsSelectFields(selectFields ...ProviderWebhookEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ProviderWebhookEventsField string
type ProviderWebhookEventsFieldList []ProviderWebhookEventsField

type ProviderWebhookEventsSelectFields struct {
}

func (ss ProviderWebhookEventsSelectFields) Id() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("id")
}

func (ss ProviderWebhookEventsSelectFields) WebhookEndpointId() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("webhook_endpoint_id")
}

func (ss ProviderWebhookEventsSelectFields) EndpointKey() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("endpoint_key")
}

func (ss ProviderWebhookEventsSelectFields) ProviderAccountId() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("provider_account_id")
}

func (ss ProviderWebhookEventsSelectFields) ProviderCode() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("provider_code")
}

func (ss ProviderWebhookEventsSelectFields) EventId() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("event_id")
}

func (ss ProviderWebhookEventsSelectFields) EventType() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("event_type")
}

func (ss ProviderWebhookEventsSelectFields) ProviderReference() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("provider_reference")
}

func (ss ProviderWebhookEventsSelectFields) ProviderStatus() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("provider_status")
}

func (ss ProviderWebhookEventsSelectFields) SignatureValid() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("signature_valid")
}

func (ss ProviderWebhookEventsSelectFields) SignatureAlgorithm() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("signature_algorithm")
}

func (ss ProviderWebhookEventsSelectFields) Headers() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("headers")
}

func (ss ProviderWebhookEventsSelectFields) RawBody() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("raw_body")
}

func (ss ProviderWebhookEventsSelectFields) RawBodySha256() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("raw_body_sha256")
}

func (ss ProviderWebhookEventsSelectFields) ParsedBody() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("parsed_body")
}

func (ss ProviderWebhookEventsSelectFields) ProcessingStatus() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("processing_status")
}

func (ss ProviderWebhookEventsSelectFields) RetryCount() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("retry_count")
}

func (ss ProviderWebhookEventsSelectFields) NextRetryAt() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("next_retry_at")
}

func (ss ProviderWebhookEventsSelectFields) LockedUntil() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("locked_until")
}

func (ss ProviderWebhookEventsSelectFields) ReceivedAt() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("received_at")
}

func (ss ProviderWebhookEventsSelectFields) ProcessedAt() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("processed_at")
}

func (ss ProviderWebhookEventsSelectFields) ErrorCode() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("error_code")
}

func (ss ProviderWebhookEventsSelectFields) ErrorMessage() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("error_message")
}

func (ss ProviderWebhookEventsSelectFields) Metadata() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("metadata")
}

func (ss ProviderWebhookEventsSelectFields) MetaCreatedAt() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("meta_created_at")
}

func (ss ProviderWebhookEventsSelectFields) MetaCreatedBy() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("meta_created_by")
}

func (ss ProviderWebhookEventsSelectFields) MetaUpdatedAt() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("meta_updated_at")
}

func (ss ProviderWebhookEventsSelectFields) MetaUpdatedBy() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("meta_updated_by")
}

func (ss ProviderWebhookEventsSelectFields) MetaDeletedAt() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("meta_deleted_at")
}

func (ss ProviderWebhookEventsSelectFields) MetaDeletedBy() ProviderWebhookEventsField {
	return ProviderWebhookEventsField("meta_deleted_by")
}

func (ss ProviderWebhookEventsSelectFields) All() ProviderWebhookEventsFieldList {
	return []ProviderWebhookEventsField{
		ss.Id(),
		ss.WebhookEndpointId(),
		ss.EndpointKey(),
		ss.ProviderAccountId(),
		ss.ProviderCode(),
		ss.EventId(),
		ss.EventType(),
		ss.ProviderReference(),
		ss.ProviderStatus(),
		ss.SignatureValid(),
		ss.SignatureAlgorithm(),
		ss.Headers(),
		ss.RawBody(),
		ss.RawBodySha256(),
		ss.ParsedBody(),
		ss.ProcessingStatus(),
		ss.RetryCount(),
		ss.NextRetryAt(),
		ss.LockedUntil(),
		ss.ReceivedAt(),
		ss.ProcessedAt(),
		ss.ErrorCode(),
		ss.ErrorMessage(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewProviderWebhookEventsSelectFields() ProviderWebhookEventsSelectFields {
	return ProviderWebhookEventsSelectFields{}
}

type ProviderWebhookEventsUpdateFieldOption struct {
	useIncrement bool
}
type ProviderWebhookEventsUpdateField struct {
	providerWebhookEventsField ProviderWebhookEventsField
	opt                        ProviderWebhookEventsUpdateFieldOption
	value                      interface{}
}
type ProviderWebhookEventsUpdateFieldList []ProviderWebhookEventsUpdateField

func defaultProviderWebhookEventsUpdateFieldOption() ProviderWebhookEventsUpdateFieldOption {
	return ProviderWebhookEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementProviderWebhookEventsOption(useIncrement bool) func(*ProviderWebhookEventsUpdateFieldOption) {
	return func(pcufo *ProviderWebhookEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewProviderWebhookEventsUpdateField(field ProviderWebhookEventsField, val interface{}, opts ...func(*ProviderWebhookEventsUpdateFieldOption)) ProviderWebhookEventsUpdateField {
	defaultOpt := defaultProviderWebhookEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ProviderWebhookEventsUpdateField{
		providerWebhookEventsField: field,
		value:                      val,
		opt:                        defaultOpt,
	}
}
func defaultProviderWebhookEventsUpdateFields(providerWebhookEvents model.ProviderWebhookEvents) (providerWebhookEventsUpdateFieldList ProviderWebhookEventsUpdateFieldList) {
	selectFields := NewProviderWebhookEventsSelectFields()
	providerWebhookEventsUpdateFieldList = append(providerWebhookEventsUpdateFieldList,
		NewProviderWebhookEventsUpdateField(selectFields.Id(), providerWebhookEvents.Id),
		NewProviderWebhookEventsUpdateField(selectFields.WebhookEndpointId(), providerWebhookEvents.WebhookEndpointId),
		NewProviderWebhookEventsUpdateField(selectFields.EndpointKey(), providerWebhookEvents.EndpointKey),
		NewProviderWebhookEventsUpdateField(selectFields.ProviderAccountId(), providerWebhookEvents.ProviderAccountId),
		NewProviderWebhookEventsUpdateField(selectFields.ProviderCode(), providerWebhookEvents.ProviderCode),
		NewProviderWebhookEventsUpdateField(selectFields.EventId(), providerWebhookEvents.EventId),
		NewProviderWebhookEventsUpdateField(selectFields.EventType(), providerWebhookEvents.EventType),
		NewProviderWebhookEventsUpdateField(selectFields.ProviderReference(), providerWebhookEvents.ProviderReference),
		NewProviderWebhookEventsUpdateField(selectFields.ProviderStatus(), providerWebhookEvents.ProviderStatus),
		NewProviderWebhookEventsUpdateField(selectFields.SignatureValid(), providerWebhookEvents.SignatureValid),
		NewProviderWebhookEventsUpdateField(selectFields.SignatureAlgorithm(), providerWebhookEvents.SignatureAlgorithm),
		NewProviderWebhookEventsUpdateField(selectFields.Headers(), providerWebhookEvents.Headers),
		NewProviderWebhookEventsUpdateField(selectFields.RawBody(), providerWebhookEvents.RawBody),
		NewProviderWebhookEventsUpdateField(selectFields.RawBodySha256(), providerWebhookEvents.RawBodySha256),
		NewProviderWebhookEventsUpdateField(selectFields.ParsedBody(), providerWebhookEvents.ParsedBody),
		NewProviderWebhookEventsUpdateField(selectFields.ProcessingStatus(), providerWebhookEvents.ProcessingStatus),
		NewProviderWebhookEventsUpdateField(selectFields.RetryCount(), providerWebhookEvents.RetryCount),
		NewProviderWebhookEventsUpdateField(selectFields.NextRetryAt(), providerWebhookEvents.NextRetryAt),
		NewProviderWebhookEventsUpdateField(selectFields.LockedUntil(), providerWebhookEvents.LockedUntil),
		NewProviderWebhookEventsUpdateField(selectFields.ReceivedAt(), providerWebhookEvents.ReceivedAt),
		NewProviderWebhookEventsUpdateField(selectFields.ProcessedAt(), providerWebhookEvents.ProcessedAt),
		NewProviderWebhookEventsUpdateField(selectFields.ErrorCode(), providerWebhookEvents.ErrorCode),
		NewProviderWebhookEventsUpdateField(selectFields.ErrorMessage(), providerWebhookEvents.ErrorMessage),
		NewProviderWebhookEventsUpdateField(selectFields.Metadata(), providerWebhookEvents.Metadata),
		NewProviderWebhookEventsUpdateField(selectFields.MetaCreatedAt(), providerWebhookEvents.MetaCreatedAt),
		NewProviderWebhookEventsUpdateField(selectFields.MetaCreatedBy(), providerWebhookEvents.MetaCreatedBy),
		NewProviderWebhookEventsUpdateField(selectFields.MetaUpdatedAt(), providerWebhookEvents.MetaUpdatedAt),
		NewProviderWebhookEventsUpdateField(selectFields.MetaUpdatedBy(), providerWebhookEvents.MetaUpdatedBy),
		NewProviderWebhookEventsUpdateField(selectFields.MetaDeletedAt(), providerWebhookEvents.MetaDeletedAt),
		NewProviderWebhookEventsUpdateField(selectFields.MetaDeletedBy(), providerWebhookEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsProviderWebhookEventsCommand(providerWebhookEventsUpdateFieldList ProviderWebhookEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range providerWebhookEventsUpdateFieldList {
		field := string(updateField.providerWebhookEventsField)
		valueParam := fmt.Sprintf("\"%s\" = $%d", field, index)
		if updateField.opt.useIncrement {
			valueParam = fmt.Sprintf("\"%s\" = %s + $%d", field, field, index)
		}
		updatedFieldsQuery = append(updatedFieldsQuery, valueParam)
		index++
		args = append(args, updateField.value)
	}

	return updatedFieldsQuery, args
}

func (repo *RepositoryImpl) BulkCreateProviderWebhookEvents(ctx context.Context, providerWebhookEventsList []*model.ProviderWebhookEvents, fieldsInsert ...ProviderWebhookEventsField) (err error) {
	var (
		fieldsStr                      string
		valueListStr                   []string
		argsList                       []interface{}
		primaryIds                     []model.ProviderWebhookEventsPrimaryID
		providerWebhookEventsValueList []model.ProviderWebhookEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewProviderWebhookEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, providerWebhookEvents := range providerWebhookEventsList {

		primaryIds = append(primaryIds, providerWebhookEvents.ToProviderWebhookEventsPrimaryID())

		providerWebhookEventsValueList = append(providerWebhookEventsValueList, *providerWebhookEvents)
	}

	_, notFoundIds, err := repo.IsExistProviderWebhookEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderWebhookEvents] failed checking providerWebhookEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ProviderWebhookEventsPrimaryID{}
		mapNotFoundIds := map[model.ProviderWebhookEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "providerWebhookEvents", fmt.Sprintf("providerWebhookEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsProviderWebhookEvents(providerWebhookEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(providerWebhookEventsQueries.insertProviderWebhookEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderWebhookEvents] failed exec create providerWebhookEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteProviderWebhookEventsByIDs(ctx context.Context, primaryIDs []model.ProviderWebhookEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistProviderWebhookEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderWebhookEventsByIDs] failed checking providerWebhookEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_webhook_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := providerWebhookEventsQueries.deleteProviderWebhookEvents + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderWebhookEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderWebhookEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistProviderWebhookEventsByIDs(ctx context.Context, ids []model.ProviderWebhookEventsPrimaryID) (exists bool, notFoundIds []model.ProviderWebhookEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_webhook_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(providerWebhookEventsQueries.selectProviderWebhookEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderWebhookEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ProviderWebhookEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderWebhookEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ProviderWebhookEventsPrimaryID]bool{}
	for _, id := range resIds {
		mapReqIds[id] = true
	}
	for _, reqId := range ids {
		if exist := mapReqIds[reqId]; !exist {
			notFoundIds = append(notFoundIds, reqId)
		}
	}
	return len(notFoundIds) == 0, notFoundIds, nil
}

// BulkUpdateProviderWebhookEvents is used to bulk update providerWebhookEvents, by default it will update all field
// if want to update specific field, then fill providerWebhookEventssMapUpdateFieldsRequest else please fill providerWebhookEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateProviderWebhookEvents(ctx context.Context, providerWebhookEventssMap map[model.ProviderWebhookEventsPrimaryID]*model.ProviderWebhookEvents, providerWebhookEventssMapUpdateFieldsRequest map[model.ProviderWebhookEventsPrimaryID]ProviderWebhookEventsUpdateFieldList) (err error) {
	if len(providerWebhookEventssMap) == 0 && len(providerWebhookEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		providerWebhookEventssMapUpdateField map[model.ProviderWebhookEventsPrimaryID]ProviderWebhookEventsUpdateFieldList = map[model.ProviderWebhookEventsPrimaryID]ProviderWebhookEventsUpdateFieldList{}
		asTableValues                        string                                                                        = "myvalues"
	)

	if len(providerWebhookEventssMap) > 0 {
		for id, providerWebhookEvents := range providerWebhookEventssMap {
			if providerWebhookEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdateProviderWebhookEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			providerWebhookEventssMapUpdateField[id] = defaultProviderWebhookEventsUpdateFields(*providerWebhookEvents)
		}
	} else {
		providerWebhookEventssMapUpdateField = providerWebhookEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateProviderWebhookEventsQuery(providerWebhookEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistProviderWebhookEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderWebhookEvents] failed checking providerWebhookEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeProviderWebhookEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"provider_webhook_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderWebhookEvents] failed exec query")
	}
	return
}

type ProviderWebhookEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewProviderWebhookEventsFieldParameter(param string, args ...interface{}) ProviderWebhookEventsFieldParameter {
	return ProviderWebhookEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateProviderWebhookEventsQuery(mapProviderWebhookEventss map[model.ProviderWebhookEventsPrimaryID]ProviderWebhookEventsUpdateFieldList, asTableValues string) (primaryIDs []model.ProviderWebhookEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ProviderWebhookEventsPrimaryID]map[string]interface{}{}
	providerWebhookEventsSelectFields := NewProviderWebhookEventsSelectFields()
	for id, updateFields := range mapProviderWebhookEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.providerWebhookEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapProviderWebhookEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetProviderWebhookEventsFieldType(updateField.providerWebhookEventsField)))
			args = append(args, fields[string(updateField.providerWebhookEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.providerWebhookEventsField))
		if updateField.providerWebhookEventsField == providerWebhookEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.providerWebhookEventsField, asTableValues, updateField.providerWebhookEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.providerWebhookEventsField,
				"\"provider_webhook_events\"", updateField.providerWebhookEventsField,
				asTableValues, updateField.providerWebhookEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeProviderWebhookEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ProviderWebhookEventsPrimaryID, asTableValue string) (whereQry string) {
	providerWebhookEventsSelectFields := NewProviderWebhookEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"provider_webhook_events\".\"id\" = %s.\"id\"::"+GetProviderWebhookEventsFieldType(providerWebhookEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetProviderWebhookEventsFieldType(providerWebhookEventsField ProviderWebhookEventsField) string {
	selectProviderWebhookEventsFields := NewProviderWebhookEventsSelectFields()
	switch providerWebhookEventsField {

	case selectProviderWebhookEventsFields.Id():
		return "uuid"

	case selectProviderWebhookEventsFields.WebhookEndpointId():
		return "uuid"

	case selectProviderWebhookEventsFields.EndpointKey():
		return "text"

	case selectProviderWebhookEventsFields.ProviderAccountId():
		return "uuid"

	case selectProviderWebhookEventsFields.ProviderCode():
		return "text"

	case selectProviderWebhookEventsFields.EventId():
		return "text"

	case selectProviderWebhookEventsFields.EventType():
		return "text"

	case selectProviderWebhookEventsFields.ProviderReference():
		return "text"

	case selectProviderWebhookEventsFields.ProviderStatus():
		return "text"

	case selectProviderWebhookEventsFields.SignatureValid():
		return "bool"

	case selectProviderWebhookEventsFields.SignatureAlgorithm():
		return "text"

	case selectProviderWebhookEventsFields.Headers():
		return "jsonb"

	case selectProviderWebhookEventsFields.RawBody():
		return "bytea"

	case selectProviderWebhookEventsFields.RawBodySha256():
		return "text"

	case selectProviderWebhookEventsFields.ParsedBody():
		return "jsonb"

	case selectProviderWebhookEventsFields.ProcessingStatus():
		return "webhook_processing_status_enum"

	case selectProviderWebhookEventsFields.RetryCount():
		return "int4"

	case selectProviderWebhookEventsFields.NextRetryAt():
		return "timestamptz"

	case selectProviderWebhookEventsFields.LockedUntil():
		return "timestamptz"

	case selectProviderWebhookEventsFields.ReceivedAt():
		return "timestamptz"

	case selectProviderWebhookEventsFields.ProcessedAt():
		return "timestamptz"

	case selectProviderWebhookEventsFields.ErrorCode():
		return "text"

	case selectProviderWebhookEventsFields.ErrorMessage():
		return "text"

	case selectProviderWebhookEventsFields.Metadata():
		return "jsonb"

	case selectProviderWebhookEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectProviderWebhookEventsFields.MetaCreatedBy():
		return "uuid"

	case selectProviderWebhookEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectProviderWebhookEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectProviderWebhookEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectProviderWebhookEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateProviderWebhookEvents(ctx context.Context, providerWebhookEvents *model.ProviderWebhookEvents, fieldsInsert ...ProviderWebhookEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewProviderWebhookEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ProviderWebhookEventsPrimaryID{
		Id: providerWebhookEvents.Id,
	}
	exists, err := repo.IsExistProviderWebhookEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderWebhookEvents] failed checking providerWebhookEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "providerWebhookEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsProviderWebhookEvents([]model.ProviderWebhookEvents{*providerWebhookEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(providerWebhookEventsQueries.insertProviderWebhookEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderWebhookEvents] failed exec create providerWebhookEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteProviderWebhookEventsByID(ctx context.Context, primaryID model.ProviderWebhookEventsPrimaryID) (err error) {
	exists, err := repo.IsExistProviderWebhookEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderWebhookEventsByID] failed checking providerWebhookEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeProviderWebhookEventsCompositePrimaryKeyWhere([]model.ProviderWebhookEventsPrimaryID{primaryID})
	commandQuery := providerWebhookEventsQueries.deleteProviderWebhookEvents + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderWebhookEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderWebhookEventsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderWebhookEventsFilterResult, err error) {
	query, args, err := composeProviderWebhookEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderWebhookEventsByFilter] failed compose providerWebhookEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderWebhookEventsByFilter] failed get providerWebhookEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeProviderWebhookEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ProviderWebhookEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeProviderWebhookEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeProviderWebhookEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeProviderWebhookEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 30 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewProviderWebhookEventsFilterFieldSpecFromStr(sourceField)
		if !found || !spec.Selectable || spec.Relation != "" {
			return failure.BadRequestFromString(fmt.Sprintf("field %s is not selectable", sourceField))
		}
		if _, selected := selectedColumns[spec.Column]; selected {
			return nil
		}
		selectColumns = append(selectColumns, fmt.Sprintf("base.\"%s\"", spec.Column))
		selectedColumns[spec.Column] = struct{}{}
		return nil
	}

	if len(filter.SelectFields) == 0 {
		selectColumns = make([]string, 0, 30+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["webhook_endpoint_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"webhook_endpoint_id\"")
			selectedColumns["webhook_endpoint_id"] = struct{}{}
		}
		if _, selected := selectedColumns["endpoint_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"endpoint_key\"")
			selectedColumns["endpoint_key"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_code\"")
			selectedColumns["provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["event_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_id\"")
			selectedColumns["event_id"] = struct{}{}
		}
		if _, selected := selectedColumns["event_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_type\"")
			selectedColumns["event_type"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_reference"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_reference\"")
			selectedColumns["provider_reference"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_status\"")
			selectedColumns["provider_status"] = struct{}{}
		}
		if _, selected := selectedColumns["signature_valid"]; !selected {
			selectColumns = append(selectColumns, "base.\"signature_valid\"")
			selectedColumns["signature_valid"] = struct{}{}
		}
		if _, selected := selectedColumns["signature_algorithm"]; !selected {
			selectColumns = append(selectColumns, "base.\"signature_algorithm\"")
			selectedColumns["signature_algorithm"] = struct{}{}
		}
		if _, selected := selectedColumns["headers"]; !selected {
			selectColumns = append(selectColumns, "base.\"headers\"")
			selectedColumns["headers"] = struct{}{}
		}
		if _, selected := selectedColumns["raw_body"]; !selected {
			selectColumns = append(selectColumns, "base.\"raw_body\"")
			selectedColumns["raw_body"] = struct{}{}
		}
		if _, selected := selectedColumns["raw_body_sha256"]; !selected {
			selectColumns = append(selectColumns, "base.\"raw_body_sha256\"")
			selectedColumns["raw_body_sha256"] = struct{}{}
		}
		if _, selected := selectedColumns["parsed_body"]; !selected {
			selectColumns = append(selectColumns, "base.\"parsed_body\"")
			selectedColumns["parsed_body"] = struct{}{}
		}
		if _, selected := selectedColumns["processing_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"processing_status\"")
			selectedColumns["processing_status"] = struct{}{}
		}
		if _, selected := selectedColumns["retry_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"retry_count\"")
			selectedColumns["retry_count"] = struct{}{}
		}
		if _, selected := selectedColumns["next_retry_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"next_retry_at\"")
			selectedColumns["next_retry_at"] = struct{}{}
		}
		if _, selected := selectedColumns["locked_until"]; !selected {
			selectColumns = append(selectColumns, "base.\"locked_until\"")
			selectedColumns["locked_until"] = struct{}{}
		}
		if _, selected := selectedColumns["received_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"received_at\"")
			selectedColumns["received_at"] = struct{}{}
		}
		if _, selected := selectedColumns["processed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"processed_at\"")
			selectedColumns["processed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["error_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"error_code\"")
			selectedColumns["error_code"] = struct{}{}
		}
		if _, selected := selectedColumns["error_message"]; !selected {
			selectColumns = append(selectColumns, "base.\"error_message\"")
			selectedColumns["error_message"] = struct{}{}
		}
		if _, selected := selectedColumns["metadata"]; !selected {
			selectColumns = append(selectColumns, "base.\"metadata\"")
			selectedColumns["metadata"] = struct{}{}
		}
		if _, selected := selectedColumns["meta_created_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"meta_created_at\"")
			selectedColumns["meta_created_at"] = struct{}{}
		}
		if _, selected := selectedColumns["meta_created_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"meta_created_by\"")
			selectedColumns["meta_created_by"] = struct{}{}
		}
		if _, selected := selectedColumns["meta_updated_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"meta_updated_at\"")
			selectedColumns["meta_updated_at"] = struct{}{}
		}
		if _, selected := selectedColumns["meta_updated_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"meta_updated_by\"")
			selectedColumns["meta_updated_by"] = struct{}{}
		}
		if _, selected := selectedColumns["meta_deleted_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"meta_deleted_at\"")
			selectedColumns["meta_deleted_at"] = struct{}{}
		}
		if _, selected := selectedColumns["meta_deleted_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"meta_deleted_by\"")
			selectedColumns["meta_deleted_by"] = struct{}{}
		}

	} else {
		selectColumns = make([]string, 0, len(filter.SelectFields)+1)
		for _, field := range filter.SelectFields {
			if err = addColumn(field); err != nil {
				return
			}
		}
	}

	if _, selected := selectedColumns["id"]; isCursorMode && !selected {
		selectColumns = append(selectColumns, "base.\"id\"")
		selectedColumns["id"] = struct{}{}
	}

	return
}

type providerWebhookEventsFilterPlaceholder struct {
	index int
}

func (p *providerWebhookEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeProviderWebhookEventsFilterPredicate(filterField model.FilterField, placeholders *providerWebhookEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewProviderWebhookEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeProviderWebhookEventsFilterSQLExpr(spec)
	if err != nil {
		return "", err
	}
	if spec.Relation != "" {
		requiredJoins[spec.Relation] = true
	}
	switch filterField.Operator {
	case model.OperatorEqual:
		placeholder := placeholders.Next()
		*args = append(*args, filterField.Value)
		return fmt.Sprintf("%s = %s", sqlExpr, placeholder), nil
	case model.OperatorRange:
		valueArray, ok := filterField.Value.([]interface{})
		if !ok || len(valueArray) != 2 {
			return "", failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
		start := placeholders.Next()
		end := placeholders.Next()
		*args = append(*args, valueArray...)
		return fmt.Sprintf("%s BETWEEN %s AND %s", sqlExpr, start, end), nil
	case model.OperatorIn, model.OperatorNotIn:
		valueArray, ok := filterField.Value.([]interface{})
		if !ok || len(valueArray) == 0 {
			return "", failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
		placeholder := []string{}
		for range valueArray {
			placeholder = append(placeholder, placeholders.Next())
		}
		operator := "IN"
		if filterField.Operator == model.OperatorNotIn {
			operator = "NOT IN"
		}
		*args = append(*args, valueArray...)
		return fmt.Sprintf("%s %s (%s)", sqlExpr, operator, strings.Join(placeholder, ",")), nil
	case model.OperatorIsNull:
		value, ok := filterField.Value.(bool)
		if !ok {
			return "", failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
		if value {
			return fmt.Sprintf("%s IS NULL", sqlExpr), nil
		}
		return fmt.Sprintf("%s IS NOT NULL", sqlExpr), nil
	case model.OperatorNot:
		placeholder := placeholders.Next()
		*args = append(*args, filterField.Value)
		return fmt.Sprintf("%s != %s", sqlExpr, placeholder), nil
	case model.OperatorGT:
		placeholder := placeholders.Next()
		*args = append(*args, filterField.Value)
		return fmt.Sprintf("%s > %s", sqlExpr, placeholder), nil
	case model.OperatorGTE:
		placeholder := placeholders.Next()
		*args = append(*args, filterField.Value)
		return fmt.Sprintf("%s >= %s", sqlExpr, placeholder), nil
	case model.OperatorLT:
		placeholder := placeholders.Next()
		*args = append(*args, filterField.Value)
		return fmt.Sprintf("%s < %s", sqlExpr, placeholder), nil
	case model.OperatorLTE:
		placeholder := placeholders.Next()
		*args = append(*args, filterField.Value)
		return fmt.Sprintf("%s <= %s", sqlExpr, placeholder), nil
	case model.OperatorLike:
		value, ok := filterField.Value.(string)
		if !ok {
			return "", failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
		placeholder := placeholders.Next()
		*args = append(*args, "%"+strings.TrimSpace(value)+"%")
		return fmt.Sprintf("%s ILIKE %s", sqlExpr, placeholder), nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("operator %s is not allowed", filterField.Operator))
	}
}

func composeProviderWebhookEventsFilterGroup(group model.FilterGroup, placeholders *providerWebhookEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	logic := strings.ToUpper(strings.TrimSpace(group.Logic))
	if logic == "" {
		logic = "AND"
	}
	switch logic {
	case "AND", "OR":
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("filter logic %s is not allowed", group.Logic))
	}

	parts := []string{}
	for _, filterField := range group.FilterFields {
		predicate, err := composeProviderWebhookEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeProviderWebhookEventsFilterGroup(child, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if childQuery != "" {
			parts = append(parts, childQuery)
		}
	}
	if len(parts) == 0 {
		return "", nil
	}
	if len(parts) == 1 {
		return parts[0], nil
	}
	return "(" + strings.Join(parts, " "+logic+" ") + ")", nil
}

func composeProviderWebhookEventsFilterWhereQueries(filter model.Filter, placeholders *providerWebhookEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeProviderWebhookEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeProviderWebhookEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
		if groupErr != nil {
			err = groupErr
			return
		}
		if groupQuery != "" {
			whereQueries = append(whereQueries, groupQuery)
		}
	}
	return
}

func composeProviderWebhookEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateProviderWebhookEventsFieldNameFilter(filter)
	if err != nil {
		return
	}
	isCursorMode := filter.Pagination.IsCursorMode()
	requiredJoins := map[string]bool{}
	cursorOperator := ">"
	cursorSortOrder := model.SortAsc
	if filter.Pagination.Direction == model.CursorDirectionPrev {
		cursorOperator = "<"
		cursorSortOrder = model.SortDesc
	}
	if isCursorMode {
		if len(filter.Sorts) > 1 {
			err = failure.BadRequestFromString("cursor pagination only supports the default primary-key sort")
			return
		}
		if len(filter.Sorts) == 1 {
			sortOrder, sortErr := normalizeProviderWebhookEventsSortOrder(filter.Sorts[0].Order)
			if sortErr != nil {
				err = sortErr
				return
			}
			if filter.Sorts[0].Field != "id" || sortOrder != model.SortAsc {
				err = failure.BadRequestFromString("cursor pagination only supports the default primary-key sort")
				return
			}
		}
	}

	selectColumns, err := composeProviderWebhookEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := providerWebhookEventsFilterPlaceholder{index: 1}
	whereQueries, err := composeProviderWebhookEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
	if err != nil {
		return
	}

	if isCursorMode && filter.Pagination.Cursor != nil {
		whereQueries = append(whereQueries, fmt.Sprintf("base.\"id\" %s %s", cursorOperator, placeholders.Next()))
		args = append(args, filter.Pagination.Cursor)
	}

	sortQuery := []string{}
	if isCursorMode {
		sortQuery = append(sortQuery, fmt.Sprintf("base.\"id\" %s", cursorSortOrder))

	} else {
		for _, sort := range filter.Sorts {
			spec, found := model.NewProviderWebhookEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeProviderWebhookEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeProviderWebhookEventsSortOrder(sort.Order)
			if sortErr != nil {
				err = sortErr
				return
			}
			sortQuery = append(sortQuery, fmt.Sprintf("%s %s", sqlExpr, sortOrder))
		}
		if len(sortQuery) == 0 {
			sortQuery = append(sortQuery, "base.\"id\" ASC")
		}

	}

	query = fmt.Sprintf("SELECT %s FROM \"provider_webhook_events\" base%s", strings.Join(selectColumns, ","), composeProviderWebhookEventsFilterJoins(requiredJoins))
	if len(whereQueries) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND "))
	}
	if len(sortQuery) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
	}
	if filter.Pagination.PageSize > 0 {
		limit := filter.Pagination.PageSize
		if isCursorMode {
			limit = filter.Pagination.PageSize + 1
		}
		query += fmt.Sprintf(" LIMIT %d", limit)
		if !isCursorMode && filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
	}

	return
}

func (repo *RepositoryImpl) IsExistProviderWebhookEventsByID(ctx context.Context, primaryID model.ProviderWebhookEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeProviderWebhookEventsCompositePrimaryKeyWhere([]model.ProviderWebhookEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", providerWebhookEventsQueries.selectCountProviderWebhookEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderWebhookEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderWebhookEvents(ctx context.Context, selectFields ...ProviderWebhookEventsField) (providerWebhookEventsList model.ProviderWebhookEventsList, err error) {
	var (
		defaultProviderWebhookEventsSelectFields = defaultProviderWebhookEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderWebhookEventsSelectFields = composeProviderWebhookEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(providerWebhookEventsQueries.selectProviderWebhookEvents, defaultProviderWebhookEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &providerWebhookEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderWebhookEvents] failed get providerWebhookEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderWebhookEventsByID(ctx context.Context, primaryID model.ProviderWebhookEventsPrimaryID, selectFields ...ProviderWebhookEventsField) (providerWebhookEvents model.ProviderWebhookEvents, err error) {
	var (
		defaultProviderWebhookEventsSelectFields = defaultProviderWebhookEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderWebhookEventsSelectFields = composeProviderWebhookEventsSelectFields(selectFields...)
	}
	whereQry, params := composeProviderWebhookEventsCompositePrimaryKeyWhere([]model.ProviderWebhookEventsPrimaryID{primaryID})
	query := fmt.Sprintf(providerWebhookEventsQueries.selectProviderWebhookEvents+" WHERE "+whereQry, defaultProviderWebhookEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &providerWebhookEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("providerWebhookEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveProviderWebhookEventsByID] failed get providerWebhookEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateProviderWebhookEventsByID(ctx context.Context, primaryID model.ProviderWebhookEventsPrimaryID, providerWebhookEvents *model.ProviderWebhookEvents, providerWebhookEventsUpdateFields ...ProviderWebhookEventsUpdateField) (err error) {
	exists, err := repo.IsExistProviderWebhookEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEvents] failed checking providerWebhookEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if providerWebhookEvents == nil {
		if len(providerWebhookEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateProviderWebhookEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		providerWebhookEvents = &model.ProviderWebhookEvents{}
	}
	var (
		defaultProviderWebhookEventsUpdateFields = defaultProviderWebhookEventsUpdateFields(*providerWebhookEvents)
		tempUpdateField                          ProviderWebhookEventsUpdateFieldList
		selectFields                             = NewProviderWebhookEventsSelectFields()
	)
	if len(providerWebhookEventsUpdateFields) > 0 {
		for _, updateField := range providerWebhookEventsUpdateFields {
			if updateField.providerWebhookEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultProviderWebhookEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeProviderWebhookEventsCompositePrimaryKeyWhere([]model.ProviderWebhookEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsProviderWebhookEventsCommand(defaultProviderWebhookEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(providerWebhookEventsQueries.updateProviderWebhookEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEvents] error when try to update providerWebhookEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateProviderWebhookEventsByFilter(ctx context.Context, filter model.Filter, providerWebhookEventsUpdateFields ...ProviderWebhookEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(providerWebhookEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ProviderWebhookEventsUpdateFieldList
		selectFields = NewProviderWebhookEventsSelectFields()
	)
	for _, updateField := range providerWebhookEventsUpdateFields {
		if updateField.providerWebhookEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsProviderWebhookEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := providerWebhookEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeProviderWebhookEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
	if err != nil {
		return
	}
	if len(whereQueries) == 0 {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(requiredJoins) > 0 {
		err = failure.BadRequestFromString("update by filter does not support join fields")
		return
	}

	commandQuery := fmt.Sprintf("UPDATE \"provider_webhook_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEventsByFilter] error when try to update providerWebhookEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEventsByFilter] failed get rows affected")
	}
	return
}

var (
	providerWebhookEventsQueries = struct {
		selectProviderWebhookEvents      string
		selectCountProviderWebhookEvents string
		deleteProviderWebhookEvents      string
		updateProviderWebhookEvents      string
		insertProviderWebhookEvents      string
	}{
		selectProviderWebhookEvents:      "SELECT %s FROM \"provider_webhook_events\"",
		selectCountProviderWebhookEvents: "SELECT COUNT(\"id\") FROM \"provider_webhook_events\"",
		deleteProviderWebhookEvents:      "DELETE FROM \"provider_webhook_events\"",
		updateProviderWebhookEvents:      "UPDATE \"provider_webhook_events\" SET %s ",
		insertProviderWebhookEvents:      "INSERT INTO \"provider_webhook_events\" %s VALUES %s",
	}
)

type ProviderWebhookEventsRepository interface {
	CreateProviderWebhookEvents(ctx context.Context, providerWebhookEvents *model.ProviderWebhookEvents, fieldsInsert ...ProviderWebhookEventsField) error
	BulkCreateProviderWebhookEvents(ctx context.Context, providerWebhookEventsList []*model.ProviderWebhookEvents, fieldsInsert ...ProviderWebhookEventsField) error
	ResolveProviderWebhookEvents(ctx context.Context, selectFields ...ProviderWebhookEventsField) (model.ProviderWebhookEventsList, error)
	ResolveProviderWebhookEventsByID(ctx context.Context, primaryID model.ProviderWebhookEventsPrimaryID, selectFields ...ProviderWebhookEventsField) (model.ProviderWebhookEvents, error)
	UpdateProviderWebhookEventsByID(ctx context.Context, id model.ProviderWebhookEventsPrimaryID, providerWebhookEvents *model.ProviderWebhookEvents, providerWebhookEventsUpdateFields ...ProviderWebhookEventsUpdateField) error
	UpdateProviderWebhookEventsByFilter(ctx context.Context, filter model.Filter, providerWebhookEventsUpdateFields ...ProviderWebhookEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdateProviderWebhookEvents(ctx context.Context, providerWebhookEventsListMap map[model.ProviderWebhookEventsPrimaryID]*model.ProviderWebhookEvents, ProviderWebhookEventssMapUpdateFieldsRequest map[model.ProviderWebhookEventsPrimaryID]ProviderWebhookEventsUpdateFieldList) (err error)
	DeleteProviderWebhookEventsByID(ctx context.Context, id model.ProviderWebhookEventsPrimaryID) error
	BulkDeleteProviderWebhookEventsByIDs(ctx context.Context, ids []model.ProviderWebhookEventsPrimaryID) error
	ResolveProviderWebhookEventsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderWebhookEventsFilterResult, err error)
	IsExistProviderWebhookEventsByIDs(ctx context.Context, ids []model.ProviderWebhookEventsPrimaryID) (exists bool, notFoundIds []model.ProviderWebhookEventsPrimaryID, err error)
	IsExistProviderWebhookEventsByID(ctx context.Context, id model.ProviderWebhookEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
