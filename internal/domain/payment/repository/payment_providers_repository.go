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

func composeInsertFieldsAndParamsPaymentProviders(paymentProvidersList []model.PaymentProviders, fieldsInsert ...PaymentProvidersField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentProvidersSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentProviders := range paymentProvidersList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentProviders.Id)
			case selectField.Code():
				args = append(args, paymentProviders.Code)
			case selectField.Name():
				args = append(args, paymentProviders.Name)
			case selectField.ProviderType():
				args = append(args, paymentProviders.ProviderType)
			case selectField.Status():
				args = append(args, paymentProviders.Status)
			case selectField.SupportsRefund():
				args = append(args, paymentProviders.SupportsRefund)
			case selectField.SupportsPartialRefund():
				args = append(args, paymentProviders.SupportsPartialRefund)
			case selectField.SupportsAuthorization():
				args = append(args, paymentProviders.SupportsAuthorization)
			case selectField.SupportsCapture():
				args = append(args, paymentProviders.SupportsCapture)
			case selectField.SupportsVoid():
				args = append(args, paymentProviders.SupportsVoid)
			case selectField.SupportsWebhook():
				args = append(args, paymentProviders.SupportsWebhook)
			case selectField.Metadata():
				args = append(args, paymentProviders.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentProviders.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentProviders.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentProviders.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentProviders.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentProviders.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentProviders.MetaDeletedBy)

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

