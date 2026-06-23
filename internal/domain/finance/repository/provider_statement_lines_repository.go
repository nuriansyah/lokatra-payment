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

func composeInsertFieldsAndParamsProviderStatementLines(providerStatementLinesList []model.ProviderStatementLines, fieldsInsert ...ProviderStatementLinesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewProviderStatementLinesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, providerStatementLines := range providerStatementLinesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, providerStatementLines.Id)
			case selectField.StatementFileId():
				args = append(args, providerStatementLines.StatementFileId)
			case selectField.LineNo():
				args = append(args, providerStatementLines.LineNo)
			case selectField.ProviderRef():
				args = append(args, providerStatementLines.ProviderRef)
			case selectField.TransactionType():
				args = append(args, providerStatementLines.TransactionType)
			case selectField.BookingDate():
				args = append(args, providerStatementLines.BookingDate)
			case selectField.ValueDate():
				args = append(args, providerStatementLines.ValueDate)
			case selectField.CurrencyCode():
				args = append(args, providerStatementLines.CurrencyCode)
			case selectField.GrossAmount():
				args = append(args, providerStatementLines.GrossAmount)
			case selectField.FeeAmount():
				args = append(args, providerStatementLines.FeeAmount)
			case selectField.NetAmount():
				args = append(args, providerStatementLines.NetAmount)
			case selectField.RawLine():
				args = append(args, providerStatementLines.RawLine)
			case selectField.LineHash():
				args = append(args, providerStatementLines.LineHash)
			case selectField.Metadata():
				args = append(args, providerStatementLines.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, providerStatementLines.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, providerStatementLines.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, providerStatementLines.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, providerStatementLines.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, providerStatementLines.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, providerStatementLines.MetaDeletedBy)

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

