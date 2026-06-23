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

func composeInsertFieldsAndParamsIdempotencyKeys(idempotencyKeysList []model.IdempotencyKeys, fieldsInsert ...IdempotencyKeysField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewIdempotencyKeysSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, idempotencyKeys := range idempotencyKeysList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, idempotencyKeys.Id)
			case selectField.Key():
				args = append(args, idempotencyKeys.Key)
			case selectField.ActorType():
				args = append(args, idempotencyKeys.ActorType)
			case selectField.ActorId():
				args = append(args, idempotencyKeys.ActorId)
			case selectField.RequestHash():
				args = append(args, idempotencyKeys.RequestHash)
			case selectField.Status():
				args = append(args, idempotencyKeys.Status)
			case selectField.ResourceType():
				args = append(args, idempotencyKeys.ResourceType)
			case selectField.ResourceId():
				args = append(args, idempotencyKeys.ResourceId)
			case selectField.ResponseStatus():
				args = append(args, idempotencyKeys.ResponseStatus)
			case selectField.ResponseBody():
				args = append(args, idempotencyKeys.ResponseBody)
			case selectField.LockedUntil():
				args = append(args, idempotencyKeys.LockedUntil)
			case selectField.CompletedAt():
				args = append(args, idempotencyKeys.CompletedAt)
			case selectField.Metadata():
				args = append(args, idempotencyKeys.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, idempotencyKeys.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, idempotencyKeys.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, idempotencyKeys.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, idempotencyKeys.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, idempotencyKeys.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, idempotencyKeys.MetaDeletedBy)

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

