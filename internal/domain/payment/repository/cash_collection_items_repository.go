package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsCashCollectionItems(cashCollectionItemsList []model.CashCollectionItems, fieldsInsert ...CashCollectionItemsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewCashCollectionItemsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, cashCollectionItems := range cashCollectionItemsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, cashCollectionItems.Id)
			case selectField.CashCollectionSessionId():
				args = append(args, cashCollectionItems.CashCollectionSessionId)
			case selectField.PaymentIntentId():
				args = append(args, cashCollectionItems.PaymentIntentId)
			case selectField.PaymentAttemptId():
				args = append(args, cashCollectionItems.PaymentAttemptId)
			case selectField.CollectionType():
				args = append(args, cashCollectionItems.CollectionType)
			case selectField.Amount():
				args = append(args, cashCollectionItems.Amount)
			case selectField.Currency():
				args = append(args, cashCollectionItems.Currency)
			case selectField.Status():
				args = append(args, cashCollectionItems.Status)
			case selectField.CollectedAt():
				args = append(args, cashCollectionItems.CollectedAt)
			case selectField.VoidedAt():
				args = append(args, cashCollectionItems.VoidedAt)
			case selectField.VoidReason():
				args = append(args, cashCollectionItems.VoidReason)
			case selectField.Notes():
				args = append(args, cashCollectionItems.Notes)
			case selectField.Metadata():
				args = append(args, cashCollectionItems.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, cashCollectionItems.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, cashCollectionItems.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, cashCollectionItems.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, cashCollectionItems.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, cashCollectionItems.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, cashCollectionItems.MetaDeletedBy)

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

