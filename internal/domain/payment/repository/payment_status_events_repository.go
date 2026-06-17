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

func composeInsertFieldsAndParamsPaymentStatusEvents(paymentStatusEventsList []model.PaymentStatusEvents, fieldsInsert ...PaymentStatusEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentStatusEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentStatusEvents := range paymentStatusEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentStatusEvents.Id)
			case selectField.PaymentIntentId():
				args = append(args, paymentStatusEvents.PaymentIntentId)
			case selectField.PaymentAttemptId():
				args = append(args, paymentStatusEvents.PaymentAttemptId)
			case selectField.ProviderWebhookEventId():
				args = append(args, paymentStatusEvents.ProviderWebhookEventId)
			case selectField.SourceType():
				args = append(args, paymentStatusEvents.SourceType)
			case selectField.EventType():
				args = append(args, paymentStatusEvents.EventType)
			case selectField.OldIntentStatus():
				args = append(args, paymentStatusEvents.OldIntentStatus)
			case selectField.NewIntentStatus():
				args = append(args, paymentStatusEvents.NewIntentStatus)
			case selectField.OldAttemptStatus():
				args = append(args, paymentStatusEvents.OldAttemptStatus)
			case selectField.NewAttemptStatus():
				args = append(args, paymentStatusEvents.NewAttemptStatus)
			case selectField.ProviderStatus():
				args = append(args, paymentStatusEvents.ProviderStatus)
			case selectField.Reason():
				args = append(args, paymentStatusEvents.Reason)
			case selectField.OccurredAt():
				args = append(args, paymentStatusEvents.OccurredAt)
			case selectField.Metadata():
				args = append(args, paymentStatusEvents.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentStatusEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentStatusEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentStatusEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentStatusEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentStatusEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentStatusEvents.MetaDeletedBy)

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

