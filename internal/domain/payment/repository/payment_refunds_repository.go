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

func composeInsertFieldsAndParamsPaymentRefunds(paymentRefundsList []model.PaymentRefunds, fieldsInsert ...PaymentRefundsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentRefundsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentRefunds := range paymentRefundsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentRefunds.Id)
			case selectField.PaymentIntentId():
				args = append(args, paymentRefunds.PaymentIntentId)
			case selectField.PaymentAttemptId():
				args = append(args, paymentRefunds.PaymentAttemptId)
			case selectField.RefundCode():
				args = append(args, paymentRefunds.RefundCode)
			case selectField.Amount():
				args = append(args, paymentRefunds.Amount)
			case selectField.Currency():
				args = append(args, paymentRefunds.Currency)
			case selectField.Reason():
				args = append(args, paymentRefunds.Reason)
			case selectField.Status():
				args = append(args, paymentRefunds.Status)
			case selectField.ProviderRefundId():
				args = append(args, paymentRefunds.ProviderRefundId)
			case selectField.ProviderReference():
				args = append(args, paymentRefunds.ProviderReference)
			case selectField.RequestedBy():
				args = append(args, paymentRefunds.RequestedBy)
			case selectField.RequestedAt():
				args = append(args, paymentRefunds.RequestedAt)
			case selectField.ApprovedBy():
				args = append(args, paymentRefunds.ApprovedBy)
			case selectField.ApprovedAt():
				args = append(args, paymentRefunds.ApprovedAt)
			case selectField.RejectedBy():
				args = append(args, paymentRefunds.RejectedBy)
			case selectField.RejectedAt():
				args = append(args, paymentRefunds.RejectedAt)
			case selectField.RejectionReason():
				args = append(args, paymentRefunds.RejectionReason)
			case selectField.ProcessingAt():
				args = append(args, paymentRefunds.ProcessingAt)
			case selectField.SucceededAt():
				args = append(args, paymentRefunds.SucceededAt)
			case selectField.FailedAt():
				args = append(args, paymentRefunds.FailedAt)
			case selectField.FailureCode():
				args = append(args, paymentRefunds.FailureCode)
			case selectField.FailureMessage():
				args = append(args, paymentRefunds.FailureMessage)
			case selectField.RawRequest():
				args = append(args, paymentRefunds.RawRequest)
			case selectField.RawResponse():
				args = append(args, paymentRefunds.RawResponse)
			case selectField.Metadata():
				args = append(args, paymentRefunds.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentRefunds.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentRefunds.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentRefunds.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentRefunds.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentRefunds.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentRefunds.MetaDeletedBy)

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

