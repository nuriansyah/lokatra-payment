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

func composeInsertFieldsAndParamsProviderAccounts(providerAccountsList []model.ProviderAccounts, fieldsInsert ...ProviderAccountsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewProviderAccountsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, providerAccounts := range providerAccountsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, providerAccounts.Id)
			case selectField.ProviderId():
				args = append(args, providerAccounts.ProviderId)
			case selectField.AccountName():
				args = append(args, providerAccounts.AccountName)
			case selectField.Environment():
				args = append(args, providerAccounts.Environment)
			case selectField.OwnerType():
				args = append(args, providerAccounts.OwnerType)
			case selectField.OwnerId():
				args = append(args, providerAccounts.OwnerId)
			case selectField.MerchantRef():
				args = append(args, providerAccounts.MerchantRef)
			case selectField.CredentialSecretRef():
				args = append(args, providerAccounts.CredentialSecretRef)
			case selectField.WebhookSecretRef():
				args = append(args, providerAccounts.WebhookSecretRef)
			case selectField.PublicKeyRef():
				args = append(args, providerAccounts.PublicKeyRef)
			case selectField.Status():
				args = append(args, providerAccounts.Status)
			case selectField.Config():
				args = append(args, providerAccounts.Config)
			case selectField.Metadata():
				args = append(args, providerAccounts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, providerAccounts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, providerAccounts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, providerAccounts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, providerAccounts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, providerAccounts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, providerAccounts.MetaDeletedBy)

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

