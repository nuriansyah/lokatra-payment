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

func composeInsertFieldsAndParamsReconciliationMatches(reconciliationMatchesList []model.ReconciliationMatches, fieldsInsert ...ReconciliationMatchesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewReconciliationMatchesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, reconciliationMatches := range reconciliationMatchesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, reconciliationMatches.Id)
			case selectField.ReconciliationRunId():
				args = append(args, reconciliationMatches.ReconciliationRunId)
			case selectField.LedgerJournalId():
				args = append(args, reconciliationMatches.LedgerJournalId)
			case selectField.StatementLineId():
				args = append(args, reconciliationMatches.StatementLineId)
			case selectField.MatchType():
				args = append(args, reconciliationMatches.MatchType)
			case selectField.MatchStatus():
				args = append(args, reconciliationMatches.MatchStatus)
			case selectField.AmountDifference():
				args = append(args, reconciliationMatches.AmountDifference)
			case selectField.MatchedAt():
				args = append(args, reconciliationMatches.MatchedAt)
			case selectField.Metadata():
				args = append(args, reconciliationMatches.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, reconciliationMatches.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, reconciliationMatches.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, reconciliationMatches.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, reconciliationMatches.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, reconciliationMatches.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, reconciliationMatches.MetaDeletedBy)

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

func composeReconciliationMatchesCompositePrimaryKeyWhere(primaryIDs []model.ReconciliationMatchesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"reconciliation_matches\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultReconciliationMatchesSelectFields() string {
	fields := NewReconciliationMatchesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeReconciliationMatchesSelectFields(selectFields ...ReconciliationMatchesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ReconciliationMatchesField string
type ReconciliationMatchesFieldList []ReconciliationMatchesField

type ReconciliationMatchesSelectFields struct {
}

func (ss ReconciliationMatchesSelectFields) Id() ReconciliationMatchesField {
	return ReconciliationMatchesField("id")
}

func (ss ReconciliationMatchesSelectFields) ReconciliationRunId() ReconciliationMatchesField {
	return ReconciliationMatchesField("reconciliation_run_id")
}

func (ss ReconciliationMatchesSelectFields) LedgerJournalId() ReconciliationMatchesField {
	return ReconciliationMatchesField("ledger_journal_id")
}

func (ss ReconciliationMatchesSelectFields) StatementLineId() ReconciliationMatchesField {
	return ReconciliationMatchesField("statement_line_id")
}

func (ss ReconciliationMatchesSelectFields) MatchType() ReconciliationMatchesField {
	return ReconciliationMatchesField("match_type")
}

func (ss ReconciliationMatchesSelectFields) MatchStatus() ReconciliationMatchesField {
	return ReconciliationMatchesField("match_status")
}

func (ss ReconciliationMatchesSelectFields) AmountDifference() ReconciliationMatchesField {
	return ReconciliationMatchesField("amount_difference")
}

func (ss ReconciliationMatchesSelectFields) MatchedAt() ReconciliationMatchesField {
	return ReconciliationMatchesField("matched_at")
}

func (ss ReconciliationMatchesSelectFields) Metadata() ReconciliationMatchesField {
	return ReconciliationMatchesField("metadata")
}

func (ss ReconciliationMatchesSelectFields) MetaCreatedAt() ReconciliationMatchesField {
	return ReconciliationMatchesField("meta_created_at")
}

func (ss ReconciliationMatchesSelectFields) MetaCreatedBy() ReconciliationMatchesField {
	return ReconciliationMatchesField("meta_created_by")
}

func (ss ReconciliationMatchesSelectFields) MetaUpdatedAt() ReconciliationMatchesField {
	return ReconciliationMatchesField("meta_updated_at")
}

func (ss ReconciliationMatchesSelectFields) MetaUpdatedBy() ReconciliationMatchesField {
	return ReconciliationMatchesField("meta_updated_by")
}

func (ss ReconciliationMatchesSelectFields) MetaDeletedAt() ReconciliationMatchesField {
	return ReconciliationMatchesField("meta_deleted_at")
}

func (ss ReconciliationMatchesSelectFields) MetaDeletedBy() ReconciliationMatchesField {
	return ReconciliationMatchesField("meta_deleted_by")
}

func (ss ReconciliationMatchesSelectFields) All() ReconciliationMatchesFieldList {
	return []ReconciliationMatchesField{
		ss.Id(),
		ss.ReconciliationRunId(),
		ss.LedgerJournalId(),
		ss.StatementLineId(),
		ss.MatchType(),
		ss.MatchStatus(),
		ss.AmountDifference(),
		ss.MatchedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewReconciliationMatchesSelectFields() ReconciliationMatchesSelectFields {
	return ReconciliationMatchesSelectFields{}
}

type ReconciliationMatchesUpdateFieldOption struct {
	useIncrement bool
}
type ReconciliationMatchesUpdateField struct {
	reconciliationMatchesField ReconciliationMatchesField
	opt                        ReconciliationMatchesUpdateFieldOption
	value                      interface{}
}
type ReconciliationMatchesUpdateFieldList []ReconciliationMatchesUpdateField

func defaultReconciliationMatchesUpdateFieldOption() ReconciliationMatchesUpdateFieldOption {
	return ReconciliationMatchesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementReconciliationMatchesOption(useIncrement bool) func(*ReconciliationMatchesUpdateFieldOption) {
	return func(pcufo *ReconciliationMatchesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewReconciliationMatchesUpdateField(field ReconciliationMatchesField, val interface{}, opts ...func(*ReconciliationMatchesUpdateFieldOption)) ReconciliationMatchesUpdateField {
	defaultOpt := defaultReconciliationMatchesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ReconciliationMatchesUpdateField{
		reconciliationMatchesField: field,
		value:                      val,
		opt:                        defaultOpt,
	}
}
func defaultReconciliationMatchesUpdateFields(reconciliationMatches model.ReconciliationMatches) (reconciliationMatchesUpdateFieldList ReconciliationMatchesUpdateFieldList) {
	selectFields := NewReconciliationMatchesSelectFields()
	reconciliationMatchesUpdateFieldList = append(reconciliationMatchesUpdateFieldList,
		NewReconciliationMatchesUpdateField(selectFields.Id(), reconciliationMatches.Id),
		NewReconciliationMatchesUpdateField(selectFields.ReconciliationRunId(), reconciliationMatches.ReconciliationRunId),
		NewReconciliationMatchesUpdateField(selectFields.LedgerJournalId(), reconciliationMatches.LedgerJournalId),
		NewReconciliationMatchesUpdateField(selectFields.StatementLineId(), reconciliationMatches.StatementLineId),
		NewReconciliationMatchesUpdateField(selectFields.MatchType(), reconciliationMatches.MatchType),
		NewReconciliationMatchesUpdateField(selectFields.MatchStatus(), reconciliationMatches.MatchStatus),
		NewReconciliationMatchesUpdateField(selectFields.AmountDifference(), reconciliationMatches.AmountDifference),
		NewReconciliationMatchesUpdateField(selectFields.MatchedAt(), reconciliationMatches.MatchedAt),
		NewReconciliationMatchesUpdateField(selectFields.Metadata(), reconciliationMatches.Metadata),
		NewReconciliationMatchesUpdateField(selectFields.MetaCreatedAt(), reconciliationMatches.MetaCreatedAt),
		NewReconciliationMatchesUpdateField(selectFields.MetaCreatedBy(), reconciliationMatches.MetaCreatedBy),
		NewReconciliationMatchesUpdateField(selectFields.MetaUpdatedAt(), reconciliationMatches.MetaUpdatedAt),
		NewReconciliationMatchesUpdateField(selectFields.MetaUpdatedBy(), reconciliationMatches.MetaUpdatedBy),
		NewReconciliationMatchesUpdateField(selectFields.MetaDeletedAt(), reconciliationMatches.MetaDeletedAt),
		NewReconciliationMatchesUpdateField(selectFields.MetaDeletedBy(), reconciliationMatches.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsReconciliationMatchesCommand(reconciliationMatchesUpdateFieldList ReconciliationMatchesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range reconciliationMatchesUpdateFieldList {
		field := string(updateField.reconciliationMatchesField)
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

func (repo *RepositoryImpl) BulkCreateReconciliationMatches(ctx context.Context, reconciliationMatchesList []*model.ReconciliationMatches, fieldsInsert ...ReconciliationMatchesField) (err error) {
	var (
		fieldsStr                      string
		valueListStr                   []string
		argsList                       []interface{}
		primaryIds                     []model.ReconciliationMatchesPrimaryID
		reconciliationMatchesValueList []model.ReconciliationMatches
	)

	if len(fieldsInsert) == 0 {
		selectField := NewReconciliationMatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, reconciliationMatches := range reconciliationMatchesList {

		primaryIds = append(primaryIds, reconciliationMatches.ToReconciliationMatchesPrimaryID())

		reconciliationMatchesValueList = append(reconciliationMatchesValueList, *reconciliationMatches)
	}

	_, notFoundIds, err := repo.IsExistReconciliationMatchesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReconciliationMatches] failed checking reconciliationMatches whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ReconciliationMatchesPrimaryID{}
		mapNotFoundIds := map[model.ReconciliationMatchesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "reconciliationMatches", fmt.Sprintf("reconciliationMatches with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsReconciliationMatches(reconciliationMatchesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(reconciliationMatchesQueries.insertReconciliationMatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReconciliationMatches] failed exec create reconciliationMatches query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteReconciliationMatchesByIDs(ctx context.Context, primaryIDs []model.ReconciliationMatchesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistReconciliationMatchesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationMatchesByIDs] failed checking reconciliationMatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationMatches with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reconciliation_matches\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(reconciliationMatchesQueries.deleteReconciliationMatches + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationMatchesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReconciliationMatchesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistReconciliationMatchesByIDs(ctx context.Context, ids []model.ReconciliationMatchesPrimaryID) (exists bool, notFoundIds []model.ReconciliationMatchesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reconciliation_matches\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(reconciliationMatchesQueries.selectReconciliationMatches, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationMatchesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ReconciliationMatchesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationMatchesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ReconciliationMatchesPrimaryID]bool{}
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

// BulkUpdateReconciliationMatches is used to bulk update reconciliationMatches, by default it will update all field
// if want to update specific field, then fill reconciliationMatchessMapUpdateFieldsRequest else please fill reconciliationMatchessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateReconciliationMatches(ctx context.Context, reconciliationMatchessMap map[model.ReconciliationMatchesPrimaryID]*model.ReconciliationMatches, reconciliationMatchessMapUpdateFieldsRequest map[model.ReconciliationMatchesPrimaryID]ReconciliationMatchesUpdateFieldList) (err error) {
	if len(reconciliationMatchessMap) == 0 && len(reconciliationMatchessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		reconciliationMatchessMapUpdateField map[model.ReconciliationMatchesPrimaryID]ReconciliationMatchesUpdateFieldList = map[model.ReconciliationMatchesPrimaryID]ReconciliationMatchesUpdateFieldList{}
		asTableValues                        string                                                                        = "myvalues"
	)

	if len(reconciliationMatchessMap) > 0 {
		for id, reconciliationMatches := range reconciliationMatchessMap {
			if reconciliationMatches == nil {
				log.Error().Err(err).Msg("[BulkUpdateReconciliationMatches] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			reconciliationMatchessMapUpdateField[id] = defaultReconciliationMatchesUpdateFields(*reconciliationMatches)
		}
	} else {
		reconciliationMatchessMapUpdateField = reconciliationMatchessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateReconciliationMatchesQuery(reconciliationMatchessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistReconciliationMatchesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReconciliationMatches] failed checking reconciliationMatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationMatches with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeReconciliationMatchesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"reconciliation_matches\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReconciliationMatches] failed exec query")
	}
	return
}

type ReconciliationMatchesFieldParameter struct {
	param string
	args  []interface{}
}

func NewReconciliationMatchesFieldParameter(param string, args ...interface{}) ReconciliationMatchesFieldParameter {
	return ReconciliationMatchesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateReconciliationMatchesQuery(mapReconciliationMatchess map[model.ReconciliationMatchesPrimaryID]ReconciliationMatchesUpdateFieldList, asTableValues string) (primaryIDs []model.ReconciliationMatchesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ReconciliationMatchesPrimaryID]map[string]interface{}{}
	reconciliationMatchesSelectFields := NewReconciliationMatchesSelectFields()
	for id, updateFields := range mapReconciliationMatchess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.reconciliationMatchesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapReconciliationMatchess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetReconciliationMatchesFieldType(updateField.reconciliationMatchesField)))
			args = append(args, fields[string(updateField.reconciliationMatchesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.reconciliationMatchesField))
		if updateField.reconciliationMatchesField == reconciliationMatchesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.reconciliationMatchesField, asTableValues, updateField.reconciliationMatchesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.reconciliationMatchesField,
				"\"reconciliation_matches\"", updateField.reconciliationMatchesField,
				asTableValues, updateField.reconciliationMatchesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeReconciliationMatchesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ReconciliationMatchesPrimaryID, asTableValue string) (whereQry string) {
	reconciliationMatchesSelectFields := NewReconciliationMatchesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"reconciliation_matches\".\"id\" = %s.\"id\"::"+GetReconciliationMatchesFieldType(reconciliationMatchesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetReconciliationMatchesFieldType(reconciliationMatchesField ReconciliationMatchesField) string {
	selectReconciliationMatchesFields := NewReconciliationMatchesSelectFields()
	switch reconciliationMatchesField {

	case selectReconciliationMatchesFields.Id():
		return "uuid"

	case selectReconciliationMatchesFields.ReconciliationRunId():
		return "uuid"

	case selectReconciliationMatchesFields.LedgerJournalId():
		return "uuid"

	case selectReconciliationMatchesFields.StatementLineId():
		return "uuid"

	case selectReconciliationMatchesFields.MatchType():
		return "match_type_enum"

	case selectReconciliationMatchesFields.MatchStatus():
		return "match_status_enum"

	case selectReconciliationMatchesFields.AmountDifference():
		return "numeric"

	case selectReconciliationMatchesFields.MatchedAt():
		return "timestamptz"

	case selectReconciliationMatchesFields.Metadata():
		return "jsonb"

	case selectReconciliationMatchesFields.MetaCreatedAt():
		return "timestamptz"

	case selectReconciliationMatchesFields.MetaCreatedBy():
		return "uuid"

	case selectReconciliationMatchesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectReconciliationMatchesFields.MetaUpdatedBy():
		return "uuid"

	case selectReconciliationMatchesFields.MetaDeletedAt():
		return "timestamptz"

	case selectReconciliationMatchesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateReconciliationMatches(ctx context.Context, reconciliationMatches *model.ReconciliationMatches, fieldsInsert ...ReconciliationMatchesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewReconciliationMatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ReconciliationMatchesPrimaryID{
		Id: reconciliationMatches.Id,
	}
	exists, err := repo.IsExistReconciliationMatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReconciliationMatches] failed checking reconciliationMatches whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "reconciliationMatches", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsReconciliationMatches([]model.ReconciliationMatches{*reconciliationMatches}, fieldsInsert...)
	commandQuery := fmt.Sprintf(reconciliationMatchesQueries.insertReconciliationMatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReconciliationMatches] failed exec create reconciliationMatches query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteReconciliationMatchesByID(ctx context.Context, primaryID model.ReconciliationMatchesPrimaryID) (err error) {
	exists, err := repo.IsExistReconciliationMatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReconciliationMatchesByID] failed checking reconciliationMatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationMatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeReconciliationMatchesCompositePrimaryKeyWhere([]model.ReconciliationMatchesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(reconciliationMatchesQueries.deleteReconciliationMatches + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReconciliationMatchesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationMatchesByFilter(ctx context.Context, filter model.Filter) (result []model.ReconciliationMatchesFilterResult, err error) {
	query, args, err := composeReconciliationMatchesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationMatchesByFilter] failed compose reconciliationMatches filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationMatchesByFilter] failed get reconciliationMatches by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeReconciliationMatchesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ReconciliationMatchesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeReconciliationMatchesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeReconciliationMatchesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeReconciliationMatchesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewReconciliationMatchesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["ledger_journal_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"ledger_journal_id\"")
			selectedColumns["ledger_journal_id"] = struct{}{}
		}
		if _, selected := selectedColumns["statement_line_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"statement_line_id\"")
			selectedColumns["statement_line_id"] = struct{}{}
		}
		if _, selected := selectedColumns["match_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"match_type\"")
			selectedColumns["match_type"] = struct{}{}
		}
		if _, selected := selectedColumns["match_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"match_status\"")
			selectedColumns["match_status"] = struct{}{}
		}
		if _, selected := selectedColumns["amount_difference"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount_difference\"")
			selectedColumns["amount_difference"] = struct{}{}
		}
		if _, selected := selectedColumns["matched_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"matched_at\"")
			selectedColumns["matched_at"] = struct{}{}
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

type reconciliationMatchesFilterPlaceholder struct {
	index int
}

func (p *reconciliationMatchesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeReconciliationMatchesFilterPredicate(filterField model.FilterField, placeholders *reconciliationMatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewReconciliationMatchesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeReconciliationMatchesFilterSQLExpr(spec)
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

func composeReconciliationMatchesFilterGroup(group model.FilterGroup, placeholders *reconciliationMatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeReconciliationMatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeReconciliationMatchesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeReconciliationMatchesFilterWhereQueries(filter model.Filter, placeholders *reconciliationMatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeReconciliationMatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeReconciliationMatchesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeReconciliationMatchesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateReconciliationMatchesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeReconciliationMatchesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeReconciliationMatchesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := reconciliationMatchesFilterPlaceholder{index: 1}
	whereQueries, err := composeReconciliationMatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewReconciliationMatchesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeReconciliationMatchesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeReconciliationMatchesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"reconciliation_matches\" base%s", strings.Join(selectColumns, ","), composeReconciliationMatchesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistReconciliationMatchesByID(ctx context.Context, primaryID model.ReconciliationMatchesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeReconciliationMatchesCompositePrimaryKeyWhere([]model.ReconciliationMatchesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", reconciliationMatchesQueries.selectCountReconciliationMatches, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReconciliationMatchesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationMatches(ctx context.Context, selectFields ...ReconciliationMatchesField) (reconciliationMatchesList model.ReconciliationMatchesList, err error) {
	var (
		defaultReconciliationMatchesSelectFields = defaultReconciliationMatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReconciliationMatchesSelectFields = composeReconciliationMatchesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(reconciliationMatchesQueries.selectReconciliationMatches, defaultReconciliationMatchesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &reconciliationMatchesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReconciliationMatches] failed get reconciliationMatches list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReconciliationMatchesByID(ctx context.Context, primaryID model.ReconciliationMatchesPrimaryID, selectFields ...ReconciliationMatchesField) (reconciliationMatches model.ReconciliationMatches, err error) {
	var (
		defaultReconciliationMatchesSelectFields = defaultReconciliationMatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReconciliationMatchesSelectFields = composeReconciliationMatchesSelectFields(selectFields...)
	}
	whereQry, params := composeReconciliationMatchesCompositePrimaryKeyWhere([]model.ReconciliationMatchesPrimaryID{primaryID})
	query := fmt.Sprintf(reconciliationMatchesQueries.selectReconciliationMatches+" WHERE "+whereQry, defaultReconciliationMatchesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &reconciliationMatches, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("reconciliationMatches with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveReconciliationMatchesByID] failed get reconciliationMatches")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateReconciliationMatchesByID(ctx context.Context, primaryID model.ReconciliationMatchesPrimaryID, reconciliationMatches *model.ReconciliationMatches, reconciliationMatchesUpdateFields ...ReconciliationMatchesUpdateField) (err error) {
	exists, err := repo.IsExistReconciliationMatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationMatches] failed checking reconciliationMatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reconciliationMatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if reconciliationMatches == nil {
		if len(reconciliationMatchesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateReconciliationMatchesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		reconciliationMatches = &model.ReconciliationMatches{}
	}
	var (
		defaultReconciliationMatchesUpdateFields = defaultReconciliationMatchesUpdateFields(*reconciliationMatches)
		tempUpdateField                          ReconciliationMatchesUpdateFieldList
		selectFields                             = NewReconciliationMatchesSelectFields()
	)
	if len(reconciliationMatchesUpdateFields) > 0 {
		for _, updateField := range reconciliationMatchesUpdateFields {
			if updateField.reconciliationMatchesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultReconciliationMatchesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeReconciliationMatchesCompositePrimaryKeyWhere([]model.ReconciliationMatchesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsReconciliationMatchesCommand(defaultReconciliationMatchesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(reconciliationMatchesQueries.updateReconciliationMatches+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationMatches] error when try to update reconciliationMatches by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateReconciliationMatchesByFilter(ctx context.Context, filter model.Filter, reconciliationMatchesUpdateFields ...ReconciliationMatchesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(reconciliationMatchesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ReconciliationMatchesUpdateFieldList
		selectFields = NewReconciliationMatchesSelectFields()
	)
	for _, updateField := range reconciliationMatchesUpdateFields {
		if updateField.reconciliationMatchesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsReconciliationMatchesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := reconciliationMatchesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeReconciliationMatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"reconciliation_matches\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationMatchesByFilter] error when try to update reconciliationMatches by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReconciliationMatchesByFilter] failed get rows affected")
	}
	return
}

var (
	reconciliationMatchesQueries = struct {
		selectReconciliationMatches      string
		selectCountReconciliationMatches string
		deleteReconciliationMatches      string
		updateReconciliationMatches      string
		insertReconciliationMatches      string
	}{
		selectReconciliationMatches:      "SELECT %s FROM \"reconciliation_matches\"",
		selectCountReconciliationMatches: "SELECT COUNT(\"id\") FROM \"reconciliation_matches\"",
		deleteReconciliationMatches:      "DELETE FROM \"reconciliation_matches\"",
		updateReconciliationMatches:      "UPDATE \"reconciliation_matches\" SET %s ",
		insertReconciliationMatches:      "INSERT INTO \"reconciliation_matches\" %s VALUES %s",
	}
)

type ReconciliationMatchesRepository interface {
	CreateReconciliationMatches(ctx context.Context, reconciliationMatches *model.ReconciliationMatches, fieldsInsert ...ReconciliationMatchesField) error
	BulkCreateReconciliationMatches(ctx context.Context, reconciliationMatchesList []*model.ReconciliationMatches, fieldsInsert ...ReconciliationMatchesField) error
	ResolveReconciliationMatches(ctx context.Context, selectFields ...ReconciliationMatchesField) (model.ReconciliationMatchesList, error)
	ResolveReconciliationMatchesByID(ctx context.Context, primaryID model.ReconciliationMatchesPrimaryID, selectFields ...ReconciliationMatchesField) (model.ReconciliationMatches, error)
	UpdateReconciliationMatchesByID(ctx context.Context, id model.ReconciliationMatchesPrimaryID, reconciliationMatches *model.ReconciliationMatches, reconciliationMatchesUpdateFields ...ReconciliationMatchesUpdateField) error
	UpdateReconciliationMatchesByFilter(ctx context.Context, filter model.Filter, reconciliationMatchesUpdateFields ...ReconciliationMatchesUpdateField) (rowsAffected int64, err error)
	BulkUpdateReconciliationMatches(ctx context.Context, reconciliationMatchesListMap map[model.ReconciliationMatchesPrimaryID]*model.ReconciliationMatches, ReconciliationMatchessMapUpdateFieldsRequest map[model.ReconciliationMatchesPrimaryID]ReconciliationMatchesUpdateFieldList) (err error)
	DeleteReconciliationMatchesByID(ctx context.Context, id model.ReconciliationMatchesPrimaryID) error
	BulkDeleteReconciliationMatchesByIDs(ctx context.Context, ids []model.ReconciliationMatchesPrimaryID) error
	ResolveReconciliationMatchesByFilter(ctx context.Context, filter model.Filter) (result []model.ReconciliationMatchesFilterResult, err error)
	IsExistReconciliationMatchesByIDs(ctx context.Context, ids []model.ReconciliationMatchesPrimaryID) (exists bool, notFoundIds []model.ReconciliationMatchesPrimaryID, err error)
	IsExistReconciliationMatchesByID(ctx context.Context, id model.ReconciliationMatchesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
