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

func composeInsertFieldsAndParamsAccountingPeriods(accountingPeriodsList []model.AccountingPeriods, fieldsInsert ...AccountingPeriodsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewAccountingPeriodsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, accountingPeriods := range accountingPeriodsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, accountingPeriods.Id)
			case selectField.BookId():
				args = append(args, accountingPeriods.BookId)
			case selectField.PeriodCode():
				args = append(args, accountingPeriods.PeriodCode)
			case selectField.PeriodType():
				args = append(args, accountingPeriods.PeriodType)
			case selectField.PeriodStart():
				args = append(args, accountingPeriods.PeriodStart)
			case selectField.PeriodEnd():
				args = append(args, accountingPeriods.PeriodEnd)
			case selectField.PeriodStatus():
				args = append(args, accountingPeriods.PeriodStatus)
			case selectField.ClosedAt():
				args = append(args, accountingPeriods.ClosedAt)
			case selectField.ClosedBy():
				args = append(args, accountingPeriods.ClosedBy)
			case selectField.LockReason():
				args = append(args, accountingPeriods.LockReason)
			case selectField.ClosingHash():
				args = append(args, accountingPeriods.ClosingHash)
			case selectField.Metadata():
				args = append(args, accountingPeriods.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, accountingPeriods.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, accountingPeriods.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, accountingPeriods.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, accountingPeriods.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, accountingPeriods.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, accountingPeriods.MetaDeletedBy)

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

