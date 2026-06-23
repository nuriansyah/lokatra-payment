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

func composeInsertFieldsAndParamsPayoutAttempts(payoutAttemptsList []model.PayoutAttempts, fieldsInsert ...PayoutAttemptsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPayoutAttemptsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, payoutAttempts := range payoutAttemptsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, payoutAttempts.Id)
			case selectField.PayoutId():
				args = append(args, payoutAttempts.PayoutId)
			case selectField.AttemptNo():
				args = append(args, payoutAttempts.AttemptNo)
			case selectField.AttemptType():
				args = append(args, payoutAttempts.AttemptType)
			case selectField.ProviderAccountId():
				args = append(args, payoutAttempts.ProviderAccountId)
			case selectField.Amount():
				args = append(args, payoutAttempts.Amount)
			case selectField.CurrencyCode():
				args = append(args, payoutAttempts.CurrencyCode)
			case selectField.AttemptStatus():
				args = append(args, payoutAttempts.AttemptStatus)
			case selectField.ProviderPayoutRef():
				args = append(args, payoutAttempts.ProviderPayoutRef)
			case selectField.FailureCode():
				args = append(args, payoutAttempts.FailureCode)
			case selectField.FailureReason():
				args = append(args, payoutAttempts.FailureReason)
			case selectField.RawRequest():
				args = append(args, payoutAttempts.RawRequest)
			case selectField.RawResponse():
				args = append(args, payoutAttempts.RawResponse)
			case selectField.Metadata():
				args = append(args, payoutAttempts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, payoutAttempts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, payoutAttempts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, payoutAttempts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, payoutAttempts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, payoutAttempts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, payoutAttempts.MetaDeletedBy)

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

