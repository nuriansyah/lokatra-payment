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

func composeInsertFieldsAndParamsReconciliationCandidates(reconciliationCandidatesList []model.ReconciliationCandidates, fieldsInsert ...ReconciliationCandidatesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewReconciliationCandidatesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, reconciliationCandidates := range reconciliationCandidatesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, reconciliationCandidates.Id)
			case selectField.ReconciliationRunId():
				args = append(args, reconciliationCandidates.ReconciliationRunId)
			case selectField.SourceSystem():
				args = append(args, reconciliationCandidates.SourceSystem)
			case selectField.SourceRefId():
				args = append(args, reconciliationCandidates.SourceRefId)
			case selectField.CandidateKey():
				args = append(args, reconciliationCandidates.CandidateKey)
			case selectField.Amount():
				args = append(args, reconciliationCandidates.Amount)
			case selectField.OccurredAt():
				args = append(args, reconciliationCandidates.OccurredAt)
			case selectField.NormalizedPayload():
				args = append(args, reconciliationCandidates.NormalizedPayload)
			case selectField.Metadata():
				args = append(args, reconciliationCandidates.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, reconciliationCandidates.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, reconciliationCandidates.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, reconciliationCandidates.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, reconciliationCandidates.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, reconciliationCandidates.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, reconciliationCandidates.MetaDeletedBy)

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

func composeReconciliationCandidatesCompositePrimaryKeyWhere(primaryIDs []model.ReconciliationCandidatesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"reconciliation_candidates\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultReconciliationCandidatesSelectFields() string {
	fields := NewReconciliationCandidatesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeReconciliationCandidatesSelectFields(selectFields ...ReconciliationCandidatesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ReconciliationCandidatesField string
type ReconciliationCandidatesFieldList []ReconciliationCandidatesField

type ReconciliationCandidatesSelectFields struct {
}

func (ss ReconciliationCandidatesSelectFields) Id() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("id")
}

func (ss ReconciliationCandidatesSelectFields) ReconciliationRunId() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("reconciliation_run_id")
}

func (ss ReconciliationCandidatesSelectFields) SourceSystem() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("source_system")
}

func (ss ReconciliationCandidatesSelectFields) SourceRefId() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("source_ref_id")
}

func (ss ReconciliationCandidatesSelectFields) CandidateKey() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("candidate_key")
}

func (ss ReconciliationCandidatesSelectFields) Amount() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("amount")
}

func (ss ReconciliationCandidatesSelectFields) OccurredAt() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("occurred_at")
}

func (ss ReconciliationCandidatesSelectFields) NormalizedPayload() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("normalized_payload")
}

func (ss ReconciliationCandidatesSelectFields) Metadata() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("metadata")
}

func (ss ReconciliationCandidatesSelectFields) MetaCreatedAt() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("meta_created_at")
}

func (ss ReconciliationCandidatesSelectFields) MetaCreatedBy() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("meta_created_by")
}

func (ss ReconciliationCandidatesSelectFields) MetaUpdatedAt() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("meta_updated_at")
}

func (ss ReconciliationCandidatesSelectFields) MetaUpdatedBy() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("meta_updated_by")
}

func (ss ReconciliationCandidatesSelectFields) MetaDeletedAt() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("meta_deleted_at")
}

func (ss ReconciliationCandidatesSelectFields) MetaDeletedBy() ReconciliationCandidatesField {
	return ReconciliationCandidatesField("meta_deleted_by")
}

