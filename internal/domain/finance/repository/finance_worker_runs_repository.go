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

func composeInsertFieldsAndParamsFinanceWorkerRuns(financeWorkerRunsList []model.FinanceWorkerRuns, fieldsInsert ...FinanceWorkerRunsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceWorkerRunsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeWorkerRuns := range financeWorkerRunsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeWorkerRuns.Id)
			case selectField.WorkerName():
				args = append(args, financeWorkerRuns.WorkerName)
			case selectField.RunKey():
				args = append(args, financeWorkerRuns.RunKey)
			case selectField.RunStatus():
				args = append(args, financeWorkerRuns.RunStatus)
			case selectField.StartedAt():
				args = append(args, financeWorkerRuns.StartedAt)
			case selectField.FinishedAt():
				args = append(args, financeWorkerRuns.FinishedAt)
			case selectField.ProcessedCount():
				args = append(args, financeWorkerRuns.ProcessedCount)
			case selectField.FailedCount():
				args = append(args, financeWorkerRuns.FailedCount)
			case selectField.ErrorCode():
				args = append(args, financeWorkerRuns.ErrorCode)
			case selectField.ErrorDetail():
				args = append(args, financeWorkerRuns.ErrorDetail)
			case selectField.Metadata():
				args = append(args, financeWorkerRuns.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeWorkerRuns.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeWorkerRuns.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeWorkerRuns.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeWorkerRuns.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeWorkerRuns.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeWorkerRuns.MetaDeletedBy)

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

func composeFinanceWorkerRunsCompositePrimaryKeyWhere(primaryIDs []model.FinanceWorkerRunsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_worker_runs\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceWorkerRunsSelectFields() string {
	fields := NewFinanceWorkerRunsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceWorkerRunsSelectFields(selectFields ...FinanceWorkerRunsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceWorkerRunsField string
type FinanceWorkerRunsFieldList []FinanceWorkerRunsField

type FinanceWorkerRunsSelectFields struct {
}

func (ss FinanceWorkerRunsSelectFields) Id() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("id")
}

func (ss FinanceWorkerRunsSelectFields) WorkerName() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("worker_name")
}

func (ss FinanceWorkerRunsSelectFields) RunKey() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("run_key")
}

func (ss FinanceWorkerRunsSelectFields) RunStatus() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("run_status")
}

func (ss FinanceWorkerRunsSelectFields) StartedAt() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("started_at")
}

func (ss FinanceWorkerRunsSelectFields) FinishedAt() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("finished_at")
}

func (ss FinanceWorkerRunsSelectFields) ProcessedCount() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("processed_count")
}

func (ss FinanceWorkerRunsSelectFields) FailedCount() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("failed_count")
}

func (ss FinanceWorkerRunsSelectFields) ErrorCode() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("error_code")
}

func (ss FinanceWorkerRunsSelectFields) ErrorDetail() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("error_detail")
}

func (ss FinanceWorkerRunsSelectFields) Metadata() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("metadata")
}

func (ss FinanceWorkerRunsSelectFields) MetaCreatedAt() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("meta_created_at")
}

func (ss FinanceWorkerRunsSelectFields) MetaCreatedBy() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("meta_created_by")
}

func (ss FinanceWorkerRunsSelectFields) MetaUpdatedAt() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("meta_updated_at")
}

func (ss FinanceWorkerRunsSelectFields) MetaUpdatedBy() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("meta_updated_by")
}

func (ss FinanceWorkerRunsSelectFields) MetaDeletedAt() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("meta_deleted_at")
}

func (ss FinanceWorkerRunsSelectFields) MetaDeletedBy() FinanceWorkerRunsField {
	return FinanceWorkerRunsField("meta_deleted_by")
}