func composePayoutAttemptsCompositePrimaryKeyWhere(primaryIDs []model.PayoutAttemptsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payout_attempts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPayoutAttemptsSelectFields() string {
	fields := NewPayoutAttemptsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePayoutAttemptsSelectFields(selectFields ...PayoutAttemptsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PayoutAttemptsField string
type PayoutAttemptsFieldList []PayoutAttemptsField

type PayoutAttemptsSelectFields struct {
}

func (ss PayoutAttemptsSelectFields) Id() PayoutAttemptsField {
	return PayoutAttemptsField("id")
}

func (ss PayoutAttemptsSelectFields) PayoutId() PayoutAttemptsField {
	return PayoutAttemptsField("payout_id")
}

func (ss PayoutAttemptsSelectFields) AttemptNo() PayoutAttemptsField {
	return PayoutAttemptsField("attempt_no")
}

func (ss PayoutAttemptsSelectFields) AttemptType() PayoutAttemptsField {
	return PayoutAttemptsField("attempt_type")
}

func (ss PayoutAttemptsSelectFields) ProviderAccountId() PayoutAttemptsField {
	return PayoutAttemptsField("provider_account_id")
}

func (ss PayoutAttemptsSelectFields) Amount() PayoutAttemptsField {
	return PayoutAttemptsField("amount")
}

func (ss PayoutAttemptsSelectFields) CurrencyCode() PayoutAttemptsField {
	return PayoutAttemptsField("currency_code")
}

func (ss PayoutAttemptsSelectFields) AttemptStatus() PayoutAttemptsField {
	return PayoutAttemptsField("attempt_status")
}

func (ss PayoutAttemptsSelectFields) ProviderPayoutRef() PayoutAttemptsField {
	return PayoutAttemptsField("provider_payout_ref")
}

func (ss PayoutAttemptsSelectFields) FailureCode() PayoutAttemptsField {
	return PayoutAttemptsField("failure_code")
}

func (ss PayoutAttemptsSelectFields) FailureReason() PayoutAttemptsField {
	return PayoutAttemptsField("failure_reason")
}

func (ss PayoutAttemptsSelectFields) RawRequest() PayoutAttemptsField {
	return PayoutAttemptsField("raw_request")
}

func (ss PayoutAttemptsSelectFields) RawResponse() PayoutAttemptsField {
	return PayoutAttemptsField("raw_response")
}

func (ss PayoutAttemptsSelectFields) Metadata() PayoutAttemptsField {
	return PayoutAttemptsField("metadata")
}

func (ss PayoutAttemptsSelectFields) MetaCreatedAt() PayoutAttemptsField {
	return PayoutAttemptsField("meta_created_at")
}

func (ss PayoutAttemptsSelectFields) MetaCreatedBy() PayoutAttemptsField {
	return PayoutAttemptsField("meta_created_by")
}

func (ss PayoutAttemptsSelectFields) MetaUpdatedAt() PayoutAttemptsField {
	return PayoutAttemptsField("meta_updated_at")
}

func (ss PayoutAttemptsSelectFields) MetaUpdatedBy() PayoutAttemptsField {
	return PayoutAttemptsField("meta_updated_by")
}

func (ss PayoutAttemptsSelectFields) MetaDeletedAt() PayoutAttemptsField {
	return PayoutAttemptsField("meta_deleted_at")
}

func (ss PayoutAttemptsSelectFields) MetaDeletedBy() PayoutAttemptsField {
	return PayoutAttemptsField("meta_deleted_by")
}

func (ss PayoutAttemptsSelectFields) All() PayoutAttemptsFieldList {
	return []PayoutAttemptsField{
		ss.Id(),
		ss.PayoutId(),
		ss.AttemptNo(),
		ss.AttemptType(),
		ss.ProviderAccountId(),
		ss.Amount(),
		ss.CurrencyCode(),
		ss.AttemptStatus(),
		ss.ProviderPayoutRef(),
		ss.FailureCode(),
		ss.FailureReason(),
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

func NewPayoutAttemptsSelectFields() PayoutAttemptsSelectFields {
	return PayoutAttemptsSelectFields{}
}

type PayoutAttemptsUpdateFieldOption struct {
	useIncrement bool
}
type PayoutAttemptsUpdateField struct {
	payoutAttemptsField PayoutAttemptsField
	opt                 PayoutAttemptsUpdateFieldOption
	value               interface{}
}
type PayoutAttemptsUpdateFieldList []PayoutAttemptsUpdateField

func defaultPayoutAttemptsUpdateFieldOption() PayoutAttemptsUpdateFieldOption {
	return PayoutAttemptsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPayoutAttemptsOption(useIncrement bool) func(*PayoutAttemptsUpdateFieldOption) {
	return func(pcufo *PayoutAttemptsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPayoutAttemptsUpdateField(field PayoutAttemptsField, val interface{}, opts ...func(*PayoutAttemptsUpdateFieldOption)) PayoutAttemptsUpdateField {
	defaultOpt := defaultPayoutAttemptsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PayoutAttemptsUpdateField{
		payoutAttemptsField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultPayoutAttemptsUpdateFields(payoutAttempts model.PayoutAttempts) (payoutAttemptsUpdateFieldList PayoutAttemptsUpdateFieldList) {
	selectFields := NewPayoutAttemptsSelectFields()
	payoutAttemptsUpdateFieldList = append(payoutAttemptsUpdateFieldList,
		NewPayoutAttemptsUpdateField(selectFields.Id(), payoutAttempts.Id),
		NewPayoutAttemptsUpdateField(selectFields.PayoutId(), payoutAttempts.PayoutId),
		NewPayoutAttemptsUpdateField(selectFields.AttemptNo(), payoutAttempts.AttemptNo),
		NewPayoutAttemptsUpdateField(selectFields.AttemptType(), payoutAttempts.AttemptType),
		NewPayoutAttemptsUpdateField(selectFields.ProviderAccountId(), payoutAttempts.ProviderAccountId),
		NewPayoutAttemptsUpdateField(selectFields.Amount(), payoutAttempts.Amount),
		NewPayoutAttemptsUpdateField(selectFields.CurrencyCode(), payoutAttempts.CurrencyCode),
		NewPayoutAttemptsUpdateField(selectFields.AttemptStatus(), payoutAttempts.AttemptStatus),
		NewPayoutAttemptsUpdateField(selectFields.ProviderPayoutRef(), payoutAttempts.ProviderPayoutRef),
		NewPayoutAttemptsUpdateField(selectFields.FailureCode(), payoutAttempts.FailureCode),
		NewPayoutAttemptsUpdateField(selectFields.FailureReason(), payoutAttempts.FailureReason),
		NewPayoutAttemptsUpdateField(selectFields.RawRequest(), payoutAttempts.RawRequest),
		NewPayoutAttemptsUpdateField(selectFields.RawResponse(), payoutAttempts.RawResponse),
		NewPayoutAttemptsUpdateField(selectFields.Metadata(), payoutAttempts.Metadata),
		NewPayoutAttemptsUpdateField(selectFields.MetaCreatedAt(), payoutAttempts.MetaCreatedAt),
		NewPayoutAttemptsUpdateField(selectFields.MetaCreatedBy(), payoutAttempts.MetaCreatedBy),
		NewPayoutAttemptsUpdateField(selectFields.MetaUpdatedAt(), payoutAttempts.MetaUpdatedAt),
		NewPayoutAttemptsUpdateField(selectFields.MetaUpdatedBy(), payoutAttempts.MetaUpdatedBy),
		NewPayoutAttemptsUpdateField(selectFields.MetaDeletedAt(), payoutAttempts.MetaDeletedAt),
		NewPayoutAttemptsUpdateField(selectFields.MetaDeletedBy(), payoutAttempts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPayoutAttemptsCommand(payoutAttemptsUpdateFieldList PayoutAttemptsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range payoutAttemptsUpdateFieldList {
		field := string(updateField.payoutAttemptsField)
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

func (repo *RepositoryImpl) BulkCreatePayoutAttempts(ctx context.Context, payoutAttemptsList []*model.PayoutAttempts, fieldsInsert ...PayoutAttemptsField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.PayoutAttemptsPrimaryID
		payoutAttemptsValueList []model.PayoutAttempts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPayoutAttemptsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, payoutAttempts := range payoutAttemptsList {

		primaryIds = append(primaryIds, payoutAttempts.ToPayoutAttemptsPrimaryID())

		payoutAttemptsValueList = append(payoutAttemptsValueList, *payoutAttempts)
	}

	_, notFoundIds, err := repo.IsExistPayoutAttemptsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutAttempts] failed checking payoutAttempts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PayoutAttemptsPrimaryID{}
		mapNotFoundIds := map[model.PayoutAttemptsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "payoutAttempts", fmt.Sprintf("payoutAttempts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPayoutAttempts(payoutAttemptsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(payoutAttemptsQueries.insertPayoutAttempts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutAttempts] failed exec create payoutAttempts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePayoutAttemptsByIDs(ctx context.Context, primaryIDs []model.PayoutAttemptsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPayoutAttemptsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutAttemptsByIDs] failed checking payoutAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutAttempts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_attempts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(payoutAttemptsQueries.deletePayoutAttempts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutAttemptsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutAttemptsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPayoutAttemptsByIDs(ctx context.Context, ids []model.PayoutAttemptsPrimaryID) (exists bool, notFoundIds []model.PayoutAttemptsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_attempts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(payoutAttemptsQueries.selectPayoutAttempts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutAttemptsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PayoutAttemptsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutAttemptsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PayoutAttemptsPrimaryID]bool{}
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

// BulkUpdatePayoutAttempts is used to bulk update payoutAttempts, by default it will update all field
// if want to update specific field, then fill payoutAttemptssMapUpdateFieldsRequest else please fill payoutAttemptssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePayoutAttempts(ctx context.Context, payoutAttemptssMap map[model.PayoutAttemptsPrimaryID]*model.PayoutAttempts, payoutAttemptssMapUpdateFieldsRequest map[model.PayoutAttemptsPrimaryID]PayoutAttemptsUpdateFieldList) (err error) {
	if len(payoutAttemptssMap) == 0 && len(payoutAttemptssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		payoutAttemptssMapUpdateField map[model.PayoutAttemptsPrimaryID]PayoutAttemptsUpdateFieldList = map[model.PayoutAttemptsPrimaryID]PayoutAttemptsUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(payoutAttemptssMap) > 0 {
		for id, payoutAttempts := range payoutAttemptssMap {
			if payoutAttempts == nil {
				log.Error().Err(err).Msg("[BulkUpdatePayoutAttempts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			payoutAttemptssMapUpdateField[id] = defaultPayoutAttemptsUpdateFields(*payoutAttempts)
		}
	} else {
		payoutAttemptssMapUpdateField = payoutAttemptssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePayoutAttemptsQuery(payoutAttemptssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPayoutAttemptsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutAttempts] failed checking payoutAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutAttempts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePayoutAttemptsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payout_attempts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutAttempts] failed exec query")
	}
	return
}

type PayoutAttemptsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPayoutAttemptsFieldParameter(param string, args ...interface{}) PayoutAttemptsFieldParameter {
	return PayoutAttemptsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePayoutAttemptsQuery(mapPayoutAttemptss map[model.PayoutAttemptsPrimaryID]PayoutAttemptsUpdateFieldList, asTableValues string) (primaryIDs []model.PayoutAttemptsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PayoutAttemptsPrimaryID]map[string]interface{}{}
	payoutAttemptsSelectFields := NewPayoutAttemptsSelectFields()
	for id, updateFields := range mapPayoutAttemptss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.payoutAttemptsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPayoutAttemptss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPayoutAttemptsFieldType(updateField.payoutAttemptsField)))
			args = append(args, fields[string(updateField.payoutAttemptsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.payoutAttemptsField))
		if updateField.payoutAttemptsField == payoutAttemptsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.payoutAttemptsField, asTableValues, updateField.payoutAttemptsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.payoutAttemptsField,
				"\"payout_attempts\"", updateField.payoutAttemptsField,
				asTableValues, updateField.payoutAttemptsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePayoutAttemptsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PayoutAttemptsPrimaryID, asTableValue string) (whereQry string) {
	payoutAttemptsSelectFields := NewPayoutAttemptsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payout_attempts\".\"id\" = %s.\"id\"::"+GetPayoutAttemptsFieldType(payoutAttemptsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPayoutAttemptsFieldType(payoutAttemptsField PayoutAttemptsField) string {
	selectPayoutAttemptsFields := NewPayoutAttemptsSelectFields()
	switch payoutAttemptsField {

	case selectPayoutAttemptsFields.Id():
		return "uuid"

	case selectPayoutAttemptsFields.PayoutId():
		return "uuid"

	case selectPayoutAttemptsFields.AttemptNo():
		return "int4"

	case selectPayoutAttemptsFields.AttemptType():
		return "payout_attempts_attempt_type_enum"

	case selectPayoutAttemptsFields.ProviderAccountId():
		return "uuid"

	case selectPayoutAttemptsFields.Amount():
		return "numeric"

	case selectPayoutAttemptsFields.CurrencyCode():
		return "text"

	case selectPayoutAttemptsFields.AttemptStatus():
		return "payout_attempts_attempt_status_enum"

	case selectPayoutAttemptsFields.ProviderPayoutRef():
		return "text"

	case selectPayoutAttemptsFields.FailureCode():
		return "text"

	case selectPayoutAttemptsFields.FailureReason():
		return "text"

	case selectPayoutAttemptsFields.RawRequest():
		return "jsonb"

	case selectPayoutAttemptsFields.RawResponse():
		return "jsonb"

	case selectPayoutAttemptsFields.Metadata():
		return "jsonb"

	case selectPayoutAttemptsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPayoutAttemptsFields.MetaCreatedBy():
		return "uuid"

	case selectPayoutAttemptsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPayoutAttemptsFields.MetaUpdatedBy():
		return "uuid"

	case selectPayoutAttemptsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPayoutAttemptsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePayoutAttempts(ctx context.Context, payoutAttempts *model.PayoutAttempts, fieldsInsert ...PayoutAttemptsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPayoutAttemptsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PayoutAttemptsPrimaryID{
		Id: payoutAttempts.Id,
	}
	exists, err := repo.IsExistPayoutAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutAttempts] failed checking payoutAttempts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "payoutAttempts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPayoutAttempts([]model.PayoutAttempts{*payoutAttempts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(payoutAttemptsQueries.insertPayoutAttempts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutAttempts] failed exec create payoutAttempts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePayoutAttemptsByID(ctx context.Context, primaryID model.PayoutAttemptsPrimaryID) (err error) {
	exists, err := repo.IsExistPayoutAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutAttemptsByID] failed checking payoutAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutAttempts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePayoutAttemptsCompositePrimaryKeyWhere([]model.PayoutAttemptsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(payoutAttemptsQueries.deletePayoutAttempts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutAttemptsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutAttemptsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutAttemptsFilterResult, err error) {
	query, args, err := composePayoutAttemptsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutAttemptsByFilter] failed compose payoutAttempts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutAttemptsByFilter] failed get payoutAttempts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePayoutAttemptsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PayoutAttemptsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePayoutAttemptsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePayoutAttemptsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePayoutAttemptsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPayoutAttemptsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["payout_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_id\"")
			selectedColumns["payout_id"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_no\"")
			selectedColumns["attempt_no"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_type\"")
			selectedColumns["attempt_type"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_status\"")
			selectedColumns["attempt_status"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_payout_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_payout_ref\"")
			selectedColumns["provider_payout_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_code\"")
			selectedColumns["failure_code"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_reason\"")
			selectedColumns["failure_reason"] = struct{}{}
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

type payoutAttemptsFilterPlaceholder struct {
	index int
}

func (p *payoutAttemptsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePayoutAttemptsFilterPredicate(filterField model.FilterField, placeholders *payoutAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPayoutAttemptsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePayoutAttemptsFilterSQLExpr(spec)
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

func composePayoutAttemptsFilterGroup(group model.FilterGroup, placeholders *payoutAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePayoutAttemptsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePayoutAttemptsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePayoutAttemptsFilterWhereQueries(filter model.Filter, placeholders *payoutAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePayoutAttemptsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePayoutAttemptsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePayoutAttemptsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePayoutAttemptsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePayoutAttemptsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePayoutAttemptsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := payoutAttemptsFilterPlaceholder{index: 1}
	whereQueries, err := composePayoutAttemptsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPayoutAttemptsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePayoutAttemptsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePayoutAttemptsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payout_attempts\" base%s", strings.Join(selectColumns, ","), composePayoutAttemptsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPayoutAttemptsByID(ctx context.Context, primaryID model.PayoutAttemptsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePayoutAttemptsCompositePrimaryKeyWhere([]model.PayoutAttemptsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", payoutAttemptsQueries.selectCountPayoutAttempts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutAttemptsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutAttempts(ctx context.Context, selectFields ...PayoutAttemptsField) (payoutAttemptsList model.PayoutAttemptsList, err error) {
	var (
		defaultPayoutAttemptsSelectFields = defaultPayoutAttemptsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutAttemptsSelectFields = composePayoutAttemptsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(payoutAttemptsQueries.selectPayoutAttempts, defaultPayoutAttemptsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &payoutAttemptsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutAttempts] failed get payoutAttempts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutAttemptsByID(ctx context.Context, primaryID model.PayoutAttemptsPrimaryID, selectFields ...PayoutAttemptsField) (payoutAttempts model.PayoutAttempts, err error) {
	var (
		defaultPayoutAttemptsSelectFields = defaultPayoutAttemptsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutAttemptsSelectFields = composePayoutAttemptsSelectFields(selectFields...)
	}
	whereQry, params := composePayoutAttemptsCompositePrimaryKeyWhere([]model.PayoutAttemptsPrimaryID{primaryID})
	query := fmt.Sprintf(payoutAttemptsQueries.selectPayoutAttempts+" WHERE "+whereQry, defaultPayoutAttemptsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &payoutAttempts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("payoutAttempts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePayoutAttemptsByID] failed get payoutAttempts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePayoutAttemptsByID(ctx context.Context, primaryID model.PayoutAttemptsPrimaryID, payoutAttempts *model.PayoutAttempts, payoutAttemptsUpdateFields ...PayoutAttemptsUpdateField) (err error) {
	exists, err := repo.IsExistPayoutAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutAttempts] failed checking payoutAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutAttempts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if payoutAttempts == nil {
		if len(payoutAttemptsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePayoutAttemptsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		payoutAttempts = &model.PayoutAttempts{}
	}
	var (
		defaultPayoutAttemptsUpdateFields = defaultPayoutAttemptsUpdateFields(*payoutAttempts)
		tempUpdateField                   PayoutAttemptsUpdateFieldList
		selectFields                      = NewPayoutAttemptsSelectFields()
	)
	if len(payoutAttemptsUpdateFields) > 0 {
		for _, updateField := range payoutAttemptsUpdateFields {
			if updateField.payoutAttemptsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPayoutAttemptsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePayoutAttemptsCompositePrimaryKeyWhere([]model.PayoutAttemptsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPayoutAttemptsCommand(defaultPayoutAttemptsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(payoutAttemptsQueries.updatePayoutAttempts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutAttempts] error when try to update payoutAttempts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePayoutAttemptsByFilter(ctx context.Context, filter model.Filter, payoutAttemptsUpdateFields ...PayoutAttemptsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(payoutAttemptsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PayoutAttemptsUpdateFieldList
		selectFields = NewPayoutAttemptsSelectFields()
	)
	for _, updateField := range payoutAttemptsUpdateFields {
		if updateField.payoutAttemptsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPayoutAttemptsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := payoutAttemptsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePayoutAttemptsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payout_attempts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutAttemptsByFilter] error when try to update payoutAttempts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutAttemptsByFilter] failed get rows affected")
	}
	return
}

var (
	payoutAttemptsQueries = struct {
		selectPayoutAttempts      string
		selectCountPayoutAttempts string
		deletePayoutAttempts      string
		updatePayoutAttempts      string
		insertPayoutAttempts      string
	}{
		selectPayoutAttempts:      "SELECT %s FROM \"payout_attempts\"",
		selectCountPayoutAttempts: "SELECT COUNT(\"id\") FROM \"payout_attempts\"",
		deletePayoutAttempts:      "DELETE FROM \"payout_attempts\"",
		updatePayoutAttempts:      "UPDATE \"payout_attempts\" SET %s ",
		insertPayoutAttempts:      "INSERT INTO \"payout_attempts\" %s VALUES %s",
	}
)

type PayoutAttemptsRepository interface {
	CreatePayoutAttempts(ctx context.Context, payoutAttempts *model.PayoutAttempts, fieldsInsert ...PayoutAttemptsField) error
	BulkCreatePayoutAttempts(ctx context.Context, payoutAttemptsList []*model.PayoutAttempts, fieldsInsert ...PayoutAttemptsField) error
	ResolvePayoutAttempts(ctx context.Context, selectFields ...PayoutAttemptsField) (model.PayoutAttemptsList, error)
	ResolvePayoutAttemptsByID(ctx context.Context, primaryID model.PayoutAttemptsPrimaryID, selectFields ...PayoutAttemptsField) (model.PayoutAttempts, error)
	UpdatePayoutAttemptsByID(ctx context.Context, id model.PayoutAttemptsPrimaryID, payoutAttempts *model.PayoutAttempts, payoutAttemptsUpdateFields ...PayoutAttemptsUpdateField) error
	UpdatePayoutAttemptsByFilter(ctx context.Context, filter model.Filter, payoutAttemptsUpdateFields ...PayoutAttemptsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePayoutAttempts(ctx context.Context, payoutAttemptsListMap map[model.PayoutAttemptsPrimaryID]*model.PayoutAttempts, PayoutAttemptssMapUpdateFieldsRequest map[model.PayoutAttemptsPrimaryID]PayoutAttemptsUpdateFieldList) (err error)
	DeletePayoutAttemptsByID(ctx context.Context, id model.PayoutAttemptsPrimaryID) error
	BulkDeletePayoutAttemptsByIDs(ctx context.Context, ids []model.PayoutAttemptsPrimaryID) error
	ResolvePayoutAttemptsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutAttemptsFilterResult, err error)
	IsExistPayoutAttemptsByIDs(ctx context.Context, ids []model.PayoutAttemptsPrimaryID) (exists bool, notFoundIds []model.PayoutAttemptsPrimaryID, err error)
	IsExistPayoutAttemptsByID(ctx context.Context, id model.PayoutAttemptsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
