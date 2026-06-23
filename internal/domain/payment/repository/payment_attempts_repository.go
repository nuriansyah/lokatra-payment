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

func composeInsertFieldsAndParamsPaymentAttempts(paymentAttemptsList []model.PaymentAttempts, fieldsInsert ...PaymentAttemptsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentAttemptsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentAttempts := range paymentAttemptsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentAttempts.Id)
			case selectField.PaymentIntentId():
				args = append(args, paymentAttempts.PaymentIntentId)
			case selectField.AttemptNo():
				args = append(args, paymentAttempts.AttemptNo)
			case selectField.ProviderAccountId():
				args = append(args, paymentAttempts.ProviderAccountId)
			case selectField.RouteDecisionId():
				args = append(args, paymentAttempts.RouteDecisionId)
			case selectField.ProviderCode():
				args = append(args, paymentAttempts.ProviderCode)
			case selectField.MethodCode():
				args = append(args, paymentAttempts.MethodCode)
			case selectField.ChannelCode():
				args = append(args, paymentAttempts.ChannelCode)
			case selectField.Amount():
				args = append(args, paymentAttempts.Amount)
			case selectField.Currency():
				args = append(args, paymentAttempts.Currency)
			case selectField.Status():
				args = append(args, paymentAttempts.Status)
			case selectField.ProviderReference():
				args = append(args, paymentAttempts.ProviderReference)
			case selectField.ProviderTransactionId():
				args = append(args, paymentAttempts.ProviderTransactionId)
			case selectField.ProviderOrderId():
				args = append(args, paymentAttempts.ProviderOrderId)
			case selectField.ProviderPaymentId():
				args = append(args, paymentAttempts.ProviderPaymentId)
			case selectField.FailureCode():
				args = append(args, paymentAttempts.FailureCode)
			case selectField.FailureMessage():
				args = append(args, paymentAttempts.FailureMessage)
			case selectField.ExpiresAt():
				args = append(args, paymentAttempts.ExpiresAt)
			case selectField.AuthorizedAt():
				args = append(args, paymentAttempts.AuthorizedAt)
			case selectField.CapturedAt():
				args = append(args, paymentAttempts.CapturedAt)
			case selectField.PaidAt():
				args = append(args, paymentAttempts.PaidAt)
			case selectField.FailedAt():
				args = append(args, paymentAttempts.FailedAt)
			case selectField.CanceledAt():
				args = append(args, paymentAttempts.CanceledAt)
			case selectField.StatusSyncRequiredAt():
				args = append(args, paymentAttempts.StatusSyncRequiredAt)
			case selectField.LastStatusSyncAt():
				args = append(args, paymentAttempts.LastStatusSyncAt)
			case selectField.RawRequest():
				args = append(args, paymentAttempts.RawRequest)
			case selectField.RawResponse():
				args = append(args, paymentAttempts.RawResponse)
			case selectField.Metadata():
				args = append(args, paymentAttempts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentAttempts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentAttempts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentAttempts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentAttempts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentAttempts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentAttempts.MetaDeletedBy)

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

