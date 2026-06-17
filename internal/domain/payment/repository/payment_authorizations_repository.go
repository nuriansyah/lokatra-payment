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

func composeInsertFieldsAndParamsPaymentAuthorizations(paymentAuthorizationsList []model.PaymentAuthorizations, fieldsInsert ...PaymentAuthorizationsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentAuthorizationsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentAuthorizations := range paymentAuthorizationsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentAuthorizations.Id)
			case selectField.PaymentIntentId():
				args = append(args, paymentAuthorizations.PaymentIntentId)
			case selectField.PaymentAttemptId():
				args = append(args, paymentAuthorizations.PaymentAttemptId)
			case selectField.ProviderAccountId():
				args = append(args, paymentAuthorizations.ProviderAccountId)
			case selectField.ProviderAuthorizationId():
				args = append(args, paymentAuthorizations.ProviderAuthorizationId)
			case selectField.Amount():
				args = append(args, paymentAuthorizations.Amount)
			case selectField.Currency():
				args = append(args, paymentAuthorizations.Currency)
			case selectField.Status():
				args = append(args, paymentAuthorizations.Status)
			case selectField.AuthorizedAt():
				args = append(args, paymentAuthorizations.AuthorizedAt)
			case selectField.ExpiresAt():
				args = append(args, paymentAuthorizations.ExpiresAt)
			case selectField.CapturedAmount():
				args = append(args, paymentAuthorizations.CapturedAmount)
			case selectField.FailureCode():
				args = append(args, paymentAuthorizations.FailureCode)
			case selectField.FailureMessage():
				args = append(args, paymentAuthorizations.FailureMessage)
			case selectField.RawRequest():
				args = append(args, paymentAuthorizations.RawRequest)
			case selectField.RawResponse():
				args = append(args, paymentAuthorizations.RawResponse)
			case selectField.Metadata():
				args = append(args, paymentAuthorizations.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentAuthorizations.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentAuthorizations.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentAuthorizations.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentAuthorizations.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentAuthorizations.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentAuthorizations.MetaDeletedBy)

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

func composePaymentAuthorizationsCompositePrimaryKeyWhere(primaryIDs []model.PaymentAuthorizationsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_authorizations\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentAuthorizationsSelectFields() string {
	fields := NewPaymentAuthorizationsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentAuthorizationsSelectFields(selectFields ...PaymentAuthorizationsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentAuthorizationsField string
type PaymentAuthorizationsFieldList []PaymentAuthorizationsField

type PaymentAuthorizationsSelectFields struct {
}

func (ss PaymentAuthorizationsSelectFields) Id() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("id")
}

func (ss PaymentAuthorizationsSelectFields) PaymentIntentId() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("payment_intent_id")
}

func (ss PaymentAuthorizationsSelectFields) PaymentAttemptId() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("payment_attempt_id")
}

func (ss PaymentAuthorizationsSelectFields) ProviderAccountId() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("provider_account_id")
}

func (ss PaymentAuthorizationsSelectFields) ProviderAuthorizationId() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("provider_authorization_id")
}

func (ss PaymentAuthorizationsSelectFields) Amount() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("amount")
}

func (ss PaymentAuthorizationsSelectFields) Currency() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("currency")
}

func (ss PaymentAuthorizationsSelectFields) Status() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("status")
}

func (ss PaymentAuthorizationsSelectFields) AuthorizedAt() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("authorized_at")
}

func (ss PaymentAuthorizationsSelectFields) ExpiresAt() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("expires_at")
}

func (ss PaymentAuthorizationsSelectFields) CapturedAmount() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("captured_amount")
}

func (ss PaymentAuthorizationsSelectFields) FailureCode() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("failure_code")
}

func (ss PaymentAuthorizationsSelectFields) FailureMessage() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("failure_message")
}

func (ss PaymentAuthorizationsSelectFields) RawRequest() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("raw_request")
}

func (ss PaymentAuthorizationsSelectFields) RawResponse() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("raw_response")
}

func (ss PaymentAuthorizationsSelectFields) Metadata() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("metadata")
}

func (ss PaymentAuthorizationsSelectFields) MetaCreatedAt() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("meta_created_at")
}

func (ss PaymentAuthorizationsSelectFields) MetaCreatedBy() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("meta_created_by")
}

