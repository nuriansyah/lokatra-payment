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

func composeInsertFieldsAndParamsPayments(paymentsList []model.Payments, fieldsInsert ...PaymentsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, payments := range paymentsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, payments.Id)
			case selectField.PaymentCode():
				args = append(args, payments.PaymentCode)
			case selectField.IntentId():
				args = append(args, payments.IntentId)
			case selectField.AttemptNumber():
				args = append(args, payments.AttemptNumber)
			case selectField.Psp():
				args = append(args, payments.Psp)
			case selectField.PspTransactionId():
				args = append(args, payments.PspTransactionId)
			case selectField.PspReference():
				args = append(args, payments.PspReference)
			case selectField.PspRawRequest():
				args = append(args, payments.PspRawRequest)
			case selectField.PspRawResponse():
				args = append(args, payments.PspRawResponse)
			case selectField.Amount():
				args = append(args, payments.Amount)
			case selectField.Currency():
				args = append(args, payments.Currency)
			case selectField.AmountInSettlementCurrency():
				args = append(args, payments.AmountInSettlementCurrency)
			case selectField.SettlementCurrency():
				args = append(args, payments.SettlementCurrency)
			case selectField.FxRate():
				args = append(args, payments.FxRate)
			case selectField.FxRateSnapshotId():
				args = append(args, payments.FxRateSnapshotId)
			case selectField.PaymentMethodId():
				args = append(args, payments.PaymentMethodId)
			case selectField.PaymentMethodType():
				args = append(args, payments.PaymentMethodType)
			case selectField.Status():
				args = append(args, payments.Status)
			case selectField.FailureCode():
				args = append(args, payments.FailureCode)
			case selectField.FailureMessage():
				args = append(args, payments.FailureMessage)
			case selectField.FailureCategory():
				args = append(args, payments.FailureCategory)
			case selectField.AuthorisedAt():
				args = append(args, payments.AuthorisedAt)
			case selectField.AuthorisedAmount():
				args = append(args, payments.AuthorisedAmount)
			case selectField.CapturedAt():
				args = append(args, payments.CapturedAt)
			case selectField.CapturedAmount():
				args = append(args, payments.CapturedAmount)
			case selectField.ProcessingFee():
				args = append(args, payments.ProcessingFee)
			case selectField.ProcessingFeeCurrency():
				args = append(args, payments.ProcessingFeeCurrency)
			case selectField.RiskScoreId():
				args = append(args, payments.RiskScoreId)
			case selectField.Description():
				args = append(args, payments.Description)
			case selectField.Metadata():
				args = append(args, payments.Metadata)
			case selectField.CompletedAt():
				args = append(args, payments.CompletedAt)
			case selectField.CancelledAt():
				args = append(args, payments.CancelledAt)
			case selectField.ExpiredAt():
				args = append(args, payments.ExpiredAt)
			case selectField.MetaCreatedAt():
				args = append(args, payments.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, payments.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, payments.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, payments.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, payments.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, payments.MetaDeletedBy)

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

