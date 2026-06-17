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

func composeInsertFieldsAndParamsPaymentInstructions(paymentInstructionsList []model.PaymentInstructions, fieldsInsert ...PaymentInstructionsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentInstructionsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentInstructions := range paymentInstructionsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentInstructions.Id)
			case selectField.PaymentAttemptId():
				args = append(args, paymentInstructions.PaymentAttemptId)
			case selectField.InstructionType():
				args = append(args, paymentInstructions.InstructionType)
			case selectField.IsActive():
				args = append(args, paymentInstructions.IsActive)
			case selectField.DisplayName():
				args = append(args, paymentInstructions.DisplayName)
			case selectField.AccountNumber():
				args = append(args, paymentInstructions.AccountNumber)
			case selectField.AccountNumberMasked():
				args = append(args, paymentInstructions.AccountNumberMasked)
			case selectField.AccountHolderName():
				args = append(args, paymentInstructions.AccountHolderName)
			case selectField.BankCode():
				args = append(args, paymentInstructions.BankCode)
			case selectField.BillerCode():
				args = append(args, paymentInstructions.BillerCode)
			case selectField.PaymentCode():
				args = append(args, paymentInstructions.PaymentCode)
			case selectField.QrString():
				args = append(args, paymentInstructions.QrString)
			case selectField.QrImageUrl():
				args = append(args, paymentInstructions.QrImageUrl)
			case selectField.CheckoutUrl():
				args = append(args, paymentInstructions.CheckoutUrl)
			case selectField.DeeplinkUrl():
				args = append(args, paymentInstructions.DeeplinkUrl)
			case selectField.RetailOutletCode():
				args = append(args, paymentInstructions.RetailOutletCode)
			case selectField.ExpiresAt():
				args = append(args, paymentInstructions.ExpiresAt)
			case selectField.Metadata():
				args = append(args, paymentInstructions.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentInstructions.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentInstructions.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentInstructions.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentInstructions.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentInstructions.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentInstructions.MetaDeletedBy)

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

