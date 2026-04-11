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

func composeInsertFieldsAndParamsRoutingRules(routingRulesList []model.RoutingRules, fieldsInsert ...RoutingRulesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRoutingRulesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, routingRules := range routingRulesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, routingRules.Id)
			case selectField.ProfileId():
				args = append(args, routingRules.ProfileId)
			case selectField.Priority():
				args = append(args, routingRules.Priority)
			case selectField.Name():
				args = append(args, routingRules.Name)
			case selectField.IsActive():
				args = append(args, routingRules.IsActive)
			case selectField.MatchPaymentMethod():
				args = append(args, routingRules.MatchPaymentMethod)
			case selectField.MatchCurrency():
				args = append(args, routingRules.MatchCurrency)
			case selectField.MatchAmountMin():
				args = append(args, routingRules.MatchAmountMin)
			case selectField.MatchAmountMax():
				args = append(args, routingRules.MatchAmountMax)
			case selectField.MatchUserCountry():
				args = append(args, routingRules.MatchUserCountry)
			case selectField.MatchCardBin():
				args = append(args, routingRules.MatchCardBin)
			case selectField.MatchProductType():
				args = append(args, routingRules.MatchProductType)
			case selectField.CostWeight():
				args = append(args, routingRules.CostWeight)
			case selectField.Notes():
				args = append(args, routingRules.Notes)
			case selectField.MetaCreatedAt():
				args = append(args, routingRules.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, routingRules.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, routingRules.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, routingRules.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, routingRules.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, routingRules.MetaDeletedBy)

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

