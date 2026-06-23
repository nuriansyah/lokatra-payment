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

func composeInsertFieldsAndParamsPaymentRouteCandidatesRuntime(paymentRouteCandidatesRuntimeList []model.PaymentRouteCandidatesRuntime, fieldsInsert ...PaymentRouteCandidatesRuntimeField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentRouteCandidatesRuntimeSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentRouteCandidatesRuntime := range paymentRouteCandidatesRuntimeList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentRouteCandidatesRuntime.Id)
			case selectField.ScopeType():
				args = append(args, paymentRouteCandidatesRuntime.ScopeType)
			case selectField.ScopeId():
				args = append(args, paymentRouteCandidatesRuntime.ScopeId)
			case selectField.MerchantId():
				args = append(args, paymentRouteCandidatesRuntime.MerchantId)
			case selectField.MethodCode():
				args = append(args, paymentRouteCandidatesRuntime.MethodCode)
			case selectField.ChannelCode():
				args = append(args, paymentRouteCandidatesRuntime.ChannelCode)
			case selectField.Currency():
				args = append(args, paymentRouteCandidatesRuntime.Currency)
			case selectField.MinAmount():
				args = append(args, paymentRouteCandidatesRuntime.MinAmount)
			case selectField.MaxAmount():
				args = append(args, paymentRouteCandidatesRuntime.MaxAmount)
			case selectField.ProviderAccountId():
				args = append(args, paymentRouteCandidatesRuntime.ProviderAccountId)
			case selectField.ProviderMethodCode():
				args = append(args, paymentRouteCandidatesRuntime.ProviderMethodCode)
			case selectField.ProviderChannelCode():
				args = append(args, paymentRouteCandidatesRuntime.ProviderChannelCode)
			case selectField.Priority():
				args = append(args, paymentRouteCandidatesRuntime.Priority)
			case selectField.IsFallback():
				args = append(args, paymentRouteCandidatesRuntime.IsFallback)
			case selectField.TrafficWeight():
				args = append(args, paymentRouteCandidatesRuntime.TrafficWeight)
			case selectField.TimeoutMs():
				args = append(args, paymentRouteCandidatesRuntime.TimeoutMs)
			case selectField.MaxAttempts():
				args = append(args, paymentRouteCandidatesRuntime.MaxAttempts)
			case selectField.IsEnabled():
				args = append(args, paymentRouteCandidatesRuntime.IsEnabled)
			case selectField.ConditionExpr():
				args = append(args, paymentRouteCandidatesRuntime.ConditionExpr)
			case selectField.Metadata():
				args = append(args, paymentRouteCandidatesRuntime.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentRouteCandidatesRuntime.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentRouteCandidatesRuntime.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentRouteCandidatesRuntime.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentRouteCandidatesRuntime.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentRouteCandidatesRuntime.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentRouteCandidatesRuntime.MetaDeletedBy)

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

func composePaymentRouteCandidatesRuntimeCompositePrimaryKeyWhere(primaryIDs []model.PaymentRouteCandidatesRuntimePrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_route_candidates_runtime\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentRouteCandidatesRuntimeSelectFields() string {
	fields := NewPaymentRouteCandidatesRuntimeSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentRouteCandidatesRuntimeSelectFields(selectFields ...PaymentRouteCandidatesRuntimeField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentRouteCandidatesRuntimeField string
type PaymentRouteCandidatesRuntimeFieldList []PaymentRouteCandidatesRuntimeField

type PaymentRouteCandidatesRuntimeSelectFields struct {
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) Id() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("id")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) ScopeType() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("scope_type")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) ScopeId() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("scope_id")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MerchantId() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("merchant_id")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MethodCode() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("method_code")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) ChannelCode() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("channel_code")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) Currency() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("currency")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MinAmount() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("min_amount")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MaxAmount() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("max_amount")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) ProviderAccountId() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("provider_account_id")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) ProviderMethodCode() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("provider_method_code")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) ProviderChannelCode() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("provider_channel_code")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) Priority() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("priority")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) IsFallback() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("is_fallback")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) TrafficWeight() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("traffic_weight")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) TimeoutMs() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("timeout_ms")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MaxAttempts() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("max_attempts")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) IsEnabled() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("is_enabled")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) ConditionExpr() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("condition_expr")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) Metadata() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("metadata")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MetaCreatedAt() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("meta_created_at")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MetaCreatedBy() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("meta_created_by")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MetaUpdatedAt() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("meta_updated_at")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MetaUpdatedBy() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("meta_updated_by")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MetaDeletedAt() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("meta_deleted_at")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) MetaDeletedBy() PaymentRouteCandidatesRuntimeField {
	return PaymentRouteCandidatesRuntimeField("meta_deleted_by")
}

