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

func composeInsertFieldsAndParamsProviderCircuitBreakers(providerCircuitBreakersList []model.ProviderCircuitBreakers, fieldsInsert ...ProviderCircuitBreakersField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewProviderCircuitBreakersSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, providerCircuitBreakers := range providerCircuitBreakersList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, providerCircuitBreakers.Id)
			case selectField.ProviderAccountId():
				args = append(args, providerCircuitBreakers.ProviderAccountId)
			case selectField.MethodCode():
				args = append(args, providerCircuitBreakers.MethodCode)
			case selectField.ChannelCode():
				args = append(args, providerCircuitBreakers.ChannelCode)
			case selectField.Status():
				args = append(args, providerCircuitBreakers.Status)
			case selectField.FailureCount():
				args = append(args, providerCircuitBreakers.FailureCount)
			case selectField.SuccessCount():
				args = append(args, providerCircuitBreakers.SuccessCount)
			case selectField.LastFailureAt():
				args = append(args, providerCircuitBreakers.LastFailureAt)
			case selectField.LastSuccessAt():
				args = append(args, providerCircuitBreakers.LastSuccessAt)
			case selectField.OpenedAt():
				args = append(args, providerCircuitBreakers.OpenedAt)
			case selectField.OpenUntil():
				args = append(args, providerCircuitBreakers.OpenUntil)
			case selectField.HalfOpenAt():
				args = append(args, providerCircuitBreakers.HalfOpenAt)
			case selectField.Reason():
				args = append(args, providerCircuitBreakers.Reason)
			case selectField.Metadata():
				args = append(args, providerCircuitBreakers.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, providerCircuitBreakers.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, providerCircuitBreakers.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, providerCircuitBreakers.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, providerCircuitBreakers.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, providerCircuitBreakers.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, providerCircuitBreakers.MetaDeletedBy)

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

func composeProviderCircuitBreakersCompositePrimaryKeyWhere(primaryIDs []model.ProviderCircuitBreakersPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"provider_circuit_breakers\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultProviderCircuitBreakersSelectFields() string {
	fields := NewProviderCircuitBreakersSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeProviderCircuitBreakersSelectFields(selectFields ...ProviderCircuitBreakersField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ProviderCircuitBreakersField string
type ProviderCircuitBreakersFieldList []ProviderCircuitBreakersField

type ProviderCircuitBreakersSelectFields struct {
}

func (ss ProviderCircuitBreakersSelectFields) Id() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("id")
}

func (ss ProviderCircuitBreakersSelectFields) ProviderAccountId() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("provider_account_id")
}

func (ss ProviderCircuitBreakersSelectFields) MethodCode() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("method_code")
}

func (ss ProviderCircuitBreakersSelectFields) ChannelCode() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("channel_code")
}

func (ss ProviderCircuitBreakersSelectFields) Status() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("status")
}

func (ss ProviderCircuitBreakersSelectFields) FailureCount() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("failure_count")
}

func (ss ProviderCircuitBreakersSelectFields) SuccessCount() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("success_count")
}

func (ss ProviderCircuitBreakersSelectFields) LastFailureAt() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("last_failure_at")
}

func (ss ProviderCircuitBreakersSelectFields) LastSuccessAt() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("last_success_at")
}

func (ss ProviderCircuitBreakersSelectFields) OpenedAt() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("opened_at")
}

func (ss ProviderCircuitBreakersSelectFields) OpenUntil() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("open_until")
}

func (ss ProviderCircuitBreakersSelectFields) HalfOpenAt() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("half_open_at")
}

func (ss ProviderCircuitBreakersSelectFields) Reason() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("reason")
}

func (ss ProviderCircuitBreakersSelectFields) Metadata() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("metadata")
}

func (ss ProviderCircuitBreakersSelectFields) MetaCreatedAt() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("meta_created_at")
}

func (ss ProviderCircuitBreakersSelectFields) MetaCreatedBy() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("meta_created_by")
}

func (ss ProviderCircuitBreakersSelectFields) MetaUpdatedAt() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("meta_updated_at")
}

func (ss ProviderCircuitBreakersSelectFields) MetaUpdatedBy() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("meta_updated_by")
}

