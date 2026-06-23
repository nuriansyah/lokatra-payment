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

func composeInsertFieldsAndParamsFeeRuleSets(feeRuleSetsList []model.FeeRuleSets, fieldsInsert ...FeeRuleSetsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFeeRuleSetsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, feeRuleSets := range feeRuleSetsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, feeRuleSets.Id)
			case selectField.FeeProfileId():
				args = append(args, feeRuleSets.FeeProfileId)
			case selectField.RuleSetCode():
				args = append(args, feeRuleSets.RuleSetCode)
			case selectField.Precedence():
				args = append(args, feeRuleSets.Precedence)
			case selectField.EffectiveFrom():
				args = append(args, feeRuleSets.EffectiveFrom)
			case selectField.EffectiveUntil():
				args = append(args, feeRuleSets.EffectiveUntil)
			case selectField.RuleSetStatus():
				args = append(args, feeRuleSets.RuleSetStatus)
			case selectField.Metadata():
				args = append(args, feeRuleSets.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, feeRuleSets.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, feeRuleSets.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, feeRuleSets.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, feeRuleSets.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, feeRuleSets.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, feeRuleSets.MetaDeletedBy)

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

func composeFeeRuleSetsCompositePrimaryKeyWhere(primaryIDs []model.FeeRuleSetsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"fee_rule_sets\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFeeRuleSetsSelectFields() string {
	fields := NewFeeRuleSetsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFeeRuleSetsSelectFields(selectFields ...FeeRuleSetsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FeeRuleSetsField string
type FeeRuleSetsFieldList []FeeRuleSetsField

type FeeRuleSetsSelectFields struct {
}

func (ss FeeRuleSetsSelectFields) Id() FeeRuleSetsField {
	return FeeRuleSetsField("id")
}

func (ss FeeRuleSetsSelectFields) FeeProfileId() FeeRuleSetsField {
	return FeeRuleSetsField("fee_profile_id")
}

func (ss FeeRuleSetsSelectFields) RuleSetCode() FeeRuleSetsField {
	return FeeRuleSetsField("rule_set_code")
}

func (ss FeeRuleSetsSelectFields) Precedence() FeeRuleSetsField {
	return FeeRuleSetsField("precedence")
}

func (ss FeeRuleSetsSelectFields) EffectiveFrom() FeeRuleSetsField {
	return FeeRuleSetsField("effective_from")
}

func (ss FeeRuleSetsSelectFields) EffectiveUntil() FeeRuleSetsField {
	return FeeRuleSetsField("effective_until")
}

func (ss FeeRuleSetsSelectFields) RuleSetStatus() FeeRuleSetsField {
	return FeeRuleSetsField("rule_set_status")
}

func (ss FeeRuleSetsSelectFields) Metadata() FeeRuleSetsField {
	return FeeRuleSetsField("metadata")
}

func (ss FeeRuleSetsSelectFields) MetaCreatedAt() FeeRuleSetsField {
	return FeeRuleSetsField("meta_created_at")
}

func (ss FeeRuleSetsSelectFields) MetaCreatedBy() FeeRuleSetsField {
	return FeeRuleSetsField("meta_created_by")
}

func (ss FeeRuleSetsSelectFields) MetaUpdatedAt() FeeRuleSetsField {
	return FeeRuleSetsField("meta_updated_at")
}

func (ss FeeRuleSetsSelectFields) MetaUpdatedBy() FeeRuleSetsField {
	return FeeRuleSetsField("meta_updated_by")
}

func (ss FeeRuleSetsSelectFields) MetaDeletedAt() FeeRuleSetsField {
	return FeeRuleSetsField("meta_deleted_at")
}

func (ss FeeRuleSetsSelectFields) MetaDeletedBy() FeeRuleSetsField {
	return FeeRuleSetsField("meta_deleted_by")
}

func (ss FeeRuleSetsSelectFields) All() FeeRuleSetsFieldList {
	return []FeeRuleSetsField{
		ss.Id(),
		ss.FeeProfileId(),
		ss.RuleSetCode(),
		ss.Precedence(),
		ss.EffectiveFrom(),
		ss.EffectiveUntil(),
		ss.RuleSetStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFeeRuleSetsSelectFields() FeeRuleSetsSelectFields {
	return FeeRuleSetsSelectFields{}
}

type FeeRuleSetsUpdateFieldOption struct {
	useIncrement bool
}
type FeeRuleSetsUpdateField struct {
	feeRuleSetsField FeeRuleSetsField
	opt              FeeRuleSetsUpdateFieldOption
	value            interface{}
}
type FeeRuleSetsUpdateFieldList []FeeRuleSetsUpdateField

func defaultFeeRuleSetsUpdateFieldOption() FeeRuleSetsUpdateFieldOption {
	return FeeRuleSetsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFeeRuleSetsOption(useIncrement bool) func(*FeeRuleSetsUpdateFieldOption) {
	return func(pcufo *FeeRuleSetsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFeeRuleSetsUpdateField(field FeeRuleSetsField, val interface{}, opts ...func(*FeeRuleSetsUpdateFieldOption)) FeeRuleSetsUpdateField {
	defaultOpt := defaultFeeRuleSetsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FeeRuleSetsUpdateField{
		feeRuleSetsField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultFeeRuleSetsUpdateFields(feeRuleSets model.FeeRuleSets) (feeRuleSetsUpdateFieldList FeeRuleSetsUpdateFieldList) {
	selectFields := NewFeeRuleSetsSelectFields()
	feeRuleSetsUpdateFieldList = append(feeRuleSetsUpdateFieldList,
		NewFeeRuleSetsUpdateField(selectFields.Id(), feeRuleSets.Id),
		NewFeeRuleSetsUpdateField(selectFields.FeeProfileId(), feeRuleSets.FeeProfileId),
		NewFeeRuleSetsUpdateField(selectFields.RuleSetCode(), feeRuleSets.RuleSetCode),
		NewFeeRuleSetsUpdateField(selectFields.Precedence(), feeRuleSets.Precedence),
		NewFeeRuleSetsUpdateField(selectFields.EffectiveFrom(), feeRuleSets.EffectiveFrom),
		NewFeeRuleSetsUpdateField(selectFields.EffectiveUntil(), feeRuleSets.EffectiveUntil),
		NewFeeRuleSetsUpdateField(selectFields.RuleSetStatus(), feeRuleSets.RuleSetStatus),
		NewFeeRuleSetsUpdateField(selectFields.Metadata(), feeRuleSets.Metadata),
		NewFeeRuleSetsUpdateField(selectFields.MetaCreatedAt(), feeRuleSets.MetaCreatedAt),
		NewFeeRuleSetsUpdateField(selectFields.MetaCreatedBy(), feeRuleSets.MetaCreatedBy),
		NewFeeRuleSetsUpdateField(selectFields.MetaUpdatedAt(), feeRuleSets.MetaUpdatedAt),
		NewFeeRuleSetsUpdateField(selectFields.MetaUpdatedBy(), feeRuleSets.MetaUpdatedBy),
		NewFeeRuleSetsUpdateField(selectFields.MetaDeletedAt(), feeRuleSets.MetaDeletedAt),
		NewFeeRuleSetsUpdateField(selectFields.MetaDeletedBy(), feeRuleSets.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFeeRuleSetsCommand(feeRuleSetsUpdateFieldList FeeRuleSetsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range feeRuleSetsUpdateFieldList {
		field := string(updateField.feeRuleSetsField)
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

func (repo *RepositoryImpl) BulkCreateFeeRuleSets(ctx context.Context, feeRuleSetsList []*model.FeeRuleSets, fieldsInsert ...FeeRuleSetsField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.FeeRuleSetsPrimaryID
		feeRuleSetsValueList []model.FeeRuleSets
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFeeRuleSetsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, feeRuleSets := range feeRuleSetsList {

		primaryIds = append(primaryIds, feeRuleSets.ToFeeRuleSetsPrimaryID())

		feeRuleSetsValueList = append(feeRuleSetsValueList, *feeRuleSets)
	}

	_, notFoundIds, err := repo.IsExistFeeRuleSetsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeRuleSets] failed checking feeRuleSets whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FeeRuleSetsPrimaryID{}
		mapNotFoundIds := map[model.FeeRuleSetsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "feeRuleSets", fmt.Sprintf("feeRuleSets with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFeeRuleSets(feeRuleSetsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(feeRuleSetsQueries.insertFeeRuleSets, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeRuleSets] failed exec create feeRuleSets query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFeeRuleSetsByIDs(ctx context.Context, primaryIDs []model.FeeRuleSetsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFeeRuleSetsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRuleSetsByIDs] failed checking feeRuleSets whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleSets with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_rule_sets\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(feeRuleSetsQueries.deleteFeeRuleSets + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRuleSetsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRuleSetsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFeeRuleSetsByIDs(ctx context.Context, ids []model.FeeRuleSetsPrimaryID) (exists bool, notFoundIds []model.FeeRuleSetsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_rule_sets\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(feeRuleSetsQueries.selectFeeRuleSets, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRuleSetsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FeeRuleSetsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRuleSetsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FeeRuleSetsPrimaryID]bool{}
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

// BulkUpdateFeeRuleSets is used to bulk update feeRuleSets, by default it will update all field
// if want to update specific field, then fill feeRuleSetssMapUpdateFieldsRequest else please fill feeRuleSetssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFeeRuleSets(ctx context.Context, feeRuleSetssMap map[model.FeeRuleSetsPrimaryID]*model.FeeRuleSets, feeRuleSetssMapUpdateFieldsRequest map[model.FeeRuleSetsPrimaryID]FeeRuleSetsUpdateFieldList) (err error) {
	if len(feeRuleSetssMap) == 0 && len(feeRuleSetssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		feeRuleSetssMapUpdateField map[model.FeeRuleSetsPrimaryID]FeeRuleSetsUpdateFieldList = map[model.FeeRuleSetsPrimaryID]FeeRuleSetsUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(feeRuleSetssMap) > 0 {
		for id, feeRuleSets := range feeRuleSetssMap {
			if feeRuleSets == nil {
				log.Error().Err(err).Msg("[BulkUpdateFeeRuleSets] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			feeRuleSetssMapUpdateField[id] = defaultFeeRuleSetsUpdateFields(*feeRuleSets)
		}
	} else {
		feeRuleSetssMapUpdateField = feeRuleSetssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFeeRuleSetsQuery(feeRuleSetssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFeeRuleSetsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeRuleSets] failed checking feeRuleSets whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleSets with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFeeRuleSetsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"fee_rule_sets\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeRuleSets] failed exec query")
	}
	return
}

type FeeRuleSetsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFeeRuleSetsFieldParameter(param string, args ...interface{}) FeeRuleSetsFieldParameter {
	return FeeRuleSetsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFeeRuleSetsQuery(mapFeeRuleSetss map[model.FeeRuleSetsPrimaryID]FeeRuleSetsUpdateFieldList, asTableValues string) (primaryIDs []model.FeeRuleSetsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FeeRuleSetsPrimaryID]map[string]interface{}{}
	feeRuleSetsSelectFields := NewFeeRuleSetsSelectFields()
	for id, updateFields := range mapFeeRuleSetss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.feeRuleSetsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFeeRuleSetss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFeeRuleSetsFieldType(updateField.feeRuleSetsField)))
			args = append(args, fields[string(updateField.feeRuleSetsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.feeRuleSetsField))
		if updateField.feeRuleSetsField == feeRuleSetsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.feeRuleSetsField, asTableValues, updateField.feeRuleSetsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.feeRuleSetsField,
				"\"fee_rule_sets\"", updateField.feeRuleSetsField,
				asTableValues, updateField.feeRuleSetsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFeeRuleSetsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FeeRuleSetsPrimaryID, asTableValue string) (whereQry string) {
	feeRuleSetsSelectFields := NewFeeRuleSetsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"fee_rule_sets\".\"id\" = %s.\"id\"::"+GetFeeRuleSetsFieldType(feeRuleSetsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFeeRuleSetsFieldType(feeRuleSetsField FeeRuleSetsField) string {
	selectFeeRuleSetsFields := NewFeeRuleSetsSelectFields()
	switch feeRuleSetsField {

	case selectFeeRuleSetsFields.Id():
		return "uuid"

	case selectFeeRuleSetsFields.FeeProfileId():
		return "uuid"

	case selectFeeRuleSetsFields.RuleSetCode():
		return "text"

	case selectFeeRuleSetsFields.Precedence():
		return "int4"

	case selectFeeRuleSetsFields.EffectiveFrom():
		return "timestamptz"

	case selectFeeRuleSetsFields.EffectiveUntil():
		return "timestamptz"

	case selectFeeRuleSetsFields.RuleSetStatus():
		return "rule_set_status_enum"

	case selectFeeRuleSetsFields.Metadata():
		return "jsonb"

	case selectFeeRuleSetsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFeeRuleSetsFields.MetaCreatedBy():
		return "uuid"

	case selectFeeRuleSetsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFeeRuleSetsFields.MetaUpdatedBy():
		return "uuid"

	case selectFeeRuleSetsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFeeRuleSetsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFeeRuleSets(ctx context.Context, feeRuleSets *model.FeeRuleSets, fieldsInsert ...FeeRuleSetsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFeeRuleSetsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FeeRuleSetsPrimaryID{
		Id: feeRuleSets.Id,
	}
	exists, err := repo.IsExistFeeRuleSetsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeRuleSets] failed checking feeRuleSets whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "feeRuleSets", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFeeRuleSets([]model.FeeRuleSets{*feeRuleSets}, fieldsInsert...)
	commandQuery := fmt.Sprintf(feeRuleSetsQueries.insertFeeRuleSets, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeRuleSets] failed exec create feeRuleSets query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFeeRuleSetsByID(ctx context.Context, primaryID model.FeeRuleSetsPrimaryID) (err error) {
	exists, err := repo.IsExistFeeRuleSetsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeRuleSetsByID] failed checking feeRuleSets whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleSets with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFeeRuleSetsCompositePrimaryKeyWhere([]model.FeeRuleSetsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(feeRuleSetsQueries.deleteFeeRuleSets + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeRuleSetsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRuleSetsByFilter(ctx context.Context, filter model.Filter) (result []model.FeeRuleSetsFilterResult, err error) {
	query, args, err := composeFeeRuleSetsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRuleSetsByFilter] failed compose feeRuleSets filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRuleSetsByFilter] failed get feeRuleSets by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFeeRuleSetsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FeeRuleSetsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFeeRuleSetsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFeeRuleSetsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFeeRuleSetsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFeeRuleSetsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 14+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_profile_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_profile_id\"")
			selectedColumns["fee_profile_id"] = struct{}{}
		}
		if _, selected := selectedColumns["rule_set_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"rule_set_code\"")
			selectedColumns["rule_set_code"] = struct{}{}
		}
		if _, selected := selectedColumns["precedence"]; !selected {
			selectColumns = append(selectColumns, "base.\"precedence\"")
			selectedColumns["precedence"] = struct{}{}
		}
		if _, selected := selectedColumns["effective_from"]; !selected {
			selectColumns = append(selectColumns, "base.\"effective_from\"")
			selectedColumns["effective_from"] = struct{}{}
		}
		if _, selected := selectedColumns["effective_until"]; !selected {
			selectColumns = append(selectColumns, "base.\"effective_until\"")
			selectedColumns["effective_until"] = struct{}{}
		}
		if _, selected := selectedColumns["rule_set_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"rule_set_status\"")
			selectedColumns["rule_set_status"] = struct{}{}
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

type feeRuleSetsFilterPlaceholder struct {
	index int
}

func (p *feeRuleSetsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFeeRuleSetsFilterPredicate(filterField model.FilterField, placeholders *feeRuleSetsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFeeRuleSetsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFeeRuleSetsFilterSQLExpr(spec)
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

func composeFeeRuleSetsFilterGroup(group model.FilterGroup, placeholders *feeRuleSetsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFeeRuleSetsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFeeRuleSetsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFeeRuleSetsFilterWhereQueries(filter model.Filter, placeholders *feeRuleSetsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFeeRuleSetsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFeeRuleSetsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFeeRuleSetsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFeeRuleSetsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFeeRuleSetsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFeeRuleSetsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := feeRuleSetsFilterPlaceholder{index: 1}
	whereQueries, err := composeFeeRuleSetsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFeeRuleSetsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFeeRuleSetsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFeeRuleSetsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"fee_rule_sets\" base%s", strings.Join(selectColumns, ","), composeFeeRuleSetsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFeeRuleSetsByID(ctx context.Context, primaryID model.FeeRuleSetsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFeeRuleSetsCompositePrimaryKeyWhere([]model.FeeRuleSetsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", feeRuleSetsQueries.selectCountFeeRuleSets, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRuleSetsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRuleSets(ctx context.Context, selectFields ...FeeRuleSetsField) (feeRuleSetsList model.FeeRuleSetsList, err error) {
	var (
		defaultFeeRuleSetsSelectFields = defaultFeeRuleSetsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeRuleSetsSelectFields = composeFeeRuleSetsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(feeRuleSetsQueries.selectFeeRuleSets, defaultFeeRuleSetsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &feeRuleSetsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRuleSets] failed get feeRuleSets list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRuleSetsByID(ctx context.Context, primaryID model.FeeRuleSetsPrimaryID, selectFields ...FeeRuleSetsField) (feeRuleSets model.FeeRuleSets, err error) {
	var (
		defaultFeeRuleSetsSelectFields = defaultFeeRuleSetsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeRuleSetsSelectFields = composeFeeRuleSetsSelectFields(selectFields...)
	}
	whereQry, params := composeFeeRuleSetsCompositePrimaryKeyWhere([]model.FeeRuleSetsPrimaryID{primaryID})
	query := fmt.Sprintf(feeRuleSetsQueries.selectFeeRuleSets+" WHERE "+whereQry, defaultFeeRuleSetsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &feeRuleSets, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("feeRuleSets with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFeeRuleSetsByID] failed get feeRuleSets")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFeeRuleSetsByID(ctx context.Context, primaryID model.FeeRuleSetsPrimaryID, feeRuleSets *model.FeeRuleSets, feeRuleSetsUpdateFields ...FeeRuleSetsUpdateField) (err error) {
	exists, err := repo.IsExistFeeRuleSetsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleSets] failed checking feeRuleSets whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleSets with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if feeRuleSets == nil {
		if len(feeRuleSetsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFeeRuleSetsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		feeRuleSets = &model.FeeRuleSets{}
	}
	var (
		defaultFeeRuleSetsUpdateFields = defaultFeeRuleSetsUpdateFields(*feeRuleSets)
		tempUpdateField                FeeRuleSetsUpdateFieldList
		selectFields                   = NewFeeRuleSetsSelectFields()
	)
	if len(feeRuleSetsUpdateFields) > 0 {
		for _, updateField := range feeRuleSetsUpdateFields {
			if updateField.feeRuleSetsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFeeRuleSetsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFeeRuleSetsCompositePrimaryKeyWhere([]model.FeeRuleSetsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFeeRuleSetsCommand(defaultFeeRuleSetsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(feeRuleSetsQueries.updateFeeRuleSets+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleSets] error when try to update feeRuleSets by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFeeRuleSetsByFilter(ctx context.Context, filter model.Filter, feeRuleSetsUpdateFields ...FeeRuleSetsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(feeRuleSetsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FeeRuleSetsUpdateFieldList
		selectFields = NewFeeRuleSetsSelectFields()
	)
	for _, updateField := range feeRuleSetsUpdateFields {
		if updateField.feeRuleSetsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFeeRuleSetsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := feeRuleSetsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFeeRuleSetsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"fee_rule_sets\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleSetsByFilter] error when try to update feeRuleSets by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleSetsByFilter] failed get rows affected")
	}
	return
}

var (
	feeRuleSetsQueries = struct {
		selectFeeRuleSets      string
		selectCountFeeRuleSets string
		deleteFeeRuleSets      string
		updateFeeRuleSets      string
		insertFeeRuleSets      string
	}{
		selectFeeRuleSets:      "SELECT %s FROM \"fee_rule_sets\"",
		selectCountFeeRuleSets: "SELECT COUNT(\"id\") FROM \"fee_rule_sets\"",
		deleteFeeRuleSets:      "DELETE FROM \"fee_rule_sets\"",
		updateFeeRuleSets:      "UPDATE \"fee_rule_sets\" SET %s ",
		insertFeeRuleSets:      "INSERT INTO \"fee_rule_sets\" %s VALUES %s",
	}
)

type FeeRuleSetsRepository interface {
	CreateFeeRuleSets(ctx context.Context, feeRuleSets *model.FeeRuleSets, fieldsInsert ...FeeRuleSetsField) error
	BulkCreateFeeRuleSets(ctx context.Context, feeRuleSetsList []*model.FeeRuleSets, fieldsInsert ...FeeRuleSetsField) error
	ResolveFeeRuleSets(ctx context.Context, selectFields ...FeeRuleSetsField) (model.FeeRuleSetsList, error)
	ResolveFeeRuleSetsByID(ctx context.Context, primaryID model.FeeRuleSetsPrimaryID, selectFields ...FeeRuleSetsField) (model.FeeRuleSets, error)
	UpdateFeeRuleSetsByID(ctx context.Context, id model.FeeRuleSetsPrimaryID, feeRuleSets *model.FeeRuleSets, feeRuleSetsUpdateFields ...FeeRuleSetsUpdateField) error
	UpdateFeeRuleSetsByFilter(ctx context.Context, filter model.Filter, feeRuleSetsUpdateFields ...FeeRuleSetsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFeeRuleSets(ctx context.Context, feeRuleSetsListMap map[model.FeeRuleSetsPrimaryID]*model.FeeRuleSets, FeeRuleSetssMapUpdateFieldsRequest map[model.FeeRuleSetsPrimaryID]FeeRuleSetsUpdateFieldList) (err error)
	DeleteFeeRuleSetsByID(ctx context.Context, id model.FeeRuleSetsPrimaryID) error
	BulkDeleteFeeRuleSetsByIDs(ctx context.Context, ids []model.FeeRuleSetsPrimaryID) error
	ResolveFeeRuleSetsByFilter(ctx context.Context, filter model.Filter) (result []model.FeeRuleSetsFilterResult, err error)
	IsExistFeeRuleSetsByIDs(ctx context.Context, ids []model.FeeRuleSetsPrimaryID) (exists bool, notFoundIds []model.FeeRuleSetsPrimaryID, err error)
	IsExistFeeRuleSetsByID(ctx context.Context, id model.FeeRuleSetsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
