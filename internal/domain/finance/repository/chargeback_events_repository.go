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

func composeInsertFieldsAndParamsChargebackEvents(chargebackEventsList []model.ChargebackEvents, fieldsInsert ...ChargebackEventsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewChargebackEventsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, chargebackEvents := range chargebackEventsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, chargebackEvents.Id)
			case selectField.DisputeId():
				args = append(args, chargebackEvents.DisputeId)
			case selectField.EventType():
				args = append(args, chargebackEvents.EventType)
			case selectField.JournalEntryId():
				args = append(args, chargebackEvents.JournalEntryId)
			case selectField.Amount():
				args = append(args, chargebackEvents.Amount)
			case selectField.CurrencyCode():
				args = append(args, chargebackEvents.CurrencyCode)
			case selectField.ProviderEventRef():
				args = append(args, chargebackEvents.ProviderEventRef)
			case selectField.EventPayload():
				args = append(args, chargebackEvents.EventPayload)
			case selectField.OccurredAt():
				args = append(args, chargebackEvents.OccurredAt)
			case selectField.Metadata():
				args = append(args, chargebackEvents.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, chargebackEvents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, chargebackEvents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, chargebackEvents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, chargebackEvents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, chargebackEvents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, chargebackEvents.MetaDeletedBy)

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