func (ss ReconciliationCandidatesSelectFields) All() ReconciliationCandidatesFieldList {
	return []ReconciliationCandidatesField{
		ss.Id(),
		ss.ReconciliationRunId(),
		ss.SourceSystem(),
		ss.SourceRefId(),
		ss.CandidateKey(),
		ss.Amount(),
		ss.OccurredAt(),
		ss.NormalizedPayload(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewReconciliationCandidatesSelectFields() ReconciliationCandidatesSelectFields {
	return ReconciliationCandidatesSelectFields{}
}

type ReconciliationCandidatesUpdateFieldOption struct {
	useIncrement bool
}
type ReconciliationCandidatesUpdateField struct {
	reconciliationCandidatesField ReconciliationCandidatesField
	opt                           ReconciliationCandidatesUpdateFieldOption
	value                         interface{}
}
type ReconciliationCandidatesUpdateFieldList []ReconciliationCandidatesUpdateField

func defaultReconciliationCandidatesUpdateFieldOption() ReconciliationCandidatesUpdateFieldOption {
	return ReconciliationCandidatesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementReconciliationCandidatesOption(useIncrement bool) func(*ReconciliationCandidatesUpdateFieldOption) {
	return func(pcufo *ReconciliationCandidatesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewReconciliationCandidatesUpdateField(field ReconciliationCandidatesField, val interface{}, opts ...func(*ReconciliationCandidatesUpdateFieldOption)) ReconciliationCandidatesUpdateField {
	defaultOpt := defaultReconciliationCandidatesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ReconciliationCandidatesUpdateField{
		reconciliationCandidatesField: field,
		value:                         val,
		opt:                           defaultOpt,
	}
}
func defaultReconciliationCandidatesUpdateFields(reconciliationCandidates model.ReconciliationCandidates) (reconciliationCandidatesUpdateFieldList ReconciliationCandidatesUpdateFieldList) {
	selectFields := NewReconciliationCandidatesSelectFields()
	reconciliationCandidatesUpdateFieldList = append(reconciliationCandidatesUpdateFieldList,
		NewReconciliationCandidatesUpdateField(selectFields.Id(), reconciliationCandidates.Id),
		NewReconciliationCandidatesUpdateField(selectFields.ReconciliationRunId(), reconciliationCandidates.ReconciliationRunId),
		NewReconciliationCandidatesUpdateField(selectFields.SourceSystem(), reconciliationCandidates.SourceSystem),
		NewReconciliationCandidatesUpdateField(selectFields.SourceRefId(), reconciliationCandidates.SourceRefId),
		NewReconciliationCandidatesUpdateField(selectFields.CandidateKey(), reconciliationCandidates.CandidateKey),
		NewReconciliationCandidatesUpdateField(selectFields.Amount(), reconciliationCandidates.Amount),
		NewReconciliationCandidatesUpdateField(selectFields.OccurredAt(), reconciliationCandidates.OccurredAt),
		NewReconciliationCandidatesUpdateField(selectFields.NormalizedPayload(), reconciliationCandidates.NormalizedPayload),
		NewReconciliationCandidatesUpdateField(selectFields.Metadata(), reconciliationCandidates.Metadata),
		NewReconciliationCandidatesUpdateField(selectFields.MetaCreatedAt(), reconciliationCandidates.MetaCreatedAt),
		NewReconciliationCandidatesUpdateField(selectFields.MetaCreatedBy(), reconciliationCandidates.MetaCreatedBy),
		NewReconciliationCandidatesUpdateField(selectFields.MetaUpdatedAt(), reconciliationCandidates.MetaUpdatedAt),
		NewReconciliationCandidatesUpdateField(selectFields.MetaUpdatedBy(), reconciliationCandidates.MetaUpdatedBy),
		NewReconciliationCandidatesUpdateField(selectFields.MetaDeletedAt(), reconciliationCandidates.MetaDeletedAt),
		NewReconciliationCandidatesUpdateField(selectFields.MetaDeletedBy(), reconciliationCandidates.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsReconciliationCandidatesCommand(reconciliationCandidatesUpdateFieldList ReconciliationCandidatesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range reconciliationCandidatesUpdateFieldList {
		field := string(updateField.reconciliationCandidatesField)
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

func (repo *RepositoryImpl) BulkCreateReconciliationCandidates(ctx context.Context, reconciliationCandidatesList []*model.ReconciliationCandidates, fieldsInsert ...ReconciliationCandidatesField) (err error) {
	var (
		fieldsStr                         string
		valueListStr                      []string
		argsList                          []interface{}
		primaryIds                        []model.ReconciliationCandidatesPrimaryID
		reconciliationCandidatesValueList []model.ReconciliationCandidates
	)

	if len(fieldsInsert) == 0 {
		selectField := NewReconciliationCandidatesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, reconciliationCandidates := range reconciliationCandidatesList {

		primaryIds = append(primaryIds, reconciliationCandidates.ToReconciliationCandidatesPrimaryID())

		reconciliationCandidatesValueList = append(reconciliationCandidatesValueList, *reconciliationCandidates)
	}

	_, notFoundIds, err := repo.IsExistReconciliationCandidatesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReconciliationCandidates] failed checking reconciliationCandidates whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ReconciliationCandidatesPrimaryID{}
		mapNotFoundIds := map[model.ReconciliationCandidatesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "reconciliationCandidates", fmt.Sprintf("reconciliationCandidates with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsReconciliationCandidates(reconciliationCandidatesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(reconciliationCandidatesQueries.insertReconciliationCandidates, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReconciliationCandidates] failed exec create reconciliationCandidates query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteReconciliationCandidatesByIDs(ctx context.Context, primaryIDs []model.ReconciliationCandidatesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistReconciliationCandidatesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationCandidatesByIDs] failed checking reconciliationCandidates whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationCandidates with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reconciliation_candidates\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(reconciliationCandidatesQueries.deleteReconciliationCandidates + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationCandidatesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationCandidatesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistReconciliationCandidatesByIDs(ctx context.Context, ids []model.ReconciliationCandidatesPrimaryID) (exists bool, notFoundIds []model.ReconciliationCandidatesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reconciliation_candidates\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(reconciliationCandidatesQueries.selectReconciliationCandidates, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationCandidatesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ReconciliationCandidatesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationCandidatesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ReconciliationCandidatesPrimaryID]bool{}
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

// BulkUpdateReconciliationCandidates is used to bulk update reconciliationCandidates, by default it will update all field
// if want to update specific field, then fill reconciliationCandidatessMapUpdateFieldsRequest else please fill reconciliationCandidatessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateReconciliationCandidates(ctx context.Context, reconciliationCandidatessMap map[model.ReconciliationCandidatesPrimaryID]*model.ReconciliationCandidates, reconciliationCandidatessMapUpdateFieldsRequest map[model.ReconciliationCandidatesPrimaryID]ReconciliationCandidatesUpdateFieldList) (err error) {
	if len(reconciliationCandidatessMap) == 0 && len(reconciliationCandidatessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		reconciliationCandidatessMapUpdateField map[model.ReconciliationCandidatesPrimaryID]ReconciliationCandidatesUpdateFieldList = map[model.ReconciliationCandidatesPrimaryID]ReconciliationCandidatesUpdateFieldList{}
		asTableValues                           string                                                                              = "myvalues"
	)

	if len(reconciliationCandidatessMap) > 0 {
		for id, reconciliationCandidates := range reconciliationCandidatessMap {
			if reconciliationCandidates == nil {
				log.Error().Err(err).Msg("[BulkUpdateReconciliationCandidates] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			reconciliationCandidatessMapUpdateField[id] = defaultReconciliationCandidatesUpdateFields(*reconciliationCandidates)
		}
	} else {
		reconciliationCandidatessMapUpdateField = reconciliationCandidatessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateReconciliationCandidatesQuery(reconciliationCandidatessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistReconciliationCandidatesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReconciliationCandidates] failed checking reconciliationCandidates whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationCandidates with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeReconciliationCandidatesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"reconciliation_candidates\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReconciliationCandidates] failed exec query")
	}
	return
}

type ReconciliationCandidatesFieldParameter struct {
	param string
	args  []interface{}
}

func NewReconciliationCandidatesFieldParameter(param string, args ...interface{}) ReconciliationCandidatesFieldParameter {
	return ReconciliationCandidatesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateReconciliationCandidatesQuery(mapReconciliationCandidatess map[model.ReconciliationCandidatesPrimaryID]ReconciliationCandidatesUpdateFieldList, asTableValues string) (primaryIDs []model.ReconciliationCandidatesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ReconciliationCandidatesPrimaryID]map[string]interface{}{}
	reconciliationCandidatesSelectFields := NewReconciliationCandidatesSelectFields()
	for id, updateFields := range mapReconciliationCandidatess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.reconciliationCandidatesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapReconciliationCandidatess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetReconciliationCandidatesFieldType(updateField.reconciliationCandidatesField)))
			args = append(args, fields[string(updateField.reconciliationCandidatesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.reconciliationCandidatesField))
		if updateField.reconciliationCandidatesField == reconciliationCandidatesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.reconciliationCandidatesField, asTableValues, updateField.reconciliationCandidatesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.reconciliationCandidatesField,
				"\"reconciliation_candidates\"", updateField.reconciliationCandidatesField,
				asTableValues, updateField.reconciliationCandidatesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeReconciliationCandidatesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ReconciliationCandidatesPrimaryID, asTableValue string) (whereQry string) {
	reconciliationCandidatesSelectFields := NewReconciliationCandidatesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"reconciliation_candidates\".\"id\" = %s.\"id\"::"+GetReconciliationCandidatesFieldType(reconciliationCandidatesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetReconciliationCandidatesFieldType(reconciliationCandidatesField ReconciliationCandidatesField) string {
	selectReconciliationCandidatesFields := NewReconciliationCandidatesSelectFields()
	switch reconciliationCandidatesField {

	case selectReconciliationCandidatesFields.Id():
		return "uuid"

	case selectReconciliationCandidatesFields.ReconciliationRunId():
		return "uuid"

	case selectReconciliationCandidatesFields.SourceSystem():
		return "source_system_enum"

	case selectReconciliationCandidatesFields.SourceRefId():
		return "uuid"

	case selectReconciliationCandidatesFields.CandidateKey():
		return "text"

	case selectReconciliationCandidatesFields.Amount():
		return "numeric"

	case selectReconciliationCandidatesFields.OccurredAt():
		return "timestamptz"

	case selectReconciliationCandidatesFields.NormalizedPayload():
		return "jsonb"

	case selectReconciliationCandidatesFields.Metadata():
		return "jsonb"

	case selectReconciliationCandidatesFields.MetaCreatedAt():
		return "timestamptz"

	case selectReconciliationCandidatesFields.MetaCreatedBy():
		return "uuid"

	case selectReconciliationCandidatesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectReconciliationCandidatesFields.MetaUpdatedBy():
		return "uuid"

	case selectReconciliationCandidatesFields.MetaDeletedAt():
		return "timestamptz"

	case selectReconciliationCandidatesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateReconciliationCandidates(ctx context.Context, reconciliationCandidates *model.ReconciliationCandidates, fieldsInsert ...ReconciliationCandidatesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewReconciliationCandidatesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ReconciliationCandidatesPrimaryID{
		Id: reconciliationCandidates.Id,
	}
	exists, err := repo.IsExistReconciliationCandidatesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReconciliationCandidates] failed checking reconciliationCandidates whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "reconciliationCandidates", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsReconciliationCandidates([]model.ReconciliationCandidates{*reconciliationCandidates}, fieldsInsert...)
	commandQuery := fmt.Sprintf(reconciliationCandidatesQueries.insertReconciliationCandidates, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReconciliationCandidates] failed exec create reconciliationCandidates query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteReconciliationCandidatesByID(ctx context.Context, primaryID model.ReconciliationCandidatesPrimaryID) (err error) {
	exists, err := repo.IsExistReconciliationCandidatesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReconciliationCandidatesByID] failed checking reconciliationCandidates whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationCandidates with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeReconciliationCandidatesCompositePrimaryKeyWhere([]model.ReconciliationCandidatesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(reconciliationCandidatesQueries.deleteReconciliationCandidates + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReconciliationCandidatesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationCandidatesByFilter(ctx context.Context, filter model.Filter) (result []model.ReconciliationCandidatesFilterResult, err error) {
	query, args, err := composeReconciliationCandidatesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationCandidatesByFilter] failed compose reconciliationCandidates filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationCandidatesByFilter] failed get reconciliationCandidates by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeReconciliationCandidatesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ReconciliationCandidatesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeReconciliationCandidatesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeReconciliationCandidatesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeReconciliationCandidatesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewReconciliationCandidatesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["reconciliation_run_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"reconciliation_run_id\"")
			selectedColumns["reconciliation_run_id"] = struct{}{}
		}
		if _, selected := selectedColumns["source_system"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_system\"")
			selectedColumns["source_system"] = struct{}{}
		}
		if _, selected := selectedColumns["source_ref_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_ref_id\"")
			selectedColumns["source_ref_id"] = struct{}{}
		}
		if _, selected := selectedColumns["candidate_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"candidate_key\"")
			selectedColumns["candidate_key"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["occurred_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"occurred_at\"")
			selectedColumns["occurred_at"] = struct{}{}
		}
		if _, selected := selectedColumns["normalized_payload"]; !selected {
			selectColumns = append(selectColumns, "base.\"normalized_payload\"")
			selectedColumns["normalized_payload"] = struct{}{}
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

type reconciliationCandidatesFilterPlaceholder struct {
	index int
}

func (p *reconciliationCandidatesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeReconciliationCandidatesFilterPredicate(filterField model.FilterField, placeholders *reconciliationCandidatesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewReconciliationCandidatesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeReconciliationCandidatesFilterSQLExpr(spec)
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

func composeReconciliationCandidatesFilterGroup(group model.FilterGroup, placeholders *reconciliationCandidatesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeReconciliationCandidatesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeReconciliationCandidatesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeReconciliationCandidatesFilterWhereQueries(filter model.Filter, placeholders *reconciliationCandidatesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeReconciliationCandidatesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeReconciliationCandidatesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeReconciliationCandidatesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateReconciliationCandidatesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeReconciliationCandidatesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeReconciliationCandidatesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := reconciliationCandidatesFilterPlaceholder{index: 1}
	whereQueries, err := composeReconciliationCandidatesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewReconciliationCandidatesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeReconciliationCandidatesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeReconciliationCandidatesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"reconciliation_candidates\" base%s", strings.Join(selectColumns, ","), composeReconciliationCandidatesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistReconciliationCandidatesByID(ctx context.Context, primaryID model.ReconciliationCandidatesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeReconciliationCandidatesCompositePrimaryKeyWhere([]model.ReconciliationCandidatesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", reconciliationCandidatesQueries.selectCountReconciliationCandidates, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationCandidatesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationCandidates(ctx context.Context, selectFields ...ReconciliationCandidatesField) (reconciliationCandidatesList model.ReconciliationCandidatesList, err error) {
	var (
		defaultReconciliationCandidatesSelectFields = defaultReconciliationCandidatesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReconciliationCandidatesSelectFields = composeReconciliationCandidatesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(reconciliationCandidatesQueries.selectReconciliationCandidates, defaultReconciliationCandidatesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &reconciliationCandidatesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationCandidates] failed get reconciliationCandidates list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationCandidatesByID(ctx context.Context, primaryID model.ReconciliationCandidatesPrimaryID, selectFields ...ReconciliationCandidatesField) (reconciliationCandidates model.ReconciliationCandidates, err error) {
	var (
		defaultReconciliationCandidatesSelectFields = defaultReconciliationCandidatesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReconciliationCandidatesSelectFields = composeReconciliationCandidatesSelectFields(selectFields...)
	}
	whereQry, params := composeReconciliationCandidatesCompositePrimaryKeyWhere([]model.ReconciliationCandidatesPrimaryID{primaryID})
	query := fmt.Sprintf(reconciliationCandidatesQueries.selectReconciliationCandidates+" WHERE "+whereQry, defaultReconciliationCandidatesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &reconciliationCandidates, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("reconciliationCandidates with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveReconciliationCandidatesByID] failed get reconciliationCandidates")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateReconciliationCandidatesByID(ctx context.Context, primaryID model.ReconciliationCandidatesPrimaryID, reconciliationCandidates *model.ReconciliationCandidates, reconciliationCandidatesUpdateFields ...ReconciliationCandidatesUpdateField) (err error) {
	exists, err := repo.IsExistReconciliationCandidatesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationCandidates] failed checking reconciliationCandidates whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationCandidates with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if reconciliationCandidates == nil {
		if len(reconciliationCandidatesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateReconciliationCandidatesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		reconciliationCandidates = &model.ReconciliationCandidates{}
	}
	var (
		defaultReconciliationCandidatesUpdateFields = defaultReconciliationCandidatesUpdateFields(*reconciliationCandidates)
		tempUpdateField                             ReconciliationCandidatesUpdateFieldList
		selectFields                                = NewReconciliationCandidatesSelectFields()
	)
	if len(reconciliationCandidatesUpdateFields) > 0 {
		for _, updateField := range reconciliationCandidatesUpdateFields {
			if updateField.reconciliationCandidatesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultReconciliationCandidatesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeReconciliationCandidatesCompositePrimaryKeyWhere([]model.ReconciliationCandidatesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsReconciliationCandidatesCommand(defaultReconciliationCandidatesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(reconciliationCandidatesQueries.updateReconciliationCandidates+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationCandidates] error when try to update reconciliationCandidates by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateReconciliationCandidatesByFilter(ctx context.Context, filter model.Filter, reconciliationCandidatesUpdateFields ...ReconciliationCandidatesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(reconciliationCandidatesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ReconciliationCandidatesUpdateFieldList
		selectFields = NewReconciliationCandidatesSelectFields()
	)
	for _, updateField := range reconciliationCandidatesUpdateFields {
		if updateField.reconciliationCandidatesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsReconciliationCandidatesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := reconciliationCandidatesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeReconciliationCandidatesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"reconciliation_candidates\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationCandidatesByFilter] error when try to update reconciliationCandidates by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationCandidatesByFilter] failed get rows affected")
	}
	return
}

var (
	reconciliationCandidatesQueries = struct {
		selectReconciliationCandidates      string
		selectCountReconciliationCandidates string
		deleteReconciliationCandidates      string
		updateReconciliationCandidates      string
		insertReconciliationCandidates      string
	}{
		selectReconciliationCandidates:      "SELECT %s FROM \"reconciliation_candidates\"",
		selectCountReconciliationCandidates: "SELECT COUNT(\"id\") FROM \"reconciliation_candidates\"",
		deleteReconciliationCandidates:      "DELETE FROM \"reconciliation_candidates\"",
		updateReconciliationCandidates:      "UPDATE \"reconciliation_candidates\" SET %s ",
		insertReconciliationCandidates:      "INSERT INTO \"reconciliation_candidates\" %s VALUES %s",
	}
)

type ReconciliationCandidatesRepository interface {
	CreateReconciliationCandidates(ctx context.Context, reconciliationCandidates *model.ReconciliationCandidates, fieldsInsert ...ReconciliationCandidatesField) error
	BulkCreateReconciliationCandidates(ctx context.Context, reconciliationCandidatesList []*model.ReconciliationCandidates, fieldsInsert ...ReconciliationCandidatesField) error
	ResolveReconciliationCandidates(ctx context.Context, selectFields ...ReconciliationCandidatesField) (model.ReconciliationCandidatesList, error)
	ResolveReconciliationCandidatesByID(ctx context.Context, primaryID model.ReconciliationCandidatesPrimaryID, selectFields ...ReconciliationCandidatesField) (model.ReconciliationCandidates, error)
	UpdateReconciliationCandidatesByID(ctx context.Context, id model.ReconciliationCandidatesPrimaryID, reconciliationCandidates *model.ReconciliationCandidates, reconciliationCandidatesUpdateFields ...ReconciliationCandidatesUpdateField) error
	UpdateReconciliationCandidatesByFilter(ctx context.Context, filter model.Filter, reconciliationCandidatesUpdateFields ...ReconciliationCandidatesUpdateField) (rowsAffected int64, err error)
	BulkUpdateReconciliationCandidates(ctx context.Context, reconciliationCandidatesListMap map[model.ReconciliationCandidatesPrimaryID]*model.ReconciliationCandidates, ReconciliationCandidatessMapUpdateFieldsRequest map[model.ReconciliationCandidatesPrimaryID]ReconciliationCandidatesUpdateFieldList) (err error)
	DeleteReconciliationCandidatesByID(ctx context.Context, id model.ReconciliationCandidatesPrimaryID) error
	BulkDeleteReconciliationCandidatesByIDs(ctx context.Context, ids []model.ReconciliationCandidatesPrimaryID) error
	ResolveReconciliationCandidatesByFilter(ctx context.Context, filter model.Filter) (result []model.ReconciliationCandidatesFilterResult, err error)
	IsExistReconciliationCandidatesByIDs(ctx context.Context, ids []model.ReconciliationCandidatesPrimaryID) (exists bool, notFoundIds []model.ReconciliationCandidatesPrimaryID, err error)
	IsExistReconciliationCandidatesByID(ctx context.Context, id model.ReconciliationCandidatesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
