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

func composeInsertFieldsAndParamsPayoutApprovals(payoutApprovalsList []model.PayoutApprovals, fieldsInsert ...PayoutApprovalsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPayoutApprovalsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, payoutApprovals := range payoutApprovalsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, payoutApprovals.Id)
			case selectField.PayoutId():
				args = append(args, payoutApprovals.PayoutId)
			case selectField.ApprovalStatus():
				args = append(args, payoutApprovals.ApprovalStatus)
			case selectField.ReasonCode():
				args = append(args, payoutApprovals.ReasonCode)
			case selectField.ReasonDetail():
				args = append(args, payoutApprovals.ReasonDetail)
			case selectField.ApprovedBy():
				args = append(args, payoutApprovals.ApprovedBy)
			case selectField.ApprovedAt():
				args = append(args, payoutApprovals.ApprovedAt)
			case selectField.Metadata():
				args = append(args, payoutApprovals.Metadata)
			case selectField.ApprovedAmountSnapshot():
				args = append(args, payoutApprovals.ApprovedAmountSnapshot)
			case selectField.CurrencyCodeSnapshot():
				args = append(args, payoutApprovals.CurrencyCodeSnapshot)
			case selectField.PayoutRevisionHash():
				args = append(args, payoutApprovals.PayoutRevisionHash)
			case selectField.MetaCreatedAt():
				args = append(args, payoutApprovals.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, payoutApprovals.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, payoutApprovals.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, payoutApprovals.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, payoutApprovals.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, payoutApprovals.MetaDeletedBy)

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

