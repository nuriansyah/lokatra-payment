package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/idempotency/model"
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
			case selectField.IdempotencyKey():
				args = append(args, idempotencyKeys.IdempotencyKey)
			case selectField.MerchantId():
				args = append(args, idempotencyKeys.MerchantId)
			case selectField.RequestPath():
				args = append(args, idempotencyKeys.RequestPath)
			case selectField.RequestBodyHash():
				args = append(args, idempotencyKeys.RequestBodyHash)
			case selectField.ResponseStatus():
				args = append(args, idempotencyKeys.ResponseStatus)
			case selectField.ResponseBody():
				args = append(args, idempotencyKeys.ResponseBody)
			case selectField.LockedAt():
				args = append(args, idempotencyKeys.LockedAt)
			case selectField.LockedUntil():
				args = append(args, idempotencyKeys.LockedUntil)
			case selectField.CompletedAt():
				args = append(args, idempotencyKeys.CompletedAt)
			case selectField.ExpiresAt():
				args = append(args, idempotencyKeys.ExpiresAt)
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

func (ss IdempotencyKeysSelectFields) IdempotencyKey() IdempotencyKeysField {
	return IdempotencyKeysField("idempotency_key")
}

func (ss IdempotencyKeysSelectFields) MerchantId() IdempotencyKeysField {
	return IdempotencyKeysField("merchant_id")
}

func (ss IdempotencyKeysSelectFields) RequestPath() IdempotencyKeysField {
	return IdempotencyKeysField("request_path")
}

func (ss IdempotencyKeysSelectFields) RequestBodyHash() IdempotencyKeysField {
	return IdempotencyKeysField("request_body_hash")
}

func (ss IdempotencyKeysSelectFields) ResponseStatus() IdempotencyKeysField {
	return IdempotencyKeysField("response_status")
}

func (ss IdempotencyKeysSelectFields) ResponseBody() IdempotencyKeysField {
	return IdempotencyKeysField("response_body")
}

func (ss IdempotencyKeysSelectFields) LockedAt() IdempotencyKeysField {
	return IdempotencyKeysField("locked_at")
}

func (ss IdempotencyKeysSelectFields) LockedUntil() IdempotencyKeysField {
	return IdempotencyKeysField("locked_until")
}

func (ss IdempotencyKeysSelectFields) CompletedAt() IdempotencyKeysField {
	return IdempotencyKeysField("completed_at")
}

func (ss IdempotencyKeysSelectFields) ExpiresAt() IdempotencyKeysField {
	return IdempotencyKeysField("expires_at")
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
		ss.IdempotencyKey(),
		ss.MerchantId(),
		ss.RequestPath(),
		ss.RequestBodyHash(),
		ss.ResponseStatus(),
		ss.ResponseBody(),
		ss.LockedAt(),
		ss.LockedUntil(),
		ss.CompletedAt(),
		ss.ExpiresAt(),
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
		NewIdempotencyKeysUpdateField(selectFields.IdempotencyKey(), idempotencyKeys.IdempotencyKey),
		NewIdempotencyKeysUpdateField(selectFields.MerchantId(), idempotencyKeys.MerchantId),
		NewIdempotencyKeysUpdateField(selectFields.RequestPath(), idempotencyKeys.RequestPath),
		NewIdempotencyKeysUpdateField(selectFields.RequestBodyHash(), idempotencyKeys.RequestBodyHash),
		NewIdempotencyKeysUpdateField(selectFields.ResponseStatus(), idempotencyKeys.ResponseStatus),
		NewIdempotencyKeysUpdateField(selectFields.ResponseBody(), idempotencyKeys.ResponseBody),
		NewIdempotencyKeysUpdateField(selectFields.LockedAt(), idempotencyKeys.LockedAt),
		NewIdempotencyKeysUpdateField(selectFields.LockedUntil(), idempotencyKeys.LockedUntil),
		NewIdempotencyKeysUpdateField(selectFields.CompletedAt(), idempotencyKeys.CompletedAt),
		NewIdempotencyKeysUpdateField(selectFields.ExpiresAt(), idempotencyKeys.ExpiresAt),
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

	commandQuery := fmt.Sprintf(idempotencyKeysQueries.deleteIdempotencyKeys + " WHERE " + whereQuery)

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
	err = repo.db.Read.Select(&resIds, query, params...)
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

	case selectIdempotencyKeysFields.IdempotencyKey():
		return "text"

	case selectIdempotencyKeysFields.MerchantId():
		return "uuid"

	case selectIdempotencyKeysFields.RequestPath():
		return "text"

	case selectIdempotencyKeysFields.RequestBodyHash():
		return "text"

	case selectIdempotencyKeysFields.ResponseStatus():
		return "int2"

	case selectIdempotencyKeysFields.ResponseBody():
		return "jsonb"

	case selectIdempotencyKeysFields.LockedAt():
		return "timestamptz"

	case selectIdempotencyKeysFields.LockedUntil():
		return "timestamptz"

	case selectIdempotencyKeysFields.CompletedAt():
		return "timestamptz"

	case selectIdempotencyKeysFields.ExpiresAt():
		return "timestamptz"

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
	commandQuery := fmt.Sprintf(idempotencyKeysQueries.deleteIdempotencyKeys + " WHERE " + whereQuery)
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
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveIdempotencyKeysByFilter] failed get idempotencyKeys by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeIdempotencyKeysFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateIdempotencyKeysFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultIdempotencyKeysSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := IdempotencyKeysFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, IdempotencyKeysField(filterSelectField))
		}
		selectFields = composeIdempotencyKeysSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(idempotencyKeysQueries.selectIdempotencyKeys, selectFields)

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

