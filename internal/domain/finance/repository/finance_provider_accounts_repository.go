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

func composeInsertFieldsAndParamsFinanceProviderAccounts(financeProviderAccountsList []model.FinanceProviderAccounts, fieldsInsert ...FinanceProviderAccountsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceProviderAccountsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeProviderAccounts := range financeProviderAccountsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeProviderAccounts.Id)
			case selectField.ProviderCode():
				args = append(args, financeProviderAccounts.ProviderCode)
			case selectField.ProviderName():
				args = append(args, financeProviderAccounts.ProviderName)
			case selectField.OwnerPartyId():
				args = append(args, financeProviderAccounts.OwnerPartyId)
			case selectField.Environment():
				args = append(args, financeProviderAccounts.Environment)
			case selectField.ApiBaseUrl():
				args = append(args, financeProviderAccounts.ApiBaseUrl)
			case selectField.MerchantRef():
				args = append(args, financeProviderAccounts.MerchantRef)
			case selectField.SettlementCurrency():
				args = append(args, financeProviderAccounts.SettlementCurrency)
			case selectField.VaultSecretRef():
				args = append(args, financeProviderAccounts.VaultSecretRef)
			case selectField.WebhookSecretRef():
				args = append(args, financeProviderAccounts.WebhookSecretRef)
			case selectField.ProviderStatus():
				args = append(args, financeProviderAccounts.ProviderStatus)
			case selectField.Capabilities():
				args = append(args, financeProviderAccounts.Capabilities)
			case selectField.Metadata():
				args = append(args, financeProviderAccounts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeProviderAccounts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeProviderAccounts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeProviderAccounts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeProviderAccounts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeProviderAccounts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeProviderAccounts.MetaDeletedBy)

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

func composeFinanceProviderAccountsCompositePrimaryKeyWhere(primaryIDs []model.FinanceProviderAccountsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_provider_accounts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceProviderAccountsSelectFields() string {
	fields := NewFinanceProviderAccountsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceProviderAccountsSelectFields(selectFields ...FinanceProviderAccountsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceProviderAccountsField string
type FinanceProviderAccountsFieldList []FinanceProviderAccountsField

type FinanceProviderAccountsSelectFields struct {
}

func (ss FinanceProviderAccountsSelectFields) Id() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("id")
}

func (ss FinanceProviderAccountsSelectFields) ProviderCode() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("provider_code")
}

func (ss FinanceProviderAccountsSelectFields) ProviderName() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("provider_name")
}

func (ss FinanceProviderAccountsSelectFields) OwnerPartyId() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("owner_party_id")
}

func (ss FinanceProviderAccountsSelectFields) Environment() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("environment")
}

func (ss FinanceProviderAccountsSelectFields) ApiBaseUrl() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("api_base_url")
}

func (ss FinanceProviderAccountsSelectFields) MerchantRef() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("merchant_ref")
}

func (ss FinanceProviderAccountsSelectFields) SettlementCurrency() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("settlement_currency")
}

func (ss FinanceProviderAccountsSelectFields) VaultSecretRef() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("vault_secret_ref")
}

func (ss FinanceProviderAccountsSelectFields) WebhookSecretRef() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("webhook_secret_ref")
}

func (ss FinanceProviderAccountsSelectFields) ProviderStatus() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("provider_status")
}

func (ss FinanceProviderAccountsSelectFields) Capabilities() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("capabilities")
}

func (ss FinanceProviderAccountsSelectFields) Metadata() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("metadata")
}

func (ss FinanceProviderAccountsSelectFields) MetaCreatedAt() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("meta_created_at")
}

func (ss FinanceProviderAccountsSelectFields) MetaCreatedBy() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("meta_created_by")
}

func (ss FinanceProviderAccountsSelectFields) MetaUpdatedAt() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("meta_updated_at")
}

func (ss FinanceProviderAccountsSelectFields) MetaUpdatedBy() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("meta_updated_by")
}

func (ss FinanceProviderAccountsSelectFields) MetaDeletedAt() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("meta_deleted_at")
}

func (ss FinanceProviderAccountsSelectFields) MetaDeletedBy() FinanceProviderAccountsField {
	return FinanceProviderAccountsField("meta_deleted_by")
}

