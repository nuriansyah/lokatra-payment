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

func composeInsertFieldsAndParamsFinanceCurrencies(financeCurrenciesList []model.FinanceCurrencies, fieldsInsert ...FinanceCurrenciesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceCurrenciesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeCurrencies := range financeCurrenciesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Code():
				args = append(args, financeCurrencies.Code)
			case selectField.DecimalCode():
				args = append(args, financeCurrencies.DecimalCode)
			case selectField.Exponent():
				args = append(args, financeCurrencies.Exponent)
			case selectField.IsActive():
				args = append(args, financeCurrencies.IsActive)
			case selectField.Metadata():
				args = append(args, financeCurrencies.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeCurrencies.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeCurrencies.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeCurrencies.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeCurrencies.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeCurrencies.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeCurrencies.MetaDeletedBy)

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

func composeFinanceCurrenciesCompositePrimaryKeyWhere(primaryIDs []model.FinanceCurrenciesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		code := "\"finance_currencies\".\"code\" = ?"
		params = append(params, primaryID.Code)
		arrWhereQry = append(arrWhereQry, code)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceCurrenciesSelectFields() string {
	fields := NewFinanceCurrenciesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceCurrenciesSelectFields(selectFields ...FinanceCurrenciesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceCurrenciesField string
type FinanceCurrenciesFieldList []FinanceCurrenciesField

type FinanceCurrenciesSelectFields struct {
}

func (ss FinanceCurrenciesSelectFields) Code() FinanceCurrenciesField {
	return FinanceCurrenciesField("code")
}

func (ss FinanceCurrenciesSelectFields) DecimalCode() FinanceCurrenciesField {
	return FinanceCurrenciesField("decimal_code")
}

func (ss FinanceCurrenciesSelectFields) Exponent() FinanceCurrenciesField {
	return FinanceCurrenciesField("exponent")
}

func (ss FinanceCurrenciesSelectFields) IsActive() FinanceCurrenciesField {
	return FinanceCurrenciesField("is_active")
}

func (ss FinanceCurrenciesSelectFields) Metadata() FinanceCurrenciesField {
	return FinanceCurrenciesField("metadata")
}

func (ss FinanceCurrenciesSelectFields) MetaCreatedAt() FinanceCurrenciesField {
	return FinanceCurrenciesField("meta_created_at")
}

func (ss FinanceCurrenciesSelectFields) MetaCreatedBy() FinanceCurrenciesField {
	return FinanceCurrenciesField("meta_created_by")
}

func (ss FinanceCurrenciesSelectFields) MetaUpdatedAt() FinanceCurrenciesField {
	return FinanceCurrenciesField("meta_updated_at")
}

func (ss FinanceCurrenciesSelectFields) MetaUpdatedBy() FinanceCurrenciesField {
	return FinanceCurrenciesField("meta_updated_by")
}

func (ss FinanceCurrenciesSelectFields) MetaDeletedAt() FinanceCurrenciesField {
	return FinanceCurrenciesField("meta_deleted_at")
}

func (ss FinanceCurrenciesSelectFields) MetaDeletedBy() FinanceCurrenciesField {
	return FinanceCurrenciesField("meta_deleted_by")
}

func (ss FinanceCurrenciesSelectFields) All() FinanceCurrenciesFieldList {
	return []FinanceCurrenciesField{
		ss.Code(),
		ss.DecimalCode(),
		ss.Exponent(),
		ss.IsActive(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinanceCurrenciesSelectFields() FinanceCurrenciesSelectFields {
	return FinanceCurrenciesSelectFields{}
}

type FinanceCurrenciesUpdateFieldOption struct {
	useIncrement bool
}
type FinanceCurrenciesUpdateField struct {
	financeCurrenciesField FinanceCurrenciesField
	opt                    FinanceCurrenciesUpdateFieldOption
	value                  interface{}
}
type FinanceCurrenciesUpdateFieldList []FinanceCurrenciesUpdateField

func defaultFinanceCurrenciesUpdateFieldOption() FinanceCurrenciesUpdateFieldOption {
	return FinanceCurrenciesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceCurrenciesOption(useIncrement bool) func(*FinanceCurrenciesUpdateFieldOption) {
	return func(pcufo *FinanceCurrenciesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceCurrenciesUpdateField(field FinanceCurrenciesField, val interface{}, opts ...func(*FinanceCurrenciesUpdateFieldOption)) FinanceCurrenciesUpdateField {
	defaultOpt := defaultFinanceCurrenciesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceCurrenciesUpdateField{
		financeCurrenciesField: field,
		value:                  val,
		opt:                    defaultOpt,
	}
}
func defaultFinanceCurrenciesUpdateFields(financeCurrencies model.FinanceCurrencies) (financeCurrenciesUpdateFieldList FinanceCurrenciesUpdateFieldList) {
	selectFields := NewFinanceCurrenciesSelectFields()
	financeCurrenciesUpdateFieldList = append(financeCurrenciesUpdateFieldList,
		NewFinanceCurrenciesUpdateField(selectFields.Code(), financeCurrencies.Code),
		NewFinanceCurrenciesUpdateField(selectFields.DecimalCode(), financeCurrencies.DecimalCode),
		NewFinanceCurrenciesUpdateField(selectFields.Exponent(), financeCurrencies.Exponent),
		NewFinanceCurrenciesUpdateField(selectFields.IsActive(), financeCurrencies.IsActive),
		NewFinanceCurrenciesUpdateField(selectFields.Metadata(), financeCurrencies.Metadata),
		NewFinanceCurrenciesUpdateField(selectFields.MetaCreatedAt(), financeCurrencies.MetaCreatedAt),
		NewFinanceCurrenciesUpdateField(selectFields.MetaCreatedBy(), financeCurrencies.MetaCreatedBy),
		NewFinanceCurrenciesUpdateField(selectFields.MetaUpdatedAt(), financeCurrencies.MetaUpdatedAt),
		NewFinanceCurrenciesUpdateField(selectFields.MetaUpdatedBy(), financeCurrencies.MetaUpdatedBy),
		NewFinanceCurrenciesUpdateField(selectFields.MetaDeletedAt(), financeCurrencies.MetaDeletedAt),
		NewFinanceCurrenciesUpdateField(selectFields.MetaDeletedBy(), financeCurrencies.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceCurrenciesCommand(financeCurrenciesUpdateFieldList FinanceCurrenciesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeCurrenciesUpdateFieldList {
		field := string(updateField.financeCurrenciesField)
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

func (repo *RepositoryImpl) BulkCreateFinanceCurrencies(ctx context.Context, financeCurrenciesList []*model.FinanceCurrencies, fieldsInsert ...FinanceCurrenciesField) (err error) {
	var (
		fieldsStr                  string
		valueListStr               []string
		argsList                   []interface{}
		primaryIds                 []model.FinanceCurrenciesPrimaryID
		financeCurrenciesValueList []model.FinanceCurrencies
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceCurrenciesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeCurrencies := range financeCurrenciesList {

		primaryIds = append(primaryIds, financeCurrencies.ToFinanceCurrenciesPrimaryID())

		financeCurrenciesValueList = append(financeCurrenciesValueList, *financeCurrencies)
	}

	_, notFoundIds, err := repo.IsExistFinanceCurrenciesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceCurrencies] failed checking financeCurrencies whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceCurrenciesPrimaryID{}
		mapNotFoundIds := map[model.FinanceCurrenciesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeCurrencies", fmt.Sprintf("financeCurrencies with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceCurrencies(financeCurrenciesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeCurrenciesQueries.insertFinanceCurrencies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceCurrencies] failed exec create financeCurrencies query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceCurrenciesByIDs(ctx context.Context, primaryIDs []model.FinanceCurrenciesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceCurrenciesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceCurrenciesByIDs] failed checking financeCurrencies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeCurrencies with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_currencies\".\"code\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Code)
	}

	commandQuery := fmt.Sprintf(financeCurrenciesQueries.deleteFinanceCurrencies + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceCurrenciesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceCurrenciesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceCurrenciesByIDs(ctx context.Context, ids []model.FinanceCurrenciesPrimaryID) (exists bool, notFoundIds []model.FinanceCurrenciesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_currencies\".\"code\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Code)
	}

	query := fmt.Sprintf(financeCurrenciesQueries.selectFinanceCurrencies, " \"code\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceCurrenciesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceCurrenciesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceCurrenciesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceCurrenciesPrimaryID]bool{}
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

// BulkUpdateFinanceCurrencies is used to bulk update financeCurrencies, by default it will update all field
// if want to update specific field, then fill financeCurrenciessMapUpdateFieldsRequest else please fill financeCurrenciessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceCurrencies(ctx context.Context, financeCurrenciessMap map[model.FinanceCurrenciesPrimaryID]*model.FinanceCurrencies, financeCurrenciessMapUpdateFieldsRequest map[model.FinanceCurrenciesPrimaryID]FinanceCurrenciesUpdateFieldList) (err error) {
	if len(financeCurrenciessMap) == 0 && len(financeCurrenciessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeCurrenciessMapUpdateField map[model.FinanceCurrenciesPrimaryID]FinanceCurrenciesUpdateFieldList = map[model.FinanceCurrenciesPrimaryID]FinanceCurrenciesUpdateFieldList{}
		asTableValues                    string                                                                = "myvalues"
	)

	if len(financeCurrenciessMap) > 0 {
		for id, financeCurrencies := range financeCurrenciessMap {
			if financeCurrencies == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceCurrencies] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeCurrenciessMapUpdateField[id] = defaultFinanceCurrenciesUpdateFields(*financeCurrencies)
		}
	} else {
		financeCurrenciessMapUpdateField = financeCurrenciessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceCurrenciesQuery(financeCurrenciessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceCurrenciesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceCurrencies] failed checking financeCurrencies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeCurrencies with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceCurrenciesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_currencies\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceCurrencies] failed exec query")
	}
	return
}

type FinanceCurrenciesFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceCurrenciesFieldParameter(param string, args ...interface{}) FinanceCurrenciesFieldParameter {
	return FinanceCurrenciesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceCurrenciesQuery(mapFinanceCurrenciess map[model.FinanceCurrenciesPrimaryID]FinanceCurrenciesUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceCurrenciesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceCurrenciesPrimaryID]map[string]interface{}{}
	financeCurrenciesSelectFields := NewFinanceCurrenciesSelectFields()
	for id, updateFields := range mapFinanceCurrenciess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeCurrenciesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceCurrenciess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceCurrenciesFieldType(updateField.financeCurrenciesField)))
			args = append(args, fields[string(updateField.financeCurrenciesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeCurrenciesField))
		if updateField.financeCurrenciesField == financeCurrenciesSelectFields.Code() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeCurrenciesField, asTableValues, updateField.financeCurrenciesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeCurrenciesField,
				"\"finance_currencies\"", updateField.financeCurrenciesField,
				asTableValues, updateField.financeCurrenciesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceCurrenciesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceCurrenciesPrimaryID, asTableValue string) (whereQry string) {
	financeCurrenciesSelectFields := NewFinanceCurrenciesSelectFields()
	var arrWhereQry []string
	code := fmt.Sprintf("\"finance_currencies\".\"code\" = %s.\"code\"::"+GetFinanceCurrenciesFieldType(financeCurrenciesSelectFields.Code()), asTableValue)
	arrWhereQry = append(arrWhereQry, code)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceCurrenciesFieldType(financeCurrenciesField FinanceCurrenciesField) string {
	selectFinanceCurrenciesFields := NewFinanceCurrenciesSelectFields()
	switch financeCurrenciesField {

	case selectFinanceCurrenciesFields.Code():
		return "text"

	case selectFinanceCurrenciesFields.DecimalCode():
		return "int2"

	case selectFinanceCurrenciesFields.Exponent():
		return "int2"

	case selectFinanceCurrenciesFields.IsActive():
		return "bool"

	case selectFinanceCurrenciesFields.Metadata():
		return "jsonb"

	case selectFinanceCurrenciesFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceCurrenciesFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceCurrenciesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceCurrenciesFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceCurrenciesFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceCurrenciesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceCurrencies(ctx context.Context, financeCurrencies *model.FinanceCurrencies, fieldsInsert ...FinanceCurrenciesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceCurrenciesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceCurrenciesPrimaryID{
		Code: financeCurrencies.Code,
	}
	exists, err := repo.IsExistFinanceCurrenciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceCurrencies] failed checking financeCurrencies whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeCurrencies", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceCurrencies([]model.FinanceCurrencies{*financeCurrencies}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeCurrenciesQueries.insertFinanceCurrencies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceCurrencies] failed exec create financeCurrencies query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceCurrenciesByID(ctx context.Context, primaryID model.FinanceCurrenciesPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceCurrenciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceCurrenciesByID] failed checking financeCurrencies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeCurrencies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceCurrenciesCompositePrimaryKeyWhere([]model.FinanceCurrenciesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeCurrenciesQueries.deleteFinanceCurrencies + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceCurrenciesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceCurrenciesByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceCurrenciesFilterResult, err error) {
	query, args, err := composeFinanceCurrenciesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceCurrenciesByFilter] failed compose financeCurrencies filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceCurrenciesByFilter] failed get financeCurrencies by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceCurrenciesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceCurrenciesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceCurrenciesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceCurrenciesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceCurrenciesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 11 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceCurrenciesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 11+1)
		if _, selected := selectedColumns["code"]; !selected {
			selectColumns = append(selectColumns, "base.\"code\"")
			selectedColumns["code"] = struct{}{}
		}
		if _, selected := selectedColumns["decimal_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"decimal_code\"")
			selectedColumns["decimal_code"] = struct{}{}
		}
		if _, selected := selectedColumns["exponent"]; !selected {
			selectColumns = append(selectColumns, "base.\"exponent\"")
			selectedColumns["exponent"] = struct{}{}
		}
		if _, selected := selectedColumns["is_active"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_active\"")
			selectedColumns["is_active"] = struct{}{}
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

	if _, selected := selectedColumns["code"]; isCursorMode && !selected {
		selectColumns = append(selectColumns, "base.\"code\"")
		selectedColumns["code"] = struct{}{}
	}

	return
}

type financeCurrenciesFilterPlaceholder struct {
	index int
}

func (p *financeCurrenciesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceCurrenciesFilterPredicate(filterField model.FilterField, placeholders *financeCurrenciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceCurrenciesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceCurrenciesFilterSQLExpr(spec)
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

func composeFinanceCurrenciesFilterGroup(group model.FilterGroup, placeholders *financeCurrenciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceCurrenciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceCurrenciesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceCurrenciesFilterWhereQueries(filter model.Filter, placeholders *financeCurrenciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceCurrenciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceCurrenciesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceCurrenciesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceCurrenciesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceCurrenciesSortOrder(filter.Sorts[0].Order)
			if sortErr != nil {
				err = sortErr
				return
			}
			if filter.Sorts[0].Field != "code" || sortOrder != model.SortAsc {
				err = failure.BadRequestFromString("cursor pagination only supports the default primary-key sort")
				return
			}
		}
	}

	selectColumns, err := composeFinanceCurrenciesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeCurrenciesFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceCurrenciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
	if err != nil {
		return
	}

	if isCursorMode && filter.Pagination.Cursor != nil {
		whereQueries = append(whereQueries, fmt.Sprintf("base.\"code\" %s %s", cursorOperator, placeholders.Next()))
		args = append(args, filter.Pagination.Cursor)
	}

	sortQuery := []string{}
	if isCursorMode {
		sortQuery = append(sortQuery, fmt.Sprintf("base.\"code\" %s", cursorSortOrder))

	} else {
		for _, sort := range filter.Sorts {
			spec, found := model.NewFinanceCurrenciesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceCurrenciesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceCurrenciesSortOrder(sort.Order)
			if sortErr != nil {
				err = sortErr
				return
			}
			sortQuery = append(sortQuery, fmt.Sprintf("%s %s", sqlExpr, sortOrder))
		}
		if len(sortQuery) == 0 {
			sortQuery = append(sortQuery, "base.\"code\" ASC")
		}

	}

	query = fmt.Sprintf("SELECT %s FROM \"finance_currencies\" base%s", strings.Join(selectColumns, ","), composeFinanceCurrenciesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceCurrenciesByID(ctx context.Context, primaryID model.FinanceCurrenciesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceCurrenciesCompositePrimaryKeyWhere([]model.FinanceCurrenciesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeCurrenciesQueries.selectCountFinanceCurrencies, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceCurrenciesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceCurrencies(ctx context.Context, selectFields ...FinanceCurrenciesField) (financeCurrenciesList model.FinanceCurrenciesList, err error) {
	var (
		defaultFinanceCurrenciesSelectFields = defaultFinanceCurrenciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceCurrenciesSelectFields = composeFinanceCurrenciesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeCurrenciesQueries.selectFinanceCurrencies, defaultFinanceCurrenciesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeCurrenciesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceCurrencies] failed get financeCurrencies list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceCurrenciesByID(ctx context.Context, primaryID model.FinanceCurrenciesPrimaryID, selectFields ...FinanceCurrenciesField) (financeCurrencies model.FinanceCurrencies, err error) {
	var (
		defaultFinanceCurrenciesSelectFields = defaultFinanceCurrenciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceCurrenciesSelectFields = composeFinanceCurrenciesSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceCurrenciesCompositePrimaryKeyWhere([]model.FinanceCurrenciesPrimaryID{primaryID})
	query := fmt.Sprintf(financeCurrenciesQueries.selectFinanceCurrencies+" WHERE "+whereQry, defaultFinanceCurrenciesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeCurrencies, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeCurrencies with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceCurrenciesByID] failed get financeCurrencies")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceCurrenciesByID(ctx context.Context, primaryID model.FinanceCurrenciesPrimaryID, financeCurrencies *model.FinanceCurrencies, financeCurrenciesUpdateFields ...FinanceCurrenciesUpdateField) (err error) {
	exists, err := repo.IsExistFinanceCurrenciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceCurrencies] failed checking financeCurrencies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeCurrencies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeCurrencies == nil {
		if len(financeCurrenciesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceCurrenciesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeCurrencies = &model.FinanceCurrencies{}
	}
	var (
		defaultFinanceCurrenciesUpdateFields = defaultFinanceCurrenciesUpdateFields(*financeCurrencies)
		tempUpdateField                      FinanceCurrenciesUpdateFieldList
		selectFields                         = NewFinanceCurrenciesSelectFields()
	)
	if len(financeCurrenciesUpdateFields) > 0 {
		for _, updateField := range financeCurrenciesUpdateFields {
			if updateField.financeCurrenciesField == selectFields.Code() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceCurrenciesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceCurrenciesCompositePrimaryKeyWhere([]model.FinanceCurrenciesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceCurrenciesCommand(defaultFinanceCurrenciesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeCurrenciesQueries.updateFinanceCurrencies+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceCurrencies] error when try to update financeCurrencies by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceCurrenciesByFilter(ctx context.Context, filter model.Filter, financeCurrenciesUpdateFields ...FinanceCurrenciesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeCurrenciesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceCurrenciesUpdateFieldList
		selectFields = NewFinanceCurrenciesSelectFields()
	)
	for _, updateField := range financeCurrenciesUpdateFields {
		if updateField.financeCurrenciesField == selectFields.Code() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceCurrenciesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeCurrenciesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceCurrenciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_currencies\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceCurrenciesByFilter] error when try to update financeCurrencies by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceCurrenciesByFilter] failed get rows affected")
	}
	return
}

var (
	financeCurrenciesQueries = struct {
		selectFinanceCurrencies      string
		selectCountFinanceCurrencies string
		deleteFinanceCurrencies      string
		updateFinanceCurrencies      string
		insertFinanceCurrencies      string
	}{
		selectFinanceCurrencies:      "SELECT %s FROM \"finance_currencies\"",
		selectCountFinanceCurrencies: "SELECT COUNT(\"code\") FROM \"finance_currencies\"",
		deleteFinanceCurrencies:      "DELETE FROM \"finance_currencies\"",
		updateFinanceCurrencies:      "UPDATE \"finance_currencies\" SET %s ",
		insertFinanceCurrencies:      "INSERT INTO \"finance_currencies\" %s VALUES %s",
	}
)

type FinanceCurrenciesRepository interface {
	CreateFinanceCurrencies(ctx context.Context, financeCurrencies *model.FinanceCurrencies, fieldsInsert ...FinanceCurrenciesField) error
	BulkCreateFinanceCurrencies(ctx context.Context, financeCurrenciesList []*model.FinanceCurrencies, fieldsInsert ...FinanceCurrenciesField) error
	ResolveFinanceCurrencies(ctx context.Context, selectFields ...FinanceCurrenciesField) (model.FinanceCurrenciesList, error)
	ResolveFinanceCurrenciesByID(ctx context.Context, primaryID model.FinanceCurrenciesPrimaryID, selectFields ...FinanceCurrenciesField) (model.FinanceCurrencies, error)
	UpdateFinanceCurrenciesByID(ctx context.Context, id model.FinanceCurrenciesPrimaryID, financeCurrencies *model.FinanceCurrencies, financeCurrenciesUpdateFields ...FinanceCurrenciesUpdateField) error
	UpdateFinanceCurrenciesByFilter(ctx context.Context, filter model.Filter, financeCurrenciesUpdateFields ...FinanceCurrenciesUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceCurrencies(ctx context.Context, financeCurrenciesListMap map[model.FinanceCurrenciesPrimaryID]*model.FinanceCurrencies, FinanceCurrenciessMapUpdateFieldsRequest map[model.FinanceCurrenciesPrimaryID]FinanceCurrenciesUpdateFieldList) (err error)
	DeleteFinanceCurrenciesByID(ctx context.Context, id model.FinanceCurrenciesPrimaryID) error
	BulkDeleteFinanceCurrenciesByIDs(ctx context.Context, ids []model.FinanceCurrenciesPrimaryID) error
	ResolveFinanceCurrenciesByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceCurrenciesFilterResult, err error)
	IsExistFinanceCurrenciesByIDs(ctx context.Context, ids []model.FinanceCurrenciesPrimaryID) (exists bool, notFoundIds []model.FinanceCurrenciesPrimaryID, err error)
	IsExistFinanceCurrenciesByID(ctx context.Context, id model.FinanceCurrenciesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
