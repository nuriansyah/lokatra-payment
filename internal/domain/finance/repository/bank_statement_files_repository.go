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

func composeInsertFieldsAndParamsBankStatementFiles(bankStatementFilesList []model.BankStatementFiles, fieldsInsert ...BankStatementFilesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewBankStatementFilesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, bankStatementFiles := range bankStatementFilesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, bankStatementFiles.Id)
			case selectField.BankCode():
				args = append(args, bankStatementFiles.BankCode)
			case selectField.AccountNumberMasked():
				args = append(args, bankStatementFiles.AccountNumberMasked)
			case selectField.StatementPeriodStart():
				args = append(args, bankStatementFiles.StatementPeriodStart)
			case selectField.StatementPeriodEnd():
				args = append(args, bankStatementFiles.StatementPeriodEnd)
			case selectField.CurrencyCode():
				args = append(args, bankStatementFiles.CurrencyCode)
			case selectField.StorageUri():
				args = append(args, bankStatementFiles.StorageUri)
			case selectField.FileHash():
				args = append(args, bankStatementFiles.FileHash)
			case selectField.ImportStatus():
				args = append(args, bankStatementFiles.ImportStatus)
			case selectField.ImportedAt():
				args = append(args, bankStatementFiles.ImportedAt)
			case selectField.RowCount():
				args = append(args, bankStatementFiles.RowCount)
			case selectField.Metadata():
				args = append(args, bankStatementFiles.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, bankStatementFiles.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, bankStatementFiles.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, bankStatementFiles.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, bankStatementFiles.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, bankStatementFiles.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, bankStatementFiles.MetaDeletedBy)

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

