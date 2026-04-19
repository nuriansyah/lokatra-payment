package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/netip"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/inetaddr"
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
			case selectField.MerchantId():
				args = append(args, paymentIntents.MerchantId)
			case selectField.OrderId():
				args = append(args, paymentIntents.OrderId)
			case selectField.OrderType():
				args = append(args, paymentIntents.OrderType)
			case selectField.Amount():
				args = append(args, paymentIntents.Amount)
			case selectField.Currency():
				args = append(args, paymentIntents.Currency)
			case selectField.TaxAmount():
				args = append(args, paymentIntents.TaxAmount)
			case selectField.DiscountAmount():
				args = append(args, paymentIntents.DiscountAmount)
			case selectField.TipAmount():
				args = append(args, paymentIntents.TipAmount)
			case selectField.UserId():
				args = append(args, paymentIntents.UserId)
			case selectField.CustomerName():
				args = append(args, paymentIntents.CustomerName)
			case selectField.CustomerEmail():
				args = append(args, paymentIntents.CustomerEmail)
			case selectField.CustomerPhone():
				args = append(args, paymentIntents.CustomerPhone)
			case selectField.CustomerIp():
				args = append(args, customerIPValue(paymentIntents.CustomerIp))
			case selectField.CustomerCountry():
				args = append(args, paymentIntents.CustomerCountry)
			case selectField.PaymentMethodId():
				args = append(args, paymentIntents.PaymentMethodId)
			case selectField.PaymentMethodType():
				args = append(args, paymentIntents.PaymentMethodType)
			case selectField.Status():
				args = append(args, paymentIntents.Status)
			case selectField.RoutingProfileId():
				args = append(args, paymentIntents.RoutingProfileId)
			case selectField.ExpiresAt():
				args = append(args, paymentIntents.ExpiresAt)
			case selectField.Requires3ds():
				args = append(args, paymentIntents.Requires3ds)
			case selectField.ThreeDsVersion():
				args = append(args, paymentIntents.ThreeDsVersion)
			case selectField.Description():
				args = append(args, paymentIntents.Description)
			case selectField.StatementDescriptor():
				args = append(args, paymentIntents.StatementDescriptor)
			case selectField.Metadata():
				args = append(args, paymentIntents.Metadata)
			case selectField.PromoCode():
				args = append(args, paymentIntents.PromoCode)
			case selectField.PromoDiscountAmount():
				args = append(args, paymentIntents.PromoDiscountAmount)
			case selectField.IdempotencyKeyId():
				args = append(args, paymentIntents.IdempotencyKeyId)
			case selectField.ConfirmedAt():
				args = append(args, paymentIntents.ConfirmedAt)
			case selectField.CancelledAt():
				args = append(args, paymentIntents.CancelledAt)
			case selectField.CancellationReason():
				args = append(args, paymentIntents.CancellationReason)
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

