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

func composeInsertFieldsAndParamsLedgerBooks(ledgerBooksList []model.LedgerBooks, fieldsInsert ...LedgerBooksField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerBooksSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerBooks := range ledgerBooksList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerBooks.Id)
			case selectField.BookCode():
				args = append(args, ledgerBooks.BookCode)
			case selectField.BookTypeCode():
				args = append(args, ledgerBooks.BookTypeCode)
			case selectField.OwnerPartyId():
				args = append(args, ledgerBooks.OwnerPartyId)
			case selectField.CurrencyCode():
				args = append(args, ledgerBooks.CurrencyCode)
			case selectField.BookStatus():
				args = append(args, ledgerBooks.BookStatus)
			case selectField.CloseCutoffTz():
				args = append(args, ledgerBooks.CloseCutoffTz)
			case selectField.Metadata():
				args = append(args, ledgerBooks.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerBooks.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerBooks.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerBooks.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerBooks.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerBooks.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerBooks.MetaDeletedBy)

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

func composeLedgerBooksCompositePrimaryKeyWhere(primaryIDs []model.LedgerBooksPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_books\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerBooksSelectFields() string {
	fields := NewLedgerBooksSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerBooksSelectFields(selectFields ...LedgerBooksField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerBooksField string
type LedgerBooksFieldList []LedgerBooksField

type LedgerBooksSelectFields struct {
}

func (ss LedgerBooksSelectFields) Id() LedgerBooksField {
	return LedgerBooksField("id")
}

func (ss LedgerBooksSelectFields) BookCode() LedgerBooksField {
	return LedgerBooksField("book_code")
}

func (ss LedgerBooksSelectFields) BookTypeCode() LedgerBooksField {
	return LedgerBooksField("book_type_code")
}

func (ss LedgerBooksSelectFields) OwnerPartyId() LedgerBooksField {
	return LedgerBooksField("owner_party_id")
}

func (ss LedgerBooksSelectFields) CurrencyCode() LedgerBooksField {
	return LedgerBooksField("currency_code")
}

func (ss LedgerBooksSelectFields) BookStatus() LedgerBooksField {
	return LedgerBooksField("book_status")
}

func (ss LedgerBooksSelectFields) CloseCutoffTz() LedgerBooksField {
	return LedgerBooksField("close_cutoff_tz")
}

func (ss LedgerBooksSelectFields) Metadata() LedgerBooksField {
	return LedgerBooksField("metadata")
}

func (ss LedgerBooksSelectFields) MetaCreatedAt() LedgerBooksField {
	return LedgerBooksField("meta_created_at")
}

func (ss LedgerBooksSelectFields) MetaCreatedBy() LedgerBooksField {
	return LedgerBooksField("meta_created_by")
}

func (ss LedgerBooksSelectFields) MetaUpdatedAt() LedgerBooksField {
	return LedgerBooksField("meta_updated_at")
}

func (ss LedgerBooksSelectFields) MetaUpdatedBy() LedgerBooksField {
	return LedgerBooksField("meta_updated_by")
}

func (ss LedgerBooksSelectFields) MetaDeletedAt() LedgerBooksField {
	return LedgerBooksField("meta_deleted_at")
}

func (ss LedgerBooksSelectFields) MetaDeletedBy() LedgerBooksField {
	return LedgerBooksField("meta_deleted_by")
}

func (ss LedgerBooksSelectFields) All() LedgerBooksFieldList {
	return []LedgerBooksField{
		ss.Id(),
		ss.BookCode(),
		ss.BookTypeCode(),
		ss.OwnerPartyId(),
		ss.CurrencyCode(),
		ss.BookStatus(),
		ss.CloseCutoffTz(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerBooksSelectFields() LedgerBooksSelectFields {
	return LedgerBooksSelectFields{}
}

type LedgerBooksUpdateFieldOption struct {
	useIncrement bool
}
type LedgerBooksUpdateField struct {
	ledgerBooksField LedgerBooksField
	opt              LedgerBooksUpdateFieldOption
	value            interface{}
}
type LedgerBooksUpdateFieldList []LedgerBooksUpdateField

func defaultLedgerBooksUpdateFieldOption() LedgerBooksUpdateFieldOption {
	return LedgerBooksUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerBooksOption(useIncrement bool) func(*LedgerBooksUpdateFieldOption) {
	return func(pcufo *LedgerBooksUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerBooksUpdateField(field LedgerBooksField, val interface{}, opts ...func(*LedgerBooksUpdateFieldOption)) LedgerBooksUpdateField {
	defaultOpt := defaultLedgerBooksUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerBooksUpdateField{
		ledgerBooksField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultLedgerBooksUpdateFields(ledgerBooks model.LedgerBooks) (ledgerBooksUpdateFieldList LedgerBooksUpdateFieldList) {
	selectFields := NewLedgerBooksSelectFields()
	ledgerBooksUpdateFieldList = append(ledgerBooksUpdateFieldList,
		NewLedgerBooksUpdateField(selectFields.Id(), ledgerBooks.Id),
		NewLedgerBooksUpdateField(selectFields.BookCode(), ledgerBooks.BookCode),
		NewLedgerBooksUpdateField(selectFields.BookTypeCode(), ledgerBooks.BookTypeCode),
		NewLedgerBooksUpdateField(selectFields.OwnerPartyId(), ledgerBooks.OwnerPartyId),
		NewLedgerBooksUpdateField(selectFields.CurrencyCode(), ledgerBooks.CurrencyCode),
		NewLedgerBooksUpdateField(selectFields.BookStatus(), ledgerBooks.BookStatus),
		NewLedgerBooksUpdateField(selectFields.CloseCutoffTz(), ledgerBooks.CloseCutoffTz),
		NewLedgerBooksUpdateField(selectFields.Metadata(), ledgerBooks.Metadata),
		NewLedgerBooksUpdateField(selectFields.MetaCreatedAt(), ledgerBooks.MetaCreatedAt),
		NewLedgerBooksUpdateField(selectFields.MetaCreatedBy(), ledgerBooks.MetaCreatedBy),
		NewLedgerBooksUpdateField(selectFields.MetaUpdatedAt(), ledgerBooks.MetaUpdatedAt),
		NewLedgerBooksUpdateField(selectFields.MetaUpdatedBy(), ledgerBooks.MetaUpdatedBy),
		NewLedgerBooksUpdateField(selectFields.MetaDeletedAt(), ledgerBooks.MetaDeletedAt),
		NewLedgerBooksUpdateField(selectFields.MetaDeletedBy(), ledgerBooks.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerBooksCommand(ledgerBooksUpdateFieldList LedgerBooksUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerBooksUpdateFieldList {
		field := string(updateField.ledgerBooksField)
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

func (repo *RepositoryImpl) BulkCreateLedgerBooks(ctx context.Context, ledgerBooksList []*model.LedgerBooks, fieldsInsert ...LedgerBooksField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.LedgerBooksPrimaryID
		ledgerBooksValueList []model.LedgerBooks
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerBooksSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerBooks := range ledgerBooksList {

		primaryIds = append(primaryIds, ledgerBooks.ToLedgerBooksPrimaryID())

		ledgerBooksValueList = append(ledgerBooksValueList, *ledgerBooks)
	}

	_, notFoundIds, err := repo.IsExistLedgerBooksByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerBooks] failed checking ledgerBooks whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerBooksPrimaryID{}
		mapNotFoundIds := map[model.LedgerBooksPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerBooks", fmt.Sprintf("ledgerBooks with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerBooks(ledgerBooksValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerBooksQueries.insertLedgerBooks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerBooks] failed exec create ledgerBooks query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerBooksByIDs(ctx context.Context, primaryIDs []model.LedgerBooksPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerBooksByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerBooksByIDs] failed checking ledgerBooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBooks with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_books\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerBooksQueries.deleteLedgerBooks + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerBooksByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerBooksByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerBooksByIDs(ctx context.Context, ids []model.LedgerBooksPrimaryID) (exists bool, notFoundIds []model.LedgerBooksPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_books\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerBooksQueries.selectLedgerBooks, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerBooksByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerBooksPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerBooksByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerBooksPrimaryID]bool{}
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

// BulkUpdateLedgerBooks is used to bulk update ledgerBooks, by default it will update all field
// if want to update specific field, then fill ledgerBookssMapUpdateFieldsRequest else please fill ledgerBookssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerBooks(ctx context.Context, ledgerBookssMap map[model.LedgerBooksPrimaryID]*model.LedgerBooks, ledgerBookssMapUpdateFieldsRequest map[model.LedgerBooksPrimaryID]LedgerBooksUpdateFieldList) (err error) {
	if len(ledgerBookssMap) == 0 && len(ledgerBookssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerBookssMapUpdateField map[model.LedgerBooksPrimaryID]LedgerBooksUpdateFieldList = map[model.LedgerBooksPrimaryID]LedgerBooksUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(ledgerBookssMap) > 0 {
		for id, ledgerBooks := range ledgerBookssMap {
			if ledgerBooks == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerBooks] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerBookssMapUpdateField[id] = defaultLedgerBooksUpdateFields(*ledgerBooks)
		}
	} else {
		ledgerBookssMapUpdateField = ledgerBookssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerBooksQuery(ledgerBookssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerBooksByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerBooks] failed checking ledgerBooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBooks with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerBooksCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_books\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerBooks] failed exec query")
	}
	return
}

type LedgerBooksFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerBooksFieldParameter(param string, args ...interface{}) LedgerBooksFieldParameter {
	return LedgerBooksFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerBooksQuery(mapLedgerBookss map[model.LedgerBooksPrimaryID]LedgerBooksUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerBooksPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerBooksPrimaryID]map[string]interface{}{}
	ledgerBooksSelectFields := NewLedgerBooksSelectFields()
	for id, updateFields := range mapLedgerBookss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerBooksField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerBookss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerBooksFieldType(updateField.ledgerBooksField)))
			args = append(args, fields[string(updateField.ledgerBooksField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerBooksField))
		if updateField.ledgerBooksField == ledgerBooksSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerBooksField, asTableValues, updateField.ledgerBooksField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerBooksField,
				"\"ledger_books\"", updateField.ledgerBooksField,
				asTableValues, updateField.ledgerBooksField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerBooksCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerBooksPrimaryID, asTableValue string) (whereQry string) {
	ledgerBooksSelectFields := NewLedgerBooksSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_books\".\"id\" = %s.\"id\"::"+GetLedgerBooksFieldType(ledgerBooksSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerBooksFieldType(ledgerBooksField LedgerBooksField) string {
	selectLedgerBooksFields := NewLedgerBooksSelectFields()
	switch ledgerBooksField {

	case selectLedgerBooksFields.Id():
		return "uuid"

	case selectLedgerBooksFields.BookCode():
		return "text"

	case selectLedgerBooksFields.BookTypeCode():
		return "text"

	case selectLedgerBooksFields.OwnerPartyId():
		return "uuid"

	case selectLedgerBooksFields.CurrencyCode():
		return "text"

	case selectLedgerBooksFields.BookStatus():
		return "book_status_enum"

	case selectLedgerBooksFields.CloseCutoffTz():
		return "text"

	case selectLedgerBooksFields.Metadata():
		return "jsonb"

	case selectLedgerBooksFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerBooksFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerBooksFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerBooksFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerBooksFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerBooksFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerBooks(ctx context.Context, ledgerBooks *model.LedgerBooks, fieldsInsert ...LedgerBooksField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerBooksSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerBooksPrimaryID{
		Id: ledgerBooks.Id,
	}
	exists, err := repo.IsExistLedgerBooksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerBooks] failed checking ledgerBooks whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerBooks", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerBooks([]model.LedgerBooks{*ledgerBooks}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerBooksQueries.insertLedgerBooks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerBooks] failed exec create ledgerBooks query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerBooksByID(ctx context.Context, primaryID model.LedgerBooksPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerBooksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerBooksByID] failed checking ledgerBooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBooks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerBooksCompositePrimaryKeyWhere([]model.LedgerBooksPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerBooksQueries.deleteLedgerBooks + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerBooksByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerBooksByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerBooksFilterResult, err error) {
	query, args, err := composeLedgerBooksFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerBooksByFilter] failed compose ledgerBooks filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerBooksByFilter] failed get ledgerBooks by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerBooksFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerBooksFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerBooksFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerBooksSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerBooksFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerBooksFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 14+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["book_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_code\"")
			selectedColumns["book_code"] = struct{}{}
		}
		if _, selected := selectedColumns["book_type_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_type_code\"")
			selectedColumns["book_type_code"] = struct{}{}
		}
		if _, selected := selectedColumns["owner_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"owner_party_id\"")
			selectedColumns["owner_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["book_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_status\"")
			selectedColumns["book_status"] = struct{}{}
		}
		if _, selected := selectedColumns["close_cutoff_tz"]; !selected {
			selectColumns = append(selectColumns, "base.\"close_cutoff_tz\"")
			selectedColumns["close_cutoff_tz"] = struct{}{}
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

type ledgerBooksFilterPlaceholder struct {
	index int
}

func (p *ledgerBooksFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerBooksFilterPredicate(filterField model.FilterField, placeholders *ledgerBooksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerBooksFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerBooksFilterSQLExpr(spec)
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

func composeLedgerBooksFilterGroup(group model.FilterGroup, placeholders *ledgerBooksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerBooksFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerBooksFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerBooksFilterWhereQueries(filter model.Filter, placeholders *ledgerBooksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerBooksFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerBooksFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerBooksFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerBooksFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerBooksSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerBooksFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerBooksFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerBooksFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerBooksFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerBooksFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerBooksSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_books\" base%s", strings.Join(selectColumns, ","), composeLedgerBooksFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerBooksByID(ctx context.Context, primaryID model.LedgerBooksPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerBooksCompositePrimaryKeyWhere([]model.LedgerBooksPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerBooksQueries.selectCountLedgerBooks, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerBooksByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerBooks(ctx context.Context, selectFields ...LedgerBooksField) (ledgerBooksList model.LedgerBooksList, err error) {
	var (
		defaultLedgerBooksSelectFields = defaultLedgerBooksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerBooksSelectFields = composeLedgerBooksSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerBooksQueries.selectLedgerBooks, defaultLedgerBooksSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerBooksList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerBooks] failed get ledgerBooks list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerBooksByID(ctx context.Context, primaryID model.LedgerBooksPrimaryID, selectFields ...LedgerBooksField) (ledgerBooks model.LedgerBooks, err error) {
	var (
		defaultLedgerBooksSelectFields = defaultLedgerBooksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerBooksSelectFields = composeLedgerBooksSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerBooksCompositePrimaryKeyWhere([]model.LedgerBooksPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerBooksQueries.selectLedgerBooks+" WHERE "+whereQry, defaultLedgerBooksSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerBooks, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerBooks with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerBooksByID] failed get ledgerBooks")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerBooksByID(ctx context.Context, primaryID model.LedgerBooksPrimaryID, ledgerBooks *model.LedgerBooks, ledgerBooksUpdateFields ...LedgerBooksUpdateField) (err error) {
	exists, err := repo.IsExistLedgerBooksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBooks] failed checking ledgerBooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerBooks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerBooks == nil {
		if len(ledgerBooksUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerBooksByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerBooks = &model.LedgerBooks{}
	}
	var (
		defaultLedgerBooksUpdateFields = defaultLedgerBooksUpdateFields(*ledgerBooks)
		tempUpdateField                LedgerBooksUpdateFieldList
		selectFields                   = NewLedgerBooksSelectFields()
	)
	if len(ledgerBooksUpdateFields) > 0 {
		for _, updateField := range ledgerBooksUpdateFields {
			if updateField.ledgerBooksField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerBooksUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerBooksCompositePrimaryKeyWhere([]model.LedgerBooksPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerBooksCommand(defaultLedgerBooksUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerBooksQueries.updateLedgerBooks+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBooks] error when try to update ledgerBooks by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerBooksByFilter(ctx context.Context, filter model.Filter, ledgerBooksUpdateFields ...LedgerBooksUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerBooksUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerBooksUpdateFieldList
		selectFields = NewLedgerBooksSelectFields()
	)
	for _, updateField := range ledgerBooksUpdateFields {
		if updateField.ledgerBooksField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerBooksCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerBooksFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerBooksFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_books\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBooksByFilter] error when try to update ledgerBooks by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerBooksByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerBooksQueries = struct {
		selectLedgerBooks      string
		selectCountLedgerBooks string
		deleteLedgerBooks      string
		updateLedgerBooks      string
		insertLedgerBooks      string
	}{
		selectLedgerBooks:      "SELECT %s FROM \"ledger_books\"",
		selectCountLedgerBooks: "SELECT COUNT(\"id\") FROM \"ledger_books\"",
		deleteLedgerBooks:      "DELETE FROM \"ledger_books\"",
		updateLedgerBooks:      "UPDATE \"ledger_books\" SET %s ",
		insertLedgerBooks:      "INSERT INTO \"ledger_books\" %s VALUES %s",
	}
)

type LedgerBooksRepository interface {
	CreateLedgerBooks(ctx context.Context, ledgerBooks *model.LedgerBooks, fieldsInsert ...LedgerBooksField) error
	BulkCreateLedgerBooks(ctx context.Context, ledgerBooksList []*model.LedgerBooks, fieldsInsert ...LedgerBooksField) error
	ResolveLedgerBooks(ctx context.Context, selectFields ...LedgerBooksField) (model.LedgerBooksList, error)
	ResolveLedgerBooksByID(ctx context.Context, primaryID model.LedgerBooksPrimaryID, selectFields ...LedgerBooksField) (model.LedgerBooks, error)
	UpdateLedgerBooksByID(ctx context.Context, id model.LedgerBooksPrimaryID, ledgerBooks *model.LedgerBooks, ledgerBooksUpdateFields ...LedgerBooksUpdateField) error
	UpdateLedgerBooksByFilter(ctx context.Context, filter model.Filter, ledgerBooksUpdateFields ...LedgerBooksUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerBooks(ctx context.Context, ledgerBooksListMap map[model.LedgerBooksPrimaryID]*model.LedgerBooks, LedgerBookssMapUpdateFieldsRequest map[model.LedgerBooksPrimaryID]LedgerBooksUpdateFieldList) (err error)
	DeleteLedgerBooksByID(ctx context.Context, id model.LedgerBooksPrimaryID) error
	BulkDeleteLedgerBooksByIDs(ctx context.Context, ids []model.LedgerBooksPrimaryID) error
	ResolveLedgerBooksByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerBooksFilterResult, err error)
	IsExistLedgerBooksByIDs(ctx context.Context, ids []model.LedgerBooksPrimaryID) (exists bool, notFoundIds []model.LedgerBooksPrimaryID, err error)
	IsExistLedgerBooksByID(ctx context.Context, id model.LedgerBooksPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
