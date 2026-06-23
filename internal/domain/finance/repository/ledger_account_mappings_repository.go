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

func composeInsertFieldsAndParamsLedgerAccountMappings(ledgerAccountMappingsList []model.LedgerAccountMappings, fieldsInsert ...LedgerAccountMappingsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerAccountMappingsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerAccountMappings := range ledgerAccountMappingsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerAccountMappings.Id)
			case selectField.MappingCode():
				args = append(args, ledgerAccountMappings.MappingCode)
			case selectField.BookId():
				args = append(args, ledgerAccountMappings.BookId)
			case selectField.SourceType():
				args = append(args, ledgerAccountMappings.SourceType)
			case selectField.SourceSubtype():
				args = append(args, ledgerAccountMappings.SourceSubtype)
			case selectField.AccountId():
				args = append(args, ledgerAccountMappings.AccountId)
			case selectField.Priority():
				args = append(args, ledgerAccountMappings.Priority)
			case selectField.IsActive():
				args = append(args, ledgerAccountMappings.IsActive)
			case selectField.Conditions():
				args = append(args, ledgerAccountMappings.Conditions)
			case selectField.Metadata():
				args = append(args, ledgerAccountMappings.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerAccountMappings.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerAccountMappings.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerAccountMappings.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerAccountMappings.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerAccountMappings.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerAccountMappings.MetaDeletedBy)

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

func composeLedgerAccountMappingsCompositePrimaryKeyWhere(primaryIDs []model.LedgerAccountMappingsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_account_mappings\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerAccountMappingsSelectFields() string {
	fields := NewLedgerAccountMappingsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerAccountMappingsSelectFields(selectFields ...LedgerAccountMappingsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerAccountMappingsField string
type LedgerAccountMappingsFieldList []LedgerAccountMappingsField

type LedgerAccountMappingsSelectFields struct {
}

func (ss LedgerAccountMappingsSelectFields) Id() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("id")
}

func (ss LedgerAccountMappingsSelectFields) MappingCode() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("mapping_code")
}

func (ss LedgerAccountMappingsSelectFields) BookId() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("book_id")
}

func (ss LedgerAccountMappingsSelectFields) SourceType() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("source_type")
}

func (ss LedgerAccountMappingsSelectFields) SourceSubtype() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("source_subtype")
}

func (ss LedgerAccountMappingsSelectFields) AccountId() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("account_id")
}

func (ss LedgerAccountMappingsSelectFields) Priority() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("priority")
}

func (ss LedgerAccountMappingsSelectFields) IsActive() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("is_active")
}

func (ss LedgerAccountMappingsSelectFields) Conditions() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("conditions")
}

func (ss LedgerAccountMappingsSelectFields) Metadata() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("metadata")
}

func (ss LedgerAccountMappingsSelectFields) MetaCreatedAt() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("meta_created_at")
}

func (ss LedgerAccountMappingsSelectFields) MetaCreatedBy() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("meta_created_by")
}

func (ss LedgerAccountMappingsSelectFields) MetaUpdatedAt() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("meta_updated_at")
}

func (ss LedgerAccountMappingsSelectFields) MetaUpdatedBy() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("meta_updated_by")
}

func (ss LedgerAccountMappingsSelectFields) MetaDeletedAt() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("meta_deleted_at")
}

func (ss LedgerAccountMappingsSelectFields) MetaDeletedBy() LedgerAccountMappingsField {
	return LedgerAccountMappingsField("meta_deleted_by")
}