func customerIPValue(addr inetaddr.NullIP) interface{} {
	if !addr.Valid {
		return nil
	}
	return addr.String()
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

func (ss PaymentIntentsSelectFields) MerchantId() PaymentIntentsField {
	return PaymentIntentsField("merchant_id")
}

func (ss PaymentIntentsSelectFields) OrderId() PaymentIntentsField {
	return PaymentIntentsField("order_id")
}

func (ss PaymentIntentsSelectFields) OrderType() PaymentIntentsField {
	return PaymentIntentsField("order_type")
}

func (ss PaymentIntentsSelectFields) Amount() PaymentIntentsField {
	return PaymentIntentsField("amount")
}

func (ss PaymentIntentsSelectFields) Currency() PaymentIntentsField {
	return PaymentIntentsField("currency")
}

func (ss PaymentIntentsSelectFields) TaxAmount() PaymentIntentsField {
	return PaymentIntentsField("tax_amount")
}

func (ss PaymentIntentsSelectFields) DiscountAmount() PaymentIntentsField {
	return PaymentIntentsField("discount_amount")
}

func (ss PaymentIntentsSelectFields) TipAmount() PaymentIntentsField {
	return PaymentIntentsField("tip_amount")
}

func (ss PaymentIntentsSelectFields) UserId() PaymentIntentsField {
	return PaymentIntentsField("user_id")
}

func (ss PaymentIntentsSelectFields) CustomerName() PaymentIntentsField {
	return PaymentIntentsField("customer_name")
}

func (ss PaymentIntentsSelectFields) CustomerEmail() PaymentIntentsField {
	return PaymentIntentsField("customer_email")
}

func (ss PaymentIntentsSelectFields) CustomerPhone() PaymentIntentsField {
	return PaymentIntentsField("customer_phone")
}

func (ss PaymentIntentsSelectFields) CustomerIp() PaymentIntentsField {
	return PaymentIntentsField("customer_ip")
}

func (ss PaymentIntentsSelectFields) CustomerCountry() PaymentIntentsField {
	return PaymentIntentsField("customer_country")
}

func (ss PaymentIntentsSelectFields) PaymentMethodId() PaymentIntentsField {
	return PaymentIntentsField("payment_method_id")
}

func (ss PaymentIntentsSelectFields) PaymentMethodType() PaymentIntentsField {
	return PaymentIntentsField("payment_method_type")
}

func (ss PaymentIntentsSelectFields) Status() PaymentIntentsField {
	return PaymentIntentsField("status")
}

func (ss PaymentIntentsSelectFields) RoutingProfileId() PaymentIntentsField {
	return PaymentIntentsField("routing_profile_id")
}

func (ss PaymentIntentsSelectFields) ExpiresAt() PaymentIntentsField {
	return PaymentIntentsField("expires_at")
}

func (ss PaymentIntentsSelectFields) Requires3ds() PaymentIntentsField {
	return PaymentIntentsField("requires_3ds")
}

func (ss PaymentIntentsSelectFields) ThreeDsVersion() PaymentIntentsField {
	return PaymentIntentsField("three_ds_version")
}

func (ss PaymentIntentsSelectFields) Description() PaymentIntentsField {
	return PaymentIntentsField("description")
}

func (ss PaymentIntentsSelectFields) StatementDescriptor() PaymentIntentsField {
	return PaymentIntentsField("statement_descriptor")
}

func (ss PaymentIntentsSelectFields) Metadata() PaymentIntentsField {
	return PaymentIntentsField("metadata")
}

func (ss PaymentIntentsSelectFields) PromoCode() PaymentIntentsField {
	return PaymentIntentsField("promo_code")
}

func (ss PaymentIntentsSelectFields) PromoDiscountAmount() PaymentIntentsField {
	return PaymentIntentsField("promo_discount_amount")
}

func (ss PaymentIntentsSelectFields) IdempotencyKeyId() PaymentIntentsField {
	return PaymentIntentsField("idempotency_key_id")
}

func (ss PaymentIntentsSelectFields) ConfirmedAt() PaymentIntentsField {
	return PaymentIntentsField("confirmed_at")
}

func (ss PaymentIntentsSelectFields) CancelledAt() PaymentIntentsField {
	return PaymentIntentsField("cancelled_at")
}

func (ss PaymentIntentsSelectFields) CancellationReason() PaymentIntentsField {
	return PaymentIntentsField("cancellation_reason")
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
		ss.MerchantId(),
		ss.OrderId(),
		ss.OrderType(),
		ss.Amount(),
		ss.Currency(),
		ss.TaxAmount(),
		ss.DiscountAmount(),
		ss.TipAmount(),
		ss.UserId(),
		ss.CustomerName(),
		ss.CustomerEmail(),
		ss.CustomerPhone(),
		ss.CustomerIp(),
		ss.CustomerCountry(),
		ss.PaymentMethodId(),
		ss.PaymentMethodType(),
		ss.Status(),
		ss.RoutingProfileId(),
		ss.ExpiresAt(),
		ss.Requires3ds(),
		ss.ThreeDsVersion(),
		ss.Description(),
		ss.StatementDescriptor(),
		ss.Metadata(),
		ss.PromoCode(),
		ss.PromoDiscountAmount(),
		ss.IdempotencyKeyId(),
		ss.ConfirmedAt(),
		ss.CancelledAt(),
		ss.CancellationReason(),
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
		NewPaymentIntentsUpdateField(selectFields.MerchantId(), paymentIntents.MerchantId),
		NewPaymentIntentsUpdateField(selectFields.OrderId(), paymentIntents.OrderId),
		NewPaymentIntentsUpdateField(selectFields.OrderType(), paymentIntents.OrderType),
		NewPaymentIntentsUpdateField(selectFields.Amount(), paymentIntents.Amount),
		NewPaymentIntentsUpdateField(selectFields.Currency(), paymentIntents.Currency),
		NewPaymentIntentsUpdateField(selectFields.TaxAmount(), paymentIntents.TaxAmount),
		NewPaymentIntentsUpdateField(selectFields.DiscountAmount(), paymentIntents.DiscountAmount),
		NewPaymentIntentsUpdateField(selectFields.TipAmount(), paymentIntents.TipAmount),
		NewPaymentIntentsUpdateField(selectFields.UserId(), paymentIntents.UserId),
		NewPaymentIntentsUpdateField(selectFields.CustomerName(), paymentIntents.CustomerName),
		NewPaymentIntentsUpdateField(selectFields.CustomerEmail(), paymentIntents.CustomerEmail),
		NewPaymentIntentsUpdateField(selectFields.CustomerPhone(), paymentIntents.CustomerPhone),
		NewPaymentIntentsUpdateField(selectFields.CustomerIp(), customerIPValue(paymentIntents.CustomerIp)),
		NewPaymentIntentsUpdateField(selectFields.CustomerCountry(), paymentIntents.CustomerCountry),
		NewPaymentIntentsUpdateField(selectFields.PaymentMethodId(), paymentIntents.PaymentMethodId),
		NewPaymentIntentsUpdateField(selectFields.PaymentMethodType(), paymentIntents.PaymentMethodType),
		NewPaymentIntentsUpdateField(selectFields.Status(), paymentIntents.Status),
		NewPaymentIntentsUpdateField(selectFields.RoutingProfileId(), paymentIntents.RoutingProfileId),
		NewPaymentIntentsUpdateField(selectFields.ExpiresAt(), paymentIntents.ExpiresAt),
		NewPaymentIntentsUpdateField(selectFields.Requires3ds(), paymentIntents.Requires3ds),
		NewPaymentIntentsUpdateField(selectFields.ThreeDsVersion(), paymentIntents.ThreeDsVersion),
		NewPaymentIntentsUpdateField(selectFields.Description(), paymentIntents.Description),
		NewPaymentIntentsUpdateField(selectFields.StatementDescriptor(), paymentIntents.StatementDescriptor),
		NewPaymentIntentsUpdateField(selectFields.Metadata(), paymentIntents.Metadata),
		NewPaymentIntentsUpdateField(selectFields.PromoCode(), paymentIntents.PromoCode),
		NewPaymentIntentsUpdateField(selectFields.PromoDiscountAmount(), paymentIntents.PromoDiscountAmount),
		NewPaymentIntentsUpdateField(selectFields.IdempotencyKeyId(), paymentIntents.IdempotencyKeyId),
		NewPaymentIntentsUpdateField(selectFields.ConfirmedAt(), paymentIntents.ConfirmedAt),
		NewPaymentIntentsUpdateField(selectFields.CancelledAt(), paymentIntents.CancelledAt),
		NewPaymentIntentsUpdateField(selectFields.CancellationReason(), paymentIntents.CancellationReason),
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
	err = repo.db.Read.Select(&resIds, query, params...)
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

	case selectPaymentIntentsFields.MerchantId():
		return "uuid"

	case selectPaymentIntentsFields.OrderId():
		return "uuid"

	case selectPaymentIntentsFields.OrderType():
		return "text"

	case selectPaymentIntentsFields.Amount():
		return "numeric"

	case selectPaymentIntentsFields.Currency():
		return "payment_currency"

	case selectPaymentIntentsFields.TaxAmount():
		return "numeric"

	case selectPaymentIntentsFields.DiscountAmount():
		return "numeric"

	case selectPaymentIntentsFields.TipAmount():
		return "numeric"

	case selectPaymentIntentsFields.UserId():
		return "uuid"

	case selectPaymentIntentsFields.CustomerName():
		return "text"

	case selectPaymentIntentsFields.CustomerEmail():
		return "text"

	case selectPaymentIntentsFields.CustomerPhone():
		return "text"

	case selectPaymentIntentsFields.CustomerIp():
		return "inet"

	case selectPaymentIntentsFields.CustomerCountry():
		return "text"

	case selectPaymentIntentsFields.PaymentMethodId():
		return "uuid"

	case selectPaymentIntentsFields.PaymentMethodType():
		return "payment_method_type_enum"

	case selectPaymentIntentsFields.Status():
		return "payment_status_enum"

	case selectPaymentIntentsFields.RoutingProfileId():
		return "uuid"

	case selectPaymentIntentsFields.ExpiresAt():
		return "timestamptz"

	case selectPaymentIntentsFields.Requires3ds():
		return "bool"

	case selectPaymentIntentsFields.ThreeDsVersion():
		return "text"

	case selectPaymentIntentsFields.Description():
		return "text"

	case selectPaymentIntentsFields.StatementDescriptor():
		return "text"

	case selectPaymentIntentsFields.Metadata():
		return "jsonb"

	case selectPaymentIntentsFields.PromoCode():
		return "text"

	case selectPaymentIntentsFields.PromoDiscountAmount():
		return "numeric"

	case selectPaymentIntentsFields.IdempotencyKeyId():
		return "uuid"

	case selectPaymentIntentsFields.ConfirmedAt():
		return "timestamptz"

	case selectPaymentIntentsFields.CancelledAt():
		return "timestamptz"

	case selectPaymentIntentsFields.CancellationReason():
		return "text"

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
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentIntentsByFilter] failed get paymentIntents by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentIntentsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentIntentsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultPaymentIntentsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := PaymentIntentsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, PaymentIntentsField(filterSelectField))
		}
		selectFields = composePaymentIntentsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(paymentIntentsQueries.selectPaymentIntents, selectFields)

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