func (ss ProviderCircuitBreakersSelectFields) MetaDeletedAt() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("meta_deleted_at")
}

func (ss ProviderCircuitBreakersSelectFields) MetaDeletedBy() ProviderCircuitBreakersField {
	return ProviderCircuitBreakersField("meta_deleted_by")
}

func (ss ProviderCircuitBreakersSelectFields) All() ProviderCircuitBreakersFieldList {
	return []ProviderCircuitBreakersField{
		ss.Id(),
		ss.ProviderAccountId(),
		ss.MethodCode(),
		ss.ChannelCode(),
		ss.Status(),
		ss.FailureCount(),
		ss.SuccessCount(),
		ss.LastFailureAt(),
		ss.LastSuccessAt(),
		ss.OpenedAt(),
		ss.OpenUntil(),
		ss.HalfOpenAt(),
		ss.Reason(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewProviderCircuitBreakersSelectFields() ProviderCircuitBreakersSelectFields {
	return ProviderCircuitBreakersSelectFields{}
}

type ProviderCircuitBreakersUpdateFieldOption struct {
	useIncrement bool
}
type ProviderCircuitBreakersUpdateField struct {
	providerCircuitBreakersField ProviderCircuitBreakersField
	opt                          ProviderCircuitBreakersUpdateFieldOption
	value                        interface{}
}
type ProviderCircuitBreakersUpdateFieldList []ProviderCircuitBreakersUpdateField

func defaultProviderCircuitBreakersUpdateFieldOption() ProviderCircuitBreakersUpdateFieldOption {
	return ProviderCircuitBreakersUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementProviderCircuitBreakersOption(useIncrement bool) func(*ProviderCircuitBreakersUpdateFieldOption) {
	return func(pcufo *ProviderCircuitBreakersUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewProviderCircuitBreakersUpdateField(field ProviderCircuitBreakersField, val interface{}, opts ...func(*ProviderCircuitBreakersUpdateFieldOption)) ProviderCircuitBreakersUpdateField {
	defaultOpt := defaultProviderCircuitBreakersUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ProviderCircuitBreakersUpdateField{
		providerCircuitBreakersField: field,
		value:                        val,
		opt:                          defaultOpt,
	}
}
func defaultProviderCircuitBreakersUpdateFields(providerCircuitBreakers model.ProviderCircuitBreakers) (providerCircuitBreakersUpdateFieldList ProviderCircuitBreakersUpdateFieldList) {
	selectFields := NewProviderCircuitBreakersSelectFields()
	providerCircuitBreakersUpdateFieldList = append(providerCircuitBreakersUpdateFieldList,
		NewProviderCircuitBreakersUpdateField(selectFields.Id(), providerCircuitBreakers.Id),
		NewProviderCircuitBreakersUpdateField(selectFields.ProviderAccountId(), providerCircuitBreakers.ProviderAccountId),
		NewProviderCircuitBreakersUpdateField(selectFields.MethodCode(), providerCircuitBreakers.MethodCode),
		NewProviderCircuitBreakersUpdateField(selectFields.ChannelCode(), providerCircuitBreakers.ChannelCode),
		NewProviderCircuitBreakersUpdateField(selectFields.Status(), providerCircuitBreakers.Status),
		NewProviderCircuitBreakersUpdateField(selectFields.FailureCount(), providerCircuitBreakers.FailureCount),
		NewProviderCircuitBreakersUpdateField(selectFields.SuccessCount(), providerCircuitBreakers.SuccessCount),
		NewProviderCircuitBreakersUpdateField(selectFields.LastFailureAt(), providerCircuitBreakers.LastFailureAt),
		NewProviderCircuitBreakersUpdateField(selectFields.LastSuccessAt(), providerCircuitBreakers.LastSuccessAt),
		NewProviderCircuitBreakersUpdateField(selectFields.OpenedAt(), providerCircuitBreakers.OpenedAt),
		NewProviderCircuitBreakersUpdateField(selectFields.OpenUntil(), providerCircuitBreakers.OpenUntil),
		NewProviderCircuitBreakersUpdateField(selectFields.HalfOpenAt(), providerCircuitBreakers.HalfOpenAt),
		NewProviderCircuitBreakersUpdateField(selectFields.Reason(), providerCircuitBreakers.Reason),
		NewProviderCircuitBreakersUpdateField(selectFields.Metadata(), providerCircuitBreakers.Metadata),
		NewProviderCircuitBreakersUpdateField(selectFields.MetaCreatedAt(), providerCircuitBreakers.MetaCreatedAt),
		NewProviderCircuitBreakersUpdateField(selectFields.MetaCreatedBy(), providerCircuitBreakers.MetaCreatedBy),
		NewProviderCircuitBreakersUpdateField(selectFields.MetaUpdatedAt(), providerCircuitBreakers.MetaUpdatedAt),
		NewProviderCircuitBreakersUpdateField(selectFields.MetaUpdatedBy(), providerCircuitBreakers.MetaUpdatedBy),
		NewProviderCircuitBreakersUpdateField(selectFields.MetaDeletedAt(), providerCircuitBreakers.MetaDeletedAt),
		NewProviderCircuitBreakersUpdateField(selectFields.MetaDeletedBy(), providerCircuitBreakers.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsProviderCircuitBreakersCommand(providerCircuitBreakersUpdateFieldList ProviderCircuitBreakersUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range providerCircuitBreakersUpdateFieldList {
		field := string(updateField.providerCircuitBreakersField)
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

func (repo *RepositoryImpl) BulkCreateProviderCircuitBreakers(ctx context.Context, providerCircuitBreakersList []*model.ProviderCircuitBreakers, fieldsInsert ...ProviderCircuitBreakersField) (err error) {
	var (
		fieldsStr                        string
		valueListStr                     []string
		argsList                         []interface{}
		primaryIds                       []model.ProviderCircuitBreakersPrimaryID
		providerCircuitBreakersValueList []model.ProviderCircuitBreakers
	)

	if len(fieldsInsert) == 0 {
		selectField := NewProviderCircuitBreakersSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, providerCircuitBreakers := range providerCircuitBreakersList {

		primaryIds = append(primaryIds, providerCircuitBreakers.ToProviderCircuitBreakersPrimaryID())

		providerCircuitBreakersValueList = append(providerCircuitBreakersValueList, *providerCircuitBreakers)
	}

	_, notFoundIds, err := repo.IsExistProviderCircuitBreakersByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderCircuitBreakers] failed checking providerCircuitBreakers whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ProviderCircuitBreakersPrimaryID{}
		mapNotFoundIds := map[model.ProviderCircuitBreakersPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "providerCircuitBreakers", fmt.Sprintf("providerCircuitBreakers with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsProviderCircuitBreakers(providerCircuitBreakersValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(providerCircuitBreakersQueries.insertProviderCircuitBreakers, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderCircuitBreakers] failed exec create providerCircuitBreakers query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteProviderCircuitBreakersByIDs(ctx context.Context, primaryIDs []model.ProviderCircuitBreakersPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistProviderCircuitBreakersByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderCircuitBreakersByIDs] failed checking providerCircuitBreakers whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerCircuitBreakers with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_circuit_breakers\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(providerCircuitBreakersQueries.deleteProviderCircuitBreakers + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderCircuitBreakersByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderCircuitBreakersByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistProviderCircuitBreakersByIDs(ctx context.Context, ids []model.ProviderCircuitBreakersPrimaryID) (exists bool, notFoundIds []model.ProviderCircuitBreakersPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_circuit_breakers\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(providerCircuitBreakersQueries.selectProviderCircuitBreakers, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderCircuitBreakersByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ProviderCircuitBreakersPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderCircuitBreakersByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ProviderCircuitBreakersPrimaryID]bool{}
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

// BulkUpdateProviderCircuitBreakers is used to bulk update providerCircuitBreakers, by default it will update all field
// if want to update specific field, then fill providerCircuitBreakerssMapUpdateFieldsRequest else please fill providerCircuitBreakerssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateProviderCircuitBreakers(ctx context.Context, providerCircuitBreakerssMap map[model.ProviderCircuitBreakersPrimaryID]*model.ProviderCircuitBreakers, providerCircuitBreakerssMapUpdateFieldsRequest map[model.ProviderCircuitBreakersPrimaryID]ProviderCircuitBreakersUpdateFieldList) (err error) {
	if len(providerCircuitBreakerssMap) == 0 && len(providerCircuitBreakerssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		providerCircuitBreakerssMapUpdateField map[model.ProviderCircuitBreakersPrimaryID]ProviderCircuitBreakersUpdateFieldList = map[model.ProviderCircuitBreakersPrimaryID]ProviderCircuitBreakersUpdateFieldList{}
		asTableValues                          string                                                                            = "myvalues"
	)

	if len(providerCircuitBreakerssMap) > 0 {
		for id, providerCircuitBreakers := range providerCircuitBreakerssMap {
			if providerCircuitBreakers == nil {
				log.Error().Err(err).Msg("[BulkUpdateProviderCircuitBreakers] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			providerCircuitBreakerssMapUpdateField[id] = defaultProviderCircuitBreakersUpdateFields(*providerCircuitBreakers)
		}
	} else {
		providerCircuitBreakerssMapUpdateField = providerCircuitBreakerssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateProviderCircuitBreakersQuery(providerCircuitBreakerssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistProviderCircuitBreakersByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderCircuitBreakers] failed checking providerCircuitBreakers whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerCircuitBreakers with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeProviderCircuitBreakersCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"provider_circuit_breakers\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderCircuitBreakers] failed exec query")
	}
	return
}

type ProviderCircuitBreakersFieldParameter struct {
	param string
	args  []interface{}
}

func NewProviderCircuitBreakersFieldParameter(param string, args ...interface{}) ProviderCircuitBreakersFieldParameter {
	return ProviderCircuitBreakersFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateProviderCircuitBreakersQuery(mapProviderCircuitBreakerss map[model.ProviderCircuitBreakersPrimaryID]ProviderCircuitBreakersUpdateFieldList, asTableValues string) (primaryIDs []model.ProviderCircuitBreakersPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ProviderCircuitBreakersPrimaryID]map[string]interface{}{}
	providerCircuitBreakersSelectFields := NewProviderCircuitBreakersSelectFields()
	for id, updateFields := range mapProviderCircuitBreakerss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.providerCircuitBreakersField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapProviderCircuitBreakerss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetProviderCircuitBreakersFieldType(updateField.providerCircuitBreakersField)))
			args = append(args, fields[string(updateField.providerCircuitBreakersField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.providerCircuitBreakersField))
		if updateField.providerCircuitBreakersField == providerCircuitBreakersSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.providerCircuitBreakersField, asTableValues, updateField.providerCircuitBreakersField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.providerCircuitBreakersField,
				"\"provider_circuit_breakers\"", updateField.providerCircuitBreakersField,
				asTableValues, updateField.providerCircuitBreakersField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeProviderCircuitBreakersCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ProviderCircuitBreakersPrimaryID, asTableValue string) (whereQry string) {
	providerCircuitBreakersSelectFields := NewProviderCircuitBreakersSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"provider_circuit_breakers\".\"id\" = %s.\"id\"::"+GetProviderCircuitBreakersFieldType(providerCircuitBreakersSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetProviderCircuitBreakersFieldType(providerCircuitBreakersField ProviderCircuitBreakersField) string {
	selectProviderCircuitBreakersFields := NewProviderCircuitBreakersSelectFields()
	switch providerCircuitBreakersField {

	case selectProviderCircuitBreakersFields.Id():
		return "uuid"

	case selectProviderCircuitBreakersFields.ProviderAccountId():
		return "uuid"

	case selectProviderCircuitBreakersFields.MethodCode():
		return "text"

	case selectProviderCircuitBreakersFields.ChannelCode():
		return "text"

	case selectProviderCircuitBreakersFields.Status():
		return "circuit_status_enum"

	case selectProviderCircuitBreakersFields.FailureCount():
		return "int4"

	case selectProviderCircuitBreakersFields.SuccessCount():
		return "int4"

	case selectProviderCircuitBreakersFields.LastFailureAt():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.LastSuccessAt():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.OpenedAt():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.OpenUntil():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.HalfOpenAt():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.Reason():
		return "text"

	case selectProviderCircuitBreakersFields.Metadata():
		return "jsonb"

	case selectProviderCircuitBreakersFields.MetaCreatedAt():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.MetaCreatedBy():
		return "uuid"

	case selectProviderCircuitBreakersFields.MetaUpdatedAt():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.MetaUpdatedBy():
		return "uuid"

	case selectProviderCircuitBreakersFields.MetaDeletedAt():
		return "timestamptz"

	case selectProviderCircuitBreakersFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateProviderCircuitBreakers(ctx context.Context, providerCircuitBreakers *model.ProviderCircuitBreakers, fieldsInsert ...ProviderCircuitBreakersField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewProviderCircuitBreakersSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ProviderCircuitBreakersPrimaryID{
		Id: providerCircuitBreakers.Id,
	}
	exists, err := repo.IsExistProviderCircuitBreakersByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderCircuitBreakers] failed checking providerCircuitBreakers whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "providerCircuitBreakers", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsProviderCircuitBreakers([]model.ProviderCircuitBreakers{*providerCircuitBreakers}, fieldsInsert...)
	commandQuery := fmt.Sprintf(providerCircuitBreakersQueries.insertProviderCircuitBreakers, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderCircuitBreakers] failed exec create providerCircuitBreakers query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteProviderCircuitBreakersByID(ctx context.Context, primaryID model.ProviderCircuitBreakersPrimaryID) (err error) {
	exists, err := repo.IsExistProviderCircuitBreakersByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderCircuitBreakersByID] failed checking providerCircuitBreakers whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerCircuitBreakers with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeProviderCircuitBreakersCompositePrimaryKeyWhere([]model.ProviderCircuitBreakersPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(providerCircuitBreakersQueries.deleteProviderCircuitBreakers + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderCircuitBreakersByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderCircuitBreakersByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderCircuitBreakersFilterResult, err error) {
	query, args, err := composeProviderCircuitBreakersFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderCircuitBreakersByFilter] failed compose providerCircuitBreakers filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderCircuitBreakersByFilter] failed get providerCircuitBreakers by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeProviderCircuitBreakersFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ProviderCircuitBreakersFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeProviderCircuitBreakersFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeProviderCircuitBreakersSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeProviderCircuitBreakersFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewProviderCircuitBreakersFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 20+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["method_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_code\"")
			selectedColumns["method_code"] = struct{}{}
		}
		if _, selected := selectedColumns["channel_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"channel_code\"")
			selectedColumns["channel_code"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_count\"")
			selectedColumns["failure_count"] = struct{}{}
		}
		if _, selected := selectedColumns["success_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"success_count\"")
			selectedColumns["success_count"] = struct{}{}
		}
		if _, selected := selectedColumns["last_failure_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"last_failure_at\"")
			selectedColumns["last_failure_at"] = struct{}{}
		}
		if _, selected := selectedColumns["last_success_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"last_success_at\"")
			selectedColumns["last_success_at"] = struct{}{}
		}
		if _, selected := selectedColumns["opened_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"opened_at\"")
			selectedColumns["opened_at"] = struct{}{}
		}
		if _, selected := selectedColumns["open_until"]; !selected {
			selectColumns = append(selectColumns, "base.\"open_until\"")
			selectedColumns["open_until"] = struct{}{}
		}
		if _, selected := selectedColumns["half_open_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"half_open_at\"")
			selectedColumns["half_open_at"] = struct{}{}
		}
		if _, selected := selectedColumns["reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason\"")
			selectedColumns["reason"] = struct{}{}
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

type providerCircuitBreakersFilterPlaceholder struct {
	index int
}

func (p *providerCircuitBreakersFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeProviderCircuitBreakersFilterPredicate(filterField model.FilterField, placeholders *providerCircuitBreakersFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewProviderCircuitBreakersFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeProviderCircuitBreakersFilterSQLExpr(spec)
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

func composeProviderCircuitBreakersFilterGroup(group model.FilterGroup, placeholders *providerCircuitBreakersFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeProviderCircuitBreakersFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeProviderCircuitBreakersFilterGroup(child, placeholders, args, requiredJoins)
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

func composeProviderCircuitBreakersFilterWhereQueries(filter model.Filter, placeholders *providerCircuitBreakersFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeProviderCircuitBreakersFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeProviderCircuitBreakersFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeProviderCircuitBreakersFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateProviderCircuitBreakersFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeProviderCircuitBreakersSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeProviderCircuitBreakersFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := providerCircuitBreakersFilterPlaceholder{index: 1}
	whereQueries, err := composeProviderCircuitBreakersFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewProviderCircuitBreakersFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeProviderCircuitBreakersFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeProviderCircuitBreakersSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"provider_circuit_breakers\" base%s", strings.Join(selectColumns, ","), composeProviderCircuitBreakersFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistProviderCircuitBreakersByID(ctx context.Context, primaryID model.ProviderCircuitBreakersPrimaryID) (exists bool, err error) {
	whereQuery, params := composeProviderCircuitBreakersCompositePrimaryKeyWhere([]model.ProviderCircuitBreakersPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", providerCircuitBreakersQueries.selectCountProviderCircuitBreakers, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderCircuitBreakersByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderCircuitBreakers(ctx context.Context, selectFields ...ProviderCircuitBreakersField) (providerCircuitBreakersList model.ProviderCircuitBreakersList, err error) {
	var (
		defaultProviderCircuitBreakersSelectFields = defaultProviderCircuitBreakersSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderCircuitBreakersSelectFields = composeProviderCircuitBreakersSelectFields(selectFields...)
	}
	query := fmt.Sprintf(providerCircuitBreakersQueries.selectProviderCircuitBreakers, defaultProviderCircuitBreakersSelectFields)

	err = repo.db.Read.SelectContext(ctx, &providerCircuitBreakersList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderCircuitBreakers] failed get providerCircuitBreakers list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderCircuitBreakersByID(ctx context.Context, primaryID model.ProviderCircuitBreakersPrimaryID, selectFields ...ProviderCircuitBreakersField) (providerCircuitBreakers model.ProviderCircuitBreakers, err error) {
	var (
		defaultProviderCircuitBreakersSelectFields = defaultProviderCircuitBreakersSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderCircuitBreakersSelectFields = composeProviderCircuitBreakersSelectFields(selectFields...)
	}
	whereQry, params := composeProviderCircuitBreakersCompositePrimaryKeyWhere([]model.ProviderCircuitBreakersPrimaryID{primaryID})
	query := fmt.Sprintf(providerCircuitBreakersQueries.selectProviderCircuitBreakers+" WHERE "+whereQry, defaultProviderCircuitBreakersSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &providerCircuitBreakers, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("providerCircuitBreakers with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveProviderCircuitBreakersByID] failed get providerCircuitBreakers")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateProviderCircuitBreakersByID(ctx context.Context, primaryID model.ProviderCircuitBreakersPrimaryID, providerCircuitBreakers *model.ProviderCircuitBreakers, providerCircuitBreakersUpdateFields ...ProviderCircuitBreakersUpdateField) (err error) {
	exists, err := repo.IsExistProviderCircuitBreakersByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderCircuitBreakers] failed checking providerCircuitBreakers whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerCircuitBreakers with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if providerCircuitBreakers == nil {
		if len(providerCircuitBreakersUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateProviderCircuitBreakersByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		providerCircuitBreakers = &model.ProviderCircuitBreakers{}
	}
	var (
		defaultProviderCircuitBreakersUpdateFields = defaultProviderCircuitBreakersUpdateFields(*providerCircuitBreakers)
		tempUpdateField                            ProviderCircuitBreakersUpdateFieldList
		selectFields                               = NewProviderCircuitBreakersSelectFields()
	)
	if len(providerCircuitBreakersUpdateFields) > 0 {
		for _, updateField := range providerCircuitBreakersUpdateFields {
			if updateField.providerCircuitBreakersField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultProviderCircuitBreakersUpdateFields = tempUpdateField
	}
	whereQuery, params := composeProviderCircuitBreakersCompositePrimaryKeyWhere([]model.ProviderCircuitBreakersPrimaryID{primaryID})
	fields, args := composeUpdateFieldsProviderCircuitBreakersCommand(defaultProviderCircuitBreakersUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(providerCircuitBreakersQueries.updateProviderCircuitBreakers+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderCircuitBreakers] error when try to update providerCircuitBreakers by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateProviderCircuitBreakersByFilter(ctx context.Context, filter model.Filter, providerCircuitBreakersUpdateFields ...ProviderCircuitBreakersUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(providerCircuitBreakersUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ProviderCircuitBreakersUpdateFieldList
		selectFields = NewProviderCircuitBreakersSelectFields()
	)
	for _, updateField := range providerCircuitBreakersUpdateFields {
		if updateField.providerCircuitBreakersField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsProviderCircuitBreakersCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := providerCircuitBreakersFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeProviderCircuitBreakersFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"provider_circuit_breakers\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderCircuitBreakersByFilter] error when try to update providerCircuitBreakers by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderCircuitBreakersByFilter] failed get rows affected")
	}
	return
}

var (
	providerCircuitBreakersQueries = struct {
		selectProviderCircuitBreakers      string
		selectCountProviderCircuitBreakers string
		deleteProviderCircuitBreakers      string
		updateProviderCircuitBreakers      string
		insertProviderCircuitBreakers      string
	}{
		selectProviderCircuitBreakers:      "SELECT %s FROM \"provider_circuit_breakers\"",
		selectCountProviderCircuitBreakers: "SELECT COUNT(\"id\") FROM \"provider_circuit_breakers\"",
		deleteProviderCircuitBreakers:      "DELETE FROM \"provider_circuit_breakers\"",
		updateProviderCircuitBreakers:      "UPDATE \"provider_circuit_breakers\" SET %s ",
		insertProviderCircuitBreakers:      "INSERT INTO \"provider_circuit_breakers\" %s VALUES %s",
	}
)

type ProviderCircuitBreakersRepository interface {
	CreateProviderCircuitBreakers(ctx context.Context, providerCircuitBreakers *model.ProviderCircuitBreakers, fieldsInsert ...ProviderCircuitBreakersField) error
	BulkCreateProviderCircuitBreakers(ctx context.Context, providerCircuitBreakersList []*model.ProviderCircuitBreakers, fieldsInsert ...ProviderCircuitBreakersField) error
	ResolveProviderCircuitBreakers(ctx context.Context, selectFields ...ProviderCircuitBreakersField) (model.ProviderCircuitBreakersList, error)
	ResolveProviderCircuitBreakersByID(ctx context.Context, primaryID model.ProviderCircuitBreakersPrimaryID, selectFields ...ProviderCircuitBreakersField) (model.ProviderCircuitBreakers, error)
	UpdateProviderCircuitBreakersByID(ctx context.Context, id model.ProviderCircuitBreakersPrimaryID, providerCircuitBreakers *model.ProviderCircuitBreakers, providerCircuitBreakersUpdateFields ...ProviderCircuitBreakersUpdateField) error
	UpdateProviderCircuitBreakersByFilter(ctx context.Context, filter model.Filter, providerCircuitBreakersUpdateFields ...ProviderCircuitBreakersUpdateField) (rowsAffected int64, err error)
	BulkUpdateProviderCircuitBreakers(ctx context.Context, providerCircuitBreakersListMap map[model.ProviderCircuitBreakersPrimaryID]*model.ProviderCircuitBreakers, ProviderCircuitBreakerssMapUpdateFieldsRequest map[model.ProviderCircuitBreakersPrimaryID]ProviderCircuitBreakersUpdateFieldList) (err error)
	DeleteProviderCircuitBreakersByID(ctx context.Context, id model.ProviderCircuitBreakersPrimaryID) error
	BulkDeleteProviderCircuitBreakersByIDs(ctx context.Context, ids []model.ProviderCircuitBreakersPrimaryID) error
	ResolveProviderCircuitBreakersByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderCircuitBreakersFilterResult, err error)
	IsExistProviderCircuitBreakersByIDs(ctx context.Context, ids []model.ProviderCircuitBreakersPrimaryID) (exists bool, notFoundIds []model.ProviderCircuitBreakersPrimaryID, err error)
	IsExistProviderCircuitBreakersByID(ctx context.Context, id model.ProviderCircuitBreakersPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
