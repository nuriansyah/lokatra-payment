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

func composeInsertFieldsAndParamsPaymentOverpayments(paymentOverpaymentsList []model.PaymentOverpayments, fieldsInsert ...PaymentOverpaymentsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentOverpaymentsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentOverpayments := range paymentOverpaymentsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentOverpayments.Id)
			case selectField.PaymentIntentId():
				args = append(args, paymentOverpayments.PaymentIntentId)
			case selectField.PaidAttemptId():
				args = append(args, paymentOverpayments.PaidAttemptId)
			case selectField.OverpaidAttemptId():
				args = append(args, paymentOverpayments.OverpaidAttemptId)
			case selectField.ExpectedAmount():
				args = append(args, paymentOverpayments.ExpectedAmount)
			case selectField.ReceivedAmount():
				args = append(args, paymentOverpayments.ReceivedAmount)
			case selectField.OverpaidAmount():
				args = append(args, paymentOverpayments.OverpaidAmount)
			case selectField.Currency():
				args = append(args, paymentOverpayments.Currency)
			case selectField.Status():
				args = append(args, paymentOverpayments.Status)
			case selectField.ResolutionAction():
				args = append(args, paymentOverpayments.ResolutionAction)
			case selectField.ResolutionNotes():
				args = append(args, paymentOverpayments.ResolutionNotes)
			case selectField.ResolvedAt():
				args = append(args, paymentOverpayments.ResolvedAt)
			case selectField.ResolvedBy():
				args = append(args, paymentOverpayments.ResolvedBy)
			case selectField.Metadata():
				args = append(args, paymentOverpayments.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentOverpayments.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentOverpayments.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentOverpayments.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentOverpayments.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentOverpayments.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentOverpayments.MetaDeletedBy)

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

func composePaymentOverpaymentsCompositePrimaryKeyWhere(primaryIDs []model.PaymentOverpaymentsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_overpayments\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentOverpaymentsSelectFields() string {
	fields := NewPaymentOverpaymentsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentOverpaymentsSelectFields(selectFields ...PaymentOverpaymentsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentOverpaymentsField string
type PaymentOverpaymentsFieldList []PaymentOverpaymentsField

type PaymentOverpaymentsSelectFields struct {
}

func (ss PaymentOverpaymentsSelectFields) Id() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("id")
}

func (ss PaymentOverpaymentsSelectFields) PaymentIntentId() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("payment_intent_id")
}

func (ss PaymentOverpaymentsSelectFields) PaidAttemptId() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("paid_attempt_id")
}

func (ss PaymentOverpaymentsSelectFields) OverpaidAttemptId() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("overpaid_attempt_id")
}

func (ss PaymentOverpaymentsSelectFields) ExpectedAmount() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("expected_amount")
}

func (ss PaymentOverpaymentsSelectFields) ReceivedAmount() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("received_amount")
}

func (ss PaymentOverpaymentsSelectFields) OverpaidAmount() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("overpaid_amount")
}

func (ss PaymentOverpaymentsSelectFields) Currency() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("currency")
}

func (ss PaymentOverpaymentsSelectFields) Status() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("status")
}

func (ss PaymentOverpaymentsSelectFields) ResolutionAction() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("resolution_action")
}

func (ss PaymentOverpaymentsSelectFields) ResolutionNotes() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("resolution_notes")
}

func (ss PaymentOverpaymentsSelectFields) ResolvedAt() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("resolved_at")
}

func (ss PaymentOverpaymentsSelectFields) ResolvedBy() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("resolved_by")
}

func (ss PaymentOverpaymentsSelectFields) Metadata() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("metadata")
}

func (ss PaymentOverpaymentsSelectFields) MetaCreatedAt() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("meta_created_at")
}

func (ss PaymentOverpaymentsSelectFields) MetaCreatedBy() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("meta_created_by")
}

func (ss PaymentOverpaymentsSelectFields) MetaUpdatedAt() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("meta_updated_at")
}

func (ss PaymentOverpaymentsSelectFields) MetaUpdatedBy() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("meta_updated_by")
}

func (ss PaymentOverpaymentsSelectFields) MetaDeletedAt() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("meta_deleted_at")
}

