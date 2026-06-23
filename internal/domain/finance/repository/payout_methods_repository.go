package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsPayoutMethods(payoutMethodsList []model.PayoutMethods, fieldsInsert ...PayoutMethodsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPayoutMethodsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, payoutMethods := range payoutMethodsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, payoutMethods.Id)
			case selectField.MerchantPartyId():
				args = append(args, payoutMethods.MerchantPartyId)
			case selectField.MethodType():
				args = append(args, payoutMethods.MethodType)
			case selectField.ProviderCode():
				args = append(args, payoutMethods.ProviderCode)
			case selectField.CountryCode():
				args = append(args, payoutMethods.CountryCode)
			case selectField.CurrencyCode():
				args = append(args, payoutMethods.CurrencyCode)
			case selectField.AccountHolderName():
				args = append(args, payoutMethods.AccountHolderName)
			case selectField.AccountNoLast4():
				args = append(args, payoutMethods.AccountNoLast4)
			case selectField.BankCode():
				args = append(args, payoutMethods.BankCode)
			case selectField.BankName():
				args = append(args, payoutMethods.BankName)
			case selectField.AccountTokenRef():
				args = append(args, payoutMethods.AccountTokenRef)
			case selectField.AccountEncryptedBlob():
				args = append(args, payoutMethods.AccountEncryptedBlob)
			case selectField.VerificationStatus():
				args = append(args, payoutMethods.VerificationStatus)
			case selectField.IsDefault():
				args = append(args, payoutMethods.IsDefault)
			case selectField.MethodStatus():
				args = append(args, payoutMethods.MethodStatus)
			case selectField.Metadata():
				args = append(args, payoutMethods.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, payoutMethods.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, payoutMethods.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, payoutMethods.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, payoutMethods.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, payoutMethods.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, payoutMethods.MetaDeletedBy)

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

