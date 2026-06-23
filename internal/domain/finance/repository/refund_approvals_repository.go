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

func composeInsertFieldsAndParamsRefundApprovals(refundApprovalsList []model.RefundApprovals, fieldsInsert ...RefundApprovalsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRefundApprovalsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, refundApprovals := range refundApprovalsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, refundApprovals.Id)
			case selectField.RefundId():
				args = append(args, refundApprovals.RefundId)
			case selectField.ApprovalStatus():
				args = append(args, refundApprovals.ApprovalStatus)
			case selectField.ApprovedAmount():
				args = append(args, refundApprovals.ApprovedAmount)
			case selectField.ReasonCode():
				args = append(args, refundApprovals.ReasonCode)
			case selectField.ReasonDetail():
				args = append(args, refundApprovals.ReasonDetail)
			case selectField.ApprovedBy():
				args = append(args, refundApprovals.ApprovedBy)
			case selectField.ApprovedAt():
				args = append(args, refundApprovals.ApprovedAt)
			case selectField.Metadata():
				args = append(args, refundApprovals.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, refundApprovals.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, refundApprovals.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, refundApprovals.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, refundApprovals.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, refundApprovals.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, refundApprovals.MetaDeletedBy)

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

func composeRefundApprovalsCompositePrimaryKeyWhere(primaryIDs []model.RefundApprovalsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"refund_approvals\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRefundApprovalsSelectFields() string {
	fields := NewRefundApprovalsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRefundApprovalsSelectFields(selectFields ...RefundApprovalsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RefundApprovalsField string
type RefundApprovalsFieldList []RefundApprovalsField

type RefundApprovalsSelectFields struct {
}

func (ss RefundApprovalsSelectFields) Id() RefundApprovalsField {
	return RefundApprovalsField("id")
}

func (ss RefundApprovalsSelectFields) RefundId() RefundApprovalsField {
	return RefundApprovalsField("refund_id")
}

func (ss RefundApprovalsSelectFields) ApprovalStatus() RefundApprovalsField {
	return RefundApprovalsField("approval_status")
}

func (ss RefundApprovalsSelectFields) ApprovedAmount() RefundApprovalsField {
	return RefundApprovalsField("approved_amount")
}

func (ss RefundApprovalsSelectFields) ReasonCode() RefundApprovalsField {
	return RefundApprovalsField("reason_code")
}

func (ss RefundApprovalsSelectFields) ReasonDetail() RefundApprovalsField {
	return RefundApprovalsField("reason_detail")
}

func (ss RefundApprovalsSelectFields) ApprovedBy() RefundApprovalsField {
	return RefundApprovalsField("approved_by")
}

func (ss RefundApprovalsSelectFields) ApprovedAt() RefundApprovalsField {
	return RefundApprovalsField("approved_at")
}

func (ss RefundApprovalsSelectFields) Metadata() RefundApprovalsField {
	return RefundApprovalsField("metadata")
}

func (ss RefundApprovalsSelectFields) MetaCreatedAt() RefundApprovalsField {
	return RefundApprovalsField("meta_created_at")
}

func (ss RefundApprovalsSelectFields) MetaCreatedBy() RefundApprovalsField {
	return RefundApprovalsField("meta_created_by")
}

func (ss RefundApprovalsSelectFields) MetaUpdatedAt() RefundApprovalsField {
	return RefundApprovalsField("meta_updated_at")
}

func (ss RefundApprovalsSelectFields) MetaUpdatedBy() RefundApprovalsField {
	return RefundApprovalsField("meta_updated_by")
}

func (ss RefundApprovalsSelectFields) MetaDeletedAt() RefundApprovalsField {
	return RefundApprovalsField("meta_deleted_at")
}

func (ss RefundApprovalsSelectFields) MetaDeletedBy() RefundApprovalsField {
	return RefundApprovalsField("meta_deleted_by")
}

func (ss RefundApprovalsSelectFields) All() RefundApprovalsFieldList {
	return []RefundApprovalsField{
		ss.Id(),
		ss.RefundId(),
		ss.ApprovalStatus(),
		ss.ApprovedAmount(),
		ss.ReasonCode(),
		ss.ReasonDetail(),
		ss.ApprovedBy(),
		ss.ApprovedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRefundApprovalsSelectFields() RefundApprovalsSelectFields {
	return RefundApprovalsSelectFields{}
}

type RefundApprovalsUpdateFieldOption struct {
	useIncrement bool
}
type RefundApprovalsUpdateField struct {
	refundApprovalsField RefundApprovalsField
	opt                  RefundApprovalsUpdateFieldOption
	value                interface{}
}
type RefundApprovalsUpdateFieldList []RefundApprovalsUpdateField

func defaultRefundApprovalsUpdateFieldOption() RefundApprovalsUpdateFieldOption {
	return RefundApprovalsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRefundApprovalsOption(useIncrement bool) func(*RefundApprovalsUpdateFieldOption) {
	return func(pcufo *RefundApprovalsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRefundApprovalsUpdateField(field RefundApprovalsField, val interface{}, opts ...func(*RefundApprovalsUpdateFieldOption)) RefundApprovalsUpdateField {
	defaultOpt := defaultRefundApprovalsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RefundApprovalsUpdateField{
		refundApprovalsField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultRefundApprovalsUpdateFields(refundApprovals model.RefundApprovals) (refundApprovalsUpdateFieldList RefundApprovalsUpdateFieldList) {
	selectFields := NewRefundApprovalsSelectFields()
	refundApprovalsUpdateFieldList = append(refundApprovalsUpdateFieldList,
		NewRefundApprovalsUpdateField(selectFields.Id(), refundApprovals.Id),
		NewRefundApprovalsUpdateField(selectFields.RefundId(), refundApprovals.RefundId),
		NewRefundApprovalsUpdateField(selectFields.ApprovalStatus(), refundApprovals.ApprovalStatus),
		NewRefundApprovalsUpdateField(selectFields.ApprovedAmount(), refundApprovals.ApprovedAmount),
		NewRefundApprovalsUpdateField(selectFields.ReasonCode(), refundApprovals.ReasonCode),
		NewRefundApprovalsUpdateField(selectFields.ReasonDetail(), refundApprovals.ReasonDetail),
		NewRefundApprovalsUpdateField(selectFields.ApprovedBy(), refundApprovals.ApprovedBy),
		NewRefundApprovalsUpdateField(selectFields.ApprovedAt(), refundApprovals.ApprovedAt),
		NewRefundApprovalsUpdateField(selectFields.Metadata(), refundApprovals.Metadata),
		NewRefundApprovalsUpdateField(selectFields.MetaCreatedAt(), refundApprovals.MetaCreatedAt),
		NewRefundApprovalsUpdateField(selectFields.MetaCreatedBy(), refundApprovals.MetaCreatedBy),
		NewRefundApprovalsUpdateField(selectFields.MetaUpdatedAt(), refundApprovals.MetaUpdatedAt),
		NewRefundApprovalsUpdateField(selectFields.MetaUpdatedBy(), refundApprovals.MetaUpdatedBy),
		NewRefundApprovalsUpdateField(selectFields.MetaDeletedAt(), refundApprovals.MetaDeletedAt),
		NewRefundApprovalsUpdateField(selectFields.MetaDeletedBy(), refundApprovals.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRefundApprovalsCommand(refundApprovalsUpdateFieldList RefundApprovalsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range refundApprovalsUpdateFieldList {
		field := string(updateField.refundApprovalsField)
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

func (repo *RepositoryImpl) BulkCreateRefundApprovals(ctx context.Context, refundApprovalsList []*model.RefundApprovals, fieldsInsert ...RefundApprovalsField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.RefundApprovalsPrimaryID
		refundApprovalsValueList []model.RefundApprovals
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRefundApprovalsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, refundApprovals := range refundApprovalsList {

		primaryIds = append(primaryIds, refundApprovals.ToRefundApprovalsPrimaryID())

		refundApprovalsValueList = append(refundApprovalsValueList, *refundApprovals)
	}

	_, notFoundIds, err := repo.IsExistRefundApprovalsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundApprovals] failed checking refundApprovals whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RefundApprovalsPrimaryID{}
		mapNotFoundIds := map[model.RefundApprovalsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "refundApprovals", fmt.Sprintf("refundApprovals with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRefundApprovals(refundApprovalsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(refundApprovalsQueries.insertRefundApprovals, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundApprovals] failed exec create refundApprovals query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRefundApprovalsByIDs(ctx context.Context, primaryIDs []model.RefundApprovalsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRefundApprovalsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundApprovalsByIDs] failed checking refundApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundApprovals with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_approvals\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(refundApprovalsQueries.deleteRefundApprovals + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundApprovalsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundApprovalsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRefundApprovalsByIDs(ctx context.Context, ids []model.RefundApprovalsPrimaryID) (exists bool, notFoundIds []model.RefundApprovalsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_approvals\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(refundApprovalsQueries.selectRefundApprovals, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundApprovalsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RefundApprovalsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundApprovalsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RefundApprovalsPrimaryID]bool{}
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

// BulkUpdateRefundApprovals is used to bulk update refundApprovals, by default it will update all field
// if want to update specific field, then fill refundApprovalssMapUpdateFieldsRequest else please fill refundApprovalssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRefundApprovals(ctx context.Context, refundApprovalssMap map[model.RefundApprovalsPrimaryID]*model.RefundApprovals, refundApprovalssMapUpdateFieldsRequest map[model.RefundApprovalsPrimaryID]RefundApprovalsUpdateFieldList) (err error) {
	if len(refundApprovalssMap) == 0 && len(refundApprovalssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		refundApprovalssMapUpdateField map[model.RefundApprovalsPrimaryID]RefundApprovalsUpdateFieldList = map[model.RefundApprovalsPrimaryID]RefundApprovalsUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(refundApprovalssMap) > 0 {
		for id, refundApprovals := range refundApprovalssMap {
			if refundApprovals == nil {
				log.Error().Err(err).Msg("[BulkUpdateRefundApprovals] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			refundApprovalssMapUpdateField[id] = defaultRefundApprovalsUpdateFields(*refundApprovals)
		}
	} else {
		refundApprovalssMapUpdateField = refundApprovalssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRefundApprovalsQuery(refundApprovalssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRefundApprovalsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundApprovals] failed checking refundApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundApprovals with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRefundApprovalsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"refund_approvals\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundApprovals] failed exec query")
	}
	return
}

type RefundApprovalsFieldParameter struct {
	param string
	args  []interface{}
}

func NewRefundApprovalsFieldParameter(param string, args ...interface{}) RefundApprovalsFieldParameter {
	return RefundApprovalsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRefundApprovalsQuery(mapRefundApprovalss map[model.RefundApprovalsPrimaryID]RefundApprovalsUpdateFieldList, asTableValues string) (primaryIDs []model.RefundApprovalsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RefundApprovalsPrimaryID]map[string]interface{}{}
	refundApprovalsSelectFields := NewRefundApprovalsSelectFields()
	for id, updateFields := range mapRefundApprovalss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.refundApprovalsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRefundApprovalss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRefundApprovalsFieldType(updateField.refundApprovalsField)))
			args = append(args, fields[string(updateField.refundApprovalsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.refundApprovalsField))
		if updateField.refundApprovalsField == refundApprovalsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.refundApprovalsField, asTableValues, updateField.refundApprovalsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.refundApprovalsField,
				"\"refund_approvals\"", updateField.refundApprovalsField,
				asTableValues, updateField.refundApprovalsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRefundApprovalsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RefundApprovalsPrimaryID, asTableValue string) (whereQry string) {
	refundApprovalsSelectFields := NewRefundApprovalsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"refund_approvals\".\"id\" = %s.\"id\"::"+GetRefundApprovalsFieldType(refundApprovalsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRefundApprovalsFieldType(refundApprovalsField RefundApprovalsField) string {
	selectRefundApprovalsFields := NewRefundApprovalsSelectFields()
	switch refundApprovalsField {

	case selectRefundApprovalsFields.Id():
		return "uuid"

	case selectRefundApprovalsFields.RefundId():
		return "uuid"

	case selectRefundApprovalsFields.ApprovalStatus():
		return "refund_approvals_approval_status_enum"

	case selectRefundApprovalsFields.ApprovedAmount():
		return "numeric"

	case selectRefundApprovalsFields.ReasonCode():
		return "text"

	case selectRefundApprovalsFields.ReasonDetail():
		return "text"

	case selectRefundApprovalsFields.ApprovedBy():
		return "uuid"

	case selectRefundApprovalsFields.ApprovedAt():
		return "timestamptz"

	case selectRefundApprovalsFields.Metadata():
		return "jsonb"

	case selectRefundApprovalsFields.MetaCreatedAt():
		return "timestamptz"

	case selectRefundApprovalsFields.MetaCreatedBy():
		return "uuid"

	case selectRefundApprovalsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRefundApprovalsFields.MetaUpdatedBy():
		return "uuid"

	case selectRefundApprovalsFields.MetaDeletedAt():
		return "timestamptz"

	case selectRefundApprovalsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRefundApprovals(ctx context.Context, refundApprovals *model.RefundApprovals, fieldsInsert ...RefundApprovalsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRefundApprovalsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RefundApprovalsPrimaryID{
		Id: refundApprovals.Id,
	}
	exists, err := repo.IsExistRefundApprovalsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundApprovals] failed checking refundApprovals whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "refundApprovals", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRefundApprovals([]model.RefundApprovals{*refundApprovals}, fieldsInsert...)
	commandQuery := fmt.Sprintf(refundApprovalsQueries.insertRefundApprovals, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundApprovals] failed exec create refundApprovals query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRefundApprovalsByID(ctx context.Context, primaryID model.RefundApprovalsPrimaryID) (err error) {
	exists, err := repo.IsExistRefundApprovalsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundApprovalsByID] failed checking refundApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundApprovals with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRefundApprovalsCompositePrimaryKeyWhere([]model.RefundApprovalsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(refundApprovalsQueries.deleteRefundApprovals + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundApprovalsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundApprovalsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundApprovalsFilterResult, err error) {
	query, args, err := composeRefundApprovalsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundApprovalsByFilter] failed compose refundApprovals filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundApprovalsByFilter] failed get refundApprovals by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRefundApprovalsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.RefundApprovalsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeRefundApprovalsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeRefundApprovalsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeRefundApprovalsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewRefundApprovalsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["refund_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_id\"")
			selectedColumns["refund_id"] = struct{}{}
		}
		if _, selected := selectedColumns["approval_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"approval_status\"")
			selectedColumns["approval_status"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_amount\"")
			selectedColumns["approved_amount"] = struct{}{}
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

type refundApprovalsFilterPlaceholder struct {
	index int
}

func (p *refundApprovalsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeRefundApprovalsFilterPredicate(filterField model.FilterField, placeholders *refundApprovalsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewRefundApprovalsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeRefundApprovalsFilterSQLExpr(spec)
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

func composeRefundApprovalsFilterGroup(group model.FilterGroup, placeholders *refundApprovalsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeRefundApprovalsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeRefundApprovalsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeRefundApprovalsFilterWhereQueries(filter model.Filter, placeholders *refundApprovalsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeRefundApprovalsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeRefundApprovalsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeRefundApprovalsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRefundApprovalsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeRefundApprovalsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeRefundApprovalsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := refundApprovalsFilterPlaceholder{index: 1}
	whereQueries, err := composeRefundApprovalsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewRefundApprovalsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeRefundApprovalsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeRefundApprovalsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"refund_approvals\" base%s", strings.Join(selectColumns, ","), composeRefundApprovalsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistRefundApprovalsByID(ctx context.Context, primaryID model.RefundApprovalsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRefundApprovalsCompositePrimaryKeyWhere([]model.RefundApprovalsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", refundApprovalsQueries.selectCountRefundApprovals, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundApprovalsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundApprovals(ctx context.Context, selectFields ...RefundApprovalsField) (refundApprovalsList model.RefundApprovalsList, err error) {
	var (
		defaultRefundApprovalsSelectFields = defaultRefundApprovalsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundApprovalsSelectFields = composeRefundApprovalsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(refundApprovalsQueries.selectRefundApprovals, defaultRefundApprovalsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &refundApprovalsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundApprovals] failed get refundApprovals list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundApprovalsByID(ctx context.Context, primaryID model.RefundApprovalsPrimaryID, selectFields ...RefundApprovalsField) (refundApprovals model.RefundApprovals, err error) {
	var (
		defaultRefundApprovalsSelectFields = defaultRefundApprovalsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundApprovalsSelectFields = composeRefundApprovalsSelectFields(selectFields...)
	}
	whereQry, params := composeRefundApprovalsCompositePrimaryKeyWhere([]model.RefundApprovalsPrimaryID{primaryID})
	query := fmt.Sprintf(refundApprovalsQueries.selectRefundApprovals+" WHERE "+whereQry, defaultRefundApprovalsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &refundApprovals, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("refundApprovals with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRefundApprovalsByID] failed get refundApprovals")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRefundApprovalsByID(ctx context.Context, primaryID model.RefundApprovalsPrimaryID, refundApprovals *model.RefundApprovals, refundApprovalsUpdateFields ...RefundApprovalsUpdateField) (err error) {
	exists, err := repo.IsExistRefundApprovalsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundApprovals] failed checking refundApprovals whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundApprovals with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if refundApprovals == nil {
		if len(refundApprovalsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRefundApprovalsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		refundApprovals = &model.RefundApprovals{}
	}
	var (
		defaultRefundApprovalsUpdateFields = defaultRefundApprovalsUpdateFields(*refundApprovals)
		tempUpdateField                    RefundApprovalsUpdateFieldList
		selectFields                       = NewRefundApprovalsSelectFields()
	)
	if len(refundApprovalsUpdateFields) > 0 {
		for _, updateField := range refundApprovalsUpdateFields {
			if updateField.refundApprovalsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRefundApprovalsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRefundApprovalsCompositePrimaryKeyWhere([]model.RefundApprovalsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRefundApprovalsCommand(defaultRefundApprovalsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(refundApprovalsQueries.updateRefundApprovals+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundApprovals] error when try to update refundApprovals by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateRefundApprovalsByFilter(ctx context.Context, filter model.Filter, refundApprovalsUpdateFields ...RefundApprovalsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(refundApprovalsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields RefundApprovalsUpdateFieldList
		selectFields = NewRefundApprovalsSelectFields()
	)
	for _, updateField := range refundApprovalsUpdateFields {
		if updateField.refundApprovalsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsRefundApprovalsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := refundApprovalsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeRefundApprovalsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"refund_approvals\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundApprovalsByFilter] error when try to update refundApprovals by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundApprovalsByFilter] failed get rows affected")
	}
	return
}

var (
	refundApprovalsQueries = struct {
		selectRefundApprovals      string
		selectCountRefundApprovals string
		deleteRefundApprovals      string
		updateRefundApprovals      string
		insertRefundApprovals      string
	}{
		selectRefundApprovals:      "SELECT %s FROM \"refund_approvals\"",
		selectCountRefundApprovals: "SELECT COUNT(\"id\") FROM \"refund_approvals\"",
		deleteRefundApprovals:      "DELETE FROM \"refund_approvals\"",
		updateRefundApprovals:      "UPDATE \"refund_approvals\" SET %s ",
		insertRefundApprovals:      "INSERT INTO \"refund_approvals\" %s VALUES %s",
	}
)

type RefundApprovalsRepository interface {
	CreateRefundApprovals(ctx context.Context, refundApprovals *model.RefundApprovals, fieldsInsert ...RefundApprovalsField) error
	BulkCreateRefundApprovals(ctx context.Context, refundApprovalsList []*model.RefundApprovals, fieldsInsert ...RefundApprovalsField) error
	ResolveRefundApprovals(ctx context.Context, selectFields ...RefundApprovalsField) (model.RefundApprovalsList, error)
	ResolveRefundApprovalsByID(ctx context.Context, primaryID model.RefundApprovalsPrimaryID, selectFields ...RefundApprovalsField) (model.RefundApprovals, error)
	UpdateRefundApprovalsByID(ctx context.Context, id model.RefundApprovalsPrimaryID, refundApprovals *model.RefundApprovals, refundApprovalsUpdateFields ...RefundApprovalsUpdateField) error
	UpdateRefundApprovalsByFilter(ctx context.Context, filter model.Filter, refundApprovalsUpdateFields ...RefundApprovalsUpdateField) (rowsAffected int64, err error)
	BulkUpdateRefundApprovals(ctx context.Context, refundApprovalsListMap map[model.RefundApprovalsPrimaryID]*model.RefundApprovals, RefundApprovalssMapUpdateFieldsRequest map[model.RefundApprovalsPrimaryID]RefundApprovalsUpdateFieldList) (err error)
	DeleteRefundApprovalsByID(ctx context.Context, id model.RefundApprovalsPrimaryID) error
	BulkDeleteRefundApprovalsByIDs(ctx context.Context, ids []model.RefundApprovalsPrimaryID) error
	ResolveRefundApprovalsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundApprovalsFilterResult, err error)
	IsExistRefundApprovalsByIDs(ctx context.Context, ids []model.RefundApprovalsPrimaryID) (exists bool, notFoundIds []model.RefundApprovalsPrimaryID, err error)
	IsExistRefundApprovalsByID(ctx context.Context, id model.RefundApprovalsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
