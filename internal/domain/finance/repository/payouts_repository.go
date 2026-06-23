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

func composeInsertFieldsAndParamsPayouts(payoutsList []model.Payouts, fieldsInsert ...PayoutsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPayoutsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, payouts := range payoutsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, payouts.Id)
			case selectField.PayoutCode():
				args = append(args, payouts.PayoutCode)
			case selectField.PayoutBatchId():
				args = append(args, payouts.PayoutBatchId)
			case selectField.SettlementBatchId():
				args = append(args, payouts.SettlementBatchId)
			case selectField.MerchantPartyId():
				args = append(args, payouts.MerchantPartyId)
			case selectField.PayoutMethodId():
				args = append(args, payouts.PayoutMethodId)
			case selectField.ProviderAccountId():
				args = append(args, payouts.ProviderAccountId)
			case selectField.CurrencyCode():
				args = append(args, payouts.CurrencyCode)
			case selectField.Amount():
				args = append(args, payouts.Amount)
			case selectField.FeeAmount():
				args = append(args, payouts.FeeAmount)
			case selectField.NetSentAmount():
				args = append(args, payouts.NetSentAmount)
			case selectField.IdempotencyKey():
				args = append(args, payouts.IdempotencyKey)
			case selectField.ProviderPayoutRef():
				args = append(args, payouts.ProviderPayoutRef)
			case selectField.PayoutStatus():
				args = append(args, payouts.PayoutStatus)
			case selectField.HoldReasonCode():
				args = append(args, payouts.HoldReasonCode)
			case selectField.InitiatedAt():
				args = append(args, payouts.InitiatedAt)
			case selectField.CompletedAt():
				args = append(args, payouts.CompletedAt)
			case selectField.FailedAt():
				args = append(args, payouts.FailedAt)
			case selectField.FailureCode():
				args = append(args, payouts.FailureCode)
			case selectField.FailureReason():
				args = append(args, payouts.FailureReason)
			case selectField.Metadata():
				args = append(args, payouts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, payouts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, payouts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, payouts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, payouts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, payouts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, payouts.MetaDeletedBy)

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

