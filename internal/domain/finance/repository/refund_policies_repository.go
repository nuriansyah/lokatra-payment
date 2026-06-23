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

func composeInsertFieldsAndParamsRefundPolicies(refundPoliciesList []model.RefundPolicies, fieldsInsert ...RefundPoliciesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRefundPoliciesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, refundPolicies := range refundPoliciesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, refundPolicies.Id)
			case selectField.PolicyCode():
				args = append(args, refundPolicies.PolicyCode)
			case selectField.MerchantPartyId():
				args = append(args, refundPolicies.MerchantPartyId)
			case selectField.PolicyScope():
				args = append(args, refundPolicies.PolicyScope)
			case selectField.AllowPartial():
				args = append(args, refundPolicies.AllowPartial)
			case selectField.AllowPostPayout():
				args = append(args, refundPolicies.AllowPostPayout)
			case selectField.FeeReturnMode():
				args = append(args, refundPolicies.FeeReturnMode)
			case selectField.TaxReturnMode():
				args = append(args, refundPolicies.TaxReturnMode)
			case selectField.RequiresApprovalOverAmount():
				args = append(args, refundPolicies.RequiresApprovalOverAmount)
			case selectField.PolicyStatus():
				args = append(args, refundPolicies.PolicyStatus)
			case selectField.Metadata():
				args = append(args, refundPolicies.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, refundPolicies.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, refundPolicies.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, refundPolicies.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, refundPolicies.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, refundPolicies.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, refundPolicies.MetaDeletedBy)

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

