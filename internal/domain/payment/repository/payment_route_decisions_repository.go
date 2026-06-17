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

func composeInsertFieldsAndParamsPaymentRouteDecisions(paymentRouteDecisionsList []model.PaymentRouteDecisions, fieldsInsert ...PaymentRouteDecisionsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentRouteDecisionsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentRouteDecisions := range paymentRouteDecisionsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentRouteDecisions.Id)
			case selectField.PaymentIntentId():
				args = append(args, paymentRouteDecisions.PaymentIntentId)
			case selectField.SelectedProviderAccountId():
				args = append(args, paymentRouteDecisions.SelectedProviderAccountId)
			case selectField.SelectedProviderCode():
				args = append(args, paymentRouteDecisions.SelectedProviderCode)
			case selectField.MethodCode():
				args = append(args, paymentRouteDecisions.MethodCode)
			case selectField.ChannelCode():
				args = append(args, paymentRouteDecisions.ChannelCode)
			case selectField.Reason():
				args = append(args, paymentRouteDecisions.Reason)
			case selectField.EvaluatedContext():
				args = append(args, paymentRouteDecisions.EvaluatedContext)
			case selectField.Candidates():
				args = append(args, paymentRouteDecisions.Candidates)
			case selectField.Metadata():
				args = append(args, paymentRouteDecisions.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentRouteDecisions.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentRouteDecisions.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentRouteDecisions.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentRouteDecisions.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentRouteDecisions.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentRouteDecisions.MetaDeletedBy)

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

func composePaymentRouteDecisionsCompositePrimaryKeyWhere(primaryIDs []model.PaymentRouteDecisionsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_route_decisions\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentRouteDecisionsSelectFields() string {
	fields := NewPaymentRouteDecisionsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentRouteDecisionsSelectFields(selectFields ...PaymentRouteDecisionsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentRouteDecisionsField string
type PaymentRouteDecisionsFieldList []PaymentRouteDecisionsField

type PaymentRouteDecisionsSelectFields struct {
}

func (ss PaymentRouteDecisionsSelectFields) Id() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("id")
}

func (ss PaymentRouteDecisionsSelectFields) PaymentIntentId() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("payment_intent_id")
}

func (ss PaymentRouteDecisionsSelectFields) SelectedProviderAccountId() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("selected_provider_account_id")
}

func (ss PaymentRouteDecisionsSelectFields) SelectedProviderCode() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("selected_provider_code")
}

func (ss PaymentRouteDecisionsSelectFields) MethodCode() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("method_code")
}

func (ss PaymentRouteDecisionsSelectFields) ChannelCode() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("channel_code")
}

func (ss PaymentRouteDecisionsSelectFields) Reason() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("reason")
}

func (ss PaymentRouteDecisionsSelectFields) EvaluatedContext() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("evaluated_context")
}

func (ss PaymentRouteDecisionsSelectFields) Candidates() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("candidates")
}

func (ss PaymentRouteDecisionsSelectFields) Metadata() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("metadata")
}

func (ss PaymentRouteDecisionsSelectFields) MetaCreatedAt() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("meta_created_at")
}

func (ss PaymentRouteDecisionsSelectFields) MetaCreatedBy() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("meta_created_by")
}

func (ss PaymentRouteDecisionsSelectFields) MetaUpdatedAt() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("meta_updated_at")
}

func (ss PaymentRouteDecisionsSelectFields) MetaUpdatedBy() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("meta_updated_by")
}

func (ss PaymentRouteDecisionsSelectFields) MetaDeletedAt() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("meta_deleted_at")
}

func (ss PaymentRouteDecisionsSelectFields) MetaDeletedBy() PaymentRouteDecisionsField {
	return PaymentRouteDecisionsField("meta_deleted_by")
}

