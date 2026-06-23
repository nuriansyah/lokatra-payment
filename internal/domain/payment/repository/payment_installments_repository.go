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

func composeInsertFieldsAndParamsPaymentInstallments(paymentInstallmentsList []model.PaymentInstallments, fieldsInsert ...PaymentInstallmentsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentInstallmentsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentInstallments := range paymentInstallmentsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentInstallments.Id)
			case selectField.PaymentPlanId():
				args = append(args, paymentInstallments.PaymentPlanId)
			case selectField.PaymentIntentId():
				args = append(args, paymentInstallments.PaymentIntentId)
			case selectField.InstallmentNo():
				args = append(args, paymentInstallments.InstallmentNo)
			case selectField.DueAmount():
				args = append(args, paymentInstallments.DueAmount)
			case selectField.PaidAmount():
				args = append(args, paymentInstallments.PaidAmount)
			case selectField.Currency():
				args = append(args, paymentInstallments.Currency)
			case selectField.DueAt():
				args = append(args, paymentInstallments.DueAt)
			case selectField.Status():
				args = append(args, paymentInstallments.Status)
			case selectField.PaidAt():
				args = append(args, paymentInstallments.PaidAt)
			case selectField.OverdueAt():
				args = append(args, paymentInstallments.OverdueAt)
			case selectField.Metadata():
				args = append(args, paymentInstallments.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentInstallments.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentInstallments.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentInstallments.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentInstallments.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentInstallments.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentInstallments.MetaDeletedBy)

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