func (ss LedgerAccountMappingsSelectFields) All() LedgerAccountMappingsFieldList {
	return []LedgerAccountMappingsField{
		ss.Id(),
		ss.MappingCode(),
		ss.BookId(),
		ss.SourceType(),
		ss.SourceSubtype(),
		ss.AccountId(),
		ss.Priority(),
		ss.IsActive(),
		ss.Conditions(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerAccountMappingsSelectFields() LedgerAccountMappingsSelectFields {
	return LedgerAccountMappingsSelectFields{}
}

type LedgerAccountMappingsUpdateFieldOption struct {
	useIncrement bool
}
type LedgerAccountMappingsUpdateField struct {
	ledgerAccountMappingsField LedgerAccountMappingsField
	opt                        LedgerAccountMappingsUpdateFieldOption
	value                      interface{}
}
type LedgerAccountMappingsUpdateFieldList []LedgerAccountMappingsUpdateField

func defaultLedgerAccountMappingsUpdateFieldOption() LedgerAccountMappingsUpdateFieldOption {
	return LedgerAccountMappingsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerAccountMappingsOption(useIncrement bool) func(*LedgerAccountMappingsUpdateFieldOption) {
	return func(pcufo *LedgerAccountMappingsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerAccountMappingsUpdateField(field LedgerAccountMappingsField, val interface{}, opts ...func(*LedgerAccountMappingsUpdateFieldOption)) LedgerAccountMappingsUpdateField {
	defaultOpt := defaultLedgerAccountMappingsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerAccountMappingsUpdateField{
		ledgerAccountMappingsField: field,
		value:                      val,
		opt:                        defaultOpt,
	}
}
func defaultLedgerAccountMappingsUpdateFields(ledgerAccountMappings model.LedgerAccountMappings) (ledgerAccountMappingsUpdateFieldList LedgerAccountMappingsUpdateFieldList) {
	selectFields := NewLedgerAccountMappingsSelectFields()
	ledgerAccountMappingsUpdateFieldList = append(ledgerAccountMappingsUpdateFieldList,
		NewLedgerAccountMappingsUpdateField(selectFields.Id(), ledgerAccountMappings.Id),
		NewLedgerAccountMappingsUpdateField(selectFields.MappingCode(), ledgerAccountMappings.MappingCode),
		NewLedgerAccountMappingsUpdateField(selectFields.BookId(), ledgerAccountMappings.BookId),
		NewLedgerAccountMappingsUpdateField(selectFields.SourceType(), ledgerAccountMappings.SourceType),
		NewLedgerAccountMappingsUpdateField(selectFields.SourceSubtype(), ledgerAccountMappings.SourceSubtype),
		NewLedgerAccountMappingsUpdateField(selectFields.AccountId(), ledgerAccountMappings.AccountId),
		NewLedgerAccountMappingsUpdateField(selectFields.Priority(), ledgerAccountMappings.Priority),
		NewLedgerAccountMappingsUpdateField(selectFields.IsActive(), ledgerAccountMappings.IsActive),
		NewLedgerAccountMappingsUpdateField(selectFields.Conditions(), ledgerAccountMappings.Conditions),
		NewLedgerAccountMappingsUpdateField(selectFields.Metadata(), ledgerAccountMappings.Metadata),
		NewLedgerAccountMappingsUpdateField(selectFields.MetaCreatedAt(), ledgerAccountMappings.MetaCreatedAt),
		NewLedgerAccountMappingsUpdateField(selectFields.MetaCreatedBy(), ledgerAccountMappings.MetaCreatedBy),
		NewLedgerAccountMappingsUpdateField(selectFields.MetaUpdatedAt(), ledgerAccountMappings.MetaUpdatedAt),
		NewLedgerAccountMappingsUpdateField(selectFields.MetaUpdatedBy(), ledgerAccountMappings.MetaUpdatedBy),
		NewLedgerAccountMappingsUpdateField(selectFields.MetaDeletedAt(), ledgerAccountMappings.MetaDeletedAt),
		NewLedgerAccountMappingsUpdateField(selectFields.MetaDeletedBy(), ledgerAccountMappings.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerAccountMappingsCommand(ledgerAccountMappingsUpdateFieldList LedgerAccountMappingsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerAccountMappingsUpdateFieldList {
		field := string(updateField.ledgerAccountMappingsField)
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

func (repo *RepositoryImpl) BulkCreateLedgerAccountMappings(ctx context.Context, ledgerAccountMappingsList []*model.LedgerAccountMappings, fieldsInsert ...LedgerAccountMappingsField) (err error) {
	var (
		fieldsStr                      string
		valueListStr                   []string
		argsList                       []interface{}
		primaryIds                     []model.LedgerAccountMappingsPrimaryID
		ledgerAccountMappingsValueList []model.LedgerAccountMappings
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountMappingsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerAccountMappings := range ledgerAccountMappingsList {

		primaryIds = append(primaryIds, ledgerAccountMappings.ToLedgerAccountMappingsPrimaryID())

		ledgerAccountMappingsValueList = append(ledgerAccountMappingsValueList, *ledgerAccountMappings)
	}

	_, notFoundIds, err := repo.IsExistLedgerAccountMappingsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccountMappings] failed checking ledgerAccountMappings whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerAccountMappingsPrimaryID{}
		mapNotFoundIds := map[model.LedgerAccountMappingsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerAccountMappings", fmt.Sprintf("ledgerAccountMappings with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerAccountMappings(ledgerAccountMappingsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerAccountMappingsQueries.insertLedgerAccountMappings, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerAccountMappings] failed exec create ledgerAccountMappings query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerAccountMappingsByIDs(ctx context.Context, primaryIDs []model.LedgerAccountMappingsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerAccountMappingsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountMappingsByIDs] failed checking ledgerAccountMappings whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountMappings with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_account_mappings\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerAccountMappingsQueries.deleteLedgerAccountMappings + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountMappingsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerAccountMappingsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerAccountMappingsByIDs(ctx context.Context, ids []model.LedgerAccountMappingsPrimaryID) (exists bool, notFoundIds []model.LedgerAccountMappingsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_account_mappings\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerAccountMappingsQueries.selectLedgerAccountMappings, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountMappingsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerAccountMappingsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountMappingsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerAccountMappingsPrimaryID]bool{}
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

// BulkUpdateLedgerAccountMappings is used to bulk update ledgerAccountMappings, by default it will update all field
// if want to update specific field, then fill ledgerAccountMappingssMapUpdateFieldsRequest else please fill ledgerAccountMappingssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerAccountMappings(ctx context.Context, ledgerAccountMappingssMap map[model.LedgerAccountMappingsPrimaryID]*model.LedgerAccountMappings, ledgerAccountMappingssMapUpdateFieldsRequest map[model.LedgerAccountMappingsPrimaryID]LedgerAccountMappingsUpdateFieldList) (err error) {
	if len(ledgerAccountMappingssMap) == 0 && len(ledgerAccountMappingssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerAccountMappingssMapUpdateField map[model.LedgerAccountMappingsPrimaryID]LedgerAccountMappingsUpdateFieldList = map[model.LedgerAccountMappingsPrimaryID]LedgerAccountMappingsUpdateFieldList{}
		asTableValues                        string                                                                        = "myvalues"
	)

	if len(ledgerAccountMappingssMap) > 0 {
		for id, ledgerAccountMappings := range ledgerAccountMappingssMap {
			if ledgerAccountMappings == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerAccountMappings] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerAccountMappingssMapUpdateField[id] = defaultLedgerAccountMappingsUpdateFields(*ledgerAccountMappings)
		}
	} else {
		ledgerAccountMappingssMapUpdateField = ledgerAccountMappingssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerAccountMappingsQuery(ledgerAccountMappingssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerAccountMappingsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccountMappings] failed checking ledgerAccountMappings whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountMappings with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerAccountMappingsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_account_mappings\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerAccountMappings] failed exec query")
	}
	return
}

type LedgerAccountMappingsFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerAccountMappingsFieldParameter(param string, args ...interface{}) LedgerAccountMappingsFieldParameter {
	return LedgerAccountMappingsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerAccountMappingsQuery(mapLedgerAccountMappingss map[model.LedgerAccountMappingsPrimaryID]LedgerAccountMappingsUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerAccountMappingsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerAccountMappingsPrimaryID]map[string]interface{}{}
	ledgerAccountMappingsSelectFields := NewLedgerAccountMappingsSelectFields()
	for id, updateFields := range mapLedgerAccountMappingss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerAccountMappingsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerAccountMappingss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerAccountMappingsFieldType(updateField.ledgerAccountMappingsField)))
			args = append(args, fields[string(updateField.ledgerAccountMappingsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerAccountMappingsField))
		if updateField.ledgerAccountMappingsField == ledgerAccountMappingsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerAccountMappingsField, asTableValues, updateField.ledgerAccountMappingsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerAccountMappingsField,
				"\"ledger_account_mappings\"", updateField.ledgerAccountMappingsField,
				asTableValues, updateField.ledgerAccountMappingsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerAccountMappingsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerAccountMappingsPrimaryID, asTableValue string) (whereQry string) {
	ledgerAccountMappingsSelectFields := NewLedgerAccountMappingsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_account_mappings\".\"id\" = %s.\"id\"::"+GetLedgerAccountMappingsFieldType(ledgerAccountMappingsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerAccountMappingsFieldType(ledgerAccountMappingsField LedgerAccountMappingsField) string {
	selectLedgerAccountMappingsFields := NewLedgerAccountMappingsSelectFields()
	switch ledgerAccountMappingsField {

	case selectLedgerAccountMappingsFields.Id():
		return "uuid"

	case selectLedgerAccountMappingsFields.MappingCode():
		return "text"

	case selectLedgerAccountMappingsFields.BookId():
		return "uuid"

	case selectLedgerAccountMappingsFields.SourceType():
		return "text"

	case selectLedgerAccountMappingsFields.SourceSubtype():
		return "text"

	case selectLedgerAccountMappingsFields.AccountId():
		return "uuid"

	case selectLedgerAccountMappingsFields.Priority():
		return "int4"

	case selectLedgerAccountMappingsFields.IsActive():
		return "bool"

	case selectLedgerAccountMappingsFields.Conditions():
		return "jsonb"

	case selectLedgerAccountMappingsFields.Metadata():
		return "jsonb"

	case selectLedgerAccountMappingsFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerAccountMappingsFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerAccountMappingsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerAccountMappingsFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerAccountMappingsFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerAccountMappingsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerAccountMappings(ctx context.Context, ledgerAccountMappings *model.LedgerAccountMappings, fieldsInsert ...LedgerAccountMappingsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerAccountMappingsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerAccountMappingsPrimaryID{
		Id: ledgerAccountMappings.Id,
	}
	exists, err := repo.IsExistLedgerAccountMappingsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccountMappings] failed checking ledgerAccountMappings whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerAccountMappings", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerAccountMappings([]model.LedgerAccountMappings{*ledgerAccountMappings}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerAccountMappingsQueries.insertLedgerAccountMappings, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerAccountMappings] failed exec create ledgerAccountMappings query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerAccountMappingsByID(ctx context.Context, primaryID model.LedgerAccountMappingsPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerAccountMappingsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountMappingsByID] failed checking ledgerAccountMappings whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountMappings with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerAccountMappingsCompositePrimaryKeyWhere([]model.LedgerAccountMappingsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerAccountMappingsQueries.deleteLedgerAccountMappings + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerAccountMappingsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountMappingsByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountMappingsFilterResult, err error) {
	query, args, err := composeLedgerAccountMappingsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountMappingsByFilter] failed compose ledgerAccountMappings filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountMappingsByFilter] failed get ledgerAccountMappings by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerAccountMappingsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerAccountMappingsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerAccountMappingsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerAccountMappingsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerAccountMappingsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 16 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerAccountMappingsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["mapping_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"mapping_code\"")
			selectedColumns["mapping_code"] = struct{}{}
		}
		if _, selected := selectedColumns["book_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_id\"")
			selectedColumns["book_id"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_subtype"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_subtype\"")
			selectedColumns["source_subtype"] = struct{}{}
		}
		if _, selected := selectedColumns["account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_id\"")
			selectedColumns["account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["priority"]; !selected {
			selectColumns = append(selectColumns, "base.\"priority\"")
			selectedColumns["priority"] = struct{}{}
		}
		if _, selected := selectedColumns["is_active"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_active\"")
			selectedColumns["is_active"] = struct{}{}
		}
		if _, selected := selectedColumns["conditions"]; !selected {
			selectColumns = append(selectColumns, "base.\"conditions\"")
			selectedColumns["conditions"] = struct{}{}
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

type ledgerAccountMappingsFilterPlaceholder struct {
	index int
}

func (p *ledgerAccountMappingsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerAccountMappingsFilterPredicate(filterField model.FilterField, placeholders *ledgerAccountMappingsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerAccountMappingsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerAccountMappingsFilterSQLExpr(spec)
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

func composeLedgerAccountMappingsFilterGroup(group model.FilterGroup, placeholders *ledgerAccountMappingsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerAccountMappingsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerAccountMappingsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerAccountMappingsFilterWhereQueries(filter model.Filter, placeholders *ledgerAccountMappingsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerAccountMappingsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerAccountMappingsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerAccountMappingsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerAccountMappingsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerAccountMappingsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerAccountMappingsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerAccountMappingsFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerAccountMappingsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerAccountMappingsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerAccountMappingsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerAccountMappingsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_account_mappings\" base%s", strings.Join(selectColumns, ","), composeLedgerAccountMappingsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerAccountMappingsByID(ctx context.Context, primaryID model.LedgerAccountMappingsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerAccountMappingsCompositePrimaryKeyWhere([]model.LedgerAccountMappingsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerAccountMappingsQueries.selectCountLedgerAccountMappings, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerAccountMappingsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountMappings(ctx context.Context, selectFields ...LedgerAccountMappingsField) (ledgerAccountMappingsList model.LedgerAccountMappingsList, err error) {
	var (
		defaultLedgerAccountMappingsSelectFields = defaultLedgerAccountMappingsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountMappingsSelectFields = composeLedgerAccountMappingsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerAccountMappingsQueries.selectLedgerAccountMappings, defaultLedgerAccountMappingsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerAccountMappingsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerAccountMappings] failed get ledgerAccountMappings list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerAccountMappingsByID(ctx context.Context, primaryID model.LedgerAccountMappingsPrimaryID, selectFields ...LedgerAccountMappingsField) (ledgerAccountMappings model.LedgerAccountMappings, err error) {
	var (
		defaultLedgerAccountMappingsSelectFields = defaultLedgerAccountMappingsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerAccountMappingsSelectFields = composeLedgerAccountMappingsSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerAccountMappingsCompositePrimaryKeyWhere([]model.LedgerAccountMappingsPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerAccountMappingsQueries.selectLedgerAccountMappings+" WHERE "+whereQry, defaultLedgerAccountMappingsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerAccountMappings, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerAccountMappings with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerAccountMappingsByID] failed get ledgerAccountMappings")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerAccountMappingsByID(ctx context.Context, primaryID model.LedgerAccountMappingsPrimaryID, ledgerAccountMappings *model.LedgerAccountMappings, ledgerAccountMappingsUpdateFields ...LedgerAccountMappingsUpdateField) (err error) {
	exists, err := repo.IsExistLedgerAccountMappingsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountMappings] failed checking ledgerAccountMappings whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerAccountMappings with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerAccountMappings == nil {
		if len(ledgerAccountMappingsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerAccountMappingsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerAccountMappings = &model.LedgerAccountMappings{}
	}
	var (
		defaultLedgerAccountMappingsUpdateFields = defaultLedgerAccountMappingsUpdateFields(*ledgerAccountMappings)
		tempUpdateField                          LedgerAccountMappingsUpdateFieldList
		selectFields                             = NewLedgerAccountMappingsSelectFields()
	)
	if len(ledgerAccountMappingsUpdateFields) > 0 {
		for _, updateField := range ledgerAccountMappingsUpdateFields {
			if updateField.ledgerAccountMappingsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerAccountMappingsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerAccountMappingsCompositePrimaryKeyWhere([]model.LedgerAccountMappingsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerAccountMappingsCommand(defaultLedgerAccountMappingsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerAccountMappingsQueries.updateLedgerAccountMappings+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountMappings] error when try to update ledgerAccountMappings by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerAccountMappingsByFilter(ctx context.Context, filter model.Filter, ledgerAccountMappingsUpdateFields ...LedgerAccountMappingsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerAccountMappingsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerAccountMappingsUpdateFieldList
		selectFields = NewLedgerAccountMappingsSelectFields()
	)
	for _, updateField := range ledgerAccountMappingsUpdateFields {
		if updateField.ledgerAccountMappingsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerAccountMappingsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerAccountMappingsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerAccountMappingsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_account_mappings\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountMappingsByFilter] error when try to update ledgerAccountMappings by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerAccountMappingsByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerAccountMappingsQueries = struct {
		selectLedgerAccountMappings      string
		selectCountLedgerAccountMappings string
		deleteLedgerAccountMappings      string
		updateLedgerAccountMappings      string
		insertLedgerAccountMappings      string
	}{
		selectLedgerAccountMappings:      "SELECT %s FROM \"ledger_account_mappings\"",
		selectCountLedgerAccountMappings: "SELECT COUNT(\"id\") FROM \"ledger_account_mappings\"",
		deleteLedgerAccountMappings:      "DELETE FROM \"ledger_account_mappings\"",
		updateLedgerAccountMappings:      "UPDATE \"ledger_account_mappings\" SET %s ",
		insertLedgerAccountMappings:      "INSERT INTO \"ledger_account_mappings\" %s VALUES %s",
	}
)

type LedgerAccountMappingsRepository interface {
	CreateLedgerAccountMappings(ctx context.Context, ledgerAccountMappings *model.LedgerAccountMappings, fieldsInsert ...LedgerAccountMappingsField) error
	BulkCreateLedgerAccountMappings(ctx context.Context, ledgerAccountMappingsList []*model.LedgerAccountMappings, fieldsInsert ...LedgerAccountMappingsField) error
	ResolveLedgerAccountMappings(ctx context.Context, selectFields ...LedgerAccountMappingsField) (model.LedgerAccountMappingsList, error)
	ResolveLedgerAccountMappingsByID(ctx context.Context, primaryID model.LedgerAccountMappingsPrimaryID, selectFields ...LedgerAccountMappingsField) (model.LedgerAccountMappings, error)
	UpdateLedgerAccountMappingsByID(ctx context.Context, id model.LedgerAccountMappingsPrimaryID, ledgerAccountMappings *model.LedgerAccountMappings, ledgerAccountMappingsUpdateFields ...LedgerAccountMappingsUpdateField) error
	UpdateLedgerAccountMappingsByFilter(ctx context.Context, filter model.Filter, ledgerAccountMappingsUpdateFields ...LedgerAccountMappingsUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerAccountMappings(ctx context.Context, ledgerAccountMappingsListMap map[model.LedgerAccountMappingsPrimaryID]*model.LedgerAccountMappings, LedgerAccountMappingssMapUpdateFieldsRequest map[model.LedgerAccountMappingsPrimaryID]LedgerAccountMappingsUpdateFieldList) (err error)
	DeleteLedgerAccountMappingsByID(ctx context.Context, id model.LedgerAccountMappingsPrimaryID) error
	BulkDeleteLedgerAccountMappingsByIDs(ctx context.Context, ids []model.LedgerAccountMappingsPrimaryID) error
	ResolveLedgerAccountMappingsByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerAccountMappingsFilterResult, err error)
	IsExistLedgerAccountMappingsByIDs(ctx context.Context, ids []model.LedgerAccountMappingsPrimaryID) (exists bool, notFoundIds []model.LedgerAccountMappingsPrimaryID, err error)
	IsExistLedgerAccountMappingsByID(ctx context.Context, id model.LedgerAccountMappingsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
