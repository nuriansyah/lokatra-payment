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

func composeInsertFieldsAndParamsRefundAttempts(refundAttemptsList []model.RefundAttempts, fieldsInsert ...RefundAttemptsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRefundAttemptsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, refundAttempts := range refundAttemptsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, refundAttempts.Id)
			case selectField.RefundId():
				args = append(args, refundAttempts.RefundId)
			case selectField.AttemptNo():
				args = append(args, refundAttempts.AttemptNo)
			case selectField.AttemptType():
				args = append(args, refundAttempts.AttemptType)
			case selectField.ProviderAccountId():
				args = append(args, refundAttempts.ProviderAccountId)
			case selectField.Amount():
				args = append(args, refundAttempts.Amount)
			case selectField.CurrencyCode():
				args = append(args, refundAttempts.CurrencyCode)
			case selectField.AttemptStatus():
				args = append(args, refundAttempts.AttemptStatus)
			case selectField.ProviderRefundRef():
				args = append(args, refundAttempts.ProviderRefundRef)
			case selectField.FailureCode():
				args = append(args, refundAttempts.FailureCode)
			case selectField.FailureReason():
				args = append(args, refundAttempts.FailureReason)
			case selectField.RawRequest():
				args = append(args, refundAttempts.RawRequest)
			case selectField.RawResponse():
				args = append(args, refundAttempts.RawResponse)
			case selectField.Metadata():
				args = append(args, refundAttempts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, refundAttempts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, refundAttempts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, refundAttempts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, refundAttempts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, refundAttempts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, refundAttempts.MetaDeletedBy)

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