func composePayoutsCompositePrimaryKeyWhere(primaryIDs []model.PayoutsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payouts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPayoutsSelectFields() string {
	fields := NewPayoutsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePayoutsSelectFields(selectFields ...PayoutsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PayoutsField string
type PayoutsFieldList []PayoutsField

type PayoutsSelectFields struct {
}

func (ss PayoutsSelectFields) Id() PayoutsField {
	return PayoutsField("id")
}

func (ss PayoutsSelectFields) PayoutCode() PayoutsField {
	return PayoutsField("payout_code")
}

func (ss PayoutsSelectFields) PayoutBatchId() PayoutsField {
	return PayoutsField("payout_batch_id")
}

func (ss PayoutsSelectFields) SettlementBatchId() PayoutsField {
	return PayoutsField("settlement_batch_id")
}

func (ss PayoutsSelectFields) MerchantPartyId() PayoutsField {
	return PayoutsField("merchant_party_id")
}

func (ss PayoutsSelectFields) PayoutMethodId() PayoutsField {
	return PayoutsField("payout_method_id")
}

func (ss PayoutsSelectFields) ProviderAccountId() PayoutsField {
	return PayoutsField("provider_account_id")
}

func (ss PayoutsSelectFields) CurrencyCode() PayoutsField {
	return PayoutsField("currency_code")
}

func (ss PayoutsSelectFields) Amount() PayoutsField {
	return PayoutsField("amount")
}

func (ss PayoutsSelectFields) FeeAmount() PayoutsField {
	return PayoutsField("fee_amount")
}

func (ss PayoutsSelectFields) NetSentAmount() PayoutsField {
	return PayoutsField("net_sent_amount")
}

func (ss PayoutsSelectFields) IdempotencyKey() PayoutsField {
	return PayoutsField("idempotency_key")
}

func (ss PayoutsSelectFields) ProviderPayoutRef() PayoutsField {
	return PayoutsField("provider_payout_ref")
}

func (ss PayoutsSelectFields) PayoutStatus() PayoutsField {
	return PayoutsField("payout_status")
}

func (ss PayoutsSelectFields) HoldReasonCode() PayoutsField {
	return PayoutsField("hold_reason_code")
}

func (ss PayoutsSelectFields) InitiatedAt() PayoutsField {
	return PayoutsField("initiated_at")
}

func (ss PayoutsSelectFields) CompletedAt() PayoutsField {
	return PayoutsField("completed_at")
}

func (ss PayoutsSelectFields) FailedAt() PayoutsField {
	return PayoutsField("failed_at")
}

func (ss PayoutsSelectFields) FailureCode() PayoutsField {
	return PayoutsField("failure_code")
}

func (ss PayoutsSelectFields) FailureReason() PayoutsField {
	return PayoutsField("failure_reason")
}

func (ss PayoutsSelectFields) Metadata() PayoutsField {
	return PayoutsField("metadata")
}

func (ss PayoutsSelectFields) MetaCreatedAt() PayoutsField {
	return PayoutsField("meta_created_at")
}

func (ss PayoutsSelectFields) MetaCreatedBy() PayoutsField {
	return PayoutsField("meta_created_by")
}

func (ss PayoutsSelectFields) MetaUpdatedAt() PayoutsField {
	return PayoutsField("meta_updated_at")
}

func (ss PayoutsSelectFields) MetaUpdatedBy() PayoutsField {
	return PayoutsField("meta_updated_by")
}

func (ss PayoutsSelectFields) MetaDeletedAt() PayoutsField {
	return PayoutsField("meta_deleted_at")
}

func (ss PayoutsSelectFields) MetaDeletedBy() PayoutsField {
	return PayoutsField("meta_deleted_by")
}

func (ss PayoutsSelectFields) All() PayoutsFieldList {
	return []PayoutsField{
		ss.Id(),
		ss.PayoutCode(),
		ss.PayoutBatchId(),
		ss.SettlementBatchId(),
		ss.MerchantPartyId(),
		ss.PayoutMethodId(),
		ss.ProviderAccountId(),
		ss.CurrencyCode(),
		ss.Amount(),
		ss.FeeAmount(),
		ss.NetSentAmount(),
		ss.IdempotencyKey(),
		ss.ProviderPayoutRef(),
		ss.PayoutStatus(),
		ss.HoldReasonCode(),
		ss.InitiatedAt(),
		ss.CompletedAt(),
		ss.FailedAt(),
		ss.FailureCode(),
		ss.FailureReason(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPayoutsSelectFields() PayoutsSelectFields {
	return PayoutsSelectFields{}
}

type PayoutsUpdateFieldOption struct {
	useIncrement bool
}
type PayoutsUpdateField struct {
	payoutsField PayoutsField
	opt          PayoutsUpdateFieldOption
	value        interface{}
}
type PayoutsUpdateFieldList []PayoutsUpdateField

func defaultPayoutsUpdateFieldOption() PayoutsUpdateFieldOption {
	return PayoutsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPayoutsOption(useIncrement bool) func(*PayoutsUpdateFieldOption) {
	return func(pcufo *PayoutsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPayoutsUpdateField(field PayoutsField, val interface{}, opts ...func(*PayoutsUpdateFieldOption)) PayoutsUpdateField {
	defaultOpt := defaultPayoutsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PayoutsUpdateField{
		payoutsField: field,
		value:        val,
		opt:          defaultOpt,
	}
}
func defaultPayoutsUpdateFields(payouts model.Payouts) (payoutsUpdateFieldList PayoutsUpdateFieldList) {
	selectFields := NewPayoutsSelectFields()
	payoutsUpdateFieldList = append(payoutsUpdateFieldList,
		NewPayoutsUpdateField(selectFields.Id(), payouts.Id),
		NewPayoutsUpdateField(selectFields.PayoutCode(), payouts.PayoutCode),
		NewPayoutsUpdateField(selectFields.PayoutBatchId(), payouts.PayoutBatchId),
		NewPayoutsUpdateField(selectFields.SettlementBatchId(), payouts.SettlementBatchId),
		NewPayoutsUpdateField(selectFields.MerchantPartyId(), payouts.MerchantPartyId),
		NewPayoutsUpdateField(selectFields.PayoutMethodId(), payouts.PayoutMethodId),
		NewPayoutsUpdateField(selectFields.ProviderAccountId(), payouts.ProviderAccountId),
		NewPayoutsUpdateField(selectFields.CurrencyCode(), payouts.CurrencyCode),
		NewPayoutsUpdateField(selectFields.Amount(), payouts.Amount),
		NewPayoutsUpdateField(selectFields.FeeAmount(), payouts.FeeAmount),
		NewPayoutsUpdateField(selectFields.NetSentAmount(), payouts.NetSentAmount),
		NewPayoutsUpdateField(selectFields.IdempotencyKey(), payouts.IdempotencyKey),
		NewPayoutsUpdateField(selectFields.ProviderPayoutRef(), payouts.ProviderPayoutRef),
		NewPayoutsUpdateField(selectFields.PayoutStatus(), payouts.PayoutStatus),
		NewPayoutsUpdateField(selectFields.HoldReasonCode(), payouts.HoldReasonCode),
		NewPayoutsUpdateField(selectFields.InitiatedAt(), payouts.InitiatedAt),
		NewPayoutsUpdateField(selectFields.CompletedAt(), payouts.CompletedAt),
		NewPayoutsUpdateField(selectFields.FailedAt(), payouts.FailedAt),
		NewPayoutsUpdateField(selectFields.FailureCode(), payouts.FailureCode),
		NewPayoutsUpdateField(selectFields.FailureReason(), payouts.FailureReason),
		NewPayoutsUpdateField(selectFields.Metadata(), payouts.Metadata),
		NewPayoutsUpdateField(selectFields.MetaCreatedAt(), payouts.MetaCreatedAt),
		NewPayoutsUpdateField(selectFields.MetaCreatedBy(), payouts.MetaCreatedBy),
		NewPayoutsUpdateField(selectFields.MetaUpdatedAt(), payouts.MetaUpdatedAt),
		NewPayoutsUpdateField(selectFields.MetaUpdatedBy(), payouts.MetaUpdatedBy),
		NewPayoutsUpdateField(selectFields.MetaDeletedAt(), payouts.MetaDeletedAt),
		NewPayoutsUpdateField(selectFields.MetaDeletedBy(), payouts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPayoutsCommand(payoutsUpdateFieldList PayoutsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range payoutsUpdateFieldList {
		field := string(updateField.payoutsField)
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

func (repo *RepositoryImpl) BulkCreatePayouts(ctx context.Context, payoutsList []*model.Payouts, fieldsInsert ...PayoutsField) (err error) {
	var (
		fieldsStr        string
		valueListStr     []string
		argsList         []interface{}
		primaryIds       []model.PayoutsPrimaryID
		payoutsValueList []model.Payouts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPayoutsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, payouts := range payoutsList {

		primaryIds = append(primaryIds, payouts.ToPayoutsPrimaryID())

		payoutsValueList = append(payoutsValueList, *payouts)
	}

	_, notFoundIds, err := repo.IsExistPayoutsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayouts] failed checking payouts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PayoutsPrimaryID{}
		mapNotFoundIds := map[model.PayoutsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "payouts", fmt.Sprintf("payouts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPayouts(payoutsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(payoutsQueries.insertPayouts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayouts] failed exec create payouts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePayoutsByIDs(ctx context.Context, primaryIDs []model.PayoutsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPayoutsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutsByIDs] failed checking payouts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payouts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payouts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(payoutsQueries.deletePayouts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPayoutsByIDs(ctx context.Context, ids []model.PayoutsPrimaryID) (exists bool, notFoundIds []model.PayoutsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payouts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(payoutsQueries.selectPayouts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PayoutsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PayoutsPrimaryID]bool{}
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

// BulkUpdatePayouts is used to bulk update payouts, by default it will update all field
// if want to update specific field, then fill payoutssMapUpdateFieldsRequest else please fill payoutssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePayouts(ctx context.Context, payoutssMap map[model.PayoutsPrimaryID]*model.Payouts, payoutssMapUpdateFieldsRequest map[model.PayoutsPrimaryID]PayoutsUpdateFieldList) (err error) {
	if len(payoutssMap) == 0 && len(payoutssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		payoutssMapUpdateField map[model.PayoutsPrimaryID]PayoutsUpdateFieldList = map[model.PayoutsPrimaryID]PayoutsUpdateFieldList{}
		asTableValues          string                                            = "myvalues"
	)

	if len(payoutssMap) > 0 {
		for id, payouts := range payoutssMap {
			if payouts == nil {
				log.Error().Err(err).Msg("[BulkUpdatePayouts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			payoutssMapUpdateField[id] = defaultPayoutsUpdateFields(*payouts)
		}
	} else {
		payoutssMapUpdateField = payoutssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePayoutsQuery(payoutssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPayoutsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayouts] failed checking payouts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payouts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePayoutsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payouts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayouts] failed exec query")
	}
	return
}

type PayoutsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPayoutsFieldParameter(param string, args ...interface{}) PayoutsFieldParameter {
	return PayoutsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePayoutsQuery(mapPayoutss map[model.PayoutsPrimaryID]PayoutsUpdateFieldList, asTableValues string) (primaryIDs []model.PayoutsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PayoutsPrimaryID]map[string]interface{}{}
	payoutsSelectFields := NewPayoutsSelectFields()
	for id, updateFields := range mapPayoutss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.payoutsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPayoutss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPayoutsFieldType(updateField.payoutsField)))
			args = append(args, fields[string(updateField.payoutsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.payoutsField))
		if updateField.payoutsField == payoutsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.payoutsField, asTableValues, updateField.payoutsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.payoutsField,
				"\"payouts\"", updateField.payoutsField,
				asTableValues, updateField.payoutsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePayoutsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PayoutsPrimaryID, asTableValue string) (whereQry string) {
	payoutsSelectFields := NewPayoutsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payouts\".\"id\" = %s.\"id\"::"+GetPayoutsFieldType(payoutsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPayoutsFieldType(payoutsField PayoutsField) string {
	selectPayoutsFields := NewPayoutsSelectFields()
	switch payoutsField {

	case selectPayoutsFields.Id():
		return "uuid"

	case selectPayoutsFields.PayoutCode():
		return "text"

	case selectPayoutsFields.PayoutBatchId():
		return "uuid"

	case selectPayoutsFields.SettlementBatchId():
		return "uuid"

	case selectPayoutsFields.MerchantPartyId():
		return "uuid"

	case selectPayoutsFields.PayoutMethodId():
		return "uuid"

	case selectPayoutsFields.ProviderAccountId():
		return "uuid"

	case selectPayoutsFields.CurrencyCode():
		return "text"

	case selectPayoutsFields.Amount():
		return "numeric"

	case selectPayoutsFields.FeeAmount():
		return "numeric"

	case selectPayoutsFields.NetSentAmount():
		return "numeric"

	case selectPayoutsFields.IdempotencyKey():
		return "text"

	case selectPayoutsFields.ProviderPayoutRef():
		return "text"

	case selectPayoutsFields.PayoutStatus():
		return "payout_status_enum"

	case selectPayoutsFields.HoldReasonCode():
		return "text"

	case selectPayoutsFields.InitiatedAt():
		return "timestamptz"

	case selectPayoutsFields.CompletedAt():
		return "timestamptz"

	case selectPayoutsFields.FailedAt():
		return "timestamptz"

	case selectPayoutsFields.FailureCode():
		return "text"

	case selectPayoutsFields.FailureReason():
		return "text"

	case selectPayoutsFields.Metadata():
		return "jsonb"

	case selectPayoutsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPayoutsFields.MetaCreatedBy():
		return "uuid"

	case selectPayoutsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPayoutsFields.MetaUpdatedBy():
		return "uuid"

	case selectPayoutsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPayoutsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePayouts(ctx context.Context, payouts *model.Payouts, fieldsInsert ...PayoutsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPayoutsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PayoutsPrimaryID{
		Id: payouts.Id,
	}
	exists, err := repo.IsExistPayoutsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayouts] failed checking payouts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "payouts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPayouts([]model.Payouts{*payouts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(payoutsQueries.insertPayouts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayouts] failed exec create payouts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePayoutsByID(ctx context.Context, primaryID model.PayoutsPrimaryID) (err error) {
	exists, err := repo.IsExistPayoutsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutsByID] failed checking payouts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payouts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePayoutsCompositePrimaryKeyWhere([]model.PayoutsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(payoutsQueries.deletePayouts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutsFilterResult, err error) {
	query, args, err := composePayoutsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutsByFilter] failed compose payouts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutsByFilter] failed get payouts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePayoutsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PayoutsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePayoutsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePayoutsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePayoutsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 27 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPayoutsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 27+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_code\"")
			selectedColumns["payout_code"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_batch_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_batch_id\"")
			selectedColumns["payout_batch_id"] = struct{}{}
		}
		if _, selected := selectedColumns["settlement_batch_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"settlement_batch_id\"")
			selectedColumns["settlement_batch_id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_method_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_method_id\"")
			selectedColumns["payout_method_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_amount\"")
			selectedColumns["fee_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["net_sent_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"net_sent_amount\"")
			selectedColumns["net_sent_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_payout_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_payout_ref\"")
			selectedColumns["provider_payout_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_status\"")
			selectedColumns["payout_status"] = struct{}{}
		}
		if _, selected := selectedColumns["hold_reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"hold_reason_code\"")
			selectedColumns["hold_reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["initiated_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"initiated_at\"")
			selectedColumns["initiated_at"] = struct{}{}
		}
		if _, selected := selectedColumns["completed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"completed_at\"")
			selectedColumns["completed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["failed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"failed_at\"")
			selectedColumns["failed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_code\"")
			selectedColumns["failure_code"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_reason\"")
			selectedColumns["failure_reason"] = struct{}{}
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

type payoutsFilterPlaceholder struct {
	index int
}

func (p *payoutsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePayoutsFilterPredicate(filterField model.FilterField, placeholders *payoutsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPayoutsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePayoutsFilterSQLExpr(spec)
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

func composePayoutsFilterGroup(group model.FilterGroup, placeholders *payoutsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePayoutsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePayoutsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePayoutsFilterWhereQueries(filter model.Filter, placeholders *payoutsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePayoutsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePayoutsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePayoutsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePayoutsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePayoutsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePayoutsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := payoutsFilterPlaceholder{index: 1}
	whereQueries, err := composePayoutsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPayoutsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePayoutsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePayoutsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payouts\" base%s", strings.Join(selectColumns, ","), composePayoutsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPayoutsByID(ctx context.Context, primaryID model.PayoutsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePayoutsCompositePrimaryKeyWhere([]model.PayoutsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", payoutsQueries.selectCountPayouts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayouts(ctx context.Context, selectFields ...PayoutsField) (payoutsList model.PayoutsList, err error) {
	var (
		defaultPayoutsSelectFields = defaultPayoutsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutsSelectFields = composePayoutsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(payoutsQueries.selectPayouts, defaultPayoutsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &payoutsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayouts] failed get payouts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutsByID(ctx context.Context, primaryID model.PayoutsPrimaryID, selectFields ...PayoutsField) (payouts model.Payouts, err error) {
	var (
		defaultPayoutsSelectFields = defaultPayoutsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutsSelectFields = composePayoutsSelectFields(selectFields...)
	}
	whereQry, params := composePayoutsCompositePrimaryKeyWhere([]model.PayoutsPrimaryID{primaryID})
	query := fmt.Sprintf(payoutsQueries.selectPayouts+" WHERE "+whereQry, defaultPayoutsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &payouts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("payouts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePayoutsByID] failed get payouts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePayoutsByID(ctx context.Context, primaryID model.PayoutsPrimaryID, payouts *model.Payouts, payoutsUpdateFields ...PayoutsUpdateField) (err error) {
	exists, err := repo.IsExistPayoutsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayouts] failed checking payouts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payouts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if payouts == nil {
		if len(payoutsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePayoutsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		payouts = &model.Payouts{}
	}
	var (
		defaultPayoutsUpdateFields = defaultPayoutsUpdateFields(*payouts)
		tempUpdateField            PayoutsUpdateFieldList
		selectFields               = NewPayoutsSelectFields()
	)
	if len(payoutsUpdateFields) > 0 {
		for _, updateField := range payoutsUpdateFields {
			if updateField.payoutsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPayoutsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePayoutsCompositePrimaryKeyWhere([]model.PayoutsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPayoutsCommand(defaultPayoutsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(payoutsQueries.updatePayouts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayouts] error when try to update payouts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePayoutsByFilter(ctx context.Context, filter model.Filter, payoutsUpdateFields ...PayoutsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(payoutsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PayoutsUpdateFieldList
		selectFields = NewPayoutsSelectFields()
	)
	for _, updateField := range payoutsUpdateFields {
		if updateField.payoutsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPayoutsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := payoutsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePayoutsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payouts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutsByFilter] error when try to update payouts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutsByFilter] failed get rows affected")
	}
	return
}

var (
	payoutsQueries = struct {
		selectPayouts      string
		selectCountPayouts string
		deletePayouts      string
		updatePayouts      string
		insertPayouts      string
	}{
		selectPayouts:      "SELECT %s FROM \"payouts\"",
		selectCountPayouts: "SELECT COUNT(\"id\") FROM \"payouts\"",
		deletePayouts:      "DELETE FROM \"payouts\"",
		updatePayouts:      "UPDATE \"payouts\" SET %s ",
		insertPayouts:      "INSERT INTO \"payouts\" %s VALUES %s",
	}
)

type PayoutsRepository interface {
	CreatePayouts(ctx context.Context, payouts *model.Payouts, fieldsInsert ...PayoutsField) error
	BulkCreatePayouts(ctx context.Context, payoutsList []*model.Payouts, fieldsInsert ...PayoutsField) error
	ResolvePayouts(ctx context.Context, selectFields ...PayoutsField) (model.PayoutsList, error)
	ResolvePayoutsByID(ctx context.Context, primaryID model.PayoutsPrimaryID, selectFields ...PayoutsField) (model.Payouts, error)
	UpdatePayoutsByID(ctx context.Context, id model.PayoutsPrimaryID, payouts *model.Payouts, payoutsUpdateFields ...PayoutsUpdateField) error
	UpdatePayoutsByFilter(ctx context.Context, filter model.Filter, payoutsUpdateFields ...PayoutsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePayouts(ctx context.Context, payoutsListMap map[model.PayoutsPrimaryID]*model.Payouts, PayoutssMapUpdateFieldsRequest map[model.PayoutsPrimaryID]PayoutsUpdateFieldList) (err error)
	DeletePayoutsByID(ctx context.Context, id model.PayoutsPrimaryID) error
	BulkDeletePayoutsByIDs(ctx context.Context, ids []model.PayoutsPrimaryID) error
	ResolvePayoutsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutsFilterResult, err error)
	IsExistPayoutsByIDs(ctx context.Context, ids []model.PayoutsPrimaryID) (exists bool, notFoundIds []model.PayoutsPrimaryID, err error)
	IsExistPayoutsByID(ctx context.Context, id model.PayoutsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
