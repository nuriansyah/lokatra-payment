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

func composeInsertFieldsAndParamsFinanceIdempotencyKeys(financeIdempotencyKeysList []model.FinanceIdempotencyKeys, fieldsInsert ...FinanceIdempotencyKeysField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceIdempotencyKeysSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeIdempotencyKeys := range financeIdempotencyKeysList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeIdempotencyKeys.Id)
			case selectField.Scope():
				args = append(args, financeIdempotencyKeys.Scope)
			case selectField.Operation():
				args = append(args, financeIdempotencyKeys.Operation)
			case selectField.IdempotencyKey():
				args = append(args, financeIdempotencyKeys.IdempotencyKey)
			case selectField.RequestHash():
				args = append(args, financeIdempotencyKeys.RequestHash)
			case selectField.ResourceType():
				args = append(args, financeIdempotencyKeys.ResourceType)
			case selectField.ResourceId():
				args = append(args, financeIdempotencyKeys.ResourceId)
			case selectField.ResponseStatus():
				args = append(args, financeIdempotencyKeys.ResponseStatus)
			case selectField.ResponseBody():
				args = append(args, financeIdempotencyKeys.ResponseBody)
			case selectField.LockedUntil():
				args = append(args, financeIdempotencyKeys.LockedUntil)
			case selectField.CompletedAt():
				args = append(args, financeIdempotencyKeys.CompletedAt)
			case selectField.Metadata():
				args = append(args, financeIdempotencyKeys.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeIdempotencyKeys.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeIdempotencyKeys.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeIdempotencyKeys.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeIdempotencyKeys.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeIdempotencyKeys.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeIdempotencyKeys.MetaDeletedBy)

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

func composeFinanceIdempotencyKeysCompositePrimaryKeyWhere(primaryIDs []model.FinanceIdempotencyKeysPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_idempotency_keys\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceIdempotencyKeysSelectFields() string {
	fields := NewFinanceIdempotencyKeysSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceIdempotencyKeysSelectFields(selectFields ...FinanceIdempotencyKeysField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceIdempotencyKeysField string
type FinanceIdempotencyKeysFieldList []FinanceIdempotencyKeysField

type FinanceIdempotencyKeysSelectFields struct {
}

func (ss FinanceIdempotencyKeysSelectFields) Id() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("id")
}

func (ss FinanceIdempotencyKeysSelectFields) Scope() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("scope")
}

func (ss FinanceIdempotencyKeysSelectFields) Operation() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("operation")
}

func (ss FinanceIdempotencyKeysSelectFields) IdempotencyKey() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("idempotency_key")
}

func (ss FinanceIdempotencyKeysSelectFields) RequestHash() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("request_hash")
}

func (ss FinanceIdempotencyKeysSelectFields) ResourceType() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("resource_type")
}

func (ss FinanceIdempotencyKeysSelectFields) ResourceId() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("resource_id")
}

func (ss FinanceIdempotencyKeysSelectFields) ResponseStatus() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("response_status")
}

func (ss FinanceIdempotencyKeysSelectFields) ResponseBody() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("response_body")
}

func (ss FinanceIdempotencyKeysSelectFields) LockedUntil() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("locked_until")
}

func (ss FinanceIdempotencyKeysSelectFields) CompletedAt() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("completed_at")
}

func (ss FinanceIdempotencyKeysSelectFields) Metadata() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("metadata")
}

func (ss FinanceIdempotencyKeysSelectFields) MetaCreatedAt() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("meta_created_at")
}

func (ss FinanceIdempotencyKeysSelectFields) MetaCreatedBy() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("meta_created_by")
}

func (ss FinanceIdempotencyKeysSelectFields) MetaUpdatedAt() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("meta_updated_at")
}

func (ss FinanceIdempotencyKeysSelectFields) MetaUpdatedBy() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("meta_updated_by")
}

func (ss FinanceIdempotencyKeysSelectFields) MetaDeletedAt() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("meta_deleted_at")
}

func (ss FinanceIdempotencyKeysSelectFields) MetaDeletedBy() FinanceIdempotencyKeysField {
	return FinanceIdempotencyKeysField("meta_deleted_by")
}

