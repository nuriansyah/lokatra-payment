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

func composeInsertFieldsAndParamsRefundStatusEvents(refundStatusEventsList []model.RefundStatusEvents, fieldsInsert ...RefundStatusEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRefundStatusEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, refundStatusEvents := range refundStatusEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, refundStatusEvents.Id)
			case selectField.RefundId():
				args = append(args, refundStatusEvents.RefundId)
			case selectField.FromStatus():
				args = append(args, refundStatusEvents.FromStatus)
			case selectField.ToStatus():
				args = append(args, refundStatusEvents.ToStatus)
			case selectField.ReasonCode():
				args = append(args, refundStatusEvents.ReasonCode)
			case selectField.ProviderRef():
				args = append(args, refundStatusEvents.ProviderRef)
			case selectField.EventPayload():
				args = append(args, refundStatusEvents.EventPayload)
			case selectField.OccurredAt():
				args = append(args, refundStatusEvents.OccurredAt)
			case selectField.MetaCreatedAt():
				args = append(args, refundStatusEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, refundStatusEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, refundStatusEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, refundStatusEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, refundStatusEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, refundStatusEvents.MetaDeletedBy)

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

func composeRefundStatusEventsCompositePrimaryKeyWhere(primaryIDs []model.RefundStatusEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"refund_status_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRefundStatusEventsSelectFields() string {
	fields := NewRefundStatusEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRefundStatusEventsSelectFields(selectFields ...RefundStatusEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RefundStatusEventsField string
type RefundStatusEventsFieldList []RefundStatusEventsField

type RefundStatusEventsSelectFields struct {
}

func (ss RefundStatusEventsSelectFields) Id() RefundStatusEventsField {
	return RefundStatusEventsField("id")
}

func (ss RefundStatusEventsSelectFields) RefundId() RefundStatusEventsField {
	return RefundStatusEventsField("refund_id")
}

func (ss RefundStatusEventsSelectFields) FromStatus() RefundStatusEventsField {
	return RefundStatusEventsField("from_status")
}

func (ss RefundStatusEventsSelectFields) ToStatus() RefundStatusEventsField {
	return RefundStatusEventsField("to_status")
}

func (ss RefundStatusEventsSelectFields) ReasonCode() RefundStatusEventsField {
	return RefundStatusEventsField("reason_code")
}

func (ss RefundStatusEventsSelectFields) ProviderRef() RefundStatusEventsField {
	return RefundStatusEventsField("provider_ref")
}

func (ss RefundStatusEventsSelectFields) EventPayload() RefundStatusEventsField {
	return RefundStatusEventsField("event_payload")
}

func (ss RefundStatusEventsSelectFields) OccurredAt() RefundStatusEventsField {
	return RefundStatusEventsField("occurred_at")
}

func (ss RefundStatusEventsSelectFields) MetaCreatedAt() RefundStatusEventsField {
	return RefundStatusEventsField("meta_created_at")
}

func (ss RefundStatusEventsSelectFields) MetaCreatedBy() RefundStatusEventsField {
	return RefundStatusEventsField("meta_created_by")
}

func (ss RefundStatusEventsSelectFields) MetaUpdatedAt() RefundStatusEventsField {
	return RefundStatusEventsField("meta_updated_at")
}

func (ss RefundStatusEventsSelectFields) MetaUpdatedBy() RefundStatusEventsField {
	return RefundStatusEventsField("meta_updated_by")
}

func (ss RefundStatusEventsSelectFields) MetaDeletedAt() RefundStatusEventsField {
	return RefundStatusEventsField("meta_deleted_at")
}

func (ss RefundStatusEventsSelectFields) MetaDeletedBy() RefundStatusEventsField {
	return RefundStatusEventsField("meta_deleted_by")
}

func (ss RefundStatusEventsSelectFields) All() RefundStatusEventsFieldList {
	return []RefundStatusEventsField{
		ss.Id(),
		ss.RefundId(),
		ss.FromStatus(),
		ss.ToStatus(),
		ss.ReasonCode(),
		ss.ProviderRef(),
		ss.EventPayload(),
		ss.OccurredAt(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRefundStatusEventsSelectFields() RefundStatusEventsSelectFields {
	return RefundStatusEventsSelectFields{}
}

type RefundStatusEventsUpdateFieldOption struct {
	useIncrement bool
}
type RefundStatusEventsUpdateField struct {
	refundStatusEventsField RefundStatusEventsField
	opt                     RefundStatusEventsUpdateFieldOption
	value                   interface{}
}
type RefundStatusEventsUpdateFieldList []RefundStatusEventsUpdateField

func defaultRefundStatusEventsUpdateFieldOption() RefundStatusEventsUpdateFieldOption {
	return RefundStatusEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRefundStatusEventsOption(useIncrement bool) func(*RefundStatusEventsUpdateFieldOption) {
	return func(pcufo *RefundStatusEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRefundStatusEventsUpdateField(field RefundStatusEventsField, val interface{}, opts ...func(*RefundStatusEventsUpdateFieldOption)) RefundStatusEventsUpdateField {
	defaultOpt := defaultRefundStatusEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RefundStatusEventsUpdateField{
		refundStatusEventsField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultRefundStatusEventsUpdateFields(refundStatusEvents model.RefundStatusEvents) (refundStatusEventsUpdateFieldList RefundStatusEventsUpdateFieldList) {
	selectFields := NewRefundStatusEventsSelectFields()
	refundStatusEventsUpdateFieldList = append(refundStatusEventsUpdateFieldList,
		NewRefundStatusEventsUpdateField(selectFields.Id(), refundStatusEvents.Id),
		NewRefundStatusEventsUpdateField(selectFields.RefundId(), refundStatusEvents.RefundId),
		NewRefundStatusEventsUpdateField(selectFields.FromStatus(), refundStatusEvents.FromStatus),
		NewRefundStatusEventsUpdateField(selectFields.ToStatus(), refundStatusEvents.ToStatus),
		NewRefundStatusEventsUpdateField(selectFields.ReasonCode(), refundStatusEvents.ReasonCode),
		NewRefundStatusEventsUpdateField(selectFields.ProviderRef(), refundStatusEvents.ProviderRef),
		NewRefundStatusEventsUpdateField(selectFields.EventPayload(), refundStatusEvents.EventPayload),
		NewRefundStatusEventsUpdateField(selectFields.OccurredAt(), refundStatusEvents.OccurredAt),
		NewRefundStatusEventsUpdateField(selectFields.MetaCreatedAt(), refundStatusEvents.MetaCreatedAt),
		NewRefundStatusEventsUpdateField(selectFields.MetaCreatedBy(), refundStatusEvents.MetaCreatedBy),
		NewRefundStatusEventsUpdateField(selectFields.MetaUpdatedAt(), refundStatusEvents.MetaUpdatedAt),
		NewRefundStatusEventsUpdateField(selectFields.MetaUpdatedBy(), refundStatusEvents.MetaUpdatedBy),
		NewRefundStatusEventsUpdateField(selectFields.MetaDeletedAt(), refundStatusEvents.MetaDeletedAt),
		NewRefundStatusEventsUpdateField(selectFields.MetaDeletedBy(), refundStatusEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRefundStatusEventsCommand(refundStatusEventsUpdateFieldList RefundStatusEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range refundStatusEventsUpdateFieldList {
		field := string(updateField.refundStatusEventsField)
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

func (repo *RepositoryImpl) BulkCreateRefundStatusEvents(ctx context.Context, refundStatusEventsList []*model.RefundStatusEvents, fieldsInsert ...RefundStatusEventsField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.RefundStatusEventsPrimaryID
		refundStatusEventsValueList []model.RefundStatusEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRefundStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, refundStatusEvents := range refundStatusEventsList {

		primaryIds = append(primaryIds, refundStatusEvents.ToRefundStatusEventsPrimaryID())

		refundStatusEventsValueList = append(refundStatusEventsValueList, *refundStatusEvents)
	}

	_, notFoundIds, err := repo.IsExistRefundStatusEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundStatusEvents] failed checking refundStatusEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RefundStatusEventsPrimaryID{}
		mapNotFoundIds := map[model.RefundStatusEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "refundStatusEvents", fmt.Sprintf("refundStatusEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRefundStatusEvents(refundStatusEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(refundStatusEventsQueries.insertRefundStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundStatusEvents] failed exec create refundStatusEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRefundStatusEventsByIDs(ctx context.Context, primaryIDs []model.RefundStatusEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRefundStatusEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundStatusEventsByIDs] failed checking refundStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundStatusEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_status_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(refundStatusEventsQueries.deleteRefundStatusEvents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundStatusEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundStatusEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRefundStatusEventsByIDs(ctx context.Context, ids []model.RefundStatusEventsPrimaryID) (exists bool, notFoundIds []model.RefundStatusEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_status_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(refundStatusEventsQueries.selectRefundStatusEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundStatusEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RefundStatusEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundStatusEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RefundStatusEventsPrimaryID]bool{}
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

// BulkUpdateRefundStatusEvents is used to bulk update refundStatusEvents, by default it will update all field
// if want to update specific field, then fill refundStatusEventssMapUpdateFieldsRequest else please fill refundStatusEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRefundStatusEvents(ctx context.Context, refundStatusEventssMap map[model.RefundStatusEventsPrimaryID]*model.RefundStatusEvents, refundStatusEventssMapUpdateFieldsRequest map[model.RefundStatusEventsPrimaryID]RefundStatusEventsUpdateFieldList) (err error) {
	if len(refundStatusEventssMap) == 0 && len(refundStatusEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		refundStatusEventssMapUpdateField map[model.RefundStatusEventsPrimaryID]RefundStatusEventsUpdateFieldList = map[model.RefundStatusEventsPrimaryID]RefundStatusEventsUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(refundStatusEventssMap) > 0 {
		for id, refundStatusEvents := range refundStatusEventssMap {
			if refundStatusEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdateRefundStatusEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			refundStatusEventssMapUpdateField[id] = defaultRefundStatusEventsUpdateFields(*refundStatusEvents)
		}
	} else {
		refundStatusEventssMapUpdateField = refundStatusEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRefundStatusEventsQuery(refundStatusEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRefundStatusEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundStatusEvents] failed checking refundStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundStatusEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRefundStatusEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"refund_status_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundStatusEvents] failed exec query")
	}
	return
}

type RefundStatusEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewRefundStatusEventsFieldParameter(param string, args ...interface{}) RefundStatusEventsFieldParameter {
	return RefundStatusEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRefundStatusEventsQuery(mapRefundStatusEventss map[model.RefundStatusEventsPrimaryID]RefundStatusEventsUpdateFieldList, asTableValues string) (primaryIDs []model.RefundStatusEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RefundStatusEventsPrimaryID]map[string]interface{}{}
	refundStatusEventsSelectFields := NewRefundStatusEventsSelectFields()
	for id, updateFields := range mapRefundStatusEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.refundStatusEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRefundStatusEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRefundStatusEventsFieldType(updateField.refundStatusEventsField)))
			args = append(args, fields[string(updateField.refundStatusEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.refundStatusEventsField))
		if updateField.refundStatusEventsField == refundStatusEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.refundStatusEventsField, asTableValues, updateField.refundStatusEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.refundStatusEventsField,
				"\"refund_status_events\"", updateField.refundStatusEventsField,
				asTableValues, updateField.refundStatusEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRefundStatusEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RefundStatusEventsPrimaryID, asTableValue string) (whereQry string) {
	refundStatusEventsSelectFields := NewRefundStatusEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"refund_status_events\".\"id\" = %s.\"id\"::"+GetRefundStatusEventsFieldType(refundStatusEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRefundStatusEventsFieldType(refundStatusEventsField RefundStatusEventsField) string {
	selectRefundStatusEventsFields := NewRefundStatusEventsSelectFields()
	switch refundStatusEventsField {

	case selectRefundStatusEventsFields.Id():
		return "uuid"

	case selectRefundStatusEventsFields.RefundId():
		return "uuid"

	case selectRefundStatusEventsFields.FromStatus():
		return "text"

	case selectRefundStatusEventsFields.ToStatus():
		return "text"

	case selectRefundStatusEventsFields.ReasonCode():
		return "text"

	case selectRefundStatusEventsFields.ProviderRef():
		return "text"

	case selectRefundStatusEventsFields.EventPayload():
		return "jsonb"

	case selectRefundStatusEventsFields.OccurredAt():
		return "timestamptz"

	case selectRefundStatusEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectRefundStatusEventsFields.MetaCreatedBy():
		return "uuid"

	case selectRefundStatusEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRefundStatusEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectRefundStatusEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectRefundStatusEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRefundStatusEvents(ctx context.Context, refundStatusEvents *model.RefundStatusEvents, fieldsInsert ...RefundStatusEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRefundStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RefundStatusEventsPrimaryID{
		Id: refundStatusEvents.Id,
	}
	exists, err := repo.IsExistRefundStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundStatusEvents] failed checking refundStatusEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "refundStatusEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRefundStatusEvents([]model.RefundStatusEvents{*refundStatusEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(refundStatusEventsQueries.insertRefundStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundStatusEvents] failed exec create refundStatusEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRefundStatusEventsByID(ctx context.Context, primaryID model.RefundStatusEventsPrimaryID) (err error) {
	exists, err := repo.IsExistRefundStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundStatusEventsByID] failed checking refundStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRefundStatusEventsCompositePrimaryKeyWhere([]model.RefundStatusEventsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(refundStatusEventsQueries.deleteRefundStatusEvents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundStatusEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundStatusEventsFilterResult, err error) {
	query, args, err := composeRefundStatusEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundStatusEventsByFilter] failed compose refundStatusEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundStatusEventsByFilter] failed get refundStatusEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRefundStatusEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.RefundStatusEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeRefundStatusEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeRefundStatusEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeRefundStatusEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewRefundStatusEventsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["refund_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_id\"")
			selectedColumns["refund_id"] = struct{}{}
		}
		if _, selected := selectedColumns["from_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"from_status\"")
			selectedColumns["from_status"] = struct{}{}
		}
		if _, selected := selectedColumns["to_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"to_status\"")
			selectedColumns["to_status"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_ref\"")
			selectedColumns["provider_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["event_payload"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_payload\"")
			selectedColumns["event_payload"] = struct{}{}
		}
		if _, selected := selectedColumns["occurred_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"occurred_at\"")
			selectedColumns["occurred_at"] = struct{}{}
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

type refundStatusEventsFilterPlaceholder struct {
	index int
}

func (p *refundStatusEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeRefundStatusEventsFilterPredicate(filterField model.FilterField, placeholders *refundStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewRefundStatusEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeRefundStatusEventsFilterSQLExpr(spec)
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

func composeRefundStatusEventsFilterGroup(group model.FilterGroup, placeholders *refundStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeRefundStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeRefundStatusEventsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeRefundStatusEventsFilterWhereQueries(filter model.Filter, placeholders *refundStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeRefundStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeRefundStatusEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeRefundStatusEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRefundStatusEventsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeRefundStatusEventsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeRefundStatusEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := refundStatusEventsFilterPlaceholder{index: 1}
	whereQueries, err := composeRefundStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewRefundStatusEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeRefundStatusEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeRefundStatusEventsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"refund_status_events\" base%s", strings.Join(selectColumns, ","), composeRefundStatusEventsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistRefundStatusEventsByID(ctx context.Context, primaryID model.RefundStatusEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRefundStatusEventsCompositePrimaryKeyWhere([]model.RefundStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", refundStatusEventsQueries.selectCountRefundStatusEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundStatusEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundStatusEvents(ctx context.Context, selectFields ...RefundStatusEventsField) (refundStatusEventsList model.RefundStatusEventsList, err error) {
	var (
		defaultRefundStatusEventsSelectFields = defaultRefundStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundStatusEventsSelectFields = composeRefundStatusEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(refundStatusEventsQueries.selectRefundStatusEvents, defaultRefundStatusEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &refundStatusEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundStatusEvents] failed get refundStatusEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundStatusEventsByID(ctx context.Context, primaryID model.RefundStatusEventsPrimaryID, selectFields ...RefundStatusEventsField) (refundStatusEvents model.RefundStatusEvents, err error) {
	var (
		defaultRefundStatusEventsSelectFields = defaultRefundStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundStatusEventsSelectFields = composeRefundStatusEventsSelectFields(selectFields...)
	}
	whereQry, params := composeRefundStatusEventsCompositePrimaryKeyWhere([]model.RefundStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf(refundStatusEventsQueries.selectRefundStatusEvents+" WHERE "+whereQry, defaultRefundStatusEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &refundStatusEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("refundStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRefundStatusEventsByID] failed get refundStatusEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRefundStatusEventsByID(ctx context.Context, primaryID model.RefundStatusEventsPrimaryID, refundStatusEvents *model.RefundStatusEvents, refundStatusEventsUpdateFields ...RefundStatusEventsUpdateField) (err error) {
	exists, err := repo.IsExistRefundStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundStatusEvents] failed checking refundStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if refundStatusEvents == nil {
		if len(refundStatusEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRefundStatusEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		refundStatusEvents = &model.RefundStatusEvents{}
	}
	var (
		defaultRefundStatusEventsUpdateFields = defaultRefundStatusEventsUpdateFields(*refundStatusEvents)
		tempUpdateField                       RefundStatusEventsUpdateFieldList
		selectFields                          = NewRefundStatusEventsSelectFields()
	)
	if len(refundStatusEventsUpdateFields) > 0 {
		for _, updateField := range refundStatusEventsUpdateFields {
			if updateField.refundStatusEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRefundStatusEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRefundStatusEventsCompositePrimaryKeyWhere([]model.RefundStatusEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRefundStatusEventsCommand(defaultRefundStatusEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(refundStatusEventsQueries.updateRefundStatusEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundStatusEvents] error when try to update refundStatusEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateRefundStatusEventsByFilter(ctx context.Context, filter model.Filter, refundStatusEventsUpdateFields ...RefundStatusEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(refundStatusEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields RefundStatusEventsUpdateFieldList
		selectFields = NewRefundStatusEventsSelectFields()
	)
	for _, updateField := range refundStatusEventsUpdateFields {
		if updateField.refundStatusEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsRefundStatusEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := refundStatusEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeRefundStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"refund_status_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundStatusEventsByFilter] error when try to update refundStatusEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundStatusEventsByFilter] failed get rows affected")
	}
	return
}

var (
	refundStatusEventsQueries = struct {
		selectRefundStatusEvents      string
		selectCountRefundStatusEvents string
		deleteRefundStatusEvents      string
		updateRefundStatusEvents      string
		insertRefundStatusEvents      string
	}{
		selectRefundStatusEvents:      "SELECT %s FROM \"refund_status_events\"",
		selectCountRefundStatusEvents: "SELECT COUNT(\"id\") FROM \"refund_status_events\"",
		deleteRefundStatusEvents:      "DELETE FROM \"refund_status_events\"",
		updateRefundStatusEvents:      "UPDATE \"refund_status_events\" SET %s ",
		insertRefundStatusEvents:      "INSERT INTO \"refund_status_events\" %s VALUES %s",
	}
)

type RefundStatusEventsRepository interface {
	CreateRefundStatusEvents(ctx context.Context, refundStatusEvents *model.RefundStatusEvents, fieldsInsert ...RefundStatusEventsField) error
	BulkCreateRefundStatusEvents(ctx context.Context, refundStatusEventsList []*model.RefundStatusEvents, fieldsInsert ...RefundStatusEventsField) error
	ResolveRefundStatusEvents(ctx context.Context, selectFields ...RefundStatusEventsField) (model.RefundStatusEventsList, error)
	ResolveRefundStatusEventsByID(ctx context.Context, primaryID model.RefundStatusEventsPrimaryID, selectFields ...RefundStatusEventsField) (model.RefundStatusEvents, error)
	UpdateRefundStatusEventsByID(ctx context.Context, id model.RefundStatusEventsPrimaryID, refundStatusEvents *model.RefundStatusEvents, refundStatusEventsUpdateFields ...RefundStatusEventsUpdateField) error
	UpdateRefundStatusEventsByFilter(ctx context.Context, filter model.Filter, refundStatusEventsUpdateFields ...RefundStatusEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdateRefundStatusEvents(ctx context.Context, refundStatusEventsListMap map[model.RefundStatusEventsPrimaryID]*model.RefundStatusEvents, RefundStatusEventssMapUpdateFieldsRequest map[model.RefundStatusEventsPrimaryID]RefundStatusEventsUpdateFieldList) (err error)
	DeleteRefundStatusEventsByID(ctx context.Context, id model.RefundStatusEventsPrimaryID) error
	BulkDeleteRefundStatusEventsByIDs(ctx context.Context, ids []model.RefundStatusEventsPrimaryID) error
	ResolveRefundStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundStatusEventsFilterResult, err error)
	IsExistRefundStatusEventsByIDs(ctx context.Context, ids []model.RefundStatusEventsPrimaryID) (exists bool, notFoundIds []model.RefundStatusEventsPrimaryID, err error)
	IsExistRefundStatusEventsByID(ctx context.Context, id model.RefundStatusEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
