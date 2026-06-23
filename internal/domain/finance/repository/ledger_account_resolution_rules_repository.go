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

func composeInsertFieldsAndParamsLedgerAccountResolutionRules(ledgerAccountResolutionRulesList []model.LedgerAccountResolutionRules, fieldsInsert ...LedgerAccountResolutionRulesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerAccountResolutionRulesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerAccountResolutionRules := range ledgerAccountResolutionRulesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerAccountResolutionRules.Id)
			case selectField.BookId():
				args = append(args, ledgerAccountResolutionRules.BookId)
			case selectField.Purpose():
				args = append(args, ledgerAccountResolutionRules.Purpose)
			case selectField.MerchantPartyId():
				args = append(args, ledgerAccountResolutionRules.MerchantPartyId)
			case selectField.ProviderCode():
				args = append(args, ledgerAccountResolutionRules.ProviderCode)
			case selectField.CurrencyCode():
				args = append(args, ledgerAccountResolutionRules.CurrencyCode)
			case selectField.SourceType():
				args = append(args, ledgerAccountResolutionRules.SourceType)
			case selectField.SourceSubtype():
				args = append(args, ledgerAccountResolutionRules.SourceSubtype)
			case selectField.AccountId():
				args = append(args, ledgerAccountResolutionRules.AccountId)
			case selectField.Priority():
				args = append(args, ledgerAccountResolutionRules.Priority)
			case selectField.IsActive():
				args = append(args, ledgerAccountResolutionRules.IsActive)
			case selectField.Conditions():
				args = append(args, ledgerAccountResolutionRules.Conditions)
			case selectField.Metadata():
				args = append(args, ledgerAccountResolutionRules.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerAccountResolutionRules.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerAccountResolutionRules.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerAccountResolutionRules.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerAccountResolutionRules.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerAccountResolutionRules.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerAccountResolutionRules.MetaDeletedBy)

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

func composeLedgerAccountResolutionRulesCompositePrimaryKeyWhere(primaryIDs []model.LedgerAccountResolutionRulesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_account_resolution_rules\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerAccountResolutionRulesSelectFields() string {
	fields := NewLedgerAccountResolutionRulesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerAccountResolutionRulesSelectFields(selectFields ...LedgerAccountResolutionRulesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerAccountResolutionRulesField string
type LedgerAccountResolutionRulesFieldList []LedgerAccountResolutionRulesField

type LedgerAccountResolutionRulesSelectFields struct {
}

func (ss LedgerAccountResolutionRulesSelectFields) Id() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("id")
}

func (ss LedgerAccountResolutionRulesSelectFields) BookId() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("book_id")
}

func (ss LedgerAccountResolutionRulesSelectFields) Purpose() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("purpose")
}

func (ss LedgerAccountResolutionRulesSelectFields) MerchantPartyId() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("merchant_party_id")
}

func (ss LedgerAccountResolutionRulesSelectFields) ProviderCode() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("provider_code")
}

func (ss LedgerAccountResolutionRulesSelectFields) CurrencyCode() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("currency_code")
}

func (ss LedgerAccountResolutionRulesSelectFields) SourceType() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("source_type")
}

func (ss LedgerAccountResolutionRulesSelectFields) SourceSubtype() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("source_subtype")
}

func (ss LedgerAccountResolutionRulesSelectFields) AccountId() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("account_id")
}

func (ss LedgerAccountResolutionRulesSelectFields) Priority() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("priority")
}

func (ss LedgerAccountResolutionRulesSelectFields) IsActive() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("is_active")
}

func (ss LedgerAccountResolutionRulesSelectFields) Conditions() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("conditions")
}

func (ss LedgerAccountResolutionRulesSelectFields) Metadata() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("metadata")
}

func (ss LedgerAccountResolutionRulesSelectFields) MetaCreatedAt() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("meta_created_at")
}

func (ss LedgerAccountResolutionRulesSelectFields) MetaCreatedBy() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("meta_created_by")
}