func composePayoutApprovalsCompositePrimaryKeyWhere(primaryIDs []model.PayoutApprovalsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"payout_approvals\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPayoutApprovalsSelectFields() string {
	fields := NewPayoutApprovalsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePayoutApprovalsSelectFields(selectFields ...PayoutApprovalsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PayoutApprovalsField string
type PayoutApprovalsFieldList []PayoutApprovalsField

type PayoutApprovalsSelectFields struct {
}

func (ss PayoutApprovalsSelectFields) Id() PayoutApprovalsField {
	return PayoutApprovalsField("id")
}

func (ss PayoutApprovalsSelectFields) PayoutId() PayoutApprovalsField {
	return PayoutApprovalsField("payout_id")
}

func (ss PayoutApprovalsSelectFields) ApprovalStatus() PayoutApprovalsField {
	return PayoutApprovalsField("approval_status")
}

func (ss PayoutApprovalsSelectFields) ReasonCode() PayoutApprovalsField {
	return PayoutApprovalsField("reason_code")
}

func (ss PayoutApprovalsSelectFields) ReasonDetail() PayoutApprovalsField {
	return PayoutApprovalsField("reason_detail")
}

func (ss PayoutApprovalsSelectFields) ApprovedBy() PayoutApprovalsField {
	return PayoutApprovalsField("approved_by")
}

func (ss PayoutApprovalsSelectFields) ApprovedAt() PayoutApprovalsField {
	return PayoutApprovalsField("approved_at")
}

func (ss PayoutApprovalsSelectFields) Metadata() PayoutApprovalsField {
	return PayoutApprovalsField("metadata")
}

func (ss PayoutApprovalsSelectFields) ApprovedAmountSnapshot() PayoutApprovalsField {
	return PayoutApprovalsField("approved_amount_snapshot")
}

func (ss PayoutApprovalsSelectFields) CurrencyCodeSnapshot() PayoutApprovalsField {
	return PayoutApprovalsField("currency_code_snapshot")
}

func (ss PayoutApprovalsSelectFields) PayoutRevisionHash() PayoutApprovalsField {
	return PayoutApprovalsField("payout_revision_hash")
}

func (ss PayoutApprovalsSelectFields) MetaCreatedAt() PayoutApprovalsField {
	return PayoutApprovalsField("meta_created_at")
}

func (ss PayoutApprovalsSelectFields) MetaCreatedBy() PayoutApprovalsField {
	return PayoutApprovalsField("meta_created_by")
}

func (ss PayoutApprovalsSelectFields) MetaUpdatedAt() PayoutApprovalsField {
	return PayoutApprovalsField("meta_updated_at")
}

func (ss PayoutApprovalsSelectFields) MetaUpdatedBy() PayoutApprovalsField {
	return PayoutApprovalsField("meta_updated_by")
}

func (ss PayoutApprovalsSelectFields) MetaDeletedAt() PayoutApprovalsField {
	return PayoutApprovalsField("meta_deleted_at")
}

func (ss PayoutApprovalsSelectFields) MetaDeletedBy() PayoutApprovalsField {
	return PayoutApprovalsField("meta_deleted_by")
}

func (ss PayoutApprovalsSelectFields) All() PayoutApprovalsFieldList {
	return []PayoutApprovalsField{
		ss.Id(),
		ss.PayoutId(),
		ss.ApprovalStatus(),
		ss.ReasonCode(),
		ss.ReasonDetail(),
		ss.ApprovedBy(),
		ss.ApprovedAt(),
		ss.Metadata(),
		ss.ApprovedAmountSnapshot(),
		ss.CurrencyCodeSnapshot(),
		ss.PayoutRevisionHash(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewPayoutApprovalsSelectFields() PayoutApprovalsSelectFields {
	return PayoutApprovalsSelectFields{}
}

type PayoutApprovalsUpdateFieldOption struct {
	useIncrement bool
}
type PayoutApprovalsUpdateField struct {
	payoutApprovalsField PayoutApprovalsField
	opt                  PayoutApprovalsUpdateFieldOption
	value                interface{}
}
type PayoutApprovalsUpdateFieldList []PayoutApprovalsUpdateField

func defaultPayoutApprovalsUpdateFieldOption() PayoutApprovalsUpdateFieldOption {
	return PayoutApprovalsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPayoutApprovalsOption(useIncrement bool) func(*PayoutApprovalsUpdateFieldOption) {
	return func(pcufo *PayoutApprovalsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPayoutApprovalsUpdateField(field PayoutApprovalsField, val interface{}, opts ...func(*PayoutApprovalsUpdateFieldOption)) PayoutApprovalsUpdateField {
	defaultOpt := defaultPayoutApprovalsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PayoutApprovalsUpdateField{
		payoutApprovalsField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultPayoutApprovalsUpdateFields(payoutApprovals model.PayoutApprovals) (payoutApprovalsUpdateFieldList PayoutApprovalsUpdateFieldList) {
	selectFields := NewPayoutApprovalsSelectFields()
	payoutApprovalsUpdateFieldList = append(payoutApprovalsUpdateFieldList,
		NewPayoutApprovalsUpdateField(selectFields.Id(), payoutApprovals.Id),
		NewPayoutApprovalsUpdateField(selectFields.PayoutId(), payoutApprovals.PayoutId),
		NewPayoutApprovalsUpdateField(selectFields.ApprovalStatus(), payoutApprovals.ApprovalStatus),
		NewPayoutApprovalsUpdateField(selectFields.ReasonCode(), payoutApprovals.ReasonCode),
		NewPayoutApprovalsUpdateField(selectFields.ReasonDetail(), payoutApprovals.ReasonDetail),
		NewPayoutApprovalsUpdateField(selectFields.ApprovedBy(), payoutApprovals.ApprovedBy),
		NewPayoutApprovalsUpdateField(selectFields.ApprovedAt(), payoutApprovals.ApprovedAt),
		NewPayoutApprovalsUpdateField(selectFields.Metadata(), payoutApprovals.Metadata),
		NewPayoutApprovalsUpdateField(selectFields.ApprovedAmountSnapshot(), payoutApprovals.ApprovedAmountSnapshot),
		NewPayoutApprovalsUpdateField(selectFields.CurrencyCodeSnapshot(), payoutApprovals.CurrencyCodeSnapshot),
		NewPayoutApprovalsUpdateField(selectFields.PayoutRevisionHash(), payoutApprovals.PayoutRevisionHash),
		NewPayoutApprovalsUpdateField(selectFields.MetaCreatedAt(), payoutApprovals.MetaCreatedAt),
		NewPayoutApprovalsUpdateField(selectFields.MetaCreatedBy(), payoutApprovals.MetaCreatedBy),
		NewPayoutApprovalsUpdateField(selectFields.MetaUpdatedAt(), payoutApprovals.MetaUpdatedAt),
		NewPayoutApprovalsUpdateField(selectFields.MetaUpdatedBy(), payoutApprovals.MetaUpdatedBy),
		NewPayoutApprovalsUpdateField(selectFields.MetaDeletedAt(), payoutApprovals.MetaDeletedAt),
		NewPayoutApprovalsUpdateField(selectFields.MetaDeletedBy(), payoutApprovals.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsPayoutApprovalsCommand(payoutApprovalsUpdateFieldList PayoutApprovalsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range payoutApprovalsUpdateFieldList {
		field := string(updateField.payoutApprovalsField)
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

func (repo *RepositoryImpl) BulkCreatePayoutApprovals(ctx context.Context, payoutApprovalsList []*model.PayoutApprovals, fieldsInsert ...PayoutApprovalsField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.PayoutApprovalsPrimaryID
		payoutApprovalsValueList []model.PayoutApprovals
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPayoutApprovalsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, payoutApprovals := range payoutApprovalsList {

		primaryIds = append(primaryIds, payoutApprovals.ToPayoutApprovalsPrimaryID())

		payoutApprovalsValueList = append(payoutApprovalsValueList, *payoutApprovals)
	}

	_, notFoundIds, err := repo.IsExistPayoutApprovalsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutApprovals] failed checking payoutApprovals whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PayoutApprovalsPrimaryID{}
		mapNotFoundIds := map[model.PayoutApprovalsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "payoutApprovals", fmt.Sprintf("payoutApprovals with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPayoutApprovals(payoutApprovalsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(payoutApprovalsQueries.insertPayoutApprovals, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePayoutApprovals] failed exec create payoutApprovals query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePayoutApprovalsByIDs(ctx context.Context, primaryIDs []model.PayoutApprovalsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPayoutApprovalsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutApprovalsByIDs] failed checking payoutApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutApprovals with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_approvals\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(payoutApprovalsQueries.deletePayoutApprovals + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutApprovalsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePayoutApprovalsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPayoutApprovalsByIDs(ctx context.Context, ids []model.PayoutApprovalsPrimaryID) (exists bool, notFoundIds []model.PayoutApprovalsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"payout_approvals\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(payoutApprovalsQueries.selectPayoutApprovals, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutApprovalsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PayoutApprovalsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutApprovalsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PayoutApprovalsPrimaryID]bool{}
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

// BulkUpdatePayoutApprovals is used to bulk update payoutApprovals, by default it will update all field
// if want to update specific field, then fill payoutApprovalssMapUpdateFieldsRequest else please fill payoutApprovalssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePayoutApprovals(ctx context.Context, payoutApprovalssMap map[model.PayoutApprovalsPrimaryID]*model.PayoutApprovals, payoutApprovalssMapUpdateFieldsRequest map[model.PayoutApprovalsPrimaryID]PayoutApprovalsUpdateFieldList) (err error) {
	if len(payoutApprovalssMap) == 0 && len(payoutApprovalssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		payoutApprovalssMapUpdateField map[model.PayoutApprovalsPrimaryID]PayoutApprovalsUpdateFieldList = map[model.PayoutApprovalsPrimaryID]PayoutApprovalsUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(payoutApprovalssMap) > 0 {
		for id, payoutApprovals := range payoutApprovalssMap {
			if payoutApprovals == nil {
				log.Error().Err(err).Msg("[BulkUpdatePayoutApprovals] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			payoutApprovalssMapUpdateField[id] = defaultPayoutApprovalsUpdateFields(*payoutApprovals)
		}
	} else {
		payoutApprovalssMapUpdateField = payoutApprovalssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePayoutApprovalsQuery(payoutApprovalssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPayoutApprovalsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutApprovals] failed checking payoutApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutApprovals with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePayoutApprovalsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"payout_approvals\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePayoutApprovals] failed exec query")
	}
	return
}

type PayoutApprovalsFieldParameter struct {
	param string
	args  []interface{}
}

func NewPayoutApprovalsFieldParameter(param string, args ...interface{}) PayoutApprovalsFieldParameter {
	return PayoutApprovalsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePayoutApprovalsQuery(mapPayoutApprovalss map[model.PayoutApprovalsPrimaryID]PayoutApprovalsUpdateFieldList, asTableValues string) (primaryIDs []model.PayoutApprovalsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PayoutApprovalsPrimaryID]map[string]interface{}{}
	payoutApprovalsSelectFields := NewPayoutApprovalsSelectFields()
	for id, updateFields := range mapPayoutApprovalss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.payoutApprovalsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPayoutApprovalss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPayoutApprovalsFieldType(updateField.payoutApprovalsField)))
			args = append(args, fields[string(updateField.payoutApprovalsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.payoutApprovalsField))
		if updateField.payoutApprovalsField == payoutApprovalsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.payoutApprovalsField, asTableValues, updateField.payoutApprovalsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.payoutApprovalsField,
				"\"payout_approvals\"", updateField.payoutApprovalsField,
				asTableValues, updateField.payoutApprovalsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePayoutApprovalsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PayoutApprovalsPrimaryID, asTableValue string) (whereQry string) {
	payoutApprovalsSelectFields := NewPayoutApprovalsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"payout_approvals\".\"id\" = %s.\"id\"::"+GetPayoutApprovalsFieldType(payoutApprovalsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPayoutApprovalsFieldType(payoutApprovalsField PayoutApprovalsField) string {
	selectPayoutApprovalsFields := NewPayoutApprovalsSelectFields()
	switch payoutApprovalsField {

	case selectPayoutApprovalsFields.Id():
		return "uuid"

	case selectPayoutApprovalsFields.PayoutId():
		return "uuid"

	case selectPayoutApprovalsFields.ApprovalStatus():
		return "payout_approvals_approval_status_enum"

	case selectPayoutApprovalsFields.ReasonCode():
		return "text"

	case selectPayoutApprovalsFields.ReasonDetail():
		return "text"

	case selectPayoutApprovalsFields.ApprovedBy():
		return "uuid"

	case selectPayoutApprovalsFields.ApprovedAt():
		return "timestamptz"

	case selectPayoutApprovalsFields.Metadata():
		return "jsonb"

	case selectPayoutApprovalsFields.ApprovedAmountSnapshot():
		return "numeric"

	case selectPayoutApprovalsFields.CurrencyCodeSnapshot():
		return "text"

	case selectPayoutApprovalsFields.PayoutRevisionHash():
		return "text"

	case selectPayoutApprovalsFields.MetaCreatedAt():
		return "timestamptz"

	case selectPayoutApprovalsFields.MetaCreatedBy():
		return "uuid"

	case selectPayoutApprovalsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPayoutApprovalsFields.MetaUpdatedBy():
		return "uuid"

	case selectPayoutApprovalsFields.MetaDeletedAt():
		return "timestamptz"

	case selectPayoutApprovalsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePayoutApprovals(ctx context.Context, payoutApprovals *model.PayoutApprovals, fieldsInsert ...PayoutApprovalsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPayoutApprovalsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PayoutApprovalsPrimaryID{
		Id: payoutApprovals.Id,
	}
	exists, err := repo.IsExistPayoutApprovalsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutApprovals] failed checking payoutApprovals whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "payoutApprovals", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPayoutApprovals([]model.PayoutApprovals{*payoutApprovals}, fieldsInsert...)
	commandQuery := fmt.Sprintf(payoutApprovalsQueries.insertPayoutApprovals, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePayoutApprovals] failed exec create payoutApprovals query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePayoutApprovalsByID(ctx context.Context, primaryID model.PayoutApprovalsPrimaryID) (err error) {
	exists, err := repo.IsExistPayoutApprovalsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutApprovalsByID] failed checking payoutApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutApprovals with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePayoutApprovalsCompositePrimaryKeyWhere([]model.PayoutApprovalsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(payoutApprovalsQueries.deletePayoutApprovals + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePayoutApprovalsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutApprovalsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutApprovalsFilterResult, err error) {
	query, args, err := composePayoutApprovalsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutApprovalsByFilter] failed compose payoutApprovals filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutApprovalsByFilter] failed get payoutApprovals by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePayoutApprovalsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.PayoutApprovalsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composePayoutApprovalsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizePayoutApprovalsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composePayoutApprovalsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 17 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewPayoutApprovalsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 17+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_id\"")
			selectedColumns["payout_id"] = struct{}{}
		}
		if _, selected := selectedColumns["approval_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"approval_status\"")
			selectedColumns["approval_status"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_detail"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_detail\"")
			selectedColumns["reason_detail"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_by\"")
			selectedColumns["approved_by"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_at\"")
			selectedColumns["approved_at"] = struct{}{}
		}
		if _, selected := selectedColumns["metadata"]; !selected {
			selectColumns = append(selectColumns, "base.\"metadata\"")
			selectedColumns["metadata"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_amount_snapshot"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_amount_snapshot\"")
			selectedColumns["approved_amount_snapshot"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code_snapshot"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code_snapshot\"")
			selectedColumns["currency_code_snapshot"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_revision_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_revision_hash\"")
			selectedColumns["payout_revision_hash"] = struct{}{}
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

type payoutApprovalsFilterPlaceholder struct {
	index int
}

func (p *payoutApprovalsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composePayoutApprovalsFilterPredicate(filterField model.FilterField, placeholders *payoutApprovalsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewPayoutApprovalsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composePayoutApprovalsFilterSQLExpr(spec)
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

func composePayoutApprovalsFilterGroup(group model.FilterGroup, placeholders *payoutApprovalsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composePayoutApprovalsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composePayoutApprovalsFilterGroup(child, placeholders, args, requiredJoins)
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

func composePayoutApprovalsFilterWhereQueries(filter model.Filter, placeholders *payoutApprovalsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composePayoutApprovalsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composePayoutApprovalsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composePayoutApprovalsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePayoutApprovalsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizePayoutApprovalsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composePayoutApprovalsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := payoutApprovalsFilterPlaceholder{index: 1}
	whereQueries, err := composePayoutApprovalsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewPayoutApprovalsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composePayoutApprovalsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizePayoutApprovalsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"payout_approvals\" base%s", strings.Join(selectColumns, ","), composePayoutApprovalsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistPayoutApprovalsByID(ctx context.Context, primaryID model.PayoutApprovalsPrimaryID) (exists bool, err error) {
	whereQuery, params := composePayoutApprovalsCompositePrimaryKeyWhere([]model.PayoutApprovalsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", payoutApprovalsQueries.selectCountPayoutApprovals, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPayoutApprovalsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutApprovals(ctx context.Context, selectFields ...PayoutApprovalsField) (payoutApprovalsList model.PayoutApprovalsList, err error) {
	var (
		defaultPayoutApprovalsSelectFields = defaultPayoutApprovalsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutApprovalsSelectFields = composePayoutApprovalsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(payoutApprovalsQueries.selectPayoutApprovals, defaultPayoutApprovalsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &payoutApprovalsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePayoutApprovals] failed get payoutApprovals list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePayoutApprovalsByID(ctx context.Context, primaryID model.PayoutApprovalsPrimaryID, selectFields ...PayoutApprovalsField) (payoutApprovals model.PayoutApprovals, err error) {
	var (
		defaultPayoutApprovalsSelectFields = defaultPayoutApprovalsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPayoutApprovalsSelectFields = composePayoutApprovalsSelectFields(selectFields...)
	}
	whereQry, params := composePayoutApprovalsCompositePrimaryKeyWhere([]model.PayoutApprovalsPrimaryID{primaryID})
	query := fmt.Sprintf(payoutApprovalsQueries.selectPayoutApprovals+" WHERE "+whereQry, defaultPayoutApprovalsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &payoutApprovals, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("payoutApprovals with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePayoutApprovalsByID] failed get payoutApprovals")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePayoutApprovalsByID(ctx context.Context, primaryID model.PayoutApprovalsPrimaryID, payoutApprovals *model.PayoutApprovals, payoutApprovalsUpdateFields ...PayoutApprovalsUpdateField) (err error) {
	exists, err := repo.IsExistPayoutApprovalsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutApprovals] failed checking payoutApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("payoutApprovals with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if payoutApprovals == nil {
		if len(payoutApprovalsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePayoutApprovalsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		payoutApprovals = &model.PayoutApprovals{}
	}
	var (
		defaultPayoutApprovalsUpdateFields = defaultPayoutApprovalsUpdateFields(*payoutApprovals)
		tempUpdateField                    PayoutApprovalsUpdateFieldList
		selectFields                       = NewPayoutApprovalsSelectFields()
	)
	if len(payoutApprovalsUpdateFields) > 0 {
		for _, updateField := range payoutApprovalsUpdateFields {
			if updateField.payoutApprovalsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPayoutApprovalsUpdateFields = tempUpdateField
	}
	whereQuery, params := composePayoutApprovalsCompositePrimaryKeyWhere([]model.PayoutApprovalsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPayoutApprovalsCommand(defaultPayoutApprovalsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(payoutApprovalsQueries.updatePayoutApprovals+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutApprovals] error when try to update payoutApprovals by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdatePayoutApprovalsByFilter(ctx context.Context, filter model.Filter, payoutApprovalsUpdateFields ...PayoutApprovalsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(payoutApprovalsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields PayoutApprovalsUpdateFieldList
		selectFields = NewPayoutApprovalsSelectFields()
	)
	for _, updateField := range payoutApprovalsUpdateFields {
		if updateField.payoutApprovalsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsPayoutApprovalsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := payoutApprovalsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composePayoutApprovalsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"payout_approvals\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutApprovalsByFilter] error when try to update payoutApprovals by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePayoutApprovalsByFilter] failed get rows affected")
	}
	return
}

var (
	payoutApprovalsQueries = struct {
		selectPayoutApprovals      string
		selectCountPayoutApprovals string
		deletePayoutApprovals      string
		updatePayoutApprovals      string
		insertPayoutApprovals      string
	}{
		selectPayoutApprovals:      "SELECT %s FROM \"payout_approvals\"",
		selectCountPayoutApprovals: "SELECT COUNT(\"id\") FROM \"payout_approvals\"",
		deletePayoutApprovals:      "DELETE FROM \"payout_approvals\"",
		updatePayoutApprovals:      "UPDATE \"payout_approvals\" SET %s ",
		insertPayoutApprovals:      "INSERT INTO \"payout_approvals\" %s VALUES %s",
	}
)

type PayoutApprovalsRepository interface {
	CreatePayoutApprovals(ctx context.Context, payoutApprovals *model.PayoutApprovals, fieldsInsert ...PayoutApprovalsField) error
	BulkCreatePayoutApprovals(ctx context.Context, payoutApprovalsList []*model.PayoutApprovals, fieldsInsert ...PayoutApprovalsField) error
	ResolvePayoutApprovals(ctx context.Context, selectFields ...PayoutApprovalsField) (model.PayoutApprovalsList, error)
	ResolvePayoutApprovalsByID(ctx context.Context, primaryID model.PayoutApprovalsPrimaryID, selectFields ...PayoutApprovalsField) (model.PayoutApprovals, error)
	UpdatePayoutApprovalsByID(ctx context.Context, id model.PayoutApprovalsPrimaryID, payoutApprovals *model.PayoutApprovals, payoutApprovalsUpdateFields ...PayoutApprovalsUpdateField) error
	UpdatePayoutApprovalsByFilter(ctx context.Context, filter model.Filter, payoutApprovalsUpdateFields ...PayoutApprovalsUpdateField) (rowsAffected int64, err error)
	BulkUpdatePayoutApprovals(ctx context.Context, payoutApprovalsListMap map[model.PayoutApprovalsPrimaryID]*model.PayoutApprovals, PayoutApprovalssMapUpdateFieldsRequest map[model.PayoutApprovalsPrimaryID]PayoutApprovalsUpdateFieldList) (err error)
	DeletePayoutApprovalsByID(ctx context.Context, id model.PayoutApprovalsPrimaryID) error
	BulkDeletePayoutApprovalsByIDs(ctx context.Context, ids []model.PayoutApprovalsPrimaryID) error
	ResolvePayoutApprovalsByFilter(ctx context.Context, filter model.Filter) (result []model.PayoutApprovalsFilterResult, err error)
	IsExistPayoutApprovalsByIDs(ctx context.Context, ids []model.PayoutApprovalsPrimaryID) (exists bool, notFoundIds []model.PayoutApprovalsPrimaryID, err error)
	IsExistPayoutApprovalsByID(ctx context.Context, id model.PayoutApprovalsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