func composeCashCollectionItemsCompositePrimaryKeyWhere(primaryIDs []model.CashCollectionItemsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"cash_collection_items\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultCashCollectionItemsSelectFields() string {
	fields := NewCashCollectionItemsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeCashCollectionItemsSelectFields(selectFields ...CashCollectionItemsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type CashCollectionItemsField string
type CashCollectionItemsFieldList []CashCollectionItemsField

type CashCollectionItemsSelectFields struct {
}

func (ss CashCollectionItemsSelectFields) Id() CashCollectionItemsField {
	return CashCollectionItemsField("id")
}

func (ss CashCollectionItemsSelectFields) CashCollectionSessionId() CashCollectionItemsField {
	return CashCollectionItemsField("cash_collection_session_id")
}

func (ss CashCollectionItemsSelectFields) PaymentIntentId() CashCollectionItemsField {
	return CashCollectionItemsField("payment_intent_id")
}

func (ss CashCollectionItemsSelectFields) PaymentAttemptId() CashCollectionItemsField {
	return CashCollectionItemsField("payment_attempt_id")
}

func (ss CashCollectionItemsSelectFields) CollectionType() CashCollectionItemsField {
	return CashCollectionItemsField("collection_type")
}

func (ss CashCollectionItemsSelectFields) Amount() CashCollectionItemsField {
	return CashCollectionItemsField("amount")
}

func (ss CashCollectionItemsSelectFields) Currency() CashCollectionItemsField {
	return CashCollectionItemsField("currency")
}

func (ss CashCollectionItemsSelectFields) Status() CashCollectionItemsField {
	return CashCollectionItemsField("status")
}

func (ss CashCollectionItemsSelectFields) CollectedAt() CashCollectionItemsField {
	return CashCollectionItemsField("collected_at")
}

func (ss CashCollectionItemsSelectFields) VoidedAt() CashCollectionItemsField {
	return CashCollectionItemsField("voided_at")
}

func (ss CashCollectionItemsSelectFields) VoidReason() CashCollectionItemsField {
	return CashCollectionItemsField("void_reason")
}

func (ss CashCollectionItemsSelectFields) Notes() CashCollectionItemsField {
	return CashCollectionItemsField("notes")
}

func (ss CashCollectionItemsSelectFields) Metadata() CashCollectionItemsField {
	return CashCollectionItemsField("metadata")
}

func (ss CashCollectionItemsSelectFields) MetaCreatedAt() CashCollectionItemsField {
	return CashCollectionItemsField("meta_created_at")
}

func (ss CashCollectionItemsSelectFields) MetaCreatedBy() CashCollectionItemsField {
	return CashCollectionItemsField("meta_created_by")
}

func (ss CashCollectionItemsSelectFields) MetaUpdatedAt() CashCollectionItemsField {
	return CashCollectionItemsField("meta_updated_at")
}

func (ss CashCollectionItemsSelectFields) MetaUpdatedBy() CashCollectionItemsField {
	return CashCollectionItemsField("meta_updated_by")
}

func (ss CashCollectionItemsSelectFields) MetaDeletedAt() CashCollectionItemsField {
	return CashCollectionItemsField("meta_deleted_at")
}

func (ss CashCollectionItemsSelectFields) MetaDeletedBy() CashCollectionItemsField {
	return CashCollectionItemsField("meta_deleted_by")
}

func (ss CashCollectionItemsSelectFields) All() CashCollectionItemsFieldList {
	return []CashCollectionItemsField{
		ss.Id(),
		ss.CashCollectionSessionId(),
		ss.PaymentIntentId(),
		ss.PaymentAttemptId(),
		ss.CollectionType(),
		ss.Amount(),
		ss.Currency(),
		ss.Status(),
		ss.CollectedAt(),
		ss.VoidedAt(),
		ss.VoidReason(),
		ss.Notes(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewCashCollectionItemsSelectFields() CashCollectionItemsSelectFields {
	return CashCollectionItemsSelectFields{}
}

type CashCollectionItemsUpdateFieldOption struct {
	useIncrement bool
}
type CashCollectionItemsUpdateField struct {
	cashCollectionItemsField CashCollectionItemsField
	opt                      CashCollectionItemsUpdateFieldOption
	value                    interface{}
}
type CashCollectionItemsUpdateFieldList []CashCollectionItemsUpdateField

func defaultCashCollectionItemsUpdateFieldOption() CashCollectionItemsUpdateFieldOption {
	return CashCollectionItemsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementCashCollectionItemsOption(useIncrement bool) func(*CashCollectionItemsUpdateFieldOption) {
	return func(pcufo *CashCollectionItemsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewCashCollectionItemsUpdateField(field CashCollectionItemsField, val interface{}, opts ...func(*CashCollectionItemsUpdateFieldOption)) CashCollectionItemsUpdateField {
	defaultOpt := defaultCashCollectionItemsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return CashCollectionItemsUpdateField{
		cashCollectionItemsField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultCashCollectionItemsUpdateFields(cashCollectionItems model.CashCollectionItems) (cashCollectionItemsUpdateFieldList CashCollectionItemsUpdateFieldList) {
	selectFields := NewCashCollectionItemsSelectFields()
	cashCollectionItemsUpdateFieldList = append(cashCollectionItemsUpdateFieldList,
		NewCashCollectionItemsUpdateField(selectFields.Id(), cashCollectionItems.Id),
		NewCashCollectionItemsUpdateField(selectFields.CashCollectionSessionId(), cashCollectionItems.CashCollectionSessionId),
		NewCashCollectionItemsUpdateField(selectFields.PaymentIntentId(), cashCollectionItems.PaymentIntentId),
		NewCashCollectionItemsUpdateField(selectFields.PaymentAttemptId(), cashCollectionItems.PaymentAttemptId),
		NewCashCollectionItemsUpdateField(selectFields.CollectionType(), cashCollectionItems.CollectionType),
		NewCashCollectionItemsUpdateField(selectFields.Amount(), cashCollectionItems.Amount),
		NewCashCollectionItemsUpdateField(selectFields.Currency(), cashCollectionItems.Currency),
		NewCashCollectionItemsUpdateField(selectFields.Status(), cashCollectionItems.Status),
		NewCashCollectionItemsUpdateField(selectFields.CollectedAt(), cashCollectionItems.CollectedAt),
		NewCashCollectionItemsUpdateField(selectFields.VoidedAt(), cashCollectionItems.VoidedAt),
		NewCashCollectionItemsUpdateField(selectFields.VoidReason(), cashCollectionItems.VoidReason),
		NewCashCollectionItemsUpdateField(selectFields.Notes(), cashCollectionItems.Notes),
		NewCashCollectionItemsUpdateField(selectFields.Metadata(), cashCollectionItems.Metadata),
		NewCashCollectionItemsUpdateField(selectFields.MetaCreatedAt(), cashCollectionItems.MetaCreatedAt),
		NewCashCollectionItemsUpdateField(selectFields.MetaCreatedBy(), cashCollectionItems.MetaCreatedBy),
		NewCashCollectionItemsUpdateField(selectFields.MetaUpdatedAt(), cashCollectionItems.MetaUpdatedAt),
		NewCashCollectionItemsUpdateField(selectFields.MetaUpdatedBy(), cashCollectionItems.MetaUpdatedBy),
		NewCashCollectionItemsUpdateField(selectFields.MetaDeletedAt(), cashCollectionItems.MetaDeletedAt),
		NewCashCollectionItemsUpdateField(selectFields.MetaDeletedBy(), cashCollectionItems.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsCashCollectionItemsCommand(cashCollectionItemsUpdateFieldList CashCollectionItemsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range cashCollectionItemsUpdateFieldList {
		field := string(updateField.cashCollectionItemsField)
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

func (repo *RepositoryImpl) BulkCreateCashCollectionItems(ctx context.Context, cashCollectionItemsList []*model.CashCollectionItems, fieldsInsert ...CashCollectionItemsField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.CashCollectionItemsPrimaryID
		cashCollectionItemsValueList []model.CashCollectionItems
	)

	if len(fieldsInsert) == 0 {
		selectField := NewCashCollectionItemsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, cashCollectionItems := range cashCollectionItemsList {

		primaryIds = append(primaryIds, cashCollectionItems.ToCashCollectionItemsPrimaryID())

		cashCollectionItemsValueList = append(cashCollectionItemsValueList, *cashCollectionItems)
	}

	_, notFoundIds, err := repo.IsExistCashCollectionItemsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateCashCollectionItems] failed checking cashCollectionItems whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.CashCollectionItemsPrimaryID{}
		mapNotFoundIds := map[model.CashCollectionItemsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "cashCollectionItems", fmt.Sprintf("cashCollectionItems with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsCashCollectionItems(cashCollectionItemsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(cashCollectionItemsQueries.insertCashCollectionItems, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateCashCollectionItems] failed exec create cashCollectionItems query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteCashCollectionItemsByIDs(ctx context.Context, primaryIDs []model.CashCollectionItemsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistCashCollectionItemsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCashCollectionItemsByIDs] failed checking cashCollectionItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionItems with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"cash_collection_items\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(cashCollectionItemsQueries.deleteCashCollectionItems + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCashCollectionItemsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCashCollectionItemsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistCashCollectionItemsByIDs(ctx context.Context, ids []model.CashCollectionItemsPrimaryID) (exists bool, notFoundIds []model.CashCollectionItemsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"cash_collection_items\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(cashCollectionItemsQueries.selectCashCollectionItems, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCashCollectionItemsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.CashCollectionItemsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCashCollectionItemsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.CashCollectionItemsPrimaryID]bool{}
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

// BulkUpdateCashCollectionItems is used to bulk update cashCollectionItems, by default it will update all field
// if want to update specific field, then fill cashCollectionItemssMapUpdateFieldsRequest else please fill cashCollectionItemssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateCashCollectionItems(ctx context.Context, cashCollectionItemssMap map[model.CashCollectionItemsPrimaryID]*model.CashCollectionItems, cashCollectionItemssMapUpdateFieldsRequest map[model.CashCollectionItemsPrimaryID]CashCollectionItemsUpdateFieldList) (err error) {
	if len(cashCollectionItemssMap) == 0 && len(cashCollectionItemssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		cashCollectionItemssMapUpdateField map[model.CashCollectionItemsPrimaryID]CashCollectionItemsUpdateFieldList = map[model.CashCollectionItemsPrimaryID]CashCollectionItemsUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(cashCollectionItemssMap) > 0 {
		for id, cashCollectionItems := range cashCollectionItemssMap {
			if cashCollectionItems == nil {
				log.Error().Err(err).Msg("[BulkUpdateCashCollectionItems] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			cashCollectionItemssMapUpdateField[id] = defaultCashCollectionItemsUpdateFields(*cashCollectionItems)
		}
	} else {
		cashCollectionItemssMapUpdateField = cashCollectionItemssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateCashCollectionItemsQuery(cashCollectionItemssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistCashCollectionItemsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateCashCollectionItems] failed checking cashCollectionItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionItems with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeCashCollectionItemsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"cash_collection_items\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateCashCollectionItems] failed exec query")
	}
	return
}

type CashCollectionItemsFieldParameter struct {
	param string
	args  []interface{}
}

func NewCashCollectionItemsFieldParameter(param string, args ...interface{}) CashCollectionItemsFieldParameter {
	return CashCollectionItemsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateCashCollectionItemsQuery(mapCashCollectionItemss map[model.CashCollectionItemsPrimaryID]CashCollectionItemsUpdateFieldList, asTableValues string) (primaryIDs []model.CashCollectionItemsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.CashCollectionItemsPrimaryID]map[string]interface{}{}
	cashCollectionItemsSelectFields := NewCashCollectionItemsSelectFields()
	for id, updateFields := range mapCashCollectionItemss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.cashCollectionItemsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapCashCollectionItemss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetCashCollectionItemsFieldType(updateField.cashCollectionItemsField)))
			args = append(args, fields[string(updateField.cashCollectionItemsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.cashCollectionItemsField))
		if updateField.cashCollectionItemsField == cashCollectionItemsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.cashCollectionItemsField, asTableValues, updateField.cashCollectionItemsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.cashCollectionItemsField,
				"\"cash_collection_items\"", updateField.cashCollectionItemsField,
				asTableValues, updateField.cashCollectionItemsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeCashCollectionItemsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.CashCollectionItemsPrimaryID, asTableValue string) (whereQry string) {
	cashCollectionItemsSelectFields := NewCashCollectionItemsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"cash_collection_items\".\"id\" = %s.\"id\"::"+GetCashCollectionItemsFieldType(cashCollectionItemsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetCashCollectionItemsFieldType(cashCollectionItemsField CashCollectionItemsField) string {
	selectCashCollectionItemsFields := NewCashCollectionItemsSelectFields()
	switch cashCollectionItemsField {

	case selectCashCollectionItemsFields.Id():
		return "uuid"

	case selectCashCollectionItemsFields.CashCollectionSessionId():
		return "uuid"

	case selectCashCollectionItemsFields.PaymentIntentId():
		return "uuid"

	case selectCashCollectionItemsFields.PaymentAttemptId():
		return "uuid"

	case selectCashCollectionItemsFields.CollectionType():
		return "text"

	case selectCashCollectionItemsFields.Amount():
		return "numeric"

	case selectCashCollectionItemsFields.Currency():
		return "text"

	case selectCashCollectionItemsFields.Status():
		return "cash_item_status_enum"

	case selectCashCollectionItemsFields.CollectedAt():
		return "timestamptz"

	case selectCashCollectionItemsFields.VoidedAt():
		return "timestamptz"

	case selectCashCollectionItemsFields.VoidReason():
		return "text"

	case selectCashCollectionItemsFields.Notes():
		return "text"

	case selectCashCollectionItemsFields.Metadata():
		return "jsonb"

	case selectCashCollectionItemsFields.MetaCreatedAt():
		return "timestamptz"

	case selectCashCollectionItemsFields.MetaCreatedBy():
		return "uuid"

	case selectCashCollectionItemsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectCashCollectionItemsFields.MetaUpdatedBy():
		return "uuid"

	case selectCashCollectionItemsFields.MetaDeletedAt():
		return "timestamptz"

	case selectCashCollectionItemsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateCashCollectionItems(ctx context.Context, cashCollectionItems *model.CashCollectionItems, fieldsInsert ...CashCollectionItemsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewCashCollectionItemsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.CashCollectionItemsPrimaryID{
		Id: cashCollectionItems.Id,
	}
	exists, err := repo.IsExistCashCollectionItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateCashCollectionItems] failed checking cashCollectionItems whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "cashCollectionItems", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsCashCollectionItems([]model.CashCollectionItems{*cashCollectionItems}, fieldsInsert...)
	commandQuery := fmt.Sprintf(cashCollectionItemsQueries.insertCashCollectionItems, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateCashCollectionItems] failed exec create cashCollectionItems query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteCashCollectionItemsByID(ctx context.Context, primaryID model.CashCollectionItemsPrimaryID) (err error) {
	exists, err := repo.IsExistCashCollectionItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteCashCollectionItemsByID] failed checking cashCollectionItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionItems with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeCashCollectionItemsCompositePrimaryKeyWhere([]model.CashCollectionItemsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(cashCollectionItemsQueries.deleteCashCollectionItems + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteCashCollectionItemsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveCashCollectionItemsByFilter(ctx context.Context, filter model.Filter) (result []model.CashCollectionItemsFilterResult, err error) {
	query, args, err := composeCashCollectionItemsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCashCollectionItemsByFilter] failed compose cashCollectionItems filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCashCollectionItemsByFilter] failed get cashCollectionItems by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeCashCollectionItemsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.CashCollectionItemsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeCashCollectionItemsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeCashCollectionItemsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeCashCollectionItemsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewCashCollectionItemsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 19+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["cash_collection_session_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"cash_collection_session_id\"")
			selectedColumns["cash_collection_session_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_attempt_id\"")
			selectedColumns["payment_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["collection_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"collection_type\"")
			selectedColumns["collection_type"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["collected_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"collected_at\"")
			selectedColumns["collected_at"] = struct{}{}
		}
		if _, selected := selectedColumns["voided_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"voided_at\"")
			selectedColumns["voided_at"] = struct{}{}
		}
		if _, selected := selectedColumns["void_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"void_reason\"")
			selectedColumns["void_reason"] = struct{}{}
		}
		if _, selected := selectedColumns["notes"]; !selected {
			selectColumns = append(selectColumns, "base.\"notes\"")
			selectedColumns["notes"] = struct{}{}
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

type cashCollectionItemsFilterPlaceholder struct {
	index int
}

func (p *cashCollectionItemsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeCashCollectionItemsFilterPredicate(filterField model.FilterField, placeholders *cashCollectionItemsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewCashCollectionItemsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeCashCollectionItemsFilterSQLExpr(spec)
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

func composeCashCollectionItemsFilterGroup(group model.FilterGroup, placeholders *cashCollectionItemsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeCashCollectionItemsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeCashCollectionItemsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeCashCollectionItemsFilterWhereQueries(filter model.Filter, placeholders *cashCollectionItemsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeCashCollectionItemsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeCashCollectionItemsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeCashCollectionItemsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateCashCollectionItemsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeCashCollectionItemsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeCashCollectionItemsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := cashCollectionItemsFilterPlaceholder{index: 1}
	whereQueries, err := composeCashCollectionItemsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewCashCollectionItemsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeCashCollectionItemsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeCashCollectionItemsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"cash_collection_items\" base%s", strings.Join(selectColumns, ","), composeCashCollectionItemsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistCashCollectionItemsByID(ctx context.Context, primaryID model.CashCollectionItemsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeCashCollectionItemsCompositePrimaryKeyWhere([]model.CashCollectionItemsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", cashCollectionItemsQueries.selectCountCashCollectionItems, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCashCollectionItemsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveCashCollectionItems(ctx context.Context, selectFields ...CashCollectionItemsField) (cashCollectionItemsList model.CashCollectionItemsList, err error) {
	var (
		defaultCashCollectionItemsSelectFields = defaultCashCollectionItemsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultCashCollectionItemsSelectFields = composeCashCollectionItemsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(cashCollectionItemsQueries.selectCashCollectionItems, defaultCashCollectionItemsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &cashCollectionItemsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCashCollectionItems] failed get cashCollectionItems list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveCashCollectionItemsByID(ctx context.Context, primaryID model.CashCollectionItemsPrimaryID, selectFields ...CashCollectionItemsField) (cashCollectionItems model.CashCollectionItems, err error) {
	var (
		defaultCashCollectionItemsSelectFields = defaultCashCollectionItemsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultCashCollectionItemsSelectFields = composeCashCollectionItemsSelectFields(selectFields...)
	}
	whereQry, params := composeCashCollectionItemsCompositePrimaryKeyWhere([]model.CashCollectionItemsPrimaryID{primaryID})
	query := fmt.Sprintf(cashCollectionItemsQueries.selectCashCollectionItems+" WHERE "+whereQry, defaultCashCollectionItemsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &cashCollectionItems, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("cashCollectionItems with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveCashCollectionItemsByID] failed get cashCollectionItems")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateCashCollectionItemsByID(ctx context.Context, primaryID model.CashCollectionItemsPrimaryID, cashCollectionItems *model.CashCollectionItems, cashCollectionItemsUpdateFields ...CashCollectionItemsUpdateField) (err error) {
	exists, err := repo.IsExistCashCollectionItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionItems] failed checking cashCollectionItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionItems with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if cashCollectionItems == nil {
		if len(cashCollectionItemsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateCashCollectionItemsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		cashCollectionItems = &model.CashCollectionItems{}
	}
	var (
		defaultCashCollectionItemsUpdateFields = defaultCashCollectionItemsUpdateFields(*cashCollectionItems)
		tempUpdateField                        CashCollectionItemsUpdateFieldList
		selectFields                           = NewCashCollectionItemsSelectFields()
	)
	if len(cashCollectionItemsUpdateFields) > 0 {
		for _, updateField := range cashCollectionItemsUpdateFields {
			if updateField.cashCollectionItemsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultCashCollectionItemsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeCashCollectionItemsCompositePrimaryKeyWhere([]model.CashCollectionItemsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsCashCollectionItemsCommand(defaultCashCollectionItemsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(cashCollectionItemsQueries.updateCashCollectionItems+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionItems] error when try to update cashCollectionItems by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateCashCollectionItemsByFilter(ctx context.Context, filter model.Filter, cashCollectionItemsUpdateFields ...CashCollectionItemsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(cashCollectionItemsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields CashCollectionItemsUpdateFieldList
		selectFields = NewCashCollectionItemsSelectFields()
	)
	for _, updateField := range cashCollectionItemsUpdateFields {
		if updateField.cashCollectionItemsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsCashCollectionItemsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := cashCollectionItemsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeCashCollectionItemsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"cash_collection_items\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionItemsByFilter] error when try to update cashCollectionItems by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionItemsByFilter] failed get rows affected")
	}
	return
}

var (
	cashCollectionItemsQueries = struct {
		selectCashCollectionItems      string
		selectCountCashCollectionItems string
		deleteCashCollectionItems      string
		updateCashCollectionItems      string
		insertCashCollectionItems      string
	}{
		selectCashCollectionItems:      "SELECT %s FROM \"cash_collection_items\"",
		selectCountCashCollectionItems: "SELECT COUNT(\"id\") FROM \"cash_collection_items\"",
		deleteCashCollectionItems:      "DELETE FROM \"cash_collection_items\"",
		updateCashCollectionItems:      "UPDATE \"cash_collection_items\" SET %s ",
		insertCashCollectionItems:      "INSERT INTO \"cash_collection_items\" %s VALUES %s",
	}
)

type CashCollectionItemsRepository interface {
	CreateCashCollectionItems(ctx context.Context, cashCollectionItems *model.CashCollectionItems, fieldsInsert ...CashCollectionItemsField) error
	BulkCreateCashCollectionItems(ctx context.Context, cashCollectionItemsList []*model.CashCollectionItems, fieldsInsert ...CashCollectionItemsField) error
	ResolveCashCollectionItems(ctx context.Context, selectFields ...CashCollectionItemsField) (model.CashCollectionItemsList, error)
	ResolveCashCollectionItemsByID(ctx context.Context, primaryID model.CashCollectionItemsPrimaryID, selectFields ...CashCollectionItemsField) (model.CashCollectionItems, error)
	UpdateCashCollectionItemsByID(ctx context.Context, id model.CashCollectionItemsPrimaryID, cashCollectionItems *model.CashCollectionItems, cashCollectionItemsUpdateFields ...CashCollectionItemsUpdateField) error
	UpdateCashCollectionItemsByFilter(ctx context.Context, filter model.Filter, cashCollectionItemsUpdateFields ...CashCollectionItemsUpdateField) (rowsAffected int64, err error)
	BulkUpdateCashCollectionItems(ctx context.Context, cashCollectionItemsListMap map[model.CashCollectionItemsPrimaryID]*model.CashCollectionItems, CashCollectionItemssMapUpdateFieldsRequest map[model.CashCollectionItemsPrimaryID]CashCollectionItemsUpdateFieldList) (err error)
	DeleteCashCollectionItemsByID(ctx context.Context, id model.CashCollectionItemsPrimaryID) error
	BulkDeleteCashCollectionItemsByIDs(ctx context.Context, ids []model.CashCollectionItemsPrimaryID) error
	ResolveCashCollectionItemsByFilter(ctx context.Context, filter model.Filter) (result []model.CashCollectionItemsFilterResult, err error)
	IsExistCashCollectionItemsByIDs(ctx context.Context, ids []model.CashCollectionItemsPrimaryID) (exists bool, notFoundIds []model.CashCollectionItemsPrimaryID, err error)
	IsExistCashCollectionItemsByID(ctx context.Context, id model.CashCollectionItemsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