func (repo *RepositoryImpl) IsExistIdempotencyKeysByID(ctx context.Context, primaryID model.IdempotencyKeysPrimaryID) (exists bool, err error) {
	whereQuery, params := composeIdempotencyKeysCompositePrimaryKeyWhere([]model.IdempotencyKeysPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", idempotencyKeysQueries.selectCountIdempotencyKeys, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
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

	err = repo.db.Read.Select(&idempotencyKeysList, query)
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
	err = repo.db.Read.Get(&idempotencyKeys, query, params...)
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
	ClaimExecuteIdempotency(ctx context.Context, idempotencyKey string, merchantID uuid.UUID, requestPath string, requestBodyHash string, actorID uuid.UUID, lockUntil time.Time) (record model.IdempotencyKeys, claimed bool, err error)
	CompleteExecuteIdempotency(ctx context.Context, id uuid.UUID, responseStatus int, responseBody json.RawMessage, actorID uuid.UUID) (updated bool, err error)
	CreateIdempotencyKeys(ctx context.Context, idempotencyKeys *model.IdempotencyKeys, fieldsInsert ...IdempotencyKeysField) error
	BulkCreateIdempotencyKeys(ctx context.Context, idempotencyKeysList []*model.IdempotencyKeys, fieldsInsert ...IdempotencyKeysField) error
	ResolveIdempotencyKeys(ctx context.Context, selectFields ...IdempotencyKeysField) (model.IdempotencyKeysList, error)
	ResolveIdempotencyKeysByID(ctx context.Context, primaryID model.IdempotencyKeysPrimaryID, selectFields ...IdempotencyKeysField) (model.IdempotencyKeys, error)
	UpdateIdempotencyKeysByID(ctx context.Context, id model.IdempotencyKeysPrimaryID, idempotencyKeys *model.IdempotencyKeys, idempotencyKeysUpdateFields ...IdempotencyKeysUpdateField) error
	BulkUpdateIdempotencyKeys(ctx context.Context, idempotencyKeysListMap map[model.IdempotencyKeysPrimaryID]*model.IdempotencyKeys, IdempotencyKeyssMapUpdateFieldsRequest map[model.IdempotencyKeysPrimaryID]IdempotencyKeysUpdateFieldList) (err error)
	DeleteIdempotencyKeysByID(ctx context.Context, id model.IdempotencyKeysPrimaryID) error
	BulkDeleteIdempotencyKeysByIDs(ctx context.Context, ids []model.IdempotencyKeysPrimaryID) error
	ResolveIdempotencyKeysByFilter(ctx context.Context, filter model.Filter) (result []model.IdempotencyKeysFilterResult, err error)
	IsExistIdempotencyKeysByIDs(ctx context.Context, ids []model.IdempotencyKeysPrimaryID) (exists bool, notFoundIds []model.IdempotencyKeysPrimaryID, err error)
	IsExistIdempotencyKeysByID(ctx context.Context, id model.IdempotencyKeysPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