func composePaymentInstructionsCompositePrimaryKeyWhere(primaryIDs []model.PaymentInstructionsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_instructions\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentInstructionsSelectFields() string {
	fields := NewPaymentInstructionsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentInstructionsSelectFields(selectFields ...PaymentInstructionsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentInstructionsField string
type PaymentInstructionsFieldList []PaymentInstructionsField

type PaymentInstructionsSelectFields struct {
}

func (ss PaymentInstructionsSelectFields) Id() PaymentInstructionsField {
	return PaymentInstructionsField("id")
}

func (ss PaymentInstructionsSelectFields) PaymentAttemptId() PaymentInstructionsField {
	return PaymentInstructionsField("payment_attempt_id")
}

func (ss PaymentInstructionsSelectFields) InstructionType() PaymentInstructionsField {
	return PaymentInstructionsField("instruction_type")
}

func (ss PaymentInstructionsSelectFields) IsActive() PaymentInstructionsField {
	return PaymentInstructionsField("is_active")
}

func (ss PaymentInstructionsSelectFields) DisplayName() PaymentInstructionsField {
	return PaymentInstructionsField("display_name")
}

func (ss PaymentInstructionsSelectFields) AccountNumber() PaymentInstructionsField {
	return PaymentInstructionsField("account_number")
}

func (ss PaymentInstructionsSelectFields) AccountNumberMasked() PaymentInstructionsField {
	return PaymentInstructionsField("account_number_masked")
}

func (ss PaymentInstructionsSelectFields) AccountHolderName() PaymentInstructionsField {
	return PaymentInstructionsField("account_holder_name")
}

func (ss PaymentInstructionsSelectFields) BankCode() PaymentInstructionsField {
	return PaymentInstructionsField("bank_code")
}

func (ss PaymentInstructionsSelectFields) BillerCode() PaymentInstructionsField {
	return PaymentInstructionsField("biller_code")
}

func (ss PaymentInstructionsSelectFields) PaymentCode() PaymentInstructionsField {
	return PaymentInstructionsField("payment_code")
}

func (ss PaymentInstructionsSelectFields) QrString() PaymentInstructionsField {
	return PaymentInstructionsField("qr_string")
}

func (ss PaymentInstructionsSelectFields) QrImageUrl() PaymentInstructionsField {
	return PaymentInstructionsField("qr_image_url")
}

func (ss PaymentInstructionsSelectFields) CheckoutUrl() PaymentInstructionsField {
	return PaymentInstructionsField("checkout_url")
}

func (ss PaymentInstructionsSelectFields) DeeplinkUrl() PaymentInstructionsField {
	return PaymentInstructionsField("deeplink_url")
}

func (ss PaymentInstructionsSelectFields) RetailOutletCode() PaymentInstructionsField {
	return PaymentInstructionsField("retail_outlet_code")
}

func (ss PaymentInstructionsSelectFields) ExpiresAt() PaymentInstructionsField {
	return PaymentInstructionsField("expires_at")
}

func (ss PaymentInstructionsSelectFields) Metadata() PaymentInstructionsField {
	return PaymentInstructionsField("metadata")
}

func (ss PaymentInstructionsSelectFields) MetaCreatedAt() PaymentInstructionsField {
	return PaymentInstructionsField("meta_created_at")
}

func (ss PaymentInstructionsSelectFields) MetaCreatedBy() PaymentInstructionsField {
	return PaymentInstructionsField("meta_created_by")
}

func (ss PaymentInstructionsSelectFields) MetaUpdatedAt() PaymentInstructionsField {
	return PaymentInstructionsField("meta_updated_at")
}

func (ss PaymentInstructionsSelectFields) MetaUpdatedBy() PaymentInstructionsField {
	return PaymentInstructionsField("meta_updated_by")
}

func (ss PaymentInstructionsSelectFields) MetaDeletedAt() PaymentInstructionsField {
	return PaymentInstructionsField("meta_deleted_at")
}

func (ss PaymentInstructionsSelectFields) MetaDeletedBy() PaymentInstructionsField {
	return PaymentInstructionsField("meta_deleted_by")
}

func (ss PaymentInstructionsSelectFields) All() PaymentInstructionsFieldList {
	return []PaymentInstructionsField{
		ss.Id(),
		ss.PaymentAttemptId(),
		ss.InstructionType(),
		ss.IsActive(),
		ss.DisplayName(),
		ss.AccountNumber(),
		ss.AccountNumberMasked(),
		ss.AccountHolderName(),
		ss.BankCode(),
		ss.BillerCode(),
		ss.PaymentCode(),
		ss.QrString(),
		ss.QrImageUrl(),
		ss.CheckoutUrl(),
		ss.DeeplinkUrl(),
		ss.RetailOutletCode(),
		ss.ExpiresAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentInstructionsSelectFields() PaymentInstructionsSelectFields {
	return PaymentInstructionsSelectFields{}
}

type PaymentInstructionsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentInstructionsUpdateField struct {
	paymentInstructionsField PaymentInstructionsField
	opt                      PaymentInstructionsUpdateFieldOption
	value                    interface{}
}
type PaymentInstructionsUpdateFieldList []PaymentInstructionsUpdateField

func defaultPaymentInstructionsUpdateFieldOption() PaymentInstructionsUpdateFieldOption {
	return PaymentInstructionsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentInstructionsOption(useIncrement bool) func(*PaymentInstructionsUpdateFieldOption) {
	return func(pcufo *PaymentInstructionsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentInstructionsUpdateField(field PaymentInstructionsField, val interface{}, opts ...func(*PaymentInstructionsUpdateFieldOption)) PaymentInstructionsUpdateField {
	defaultOpt := defaultPaymentInstructionsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentInstructionsUpdateField{
		paymentInstructionsField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultPaymentInstructionsUpdateFields(paymentInstructions model.PaymentInstructions) (paymentInstructionsUpdateFieldList PaymentInstructionsUpdateFieldList) {
	selectFields := NewPaymentInstructionsSelectFields()
	paymentInstructionsUpdateFieldList = append(paymentInstructionsUpdateFieldList,
		NewPaymentInstructionsUpdateField(selectFields.Id(), paymentInstructions.Id),
		NewPaymentInstructionsUpdateField(selectFields.PaymentAttemptId(), paymentInstructions.PaymentAttemptId),
		NewPaymentInstructionsUpdateField(selectFields.InstructionType(), paymentInstructions.InstructionType),
		NewPaymentInstructionsUpdateField(selectFields.IsActive(), paymentInstructions.IsActive),
		NewPaymentInstructionsUpdateField(selectFields.DisplayName(), paymentInstructions.DisplayName),
		NewPaymentInstructionsUpdateField(selectFields.AccountNumber(), paymentInstructions.AccountNumber),
		NewPaymentInstructionsUpdateField(selectFields.AccountNumberMasked(), paymentInstructions.AccountNumberMasked),
		NewPaymentInstructionsUpdateField(selectFields.AccountHolderName(), paymentInstructions.AccountHolderName),
		NewPaymentInstructionsUpdateField(selectFields.BankCode(), paymentInstructions.BankCode),
		NewPaymentInstructionsUpdateField(selectFields.BillerCode(), paymentInstructions.BillerCode),
		NewPaymentInstructionsUpdateField(selectFields.PaymentCode(), paymentInstructions.PaymentCode),
		NewPaymentInstructionsUpdateField(selectFields.QrString(), paymentInstructions.QrString),
		NewPaymentInstructionsUpdateField(selectFields.QrImageUrl(), paymentInstructions.QrImageUrl),
		NewPaymentInstructionsUpdateField(selectFields.CheckoutUrl(), paymentInstructions.CheckoutUrl),
		NewPaymentInstructionsUpdateField(selectFields.DeeplinkUrl(), paymentInstructions.DeeplinkUrl),
		NewPaymentInstructionsUpdateField(selectFields.RetailOutletCode(), paymentInstructions.RetailOutletCode),
		NewPaymentInstructionsUpdateField(selectFields.ExpiresAt(), paymentInstructions.ExpiresAt),
		NewPaymentInstructionsUpdateField(selectFields.Metadata(), paymentInstructions.Metadata),
		NewPaymentInstructionsUpdateField(selectFields.MetaCreatedAt(), paymentInstructions.MetaCreatedAt),
		NewPaymentInstructionsUpdateField(selectFields.MetaCreatedBy(), paymentInstructions.MetaCreatedBy),
		NewPaymentInstructionsUpdateField(selectFields.MetaUpdatedAt(), paymentInstructions.MetaUpdatedAt),
		NewPaymentInstructionsUpdateField(selectFields.MetaUpdatedBy(), paymentInstructions.MetaUpdatedBy),
		NewPaymentInstructionsUpdateField(selectFields.MetaDeletedAt(), paymentInstructions.MetaDeletedAt),
		NewPaymentInstructionsUpdateField(selectFields.MetaDeletedBy(), paymentInstructions.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentInstructionsCommand(paymentInstructionsUpdateFieldList PaymentInstructionsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentInstructionsUpdateFieldList {
		field := string(updateField.paymentInstructionsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentInstructions(ctx context.Context, paymentInstructionsList []*model.PaymentInstructions, fieldsInsert ...PaymentInstructionsField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.PaymentInstructionsPrimaryID
		paymentInstructionsValueList []model.PaymentInstructions
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentInstructionsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentInstructions := range paymentInstructionsList {

		primaryIds = append(primaryIds, paymentInstructions.ToPaymentInstructionsPrimaryID())

		paymentInstructionsValueList = append(paymentInstructionsValueList, *paymentInstructions)
	}

	_, notFoundIds, err := repo.IsExistPaymentInstructionsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentInstructions] failed checking paymentInstructions whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentInstructionsPrimaryID{}
		mapNotFoundIds := map[model.PaymentInstructionsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentInstructions", fmt.Sprintf("paymentInstructions with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentInstructions(paymentInstructionsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentInstructionsQueries.insertPaymentInstructions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentInstructions] failed exec create paymentInstructions query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentInstructionsByIDs(ctx context.Context, primaryIDs []model.PaymentInstructionsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentInstructionsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentInstructionsByIDs] failed checking paymentInstructions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstructions with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_instructions\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentInstructionsQueries.deletePaymentInstructions + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentInstructionsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentInstructionsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentInstructionsByIDs(ctx context.Context, ids []model.PaymentInstructionsPrimaryID) (exists bool, notFoundIds []model.PaymentInstructionsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_instructions\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentInstructionsQueries.selectPaymentInstructions, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentInstructionsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentInstructionsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentInstructionsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentInstructionsPrimaryID]bool{}
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

// BulkUpdatePaymentInstructions is used to bulk update paymentInstructions, by default it will update all field
// if want to update specific field, then fill paymentInstructionssMapUpdateFieldsRequest else please fill paymentInstructionssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentInstructions(ctx context.Context, paymentInstructionssMap map[model.PaymentInstructionsPrimaryID]*model.PaymentInstructions, paymentInstructionssMapUpdateFieldsRequest map[model.PaymentInstructionsPrimaryID]PaymentInstructionsUpdateFieldList) (err error) {
	if len(paymentInstructionssMap) == 0 && len(paymentInstructionssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentInstructionssMapUpdateField map[model.PaymentInstructionsPrimaryID]PaymentInstructionsUpdateFieldList = map[model.PaymentInstructionsPrimaryID]PaymentInstructionsUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(paymentInstructionssMap) > 0 {
		for id, paymentInstructions := range paymentInstructionssMap {
			if paymentInstructions == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentInstructions] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentInstructionssMapUpdateField[id] = defaultPaymentInstructionsUpdateFields(*paymentInstructions)
		}
	} else {
		paymentInstructionssMapUpdateField = paymentInstructionssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentInstructionsQuery(paymentInstructionssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentInstructionsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentInstructions] failed checking paymentInstructions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstructions with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentInstructionsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_instructions\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentInstructions] failed exec query")
	}
	return
}

type PaymentInstructionsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentInstructionsFieldParameter(param string, args ...interface{}) PaymentInstructionsFieldParameter {
	return PaymentInstructionsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentInstructionsQuery(mapPaymentInstructionss map[model.PaymentInstructionsPrimaryID]PaymentInstructionsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentInstructionsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentInstructionsPrimaryID]map[string]interface{}{}
	paymentInstructionsSelectFields := NewPaymentInstructionsSelectFields()
	for id, updateFields := range mapPaymentInstructionss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentInstructionsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentInstructionss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentInstructionsFieldType(updateField.paymentInstructionsField)))
			args = append(args, fields[string(updateField.paymentInstructionsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentInstructionsField))
		if updateField.paymentInstructionsField == paymentInstructionsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentInstructionsField, asTableValues, updateField.paymentInstructionsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentInstructionsField,
				"\"payment_instructions\"", updateField.paymentInstructionsField,
				asTableValues, updateField.paymentInstructionsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentInstructionsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentInstructionsPrimaryID, asTableValue string) (whereQry string) {
	paymentInstructionsSelectFields := NewPaymentInstructionsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_instructions\".\"id\" = %s.\"id\"::"+GetPaymentInstructionsFieldType(paymentInstructionsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentInstructionsFieldType(paymentInstructionsField PaymentInstructionsField) string {
	selectPaymentInstructionsFields := NewPaymentInstructionsSelectFields()
	switch paymentInstructionsField {

	case selectPaymentInstructionsFields.Id():
		return "uuid"

	case selectPaymentInstructionsFields.PaymentAttemptId():
		return "uuid"

	case selectPaymentInstructionsFields.InstructionType():
		return "text"

	case selectPaymentInstructionsFields.IsActive():
		return "bool"

	case selectPaymentInstructionsFields.DisplayName():
		return "text"

	case selectPaymentInstructionsFields.AccountNumber():
		return "text"

	case selectPaymentInstructionsFields.AccountNumberMasked():
		return "text"

	case selectPaymentInstructionsFields.AccountHolderName():
		return "text"

	case selectPaymentInstructionsFields.BankCode():
		return "text"

	case selectPaymentInstructionsFields.BillerCode():
		return "text"

	case selectPaymentInstructionsFields.PaymentCode():
		return "text"

	case selectPaymentInstructionsFields.QrString():
		return "text"

	case selectPaymentInstructionsFields.QrImageUrl():
		return "text"

	case selectPaymentInstructionsFields.CheckoutUrl():
		return "text"

	case selectPaymentInstructionsFields.DeeplinkUrl():
		return "text"

	case selectPaymentInstructionsFields.RetailOutletCode():
		return "text"

	case selectPaymentInstructionsFields.ExpiresAt():
		return "timestamptz"

	case selectPaymentInstructionsFields.Metadata():
		return "jsonb"

	case selectPaymentInstructionsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentInstructionsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentInstructionsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentInstructionsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentInstructionsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentInstructionsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentInstructions(ctx context.Context, paymentInstructions *model.PaymentInstructions, fieldsInsert ...PaymentInstructionsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentInstructionsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentInstructionsPrimaryID{
		Id: paymentInstructions.Id,
	}
	exists, err := repo.IsExistPaymentInstructionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentInstructions] failed checking paymentInstructions whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentInstructions", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentInstructions([]model.PaymentInstructions{*paymentInstructions}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentInstructionsQueries.insertPaymentInstructions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentInstructions] failed exec create paymentInstructions query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentInstructionsByID(ctx context.Context, primaryID model.PaymentInstructionsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentInstructionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentInstructionsByID] failed checking paymentInstructions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstructions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentInstructionsCompositePrimaryKeyWhere([]model.PaymentInstructionsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentInstructionsQueries.deletePaymentInstructions + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentInstructionsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentInstructionsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentInstructionsFilterResult, err error) {
	query, args, err := composePaymentInstructionsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentInstructionsByFilter] failed compose paymentInstructions filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentInstructionsByFilter] failed get paymentInstructions by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentInstructionsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentInstructionsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentInstructionsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentInstructionsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentInstructionsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 24 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentInstructionsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 24+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_attempt_id\"")
			selectedColumns["payment_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["instruction_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"instruction_type\"")
			selectedColumns["instruction_type"] = struct{}{}
		}
		if _, selected := selectedColumns["is_active"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_active\"")
			selectedColumns["is_active"] = struct{}{}
		}
		if _, selected := selectedColumns["display_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"display_name\"")
			selectedColumns["display_name"] = struct{}{}
		}
		if _, selected := selectedColumns["account_number"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_number\"")
			selectedColumns["account_number"] = struct{}{}
		}
		if _, selected := selectedColumns["account_number_masked"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_number_masked\"")
			selectedColumns["account_number_masked"] = struct{}{}
		}
		if _, selected := selectedColumns["account_holder_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_holder_name\"")
			selectedColumns["account_holder_name"] = struct{}{}
		}
		if _, selected := selectedColumns["bank_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"bank_code\"")
			selectedColumns["bank_code"] = struct{}{}
		}
		if _, selected := selectedColumns["biller_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"biller_code\"")
			selectedColumns["biller_code"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_code\"")
			selectedColumns["payment_code"] = struct{}{}
		}
		if _, selected := selectedColumns["qr_string"]; !selected {
			selectColumns = append(selectColumns, "base.\"qr_string\"")
			selectedColumns["qr_string"] = struct{}{}
		}
		if _, selected := selectedColumns["qr_image_url"]; !selected {
			selectColumns = append(selectColumns, "base.\"qr_image_url\"")
			selectedColumns["qr_image_url"] = struct{}{}
		}
		if _, selected := selectedColumns["checkout_url"]; !selected {
			selectColumns = append(selectColumns, "base.\"checkout_url\"")
			selectedColumns["checkout_url"] = struct{}{}
		}
		if _, selected := selectedColumns["deeplink_url"]; !selected {
			selectColumns = append(selectColumns, "base.\"deeplink_url\"")
			selectedColumns["deeplink_url"] = struct{}{}
		}
		if _, selected := selectedColumns["retail_outlet_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"retail_outlet_code\"")
			selectedColumns["retail_outlet_code"] = struct{}{}
		}
		if _, selected := selectedColumns["expires_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"expires_at\"")
			selectedColumns["expires_at"] = struct{}{}
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

type paymentInstructionsFilterPlaceholder struct {
	index int
}

func (p *paymentInstructionsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentInstructionsFilterPredicate(filterField model.FilterField, placeholders *paymentInstructionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentInstructionsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentInstructionsFilterSQLExpr(spec)
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

func composePaymentInstructionsFilterGroup(group model.FilterGroup, placeholders *paymentInstructionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentInstructionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentInstructionsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentInstructionsFilterWhereQueries(filter model.Filter, placeholders *paymentInstructionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentInstructionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentInstructionsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentInstructionsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentInstructionsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentInstructionsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentInstructionsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentInstructionsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentInstructionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentInstructionsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentInstructionsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentInstructionsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_instructions\" base%s", strings.Join(selectColumns, ","), composePaymentInstructionsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentInstructionsByID(ctx context.Context, primaryID model.PaymentInstructionsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentInstructionsCompositePrimaryKeyWhere([]model.PaymentInstructionsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentInstructionsQueries.selectCountPaymentInstructions, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentInstructionsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentInstructions(ctx context.Context, selectFields ...PaymentInstructionsField) (paymentInstructionsList model.PaymentInstructionsList, err error) {
	var (
		defaultPaymentInstructionsSelectFields = defaultPaymentInstructionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentInstructionsSelectFields = composePaymentInstructionsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentInstructionsQueries.selectPaymentInstructions, defaultPaymentInstructionsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentInstructionsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentInstructions] failed get paymentInstructions list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentInstructionsByID(ctx context.Context, primaryID model.PaymentInstructionsPrimaryID, selectFields ...PaymentInstructionsField) (paymentInstructions model.PaymentInstructions, err error) {
	var (
		defaultPaymentInstructionsSelectFields = defaultPaymentInstructionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentInstructionsSelectFields = composePaymentInstructionsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentInstructionsCompositePrimaryKeyWhere([]model.PaymentInstructionsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentInstructionsQueries.selectPaymentInstructions+" WHERE "+whereQry, defaultPaymentInstructionsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentInstructions, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentInstructions with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentInstructionsByID] failed get paymentInstructions")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentInstructionsByID(ctx context.Context, primaryID model.PaymentInstructionsPrimaryID, paymentInstructions *model.PaymentInstructions, paymentInstructionsUpdateFields ...PaymentInstructionsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentInstructionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstructions] failed checking paymentInstructions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentInstructions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentInstructions == nil {
		if len(paymentInstructionsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentInstructionsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentInstructions = &model.PaymentInstructions{}
	}
	var (
		defaultPaymentInstructionsUpdateFields = defaultPaymentInstructionsUpdateFields(*paymentInstructions)
		tempUpdateField                        PaymentInstructionsUpdateFieldList
		selectFields                           = NewPaymentInstructionsSelectFields()
	)
	if len(paymentInstructionsUpdateFields) > 0 {
		for _, updateField := range paymentInstructionsUpdateFields {
			if updateField.paymentInstructionsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentInstructionsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentInstructionsCompositePrimaryKeyWhere([]model.PaymentInstructionsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentInstructionsCommand(defaultPaymentInstructionsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentInstructionsQueries.updatePaymentInstructions+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstructions] error when try to update paymentInstructions by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentInstructionsByFilter(ctx context.Context, filter model.Filter, paymentInstructionsUpdateFields ...PaymentInstructionsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentInstructionsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentInstructionsUpdateFieldList
		selectFields = NewPaymentInstructionsSelectFields()
	)
	for _, updateField := range paymentInstructionsUpdateFields {
		if updateField.paymentInstructionsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentInstructionsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentInstructionsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentInstructionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_instructions\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstructionsByFilter] error when try to update paymentInstructions by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentInstructionsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentInstructionsQueries = struct {
		selectPaymentInstructions      string
		selectCountPaymentInstructions string
		deletePaymentInstructions      string
		updatePaymentInstructions      string
		insertPaymentInstructions      string
	}{
		selectPaymentInstructions:      "SELECT %s FROM \"payment_instructions\"",
		selectCountPaymentInstructions: "SELECT COUNT(\"id\") FROM \"payment_instructions\"",
		deletePaymentInstructions:      "DELETE FROM \"payment_instructions\"",
		updatePaymentInstructions:      "UPDATE \"payment_instructions\" SET %s ",
		insertPaymentInstructions:      "INSERT INTO \"payment_instructions\" %s VALUES %s",
	}
)

type PaymentInstructionsRepository interface {
	CreatePaymentInstructions(ctx context.Context, paymentInstructions *model.PaymentInstructions, fieldsInsert ...PaymentInstructionsField) error
	BulkCreatePaymentInstructions(ctx context.Context, paymentInstructionsList []*model.PaymentInstructions, fieldsInsert ...PaymentInstructionsField) error
	ResolvePaymentInstructions(ctx context.Context, selectFields ...PaymentInstructionsField) (model.PaymentInstructionsList, error)
	ResolvePaymentInstructionsByID(ctx context.Context, primaryID model.PaymentInstructionsPrimaryID, selectFields ...PaymentInstructionsField) (model.PaymentInstructions, error)
	UpdatePaymentInstructionsByID(ctx context.Context, id model.PaymentInstructionsPrimaryID, paymentInstructions *model.PaymentInstructions, paymentInstructionsUpdateFields ...PaymentInstructionsUpdateField) error
	UpdatePaymentInstructionsByFilter(ctx context.Context, filter model.Filter, paymentInstructionsUpdateFields ...PaymentInstructionsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentInstructions(ctx context.Context, paymentInstructionsListMap map[model.PaymentInstructionsPrimaryID]*model.PaymentInstructions, PaymentInstructionssMapUpdateFieldsRequest map[model.PaymentInstructionsPrimaryID]PaymentInstructionsUpdateFieldList) (err error)
	DeletePaymentInstructionsByID(ctx context.Context, id model.PaymentInstructionsPrimaryID) error
	BulkDeletePaymentInstructionsByIDs(ctx context.Context, ids []model.PaymentInstructionsPrimaryID) error
	ResolvePaymentInstructionsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentInstructionsFilterResult, err error)
	IsExistPaymentInstructionsByIDs(ctx context.Context, ids []model.PaymentInstructionsPrimaryID) (exists bool, notFoundIds []model.PaymentInstructionsPrimaryID, err error)
	IsExistPaymentInstructionsByID(ctx context.Context, id model.PaymentInstructionsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