func composeRefundPoliciesCompositePrimaryKeyWhere(primaryIDs []model.RefundPoliciesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"refund_policies\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRefundPoliciesSelectFields() string {
	fields := NewRefundPoliciesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRefundPoliciesSelectFields(selectFields ...RefundPoliciesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RefundPoliciesField string
type RefundPoliciesFieldList []RefundPoliciesField

type RefundPoliciesSelectFields struct {
}

func (ss RefundPoliciesSelectFields) Id() RefundPoliciesField {
	return RefundPoliciesField("id")
}

func (ss RefundPoliciesSelectFields) PolicyCode() RefundPoliciesField {
	return RefundPoliciesField("policy_code")
}

func (ss RefundPoliciesSelectFields) MerchantPartyId() RefundPoliciesField {
	return RefundPoliciesField("merchant_party_id")
}

func (ss RefundPoliciesSelectFields) PolicyScope() RefundPoliciesField {
	return RefundPoliciesField("policy_scope")
}

func (ss RefundPoliciesSelectFields) AllowPartial() RefundPoliciesField {
	return RefundPoliciesField("allow_partial")
}

func (ss RefundPoliciesSelectFields) AllowPostPayout() RefundPoliciesField {
	return RefundPoliciesField("allow_post_payout")
}

func (ss RefundPoliciesSelectFields) FeeReturnMode() RefundPoliciesField {
	return RefundPoliciesField("fee_return_mode")
}

func (ss RefundPoliciesSelectFields) TaxReturnMode() RefundPoliciesField {
	return RefundPoliciesField("tax_return_mode")
}

func (ss RefundPoliciesSelectFields) RequiresApprovalOverAmount() RefundPoliciesField {
	return RefundPoliciesField("requires_approval_over_amount")
}

func (ss RefundPoliciesSelectFields) PolicyStatus() RefundPoliciesField {
	return RefundPoliciesField("policy_status")
}

func (ss RefundPoliciesSelectFields) Metadata() RefundPoliciesField {
	return RefundPoliciesField("metadata")
}

func (ss RefundPoliciesSelectFields) MetaCreatedAt() RefundPoliciesField {
	return RefundPoliciesField("meta_created_at")
}

func (ss RefundPoliciesSelectFields) MetaCreatedBy() RefundPoliciesField {
	return RefundPoliciesField("meta_created_by")
}

func (ss RefundPoliciesSelectFields) MetaUpdatedAt() RefundPoliciesField {
	return RefundPoliciesField("meta_updated_at")
}

func (ss RefundPoliciesSelectFields) MetaUpdatedBy() RefundPoliciesField {
	return RefundPoliciesField("meta_updated_by")
}

func (ss RefundPoliciesSelectFields) MetaDeletedAt() RefundPoliciesField {
	return RefundPoliciesField("meta_deleted_at")
}

func (ss RefundPoliciesSelectFields) MetaDeletedBy() RefundPoliciesField {
	return RefundPoliciesField("meta_deleted_by")
}

func (ss RefundPoliciesSelectFields) All() RefundPoliciesFieldList {
	return []RefundPoliciesField{
		ss.Id(),
		ss.PolicyCode(),
		ss.MerchantPartyId(),
		ss.PolicyScope(),
		ss.AllowPartial(),
		ss.AllowPostPayout(),
		ss.FeeReturnMode(),
		ss.TaxReturnMode(),
		ss.RequiresApprovalOverAmount(),
		ss.PolicyStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRefundPoliciesSelectFields() RefundPoliciesSelectFields {
	return RefundPoliciesSelectFields{}
}

type RefundPoliciesUpdateFieldOption struct {
	useIncrement bool
}
type RefundPoliciesUpdateField struct {
	refundPoliciesField RefundPoliciesField
	opt                 RefundPoliciesUpdateFieldOption
	value               interface{}
}
type RefundPoliciesUpdateFieldList []RefundPoliciesUpdateField

func defaultRefundPoliciesUpdateFieldOption() RefundPoliciesUpdateFieldOption {
	return RefundPoliciesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRefundPoliciesOption(useIncrement bool) func(*RefundPoliciesUpdateFieldOption) {
	return func(pcufo *RefundPoliciesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRefundPoliciesUpdateField(field RefundPoliciesField, val interface{}, opts ...func(*RefundPoliciesUpdateFieldOption)) RefundPoliciesUpdateField {
	defaultOpt := defaultRefundPoliciesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RefundPoliciesUpdateField{
		refundPoliciesField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultRefundPoliciesUpdateFields(refundPolicies model.RefundPolicies) (refundPoliciesUpdateFieldList RefundPoliciesUpdateFieldList) {
	selectFields := NewRefundPoliciesSelectFields()
	refundPoliciesUpdateFieldList = append(refundPoliciesUpdateFieldList,
		NewRefundPoliciesUpdateField(selectFields.Id(), refundPolicies.Id),
		NewRefundPoliciesUpdateField(selectFields.PolicyCode(), refundPolicies.PolicyCode),
		NewRefundPoliciesUpdateField(selectFields.MerchantPartyId(), refundPolicies.MerchantPartyId),
		NewRefundPoliciesUpdateField(selectFields.PolicyScope(), refundPolicies.PolicyScope),
		NewRefundPoliciesUpdateField(selectFields.AllowPartial(), refundPolicies.AllowPartial),
		NewRefundPoliciesUpdateField(selectFields.AllowPostPayout(), refundPolicies.AllowPostPayout),
		NewRefundPoliciesUpdateField(selectFields.FeeReturnMode(), refundPolicies.FeeReturnMode),
		NewRefundPoliciesUpdateField(selectFields.TaxReturnMode(), refundPolicies.TaxReturnMode),
		NewRefundPoliciesUpdateField(selectFields.RequiresApprovalOverAmount(), refundPolicies.RequiresApprovalOverAmount),
		NewRefundPoliciesUpdateField(selectFields.PolicyStatus(), refundPolicies.PolicyStatus),
		NewRefundPoliciesUpdateField(selectFields.Metadata(), refundPolicies.Metadata),
		NewRefundPoliciesUpdateField(selectFields.MetaCreatedAt(), refundPolicies.MetaCreatedAt),
		NewRefundPoliciesUpdateField(selectFields.MetaCreatedBy(), refundPolicies.MetaCreatedBy),
		NewRefundPoliciesUpdateField(selectFields.MetaUpdatedAt(), refundPolicies.MetaUpdatedAt),
		NewRefundPoliciesUpdateField(selectFields.MetaUpdatedBy(), refundPolicies.MetaUpdatedBy),
		NewRefundPoliciesUpdateField(selectFields.MetaDeletedAt(), refundPolicies.MetaDeletedAt),
		NewRefundPoliciesUpdateField(selectFields.MetaDeletedBy(), refundPolicies.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRefundPoliciesCommand(refundPoliciesUpdateFieldList RefundPoliciesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range refundPoliciesUpdateFieldList {
		field := string(updateField.refundPoliciesField)
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

func (repo *RepositoryImpl) BulkCreateRefundPolicies(ctx context.Context, refundPoliciesList []*model.RefundPolicies, fieldsInsert ...RefundPoliciesField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.RefundPoliciesPrimaryID
		refundPoliciesValueList []model.RefundPolicies
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRefundPoliciesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, refundPolicies := range refundPoliciesList {

		primaryIds = append(primaryIds, refundPolicies.ToRefundPoliciesPrimaryID())

		refundPoliciesValueList = append(refundPoliciesValueList, *refundPolicies)
	}

	_, notFoundIds, err := repo.IsExistRefundPoliciesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundPolicies] failed checking refundPolicies whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RefundPoliciesPrimaryID{}
		mapNotFoundIds := map[model.RefundPoliciesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "refundPolicies", fmt.Sprintf("refundPolicies with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRefundPolicies(refundPoliciesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(refundPoliciesQueries.insertRefundPolicies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundPolicies] failed exec create refundPolicies query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRefundPoliciesByIDs(ctx context.Context, primaryIDs []model.RefundPoliciesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRefundPoliciesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundPoliciesByIDs] failed checking refundPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundPolicies with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_policies\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(refundPoliciesQueries.deleteRefundPolicies + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundPoliciesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundPoliciesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRefundPoliciesByIDs(ctx context.Context, ids []model.RefundPoliciesPrimaryID) (exists bool, notFoundIds []model.RefundPoliciesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_policies\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(refundPoliciesQueries.selectRefundPolicies, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundPoliciesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RefundPoliciesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundPoliciesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RefundPoliciesPrimaryID]bool{}
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

// BulkUpdateRefundPolicies is used to bulk update refundPolicies, by default it will update all field
// if want to update specific field, then fill refundPoliciessMapUpdateFieldsRequest else please fill refundPoliciessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRefundPolicies(ctx context.Context, refundPoliciessMap map[model.RefundPoliciesPrimaryID]*model.RefundPolicies, refundPoliciessMapUpdateFieldsRequest map[model.RefundPoliciesPrimaryID]RefundPoliciesUpdateFieldList) (err error) {
	if len(refundPoliciessMap) == 0 && len(refundPoliciessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		refundPoliciessMapUpdateField map[model.RefundPoliciesPrimaryID]RefundPoliciesUpdateFieldList = map[model.RefundPoliciesPrimaryID]RefundPoliciesUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(refundPoliciessMap) > 0 {
		for id, refundPolicies := range refundPoliciessMap {
			if refundPolicies == nil {
				log.Error().Err(err).Msg("[BulkUpdateRefundPolicies] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			refundPoliciessMapUpdateField[id] = defaultRefundPoliciesUpdateFields(*refundPolicies)
		}
	} else {
		refundPoliciessMapUpdateField = refundPoliciessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRefundPoliciesQuery(refundPoliciessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRefundPoliciesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundPolicies] failed checking refundPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundPolicies with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRefundPoliciesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"refund_policies\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundPolicies] failed exec query")
	}
	return
}

type RefundPoliciesFieldParameter struct {
	param string
	args  []interface{}
}

func NewRefundPoliciesFieldParameter(param string, args ...interface{}) RefundPoliciesFieldParameter {
	return RefundPoliciesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRefundPoliciesQuery(mapRefundPoliciess map[model.RefundPoliciesPrimaryID]RefundPoliciesUpdateFieldList, asTableValues string) (primaryIDs []model.RefundPoliciesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RefundPoliciesPrimaryID]map[string]interface{}{}
	refundPoliciesSelectFields := NewRefundPoliciesSelectFields()
	for id, updateFields := range mapRefundPoliciess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.refundPoliciesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRefundPoliciess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRefundPoliciesFieldType(updateField.refundPoliciesField)))
			args = append(args, fields[string(updateField.refundPoliciesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.refundPoliciesField))
		if updateField.refundPoliciesField == refundPoliciesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.refundPoliciesField, asTableValues, updateField.refundPoliciesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.refundPoliciesField,
				"\"refund_policies\"", updateField.refundPoliciesField,
				asTableValues, updateField.refundPoliciesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRefundPoliciesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RefundPoliciesPrimaryID, asTableValue string) (whereQry string) {
	refundPoliciesSelectFields := NewRefundPoliciesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"refund_policies\".\"id\" = %s.\"id\"::"+GetRefundPoliciesFieldType(refundPoliciesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRefundPoliciesFieldType(refundPoliciesField RefundPoliciesField) string {
	selectRefundPoliciesFields := NewRefundPoliciesSelectFields()
	switch refundPoliciesField {

	case selectRefundPoliciesFields.Id():
		return "uuid"

	case selectRefundPoliciesFields.PolicyCode():
		return "text"

	case selectRefundPoliciesFields.MerchantPartyId():
		return "uuid"

	case selectRefundPoliciesFields.PolicyScope():
		return "refund_policies_policy_scope_enum"

	case selectRefundPoliciesFields.AllowPartial():
		return "bool"

	case selectRefundPoliciesFields.AllowPostPayout():
		return "bool"

	case selectRefundPoliciesFields.FeeReturnMode():
		return "fee_return_mode_enum"

	case selectRefundPoliciesFields.TaxReturnMode():
		return "tax_return_mode_enum"

	case selectRefundPoliciesFields.RequiresApprovalOverAmount():
		return "numeric"

	case selectRefundPoliciesFields.PolicyStatus():
		return "refund_policies_policy_status_enum"

	case selectRefundPoliciesFields.Metadata():
		return "jsonb"

	case selectRefundPoliciesFields.MetaCreatedAt():
		return "timestamptz"

	case selectRefundPoliciesFields.MetaCreatedBy():
		return "uuid"

	case selectRefundPoliciesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRefundPoliciesFields.MetaUpdatedBy():
		return "uuid"

	case selectRefundPoliciesFields.MetaDeletedAt():
		return "timestamptz"

	case selectRefundPoliciesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRefundPolicies(ctx context.Context, refundPolicies *model.RefundPolicies, fieldsInsert ...RefundPoliciesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRefundPoliciesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RefundPoliciesPrimaryID{
		Id: refundPolicies.Id,
	}
	exists, err := repo.IsExistRefundPoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundPolicies] failed checking refundPolicies whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "refundPolicies", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRefundPolicies([]model.RefundPolicies{*refundPolicies}, fieldsInsert...)
	commandQuery := fmt.Sprintf(refundPoliciesQueries.insertRefundPolicies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundPolicies] failed exec create refundPolicies query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRefundPoliciesByID(ctx context.Context, primaryID model.RefundPoliciesPrimaryID) (err error) {
	exists, err := repo.IsExistRefundPoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundPoliciesByID] failed checking refundPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundPolicies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRefundPoliciesCompositePrimaryKeyWhere([]model.RefundPoliciesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(refundPoliciesQueries.deleteRefundPolicies + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundPoliciesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundPoliciesByFilter(ctx context.Context, filter model.Filter) (result []model.RefundPoliciesFilterResult, err error) {
	query, args, err := composeRefundPoliciesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundPoliciesByFilter] failed compose refundPolicies filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundPoliciesByFilter] failed get refundPolicies by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRefundPoliciesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.RefundPoliciesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeRefundPoliciesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeRefundPoliciesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeRefundPoliciesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 17 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewRefundPoliciesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["policy_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"policy_code\"")
			selectedColumns["policy_code"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["policy_scope"]; !selected {
			selectColumns = append(selectColumns, "base.\"policy_scope\"")
			selectedColumns["policy_scope"] = struct{}{}
		}
		if _, selected := selectedColumns["allow_partial"]; !selected {
			selectColumns = append(selectColumns, "base.\"allow_partial\"")
			selectedColumns["allow_partial"] = struct{}{}
		}
		if _, selected := selectedColumns["allow_post_payout"]; !selected {
			selectColumns = append(selectColumns, "base.\"allow_post_payout\"")
			selectedColumns["allow_post_payout"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_return_mode"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_return_mode\"")
			selectedColumns["fee_return_mode"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_return_mode"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_return_mode\"")
			selectedColumns["tax_return_mode"] = struct{}{}
		}
		if _, selected := selectedColumns["requires_approval_over_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"requires_approval_over_amount\"")
			selectedColumns["requires_approval_over_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["policy_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"policy_status\"")
			selectedColumns["policy_status"] = struct{}{}
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

type refundPoliciesFilterPlaceholder struct {
	index int
}

func (p *refundPoliciesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeRefundPoliciesFilterPredicate(filterField model.FilterField, placeholders *refundPoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewRefundPoliciesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeRefundPoliciesFilterSQLExpr(spec)
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

func composeRefundPoliciesFilterGroup(group model.FilterGroup, placeholders *refundPoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeRefundPoliciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeRefundPoliciesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeRefundPoliciesFilterWhereQueries(filter model.Filter, placeholders *refundPoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeRefundPoliciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeRefundPoliciesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeRefundPoliciesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRefundPoliciesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeRefundPoliciesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeRefundPoliciesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := refundPoliciesFilterPlaceholder{index: 1}
	whereQueries, err := composeRefundPoliciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewRefundPoliciesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeRefundPoliciesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeRefundPoliciesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"refund_policies\" base%s", strings.Join(selectColumns, ","), composeRefundPoliciesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistRefundPoliciesByID(ctx context.Context, primaryID model.RefundPoliciesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRefundPoliciesCompositePrimaryKeyWhere([]model.RefundPoliciesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", refundPoliciesQueries.selectCountRefundPolicies, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundPoliciesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundPolicies(ctx context.Context, selectFields ...RefundPoliciesField) (refundPoliciesList model.RefundPoliciesList, err error) {
	var (
		defaultRefundPoliciesSelectFields = defaultRefundPoliciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundPoliciesSelectFields = composeRefundPoliciesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(refundPoliciesQueries.selectRefundPolicies, defaultRefundPoliciesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &refundPoliciesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundPolicies] failed get refundPolicies list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundPoliciesByID(ctx context.Context, primaryID model.RefundPoliciesPrimaryID, selectFields ...RefundPoliciesField) (refundPolicies model.RefundPolicies, err error) {
	var (
		defaultRefundPoliciesSelectFields = defaultRefundPoliciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundPoliciesSelectFields = composeRefundPoliciesSelectFields(selectFields...)
	}
	whereQry, params := composeRefundPoliciesCompositePrimaryKeyWhere([]model.RefundPoliciesPrimaryID{primaryID})
	query := fmt.Sprintf(refundPoliciesQueries.selectRefundPolicies+" WHERE "+whereQry, defaultRefundPoliciesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &refundPolicies, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("refundPolicies with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRefundPoliciesByID] failed get refundPolicies")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRefundPoliciesByID(ctx context.Context, primaryID model.RefundPoliciesPrimaryID, refundPolicies *model.RefundPolicies, refundPoliciesUpdateFields ...RefundPoliciesUpdateField) (err error) {
	exists, err := repo.IsExistRefundPoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundPolicies] failed checking refundPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundPolicies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if refundPolicies == nil {
		if len(refundPoliciesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRefundPoliciesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		refundPolicies = &model.RefundPolicies{}
	}
	var (
		defaultRefundPoliciesUpdateFields = defaultRefundPoliciesUpdateFields(*refundPolicies)
		tempUpdateField                   RefundPoliciesUpdateFieldList
		selectFields                      = NewRefundPoliciesSelectFields()
	)
	if len(refundPoliciesUpdateFields) > 0 {
		for _, updateField := range refundPoliciesUpdateFields {
			if updateField.refundPoliciesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRefundPoliciesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRefundPoliciesCompositePrimaryKeyWhere([]model.RefundPoliciesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRefundPoliciesCommand(defaultRefundPoliciesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(refundPoliciesQueries.updateRefundPolicies+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundPolicies] error when try to update refundPolicies by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateRefundPoliciesByFilter(ctx context.Context, filter model.Filter, refundPoliciesUpdateFields ...RefundPoliciesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(refundPoliciesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields RefundPoliciesUpdateFieldList
		selectFields = NewRefundPoliciesSelectFields()
	)
	for _, updateField := range refundPoliciesUpdateFields {
		if updateField.refundPoliciesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsRefundPoliciesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := refundPoliciesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeRefundPoliciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"refund_policies\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundPoliciesByFilter] error when try to update refundPolicies by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundPoliciesByFilter] failed get rows affected")
	}
	return
}

var (
	refundPoliciesQueries = struct {
		selectRefundPolicies      string
		selectCountRefundPolicies string
		deleteRefundPolicies      string
		updateRefundPolicies      string
		insertRefundPolicies      string
	}{
		selectRefundPolicies:      "SELECT %s FROM \"refund_policies\"",
		selectCountRefundPolicies: "SELECT COUNT(\"id\") FROM \"refund_policies\"",
		deleteRefundPolicies:      "DELETE FROM \"refund_policies\"",
		updateRefundPolicies:      "UPDATE \"refund_policies\" SET %s ",
		insertRefundPolicies:      "INSERT INTO \"refund_policies\" %s VALUES %s",
	}
)

type RefundPoliciesRepository interface {
	CreateRefundPolicies(ctx context.Context, refundPolicies *model.RefundPolicies, fieldsInsert ...RefundPoliciesField) error
	BulkCreateRefundPolicies(ctx context.Context, refundPoliciesList []*model.RefundPolicies, fieldsInsert ...RefundPoliciesField) error
	ResolveRefundPolicies(ctx context.Context, selectFields ...RefundPoliciesField) (model.RefundPoliciesList, error)
	ResolveRefundPoliciesByID(ctx context.Context, primaryID model.RefundPoliciesPrimaryID, selectFields ...RefundPoliciesField) (model.RefundPolicies, error)
	UpdateRefundPoliciesByID(ctx context.Context, id model.RefundPoliciesPrimaryID, refundPolicies *model.RefundPolicies, refundPoliciesUpdateFields ...RefundPoliciesUpdateField) error
	UpdateRefundPoliciesByFilter(ctx context.Context, filter model.Filter, refundPoliciesUpdateFields ...RefundPoliciesUpdateField) (rowsAffected int64, err error)
	BulkUpdateRefundPolicies(ctx context.Context, refundPoliciesListMap map[model.RefundPoliciesPrimaryID]*model.RefundPolicies, RefundPoliciessMapUpdateFieldsRequest map[model.RefundPoliciesPrimaryID]RefundPoliciesUpdateFieldList) (err error)
	DeleteRefundPoliciesByID(ctx context.Context, id model.RefundPoliciesPrimaryID) error
	BulkDeleteRefundPoliciesByIDs(ctx context.Context, ids []model.RefundPoliciesPrimaryID) error
	ResolveRefundPoliciesByFilter(ctx context.Context, filter model.Filter) (result []model.RefundPoliciesFilterResult, err error)
	IsExistRefundPoliciesByIDs(ctx context.Context, ids []model.RefundPoliciesPrimaryID) (exists bool, notFoundIds []model.RefundPoliciesPrimaryID, err error)
	IsExistRefundPoliciesByID(ctx context.Context, id model.RefundPoliciesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