func composeRefundAttemptsCompositePrimaryKeyWhere(primaryIDs []model.RefundAttemptsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"refund_attempts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRefundAttemptsSelectFields() string {
	fields := NewRefundAttemptsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRefundAttemptsSelectFields(selectFields ...RefundAttemptsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RefundAttemptsField string
type RefundAttemptsFieldList []RefundAttemptsField

type RefundAttemptsSelectFields struct {
}

func (ss RefundAttemptsSelectFields) Id() RefundAttemptsField {
	return RefundAttemptsField("id")
}

func (ss RefundAttemptsSelectFields) RefundId() RefundAttemptsField {
	return RefundAttemptsField("refund_id")
}

func (ss RefundAttemptsSelectFields) AttemptNo() RefundAttemptsField {
	return RefundAttemptsField("attempt_no")
}

func (ss RefundAttemptsSelectFields) AttemptType() RefundAttemptsField {
	return RefundAttemptsField("attempt_type")
}

func (ss RefundAttemptsSelectFields) ProviderAccountId() RefundAttemptsField {
	return RefundAttemptsField("provider_account_id")
}

func (ss RefundAttemptsSelectFields) Amount() RefundAttemptsField {
	return RefundAttemptsField("amount")
}

func (ss RefundAttemptsSelectFields) CurrencyCode() RefundAttemptsField {
	return RefundAttemptsField("currency_code")
}

func (ss RefundAttemptsSelectFields) AttemptStatus() RefundAttemptsField {
	return RefundAttemptsField("attempt_status")
}

func (ss RefundAttemptsSelectFields) ProviderRefundRef() RefundAttemptsField {
	return RefundAttemptsField("provider_refund_ref")
}

func (ss RefundAttemptsSelectFields) FailureCode() RefundAttemptsField {
	return RefundAttemptsField("failure_code")
}

func (ss RefundAttemptsSelectFields) FailureReason() RefundAttemptsField {
	return RefundAttemptsField("failure_reason")
}

func (ss RefundAttemptsSelectFields) RawRequest() RefundAttemptsField {
	return RefundAttemptsField("raw_request")
}

func (ss RefundAttemptsSelectFields) RawResponse() RefundAttemptsField {
	return RefundAttemptsField("raw_response")
}

func (ss RefundAttemptsSelectFields) Metadata() RefundAttemptsField {
	return RefundAttemptsField("metadata")
}

func (ss RefundAttemptsSelectFields) MetaCreatedAt() RefundAttemptsField {
	return RefundAttemptsField("meta_created_at")
}

func (ss RefundAttemptsSelectFields) MetaCreatedBy() RefundAttemptsField {
	return RefundAttemptsField("meta_created_by")
}

func (ss RefundAttemptsSelectFields) MetaUpdatedAt() RefundAttemptsField {
	return RefundAttemptsField("meta_updated_at")
}

func (ss RefundAttemptsSelectFields) MetaUpdatedBy() RefundAttemptsField {
	return RefundAttemptsField("meta_updated_by")
}

func (ss RefundAttemptsSelectFields) MetaDeletedAt() RefundAttemptsField {
	return RefundAttemptsField("meta_deleted_at")
}

func (ss RefundAttemptsSelectFields) MetaDeletedBy() RefundAttemptsField {
	return RefundAttemptsField("meta_deleted_by")
}

func (ss RefundAttemptsSelectFields) All() RefundAttemptsFieldList {
	return []RefundAttemptsField{
		ss.Id(),
		ss.RefundId(),
		ss.AttemptNo(),
		ss.AttemptType(),
		ss.ProviderAccountId(),
		ss.Amount(),
		ss.CurrencyCode(),
		ss.AttemptStatus(),
		ss.ProviderRefundRef(),
		ss.FailureCode(),
		ss.FailureReason(),
		ss.RawRequest(),
		ss.RawResponse(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRefundAttemptsSelectFields() RefundAttemptsSelectFields {
	return RefundAttemptsSelectFields{}
}

type RefundAttemptsUpdateFieldOption struct {
	useIncrement bool
}
type RefundAttemptsUpdateField struct {
	refundAttemptsField RefundAttemptsField
	opt                 RefundAttemptsUpdateFieldOption
	value               interface{}
}
type RefundAttemptsUpdateFieldList []RefundAttemptsUpdateField

func defaultRefundAttemptsUpdateFieldOption() RefundAttemptsUpdateFieldOption {
	return RefundAttemptsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRefundAttemptsOption(useIncrement bool) func(*RefundAttemptsUpdateFieldOption) {
	return func(pcufo *RefundAttemptsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRefundAttemptsUpdateField(field RefundAttemptsField, val interface{}, opts ...func(*RefundAttemptsUpdateFieldOption)) RefundAttemptsUpdateField {
	defaultOpt := defaultRefundAttemptsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RefundAttemptsUpdateField{
		refundAttemptsField: field,
		value:               val,
		opt:                 defaultOpt,
	}
}
func defaultRefundAttemptsUpdateFields(refundAttempts model.RefundAttempts) (refundAttemptsUpdateFieldList RefundAttemptsUpdateFieldList) {
	selectFields := NewRefundAttemptsSelectFields()
	refundAttemptsUpdateFieldList = append(refundAttemptsUpdateFieldList,
		NewRefundAttemptsUpdateField(selectFields.Id(), refundAttempts.Id),
		NewRefundAttemptsUpdateField(selectFields.RefundId(), refundAttempts.RefundId),
		NewRefundAttemptsUpdateField(selectFields.AttemptNo(), refundAttempts.AttemptNo),
		NewRefundAttemptsUpdateField(selectFields.AttemptType(), refundAttempts.AttemptType),
		NewRefundAttemptsUpdateField(selectFields.ProviderAccountId(), refundAttempts.ProviderAccountId),
		NewRefundAttemptsUpdateField(selectFields.Amount(), refundAttempts.Amount),
		NewRefundAttemptsUpdateField(selectFields.CurrencyCode(), refundAttempts.CurrencyCode),
		NewRefundAttemptsUpdateField(selectFields.AttemptStatus(), refundAttempts.AttemptStatus),
		NewRefundAttemptsUpdateField(selectFields.ProviderRefundRef(), refundAttempts.ProviderRefundRef),
		NewRefundAttemptsUpdateField(selectFields.FailureCode(), refundAttempts.FailureCode),
		NewRefundAttemptsUpdateField(selectFields.FailureReason(), refundAttempts.FailureReason),
		NewRefundAttemptsUpdateField(selectFields.RawRequest(), refundAttempts.RawRequest),
		NewRefundAttemptsUpdateField(selectFields.RawResponse(), refundAttempts.RawResponse),
		NewRefundAttemptsUpdateField(selectFields.Metadata(), refundAttempts.Metadata),
		NewRefundAttemptsUpdateField(selectFields.MetaCreatedAt(), refundAttempts.MetaCreatedAt),
		NewRefundAttemptsUpdateField(selectFields.MetaCreatedBy(), refundAttempts.MetaCreatedBy),
		NewRefundAttemptsUpdateField(selectFields.MetaUpdatedAt(), refundAttempts.MetaUpdatedAt),
		NewRefundAttemptsUpdateField(selectFields.MetaUpdatedBy(), refundAttempts.MetaUpdatedBy),
		NewRefundAttemptsUpdateField(selectFields.MetaDeletedAt(), refundAttempts.MetaDeletedAt),
		NewRefundAttemptsUpdateField(selectFields.MetaDeletedBy(), refundAttempts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRefundAttemptsCommand(refundAttemptsUpdateFieldList RefundAttemptsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range refundAttemptsUpdateFieldList {
		field := string(updateField.refundAttemptsField)
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

func (repo *RepositoryImpl) BulkCreateRefundAttempts(ctx context.Context, refundAttemptsList []*model.RefundAttempts, fieldsInsert ...RefundAttemptsField) (err error) {
	var (
		fieldsStr               string
		valueListStr            []string
		argsList                []interface{}
		primaryIds              []model.RefundAttemptsPrimaryID
		refundAttemptsValueList []model.RefundAttempts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRefundAttemptsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, refundAttempts := range refundAttemptsList {

		primaryIds = append(primaryIds, refundAttempts.ToRefundAttemptsPrimaryID())

		refundAttemptsValueList = append(refundAttemptsValueList, *refundAttempts)
	}

	_, notFoundIds, err := repo.IsExistRefundAttemptsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundAttempts] failed checking refundAttempts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RefundAttemptsPrimaryID{}
		mapNotFoundIds := map[model.RefundAttemptsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "refundAttempts", fmt.Sprintf("refundAttempts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRefundAttempts(refundAttemptsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(refundAttemptsQueries.insertRefundAttempts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundAttempts] failed exec create refundAttempts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRefundAttemptsByIDs(ctx context.Context, primaryIDs []model.RefundAttemptsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRefundAttemptsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundAttemptsByIDs] failed checking refundAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAttempts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_attempts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(refundAttemptsQueries.deleteRefundAttempts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundAttemptsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundAttemptsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRefundAttemptsByIDs(ctx context.Context, ids []model.RefundAttemptsPrimaryID) (exists bool, notFoundIds []model.RefundAttemptsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_attempts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(refundAttemptsQueries.selectRefundAttempts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundAttemptsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RefundAttemptsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundAttemptsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RefundAttemptsPrimaryID]bool{}
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

// BulkUpdateRefundAttempts is used to bulk update refundAttempts, by default it will update all field
// if want to update specific field, then fill refundAttemptssMapUpdateFieldsRequest else please fill refundAttemptssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRefundAttempts(ctx context.Context, refundAttemptssMap map[model.RefundAttemptsPrimaryID]*model.RefundAttempts, refundAttemptssMapUpdateFieldsRequest map[model.RefundAttemptsPrimaryID]RefundAttemptsUpdateFieldList) (err error) {
	if len(refundAttemptssMap) == 0 && len(refundAttemptssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		refundAttemptssMapUpdateField map[model.RefundAttemptsPrimaryID]RefundAttemptsUpdateFieldList = map[model.RefundAttemptsPrimaryID]RefundAttemptsUpdateFieldList{}
		asTableValues                 string                                                          = "myvalues"
	)

	if len(refundAttemptssMap) > 0 {
		for id, refundAttempts := range refundAttemptssMap {
			if refundAttempts == nil {
				log.Error().Err(err).Msg("[BulkUpdateRefundAttempts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			refundAttemptssMapUpdateField[id] = defaultRefundAttemptsUpdateFields(*refundAttempts)
		}
	} else {
		refundAttemptssMapUpdateField = refundAttemptssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRefundAttemptsQuery(refundAttemptssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRefundAttemptsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundAttempts] failed checking refundAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAttempts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRefundAttemptsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"refund_attempts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundAttempts] failed exec query")
	}
	return
}

type RefundAttemptsFieldParameter struct {
	param string
	args  []interface{}
}

func NewRefundAttemptsFieldParameter(param string, args ...interface{}) RefundAttemptsFieldParameter {
	return RefundAttemptsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRefundAttemptsQuery(mapRefundAttemptss map[model.RefundAttemptsPrimaryID]RefundAttemptsUpdateFieldList, asTableValues string) (primaryIDs []model.RefundAttemptsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RefundAttemptsPrimaryID]map[string]interface{}{}
	refundAttemptsSelectFields := NewRefundAttemptsSelectFields()
	for id, updateFields := range mapRefundAttemptss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.refundAttemptsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRefundAttemptss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRefundAttemptsFieldType(updateField.refundAttemptsField)))
			args = append(args, fields[string(updateField.refundAttemptsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.refundAttemptsField))
		if updateField.refundAttemptsField == refundAttemptsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.refundAttemptsField, asTableValues, updateField.refundAttemptsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.refundAttemptsField,
				"\"refund_attempts\"", updateField.refundAttemptsField,
				asTableValues, updateField.refundAttemptsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRefundAttemptsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RefundAttemptsPrimaryID, asTableValue string) (whereQry string) {
	refundAttemptsSelectFields := NewRefundAttemptsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"refund_attempts\".\"id\" = %s.\"id\"::"+GetRefundAttemptsFieldType(refundAttemptsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRefundAttemptsFieldType(refundAttemptsField RefundAttemptsField) string {
	selectRefundAttemptsFields := NewRefundAttemptsSelectFields()
	switch refundAttemptsField {

	case selectRefundAttemptsFields.Id():
		return "uuid"

	case selectRefundAttemptsFields.RefundId():
		return "uuid"

	case selectRefundAttemptsFields.AttemptNo():
		return "int4"

	case selectRefundAttemptsFields.AttemptType():
		return "refund_attempts_attempt_type_enum"

	case selectRefundAttemptsFields.ProviderAccountId():
		return "uuid"

	case selectRefundAttemptsFields.Amount():
		return "numeric"

	case selectRefundAttemptsFields.CurrencyCode():
		return "text"

	case selectRefundAttemptsFields.AttemptStatus():
		return "refund_attempts_attempt_status_enum"

	case selectRefundAttemptsFields.ProviderRefundRef():
		return "text"

	case selectRefundAttemptsFields.FailureCode():
		return "text"

	case selectRefundAttemptsFields.FailureReason():
		return "text"

	case selectRefundAttemptsFields.RawRequest():
		return "jsonb"

	case selectRefundAttemptsFields.RawResponse():
		return "jsonb"

	case selectRefundAttemptsFields.Metadata():
		return "jsonb"

	case selectRefundAttemptsFields.MetaCreatedAt():
		return "timestamptz"

	case selectRefundAttemptsFields.MetaCreatedBy():
		return "uuid"

	case selectRefundAttemptsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRefundAttemptsFields.MetaUpdatedBy():
		return "uuid"

	case selectRefundAttemptsFields.MetaDeletedAt():
		return "timestamptz"

	case selectRefundAttemptsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRefundAttempts(ctx context.Context, refundAttempts *model.RefundAttempts, fieldsInsert ...RefundAttemptsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRefundAttemptsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RefundAttemptsPrimaryID{
		Id: refundAttempts.Id,
	}
	exists, err := repo.IsExistRefundAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundAttempts] failed checking refundAttempts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "refundAttempts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRefundAttempts([]model.RefundAttempts{*refundAttempts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(refundAttemptsQueries.insertRefundAttempts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundAttempts] failed exec create refundAttempts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRefundAttemptsByID(ctx context.Context, primaryID model.RefundAttemptsPrimaryID) (err error) {
	exists, err := repo.IsExistRefundAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundAttemptsByID] failed checking refundAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAttempts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRefundAttemptsCompositePrimaryKeyWhere([]model.RefundAttemptsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(refundAttemptsQueries.deleteRefundAttempts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundAttemptsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundAttemptsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundAttemptsFilterResult, err error) {
	query, args, err := composeRefundAttemptsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundAttemptsByFilter] failed compose refundAttempts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundAttemptsByFilter] failed get refundAttempts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRefundAttemptsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.RefundAttemptsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeRefundAttemptsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeRefundAttemptsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeRefundAttemptsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewRefundAttemptsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 20+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["refund_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_id\"")
			selectedColumns["refund_id"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_no\"")
			selectedColumns["attempt_no"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_type\"")
			selectedColumns["attempt_type"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["attempt_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"attempt_status\"")
			selectedColumns["attempt_status"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_refund_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_refund_ref\"")
			selectedColumns["provider_refund_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_code\"")
			selectedColumns["failure_code"] = struct{}{}
		}
		if _, selected := selectedColumns["failure_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"failure_reason\"")
			selectedColumns["failure_reason"] = struct{}{}
		}
		if _, selected := selectedColumns["raw_request"]; !selected {
			selectColumns = append(selectColumns, "base.\"raw_request\"")
			selectedColumns["raw_request"] = struct{}{}
		}
		if _, selected := selectedColumns["raw_response"]; !selected {
			selectColumns = append(selectColumns, "base.\"raw_response\"")
			selectedColumns["raw_response"] = struct{}{}
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

type refundAttemptsFilterPlaceholder struct {
	index int
}

func (p *refundAttemptsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeRefundAttemptsFilterPredicate(filterField model.FilterField, placeholders *refundAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewRefundAttemptsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeRefundAttemptsFilterSQLExpr(spec)
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

func composeRefundAttemptsFilterGroup(group model.FilterGroup, placeholders *refundAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeRefundAttemptsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeRefundAttemptsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeRefundAttemptsFilterWhereQueries(filter model.Filter, placeholders *refundAttemptsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeRefundAttemptsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeRefundAttemptsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeRefundAttemptsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRefundAttemptsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeRefundAttemptsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeRefundAttemptsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := refundAttemptsFilterPlaceholder{index: 1}
	whereQueries, err := composeRefundAttemptsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewRefundAttemptsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeRefundAttemptsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeRefundAttemptsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"refund_attempts\" base%s", strings.Join(selectColumns, ","), composeRefundAttemptsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistRefundAttemptsByID(ctx context.Context, primaryID model.RefundAttemptsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRefundAttemptsCompositePrimaryKeyWhere([]model.RefundAttemptsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", refundAttemptsQueries.selectCountRefundAttempts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundAttemptsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundAttempts(ctx context.Context, selectFields ...RefundAttemptsField) (refundAttemptsList model.RefundAttemptsList, err error) {
	var (
		defaultRefundAttemptsSelectFields = defaultRefundAttemptsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundAttemptsSelectFields = composeRefundAttemptsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(refundAttemptsQueries.selectRefundAttempts, defaultRefundAttemptsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &refundAttemptsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundAttempts] failed get refundAttempts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundAttemptsByID(ctx context.Context, primaryID model.RefundAttemptsPrimaryID, selectFields ...RefundAttemptsField) (refundAttempts model.RefundAttempts, err error) {
	var (
		defaultRefundAttemptsSelectFields = defaultRefundAttemptsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundAttemptsSelectFields = composeRefundAttemptsSelectFields(selectFields...)
	}
	whereQry, params := composeRefundAttemptsCompositePrimaryKeyWhere([]model.RefundAttemptsPrimaryID{primaryID})
	query := fmt.Sprintf(refundAttemptsQueries.selectRefundAttempts+" WHERE "+whereQry, defaultRefundAttemptsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &refundAttempts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("refundAttempts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRefundAttemptsByID] failed get refundAttempts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRefundAttemptsByID(ctx context.Context, primaryID model.RefundAttemptsPrimaryID, refundAttempts *model.RefundAttempts, refundAttemptsUpdateFields ...RefundAttemptsUpdateField) (err error) {
	exists, err := repo.IsExistRefundAttemptsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAttempts] failed checking refundAttempts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAttempts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if refundAttempts == nil {
		if len(refundAttemptsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRefundAttemptsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		refundAttempts = &model.RefundAttempts{}
	}
	var (
		defaultRefundAttemptsUpdateFields = defaultRefundAttemptsUpdateFields(*refundAttempts)
		tempUpdateField                   RefundAttemptsUpdateFieldList
		selectFields                      = NewRefundAttemptsSelectFields()
	)
	if len(refundAttemptsUpdateFields) > 0 {
		for _, updateField := range refundAttemptsUpdateFields {
			if updateField.refundAttemptsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRefundAttemptsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRefundAttemptsCompositePrimaryKeyWhere([]model.RefundAttemptsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRefundAttemptsCommand(defaultRefundAttemptsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(refundAttemptsQueries.updateRefundAttempts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAttempts] error when try to update refundAttempts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateRefundAttemptsByFilter(ctx context.Context, filter model.Filter, refundAttemptsUpdateFields ...RefundAttemptsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(refundAttemptsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields RefundAttemptsUpdateFieldList
		selectFields = NewRefundAttemptsSelectFields()
	)
	for _, updateField := range refundAttemptsUpdateFields {
		if updateField.refundAttemptsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsRefundAttemptsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := refundAttemptsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeRefundAttemptsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"refund_attempts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAttemptsByFilter] error when try to update refundAttempts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAttemptsByFilter] failed get rows affected")
	}
	return
}

var (
	refundAttemptsQueries = struct {
		selectRefundAttempts      string
		selectCountRefundAttempts string
		deleteRefundAttempts      string
		updateRefundAttempts      string
		insertRefundAttempts      string
	}{
		selectRefundAttempts:      "SELECT %s FROM \"refund_attempts\"",
		selectCountRefundAttempts: "SELECT COUNT(\"id\") FROM \"refund_attempts\"",
		deleteRefundAttempts:      "DELETE FROM \"refund_attempts\"",
		updateRefundAttempts:      "UPDATE \"refund_attempts\" SET %s ",
		insertRefundAttempts:      "INSERT INTO \"refund_attempts\" %s VALUES %s",
	}
)

type RefundAttemptsRepository interface {
	CreateRefundAttempts(ctx context.Context, refundAttempts *model.RefundAttempts, fieldsInsert ...RefundAttemptsField) error
	BulkCreateRefundAttempts(ctx context.Context, refundAttemptsList []*model.RefundAttempts, fieldsInsert ...RefundAttemptsField) error
	ResolveRefundAttempts(ctx context.Context, selectFields ...RefundAttemptsField) (model.RefundAttemptsList, error)
	ResolveRefundAttemptsByID(ctx context.Context, primaryID model.RefundAttemptsPrimaryID, selectFields ...RefundAttemptsField) (model.RefundAttempts, error)
	UpdateRefundAttemptsByID(ctx context.Context, id model.RefundAttemptsPrimaryID, refundAttempts *model.RefundAttempts, refundAttemptsUpdateFields ...RefundAttemptsUpdateField) error
	UpdateRefundAttemptsByFilter(ctx context.Context, filter model.Filter, refundAttemptsUpdateFields ...RefundAttemptsUpdateField) (rowsAffected int64, err error)
	BulkUpdateRefundAttempts(ctx context.Context, refundAttemptsListMap map[model.RefundAttemptsPrimaryID]*model.RefundAttempts, RefundAttemptssMapUpdateFieldsRequest map[model.RefundAttemptsPrimaryID]RefundAttemptsUpdateFieldList) (err error)
	DeleteRefundAttemptsByID(ctx context.Context, id model.RefundAttemptsPrimaryID) error
	BulkDeleteRefundAttemptsByIDs(ctx context.Context, ids []model.RefundAttemptsPrimaryID) error
	ResolveRefundAttemptsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundAttemptsFilterResult, err error)
	IsExistRefundAttemptsByIDs(ctx context.Context, ids []model.RefundAttemptsPrimaryID) (exists bool, notFoundIds []model.RefundAttemptsPrimaryID, err error)
	IsExistRefundAttemptsByID(ctx context.Context, id model.RefundAttemptsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