func composeChargebackEventsCompositePrimaryKeyWhere(primaryIDs []model.ChargebackEventsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"chargeback_events\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultChargebackEventsSelectFields() string {
	fields := NewChargebackEventsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeChargebackEventsSelectFields(selectFields ...ChargebackEventsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ChargebackEventsField string
type ChargebackEventsFieldList []ChargebackEventsField

type ChargebackEventsSelectFields struct {
}

func (ss ChargebackEventsSelectFields) Id() ChargebackEventsField {
	return ChargebackEventsField("id")
}

func (ss ChargebackEventsSelectFields) DisputeId() ChargebackEventsField {
	return ChargebackEventsField("dispute_id")
}

func (ss ChargebackEventsSelectFields) EventType() ChargebackEventsField {
	return ChargebackEventsField("event_type")
}

func (ss ChargebackEventsSelectFields) JournalEntryId() ChargebackEventsField {
	return ChargebackEventsField("journal_entry_id")
}

func (ss ChargebackEventsSelectFields) Amount() ChargebackEventsField {
	return ChargebackEventsField("amount")
}

func (ss ChargebackEventsSelectFields) CurrencyCode() ChargebackEventsField {
	return ChargebackEventsField("currency_code")
}

func (ss ChargebackEventsSelectFields) ProviderEventRef() ChargebackEventsField {
	return ChargebackEventsField("provider_event_ref")
}

func (ss ChargebackEventsSelectFields) EventPayload() ChargebackEventsField {
	return ChargebackEventsField("event_payload")
}

func (ss ChargebackEventsSelectFields) OccurredAt() ChargebackEventsField {
	return ChargebackEventsField("occurred_at")
}

func (ss ChargebackEventsSelectFields) Metadata() ChargebackEventsField {
	return ChargebackEventsField("metadata")
}

func (ss ChargebackEventsSelectFields) MetaCreatedAt() ChargebackEventsField {
	return ChargebackEventsField("meta_created_at")
}

func (ss ChargebackEventsSelectFields) MetaCreatedBy() ChargebackEventsField {
	return ChargebackEventsField("meta_created_by")
}

func (ss ChargebackEventsSelectFields) MetaUpdatedAt() ChargebackEventsField {
	return ChargebackEventsField("meta_updated_at")
}

func (ss ChargebackEventsSelectFields) MetaUpdatedBy() ChargebackEventsField {
	return ChargebackEventsField("meta_updated_by")
}

func (ss ChargebackEventsSelectFields) MetaDeletedAt() ChargebackEventsField {
	return ChargebackEventsField("meta_deleted_at")
}

func (ss ChargebackEventsSelectFields) MetaDeletedBy() ChargebackEventsField {
	return ChargebackEventsField("meta_deleted_by")
}

func (ss ChargebackEventsSelectFields) All() ChargebackEventsFieldList {
	return []ChargebackEventsField{
		ss.Id(),
		ss.DisputeId(),
		ss.EventType(),
		ss.JournalEntryId(),
		ss.Amount(),
		ss.CurrencyCode(),
		ss.ProviderEventRef(),
		ss.EventPayload(),
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

func NewChargebackEventsSelectFields() ChargebackEventsSelectFields {
	return ChargebackEventsSelectFields{}
}

type ChargebackEventsUpdateFieldOption struct {
	useIncrement bool
}
type ChargebackEventsUpdateField struct {
	chargebackEventsField ChargebackEventsField
	opt                   ChargebackEventsUpdateFieldOption
	value                 interface{}
}
type ChargebackEventsUpdateFieldList []ChargebackEventsUpdateField

func defaultChargebackEventsUpdateFieldOption() ChargebackEventsUpdateFieldOption {
	return ChargebackEventsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementChargebackEventsOption(useIncrement bool) func(*ChargebackEventsUpdateFieldOption) {
	return func(pcufo *ChargebackEventsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewChargebackEventsUpdateField(field ChargebackEventsField, val interface{}, opts ...func(*ChargebackEventsUpdateFieldOption)) ChargebackEventsUpdateField {
	defaultOpt := defaultChargebackEventsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ChargebackEventsUpdateField{
		chargebackEventsField: field,
		value:                 val,
		opt:                   defaultOpt,
	}
}
func defaultChargebackEventsUpdateFields(chargebackEvents model.ChargebackEvents) (chargebackEventsUpdateFieldList ChargebackEventsUpdateFieldList) {
	selectFields := NewChargebackEventsSelectFields()
	chargebackEventsUpdateFieldList = append(chargebackEventsUpdateFieldList,
		NewChargebackEventsUpdateField(selectFields.Id(), chargebackEvents.Id),
		NewChargebackEventsUpdateField(selectFields.DisputeId(), chargebackEvents.DisputeId),
		NewChargebackEventsUpdateField(selectFields.EventType(), chargebackEvents.EventType),
		NewChargebackEventsUpdateField(selectFields.JournalEntryId(), chargebackEvents.JournalEntryId),
		NewChargebackEventsUpdateField(selectFields.Amount(), chargebackEvents.Amount),
		NewChargebackEventsUpdateField(selectFields.CurrencyCode(), chargebackEvents.CurrencyCode),
		NewChargebackEventsUpdateField(selectFields.ProviderEventRef(), chargebackEvents.ProviderEventRef),
		NewChargebackEventsUpdateField(selectFields.EventPayload(), chargebackEvents.EventPayload),
		NewChargebackEventsUpdateField(selectFields.OccurredAt(), chargebackEvents.OccurredAt),
		NewChargebackEventsUpdateField(selectFields.Metadata(), chargebackEvents.Metadata),
		NewChargebackEventsUpdateField(selectFields.MetaCreatedAt(), chargebackEvents.MetaCreatedAt),
		NewChargebackEventsUpdateField(selectFields.MetaCreatedBy(), chargebackEvents.MetaCreatedBy),
		NewChargebackEventsUpdateField(selectFields.MetaUpdatedAt(), chargebackEvents.MetaUpdatedAt),
		NewChargebackEventsUpdateField(selectFields.MetaUpdatedBy(), chargebackEvents.MetaUpdatedBy),
		NewChargebackEventsUpdateField(selectFields.MetaDeletedAt(), chargebackEvents.MetaDeletedAt),
		NewChargebackEventsUpdateField(selectFields.MetaDeletedBy(), chargebackEvents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsChargebackEventsCommand(chargebackEventsUpdateFieldList ChargebackEventsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range chargebackEventsUpdateFieldList {
		field := string(updateField.chargebackEventsField)
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

func (repo *RepositoryImpl) BulkCreateChargebackEvents(ctx context.Context, chargebackEventsList []*model.ChargebackEvents, fieldsInsert ...ChargebackEventsField) (err error) {
	var (
		fieldsStr                 string
		valueListStr              []string
		argsList                  []interface{}
		primaryIds                []model.ChargebackEventsPrimaryID
		chargebackEventsValueList []model.ChargebackEvents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewChargebackEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, chargebackEvents := range chargebackEventsList {

		primaryIds = append(primaryIds, chargebackEvents.ToChargebackEventsPrimaryID())

		chargebackEventsValueList = append(chargebackEventsValueList, *chargebackEvents)
	}

	_, notFoundIds, err := repo.IsExistChargebackEventsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateChargebackEvents] failed checking chargebackEvents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ChargebackEventsPrimaryID{}
		mapNotFoundIds := map[model.ChargebackEventsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "chargebackEvents", fmt.Sprintf("chargebackEvents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsChargebackEvents(chargebackEventsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(chargebackEventsQueries.insertChargebackEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateChargebackEvents] failed exec create chargebackEvents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteChargebackEventsByIDs(ctx context.Context, primaryIDs []model.ChargebackEventsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistChargebackEventsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteChargebackEventsByIDs] failed checking chargebackEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebackEvents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"chargeback_events\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(chargebackEventsQueries.deleteChargebackEvents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteChargebackEventsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteChargebackEventsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistChargebackEventsByIDs(ctx context.Context, ids []model.ChargebackEventsPrimaryID) (exists bool, notFoundIds []model.ChargebackEventsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"chargeback_events\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(chargebackEventsQueries.selectChargebackEvents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistChargebackEventsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ChargebackEventsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistChargebackEventsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ChargebackEventsPrimaryID]bool{}
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

// BulkUpdateChargebackEvents is used to bulk update chargebackEvents, by default it will update all field
// if want to update specific field, then fill chargebackEventssMapUpdateFieldsRequest else please fill chargebackEventssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateChargebackEvents(ctx context.Context, chargebackEventssMap map[model.ChargebackEventsPrimaryID]*model.ChargebackEvents, chargebackEventssMapUpdateFieldsRequest map[model.ChargebackEventsPrimaryID]ChargebackEventsUpdateFieldList) (err error) {
	if len(chargebackEventssMap) == 0 && len(chargebackEventssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		chargebackEventssMapUpdateField map[model.ChargebackEventsPrimaryID]ChargebackEventsUpdateFieldList = map[model.ChargebackEventsPrimaryID]ChargebackEventsUpdateFieldList{}
		asTableValues                   string                                                              = "myvalues"
	)

	if len(chargebackEventssMap) > 0 {
		for id, chargebackEvents := range chargebackEventssMap {
			if chargebackEvents == nil {
				log.Error().Err(err).Msg("[BulkUpdateChargebackEvents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			chargebackEventssMapUpdateField[id] = defaultChargebackEventsUpdateFields(*chargebackEvents)
		}
	} else {
		chargebackEventssMapUpdateField = chargebackEventssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateChargebackEventsQuery(chargebackEventssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistChargebackEventsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateChargebackEvents] failed checking chargebackEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebackEvents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeChargebackEventsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"chargeback_events\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateChargebackEvents] failed exec query")
	}
	return
}

type ChargebackEventsFieldParameter struct {
	param string
	args  []interface{}
}

func NewChargebackEventsFieldParameter(param string, args ...interface{}) ChargebackEventsFieldParameter {
	return ChargebackEventsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateChargebackEventsQuery(mapChargebackEventss map[model.ChargebackEventsPrimaryID]ChargebackEventsUpdateFieldList, asTableValues string) (primaryIDs []model.ChargebackEventsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ChargebackEventsPrimaryID]map[string]interface{}{}
	chargebackEventsSelectFields := NewChargebackEventsSelectFields()
	for id, updateFields := range mapChargebackEventss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.chargebackEventsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapChargebackEventss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetChargebackEventsFieldType(updateField.chargebackEventsField)))
			args = append(args, fields[string(updateField.chargebackEventsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.chargebackEventsField))
		if updateField.chargebackEventsField == chargebackEventsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.chargebackEventsField, asTableValues, updateField.chargebackEventsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.chargebackEventsField,
				"\"chargeback_events\"", updateField.chargebackEventsField,
				asTableValues, updateField.chargebackEventsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeChargebackEventsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ChargebackEventsPrimaryID, asTableValue string) (whereQry string) {
	chargebackEventsSelectFields := NewChargebackEventsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"chargeback_events\".\"id\" = %s.\"id\"::"+GetChargebackEventsFieldType(chargebackEventsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetChargebackEventsFieldType(chargebackEventsField ChargebackEventsField) string {
	selectChargebackEventsFields := NewChargebackEventsSelectFields()
	switch chargebackEventsField {

	case selectChargebackEventsFields.Id():
		return "uuid"

	case selectChargebackEventsFields.DisputeId():
		return "uuid"

	case selectChargebackEventsFields.EventType():
		return "event_type_enum"

	case selectChargebackEventsFields.JournalEntryId():
		return "uuid"

	case selectChargebackEventsFields.Amount():
		return "numeric"

	case selectChargebackEventsFields.CurrencyCode():
		return "text"

	case selectChargebackEventsFields.ProviderEventRef():
		return "text"

	case selectChargebackEventsFields.EventPayload():
		return "jsonb"

	case selectChargebackEventsFields.OccurredAt():
		return "timestamptz"

	case selectChargebackEventsFields.Metadata():
		return "jsonb"

	case selectChargebackEventsFields.MetaCreatedAt():
		return "timestamptz"

	case selectChargebackEventsFields.MetaCreatedBy():
		return "uuid"

	case selectChargebackEventsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectChargebackEventsFields.MetaUpdatedBy():
		return "uuid"

	case selectChargebackEventsFields.MetaDeletedAt():
		return "timestamptz"

	case selectChargebackEventsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateChargebackEvents(ctx context.Context, chargebackEvents *model.ChargebackEvents, fieldsInsert ...ChargebackEventsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewChargebackEventsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ChargebackEventsPrimaryID{
		Id: chargebackEvents.Id,
	}
	exists, err := repo.IsExistChargebackEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateChargebackEvents] failed checking chargebackEvents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "chargebackEvents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsChargebackEvents([]model.ChargebackEvents{*chargebackEvents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(chargebackEventsQueries.insertChargebackEvents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateChargebackEvents] failed exec create chargebackEvents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteChargebackEventsByID(ctx context.Context, primaryID model.ChargebackEventsPrimaryID) (err error) {
	exists, err := repo.IsExistChargebackEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteChargebackEventsByID] failed checking chargebackEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebackEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeChargebackEventsCompositePrimaryKeyWhere([]model.ChargebackEventsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(chargebackEventsQueries.deleteChargebackEvents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteChargebackEventsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveChargebackEventsByFilter(ctx context.Context, filter model.Filter) (result []model.ChargebackEventsFilterResult, err error) {
	query, args, err := composeChargebackEventsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveChargebackEventsByFilter] failed compose chargebackEvents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveChargebackEventsByFilter] failed get chargebackEvents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeChargebackEventsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ChargebackEventsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeChargebackEventsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeChargebackEventsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeChargebackEventsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 16 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewChargebackEventsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 16+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["dispute_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"dispute_id\"")
			selectedColumns["dispute_id"] = struct{}{}
		}
		if _, selected := selectedColumns["event_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_type\"")
			selectedColumns["event_type"] = struct{}{}
		}
		if _, selected := selectedColumns["journal_entry_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"journal_entry_id\"")
			selectedColumns["journal_entry_id"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_event_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_event_ref\"")
			selectedColumns["provider_event_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["event_payload"]; !selected {
			selectColumns = append(selectColumns, "base.\"event_payload\"")
			selectedColumns["event_payload"] = struct{}{}
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

type chargebackEventsFilterPlaceholder struct {
	index int
}

func (p *chargebackEventsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeChargebackEventsFilterPredicate(filterField model.FilterField, placeholders *chargebackEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewChargebackEventsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeChargebackEventsFilterSQLExpr(spec)
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

func composeChargebackEventsFilterGroup(group model.FilterGroup, placeholders *chargebackEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeChargebackEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeChargebackEventsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeChargebackEventsFilterWhereQueries(filter model.Filter, placeholders *chargebackEventsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeChargebackEventsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeChargebackEventsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeChargebackEventsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateChargebackEventsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeChargebackEventsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeChargebackEventsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := chargebackEventsFilterPlaceholder{index: 1}
	whereQueries, err := composeChargebackEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewChargebackEventsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeChargebackEventsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeChargebackEventsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"chargeback_events\" base%s", strings.Join(selectColumns, ","), composeChargebackEventsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistChargebackEventsByID(ctx context.Context, primaryID model.ChargebackEventsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeChargebackEventsCompositePrimaryKeyWhere([]model.ChargebackEventsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", chargebackEventsQueries.selectCountChargebackEvents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistChargebackEventsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveChargebackEvents(ctx context.Context, selectFields ...ChargebackEventsField) (chargebackEventsList model.ChargebackEventsList, err error) {
	var (
		defaultChargebackEventsSelectFields = defaultChargebackEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultChargebackEventsSelectFields = composeChargebackEventsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(chargebackEventsQueries.selectChargebackEvents, defaultChargebackEventsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &chargebackEventsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveChargebackEvents] failed get chargebackEvents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveChargebackEventsByID(ctx context.Context, primaryID model.ChargebackEventsPrimaryID, selectFields ...ChargebackEventsField) (chargebackEvents model.ChargebackEvents, err error) {
	var (
		defaultChargebackEventsSelectFields = defaultChargebackEventsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultChargebackEventsSelectFields = composeChargebackEventsSelectFields(selectFields...)
	}
	whereQry, params := composeChargebackEventsCompositePrimaryKeyWhere([]model.ChargebackEventsPrimaryID{primaryID})
	query := fmt.Sprintf(chargebackEventsQueries.selectChargebackEvents+" WHERE "+whereQry, defaultChargebackEventsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &chargebackEvents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("chargebackEvents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveChargebackEventsByID] failed get chargebackEvents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateChargebackEventsByID(ctx context.Context, primaryID model.ChargebackEventsPrimaryID, chargebackEvents *model.ChargebackEvents, chargebackEventsUpdateFields ...ChargebackEventsUpdateField) (err error) {
	exists, err := repo.IsExistChargebackEventsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebackEvents] failed checking chargebackEvents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebackEvents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if chargebackEvents == nil {
		if len(chargebackEventsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateChargebackEventsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		chargebackEvents = &model.ChargebackEvents{}
	}
	var (
		defaultChargebackEventsUpdateFields = defaultChargebackEventsUpdateFields(*chargebackEvents)
		tempUpdateField                     ChargebackEventsUpdateFieldList
		selectFields                        = NewChargebackEventsSelectFields()
	)
	if len(chargebackEventsUpdateFields) > 0 {
		for _, updateField := range chargebackEventsUpdateFields {
			if updateField.chargebackEventsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultChargebackEventsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeChargebackEventsCompositePrimaryKeyWhere([]model.ChargebackEventsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsChargebackEventsCommand(defaultChargebackEventsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(chargebackEventsQueries.updateChargebackEvents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebackEvents] error when try to update chargebackEvents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateChargebackEventsByFilter(ctx context.Context, filter model.Filter, chargebackEventsUpdateFields ...ChargebackEventsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(chargebackEventsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ChargebackEventsUpdateFieldList
		selectFields = NewChargebackEventsSelectFields()
	)
	for _, updateField := range chargebackEventsUpdateFields {
		if updateField.chargebackEventsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsChargebackEventsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := chargebackEventsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeChargebackEventsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"chargeback_events\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebackEventsByFilter] error when try to update chargebackEvents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebackEventsByFilter] failed get rows affected")
	}
	return
}

var (
	chargebackEventsQueries = struct {
		selectChargebackEvents      string
		selectCountChargebackEvents string
		deleteChargebackEvents      string
		updateChargebackEvents      string
		insertChargebackEvents      string
	}{
		selectChargebackEvents:      "SELECT %s FROM \"chargeback_events\"",
		selectCountChargebackEvents: "SELECT COUNT(\"id\") FROM \"chargeback_events\"",
		deleteChargebackEvents:      "DELETE FROM \"chargeback_events\"",
		updateChargebackEvents:      "UPDATE \"chargeback_events\" SET %s ",
		insertChargebackEvents:      "INSERT INTO \"chargeback_events\" %s VALUES %s",
	}
)

type ChargebackEventsRepository interface {
	CreateChargebackEvents(ctx context.Context, chargebackEvents *model.ChargebackEvents, fieldsInsert ...ChargebackEventsField) error
	BulkCreateChargebackEvents(ctx context.Context, chargebackEventsList []*model.ChargebackEvents, fieldsInsert ...ChargebackEventsField) error
	ResolveChargebackEvents(ctx context.Context, selectFields ...ChargebackEventsField) (model.ChargebackEventsList, error)
	ResolveChargebackEventsByID(ctx context.Context, primaryID model.ChargebackEventsPrimaryID, selectFields ...ChargebackEventsField) (model.ChargebackEvents, error)
	UpdateChargebackEventsByID(ctx context.Context, id model.ChargebackEventsPrimaryID, chargebackEvents *model.ChargebackEvents, chargebackEventsUpdateFields ...ChargebackEventsUpdateField) error
	UpdateChargebackEventsByFilter(ctx context.Context, filter model.Filter, chargebackEventsUpdateFields ...ChargebackEventsUpdateField) (rowsAffected int64, err error)
	BulkUpdateChargebackEvents(ctx context.Context, chargebackEventsListMap map[model.ChargebackEventsPrimaryID]*model.ChargebackEvents, ChargebackEventssMapUpdateFieldsRequest map[model.ChargebackEventsPrimaryID]ChargebackEventsUpdateFieldList) (err error)
	DeleteChargebackEventsByID(ctx context.Context, id model.ChargebackEventsPrimaryID) error
	BulkDeleteChargebackEventsByIDs(ctx context.Context, ids []model.ChargebackEventsPrimaryID) error
	ResolveChargebackEventsByFilter(ctx context.Context, filter model.Filter) (result []model.ChargebackEventsFilterResult, err error)
	IsExistChargebackEventsByIDs(ctx context.Context, ids []model.ChargebackEventsPrimaryID) (exists bool, notFoundIds []model.ChargebackEventsPrimaryID, err error)
	IsExistChargebackEventsByID(ctx context.Context, id model.ChargebackEventsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
