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

func composeInsertFieldsAndParamsFinanceDashboardRefreshJobs(financeDashboardRefreshJobsList []model.FinanceDashboardRefreshJobs, fieldsInsert ...FinanceDashboardRefreshJobsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceDashboardRefreshJobsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeDashboardRefreshJobs := range financeDashboardRefreshJobsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeDashboardRefreshJobs.Id)
			case selectField.JobKey():
				args = append(args, financeDashboardRefreshJobs.JobKey)
			case selectField.RefreshScope():
				args = append(args, financeDashboardRefreshJobs.RefreshScope)
			case selectField.ScopeRef():
				args = append(args, financeDashboardRefreshJobs.ScopeRef)
			case selectField.CurrencyCode():
				args = append(args, financeDashboardRefreshJobs.CurrencyCode)
			case selectField.IdempotencyKey():
				args = append(args, financeDashboardRefreshJobs.IdempotencyKey)
			case selectField.RefreshStatus():
				args = append(args, financeDashboardRefreshJobs.RefreshStatus)
			case selectField.RequestedAt():
				args = append(args, financeDashboardRefreshJobs.RequestedAt)
			case selectField.StartedAt():
				args = append(args, financeDashboardRefreshJobs.StartedAt)
			case selectField.FinishedAt():
				args = append(args, financeDashboardRefreshJobs.FinishedAt)
			case selectField.ErrorCode():
				args = append(args, financeDashboardRefreshJobs.ErrorCode)
			case selectField.ErrorDetail():
				args = append(args, financeDashboardRefreshJobs.ErrorDetail)
			case selectField.Metadata():
				args = append(args, financeDashboardRefreshJobs.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeDashboardRefreshJobs.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeDashboardRefreshJobs.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeDashboardRefreshJobs.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeDashboardRefreshJobs.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeDashboardRefreshJobs.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeDashboardRefreshJobs.MetaDeletedBy)

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

func composeFinanceDashboardRefreshJobsCompositePrimaryKeyWhere(primaryIDs []model.FinanceDashboardRefreshJobsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_dashboard_refresh_jobs\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceDashboardRefreshJobsSelectFields() string {
	fields := NewFinanceDashboardRefreshJobsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceDashboardRefreshJobsSelectFields(selectFields ...FinanceDashboardRefreshJobsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceDashboardRefreshJobsField string
type FinanceDashboardRefreshJobsFieldList []FinanceDashboardRefreshJobsField

type FinanceDashboardRefreshJobsSelectFields struct {
}

func (ss FinanceDashboardRefreshJobsSelectFields) Id() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("id")
}

func (ss FinanceDashboardRefreshJobsSelectFields) JobKey() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("job_key")
}

func (ss FinanceDashboardRefreshJobsSelectFields) RefreshScope() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("refresh_scope")
}

func (ss FinanceDashboardRefreshJobsSelectFields) ScopeRef() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("scope_ref")
}

func (ss FinanceDashboardRefreshJobsSelectFields) CurrencyCode() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("currency_code")
}

func (ss FinanceDashboardRefreshJobsSelectFields) IdempotencyKey() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("idempotency_key")
}

func (ss FinanceDashboardRefreshJobsSelectFields) RefreshStatus() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("refresh_status")
}

func (ss FinanceDashboardRefreshJobsSelectFields) RequestedAt() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("requested_at")
}

func (ss FinanceDashboardRefreshJobsSelectFields) StartedAt() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("started_at")
}

func (ss FinanceDashboardRefreshJobsSelectFields) FinishedAt() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("finished_at")
}

func (ss FinanceDashboardRefreshJobsSelectFields) ErrorCode() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("error_code")
}

func (ss FinanceDashboardRefreshJobsSelectFields) ErrorDetail() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("error_detail")
}

func (ss FinanceDashboardRefreshJobsSelectFields) Metadata() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("metadata")
}

func (ss FinanceDashboardRefreshJobsSelectFields) MetaCreatedAt() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("meta_created_at")
}

func (ss FinanceDashboardRefreshJobsSelectFields) MetaCreatedBy() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("meta_created_by")
}

func (ss FinanceDashboardRefreshJobsSelectFields) MetaUpdatedAt() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("meta_updated_at")
}

