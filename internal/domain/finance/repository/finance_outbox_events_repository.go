package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsFinanceOutboxEvents(financeOutboxEventsList []model.FinanceOutboxEvents, fieldsInsert ...FinanceOutboxEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceOutboxEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeOutboxEvents := range financeOutboxEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeOutboxEvents.Id)
			case selectField.AggregateType():
				args = append(args, financeOutboxEvents.AggregateType)
			case selectField.AggregateId():
				args = append(args, financeOutboxEvents.AggregateId)
			case selectField.EventType():
				args = append(args, financeOutboxEvents.EventType)
			case selectField.EventVersion():
				args = append(args, financeOutboxEvents.EventVersion)
			case selectField.IdempotencyKey():
				args = append(args, financeOutboxEvents.IdempotencyKey)
			case selectField.CorrelationId():
				args = append(args, financeOutboxEvents.CorrelationId)
			case selectField.Payload():
				args = append(args, financeOutboxEvents.Payload)
			case selectField.Headers():
				args = append(args, financeOutboxEvents.Headers)
			case selectField.PublishStatus():
				args = append(args, financeOutboxEvents.PublishStatus)
			case selectField.AttemptCount():
				args = append(args, financeOutboxEvents.AttemptCount)
			case selectField.NextAttemptAt():
				args = append(args, financeOutboxEvents.NextAttemptAt)
			case selectField.PublishedAt():
				args = append(args, financeOutboxEvents.PublishedAt)
			case selectField.LastErrorCode():
				args = append(args, financeOutboxEvents.LastErrorCode)
			case selectField.LastErrorDetail():
				args = append(args, financeOutboxEvents.LastErrorDetail)
			case selectField.Metadata():
				args = append(args, financeOutboxEvents.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeOutboxEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeOutboxEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeOutboxEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeOutboxEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeOutboxEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeOutboxEvents.MetaDeletedBy)

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

func composeFinanceOutboxEventsCompositePrimaryKeyWhere(primaryIDs []model.FinanceOutboxEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_outbox_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceOutboxEventsSelectFields() string {
	fields := NewFinanceOutboxEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceOutboxEventsSelectFields(selectFields ...FinanceOutboxEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceOutboxEventsField string
type FinanceOutboxEventsFieldList []FinanceOutboxEventsField

type FinanceOutboxEventsSelectFields struct {
}

func (ss FinanceOutboxEventsSelectFields) Id() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("id")
}

func (ss FinanceOutboxEventsSelectFields) AggregateType() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("aggregate_type")
}

func (ss FinanceOutboxEventsSelectFields) AggregateId() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("aggregate_id")
}

func (ss FinanceOutboxEventsSelectFields) EventType() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("event_type")
}

func (ss FinanceOutboxEventsSelectFields) EventVersion() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("event_version")
}

func (ss FinanceOutboxEventsSelectFields) IdempotencyKey() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("idempotency_key")
}

func (ss FinanceOutboxEventsSelectFields) CorrelationId() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("correlation_id")
}

func (ss FinanceOutboxEventsSelectFields) Payload() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("payload")
}

func (ss FinanceOutboxEventsSelectFields) Headers() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("headers")
}

func (ss FinanceOutboxEventsSelectFields) PublishStatus() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("publish_status")
}

func (ss FinanceOutboxEventsSelectFields) AttemptCount() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("attempt_count")
}

func (ss FinanceOutboxEventsSelectFields) NextAttemptAt() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("next_attempt_at")
}

func (ss FinanceOutboxEventsSelectFields) PublishedAt() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("published_at")
}

func (ss FinanceOutboxEventsSelectFields) LastErrorCode() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("last_error_code")
}

func (ss FinanceOutboxEventsSelectFields) LastErrorDetail() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("last_error_detail")
}

func (ss FinanceOutboxEventsSelectFields) Metadata() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("metadata")
}

func (ss FinanceOutboxEventsSelectFields) MetaCreatedAt() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("meta_created_at")
}

func (ss FinanceOutboxEventsSelectFields) MetaCreatedBy() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("meta_created_by")
}

func (ss FinanceOutboxEventsSelectFields) MetaUpdatedAt() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("meta_updated_at")
}

func (ss FinanceOutboxEventsSelectFields) MetaUpdatedBy() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("meta_updated_by")
}

func (ss FinanceOutboxEventsSelectFields) MetaDeletedAt() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("meta_deleted_at")
}

