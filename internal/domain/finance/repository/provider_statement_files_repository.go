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

func composeInsertFieldsAndParamsProviderStatementFiles(providerStatementFilesList []model.ProviderStatementFiles, fieldsInsert ...ProviderStatementFilesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewProviderStatementFilesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, providerStatementFiles := range providerStatementFilesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, providerStatementFiles.Id)
			case selectField.ProviderAccountId():
				args = append(args, providerStatementFiles.ProviderAccountId)
			case selectField.StatementType():
				args = append(args, providerStatementFiles.StatementType)
			case selectField.StatementDate():
				args = append(args, providerStatementFiles.StatementDate)
			case selectField.FileName():
				args = append(args, providerStatementFiles.FileName)
			case selectField.StorageUrl():
				args = append(args, providerStatementFiles.StorageUrl)
			case selectField.FileHash():
				args = append(args, providerStatementFiles.FileHash)
			case selectField.ParseStatus():
				args = append(args, providerStatementFiles.ParseStatus)
			case selectField.Metadata():
				args = append(args, providerStatementFiles.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, providerStatementFiles.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, providerStatementFiles.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, providerStatementFiles.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, providerStatementFiles.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, providerStatementFiles.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, providerStatementFiles.MetaDeletedBy)

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

func composeProviderStatementFilesCompositePrimaryKeyWhere(primaryIDs []model.ProviderStatementFilesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"provider_statement_files\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultProviderStatementFilesSelectFields() string {
	fields := NewProviderStatementFilesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeProviderStatementFilesSelectFields(selectFields ...ProviderStatementFilesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ProviderStatementFilesField string
type ProviderStatementFilesFieldList []ProviderStatementFilesField

type ProviderStatementFilesSelectFields struct {
}

func (ss ProviderStatementFilesSelectFields) Id() ProviderStatementFilesField {
	return ProviderStatementFilesField("id")
}

func (ss ProviderStatementFilesSelectFields) ProviderAccountId() ProviderStatementFilesField {
	return ProviderStatementFilesField("provider_account_id")
}

func (ss ProviderStatementFilesSelectFields) StatementType() ProviderStatementFilesField {
	return ProviderStatementFilesField("statement_type")
}

func (ss ProviderStatementFilesSelectFields) StatementDate() ProviderStatementFilesField {
	return ProviderStatementFilesField("statement_date")
}

func (ss ProviderStatementFilesSelectFields) FileName() ProviderStatementFilesField {
	return ProviderStatementFilesField("file_name")
}

func (ss ProviderStatementFilesSelectFields) StorageUrl() ProviderStatementFilesField {
	return ProviderStatementFilesField("storage_url")
}

func (ss ProviderStatementFilesSelectFields) FileHash() ProviderStatementFilesField {
	return ProviderStatementFilesField("file_hash")
}

func (ss ProviderStatementFilesSelectFields) ParseStatus() ProviderStatementFilesField {
	return ProviderStatementFilesField("parse_status")
}

func (ss ProviderStatementFilesSelectFields) Metadata() ProviderStatementFilesField {
	return ProviderStatementFilesField("metadata")
}

func (ss ProviderStatementFilesSelectFields) MetaCreatedAt() ProviderStatementFilesField {
	return ProviderStatementFilesField("meta_created_at")
}

func (ss ProviderStatementFilesSelectFields) MetaCreatedBy() ProviderStatementFilesField {
	return ProviderStatementFilesField("meta_created_by")
}

func (ss ProviderStatementFilesSelectFields) MetaUpdatedAt() ProviderStatementFilesField {
	return ProviderStatementFilesField("meta_updated_at")
}

func (ss ProviderStatementFilesSelectFields) MetaUpdatedBy() ProviderStatementFilesField {
	return ProviderStatementFilesField("meta_updated_by")
}

func (ss ProviderStatementFilesSelectFields) MetaDeletedAt() ProviderStatementFilesField {
	return ProviderStatementFilesField("meta_deleted_at")
}

func (ss ProviderStatementFilesSelectFields) MetaDeletedBy() ProviderStatementFilesField {
	return ProviderStatementFilesField("meta_deleted_by")
}

func (ss ProviderStatementFilesSelectFields) All() ProviderStatementFilesFieldList {
	return []ProviderStatementFilesField{
		ss.Id(),
		ss.ProviderAccountId(),
		ss.StatementType(),
		ss.StatementDate(),
		ss.FileName(),
		ss.StorageUrl(),
		ss.FileHash(),
		ss.ParseStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewProviderStatementFilesSelectFields() ProviderStatementFilesSelectFields {
	return ProviderStatementFilesSelectFields{}
}

type ProviderStatementFilesUpdateFieldOption struct {
	useIncrement bool
}
type ProviderStatementFilesUpdateField struct {
	providerStatementFilesField ProviderStatementFilesField
	opt                         ProviderStatementFilesUpdateFieldOption
	value                       interface{}
}
type ProviderStatementFilesUpdateFieldList []ProviderStatementFilesUpdateField

func defaultProviderStatementFilesUpdateFieldOption() ProviderStatementFilesUpdateFieldOption {
	return ProviderStatementFilesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementProviderStatementFilesOption(useIncrement bool) func(*ProviderStatementFilesUpdateFieldOption) {
	return func(pcufo *ProviderStatementFilesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewProviderStatementFilesUpdateField(field ProviderStatementFilesField, val interface{}, opts ...func(*ProviderStatementFilesUpdateFieldOption)) ProviderStatementFilesUpdateField {
	defaultOpt := defaultProviderStatementFilesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ProviderStatementFilesUpdateField{
		providerStatementFilesField: field,
		value:                       val,
		opt:                         defaultOpt,
	}
}
func defaultProviderStatementFilesUpdateFields(providerStatementFiles model.ProviderStatementFiles) (providerStatementFilesUpdateFieldList ProviderStatementFilesUpdateFieldList) {
	selectFields := NewProviderStatementFilesSelectFields()
	providerStatementFilesUpdateFieldList = append(providerStatementFilesUpdateFieldList,
		NewProviderStatementFilesUpdateField(selectFields.Id(), providerStatementFiles.Id),
		NewProviderStatementFilesUpdateField(selectFields.ProviderAccountId(), providerStatementFiles.ProviderAccountId),
		NewProviderStatementFilesUpdateField(selectFields.StatementType(), providerStatementFiles.StatementType),
		NewProviderStatementFilesUpdateField(selectFields.StatementDate(), providerStatementFiles.StatementDate),
		NewProviderStatementFilesUpdateField(selectFields.FileName(), providerStatementFiles.FileName),
		NewProviderStatementFilesUpdateField(selectFields.StorageUrl(), providerStatementFiles.StorageUrl),
		NewProviderStatementFilesUpdateField(selectFields.FileHash(), providerStatementFiles.FileHash),
		NewProviderStatementFilesUpdateField(selectFields.ParseStatus(), providerStatementFiles.ParseStatus),
		NewProviderStatementFilesUpdateField(selectFields.Metadata(), providerStatementFiles.Metadata),
		NewProviderStatementFilesUpdateField(selectFields.MetaCreatedAt(), providerStatementFiles.MetaCreatedAt),
		NewProviderStatementFilesUpdateField(selectFields.MetaCreatedBy(), providerStatementFiles.MetaCreatedBy),
		NewProviderStatementFilesUpdateField(selectFields.MetaUpdatedAt(), providerStatementFiles.MetaUpdatedAt),
		NewProviderStatementFilesUpdateField(selectFields.MetaUpdatedBy(), providerStatementFiles.MetaUpdatedBy),
		NewProviderStatementFilesUpdateField(selectFields.MetaDeletedAt(), providerStatementFiles.MetaDeletedAt),
		NewProviderStatementFilesUpdateField(selectFields.MetaDeletedBy(), providerStatementFiles.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsProviderStatementFilesCommand(providerStatementFilesUpdateFieldList ProviderStatementFilesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range providerStatementFilesUpdateFieldList {
		field := string(updateField.providerStatementFilesField)
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

func (repo *RepositoryImpl) BulkCreateProviderStatementFiles(ctx context.Context, providerStatementFilesList []*model.ProviderStatementFiles, fieldsInsert ...ProviderStatementFilesField) (err error) {
	var (
		fieldsStr                       string
		valueListStr                    []string
		argsList                        []interface{}
		primaryIds                      []model.ProviderStatementFilesPrimaryID
		providerStatementFilesValueList []model.ProviderStatementFiles
	)

	if len(fieldsInsert) == 0 {
		selectField := NewProviderStatementFilesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, providerStatementFiles := range providerStatementFilesList {

		primaryIds = append(primaryIds, providerStatementFiles.ToProviderStatementFilesPrimaryID())

		providerStatementFilesValueList = append(providerStatementFilesValueList, *providerStatementFiles)
	}

	_, notFoundIds, err := repo.IsExistProviderStatementFilesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderStatementFiles] failed checking providerStatementFiles whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ProviderStatementFilesPrimaryID{}
		mapNotFoundIds := map[model.ProviderStatementFilesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "providerStatementFiles", fmt.Sprintf("providerStatementFiles with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsProviderStatementFiles(providerStatementFilesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(providerStatementFilesQueries.insertProviderStatementFiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderStatementFiles] failed exec create providerStatementFiles query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteProviderStatementFilesByIDs(ctx context.Context, primaryIDs []model.ProviderStatementFilesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistProviderStatementFilesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderStatementFilesByIDs] failed checking providerStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementFiles with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_statement_files\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(providerStatementFilesQueries.deleteProviderStatementFiles + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderStatementFilesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderStatementFilesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistProviderStatementFilesByIDs(ctx context.Context, ids []model.ProviderStatementFilesPrimaryID) (exists bool, notFoundIds []model.ProviderStatementFilesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_statement_files\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(providerStatementFilesQueries.selectProviderStatementFiles, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderStatementFilesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ProviderStatementFilesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderStatementFilesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ProviderStatementFilesPrimaryID]bool{}
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

// BulkUpdateProviderStatementFiles is used to bulk update providerStatementFiles, by default it will update all field
// if want to update specific field, then fill providerStatementFilessMapUpdateFieldsRequest else please fill providerStatementFilessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateProviderStatementFiles(ctx context.Context, providerStatementFilessMap map[model.ProviderStatementFilesPrimaryID]*model.ProviderStatementFiles, providerStatementFilessMapUpdateFieldsRequest map[model.ProviderStatementFilesPrimaryID]ProviderStatementFilesUpdateFieldList) (err error) {
	if len(providerStatementFilessMap) == 0 && len(providerStatementFilessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		providerStatementFilessMapUpdateField map[model.ProviderStatementFilesPrimaryID]ProviderStatementFilesUpdateFieldList = map[model.ProviderStatementFilesPrimaryID]ProviderStatementFilesUpdateFieldList{}
		asTableValues                         string                                                                          = "myvalues"
	)

	if len(providerStatementFilessMap) > 0 {
		for id, providerStatementFiles := range providerStatementFilessMap {
			if providerStatementFiles == nil {
				log.Error().Err(err).Msg("[BulkUpdateProviderStatementFiles] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			providerStatementFilessMapUpdateField[id] = defaultProviderStatementFilesUpdateFields(*providerStatementFiles)
		}
	} else {
		providerStatementFilessMapUpdateField = providerStatementFilessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateProviderStatementFilesQuery(providerStatementFilessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistProviderStatementFilesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderStatementFiles] failed checking providerStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementFiles with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeProviderStatementFilesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"provider_statement_files\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderStatementFiles] failed exec query")
	}
	return
}

type ProviderStatementFilesFieldParameter struct {
	param string
	args  []interface{}
}

func NewProviderStatementFilesFieldParameter(param string, args ...interface{}) ProviderStatementFilesFieldParameter {
	return ProviderStatementFilesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateProviderStatementFilesQuery(mapProviderStatementFiless map[model.ProviderStatementFilesPrimaryID]ProviderStatementFilesUpdateFieldList, asTableValues string) (primaryIDs []model.ProviderStatementFilesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ProviderStatementFilesPrimaryID]map[string]interface{}{}
	providerStatementFilesSelectFields := NewProviderStatementFilesSelectFields()
	for id, updateFields := range mapProviderStatementFiless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.providerStatementFilesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapProviderStatementFiless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetProviderStatementFilesFieldType(updateField.providerStatementFilesField)))
			args = append(args, fields[string(updateField.providerStatementFilesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.providerStatementFilesField))
		if updateField.providerStatementFilesField == providerStatementFilesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.providerStatementFilesField, asTableValues, updateField.providerStatementFilesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.providerStatementFilesField,
				"\"provider_statement_files\"", updateField.providerStatementFilesField,
				asTableValues, updateField.providerStatementFilesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeProviderStatementFilesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ProviderStatementFilesPrimaryID, asTableValue string) (whereQry string) {
	providerStatementFilesSelectFields := NewProviderStatementFilesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"provider_statement_files\".\"id\" = %s.\"id\"::"+GetProviderStatementFilesFieldType(providerStatementFilesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetProviderStatementFilesFieldType(providerStatementFilesField ProviderStatementFilesField) string {
	selectProviderStatementFilesFields := NewProviderStatementFilesSelectFields()
	switch providerStatementFilesField {

	case selectProviderStatementFilesFields.Id():
		return "uuid"

	case selectProviderStatementFilesFields.ProviderAccountId():
		return "uuid"

	case selectProviderStatementFilesFields.StatementType():
		return "statement_type_enum"

	case selectProviderStatementFilesFields.StatementDate():
		return "date"

	case selectProviderStatementFilesFields.FileName():
		return "text"

	case selectProviderStatementFilesFields.StorageUrl():
		return "text"

	case selectProviderStatementFilesFields.FileHash():
		return "text"

	case selectProviderStatementFilesFields.ParseStatus():
		return "parse_status_enum"

	case selectProviderStatementFilesFields.Metadata():
		return "jsonb"

	case selectProviderStatementFilesFields.MetaCreatedAt():
		return "timestamptz"

	case selectProviderStatementFilesFields.MetaCreatedBy():
		return "uuid"

	case selectProviderStatementFilesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectProviderStatementFilesFields.MetaUpdatedBy():
		return "uuid"

	case selectProviderStatementFilesFields.MetaDeletedAt():
		return "timestamptz"

	case selectProviderStatementFilesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateProviderStatementFiles(ctx context.Context, providerStatementFiles *model.ProviderStatementFiles, fieldsInsert ...ProviderStatementFilesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewProviderStatementFilesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ProviderStatementFilesPrimaryID{
		Id: providerStatementFiles.Id,
	}
	exists, err := repo.IsExistProviderStatementFilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderStatementFiles] failed checking providerStatementFiles whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "providerStatementFiles", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsProviderStatementFiles([]model.ProviderStatementFiles{*providerStatementFiles}, fieldsInsert...)
	commandQuery := fmt.Sprintf(providerStatementFilesQueries.insertProviderStatementFiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderStatementFiles] failed exec create providerStatementFiles query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteProviderStatementFilesByID(ctx context.Context, primaryID model.ProviderStatementFilesPrimaryID) (err error) {
	exists, err := repo.IsExistProviderStatementFilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderStatementFilesByID] failed checking providerStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementFiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeProviderStatementFilesCompositePrimaryKeyWhere([]model.ProviderStatementFilesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(providerStatementFilesQueries.deleteProviderStatementFiles + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderStatementFilesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderStatementFilesByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderStatementFilesFilterResult, err error) {
	query, args, err := composeProviderStatementFilesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderStatementFilesByFilter] failed compose providerStatementFiles filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderStatementFilesByFilter] failed get providerStatementFiles by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeProviderStatementFilesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ProviderStatementFilesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeProviderStatementFilesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeProviderStatementFilesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeProviderStatementFilesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewProviderStatementFilesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["statement_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"statement_type\"")
			selectedColumns["statement_type"] = struct{}{}
		}
		if _, selected := selectedColumns["statement_date"]; !selected {
			selectColumns = append(selectColumns, "base.\"statement_date\"")
			selectedColumns["statement_date"] = struct{}{}
		}
		if _, selected := selectedColumns["file_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"file_name\"")
			selectedColumns["file_name"] = struct{}{}
		}
		if _, selected := selectedColumns["storage_url"]; !selected {
			selectColumns = append(selectColumns, "base.\"storage_url\"")
			selectedColumns["storage_url"] = struct{}{}
		}
		if _, selected := selectedColumns["file_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"file_hash\"")
			selectedColumns["file_hash"] = struct{}{}
		}
		if _, selected := selectedColumns["parse_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"parse_status\"")
			selectedColumns["parse_status"] = struct{}{}
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

type providerStatementFilesFilterPlaceholder struct {
	index int
}

func (p *providerStatementFilesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeProviderStatementFilesFilterPredicate(filterField model.FilterField, placeholders *providerStatementFilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewProviderStatementFilesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeProviderStatementFilesFilterSQLExpr(spec)
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

func composeProviderStatementFilesFilterGroup(group model.FilterGroup, placeholders *providerStatementFilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeProviderStatementFilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeProviderStatementFilesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeProviderStatementFilesFilterWhereQueries(filter model.Filter, placeholders *providerStatementFilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeProviderStatementFilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeProviderStatementFilesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeProviderStatementFilesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateProviderStatementFilesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeProviderStatementFilesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeProviderStatementFilesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := providerStatementFilesFilterPlaceholder{index: 1}
	whereQueries, err := composeProviderStatementFilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewProviderStatementFilesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeProviderStatementFilesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeProviderStatementFilesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"provider_statement_files\" base%s", strings.Join(selectColumns, ","), composeProviderStatementFilesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistProviderStatementFilesByID(ctx context.Context, primaryID model.ProviderStatementFilesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeProviderStatementFilesCompositePrimaryKeyWhere([]model.ProviderStatementFilesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", providerStatementFilesQueries.selectCountProviderStatementFiles, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderStatementFilesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderStatementFiles(ctx context.Context, selectFields ...ProviderStatementFilesField) (providerStatementFilesList model.ProviderStatementFilesList, err error) {
	var (
		defaultProviderStatementFilesSelectFields = defaultProviderStatementFilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderStatementFilesSelectFields = composeProviderStatementFilesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(providerStatementFilesQueries.selectProviderStatementFiles, defaultProviderStatementFilesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &providerStatementFilesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderStatementFiles] failed get providerStatementFiles list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderStatementFilesByID(ctx context.Context, primaryID model.ProviderStatementFilesPrimaryID, selectFields ...ProviderStatementFilesField) (providerStatementFiles model.ProviderStatementFiles, err error) {
	var (
		defaultProviderStatementFilesSelectFields = defaultProviderStatementFilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderStatementFilesSelectFields = composeProviderStatementFilesSelectFields(selectFields...)
	}
	whereQry, params := composeProviderStatementFilesCompositePrimaryKeyWhere([]model.ProviderStatementFilesPrimaryID{primaryID})
	query := fmt.Sprintf(providerStatementFilesQueries.selectProviderStatementFiles+" WHERE "+whereQry, defaultProviderStatementFilesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &providerStatementFiles, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("providerStatementFiles with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveProviderStatementFilesByID] failed get providerStatementFiles")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateProviderStatementFilesByID(ctx context.Context, primaryID model.ProviderStatementFilesPrimaryID, providerStatementFiles *model.ProviderStatementFiles, providerStatementFilesUpdateFields ...ProviderStatementFilesUpdateField) (err error) {
	exists, err := repo.IsExistProviderStatementFilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementFiles] failed checking providerStatementFiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementFiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if providerStatementFiles == nil {
		if len(providerStatementFilesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateProviderStatementFilesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		providerStatementFiles = &model.ProviderStatementFiles{}
	}
	var (
		defaultProviderStatementFilesUpdateFields = defaultProviderStatementFilesUpdateFields(*providerStatementFiles)
		tempUpdateField                           ProviderStatementFilesUpdateFieldList
		selectFields                              = NewProviderStatementFilesSelectFields()
	)
	if len(providerStatementFilesUpdateFields) > 0 {
		for _, updateField := range providerStatementFilesUpdateFields {
			if updateField.providerStatementFilesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultProviderStatementFilesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeProviderStatementFilesCompositePrimaryKeyWhere([]model.ProviderStatementFilesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsProviderStatementFilesCommand(defaultProviderStatementFilesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(providerStatementFilesQueries.updateProviderStatementFiles+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementFiles] error when try to update providerStatementFiles by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateProviderStatementFilesByFilter(ctx context.Context, filter model.Filter, providerStatementFilesUpdateFields ...ProviderStatementFilesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(providerStatementFilesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ProviderStatementFilesUpdateFieldList
		selectFields = NewProviderStatementFilesSelectFields()
	)
	for _, updateField := range providerStatementFilesUpdateFields {
		if updateField.providerStatementFilesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsProviderStatementFilesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := providerStatementFilesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeProviderStatementFilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"provider_statement_files\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementFilesByFilter] error when try to update providerStatementFiles by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementFilesByFilter] failed get rows affected")
	}
	return
}

var (
	providerStatementFilesQueries = struct {
		selectProviderStatementFiles      string
		selectCountProviderStatementFiles string
		deleteProviderStatementFiles      string
		updateProviderStatementFiles      string
		insertProviderStatementFiles      string
	}{
		selectProviderStatementFiles:      "SELECT %s FROM \"provider_statement_files\"",
		selectCountProviderStatementFiles: "SELECT COUNT(\"id\") FROM \"provider_statement_files\"",
		deleteProviderStatementFiles:      "DELETE FROM \"provider_statement_files\"",
		updateProviderStatementFiles:      "UPDATE \"provider_statement_files\" SET %s ",
		insertProviderStatementFiles:      "INSERT INTO \"provider_statement_files\" %s VALUES %s",
	}
)

type ProviderStatementFilesRepository interface {
	CreateProviderStatementFiles(ctx context.Context, providerStatementFiles *model.ProviderStatementFiles, fieldsInsert ...ProviderStatementFilesField) error
	BulkCreateProviderStatementFiles(ctx context.Context, providerStatementFilesList []*model.ProviderStatementFiles, fieldsInsert ...ProviderStatementFilesField) error
	ResolveProviderStatementFiles(ctx context.Context, selectFields ...ProviderStatementFilesField) (model.ProviderStatementFilesList, error)
	ResolveProviderStatementFilesByID(ctx context.Context, primaryID model.ProviderStatementFilesPrimaryID, selectFields ...ProviderStatementFilesField) (model.ProviderStatementFiles, error)
	UpdateProviderStatementFilesByID(ctx context.Context, id model.ProviderStatementFilesPrimaryID, providerStatementFiles *model.ProviderStatementFiles, providerStatementFilesUpdateFields ...ProviderStatementFilesUpdateField) error
	UpdateProviderStatementFilesByFilter(ctx context.Context, filter model.Filter, providerStatementFilesUpdateFields ...ProviderStatementFilesUpdateField) (rowsAffected int64, err error)
	BulkUpdateProviderStatementFiles(ctx context.Context, providerStatementFilesListMap map[model.ProviderStatementFilesPrimaryID]*model.ProviderStatementFiles, ProviderStatementFilessMapUpdateFieldsRequest map[model.ProviderStatementFilesPrimaryID]ProviderStatementFilesUpdateFieldList) (err error)
	DeleteProviderStatementFilesByID(ctx context.Context, id model.ProviderStatementFilesPrimaryID) error
	BulkDeleteProviderStatementFilesByIDs(ctx context.Context, ids []model.ProviderStatementFilesPrimaryID) error
	ResolveProviderStatementFilesByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderStatementFilesFilterResult, err error)
	IsExistProviderStatementFilesByIDs(ctx context.Context, ids []model.ProviderStatementFilesPrimaryID) (exists bool, notFoundIds []model.ProviderStatementFilesPrimaryID, err error)
	IsExistProviderStatementFilesByID(ctx context.Context, id model.ProviderStatementFilesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
