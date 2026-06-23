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

func composeInsertFieldsAndParamsFinanceAuditEvents(financeAuditEventsList []model.FinanceAuditEvents, fieldsInsert ...FinanceAuditEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceAuditEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeAuditEvents := range financeAuditEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeAuditEvents.Id)
			case selectField.AggregateType():
				args = append(args, financeAuditEvents.AggregateType)
			case selectField.AggregateId():
				args = append(args, financeAuditEvents.AggregateId)
			case selectField.EventType():
				args = append(args, financeAuditEvents.EventType)
			case selectField.ActorUserId():
				args = append(args, financeAuditEvents.ActorUserId)
			case selectField.ActorType():
				args = append(args, financeAuditEvents.ActorType)
			case selectField.EventAt():
				args = append(args, financeAuditEvents.EventAt)
			case selectField.OldState():
				args = append(args, financeAuditEvents.OldState)
			case selectField.NewState():
				args = append(args, financeAuditEvents.NewState)
			case selectField.CorrelationId():
				args = append(args, financeAuditEvents.CorrelationId)
			case selectField.CausationId():
				args = append(args, financeAuditEvents.CausationId)
			case selectField.Metadata():
				args = append(args, financeAuditEvents.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeAuditEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeAuditEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeAuditEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeAuditEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeAuditEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeAuditEvents.MetaDeletedBy)

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

