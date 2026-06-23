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

func composeInsertFieldsAndParamsTaxDocumentSequences(taxDocumentSequencesList []model.TaxDocumentSequences, fieldsInsert ...TaxDocumentSequencesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewTaxDocumentSequencesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, taxDocumentSequences := range taxDocumentSequencesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.SequenceCode():
				args = append(args, taxDocumentSequences.SequenceCode)
			case selectField.CurrentValue():
				args = append(args, taxDocumentSequences.CurrentValue)
			case selectField.Metadata():
				args = append(args, taxDocumentSequences.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, taxDocumentSequences.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, taxDocumentSequences.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, taxDocumentSequences.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, taxDocumentSequences.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, taxDocumentSequences.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, taxDocumentSequences.MetaDeletedBy)

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

func composeTaxDocumentSequencesCompositePrimaryKeyWhere(primaryIDs []model.TaxDocumentSequencesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		sequenceCode := "\"tax_document_sequences\".\"sequence_code\" = ?"
		params = append(params, primaryID.SequenceCode)
		arrWhereQry = append(arrWhereQry, sequenceCode)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultTaxDocumentSequencesSelectFields() string {
	fields := NewTaxDocumentSequencesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeTaxDocumentSequencesSelectFields(selectFields ...TaxDocumentSequencesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type TaxDocumentSequencesField string
type TaxDocumentSequencesFieldList []TaxDocumentSequencesField

type TaxDocumentSequencesSelectFields struct {
}

func (ss TaxDocumentSequencesSelectFields) SequenceCode() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("sequence_code")
}

func (ss TaxDocumentSequencesSelectFields) CurrentValue() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("current_value")
}

func (ss TaxDocumentSequencesSelectFields) Metadata() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("metadata")
}

func (ss TaxDocumentSequencesSelectFields) MetaCreatedAt() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("meta_created_at")
}

func (ss TaxDocumentSequencesSelectFields) MetaCreatedBy() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("meta_created_by")
}

func (ss TaxDocumentSequencesSelectFields) MetaUpdatedAt() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("meta_updated_at")
}

func (ss TaxDocumentSequencesSelectFields) MetaUpdatedBy() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("meta_updated_by")
}

func (ss TaxDocumentSequencesSelectFields) MetaDeletedAt() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("meta_deleted_at")
}

func (ss TaxDocumentSequencesSelectFields) MetaDeletedBy() TaxDocumentSequencesField {
	return TaxDocumentSequencesField("meta_deleted_by")
}