func composeIdempotencyKeysCompositePrimaryKeyWhere(primaryIDs []model.IdempotencyKeysPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"idempotency_keys\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultIdempotencyKeysSelectFields() string {
	fields := NewIdempotencyKeysSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeIdempotencyKeysSelectFields(selectFields ...IdempotencyKeysField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type IdempotencyKeysField string
type IdempotencyKeysFieldList []IdempotencyKeysField

type IdempotencyKeysSelectFields struct {
}

func (ss IdempotencyKeysSelectFields) Id() IdempotencyKeysField {
	return IdempotencyKeysField("id")
}

func (ss IdempotencyKeysSelectFields) Key() IdempotencyKeysField {
	return IdempotencyKeysField("key")
}

func (ss IdempotencyKeysSelectFields) ActorType() IdempotencyKeysField {
	return IdempotencyKeysField("actor_type")
}

func (ss IdempotencyKeysSelectFields) ActorId() IdempotencyKeysField {
	return IdempotencyKeysField("actor_id")
}

func (ss IdempotencyKeysSelectFields) RequestHash() IdempotencyKeysField {
	return IdempotencyKeysField("request_hash")
}

func (ss IdempotencyKeysSelectFields) Status() IdempotencyKeysField {
	return IdempotencyKeysField("status")
}

func (ss IdempotencyKeysSelectFields) ResourceType() IdempotencyKeysField {
	return IdempotencyKeysField("resource_type")
}

func (ss IdempotencyKeysSelectFields) ResourceId() IdempotencyKeysField {
	return IdempotencyKeysField("resource_id")
}

func (ss IdempotencyKeysSelectFields) ResponseStatus() IdempotencyKeysField {
	return IdempotencyKeysField("response_status")
}

func (ss IdempotencyKeysSelectFields) ResponseBody() IdempotencyKeysField {
	return IdempotencyKeysField("response_body")
}

func (ss IdempotencyKeysSelectFields) LockedUntil() IdempotencyKeysField {
	return IdempotencyKeysField("locked_until")
}

func (ss IdempotencyKeysSelectFields) CompletedAt() IdempotencyKeysField {
	return IdempotencyKeysField("completed_at")
}

func (ss IdempotencyKeysSelectFields) Metadata() IdempotencyKeysField {
	return IdempotencyKeysField("metadata")
}

func (ss IdempotencyKeysSelectFields) MetaCreatedAt() IdempotencyKeysField {
	return IdempotencyKeysField("meta_created_at")
}

func (ss IdempotencyKeysSelectFields) MetaCreatedBy() IdempotencyKeysField {
	return IdempotencyKeysField("meta_created_by")
}

func (ss IdempotencyKeysSelectFields) MetaUpdatedAt() IdempotencyKeysField {
	return IdempotencyKeysField("meta_updated_at")
}

func (ss IdempotencyKeysSelectFields) MetaUpdatedBy() IdempotencyKeysField {
	return IdempotencyKeysField("meta_updated_by")
}

func (ss IdempotencyKeysSelectFields) MetaDeletedAt() IdempotencyKeysField {
	return IdempotencyKeysField("meta_deleted_at")
}

func (ss IdempotencyKeysSelectFields) MetaDeletedBy() IdempotencyKeysField {
	return IdempotencyKeysField("meta_deleted_by")
}

func (ss IdempotencyKeysSelectFields) All() IdempotencyKeysFieldList {
	return []IdempotencyKeysField{
		ss.Id(),
		ss.Key(),
		ss.ActorType(),
		ss.ActorId(),
		ss.RequestHash(),
		ss.Status(),
		ss.ResourceType(),
		ss.ResourceId(),
		ss.ResponseStatus(),
		ss.ResponseBody(),
		ss.LockedUntil(),
		ss.CompletedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewIdempotencyKeysSelectFields() IdempotencyKeysSelectFields {
	return IdempotencyKeysSelectFields{}
}

type IdempotencyKeysUpdateFieldOption struct {
	useIncrement bool
}
type IdempotencyKeysUpdateField struct {
	idempotencyKeysField IdempotencyKeysField
	opt                  IdempotencyKeysUpdateFieldOption
	value                interface{}
}
type IdempotencyKeysUpdateFieldList []IdempotencyKeysUpdateField

func defaultIdempotencyKeysUpdateFieldOption() IdempotencyKeysUpdateFieldOption {
	return IdempotencyKeysUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementIdempotencyKeysOption(useIncrement bool) func(*IdempotencyKeysUpdateFieldOption) {
	return func(pcufo *IdempotencyKeysUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewIdempotencyKeysUpdateField(field IdempotencyKeysField, val interface{}, opts ...func(*IdempotencyKeysUpdateFieldOption)) IdempotencyKeysUpdateField {
	defaultOpt := defaultIdempotencyKeysUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return IdempotencyKeysUpdateField{
		idempotencyKeysField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultIdempotencyKeysUpdateFields(idempotencyKeys model.IdempotencyKeys) (idempotencyKeysUpdateFieldList IdempotencyKeysUpdateFieldList) {
	selectFields := NewIdempotencyKeysSelectFields()
	idempotencyKeysUpdateFieldList = append(idempotencyKeysUpdateFieldList,
		NewIdempotencyKeysUpdateField(selectFields.Id(), idempotencyKeys.Id),
		NewIdempotencyKeysUpdateField(selectFields.Key(), idempotencyKeys.Key),
		NewIdempotencyKeysUpdateField(selectFields.ActorType(), idempotencyKeys.ActorType),
		NewIdempotencyKeysUpdateField(selectFields.ActorId(), idempotencyKeys.ActorId),
		NewIdempotencyKeysUpdateField(selectFields.RequestHash(), idempotencyKeys.RequestHash),
		NewIdempotencyKeysUpdateField(selectFields.Status(), idempotencyKeys.Status),
		NewIdempotencyKeysUpdateField(selectFields.ResourceType(), idempotencyKeys.ResourceType),
		NewIdempotencyKeysUpdateField(selectFields.ResourceId(), idempotencyKeys.ResourceId),
		NewIdempotencyKeysUpdateField(selectFields.ResponseStatus(), idempotencyKeys.ResponseStatus),
		NewIdempotencyKeysUpdateField(selectFields.ResponseBody(), idempotencyKeys.ResponseBody),
		NewIdempotencyKeysUpdateField(selectFields.LockedUntil(), idempotencyKeys.LockedUntil),
		NewIdempotencyKeysUpdateField(selectFields.CompletedAt(), idempotencyKeys.CompletedAt),
		NewIdempotencyKeysUpdateField(selectFields.Metadata(), idempotencyKeys.Metadata),
		NewIdempotencyKeysUpdateField(selectFields.MetaCreatedAt(), idempotencyKeys.MetaCreatedAt),
		NewIdempotencyKeysUpdateField(selectFields.MetaCreatedBy(), idempotencyKeys.MetaCreatedBy),
		NewIdempotencyKeysUpdateField(selectFields.MetaUpdatedAt(), idempotencyKeys.MetaUpdatedAt),
		NewIdempotencyKeysUpdateField(selectFields.MetaUpdatedBy(), idempotencyKeys.MetaUpdatedBy),
		NewIdempotencyKeysUpdateField(selectFields.MetaDeletedAt(), idempotencyKeys.MetaDeletedAt),
		NewIdempotencyKeysUpdateField(selectFields.MetaDeletedBy(), idempotencyKeys.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsIdempotencyKeysCommand(idempotencyKeysUpdateFieldList IdempotencyKeysUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range idempotencyKeysUpdateFieldList {
		field := string(updateField.idempotencyKeysField)
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

func (repo *RepositoryImpl) BulkCreateIdempotencyKeys(ctx context.Context, idempotencyKeysList []*model.IdempotencyKeys, fieldsInsert ...IdempotencyKeysField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.IdempotencyKeysPrimaryID
		idempotencyKeysValueList []model.IdempotencyKeys
	)

	if len(fieldsInsert) == 0 {
		selectField := NewIdempotencyKeysSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, idempotencyKeys := range idempotencyKeysList {

		primaryIds = append(primaryIds, idempotencyKeys.ToIdempotencyKeysPrimaryID())

		idempotencyKeysValueList = append(idempotencyKeysValueList, *idempotencyKeys)
	}

	_, notFoundIds, err := repo.IsExistIdempotencyKeysByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateIdempotencyKeys] failed checking idempotencyKeys whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.IdempotencyKeysPrimaryID{}
		mapNotFoundIds := map[model.IdempotencyKeysPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "idempotencyKeys", fmt.Sprintf("idempotencyKeys with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsIdempotencyKeys(idempotencyKeysValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(idempotencyKeysQueries.insertIdempotencyKeys, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateIdempotencyKeys] failed exec create idempotencyKeys query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteIdempotencyKeysByIDs(ctx context.Context, primaryIDs []model.IdempotencyKeysPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistIdempotencyKeysByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteIdempotencyKeysByIDs] failed checking idempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("idempotencyKeys with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"idempotency_keys\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := idempotencyKeysQueries.deleteIdempotencyKeys + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteIdempotencyKeysByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteIdempotencyKeysByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistIdempotencyKeysByIDs(ctx context.Context, ids []model.IdempotencyKeysPrimaryID) (exists bool, notFoundIds []model.IdempotencyKeysPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"idempotency_keys\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(idempotencyKeysQueries.selectIdempotencyKeys, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistIdempotencyKeysByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.IdempotencyKeysPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistIdempotencyKeysByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.IdempotencyKeysPrimaryID]bool{}
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

// BulkUpdateIdempotencyKeys is used to bulk update idempotencyKeys, by default it will update all field
// if want to update specific field, then fill idempotencyKeyssMapUpdateFieldsRequest else please fill idempotencyKeyssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateIdempotencyKeys(ctx context.Context, idempotencyKeyssMap map[model.IdempotencyKeysPrimaryID]*model.IdempotencyKeys, idempotencyKeyssMapUpdateFieldsRequest map[model.IdempotencyKeysPrimaryID]IdempotencyKeysUpdateFieldList) (err error) {
	if len(idempotencyKeyssMap) == 0 && len(idempotencyKeyssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		idempotencyKeyssMapUpdateField map[model.IdempotencyKeysPrimaryID]IdempotencyKeysUpdateFieldList = map[model.IdempotencyKeysPrimaryID]IdempotencyKeysUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(idempotencyKeyssMap) > 0 {
		for id, idempotencyKeys := range idempotencyKeyssMap {
			if idempotencyKeys == nil {
				log.Error().Err(err).Msg("[BulkUpdateIdempotencyKeys] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			idempotencyKeyssMapUpdateField[id] = defaultIdempotencyKeysUpdateFields(*idempotencyKeys)
		}
	} else {
		idempotencyKeyssMapUpdateField = idempotencyKeyssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateIdempotencyKeysQuery(idempotencyKeyssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistIdempotencyKeysByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateIdempotencyKeys] failed checking idempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("idempotencyKeys with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeIdempotencyKeysCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"idempotency_keys\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateIdempotencyKeys] failed exec query")
	}
	return
}

type IdempotencyKeysFieldParameter struct {
	param string
	args  []interface{}
}

func NewIdempotencyKeysFieldParameter(param string, args ...interface{}) IdempotencyKeysFieldParameter {
	return IdempotencyKeysFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateIdempotencyKeysQuery(mapIdempotencyKeyss map[model.IdempotencyKeysPrimaryID]IdempotencyKeysUpdateFieldList, asTableValues string) (primaryIDs []model.IdempotencyKeysPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.IdempotencyKeysPrimaryID]map[string]interface{}{}
	idempotencyKeysSelectFields := NewIdempotencyKeysSelectFields()
	for id, updateFields := range mapIdempotencyKeyss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.idempotencyKeysField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapIdempotencyKeyss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetIdempotencyKeysFieldType(updateField.idempotencyKeysField)))
			args = append(args, fields[string(updateField.idempotencyKeysField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.idempotencyKeysField))
		if updateField.idempotencyKeysField == idempotencyKeysSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.idempotencyKeysField, asTableValues, updateField.idempotencyKeysField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.idempotencyKeysField,
				"\"idempotency_keys\"", updateField.idempotencyKeysField,
				asTableValues, updateField.idempotencyKeysField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeIdempotencyKeysCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.IdempotencyKeysPrimaryID, asTableValue string) (whereQry string) {
	idempotencyKeysSelectFields := NewIdempotencyKeysSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"idempotency_keys\".\"id\" = %s.\"id\"::"+GetIdempotencyKeysFieldType(idempotencyKeysSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetIdempotencyKeysFieldType(idempotencyKeysField IdempotencyKeysField) string {
	selectIdempotencyKeysFields := NewIdempotencyKeysSelectFields()
	switch idempotencyKeysField {

	case selectIdempotencyKeysFields.Id():
		return "uuid"

	case selectIdempotencyKeysFields.Key():
		return "text"

	case selectIdempotencyKeysFields.ActorType():
		return "text"

	case selectIdempotencyKeysFields.ActorId():
		return "uuid"

	case selectIdempotencyKeysFields.RequestHash():
		return "text"

	case selectIdempotencyKeysFields.Status():
		return "idempotency_status_enum"

	case selectIdempotencyKeysFields.ResourceType():
		return "text"

	case selectIdempotencyKeysFields.ResourceId():
		return "uuid"

	case selectIdempotencyKeysFields.ResponseStatus():
		return "int4"

	case selectIdempotencyKeysFields.ResponseBody():
		return "jsonb"

	case selectIdempotencyKeysFields.LockedUntil():
		return "timestamptz"

	case selectIdempotencyKeysFields.CompletedAt():
		return "timestamptz"

	case selectIdempotencyKeysFields.Metadata():
		return "jsonb"

	case selectIdempotencyKeysFields.MetaCreatedAt():
		return "timestamptz"

	case selectIdempotencyKeysFields.MetaCreatedBy():
		return "uuid"

	case selectIdempotencyKeysFields.MetaUpdatedAt():
		return "timestamptz"

	case selectIdempotencyKeysFields.MetaUpdatedBy():
		return "uuid"

	case selectIdempotencyKeysFields.MetaDeletedAt():
		return "timestamptz"

	case selectIdempotencyKeysFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateIdempotencyKeys(ctx context.Context, idempotencyKeys *model.IdempotencyKeys, fieldsInsert ...IdempotencyKeysField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewIdempotencyKeysSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.IdempotencyKeysPrimaryID{
		Id: idempotencyKeys.Id,
	}
	exists, err := repo.IsExistIdempotencyKeysByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateIdempotencyKeys] failed checking idempotencyKeys whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "idempotencyKeys", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsIdempotencyKeys([]model.IdempotencyKeys{*idempotencyKeys}, fieldsInsert...)
	commandQuery := fmt.Sprintf(idempotencyKeysQueries.insertIdempotencyKeys, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateIdempotencyKeys] failed exec create idempotencyKeys query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteIdempotencyKeysByID(ctx context.Context, primaryID model.IdempotencyKeysPrimaryID) (err error) {
	exists, err := repo.IsExistIdempotencyKeysByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteIdempotencyKeysByID] failed checking idempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("idempotencyKeys with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeIdempotencyKeysCompositePrimaryKeyWhere([]model.IdempotencyKeysPrimaryID{primaryID})
	commandQuery := idempotencyKeysQueries.deleteIdempotencyKeys + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteIdempotencyKeysByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveIdempotencyKeysByFilter(ctx context.Context, filter model.Filter) (result []model.IdempotencyKeysFilterResult, err error) {
	query, args, err := composeIdempotencyKeysFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveIdempotencyKeysByFilter] failed compose idempotencyKeys filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveIdempotencyKeysByFilter] failed get idempotencyKeys by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeIdempotencyKeysFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.IdempotencyKeysFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeIdempotencyKeysFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeIdempotencyKeysSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeIdempotencyKeysFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewIdempotencyKeysFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 19+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["key"]; !selected {
			selectColumns = append(selectColumns, "base.\"key\"")
			selectedColumns["key"] = struct{}{}
		}
		if _, selected := selectedColumns["actor_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"actor_type\"")
			selectedColumns["actor_type"] = struct{}{}
		}
		if _, selected := selectedColumns["actor_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"actor_id\"")
			selectedColumns["actor_id"] = struct{}{}
		}
		if _, selected := selectedColumns["request_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"request_hash\"")
			selectedColumns["request_hash"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["resource_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"resource_type\"")
			selectedColumns["resource_type"] = struct{}{}
		}
		if _, selected := selectedColumns["resource_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"resource_id\"")
			selectedColumns["resource_id"] = struct{}{}
		}
		if _, selected := selectedColumns["response_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"response_status\"")
			selectedColumns["response_status"] = struct{}{}
		}
		if _, selected := selectedColumns["response_body"]; !selected {
			selectColumns = append(selectColumns, "base.\"response_body\"")
			selectedColumns["response_body"] = struct{}{}
		}
		if _, selected := selectedColumns["locked_until"]; !selected {
			selectColumns = append(selectColumns, "base.\"locked_until\"")
			selectedColumns["locked_until"] = struct{}{}
		}
		if _, selected := selectedColumns["completed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"completed_at\"")
			selectedColumns["completed_at"] = struct{}{}
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

type idempotencyKeysFilterPlaceholder struct {
	index int
}

func (p *idempotencyKeysFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeIdempotencyKeysFilterPredicate(filterField model.FilterField, placeholders *idempotencyKeysFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewIdempotencyKeysFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeIdempotencyKeysFilterSQLExpr(spec)
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

func composeIdempotencyKeysFilterGroup(group model.FilterGroup, placeholders *idempotencyKeysFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeIdempotencyKeysFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeIdempotencyKeysFilterGroup(child, placeholders, args, requiredJoins)
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

func composeIdempotencyKeysFilterWhereQueries(filter model.Filter, placeholders *idempotencyKeysFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeIdempotencyKeysFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeIdempotencyKeysFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeIdempotencyKeysFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateIdempotencyKeysFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeIdempotencyKeysSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeIdempotencyKeysFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := idempotencyKeysFilterPlaceholder{index: 1}
	whereQueries, err := composeIdempotencyKeysFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewIdempotencyKeysFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeIdempotencyKeysFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeIdempotencyKeysSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"idempotency_keys\" base%s", strings.Join(selectColumns, ","), composeIdempotencyKeysFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistIdempotencyKeysByID(ctx context.Context, primaryID model.IdempotencyKeysPrimaryID) (exists bool, err error) {
	whereQuery, params := composeIdempotencyKeysCompositePrimaryKeyWhere([]model.IdempotencyKeysPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", idempotencyKeysQueries.selectCountIdempotencyKeys, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistIdempotencyKeysByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveIdempotencyKeys(ctx context.Context, selectFields ...IdempotencyKeysField) (idempotencyKeysList model.IdempotencyKeysList, err error) {
	var (
		defaultIdempotencyKeysSelectFields = defaultIdempotencyKeysSelectFields()
	)
	if len(selectFields) > 0 {
		defaultIdempotencyKeysSelectFields = composeIdempotencyKeysSelectFields(selectFields...)
	}
	query := fmt.Sprintf(idempotencyKeysQueries.selectIdempotencyKeys, defaultIdempotencyKeysSelectFields)

	err = repo.db.Read.SelectContext(ctx, &idempotencyKeysList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveIdempotencyKeys] failed get idempotencyKeys list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveIdempotencyKeysByID(ctx context.Context, primaryID model.IdempotencyKeysPrimaryID, selectFields ...IdempotencyKeysField) (idempotencyKeys model.IdempotencyKeys, err error) {
	var (
		defaultIdempotencyKeysSelectFields = defaultIdempotencyKeysSelectFields()
	)
	if len(selectFields) > 0 {
		defaultIdempotencyKeysSelectFields = composeIdempotencyKeysSelectFields(selectFields...)
	}
	whereQry, params := composeIdempotencyKeysCompositePrimaryKeyWhere([]model.IdempotencyKeysPrimaryID{primaryID})
	query := fmt.Sprintf(idempotencyKeysQueries.selectIdempotencyKeys+" WHERE "+whereQry, defaultIdempotencyKeysSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &idempotencyKeys, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("idempotencyKeys with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveIdempotencyKeysByID] failed get idempotencyKeys")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateIdempotencyKeysByID(ctx context.Context, primaryID model.IdempotencyKeysPrimaryID, idempotencyKeys *model.IdempotencyKeys, idempotencyKeysUpdateFields ...IdempotencyKeysUpdateField) (err error) {
	exists, err := repo.IsExistIdempotencyKeysByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateIdempotencyKeys] failed checking idempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("idempotencyKeys with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if idempotencyKeys == nil {
		if len(idempotencyKeysUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateIdempotencyKeysByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		idempotencyKeys = &model.IdempotencyKeys{}
	}
	var (
		defaultIdempotencyKeysUpdateFields = defaultIdempotencyKeysUpdateFields(*idempotencyKeys)
		tempUpdateField                    IdempotencyKeysUpdateFieldList
		selectFields                       = NewIdempotencyKeysSelectFields()
	)
	if len(idempotencyKeysUpdateFields) > 0 {
		for _, updateField := range idempotencyKeysUpdateFields {
			if updateField.idempotencyKeysField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultIdempotencyKeysUpdateFields = tempUpdateField
	}
	whereQuery, params := composeIdempotencyKeysCompositePrimaryKeyWhere([]model.IdempotencyKeysPrimaryID{primaryID})
	fields, args := composeUpdateFieldsIdempotencyKeysCommand(defaultIdempotencyKeysUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(idempotencyKeysQueries.updateIdempotencyKeys+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateIdempotencyKeys] error when try to update idempotencyKeys by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateIdempotencyKeysByFilter(ctx context.Context, filter model.Filter, idempotencyKeysUpdateFields ...IdempotencyKeysUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(idempotencyKeysUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields IdempotencyKeysUpdateFieldList
		selectFields = NewIdempotencyKeysSelectFields()
	)
	for _, updateField := range idempotencyKeysUpdateFields {
		if updateField.idempotencyKeysField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsIdempotencyKeysCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := idempotencyKeysFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeIdempotencyKeysFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"idempotency_keys\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateIdempotencyKeysByFilter] error when try to update idempotencyKeys by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateIdempotencyKeysByFilter] failed get rows affected")
	}
	return
}

var (
	idempotencyKeysQueries = struct {
		selectIdempotencyKeys      string
		selectCountIdempotencyKeys string
		deleteIdempotencyKeys      string
		updateIdempotencyKeys      string
		insertIdempotencyKeys      string
	}{
		selectIdempotencyKeys:      "SELECT %s FROM \"idempotency_keys\"",
		selectCountIdempotencyKeys: "SELECT COUNT(\"id\") FROM \"idempotency_keys\"",
		deleteIdempotencyKeys:      "DELETE FROM \"idempotency_keys\"",
		updateIdempotencyKeys:      "UPDATE \"idempotency_keys\" SET %s ",
		insertIdempotencyKeys:      "INSERT INTO \"idempotency_keys\" %s VALUES %s",
	}
)

type IdempotencyKeysRepository interface {
	CreateIdempotencyKeys(ctx context.Context, idempotencyKeys *model.IdempotencyKeys, fieldsInsert ...IdempotencyKeysField) error
	BulkCreateIdempotencyKeys(ctx context.Context, idempotencyKeysList []*model.IdempotencyKeys, fieldsInsert ...IdempotencyKeysField) error
	ResolveIdempotencyKeys(ctx context.Context, selectFields ...IdempotencyKeysField) (model.IdempotencyKeysList, error)
	ResolveIdempotencyKeysByID(ctx context.Context, primaryID model.IdempotencyKeysPrimaryID, selectFields ...IdempotencyKeysField) (model.IdempotencyKeys, error)
	UpdateIdempotencyKeysByID(ctx context.Context, id model.IdempotencyKeysPrimaryID, idempotencyKeys *model.IdempotencyKeys, idempotencyKeysUpdateFields ...IdempotencyKeysUpdateField) error
	UpdateIdempotencyKeysByFilter(ctx context.Context, filter model.Filter, idempotencyKeysUpdateFields ...IdempotencyKeysUpdateField) (rowsAffected int64, err error)
	BulkUpdateIdempotencyKeys(ctx context.Context, idempotencyKeysListMap map[model.IdempotencyKeysPrimaryID]*model.IdempotencyKeys, IdempotencyKeyssMapUpdateFieldsRequest map[model.IdempotencyKeysPrimaryID]IdempotencyKeysUpdateFieldList) (err error)
	DeleteIdempotencyKeysByID(ctx context.Context, id model.IdempotencyKeysPrimaryID) error
	BulkDeleteIdempotencyKeysByIDs(ctx context.Context, ids []model.IdempotencyKeysPrimaryID) error
	ResolveIdempotencyKeysByFilter(ctx context.Context, filter model.Filter) (result []model.IdempotencyKeysFilterResult, err error)
	IsExistIdempotencyKeysByIDs(ctx context.Context, ids []model.IdempotencyKeysPrimaryID) (exists bool, notFoundIds []model.IdempotencyKeysPrimaryID, err error)
	IsExistIdempotencyKeysByID(ctx context.Context, id model.IdempotencyKeysPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
