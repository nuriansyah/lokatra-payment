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

func composeInsertFieldsAndParamsMerchantBalanceTransactions(merchantBalanceTransactionsList []model.MerchantBalanceTransactions, fieldsInsert ...MerchantBalanceTransactionsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewMerchantBalanceTransactionsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, merchantBalanceTransactions := range merchantBalanceTransactionsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, merchantBalanceTransactions.Id)
			case selectField.BalanceAccountId():
				args = append(args, merchantBalanceTransactions.BalanceAccountId)
			case selectField.MerchantPartyId():
				args = append(args, merchantBalanceTransactions.MerchantPartyId)
			case selectField.BalanceType():
				args = append(args, merchantBalanceTransactions.BalanceType)
			case selectField.CurrencyCode():
				args = append(args, merchantBalanceTransactions.CurrencyCode)
			case selectField.SourceType():
				args = append(args, merchantBalanceTransactions.SourceType)
			case selectField.SourceId():
				args = append(args, merchantBalanceTransactions.SourceId)
			case selectField.Direction():
				args = append(args, merchantBalanceTransactions.Direction)
			case selectField.Amount():
				args = append(args, merchantBalanceTransactions.Amount)
			case selectField.BalanceBefore():
				args = append(args, merchantBalanceTransactions.BalanceBefore)
			case selectField.BalanceAfter():
				args = append(args, merchantBalanceTransactions.BalanceAfter)
			case selectField.ReasonCode():
				args = append(args, merchantBalanceTransactions.ReasonCode)
			case selectField.Metadata():
				args = append(args, merchantBalanceTransactions.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, merchantBalanceTransactions.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, merchantBalanceTransactions.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, merchantBalanceTransactions.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, merchantBalanceTransactions.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, merchantBalanceTransactions.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, merchantBalanceTransactions.MetaDeletedBy)

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

func composeMerchantBalanceTransactionsCompositePrimaryKeyWhere(primaryIDs []model.MerchantBalanceTransactionsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"merchant_balance_transactions\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultMerchantBalanceTransactionsSelectFields() string {
	fields := NewMerchantBalanceTransactionsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeMerchantBalanceTransactionsSelectFields(selectFields ...MerchantBalanceTransactionsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type MerchantBalanceTransactionsField string
type MerchantBalanceTransactionsFieldList []MerchantBalanceTransactionsField

type MerchantBalanceTransactionsSelectFields struct {
}

func (ss MerchantBalanceTransactionsSelectFields) Id() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("id")
}

func (ss MerchantBalanceTransactionsSelectFields) BalanceAccountId() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("balance_account_id")
}

func (ss MerchantBalanceTransactionsSelectFields) MerchantPartyId() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("merchant_party_id")
}

func (ss MerchantBalanceTransactionsSelectFields) BalanceType() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("balance_type")
}

func (ss MerchantBalanceTransactionsSelectFields) CurrencyCode() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("currency_code")
}

func (ss MerchantBalanceTransactionsSelectFields) SourceType() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("source_type")
}

func (ss MerchantBalanceTransactionsSelectFields) SourceId() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("source_id")
}

func (ss MerchantBalanceTransactionsSelectFields) Direction() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("direction")
}

func (ss MerchantBalanceTransactionsSelectFields) Amount() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("amount")
}

func (ss MerchantBalanceTransactionsSelectFields) BalanceBefore() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("balance_before")
}

func (ss MerchantBalanceTransactionsSelectFields) BalanceAfter() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("balance_after")
}

func (ss MerchantBalanceTransactionsSelectFields) ReasonCode() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("reason_code")
}

func (ss MerchantBalanceTransactionsSelectFields) Metadata() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("metadata")
}

func (ss MerchantBalanceTransactionsSelectFields) MetaCreatedAt() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("meta_created_at")
}

func (ss MerchantBalanceTransactionsSelectFields) MetaCreatedBy() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("meta_created_by")
}

func (ss MerchantBalanceTransactionsSelectFields) MetaUpdatedAt() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("meta_updated_at")
}

func (ss MerchantBalanceTransactionsSelectFields) MetaUpdatedBy() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("meta_updated_by")
}

