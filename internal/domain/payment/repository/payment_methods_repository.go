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

func composeInsertFieldsAndParamsPaymentMethods(paymentMethodsList []model.PaymentMethods, fieldsInsert ...PaymentMethodsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentMethodsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentMethods := range paymentMethodsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentMethods.Id)
			case selectField.UserId():
				args = append(args, paymentMethods.UserId)
			case selectField.MerchantId():
				args = append(args, paymentMethods.MerchantId)
			case selectField.MethodType():
				args = append(args, paymentMethods.MethodType)
			case selectField.Psp():
				args = append(args, paymentMethods.Psp)
			case selectField.TokenRef():
				args = append(args, paymentMethods.TokenRef)
			case selectField.TokenType():
				args = append(args, paymentMethods.TokenType)
			case selectField.TokenExpiresAt():
				args = append(args, paymentMethods.TokenExpiresAt)
			case selectField.CardBrand():
				args = append(args, paymentMethods.CardBrand)
			case selectField.CardLastFour():
				args = append(args, paymentMethods.CardLastFour)
			case selectField.CardExpMonth():
				args = append(args, paymentMethods.CardExpMonth)
			case selectField.CardExpYear():
				args = append(args, paymentMethods.CardExpYear)
			case selectField.CardCountry():
				args = append(args, paymentMethods.CardCountry)
			case selectField.CardFundingType():
				args = append(args, paymentMethods.CardFundingType)
			case selectField.CardBin():
				args = append(args, paymentMethods.CardBin)
			case selectField.WalletAccountRef():
				args = append(args, paymentMethods.WalletAccountRef)
			case selectField.VaBankCode():
				args = append(args, paymentMethods.VaBankCode)
			case selectField.DisplayLabel():
				args = append(args, paymentMethods.DisplayLabel)
			case selectField.IsDefault():
				args = append(args, paymentMethods.IsDefault)
			case selectField.IsActive():
				args = append(args, paymentMethods.IsActive)
			case selectField.VerifiedAt():
				args = append(args, paymentMethods.VerifiedAt)
			case selectField.Fingerprint():
				args = append(args, paymentMethods.Fingerprint)
			case selectField.GdprErasureRequestedAt():
				args = append(args, paymentMethods.GdprErasureRequestedAt)
			case selectField.GdprErasedAt():
				args = append(args, paymentMethods.GdprErasedAt)
			case selectField.MetaCreatedAt():
				args = append(args, paymentMethods.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentMethods.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentMethods.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentMethods.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentMethods.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentMethods.MetaDeletedBy)

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

