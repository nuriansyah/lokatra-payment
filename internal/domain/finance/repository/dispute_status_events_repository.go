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

func composeInsertFieldsAndParamsDisputeStatusEvents(disputeStatusEventsList []model.DisputeStatusEvents, fieldsInsert ...DisputeStatusEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewDisputeStatusEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, disputeStatusEvents := range disputeStatusEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, disputeStatusEvents.Id)
			case selectField.DisputeId():
				args = append(args, disputeStatusEvents.DisputeId)
			case selectField.PreviousStatus():
				args = append(args, disputeStatusEvents.PreviousStatus)
			case selectField.NextStatus():
				args = append(args, disputeStatusEvents.NextStatus)
			case selectField.ReasonCode():
				args = append(args, disputeStatusEvents.ReasonCode)
			case selectField.ActorId():
				args = append(args, disputeStatusEvents.ActorId)
			case selectField.OccurredAt():
				args = append(args, disputeStatusEvents.OccurredAt)
			case selectField.Metadata():
				args = append(args, disputeStatusEvents.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, disputeStatusEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, disputeStatusEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, disputeStatusEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, disputeStatusEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, disputeStatusEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, disputeStatusEvents.MetaDeletedBy)

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

func composeDisputeStatusEventsCompositePrimaryKeyWhere(primaryIDs []model.DisputeStatusEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"dispute_status_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultDisputeStatusEventsSelectFields() string {
	fields := NewDisputeStatusEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeDisputeStatusEventsSelectFields(selectFields ...DisputeStatusEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type DisputeStatusEventsField string
type DisputeStatusEventsFieldList []DisputeStatusEventsField

type DisputeStatusEventsSelectFields struct {
}

func (ss DisputeStatusEventsSelectFields) Id() DisputeStatusEventsField {
	return DisputeStatusEventsField("id")
}

func (ss DisputeStatusEventsSelectFields) DisputeId() DisputeStatusEventsField {
	return DisputeStatusEventsField("dispute_id")
}

func (ss DisputeStatusEventsSelectFields) PreviousStatus() DisputeStatusEventsField {
	return DisputeStatusEventsField("previous_status")
}

func (ss DisputeStatusEventsSelectFields) NextStatus() DisputeStatusEventsField {
	return DisputeStatusEventsField("next_status")
}

func (ss DisputeStatusEventsSelectFields) ReasonCode() DisputeStatusEventsField {
	return DisputeStatusEventsField("reason_code")
}

func (ss DisputeStatusEventsSelectFields) ActorId() DisputeStatusEventsField {
	return DisputeStatusEventsField("actor_id")
}

func (ss DisputeStatusEventsSelectFields) OccurredAt() DisputeStatusEventsField {
	return DisputeStatusEventsField("occurred_at")
}

func (ss DisputeStatusEventsSelectFields) Metadata() DisputeStatusEventsField {
	return DisputeStatusEventsField("metadata")
}

func (ss DisputeStatusEventsSelectFields) MetaCreatedAt() DisputeStatusEventsField {
	return DisputeStatusEventsField("meta_created_at")
}

func (ss DisputeStatusEventsSelectFields) MetaCreatedBy() DisputeStatusEventsField {
	return DisputeStatusEventsField("meta_created_by")
}

func (ss DisputeStatusEventsSelectFields) MetaUpdatedAt() DisputeStatusEventsField {
	return DisputeStatusEventsField("meta_updated_at")
}

func (ss DisputeStatusEventsSelectFields) MetaUpdatedBy() DisputeStatusEventsField {
	return DisputeStatusEventsField("meta_updated_by")
}

func (ss DisputeStatusEventsSelectFields) MetaDeletedAt() DisputeStatusEventsField {
	return DisputeStatusEventsField("meta_deleted_at")
}

func (ss DisputeStatusEventsSelectFields) MetaDeletedBy() DisputeStatusEventsField {
	return DisputeStatusEventsField("meta_deleted_by")
}

func (ss DisputeStatusEventsSelectFields) All() DisputeStatusEventsFieldList {
	return []DisputeStatusEventsField{
		ss.Id(),
		ss.DisputeId(),
		ss.PreviousStatus(),
		ss.NextStatus(),
		ss.ReasonCode(),
		ss.ActorId(),
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

func NewDisputeStatusEventsSelectFields() DisputeStatusEventsSelectFields {
	return DisputeStatusEventsSelectFields{}
}

type DisputeStatusEventsUpdateFieldOption struct {
	useIncrement bool
}
type DisputeStatusEventsUpdateField struct {
	disputeStatusEventsField DisputeStatusEventsField
	opt                      DisputeStatusEventsUpdateFieldOption
	value                    interface{}
}
type DisputeStatusEventsUpdateFieldList []DisputeStatusEventsUpdateField

func defaultDisputeStatusEventsUpdateFieldOption() DisputeStatusEventsUpdateFieldOption {
	return DisputeStatusEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementDisputeStatusEventsOption(useIncrement bool) func(*DisputeStatusEventsUpdateFieldOption) {
	return func(pcufo *DisputeStatusEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewDisputeStatusEventsUpdateField(field DisputeStatusEventsField, val interface{}, opts ...func(*DisputeStatusEventsUpdateFieldOption)) DisputeStatusEventsUpdateField {
	defaultOpt := defaultDisputeStatusEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return DisputeStatusEventsUpdateField{
		disputeStatusEventsField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultDisputeStatusEventsUpdateFields(disputeStatusEvents model.DisputeStatusEvents) (disputeStatusEventsUpdateFieldList DisputeStatusEventsUpdateFieldList) {
	selectFields := NewDisputeStatusEventsSelectFields()
	disputeStatusEventsUpdateFieldList = append(disputeStatusEventsUpdateFieldList,
		NewDisputeStatusEventsUpdateField(selectFields.Id(), disputeStatusEvents.Id),
		NewDisputeStatusEventsUpdateField(selectFields.DisputeId(), disputeStatusEvents.DisputeId),
		NewDisputeStatusEventsUpdateField(selectFields.PreviousStatus(), disputeStatusEvents.PreviousStatus),
		NewDisputeStatusEventsUpdateField(selectFields.NextStatus(), disputeStatusEvents.NextStatus),
		NewDisputeStatusEventsUpdateField(selectFields.ReasonCode(), disputeStatusEvents.ReasonCode),
		NewDisputeStatusEventsUpdateField(selectFields.ActorId(), disputeStatusEvents.ActorId),
		NewDisputeStatusEventsUpdateField(selectFields.OccurredAt(), disputeStatusEvents.OccurredAt),
		NewDisputeStatusEventsUpdateField(selectFields.Metadata(), disputeStatusEvents.Metadata),
		NewDisputeStatusEventsUpdateField(selectFields.MetaCreatedAt(), disputeStatusEvents.MetaCreatedAt),
		NewDisputeStatusEventsUpdateField(selectFields.MetaCreatedBy(), disputeStatusEvents.MetaCreatedBy),
		NewDisputeStatusEventsUpdateField(selectFields.MetaUpdatedAt(), disputeStatusEvents.MetaUpdatedAt),
		NewDisputeStatusEventsUpdateField(selectFields.MetaUpdatedBy(), disputeStatusEvents.MetaUpdatedBy),
		NewDisputeStatusEventsUpdateField(selectFields.MetaDeletedAt(), disputeStatusEvents.MetaDeletedAt),
		NewDisputeStatusEventsUpdateField(selectFields.MetaDeletedBy(), disputeStatusEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsDisputeStatusEventsCommand(disputeStatusEventsUpdateFieldList DisputeStatusEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range disputeStatusEventsUpdateFieldList {
		field := string(updateField.disputeStatusEventsField)
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

func (repo *RepositoryImpl) BulkCreateDisputeStatusEvents(ctx context.Context, disputeStatusEventsList []*model.DisputeStatusEvents, fieldsInsert ...DisputeStatusEventsField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.DisputeStatusEventsPrimaryID
		disputeStatusEventsValueList []model.DisputeStatusEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewDisputeStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, disputeStatusEvents := range disputeStatusEventsList {

		primaryIds = append(primaryIds, disputeStatusEvents.ToDisputeStatusEventsPrimaryID())

		disputeStatusEventsValueList = append(disputeStatusEventsValueList, *disputeStatusEvents)
	}

	_, notFoundIds, err := repo.IsExistDisputeStatusEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputeStatusEvents] failed checking disputeStatusEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.DisputeStatusEventsPrimaryID{}
		mapNotFoundIds := map[model.DisputeStatusEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "disputeStatusEvents", fmt.Sprintf("disputeStatusEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsDisputeStatusEvents(disputeStatusEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(disputeStatusEventsQueries.insertDisputeStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputeStatusEvents] failed exec create disputeStatusEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteDisputeStatusEventsByIDs(ctx context.Context, primaryIDs []model.DisputeStatusEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistDisputeStatusEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeStatusEventsByIDs] failed checking disputeStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeStatusEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"dispute_status_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(disputeStatusEventsQueries.deleteDisputeStatusEvents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeStatusEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeStatusEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistDisputeStatusEventsByIDs(ctx context.Context, ids []model.DisputeStatusEventsPrimaryID) (exists bool, notFoundIds []model.DisputeStatusEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"dispute_status_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(disputeStatusEventsQueries.selectDisputeStatusEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeStatusEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.DisputeStatusEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeStatusEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.DisputeStatusEventsPrimaryID]bool{}
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

// BulkUpdateDisputeStatusEvents is used to bulk update disputeStatusEvents, by default it will update all field
// if want to update specific field, then fill disputeStatusEventssMapUpdateFieldsRequest else please fill disputeStatusEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateDisputeStatusEvents(ctx context.Context, disputeStatusEventssMap map[model.DisputeStatusEventsPrimaryID]*model.DisputeStatusEvents, disputeStatusEventssMapUpdateFieldsRequest map[model.DisputeStatusEventsPrimaryID]DisputeStatusEventsUpdateFieldList) (err error) {
	if len(disputeStatusEventssMap) == 0 && len(disputeStatusEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		disputeStatusEventssMapUpdateField map[model.DisputeStatusEventsPrimaryID]DisputeStatusEventsUpdateFieldList = map[model.DisputeStatusEventsPrimaryID]DisputeStatusEventsUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(disputeStatusEventssMap) > 0 {
		for id, disputeStatusEvents := range disputeStatusEventssMap {
			if disputeStatusEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdateDisputeStatusEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			disputeStatusEventssMapUpdateField[id] = defaultDisputeStatusEventsUpdateFields(*disputeStatusEvents)
		}
	} else {
		disputeStatusEventssMapUpdateField = disputeStatusEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateDisputeStatusEventsQuery(disputeStatusEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistDisputeStatusEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputeStatusEvents] failed checking disputeStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeStatusEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeDisputeStatusEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"dispute_status_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputeStatusEvents] failed exec query")
	}
	return
}

type DisputeStatusEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewDisputeStatusEventsFieldParameter(param string, args ...interface{}) DisputeStatusEventsFieldParameter {
	return DisputeStatusEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateDisputeStatusEventsQuery(mapDisputeStatusEventss map[model.DisputeStatusEventsPrimaryID]DisputeStatusEventsUpdateFieldList, asTableValues string) (primaryIDs []model.DisputeStatusEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.DisputeStatusEventsPrimaryID]map[string]interface{}{}
	disputeStatusEventsSelectFields := NewDisputeStatusEventsSelectFields()
	for id, updateFields := range mapDisputeStatusEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.disputeStatusEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapDisputeStatusEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetDisputeStatusEventsFieldType(updateField.disputeStatusEventsField)))
			args = append(args, fields[string(updateField.disputeStatusEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.disputeStatusEventsField))
		if updateField.disputeStatusEventsField == disputeStatusEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.disputeStatusEventsField, asTableValues, updateField.disputeStatusEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.disputeStatusEventsField,
				"\"dispute_status_events\"", updateField.disputeStatusEventsField,
				asTableValues, updateField.disputeStatusEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeDisputeStatusEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.DisputeStatusEventsPrimaryID, asTableValue string) (whereQry string) {
	disputeStatusEventsSelectFields := NewDisputeStatusEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"dispute_status_events\".\"id\" = %s.\"id\"::"+GetDisputeStatusEventsFieldType(disputeStatusEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetDisputeStatusEventsFieldType(disputeStatusEventsField DisputeStatusEventsField) string {
	selectDisputeStatusEventsFields := NewDisputeStatusEventsSelectFields()
	switch disputeStatusEventsField {

	case selectDisputeStatusEventsFields.Id():
		return "uuid"

	case selectDisputeStatusEventsFields.DisputeId():
		return "uuid"

	case selectDisputeStatusEventsFields.PreviousStatus():
		return "text"

	case selectDisputeStatusEventsFields.NextStatus():
		return "text"

	case selectDisputeStatusEventsFields.ReasonCode():
		return "text"

	case selectDisputeStatusEventsFields.ActorId():
		return "uuid"

	case selectDisputeStatusEventsFields.OccurredAt():
		return "timestamptz"

	case selectDisputeStatusEventsFields.Metadata():
		return "jsonb"

	case selectDisputeStatusEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectDisputeStatusEventsFields.MetaCreatedBy():
		return "uuid"

	case selectDisputeStatusEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectDisputeStatusEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectDisputeStatusEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectDisputeStatusEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateDisputeStatusEvents(ctx context.Context, disputeStatusEvents *model.DisputeStatusEvents, fieldsInsert ...DisputeStatusEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewDisputeStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.DisputeStatusEventsPrimaryID{
		Id: disputeStatusEvents.Id,
	}
	exists, err := repo.IsExistDisputeStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputeStatusEvents] failed checking disputeStatusEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "disputeStatusEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsDisputeStatusEvents([]model.DisputeStatusEvents{*disputeStatusEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(disputeStatusEventsQueries.insertDisputeStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputeStatusEvents] failed exec create disputeStatusEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteDisputeStatusEventsByID(ctx context.Context, primaryID model.DisputeStatusEventsPrimaryID) (err error) {
	exists, err := repo.IsExistDisputeStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputeStatusEventsByID] failed checking disputeStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeDisputeStatusEventsCompositePrimaryKeyWhere([]model.DisputeStatusEventsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(disputeStatusEventsQueries.deleteDisputeStatusEvents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputeStatusEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.DisputeStatusEventsFilterResult, err error) {
	query, args, err := composeDisputeStatusEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeStatusEventsByFilter] failed compose disputeStatusEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeStatusEventsByFilter] failed get disputeStatusEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeDisputeStatusEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.DisputeStatusEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeDisputeStatusEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeDisputeStatusEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeDisputeStatusEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewDisputeStatusEventsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 14+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["dispute_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"dispute_id\"")
			selectedColumns["dispute_id"] = struct{}{}
		}
		if _, selected := selectedColumns["previous_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"previous_status\"")
			selectedColumns["previous_status"] = struct{}{}
		}
		if _, selected := selectedColumns["next_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"next_status\"")
			selectedColumns["next_status"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["actor_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"actor_id\"")
			selectedColumns["actor_id"] = struct{}{}
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

type disputeStatusEventsFilterPlaceholder struct {
	index int
}

func (p *disputeStatusEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeDisputeStatusEventsFilterPredicate(filterField model.FilterField, placeholders *disputeStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewDisputeStatusEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeDisputeStatusEventsFilterSQLExpr(spec)
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

func composeDisputeStatusEventsFilterGroup(group model.FilterGroup, placeholders *disputeStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeDisputeStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeDisputeStatusEventsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeDisputeStatusEventsFilterWhereQueries(filter model.Filter, placeholders *disputeStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeDisputeStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeDisputeStatusEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeDisputeStatusEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateDisputeStatusEventsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeDisputeStatusEventsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeDisputeStatusEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := disputeStatusEventsFilterPlaceholder{index: 1}
	whereQueries, err := composeDisputeStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewDisputeStatusEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeDisputeStatusEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeDisputeStatusEventsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"dispute_status_events\" base%s", strings.Join(selectColumns, ","), composeDisputeStatusEventsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistDisputeStatusEventsByID(ctx context.Context, primaryID model.DisputeStatusEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeDisputeStatusEventsCompositePrimaryKeyWhere([]model.DisputeStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", disputeStatusEventsQueries.selectCountDisputeStatusEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeStatusEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeStatusEvents(ctx context.Context, selectFields ...DisputeStatusEventsField) (disputeStatusEventsList model.DisputeStatusEventsList, err error) {
	var (
		defaultDisputeStatusEventsSelectFields = defaultDisputeStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputeStatusEventsSelectFields = composeDisputeStatusEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(disputeStatusEventsQueries.selectDisputeStatusEvents, defaultDisputeStatusEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &disputeStatusEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeStatusEvents] failed get disputeStatusEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeStatusEventsByID(ctx context.Context, primaryID model.DisputeStatusEventsPrimaryID, selectFields ...DisputeStatusEventsField) (disputeStatusEvents model.DisputeStatusEvents, err error) {
	var (
		defaultDisputeStatusEventsSelectFields = defaultDisputeStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputeStatusEventsSelectFields = composeDisputeStatusEventsSelectFields(selectFields...)
	}
	whereQry, params := composeDisputeStatusEventsCompositePrimaryKeyWhere([]model.DisputeStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf(disputeStatusEventsQueries.selectDisputeStatusEvents+" WHERE "+whereQry, defaultDisputeStatusEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &disputeStatusEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("disputeStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveDisputeStatusEventsByID] failed get disputeStatusEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateDisputeStatusEventsByID(ctx context.Context, primaryID model.DisputeStatusEventsPrimaryID, disputeStatusEvents *model.DisputeStatusEvents, disputeStatusEventsUpdateFields ...DisputeStatusEventsUpdateField) (err error) {
	exists, err := repo.IsExistDisputeStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeStatusEvents] failed checking disputeStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if disputeStatusEvents == nil {
		if len(disputeStatusEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateDisputeStatusEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		disputeStatusEvents = &model.DisputeStatusEvents{}
	}
	var (
		defaultDisputeStatusEventsUpdateFields = defaultDisputeStatusEventsUpdateFields(*disputeStatusEvents)
		tempUpdateField                        DisputeStatusEventsUpdateFieldList
		selectFields                           = NewDisputeStatusEventsSelectFields()
	)
	if len(disputeStatusEventsUpdateFields) > 0 {
		for _, updateField := range disputeStatusEventsUpdateFields {
			if updateField.disputeStatusEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultDisputeStatusEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeDisputeStatusEventsCompositePrimaryKeyWhere([]model.DisputeStatusEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsDisputeStatusEventsCommand(defaultDisputeStatusEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(disputeStatusEventsQueries.updateDisputeStatusEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeStatusEvents] error when try to update disputeStatusEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateDisputeStatusEventsByFilter(ctx context.Context, filter model.Filter, disputeStatusEventsUpdateFields ...DisputeStatusEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(disputeStatusEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields DisputeStatusEventsUpdateFieldList
		selectFields = NewDisputeStatusEventsSelectFields()
	)
	for _, updateField := range disputeStatusEventsUpdateFields {
		if updateField.disputeStatusEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsDisputeStatusEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := disputeStatusEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeDisputeStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"dispute_status_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeStatusEventsByFilter] error when try to update disputeStatusEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeStatusEventsByFilter] failed get rows affected")
	}
	return
}

var (
	disputeStatusEventsQueries = struct {
		selectDisputeStatusEvents      string
		selectCountDisputeStatusEvents string
		deleteDisputeStatusEvents      string
		updateDisputeStatusEvents      string
		insertDisputeStatusEvents      string
	}{
		selectDisputeStatusEvents:      "SELECT %s FROM \"dispute_status_events\"",
		selectCountDisputeStatusEvents: "SELECT COUNT(\"id\") FROM \"dispute_status_events\"",
		deleteDisputeStatusEvents:      "DELETE FROM \"dispute_status_events\"",
		updateDisputeStatusEvents:      "UPDATE \"dispute_status_events\" SET %s ",
		insertDisputeStatusEvents:      "INSERT INTO \"dispute_status_events\" %s VALUES %s",
	}
)

type DisputeStatusEventsRepository interface {
	CreateDisputeStatusEvents(ctx context.Context, disputeStatusEvents *model.DisputeStatusEvents, fieldsInsert ...DisputeStatusEventsField) error
	BulkCreateDisputeStatusEvents(ctx context.Context, disputeStatusEventsList []*model.DisputeStatusEvents, fieldsInsert ...DisputeStatusEventsField) error
	ResolveDisputeStatusEvents(ctx context.Context, selectFields ...DisputeStatusEventsField) (model.DisputeStatusEventsList, error)
	ResolveDisputeStatusEventsByID(ctx context.Context, primaryID model.DisputeStatusEventsPrimaryID, selectFields ...DisputeStatusEventsField) (model.DisputeStatusEvents, error)
	UpdateDisputeStatusEventsByID(ctx context.Context, id model.DisputeStatusEventsPrimaryID, disputeStatusEvents *model.DisputeStatusEvents, disputeStatusEventsUpdateFields ...DisputeStatusEventsUpdateField) error
	UpdateDisputeStatusEventsByFilter(ctx context.Context, filter model.Filter, disputeStatusEventsUpdateFields ...DisputeStatusEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdateDisputeStatusEvents(ctx context.Context, disputeStatusEventsListMap map[model.DisputeStatusEventsPrimaryID]*model.DisputeStatusEvents, DisputeStatusEventssMapUpdateFieldsRequest map[model.DisputeStatusEventsPrimaryID]DisputeStatusEventsUpdateFieldList) (err error)
	DeleteDisputeStatusEventsByID(ctx context.Context, id model.DisputeStatusEventsPrimaryID) error
	BulkDeleteDisputeStatusEventsByIDs(ctx context.Context, ids []model.DisputeStatusEventsPrimaryID) error
	ResolveDisputeStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.DisputeStatusEventsFilterResult, err error)
	IsExistDisputeStatusEventsByIDs(ctx context.Context, ids []model.DisputeStatusEventsPrimaryID) (exists bool, notFoundIds []model.DisputeStatusEventsPrimaryID, err error)
	IsExistDisputeStatusEventsByID(ctx context.Context, id model.DisputeStatusEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
