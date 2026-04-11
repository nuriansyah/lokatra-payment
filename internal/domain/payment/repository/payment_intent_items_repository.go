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

func composeInsertFieldsAndParamsPaymentIntentItems(paymentIntentItemsList []model.PaymentIntentItems, fieldsInsert ...PaymentIntentItemsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentIntentItemsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentIntentItems := range paymentIntentItemsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentIntentItems.Id)
			case selectField.IntentId():
				args = append(args, paymentIntentItems.IntentId)
			case selectField.ProductId():
				args = append(args, paymentIntentItems.ProductId)
			case selectField.ProductType():
				args = append(args, paymentIntentItems.ProductType)
			case selectField.ProductName():
				args = append(args, paymentIntentItems.ProductName)
			case selectField.Quantity():
				args = append(args, paymentIntentItems.Quantity)
			case selectField.UnitPrice():
				args = append(args, paymentIntentItems.UnitPrice)
			case selectField.DiscountAmount():
				args = append(args, paymentIntentItems.DiscountAmount)
			case selectField.TotalPrice():
				args = append(args, paymentIntentItems.TotalPrice)
			case selectField.SellerMerchantId():
				args = append(args, paymentIntentItems.SellerMerchantId)
			case selectField.Metadata():
				args = append(args, paymentIntentItems.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentIntentItems.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentIntentItems.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentIntentItems.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentIntentItems.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentIntentItems.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentIntentItems.MetaDeletedBy)

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