func composePaymentInstallmentsCompositePrimaryKeyWhere(primaryIDs []model.PaymentInstallmentsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_installments\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentInstallmentsSelectFields() string {
	fields := NewPaymentInstallmentsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentInstallmentsSelectFields(selectFields ...PaymentInstallmentsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentInstallmentsField string
type PaymentInstallmentsFieldList []PaymentInstallmentsField

type PaymentInstallmentsSelectFields struct {
}

func (ss PaymentInstallmentsSelectFields) Id() PaymentInstallmentsField {
	return PaymentInstallmentsField("id")
}

func (ss PaymentInstallmentsSelectFields) PaymentPlanId() PaymentInstallmentsField {
	return PaymentInstallmentsField("payment_plan_id")
}

func (ss PaymentInstallmentsSelectFields) PaymentIntentId() PaymentInstallmentsField {
	return PaymentInstallmentsField("payment_intent_id")
}

func (ss PaymentInstallmentsSelectFields) InstallmentNo() PaymentInstallmentsField {
	return PaymentInstallmentsField("installment_no")
}

func (ss PaymentInstallmentsSelectFields) DueAmount() PaymentInstallmentsField {
	return PaymentInstallmentsField("due_amount")
}

func (ss PaymentInstallmentsSelectFields) PaidAmount() PaymentInstallmentsField {
	return PaymentInstallmentsField("paid_amount")
}

func (ss PaymentInstallmentsSelectFields) Currency() PaymentInstallmentsField {
	return PaymentInstallmentsField("currency")
}

func (ss PaymentInstallmentsSelectFields) DueAt() PaymentInstallmentsField {
	return PaymentInstallmentsField("due_at")
}

func (ss PaymentInstallmentsSelectFields) Status() PaymentInstallmentsField {
	return PaymentInstallmentsField("status")
}

func (ss PaymentInstallmentsSelectFields) PaidAt() PaymentInstallmentsField {
	return PaymentInstallmentsField("paid_at")
}

func (ss PaymentInstallmentsSelectFields) OverdueAt() PaymentInstallmentsField {
	return PaymentInstallmentsField("overdue_at")
}

func (ss PaymentInstallmentsSelectFields) Metadata() PaymentInstallmentsField {
	return PaymentInstallmentsField("metadata")
}

func (ss PaymentInstallmentsSelectFields) MetaCreatedAt() PaymentInstallmentsField {
	return PaymentInstallmentsField("meta_created_at")
}

func (ss PaymentInstallmentsSelectFields) MetaCreatedBy() PaymentInstallmentsField {
	return PaymentInstallmentsField("meta_created_by")
}

func (ss PaymentInstallmentsSelectFields) MetaUpdatedAt() PaymentInstallmentsField {
	return PaymentInstallmentsField("meta_updated_at")
}

func (ss PaymentInstallmentsSelectFields) MetaUpdatedBy() PaymentInstallmentsField {
	return PaymentInstallmentsField("meta_updated_by")
}

func (ss PaymentInstallmentsSelectFields) MetaDeletedAt() PaymentInstallmentsField {
	return PaymentInstallmentsField("meta_deleted_at")
}

func (ss PaymentInstallmentsSelectFields) MetaDeletedBy() PaymentInstallmentsField {
	return PaymentInstallmentsField("meta_deleted_by")
}

func (ss PaymentInstallmentsSelectFields) All() PaymentInstallmentsFieldList {
	return []PaymentInstallmentsField{
		ss.Id(),
		ss.PaymentPlanId(),
		ss.PaymentIntentId(),
		ss.InstallmentNo(),
		ss.DueAmount(),
		ss.PaidAmount(),
		ss.Currency(),
		ss.DueAt(),
		ss.Status(),
		ss.PaidAt(),
		ss.OverdueAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentInstallmentsSelectFields() PaymentInstallmentsSelectFields {
	return PaymentInstallmentsSelectFields{}
}

type PaymentInstallmentsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentInstallmentsUpdateField struct {
	paymentInstallmentsField PaymentInstallmentsField
	opt                      PaymentInstallmentsUpdateFieldOption
	value                    interface{}
}
type PaymentInstallmentsUpdateFieldList []PaymentInstallmentsUpdateField

func defaultPaymentInstallmentsUpdateFieldOption() PaymentInstallmentsUpdateFieldOption {
	return PaymentInstallmentsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentInstallmentsOption(useIncrement bool) func(*PaymentInstallmentsUpdateFieldOption) {
	return func(pcufo *PaymentInstallmentsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentInstallmentsUpdateField(field PaymentInstallmentsField, val interface{}, opts ...func(*PaymentInstallmentsUpdateFieldOption)) PaymentInstallmentsUpdateField {
	defaultOpt := defaultPaymentInstallmentsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentInstallmentsUpdateField{
		paymentInstallmentsField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultPaymentInstallmentsUpdateFields(paymentInstallments model.PaymentInstallments) (paymentInstallmentsUpdateFieldList PaymentInstallmentsUpdateFieldList) {
	selectFields := NewPaymentInstallmentsSelectFields()
	paymentInstallmentsUpdateFieldList = append(paymentInstallmentsUpdateFieldList,
		NewPaymentInstallmentsUpdateField(selectFields.Id(), paymentInstallments.Id),
		NewPaymentInstallmentsUpdateField(selectFields.PaymentPlanId(), paymentInstallments.PaymentPlanId),
		NewPaymentInstallmentsUpdateField(selectFields.PaymentIntentId(), paymentInstallments.PaymentIntentId),
		NewPaymentInstallmentsUpdateField(selectFields.InstallmentNo(), paymentInstallments.InstallmentNo),
		NewPaymentInstallmentsUpdateField(selectFields.DueAmount(), paymentInstallments.DueAmount),
		NewPaymentInstallmentsUpdateField(selectFields.PaidAmount(), paymentInstallments.PaidAmount),
		NewPaymentInstallmentsUpdateField(selectFields.Currency(), paymentInstallments.Currency),
		NewPaymentInstallmentsUpdateField(selectFields.DueAt(), paymentInstallments.DueAt),
		NewPaymentInstallmentsUpdateField(selectFields.Status(), paymentInstallments.Status),
		NewPaymentInstallmentsUpdateField(selectFields.PaidAt(), paymentInstallments.PaidAt),
		NewPaymentInstallmentsUpdateField(selectFields.OverdueAt(), paymentInstallments.OverdueAt),
		NewPaymentInstallmentsUpdateField(selectFields.Metadata(), paymentInstallments.Metadata),
		NewPaymentInstallmentsUpdateField(selectFields.MetaCreatedAt(), paymentInstallments.MetaCreatedAt),
		NewPaymentInstallmentsUpdateField(selectFields.MetaCreatedBy(), paymentInstallments.MetaCreatedBy),
		NewPaymentInstallmentsUpdateField(selectFields.MetaUpdatedAt(), paymentInstallments.MetaUpdatedAt),
		NewPaymentInstallmentsUpdateField(selectFields.MetaUpdatedBy(), paymentInstallments.MetaUpdatedBy),
		NewPaymentInstallmentsUpdateField(selectFields.MetaDeletedAt(), paymentInstallments.MetaDeletedAt),
		NewPaymentInstallmentsUpdateField(selectFields.MetaDeletedBy(), paymentInstallments.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentInstallmentsCommand(paymentInstallmentsUpdateFieldList PaymentInstallmentsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentInstallmentsUpdateFieldList {
		field := string(updateField.paymentInstallmentsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentInstallments(ctx context.Context, paymentInstallmentsList []*model.PaymentInstallments, fieldsInsert ...PaymentInstallmentsField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.PaymentInstallmentsPrimaryID
		paymentInstallmentsValueList []model.PaymentInstallments
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentInstallmentsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentInstallments := range paymentInstallmentsList {

		primaryIds = append(primaryIds, paymentInstallments.ToPaymentInstallmentsPrimaryID())

		paymentInstallmentsValueList = append(paymentInstallmentsValueList, *paymentInstallments)
	}

	_, notFoundIds, err := repo.IsExistPaymentInstallmentsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentInstallments] failed checking paymentInstallments whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentInstallmentsPrimaryID{}
		mapNotFoundIds := map[model.PaymentInstallmentsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentInstallments", fmt.Sprintf("paymentInstallments with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentInstallments(paymentInstallmentsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentInstallmentsQueries.insertPaymentInstallments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentInstallments] failed exec create paymentInstallments query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentInstallmentsByIDs(ctx context.Context, primaryIDs []model.PaymentInstallmentsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentInstallmentsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentInstallmentsByIDs] failed checking paymentInstallments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstallments with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_installments\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentInstallmentsQueries.deletePaymentInstallments + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentInstallmentsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentInstallmentsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentInstallmentsByIDs(ctx context.Context, ids []model.PaymentInstallmentsPrimaryID) (exists bool, notFoundIds []model.PaymentInstallmentsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_installments\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentInstallmentsQueries.selectPaymentInstallments, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentInstallmentsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentInstallmentsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentInstallmentsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentInstallmentsPrimaryID]bool{}
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

// BulkUpdatePaymentInstallments is used to bulk update paymentInstallments, by default it will update all field
// if want to update specific field, then fill paymentInstallmentssMapUpdateFieldsRequest else please fill paymentInstallmentssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentInstallments(ctx context.Context, paymentInstallmentssMap map[model.PaymentInstallmentsPrimaryID]*model.PaymentInstallments, paymentInstallmentssMapUpdateFieldsRequest map[model.PaymentInstallmentsPrimaryID]PaymentInstallmentsUpdateFieldList) (err error) {
	if len(paymentInstallmentssMap) == 0 && len(paymentInstallmentssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentInstallmentssMapUpdateField map[model.PaymentInstallmentsPrimaryID]PaymentInstallmentsUpdateFieldList = map[model.PaymentInstallmentsPrimaryID]PaymentInstallmentsUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(paymentInstallmentssMap) > 0 {
		for id, paymentInstallments := range paymentInstallmentssMap {
			if paymentInstallments == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentInstallments] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentInstallmentssMapUpdateField[id] = defaultPaymentInstallmentsUpdateFields(*paymentInstallments)
		}
	} else {
		paymentInstallmentssMapUpdateField = paymentInstallmentssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentInstallmentsQuery(paymentInstallmentssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentInstallmentsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentInstallments] failed checking paymentInstallments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstallments with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentInstallmentsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_installments\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentInstallments] failed exec query")
	}
	return
}

type PaymentInstallmentsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentInstallmentsFieldParameter(param string, args ...interface{}) PaymentInstallmentsFieldParameter {
	return PaymentInstallmentsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentInstallmentsQuery(mapPaymentInstallmentss map[model.PaymentInstallmentsPrimaryID]PaymentInstallmentsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentInstallmentsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentInstallmentsPrimaryID]map[string]interface{}{}
	paymentInstallmentsSelectFields := NewPaymentInstallmentsSelectFields()
	for id, updateFields := range mapPaymentInstallmentss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentInstallmentsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentInstallmentss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentInstallmentsFieldType(updateField.paymentInstallmentsField)))
			args = append(args, fields[string(updateField.paymentInstallmentsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentInstallmentsField))
		if updateField.paymentInstallmentsField == paymentInstallmentsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentInstallmentsField, asTableValues, updateField.paymentInstallmentsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentInstallmentsField,
				"\"payment_installments\"", updateField.paymentInstallmentsField,
				asTableValues, updateField.paymentInstallmentsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentInstallmentsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentInstallmentsPrimaryID, asTableValue string) (whereQry string) {
	paymentInstallmentsSelectFields := NewPaymentInstallmentsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_installments\".\"id\" = %s.\"id\"::"+GetPaymentInstallmentsFieldType(paymentInstallmentsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentInstallmentsFieldType(paymentInstallmentsField PaymentInstallmentsField) string {
	selectPaymentInstallmentsFields := NewPaymentInstallmentsSelectFields()
	switch paymentInstallmentsField {

	case selectPaymentInstallmentsFields.Id():
		return "uuid"

	case selectPaymentInstallmentsFields.PaymentPlanId():
		return "uuid"

	case selectPaymentInstallmentsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentInstallmentsFields.InstallmentNo():
		return "int4"

	case selectPaymentInstallmentsFields.DueAmount():
		return "numeric"

	case selectPaymentInstallmentsFields.PaidAmount():
		return "numeric"

	case selectPaymentInstallmentsFields.Currency():
		return "text"

	case selectPaymentInstallmentsFields.DueAt():
		return "timestamptz"

	case selectPaymentInstallmentsFields.Status():
		return "payment_installment_status_enum"

	case selectPaymentInstallmentsFields.PaidAt():
		return "timestamptz"

	case selectPaymentInstallmentsFields.OverdueAt():
		return "timestamptz"

	case selectPaymentInstallmentsFields.Metadata():
		return "jsonb"

	case selectPaymentInstallmentsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentInstallmentsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentInstallmentsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentInstallmentsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentInstallmentsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentInstallmentsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentInstallments(ctx context.Context, paymentInstallments *model.PaymentInstallments, fieldsInsert ...PaymentInstallmentsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentInstallmentsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentInstallmentsPrimaryID{
		Id: paymentInstallments.Id,
	}
	exists, err := repo.IsExistPaymentInstallmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentInstallments] failed checking paymentInstallments whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentInstallments", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentInstallments([]model.PaymentInstallments{*paymentInstallments}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentInstallmentsQueries.insertPaymentInstallments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentInstallments] failed exec create paymentInstallments query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentInstallmentsByID(ctx context.Context, primaryID model.PaymentInstallmentsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentInstallmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentInstallmentsByID] failed checking paymentInstallments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstallments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentInstallmentsCompositePrimaryKeyWhere([]model.PaymentInstallmentsPrimaryID{primaryID})
	commandQuery := paymentInstallmentsQueries.deletePaymentInstallments + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentInstallmentsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentInstallmentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentInstallmentsFilterResult, err error) {
	query, args, err := composePaymentInstallmentsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentInstallmentsByFilter] failed compose paymentInstallments filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentInstallmentsByFilter] failed get paymentInstallments by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentInstallmentsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentInstallmentsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentInstallmentsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentInstallmentsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentInstallmentsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentInstallmentsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 18+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_plan_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_plan_id\"")
			selectedColumns["payment_plan_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["installment_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"installment_no\"")
			selectedColumns["installment_no"] = struct{}{}
		}
		if _, selected := selectedColumns["due_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"due_amount\"")
			selectedColumns["due_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["paid_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"paid_amount\"")
			selectedColumns["paid_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["due_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"due_at\"")
			selectedColumns["due_at"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["paid_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"paid_at\"")
			selectedColumns["paid_at"] = struct{}{}
		}
		if _, selected := selectedColumns["overdue_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"overdue_at\"")
			selectedColumns["overdue_at"] = struct{}{}
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

type paymentInstallmentsFilterPlaceholder struct {
	index int
}

func (p *paymentInstallmentsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentInstallmentsFilterPredicate(filterField model.FilterField, placeholders *paymentInstallmentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentInstallmentsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentInstallmentsFilterSQLExpr(spec)
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

func composePaymentInstallmentsFilterGroup(group model.FilterGroup, placeholders *paymentInstallmentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentInstallmentsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentInstallmentsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentInstallmentsFilterWhereQueries(filter model.Filter, placeholders *paymentInstallmentsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentInstallmentsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentInstallmentsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentInstallmentsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentInstallmentsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentInstallmentsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentInstallmentsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentInstallmentsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentInstallmentsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentInstallmentsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentInstallmentsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentInstallmentsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_installments\" base%s", strings.Join(selectColumns, ","), composePaymentInstallmentsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentInstallmentsByID(ctx context.Context, primaryID model.PaymentInstallmentsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentInstallmentsCompositePrimaryKeyWhere([]model.PaymentInstallmentsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentInstallmentsQueries.selectCountPaymentInstallments, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentInstallmentsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentInstallments(ctx context.Context, selectFields ...PaymentInstallmentsField) (paymentInstallmentsList model.PaymentInstallmentsList, err error) {
	var (
		defaultPaymentInstallmentsSelectFields = defaultPaymentInstallmentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentInstallmentsSelectFields = composePaymentInstallmentsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentInstallmentsQueries.selectPaymentInstallments, defaultPaymentInstallmentsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentInstallmentsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentInstallments] failed get paymentInstallments list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentInstallmentsByID(ctx context.Context, primaryID model.PaymentInstallmentsPrimaryID, selectFields ...PaymentInstallmentsField) (paymentInstallments model.PaymentInstallments, err error) {
	var (
		defaultPaymentInstallmentsSelectFields = defaultPaymentInstallmentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentInstallmentsSelectFields = composePaymentInstallmentsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentInstallmentsCompositePrimaryKeyWhere([]model.PaymentInstallmentsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentInstallmentsQueries.selectPaymentInstallments+" WHERE "+whereQry, defaultPaymentInstallmentsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentInstallments, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentInstallments with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentInstallmentsByID] failed get paymentInstallments")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentInstallmentsByID(ctx context.Context, primaryID model.PaymentInstallmentsPrimaryID, paymentInstallments *model.PaymentInstallments, paymentInstallmentsUpdateFields ...PaymentInstallmentsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentInstallmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstallments] failed checking paymentInstallments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstallments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentInstallments == nil {
		if len(paymentInstallmentsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentInstallmentsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentInstallments = &model.PaymentInstallments{}
	}
	var (
		defaultPaymentInstallmentsUpdateFields = defaultPaymentInstallmentsUpdateFields(*paymentInstallments)
		tempUpdateField                        PaymentInstallmentsUpdateFieldList
		selectFields                           = NewPaymentInstallmentsSelectFields()
	)
	if len(paymentInstallmentsUpdateFields) > 0 {
		for _, updateField := range paymentInstallmentsUpdateFields {
			if updateField.paymentInstallmentsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentInstallmentsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentInstallmentsCompositePrimaryKeyWhere([]model.PaymentInstallmentsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentInstallmentsCommand(defaultPaymentInstallmentsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentInstallmentsQueries.updatePaymentInstallments+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstallments] error when try to update paymentInstallments by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentInstallmentsByFilter(ctx context.Context, filter model.Filter, paymentInstallmentsUpdateFields ...PaymentInstallmentsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentInstallmentsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentInstallmentsUpdateFieldList
		selectFields = NewPaymentInstallmentsSelectFields()
	)
	for _, updateField := range paymentInstallmentsUpdateFields {
		if updateField.paymentInstallmentsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentInstallmentsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentInstallmentsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentInstallmentsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_installments\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstallmentsByFilter] error when try to update paymentInstallments by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstallmentsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentInstallmentsQueries = struct {
		selectPaymentInstallments      string
		selectCountPaymentInstallments string
		deletePaymentInstallments      string
		updatePaymentInstallments      string
		insertPaymentInstallments      string
	}{
		selectPaymentInstallments:      "SELECT %s FROM \"payment_installments\"",
		selectCountPaymentInstallments: "SELECT COUNT(\"id\") FROM \"payment_installments\"",
		deletePaymentInstallments:      "DELETE FROM \"payment_installments\"",
		updatePaymentInstallments:      "UPDATE \"payment_installments\" SET %s ",
		insertPaymentInstallments:      "INSERT INTO \"payment_installments\" %s VALUES %s",
	}
)

type PaymentInstallmentsRepository interface {
	CreatePaymentInstallments(ctx context.Context, paymentInstallments *model.PaymentInstallments, fieldsInsert ...PaymentInstallmentsField) error
	BulkCreatePaymentInstallments(ctx context.Context, paymentInstallmentsList []*model.PaymentInstallments, fieldsInsert ...PaymentInstallmentsField) error
	ResolvePaymentInstallments(ctx context.Context, selectFields ...PaymentInstallmentsField) (model.PaymentInstallmentsList, error)
	ResolvePaymentInstallmentsByID(ctx context.Context, primaryID model.PaymentInstallmentsPrimaryID, selectFields ...PaymentInstallmentsField) (model.PaymentInstallments, error)
	UpdatePaymentInstallmentsByID(ctx context.Context, id model.PaymentInstallmentsPrimaryID, paymentInstallments *model.PaymentInstallments, paymentInstallmentsUpdateFields ...PaymentInstallmentsUpdateField) error
	UpdatePaymentInstallmentsByFilter(ctx context.Context, filter model.Filter, paymentInstallmentsUpdateFields ...PaymentInstallmentsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentInstallments(ctx context.Context, paymentInstallmentsListMap map[model.PaymentInstallmentsPrimaryID]*model.PaymentInstallments, PaymentInstallmentssMapUpdateFieldsRequest map[model.PaymentInstallmentsPrimaryID]PaymentInstallmentsUpdateFieldList) (err error)
	DeletePaymentInstallmentsByID(ctx context.Context, id model.PaymentInstallmentsPrimaryID) error
	BulkDeletePaymentInstallmentsByIDs(ctx context.Context, ids []model.PaymentInstallmentsPrimaryID) error
	ResolvePaymentInstallmentsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentInstallmentsFilterResult, err error)
	IsExistPaymentInstallmentsByIDs(ctx context.Context, ids []model.PaymentInstallmentsPrimaryID) (exists bool, notFoundIds []model.PaymentInstallmentsPrimaryID, err error)
	IsExistPaymentInstallmentsByID(ctx context.Context, id model.PaymentInstallmentsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
