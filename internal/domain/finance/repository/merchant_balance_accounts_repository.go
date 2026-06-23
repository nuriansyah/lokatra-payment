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

func composeInsertFieldsAndParamsMerchantBalanceAccounts(merchantBalanceAccountsList []model.MerchantBalanceAccounts, fieldsInsert ...MerchantBalanceAccountsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewMerchantBalanceAccountsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, merchantBalanceAccounts := range merchantBalanceAccountsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, merchantBalanceAccounts.Id)
			case selectField.MerchantPartyId():
				args = append(args, merchantBalanceAccounts.MerchantPartyId)
			case selectField.BalanceType():
				args = append(args, merchantBalanceAccounts.BalanceType)
			case selectField.CurrencyCode():
				args = append(args, merchantBalanceAccounts.CurrencyCode)
			case selectField.LinkedLedgerAccountId():
				args = append(args, merchantBalanceAccounts.LinkedLedgerAccountId)
			case selectField.AccountStatus():
				args = append(args, merchantBalanceAccounts.AccountStatus)
			case selectField.Metadata():
				args = append(args, merchantBalanceAccounts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, merchantBalanceAccounts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, merchantBalanceAccounts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, merchantBalanceAccounts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, merchantBalanceAccounts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, merchantBalanceAccounts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, merchantBalanceAccounts.MetaDeletedBy)

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

func composeMerchantBalanceAccountsCompositePrimaryKeyWhere(primaryIDs []model.MerchantBalanceAccountsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"merchant_balance_accounts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultMerchantBalanceAccountsSelectFields() string {
	fields := NewMerchantBalanceAccountsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeMerchantBalanceAccountsSelectFields(selectFields ...MerchantBalanceAccountsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type MerchantBalanceAccountsField string
type MerchantBalanceAccountsFieldList []MerchantBalanceAccountsField

type MerchantBalanceAccountsSelectFields struct {
}

func (ss MerchantBalanceAccountsSelectFields) Id() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("id")
}

func (ss MerchantBalanceAccountsSelectFields) MerchantPartyId() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("merchant_party_id")
}

func (ss MerchantBalanceAccountsSelectFields) BalanceType() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("balance_type")
}

func (ss MerchantBalanceAccountsSelectFields) CurrencyCode() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("currency_code")
}

func (ss MerchantBalanceAccountsSelectFields) LinkedLedgerAccountId() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("linked_ledger_account_id")
}

func (ss MerchantBalanceAccountsSelectFields) AccountStatus() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("account_status")
}

func (ss MerchantBalanceAccountsSelectFields) Metadata() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("metadata")
}

func (ss MerchantBalanceAccountsSelectFields) MetaCreatedAt() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("meta_created_at")
}

func (ss MerchantBalanceAccountsSelectFields) MetaCreatedBy() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("meta_created_by")
}

func (ss MerchantBalanceAccountsSelectFields) MetaUpdatedAt() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("meta_updated_at")
}

func (ss MerchantBalanceAccountsSelectFields) MetaUpdatedBy() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("meta_updated_by")
}

func (ss MerchantBalanceAccountsSelectFields) MetaDeletedAt() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("meta_deleted_at")
}

func (ss MerchantBalanceAccountsSelectFields) MetaDeletedBy() MerchantBalanceAccountsField {
	return MerchantBalanceAccountsField("meta_deleted_by")
}