func composePaymentIntentItemsCompositePrimaryKeyWhere(primaryIDs []model.PaymentIntentItemsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_intent_items\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentIntentItemsSelectFields() string {
	fields := NewPaymentIntentItemsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentIntentItemsSelectFields(selectFields ...PaymentIntentItemsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentIntentItemsField string
type PaymentIntentItemsFieldList []PaymentIntentItemsField

type PaymentIntentItemsSelectFields struct {
}

func (ss PaymentIntentItemsSelectFields) Id() PaymentIntentItemsField {
	return PaymentIntentItemsField("id")
}

func (ss PaymentIntentItemsSelectFields) IntentId() PaymentIntentItemsField {
	return PaymentIntentItemsField("intent_id")
}

func (ss PaymentIntentItemsSelectFields) ProductId() PaymentIntentItemsField {
	return PaymentIntentItemsField("product_id")
}

func (ss PaymentIntentItemsSelectFields) ProductType() PaymentIntentItemsField {
	return PaymentIntentItemsField("product_type")
}

func (ss PaymentIntentItemsSelectFields) ProductName() PaymentIntentItemsField {
	return PaymentIntentItemsField("product_name")
}

func (ss PaymentIntentItemsSelectFields) Quantity() PaymentIntentItemsField {
	return PaymentIntentItemsField("quantity")
}

func (ss PaymentIntentItemsSelectFields) UnitPrice() PaymentIntentItemsField {
	return PaymentIntentItemsField("unit_price")
}

func (ss PaymentIntentItemsSelectFields) DiscountAmount() PaymentIntentItemsField {
	return PaymentIntentItemsField("discount_amount")
}

func (ss PaymentIntentItemsSelectFields) TotalPrice() PaymentIntentItemsField {
	return PaymentIntentItemsField("total_price")
}

func (ss PaymentIntentItemsSelectFields) SellerMerchantId() PaymentIntentItemsField {
	return PaymentIntentItemsField("seller_merchant_id")
}

func (ss PaymentIntentItemsSelectFields) Metadata() PaymentIntentItemsField {
	return PaymentIntentItemsField("metadata")
}

func (ss PaymentIntentItemsSelectFields) MetaCreatedAt() PaymentIntentItemsField {
	return PaymentIntentItemsField("meta_created_at")
}

func (ss PaymentIntentItemsSelectFields) MetaCreatedBy() PaymentIntentItemsField {
	return PaymentIntentItemsField("meta_created_by")
}

func (ss PaymentIntentItemsSelectFields) MetaUpdatedAt() PaymentIntentItemsField {
	return PaymentIntentItemsField("meta_updated_at")
}

func (ss PaymentIntentItemsSelectFields) MetaUpdatedBy() PaymentIntentItemsField {
	return PaymentIntentItemsField("meta_updated_by")
}

func (ss PaymentIntentItemsSelectFields) MetaDeletedAt() PaymentIntentItemsField {
	return PaymentIntentItemsField("meta_deleted_at")
}

func (ss PaymentIntentItemsSelectFields) MetaDeletedBy() PaymentIntentItemsField {
	return PaymentIntentItemsField("meta_deleted_by")
}

func (ss PaymentIntentItemsSelectFields) All() PaymentIntentItemsFieldList {
	return []PaymentIntentItemsField{
		ss.Id(),
		ss.IntentId(),
		ss.ProductId(),
		ss.ProductType(),
		ss.ProductName(),
		ss.Quantity(),
		ss.UnitPrice(),
		ss.DiscountAmount(),
		ss.TotalPrice(),
		ss.SellerMerchantId(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentIntentItemsSelectFields() PaymentIntentItemsSelectFields {
	return PaymentIntentItemsSelectFields{}
}

type PaymentIntentItemsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentIntentItemsUpdateField struct {
	paymentIntentItemsField PaymentIntentItemsField
	opt                     PaymentIntentItemsUpdateFieldOption
	value                   interface{}
}
type PaymentIntentItemsUpdateFieldList []PaymentIntentItemsUpdateField

func defaultPaymentIntentItemsUpdateFieldOption() PaymentIntentItemsUpdateFieldOption {
	return PaymentIntentItemsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentIntentItemsOption(useIncrement bool) func(*PaymentIntentItemsUpdateFieldOption) {
	return func(pcufo *PaymentIntentItemsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentIntentItemsUpdateField(field PaymentIntentItemsField, val interface{}, opts ...func(*PaymentIntentItemsUpdateFieldOption)) PaymentIntentItemsUpdateField {
	defaultOpt := defaultPaymentIntentItemsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentIntentItemsUpdateField{
		paymentIntentItemsField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultPaymentIntentItemsUpdateFields(paymentIntentItems model.PaymentIntentItems) (paymentIntentItemsUpdateFieldList PaymentIntentItemsUpdateFieldList) {
	selectFields := NewPaymentIntentItemsSelectFields()
	paymentIntentItemsUpdateFieldList = append(paymentIntentItemsUpdateFieldList,
		NewPaymentIntentItemsUpdateField(selectFields.Id(), paymentIntentItems.Id),
		NewPaymentIntentItemsUpdateField(selectFields.IntentId(), paymentIntentItems.IntentId),
		NewPaymentIntentItemsUpdateField(selectFields.ProductId(), paymentIntentItems.ProductId),
		NewPaymentIntentItemsUpdateField(selectFields.ProductType(), paymentIntentItems.ProductType),
		NewPaymentIntentItemsUpdateField(selectFields.ProductName(), paymentIntentItems.ProductName),
		NewPaymentIntentItemsUpdateField(selectFields.Quantity(), paymentIntentItems.Quantity),
		NewPaymentIntentItemsUpdateField(selectFields.UnitPrice(), paymentIntentItems.UnitPrice),
		NewPaymentIntentItemsUpdateField(selectFields.DiscountAmount(), paymentIntentItems.DiscountAmount),
		NewPaymentIntentItemsUpdateField(selectFields.TotalPrice(), paymentIntentItems.TotalPrice),
		NewPaymentIntentItemsUpdateField(selectFields.SellerMerchantId(), paymentIntentItems.SellerMerchantId),
		NewPaymentIntentItemsUpdateField(selectFields.Metadata(), paymentIntentItems.Metadata),
		NewPaymentIntentItemsUpdateField(selectFields.MetaCreatedAt(), paymentIntentItems.MetaCreatedAt),
		NewPaymentIntentItemsUpdateField(selectFields.MetaCreatedBy(), paymentIntentItems.MetaCreatedBy),
		NewPaymentIntentItemsUpdateField(selectFields.MetaUpdatedAt(), paymentIntentItems.MetaUpdatedAt),
		NewPaymentIntentItemsUpdateField(selectFields.MetaUpdatedBy(), paymentIntentItems.MetaUpdatedBy),
		NewPaymentIntentItemsUpdateField(selectFields.MetaDeletedAt(), paymentIntentItems.MetaDeletedAt),
		NewPaymentIntentItemsUpdateField(selectFields.MetaDeletedBy(), paymentIntentItems.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentIntentItemsCommand(paymentIntentItemsUpdateFieldList PaymentIntentItemsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentIntentItemsUpdateFieldList {
		field := string(updateField.paymentIntentItemsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentIntentItems(ctx context.Context, paymentIntentItemsList []*model.PaymentIntentItems, fieldsInsert ...PaymentIntentItemsField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.PaymentIntentItemsPrimaryID
		paymentIntentItemsValueList []model.PaymentIntentItems
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentIntentItemsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentIntentItems := range paymentIntentItemsList {

		primaryIds = append(primaryIds, paymentIntentItems.ToPaymentIntentItemsPrimaryID())

		paymentIntentItemsValueList = append(paymentIntentItemsValueList, *paymentIntentItems)
	}

	_, notFoundIds, err := repo.IsExistPaymentIntentItemsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentIntentItems] failed checking paymentIntentItems whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentIntentItemsPrimaryID{}
		mapNotFoundIds := map[model.PaymentIntentItemsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentIntentItems", fmt.Sprintf("paymentIntentItems with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentIntentItems(paymentIntentItemsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentIntentItemsQueries.insertPaymentIntentItems, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentIntentItems] failed exec create paymentIntentItems query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentIntentItemsByIDs(ctx context.Context, primaryIDs []model.PaymentIntentItemsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentIntentItemsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentIntentItemsByIDs] failed checking paymentIntentItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntentItems with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_intent_items\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentIntentItemsQueries.deletePaymentIntentItems + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentIntentItemsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentIntentItemsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentIntentItemsByIDs(ctx context.Context, ids []model.PaymentIntentItemsPrimaryID) (exists bool, notFoundIds []model.PaymentIntentItemsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_intent_items\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentIntentItemsQueries.selectPaymentIntentItems, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentIntentItemsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentIntentItemsPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentIntentItemsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentIntentItemsPrimaryID]bool{}
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

// BulkUpdatePaymentIntentItems is used to bulk update paymentIntentItems, by default it will update all field
// if want to update specific field, then fill paymentIntentItemssMapUpdateFieldsRequest else please fill paymentIntentItemssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentIntentItems(ctx context.Context, paymentIntentItemssMap map[model.PaymentIntentItemsPrimaryID]*model.PaymentIntentItems, paymentIntentItemssMapUpdateFieldsRequest map[model.PaymentIntentItemsPrimaryID]PaymentIntentItemsUpdateFieldList) (err error) {
	if len(paymentIntentItemssMap) == 0 && len(paymentIntentItemssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentIntentItemssMapUpdateField map[model.PaymentIntentItemsPrimaryID]PaymentIntentItemsUpdateFieldList = map[model.PaymentIntentItemsPrimaryID]PaymentIntentItemsUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(paymentIntentItemssMap) > 0 {
		for id, paymentIntentItems := range paymentIntentItemssMap {
			if paymentIntentItems == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentIntentItems] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentIntentItemssMapUpdateField[id] = defaultPaymentIntentItemsUpdateFields(*paymentIntentItems)
		}
	} else {
		paymentIntentItemssMapUpdateField = paymentIntentItemssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentIntentItemsQuery(paymentIntentItemssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentIntentItemsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentIntentItems] failed checking paymentIntentItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntentItems with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentIntentItemsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_intent_items\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentIntentItems] failed exec query")
	}
	return
}

type PaymentIntentItemsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentIntentItemsFieldParameter(param string, args ...interface{}) PaymentIntentItemsFieldParameter {
	return PaymentIntentItemsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentIntentItemsQuery(mapPaymentIntentItemss map[model.PaymentIntentItemsPrimaryID]PaymentIntentItemsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentIntentItemsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentIntentItemsPrimaryID]map[string]interface{}{}
	paymentIntentItemsSelectFields := NewPaymentIntentItemsSelectFields()
	for id, updateFields := range mapPaymentIntentItemss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentIntentItemsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentIntentItemss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentIntentItemsFieldType(updateField.paymentIntentItemsField)))
			args = append(args, fields[string(updateField.paymentIntentItemsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentIntentItemsField))
		if updateField.paymentIntentItemsField == paymentIntentItemsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentIntentItemsField, asTableValues, updateField.paymentIntentItemsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentIntentItemsField,
				"\"payment_intent_items\"", updateField.paymentIntentItemsField,
				asTableValues, updateField.paymentIntentItemsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentIntentItemsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentIntentItemsPrimaryID, asTableValue string) (whereQry string) {
	paymentIntentItemsSelectFields := NewPaymentIntentItemsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_intent_items\".\"id\" = %s.\"id\"::"+GetPaymentIntentItemsFieldType(paymentIntentItemsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentIntentItemsFieldType(paymentIntentItemsField PaymentIntentItemsField) string {
	selectPaymentIntentItemsFields := NewPaymentIntentItemsSelectFields()
	switch paymentIntentItemsField {

	case selectPaymentIntentItemsFields.Id():
		return "uuid"

	case selectPaymentIntentItemsFields.IntentId():
		return "uuid"

	case selectPaymentIntentItemsFields.ProductId():
		return "uuid"

	case selectPaymentIntentItemsFields.ProductType():
		return "text"

	case selectPaymentIntentItemsFields.ProductName():
		return "text"

	case selectPaymentIntentItemsFields.Quantity():
		return "int4"

	case selectPaymentIntentItemsFields.UnitPrice():
		return "numeric"

	case selectPaymentIntentItemsFields.DiscountAmount():
		return "numeric"

	case selectPaymentIntentItemsFields.TotalPrice():
		return "numeric"

	case selectPaymentIntentItemsFields.SellerMerchantId():
		return "uuid"

	case selectPaymentIntentItemsFields.Metadata():
		return "jsonb"

	case selectPaymentIntentItemsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentIntentItemsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentIntentItemsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentIntentItemsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentIntentItemsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentIntentItemsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentIntentItems(ctx context.Context, paymentIntentItems *model.PaymentIntentItems, fieldsInsert ...PaymentIntentItemsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentIntentItemsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentIntentItemsPrimaryID{
		Id: paymentIntentItems.Id,
	}
	exists, err := repo.IsExistPaymentIntentItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentIntentItems] failed checking paymentIntentItems whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentIntentItems", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentIntentItems([]model.PaymentIntentItems{*paymentIntentItems}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentIntentItemsQueries.insertPaymentIntentItems, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentIntentItems] failed exec create paymentIntentItems query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentIntentItemsByID(ctx context.Context, primaryID model.PaymentIntentItemsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentIntentItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentIntentItemsByID] failed checking paymentIntentItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntentItems with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentIntentItemsCompositePrimaryKeyWhere([]model.PaymentIntentItemsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentIntentItemsQueries.deletePaymentIntentItems + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentIntentItemsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentIntentItemsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentIntentItemsFilterResult, err error) {
	query, args, err := composePaymentIntentItemsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentIntentItemsByFilter] failed compose paymentIntentItems filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentIntentItemsByFilter] failed get paymentIntentItems by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentIntentItemsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentIntentItemsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultPaymentIntentItemsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := PaymentIntentItemsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, PaymentIntentItemsField(filterSelectField))
		}
		selectFields = composePaymentIntentItemsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(paymentIntentItemsQueries.selectPaymentIntentItems, selectFields)

	if len(filter.FilterFields) > 0 {
		var (
			whereQueries []string
			whereArgs    []interface{}
		)
		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorEqual:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" = $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			case model.OperatorRange:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" BETWEEN $%d AND $%d", filterField.Field, index, index+1))
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				whereArgs = append(whereArgs, valueArray...)
				index += 2
			case model.OperatorIn:
				valueArray, ok := filterField.Value.([]interface{})
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				var placeholder []string
				for range valueArray {
					placeholder = append(placeholder, fmt.Sprintf("$%d", index))
					index++
				}
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IN (%s)", filterField.Field, strings.Join(placeholder, ",")))
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIsNull:
				value, ok := filterField.Value.(bool)
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				if value {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NULL", filterField.Field))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NOT NULL", filterField.Field))
				}
			case model.OperatorNot:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" != $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			}
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND "))
		args = append(args, whereArgs...)
	}

	sortQuery := []string{}
	for _, sort := range filter.Sorts {
		sortQuery = append(sortQuery, fmt.Sprintf("\"%s\" %s", sort.Field, sort.Order))
	}
	if len(sortQuery) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
	}
	if filter.Pagination.PageSize > 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Pagination.PageSize)
		if filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
	}

	return
}

func (repo *RepositoryImpl) IsExistPaymentIntentItemsByID(ctx context.Context, primaryID model.PaymentIntentItemsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentIntentItemsCompositePrimaryKeyWhere([]model.PaymentIntentItemsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentIntentItemsQueries.selectCountPaymentIntentItems, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentIntentItemsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentIntentItems(ctx context.Context, selectFields ...PaymentIntentItemsField) (paymentIntentItemsList model.PaymentIntentItemsList, err error) {
	var (
		defaultPaymentIntentItemsSelectFields = defaultPaymentIntentItemsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentIntentItemsSelectFields = composePaymentIntentItemsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentIntentItemsQueries.selectPaymentIntentItems, defaultPaymentIntentItemsSelectFields)

	err = repo.db.Read.Select(&paymentIntentItemsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentIntentItems] failed get paymentIntentItems list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentIntentItemsByID(ctx context.Context, primaryID model.PaymentIntentItemsPrimaryID, selectFields ...PaymentIntentItemsField) (paymentIntentItems model.PaymentIntentItems, err error) {
	var (
		defaultPaymentIntentItemsSelectFields = defaultPaymentIntentItemsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentIntentItemsSelectFields = composePaymentIntentItemsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentIntentItemsCompositePrimaryKeyWhere([]model.PaymentIntentItemsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentIntentItemsQueries.selectPaymentIntentItems+" WHERE "+whereQry, defaultPaymentIntentItemsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&paymentIntentItems, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentIntentItems with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentIntentItemsByID] failed get paymentIntentItems")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentIntentItemsByID(ctx context.Context, primaryID model.PaymentIntentItemsPrimaryID, paymentIntentItems *model.PaymentIntentItems, paymentIntentItemsUpdateFields ...PaymentIntentItemsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentIntentItemsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentIntentItems] failed checking paymentIntentItems whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntentItems with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentIntentItems == nil {
		if len(paymentIntentItemsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentIntentItemsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentIntentItems = &model.PaymentIntentItems{}
	}
	var (
		defaultPaymentIntentItemsUpdateFields = defaultPaymentIntentItemsUpdateFields(*paymentIntentItems)
		tempUpdateField                       PaymentIntentItemsUpdateFieldList
		selectFields                          = NewPaymentIntentItemsSelectFields()
	)
	if len(paymentIntentItemsUpdateFields) > 0 {
		for _, updateField := range paymentIntentItemsUpdateFields {
			if updateField.paymentIntentItemsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentIntentItemsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentIntentItemsCompositePrimaryKeyWhere([]model.PaymentIntentItemsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentIntentItemsCommand(defaultPaymentIntentItemsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentIntentItemsQueries.updatePaymentIntentItems+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentIntentItems] error when try to update paymentIntentItems by id")
	}
	return err
}

var (
	paymentIntentItemsQueries = struct {
		selectPaymentIntentItems      string
		selectCountPaymentIntentItems string
		deletePaymentIntentItems      string
		updatePaymentIntentItems      string
		insertPaymentIntentItems      string
	}{
		selectPaymentIntentItems:      "SELECT %s FROM \"payment_intent_items\"",
		selectCountPaymentIntentItems: "SELECT COUNT(\"id\") FROM \"payment_intent_items\"",
		deletePaymentIntentItems:      "DELETE FROM \"payment_intent_items\"",
		updatePaymentIntentItems:      "UPDATE \"payment_intent_items\" SET %s ",
		insertPaymentIntentItems:      "INSERT INTO \"payment_intent_items\" %s VALUES %s",
	}
)

type PaymentIntentItemsRepository interface {
	CreatePaymentIntentItems(ctx context.Context, paymentIntentItems *model.PaymentIntentItems, fieldsInsert ...PaymentIntentItemsField) error
	BulkCreatePaymentIntentItems(ctx context.Context, paymentIntentItemsList []*model.PaymentIntentItems, fieldsInsert ...PaymentIntentItemsField) error
	ResolvePaymentIntentItems(ctx context.Context, selectFields ...PaymentIntentItemsField) (model.PaymentIntentItemsList, error)
	ResolvePaymentIntentItemsByID(ctx context.Context, primaryID model.PaymentIntentItemsPrimaryID, selectFields ...PaymentIntentItemsField) (model.PaymentIntentItems, error)
	UpdatePaymentIntentItemsByID(ctx context.Context, id model.PaymentIntentItemsPrimaryID, paymentIntentItems *model.PaymentIntentItems, paymentIntentItemsUpdateFields ...PaymentIntentItemsUpdateField) error
	BulkUpdatePaymentIntentItems(ctx context.Context, paymentIntentItemsListMap map[model.PaymentIntentItemsPrimaryID]*model.PaymentIntentItems, PaymentIntentItemssMapUpdateFieldsRequest map[model.PaymentIntentItemsPrimaryID]PaymentIntentItemsUpdateFieldList) (err error)
	DeletePaymentIntentItemsByID(ctx context.Context, id model.PaymentIntentItemsPrimaryID) error
	BulkDeletePaymentIntentItemsByIDs(ctx context.Context, ids []model.PaymentIntentItemsPrimaryID) error
	ResolvePaymentIntentItemsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentIntentItemsFilterResult, err error)
	IsExistPaymentIntentItemsByIDs(ctx context.Context, ids []model.PaymentIntentItemsPrimaryID) (exists bool, notFoundIds []model.PaymentIntentItemsPrimaryID, err error)
	IsExistPaymentIntentItemsByID(ctx context.Context, id model.PaymentIntentItemsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