func composeRoutingRulesCompositePrimaryKeyWhere(primaryIDs []model.RoutingRulesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"routing_rules\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRoutingRulesSelectFields() string {
	fields := NewRoutingRulesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRoutingRulesSelectFields(selectFields ...RoutingRulesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RoutingRulesField string
type RoutingRulesFieldList []RoutingRulesField

type RoutingRulesSelectFields struct {
}

func (ss RoutingRulesSelectFields) Id() RoutingRulesField {
	return RoutingRulesField("id")
}

func (ss RoutingRulesSelectFields) ProfileId() RoutingRulesField {
	return RoutingRulesField("profile_id")
}

func (ss RoutingRulesSelectFields) Priority() RoutingRulesField {
	return RoutingRulesField("priority")
}

func (ss RoutingRulesSelectFields) Name() RoutingRulesField {
	return RoutingRulesField("name")
}

func (ss RoutingRulesSelectFields) IsActive() RoutingRulesField {
	return RoutingRulesField("is_active")
}

func (ss RoutingRulesSelectFields) MatchPaymentMethod() RoutingRulesField {
	return RoutingRulesField("match_payment_method")
}

func (ss RoutingRulesSelectFields) MatchCurrency() RoutingRulesField {
	return RoutingRulesField("match_currency")
}

func (ss RoutingRulesSelectFields) MatchAmountMin() RoutingRulesField {
	return RoutingRulesField("match_amount_min")
}

func (ss RoutingRulesSelectFields) MatchAmountMax() RoutingRulesField {
	return RoutingRulesField("match_amount_max")
}

func (ss RoutingRulesSelectFields) MatchUserCountry() RoutingRulesField {
	return RoutingRulesField("match_user_country")
}

func (ss RoutingRulesSelectFields) MatchCardBin() RoutingRulesField {
	return RoutingRulesField("match_card_bin")
}

func (ss RoutingRulesSelectFields) MatchProductType() RoutingRulesField {
	return RoutingRulesField("match_product_type")
}

func (ss RoutingRulesSelectFields) CostWeight() RoutingRulesField {
	return RoutingRulesField("cost_weight")
}

func (ss RoutingRulesSelectFields) Notes() RoutingRulesField {
	return RoutingRulesField("notes")
}

func (ss RoutingRulesSelectFields) MetaCreatedAt() RoutingRulesField {
	return RoutingRulesField("meta_created_at")
}

func (ss RoutingRulesSelectFields) MetaCreatedBy() RoutingRulesField {
	return RoutingRulesField("meta_created_by")
}

func (ss RoutingRulesSelectFields) MetaUpdatedAt() RoutingRulesField {
	return RoutingRulesField("meta_updated_at")
}

func (ss RoutingRulesSelectFields) MetaUpdatedBy() RoutingRulesField {
	return RoutingRulesField("meta_updated_by")
}

func (ss RoutingRulesSelectFields) MetaDeletedAt() RoutingRulesField {
	return RoutingRulesField("meta_deleted_at")
}

func (ss RoutingRulesSelectFields) MetaDeletedBy() RoutingRulesField {
	return RoutingRulesField("meta_deleted_by")
}

func (ss RoutingRulesSelectFields) All() RoutingRulesFieldList {
	return []RoutingRulesField{
		ss.Id(),
		ss.ProfileId(),
		ss.Priority(),
		ss.Name(),
		ss.IsActive(),
		ss.MatchPaymentMethod(),
		ss.MatchCurrency(),
		ss.MatchAmountMin(),
		ss.MatchAmountMax(),
		ss.MatchUserCountry(),
		ss.MatchCardBin(),
		ss.MatchProductType(),
		ss.CostWeight(),
		ss.Notes(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRoutingRulesSelectFields() RoutingRulesSelectFields {
	return RoutingRulesSelectFields{}
}

type RoutingRulesUpdateFieldOption struct {
	useIncrement bool
}
type RoutingRulesUpdateField struct {
	routingRulesField RoutingRulesField
	opt               RoutingRulesUpdateFieldOption
	value             interface{}
}
type RoutingRulesUpdateFieldList []RoutingRulesUpdateField

func defaultRoutingRulesUpdateFieldOption() RoutingRulesUpdateFieldOption {
	return RoutingRulesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRoutingRulesOption(useIncrement bool) func(*RoutingRulesUpdateFieldOption) {
	return func(pcufo *RoutingRulesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRoutingRulesUpdateField(field RoutingRulesField, val interface{}, opts ...func(*RoutingRulesUpdateFieldOption)) RoutingRulesUpdateField {
	defaultOpt := defaultRoutingRulesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RoutingRulesUpdateField{
		routingRulesField: field,
		value:             val,
		opt:               defaultOpt,
	}
}
func defaultRoutingRulesUpdateFields(routingRules model.RoutingRules) (routingRulesUpdateFieldList RoutingRulesUpdateFieldList) {
	selectFields := NewRoutingRulesSelectFields()
	routingRulesUpdateFieldList = append(routingRulesUpdateFieldList,
		NewRoutingRulesUpdateField(selectFields.Id(), routingRules.Id),
		NewRoutingRulesUpdateField(selectFields.ProfileId(), routingRules.ProfileId),
		NewRoutingRulesUpdateField(selectFields.Priority(), routingRules.Priority),
		NewRoutingRulesUpdateField(selectFields.Name(), routingRules.Name),
		NewRoutingRulesUpdateField(selectFields.IsActive(), routingRules.IsActive),
		NewRoutingRulesUpdateField(selectFields.MatchPaymentMethod(), routingRules.MatchPaymentMethod),
		NewRoutingRulesUpdateField(selectFields.MatchCurrency(), routingRules.MatchCurrency),
		NewRoutingRulesUpdateField(selectFields.MatchAmountMin(), routingRules.MatchAmountMin),
		NewRoutingRulesUpdateField(selectFields.MatchAmountMax(), routingRules.MatchAmountMax),
		NewRoutingRulesUpdateField(selectFields.MatchUserCountry(), routingRules.MatchUserCountry),
		NewRoutingRulesUpdateField(selectFields.MatchCardBin(), routingRules.MatchCardBin),
		NewRoutingRulesUpdateField(selectFields.MatchProductType(), routingRules.MatchProductType),
		NewRoutingRulesUpdateField(selectFields.CostWeight(), routingRules.CostWeight),
		NewRoutingRulesUpdateField(selectFields.Notes(), routingRules.Notes),
		NewRoutingRulesUpdateField(selectFields.MetaCreatedAt(), routingRules.MetaCreatedAt),
		NewRoutingRulesUpdateField(selectFields.MetaCreatedBy(), routingRules.MetaCreatedBy),
		NewRoutingRulesUpdateField(selectFields.MetaUpdatedAt(), routingRules.MetaUpdatedAt),
		NewRoutingRulesUpdateField(selectFields.MetaUpdatedBy(), routingRules.MetaUpdatedBy),
		NewRoutingRulesUpdateField(selectFields.MetaDeletedAt(), routingRules.MetaDeletedAt),
		NewRoutingRulesUpdateField(selectFields.MetaDeletedBy(), routingRules.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRoutingRulesCommand(routingRulesUpdateFieldList RoutingRulesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range routingRulesUpdateFieldList {
		field := string(updateField.routingRulesField)
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

func (repo *RepositoryImpl) BulkCreateRoutingRules(ctx context.Context, routingRulesList []*model.RoutingRules, fieldsInsert ...RoutingRulesField) (err error) {
	var (
		fieldsStr             string
		valueListStr          []string
		argsList              []interface{}
		primaryIds            []model.RoutingRulesPrimaryID
		routingRulesValueList []model.RoutingRules
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRoutingRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, routingRules := range routingRulesList {

		primaryIds = append(primaryIds, routingRules.ToRoutingRulesPrimaryID())

		routingRulesValueList = append(routingRulesValueList, *routingRules)
	}

	_, notFoundIds, err := repo.IsExistRoutingRulesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRoutingRules] failed checking routingRules whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RoutingRulesPrimaryID{}
		mapNotFoundIds := map[model.RoutingRulesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "routingRules", fmt.Sprintf("routingRules with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRoutingRules(routingRulesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(routingRulesQueries.insertRoutingRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRoutingRules] failed exec create routingRules query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRoutingRulesByIDs(ctx context.Context, primaryIDs []model.RoutingRulesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRoutingRulesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingRulesByIDs] failed checking routingRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingRules with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"routing_rules\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(routingRulesQueries.deleteRoutingRules + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingRulesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingRulesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRoutingRulesByIDs(ctx context.Context, ids []model.RoutingRulesPrimaryID) (exists bool, notFoundIds []model.RoutingRulesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"routing_rules\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(routingRulesQueries.selectRoutingRules, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingRulesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RoutingRulesPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingRulesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RoutingRulesPrimaryID]bool{}
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

// BulkUpdateRoutingRules is used to bulk update routingRules, by default it will update all field
// if want to update specific field, then fill routingRulessMapUpdateFieldsRequest else please fill routingRulessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRoutingRules(ctx context.Context, routingRulessMap map[model.RoutingRulesPrimaryID]*model.RoutingRules, routingRulessMapUpdateFieldsRequest map[model.RoutingRulesPrimaryID]RoutingRulesUpdateFieldList) (err error) {
	if len(routingRulessMap) == 0 && len(routingRulessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		routingRulessMapUpdateField map[model.RoutingRulesPrimaryID]RoutingRulesUpdateFieldList = map[model.RoutingRulesPrimaryID]RoutingRulesUpdateFieldList{}
		asTableValues               string                                                      = "myvalues"
	)

	if len(routingRulessMap) > 0 {
		for id, routingRules := range routingRulessMap {
			if routingRules == nil {
				log.Error().Err(err).Msg("[BulkUpdateRoutingRules] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			routingRulessMapUpdateField[id] = defaultRoutingRulesUpdateFields(*routingRules)
		}
	} else {
		routingRulessMapUpdateField = routingRulessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRoutingRulesQuery(routingRulessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRoutingRulesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRoutingRules] failed checking routingRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingRules with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRoutingRulesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"routing_rules\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRoutingRules] failed exec query")
	}
	return
}

type RoutingRulesFieldParameter struct {
	param string
	args  []interface{}
}

func NewRoutingRulesFieldParameter(param string, args ...interface{}) RoutingRulesFieldParameter {
	return RoutingRulesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRoutingRulesQuery(mapRoutingRuless map[model.RoutingRulesPrimaryID]RoutingRulesUpdateFieldList, asTableValues string) (primaryIDs []model.RoutingRulesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RoutingRulesPrimaryID]map[string]interface{}{}
	routingRulesSelectFields := NewRoutingRulesSelectFields()
	for id, updateFields := range mapRoutingRuless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.routingRulesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRoutingRuless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRoutingRulesFieldType(updateField.routingRulesField)))
			args = append(args, fields[string(updateField.routingRulesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.routingRulesField))
		if updateField.routingRulesField == routingRulesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.routingRulesField, asTableValues, updateField.routingRulesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.routingRulesField,
				"\"routing_rules\"", updateField.routingRulesField,
				asTableValues, updateField.routingRulesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRoutingRulesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RoutingRulesPrimaryID, asTableValue string) (whereQry string) {
	routingRulesSelectFields := NewRoutingRulesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"routing_rules\".\"id\" = %s.\"id\"::"+GetRoutingRulesFieldType(routingRulesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRoutingRulesFieldType(routingRulesField RoutingRulesField) string {
	selectRoutingRulesFields := NewRoutingRulesSelectFields()
	switch routingRulesField {

	case selectRoutingRulesFields.Id():
		return "uuid"

	case selectRoutingRulesFields.ProfileId():
		return "uuid"

	case selectRoutingRulesFields.Priority():
		return "int2"

	case selectRoutingRulesFields.Name():
		return "text"

	case selectRoutingRulesFields.IsActive():
		return "bool"

	case selectRoutingRulesFields.MatchPaymentMethod():
		return "payment_method_type_enum"

	case selectRoutingRulesFields.MatchCurrency():
		return "payment_currency"

	case selectRoutingRulesFields.MatchAmountMin():
		return "numeric"

	case selectRoutingRulesFields.MatchAmountMax():
		return "numeric"

	case selectRoutingRulesFields.MatchUserCountry():
		return "text"

	case selectRoutingRulesFields.MatchCardBin():
		return "text"

	case selectRoutingRulesFields.MatchProductType():
		return "text"

	case selectRoutingRulesFields.CostWeight():
		return "numeric"

	case selectRoutingRulesFields.Notes():
		return "text"

	case selectRoutingRulesFields.MetaCreatedAt():
		return "timestamptz"

	case selectRoutingRulesFields.MetaCreatedBy():
		return "uuid"

	case selectRoutingRulesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRoutingRulesFields.MetaUpdatedBy():
		return "uuid"

	case selectRoutingRulesFields.MetaDeletedAt():
		return "timestamptz"

	case selectRoutingRulesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRoutingRules(ctx context.Context, routingRules *model.RoutingRules, fieldsInsert ...RoutingRulesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRoutingRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RoutingRulesPrimaryID{
		Id: routingRules.Id,
	}
	exists, err := repo.IsExistRoutingRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRoutingRules] failed checking routingRules whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "routingRules", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRoutingRules([]model.RoutingRules{*routingRules}, fieldsInsert...)
	commandQuery := fmt.Sprintf(routingRulesQueries.insertRoutingRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRoutingRules] failed exec create routingRules query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRoutingRulesByID(ctx context.Context, primaryID model.RoutingRulesPrimaryID) (err error) {
	exists, err := repo.IsExistRoutingRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRoutingRulesByID] failed checking routingRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRoutingRulesCompositePrimaryKeyWhere([]model.RoutingRulesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(routingRulesQueries.deleteRoutingRules + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRoutingRulesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingRulesByFilter(ctx context.Context, filter model.Filter) (result []model.RoutingRulesFilterResult, err error) {
	query, args, err := composeRoutingRulesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingRulesByFilter] failed compose routingRules filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingRulesByFilter] failed get routingRules by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRoutingRulesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRoutingRulesFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultRoutingRulesSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := RoutingRulesFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, RoutingRulesField(filterSelectField))
		}
		selectFields = composeRoutingRulesSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(routingRulesQueries.selectRoutingRules, selectFields)

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

func (repo *RepositoryImpl) IsExistRoutingRulesByID(ctx context.Context, primaryID model.RoutingRulesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRoutingRulesCompositePrimaryKeyWhere([]model.RoutingRulesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", routingRulesQueries.selectCountRoutingRules, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingRulesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingRules(ctx context.Context, selectFields ...RoutingRulesField) (routingRulesList model.RoutingRulesList, err error) {
	var (
		defaultRoutingRulesSelectFields = defaultRoutingRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRoutingRulesSelectFields = composeRoutingRulesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(routingRulesQueries.selectRoutingRules, defaultRoutingRulesSelectFields)

	err = repo.db.Read.Select(&routingRulesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingRules] failed get routingRules list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingRulesByID(ctx context.Context, primaryID model.RoutingRulesPrimaryID, selectFields ...RoutingRulesField) (routingRules model.RoutingRules, err error) {
	var (
		defaultRoutingRulesSelectFields = defaultRoutingRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRoutingRulesSelectFields = composeRoutingRulesSelectFields(selectFields...)
	}
	whereQry, params := composeRoutingRulesCompositePrimaryKeyWhere([]model.RoutingRulesPrimaryID{primaryID})
	query := fmt.Sprintf(routingRulesQueries.selectRoutingRules+" WHERE "+whereQry, defaultRoutingRulesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&routingRules, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("routingRules with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRoutingRulesByID] failed get routingRules")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRoutingRulesByID(ctx context.Context, primaryID model.RoutingRulesPrimaryID, routingRules *model.RoutingRules, routingRulesUpdateFields ...RoutingRulesUpdateField) (err error) {
	exists, err := repo.IsExistRoutingRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRoutingRules] failed checking routingRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if routingRules == nil {
		if len(routingRulesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRoutingRulesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		routingRules = &model.RoutingRules{}
	}
	var (
		defaultRoutingRulesUpdateFields = defaultRoutingRulesUpdateFields(*routingRules)
		tempUpdateField                 RoutingRulesUpdateFieldList
		selectFields                    = NewRoutingRulesSelectFields()
	)
	if len(routingRulesUpdateFields) > 0 {
		for _, updateField := range routingRulesUpdateFields {
			if updateField.routingRulesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRoutingRulesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRoutingRulesCompositePrimaryKeyWhere([]model.RoutingRulesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRoutingRulesCommand(defaultRoutingRulesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(routingRulesQueries.updateRoutingRules+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRoutingRules] error when try to update routingRules by id")
	}
	return err
}

var (
	routingRulesQueries = struct {
		selectRoutingRules      string
		selectCountRoutingRules string
		deleteRoutingRules      string
		updateRoutingRules      string
		insertRoutingRules      string
	}{
		selectRoutingRules:      "SELECT %s FROM \"routing_rules\"",
		selectCountRoutingRules: "SELECT COUNT(\"id\") FROM \"routing_rules\"",
		deleteRoutingRules:      "DELETE FROM \"routing_rules\"",
		updateRoutingRules:      "UPDATE \"routing_rules\" SET %s ",
		insertRoutingRules:      "INSERT INTO \"routing_rules\" %s VALUES %s",
	}
)

type RoutingRulesRepository interface {
	CreateRoutingRules(ctx context.Context, routingRules *model.RoutingRules, fieldsInsert ...RoutingRulesField) error
	BulkCreateRoutingRules(ctx context.Context, routingRulesList []*model.RoutingRules, fieldsInsert ...RoutingRulesField) error
	ResolveRoutingRules(ctx context.Context, selectFields ...RoutingRulesField) (model.RoutingRulesList, error)
	ResolveRoutingRulesByID(ctx context.Context, primaryID model.RoutingRulesPrimaryID, selectFields ...RoutingRulesField) (model.RoutingRules, error)
	UpdateRoutingRulesByID(ctx context.Context, id model.RoutingRulesPrimaryID, routingRules *model.RoutingRules, routingRulesUpdateFields ...RoutingRulesUpdateField) error
	BulkUpdateRoutingRules(ctx context.Context, routingRulesListMap map[model.RoutingRulesPrimaryID]*model.RoutingRules, RoutingRulessMapUpdateFieldsRequest map[model.RoutingRulesPrimaryID]RoutingRulesUpdateFieldList) (err error)
	DeleteRoutingRulesByID(ctx context.Context, id model.RoutingRulesPrimaryID) error
	BulkDeleteRoutingRulesByIDs(ctx context.Context, ids []model.RoutingRulesPrimaryID) error
	ResolveRoutingRulesByFilter(ctx context.Context, filter model.Filter) (result []model.RoutingRulesFilterResult, err error)
	IsExistRoutingRulesByIDs(ctx context.Context, ids []model.RoutingRulesPrimaryID) (exists bool, notFoundIds []model.RoutingRulesPrimaryID, err error)
	IsExistRoutingRulesByID(ctx context.Context, id model.RoutingRulesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
