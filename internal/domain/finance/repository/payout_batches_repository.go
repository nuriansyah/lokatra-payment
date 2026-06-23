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

func composeInsertFieldsAndParamsPayoutBatches(payoutBatchesList []model.PayoutBatches, fieldsInsert ...PayoutBatchesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPayoutBatchesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, payoutBatches := range payoutBatchesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, payoutBatches.Id)
			case selectField.PayoutBatchCode():
				args = append(args, payoutBatches.PayoutBatchCode)
			case selectField.ProviderAccountId():
				args = append(args, payoutBatches.ProviderAccountId)
			case selectField.ScheduledFor():
				args = append(args, payoutBatches.ScheduledFor)
			case selectField.CurrencyCode():
				args = append(args, payoutBatches.CurrencyCode)
			case selectField.BatchStatus():
				args = append(args, payoutBatches.BatchStatus)
			case selectField.TotalCount():
				args = append(args, payoutBatches.TotalCount)
			case selectField.TotalAmount():
				args = append(args, payoutBatches.TotalAmount)
			case selectField.Metadata():
				args = append(args, payoutBatches.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, payoutBatches.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, payoutBatches.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, payoutBatches.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, payoutBatches.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, payoutBatches.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, payoutBatches.MetaDeletedBy)

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

func composePayoutBatchesCompositePrimaryKeyWhere(primaryIDs []model.PayoutBatchesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payout_batches\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPayoutBatchesSelectFields() string {
	fields := NewPayoutBatchesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePayoutBatchesSelectFields(selectFields ...PayoutBatchesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PayoutBatchesField string
type PayoutBatchesFieldList []PayoutBatchesField

type PayoutBatchesSelectFields struct {
}

func (ss PayoutBatchesSelectFields) Id() PayoutBatchesField {
	return PayoutBatchesField("id")
}

func (ss PayoutBatchesSelectFields) PayoutBatchCode() PayoutBatchesField {
	return PayoutBatchesField("payout_batch_code")
}

func (ss PayoutBatchesSelectFields) ProviderAccountId() PayoutBatchesField {
	return PayoutBatchesField("provider_account_id")
}

func (ss PayoutBatchesSelectFields) ScheduledFor() PayoutBatchesField {
	return PayoutBatchesField("scheduled_for")
}

func (ss PayoutBatchesSelectFields) CurrencyCode() PayoutBatchesField {
	return PayoutBatchesField("currency_code")
}

func (ss PayoutBatchesSelectFields) BatchStatus() PayoutBatchesField {
	return PayoutBatchesField("batch_status")
}

func (ss PayoutBatchesSelectFields) TotalCount() PayoutBatchesField {
	return PayoutBatchesField("total_count")
}

func (ss PayoutBatchesSelectFields) TotalAmount() PayoutBatchesField {
	return PayoutBatchesField("total_amount")
}

func (ss PayoutBatchesSelectFields) Metadata() PayoutBatchesField {
	return PayoutBatchesField("metadata")
}

func (ss PayoutBatchesSelectFields) MetaCreatedAt() PayoutBatchesField {
	return PayoutBatchesField("meta_created_at")
}

func (ss PayoutBatchesSelectFields) MetaCreatedBy() PayoutBatchesField {
	return PayoutBatchesField("meta_created_by")
}

func (ss PayoutBatchesSelectFields) MetaUpdatedAt() PayoutBatchesField {
	return PayoutBatchesField("meta_updated_at")
}

func (ss PayoutBatchesSelectFields) MetaUpdatedBy() PayoutBatchesField {
	return PayoutBatchesField("meta_updated_by")
}

func (ss PayoutBatchesSelectFields) MetaDeletedAt() PayoutBatchesField {
	return PayoutBatchesField("meta_deleted_at")
}

func (ss PayoutBatchesSelectFields) MetaDeletedBy() PayoutBatchesField {
	return PayoutBatchesField("meta_deleted_by")
}

func (ss PayoutBatchesSelectFields) All() PayoutBatchesFieldList {
	return []PayoutBatchesField{
		ss.Id(),
		ss.PayoutBatchCode(),
		ss.ProviderAccountId(),
		ss.ScheduledFor(),
		ss.CurrencyCode(),
		ss.BatchStatus(),
		ss.TotalCount(),
		ss.TotalAmount(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPayoutBatchesSelectFields() PayoutBatchesSelectFields {
	return PayoutBatchesSelectFields{}
}

type PayoutBatchesUpdateFieldOption struct {
	useIncrement bool
}
type PayoutBatchesUpdateField struct {
	payoutBatchesField PayoutBatchesField
	opt                PayoutBatchesUpdateFieldOption
	value              interface{}
}
type PayoutBatchesUpdateFieldList []PayoutBatchesUpdateField

func defaultPayoutBatchesUpdateFieldOption() PayoutBatchesUpdateFieldOption {
	return PayoutBatchesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPayoutBatchesOption(useIncrement bool) func(*PayoutBatchesUpdateFieldOption) {
	return func(pcufo *PayoutBatchesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPayoutBatchesUpdateField(field PayoutBatchesField, val interface{}, opts ...func(*PayoutBatchesUpdateFieldOption)) PayoutBatchesUpdateField {
	defaultOpt := defaultPayoutBatchesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PayoutBatchesUpdateField{
		payoutBatchesField: field,
		value:              val,
		opt:                defaultOpt,
	}
}
func defaultPayoutBatchesUpdateFields(payoutBatches model.PayoutBatches) (payoutBatchesUpdateFieldList PayoutBatchesUpdateFieldList) {
	selectFields := NewPayoutBatchesSelectFields()
	payoutBatchesUpdateFieldList = append(payoutBatchesUpdateFieldList,
		NewPayoutBatchesUpdateField(selectFields.Id(), payoutBatches.Id),
		NewPayoutBatchesUpdateField(selectFields.PayoutBatchCode(), payoutBatches.PayoutBatchCode),
		NewPayoutBatchesUpdateField(selectFields.ProviderAccountId(), payoutBatches.ProviderAccountId),
		NewPayoutBatchesUpdateField(selectFields.ScheduledFor(), payoutBatches.ScheduledFor),
		NewPayoutBatchesUpdateField(selectFields.CurrencyCode(), payoutBatches.CurrencyCode),
		NewPayoutBatchesUpdateField(selectFields.BatchStatus(), payoutBatches.BatchStatus),
		NewPayoutBatchesUpdateField(selectFields.TotalCount(), payoutBatches.TotalCount),
		NewPayoutBatchesUpdateField(selectFields.TotalAmount(), payoutBatches.TotalAmount),
		NewPayoutBatchesUpdateField(selectFields.Metadata(), payoutBatches.Metadata),
		NewPayoutBatchesUpdateField(selectFields.MetaCreatedAt(), payoutBatches.MetaCreatedAt),
		NewPayoutBatchesUpdateField(selectFields.MetaCreatedBy(), payoutBatches.MetaCreatedBy),
		NewPayoutBatchesUpdateField(selectFields.MetaUpdatedAt(), payoutBatches.MetaUpdatedAt),
		NewPayoutBatchesUpdateField(selectFields.MetaUpdatedBy(), payoutBatches.MetaUpdatedBy),
		NewPayoutBatchesUpdateField(selectFields.MetaDeletedAt(), payoutBatches.MetaDeletedAt),
		NewPayoutBatchesUpdateField(selectFields.MetaDeletedBy(), payoutBatches.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPayoutBatchesCommand(payoutBatchesUpdateFieldList PayoutBatchesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range payoutBatchesUpdateFieldList {
		field := string(updateField.payoutBatchesField)
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

func (repo *RepositoryImpl) BulkCreatePayoutBatches(ctx context.Context, payoutBatchesList []*model.PayoutBatches, fieldsInsert ...PayoutBatchesField) (err error) {
	var (
		fieldsStr              string
		valueListStr           []string
		argsList               []interface{}
		primaryIds             []model.PayoutBatchesPrimaryID
		payoutBatchesValueList []model.PayoutBatches
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPayoutBatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, payoutBatches := range payoutBatchesList {

		primaryIds = append(primaryIds, payoutBatches.ToPayoutBatchesPrimaryID())

		payoutBatchesValueList = append(payoutBatchesValueList, *payoutBatches)
	}

	_, notFoundIds, err := repo.IsExistPayoutBatchesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutBatches] failed checking payoutBatches whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PayoutBatchesPrimaryID{}
		mapNotFoundIds := map[model.PayoutBatchesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "payoutBatches", fmt.Sprintf("payoutBatches with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPayoutBatches(payoutBatchesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(payoutBatchesQueries.insertPayoutBatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutBatches] failed exec create payoutBatches query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePayoutBatchesByIDs(ctx context.Context, primaryIDs []model.PayoutBatchesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPayoutBatchesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutBatchesByIDs] failed checking payoutBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutBatches with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_batches\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(payoutBatchesQueries.deletePayoutBatches + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutBatchesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutBatchesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPayoutBatchesByIDs(ctx context.Context, ids []model.PayoutBatchesPrimaryID) (exists bool, notFoundIds []model.PayoutBatchesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_batches\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(payoutBatchesQueries.selectPayoutBatches, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutBatchesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PayoutBatchesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutBatchesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PayoutBatchesPrimaryID]bool{}
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

// BulkUpdatePayoutBatches is used to bulk update payoutBatches, by default it will update all field
// if want to update specific field, then fill payoutBatchessMapUpdateFieldsRequest else please fill payoutBatchessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePayoutBatches(ctx context.Context, payoutBatchessMap map[model.PayoutBatchesPrimaryID]*model.PayoutBatches, payoutBatchessMapUpdateFieldsRequest map[model.PayoutBatchesPrimaryID]PayoutBatchesUpdateFieldList) (err error) {
	if len(payoutBatchessMap) == 0 && len(payoutBatchessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		payoutBatchessMapUpdateField map[model.PayoutBatchesPrimaryID]PayoutBatchesUpdateFieldList = map[model.PayoutBatchesPrimaryID]PayoutBatchesUpdateFieldList{}
		asTableValues                string                                                        = "myvalues"
	)

	if len(payoutBatchessMap) > 0 {
		for id, payoutBatches := range payoutBatchessMap {
			if payoutBatches == nil {
				log.Error().Err(err).Msg("[BulkUpdatePayoutBatches] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			payoutBatchessMapUpdateField[id] = defaultPayoutBatchesUpdateFields(*payoutBatches)
		}
	} else {
		payoutBatchessMapUpdateField = payoutBatchessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePayoutBatchesQuery(payoutBatchessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPayoutBatchesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutBatches] failed checking payoutBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutBatches with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePayoutBatchesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payout_batches\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutBatches] failed exec query")
	}
	return
}

type PayoutBatchesFieldParameter struct {
	param string
	args  []interface{}
}

func NewPayoutBatchesFieldParameter(param string, args ...interface{}) PayoutBatchesFieldParameter {
	return PayoutBatchesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePayoutBatchesQuery(mapPayoutBatchess map[model.PayoutBatchesPrimaryID]PayoutBatchesUpdateFieldList, asTableValues string) (primaryIDs []model.PayoutBatchesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PayoutBatchesPrimaryID]map[string]interface{}{}
	payoutBatchesSelectFields := NewPayoutBatchesSelectFields()
	for id, updateFields := range mapPayoutBatchess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.payoutBatchesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPayoutBatchess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPayoutBatchesFieldType(updateField.payoutBatchesField)))
			args = append(args, fields[string(updateField.payoutBatchesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.payoutBatchesField))
		if updateField.payoutBatchesField == payoutBatchesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.payoutBatchesField, asTableValues, updateField.payoutBatchesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.payoutBatchesField,
				"\"payout_batches\"", updateField.payoutBatchesField,
				asTableValues, updateField.payoutBatchesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePayoutBatchesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PayoutBatchesPrimaryID, asTableValue string) (whereQry string) {
	payoutBatchesSelectFields := NewPayoutBatchesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payout_batches\".\"id\" = %s.\"id\"::"+GetPayoutBatchesFieldType(payoutBatchesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPayoutBatchesFieldType(payoutBatchesField PayoutBatchesField) string {
	selectPayoutBatchesFields := NewPayoutBatchesSelectFields()
	switch payoutBatchesField {

	case selectPayoutBatchesFields.Id():
		return "uuid"

	case selectPayoutBatchesFields.PayoutBatchCode():
		return "text"

	case selectPayoutBatchesFields.ProviderAccountId():
		return "uuid"

	case selectPayoutBatchesFields.ScheduledFor():
		return "timestamptz"

	case selectPayoutBatchesFields.CurrencyCode():
		return "text"

	case selectPayoutBatchesFields.BatchStatus():
		return "payout_batches_batch_status_enum"

	case selectPayoutBatchesFields.TotalCount():
		return "int4"

	case selectPayoutBatchesFields.TotalAmount():
		return "numeric"

	case selectPayoutBatchesFields.Metadata():
		return "jsonb"

	case selectPayoutBatchesFields.MetaCreatedAt():
		return "timestamptz"

	case selectPayoutBatchesFields.MetaCreatedBy():
		return "uuid"

	case selectPayoutBatchesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPayoutBatchesFields.MetaUpdatedBy():
		return "uuid"

	case selectPayoutBatchesFields.MetaDeletedAt():
		return "timestamptz"

	case selectPayoutBatchesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePayoutBatches(ctx context.Context, payoutBatches *model.PayoutBatches, fieldsInsert ...PayoutBatchesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPayoutBatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PayoutBatchesPrimaryID{
		Id: payoutBatches.Id,
	}
	exists, err := repo.IsExistPayoutBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutBatches] failed checking payoutBatches whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "payoutBatches", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPayoutBatches([]model.PayoutBatches{*payoutBatches}, fieldsInsert...)
	commandQuery := fmt.Sprintf(payoutBatchesQueries.insertPayoutBatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutBatches] failed exec create payoutBatches query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePayoutBatchesByID(ctx context.Context, primaryID model.PayoutBatchesPrimaryID) (err error) {
	exists, err := repo.IsExistPayoutBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutBatchesByID] failed checking payoutBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutBatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePayoutBatchesCompositePrimaryKeyWhere([]model.PayoutBatchesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(payoutBatchesQueries.deletePayoutBatches + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutBatchesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutBatchesByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutBatchesFilterResult, err error) {
	query, args, err := composePayoutBatchesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutBatchesByFilter] failed compose payoutBatches filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutBatchesByFilter] failed get payoutBatches by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePayoutBatchesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PayoutBatchesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePayoutBatchesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePayoutBatchesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePayoutBatchesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPayoutBatchesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 15+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_batch_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_batch_code\"")
			selectedColumns["payout_batch_code"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["scheduled_for"]; !selected {
			selectColumns = append(selectColumns, "base.\"scheduled_for\"")
			selectedColumns["scheduled_for"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["batch_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"batch_status\"")
			selectedColumns["batch_status"] = struct{}{}
		}
		if _, selected := selectedColumns["total_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"total_count\"")
			selectedColumns["total_count"] = struct{}{}
		}
		if _, selected := selectedColumns["total_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"total_amount\"")
			selectedColumns["total_amount"] = struct{}{}
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

type payoutBatchesFilterPlaceholder struct {
	index int
}

func (p *payoutBatchesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePayoutBatchesFilterPredicate(filterField model.FilterField, placeholders *payoutBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPayoutBatchesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePayoutBatchesFilterSQLExpr(spec)
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

func composePayoutBatchesFilterGroup(group model.FilterGroup, placeholders *payoutBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePayoutBatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePayoutBatchesFilterGroup(child, placeholders, args, requiredJoins)
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

func composePayoutBatchesFilterWhereQueries(filter model.Filter, placeholders *payoutBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePayoutBatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePayoutBatchesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePayoutBatchesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePayoutBatchesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePayoutBatchesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePayoutBatchesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := payoutBatchesFilterPlaceholder{index: 1}
	whereQueries, err := composePayoutBatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPayoutBatchesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePayoutBatchesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePayoutBatchesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payout_batches\" base%s", strings.Join(selectColumns, ","), composePayoutBatchesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPayoutBatchesByID(ctx context.Context, primaryID model.PayoutBatchesPrimaryID) (exists bool, err error) {
	whereQuery, params := composePayoutBatchesCompositePrimaryKeyWhere([]model.PayoutBatchesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", payoutBatchesQueries.selectCountPayoutBatches, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutBatchesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutBatches(ctx context.Context, selectFields ...PayoutBatchesField) (payoutBatchesList model.PayoutBatchesList, err error) {
	var (
		defaultPayoutBatchesSelectFields = defaultPayoutBatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutBatchesSelectFields = composePayoutBatchesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(payoutBatchesQueries.selectPayoutBatches, defaultPayoutBatchesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &payoutBatchesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutBatches] failed get payoutBatches list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutBatchesByID(ctx context.Context, primaryID model.PayoutBatchesPrimaryID, selectFields ...PayoutBatchesField) (payoutBatches model.PayoutBatches, err error) {
	var (
		defaultPayoutBatchesSelectFields = defaultPayoutBatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutBatchesSelectFields = composePayoutBatchesSelectFields(selectFields...)
	}
	whereQry, params := composePayoutBatchesCompositePrimaryKeyWhere([]model.PayoutBatchesPrimaryID{primaryID})
	query := fmt.Sprintf(payoutBatchesQueries.selectPayoutBatches+" WHERE "+whereQry, defaultPayoutBatchesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &payoutBatches, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("payoutBatches with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePayoutBatchesByID] failed get payoutBatches")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePayoutBatchesByID(ctx context.Context, primaryID model.PayoutBatchesPrimaryID, payoutBatches *model.PayoutBatches, payoutBatchesUpdateFields ...PayoutBatchesUpdateField) (err error) {
	exists, err := repo.IsExistPayoutBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutBatches] failed checking payoutBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutBatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if payoutBatches == nil {
		if len(payoutBatchesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePayoutBatchesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		payoutBatches = &model.PayoutBatches{}
	}
	var (
		defaultPayoutBatchesUpdateFields = defaultPayoutBatchesUpdateFields(*payoutBatches)
		tempUpdateField                  PayoutBatchesUpdateFieldList
		selectFields                     = NewPayoutBatchesSelectFields()
	)
	if len(payoutBatchesUpdateFields) > 0 {
		for _, updateField := range payoutBatchesUpdateFields {
			if updateField.payoutBatchesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPayoutBatchesUpdateFields = tempUpdateField
	}
	whereQuery, params := composePayoutBatchesCompositePrimaryKeyWhere([]model.PayoutBatchesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPayoutBatchesCommand(defaultPayoutBatchesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(payoutBatchesQueries.updatePayoutBatches+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutBatches] error when try to update payoutBatches by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePayoutBatchesByFilter(ctx context.Context, filter model.Filter, payoutBatchesUpdateFields ...PayoutBatchesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(payoutBatchesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PayoutBatchesUpdateFieldList
		selectFields = NewPayoutBatchesSelectFields()
	)
	for _, updateField := range payoutBatchesUpdateFields {
		if updateField.payoutBatchesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPayoutBatchesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := payoutBatchesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePayoutBatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payout_batches\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutBatchesByFilter] error when try to update payoutBatches by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutBatchesByFilter] failed get rows affected")
	}
	return
}

var (
	payoutBatchesQueries = struct {
		selectPayoutBatches      string
		selectCountPayoutBatches string
		deletePayoutBatches      string
		updatePayoutBatches      string
		insertPayoutBatches      string
	}{
		selectPayoutBatches:      "SELECT %s FROM \"payout_batches\"",
		selectCountPayoutBatches: "SELECT COUNT(\"id\") FROM \"payout_batches\"",
		deletePayoutBatches:      "DELETE FROM \"payout_batches\"",
		updatePayoutBatches:      "UPDATE \"payout_batches\" SET %s ",
		insertPayoutBatches:      "INSERT INTO \"payout_batches\" %s VALUES %s",
	}
)

type PayoutBatchesRepository interface {
	CreatePayoutBatches(ctx context.Context, payoutBatches *model.PayoutBatches, fieldsInsert ...PayoutBatchesField) error
	BulkCreatePayoutBatches(ctx context.Context, payoutBatchesList []*model.PayoutBatches, fieldsInsert ...PayoutBatchesField) error
	ResolvePayoutBatches(ctx context.Context, selectFields ...PayoutBatchesField) (model.PayoutBatchesList, error)
	ResolvePayoutBatchesByID(ctx context.Context, primaryID model.PayoutBatchesPrimaryID, selectFields ...PayoutBatchesField) (model.PayoutBatches, error)
	UpdatePayoutBatchesByID(ctx context.Context, id model.PayoutBatchesPrimaryID, payoutBatches *model.PayoutBatches, payoutBatchesUpdateFields ...PayoutBatchesUpdateField) error
	UpdatePayoutBatchesByFilter(ctx context.Context, filter model.Filter, payoutBatchesUpdateFields ...PayoutBatchesUpdateField) (rowsAffected int64, err error)
	BulkUpdatePayoutBatches(ctx context.Context, payoutBatchesListMap map[model.PayoutBatchesPrimaryID]*model.PayoutBatches, PayoutBatchessMapUpdateFieldsRequest map[model.PayoutBatchesPrimaryID]PayoutBatchesUpdateFieldList) (err error)
	DeletePayoutBatchesByID(ctx context.Context, id model.PayoutBatchesPrimaryID) error
	BulkDeletePayoutBatchesByIDs(ctx context.Context, ids []model.PayoutBatchesPrimaryID) error
	ResolvePayoutBatchesByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutBatchesFilterResult, err error)
	IsExistPayoutBatchesByIDs(ctx context.Context, ids []model.PayoutBatchesPrimaryID) (exists bool, notFoundIds []model.PayoutBatchesPrimaryID, err error)
	IsExistPayoutBatchesByID(ctx context.Context, id model.PayoutBatchesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
