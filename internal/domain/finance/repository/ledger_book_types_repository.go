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

func composeInsertFieldsAndParamsLedgerBookTypes(ledgerBookTypesList []model.LedgerBookTypes, fieldsInsert ...LedgerBookTypesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerBookTypesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerBookTypes := range ledgerBookTypesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Code():
				args = append(args, ledgerBookTypes.Code)
			case selectField.Description():
				args = append(args, ledgerBookTypes.Description)
			case selectField.Metadata():
				args = append(args, ledgerBookTypes.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerBookTypes.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerBookTypes.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerBookTypes.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerBookTypes.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerBookTypes.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerBookTypes.MetaDeletedBy)

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

func composeLedgerBookTypesCompositePrimaryKeyWhere(primaryIDs []model.LedgerBookTypesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		code := "\"ledger_book_types\".\"code\" = ?"
		params = append(params, primaryID.Code)
		arrWhereQry = append(arrWhereQry, code)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerBookTypesSelectFields() string {
	fields := NewLedgerBookTypesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerBookTypesSelectFields(selectFields ...LedgerBookTypesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerBookTypesField string
type LedgerBookTypesFieldList []LedgerBookTypesField

type LedgerBookTypesSelectFields struct {
}

func (ss LedgerBookTypesSelectFields) Code() LedgerBookTypesField {
	return LedgerBookTypesField("code")
}

func (ss LedgerBookTypesSelectFields) Description() LedgerBookTypesField {
	return LedgerBookTypesField("description")
}

func (ss LedgerBookTypesSelectFields) Metadata() LedgerBookTypesField {
	return LedgerBookTypesField("metadata")
}

func (ss LedgerBookTypesSelectFields) MetaCreatedAt() LedgerBookTypesField {
	return LedgerBookTypesField("meta_created_at")
}

func (ss LedgerBookTypesSelectFields) MetaCreatedBy() LedgerBookTypesField {
	return LedgerBookTypesField("meta_created_by")
}

func (ss LedgerBookTypesSelectFields) MetaUpdatedAt() LedgerBookTypesField {
	return LedgerBookTypesField("meta_updated_at")
}

func (ss LedgerBookTypesSelectFields) MetaUpdatedBy() LedgerBookTypesField {
	return LedgerBookTypesField("meta_updated_by")
}

func (ss LedgerBookTypesSelectFields) MetaDeletedAt() LedgerBookTypesField {
	return LedgerBookTypesField("meta_deleted_at")
}

func (ss LedgerBookTypesSelectFields) MetaDeletedBy() LedgerBookTypesField {
	return LedgerBookTypesField("meta_deleted_by")
}

func (ss LedgerBookTypesSelectFields) All() LedgerBookTypesFieldList {
	return []LedgerBookTypesField{
		ss.Code(),
		ss.Description(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerBookTypesSelectFields() LedgerBookTypesSelectFields {
	return LedgerBookTypesSelectFields{}
}

type LedgerBookTypesUpdateFieldOption struct {
	useIncrement bool
}
type LedgerBookTypesUpdateField struct {
	ledgerBookTypesField LedgerBookTypesField
	opt                  LedgerBookTypesUpdateFieldOption
	value                interface{}
}
type LedgerBookTypesUpdateFieldList []LedgerBookTypesUpdateField

func defaultLedgerBookTypesUpdateFieldOption() LedgerBookTypesUpdateFieldOption {
	return LedgerBookTypesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerBookTypesOption(useIncrement bool) func(*LedgerBookTypesUpdateFieldOption) {
	return func(pcufo *LedgerBookTypesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerBookTypesUpdateField(field LedgerBookTypesField, val interface{}, opts ...func(*LedgerBookTypesUpdateFieldOption)) LedgerBookTypesUpdateField {
	defaultOpt := defaultLedgerBookTypesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerBookTypesUpdateField{
		ledgerBookTypesField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultLedgerBookTypesUpdateFields(ledgerBookTypes model.LedgerBookTypes) (ledgerBookTypesUpdateFieldList LedgerBookTypesUpdateFieldList) {
	selectFields := NewLedgerBookTypesSelectFields()
	ledgerBookTypesUpdateFieldList = append(ledgerBookTypesUpdateFieldList,
		NewLedgerBookTypesUpdateField(selectFields.Code(), ledgerBookTypes.Code),
		NewLedgerBookTypesUpdateField(selectFields.Description(), ledgerBookTypes.Description),
		NewLedgerBookTypesUpdateField(selectFields.Metadata(), ledgerBookTypes.Metadata),
		NewLedgerBookTypesUpdateField(selectFields.MetaCreatedAt(), ledgerBookTypes.MetaCreatedAt),
		NewLedgerBookTypesUpdateField(selectFields.MetaCreatedBy(), ledgerBookTypes.MetaCreatedBy),
		NewLedgerBookTypesUpdateField(selectFields.MetaUpdatedAt(), ledgerBookTypes.MetaUpdatedAt),
		NewLedgerBookTypesUpdateField(selectFields.MetaUpdatedBy(), ledgerBookTypes.MetaUpdatedBy),
		NewLedgerBookTypesUpdateField(selectFields.MetaDeletedAt(), ledgerBookTypes.MetaDeletedAt),
		NewLedgerBookTypesUpdateField(selectFields.MetaDeletedBy(), ledgerBookTypes.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerBookTypesCommand(ledgerBookTypesUpdateFieldList LedgerBookTypesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerBookTypesUpdateFieldList {
		field := string(updateField.ledgerBookTypesField)
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

func (repo *RepositoryImpl) BulkCreateLedgerBookTypes(ctx context.Context, ledgerBookTypesList []*model.LedgerBookTypes, fieldsInsert ...LedgerBookTypesField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.LedgerBookTypesPrimaryID
		ledgerBookTypesValueList []model.LedgerBookTypes
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerBookTypesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerBookTypes := range ledgerBookTypesList {

		primaryIds = append(primaryIds, ledgerBookTypes.ToLedgerBookTypesPrimaryID())

		ledgerBookTypesValueList = append(ledgerBookTypesValueList, *ledgerBookTypes)
	}

	_, notFoundIds, err := repo.IsExistLedgerBookTypesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerBookTypes] failed checking ledgerBookTypes whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerBookTypesPrimaryID{}
		mapNotFoundIds := map[model.LedgerBookTypesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerBookTypes", fmt.Sprintf("ledgerBookTypes with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerBookTypes(ledgerBookTypesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerBookTypesQueries.insertLedgerBookTypes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerBookTypes] failed exec create ledgerBookTypes query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerBookTypesByIDs(ctx context.Context, primaryIDs []model.LedgerBookTypesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerBookTypesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerBookTypesByIDs] failed checking ledgerBookTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBookTypes with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_book_types\".\"code\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Code)
	}

	commandQuery := fmt.Sprintf(ledgerBookTypesQueries.deleteLedgerBookTypes + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerBookTypesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerBookTypesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerBookTypesByIDs(ctx context.Context, ids []model.LedgerBookTypesPrimaryID) (exists bool, notFoundIds []model.LedgerBookTypesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_book_types\".\"code\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Code)
	}

	query := fmt.Sprintf(ledgerBookTypesQueries.selectLedgerBookTypes, " \"code\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerBookTypesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerBookTypesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerBookTypesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerBookTypesPrimaryID]bool{}
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

// BulkUpdateLedgerBookTypes is used to bulk update ledgerBookTypes, by default it will update all field
// if want to update specific field, then fill ledgerBookTypessMapUpdateFieldsRequest else please fill ledgerBookTypessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerBookTypes(ctx context.Context, ledgerBookTypessMap map[model.LedgerBookTypesPrimaryID]*model.LedgerBookTypes, ledgerBookTypessMapUpdateFieldsRequest map[model.LedgerBookTypesPrimaryID]LedgerBookTypesUpdateFieldList) (err error) {
	if len(ledgerBookTypessMap) == 0 && len(ledgerBookTypessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerBookTypessMapUpdateField map[model.LedgerBookTypesPrimaryID]LedgerBookTypesUpdateFieldList = map[model.LedgerBookTypesPrimaryID]LedgerBookTypesUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(ledgerBookTypessMap) > 0 {
		for id, ledgerBookTypes := range ledgerBookTypessMap {
			if ledgerBookTypes == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerBookTypes] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerBookTypessMapUpdateField[id] = defaultLedgerBookTypesUpdateFields(*ledgerBookTypes)
		}
	} else {
		ledgerBookTypessMapUpdateField = ledgerBookTypessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerBookTypesQuery(ledgerBookTypessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerBookTypesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerBookTypes] failed checking ledgerBookTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBookTypes with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerBookTypesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_book_types\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerBookTypes] failed exec query")
	}
	return
}

type LedgerBookTypesFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerBookTypesFieldParameter(param string, args ...interface{}) LedgerBookTypesFieldParameter {
	return LedgerBookTypesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerBookTypesQuery(mapLedgerBookTypess map[model.LedgerBookTypesPrimaryID]LedgerBookTypesUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerBookTypesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerBookTypesPrimaryID]map[string]interface{}{}
	ledgerBookTypesSelectFields := NewLedgerBookTypesSelectFields()
	for id, updateFields := range mapLedgerBookTypess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerBookTypesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerBookTypess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerBookTypesFieldType(updateField.ledgerBookTypesField)))
			args = append(args, fields[string(updateField.ledgerBookTypesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerBookTypesField))
		if updateField.ledgerBookTypesField == ledgerBookTypesSelectFields.Code() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerBookTypesField, asTableValues, updateField.ledgerBookTypesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerBookTypesField,
				"\"ledger_book_types\"", updateField.ledgerBookTypesField,
				asTableValues, updateField.ledgerBookTypesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerBookTypesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerBookTypesPrimaryID, asTableValue string) (whereQry string) {
	ledgerBookTypesSelectFields := NewLedgerBookTypesSelectFields()
	var arrWhereQry []string
	code := fmt.Sprintf("\"ledger_book_types\".\"code\" = %s.\"code\"::"+GetLedgerBookTypesFieldType(ledgerBookTypesSelectFields.Code()), asTableValue)
	arrWhereQry = append(arrWhereQry, code)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerBookTypesFieldType(ledgerBookTypesField LedgerBookTypesField) string {
	selectLedgerBookTypesFields := NewLedgerBookTypesSelectFields()
	switch ledgerBookTypesField {

	case selectLedgerBookTypesFields.Code():
		return "text"

	case selectLedgerBookTypesFields.Description():
		return "text"

	case selectLedgerBookTypesFields.Metadata():
		return "jsonb"

	case selectLedgerBookTypesFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerBookTypesFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerBookTypesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerBookTypesFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerBookTypesFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerBookTypesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerBookTypes(ctx context.Context, ledgerBookTypes *model.LedgerBookTypes, fieldsInsert ...LedgerBookTypesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerBookTypesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerBookTypesPrimaryID{
		Code: ledgerBookTypes.Code,
	}
	exists, err := repo.IsExistLedgerBookTypesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerBookTypes] failed checking ledgerBookTypes whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerBookTypes", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerBookTypes([]model.LedgerBookTypes{*ledgerBookTypes}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerBookTypesQueries.insertLedgerBookTypes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerBookTypes] failed exec create ledgerBookTypes query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerBookTypesByID(ctx context.Context, primaryID model.LedgerBookTypesPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerBookTypesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerBookTypesByID] failed checking ledgerBookTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBookTypes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerBookTypesCompositePrimaryKeyWhere([]model.LedgerBookTypesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerBookTypesQueries.deleteLedgerBookTypes + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerBookTypesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerBookTypesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerBookTypesFilterResult, err error) {
	query, args, err := composeLedgerBookTypesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerBookTypesByFilter] failed compose ledgerBookTypes filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerBookTypesByFilter] failed get ledgerBookTypes by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerBookTypesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerBookTypesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerBookTypesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerBookTypesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerBookTypesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 9 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerBookTypesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["code"]; !selected {
			selectColumns = append(selectColumns, "base.\"code\"")
			selectedColumns["code"] = struct{}{}
		}
		if _, selected := selectedColumns["description"]; !selected {
			selectColumns = append(selectColumns, "base.\"description\"")
			selectedColumns["description"] = struct{}{}
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

type ledgerBookTypesFilterPlaceholder struct {
	index int
}

func (p *ledgerBookTypesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerBookTypesFilterPredicate(filterField model.FilterField, placeholders *ledgerBookTypesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerBookTypesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerBookTypesFilterSQLExpr(spec)
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

func composeLedgerBookTypesFilterGroup(group model.FilterGroup, placeholders *ledgerBookTypesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerBookTypesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerBookTypesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerBookTypesFilterWhereQueries(filter model.Filter, placeholders *ledgerBookTypesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerBookTypesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerBookTypesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerBookTypesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerBookTypesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerBookTypesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerBookTypesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerBookTypesFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerBookTypesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerBookTypesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerBookTypesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerBookTypesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_book_types\" base%s", strings.Join(selectColumns, ","), composeLedgerBookTypesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerBookTypesByID(ctx context.Context, primaryID model.LedgerBookTypesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerBookTypesCompositePrimaryKeyWhere([]model.LedgerBookTypesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerBookTypesQueries.selectCountLedgerBookTypes, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerBookTypesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerBookTypes(ctx context.Context, selectFields ...LedgerBookTypesField) (ledgerBookTypesList model.LedgerBookTypesList, err error) {
	var (
		defaultLedgerBookTypesSelectFields = defaultLedgerBookTypesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerBookTypesSelectFields = composeLedgerBookTypesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerBookTypesQueries.selectLedgerBookTypes, defaultLedgerBookTypesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerBookTypesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerBookTypes] failed get ledgerBookTypes list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerBookTypesByID(ctx context.Context, primaryID model.LedgerBookTypesPrimaryID, selectFields ...LedgerBookTypesField) (ledgerBookTypes model.LedgerBookTypes, err error) {
	var (
		defaultLedgerBookTypesSelectFields = defaultLedgerBookTypesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerBookTypesSelectFields = composeLedgerBookTypesSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerBookTypesCompositePrimaryKeyWhere([]model.LedgerBookTypesPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerBookTypesQueries.selectLedgerBookTypes+" WHERE "+whereQry, defaultLedgerBookTypesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerBookTypes, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerBookTypes with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerBookTypesByID] failed get ledgerBookTypes")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerBookTypesByID(ctx context.Context, primaryID model.LedgerBookTypesPrimaryID, ledgerBookTypes *model.LedgerBookTypes, ledgerBookTypesUpdateFields ...LedgerBookTypesUpdateField) (err error) {
	exists, err := repo.IsExistLedgerBookTypesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBookTypes] failed checking ledgerBookTypes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBookTypes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerBookTypes == nil {
		if len(ledgerBookTypesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerBookTypesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerBookTypes = &model.LedgerBookTypes{}
	}
	var (
		defaultLedgerBookTypesUpdateFields = defaultLedgerBookTypesUpdateFields(*ledgerBookTypes)
		tempUpdateField                    LedgerBookTypesUpdateFieldList
		selectFields                       = NewLedgerBookTypesSelectFields()
	)
	if len(ledgerBookTypesUpdateFields) > 0 {
		for _, updateField := range ledgerBookTypesUpdateFields {
			if updateField.ledgerBookTypesField == selectFields.Code() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerBookTypesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerBookTypesCompositePrimaryKeyWhere([]model.LedgerBookTypesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerBookTypesCommand(defaultLedgerBookTypesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerBookTypesQueries.updateLedgerBookTypes+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBookTypes] error when try to update ledgerBookTypes by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerBookTypesByFilter(ctx context.Context, filter model.Filter, ledgerBookTypesUpdateFields ...LedgerBookTypesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerBookTypesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerBookTypesUpdateFieldList
		selectFields = NewLedgerBookTypesSelectFields()
	)
	for _, updateField := range ledgerBookTypesUpdateFields {
		if updateField.ledgerBookTypesField == selectFields.Code() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerBookTypesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerBookTypesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerBookTypesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_book_types\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBookTypesByFilter] error when try to update ledgerBookTypes by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBookTypesByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerBookTypesQueries = struct {
		selectLedgerBookTypes      string
		selectCountLedgerBookTypes string
		deleteLedgerBookTypes      string
		updateLedgerBookTypes      string
		insertLedgerBookTypes      string
	}{
		selectLedgerBookTypes:      "SELECT %s FROM \"ledger_book_types\"",
		selectCountLedgerBookTypes: "SELECT COUNT(\"code\") FROM \"ledger_book_types\"",
		deleteLedgerBookTypes:      "DELETE FROM \"ledger_book_types\"",
		updateLedgerBookTypes:      "UPDATE \"ledger_book_types\" SET %s ",
		insertLedgerBookTypes:      "INSERT INTO \"ledger_book_types\" %s VALUES %s",
	}
)

type LedgerBookTypesRepository interface {
	CreateLedgerBookTypes(ctx context.Context, ledgerBookTypes *model.LedgerBookTypes, fieldsInsert ...LedgerBookTypesField) error
	BulkCreateLedgerBookTypes(ctx context.Context, ledgerBookTypesList []*model.LedgerBookTypes, fieldsInsert ...LedgerBookTypesField) error
	ResolveLedgerBookTypes(ctx context.Context, selectFields ...LedgerBookTypesField) (model.LedgerBookTypesList, error)
	ResolveLedgerBookTypesByID(ctx context.Context, primaryID model.LedgerBookTypesPrimaryID, selectFields ...LedgerBookTypesField) (model.LedgerBookTypes, error)
	UpdateLedgerBookTypesByID(ctx context.Context, id model.LedgerBookTypesPrimaryID, ledgerBookTypes *model.LedgerBookTypes, ledgerBookTypesUpdateFields ...LedgerBookTypesUpdateField) error
	UpdateLedgerBookTypesByFilter(ctx context.Context, filter model.Filter, ledgerBookTypesUpdateFields ...LedgerBookTypesUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerBookTypes(ctx context.Context, ledgerBookTypesListMap map[model.LedgerBookTypesPrimaryID]*model.LedgerBookTypes, LedgerBookTypessMapUpdateFieldsRequest map[model.LedgerBookTypesPrimaryID]LedgerBookTypesUpdateFieldList) (err error)
	DeleteLedgerBookTypesByID(ctx context.Context, id model.LedgerBookTypesPrimaryID) error
	BulkDeleteLedgerBookTypesByIDs(ctx context.Context, ids []model.LedgerBookTypesPrimaryID) error
	ResolveLedgerBookTypesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerBookTypesFilterResult, err error)
	IsExistLedgerBookTypesByIDs(ctx context.Context, ids []model.LedgerBookTypesPrimaryID) (exists bool, notFoundIds []model.LedgerBookTypesPrimaryID, err error)
	IsExistLedgerBookTypesByID(ctx context.Context, id model.LedgerBookTypesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
