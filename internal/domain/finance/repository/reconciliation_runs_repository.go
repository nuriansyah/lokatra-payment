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

func composeInsertFieldsAndParamsReconciliationRuns(reconciliationRunsList []model.ReconciliationRuns, fieldsInsert ...ReconciliationRunsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewReconciliationRunsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, reconciliationRuns := range reconciliationRunsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, reconciliationRuns.Id)
			case selectField.ReconCode():
				args = append(args, reconciliationRuns.ReconCode)
			case selectField.ReconType():
				args = append(args, reconciliationRuns.ReconType)
			case selectField.PeriodStart():
				args = append(args, reconciliationRuns.PeriodStart)
			case selectField.PeriodEnd():
				args = append(args, reconciliationRuns.PeriodEnd)
			case selectField.CurrencyCode():
				args = append(args, reconciliationRuns.CurrencyCode)
			case selectField.RunStatus():
				args = append(args, reconciliationRuns.RunStatus)
			case selectField.ToleranceAmount():
				args = append(args, reconciliationRuns.ToleranceAmount)
			case selectField.ToleranceDays():
				args = append(args, reconciliationRuns.ToleranceDays)
			case selectField.StartedAt():
				args = append(args, reconciliationRuns.StartedAt)
			case selectField.CompletedAt():
				args = append(args, reconciliationRuns.CompletedAt)
			case selectField.Summary():
				args = append(args, reconciliationRuns.Summary)
			case selectField.MetaCreatedAt():
				args = append(args, reconciliationRuns.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, reconciliationRuns.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, reconciliationRuns.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, reconciliationRuns.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, reconciliationRuns.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, reconciliationRuns.MetaDeletedBy)

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