func (ss FinanceOutboxEventsSelectFields) MetaDeletedBy() FinanceOutboxEventsField {
	return FinanceOutboxEventsField("meta_deleted_by")
}

func (ss FinanceOutboxEventsSelectFields) All() FinanceOutboxEventsFieldList {
	return []FinanceOutboxEventsField{
		ss.Id(),
		ss.AggregateType(),
		ss.AggregateId(),
		ss.EventType(),
		ss.EventVersion(),
		ss.IdempotencyKey(),
		ss.CorrelationId(),
		ss.Payload(),
		ss.Headers(),
		ss.PublishStatus(),
		ss.AttemptCount(),
		ss.NextAttemptAt(),
		ss.PublishedAt(),
		ss.LastErrorCode(),
		ss.LastErrorDetail(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinanceOutboxEventsSelectFields() FinanceOutboxEventsSelectFields {
	return FinanceOutboxEventsSelectFields{}
}

type FinanceOutboxEventsUpdateFieldOption struct {
	useIncrement bool
}
type FinanceOutboxEventsUpdateField struct {
	financeOutboxEventsField FinanceOutboxEventsField
	opt                      FinanceOutboxEventsUpdateFieldOption
	value                    interface{}
}
type FinanceOutboxEventsUpdateFieldList []FinanceOutboxEventsUpdateField

func defaultFinanceOutboxEventsUpdateFieldOption() FinanceOutboxEventsUpdateFieldOption {
	return FinanceOutboxEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceOutboxEventsOption(useIncrement bool) func(*FinanceOutboxEventsUpdateFieldOption) {
	return func(pcufo *FinanceOutboxEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceOutboxEventsUpdateField(field FinanceOutboxEventsField, val interface{}, opts ...func(*FinanceOutboxEventsUpdateFieldOption)) FinanceOutboxEventsUpdateField {
	defaultOpt := defaultFinanceOutboxEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceOutboxEventsUpdateField{
		financeOutboxEventsField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultFinanceOutboxEventsUpdateFields(financeOutboxEvents model.FinanceOutboxEvents) (financeOutboxEventsUpdateFieldList FinanceOutboxEventsUpdateFieldList) {
	selectFields := NewFinanceOutboxEventsSelectFields()
	financeOutboxEventsUpdateFieldList = append(financeOutboxEventsUpdateFieldList,
		NewFinanceOutboxEventsUpdateField(selectFields.Id(), financeOutboxEvents.Id),
		NewFinanceOutboxEventsUpdateField(selectFields.AggregateType(), financeOutboxEvents.AggregateType),
		NewFinanceOutboxEventsUpdateField(selectFields.AggregateId(), financeOutboxEvents.AggregateId),
		NewFinanceOutboxEventsUpdateField(selectFields.EventType(), financeOutboxEvents.EventType),
		NewFinanceOutboxEventsUpdateField(selectFields.EventVersion(), financeOutboxEvents.EventVersion),
		NewFinanceOutboxEventsUpdateField(selectFields.IdempotencyKey(), financeOutboxEvents.IdempotencyKey),
		NewFinanceOutboxEventsUpdateField(selectFields.CorrelationId(), financeOutboxEvents.CorrelationId),
		NewFinanceOutboxEventsUpdateField(selectFields.Payload(), financeOutboxEvents.Payload),
		NewFinanceOutboxEventsUpdateField(selectFields.Headers(), financeOutboxEvents.Headers),
		NewFinanceOutboxEventsUpdateField(selectFields.PublishStatus(), financeOutboxEvents.PublishStatus),
		NewFinanceOutboxEventsUpdateField(selectFields.AttemptCount(), financeOutboxEvents.AttemptCount),
		NewFinanceOutboxEventsUpdateField(selectFields.NextAttemptAt(), financeOutboxEvents.NextAttemptAt),
		NewFinanceOutboxEventsUpdateField(selectFields.PublishedAt(), financeOutboxEvents.PublishedAt),
		NewFinanceOutboxEventsUpdateField(selectFields.LastErrorCode(), financeOutboxEvents.LastErrorCode),
		NewFinanceOutboxEventsUpdateField(selectFields.LastErrorDetail(), financeOutboxEvents.LastErrorDetail),
		NewFinanceOutboxEventsUpdateField(selectFields.Metadata(), financeOutboxEvents.Metadata),
		NewFinanceOutboxEventsUpdateField(selectFields.MetaCreatedAt(), financeOutboxEvents.MetaCreatedAt),
		NewFinanceOutboxEventsUpdateField(selectFields.MetaCreatedBy(), financeOutboxEvents.MetaCreatedBy),
		NewFinanceOutboxEventsUpdateField(selectFields.MetaUpdatedAt(), financeOutboxEvents.MetaUpdatedAt),
		NewFinanceOutboxEventsUpdateField(selectFields.MetaUpdatedBy(), financeOutboxEvents.MetaUpdatedBy),
		NewFinanceOutboxEventsUpdateField(selectFields.MetaDeletedAt(), financeOutboxEvents.MetaDeletedAt),
		NewFinanceOutboxEventsUpdateField(selectFields.MetaDeletedBy(), financeOutboxEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceOutboxEventsCommand(financeOutboxEventsUpdateFieldList FinanceOutboxEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeOutboxEventsUpdateFieldList {
		field := string(updateField.financeOutboxEventsField)
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

func (repo *RepositoryImpl) BulkCreateFinanceOutboxEvents(ctx context.Context, financeOutboxEventsList []*model.FinanceOutboxEvents, fieldsInsert ...FinanceOutboxEventsField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.FinanceOutboxEventsPrimaryID
		financeOutboxEventsValueList []model.FinanceOutboxEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceOutboxEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeOutboxEvents := range financeOutboxEventsList {

		primaryIds = append(primaryIds, financeOutboxEvents.ToFinanceOutboxEventsPrimaryID())

		financeOutboxEventsValueList = append(financeOutboxEventsValueList, *financeOutboxEvents)
	}

	_, notFoundIds, err := repo.IsExistFinanceOutboxEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceOutboxEvents] failed checking financeOutboxEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceOutboxEventsPrimaryID{}
		mapNotFoundIds := map[model.FinanceOutboxEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeOutboxEvents", fmt.Sprintf("financeOutboxEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceOutboxEvents(financeOutboxEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeOutboxEventsQueries.insertFinanceOutboxEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceOutboxEvents] failed exec create financeOutboxEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceOutboxEventsByIDs(ctx context.Context, primaryIDs []model.FinanceOutboxEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceOutboxEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceOutboxEventsByIDs] failed checking financeOutboxEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeOutboxEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_outbox_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeOutboxEventsQueries.deleteFinanceOutboxEvents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceOutboxEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceOutboxEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceOutboxEventsByIDs(ctx context.Context, ids []model.FinanceOutboxEventsPrimaryID) (exists bool, notFoundIds []model.FinanceOutboxEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_outbox_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeOutboxEventsQueries.selectFinanceOutboxEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceOutboxEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceOutboxEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceOutboxEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceOutboxEventsPrimaryID]bool{}
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

// BulkUpdateFinanceOutboxEvents is used to bulk update financeOutboxEvents, by default it will update all field
// if want to update specific field, then fill financeOutboxEventssMapUpdateFieldsRequest else please fill financeOutboxEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceOutboxEvents(ctx context.Context, financeOutboxEventssMap map[model.FinanceOutboxEventsPrimaryID]*model.FinanceOutboxEvents, financeOutboxEventssMapUpdateFieldsRequest map[model.FinanceOutboxEventsPrimaryID]FinanceOutboxEventsUpdateFieldList) (err error) {
	if len(financeOutboxEventssMap) == 0 && len(financeOutboxEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeOutboxEventssMapUpdateField map[model.FinanceOutboxEventsPrimaryID]FinanceOutboxEventsUpdateFieldList = map[model.FinanceOutboxEventsPrimaryID]FinanceOutboxEventsUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(financeOutboxEventssMap) > 0 {
		for id, financeOutboxEvents := range financeOutboxEventssMap {
			if financeOutboxEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceOutboxEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeOutboxEventssMapUpdateField[id] = defaultFinanceOutboxEventsUpdateFields(*financeOutboxEvents)
		}
	} else {
		financeOutboxEventssMapUpdateField = financeOutboxEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceOutboxEventsQuery(financeOutboxEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceOutboxEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceOutboxEvents] failed checking financeOutboxEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeOutboxEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceOutboxEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_outbox_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceOutboxEvents] failed exec query")
	}
	return
}

type FinanceOutboxEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceOutboxEventsFieldParameter(param string, args ...interface{}) FinanceOutboxEventsFieldParameter {
	return FinanceOutboxEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceOutboxEventsQuery(mapFinanceOutboxEventss map[model.FinanceOutboxEventsPrimaryID]FinanceOutboxEventsUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceOutboxEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceOutboxEventsPrimaryID]map[string]interface{}{}
	financeOutboxEventsSelectFields := NewFinanceOutboxEventsSelectFields()
	for id, updateFields := range mapFinanceOutboxEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeOutboxEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceOutboxEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceOutboxEventsFieldType(updateField.financeOutboxEventsField)))
			args = append(args, fields[string(updateField.financeOutboxEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeOutboxEventsField))
		if updateField.financeOutboxEventsField == financeOutboxEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeOutboxEventsField, asTableValues, updateField.financeOutboxEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeOutboxEventsField,
				"\"finance_outbox_events\"", updateField.financeOutboxEventsField,
				asTableValues, updateField.financeOutboxEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceOutboxEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceOutboxEventsPrimaryID, asTableValue string) (whereQry string) {
	financeOutboxEventsSelectFields := NewFinanceOutboxEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_outbox_events\".\"id\" = %s.\"id\"::"+GetFinanceOutboxEventsFieldType(financeOutboxEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceOutboxEventsFieldType(financeOutboxEventsField FinanceOutboxEventsField) string {
	selectFinanceOutboxEventsFields := NewFinanceOutboxEventsSelectFields()
	switch financeOutboxEventsField {

	case selectFinanceOutboxEventsFields.Id():
		return "uuid"

	case selectFinanceOutboxEventsFields.AggregateType():
		return "text"

	case selectFinanceOutboxEventsFields.AggregateId():
		return "uuid"

	case selectFinanceOutboxEventsFields.EventType():
		return "text"

	case selectFinanceOutboxEventsFields.EventVersion():
		return "int4"

	case selectFinanceOutboxEventsFields.IdempotencyKey():
		return "text"

	case selectFinanceOutboxEventsFields.CorrelationId():
		return "uuid"

	case selectFinanceOutboxEventsFields.Payload():
		return "jsonb"

	case selectFinanceOutboxEventsFields.Headers():
		return "jsonb"

	case selectFinanceOutboxEventsFields.PublishStatus():
		return "publish_status_enum"

	case selectFinanceOutboxEventsFields.AttemptCount():
		return "int4"

	case selectFinanceOutboxEventsFields.NextAttemptAt():
		return "timestamptz"

	case selectFinanceOutboxEventsFields.PublishedAt():
		return "timestamptz"

	case selectFinanceOutboxEventsFields.LastErrorCode():
		return "text"

	case selectFinanceOutboxEventsFields.LastErrorDetail():
		return "text"

	case selectFinanceOutboxEventsFields.Metadata():
		return "jsonb"

	case selectFinanceOutboxEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceOutboxEventsFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceOutboxEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceOutboxEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceOutboxEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceOutboxEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceOutboxEvents(ctx context.Context, financeOutboxEvents *model.FinanceOutboxEvents, fieldsInsert ...FinanceOutboxEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceOutboxEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceOutboxEventsPrimaryID{
		Id: financeOutboxEvents.Id,
	}
	exists, err := repo.IsExistFinanceOutboxEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceOutboxEvents] failed checking financeOutboxEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeOutboxEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceOutboxEvents([]model.FinanceOutboxEvents{*financeOutboxEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeOutboxEventsQueries.insertFinanceOutboxEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceOutboxEvents] failed exec create financeOutboxEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceOutboxEventsByID(ctx context.Context, primaryID model.FinanceOutboxEventsPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceOutboxEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceOutboxEventsByID] failed checking financeOutboxEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeOutboxEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceOutboxEventsCompositePrimaryKeyWhere([]model.FinanceOutboxEventsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeOutboxEventsQueries.deleteFinanceOutboxEvents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceOutboxEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceOutboxEventsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceOutboxEventsFilterResult, err error) {
	query, args, err := composeFinanceOutboxEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceOutboxEventsByFilter] failed compose financeOutboxEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceOutboxEventsByFilter] failed get financeOutboxEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceOutboxEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceOutboxEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceOutboxEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceOutboxEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceOutboxEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 22 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceOutboxEventsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 22+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["aggregate_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"aggregate_type\"")
			selectedColumns["aggregate_type"] = struct{}{}
		}
		if _, selected := selectedColumns["aggregate_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"aggregate_id\"")
			selectedColumns["aggregate_id"] = struct{}{}
		}
		if _, selected := selectedColumns["event_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_type\"")
			selectedColumns["event_type"] = struct{}{}
		}
		if _, selected := selectedColumns["event_version"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_version\"")
			selectedColumns["event_version"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["correlation_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"correlation_id\"")
			selectedColumns["correlation_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payload"]; !selected {
			selectColumns = append(selectColumns, "base.\"payload\"")
			selectedColumns["payload"] = struct{}{}
		}
		if _, selected := selectedColumns["headers"]; !selected {
			selectColumns = append(selectColumns, "base.\"headers\"")
			selectedColumns["headers"] = struct{}{}
		}
		if _, selected := selectedColumns["publish_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"publish_status\"")
			selectedColumns["publish_status"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_count\"")
			selectedColumns["attempt_count"] = struct{}{}
		}
		if _, selected := selectedColumns["next_attempt_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"next_attempt_at\"")
			selectedColumns["next_attempt_at"] = struct{}{}
		}
		if _, selected := selectedColumns["published_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"published_at\"")
			selectedColumns["published_at"] = struct{}{}
		}
		if _, selected := selectedColumns["last_error_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"last_error_code\"")
			selectedColumns["last_error_code"] = struct{}{}
		}
		if _, selected := selectedColumns["last_error_detail"]; !selected {
			selectColumns = append(selectColumns, "base.\"last_error_detail\"")
			selectedColumns["last_error_detail"] = struct{}{}
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

type financeOutboxEventsFilterPlaceholder struct {
	index int
}

func (p *financeOutboxEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceOutboxEventsFilterPredicate(filterField model.FilterField, placeholders *financeOutboxEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceOutboxEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceOutboxEventsFilterSQLExpr(spec)
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

func composeFinanceOutboxEventsFilterGroup(group model.FilterGroup, placeholders *financeOutboxEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceOutboxEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceOutboxEventsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceOutboxEventsFilterWhereQueries(filter model.Filter, placeholders *financeOutboxEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceOutboxEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceOutboxEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceOutboxEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceOutboxEventsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceOutboxEventsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceOutboxEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeOutboxEventsFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceOutboxEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceOutboxEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceOutboxEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceOutboxEventsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_outbox_events\" base%s", strings.Join(selectColumns, ","), composeFinanceOutboxEventsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceOutboxEventsByID(ctx context.Context, primaryID model.FinanceOutboxEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceOutboxEventsCompositePrimaryKeyWhere([]model.FinanceOutboxEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeOutboxEventsQueries.selectCountFinanceOutboxEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceOutboxEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceOutboxEvents(ctx context.Context, selectFields ...FinanceOutboxEventsField) (financeOutboxEventsList model.FinanceOutboxEventsList, err error) {
	var (
		defaultFinanceOutboxEventsSelectFields = defaultFinanceOutboxEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceOutboxEventsSelectFields = composeFinanceOutboxEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeOutboxEventsQueries.selectFinanceOutboxEvents, defaultFinanceOutboxEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeOutboxEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceOutboxEvents] failed get financeOutboxEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceOutboxEventsByID(ctx context.Context, primaryID model.FinanceOutboxEventsPrimaryID, selectFields ...FinanceOutboxEventsField) (financeOutboxEvents model.FinanceOutboxEvents, err error) {
	var (
		defaultFinanceOutboxEventsSelectFields = defaultFinanceOutboxEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceOutboxEventsSelectFields = composeFinanceOutboxEventsSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceOutboxEventsCompositePrimaryKeyWhere([]model.FinanceOutboxEventsPrimaryID{primaryID})
	query := fmt.Sprintf(financeOutboxEventsQueries.selectFinanceOutboxEvents+" WHERE "+whereQry, defaultFinanceOutboxEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeOutboxEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeOutboxEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceOutboxEventsByID] failed get financeOutboxEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceOutboxEventsByID(ctx context.Context, primaryID model.FinanceOutboxEventsPrimaryID, financeOutboxEvents *model.FinanceOutboxEvents, financeOutboxEventsUpdateFields ...FinanceOutboxEventsUpdateField) (err error) {
	exists, err := repo.IsExistFinanceOutboxEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceOutboxEvents] failed checking financeOutboxEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeOutboxEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeOutboxEvents == nil {
		if len(financeOutboxEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceOutboxEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeOutboxEvents = &model.FinanceOutboxEvents{}
	}
	var (
		defaultFinanceOutboxEventsUpdateFields = defaultFinanceOutboxEventsUpdateFields(*financeOutboxEvents)
		tempUpdateField                        FinanceOutboxEventsUpdateFieldList
		selectFields                           = NewFinanceOutboxEventsSelectFields()
	)
	if len(financeOutboxEventsUpdateFields) > 0 {
		for _, updateField := range financeOutboxEventsUpdateFields {
			if updateField.financeOutboxEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceOutboxEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceOutboxEventsCompositePrimaryKeyWhere([]model.FinanceOutboxEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceOutboxEventsCommand(defaultFinanceOutboxEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeOutboxEventsQueries.updateFinanceOutboxEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceOutboxEvents] error when try to update financeOutboxEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceOutboxEventsByFilter(ctx context.Context, filter model.Filter, financeOutboxEventsUpdateFields ...FinanceOutboxEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeOutboxEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceOutboxEventsUpdateFieldList
		selectFields = NewFinanceOutboxEventsSelectFields()
	)
	for _, updateField := range financeOutboxEventsUpdateFields {
		if updateField.financeOutboxEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceOutboxEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeOutboxEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceOutboxEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_outbox_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceOutboxEventsByFilter] error when try to update financeOutboxEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceOutboxEventsByFilter] failed get rows affected")
	}
	return
}

var (
	financeOutboxEventsQueries = struct {
		selectFinanceOutboxEvents      string
		selectCountFinanceOutboxEvents string
		deleteFinanceOutboxEvents      string
		updateFinanceOutboxEvents      string
		insertFinanceOutboxEvents      string
	}{
		selectFinanceOutboxEvents:      "SELECT %s FROM \"finance_outbox_events\"",
		selectCountFinanceOutboxEvents: "SELECT COUNT(\"id\") FROM \"finance_outbox_events\"",
		deleteFinanceOutboxEvents:      "DELETE FROM \"finance_outbox_events\"",
		updateFinanceOutboxEvents:      "UPDATE \"finance_outbox_events\" SET %s ",
		insertFinanceOutboxEvents:      "INSERT INTO \"finance_outbox_events\" %s VALUES %s",
	}
)

type FinanceOutboxEventsRepository interface {
	CreateFinanceOutboxEvents(ctx context.Context, financeOutboxEvents *model.FinanceOutboxEvents, fieldsInsert ...FinanceOutboxEventsField) error
	BulkCreateFinanceOutboxEvents(ctx context.Context, financeOutboxEventsList []*model.FinanceOutboxEvents, fieldsInsert ...FinanceOutboxEventsField) error
	ResolveFinanceOutboxEvents(ctx context.Context, selectFields ...FinanceOutboxEventsField) (model.FinanceOutboxEventsList, error)
	ResolveFinanceOutboxEventsByID(ctx context.Context, primaryID model.FinanceOutboxEventsPrimaryID, selectFields ...FinanceOutboxEventsField) (model.FinanceOutboxEvents, error)
	UpdateFinanceOutboxEventsByID(ctx context.Context, id model.FinanceOutboxEventsPrimaryID, financeOutboxEvents *model.FinanceOutboxEvents, financeOutboxEventsUpdateFields ...FinanceOutboxEventsUpdateField) error
	UpdateFinanceOutboxEventsByFilter(ctx context.Context, filter model.Filter, financeOutboxEventsUpdateFields ...FinanceOutboxEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceOutboxEvents(ctx context.Context, financeOutboxEventsListMap map[model.FinanceOutboxEventsPrimaryID]*model.FinanceOutboxEvents, FinanceOutboxEventssMapUpdateFieldsRequest map[model.FinanceOutboxEventsPrimaryID]FinanceOutboxEventsUpdateFieldList) (err error)
	DeleteFinanceOutboxEventsByID(ctx context.Context, id model.FinanceOutboxEventsPrimaryID) error
	BulkDeleteFinanceOutboxEventsByIDs(ctx context.Context, ids []model.FinanceOutboxEventsPrimaryID) error
	ResolveFinanceOutboxEventsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceOutboxEventsFilterResult, err error)
	IsExistFinanceOutboxEventsByIDs(ctx context.Context, ids []model.FinanceOutboxEventsPrimaryID) (exists bool, notFoundIds []model.FinanceOutboxEventsPrimaryID, err error)
	IsExistFinanceOutboxEventsByID(ctx context.Context, id model.FinanceOutboxEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