func (ss PaymentAuthorizationsSelectFields) MetaUpdatedAt() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("meta_updated_at")
}

func (ss PaymentAuthorizationsSelectFields) MetaUpdatedBy() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("meta_updated_by")
}

func (ss PaymentAuthorizationsSelectFields) MetaDeletedAt() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("meta_deleted_at")
}

func (ss PaymentAuthorizationsSelectFields) MetaDeletedBy() PaymentAuthorizationsField {
	return PaymentAuthorizationsField("meta_deleted_by")
}

func (ss PaymentAuthorizationsSelectFields) All() PaymentAuthorizationsFieldList {
	return []PaymentAuthorizationsField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.PaymentAttemptId(),
		ss.ProviderAccountId(),
		ss.ProviderAuthorizationId(),
		ss.Amount(),
		ss.Currency(),
		ss.Status(),
		ss.AuthorizedAt(),
		ss.ExpiresAt(),
		ss.CapturedAmount(),
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

func NewPaymentAuthorizationsSelectFields() PaymentAuthorizationsSelectFields {
	return PaymentAuthorizationsSelectFields{}
}

type PaymentAuthorizationsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentAuthorizationsUpdateField struct {
	paymentAuthorizationsField PaymentAuthorizationsField
	opt                        PaymentAuthorizationsUpdateFieldOption
	value                      interface{}
}
type PaymentAuthorizationsUpdateFieldList []PaymentAuthorizationsUpdateField

func defaultPaymentAuthorizationsUpdateFieldOption() PaymentAuthorizationsUpdateFieldOption {
	return PaymentAuthorizationsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentAuthorizationsOption(useIncrement bool) func(*PaymentAuthorizationsUpdateFieldOption) {
	return func(pcufo *PaymentAuthorizationsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentAuthorizationsUpdateField(field PaymentAuthorizationsField, val interface{}, opts ...func(*PaymentAuthorizationsUpdateFieldOption)) PaymentAuthorizationsUpdateField {
	defaultOpt := defaultPaymentAuthorizationsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentAuthorizationsUpdateField{
		paymentAuthorizationsField: field,
		value:                      val,
		opt:                        defaultOpt,
	}
}
func defaultPaymentAuthorizationsUpdateFields(paymentAuthorizations model.PaymentAuthorizations) (paymentAuthorizationsUpdateFieldList PaymentAuthorizationsUpdateFieldList) {
	selectFields := NewPaymentAuthorizationsSelectFields()
	paymentAuthorizationsUpdateFieldList = append(paymentAuthorizationsUpdateFieldList,
		NewPaymentAuthorizationsUpdateField(selectFields.Id(), paymentAuthorizations.Id),
		NewPaymentAuthorizationsUpdateField(selectFields.PaymentIntentId(), paymentAuthorizations.PaymentIntentId),
		NewPaymentAuthorizationsUpdateField(selectFields.PaymentAttemptId(), paymentAuthorizations.PaymentAttemptId),
		NewPaymentAuthorizationsUpdateField(selectFields.ProviderAccountId(), paymentAuthorizations.ProviderAccountId),
		NewPaymentAuthorizationsUpdateField(selectFields.ProviderAuthorizationId(), paymentAuthorizations.ProviderAuthorizationId),
		NewPaymentAuthorizationsUpdateField(selectFields.Amount(), paymentAuthorizations.Amount),
		NewPaymentAuthorizationsUpdateField(selectFields.Currency(), paymentAuthorizations.Currency),
		NewPaymentAuthorizationsUpdateField(selectFields.Status(), paymentAuthorizations.Status),
		NewPaymentAuthorizationsUpdateField(selectFields.AuthorizedAt(), paymentAuthorizations.AuthorizedAt),
		NewPaymentAuthorizationsUpdateField(selectFields.ExpiresAt(), paymentAuthorizations.ExpiresAt),
		NewPaymentAuthorizationsUpdateField(selectFields.CapturedAmount(), paymentAuthorizations.CapturedAmount),
		NewPaymentAuthorizationsUpdateField(selectFields.FailureCode(), paymentAuthorizations.FailureCode),
		NewPaymentAuthorizationsUpdateField(selectFields.FailureMessage(), paymentAuthorizations.FailureMessage),
		NewPaymentAuthorizationsUpdateField(selectFields.RawRequest(), paymentAuthorizations.RawRequest),
		NewPaymentAuthorizationsUpdateField(selectFields.RawResponse(), paymentAuthorizations.RawResponse),
		NewPaymentAuthorizationsUpdateField(selectFields.Metadata(), paymentAuthorizations.Metadata),
		NewPaymentAuthorizationsUpdateField(selectFields.MetaCreatedAt(), paymentAuthorizations.MetaCreatedAt),
		NewPaymentAuthorizationsUpdateField(selectFields.MetaCreatedBy(), paymentAuthorizations.MetaCreatedBy),
		NewPaymentAuthorizationsUpdateField(selectFields.MetaUpdatedAt(), paymentAuthorizations.MetaUpdatedAt),
		NewPaymentAuthorizationsUpdateField(selectFields.MetaUpdatedBy(), paymentAuthorizations.MetaUpdatedBy),
		NewPaymentAuthorizationsUpdateField(selectFields.MetaDeletedAt(), paymentAuthorizations.MetaDeletedAt),
		NewPaymentAuthorizationsUpdateField(selectFields.MetaDeletedBy(), paymentAuthorizations.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentAuthorizationsCommand(paymentAuthorizationsUpdateFieldList PaymentAuthorizationsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentAuthorizationsUpdateFieldList {
		field := string(updateField.paymentAuthorizationsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentAuthorizations(ctx context.Context, paymentAuthorizationsList []*model.PaymentAuthorizations, fieldsInsert ...PaymentAuthorizationsField) (err error) {
	var (
		fieldsStr                      string
		valueListStr                   []string
		argsList                       []interface{}
		primaryIds                     []model.PaymentAuthorizationsPrimaryID
		paymentAuthorizationsValueList []model.PaymentAuthorizations
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentAuthorizationsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentAuthorizations := range paymentAuthorizationsList {

		primaryIds = append(primaryIds, paymentAuthorizations.ToPaymentAuthorizationsPrimaryID())

		paymentAuthorizationsValueList = append(paymentAuthorizationsValueList, *paymentAuthorizations)
	}

	_, notFoundIds, err := repo.IsExistPaymentAuthorizationsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentAuthorizations] failed checking paymentAuthorizations whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentAuthorizationsPrimaryID{}
		mapNotFoundIds := map[model.PaymentAuthorizationsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentAuthorizations", fmt.Sprintf("paymentAuthorizations with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentAuthorizations(paymentAuthorizationsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentAuthorizationsQueries.insertPaymentAuthorizations, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentAuthorizations] failed exec create paymentAuthorizations query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentAuthorizationsByIDs(ctx context.Context, primaryIDs []model.PaymentAuthorizationsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentAuthorizationsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentAuthorizationsByIDs] failed checking paymentAuthorizations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAuthorizations with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_authorizations\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentAuthorizationsQueries.deletePaymentAuthorizations + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentAuthorizationsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentAuthorizationsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentAuthorizationsByIDs(ctx context.Context, ids []model.PaymentAuthorizationsPrimaryID) (exists bool, notFoundIds []model.PaymentAuthorizationsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_authorizations\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentAuthorizationsQueries.selectPaymentAuthorizations, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentAuthorizationsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentAuthorizationsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentAuthorizationsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentAuthorizationsPrimaryID]bool{}
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

// BulkUpdatePaymentAuthorizations is used to bulk update paymentAuthorizations, by default it will update all field
// if want to update specific field, then fill paymentAuthorizationssMapUpdateFieldsRequest else please fill paymentAuthorizationssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentAuthorizations(ctx context.Context, paymentAuthorizationssMap map[model.PaymentAuthorizationsPrimaryID]*model.PaymentAuthorizations, paymentAuthorizationssMapUpdateFieldsRequest map[model.PaymentAuthorizationsPrimaryID]PaymentAuthorizationsUpdateFieldList) (err error) {
	if len(paymentAuthorizationssMap) == 0 && len(paymentAuthorizationssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentAuthorizationssMapUpdateField map[model.PaymentAuthorizationsPrimaryID]PaymentAuthorizationsUpdateFieldList = map[model.PaymentAuthorizationsPrimaryID]PaymentAuthorizationsUpdateFieldList{}
		asTableValues                        string                                                                        = "myvalues"
	)

	if len(paymentAuthorizationssMap) > 0 {
		for id, paymentAuthorizations := range paymentAuthorizationssMap {
			if paymentAuthorizations == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentAuthorizations] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentAuthorizationssMapUpdateField[id] = defaultPaymentAuthorizationsUpdateFields(*paymentAuthorizations)
		}
	} else {
		paymentAuthorizationssMapUpdateField = paymentAuthorizationssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentAuthorizationsQuery(paymentAuthorizationssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentAuthorizationsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentAuthorizations] failed checking paymentAuthorizations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAuthorizations with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentAuthorizationsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_authorizations\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentAuthorizations] failed exec query")
	}
	return
}

type PaymentAuthorizationsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentAuthorizationsFieldParameter(param string, args ...interface{}) PaymentAuthorizationsFieldParameter {
	return PaymentAuthorizationsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentAuthorizationsQuery(mapPaymentAuthorizationss map[model.PaymentAuthorizationsPrimaryID]PaymentAuthorizationsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentAuthorizationsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentAuthorizationsPrimaryID]map[string]interface{}{}
	paymentAuthorizationsSelectFields := NewPaymentAuthorizationsSelectFields()
	for id, updateFields := range mapPaymentAuthorizationss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentAuthorizationsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentAuthorizationss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentAuthorizationsFieldType(updateField.paymentAuthorizationsField)))
			args = append(args, fields[string(updateField.paymentAuthorizationsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentAuthorizationsField))
		if updateField.paymentAuthorizationsField == paymentAuthorizationsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentAuthorizationsField, asTableValues, updateField.paymentAuthorizationsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentAuthorizationsField,
				"\"payment_authorizations\"", updateField.paymentAuthorizationsField,
				asTableValues, updateField.paymentAuthorizationsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentAuthorizationsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentAuthorizationsPrimaryID, asTableValue string) (whereQry string) {
	paymentAuthorizationsSelectFields := NewPaymentAuthorizationsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_authorizations\".\"id\" = %s.\"id\"::"+GetPaymentAuthorizationsFieldType(paymentAuthorizationsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentAuthorizationsFieldType(paymentAuthorizationsField PaymentAuthorizationsField) string {
	selectPaymentAuthorizationsFields := NewPaymentAuthorizationsSelectFields()
	switch paymentAuthorizationsField {

	case selectPaymentAuthorizationsFields.Id():
		return "uuid"

	case selectPaymentAuthorizationsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentAuthorizationsFields.PaymentAttemptId():
		return "uuid"

	case selectPaymentAuthorizationsFields.ProviderAccountId():
		return "uuid"

	case selectPaymentAuthorizationsFields.ProviderAuthorizationId():
		return "text"

	case selectPaymentAuthorizationsFields.Amount():
		return "numeric"

	case selectPaymentAuthorizationsFields.Currency():
		return "text"

	case selectPaymentAuthorizationsFields.Status():
		return "payment_authorization_status_enum"

	case selectPaymentAuthorizationsFields.AuthorizedAt():
		return "timestamptz"

	case selectPaymentAuthorizationsFields.ExpiresAt():
		return "timestamptz"

	case selectPaymentAuthorizationsFields.CapturedAmount():
		return "numeric"

	case selectPaymentAuthorizationsFields.FailureCode():
		return "text"

	case selectPaymentAuthorizationsFields.FailureMessage():
		return "text"

	case selectPaymentAuthorizationsFields.RawRequest():
		return "jsonb"

	case selectPaymentAuthorizationsFields.RawResponse():
		return "jsonb"

	case selectPaymentAuthorizationsFields.Metadata():
		return "jsonb"

	case selectPaymentAuthorizationsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentAuthorizationsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentAuthorizationsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentAuthorizationsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentAuthorizationsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentAuthorizationsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentAuthorizations(ctx context.Context, paymentAuthorizations *model.PaymentAuthorizations, fieldsInsert ...PaymentAuthorizationsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentAuthorizationsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentAuthorizationsPrimaryID{
		Id: paymentAuthorizations.Id,
	}
	exists, err := repo.IsExistPaymentAuthorizationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentAuthorizations] failed checking paymentAuthorizations whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentAuthorizations", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentAuthorizations([]model.PaymentAuthorizations{*paymentAuthorizations}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentAuthorizationsQueries.insertPaymentAuthorizations, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentAuthorizations] failed exec create paymentAuthorizations query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentAuthorizationsByID(ctx context.Context, primaryID model.PaymentAuthorizationsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentAuthorizationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentAuthorizationsByID] failed checking paymentAuthorizations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAuthorizations with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentAuthorizationsCompositePrimaryKeyWhere([]model.PaymentAuthorizationsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentAuthorizationsQueries.deletePaymentAuthorizations + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentAuthorizationsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentAuthorizationsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentAuthorizationsFilterResult, err error) {
	query, args, err := composePaymentAuthorizationsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentAuthorizationsByFilter] failed compose paymentAuthorizations filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentAuthorizationsByFilter] failed get paymentAuthorizations by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentAuthorizationsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentAuthorizationsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentAuthorizationsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentAuthorizationsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentAuthorizationsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 22 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentAuthorizationsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_attempt_id\"")
			selectedColumns["payment_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_authorization_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_authorization_id\"")
			selectedColumns["provider_authorization_id"] = struct{}{}
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
		if _, selected := selectedColumns["authorized_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"authorized_at\"")
			selectedColumns["authorized_at"] = struct{}{}
		}
		if _, selected := selectedColumns["expires_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"expires_at\"")
			selectedColumns["expires_at"] = struct{}{}
		}
		if _, selected := selectedColumns["captured_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"captured_amount\"")
			selectedColumns["captured_amount"] = struct{}{}
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

type paymentAuthorizationsFilterPlaceholder struct {
	index int
}

func (p *paymentAuthorizationsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentAuthorizationsFilterPredicate(filterField model.FilterField, placeholders *paymentAuthorizationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentAuthorizationsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentAuthorizationsFilterSQLExpr(spec)
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

func composePaymentAuthorizationsFilterGroup(group model.FilterGroup, placeholders *paymentAuthorizationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentAuthorizationsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentAuthorizationsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentAuthorizationsFilterWhereQueries(filter model.Filter, placeholders *paymentAuthorizationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentAuthorizationsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentAuthorizationsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentAuthorizationsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentAuthorizationsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentAuthorizationsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentAuthorizationsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentAuthorizationsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentAuthorizationsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentAuthorizationsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentAuthorizationsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentAuthorizationsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_authorizations\" base%s", strings.Join(selectColumns, ","), composePaymentAuthorizationsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentAuthorizationsByID(ctx context.Context, primaryID model.PaymentAuthorizationsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentAuthorizationsCompositePrimaryKeyWhere([]model.PaymentAuthorizationsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentAuthorizationsQueries.selectCountPaymentAuthorizations, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentAuthorizationsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentAuthorizations(ctx context.Context, selectFields ...PaymentAuthorizationsField) (paymentAuthorizationsList model.PaymentAuthorizationsList, err error) {
	var (
		defaultPaymentAuthorizationsSelectFields = defaultPaymentAuthorizationsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentAuthorizationsSelectFields = composePaymentAuthorizationsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentAuthorizationsQueries.selectPaymentAuthorizations, defaultPaymentAuthorizationsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentAuthorizationsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentAuthorizations] failed get paymentAuthorizations list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentAuthorizationsByID(ctx context.Context, primaryID model.PaymentAuthorizationsPrimaryID, selectFields ...PaymentAuthorizationsField) (paymentAuthorizations model.PaymentAuthorizations, err error) {
	var (
		defaultPaymentAuthorizationsSelectFields = defaultPaymentAuthorizationsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentAuthorizationsSelectFields = composePaymentAuthorizationsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentAuthorizationsCompositePrimaryKeyWhere([]model.PaymentAuthorizationsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentAuthorizationsQueries.selectPaymentAuthorizations+" WHERE "+whereQry, defaultPaymentAuthorizationsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentAuthorizations, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentAuthorizations with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentAuthorizationsByID] failed get paymentAuthorizations")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentAuthorizationsByID(ctx context.Context, primaryID model.PaymentAuthorizationsPrimaryID, paymentAuthorizations *model.PaymentAuthorizations, paymentAuthorizationsUpdateFields ...PaymentAuthorizationsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentAuthorizationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAuthorizations] failed checking paymentAuthorizations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentAuthorizations with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentAuthorizations == nil {
		if len(paymentAuthorizationsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentAuthorizationsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentAuthorizations = &model.PaymentAuthorizations{}
	}
	var (
		defaultPaymentAuthorizationsUpdateFields = defaultPaymentAuthorizationsUpdateFields(*paymentAuthorizations)
		tempUpdateField                          PaymentAuthorizationsUpdateFieldList
		selectFields                             = NewPaymentAuthorizationsSelectFields()
	)
	if len(paymentAuthorizationsUpdateFields) > 0 {
		for _, updateField := range paymentAuthorizationsUpdateFields {
			if updateField.paymentAuthorizationsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentAuthorizationsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentAuthorizationsCompositePrimaryKeyWhere([]model.PaymentAuthorizationsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentAuthorizationsCommand(defaultPaymentAuthorizationsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentAuthorizationsQueries.updatePaymentAuthorizations+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAuthorizations] error when try to update paymentAuthorizations by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentAuthorizationsByFilter(ctx context.Context, filter model.Filter, paymentAuthorizationsUpdateFields ...PaymentAuthorizationsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentAuthorizationsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentAuthorizationsUpdateFieldList
		selectFields = NewPaymentAuthorizationsSelectFields()
	)
	for _, updateField := range paymentAuthorizationsUpdateFields {
		if updateField.paymentAuthorizationsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentAuthorizationsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentAuthorizationsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentAuthorizationsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_authorizations\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAuthorizationsByFilter] error when try to update paymentAuthorizations by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentAuthorizationsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentAuthorizationsQueries = struct {
		selectPaymentAuthorizations      string
		selectCountPaymentAuthorizations string
		deletePaymentAuthorizations      string
		updatePaymentAuthorizations      string
		insertPaymentAuthorizations      string
	}{
		selectPaymentAuthorizations:      "SELECT %s FROM \"payment_authorizations\"",
		selectCountPaymentAuthorizations: "SELECT COUNT(\"id\") FROM \"payment_authorizations\"",
		deletePaymentAuthorizations:      "DELETE FROM \"payment_authorizations\"",
		updatePaymentAuthorizations:      "UPDATE \"payment_authorizations\" SET %s ",
		insertPaymentAuthorizations:      "INSERT INTO \"payment_authorizations\" %s VALUES %s",
	}
)

type PaymentAuthorizationsRepository interface {
	CreatePaymentAuthorizations(ctx context.Context, paymentAuthorizations *model.PaymentAuthorizations, fieldsInsert ...PaymentAuthorizationsField) error
	BulkCreatePaymentAuthorizations(ctx context.Context, paymentAuthorizationsList []*model.PaymentAuthorizations, fieldsInsert ...PaymentAuthorizationsField) error
	ResolvePaymentAuthorizations(ctx context.Context, selectFields ...PaymentAuthorizationsField) (model.PaymentAuthorizationsList, error)
	ResolvePaymentAuthorizationsByID(ctx context.Context, primaryID model.PaymentAuthorizationsPrimaryID, selectFields ...PaymentAuthorizationsField) (model.PaymentAuthorizations, error)
	UpdatePaymentAuthorizationsByID(ctx context.Context, id model.PaymentAuthorizationsPrimaryID, paymentAuthorizations *model.PaymentAuthorizations, paymentAuthorizationsUpdateFields ...PaymentAuthorizationsUpdateField) error
	UpdatePaymentAuthorizationsByFilter(ctx context.Context, filter model.Filter, paymentAuthorizationsUpdateFields ...PaymentAuthorizationsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentAuthorizations(ctx context.Context, paymentAuthorizationsListMap map[model.PaymentAuthorizationsPrimaryID]*model.PaymentAuthorizations, PaymentAuthorizationssMapUpdateFieldsRequest map[model.PaymentAuthorizationsPrimaryID]PaymentAuthorizationsUpdateFieldList) (err error)
	DeletePaymentAuthorizationsByID(ctx context.Context, id model.PaymentAuthorizationsPrimaryID) error
	BulkDeletePaymentAuthorizationsByIDs(ctx context.Context, ids []model.PaymentAuthorizationsPrimaryID) error
	ResolvePaymentAuthorizationsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentAuthorizationsFilterResult, err error)
	IsExistPaymentAuthorizationsByIDs(ctx context.Context, ids []model.PaymentAuthorizationsPrimaryID) (exists bool, notFoundIds []model.PaymentAuthorizationsPrimaryID, err error)
	IsExistPaymentAuthorizationsByID(ctx context.Context, id model.PaymentAuthorizationsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
