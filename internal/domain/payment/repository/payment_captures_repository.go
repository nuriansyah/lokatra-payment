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

func composeInsertFieldsAndParamsPaymentCaptures(paymentCapturesList []model.PaymentCaptures, fieldsInsert ...PaymentCapturesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentCapturesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentCaptures := range paymentCapturesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentCaptures.Id)
			case selectField.PaymentAuthorizationId():
				args = append(args, paymentCaptures.PaymentAuthorizationId)
			case selectField.PaymentIntentId():
				args = append(args, paymentCaptures.PaymentIntentId)
			case selectField.Amount():
				args = append(args, paymentCaptures.Amount)
			case selectField.Currency():
				args = append(args, paymentCaptures.Currency)
			case selectField.Status():
				args = append(args, paymentCaptures.Status)
			case selectField.ProviderCaptureId():
				args = append(args, paymentCaptures.ProviderCaptureId)
			case selectField.CapturedAt():
				args = append(args, paymentCaptures.CapturedAt)
			case selectField.FailureCode():
				args = append(args, paymentCaptures.FailureCode)
			case selectField.FailureMessage():
				args = append(args, paymentCaptures.FailureMessage)
			case selectField.RawRequest():
				args = append(args, paymentCaptures.RawRequest)
			case selectField.RawResponse():
				args = append(args, paymentCaptures.RawResponse)
			case selectField.Metadata():
				args = append(args, paymentCaptures.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentCaptures.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentCaptures.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentCaptures.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentCaptures.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentCaptures.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentCaptures.MetaDeletedBy)

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