func (ss TaxDocumentSequencesSelectFields) All() TaxDocumentSequencesFieldList {
	return []TaxDocumentSequencesField{
		ss.SequenceCode(),
		ss.CurrentValue(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewTaxDocumentSequencesSelectFields() TaxDocumentSequencesSelectFields {
	return TaxDocumentSequencesSelectFields{}
}

type TaxDocumentSequencesUpdateFieldOption struct {
	useIncrement bool
}
type TaxDocumentSequencesUpdateField struct {
	taxDocumentSequencesField TaxDocumentSequencesField
	opt                       TaxDocumentSequencesUpdateFieldOption
	value                     interface{}
}
type TaxDocumentSequencesUpdateFieldList []TaxDocumentSequencesUpdateField

func defaultTaxDocumentSequencesUpdateFieldOption() TaxDocumentSequencesUpdateFieldOption {
	return TaxDocumentSequencesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementTaxDocumentSequencesOption(useIncrement bool) func(*TaxDocumentSequencesUpdateFieldOption) {
	return func(pcufo *TaxDocumentSequencesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewTaxDocumentSequencesUpdateField(field TaxDocumentSequencesField, val interface{}, opts ...func(*TaxDocumentSequencesUpdateFieldOption)) TaxDocumentSequencesUpdateField {
	defaultOpt := defaultTaxDocumentSequencesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return TaxDocumentSequencesUpdateField{
		taxDocumentSequencesField: field,
		value:                     val,
		opt:                       defaultOpt,
	}
}
func defaultTaxDocumentSequencesUpdateFields(taxDocumentSequences model.TaxDocumentSequences) (taxDocumentSequencesUpdateFieldList TaxDocumentSequencesUpdateFieldList) {
	selectFields := NewTaxDocumentSequencesSelectFields()
	taxDocumentSequencesUpdateFieldList = append(taxDocumentSequencesUpdateFieldList,
		NewTaxDocumentSequencesUpdateField(selectFields.SequenceCode(), taxDocumentSequences.SequenceCode),
		NewTaxDocumentSequencesUpdateField(selectFields.CurrentValue(), taxDocumentSequences.CurrentValue),
		NewTaxDocumentSequencesUpdateField(selectFields.Metadata(), taxDocumentSequences.Metadata),
		NewTaxDocumentSequencesUpdateField(selectFields.MetaCreatedAt(), taxDocumentSequences.MetaCreatedAt),
		NewTaxDocumentSequencesUpdateField(selectFields.MetaCreatedBy(), taxDocumentSequences.MetaCreatedBy),
		NewTaxDocumentSequencesUpdateField(selectFields.MetaUpdatedAt(), taxDocumentSequences.MetaUpdatedAt),
		NewTaxDocumentSequencesUpdateField(selectFields.MetaUpdatedBy(), taxDocumentSequences.MetaUpdatedBy),
		NewTaxDocumentSequencesUpdateField(selectFields.MetaDeletedAt(), taxDocumentSequences.MetaDeletedAt),
		NewTaxDocumentSequencesUpdateField(selectFields.MetaDeletedBy(), taxDocumentSequences.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsTaxDocumentSequencesCommand(taxDocumentSequencesUpdateFieldList TaxDocumentSequencesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range taxDocumentSequencesUpdateFieldList {
		field := string(updateField.taxDocumentSequencesField)
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

func (repo *RepositoryImpl) BulkCreateTaxDocumentSequences(ctx context.Context, taxDocumentSequencesList []*model.TaxDocumentSequences, fieldsInsert ...TaxDocumentSequencesField) (err error) {
	var (
		fieldsStr                     string
		valueListStr                  []string
		argsList                      []interface{}
		primaryIds                    []model.TaxDocumentSequencesPrimaryID
		taxDocumentSequencesValueList []model.TaxDocumentSequences
	)

	if len(fieldsInsert) == 0 {
		selectField := NewTaxDocumentSequencesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, taxDocumentSequences := range taxDocumentSequencesList {

		primaryIds = append(primaryIds, taxDocumentSequences.ToTaxDocumentSequencesPrimaryID())

		taxDocumentSequencesValueList = append(taxDocumentSequencesValueList, *taxDocumentSequences)
	}

	_, notFoundIds, err := repo.IsExistTaxDocumentSequencesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxDocumentSequences] failed checking taxDocumentSequences whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.TaxDocumentSequencesPrimaryID{}
		mapNotFoundIds := map[model.TaxDocumentSequencesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "taxDocumentSequences", fmt.Sprintf("taxDocumentSequences with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsTaxDocumentSequences(taxDocumentSequencesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(taxDocumentSequencesQueries.insertTaxDocumentSequences, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxDocumentSequences] failed exec create taxDocumentSequences query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteTaxDocumentSequencesByIDs(ctx context.Context, primaryIDs []model.TaxDocumentSequencesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistTaxDocumentSequencesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxDocumentSequencesByIDs] failed checking taxDocumentSequences whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentSequences with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_document_sequences\".\"sequence_code\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.SequenceCode)
	}

	commandQuery := fmt.Sprintf(taxDocumentSequencesQueries.deleteTaxDocumentSequences + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxDocumentSequencesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxDocumentSequencesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistTaxDocumentSequencesByIDs(ctx context.Context, ids []model.TaxDocumentSequencesPrimaryID) (exists bool, notFoundIds []model.TaxDocumentSequencesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_document_sequences\".\"sequence_code\" IN (?)"
	for _, id := range ids {
		params = append(params, id.SequenceCode)
	}

	query := fmt.Sprintf(taxDocumentSequencesQueries.selectTaxDocumentSequences, " \"sequence_code\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxDocumentSequencesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.TaxDocumentSequencesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxDocumentSequencesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.TaxDocumentSequencesPrimaryID]bool{}
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

// BulkUpdateTaxDocumentSequences is used to bulk update taxDocumentSequences, by default it will update all field
// if want to update specific field, then fill taxDocumentSequencessMapUpdateFieldsRequest else please fill taxDocumentSequencessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateTaxDocumentSequences(ctx context.Context, taxDocumentSequencessMap map[model.TaxDocumentSequencesPrimaryID]*model.TaxDocumentSequences, taxDocumentSequencessMapUpdateFieldsRequest map[model.TaxDocumentSequencesPrimaryID]TaxDocumentSequencesUpdateFieldList) (err error) {
	if len(taxDocumentSequencessMap) == 0 && len(taxDocumentSequencessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		taxDocumentSequencessMapUpdateField map[model.TaxDocumentSequencesPrimaryID]TaxDocumentSequencesUpdateFieldList = map[model.TaxDocumentSequencesPrimaryID]TaxDocumentSequencesUpdateFieldList{}
		asTableValues                       string                                                                      = "myvalues"
	)

	if len(taxDocumentSequencessMap) > 0 {
		for id, taxDocumentSequences := range taxDocumentSequencessMap {
			if taxDocumentSequences == nil {
				log.Error().Err(err).Msg("[BulkUpdateTaxDocumentSequences] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			taxDocumentSequencessMapUpdateField[id] = defaultTaxDocumentSequencesUpdateFields(*taxDocumentSequences)
		}
	} else {
		taxDocumentSequencessMapUpdateField = taxDocumentSequencessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateTaxDocumentSequencesQuery(taxDocumentSequencessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistTaxDocumentSequencesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxDocumentSequences] failed checking taxDocumentSequences whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentSequences with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeTaxDocumentSequencesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"tax_document_sequences\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxDocumentSequences] failed exec query")
	}
	return
}

type TaxDocumentSequencesFieldParameter struct {
	param string
	args  []interface{}
}

func NewTaxDocumentSequencesFieldParameter(param string, args ...interface{}) TaxDocumentSequencesFieldParameter {
	return TaxDocumentSequencesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateTaxDocumentSequencesQuery(mapTaxDocumentSequencess map[model.TaxDocumentSequencesPrimaryID]TaxDocumentSequencesUpdateFieldList, asTableValues string) (primaryIDs []model.TaxDocumentSequencesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.TaxDocumentSequencesPrimaryID]map[string]interface{}{}
	taxDocumentSequencesSelectFields := NewTaxDocumentSequencesSelectFields()
	for id, updateFields := range mapTaxDocumentSequencess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.taxDocumentSequencesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapTaxDocumentSequencess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetTaxDocumentSequencesFieldType(updateField.taxDocumentSequencesField)))
			args = append(args, fields[string(updateField.taxDocumentSequencesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.taxDocumentSequencesField))
		if updateField.taxDocumentSequencesField == taxDocumentSequencesSelectFields.SequenceCode() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.taxDocumentSequencesField, asTableValues, updateField.taxDocumentSequencesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.taxDocumentSequencesField,
				"\"tax_document_sequences\"", updateField.taxDocumentSequencesField,
				asTableValues, updateField.taxDocumentSequencesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeTaxDocumentSequencesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.TaxDocumentSequencesPrimaryID, asTableValue string) (whereQry string) {
	taxDocumentSequencesSelectFields := NewTaxDocumentSequencesSelectFields()
	var arrWhereQry []string
	sequenceCode := fmt.Sprintf("\"tax_document_sequences\".\"sequence_code\" = %s.\"sequence_code\"::"+GetTaxDocumentSequencesFieldType(taxDocumentSequencesSelectFields.SequenceCode()), asTableValue)
	arrWhereQry = append(arrWhereQry, sequenceCode)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetTaxDocumentSequencesFieldType(taxDocumentSequencesField TaxDocumentSequencesField) string {
	selectTaxDocumentSequencesFields := NewTaxDocumentSequencesSelectFields()
	switch taxDocumentSequencesField {

	case selectTaxDocumentSequencesFields.SequenceCode():
		return "text"

	case selectTaxDocumentSequencesFields.CurrentValue():
		return "int8"

	case selectTaxDocumentSequencesFields.Metadata():
		return "jsonb"

	case selectTaxDocumentSequencesFields.MetaCreatedAt():
		return "timestamptz"

	case selectTaxDocumentSequencesFields.MetaCreatedBy():
		return "uuid"

	case selectTaxDocumentSequencesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectTaxDocumentSequencesFields.MetaUpdatedBy():
		return "uuid"

	case selectTaxDocumentSequencesFields.MetaDeletedAt():
		return "timestamptz"

	case selectTaxDocumentSequencesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateTaxDocumentSequences(ctx context.Context, taxDocumentSequences *model.TaxDocumentSequences, fieldsInsert ...TaxDocumentSequencesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewTaxDocumentSequencesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.TaxDocumentSequencesPrimaryID{
		SequenceCode: taxDocumentSequences.SequenceCode,
	}
	exists, err := repo.IsExistTaxDocumentSequencesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxDocumentSequences] failed checking taxDocumentSequences whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "taxDocumentSequences", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsTaxDocumentSequences([]model.TaxDocumentSequences{*taxDocumentSequences}, fieldsInsert...)
	commandQuery := fmt.Sprintf(taxDocumentSequencesQueries.insertTaxDocumentSequences, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxDocumentSequences] failed exec create taxDocumentSequences query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteTaxDocumentSequencesByID(ctx context.Context, primaryID model.TaxDocumentSequencesPrimaryID) (err error) {
	exists, err := repo.IsExistTaxDocumentSequencesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxDocumentSequencesByID] failed checking taxDocumentSequences whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentSequences with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeTaxDocumentSequencesCompositePrimaryKeyWhere([]model.TaxDocumentSequencesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(taxDocumentSequencesQueries.deleteTaxDocumentSequences + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxDocumentSequencesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxDocumentSequencesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxDocumentSequencesFilterResult, err error) {
	query, args, err := composeTaxDocumentSequencesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxDocumentSequencesByFilter] failed compose taxDocumentSequences filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxDocumentSequencesByFilter] failed get taxDocumentSequences by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeTaxDocumentSequencesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.TaxDocumentSequencesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeTaxDocumentSequencesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeTaxDocumentSequencesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeTaxDocumentSequencesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 9 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewTaxDocumentSequencesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 9+1)
		if _, selected := selectedColumns["sequence_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"sequence_code\"")
			selectedColumns["sequence_code"] = struct{}{}
		}
		if _, selected := selectedColumns["current_value"]; !selected {
			selectColumns = append(selectColumns, "base.\"current_value\"")
			selectedColumns["current_value"] = struct{}{}
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

	if _, selected := selectedColumns["sequence_code"]; isCursorMode && !selected {
		selectColumns = append(selectColumns, "base.\"sequence_code\"")
		selectedColumns["sequence_code"] = struct{}{}
	}

	return
}

type taxDocumentSequencesFilterPlaceholder struct {
	index int
}

func (p *taxDocumentSequencesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeTaxDocumentSequencesFilterPredicate(filterField model.FilterField, placeholders *taxDocumentSequencesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewTaxDocumentSequencesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeTaxDocumentSequencesFilterSQLExpr(spec)
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

func composeTaxDocumentSequencesFilterGroup(group model.FilterGroup, placeholders *taxDocumentSequencesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeTaxDocumentSequencesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeTaxDocumentSequencesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeTaxDocumentSequencesFilterWhereQueries(filter model.Filter, placeholders *taxDocumentSequencesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeTaxDocumentSequencesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeTaxDocumentSequencesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeTaxDocumentSequencesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateTaxDocumentSequencesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeTaxDocumentSequencesSortOrder(filter.Sorts[0].Order)
			if sortErr != nil {
				err = sortErr
				return
			}
			if filter.Sorts[0].Field != "sequence_code" || sortOrder != model.SortAsc {
				err = failure.BadRequestFromString("cursor pagination only supports the default primary-key sort")
				return
			}
		}
	}

	selectColumns, err := composeTaxDocumentSequencesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := taxDocumentSequencesFilterPlaceholder{index: 1}
	whereQueries, err := composeTaxDocumentSequencesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
	if err != nil {
		return
	}

	if isCursorMode && filter.Pagination.Cursor != nil {
		whereQueries = append(whereQueries, fmt.Sprintf("base.\"sequence_code\" %s %s", cursorOperator, placeholders.Next()))
		args = append(args, filter.Pagination.Cursor)
	}

	sortQuery := []string{}
	if isCursorMode {
		sortQuery = append(sortQuery, fmt.Sprintf("base.\"sequence_code\" %s", cursorSortOrder))

	} else {
		for _, sort := range filter.Sorts {
			spec, found := model.NewTaxDocumentSequencesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeTaxDocumentSequencesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeTaxDocumentSequencesSortOrder(sort.Order)
			if sortErr != nil {
				err = sortErr
				return
			}
			sortQuery = append(sortQuery, fmt.Sprintf("%s %s", sqlExpr, sortOrder))
		}
		if len(sortQuery) == 0 {
			sortQuery = append(sortQuery, "base.\"sequence_code\" ASC")
		}

	}

	query = fmt.Sprintf("SELECT %s FROM \"tax_document_sequences\" base%s", strings.Join(selectColumns, ","), composeTaxDocumentSequencesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistTaxDocumentSequencesByID(ctx context.Context, primaryID model.TaxDocumentSequencesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeTaxDocumentSequencesCompositePrimaryKeyWhere([]model.TaxDocumentSequencesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", taxDocumentSequencesQueries.selectCountTaxDocumentSequences, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxDocumentSequencesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxDocumentSequences(ctx context.Context, selectFields ...TaxDocumentSequencesField) (taxDocumentSequencesList model.TaxDocumentSequencesList, err error) {
	var (
		defaultTaxDocumentSequencesSelectFields = defaultTaxDocumentSequencesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxDocumentSequencesSelectFields = composeTaxDocumentSequencesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(taxDocumentSequencesQueries.selectTaxDocumentSequences, defaultTaxDocumentSequencesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &taxDocumentSequencesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxDocumentSequences] failed get taxDocumentSequences list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxDocumentSequencesByID(ctx context.Context, primaryID model.TaxDocumentSequencesPrimaryID, selectFields ...TaxDocumentSequencesField) (taxDocumentSequences model.TaxDocumentSequences, err error) {
	var (
		defaultTaxDocumentSequencesSelectFields = defaultTaxDocumentSequencesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxDocumentSequencesSelectFields = composeTaxDocumentSequencesSelectFields(selectFields...)
	}
	whereQry, params := composeTaxDocumentSequencesCompositePrimaryKeyWhere([]model.TaxDocumentSequencesPrimaryID{primaryID})
	query := fmt.Sprintf(taxDocumentSequencesQueries.selectTaxDocumentSequences+" WHERE "+whereQry, defaultTaxDocumentSequencesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &taxDocumentSequences, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("taxDocumentSequences with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveTaxDocumentSequencesByID] failed get taxDocumentSequences")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateTaxDocumentSequencesByID(ctx context.Context, primaryID model.TaxDocumentSequencesPrimaryID, taxDocumentSequences *model.TaxDocumentSequences, taxDocumentSequencesUpdateFields ...TaxDocumentSequencesUpdateField) (err error) {
	exists, err := repo.IsExistTaxDocumentSequencesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentSequences] failed checking taxDocumentSequences whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentSequences with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if taxDocumentSequences == nil {
		if len(taxDocumentSequencesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateTaxDocumentSequencesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		taxDocumentSequences = &model.TaxDocumentSequences{}
	}
	var (
		defaultTaxDocumentSequencesUpdateFields = defaultTaxDocumentSequencesUpdateFields(*taxDocumentSequences)
		tempUpdateField                         TaxDocumentSequencesUpdateFieldList
		selectFields                            = NewTaxDocumentSequencesSelectFields()
	)
	if len(taxDocumentSequencesUpdateFields) > 0 {
		for _, updateField := range taxDocumentSequencesUpdateFields {
			if updateField.taxDocumentSequencesField == selectFields.SequenceCode() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultTaxDocumentSequencesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeTaxDocumentSequencesCompositePrimaryKeyWhere([]model.TaxDocumentSequencesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsTaxDocumentSequencesCommand(defaultTaxDocumentSequencesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(taxDocumentSequencesQueries.updateTaxDocumentSequences+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentSequences] error when try to update taxDocumentSequences by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateTaxDocumentSequencesByFilter(ctx context.Context, filter model.Filter, taxDocumentSequencesUpdateFields ...TaxDocumentSequencesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(taxDocumentSequencesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields TaxDocumentSequencesUpdateFieldList
		selectFields = NewTaxDocumentSequencesSelectFields()
	)
	for _, updateField := range taxDocumentSequencesUpdateFields {
		if updateField.taxDocumentSequencesField == selectFields.SequenceCode() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsTaxDocumentSequencesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := taxDocumentSequencesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeTaxDocumentSequencesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"tax_document_sequences\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentSequencesByFilter] error when try to update taxDocumentSequences by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentSequencesByFilter] failed get rows affected")
	}
	return
}

var (
	taxDocumentSequencesQueries = struct {
		selectTaxDocumentSequences      string
		selectCountTaxDocumentSequences string
		deleteTaxDocumentSequences      string
		updateTaxDocumentSequences      string
		insertTaxDocumentSequences      string
	}{
		selectTaxDocumentSequences:      "SELECT %s FROM \"tax_document_sequences\"",
		selectCountTaxDocumentSequences: "SELECT COUNT(\"sequence_code\") FROM \"tax_document_sequences\"",
		deleteTaxDocumentSequences:      "DELETE FROM \"tax_document_sequences\"",
		updateTaxDocumentSequences:      "UPDATE \"tax_document_sequences\" SET %s ",
		insertTaxDocumentSequences:      "INSERT INTO \"tax_document_sequences\" %s VALUES %s",
	}
)

type TaxDocumentSequencesRepository interface {
	CreateTaxDocumentSequences(ctx context.Context, taxDocumentSequences *model.TaxDocumentSequences, fieldsInsert ...TaxDocumentSequencesField) error
	BulkCreateTaxDocumentSequences(ctx context.Context, taxDocumentSequencesList []*model.TaxDocumentSequences, fieldsInsert ...TaxDocumentSequencesField) error
	ResolveTaxDocumentSequences(ctx context.Context, selectFields ...TaxDocumentSequencesField) (model.TaxDocumentSequencesList, error)
	ResolveTaxDocumentSequencesByID(ctx context.Context, primaryID model.TaxDocumentSequencesPrimaryID, selectFields ...TaxDocumentSequencesField) (model.TaxDocumentSequences, error)
	UpdateTaxDocumentSequencesByID(ctx context.Context, id model.TaxDocumentSequencesPrimaryID, taxDocumentSequences *model.TaxDocumentSequences, taxDocumentSequencesUpdateFields ...TaxDocumentSequencesUpdateField) error
	UpdateTaxDocumentSequencesByFilter(ctx context.Context, filter model.Filter, taxDocumentSequencesUpdateFields ...TaxDocumentSequencesUpdateField) (rowsAffected int64, err error)
	BulkUpdateTaxDocumentSequences(ctx context.Context, taxDocumentSequencesListMap map[model.TaxDocumentSequencesPrimaryID]*model.TaxDocumentSequences, TaxDocumentSequencessMapUpdateFieldsRequest map[model.TaxDocumentSequencesPrimaryID]TaxDocumentSequencesUpdateFieldList) (err error)
	DeleteTaxDocumentSequencesByID(ctx context.Context, id model.TaxDocumentSequencesPrimaryID) error
	BulkDeleteTaxDocumentSequencesByIDs(ctx context.Context, ids []model.TaxDocumentSequencesPrimaryID) error
	ResolveTaxDocumentSequencesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxDocumentSequencesFilterResult, err error)
	IsExistTaxDocumentSequencesByIDs(ctx context.Context, ids []model.TaxDocumentSequencesPrimaryID) (exists bool, notFoundIds []model.TaxDocumentSequencesPrimaryID, err error)
	IsExistTaxDocumentSequencesByID(ctx context.Context, id model.TaxDocumentSequencesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