func (ss PaymentOverpaymentsSelectFields) MetaDeletedBy() PaymentOverpaymentsField {
	return PaymentOverpaymentsField("meta_deleted_by")
}

func (ss PaymentOverpaymentsSelectFields) All() PaymentOverpaymentsFieldList {
	return []PaymentOverpaymentsField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.PaidAttemptId(),
		ss.OverpaidAttemptId(),
		ss.ExpectedAmount(),
		ss.ReceivedAmount(),
		ss.OverpaidAmount(),
		ss.Currency(),
		ss.Status(),
		ss.ResolutionAction(),
		ss.ResolutionNotes(),
		ss.ResolvedAt(),
		ss.ResolvedBy(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentOverpaymentsSelectFields() PaymentOverpaymentsSelectFields {
	return PaymentOverpaymentsSelectFields{}
}

type PaymentOverpaymentsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentOverpaymentsUpdateField struct {
	paymentOverpaymentsField PaymentOverpaymentsField
	opt                      PaymentOverpaymentsUpdateFieldOption
	value                    interface{}
}
type PaymentOverpaymentsUpdateFieldList []PaymentOverpaymentsUpdateField

func defaultPaymentOverpaymentsUpdateFieldOption() PaymentOverpaymentsUpdateFieldOption {
	return PaymentOverpaymentsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentOverpaymentsOption(useIncrement bool) func(*PaymentOverpaymentsUpdateFieldOption) {
	return func(pcufo *PaymentOverpaymentsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentOverpaymentsUpdateField(field PaymentOverpaymentsField, val interface{}, opts ...func(*PaymentOverpaymentsUpdateFieldOption)) PaymentOverpaymentsUpdateField {
	defaultOpt := defaultPaymentOverpaymentsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentOverpaymentsUpdateField{
		paymentOverpaymentsField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultPaymentOverpaymentsUpdateFields(paymentOverpayments model.PaymentOverpayments) (paymentOverpaymentsUpdateFieldList PaymentOverpaymentsUpdateFieldList) {
	selectFields := NewPaymentOverpaymentsSelectFields()
	paymentOverpaymentsUpdateFieldList = append(paymentOverpaymentsUpdateFieldList,
		NewPaymentOverpaymentsUpdateField(selectFields.Id(), paymentOverpayments.Id),
		NewPaymentOverpaymentsUpdateField(selectFields.PaymentIntentId(), paymentOverpayments.PaymentIntentId),
		NewPaymentOverpaymentsUpdateField(selectFields.PaidAttemptId(), paymentOverpayments.PaidAttemptId),
		NewPaymentOverpaymentsUpdateField(selectFields.OverpaidAttemptId(), paymentOverpayments.OverpaidAttemptId),
		NewPaymentOverpaymentsUpdateField(selectFields.ExpectedAmount(), paymentOverpayments.ExpectedAmount),
		NewPaymentOverpaymentsUpdateField(selectFields.ReceivedAmount(), paymentOverpayments.ReceivedAmount),
		NewPaymentOverpaymentsUpdateField(selectFields.OverpaidAmount(), paymentOverpayments.OverpaidAmount),
		NewPaymentOverpaymentsUpdateField(selectFields.Currency(), paymentOverpayments.Currency),
		NewPaymentOverpaymentsUpdateField(selectFields.Status(), paymentOverpayments.Status),
		NewPaymentOverpaymentsUpdateField(selectFields.ResolutionAction(), paymentOverpayments.ResolutionAction),
		NewPaymentOverpaymentsUpdateField(selectFields.ResolutionNotes(), paymentOverpayments.ResolutionNotes),
		NewPaymentOverpaymentsUpdateField(selectFields.ResolvedAt(), paymentOverpayments.ResolvedAt),
		NewPaymentOverpaymentsUpdateField(selectFields.ResolvedBy(), paymentOverpayments.ResolvedBy),
		NewPaymentOverpaymentsUpdateField(selectFields.Metadata(), paymentOverpayments.Metadata),
		NewPaymentOverpaymentsUpdateField(selectFields.MetaCreatedAt(), paymentOverpayments.MetaCreatedAt),
		NewPaymentOverpaymentsUpdateField(selectFields.MetaCreatedBy(), paymentOverpayments.MetaCreatedBy),
		NewPaymentOverpaymentsUpdateField(selectFields.MetaUpdatedAt(), paymentOverpayments.MetaUpdatedAt),
		NewPaymentOverpaymentsUpdateField(selectFields.MetaUpdatedBy(), paymentOverpayments.MetaUpdatedBy),
		NewPaymentOverpaymentsUpdateField(selectFields.MetaDeletedAt(), paymentOverpayments.MetaDeletedAt),
		NewPaymentOverpaymentsUpdateField(selectFields.MetaDeletedBy(), paymentOverpayments.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentOverpaymentsCommand(paymentOverpaymentsUpdateFieldList PaymentOverpaymentsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentOverpaymentsUpdateFieldList {
		field := string(updateField.paymentOverpaymentsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentOverpayments(ctx context.Context, paymentOverpaymentsList []*model.PaymentOverpayments, fieldsInsert ...PaymentOverpaymentsField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.PaymentOverpaymentsPrimaryID
		paymentOverpaymentsValueList []model.PaymentOverpayments
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentOverpaymentsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentOverpayments := range paymentOverpaymentsList {

		primaryIds = append(primaryIds, paymentOverpayments.ToPaymentOverpaymentsPrimaryID())

		paymentOverpaymentsValueList = append(paymentOverpaymentsValueList, *paymentOverpayments)
	}

	_, notFoundIds, err := repo.IsExistPaymentOverpaymentsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentOverpayments] failed checking paymentOverpayments whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentOverpaymentsPrimaryID{}
		mapNotFoundIds := map[model.PaymentOverpaymentsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentOverpayments", fmt.Sprintf("paymentOverpayments with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentOverpayments(paymentOverpaymentsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentOverpaymentsQueries.insertPaymentOverpayments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentOverpayments] failed exec create paymentOverpayments query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentOverpaymentsByIDs(ctx context.Context, primaryIDs []model.PaymentOverpaymentsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentOverpaymentsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentOverpaymentsByIDs] failed checking paymentOverpayments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentOverpayments with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_overpayments\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentOverpaymentsQueries.deletePaymentOverpayments + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentOverpaymentsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentOverpaymentsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentOverpaymentsByIDs(ctx context.Context, ids []model.PaymentOverpaymentsPrimaryID) (exists bool, notFoundIds []model.PaymentOverpaymentsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_overpayments\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentOverpaymentsQueries.selectPaymentOverpayments, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentOverpaymentsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentOverpaymentsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentOverpaymentsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentOverpaymentsPrimaryID]bool{}
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

// BulkUpdatePaymentOverpayments is used to bulk update paymentOverpayments, by default it will update all field
// if want to update specific field, then fill paymentOverpaymentssMapUpdateFieldsRequest else please fill paymentOverpaymentssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentOverpayments(ctx context.Context, paymentOverpaymentssMap map[model.PaymentOverpaymentsPrimaryID]*model.PaymentOverpayments, paymentOverpaymentssMapUpdateFieldsRequest map[model.PaymentOverpaymentsPrimaryID]PaymentOverpaymentsUpdateFieldList) (err error) {
	if len(paymentOverpaymentssMap) == 0 && len(paymentOverpaymentssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentOverpaymentssMapUpdateField map[model.PaymentOverpaymentsPrimaryID]PaymentOverpaymentsUpdateFieldList = map[model.PaymentOverpaymentsPrimaryID]PaymentOverpaymentsUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(paymentOverpaymentssMap) > 0 {
		for id, paymentOverpayments := range paymentOverpaymentssMap {
			if paymentOverpayments == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentOverpayments] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentOverpaymentssMapUpdateField[id] = defaultPaymentOverpaymentsUpdateFields(*paymentOverpayments)
		}
	} else {
		paymentOverpaymentssMapUpdateField = paymentOverpaymentssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentOverpaymentsQuery(paymentOverpaymentssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentOverpaymentsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentOverpayments] failed checking paymentOverpayments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentOverpayments with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentOverpaymentsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_overpayments\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentOverpayments] failed exec query")
	}
	return
}

type PaymentOverpaymentsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentOverpaymentsFieldParameter(param string, args ...interface{}) PaymentOverpaymentsFieldParameter {
	return PaymentOverpaymentsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentOverpaymentsQuery(mapPaymentOverpaymentss map[model.PaymentOverpaymentsPrimaryID]PaymentOverpaymentsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentOverpaymentsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentOverpaymentsPrimaryID]map[string]interface{}{}
	paymentOverpaymentsSelectFields := NewPaymentOverpaymentsSelectFields()
	for id, updateFields := range mapPaymentOverpaymentss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentOverpaymentsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentOverpaymentss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentOverpaymentsFieldType(updateField.paymentOverpaymentsField)))
			args = append(args, fields[string(updateField.paymentOverpaymentsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentOverpaymentsField))
		if updateField.paymentOverpaymentsField == paymentOverpaymentsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentOverpaymentsField, asTableValues, updateField.paymentOverpaymentsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentOverpaymentsField,
				"\"payment_overpayments\"", updateField.paymentOverpaymentsField,
				asTableValues, updateField.paymentOverpaymentsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentOverpaymentsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentOverpaymentsPrimaryID, asTableValue string) (whereQry string) {
	paymentOverpaymentsSelectFields := NewPaymentOverpaymentsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_overpayments\".\"id\" = %s.\"id\"::"+GetPaymentOverpaymentsFieldType(paymentOverpaymentsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentOverpaymentsFieldType(paymentOverpaymentsField PaymentOverpaymentsField) string {
	selectPaymentOverpaymentsFields := NewPaymentOverpaymentsSelectFields()
	switch paymentOverpaymentsField {

	case selectPaymentOverpaymentsFields.Id():
		return "uuid"

	case selectPaymentOverpaymentsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentOverpaymentsFields.PaidAttemptId():
		return "uuid"

	case selectPaymentOverpaymentsFields.OverpaidAttemptId():
		return "uuid"

	case selectPaymentOverpaymentsFields.ExpectedAmount():
		return "numeric"

	case selectPaymentOverpaymentsFields.ReceivedAmount():
		return "numeric"

	case selectPaymentOverpaymentsFields.OverpaidAmount():
		return "numeric"

	case selectPaymentOverpaymentsFields.Currency():
		return "text"

	case selectPaymentOverpaymentsFields.Status():
		return "text"

	case selectPaymentOverpaymentsFields.ResolutionAction():
		return "text"

	case selectPaymentOverpaymentsFields.ResolutionNotes():
		return "text"

	case selectPaymentOverpaymentsFields.ResolvedAt():
		return "timestamptz"

	case selectPaymentOverpaymentsFields.ResolvedBy():
		return "uuid"

	case selectPaymentOverpaymentsFields.Metadata():
		return "jsonb"

	case selectPaymentOverpaymentsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentOverpaymentsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentOverpaymentsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentOverpaymentsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentOverpaymentsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentOverpaymentsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentOverpayments(ctx context.Context, paymentOverpayments *model.PaymentOverpayments, fieldsInsert ...PaymentOverpaymentsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentOverpaymentsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentOverpaymentsPrimaryID{
		Id: paymentOverpayments.Id,
	}
	exists, err := repo.IsExistPaymentOverpaymentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentOverpayments] failed checking paymentOverpayments whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentOverpayments", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentOverpayments([]model.PaymentOverpayments{*paymentOverpayments}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentOverpaymentsQueries.insertPaymentOverpayments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentOverpayments] failed exec create paymentOverpayments query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentOverpaymentsByID(ctx context.Context, primaryID model.PaymentOverpaymentsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentOverpaymentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentOverpaymentsByID] failed checking paymentOverpayments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentOverpayments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentOverpaymentsCompositePrimaryKeyWhere([]model.PaymentOverpaymentsPrimaryID{primaryID})
	commandQuery := paymentOverpaymentsQueries.deletePaymentOverpayments + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentOverpaymentsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentOverpaymentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentOverpaymentsFilterResult, err error) {
	query, args, err := composePaymentOverpaymentsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentOverpaymentsByFilter] failed compose paymentOverpayments filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentOverpaymentsByFilter] failed get paymentOverpayments by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentOverpaymentsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentOverpaymentsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentOverpaymentsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentOverpaymentsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentOverpaymentsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentOverpaymentsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["paid_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"paid_attempt_id\"")
			selectedColumns["paid_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["overpaid_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"overpaid_attempt_id\"")
			selectedColumns["overpaid_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["expected_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"expected_amount\"")
			selectedColumns["expected_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["received_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"received_amount\"")
			selectedColumns["received_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["overpaid_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"overpaid_amount\"")
			selectedColumns["overpaid_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["resolution_action"]; !selected {
			selectColumns = append(selectColumns, "base.\"resolution_action\"")
			selectedColumns["resolution_action"] = struct{}{}
		}
		if _, selected := selectedColumns["resolution_notes"]; !selected {
			selectColumns = append(selectColumns, "base.\"resolution_notes\"")
			selectedColumns["resolution_notes"] = struct{}{}
		}
		if _, selected := selectedColumns["resolved_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"resolved_at\"")
			selectedColumns["resolved_at"] = struct{}{}
		}
		if _, selected := selectedColumns["resolved_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"resolved_by\"")
			selectedColumns["resolved_by"] = struct{}{}
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

type paymentOverpaymentsFilterPlaceholder struct {
	index int
}

func (p *paymentOverpaymentsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentOverpaymentsFilterPredicate(filterField model.FilterField, placeholders *paymentOverpaymentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentOverpaymentsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentOverpaymentsFilterSQLExpr(spec)
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

func composePaymentOverpaymentsFilterGroup(group model.FilterGroup, placeholders *paymentOverpaymentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentOverpaymentsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentOverpaymentsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentOverpaymentsFilterWhereQueries(filter model.Filter, placeholders *paymentOverpaymentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentOverpaymentsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentOverpaymentsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentOverpaymentsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentOverpaymentsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentOverpaymentsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentOverpaymentsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentOverpaymentsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentOverpaymentsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentOverpaymentsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentOverpaymentsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentOverpaymentsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_overpayments\" base%s", strings.Join(selectColumns, ","), composePaymentOverpaymentsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentOverpaymentsByID(ctx context.Context, primaryID model.PaymentOverpaymentsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentOverpaymentsCompositePrimaryKeyWhere([]model.PaymentOverpaymentsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentOverpaymentsQueries.selectCountPaymentOverpayments, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentOverpaymentsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentOverpayments(ctx context.Context, selectFields ...PaymentOverpaymentsField) (paymentOverpaymentsList model.PaymentOverpaymentsList, err error) {
	var (
		defaultPaymentOverpaymentsSelectFields = defaultPaymentOverpaymentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentOverpaymentsSelectFields = composePaymentOverpaymentsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentOverpaymentsQueries.selectPaymentOverpayments, defaultPaymentOverpaymentsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentOverpaymentsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentOverpayments] failed get paymentOverpayments list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentOverpaymentsByID(ctx context.Context, primaryID model.PaymentOverpaymentsPrimaryID, selectFields ...PaymentOverpaymentsField) (paymentOverpayments model.PaymentOverpayments, err error) {
	var (
		defaultPaymentOverpaymentsSelectFields = defaultPaymentOverpaymentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentOverpaymentsSelectFields = composePaymentOverpaymentsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentOverpaymentsCompositePrimaryKeyWhere([]model.PaymentOverpaymentsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentOverpaymentsQueries.selectPaymentOverpayments+" WHERE "+whereQry, defaultPaymentOverpaymentsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentOverpayments, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentOverpayments with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentOverpaymentsByID] failed get paymentOverpayments")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentOverpaymentsByID(ctx context.Context, primaryID model.PaymentOverpaymentsPrimaryID, paymentOverpayments *model.PaymentOverpayments, paymentOverpaymentsUpdateFields ...PaymentOverpaymentsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentOverpaymentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentOverpayments] failed checking paymentOverpayments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentOverpayments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentOverpayments == nil {
		if len(paymentOverpaymentsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentOverpaymentsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentOverpayments = &model.PaymentOverpayments{}
	}
	var (
		defaultPaymentOverpaymentsUpdateFields = defaultPaymentOverpaymentsUpdateFields(*paymentOverpayments)
		tempUpdateField                        PaymentOverpaymentsUpdateFieldList
		selectFields                           = NewPaymentOverpaymentsSelectFields()
	)
	if len(paymentOverpaymentsUpdateFields) > 0 {
		for _, updateField := range paymentOverpaymentsUpdateFields {
			if updateField.paymentOverpaymentsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentOverpaymentsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentOverpaymentsCompositePrimaryKeyWhere([]model.PaymentOverpaymentsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentOverpaymentsCommand(defaultPaymentOverpaymentsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentOverpaymentsQueries.updatePaymentOverpayments+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentOverpayments] error when try to update paymentOverpayments by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentOverpaymentsByFilter(ctx context.Context, filter model.Filter, paymentOverpaymentsUpdateFields ...PaymentOverpaymentsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentOverpaymentsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentOverpaymentsUpdateFieldList
		selectFields = NewPaymentOverpaymentsSelectFields()
	)
	for _, updateField := range paymentOverpaymentsUpdateFields {
		if updateField.paymentOverpaymentsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentOverpaymentsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentOverpaymentsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentOverpaymentsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_overpayments\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentOverpaymentsByFilter] error when try to update paymentOverpayments by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentOverpaymentsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentOverpaymentsQueries = struct {
		selectPaymentOverpayments      string
		selectCountPaymentOverpayments string
		deletePaymentOverpayments      string
		updatePaymentOverpayments      string
		insertPaymentOverpayments      string
	}{
		selectPaymentOverpayments:      "SELECT %s FROM \"payment_overpayments\"",
		selectCountPaymentOverpayments: "SELECT COUNT(\"id\") FROM \"payment_overpayments\"",
		deletePaymentOverpayments:      "DELETE FROM \"payment_overpayments\"",
		updatePaymentOverpayments:      "UPDATE \"payment_overpayments\" SET %s ",
		insertPaymentOverpayments:      "INSERT INTO \"payment_overpayments\" %s VALUES %s",
	}
)

type PaymentOverpaymentsRepository interface {
	CreatePaymentOverpayments(ctx context.Context, paymentOverpayments *model.PaymentOverpayments, fieldsInsert ...PaymentOverpaymentsField) error
	BulkCreatePaymentOverpayments(ctx context.Context, paymentOverpaymentsList []*model.PaymentOverpayments, fieldsInsert ...PaymentOverpaymentsField) error
	ResolvePaymentOverpayments(ctx context.Context, selectFields ...PaymentOverpaymentsField) (model.PaymentOverpaymentsList, error)
	ResolvePaymentOverpaymentsByID(ctx context.Context, primaryID model.PaymentOverpaymentsPrimaryID, selectFields ...PaymentOverpaymentsField) (model.PaymentOverpayments, error)
	UpdatePaymentOverpaymentsByID(ctx context.Context, id model.PaymentOverpaymentsPrimaryID, paymentOverpayments *model.PaymentOverpayments, paymentOverpaymentsUpdateFields ...PaymentOverpaymentsUpdateField) error
	UpdatePaymentOverpaymentsByFilter(ctx context.Context, filter model.Filter, paymentOverpaymentsUpdateFields ...PaymentOverpaymentsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentOverpayments(ctx context.Context, paymentOverpaymentsListMap map[model.PaymentOverpaymentsPrimaryID]*model.PaymentOverpayments, PaymentOverpaymentssMapUpdateFieldsRequest map[model.PaymentOverpaymentsPrimaryID]PaymentOverpaymentsUpdateFieldList) (err error)
	DeletePaymentOverpaymentsByID(ctx context.Context, id model.PaymentOverpaymentsPrimaryID) error
	BulkDeletePaymentOverpaymentsByIDs(ctx context.Context, ids []model.PaymentOverpaymentsPrimaryID) error
	ResolvePaymentOverpaymentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentOverpaymentsFilterResult, err error)
	IsExistPaymentOverpaymentsByIDs(ctx context.Context, ids []model.PaymentOverpaymentsPrimaryID) (exists bool, notFoundIds []model.PaymentOverpaymentsPrimaryID, err error)
	IsExistPaymentOverpaymentsByID(ctx context.Context, id model.PaymentOverpaymentsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