func (ss FinanceDashboardRefreshJobsSelectFields) MetaUpdatedBy() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("meta_updated_by")
}

func (ss FinanceDashboardRefreshJobsSelectFields) MetaDeletedAt() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("meta_deleted_at")
}

func (ss FinanceDashboardRefreshJobsSelectFields) MetaDeletedBy() FinanceDashboardRefreshJobsField {
	return FinanceDashboardRefreshJobsField("meta_deleted_by")
}

func (ss FinanceDashboardRefreshJobsSelectFields) All() FinanceDashboardRefreshJobsFieldList {
	return []FinanceDashboardRefreshJobsField{
		ss.Id(),
		ss.JobKey(),
		ss.RefreshScope(),
		ss.ScopeRef(),
		ss.CurrencyCode(),
		ss.IdempotencyKey(),
		ss.RefreshStatus(),
		ss.RequestedAt(),
		ss.StartedAt(),
		ss.FinishedAt(),
		ss.ErrorCode(),
		ss.ErrorDetail(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinanceDashboardRefreshJobsSelectFields() FinanceDashboardRefreshJobsSelectFields {
	return FinanceDashboardRefreshJobsSelectFields{}
}

type FinanceDashboardRefreshJobsUpdateFieldOption struct {
	useIncrement bool
}
type FinanceDashboardRefreshJobsUpdateField struct {
	financeDashboardRefreshJobsField FinanceDashboardRefreshJobsField
	opt                              FinanceDashboardRefreshJobsUpdateFieldOption
	value                            interface{}
}
type FinanceDashboardRefreshJobsUpdateFieldList []FinanceDashboardRefreshJobsUpdateField

func defaultFinanceDashboardRefreshJobsUpdateFieldOption() FinanceDashboardRefreshJobsUpdateFieldOption {
	return FinanceDashboardRefreshJobsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceDashboardRefreshJobsOption(useIncrement bool) func(*FinanceDashboardRefreshJobsUpdateFieldOption) {
	return func(pcufo *FinanceDashboardRefreshJobsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceDashboardRefreshJobsUpdateField(field FinanceDashboardRefreshJobsField, val interface{}, opts ...func(*FinanceDashboardRefreshJobsUpdateFieldOption)) FinanceDashboardRefreshJobsUpdateField {
	defaultOpt := defaultFinanceDashboardRefreshJobsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceDashboardRefreshJobsUpdateField{
		financeDashboardRefreshJobsField: field,
		value:                            val,
		opt:                              defaultOpt,
	}
}
func defaultFinanceDashboardRefreshJobsUpdateFields(financeDashboardRefreshJobs model.FinanceDashboardRefreshJobs) (financeDashboardRefreshJobsUpdateFieldList FinanceDashboardRefreshJobsUpdateFieldList) {
	selectFields := NewFinanceDashboardRefreshJobsSelectFields()
	financeDashboardRefreshJobsUpdateFieldList = append(financeDashboardRefreshJobsUpdateFieldList,
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.Id(), financeDashboardRefreshJobs.Id),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.JobKey(), financeDashboardRefreshJobs.JobKey),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.RefreshScope(), financeDashboardRefreshJobs.RefreshScope),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.ScopeRef(), financeDashboardRefreshJobs.ScopeRef),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.CurrencyCode(), financeDashboardRefreshJobs.CurrencyCode),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.IdempotencyKey(), financeDashboardRefreshJobs.IdempotencyKey),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.RefreshStatus(), financeDashboardRefreshJobs.RefreshStatus),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.RequestedAt(), financeDashboardRefreshJobs.RequestedAt),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.StartedAt(), financeDashboardRefreshJobs.StartedAt),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.FinishedAt(), financeDashboardRefreshJobs.FinishedAt),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.ErrorCode(), financeDashboardRefreshJobs.ErrorCode),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.ErrorDetail(), financeDashboardRefreshJobs.ErrorDetail),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.Metadata(), financeDashboardRefreshJobs.Metadata),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.MetaCreatedAt(), financeDashboardRefreshJobs.MetaCreatedAt),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.MetaCreatedBy(), financeDashboardRefreshJobs.MetaCreatedBy),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.MetaUpdatedAt(), financeDashboardRefreshJobs.MetaUpdatedAt),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.MetaUpdatedBy(), financeDashboardRefreshJobs.MetaUpdatedBy),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.MetaDeletedAt(), financeDashboardRefreshJobs.MetaDeletedAt),
		NewFinanceDashboardRefreshJobsUpdateField(selectFields.MetaDeletedBy(), financeDashboardRefreshJobs.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceDashboardRefreshJobsCommand(financeDashboardRefreshJobsUpdateFieldList FinanceDashboardRefreshJobsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeDashboardRefreshJobsUpdateFieldList {
		field := string(updateField.financeDashboardRefreshJobsField)
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

func (repo *RepositoryImpl) BulkCreateFinanceDashboardRefreshJobs(ctx context.Context, financeDashboardRefreshJobsList []*model.FinanceDashboardRefreshJobs, fieldsInsert ...FinanceDashboardRefreshJobsField) (err error) {
	var (
		fieldsStr                            string
		valueListStr                         []string
		argsList                             []interface{}
		primaryIds                           []model.FinanceDashboardRefreshJobsPrimaryID
		financeDashboardRefreshJobsValueList []model.FinanceDashboardRefreshJobs
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceDashboardRefreshJobsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeDashboardRefreshJobs := range financeDashboardRefreshJobsList {

		primaryIds = append(primaryIds, financeDashboardRefreshJobs.ToFinanceDashboardRefreshJobsPrimaryID())

		financeDashboardRefreshJobsValueList = append(financeDashboardRefreshJobsValueList, *financeDashboardRefreshJobs)
	}

	_, notFoundIds, err := repo.IsExistFinanceDashboardRefreshJobsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceDashboardRefreshJobs] failed checking financeDashboardRefreshJobs whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceDashboardRefreshJobsPrimaryID{}
		mapNotFoundIds := map[model.FinanceDashboardRefreshJobsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeDashboardRefreshJobs", fmt.Sprintf("financeDashboardRefreshJobs with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceDashboardRefreshJobs(financeDashboardRefreshJobsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeDashboardRefreshJobsQueries.insertFinanceDashboardRefreshJobs, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceDashboardRefreshJobs] failed exec create financeDashboardRefreshJobs query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceDashboardRefreshJobsByIDs(ctx context.Context, primaryIDs []model.FinanceDashboardRefreshJobsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceDashboardRefreshJobsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceDashboardRefreshJobsByIDs] failed checking financeDashboardRefreshJobs whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeDashboardRefreshJobs with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_dashboard_refresh_jobs\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeDashboardRefreshJobsQueries.deleteFinanceDashboardRefreshJobs + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceDashboardRefreshJobsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceDashboardRefreshJobsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceDashboardRefreshJobsByIDs(ctx context.Context, ids []model.FinanceDashboardRefreshJobsPrimaryID) (exists bool, notFoundIds []model.FinanceDashboardRefreshJobsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_dashboard_refresh_jobs\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeDashboardRefreshJobsQueries.selectFinanceDashboardRefreshJobs, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceDashboardRefreshJobsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceDashboardRefreshJobsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceDashboardRefreshJobsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceDashboardRefreshJobsPrimaryID]bool{}
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

// BulkUpdateFinanceDashboardRefreshJobs is used to bulk update financeDashboardRefreshJobs, by default it will update all field
// if want to update specific field, then fill financeDashboardRefreshJobssMapUpdateFieldsRequest else please fill financeDashboardRefreshJobssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceDashboardRefreshJobs(ctx context.Context, financeDashboardRefreshJobssMap map[model.FinanceDashboardRefreshJobsPrimaryID]*model.FinanceDashboardRefreshJobs, financeDashboardRefreshJobssMapUpdateFieldsRequest map[model.FinanceDashboardRefreshJobsPrimaryID]FinanceDashboardRefreshJobsUpdateFieldList) (err error) {
	if len(financeDashboardRefreshJobssMap) == 0 && len(financeDashboardRefreshJobssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeDashboardRefreshJobssMapUpdateField map[model.FinanceDashboardRefreshJobsPrimaryID]FinanceDashboardRefreshJobsUpdateFieldList = map[model.FinanceDashboardRefreshJobsPrimaryID]FinanceDashboardRefreshJobsUpdateFieldList{}
		asTableValues                              string                                                                                    = "myvalues"
	)

	if len(financeDashboardRefreshJobssMap) > 0 {
		for id, financeDashboardRefreshJobs := range financeDashboardRefreshJobssMap {
			if financeDashboardRefreshJobs == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceDashboardRefreshJobs] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeDashboardRefreshJobssMapUpdateField[id] = defaultFinanceDashboardRefreshJobsUpdateFields(*financeDashboardRefreshJobs)
		}
	} else {
		financeDashboardRefreshJobssMapUpdateField = financeDashboardRefreshJobssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceDashboardRefreshJobsQuery(financeDashboardRefreshJobssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceDashboardRefreshJobsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceDashboardRefreshJobs] failed checking financeDashboardRefreshJobs whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeDashboardRefreshJobs with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceDashboardRefreshJobsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_dashboard_refresh_jobs\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceDashboardRefreshJobs] failed exec query")
	}
	return
}

type FinanceDashboardRefreshJobsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceDashboardRefreshJobsFieldParameter(param string, args ...interface{}) FinanceDashboardRefreshJobsFieldParameter {
	return FinanceDashboardRefreshJobsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceDashboardRefreshJobsQuery(mapFinanceDashboardRefreshJobss map[model.FinanceDashboardRefreshJobsPrimaryID]FinanceDashboardRefreshJobsUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceDashboardRefreshJobsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceDashboardRefreshJobsPrimaryID]map[string]interface{}{}
	financeDashboardRefreshJobsSelectFields := NewFinanceDashboardRefreshJobsSelectFields()
	for id, updateFields := range mapFinanceDashboardRefreshJobss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeDashboardRefreshJobsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceDashboardRefreshJobss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceDashboardRefreshJobsFieldType(updateField.financeDashboardRefreshJobsField)))
			args = append(args, fields[string(updateField.financeDashboardRefreshJobsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeDashboardRefreshJobsField))
		if updateField.financeDashboardRefreshJobsField == financeDashboardRefreshJobsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeDashboardRefreshJobsField, asTableValues, updateField.financeDashboardRefreshJobsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeDashboardRefreshJobsField,
				"\"finance_dashboard_refresh_jobs\"", updateField.financeDashboardRefreshJobsField,
				asTableValues, updateField.financeDashboardRefreshJobsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceDashboardRefreshJobsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceDashboardRefreshJobsPrimaryID, asTableValue string) (whereQry string) {
	financeDashboardRefreshJobsSelectFields := NewFinanceDashboardRefreshJobsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_dashboard_refresh_jobs\".\"id\" = %s.\"id\"::"+GetFinanceDashboardRefreshJobsFieldType(financeDashboardRefreshJobsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceDashboardRefreshJobsFieldType(financeDashboardRefreshJobsField FinanceDashboardRefreshJobsField) string {
	selectFinanceDashboardRefreshJobsFields := NewFinanceDashboardRefreshJobsSelectFields()
	switch financeDashboardRefreshJobsField {

	case selectFinanceDashboardRefreshJobsFields.Id():
		return "uuid"

	case selectFinanceDashboardRefreshJobsFields.JobKey():
		return "text"

	case selectFinanceDashboardRefreshJobsFields.RefreshScope():
		return "text"

	case selectFinanceDashboardRefreshJobsFields.ScopeRef():
		return "uuid"

	case selectFinanceDashboardRefreshJobsFields.CurrencyCode():
		return "text"

	case selectFinanceDashboardRefreshJobsFields.IdempotencyKey():
		return "text"

	case selectFinanceDashboardRefreshJobsFields.RefreshStatus():
		return "refresh_status_enum"

	case selectFinanceDashboardRefreshJobsFields.RequestedAt():
		return "timestamptz"

	case selectFinanceDashboardRefreshJobsFields.StartedAt():
		return "timestamptz"

	case selectFinanceDashboardRefreshJobsFields.FinishedAt():
		return "timestamptz"

	case selectFinanceDashboardRefreshJobsFields.ErrorCode():
		return "text"

	case selectFinanceDashboardRefreshJobsFields.ErrorDetail():
		return "text"

	case selectFinanceDashboardRefreshJobsFields.Metadata():
		return "jsonb"

	case selectFinanceDashboardRefreshJobsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceDashboardRefreshJobsFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceDashboardRefreshJobsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceDashboardRefreshJobsFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceDashboardRefreshJobsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceDashboardRefreshJobsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceDashboardRefreshJobs(ctx context.Context, financeDashboardRefreshJobs *model.FinanceDashboardRefreshJobs, fieldsInsert ...FinanceDashboardRefreshJobsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceDashboardRefreshJobsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceDashboardRefreshJobsPrimaryID{
		Id: financeDashboardRefreshJobs.Id,
	}
	exists, err := repo.IsExistFinanceDashboardRefreshJobsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceDashboardRefreshJobs] failed checking financeDashboardRefreshJobs whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeDashboardRefreshJobs", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceDashboardRefreshJobs([]model.FinanceDashboardRefreshJobs{*financeDashboardRefreshJobs}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeDashboardRefreshJobsQueries.insertFinanceDashboardRefreshJobs, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceDashboardRefreshJobs] failed exec create financeDashboardRefreshJobs query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceDashboardRefreshJobsByID(ctx context.Context, primaryID model.FinanceDashboardRefreshJobsPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceDashboardRefreshJobsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceDashboardRefreshJobsByID] failed checking financeDashboardRefreshJobs whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeDashboardRefreshJobs with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceDashboardRefreshJobsCompositePrimaryKeyWhere([]model.FinanceDashboardRefreshJobsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeDashboardRefreshJobsQueries.deleteFinanceDashboardRefreshJobs + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceDashboardRefreshJobsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceDashboardRefreshJobsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceDashboardRefreshJobsFilterResult, err error) {
	query, args, err := composeFinanceDashboardRefreshJobsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceDashboardRefreshJobsByFilter] failed compose financeDashboardRefreshJobs filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceDashboardRefreshJobsByFilter] failed get financeDashboardRefreshJobs by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceDashboardRefreshJobsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceDashboardRefreshJobsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceDashboardRefreshJobsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceDashboardRefreshJobsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceDashboardRefreshJobsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 19+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["job_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"job_key\"")
			selectedColumns["job_key"] = struct{}{}
		}
		if _, selected := selectedColumns["refresh_scope"]; !selected {
			selectColumns = append(selectColumns, "base.\"refresh_scope\"")
			selectedColumns["refresh_scope"] = struct{}{}
		}
		if _, selected := selectedColumns["scope_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"scope_ref\"")
			selectedColumns["scope_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["refresh_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"refresh_status\"")
			selectedColumns["refresh_status"] = struct{}{}
		}
		if _, selected := selectedColumns["requested_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"requested_at\"")
			selectedColumns["requested_at"] = struct{}{}
		}
		if _, selected := selectedColumns["started_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"started_at\"")
			selectedColumns["started_at"] = struct{}{}
		}
		if _, selected := selectedColumns["finished_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"finished_at\"")
			selectedColumns["finished_at"] = struct{}{}
		}
		if _, selected := selectedColumns["error_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"error_code\"")
			selectedColumns["error_code"] = struct{}{}
		}
		if _, selected := selectedColumns["error_detail"]; !selected {
			selectColumns = append(selectColumns, "base.\"error_detail\"")
			selectedColumns["error_detail"] = struct{}{}
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

type financeDashboardRefreshJobsFilterPlaceholder struct {
	index int
}

func (p *financeDashboardRefreshJobsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceDashboardRefreshJobsFilterPredicate(filterField model.FilterField, placeholders *financeDashboardRefreshJobsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceDashboardRefreshJobsFilterSQLExpr(spec)
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

func composeFinanceDashboardRefreshJobsFilterGroup(group model.FilterGroup, placeholders *financeDashboardRefreshJobsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceDashboardRefreshJobsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceDashboardRefreshJobsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceDashboardRefreshJobsFilterWhereQueries(filter model.Filter, placeholders *financeDashboardRefreshJobsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceDashboardRefreshJobsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceDashboardRefreshJobsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceDashboardRefreshJobsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceDashboardRefreshJobsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceDashboardRefreshJobsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceDashboardRefreshJobsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeDashboardRefreshJobsFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceDashboardRefreshJobsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceDashboardRefreshJobsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceDashboardRefreshJobsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceDashboardRefreshJobsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_dashboard_refresh_jobs\" base%s", strings.Join(selectColumns, ","), composeFinanceDashboardRefreshJobsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceDashboardRefreshJobsByID(ctx context.Context, primaryID model.FinanceDashboardRefreshJobsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceDashboardRefreshJobsCompositePrimaryKeyWhere([]model.FinanceDashboardRefreshJobsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeDashboardRefreshJobsQueries.selectCountFinanceDashboardRefreshJobs, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceDashboardRefreshJobsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceDashboardRefreshJobs(ctx context.Context, selectFields ...FinanceDashboardRefreshJobsField) (financeDashboardRefreshJobsList model.FinanceDashboardRefreshJobsList, err error) {
	var (
		defaultFinanceDashboardRefreshJobsSelectFields = defaultFinanceDashboardRefreshJobsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceDashboardRefreshJobsSelectFields = composeFinanceDashboardRefreshJobsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeDashboardRefreshJobsQueries.selectFinanceDashboardRefreshJobs, defaultFinanceDashboardRefreshJobsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeDashboardRefreshJobsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceDashboardRefreshJobs] failed get financeDashboardRefreshJobs list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceDashboardRefreshJobsByID(ctx context.Context, primaryID model.FinanceDashboardRefreshJobsPrimaryID, selectFields ...FinanceDashboardRefreshJobsField) (financeDashboardRefreshJobs model.FinanceDashboardRefreshJobs, err error) {
	var (
		defaultFinanceDashboardRefreshJobsSelectFields = defaultFinanceDashboardRefreshJobsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceDashboardRefreshJobsSelectFields = composeFinanceDashboardRefreshJobsSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceDashboardRefreshJobsCompositePrimaryKeyWhere([]model.FinanceDashboardRefreshJobsPrimaryID{primaryID})
	query := fmt.Sprintf(financeDashboardRefreshJobsQueries.selectFinanceDashboardRefreshJobs+" WHERE "+whereQry, defaultFinanceDashboardRefreshJobsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeDashboardRefreshJobs, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeDashboardRefreshJobs with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceDashboardRefreshJobsByID] failed get financeDashboardRefreshJobs")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceDashboardRefreshJobsByID(ctx context.Context, primaryID model.FinanceDashboardRefreshJobsPrimaryID, financeDashboardRefreshJobs *model.FinanceDashboardRefreshJobs, financeDashboardRefreshJobsUpdateFields ...FinanceDashboardRefreshJobsUpdateField) (err error) {
	exists, err := repo.IsExistFinanceDashboardRefreshJobsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceDashboardRefreshJobs] failed checking financeDashboardRefreshJobs whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeDashboardRefreshJobs with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeDashboardRefreshJobs == nil {
		if len(financeDashboardRefreshJobsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceDashboardRefreshJobsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeDashboardRefreshJobs = &model.FinanceDashboardRefreshJobs{}
	}
	var (
		defaultFinanceDashboardRefreshJobsUpdateFields = defaultFinanceDashboardRefreshJobsUpdateFields(*financeDashboardRefreshJobs)
		tempUpdateField                                FinanceDashboardRefreshJobsUpdateFieldList
		selectFields                                   = NewFinanceDashboardRefreshJobsSelectFields()
	)
	if len(financeDashboardRefreshJobsUpdateFields) > 0 {
		for _, updateField := range financeDashboardRefreshJobsUpdateFields {
			if updateField.financeDashboardRefreshJobsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceDashboardRefreshJobsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceDashboardRefreshJobsCompositePrimaryKeyWhere([]model.FinanceDashboardRefreshJobsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceDashboardRefreshJobsCommand(defaultFinanceDashboardRefreshJobsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeDashboardRefreshJobsQueries.updateFinanceDashboardRefreshJobs+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceDashboardRefreshJobs] error when try to update financeDashboardRefreshJobs by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceDashboardRefreshJobsByFilter(ctx context.Context, filter model.Filter, financeDashboardRefreshJobsUpdateFields ...FinanceDashboardRefreshJobsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeDashboardRefreshJobsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceDashboardRefreshJobsUpdateFieldList
		selectFields = NewFinanceDashboardRefreshJobsSelectFields()
	)
	for _, updateField := range financeDashboardRefreshJobsUpdateFields {
		if updateField.financeDashboardRefreshJobsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceDashboardRefreshJobsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeDashboardRefreshJobsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceDashboardRefreshJobsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_dashboard_refresh_jobs\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceDashboardRefreshJobsByFilter] error when try to update financeDashboardRefreshJobs by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceDashboardRefreshJobsByFilter] failed get rows affected")
	}
	return
}

var (
	financeDashboardRefreshJobsQueries = struct {
		selectFinanceDashboardRefreshJobs      string
		selectCountFinanceDashboardRefreshJobs string
		deleteFinanceDashboardRefreshJobs      string
		updateFinanceDashboardRefreshJobs      string
		insertFinanceDashboardRefreshJobs      string
	}{
		selectFinanceDashboardRefreshJobs:      "SELECT %s FROM \"finance_dashboard_refresh_jobs\"",
		selectCountFinanceDashboardRefreshJobs: "SELECT COUNT(\"id\") FROM \"finance_dashboard_refresh_jobs\"",
		deleteFinanceDashboardRefreshJobs:      "DELETE FROM \"finance_dashboard_refresh_jobs\"",
		updateFinanceDashboardRefreshJobs:      "UPDATE \"finance_dashboard_refresh_jobs\" SET %s ",
		insertFinanceDashboardRefreshJobs:      "INSERT INTO \"finance_dashboard_refresh_jobs\" %s VALUES %s",
	}
)

type FinanceDashboardRefreshJobsRepository interface {
	CreateFinanceDashboardRefreshJobs(ctx context.Context, financeDashboardRefreshJobs *model.FinanceDashboardRefreshJobs, fieldsInsert ...FinanceDashboardRefreshJobsField) error
	BulkCreateFinanceDashboardRefreshJobs(ctx context.Context, financeDashboardRefreshJobsList []*model.FinanceDashboardRefreshJobs, fieldsInsert ...FinanceDashboardRefreshJobsField) error
	ResolveFinanceDashboardRefreshJobs(ctx context.Context, selectFields ...FinanceDashboardRefreshJobsField) (model.FinanceDashboardRefreshJobsList, error)
	ResolveFinanceDashboardRefreshJobsByID(ctx context.Context, primaryID model.FinanceDashboardRefreshJobsPrimaryID, selectFields ...FinanceDashboardRefreshJobsField) (model.FinanceDashboardRefreshJobs, error)
	UpdateFinanceDashboardRefreshJobsByID(ctx context.Context, id model.FinanceDashboardRefreshJobsPrimaryID, financeDashboardRefreshJobs *model.FinanceDashboardRefreshJobs, financeDashboardRefreshJobsUpdateFields ...FinanceDashboardRefreshJobsUpdateField) error
	UpdateFinanceDashboardRefreshJobsByFilter(ctx context.Context, filter model.Filter, financeDashboardRefreshJobsUpdateFields ...FinanceDashboardRefreshJobsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceDashboardRefreshJobs(ctx context.Context, financeDashboardRefreshJobsListMap map[model.FinanceDashboardRefreshJobsPrimaryID]*model.FinanceDashboardRefreshJobs, FinanceDashboardRefreshJobssMapUpdateFieldsRequest map[model.FinanceDashboardRefreshJobsPrimaryID]FinanceDashboardRefreshJobsUpdateFieldList) (err error)
	DeleteFinanceDashboardRefreshJobsByID(ctx context.Context, id model.FinanceDashboardRefreshJobsPrimaryID) error
	BulkDeleteFinanceDashboardRefreshJobsByIDs(ctx context.Context, ids []model.FinanceDashboardRefreshJobsPrimaryID) error
	ResolveFinanceDashboardRefreshJobsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceDashboardRefreshJobsFilterResult, err error)
	IsExistFinanceDashboardRefreshJobsByIDs(ctx context.Context, ids []model.FinanceDashboardRefreshJobsPrimaryID) (exists bool, notFoundIds []model.FinanceDashboardRefreshJobsPrimaryID, err error)
	IsExistFinanceDashboardRefreshJobsByID(ctx context.Context, id model.FinanceDashboardRefreshJobsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
