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

func composeInsertFieldsAndParamsFinanceParties(financePartiesList []model.FinanceParties, fieldsInsert ...FinancePartiesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinancePartiesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeParties := range financePartiesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeParties.Id)
			case selectField.PartyType():
				args = append(args, financeParties.PartyType)
			case selectField.ExternalRef():
				args = append(args, financeParties.ExternalRef)
			case selectField.LegalName():
				args = append(args, financeParties.LegalName)
			case selectField.DisplayName():
				args = append(args, financeParties.DisplayName)
			case selectField.CountryCode():
				args = append(args, financeParties.CountryCode)
			case selectField.Email():
				args = append(args, financeParties.Email)
			case selectField.Phone():
				args = append(args, financeParties.Phone)
			case selectField.PartyStatus():
				args = append(args, financeParties.PartyStatus)
			case selectField.Metadata():
				args = append(args, financeParties.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeParties.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeParties.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeParties.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeParties.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeParties.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeParties.MetaDeletedBy)

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

func composeFinancePartiesCompositePrimaryKeyWhere(primaryIDs []model.FinancePartiesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_parties\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinancePartiesSelectFields() string {
	fields := NewFinancePartiesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinancePartiesSelectFields(selectFields ...FinancePartiesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinancePartiesField string
type FinancePartiesFieldList []FinancePartiesField

type FinancePartiesSelectFields struct {
}

func (ss FinancePartiesSelectFields) Id() FinancePartiesField {
	return FinancePartiesField("id")
}

func (ss FinancePartiesSelectFields) PartyType() FinancePartiesField {
	return FinancePartiesField("party_type")
}

func (ss FinancePartiesSelectFields) ExternalRef() FinancePartiesField {
	return FinancePartiesField("external_ref")
}

func (ss FinancePartiesSelectFields) LegalName() FinancePartiesField {
	return FinancePartiesField("legal_name")
}

func (ss FinancePartiesSelectFields) DisplayName() FinancePartiesField {
	return FinancePartiesField("display_name")
}

func (ss FinancePartiesSelectFields) CountryCode() FinancePartiesField {
	return FinancePartiesField("country_code")
}

func (ss FinancePartiesSelectFields) Email() FinancePartiesField {
	return FinancePartiesField("email")
}

func (ss FinancePartiesSelectFields) Phone() FinancePartiesField {
	return FinancePartiesField("phone")
}

func (ss FinancePartiesSelectFields) PartyStatus() FinancePartiesField {
	return FinancePartiesField("party_status")
}

func (ss FinancePartiesSelectFields) Metadata() FinancePartiesField {
	return FinancePartiesField("metadata")
}

func (ss FinancePartiesSelectFields) MetaCreatedAt() FinancePartiesField {
	return FinancePartiesField("meta_created_at")
}

func (ss FinancePartiesSelectFields) MetaCreatedBy() FinancePartiesField {
	return FinancePartiesField("meta_created_by")
}

func (ss FinancePartiesSelectFields) MetaUpdatedAt() FinancePartiesField {
	return FinancePartiesField("meta_updated_at")
}

func (ss FinancePartiesSelectFields) MetaUpdatedBy() FinancePartiesField {
	return FinancePartiesField("meta_updated_by")
}

func (ss FinancePartiesSelectFields) MetaDeletedAt() FinancePartiesField {
	return FinancePartiesField("meta_deleted_at")
}

func (ss FinancePartiesSelectFields) MetaDeletedBy() FinancePartiesField {
	return FinancePartiesField("meta_deleted_by")
}

func (ss FinancePartiesSelectFields) All() FinancePartiesFieldList {
	return []FinancePartiesField{
		ss.Id(),
		ss.PartyType(),
		ss.ExternalRef(),
		ss.LegalName(),
		ss.DisplayName(),
		ss.CountryCode(),
		ss.Email(),
		ss.Phone(),
		ss.PartyStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinancePartiesSelectFields() FinancePartiesSelectFields {
	return FinancePartiesSelectFields{}
}

type FinancePartiesUpdateFieldOption struct {
	useIncrement bool
}
type FinancePartiesUpdateField struct {
	financePartiesField FinancePartiesField
	opt                 FinancePartiesUpdateFieldOption
	value               interface{}
}
type FinancePartiesUpdateFieldList []FinancePartiesUpdateField

func defaultFinancePartiesUpdateFieldOption() FinancePartiesUpdateFieldOption {
	return FinancePartiesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinancePartiesOption(useIncrement bool) func(*FinancePartiesUpdateFieldOption) {
	return func(pcufo *FinancePartiesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinancePartiesUpdateField(field FinancePartiesField, val interface{}, opts ...func(*FinancePartiesUpdateFieldOption)) FinancePartiesUpdateField {
	defaultOpt := defaultFinancePartiesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinancePartiesUpdateField{
		financePartiesField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultFinancePartiesUpdateFields(financeParties model.FinanceParties) (financePartiesUpdateFieldList FinancePartiesUpdateFieldList) {
	selectFields := NewFinancePartiesSelectFields()
	financePartiesUpdateFieldList = append(financePartiesUpdateFieldList,
		NewFinancePartiesUpdateField(selectFields.Id(), financeParties.Id),
		NewFinancePartiesUpdateField(selectFields.PartyType(), financeParties.PartyType),
		NewFinancePartiesUpdateField(selectFields.ExternalRef(), financeParties.ExternalRef),
		NewFinancePartiesUpdateField(selectFields.LegalName(), financeParties.LegalName),
		NewFinancePartiesUpdateField(selectFields.DisplayName(), financeParties.DisplayName),
		NewFinancePartiesUpdateField(selectFields.CountryCode(), financeParties.CountryCode),
		NewFinancePartiesUpdateField(selectFields.Email(), financeParties.Email),
		NewFinancePartiesUpdateField(selectFields.Phone(), financeParties.Phone),
		NewFinancePartiesUpdateField(selectFields.PartyStatus(), financeParties.PartyStatus),
		NewFinancePartiesUpdateField(selectFields.Metadata(), financeParties.Metadata),
		NewFinancePartiesUpdateField(selectFields.MetaCreatedAt(), financeParties.MetaCreatedAt),
		NewFinancePartiesUpdateField(selectFields.MetaCreatedBy(), financeParties.MetaCreatedBy),
		NewFinancePartiesUpdateField(selectFields.MetaUpdatedAt(), financeParties.MetaUpdatedAt),
		NewFinancePartiesUpdateField(selectFields.MetaUpdatedBy(), financeParties.MetaUpdatedBy),
		NewFinancePartiesUpdateField(selectFields.MetaDeletedAt(), financeParties.MetaDeletedAt),
		NewFinancePartiesUpdateField(selectFields.MetaDeletedBy(), financeParties.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinancePartiesCommand(financePartiesUpdateFieldList FinancePartiesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financePartiesUpdateFieldList {
		field := string(updateField.financePartiesField)
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

func (repo *RepositoryImpl) BulkCreateFinanceParties(ctx context.Context, financePartiesList []*model.FinanceParties, fieldsInsert ...FinancePartiesField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.FinancePartiesPrimaryID
		financePartiesValueList []model.FinanceParties
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinancePartiesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeParties := range financePartiesList {

		primaryIds = append(primaryIds, financeParties.ToFinancePartiesPrimaryID())

		financePartiesValueList = append(financePartiesValueList, *financeParties)
	}

	_, notFoundIds, err := repo.IsExistFinancePartiesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceParties] failed checking financeParties whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinancePartiesPrimaryID{}
		mapNotFoundIds := map[model.FinancePartiesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeParties", fmt.Sprintf("financeParties with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceParties(financePartiesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financePartiesQueries.insertFinanceParties, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceParties] failed exec create financeParties query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinancePartiesByIDs(ctx context.Context, primaryIDs []model.FinancePartiesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinancePartiesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinancePartiesByIDs] failed checking financeParties whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeParties with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_parties\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financePartiesQueries.deleteFinanceParties + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinancePartiesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinancePartiesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinancePartiesByIDs(ctx context.Context, ids []model.FinancePartiesPrimaryID) (exists bool, notFoundIds []model.FinancePartiesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_parties\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financePartiesQueries.selectFinanceParties, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinancePartiesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinancePartiesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinancePartiesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinancePartiesPrimaryID]bool{}
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

// BulkUpdateFinanceParties is used to bulk update financeParties, by default it will update all field
// if want to update specific field, then fill financePartiessMapUpdateFieldsRequest else please fill financePartiessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceParties(ctx context.Context, financePartiessMap map[model.FinancePartiesPrimaryID]*model.FinanceParties, financePartiessMapUpdateFieldsRequest map[model.FinancePartiesPrimaryID]FinancePartiesUpdateFieldList) (err error) {
	if len(financePartiessMap) == 0 && len(financePartiessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financePartiessMapUpdateField map[model.FinancePartiesPrimaryID]FinancePartiesUpdateFieldList = map[model.FinancePartiesPrimaryID]FinancePartiesUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(financePartiessMap) > 0 {
		for id, financeParties := range financePartiessMap {
			if financeParties == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceParties] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financePartiessMapUpdateField[id] = defaultFinancePartiesUpdateFields(*financeParties)
		}
	} else {
		financePartiessMapUpdateField = financePartiessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinancePartiesQuery(financePartiessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinancePartiesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceParties] failed checking financeParties whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeParties with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinancePartiesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_parties\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceParties] failed exec query")
	}
	return
}

type FinancePartiesFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinancePartiesFieldParameter(param string, args ...interface{}) FinancePartiesFieldParameter {
	return FinancePartiesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinancePartiesQuery(mapFinancePartiess map[model.FinancePartiesPrimaryID]FinancePartiesUpdateFieldList, asTableValues string) (primaryIDs []model.FinancePartiesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinancePartiesPrimaryID]map[string]interface{}{}
	financePartiesSelectFields := NewFinancePartiesSelectFields()
	for id, updateFields := range mapFinancePartiess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financePartiesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinancePartiess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinancePartiesFieldType(updateField.financePartiesField)))
			args = append(args, fields[string(updateField.financePartiesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financePartiesField))
		if updateField.financePartiesField == financePartiesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financePartiesField, asTableValues, updateField.financePartiesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financePartiesField,
				"\"finance_parties\"", updateField.financePartiesField,
				asTableValues, updateField.financePartiesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinancePartiesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinancePartiesPrimaryID, asTableValue string) (whereQry string) {
	financePartiesSelectFields := NewFinancePartiesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_parties\".\"id\" = %s.\"id\"::"+GetFinancePartiesFieldType(financePartiesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinancePartiesFieldType(financePartiesField FinancePartiesField) string {
	selectFinancePartiesFields := NewFinancePartiesSelectFields()
	switch financePartiesField {

	case selectFinancePartiesFields.Id():
		return "uuid"

	case selectFinancePartiesFields.PartyType():
		return "party_type_enum"

	case selectFinancePartiesFields.ExternalRef():
		return "text"

	case selectFinancePartiesFields.LegalName():
		return "text"

	case selectFinancePartiesFields.DisplayName():
		return "text"

	case selectFinancePartiesFields.CountryCode():
		return "text"

	case selectFinancePartiesFields.Email():
		return "text"

	case selectFinancePartiesFields.Phone():
		return "text"

	case selectFinancePartiesFields.PartyStatus():
		return "party_status_enum"

	case selectFinancePartiesFields.Metadata():
		return "jsonb"

	case selectFinancePartiesFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinancePartiesFields.MetaCreatedBy():
		return "uuid"

	case selectFinancePartiesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinancePartiesFields.MetaUpdatedBy():
		return "uuid"

	case selectFinancePartiesFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinancePartiesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceParties(ctx context.Context, financeParties *model.FinanceParties, fieldsInsert ...FinancePartiesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinancePartiesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinancePartiesPrimaryID{
		Id: financeParties.Id,
	}
	exists, err := repo.IsExistFinancePartiesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceParties] failed checking financeParties whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeParties", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceParties([]model.FinanceParties{*financeParties}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financePartiesQueries.insertFinanceParties, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceParties] failed exec create financeParties query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinancePartiesByID(ctx context.Context, primaryID model.FinancePartiesPrimaryID) (err error) {
	exists, err := repo.IsExistFinancePartiesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinancePartiesByID] failed checking financeParties whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeParties with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinancePartiesCompositePrimaryKeyWhere([]model.FinancePartiesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financePartiesQueries.deleteFinanceParties + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinancePartiesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinancePartiesByFilter(ctx context.Context, filter model.Filter) (result []model.FinancePartiesFilterResult, err error) {
	query, args, err := composeFinancePartiesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinancePartiesByFilter] failed compose financeParties filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinancePartiesByFilter] failed get financeParties by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinancePartiesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinancePartiesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinancePartiesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinancePartiesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinancePartiesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 16 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinancePartiesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 16+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["party_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"party_type\"")
			selectedColumns["party_type"] = struct{}{}
		}
		if _, selected := selectedColumns["external_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"external_ref\"")
			selectedColumns["external_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["legal_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"legal_name\"")
			selectedColumns["legal_name"] = struct{}{}
		}
		if _, selected := selectedColumns["display_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"display_name\"")
			selectedColumns["display_name"] = struct{}{}
		}
		if _, selected := selectedColumns["country_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"country_code\"")
			selectedColumns["country_code"] = struct{}{}
		}
		if _, selected := selectedColumns["email"]; !selected {
			selectColumns = append(selectColumns, "base.\"email\"")
			selectedColumns["email"] = struct{}{}
		}
		if _, selected := selectedColumns["phone"]; !selected {
			selectColumns = append(selectColumns, "base.\"phone\"")
			selectedColumns["phone"] = struct{}{}
		}
		if _, selected := selectedColumns["party_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"party_status\"")
			selectedColumns["party_status"] = struct{}{}
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

type financePartiesFilterPlaceholder struct {
	index int
}

func (p *financePartiesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinancePartiesFilterPredicate(filterField model.FilterField, placeholders *financePartiesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinancePartiesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinancePartiesFilterSQLExpr(spec)
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

func composeFinancePartiesFilterGroup(group model.FilterGroup, placeholders *financePartiesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinancePartiesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinancePartiesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinancePartiesFilterWhereQueries(filter model.Filter, placeholders *financePartiesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinancePartiesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinancePartiesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinancePartiesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinancePartiesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinancePartiesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinancePartiesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financePartiesFilterPlaceholder{index: 1}
	whereQueries, err := composeFinancePartiesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinancePartiesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinancePartiesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinancePartiesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_parties\" base%s", strings.Join(selectColumns, ","), composeFinancePartiesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinancePartiesByID(ctx context.Context, primaryID model.FinancePartiesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinancePartiesCompositePrimaryKeyWhere([]model.FinancePartiesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financePartiesQueries.selectCountFinanceParties, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinancePartiesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceParties(ctx context.Context, selectFields ...FinancePartiesField) (financePartiesList model.FinancePartiesList, err error) {
	var (
		defaultFinancePartiesSelectFields = defaultFinancePartiesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinancePartiesSelectFields = composeFinancePartiesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financePartiesQueries.selectFinanceParties, defaultFinancePartiesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financePartiesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceParties] failed get financeParties list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinancePartiesByID(ctx context.Context, primaryID model.FinancePartiesPrimaryID, selectFields ...FinancePartiesField) (financeParties model.FinanceParties, err error) {
	var (
		defaultFinancePartiesSelectFields = defaultFinancePartiesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinancePartiesSelectFields = composeFinancePartiesSelectFields(selectFields...)
	}
	whereQry, params := composeFinancePartiesCompositePrimaryKeyWhere([]model.FinancePartiesPrimaryID{primaryID})
	query := fmt.Sprintf(financePartiesQueries.selectFinanceParties+" WHERE "+whereQry, defaultFinancePartiesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeParties, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeParties with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinancePartiesByID] failed get financeParties")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinancePartiesByID(ctx context.Context, primaryID model.FinancePartiesPrimaryID, financeParties *model.FinanceParties, financePartiesUpdateFields ...FinancePartiesUpdateField) (err error) {
	exists, err := repo.IsExistFinancePartiesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceParties] failed checking financeParties whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeParties with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeParties == nil {
		if len(financePartiesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinancePartiesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeParties = &model.FinanceParties{}
	}
	var (
		defaultFinancePartiesUpdateFields = defaultFinancePartiesUpdateFields(*financeParties)
		tempUpdateField                   FinancePartiesUpdateFieldList
		selectFields                      = NewFinancePartiesSelectFields()
	)
	if len(financePartiesUpdateFields) > 0 {
		for _, updateField := range financePartiesUpdateFields {
			if updateField.financePartiesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinancePartiesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinancePartiesCompositePrimaryKeyWhere([]model.FinancePartiesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinancePartiesCommand(defaultFinancePartiesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financePartiesQueries.updateFinanceParties+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceParties] error when try to update financeParties by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinancePartiesByFilter(ctx context.Context, filter model.Filter, financePartiesUpdateFields ...FinancePartiesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financePartiesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinancePartiesUpdateFieldList
		selectFields = NewFinancePartiesSelectFields()
	)
	for _, updateField := range financePartiesUpdateFields {
		if updateField.financePartiesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinancePartiesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financePartiesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinancePartiesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_parties\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinancePartiesByFilter] error when try to update financeParties by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinancePartiesByFilter] failed get rows affected")
	}
	return
}

var (
	financePartiesQueries = struct {
		selectFinanceParties      string
		selectCountFinanceParties string
		deleteFinanceParties      string
		updateFinanceParties      string
		insertFinanceParties      string
	}{
		selectFinanceParties:      "SELECT %s FROM \"finance_parties\"",
		selectCountFinanceParties: "SELECT COUNT(\"id\") FROM \"finance_parties\"",
		deleteFinanceParties:      "DELETE FROM \"finance_parties\"",
		updateFinanceParties:      "UPDATE \"finance_parties\" SET %s ",
		insertFinanceParties:      "INSERT INTO \"finance_parties\" %s VALUES %s",
	}
)

type FinancePartiesRepository interface {
	CreateFinanceParties(ctx context.Context, financeParties *model.FinanceParties, fieldsInsert ...FinancePartiesField) error
	BulkCreateFinanceParties(ctx context.Context, financePartiesList []*model.FinanceParties, fieldsInsert ...FinancePartiesField) error
	ResolveFinanceParties(ctx context.Context, selectFields ...FinancePartiesField) (model.FinancePartiesList, error)
	ResolveFinancePartiesByID(ctx context.Context, primaryID model.FinancePartiesPrimaryID, selectFields ...FinancePartiesField) (model.FinanceParties, error)
	UpdateFinancePartiesByID(ctx context.Context, id model.FinancePartiesPrimaryID, financeParties *model.FinanceParties, financePartiesUpdateFields ...FinancePartiesUpdateField) error
	UpdateFinancePartiesByFilter(ctx context.Context, filter model.Filter, financePartiesUpdateFields ...FinancePartiesUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceParties(ctx context.Context, financePartiesListMap map[model.FinancePartiesPrimaryID]*model.FinanceParties, FinancePartiessMapUpdateFieldsRequest map[model.FinancePartiesPrimaryID]FinancePartiesUpdateFieldList) (err error)
	DeleteFinancePartiesByID(ctx context.Context, id model.FinancePartiesPrimaryID) error
	BulkDeleteFinancePartiesByIDs(ctx context.Context, ids []model.FinancePartiesPrimaryID) error
	ResolveFinancePartiesByFilter(ctx context.Context, filter model.Filter) (result []model.FinancePartiesFilterResult, err error)
	IsExistFinancePartiesByIDs(ctx context.Context, ids []model.FinancePartiesPrimaryID) (exists bool, notFoundIds []model.FinancePartiesPrimaryID, err error)
	IsExistFinancePartiesByID(ctx context.Context, id model.FinancePartiesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
