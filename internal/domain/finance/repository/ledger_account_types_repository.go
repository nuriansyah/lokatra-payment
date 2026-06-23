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

func composeInsertFieldsAndParamsLedgerAccountTypes(ledgerAccountTypesList []model.LedgerAccountTypes, fieldsInsert ...LedgerAccountTypesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerAccountTypesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerAccountTypes := range ledgerAccountTypesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Code():
				args = append(args, ledgerAccountTypes.Code)
			case selectField.NormalSide():
				args = append(args, ledgerAccountTypes.NormalSide)
			case selectField.Category():
				args = append(args, ledgerAccountTypes.Category)
			case selectField.Description():
				args = append(args, ledgerAccountTypes.Description)
			case selectField.IsControlAccount():
				args = append(args, ledgerAccountTypes.IsControlAccount)
			case selectField.Metadata():
				args = append(args, ledgerAccountTypes.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerAccountTypes.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerAccountTypes.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerAccountTypes.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerAccountTypes.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerAccountTypes.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerAccountTypes.MetaDeletedBy)

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

func composeLedgerAccountTypesCompositePrimaryKeyWhere(primaryIDs []model.LedgerAccountTypesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		code := "\"ledger_account_types\".\"code\" = ?"
		params = append(params, primaryID.Code)
		arrWhereQry = append(arrWhereQry, code)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerAccountTypesSelectFields() string {
	fields := NewLedgerAccountTypesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerAccountTypesSelectFields(selectFields ...LedgerAccountTypesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerAccountTypesField string
type LedgerAccountTypesFieldList []LedgerAccountTypesField

type LedgerAccountTypesSelectFields struct {
}

func (ss LedgerAccountTypesSelectFields) Code() LedgerAccountTypesField {
	return LedgerAccountTypesField("code")
}

func (ss LedgerAccountTypesSelectFields) NormalSide() LedgerAccountTypesField {
	return LedgerAccountTypesField("normal_side")
}

func (ss LedgerAccountTypesSelectFields) Category() LedgerAccountTypesField {
	return LedgerAccountTypesField("category")
}

func (ss LedgerAccountTypesSelectFields) Description() LedgerAccountTypesField {
	return LedgerAccountTypesField("description")
}

func (ss LedgerAccountTypesSelectFields) IsControlAccount() LedgerAccountTypesField {
	return LedgerAccountTypesField("is_control_account")
}

func (ss LedgerAccountTypesSelectFields) Metadata() LedgerAccountTypesField {
	return LedgerAccountTypesField("metadata")
}

func (ss LedgerAccountTypesSelectFields) MetaCreatedAt() LedgerAccountTypesField {
	return LedgerAccountTypesField("meta_created_at")
}

func (ss LedgerAccountTypesSelectFields) MetaCreatedBy() LedgerAccountTypesField {
	return LedgerAccountTypesField("meta_created_by")
}

func (ss LedgerAccountTypesSelectFields) MetaUpdatedAt() LedgerAccountTypesField {
	return LedgerAccountTypesField("meta_updated_at")
}

func (ss LedgerAccountTypesSelectFields) MetaUpdatedBy() LedgerAccountTypesField {
	return LedgerAccountTypesField("meta_updated_by")
}

func (ss LedgerAccountTypesSelectFields) MetaDeletedAt() LedgerAccountTypesField {
	return LedgerAccountTypesField("meta_deleted_at")
}

func (ss LedgerAccountTypesSelectFields) MetaDeletedBy() LedgerAccountTypesField {
	return LedgerAccountTypesField("meta_deleted_by")
}

func (ss LedgerAccountTypesSelectFields) All() LedgerAccountTypesFieldList {
	return []LedgerAccountTypesField{
		ss.Code(),
		ss.NormalSide(),
		ss.Category(),
		ss.Description(),
		ss.IsControlAccount(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerAccountTypesSelectFields() LedgerAccountTypesSelectFields {
	return LedgerAccountTypesSelectFields{}
}

type LedgerAccountTypesUpdateFieldOption struct {
	useIncrement bool
}
type LedgerAccountTypesUpdateField struct {
	ledgerAccountTypesField LedgerAccountTypesField
	opt                     LedgerAccountTypesUpdateFieldOption
	value                   interface{}
}
type LedgerAccountTypesUpdateFieldList []LedgerAccountTypesUpdateField

func defaultLedgerAccountTypesUpdateFieldOption() LedgerAccountTypesUpdateFieldOption {
	return LedgerAccountTypesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerAccountTypesOption(useIncrement bool) func(*LedgerAccountTypesUpdateFieldOption) {
	return func(pcufo *LedgerAccountTypesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerAccountTypesUpdateField(field LedgerAccountTypesField, val interface{}, opts ...func(*LedgerAccountTypesUpdateFieldOption)) LedgerAccountTypesUpdateField {
	defaultOpt := defaultLedgerAccountTypesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerAccountTypesUpdateField{
		ledgerAccountTypesField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultLedgerAccountTypesUpdateFields(ledgerAccountTypes model.LedgerAccountTypes) (ledgerAccountTypesUpdateFieldList LedgerAccountTypesUpdateFieldList) {
	selectFields := NewLedgerAccountTypesSelectFields()
	ledgerAccountTypesUpdateFieldList = append(ledgerAccountTypesUpdateFieldList,
		NewLedgerAccountTypesUpdateField(selectFields.Code(), ledgerAccountTypes.Code),
		NewLedgerAccountTypesUpdateField(selectFields.NormalSide(), ledgerAccountTypes.NormalSide),
		NewLedgerAccountTypesUpdateField(selectFields.Category(), ledgerAccountTypes.Category),
		NewLedgerAccountTypesUpdateField(selectFields.Description(), ledgerAccountTypes.Description),
		NewLedgerAccountTypesUpdateField(selectFields.IsControlAccount(), ledgerAccountTypes.IsControlAccount),
		NewLedgerAccountTypesUpdateField(selectFields.Metadata(), ledgerAccountTypes.Metadata),
		NewLedgerAccountTypesUpdateField(selectFields.MetaCreatedAt(), ledgerAccountTypes.MetaCreatedAt),
		NewLedgerAccountTypesUpdateField(selectFields.MetaCreatedBy(), ledgerAccountTypes.MetaCreatedBy),
		NewLedgerAccountTypesUpdateField(selectFields.MetaUpdatedAt(), ledgerAccountTypes.MetaUpdatedAt),
		NewLedgerAccountTypesUpdateField(selectFields.MetaUpdatedBy(), ledgerAccountTypes.MetaUpdatedBy),
		NewLedgerAccountTypesUpdateField(selectFields.MetaDeletedAt(), ledgerAccountTypes.MetaDeletedAt),
		NewLedgerAccountTypesUpdateField(selectFields.MetaDeletedBy(), ledgerAccountTypes.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerAccountTypesCommand(ledgerAccountTypesUpdateFieldList LedgerAccountTypesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerAccountTypesUpdateFieldList {
		field := string(updateField.ledgerAccountTypesField)
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

func (repo *RepositoryImpl) BulkCreateLedgerAccountTypes(ctx context.Context, ledgerAccountTypesList []*model.LedgerAccountTypes, fieldsInsert ...LedgerAccountTypesField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.LedgerAccountTypesPrimaryID
		ledgerAccountTypesValueList []model.LedgerAccountTypes
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountTypesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerAccountTypes := range ledgerAccountTypesList {

		primaryIds = append(primaryIds, ledgerAccountTypes.ToLedgerAccountTypesPrimaryID())

		ledgerAccountTypesValueList = append(ledgerAccountTypesValueList, *ledgerAccountTypes)
	}

	_, notFoundIds, err := repo.IsExistLedgerAccountTypesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccountTypes] failed checking ledgerAccountTypes whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerAccountTypesPrimaryID{}
		mapNotFoundIds := map[model.LedgerAccountTypesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerAccountTypes", fmt.Sprintf("ledgerAccountTypes with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerAccountTypes(ledgerAccountTypesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerAccountTypesQueries.insertLedgerAccountTypes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccountTypes] failed exec create ledgerAccountTypes query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerAccountTypesByIDs(ctx context.Context, primaryIDs []model.LedgerAccountTypesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerAccountTypesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountTypesByIDs] failed checking ledgerAccountTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountTypes with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_account_types\".\"code\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Code)
	}

	commandQuery := fmt.Sprintf(ledgerAccountTypesQueries.deleteLedgerAccountTypes + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountTypesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountTypesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerAccountTypesByIDs(ctx context.Context, ids []model.LedgerAccountTypesPrimaryID) (exists bool, notFoundIds []model.LedgerAccountTypesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_account_types\".\"code\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Code)
	}

	query := fmt.Sprintf(ledgerAccountTypesQueries.selectLedgerAccountTypes, " \"code\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountTypesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerAccountTypesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountTypesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerAccountTypesPrimaryID]bool{}
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

// BulkUpdateLedgerAccountTypes is used to bulk update ledgerAccountTypes, by default it will update all field
// if want to update specific field, then fill ledgerAccountTypessMapUpdateFieldsRequest else please fill ledgerAccountTypessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerAccountTypes(ctx context.Context, ledgerAccountTypessMap map[model.LedgerAccountTypesPrimaryID]*model.LedgerAccountTypes, ledgerAccountTypessMapUpdateFieldsRequest map[model.LedgerAccountTypesPrimaryID]LedgerAccountTypesUpdateFieldList) (err error) {
	if len(ledgerAccountTypessMap) == 0 && len(ledgerAccountTypessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerAccountTypessMapUpdateField map[model.LedgerAccountTypesPrimaryID]LedgerAccountTypesUpdateFieldList = map[model.LedgerAccountTypesPrimaryID]LedgerAccountTypesUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(ledgerAccountTypessMap) > 0 {
		for id, ledgerAccountTypes := range ledgerAccountTypessMap {
			if ledgerAccountTypes == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerAccountTypes] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerAccountTypessMapUpdateField[id] = defaultLedgerAccountTypesUpdateFields(*ledgerAccountTypes)
		}
	} else {
		ledgerAccountTypessMapUpdateField = ledgerAccountTypessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerAccountTypesQuery(ledgerAccountTypessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerAccountTypesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccountTypes] failed checking ledgerAccountTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountTypes with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerAccountTypesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_account_types\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccountTypes] failed exec query")
	}
	return
}

type LedgerAccountTypesFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerAccountTypesFieldParameter(param string, args ...interface{}) LedgerAccountTypesFieldParameter {
	return LedgerAccountTypesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerAccountTypesQuery(mapLedgerAccountTypess map[model.LedgerAccountTypesPrimaryID]LedgerAccountTypesUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerAccountTypesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerAccountTypesPrimaryID]map[string]interface{}{}
	ledgerAccountTypesSelectFields := NewLedgerAccountTypesSelectFields()
	for id, updateFields := range mapLedgerAccountTypess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerAccountTypesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerAccountTypess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerAccountTypesFieldType(updateField.ledgerAccountTypesField)))
			args = append(args, fields[string(updateField.ledgerAccountTypesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerAccountTypesField))
		if updateField.ledgerAccountTypesField == ledgerAccountTypesSelectFields.Code() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerAccountTypesField, asTableValues, updateField.ledgerAccountTypesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerAccountTypesField,
				"\"ledger_account_types\"", updateField.ledgerAccountTypesField,
				asTableValues, updateField.ledgerAccountTypesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerAccountTypesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerAccountTypesPrimaryID, asTableValue string) (whereQry string) {
	ledgerAccountTypesSelectFields := NewLedgerAccountTypesSelectFields()
	var arrWhereQry []string
	code := fmt.Sprintf("\"ledger_account_types\".\"code\" = %s.\"code\"::"+GetLedgerAccountTypesFieldType(ledgerAccountTypesSelectFields.Code()), asTableValue)
	arrWhereQry = append(arrWhereQry, code)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerAccountTypesFieldType(ledgerAccountTypesField LedgerAccountTypesField) string {
	selectLedgerAccountTypesFields := NewLedgerAccountTypesSelectFields()
	switch ledgerAccountTypesField {

	case selectLedgerAccountTypesFields.Code():
		return "text"

	case selectLedgerAccountTypesFields.NormalSide():
		return "normal_side_enum"

	case selectLedgerAccountTypesFields.Category():
		return "category_enum"

	case selectLedgerAccountTypesFields.Description():
		return "text"

	case selectLedgerAccountTypesFields.IsControlAccount():
		return "bool"

	case selectLedgerAccountTypesFields.Metadata():
		return "jsonb"

	case selectLedgerAccountTypesFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerAccountTypesFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerAccountTypesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerAccountTypesFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerAccountTypesFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerAccountTypesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerAccountTypes(ctx context.Context, ledgerAccountTypes *model.LedgerAccountTypes, fieldsInsert ...LedgerAccountTypesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountTypesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerAccountTypesPrimaryID{
		Code: ledgerAccountTypes.Code,
	}
	exists, err := repo.IsExistLedgerAccountTypesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccountTypes] failed checking ledgerAccountTypes whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerAccountTypes", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerAccountTypes([]model.LedgerAccountTypes{*ledgerAccountTypes}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerAccountTypesQueries.insertLedgerAccountTypes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccountTypes] failed exec create ledgerAccountTypes query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerAccountTypesByID(ctx context.Context, primaryID model.LedgerAccountTypesPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerAccountTypesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountTypesByID] failed checking ledgerAccountTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountTypes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerAccountTypesCompositePrimaryKeyWhere([]model.LedgerAccountTypesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerAccountTypesQueries.deleteLedgerAccountTypes + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountTypesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountTypesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountTypesFilterResult, err error) {
	query, args, err := composeLedgerAccountTypesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountTypesByFilter] failed compose ledgerAccountTypes filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountTypesByFilter] failed get ledgerAccountTypes by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerAccountTypesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerAccountTypesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerAccountTypesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerAccountTypesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerAccountTypesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 12 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerAccountTypesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 12+1)
		if _, selected := selectedColumns["code"]; !selected {
			selectColumns = append(selectColumns, "base.\"code\"")
			selectedColumns["code"] = struct{}{}
		}
		if _, selected := selectedColumns["normal_side"]; !selected {
			selectColumns = append(selectColumns, "base.\"normal_side\"")
			selectedColumns["normal_side"] = struct{}{}
		}
		if _, selected := selectedColumns["category"]; !selected {
			selectColumns = append(selectColumns, "base.\"category\"")
			selectedColumns["category"] = struct{}{}
		}
		if _, selected := selectedColumns["description"]; !selected {
			selectColumns = append(selectColumns, "base.\"description\"")
			selectedColumns["description"] = struct{}{}
		}
		if _, selected := selectedColumns["is_control_account"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_control_account\"")
			selectedColumns["is_control_account"] = struct{}{}
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

type ledgerAccountTypesFilterPlaceholder struct {
	index int
}

func (p *ledgerAccountTypesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerAccountTypesFilterPredicate(filterField model.FilterField, placeholders *ledgerAccountTypesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerAccountTypesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerAccountTypesFilterSQLExpr(spec)
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

func composeLedgerAccountTypesFilterGroup(group model.FilterGroup, placeholders *ledgerAccountTypesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerAccountTypesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerAccountTypesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerAccountTypesFilterWhereQueries(filter model.Filter, placeholders *ledgerAccountTypesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerAccountTypesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerAccountTypesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerAccountTypesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerAccountTypesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerAccountTypesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerAccountTypesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerAccountTypesFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerAccountTypesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerAccountTypesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerAccountTypesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerAccountTypesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_account_types\" base%s", strings.Join(selectColumns, ","), composeLedgerAccountTypesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerAccountTypesByID(ctx context.Context, primaryID model.LedgerAccountTypesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerAccountTypesCompositePrimaryKeyWhere([]model.LedgerAccountTypesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerAccountTypesQueries.selectCountLedgerAccountTypes, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountTypesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountTypes(ctx context.Context, selectFields ...LedgerAccountTypesField) (ledgerAccountTypesList model.LedgerAccountTypesList, err error) {
	var (
		defaultLedgerAccountTypesSelectFields = defaultLedgerAccountTypesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountTypesSelectFields = composeLedgerAccountTypesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerAccountTypesQueries.selectLedgerAccountTypes, defaultLedgerAccountTypesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerAccountTypesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountTypes] failed get ledgerAccountTypes list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountTypesByID(ctx context.Context, primaryID model.LedgerAccountTypesPrimaryID, selectFields ...LedgerAccountTypesField) (ledgerAccountTypes model.LedgerAccountTypes, err error) {
	var (
		defaultLedgerAccountTypesSelectFields = defaultLedgerAccountTypesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountTypesSelectFields = composeLedgerAccountTypesSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerAccountTypesCompositePrimaryKeyWhere([]model.LedgerAccountTypesPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerAccountTypesQueries.selectLedgerAccountTypes+" WHERE "+whereQry, defaultLedgerAccountTypesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerAccountTypes, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerAccountTypes with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerAccountTypesByID] failed get ledgerAccountTypes")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerAccountTypesByID(ctx context.Context, primaryID model.LedgerAccountTypesPrimaryID, ledgerAccountTypes *model.LedgerAccountTypes, ledgerAccountTypesUpdateFields ...LedgerAccountTypesUpdateField) (err error) {
	exists, err := repo.IsExistLedgerAccountTypesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountTypes] failed checking ledgerAccountTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountTypes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerAccountTypes == nil {
		if len(ledgerAccountTypesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerAccountTypesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerAccountTypes = &model.LedgerAccountTypes{}
	}
	var (
		defaultLedgerAccountTypesUpdateFields = defaultLedgerAccountTypesUpdateFields(*ledgerAccountTypes)
		tempUpdateField                       LedgerAccountTypesUpdateFieldList
		selectFields                          = NewLedgerAccountTypesSelectFields()
	)
	if len(ledgerAccountTypesUpdateFields) > 0 {
		for _, updateField := range ledgerAccountTypesUpdateFields {
			if updateField.ledgerAccountTypesField == selectFields.Code() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerAccountTypesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerAccountTypesCompositePrimaryKeyWhere([]model.LedgerAccountTypesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerAccountTypesCommand(defaultLedgerAccountTypesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerAccountTypesQueries.updateLedgerAccountTypes+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountTypes] error when try to update ledgerAccountTypes by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerAccountTypesByFilter(ctx context.Context, filter model.Filter, ledgerAccountTypesUpdateFields ...LedgerAccountTypesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerAccountTypesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerAccountTypesUpdateFieldList
		selectFields = NewLedgerAccountTypesSelectFields()
	)
	for _, updateField := range ledgerAccountTypesUpdateFields {
		if updateField.ledgerAccountTypesField == selectFields.Code() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerAccountTypesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerAccountTypesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerAccountTypesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_account_types\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountTypesByFilter] error when try to update ledgerAccountTypes by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountTypesByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerAccountTypesQueries = struct {
		selectLedgerAccountTypes      string
		selectCountLedgerAccountTypes string
		deleteLedgerAccountTypes      string
		updateLedgerAccountTypes      string
		insertLedgerAccountTypes      string
	}{
		selectLedgerAccountTypes:      "SELECT %s FROM \"ledger_account_types\"",
		selectCountLedgerAccountTypes: "SELECT COUNT(\"code\") FROM \"ledger_account_types\"",
		deleteLedgerAccountTypes:      "DELETE FROM \"ledger_account_types\"",
		updateLedgerAccountTypes:      "UPDATE \"ledger_account_types\" SET %s ",
		insertLedgerAccountTypes:      "INSERT INTO \"ledger_account_types\" %s VALUES %s",
	}
)

type LedgerAccountTypesRepository interface {
	CreateLedgerAccountTypes(ctx context.Context, ledgerAccountTypes *model.LedgerAccountTypes, fieldsInsert ...LedgerAccountTypesField) error
	BulkCreateLedgerAccountTypes(ctx context.Context, ledgerAccountTypesList []*model.LedgerAccountTypes, fieldsInsert ...LedgerAccountTypesField) error
	ResolveLedgerAccountTypes(ctx context.Context, selectFields ...LedgerAccountTypesField) (model.LedgerAccountTypesList, error)
	ResolveLedgerAccountTypesByID(ctx context.Context, primaryID model.LedgerAccountTypesPrimaryID, selectFields ...LedgerAccountTypesField) (model.LedgerAccountTypes, error)
	UpdateLedgerAccountTypesByID(ctx context.Context, id model.LedgerAccountTypesPrimaryID, ledgerAccountTypes *model.LedgerAccountTypes, ledgerAccountTypesUpdateFields ...LedgerAccountTypesUpdateField) error
	UpdateLedgerAccountTypesByFilter(ctx context.Context, filter model.Filter, ledgerAccountTypesUpdateFields ...LedgerAccountTypesUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerAccountTypes(ctx context.Context, ledgerAccountTypesListMap map[model.LedgerAccountTypesPrimaryID]*model.LedgerAccountTypes, LedgerAccountTypessMapUpdateFieldsRequest map[model.LedgerAccountTypesPrimaryID]LedgerAccountTypesUpdateFieldList) (err error)
	DeleteLedgerAccountTypesByID(ctx context.Context, id model.LedgerAccountTypesPrimaryID) error
	BulkDeleteLedgerAccountTypesByIDs(ctx context.Context, ids []model.LedgerAccountTypesPrimaryID) error
	ResolveLedgerAccountTypesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountTypesFilterResult, err error)
	IsExistLedgerAccountTypesByIDs(ctx context.Context, ids []model.LedgerAccountTypesPrimaryID) (exists bool, notFoundIds []model.LedgerAccountTypesPrimaryID, err error)
	IsExistLedgerAccountTypesByID(ctx context.Context, id model.LedgerAccountTypesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
