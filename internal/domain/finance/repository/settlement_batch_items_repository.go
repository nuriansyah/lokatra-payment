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

func composeInsertFieldsAndParamsSettlementBatchItems(settlementBatchItemsList []model.SettlementBatchItems, fieldsInsert ...SettlementBatchItemsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewSettlementBatchItemsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, settlementBatchItems := range settlementBatchItemsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, settlementBatchItems.Id)
			case selectField.SettlementBatchId():
				args = append(args, settlementBatchItems.SettlementBatchId)
			case selectField.SourceType():
				args = append(args, settlementBatchItems.SourceType)
			case selectField.SourceId():
				args = append(args, settlementBatchItems.SourceId)
			case selectField.MerchantPartyId():
				args = append(args, settlementBatchItems.MerchantPartyId)
			case selectField.CurrencyCode():
				args = append(args, settlementBatchItems.CurrencyCode)
			case selectField.GrossAmount():
				args = append(args, settlementBatchItems.GrossAmount)
			case selectField.FeeAmount():
				args = append(args, settlementBatchItems.FeeAmount)
			case selectField.TaxAmount():
				args = append(args, settlementBatchItems.TaxAmount)
			case selectField.ReserveAmount():
				args = append(args, settlementBatchItems.ReserveAmount)
			case selectField.NetAmount():
				args = append(args, settlementBatchItems.NetAmount)
			case selectField.LinkedJournalEntryId():
				args = append(args, settlementBatchItems.LinkedJournalEntryId)
			case selectField.ItemStatus():
				args = append(args, settlementBatchItems.ItemStatus)
			case selectField.Metadata():
				args = append(args, settlementBatchItems.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, settlementBatchItems.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, settlementBatchItems.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, settlementBatchItems.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, settlementBatchItems.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, settlementBatchItems.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, settlementBatchItems.MetaDeletedBy)

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