func composeProviderStatementLinesCompositePrimaryKeyWhere(primaryIDs []model.ProviderStatementLinesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"provider_statement_lines\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultProviderStatementLinesSelectFields() string {
	fields := NewProviderStatementLinesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeProviderStatementLinesSelectFields(selectFields ...ProviderStatementLinesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ProviderStatementLinesField string
type ProviderStatementLinesFieldList []ProviderStatementLinesField

type ProviderStatementLinesSelectFields struct {
}

func (ss ProviderStatementLinesSelectFields) Id() ProviderStatementLinesField {
	return ProviderStatementLinesField("id")
}

func (ss ProviderStatementLinesSelectFields) StatementFileId() ProviderStatementLinesField {
	return ProviderStatementLinesField("statement_file_id")
}

func (ss ProviderStatementLinesSelectFields) LineNo() ProviderStatementLinesField {
	return ProviderStatementLinesField("line_no")
}

func (ss ProviderStatementLinesSelectFields) ProviderRef() ProviderStatementLinesField {
	return ProviderStatementLinesField("provider_ref")
}

func (ss ProviderStatementLinesSelectFields) TransactionType() ProviderStatementLinesField {
	return ProviderStatementLinesField("transaction_type")
}

func (ss ProviderStatementLinesSelectFields) BookingDate() ProviderStatementLinesField {
	return ProviderStatementLinesField("booking_date")
}

func (ss ProviderStatementLinesSelectFields) ValueDate() ProviderStatementLinesField {
	return ProviderStatementLinesField("value_date")
}

func (ss ProviderStatementLinesSelectFields) CurrencyCode() ProviderStatementLinesField {
	return ProviderStatementLinesField("currency_code")
}

func (ss ProviderStatementLinesSelectFields) GrossAmount() ProviderStatementLinesField {
	return ProviderStatementLinesField("gross_amount")
}

func (ss ProviderStatementLinesSelectFields) FeeAmount() ProviderStatementLinesField {
	return ProviderStatementLinesField("fee_amount")
}

func (ss ProviderStatementLinesSelectFields) NetAmount() ProviderStatementLinesField {
	return ProviderStatementLinesField("net_amount")
}

func (ss ProviderStatementLinesSelectFields) RawLine() ProviderStatementLinesField {
	return ProviderStatementLinesField("raw_line")
}

func (ss ProviderStatementLinesSelectFields) LineHash() ProviderStatementLinesField {
	return ProviderStatementLinesField("line_hash")
}

func (ss ProviderStatementLinesSelectFields) Metadata() ProviderStatementLinesField {
	return ProviderStatementLinesField("metadata")
}

func (ss ProviderStatementLinesSelectFields) MetaCreatedAt() ProviderStatementLinesField {
	return ProviderStatementLinesField("meta_created_at")
}

func (ss ProviderStatementLinesSelectFields) MetaCreatedBy() ProviderStatementLinesField {
	return ProviderStatementLinesField("meta_created_by")
}

func (ss ProviderStatementLinesSelectFields) MetaUpdatedAt() ProviderStatementLinesField {
	return ProviderStatementLinesField("meta_updated_at")
}

func (ss ProviderStatementLinesSelectFields) MetaUpdatedBy() ProviderStatementLinesField {
	return ProviderStatementLinesField("meta_updated_by")
}

func (ss ProviderStatementLinesSelectFields) MetaDeletedAt() ProviderStatementLinesField {
	return ProviderStatementLinesField("meta_deleted_at")
}

func (ss ProviderStatementLinesSelectFields) MetaDeletedBy() ProviderStatementLinesField {
	return ProviderStatementLinesField("meta_deleted_by")
}

func (ss ProviderStatementLinesSelectFields) All() ProviderStatementLinesFieldList {
	return []ProviderStatementLinesField{
		ss.Id(),
		ss.StatementFileId(),
		ss.LineNo(),
		ss.ProviderRef(),
		ss.TransactionType(),
		ss.BookingDate(),
		ss.ValueDate(),
		ss.CurrencyCode(),
		ss.GrossAmount(),
		ss.FeeAmount(),
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

func NewProviderStatementLinesSelectFields() ProviderStatementLinesSelectFields {
	return ProviderStatementLinesSelectFields{}
}

type ProviderStatementLinesUpdateFieldOption struct {
	useIncrement bool
}
type ProviderStatementLinesUpdateField struct {
	providerStatementLinesField ProviderStatementLinesField
	opt                         ProviderStatementLinesUpdateFieldOption
	value                       interface{}
}
type ProviderStatementLinesUpdateFieldList []ProviderStatementLinesUpdateField

func defaultProviderStatementLinesUpdateFieldOption() ProviderStatementLinesUpdateFieldOption {
	return ProviderStatementLinesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementProviderStatementLinesOption(useIncrement bool) func(*ProviderStatementLinesUpdateFieldOption) {
	return func(pcufo *ProviderStatementLinesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewProviderStatementLinesUpdateField(field ProviderStatementLinesField, val interface{}, opts ...func(*ProviderStatementLinesUpdateFieldOption)) ProviderStatementLinesUpdateField {
	defaultOpt := defaultProviderStatementLinesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ProviderStatementLinesUpdateField{
		providerStatementLinesField: field,
		value:                       val,
		opt:                         defaultOpt,
	}
}
func defaultProviderStatementLinesUpdateFields(providerStatementLines model.ProviderStatementLines) (providerStatementLinesUpdateFieldList ProviderStatementLinesUpdateFieldList) {
	selectFields := NewProviderStatementLinesSelectFields()
	providerStatementLinesUpdateFieldList = append(providerStatementLinesUpdateFieldList,
		NewProviderStatementLinesUpdateField(selectFields.Id(), providerStatementLines.Id),
		NewProviderStatementLinesUpdateField(selectFields.StatementFileId(), providerStatementLines.StatementFileId),
		NewProviderStatementLinesUpdateField(selectFields.LineNo(), providerStatementLines.LineNo),
		NewProviderStatementLinesUpdateField(selectFields.ProviderRef(), providerStatementLines.ProviderRef),
		NewProviderStatementLinesUpdateField(selectFields.TransactionType(), providerStatementLines.TransactionType),
		NewProviderStatementLinesUpdateField(selectFields.BookingDate(), providerStatementLines.BookingDate),
		NewProviderStatementLinesUpdateField(selectFields.ValueDate(), providerStatementLines.ValueDate),
		NewProviderStatementLinesUpdateField(selectFields.CurrencyCode(), providerStatementLines.CurrencyCode),
		NewProviderStatementLinesUpdateField(selectFields.GrossAmount(), providerStatementLines.GrossAmount),
		NewProviderStatementLinesUpdateField(selectFields.FeeAmount(), providerStatementLines.FeeAmount),
		NewProviderStatementLinesUpdateField(selectFields.NetAmount(), providerStatementLines.NetAmount),
		NewProviderStatementLinesUpdateField(selectFields.RawLine(), providerStatementLines.RawLine),
		NewProviderStatementLinesUpdateField(selectFields.LineHash(), providerStatementLines.LineHash),
		NewProviderStatementLinesUpdateField(selectFields.Metadata(), providerStatementLines.Metadata),
		NewProviderStatementLinesUpdateField(selectFields.MetaCreatedAt(), providerStatementLines.MetaCreatedAt),
		NewProviderStatementLinesUpdateField(selectFields.MetaCreatedBy(), providerStatementLines.MetaCreatedBy),
		NewProviderStatementLinesUpdateField(selectFields.MetaUpdatedAt(), providerStatementLines.MetaUpdatedAt),
		NewProviderStatementLinesUpdateField(selectFields.MetaUpdatedBy(), providerStatementLines.MetaUpdatedBy),
		NewProviderStatementLinesUpdateField(selectFields.MetaDeletedAt(), providerStatementLines.MetaDeletedAt),
		NewProviderStatementLinesUpdateField(selectFields.MetaDeletedBy(), providerStatementLines.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsProviderStatementLinesCommand(providerStatementLinesUpdateFieldList ProviderStatementLinesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range providerStatementLinesUpdateFieldList {
		field := string(updateField.providerStatementLinesField)
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

func (repo *RepositoryImpl) BulkCreateProviderStatementLines(ctx context.Context, providerStatementLinesList []*model.ProviderStatementLines, fieldsInsert ...ProviderStatementLinesField) (err error) {
	var (
		fieldsStr                       string
		valueListStr                    []string
		argsList                        []interface{}
		primaryIds                      []model.ProviderStatementLinesPrimaryID
		providerStatementLinesValueList []model.ProviderStatementLines
	)

	if len(fieldsInsert) == 0 {
		selectField := NewProviderStatementLinesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, providerStatementLines := range providerStatementLinesList {

		primaryIds = append(primaryIds, providerStatementLines.ToProviderStatementLinesPrimaryID())

		providerStatementLinesValueList = append(providerStatementLinesValueList, *providerStatementLines)
	}

	_, notFoundIds, err := repo.IsExistProviderStatementLinesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderStatementLines] failed checking providerStatementLines whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ProviderStatementLinesPrimaryID{}
		mapNotFoundIds := map[model.ProviderStatementLinesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "providerStatementLines", fmt.Sprintf("providerStatementLines with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsProviderStatementLines(providerStatementLinesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(providerStatementLinesQueries.insertProviderStatementLines, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderStatementLines] failed exec create providerStatementLines query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteProviderStatementLinesByIDs(ctx context.Context, primaryIDs []model.ProviderStatementLinesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistProviderStatementLinesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderStatementLinesByIDs] failed checking providerStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementLines with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_statement_lines\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(providerStatementLinesQueries.deleteProviderStatementLines + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderStatementLinesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderStatementLinesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistProviderStatementLinesByIDs(ctx context.Context, ids []model.ProviderStatementLinesPrimaryID) (exists bool, notFoundIds []model.ProviderStatementLinesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_statement_lines\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(providerStatementLinesQueries.selectProviderStatementLines, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderStatementLinesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ProviderStatementLinesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderStatementLinesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ProviderStatementLinesPrimaryID]bool{}
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

// BulkUpdateProviderStatementLines is used to bulk update providerStatementLines, by default it will update all field
// if want to update specific field, then fill providerStatementLinessMapUpdateFieldsRequest else please fill providerStatementLinessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateProviderStatementLines(ctx context.Context, providerStatementLinessMap map[model.ProviderStatementLinesPrimaryID]*model.ProviderStatementLines, providerStatementLinessMapUpdateFieldsRequest map[model.ProviderStatementLinesPrimaryID]ProviderStatementLinesUpdateFieldList) (err error) {
	if len(providerStatementLinessMap) == 0 && len(providerStatementLinessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		providerStatementLinessMapUpdateField map[model.ProviderStatementLinesPrimaryID]ProviderStatementLinesUpdateFieldList = map[model.ProviderStatementLinesPrimaryID]ProviderStatementLinesUpdateFieldList{}
		asTableValues                         string                                                                          = "myvalues"
	)

	if len(providerStatementLinessMap) > 0 {
		for id, providerStatementLines := range providerStatementLinessMap {
			if providerStatementLines == nil {
				log.Error().Err(err).Msg("[BulkUpdateProviderStatementLines] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			providerStatementLinessMapUpdateField[id] = defaultProviderStatementLinesUpdateFields(*providerStatementLines)
		}
	} else {
		providerStatementLinessMapUpdateField = providerStatementLinessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateProviderStatementLinesQuery(providerStatementLinessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistProviderStatementLinesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderStatementLines] failed checking providerStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementLines with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeProviderStatementLinesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"provider_statement_lines\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderStatementLines] failed exec query")
	}
	return
}

type ProviderStatementLinesFieldParameter struct {
	param string
	args  []interface{}
}

func NewProviderStatementLinesFieldParameter(param string, args ...interface{}) ProviderStatementLinesFieldParameter {
	return ProviderStatementLinesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateProviderStatementLinesQuery(mapProviderStatementLiness map[model.ProviderStatementLinesPrimaryID]ProviderStatementLinesUpdateFieldList, asTableValues string) (primaryIDs []model.ProviderStatementLinesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ProviderStatementLinesPrimaryID]map[string]interface{}{}
	providerStatementLinesSelectFields := NewProviderStatementLinesSelectFields()
	for id, updateFields := range mapProviderStatementLiness {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.providerStatementLinesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapProviderStatementLiness[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetProviderStatementLinesFieldType(updateField.providerStatementLinesField)))
			args = append(args, fields[string(updateField.providerStatementLinesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.providerStatementLinesField))
		if updateField.providerStatementLinesField == providerStatementLinesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.providerStatementLinesField, asTableValues, updateField.providerStatementLinesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.providerStatementLinesField,
				"\"provider_statement_lines\"", updateField.providerStatementLinesField,
				asTableValues, updateField.providerStatementLinesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeProviderStatementLinesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ProviderStatementLinesPrimaryID, asTableValue string) (whereQry string) {
	providerStatementLinesSelectFields := NewProviderStatementLinesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"provider_statement_lines\".\"id\" = %s.\"id\"::"+GetProviderStatementLinesFieldType(providerStatementLinesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetProviderStatementLinesFieldType(providerStatementLinesField ProviderStatementLinesField) string {
	selectProviderStatementLinesFields := NewProviderStatementLinesSelectFields()
	switch providerStatementLinesField {

	case selectProviderStatementLinesFields.Id():
		return "uuid"

	case selectProviderStatementLinesFields.StatementFileId():
		return "uuid"

	case selectProviderStatementLinesFields.LineNo():
		return "int4"

	case selectProviderStatementLinesFields.ProviderRef():
		return "text"

	case selectProviderStatementLinesFields.TransactionType():
		return "text"

	case selectProviderStatementLinesFields.BookingDate():
		return "timestamptz"

	case selectProviderStatementLinesFields.ValueDate():
		return "timestamptz"

	case selectProviderStatementLinesFields.CurrencyCode():
		return "text"

	case selectProviderStatementLinesFields.GrossAmount():
		return "numeric"

	case selectProviderStatementLinesFields.FeeAmount():
		return "numeric"

	case selectProviderStatementLinesFields.NetAmount():
		return "numeric"

	case selectProviderStatementLinesFields.RawLine():
		return "jsonb"

	case selectProviderStatementLinesFields.LineHash():
		return "text"

	case selectProviderStatementLinesFields.Metadata():
		return "jsonb"

	case selectProviderStatementLinesFields.MetaCreatedAt():
		return "timestamptz"

	case selectProviderStatementLinesFields.MetaCreatedBy():
		return "uuid"

	case selectProviderStatementLinesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectProviderStatementLinesFields.MetaUpdatedBy():
		return "uuid"

	case selectProviderStatementLinesFields.MetaDeletedAt():
		return "timestamptz"

	case selectProviderStatementLinesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateProviderStatementLines(ctx context.Context, providerStatementLines *model.ProviderStatementLines, fieldsInsert ...ProviderStatementLinesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewProviderStatementLinesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ProviderStatementLinesPrimaryID{
		Id: providerStatementLines.Id,
	}
	exists, err := repo.IsExistProviderStatementLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderStatementLines] failed checking providerStatementLines whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "providerStatementLines", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsProviderStatementLines([]model.ProviderStatementLines{*providerStatementLines}, fieldsInsert...)
	commandQuery := fmt.Sprintf(providerStatementLinesQueries.insertProviderStatementLines, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderStatementLines] failed exec create providerStatementLines query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteProviderStatementLinesByID(ctx context.Context, primaryID model.ProviderStatementLinesPrimaryID) (err error) {
	exists, err := repo.IsExistProviderStatementLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderStatementLinesByID] failed checking providerStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementLines with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeProviderStatementLinesCompositePrimaryKeyWhere([]model.ProviderStatementLinesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(providerStatementLinesQueries.deleteProviderStatementLines + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderStatementLinesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderStatementLinesByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderStatementLinesFilterResult, err error) {
	query, args, err := composeProviderStatementLinesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderStatementLinesByFilter] failed compose providerStatementLines filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderStatementLinesByFilter] failed get providerStatementLines by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeProviderStatementLinesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ProviderStatementLinesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeProviderStatementLinesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeProviderStatementLinesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeProviderStatementLinesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewProviderStatementLinesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["provider_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_ref\"")
			selectedColumns["provider_ref"] = struct{}{}
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
		if _, selected := selectedColumns["gross_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"gross_amount\"")
			selectedColumns["gross_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_amount\"")
			selectedColumns["fee_amount"] = struct{}{}
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

type providerStatementLinesFilterPlaceholder struct {
	index int
}

func (p *providerStatementLinesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeProviderStatementLinesFilterPredicate(filterField model.FilterField, placeholders *providerStatementLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewProviderStatementLinesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeProviderStatementLinesFilterSQLExpr(spec)
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

func composeProviderStatementLinesFilterGroup(group model.FilterGroup, placeholders *providerStatementLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeProviderStatementLinesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeProviderStatementLinesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeProviderStatementLinesFilterWhereQueries(filter model.Filter, placeholders *providerStatementLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeProviderStatementLinesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeProviderStatementLinesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeProviderStatementLinesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateProviderStatementLinesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeProviderStatementLinesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeProviderStatementLinesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := providerStatementLinesFilterPlaceholder{index: 1}
	whereQueries, err := composeProviderStatementLinesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewProviderStatementLinesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeProviderStatementLinesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeProviderStatementLinesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"provider_statement_lines\" base%s", strings.Join(selectColumns, ","), composeProviderStatementLinesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistProviderStatementLinesByID(ctx context.Context, primaryID model.ProviderStatementLinesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeProviderStatementLinesCompositePrimaryKeyWhere([]model.ProviderStatementLinesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", providerStatementLinesQueries.selectCountProviderStatementLines, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderStatementLinesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderStatementLines(ctx context.Context, selectFields ...ProviderStatementLinesField) (providerStatementLinesList model.ProviderStatementLinesList, err error) {
	var (
		defaultProviderStatementLinesSelectFields = defaultProviderStatementLinesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderStatementLinesSelectFields = composeProviderStatementLinesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(providerStatementLinesQueries.selectProviderStatementLines, defaultProviderStatementLinesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &providerStatementLinesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderStatementLines] failed get providerStatementLines list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderStatementLinesByID(ctx context.Context, primaryID model.ProviderStatementLinesPrimaryID, selectFields ...ProviderStatementLinesField) (providerStatementLines model.ProviderStatementLines, err error) {
	var (
		defaultProviderStatementLinesSelectFields = defaultProviderStatementLinesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderStatementLinesSelectFields = composeProviderStatementLinesSelectFields(selectFields...)
	}
	whereQry, params := composeProviderStatementLinesCompositePrimaryKeyWhere([]model.ProviderStatementLinesPrimaryID{primaryID})
	query := fmt.Sprintf(providerStatementLinesQueries.selectProviderStatementLines+" WHERE "+whereQry, defaultProviderStatementLinesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &providerStatementLines, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("providerStatementLines with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveProviderStatementLinesByID] failed get providerStatementLines")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateProviderStatementLinesByID(ctx context.Context, primaryID model.ProviderStatementLinesPrimaryID, providerStatementLines *model.ProviderStatementLines, providerStatementLinesUpdateFields ...ProviderStatementLinesUpdateField) (err error) {
	exists, err := repo.IsExistProviderStatementLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementLines] failed checking providerStatementLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerStatementLines with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if providerStatementLines == nil {
		if len(providerStatementLinesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateProviderStatementLinesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		providerStatementLines = &model.ProviderStatementLines{}
	}
	var (
		defaultProviderStatementLinesUpdateFields = defaultProviderStatementLinesUpdateFields(*providerStatementLines)
		tempUpdateField                           ProviderStatementLinesUpdateFieldList
		selectFields                              = NewProviderStatementLinesSelectFields()
	)
	if len(providerStatementLinesUpdateFields) > 0 {
		for _, updateField := range providerStatementLinesUpdateFields {
			if updateField.providerStatementLinesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultProviderStatementLinesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeProviderStatementLinesCompositePrimaryKeyWhere([]model.ProviderStatementLinesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsProviderStatementLinesCommand(defaultProviderStatementLinesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(providerStatementLinesQueries.updateProviderStatementLines+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementLines] error when try to update providerStatementLines by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateProviderStatementLinesByFilter(ctx context.Context, filter model.Filter, providerStatementLinesUpdateFields ...ProviderStatementLinesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(providerStatementLinesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ProviderStatementLinesUpdateFieldList
		selectFields = NewProviderStatementLinesSelectFields()
	)
	for _, updateField := range providerStatementLinesUpdateFields {
		if updateField.providerStatementLinesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsProviderStatementLinesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := providerStatementLinesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeProviderStatementLinesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"provider_statement_lines\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementLinesByFilter] error when try to update providerStatementLines by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderStatementLinesByFilter] failed get rows affected")
	}
	return
}

var (
	providerStatementLinesQueries = struct {
		selectProviderStatementLines      string
		selectCountProviderStatementLines string
		deleteProviderStatementLines      string
		updateProviderStatementLines      string
		insertProviderStatementLines      string
	}{
		selectProviderStatementLines:      "SELECT %s FROM \"provider_statement_lines\"",
		selectCountProviderStatementLines: "SELECT COUNT(\"id\") FROM \"provider_statement_lines\"",
		deleteProviderStatementLines:      "DELETE FROM \"provider_statement_lines\"",
		updateProviderStatementLines:      "UPDATE \"provider_statement_lines\" SET %s ",
		insertProviderStatementLines:      "INSERT INTO \"provider_statement_lines\" %s VALUES %s",
	}
)

type ProviderStatementLinesRepository interface {
	CreateProviderStatementLines(ctx context.Context, providerStatementLines *model.ProviderStatementLines, fieldsInsert ...ProviderStatementLinesField) error
	BulkCreateProviderStatementLines(ctx context.Context, providerStatementLinesList []*model.ProviderStatementLines, fieldsInsert ...ProviderStatementLinesField) error
	ResolveProviderStatementLines(ctx context.Context, selectFields ...ProviderStatementLinesField) (model.ProviderStatementLinesList, error)
	ResolveProviderStatementLinesByID(ctx context.Context, primaryID model.ProviderStatementLinesPrimaryID, selectFields ...ProviderStatementLinesField) (model.ProviderStatementLines, error)
	UpdateProviderStatementLinesByID(ctx context.Context, id model.ProviderStatementLinesPrimaryID, providerStatementLines *model.ProviderStatementLines, providerStatementLinesUpdateFields ...ProviderStatementLinesUpdateField) error
	UpdateProviderStatementLinesByFilter(ctx context.Context, filter model.Filter, providerStatementLinesUpdateFields ...ProviderStatementLinesUpdateField) (rowsAffected int64, err error)
	BulkUpdateProviderStatementLines(ctx context.Context, providerStatementLinesListMap map[model.ProviderStatementLinesPrimaryID]*model.ProviderStatementLines, ProviderStatementLinessMapUpdateFieldsRequest map[model.ProviderStatementLinesPrimaryID]ProviderStatementLinesUpdateFieldList) (err error)
	DeleteProviderStatementLinesByID(ctx context.Context, id model.ProviderStatementLinesPrimaryID) error
	BulkDeleteProviderStatementLinesByIDs(ctx context.Context, ids []model.ProviderStatementLinesPrimaryID) error
	ResolveProviderStatementLinesByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderStatementLinesFilterResult, err error)
	IsExistProviderStatementLinesByIDs(ctx context.Context, ids []model.ProviderStatementLinesPrimaryID) (exists bool, notFoundIds []model.ProviderStatementLinesPrimaryID, err error)
	IsExistProviderStatementLinesByID(ctx context.Context, id model.ProviderStatementLinesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
