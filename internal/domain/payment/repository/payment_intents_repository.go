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

func composeInsertFieldsAndParamsPaymentIntents(paymentIntentsList []model.PaymentIntents, fieldsInsert ...PaymentIntentsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentIntentsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentIntents := range paymentIntentsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentIntents.Id)
			case selectField.IntentCode():
				args = append(args, paymentIntents.IntentCode)
			case selectField.SourceService():
				args = append(args, paymentIntents.SourceService)
			case selectField.SourceType():
				args = append(args, paymentIntents.SourceType)
			case selectField.SourceId():
				args = append(args, paymentIntents.SourceId)
			case selectField.MerchantId():
				args = append(args, paymentIntents.MerchantId)
			case selectField.CustomerId():
				args = append(args, paymentIntents.CustomerId)
			case selectField.Amount():
				args = append(args, paymentIntents.Amount)
			case selectField.Currency():
				args = append(args, paymentIntents.Currency)
			case selectField.Status():
				args = append(args, paymentIntents.Status)
			case selectField.SelectedMethodCode():
				args = append(args, paymentIntents.SelectedMethodCode)
			case selectField.SelectedChannelCode():
				args = append(args, paymentIntents.SelectedChannelCode)
			case selectField.Description():
				args = append(args, paymentIntents.Description)
			case selectField.ExpiresAt():
				args = append(args, paymentIntents.ExpiresAt)
			case selectField.PaidAt():
				args = append(args, paymentIntents.PaidAt)
			case selectField.CanceledAt():
				args = append(args, paymentIntents.CanceledAt)
			case selectField.CancellationReason():
				args = append(args, paymentIntents.CancellationReason)
			case selectField.IdempotencyKey():
				args = append(args, paymentIntents.IdempotencyKey)
			case selectField.SourceSnapshot():
				args = append(args, paymentIntents.SourceSnapshot)
			case selectField.Metadata():
				args = append(args, paymentIntents.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentIntents.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentIntents.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentIntents.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentIntents.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentIntents.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentIntents.MetaDeletedBy)

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

