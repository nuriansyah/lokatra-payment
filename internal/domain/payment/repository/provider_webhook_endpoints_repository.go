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

func composeInsertFieldsAndParamsProviderWebhookEndpoints(providerWebhookEndpointsList []model.ProviderWebhookEndpoints, fieldsInsert ...ProviderWebhookEndpointsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewProviderWebhookEndpointsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, providerWebhookEndpoints := range providerWebhookEndpointsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, providerWebhookEndpoints.Id)
			case selectField.ProviderAccountId():
				args = append(args, providerWebhookEndpoints.ProviderAccountId)
			case selectField.ProviderCode():
				args = append(args, providerWebhookEndpoints.ProviderCode)
			case selectField.EndpointKey():
				args = append(args, providerWebhookEndpoints.EndpointKey)
			case selectField.Environment():
				args = append(args, providerWebhookEndpoints.Environment)
			case selectField.SecretRef():
				args = append(args, providerWebhookEndpoints.SecretRef)
			case selectField.SignatureAlgorithm():
				args = append(args, providerWebhookEndpoints.SignatureAlgorithm)
			case selectField.IsActive():
				args = append(args, providerWebhookEndpoints.IsActive)
			case selectField.Metadata():
				args = append(args, providerWebhookEndpoints.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, providerWebhookEndpoints.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, providerWebhookEndpoints.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, providerWebhookEndpoints.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, providerWebhookEndpoints.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, providerWebhookEndpoints.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, providerWebhookEndpoints.MetaDeletedBy)

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

func composeProviderWebhookEndpointsCompositePrimaryKeyWhere(primaryIDs []model.ProviderWebhookEndpointsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"provider_webhook_endpoints\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultProviderWebhookEndpointsSelectFields() string {
	fields := NewProviderWebhookEndpointsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeProviderWebhookEndpointsSelectFields(selectFields ...ProviderWebhookEndpointsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ProviderWebhookEndpointsField string
type ProviderWebhookEndpointsFieldList []ProviderWebhookEndpointsField

type ProviderWebhookEndpointsSelectFields struct {
}

func (ss ProviderWebhookEndpointsSelectFields) Id() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("id")
}

func (ss ProviderWebhookEndpointsSelectFields) ProviderAccountId() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("provider_account_id")
}

func (ss ProviderWebhookEndpointsSelectFields) ProviderCode() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("provider_code")
}

func (ss ProviderWebhookEndpointsSelectFields) EndpointKey() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("endpoint_key")
}

func (ss ProviderWebhookEndpointsSelectFields) Environment() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("environment")
}

func (ss ProviderWebhookEndpointsSelectFields) SecretRef() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("secret_ref")
}

func (ss ProviderWebhookEndpointsSelectFields) SignatureAlgorithm() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("signature_algorithm")
}

func (ss ProviderWebhookEndpointsSelectFields) IsActive() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("is_active")
}

func (ss ProviderWebhookEndpointsSelectFields) Metadata() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("metadata")
}

func (ss ProviderWebhookEndpointsSelectFields) MetaCreatedAt() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("meta_created_at")
}

func (ss ProviderWebhookEndpointsSelectFields) MetaCreatedBy() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("meta_created_by")
}

func (ss ProviderWebhookEndpointsSelectFields) MetaUpdatedAt() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("meta_updated_at")
}

func (ss ProviderWebhookEndpointsSelectFields) MetaUpdatedBy() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("meta_updated_by")
}

func (ss ProviderWebhookEndpointsSelectFields) MetaDeletedAt() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("meta_deleted_at")
}

func (ss ProviderWebhookEndpointsSelectFields) MetaDeletedBy() ProviderWebhookEndpointsField {
	return ProviderWebhookEndpointsField("meta_deleted_by")
}