func composeSettlementBatchItemsCompositePrimaryKeyWhere(primaryIDs []model.SettlementBatchItemsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"settlement_batch_items\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultSettlementBatchItemsSelectFields() string {
	fields := NewSettlementBatchItemsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeSettlementBatchItemsSelectFields(selectFields ...SettlementBatchItemsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type SettlementBatchItemsField string
type SettlementBatchItemsFieldList []SettlementBatchItemsField

type SettlementBatchItemsSelectFields struct {
}

func (ss SettlementBatchItemsSelectFields) Id() SettlementBatchItemsField {
	return SettlementBatchItemsField("id")
}

func (ss SettlementBatchItemsSelectFields) SettlementBatchId() SettlementBatchItemsField {
	return SettlementBatchItemsField("settlement_batch_id")
}

func (ss SettlementBatchItemsSelectFields) SourceType() SettlementBatchItemsField {
	return SettlementBatchItemsField("source_type")
}

func (ss SettlementBatchItemsSelectFields) SourceId() SettlementBatchItemsField {
	return SettlementBatchItemsField("source_id")
}

func (ss SettlementBatchItemsSelectFields) MerchantPartyId() SettlementBatchItemsField {
	return SettlementBatchItemsField("merchant_party_id")
}

func (ss SettlementBatchItemsSelectFields) CurrencyCode() SettlementBatchItemsField {
	return SettlementBatchItemsField("currency_code")
}

func (ss SettlementBatchItemsSelectFields) GrossAmount() SettlementBatchItemsField {
	return SettlementBatchItemsField("gross_amount")
}

func (ss SettlementBatchItemsSelectFields) FeeAmount() SettlementBatchItemsField {
	return SettlementBatchItemsField("fee_amount")
}

func (ss SettlementBatchItemsSelectFields) TaxAmount() SettlementBatchItemsField {
	return SettlementBatchItemsField("tax_amount")
}

func (ss SettlementBatchItemsSelectFields) ReserveAmount() SettlementBatchItemsField {
	return SettlementBatchItemsField("reserve_amount")
}

func (ss SettlementBatchItemsSelectFields) NetAmount() SettlementBatchItemsField {
	return SettlementBatchItemsField("net_amount")
}

func (ss SettlementBatchItemsSelectFields) LinkedJournalEntryId() SettlementBatchItemsField {
	return SettlementBatchItemsField("linked_journal_entry_id")
}

func (ss SettlementBatchItemsSelectFields) ItemStatus() SettlementBatchItemsField {
	return SettlementBatchItemsField("item_status")
}

func (ss SettlementBatchItemsSelectFields) Metadata() SettlementBatchItemsField {
	return SettlementBatchItemsField("metadata")
}

func (ss SettlementBatchItemsSelectFields) MetaCreatedAt() SettlementBatchItemsField {
	return SettlementBatchItemsField("meta_created_at")
}

func (ss SettlementBatchItemsSelectFields) MetaCreatedBy() SettlementBatchItemsField {
	return SettlementBatchItemsField("meta_created_by")
}

func (ss SettlementBatchItemsSelectFields) MetaUpdatedAt() SettlementBatchItemsField {
	return SettlementBatchItemsField("meta_updated_at")
}

func (ss SettlementBatchItemsSelectFields) MetaUpdatedBy() SettlementBatchItemsField {
	return SettlementBatchItemsField("meta_updated_by")
}

func (ss SettlementBatchItemsSelectFields) MetaDeletedAt() SettlementBatchItemsField {
	return SettlementBatchItemsField("meta_deleted_at")
}

func (ss SettlementBatchItemsSelectFields) MetaDeletedBy() SettlementBatchItemsField {
	return SettlementBatchItemsField("meta_deleted_by")
}

func (ss SettlementBatchItemsSelectFields) All() SettlementBatchItemsFieldList {
	return []SettlementBatchItemsField{
		ss.Id(),
		ss.SettlementBatchId(),
		ss.SourceType(),
		ss.SourceId(),
		ss.MerchantPartyId(),
		ss.CurrencyCode(),
		ss.GrossAmount(),
		ss.FeeAmount(),
		ss.TaxAmount(),
		ss.ReserveAmount(),
		ss.NetAmount(),
		ss.LinkedJournalEntryId(),
		ss.ItemStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewSettlementBatchItemsSelectFields() SettlementBatchItemsSelectFields {
	return SettlementBatchItemsSelectFields{}
}

type SettlementBatchItemsUpdateFieldOption struct {
	useIncrement bool
}
type SettlementBatchItemsUpdateField struct {
	settlementBatchItemsField SettlementBatchItemsField
	opt                       SettlementBatchItemsUpdateFieldOption
	value                     interface{}
}
type SettlementBatchItemsUpdateFieldList []SettlementBatchItemsUpdateField

func defaultSettlementBatchItemsUpdateFieldOption() SettlementBatchItemsUpdateFieldOption {
	return SettlementBatchItemsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementSettlementBatchItemsOption(useIncrement bool) func(*SettlementBatchItemsUpdateFieldOption) {
	return func(pcufo *SettlementBatchItemsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewSettlementBatchItemsUpdateField(field SettlementBatchItemsField, val interface{}, opts ...func(*SettlementBatchItemsUpdateFieldOption)) SettlementBatchItemsUpdateField {
	defaultOpt := defaultSettlementBatchItemsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return SettlementBatchItemsUpdateField{
		settlementBatchItemsField: field,
		value:                     val,
		opt:                       defaultOpt,
	}
}
func defaultSettlementBatchItemsUpdateFields(settlementBatchItems model.SettlementBatchItems) (settlementBatchItemsUpdateFieldList SettlementBatchItemsUpdateFieldList) {
	selectFields := NewSettlementBatchItemsSelectFields()
	settlementBatchItemsUpdateFieldList = append(settlementBatchItemsUpdateFieldList,
		NewSettlementBatchItemsUpdateField(selectFields.Id(), settlementBatchItems.Id),
		NewSettlementBatchItemsUpdateField(selectFields.SettlementBatchId(), settlementBatchItems.SettlementBatchId),
		NewSettlementBatchItemsUpdateField(selectFields.SourceType(), settlementBatchItems.SourceType),
		NewSettlementBatchItemsUpdateField(selectFields.SourceId(), settlementBatchItems.SourceId),
		NewSettlementBatchItemsUpdateField(selectFields.MerchantPartyId(), settlementBatchItems.MerchantPartyId),
		NewSettlementBatchItemsUpdateField(selectFields.CurrencyCode(), settlementBatchItems.CurrencyCode),
		NewSettlementBatchItemsUpdateField(selectFields.GrossAmount(), settlementBatchItems.GrossAmount),
		NewSettlementBatchItemsUpdateField(selectFields.FeeAmount(), settlementBatchItems.FeeAmount),
		NewSettlementBatchItemsUpdateField(selectFields.TaxAmount(), settlementBatchItems.TaxAmount),
		NewSettlementBatchItemsUpdateField(selectFields.ReserveAmount(), settlementBatchItems.ReserveAmount),
		NewSettlementBatchItemsUpdateField(selectFields.NetAmount(), settlementBatchItems.NetAmount),
		NewSettlementBatchItemsUpdateField(selectFields.LinkedJournalEntryId(), settlementBatchItems.LinkedJournalEntryId),
		NewSettlementBatchItemsUpdateField(selectFields.ItemStatus(), settlementBatchItems.ItemStatus),
		NewSettlementBatchItemsUpdateField(selectFields.Metadata(), settlementBatchItems.Metadata),
		NewSettlementBatchItemsUpdateField(selectFields.MetaCreatedAt(), settlementBatchItems.MetaCreatedAt),
		NewSettlementBatchItemsUpdateField(selectFields.MetaCreatedBy(), settlementBatchItems.MetaCreatedBy),
		NewSettlementBatchItemsUpdateField(selectFields.MetaUpdatedAt(), settlementBatchItems.MetaUpdatedAt),
		NewSettlementBatchItemsUpdateField(selectFields.MetaUpdatedBy(), settlementBatchItems.MetaUpdatedBy),
		NewSettlementBatchItemsUpdateField(selectFields.MetaDeletedAt(), settlementBatchItems.MetaDeletedAt),
		NewSettlementBatchItemsUpdateField(selectFields.MetaDeletedBy(), settlementBatchItems.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsSettlementBatchItemsCommand(settlementBatchItemsUpdateFieldList SettlementBatchItemsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range settlementBatchItemsUpdateFieldList {
		field := string(updateField.settlementBatchItemsField)
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

func (repo *RepositoryImpl) BulkCreateSettlementBatchItems(ctx context.Context, settlementBatchItemsList []*model.SettlementBatchItems, fieldsInsert ...SettlementBatchItemsField) (err error) {
	var (
		fieldsStr                     string
		valueListStr                  []string
		argsList                      []interface{}
		primaryIds                    []model.SettlementBatchItemsPrimaryID
		settlementBatchItemsValueList []model.SettlementBatchItems
	)

	if len(fieldsInsert) == 0 {
		selectField := NewSettlementBatchItemsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, settlementBatchItems := range settlementBatchItemsList {

		primaryIds = append(primaryIds, settlementBatchItems.ToSettlementBatchItemsPrimaryID())

		settlementBatchItemsValueList = append(settlementBatchItemsValueList, *settlementBatchItems)
	}

	_, notFoundIds, err := repo.IsExistSettlementBatchItemsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementBatchItems] failed checking settlementBatchItems whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.SettlementBatchItemsPrimaryID{}
		mapNotFoundIds := map[model.SettlementBatchItemsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "settlementBatchItems", fmt.Sprintf("settlementBatchItems with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsSettlementBatchItems(settlementBatchItemsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(settlementBatchItemsQueries.insertSettlementBatchItems, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementBatchItems] failed exec create settlementBatchItems query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteSettlementBatchItemsByIDs(ctx context.Context, primaryIDs []model.SettlementBatchItemsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistSettlementBatchItemsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementBatchItemsByIDs] failed checking settlementBatchItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatchItems with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_batch_items\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(settlementBatchItemsQueries.deleteSettlementBatchItems + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementBatchItemsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementBatchItemsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistSettlementBatchItemsByIDs(ctx context.Context, ids []model.SettlementBatchItemsPrimaryID) (exists bool, notFoundIds []model.SettlementBatchItemsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_batch_items\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(settlementBatchItemsQueries.selectSettlementBatchItems, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementBatchItemsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.SettlementBatchItemsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementBatchItemsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.SettlementBatchItemsPrimaryID]bool{}
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

// BulkUpdateSettlementBatchItems is used to bulk update settlementBatchItems, by default it will update all field
// if want to update specific field, then fill settlementBatchItemssMapUpdateFieldsRequest else please fill settlementBatchItemssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateSettlementBatchItems(ctx context.Context, settlementBatchItemssMap map[model.SettlementBatchItemsPrimaryID]*model.SettlementBatchItems, settlementBatchItemssMapUpdateFieldsRequest map[model.SettlementBatchItemsPrimaryID]SettlementBatchItemsUpdateFieldList) (err error) {
	if len(settlementBatchItemssMap) == 0 && len(settlementBatchItemssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		settlementBatchItemssMapUpdateField map[model.SettlementBatchItemsPrimaryID]SettlementBatchItemsUpdateFieldList = map[model.SettlementBatchItemsPrimaryID]SettlementBatchItemsUpdateFieldList{}
		asTableValues                       string                                                                      = "myvalues"
	)

	if len(settlementBatchItemssMap) > 0 {
		for id, settlementBatchItems := range settlementBatchItemssMap {
			if settlementBatchItems == nil {
				log.Error().Err(err).Msg("[BulkUpdateSettlementBatchItems] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			settlementBatchItemssMapUpdateField[id] = defaultSettlementBatchItemsUpdateFields(*settlementBatchItems)
		}
	} else {
		settlementBatchItemssMapUpdateField = settlementBatchItemssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateSettlementBatchItemsQuery(settlementBatchItemssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistSettlementBatchItemsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementBatchItems] failed checking settlementBatchItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatchItems with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeSettlementBatchItemsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"settlement_batch_items\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementBatchItems] failed exec query")
	}
	return
}

type SettlementBatchItemsFieldParameter struct {
	param string
	args  []interface{}
}

func NewSettlementBatchItemsFieldParameter(param string, args ...interface{}) SettlementBatchItemsFieldParameter {
	return SettlementBatchItemsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateSettlementBatchItemsQuery(mapSettlementBatchItemss map[model.SettlementBatchItemsPrimaryID]SettlementBatchItemsUpdateFieldList, asTableValues string) (primaryIDs []model.SettlementBatchItemsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.SettlementBatchItemsPrimaryID]map[string]interface{}{}
	settlementBatchItemsSelectFields := NewSettlementBatchItemsSelectFields()
	for id, updateFields := range mapSettlementBatchItemss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.settlementBatchItemsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapSettlementBatchItemss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetSettlementBatchItemsFieldType(updateField.settlementBatchItemsField)))
			args = append(args, fields[string(updateField.settlementBatchItemsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.settlementBatchItemsField))
		if updateField.settlementBatchItemsField == settlementBatchItemsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.settlementBatchItemsField, asTableValues, updateField.settlementBatchItemsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.settlementBatchItemsField,
				"\"settlement_batch_items\"", updateField.settlementBatchItemsField,
				asTableValues, updateField.settlementBatchItemsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeSettlementBatchItemsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.SettlementBatchItemsPrimaryID, asTableValue string) (whereQry string) {
	settlementBatchItemsSelectFields := NewSettlementBatchItemsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"settlement_batch_items\".\"id\" = %s.\"id\"::"+GetSettlementBatchItemsFieldType(settlementBatchItemsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetSettlementBatchItemsFieldType(settlementBatchItemsField SettlementBatchItemsField) string {
	selectSettlementBatchItemsFields := NewSettlementBatchItemsSelectFields()
	switch settlementBatchItemsField {

	case selectSettlementBatchItemsFields.Id():
		return "uuid"

	case selectSettlementBatchItemsFields.SettlementBatchId():
		return "uuid"

	case selectSettlementBatchItemsFields.SourceType():
		return "text"

	case selectSettlementBatchItemsFields.SourceId():
		return "uuid"

	case selectSettlementBatchItemsFields.MerchantPartyId():
		return "uuid"

	case selectSettlementBatchItemsFields.CurrencyCode():
		return "text"

	case selectSettlementBatchItemsFields.GrossAmount():
		return "numeric"

	case selectSettlementBatchItemsFields.FeeAmount():
		return "numeric"

	case selectSettlementBatchItemsFields.TaxAmount():
		return "numeric"

	case selectSettlementBatchItemsFields.ReserveAmount():
		return "numeric"

	case selectSettlementBatchItemsFields.NetAmount():
		return "numeric"

	case selectSettlementBatchItemsFields.LinkedJournalEntryId():
		return "uuid"

	case selectSettlementBatchItemsFields.ItemStatus():
		return "item_status_enum"

	case selectSettlementBatchItemsFields.Metadata():
		return "jsonb"

	case selectSettlementBatchItemsFields.MetaCreatedAt():
		return "timestamptz"

	case selectSettlementBatchItemsFields.MetaCreatedBy():
		return "uuid"

	case selectSettlementBatchItemsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectSettlementBatchItemsFields.MetaUpdatedBy():
		return "uuid"

	case selectSettlementBatchItemsFields.MetaDeletedAt():
		return "timestamptz"

	case selectSettlementBatchItemsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateSettlementBatchItems(ctx context.Context, settlementBatchItems *model.SettlementBatchItems, fieldsInsert ...SettlementBatchItemsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewSettlementBatchItemsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.SettlementBatchItemsPrimaryID{
		Id: settlementBatchItems.Id,
	}
	exists, err := repo.IsExistSettlementBatchItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementBatchItems] failed checking settlementBatchItems whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "settlementBatchItems", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsSettlementBatchItems([]model.SettlementBatchItems{*settlementBatchItems}, fieldsInsert...)
	commandQuery := fmt.Sprintf(settlementBatchItemsQueries.insertSettlementBatchItems, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementBatchItems] failed exec create settlementBatchItems query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteSettlementBatchItemsByID(ctx context.Context, primaryID model.SettlementBatchItemsPrimaryID) (err error) {
	exists, err := repo.IsExistSettlementBatchItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementBatchItemsByID] failed checking settlementBatchItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatchItems with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeSettlementBatchItemsCompositePrimaryKeyWhere([]model.SettlementBatchItemsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(settlementBatchItemsQueries.deleteSettlementBatchItems + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementBatchItemsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementBatchItemsByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementBatchItemsFilterResult, err error) {
	query, args, err := composeSettlementBatchItemsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementBatchItemsByFilter] failed compose settlementBatchItems filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementBatchItemsByFilter] failed get settlementBatchItems by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeSettlementBatchItemsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.SettlementBatchItemsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeSettlementBatchItemsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeSettlementBatchItemsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeSettlementBatchItemsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewSettlementBatchItemsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["settlement_batch_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"settlement_batch_id\"")
			selectedColumns["settlement_batch_id"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
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
		if _, selected := selectedColumns["tax_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_amount\"")
			selectedColumns["tax_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["reserve_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"reserve_amount\"")
			selectedColumns["reserve_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["net_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"net_amount\"")
			selectedColumns["net_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["linked_journal_entry_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"linked_journal_entry_id\"")
			selectedColumns["linked_journal_entry_id"] = struct{}{}
		}
		if _, selected := selectedColumns["item_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"item_status\"")
			selectedColumns["item_status"] = struct{}{}
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

type settlementBatchItemsFilterPlaceholder struct {
	index int
}

func (p *settlementBatchItemsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeSettlementBatchItemsFilterPredicate(filterField model.FilterField, placeholders *settlementBatchItemsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewSettlementBatchItemsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeSettlementBatchItemsFilterSQLExpr(spec)
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

func composeSettlementBatchItemsFilterGroup(group model.FilterGroup, placeholders *settlementBatchItemsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeSettlementBatchItemsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeSettlementBatchItemsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeSettlementBatchItemsFilterWhereQueries(filter model.Filter, placeholders *settlementBatchItemsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeSettlementBatchItemsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeSettlementBatchItemsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeSettlementBatchItemsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateSettlementBatchItemsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeSettlementBatchItemsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeSettlementBatchItemsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := settlementBatchItemsFilterPlaceholder{index: 1}
	whereQueries, err := composeSettlementBatchItemsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewSettlementBatchItemsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeSettlementBatchItemsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeSettlementBatchItemsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"settlement_batch_items\" base%s", strings.Join(selectColumns, ","), composeSettlementBatchItemsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistSettlementBatchItemsByID(ctx context.Context, primaryID model.SettlementBatchItemsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeSettlementBatchItemsCompositePrimaryKeyWhere([]model.SettlementBatchItemsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", settlementBatchItemsQueries.selectCountSettlementBatchItems, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementBatchItemsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementBatchItems(ctx context.Context, selectFields ...SettlementBatchItemsField) (settlementBatchItemsList model.SettlementBatchItemsList, err error) {
	var (
		defaultSettlementBatchItemsSelectFields = defaultSettlementBatchItemsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementBatchItemsSelectFields = composeSettlementBatchItemsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(settlementBatchItemsQueries.selectSettlementBatchItems, defaultSettlementBatchItemsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &settlementBatchItemsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementBatchItems] failed get settlementBatchItems list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementBatchItemsByID(ctx context.Context, primaryID model.SettlementBatchItemsPrimaryID, selectFields ...SettlementBatchItemsField) (settlementBatchItems model.SettlementBatchItems, err error) {
	var (
		defaultSettlementBatchItemsSelectFields = defaultSettlementBatchItemsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementBatchItemsSelectFields = composeSettlementBatchItemsSelectFields(selectFields...)
	}
	whereQry, params := composeSettlementBatchItemsCompositePrimaryKeyWhere([]model.SettlementBatchItemsPrimaryID{primaryID})
	query := fmt.Sprintf(settlementBatchItemsQueries.selectSettlementBatchItems+" WHERE "+whereQry, defaultSettlementBatchItemsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &settlementBatchItems, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("settlementBatchItems with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveSettlementBatchItemsByID] failed get settlementBatchItems")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateSettlementBatchItemsByID(ctx context.Context, primaryID model.SettlementBatchItemsPrimaryID, settlementBatchItems *model.SettlementBatchItems, settlementBatchItemsUpdateFields ...SettlementBatchItemsUpdateField) (err error) {
	exists, err := repo.IsExistSettlementBatchItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatchItems] failed checking settlementBatchItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatchItems with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if settlementBatchItems == nil {
		if len(settlementBatchItemsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateSettlementBatchItemsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		settlementBatchItems = &model.SettlementBatchItems{}
	}
	var (
		defaultSettlementBatchItemsUpdateFields = defaultSettlementBatchItemsUpdateFields(*settlementBatchItems)
		tempUpdateField                         SettlementBatchItemsUpdateFieldList
		selectFields                            = NewSettlementBatchItemsSelectFields()
	)
	if len(settlementBatchItemsUpdateFields) > 0 {
		for _, updateField := range settlementBatchItemsUpdateFields {
			if updateField.settlementBatchItemsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultSettlementBatchItemsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeSettlementBatchItemsCompositePrimaryKeyWhere([]model.SettlementBatchItemsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsSettlementBatchItemsCommand(defaultSettlementBatchItemsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(settlementBatchItemsQueries.updateSettlementBatchItems+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatchItems] error when try to update settlementBatchItems by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateSettlementBatchItemsByFilter(ctx context.Context, filter model.Filter, settlementBatchItemsUpdateFields ...SettlementBatchItemsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(settlementBatchItemsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields SettlementBatchItemsUpdateFieldList
		selectFields = NewSettlementBatchItemsSelectFields()
	)
	for _, updateField := range settlementBatchItemsUpdateFields {
		if updateField.settlementBatchItemsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsSettlementBatchItemsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := settlementBatchItemsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeSettlementBatchItemsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"settlement_batch_items\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatchItemsByFilter] error when try to update settlementBatchItems by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatchItemsByFilter] failed get rows affected")
	}
	return
}

var (
	settlementBatchItemsQueries = struct {
		selectSettlementBatchItems      string
		selectCountSettlementBatchItems string
		deleteSettlementBatchItems      string
		updateSettlementBatchItems      string
		insertSettlementBatchItems      string
	}{
		selectSettlementBatchItems:      "SELECT %s FROM \"settlement_batch_items\"",
		selectCountSettlementBatchItems: "SELECT COUNT(\"id\") FROM \"settlement_batch_items\"",
		deleteSettlementBatchItems:      "DELETE FROM \"settlement_batch_items\"",
		updateSettlementBatchItems:      "UPDATE \"settlement_batch_items\" SET %s ",
		insertSettlementBatchItems:      "INSERT INTO \"settlement_batch_items\" %s VALUES %s",
	}
)

type SettlementBatchItemsRepository interface {
	CreateSettlementBatchItems(ctx context.Context, settlementBatchItems *model.SettlementBatchItems, fieldsInsert ...SettlementBatchItemsField) error
	BulkCreateSettlementBatchItems(ctx context.Context, settlementBatchItemsList []*model.SettlementBatchItems, fieldsInsert ...SettlementBatchItemsField) error
	ResolveSettlementBatchItems(ctx context.Context, selectFields ...SettlementBatchItemsField) (model.SettlementBatchItemsList, error)
	ResolveSettlementBatchItemsByID(ctx context.Context, primaryID model.SettlementBatchItemsPrimaryID, selectFields ...SettlementBatchItemsField) (model.SettlementBatchItems, error)
	UpdateSettlementBatchItemsByID(ctx context.Context, id model.SettlementBatchItemsPrimaryID, settlementBatchItems *model.SettlementBatchItems, settlementBatchItemsUpdateFields ...SettlementBatchItemsUpdateField) error
	UpdateSettlementBatchItemsByFilter(ctx context.Context, filter model.Filter, settlementBatchItemsUpdateFields ...SettlementBatchItemsUpdateField) (rowsAffected int64, err error)
	BulkUpdateSettlementBatchItems(ctx context.Context, settlementBatchItemsListMap map[model.SettlementBatchItemsPrimaryID]*model.SettlementBatchItems, SettlementBatchItemssMapUpdateFieldsRequest map[model.SettlementBatchItemsPrimaryID]SettlementBatchItemsUpdateFieldList) (err error)
	DeleteSettlementBatchItemsByID(ctx context.Context, id model.SettlementBatchItemsPrimaryID) error
	BulkDeleteSettlementBatchItemsByIDs(ctx context.Context, ids []model.SettlementBatchItemsPrimaryID) error
	ResolveSettlementBatchItemsByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementBatchItemsFilterResult, err error)
	IsExistSettlementBatchItemsByIDs(ctx context.Context, ids []model.SettlementBatchItemsPrimaryID) (exists bool, notFoundIds []model.SettlementBatchItemsPrimaryID, err error)
	IsExistSettlementBatchItemsByID(ctx context.Context, id model.SettlementBatchItemsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