func composePaymentCapturesCompositePrimaryKeyWhere(primaryIDs []model.PaymentCapturesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_captures\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentCapturesSelectFields() string {
	fields := NewPaymentCapturesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentCapturesSelectFields(selectFields ...PaymentCapturesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentCapturesField string
type PaymentCapturesFieldList []PaymentCapturesField

type PaymentCapturesSelectFields struct {
}

func (ss PaymentCapturesSelectFields) Id() PaymentCapturesField {
	return PaymentCapturesField("id")
}

func (ss PaymentCapturesSelectFields) PaymentAuthorizationId() PaymentCapturesField {
	return PaymentCapturesField("payment_authorization_id")
}

func (ss PaymentCapturesSelectFields) PaymentIntentId() PaymentCapturesField {
	return PaymentCapturesField("payment_intent_id")
}

func (ss PaymentCapturesSelectFields) Amount() PaymentCapturesField {
	return PaymentCapturesField("amount")
}

func (ss PaymentCapturesSelectFields) Currency() PaymentCapturesField {
	return PaymentCapturesField("currency")
}

func (ss PaymentCapturesSelectFields) Status() PaymentCapturesField {
	return PaymentCapturesField("status")
}

func (ss PaymentCapturesSelectFields) ProviderCaptureId() PaymentCapturesField {
	return PaymentCapturesField("provider_capture_id")
}

func (ss PaymentCapturesSelectFields) CapturedAt() PaymentCapturesField {
	return PaymentCapturesField("captured_at")
}

func (ss PaymentCapturesSelectFields) FailureCode() PaymentCapturesField {
	return PaymentCapturesField("failure_code")
}

func (ss PaymentCapturesSelectFields) FailureMessage() PaymentCapturesField {
	return PaymentCapturesField("failure_message")
}

func (ss PaymentCapturesSelectFields) RawRequest() PaymentCapturesField {
	return PaymentCapturesField("raw_request")
}

func (ss PaymentCapturesSelectFields) RawResponse() PaymentCapturesField {
	return PaymentCapturesField("raw_response")
}

func (ss PaymentCapturesSelectFields) Metadata() PaymentCapturesField {
	return PaymentCapturesField("metadata")
}

func (ss PaymentCapturesSelectFields) MetaCreatedAt() PaymentCapturesField {
	return PaymentCapturesField("meta_created_at")
}

func (ss PaymentCapturesSelectFields) MetaCreatedBy() PaymentCapturesField {
	return PaymentCapturesField("meta_created_by")
}

func (ss PaymentCapturesSelectFields) MetaUpdatedAt() PaymentCapturesField {
	return PaymentCapturesField("meta_updated_at")
}

func (ss PaymentCapturesSelectFields) MetaUpdatedBy() PaymentCapturesField {
	return PaymentCapturesField("meta_updated_by")
}

func (ss PaymentCapturesSelectFields) MetaDeletedAt() PaymentCapturesField {
	return PaymentCapturesField("meta_deleted_at")
}

func (ss PaymentCapturesSelectFields) MetaDeletedBy() PaymentCapturesField {
	return PaymentCapturesField("meta_deleted_by")
}

func (ss PaymentCapturesSelectFields) All() PaymentCapturesFieldList {
	return []PaymentCapturesField{
		ss.Id(),
		ss.PaymentAuthorizationId(),
		ss.PaymentIntentId(),
		ss.Amount(),
		ss.Currency(),
		ss.Status(),
		ss.ProviderCaptureId(),
		ss.CapturedAt(),
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

func NewPaymentCapturesSelectFields() PaymentCapturesSelectFields {
	return PaymentCapturesSelectFields{}
}

type PaymentCapturesUpdateFieldOption struct {
	useIncrement bool
}
type PaymentCapturesUpdateField struct {
	paymentCapturesField PaymentCapturesField
	opt                  PaymentCapturesUpdateFieldOption
	value                interface{}
}
type PaymentCapturesUpdateFieldList []PaymentCapturesUpdateField

func defaultPaymentCapturesUpdateFieldOption() PaymentCapturesUpdateFieldOption {
	return PaymentCapturesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentCapturesOption(useIncrement bool) func(*PaymentCapturesUpdateFieldOption) {
	return func(pcufo *PaymentCapturesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentCapturesUpdateField(field PaymentCapturesField, val interface{}, opts ...func(*PaymentCapturesUpdateFieldOption)) PaymentCapturesUpdateField {
	defaultOpt := defaultPaymentCapturesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentCapturesUpdateField{
		paymentCapturesField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultPaymentCapturesUpdateFields(paymentCaptures model.PaymentCaptures) (paymentCapturesUpdateFieldList PaymentCapturesUpdateFieldList) {
	selectFields := NewPaymentCapturesSelectFields()
	paymentCapturesUpdateFieldList = append(paymentCapturesUpdateFieldList,
		NewPaymentCapturesUpdateField(selectFields.Id(), paymentCaptures.Id),
		NewPaymentCapturesUpdateField(selectFields.PaymentAuthorizationId(), paymentCaptures.PaymentAuthorizationId),
		NewPaymentCapturesUpdateField(selectFields.PaymentIntentId(), paymentCaptures.PaymentIntentId),
		NewPaymentCapturesUpdateField(selectFields.Amount(), paymentCaptures.Amount),
		NewPaymentCapturesUpdateField(selectFields.Currency(), paymentCaptures.Currency),
		NewPaymentCapturesUpdateField(selectFields.Status(), paymentCaptures.Status),
		NewPaymentCapturesUpdateField(selectFields.ProviderCaptureId(), paymentCaptures.ProviderCaptureId),
		NewPaymentCapturesUpdateField(selectFields.CapturedAt(), paymentCaptures.CapturedAt),
		NewPaymentCapturesUpdateField(selectFields.FailureCode(), paymentCaptures.FailureCode),
		NewPaymentCapturesUpdateField(selectFields.FailureMessage(), paymentCaptures.FailureMessage),
		NewPaymentCapturesUpdateField(selectFields.RawRequest(), paymentCaptures.RawRequest),
		NewPaymentCapturesUpdateField(selectFields.RawResponse(), paymentCaptures.RawResponse),
		NewPaymentCapturesUpdateField(selectFields.Metadata(), paymentCaptures.Metadata),
		NewPaymentCapturesUpdateField(selectFields.MetaCreatedAt(), paymentCaptures.MetaCreatedAt),
		NewPaymentCapturesUpdateField(selectFields.MetaCreatedBy(), paymentCaptures.MetaCreatedBy),
		NewPaymentCapturesUpdateField(selectFields.MetaUpdatedAt(), paymentCaptures.MetaUpdatedAt),
		NewPaymentCapturesUpdateField(selectFields.MetaUpdatedBy(), paymentCaptures.MetaUpdatedBy),
		NewPaymentCapturesUpdateField(selectFields.MetaDeletedAt(), paymentCaptures.MetaDeletedAt),
		NewPaymentCapturesUpdateField(selectFields.MetaDeletedBy(), paymentCaptures.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentCapturesCommand(paymentCapturesUpdateFieldList PaymentCapturesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentCapturesUpdateFieldList {
		field := string(updateField.paymentCapturesField)
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

func (repo *RepositoryImpl) BulkCreatePaymentCaptures(ctx context.Context, paymentCapturesList []*model.PaymentCaptures, fieldsInsert ...PaymentCapturesField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.PaymentCapturesPrimaryID
		paymentCapturesValueList []model.PaymentCaptures
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentCapturesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentCaptures := range paymentCapturesList {

		primaryIds = append(primaryIds, paymentCaptures.ToPaymentCapturesPrimaryID())

		paymentCapturesValueList = append(paymentCapturesValueList, *paymentCaptures)
	}

	_, notFoundIds, err := repo.IsExistPaymentCapturesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentCaptures] failed checking paymentCaptures whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentCapturesPrimaryID{}
		mapNotFoundIds := map[model.PaymentCapturesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentCaptures", fmt.Sprintf("paymentCaptures with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentCaptures(paymentCapturesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentCapturesQueries.insertPaymentCaptures, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentCaptures] failed exec create paymentCaptures query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentCapturesByIDs(ctx context.Context, primaryIDs []model.PaymentCapturesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentCapturesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentCapturesByIDs] failed checking paymentCaptures whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentCaptures with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_captures\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentCapturesQueries.deletePaymentCaptures + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentCapturesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentCapturesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentCapturesByIDs(ctx context.Context, ids []model.PaymentCapturesPrimaryID) (exists bool, notFoundIds []model.PaymentCapturesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_captures\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentCapturesQueries.selectPaymentCaptures, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentCapturesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentCapturesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentCapturesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentCapturesPrimaryID]bool{}
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

// BulkUpdatePaymentCaptures is used to bulk update paymentCaptures, by default it will update all field
// if want to update specific field, then fill paymentCapturessMapUpdateFieldsRequest else please fill paymentCapturessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentCaptures(ctx context.Context, paymentCapturessMap map[model.PaymentCapturesPrimaryID]*model.PaymentCaptures, paymentCapturessMapUpdateFieldsRequest map[model.PaymentCapturesPrimaryID]PaymentCapturesUpdateFieldList) (err error) {
	if len(paymentCapturessMap) == 0 && len(paymentCapturessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentCapturessMapUpdateField map[model.PaymentCapturesPrimaryID]PaymentCapturesUpdateFieldList = map[model.PaymentCapturesPrimaryID]PaymentCapturesUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(paymentCapturessMap) > 0 {
		for id, paymentCaptures := range paymentCapturessMap {
			if paymentCaptures == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentCaptures] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentCapturessMapUpdateField[id] = defaultPaymentCapturesUpdateFields(*paymentCaptures)
		}
	} else {
		paymentCapturessMapUpdateField = paymentCapturessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentCapturesQuery(paymentCapturessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentCapturesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentCaptures] failed checking paymentCaptures whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentCaptures with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentCapturesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_captures\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentCaptures] failed exec query")
	}
	return
}

type PaymentCapturesFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentCapturesFieldParameter(param string, args ...interface{}) PaymentCapturesFieldParameter {
	return PaymentCapturesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentCapturesQuery(mapPaymentCapturess map[model.PaymentCapturesPrimaryID]PaymentCapturesUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentCapturesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentCapturesPrimaryID]map[string]interface{}{}
	paymentCapturesSelectFields := NewPaymentCapturesSelectFields()
	for id, updateFields := range mapPaymentCapturess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentCapturesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentCapturess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentCapturesFieldType(updateField.paymentCapturesField)))
			args = append(args, fields[string(updateField.paymentCapturesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentCapturesField))
		if updateField.paymentCapturesField == paymentCapturesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentCapturesField, asTableValues, updateField.paymentCapturesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentCapturesField,
				"\"payment_captures\"", updateField.paymentCapturesField,
				asTableValues, updateField.paymentCapturesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentCapturesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentCapturesPrimaryID, asTableValue string) (whereQry string) {
	paymentCapturesSelectFields := NewPaymentCapturesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_captures\".\"id\" = %s.\"id\"::"+GetPaymentCapturesFieldType(paymentCapturesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentCapturesFieldType(paymentCapturesField PaymentCapturesField) string {
	selectPaymentCapturesFields := NewPaymentCapturesSelectFields()
	switch paymentCapturesField {

	case selectPaymentCapturesFields.Id():
		return "uuid"

	case selectPaymentCapturesFields.PaymentAuthorizationId():
		return "uuid"

	case selectPaymentCapturesFields.PaymentIntentId():
		return "uuid"

	case selectPaymentCapturesFields.Amount():
		return "numeric"

	case selectPaymentCapturesFields.Currency():
		return "text"

	case selectPaymentCapturesFields.Status():
		return "payment_capture_status_enum"

	case selectPaymentCapturesFields.ProviderCaptureId():
		return "text"

	case selectPaymentCapturesFields.CapturedAt():
		return "timestamptz"

	case selectPaymentCapturesFields.FailureCode():
		return "text"

	case selectPaymentCapturesFields.FailureMessage():
		return "text"

	case selectPaymentCapturesFields.RawRequest():
		return "jsonb"

	case selectPaymentCapturesFields.RawResponse():
		return "jsonb"

	case selectPaymentCapturesFields.Metadata():
		return "jsonb"

	case selectPaymentCapturesFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentCapturesFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentCapturesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentCapturesFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentCapturesFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentCapturesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentCaptures(ctx context.Context, paymentCaptures *model.PaymentCaptures, fieldsInsert ...PaymentCapturesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentCapturesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentCapturesPrimaryID{
		Id: paymentCaptures.Id,
	}
	exists, err := repo.IsExistPaymentCapturesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentCaptures] failed checking paymentCaptures whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentCaptures", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentCaptures([]model.PaymentCaptures{*paymentCaptures}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentCapturesQueries.insertPaymentCaptures, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentCaptures] failed exec create paymentCaptures query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentCapturesByID(ctx context.Context, primaryID model.PaymentCapturesPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentCapturesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentCapturesByID] failed checking paymentCaptures whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentCaptures with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentCapturesCompositePrimaryKeyWhere([]model.PaymentCapturesPrimaryID{primaryID})
	commandQuery := paymentCapturesQueries.deletePaymentCaptures + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentCapturesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentCapturesByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentCapturesFilterResult, err error) {
	query, args, err := composePaymentCapturesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentCapturesByFilter] failed compose paymentCaptures filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentCapturesByFilter] failed get paymentCaptures by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentCapturesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentCapturesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentCapturesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentCapturesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentCapturesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentCapturesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["provider_capture_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_capture_id\"")
			selectedColumns["provider_capture_id"] = struct{}{}
		}
		if _, selected := selectedColumns["captured_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"captured_at\"")
			selectedColumns["captured_at"] = struct{}{}
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

type paymentCapturesFilterPlaceholder struct {
	index int
}

func (p *paymentCapturesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentCapturesFilterPredicate(filterField model.FilterField, placeholders *paymentCapturesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentCapturesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentCapturesFilterSQLExpr(spec)
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

func composePaymentCapturesFilterGroup(group model.FilterGroup, placeholders *paymentCapturesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentCapturesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentCapturesFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentCapturesFilterWhereQueries(filter model.Filter, placeholders *paymentCapturesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentCapturesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentCapturesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentCapturesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentCapturesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentCapturesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentCapturesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentCapturesFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentCapturesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentCapturesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentCapturesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentCapturesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_captures\" base%s", strings.Join(selectColumns, ","), composePaymentCapturesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentCapturesByID(ctx context.Context, primaryID model.PaymentCapturesPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentCapturesCompositePrimaryKeyWhere([]model.PaymentCapturesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentCapturesQueries.selectCountPaymentCaptures, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentCapturesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentCaptures(ctx context.Context, selectFields ...PaymentCapturesField) (paymentCapturesList model.PaymentCapturesList, err error) {
	var (
		defaultPaymentCapturesSelectFields = defaultPaymentCapturesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentCapturesSelectFields = composePaymentCapturesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentCapturesQueries.selectPaymentCaptures, defaultPaymentCapturesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentCapturesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentCaptures] failed get paymentCaptures list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentCapturesByID(ctx context.Context, primaryID model.PaymentCapturesPrimaryID, selectFields ...PaymentCapturesField) (paymentCaptures model.PaymentCaptures, err error) {
	var (
		defaultPaymentCapturesSelectFields = defaultPaymentCapturesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentCapturesSelectFields = composePaymentCapturesSelectFields(selectFields...)
	}
	whereQry, params := composePaymentCapturesCompositePrimaryKeyWhere([]model.PaymentCapturesPrimaryID{primaryID})
	query := fmt.Sprintf(paymentCapturesQueries.selectPaymentCaptures+" WHERE "+whereQry, defaultPaymentCapturesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentCaptures, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentCaptures with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentCapturesByID] failed get paymentCaptures")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentCapturesByID(ctx context.Context, primaryID model.PaymentCapturesPrimaryID, paymentCaptures *model.PaymentCaptures, paymentCapturesUpdateFields ...PaymentCapturesUpdateField) (err error) {
	exists, err := repo.IsExistPaymentCapturesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentCaptures] failed checking paymentCaptures whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentCaptures with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentCaptures == nil {
		if len(paymentCapturesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentCapturesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentCaptures = &model.PaymentCaptures{}
	}
	var (
		defaultPaymentCapturesUpdateFields = defaultPaymentCapturesUpdateFields(*paymentCaptures)
		tempUpdateField                    PaymentCapturesUpdateFieldList
		selectFields                       = NewPaymentCapturesSelectFields()
	)
	if len(paymentCapturesUpdateFields) > 0 {
		for _, updateField := range paymentCapturesUpdateFields {
			if updateField.paymentCapturesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentCapturesUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentCapturesCompositePrimaryKeyWhere([]model.PaymentCapturesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentCapturesCommand(defaultPaymentCapturesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentCapturesQueries.updatePaymentCaptures+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentCaptures] error when try to update paymentCaptures by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentCapturesByFilter(ctx context.Context, filter model.Filter, paymentCapturesUpdateFields ...PaymentCapturesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentCapturesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentCapturesUpdateFieldList
		selectFields = NewPaymentCapturesSelectFields()
	)
	for _, updateField := range paymentCapturesUpdateFields {
		if updateField.paymentCapturesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentCapturesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentCapturesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentCapturesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_captures\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentCapturesByFilter] error when try to update paymentCaptures by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentCapturesByFilter] failed get rows affected")
	}
	return
}

var (
	paymentCapturesQueries = struct {
		selectPaymentCaptures      string
		selectCountPaymentCaptures string
		deletePaymentCaptures      string
		updatePaymentCaptures      string
		insertPaymentCaptures      string
	}{
		selectPaymentCaptures:      "SELECT %s FROM \"payment_captures\"",
		selectCountPaymentCaptures: "SELECT COUNT(\"id\") FROM \"payment_captures\"",
		deletePaymentCaptures:      "DELETE FROM \"payment_captures\"",
		updatePaymentCaptures:      "UPDATE \"payment_captures\" SET %s ",
		insertPaymentCaptures:      "INSERT INTO \"payment_captures\" %s VALUES %s",
	}
)

type PaymentCapturesRepository interface {
	CreatePaymentCaptures(ctx context.Context, paymentCaptures *model.PaymentCaptures, fieldsInsert ...PaymentCapturesField) error
	BulkCreatePaymentCaptures(ctx context.Context, paymentCapturesList []*model.PaymentCaptures, fieldsInsert ...PaymentCapturesField) error
	ResolvePaymentCaptures(ctx context.Context, selectFields ...PaymentCapturesField) (model.PaymentCapturesList, error)
	ResolvePaymentCapturesByID(ctx context.Context, primaryID model.PaymentCapturesPrimaryID, selectFields ...PaymentCapturesField) (model.PaymentCaptures, error)
	UpdatePaymentCapturesByID(ctx context.Context, id model.PaymentCapturesPrimaryID, paymentCaptures *model.PaymentCaptures, paymentCapturesUpdateFields ...PaymentCapturesUpdateField) error
	UpdatePaymentCapturesByFilter(ctx context.Context, filter model.Filter, paymentCapturesUpdateFields ...PaymentCapturesUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentCaptures(ctx context.Context, paymentCapturesListMap map[model.PaymentCapturesPrimaryID]*model.PaymentCaptures, PaymentCapturessMapUpdateFieldsRequest map[model.PaymentCapturesPrimaryID]PaymentCapturesUpdateFieldList) (err error)
	DeletePaymentCapturesByID(ctx context.Context, id model.PaymentCapturesPrimaryID) error
	BulkDeletePaymentCapturesByIDs(ctx context.Context, ids []model.PaymentCapturesPrimaryID) error
	ResolvePaymentCapturesByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentCapturesFilterResult, err error)
	IsExistPaymentCapturesByIDs(ctx context.Context, ids []model.PaymentCapturesPrimaryID) (exists bool, notFoundIds []model.PaymentCapturesPrimaryID, err error)
	IsExistPaymentCapturesByID(ctx context.Context, id model.PaymentCapturesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