func (ss FinanceIdempotencyKeysSelectFields) All() FinanceIdempotencyKeysFieldList {
	return []FinanceIdempotencyKeysField{
		ss.Id(),
		ss.Scope(),
		ss.Operation(),
		ss.IdempotencyKey(),
		ss.RequestHash(),
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

func NewFinanceIdempotencyKeysSelectFields() FinanceIdempotencyKeysSelectFields {
	return FinanceIdempotencyKeysSelectFields{}
}

type FinanceIdempotencyKeysUpdateFieldOption struct {
	useIncrement bool
}
type FinanceIdempotencyKeysUpdateField struct {
	financeIdempotencyKeysField FinanceIdempotencyKeysField
	opt                         FinanceIdempotencyKeysUpdateFieldOption
	value                       interface{}
}
type FinanceIdempotencyKeysUpdateFieldList []FinanceIdempotencyKeysUpdateField

func defaultFinanceIdempotencyKeysUpdateFieldOption() FinanceIdempotencyKeysUpdateFieldOption {
	return FinanceIdempotencyKeysUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceIdempotencyKeysOption(useIncrement bool) func(*FinanceIdempotencyKeysUpdateFieldOption) {
	return func(pcufo *FinanceIdempotencyKeysUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceIdempotencyKeysUpdateField(field FinanceIdempotencyKeysField, val interface{}, opts ...func(*FinanceIdempotencyKeysUpdateFieldOption)) FinanceIdempotencyKeysUpdateField {
	defaultOpt := defaultFinanceIdempotencyKeysUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceIdempotencyKeysUpdateField{
		financeIdempotencyKeysField: field,
		value:                       val,
		opt:                         defaultOpt,
	}
}
func defaultFinanceIdempotencyKeysUpdateFields(financeIdempotencyKeys model.FinanceIdempotencyKeys) (financeIdempotencyKeysUpdateFieldList FinanceIdempotencyKeysUpdateFieldList) {
	selectFields := NewFinanceIdempotencyKeysSelectFields()
	financeIdempotencyKeysUpdateFieldList = append(financeIdempotencyKeysUpdateFieldList,
		NewFinanceIdempotencyKeysUpdateField(selectFields.Id(), financeIdempotencyKeys.Id),
		NewFinanceIdempotencyKeysUpdateField(selectFields.Scope(), financeIdempotencyKeys.Scope),
		NewFinanceIdempotencyKeysUpdateField(selectFields.Operation(), financeIdempotencyKeys.Operation),
		NewFinanceIdempotencyKeysUpdateField(selectFields.IdempotencyKey(), financeIdempotencyKeys.IdempotencyKey),
		NewFinanceIdempotencyKeysUpdateField(selectFields.RequestHash(), financeIdempotencyKeys.RequestHash),
		NewFinanceIdempotencyKeysUpdateField(selectFields.ResourceType(), financeIdempotencyKeys.ResourceType),
		NewFinanceIdempotencyKeysUpdateField(selectFields.ResourceId(), financeIdempotencyKeys.ResourceId),
		NewFinanceIdempotencyKeysUpdateField(selectFields.ResponseStatus(), financeIdempotencyKeys.ResponseStatus),
		NewFinanceIdempotencyKeysUpdateField(selectFields.ResponseBody(), financeIdempotencyKeys.ResponseBody),
		NewFinanceIdempotencyKeysUpdateField(selectFields.LockedUntil(), financeIdempotencyKeys.LockedUntil),
		NewFinanceIdempotencyKeysUpdateField(selectFields.CompletedAt(), financeIdempotencyKeys.CompletedAt),
		NewFinanceIdempotencyKeysUpdateField(selectFields.Metadata(), financeIdempotencyKeys.Metadata),
		NewFinanceIdempotencyKeysUpdateField(selectFields.MetaCreatedAt(), financeIdempotencyKeys.MetaCreatedAt),
		NewFinanceIdempotencyKeysUpdateField(selectFields.MetaCreatedBy(), financeIdempotencyKeys.MetaCreatedBy),
		NewFinanceIdempotencyKeysUpdateField(selectFields.MetaUpdatedAt(), financeIdempotencyKeys.MetaUpdatedAt),
		NewFinanceIdempotencyKeysUpdateField(selectFields.MetaUpdatedBy(), financeIdempotencyKeys.MetaUpdatedBy),
		NewFinanceIdempotencyKeysUpdateField(selectFields.MetaDeletedAt(), financeIdempotencyKeys.MetaDeletedAt),
		NewFinanceIdempotencyKeysUpdateField(selectFields.MetaDeletedBy(), financeIdempotencyKeys.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceIdempotencyKeysCommand(financeIdempotencyKeysUpdateFieldList FinanceIdempotencyKeysUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeIdempotencyKeysUpdateFieldList {
		field := string(updateField.financeIdempotencyKeysField)
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

func (repo *RepositoryImpl) BulkCreateFinanceIdempotencyKeys(ctx context.Context, financeIdempotencyKeysList []*model.FinanceIdempotencyKeys, fieldsInsert ...FinanceIdempotencyKeysField) (err error) {
	var (
		fieldsStr                       string
		valueListStr                    []string
		argsList                        []interface{}
		primaryIds                      []model.FinanceIdempotencyKeysPrimaryID
		financeIdempotencyKeysValueList []model.FinanceIdempotencyKeys
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceIdempotencyKeysSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeIdempotencyKeys := range financeIdempotencyKeysList {

		primaryIds = append(primaryIds, financeIdempotencyKeys.ToFinanceIdempotencyKeysPrimaryID())

		financeIdempotencyKeysValueList = append(financeIdempotencyKeysValueList, *financeIdempotencyKeys)
	}

	_, notFoundIds, err := repo.IsExistFinanceIdempotencyKeysByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceIdempotencyKeys] failed checking financeIdempotencyKeys whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceIdempotencyKeysPrimaryID{}
		mapNotFoundIds := map[model.FinanceIdempotencyKeysPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeIdempotencyKeys", fmt.Sprintf("financeIdempotencyKeys with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceIdempotencyKeys(financeIdempotencyKeysValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeIdempotencyKeysQueries.insertFinanceIdempotencyKeys, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceIdempotencyKeys] failed exec create financeIdempotencyKeys query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceIdempotencyKeysByIDs(ctx context.Context, primaryIDs []model.FinanceIdempotencyKeysPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceIdempotencyKeysByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceIdempotencyKeysByIDs] failed checking financeIdempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeIdempotencyKeys with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_idempotency_keys\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeIdempotencyKeysQueries.deleteFinanceIdempotencyKeys + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceIdempotencyKeysByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceIdempotencyKeysByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceIdempotencyKeysByIDs(ctx context.Context, ids []model.FinanceIdempotencyKeysPrimaryID) (exists bool, notFoundIds []model.FinanceIdempotencyKeysPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_idempotency_keys\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeIdempotencyKeysQueries.selectFinanceIdempotencyKeys, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceIdempotencyKeysByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceIdempotencyKeysPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceIdempotencyKeysByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceIdempotencyKeysPrimaryID]bool{}
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

// BulkUpdateFinanceIdempotencyKeys is used to bulk update financeIdempotencyKeys, by default it will update all field
// if want to update specific field, then fill financeIdempotencyKeyssMapUpdateFieldsRequest else please fill financeIdempotencyKeyssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceIdempotencyKeys(ctx context.Context, financeIdempotencyKeyssMap map[model.FinanceIdempotencyKeysPrimaryID]*model.FinanceIdempotencyKeys, financeIdempotencyKeyssMapUpdateFieldsRequest map[model.FinanceIdempotencyKeysPrimaryID]FinanceIdempotencyKeysUpdateFieldList) (err error) {
	if len(financeIdempotencyKeyssMap) == 0 && len(financeIdempotencyKeyssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeIdempotencyKeyssMapUpdateField map[model.FinanceIdempotencyKeysPrimaryID]FinanceIdempotencyKeysUpdateFieldList = map[model.FinanceIdempotencyKeysPrimaryID]FinanceIdempotencyKeysUpdateFieldList{}
		asTableValues                         string                                                                          = "myvalues"
	)

	if len(financeIdempotencyKeyssMap) > 0 {
		for id, financeIdempotencyKeys := range financeIdempotencyKeyssMap {
			if financeIdempotencyKeys == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceIdempotencyKeys] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeIdempotencyKeyssMapUpdateField[id] = defaultFinanceIdempotencyKeysUpdateFields(*financeIdempotencyKeys)
		}
	} else {
		financeIdempotencyKeyssMapUpdateField = financeIdempotencyKeyssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceIdempotencyKeysQuery(financeIdempotencyKeyssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceIdempotencyKeysByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceIdempotencyKeys] failed checking financeIdempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeIdempotencyKeys with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceIdempotencyKeysCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_idempotency_keys\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceIdempotencyKeys] failed exec query")
	}
	return
}

type FinanceIdempotencyKeysFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceIdempotencyKeysFieldParameter(param string, args ...interface{}) FinanceIdempotencyKeysFieldParameter {
	return FinanceIdempotencyKeysFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceIdempotencyKeysQuery(mapFinanceIdempotencyKeyss map[model.FinanceIdempotencyKeysPrimaryID]FinanceIdempotencyKeysUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceIdempotencyKeysPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceIdempotencyKeysPrimaryID]map[string]interface{}{}
	financeIdempotencyKeysSelectFields := NewFinanceIdempotencyKeysSelectFields()
	for id, updateFields := range mapFinanceIdempotencyKeyss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeIdempotencyKeysField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceIdempotencyKeyss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceIdempotencyKeysFieldType(updateField.financeIdempotencyKeysField)))
			args = append(args, fields[string(updateField.financeIdempotencyKeysField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeIdempotencyKeysField))
		if updateField.financeIdempotencyKeysField == financeIdempotencyKeysSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeIdempotencyKeysField, asTableValues, updateField.financeIdempotencyKeysField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeIdempotencyKeysField,
				"\"finance_idempotency_keys\"", updateField.financeIdempotencyKeysField,
				asTableValues, updateField.financeIdempotencyKeysField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceIdempotencyKeysCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceIdempotencyKeysPrimaryID, asTableValue string) (whereQry string) {
	financeIdempotencyKeysSelectFields := NewFinanceIdempotencyKeysSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_idempotency_keys\".\"id\" = %s.\"id\"::"+GetFinanceIdempotencyKeysFieldType(financeIdempotencyKeysSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceIdempotencyKeysFieldType(financeIdempotencyKeysField FinanceIdempotencyKeysField) string {
	selectFinanceIdempotencyKeysFields := NewFinanceIdempotencyKeysSelectFields()
	switch financeIdempotencyKeysField {

	case selectFinanceIdempotencyKeysFields.Id():
		return "uuid"

	case selectFinanceIdempotencyKeysFields.Scope():
		return "text"

	case selectFinanceIdempotencyKeysFields.Operation():
		return "text"

	case selectFinanceIdempotencyKeysFields.IdempotencyKey():
		return "text"

	case selectFinanceIdempotencyKeysFields.RequestHash():
		return "text"

	case selectFinanceIdempotencyKeysFields.ResourceType():
		return "text"

	case selectFinanceIdempotencyKeysFields.ResourceId():
		return "uuid"

	case selectFinanceIdempotencyKeysFields.ResponseStatus():
		return "int4"

	case selectFinanceIdempotencyKeysFields.ResponseBody():
		return "jsonb"

	case selectFinanceIdempotencyKeysFields.LockedUntil():
		return "timestamptz"

	case selectFinanceIdempotencyKeysFields.CompletedAt():
		return "timestamptz"

	case selectFinanceIdempotencyKeysFields.Metadata():
		return "jsonb"

	case selectFinanceIdempotencyKeysFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceIdempotencyKeysFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceIdempotencyKeysFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceIdempotencyKeysFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceIdempotencyKeysFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceIdempotencyKeysFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceIdempotencyKeys(ctx context.Context, financeIdempotencyKeys *model.FinanceIdempotencyKeys, fieldsInsert ...FinanceIdempotencyKeysField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceIdempotencyKeysSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceIdempotencyKeysPrimaryID{
		Id: financeIdempotencyKeys.Id,
	}
	exists, err := repo.IsExistFinanceIdempotencyKeysByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceIdempotencyKeys] failed checking financeIdempotencyKeys whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeIdempotencyKeys", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceIdempotencyKeys([]model.FinanceIdempotencyKeys{*financeIdempotencyKeys}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeIdempotencyKeysQueries.insertFinanceIdempotencyKeys, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceIdempotencyKeys] failed exec create financeIdempotencyKeys query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceIdempotencyKeysByID(ctx context.Context, primaryID model.FinanceIdempotencyKeysPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceIdempotencyKeysByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceIdempotencyKeysByID] failed checking financeIdempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeIdempotencyKeys with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceIdempotencyKeysCompositePrimaryKeyWhere([]model.FinanceIdempotencyKeysPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeIdempotencyKeysQueries.deleteFinanceIdempotencyKeys + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceIdempotencyKeysByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceIdempotencyKeysByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceIdempotencyKeysFilterResult, err error) {
	query, args, err := composeFinanceIdempotencyKeysFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceIdempotencyKeysByFilter] failed compose financeIdempotencyKeys filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceIdempotencyKeysByFilter] failed get financeIdempotencyKeys by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceIdempotencyKeysFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceIdempotencyKeysFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceIdempotencyKeysFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceIdempotencyKeysSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceIdempotencyKeysFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceIdempotencyKeysFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 18+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["scope"]; !selected {
			selectColumns = append(selectColumns, "base.\"scope\"")
			selectedColumns["scope"] = struct{}{}
		}
		if _, selected := selectedColumns["operation"]; !selected {
			selectColumns = append(selectColumns, "base.\"operation\"")
			selectedColumns["operation"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["request_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"request_hash\"")
			selectedColumns["request_hash"] = struct{}{}
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

type financeIdempotencyKeysFilterPlaceholder struct {
	index int
}

func (p *financeIdempotencyKeysFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceIdempotencyKeysFilterPredicate(filterField model.FilterField, placeholders *financeIdempotencyKeysFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceIdempotencyKeysFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceIdempotencyKeysFilterSQLExpr(spec)
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

func composeFinanceIdempotencyKeysFilterGroup(group model.FilterGroup, placeholders *financeIdempotencyKeysFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceIdempotencyKeysFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceIdempotencyKeysFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceIdempotencyKeysFilterWhereQueries(filter model.Filter, placeholders *financeIdempotencyKeysFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceIdempotencyKeysFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceIdempotencyKeysFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceIdempotencyKeysFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceIdempotencyKeysFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceIdempotencyKeysSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceIdempotencyKeysFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeIdempotencyKeysFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceIdempotencyKeysFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceIdempotencyKeysFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceIdempotencyKeysFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceIdempotencyKeysSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_idempotency_keys\" base%s", strings.Join(selectColumns, ","), composeFinanceIdempotencyKeysFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceIdempotencyKeysByID(ctx context.Context, primaryID model.FinanceIdempotencyKeysPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceIdempotencyKeysCompositePrimaryKeyWhere([]model.FinanceIdempotencyKeysPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeIdempotencyKeysQueries.selectCountFinanceIdempotencyKeys, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceIdempotencyKeysByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceIdempotencyKeys(ctx context.Context, selectFields ...FinanceIdempotencyKeysField) (financeIdempotencyKeysList model.FinanceIdempotencyKeysList, err error) {
	var (
		defaultFinanceIdempotencyKeysSelectFields = defaultFinanceIdempotencyKeysSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceIdempotencyKeysSelectFields = composeFinanceIdempotencyKeysSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeIdempotencyKeysQueries.selectFinanceIdempotencyKeys, defaultFinanceIdempotencyKeysSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeIdempotencyKeysList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceIdempotencyKeys] failed get financeIdempotencyKeys list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceIdempotencyKeysByID(ctx context.Context, primaryID model.FinanceIdempotencyKeysPrimaryID, selectFields ...FinanceIdempotencyKeysField) (financeIdempotencyKeys model.FinanceIdempotencyKeys, err error) {
	var (
		defaultFinanceIdempotencyKeysSelectFields = defaultFinanceIdempotencyKeysSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceIdempotencyKeysSelectFields = composeFinanceIdempotencyKeysSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceIdempotencyKeysCompositePrimaryKeyWhere([]model.FinanceIdempotencyKeysPrimaryID{primaryID})
	query := fmt.Sprintf(financeIdempotencyKeysQueries.selectFinanceIdempotencyKeys+" WHERE "+whereQry, defaultFinanceIdempotencyKeysSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeIdempotencyKeys, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeIdempotencyKeys with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceIdempotencyKeysByID] failed get financeIdempotencyKeys")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceIdempotencyKeysByID(ctx context.Context, primaryID model.FinanceIdempotencyKeysPrimaryID, financeIdempotencyKeys *model.FinanceIdempotencyKeys, financeIdempotencyKeysUpdateFields ...FinanceIdempotencyKeysUpdateField) (err error) {
	exists, err := repo.IsExistFinanceIdempotencyKeysByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceIdempotencyKeys] failed checking financeIdempotencyKeys whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeIdempotencyKeys with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeIdempotencyKeys == nil {
		if len(financeIdempotencyKeysUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceIdempotencyKeysByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeIdempotencyKeys = &model.FinanceIdempotencyKeys{}
	}
	var (
		defaultFinanceIdempotencyKeysUpdateFields = defaultFinanceIdempotencyKeysUpdateFields(*financeIdempotencyKeys)
		tempUpdateField                           FinanceIdempotencyKeysUpdateFieldList
		selectFields                              = NewFinanceIdempotencyKeysSelectFields()
	)
	if len(financeIdempotencyKeysUpdateFields) > 0 {
		for _, updateField := range financeIdempotencyKeysUpdateFields {
			if updateField.financeIdempotencyKeysField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceIdempotencyKeysUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceIdempotencyKeysCompositePrimaryKeyWhere([]model.FinanceIdempotencyKeysPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceIdempotencyKeysCommand(defaultFinanceIdempotencyKeysUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeIdempotencyKeysQueries.updateFinanceIdempotencyKeys+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceIdempotencyKeys] error when try to update financeIdempotencyKeys by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceIdempotencyKeysByFilter(ctx context.Context, filter model.Filter, financeIdempotencyKeysUpdateFields ...FinanceIdempotencyKeysUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeIdempotencyKeysUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceIdempotencyKeysUpdateFieldList
		selectFields = NewFinanceIdempotencyKeysSelectFields()
	)
	for _, updateField := range financeIdempotencyKeysUpdateFields {
		if updateField.financeIdempotencyKeysField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceIdempotencyKeysCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeIdempotencyKeysFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceIdempotencyKeysFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_idempotency_keys\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceIdempotencyKeysByFilter] error when try to update financeIdempotencyKeys by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceIdempotencyKeysByFilter] failed get rows affected")
	}
	return
}

var (
	financeIdempotencyKeysQueries = struct {
		selectFinanceIdempotencyKeys      string
		selectCountFinanceIdempotencyKeys string
		deleteFinanceIdempotencyKeys      string
		updateFinanceIdempotencyKeys      string
		insertFinanceIdempotencyKeys      string
	}{
		selectFinanceIdempotencyKeys:      "SELECT %s FROM \"finance_idempotency_keys\"",
		selectCountFinanceIdempotencyKeys: "SELECT COUNT(\"id\") FROM \"finance_idempotency_keys\"",
		deleteFinanceIdempotencyKeys:      "DELETE FROM \"finance_idempotency_keys\"",
		updateFinanceIdempotencyKeys:      "UPDATE \"finance_idempotency_keys\" SET %s ",
		insertFinanceIdempotencyKeys:      "INSERT INTO \"finance_idempotency_keys\" %s VALUES %s",
	}
)

type FinanceIdempotencyKeysRepository interface {
	CreateFinanceIdempotencyKeys(ctx context.Context, financeIdempotencyKeys *model.FinanceIdempotencyKeys, fieldsInsert ...FinanceIdempotencyKeysField) error
	BulkCreateFinanceIdempotencyKeys(ctx context.Context, financeIdempotencyKeysList []*model.FinanceIdempotencyKeys, fieldsInsert ...FinanceIdempotencyKeysField) error
	ResolveFinanceIdempotencyKeys(ctx context.Context, selectFields ...FinanceIdempotencyKeysField) (model.FinanceIdempotencyKeysList, error)
	ResolveFinanceIdempotencyKeysByID(ctx context.Context, primaryID model.FinanceIdempotencyKeysPrimaryID, selectFields ...FinanceIdempotencyKeysField) (model.FinanceIdempotencyKeys, error)
	UpdateFinanceIdempotencyKeysByID(ctx context.Context, id model.FinanceIdempotencyKeysPrimaryID, financeIdempotencyKeys *model.FinanceIdempotencyKeys, financeIdempotencyKeysUpdateFields ...FinanceIdempotencyKeysUpdateField) error
	UpdateFinanceIdempotencyKeysByFilter(ctx context.Context, filter model.Filter, financeIdempotencyKeysUpdateFields ...FinanceIdempotencyKeysUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceIdempotencyKeys(ctx context.Context, financeIdempotencyKeysListMap map[model.FinanceIdempotencyKeysPrimaryID]*model.FinanceIdempotencyKeys, FinanceIdempotencyKeyssMapUpdateFieldsRequest map[model.FinanceIdempotencyKeysPrimaryID]FinanceIdempotencyKeysUpdateFieldList) (err error)
	DeleteFinanceIdempotencyKeysByID(ctx context.Context, id model.FinanceIdempotencyKeysPrimaryID) error
	BulkDeleteFinanceIdempotencyKeysByIDs(ctx context.Context, ids []model.FinanceIdempotencyKeysPrimaryID) error
	ResolveFinanceIdempotencyKeysByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceIdempotencyKeysFilterResult, err error)
	IsExistFinanceIdempotencyKeysByIDs(ctx context.Context, ids []model.FinanceIdempotencyKeysPrimaryID) (exists bool, notFoundIds []model.FinanceIdempotencyKeysPrimaryID, err error)
	IsExistFinanceIdempotencyKeysByID(ctx context.Context, id model.FinanceIdempotencyKeysPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