func (ss PaymentRouteDecisionsSelectFields) All() PaymentRouteDecisionsFieldList {
	return []PaymentRouteDecisionsField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.SelectedProviderAccountId(),
		ss.SelectedProviderCode(),
		ss.MethodCode(),
		ss.ChannelCode(),
		ss.Reason(),
		ss.EvaluatedContext(),
		ss.Candidates(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentRouteDecisionsSelectFields() PaymentRouteDecisionsSelectFields {
	return PaymentRouteDecisionsSelectFields{}
}

type PaymentRouteDecisionsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentRouteDecisionsUpdateField struct {
	paymentRouteDecisionsField PaymentRouteDecisionsField
	opt                        PaymentRouteDecisionsUpdateFieldOption
	value                      interface{}
}
type PaymentRouteDecisionsUpdateFieldList []PaymentRouteDecisionsUpdateField

func defaultPaymentRouteDecisionsUpdateFieldOption() PaymentRouteDecisionsUpdateFieldOption {
	return PaymentRouteDecisionsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentRouteDecisionsOption(useIncrement bool) func(*PaymentRouteDecisionsUpdateFieldOption) {
	return func(pcufo *PaymentRouteDecisionsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentRouteDecisionsUpdateField(field PaymentRouteDecisionsField, val interface{}, opts ...func(*PaymentRouteDecisionsUpdateFieldOption)) PaymentRouteDecisionsUpdateField {
	defaultOpt := defaultPaymentRouteDecisionsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentRouteDecisionsUpdateField{
		paymentRouteDecisionsField: field,
		value:                      val,
		opt:                        defaultOpt,
	}
}
func defaultPaymentRouteDecisionsUpdateFields(paymentRouteDecisions model.PaymentRouteDecisions) (paymentRouteDecisionsUpdateFieldList PaymentRouteDecisionsUpdateFieldList) {
	selectFields := NewPaymentRouteDecisionsSelectFields()
	paymentRouteDecisionsUpdateFieldList = append(paymentRouteDecisionsUpdateFieldList,
		NewPaymentRouteDecisionsUpdateField(selectFields.Id(), paymentRouteDecisions.Id),
		NewPaymentRouteDecisionsUpdateField(selectFields.PaymentIntentId(), paymentRouteDecisions.PaymentIntentId),
		NewPaymentRouteDecisionsUpdateField(selectFields.SelectedProviderAccountId(), paymentRouteDecisions.SelectedProviderAccountId),
		NewPaymentRouteDecisionsUpdateField(selectFields.SelectedProviderCode(), paymentRouteDecisions.SelectedProviderCode),
		NewPaymentRouteDecisionsUpdateField(selectFields.MethodCode(), paymentRouteDecisions.MethodCode),
		NewPaymentRouteDecisionsUpdateField(selectFields.ChannelCode(), paymentRouteDecisions.ChannelCode),
		NewPaymentRouteDecisionsUpdateField(selectFields.Reason(), paymentRouteDecisions.Reason),
		NewPaymentRouteDecisionsUpdateField(selectFields.EvaluatedContext(), paymentRouteDecisions.EvaluatedContext),
		NewPaymentRouteDecisionsUpdateField(selectFields.Candidates(), paymentRouteDecisions.Candidates),
		NewPaymentRouteDecisionsUpdateField(selectFields.Metadata(), paymentRouteDecisions.Metadata),
		NewPaymentRouteDecisionsUpdateField(selectFields.MetaCreatedAt(), paymentRouteDecisions.MetaCreatedAt),
		NewPaymentRouteDecisionsUpdateField(selectFields.MetaCreatedBy(), paymentRouteDecisions.MetaCreatedBy),
		NewPaymentRouteDecisionsUpdateField(selectFields.MetaUpdatedAt(), paymentRouteDecisions.MetaUpdatedAt),
		NewPaymentRouteDecisionsUpdateField(selectFields.MetaUpdatedBy(), paymentRouteDecisions.MetaUpdatedBy),
		NewPaymentRouteDecisionsUpdateField(selectFields.MetaDeletedAt(), paymentRouteDecisions.MetaDeletedAt),
		NewPaymentRouteDecisionsUpdateField(selectFields.MetaDeletedBy(), paymentRouteDecisions.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentRouteDecisionsCommand(paymentRouteDecisionsUpdateFieldList PaymentRouteDecisionsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentRouteDecisionsUpdateFieldList {
		field := string(updateField.paymentRouteDecisionsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentRouteDecisions(ctx context.Context, paymentRouteDecisionsList []*model.PaymentRouteDecisions, fieldsInsert ...PaymentRouteDecisionsField) (err error) {
	var (
		fieldsStr                      string
		valueListStr                   []string
		argsList                       []interface{}
		primaryIds                     []model.PaymentRouteDecisionsPrimaryID
		paymentRouteDecisionsValueList []model.PaymentRouteDecisions
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentRouteDecisionsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentRouteDecisions := range paymentRouteDecisionsList {

		primaryIds = append(primaryIds, paymentRouteDecisions.ToPaymentRouteDecisionsPrimaryID())

		paymentRouteDecisionsValueList = append(paymentRouteDecisionsValueList, *paymentRouteDecisions)
	}

	_, notFoundIds, err := repo.IsExistPaymentRouteDecisionsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentRouteDecisions] failed checking paymentRouteDecisions whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentRouteDecisionsPrimaryID{}
		mapNotFoundIds := map[model.PaymentRouteDecisionsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentRouteDecisions", fmt.Sprintf("paymentRouteDecisions with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentRouteDecisions(paymentRouteDecisionsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentRouteDecisionsQueries.insertPaymentRouteDecisions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentRouteDecisions] failed exec create paymentRouteDecisions query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentRouteDecisionsByIDs(ctx context.Context, primaryIDs []model.PaymentRouteDecisionsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentRouteDecisionsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRouteDecisionsByIDs] failed checking paymentRouteDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteDecisions with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_route_decisions\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentRouteDecisionsQueries.deletePaymentRouteDecisions + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRouteDecisionsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentRouteDecisionsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentRouteDecisionsByIDs(ctx context.Context, ids []model.PaymentRouteDecisionsPrimaryID) (exists bool, notFoundIds []model.PaymentRouteDecisionsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_route_decisions\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentRouteDecisionsQueries.selectPaymentRouteDecisions, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRouteDecisionsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentRouteDecisionsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRouteDecisionsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentRouteDecisionsPrimaryID]bool{}
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

// BulkUpdatePaymentRouteDecisions is used to bulk update paymentRouteDecisions, by default it will update all field
// if want to update specific field, then fill paymentRouteDecisionssMapUpdateFieldsRequest else please fill paymentRouteDecisionssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentRouteDecisions(ctx context.Context, paymentRouteDecisionssMap map[model.PaymentRouteDecisionsPrimaryID]*model.PaymentRouteDecisions, paymentRouteDecisionssMapUpdateFieldsRequest map[model.PaymentRouteDecisionsPrimaryID]PaymentRouteDecisionsUpdateFieldList) (err error) {
	if len(paymentRouteDecisionssMap) == 0 && len(paymentRouteDecisionssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentRouteDecisionssMapUpdateField map[model.PaymentRouteDecisionsPrimaryID]PaymentRouteDecisionsUpdateFieldList = map[model.PaymentRouteDecisionsPrimaryID]PaymentRouteDecisionsUpdateFieldList{}
		asTableValues                        string                                                                        = "myvalues"
	)

	if len(paymentRouteDecisionssMap) > 0 {
		for id, paymentRouteDecisions := range paymentRouteDecisionssMap {
			if paymentRouteDecisions == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentRouteDecisions] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentRouteDecisionssMapUpdateField[id] = defaultPaymentRouteDecisionsUpdateFields(*paymentRouteDecisions)
		}
	} else {
		paymentRouteDecisionssMapUpdateField = paymentRouteDecisionssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentRouteDecisionsQuery(paymentRouteDecisionssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentRouteDecisionsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentRouteDecisions] failed checking paymentRouteDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteDecisions with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentRouteDecisionsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_route_decisions\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentRouteDecisions] failed exec query")
	}
	return
}

type PaymentRouteDecisionsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentRouteDecisionsFieldParameter(param string, args ...interface{}) PaymentRouteDecisionsFieldParameter {
	return PaymentRouteDecisionsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentRouteDecisionsQuery(mapPaymentRouteDecisionss map[model.PaymentRouteDecisionsPrimaryID]PaymentRouteDecisionsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentRouteDecisionsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentRouteDecisionsPrimaryID]map[string]interface{}{}
	paymentRouteDecisionsSelectFields := NewPaymentRouteDecisionsSelectFields()
	for id, updateFields := range mapPaymentRouteDecisionss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentRouteDecisionsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentRouteDecisionss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentRouteDecisionsFieldType(updateField.paymentRouteDecisionsField)))
			args = append(args, fields[string(updateField.paymentRouteDecisionsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentRouteDecisionsField))
		if updateField.paymentRouteDecisionsField == paymentRouteDecisionsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentRouteDecisionsField, asTableValues, updateField.paymentRouteDecisionsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentRouteDecisionsField,
				"\"payment_route_decisions\"", updateField.paymentRouteDecisionsField,
				asTableValues, updateField.paymentRouteDecisionsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentRouteDecisionsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentRouteDecisionsPrimaryID, asTableValue string) (whereQry string) {
	paymentRouteDecisionsSelectFields := NewPaymentRouteDecisionsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_route_decisions\".\"id\" = %s.\"id\"::"+GetPaymentRouteDecisionsFieldType(paymentRouteDecisionsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentRouteDecisionsFieldType(paymentRouteDecisionsField PaymentRouteDecisionsField) string {
	selectPaymentRouteDecisionsFields := NewPaymentRouteDecisionsSelectFields()
	switch paymentRouteDecisionsField {

	case selectPaymentRouteDecisionsFields.Id():
		return "uuid"

	case selectPaymentRouteDecisionsFields.PaymentIntentId():
		return "uuid"

	case selectPaymentRouteDecisionsFields.SelectedProviderAccountId():
		return "uuid"

	case selectPaymentRouteDecisionsFields.SelectedProviderCode():
		return "text"

	case selectPaymentRouteDecisionsFields.MethodCode():
		return "text"

	case selectPaymentRouteDecisionsFields.ChannelCode():
		return "text"

	case selectPaymentRouteDecisionsFields.Reason():
		return "text"

	case selectPaymentRouteDecisionsFields.EvaluatedContext():
		return "jsonb"

	case selectPaymentRouteDecisionsFields.Candidates():
		return "jsonb"

	case selectPaymentRouteDecisionsFields.Metadata():
		return "jsonb"

	case selectPaymentRouteDecisionsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentRouteDecisionsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentRouteDecisionsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentRouteDecisionsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentRouteDecisionsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentRouteDecisionsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentRouteDecisions(ctx context.Context, paymentRouteDecisions *model.PaymentRouteDecisions, fieldsInsert ...PaymentRouteDecisionsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentRouteDecisionsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentRouteDecisionsPrimaryID{
		Id: paymentRouteDecisions.Id,
	}
	exists, err := repo.IsExistPaymentRouteDecisionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentRouteDecisions] failed checking paymentRouteDecisions whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentRouteDecisions", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentRouteDecisions([]model.PaymentRouteDecisions{*paymentRouteDecisions}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentRouteDecisionsQueries.insertPaymentRouteDecisions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentRouteDecisions] failed exec create paymentRouteDecisions query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentRouteDecisionsByID(ctx context.Context, primaryID model.PaymentRouteDecisionsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentRouteDecisionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentRouteDecisionsByID] failed checking paymentRouteDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteDecisions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentRouteDecisionsCompositePrimaryKeyWhere([]model.PaymentRouteDecisionsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentRouteDecisionsQueries.deletePaymentRouteDecisions + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentRouteDecisionsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRouteDecisionsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentRouteDecisionsFilterResult, err error) {
	query, args, err := composePaymentRouteDecisionsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRouteDecisionsByFilter] failed compose paymentRouteDecisions filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRouteDecisionsByFilter] failed get paymentRouteDecisions by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentRouteDecisionsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentRouteDecisionsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentRouteDecisionsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentRouteDecisionsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentRouteDecisionsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 16 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentRouteDecisionsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 16+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["selected_provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"selected_provider_account_id\"")
			selectedColumns["selected_provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["selected_provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"selected_provider_code\"")
			selectedColumns["selected_provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["method_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_code\"")
			selectedColumns["method_code"] = struct{}{}
		}
		if _, selected := selectedColumns["channel_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"channel_code\"")
			selectedColumns["channel_code"] = struct{}{}
		}
		if _, selected := selectedColumns["reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason\"")
			selectedColumns["reason"] = struct{}{}
		}
		if _, selected := selectedColumns["evaluated_context"]; !selected {
			selectColumns = append(selectColumns, "base.\"evaluated_context\"")
			selectedColumns["evaluated_context"] = struct{}{}
		}
		if _, selected := selectedColumns["candidates"]; !selected {
			selectColumns = append(selectColumns, "base.\"candidates\"")
			selectedColumns["candidates"] = struct{}{}
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

type paymentRouteDecisionsFilterPlaceholder struct {
	index int
}

func (p *paymentRouteDecisionsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentRouteDecisionsFilterPredicate(filterField model.FilterField, placeholders *paymentRouteDecisionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentRouteDecisionsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentRouteDecisionsFilterSQLExpr(spec)
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

func composePaymentRouteDecisionsFilterGroup(group model.FilterGroup, placeholders *paymentRouteDecisionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentRouteDecisionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentRouteDecisionsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentRouteDecisionsFilterWhereQueries(filter model.Filter, placeholders *paymentRouteDecisionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentRouteDecisionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentRouteDecisionsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentRouteDecisionsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentRouteDecisionsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentRouteDecisionsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentRouteDecisionsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentRouteDecisionsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentRouteDecisionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentRouteDecisionsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentRouteDecisionsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentRouteDecisionsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_route_decisions\" base%s", strings.Join(selectColumns, ","), composePaymentRouteDecisionsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentRouteDecisionsByID(ctx context.Context, primaryID model.PaymentRouteDecisionsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentRouteDecisionsCompositePrimaryKeyWhere([]model.PaymentRouteDecisionsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentRouteDecisionsQueries.selectCountPaymentRouteDecisions, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentRouteDecisionsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRouteDecisions(ctx context.Context, selectFields ...PaymentRouteDecisionsField) (paymentRouteDecisionsList model.PaymentRouteDecisionsList, err error) {
	var (
		defaultPaymentRouteDecisionsSelectFields = defaultPaymentRouteDecisionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentRouteDecisionsSelectFields = composePaymentRouteDecisionsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentRouteDecisionsQueries.selectPaymentRouteDecisions, defaultPaymentRouteDecisionsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentRouteDecisionsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentRouteDecisions] failed get paymentRouteDecisions list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentRouteDecisionsByID(ctx context.Context, primaryID model.PaymentRouteDecisionsPrimaryID, selectFields ...PaymentRouteDecisionsField) (paymentRouteDecisions model.PaymentRouteDecisions, err error) {
	var (
		defaultPaymentRouteDecisionsSelectFields = defaultPaymentRouteDecisionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentRouteDecisionsSelectFields = composePaymentRouteDecisionsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentRouteDecisionsCompositePrimaryKeyWhere([]model.PaymentRouteDecisionsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentRouteDecisionsQueries.selectPaymentRouteDecisions+" WHERE "+whereQry, defaultPaymentRouteDecisionsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentRouteDecisions, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentRouteDecisions with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentRouteDecisionsByID] failed get paymentRouteDecisions")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentRouteDecisionsByID(ctx context.Context, primaryID model.PaymentRouteDecisionsPrimaryID, paymentRouteDecisions *model.PaymentRouteDecisions, paymentRouteDecisionsUpdateFields ...PaymentRouteDecisionsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentRouteDecisionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteDecisions] failed checking paymentRouteDecisions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentRouteDecisions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentRouteDecisions == nil {
		if len(paymentRouteDecisionsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentRouteDecisionsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentRouteDecisions = &model.PaymentRouteDecisions{}
	}
	var (
		defaultPaymentRouteDecisionsUpdateFields = defaultPaymentRouteDecisionsUpdateFields(*paymentRouteDecisions)
		tempUpdateField                          PaymentRouteDecisionsUpdateFieldList
		selectFields                             = NewPaymentRouteDecisionsSelectFields()
	)
	if len(paymentRouteDecisionsUpdateFields) > 0 {
		for _, updateField := range paymentRouteDecisionsUpdateFields {
			if updateField.paymentRouteDecisionsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentRouteDecisionsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentRouteDecisionsCompositePrimaryKeyWhere([]model.PaymentRouteDecisionsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentRouteDecisionsCommand(defaultPaymentRouteDecisionsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentRouteDecisionsQueries.updatePaymentRouteDecisions+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteDecisions] error when try to update paymentRouteDecisions by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentRouteDecisionsByFilter(ctx context.Context, filter model.Filter, paymentRouteDecisionsUpdateFields ...PaymentRouteDecisionsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentRouteDecisionsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentRouteDecisionsUpdateFieldList
		selectFields = NewPaymentRouteDecisionsSelectFields()
	)
	for _, updateField := range paymentRouteDecisionsUpdateFields {
		if updateField.paymentRouteDecisionsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentRouteDecisionsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentRouteDecisionsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentRouteDecisionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_route_decisions\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteDecisionsByFilter] error when try to update paymentRouteDecisions by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentRouteDecisionsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentRouteDecisionsQueries = struct {
		selectPaymentRouteDecisions      string
		selectCountPaymentRouteDecisions string
		deletePaymentRouteDecisions      string
		updatePaymentRouteDecisions      string
		insertPaymentRouteDecisions      string
	}{
		selectPaymentRouteDecisions:      "SELECT %s FROM \"payment_route_decisions\"",
		selectCountPaymentRouteDecisions: "SELECT COUNT(\"id\") FROM \"payment_route_decisions\"",
		deletePaymentRouteDecisions:      "DELETE FROM \"payment_route_decisions\"",
		updatePaymentRouteDecisions:      "UPDATE \"payment_route_decisions\" SET %s ",
		insertPaymentRouteDecisions:      "INSERT INTO \"payment_route_decisions\" %s VALUES %s",
	}
)

type PaymentRouteDecisionsRepository interface {
	CreatePaymentRouteDecisions(ctx context.Context, paymentRouteDecisions *model.PaymentRouteDecisions, fieldsInsert ...PaymentRouteDecisionsField) error
	BulkCreatePaymentRouteDecisions(ctx context.Context, paymentRouteDecisionsList []*model.PaymentRouteDecisions, fieldsInsert ...PaymentRouteDecisionsField) error
	ResolvePaymentRouteDecisions(ctx context.Context, selectFields ...PaymentRouteDecisionsField) (model.PaymentRouteDecisionsList, error)
	ResolvePaymentRouteDecisionsByID(ctx context.Context, primaryID model.PaymentRouteDecisionsPrimaryID, selectFields ...PaymentRouteDecisionsField) (model.PaymentRouteDecisions, error)
	UpdatePaymentRouteDecisionsByID(ctx context.Context, id model.PaymentRouteDecisionsPrimaryID, paymentRouteDecisions *model.PaymentRouteDecisions, paymentRouteDecisionsUpdateFields ...PaymentRouteDecisionsUpdateField) error
	UpdatePaymentRouteDecisionsByFilter(ctx context.Context, filter model.Filter, paymentRouteDecisionsUpdateFields ...PaymentRouteDecisionsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentRouteDecisions(ctx context.Context, paymentRouteDecisionsListMap map[model.PaymentRouteDecisionsPrimaryID]*model.PaymentRouteDecisions, PaymentRouteDecisionssMapUpdateFieldsRequest map[model.PaymentRouteDecisionsPrimaryID]PaymentRouteDecisionsUpdateFieldList) (err error)
	DeletePaymentRouteDecisionsByID(ctx context.Context, id model.PaymentRouteDecisionsPrimaryID) error
	BulkDeletePaymentRouteDecisionsByIDs(ctx context.Context, ids []model.PaymentRouteDecisionsPrimaryID) error
	ResolvePaymentRouteDecisionsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentRouteDecisionsFilterResult, err error)
	IsExistPaymentRouteDecisionsByIDs(ctx context.Context, ids []model.PaymentRouteDecisionsPrimaryID) (exists bool, notFoundIds []model.PaymentRouteDecisionsPrimaryID, err error)
	IsExistPaymentRouteDecisionsByID(ctx context.Context, id model.PaymentRouteDecisionsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