func composePaymentMethodsCompositePrimaryKeyWhere(primaryIDs []model.PaymentMethodsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_methods\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentMethodsSelectFields() string {
	fields := NewPaymentMethodsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentMethodsSelectFields(selectFields ...PaymentMethodsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentMethodsField string
type PaymentMethodsFieldList []PaymentMethodsField

type PaymentMethodsSelectFields struct {
}

func (ss PaymentMethodsSelectFields) Id() PaymentMethodsField {
	return PaymentMethodsField("id")
}

func (ss PaymentMethodsSelectFields) UserId() PaymentMethodsField {
	return PaymentMethodsField("user_id")
}

func (ss PaymentMethodsSelectFields) MerchantId() PaymentMethodsField {
	return PaymentMethodsField("merchant_id")
}

func (ss PaymentMethodsSelectFields) MethodType() PaymentMethodsField {
	return PaymentMethodsField("method_type")
}

func (ss PaymentMethodsSelectFields) Psp() PaymentMethodsField {
	return PaymentMethodsField("psp")
}

func (ss PaymentMethodsSelectFields) TokenRef() PaymentMethodsField {
	return PaymentMethodsField("token_ref")
}

func (ss PaymentMethodsSelectFields) TokenType() PaymentMethodsField {
	return PaymentMethodsField("token_type")
}

func (ss PaymentMethodsSelectFields) TokenExpiresAt() PaymentMethodsField {
	return PaymentMethodsField("token_expires_at")
}

func (ss PaymentMethodsSelectFields) CardBrand() PaymentMethodsField {
	return PaymentMethodsField("card_brand")
}

func (ss PaymentMethodsSelectFields) CardLastFour() PaymentMethodsField {
	return PaymentMethodsField("card_last_four")
}

func (ss PaymentMethodsSelectFields) CardExpMonth() PaymentMethodsField {
	return PaymentMethodsField("card_exp_month")
}

func (ss PaymentMethodsSelectFields) CardExpYear() PaymentMethodsField {
	return PaymentMethodsField("card_exp_year")
}

func (ss PaymentMethodsSelectFields) CardCountry() PaymentMethodsField {
	return PaymentMethodsField("card_country")
}

func (ss PaymentMethodsSelectFields) CardFundingType() PaymentMethodsField {
	return PaymentMethodsField("card_funding_type")
}

func (ss PaymentMethodsSelectFields) CardBin() PaymentMethodsField {
	return PaymentMethodsField("card_bin")
}

func (ss PaymentMethodsSelectFields) WalletAccountRef() PaymentMethodsField {
	return PaymentMethodsField("wallet_account_ref")
}

func (ss PaymentMethodsSelectFields) VaBankCode() PaymentMethodsField {
	return PaymentMethodsField("va_bank_code")
}

func (ss PaymentMethodsSelectFields) DisplayLabel() PaymentMethodsField {
	return PaymentMethodsField("display_label")
}

func (ss PaymentMethodsSelectFields) IsDefault() PaymentMethodsField {
	return PaymentMethodsField("is_default")
}

func (ss PaymentMethodsSelectFields) IsActive() PaymentMethodsField {
	return PaymentMethodsField("is_active")
}

func (ss PaymentMethodsSelectFields) VerifiedAt() PaymentMethodsField {
	return PaymentMethodsField("verified_at")
}

func (ss PaymentMethodsSelectFields) Fingerprint() PaymentMethodsField {
	return PaymentMethodsField("fingerprint")
}

func (ss PaymentMethodsSelectFields) GdprErasureRequestedAt() PaymentMethodsField {
	return PaymentMethodsField("gdpr_erasure_requested_at")
}

func (ss PaymentMethodsSelectFields) GdprErasedAt() PaymentMethodsField {
	return PaymentMethodsField("gdpr_erased_at")
}

func (ss PaymentMethodsSelectFields) MetaCreatedAt() PaymentMethodsField {
	return PaymentMethodsField("meta_created_at")
}

func (ss PaymentMethodsSelectFields) MetaCreatedBy() PaymentMethodsField {
	return PaymentMethodsField("meta_created_by")
}

func (ss PaymentMethodsSelectFields) MetaUpdatedAt() PaymentMethodsField {
	return PaymentMethodsField("meta_updated_at")
}

func (ss PaymentMethodsSelectFields) MetaUpdatedBy() PaymentMethodsField {
	return PaymentMethodsField("meta_updated_by")
}

func (ss PaymentMethodsSelectFields) MetaDeletedAt() PaymentMethodsField {
	return PaymentMethodsField("meta_deleted_at")
}

func (ss PaymentMethodsSelectFields) MetaDeletedBy() PaymentMethodsField {
	return PaymentMethodsField("meta_deleted_by")
}

func (ss PaymentMethodsSelectFields) All() PaymentMethodsFieldList {
	return []PaymentMethodsField{
		ss.Id(),
		ss.UserId(),
		ss.MerchantId(),
		ss.MethodType(),
		ss.Psp(),
		ss.TokenRef(),
		ss.TokenType(),
		ss.TokenExpiresAt(),
		ss.CardBrand(),
		ss.CardLastFour(),
		ss.CardExpMonth(),
		ss.CardExpYear(),
		ss.CardCountry(),
		ss.CardFundingType(),
		ss.CardBin(),
		ss.WalletAccountRef(),
		ss.VaBankCode(),
		ss.DisplayLabel(),
		ss.IsDefault(),
		ss.IsActive(),
		ss.VerifiedAt(),
		ss.Fingerprint(),
		ss.GdprErasureRequestedAt(),
		ss.GdprErasedAt(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentMethodsSelectFields() PaymentMethodsSelectFields {
	return PaymentMethodsSelectFields{}
}

type PaymentMethodsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentMethodsUpdateField struct {
	paymentMethodsField PaymentMethodsField
	opt                 PaymentMethodsUpdateFieldOption
	value               interface{}
}
type PaymentMethodsUpdateFieldList []PaymentMethodsUpdateField

func defaultPaymentMethodsUpdateFieldOption() PaymentMethodsUpdateFieldOption {
	return PaymentMethodsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentMethodsOption(useIncrement bool) func(*PaymentMethodsUpdateFieldOption) {
	return func(pcufo *PaymentMethodsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentMethodsUpdateField(field PaymentMethodsField, val interface{}, opts ...func(*PaymentMethodsUpdateFieldOption)) PaymentMethodsUpdateField {
	defaultOpt := defaultPaymentMethodsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentMethodsUpdateField{
		paymentMethodsField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultPaymentMethodsUpdateFields(paymentMethods model.PaymentMethods) (paymentMethodsUpdateFieldList PaymentMethodsUpdateFieldList) {
	selectFields := NewPaymentMethodsSelectFields()
	paymentMethodsUpdateFieldList = append(paymentMethodsUpdateFieldList,
		NewPaymentMethodsUpdateField(selectFields.Id(), paymentMethods.Id),
		NewPaymentMethodsUpdateField(selectFields.UserId(), paymentMethods.UserId),
		NewPaymentMethodsUpdateField(selectFields.MerchantId(), paymentMethods.MerchantId),
		NewPaymentMethodsUpdateField(selectFields.MethodType(), paymentMethods.MethodType),
		NewPaymentMethodsUpdateField(selectFields.Psp(), paymentMethods.Psp),
		NewPaymentMethodsUpdateField(selectFields.TokenRef(), paymentMethods.TokenRef),
		NewPaymentMethodsUpdateField(selectFields.TokenType(), paymentMethods.TokenType),
		NewPaymentMethodsUpdateField(selectFields.TokenExpiresAt(), paymentMethods.TokenExpiresAt),
		NewPaymentMethodsUpdateField(selectFields.CardBrand(), paymentMethods.CardBrand),
		NewPaymentMethodsUpdateField(selectFields.CardLastFour(), paymentMethods.CardLastFour),
		NewPaymentMethodsUpdateField(selectFields.CardExpMonth(), paymentMethods.CardExpMonth),
		NewPaymentMethodsUpdateField(selectFields.CardExpYear(), paymentMethods.CardExpYear),
		NewPaymentMethodsUpdateField(selectFields.CardCountry(), paymentMethods.CardCountry),
		NewPaymentMethodsUpdateField(selectFields.CardFundingType(), paymentMethods.CardFundingType),
		NewPaymentMethodsUpdateField(selectFields.CardBin(), paymentMethods.CardBin),
		NewPaymentMethodsUpdateField(selectFields.WalletAccountRef(), paymentMethods.WalletAccountRef),
		NewPaymentMethodsUpdateField(selectFields.VaBankCode(), paymentMethods.VaBankCode),
		NewPaymentMethodsUpdateField(selectFields.DisplayLabel(), paymentMethods.DisplayLabel),
		NewPaymentMethodsUpdateField(selectFields.IsDefault(), paymentMethods.IsDefault),
		NewPaymentMethodsUpdateField(selectFields.IsActive(), paymentMethods.IsActive),
		NewPaymentMethodsUpdateField(selectFields.VerifiedAt(), paymentMethods.VerifiedAt),
		NewPaymentMethodsUpdateField(selectFields.Fingerprint(), paymentMethods.Fingerprint),
		NewPaymentMethodsUpdateField(selectFields.GdprErasureRequestedAt(), paymentMethods.GdprErasureRequestedAt),
		NewPaymentMethodsUpdateField(selectFields.GdprErasedAt(), paymentMethods.GdprErasedAt),
		NewPaymentMethodsUpdateField(selectFields.MetaCreatedAt(), paymentMethods.MetaCreatedAt),
		NewPaymentMethodsUpdateField(selectFields.MetaCreatedBy(), paymentMethods.MetaCreatedBy),
		NewPaymentMethodsUpdateField(selectFields.MetaUpdatedAt(), paymentMethods.MetaUpdatedAt),
		NewPaymentMethodsUpdateField(selectFields.MetaUpdatedBy(), paymentMethods.MetaUpdatedBy),
		NewPaymentMethodsUpdateField(selectFields.MetaDeletedAt(), paymentMethods.MetaDeletedAt),
		NewPaymentMethodsUpdateField(selectFields.MetaDeletedBy(), paymentMethods.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentMethodsCommand(paymentMethodsUpdateFieldList PaymentMethodsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentMethodsUpdateFieldList {
		field := string(updateField.paymentMethodsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentMethods(ctx context.Context, paymentMethodsList []*model.PaymentMethods, fieldsInsert ...PaymentMethodsField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.PaymentMethodsPrimaryID
		paymentMethodsValueList []model.PaymentMethods
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentMethodsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentMethods := range paymentMethodsList {

		primaryIds = append(primaryIds, paymentMethods.ToPaymentMethodsPrimaryID())

		paymentMethodsValueList = append(paymentMethodsValueList, *paymentMethods)
	}

	_, notFoundIds, err := repo.IsExistPaymentMethodsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentMethods] failed checking paymentMethods whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentMethodsPrimaryID{}
		mapNotFoundIds := map[model.PaymentMethodsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentMethods", fmt.Sprintf("paymentMethods with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentMethods(paymentMethodsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentMethodsQueries.insertPaymentMethods, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentMethods] failed exec create paymentMethods query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentMethodsByIDs(ctx context.Context, primaryIDs []model.PaymentMethodsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentMethodsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentMethodsByIDs] failed checking paymentMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentMethods with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_methods\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentMethodsQueries.deletePaymentMethods + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentMethodsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentMethodsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentMethodsByIDs(ctx context.Context, ids []model.PaymentMethodsPrimaryID) (exists bool, notFoundIds []model.PaymentMethodsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_methods\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentMethodsQueries.selectPaymentMethods, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentMethodsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentMethodsPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentMethodsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentMethodsPrimaryID]bool{}
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

// BulkUpdatePaymentMethods is used to bulk update paymentMethods, by default it will update all field
// if want to update specific field, then fill paymentMethodssMapUpdateFieldsRequest else please fill paymentMethodssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentMethods(ctx context.Context, paymentMethodssMap map[model.PaymentMethodsPrimaryID]*model.PaymentMethods, paymentMethodssMapUpdateFieldsRequest map[model.PaymentMethodsPrimaryID]PaymentMethodsUpdateFieldList) (err error) {
	if len(paymentMethodssMap) == 0 && len(paymentMethodssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentMethodssMapUpdateField map[model.PaymentMethodsPrimaryID]PaymentMethodsUpdateFieldList = map[model.PaymentMethodsPrimaryID]PaymentMethodsUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(paymentMethodssMap) > 0 {
		for id, paymentMethods := range paymentMethodssMap {
			if paymentMethods == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentMethods] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentMethodssMapUpdateField[id] = defaultPaymentMethodsUpdateFields(*paymentMethods)
		}
	} else {
		paymentMethodssMapUpdateField = paymentMethodssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentMethodsQuery(paymentMethodssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentMethodsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentMethods] failed checking paymentMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentMethods with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentMethodsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_methods\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentMethods] failed exec query")
	}
	return
}

type PaymentMethodsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentMethodsFieldParameter(param string, args ...interface{}) PaymentMethodsFieldParameter {
	return PaymentMethodsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentMethodsQuery(mapPaymentMethodss map[model.PaymentMethodsPrimaryID]PaymentMethodsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentMethodsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentMethodsPrimaryID]map[string]interface{}{}
	paymentMethodsSelectFields := NewPaymentMethodsSelectFields()
	for id, updateFields := range mapPaymentMethodss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentMethodsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentMethodss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentMethodsFieldType(updateField.paymentMethodsField)))
			args = append(args, fields[string(updateField.paymentMethodsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentMethodsField))
		if updateField.paymentMethodsField == paymentMethodsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentMethodsField, asTableValues, updateField.paymentMethodsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentMethodsField,
				"\"payment_methods\"", updateField.paymentMethodsField,
				asTableValues, updateField.paymentMethodsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentMethodsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentMethodsPrimaryID, asTableValue string) (whereQry string) {
	paymentMethodsSelectFields := NewPaymentMethodsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_methods\".\"id\" = %s.\"id\"::"+GetPaymentMethodsFieldType(paymentMethodsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentMethodsFieldType(paymentMethodsField PaymentMethodsField) string {
	selectPaymentMethodsFields := NewPaymentMethodsSelectFields()
	switch paymentMethodsField {

	case selectPaymentMethodsFields.Id():
		return "uuid"

	case selectPaymentMethodsFields.UserId():
		return "uuid"

	case selectPaymentMethodsFields.MerchantId():
		return "uuid"

	case selectPaymentMethodsFields.MethodType():
		return "payment_method_type_enum"

	case selectPaymentMethodsFields.Psp():
		return "psp_enum"

	case selectPaymentMethodsFields.TokenRef():
		return "text"

	case selectPaymentMethodsFields.TokenType():
		return "text"

	case selectPaymentMethodsFields.TokenExpiresAt():
		return "timestamptz"

	case selectPaymentMethodsFields.CardBrand():
		return "text"

	case selectPaymentMethodsFields.CardLastFour():
		return "text"

	case selectPaymentMethodsFields.CardExpMonth():
		return "int2"

	case selectPaymentMethodsFields.CardExpYear():
		return "int2"

	case selectPaymentMethodsFields.CardCountry():
		return "text"

	case selectPaymentMethodsFields.CardFundingType():
		return "text"

	case selectPaymentMethodsFields.CardBin():
		return "text"

	case selectPaymentMethodsFields.WalletAccountRef():
		return "text"

	case selectPaymentMethodsFields.VaBankCode():
		return "text"

	case selectPaymentMethodsFields.DisplayLabel():
		return "text"

	case selectPaymentMethodsFields.IsDefault():
		return "bool"

	case selectPaymentMethodsFields.IsActive():
		return "bool"

	case selectPaymentMethodsFields.VerifiedAt():
		return "timestamptz"

	case selectPaymentMethodsFields.Fingerprint():
		return "text"

	case selectPaymentMethodsFields.GdprErasureRequestedAt():
		return "timestamptz"

	case selectPaymentMethodsFields.GdprErasedAt():
		return "timestamptz"

	case selectPaymentMethodsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentMethodsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentMethodsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentMethodsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentMethodsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentMethodsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentMethods(ctx context.Context, paymentMethods *model.PaymentMethods, fieldsInsert ...PaymentMethodsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentMethodsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentMethodsPrimaryID{
		Id: paymentMethods.Id,
	}
	exists, err := repo.IsExistPaymentMethodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentMethods] failed checking paymentMethods whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentMethods", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentMethods([]model.PaymentMethods{*paymentMethods}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentMethodsQueries.insertPaymentMethods, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentMethods] failed exec create paymentMethods query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentMethodsByID(ctx context.Context, primaryID model.PaymentMethodsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentMethodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentMethodsByID] failed checking paymentMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentMethods with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentMethodsCompositePrimaryKeyWhere([]model.PaymentMethodsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentMethodsQueries.deletePaymentMethods + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentMethodsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentMethodsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentMethodsFilterResult, err error) {
	query, args, err := composePaymentMethodsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentMethodsByFilter] failed compose paymentMethods filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentMethodsByFilter] failed get paymentMethods by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentMethodsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentMethodsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultPaymentMethodsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := PaymentMethodsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, PaymentMethodsField(filterSelectField))
		}
		selectFields = composePaymentMethodsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(paymentMethodsQueries.selectPaymentMethods, selectFields)

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

func (repo *RepositoryImpl) IsExistPaymentMethodsByID(ctx context.Context, primaryID model.PaymentMethodsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentMethodsCompositePrimaryKeyWhere([]model.PaymentMethodsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentMethodsQueries.selectCountPaymentMethods, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentMethodsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentMethods(ctx context.Context, selectFields ...PaymentMethodsField) (paymentMethodsList model.PaymentMethodsList, err error) {
	var (
		defaultPaymentMethodsSelectFields = defaultPaymentMethodsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentMethodsSelectFields = composePaymentMethodsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentMethodsQueries.selectPaymentMethods, defaultPaymentMethodsSelectFields)

	err = repo.db.Read.Select(&paymentMethodsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentMethods] failed get paymentMethods list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentMethodsByID(ctx context.Context, primaryID model.PaymentMethodsPrimaryID, selectFields ...PaymentMethodsField) (paymentMethods model.PaymentMethods, err error) {
	var (
		defaultPaymentMethodsSelectFields = defaultPaymentMethodsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentMethodsSelectFields = composePaymentMethodsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentMethodsCompositePrimaryKeyWhere([]model.PaymentMethodsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentMethodsQueries.selectPaymentMethods+" WHERE "+whereQry, defaultPaymentMethodsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&paymentMethods, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentMethods with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentMethodsByID] failed get paymentMethods")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentMethodsByID(ctx context.Context, primaryID model.PaymentMethodsPrimaryID, paymentMethods *model.PaymentMethods, paymentMethodsUpdateFields ...PaymentMethodsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentMethodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentMethods] failed checking paymentMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentMethods with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentMethods == nil {
		if len(paymentMethodsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentMethodsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentMethods = &model.PaymentMethods{}
	}
	var (
		defaultPaymentMethodsUpdateFields = defaultPaymentMethodsUpdateFields(*paymentMethods)
		tempUpdateField                   PaymentMethodsUpdateFieldList
		selectFields                      = NewPaymentMethodsSelectFields()
	)
	if len(paymentMethodsUpdateFields) > 0 {
		for _, updateField := range paymentMethodsUpdateFields {
			if updateField.paymentMethodsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentMethodsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentMethodsCompositePrimaryKeyWhere([]model.PaymentMethodsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentMethodsCommand(defaultPaymentMethodsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentMethodsQueries.updatePaymentMethods+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentMethods] error when try to update paymentMethods by id")
	}
	return err
}

var (
	paymentMethodsQueries = struct {
		selectPaymentMethods      string
		selectCountPaymentMethods string
		deletePaymentMethods      string
		updatePaymentMethods      string
		insertPaymentMethods      string
	}{
		selectPaymentMethods:      "SELECT %s FROM \"payment_methods\"",
		selectCountPaymentMethods: "SELECT COUNT(\"id\") FROM \"payment_methods\"",
		deletePaymentMethods:      "DELETE FROM \"payment_methods\"",
		updatePaymentMethods:      "UPDATE \"payment_methods\" SET %s ",
		insertPaymentMethods:      "INSERT INTO \"payment_methods\" %s VALUES %s",
	}
)

type PaymentMethodsRepository interface {
	CreatePaymentMethods(ctx context.Context, paymentMethods *model.PaymentMethods, fieldsInsert ...PaymentMethodsField) error
	BulkCreatePaymentMethods(ctx context.Context, paymentMethodsList []*model.PaymentMethods, fieldsInsert ...PaymentMethodsField) error
	ResolvePaymentMethods(ctx context.Context, selectFields ...PaymentMethodsField) (model.PaymentMethodsList, error)
	ResolvePaymentMethodsByID(ctx context.Context, primaryID model.PaymentMethodsPrimaryID, selectFields ...PaymentMethodsField) (model.PaymentMethods, error)
	UpdatePaymentMethodsByID(ctx context.Context, id model.PaymentMethodsPrimaryID, paymentMethods *model.PaymentMethods, paymentMethodsUpdateFields ...PaymentMethodsUpdateField) error
	BulkUpdatePaymentMethods(ctx context.Context, paymentMethodsListMap map[model.PaymentMethodsPrimaryID]*model.PaymentMethods, PaymentMethodssMapUpdateFieldsRequest map[model.PaymentMethodsPrimaryID]PaymentMethodsUpdateFieldList) (err error)
	DeletePaymentMethodsByID(ctx context.Context, id model.PaymentMethodsPrimaryID) error
	BulkDeletePaymentMethodsByIDs(ctx context.Context, ids []model.PaymentMethodsPrimaryID) error
	ResolvePaymentMethodsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentMethodsFilterResult, err error)
	IsExistPaymentMethodsByIDs(ctx context.Context, ids []model.PaymentMethodsPrimaryID) (exists bool, notFoundIds []model.PaymentMethodsPrimaryID, err error)
	IsExistPaymentMethodsByID(ctx context.Context, id model.PaymentMethodsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
