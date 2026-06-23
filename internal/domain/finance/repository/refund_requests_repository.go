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

func composeInsertFieldsAndParamsRefundRequests(refundRequestsList []model.RefundRequests, fieldsInsert ...RefundRequestsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRefundRequestsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, refundRequests := range refundRequestsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, refundRequests.Id)
			case selectField.RefundCode():
				args = append(args, refundRequests.RefundCode)
			case selectField.PaymentRefId():
				args = append(args, refundRequests.PaymentRefId)
			case selectField.MerchantPartyId():
				args = append(args, refundRequests.MerchantPartyId)
			case selectField.CustomerPartyId():
				args = append(args, refundRequests.CustomerPartyId)
			case selectField.RefundPolicyId():
				args = append(args, refundRequests.RefundPolicyId)
			case selectField.CurrencyCode():
				args = append(args, refundRequests.CurrencyCode)
			case selectField.RequestedAmount():
				args = append(args, refundRequests.RequestedAmount)
			case selectField.ApprovedAmount():
				args = append(args, refundRequests.ApprovedAmount)
			case selectField.RefundReasonCode():
				args = append(args, refundRequests.RefundReasonCode)
			case selectField.RefundStatus():
				args = append(args, refundRequests.RefundStatus)
			case selectField.RequestedAt():
				args = append(args, refundRequests.RequestedAt)
			case selectField.ApprovedAt():
				args = append(args, refundRequests.ApprovedAt)
			case selectField.SettledFinanciallyAt():
				args = append(args, refundRequests.SettledFinanciallyAt)
			case selectField.Metadata():
				args = append(args, refundRequests.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, refundRequests.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, refundRequests.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, refundRequests.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, refundRequests.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, refundRequests.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, refundRequests.MetaDeletedBy)

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