func composeBankStatementFilesCompositePrimaryKeyWhere(primaryIDs []model.BankStatementFilesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"bank_statement_files\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultBankStatementFilesSelectFields() string {
	fields := NewBankStatementFilesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeBankStatementFilesSelectFields(selectFields ...BankStatementFilesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type BankStatementFilesField string
type BankStatementFilesFieldList []BankStatementFilesField

type BankStatementFilesSelectFields struct {
}

func (ss BankStatementFilesSelectFields) Id() BankStatementFilesField {
	return BankStatementFilesField("id")
}

func (ss BankStatementFilesSelectFields) BankCode() BankStatementFilesField {
	return BankStatementFilesField("bank_code")
}

func (ss BankStatementFilesSelectFields) AccountNumberMasked() BankStatementFilesField {
	return BankStatementFilesField("account_number_masked")
}

func (ss BankStatementFilesSelectFields) StatementPeriodStart() BankStatementFilesField {
	return BankStatementFilesField("statement_period_start")
}

func (ss BankStatementFilesSelectFields) StatementPeriodEnd() BankStatementFilesField {
	return BankStatementFilesField("statement_period_end")
}

func (ss BankStatementFilesSelectFields) CurrencyCode() BankStatementFilesField {
	return BankStatementFilesField("currency_code")
}

func (ss BankStatementFilesSelectFields) StorageUri() BankStatementFilesField {
	return BankStatementFilesField("storage_uri")
}

func (ss BankStatementFilesSelectFields) FileHash() BankStatementFilesField {
	return BankStatementFilesField("file_hash")
}

func (ss BankStatementFilesSelectFields) ImportStatus() BankStatementFilesField {
	return BankStatementFilesField("import_status")
}

func (ss BankStatementFilesSelectFields) ImportedAt() BankStatementFilesField {
	return BankStatementFilesField("imported_at")
}

func (ss BankStatementFilesSelectFields) RowCount() BankStatementFilesField {
	return BankStatementFilesField("row_count")
}

func (ss BankStatementFilesSelectFields) Metadata() BankStatementFilesField {
	return BankStatementFilesField("metadata")
}

func (ss BankStatementFilesSelectFields) MetaCreatedAt() BankStatementFilesField {
	return BankStatementFilesField("meta_created_at")
}

func (ss BankStatementFilesSelectFields) MetaCreatedBy() BankStatementFilesField {
	return BankStatementFilesField("meta_created_by")
}

func (ss BankStatementFilesSelectFields) MetaUpdatedAt() BankStatementFilesField {
	return BankStatementFilesField("meta_updated_at")
}

func (ss BankStatementFilesSelectFields) MetaUpdatedBy() BankStatementFilesField {
	return BankStatementFilesField("meta_updated_by")
}

func (ss BankStatementFilesSelectFields) MetaDeletedAt() BankStatementFilesField {
	return BankStatementFilesField("meta_deleted_at")
}

func (ss BankStatementFilesSelectFields) MetaDeletedBy() BankStatementFilesField {
	return BankStatementFilesField("meta_deleted_by")
}

func (ss BankStatementFilesSelectFields) All() BankStatementFilesFieldList {
	return []BankStatementFilesField{
		ss.Id(),
		ss.BankCode(),
		ss.AccountNumberMasked(),
		ss.StatementPeriodStart(),
		ss.StatementPeriodEnd(),
		ss.CurrencyCode(),
		ss.StorageUri(),
		ss.FileHash(),
		ss.ImportStatus(),
		ss.ImportedAt(),
		ss.RowCount(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewBankStatementFilesSelectFields() BankStatementFilesSelectFields {
	return BankStatementFilesSelectFields{}
}

type BankStatementFilesUpdateFieldOption struct {
	useIncrement bool
}
type BankStatementFilesUpdateField struct {
	bankStatementFilesField BankStatementFilesField
	opt                     BankStatementFilesUpdateFieldOption
	value                   interface{}
}
type BankStatementFilesUpdateFieldList []BankStatementFilesUpdateField

func defaultBankStatementFilesUpdateFieldOption() BankStatementFilesUpdateFieldOption {
	return BankStatementFilesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementBankStatementFilesOption(useIncrement bool) func(*BankStatementFilesUpdateFieldOption) {
	return func(pcufo *BankStatementFilesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewBankStatementFilesUpdateField(field BankStatementFilesField, val interface{}, opts ...func(*BankStatementFilesUpdateFieldOption)) BankStatementFilesUpdateField {
	defaultOpt := defaultBankStatementFilesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return BankStatementFilesUpdateField{
		bankStatementFilesField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultBankStatementFilesUpdateFields(bankStatementFiles model.BankStatementFiles) (bankStatementFilesUpdateFieldList BankStatementFilesUpdateFieldList) {
	selectFields := NewBankStatementFilesSelectFields()
	bankStatementFilesUpdateFieldList = append(bankStatementFilesUpdateFieldList,
		NewBankStatementFilesUpdateField(selectFields.Id(), bankStatementFiles.Id),
		NewBankStatementFilesUpdateField(selectFields.BankCode(), bankStatementFiles.BankCode),
		NewBankStatementFilesUpdateField(selectFields.AccountNumberMasked(), bankStatementFiles.AccountNumberMasked),
		NewBankStatementFilesUpdateField(selectFields.StatementPeriodStart(), bankStatementFiles.StatementPeriodStart),
		NewBankStatementFilesUpdateField(selectFields.StatementPeriodEnd(), bankStatementFiles.StatementPeriodEnd),
		NewBankStatementFilesUpdateField(selectFields.CurrencyCode(), bankStatementFiles.CurrencyCode),
		NewBankStatementFilesUpdateField(selectFields.StorageUri(), bankStatementFiles.StorageUri),
		NewBankStatementFilesUpdateField(selectFields.FileHash(), bankStatementFiles.FileHash),
		NewBankStatementFilesUpdateField(selectFields.ImportStatus(), bankStatementFiles.ImportStatus),
		NewBankStatementFilesUpdateField(selectFields.ImportedAt(), bankStatementFiles.ImportedAt),
		NewBankStatementFilesUpdateField(selectFields.RowCount(), bankStatementFiles.RowCount),
		NewBankStatementFilesUpdateField(selectFields.Metadata(), bankStatementFiles.Metadata),
		NewBankStatementFilesUpdateField(selectFields.MetaCreatedAt(), bankStatementFiles.MetaCreatedAt),
		NewBankStatementFilesUpdateField(selectFields.MetaCreatedBy(), bankStatementFiles.MetaCreatedBy),
		NewBankStatementFilesUpdateField(selectFields.MetaUpdatedAt(), bankStatementFiles.MetaUpdatedAt),
		NewBankStatementFilesUpdateField(selectFields.MetaUpdatedBy(), bankStatementFiles.MetaUpdatedBy),
		NewBankStatementFilesUpdateField(selectFields.MetaDeletedAt(), bankStatementFiles.MetaDeletedAt),
		NewBankStatementFilesUpdateField(selectFields.MetaDeletedBy(), bankStatementFiles.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsBankStatementFilesCommand(bankStatementFilesUpdateFieldList BankStatementFilesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range bankStatementFilesUpdateFieldList {
		field := string(updateField.bankStatementFilesField)
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

func (repo *RepositoryImpl) BulkCreateBankStatementFiles(ctx context.Context, bankStatementFilesList []*model.BankStatementFiles, fieldsInsert ...BankStatementFilesField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.BankStatementFilesPrimaryID
		bankStatementFilesValueList []model.BankStatementFiles
	)

	if len(fieldsInsert) == 0 {
		selectField := NewBankStatementFilesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, bankStatementFiles := range bankStatementFilesList {

		primaryIds = append(primaryIds, bankStatementFiles.ToBankStatementFilesPrimaryID())

		bankStatementFilesValueList = append(bankStatementFilesValueList, *bankStatementFiles)
	}

	_, notFoundIds, err := repo.IsExistBankStatementFilesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateBankStatementFiles] failed checking bankStatementFiles whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.BankStatementFilesPrimaryID{}
		mapNotFoundIds := map[model.BankStatementFilesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "bankStatementFiles", fmt.Sprintf("bankStatementFiles with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsBankStatementFiles(bankStatementFilesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(bankStatementFilesQueries.insertBankStatementFiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateBankStatementFiles] failed exec create bankStatementFiles query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteBankStatementFilesByIDs(ctx context.Context, primaryIDs []model.BankStatementFilesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistBankStatementFilesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteBankStatementFilesByIDs] failed checking bankStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementFiles with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"bank_statement_files\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(bankStatementFilesQueries.deleteBankStatementFiles + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteBankStatementFilesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteBankStatementFilesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistBankStatementFilesByIDs(ctx context.Context, ids []model.BankStatementFilesPrimaryID) (exists bool, notFoundIds []model.BankStatementFilesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"bank_statement_files\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(bankStatementFilesQueries.selectBankStatementFiles, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistBankStatementFilesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.BankStatementFilesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistBankStatementFilesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.BankStatementFilesPrimaryID]bool{}
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

// BulkUpdateBankStatementFiles is used to bulk update bankStatementFiles, by default it will update all field
// if want to update specific field, then fill bankStatementFilessMapUpdateFieldsRequest else please fill bankStatementFilessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateBankStatementFiles(ctx context.Context, bankStatementFilessMap map[model.BankStatementFilesPrimaryID]*model.BankStatementFiles, bankStatementFilessMapUpdateFieldsRequest map[model.BankStatementFilesPrimaryID]BankStatementFilesUpdateFieldList) (err error) {
	if len(bankStatementFilessMap) == 0 && len(bankStatementFilessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		bankStatementFilessMapUpdateField map[model.BankStatementFilesPrimaryID]BankStatementFilesUpdateFieldList = map[model.BankStatementFilesPrimaryID]BankStatementFilesUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(bankStatementFilessMap) > 0 {
		for id, bankStatementFiles := range bankStatementFilessMap {
			if bankStatementFiles == nil {
				log.Error().Err(err).Msg("[BulkUpdateBankStatementFiles] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			bankStatementFilessMapUpdateField[id] = defaultBankStatementFilesUpdateFields(*bankStatementFiles)
		}
	} else {
		bankStatementFilessMapUpdateField = bankStatementFilessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateBankStatementFilesQuery(bankStatementFilessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistBankStatementFilesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateBankStatementFiles] failed checking bankStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementFiles with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeBankStatementFilesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"bank_statement_files\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateBankStatementFiles] failed exec query")
	}
	return
}

type BankStatementFilesFieldParameter struct {
	param string
	args  []interface{}
}

func NewBankStatementFilesFieldParameter(param string, args ...interface{}) BankStatementFilesFieldParameter {
	return BankStatementFilesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateBankStatementFilesQuery(mapBankStatementFiless map[model.BankStatementFilesPrimaryID]BankStatementFilesUpdateFieldList, asTableValues string) (primaryIDs []model.BankStatementFilesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.BankStatementFilesPrimaryID]map[string]interface{}{}
	bankStatementFilesSelectFields := NewBankStatementFilesSelectFields()
	for id, updateFields := range mapBankStatementFiless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.bankStatementFilesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapBankStatementFiless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetBankStatementFilesFieldType(updateField.bankStatementFilesField)))
			args = append(args, fields[string(updateField.bankStatementFilesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.bankStatementFilesField))
		if updateField.bankStatementFilesField == bankStatementFilesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.bankStatementFilesField, asTableValues, updateField.bankStatementFilesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.bankStatementFilesField,
				"\"bank_statement_files\"", updateField.bankStatementFilesField,
				asTableValues, updateField.bankStatementFilesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeBankStatementFilesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.BankStatementFilesPrimaryID, asTableValue string) (whereQry string) {
	bankStatementFilesSelectFields := NewBankStatementFilesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"bank_statement_files\".\"id\" = %s.\"id\"::"+GetBankStatementFilesFieldType(bankStatementFilesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetBankStatementFilesFieldType(bankStatementFilesField BankStatementFilesField) string {
	selectBankStatementFilesFields := NewBankStatementFilesSelectFields()
	switch bankStatementFilesField {

	case selectBankStatementFilesFields.Id():
		return "uuid"

	case selectBankStatementFilesFields.BankCode():
		return "text"

	case selectBankStatementFilesFields.AccountNumberMasked():
		return "text"

	case selectBankStatementFilesFields.StatementPeriodStart():
		return "date"

	case selectBankStatementFilesFields.StatementPeriodEnd():
		return "date"

	case selectBankStatementFilesFields.CurrencyCode():
		return "text"

	case selectBankStatementFilesFields.StorageUri():
		return "text"

	case selectBankStatementFilesFields.FileHash():
		return "text"

	case selectBankStatementFilesFields.ImportStatus():
		return "import_status_enum"

	case selectBankStatementFilesFields.ImportedAt():
		return "timestamptz"

	case selectBankStatementFilesFields.RowCount():
		return "int4"

	case selectBankStatementFilesFields.Metadata():
		return "jsonb"

	case selectBankStatementFilesFields.MetaCreatedAt():
		return "timestamptz"

	case selectBankStatementFilesFields.MetaCreatedBy():
		return "uuid"

	case selectBankStatementFilesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectBankStatementFilesFields.MetaUpdatedBy():
		return "uuid"

	case selectBankStatementFilesFields.MetaDeletedAt():
		return "timestamptz"

	case selectBankStatementFilesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateBankStatementFiles(ctx context.Context, bankStatementFiles *model.BankStatementFiles, fieldsInsert ...BankStatementFilesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewBankStatementFilesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.BankStatementFilesPrimaryID{
		Id: bankStatementFiles.Id,
	}
	exists, err := repo.IsExistBankStatementFilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateBankStatementFiles] failed checking bankStatementFiles whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "bankStatementFiles", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsBankStatementFiles([]model.BankStatementFiles{*bankStatementFiles}, fieldsInsert...)
	commandQuery := fmt.Sprintf(bankStatementFilesQueries.insertBankStatementFiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateBankStatementFiles] failed exec create bankStatementFiles query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteBankStatementFilesByID(ctx context.Context, primaryID model.BankStatementFilesPrimaryID) (err error) {
	exists, err := repo.IsExistBankStatementFilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteBankStatementFilesByID] failed checking bankStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementFiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeBankStatementFilesCompositePrimaryKeyWhere([]model.BankStatementFilesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(bankStatementFilesQueries.deleteBankStatementFiles + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteBankStatementFilesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveBankStatementFilesByFilter(ctx context.Context, filter model.Filter) (result []model.BankStatementFilesFilterResult, err error) {
	query, args, err := composeBankStatementFilesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveBankStatementFilesByFilter] failed compose bankStatementFiles filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveBankStatementFilesByFilter] failed get bankStatementFiles by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeBankStatementFilesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.BankStatementFilesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeBankStatementFilesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeBankStatementFilesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeBankStatementFilesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewBankStatementFilesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["bank_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"bank_code\"")
			selectedColumns["bank_code"] = struct{}{}
		}
		if _, selected := selectedColumns["account_number_masked"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_number_masked\"")
			selectedColumns["account_number_masked"] = struct{}{}
		}
		if _, selected := selectedColumns["statement_period_start"]; !selected {
			selectColumns = append(selectColumns, "base.\"statement_period_start\"")
			selectedColumns["statement_period_start"] = struct{}{}
		}
		if _, selected := selectedColumns["statement_period_end"]; !selected {
			selectColumns = append(selectColumns, "base.\"statement_period_end\"")
			selectedColumns["statement_period_end"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["storage_uri"]; !selected {
			selectColumns = append(selectColumns, "base.\"storage_uri\"")
			selectedColumns["storage_uri"] = struct{}{}
		}
		if _, selected := selectedColumns["file_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"file_hash\"")
			selectedColumns["file_hash"] = struct{}{}
		}
		if _, selected := selectedColumns["import_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"import_status\"")
			selectedColumns["import_status"] = struct{}{}
		}
		if _, selected := selectedColumns["imported_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"imported_at\"")
			selectedColumns["imported_at"] = struct{}{}
		}
		if _, selected := selectedColumns["row_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"row_count\"")
			selectedColumns["row_count"] = struct{}{}
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

type bankStatementFilesFilterPlaceholder struct {
	index int
}

func (p *bankStatementFilesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeBankStatementFilesFilterPredicate(filterField model.FilterField, placeholders *bankStatementFilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewBankStatementFilesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeBankStatementFilesFilterSQLExpr(spec)
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

func composeBankStatementFilesFilterGroup(group model.FilterGroup, placeholders *bankStatementFilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeBankStatementFilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeBankStatementFilesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeBankStatementFilesFilterWhereQueries(filter model.Filter, placeholders *bankStatementFilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeBankStatementFilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeBankStatementFilesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeBankStatementFilesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateBankStatementFilesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeBankStatementFilesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeBankStatementFilesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := bankStatementFilesFilterPlaceholder{index: 1}
	whereQueries, err := composeBankStatementFilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewBankStatementFilesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeBankStatementFilesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeBankStatementFilesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"bank_statement_files\" base%s", strings.Join(selectColumns, ","), composeBankStatementFilesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistBankStatementFilesByID(ctx context.Context, primaryID model.BankStatementFilesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeBankStatementFilesCompositePrimaryKeyWhere([]model.BankStatementFilesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", bankStatementFilesQueries.selectCountBankStatementFiles, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistBankStatementFilesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveBankStatementFiles(ctx context.Context, selectFields ...BankStatementFilesField) (bankStatementFilesList model.BankStatementFilesList, err error) {
	var (
		defaultBankStatementFilesSelectFields = defaultBankStatementFilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultBankStatementFilesSelectFields = composeBankStatementFilesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(bankStatementFilesQueries.selectBankStatementFiles, defaultBankStatementFilesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &bankStatementFilesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveBankStatementFiles] failed get bankStatementFiles list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveBankStatementFilesByID(ctx context.Context, primaryID model.BankStatementFilesPrimaryID, selectFields ...BankStatementFilesField) (bankStatementFiles model.BankStatementFiles, err error) {
	var (
		defaultBankStatementFilesSelectFields = defaultBankStatementFilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultBankStatementFilesSelectFields = composeBankStatementFilesSelectFields(selectFields...)
	}
	whereQry, params := composeBankStatementFilesCompositePrimaryKeyWhere([]model.BankStatementFilesPrimaryID{primaryID})
	query := fmt.Sprintf(bankStatementFilesQueries.selectBankStatementFiles+" WHERE "+whereQry, defaultBankStatementFilesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &bankStatementFiles, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("bankStatementFiles with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveBankStatementFilesByID] failed get bankStatementFiles")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateBankStatementFilesByID(ctx context.Context, primaryID model.BankStatementFilesPrimaryID, bankStatementFiles *model.BankStatementFiles, bankStatementFilesUpdateFields ...BankStatementFilesUpdateField) (err error) {
	exists, err := repo.IsExistBankStatementFilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementFiles] failed checking bankStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("bankStatementFiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if bankStatementFiles == nil {
		if len(bankStatementFilesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateBankStatementFilesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		bankStatementFiles = &model.BankStatementFiles{}
	}
	var (
		defaultBankStatementFilesUpdateFields = defaultBankStatementFilesUpdateFields(*bankStatementFiles)
		tempUpdateField                       BankStatementFilesUpdateFieldList
		selectFields                          = NewBankStatementFilesSelectFields()
	)
	if len(bankStatementFilesUpdateFields) > 0 {
		for _, updateField := range bankStatementFilesUpdateFields {
			if updateField.bankStatementFilesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultBankStatementFilesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeBankStatementFilesCompositePrimaryKeyWhere([]model.BankStatementFilesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsBankStatementFilesCommand(defaultBankStatementFilesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(bankStatementFilesQueries.updateBankStatementFiles+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementFiles] error when try to update bankStatementFiles by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateBankStatementFilesByFilter(ctx context.Context, filter model.Filter, bankStatementFilesUpdateFields ...BankStatementFilesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(bankStatementFilesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields BankStatementFilesUpdateFieldList
		selectFields = NewBankStatementFilesSelectFields()
	)
	for _, updateField := range bankStatementFilesUpdateFields {
		if updateField.bankStatementFilesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsBankStatementFilesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := bankStatementFilesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeBankStatementFilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"bank_statement_files\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementFilesByFilter] error when try to update bankStatementFiles by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBankStatementFilesByFilter] failed get rows affected")
	}
	return
}

var (
	bankStatementFilesQueries = struct {
		selectBankStatementFiles      string
		selectCountBankStatementFiles string
		deleteBankStatementFiles      string
		updateBankStatementFiles      string
		insertBankStatementFiles      string
	}{
		selectBankStatementFiles:      "SELECT %s FROM \"bank_statement_files\"",
		selectCountBankStatementFiles: "SELECT COUNT(\"id\") FROM \"bank_statement_files\"",
		deleteBankStatementFiles:      "DELETE FROM \"bank_statement_files\"",
		updateBankStatementFiles:      "UPDATE \"bank_statement_files\" SET %s ",
		insertBankStatementFiles:      "INSERT INTO \"bank_statement_files\" %s VALUES %s",
	}
)

type BankStatementFilesRepository interface {
	CreateBankStatementFiles(ctx context.Context, bankStatementFiles *model.BankStatementFiles, fieldsInsert ...BankStatementFilesField) error
	BulkCreateBankStatementFiles(ctx context.Context, bankStatementFilesList []*model.BankStatementFiles, fieldsInsert ...BankStatementFilesField) error
	ResolveBankStatementFiles(ctx context.Context, selectFields ...BankStatementFilesField) (model.BankStatementFilesList, error)
	ResolveBankStatementFilesByID(ctx context.Context, primaryID model.BankStatementFilesPrimaryID, selectFields ...BankStatementFilesField) (model.BankStatementFiles, error)
	UpdateBankStatementFilesByID(ctx context.Context, id model.BankStatementFilesPrimaryID, bankStatementFiles *model.BankStatementFiles, bankStatementFilesUpdateFields ...BankStatementFilesUpdateField) error
	UpdateBankStatementFilesByFilter(ctx context.Context, filter model.Filter, bankStatementFilesUpdateFields ...BankStatementFilesUpdateField) (rowsAffected int64, err error)
	BulkUpdateBankStatementFiles(ctx context.Context, bankStatementFilesListMap map[model.BankStatementFilesPrimaryID]*model.BankStatementFiles, BankStatementFilessMapUpdateFieldsRequest map[model.BankStatementFilesPrimaryID]BankStatementFilesUpdateFieldList) (err error)
	DeleteBankStatementFilesByID(ctx context.Context, id model.BankStatementFilesPrimaryID) error
	BulkDeleteBankStatementFilesByIDs(ctx context.Context, ids []model.BankStatementFilesPrimaryID) error
	ResolveBankStatementFilesByFilter(ctx context.Context, filter model.Filter) (result []model.BankStatementFilesFilterResult, err error)
	IsExistBankStatementFilesByIDs(ctx context.Context, ids []model.BankStatementFilesPrimaryID) (exists bool, notFoundIds []model.BankStatementFilesPrimaryID, err error)
	IsExistBankStatementFilesByID(ctx context.Context, id model.BankStatementFilesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