func composeProviderAccountsCompositePrimaryKeyWhere(primaryIDs []model.ProviderAccountsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"provider_accounts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultProviderAccountsSelectFields() string {
	fields := NewProviderAccountsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeProviderAccountsSelectFields(selectFields ...ProviderAccountsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ProviderAccountsField string
type ProviderAccountsFieldList []ProviderAccountsField

type ProviderAccountsSelectFields struct {
}

func (ss ProviderAccountsSelectFields) Id() ProviderAccountsField {
	return ProviderAccountsField("id")
}

func (ss ProviderAccountsSelectFields) ProviderId() ProviderAccountsField {
	return ProviderAccountsField("provider_id")
}

func (ss ProviderAccountsSelectFields) AccountName() ProviderAccountsField {
	return ProviderAccountsField("account_name")
}

func (ss ProviderAccountsSelectFields) Environment() ProviderAccountsField {
	return ProviderAccountsField("environment")
}

func (ss ProviderAccountsSelectFields) OwnerType() ProviderAccountsField {
	return ProviderAccountsField("owner_type")
}

func (ss ProviderAccountsSelectFields) OwnerId() ProviderAccountsField {
	return ProviderAccountsField("owner_id")
}

func (ss ProviderAccountsSelectFields) MerchantRef() ProviderAccountsField {
	return ProviderAccountsField("merchant_ref")
}

func (ss ProviderAccountsSelectFields) CredentialSecretRef() ProviderAccountsField {
	return ProviderAccountsField("credential_secret_ref")
}

func (ss ProviderAccountsSelectFields) WebhookSecretRef() ProviderAccountsField {
	return ProviderAccountsField("webhook_secret_ref")
}

func (ss ProviderAccountsSelectFields) PublicKeyRef() ProviderAccountsField {
	return ProviderAccountsField("public_key_ref")
}

func (ss ProviderAccountsSelectFields) Status() ProviderAccountsField {
	return ProviderAccountsField("status")
}

func (ss ProviderAccountsSelectFields) Config() ProviderAccountsField {
	return ProviderAccountsField("config")
}

func (ss ProviderAccountsSelectFields) Metadata() ProviderAccountsField {
	return ProviderAccountsField("metadata")
}

func (ss ProviderAccountsSelectFields) MetaCreatedAt() ProviderAccountsField {
	return ProviderAccountsField("meta_created_at")
}

func (ss ProviderAccountsSelectFields) MetaCreatedBy() ProviderAccountsField {
	return ProviderAccountsField("meta_created_by")
}

func (ss ProviderAccountsSelectFields) MetaUpdatedAt() ProviderAccountsField {
	return ProviderAccountsField("meta_updated_at")
}

func (ss ProviderAccountsSelectFields) MetaUpdatedBy() ProviderAccountsField {
	return ProviderAccountsField("meta_updated_by")
}

func (ss ProviderAccountsSelectFields) MetaDeletedAt() ProviderAccountsField {
	return ProviderAccountsField("meta_deleted_at")
}

func (ss ProviderAccountsSelectFields) MetaDeletedBy() ProviderAccountsField {
	return ProviderAccountsField("meta_deleted_by")
}

func (ss ProviderAccountsSelectFields) All() ProviderAccountsFieldList {
	return []ProviderAccountsField{
		ss.Id(),
		ss.ProviderId(),
		ss.AccountName(),
		ss.Environment(),
		ss.OwnerType(),
		ss.OwnerId(),
		ss.MerchantRef(),
		ss.CredentialSecretRef(),
		ss.WebhookSecretRef(),
		ss.PublicKeyRef(),
		ss.Status(),
		ss.Config(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewProviderAccountsSelectFields() ProviderAccountsSelectFields {
	return ProviderAccountsSelectFields{}
}

type ProviderAccountsUpdateFieldOption struct {
	useIncrement bool
}
type ProviderAccountsUpdateField struct {
	providerAccountsField ProviderAccountsField
	opt                   ProviderAccountsUpdateFieldOption
	value                 interface{}
}
type ProviderAccountsUpdateFieldList []ProviderAccountsUpdateField

func defaultProviderAccountsUpdateFieldOption() ProviderAccountsUpdateFieldOption {
	return ProviderAccountsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementProviderAccountsOption(useIncrement bool) func(*ProviderAccountsUpdateFieldOption) {
	return func(pcufo *ProviderAccountsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewProviderAccountsUpdateField(field ProviderAccountsField, val interface{}, opts ...func(*ProviderAccountsUpdateFieldOption)) ProviderAccountsUpdateField {
	defaultOpt := defaultProviderAccountsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ProviderAccountsUpdateField{
		providerAccountsField: field,
		value:                 val,
		opt:                   defaultOpt,
	}
}
func defaultProviderAccountsUpdateFields(providerAccounts model.ProviderAccounts) (providerAccountsUpdateFieldList ProviderAccountsUpdateFieldList) {
	selectFields := NewProviderAccountsSelectFields()
	providerAccountsUpdateFieldList = append(providerAccountsUpdateFieldList,
		NewProviderAccountsUpdateField(selectFields.Id(), providerAccounts.Id),
		NewProviderAccountsUpdateField(selectFields.ProviderId(), providerAccounts.ProviderId),
		NewProviderAccountsUpdateField(selectFields.AccountName(), providerAccounts.AccountName),
		NewProviderAccountsUpdateField(selectFields.Environment(), providerAccounts.Environment),
		NewProviderAccountsUpdateField(selectFields.OwnerType(), providerAccounts.OwnerType),
		NewProviderAccountsUpdateField(selectFields.OwnerId(), providerAccounts.OwnerId),
		NewProviderAccountsUpdateField(selectFields.MerchantRef(), providerAccounts.MerchantRef),
		NewProviderAccountsUpdateField(selectFields.CredentialSecretRef(), providerAccounts.CredentialSecretRef),
		NewProviderAccountsUpdateField(selectFields.WebhookSecretRef(), providerAccounts.WebhookSecretRef),
		NewProviderAccountsUpdateField(selectFields.PublicKeyRef(), providerAccounts.PublicKeyRef),
		NewProviderAccountsUpdateField(selectFields.Status(), providerAccounts.Status),
		NewProviderAccountsUpdateField(selectFields.Config(), providerAccounts.Config),
		NewProviderAccountsUpdateField(selectFields.Metadata(), providerAccounts.Metadata),
		NewProviderAccountsUpdateField(selectFields.MetaCreatedAt(), providerAccounts.MetaCreatedAt),
		NewProviderAccountsUpdateField(selectFields.MetaCreatedBy(), providerAccounts.MetaCreatedBy),
		NewProviderAccountsUpdateField(selectFields.MetaUpdatedAt(), providerAccounts.MetaUpdatedAt),
		NewProviderAccountsUpdateField(selectFields.MetaUpdatedBy(), providerAccounts.MetaUpdatedBy),
		NewProviderAccountsUpdateField(selectFields.MetaDeletedAt(), providerAccounts.MetaDeletedAt),
		NewProviderAccountsUpdateField(selectFields.MetaDeletedBy(), providerAccounts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsProviderAccountsCommand(providerAccountsUpdateFieldList ProviderAccountsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range providerAccountsUpdateFieldList {
		field := string(updateField.providerAccountsField)
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

func (repo *RepositoryImpl) BulkCreateProviderAccounts(ctx context.Context, providerAccountsList []*model.ProviderAccounts, fieldsInsert ...ProviderAccountsField) (err error) {
	var (
		fieldsStr                 string
		valueListStr              []string
		argsList                  []interface{}
		primaryIds                []model.ProviderAccountsPrimaryID
		providerAccountsValueList []model.ProviderAccounts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewProviderAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, providerAccounts := range providerAccountsList {

		primaryIds = append(primaryIds, providerAccounts.ToProviderAccountsPrimaryID())

		providerAccountsValueList = append(providerAccountsValueList, *providerAccounts)
	}

	_, notFoundIds, err := repo.IsExistProviderAccountsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderAccounts] failed checking providerAccounts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ProviderAccountsPrimaryID{}
		mapNotFoundIds := map[model.ProviderAccountsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "providerAccounts", fmt.Sprintf("providerAccounts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsProviderAccounts(providerAccountsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(providerAccountsQueries.insertProviderAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderAccounts] failed exec create providerAccounts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteProviderAccountsByIDs(ctx context.Context, primaryIDs []model.ProviderAccountsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistProviderAccountsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderAccountsByIDs] failed checking providerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerAccounts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_accounts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(providerAccountsQueries.deleteProviderAccounts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderAccountsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderAccountsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistProviderAccountsByIDs(ctx context.Context, ids []model.ProviderAccountsPrimaryID) (exists bool, notFoundIds []model.ProviderAccountsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_accounts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(providerAccountsQueries.selectProviderAccounts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderAccountsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ProviderAccountsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderAccountsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ProviderAccountsPrimaryID]bool{}
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

// BulkUpdateProviderAccounts is used to bulk update providerAccounts, by default it will update all field
// if want to update specific field, then fill providerAccountssMapUpdateFieldsRequest else please fill providerAccountssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateProviderAccounts(ctx context.Context, providerAccountssMap map[model.ProviderAccountsPrimaryID]*model.ProviderAccounts, providerAccountssMapUpdateFieldsRequest map[model.ProviderAccountsPrimaryID]ProviderAccountsUpdateFieldList) (err error) {
	if len(providerAccountssMap) == 0 && len(providerAccountssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		providerAccountssMapUpdateField map[model.ProviderAccountsPrimaryID]ProviderAccountsUpdateFieldList = map[model.ProviderAccountsPrimaryID]ProviderAccountsUpdateFieldList{}
		asTableValues                   string                                                              = "myvalues"
	)

	if len(providerAccountssMap) > 0 {
		for id, providerAccounts := range providerAccountssMap {
			if providerAccounts == nil {
				log.Error().Err(err).Msg("[BulkUpdateProviderAccounts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			providerAccountssMapUpdateField[id] = defaultProviderAccountsUpdateFields(*providerAccounts)
		}
	} else {
		providerAccountssMapUpdateField = providerAccountssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateProviderAccountsQuery(providerAccountssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistProviderAccountsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderAccounts] failed checking providerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerAccounts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeProviderAccountsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"provider_accounts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderAccounts] failed exec query")
	}
	return
}

type ProviderAccountsFieldParameter struct {
	param string
	args  []interface{}
}

func NewProviderAccountsFieldParameter(param string, args ...interface{}) ProviderAccountsFieldParameter {
	return ProviderAccountsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateProviderAccountsQuery(mapProviderAccountss map[model.ProviderAccountsPrimaryID]ProviderAccountsUpdateFieldList, asTableValues string) (primaryIDs []model.ProviderAccountsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ProviderAccountsPrimaryID]map[string]interface{}{}
	providerAccountsSelectFields := NewProviderAccountsSelectFields()
	for id, updateFields := range mapProviderAccountss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.providerAccountsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapProviderAccountss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetProviderAccountsFieldType(updateField.providerAccountsField)))
			args = append(args, fields[string(updateField.providerAccountsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.providerAccountsField))
		if updateField.providerAccountsField == providerAccountsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.providerAccountsField, asTableValues, updateField.providerAccountsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.providerAccountsField,
				"\"provider_accounts\"", updateField.providerAccountsField,
				asTableValues, updateField.providerAccountsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeProviderAccountsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ProviderAccountsPrimaryID, asTableValue string) (whereQry string) {
	providerAccountsSelectFields := NewProviderAccountsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"provider_accounts\".\"id\" = %s.\"id\"::"+GetProviderAccountsFieldType(providerAccountsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetProviderAccountsFieldType(providerAccountsField ProviderAccountsField) string {
	selectProviderAccountsFields := NewProviderAccountsSelectFields()
	switch providerAccountsField {

	case selectProviderAccountsFields.Id():
		return "uuid"

	case selectProviderAccountsFields.ProviderId():
		return "uuid"

	case selectProviderAccountsFields.AccountName():
		return "text"

	case selectProviderAccountsFields.Environment():
		return "text"

	case selectProviderAccountsFields.OwnerType():
		return "text"

	case selectProviderAccountsFields.OwnerId():
		return "uuid"

	case selectProviderAccountsFields.MerchantRef():
		return "text"

	case selectProviderAccountsFields.CredentialSecretRef():
		return "text"

	case selectProviderAccountsFields.WebhookSecretRef():
		return "text"

	case selectProviderAccountsFields.PublicKeyRef():
		return "text"

	case selectProviderAccountsFields.Status():
		return "provider_account_status_enum"

	case selectProviderAccountsFields.Config():
		return "jsonb"

	case selectProviderAccountsFields.Metadata():
		return "jsonb"

	case selectProviderAccountsFields.MetaCreatedAt():
		return "timestamptz"

	case selectProviderAccountsFields.MetaCreatedBy():
		return "uuid"

	case selectProviderAccountsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectProviderAccountsFields.MetaUpdatedBy():
		return "uuid"

	case selectProviderAccountsFields.MetaDeletedAt():
		return "timestamptz"

	case selectProviderAccountsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateProviderAccounts(ctx context.Context, providerAccounts *model.ProviderAccounts, fieldsInsert ...ProviderAccountsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewProviderAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ProviderAccountsPrimaryID{
		Id: providerAccounts.Id,
	}
	exists, err := repo.IsExistProviderAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderAccounts] failed checking providerAccounts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "providerAccounts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsProviderAccounts([]model.ProviderAccounts{*providerAccounts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(providerAccountsQueries.insertProviderAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderAccounts] failed exec create providerAccounts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteProviderAccountsByID(ctx context.Context, primaryID model.ProviderAccountsPrimaryID) (err error) {
	exists, err := repo.IsExistProviderAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderAccountsByID] failed checking providerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeProviderAccountsCompositePrimaryKeyWhere([]model.ProviderAccountsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(providerAccountsQueries.deleteProviderAccounts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderAccountsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderAccountsFilterResult, err error) {
	query, args, err := composeProviderAccountsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderAccountsByFilter] failed compose providerAccounts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderAccountsByFilter] failed get providerAccounts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeProviderAccountsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ProviderAccountsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeProviderAccountsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeProviderAccountsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeProviderAccountsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewProviderAccountsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["provider_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_id\"")
			selectedColumns["provider_id"] = struct{}{}
		}
		if _, selected := selectedColumns["account_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_name\"")
			selectedColumns["account_name"] = struct{}{}
		}
		if _, selected := selectedColumns["environment"]; !selected {
			selectColumns = append(selectColumns, "base.\"environment\"")
			selectedColumns["environment"] = struct{}{}
		}
		if _, selected := selectedColumns["owner_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"owner_type\"")
			selectedColumns["owner_type"] = struct{}{}
		}
		if _, selected := selectedColumns["owner_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"owner_id\"")
			selectedColumns["owner_id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_ref\"")
			selectedColumns["merchant_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["credential_secret_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"credential_secret_ref\"")
			selectedColumns["credential_secret_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["webhook_secret_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"webhook_secret_ref\"")
			selectedColumns["webhook_secret_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["public_key_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"public_key_ref\"")
			selectedColumns["public_key_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["config"]; !selected {
			selectColumns = append(selectColumns, "base.\"config\"")
			selectedColumns["config"] = struct{}{}
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

type providerAccountsFilterPlaceholder struct {
	index int
}

func (p *providerAccountsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeProviderAccountsFilterPredicate(filterField model.FilterField, placeholders *providerAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewProviderAccountsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeProviderAccountsFilterSQLExpr(spec)
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

func composeProviderAccountsFilterGroup(group model.FilterGroup, placeholders *providerAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeProviderAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeProviderAccountsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeProviderAccountsFilterWhereQueries(filter model.Filter, placeholders *providerAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeProviderAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeProviderAccountsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeProviderAccountsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateProviderAccountsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeProviderAccountsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeProviderAccountsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := providerAccountsFilterPlaceholder{index: 1}
	whereQueries, err := composeProviderAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewProviderAccountsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeProviderAccountsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeProviderAccountsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"provider_accounts\" base%s", strings.Join(selectColumns, ","), composeProviderAccountsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistProviderAccountsByID(ctx context.Context, primaryID model.ProviderAccountsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeProviderAccountsCompositePrimaryKeyWhere([]model.ProviderAccountsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", providerAccountsQueries.selectCountProviderAccounts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderAccountsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderAccounts(ctx context.Context, selectFields ...ProviderAccountsField) (providerAccountsList model.ProviderAccountsList, err error) {
	var (
		defaultProviderAccountsSelectFields = defaultProviderAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderAccountsSelectFields = composeProviderAccountsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(providerAccountsQueries.selectProviderAccounts, defaultProviderAccountsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &providerAccountsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderAccounts] failed get providerAccounts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderAccountsByID(ctx context.Context, primaryID model.ProviderAccountsPrimaryID, selectFields ...ProviderAccountsField) (providerAccounts model.ProviderAccounts, err error) {
	var (
		defaultProviderAccountsSelectFields = defaultProviderAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderAccountsSelectFields = composeProviderAccountsSelectFields(selectFields...)
	}
	whereQry, params := composeProviderAccountsCompositePrimaryKeyWhere([]model.ProviderAccountsPrimaryID{primaryID})
	query := fmt.Sprintf(providerAccountsQueries.selectProviderAccounts+" WHERE "+whereQry, defaultProviderAccountsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &providerAccounts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("providerAccounts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveProviderAccountsByID] failed get providerAccounts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateProviderAccountsByID(ctx context.Context, primaryID model.ProviderAccountsPrimaryID, providerAccounts *model.ProviderAccounts, providerAccountsUpdateFields ...ProviderAccountsUpdateField) (err error) {
	exists, err := repo.IsExistProviderAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderAccounts] failed checking providerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if providerAccounts == nil {
		if len(providerAccountsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateProviderAccountsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		providerAccounts = &model.ProviderAccounts{}
	}
	var (
		defaultProviderAccountsUpdateFields = defaultProviderAccountsUpdateFields(*providerAccounts)
		tempUpdateField                     ProviderAccountsUpdateFieldList
		selectFields                        = NewProviderAccountsSelectFields()
	)
	if len(providerAccountsUpdateFields) > 0 {
		for _, updateField := range providerAccountsUpdateFields {
			if updateField.providerAccountsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultProviderAccountsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeProviderAccountsCompositePrimaryKeyWhere([]model.ProviderAccountsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsProviderAccountsCommand(defaultProviderAccountsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(providerAccountsQueries.updateProviderAccounts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderAccounts] error when try to update providerAccounts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateProviderAccountsByFilter(ctx context.Context, filter model.Filter, providerAccountsUpdateFields ...ProviderAccountsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(providerAccountsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ProviderAccountsUpdateFieldList
		selectFields = NewProviderAccountsSelectFields()
	)
	for _, updateField := range providerAccountsUpdateFields {
		if updateField.providerAccountsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsProviderAccountsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := providerAccountsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeProviderAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"provider_accounts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderAccountsByFilter] error when try to update providerAccounts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderAccountsByFilter] failed get rows affected")
	}
	return
}

var (
	providerAccountsQueries = struct {
		selectProviderAccounts      string
		selectCountProviderAccounts string
		deleteProviderAccounts      string
		updateProviderAccounts      string
		insertProviderAccounts      string
	}{
		selectProviderAccounts:      "SELECT %s FROM \"provider_accounts\"",
		selectCountProviderAccounts: "SELECT COUNT(\"id\") FROM \"provider_accounts\"",
		deleteProviderAccounts:      "DELETE FROM \"provider_accounts\"",
		updateProviderAccounts:      "UPDATE \"provider_accounts\" SET %s ",
		insertProviderAccounts:      "INSERT INTO \"provider_accounts\" %s VALUES %s",
	}
)

type ProviderAccountsRepository interface {
	CreateProviderAccounts(ctx context.Context, providerAccounts *model.ProviderAccounts, fieldsInsert ...ProviderAccountsField) error
	BulkCreateProviderAccounts(ctx context.Context, providerAccountsList []*model.ProviderAccounts, fieldsInsert ...ProviderAccountsField) error
	ResolveProviderAccounts(ctx context.Context, selectFields ...ProviderAccountsField) (model.ProviderAccountsList, error)
	ResolveProviderAccountsByID(ctx context.Context, primaryID model.ProviderAccountsPrimaryID, selectFields ...ProviderAccountsField) (model.ProviderAccounts, error)
	UpdateProviderAccountsByID(ctx context.Context, id model.ProviderAccountsPrimaryID, providerAccounts *model.ProviderAccounts, providerAccountsUpdateFields ...ProviderAccountsUpdateField) error
	UpdateProviderAccountsByFilter(ctx context.Context, filter model.Filter, providerAccountsUpdateFields ...ProviderAccountsUpdateField) (rowsAffected int64, err error)
	BulkUpdateProviderAccounts(ctx context.Context, providerAccountsListMap map[model.ProviderAccountsPrimaryID]*model.ProviderAccounts, ProviderAccountssMapUpdateFieldsRequest map[model.ProviderAccountsPrimaryID]ProviderAccountsUpdateFieldList) (err error)
	DeleteProviderAccountsByID(ctx context.Context, id model.ProviderAccountsPrimaryID) error
	BulkDeleteProviderAccountsByIDs(ctx context.Context, ids []model.ProviderAccountsPrimaryID) error
	ResolveProviderAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderAccountsFilterResult, err error)
	IsExistProviderAccountsByIDs(ctx context.Context, ids []model.ProviderAccountsPrimaryID) (exists bool, notFoundIds []model.ProviderAccountsPrimaryID, err error)
	IsExistProviderAccountsByID(ctx context.Context, id model.ProviderAccountsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