func composePayoutMethodsCompositePrimaryKeyWhere(primaryIDs []model.PayoutMethodsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payout_methods\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPayoutMethodsSelectFields() string {
	fields := NewPayoutMethodsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePayoutMethodsSelectFields(selectFields ...PayoutMethodsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PayoutMethodsField string
type PayoutMethodsFieldList []PayoutMethodsField

type PayoutMethodsSelectFields struct {
}

func (ss PayoutMethodsSelectFields) Id() PayoutMethodsField {
	return PayoutMethodsField("id")
}

func (ss PayoutMethodsSelectFields) MerchantPartyId() PayoutMethodsField {
	return PayoutMethodsField("merchant_party_id")
}

func (ss PayoutMethodsSelectFields) MethodType() PayoutMethodsField {
	return PayoutMethodsField("method_type")
}

func (ss PayoutMethodsSelectFields) ProviderCode() PayoutMethodsField {
	return PayoutMethodsField("provider_code")
}

func (ss PayoutMethodsSelectFields) CountryCode() PayoutMethodsField {
	return PayoutMethodsField("country_code")
}

func (ss PayoutMethodsSelectFields) CurrencyCode() PayoutMethodsField {
	return PayoutMethodsField("currency_code")
}

func (ss PayoutMethodsSelectFields) AccountHolderName() PayoutMethodsField {
	return PayoutMethodsField("account_holder_name")
}

func (ss PayoutMethodsSelectFields) AccountNoLast4() PayoutMethodsField {
	return PayoutMethodsField("account_no_last4")
}

func (ss PayoutMethodsSelectFields) BankCode() PayoutMethodsField {
	return PayoutMethodsField("bank_code")
}

func (ss PayoutMethodsSelectFields) BankName() PayoutMethodsField {
	return PayoutMethodsField("bank_name")
}

func (ss PayoutMethodsSelectFields) AccountTokenRef() PayoutMethodsField {
	return PayoutMethodsField("account_token_ref")
}

func (ss PayoutMethodsSelectFields) AccountEncryptedBlob() PayoutMethodsField {
	return PayoutMethodsField("account_encrypted_blob")
}

func (ss PayoutMethodsSelectFields) VerificationStatus() PayoutMethodsField {
	return PayoutMethodsField("verification_status")
}

func (ss PayoutMethodsSelectFields) IsDefault() PayoutMethodsField {
	return PayoutMethodsField("is_default")
}

func (ss PayoutMethodsSelectFields) MethodStatus() PayoutMethodsField {
	return PayoutMethodsField("method_status")
}

func (ss PayoutMethodsSelectFields) Metadata() PayoutMethodsField {
	return PayoutMethodsField("metadata")
}

func (ss PayoutMethodsSelectFields) MetaCreatedAt() PayoutMethodsField {
	return PayoutMethodsField("meta_created_at")
}

func (ss PayoutMethodsSelectFields) MetaCreatedBy() PayoutMethodsField {
	return PayoutMethodsField("meta_created_by")
}

func (ss PayoutMethodsSelectFields) MetaUpdatedAt() PayoutMethodsField {
	return PayoutMethodsField("meta_updated_at")
}

func (ss PayoutMethodsSelectFields) MetaUpdatedBy() PayoutMethodsField {
	return PayoutMethodsField("meta_updated_by")
}

func (ss PayoutMethodsSelectFields) MetaDeletedAt() PayoutMethodsField {
	return PayoutMethodsField("meta_deleted_at")
}

func (ss PayoutMethodsSelectFields) MetaDeletedBy() PayoutMethodsField {
	return PayoutMethodsField("meta_deleted_by")
}

func (ss PayoutMethodsSelectFields) All() PayoutMethodsFieldList {
	return []PayoutMethodsField{
		ss.Id(),
		ss.MerchantPartyId(),
		ss.MethodType(),
		ss.ProviderCode(),
		ss.CountryCode(),
		ss.CurrencyCode(),
		ss.AccountHolderName(),
		ss.AccountNoLast4(),
		ss.BankCode(),
		ss.BankName(),
		ss.AccountTokenRef(),
		ss.AccountEncryptedBlob(),
		ss.VerificationStatus(),
		ss.IsDefault(),
		ss.MethodStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPayoutMethodsSelectFields() PayoutMethodsSelectFields {
	return PayoutMethodsSelectFields{}
}

type PayoutMethodsUpdateFieldOption struct {
	useIncrement bool
}
type PayoutMethodsUpdateField struct {
	payoutMethodsField PayoutMethodsField
	opt                PayoutMethodsUpdateFieldOption
	value              interface{}
}
type PayoutMethodsUpdateFieldList []PayoutMethodsUpdateField

func defaultPayoutMethodsUpdateFieldOption() PayoutMethodsUpdateFieldOption {
	return PayoutMethodsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPayoutMethodsOption(useIncrement bool) func(*PayoutMethodsUpdateFieldOption) {
	return func(pcufo *PayoutMethodsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPayoutMethodsUpdateField(field PayoutMethodsField, val interface{}, opts ...func(*PayoutMethodsUpdateFieldOption)) PayoutMethodsUpdateField {
	defaultOpt := defaultPayoutMethodsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PayoutMethodsUpdateField{
		payoutMethodsField: field,
		value:              val,
		opt:                defaultOpt,
	}
}
func defaultPayoutMethodsUpdateFields(payoutMethods model.PayoutMethods) (payoutMethodsUpdateFieldList PayoutMethodsUpdateFieldList) {
	selectFields := NewPayoutMethodsSelectFields()
	payoutMethodsUpdateFieldList = append(payoutMethodsUpdateFieldList,
		NewPayoutMethodsUpdateField(selectFields.Id(), payoutMethods.Id),
		NewPayoutMethodsUpdateField(selectFields.MerchantPartyId(), payoutMethods.MerchantPartyId),
		NewPayoutMethodsUpdateField(selectFields.MethodType(), payoutMethods.MethodType),
		NewPayoutMethodsUpdateField(selectFields.ProviderCode(), payoutMethods.ProviderCode),
		NewPayoutMethodsUpdateField(selectFields.CountryCode(), payoutMethods.CountryCode),
		NewPayoutMethodsUpdateField(selectFields.CurrencyCode(), payoutMethods.CurrencyCode),
		NewPayoutMethodsUpdateField(selectFields.AccountHolderName(), payoutMethods.AccountHolderName),
		NewPayoutMethodsUpdateField(selectFields.AccountNoLast4(), payoutMethods.AccountNoLast4),
		NewPayoutMethodsUpdateField(selectFields.BankCode(), payoutMethods.BankCode),
		NewPayoutMethodsUpdateField(selectFields.BankName(), payoutMethods.BankName),
		NewPayoutMethodsUpdateField(selectFields.AccountTokenRef(), payoutMethods.AccountTokenRef),
		NewPayoutMethodsUpdateField(selectFields.AccountEncryptedBlob(), payoutMethods.AccountEncryptedBlob),
		NewPayoutMethodsUpdateField(selectFields.VerificationStatus(), payoutMethods.VerificationStatus),
		NewPayoutMethodsUpdateField(selectFields.IsDefault(), payoutMethods.IsDefault),
		NewPayoutMethodsUpdateField(selectFields.MethodStatus(), payoutMethods.MethodStatus),
		NewPayoutMethodsUpdateField(selectFields.Metadata(), payoutMethods.Metadata),
		NewPayoutMethodsUpdateField(selectFields.MetaCreatedAt(), payoutMethods.MetaCreatedAt),
		NewPayoutMethodsUpdateField(selectFields.MetaCreatedBy(), payoutMethods.MetaCreatedBy),
		NewPayoutMethodsUpdateField(selectFields.MetaUpdatedAt(), payoutMethods.MetaUpdatedAt),
		NewPayoutMethodsUpdateField(selectFields.MetaUpdatedBy(), payoutMethods.MetaUpdatedBy),
		NewPayoutMethodsUpdateField(selectFields.MetaDeletedAt(), payoutMethods.MetaDeletedAt),
		NewPayoutMethodsUpdateField(selectFields.MetaDeletedBy(), payoutMethods.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPayoutMethodsCommand(payoutMethodsUpdateFieldList PayoutMethodsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range payoutMethodsUpdateFieldList {
		field := string(updateField.payoutMethodsField)
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

func (repo *RepositoryImpl) BulkCreatePayoutMethods(ctx context.Context, payoutMethodsList []*model.PayoutMethods, fieldsInsert ...PayoutMethodsField) (err error) {
	var (
		fieldsStr              string
		valueListStr           []string
		argsList               []interface{}
		primaryIds             []model.PayoutMethodsPrimaryID
		payoutMethodsValueList []model.PayoutMethods
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPayoutMethodsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, payoutMethods := range payoutMethodsList {

		primaryIds = append(primaryIds, payoutMethods.ToPayoutMethodsPrimaryID())

		payoutMethodsValueList = append(payoutMethodsValueList, *payoutMethods)
	}

	_, notFoundIds, err := repo.IsExistPayoutMethodsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutMethods] failed checking payoutMethods whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PayoutMethodsPrimaryID{}
		mapNotFoundIds := map[model.PayoutMethodsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "payoutMethods", fmt.Sprintf("payoutMethods with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPayoutMethods(payoutMethodsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(payoutMethodsQueries.insertPayoutMethods, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutMethods] failed exec create payoutMethods query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePayoutMethodsByIDs(ctx context.Context, primaryIDs []model.PayoutMethodsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPayoutMethodsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutMethodsByIDs] failed checking payoutMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutMethods with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_methods\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(payoutMethodsQueries.deletePayoutMethods + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutMethodsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutMethodsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPayoutMethodsByIDs(ctx context.Context, ids []model.PayoutMethodsPrimaryID) (exists bool, notFoundIds []model.PayoutMethodsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_methods\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(payoutMethodsQueries.selectPayoutMethods, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutMethodsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PayoutMethodsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutMethodsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PayoutMethodsPrimaryID]bool{}
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

// BulkUpdatePayoutMethods is used to bulk update payoutMethods, by default it will update all field
// if want to update specific field, then fill payoutMethodssMapUpdateFieldsRequest else please fill payoutMethodssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePayoutMethods(ctx context.Context, payoutMethodssMap map[model.PayoutMethodsPrimaryID]*model.PayoutMethods, payoutMethodssMapUpdateFieldsRequest map[model.PayoutMethodsPrimaryID]PayoutMethodsUpdateFieldList) (err error) {
	if len(payoutMethodssMap) == 0 && len(payoutMethodssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		payoutMethodssMapUpdateField map[model.PayoutMethodsPrimaryID]PayoutMethodsUpdateFieldList = map[model.PayoutMethodsPrimaryID]PayoutMethodsUpdateFieldList{}
		asTableValues                string                                                        = "myvalues"
	)

	if len(payoutMethodssMap) > 0 {
		for id, payoutMethods := range payoutMethodssMap {
			if payoutMethods == nil {
				log.Error().Err(err).Msg("[BulkUpdatePayoutMethods] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			payoutMethodssMapUpdateField[id] = defaultPayoutMethodsUpdateFields(*payoutMethods)
		}
	} else {
		payoutMethodssMapUpdateField = payoutMethodssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePayoutMethodsQuery(payoutMethodssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPayoutMethodsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutMethods] failed checking payoutMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutMethods with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePayoutMethodsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payout_methods\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutMethods] failed exec query")
	}
	return
}

type PayoutMethodsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPayoutMethodsFieldParameter(param string, args ...interface{}) PayoutMethodsFieldParameter {
	return PayoutMethodsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePayoutMethodsQuery(mapPayoutMethodss map[model.PayoutMethodsPrimaryID]PayoutMethodsUpdateFieldList, asTableValues string) (primaryIDs []model.PayoutMethodsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PayoutMethodsPrimaryID]map[string]interface{}{}
	payoutMethodsSelectFields := NewPayoutMethodsSelectFields()
	for id, updateFields := range mapPayoutMethodss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.payoutMethodsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPayoutMethodss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPayoutMethodsFieldType(updateField.payoutMethodsField)))
			args = append(args, fields[string(updateField.payoutMethodsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.payoutMethodsField))
		if updateField.payoutMethodsField == payoutMethodsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.payoutMethodsField, asTableValues, updateField.payoutMethodsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.payoutMethodsField,
				"\"payout_methods\"", updateField.payoutMethodsField,
				asTableValues, updateField.payoutMethodsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePayoutMethodsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PayoutMethodsPrimaryID, asTableValue string) (whereQry string) {
	payoutMethodsSelectFields := NewPayoutMethodsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payout_methods\".\"id\" = %s.\"id\"::"+GetPayoutMethodsFieldType(payoutMethodsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPayoutMethodsFieldType(payoutMethodsField PayoutMethodsField) string {
	selectPayoutMethodsFields := NewPayoutMethodsSelectFields()
	switch payoutMethodsField {

	case selectPayoutMethodsFields.Id():
		return "uuid"

	case selectPayoutMethodsFields.MerchantPartyId():
		return "uuid"

	case selectPayoutMethodsFields.MethodType():
		return "method_type_enum"

	case selectPayoutMethodsFields.ProviderCode():
		return "text"

	case selectPayoutMethodsFields.CountryCode():
		return "text"

	case selectPayoutMethodsFields.CurrencyCode():
		return "text"

	case selectPayoutMethodsFields.AccountHolderName():
		return "text"

	case selectPayoutMethodsFields.AccountNoLast4():
		return "text"

	case selectPayoutMethodsFields.BankCode():
		return "text"

	case selectPayoutMethodsFields.BankName():
		return "text"

	case selectPayoutMethodsFields.AccountTokenRef():
		return "text"

	case selectPayoutMethodsFields.AccountEncryptedBlob():
		return "text"

	case selectPayoutMethodsFields.VerificationStatus():
		return "verification_status_enum"

	case selectPayoutMethodsFields.IsDefault():
		return "bool"

	case selectPayoutMethodsFields.MethodStatus():
		return "method_status_enum"

	case selectPayoutMethodsFields.Metadata():
		return "jsonb"

	case selectPayoutMethodsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPayoutMethodsFields.MetaCreatedBy():
		return "uuid"

	case selectPayoutMethodsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPayoutMethodsFields.MetaUpdatedBy():
		return "uuid"

	case selectPayoutMethodsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPayoutMethodsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePayoutMethods(ctx context.Context, payoutMethods *model.PayoutMethods, fieldsInsert ...PayoutMethodsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPayoutMethodsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PayoutMethodsPrimaryID{
		Id: payoutMethods.Id,
	}
	exists, err := repo.IsExistPayoutMethodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutMethods] failed checking payoutMethods whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "payoutMethods", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPayoutMethods([]model.PayoutMethods{*payoutMethods}, fieldsInsert...)
	commandQuery := fmt.Sprintf(payoutMethodsQueries.insertPayoutMethods, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutMethods] failed exec create payoutMethods query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePayoutMethodsByID(ctx context.Context, primaryID model.PayoutMethodsPrimaryID) (err error) {
	exists, err := repo.IsExistPayoutMethodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutMethodsByID] failed checking payoutMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutMethods with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePayoutMethodsCompositePrimaryKeyWhere([]model.PayoutMethodsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(payoutMethodsQueries.deletePayoutMethods + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutMethodsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutMethodsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutMethodsFilterResult, err error) {
	query, args, err := composePayoutMethodsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutMethodsByFilter] failed compose payoutMethods filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutMethodsByFilter] failed get payoutMethods by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePayoutMethodsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PayoutMethodsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePayoutMethodsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePayoutMethodsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePayoutMethodsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 22 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPayoutMethodsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 22+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["method_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_type\"")
			selectedColumns["method_type"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_code\"")
			selectedColumns["provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["country_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"country_code\"")
			selectedColumns["country_code"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["account_holder_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_holder_name\"")
			selectedColumns["account_holder_name"] = struct{}{}
		}
		if _, selected := selectedColumns["account_no_last4"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_no_last4\"")
			selectedColumns["account_no_last4"] = struct{}{}
		}
		if _, selected := selectedColumns["bank_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"bank_code\"")
			selectedColumns["bank_code"] = struct{}{}
		}
		if _, selected := selectedColumns["bank_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"bank_name\"")
			selectedColumns["bank_name"] = struct{}{}
		}
		if _, selected := selectedColumns["account_token_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_token_ref\"")
			selectedColumns["account_token_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["account_encrypted_blob"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_encrypted_blob\"")
			selectedColumns["account_encrypted_blob"] = struct{}{}
		}
		if _, selected := selectedColumns["verification_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"verification_status\"")
			selectedColumns["verification_status"] = struct{}{}
		}
		if _, selected := selectedColumns["is_default"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_default\"")
			selectedColumns["is_default"] = struct{}{}
		}
		if _, selected := selectedColumns["method_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_status\"")
			selectedColumns["method_status"] = struct{}{}
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

type payoutMethodsFilterPlaceholder struct {
	index int
}

func (p *payoutMethodsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePayoutMethodsFilterPredicate(filterField model.FilterField, placeholders *payoutMethodsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPayoutMethodsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePayoutMethodsFilterSQLExpr(spec)
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

func composePayoutMethodsFilterGroup(group model.FilterGroup, placeholders *payoutMethodsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePayoutMethodsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePayoutMethodsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePayoutMethodsFilterWhereQueries(filter model.Filter, placeholders *payoutMethodsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePayoutMethodsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePayoutMethodsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePayoutMethodsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePayoutMethodsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePayoutMethodsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePayoutMethodsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := payoutMethodsFilterPlaceholder{index: 1}
	whereQueries, err := composePayoutMethodsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPayoutMethodsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePayoutMethodsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePayoutMethodsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payout_methods\" base%s", strings.Join(selectColumns, ","), composePayoutMethodsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPayoutMethodsByID(ctx context.Context, primaryID model.PayoutMethodsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePayoutMethodsCompositePrimaryKeyWhere([]model.PayoutMethodsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", payoutMethodsQueries.selectCountPayoutMethods, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutMethodsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutMethods(ctx context.Context, selectFields ...PayoutMethodsField) (payoutMethodsList model.PayoutMethodsList, err error) {
	var (
		defaultPayoutMethodsSelectFields = defaultPayoutMethodsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutMethodsSelectFields = composePayoutMethodsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(payoutMethodsQueries.selectPayoutMethods, defaultPayoutMethodsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &payoutMethodsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutMethods] failed get payoutMethods list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutMethodsByID(ctx context.Context, primaryID model.PayoutMethodsPrimaryID, selectFields ...PayoutMethodsField) (payoutMethods model.PayoutMethods, err error) {
	var (
		defaultPayoutMethodsSelectFields = defaultPayoutMethodsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutMethodsSelectFields = composePayoutMethodsSelectFields(selectFields...)
	}
	whereQry, params := composePayoutMethodsCompositePrimaryKeyWhere([]model.PayoutMethodsPrimaryID{primaryID})
	query := fmt.Sprintf(payoutMethodsQueries.selectPayoutMethods+" WHERE "+whereQry, defaultPayoutMethodsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &payoutMethods, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("payoutMethods with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePayoutMethodsByID] failed get payoutMethods")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePayoutMethodsByID(ctx context.Context, primaryID model.PayoutMethodsPrimaryID, payoutMethods *model.PayoutMethods, payoutMethodsUpdateFields ...PayoutMethodsUpdateField) (err error) {
	exists, err := repo.IsExistPayoutMethodsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutMethods] failed checking payoutMethods whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutMethods with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if payoutMethods == nil {
		if len(payoutMethodsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePayoutMethodsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		payoutMethods = &model.PayoutMethods{}
	}
	var (
		defaultPayoutMethodsUpdateFields = defaultPayoutMethodsUpdateFields(*payoutMethods)
		tempUpdateField                  PayoutMethodsUpdateFieldList
		selectFields                     = NewPayoutMethodsSelectFields()
	)
	if len(payoutMethodsUpdateFields) > 0 {
		for _, updateField := range payoutMethodsUpdateFields {
			if updateField.payoutMethodsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPayoutMethodsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePayoutMethodsCompositePrimaryKeyWhere([]model.PayoutMethodsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPayoutMethodsCommand(defaultPayoutMethodsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(payoutMethodsQueries.updatePayoutMethods+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutMethods] error when try to update payoutMethods by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePayoutMethodsByFilter(ctx context.Context, filter model.Filter, payoutMethodsUpdateFields ...PayoutMethodsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(payoutMethodsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PayoutMethodsUpdateFieldList
		selectFields = NewPayoutMethodsSelectFields()
	)
	for _, updateField := range payoutMethodsUpdateFields {
		if updateField.payoutMethodsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPayoutMethodsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := payoutMethodsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePayoutMethodsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payout_methods\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutMethodsByFilter] error when try to update payoutMethods by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutMethodsByFilter] failed get rows affected")
	}
	return
}

var (
	payoutMethodsQueries = struct {
		selectPayoutMethods      string
		selectCountPayoutMethods string
		deletePayoutMethods      string
		updatePayoutMethods      string
		insertPayoutMethods      string
	}{
		selectPayoutMethods:      "SELECT %s FROM \"payout_methods\"",
		selectCountPayoutMethods: "SELECT COUNT(\"id\") FROM \"payout_methods\"",
		deletePayoutMethods:      "DELETE FROM \"payout_methods\"",
		updatePayoutMethods:      "UPDATE \"payout_methods\" SET %s ",
		insertPayoutMethods:      "INSERT INTO \"payout_methods\" %s VALUES %s",
	}
)

type PayoutMethodsRepository interface {
	CreatePayoutMethods(ctx context.Context, payoutMethods *model.PayoutMethods, fieldsInsert ...PayoutMethodsField) error
	BulkCreatePayoutMethods(ctx context.Context, payoutMethodsList []*model.PayoutMethods, fieldsInsert ...PayoutMethodsField) error
	ResolvePayoutMethods(ctx context.Context, selectFields ...PayoutMethodsField) (model.PayoutMethodsList, error)
	ResolvePayoutMethodsByID(ctx context.Context, primaryID model.PayoutMethodsPrimaryID, selectFields ...PayoutMethodsField) (model.PayoutMethods, error)
	UpdatePayoutMethodsByID(ctx context.Context, id model.PayoutMethodsPrimaryID, payoutMethods *model.PayoutMethods, payoutMethodsUpdateFields ...PayoutMethodsUpdateField) error
	UpdatePayoutMethodsByFilter(ctx context.Context, filter model.Filter, payoutMethodsUpdateFields ...PayoutMethodsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePayoutMethods(ctx context.Context, payoutMethodsListMap map[model.PayoutMethodsPrimaryID]*model.PayoutMethods, PayoutMethodssMapUpdateFieldsRequest map[model.PayoutMethodsPrimaryID]PayoutMethodsUpdateFieldList) (err error)
	DeletePayoutMethodsByID(ctx context.Context, id model.PayoutMethodsPrimaryID) error
	BulkDeletePayoutMethodsByIDs(ctx context.Context, ids []model.PayoutMethodsPrimaryID) error
	ResolvePayoutMethodsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutMethodsFilterResult, err error)
	IsExistPayoutMethodsByIDs(ctx context.Context, ids []model.PayoutMethodsPrimaryID) (exists bool, notFoundIds []model.PayoutMethodsPrimaryID, err error)
	IsExistPayoutMethodsByID(ctx context.Context, id model.PayoutMethodsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