func (ss PaymentRouteCandidatesRuntimeSelectFields) All() PaymentRouteCandidatesRuntimeFieldList {
	return []PaymentRouteCandidatesRuntimeField{
		ss.Id(),
		ss.ScopeType(),
		ss.ScopeId(),
		ss.MerchantId(),
		ss.MethodCode(),
		ss.ChannelCode(),
		ss.Currency(),
		ss.MinAmount(),
		ss.MaxAmount(),
		ss.ProviderAccountId(),
		ss.ProviderMethodCode(),
		ss.ProviderChannelCode(),
		ss.Priority(),
		ss.IsFallback(),
		ss.TrafficWeight(),
		ss.TimeoutMs(),
		ss.MaxAttempts(),
		ss.IsEnabled(),
		ss.ConditionExpr(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentRouteCandidatesRuntimeSelectFields() PaymentRouteCandidatesRuntimeSelectFields {
	return PaymentRouteCandidatesRuntimeSelectFields{}
}

type PaymentRouteCandidatesRuntimeUpdateFieldOption struct {
	useIncrement bool
}
type PaymentRouteCandidatesRuntimeUpdateField struct {
	paymentRouteCandidatesRuntimeField PaymentRouteCandidatesRuntimeField
	opt                                PaymentRouteCandidatesRuntimeUpdateFieldOption
	value                              interface{}
}
type PaymentRouteCandidatesRuntimeUpdateFieldList []PaymentRouteCandidatesRuntimeUpdateField

func defaultPaymentRouteCandidatesRuntimeUpdateFieldOption() PaymentRouteCandidatesRuntimeUpdateFieldOption {
	return PaymentRouteCandidatesRuntimeUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentRouteCandidatesRuntimeOption(useIncrement bool) func(*PaymentRouteCandidatesRuntimeUpdateFieldOption) {
	return func(pcufo *PaymentRouteCandidatesRuntimeUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentRouteCandidatesRuntimeUpdateField(field PaymentRouteCandidatesRuntimeField, val interface{}, opts ...func(*PaymentRouteCandidatesRuntimeUpdateFieldOption)) PaymentRouteCandidatesRuntimeUpdateField {
	defaultOpt := defaultPaymentRouteCandidatesRuntimeUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentRouteCandidatesRuntimeUpdateField{
		paymentRouteCandidatesRuntimeField: field,
		value:                              val,
		opt:                                defaultOpt,
	}
}
func defaultPaymentRouteCandidatesRuntimeUpdateFields(paymentRouteCandidatesRuntime model.PaymentRouteCandidatesRuntime) (paymentRouteCandidatesRuntimeUpdateFieldList PaymentRouteCandidatesRuntimeUpdateFieldList) {
	selectFields := NewPaymentRouteCandidatesRuntimeSelectFields()
	paymentRouteCandidatesRuntimeUpdateFieldList = append(paymentRouteCandidatesRuntimeUpdateFieldList,
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.Id(), paymentRouteCandidatesRuntime.Id),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.ScopeType(), paymentRouteCandidatesRuntime.ScopeType),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.ScopeId(), paymentRouteCandidatesRuntime.ScopeId),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MerchantId(), paymentRouteCandidatesRuntime.MerchantId),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MethodCode(), paymentRouteCandidatesRuntime.MethodCode),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.ChannelCode(), paymentRouteCandidatesRuntime.ChannelCode),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.Currency(), paymentRouteCandidatesRuntime.Currency),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MinAmount(), paymentRouteCandidatesRuntime.MinAmount),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MaxAmount(), paymentRouteCandidatesRuntime.MaxAmount),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.ProviderAccountId(), paymentRouteCandidatesRuntime.ProviderAccountId),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.ProviderMethodCode(), paymentRouteCandidatesRuntime.ProviderMethodCode),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.ProviderChannelCode(), paymentRouteCandidatesRuntime.ProviderChannelCode),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.Priority(), paymentRouteCandidatesRuntime.Priority),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.IsFallback(), paymentRouteCandidatesRuntime.IsFallback),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.TrafficWeight(), paymentRouteCandidatesRuntime.TrafficWeight),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.TimeoutMs(), paymentRouteCandidatesRuntime.TimeoutMs),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MaxAttempts(), paymentRouteCandidatesRuntime.MaxAttempts),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.IsEnabled(), paymentRouteCandidatesRuntime.IsEnabled),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.ConditionExpr(), paymentRouteCandidatesRuntime.ConditionExpr),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.Metadata(), paymentRouteCandidatesRuntime.Metadata),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MetaCreatedAt(), paymentRouteCandidatesRuntime.MetaCreatedAt),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MetaCreatedBy(), paymentRouteCandidatesRuntime.MetaCreatedBy),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MetaUpdatedAt(), paymentRouteCandidatesRuntime.MetaUpdatedAt),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MetaUpdatedBy(), paymentRouteCandidatesRuntime.MetaUpdatedBy),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MetaDeletedAt(), paymentRouteCandidatesRuntime.MetaDeletedAt),
		NewPaymentRouteCandidatesRuntimeUpdateField(selectFields.MetaDeletedBy(), paymentRouteCandidatesRuntime.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentRouteCandidatesRuntimeCommand(paymentRouteCandidatesRuntimeUpdateFieldList PaymentRouteCandidatesRuntimeUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentRouteCandidatesRuntimeUpdateFieldList {
		field := string(updateField.paymentRouteCandidatesRuntimeField)
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

func (repo *RepositoryImpl) BulkCreatePaymentRouteCandidatesRuntime(ctx context.Context, paymentRouteCandidatesRuntimeList []*model.PaymentRouteCandidatesRuntime, fieldsInsert ...PaymentRouteCandidatesRuntimeField) (err error) {
	var (
		fieldsStr                              string
		valueListStr                           []string
		argsList                               []interface{}
		primaryIds                             []model.PaymentRouteCandidatesRuntimePrimaryID
		paymentRouteCandidatesRuntimeValueList []model.PaymentRouteCandidatesRuntime
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentRouteCandidatesRuntimeSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentRouteCandidatesRuntime := range paymentRouteCandidatesRuntimeList {

		primaryIds = append(primaryIds, paymentRouteCandidatesRuntime.ToPaymentRouteCandidatesRuntimePrimaryID())

		paymentRouteCandidatesRuntimeValueList = append(paymentRouteCandidatesRuntimeValueList, *paymentRouteCandidatesRuntime)
	}

	_, notFoundIds, err := repo.IsExistPaymentRouteCandidatesRuntimeByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentRouteCandidatesRuntime] failed checking paymentRouteCandidatesRuntime whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentRouteCandidatesRuntimePrimaryID{}
		mapNotFoundIds := map[model.PaymentRouteCandidatesRuntimePrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentRouteCandidatesRuntime", fmt.Sprintf("paymentRouteCandidatesRuntime with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentRouteCandidatesRuntime(paymentRouteCandidatesRuntimeValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentRouteCandidatesRuntimeQueries.insertPaymentRouteCandidatesRuntime, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentRouteCandidatesRuntime] failed exec create paymentRouteCandidatesRuntime query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentRouteCandidatesRuntimeByIDs(ctx context.Context, primaryIDs []model.PaymentRouteCandidatesRuntimePrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentRouteCandidatesRuntimeByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRouteCandidatesRuntimeByIDs] failed checking paymentRouteCandidatesRuntime whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteCandidatesRuntime with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_route_candidates_runtime\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentRouteCandidatesRuntimeQueries.deletePaymentRouteCandidatesRuntime + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRouteCandidatesRuntimeByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRouteCandidatesRuntimeByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentRouteCandidatesRuntimeByIDs(ctx context.Context, ids []model.PaymentRouteCandidatesRuntimePrimaryID) (exists bool, notFoundIds []model.PaymentRouteCandidatesRuntimePrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_route_candidates_runtime\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentRouteCandidatesRuntimeQueries.selectPaymentRouteCandidatesRuntime, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRouteCandidatesRuntimeByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentRouteCandidatesRuntimePrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRouteCandidatesRuntimeByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentRouteCandidatesRuntimePrimaryID]bool{}
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

// BulkUpdatePaymentRouteCandidatesRuntime is used to bulk update paymentRouteCandidatesRuntime, by default it will update all field
// if want to update specific field, then fill paymentRouteCandidatesRuntimesMapUpdateFieldsRequest else please fill paymentRouteCandidatesRuntimesMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentRouteCandidatesRuntime(ctx context.Context, paymentRouteCandidatesRuntimesMap map[model.PaymentRouteCandidatesRuntimePrimaryID]*model.PaymentRouteCandidatesRuntime, paymentRouteCandidatesRuntimesMapUpdateFieldsRequest map[model.PaymentRouteCandidatesRuntimePrimaryID]PaymentRouteCandidatesRuntimeUpdateFieldList) (err error) {
	if len(paymentRouteCandidatesRuntimesMap) == 0 && len(paymentRouteCandidatesRuntimesMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentRouteCandidatesRuntimesMapUpdateField map[model.PaymentRouteCandidatesRuntimePrimaryID]PaymentRouteCandidatesRuntimeUpdateFieldList = map[model.PaymentRouteCandidatesRuntimePrimaryID]PaymentRouteCandidatesRuntimeUpdateFieldList{}
		asTableValues                                string                                                                                        = "myvalues"
	)

	if len(paymentRouteCandidatesRuntimesMap) > 0 {
		for id, paymentRouteCandidatesRuntime := range paymentRouteCandidatesRuntimesMap {
			if paymentRouteCandidatesRuntime == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentRouteCandidatesRuntime] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentRouteCandidatesRuntimesMapUpdateField[id] = defaultPaymentRouteCandidatesRuntimeUpdateFields(*paymentRouteCandidatesRuntime)
		}
	} else {
		paymentRouteCandidatesRuntimesMapUpdateField = paymentRouteCandidatesRuntimesMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentRouteCandidatesRuntimeQuery(paymentRouteCandidatesRuntimesMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentRouteCandidatesRuntimeByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentRouteCandidatesRuntime] failed checking paymentRouteCandidatesRuntime whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteCandidatesRuntime with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentRouteCandidatesRuntimeCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_route_candidates_runtime\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentRouteCandidatesRuntime] failed exec query")
	}
	return
}

type PaymentRouteCandidatesRuntimeFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentRouteCandidatesRuntimeFieldParameter(param string, args ...interface{}) PaymentRouteCandidatesRuntimeFieldParameter {
	return PaymentRouteCandidatesRuntimeFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentRouteCandidatesRuntimeQuery(mapPaymentRouteCandidatesRuntimes map[model.PaymentRouteCandidatesRuntimePrimaryID]PaymentRouteCandidatesRuntimeUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentRouteCandidatesRuntimePrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentRouteCandidatesRuntimePrimaryID]map[string]interface{}{}
	paymentRouteCandidatesRuntimeSelectFields := NewPaymentRouteCandidatesRuntimeSelectFields()
	for id, updateFields := range mapPaymentRouteCandidatesRuntimes {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentRouteCandidatesRuntimeField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentRouteCandidatesRuntimes[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentRouteCandidatesRuntimeFieldType(updateField.paymentRouteCandidatesRuntimeField)))
			args = append(args, fields[string(updateField.paymentRouteCandidatesRuntimeField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentRouteCandidatesRuntimeField))
		if updateField.paymentRouteCandidatesRuntimeField == paymentRouteCandidatesRuntimeSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentRouteCandidatesRuntimeField, asTableValues, updateField.paymentRouteCandidatesRuntimeField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentRouteCandidatesRuntimeField,
				"\"payment_route_candidates_runtime\"", updateField.paymentRouteCandidatesRuntimeField,
				asTableValues, updateField.paymentRouteCandidatesRuntimeField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentRouteCandidatesRuntimeCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentRouteCandidatesRuntimePrimaryID, asTableValue string) (whereQry string) {
	paymentRouteCandidatesRuntimeSelectFields := NewPaymentRouteCandidatesRuntimeSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_route_candidates_runtime\".\"id\" = %s.\"id\"::"+GetPaymentRouteCandidatesRuntimeFieldType(paymentRouteCandidatesRuntimeSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentRouteCandidatesRuntimeFieldType(paymentRouteCandidatesRuntimeField PaymentRouteCandidatesRuntimeField) string {
	selectPaymentRouteCandidatesRuntimeFields := NewPaymentRouteCandidatesRuntimeSelectFields()
	switch paymentRouteCandidatesRuntimeField {

	case selectPaymentRouteCandidatesRuntimeFields.Id():
		return "uuid"

	case selectPaymentRouteCandidatesRuntimeFields.ScopeType():
		return "scope_type_enum"

	case selectPaymentRouteCandidatesRuntimeFields.ScopeId():
		return "uuid"

	case selectPaymentRouteCandidatesRuntimeFields.MerchantId():
		return "uuid"

	case selectPaymentRouteCandidatesRuntimeFields.MethodCode():
		return "text"

	case selectPaymentRouteCandidatesRuntimeFields.ChannelCode():
		return "text"

	case selectPaymentRouteCandidatesRuntimeFields.Currency():
		return "text"

	case selectPaymentRouteCandidatesRuntimeFields.MinAmount():
		return "numeric"

	case selectPaymentRouteCandidatesRuntimeFields.MaxAmount():
		return "numeric"

	case selectPaymentRouteCandidatesRuntimeFields.ProviderAccountId():
		return "uuid"

	case selectPaymentRouteCandidatesRuntimeFields.ProviderMethodCode():
		return "text"

	case selectPaymentRouteCandidatesRuntimeFields.ProviderChannelCode():
		return "text"

	case selectPaymentRouteCandidatesRuntimeFields.Priority():
		return "int4"

	case selectPaymentRouteCandidatesRuntimeFields.IsFallback():
		return "bool"

	case selectPaymentRouteCandidatesRuntimeFields.TrafficWeight():
		return "int4"

	case selectPaymentRouteCandidatesRuntimeFields.TimeoutMs():
		return "int4"

	case selectPaymentRouteCandidatesRuntimeFields.MaxAttempts():
		return "int4"

	case selectPaymentRouteCandidatesRuntimeFields.IsEnabled():
		return "bool"

	case selectPaymentRouteCandidatesRuntimeFields.ConditionExpr():
		return "jsonb"

	case selectPaymentRouteCandidatesRuntimeFields.Metadata():
		return "jsonb"

	case selectPaymentRouteCandidatesRuntimeFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentRouteCandidatesRuntimeFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentRouteCandidatesRuntimeFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentRouteCandidatesRuntimeFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentRouteCandidatesRuntimeFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentRouteCandidatesRuntimeFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentRouteCandidatesRuntime(ctx context.Context, paymentRouteCandidatesRuntime *model.PaymentRouteCandidatesRuntime, fieldsInsert ...PaymentRouteCandidatesRuntimeField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentRouteCandidatesRuntimeSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentRouteCandidatesRuntimePrimaryID{
		Id: paymentRouteCandidatesRuntime.Id,
	}
	exists, err := repo.IsExistPaymentRouteCandidatesRuntimeByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentRouteCandidatesRuntime] failed checking paymentRouteCandidatesRuntime whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentRouteCandidatesRuntime", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentRouteCandidatesRuntime([]model.PaymentRouteCandidatesRuntime{*paymentRouteCandidatesRuntime}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentRouteCandidatesRuntimeQueries.insertPaymentRouteCandidatesRuntime, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentRouteCandidatesRuntime] failed exec create paymentRouteCandidatesRuntime query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentRouteCandidatesRuntimeByID(ctx context.Context, primaryID model.PaymentRouteCandidatesRuntimePrimaryID) (err error) {
	exists, err := repo.IsExistPaymentRouteCandidatesRuntimeByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentRouteCandidatesRuntimeByID] failed checking paymentRouteCandidatesRuntime whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteCandidatesRuntime with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentRouteCandidatesRuntimeCompositePrimaryKeyWhere([]model.PaymentRouteCandidatesRuntimePrimaryID{primaryID})
	commandQuery := paymentRouteCandidatesRuntimeQueries.deletePaymentRouteCandidatesRuntime + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentRouteCandidatesRuntimeByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRouteCandidatesRuntimeByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentRouteCandidatesRuntimeFilterResult, err error) {
	query, args, err := composePaymentRouteCandidatesRuntimeFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRouteCandidatesRuntimeByFilter] failed compose paymentRouteCandidatesRuntime filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRouteCandidatesRuntimeByFilter] failed get paymentRouteCandidatesRuntime by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentRouteCandidatesRuntimeFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentRouteCandidatesRuntimeFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentRouteCandidatesRuntimeFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentRouteCandidatesRuntimeSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentRouteCandidatesRuntimeFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 26 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 26+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["scope_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"scope_type\"")
			selectedColumns["scope_type"] = struct{}{}
		}
		if _, selected := selectedColumns["scope_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"scope_id\"")
			selectedColumns["scope_id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_id\"")
			selectedColumns["merchant_id"] = struct{}{}
		}
		if _, selected := selectedColumns["method_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_code\"")
			selectedColumns["method_code"] = struct{}{}
		}
		if _, selected := selectedColumns["channel_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"channel_code\"")
			selectedColumns["channel_code"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["min_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"min_amount\"")
			selectedColumns["min_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["max_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"max_amount\"")
			selectedColumns["max_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_method_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_method_code\"")
			selectedColumns["provider_method_code"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_channel_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_channel_code\"")
			selectedColumns["provider_channel_code"] = struct{}{}
		}
		if _, selected := selectedColumns["priority"]; !selected {
			selectColumns = append(selectColumns, "base.\"priority\"")
			selectedColumns["priority"] = struct{}{}
		}
		if _, selected := selectedColumns["is_fallback"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_fallback\"")
			selectedColumns["is_fallback"] = struct{}{}
		}
		if _, selected := selectedColumns["traffic_weight"]; !selected {
			selectColumns = append(selectColumns, "base.\"traffic_weight\"")
			selectedColumns["traffic_weight"] = struct{}{}
		}
		if _, selected := selectedColumns["timeout_ms"]; !selected {
			selectColumns = append(selectColumns, "base.\"timeout_ms\"")
			selectedColumns["timeout_ms"] = struct{}{}
		}
		if _, selected := selectedColumns["max_attempts"]; !selected {
			selectColumns = append(selectColumns, "base.\"max_attempts\"")
			selectedColumns["max_attempts"] = struct{}{}
		}
		if _, selected := selectedColumns["is_enabled"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_enabled\"")
			selectedColumns["is_enabled"] = struct{}{}
		}
		if _, selected := selectedColumns["condition_expr"]; !selected {
			selectColumns = append(selectColumns, "base.\"condition_expr\"")
			selectedColumns["condition_expr"] = struct{}{}
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

type paymentRouteCandidatesRuntimeFilterPlaceholder struct {
	index int
}

func (p *paymentRouteCandidatesRuntimeFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentRouteCandidatesRuntimeFilterPredicate(filterField model.FilterField, placeholders *paymentRouteCandidatesRuntimeFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentRouteCandidatesRuntimeFilterSQLExpr(spec)
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

func composePaymentRouteCandidatesRuntimeFilterGroup(group model.FilterGroup, placeholders *paymentRouteCandidatesRuntimeFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentRouteCandidatesRuntimeFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentRouteCandidatesRuntimeFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentRouteCandidatesRuntimeFilterWhereQueries(filter model.Filter, placeholders *paymentRouteCandidatesRuntimeFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentRouteCandidatesRuntimeFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentRouteCandidatesRuntimeFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentRouteCandidatesRuntimeFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentRouteCandidatesRuntimeFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentRouteCandidatesRuntimeSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentRouteCandidatesRuntimeFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentRouteCandidatesRuntimeFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentRouteCandidatesRuntimeFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentRouteCandidatesRuntimeFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentRouteCandidatesRuntimeFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentRouteCandidatesRuntimeSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_route_candidates_runtime\" base%s", strings.Join(selectColumns, ","), composePaymentRouteCandidatesRuntimeFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentRouteCandidatesRuntimeByID(ctx context.Context, primaryID model.PaymentRouteCandidatesRuntimePrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentRouteCandidatesRuntimeCompositePrimaryKeyWhere([]model.PaymentRouteCandidatesRuntimePrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentRouteCandidatesRuntimeQueries.selectCountPaymentRouteCandidatesRuntime, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRouteCandidatesRuntimeByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRouteCandidatesRuntime(ctx context.Context, selectFields ...PaymentRouteCandidatesRuntimeField) (paymentRouteCandidatesRuntimeList model.PaymentRouteCandidatesRuntimeList, err error) {
	var (
		defaultPaymentRouteCandidatesRuntimeSelectFields = defaultPaymentRouteCandidatesRuntimeSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentRouteCandidatesRuntimeSelectFields = composePaymentRouteCandidatesRuntimeSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentRouteCandidatesRuntimeQueries.selectPaymentRouteCandidatesRuntime, defaultPaymentRouteCandidatesRuntimeSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentRouteCandidatesRuntimeList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRouteCandidatesRuntime] failed get paymentRouteCandidatesRuntime list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRouteCandidatesRuntimeByID(ctx context.Context, primaryID model.PaymentRouteCandidatesRuntimePrimaryID, selectFields ...PaymentRouteCandidatesRuntimeField) (paymentRouteCandidatesRuntime model.PaymentRouteCandidatesRuntime, err error) {
	var (
		defaultPaymentRouteCandidatesRuntimeSelectFields = defaultPaymentRouteCandidatesRuntimeSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentRouteCandidatesRuntimeSelectFields = composePaymentRouteCandidatesRuntimeSelectFields(selectFields...)
	}
	whereQry, params := composePaymentRouteCandidatesRuntimeCompositePrimaryKeyWhere([]model.PaymentRouteCandidatesRuntimePrimaryID{primaryID})
	query := fmt.Sprintf(paymentRouteCandidatesRuntimeQueries.selectPaymentRouteCandidatesRuntime+" WHERE "+whereQry, defaultPaymentRouteCandidatesRuntimeSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentRouteCandidatesRuntime, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentRouteCandidatesRuntime with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentRouteCandidatesRuntimeByID] failed get paymentRouteCandidatesRuntime")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentRouteCandidatesRuntimeByID(ctx context.Context, primaryID model.PaymentRouteCandidatesRuntimePrimaryID, paymentRouteCandidatesRuntime *model.PaymentRouteCandidatesRuntime, paymentRouteCandidatesRuntimeUpdateFields ...PaymentRouteCandidatesRuntimeUpdateField) (err error) {
	exists, err := repo.IsExistPaymentRouteCandidatesRuntimeByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteCandidatesRuntime] failed checking paymentRouteCandidatesRuntime whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteCandidatesRuntime with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentRouteCandidatesRuntime == nil {
		if len(paymentRouteCandidatesRuntimeUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentRouteCandidatesRuntimeByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentRouteCandidatesRuntime = &model.PaymentRouteCandidatesRuntime{}
	}
	var (
		defaultPaymentRouteCandidatesRuntimeUpdateFields = defaultPaymentRouteCandidatesRuntimeUpdateFields(*paymentRouteCandidatesRuntime)
		tempUpdateField                                  PaymentRouteCandidatesRuntimeUpdateFieldList
		selectFields                                     = NewPaymentRouteCandidatesRuntimeSelectFields()
	)
	if len(paymentRouteCandidatesRuntimeUpdateFields) > 0 {
		for _, updateField := range paymentRouteCandidatesRuntimeUpdateFields {
			if updateField.paymentRouteCandidatesRuntimeField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentRouteCandidatesRuntimeUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentRouteCandidatesRuntimeCompositePrimaryKeyWhere([]model.PaymentRouteCandidatesRuntimePrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentRouteCandidatesRuntimeCommand(defaultPaymentRouteCandidatesRuntimeUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentRouteCandidatesRuntimeQueries.updatePaymentRouteCandidatesRuntime+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteCandidatesRuntime] error when try to update paymentRouteCandidatesRuntime by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentRouteCandidatesRuntimeByFilter(ctx context.Context, filter model.Filter, paymentRouteCandidatesRuntimeUpdateFields ...PaymentRouteCandidatesRuntimeUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentRouteCandidatesRuntimeUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentRouteCandidatesRuntimeUpdateFieldList
		selectFields = NewPaymentRouteCandidatesRuntimeSelectFields()
	)
	for _, updateField := range paymentRouteCandidatesRuntimeUpdateFields {
		if updateField.paymentRouteCandidatesRuntimeField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentRouteCandidatesRuntimeCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentRouteCandidatesRuntimeFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentRouteCandidatesRuntimeFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_route_candidates_runtime\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteCandidatesRuntimeByFilter] error when try to update paymentRouteCandidatesRuntime by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteCandidatesRuntimeByFilter] failed get rows affected")
	}
	return
}

var (
	paymentRouteCandidatesRuntimeQueries = struct {
		selectPaymentRouteCandidatesRuntime      string
		selectCountPaymentRouteCandidatesRuntime string
		deletePaymentRouteCandidatesRuntime      string
		updatePaymentRouteCandidatesRuntime      string
		insertPaymentRouteCandidatesRuntime      string
	}{
		selectPaymentRouteCandidatesRuntime:      "SELECT %s FROM \"payment_route_candidates_runtime\"",
		selectCountPaymentRouteCandidatesRuntime: "SELECT COUNT(\"id\") FROM \"payment_route_candidates_runtime\"",
		deletePaymentRouteCandidatesRuntime:      "DELETE FROM \"payment_route_candidates_runtime\"",
		updatePaymentRouteCandidatesRuntime:      "UPDATE \"payment_route_candidates_runtime\" SET %s ",
		insertPaymentRouteCandidatesRuntime:      "INSERT INTO \"payment_route_candidates_runtime\" %s VALUES %s",
	}
)

type PaymentRouteCandidatesRuntimeRepository interface {
	CreatePaymentRouteCandidatesRuntime(ctx context.Context, paymentRouteCandidatesRuntime *model.PaymentRouteCandidatesRuntime, fieldsInsert ...PaymentRouteCandidatesRuntimeField) error
	BulkCreatePaymentRouteCandidatesRuntime(ctx context.Context, paymentRouteCandidatesRuntimeList []*model.PaymentRouteCandidatesRuntime, fieldsInsert ...PaymentRouteCandidatesRuntimeField) error
	ResolvePaymentRouteCandidatesRuntime(ctx context.Context, selectFields ...PaymentRouteCandidatesRuntimeField) (model.PaymentRouteCandidatesRuntimeList, error)
	ResolvePaymentRouteCandidatesRuntimeByID(ctx context.Context, primaryID model.PaymentRouteCandidatesRuntimePrimaryID, selectFields ...PaymentRouteCandidatesRuntimeField) (model.PaymentRouteCandidatesRuntime, error)
	UpdatePaymentRouteCandidatesRuntimeByID(ctx context.Context, id model.PaymentRouteCandidatesRuntimePrimaryID, paymentRouteCandidatesRuntime *model.PaymentRouteCandidatesRuntime, paymentRouteCandidatesRuntimeUpdateFields ...PaymentRouteCandidatesRuntimeUpdateField) error
	UpdatePaymentRouteCandidatesRuntimeByFilter(ctx context.Context, filter model.Filter, paymentRouteCandidatesRuntimeUpdateFields ...PaymentRouteCandidatesRuntimeUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentRouteCandidatesRuntime(ctx context.Context, paymentRouteCandidatesRuntimeListMap map[model.PaymentRouteCandidatesRuntimePrimaryID]*model.PaymentRouteCandidatesRuntime, PaymentRouteCandidatesRuntimesMapUpdateFieldsRequest map[model.PaymentRouteCandidatesRuntimePrimaryID]PaymentRouteCandidatesRuntimeUpdateFieldList) (err error)
	DeletePaymentRouteCandidatesRuntimeByID(ctx context.Context, id model.PaymentRouteCandidatesRuntimePrimaryID) error
	BulkDeletePaymentRouteCandidatesRuntimeByIDs(ctx context.Context, ids []model.PaymentRouteCandidatesRuntimePrimaryID) error
	ResolvePaymentRouteCandidatesRuntimeByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentRouteCandidatesRuntimeFilterResult, err error)
	IsExistPaymentRouteCandidatesRuntimeByIDs(ctx context.Context, ids []model.PaymentRouteCandidatesRuntimePrimaryID) (exists bool, notFoundIds []model.PaymentRouteCandidatesRuntimePrimaryID, err error)
	IsExistPaymentRouteCandidatesRuntimeByID(ctx context.Context, id model.PaymentRouteCandidatesRuntimePrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