func composeFinanceAuditEventsCompositePrimaryKeyWhere(primaryIDs []model.FinanceAuditEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_audit_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceAuditEventsSelectFields() string {
	fields := NewFinanceAuditEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceAuditEventsSelectFields(selectFields ...FinanceAuditEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceAuditEventsField string
type FinanceAuditEventsFieldList []FinanceAuditEventsField

type FinanceAuditEventsSelectFields struct {
}

func (ss FinanceAuditEventsSelectFields) Id() FinanceAuditEventsField {
	return FinanceAuditEventsField("id")
}

func (ss FinanceAuditEventsSelectFields) AggregateType() FinanceAuditEventsField {
	return FinanceAuditEventsField("aggregate_type")
}

func (ss FinanceAuditEventsSelectFields) AggregateId() FinanceAuditEventsField {
	return FinanceAuditEventsField("aggregate_id")
}

func (ss FinanceAuditEventsSelectFields) EventType() FinanceAuditEventsField {
	return FinanceAuditEventsField("event_type")
}

func (ss FinanceAuditEventsSelectFields) ActorUserId() FinanceAuditEventsField {
	return FinanceAuditEventsField("actor_user_id")
}

func (ss FinanceAuditEventsSelectFields) ActorType() FinanceAuditEventsField {
	return FinanceAuditEventsField("actor_type")
}

func (ss FinanceAuditEventsSelectFields) EventAt() FinanceAuditEventsField {
	return FinanceAuditEventsField("event_at")
}

func (ss FinanceAuditEventsSelectFields) OldState() FinanceAuditEventsField {
	return FinanceAuditEventsField("old_state")
}

func (ss FinanceAuditEventsSelectFields) NewState() FinanceAuditEventsField {
	return FinanceAuditEventsField("new_state")
}

func (ss FinanceAuditEventsSelectFields) CorrelationId() FinanceAuditEventsField {
	return FinanceAuditEventsField("correlation_id")
}

func (ss FinanceAuditEventsSelectFields) CausationId() FinanceAuditEventsField {
	return FinanceAuditEventsField("causation_id")
}

func (ss FinanceAuditEventsSelectFields) Metadata() FinanceAuditEventsField {
	return FinanceAuditEventsField("metadata")
}

func (ss FinanceAuditEventsSelectFields) MetaCreatedAt() FinanceAuditEventsField {
	return FinanceAuditEventsField("meta_created_at")
}

func (ss FinanceAuditEventsSelectFields) MetaCreatedBy() FinanceAuditEventsField {
	return FinanceAuditEventsField("meta_created_by")
}

func (ss FinanceAuditEventsSelectFields) MetaUpdatedAt() FinanceAuditEventsField {
	return FinanceAuditEventsField("meta_updated_at")
}

func (ss FinanceAuditEventsSelectFields) MetaUpdatedBy() FinanceAuditEventsField {
	return FinanceAuditEventsField("meta_updated_by")
}

func (ss FinanceAuditEventsSelectFields) MetaDeletedAt() FinanceAuditEventsField {
	return FinanceAuditEventsField("meta_deleted_at")
}

func (ss FinanceAuditEventsSelectFields) MetaDeletedBy() FinanceAuditEventsField {
	return FinanceAuditEventsField("meta_deleted_by")
}

func (ss FinanceAuditEventsSelectFields) All() FinanceAuditEventsFieldList {
	return []FinanceAuditEventsField{
		ss.Id(),
		ss.AggregateType(),
		ss.AggregateId(),
		ss.EventType(),
		ss.ActorUserId(),
		ss.ActorType(),
		ss.EventAt(),
		ss.OldState(),
		ss.NewState(),
		ss.CorrelationId(),
		ss.CausationId(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinanceAuditEventsSelectFields() FinanceAuditEventsSelectFields {
	return FinanceAuditEventsSelectFields{}
}

type FinanceAuditEventsUpdateFieldOption struct {
	useIncrement bool
}
type FinanceAuditEventsUpdateField struct {
	financeAuditEventsField FinanceAuditEventsField
	opt                     FinanceAuditEventsUpdateFieldOption
	value                   interface{}
}
type FinanceAuditEventsUpdateFieldList []FinanceAuditEventsUpdateField

func defaultFinanceAuditEventsUpdateFieldOption() FinanceAuditEventsUpdateFieldOption {
	return FinanceAuditEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceAuditEventsOption(useIncrement bool) func(*FinanceAuditEventsUpdateFieldOption) {
	return func(pcufo *FinanceAuditEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceAuditEventsUpdateField(field FinanceAuditEventsField, val interface{}, opts ...func(*FinanceAuditEventsUpdateFieldOption)) FinanceAuditEventsUpdateField {
	defaultOpt := defaultFinanceAuditEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceAuditEventsUpdateField{
		financeAuditEventsField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultFinanceAuditEventsUpdateFields(financeAuditEvents model.FinanceAuditEvents) (financeAuditEventsUpdateFieldList FinanceAuditEventsUpdateFieldList) {
	selectFields := NewFinanceAuditEventsSelectFields()
	financeAuditEventsUpdateFieldList = append(financeAuditEventsUpdateFieldList,
		NewFinanceAuditEventsUpdateField(selectFields.Id(), financeAuditEvents.Id),
		NewFinanceAuditEventsUpdateField(selectFields.AggregateType(), financeAuditEvents.AggregateType),
		NewFinanceAuditEventsUpdateField(selectFields.AggregateId(), financeAuditEvents.AggregateId),
		NewFinanceAuditEventsUpdateField(selectFields.EventType(), financeAuditEvents.EventType),
		NewFinanceAuditEventsUpdateField(selectFields.ActorUserId(), financeAuditEvents.ActorUserId),
		NewFinanceAuditEventsUpdateField(selectFields.ActorType(), financeAuditEvents.ActorType),
		NewFinanceAuditEventsUpdateField(selectFields.EventAt(), financeAuditEvents.EventAt),
		NewFinanceAuditEventsUpdateField(selectFields.OldState(), financeAuditEvents.OldState),
		NewFinanceAuditEventsUpdateField(selectFields.NewState(), financeAuditEvents.NewState),
		NewFinanceAuditEventsUpdateField(selectFields.CorrelationId(), financeAuditEvents.CorrelationId),
		NewFinanceAuditEventsUpdateField(selectFields.CausationId(), financeAuditEvents.CausationId),
		NewFinanceAuditEventsUpdateField(selectFields.Metadata(), financeAuditEvents.Metadata),
		NewFinanceAuditEventsUpdateField(selectFields.MetaCreatedAt(), financeAuditEvents.MetaCreatedAt),
		NewFinanceAuditEventsUpdateField(selectFields.MetaCreatedBy(), financeAuditEvents.MetaCreatedBy),
		NewFinanceAuditEventsUpdateField(selectFields.MetaUpdatedAt(), financeAuditEvents.MetaUpdatedAt),
		NewFinanceAuditEventsUpdateField(selectFields.MetaUpdatedBy(), financeAuditEvents.MetaUpdatedBy),
		NewFinanceAuditEventsUpdateField(selectFields.MetaDeletedAt(), financeAuditEvents.MetaDeletedAt),
		NewFinanceAuditEventsUpdateField(selectFields.MetaDeletedBy(), financeAuditEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceAuditEventsCommand(financeAuditEventsUpdateFieldList FinanceAuditEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeAuditEventsUpdateFieldList {
		field := string(updateField.financeAuditEventsField)
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

func (repo *RepositoryImpl) BulkCreateFinanceAuditEvents(ctx context.Context, financeAuditEventsList []*model.FinanceAuditEvents, fieldsInsert ...FinanceAuditEventsField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.FinanceAuditEventsPrimaryID
		financeAuditEventsValueList []model.FinanceAuditEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceAuditEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeAuditEvents := range financeAuditEventsList {

		primaryIds = append(primaryIds, financeAuditEvents.ToFinanceAuditEventsPrimaryID())

		financeAuditEventsValueList = append(financeAuditEventsValueList, *financeAuditEvents)
	}

	_, notFoundIds, err := repo.IsExistFinanceAuditEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceAuditEvents] failed checking financeAuditEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceAuditEventsPrimaryID{}
		mapNotFoundIds := map[model.FinanceAuditEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeAuditEvents", fmt.Sprintf("financeAuditEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceAuditEvents(financeAuditEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeAuditEventsQueries.insertFinanceAuditEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceAuditEvents] failed exec create financeAuditEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceAuditEventsByIDs(ctx context.Context, primaryIDs []model.FinanceAuditEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceAuditEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceAuditEventsByIDs] failed checking financeAuditEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeAuditEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_audit_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeAuditEventsQueries.deleteFinanceAuditEvents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceAuditEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceAuditEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceAuditEventsByIDs(ctx context.Context, ids []model.FinanceAuditEventsPrimaryID) (exists bool, notFoundIds []model.FinanceAuditEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_audit_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeAuditEventsQueries.selectFinanceAuditEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceAuditEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceAuditEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceAuditEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceAuditEventsPrimaryID]bool{}
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

// BulkUpdateFinanceAuditEvents is used to bulk update financeAuditEvents, by default it will update all field
// if want to update specific field, then fill financeAuditEventssMapUpdateFieldsRequest else please fill financeAuditEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceAuditEvents(ctx context.Context, financeAuditEventssMap map[model.FinanceAuditEventsPrimaryID]*model.FinanceAuditEvents, financeAuditEventssMapUpdateFieldsRequest map[model.FinanceAuditEventsPrimaryID]FinanceAuditEventsUpdateFieldList) (err error) {
	if len(financeAuditEventssMap) == 0 && len(financeAuditEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeAuditEventssMapUpdateField map[model.FinanceAuditEventsPrimaryID]FinanceAuditEventsUpdateFieldList = map[model.FinanceAuditEventsPrimaryID]FinanceAuditEventsUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(financeAuditEventssMap) > 0 {
		for id, financeAuditEvents := range financeAuditEventssMap {
			if financeAuditEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceAuditEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeAuditEventssMapUpdateField[id] = defaultFinanceAuditEventsUpdateFields(*financeAuditEvents)
		}
	} else {
		financeAuditEventssMapUpdateField = financeAuditEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceAuditEventsQuery(financeAuditEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceAuditEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceAuditEvents] failed checking financeAuditEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeAuditEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceAuditEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_audit_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceAuditEvents] failed exec query")
	}
	return
}

type FinanceAuditEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceAuditEventsFieldParameter(param string, args ...interface{}) FinanceAuditEventsFieldParameter {
	return FinanceAuditEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceAuditEventsQuery(mapFinanceAuditEventss map[model.FinanceAuditEventsPrimaryID]FinanceAuditEventsUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceAuditEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceAuditEventsPrimaryID]map[string]interface{}{}
	financeAuditEventsSelectFields := NewFinanceAuditEventsSelectFields()
	for id, updateFields := range mapFinanceAuditEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeAuditEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceAuditEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceAuditEventsFieldType(updateField.financeAuditEventsField)))
			args = append(args, fields[string(updateField.financeAuditEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeAuditEventsField))
		if updateField.financeAuditEventsField == financeAuditEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeAuditEventsField, asTableValues, updateField.financeAuditEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeAuditEventsField,
				"\"finance_audit_events\"", updateField.financeAuditEventsField,
				asTableValues, updateField.financeAuditEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceAuditEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceAuditEventsPrimaryID, asTableValue string) (whereQry string) {
	financeAuditEventsSelectFields := NewFinanceAuditEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_audit_events\".\"id\" = %s.\"id\"::"+GetFinanceAuditEventsFieldType(financeAuditEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceAuditEventsFieldType(financeAuditEventsField FinanceAuditEventsField) string {
	selectFinanceAuditEventsFields := NewFinanceAuditEventsSelectFields()
	switch financeAuditEventsField {

	case selectFinanceAuditEventsFields.Id():
		return "uuid"

	case selectFinanceAuditEventsFields.AggregateType():
		return "text"

	case selectFinanceAuditEventsFields.AggregateId():
		return "uuid"

	case selectFinanceAuditEventsFields.EventType():
		return "text"

	case selectFinanceAuditEventsFields.ActorUserId():
		return "uuid"

	case selectFinanceAuditEventsFields.ActorType():
		return "actor_type_enum"

	case selectFinanceAuditEventsFields.EventAt():
		return "timestamptz"

	case selectFinanceAuditEventsFields.OldState():
		return "jsonb"

	case selectFinanceAuditEventsFields.NewState():
		return "jsonb"

	case selectFinanceAuditEventsFields.CorrelationId():
		return "text"

	case selectFinanceAuditEventsFields.CausationId():
		return "text"

	case selectFinanceAuditEventsFields.Metadata():
		return "jsonb"

	case selectFinanceAuditEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceAuditEventsFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceAuditEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceAuditEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceAuditEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceAuditEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceAuditEvents(ctx context.Context, financeAuditEvents *model.FinanceAuditEvents, fieldsInsert ...FinanceAuditEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceAuditEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceAuditEventsPrimaryID{
		Id: financeAuditEvents.Id,
	}
	exists, err := repo.IsExistFinanceAuditEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceAuditEvents] failed checking financeAuditEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeAuditEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceAuditEvents([]model.FinanceAuditEvents{*financeAuditEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeAuditEventsQueries.insertFinanceAuditEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceAuditEvents] failed exec create financeAuditEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceAuditEventsByID(ctx context.Context, primaryID model.FinanceAuditEventsPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceAuditEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceAuditEventsByID] failed checking financeAuditEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeAuditEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceAuditEventsCompositePrimaryKeyWhere([]model.FinanceAuditEventsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeAuditEventsQueries.deleteFinanceAuditEvents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceAuditEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceAuditEventsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceAuditEventsFilterResult, err error) {
	query, args, err := composeFinanceAuditEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceAuditEventsByFilter] failed compose financeAuditEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceAuditEventsByFilter] failed get financeAuditEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceAuditEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceAuditEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceAuditEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceAuditEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceAuditEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceAuditEventsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 18+1)
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
		if _, selected := selectedColumns["actor_user_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"actor_user_id\"")
			selectedColumns["actor_user_id"] = struct{}{}
		}
		if _, selected := selectedColumns["actor_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"actor_type\"")
			selectedColumns["actor_type"] = struct{}{}
		}
		if _, selected := selectedColumns["event_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_at\"")
			selectedColumns["event_at"] = struct{}{}
		}
		if _, selected := selectedColumns["old_state"]; !selected {
			selectColumns = append(selectColumns, "base.\"old_state\"")
			selectedColumns["old_state"] = struct{}{}
		}
		if _, selected := selectedColumns["new_state"]; !selected {
			selectColumns = append(selectColumns, "base.\"new_state\"")
			selectedColumns["new_state"] = struct{}{}
		}
		if _, selected := selectedColumns["correlation_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"correlation_id\"")
			selectedColumns["correlation_id"] = struct{}{}
		}
		if _, selected := selectedColumns["causation_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"causation_id\"")
			selectedColumns["causation_id"] = struct{}{}
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

type financeAuditEventsFilterPlaceholder struct {
	index int
}

func (p *financeAuditEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceAuditEventsFilterPredicate(filterField model.FilterField, placeholders *financeAuditEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceAuditEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceAuditEventsFilterSQLExpr(spec)
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

func composeFinanceAuditEventsFilterGroup(group model.FilterGroup, placeholders *financeAuditEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceAuditEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceAuditEventsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceAuditEventsFilterWhereQueries(filter model.Filter, placeholders *financeAuditEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceAuditEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceAuditEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceAuditEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceAuditEventsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceAuditEventsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceAuditEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeAuditEventsFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceAuditEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceAuditEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceAuditEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceAuditEventsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_audit_events\" base%s", strings.Join(selectColumns, ","), composeFinanceAuditEventsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceAuditEventsByID(ctx context.Context, primaryID model.FinanceAuditEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceAuditEventsCompositePrimaryKeyWhere([]model.FinanceAuditEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeAuditEventsQueries.selectCountFinanceAuditEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceAuditEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceAuditEvents(ctx context.Context, selectFields ...FinanceAuditEventsField) (financeAuditEventsList model.FinanceAuditEventsList, err error) {
	var (
		defaultFinanceAuditEventsSelectFields = defaultFinanceAuditEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceAuditEventsSelectFields = composeFinanceAuditEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeAuditEventsQueries.selectFinanceAuditEvents, defaultFinanceAuditEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeAuditEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceAuditEvents] failed get financeAuditEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceAuditEventsByID(ctx context.Context, primaryID model.FinanceAuditEventsPrimaryID, selectFields ...FinanceAuditEventsField) (financeAuditEvents model.FinanceAuditEvents, err error) {
	var (
		defaultFinanceAuditEventsSelectFields = defaultFinanceAuditEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceAuditEventsSelectFields = composeFinanceAuditEventsSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceAuditEventsCompositePrimaryKeyWhere([]model.FinanceAuditEventsPrimaryID{primaryID})
	query := fmt.Sprintf(financeAuditEventsQueries.selectFinanceAuditEvents+" WHERE "+whereQry, defaultFinanceAuditEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeAuditEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeAuditEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceAuditEventsByID] failed get financeAuditEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceAuditEventsByID(ctx context.Context, primaryID model.FinanceAuditEventsPrimaryID, financeAuditEvents *model.FinanceAuditEvents, financeAuditEventsUpdateFields ...FinanceAuditEventsUpdateField) (err error) {
	exists, err := repo.IsExistFinanceAuditEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceAuditEvents] failed checking financeAuditEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeAuditEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeAuditEvents == nil {
		if len(financeAuditEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceAuditEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeAuditEvents = &model.FinanceAuditEvents{}
	}
	var (
		defaultFinanceAuditEventsUpdateFields = defaultFinanceAuditEventsUpdateFields(*financeAuditEvents)
		tempUpdateField                       FinanceAuditEventsUpdateFieldList
		selectFields                          = NewFinanceAuditEventsSelectFields()
	)
	if len(financeAuditEventsUpdateFields) > 0 {
		for _, updateField := range financeAuditEventsUpdateFields {
			if updateField.financeAuditEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceAuditEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceAuditEventsCompositePrimaryKeyWhere([]model.FinanceAuditEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceAuditEventsCommand(defaultFinanceAuditEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeAuditEventsQueries.updateFinanceAuditEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceAuditEvents] error when try to update financeAuditEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceAuditEventsByFilter(ctx context.Context, filter model.Filter, financeAuditEventsUpdateFields ...FinanceAuditEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeAuditEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceAuditEventsUpdateFieldList
		selectFields = NewFinanceAuditEventsSelectFields()
	)
	for _, updateField := range financeAuditEventsUpdateFields {
		if updateField.financeAuditEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceAuditEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeAuditEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceAuditEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_audit_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceAuditEventsByFilter] error when try to update financeAuditEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceAuditEventsByFilter] failed get rows affected")
	}
	return
}

var (
	financeAuditEventsQueries = struct {
		selectFinanceAuditEvents      string
		selectCountFinanceAuditEvents string
		deleteFinanceAuditEvents      string
		updateFinanceAuditEvents      string
		insertFinanceAuditEvents      string
	}{
		selectFinanceAuditEvents:      "SELECT %s FROM \"finance_audit_events\"",
		selectCountFinanceAuditEvents: "SELECT COUNT(\"id\") FROM \"finance_audit_events\"",
		deleteFinanceAuditEvents:      "DELETE FROM \"finance_audit_events\"",
		updateFinanceAuditEvents:      "UPDATE \"finance_audit_events\" SET %s ",
		insertFinanceAuditEvents:      "INSERT INTO \"finance_audit_events\" %s VALUES %s",
	}
)

type FinanceAuditEventsRepository interface {
	CreateFinanceAuditEvents(ctx context.Context, financeAuditEvents *model.FinanceAuditEvents, fieldsInsert ...FinanceAuditEventsField) error
	BulkCreateFinanceAuditEvents(ctx context.Context, financeAuditEventsList []*model.FinanceAuditEvents, fieldsInsert ...FinanceAuditEventsField) error
	ResolveFinanceAuditEvents(ctx context.Context, selectFields ...FinanceAuditEventsField) (model.FinanceAuditEventsList, error)
	ResolveFinanceAuditEventsByID(ctx context.Context, primaryID model.FinanceAuditEventsPrimaryID, selectFields ...FinanceAuditEventsField) (model.FinanceAuditEvents, error)
	UpdateFinanceAuditEventsByID(ctx context.Context, id model.FinanceAuditEventsPrimaryID, financeAuditEvents *model.FinanceAuditEvents, financeAuditEventsUpdateFields ...FinanceAuditEventsUpdateField) error
	UpdateFinanceAuditEventsByFilter(ctx context.Context, filter model.Filter, financeAuditEventsUpdateFields ...FinanceAuditEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceAuditEvents(ctx context.Context, financeAuditEventsListMap map[model.FinanceAuditEventsPrimaryID]*model.FinanceAuditEvents, FinanceAuditEventssMapUpdateFieldsRequest map[model.FinanceAuditEventsPrimaryID]FinanceAuditEventsUpdateFieldList) (err error)
	DeleteFinanceAuditEventsByID(ctx context.Context, id model.FinanceAuditEventsPrimaryID) error
	BulkDeleteFinanceAuditEventsByIDs(ctx context.Context, ids []model.FinanceAuditEventsPrimaryID) error
	ResolveFinanceAuditEventsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceAuditEventsFilterResult, err error)
	IsExistFinanceAuditEventsByIDs(ctx context.Context, ids []model.FinanceAuditEventsPrimaryID) (exists bool, notFoundIds []model.FinanceAuditEventsPrimaryID, err error)
	IsExistFinanceAuditEventsByID(ctx context.Context, id model.FinanceAuditEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
