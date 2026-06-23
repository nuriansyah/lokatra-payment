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

func composeInsertFieldsAndParamsPayoutStatusEvents(payoutStatusEventsList []model.PayoutStatusEvents, fieldsInsert ...PayoutStatusEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPayoutStatusEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, payoutStatusEvents := range payoutStatusEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, payoutStatusEvents.Id)
			case selectField.PayoutId():
				args = append(args, payoutStatusEvents.PayoutId)
			case selectField.FromStatus():
				args = append(args, payoutStatusEvents.FromStatus)
			case selectField.ToStatus():
				args = append(args, payoutStatusEvents.ToStatus)
			case selectField.ReasonCode():
				args = append(args, payoutStatusEvents.ReasonCode)
			case selectField.ProviderRef():
				args = append(args, payoutStatusEvents.ProviderRef)
			case selectField.EventPayload():
				args = append(args, payoutStatusEvents.EventPayload)
			case selectField.OccurredAt():
				args = append(args, payoutStatusEvents.OccurredAt)
			case selectField.MetaCreatedAt():
				args = append(args, payoutStatusEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, payoutStatusEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, payoutStatusEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, payoutStatusEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, payoutStatusEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, payoutStatusEvents.MetaDeletedBy)

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

func composePayoutStatusEventsCompositePrimaryKeyWhere(primaryIDs []model.PayoutStatusEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payout_status_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPayoutStatusEventsSelectFields() string {
	fields := NewPayoutStatusEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePayoutStatusEventsSelectFields(selectFields ...PayoutStatusEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PayoutStatusEventsField string
type PayoutStatusEventsFieldList []PayoutStatusEventsField

type PayoutStatusEventsSelectFields struct {
}

func (ss PayoutStatusEventsSelectFields) Id() PayoutStatusEventsField {
	return PayoutStatusEventsField("id")
}

func (ss PayoutStatusEventsSelectFields) PayoutId() PayoutStatusEventsField {
	return PayoutStatusEventsField("payout_id")
}

func (ss PayoutStatusEventsSelectFields) FromStatus() PayoutStatusEventsField {
	return PayoutStatusEventsField("from_status")
}

func (ss PayoutStatusEventsSelectFields) ToStatus() PayoutStatusEventsField {
	return PayoutStatusEventsField("to_status")
}

func (ss PayoutStatusEventsSelectFields) ReasonCode() PayoutStatusEventsField {
	return PayoutStatusEventsField("reason_code")
}

func (ss PayoutStatusEventsSelectFields) ProviderRef() PayoutStatusEventsField {
	return PayoutStatusEventsField("provider_ref")
}

func (ss PayoutStatusEventsSelectFields) EventPayload() PayoutStatusEventsField {
	return PayoutStatusEventsField("event_payload")
}

func (ss PayoutStatusEventsSelectFields) OccurredAt() PayoutStatusEventsField {
	return PayoutStatusEventsField("occurred_at")
}

func (ss PayoutStatusEventsSelectFields) MetaCreatedAt() PayoutStatusEventsField {
	return PayoutStatusEventsField("meta_created_at")
}

func (ss PayoutStatusEventsSelectFields) MetaCreatedBy() PayoutStatusEventsField {
	return PayoutStatusEventsField("meta_created_by")
}

func (ss PayoutStatusEventsSelectFields) MetaUpdatedAt() PayoutStatusEventsField {
	return PayoutStatusEventsField("meta_updated_at")
}

func (ss PayoutStatusEventsSelectFields) MetaUpdatedBy() PayoutStatusEventsField {
	return PayoutStatusEventsField("meta_updated_by")
}

func (ss PayoutStatusEventsSelectFields) MetaDeletedAt() PayoutStatusEventsField {
	return PayoutStatusEventsField("meta_deleted_at")
}

func (ss PayoutStatusEventsSelectFields) MetaDeletedBy() PayoutStatusEventsField {
	return PayoutStatusEventsField("meta_deleted_by")
}

func (ss PayoutStatusEventsSelectFields) All() PayoutStatusEventsFieldList {
	return []PayoutStatusEventsField{
		ss.Id(),
		ss.PayoutId(),
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

func NewPayoutStatusEventsSelectFields() PayoutStatusEventsSelectFields {
	return PayoutStatusEventsSelectFields{}
}

type PayoutStatusEventsUpdateFieldOption struct {
	useIncrement bool
}
type PayoutStatusEventsUpdateField struct {
	payoutStatusEventsField PayoutStatusEventsField
	opt                     PayoutStatusEventsUpdateFieldOption
	value                   interface{}
}
type PayoutStatusEventsUpdateFieldList []PayoutStatusEventsUpdateField

func defaultPayoutStatusEventsUpdateFieldOption() PayoutStatusEventsUpdateFieldOption {
	return PayoutStatusEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPayoutStatusEventsOption(useIncrement bool) func(*PayoutStatusEventsUpdateFieldOption) {
	return func(pcufo *PayoutStatusEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPayoutStatusEventsUpdateField(field PayoutStatusEventsField, val interface{}, opts ...func(*PayoutStatusEventsUpdateFieldOption)) PayoutStatusEventsUpdateField {
	defaultOpt := defaultPayoutStatusEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PayoutStatusEventsUpdateField{
		payoutStatusEventsField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultPayoutStatusEventsUpdateFields(payoutStatusEvents model.PayoutStatusEvents) (payoutStatusEventsUpdateFieldList PayoutStatusEventsUpdateFieldList) {
	selectFields := NewPayoutStatusEventsSelectFields()
	payoutStatusEventsUpdateFieldList = append(payoutStatusEventsUpdateFieldList,
		NewPayoutStatusEventsUpdateField(selectFields.Id(), payoutStatusEvents.Id),
		NewPayoutStatusEventsUpdateField(selectFields.PayoutId(), payoutStatusEvents.PayoutId),
		NewPayoutStatusEventsUpdateField(selectFields.FromStatus(), payoutStatusEvents.FromStatus),
		NewPayoutStatusEventsUpdateField(selectFields.ToStatus(), payoutStatusEvents.ToStatus),
		NewPayoutStatusEventsUpdateField(selectFields.ReasonCode(), payoutStatusEvents.ReasonCode),
		NewPayoutStatusEventsUpdateField(selectFields.ProviderRef(), payoutStatusEvents.ProviderRef),
		NewPayoutStatusEventsUpdateField(selectFields.EventPayload(), payoutStatusEvents.EventPayload),
		NewPayoutStatusEventsUpdateField(selectFields.OccurredAt(), payoutStatusEvents.OccurredAt),
		NewPayoutStatusEventsUpdateField(selectFields.MetaCreatedAt(), payoutStatusEvents.MetaCreatedAt),
		NewPayoutStatusEventsUpdateField(selectFields.MetaCreatedBy(), payoutStatusEvents.MetaCreatedBy),
		NewPayoutStatusEventsUpdateField(selectFields.MetaUpdatedAt(), payoutStatusEvents.MetaUpdatedAt),
		NewPayoutStatusEventsUpdateField(selectFields.MetaUpdatedBy(), payoutStatusEvents.MetaUpdatedBy),
		NewPayoutStatusEventsUpdateField(selectFields.MetaDeletedAt(), payoutStatusEvents.MetaDeletedAt),
		NewPayoutStatusEventsUpdateField(selectFields.MetaDeletedBy(), payoutStatusEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPayoutStatusEventsCommand(payoutStatusEventsUpdateFieldList PayoutStatusEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range payoutStatusEventsUpdateFieldList {
		field := string(updateField.payoutStatusEventsField)
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

func (repo *RepositoryImpl) BulkCreatePayoutStatusEvents(ctx context.Context, payoutStatusEventsList []*model.PayoutStatusEvents, fieldsInsert ...PayoutStatusEventsField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.PayoutStatusEventsPrimaryID
		payoutStatusEventsValueList []model.PayoutStatusEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPayoutStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, payoutStatusEvents := range payoutStatusEventsList {

		primaryIds = append(primaryIds, payoutStatusEvents.ToPayoutStatusEventsPrimaryID())

		payoutStatusEventsValueList = append(payoutStatusEventsValueList, *payoutStatusEvents)
	}

	_, notFoundIds, err := repo.IsExistPayoutStatusEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutStatusEvents] failed checking payoutStatusEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PayoutStatusEventsPrimaryID{}
		mapNotFoundIds := map[model.PayoutStatusEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "payoutStatusEvents", fmt.Sprintf("payoutStatusEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPayoutStatusEvents(payoutStatusEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(payoutStatusEventsQueries.insertPayoutStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutStatusEvents] failed exec create payoutStatusEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePayoutStatusEventsByIDs(ctx context.Context, primaryIDs []model.PayoutStatusEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPayoutStatusEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutStatusEventsByIDs] failed checking payoutStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutStatusEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_status_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(payoutStatusEventsQueries.deletePayoutStatusEvents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutStatusEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutStatusEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPayoutStatusEventsByIDs(ctx context.Context, ids []model.PayoutStatusEventsPrimaryID) (exists bool, notFoundIds []model.PayoutStatusEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_status_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(payoutStatusEventsQueries.selectPayoutStatusEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutStatusEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PayoutStatusEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutStatusEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PayoutStatusEventsPrimaryID]bool{}
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

// BulkUpdatePayoutStatusEvents is used to bulk update payoutStatusEvents, by default it will update all field
// if want to update specific field, then fill payoutStatusEventssMapUpdateFieldsRequest else please fill payoutStatusEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePayoutStatusEvents(ctx context.Context, payoutStatusEventssMap map[model.PayoutStatusEventsPrimaryID]*model.PayoutStatusEvents, payoutStatusEventssMapUpdateFieldsRequest map[model.PayoutStatusEventsPrimaryID]PayoutStatusEventsUpdateFieldList) (err error) {
	if len(payoutStatusEventssMap) == 0 && len(payoutStatusEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		payoutStatusEventssMapUpdateField map[model.PayoutStatusEventsPrimaryID]PayoutStatusEventsUpdateFieldList = map[model.PayoutStatusEventsPrimaryID]PayoutStatusEventsUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(payoutStatusEventssMap) > 0 {
		for id, payoutStatusEvents := range payoutStatusEventssMap {
			if payoutStatusEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdatePayoutStatusEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			payoutStatusEventssMapUpdateField[id] = defaultPayoutStatusEventsUpdateFields(*payoutStatusEvents)
		}
	} else {
		payoutStatusEventssMapUpdateField = payoutStatusEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePayoutStatusEventsQuery(payoutStatusEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPayoutStatusEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutStatusEvents] failed checking payoutStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutStatusEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePayoutStatusEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payout_status_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutStatusEvents] failed exec query")
	}
	return
}

type PayoutStatusEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPayoutStatusEventsFieldParameter(param string, args ...interface{}) PayoutStatusEventsFieldParameter {
	return PayoutStatusEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePayoutStatusEventsQuery(mapPayoutStatusEventss map[model.PayoutStatusEventsPrimaryID]PayoutStatusEventsUpdateFieldList, asTableValues string) (primaryIDs []model.PayoutStatusEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PayoutStatusEventsPrimaryID]map[string]interface{}{}
	payoutStatusEventsSelectFields := NewPayoutStatusEventsSelectFields()
	for id, updateFields := range mapPayoutStatusEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.payoutStatusEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPayoutStatusEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPayoutStatusEventsFieldType(updateField.payoutStatusEventsField)))
			args = append(args, fields[string(updateField.payoutStatusEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.payoutStatusEventsField))
		if updateField.payoutStatusEventsField == payoutStatusEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.payoutStatusEventsField, asTableValues, updateField.payoutStatusEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.payoutStatusEventsField,
				"\"payout_status_events\"", updateField.payoutStatusEventsField,
				asTableValues, updateField.payoutStatusEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePayoutStatusEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PayoutStatusEventsPrimaryID, asTableValue string) (whereQry string) {
	payoutStatusEventsSelectFields := NewPayoutStatusEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payout_status_events\".\"id\" = %s.\"id\"::"+GetPayoutStatusEventsFieldType(payoutStatusEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPayoutStatusEventsFieldType(payoutStatusEventsField PayoutStatusEventsField) string {
	selectPayoutStatusEventsFields := NewPayoutStatusEventsSelectFields()
	switch payoutStatusEventsField {

	case selectPayoutStatusEventsFields.Id():
		return "uuid"

	case selectPayoutStatusEventsFields.PayoutId():
		return "uuid"

	case selectPayoutStatusEventsFields.FromStatus():
		return "text"

	case selectPayoutStatusEventsFields.ToStatus():
		return "text"

	case selectPayoutStatusEventsFields.ReasonCode():
		return "text"

	case selectPayoutStatusEventsFields.ProviderRef():
		return "text"

	case selectPayoutStatusEventsFields.EventPayload():
		return "jsonb"

	case selectPayoutStatusEventsFields.OccurredAt():
		return "timestamptz"

	case selectPayoutStatusEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPayoutStatusEventsFields.MetaCreatedBy():
		return "uuid"

	case selectPayoutStatusEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPayoutStatusEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectPayoutStatusEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPayoutStatusEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePayoutStatusEvents(ctx context.Context, payoutStatusEvents *model.PayoutStatusEvents, fieldsInsert ...PayoutStatusEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPayoutStatusEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PayoutStatusEventsPrimaryID{
		Id: payoutStatusEvents.Id,
	}
	exists, err := repo.IsExistPayoutStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutStatusEvents] failed checking payoutStatusEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "payoutStatusEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPayoutStatusEvents([]model.PayoutStatusEvents{*payoutStatusEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(payoutStatusEventsQueries.insertPayoutStatusEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutStatusEvents] failed exec create payoutStatusEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePayoutStatusEventsByID(ctx context.Context, primaryID model.PayoutStatusEventsPrimaryID) (err error) {
	exists, err := repo.IsExistPayoutStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutStatusEventsByID] failed checking payoutStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePayoutStatusEventsCompositePrimaryKeyWhere([]model.PayoutStatusEventsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(payoutStatusEventsQueries.deletePayoutStatusEvents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutStatusEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutStatusEventsFilterResult, err error) {
	query, args, err := composePayoutStatusEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutStatusEventsByFilter] failed compose payoutStatusEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutStatusEventsByFilter] failed get payoutStatusEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePayoutStatusEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PayoutStatusEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePayoutStatusEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePayoutStatusEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePayoutStatusEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPayoutStatusEventsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["payout_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_id\"")
			selectedColumns["payout_id"] = struct{}{}
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

type payoutStatusEventsFilterPlaceholder struct {
	index int
}

func (p *payoutStatusEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePayoutStatusEventsFilterPredicate(filterField model.FilterField, placeholders *payoutStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPayoutStatusEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePayoutStatusEventsFilterSQLExpr(spec)
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

func composePayoutStatusEventsFilterGroup(group model.FilterGroup, placeholders *payoutStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePayoutStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePayoutStatusEventsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePayoutStatusEventsFilterWhereQueries(filter model.Filter, placeholders *payoutStatusEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePayoutStatusEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePayoutStatusEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePayoutStatusEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePayoutStatusEventsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePayoutStatusEventsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePayoutStatusEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := payoutStatusEventsFilterPlaceholder{index: 1}
	whereQueries, err := composePayoutStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPayoutStatusEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePayoutStatusEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePayoutStatusEventsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payout_status_events\" base%s", strings.Join(selectColumns, ","), composePayoutStatusEventsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPayoutStatusEventsByID(ctx context.Context, primaryID model.PayoutStatusEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePayoutStatusEventsCompositePrimaryKeyWhere([]model.PayoutStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", payoutStatusEventsQueries.selectCountPayoutStatusEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutStatusEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutStatusEvents(ctx context.Context, selectFields ...PayoutStatusEventsField) (payoutStatusEventsList model.PayoutStatusEventsList, err error) {
	var (
		defaultPayoutStatusEventsSelectFields = defaultPayoutStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutStatusEventsSelectFields = composePayoutStatusEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(payoutStatusEventsQueries.selectPayoutStatusEvents, defaultPayoutStatusEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &payoutStatusEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutStatusEvents] failed get payoutStatusEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutStatusEventsByID(ctx context.Context, primaryID model.PayoutStatusEventsPrimaryID, selectFields ...PayoutStatusEventsField) (payoutStatusEvents model.PayoutStatusEvents, err error) {
	var (
		defaultPayoutStatusEventsSelectFields = defaultPayoutStatusEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutStatusEventsSelectFields = composePayoutStatusEventsSelectFields(selectFields...)
	}
	whereQry, params := composePayoutStatusEventsCompositePrimaryKeyWhere([]model.PayoutStatusEventsPrimaryID{primaryID})
	query := fmt.Sprintf(payoutStatusEventsQueries.selectPayoutStatusEvents+" WHERE "+whereQry, defaultPayoutStatusEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &payoutStatusEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("payoutStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePayoutStatusEventsByID] failed get payoutStatusEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePayoutStatusEventsByID(ctx context.Context, primaryID model.PayoutStatusEventsPrimaryID, payoutStatusEvents *model.PayoutStatusEvents, payoutStatusEventsUpdateFields ...PayoutStatusEventsUpdateField) (err error) {
	exists, err := repo.IsExistPayoutStatusEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutStatusEvents] failed checking payoutStatusEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutStatusEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if payoutStatusEvents == nil {
		if len(payoutStatusEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePayoutStatusEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		payoutStatusEvents = &model.PayoutStatusEvents{}
	}
	var (
		defaultPayoutStatusEventsUpdateFields = defaultPayoutStatusEventsUpdateFields(*payoutStatusEvents)
		tempUpdateField                       PayoutStatusEventsUpdateFieldList
		selectFields                          = NewPayoutStatusEventsSelectFields()
	)
	if len(payoutStatusEventsUpdateFields) > 0 {
		for _, updateField := range payoutStatusEventsUpdateFields {
			if updateField.payoutStatusEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPayoutStatusEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePayoutStatusEventsCompositePrimaryKeyWhere([]model.PayoutStatusEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPayoutStatusEventsCommand(defaultPayoutStatusEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(payoutStatusEventsQueries.updatePayoutStatusEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutStatusEvents] error when try to update payoutStatusEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePayoutStatusEventsByFilter(ctx context.Context, filter model.Filter, payoutStatusEventsUpdateFields ...PayoutStatusEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(payoutStatusEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PayoutStatusEventsUpdateFieldList
		selectFields = NewPayoutStatusEventsSelectFields()
	)
	for _, updateField := range payoutStatusEventsUpdateFields {
		if updateField.payoutStatusEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPayoutStatusEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := payoutStatusEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePayoutStatusEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payout_status_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutStatusEventsByFilter] error when try to update payoutStatusEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutStatusEventsByFilter] failed get rows affected")
	}
	return
}

var (
	payoutStatusEventsQueries = struct {
		selectPayoutStatusEvents      string
		selectCountPayoutStatusEvents string
		deletePayoutStatusEvents      string
		updatePayoutStatusEvents      string
		insertPayoutStatusEvents      string
	}{
		selectPayoutStatusEvents:      "SELECT %s FROM \"payout_status_events\"",
		selectCountPayoutStatusEvents: "SELECT COUNT(\"id\") FROM \"payout_status_events\"",
		deletePayoutStatusEvents:      "DELETE FROM \"payout_status_events\"",
		updatePayoutStatusEvents:      "UPDATE \"payout_status_events\" SET %s ",
		insertPayoutStatusEvents:      "INSERT INTO \"payout_status_events\" %s VALUES %s",
	}
)

type PayoutStatusEventsRepository interface {
	CreatePayoutStatusEvents(ctx context.Context, payoutStatusEvents *model.PayoutStatusEvents, fieldsInsert ...PayoutStatusEventsField) error
	BulkCreatePayoutStatusEvents(ctx context.Context, payoutStatusEventsList []*model.PayoutStatusEvents, fieldsInsert ...PayoutStatusEventsField) error
	ResolvePayoutStatusEvents(ctx context.Context, selectFields ...PayoutStatusEventsField) (model.PayoutStatusEventsList, error)
	ResolvePayoutStatusEventsByID(ctx context.Context, primaryID model.PayoutStatusEventsPrimaryID, selectFields ...PayoutStatusEventsField) (model.PayoutStatusEvents, error)
	UpdatePayoutStatusEventsByID(ctx context.Context, id model.PayoutStatusEventsPrimaryID, payoutStatusEvents *model.PayoutStatusEvents, payoutStatusEventsUpdateFields ...PayoutStatusEventsUpdateField) error
	UpdatePayoutStatusEventsByFilter(ctx context.Context, filter model.Filter, payoutStatusEventsUpdateFields ...PayoutStatusEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePayoutStatusEvents(ctx context.Context, payoutStatusEventsListMap map[model.PayoutStatusEventsPrimaryID]*model.PayoutStatusEvents, PayoutStatusEventssMapUpdateFieldsRequest map[model.PayoutStatusEventsPrimaryID]PayoutStatusEventsUpdateFieldList) (err error)
	DeletePayoutStatusEventsByID(ctx context.Context, id model.PayoutStatusEventsPrimaryID) error
	BulkDeletePayoutStatusEventsByIDs(ctx context.Context, ids []model.PayoutStatusEventsPrimaryID) error
	ResolvePayoutStatusEventsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutStatusEventsFilterResult, err error)
	IsExistPayoutStatusEventsByIDs(ctx context.Context, ids []model.PayoutStatusEventsPrimaryID) (exists bool, notFoundIds []model.PayoutStatusEventsPrimaryID, err error)
	IsExistPayoutStatusEventsByID(ctx context.Context, id model.PayoutStatusEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
