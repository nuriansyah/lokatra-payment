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

func composeInsertFieldsAndParamsRoutingProfiles(routingProfilesList []model.RoutingProfiles, fieldsInsert ...RoutingProfilesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRoutingProfilesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, routingProfiles := range routingProfilesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, routingProfiles.Id)
			case selectField.MerchantId():
				args = append(args, routingProfiles.MerchantId)
			case selectField.Name():
				args = append(args, routingProfiles.Name)
			case selectField.Strategy():
				args = append(args, routingProfiles.Strategy)
			case selectField.IsActive():
				args = append(args, routingProfiles.IsActive)
			case selectField.FallbackProfileId():
				args = append(args, routingProfiles.FallbackProfileId)
			case selectField.Notes():
				args = append(args, routingProfiles.Notes)
			case selectField.MetaCreatedAt():
				args = append(args, routingProfiles.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, routingProfiles.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, routingProfiles.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, routingProfiles.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, routingProfiles.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, routingProfiles.MetaDeletedBy)

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

func composeRoutingProfilesCompositePrimaryKeyWhere(primaryIDs []model.RoutingProfilesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"routing_profiles\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRoutingProfilesSelectFields() string {
	fields := NewRoutingProfilesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRoutingProfilesSelectFields(selectFields ...RoutingProfilesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RoutingProfilesField string
type RoutingProfilesFieldList []RoutingProfilesField

type RoutingProfilesSelectFields struct {
}

func (ss RoutingProfilesSelectFields) Id() RoutingProfilesField {
	return RoutingProfilesField("id")
}

func (ss RoutingProfilesSelectFields) MerchantId() RoutingProfilesField {
	return RoutingProfilesField("merchant_id")
}

func (ss RoutingProfilesSelectFields) Name() RoutingProfilesField {
	return RoutingProfilesField("name")
}

func (ss RoutingProfilesSelectFields) Strategy() RoutingProfilesField {
	return RoutingProfilesField("strategy")
}

func (ss RoutingProfilesSelectFields) IsActive() RoutingProfilesField {
	return RoutingProfilesField("is_active")
}

func (ss RoutingProfilesSelectFields) FallbackProfileId() RoutingProfilesField {
	return RoutingProfilesField("fallback_profile_id")
}

func (ss RoutingProfilesSelectFields) Notes() RoutingProfilesField {
	return RoutingProfilesField("notes")
}

func (ss RoutingProfilesSelectFields) MetaCreatedAt() RoutingProfilesField {
	return RoutingProfilesField("meta_created_at")
}

func (ss RoutingProfilesSelectFields) MetaCreatedBy() RoutingProfilesField {
	return RoutingProfilesField("meta_created_by")
}

func (ss RoutingProfilesSelectFields) MetaUpdatedAt() RoutingProfilesField {
	return RoutingProfilesField("meta_updated_at")
}

func (ss RoutingProfilesSelectFields) MetaUpdatedBy() RoutingProfilesField {
	return RoutingProfilesField("meta_updated_by")
}

func (ss RoutingProfilesSelectFields) MetaDeletedAt() RoutingProfilesField {
	return RoutingProfilesField("meta_deleted_at")
}

func (ss RoutingProfilesSelectFields) MetaDeletedBy() RoutingProfilesField {
	return RoutingProfilesField("meta_deleted_by")
}

func (ss RoutingProfilesSelectFields) All() RoutingProfilesFieldList {
	return []RoutingProfilesField{
		ss.Id(),
		ss.MerchantId(),
		ss.Name(),
		ss.Strategy(),
		ss.IsActive(),
		ss.FallbackProfileId(),
		ss.Notes(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRoutingProfilesSelectFields() RoutingProfilesSelectFields {
	return RoutingProfilesSelectFields{}
}

type RoutingProfilesUpdateFieldOption struct {
	useIncrement bool
}
type RoutingProfilesUpdateField struct {
	routingProfilesField RoutingProfilesField
	opt                  RoutingProfilesUpdateFieldOption
	value                interface{}
}
type RoutingProfilesUpdateFieldList []RoutingProfilesUpdateField

func defaultRoutingProfilesUpdateFieldOption() RoutingProfilesUpdateFieldOption {
	return RoutingProfilesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRoutingProfilesOption(useIncrement bool) func(*RoutingProfilesUpdateFieldOption) {
	return func(pcufo *RoutingProfilesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRoutingProfilesUpdateField(field RoutingProfilesField, val interface{}, opts ...func(*RoutingProfilesUpdateFieldOption)) RoutingProfilesUpdateField {
	defaultOpt := defaultRoutingProfilesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RoutingProfilesUpdateField{
		routingProfilesField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultRoutingProfilesUpdateFields(routingProfiles model.RoutingProfiles) (routingProfilesUpdateFieldList RoutingProfilesUpdateFieldList) {
	selectFields := NewRoutingProfilesSelectFields()
	routingProfilesUpdateFieldList = append(routingProfilesUpdateFieldList,
		NewRoutingProfilesUpdateField(selectFields.Id(), routingProfiles.Id),
		NewRoutingProfilesUpdateField(selectFields.MerchantId(), routingProfiles.MerchantId),
		NewRoutingProfilesUpdateField(selectFields.Name(), routingProfiles.Name),
		NewRoutingProfilesUpdateField(selectFields.Strategy(), routingProfiles.Strategy),
		NewRoutingProfilesUpdateField(selectFields.IsActive(), routingProfiles.IsActive),
		NewRoutingProfilesUpdateField(selectFields.FallbackProfileId(), routingProfiles.FallbackProfileId),
		NewRoutingProfilesUpdateField(selectFields.Notes(), routingProfiles.Notes),
		NewRoutingProfilesUpdateField(selectFields.MetaCreatedAt(), routingProfiles.MetaCreatedAt),
		NewRoutingProfilesUpdateField(selectFields.MetaCreatedBy(), routingProfiles.MetaCreatedBy),
		NewRoutingProfilesUpdateField(selectFields.MetaUpdatedAt(), routingProfiles.MetaUpdatedAt),
		NewRoutingProfilesUpdateField(selectFields.MetaUpdatedBy(), routingProfiles.MetaUpdatedBy),
		NewRoutingProfilesUpdateField(selectFields.MetaDeletedAt(), routingProfiles.MetaDeletedAt),
		NewRoutingProfilesUpdateField(selectFields.MetaDeletedBy(), routingProfiles.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRoutingProfilesCommand(routingProfilesUpdateFieldList RoutingProfilesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range routingProfilesUpdateFieldList {
		field := string(updateField.routingProfilesField)
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

func (repo *RepositoryImpl) BulkCreateRoutingProfiles(ctx context.Context, routingProfilesList []*model.RoutingProfiles, fieldsInsert ...RoutingProfilesField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.RoutingProfilesPrimaryID
		routingProfilesValueList []model.RoutingProfiles
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRoutingProfilesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, routingProfiles := range routingProfilesList {

		primaryIds = append(primaryIds, routingProfiles.ToRoutingProfilesPrimaryID())

		routingProfilesValueList = append(routingProfilesValueList, *routingProfiles)
	}

	_, notFoundIds, err := repo.IsExistRoutingProfilesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRoutingProfiles] failed checking routingProfiles whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RoutingProfilesPrimaryID{}
		mapNotFoundIds := map[model.RoutingProfilesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "routingProfiles", fmt.Sprintf("routingProfiles with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRoutingProfiles(routingProfilesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(routingProfilesQueries.insertRoutingProfiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRoutingProfiles] failed exec create routingProfiles query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRoutingProfilesByIDs(ctx context.Context, primaryIDs []model.RoutingProfilesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRoutingProfilesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingProfilesByIDs] failed checking routingProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingProfiles with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"routing_profiles\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(routingProfilesQueries.deleteRoutingProfiles + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingProfilesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRoutingProfilesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRoutingProfilesByIDs(ctx context.Context, ids []model.RoutingProfilesPrimaryID) (exists bool, notFoundIds []model.RoutingProfilesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"routing_profiles\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(routingProfilesQueries.selectRoutingProfiles, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingProfilesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RoutingProfilesPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingProfilesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RoutingProfilesPrimaryID]bool{}
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

// BulkUpdateRoutingProfiles is used to bulk update routingProfiles, by default it will update all field
// if want to update specific field, then fill routingProfilessMapUpdateFieldsRequest else please fill routingProfilessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRoutingProfiles(ctx context.Context, routingProfilessMap map[model.RoutingProfilesPrimaryID]*model.RoutingProfiles, routingProfilessMapUpdateFieldsRequest map[model.RoutingProfilesPrimaryID]RoutingProfilesUpdateFieldList) (err error) {
	if len(routingProfilessMap) == 0 && len(routingProfilessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		routingProfilessMapUpdateField map[model.RoutingProfilesPrimaryID]RoutingProfilesUpdateFieldList = map[model.RoutingProfilesPrimaryID]RoutingProfilesUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(routingProfilessMap) > 0 {
		for id, routingProfiles := range routingProfilessMap {
			if routingProfiles == nil {
				log.Error().Err(err).Msg("[BulkUpdateRoutingProfiles] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			routingProfilessMapUpdateField[id] = defaultRoutingProfilesUpdateFields(*routingProfiles)
		}
	} else {
		routingProfilessMapUpdateField = routingProfilessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRoutingProfilesQuery(routingProfilessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRoutingProfilesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRoutingProfiles] failed checking routingProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingProfiles with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRoutingProfilesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"routing_profiles\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRoutingProfiles] failed exec query")
	}
	return
}

type RoutingProfilesFieldParameter struct {
	param string
	args  []interface{}
}

func NewRoutingProfilesFieldParameter(param string, args ...interface{}) RoutingProfilesFieldParameter {
	return RoutingProfilesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRoutingProfilesQuery(mapRoutingProfiless map[model.RoutingProfilesPrimaryID]RoutingProfilesUpdateFieldList, asTableValues string) (primaryIDs []model.RoutingProfilesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RoutingProfilesPrimaryID]map[string]interface{}{}
	routingProfilesSelectFields := NewRoutingProfilesSelectFields()
	for id, updateFields := range mapRoutingProfiless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.routingProfilesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRoutingProfiless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRoutingProfilesFieldType(updateField.routingProfilesField)))
			args = append(args, fields[string(updateField.routingProfilesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.routingProfilesField))
		if updateField.routingProfilesField == routingProfilesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.routingProfilesField, asTableValues, updateField.routingProfilesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.routingProfilesField,
				"\"routing_profiles\"", updateField.routingProfilesField,
				asTableValues, updateField.routingProfilesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRoutingProfilesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RoutingProfilesPrimaryID, asTableValue string) (whereQry string) {
	routingProfilesSelectFields := NewRoutingProfilesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"routing_profiles\".\"id\" = %s.\"id\"::"+GetRoutingProfilesFieldType(routingProfilesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRoutingProfilesFieldType(routingProfilesField RoutingProfilesField) string {
	selectRoutingProfilesFields := NewRoutingProfilesSelectFields()
	switch routingProfilesField {

	case selectRoutingProfilesFields.Id():
		return "uuid"

	case selectRoutingProfilesFields.MerchantId():
		return "uuid"

	case selectRoutingProfilesFields.Name():
		return "text"

	case selectRoutingProfilesFields.Strategy():
		return "routing_strategy_enum"

	case selectRoutingProfilesFields.IsActive():
		return "bool"

	case selectRoutingProfilesFields.FallbackProfileId():
		return "uuid"

	case selectRoutingProfilesFields.Notes():
		return "text"

	case selectRoutingProfilesFields.MetaCreatedAt():
		return "timestamptz"

	case selectRoutingProfilesFields.MetaCreatedBy():
		return "uuid"

	case selectRoutingProfilesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRoutingProfilesFields.MetaUpdatedBy():
		return "uuid"

	case selectRoutingProfilesFields.MetaDeletedAt():
		return "timestamptz"

	case selectRoutingProfilesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRoutingProfiles(ctx context.Context, routingProfiles *model.RoutingProfiles, fieldsInsert ...RoutingProfilesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRoutingProfilesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RoutingProfilesPrimaryID{
		Id: routingProfiles.Id,
	}
	exists, err := repo.IsExistRoutingProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRoutingProfiles] failed checking routingProfiles whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "routingProfiles", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRoutingProfiles([]model.RoutingProfiles{*routingProfiles}, fieldsInsert...)
	commandQuery := fmt.Sprintf(routingProfilesQueries.insertRoutingProfiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRoutingProfiles] failed exec create routingProfiles query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRoutingProfilesByID(ctx context.Context, primaryID model.RoutingProfilesPrimaryID) (err error) {
	exists, err := repo.IsExistRoutingProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRoutingProfilesByID] failed checking routingProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingProfiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRoutingProfilesCompositePrimaryKeyWhere([]model.RoutingProfilesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(routingProfilesQueries.deleteRoutingProfiles + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRoutingProfilesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingProfilesByFilter(ctx context.Context, filter model.Filter) (result []model.RoutingProfilesFilterResult, err error) {
	query, args, err := composeRoutingProfilesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingProfilesByFilter] failed compose routingProfiles filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingProfilesByFilter] failed get routingProfiles by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRoutingProfilesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRoutingProfilesFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultRoutingProfilesSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := RoutingProfilesFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, RoutingProfilesField(filterSelectField))
		}
		selectFields = composeRoutingProfilesSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(routingProfilesQueries.selectRoutingProfiles, selectFields)

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

func (repo *RepositoryImpl) IsExistRoutingProfilesByID(ctx context.Context, primaryID model.RoutingProfilesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRoutingProfilesCompositePrimaryKeyWhere([]model.RoutingProfilesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", routingProfilesQueries.selectCountRoutingProfiles, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRoutingProfilesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingProfiles(ctx context.Context, selectFields ...RoutingProfilesField) (routingProfilesList model.RoutingProfilesList, err error) {
	var (
		defaultRoutingProfilesSelectFields = defaultRoutingProfilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRoutingProfilesSelectFields = composeRoutingProfilesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(routingProfilesQueries.selectRoutingProfiles, defaultRoutingProfilesSelectFields)

	err = repo.db.Read.Select(&routingProfilesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRoutingProfiles] failed get routingProfiles list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRoutingProfilesByID(ctx context.Context, primaryID model.RoutingProfilesPrimaryID, selectFields ...RoutingProfilesField) (routingProfiles model.RoutingProfiles, err error) {
	var (
		defaultRoutingProfilesSelectFields = defaultRoutingProfilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRoutingProfilesSelectFields = composeRoutingProfilesSelectFields(selectFields...)
	}
	whereQry, params := composeRoutingProfilesCompositePrimaryKeyWhere([]model.RoutingProfilesPrimaryID{primaryID})
	query := fmt.Sprintf(routingProfilesQueries.selectRoutingProfiles+" WHERE "+whereQry, defaultRoutingProfilesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&routingProfiles, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("routingProfiles with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRoutingProfilesByID] failed get routingProfiles")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRoutingProfilesByID(ctx context.Context, primaryID model.RoutingProfilesPrimaryID, routingProfiles *model.RoutingProfiles, routingProfilesUpdateFields ...RoutingProfilesUpdateField) (err error) {
	exists, err := repo.IsExistRoutingProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRoutingProfiles] failed checking routingProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("routingProfiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if routingProfiles == nil {
		if len(routingProfilesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRoutingProfilesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		routingProfiles = &model.RoutingProfiles{}
	}
	var (
		defaultRoutingProfilesUpdateFields = defaultRoutingProfilesUpdateFields(*routingProfiles)
		tempUpdateField                    RoutingProfilesUpdateFieldList
		selectFields                       = NewRoutingProfilesSelectFields()
	)
	if len(routingProfilesUpdateFields) > 0 {
		for _, updateField := range routingProfilesUpdateFields {
			if updateField.routingProfilesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRoutingProfilesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRoutingProfilesCompositePrimaryKeyWhere([]model.RoutingProfilesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRoutingProfilesCommand(defaultRoutingProfilesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(routingProfilesQueries.updateRoutingProfiles+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRoutingProfiles] error when try to update routingProfiles by id")
	}
	return err
}

var (
	routingProfilesQueries = struct {
		selectRoutingProfiles      string
		selectCountRoutingProfiles string
		deleteRoutingProfiles      string
		updateRoutingProfiles      string
		insertRoutingProfiles      string
	}{
		selectRoutingProfiles:      "SELECT %s FROM \"routing_profiles\"",
		selectCountRoutingProfiles: "SELECT COUNT(\"id\") FROM \"routing_profiles\"",
		deleteRoutingProfiles:      "DELETE FROM \"routing_profiles\"",
		updateRoutingProfiles:      "UPDATE \"routing_profiles\" SET %s ",
		insertRoutingProfiles:      "INSERT INTO \"routing_profiles\" %s VALUES %s",
	}
)

type RoutingProfilesRepository interface {
	CreateRoutingProfiles(ctx context.Context, routingProfiles *model.RoutingProfiles, fieldsInsert ...RoutingProfilesField) error
	BulkCreateRoutingProfiles(ctx context.Context, routingProfilesList []*model.RoutingProfiles, fieldsInsert ...RoutingProfilesField) error
	ResolveRoutingProfiles(ctx context.Context, selectFields ...RoutingProfilesField) (model.RoutingProfilesList, error)
	ResolveRoutingProfilesByID(ctx context.Context, primaryID model.RoutingProfilesPrimaryID, selectFields ...RoutingProfilesField) (model.RoutingProfiles, error)
	UpdateRoutingProfilesByID(ctx context.Context, id model.RoutingProfilesPrimaryID, routingProfiles *model.RoutingProfiles, routingProfilesUpdateFields ...RoutingProfilesUpdateField) error
	BulkUpdateRoutingProfiles(ctx context.Context, routingProfilesListMap map[model.RoutingProfilesPrimaryID]*model.RoutingProfiles, RoutingProfilessMapUpdateFieldsRequest map[model.RoutingProfilesPrimaryID]RoutingProfilesUpdateFieldList) (err error)
	DeleteRoutingProfilesByID(ctx context.Context, id model.RoutingProfilesPrimaryID) error
	BulkDeleteRoutingProfilesByIDs(ctx context.Context, ids []model.RoutingProfilesPrimaryID) error
	ResolveRoutingProfilesByFilter(ctx context.Context, filter model.Filter) (result []model.RoutingProfilesFilterResult, err error)
	IsExistRoutingProfilesByIDs(ctx context.Context, ids []model.RoutingProfilesPrimaryID) (exists bool, notFoundIds []model.RoutingProfilesPrimaryID, err error)
	IsExistRoutingProfilesByID(ctx context.Context, id model.RoutingProfilesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
