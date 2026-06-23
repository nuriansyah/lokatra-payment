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

func composeInsertFieldsAndParamsLedgerAccounts(ledgerAccountsList []model.LedgerAccounts, fieldsInsert ...LedgerAccountsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerAccountsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerAccounts := range ledgerAccountsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerAccounts.Id)
			case selectField.BookId():
				args = append(args, ledgerAccounts.BookId)
			case selectField.AccountCode():
				args = append(args, ledgerAccounts.AccountCode)
			case selectField.AccountName():
				args = append(args, ledgerAccounts.AccountName)
			case selectField.AccountTypeCode():
				args = append(args, ledgerAccounts.AccountTypeCode)
			case selectField.OwnerPartyId():
				args = append(args, ledgerAccounts.OwnerPartyId)
			case selectField.ParentAccountId():
				args = append(args, ledgerAccounts.ParentAccountId)
			case selectField.CurrencyCode():
				args = append(args, ledgerAccounts.CurrencyCode)
			case selectField.AllowManualPosting():
				args = append(args, ledgerAccounts.AllowManualPosting)
			case selectField.AccountStatus():
				args = append(args, ledgerAccounts.AccountStatus)
			case selectField.Metadata():
				args = append(args, ledgerAccounts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerAccounts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerAccounts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerAccounts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerAccounts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerAccounts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerAccounts.MetaDeletedBy)

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