func (repo *RepositoryImpl) IsExistPaymentIntentsByID(ctx context.Context, primaryID model.PaymentIntentsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentIntentsCompositePrimaryKeyWhere([]model.PaymentIntentsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentIntentsQueries.selectCountPaymentIntents, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
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

	err = repo.db.Read.Select(&paymentIntentsList, query)
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
	err = repo.db.Read.Get(&paymentIntents, query, params...)
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
	BulkUpdatePaymentIntents(ctx context.Context, paymentIntentsListMap map[model.PaymentIntentsPrimaryID]*model.PaymentIntents, PaymentIntentssMapUpdateFieldsRequest map[model.PaymentIntentsPrimaryID]PaymentIntentsUpdateFieldList) (err error)
	DeletePaymentIntentsByID(ctx context.Context, id model.PaymentIntentsPrimaryID) error
	BulkDeletePaymentIntentsByIDs(ctx context.Context, ids []model.PaymentIntentsPrimaryID) error
	ResolvePaymentIntentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentIntentsFilterResult, err error)
	IsExistPaymentIntentsByIDs(ctx context.Context, ids []model.PaymentIntentsPrimaryID) (exists bool, notFoundIds []model.PaymentIntentsPrimaryID, err error)
	IsExistPaymentIntentsByID(ctx context.Context, id model.PaymentIntentsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
