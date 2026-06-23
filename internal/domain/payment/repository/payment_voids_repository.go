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

func composeInsertFieldsAndParamsPaymentVoids(paymentVoidsList []model.PaymentVoids, fieldsInsert ...PaymentVoidsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentVoidsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentVoids := range paymentVoidsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentVoids.Id)
			case selectField.PaymentAuthorizationId():
				args = append(args, paymentVoids.PaymentAuthorizationId)
			case selectField.PaymentIntentId():
				args = append(args, paymentVoids.PaymentIntentId)
			case selectField.Amount():
				args = append(args, paymentVoids.Amount)
			case selectField.Currency():
				args = append(args, paymentVoids.Currency)
			case selectField.Status():
				args = append(args, paymentVoids.Status)
			case selectField.ProviderVoidId():
				args = append(args, paymentVoids.ProviderVoidId)
			case selectField.VoidedAt():
				args = append(args, paymentVoids.VoidedAt)
			case selectField.FailureCode():
				args = append(args, paymentVoids.FailureCode)
			case selectField.FailureMessage():
				args = append(args, paymentVoids.FailureMessage)
			case selectField.RawRequest():
				args = append(args, paymentVoids.RawRequest)
			case selectField.RawResponse():
				args = append(args, paymentVoids.RawResponse)
			case selectField.Metadata():
				args = append(args, paymentVoids.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentVoids.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentVoids.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentVoids.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentVoids.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentVoids.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentVoids.MetaDeletedBy)

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