func composeReconciliationRunsCompositePrimaryKeyWhere(primaryIDs []model.ReconciliationRunsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"reconciliation_runs\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultReconciliationRunsSelectFields() string {
	fields := NewReconciliationRunsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeReconciliationRunsSelectFields(selectFields ...ReconciliationRunsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ReconciliationRunsField string
type ReconciliationRunsFieldList []ReconciliationRunsField

type ReconciliationRunsSelectFields struct {
}

func (ss ReconciliationRunsSelectFields) Id() ReconciliationRunsField {
	return ReconciliationRunsField("id")
}

func (ss ReconciliationRunsSelectFields) ReconCode() ReconciliationRunsField {
	return ReconciliationRunsField("recon_code")
}

func (ss ReconciliationRunsSelectFields) ReconType() ReconciliationRunsField {
	return ReconciliationRunsField("recon_type")
}

func (ss ReconciliationRunsSelectFields) PeriodStart() ReconciliationRunsField {
	return ReconciliationRunsField("period_start")
}

func (ss ReconciliationRunsSelectFields) PeriodEnd() ReconciliationRunsField {
	return ReconciliationRunsField("period_end")
}

func (ss ReconciliationRunsSelectFields) CurrencyCode() ReconciliationRunsField {
	return ReconciliationRunsField("currency_code")
}

func (ss ReconciliationRunsSelectFields) RunStatus() ReconciliationRunsField {
	return ReconciliationRunsField("run_status")
}

func (ss ReconciliationRunsSelectFields) ToleranceAmount() ReconciliationRunsField {
	return ReconciliationRunsField("tolerance_amount")
}

func (ss ReconciliationRunsSelectFields) ToleranceDays() ReconciliationRunsField {
	return ReconciliationRunsField("tolerance_days")
}

func (ss ReconciliationRunsSelectFields) StartedAt() ReconciliationRunsField {
	return ReconciliationRunsField("started_at")
}

func (ss ReconciliationRunsSelectFields) CompletedAt() ReconciliationRunsField {
	return ReconciliationRunsField("completed_at")
}

func (ss ReconciliationRunsSelectFields) Summary() ReconciliationRunsField {
	return ReconciliationRunsField("summary")
}

func (ss ReconciliationRunsSelectFields) MetaCreatedAt() ReconciliationRunsField {
	return ReconciliationRunsField("meta_created_at")
}

func (ss ReconciliationRunsSelectFields) MetaCreatedBy() ReconciliationRunsField {
	return ReconciliationRunsField("meta_created_by")
}

func (ss ReconciliationRunsSelectFields) MetaUpdatedAt() ReconciliationRunsField {
	return ReconciliationRunsField("meta_updated_at")
}

func (ss ReconciliationRunsSelectFields) MetaUpdatedBy() ReconciliationRunsField {
	return ReconciliationRunsField("meta_updated_by")
}

func (ss ReconciliationRunsSelectFields) MetaDeletedAt() ReconciliationRunsField {
	return ReconciliationRunsField("meta_deleted_at")
}

func (ss ReconciliationRunsSelectFields) MetaDeletedBy() ReconciliationRunsField {
	return ReconciliationRunsField("meta_deleted_by")
}

func (ss ReconciliationRunsSelectFields) All() ReconciliationRunsFieldList {
	return []ReconciliationRunsField{
		ss.Id(),
		ss.ReconCode(),
		ss.ReconType(),
		ss.PeriodStart(),
		ss.PeriodEnd(),
		ss.CurrencyCode(),
		ss.RunStatus(),
		ss.ToleranceAmount(),
		ss.ToleranceDays(),
		ss.StartedAt(),
		ss.CompletedAt(),
		ss.Summary(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewReconciliationRunsSelectFields() ReconciliationRunsSelectFields {
	return ReconciliationRunsSelectFields{}
}

type ReconciliationRunsUpdateFieldOption struct {
	useIncrement bool
}
type ReconciliationRunsUpdateField struct {
	reconciliationRunsField ReconciliationRunsField
	opt                     ReconciliationRunsUpdateFieldOption
	value                   interface{}
}
type ReconciliationRunsUpdateFieldList []ReconciliationRunsUpdateField

func defaultReconciliationRunsUpdateFieldOption() ReconciliationRunsUpdateFieldOption {
	return ReconciliationRunsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementReconciliationRunsOption(useIncrement bool) func(*ReconciliationRunsUpdateFieldOption) {
	return func(pcufo *ReconciliationRunsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewReconciliationRunsUpdateField(field ReconciliationRunsField, val interface{}, opts ...func(*ReconciliationRunsUpdateFieldOption)) ReconciliationRunsUpdateField {
	defaultOpt := defaultReconciliationRunsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ReconciliationRunsUpdateField{
		reconciliationRunsField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultReconciliationRunsUpdateFields(reconciliationRuns model.ReconciliationRuns) (reconciliationRunsUpdateFieldList ReconciliationRunsUpdateFieldList) {
	selectFields := NewReconciliationRunsSelectFields()
	reconciliationRunsUpdateFieldList = append(reconciliationRunsUpdateFieldList,
		NewReconciliationRunsUpdateField(selectFields.Id(), reconciliationRuns.Id),
		NewReconciliationRunsUpdateField(selectFields.ReconCode(), reconciliationRuns.ReconCode),
		NewReconciliationRunsUpdateField(selectFields.ReconType(), reconciliationRuns.ReconType),
		NewReconciliationRunsUpdateField(selectFields.PeriodStart(), reconciliationRuns.PeriodStart),
		NewReconciliationRunsUpdateField(selectFields.PeriodEnd(), reconciliationRuns.PeriodEnd),
		NewReconciliationRunsUpdateField(selectFields.CurrencyCode(), reconciliationRuns.CurrencyCode),
		NewReconciliationRunsUpdateField(selectFields.RunStatus(), reconciliationRuns.RunStatus),
		NewReconciliationRunsUpdateField(selectFields.ToleranceAmount(), reconciliationRuns.ToleranceAmount),
		NewReconciliationRunsUpdateField(selectFields.ToleranceDays(), reconciliationRuns.ToleranceDays),
		NewReconciliationRunsUpdateField(selectFields.StartedAt(), reconciliationRuns.StartedAt),
		NewReconciliationRunsUpdateField(selectFields.CompletedAt(), reconciliationRuns.CompletedAt),
		NewReconciliationRunsUpdateField(selectFields.Summary(), reconciliationRuns.Summary),
		NewReconciliationRunsUpdateField(selectFields.MetaCreatedAt(), reconciliationRuns.MetaCreatedAt),
		NewReconciliationRunsUpdateField(selectFields.MetaCreatedBy(), reconciliationRuns.MetaCreatedBy),
		NewReconciliationRunsUpdateField(selectFields.MetaUpdatedAt(), reconciliationRuns.MetaUpdatedAt),
		NewReconciliationRunsUpdateField(selectFields.MetaUpdatedBy(), reconciliationRuns.MetaUpdatedBy),
		NewReconciliationRunsUpdateField(selectFields.MetaDeletedAt(), reconciliationRuns.MetaDeletedAt),
		NewReconciliationRunsUpdateField(selectFields.MetaDeletedBy(), reconciliationRuns.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsReconciliationRunsCommand(reconciliationRunsUpdateFieldList ReconciliationRunsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range reconciliationRunsUpdateFieldList {
		field := string(updateField.reconciliationRunsField)
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

func (repo *RepositoryImpl) BulkCreateReconciliationRuns(ctx context.Context, reconciliationRunsList []*model.ReconciliationRuns, fieldsInsert ...ReconciliationRunsField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.ReconciliationRunsPrimaryID
		reconciliationRunsValueList []model.ReconciliationRuns
	)

	if len(fieldsInsert) == 0 {
		selectField := NewReconciliationRunsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, reconciliationRuns := range reconciliationRunsList {

		primaryIds = append(primaryIds, reconciliationRuns.ToReconciliationRunsPrimaryID())

		reconciliationRunsValueList = append(reconciliationRunsValueList, *reconciliationRuns)
	}

	_, notFoundIds, err := repo.IsExistReconciliationRunsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReconciliationRuns] failed checking reconciliationRuns whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ReconciliationRunsPrimaryID{}
		mapNotFoundIds := map[model.ReconciliationRunsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "reconciliationRuns", fmt.Sprintf("reconciliationRuns with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsReconciliationRuns(reconciliationRunsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(reconciliationRunsQueries.insertReconciliationRuns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReconciliationRuns] failed exec create reconciliationRuns query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteReconciliationRunsByIDs(ctx context.Context, primaryIDs []model.ReconciliationRunsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistReconciliationRunsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationRunsByIDs] failed checking reconciliationRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationRuns with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reconciliation_runs\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(reconciliationRunsQueries.deleteReconciliationRuns + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationRunsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationRunsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistReconciliationRunsByIDs(ctx context.Context, ids []model.ReconciliationRunsPrimaryID) (exists bool, notFoundIds []model.ReconciliationRunsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reconciliation_runs\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(reconciliationRunsQueries.selectReconciliationRuns, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationRunsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ReconciliationRunsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationRunsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ReconciliationRunsPrimaryID]bool{}
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

// BulkUpdateReconciliationRuns is used to bulk update reconciliationRuns, by default it will update all field
// if want to update specific field, then fill reconciliationRunssMapUpdateFieldsRequest else please fill reconciliationRunssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateReconciliationRuns(ctx context.Context, reconciliationRunssMap map[model.ReconciliationRunsPrimaryID]*model.ReconciliationRuns, reconciliationRunssMapUpdateFieldsRequest map[model.ReconciliationRunsPrimaryID]ReconciliationRunsUpdateFieldList) (err error) {
	if len(reconciliationRunssMap) == 0 && len(reconciliationRunssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		reconciliationRunssMapUpdateField map[model.ReconciliationRunsPrimaryID]ReconciliationRunsUpdateFieldList = map[model.ReconciliationRunsPrimaryID]ReconciliationRunsUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(reconciliationRunssMap) > 0 {
		for id, reconciliationRuns := range reconciliationRunssMap {
			if reconciliationRuns == nil {
				log.Error().Err(err).Msg("[BulkUpdateReconciliationRuns] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			reconciliationRunssMapUpdateField[id] = defaultReconciliationRunsUpdateFields(*reconciliationRuns)
		}
	} else {
		reconciliationRunssMapUpdateField = reconciliationRunssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateReconciliationRunsQuery(reconciliationRunssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistReconciliationRunsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReconciliationRuns] failed checking reconciliationRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationRuns with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeReconciliationRunsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"reconciliation_runs\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReconciliationRuns] failed exec query")
	}
	return
}

type ReconciliationRunsFieldParameter struct {
	param string
	args  []interface{}
}

func NewReconciliationRunsFieldParameter(param string, args ...interface{}) ReconciliationRunsFieldParameter {
	return ReconciliationRunsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateReconciliationRunsQuery(mapReconciliationRunss map[model.ReconciliationRunsPrimaryID]ReconciliationRunsUpdateFieldList, asTableValues string) (primaryIDs []model.ReconciliationRunsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ReconciliationRunsPrimaryID]map[string]interface{}{}
	reconciliationRunsSelectFields := NewReconciliationRunsSelectFields()
	for id, updateFields := range mapReconciliationRunss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.reconciliationRunsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapReconciliationRunss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetReconciliationRunsFieldType(updateField.reconciliationRunsField)))
			args = append(args, fields[string(updateField.reconciliationRunsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.reconciliationRunsField))
		if updateField.reconciliationRunsField == reconciliationRunsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.reconciliationRunsField, asTableValues, updateField.reconciliationRunsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.reconciliationRunsField,
				"\"reconciliation_runs\"", updateField.reconciliationRunsField,
				asTableValues, updateField.reconciliationRunsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeReconciliationRunsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ReconciliationRunsPrimaryID, asTableValue string) (whereQry string) {
	reconciliationRunsSelectFields := NewReconciliationRunsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"reconciliation_runs\".\"id\" = %s.\"id\"::"+GetReconciliationRunsFieldType(reconciliationRunsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetReconciliationRunsFieldType(reconciliationRunsField ReconciliationRunsField) string {
	selectReconciliationRunsFields := NewReconciliationRunsSelectFields()
	switch reconciliationRunsField {

	case selectReconciliationRunsFields.Id():
		return "uuid"

	case selectReconciliationRunsFields.ReconCode():
		return "text"

	case selectReconciliationRunsFields.ReconType():
		return "recon_type_enum"

	case selectReconciliationRunsFields.PeriodStart():
		return "timestamptz"

	case selectReconciliationRunsFields.PeriodEnd():
		return "timestamptz"

	case selectReconciliationRunsFields.CurrencyCode():
		return "text"

	case selectReconciliationRunsFields.RunStatus():
		return "reconciliation_runs_run_status_enum"

	case selectReconciliationRunsFields.ToleranceAmount():
		return "numeric"

	case selectReconciliationRunsFields.ToleranceDays():
		return "int4"

	case selectReconciliationRunsFields.StartedAt():
		return "timestamptz"

	case selectReconciliationRunsFields.CompletedAt():
		return "timestamptz"

	case selectReconciliationRunsFields.Summary():
		return "jsonb"

	case selectReconciliationRunsFields.MetaCreatedAt():
		return "timestamptz"

	case selectReconciliationRunsFields.MetaCreatedBy():
		return "uuid"

	case selectReconciliationRunsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectReconciliationRunsFields.MetaUpdatedBy():
		return "uuid"

	case selectReconciliationRunsFields.MetaDeletedAt():
		return "timestamptz"

	case selectReconciliationRunsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateReconciliationRuns(ctx context.Context, reconciliationRuns *model.ReconciliationRuns, fieldsInsert ...ReconciliationRunsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewReconciliationRunsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ReconciliationRunsPrimaryID{
		Id: reconciliationRuns.Id,
	}
	exists, err := repo.IsExistReconciliationRunsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReconciliationRuns] failed checking reconciliationRuns whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "reconciliationRuns", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsReconciliationRuns([]model.ReconciliationRuns{*reconciliationRuns}, fieldsInsert...)
	commandQuery := fmt.Sprintf(reconciliationRunsQueries.insertReconciliationRuns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReconciliationRuns] failed exec create reconciliationRuns query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteReconciliationRunsByID(ctx context.Context, primaryID model.ReconciliationRunsPrimaryID) (err error) {
	exists, err := repo.IsExistReconciliationRunsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReconciliationRunsByID] failed checking reconciliationRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationRuns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeReconciliationRunsCompositePrimaryKeyWhere([]model.ReconciliationRunsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(reconciliationRunsQueries.deleteReconciliationRuns + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReconciliationRunsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationRunsByFilter(ctx context.Context, filter model.Filter) (result []model.ReconciliationRunsFilterResult, err error) {
	query, args, err := composeReconciliationRunsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationRunsByFilter] failed compose reconciliationRuns filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationRunsByFilter] failed get reconciliationRuns by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeReconciliationRunsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ReconciliationRunsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeReconciliationRunsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeReconciliationRunsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeReconciliationRunsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewReconciliationRunsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 18+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["recon_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"recon_code\"")
			selectedColumns["recon_code"] = struct{}{}
		}
		if _, selected := selectedColumns["recon_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"recon_type\"")
			selectedColumns["recon_type"] = struct{}{}
		}
		if _, selected := selectedColumns["period_start"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_start\"")
			selectedColumns["period_start"] = struct{}{}
		}
		if _, selected := selectedColumns["period_end"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_end\"")
			selectedColumns["period_end"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["run_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"run_status\"")
			selectedColumns["run_status"] = struct{}{}
		}
		if _, selected := selectedColumns["tolerance_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"tolerance_amount\"")
			selectedColumns["tolerance_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["tolerance_days"]; !selected {
			selectColumns = append(selectColumns, "base.\"tolerance_days\"")
			selectedColumns["tolerance_days"] = struct{}{}
		}
		if _, selected := selectedColumns["started_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"started_at\"")
			selectedColumns["started_at"] = struct{}{}
		}
		if _, selected := selectedColumns["completed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"completed_at\"")
			selectedColumns["completed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["summary"]; !selected {
			selectColumns = append(selectColumns, "base.\"summary\"")
			selectedColumns["summary"] = struct{}{}
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

type reconciliationRunsFilterPlaceholder struct {
	index int
}

func (p *reconciliationRunsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeReconciliationRunsFilterPredicate(filterField model.FilterField, placeholders *reconciliationRunsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewReconciliationRunsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeReconciliationRunsFilterSQLExpr(spec)
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

func composeReconciliationRunsFilterGroup(group model.FilterGroup, placeholders *reconciliationRunsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeReconciliationRunsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeReconciliationRunsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeReconciliationRunsFilterWhereQueries(filter model.Filter, placeholders *reconciliationRunsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeReconciliationRunsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeReconciliationRunsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeReconciliationRunsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateReconciliationRunsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeReconciliationRunsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeReconciliationRunsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := reconciliationRunsFilterPlaceholder{index: 1}
	whereQueries, err := composeReconciliationRunsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewReconciliationRunsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeReconciliationRunsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeReconciliationRunsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"reconciliation_runs\" base%s", strings.Join(selectColumns, ","), composeReconciliationRunsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistReconciliationRunsByID(ctx context.Context, primaryID model.ReconciliationRunsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeReconciliationRunsCompositePrimaryKeyWhere([]model.ReconciliationRunsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", reconciliationRunsQueries.selectCountReconciliationRuns, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationRunsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationRuns(ctx context.Context, selectFields ...ReconciliationRunsField) (reconciliationRunsList model.ReconciliationRunsList, err error) {
	var (
		defaultReconciliationRunsSelectFields = defaultReconciliationRunsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReconciliationRunsSelectFields = composeReconciliationRunsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(reconciliationRunsQueries.selectReconciliationRuns, defaultReconciliationRunsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &reconciliationRunsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationRuns] failed get reconciliationRuns list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationRunsByID(ctx context.Context, primaryID model.ReconciliationRunsPrimaryID, selectFields ...ReconciliationRunsField) (reconciliationRuns model.ReconciliationRuns, err error) {
	var (
		defaultReconciliationRunsSelectFields = defaultReconciliationRunsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReconciliationRunsSelectFields = composeReconciliationRunsSelectFields(selectFields...)
	}
	whereQry, params := composeReconciliationRunsCompositePrimaryKeyWhere([]model.ReconciliationRunsPrimaryID{primaryID})
	query := fmt.Sprintf(reconciliationRunsQueries.selectReconciliationRuns+" WHERE "+whereQry, defaultReconciliationRunsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &reconciliationRuns, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("reconciliationRuns with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveReconciliationRunsByID] failed get reconciliationRuns")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateReconciliationRunsByID(ctx context.Context, primaryID model.ReconciliationRunsPrimaryID, reconciliationRuns *model.ReconciliationRuns, reconciliationRunsUpdateFields ...ReconciliationRunsUpdateField) (err error) {
	exists, err := repo.IsExistReconciliationRunsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationRuns] failed checking reconciliationRuns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationRuns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if reconciliationRuns == nil {
		if len(reconciliationRunsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateReconciliationRunsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		reconciliationRuns = &model.ReconciliationRuns{}
	}
	var (
		defaultReconciliationRunsUpdateFields = defaultReconciliationRunsUpdateFields(*reconciliationRuns)
		tempUpdateField                       ReconciliationRunsUpdateFieldList
		selectFields                          = NewReconciliationRunsSelectFields()
	)
	if len(reconciliationRunsUpdateFields) > 0 {
		for _, updateField := range reconciliationRunsUpdateFields {
			if updateField.reconciliationRunsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultReconciliationRunsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeReconciliationRunsCompositePrimaryKeyWhere([]model.ReconciliationRunsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsReconciliationRunsCommand(defaultReconciliationRunsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(reconciliationRunsQueries.updateReconciliationRuns+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationRuns] error when try to update reconciliationRuns by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateReconciliationRunsByFilter(ctx context.Context, filter model.Filter, reconciliationRunsUpdateFields ...ReconciliationRunsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(reconciliationRunsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ReconciliationRunsUpdateFieldList
		selectFields = NewReconciliationRunsSelectFields()
	)
	for _, updateField := range reconciliationRunsUpdateFields {
		if updateField.reconciliationRunsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsReconciliationRunsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := reconciliationRunsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeReconciliationRunsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"reconciliation_runs\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationRunsByFilter] error when try to update reconciliationRuns by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationRunsByFilter] failed get rows affected")
	}
	return
}

var (
	reconciliationRunsQueries = struct {
		selectReconciliationRuns      string
		selectCountReconciliationRuns string
		deleteReconciliationRuns      string
		updateReconciliationRuns      string
		insertReconciliationRuns      string
	}{
		selectReconciliationRuns:      "SELECT %s FROM \"reconciliation_runs\"",
		selectCountReconciliationRuns: "SELECT COUNT(\"id\") FROM \"reconciliation_runs\"",
		deleteReconciliationRuns:      "DELETE FROM \"reconciliation_runs\"",
		updateReconciliationRuns:      "UPDATE \"reconciliation_runs\" SET %s ",
		insertReconciliationRuns:      "INSERT INTO \"reconciliation_runs\" %s VALUES %s",
	}
)

type ReconciliationRunsRepository interface {
	CreateReconciliationRuns(ctx context.Context, reconciliationRuns *model.ReconciliationRuns, fieldsInsert ...ReconciliationRunsField) error
	BulkCreateReconciliationRuns(ctx context.Context, reconciliationRunsList []*model.ReconciliationRuns, fieldsInsert ...ReconciliationRunsField) error
	ResolveReconciliationRuns(ctx context.Context, selectFields ...ReconciliationRunsField) (model.ReconciliationRunsList, error)
	ResolveReconciliationRunsByID(ctx context.Context, primaryID model.ReconciliationRunsPrimaryID, selectFields ...ReconciliationRunsField) (model.ReconciliationRuns, error)
	UpdateReconciliationRunsByID(ctx context.Context, id model.ReconciliationRunsPrimaryID, reconciliationRuns *model.ReconciliationRuns, reconciliationRunsUpdateFields ...ReconciliationRunsUpdateField) error
	UpdateReconciliationRunsByFilter(ctx context.Context, filter model.Filter, reconciliationRunsUpdateFields ...ReconciliationRunsUpdateField) (rowsAffected int64, err error)
	BulkUpdateReconciliationRuns(ctx context.Context, reconciliationRunsListMap map[model.ReconciliationRunsPrimaryID]*model.ReconciliationRuns, ReconciliationRunssMapUpdateFieldsRequest map[model.ReconciliationRunsPrimaryID]ReconciliationRunsUpdateFieldList) (err error)
	DeleteReconciliationRunsByID(ctx context.Context, id model.ReconciliationRunsPrimaryID) error
	BulkDeleteReconciliationRunsByIDs(ctx context.Context, ids []model.ReconciliationRunsPrimaryID) error
	ResolveReconciliationRunsByFilter(ctx context.Context, filter model.Filter) (result []model.ReconciliationRunsFilterResult, err error)
	IsExistReconciliationRunsByIDs(ctx context.Context, ids []model.ReconciliationRunsPrimaryID) (exists bool, notFoundIds []model.ReconciliationRunsPrimaryID, err error)
	IsExistReconciliationRunsByID(ctx context.Context, id model.ReconciliationRunsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