func composePaymentsCompositePrimaryKeyWhere(primaryIDs []model.PaymentsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payments\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentsSelectFields() string {
	fields := NewPaymentsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentsSelectFields(selectFields ...PaymentsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentsField string
type PaymentsFieldList []PaymentsField

type PaymentsSelectFields struct {
}

func (ss PaymentsSelectFields) Id() PaymentsField {
	return PaymentsField("id")
}

func (ss PaymentsSelectFields) PaymentCode() PaymentsField {
	return PaymentsField("payment_code")
}

func (ss PaymentsSelectFields) IntentId() PaymentsField {
	return PaymentsField("intent_id")
}

func (ss PaymentsSelectFields) AttemptNumber() PaymentsField {
	return PaymentsField("attempt_number")
}

func (ss PaymentsSelectFields) Psp() PaymentsField {
	return PaymentsField("psp")
}

func (ss PaymentsSelectFields) PspTransactionId() PaymentsField {
	return PaymentsField("psp_transaction_id")
}

func (ss PaymentsSelectFields) PspReference() PaymentsField {
	return PaymentsField("psp_reference")
}

func (ss PaymentsSelectFields) PspRawRequest() PaymentsField {
	return PaymentsField("psp_raw_request")
}

func (ss PaymentsSelectFields) PspRawResponse() PaymentsField {
	return PaymentsField("psp_raw_response")
}

func (ss PaymentsSelectFields) Amount() PaymentsField {
	return PaymentsField("amount")
}

func (ss PaymentsSelectFields) Currency() PaymentsField {
	return PaymentsField("currency")
}

func (ss PaymentsSelectFields) AmountInSettlementCurrency() PaymentsField {
	return PaymentsField("amount_in_settlement_currency")
}

func (ss PaymentsSelectFields) SettlementCurrency() PaymentsField {
	return PaymentsField("settlement_currency")
}

func (ss PaymentsSelectFields) FxRate() PaymentsField {
	return PaymentsField("fx_rate")
}

func (ss PaymentsSelectFields) FxRateSnapshotId() PaymentsField {
	return PaymentsField("fx_rate_snapshot_id")
}

func (ss PaymentsSelectFields) PaymentMethodId() PaymentsField {
	return PaymentsField("payment_method_id")
}

func (ss PaymentsSelectFields) PaymentMethodType() PaymentsField {
	return PaymentsField("payment_method_type")
}

func (ss PaymentsSelectFields) Status() PaymentsField {
	return PaymentsField("status")
}

func (ss PaymentsSelectFields) FailureCode() PaymentsField {
	return PaymentsField("failure_code")
}

func (ss PaymentsSelectFields) FailureMessage() PaymentsField {
	return PaymentsField("failure_message")
}

func (ss PaymentsSelectFields) FailureCategory() PaymentsField {
	return PaymentsField("failure_category")
}

func (ss PaymentsSelectFields) AuthorisedAt() PaymentsField {
	return PaymentsField("authorised_at")
}

func (ss PaymentsSelectFields) AuthorisedAmount() PaymentsField {
	return PaymentsField("authorised_amount")
}

func (ss PaymentsSelectFields) CapturedAt() PaymentsField {
	return PaymentsField("captured_at")
}

func (ss PaymentsSelectFields) CapturedAmount() PaymentsField {
	return PaymentsField("captured_amount")
}

func (ss PaymentsSelectFields) ProcessingFee() PaymentsField {
	return PaymentsField("processing_fee")
}

func (ss PaymentsSelectFields) ProcessingFeeCurrency() PaymentsField {
	return PaymentsField("processing_fee_currency")
}

func (ss PaymentsSelectFields) RiskScoreId() PaymentsField {
	return PaymentsField("risk_score_id")
}

func (ss PaymentsSelectFields) Description() PaymentsField {
	return PaymentsField("description")
}

func (ss PaymentsSelectFields) Metadata() PaymentsField {
	return PaymentsField("metadata")
}

func (ss PaymentsSelectFields) CompletedAt() PaymentsField {
	return PaymentsField("completed_at")
}

func (ss PaymentsSelectFields) CancelledAt() PaymentsField {
	return PaymentsField("cancelled_at")
}

func (ss PaymentsSelectFields) ExpiredAt() PaymentsField {
	return PaymentsField("expired_at")
}

func (ss PaymentsSelectFields) MetaCreatedAt() PaymentsField {
	return PaymentsField("meta_created_at")
}

func (ss PaymentsSelectFields) MetaCreatedBy() PaymentsField {
	return PaymentsField("meta_created_by")
}

func (ss PaymentsSelectFields) MetaUpdatedAt() PaymentsField {
	return PaymentsField("meta_updated_at")
}

func (ss PaymentsSelectFields) MetaUpdatedBy() PaymentsField {
	return PaymentsField("meta_updated_by")
}

func (ss PaymentsSelectFields) MetaDeletedAt() PaymentsField {
	return PaymentsField("meta_deleted_at")
}

func (ss PaymentsSelectFields) MetaDeletedBy() PaymentsField {
	return PaymentsField("meta_deleted_by")
}

func (ss PaymentsSelectFields) All() PaymentsFieldList {
	return []PaymentsField{
		ss.Id(),
		ss.PaymentCode(),
		ss.IntentId(),
		ss.AttemptNumber(),
		ss.Psp(),
		ss.PspTransactionId(),
		ss.PspReference(),
		ss.PspRawRequest(),
		ss.PspRawResponse(),
		ss.Amount(),
		ss.Currency(),
		ss.AmountInSettlementCurrency(),
		ss.SettlementCurrency(),
		ss.FxRate(),
		ss.FxRateSnapshotId(),
		ss.PaymentMethodId(),
		ss.PaymentMethodType(),
		ss.Status(),
		ss.FailureCode(),
		ss.FailureMessage(),
		ss.FailureCategory(),
		ss.AuthorisedAt(),
		ss.AuthorisedAmount(),
		ss.CapturedAt(),
		ss.CapturedAmount(),
		ss.ProcessingFee(),
		ss.ProcessingFeeCurrency(),
		ss.RiskScoreId(),
		ss.Description(),
		ss.Metadata(),
		ss.CompletedAt(),
		ss.CancelledAt(),
		ss.ExpiredAt(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentsSelectFields() PaymentsSelectFields {
	return PaymentsSelectFields{}
}

type PaymentsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentsUpdateField struct {
	paymentsField PaymentsField
	opt           PaymentsUpdateFieldOption
	value         interface{}
}
type PaymentsUpdateFieldList []PaymentsUpdateField

func defaultPaymentsUpdateFieldOption() PaymentsUpdateFieldOption {
	return PaymentsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentsOption(useIncrement bool) func(*PaymentsUpdateFieldOption) {
	return func(pcufo *PaymentsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentsUpdateField(field PaymentsField, val interface{}, opts ...func(*PaymentsUpdateFieldOption)) PaymentsUpdateField {
	defaultOpt := defaultPaymentsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentsUpdateField{
		paymentsField: field,
		value:         val,
		opt:           defaultOpt,
	}
}
func defaultPaymentsUpdateFields(payments model.Payments) (paymentsUpdateFieldList PaymentsUpdateFieldList) {
	selectFields := NewPaymentsSelectFields()
	paymentsUpdateFieldList = append(paymentsUpdateFieldList,
		NewPaymentsUpdateField(selectFields.Id(), payments.Id),
		NewPaymentsUpdateField(selectFields.PaymentCode(), payments.PaymentCode),
		NewPaymentsUpdateField(selectFields.IntentId(), payments.IntentId),
		NewPaymentsUpdateField(selectFields.AttemptNumber(), payments.AttemptNumber),
		NewPaymentsUpdateField(selectFields.Psp(), payments.Psp),
		NewPaymentsUpdateField(selectFields.PspTransactionId(), payments.PspTransactionId),
		NewPaymentsUpdateField(selectFields.PspReference(), payments.PspReference),
		NewPaymentsUpdateField(selectFields.PspRawRequest(), payments.PspRawRequest),
		NewPaymentsUpdateField(selectFields.PspRawResponse(), payments.PspRawResponse),
		NewPaymentsUpdateField(selectFields.Amount(), payments.Amount),
		NewPaymentsUpdateField(selectFields.Currency(), payments.Currency),
		NewPaymentsUpdateField(selectFields.AmountInSettlementCurrency(), payments.AmountInSettlementCurrency),
		NewPaymentsUpdateField(selectFields.SettlementCurrency(), payments.SettlementCurrency),
		NewPaymentsUpdateField(selectFields.FxRate(), payments.FxRate),
		NewPaymentsUpdateField(selectFields.FxRateSnapshotId(), payments.FxRateSnapshotId),
		NewPaymentsUpdateField(selectFields.PaymentMethodId(), payments.PaymentMethodId),
		NewPaymentsUpdateField(selectFields.PaymentMethodType(), payments.PaymentMethodType),
		NewPaymentsUpdateField(selectFields.Status(), payments.Status),
		NewPaymentsUpdateField(selectFields.FailureCode(), payments.FailureCode),
		NewPaymentsUpdateField(selectFields.FailureMessage(), payments.FailureMessage),
		NewPaymentsUpdateField(selectFields.FailureCategory(), payments.FailureCategory),
		NewPaymentsUpdateField(selectFields.AuthorisedAt(), payments.AuthorisedAt),
		NewPaymentsUpdateField(selectFields.AuthorisedAmount(), payments.AuthorisedAmount),
		NewPaymentsUpdateField(selectFields.CapturedAt(), payments.CapturedAt),
		NewPaymentsUpdateField(selectFields.CapturedAmount(), payments.CapturedAmount),
		NewPaymentsUpdateField(selectFields.ProcessingFee(), payments.ProcessingFee),
		NewPaymentsUpdateField(selectFields.ProcessingFeeCurrency(), payments.ProcessingFeeCurrency),
		NewPaymentsUpdateField(selectFields.RiskScoreId(), payments.RiskScoreId),
		NewPaymentsUpdateField(selectFields.Description(), payments.Description),
		NewPaymentsUpdateField(selectFields.Metadata(), payments.Metadata),
		NewPaymentsUpdateField(selectFields.CompletedAt(), payments.CompletedAt),
		NewPaymentsUpdateField(selectFields.CancelledAt(), payments.CancelledAt),
		NewPaymentsUpdateField(selectFields.ExpiredAt(), payments.ExpiredAt),
		NewPaymentsUpdateField(selectFields.MetaCreatedAt(), payments.MetaCreatedAt),
		NewPaymentsUpdateField(selectFields.MetaCreatedBy(), payments.MetaCreatedBy),
		NewPaymentsUpdateField(selectFields.MetaUpdatedAt(), payments.MetaUpdatedAt),
		NewPaymentsUpdateField(selectFields.MetaUpdatedBy(), payments.MetaUpdatedBy),
		NewPaymentsUpdateField(selectFields.MetaDeletedAt(), payments.MetaDeletedAt),
		NewPaymentsUpdateField(selectFields.MetaDeletedBy(), payments.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentsCommand(paymentsUpdateFieldList PaymentsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentsUpdateFieldList {
		field := string(updateField.paymentsField)
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

func (repo *RepositoryImpl) BulkCreatePayments(ctx context.Context, paymentsList []*model.Payments, fieldsInsert ...PaymentsField) (err error) {
	var (
		fieldsStr         string
		valueListStr      []string
		argsList          []interface{}
		primaryIds        []model.PaymentsPrimaryID
		paymentsValueList []model.Payments
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, payments := range paymentsList {

		primaryIds = append(primaryIds, payments.ToPaymentsPrimaryID())

		paymentsValueList = append(paymentsValueList, *payments)
	}

	_, notFoundIds, err := repo.IsExistPaymentsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayments] failed checking payments whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentsPrimaryID{}
		mapNotFoundIds := map[model.PaymentsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "payments", fmt.Sprintf("payments with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPayments(paymentsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentsQueries.insertPayments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayments] failed exec create payments query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentsByIDs(ctx context.Context, primaryIDs []model.PaymentsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentsByIDs] failed checking payments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payments with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payments\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentsQueries.deletePayments + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentsByIDs(ctx context.Context, ids []model.PaymentsPrimaryID) (exists bool, notFoundIds []model.PaymentsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payments\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentsQueries.selectPayments, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentsPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentsPrimaryID]bool{}
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

// BulkUpdatePayments is used to bulk update payments, by default it will update all field
// if want to update specific field, then fill paymentssMapUpdateFieldsRequest else please fill paymentssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePayments(ctx context.Context, paymentssMap map[model.PaymentsPrimaryID]*model.Payments, paymentssMapUpdateFieldsRequest map[model.PaymentsPrimaryID]PaymentsUpdateFieldList) (err error) {
	if len(paymentssMap) == 0 && len(paymentssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentssMapUpdateField map[model.PaymentsPrimaryID]PaymentsUpdateFieldList = map[model.PaymentsPrimaryID]PaymentsUpdateFieldList{}
		asTableValues           string                                              = "myvalues"
	)

	if len(paymentssMap) > 0 {
		for id, payments := range paymentssMap {
			if payments == nil {
				log.Error().Err(err).Msg("[BulkUpdatePayments] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentssMapUpdateField[id] = defaultPaymentsUpdateFields(*payments)
		}
	} else {
		paymentssMapUpdateField = paymentssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentsQuery(paymentssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayments] failed checking payments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payments with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payments\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayments] failed exec query")
	}
	return
}

type PaymentsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentsFieldParameter(param string, args ...interface{}) PaymentsFieldParameter {
	return PaymentsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentsQuery(mapPaymentss map[model.PaymentsPrimaryID]PaymentsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentsPrimaryID]map[string]interface{}{}
	paymentsSelectFields := NewPaymentsSelectFields()
	for id, updateFields := range mapPaymentss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentsFieldType(updateField.paymentsField)))
			args = append(args, fields[string(updateField.paymentsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentsField))
		if updateField.paymentsField == paymentsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentsField, asTableValues, updateField.paymentsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentsField,
				"\"payments\"", updateField.paymentsField,
				asTableValues, updateField.paymentsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentsPrimaryID, asTableValue string) (whereQry string) {
	paymentsSelectFields := NewPaymentsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payments\".\"id\" = %s.\"id\"::"+GetPaymentsFieldType(paymentsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentsFieldType(paymentsField PaymentsField) string {
	selectPaymentsFields := NewPaymentsSelectFields()
	switch paymentsField {

	case selectPaymentsFields.Id():
		return "uuid"

	case selectPaymentsFields.PaymentCode():
		return "text"

	case selectPaymentsFields.IntentId():
		return "uuid"

	case selectPaymentsFields.AttemptNumber():
		return "int2"

	case selectPaymentsFields.Psp():
		return "psp_enum"

	case selectPaymentsFields.PspTransactionId():
		return "text"

	case selectPaymentsFields.PspReference():
		return "text"

	case selectPaymentsFields.PspRawRequest():
		return "jsonb"

	case selectPaymentsFields.PspRawResponse():
		return "jsonb"

	case selectPaymentsFields.Amount():
		return "numeric"

	case selectPaymentsFields.Currency():
		return "payment_currency"

	case selectPaymentsFields.AmountInSettlementCurrency():
		return "numeric"

	case selectPaymentsFields.SettlementCurrency():
		return "payment_currency"

	case selectPaymentsFields.FxRate():
		return "numeric"

	case selectPaymentsFields.FxRateSnapshotId():
		return "uuid"

	case selectPaymentsFields.PaymentMethodId():
		return "uuid"

	case selectPaymentsFields.PaymentMethodType():
		return "payment_method_type_enum"

	case selectPaymentsFields.Status():
		return "payment_status_enum"

	case selectPaymentsFields.FailureCode():
		return "text"

	case selectPaymentsFields.FailureMessage():
		return "text"

	case selectPaymentsFields.FailureCategory():
		return "text"

	case selectPaymentsFields.AuthorisedAt():
		return "timestamptz"

	case selectPaymentsFields.AuthorisedAmount():
		return "numeric"

	case selectPaymentsFields.CapturedAt():
		return "timestamptz"

	case selectPaymentsFields.CapturedAmount():
		return "numeric"

	case selectPaymentsFields.ProcessingFee():
		return "numeric"

	case selectPaymentsFields.ProcessingFeeCurrency():
		return "payment_currency"

	case selectPaymentsFields.RiskScoreId():
		return "uuid"

	case selectPaymentsFields.Description():
		return "text"

	case selectPaymentsFields.Metadata():
		return "jsonb"

	case selectPaymentsFields.CompletedAt():
		return "timestamptz"

	case selectPaymentsFields.CancelledAt():
		return "timestamptz"

	case selectPaymentsFields.ExpiredAt():
		return "timestamptz"

	case selectPaymentsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePayments(ctx context.Context, payments *model.Payments, fieldsInsert ...PaymentsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentsPrimaryID{
		Id: payments.Id,
	}
	exists, err := repo.IsExistPaymentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayments] failed checking payments whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "payments", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPayments([]model.Payments{*payments}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentsQueries.insertPayments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayments] failed exec create payments query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentsByID(ctx context.Context, primaryID model.PaymentsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentsByID] failed checking payments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentsCompositePrimaryKeyWhere([]model.PaymentsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentsQueries.deletePayments + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentsFilterResult, err error) {
	query, args, err := composePaymentsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentsByFilter] failed compose payments filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentsByFilter] failed get payments by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultPaymentsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := PaymentsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, PaymentsField(filterSelectField))
		}
		selectFields = composePaymentsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(paymentsQueries.selectPayments, selectFields)

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

func (repo *RepositoryImpl) IsExistPaymentsByID(ctx context.Context, primaryID model.PaymentsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentsCompositePrimaryKeyWhere([]model.PaymentsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentsQueries.selectCountPayments, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayments(ctx context.Context, selectFields ...PaymentsField) (paymentsList model.PaymentsList, err error) {
	var (
		defaultPaymentsSelectFields = defaultPaymentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentsSelectFields = composePaymentsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentsQueries.selectPayments, defaultPaymentsSelectFields)

	err = repo.db.Read.Select(&paymentsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayments] failed get payments list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentsByID(ctx context.Context, primaryID model.PaymentsPrimaryID, selectFields ...PaymentsField) (payments model.Payments, err error) {
	var (
		defaultPaymentsSelectFields = defaultPaymentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentsSelectFields = composePaymentsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentsCompositePrimaryKeyWhere([]model.PaymentsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentsQueries.selectPayments+" WHERE "+whereQry, defaultPaymentsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&payments, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("payments with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentsByID] failed get payments")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentsByID(ctx context.Context, primaryID model.PaymentsPrimaryID, payments *model.Payments, paymentsUpdateFields ...PaymentsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayments] failed checking payments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if payments == nil {
		if len(paymentsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		payments = &model.Payments{}
	}
	var (
		defaultPaymentsUpdateFields = defaultPaymentsUpdateFields(*payments)
		tempUpdateField             PaymentsUpdateFieldList
		selectFields                = NewPaymentsSelectFields()
	)
	if len(paymentsUpdateFields) > 0 {
		for _, updateField := range paymentsUpdateFields {
			if updateField.paymentsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentsCompositePrimaryKeyWhere([]model.PaymentsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentsCommand(defaultPaymentsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentsQueries.updatePayments+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayments] error when try to update payments by id")
	}
	return err
}

var (
	paymentsQueries = struct {
		selectPayments      string
		selectCountPayments string
		deletePayments      string
		updatePayments      string
		insertPayments      string
	}{
		selectPayments:      "SELECT %s FROM \"payments\"",
		selectCountPayments: "SELECT COUNT(\"id\") FROM \"payments\"",
		deletePayments:      "DELETE FROM \"payments\"",
		updatePayments:      "UPDATE \"payments\" SET %s ",
		insertPayments:      "INSERT INTO \"payments\" %s VALUES %s",
	}
)

type PaymentsRepository interface {
	CreatePayments(ctx context.Context, payments *model.Payments, fieldsInsert ...PaymentsField) error
	BulkCreatePayments(ctx context.Context, paymentsList []*model.Payments, fieldsInsert ...PaymentsField) error
	ResolvePayments(ctx context.Context, selectFields ...PaymentsField) (model.PaymentsList, error)
	ResolvePaymentsByID(ctx context.Context, primaryID model.PaymentsPrimaryID, selectFields ...PaymentsField) (model.Payments, error)
	UpdatePaymentsByID(ctx context.Context, id model.PaymentsPrimaryID, payments *model.Payments, paymentsUpdateFields ...PaymentsUpdateField) error
	BulkUpdatePayments(ctx context.Context, paymentsListMap map[model.PaymentsPrimaryID]*model.Payments, PaymentssMapUpdateFieldsRequest map[model.PaymentsPrimaryID]PaymentsUpdateFieldList) (err error)
	DeletePaymentsByID(ctx context.Context, id model.PaymentsPrimaryID) error
	BulkDeletePaymentsByIDs(ctx context.Context, ids []model.PaymentsPrimaryID) error
	ResolvePaymentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentsFilterResult, err error)
	IsExistPaymentsByIDs(ctx context.Context, ids []model.PaymentsPrimaryID) (exists bool, notFoundIds []model.PaymentsPrimaryID, err error)
	IsExistPaymentsByID(ctx context.Context, id model.PaymentsPrimaryID) (exists bool, err error)
}