func composePaymentAttemptsCompositePrimaryKeyWhere(primaryIDs []model.PaymentAttemptsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_attempts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentAttemptsSelectFields() string {
	fields := NewPaymentAttemptsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentAttemptsSelectFields(selectFields ...PaymentAttemptsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentAttemptsField string
type PaymentAttemptsFieldList []PaymentAttemptsField

type PaymentAttemptsSelectFields struct {
}

func (ss PaymentAttemptsSelectFields) Id() PaymentAttemptsField {
	return PaymentAttemptsField("id")
}

func (ss PaymentAttemptsSelectFields) PaymentIntentId() PaymentAttemptsField {
	return PaymentAttemptsField("payment_intent_id")
}

func (ss PaymentAttemptsSelectFields) AttemptNo() PaymentAttemptsField {
	return PaymentAttemptsField("attempt_no")
}

func (ss PaymentAttemptsSelectFields) ProviderAccountId() PaymentAttemptsField {
	return PaymentAttemptsField("provider_account_id")
}

func (ss PaymentAttemptsSelectFields) RouteDecisionId() PaymentAttemptsField {
	return PaymentAttemptsField("route_decision_id")
}

func (ss PaymentAttemptsSelectFields) ProviderCode() PaymentAttemptsField {
	return PaymentAttemptsField("provider_code")
}

func (ss PaymentAttemptsSelectFields) MethodCode() PaymentAttemptsField {
	return PaymentAttemptsField("method_code")
}

func (ss PaymentAttemptsSelectFields) ChannelCode() PaymentAttemptsField {
	return PaymentAttemptsField("channel_code")
}

func (ss PaymentAttemptsSelectFields) Amount() PaymentAttemptsField {
	return PaymentAttemptsField("amount")
}

func (ss PaymentAttemptsSelectFields) Currency() PaymentAttemptsField {
	return PaymentAttemptsField("currency")
}

func (ss PaymentAttemptsSelectFields) Status() PaymentAttemptsField {
	return PaymentAttemptsField("status")
}

func (ss PaymentAttemptsSelectFields) ProviderReference() PaymentAttemptsField {
	return PaymentAttemptsField("provider_reference")
}

func (ss PaymentAttemptsSelectFields) ProviderTransactionId() PaymentAttemptsField {
	return PaymentAttemptsField("provider_transaction_id")
}

func (ss PaymentAttemptsSelectFields) ProviderOrderId() PaymentAttemptsField {
	return PaymentAttemptsField("provider_order_id")
}

func (ss PaymentAttemptsSelectFields) ProviderPaymentId() PaymentAttemptsField {
	return PaymentAttemptsField("provider_payment_id")
}

func (ss PaymentAttemptsSelectFields) FailureCode() PaymentAttemptsField {
	return PaymentAttemptsField("failure_code")
}

func (ss PaymentAttemptsSelectFields) FailureMessage() PaymentAttemptsField {
	return PaymentAttemptsField("failure_message")
}

func (ss PaymentAttemptsSelectFields) ExpiresAt() PaymentAttemptsField {
	return PaymentAttemptsField("expires_at")
}

func (ss PaymentAttemptsSelectFields) AuthorizedAt() PaymentAttemptsField {
	return PaymentAttemptsField("authorized_at")
}

func (ss PaymentAttemptsSelectFields) CapturedAt() PaymentAttemptsField {
	return PaymentAttemptsField("captured_at")
}

func (ss PaymentAttemptsSelectFields) PaidAt() PaymentAttemptsField {
	return PaymentAttemptsField("paid_at")
}

func (ss PaymentAttemptsSelectFields) FailedAt() PaymentAttemptsField {
	return PaymentAttemptsField("failed_at")
}

func (ss PaymentAttemptsSelectFields) CanceledAt() PaymentAttemptsField {
	return PaymentAttemptsField("canceled_at")
}

func (ss PaymentAttemptsSelectFields) StatusSyncRequiredAt() PaymentAttemptsField {
	return PaymentAttemptsField("status_sync_required_at")
}

func (ss PaymentAttemptsSelectFields) LastStatusSyncAt() PaymentAttemptsField {
	return PaymentAttemptsField("last_status_sync_at")
}

func (ss PaymentAttemptsSelectFields) RawRequest() PaymentAttemptsField {
	return PaymentAttemptsField("raw_request")
}

func (ss PaymentAttemptsSelectFields) RawResponse() PaymentAttemptsField {
	return PaymentAttemptsField("raw_response")
}

func (ss PaymentAttemptsSelectFields) Metadata() PaymentAttemptsField {
	return PaymentAttemptsField("metadata")
}

func (ss PaymentAttemptsSelectFields) MetaCreatedAt() PaymentAttemptsField {
	return PaymentAttemptsField("meta_created_at")
}

func (ss PaymentAttemptsSelectFields) MetaCreatedBy() PaymentAttemptsField {
	return PaymentAttemptsField("meta_created_by")
}

func (ss PaymentAttemptsSelectFields) MetaUpdatedAt() PaymentAttemptsField {
	return PaymentAttemptsField("meta_updated_at")
}

func (ss PaymentAttemptsSelectFields) MetaUpdatedBy() PaymentAttemptsField {
	return PaymentAttemptsField("meta_updated_by")
}

func (ss PaymentAttemptsSelectFields) MetaDeletedAt() PaymentAttemptsField {
	return PaymentAttemptsField("meta_deleted_at")
}

func (ss PaymentAttemptsSelectFields) MetaDeletedBy() PaymentAttemptsField {
	return PaymentAttemptsField("meta_deleted_by")
}

func (ss PaymentAttemptsSelectFields) All() PaymentAttemptsFieldList {
	return []PaymentAttemptsField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.AttemptNo(),
		ss.ProviderAccountId(),
		ss.RouteDecisionId(),
		ss.ProviderCode(),
		ss.MethodCode(),
		ss.ChannelCode(),
		ss.Amount(),
		ss.Currency(),
		ss.Status(),
		ss.ProviderReference(),
		ss.ProviderTransactionId(),
		ss.ProviderOrderId(),
		ss.ProviderPaymentId(),
		ss.FailureCode(),
		ss.FailureMessage(),
		ss.ExpiresAt(),
		ss.AuthorizedAt(),
		ss.CapturedAt(),
		ss.PaidAt(),
		ss.FailedAt(),
		ss.CanceledAt(),
		ss.StatusSyncRequiredAt(),
		ss.LastStatusSyncAt(),
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

func NewPaymentAttemptsSelectFields() PaymentAttemptsSelectFields {
	return PaymentAttemptsSelectFields{}
}

type PaymentAttemptsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentAttemptsUpdateField struct {
	paymentAttemptsField PaymentAttemptsField
	opt                  PaymentAttemptsUpdateFieldOption
	value                interface{}
}
type PaymentAttemptsUpdateFieldList []PaymentAttemptsUpdateField

func defaultPaymentAttemptsUpdateFieldOption() PaymentAttemptsUpdateFieldOption {
	return PaymentAttemptsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentAttemptsOption(useIncrement bool) func(*PaymentAttemptsUpdateFieldOption) {
	return func(pcufo *PaymentAttemptsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentAttemptsUpdateField(field PaymentAttemptsField, val interface{}, opts ...func(*PaymentAttemptsUpdateFieldOption)) PaymentAttemptsUpdateField {
	defaultOpt := defaultPaymentAttemptsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentAttemptsUpdateField{
		paymentAttemptsField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultPaymentAttemptsUpdateFields(paymentAttempts model.PaymentAttempts) (paymentAttemptsUpdateFieldList PaymentAttemptsUpdateFieldList) {
	selectFields := NewPaymentAttemptsSelectFields()
	paymentAttemptsUpdateFieldList = append(paymentAttemptsUpdateFieldList,
		NewPaymentAttemptsUpdateField(selectFields.Id(), paymentAttempts.Id),
		NewPaymentAttemptsUpdateField(selectFields.PaymentIntentId(), paymentAttempts.PaymentIntentId),
		NewPaymentAttemptsUpdateField(selectFields.AttemptNo(), paymentAttempts.AttemptNo),
		NewPaymentAttemptsUpdateField(selectFields.ProviderAccountId(), paymentAttempts.ProviderAccountId),
		NewPaymentAttemptsUpdateField(selectFields.RouteDecisionId(), paymentAttempts.RouteDecisionId),
		NewPaymentAttemptsUpdateField(selectFields.ProviderCode(), paymentAttempts.ProviderCode),
		NewPaymentAttemptsUpdateField(selectFields.MethodCode(), paymentAttempts.MethodCode),
		NewPaymentAttemptsUpdateField(selectFields.ChannelCode(), paymentAttempts.ChannelCode),
		NewPaymentAttemptsUpdateField(selectFields.Amount(), paymentAttempts.Amount),
		NewPaymentAttemptsUpdateField(selectFields.Currency(), paymentAttempts.Currency),
		NewPaymentAttemptsUpdateField(selectFields.Status(), paymentAttempts.Status),
		NewPaymentAttemptsUpdateField(selectFields.ProviderReference(), paymentAttempts.ProviderReference),
		NewPaymentAttemptsUpdateField(selectFields.ProviderTransactionId(), paymentAttempts.ProviderTransactionId),
		NewPaymentAttemptsUpdateField(selectFields.ProviderOrderId(), paymentAttempts.ProviderOrderId),
		NewPaymentAttemptsUpdateField(selectFields.ProviderPaymentId(), paymentAttempts.ProviderPaymentId),
		NewPaymentAttemptsUpdateField(selectFields.FailureCode(), paymentAttempts.FailureCode),
		NewPaymentAttemptsUpdateField(selectFields.FailureMessage(), paymentAttempts.FailureMessage),
		NewPaymentAttemptsUpdateField(selectFields.ExpiresAt(), paymentAttempts.ExpiresAt),
		NewPaymentAttemptsUpdateField(selectFields.AuthorizedAt(), paymentAttempts.AuthorizedAt),
		NewPaymentAttemptsUpdateField(selectFields.CapturedAt(), paymentAttempts.CapturedAt),
		NewPaymentAttemptsUpdateField(selectFields.PaidAt(), paymentAttempts.PaidAt),
		NewPaymentAttemptsUpdateField(selectFields.FailedAt(), paymentAttempts.FailedAt),
		NewPaymentAttemptsUpdateField(selectFields.CanceledAt(), paymentAttempts.CanceledAt),
		NewPaymentAttemptsUpdateField(selectFields.StatusSyncRequiredAt(), paymentAttempts.StatusSyncRequiredAt),
		NewPaymentAttemptsUpdateField(selectFields.LastStatusSyncAt(), paymentAttempts.LastStatusSyncAt),
		NewPaymentAttemptsUpdateField(selectFields.RawRequest(), paymentAttempts.RawRequest),
		NewPaymentAttemptsUpdateField(selectFields.RawResponse(), paymentAttempts.RawResponse),
		NewPaymentAttemptsUpdateField(selectFields.Metadata(), paymentAttempts.Metadata),
		NewPaymentAttemptsUpdateField(selectFields.MetaCreatedAt(), paymentAttempts.MetaCreatedAt),
		NewPaymentAttemptsUpdateField(selectFields.MetaCreatedBy(), paymentAttempts.MetaCreatedBy),
		NewPaymentAttemptsUpdateField(selectFields.MetaUpdatedAt(), paymentAttempts.MetaUpdatedAt),
		NewPaymentAttemptsUpdateField(selectFields.MetaUpdatedBy(), paymentAttempts.MetaUpdatedBy),
		NewPaymentAttemptsUpdateField(selectFields.MetaDeletedAt(), paymentAttempts.MetaDeletedAt),
		NewPaymentAttemptsUpdateField(selectFields.MetaDeletedBy(), paymentAttempts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentAttemptsCommand(paymentAttemptsUpdateFieldList PaymentAttemptsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentAttemptsUpdateFieldList {
		field := string(updateField.paymentAttemptsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentAttempts(ctx context.Context, paymentAttemptsList []*model.PaymentAttempts, fieldsInsert ...PaymentAttemptsField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.PaymentAttemptsPrimaryID
		paymentAttemptsValueList []model.PaymentAttempts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentAttemptsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentAttempts := range paymentAttemptsList {

		primaryIds = append(primaryIds, paymentAttempts.ToPaymentAttemptsPrimaryID())

		paymentAttemptsValueList = append(paymentAttemptsValueList, *paymentAttempts)
	}

	_, notFoundIds, err := repo.IsExistPaymentAttemptsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentAttempts] failed checking paymentAttempts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentAttemptsPrimaryID{}
		mapNotFoundIds := map[model.PaymentAttemptsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentAttempts", fmt.Sprintf("paymentAttempts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentAttempts(paymentAttemptsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentAttemptsQueries.insertPaymentAttempts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentAttempts] failed exec create paymentAttempts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentAttemptsByIDs(ctx context.Context, primaryIDs []model.PaymentAttemptsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentAttemptsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentAttemptsByIDs] failed checking paymentAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAttempts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_attempts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentAttemptsQueries.deletePaymentAttempts + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentAttemptsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentAttemptsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentAttemptsByIDs(ctx context.Context, ids []model.PaymentAttemptsPrimaryID) (exists bool, notFoundIds []model.PaymentAttemptsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_attempts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentAttemptsQueries.selectPaymentAttempts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentAttemptsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentAttemptsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentAttemptsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentAttemptsPrimaryID]bool{}
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

// BulkUpdatePaymentAttempts is used to bulk update paymentAttempts, by default it will update all field
// if want to update specific field, then fill paymentAttemptssMapUpdateFieldsRequest else please fill paymentAttemptssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentAttempts(ctx context.Context, paymentAttemptssMap map[model.PaymentAttemptsPrimaryID]*model.PaymentAttempts, paymentAttemptssMapUpdateFieldsRequest map[model.PaymentAttemptsPrimaryID]PaymentAttemptsUpdateFieldList) (err error) {
	if len(paymentAttemptssMap) == 0 && len(paymentAttemptssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentAttemptssMapUpdateField map[model.PaymentAttemptsPrimaryID]PaymentAttemptsUpdateFieldList = map[model.PaymentAttemptsPrimaryID]PaymentAttemptsUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(paymentAttemptssMap) > 0 {
		for id, paymentAttempts := range paymentAttemptssMap {
			if paymentAttempts == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentAttempts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentAttemptssMapUpdateField[id] = defaultPaymentAttemptsUpdateFields(*paymentAttempts)
		}
	} else {
		paymentAttemptssMapUpdateField = paymentAttemptssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentAttemptsQuery(paymentAttemptssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentAttemptsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentAttempts] failed checking paymentAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAttempts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentAttemptsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_attempts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentAttempts] failed exec query")
	}
	return
}

type PaymentAttemptsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentAttemptsFieldParameter(param string, args ...interface{}) PaymentAttemptsFieldParameter {
	return PaymentAttemptsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentAttemptsQuery(mapPaymentAttemptss map[model.PaymentAttemptsPrimaryID]PaymentAttemptsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentAttemptsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentAttemptsPrimaryID]map[string]interface{}{}
	paymentAttemptsSelectFields := NewPaymentAttemptsSelectFields()
	for id, updateFields := range mapPaymentAttemptss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentAttemptsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentAttemptss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentAttemptsFieldType(updateField.paymentAttemptsField)))
			args = append(args, fields[string(updateField.paymentAttemptsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentAttemptsField))
		if updateField.paymentAttemptsField == paymentAttemptsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentAttemptsField, asTableValues, updateField.paymentAttemptsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentAttemptsField,
				"\"payment_attempts\"", updateField.paymentAttemptsField,
				asTableValues, updateField.paymentAttemptsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentAttemptsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentAttemptsPrimaryID, asTableValue string) (whereQry string) {
	paymentAttemptsSelectFields := NewPaymentAttemptsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_attempts\".\"id\" = %s.\"id\"::"+GetPaymentAttemptsFieldType(paymentAttemptsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentAttemptsFieldType(paymentAttemptsField PaymentAttemptsField) string {
	selectPaymentAttemptsFields := NewPaymentAttemptsSelectFields()
	switch paymentAttemptsField {

	case selectPaymentAttemptsFields.Id():
		return "uuid"

	case selectPaymentAttemptsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentAttemptsFields.AttemptNo():
		return "int4"

	case selectPaymentAttemptsFields.ProviderAccountId():
		return "uuid"

	case selectPaymentAttemptsFields.RouteDecisionId():
		return "uuid"

	case selectPaymentAttemptsFields.ProviderCode():
		return "text"

	case selectPaymentAttemptsFields.MethodCode():
		return "text"

	case selectPaymentAttemptsFields.ChannelCode():
		return "text"

	case selectPaymentAttemptsFields.Amount():
		return "numeric"

	case selectPaymentAttemptsFields.Currency():
		return "text"

	case selectPaymentAttemptsFields.Status():
		return "payment_attempt_status_enum"

	case selectPaymentAttemptsFields.ProviderReference():
		return "text"

	case selectPaymentAttemptsFields.ProviderTransactionId():
		return "text"

	case selectPaymentAttemptsFields.ProviderOrderId():
		return "text"

	case selectPaymentAttemptsFields.ProviderPaymentId():
		return "text"

	case selectPaymentAttemptsFields.FailureCode():
		return "text"

	case selectPaymentAttemptsFields.FailureMessage():
		return "text"

	case selectPaymentAttemptsFields.ExpiresAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.AuthorizedAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.CapturedAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.PaidAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.FailedAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.CanceledAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.StatusSyncRequiredAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.LastStatusSyncAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.RawRequest():
		return "jsonb"

	case selectPaymentAttemptsFields.RawResponse():
		return "jsonb"

	case selectPaymentAttemptsFields.Metadata():
		return "jsonb"

	case selectPaymentAttemptsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentAttemptsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentAttemptsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentAttemptsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentAttempts(ctx context.Context, paymentAttempts *model.PaymentAttempts, fieldsInsert ...PaymentAttemptsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentAttemptsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentAttemptsPrimaryID{
		Id: paymentAttempts.Id,
	}
	exists, err := repo.IsExistPaymentAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentAttempts] failed checking paymentAttempts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentAttempts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentAttempts([]model.PaymentAttempts{*paymentAttempts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentAttemptsQueries.insertPaymentAttempts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentAttempts] failed exec create paymentAttempts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentAttemptsByID(ctx context.Context, primaryID model.PaymentAttemptsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentAttemptsByID] failed checking paymentAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAttempts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentAttemptsCompositePrimaryKeyWhere([]model.PaymentAttemptsPrimaryID{primaryID})
	commandQuery := paymentAttemptsQueries.deletePaymentAttempts + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentAttemptsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentAttemptsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentAttemptsFilterResult, err error) {
	query, args, err := composePaymentAttemptsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentAttemptsByFilter] failed compose paymentAttempts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentAttemptsByFilter] failed get paymentAttempts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentAttemptsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentAttemptsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentAttemptsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentAttemptsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentAttemptsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 34 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentAttemptsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 34+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_no\"")
			selectedColumns["attempt_no"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["route_decision_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"route_decision_id\"")
			selectedColumns["route_decision_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_code\"")
			selectedColumns["provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["method_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_code\"")
			selectedColumns["method_code"] = struct{}{}
		}
		if _, selected := selectedColumns["channel_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"channel_code\"")
			selectedColumns["channel_code"] = struct{}{}
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
		if _, selected := selectedColumns["provider_reference"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_reference\"")
			selectedColumns["provider_reference"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_transaction_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_transaction_id\"")
			selectedColumns["provider_transaction_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_order_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_order_id\"")
			selectedColumns["provider_order_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_payment_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_payment_id\"")
			selectedColumns["provider_payment_id"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_code\"")
			selectedColumns["failure_code"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_message"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_message\"")
			selectedColumns["failure_message"] = struct{}{}
		}
		if _, selected := selectedColumns["expires_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"expires_at\"")
			selectedColumns["expires_at"] = struct{}{}
		}
		if _, selected := selectedColumns["authorized_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"authorized_at\"")
			selectedColumns["authorized_at"] = struct{}{}
		}
		if _, selected := selectedColumns["captured_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"captured_at\"")
			selectedColumns["captured_at"] = struct{}{}
		}
		if _, selected := selectedColumns["paid_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"paid_at\"")
			selectedColumns["paid_at"] = struct{}{}
		}
		if _, selected := selectedColumns["failed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"failed_at\"")
			selectedColumns["failed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["canceled_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"canceled_at\"")
			selectedColumns["canceled_at"] = struct{}{}
		}
		if _, selected := selectedColumns["status_sync_required_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"status_sync_required_at\"")
			selectedColumns["status_sync_required_at"] = struct{}{}
		}
		if _, selected := selectedColumns["last_status_sync_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"last_status_sync_at\"")
			selectedColumns["last_status_sync_at"] = struct{}{}
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

type paymentAttemptsFilterPlaceholder struct {
	index int
}

func (p *paymentAttemptsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentAttemptsFilterPredicate(filterField model.FilterField, placeholders *paymentAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentAttemptsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentAttemptsFilterSQLExpr(spec)
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

func composePaymentAttemptsFilterGroup(group model.FilterGroup, placeholders *paymentAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentAttemptsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentAttemptsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentAttemptsFilterWhereQueries(filter model.Filter, placeholders *paymentAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentAttemptsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentAttemptsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentAttemptsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentAttemptsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentAttemptsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentAttemptsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentAttemptsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentAttemptsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentAttemptsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentAttemptsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentAttemptsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_attempts\" base%s", strings.Join(selectColumns, ","), composePaymentAttemptsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentAttemptsByID(ctx context.Context, primaryID model.PaymentAttemptsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentAttemptsCompositePrimaryKeyWhere([]model.PaymentAttemptsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentAttemptsQueries.selectCountPaymentAttempts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentAttemptsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentAttempts(ctx context.Context, selectFields ...PaymentAttemptsField) (paymentAttemptsList model.PaymentAttemptsList, err error) {
	var (
		defaultPaymentAttemptsSelectFields = defaultPaymentAttemptsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentAttemptsSelectFields = composePaymentAttemptsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentAttemptsQueries.selectPaymentAttempts, defaultPaymentAttemptsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentAttemptsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentAttempts] failed get paymentAttempts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentAttemptsByID(ctx context.Context, primaryID model.PaymentAttemptsPrimaryID, selectFields ...PaymentAttemptsField) (paymentAttempts model.PaymentAttempts, err error) {
	var (
		defaultPaymentAttemptsSelectFields = defaultPaymentAttemptsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentAttemptsSelectFields = composePaymentAttemptsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentAttemptsCompositePrimaryKeyWhere([]model.PaymentAttemptsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentAttemptsQueries.selectPaymentAttempts+" WHERE "+whereQry, defaultPaymentAttemptsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentAttempts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentAttempts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentAttemptsByID] failed get paymentAttempts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentAttemptsByID(ctx context.Context, primaryID model.PaymentAttemptsPrimaryID, paymentAttempts *model.PaymentAttempts, paymentAttemptsUpdateFields ...PaymentAttemptsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAttempts] failed checking paymentAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAttempts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentAttempts == nil {
		if len(paymentAttemptsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentAttemptsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentAttempts = &model.PaymentAttempts{}
	}
	var (
		defaultPaymentAttemptsUpdateFields = defaultPaymentAttemptsUpdateFields(*paymentAttempts)
		tempUpdateField                    PaymentAttemptsUpdateFieldList
		selectFields                       = NewPaymentAttemptsSelectFields()
	)
	if len(paymentAttemptsUpdateFields) > 0 {
		for _, updateField := range paymentAttemptsUpdateFields {
			if updateField.paymentAttemptsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentAttemptsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentAttemptsCompositePrimaryKeyWhere([]model.PaymentAttemptsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentAttemptsCommand(defaultPaymentAttemptsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentAttemptsQueries.updatePaymentAttempts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAttempts] error when try to update paymentAttempts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentAttemptsByFilter(ctx context.Context, filter model.Filter, paymentAttemptsUpdateFields ...PaymentAttemptsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentAttemptsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentAttemptsUpdateFieldList
		selectFields = NewPaymentAttemptsSelectFields()
	)
	for _, updateField := range paymentAttemptsUpdateFields {
		if updateField.paymentAttemptsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentAttemptsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentAttemptsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentAttemptsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_attempts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAttemptsByFilter] error when try to update paymentAttempts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAttemptsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentAttemptsQueries = struct {
		selectPaymentAttempts      string
		selectCountPaymentAttempts string
		deletePaymentAttempts      string
		updatePaymentAttempts      string
		insertPaymentAttempts      string
	}{
		selectPaymentAttempts:      "SELECT %s FROM \"payment_attempts\"",
		selectCountPaymentAttempts: "SELECT COUNT(\"id\") FROM \"payment_attempts\"",
		deletePaymentAttempts:      "DELETE FROM \"payment_attempts\"",
		updatePaymentAttempts:      "UPDATE \"payment_attempts\" SET %s ",
		insertPaymentAttempts:      "INSERT INTO \"payment_attempts\" %s VALUES %s",
	}
)

type PaymentAttemptsRepository interface {
	CreatePaymentAttempts(ctx context.Context, paymentAttempts *model.PaymentAttempts, fieldsInsert ...PaymentAttemptsField) error
	BulkCreatePaymentAttempts(ctx context.Context, paymentAttemptsList []*model.PaymentAttempts, fieldsInsert ...PaymentAttemptsField) error
	ResolvePaymentAttempts(ctx context.Context, selectFields ...PaymentAttemptsField) (model.PaymentAttemptsList, error)
	ResolvePaymentAttemptsByID(ctx context.Context, primaryID model.PaymentAttemptsPrimaryID, selectFields ...PaymentAttemptsField) (model.PaymentAttempts, error)
	UpdatePaymentAttemptsByID(ctx context.Context, id model.PaymentAttemptsPrimaryID, paymentAttempts *model.PaymentAttempts, paymentAttemptsUpdateFields ...PaymentAttemptsUpdateField) error
	UpdatePaymentAttemptsByFilter(ctx context.Context, filter model.Filter, paymentAttemptsUpdateFields ...PaymentAttemptsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentAttempts(ctx context.Context, paymentAttemptsListMap map[model.PaymentAttemptsPrimaryID]*model.PaymentAttempts, PaymentAttemptssMapUpdateFieldsRequest map[model.PaymentAttemptsPrimaryID]PaymentAttemptsUpdateFieldList) (err error)
	DeletePaymentAttemptsByID(ctx context.Context, id model.PaymentAttemptsPrimaryID) error
	BulkDeletePaymentAttemptsByIDs(ctx context.Context, ids []model.PaymentAttemptsPrimaryID) error
	ResolvePaymentAttemptsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentAttemptsFilterResult, err error)
	IsExistPaymentAttemptsByIDs(ctx context.Context, ids []model.PaymentAttemptsPrimaryID) (exists bool, notFoundIds []model.PaymentAttemptsPrimaryID, err error)
	IsExistPaymentAttemptsByID(ctx context.Context, id model.PaymentAttemptsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
