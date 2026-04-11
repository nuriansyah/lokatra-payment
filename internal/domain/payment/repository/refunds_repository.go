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

func composeInsertFieldsAndParamsRefunds(refundsList []model.Refunds, fieldsInsert ...RefundsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRefundsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, refunds := range refundsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, refunds.Id)
			case selectField.RefundCode():
				args = append(args, refunds.RefundCode)
			case selectField.PaymentId():
				args = append(args, refunds.PaymentId)
			case selectField.IntentId():
				args = append(args, refunds.IntentId)
			case selectField.Amount():
				args = append(args, refunds.Amount)
			case selectField.Currency():
				args = append(args, refunds.Currency)
			case selectField.Reason():
				args = append(args, refunds.Reason)
			case selectField.ReasonDetail():
				args = append(args, refunds.ReasonDetail)
			case selectField.Status():
				args = append(args, refunds.Status)
			case selectField.PspRefundId():
				args = append(args, refunds.PspRefundId)
			case selectField.PspRawResponse():
				args = append(args, refunds.PspRawResponse)
			case selectField.RequestedBy():
				args = append(args, refunds.RequestedBy)
			case selectField.ReviewedBy():
				args = append(args, refunds.ReviewedBy)
			case selectField.ReviewedAt():
				args = append(args, refunds.ReviewedAt)
			case selectField.ReviewNotes():
				args = append(args, refunds.ReviewNotes)
			case selectField.RefundedAt():
				args = append(args, refunds.RefundedAt)
			case selectField.EstimatedArrival():
				args = append(args, refunds.EstimatedArrival)
			case selectField.FailureReason():
				args = append(args, refunds.FailureReason)
			case selectField.IdempotencyKeyId():
				args = append(args, refunds.IdempotencyKeyId)
			case selectField.Metadata():
				args = append(args, refunds.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, refunds.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, refunds.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, refunds.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, refunds.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, refunds.MetaDeletedAt)

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