func composePaymentVoidsCompositePrimaryKeyWhere(primaryIDs []model.PaymentVoidsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_voids\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentVoidsSelectFields() string {
	fields := NewPaymentVoidsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentVoidsSelectFields(selectFields ...PaymentVoidsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentVoidsField string
type PaymentVoidsFieldList []PaymentVoidsField

type PaymentVoidsSelectFields struct {
}

func (ss PaymentVoidsSelectFields) Id() PaymentVoidsField {
	return PaymentVoidsField("id")
}

func (ss PaymentVoidsSelectFields) PaymentAuthorizationId() PaymentVoidsField {
	return PaymentVoidsField("payment_authorization_id")
}

func (ss PaymentVoidsSelectFields) PaymentIntentId() PaymentVoidsField {
	return PaymentVoidsField("payment_intent_id")
}

func (ss PaymentVoidsSelectFields) Amount() PaymentVoidsField {
	return PaymentVoidsField("amount")
}

func (ss PaymentVoidsSelectFields) Currency() PaymentVoidsField {
	return PaymentVoidsField("currency")
}

func (ss PaymentVoidsSelectFields) Status() PaymentVoidsField {
	return PaymentVoidsField("status")
}

func (ss PaymentVoidsSelectFields) ProviderVoidId() PaymentVoidsField {
	return PaymentVoidsField("provider_void_id")
}

func (ss PaymentVoidsSelectFields) VoidedAt() PaymentVoidsField {
	return PaymentVoidsField("voided_at")
}

func (ss PaymentVoidsSelectFields) FailureCode() PaymentVoidsField {
	return PaymentVoidsField("failure_code")
}

func (ss PaymentVoidsSelectFields) FailureMessage() PaymentVoidsField {
	return PaymentVoidsField("failure_message")
}

func (ss PaymentVoidsSelectFields) RawRequest() PaymentVoidsField {
	return PaymentVoidsField("raw_request")
}

func (ss PaymentVoidsSelectFields) RawResponse() PaymentVoidsField {
	return PaymentVoidsField("raw_response")
}

func (ss PaymentVoidsSelectFields) Metadata() PaymentVoidsField {
	return PaymentVoidsField("metadata")
}

func (ss PaymentVoidsSelectFields) MetaCreatedAt() PaymentVoidsField {
	return PaymentVoidsField("meta_created_at")
}

func (ss PaymentVoidsSelectFields) MetaCreatedBy() PaymentVoidsField {
	return PaymentVoidsField("meta_created_by")
}

func (ss PaymentVoidsSelectFields) MetaUpdatedAt() PaymentVoidsField {
	return PaymentVoidsField("meta_updated_at")
}

func (ss PaymentVoidsSelectFields) MetaUpdatedBy() PaymentVoidsField {
	return PaymentVoidsField("meta_updated_by")
}

func (ss PaymentVoidsSelectFields) MetaDeletedAt() PaymentVoidsField {
	return PaymentVoidsField("meta_deleted_at")
}

func (ss PaymentVoidsSelectFields) MetaDeletedBy() PaymentVoidsField {
	return PaymentVoidsField("meta_deleted_by")
}

func (ss PaymentVoidsSelectFields) All() PaymentVoidsFieldList {
	return []PaymentVoidsField{
		ss.Id(),
		ss.PaymentAuthorizationId(),
		ss.PaymentIntentId(),
		ss.Amount(),
		ss.Currency(),
		ss.Status(),
		ss.ProviderVoidId(),
		ss.VoidedAt(),
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

func NewPaymentVoidsSelectFields() PaymentVoidsSelectFields {
	return PaymentVoidsSelectFields{}
}

type PaymentVoidsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentVoidsUpdateField struct {
	paymentVoidsField PaymentVoidsField
	opt               PaymentVoidsUpdateFieldOption
	value             interface{}
}
type PaymentVoidsUpdateFieldList []PaymentVoidsUpdateField

func defaultPaymentVoidsUpdateFieldOption() PaymentVoidsUpdateFieldOption {
	return PaymentVoidsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentVoidsOption(useIncrement bool) func(*PaymentVoidsUpdateFieldOption) {
	return func(pcufo *PaymentVoidsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentVoidsUpdateField(field PaymentVoidsField, val interface{}, opts ...func(*PaymentVoidsUpdateFieldOption)) PaymentVoidsUpdateField {
	defaultOpt := defaultPaymentVoidsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentVoidsUpdateField{
		paymentVoidsField: field,
		value:             val,
		opt:               defaultOpt,
	}
}
func defaultPaymentVoidsUpdateFields(paymentVoids model.PaymentVoids) (paymentVoidsUpdateFieldList PaymentVoidsUpdateFieldList) {
	selectFields := NewPaymentVoidsSelectFields()
	paymentVoidsUpdateFieldList = append(paymentVoidsUpdateFieldList,
		NewPaymentVoidsUpdateField(selectFields.Id(), paymentVoids.Id),
		NewPaymentVoidsUpdateField(selectFields.PaymentAuthorizationId(), paymentVoids.PaymentAuthorizationId),
		NewPaymentVoidsUpdateField(selectFields.PaymentIntentId(), paymentVoids.PaymentIntentId),
		NewPaymentVoidsUpdateField(selectFields.Amount(), paymentVoids.Amount),
		NewPaymentVoidsUpdateField(selectFields.Currency(), paymentVoids.Currency),
		NewPaymentVoidsUpdateField(selectFields.Status(), paymentVoids.Status),
		NewPaymentVoidsUpdateField(selectFields.ProviderVoidId(), paymentVoids.ProviderVoidId),
		NewPaymentVoidsUpdateField(selectFields.VoidedAt(), paymentVoids.VoidedAt),
		NewPaymentVoidsUpdateField(selectFields.FailureCode(), paymentVoids.FailureCode),
		NewPaymentVoidsUpdateField(selectFields.FailureMessage(), paymentVoids.FailureMessage),
		NewPaymentVoidsUpdateField(selectFields.RawRequest(), paymentVoids.RawRequest),
		NewPaymentVoidsUpdateField(selectFields.RawResponse(), paymentVoids.RawResponse),
		NewPaymentVoidsUpdateField(selectFields.Metadata(), paymentVoids.Metadata),
		NewPaymentVoidsUpdateField(selectFields.MetaCreatedAt(), paymentVoids.MetaCreatedAt),
		NewPaymentVoidsUpdateField(selectFields.MetaCreatedBy(), paymentVoids.MetaCreatedBy),
		NewPaymentVoidsUpdateField(selectFields.MetaUpdatedAt(), paymentVoids.MetaUpdatedAt),
		NewPaymentVoidsUpdateField(selectFields.MetaUpdatedBy(), paymentVoids.MetaUpdatedBy),
		NewPaymentVoidsUpdateField(selectFields.MetaDeletedAt(), paymentVoids.MetaDeletedAt),
		NewPaymentVoidsUpdateField(selectFields.MetaDeletedBy(), paymentVoids.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentVoidsCommand(paymentVoidsUpdateFieldList PaymentVoidsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentVoidsUpdateFieldList {
		field := string(updateField.paymentVoidsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentVoids(ctx context.Context, paymentVoidsList []*model.PaymentVoids, fieldsInsert ...PaymentVoidsField) (err error) {
	var (
		fieldsStr             string
		valueListStr          []string
		argsList              []interface{}
		primaryIds            []model.PaymentVoidsPrimaryID
		paymentVoidsValueList []model.PaymentVoids
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentVoidsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentVoids := range paymentVoidsList {

		primaryIds = append(primaryIds, paymentVoids.ToPaymentVoidsPrimaryID())

		paymentVoidsValueList = append(paymentVoidsValueList, *paymentVoids)
	}

	_, notFoundIds, err := repo.IsExistPaymentVoidsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentVoids] failed checking paymentVoids whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentVoidsPrimaryID{}
		mapNotFoundIds := map[model.PaymentVoidsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentVoids", fmt.Sprintf("paymentVoids with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentVoids(paymentVoidsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentVoidsQueries.insertPaymentVoids, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentVoids] failed exec create paymentVoids query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentVoidsByIDs(ctx context.Context, primaryIDs []model.PaymentVoidsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentVoidsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentVoidsByIDs] failed checking paymentVoids whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentVoids with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_voids\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentVoidsQueries.deletePaymentVoids + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentVoidsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentVoidsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentVoidsByIDs(ctx context.Context, ids []model.PaymentVoidsPrimaryID) (exists bool, notFoundIds []model.PaymentVoidsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_voids\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentVoidsQueries.selectPaymentVoids, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentVoidsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentVoidsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentVoidsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentVoidsPrimaryID]bool{}
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

// BulkUpdatePaymentVoids is used to bulk update paymentVoids, by default it will update all field
// if want to update specific field, then fill paymentVoidssMapUpdateFieldsRequest else please fill paymentVoidssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentVoids(ctx context.Context, paymentVoidssMap map[model.PaymentVoidsPrimaryID]*model.PaymentVoids, paymentVoidssMapUpdateFieldsRequest map[model.PaymentVoidsPrimaryID]PaymentVoidsUpdateFieldList) (err error) {
	if len(paymentVoidssMap) == 0 && len(paymentVoidssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentVoidssMapUpdateField map[model.PaymentVoidsPrimaryID]PaymentVoidsUpdateFieldList = map[model.PaymentVoidsPrimaryID]PaymentVoidsUpdateFieldList{}
		asTableValues               string                                                      = "myvalues"
	)

	if len(paymentVoidssMap) > 0 {
		for id, paymentVoids := range paymentVoidssMap {
			if paymentVoids == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentVoids] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentVoidssMapUpdateField[id] = defaultPaymentVoidsUpdateFields(*paymentVoids)
		}
	} else {
		paymentVoidssMapUpdateField = paymentVoidssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentVoidsQuery(paymentVoidssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentVoidsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentVoids] failed checking paymentVoids whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentVoids with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentVoidsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_voids\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentVoids] failed exec query")
	}
	return
}

type PaymentVoidsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentVoidsFieldParameter(param string, args ...interface{}) PaymentVoidsFieldParameter {
	return PaymentVoidsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentVoidsQuery(mapPaymentVoidss map[model.PaymentVoidsPrimaryID]PaymentVoidsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentVoidsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentVoidsPrimaryID]map[string]interface{}{}
	paymentVoidsSelectFields := NewPaymentVoidsSelectFields()
	for id, updateFields := range mapPaymentVoidss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentVoidsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentVoidss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentVoidsFieldType(updateField.paymentVoidsField)))
			args = append(args, fields[string(updateField.paymentVoidsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentVoidsField))
		if updateField.paymentVoidsField == paymentVoidsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentVoidsField, asTableValues, updateField.paymentVoidsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentVoidsField,
				"\"payment_voids\"", updateField.paymentVoidsField,
				asTableValues, updateField.paymentVoidsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentVoidsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentVoidsPrimaryID, asTableValue string) (whereQry string) {
	paymentVoidsSelectFields := NewPaymentVoidsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_voids\".\"id\" = %s.\"id\"::"+GetPaymentVoidsFieldType(paymentVoidsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentVoidsFieldType(paymentVoidsField PaymentVoidsField) string {
	selectPaymentVoidsFields := NewPaymentVoidsSelectFields()
	switch paymentVoidsField {

	case selectPaymentVoidsFields.Id():
		return "uuid"

	case selectPaymentVoidsFields.PaymentAuthorizationId():
		return "uuid"

	case selectPaymentVoidsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentVoidsFields.Amount():
		return "numeric"

	case selectPaymentVoidsFields.Currency():
		return "text"

	case selectPaymentVoidsFields.Status():
		return "payment_void_status_enum"

	case selectPaymentVoidsFields.ProviderVoidId():
		return "text"

	case selectPaymentVoidsFields.VoidedAt():
		return "timestamptz"

	case selectPaymentVoidsFields.FailureCode():
		return "text"

	case selectPaymentVoidsFields.FailureMessage():
		return "text"

	case selectPaymentVoidsFields.RawRequest():
		return "jsonb"

	case selectPaymentVoidsFields.RawResponse():
		return "jsonb"

	case selectPaymentVoidsFields.Metadata():
		return "jsonb"

	case selectPaymentVoidsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentVoidsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentVoidsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentVoidsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentVoidsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentVoidsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentVoids(ctx context.Context, paymentVoids *model.PaymentVoids, fieldsInsert ...PaymentVoidsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentVoidsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentVoidsPrimaryID{
		Id: paymentVoids.Id,
	}
	exists, err := repo.IsExistPaymentVoidsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentVoids] failed checking paymentVoids whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentVoids", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentVoids([]model.PaymentVoids{*paymentVoids}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentVoidsQueries.insertPaymentVoids, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentVoids] failed exec create paymentVoids query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentVoidsByID(ctx context.Context, primaryID model.PaymentVoidsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentVoidsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentVoidsByID] failed checking paymentVoids whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentVoids with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentVoidsCompositePrimaryKeyWhere([]model.PaymentVoidsPrimaryID{primaryID})
	commandQuery := paymentVoidsQueries.deletePaymentVoids + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentVoidsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentVoidsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentVoidsFilterResult, err error) {
	query, args, err := composePaymentVoidsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentVoidsByFilter] failed compose paymentVoids filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentVoidsByFilter] failed get paymentVoids by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentVoidsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentVoidsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentVoidsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentVoidsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentVoidsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentVoidsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 19+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_authorization_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_authorization_id\"")
			selectedColumns["payment_authorization_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
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
		if _, selected := selectedColumns["provider_void_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_void_id\"")
			selectedColumns["provider_void_id"] = struct{}{}
		}
		if _, selected := selectedColumns["voided_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"voided_at\"")
			selectedColumns["voided_at"] = struct{}{}
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

type paymentVoidsFilterPlaceholder struct {
	index int
}

func (p *paymentVoidsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentVoidsFilterPredicate(filterField model.FilterField, placeholders *paymentVoidsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentVoidsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentVoidsFilterSQLExpr(spec)
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

func composePaymentVoidsFilterGroup(group model.FilterGroup, placeholders *paymentVoidsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentVoidsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentVoidsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentVoidsFilterWhereQueries(filter model.Filter, placeholders *paymentVoidsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentVoidsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentVoidsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentVoidsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentVoidsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentVoidsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentVoidsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentVoidsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentVoidsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentVoidsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentVoidsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentVoidsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_voids\" base%s", strings.Join(selectColumns, ","), composePaymentVoidsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentVoidsByID(ctx context.Context, primaryID model.PaymentVoidsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentVoidsCompositePrimaryKeyWhere([]model.PaymentVoidsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentVoidsQueries.selectCountPaymentVoids, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentVoidsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentVoids(ctx context.Context, selectFields ...PaymentVoidsField) (paymentVoidsList model.PaymentVoidsList, err error) {
	var (
		defaultPaymentVoidsSelectFields = defaultPaymentVoidsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentVoidsSelectFields = composePaymentVoidsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentVoidsQueries.selectPaymentVoids, defaultPaymentVoidsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentVoidsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentVoids] failed get paymentVoids list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentVoidsByID(ctx context.Context, primaryID model.PaymentVoidsPrimaryID, selectFields ...PaymentVoidsField) (paymentVoids model.PaymentVoids, err error) {
	var (
		defaultPaymentVoidsSelectFields = defaultPaymentVoidsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentVoidsSelectFields = composePaymentVoidsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentVoidsCompositePrimaryKeyWhere([]model.PaymentVoidsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentVoidsQueries.selectPaymentVoids+" WHERE "+whereQry, defaultPaymentVoidsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentVoids, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentVoids with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentVoidsByID] failed get paymentVoids")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentVoidsByID(ctx context.Context, primaryID model.PaymentVoidsPrimaryID, paymentVoids *model.PaymentVoids, paymentVoidsUpdateFields ...PaymentVoidsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentVoidsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentVoids] failed checking paymentVoids whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentVoids with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentVoids == nil {
		if len(paymentVoidsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentVoidsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentVoids = &model.PaymentVoids{}
	}
	var (
		defaultPaymentVoidsUpdateFields = defaultPaymentVoidsUpdateFields(*paymentVoids)
		tempUpdateField                 PaymentVoidsUpdateFieldList
		selectFields                    = NewPaymentVoidsSelectFields()
	)
	if len(paymentVoidsUpdateFields) > 0 {
		for _, updateField := range paymentVoidsUpdateFields {
			if updateField.paymentVoidsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentVoidsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentVoidsCompositePrimaryKeyWhere([]model.PaymentVoidsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentVoidsCommand(defaultPaymentVoidsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentVoidsQueries.updatePaymentVoids+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentVoids] error when try to update paymentVoids by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentVoidsByFilter(ctx context.Context, filter model.Filter, paymentVoidsUpdateFields ...PaymentVoidsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentVoidsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentVoidsUpdateFieldList
		selectFields = NewPaymentVoidsSelectFields()
	)
	for _, updateField := range paymentVoidsUpdateFields {
		if updateField.paymentVoidsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentVoidsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentVoidsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentVoidsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_voids\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentVoidsByFilter] error when try to update paymentVoids by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentVoidsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentVoidsQueries = struct {
		selectPaymentVoids      string
		selectCountPaymentVoids string
		deletePaymentVoids      string
		updatePaymentVoids      string
		insertPaymentVoids      string
	}{
		selectPaymentVoids:      "SELECT %s FROM \"payment_voids\"",
		selectCountPaymentVoids: "SELECT COUNT(\"id\") FROM \"payment_voids\"",
		deletePaymentVoids:      "DELETE FROM \"payment_voids\"",
		updatePaymentVoids:      "UPDATE \"payment_voids\" SET %s ",
		insertPaymentVoids:      "INSERT INTO \"payment_voids\" %s VALUES %s",
	}
)

type PaymentVoidsRepository interface {
	CreatePaymentVoids(ctx context.Context, paymentVoids *model.PaymentVoids, fieldsInsert ...PaymentVoidsField) error
	BulkCreatePaymentVoids(ctx context.Context, paymentVoidsList []*model.PaymentVoids, fieldsInsert ...PaymentVoidsField) error
	ResolvePaymentVoids(ctx context.Context, selectFields ...PaymentVoidsField) (model.PaymentVoidsList, error)
	ResolvePaymentVoidsByID(ctx context.Context, primaryID model.PaymentVoidsPrimaryID, selectFields ...PaymentVoidsField) (model.PaymentVoids, error)
	UpdatePaymentVoidsByID(ctx context.Context, id model.PaymentVoidsPrimaryID, paymentVoids *model.PaymentVoids, paymentVoidsUpdateFields ...PaymentVoidsUpdateField) error
	UpdatePaymentVoidsByFilter(ctx context.Context, filter model.Filter, paymentVoidsUpdateFields ...PaymentVoidsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentVoids(ctx context.Context, paymentVoidsListMap map[model.PaymentVoidsPrimaryID]*model.PaymentVoids, PaymentVoidssMapUpdateFieldsRequest map[model.PaymentVoidsPrimaryID]PaymentVoidsUpdateFieldList) (err error)
	DeletePaymentVoidsByID(ctx context.Context, id model.PaymentVoidsPrimaryID) error
	BulkDeletePaymentVoidsByIDs(ctx context.Context, ids []model.PaymentVoidsPrimaryID) error
	ResolvePaymentVoidsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentVoidsFilterResult, err error)
	IsExistPaymentVoidsByIDs(ctx context.Context, ids []model.PaymentVoidsPrimaryID) (exists bool, notFoundIds []model.PaymentVoidsPrimaryID, err error)
	IsExistPaymentVoidsByID(ctx context.Context, id model.PaymentVoidsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