func composePaymentStatusEventsCompositePrimaryKeyWhere(primaryIDs []model.PaymentStatusEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_status_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentStatusEventsSelectFields() string {
	fields := NewPaymentStatusEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentStatusEventsSelectFields(selectFields ...PaymentStatusEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentStatusEventsField string
type PaymentStatusEventsFieldList []PaymentStatusEventsField

type PaymentStatusEventsSelectFields struct {
}

func (ss PaymentStatusEventsSelectFields) Id() PaymentStatusEventsField {
	return PaymentStatusEventsField("id")
}

func (ss PaymentStatusEventsSelectFields) PaymentIntentId() PaymentStatusEventsField {
	return PaymentStatusEventsField("payment_intent_id")
}

func (ss PaymentStatusEventsSelectFields) PaymentAttemptId() PaymentStatusEventsField {
	return PaymentStatusEventsField("payment_attempt_id")
}

func (ss PaymentStatusEventsSelectFields) ProviderWebhookEventId() PaymentStatusEventsField {
	return PaymentStatusEventsField("provider_webhook_event_id")
}

func (ss PaymentStatusEventsSelectFields) SourceType() PaymentStatusEventsField {
	return PaymentStatusEventsField("source_type")
}

func (ss PaymentStatusEventsSelectFields) EventType() PaymentStatusEventsField {
	return PaymentStatusEventsField("event_type")
}

func (ss PaymentStatusEventsSelectFields) OldIntentStatus() PaymentStatusEventsField {
	return PaymentStatusEventsField("old_intent_status")
}

func (ss PaymentStatusEventsSelectFields) NewIntentStatus() PaymentStatusEventsField {
	return PaymentStatusEventsField("new_intent_status")
}

func (ss PaymentStatusEventsSelectFields) OldAttemptStatus() PaymentStatusEventsField {
	return PaymentStatusEventsField("old_attempt_status")
}

func (ss PaymentStatusEventsSelectFields) NewAttemptStatus() PaymentStatusEventsField {
	return PaymentStatusEventsField("new_attempt_status")
}

func (ss PaymentStatusEventsSelectFields) ProviderStatus() PaymentStatusEventsField {
	return PaymentStatusEventsField("provider_status")
}

func (ss PaymentStatusEventsSelectFields) Reason() PaymentStatusEventsField {
	return PaymentStatusEventsField("reason")
}

func (ss PaymentStatusEventsSelectFields) OccurredAt() PaymentStatusEventsField {
	return PaymentStatusEventsField("occurred_at")
}

func (ss PaymentStatusEventsSelectFields) Metadata() PaymentStatusEventsField {
	return PaymentStatusEventsField("metadata")
}

func (ss PaymentStatusEventsSelectFields) MetaCreatedAt() PaymentStatusEventsField {
	return PaymentStatusEventsField("meta_created_at")
}

func (ss PaymentStatusEventsSelectFields) MetaCreatedBy() PaymentStatusEventsField {
	return PaymentStatusEventsField("meta_created_by")
}

func (ss PaymentStatusEventsSelectFields) MetaUpdatedAt() PaymentStatusEventsField {
	return PaymentStatusEventsField("meta_updated_at")
}

func (ss PaymentStatusEventsSelectFields) MetaUpdatedBy() PaymentStatusEventsField {
	return PaymentStatusEventsField("meta_updated_by")
}

func (ss PaymentStatusEventsSelectFields) MetaDeletedAt() PaymentStatusEventsField {
	return PaymentStatusEventsField("meta_deleted_at")
}

func (ss PaymentStatusEventsSelectFields) MetaDeletedBy() PaymentStatusEventsField {
	return PaymentStatusEventsField("meta_deleted_by")
}

func (ss PaymentStatusEventsSelectFields) All() PaymentStatusEventsFieldList {
	return []PaymentStatusEventsField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.PaymentAttemptId(),
		ss.ProviderWebhookEventId(),
		ss.SourceType(),
		ss.EventType(),
		ss.OldIntentStatus(),
		ss.NewIntentStatus(),
		ss.OldAttemptStatus(),
		ss.NewAttemptStatus(),
		ss.ProviderStatus(),
		ss.Reason(),
		ss.OccurredAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentStatusEventsSelectFields() PaymentStatusEventsSelectFields {
	return PaymentStatusEventsSelectFields{}
}

type PaymentStatusEventsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentStatusEventsUpdateField struct {
	paymentStatusEventsField PaymentStatusEventsField
	opt                      PaymentStatusEventsUpdateFieldOption
	value                    interface{}
}
type PaymentStatusEventsUpdateFieldList []PaymentStatusEventsUpdateField

func defaultPaymentStatusEventsUpdateFieldOption() PaymentStatusEventsUpdateFieldOption {
	return PaymentStatusEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentStatusEventsOption(useIncrement bool) func(*PaymentStatusEventsUpdateFieldOption) {
	return func(pcufo *PaymentStatusEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentStatusEventsUpdateField(field PaymentStatusEventsField, val interface{}, opts ...func(*PaymentStatusEventsUpdateFieldOption)) PaymentStatusEventsUpdateField {
	defaultOpt := defaultPaymentStatusEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentStatusEventsUpdateField{
		paymentStatusEventsField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultPaymentStatusEventsUpdateFields(paymentStatusEvents model.PaymentStatusEvents) (paymentStatusEventsUpdateFieldList PaymentStatusEventsUpdateFieldList) {
	selectFields := NewPaymentStatusEventsSelectFields()
	paymentStatusEventsUpdateFieldList = append(paymentStatusEventsUpdateFieldList,
		NewPaymentStatusEventsUpdateField(selectFields.Id(), paymentStatusEvents.Id),
		NewPaymentStatusEventsUpdateField(selectFields.PaymentIntentId(), paymentStatusEvents.PaymentIntentId),
		NewPaymentStatusEventsUpdateField(selectFields.PaymentAttemptId(), paymentStatusEvents.PaymentAttemptId),
		NewPaymentStatusEventsUpdateField(selectFields.ProviderWebhookEventId(), paymentStatusEvents.ProviderWebhookEventId),
		NewPaymentStatusEventsUpdateField(selectFields.SourceType(), paymentStatusEvents.SourceType),
		NewPaymentStatusEventsUpdateField(selectFields.EventType(), paymentStatusEvents.EventType),
		NewPaymentStatusEventsUpdateField(selectFields.OldIntentStatus(), paymentStatusEvents.OldIntentStatus),
		NewPaymentStatusEventsUpdateField(selectFields.NewIntentStatus(), paymentStatusEvents.NewIntentStatus),
		NewPaymentStatusEventsUpdateField(selectFields.OldAttemptStatus(), paymentStatusEvents.OldAttemptStatus),
		NewPaymentStatusEventsUpdateField(selectFields.NewAttemptStatus(), paymentStatusEvents.NewAttemptStatus),
		NewPaymentStatusEventsUpdateField(selectFields.ProviderStatus(), paymentStatusEvents.ProviderStatus),
		NewPaymentStatusEventsUpdateField(selectFields.Reason(), paymentStatusEvents.Reason),
		NewPaymentStatusEventsUpdateField(selectFields.OccurredAt(), paymentStatusEvents.OccurredAt),
		NewPaymentStatusEventsUpdateField(selectFields.Metadata(), paymentStatusEvents.Metadata),
		NewPaymentStatusEventsUpdateField(selectFields.MetaCreatedAt(), paymentStatusEvents.MetaCreatedAt),
		NewPaymentStatusEventsUpdateField(selectFields.MetaCreatedBy(), paymentStatusEvents.MetaCreatedBy),
		NewPaymentStatusEventsUpdateField(selectFields.MetaUpdatedAt(), paymentStatusEvents.MetaUpdatedAt),
		NewPaymentStatusEventsUpdateField(selectFields.MetaUpdatedBy(), paymentStatusEvents.MetaUpdatedBy),
		NewPaymentStatusEventsUpdateField(selectFields.MetaDeletedAt(), paymentStatusEvents.MetaDeletedAt),
		NewPaymentStatusEventsUpdateField(selectFields.MetaDeletedBy(), paymentStatusEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentStatusEventsCommand(paymentStatusEventsUpdateFieldList PaymentStatusEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentStatusEventsUpdateFieldList {
		field := string(updateField.paymentStatusEventsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentStatusEvents(ctx context.Context, paymentStatusEventsList []*model.PaymentStatusEvents, fieldsInsert ...PaymentStatusEventsField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.PaymentStatusEventsPrimaryID
		paymentStatusEventsValueList []model.PaymentStatusEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentStatusEvents := range paymentStatusEventsList {

		primaryIds = append(primaryIds, paymentStatusEvents.ToPaymentStatusEventsPrimaryID())

		paymentStatusEventsValueList = append(paymentStatusEventsValueList, *paymentStatusEvents)
	}

	_, notFoundIds, err := repo.IsExistPaymentStatusEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentStatusEvents] failed checking paymentStatusEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentStatusEventsPrimaryID{}
		mapNotFoundIds := map[model.PaymentStatusEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentStatusEvents", fmt.Sprintf("paymentStatusEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentStatusEvents(paymentStatusEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentStatusEventsQueries.insertPaymentStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentStatusEvents] failed exec create paymentStatusEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentStatusEventsByIDs(ctx context.Context, primaryIDs []model.PaymentStatusEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentStatusEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentStatusEventsByIDs] failed checking paymentStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentStatusEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_status_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentStatusEventsQueries.deletePaymentStatusEvents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentStatusEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentStatusEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentStatusEventsByIDs(ctx context.Context, ids []model.PaymentStatusEventsPrimaryID) (exists bool, notFoundIds []model.PaymentStatusEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_status_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentStatusEventsQueries.selectPaymentStatusEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentStatusEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentStatusEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentStatusEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentStatusEventsPrimaryID]bool{}
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

// BulkUpdatePaymentStatusEvents is used to bulk update paymentStatusEvents, by default it will update all field
// if want to update specific field, then fill paymentStatusEventssMapUpdateFieldsRequest else please fill paymentStatusEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentStatusEvents(ctx context.Context, paymentStatusEventssMap map[model.PaymentStatusEventsPrimaryID]*model.PaymentStatusEvents, paymentStatusEventssMapUpdateFieldsRequest map[model.PaymentStatusEventsPrimaryID]PaymentStatusEventsUpdateFieldList) (err error) {
	if len(paymentStatusEventssMap) == 0 && len(paymentStatusEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentStatusEventssMapUpdateField map[model.PaymentStatusEventsPrimaryID]PaymentStatusEventsUpdateFieldList = map[model.PaymentStatusEventsPrimaryID]PaymentStatusEventsUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(paymentStatusEventssMap) > 0 {
		for id, paymentStatusEvents := range paymentStatusEventssMap {
			if paymentStatusEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentStatusEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentStatusEventssMapUpdateField[id] = defaultPaymentStatusEventsUpdateFields(*paymentStatusEvents)
		}
	} else {
		paymentStatusEventssMapUpdateField = paymentStatusEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentStatusEventsQuery(paymentStatusEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentStatusEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentStatusEvents] failed checking paymentStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentStatusEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentStatusEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_status_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentStatusEvents] failed exec query")
	}
	return
}

type PaymentStatusEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentStatusEventsFieldParameter(param string, args ...interface{}) PaymentStatusEventsFieldParameter {
	return PaymentStatusEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentStatusEventsQuery(mapPaymentStatusEventss map[model.PaymentStatusEventsPrimaryID]PaymentStatusEventsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentStatusEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentStatusEventsPrimaryID]map[string]interface{}{}
	paymentStatusEventsSelectFields := NewPaymentStatusEventsSelectFields()
	for id, updateFields := range mapPaymentStatusEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentStatusEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentStatusEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentStatusEventsFieldType(updateField.paymentStatusEventsField)))
			args = append(args, fields[string(updateField.paymentStatusEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentStatusEventsField))
		if updateField.paymentStatusEventsField == paymentStatusEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentStatusEventsField, asTableValues, updateField.paymentStatusEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentStatusEventsField,
				"\"payment_status_events\"", updateField.paymentStatusEventsField,
				asTableValues, updateField.paymentStatusEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentStatusEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentStatusEventsPrimaryID, asTableValue string) (whereQry string) {
	paymentStatusEventsSelectFields := NewPaymentStatusEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_status_events\".\"id\" = %s.\"id\"::"+GetPaymentStatusEventsFieldType(paymentStatusEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentStatusEventsFieldType(paymentStatusEventsField PaymentStatusEventsField) string {
	selectPaymentStatusEventsFields := NewPaymentStatusEventsSelectFields()
	switch paymentStatusEventsField {

	case selectPaymentStatusEventsFields.Id():
		return "uuid"

	case selectPaymentStatusEventsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentStatusEventsFields.PaymentAttemptId():
		return "uuid"

	case selectPaymentStatusEventsFields.ProviderWebhookEventId():
		return "uuid"

	case selectPaymentStatusEventsFields.SourceType():
		return "text"

	case selectPaymentStatusEventsFields.EventType():
		return "text"

	case selectPaymentStatusEventsFields.OldIntentStatus():
		return "text"

	case selectPaymentStatusEventsFields.NewIntentStatus():
		return "text"

	case selectPaymentStatusEventsFields.OldAttemptStatus():
		return "text"

	case selectPaymentStatusEventsFields.NewAttemptStatus():
		return "text"

	case selectPaymentStatusEventsFields.ProviderStatus():
		return "text"

	case selectPaymentStatusEventsFields.Reason():
		return "text"

	case selectPaymentStatusEventsFields.OccurredAt():
		return "timestamptz"

	case selectPaymentStatusEventsFields.Metadata():
		return "jsonb"

	case selectPaymentStatusEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentStatusEventsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentStatusEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentStatusEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentStatusEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentStatusEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentStatusEvents(ctx context.Context, paymentStatusEvents *model.PaymentStatusEvents, fieldsInsert ...PaymentStatusEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentStatusEventsPrimaryID{
		Id: paymentStatusEvents.Id,
	}
	exists, err := repo.IsExistPaymentStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentStatusEvents] failed checking paymentStatusEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentStatusEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentStatusEvents([]model.PaymentStatusEvents{*paymentStatusEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentStatusEventsQueries.insertPaymentStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentStatusEvents] failed exec create paymentStatusEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentStatusEventsByID(ctx context.Context, primaryID model.PaymentStatusEventsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentStatusEventsByID] failed checking paymentStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentStatusEventsCompositePrimaryKeyWhere([]model.PaymentStatusEventsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentStatusEventsQueries.deletePaymentStatusEvents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentStatusEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentStatusEventsFilterResult, err error) {
	query, args, err := composePaymentStatusEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentStatusEventsByFilter] failed compose paymentStatusEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentStatusEventsByFilter] failed get paymentStatusEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentStatusEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentStatusEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentStatusEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentStatusEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentStatusEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentStatusEventsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 20+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_attempt_id\"")
			selectedColumns["payment_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_webhook_event_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_webhook_event_id\"")
			selectedColumns["provider_webhook_event_id"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["event_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_type\"")
			selectedColumns["event_type"] = struct{}{}
		}
		if _, selected := selectedColumns["old_intent_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"old_intent_status\"")
			selectedColumns["old_intent_status"] = struct{}{}
		}
		if _, selected := selectedColumns["new_intent_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"new_intent_status\"")
			selectedColumns["new_intent_status"] = struct{}{}
		}
		if _, selected := selectedColumns["old_attempt_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"old_attempt_status\"")
			selectedColumns["old_attempt_status"] = struct{}{}
		}
		if _, selected := selectedColumns["new_attempt_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"new_attempt_status\"")
			selectedColumns["new_attempt_status"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_status\"")
			selectedColumns["provider_status"] = struct{}{}
		}
		if _, selected := selectedColumns["reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason\"")
			selectedColumns["reason"] = struct{}{}
		}
		if _, selected := selectedColumns["occurred_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"occurred_at\"")
			selectedColumns["occurred_at"] = struct{}{}
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

type paymentStatusEventsFilterPlaceholder struct {
	index int
}

func (p *paymentStatusEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentStatusEventsFilterPredicate(filterField model.FilterField, placeholders *paymentStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentStatusEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentStatusEventsFilterSQLExpr(spec)
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

func composePaymentStatusEventsFilterGroup(group model.FilterGroup, placeholders *paymentStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentStatusEventsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentStatusEventsFilterWhereQueries(filter model.Filter, placeholders *paymentStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentStatusEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentStatusEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentStatusEventsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentStatusEventsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentStatusEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentStatusEventsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentStatusEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentStatusEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentStatusEventsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_status_events\" base%s", strings.Join(selectColumns, ","), composePaymentStatusEventsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentStatusEventsByID(ctx context.Context, primaryID model.PaymentStatusEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentStatusEventsCompositePrimaryKeyWhere([]model.PaymentStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentStatusEventsQueries.selectCountPaymentStatusEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentStatusEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentStatusEvents(ctx context.Context, selectFields ...PaymentStatusEventsField) (paymentStatusEventsList model.PaymentStatusEventsList, err error) {
	var (
		defaultPaymentStatusEventsSelectFields = defaultPaymentStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentStatusEventsSelectFields = composePaymentStatusEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentStatusEventsQueries.selectPaymentStatusEvents, defaultPaymentStatusEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentStatusEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentStatusEvents] failed get paymentStatusEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentStatusEventsByID(ctx context.Context, primaryID model.PaymentStatusEventsPrimaryID, selectFields ...PaymentStatusEventsField) (paymentStatusEvents model.PaymentStatusEvents, err error) {
	var (
		defaultPaymentStatusEventsSelectFields = defaultPaymentStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentStatusEventsSelectFields = composePaymentStatusEventsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentStatusEventsCompositePrimaryKeyWhere([]model.PaymentStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentStatusEventsQueries.selectPaymentStatusEvents+" WHERE "+whereQry, defaultPaymentStatusEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentStatusEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentStatusEventsByID] failed get paymentStatusEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentStatusEventsByID(ctx context.Context, primaryID model.PaymentStatusEventsPrimaryID, paymentStatusEvents *model.PaymentStatusEvents, paymentStatusEventsUpdateFields ...PaymentStatusEventsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentStatusEvents] failed checking paymentStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentStatusEvents == nil {
		if len(paymentStatusEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentStatusEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentStatusEvents = &model.PaymentStatusEvents{}
	}
	var (
		defaultPaymentStatusEventsUpdateFields = defaultPaymentStatusEventsUpdateFields(*paymentStatusEvents)
		tempUpdateField                        PaymentStatusEventsUpdateFieldList
		selectFields                           = NewPaymentStatusEventsSelectFields()
	)
	if len(paymentStatusEventsUpdateFields) > 0 {
		for _, updateField := range paymentStatusEventsUpdateFields {
			if updateField.paymentStatusEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentStatusEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentStatusEventsCompositePrimaryKeyWhere([]model.PaymentStatusEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentStatusEventsCommand(defaultPaymentStatusEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentStatusEventsQueries.updatePaymentStatusEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentStatusEvents] error when try to update paymentStatusEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentStatusEventsByFilter(ctx context.Context, filter model.Filter, paymentStatusEventsUpdateFields ...PaymentStatusEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentStatusEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentStatusEventsUpdateFieldList
		selectFields = NewPaymentStatusEventsSelectFields()
	)
	for _, updateField := range paymentStatusEventsUpdateFields {
		if updateField.paymentStatusEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentStatusEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentStatusEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_status_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentStatusEventsByFilter] error when try to update paymentStatusEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentStatusEventsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentStatusEventsQueries = struct {
		selectPaymentStatusEvents      string
		selectCountPaymentStatusEvents string
		deletePaymentStatusEvents      string
		updatePaymentStatusEvents      string
		insertPaymentStatusEvents      string
	}{
		selectPaymentStatusEvents:      "SELECT %s FROM \"payment_status_events\"",
		selectCountPaymentStatusEvents: "SELECT COUNT(\"id\") FROM \"payment_status_events\"",
		deletePaymentStatusEvents:      "DELETE FROM \"payment_status_events\"",
		updatePaymentStatusEvents:      "UPDATE \"payment_status_events\" SET %s ",
		insertPaymentStatusEvents:      "INSERT INTO \"payment_status_events\" %s VALUES %s",
	}
)

type PaymentStatusEventsRepository interface {
	CreatePaymentStatusEvents(ctx context.Context, paymentStatusEvents *model.PaymentStatusEvents, fieldsInsert ...PaymentStatusEventsField) error
	BulkCreatePaymentStatusEvents(ctx context.Context, paymentStatusEventsList []*model.PaymentStatusEvents, fieldsInsert ...PaymentStatusEventsField) error
	ResolvePaymentStatusEvents(ctx context.Context, selectFields ...PaymentStatusEventsField) (model.PaymentStatusEventsList, error)
	ResolvePaymentStatusEventsByID(ctx context.Context, primaryID model.PaymentStatusEventsPrimaryID, selectFields ...PaymentStatusEventsField) (model.PaymentStatusEvents, error)
	UpdatePaymentStatusEventsByID(ctx context.Context, id model.PaymentStatusEventsPrimaryID, paymentStatusEvents *model.PaymentStatusEvents, paymentStatusEventsUpdateFields ...PaymentStatusEventsUpdateField) error
	UpdatePaymentStatusEventsByFilter(ctx context.Context, filter model.Filter, paymentStatusEventsUpdateFields ...PaymentStatusEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentStatusEvents(ctx context.Context, paymentStatusEventsListMap map[model.PaymentStatusEventsPrimaryID]*model.PaymentStatusEvents, PaymentStatusEventssMapUpdateFieldsRequest map[model.PaymentStatusEventsPrimaryID]PaymentStatusEventsUpdateFieldList) (err error)
	DeletePaymentStatusEventsByID(ctx context.Context, id model.PaymentStatusEventsPrimaryID) error
	BulkDeletePaymentStatusEventsByIDs(ctx context.Context, ids []model.PaymentStatusEventsPrimaryID) error
	ResolvePaymentStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentStatusEventsFilterResult, err error)
	IsExistPaymentStatusEventsByIDs(ctx context.Context, ids []model.PaymentStatusEventsPrimaryID) (exists bool, notFoundIds []model.PaymentStatusEventsPrimaryID, err error)
	IsExistPaymentStatusEventsByID(ctx context.Context, id model.PaymentStatusEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