func (ss MerchantBalanceTransactionsSelectFields) MetaDeletedAt() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("meta_deleted_at")
}

func (ss MerchantBalanceTransactionsSelectFields) MetaDeletedBy() MerchantBalanceTransactionsField {
	return MerchantBalanceTransactionsField("meta_deleted_by")
}

func (ss MerchantBalanceTransactionsSelectFields) All() MerchantBalanceTransactionsFieldList {
	return []MerchantBalanceTransactionsField{
		ss.Id(),
		ss.BalanceAccountId(),
		ss.MerchantPartyId(),
		ss.BalanceType(),
		ss.CurrencyCode(),
		ss.SourceType(),
		ss.SourceId(),
		ss.Direction(),
		ss.Amount(),
		ss.BalanceBefore(),
		ss.BalanceAfter(),
		ss.ReasonCode(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewMerchantBalanceTransactionsSelectFields() MerchantBalanceTransactionsSelectFields {
	return MerchantBalanceTransactionsSelectFields{}
}

type MerchantBalanceTransactionsUpdateFieldOption struct {
	useIncrement bool
}
type MerchantBalanceTransactionsUpdateField struct {
	merchantBalanceTransactionsField MerchantBalanceTransactionsField
	opt                              MerchantBalanceTransactionsUpdateFieldOption
	value                            interface{}
}
type MerchantBalanceTransactionsUpdateFieldList []MerchantBalanceTransactionsUpdateField

func defaultMerchantBalanceTransactionsUpdateFieldOption() MerchantBalanceTransactionsUpdateFieldOption {
	return MerchantBalanceTransactionsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementMerchantBalanceTransactionsOption(useIncrement bool) func(*MerchantBalanceTransactionsUpdateFieldOption) {
	return func(pcufo *MerchantBalanceTransactionsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewMerchantBalanceTransactionsUpdateField(field MerchantBalanceTransactionsField, val interface{}, opts ...func(*MerchantBalanceTransactionsUpdateFieldOption)) MerchantBalanceTransactionsUpdateField {
	defaultOpt := defaultMerchantBalanceTransactionsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return MerchantBalanceTransactionsUpdateField{
		merchantBalanceTransactionsField: field,
		value:                            val,
		opt:                              defaultOpt,
	}
}
func defaultMerchantBalanceTransactionsUpdateFields(merchantBalanceTransactions model.MerchantBalanceTransactions) (merchantBalanceTransactionsUpdateFieldList MerchantBalanceTransactionsUpdateFieldList) {
	selectFields := NewMerchantBalanceTransactionsSelectFields()
	merchantBalanceTransactionsUpdateFieldList = append(merchantBalanceTransactionsUpdateFieldList,
		NewMerchantBalanceTransactionsUpdateField(selectFields.Id(), merchantBalanceTransactions.Id),
		NewMerchantBalanceTransactionsUpdateField(selectFields.BalanceAccountId(), merchantBalanceTransactions.BalanceAccountId),
		NewMerchantBalanceTransactionsUpdateField(selectFields.MerchantPartyId(), merchantBalanceTransactions.MerchantPartyId),
		NewMerchantBalanceTransactionsUpdateField(selectFields.BalanceType(), merchantBalanceTransactions.BalanceType),
		NewMerchantBalanceTransactionsUpdateField(selectFields.CurrencyCode(), merchantBalanceTransactions.CurrencyCode),
		NewMerchantBalanceTransactionsUpdateField(selectFields.SourceType(), merchantBalanceTransactions.SourceType),
		NewMerchantBalanceTransactionsUpdateField(selectFields.SourceId(), merchantBalanceTransactions.SourceId),
		NewMerchantBalanceTransactionsUpdateField(selectFields.Direction(), merchantBalanceTransactions.Direction),
		NewMerchantBalanceTransactionsUpdateField(selectFields.Amount(), merchantBalanceTransactions.Amount),
		NewMerchantBalanceTransactionsUpdateField(selectFields.BalanceBefore(), merchantBalanceTransactions.BalanceBefore),
		NewMerchantBalanceTransactionsUpdateField(selectFields.BalanceAfter(), merchantBalanceTransactions.BalanceAfter),
		NewMerchantBalanceTransactionsUpdateField(selectFields.ReasonCode(), merchantBalanceTransactions.ReasonCode),
		NewMerchantBalanceTransactionsUpdateField(selectFields.Metadata(), merchantBalanceTransactions.Metadata),
		NewMerchantBalanceTransactionsUpdateField(selectFields.MetaCreatedAt(), merchantBalanceTransactions.MetaCreatedAt),
		NewMerchantBalanceTransactionsUpdateField(selectFields.MetaCreatedBy(), merchantBalanceTransactions.MetaCreatedBy),
		NewMerchantBalanceTransactionsUpdateField(selectFields.MetaUpdatedAt(), merchantBalanceTransactions.MetaUpdatedAt),
		NewMerchantBalanceTransactionsUpdateField(selectFields.MetaUpdatedBy(), merchantBalanceTransactions.MetaUpdatedBy),
		NewMerchantBalanceTransactionsUpdateField(selectFields.MetaDeletedAt(), merchantBalanceTransactions.MetaDeletedAt),
		NewMerchantBalanceTransactionsUpdateField(selectFields.MetaDeletedBy(), merchantBalanceTransactions.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsMerchantBalanceTransactionsCommand(merchantBalanceTransactionsUpdateFieldList MerchantBalanceTransactionsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range merchantBalanceTransactionsUpdateFieldList {
		field := string(updateField.merchantBalanceTransactionsField)
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

func (repo *RepositoryImpl) BulkCreateMerchantBalanceTransactions(ctx context.Context, merchantBalanceTransactionsList []*model.MerchantBalanceTransactions, fieldsInsert ...MerchantBalanceTransactionsField) (err error) {
	var (
		fieldsStr                            string
		valueListStr                         []string
		argsList                             []interface{}
		primaryIds                           []model.MerchantBalanceTransactionsPrimaryID
		merchantBalanceTransactionsValueList []model.MerchantBalanceTransactions
	)

	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceTransactionsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, merchantBalanceTransactions := range merchantBalanceTransactionsList {

		primaryIds = append(primaryIds, merchantBalanceTransactions.ToMerchantBalanceTransactionsPrimaryID())

		merchantBalanceTransactionsValueList = append(merchantBalanceTransactionsValueList, *merchantBalanceTransactions)
	}

	_, notFoundIds, err := repo.IsExistMerchantBalanceTransactionsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceTransactions] failed checking merchantBalanceTransactions whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.MerchantBalanceTransactionsPrimaryID{}
		mapNotFoundIds := map[model.MerchantBalanceTransactionsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "merchantBalanceTransactions", fmt.Sprintf("merchantBalanceTransactions with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsMerchantBalanceTransactions(merchantBalanceTransactionsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(merchantBalanceTransactionsQueries.insertMerchantBalanceTransactions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceTransactions] failed exec create merchantBalanceTransactions query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteMerchantBalanceTransactionsByIDs(ctx context.Context, primaryIDs []model.MerchantBalanceTransactionsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistMerchantBalanceTransactionsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceTransactionsByIDs] failed checking merchantBalanceTransactions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceTransactions with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"merchant_balance_transactions\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(merchantBalanceTransactionsQueries.deleteMerchantBalanceTransactions + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceTransactionsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceTransactionsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistMerchantBalanceTransactionsByIDs(ctx context.Context, ids []model.MerchantBalanceTransactionsPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceTransactionsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"merchant_balance_transactions\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(merchantBalanceTransactionsQueries.selectMerchantBalanceTransactions, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceTransactionsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.MerchantBalanceTransactionsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceTransactionsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.MerchantBalanceTransactionsPrimaryID]bool{}
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

// BulkUpdateMerchantBalanceTransactions is used to bulk update merchantBalanceTransactions, by default it will update all field
// if want to update specific field, then fill merchantBalanceTransactionssMapUpdateFieldsRequest else please fill merchantBalanceTransactionssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateMerchantBalanceTransactions(ctx context.Context, merchantBalanceTransactionssMap map[model.MerchantBalanceTransactionsPrimaryID]*model.MerchantBalanceTransactions, merchantBalanceTransactionssMapUpdateFieldsRequest map[model.MerchantBalanceTransactionsPrimaryID]MerchantBalanceTransactionsUpdateFieldList) (err error) {
	if len(merchantBalanceTransactionssMap) == 0 && len(merchantBalanceTransactionssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		merchantBalanceTransactionssMapUpdateField map[model.MerchantBalanceTransactionsPrimaryID]MerchantBalanceTransactionsUpdateFieldList = map[model.MerchantBalanceTransactionsPrimaryID]MerchantBalanceTransactionsUpdateFieldList{}
		asTableValues                              string                                                                                    = "myvalues"
	)

	if len(merchantBalanceTransactionssMap) > 0 {
		for id, merchantBalanceTransactions := range merchantBalanceTransactionssMap {
			if merchantBalanceTransactions == nil {
				log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceTransactions] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			merchantBalanceTransactionssMapUpdateField[id] = defaultMerchantBalanceTransactionsUpdateFields(*merchantBalanceTransactions)
		}
	} else {
		merchantBalanceTransactionssMapUpdateField = merchantBalanceTransactionssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateMerchantBalanceTransactionsQuery(merchantBalanceTransactionssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistMerchantBalanceTransactionsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceTransactions] failed checking merchantBalanceTransactions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceTransactions with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeMerchantBalanceTransactionsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"merchant_balance_transactions\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceTransactions] failed exec query")
	}
	return
}

type MerchantBalanceTransactionsFieldParameter struct {
	param string
	args  []interface{}
}

func NewMerchantBalanceTransactionsFieldParameter(param string, args ...interface{}) MerchantBalanceTransactionsFieldParameter {
	return MerchantBalanceTransactionsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateMerchantBalanceTransactionsQuery(mapMerchantBalanceTransactionss map[model.MerchantBalanceTransactionsPrimaryID]MerchantBalanceTransactionsUpdateFieldList, asTableValues string) (primaryIDs []model.MerchantBalanceTransactionsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.MerchantBalanceTransactionsPrimaryID]map[string]interface{}{}
	merchantBalanceTransactionsSelectFields := NewMerchantBalanceTransactionsSelectFields()
	for id, updateFields := range mapMerchantBalanceTransactionss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.merchantBalanceTransactionsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapMerchantBalanceTransactionss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetMerchantBalanceTransactionsFieldType(updateField.merchantBalanceTransactionsField)))
			args = append(args, fields[string(updateField.merchantBalanceTransactionsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.merchantBalanceTransactionsField))
		if updateField.merchantBalanceTransactionsField == merchantBalanceTransactionsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.merchantBalanceTransactionsField, asTableValues, updateField.merchantBalanceTransactionsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.merchantBalanceTransactionsField,
				"\"merchant_balance_transactions\"", updateField.merchantBalanceTransactionsField,
				asTableValues, updateField.merchantBalanceTransactionsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeMerchantBalanceTransactionsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.MerchantBalanceTransactionsPrimaryID, asTableValue string) (whereQry string) {
	merchantBalanceTransactionsSelectFields := NewMerchantBalanceTransactionsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"merchant_balance_transactions\".\"id\" = %s.\"id\"::"+GetMerchantBalanceTransactionsFieldType(merchantBalanceTransactionsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetMerchantBalanceTransactionsFieldType(merchantBalanceTransactionsField MerchantBalanceTransactionsField) string {
	selectMerchantBalanceTransactionsFields := NewMerchantBalanceTransactionsSelectFields()
	switch merchantBalanceTransactionsField {

	case selectMerchantBalanceTransactionsFields.Id():
		return "uuid"

	case selectMerchantBalanceTransactionsFields.BalanceAccountId():
		return "uuid"

	case selectMerchantBalanceTransactionsFields.MerchantPartyId():
		return "uuid"

	case selectMerchantBalanceTransactionsFields.BalanceType():
		return "merchant_balance_transactions_balance_type_enum"

	case selectMerchantBalanceTransactionsFields.CurrencyCode():
		return "text"

	case selectMerchantBalanceTransactionsFields.SourceType():
		return "text"

	case selectMerchantBalanceTransactionsFields.SourceId():
		return "uuid"

	case selectMerchantBalanceTransactionsFields.Direction():
		return "direction_enum"

	case selectMerchantBalanceTransactionsFields.Amount():
		return "numeric"

	case selectMerchantBalanceTransactionsFields.BalanceBefore():
		return "numeric"

	case selectMerchantBalanceTransactionsFields.BalanceAfter():
		return "numeric"

	case selectMerchantBalanceTransactionsFields.ReasonCode():
		return "text"

	case selectMerchantBalanceTransactionsFields.Metadata():
		return "jsonb"

	case selectMerchantBalanceTransactionsFields.MetaCreatedAt():
		return "timestamptz"

	case selectMerchantBalanceTransactionsFields.MetaCreatedBy():
		return "uuid"

	case selectMerchantBalanceTransactionsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectMerchantBalanceTransactionsFields.MetaUpdatedBy():
		return "uuid"

	case selectMerchantBalanceTransactionsFields.MetaDeletedAt():
		return "timestamptz"

	case selectMerchantBalanceTransactionsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateMerchantBalanceTransactions(ctx context.Context, merchantBalanceTransactions *model.MerchantBalanceTransactions, fieldsInsert ...MerchantBalanceTransactionsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceTransactionsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.MerchantBalanceTransactionsPrimaryID{
		Id: merchantBalanceTransactions.Id,
	}
	exists, err := repo.IsExistMerchantBalanceTransactionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceTransactions] failed checking merchantBalanceTransactions whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "merchantBalanceTransactions", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsMerchantBalanceTransactions([]model.MerchantBalanceTransactions{*merchantBalanceTransactions}, fieldsInsert...)
	commandQuery := fmt.Sprintf(merchantBalanceTransactionsQueries.insertMerchantBalanceTransactions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceTransactions] failed exec create merchantBalanceTransactions query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteMerchantBalanceTransactionsByID(ctx context.Context, primaryID model.MerchantBalanceTransactionsPrimaryID) (err error) {
	exists, err := repo.IsExistMerchantBalanceTransactionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceTransactionsByID] failed checking merchantBalanceTransactions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceTransactions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeMerchantBalanceTransactionsCompositePrimaryKeyWhere([]model.MerchantBalanceTransactionsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(merchantBalanceTransactionsQueries.deleteMerchantBalanceTransactions + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceTransactionsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceTransactionsByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceTransactionsFilterResult, err error) {
	query, args, err := composeMerchantBalanceTransactionsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceTransactionsByFilter] failed compose merchantBalanceTransactions filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceTransactionsByFilter] failed get merchantBalanceTransactions by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeMerchantBalanceTransactionsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.MerchantBalanceTransactionsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeMerchantBalanceTransactionsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeMerchantBalanceTransactionsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeMerchantBalanceTransactionsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewMerchantBalanceTransactionsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["balance_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"balance_account_id\"")
			selectedColumns["balance_account_id"] = struct{}{}
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
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["direction"]; !selected {
			selectColumns = append(selectColumns, "base.\"direction\"")
			selectedColumns["direction"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["balance_before"]; !selected {
			selectColumns = append(selectColumns, "base.\"balance_before\"")
			selectedColumns["balance_before"] = struct{}{}
		}
		if _, selected := selectedColumns["balance_after"]; !selected {
			selectColumns = append(selectColumns, "base.\"balance_after\"")
			selectedColumns["balance_after"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
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

type merchantBalanceTransactionsFilterPlaceholder struct {
	index int
}

func (p *merchantBalanceTransactionsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeMerchantBalanceTransactionsFilterPredicate(filterField model.FilterField, placeholders *merchantBalanceTransactionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewMerchantBalanceTransactionsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeMerchantBalanceTransactionsFilterSQLExpr(spec)
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

func composeMerchantBalanceTransactionsFilterGroup(group model.FilterGroup, placeholders *merchantBalanceTransactionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeMerchantBalanceTransactionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeMerchantBalanceTransactionsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeMerchantBalanceTransactionsFilterWhereQueries(filter model.Filter, placeholders *merchantBalanceTransactionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeMerchantBalanceTransactionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeMerchantBalanceTransactionsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeMerchantBalanceTransactionsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateMerchantBalanceTransactionsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeMerchantBalanceTransactionsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeMerchantBalanceTransactionsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := merchantBalanceTransactionsFilterPlaceholder{index: 1}
	whereQueries, err := composeMerchantBalanceTransactionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewMerchantBalanceTransactionsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeMerchantBalanceTransactionsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeMerchantBalanceTransactionsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"merchant_balance_transactions\" base%s", strings.Join(selectColumns, ","), composeMerchantBalanceTransactionsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistMerchantBalanceTransactionsByID(ctx context.Context, primaryID model.MerchantBalanceTransactionsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeMerchantBalanceTransactionsCompositePrimaryKeyWhere([]model.MerchantBalanceTransactionsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", merchantBalanceTransactionsQueries.selectCountMerchantBalanceTransactions, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceTransactionsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceTransactions(ctx context.Context, selectFields ...MerchantBalanceTransactionsField) (merchantBalanceTransactionsList model.MerchantBalanceTransactionsList, err error) {
	var (
		defaultMerchantBalanceTransactionsSelectFields = defaultMerchantBalanceTransactionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceTransactionsSelectFields = composeMerchantBalanceTransactionsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(merchantBalanceTransactionsQueries.selectMerchantBalanceTransactions, defaultMerchantBalanceTransactionsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &merchantBalanceTransactionsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceTransactions] failed get merchantBalanceTransactions list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceTransactionsByID(ctx context.Context, primaryID model.MerchantBalanceTransactionsPrimaryID, selectFields ...MerchantBalanceTransactionsField) (merchantBalanceTransactions model.MerchantBalanceTransactions, err error) {
	var (
		defaultMerchantBalanceTransactionsSelectFields = defaultMerchantBalanceTransactionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceTransactionsSelectFields = composeMerchantBalanceTransactionsSelectFields(selectFields...)
	}
	whereQry, params := composeMerchantBalanceTransactionsCompositePrimaryKeyWhere([]model.MerchantBalanceTransactionsPrimaryID{primaryID})
	query := fmt.Sprintf(merchantBalanceTransactionsQueries.selectMerchantBalanceTransactions+" WHERE "+whereQry, defaultMerchantBalanceTransactionsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &merchantBalanceTransactions, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("merchantBalanceTransactions with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveMerchantBalanceTransactionsByID] failed get merchantBalanceTransactions")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateMerchantBalanceTransactionsByID(ctx context.Context, primaryID model.MerchantBalanceTransactionsPrimaryID, merchantBalanceTransactions *model.MerchantBalanceTransactions, merchantBalanceTransactionsUpdateFields ...MerchantBalanceTransactionsUpdateField) (err error) {
	exists, err := repo.IsExistMerchantBalanceTransactionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceTransactions] failed checking merchantBalanceTransactions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceTransactions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if merchantBalanceTransactions == nil {
		if len(merchantBalanceTransactionsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateMerchantBalanceTransactionsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		merchantBalanceTransactions = &model.MerchantBalanceTransactions{}
	}
	var (
		defaultMerchantBalanceTransactionsUpdateFields = defaultMerchantBalanceTransactionsUpdateFields(*merchantBalanceTransactions)
		tempUpdateField                                MerchantBalanceTransactionsUpdateFieldList
		selectFields                                   = NewMerchantBalanceTransactionsSelectFields()
	)
	if len(merchantBalanceTransactionsUpdateFields) > 0 {
		for _, updateField := range merchantBalanceTransactionsUpdateFields {
			if updateField.merchantBalanceTransactionsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultMerchantBalanceTransactionsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeMerchantBalanceTransactionsCompositePrimaryKeyWhere([]model.MerchantBalanceTransactionsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsMerchantBalanceTransactionsCommand(defaultMerchantBalanceTransactionsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(merchantBalanceTransactionsQueries.updateMerchantBalanceTransactions+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceTransactions] error when try to update merchantBalanceTransactions by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateMerchantBalanceTransactionsByFilter(ctx context.Context, filter model.Filter, merchantBalanceTransactionsUpdateFields ...MerchantBalanceTransactionsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(merchantBalanceTransactionsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields MerchantBalanceTransactionsUpdateFieldList
		selectFields = NewMerchantBalanceTransactionsSelectFields()
	)
	for _, updateField := range merchantBalanceTransactionsUpdateFields {
		if updateField.merchantBalanceTransactionsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsMerchantBalanceTransactionsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := merchantBalanceTransactionsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeMerchantBalanceTransactionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"merchant_balance_transactions\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceTransactionsByFilter] error when try to update merchantBalanceTransactions by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceTransactionsByFilter] failed get rows affected")
	}
	return
}

var (
	merchantBalanceTransactionsQueries = struct {
		selectMerchantBalanceTransactions      string
		selectCountMerchantBalanceTransactions string
		deleteMerchantBalanceTransactions      string
		updateMerchantBalanceTransactions      string
		insertMerchantBalanceTransactions      string
	}{
		selectMerchantBalanceTransactions:      "SELECT %s FROM \"merchant_balance_transactions\"",
		selectCountMerchantBalanceTransactions: "SELECT COUNT(\"id\") FROM \"merchant_balance_transactions\"",
		deleteMerchantBalanceTransactions:      "DELETE FROM \"merchant_balance_transactions\"",
		updateMerchantBalanceTransactions:      "UPDATE \"merchant_balance_transactions\" SET %s ",
		insertMerchantBalanceTransactions:      "INSERT INTO \"merchant_balance_transactions\" %s VALUES %s",
	}
)

type MerchantBalanceTransactionsRepository interface {
	CreateMerchantBalanceTransactions(ctx context.Context, merchantBalanceTransactions *model.MerchantBalanceTransactions, fieldsInsert ...MerchantBalanceTransactionsField) error
	BulkCreateMerchantBalanceTransactions(ctx context.Context, merchantBalanceTransactionsList []*model.MerchantBalanceTransactions, fieldsInsert ...MerchantBalanceTransactionsField) error
	ResolveMerchantBalanceTransactions(ctx context.Context, selectFields ...MerchantBalanceTransactionsField) (model.MerchantBalanceTransactionsList, error)
	ResolveMerchantBalanceTransactionsByID(ctx context.Context, primaryID model.MerchantBalanceTransactionsPrimaryID, selectFields ...MerchantBalanceTransactionsField) (model.MerchantBalanceTransactions, error)
	UpdateMerchantBalanceTransactionsByID(ctx context.Context, id model.MerchantBalanceTransactionsPrimaryID, merchantBalanceTransactions *model.MerchantBalanceTransactions, merchantBalanceTransactionsUpdateFields ...MerchantBalanceTransactionsUpdateField) error
	UpdateMerchantBalanceTransactionsByFilter(ctx context.Context, filter model.Filter, merchantBalanceTransactionsUpdateFields ...MerchantBalanceTransactionsUpdateField) (rowsAffected int64, err error)
	BulkUpdateMerchantBalanceTransactions(ctx context.Context, merchantBalanceTransactionsListMap map[model.MerchantBalanceTransactionsPrimaryID]*model.MerchantBalanceTransactions, MerchantBalanceTransactionssMapUpdateFieldsRequest map[model.MerchantBalanceTransactionsPrimaryID]MerchantBalanceTransactionsUpdateFieldList) (err error)
	DeleteMerchantBalanceTransactionsByID(ctx context.Context, id model.MerchantBalanceTransactionsPrimaryID) error
	BulkDeleteMerchantBalanceTransactionsByIDs(ctx context.Context, ids []model.MerchantBalanceTransactionsPrimaryID) error
	ResolveMerchantBalanceTransactionsByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceTransactionsFilterResult, err error)
	IsExistMerchantBalanceTransactionsByIDs(ctx context.Context, ids []model.MerchantBalanceTransactionsPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceTransactionsPrimaryID, err error)
	IsExistMerchantBalanceTransactionsByID(ctx context.Context, id model.MerchantBalanceTransactionsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