func composePaymentProvidersCompositePrimaryKeyWhere(primaryIDs []model.PaymentProvidersPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_providers\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentProvidersSelectFields() string {
	fields := NewPaymentProvidersSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentProvidersSelectFields(selectFields ...PaymentProvidersField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentProvidersField string
type PaymentProvidersFieldList []PaymentProvidersField

type PaymentProvidersSelectFields struct {
}

func (ss PaymentProvidersSelectFields) Id() PaymentProvidersField {
	return PaymentProvidersField("id")
}

func (ss PaymentProvidersSelectFields) Code() PaymentProvidersField {
	return PaymentProvidersField("code")
}

func (ss PaymentProvidersSelectFields) Name() PaymentProvidersField {
	return PaymentProvidersField("name")
}

func (ss PaymentProvidersSelectFields) ProviderType() PaymentProvidersField {
	return PaymentProvidersField("provider_type")
}

func (ss PaymentProvidersSelectFields) Status() PaymentProvidersField {
	return PaymentProvidersField("status")
}

func (ss PaymentProvidersSelectFields) SupportsRefund() PaymentProvidersField {
	return PaymentProvidersField("supports_refund")
}

func (ss PaymentProvidersSelectFields) SupportsPartialRefund() PaymentProvidersField {
	return PaymentProvidersField("supports_partial_refund")
}

func (ss PaymentProvidersSelectFields) SupportsAuthorization() PaymentProvidersField {
	return PaymentProvidersField("supports_authorization")
}

func (ss PaymentProvidersSelectFields) SupportsCapture() PaymentProvidersField {
	return PaymentProvidersField("supports_capture")
}

func (ss PaymentProvidersSelectFields) SupportsVoid() PaymentProvidersField {
	return PaymentProvidersField("supports_void")
}

func (ss PaymentProvidersSelectFields) SupportsWebhook() PaymentProvidersField {
	return PaymentProvidersField("supports_webhook")
}

func (ss PaymentProvidersSelectFields) Metadata() PaymentProvidersField {
	return PaymentProvidersField("metadata")
}

func (ss PaymentProvidersSelectFields) MetaCreatedAt() PaymentProvidersField {
	return PaymentProvidersField("meta_created_at")
}

func (ss PaymentProvidersSelectFields) MetaCreatedBy() PaymentProvidersField {
	return PaymentProvidersField("meta_created_by")
}

func (ss PaymentProvidersSelectFields) MetaUpdatedAt() PaymentProvidersField {
	return PaymentProvidersField("meta_updated_at")
}

func (ss PaymentProvidersSelectFields) MetaUpdatedBy() PaymentProvidersField {
	return PaymentProvidersField("meta_updated_by")
}

func (ss PaymentProvidersSelectFields) MetaDeletedAt() PaymentProvidersField {
	return PaymentProvidersField("meta_deleted_at")
}

func (ss PaymentProvidersSelectFields) MetaDeletedBy() PaymentProvidersField {
	return PaymentProvidersField("meta_deleted_by")
}

func (ss PaymentProvidersSelectFields) All() PaymentProvidersFieldList {
	return []PaymentProvidersField{
		ss.Id(),
		ss.Code(),
		ss.Name(),
		ss.ProviderType(),
		ss.Status(),
		ss.SupportsRefund(),
		ss.SupportsPartialRefund(),
		ss.SupportsAuthorization(),
		ss.SupportsCapture(),
		ss.SupportsVoid(),
		ss.SupportsWebhook(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentProvidersSelectFields() PaymentProvidersSelectFields {
	return PaymentProvidersSelectFields{}
}

type PaymentProvidersUpdateFieldOption struct {
	useIncrement bool
}
type PaymentProvidersUpdateField struct {
	paymentProvidersField PaymentProvidersField
	opt                   PaymentProvidersUpdateFieldOption
	value                 interface{}
}
type PaymentProvidersUpdateFieldList []PaymentProvidersUpdateField

func defaultPaymentProvidersUpdateFieldOption() PaymentProvidersUpdateFieldOption {
	return PaymentProvidersUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentProvidersOption(useIncrement bool) func(*PaymentProvidersUpdateFieldOption) {
	return func(pcufo *PaymentProvidersUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentProvidersUpdateField(field PaymentProvidersField, val interface{}, opts ...func(*PaymentProvidersUpdateFieldOption)) PaymentProvidersUpdateField {
	defaultOpt := defaultPaymentProvidersUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentProvidersUpdateField{
		paymentProvidersField: field,
		value:                 val,
		opt:                   defaultOpt,
	}
}
func defaultPaymentProvidersUpdateFields(paymentProviders model.PaymentProviders) (paymentProvidersUpdateFieldList PaymentProvidersUpdateFieldList) {
	selectFields := NewPaymentProvidersSelectFields()
	paymentProvidersUpdateFieldList = append(paymentProvidersUpdateFieldList,
		NewPaymentProvidersUpdateField(selectFields.Id(), paymentProviders.Id),
		NewPaymentProvidersUpdateField(selectFields.Code(), paymentProviders.Code),
		NewPaymentProvidersUpdateField(selectFields.Name(), paymentProviders.Name),
		NewPaymentProvidersUpdateField(selectFields.ProviderType(), paymentProviders.ProviderType),
		NewPaymentProvidersUpdateField(selectFields.Status(), paymentProviders.Status),
		NewPaymentProvidersUpdateField(selectFields.SupportsRefund(), paymentProviders.SupportsRefund),
		NewPaymentProvidersUpdateField(selectFields.SupportsPartialRefund(), paymentProviders.SupportsPartialRefund),
		NewPaymentProvidersUpdateField(selectFields.SupportsAuthorization(), paymentProviders.SupportsAuthorization),
		NewPaymentProvidersUpdateField(selectFields.SupportsCapture(), paymentProviders.SupportsCapture),
		NewPaymentProvidersUpdateField(selectFields.SupportsVoid(), paymentProviders.SupportsVoid),
		NewPaymentProvidersUpdateField(selectFields.SupportsWebhook(), paymentProviders.SupportsWebhook),
		NewPaymentProvidersUpdateField(selectFields.Metadata(), paymentProviders.Metadata),
		NewPaymentProvidersUpdateField(selectFields.MetaCreatedAt(), paymentProviders.MetaCreatedAt),
		NewPaymentProvidersUpdateField(selectFields.MetaCreatedBy(), paymentProviders.MetaCreatedBy),
		NewPaymentProvidersUpdateField(selectFields.MetaUpdatedAt(), paymentProviders.MetaUpdatedAt),
		NewPaymentProvidersUpdateField(selectFields.MetaUpdatedBy(), paymentProviders.MetaUpdatedBy),
		NewPaymentProvidersUpdateField(selectFields.MetaDeletedAt(), paymentProviders.MetaDeletedAt),
		NewPaymentProvidersUpdateField(selectFields.MetaDeletedBy(), paymentProviders.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentProvidersCommand(paymentProvidersUpdateFieldList PaymentProvidersUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentProvidersUpdateFieldList {
		field := string(updateField.paymentProvidersField)
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

func (repo *RepositoryImpl) BulkCreatePaymentProviders(ctx context.Context, paymentProvidersList []*model.PaymentProviders, fieldsInsert ...PaymentProvidersField) (err error) {
	var (
		fieldsStr                 string
		valueListStr              []string
		argsList                  []interface{}
		primaryIds                []model.PaymentProvidersPrimaryID
		paymentProvidersValueList []model.PaymentProviders
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentProvidersSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentProviders := range paymentProvidersList {

		primaryIds = append(primaryIds, paymentProviders.ToPaymentProvidersPrimaryID())

		paymentProvidersValueList = append(paymentProvidersValueList, *paymentProviders)
	}

	_, notFoundIds, err := repo.IsExistPaymentProvidersByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentProviders] failed checking paymentProviders whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentProvidersPrimaryID{}
		mapNotFoundIds := map[model.PaymentProvidersPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentProviders", fmt.Sprintf("paymentProviders with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentProviders(paymentProvidersValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentProvidersQueries.insertPaymentProviders, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentProviders] failed exec create paymentProviders query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentProvidersByIDs(ctx context.Context, primaryIDs []model.PaymentProvidersPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentProvidersByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentProvidersByIDs] failed checking paymentProviders whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentProviders with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_providers\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := paymentProvidersQueries.deletePaymentProviders + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentProvidersByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentProvidersByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentProvidersByIDs(ctx context.Context, ids []model.PaymentProvidersPrimaryID) (exists bool, notFoundIds []model.PaymentProvidersPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_providers\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentProvidersQueries.selectPaymentProviders, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentProvidersByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentProvidersPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentProvidersByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentProvidersPrimaryID]bool{}
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

// BulkUpdatePaymentProviders is used to bulk update paymentProviders, by default it will update all field
// if want to update specific field, then fill paymentProviderssMapUpdateFieldsRequest else please fill paymentProviderssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentProviders(ctx context.Context, paymentProviderssMap map[model.PaymentProvidersPrimaryID]*model.PaymentProviders, paymentProviderssMapUpdateFieldsRequest map[model.PaymentProvidersPrimaryID]PaymentProvidersUpdateFieldList) (err error) {
	if len(paymentProviderssMap) == 0 && len(paymentProviderssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentProviderssMapUpdateField map[model.PaymentProvidersPrimaryID]PaymentProvidersUpdateFieldList = map[model.PaymentProvidersPrimaryID]PaymentProvidersUpdateFieldList{}
		asTableValues                   string                                                              = "myvalues"
	)

	if len(paymentProviderssMap) > 0 {
		for id, paymentProviders := range paymentProviderssMap {
			if paymentProviders == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentProviders] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentProviderssMapUpdateField[id] = defaultPaymentProvidersUpdateFields(*paymentProviders)
		}
	} else {
		paymentProviderssMapUpdateField = paymentProviderssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentProvidersQuery(paymentProviderssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentProvidersByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentProviders] failed checking paymentProviders whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentProviders with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentProvidersCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_providers\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentProviders] failed exec query")
	}
	return
}

type PaymentProvidersFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentProvidersFieldParameter(param string, args ...interface{}) PaymentProvidersFieldParameter {
	return PaymentProvidersFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentProvidersQuery(mapPaymentProviderss map[model.PaymentProvidersPrimaryID]PaymentProvidersUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentProvidersPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentProvidersPrimaryID]map[string]interface{}{}
	paymentProvidersSelectFields := NewPaymentProvidersSelectFields()
	for id, updateFields := range mapPaymentProviderss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentProvidersField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentProviderss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentProvidersFieldType(updateField.paymentProvidersField)))
			args = append(args, fields[string(updateField.paymentProvidersField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentProvidersField))
		if updateField.paymentProvidersField == paymentProvidersSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentProvidersField, asTableValues, updateField.paymentProvidersField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentProvidersField,
				"\"payment_providers\"", updateField.paymentProvidersField,
				asTableValues, updateField.paymentProvidersField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentProvidersCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentProvidersPrimaryID, asTableValue string) (whereQry string) {
	paymentProvidersSelectFields := NewPaymentProvidersSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_providers\".\"id\" = %s.\"id\"::"+GetPaymentProvidersFieldType(paymentProvidersSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentProvidersFieldType(paymentProvidersField PaymentProvidersField) string {
	selectPaymentProvidersFields := NewPaymentProvidersSelectFields()
	switch paymentProvidersField {

	case selectPaymentProvidersFields.Id():
		return "uuid"

	case selectPaymentProvidersFields.Code():
		return "text"

	case selectPaymentProvidersFields.Name():
		return "text"

	case selectPaymentProvidersFields.ProviderType():
		return "provider_type_enum"

	case selectPaymentProvidersFields.Status():
		return "provider_status_enum"

	case selectPaymentProvidersFields.SupportsRefund():
		return "bool"

	case selectPaymentProvidersFields.SupportsPartialRefund():
		return "bool"

	case selectPaymentProvidersFields.SupportsAuthorization():
		return "bool"

	case selectPaymentProvidersFields.SupportsCapture():
		return "bool"

	case selectPaymentProvidersFields.SupportsVoid():
		return "bool"

	case selectPaymentProvidersFields.SupportsWebhook():
		return "bool"

	case selectPaymentProvidersFields.Metadata():
		return "jsonb"

	case selectPaymentProvidersFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentProvidersFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentProvidersFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentProvidersFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentProvidersFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentProvidersFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentProviders(ctx context.Context, paymentProviders *model.PaymentProviders, fieldsInsert ...PaymentProvidersField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentProvidersSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentProvidersPrimaryID{
		Id: paymentProviders.Id,
	}
	exists, err := repo.IsExistPaymentProvidersByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentProviders] failed checking paymentProviders whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentProviders", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentProviders([]model.PaymentProviders{*paymentProviders}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentProvidersQueries.insertPaymentProviders, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentProviders] failed exec create paymentProviders query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentProvidersByID(ctx context.Context, primaryID model.PaymentProvidersPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentProvidersByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentProvidersByID] failed checking paymentProviders whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentProviders with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentProvidersCompositePrimaryKeyWhere([]model.PaymentProvidersPrimaryID{primaryID})
	commandQuery := paymentProvidersQueries.deletePaymentProviders + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentProvidersByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentProvidersByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentProvidersFilterResult, err error) {
	query, args, err := composePaymentProvidersFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentProvidersByFilter] failed compose paymentProviders filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentProvidersByFilter] failed get paymentProviders by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentProvidersFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentProvidersFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentProvidersFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentProvidersSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentProvidersFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentProvidersFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 18+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["code"]; !selected {
			selectColumns = append(selectColumns, "base.\"code\"")
			selectedColumns["code"] = struct{}{}
		}
		if _, selected := selectedColumns["name"]; !selected {
			selectColumns = append(selectColumns, "base.\"name\"")
			selectedColumns["name"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_type\"")
			selectedColumns["provider_type"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["supports_refund"]; !selected {
			selectColumns = append(selectColumns, "base.\"supports_refund\"")
			selectedColumns["supports_refund"] = struct{}{}
		}
		if _, selected := selectedColumns["supports_partial_refund"]; !selected {
			selectColumns = append(selectColumns, "base.\"supports_partial_refund\"")
			selectedColumns["supports_partial_refund"] = struct{}{}
		}
		if _, selected := selectedColumns["supports_authorization"]; !selected {
			selectColumns = append(selectColumns, "base.\"supports_authorization\"")
			selectedColumns["supports_authorization"] = struct{}{}
		}
		if _, selected := selectedColumns["supports_capture"]; !selected {
			selectColumns = append(selectColumns, "base.\"supports_capture\"")
			selectedColumns["supports_capture"] = struct{}{}
		}
		if _, selected := selectedColumns["supports_void"]; !selected {
			selectColumns = append(selectColumns, "base.\"supports_void\"")
			selectedColumns["supports_void"] = struct{}{}
		}
		if _, selected := selectedColumns["supports_webhook"]; !selected {
			selectColumns = append(selectColumns, "base.\"supports_webhook\"")
			selectedColumns["supports_webhook"] = struct{}{}
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

type paymentProvidersFilterPlaceholder struct {
	index int
}

func (p *paymentProvidersFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentProvidersFilterPredicate(filterField model.FilterField, placeholders *paymentProvidersFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentProvidersFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentProvidersFilterSQLExpr(spec)
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

func composePaymentProvidersFilterGroup(group model.FilterGroup, placeholders *paymentProvidersFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentProvidersFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentProvidersFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentProvidersFilterWhereQueries(filter model.Filter, placeholders *paymentProvidersFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentProvidersFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentProvidersFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentProvidersFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentProvidersFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentProvidersSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentProvidersFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentProvidersFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentProvidersFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentProvidersFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentProvidersFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentProvidersSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_providers\" base%s", strings.Join(selectColumns, ","), composePaymentProvidersFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentProvidersByID(ctx context.Context, primaryID model.PaymentProvidersPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentProvidersCompositePrimaryKeyWhere([]model.PaymentProvidersPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentProvidersQueries.selectCountPaymentProviders, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentProvidersByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentProviders(ctx context.Context, selectFields ...PaymentProvidersField) (paymentProvidersList model.PaymentProvidersList, err error) {
	var (
		defaultPaymentProvidersSelectFields = defaultPaymentProvidersSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentProvidersSelectFields = composePaymentProvidersSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentProvidersQueries.selectPaymentProviders, defaultPaymentProvidersSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentProvidersList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentProviders] failed get paymentProviders list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentProvidersByID(ctx context.Context, primaryID model.PaymentProvidersPrimaryID, selectFields ...PaymentProvidersField) (paymentProviders model.PaymentProviders, err error) {
	var (
		defaultPaymentProvidersSelectFields = defaultPaymentProvidersSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentProvidersSelectFields = composePaymentProvidersSelectFields(selectFields...)
	}
	whereQry, params := composePaymentProvidersCompositePrimaryKeyWhere([]model.PaymentProvidersPrimaryID{primaryID})
	query := fmt.Sprintf(paymentProvidersQueries.selectPaymentProviders+" WHERE "+whereQry, defaultPaymentProvidersSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentProviders, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentProviders with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentProvidersByID] failed get paymentProviders")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentProvidersByID(ctx context.Context, primaryID model.PaymentProvidersPrimaryID, paymentProviders *model.PaymentProviders, paymentProvidersUpdateFields ...PaymentProvidersUpdateField) (err error) {
	exists, err := repo.IsExistPaymentProvidersByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentProviders] failed checking paymentProviders whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentProviders with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentProviders == nil {
		if len(paymentProvidersUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentProvidersByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentProviders = &model.PaymentProviders{}
	}
	var (
		defaultPaymentProvidersUpdateFields = defaultPaymentProvidersUpdateFields(*paymentProviders)
		tempUpdateField                     PaymentProvidersUpdateFieldList
		selectFields                        = NewPaymentProvidersSelectFields()
	)
	if len(paymentProvidersUpdateFields) > 0 {
		for _, updateField := range paymentProvidersUpdateFields {
			if updateField.paymentProvidersField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentProvidersUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentProvidersCompositePrimaryKeyWhere([]model.PaymentProvidersPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentProvidersCommand(defaultPaymentProvidersUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentProvidersQueries.updatePaymentProviders+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentProviders] error when try to update paymentProviders by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentProvidersByFilter(ctx context.Context, filter model.Filter, paymentProvidersUpdateFields ...PaymentProvidersUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentProvidersUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentProvidersUpdateFieldList
		selectFields = NewPaymentProvidersSelectFields()
	)
	for _, updateField := range paymentProvidersUpdateFields {
		if updateField.paymentProvidersField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentProvidersCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentProvidersFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentProvidersFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_providers\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentProvidersByFilter] error when try to update paymentProviders by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentProvidersByFilter] failed get rows affected")
	}
	return
}

var (
	paymentProvidersQueries = struct {
		selectPaymentProviders      string
		selectCountPaymentProviders string
		deletePaymentProviders      string
		updatePaymentProviders      string
		insertPaymentProviders      string
	}{
		selectPaymentProviders:      "SELECT %s FROM \"payment_providers\"",
		selectCountPaymentProviders: "SELECT COUNT(\"id\") FROM \"payment_providers\"",
		deletePaymentProviders:      "DELETE FROM \"payment_providers\"",
		updatePaymentProviders:      "UPDATE \"payment_providers\" SET %s ",
		insertPaymentProviders:      "INSERT INTO \"payment_providers\" %s VALUES %s",
	}
)

type PaymentProvidersRepository interface {
	CreatePaymentProviders(ctx context.Context, paymentProviders *model.PaymentProviders, fieldsInsert ...PaymentProvidersField) error
	BulkCreatePaymentProviders(ctx context.Context, paymentProvidersList []*model.PaymentProviders, fieldsInsert ...PaymentProvidersField) error
	ResolvePaymentProviders(ctx context.Context, selectFields ...PaymentProvidersField) (model.PaymentProvidersList, error)
	ResolvePaymentProvidersByID(ctx context.Context, primaryID model.PaymentProvidersPrimaryID, selectFields ...PaymentProvidersField) (model.PaymentProviders, error)
	UpdatePaymentProvidersByID(ctx context.Context, id model.PaymentProvidersPrimaryID, paymentProviders *model.PaymentProviders, paymentProvidersUpdateFields ...PaymentProvidersUpdateField) error
	UpdatePaymentProvidersByFilter(ctx context.Context, filter model.Filter, paymentProvidersUpdateFields ...PaymentProvidersUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentProviders(ctx context.Context, paymentProvidersListMap map[model.PaymentProvidersPrimaryID]*model.PaymentProviders, PaymentProviderssMapUpdateFieldsRequest map[model.PaymentProvidersPrimaryID]PaymentProvidersUpdateFieldList) (err error)
	DeletePaymentProvidersByID(ctx context.Context, id model.PaymentProvidersPrimaryID) error
	BulkDeletePaymentProvidersByIDs(ctx context.Context, ids []model.PaymentProvidersPrimaryID) error
	ResolvePaymentProvidersByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentProvidersFilterResult, err error)
	IsExistPaymentProvidersByIDs(ctx context.Context, ids []model.PaymentProvidersPrimaryID) (exists bool, notFoundIds []model.PaymentProvidersPrimaryID, err error)
	IsExistPaymentProvidersByID(ctx context.Context, id model.PaymentProvidersPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
