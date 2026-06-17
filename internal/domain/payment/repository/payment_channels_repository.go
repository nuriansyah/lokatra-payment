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

func composeInsertFieldsAndParamsPaymentChannels(paymentChannelsList []model.PaymentChannels, fieldsInsert ...PaymentChannelsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPaymentChannelsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, paymentChannels := range paymentChannelsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, paymentChannels.Id)
			case selectField.MethodId():
				args = append(args, paymentChannels.MethodId)
			case selectField.Code():
				args = append(args, paymentChannels.Code)
			case selectField.Name():
				args = append(args, paymentChannels.Name)
			case selectField.CountryCode():
				args = append(args, paymentChannels.CountryCode)
			case selectField.Currency():
				args = append(args, paymentChannels.Currency)
			case selectField.Status():
				args = append(args, paymentChannels.Status)
			case selectField.Metadata():
				args = append(args, paymentChannels.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, paymentChannels.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, paymentChannels.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, paymentChannels.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, paymentChannels.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, paymentChannels.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, paymentChannels.MetaDeletedBy)

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

func composePaymentChannelsCompositePrimaryKeyWhere(primaryIDs []model.PaymentChannelsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payment_channels\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPaymentChannelsSelectFields() string {
	fields := NewPaymentChannelsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePaymentChannelsSelectFields(selectFields ...PaymentChannelsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PaymentChannelsField string
type PaymentChannelsFieldList []PaymentChannelsField

type PaymentChannelsSelectFields struct {
}

func (ss PaymentChannelsSelectFields) Id() PaymentChannelsField {
	return PaymentChannelsField("id")
}

func (ss PaymentChannelsSelectFields) MethodId() PaymentChannelsField {
	return PaymentChannelsField("method_id")
}

func (ss PaymentChannelsSelectFields) Code() PaymentChannelsField {
	return PaymentChannelsField("code")
}

func (ss PaymentChannelsSelectFields) Name() PaymentChannelsField {
	return PaymentChannelsField("name")
}

func (ss PaymentChannelsSelectFields) CountryCode() PaymentChannelsField {
	return PaymentChannelsField("country_code")
}

func (ss PaymentChannelsSelectFields) Currency() PaymentChannelsField {
	return PaymentChannelsField("currency")
}

func (ss PaymentChannelsSelectFields) Status() PaymentChannelsField {
	return PaymentChannelsField("status")
}

func (ss PaymentChannelsSelectFields) Metadata() PaymentChannelsField {
	return PaymentChannelsField("metadata")
}

func (ss PaymentChannelsSelectFields) MetaCreatedAt() PaymentChannelsField {
	return PaymentChannelsField("meta_created_at")
}

func (ss PaymentChannelsSelectFields) MetaCreatedBy() PaymentChannelsField {
	return PaymentChannelsField("meta_created_by")
}

func (ss PaymentChannelsSelectFields) MetaUpdatedAt() PaymentChannelsField {
	return PaymentChannelsField("meta_updated_at")
}

func (ss PaymentChannelsSelectFields) MetaUpdatedBy() PaymentChannelsField {
	return PaymentChannelsField("meta_updated_by")
}

func (ss PaymentChannelsSelectFields) MetaDeletedAt() PaymentChannelsField {
	return PaymentChannelsField("meta_deleted_at")
}

func (ss PaymentChannelsSelectFields) MetaDeletedBy() PaymentChannelsField {
	return PaymentChannelsField("meta_deleted_by")
}

func (ss PaymentChannelsSelectFields) All() PaymentChannelsFieldList {
	return []PaymentChannelsField{
		ss.Id(),
		ss.MethodId(),
		ss.Code(),
		ss.Name(),
		ss.CountryCode(),
		ss.Currency(),
		ss.Status(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPaymentChannelsSelectFields() PaymentChannelsSelectFields {
	return PaymentChannelsSelectFields{}
}

type PaymentChannelsUpdateFieldOption struct {
	useIncrement bool
}
type PaymentChannelsUpdateField struct {
	paymentChannelsField PaymentChannelsField
	opt                  PaymentChannelsUpdateFieldOption
	value                interface{}
}
type PaymentChannelsUpdateFieldList []PaymentChannelsUpdateField

func defaultPaymentChannelsUpdateFieldOption() PaymentChannelsUpdateFieldOption {
	return PaymentChannelsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPaymentChannelsOption(useIncrement bool) func(*PaymentChannelsUpdateFieldOption) {
	return func(pcufo *PaymentChannelsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPaymentChannelsUpdateField(field PaymentChannelsField, val interface{}, opts ...func(*PaymentChannelsUpdateFieldOption)) PaymentChannelsUpdateField {
	defaultOpt := defaultPaymentChannelsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PaymentChannelsUpdateField{
		paymentChannelsField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultPaymentChannelsUpdateFields(paymentChannels model.PaymentChannels) (paymentChannelsUpdateFieldList PaymentChannelsUpdateFieldList) {
	selectFields := NewPaymentChannelsSelectFields()
	paymentChannelsUpdateFieldList = append(paymentChannelsUpdateFieldList,
		NewPaymentChannelsUpdateField(selectFields.Id(), paymentChannels.Id),
		NewPaymentChannelsUpdateField(selectFields.MethodId(), paymentChannels.MethodId),
		NewPaymentChannelsUpdateField(selectFields.Code(), paymentChannels.Code),
		NewPaymentChannelsUpdateField(selectFields.Name(), paymentChannels.Name),
		NewPaymentChannelsUpdateField(selectFields.CountryCode(), paymentChannels.CountryCode),
		NewPaymentChannelsUpdateField(selectFields.Currency(), paymentChannels.Currency),
		NewPaymentChannelsUpdateField(selectFields.Status(), paymentChannels.Status),
		NewPaymentChannelsUpdateField(selectFields.Metadata(), paymentChannels.Metadata),
		NewPaymentChannelsUpdateField(selectFields.MetaCreatedAt(), paymentChannels.MetaCreatedAt),
		NewPaymentChannelsUpdateField(selectFields.MetaCreatedBy(), paymentChannels.MetaCreatedBy),
		NewPaymentChannelsUpdateField(selectFields.MetaUpdatedAt(), paymentChannels.MetaUpdatedAt),
		NewPaymentChannelsUpdateField(selectFields.MetaUpdatedBy(), paymentChannels.MetaUpdatedBy),
		NewPaymentChannelsUpdateField(selectFields.MetaDeletedAt(), paymentChannels.MetaDeletedAt),
		NewPaymentChannelsUpdateField(selectFields.MetaDeletedBy(), paymentChannels.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPaymentChannelsCommand(paymentChannelsUpdateFieldList PaymentChannelsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range paymentChannelsUpdateFieldList {
		field := string(updateField.paymentChannelsField)
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

func (repo *RepositoryImpl) BulkCreatePaymentChannels(ctx context.Context, paymentChannelsList []*model.PaymentChannels, fieldsInsert ...PaymentChannelsField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.PaymentChannelsPrimaryID
		paymentChannelsValueList []model.PaymentChannels
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPaymentChannelsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, paymentChannels := range paymentChannelsList {

		primaryIds = append(primaryIds, paymentChannels.ToPaymentChannelsPrimaryID())

		paymentChannelsValueList = append(paymentChannelsValueList, *paymentChannels)
	}

	_, notFoundIds, err := repo.IsExistPaymentChannelsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentChannels] failed checking paymentChannels whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PaymentChannelsPrimaryID{}
		mapNotFoundIds := map[model.PaymentChannelsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "paymentChannels", fmt.Sprintf("paymentChannels with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPaymentChannels(paymentChannelsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(paymentChannelsQueries.insertPaymentChannels, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePaymentChannels] failed exec create paymentChannels query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePaymentChannelsByIDs(ctx context.Context, primaryIDs []model.PaymentChannelsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPaymentChannelsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentChannelsByIDs] failed checking paymentChannels whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentChannels with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_channels\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(paymentChannelsQueries.deletePaymentChannels + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentChannelsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePaymentChannelsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPaymentChannelsByIDs(ctx context.Context, ids []model.PaymentChannelsPrimaryID) (exists bool, notFoundIds []model.PaymentChannelsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payment_channels\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(paymentChannelsQueries.selectPaymentChannels, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentChannelsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PaymentChannelsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentChannelsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PaymentChannelsPrimaryID]bool{}
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

// BulkUpdatePaymentChannels is used to bulk update paymentChannels, by default it will update all field
// if want to update specific field, then fill paymentChannelssMapUpdateFieldsRequest else please fill paymentChannelssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePaymentChannels(ctx context.Context, paymentChannelssMap map[model.PaymentChannelsPrimaryID]*model.PaymentChannels, paymentChannelssMapUpdateFieldsRequest map[model.PaymentChannelsPrimaryID]PaymentChannelsUpdateFieldList) (err error) {
	if len(paymentChannelssMap) == 0 && len(paymentChannelssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		paymentChannelssMapUpdateField map[model.PaymentChannelsPrimaryID]PaymentChannelsUpdateFieldList = map[model.PaymentChannelsPrimaryID]PaymentChannelsUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(paymentChannelssMap) > 0 {
		for id, paymentChannels := range paymentChannelssMap {
			if paymentChannels == nil {
				log.Error().Err(err).Msg("[BulkUpdatePaymentChannels] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			paymentChannelssMapUpdateField[id] = defaultPaymentChannelsUpdateFields(*paymentChannels)
		}
	} else {
		paymentChannelssMapUpdateField = paymentChannelssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePaymentChannelsQuery(paymentChannelssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPaymentChannelsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentChannels] failed checking paymentChannels whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentChannels with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePaymentChannelsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payment_channels\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePaymentChannels] failed exec query")
	}
	return
}

type PaymentChannelsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPaymentChannelsFieldParameter(param string, args ...interface{}) PaymentChannelsFieldParameter {
	return PaymentChannelsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePaymentChannelsQuery(mapPaymentChannelss map[model.PaymentChannelsPrimaryID]PaymentChannelsUpdateFieldList, asTableValues string) (primaryIDs []model.PaymentChannelsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PaymentChannelsPrimaryID]map[string]interface{}{}
	paymentChannelsSelectFields := NewPaymentChannelsSelectFields()
	for id, updateFields := range mapPaymentChannelss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.paymentChannelsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPaymentChannelss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPaymentChannelsFieldType(updateField.paymentChannelsField)))
			args = append(args, fields[string(updateField.paymentChannelsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.paymentChannelsField))
		if updateField.paymentChannelsField == paymentChannelsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.paymentChannelsField, asTableValues, updateField.paymentChannelsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.paymentChannelsField,
				"\"payment_channels\"", updateField.paymentChannelsField,
				asTableValues, updateField.paymentChannelsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePaymentChannelsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PaymentChannelsPrimaryID, asTableValue string) (whereQry string) {
	paymentChannelsSelectFields := NewPaymentChannelsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payment_channels\".\"id\" = %s.\"id\"::"+GetPaymentChannelsFieldType(paymentChannelsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPaymentChannelsFieldType(paymentChannelsField PaymentChannelsField) string {
	selectPaymentChannelsFields := NewPaymentChannelsSelectFields()
	switch paymentChannelsField {

	case selectPaymentChannelsFields.Id():
		return "uuid"

	case selectPaymentChannelsFields.MethodId():
		return "uuid"

	case selectPaymentChannelsFields.Code():
		return "text"

	case selectPaymentChannelsFields.Name():
		return "text"

	case selectPaymentChannelsFields.CountryCode():
		return "text"

	case selectPaymentChannelsFields.Currency():
		return "text"

	case selectPaymentChannelsFields.Status():
		return "payment_channel_status_enum"

	case selectPaymentChannelsFields.Metadata():
		return "jsonb"

	case selectPaymentChannelsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPaymentChannelsFields.MetaCreatedBy():
		return "uuid"

	case selectPaymentChannelsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPaymentChannelsFields.MetaUpdatedBy():
		return "uuid"

	case selectPaymentChannelsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPaymentChannelsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePaymentChannels(ctx context.Context, paymentChannels *model.PaymentChannels, fieldsInsert ...PaymentChannelsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPaymentChannelsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PaymentChannelsPrimaryID{
		Id: paymentChannels.Id,
	}
	exists, err := repo.IsExistPaymentChannelsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentChannels] failed checking paymentChannels whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "paymentChannels", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPaymentChannels([]model.PaymentChannels{*paymentChannels}, fieldsInsert...)
	commandQuery := fmt.Sprintf(paymentChannelsQueries.insertPaymentChannels, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePaymentChannels] failed exec create paymentChannels query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePaymentChannelsByID(ctx context.Context, primaryID model.PaymentChannelsPrimaryID) (err error) {
	exists, err := repo.IsExistPaymentChannelsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentChannelsByID] failed checking paymentChannels whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentChannels with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePaymentChannelsCompositePrimaryKeyWhere([]model.PaymentChannelsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(paymentChannelsQueries.deletePaymentChannels + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePaymentChannelsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentChannelsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentChannelsFilterResult, err error) {
	query, args, err := composePaymentChannelsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentChannelsByFilter] failed compose paymentChannels filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentChannelsByFilter] failed get paymentChannels by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePaymentChannelsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PaymentChannelsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePaymentChannelsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePaymentChannelsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePaymentChannelsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPaymentChannelsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 14+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["method_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_id\"")
			selectedColumns["method_id"] = struct{}{}
		}
		if _, selected := selectedColumns["code"]; !selected {
			selectColumns = append(selectColumns, "base.\"code\"")
			selectedColumns["code"] = struct{}{}
		}
		if _, selected := selectedColumns["name"]; !selected {
			selectColumns = append(selectColumns, "base.\"name\"")
			selectedColumns["name"] = struct{}{}
		}
		if _, selected := selectedColumns["country_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"country_code\"")
			selectedColumns["country_code"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
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

type paymentChannelsFilterPlaceholder struct {
	index int
}

func (p *paymentChannelsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePaymentChannelsFilterPredicate(filterField model.FilterField, placeholders *paymentChannelsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPaymentChannelsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePaymentChannelsFilterSQLExpr(spec)
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

func composePaymentChannelsFilterGroup(group model.FilterGroup, placeholders *paymentChannelsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePaymentChannelsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePaymentChannelsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePaymentChannelsFilterWhereQueries(filter model.Filter, placeholders *paymentChannelsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePaymentChannelsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePaymentChannelsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePaymentChannelsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePaymentChannelsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePaymentChannelsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePaymentChannelsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := paymentChannelsFilterPlaceholder{index: 1}
	whereQueries, err := composePaymentChannelsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPaymentChannelsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePaymentChannelsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePaymentChannelsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payment_channels\" base%s", strings.Join(selectColumns, ","), composePaymentChannelsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPaymentChannelsByID(ctx context.Context, primaryID model.PaymentChannelsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePaymentChannelsCompositePrimaryKeyWhere([]model.PaymentChannelsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", paymentChannelsQueries.selectCountPaymentChannels, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPaymentChannelsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentChannels(ctx context.Context, selectFields ...PaymentChannelsField) (paymentChannelsList model.PaymentChannelsList, err error) {
	var (
		defaultPaymentChannelsSelectFields = defaultPaymentChannelsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentChannelsSelectFields = composePaymentChannelsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(paymentChannelsQueries.selectPaymentChannels, defaultPaymentChannelsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &paymentChannelsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePaymentChannels] failed get paymentChannels list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePaymentChannelsByID(ctx context.Context, primaryID model.PaymentChannelsPrimaryID, selectFields ...PaymentChannelsField) (paymentChannels model.PaymentChannels, err error) {
	var (
		defaultPaymentChannelsSelectFields = defaultPaymentChannelsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPaymentChannelsSelectFields = composePaymentChannelsSelectFields(selectFields...)
	}
	whereQry, params := composePaymentChannelsCompositePrimaryKeyWhere([]model.PaymentChannelsPrimaryID{primaryID})
	query := fmt.Sprintf(paymentChannelsQueries.selectPaymentChannels+" WHERE "+whereQry, defaultPaymentChannelsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &paymentChannels, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("paymentChannels with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePaymentChannelsByID] failed get paymentChannels")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePaymentChannelsByID(ctx context.Context, primaryID model.PaymentChannelsPrimaryID, paymentChannels *model.PaymentChannels, paymentChannelsUpdateFields ...PaymentChannelsUpdateField) (err error) {
	exists, err := repo.IsExistPaymentChannelsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentChannels] failed checking paymentChannels whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("paymentChannels with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if paymentChannels == nil {
		if len(paymentChannelsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePaymentChannelsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		paymentChannels = &model.PaymentChannels{}
	}
	var (
		defaultPaymentChannelsUpdateFields = defaultPaymentChannelsUpdateFields(*paymentChannels)
		tempUpdateField                    PaymentChannelsUpdateFieldList
		selectFields                       = NewPaymentChannelsSelectFields()
	)
	if len(paymentChannelsUpdateFields) > 0 {
		for _, updateField := range paymentChannelsUpdateFields {
			if updateField.paymentChannelsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPaymentChannelsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePaymentChannelsCompositePrimaryKeyWhere([]model.PaymentChannelsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPaymentChannelsCommand(defaultPaymentChannelsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(paymentChannelsQueries.updatePaymentChannels+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentChannels] error when try to update paymentChannels by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePaymentChannelsByFilter(ctx context.Context, filter model.Filter, paymentChannelsUpdateFields ...PaymentChannelsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(paymentChannelsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PaymentChannelsUpdateFieldList
		selectFields = NewPaymentChannelsSelectFields()
	)
	for _, updateField := range paymentChannelsUpdateFields {
		if updateField.paymentChannelsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPaymentChannelsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := paymentChannelsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePaymentChannelsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payment_channels\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentChannelsByFilter] error when try to update paymentChannels by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePaymentChannelsByFilter] failed get rows affected")
	}
	return
}

var (
	paymentChannelsQueries = struct {
		selectPaymentChannels      string
		selectCountPaymentChannels string
		deletePaymentChannels      string
		updatePaymentChannels      string
		insertPaymentChannels      string
	}{
		selectPaymentChannels:      "SELECT %s FROM \"payment_channels\"",
		selectCountPaymentChannels: "SELECT COUNT(\"id\") FROM \"payment_channels\"",
		deletePaymentChannels:      "DELETE FROM \"payment_channels\"",
		updatePaymentChannels:      "UPDATE \"payment_channels\" SET %s ",
		insertPaymentChannels:      "INSERT INTO \"payment_channels\" %s VALUES %s",
	}
)

type PaymentChannelsRepository interface {
	CreatePaymentChannels(ctx context.Context, paymentChannels *model.PaymentChannels, fieldsInsert ...PaymentChannelsField) error
	BulkCreatePaymentChannels(ctx context.Context, paymentChannelsList []*model.PaymentChannels, fieldsInsert ...PaymentChannelsField) error
	ResolvePaymentChannels(ctx context.Context, selectFields ...PaymentChannelsField) (model.PaymentChannelsList, error)
	ResolvePaymentChannelsByID(ctx context.Context, primaryID model.PaymentChannelsPrimaryID, selectFields ...PaymentChannelsField) (model.PaymentChannels, error)
	UpdatePaymentChannelsByID(ctx context.Context, id model.PaymentChannelsPrimaryID, paymentChannels *model.PaymentChannels, paymentChannelsUpdateFields ...PaymentChannelsUpdateField) error
	UpdatePaymentChannelsByFilter(ctx context.Context, filter model.Filter, paymentChannelsUpdateFields ...PaymentChannelsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePaymentChannels(ctx context.Context, paymentChannelsListMap map[model.PaymentChannelsPrimaryID]*model.PaymentChannels, PaymentChannelssMapUpdateFieldsRequest map[model.PaymentChannelsPrimaryID]PaymentChannelsUpdateFieldList) (err error)
	DeletePaymentChannelsByID(ctx context.Context, id model.PaymentChannelsPrimaryID) error
	BulkDeletePaymentChannelsByIDs(ctx context.Context, ids []model.PaymentChannelsPrimaryID) error
	ResolvePaymentChannelsByFilter(ctx context.Context, filter model.Filter) (result []model.PaymentChannelsFilterResult, err error)
	IsExistPaymentChannelsByIDs(ctx context.Context, ids []model.PaymentChannelsPrimaryID) (exists bool, notFoundIds []model.PaymentChannelsPrimaryID, err error)
	IsExistPaymentChannelsByID(ctx context.Context, id model.PaymentChannelsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