func (ss MerchantBalanceAccountsSelectFields) All() MerchantBalanceAccountsFieldList {
	return []MerchantBalanceAccountsField{
		ss.Id(),
		ss.MerchantPartyId(),
		ss.BalanceType(),
		ss.CurrencyCode(),
		ss.LinkedLedgerAccountId(),
		ss.AccountStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewMerchantBalanceAccountsSelectFields() MerchantBalanceAccountsSelectFields {
	return MerchantBalanceAccountsSelectFields{}
}

type MerchantBalanceAccountsUpdateFieldOption struct {
	useIncrement bool
}
type MerchantBalanceAccountsUpdateField struct {
	merchantBalanceAccountsField MerchantBalanceAccountsField
	opt                          MerchantBalanceAccountsUpdateFieldOption
	value                        interface{}
}
type MerchantBalanceAccountsUpdateFieldList []MerchantBalanceAccountsUpdateField

func defaultMerchantBalanceAccountsUpdateFieldOption() MerchantBalanceAccountsUpdateFieldOption {
	return MerchantBalanceAccountsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementMerchantBalanceAccountsOption(useIncrement bool) func(*MerchantBalanceAccountsUpdateFieldOption) {
	return func(pcufo *MerchantBalanceAccountsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewMerchantBalanceAccountsUpdateField(field MerchantBalanceAccountsField, val interface{}, opts ...func(*MerchantBalanceAccountsUpdateFieldOption)) MerchantBalanceAccountsUpdateField {
	defaultOpt := defaultMerchantBalanceAccountsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return MerchantBalanceAccountsUpdateField{
		merchantBalanceAccountsField: field,
		value:                        val,
		opt:                          defaultOpt,
	}
}
func defaultMerchantBalanceAccountsUpdateFields(merchantBalanceAccounts model.MerchantBalanceAccounts) (merchantBalanceAccountsUpdateFieldList MerchantBalanceAccountsUpdateFieldList) {
	selectFields := NewMerchantBalanceAccountsSelectFields()
	merchantBalanceAccountsUpdateFieldList = append(merchantBalanceAccountsUpdateFieldList,
		NewMerchantBalanceAccountsUpdateField(selectFields.Id(), merchantBalanceAccounts.Id),
		NewMerchantBalanceAccountsUpdateField(selectFields.MerchantPartyId(), merchantBalanceAccounts.MerchantPartyId),
		NewMerchantBalanceAccountsUpdateField(selectFields.BalanceType(), merchantBalanceAccounts.BalanceType),
		NewMerchantBalanceAccountsUpdateField(selectFields.CurrencyCode(), merchantBalanceAccounts.CurrencyCode),
		NewMerchantBalanceAccountsUpdateField(selectFields.LinkedLedgerAccountId(), merchantBalanceAccounts.LinkedLedgerAccountId),
		NewMerchantBalanceAccountsUpdateField(selectFields.AccountStatus(), merchantBalanceAccounts.AccountStatus),
		NewMerchantBalanceAccountsUpdateField(selectFields.Metadata(), merchantBalanceAccounts.Metadata),
		NewMerchantBalanceAccountsUpdateField(selectFields.MetaCreatedAt(), merchantBalanceAccounts.MetaCreatedAt),
		NewMerchantBalanceAccountsUpdateField(selectFields.MetaCreatedBy(), merchantBalanceAccounts.MetaCreatedBy),
		NewMerchantBalanceAccountsUpdateField(selectFields.MetaUpdatedAt(), merchantBalanceAccounts.MetaUpdatedAt),
		NewMerchantBalanceAccountsUpdateField(selectFields.MetaUpdatedBy(), merchantBalanceAccounts.MetaUpdatedBy),
		NewMerchantBalanceAccountsUpdateField(selectFields.MetaDeletedAt(), merchantBalanceAccounts.MetaDeletedAt),
		NewMerchantBalanceAccountsUpdateField(selectFields.MetaDeletedBy(), merchantBalanceAccounts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsMerchantBalanceAccountsCommand(merchantBalanceAccountsUpdateFieldList MerchantBalanceAccountsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range merchantBalanceAccountsUpdateFieldList {
		field := string(updateField.merchantBalanceAccountsField)
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

func (repo *RepositoryImpl) BulkCreateMerchantBalanceAccounts(ctx context.Context, merchantBalanceAccountsList []*model.MerchantBalanceAccounts, fieldsInsert ...MerchantBalanceAccountsField) (err error) {
	var (
		fieldsStr                        string
		valueListStr                     []string
		argsList                         []interface{}
		primaryIds                       []model.MerchantBalanceAccountsPrimaryID
		merchantBalanceAccountsValueList []model.MerchantBalanceAccounts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, merchantBalanceAccounts := range merchantBalanceAccountsList {

		primaryIds = append(primaryIds, merchantBalanceAccounts.ToMerchantBalanceAccountsPrimaryID())

		merchantBalanceAccountsValueList = append(merchantBalanceAccountsValueList, *merchantBalanceAccounts)
	}

	_, notFoundIds, err := repo.IsExistMerchantBalanceAccountsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceAccounts] failed checking merchantBalanceAccounts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.MerchantBalanceAccountsPrimaryID{}
		mapNotFoundIds := map[model.MerchantBalanceAccountsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "merchantBalanceAccounts", fmt.Sprintf("merchantBalanceAccounts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsMerchantBalanceAccounts(merchantBalanceAccountsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(merchantBalanceAccountsQueries.insertMerchantBalanceAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceAccounts] failed exec create merchantBalanceAccounts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteMerchantBalanceAccountsByIDs(ctx context.Context, primaryIDs []model.MerchantBalanceAccountsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistMerchantBalanceAccountsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceAccountsByIDs] failed checking merchantBalanceAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceAccounts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"merchant_balance_accounts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(merchantBalanceAccountsQueries.deleteMerchantBalanceAccounts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceAccountsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceAccountsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistMerchantBalanceAccountsByIDs(ctx context.Context, ids []model.MerchantBalanceAccountsPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceAccountsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"merchant_balance_accounts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(merchantBalanceAccountsQueries.selectMerchantBalanceAccounts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceAccountsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.MerchantBalanceAccountsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceAccountsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.MerchantBalanceAccountsPrimaryID]bool{}
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

// BulkUpdateMerchantBalanceAccounts is used to bulk update merchantBalanceAccounts, by default it will update all field
// if want to update specific field, then fill merchantBalanceAccountssMapUpdateFieldsRequest else please fill merchantBalanceAccountssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateMerchantBalanceAccounts(ctx context.Context, merchantBalanceAccountssMap map[model.MerchantBalanceAccountsPrimaryID]*model.MerchantBalanceAccounts, merchantBalanceAccountssMapUpdateFieldsRequest map[model.MerchantBalanceAccountsPrimaryID]MerchantBalanceAccountsUpdateFieldList) (err error) {
	if len(merchantBalanceAccountssMap) == 0 && len(merchantBalanceAccountssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		merchantBalanceAccountssMapUpdateField map[model.MerchantBalanceAccountsPrimaryID]MerchantBalanceAccountsUpdateFieldList = map[model.MerchantBalanceAccountsPrimaryID]MerchantBalanceAccountsUpdateFieldList{}
		asTableValues                          string                                                                            = "myvalues"
	)

	if len(merchantBalanceAccountssMap) > 0 {
		for id, merchantBalanceAccounts := range merchantBalanceAccountssMap {
			if merchantBalanceAccounts == nil {
				log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceAccounts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			merchantBalanceAccountssMapUpdateField[id] = defaultMerchantBalanceAccountsUpdateFields(*merchantBalanceAccounts)
		}
	} else {
		merchantBalanceAccountssMapUpdateField = merchantBalanceAccountssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateMerchantBalanceAccountsQuery(merchantBalanceAccountssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistMerchantBalanceAccountsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceAccounts] failed checking merchantBalanceAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceAccounts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeMerchantBalanceAccountsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"merchant_balance_accounts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceAccounts] failed exec query")
	}
	return
}

type MerchantBalanceAccountsFieldParameter struct {
	param string
	args  []interface{}
}

func NewMerchantBalanceAccountsFieldParameter(param string, args ...interface{}) MerchantBalanceAccountsFieldParameter {
	return MerchantBalanceAccountsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateMerchantBalanceAccountsQuery(mapMerchantBalanceAccountss map[model.MerchantBalanceAccountsPrimaryID]MerchantBalanceAccountsUpdateFieldList, asTableValues string) (primaryIDs []model.MerchantBalanceAccountsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.MerchantBalanceAccountsPrimaryID]map[string]interface{}{}
	merchantBalanceAccountsSelectFields := NewMerchantBalanceAccountsSelectFields()
	for id, updateFields := range mapMerchantBalanceAccountss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.merchantBalanceAccountsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapMerchantBalanceAccountss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetMerchantBalanceAccountsFieldType(updateField.merchantBalanceAccountsField)))
			args = append(args, fields[string(updateField.merchantBalanceAccountsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.merchantBalanceAccountsField))
		if updateField.merchantBalanceAccountsField == merchantBalanceAccountsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.merchantBalanceAccountsField, asTableValues, updateField.merchantBalanceAccountsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.merchantBalanceAccountsField,
				"\"merchant_balance_accounts\"", updateField.merchantBalanceAccountsField,
				asTableValues, updateField.merchantBalanceAccountsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeMerchantBalanceAccountsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.MerchantBalanceAccountsPrimaryID, asTableValue string) (whereQry string) {
	merchantBalanceAccountsSelectFields := NewMerchantBalanceAccountsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"merchant_balance_accounts\".\"id\" = %s.\"id\"::"+GetMerchantBalanceAccountsFieldType(merchantBalanceAccountsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetMerchantBalanceAccountsFieldType(merchantBalanceAccountsField MerchantBalanceAccountsField) string {
	selectMerchantBalanceAccountsFields := NewMerchantBalanceAccountsSelectFields()
	switch merchantBalanceAccountsField {

	case selectMerchantBalanceAccountsFields.Id():
		return "uuid"

	case selectMerchantBalanceAccountsFields.MerchantPartyId():
		return "uuid"

	case selectMerchantBalanceAccountsFields.BalanceType():
		return "merchant_balance_accounts_balance_type_enum"

	case selectMerchantBalanceAccountsFields.CurrencyCode():
		return "text"

	case selectMerchantBalanceAccountsFields.LinkedLedgerAccountId():
		return "uuid"

	case selectMerchantBalanceAccountsFields.AccountStatus():
		return "merchant_balance_accounts_account_status_enum"

	case selectMerchantBalanceAccountsFields.Metadata():
		return "jsonb"

	case selectMerchantBalanceAccountsFields.MetaCreatedAt():
		return "timestamptz"

	case selectMerchantBalanceAccountsFields.MetaCreatedBy():
		return "uuid"

	case selectMerchantBalanceAccountsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectMerchantBalanceAccountsFields.MetaUpdatedBy():
		return "uuid"

	case selectMerchantBalanceAccountsFields.MetaDeletedAt():
		return "timestamptz"

	case selectMerchantBalanceAccountsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateMerchantBalanceAccounts(ctx context.Context, merchantBalanceAccounts *model.MerchantBalanceAccounts, fieldsInsert ...MerchantBalanceAccountsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.MerchantBalanceAccountsPrimaryID{
		Id: merchantBalanceAccounts.Id,
	}
	exists, err := repo.IsExistMerchantBalanceAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceAccounts] failed checking merchantBalanceAccounts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "merchantBalanceAccounts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsMerchantBalanceAccounts([]model.MerchantBalanceAccounts{*merchantBalanceAccounts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(merchantBalanceAccountsQueries.insertMerchantBalanceAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceAccounts] failed exec create merchantBalanceAccounts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteMerchantBalanceAccountsByID(ctx context.Context, primaryID model.MerchantBalanceAccountsPrimaryID) (err error) {
	exists, err := repo.IsExistMerchantBalanceAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceAccountsByID] failed checking merchantBalanceAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeMerchantBalanceAccountsCompositePrimaryKeyWhere([]model.MerchantBalanceAccountsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(merchantBalanceAccountsQueries.deleteMerchantBalanceAccounts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceAccountsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceAccountsFilterResult, err error) {
	query, args, err := composeMerchantBalanceAccountsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceAccountsByFilter] failed compose merchantBalanceAccounts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceAccountsByFilter] failed get merchantBalanceAccounts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeMerchantBalanceAccountsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.MerchantBalanceAccountsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeMerchantBalanceAccountsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeMerchantBalanceAccountsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeMerchantBalanceAccountsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 13 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewMerchantBalanceAccountsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 13+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["balance_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"balance_type\"")
			selectedColumns["balance_type"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["linked_ledger_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"linked_ledger_account_id\"")
			selectedColumns["linked_ledger_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["account_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_status\"")
			selectedColumns["account_status"] = struct{}{}
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

type merchantBalanceAccountsFilterPlaceholder struct {
	index int
}

func (p *merchantBalanceAccountsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeMerchantBalanceAccountsFilterPredicate(filterField model.FilterField, placeholders *merchantBalanceAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewMerchantBalanceAccountsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeMerchantBalanceAccountsFilterSQLExpr(spec)
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

func composeMerchantBalanceAccountsFilterGroup(group model.FilterGroup, placeholders *merchantBalanceAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeMerchantBalanceAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeMerchantBalanceAccountsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeMerchantBalanceAccountsFilterWhereQueries(filter model.Filter, placeholders *merchantBalanceAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeMerchantBalanceAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeMerchantBalanceAccountsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeMerchantBalanceAccountsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateMerchantBalanceAccountsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeMerchantBalanceAccountsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeMerchantBalanceAccountsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := merchantBalanceAccountsFilterPlaceholder{index: 1}
	whereQueries, err := composeMerchantBalanceAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewMerchantBalanceAccountsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeMerchantBalanceAccountsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeMerchantBalanceAccountsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"merchant_balance_accounts\" base%s", strings.Join(selectColumns, ","), composeMerchantBalanceAccountsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistMerchantBalanceAccountsByID(ctx context.Context, primaryID model.MerchantBalanceAccountsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeMerchantBalanceAccountsCompositePrimaryKeyWhere([]model.MerchantBalanceAccountsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", merchantBalanceAccountsQueries.selectCountMerchantBalanceAccounts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceAccountsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceAccounts(ctx context.Context, selectFields ...MerchantBalanceAccountsField) (merchantBalanceAccountsList model.MerchantBalanceAccountsList, err error) {
	var (
		defaultMerchantBalanceAccountsSelectFields = defaultMerchantBalanceAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceAccountsSelectFields = composeMerchantBalanceAccountsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(merchantBalanceAccountsQueries.selectMerchantBalanceAccounts, defaultMerchantBalanceAccountsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &merchantBalanceAccountsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceAccounts] failed get merchantBalanceAccounts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceAccountsByID(ctx context.Context, primaryID model.MerchantBalanceAccountsPrimaryID, selectFields ...MerchantBalanceAccountsField) (merchantBalanceAccounts model.MerchantBalanceAccounts, err error) {
	var (
		defaultMerchantBalanceAccountsSelectFields = defaultMerchantBalanceAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceAccountsSelectFields = composeMerchantBalanceAccountsSelectFields(selectFields...)
	}
	whereQry, params := composeMerchantBalanceAccountsCompositePrimaryKeyWhere([]model.MerchantBalanceAccountsPrimaryID{primaryID})
	query := fmt.Sprintf(merchantBalanceAccountsQueries.selectMerchantBalanceAccounts+" WHERE "+whereQry, defaultMerchantBalanceAccountsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &merchantBalanceAccounts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("merchantBalanceAccounts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveMerchantBalanceAccountsByID] failed get merchantBalanceAccounts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateMerchantBalanceAccountsByID(ctx context.Context, primaryID model.MerchantBalanceAccountsPrimaryID, merchantBalanceAccounts *model.MerchantBalanceAccounts, merchantBalanceAccountsUpdateFields ...MerchantBalanceAccountsUpdateField) (err error) {
	exists, err := repo.IsExistMerchantBalanceAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceAccounts] failed checking merchantBalanceAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if merchantBalanceAccounts == nil {
		if len(merchantBalanceAccountsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateMerchantBalanceAccountsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		merchantBalanceAccounts = &model.MerchantBalanceAccounts{}
	}
	var (
		defaultMerchantBalanceAccountsUpdateFields = defaultMerchantBalanceAccountsUpdateFields(*merchantBalanceAccounts)
		tempUpdateField                            MerchantBalanceAccountsUpdateFieldList
		selectFields                               = NewMerchantBalanceAccountsSelectFields()
	)
	if len(merchantBalanceAccountsUpdateFields) > 0 {
		for _, updateField := range merchantBalanceAccountsUpdateFields {
			if updateField.merchantBalanceAccountsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultMerchantBalanceAccountsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeMerchantBalanceAccountsCompositePrimaryKeyWhere([]model.MerchantBalanceAccountsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsMerchantBalanceAccountsCommand(defaultMerchantBalanceAccountsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(merchantBalanceAccountsQueries.updateMerchantBalanceAccounts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceAccounts] error when try to update merchantBalanceAccounts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateMerchantBalanceAccountsByFilter(ctx context.Context, filter model.Filter, merchantBalanceAccountsUpdateFields ...MerchantBalanceAccountsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(merchantBalanceAccountsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields MerchantBalanceAccountsUpdateFieldList
		selectFields = NewMerchantBalanceAccountsSelectFields()
	)
	for _, updateField := range merchantBalanceAccountsUpdateFields {
		if updateField.merchantBalanceAccountsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsMerchantBalanceAccountsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := merchantBalanceAccountsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeMerchantBalanceAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"merchant_balance_accounts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceAccountsByFilter] error when try to update merchantBalanceAccounts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceAccountsByFilter] failed get rows affected")
	}
	return
}

var (
	merchantBalanceAccountsQueries = struct {
		selectMerchantBalanceAccounts      string
		selectCountMerchantBalanceAccounts string
		deleteMerchantBalanceAccounts      string
		updateMerchantBalanceAccounts      string
		insertMerchantBalanceAccounts      string
	}{
		selectMerchantBalanceAccounts:      "SELECT %s FROM \"merchant_balance_accounts\"",
		selectCountMerchantBalanceAccounts: "SELECT COUNT(\"id\") FROM \"merchant_balance_accounts\"",
		deleteMerchantBalanceAccounts:      "DELETE FROM \"merchant_balance_accounts\"",
		updateMerchantBalanceAccounts:      "UPDATE \"merchant_balance_accounts\" SET %s ",
		insertMerchantBalanceAccounts:      "INSERT INTO \"merchant_balance_accounts\" %s VALUES %s",
	}
)

type MerchantBalanceAccountsRepository interface {
	CreateMerchantBalanceAccounts(ctx context.Context, merchantBalanceAccounts *model.MerchantBalanceAccounts, fieldsInsert ...MerchantBalanceAccountsField) error
	BulkCreateMerchantBalanceAccounts(ctx context.Context, merchantBalanceAccountsList []*model.MerchantBalanceAccounts, fieldsInsert ...MerchantBalanceAccountsField) error
	ResolveMerchantBalanceAccounts(ctx context.Context, selectFields ...MerchantBalanceAccountsField) (model.MerchantBalanceAccountsList, error)
	ResolveMerchantBalanceAccountsByID(ctx context.Context, primaryID model.MerchantBalanceAccountsPrimaryID, selectFields ...MerchantBalanceAccountsField) (model.MerchantBalanceAccounts, error)
	UpdateMerchantBalanceAccountsByID(ctx context.Context, id model.MerchantBalanceAccountsPrimaryID, merchantBalanceAccounts *model.MerchantBalanceAccounts, merchantBalanceAccountsUpdateFields ...MerchantBalanceAccountsUpdateField) error
	UpdateMerchantBalanceAccountsByFilter(ctx context.Context, filter model.Filter, merchantBalanceAccountsUpdateFields ...MerchantBalanceAccountsUpdateField) (rowsAffected int64, err error)
	BulkUpdateMerchantBalanceAccounts(ctx context.Context, merchantBalanceAccountsListMap map[model.MerchantBalanceAccountsPrimaryID]*model.MerchantBalanceAccounts, MerchantBalanceAccountssMapUpdateFieldsRequest map[model.MerchantBalanceAccountsPrimaryID]MerchantBalanceAccountsUpdateFieldList) (err error)
	DeleteMerchantBalanceAccountsByID(ctx context.Context, id model.MerchantBalanceAccountsPrimaryID) error
	BulkDeleteMerchantBalanceAccountsByIDs(ctx context.Context, ids []model.MerchantBalanceAccountsPrimaryID) error
	ResolveMerchantBalanceAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceAccountsFilterResult, err error)
	IsExistMerchantBalanceAccountsByIDs(ctx context.Context, ids []model.MerchantBalanceAccountsPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceAccountsPrimaryID, err error)
	IsExistMerchantBalanceAccountsByID(ctx context.Context, id model.MerchantBalanceAccountsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