func (ss FinanceWorkerRunsSelectFields) All() FinanceWorkerRunsFieldList {
	return []FinanceWorkerRunsField{
		ss.Id(),
		ss.WorkerName(),
		ss.RunKey(),
		ss.RunStatus(),
		ss.StartedAt(),
		ss.FinishedAt(),
		ss.ProcessedCount(),
		ss.FailedCount(),
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

func NewFinanceWorkerRunsSelectFields() FinanceWorkerRunsSelectFields {
	return FinanceWorkerRunsSelectFields{}
}

type FinanceWorkerRunsUpdateFieldOption struct {
	useIncrement bool
}
type FinanceWorkerRunsUpdateField struct {
	financeWorkerRunsField FinanceWorkerRunsField
	opt                    FinanceWorkerRunsUpdateFieldOption
	value                  interface{}
}
type FinanceWorkerRunsUpdateFieldList []FinanceWorkerRunsUpdateField

func defaultFinanceWorkerRunsUpdateFieldOption() FinanceWorkerRunsUpdateFieldOption {
	return FinanceWorkerRunsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceWorkerRunsOption(useIncrement bool) func(*FinanceWorkerRunsUpdateFieldOption) {
	return func(pcufo *FinanceWorkerRunsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceWorkerRunsUpdateField(field FinanceWorkerRunsField, val interface{}, opts ...func(*FinanceWorkerRunsUpdateFieldOption)) FinanceWorkerRunsUpdateField {
	defaultOpt := defaultFinanceWorkerRunsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceWorkerRunsUpdateField{
		financeWorkerRunsField: field,
		value:                  val,
		opt:                    defaultOpt,
	}
}
func defaultFinanceWorkerRunsUpdateFields(financeWorkerRuns model.FinanceWorkerRuns) (financeWorkerRunsUpdateFieldList FinanceWorkerRunsUpdateFieldList) {
	selectFields := NewFinanceWorkerRunsSelectFields()
	financeWorkerRunsUpdateFieldList = append(financeWorkerRunsUpdateFieldList,
		NewFinanceWorkerRunsUpdateField(selectFields.Id(), financeWorkerRuns.Id),
		NewFinanceWorkerRunsUpdateField(selectFields.WorkerName(), financeWorkerRuns.WorkerName),
		NewFinanceWorkerRunsUpdateField(selectFields.RunKey(), financeWorkerRuns.RunKey),
		NewFinanceWorkerRunsUpdateField(selectFields.RunStatus(), financeWorkerRuns.RunStatus),
		NewFinanceWorkerRunsUpdateField(selectFields.StartedAt(), financeWorkerRuns.StartedAt),
		NewFinanceWorkerRunsUpdateField(selectFields.FinishedAt(), financeWorkerRuns.FinishedAt),
		NewFinanceWorkerRunsUpdateField(selectFields.ProcessedCount(), financeWorkerRuns.ProcessedCount),
		NewFinanceWorkerRunsUpdateField(selectFields.FailedCount(), financeWorkerRuns.FailedCount),
		NewFinanceWorkerRunsUpdateField(selectFields.ErrorCode(), financeWorkerRuns.ErrorCode),
		NewFinanceWorkerRunsUpdateField(selectFields.ErrorDetail(), financeWorkerRuns.ErrorDetail),
		NewFinanceWorkerRunsUpdateField(selectFields.Metadata(), financeWorkerRuns.Metadata),
		NewFinanceWorkerRunsUpdateField(selectFields.MetaCreatedAt(), financeWorkerRuns.MetaCreatedAt),
		NewFinanceWorkerRunsUpdateField(selectFields.MetaCreatedBy(), financeWorkerRuns.MetaCreatedBy),
		NewFinanceWorkerRunsUpdateField(selectFields.MetaUpdatedAt(), financeWorkerRuns.MetaUpdatedAt),
		NewFinanceWorkerRunsUpdateField(selectFields.MetaUpdatedBy(), financeWorkerRuns.MetaUpdatedBy),
		NewFinanceWorkerRunsUpdateField(selectFields.MetaDeletedAt(), financeWorkerRuns.MetaDeletedAt),
		NewFinanceWorkerRunsUpdateField(selectFields.MetaDeletedBy(), financeWorkerRuns.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceWorkerRunsCommand(financeWorkerRunsUpdateFieldList FinanceWorkerRunsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeWorkerRunsUpdateFieldList {
		field := string(updateField.financeWorkerRunsField)
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

func (repo *RepositoryImpl) BulkCreateFinanceWorkerRuns(ctx context.Context, financeWorkerRunsList []*model.FinanceWorkerRuns, fieldsInsert ...FinanceWorkerRunsField) (err error) {
	var (
		fieldsStr                  string
		valueListStr               []string
		argsList                   []interface{}
		primaryIds                 []model.FinanceWorkerRunsPrimaryID
		financeWorkerRunsValueList []model.FinanceWorkerRuns
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceWorkerRunsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeWorkerRuns := range financeWorkerRunsList {

		primaryIds = append(primaryIds, financeWorkerRuns.ToFinanceWorkerRunsPrimaryID())

		financeWorkerRunsValueList = append(financeWorkerRunsValueList, *financeWorkerRuns)
	}

	_, notFoundIds, err := repo.IsExistFinanceWorkerRunsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceWorkerRuns] failed checking financeWorkerRuns whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceWorkerRunsPrimaryID{}
		mapNotFoundIds := map[model.FinanceWorkerRunsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeWorkerRuns", fmt.Sprintf("financeWorkerRuns with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceWorkerRuns(financeWorkerRunsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeWorkerRunsQueries.insertFinanceWorkerRuns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceWorkerRuns] failed exec create financeWorkerRuns query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceWorkerRunsByIDs(ctx context.Context, primaryIDs []model.FinanceWorkerRunsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceWorkerRunsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceWorkerRunsByIDs] failed checking financeWorkerRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeWorkerRuns with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_worker_runs\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeWorkerRunsQueries.deleteFinanceWorkerRuns + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceWorkerRunsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceWorkerRunsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceWorkerRunsByIDs(ctx context.Context, ids []model.FinanceWorkerRunsPrimaryID) (exists bool, notFoundIds []model.FinanceWorkerRunsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_worker_runs\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeWorkerRunsQueries.selectFinanceWorkerRuns, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceWorkerRunsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceWorkerRunsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceWorkerRunsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceWorkerRunsPrimaryID]bool{}
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

// BulkUpdateFinanceWorkerRuns is used to bulk update financeWorkerRuns, by default it will update all field
// if want to update specific field, then fill financeWorkerRunssMapUpdateFieldsRequest else please fill financeWorkerRunssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceWorkerRuns(ctx context.Context, financeWorkerRunssMap map[model.FinanceWorkerRunsPrimaryID]*model.FinanceWorkerRuns, financeWorkerRunssMapUpdateFieldsRequest map[model.FinanceWorkerRunsPrimaryID]FinanceWorkerRunsUpdateFieldList) (err error) {
	if len(financeWorkerRunssMap) == 0 && len(financeWorkerRunssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeWorkerRunssMapUpdateField map[model.FinanceWorkerRunsPrimaryID]FinanceWorkerRunsUpdateFieldList = map[model.FinanceWorkerRunsPrimaryID]FinanceWorkerRunsUpdateFieldList{}
		asTableValues                    string                                                                = "myvalues"
	)

	if len(financeWorkerRunssMap) > 0 {
		for id, financeWorkerRuns := range financeWorkerRunssMap {
			if financeWorkerRuns == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceWorkerRuns] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeWorkerRunssMapUpdateField[id] = defaultFinanceWorkerRunsUpdateFields(*financeWorkerRuns)
		}
	} else {
		financeWorkerRunssMapUpdateField = financeWorkerRunssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceWorkerRunsQuery(financeWorkerRunssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceWorkerRunsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceWorkerRuns] failed checking financeWorkerRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeWorkerRuns with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceWorkerRunsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_worker_runs\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceWorkerRuns] failed exec query")
	}
	return
}

type FinanceWorkerRunsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceWorkerRunsFieldParameter(param string, args ...interface{}) FinanceWorkerRunsFieldParameter {
	return FinanceWorkerRunsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceWorkerRunsQuery(mapFinanceWorkerRunss map[model.FinanceWorkerRunsPrimaryID]FinanceWorkerRunsUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceWorkerRunsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceWorkerRunsPrimaryID]map[string]interface{}{}
	financeWorkerRunsSelectFields := NewFinanceWorkerRunsSelectFields()
	for id, updateFields := range mapFinanceWorkerRunss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeWorkerRunsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceWorkerRunss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceWorkerRunsFieldType(updateField.financeWorkerRunsField)))
			args = append(args, fields[string(updateField.financeWorkerRunsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeWorkerRunsField))
		if updateField.financeWorkerRunsField == financeWorkerRunsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeWorkerRunsField, asTableValues, updateField.financeWorkerRunsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeWorkerRunsField,
				"\"finance_worker_runs\"", updateField.financeWorkerRunsField,
				asTableValues, updateField.financeWorkerRunsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceWorkerRunsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceWorkerRunsPrimaryID, asTableValue string) (whereQry string) {
	financeWorkerRunsSelectFields := NewFinanceWorkerRunsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_worker_runs\".\"id\" = %s.\"id\"::"+GetFinanceWorkerRunsFieldType(financeWorkerRunsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceWorkerRunsFieldType(financeWorkerRunsField FinanceWorkerRunsField) string {
	selectFinanceWorkerRunsFields := NewFinanceWorkerRunsSelectFields()
	switch financeWorkerRunsField {

	case selectFinanceWorkerRunsFields.Id():
		return "uuid"

	case selectFinanceWorkerRunsFields.WorkerName():
		return "text"

	case selectFinanceWorkerRunsFields.RunKey():
		return "text"

	case selectFinanceWorkerRunsFields.RunStatus():
		return "finance_worker_runs_run_status_enum"

	case selectFinanceWorkerRunsFields.StartedAt():
		return "timestamptz"

	case selectFinanceWorkerRunsFields.FinishedAt():
		return "timestamptz"

	case selectFinanceWorkerRunsFields.ProcessedCount():
		return "int4"

	case selectFinanceWorkerRunsFields.FailedCount():
		return "int4"

	case selectFinanceWorkerRunsFields.ErrorCode():
		return "text"

	case selectFinanceWorkerRunsFields.ErrorDetail():
		return "text"

	case selectFinanceWorkerRunsFields.Metadata():
		return "jsonb"

	case selectFinanceWorkerRunsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceWorkerRunsFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceWorkerRunsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceWorkerRunsFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceWorkerRunsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceWorkerRunsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceWorkerRuns(ctx context.Context, financeWorkerRuns *model.FinanceWorkerRuns, fieldsInsert ...FinanceWorkerRunsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceWorkerRunsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceWorkerRunsPrimaryID{
		Id: financeWorkerRuns.Id,
	}
	exists, err := repo.IsExistFinanceWorkerRunsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceWorkerRuns] failed checking financeWorkerRuns whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeWorkerRuns", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceWorkerRuns([]model.FinanceWorkerRuns{*financeWorkerRuns}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeWorkerRunsQueries.insertFinanceWorkerRuns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceWorkerRuns] failed exec create financeWorkerRuns query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceWorkerRunsByID(ctx context.Context, primaryID model.FinanceWorkerRunsPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceWorkerRunsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceWorkerRunsByID] failed checking financeWorkerRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeWorkerRuns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceWorkerRunsCompositePrimaryKeyWhere([]model.FinanceWorkerRunsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeWorkerRunsQueries.deleteFinanceWorkerRuns + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceWorkerRunsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceWorkerRunsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceWorkerRunsFilterResult, err error) {
	query, args, err := composeFinanceWorkerRunsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceWorkerRunsByFilter] failed compose financeWorkerRuns filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceWorkerRunsByFilter] failed get financeWorkerRuns by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceWorkerRunsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceWorkerRunsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceWorkerRunsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceWorkerRunsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceWorkerRunsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 17 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceWorkerRunsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["worker_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"worker_name\"")
			selectedColumns["worker_name"] = struct{}{}
		}
		if _, selected := selectedColumns["run_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"run_key\"")
			selectedColumns["run_key"] = struct{}{}
		}
		if _, selected := selectedColumns["run_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"run_status\"")
			selectedColumns["run_status"] = struct{}{}
		}
		if _, selected := selectedColumns["started_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"started_at\"")
			selectedColumns["started_at"] = struct{}{}
		}
		if _, selected := selectedColumns["finished_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"finished_at\"")
			selectedColumns["finished_at"] = struct{}{}
		}
		if _, selected := selectedColumns["processed_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"processed_count\"")
			selectedColumns["processed_count"] = struct{}{}
		}
		if _, selected := selectedColumns["failed_count"]; !selected {
			selectColumns = append(selectColumns, "base.\"failed_count\"")
			selectedColumns["failed_count"] = struct{}{}
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

type financeWorkerRunsFilterPlaceholder struct {
	index int
}

func (p *financeWorkerRunsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceWorkerRunsFilterPredicate(filterField model.FilterField, placeholders *financeWorkerRunsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceWorkerRunsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceWorkerRunsFilterSQLExpr(spec)
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

func composeFinanceWorkerRunsFilterGroup(group model.FilterGroup, placeholders *financeWorkerRunsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceWorkerRunsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceWorkerRunsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceWorkerRunsFilterWhereQueries(filter model.Filter, placeholders *financeWorkerRunsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceWorkerRunsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceWorkerRunsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceWorkerRunsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceWorkerRunsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceWorkerRunsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceWorkerRunsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeWorkerRunsFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceWorkerRunsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceWorkerRunsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceWorkerRunsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceWorkerRunsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_worker_runs\" base%s", strings.Join(selectColumns, ","), composeFinanceWorkerRunsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceWorkerRunsByID(ctx context.Context, primaryID model.FinanceWorkerRunsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceWorkerRunsCompositePrimaryKeyWhere([]model.FinanceWorkerRunsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeWorkerRunsQueries.selectCountFinanceWorkerRuns, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceWorkerRunsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceWorkerRuns(ctx context.Context, selectFields ...FinanceWorkerRunsField) (financeWorkerRunsList model.FinanceWorkerRunsList, err error) {
	var (
		defaultFinanceWorkerRunsSelectFields = defaultFinanceWorkerRunsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceWorkerRunsSelectFields = composeFinanceWorkerRunsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeWorkerRunsQueries.selectFinanceWorkerRuns, defaultFinanceWorkerRunsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeWorkerRunsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceWorkerRuns] failed get financeWorkerRuns list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceWorkerRunsByID(ctx context.Context, primaryID model.FinanceWorkerRunsPrimaryID, selectFields ...FinanceWorkerRunsField) (financeWorkerRuns model.FinanceWorkerRuns, err error) {
	var (
		defaultFinanceWorkerRunsSelectFields = defaultFinanceWorkerRunsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceWorkerRunsSelectFields = composeFinanceWorkerRunsSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceWorkerRunsCompositePrimaryKeyWhere([]model.FinanceWorkerRunsPrimaryID{primaryID})
	query := fmt.Sprintf(financeWorkerRunsQueries.selectFinanceWorkerRuns+" WHERE "+whereQry, defaultFinanceWorkerRunsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeWorkerRuns, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeWorkerRuns with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceWorkerRunsByID] failed get financeWorkerRuns")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceWorkerRunsByID(ctx context.Context, primaryID model.FinanceWorkerRunsPrimaryID, financeWorkerRuns *model.FinanceWorkerRuns, financeWorkerRunsUpdateFields ...FinanceWorkerRunsUpdateField) (err error) {
	exists, err := repo.IsExistFinanceWorkerRunsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceWorkerRuns] failed checking financeWorkerRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeWorkerRuns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeWorkerRuns == nil {
		if len(financeWorkerRunsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceWorkerRunsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeWorkerRuns = &model.FinanceWorkerRuns{}
	}
	var (
		defaultFinanceWorkerRunsUpdateFields = defaultFinanceWorkerRunsUpdateFields(*financeWorkerRuns)
		tempUpdateField                      FinanceWorkerRunsUpdateFieldList
		selectFields                         = NewFinanceWorkerRunsSelectFields()
	)
	if len(financeWorkerRunsUpdateFields) > 0 {
		for _, updateField := range financeWorkerRunsUpdateFields {
			if updateField.financeWorkerRunsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceWorkerRunsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceWorkerRunsCompositePrimaryKeyWhere([]model.FinanceWorkerRunsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceWorkerRunsCommand(defaultFinanceWorkerRunsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeWorkerRunsQueries.updateFinanceWorkerRuns+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceWorkerRuns] error when try to update financeWorkerRuns by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceWorkerRunsByFilter(ctx context.Context, filter model.Filter, financeWorkerRunsUpdateFields ...FinanceWorkerRunsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeWorkerRunsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceWorkerRunsUpdateFieldList
		selectFields = NewFinanceWorkerRunsSelectFields()
	)
	for _, updateField := range financeWorkerRunsUpdateFields {
		if updateField.financeWorkerRunsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceWorkerRunsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeWorkerRunsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceWorkerRunsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_worker_runs\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceWorkerRunsByFilter] error when try to update financeWorkerRuns by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceWorkerRunsByFilter] failed get rows affected")
	}
	return
}

var (
	financeWorkerRunsQueries = struct {
		selectFinanceWorkerRuns      string
		selectCountFinanceWorkerRuns string
		deleteFinanceWorkerRuns      string
		updateFinanceWorkerRuns      string
		insertFinanceWorkerRuns      string
	}{
		selectFinanceWorkerRuns:      "SELECT %s FROM \"finance_worker_runs\"",
		selectCountFinanceWorkerRuns: "SELECT COUNT(\"id\") FROM \"finance_worker_runs\"",
		deleteFinanceWorkerRuns:      "DELETE FROM \"finance_worker_runs\"",
		updateFinanceWorkerRuns:      "UPDATE \"finance_worker_runs\" SET %s ",
		insertFinanceWorkerRuns:      "INSERT INTO \"finance_worker_runs\" %s VALUES %s",
	}
)

type FinanceWorkerRunsRepository interface {
	CreateFinanceWorkerRuns(ctx context.Context, financeWorkerRuns *model.FinanceWorkerRuns, fieldsInsert ...FinanceWorkerRunsField) error
	BulkCreateFinanceWorkerRuns(ctx context.Context, financeWorkerRunsList []*model.FinanceWorkerRuns, fieldsInsert ...FinanceWorkerRunsField) error
	ResolveFinanceWorkerRuns(ctx context.Context, selectFields ...FinanceWorkerRunsField) (model.FinanceWorkerRunsList, error)
	ResolveFinanceWorkerRunsByID(ctx context.Context, primaryID model.FinanceWorkerRunsPrimaryID, selectFields ...FinanceWorkerRunsField) (model.FinanceWorkerRuns, error)
	UpdateFinanceWorkerRunsByID(ctx context.Context, id model.FinanceWorkerRunsPrimaryID, financeWorkerRuns *model.FinanceWorkerRuns, financeWorkerRunsUpdateFields ...FinanceWorkerRunsUpdateField) error
	UpdateFinanceWorkerRunsByFilter(ctx context.Context, filter model.Filter, financeWorkerRunsUpdateFields ...FinanceWorkerRunsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceWorkerRuns(ctx context.Context, financeWorkerRunsListMap map[model.FinanceWorkerRunsPrimaryID]*model.FinanceWorkerRuns, FinanceWorkerRunssMapUpdateFieldsRequest map[model.FinanceWorkerRunsPrimaryID]FinanceWorkerRunsUpdateFieldList) (err error)
	DeleteFinanceWorkerRunsByID(ctx context.Context, id model.FinanceWorkerRunsPrimaryID) error
	BulkDeleteFinanceWorkerRunsByIDs(ctx context.Context, ids []model.FinanceWorkerRunsPrimaryID) error
	ResolveFinanceWorkerRunsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceWorkerRunsFilterResult, err error)
	IsExistFinanceWorkerRunsByIDs(ctx context.Context, ids []model.FinanceWorkerRunsPrimaryID) (exists bool, notFoundIds []model.FinanceWorkerRunsPrimaryID, err error)
	IsExistFinanceWorkerRunsByID(ctx context.Context, id model.FinanceWorkerRunsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