func composeRefundRequestsCompositePrimaryKeyWhere(primaryIDs []model.RefundRequestsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"refund_requests\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRefundRequestsSelectFields() string {
	fields := NewRefundRequestsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRefundRequestsSelectFields(selectFields ...RefundRequestsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RefundRequestsField string
type RefundRequestsFieldList []RefundRequestsField

type RefundRequestsSelectFields struct {
}

func (ss RefundRequestsSelectFields) Id() RefundRequestsField {
	return RefundRequestsField("id")
}

func (ss RefundRequestsSelectFields) RefundCode() RefundRequestsField {
	return RefundRequestsField("refund_code")
}

func (ss RefundRequestsSelectFields) PaymentRefId() RefundRequestsField {
	return RefundRequestsField("payment_ref_id")
}

func (ss RefundRequestsSelectFields) MerchantPartyId() RefundRequestsField {
	return RefundRequestsField("merchant_party_id")
}

func (ss RefundRequestsSelectFields) CustomerPartyId() RefundRequestsField {
	return RefundRequestsField("customer_party_id")
}

func (ss RefundRequestsSelectFields) RefundPolicyId() RefundRequestsField {
	return RefundRequestsField("refund_policy_id")
}

func (ss RefundRequestsSelectFields) CurrencyCode() RefundRequestsField {
	return RefundRequestsField("currency_code")
}

func (ss RefundRequestsSelectFields) RequestedAmount() RefundRequestsField {
	return RefundRequestsField("requested_amount")
}

func (ss RefundRequestsSelectFields) ApprovedAmount() RefundRequestsField {
	return RefundRequestsField("approved_amount")
}

func (ss RefundRequestsSelectFields) RefundReasonCode() RefundRequestsField {
	return RefundRequestsField("refund_reason_code")
}

func (ss RefundRequestsSelectFields) RefundStatus() RefundRequestsField {
	return RefundRequestsField("refund_status")
}

func (ss RefundRequestsSelectFields) RequestedAt() RefundRequestsField {
	return RefundRequestsField("requested_at")
}

func (ss RefundRequestsSelectFields) ApprovedAt() RefundRequestsField {
	return RefundRequestsField("approved_at")
}

func (ss RefundRequestsSelectFields) SettledFinanciallyAt() RefundRequestsField {
	return RefundRequestsField("settled_financially_at")
}

func (ss RefundRequestsSelectFields) Metadata() RefundRequestsField {
	return RefundRequestsField("metadata")
}

func (ss RefundRequestsSelectFields) MetaCreatedAt() RefundRequestsField {
	return RefundRequestsField("meta_created_at")
}

func (ss RefundRequestsSelectFields) MetaCreatedBy() RefundRequestsField {
	return RefundRequestsField("meta_created_by")
}

func (ss RefundRequestsSelectFields) MetaUpdatedAt() RefundRequestsField {
	return RefundRequestsField("meta_updated_at")
}

func (ss RefundRequestsSelectFields) MetaUpdatedBy() RefundRequestsField {
	return RefundRequestsField("meta_updated_by")
}

func (ss RefundRequestsSelectFields) MetaDeletedAt() RefundRequestsField {
	return RefundRequestsField("meta_deleted_at")
}

func (ss RefundRequestsSelectFields) MetaDeletedBy() RefundRequestsField {
	return RefundRequestsField("meta_deleted_by")
}

func (ss RefundRequestsSelectFields) All() RefundRequestsFieldList {
	return []RefundRequestsField{
		ss.Id(),
		ss.RefundCode(),
		ss.PaymentRefId(),
		ss.MerchantPartyId(),
		ss.CustomerPartyId(),
		ss.RefundPolicyId(),
		ss.CurrencyCode(),
		ss.RequestedAmount(),
		ss.ApprovedAmount(),
		ss.RefundReasonCode(),
		ss.RefundStatus(),
		ss.RequestedAt(),
		ss.ApprovedAt(),
		ss.SettledFinanciallyAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRefundRequestsSelectFields() RefundRequestsSelectFields {
	return RefundRequestsSelectFields{}
}

type RefundRequestsUpdateFieldOption struct {
	useIncrement bool
}
type RefundRequestsUpdateField struct {
	refundRequestsField RefundRequestsField
	opt                 RefundRequestsUpdateFieldOption
	value               interface{}
}
type RefundRequestsUpdateFieldList []RefundRequestsUpdateField

func defaultRefundRequestsUpdateFieldOption() RefundRequestsUpdateFieldOption {
	return RefundRequestsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRefundRequestsOption(useIncrement bool) func(*RefundRequestsUpdateFieldOption) {
	return func(pcufo *RefundRequestsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRefundRequestsUpdateField(field RefundRequestsField, val interface{}, opts ...func(*RefundRequestsUpdateFieldOption)) RefundRequestsUpdateField {
	defaultOpt := defaultRefundRequestsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RefundRequestsUpdateField{
		refundRequestsField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultRefundRequestsUpdateFields(refundRequests model.RefundRequests) (refundRequestsUpdateFieldList RefundRequestsUpdateFieldList) {
	selectFields := NewRefundRequestsSelectFields()
	refundRequestsUpdateFieldList = append(refundRequestsUpdateFieldList,
		NewRefundRequestsUpdateField(selectFields.Id(), refundRequests.Id),
		NewRefundRequestsUpdateField(selectFields.RefundCode(), refundRequests.RefundCode),
		NewRefundRequestsUpdateField(selectFields.PaymentRefId(), refundRequests.PaymentRefId),
		NewRefundRequestsUpdateField(selectFields.MerchantPartyId(), refundRequests.MerchantPartyId),
		NewRefundRequestsUpdateField(selectFields.CustomerPartyId(), refundRequests.CustomerPartyId),
		NewRefundRequestsUpdateField(selectFields.RefundPolicyId(), refundRequests.RefundPolicyId),
		NewRefundRequestsUpdateField(selectFields.CurrencyCode(), refundRequests.CurrencyCode),
		NewRefundRequestsUpdateField(selectFields.RequestedAmount(), refundRequests.RequestedAmount),
		NewRefundRequestsUpdateField(selectFields.ApprovedAmount(), refundRequests.ApprovedAmount),
		NewRefundRequestsUpdateField(selectFields.RefundReasonCode(), refundRequests.RefundReasonCode),
		NewRefundRequestsUpdateField(selectFields.RefundStatus(), refundRequests.RefundStatus),
		NewRefundRequestsUpdateField(selectFields.RequestedAt(), refundRequests.RequestedAt),
		NewRefundRequestsUpdateField(selectFields.ApprovedAt(), refundRequests.ApprovedAt),
		NewRefundRequestsUpdateField(selectFields.SettledFinanciallyAt(), refundRequests.SettledFinanciallyAt),
		NewRefundRequestsUpdateField(selectFields.Metadata(), refundRequests.Metadata),
		NewRefundRequestsUpdateField(selectFields.MetaCreatedAt(), refundRequests.MetaCreatedAt),
		NewRefundRequestsUpdateField(selectFields.MetaCreatedBy(), refundRequests.MetaCreatedBy),
		NewRefundRequestsUpdateField(selectFields.MetaUpdatedAt(), refundRequests.MetaUpdatedAt),
		NewRefundRequestsUpdateField(selectFields.MetaUpdatedBy(), refundRequests.MetaUpdatedBy),
		NewRefundRequestsUpdateField(selectFields.MetaDeletedAt(), refundRequests.MetaDeletedAt),
		NewRefundRequestsUpdateField(selectFields.MetaDeletedBy(), refundRequests.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRefundRequestsCommand(refundRequestsUpdateFieldList RefundRequestsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range refundRequestsUpdateFieldList {
		field := string(updateField.refundRequestsField)
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

func (repo *RepositoryImpl) BulkCreateRefundRequests(ctx context.Context, refundRequestsList []*model.RefundRequests, fieldsInsert ...RefundRequestsField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.RefundRequestsPrimaryID
		refundRequestsValueList []model.RefundRequests
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRefundRequestsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, refundRequests := range refundRequestsList {

		primaryIds = append(primaryIds, refundRequests.ToRefundRequestsPrimaryID())

		refundRequestsValueList = append(refundRequestsValueList, *refundRequests)
	}

	_, notFoundIds, err := repo.IsExistRefundRequestsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundRequests] failed checking refundRequests whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RefundRequestsPrimaryID{}
		mapNotFoundIds := map[model.RefundRequestsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "refundRequests", fmt.Sprintf("refundRequests with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRefundRequests(refundRequestsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(refundRequestsQueries.insertRefundRequests, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundRequests] failed exec create refundRequests query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRefundRequestsByIDs(ctx context.Context, primaryIDs []model.RefundRequestsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRefundRequestsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundRequestsByIDs] failed checking refundRequests whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundRequests with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_requests\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(refundRequestsQueries.deleteRefundRequests + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundRequestsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundRequestsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRefundRequestsByIDs(ctx context.Context, ids []model.RefundRequestsPrimaryID) (exists bool, notFoundIds []model.RefundRequestsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_requests\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(refundRequestsQueries.selectRefundRequests, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundRequestsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RefundRequestsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundRequestsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RefundRequestsPrimaryID]bool{}
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

// BulkUpdateRefundRequests is used to bulk update refundRequests, by default it will update all field
// if want to update specific field, then fill refundRequestssMapUpdateFieldsRequest else please fill refundRequestssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRefundRequests(ctx context.Context, refundRequestssMap map[model.RefundRequestsPrimaryID]*model.RefundRequests, refundRequestssMapUpdateFieldsRequest map[model.RefundRequestsPrimaryID]RefundRequestsUpdateFieldList) (err error) {
	if len(refundRequestssMap) == 0 && len(refundRequestssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		refundRequestssMapUpdateField map[model.RefundRequestsPrimaryID]RefundRequestsUpdateFieldList = map[model.RefundRequestsPrimaryID]RefundRequestsUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(refundRequestssMap) > 0 {
		for id, refundRequests := range refundRequestssMap {
			if refundRequests == nil {
				log.Error().Err(err).Msg("[BulkUpdateRefundRequests] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			refundRequestssMapUpdateField[id] = defaultRefundRequestsUpdateFields(*refundRequests)
		}
	} else {
		refundRequestssMapUpdateField = refundRequestssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRefundRequestsQuery(refundRequestssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRefundRequestsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundRequests] failed checking refundRequests whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundRequests with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRefundRequestsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"refund_requests\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundRequests] failed exec query")
	}
	return
}

type RefundRequestsFieldParameter struct {
	param string
	args  []interface{}
}

func NewRefundRequestsFieldParameter(param string, args ...interface{}) RefundRequestsFieldParameter {
	return RefundRequestsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRefundRequestsQuery(mapRefundRequestss map[model.RefundRequestsPrimaryID]RefundRequestsUpdateFieldList, asTableValues string) (primaryIDs []model.RefundRequestsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RefundRequestsPrimaryID]map[string]interface{}{}
	refundRequestsSelectFields := NewRefundRequestsSelectFields()
	for id, updateFields := range mapRefundRequestss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.refundRequestsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRefundRequestss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRefundRequestsFieldType(updateField.refundRequestsField)))
			args = append(args, fields[string(updateField.refundRequestsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.refundRequestsField))
		if updateField.refundRequestsField == refundRequestsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.refundRequestsField, asTableValues, updateField.refundRequestsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.refundRequestsField,
				"\"refund_requests\"", updateField.refundRequestsField,
				asTableValues, updateField.refundRequestsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRefundRequestsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RefundRequestsPrimaryID, asTableValue string) (whereQry string) {
	refundRequestsSelectFields := NewRefundRequestsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"refund_requests\".\"id\" = %s.\"id\"::"+GetRefundRequestsFieldType(refundRequestsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRefundRequestsFieldType(refundRequestsField RefundRequestsField) string {
	selectRefundRequestsFields := NewRefundRequestsSelectFields()
	switch refundRequestsField {

	case selectRefundRequestsFields.Id():
		return "uuid"

	case selectRefundRequestsFields.RefundCode():
		return "text"

	case selectRefundRequestsFields.PaymentRefId():
		return "uuid"

	case selectRefundRequestsFields.MerchantPartyId():
		return "uuid"

	case selectRefundRequestsFields.CustomerPartyId():
		return "uuid"

	case selectRefundRequestsFields.RefundPolicyId():
		return "uuid"

	case selectRefundRequestsFields.CurrencyCode():
		return "text"

	case selectRefundRequestsFields.RequestedAmount():
		return "numeric"

	case selectRefundRequestsFields.ApprovedAmount():
		return "numeric"

	case selectRefundRequestsFields.RefundReasonCode():
		return "text"

	case selectRefundRequestsFields.RefundStatus():
		return "refund_status_enum"

	case selectRefundRequestsFields.RequestedAt():
		return "timestamptz"

	case selectRefundRequestsFields.ApprovedAt():
		return "timestamptz"

	case selectRefundRequestsFields.SettledFinanciallyAt():
		return "timestamptz"

	case selectRefundRequestsFields.Metadata():
		return "jsonb"

	case selectRefundRequestsFields.MetaCreatedAt():
		return "timestamptz"

	case selectRefundRequestsFields.MetaCreatedBy():
		return "uuid"

	case selectRefundRequestsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRefundRequestsFields.MetaUpdatedBy():
		return "uuid"

	case selectRefundRequestsFields.MetaDeletedAt():
		return "timestamptz"

	case selectRefundRequestsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRefundRequests(ctx context.Context, refundRequests *model.RefundRequests, fieldsInsert ...RefundRequestsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRefundRequestsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RefundRequestsPrimaryID{
		Id: refundRequests.Id,
	}
	exists, err := repo.IsExistRefundRequestsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundRequests] failed checking refundRequests whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "refundRequests", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRefundRequests([]model.RefundRequests{*refundRequests}, fieldsInsert...)
	commandQuery := fmt.Sprintf(refundRequestsQueries.insertRefundRequests, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundRequests] failed exec create refundRequests query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRefundRequestsByID(ctx context.Context, primaryID model.RefundRequestsPrimaryID) (err error) {
	exists, err := repo.IsExistRefundRequestsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundRequestsByID] failed checking refundRequests whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundRequests with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRefundRequestsCompositePrimaryKeyWhere([]model.RefundRequestsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(refundRequestsQueries.deleteRefundRequests + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundRequestsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundRequestsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundRequestsFilterResult, err error) {
	query, args, err := composeRefundRequestsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundRequestsByFilter] failed compose refundRequests filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundRequestsByFilter] failed get refundRequests by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRefundRequestsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.RefundRequestsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeRefundRequestsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeRefundRequestsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeRefundRequestsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 21 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewRefundRequestsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 21+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["refund_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_code\"")
			selectedColumns["refund_code"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_ref_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_ref_id\"")
			selectedColumns["payment_ref_id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["customer_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"customer_party_id\"")
			selectedColumns["customer_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["refund_policy_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_policy_id\"")
			selectedColumns["refund_policy_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["requested_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"requested_amount\"")
			selectedColumns["requested_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_amount\"")
			selectedColumns["approved_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["refund_reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_reason_code\"")
			selectedColumns["refund_reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["refund_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_status\"")
			selectedColumns["refund_status"] = struct{}{}
		}
		if _, selected := selectedColumns["requested_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"requested_at\"")
			selectedColumns["requested_at"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_at\"")
			selectedColumns["approved_at"] = struct{}{}
		}
		if _, selected := selectedColumns["settled_financially_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"settled_financially_at\"")
			selectedColumns["settled_financially_at"] = struct{}{}
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

type refundRequestsFilterPlaceholder struct {
	index int
}

func (p *refundRequestsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeRefundRequestsFilterPredicate(filterField model.FilterField, placeholders *refundRequestsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewRefundRequestsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeRefundRequestsFilterSQLExpr(spec)
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

func composeRefundRequestsFilterGroup(group model.FilterGroup, placeholders *refundRequestsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeRefundRequestsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeRefundRequestsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeRefundRequestsFilterWhereQueries(filter model.Filter, placeholders *refundRequestsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeRefundRequestsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeRefundRequestsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeRefundRequestsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRefundRequestsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeRefundRequestsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeRefundRequestsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := refundRequestsFilterPlaceholder{index: 1}
	whereQueries, err := composeRefundRequestsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewRefundRequestsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeRefundRequestsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeRefundRequestsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"refund_requests\" base%s", strings.Join(selectColumns, ","), composeRefundRequestsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistRefundRequestsByID(ctx context.Context, primaryID model.RefundRequestsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRefundRequestsCompositePrimaryKeyWhere([]model.RefundRequestsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", refundRequestsQueries.selectCountRefundRequests, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundRequestsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundRequests(ctx context.Context, selectFields ...RefundRequestsField) (refundRequestsList model.RefundRequestsList, err error) {
	var (
		defaultRefundRequestsSelectFields = defaultRefundRequestsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundRequestsSelectFields = composeRefundRequestsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(refundRequestsQueries.selectRefundRequests, defaultRefundRequestsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &refundRequestsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundRequests] failed get refundRequests list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundRequestsByID(ctx context.Context, primaryID model.RefundRequestsPrimaryID, selectFields ...RefundRequestsField) (refundRequests model.RefundRequests, err error) {
	var (
		defaultRefundRequestsSelectFields = defaultRefundRequestsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundRequestsSelectFields = composeRefundRequestsSelectFields(selectFields...)
	}
	whereQry, params := composeRefundRequestsCompositePrimaryKeyWhere([]model.RefundRequestsPrimaryID{primaryID})
	query := fmt.Sprintf(refundRequestsQueries.selectRefundRequests+" WHERE "+whereQry, defaultRefundRequestsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &refundRequests, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("refundRequests with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRefundRequestsByID] failed get refundRequests")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRefundRequestsByID(ctx context.Context, primaryID model.RefundRequestsPrimaryID, refundRequests *model.RefundRequests, refundRequestsUpdateFields ...RefundRequestsUpdateField) (err error) {
	exists, err := repo.IsExistRefundRequestsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundRequests] failed checking refundRequests whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundRequests with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if refundRequests == nil {
		if len(refundRequestsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRefundRequestsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		refundRequests = &model.RefundRequests{}
	}
	var (
		defaultRefundRequestsUpdateFields = defaultRefundRequestsUpdateFields(*refundRequests)
		tempUpdateField                   RefundRequestsUpdateFieldList
		selectFields                      = NewRefundRequestsSelectFields()
	)
	if len(refundRequestsUpdateFields) > 0 {
		for _, updateField := range refundRequestsUpdateFields {
			if updateField.refundRequestsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRefundRequestsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRefundRequestsCompositePrimaryKeyWhere([]model.RefundRequestsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRefundRequestsCommand(defaultRefundRequestsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(refundRequestsQueries.updateRefundRequests+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundRequests] error when try to update refundRequests by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateRefundRequestsByFilter(ctx context.Context, filter model.Filter, refundRequestsUpdateFields ...RefundRequestsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(refundRequestsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields RefundRequestsUpdateFieldList
		selectFields = NewRefundRequestsSelectFields()
	)
	for _, updateField := range refundRequestsUpdateFields {
		if updateField.refundRequestsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsRefundRequestsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := refundRequestsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeRefundRequestsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"refund_requests\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundRequestsByFilter] error when try to update refundRequests by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundRequestsByFilter] failed get rows affected")
	}
	return
}

var (
	refundRequestsQueries = struct {
		selectRefundRequests      string
		selectCountRefundRequests string
		deleteRefundRequests      string
		updateRefundRequests      string
		insertRefundRequests      string
	}{
		selectRefundRequests:      "SELECT %s FROM \"refund_requests\"",
		selectCountRefundRequests: "SELECT COUNT(\"id\") FROM \"refund_requests\"",
		deleteRefundRequests:      "DELETE FROM \"refund_requests\"",
		updateRefundRequests:      "UPDATE \"refund_requests\" SET %s ",
		insertRefundRequests:      "INSERT INTO \"refund_requests\" %s VALUES %s",
	}
)

type RefundRequestsRepository interface {
	CreateRefundRequests(ctx context.Context, refundRequests *model.RefundRequests, fieldsInsert ...RefundRequestsField) error
	BulkCreateRefundRequests(ctx context.Context, refundRequestsList []*model.RefundRequests, fieldsInsert ...RefundRequestsField) error
	ResolveRefundRequests(ctx context.Context, selectFields ...RefundRequestsField) (model.RefundRequestsList, error)
	ResolveRefundRequestsByID(ctx context.Context, primaryID model.RefundRequestsPrimaryID, selectFields ...RefundRequestsField) (model.RefundRequests, error)
	UpdateRefundRequestsByID(ctx context.Context, id model.RefundRequestsPrimaryID, refundRequests *model.RefundRequests, refundRequestsUpdateFields ...RefundRequestsUpdateField) error
	UpdateRefundRequestsByFilter(ctx context.Context, filter model.Filter, refundRequestsUpdateFields ...RefundRequestsUpdateField) (rowsAffected int64, err error)
	BulkUpdateRefundRequests(ctx context.Context, refundRequestsListMap map[model.RefundRequestsPrimaryID]*model.RefundRequests, RefundRequestssMapUpdateFieldsRequest map[model.RefundRequestsPrimaryID]RefundRequestsUpdateFieldList) (err error)
	DeleteRefundRequestsByID(ctx context.Context, id model.RefundRequestsPrimaryID) error
	BulkDeleteRefundRequestsByIDs(ctx context.Context, ids []model.RefundRequestsPrimaryID) error
	ResolveRefundRequestsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundRequestsFilterResult, err error)
	IsExistRefundRequestsByIDs(ctx context.Context, ids []model.RefundRequestsPrimaryID) (exists bool, notFoundIds []model.RefundRequestsPrimaryID, err error)
	IsExistRefundRequestsByID(ctx context.Context, id model.RefundRequestsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