func (ss ProviderWebhookEndpointsSelectFields) All() ProviderWebhookEndpointsFieldList {
	return []ProviderWebhookEndpointsField{
		ss.Id(),
		ss.ProviderAccountId(),
		ss.ProviderCode(),
		ss.EndpointKey(),
		ss.Environment(),
		ss.SecretRef(),
		ss.SignatureAlgorithm(),
		ss.IsActive(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewProviderWebhookEndpointsSelectFields() ProviderWebhookEndpointsSelectFields {
	return ProviderWebhookEndpointsSelectFields{}
}

type ProviderWebhookEndpointsUpdateFieldOption struct {
	useIncrement bool
}
type ProviderWebhookEndpointsUpdateField struct {
	providerWebhookEndpointsField ProviderWebhookEndpointsField
	opt                           ProviderWebhookEndpointsUpdateFieldOption
	value                         interface{}
}
type ProviderWebhookEndpointsUpdateFieldList []ProviderWebhookEndpointsUpdateField

func defaultProviderWebhookEndpointsUpdateFieldOption() ProviderWebhookEndpointsUpdateFieldOption {
	return ProviderWebhookEndpointsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementProviderWebhookEndpointsOption(useIncrement bool) func(*ProviderWebhookEndpointsUpdateFieldOption) {
	return func(pcufo *ProviderWebhookEndpointsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewProviderWebhookEndpointsUpdateField(field ProviderWebhookEndpointsField, val interface{}, opts ...func(*ProviderWebhookEndpointsUpdateFieldOption)) ProviderWebhookEndpointsUpdateField {
	defaultOpt := defaultProviderWebhookEndpointsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ProviderWebhookEndpointsUpdateField{
		providerWebhookEndpointsField: field,
		value:                         val,
		opt:                           defaultOpt,
	}
}
func defaultProviderWebhookEndpointsUpdateFields(providerWebhookEndpoints model.ProviderWebhookEndpoints) (providerWebhookEndpointsUpdateFieldList ProviderWebhookEndpointsUpdateFieldList) {
	selectFields := NewProviderWebhookEndpointsSelectFields()
	providerWebhookEndpointsUpdateFieldList = append(providerWebhookEndpointsUpdateFieldList,
		NewProviderWebhookEndpointsUpdateField(selectFields.Id(), providerWebhookEndpoints.Id),
		NewProviderWebhookEndpointsUpdateField(selectFields.ProviderAccountId(), providerWebhookEndpoints.ProviderAccountId),
		NewProviderWebhookEndpointsUpdateField(selectFields.ProviderCode(), providerWebhookEndpoints.ProviderCode),
		NewProviderWebhookEndpointsUpdateField(selectFields.EndpointKey(), providerWebhookEndpoints.EndpointKey),
		NewProviderWebhookEndpointsUpdateField(selectFields.Environment(), providerWebhookEndpoints.Environment),
		NewProviderWebhookEndpointsUpdateField(selectFields.SecretRef(), providerWebhookEndpoints.SecretRef),
		NewProviderWebhookEndpointsUpdateField(selectFields.SignatureAlgorithm(), providerWebhookEndpoints.SignatureAlgorithm),
		NewProviderWebhookEndpointsUpdateField(selectFields.IsActive(), providerWebhookEndpoints.IsActive),
		NewProviderWebhookEndpointsUpdateField(selectFields.Metadata(), providerWebhookEndpoints.Metadata),
		NewProviderWebhookEndpointsUpdateField(selectFields.MetaCreatedAt(), providerWebhookEndpoints.MetaCreatedAt),
		NewProviderWebhookEndpointsUpdateField(selectFields.MetaCreatedBy(), providerWebhookEndpoints.MetaCreatedBy),
		NewProviderWebhookEndpointsUpdateField(selectFields.MetaUpdatedAt(), providerWebhookEndpoints.MetaUpdatedAt),
		NewProviderWebhookEndpointsUpdateField(selectFields.MetaUpdatedBy(), providerWebhookEndpoints.MetaUpdatedBy),
		NewProviderWebhookEndpointsUpdateField(selectFields.MetaDeletedAt(), providerWebhookEndpoints.MetaDeletedAt),
		NewProviderWebhookEndpointsUpdateField(selectFields.MetaDeletedBy(), providerWebhookEndpoints.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsProviderWebhookEndpointsCommand(providerWebhookEndpointsUpdateFieldList ProviderWebhookEndpointsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range providerWebhookEndpointsUpdateFieldList {
		field := string(updateField.providerWebhookEndpointsField)
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

func (repo *RepositoryImpl) BulkCreateProviderWebhookEndpoints(ctx context.Context, providerWebhookEndpointsList []*model.ProviderWebhookEndpoints, fieldsInsert ...ProviderWebhookEndpointsField) (err error) {
	var (
		fieldsStr                         string
		valueListStr                      []string
		argsList                          []interface{}
		primaryIds                        []model.ProviderWebhookEndpointsPrimaryID
		providerWebhookEndpointsValueList []model.ProviderWebhookEndpoints
	)

	if len(fieldsInsert) == 0 {
		selectField := NewProviderWebhookEndpointsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, providerWebhookEndpoints := range providerWebhookEndpointsList {

		primaryIds = append(primaryIds, providerWebhookEndpoints.ToProviderWebhookEndpointsPrimaryID())

		providerWebhookEndpointsValueList = append(providerWebhookEndpointsValueList, *providerWebhookEndpoints)
	}

	_, notFoundIds, err := repo.IsExistProviderWebhookEndpointsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderWebhookEndpoints] failed checking providerWebhookEndpoints whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ProviderWebhookEndpointsPrimaryID{}
		mapNotFoundIds := map[model.ProviderWebhookEndpointsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "providerWebhookEndpoints", fmt.Sprintf("providerWebhookEndpoints with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsProviderWebhookEndpoints(providerWebhookEndpointsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(providerWebhookEndpointsQueries.insertProviderWebhookEndpoints, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderWebhookEndpoints] failed exec create providerWebhookEndpoints query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteProviderWebhookEndpointsByIDs(ctx context.Context, primaryIDs []model.ProviderWebhookEndpointsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistProviderWebhookEndpointsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderWebhookEndpointsByIDs] failed checking providerWebhookEndpoints whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEndpoints with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_webhook_endpoints\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := providerWebhookEndpointsQueries.deleteProviderWebhookEndpoints + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderWebhookEndpointsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderWebhookEndpointsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistProviderWebhookEndpointsByIDs(ctx context.Context, ids []model.ProviderWebhookEndpointsPrimaryID) (exists bool, notFoundIds []model.ProviderWebhookEndpointsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_webhook_endpoints\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(providerWebhookEndpointsQueries.selectProviderWebhookEndpoints, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderWebhookEndpointsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ProviderWebhookEndpointsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderWebhookEndpointsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ProviderWebhookEndpointsPrimaryID]bool{}
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

// BulkUpdateProviderWebhookEndpoints is used to bulk update providerWebhookEndpoints, by default it will update all field
// if want to update specific field, then fill providerWebhookEndpointssMapUpdateFieldsRequest else please fill providerWebhookEndpointssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateProviderWebhookEndpoints(ctx context.Context, providerWebhookEndpointssMap map[model.ProviderWebhookEndpointsPrimaryID]*model.ProviderWebhookEndpoints, providerWebhookEndpointssMapUpdateFieldsRequest map[model.ProviderWebhookEndpointsPrimaryID]ProviderWebhookEndpointsUpdateFieldList) (err error) {
	if len(providerWebhookEndpointssMap) == 0 && len(providerWebhookEndpointssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		providerWebhookEndpointssMapUpdateField map[model.ProviderWebhookEndpointsPrimaryID]ProviderWebhookEndpointsUpdateFieldList = map[model.ProviderWebhookEndpointsPrimaryID]ProviderWebhookEndpointsUpdateFieldList{}
		asTableValues                           string                                                                              = "myvalues"
	)

	if len(providerWebhookEndpointssMap) > 0 {
		for id, providerWebhookEndpoints := range providerWebhookEndpointssMap {
			if providerWebhookEndpoints == nil {
				log.Error().Err(err).Msg("[BulkUpdateProviderWebhookEndpoints] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			providerWebhookEndpointssMapUpdateField[id] = defaultProviderWebhookEndpointsUpdateFields(*providerWebhookEndpoints)
		}
	} else {
		providerWebhookEndpointssMapUpdateField = providerWebhookEndpointssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateProviderWebhookEndpointsQuery(providerWebhookEndpointssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistProviderWebhookEndpointsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderWebhookEndpoints] failed checking providerWebhookEndpoints whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEndpoints with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeProviderWebhookEndpointsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"provider_webhook_endpoints\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderWebhookEndpoints] failed exec query")
	}
	return
}

type ProviderWebhookEndpointsFieldParameter struct {
	param string
	args  []interface{}
}

func NewProviderWebhookEndpointsFieldParameter(param string, args ...interface{}) ProviderWebhookEndpointsFieldParameter {
	return ProviderWebhookEndpointsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateProviderWebhookEndpointsQuery(mapProviderWebhookEndpointss map[model.ProviderWebhookEndpointsPrimaryID]ProviderWebhookEndpointsUpdateFieldList, asTableValues string) (primaryIDs []model.ProviderWebhookEndpointsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ProviderWebhookEndpointsPrimaryID]map[string]interface{}{}
	providerWebhookEndpointsSelectFields := NewProviderWebhookEndpointsSelectFields()
	for id, updateFields := range mapProviderWebhookEndpointss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.providerWebhookEndpointsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapProviderWebhookEndpointss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetProviderWebhookEndpointsFieldType(updateField.providerWebhookEndpointsField)))
			args = append(args, fields[string(updateField.providerWebhookEndpointsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.providerWebhookEndpointsField))
		if updateField.providerWebhookEndpointsField == providerWebhookEndpointsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.providerWebhookEndpointsField, asTableValues, updateField.providerWebhookEndpointsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.providerWebhookEndpointsField,
				"\"provider_webhook_endpoints\"", updateField.providerWebhookEndpointsField,
				asTableValues, updateField.providerWebhookEndpointsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeProviderWebhookEndpointsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ProviderWebhookEndpointsPrimaryID, asTableValue string) (whereQry string) {
	providerWebhookEndpointsSelectFields := NewProviderWebhookEndpointsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"provider_webhook_endpoints\".\"id\" = %s.\"id\"::"+GetProviderWebhookEndpointsFieldType(providerWebhookEndpointsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetProviderWebhookEndpointsFieldType(providerWebhookEndpointsField ProviderWebhookEndpointsField) string {
	selectProviderWebhookEndpointsFields := NewProviderWebhookEndpointsSelectFields()
	switch providerWebhookEndpointsField {

	case selectProviderWebhookEndpointsFields.Id():
		return "uuid"

	case selectProviderWebhookEndpointsFields.ProviderAccountId():
		return "uuid"

	case selectProviderWebhookEndpointsFields.ProviderCode():
		return "text"

	case selectProviderWebhookEndpointsFields.EndpointKey():
		return "text"

	case selectProviderWebhookEndpointsFields.Environment():
		return "text"

	case selectProviderWebhookEndpointsFields.SecretRef():
		return "text"

	case selectProviderWebhookEndpointsFields.SignatureAlgorithm():
		return "text"

	case selectProviderWebhookEndpointsFields.IsActive():
		return "bool"

	case selectProviderWebhookEndpointsFields.Metadata():
		return "jsonb"

	case selectProviderWebhookEndpointsFields.MetaCreatedAt():
		return "timestamptz"

	case selectProviderWebhookEndpointsFields.MetaCreatedBy():
		return "uuid"

	case selectProviderWebhookEndpointsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectProviderWebhookEndpointsFields.MetaUpdatedBy():
		return "uuid"

	case selectProviderWebhookEndpointsFields.MetaDeletedAt():
		return "timestamptz"

	case selectProviderWebhookEndpointsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateProviderWebhookEndpoints(ctx context.Context, providerWebhookEndpoints *model.ProviderWebhookEndpoints, fieldsInsert ...ProviderWebhookEndpointsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewProviderWebhookEndpointsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ProviderWebhookEndpointsPrimaryID{
		Id: providerWebhookEndpoints.Id,
	}
	exists, err := repo.IsExistProviderWebhookEndpointsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderWebhookEndpoints] failed checking providerWebhookEndpoints whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "providerWebhookEndpoints", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsProviderWebhookEndpoints([]model.ProviderWebhookEndpoints{*providerWebhookEndpoints}, fieldsInsert...)
	commandQuery := fmt.Sprintf(providerWebhookEndpointsQueries.insertProviderWebhookEndpoints, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderWebhookEndpoints] failed exec create providerWebhookEndpoints query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteProviderWebhookEndpointsByID(ctx context.Context, primaryID model.ProviderWebhookEndpointsPrimaryID) (err error) {
	exists, err := repo.IsExistProviderWebhookEndpointsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderWebhookEndpointsByID] failed checking providerWebhookEndpoints whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEndpoints with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeProviderWebhookEndpointsCompositePrimaryKeyWhere([]model.ProviderWebhookEndpointsPrimaryID{primaryID})
	commandQuery := providerWebhookEndpointsQueries.deleteProviderWebhookEndpoints + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderWebhookEndpointsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderWebhookEndpointsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderWebhookEndpointsFilterResult, err error) {
	query, args, err := composeProviderWebhookEndpointsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderWebhookEndpointsByFilter] failed compose providerWebhookEndpoints filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderWebhookEndpointsByFilter] failed get providerWebhookEndpoints by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeProviderWebhookEndpointsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ProviderWebhookEndpointsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeProviderWebhookEndpointsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeProviderWebhookEndpointsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeProviderWebhookEndpointsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewProviderWebhookEndpointsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 15+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_code\"")
			selectedColumns["provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["endpoint_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"endpoint_key\"")
			selectedColumns["endpoint_key"] = struct{}{}
		}
		if _, selected := selectedColumns["environment"]; !selected {
			selectColumns = append(selectColumns, "base.\"environment\"")
			selectedColumns["environment"] = struct{}{}
		}
		if _, selected := selectedColumns["secret_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"secret_ref\"")
			selectedColumns["secret_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["signature_algorithm"]; !selected {
			selectColumns = append(selectColumns, "base.\"signature_algorithm\"")
			selectedColumns["signature_algorithm"] = struct{}{}
		}
		if _, selected := selectedColumns["is_active"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_active\"")
			selectedColumns["is_active"] = struct{}{}
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

type providerWebhookEndpointsFilterPlaceholder struct {
	index int
}

func (p *providerWebhookEndpointsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeProviderWebhookEndpointsFilterPredicate(filterField model.FilterField, placeholders *providerWebhookEndpointsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewProviderWebhookEndpointsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeProviderWebhookEndpointsFilterSQLExpr(spec)
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

func composeProviderWebhookEndpointsFilterGroup(group model.FilterGroup, placeholders *providerWebhookEndpointsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeProviderWebhookEndpointsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeProviderWebhookEndpointsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeProviderWebhookEndpointsFilterWhereQueries(filter model.Filter, placeholders *providerWebhookEndpointsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeProviderWebhookEndpointsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeProviderWebhookEndpointsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeProviderWebhookEndpointsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateProviderWebhookEndpointsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeProviderWebhookEndpointsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeProviderWebhookEndpointsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := providerWebhookEndpointsFilterPlaceholder{index: 1}
	whereQueries, err := composeProviderWebhookEndpointsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewProviderWebhookEndpointsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeProviderWebhookEndpointsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeProviderWebhookEndpointsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"provider_webhook_endpoints\" base%s", strings.Join(selectColumns, ","), composeProviderWebhookEndpointsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistProviderWebhookEndpointsByID(ctx context.Context, primaryID model.ProviderWebhookEndpointsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeProviderWebhookEndpointsCompositePrimaryKeyWhere([]model.ProviderWebhookEndpointsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", providerWebhookEndpointsQueries.selectCountProviderWebhookEndpoints, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderWebhookEndpointsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderWebhookEndpoints(ctx context.Context, selectFields ...ProviderWebhookEndpointsField) (providerWebhookEndpointsList model.ProviderWebhookEndpointsList, err error) {
	var (
		defaultProviderWebhookEndpointsSelectFields = defaultProviderWebhookEndpointsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderWebhookEndpointsSelectFields = composeProviderWebhookEndpointsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(providerWebhookEndpointsQueries.selectProviderWebhookEndpoints, defaultProviderWebhookEndpointsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &providerWebhookEndpointsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderWebhookEndpoints] failed get providerWebhookEndpoints list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderWebhookEndpointsByID(ctx context.Context, primaryID model.ProviderWebhookEndpointsPrimaryID, selectFields ...ProviderWebhookEndpointsField) (providerWebhookEndpoints model.ProviderWebhookEndpoints, err error) {
	var (
		defaultProviderWebhookEndpointsSelectFields = defaultProviderWebhookEndpointsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderWebhookEndpointsSelectFields = composeProviderWebhookEndpointsSelectFields(selectFields...)
	}
	whereQry, params := composeProviderWebhookEndpointsCompositePrimaryKeyWhere([]model.ProviderWebhookEndpointsPrimaryID{primaryID})
	query := fmt.Sprintf(providerWebhookEndpointsQueries.selectProviderWebhookEndpoints+" WHERE "+whereQry, defaultProviderWebhookEndpointsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &providerWebhookEndpoints, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("providerWebhookEndpoints with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveProviderWebhookEndpointsByID] failed get providerWebhookEndpoints")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateProviderWebhookEndpointsByID(ctx context.Context, primaryID model.ProviderWebhookEndpointsPrimaryID, providerWebhookEndpoints *model.ProviderWebhookEndpoints, providerWebhookEndpointsUpdateFields ...ProviderWebhookEndpointsUpdateField) (err error) {
	exists, err := repo.IsExistProviderWebhookEndpointsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEndpoints] failed checking providerWebhookEndpoints whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerWebhookEndpoints with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if providerWebhookEndpoints == nil {
		if len(providerWebhookEndpointsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateProviderWebhookEndpointsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		providerWebhookEndpoints = &model.ProviderWebhookEndpoints{}
	}
	var (
		defaultProviderWebhookEndpointsUpdateFields = defaultProviderWebhookEndpointsUpdateFields(*providerWebhookEndpoints)
		tempUpdateField                             ProviderWebhookEndpointsUpdateFieldList
		selectFields                                = NewProviderWebhookEndpointsSelectFields()
	)
	if len(providerWebhookEndpointsUpdateFields) > 0 {
		for _, updateField := range providerWebhookEndpointsUpdateFields {
			if updateField.providerWebhookEndpointsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultProviderWebhookEndpointsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeProviderWebhookEndpointsCompositePrimaryKeyWhere([]model.ProviderWebhookEndpointsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsProviderWebhookEndpointsCommand(defaultProviderWebhookEndpointsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(providerWebhookEndpointsQueries.updateProviderWebhookEndpoints+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEndpoints] error when try to update providerWebhookEndpoints by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateProviderWebhookEndpointsByFilter(ctx context.Context, filter model.Filter, providerWebhookEndpointsUpdateFields ...ProviderWebhookEndpointsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(providerWebhookEndpointsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ProviderWebhookEndpointsUpdateFieldList
		selectFields = NewProviderWebhookEndpointsSelectFields()
	)
	for _, updateField := range providerWebhookEndpointsUpdateFields {
		if updateField.providerWebhookEndpointsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsProviderWebhookEndpointsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := providerWebhookEndpointsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeProviderWebhookEndpointsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"provider_webhook_endpoints\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEndpointsByFilter] error when try to update providerWebhookEndpoints by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderWebhookEndpointsByFilter] failed get rows affected")
	}
	return
}

var (
	providerWebhookEndpointsQueries = struct {
		selectProviderWebhookEndpoints      string
		selectCountProviderWebhookEndpoints string
		deleteProviderWebhookEndpoints      string
		updateProviderWebhookEndpoints      string
		insertProviderWebhookEndpoints      string
	}{
		selectProviderWebhookEndpoints:      "SELECT %s FROM \"provider_webhook_endpoints\"",
		selectCountProviderWebhookEndpoints: "SELECT COUNT(\"id\") FROM \"provider_webhook_endpoints\"",
		deleteProviderWebhookEndpoints:      "DELETE FROM \"provider_webhook_endpoints\"",
		updateProviderWebhookEndpoints:      "UPDATE \"provider_webhook_endpoints\" SET %s ",
		insertProviderWebhookEndpoints:      "INSERT INTO \"provider_webhook_endpoints\" %s VALUES %s",
	}
)

type ProviderWebhookEndpointsRepository interface {
	CreateProviderWebhookEndpoints(ctx context.Context, providerWebhookEndpoints *model.ProviderWebhookEndpoints, fieldsInsert ...ProviderWebhookEndpointsField) error
	BulkCreateProviderWebhookEndpoints(ctx context.Context, providerWebhookEndpointsList []*model.ProviderWebhookEndpoints, fieldsInsert ...ProviderWebhookEndpointsField) error
	ResolveProviderWebhookEndpoints(ctx context.Context, selectFields ...ProviderWebhookEndpointsField) (model.ProviderWebhookEndpointsList, error)
	ResolveProviderWebhookEndpointsByID(ctx context.Context, primaryID model.ProviderWebhookEndpointsPrimaryID, selectFields ...ProviderWebhookEndpointsField) (model.ProviderWebhookEndpoints, error)
	UpdateProviderWebhookEndpointsByID(ctx context.Context, id model.ProviderWebhookEndpointsPrimaryID, providerWebhookEndpoints *model.ProviderWebhookEndpoints, providerWebhookEndpointsUpdateFields ...ProviderWebhookEndpointsUpdateField) error
	UpdateProviderWebhookEndpointsByFilter(ctx context.Context, filter model.Filter, providerWebhookEndpointsUpdateFields ...ProviderWebhookEndpointsUpdateField) (rowsAffected int64, err error)
	BulkUpdateProviderWebhookEndpoints(ctx context.Context, providerWebhookEndpointsListMap map[model.ProviderWebhookEndpointsPrimaryID]*model.ProviderWebhookEndpoints, ProviderWebhookEndpointssMapUpdateFieldsRequest map[model.ProviderWebhookEndpointsPrimaryID]ProviderWebhookEndpointsUpdateFieldList) (err error)
	DeleteProviderWebhookEndpointsByID(ctx context.Context, id model.ProviderWebhookEndpointsPrimaryID) error
	BulkDeleteProviderWebhookEndpointsByIDs(ctx context.Context, ids []model.ProviderWebhookEndpointsPrimaryID) error
	ResolveProviderWebhookEndpointsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderWebhookEndpointsFilterResult, err error)
	IsExistProviderWebhookEndpointsByIDs(ctx context.Context, ids []model.ProviderWebhookEndpointsPrimaryID) (exists bool, notFoundIds []model.ProviderWebhookEndpointsPrimaryID, err error)
	IsExistProviderWebhookEndpointsByID(ctx context.Context, id model.ProviderWebhookEndpointsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