func composeRefundsCompositePrimaryKeyWhere(primaryIDs []model.RefundsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"refunds\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRefundsSelectFields() string {
	fields := NewRefundsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRefundsSelectFields(selectFields ...RefundsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RefundsField string
type RefundsFieldList []RefundsField

type RefundsSelectFields struct {
}

func (ss RefundsSelectFields) Id() RefundsField {
	return RefundsField("id")
}

func (ss RefundsSelectFields) RefundCode() RefundsField {
	return RefundsField("refund_code")
}

func (ss RefundsSelectFields) PaymentId() RefundsField {
	return RefundsField("payment_id")
}

func (ss RefundsSelectFields) IntentId() RefundsField {
	return RefundsField("intent_id")
}

func (ss RefundsSelectFields) Amount() RefundsField {
	return RefundsField("amount")
}

func (ss RefundsSelectFields) Currency() RefundsField {
	return RefundsField("currency")
}

func (ss RefundsSelectFields) Reason() RefundsField {
	return RefundsField("reason")
}

func (ss RefundsSelectFields) ReasonDetail() RefundsField {
	return RefundsField("reason_detail")
}

func (ss RefundsSelectFields) Status() RefundsField {
	return RefundsField("status")
}

func (ss RefundsSelectFields) PspRefundId() RefundsField {
	return RefundsField("psp_refund_id")
}

func (ss RefundsSelectFields) PspRawResponse() RefundsField {
	return RefundsField("psp_raw_response")
}

func (ss RefundsSelectFields) RequestedBy() RefundsField {
	return RefundsField("requested_by")
}

func (ss RefundsSelectFields) ReviewedBy() RefundsField {
	return RefundsField("reviewed_by")
}

func (ss RefundsSelectFields) ReviewedAt() RefundsField {
	return RefundsField("reviewed_at")
}

func (ss RefundsSelectFields) ReviewNotes() RefundsField {
	return RefundsField("review_notes")
}

func (ss RefundsSelectFields) RefundedAt() RefundsField {
	return RefundsField("refunded_at")
}

func (ss RefundsSelectFields) EstimatedArrival() RefundsField {
	return RefundsField("estimated_arrival")
}

func (ss RefundsSelectFields) FailureReason() RefundsField {
	return RefundsField("failure_reason")
}

func (ss RefundsSelectFields) IdempotencyKeyId() RefundsField {
	return RefundsField("idempotency_key_id")
}

func (ss RefundsSelectFields) Metadata() RefundsField {
	return RefundsField("metadata")
}

func (ss RefundsSelectFields) MetaCreatedAt() RefundsField {
	return RefundsField("meta_created_at")
}

func (ss RefundsSelectFields) MetaCreatedBy() RefundsField {
	return RefundsField("meta_created_by")
}

func (ss RefundsSelectFields) MetaUpdatedAt() RefundsField {
	return RefundsField("meta_updated_at")
}

func (ss RefundsSelectFields) MetaUpdatedBy() RefundsField {
	return RefundsField("meta_updated_by")
}

func (ss RefundsSelectFields) MetaDeletedAt() RefundsField {
	return RefundsField("meta_deleted_at")
}

func (ss RefundsSelectFields) All() RefundsFieldList {
	return []RefundsField{
		ss.Id(),
		ss.RefundCode(),
		ss.PaymentId(),
		ss.IntentId(),
		ss.Amount(),
		ss.Currency(),
		ss.Reason(),
		ss.ReasonDetail(),
		ss.Status(),
		ss.PspRefundId(),
		ss.PspRawResponse(),
		ss.RequestedBy(),
		ss.ReviewedBy(),
		ss.ReviewedAt(),
		ss.ReviewNotes(),
		ss.RefundedAt(),
		ss.EstimatedArrival(),
		ss.FailureReason(),
		ss.IdempotencyKeyId(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
	}
}

func NewRefundsSelectFields() RefundsSelectFields {
	return RefundsSelectFields{}
}

type RefundsUpdateFieldOption struct {
	useIncrement bool
}
type RefundsUpdateField struct {
	refundsField RefundsField
	opt          RefundsUpdateFieldOption
	value        interface{}
}
type RefundsUpdateFieldList []RefundsUpdateField

func defaultRefundsUpdateFieldOption() RefundsUpdateFieldOption {
	return RefundsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRefundsOption(useIncrement bool) func(*RefundsUpdateFieldOption) {
	return func(pcufo *RefundsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRefundsUpdateField(field RefundsField, val interface{}, opts ...func(*RefundsUpdateFieldOption)) RefundsUpdateField {
	defaultOpt := defaultRefundsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RefundsUpdateField{
		refundsField: field,
		value:        val,
		opt:          defaultOpt,
	}
}
func defaultRefundsUpdateFields(refunds model.Refunds) (refundsUpdateFieldList RefundsUpdateFieldList) {
	selectFields := NewRefundsSelectFields()
	refundsUpdateFieldList = append(refundsUpdateFieldList,
		NewRefundsUpdateField(selectFields.Id(), refunds.Id),
		NewRefundsUpdateField(selectFields.RefundCode(), refunds.RefundCode),
		NewRefundsUpdateField(selectFields.PaymentId(), refunds.PaymentId),
		NewRefundsUpdateField(selectFields.IntentId(), refunds.IntentId),
		NewRefundsUpdateField(selectFields.Amount(), refunds.Amount),
		NewRefundsUpdateField(selectFields.Currency(), refunds.Currency),
		NewRefundsUpdateField(selectFields.Reason(), refunds.Reason),
		NewRefundsUpdateField(selectFields.ReasonDetail(), refunds.ReasonDetail),
		NewRefundsUpdateField(selectFields.Status(), refunds.Status),
		NewRefundsUpdateField(selectFields.PspRefundId(), refunds.PspRefundId),
		NewRefundsUpdateField(selectFields.PspRawResponse(), refunds.PspRawResponse),
		NewRefundsUpdateField(selectFields.RequestedBy(), refunds.RequestedBy),
		NewRefundsUpdateField(selectFields.ReviewedBy(), refunds.ReviewedBy),
		NewRefundsUpdateField(selectFields.ReviewedAt(), refunds.ReviewedAt),
		NewRefundsUpdateField(selectFields.ReviewNotes(), refunds.ReviewNotes),
		NewRefundsUpdateField(selectFields.RefundedAt(), refunds.RefundedAt),
		NewRefundsUpdateField(selectFields.EstimatedArrival(), refunds.EstimatedArrival),
		NewRefundsUpdateField(selectFields.FailureReason(), refunds.FailureReason),
		NewRefundsUpdateField(selectFields.IdempotencyKeyId(), refunds.IdempotencyKeyId),
		NewRefundsUpdateField(selectFields.Metadata(), refunds.Metadata),
		NewRefundsUpdateField(selectFields.MetaCreatedAt(), refunds.MetaCreatedAt),
		NewRefundsUpdateField(selectFields.MetaCreatedBy(), refunds.MetaCreatedBy),
		NewRefundsUpdateField(selectFields.MetaUpdatedAt(), refunds.MetaUpdatedAt),
		NewRefundsUpdateField(selectFields.MetaUpdatedBy(), refunds.MetaUpdatedBy),
		NewRefundsUpdateField(selectFields.MetaDeletedAt(), refunds.MetaDeletedAt),
	)
	return
}
func composeUpdateFieldsRefundsCommand(refundsUpdateFieldList RefundsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range refundsUpdateFieldList {
		field := string(updateField.refundsField)
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

func (repo *RepositoryImpl) BulkCreateRefunds(ctx context.Context, refundsList []*model.Refunds, fieldsInsert ...RefundsField) (err error) {
	var (
		fieldsStr        string
		valueListStr     []string
		argsList         []interface{}
		primaryIds       []model.RefundsPrimaryID
		refundsValueList []model.Refunds
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRefundsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, refunds := range refundsList {

		primaryIds = append(primaryIds, refunds.ToRefundsPrimaryID())

		refundsValueList = append(refundsValueList, *refunds)
	}

	_, notFoundIds, err := repo.IsExistRefundsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefunds] failed checking refunds whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RefundsPrimaryID{}
		mapNotFoundIds := map[model.RefundsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "refunds", fmt.Sprintf("refunds with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRefunds(refundsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(refundsQueries.insertRefunds, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefunds] failed exec create refunds query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRefundsByIDs(ctx context.Context, primaryIDs []model.RefundsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRefundsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundsByIDs] failed checking refunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refunds with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refunds\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(refundsQueries.deleteRefunds + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRefundsByIDs(ctx context.Context, ids []model.RefundsPrimaryID) (exists bool, notFoundIds []model.RefundsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refunds\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(refundsQueries.selectRefunds, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RefundsPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RefundsPrimaryID]bool{}
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

// BulkUpdateRefunds is used to bulk update refunds, by default it will update all field
// if want to update specific field, then fill refundssMapUpdateFieldsRequest else please fill refundssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRefunds(ctx context.Context, refundssMap map[model.RefundsPrimaryID]*model.Refunds, refundssMapUpdateFieldsRequest map[model.RefundsPrimaryID]RefundsUpdateFieldList) (err error) {
	if len(refundssMap) == 0 && len(refundssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		refundssMapUpdateField map[model.RefundsPrimaryID]RefundsUpdateFieldList = map[model.RefundsPrimaryID]RefundsUpdateFieldList{}
		asTableValues          string                                            = "myvalues"
	)

	if len(refundssMap) > 0 {
		for id, refunds := range refundssMap {
			if refunds == nil {
				log.Error().Err(err).Msg("[BulkUpdateRefunds] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			refundssMapUpdateField[id] = defaultRefundsUpdateFields(*refunds)
		}
	} else {
		refundssMapUpdateField = refundssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRefundsQuery(refundssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRefundsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefunds] failed checking refunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refunds with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRefundsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"refunds\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefunds] failed exec query")
	}
	return
}

type RefundsFieldParameter struct {
	param string
	args  []interface{}
}

func NewRefundsFieldParameter(param string, args ...interface{}) RefundsFieldParameter {
	return RefundsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRefundsQuery(mapRefundss map[model.RefundsPrimaryID]RefundsUpdateFieldList, asTableValues string) (primaryIDs []model.RefundsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RefundsPrimaryID]map[string]interface{}{}
	refundsSelectFields := NewRefundsSelectFields()
	for id, updateFields := range mapRefundss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.refundsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRefundss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRefundsFieldType(updateField.refundsField)))
			args = append(args, fields[string(updateField.refundsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.refundsField))
		if updateField.refundsField == refundsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.refundsField, asTableValues, updateField.refundsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.refundsField,
				"\"refunds\"", updateField.refundsField,
				asTableValues, updateField.refundsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRefundsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RefundsPrimaryID, asTableValue string) (whereQry string) {
	refundsSelectFields := NewRefundsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"refunds\".\"id\" = %s.\"id\"::"+GetRefundsFieldType(refundsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRefundsFieldType(refundsField RefundsField) string {
	selectRefundsFields := NewRefundsSelectFields()
	switch refundsField {

	case selectRefundsFields.Id():
		return "uuid"

	case selectRefundsFields.RefundCode():
		return "text"

	case selectRefundsFields.PaymentId():
		return "uuid"

	case selectRefundsFields.IntentId():
		return "uuid"

	case selectRefundsFields.Amount():
		return "numeric"

	case selectRefundsFields.Currency():
		return "payment_currency"

	case selectRefundsFields.Reason():
		return "text"

	case selectRefundsFields.ReasonDetail():
		return "text"

	case selectRefundsFields.Status():
		return "refund_status_enum"

	case selectRefundsFields.PspRefundId():
		return "text"

	case selectRefundsFields.PspRawResponse():
		return "jsonb"

	case selectRefundsFields.RequestedBy():
		return "uuid"

	case selectRefundsFields.ReviewedBy():
		return "uuid"

	case selectRefundsFields.ReviewedAt():
		return "timestamptz"

	case selectRefundsFields.ReviewNotes():
		return "text"

	case selectRefundsFields.RefundedAt():
		return "timestamptz"

	case selectRefundsFields.EstimatedArrival():
		return "timestamptz"

	case selectRefundsFields.FailureReason():
		return "text"

	case selectRefundsFields.IdempotencyKeyId():
		return "uuid"

	case selectRefundsFields.Metadata():
		return "jsonb"

	case selectRefundsFields.MetaCreatedAt():
		return "timestamptz"

	case selectRefundsFields.MetaCreatedBy():
		return "uuid"

	case selectRefundsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRefundsFields.MetaUpdatedBy():
		return "uuid"

	case selectRefundsFields.MetaDeletedAt():
		return "timestamptz"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRefunds(ctx context.Context, refunds *model.Refunds, fieldsInsert ...RefundsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRefundsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RefundsPrimaryID{
		Id: refunds.Id,
	}
	exists, err := repo.IsExistRefundsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefunds] failed checking refunds whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "refunds", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRefunds([]model.Refunds{*refunds}, fieldsInsert...)
	commandQuery := fmt.Sprintf(refundsQueries.insertRefunds, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefunds] failed exec create refunds query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRefundsByID(ctx context.Context, primaryID model.RefundsPrimaryID) (err error) {
	exists, err := repo.IsExistRefundsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundsByID] failed checking refunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refunds with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRefundsCompositePrimaryKeyWhere([]model.RefundsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(refundsQueries.deleteRefunds + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundsFilterResult, err error) {
	query, args, err := composeRefundsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundsByFilter] failed compose refunds filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundsByFilter] failed get refunds by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRefundsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRefundsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultRefundsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := RefundsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, RefundsField(filterSelectField))
		}
		selectFields = composeRefundsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(refundsQueries.selectRefunds, selectFields)

	if len(filter.FilterFields) > 0 {
		var (
			whereQueries []string
			whereArgs    []interface{}
		)
		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorEqual:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" = $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			case model.OperatorRange:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" BETWEEN $%d AND $%d", filterField.Field, index, index+1))
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				whereArgs = append(whereArgs, valueArray...)
				index += 2
			case model.OperatorIn:
				valueArray, ok := filterField.Value.([]interface{})
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				var placeholder []string
				for range valueArray {
					placeholder = append(placeholder, fmt.Sprintf("$%d", index))
					index++
				}
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IN (%s)", filterField.Field, strings.Join(placeholder, ",")))
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIsNull:
				value, ok := filterField.Value.(bool)
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				if value {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NULL", filterField.Field))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NOT NULL", filterField.Field))
				}
			case model.OperatorNot:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" != $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			}
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND "))
		args = append(args, whereArgs...)
	}

	sortQuery := []string{}
	for _, sort := range filter.Sorts {
		sortQuery = append(sortQuery, fmt.Sprintf("\"%s\" %s", sort.Field, sort.Order))
	}
	if len(sortQuery) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
	}
	if filter.Pagination.PageSize > 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Pagination.PageSize)
		if filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
	}

	return
}

func (repo *RepositoryImpl) IsExistRefundsByID(ctx context.Context, primaryID model.RefundsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRefundsCompositePrimaryKeyWhere([]model.RefundsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", refundsQueries.selectCountRefunds, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefunds(ctx context.Context, selectFields ...RefundsField) (refundsList model.RefundsList, err error) {
	var (
		defaultRefundsSelectFields = defaultRefundsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundsSelectFields = composeRefundsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(refundsQueries.selectRefunds, defaultRefundsSelectFields)

	err = repo.db.Read.Select(&refundsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefunds] failed get refunds list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundsByID(ctx context.Context, primaryID model.RefundsPrimaryID, selectFields ...RefundsField) (refunds model.Refunds, err error) {
	var (
		defaultRefundsSelectFields = defaultRefundsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundsSelectFields = composeRefundsSelectFields(selectFields...)
	}
	whereQry, params := composeRefundsCompositePrimaryKeyWhere([]model.RefundsPrimaryID{primaryID})
	query := fmt.Sprintf(refundsQueries.selectRefunds+" WHERE "+whereQry, defaultRefundsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&refunds, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("refunds with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRefundsByID] failed get refunds")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRefundsByID(ctx context.Context, primaryID model.RefundsPrimaryID, refunds *model.Refunds, refundsUpdateFields ...RefundsUpdateField) (err error) {
	exists, err := repo.IsExistRefundsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefunds] failed checking refunds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refunds with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if refunds == nil {
		if len(refundsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRefundsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		refunds = &model.Refunds{}
	}
	var (
		defaultRefundsUpdateFields = defaultRefundsUpdateFields(*refunds)
		tempUpdateField            RefundsUpdateFieldList
		selectFields               = NewRefundsSelectFields()
	)
	if len(refundsUpdateFields) > 0 {
		for _, updateField := range refundsUpdateFields {
			if updateField.refundsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRefundsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRefundsCompositePrimaryKeyWhere([]model.RefundsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRefundsCommand(defaultRefundsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(refundsQueries.updateRefunds+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefunds] error when try to update refunds by id")
	}
	return err
}

var (
	refundsQueries = struct {
		selectRefunds      string
		selectCountRefunds string
		deleteRefunds      string
		updateRefunds      string
		insertRefunds      string
	}{
		selectRefunds:      "SELECT %s FROM \"refunds\"",
		selectCountRefunds: "SELECT COUNT(\"id\") FROM \"refunds\"",
		deleteRefunds:      "DELETE FROM \"refunds\"",
		updateRefunds:      "UPDATE \"refunds\" SET %s ",
		insertRefunds:      "INSERT INTO \"refunds\" %s VALUES %s",
	}
)

type RefundsRepository interface {
	CreateRefunds(ctx context.Context, refunds *model.Refunds, fieldsInsert ...RefundsField) error
	BulkCreateRefunds(ctx context.Context, refundsList []*model.Refunds, fieldsInsert ...RefundsField) error
	ResolveRefunds(ctx context.Context, selectFields ...RefundsField) (model.RefundsList, error)
	ResolveRefundsByID(ctx context.Context, primaryID model.RefundsPrimaryID, selectFields ...RefundsField) (model.Refunds, error)
	UpdateRefundsByID(ctx context.Context, id model.RefundsPrimaryID, refunds *model.Refunds, refundsUpdateFields ...RefundsUpdateField) error
	BulkUpdateRefunds(ctx context.Context, refundsListMap map[model.RefundsPrimaryID]*model.Refunds, RefundssMapUpdateFieldsRequest map[model.RefundsPrimaryID]RefundsUpdateFieldList) (err error)
	DeleteRefundsByID(ctx context.Context, id model.RefundsPrimaryID) error
	BulkDeleteRefundsByIDs(ctx context.Context, ids []model.RefundsPrimaryID) error
	ResolveRefundsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundsFilterResult, err error)
	IsExistRefundsByIDs(ctx context.Context, ids []model.RefundsPrimaryID) (exists bool, notFoundIds []model.RefundsPrimaryID, err error)
	IsExistRefundsByID(ctx context.Context, id model.RefundsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
