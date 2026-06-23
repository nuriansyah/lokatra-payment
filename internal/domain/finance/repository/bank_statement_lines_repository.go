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

func composeInsertFieldsAndParamsBankStatementLines(bankStatementLinesList []model.BankStatementLines, fieldsInsert ...BankStatementLinesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewBankStatementLinesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, bankStatementLines := range bankStatementLinesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, bankStatementLines.Id)
			case selectField.StatementFileId():
				args = append(args, bankStatementLines.StatementFileId)
			case selectField.LineNo():
				args = append(args, bankStatementLines.LineNo)
			case selectField.BankRef():
				args = append(args, bankStatementLines.BankRef)
			case selectField.TransactionType():
				args = append(args, bankStatementLines.TransactionType)
			case selectField.BookingDate():
				args = append(args, bankStatementLines.BookingDate)
			case selectField.ValueDate():
				args = append(args, bankStatementLines.ValueDate)
			case selectField.CurrencyCode():
				args = append(args, bankStatementLines.CurrencyCode)
			case selectField.DebitAmount():
				args = append(args, bankStatementLines.DebitAmount)
			case selectField.CreditAmount():
				args = append(args, bankStatementLines.CreditAmount)
			case selectField.NetAmount():
				args = append(args, bankStatementLines.NetAmount)
			case selectField.RawLine():
				args = append(args, bankStatementLines.RawLine)
			case selectField.LineHash():
				args = append(args, bankStatementLines.LineHash)
			case selectField.Metadata():
				args = append(args, bankStatementLines.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, bankStatementLines.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, bankStatementLines.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, bankStatementLines.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, bankStatementLines.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, bankStatementLines.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, bankStatementLines.MetaDeletedBy)

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