func (ss LedgerAccountResolutionRulesSelectFields) MetaUpdatedAt() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("meta_updated_at")
}

func (ss LedgerAccountResolutionRulesSelectFields) MetaUpdatedBy() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("meta_updated_by")
}

func (ss LedgerAccountResolutionRulesSelectFields) MetaDeletedAt() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("meta_deleted_at")
}

func (ss LedgerAccountResolutionRulesSelectFields) MetaDeletedBy() LedgerAccountResolutionRulesField {
	return LedgerAccountResolutionRulesField("meta_deleted_by")
}

func (ss LedgerAccountResolutionRulesSelectFields) All() LedgerAccountResolutionRulesFieldList {
	return []LedgerAccountResolutionRulesField{
		ss.Id(),
		ss.BookId(),
		ss.Purpose(),
		ss.MerchantPartyId(),
		ss.ProviderCode(),
		ss.CurrencyCode(),
		ss.SourceType(),
		ss.SourceSubtype(),
		ss.AccountId(),
		ss.Priority(),
		ss.IsActive(),
		ss.Conditions(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerAccountResolutionRulesSelectFields() LedgerAccountResolutionRulesSelectFields {
	return LedgerAccountResolutionRulesSelectFields{}
}

type LedgerAccountResolutionRulesUpdateFieldOption struct {
	useIncrement bool
}
type LedgerAccountResolutionRulesUpdateField struct {
	ledgerAccountResolutionRulesField LedgerAccountResolutionRulesField
	opt                               LedgerAccountResolutionRulesUpdateFieldOption
	value                             interface{}
}
type LedgerAccountResolutionRulesUpdateFieldList []LedgerAccountResolutionRulesUpdateField

func defaultLedgerAccountResolutionRulesUpdateFieldOption() LedgerAccountResolutionRulesUpdateFieldOption {
	return LedgerAccountResolutionRulesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerAccountResolutionRulesOption(useIncrement bool) func(*LedgerAccountResolutionRulesUpdateFieldOption) {
	return func(pcufo *LedgerAccountResolutionRulesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerAccountResolutionRulesUpdateField(field LedgerAccountResolutionRulesField, val interface{}, opts ...func(*LedgerAccountResolutionRulesUpdateFieldOption)) LedgerAccountResolutionRulesUpdateField {
	defaultOpt := defaultLedgerAccountResolutionRulesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerAccountResolutionRulesUpdateField{
		ledgerAccountResolutionRulesField: field,
		value:                             val,
		opt:                               defaultOpt,
	}
}
func defaultLedgerAccountResolutionRulesUpdateFields(ledgerAccountResolutionRules model.LedgerAccountResolutionRules) (ledgerAccountResolutionRulesUpdateFieldList LedgerAccountResolutionRulesUpdateFieldList) {
	selectFields := NewLedgerAccountResolutionRulesSelectFields()
	ledgerAccountResolutionRulesUpdateFieldList = append(ledgerAccountResolutionRulesUpdateFieldList,
		NewLedgerAccountResolutionRulesUpdateField(selectFields.Id(), ledgerAccountResolutionRules.Id),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.BookId(), ledgerAccountResolutionRules.BookId),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.Purpose(), ledgerAccountResolutionRules.Purpose),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.MerchantPartyId(), ledgerAccountResolutionRules.MerchantPartyId),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.ProviderCode(), ledgerAccountResolutionRules.ProviderCode),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.CurrencyCode(), ledgerAccountResolutionRules.CurrencyCode),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.SourceType(), ledgerAccountResolutionRules.SourceType),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.SourceSubtype(), ledgerAccountResolutionRules.SourceSubtype),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.AccountId(), ledgerAccountResolutionRules.AccountId),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.Priority(), ledgerAccountResolutionRules.Priority),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.IsActive(), ledgerAccountResolutionRules.IsActive),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.Conditions(), ledgerAccountResolutionRules.Conditions),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.Metadata(), ledgerAccountResolutionRules.Metadata),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.MetaCreatedAt(), ledgerAccountResolutionRules.MetaCreatedAt),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.MetaCreatedBy(), ledgerAccountResolutionRules.MetaCreatedBy),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.MetaUpdatedAt(), ledgerAccountResolutionRules.MetaUpdatedAt),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.MetaUpdatedBy(), ledgerAccountResolutionRules.MetaUpdatedBy),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.MetaDeletedAt(), ledgerAccountResolutionRules.MetaDeletedAt),
		NewLedgerAccountResolutionRulesUpdateField(selectFields.MetaDeletedBy(), ledgerAccountResolutionRules.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerAccountResolutionRulesCommand(ledgerAccountResolutionRulesUpdateFieldList LedgerAccountResolutionRulesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerAccountResolutionRulesUpdateFieldList {
		field := string(updateField.ledgerAccountResolutionRulesField)
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

func (repo *RepositoryImpl) BulkCreateLedgerAccountResolutionRules(ctx context.Context, ledgerAccountResolutionRulesList []*model.LedgerAccountResolutionRules, fieldsInsert ...LedgerAccountResolutionRulesField) (err error) {
	var (
		fieldsStr                             string
		valueListStr                          []string
		argsList                              []interface{}
		primaryIds                            []model.LedgerAccountResolutionRulesPrimaryID
		ledgerAccountResolutionRulesValueList []model.LedgerAccountResolutionRules
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountResolutionRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerAccountResolutionRules := range ledgerAccountResolutionRulesList {

		primaryIds = append(primaryIds, ledgerAccountResolutionRules.ToLedgerAccountResolutionRulesPrimaryID())

		ledgerAccountResolutionRulesValueList = append(ledgerAccountResolutionRulesValueList, *ledgerAccountResolutionRules)
	}

	_, notFoundIds, err := repo.IsExistLedgerAccountResolutionRulesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccountResolutionRules] failed checking ledgerAccountResolutionRules whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerAccountResolutionRulesPrimaryID{}
		mapNotFoundIds := map[model.LedgerAccountResolutionRulesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerAccountResolutionRules", fmt.Sprintf("ledgerAccountResolutionRules with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerAccountResolutionRules(ledgerAccountResolutionRulesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerAccountResolutionRulesQueries.insertLedgerAccountResolutionRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccountResolutionRules] failed exec create ledgerAccountResolutionRules query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerAccountResolutionRulesByIDs(ctx context.Context, primaryIDs []model.LedgerAccountResolutionRulesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerAccountResolutionRulesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountResolutionRulesByIDs] failed checking ledgerAccountResolutionRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountResolutionRules with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_account_resolution_rules\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerAccountResolutionRulesQueries.deleteLedgerAccountResolutionRules + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountResolutionRulesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountResolutionRulesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerAccountResolutionRulesByIDs(ctx context.Context, ids []model.LedgerAccountResolutionRulesPrimaryID) (exists bool, notFoundIds []model.LedgerAccountResolutionRulesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_account_resolution_rules\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerAccountResolutionRulesQueries.selectLedgerAccountResolutionRules, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountResolutionRulesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerAccountResolutionRulesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountResolutionRulesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerAccountResolutionRulesPrimaryID]bool{}
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

// BulkUpdateLedgerAccountResolutionRules is used to bulk update ledgerAccountResolutionRules, by default it will update all field
// if want to update specific field, then fill ledgerAccountResolutionRulessMapUpdateFieldsRequest else please fill ledgerAccountResolutionRulessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerAccountResolutionRules(ctx context.Context, ledgerAccountResolutionRulessMap map[model.LedgerAccountResolutionRulesPrimaryID]*model.LedgerAccountResolutionRules, ledgerAccountResolutionRulessMapUpdateFieldsRequest map[model.LedgerAccountResolutionRulesPrimaryID]LedgerAccountResolutionRulesUpdateFieldList) (err error) {
	if len(ledgerAccountResolutionRulessMap) == 0 && len(ledgerAccountResolutionRulessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerAccountResolutionRulessMapUpdateField map[model.LedgerAccountResolutionRulesPrimaryID]LedgerAccountResolutionRulesUpdateFieldList = map[model.LedgerAccountResolutionRulesPrimaryID]LedgerAccountResolutionRulesUpdateFieldList{}
		asTableValues                               string                                                                                      = "myvalues"
	)

	if len(ledgerAccountResolutionRulessMap) > 0 {
		for id, ledgerAccountResolutionRules := range ledgerAccountResolutionRulessMap {
			if ledgerAccountResolutionRules == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerAccountResolutionRules] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerAccountResolutionRulessMapUpdateField[id] = defaultLedgerAccountResolutionRulesUpdateFields(*ledgerAccountResolutionRules)
		}
	} else {
		ledgerAccountResolutionRulessMapUpdateField = ledgerAccountResolutionRulessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerAccountResolutionRulesQuery(ledgerAccountResolutionRulessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerAccountResolutionRulesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccountResolutionRules] failed checking ledgerAccountResolutionRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountResolutionRules with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerAccountResolutionRulesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_account_resolution_rules\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccountResolutionRules] failed exec query")
	}
	return
}

type LedgerAccountResolutionRulesFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerAccountResolutionRulesFieldParameter(param string, args ...interface{}) LedgerAccountResolutionRulesFieldParameter {
	return LedgerAccountResolutionRulesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerAccountResolutionRulesQuery(mapLedgerAccountResolutionRuless map[model.LedgerAccountResolutionRulesPrimaryID]LedgerAccountResolutionRulesUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerAccountResolutionRulesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerAccountResolutionRulesPrimaryID]map[string]interface{}{}
	ledgerAccountResolutionRulesSelectFields := NewLedgerAccountResolutionRulesSelectFields()
	for id, updateFields := range mapLedgerAccountResolutionRuless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerAccountResolutionRulesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerAccountResolutionRuless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerAccountResolutionRulesFieldType(updateField.ledgerAccountResolutionRulesField)))
			args = append(args, fields[string(updateField.ledgerAccountResolutionRulesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerAccountResolutionRulesField))
		if updateField.ledgerAccountResolutionRulesField == ledgerAccountResolutionRulesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerAccountResolutionRulesField, asTableValues, updateField.ledgerAccountResolutionRulesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerAccountResolutionRulesField,
				"\"ledger_account_resolution_rules\"", updateField.ledgerAccountResolutionRulesField,
				asTableValues, updateField.ledgerAccountResolutionRulesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerAccountResolutionRulesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerAccountResolutionRulesPrimaryID, asTableValue string) (whereQry string) {
	ledgerAccountResolutionRulesSelectFields := NewLedgerAccountResolutionRulesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_account_resolution_rules\".\"id\" = %s.\"id\"::"+GetLedgerAccountResolutionRulesFieldType(ledgerAccountResolutionRulesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerAccountResolutionRulesFieldType(ledgerAccountResolutionRulesField LedgerAccountResolutionRulesField) string {
	selectLedgerAccountResolutionRulesFields := NewLedgerAccountResolutionRulesSelectFields()
	switch ledgerAccountResolutionRulesField {

	case selectLedgerAccountResolutionRulesFields.Id():
		return "uuid"

	case selectLedgerAccountResolutionRulesFields.BookId():
		return "uuid"

	case selectLedgerAccountResolutionRulesFields.Purpose():
		return "purpose_enum"

	case selectLedgerAccountResolutionRulesFields.MerchantPartyId():
		return "uuid"

	case selectLedgerAccountResolutionRulesFields.ProviderCode():
		return "text"

	case selectLedgerAccountResolutionRulesFields.CurrencyCode():
		return "text"

	case selectLedgerAccountResolutionRulesFields.SourceType():
		return "text"

	case selectLedgerAccountResolutionRulesFields.SourceSubtype():
		return "text"

	case selectLedgerAccountResolutionRulesFields.AccountId():
		return "uuid"

	case selectLedgerAccountResolutionRulesFields.Priority():
		return "int4"

	case selectLedgerAccountResolutionRulesFields.IsActive():
		return "bool"

	case selectLedgerAccountResolutionRulesFields.Conditions():
		return "jsonb"

	case selectLedgerAccountResolutionRulesFields.Metadata():
		return "jsonb"

	case selectLedgerAccountResolutionRulesFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerAccountResolutionRulesFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerAccountResolutionRulesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerAccountResolutionRulesFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerAccountResolutionRulesFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerAccountResolutionRulesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerAccountResolutionRules(ctx context.Context, ledgerAccountResolutionRules *model.LedgerAccountResolutionRules, fieldsInsert ...LedgerAccountResolutionRulesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountResolutionRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerAccountResolutionRulesPrimaryID{
		Id: ledgerAccountResolutionRules.Id,
	}
	exists, err := repo.IsExistLedgerAccountResolutionRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccountResolutionRules] failed checking ledgerAccountResolutionRules whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerAccountResolutionRules", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerAccountResolutionRules([]model.LedgerAccountResolutionRules{*ledgerAccountResolutionRules}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerAccountResolutionRulesQueries.insertLedgerAccountResolutionRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccountResolutionRules] failed exec create ledgerAccountResolutionRules query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerAccountResolutionRulesByID(ctx context.Context, primaryID model.LedgerAccountResolutionRulesPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerAccountResolutionRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountResolutionRulesByID] failed checking ledgerAccountResolutionRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountResolutionRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerAccountResolutionRulesCompositePrimaryKeyWhere([]model.LedgerAccountResolutionRulesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerAccountResolutionRulesQueries.deleteLedgerAccountResolutionRules + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountResolutionRulesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountResolutionRulesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountResolutionRulesFilterResult, err error) {
	query, args, err := composeLedgerAccountResolutionRulesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountResolutionRulesByFilter] failed compose ledgerAccountResolutionRules filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountResolutionRulesByFilter] failed get ledgerAccountResolutionRules by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerAccountResolutionRulesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerAccountResolutionRulesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerAccountResolutionRulesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerAccountResolutionRulesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerAccountResolutionRulesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["book_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_id\"")
			selectedColumns["book_id"] = struct{}{}
		}
		if _, selected := selectedColumns["purpose"]; !selected {
			selectColumns = append(selectColumns, "base.\"purpose\"")
			selectedColumns["purpose"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_code\"")
			selectedColumns["provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_subtype"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_subtype\"")
			selectedColumns["source_subtype"] = struct{}{}
		}
		if _, selected := selectedColumns["account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_id\"")
			selectedColumns["account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["priority"]; !selected {
			selectColumns = append(selectColumns, "base.\"priority\"")
			selectedColumns["priority"] = struct{}{}
		}
		if _, selected := selectedColumns["is_active"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_active\"")
			selectedColumns["is_active"] = struct{}{}
		}
		if _, selected := selectedColumns["conditions"]; !selected {
			selectColumns = append(selectColumns, "base.\"conditions\"")
			selectedColumns["conditions"] = struct{}{}
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

type ledgerAccountResolutionRulesFilterPlaceholder struct {
	index int
}

func (p *ledgerAccountResolutionRulesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerAccountResolutionRulesFilterPredicate(filterField model.FilterField, placeholders *ledgerAccountResolutionRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerAccountResolutionRulesFilterSQLExpr(spec)
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

func composeLedgerAccountResolutionRulesFilterGroup(group model.FilterGroup, placeholders *ledgerAccountResolutionRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerAccountResolutionRulesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerAccountResolutionRulesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerAccountResolutionRulesFilterWhereQueries(filter model.Filter, placeholders *ledgerAccountResolutionRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerAccountResolutionRulesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerAccountResolutionRulesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerAccountResolutionRulesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerAccountResolutionRulesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerAccountResolutionRulesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerAccountResolutionRulesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerAccountResolutionRulesFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerAccountResolutionRulesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerAccountResolutionRulesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerAccountResolutionRulesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerAccountResolutionRulesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_account_resolution_rules\" base%s", strings.Join(selectColumns, ","), composeLedgerAccountResolutionRulesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerAccountResolutionRulesByID(ctx context.Context, primaryID model.LedgerAccountResolutionRulesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerAccountResolutionRulesCompositePrimaryKeyWhere([]model.LedgerAccountResolutionRulesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerAccountResolutionRulesQueries.selectCountLedgerAccountResolutionRules, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountResolutionRulesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountResolutionRules(ctx context.Context, selectFields ...LedgerAccountResolutionRulesField) (ledgerAccountResolutionRulesList model.LedgerAccountResolutionRulesList, err error) {
	var (
		defaultLedgerAccountResolutionRulesSelectFields = defaultLedgerAccountResolutionRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountResolutionRulesSelectFields = composeLedgerAccountResolutionRulesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerAccountResolutionRulesQueries.selectLedgerAccountResolutionRules, defaultLedgerAccountResolutionRulesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerAccountResolutionRulesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountResolutionRules] failed get ledgerAccountResolutionRules list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountResolutionRulesByID(ctx context.Context, primaryID model.LedgerAccountResolutionRulesPrimaryID, selectFields ...LedgerAccountResolutionRulesField) (ledgerAccountResolutionRules model.LedgerAccountResolutionRules, err error) {
	var (
		defaultLedgerAccountResolutionRulesSelectFields = defaultLedgerAccountResolutionRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountResolutionRulesSelectFields = composeLedgerAccountResolutionRulesSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerAccountResolutionRulesCompositePrimaryKeyWhere([]model.LedgerAccountResolutionRulesPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerAccountResolutionRulesQueries.selectLedgerAccountResolutionRules+" WHERE "+whereQry, defaultLedgerAccountResolutionRulesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerAccountResolutionRules, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerAccountResolutionRules with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerAccountResolutionRulesByID] failed get ledgerAccountResolutionRules")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerAccountResolutionRulesByID(ctx context.Context, primaryID model.LedgerAccountResolutionRulesPrimaryID, ledgerAccountResolutionRules *model.LedgerAccountResolutionRules, ledgerAccountResolutionRulesUpdateFields ...LedgerAccountResolutionRulesUpdateField) (err error) {
	exists, err := repo.IsExistLedgerAccountResolutionRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountResolutionRules] failed checking ledgerAccountResolutionRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountResolutionRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerAccountResolutionRules == nil {
		if len(ledgerAccountResolutionRulesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerAccountResolutionRulesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerAccountResolutionRules = &model.LedgerAccountResolutionRules{}
	}
	var (
		defaultLedgerAccountResolutionRulesUpdateFields = defaultLedgerAccountResolutionRulesUpdateFields(*ledgerAccountResolutionRules)
		tempUpdateField                                 LedgerAccountResolutionRulesUpdateFieldList
		selectFields                                    = NewLedgerAccountResolutionRulesSelectFields()
	)
	if len(ledgerAccountResolutionRulesUpdateFields) > 0 {
		for _, updateField := range ledgerAccountResolutionRulesUpdateFields {
			if updateField.ledgerAccountResolutionRulesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerAccountResolutionRulesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerAccountResolutionRulesCompositePrimaryKeyWhere([]model.LedgerAccountResolutionRulesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerAccountResolutionRulesCommand(defaultLedgerAccountResolutionRulesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerAccountResolutionRulesQueries.updateLedgerAccountResolutionRules+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountResolutionRules] error when try to update ledgerAccountResolutionRules by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerAccountResolutionRulesByFilter(ctx context.Context, filter model.Filter, ledgerAccountResolutionRulesUpdateFields ...LedgerAccountResolutionRulesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerAccountResolutionRulesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerAccountResolutionRulesUpdateFieldList
		selectFields = NewLedgerAccountResolutionRulesSelectFields()
	)
	for _, updateField := range ledgerAccountResolutionRulesUpdateFields {
		if updateField.ledgerAccountResolutionRulesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerAccountResolutionRulesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerAccountResolutionRulesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerAccountResolutionRulesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_account_resolution_rules\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountResolutionRulesByFilter] error when try to update ledgerAccountResolutionRules by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountResolutionRulesByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerAccountResolutionRulesQueries = struct {
		selectLedgerAccountResolutionRules      string
		selectCountLedgerAccountResolutionRules string
		deleteLedgerAccountResolutionRules      string
		updateLedgerAccountResolutionRules      string
		insertLedgerAccountResolutionRules      string
	}{
		selectLedgerAccountResolutionRules:      "SELECT %s FROM \"ledger_account_resolution_rules\"",
		selectCountLedgerAccountResolutionRules: "SELECT COUNT(\"id\") FROM \"ledger_account_resolution_rules\"",
		deleteLedgerAccountResolutionRules:      "DELETE FROM \"ledger_account_resolution_rules\"",
		updateLedgerAccountResolutionRules:      "UPDATE \"ledger_account_resolution_rules\" SET %s ",
		insertLedgerAccountResolutionRules:      "INSERT INTO \"ledger_account_resolution_rules\" %s VALUES %s",
	}
)

type LedgerAccountResolutionRulesRepository interface {
	CreateLedgerAccountResolutionRules(ctx context.Context, ledgerAccountResolutionRules *model.LedgerAccountResolutionRules, fieldsInsert ...LedgerAccountResolutionRulesField) error
	BulkCreateLedgerAccountResolutionRules(ctx context.Context, ledgerAccountResolutionRulesList []*model.LedgerAccountResolutionRules, fieldsInsert ...LedgerAccountResolutionRulesField) error
	ResolveLedgerAccountResolutionRules(ctx context.Context, selectFields ...LedgerAccountResolutionRulesField) (model.LedgerAccountResolutionRulesList, error)
	ResolveLedgerAccountResolutionRulesByID(ctx context.Context, primaryID model.LedgerAccountResolutionRulesPrimaryID, selectFields ...LedgerAccountResolutionRulesField) (model.LedgerAccountResolutionRules, error)
	UpdateLedgerAccountResolutionRulesByID(ctx context.Context, id model.LedgerAccountResolutionRulesPrimaryID, ledgerAccountResolutionRules *model.LedgerAccountResolutionRules, ledgerAccountResolutionRulesUpdateFields ...LedgerAccountResolutionRulesUpdateField) error
	UpdateLedgerAccountResolutionRulesByFilter(ctx context.Context, filter model.Filter, ledgerAccountResolutionRulesUpdateFields ...LedgerAccountResolutionRulesUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerAccountResolutionRules(ctx context.Context, ledgerAccountResolutionRulesListMap map[model.LedgerAccountResolutionRulesPrimaryID]*model.LedgerAccountResolutionRules, LedgerAccountResolutionRulessMapUpdateFieldsRequest map[model.LedgerAccountResolutionRulesPrimaryID]LedgerAccountResolutionRulesUpdateFieldList) (err error)
	DeleteLedgerAccountResolutionRulesByID(ctx context.Context, id model.LedgerAccountResolutionRulesPrimaryID) error
	BulkDeleteLedgerAccountResolutionRulesByIDs(ctx context.Context, ids []model.LedgerAccountResolutionRulesPrimaryID) error
	ResolveLedgerAccountResolutionRulesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountResolutionRulesFilterResult, err error)
	IsExistLedgerAccountResolutionRulesByIDs(ctx context.Context, ids []model.LedgerAccountResolutionRulesPrimaryID) (exists bool, notFoundIds []model.LedgerAccountResolutionRulesPrimaryID, err error)
	IsExistLedgerAccountResolutionRulesByID(ctx context.Context, id model.LedgerAccountResolutionRulesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