func (ss FinanceProviderAccountsSelectFields) All() FinanceProviderAccountsFieldList {
	return []FinanceProviderAccountsField{
		ss.Id(),
		ss.ProviderCode(),
		ss.ProviderName(),
		ss.OwnerPartyId(),
		ss.Environment(),
		ss.ApiBaseUrl(),
		ss.MerchantRef(),
		ss.SettlementCurrency(),
		ss.VaultSecretRef(),
		ss.WebhookSecretRef(),
		ss.ProviderStatus(),
		ss.Capabilities(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinanceProviderAccountsSelectFields() FinanceProviderAccountsSelectFields {
	return FinanceProviderAccountsSelectFields{}
}

type FinanceProviderAccountsUpdateFieldOption struct {
	useIncrement bool
}
type FinanceProviderAccountsUpdateField struct {
	financeProviderAccountsField FinanceProviderAccountsField
	opt                          FinanceProviderAccountsUpdateFieldOption
	value                        interface{}
}
type FinanceProviderAccountsUpdateFieldList []FinanceProviderAccountsUpdateField

func defaultFinanceProviderAccountsUpdateFieldOption() FinanceProviderAccountsUpdateFieldOption {
	return FinanceProviderAccountsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceProviderAccountsOption(useIncrement bool) func(*FinanceProviderAccountsUpdateFieldOption) {
	return func(pcufo *FinanceProviderAccountsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceProviderAccountsUpdateField(field FinanceProviderAccountsField, val interface{}, opts ...func(*FinanceProviderAccountsUpdateFieldOption)) FinanceProviderAccountsUpdateField {
	defaultOpt := defaultFinanceProviderAccountsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceProviderAccountsUpdateField{
		financeProviderAccountsField: field,
		value:                        val,
		opt:                          defaultOpt,
	}
}
func defaultFinanceProviderAccountsUpdateFields(financeProviderAccounts model.FinanceProviderAccounts) (financeProviderAccountsUpdateFieldList FinanceProviderAccountsUpdateFieldList) {
	selectFields := NewFinanceProviderAccountsSelectFields()
	financeProviderAccountsUpdateFieldList = append(financeProviderAccountsUpdateFieldList,
		NewFinanceProviderAccountsUpdateField(selectFields.Id(), financeProviderAccounts.Id),
		NewFinanceProviderAccountsUpdateField(selectFields.ProviderCode(), financeProviderAccounts.ProviderCode),
		NewFinanceProviderAccountsUpdateField(selectFields.ProviderName(), financeProviderAccounts.ProviderName),
		NewFinanceProviderAccountsUpdateField(selectFields.OwnerPartyId(), financeProviderAccounts.OwnerPartyId),
		NewFinanceProviderAccountsUpdateField(selectFields.Environment(), financeProviderAccounts.Environment),
		NewFinanceProviderAccountsUpdateField(selectFields.ApiBaseUrl(), financeProviderAccounts.ApiBaseUrl),
		NewFinanceProviderAccountsUpdateField(selectFields.MerchantRef(), financeProviderAccounts.MerchantRef),
		NewFinanceProviderAccountsUpdateField(selectFields.SettlementCurrency(), financeProviderAccounts.SettlementCurrency),
		NewFinanceProviderAccountsUpdateField(selectFields.VaultSecretRef(), financeProviderAccounts.VaultSecretRef),
		NewFinanceProviderAccountsUpdateField(selectFields.WebhookSecretRef(), financeProviderAccounts.WebhookSecretRef),
		NewFinanceProviderAccountsUpdateField(selectFields.ProviderStatus(), financeProviderAccounts.ProviderStatus),
		NewFinanceProviderAccountsUpdateField(selectFields.Capabilities(), financeProviderAccounts.Capabilities),
		NewFinanceProviderAccountsUpdateField(selectFields.Metadata(), financeProviderAccounts.Metadata),
		NewFinanceProviderAccountsUpdateField(selectFields.MetaCreatedAt(), financeProviderAccounts.MetaCreatedAt),
		NewFinanceProviderAccountsUpdateField(selectFields.MetaCreatedBy(), financeProviderAccounts.MetaCreatedBy),
		NewFinanceProviderAccountsUpdateField(selectFields.MetaUpdatedAt(), financeProviderAccounts.MetaUpdatedAt),
		NewFinanceProviderAccountsUpdateField(selectFields.MetaUpdatedBy(), financeProviderAccounts.MetaUpdatedBy),
		NewFinanceProviderAccountsUpdateField(selectFields.MetaDeletedAt(), financeProviderAccounts.MetaDeletedAt),
		NewFinanceProviderAccountsUpdateField(selectFields.MetaDeletedBy(), financeProviderAccounts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceProviderAccountsCommand(financeProviderAccountsUpdateFieldList FinanceProviderAccountsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeProviderAccountsUpdateFieldList {
		field := string(updateField.financeProviderAccountsField)
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

func (repo *RepositoryImpl) BulkCreateFinanceProviderAccounts(ctx context.Context, financeProviderAccountsList []*model.FinanceProviderAccounts, fieldsInsert ...FinanceProviderAccountsField) (err error) {
	var (
		fieldsStr                        string
		valueListStr                     []string
		argsList                         []interface{}
		primaryIds                       []model.FinanceProviderAccountsPrimaryID
		financeProviderAccountsValueList []model.FinanceProviderAccounts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceProviderAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeProviderAccounts := range financeProviderAccountsList {

		primaryIds = append(primaryIds, financeProviderAccounts.ToFinanceProviderAccountsPrimaryID())

		financeProviderAccountsValueList = append(financeProviderAccountsValueList, *financeProviderAccounts)
	}

	_, notFoundIds, err := repo.IsExistFinanceProviderAccountsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceProviderAccounts] failed checking financeProviderAccounts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceProviderAccountsPrimaryID{}
		mapNotFoundIds := map[model.FinanceProviderAccountsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeProviderAccounts", fmt.Sprintf("financeProviderAccounts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceProviderAccounts(financeProviderAccountsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeProviderAccountsQueries.insertFinanceProviderAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceProviderAccounts] failed exec create financeProviderAccounts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceProviderAccountsByIDs(ctx context.Context, primaryIDs []model.FinanceProviderAccountsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceProviderAccountsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceProviderAccountsByIDs] failed checking financeProviderAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeProviderAccounts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_provider_accounts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeProviderAccountsQueries.deleteFinanceProviderAccounts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceProviderAccountsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceProviderAccountsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceProviderAccountsByIDs(ctx context.Context, ids []model.FinanceProviderAccountsPrimaryID) (exists bool, notFoundIds []model.FinanceProviderAccountsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_provider_accounts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeProviderAccountsQueries.selectFinanceProviderAccounts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceProviderAccountsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceProviderAccountsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceProviderAccountsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceProviderAccountsPrimaryID]bool{}
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

// BulkUpdateFinanceProviderAccounts is used to bulk update financeProviderAccounts, by default it will update all field
// if want to update specific field, then fill financeProviderAccountssMapUpdateFieldsRequest else please fill financeProviderAccountssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceProviderAccounts(ctx context.Context, financeProviderAccountssMap map[model.FinanceProviderAccountsPrimaryID]*model.FinanceProviderAccounts, financeProviderAccountssMapUpdateFieldsRequest map[model.FinanceProviderAccountsPrimaryID]FinanceProviderAccountsUpdateFieldList) (err error) {
	if len(financeProviderAccountssMap) == 0 && len(financeProviderAccountssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeProviderAccountssMapUpdateField map[model.FinanceProviderAccountsPrimaryID]FinanceProviderAccountsUpdateFieldList = map[model.FinanceProviderAccountsPrimaryID]FinanceProviderAccountsUpdateFieldList{}
		asTableValues                          string                                                                            = "myvalues"
	)

	if len(financeProviderAccountssMap) > 0 {
		for id, financeProviderAccounts := range financeProviderAccountssMap {
			if financeProviderAccounts == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceProviderAccounts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeProviderAccountssMapUpdateField[id] = defaultFinanceProviderAccountsUpdateFields(*financeProviderAccounts)
		}
	} else {
		financeProviderAccountssMapUpdateField = financeProviderAccountssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceProviderAccountsQuery(financeProviderAccountssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceProviderAccountsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceProviderAccounts] failed checking financeProviderAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeProviderAccounts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceProviderAccountsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_provider_accounts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceProviderAccounts] failed exec query")
	}
	return
}

type FinanceProviderAccountsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceProviderAccountsFieldParameter(param string, args ...interface{}) FinanceProviderAccountsFieldParameter {
	return FinanceProviderAccountsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceProviderAccountsQuery(mapFinanceProviderAccountss map[model.FinanceProviderAccountsPrimaryID]FinanceProviderAccountsUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceProviderAccountsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceProviderAccountsPrimaryID]map[string]interface{}{}
	financeProviderAccountsSelectFields := NewFinanceProviderAccountsSelectFields()
	for id, updateFields := range mapFinanceProviderAccountss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeProviderAccountsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceProviderAccountss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceProviderAccountsFieldType(updateField.financeProviderAccountsField)))
			args = append(args, fields[string(updateField.financeProviderAccountsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeProviderAccountsField))
		if updateField.financeProviderAccountsField == financeProviderAccountsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeProviderAccountsField, asTableValues, updateField.financeProviderAccountsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeProviderAccountsField,
				"\"finance_provider_accounts\"", updateField.financeProviderAccountsField,
				asTableValues, updateField.financeProviderAccountsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceProviderAccountsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceProviderAccountsPrimaryID, asTableValue string) (whereQry string) {
	financeProviderAccountsSelectFields := NewFinanceProviderAccountsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_provider_accounts\".\"id\" = %s.\"id\"::"+GetFinanceProviderAccountsFieldType(financeProviderAccountsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceProviderAccountsFieldType(financeProviderAccountsField FinanceProviderAccountsField) string {
	selectFinanceProviderAccountsFields := NewFinanceProviderAccountsSelectFields()
	switch financeProviderAccountsField {

	case selectFinanceProviderAccountsFields.Id():
		return "uuid"

	case selectFinanceProviderAccountsFields.ProviderCode():
		return "text"

	case selectFinanceProviderAccountsFields.ProviderName():
		return "text"

	case selectFinanceProviderAccountsFields.OwnerPartyId():
		return "uuid"

	case selectFinanceProviderAccountsFields.Environment():
		return "environment_enum"

	case selectFinanceProviderAccountsFields.ApiBaseUrl():
		return "text"

	case selectFinanceProviderAccountsFields.MerchantRef():
		return "text"

	case selectFinanceProviderAccountsFields.SettlementCurrency():
		return "text"

	case selectFinanceProviderAccountsFields.VaultSecretRef():
		return "text"

	case selectFinanceProviderAccountsFields.WebhookSecretRef():
		return "text"

	case selectFinanceProviderAccountsFields.ProviderStatus():
		return "provider_status_enum"

	case selectFinanceProviderAccountsFields.Capabilities():
		return "jsonb"

	case selectFinanceProviderAccountsFields.Metadata():
		return "jsonb"

	case selectFinanceProviderAccountsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceProviderAccountsFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceProviderAccountsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceProviderAccountsFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceProviderAccountsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceProviderAccountsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceProviderAccounts(ctx context.Context, financeProviderAccounts *model.FinanceProviderAccounts, fieldsInsert ...FinanceProviderAccountsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceProviderAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceProviderAccountsPrimaryID{
		Id: financeProviderAccounts.Id,
	}
	exists, err := repo.IsExistFinanceProviderAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceProviderAccounts] failed checking financeProviderAccounts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeProviderAccounts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceProviderAccounts([]model.FinanceProviderAccounts{*financeProviderAccounts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeProviderAccountsQueries.insertFinanceProviderAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceProviderAccounts] failed exec create financeProviderAccounts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceProviderAccountsByID(ctx context.Context, primaryID model.FinanceProviderAccountsPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceProviderAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceProviderAccountsByID] failed checking financeProviderAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeProviderAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceProviderAccountsCompositePrimaryKeyWhere([]model.FinanceProviderAccountsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeProviderAccountsQueries.deleteFinanceProviderAccounts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceProviderAccountsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceProviderAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceProviderAccountsFilterResult, err error) {
	query, args, err := composeFinanceProviderAccountsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceProviderAccountsByFilter] failed compose financeProviderAccounts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceProviderAccountsByFilter] failed get financeProviderAccounts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceProviderAccountsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceProviderAccountsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceProviderAccountsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceProviderAccountsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceProviderAccountsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceProviderAccountsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_code\"")
			selectedColumns["provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_name\"")
			selectedColumns["provider_name"] = struct{}{}
		}
		if _, selected := selectedColumns["owner_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"owner_party_id\"")
			selectedColumns["owner_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["environment"]; !selected {
			selectColumns = append(selectColumns, "base.\"environment\"")
			selectedColumns["environment"] = struct{}{}
		}
		if _, selected := selectedColumns["api_base_url"]; !selected {
			selectColumns = append(selectColumns, "base.\"api_base_url\"")
			selectedColumns["api_base_url"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_ref\"")
			selectedColumns["merchant_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["settlement_currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"settlement_currency\"")
			selectedColumns["settlement_currency"] = struct{}{}
		}
		if _, selected := selectedColumns["vault_secret_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"vault_secret_ref\"")
			selectedColumns["vault_secret_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["webhook_secret_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"webhook_secret_ref\"")
			selectedColumns["webhook_secret_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_status\"")
			selectedColumns["provider_status"] = struct{}{}
		}
		if _, selected := selectedColumns["capabilities"]; !selected {
			selectColumns = append(selectColumns, "base.\"capabilities\"")
			selectedColumns["capabilities"] = struct{}{}
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

type financeProviderAccountsFilterPlaceholder struct {
	index int
}

func (p *financeProviderAccountsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceProviderAccountsFilterPredicate(filterField model.FilterField, placeholders *financeProviderAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceProviderAccountsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceProviderAccountsFilterSQLExpr(spec)
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

func composeFinanceProviderAccountsFilterGroup(group model.FilterGroup, placeholders *financeProviderAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceProviderAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceProviderAccountsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceProviderAccountsFilterWhereQueries(filter model.Filter, placeholders *financeProviderAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceProviderAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceProviderAccountsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceProviderAccountsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceProviderAccountsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceProviderAccountsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceProviderAccountsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeProviderAccountsFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceProviderAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceProviderAccountsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceProviderAccountsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceProviderAccountsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_provider_accounts\" base%s", strings.Join(selectColumns, ","), composeFinanceProviderAccountsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceProviderAccountsByID(ctx context.Context, primaryID model.FinanceProviderAccountsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceProviderAccountsCompositePrimaryKeyWhere([]model.FinanceProviderAccountsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeProviderAccountsQueries.selectCountFinanceProviderAccounts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceProviderAccountsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceProviderAccounts(ctx context.Context, selectFields ...FinanceProviderAccountsField) (financeProviderAccountsList model.FinanceProviderAccountsList, err error) {
	var (
		defaultFinanceProviderAccountsSelectFields = defaultFinanceProviderAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceProviderAccountsSelectFields = composeFinanceProviderAccountsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeProviderAccountsQueries.selectFinanceProviderAccounts, defaultFinanceProviderAccountsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeProviderAccountsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceProviderAccounts] failed get financeProviderAccounts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceProviderAccountsByID(ctx context.Context, primaryID model.FinanceProviderAccountsPrimaryID, selectFields ...FinanceProviderAccountsField) (financeProviderAccounts model.FinanceProviderAccounts, err error) {
	var (
		defaultFinanceProviderAccountsSelectFields = defaultFinanceProviderAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceProviderAccountsSelectFields = composeFinanceProviderAccountsSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceProviderAccountsCompositePrimaryKeyWhere([]model.FinanceProviderAccountsPrimaryID{primaryID})
	query := fmt.Sprintf(financeProviderAccountsQueries.selectFinanceProviderAccounts+" WHERE "+whereQry, defaultFinanceProviderAccountsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeProviderAccounts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeProviderAccounts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceProviderAccountsByID] failed get financeProviderAccounts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceProviderAccountsByID(ctx context.Context, primaryID model.FinanceProviderAccountsPrimaryID, financeProviderAccounts *model.FinanceProviderAccounts, financeProviderAccountsUpdateFields ...FinanceProviderAccountsUpdateField) (err error) {
	exists, err := repo.IsExistFinanceProviderAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceProviderAccounts] failed checking financeProviderAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeProviderAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeProviderAccounts == nil {
		if len(financeProviderAccountsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceProviderAccountsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeProviderAccounts = &model.FinanceProviderAccounts{}
	}
	var (
		defaultFinanceProviderAccountsUpdateFields = defaultFinanceProviderAccountsUpdateFields(*financeProviderAccounts)
		tempUpdateField                            FinanceProviderAccountsUpdateFieldList
		selectFields                               = NewFinanceProviderAccountsSelectFields()
	)
	if len(financeProviderAccountsUpdateFields) > 0 {
		for _, updateField := range financeProviderAccountsUpdateFields {
			if updateField.financeProviderAccountsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceProviderAccountsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceProviderAccountsCompositePrimaryKeyWhere([]model.FinanceProviderAccountsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceProviderAccountsCommand(defaultFinanceProviderAccountsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeProviderAccountsQueries.updateFinanceProviderAccounts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceProviderAccounts] error when try to update financeProviderAccounts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceProviderAccountsByFilter(ctx context.Context, filter model.Filter, financeProviderAccountsUpdateFields ...FinanceProviderAccountsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeProviderAccountsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceProviderAccountsUpdateFieldList
		selectFields = NewFinanceProviderAccountsSelectFields()
	)
	for _, updateField := range financeProviderAccountsUpdateFields {
		if updateField.financeProviderAccountsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceProviderAccountsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeProviderAccountsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceProviderAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_provider_accounts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceProviderAccountsByFilter] error when try to update financeProviderAccounts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceProviderAccountsByFilter] failed get rows affected")
	}
	return
}

var (
	financeProviderAccountsQueries = struct {
		selectFinanceProviderAccounts      string
		selectCountFinanceProviderAccounts string
		deleteFinanceProviderAccounts      string
		updateFinanceProviderAccounts      string
		insertFinanceProviderAccounts      string
	}{
		selectFinanceProviderAccounts:      "SELECT %s FROM \"finance_provider_accounts\"",
		selectCountFinanceProviderAccounts: "SELECT COUNT(\"id\") FROM \"finance_provider_accounts\"",
		deleteFinanceProviderAccounts:      "DELETE FROM \"finance_provider_accounts\"",
		updateFinanceProviderAccounts:      "UPDATE \"finance_provider_accounts\" SET %s ",
		insertFinanceProviderAccounts:      "INSERT INTO \"finance_provider_accounts\" %s VALUES %s",
	}
)

type FinanceProviderAccountsRepository interface {
	CreateFinanceProviderAccounts(ctx context.Context, financeProviderAccounts *model.FinanceProviderAccounts, fieldsInsert ...FinanceProviderAccountsField) error
	BulkCreateFinanceProviderAccounts(ctx context.Context, financeProviderAccountsList []*model.FinanceProviderAccounts, fieldsInsert ...FinanceProviderAccountsField) error
	ResolveFinanceProviderAccounts(ctx context.Context, selectFields ...FinanceProviderAccountsField) (model.FinanceProviderAccountsList, error)
	ResolveFinanceProviderAccountsByID(ctx context.Context, primaryID model.FinanceProviderAccountsPrimaryID, selectFields ...FinanceProviderAccountsField) (model.FinanceProviderAccounts, error)
	UpdateFinanceProviderAccountsByID(ctx context.Context, id model.FinanceProviderAccountsPrimaryID, financeProviderAccounts *model.FinanceProviderAccounts, financeProviderAccountsUpdateFields ...FinanceProviderAccountsUpdateField) error
	UpdateFinanceProviderAccountsByFilter(ctx context.Context, filter model.Filter, financeProviderAccountsUpdateFields ...FinanceProviderAccountsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceProviderAccounts(ctx context.Context, financeProviderAccountsListMap map[model.FinanceProviderAccountsPrimaryID]*model.FinanceProviderAccounts, FinanceProviderAccountssMapUpdateFieldsRequest map[model.FinanceProviderAccountsPrimaryID]FinanceProviderAccountsUpdateFieldList) (err error)
	DeleteFinanceProviderAccountsByID(ctx context.Context, id model.FinanceProviderAccountsPrimaryID) error
	BulkDeleteFinanceProviderAccountsByIDs(ctx context.Context, ids []model.FinanceProviderAccountsPrimaryID) error
	ResolveFinanceProviderAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceProviderAccountsFilterResult, err error)
	IsExistFinanceProviderAccountsByIDs(ctx context.Context, ids []model.FinanceProviderAccountsPrimaryID) (exists bool, notFoundIds []model.FinanceProviderAccountsPrimaryID, err error)
	IsExistFinanceProviderAccountsByID(ctx context.Context, id model.FinanceProviderAccountsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
