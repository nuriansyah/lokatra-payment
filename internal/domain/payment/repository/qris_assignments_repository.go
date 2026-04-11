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

func composeInsertFieldsAndParamsQrisAssignments(qrisAssignmentsList []model.QrisAssignments, fieldsInsert ...QrisAssignmentsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewQrisAssignmentsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, qrisAssignments := range qrisAssignmentsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, qrisAssignments.Id)
			case selectField.IntentId():
				args = append(args, qrisAssignments.IntentId)
			case selectField.QrString():
				args = append(args, qrisAssignments.QrString)
			case selectField.QrUrl():
				args = append(args, qrisAssignments.QrUrl)
			case selectField.ExpiresAt():
				args = append(args, qrisAssignments.ExpiresAt)
			case selectField.PaidAt():
				args = append(args, qrisAssignments.PaidAt)
			case selectField.PspTransactionId():
				args = append(args, qrisAssignments.PspTransactionId)
			case selectField.MetaCreatedAt():
				args = append(args, qrisAssignments.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, qrisAssignments.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, qrisAssignments.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, qrisAssignments.MetaUpdatedBy)

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

func composeQrisAssignmentsCompositePrimaryKeyWhere(primaryIDs []model.QrisAssignmentsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"qris_assignments\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultQrisAssignmentsSelectFields() string {
	fields := NewQrisAssignmentsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeQrisAssignmentsSelectFields(selectFields ...QrisAssignmentsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type QrisAssignmentsField string
type QrisAssignmentsFieldList []QrisAssignmentsField

type QrisAssignmentsSelectFields struct {
}

func (ss QrisAssignmentsSelectFields) Id() QrisAssignmentsField {
	return QrisAssignmentsField("id")
}

func (ss QrisAssignmentsSelectFields) IntentId() QrisAssignmentsField {
	return QrisAssignmentsField("intent_id")
}

func (ss QrisAssignmentsSelectFields) QrString() QrisAssignmentsField {
	return QrisAssignmentsField("qr_string")
}

func (ss QrisAssignmentsSelectFields) QrUrl() QrisAssignmentsField {
	return QrisAssignmentsField("qr_url")
}

func (ss QrisAssignmentsSelectFields) ExpiresAt() QrisAssignmentsField {
	return QrisAssignmentsField("expires_at")
}

func (ss QrisAssignmentsSelectFields) PaidAt() QrisAssignmentsField {
	return QrisAssignmentsField("paid_at")
}

func (ss QrisAssignmentsSelectFields) PspTransactionId() QrisAssignmentsField {
	return QrisAssignmentsField("psp_transaction_id")
}

func (ss QrisAssignmentsSelectFields) MetaCreatedAt() QrisAssignmentsField {
	return QrisAssignmentsField("meta_created_at")
}

func (ss QrisAssignmentsSelectFields) MetaCreatedBy() QrisAssignmentsField {
	return QrisAssignmentsField("meta_created_by")
}

func (ss QrisAssignmentsSelectFields) MetaUpdatedAt() QrisAssignmentsField {
	return QrisAssignmentsField("meta_updated_at")
}

func (ss QrisAssignmentsSelectFields) MetaUpdatedBy() QrisAssignmentsField {
	return QrisAssignmentsField("meta_updated_by")
}

func (ss QrisAssignmentsSelectFields) All() QrisAssignmentsFieldList {
	return []QrisAssignmentsField{
		ss.Id(),
		ss.IntentId(),
		ss.QrString(),
		ss.QrUrl(),
		ss.ExpiresAt(),
		ss.PaidAt(),
		ss.PspTransactionId(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
	}
}

func NewQrisAssignmentsSelectFields() QrisAssignmentsSelectFields {
	return QrisAssignmentsSelectFields{}
}

type QrisAssignmentsUpdateFieldOption struct {
	useIncrement bool
}
type QrisAssignmentsUpdateField struct {
	qrisAssignmentsField QrisAssignmentsField
	opt                  QrisAssignmentsUpdateFieldOption
	value                interface{}
}
type QrisAssignmentsUpdateFieldList []QrisAssignmentsUpdateField

func defaultQrisAssignmentsUpdateFieldOption() QrisAssignmentsUpdateFieldOption {
	return QrisAssignmentsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementQrisAssignmentsOption(useIncrement bool) func(*QrisAssignmentsUpdateFieldOption) {
	return func(pcufo *QrisAssignmentsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewQrisAssignmentsUpdateField(field QrisAssignmentsField, val interface{}, opts ...func(*QrisAssignmentsUpdateFieldOption)) QrisAssignmentsUpdateField {
	defaultOpt := defaultQrisAssignmentsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return QrisAssignmentsUpdateField{
		qrisAssignmentsField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultQrisAssignmentsUpdateFields(qrisAssignments model.QrisAssignments) (qrisAssignmentsUpdateFieldList QrisAssignmentsUpdateFieldList) {
	selectFields := NewQrisAssignmentsSelectFields()
	qrisAssignmentsUpdateFieldList = append(qrisAssignmentsUpdateFieldList,
		NewQrisAssignmentsUpdateField(selectFields.Id(), qrisAssignments.Id),
		NewQrisAssignmentsUpdateField(selectFields.IntentId(), qrisAssignments.IntentId),
		NewQrisAssignmentsUpdateField(selectFields.QrString(), qrisAssignments.QrString),
		NewQrisAssignmentsUpdateField(selectFields.QrUrl(), qrisAssignments.QrUrl),
		NewQrisAssignmentsUpdateField(selectFields.ExpiresAt(), qrisAssignments.ExpiresAt),
		NewQrisAssignmentsUpdateField(selectFields.PaidAt(), qrisAssignments.PaidAt),
		NewQrisAssignmentsUpdateField(selectFields.PspTransactionId(), qrisAssignments.PspTransactionId),
		NewQrisAssignmentsUpdateField(selectFields.MetaCreatedAt(), qrisAssignments.MetaCreatedAt),
		NewQrisAssignmentsUpdateField(selectFields.MetaCreatedBy(), qrisAssignments.MetaCreatedBy),
		NewQrisAssignmentsUpdateField(selectFields.MetaUpdatedAt(), qrisAssignments.MetaUpdatedAt),
		NewQrisAssignmentsUpdateField(selectFields.MetaUpdatedBy(), qrisAssignments.MetaUpdatedBy),
	)
	return
}
func composeUpdateFieldsQrisAssignmentsCommand(qrisAssignmentsUpdateFieldList QrisAssignmentsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range qrisAssignmentsUpdateFieldList {
		field := string(updateField.qrisAssignmentsField)
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

func (repo *RepositoryImpl) BulkCreateQrisAssignments(ctx context.Context, qrisAssignmentsList []*model.QrisAssignments, fieldsInsert ...QrisAssignmentsField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.QrisAssignmentsPrimaryID
		qrisAssignmentsValueList []model.QrisAssignments
	)

	if len(fieldsInsert) == 0 {
		selectField := NewQrisAssignmentsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, qrisAssignments := range qrisAssignmentsList {

		primaryIds = append(primaryIds, qrisAssignments.ToQrisAssignmentsPrimaryID())

		qrisAssignmentsValueList = append(qrisAssignmentsValueList, *qrisAssignments)
	}

	_, notFoundIds, err := repo.IsExistQrisAssignmentsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateQrisAssignments] failed checking qrisAssignments whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.QrisAssignmentsPrimaryID{}
		mapNotFoundIds := map[model.QrisAssignmentsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "qrisAssignments", fmt.Sprintf("qrisAssignments with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsQrisAssignments(qrisAssignmentsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(qrisAssignmentsQueries.insertQrisAssignments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateQrisAssignments] failed exec create qrisAssignments query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteQrisAssignmentsByIDs(ctx context.Context, primaryIDs []model.QrisAssignmentsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistQrisAssignmentsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteQrisAssignmentsByIDs] failed checking qrisAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("qrisAssignments with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"qris_assignments\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(qrisAssignmentsQueries.deleteQrisAssignments + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteQrisAssignmentsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteQrisAssignmentsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistQrisAssignmentsByIDs(ctx context.Context, ids []model.QrisAssignmentsPrimaryID) (exists bool, notFoundIds []model.QrisAssignmentsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"qris_assignments\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(qrisAssignmentsQueries.selectQrisAssignments, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistQrisAssignmentsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.QrisAssignmentsPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistQrisAssignmentsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.QrisAssignmentsPrimaryID]bool{}
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

// BulkUpdateQrisAssignments is used to bulk update qrisAssignments, by default it will update all field
// if want to update specific field, then fill qrisAssignmentssMapUpdateFieldsRequest else please fill qrisAssignmentssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateQrisAssignments(ctx context.Context, qrisAssignmentssMap map[model.QrisAssignmentsPrimaryID]*model.QrisAssignments, qrisAssignmentssMapUpdateFieldsRequest map[model.QrisAssignmentsPrimaryID]QrisAssignmentsUpdateFieldList) (err error) {
	if len(qrisAssignmentssMap) == 0 && len(qrisAssignmentssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		qrisAssignmentssMapUpdateField map[model.QrisAssignmentsPrimaryID]QrisAssignmentsUpdateFieldList = map[model.QrisAssignmentsPrimaryID]QrisAssignmentsUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(qrisAssignmentssMap) > 0 {
		for id, qrisAssignments := range qrisAssignmentssMap {
			if qrisAssignments == nil {
				log.Error().Err(err).Msg("[BulkUpdateQrisAssignments] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			qrisAssignmentssMapUpdateField[id] = defaultQrisAssignmentsUpdateFields(*qrisAssignments)
		}
	} else {
		qrisAssignmentssMapUpdateField = qrisAssignmentssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateQrisAssignmentsQuery(qrisAssignmentssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistQrisAssignmentsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateQrisAssignments] failed checking qrisAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("qrisAssignments with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeQrisAssignmentsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"qris_assignments\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateQrisAssignments] failed exec query")
	}
	return
}

type QrisAssignmentsFieldParameter struct {
	param string
	args  []interface{}
}

func NewQrisAssignmentsFieldParameter(param string, args ...interface{}) QrisAssignmentsFieldParameter {
	return QrisAssignmentsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateQrisAssignmentsQuery(mapQrisAssignmentss map[model.QrisAssignmentsPrimaryID]QrisAssignmentsUpdateFieldList, asTableValues string) (primaryIDs []model.QrisAssignmentsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.QrisAssignmentsPrimaryID]map[string]interface{}{}
	qrisAssignmentsSelectFields := NewQrisAssignmentsSelectFields()
	for id, updateFields := range mapQrisAssignmentss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.qrisAssignmentsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapQrisAssignmentss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetQrisAssignmentsFieldType(updateField.qrisAssignmentsField)))
			args = append(args, fields[string(updateField.qrisAssignmentsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.qrisAssignmentsField))
		if updateField.qrisAssignmentsField == qrisAssignmentsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.qrisAssignmentsField, asTableValues, updateField.qrisAssignmentsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.qrisAssignmentsField,
				"\"qris_assignments\"", updateField.qrisAssignmentsField,
				asTableValues, updateField.qrisAssignmentsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeQrisAssignmentsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.QrisAssignmentsPrimaryID, asTableValue string) (whereQry string) {
	qrisAssignmentsSelectFields := NewQrisAssignmentsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"qris_assignments\".\"id\" = %s.\"id\"::"+GetQrisAssignmentsFieldType(qrisAssignmentsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetQrisAssignmentsFieldType(qrisAssignmentsField QrisAssignmentsField) string {
	selectQrisAssignmentsFields := NewQrisAssignmentsSelectFields()
	switch qrisAssignmentsField {

	case selectQrisAssignmentsFields.Id():
		return "uuid"

	case selectQrisAssignmentsFields.IntentId():
		return "uuid"

	case selectQrisAssignmentsFields.QrString():
		return "text"

	case selectQrisAssignmentsFields.QrUrl():
		return "text"

	case selectQrisAssignmentsFields.ExpiresAt():
		return "timestamptz"

	case selectQrisAssignmentsFields.PaidAt():
		return "timestamptz"

	case selectQrisAssignmentsFields.PspTransactionId():
		return "text"

	case selectQrisAssignmentsFields.MetaCreatedAt():
		return "timestamptz"

	case selectQrisAssignmentsFields.MetaCreatedBy():
		return "uuid"

	case selectQrisAssignmentsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectQrisAssignmentsFields.MetaUpdatedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateQrisAssignments(ctx context.Context, qrisAssignments *model.QrisAssignments, fieldsInsert ...QrisAssignmentsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewQrisAssignmentsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.QrisAssignmentsPrimaryID{
		Id: qrisAssignments.Id,
	}
	exists, err := repo.IsExistQrisAssignmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateQrisAssignments] failed checking qrisAssignments whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "qrisAssignments", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsQrisAssignments([]model.QrisAssignments{*qrisAssignments}, fieldsInsert...)
	commandQuery := fmt.Sprintf(qrisAssignmentsQueries.insertQrisAssignments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateQrisAssignments] failed exec create qrisAssignments query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteQrisAssignmentsByID(ctx context.Context, primaryID model.QrisAssignmentsPrimaryID) (err error) {
	exists, err := repo.IsExistQrisAssignmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteQrisAssignmentsByID] failed checking qrisAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("qrisAssignments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeQrisAssignmentsCompositePrimaryKeyWhere([]model.QrisAssignmentsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(qrisAssignmentsQueries.deleteQrisAssignments + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteQrisAssignmentsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveQrisAssignmentsByFilter(ctx context.Context, filter model.Filter) (result []model.QrisAssignmentsFilterResult, err error) {
	query, args, err := composeQrisAssignmentsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveQrisAssignmentsByFilter] failed compose qrisAssignments filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveQrisAssignmentsByFilter] failed get qrisAssignments by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeQrisAssignmentsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateQrisAssignmentsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultQrisAssignmentsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := QrisAssignmentsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, QrisAssignmentsField(filterSelectField))
		}
		selectFields = composeQrisAssignmentsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(qrisAssignmentsQueries.selectQrisAssignments, selectFields)

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

func (repo *RepositoryImpl) IsExistQrisAssignmentsByID(ctx context.Context, primaryID model.QrisAssignmentsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeQrisAssignmentsCompositePrimaryKeyWhere([]model.QrisAssignmentsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", qrisAssignmentsQueries.selectCountQrisAssignments, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistQrisAssignmentsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveQrisAssignments(ctx context.Context, selectFields ...QrisAssignmentsField) (qrisAssignmentsList model.QrisAssignmentsList, err error) {
	var (
		defaultQrisAssignmentsSelectFields = defaultQrisAssignmentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultQrisAssignmentsSelectFields = composeQrisAssignmentsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(qrisAssignmentsQueries.selectQrisAssignments, defaultQrisAssignmentsSelectFields)

	err = repo.db.Read.Select(&qrisAssignmentsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveQrisAssignments] failed get qrisAssignments list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveQrisAssignmentsByID(ctx context.Context, primaryID model.QrisAssignmentsPrimaryID, selectFields ...QrisAssignmentsField) (qrisAssignments model.QrisAssignments, err error) {
	var (
		defaultQrisAssignmentsSelectFields = defaultQrisAssignmentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultQrisAssignmentsSelectFields = composeQrisAssignmentsSelectFields(selectFields...)
	}
	whereQry, params := composeQrisAssignmentsCompositePrimaryKeyWhere([]model.QrisAssignmentsPrimaryID{primaryID})
	query := fmt.Sprintf(qrisAssignmentsQueries.selectQrisAssignments+" WHERE "+whereQry, defaultQrisAssignmentsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&qrisAssignments, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("qrisAssignments with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveQrisAssignmentsByID] failed get qrisAssignments")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateQrisAssignmentsByID(ctx context.Context, primaryID model.QrisAssignmentsPrimaryID, qrisAssignments *model.QrisAssignments, qrisAssignmentsUpdateFields ...QrisAssignmentsUpdateField) (err error) {
	exists, err := repo.IsExistQrisAssignmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateQrisAssignments] failed checking qrisAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("qrisAssignments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if qrisAssignments == nil {
		if len(qrisAssignmentsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateQrisAssignmentsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		qrisAssignments = &model.QrisAssignments{}
	}
	var (
		defaultQrisAssignmentsUpdateFields = defaultQrisAssignmentsUpdateFields(*qrisAssignments)
		tempUpdateField                    QrisAssignmentsUpdateFieldList
		selectFields                       = NewQrisAssignmentsSelectFields()
	)
	if len(qrisAssignmentsUpdateFields) > 0 {
		for _, updateField := range qrisAssignmentsUpdateFields {
			if updateField.qrisAssignmentsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultQrisAssignmentsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeQrisAssignmentsCompositePrimaryKeyWhere([]model.QrisAssignmentsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsQrisAssignmentsCommand(defaultQrisAssignmentsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(qrisAssignmentsQueries.updateQrisAssignments+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateQrisAssignments] error when try to update qrisAssignments by id")
	}
	return err
}

var (
	qrisAssignmentsQueries = struct {
		selectQrisAssignments      string
		selectCountQrisAssignments string
		deleteQrisAssignments      string
		updateQrisAssignments      string
		insertQrisAssignments      string
	}{
		selectQrisAssignments:      "SELECT %s FROM \"qris_assignments\"",
		selectCountQrisAssignments: "SELECT COUNT(\"id\") FROM \"qris_assignments\"",
		deleteQrisAssignments:      "DELETE FROM \"qris_assignments\"",
		updateQrisAssignments:      "UPDATE \"qris_assignments\" SET %s ",
		insertQrisAssignments:      "INSERT INTO \"qris_assignments\" %s VALUES %s",
	}
)

type QrisAssignmentsRepository interface {
	CreateQrisAssignments(ctx context.Context, qrisAssignments *model.QrisAssignments, fieldsInsert ...QrisAssignmentsField) error
	BulkCreateQrisAssignments(ctx context.Context, qrisAssignmentsList []*model.QrisAssignments, fieldsInsert ...QrisAssignmentsField) error
	ResolveQrisAssignments(ctx context.Context, selectFields ...QrisAssignmentsField) (model.QrisAssignmentsList, error)
	ResolveQrisAssignmentsByID(ctx context.Context, primaryID model.QrisAssignmentsPrimaryID, selectFields ...QrisAssignmentsField) (model.QrisAssignments, error)
	UpdateQrisAssignmentsByID(ctx context.Context, id model.QrisAssignmentsPrimaryID, qrisAssignments *model.QrisAssignments, qrisAssignmentsUpdateFields ...QrisAssignmentsUpdateField) error
	BulkUpdateQrisAssignments(ctx context.Context, qrisAssignmentsListMap map[model.QrisAssignmentsPrimaryID]*model.QrisAssignments, QrisAssignmentssMapUpdateFieldsRequest map[model.QrisAssignmentsPrimaryID]QrisAssignmentsUpdateFieldList) (err error)
	DeleteQrisAssignmentsByID(ctx context.Context, id model.QrisAssignmentsPrimaryID) error
	BulkDeleteQrisAssignmentsByIDs(ctx context.Context, ids []model.QrisAssignmentsPrimaryID) error
	ResolveQrisAssignmentsByFilter(ctx context.Context, filter model.Filter) (result []model.QrisAssignmentsFilterResult, err error)
	IsExistQrisAssignmentsByIDs(ctx context.Context, ids []model.QrisAssignmentsPrimaryID) (exists bool, notFoundIds []model.QrisAssignmentsPrimaryID, err error)
	IsExistQrisAssignmentsByID(ctx context.Context, id model.QrisAssignmentsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