func composeBankStatementLinesCompositePrimaryKeyWhere(primaryIDs []model.BankStatementLinesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"bank_statement_lines\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultBankStatementLinesSelectFields() string {
	fields := NewBankStatementLinesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeBankStatementLinesSelectFields(selectFields ...BankStatementLinesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type BankStatementLinesField string
type BankStatementLinesFieldList []BankStatementLinesField

type BankStatementLinesSelectFields struct {
}

func (ss BankStatementLinesSelectFields) Id() BankStatementLinesField {
	return BankStatementLinesField("id")
}

func (ss BankStatementLinesSelectFields) StatementFileId() BankStatementLinesField {
	return BankStatementLinesField("statement_file_id")
}

func (ss BankStatementLinesSelectFields) LineNo() BankStatementLinesField {
	return BankStatementLinesField("line_no")
}

func (ss BankStatementLinesSelectFields) BankRef() BankStatementLinesField {
	return BankStatementLinesField("bank_ref")
}

func (ss BankStatementLinesSelectFields) TransactionType() BankStatementLinesField {
	return BankStatementLinesField("transaction_type")
}

func (ss BankStatementLinesSelectFields) BookingDate() BankStatementLinesField {
	return BankStatementLinesField("booking_date")
}

func (ss BankStatementLinesSelectFields) ValueDate() BankStatementLinesField {
	return BankStatementLinesField("value_date")
}

func (ss BankStatementLinesSelectFields) CurrencyCode() BankStatementLinesField {
	return BankStatementLinesField("currency_code")
}

func (ss BankStatementLinesSelectFields) DebitAmount() BankStatementLinesField {
	return BankStatementLinesField("debit_amount")
}

func (ss BankStatementLinesSelectFields) CreditAmount() BankStatementLinesField {
	return BankStatementLinesField("credit_amount")
}

func (ss BankStatementLinesSelectFields) NetAmount() BankStatementLinesField {
	return BankStatementLinesField("net_amount")
}

func (ss BankStatementLinesSelectFields) RawLine() BankStatementLinesField {
	return BankStatementLinesField("raw_line")
}

func (ss BankStatementLinesSelectFields) LineHash() BankStatementLinesField {
	return BankStatementLinesField("line_hash")
}

func (ss BankStatementLinesSelectFields) Metadata() BankStatementLinesField {
	return BankStatementLinesField("metadata")
}

func (ss BankStatementLinesSelectFields) MetaCreatedAt() BankStatementLinesField {
	return BankStatementLinesField("meta_created_at")
}

func (ss BankStatementLinesSelectFields) MetaCreatedBy() BankStatementLinesField {
	return BankStatementLinesField("meta_created_by")
}

func (ss BankStatementLinesSelectFields) MetaUpdatedAt() BankStatementLinesField {
	return BankStatementLinesField("meta_updated_at")
}

func (ss BankStatementLinesSelectFields) MetaUpdatedBy() BankStatementLinesField {
	return BankStatementLinesField("meta_updated_by")
}

func (ss BankStatementLinesSelectFields) MetaDeletedAt() BankStatementLinesField {
	return BankStatementLinesField("meta_deleted_at")
}

func (ss BankStatementLinesSelectFields) MetaDeletedBy() BankStatementLinesField {
	return BankStatementLinesField("meta_deleted_by")
}

func (ss BankStatementLinesSelectFields) All() BankStatementLinesFieldList {
	return []BankStatementLinesField{
		ss.Id(),
		ss.StatementFileId(),
		ss.LineNo(),
		ss.BankRef(),
		ss.TransactionType(),
		ss.BookingDate(),
		ss.ValueDate(),
		ss.CurrencyCode(),
		ss.DebitAmount(),
		ss.CreditAmount(),
		ss.NetAmount(),
		ss.RawLine(),
		ss.LineHash(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewBankStatementLinesSelectFields() BankStatementLinesSelectFields {
	return BankStatementLinesSelectFields{}
}

type BankStatementLinesUpdateFieldOption struct {
	useIncrement bool
}
type BankStatementLinesUpdateField struct {
	bankStatementLinesField BankStatementLinesField
	opt                     BankStatementLinesUpdateFieldOption
	value                   interface{}
}
type BankStatementLinesUpdateFieldList []BankStatementLinesUpdateField

func defaultBankStatementLinesUpdateFieldOption() BankStatementLinesUpdateFieldOption {
	return BankStatementLinesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementBankStatementLinesOption(useIncrement bool) func(*BankStatementLinesUpdateFieldOption) {
	return func(pcufo *BankStatementLinesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewBankStatementLinesUpdateField(field BankStatementLinesField, val interface{}, opts ...func(*BankStatementLinesUpdateFieldOption)) BankStatementLinesUpdateField {
	defaultOpt := defaultBankStatementLinesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return BankStatementLinesUpdateField{
		bankStatementLinesField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultBankStatementLinesUpdateFields(bankStatementLines model.BankStatementLines) (bankStatementLinesUpdateFieldList BankStatementLinesUpdateFieldList) {
	selectFields := NewBankStatementLinesSelectFields()
	bankStatementLinesUpdateFieldList = append(bankStatementLinesUpdateFieldList,
		NewBankStatementLinesUpdateField(selectFields.Id(), bankStatementLines.Id),
		NewBankStatementLinesUpdateField(selectFields.StatementFileId(), bankStatementLines.StatementFileId),
		NewBankStatementLinesUpdateField(selectFields.LineNo(), bankStatementLines.LineNo),
		NewBankStatementLinesUpdateField(selectFields.BankRef(), bankStatementLines.BankRef),
		NewBankStatementLinesUpdateField(selectFields.TransactionType(), bankStatementLines.TransactionType),
		NewBankStatementLinesUpdateField(selectFields.BookingDate(), bankStatementLines.BookingDate),
		NewBankStatementLinesUpdateField(selectFields.ValueDate(), bankStatementLines.ValueDate),
		NewBankStatementLinesUpdateField(selectFields.CurrencyCode(), bankStatementLines.CurrencyCode),
		NewBankStatementLinesUpdateField(selectFields.DebitAmount(), bankStatementLines.DebitAmount),
		NewBankStatementLinesUpdateField(selectFields.CreditAmount(), bankStatementLines.CreditAmount),
		NewBankStatementLinesUpdateField(selectFields.NetAmount(), bankStatementLines.NetAmount),
		NewBankStatementLinesUpdateField(selectFields.RawLine(), bankStatementLines.RawLine),
		NewBankStatementLinesUpdateField(selectFields.LineHash(), bankStatementLines.LineHash),
		NewBankStatementLinesUpdateField(selectFields.Metadata(), bankStatementLines.Metadata),
		NewBankStatementLinesUpdateField(selectFields.MetaCreatedAt(), bankStatementLines.MetaCreatedAt),
		NewBankStatementLinesUpdateField(selectFields.MetaCreatedBy(), bankStatementLines.MetaCreatedBy),
		NewBankStatementLinesUpdateField(selectFields.MetaUpdatedAt(), bankStatementLines.MetaUpdatedAt),
		NewBankStatementLinesUpdateField(selectFields.MetaUpdatedBy(), bankStatementLines.MetaUpdatedBy),
		NewBankStatementLinesUpdateField(selectFields.MetaDeletedAt(), bankStatementLines.MetaDeletedAt),
		NewBankStatementLinesUpdateField(selectFields.MetaDeletedBy(), bankStatementLines.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsBankStatementLinesCommand(bankStatementLinesUpdateFieldList BankStatementLinesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range bankStatementLinesUpdateFieldList {
		field := string(updateField.bankStatementLinesField)
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

func (repo *RepositoryImpl) BulkCreateBankStatementLines(ctx context.Context, bankStatementLinesList []*model.BankStatementLines, fieldsInsert ...BankStatementLinesField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.BankStatementLinesPrimaryID
		bankStatementLinesValueList []model.BankStatementLines
	)

	if len(fieldsInsert) == 0 {
		selectField := NewBankStatementLinesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, bankStatementLines := range bankStatementLinesList {

		primaryIds = append(primaryIds, bankStatementLines.ToBankStatementLinesPrimaryID())

		bankStatementLinesValueList = append(bankStatementLinesValueList, *bankStatementLines)
	}

	_, notFoundIds, err := repo.IsExistBankStatementLinesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateBankStatementLines] failed checking bankStatementLines whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.BankStatementLinesPrimaryID{}
		mapNotFoundIds := map[model.BankStatementLinesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "bankStatementLines", fmt.Sprintf("bankStatementLines with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsBankStatementLines(bankStatementLinesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(bankStatementLinesQueries.insertBankStatementLines, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateBankStatementLines] failed exec create bankStatementLines query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteBankStatementLinesByIDs(ctx context.Context, primaryIDs []model.BankStatementLinesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistBankStatementLinesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteBankStatementLinesByIDs] failed checking bankStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementLines with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"bank_statement_lines\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(bankStatementLinesQueries.deleteBankStatementLines + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteBankStatementLinesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteBankStatementLinesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistBankStatementLinesByIDs(ctx context.Context, ids []model.BankStatementLinesPrimaryID) (exists bool, notFoundIds []model.BankStatementLinesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"bank_statement_lines\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(bankStatementLinesQueries.selectBankStatementLines, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistBankStatementLinesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.BankStatementLinesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistBankStatementLinesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.BankStatementLinesPrimaryID]bool{}
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

// BulkUpdateBankStatementLines is used to bulk update bankStatementLines, by default it will update all field
// if want to update specific field, then fill bankStatementLinessMapUpdateFieldsRequest else please fill bankStatementLinessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateBankStatementLines(ctx context.Context, bankStatementLinessMap map[model.BankStatementLinesPrimaryID]*model.BankStatementLines, bankStatementLinessMapUpdateFieldsRequest map[model.BankStatementLinesPrimaryID]BankStatementLinesUpdateFieldList) (err error) {
	if len(bankStatementLinessMap) == 0 && len(bankStatementLinessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		bankStatementLinessMapUpdateField map[model.BankStatementLinesPrimaryID]BankStatementLinesUpdateFieldList = map[model.BankStatementLinesPrimaryID]BankStatementLinesUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(bankStatementLinessMap) > 0 {
		for id, bankStatementLines := range bankStatementLinessMap {
			if bankStatementLines == nil {
				log.Error().Err(err).Msg("[BulkUpdateBankStatementLines] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			bankStatementLinessMapUpdateField[id] = defaultBankStatementLinesUpdateFields(*bankStatementLines)
		}
	} else {
		bankStatementLinessMapUpdateField = bankStatementLinessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateBankStatementLinesQuery(bankStatementLinessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistBankStatementLinesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateBankStatementLines] failed checking bankStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementLines with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeBankStatementLinesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"bank_statement_lines\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateBankStatementLines] failed exec query")
	}
	return
}

type BankStatementLinesFieldParameter struct {
	param string
	args  []interface{}
}

func NewBankStatementLinesFieldParameter(param string, args ...interface{}) BankStatementLinesFieldParameter {
	return BankStatementLinesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateBankStatementLinesQuery(mapBankStatementLiness map[model.BankStatementLinesPrimaryID]BankStatementLinesUpdateFieldList, asTableValues string) (primaryIDs []model.BankStatementLinesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.BankStatementLinesPrimaryID]map[string]interface{}{}
	bankStatementLinesSelectFields := NewBankStatementLinesSelectFields()
	for id, updateFields := range mapBankStatementLiness {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.bankStatementLinesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapBankStatementLiness[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetBankStatementLinesFieldType(updateField.bankStatementLinesField)))
			args = append(args, fields[string(updateField.bankStatementLinesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.bankStatementLinesField))
		if updateField.bankStatementLinesField == bankStatementLinesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.bankStatementLinesField, asTableValues, updateField.bankStatementLinesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.bankStatementLinesField,
				"\"bank_statement_lines\"", updateField.bankStatementLinesField,
				asTableValues, updateField.bankStatementLinesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeBankStatementLinesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.BankStatementLinesPrimaryID, asTableValue string) (whereQry string) {
	bankStatementLinesSelectFields := NewBankStatementLinesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"bank_statement_lines\".\"id\" = %s.\"id\"::"+GetBankStatementLinesFieldType(bankStatementLinesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetBankStatementLinesFieldType(bankStatementLinesField BankStatementLinesField) string {
	selectBankStatementLinesFields := NewBankStatementLinesSelectFields()
	switch bankStatementLinesField {

	case selectBankStatementLinesFields.Id():
		return "uuid"

	case selectBankStatementLinesFields.StatementFileId():
		return "uuid"

	case selectBankStatementLinesFields.LineNo():
		return "int4"

	case selectBankStatementLinesFields.BankRef():
		return "text"

	case selectBankStatementLinesFields.TransactionType():
		return "text"

	case selectBankStatementLinesFields.BookingDate():
		return "date"

	case selectBankStatementLinesFields.ValueDate():
		return "date"

	case selectBankStatementLinesFields.CurrencyCode():
		return "text"

	case selectBankStatementLinesFields.DebitAmount():
		return "numeric"

	case selectBankStatementLinesFields.CreditAmount():
		return "numeric"

	case selectBankStatementLinesFields.NetAmount():
		return "numeric"

	case selectBankStatementLinesFields.RawLine():
		return "jsonb"

	case selectBankStatementLinesFields.LineHash():
		return "text"

	case selectBankStatementLinesFields.Metadata():
		return "jsonb"

	case selectBankStatementLinesFields.MetaCreatedAt():
		return "timestamptz"

	case selectBankStatementLinesFields.MetaCreatedBy():
		return "uuid"

	case selectBankStatementLinesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectBankStatementLinesFields.MetaUpdatedBy():
		return "uuid"

	case selectBankStatementLinesFields.MetaDeletedAt():
		return "timestamptz"

	case selectBankStatementLinesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateBankStatementLines(ctx context.Context, bankStatementLines *model.BankStatementLines, fieldsInsert ...BankStatementLinesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewBankStatementLinesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.BankStatementLinesPrimaryID{
		Id: bankStatementLines.Id,
	}
	exists, err := repo.IsExistBankStatementLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateBankStatementLines] failed checking bankStatementLines whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "bankStatementLines", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsBankStatementLines([]model.BankStatementLines{*bankStatementLines}, fieldsInsert...)
	commandQuery := fmt.Sprintf(bankStatementLinesQueries.insertBankStatementLines, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateBankStatementLines] failed exec create bankStatementLines query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteBankStatementLinesByID(ctx context.Context, primaryID model.BankStatementLinesPrimaryID) (err error) {
	exists, err := repo.IsExistBankStatementLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteBankStatementLinesByID] failed checking bankStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementLines with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeBankStatementLinesCompositePrimaryKeyWhere([]model.BankStatementLinesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(bankStatementLinesQueries.deleteBankStatementLines + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteBankStatementLinesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveBankStatementLinesByFilter(ctx context.Context, filter model.Filter) (result []model.BankStatementLinesFilterResult, err error) {
	query, args, err := composeBankStatementLinesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveBankStatementLinesByFilter] failed compose bankStatementLines filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveBankStatementLinesByFilter] failed get bankStatementLines by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeBankStatementLinesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.BankStatementLinesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeBankStatementLinesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeBankStatementLinesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeBankStatementLinesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewBankStatementLinesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["statement_file_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"statement_file_id\"")
			selectedColumns["statement_file_id"] = struct{}{}
		}
		if _, selected := selectedColumns["line_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"line_no\"")
			selectedColumns["line_no"] = struct{}{}
		}
		if _, selected := selectedColumns["bank_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"bank_ref\"")
			selectedColumns["bank_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["transaction_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"transaction_type\"")
			selectedColumns["transaction_type"] = struct{}{}
		}
		if _, selected := selectedColumns["booking_date"]; !selected {
			selectColumns = append(selectColumns, "base.\"booking_date\"")
			selectedColumns["booking_date"] = struct{}{}
		}
		if _, selected := selectedColumns["value_date"]; !selected {
			selectColumns = append(selectColumns, "base.\"value_date\"")
			selectedColumns["value_date"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["debit_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"debit_amount\"")
			selectedColumns["debit_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["credit_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"credit_amount\"")
			selectedColumns["credit_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["net_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"net_amount\"")
			selectedColumns["net_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["raw_line"]; !selected {
			selectColumns = append(selectColumns, "base.\"raw_line\"")
			selectedColumns["raw_line"] = struct{}{}
		}
		if _, selected := selectedColumns["line_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"line_hash\"")
			selectedColumns["line_hash"] = struct{}{}
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

type bankStatementLinesFilterPlaceholder struct {
	index int
}

func (p *bankStatementLinesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeBankStatementLinesFilterPredicate(filterField model.FilterField, placeholders *bankStatementLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewBankStatementLinesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeBankStatementLinesFilterSQLExpr(spec)
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

func composeBankStatementLinesFilterGroup(group model.FilterGroup, placeholders *bankStatementLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeBankStatementLinesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeBankStatementLinesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeBankStatementLinesFilterWhereQueries(filter model.Filter, placeholders *bankStatementLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeBankStatementLinesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeBankStatementLinesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeBankStatementLinesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateBankStatementLinesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeBankStatementLinesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeBankStatementLinesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := bankStatementLinesFilterPlaceholder{index: 1}
	whereQueries, err := composeBankStatementLinesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewBankStatementLinesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeBankStatementLinesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeBankStatementLinesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"bank_statement_lines\" base%s", strings.Join(selectColumns, ","), composeBankStatementLinesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistBankStatementLinesByID(ctx context.Context, primaryID model.BankStatementLinesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeBankStatementLinesCompositePrimaryKeyWhere([]model.BankStatementLinesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", bankStatementLinesQueries.selectCountBankStatementLines, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistBankStatementLinesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveBankStatementLines(ctx context.Context, selectFields ...BankStatementLinesField) (bankStatementLinesList model.BankStatementLinesList, err error) {
	var (
		defaultBankStatementLinesSelectFields = defaultBankStatementLinesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultBankStatementLinesSelectFields = composeBankStatementLinesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(bankStatementLinesQueries.selectBankStatementLines, defaultBankStatementLinesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &bankStatementLinesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveBankStatementLines] failed get bankStatementLines list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveBankStatementLinesByID(ctx context.Context, primaryID model.BankStatementLinesPrimaryID, selectFields ...BankStatementLinesField) (bankStatementLines model.BankStatementLines, err error) {
	var (
		defaultBankStatementLinesSelectFields = defaultBankStatementLinesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultBankStatementLinesSelectFields = composeBankStatementLinesSelectFields(selectFields...)
	}
	whereQry, params := composeBankStatementLinesCompositePrimaryKeyWhere([]model.BankStatementLinesPrimaryID{primaryID})
	query := fmt.Sprintf(bankStatementLinesQueries.selectBankStatementLines+" WHERE "+whereQry, defaultBankStatementLinesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &bankStatementLines, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("bankStatementLines with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveBankStatementLinesByID] failed get bankStatementLines")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateBankStatementLinesByID(ctx context.Context, primaryID model.BankStatementLinesPrimaryID, bankStatementLines *model.BankStatementLines, bankStatementLinesUpdateFields ...BankStatementLinesUpdateField) (err error) {
	exists, err := repo.IsExistBankStatementLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementLines] failed checking bankStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementLines with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if bankStatementLines == nil {
		if len(bankStatementLinesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateBankStatementLinesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		bankStatementLines = &model.BankStatementLines{}
	}
	var (
		defaultBankStatementLinesUpdateFields = defaultBankStatementLinesUpdateFields(*bankStatementLines)
		tempUpdateField                       BankStatementLinesUpdateFieldList
		selectFields                          = NewBankStatementLinesSelectFields()
	)
	if len(bankStatementLinesUpdateFields) > 0 {
		for _, updateField := range bankStatementLinesUpdateFields {
			if updateField.bankStatementLinesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultBankStatementLinesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeBankStatementLinesCompositePrimaryKeyWhere([]model.BankStatementLinesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsBankStatementLinesCommand(defaultBankStatementLinesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(bankStatementLinesQueries.updateBankStatementLines+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementLines] error when try to update bankStatementLines by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateBankStatementLinesByFilter(ctx context.Context, filter model.Filter, bankStatementLinesUpdateFields ...BankStatementLinesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(bankStatementLinesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields BankStatementLinesUpdateFieldList
		selectFields = NewBankStatementLinesSelectFields()
	)
	for _, updateField := range bankStatementLinesUpdateFields {
		if updateField.bankStatementLinesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsBankStatementLinesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := bankStatementLinesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeBankStatementLinesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"bank_statement_lines\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementLinesByFilter] error when try to update bankStatementLines by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementLinesByFilter] failed get rows affected")
	}
	return
}

var (
	bankStatementLinesQueries = struct {
		selectBankStatementLines      string
		selectCountBankStatementLines string
		deleteBankStatementLines      string
		updateBankStatementLines      string
		insertBankStatementLines      string
	}{
		selectBankStatementLines:      "SELECT %s FROM \"bank_statement_lines\"",
		selectCountBankStatementLines: "SELECT COUNT(\"id\") FROM \"bank_statement_lines\"",
		deleteBankStatementLines:      "DELETE FROM \"bank_statement_lines\"",
		updateBankStatementLines:      "UPDATE \"bank_statement_lines\" SET %s ",
		insertBankStatementLines:      "INSERT INTO \"bank_statement_lines\" %s VALUES %s",
	}
)

type BankStatementLinesRepository interface {
	CreateBankStatementLines(ctx context.Context, bankStatementLines *model.BankStatementLines, fieldsInsert ...BankStatementLinesField) error
	BulkCreateBankStatementLines(ctx context.Context, bankStatementLinesList []*model.BankStatementLines, fieldsInsert ...BankStatementLinesField) error
	ResolveBankStatementLines(ctx context.Context, selectFields ...BankStatementLinesField) (model.BankStatementLinesList, error)
	ResolveBankStatementLinesByID(ctx context.Context, primaryID model.BankStatementLinesPrimaryID, selectFields ...BankStatementLinesField) (model.BankStatementLines, error)
	UpdateBankStatementLinesByID(ctx context.Context, id model.BankStatementLinesPrimaryID, bankStatementLines *model.BankStatementLines, bankStatementLinesUpdateFields ...BankStatementLinesUpdateField) error
	UpdateBankStatementLinesByFilter(ctx context.Context, filter model.Filter, bankStatementLinesUpdateFields ...BankStatementLinesUpdateField) (rowsAffected int64, err error)
	BulkUpdateBankStatementLines(ctx context.Context, bankStatementLinesListMap map[model.BankStatementLinesPrimaryID]*model.BankStatementLines, BankStatementLinessMapUpdateFieldsRequest map[model.BankStatementLinesPrimaryID]BankStatementLinesUpdateFieldList) (err error)
	DeleteBankStatementLinesByID(ctx context.Context, id model.BankStatementLinesPrimaryID) error
	BulkDeleteBankStatementLinesByIDs(ctx context.Context, ids []model.BankStatementLinesPrimaryID) error
	ResolveBankStatementLinesByFilter(ctx context.Context, filter model.Filter) (result []model.BankStatementLinesFilterResult, err error)
	IsExistBankStatementLinesByIDs(ctx context.Context, ids []model.BankStatementLinesPrimaryID) (exists bool, notFoundIds []model.BankStatementLinesPrimaryID, err error)
	IsExistBankStatementLinesByID(ctx context.Context, id model.BankStatementLinesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