func composePaymentRefundsCompositePrimaryKeyWhere(primaryIDs []model.PaymentRefundsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_refunds\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentRefundsSelectFields() string {
	fields := NewPaymentRefundsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentRefundsSelectFields(selectFields ...PaymentRefundsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentRefundsField string
type PaymentRefundsFieldList []PaymentRefundsField

type PaymentRefundsSelectFields struct {
}

func (ss PaymentRefundsSelectFields) Id() PaymentRefundsField {
	return PaymentRefundsField("id")
}

func (ss PaymentRefundsSelectFields) PaymentIntentId() PaymentRefundsField {
	return PaymentRefundsField("payment_intent_id")
}

func (ss PaymentRefundsSelectFields) PaymentAttemptId() PaymentRefundsField {
	return PaymentRefundsField("payment_attempt_id")
}

func (ss PaymentRefundsSelectFields) RefundCode() PaymentRefundsField {
	return PaymentRefundsField("refund_code")
}

func (ss PaymentRefundsSelectFields) Amount() PaymentRefundsField {
	return PaymentRefundsField("amount")
}

func (ss PaymentRefundsSelectFields) Currency() PaymentRefundsField {
	return PaymentRefundsField("currency")
}

func (ss PaymentRefundsSelectFields) Reason() PaymentRefundsField {
	return PaymentRefundsField("reason")
}

func (ss PaymentRefundsSelectFields) Status() PaymentRefundsField {
	return PaymentRefundsField("status")
}

func (ss PaymentRefundsSelectFields) ProviderRefundId() PaymentRefundsField {
	return PaymentRefundsField("provider_refund_id")
}

func (ss PaymentRefundsSelectFields) ProviderReference() PaymentRefundsField {
	return PaymentRefundsField("provider_reference")
}

func (ss PaymentRefundsSelectFields) RequestedBy() PaymentRefundsField {
	return PaymentRefundsField("requested_by")
}

func (ss PaymentRefundsSelectFields) RequestedAt() PaymentRefundsField {
	return PaymentRefundsField("requested_at")
}

func (ss PaymentRefundsSelectFields) ApprovedBy() PaymentRefundsField {
	return PaymentRefundsField("approved_by")
}

func (ss PaymentRefundsSelectFields) ApprovedAt() PaymentRefundsField {
	return PaymentRefundsField("approved_at")
}

func (ss PaymentRefundsSelectFields) RejectedBy() PaymentRefundsField {
	return PaymentRefundsField("rejected_by")
}

func (ss PaymentRefundsSelectFields) RejectedAt() PaymentRefundsField {
	return PaymentRefundsField("rejected_at")
}

func (ss PaymentRefundsSelectFields) RejectionReason() PaymentRefundsField {
	return PaymentRefundsField("rejection_reason")
}

func (ss PaymentRefundsSelectFields) ProcessingAt() PaymentRefundsField {
	return PaymentRefundsField("processing_at")
}

func (ss PaymentRefundsSelectFields) SucceededAt() PaymentRefundsField {
	return PaymentRefundsField("succeeded_at")
}

func (ss PaymentRefundsSelectFields) FailedAt() PaymentRefundsField {
	return PaymentRefundsField("failed_at")
}

func (ss PaymentRefundsSelectFields) FailureCode() PaymentRefundsField {
	return PaymentRefundsField("failure_code")
}

func (ss PaymentRefundsSelectFields) FailureMessage() PaymentRefundsField {
	return PaymentRefundsField("failure_message")
}

func (ss PaymentRefundsSelectFields) RawRequest() PaymentRefundsField {
	return PaymentRefundsField("raw_request")
}

func (ss PaymentRefundsSelectFields) RawResponse() PaymentRefundsField {
	return PaymentRefundsField("raw_response")
}

func (ss PaymentRefundsSelectFields) Metadata() PaymentRefundsField {
	return PaymentRefundsField("metadata")
}

func (ss PaymentRefundsSelectFields) MetaCreatedAt() PaymentRefundsField {
	return PaymentRefundsField("meta_created_at")
}

func (ss PaymentRefundsSelectFields) MetaCreatedBy() PaymentRefundsField {
	return PaymentRefundsField("meta_created_by")
}

func (ss PaymentRefundsSelectFields) MetaUpdatedAt() PaymentRefundsField {
	return PaymentRefundsField("meta_updated_at")
}

func (ss PaymentRefundsSelectFields) MetaUpdatedBy() PaymentRefundsField {
	return PaymentRefundsField("meta_updated_by")
}

func (ss PaymentRefundsSelectFields) MetaDeletedAt() PaymentRefundsField {
	return PaymentRefundsField("meta_deleted_at")
}

func (ss PaymentRefundsSelectFields) MetaDeletedBy() PaymentRefundsField {
	return PaymentRefundsField("meta_deleted_by")
}

func (ss PaymentRefundsSelectFields) All() PaymentRefundsFieldList {
	return []PaymentRefundsField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.PaymentAttemptId(),
		ss.RefundCode(),
		ss.Amount(),
		ss.Currency(),
		ss.Reason(),
		ss.Status(),
		ss.ProviderRefundId(),
		ss.ProviderReference(),
		ss.RequestedBy(),
		ss.RequestedAt(),
		ss.ApprovedBy(),
		ss.ApprovedAt(),
		ss.RejectedBy(),
		ss.RejectedAt(),
		ss.RejectionReason(),
		ss.ProcessingAt(),
		ss.SucceededAt(),
		ss.FailedAt(),
		ss.FailureCode(),
		ss.FailureMessage(),
		ss.RawRequest(),
		ss.RawResponse(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentRefundsSelectFields() PaymentRefundsSelectFields {
	return PaymentRefundsSelectFields{}
}

type PaymentRefundsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentRefundsUpdateField struct {
	paymentRefundsField PaymentRefundsField
	opt                 PaymentRefundsUpdateFieldOption
	value               interface{}
}
type PaymentRefundsUpdateFieldList []PaymentRefundsUpdateField

func defaultPaymentRefundsUpdateFieldOption() PaymentRefundsUpdateFieldOption {
	return PaymentRefundsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentRefundsOption(useIncrement bool) func(*PaymentRefundsUpdateFieldOption) {
	return func(pcufo *PaymentRefundsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentRefundsUpdateField(field PaymentRefundsField, val interface{}, opts ...func(*PaymentRefundsUpdateFieldOption)) PaymentRefundsUpdateField {
	defaultOpt := defaultPaymentRefundsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentRefundsUpdateField{
		paymentRefundsField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultPaymentRefundsUpdateFields(paymentRefunds model.PaymentRefunds) (paymentRefundsUpdateFieldList PaymentRefundsUpdateFieldList) {
	selectFields := NewPaymentRefundsSelectFields()
	paymentRefundsUpdateFieldList = append(paymentRefundsUpdateFieldList,
		NewPaymentRefundsUpdateField(selectFields.Id(), paymentRefunds.Id),
		NewPaymentRefundsUpdateField(selectFields.PaymentIntentId(), paymentRefunds.PaymentIntentId),
		NewPaymentRefundsUpdateField(selectFields.PaymentAttemptId(), paymentRefunds.PaymentAttemptId),
		NewPaymentRefundsUpdateField(selectFields.RefundCode(), paymentRefunds.RefundCode),
		NewPaymentRefundsUpdateField(selectFields.Amount(), paymentRefunds.Amount),
		NewPaymentRefundsUpdateField(selectFields.Currency(), paymentRefunds.Currency),
		NewPaymentRefundsUpdateField(selectFields.Reason(), paymentRefunds.Reason),
		NewPaymentRefundsUpdateField(selectFields.Status(), paymentRefunds.Status),
		NewPaymentRefundsUpdateField(selectFields.ProviderRefundId(), paymentRefunds.ProviderRefundId),
		NewPaymentRefundsUpdateField(selectFields.ProviderReference(), paymentRefunds.ProviderReference),
		NewPaymentRefundsUpdateField(selectFields.RequestedBy(), paymentRefunds.RequestedBy),
		NewPaymentRefundsUpdateField(selectFields.RequestedAt(), paymentRefunds.RequestedAt),
		NewPaymentRefundsUpdateField(selectFields.ApprovedBy(), paymentRefunds.ApprovedBy),
		NewPaymentRefundsUpdateField(selectFields.ApprovedAt(), paymentRefunds.ApprovedAt),
		NewPaymentRefundsUpdateField(selectFields.RejectedBy(), paymentRefunds.RejectedBy),
		NewPaymentRefundsUpdateField(selectFields.RejectedAt(), paymentRefunds.RejectedAt),
		NewPaymentRefundsUpdateField(selectFields.RejectionReason(), paymentRefunds.RejectionReason),
		NewPaymentRefundsUpdateField(selectFields.ProcessingAt(), paymentRefunds.ProcessingAt),
		NewPaymentRefundsUpdateField(selectFields.SucceededAt(), paymentRefunds.SucceededAt),
		NewPaymentRefundsUpdateField(selectFields.FailedAt(), paymentRefunds.FailedAt),
		NewPaymentRefundsUpdateField(selectFields.FailureCode(), paymentRefunds.FailureCode),
		NewPaymentRefundsUpdateField(selectFields.FailureMessage(), paymentRefunds.FailureMessage),
		NewPaymentRefundsUpdateField(selectFields.RawRequest(), paymentRefunds.RawRequest),
		NewPaymentRefundsUpdateField(selectFields.RawResponse(), paymentRefunds.RawResponse),
		NewPaymentRefundsUpdateField(selectFields.Metadata(), paymentRefunds.Metadata),
		NewPaymentRefundsUpdateField(selectFields.MetaCreatedAt(), paymentRefunds.MetaCreatedAt),
		NewPaymentRefundsUpdateField(selectFields.MetaCreatedBy(), paymentRefunds.MetaCreatedBy),
		NewPaymentRefundsUpdateField(selectFields.MetaUpdatedAt(), paymentRefunds.MetaUpdatedAt),
		NewPaymentRefundsUpdateField(selectFields.MetaUpdatedBy(), paymentRefunds.MetaUpdatedBy),
		NewPaymentRefundsUpdateField(selectFields.MetaDeletedAt(), paymentRefunds.MetaDeletedAt),
		NewPaymentRefundsUpdateField(selectFields.MetaDeletedBy(), paymentRefunds.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentRefundsCommand(paymentRefundsUpdateFieldList PaymentRefundsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentRefundsUpdateFieldList {
		field := string(updateField.paymentRefundsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentRefunds(ctx context.Context, paymentRefundsList []*model.PaymentRefunds, fieldsInsert ...PaymentRefundsField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.PaymentRefundsPrimaryID
		paymentRefundsValueList []model.PaymentRefunds
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentRefundsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentRefunds := range paymentRefundsList {

		primaryIds = append(primaryIds, paymentRefunds.ToPaymentRefundsPrimaryID())

		paymentRefundsValueList = append(paymentRefundsValueList, *paymentRefunds)
	}

	_, notFoundIds, err := repo.IsExistPaymentRefundsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentRefunds] failed checking paymentRefunds whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentRefundsPrimaryID{}
		mapNotFoundIds := map[model.PaymentRefundsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentRefunds", fmt.Sprintf("paymentRefunds with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentRefunds(paymentRefundsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentRefundsQueries.insertPaymentRefunds, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentRefunds] failed exec create paymentRefunds query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentRefundsByIDs(ctx context.Context, primaryIDs []model.PaymentRefundsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentRefundsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRefundsByIDs] failed checking paymentRefunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRefunds with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_refunds\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentRefundsQueries.deletePaymentRefunds + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRefundsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRefundsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentRefundsByIDs(ctx context.Context, ids []model.PaymentRefundsPrimaryID) (exists bool, notFoundIds []model.PaymentRefundsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_refunds\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentRefundsQueries.selectPaymentRefunds, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRefundsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentRefundsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRefundsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentRefundsPrimaryID]bool{}
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

// BulkUpdatePaymentRefunds is used to bulk update paymentRefunds, by default it will update all field
// if want to update specific field, then fill paymentRefundssMapUpdateFieldsRequest else please fill paymentRefundssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentRefunds(ctx context.Context, paymentRefundssMap map[model.PaymentRefundsPrimaryID]*model.PaymentRefunds, paymentRefundssMapUpdateFieldsRequest map[model.PaymentRefundsPrimaryID]PaymentRefundsUpdateFieldList) (err error) {
	if len(paymentRefundssMap) == 0 && len(paymentRefundssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentRefundssMapUpdateField map[model.PaymentRefundsPrimaryID]PaymentRefundsUpdateFieldList = map[model.PaymentRefundsPrimaryID]PaymentRefundsUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(paymentRefundssMap) > 0 {
		for id, paymentRefunds := range paymentRefundssMap {
			if paymentRefunds == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentRefunds] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentRefundssMapUpdateField[id] = defaultPaymentRefundsUpdateFields(*paymentRefunds)
		}
	} else {
		paymentRefundssMapUpdateField = paymentRefundssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentRefundsQuery(paymentRefundssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentRefundsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentRefunds] failed checking paymentRefunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRefunds with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentRefundsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_refunds\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentRefunds] failed exec query")
	}
	return
}

type PaymentRefundsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentRefundsFieldParameter(param string, args ...interface{}) PaymentRefundsFieldParameter {
	return PaymentRefundsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentRefundsQuery(mapPaymentRefundss map[model.PaymentRefundsPrimaryID]PaymentRefundsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentRefundsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentRefundsPrimaryID]map[string]interface{}{}
	paymentRefundsSelectFields := NewPaymentRefundsSelectFields()
	for id, updateFields := range mapPaymentRefundss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentRefundsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentRefundss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentRefundsFieldType(updateField.paymentRefundsField)))
			args = append(args, fields[string(updateField.paymentRefundsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentRefundsField))
		if updateField.paymentRefundsField == paymentRefundsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentRefundsField, asTableValues, updateField.paymentRefundsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentRefundsField,
				"\"payment_refunds\"", updateField.paymentRefundsField,
				asTableValues, updateField.paymentRefundsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentRefundsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentRefundsPrimaryID, asTableValue string) (whereQry string) {
	paymentRefundsSelectFields := NewPaymentRefundsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_refunds\".\"id\" = %s.\"id\"::"+GetPaymentRefundsFieldType(paymentRefundsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentRefundsFieldType(paymentRefundsField PaymentRefundsField) string {
	selectPaymentRefundsFields := NewPaymentRefundsSelectFields()
	switch paymentRefundsField {

	case selectPaymentRefundsFields.Id():
		return "uuid"

	case selectPaymentRefundsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentRefundsFields.PaymentAttemptId():
		return "uuid"

	case selectPaymentRefundsFields.RefundCode():
		return "text"

	case selectPaymentRefundsFields.Amount():
		return "numeric"

	case selectPaymentRefundsFields.Currency():
		return "text"

	case selectPaymentRefundsFields.Reason():
		return "text"

	case selectPaymentRefundsFields.Status():
		return "payment_refund_status_enum"

	case selectPaymentRefundsFields.ProviderRefundId():
		return "text"

	case selectPaymentRefundsFields.ProviderReference():
		return "text"

	case selectPaymentRefundsFields.RequestedBy():
		return "uuid"

	case selectPaymentRefundsFields.RequestedAt():
		return "timestamptz"

	case selectPaymentRefundsFields.ApprovedBy():
		return "uuid"

	case selectPaymentRefundsFields.ApprovedAt():
		return "timestamptz"

	case selectPaymentRefundsFields.RejectedBy():
		return "uuid"

	case selectPaymentRefundsFields.RejectedAt():
		return "timestamptz"

	case selectPaymentRefundsFields.RejectionReason():
		return "text"

	case selectPaymentRefundsFields.ProcessingAt():
		return "timestamptz"

	case selectPaymentRefundsFields.SucceededAt():
		return "timestamptz"

	case selectPaymentRefundsFields.FailedAt():
		return "timestamptz"

	case selectPaymentRefundsFields.FailureCode():
		return "text"

	case selectPaymentRefundsFields.FailureMessage():
		return "text"

	case selectPaymentRefundsFields.RawRequest():
		return "jsonb"

	case selectPaymentRefundsFields.RawResponse():
		return "jsonb"

	case selectPaymentRefundsFields.Metadata():
		return "jsonb"

	case selectPaymentRefundsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentRefundsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentRefundsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentRefundsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentRefundsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentRefundsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentRefunds(ctx context.Context, paymentRefunds *model.PaymentRefunds, fieldsInsert ...PaymentRefundsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentRefundsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentRefundsPrimaryID{
		Id: paymentRefunds.Id,
	}
	exists, err := repo.IsExistPaymentRefundsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentRefunds] failed checking paymentRefunds whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentRefunds", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentRefunds([]model.PaymentRefunds{*paymentRefunds}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentRefundsQueries.insertPaymentRefunds, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentRefunds] failed exec create paymentRefunds query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentRefundsByID(ctx context.Context, primaryID model.PaymentRefundsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentRefundsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentRefundsByID] failed checking paymentRefunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRefunds with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentRefundsCompositePrimaryKeyWhere([]model.PaymentRefundsPrimaryID{primaryID})
	commandQuery := paymentRefundsQueries.deletePaymentRefunds + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentRefundsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRefundsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentRefundsFilterResult, err error) {
	query, args, err := composePaymentRefundsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRefundsByFilter] failed compose paymentRefunds filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRefundsByFilter] failed get paymentRefunds by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentRefundsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentRefundsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentRefundsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentRefundsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentRefundsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 31 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentRefundsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 31+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_attempt_id\"")
			selectedColumns["payment_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["refund_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_code\"")
			selectedColumns["refund_code"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason\"")
			selectedColumns["reason"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_refund_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_refund_id\"")
			selectedColumns["provider_refund_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_reference"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_reference\"")
			selectedColumns["provider_reference"] = struct{}{}
		}
		if _, selected := selectedColumns["requested_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"requested_by\"")
			selectedColumns["requested_by"] = struct{}{}
		}
		if _, selected := selectedColumns["requested_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"requested_at\"")
			selectedColumns["requested_at"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_by\"")
			selectedColumns["approved_by"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_at\"")
			selectedColumns["approved_at"] = struct{}{}
		}
		if _, selected := selectedColumns["rejected_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"rejected_by\"")
			selectedColumns["rejected_by"] = struct{}{}
		}
		if _, selected := selectedColumns["rejected_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"rejected_at\"")
			selectedColumns["rejected_at"] = struct{}{}
		}
		if _, selected := selectedColumns["rejection_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"rejection_reason\"")
			selectedColumns["rejection_reason"] = struct{}{}
		}
		if _, selected := selectedColumns["processing_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"processing_at\"")
			selectedColumns["processing_at"] = struct{}{}
		}
		if _, selected := selectedColumns["succeeded_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"succeeded_at\"")
			selectedColumns["succeeded_at"] = struct{}{}
		}
		if _, selected := selectedColumns["failed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"failed_at\"")
			selectedColumns["failed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_code\"")
			selectedColumns["failure_code"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_message"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_message\"")
			selectedColumns["failure_message"] = struct{}{}
		}
		if _, selected := selectedColumns["raw_request"]; !selected {
			selectColumns = append(selectColumns, "base.\"raw_request\"")
			selectedColumns["raw_request"] = struct{}{}
		}
		if _, selected := selectedColumns["raw_response"]; !selected {
			selectColumns = append(selectColumns, "base.\"raw_response\"")
			selectedColumns["raw_response"] = struct{}{}
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

type paymentRefundsFilterPlaceholder struct {
	index int
}

func (p *paymentRefundsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentRefundsFilterPredicate(filterField model.FilterField, placeholders *paymentRefundsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentRefundsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentRefundsFilterSQLExpr(spec)
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

func composePaymentRefundsFilterGroup(group model.FilterGroup, placeholders *paymentRefundsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentRefundsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentRefundsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentRefundsFilterWhereQueries(filter model.Filter, placeholders *paymentRefundsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentRefundsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentRefundsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentRefundsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentRefundsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentRefundsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentRefundsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentRefundsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentRefundsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentRefundsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentRefundsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentRefundsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_refunds\" base%s", strings.Join(selectColumns, ","), composePaymentRefundsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentRefundsByID(ctx context.Context, primaryID model.PaymentRefundsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentRefundsCompositePrimaryKeyWhere([]model.PaymentRefundsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentRefundsQueries.selectCountPaymentRefunds, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRefundsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRefunds(ctx context.Context, selectFields ...PaymentRefundsField) (paymentRefundsList model.PaymentRefundsList, err error) {
	var (
		defaultPaymentRefundsSelectFields = defaultPaymentRefundsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentRefundsSelectFields = composePaymentRefundsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentRefundsQueries.selectPaymentRefunds, defaultPaymentRefundsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentRefundsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRefunds] failed get paymentRefunds list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRefundsByID(ctx context.Context, primaryID model.PaymentRefundsPrimaryID, selectFields ...PaymentRefundsField) (paymentRefunds model.PaymentRefunds, err error) {
	var (
		defaultPaymentRefundsSelectFields = defaultPaymentRefundsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentRefundsSelectFields = composePaymentRefundsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentRefundsCompositePrimaryKeyWhere([]model.PaymentRefundsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentRefundsQueries.selectPaymentRefunds+" WHERE "+whereQry, defaultPaymentRefundsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentRefunds, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentRefunds with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentRefundsByID] failed get paymentRefunds")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentRefundsByID(ctx context.Context, primaryID model.PaymentRefundsPrimaryID, paymentRefunds *model.PaymentRefunds, paymentRefundsUpdateFields ...PaymentRefundsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentRefundsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRefunds] failed checking paymentRefunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRefunds with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentRefunds == nil {
		if len(paymentRefundsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentRefundsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentRefunds = &model.PaymentRefunds{}
	}
	var (
		defaultPaymentRefundsUpdateFields = defaultPaymentRefundsUpdateFields(*paymentRefunds)
		tempUpdateField                   PaymentRefundsUpdateFieldList
		selectFields                      = NewPaymentRefundsSelectFields()
	)
	if len(paymentRefundsUpdateFields) > 0 {
		for _, updateField := range paymentRefundsUpdateFields {
			if updateField.paymentRefundsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentRefundsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentRefundsCompositePrimaryKeyWhere([]model.PaymentRefundsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentRefundsCommand(defaultPaymentRefundsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentRefundsQueries.updatePaymentRefunds+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRefunds] error when try to update paymentRefunds by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentRefundsByFilter(ctx context.Context, filter model.Filter, paymentRefundsUpdateFields ...PaymentRefundsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentRefundsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentRefundsUpdateFieldList
		selectFields = NewPaymentRefundsSelectFields()
	)
	for _, updateField := range paymentRefundsUpdateFields {
		if updateField.paymentRefundsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentRefundsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentRefundsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentRefundsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_refunds\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRefundsByFilter] error when try to update paymentRefunds by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRefundsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentRefundsQueries = struct {
		selectPaymentRefunds      string
		selectCountPaymentRefunds string
		deletePaymentRefunds      string
		updatePaymentRefunds      string
		insertPaymentRefunds      string
	}{
		selectPaymentRefunds:      "SELECT %s FROM \"payment_refunds\"",
		selectCountPaymentRefunds: "SELECT COUNT(\"id\") FROM \"payment_refunds\"",
		deletePaymentRefunds:      "DELETE FROM \"payment_refunds\"",
		updatePaymentRefunds:      "UPDATE \"payment_refunds\" SET %s ",
		insertPaymentRefunds:      "INSERT INTO \"payment_refunds\" %s VALUES %s",
	}
)

type PaymentRefundsRepository interface {
	CreatePaymentRefunds(ctx context.Context, paymentRefunds *model.PaymentRefunds, fieldsInsert ...PaymentRefundsField) error
	BulkCreatePaymentRefunds(ctx context.Context, paymentRefundsList []*model.PaymentRefunds, fieldsInsert ...PaymentRefundsField) error
	ResolvePaymentRefunds(ctx context.Context, selectFields ...PaymentRefundsField) (model.PaymentRefundsList, error)
	ResolvePaymentRefundsByID(ctx context.Context, primaryID model.PaymentRefundsPrimaryID, selectFields ...PaymentRefundsField) (model.PaymentRefunds, error)
	UpdatePaymentRefundsByID(ctx context.Context, id model.PaymentRefundsPrimaryID, paymentRefunds *model.PaymentRefunds, paymentRefundsUpdateFields ...PaymentRefundsUpdateField) error
	UpdatePaymentRefundsByFilter(ctx context.Context, filter model.Filter, paymentRefundsUpdateFields ...PaymentRefundsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentRefunds(ctx context.Context, paymentRefundsListMap map[model.PaymentRefundsPrimaryID]*model.PaymentRefunds, PaymentRefundssMapUpdateFieldsRequest map[model.PaymentRefundsPrimaryID]PaymentRefundsUpdateFieldList) (err error)
	DeletePaymentRefundsByID(ctx context.Context, id model.PaymentRefundsPrimaryID) error
	BulkDeletePaymentRefundsByIDs(ctx context.Context, ids []model.PaymentRefundsPrimaryID) error
	ResolvePaymentRefundsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentRefundsFilterResult, err error)
	IsExistPaymentRefundsByIDs(ctx context.Context, ids []model.PaymentRefundsPrimaryID) (exists bool, notFoundIds []model.PaymentRefundsPrimaryID, err error)
	IsExistPaymentRefundsByID(ctx context.Context, id model.PaymentRefundsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
