package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/routing/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsRoutingDecisions(routingDecisionsList []model.RoutingDecisions, fieldsInsert ...RoutingDecisionsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRoutingDecisionsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, routingDecisions := range routingDecisionsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, routingDecisions.Id)
			case selectField.PaymentIntentId():
				args = append(args, routingDecisions.PaymentIntentId)
			case selectField.ProfileId():
				args = append(args, routingDecisions.ProfileId)
			case selectField.RuleId():
				args = append(args, routingDecisions.RuleId)
			case selectField.StrategyUsed():
				args = append(args, routingDecisions.StrategyUsed)
			case selectField.CandidatePsps():
				args = append(args, routingDecisions.CandidatePsps)
			case selectField.DecisionReason():
				args = append(args, routingDecisions.DecisionReason)
			case selectField.DecidedAt():
				args = append(args, routingDecisions.DecidedAt)
			case selectField.MetaCreatedAt():
				args = append(args, routingDecisions.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, routingDecisions.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, routingDecisions.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, routingDecisions.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, routingDecisions.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, routingDecisions.MetaDeletedBy)

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

func composeRoutingDecisionsCompositePrimaryKeyWhere(primaryIDs []model.RoutingDecisionsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"routing_decisions\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRoutingDecisionsSelectFields() string {
	fields := NewRoutingDecisionsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRoutingDecisionsSelectFields(selectFields ...RoutingDecisionsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RoutingDecisionsField string
type RoutingDecisionsFieldList []RoutingDecisionsField

type RoutingDecisionsSelectFields struct {
}

func (ss RoutingDecisionsSelectFields) Id() RoutingDecisionsField {
	return RoutingDecisionsField("id")
}

func (ss RoutingDecisionsSelectFields) PaymentIntentId() RoutingDecisionsField {
	return RoutingDecisionsField("payment_intent_id")
}

func (ss RoutingDecisionsSelectFields) ProfileId() RoutingDecisionsField {
	return RoutingDecisionsField("profile_id")
}

func (ss RoutingDecisionsSelectFields) RuleId() RoutingDecisionsField {
	return RoutingDecisionsField("rule_id")
}

func (ss RoutingDecisionsSelectFields) StrategyUsed() RoutingDecisionsField {
	return RoutingDecisionsField("strategy_used")
}

func (ss RoutingDecisionsSelectFields) CandidatePsps() RoutingDecisionsField {
	return RoutingDecisionsField("candidate_psps")
}

func (ss RoutingDecisionsSelectFields) DecisionReason() RoutingDecisionsField {
	return RoutingDecisionsField("decision_reason")
}

func (ss RoutingDecisionsSelectFields) DecidedAt() RoutingDecisionsField {
	return RoutingDecisionsField("decided_at")
}

func (ss RoutingDecisionsSelectFields) MetaCreatedAt() RoutingDecisionsField {
	return RoutingDecisionsField("meta_created_at")
}

func (ss RoutingDecisionsSelectFields) MetaCreatedBy() RoutingDecisionsField {
	return RoutingDecisionsField("meta_created_by")
}

func (ss RoutingDecisionsSelectFields) MetaUpdatedAt() RoutingDecisionsField {
	return RoutingDecisionsField("meta_updated_at")
}

func (ss RoutingDecisionsSelectFields) MetaUpdatedBy() RoutingDecisionsField {
	return RoutingDecisionsField("meta_updated_by")
}

func (ss RoutingDecisionsSelectFields) MetaDeletedAt() RoutingDecisionsField {
	return RoutingDecisionsField("meta_deleted_at")
}

func (ss RoutingDecisionsSelectFields) MetaDeletedBy() RoutingDecisionsField {
	return RoutingDecisionsField("meta_deleted_by")
}

func (ss RoutingDecisionsSelectFields) All() RoutingDecisionsFieldList {
	return []RoutingDecisionsField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.ProfileId(),
		ss.RuleId(),
		ss.StrategyUsed(),
		ss.CandidatePsps(),
		ss.DecisionReason(),
		ss.DecidedAt(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRoutingDecisionsSelectFields() RoutingDecisionsSelectFields {
	return RoutingDecisionsSelectFields{}
}

type RoutingDecisionsUpdateFieldOption struct {
	useIncrement bool
}
type RoutingDecisionsUpdateField struct {
	routingDecisionsField RoutingDecisionsField
	opt                   RoutingDecisionsUpdateFieldOption
	value                 interface{}
}
type RoutingDecisionsUpdateFieldList []RoutingDecisionsUpdateField

func defaultRoutingDecisionsUpdateFieldOption() RoutingDecisionsUpdateFieldOption {
	return RoutingDecisionsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRoutingDecisionsOption(useIncrement bool) func(*RoutingDecisionsUpdateFieldOption) {
	return func(pcufo *RoutingDecisionsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRoutingDecisionsUpdateField(field RoutingDecisionsField, val interface{}, opts ...func(*RoutingDecisionsUpdateFieldOption)) RoutingDecisionsUpdateField {
	defaultOpt := defaultRoutingDecisionsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RoutingDecisionsUpdateField{
		routingDecisionsField: field,
		value:                 val,
		opt:                   defaultOpt,
	}
}
func defaultRoutingDecisionsUpdateFields(routingDecisions model.RoutingDecisions) (routingDecisionsUpdateFieldList RoutingDecisionsUpdateFieldList) {
	selectFields := NewRoutingDecisionsSelectFields()
	routingDecisionsUpdateFieldList = append(routingDecisionsUpdateFieldList,
		NewRoutingDecisionsUpdateField(selectFields.Id(), routingDecisions.Id),
		NewRoutingDecisionsUpdateField(selectFields.PaymentIntentId(), routingDecisions.PaymentIntentId),
		NewRoutingDecisionsUpdateField(selectFields.ProfileId(), routingDecisions.ProfileId),
		NewRoutingDecisionsUpdateField(selectFields.RuleId(), routingDecisions.RuleId),
		NewRoutingDecisionsUpdateField(selectFields.StrategyUsed(), routingDecisions.StrategyUsed),
		NewRoutingDecisionsUpdateField(selectFields.CandidatePsps(), routingDecisions.CandidatePsps),
		NewRoutingDecisionsUpdateField(selectFields.DecisionReason(), routingDecisions.DecisionReason),
		NewRoutingDecisionsUpdateField(selectFields.DecidedAt(), routingDecisions.DecidedAt),
		NewRoutingDecisionsUpdateField(selectFields.MetaCreatedAt(), routingDecisions.MetaCreatedAt),
		NewRoutingDecisionsUpdateField(selectFields.MetaCreatedBy(), routingDecisions.MetaCreatedBy),
		NewRoutingDecisionsUpdateField(selectFields.MetaUpdatedAt(), routingDecisions.MetaUpdatedAt),
		NewRoutingDecisionsUpdateField(selectFields.MetaUpdatedBy(), routingDecisions.MetaUpdatedBy),
		NewRoutingDecisionsUpdateField(selectFields.MetaDeletedAt(), routingDecisions.MetaDeletedAt),
		NewRoutingDecisionsUpdateField(selectFields.MetaDeletedBy(), routingDecisions.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRoutingDecisionsCommand(routingDecisionsUpdateFieldList RoutingDecisionsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range routingDecisionsUpdateFieldList {
		field := string(updateField.routingDecisionsField)
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

func (repo *RepositoryImpl) BulkCreateRoutingDecisions(ctx context.Context, routingDecisionsList []*model.RoutingDecisions, fieldsInsert ...RoutingDecisionsField) (err error) {
	var (
		fieldsStr                 string
		valueListStr              []string
		argsList                  []interface{}
		primaryIds                []model.RoutingDecisionsPrimaryID
		routingDecisionsValueList []model.RoutingDecisions
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRoutingDecisionsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, routingDecisions := range routingDecisionsList {

		primaryIds = append(primaryIds, routingDecisions.ToRoutingDecisionsPrimaryID())

		routingDecisionsValueList = append(routingDecisionsValueList, *routingDecisions)
	}

	_, notFoundIds, err := repo.IsExistRoutingDecisionsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRoutingDecisions] failed checking routingDecisions whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RoutingDecisionsPrimaryID{}
		mapNotFoundIds := map[model.RoutingDecisionsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "routingDecisions", fmt.Sprintf("routingDecisions with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRoutingDecisions(routingDecisionsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(routingDecisionsQueries.insertRoutingDecisions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRoutingDecisions] failed exec create routingDecisions query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRoutingDecisionsByIDs(ctx context.Context, primaryIDs []model.RoutingDecisionsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRoutingDecisionsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingDecisionsByIDs] failed checking routingDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingDecisions with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"routing_decisions\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(routingDecisionsQueries.deleteRoutingDecisions + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingDecisionsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingDecisionsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRoutingDecisionsByIDs(ctx context.Context, ids []model.RoutingDecisionsPrimaryID) (exists bool, notFoundIds []model.RoutingDecisionsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"routing_decisions\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(routingDecisionsQueries.selectRoutingDecisions, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingDecisionsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RoutingDecisionsPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingDecisionsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RoutingDecisionsPrimaryID]bool{}
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

// BulkUpdateRoutingDecisions is used to bulk update routingDecisions, by default it will update all field
// if want to update specific field, then fill routingDecisionssMapUpdateFieldsRequest else please fill routingDecisionssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRoutingDecisions(ctx context.Context, routingDecisionssMap map[model.RoutingDecisionsPrimaryID]*model.RoutingDecisions, routingDecisionssMapUpdateFieldsRequest map[model.RoutingDecisionsPrimaryID]RoutingDecisionsUpdateFieldList) (err error) {
	if len(routingDecisionssMap) == 0 && len(routingDecisionssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		routingDecisionssMapUpdateField map[model.RoutingDecisionsPrimaryID]RoutingDecisionsUpdateFieldList = map[model.RoutingDecisionsPrimaryID]RoutingDecisionsUpdateFieldList{}
		asTableValues                   string                                                              = "myvalues"
	)

	if len(routingDecisionssMap) > 0 {
		for id, routingDecisions := range routingDecisionssMap {
			if routingDecisions == nil {
				log.Error().Err(err).Msg("[BulkUpdateRoutingDecisions] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			routingDecisionssMapUpdateField[id] = defaultRoutingDecisionsUpdateFields(*routingDecisions)
		}
	} else {
		routingDecisionssMapUpdateField = routingDecisionssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRoutingDecisionsQuery(routingDecisionssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRoutingDecisionsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRoutingDecisions] failed checking routingDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingDecisions with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRoutingDecisionsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"routing_decisions\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRoutingDecisions] failed exec query")
	}
	return
}

type RoutingDecisionsFieldParameter struct {
	param string
	args  []interface{}
}

func NewRoutingDecisionsFieldParameter(param string, args ...interface{}) RoutingDecisionsFieldParameter {
	return RoutingDecisionsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRoutingDecisionsQuery(mapRoutingDecisionss map[model.RoutingDecisionsPrimaryID]RoutingDecisionsUpdateFieldList, asTableValues string) (primaryIDs []model.RoutingDecisionsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RoutingDecisionsPrimaryID]map[string]interface{}{}
	routingDecisionsSelectFields := NewRoutingDecisionsSelectFields()
	for id, updateFields := range mapRoutingDecisionss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.routingDecisionsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRoutingDecisionss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRoutingDecisionsFieldType(updateField.routingDecisionsField)))
			args = append(args, fields[string(updateField.routingDecisionsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.routingDecisionsField))
		if updateField.routingDecisionsField == routingDecisionsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.routingDecisionsField, asTableValues, updateField.routingDecisionsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.routingDecisionsField,
				"\"routing_decisions\"", updateField.routingDecisionsField,
				asTableValues, updateField.routingDecisionsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRoutingDecisionsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RoutingDecisionsPrimaryID, asTableValue string) (whereQry string) {
	routingDecisionsSelectFields := NewRoutingDecisionsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"routing_decisions\".\"id\" = %s.\"id\"::"+GetRoutingDecisionsFieldType(routingDecisionsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRoutingDecisionsFieldType(routingDecisionsField RoutingDecisionsField) string {
	selectRoutingDecisionsFields := NewRoutingDecisionsSelectFields()
	switch routingDecisionsField {

	case selectRoutingDecisionsFields.Id():
		return "uuid"

	case selectRoutingDecisionsFields.PaymentIntentId():
		return "uuid"

	case selectRoutingDecisionsFields.ProfileId():
		return "uuid"

	case selectRoutingDecisionsFields.RuleId():
		return "uuid"

	case selectRoutingDecisionsFields.StrategyUsed():
		return "routing_strategy_enum"

	case selectRoutingDecisionsFields.CandidatePsps():
		return "jsonb"

	case selectRoutingDecisionsFields.DecisionReason():
		return "text"

	case selectRoutingDecisionsFields.DecidedAt():
		return "timestamptz"

	case selectRoutingDecisionsFields.MetaCreatedAt():
		return "timestamptz"

	case selectRoutingDecisionsFields.MetaCreatedBy():
		return "uuid"

	case selectRoutingDecisionsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRoutingDecisionsFields.MetaUpdatedBy():
		return "uuid"

	case selectRoutingDecisionsFields.MetaDeletedAt():
		return "timestamptz"

	case selectRoutingDecisionsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRoutingDecisions(ctx context.Context, routingDecisions *model.RoutingDecisions, fieldsInsert ...RoutingDecisionsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRoutingDecisionsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RoutingDecisionsPrimaryID{
		Id: routingDecisions.Id,
	}
	exists, err := repo.IsExistRoutingDecisionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRoutingDecisions] failed checking routingDecisions whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "routingDecisions", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRoutingDecisions([]model.RoutingDecisions{*routingDecisions}, fieldsInsert...)
	commandQuery := fmt.Sprintf(routingDecisionsQueries.insertRoutingDecisions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRoutingDecisions] failed exec create routingDecisions query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRoutingDecisionsByID(ctx context.Context, primaryID model.RoutingDecisionsPrimaryID) (err error) {
	exists, err := repo.IsExistRoutingDecisionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRoutingDecisionsByID] failed checking routingDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingDecisions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRoutingDecisionsCompositePrimaryKeyWhere([]model.RoutingDecisionsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(routingDecisionsQueries.deleteRoutingDecisions + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRoutingDecisionsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingDecisionsByFilter(ctx context.Context, filter model.Filter) (result []model.RoutingDecisionsFilterResult, err error) {
	query, args, err := composeRoutingDecisionsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingDecisionsByFilter] failed compose routingDecisions filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingDecisionsByFilter] failed get routingDecisions by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRoutingDecisionsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRoutingDecisionsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultRoutingDecisionsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := RoutingDecisionsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, RoutingDecisionsField(filterSelectField))
		}
		selectFields = composeRoutingDecisionsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(routingDecisionsQueries.selectRoutingDecisions, selectFields)

	if len(filter.FilterFields) > 0 {
		var (
			whereQueries []string
			whereArgs    []interface{}
		)
		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorEqual:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" = $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			case model.OperatorRange:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" BETWEEN $%d AND $%d", filterField.Field, index, index+1))
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				whereArgs = append(whereArgs, valueArray...)
				index += 2
			case model.OperatorIn:
				valueArray, ok := filterField.Value.([]interface{})
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				var placeholder []string
				for range valueArray {
					placeholder = append(placeholder, fmt.Sprintf("$%d", index))
					index++
				}
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IN (%s)", filterField.Field, strings.Join(placeholder, ",")))
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIsNull:
				value, ok := filterField.Value.(bool)
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				if value {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NULL", filterField.Field))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NOT NULL", filterField.Field))
				}
			case model.OperatorNot:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" != $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			}
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND "))
		args = append(args, whereArgs...)
	}

	sortQuery := []string{}
	for _, sort := range filter.Sorts {
		sortQuery = append(sortQuery, fmt.Sprintf("\"%s\" %s", sort.Field, sort.Order))
	}
	if len(sortQuery) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
	}
	if filter.Pagination.PageSize > 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Pagination.PageSize)
		if filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
	}

	return
}

func (repo *RepositoryImpl) IsExistRoutingDecisionsByID(ctx context.Context, primaryID model.RoutingDecisionsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRoutingDecisionsCompositePrimaryKeyWhere([]model.RoutingDecisionsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", routingDecisionsQueries.selectCountRoutingDecisions, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingDecisionsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingDecisions(ctx context.Context, selectFields ...RoutingDecisionsField) (routingDecisionsList model.RoutingDecisionsList, err error) {
	var (
		defaultRoutingDecisionsSelectFields = defaultRoutingDecisionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRoutingDecisionsSelectFields = composeRoutingDecisionsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(routingDecisionsQueries.selectRoutingDecisions, defaultRoutingDecisionsSelectFields)

	err = repo.db.Read.Select(&routingDecisionsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingDecisions] failed get routingDecisions list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingDecisionsByID(ctx context.Context, primaryID model.RoutingDecisionsPrimaryID, selectFields ...RoutingDecisionsField) (routingDecisions model.RoutingDecisions, err error) {
	var (
		defaultRoutingDecisionsSelectFields = defaultRoutingDecisionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRoutingDecisionsSelectFields = composeRoutingDecisionsSelectFields(selectFields...)
	}
	whereQry, params := composeRoutingDecisionsCompositePrimaryKeyWhere([]model.RoutingDecisionsPrimaryID{primaryID})
	query := fmt.Sprintf(routingDecisionsQueries.selectRoutingDecisions+" WHERE "+whereQry, defaultRoutingDecisionsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&routingDecisions, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("routingDecisions with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRoutingDecisionsByID] failed get routingDecisions")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRoutingDecisionsByID(ctx context.Context, primaryID model.RoutingDecisionsPrimaryID, routingDecisions *model.RoutingDecisions, routingDecisionsUpdateFields ...RoutingDecisionsUpdateField) (err error) {
	exists, err := repo.IsExistRoutingDecisionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRoutingDecisions] failed checking routingDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingDecisions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if routingDecisions == nil {
		if len(routingDecisionsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRoutingDecisionsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		routingDecisions = &model.RoutingDecisions{}
	}
	var (
		defaultRoutingDecisionsUpdateFields = defaultRoutingDecisionsUpdateFields(*routingDecisions)
		tempUpdateField                     RoutingDecisionsUpdateFieldList
		selectFields                        = NewRoutingDecisionsSelectFields()
	)
	if len(routingDecisionsUpdateFields) > 0 {
		for _, updateField := range routingDecisionsUpdateFields {
			if updateField.routingDecisionsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRoutingDecisionsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRoutingDecisionsCompositePrimaryKeyWhere([]model.RoutingDecisionsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRoutingDecisionsCommand(defaultRoutingDecisionsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(routingDecisionsQueries.updateRoutingDecisions+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRoutingDecisions] error when try to update routingDecisions by id")
	}
	return err
}

var (
	routingDecisionsQueries = struct {
		selectRoutingDecisions      string
		selectCountRoutingDecisions string
		deleteRoutingDecisions      string
		updateRoutingDecisions      string
		insertRoutingDecisions      string
	}{
		selectRoutingDecisions:      "SELECT %s FROM \"routing_decisions\"",
		selectCountRoutingDecisions: "SELECT COUNT(\"id\") FROM \"routing_decisions\"",
		deleteRoutingDecisions:      "DELETE FROM \"routing_decisions\"",
		updateRoutingDecisions:      "UPDATE \"routing_decisions\" SET %s ",
		insertRoutingDecisions:      "INSERT INTO \"routing_decisions\" %s VALUES %s",
	}
)

type RoutingDecisionsRepository interface {
	CreateRoutingDecisions(ctx context.Context, routingDecisions *model.RoutingDecisions, fieldsInsert ...RoutingDecisionsField) error
	BulkCreateRoutingDecisions(ctx context.Context, routingDecisionsList []*model.RoutingDecisions, fieldsInsert ...RoutingDecisionsField) error
	ResolveRoutingDecisions(ctx context.Context, selectFields ...RoutingDecisionsField) (model.RoutingDecisionsList, error)
	ResolveRoutingDecisionsByID(ctx context.Context, primaryID model.RoutingDecisionsPrimaryID, selectFields ...RoutingDecisionsField) (model.RoutingDecisions, error)
	UpdateRoutingDecisionsByID(ctx context.Context, id model.RoutingDecisionsPrimaryID, routingDecisions *model.RoutingDecisions, routingDecisionsUpdateFields ...RoutingDecisionsUpdateField) error
	BulkUpdateRoutingDecisions(ctx context.Context, routingDecisionsListMap map[model.RoutingDecisionsPrimaryID]*model.RoutingDecisions, RoutingDecisionssMapUpdateFieldsRequest map[model.RoutingDecisionsPrimaryID]RoutingDecisionsUpdateFieldList) (err error)
	DeleteRoutingDecisionsByID(ctx context.Context, id model.RoutingDecisionsPrimaryID) error
	BulkDeleteRoutingDecisionsByIDs(ctx context.Context, ids []model.RoutingDecisionsPrimaryID) error
	ResolveRoutingDecisionsByFilter(ctx context.Context, filter model.Filter) (result []model.RoutingDecisionsFilterResult, err error)
	IsExistRoutingDecisionsByIDs(ctx context.Context, ids []model.RoutingDecisionsPrimaryID) (exists bool, notFoundIds []model.RoutingDecisionsPrimaryID, err error)
	IsExistRoutingDecisionsByID(ctx context.Context, id model.RoutingDecisionsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