func composePaymentIntentsCompositePrimaryKeyWhere(primaryIDs []model.PaymentIntentsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_intents\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentIntentsSelectFields() string {
	fields := NewPaymentIntentsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentIntentsSelectFields(selectFields ...PaymentIntentsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentIntentsField string
type PaymentIntentsFieldList []PaymentIntentsField

type PaymentIntentsSelectFields struct {
}

func (ss PaymentIntentsSelectFields) Id() PaymentIntentsField {
	return PaymentIntentsField("id")
}

func (ss PaymentIntentsSelectFields) IntentCode() PaymentIntentsField {
	return PaymentIntentsField("intent_code")
}

func (ss PaymentIntentsSelectFields) SourceService() PaymentIntentsField {
	return PaymentIntentsField("source_service")
}

func (ss PaymentIntentsSelectFields) SourceType() PaymentIntentsField {
	return PaymentIntentsField("source_type")
}

func (ss PaymentIntentsSelectFields) SourceId() PaymentIntentsField {
	return PaymentIntentsField("source_id")
}

func (ss PaymentIntentsSelectFields) MerchantId() PaymentIntentsField {
	return PaymentIntentsField("merchant_id")
}

func (ss PaymentIntentsSelectFields) CustomerId() PaymentIntentsField {
	return PaymentIntentsField("customer_id")
}

func (ss PaymentIntentsSelectFields) Amount() PaymentIntentsField {
	return PaymentIntentsField("amount")
}

func (ss PaymentIntentsSelectFields) Currency() PaymentIntentsField {
	return PaymentIntentsField("currency")
}

func (ss PaymentIntentsSelectFields) Status() PaymentIntentsField {
	return PaymentIntentsField("status")
}

func (ss PaymentIntentsSelectFields) SelectedMethodCode() PaymentIntentsField {
	return PaymentIntentsField("selected_method_code")
}

func (ss PaymentIntentsSelectFields) SelectedChannelCode() PaymentIntentsField {
	return PaymentIntentsField("selected_channel_code")
}

func (ss PaymentIntentsSelectFields) Description() PaymentIntentsField {
	return PaymentIntentsField("description")
}

func (ss PaymentIntentsSelectFields) ExpiresAt() PaymentIntentsField {
	return PaymentIntentsField("expires_at")
}

func (ss PaymentIntentsSelectFields) PaidAt() PaymentIntentsField {
	return PaymentIntentsField("paid_at")
}

func (ss PaymentIntentsSelectFields) CanceledAt() PaymentIntentsField {
	return PaymentIntentsField("canceled_at")
}

func (ss PaymentIntentsSelectFields) CancellationReason() PaymentIntentsField {
	return PaymentIntentsField("cancellation_reason")
}

func (ss PaymentIntentsSelectFields) IdempotencyKey() PaymentIntentsField {
	return PaymentIntentsField("idempotency_key")
}

func (ss PaymentIntentsSelectFields) SourceSnapshot() PaymentIntentsField {
	return PaymentIntentsField("source_snapshot")
}

func (ss PaymentIntentsSelectFields) Metadata() PaymentIntentsField {
	return PaymentIntentsField("metadata")
}

func (ss PaymentIntentsSelectFields) MetaCreatedAt() PaymentIntentsField {
	return PaymentIntentsField("meta_created_at")
}

func (ss PaymentIntentsSelectFields) MetaCreatedBy() PaymentIntentsField {
	return PaymentIntentsField("meta_created_by")
}

func (ss PaymentIntentsSelectFields) MetaUpdatedAt() PaymentIntentsField {
	return PaymentIntentsField("meta_updated_at")
}

func (ss PaymentIntentsSelectFields) MetaUpdatedBy() PaymentIntentsField {
	return PaymentIntentsField("meta_updated_by")
}

func (ss PaymentIntentsSelectFields) MetaDeletedAt() PaymentIntentsField {
	return PaymentIntentsField("meta_deleted_at")
}

func (ss PaymentIntentsSelectFields) MetaDeletedBy() PaymentIntentsField {
	return PaymentIntentsField("meta_deleted_by")
}

func (ss PaymentIntentsSelectFields) All() PaymentIntentsFieldList {
	return []PaymentIntentsField{
		ss.Id(),
		ss.IntentCode(),
		ss.SourceService(),
		ss.SourceType(),
		ss.SourceId(),
		ss.MerchantId(),
		ss.CustomerId(),
		ss.Amount(),
		ss.Currency(),
		ss.Status(),
		ss.SelectedMethodCode(),
		ss.SelectedChannelCode(),
		ss.Description(),
		ss.ExpiresAt(),
		ss.PaidAt(),
		ss.CanceledAt(),
		ss.CancellationReason(),
		ss.IdempotencyKey(),
		ss.SourceSnapshot(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentIntentsSelectFields() PaymentIntentsSelectFields {
	return PaymentIntentsSelectFields{}
}

type PaymentIntentsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentIntentsUpdateField struct {
	paymentIntentsField PaymentIntentsField
	opt                 PaymentIntentsUpdateFieldOption
	value               interface{}
}
type PaymentIntentsUpdateFieldList []PaymentIntentsUpdateField

func defaultPaymentIntentsUpdateFieldOption() PaymentIntentsUpdateFieldOption {
	return PaymentIntentsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentIntentsOption(useIncrement bool) func(*PaymentIntentsUpdateFieldOption) {
	return func(pcufo *PaymentIntentsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentIntentsUpdateField(field PaymentIntentsField, val interface{}, opts ...func(*PaymentIntentsUpdateFieldOption)) PaymentIntentsUpdateField {
	defaultOpt := defaultPaymentIntentsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentIntentsUpdateField{
		paymentIntentsField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultPaymentIntentsUpdateFields(paymentIntents model.PaymentIntents) (paymentIntentsUpdateFieldList PaymentIntentsUpdateFieldList) {
	selectFields := NewPaymentIntentsSelectFields()
	paymentIntentsUpdateFieldList = append(paymentIntentsUpdateFieldList,
		NewPaymentIntentsUpdateField(selectFields.Id(), paymentIntents.Id),
		NewPaymentIntentsUpdateField(selectFields.IntentCode(), paymentIntents.IntentCode),
		NewPaymentIntentsUpdateField(selectFields.SourceService(), paymentIntents.SourceService),
		NewPaymentIntentsUpdateField(selectFields.SourceType(), paymentIntents.SourceType),
		NewPaymentIntentsUpdateField(selectFields.SourceId(), paymentIntents.SourceId),
		NewPaymentIntentsUpdateField(selectFields.MerchantId(), paymentIntents.MerchantId),
		NewPaymentIntentsUpdateField(selectFields.CustomerId(), paymentIntents.CustomerId),
		NewPaymentIntentsUpdateField(selectFields.Amount(), paymentIntents.Amount),
		NewPaymentIntentsUpdateField(selectFields.Currency(), paymentIntents.Currency),
		NewPaymentIntentsUpdateField(selectFields.Status(), paymentIntents.Status),
		NewPaymentIntentsUpdateField(selectFields.SelectedMethodCode(), paymentIntents.SelectedMethodCode),
		NewPaymentIntentsUpdateField(selectFields.SelectedChannelCode(), paymentIntents.SelectedChannelCode),
		NewPaymentIntentsUpdateField(selectFields.Description(), paymentIntents.Description),
		NewPaymentIntentsUpdateField(selectFields.ExpiresAt(), paymentIntents.ExpiresAt),
		NewPaymentIntentsUpdateField(selectFields.PaidAt(), paymentIntents.PaidAt),
		NewPaymentIntentsUpdateField(selectFields.CanceledAt(), paymentIntents.CanceledAt),
		NewPaymentIntentsUpdateField(selectFields.CancellationReason(), paymentIntents.CancellationReason),
		NewPaymentIntentsUpdateField(selectFields.IdempotencyKey(), paymentIntents.IdempotencyKey),
		NewPaymentIntentsUpdateField(selectFields.SourceSnapshot(), paymentIntents.SourceSnapshot),
		NewPaymentIntentsUpdateField(selectFields.Metadata(), paymentIntents.Metadata),
		NewPaymentIntentsUpdateField(selectFields.MetaCreatedAt(), paymentIntents.MetaCreatedAt),
		NewPaymentIntentsUpdateField(selectFields.MetaCreatedBy(), paymentIntents.MetaCreatedBy),
		NewPaymentIntentsUpdateField(selectFields.MetaUpdatedAt(), paymentIntents.MetaUpdatedAt),
		NewPaymentIntentsUpdateField(selectFields.MetaUpdatedBy(), paymentIntents.MetaUpdatedBy),
		NewPaymentIntentsUpdateField(selectFields.MetaDeletedAt(), paymentIntents.MetaDeletedAt),
		NewPaymentIntentsUpdateField(selectFields.MetaDeletedBy(), paymentIntents.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentIntentsCommand(paymentIntentsUpdateFieldList PaymentIntentsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentIntentsUpdateFieldList {
		field := string(updateField.paymentIntentsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentIntents(ctx context.Context, paymentIntentsList []*model.PaymentIntents, fieldsInsert ...PaymentIntentsField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.PaymentIntentsPrimaryID
		paymentIntentsValueList []model.PaymentIntents
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentIntentsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentIntents := range paymentIntentsList {

		primaryIds = append(primaryIds, paymentIntents.ToPaymentIntentsPrimaryID())

		paymentIntentsValueList = append(paymentIntentsValueList, *paymentIntents)
	}

	_, notFoundIds, err := repo.IsExistPaymentIntentsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentIntents] failed checking paymentIntents whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentIntentsPrimaryID{}
		mapNotFoundIds := map[model.PaymentIntentsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentIntents", fmt.Sprintf("paymentIntents with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentIntents(paymentIntentsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentIntentsQueries.insertPaymentIntents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentIntents] failed exec create paymentIntents query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentIntentsByIDs(ctx context.Context, primaryIDs []model.PaymentIntentsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentIntentsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentIntentsByIDs] failed checking paymentIntents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntents with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_intents\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentIntentsQueries.deletePaymentIntents + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentIntentsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentIntentsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentIntentsByIDs(ctx context.Context, ids []model.PaymentIntentsPrimaryID) (exists bool, notFoundIds []model.PaymentIntentsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_intents\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentIntentsQueries.selectPaymentIntents, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentIntentsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentIntentsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentIntentsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentIntentsPrimaryID]bool{}
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

// BulkUpdatePaymentIntents is used to bulk update paymentIntents, by default it will update all field
// if want to update specific field, then fill paymentIntentssMapUpdateFieldsRequest else please fill paymentIntentssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentIntents(ctx context.Context, paymentIntentssMap map[model.PaymentIntentsPrimaryID]*model.PaymentIntents, paymentIntentssMapUpdateFieldsRequest map[model.PaymentIntentsPrimaryID]PaymentIntentsUpdateFieldList) (err error) {
	if len(paymentIntentssMap) == 0 && len(paymentIntentssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentIntentssMapUpdateField map[model.PaymentIntentsPrimaryID]PaymentIntentsUpdateFieldList = map[model.PaymentIntentsPrimaryID]PaymentIntentsUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(paymentIntentssMap) > 0 {
		for id, paymentIntents := range paymentIntentssMap {
			if paymentIntents == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentIntents] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentIntentssMapUpdateField[id] = defaultPaymentIntentsUpdateFields(*paymentIntents)
		}
	} else {
		paymentIntentssMapUpdateField = paymentIntentssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentIntentsQuery(paymentIntentssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentIntentsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentIntents] failed checking paymentIntents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntents with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentIntentsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_intents\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentIntents] failed exec query")
	}
	return
}

type PaymentIntentsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentIntentsFieldParameter(param string, args ...interface{}) PaymentIntentsFieldParameter {
	return PaymentIntentsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentIntentsQuery(mapPaymentIntentss map[model.PaymentIntentsPrimaryID]PaymentIntentsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentIntentsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentIntentsPrimaryID]map[string]interface{}{}
	paymentIntentsSelectFields := NewPaymentIntentsSelectFields()
	for id, updateFields := range mapPaymentIntentss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentIntentsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentIntentss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentIntentsFieldType(updateField.paymentIntentsField)))
			args = append(args, fields[string(updateField.paymentIntentsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentIntentsField))
		if updateField.paymentIntentsField == paymentIntentsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentIntentsField, asTableValues, updateField.paymentIntentsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentIntentsField,
				"\"payment_intents\"", updateField.paymentIntentsField,
				asTableValues, updateField.paymentIntentsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentIntentsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentIntentsPrimaryID, asTableValue string) (whereQry string) {
	paymentIntentsSelectFields := NewPaymentIntentsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_intents\".\"id\" = %s.\"id\"::"+GetPaymentIntentsFieldType(paymentIntentsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentIntentsFieldType(paymentIntentsField PaymentIntentsField) string {
	selectPaymentIntentsFields := NewPaymentIntentsSelectFields()
	switch paymentIntentsField {

	case selectPaymentIntentsFields.Id():
		return "uuid"

	case selectPaymentIntentsFields.IntentCode():
		return "text"

	case selectPaymentIntentsFields.SourceService():
		return "text"

	case selectPaymentIntentsFields.SourceType():
		return "text"

	case selectPaymentIntentsFields.SourceId():
		return "uuid"

	case selectPaymentIntentsFields.MerchantId():
		return "uuid"

	case selectPaymentIntentsFields.CustomerId():
		return "uuid"

	case selectPaymentIntentsFields.Amount():
		return "numeric"

	case selectPaymentIntentsFields.Currency():
		return "text"

	case selectPaymentIntentsFields.Status():
		return "payment_intent_status_enum"

	case selectPaymentIntentsFields.SelectedMethodCode():
		return "text"

	case selectPaymentIntentsFields.SelectedChannelCode():
		return "text"

	case selectPaymentIntentsFields.Description():
		return "text"

	case selectPaymentIntentsFields.ExpiresAt():
		return "timestamptz"

	case selectPaymentIntentsFields.PaidAt():
		return "timestamptz"

	case selectPaymentIntentsFields.CanceledAt():
		return "timestamptz"

	case selectPaymentIntentsFields.CancellationReason():
		return "text"

	case selectPaymentIntentsFields.IdempotencyKey():
		return "text"

	case selectPaymentIntentsFields.SourceSnapshot():
		return "jsonb"

	case selectPaymentIntentsFields.Metadata():
		return "jsonb"

	case selectPaymentIntentsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentIntentsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentIntentsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentIntentsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentIntentsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentIntentsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentIntents(ctx context.Context, paymentIntents *model.PaymentIntents, fieldsInsert ...PaymentIntentsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentIntentsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentIntentsPrimaryID{
		Id: paymentIntents.Id,
	}
	exists, err := repo.IsExistPaymentIntentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentIntents] failed checking paymentIntents whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentIntents", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentIntents([]model.PaymentIntents{*paymentIntents}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentIntentsQueries.insertPaymentIntents, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentIntents] failed exec create paymentIntents query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentIntentsByID(ctx context.Context, primaryID model.PaymentIntentsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentIntentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentIntentsByID] failed checking paymentIntents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentIntentsCompositePrimaryKeyWhere([]model.PaymentIntentsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentIntentsQueries.deletePaymentIntents + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentIntentsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentIntentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentIntentsFilterResult, err error) {
	query, args, err := composePaymentIntentsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentIntentsByFilter] failed compose paymentIntents filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentIntentsByFilter] failed get paymentIntents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentIntentsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentIntentsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentIntentsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentIntentsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentIntentsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 26 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentIntentsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 26+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["intent_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"intent_code\"")
			selectedColumns["intent_code"] = struct{}{}
		}
		if _, selected := selectedColumns["source_service"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_service\"")
			selectedColumns["source_service"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_id\"")
			selectedColumns["merchant_id"] = struct{}{}
		}
		if _, selected := selectedColumns["customer_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"customer_id\"")
			selectedColumns["customer_id"] = struct{}{}
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
		if _, selected := selectedColumns["selected_method_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"selected_method_code\"")
			selectedColumns["selected_method_code"] = struct{}{}
		}
		if _, selected := selectedColumns["selected_channel_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"selected_channel_code\"")
			selectedColumns["selected_channel_code"] = struct{}{}
		}
		if _, selected := selectedColumns["description"]; !selected {
			selectColumns = append(selectColumns, "base.\"description\"")
			selectedColumns["description"] = struct{}{}
		}
		if _, selected := selectedColumns["expires_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"expires_at\"")
			selectedColumns["expires_at"] = struct{}{}
		}
		if _, selected := selectedColumns["paid_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"paid_at\"")
			selectedColumns["paid_at"] = struct{}{}
		}
		if _, selected := selectedColumns["canceled_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"canceled_at\"")
			selectedColumns["canceled_at"] = struct{}{}
		}
		if _, selected := selectedColumns["cancellation_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"cancellation_reason\"")
			selectedColumns["cancellation_reason"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["source_snapshot"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_snapshot\"")
			selectedColumns["source_snapshot"] = struct{}{}
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

type paymentIntentsFilterPlaceholder struct {
	index int
}

func (p *paymentIntentsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentIntentsFilterPredicate(filterField model.FilterField, placeholders *paymentIntentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentIntentsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentIntentsFilterSQLExpr(spec)
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

func composePaymentIntentsFilterGroup(group model.FilterGroup, placeholders *paymentIntentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentIntentsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentIntentsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentIntentsFilterWhereQueries(filter model.Filter, placeholders *paymentIntentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentIntentsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentIntentsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentIntentsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentIntentsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentIntentsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentIntentsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentIntentsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentIntentsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentIntentsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentIntentsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentIntentsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_intents\" base%s", strings.Join(selectColumns, ","), composePaymentIntentsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentIntentsByID(ctx context.Context, primaryID model.PaymentIntentsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentIntentsCompositePrimaryKeyWhere([]model.PaymentIntentsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentIntentsQueries.selectCountPaymentIntents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentIntentsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentIntents(ctx context.Context, selectFields ...PaymentIntentsField) (paymentIntentsList model.PaymentIntentsList, err error) {
	var (
		defaultPaymentIntentsSelectFields = defaultPaymentIntentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentIntentsSelectFields = composePaymentIntentsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentIntentsQueries.selectPaymentIntents, defaultPaymentIntentsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentIntentsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentIntents] failed get paymentIntents list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentIntentsByID(ctx context.Context, primaryID model.PaymentIntentsPrimaryID, selectFields ...PaymentIntentsField) (paymentIntents model.PaymentIntents, err error) {
	var (
		defaultPaymentIntentsSelectFields = defaultPaymentIntentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentIntentsSelectFields = composePaymentIntentsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentIntentsCompositePrimaryKeyWhere([]model.PaymentIntentsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentIntentsQueries.selectPaymentIntents+" WHERE "+whereQry, defaultPaymentIntentsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentIntents, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentIntents with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentIntentsByID] failed get paymentIntents")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentIntentsByID(ctx context.Context, primaryID model.PaymentIntentsPrimaryID, paymentIntents *model.PaymentIntents, paymentIntentsUpdateFields ...PaymentIntentsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentIntentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentIntents] failed checking paymentIntents whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentIntents with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentIntents == nil {
		if len(paymentIntentsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentIntentsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentIntents = &model.PaymentIntents{}
	}
	var (
		defaultPaymentIntentsUpdateFields = defaultPaymentIntentsUpdateFields(*paymentIntents)
		tempUpdateField                   PaymentIntentsUpdateFieldList
		selectFields                      = NewPaymentIntentsSelectFields()
	)
	if len(paymentIntentsUpdateFields) > 0 {
		for _, updateField := range paymentIntentsUpdateFields {
			if updateField.paymentIntentsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentIntentsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentIntentsCompositePrimaryKeyWhere([]model.PaymentIntentsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentIntentsCommand(defaultPaymentIntentsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentIntentsQueries.updatePaymentIntents+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentIntents] error when try to update paymentIntents by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentIntentsByFilter(ctx context.Context, filter model.Filter, paymentIntentsUpdateFields ...PaymentIntentsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentIntentsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentIntentsUpdateFieldList
		selectFields = NewPaymentIntentsSelectFields()
	)
	for _, updateField := range paymentIntentsUpdateFields {
		if updateField.paymentIntentsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentIntentsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentIntentsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentIntentsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_intents\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentIntentsByFilter] error when try to update paymentIntents by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentIntentsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentIntentsQueries = struct {
		selectPaymentIntents      string
		selectCountPaymentIntents string
		deletePaymentIntents      string
		updatePaymentIntents      string
		insertPaymentIntents      string
	}{
		selectPaymentIntents:      "SELECT %s FROM \"payment_intents\"",
		selectCountPaymentIntents: "SELECT COUNT(\"id\") FROM \"payment_intents\"",
		deletePaymentIntents:      "DELETE FROM \"payment_intents\"",
		updatePaymentIntents:      "UPDATE \"payment_intents\" SET %s ",
		insertPaymentIntents:      "INSERT INTO \"payment_intents\" %s VALUES %s",
	}
)

type PaymentIntentsRepository interface {
	CreatePaymentIntents(ctx context.Context, paymentIntents *model.PaymentIntents, fieldsInsert ...PaymentIntentsField) error
	BulkCreatePaymentIntents(ctx context.Context, paymentIntentsList []*model.PaymentIntents, fieldsInsert ...PaymentIntentsField) error
	ResolvePaymentIntents(ctx context.Context, selectFields ...PaymentIntentsField) (model.PaymentIntentsList, error)
	ResolvePaymentIntentsByID(ctx context.Context, primaryID model.PaymentIntentsPrimaryID, selectFields ...PaymentIntentsField) (model.PaymentIntents, error)
	UpdatePaymentIntentsByID(ctx context.Context, id model.PaymentIntentsPrimaryID, paymentIntents *model.PaymentIntents, paymentIntentsUpdateFields ...PaymentIntentsUpdateField) error
	UpdatePaymentIntentsByFilter(ctx context.Context, filter model.Filter, paymentIntentsUpdateFields ...PaymentIntentsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentIntents(ctx context.Context, paymentIntentsListMap map[model.PaymentIntentsPrimaryID]*model.PaymentIntents, PaymentIntentssMapUpdateFieldsRequest map[model.PaymentIntentsPrimaryID]PaymentIntentsUpdateFieldList) (err error)
	DeletePaymentIntentsByID(ctx context.Context, id model.PaymentIntentsPrimaryID) error
	BulkDeletePaymentIntentsByIDs(ctx context.Context, ids []model.PaymentIntentsPrimaryID) error
	ResolvePaymentIntentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentIntentsFilterResult, err error)
	IsExistPaymentIntentsByIDs(ctx context.Context, ids []model.PaymentIntentsPrimaryID) (exists bool, notFoundIds []model.PaymentIntentsPrimaryID, err error)
	IsExistPaymentIntentsByID(ctx context.Context, id model.PaymentIntentsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