func composeAccountingPeriodsCompositePrimaryKeyWhere(primaryIDs []model.AccountingPeriodsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"accounting_periods\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultAccountingPeriodsSelectFields() string {
	fields := NewAccountingPeriodsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeAccountingPeriodsSelectFields(selectFields ...AccountingPeriodsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type AccountingPeriodsField string
type AccountingPeriodsFieldList []AccountingPeriodsField

type AccountingPeriodsSelectFields struct {
}

func (ss AccountingPeriodsSelectFields) Id() AccountingPeriodsField {
	return AccountingPeriodsField("id")
}

func (ss AccountingPeriodsSelectFields) BookId() AccountingPeriodsField {
	return AccountingPeriodsField("book_id")
}

func (ss AccountingPeriodsSelectFields) PeriodCode() AccountingPeriodsField {
	return AccountingPeriodsField("period_code")
}

func (ss AccountingPeriodsSelectFields) PeriodType() AccountingPeriodsField {
	return AccountingPeriodsField("period_type")
}

func (ss AccountingPeriodsSelectFields) PeriodStart() AccountingPeriodsField {
	return AccountingPeriodsField("period_start")
}

func (ss AccountingPeriodsSelectFields) PeriodEnd() AccountingPeriodsField {
	return AccountingPeriodsField("period_end")
}

func (ss AccountingPeriodsSelectFields) PeriodStatus() AccountingPeriodsField {
	return AccountingPeriodsField("period_status")
}

func (ss AccountingPeriodsSelectFields) ClosedAt() AccountingPeriodsField {
	return AccountingPeriodsField("closed_at")
}

func (ss AccountingPeriodsSelectFields) ClosedBy() AccountingPeriodsField {
	return AccountingPeriodsField("closed_by")
}

func (ss AccountingPeriodsSelectFields) LockReason() AccountingPeriodsField {
	return AccountingPeriodsField("lock_reason")
}

func (ss AccountingPeriodsSelectFields) ClosingHash() AccountingPeriodsField {
	return AccountingPeriodsField("closing_hash")
}

func (ss AccountingPeriodsSelectFields) Metadata() AccountingPeriodsField {
	return AccountingPeriodsField("metadata")
}

func (ss AccountingPeriodsSelectFields) MetaCreatedAt() AccountingPeriodsField {
	return AccountingPeriodsField("meta_created_at")
}

func (ss AccountingPeriodsSelectFields) MetaCreatedBy() AccountingPeriodsField {
	return AccountingPeriodsField("meta_created_by")
}

func (ss AccountingPeriodsSelectFields) MetaUpdatedAt() AccountingPeriodsField {
	return AccountingPeriodsField("meta_updated_at")
}

func (ss AccountingPeriodsSelectFields) MetaUpdatedBy() AccountingPeriodsField {
	return AccountingPeriodsField("meta_updated_by")
}

func (ss AccountingPeriodsSelectFields) MetaDeletedAt() AccountingPeriodsField {
	return AccountingPeriodsField("meta_deleted_at")
}

func (ss AccountingPeriodsSelectFields) MetaDeletedBy() AccountingPeriodsField {
	return AccountingPeriodsField("meta_deleted_by")
}

func (ss AccountingPeriodsSelectFields) All() AccountingPeriodsFieldList {
	return []AccountingPeriodsField{
		ss.Id(),
		ss.BookId(),
		ss.PeriodCode(),
		ss.PeriodType(),
		ss.PeriodStart(),
		ss.PeriodEnd(),
		ss.PeriodStatus(),
		ss.ClosedAt(),
		ss.ClosedBy(),
		ss.LockReason(),
		ss.ClosingHash(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewAccountingPeriodsSelectFields() AccountingPeriodsSelectFields {
	return AccountingPeriodsSelectFields{}
}

type AccountingPeriodsUpdateFieldOption struct {
	useIncrement bool
}
type AccountingPeriodsUpdateField struct {
	accountingPeriodsField AccountingPeriodsField
	opt                    AccountingPeriodsUpdateFieldOption
	value                  interface{}
}
type AccountingPeriodsUpdateFieldList []AccountingPeriodsUpdateField

func defaultAccountingPeriodsUpdateFieldOption() AccountingPeriodsUpdateFieldOption {
	return AccountingPeriodsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementAccountingPeriodsOption(useIncrement bool) func(*AccountingPeriodsUpdateFieldOption) {
	return func(pcufo *AccountingPeriodsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewAccountingPeriodsUpdateField(field AccountingPeriodsField, val interface{}, opts ...func(*AccountingPeriodsUpdateFieldOption)) AccountingPeriodsUpdateField {
	defaultOpt := defaultAccountingPeriodsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return AccountingPeriodsUpdateField{
		accountingPeriodsField: field,
		value:                  val,
		opt:                    defaultOpt,
	}
}
func defaultAccountingPeriodsUpdateFields(accountingPeriods model.AccountingPeriods) (accountingPeriodsUpdateFieldList AccountingPeriodsUpdateFieldList) {
	selectFields := NewAccountingPeriodsSelectFields()
	accountingPeriodsUpdateFieldList = append(accountingPeriodsUpdateFieldList,
		NewAccountingPeriodsUpdateField(selectFields.Id(), accountingPeriods.Id),
		NewAccountingPeriodsUpdateField(selectFields.BookId(), accountingPeriods.BookId),
		NewAccountingPeriodsUpdateField(selectFields.PeriodCode(), accountingPeriods.PeriodCode),
		NewAccountingPeriodsUpdateField(selectFields.PeriodType(), accountingPeriods.PeriodType),
		NewAccountingPeriodsUpdateField(selectFields.PeriodStart(), accountingPeriods.PeriodStart),
		NewAccountingPeriodsUpdateField(selectFields.PeriodEnd(), accountingPeriods.PeriodEnd),
		NewAccountingPeriodsUpdateField(selectFields.PeriodStatus(), accountingPeriods.PeriodStatus),
		NewAccountingPeriodsUpdateField(selectFields.ClosedAt(), accountingPeriods.ClosedAt),
		NewAccountingPeriodsUpdateField(selectFields.ClosedBy(), accountingPeriods.ClosedBy),
		NewAccountingPeriodsUpdateField(selectFields.LockReason(), accountingPeriods.LockReason),
		NewAccountingPeriodsUpdateField(selectFields.ClosingHash(), accountingPeriods.ClosingHash),
		NewAccountingPeriodsUpdateField(selectFields.Metadata(), accountingPeriods.Metadata),
		NewAccountingPeriodsUpdateField(selectFields.MetaCreatedAt(), accountingPeriods.MetaCreatedAt),
		NewAccountingPeriodsUpdateField(selectFields.MetaCreatedBy(), accountingPeriods.MetaCreatedBy),
		NewAccountingPeriodsUpdateField(selectFields.MetaUpdatedAt(), accountingPeriods.MetaUpdatedAt),
		NewAccountingPeriodsUpdateField(selectFields.MetaUpdatedBy(), accountingPeriods.MetaUpdatedBy),
		NewAccountingPeriodsUpdateField(selectFields.MetaDeletedAt(), accountingPeriods.MetaDeletedAt),
		NewAccountingPeriodsUpdateField(selectFields.MetaDeletedBy(), accountingPeriods.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsAccountingPeriodsCommand(accountingPeriodsUpdateFieldList AccountingPeriodsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range accountingPeriodsUpdateFieldList {
		field := string(updateField.accountingPeriodsField)
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

func (repo *RepositoryImpl) BulkCreateAccountingPeriods(ctx context.Context, accountingPeriodsList []*model.AccountingPeriods, fieldsInsert ...AccountingPeriodsField) (err error) {
	var (
		fieldsStr                  string
		valueListStr               []string
		argsList                   []interface{}
		primaryIds                 []model.AccountingPeriodsPrimaryID
		accountingPeriodsValueList []model.AccountingPeriods
	)

	if len(fieldsInsert) == 0 {
		selectField := NewAccountingPeriodsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, accountingPeriods := range accountingPeriodsList {

		primaryIds = append(primaryIds, accountingPeriods.ToAccountingPeriodsPrimaryID())

		accountingPeriodsValueList = append(accountingPeriodsValueList, *accountingPeriods)
	}

	_, notFoundIds, err := repo.IsExistAccountingPeriodsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateAccountingPeriods] failed checking accountingPeriods whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.AccountingPeriodsPrimaryID{}
		mapNotFoundIds := map[model.AccountingPeriodsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "accountingPeriods", fmt.Sprintf("accountingPeriods with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsAccountingPeriods(accountingPeriodsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(accountingPeriodsQueries.insertAccountingPeriods, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateAccountingPeriods] failed exec create accountingPeriods query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteAccountingPeriodsByIDs(ctx context.Context, primaryIDs []model.AccountingPeriodsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistAccountingPeriodsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteAccountingPeriodsByIDs] failed checking accountingPeriods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("accountingPeriods with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"accounting_periods\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(accountingPeriodsQueries.deleteAccountingPeriods + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteAccountingPeriodsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteAccountingPeriodsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistAccountingPeriodsByIDs(ctx context.Context, ids []model.AccountingPeriodsPrimaryID) (exists bool, notFoundIds []model.AccountingPeriodsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"accounting_periods\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(accountingPeriodsQueries.selectAccountingPeriods, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistAccountingPeriodsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.AccountingPeriodsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistAccountingPeriodsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.AccountingPeriodsPrimaryID]bool{}
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

// BulkUpdateAccountingPeriods is used to bulk update accountingPeriods, by default it will update all field
// if want to update specific field, then fill accountingPeriodssMapUpdateFieldsRequest else please fill accountingPeriodssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateAccountingPeriods(ctx context.Context, accountingPeriodssMap map[model.AccountingPeriodsPrimaryID]*model.AccountingPeriods, accountingPeriodssMapUpdateFieldsRequest map[model.AccountingPeriodsPrimaryID]AccountingPeriodsUpdateFieldList) (err error) {
	if len(accountingPeriodssMap) == 0 && len(accountingPeriodssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		accountingPeriodssMapUpdateField map[model.AccountingPeriodsPrimaryID]AccountingPeriodsUpdateFieldList = map[model.AccountingPeriodsPrimaryID]AccountingPeriodsUpdateFieldList{}
		asTableValues                    string                                                                = "myvalues"
	)

	if len(accountingPeriodssMap) > 0 {
		for id, accountingPeriods := range accountingPeriodssMap {
			if accountingPeriods == nil {
				log.Error().Err(err).Msg("[BulkUpdateAccountingPeriods] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			accountingPeriodssMapUpdateField[id] = defaultAccountingPeriodsUpdateFields(*accountingPeriods)
		}
	} else {
		accountingPeriodssMapUpdateField = accountingPeriodssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateAccountingPeriodsQuery(accountingPeriodssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistAccountingPeriodsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateAccountingPeriods] failed checking accountingPeriods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("accountingPeriods with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeAccountingPeriodsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"accounting_periods\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateAccountingPeriods] failed exec query")
	}
	return
}

type AccountingPeriodsFieldParameter struct {
	param string
	args  []interface{}
}

func NewAccountingPeriodsFieldParameter(param string, args ...interface{}) AccountingPeriodsFieldParameter {
	return AccountingPeriodsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateAccountingPeriodsQuery(mapAccountingPeriodss map[model.AccountingPeriodsPrimaryID]AccountingPeriodsUpdateFieldList, asTableValues string) (primaryIDs []model.AccountingPeriodsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.AccountingPeriodsPrimaryID]map[string]interface{}{}
	accountingPeriodsSelectFields := NewAccountingPeriodsSelectFields()
	for id, updateFields := range mapAccountingPeriodss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.accountingPeriodsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapAccountingPeriodss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetAccountingPeriodsFieldType(updateField.accountingPeriodsField)))
			args = append(args, fields[string(updateField.accountingPeriodsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.accountingPeriodsField))
		if updateField.accountingPeriodsField == accountingPeriodsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.accountingPeriodsField, asTableValues, updateField.accountingPeriodsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.accountingPeriodsField,
				"\"accounting_periods\"", updateField.accountingPeriodsField,
				asTableValues, updateField.accountingPeriodsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeAccountingPeriodsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.AccountingPeriodsPrimaryID, asTableValue string) (whereQry string) {
	accountingPeriodsSelectFields := NewAccountingPeriodsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"accounting_periods\".\"id\" = %s.\"id\"::"+GetAccountingPeriodsFieldType(accountingPeriodsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetAccountingPeriodsFieldType(accountingPeriodsField AccountingPeriodsField) string {
	selectAccountingPeriodsFields := NewAccountingPeriodsSelectFields()
	switch accountingPeriodsField {

	case selectAccountingPeriodsFields.Id():
		return "uuid"

	case selectAccountingPeriodsFields.BookId():
		return "uuid"

	case selectAccountingPeriodsFields.PeriodCode():
		return "text"

	case selectAccountingPeriodsFields.PeriodType():
		return "text"

	case selectAccountingPeriodsFields.PeriodStart():
		return "timestamptz"

	case selectAccountingPeriodsFields.PeriodEnd():
		return "timestamptz"

	case selectAccountingPeriodsFields.PeriodStatus():
		return "period_status_enum"

	case selectAccountingPeriodsFields.ClosedAt():
		return "timestamptz"

	case selectAccountingPeriodsFields.ClosedBy():
		return "uuid"

	case selectAccountingPeriodsFields.LockReason():
		return "text"

	case selectAccountingPeriodsFields.ClosingHash():
		return "text"

	case selectAccountingPeriodsFields.Metadata():
		return "jsonb"

	case selectAccountingPeriodsFields.MetaCreatedAt():
		return "timestamptz"

	case selectAccountingPeriodsFields.MetaCreatedBy():
		return "uuid"

	case selectAccountingPeriodsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectAccountingPeriodsFields.MetaUpdatedBy():
		return "uuid"

	case selectAccountingPeriodsFields.MetaDeletedAt():
		return "timestamptz"

	case selectAccountingPeriodsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateAccountingPeriods(ctx context.Context, accountingPeriods *model.AccountingPeriods, fieldsInsert ...AccountingPeriodsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewAccountingPeriodsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.AccountingPeriodsPrimaryID{
		Id: accountingPeriods.Id,
	}
	exists, err := repo.IsExistAccountingPeriodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateAccountingPeriods] failed checking accountingPeriods whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "accountingPeriods", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsAccountingPeriods([]model.AccountingPeriods{*accountingPeriods}, fieldsInsert...)
	commandQuery := fmt.Sprintf(accountingPeriodsQueries.insertAccountingPeriods, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateAccountingPeriods] failed exec create accountingPeriods query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteAccountingPeriodsByID(ctx context.Context, primaryID model.AccountingPeriodsPrimaryID) (err error) {
	exists, err := repo.IsExistAccountingPeriodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteAccountingPeriodsByID] failed checking accountingPeriods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("accountingPeriods with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeAccountingPeriodsCompositePrimaryKeyWhere([]model.AccountingPeriodsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(accountingPeriodsQueries.deleteAccountingPeriods + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteAccountingPeriodsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveAccountingPeriodsByFilter(ctx context.Context, filter model.Filter) (result []model.AccountingPeriodsFilterResult, err error) {
	query, args, err := composeAccountingPeriodsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveAccountingPeriodsByFilter] failed compose accountingPeriods filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveAccountingPeriodsByFilter] failed get accountingPeriods by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeAccountingPeriodsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.AccountingPeriodsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeAccountingPeriodsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeAccountingPeriodsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeAccountingPeriodsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewAccountingPeriodsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["book_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_id\"")
			selectedColumns["book_id"] = struct{}{}
		}
		if _, selected := selectedColumns["period_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_code\"")
			selectedColumns["period_code"] = struct{}{}
		}
		if _, selected := selectedColumns["period_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_type\"")
			selectedColumns["period_type"] = struct{}{}
		}
		if _, selected := selectedColumns["period_start"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_start\"")
			selectedColumns["period_start"] = struct{}{}
		}
		if _, selected := selectedColumns["period_end"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_end\"")
			selectedColumns["period_end"] = struct{}{}
		}
		if _, selected := selectedColumns["period_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_status\"")
			selectedColumns["period_status"] = struct{}{}
		}
		if _, selected := selectedColumns["closed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"closed_at\"")
			selectedColumns["closed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["closed_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"closed_by\"")
			selectedColumns["closed_by"] = struct{}{}
		}
		if _, selected := selectedColumns["lock_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"lock_reason\"")
			selectedColumns["lock_reason"] = struct{}{}
		}
		if _, selected := selectedColumns["closing_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"closing_hash\"")
			selectedColumns["closing_hash"] = struct{}{}
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

type accountingPeriodsFilterPlaceholder struct {
	index int
}

func (p *accountingPeriodsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeAccountingPeriodsFilterPredicate(filterField model.FilterField, placeholders *accountingPeriodsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewAccountingPeriodsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeAccountingPeriodsFilterSQLExpr(spec)
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

func composeAccountingPeriodsFilterGroup(group model.FilterGroup, placeholders *accountingPeriodsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeAccountingPeriodsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeAccountingPeriodsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeAccountingPeriodsFilterWhereQueries(filter model.Filter, placeholders *accountingPeriodsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeAccountingPeriodsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeAccountingPeriodsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeAccountingPeriodsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateAccountingPeriodsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeAccountingPeriodsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeAccountingPeriodsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := accountingPeriodsFilterPlaceholder{index: 1}
	whereQueries, err := composeAccountingPeriodsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewAccountingPeriodsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeAccountingPeriodsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeAccountingPeriodsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"accounting_periods\" base%s", strings.Join(selectColumns, ","), composeAccountingPeriodsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistAccountingPeriodsByID(ctx context.Context, primaryID model.AccountingPeriodsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeAccountingPeriodsCompositePrimaryKeyWhere([]model.AccountingPeriodsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", accountingPeriodsQueries.selectCountAccountingPeriods, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistAccountingPeriodsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveAccountingPeriods(ctx context.Context, selectFields ...AccountingPeriodsField) (accountingPeriodsList model.AccountingPeriodsList, err error) {
	var (
		defaultAccountingPeriodsSelectFields = defaultAccountingPeriodsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultAccountingPeriodsSelectFields = composeAccountingPeriodsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(accountingPeriodsQueries.selectAccountingPeriods, defaultAccountingPeriodsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &accountingPeriodsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveAccountingPeriods] failed get accountingPeriods list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveAccountingPeriodsByID(ctx context.Context, primaryID model.AccountingPeriodsPrimaryID, selectFields ...AccountingPeriodsField) (accountingPeriods model.AccountingPeriods, err error) {
	var (
		defaultAccountingPeriodsSelectFields = defaultAccountingPeriodsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultAccountingPeriodsSelectFields = composeAccountingPeriodsSelectFields(selectFields...)
	}
	whereQry, params := composeAccountingPeriodsCompositePrimaryKeyWhere([]model.AccountingPeriodsPrimaryID{primaryID})
	query := fmt.Sprintf(accountingPeriodsQueries.selectAccountingPeriods+" WHERE "+whereQry, defaultAccountingPeriodsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &accountingPeriods, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("accountingPeriods with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveAccountingPeriodsByID] failed get accountingPeriods")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateAccountingPeriodsByID(ctx context.Context, primaryID model.AccountingPeriodsPrimaryID, accountingPeriods *model.AccountingPeriods, accountingPeriodsUpdateFields ...AccountingPeriodsUpdateField) (err error) {
	exists, err := repo.IsExistAccountingPeriodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateAccountingPeriods] failed checking accountingPeriods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("accountingPeriods with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if accountingPeriods == nil {
		if len(accountingPeriodsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateAccountingPeriodsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		accountingPeriods = &model.AccountingPeriods{}
	}
	var (
		defaultAccountingPeriodsUpdateFields = defaultAccountingPeriodsUpdateFields(*accountingPeriods)
		tempUpdateField                      AccountingPeriodsUpdateFieldList
		selectFields                         = NewAccountingPeriodsSelectFields()
	)
	if len(accountingPeriodsUpdateFields) > 0 {
		for _, updateField := range accountingPeriodsUpdateFields {
			if updateField.accountingPeriodsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultAccountingPeriodsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeAccountingPeriodsCompositePrimaryKeyWhere([]model.AccountingPeriodsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsAccountingPeriodsCommand(defaultAccountingPeriodsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(accountingPeriodsQueries.updateAccountingPeriods+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateAccountingPeriods] error when try to update accountingPeriods by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateAccountingPeriodsByFilter(ctx context.Context, filter model.Filter, accountingPeriodsUpdateFields ...AccountingPeriodsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(accountingPeriodsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields AccountingPeriodsUpdateFieldList
		selectFields = NewAccountingPeriodsSelectFields()
	)
	for _, updateField := range accountingPeriodsUpdateFields {
		if updateField.accountingPeriodsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsAccountingPeriodsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := accountingPeriodsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeAccountingPeriodsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"accounting_periods\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateAccountingPeriodsByFilter] error when try to update accountingPeriods by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateAccountingPeriodsByFilter] failed get rows affected")
	}
	return
}

var (
	accountingPeriodsQueries = struct {
		selectAccountingPeriods      string
		selectCountAccountingPeriods string
		deleteAccountingPeriods      string
		updateAccountingPeriods      string
		insertAccountingPeriods      string
	}{
		selectAccountingPeriods:      "SELECT %s FROM \"accounting_periods\"",
		selectCountAccountingPeriods: "SELECT COUNT(\"id\") FROM \"accounting_periods\"",
		deleteAccountingPeriods:      "DELETE FROM \"accounting_periods\"",
		updateAccountingPeriods:      "UPDATE \"accounting_periods\" SET %s ",
		insertAccountingPeriods:      "INSERT INTO \"accounting_periods\" %s VALUES %s",
	}
)

type AccountingPeriodsRepository interface {
	CreateAccountingPeriods(ctx context.Context, accountingPeriods *model.AccountingPeriods, fieldsInsert ...AccountingPeriodsField) error
	BulkCreateAccountingPeriods(ctx context.Context, accountingPeriodsList []*model.AccountingPeriods, fieldsInsert ...AccountingPeriodsField) error
	ResolveAccountingPeriods(ctx context.Context, selectFields ...AccountingPeriodsField) (model.AccountingPeriodsList, error)
	ResolveAccountingPeriodsByID(ctx context.Context, primaryID model.AccountingPeriodsPrimaryID, selectFields ...AccountingPeriodsField) (model.AccountingPeriods, error)
	UpdateAccountingPeriodsByID(ctx context.Context, id model.AccountingPeriodsPrimaryID, accountingPeriods *model.AccountingPeriods, accountingPeriodsUpdateFields ...AccountingPeriodsUpdateField) error
	UpdateAccountingPeriodsByFilter(ctx context.Context, filter model.Filter, accountingPeriodsUpdateFields ...AccountingPeriodsUpdateField) (rowsAffected int64, err error)
	BulkUpdateAccountingPeriods(ctx context.Context, accountingPeriodsListMap map[model.AccountingPeriodsPrimaryID]*model.AccountingPeriods, AccountingPeriodssMapUpdateFieldsRequest map[model.AccountingPeriodsPrimaryID]AccountingPeriodsUpdateFieldList) (err error)
	DeleteAccountingPeriodsByID(ctx context.Context, id model.AccountingPeriodsPrimaryID) error
	BulkDeleteAccountingPeriodsByIDs(ctx context.Context, ids []model.AccountingPeriodsPrimaryID) error
	ResolveAccountingPeriodsByFilter(ctx context.Context, filter model.Filter) (result []model.AccountingPeriodsFilterResult, err error)
	IsExistAccountingPeriodsByIDs(ctx context.Context, ids []model.AccountingPeriodsPrimaryID) (exists bool, notFoundIds []model.AccountingPeriodsPrimaryID, err error)
	IsExistAccountingPeriodsByID(ctx context.Context, id model.AccountingPeriodsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