func composeLedgerAccountsCompositePrimaryKeyWhere(primaryIDs []model.LedgerAccountsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_accounts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerAccountsSelectFields() string {
	fields := NewLedgerAccountsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerAccountsSelectFields(selectFields ...LedgerAccountsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerAccountsField string
type LedgerAccountsFieldList []LedgerAccountsField

type LedgerAccountsSelectFields struct {
}

func (ss LedgerAccountsSelectFields) Id() LedgerAccountsField {
	return LedgerAccountsField("id")
}

func (ss LedgerAccountsSelectFields) BookId() LedgerAccountsField {
	return LedgerAccountsField("book_id")
}

func (ss LedgerAccountsSelectFields) AccountCode() LedgerAccountsField {
	return LedgerAccountsField("account_code")
}

func (ss LedgerAccountsSelectFields) AccountName() LedgerAccountsField {
	return LedgerAccountsField("account_name")
}

func (ss LedgerAccountsSelectFields) AccountTypeCode() LedgerAccountsField {
	return LedgerAccountsField("account_type_code")
}

func (ss LedgerAccountsSelectFields) OwnerPartyId() LedgerAccountsField {
	return LedgerAccountsField("owner_party_id")
}

func (ss LedgerAccountsSelectFields) ParentAccountId() LedgerAccountsField {
	return LedgerAccountsField("parent_account_id")
}

func (ss LedgerAccountsSelectFields) CurrencyCode() LedgerAccountsField {
	return LedgerAccountsField("currency_code")
}

func (ss LedgerAccountsSelectFields) AllowManualPosting() LedgerAccountsField {
	return LedgerAccountsField("allow_manual_posting")
}

func (ss LedgerAccountsSelectFields) AccountStatus() LedgerAccountsField {
	return LedgerAccountsField("account_status")
}

func (ss LedgerAccountsSelectFields) Metadata() LedgerAccountsField {
	return LedgerAccountsField("metadata")
}

func (ss LedgerAccountsSelectFields) MetaCreatedAt() LedgerAccountsField {
	return LedgerAccountsField("meta_created_at")
}

func (ss LedgerAccountsSelectFields) MetaCreatedBy() LedgerAccountsField {
	return LedgerAccountsField("meta_created_by")
}

func (ss LedgerAccountsSelectFields) MetaUpdatedAt() LedgerAccountsField {
	return LedgerAccountsField("meta_updated_at")
}

func (ss LedgerAccountsSelectFields) MetaUpdatedBy() LedgerAccountsField {
	return LedgerAccountsField("meta_updated_by")
}

func (ss LedgerAccountsSelectFields) MetaDeletedAt() LedgerAccountsField {
	return LedgerAccountsField("meta_deleted_at")
}

func (ss LedgerAccountsSelectFields) MetaDeletedBy() LedgerAccountsField {
	return LedgerAccountsField("meta_deleted_by")
}

func (ss LedgerAccountsSelectFields) All() LedgerAccountsFieldList {
	return []LedgerAccountsField{
		ss.Id(),
		ss.BookId(),
		ss.AccountCode(),
		ss.AccountName(),
		ss.AccountTypeCode(),
		ss.OwnerPartyId(),
		ss.ParentAccountId(),
		ss.CurrencyCode(),
		ss.AllowManualPosting(),
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

func NewLedgerAccountsSelectFields() LedgerAccountsSelectFields {
	return LedgerAccountsSelectFields{}
}

type LedgerAccountsUpdateFieldOption struct {
	useIncrement bool
}
type LedgerAccountsUpdateField struct {
	ledgerAccountsField LedgerAccountsField
	opt                 LedgerAccountsUpdateFieldOption
	value               interface{}
}
type LedgerAccountsUpdateFieldList []LedgerAccountsUpdateField

func defaultLedgerAccountsUpdateFieldOption() LedgerAccountsUpdateFieldOption {
	return LedgerAccountsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerAccountsOption(useIncrement bool) func(*LedgerAccountsUpdateFieldOption) {
	return func(pcufo *LedgerAccountsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerAccountsUpdateField(field LedgerAccountsField, val interface{}, opts ...func(*LedgerAccountsUpdateFieldOption)) LedgerAccountsUpdateField {
	defaultOpt := defaultLedgerAccountsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerAccountsUpdateField{
		ledgerAccountsField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultLedgerAccountsUpdateFields(ledgerAccounts model.LedgerAccounts) (ledgerAccountsUpdateFieldList LedgerAccountsUpdateFieldList) {
	selectFields := NewLedgerAccountsSelectFields()
	ledgerAccountsUpdateFieldList = append(ledgerAccountsUpdateFieldList,
		NewLedgerAccountsUpdateField(selectFields.Id(), ledgerAccounts.Id),
		NewLedgerAccountsUpdateField(selectFields.BookId(), ledgerAccounts.BookId),
		NewLedgerAccountsUpdateField(selectFields.AccountCode(), ledgerAccounts.AccountCode),
		NewLedgerAccountsUpdateField(selectFields.AccountName(), ledgerAccounts.AccountName),
		NewLedgerAccountsUpdateField(selectFields.AccountTypeCode(), ledgerAccounts.AccountTypeCode),
		NewLedgerAccountsUpdateField(selectFields.OwnerPartyId(), ledgerAccounts.OwnerPartyId),
		NewLedgerAccountsUpdateField(selectFields.ParentAccountId(), ledgerAccounts.ParentAccountId),
		NewLedgerAccountsUpdateField(selectFields.CurrencyCode(), ledgerAccounts.CurrencyCode),
		NewLedgerAccountsUpdateField(selectFields.AllowManualPosting(), ledgerAccounts.AllowManualPosting),
		NewLedgerAccountsUpdateField(selectFields.AccountStatus(), ledgerAccounts.AccountStatus),
		NewLedgerAccountsUpdateField(selectFields.Metadata(), ledgerAccounts.Metadata),
		NewLedgerAccountsUpdateField(selectFields.MetaCreatedAt(), ledgerAccounts.MetaCreatedAt),
		NewLedgerAccountsUpdateField(selectFields.MetaCreatedBy(), ledgerAccounts.MetaCreatedBy),
		NewLedgerAccountsUpdateField(selectFields.MetaUpdatedAt(), ledgerAccounts.MetaUpdatedAt),
		NewLedgerAccountsUpdateField(selectFields.MetaUpdatedBy(), ledgerAccounts.MetaUpdatedBy),
		NewLedgerAccountsUpdateField(selectFields.MetaDeletedAt(), ledgerAccounts.MetaDeletedAt),
		NewLedgerAccountsUpdateField(selectFields.MetaDeletedBy(), ledgerAccounts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerAccountsCommand(ledgerAccountsUpdateFieldList LedgerAccountsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerAccountsUpdateFieldList {
		field := string(updateField.ledgerAccountsField)
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

func (repo *RepositoryImpl) BulkCreateLedgerAccounts(ctx context.Context, ledgerAccountsList []*model.LedgerAccounts, fieldsInsert ...LedgerAccountsField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.LedgerAccountsPrimaryID
		ledgerAccountsValueList []model.LedgerAccounts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerAccounts := range ledgerAccountsList {

		primaryIds = append(primaryIds, ledgerAccounts.ToLedgerAccountsPrimaryID())

		ledgerAccountsValueList = append(ledgerAccountsValueList, *ledgerAccounts)
	}

	_, notFoundIds, err := repo.IsExistLedgerAccountsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccounts] failed checking ledgerAccounts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerAccountsPrimaryID{}
		mapNotFoundIds := map[model.LedgerAccountsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerAccounts", fmt.Sprintf("ledgerAccounts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerAccounts(ledgerAccountsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerAccountsQueries.insertLedgerAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccounts] failed exec create ledgerAccounts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerAccountsByIDs(ctx context.Context, primaryIDs []model.LedgerAccountsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerAccountsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountsByIDs] failed checking ledgerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccounts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_accounts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerAccountsQueries.deleteLedgerAccounts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerAccountsByIDs(ctx context.Context, ids []model.LedgerAccountsPrimaryID) (exists bool, notFoundIds []model.LedgerAccountsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_accounts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerAccountsQueries.selectLedgerAccounts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerAccountsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerAccountsPrimaryID]bool{}
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

// BulkUpdateLedgerAccounts is used to bulk update ledgerAccounts, by default it will update all field
// if want to update specific field, then fill ledgerAccountssMapUpdateFieldsRequest else please fill ledgerAccountssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerAccounts(ctx context.Context, ledgerAccountssMap map[model.LedgerAccountsPrimaryID]*model.LedgerAccounts, ledgerAccountssMapUpdateFieldsRequest map[model.LedgerAccountsPrimaryID]LedgerAccountsUpdateFieldList) (err error) {
	if len(ledgerAccountssMap) == 0 && len(ledgerAccountssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerAccountssMapUpdateField map[model.LedgerAccountsPrimaryID]LedgerAccountsUpdateFieldList = map[model.LedgerAccountsPrimaryID]LedgerAccountsUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(ledgerAccountssMap) > 0 {
		for id, ledgerAccounts := range ledgerAccountssMap {
			if ledgerAccounts == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerAccounts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerAccountssMapUpdateField[id] = defaultLedgerAccountsUpdateFields(*ledgerAccounts)
		}
	} else {
		ledgerAccountssMapUpdateField = ledgerAccountssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerAccountsQuery(ledgerAccountssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerAccountsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccounts] failed checking ledgerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccounts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerAccountsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_accounts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccounts] failed exec query")
	}
	return
}

type LedgerAccountsFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerAccountsFieldParameter(param string, args ...interface{}) LedgerAccountsFieldParameter {
	return LedgerAccountsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerAccountsQuery(mapLedgerAccountss map[model.LedgerAccountsPrimaryID]LedgerAccountsUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerAccountsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerAccountsPrimaryID]map[string]interface{}{}
	ledgerAccountsSelectFields := NewLedgerAccountsSelectFields()
	for id, updateFields := range mapLedgerAccountss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerAccountsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerAccountss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerAccountsFieldType(updateField.ledgerAccountsField)))
			args = append(args, fields[string(updateField.ledgerAccountsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerAccountsField))
		if updateField.ledgerAccountsField == ledgerAccountsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerAccountsField, asTableValues, updateField.ledgerAccountsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerAccountsField,
				"\"ledger_accounts\"", updateField.ledgerAccountsField,
				asTableValues, updateField.ledgerAccountsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerAccountsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerAccountsPrimaryID, asTableValue string) (whereQry string) {
	ledgerAccountsSelectFields := NewLedgerAccountsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_accounts\".\"id\" = %s.\"id\"::"+GetLedgerAccountsFieldType(ledgerAccountsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerAccountsFieldType(ledgerAccountsField LedgerAccountsField) string {
	selectLedgerAccountsFields := NewLedgerAccountsSelectFields()
	switch ledgerAccountsField {

	case selectLedgerAccountsFields.Id():
		return "uuid"

	case selectLedgerAccountsFields.BookId():
		return "uuid"

	case selectLedgerAccountsFields.AccountCode():
		return "text"

	case selectLedgerAccountsFields.AccountName():
		return "text"

	case selectLedgerAccountsFields.AccountTypeCode():
		return "text"

	case selectLedgerAccountsFields.OwnerPartyId():
		return "uuid"

	case selectLedgerAccountsFields.ParentAccountId():
		return "uuid"

	case selectLedgerAccountsFields.CurrencyCode():
		return "text"

	case selectLedgerAccountsFields.AllowManualPosting():
		return "bool"

	case selectLedgerAccountsFields.AccountStatus():
		return "ledger_accounts_account_status_enum"

	case selectLedgerAccountsFields.Metadata():
		return "jsonb"

	case selectLedgerAccountsFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerAccountsFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerAccountsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerAccountsFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerAccountsFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerAccountsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerAccounts(ctx context.Context, ledgerAccounts *model.LedgerAccounts, fieldsInsert ...LedgerAccountsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerAccountsPrimaryID{
		Id: ledgerAccounts.Id,
	}
	exists, err := repo.IsExistLedgerAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccounts] failed checking ledgerAccounts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerAccounts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerAccounts([]model.LedgerAccounts{*ledgerAccounts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerAccountsQueries.insertLedgerAccounts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccounts] failed exec create ledgerAccounts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerAccountsByID(ctx context.Context, primaryID model.LedgerAccountsPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountsByID] failed checking ledgerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerAccountsCompositePrimaryKeyWhere([]model.LedgerAccountsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerAccountsQueries.deleteLedgerAccounts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountsFilterResult, err error) {
	query, args, err := composeLedgerAccountsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountsByFilter] failed compose ledgerAccounts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountsByFilter] failed get ledgerAccounts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerAccountsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerAccountsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerAccountsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerAccountsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerAccountsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 17 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerAccountsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 17+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["book_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_id\"")
			selectedColumns["book_id"] = struct{}{}
		}
		if _, selected := selectedColumns["account_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_code\"")
			selectedColumns["account_code"] = struct{}{}
		}
		if _, selected := selectedColumns["account_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_name\"")
			selectedColumns["account_name"] = struct{}{}
		}
		if _, selected := selectedColumns["account_type_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_type_code\"")
			selectedColumns["account_type_code"] = struct{}{}
		}
		if _, selected := selectedColumns["owner_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"owner_party_id\"")
			selectedColumns["owner_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["parent_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"parent_account_id\"")
			selectedColumns["parent_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["allow_manual_posting"]; !selected {
			selectColumns = append(selectColumns, "base.\"allow_manual_posting\"")
			selectedColumns["allow_manual_posting"] = struct{}{}
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

type ledgerAccountsFilterPlaceholder struct {
	index int
}

func (p *ledgerAccountsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerAccountsFilterPredicate(filterField model.FilterField, placeholders *ledgerAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerAccountsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerAccountsFilterSQLExpr(spec)
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

func composeLedgerAccountsFilterGroup(group model.FilterGroup, placeholders *ledgerAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerAccountsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerAccountsFilterWhereQueries(filter model.Filter, placeholders *ledgerAccountsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerAccountsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerAccountsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerAccountsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerAccountsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerAccountsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerAccountsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerAccountsFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerAccountsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerAccountsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerAccountsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_accounts\" base%s", strings.Join(selectColumns, ","), composeLedgerAccountsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerAccountsByID(ctx context.Context, primaryID model.LedgerAccountsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerAccountsCompositePrimaryKeyWhere([]model.LedgerAccountsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerAccountsQueries.selectCountLedgerAccounts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccounts(ctx context.Context, selectFields ...LedgerAccountsField) (ledgerAccountsList model.LedgerAccountsList, err error) {
	var (
		defaultLedgerAccountsSelectFields = defaultLedgerAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountsSelectFields = composeLedgerAccountsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerAccountsQueries.selectLedgerAccounts, defaultLedgerAccountsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerAccountsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccounts] failed get ledgerAccounts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountsByID(ctx context.Context, primaryID model.LedgerAccountsPrimaryID, selectFields ...LedgerAccountsField) (ledgerAccounts model.LedgerAccounts, err error) {
	var (
		defaultLedgerAccountsSelectFields = defaultLedgerAccountsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountsSelectFields = composeLedgerAccountsSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerAccountsCompositePrimaryKeyWhere([]model.LedgerAccountsPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerAccountsQueries.selectLedgerAccounts+" WHERE "+whereQry, defaultLedgerAccountsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerAccounts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerAccounts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerAccountsByID] failed get ledgerAccounts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerAccountsByID(ctx context.Context, primaryID model.LedgerAccountsPrimaryID, ledgerAccounts *model.LedgerAccounts, ledgerAccountsUpdateFields ...LedgerAccountsUpdateField) (err error) {
	exists, err := repo.IsExistLedgerAccountsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccounts] failed checking ledgerAccounts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccounts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerAccounts == nil {
		if len(ledgerAccountsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerAccountsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerAccounts = &model.LedgerAccounts{}
	}
	var (
		defaultLedgerAccountsUpdateFields = defaultLedgerAccountsUpdateFields(*ledgerAccounts)
		tempUpdateField                   LedgerAccountsUpdateFieldList
		selectFields                      = NewLedgerAccountsSelectFields()
	)
	if len(ledgerAccountsUpdateFields) > 0 {
		for _, updateField := range ledgerAccountsUpdateFields {
			if updateField.ledgerAccountsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerAccountsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerAccountsCompositePrimaryKeyWhere([]model.LedgerAccountsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerAccountsCommand(defaultLedgerAccountsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerAccountsQueries.updateLedgerAccounts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccounts] error when try to update ledgerAccounts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerAccountsByFilter(ctx context.Context, filter model.Filter, ledgerAccountsUpdateFields ...LedgerAccountsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerAccountsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerAccountsUpdateFieldList
		selectFields = NewLedgerAccountsSelectFields()
	)
	for _, updateField := range ledgerAccountsUpdateFields {
		if updateField.ledgerAccountsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerAccountsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerAccountsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerAccountsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_accounts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountsByFilter] error when try to update ledgerAccounts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountsByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerAccountsQueries = struct {
		selectLedgerAccounts      string
		selectCountLedgerAccounts string
		deleteLedgerAccounts      string
		updateLedgerAccounts      string
		insertLedgerAccounts      string
	}{
		selectLedgerAccounts:      "SELECT %s FROM \"ledger_accounts\"",
		selectCountLedgerAccounts: "SELECT COUNT(\"id\") FROM \"ledger_accounts\"",
		deleteLedgerAccounts:      "DELETE FROM \"ledger_accounts\"",
		updateLedgerAccounts:      "UPDATE \"ledger_accounts\" SET %s ",
		insertLedgerAccounts:      "INSERT INTO \"ledger_accounts\" %s VALUES %s",
	}
)

type LedgerAccountsRepository interface {
	CreateLedgerAccounts(ctx context.Context, ledgerAccounts *model.LedgerAccounts, fieldsInsert ...LedgerAccountsField) error
	BulkCreateLedgerAccounts(ctx context.Context, ledgerAccountsList []*model.LedgerAccounts, fieldsInsert ...LedgerAccountsField) error
	ResolveLedgerAccounts(ctx context.Context, selectFields ...LedgerAccountsField) (model.LedgerAccountsList, error)
	ResolveLedgerAccountsByID(ctx context.Context, primaryID model.LedgerAccountsPrimaryID, selectFields ...LedgerAccountsField) (model.LedgerAccounts, error)
	UpdateLedgerAccountsByID(ctx context.Context, id model.LedgerAccountsPrimaryID, ledgerAccounts *model.LedgerAccounts, ledgerAccountsUpdateFields ...LedgerAccountsUpdateField) error
	UpdateLedgerAccountsByFilter(ctx context.Context, filter model.Filter, ledgerAccountsUpdateFields ...LedgerAccountsUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerAccounts(ctx context.Context, ledgerAccountsListMap map[model.LedgerAccountsPrimaryID]*model.LedgerAccounts, LedgerAccountssMapUpdateFieldsRequest map[model.LedgerAccountsPrimaryID]LedgerAccountsUpdateFieldList) (err error)
	DeleteLedgerAccountsByID(ctx context.Context, id model.LedgerAccountsPrimaryID) error
	BulkDeleteLedgerAccountsByIDs(ctx context.Context, ids []model.LedgerAccountsPrimaryID) error
	ResolveLedgerAccountsByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountsFilterResult, err error)
	IsExistLedgerAccountsByIDs(ctx context.Context, ids []model.LedgerAccountsPrimaryID) (exists bool, notFoundIds []model.LedgerAccountsPrimaryID, err error)
	IsExistLedgerAccountsByID(ctx context.Context, id model.LedgerAccountsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
