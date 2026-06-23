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

func composeInsertFieldsAndParamsFeeRules(feeRulesList []model.FeeRules, fieldsInsert ...FeeRulesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFeeRulesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, feeRules := range feeRulesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, feeRules.Id)
			case selectField.FeeRuleVersionId():
				args = append(args, feeRules.FeeRuleVersionId)
			case selectField.RuleName():
				args = append(args, feeRules.RuleName)
			case selectField.MinAmount():
				args = append(args, feeRules.MinAmount)
			case selectField.MaxAmount():
				args = append(args, feeRules.MaxAmount)
			case selectField.PercentageRate():
				args = append(args, feeRules.PercentageRate)
			case selectField.FlatAmount():
				args = append(args, feeRules.FlatAmount)
			case selectField.CapAmount():
				args = append(args, feeRules.CapAmount)
			case selectField.FloorAmount():
				args = append(args, feeRules.FloorAmount)
			case selectField.TaxInclusive():
				args = append(args, feeRules.TaxInclusive)
			case selectField.SortOrder():
				args = append(args, feeRules.SortOrder)
			case selectField.Metadata():
				args = append(args, feeRules.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, feeRules.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, feeRules.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, feeRules.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, feeRules.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, feeRules.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, feeRules.MetaDeletedBy)

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

func composeFeeRulesCompositePrimaryKeyWhere(primaryIDs []model.FeeRulesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"fee_rules\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFeeRulesSelectFields() string {
	fields := NewFeeRulesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFeeRulesSelectFields(selectFields ...FeeRulesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FeeRulesField string
type FeeRulesFieldList []FeeRulesField

type FeeRulesSelectFields struct {
}

func (ss FeeRulesSelectFields) Id() FeeRulesField {
	return FeeRulesField("id")
}

func (ss FeeRulesSelectFields) FeeRuleVersionId() FeeRulesField {
	return FeeRulesField("fee_rule_version_id")
}

func (ss FeeRulesSelectFields) RuleName() FeeRulesField {
	return FeeRulesField("rule_name")
}

func (ss FeeRulesSelectFields) MinAmount() FeeRulesField {
	return FeeRulesField("min_amount")
}

func (ss FeeRulesSelectFields) MaxAmount() FeeRulesField {
	return FeeRulesField("max_amount")
}

func (ss FeeRulesSelectFields) PercentageRate() FeeRulesField {
	return FeeRulesField("percentage_rate")
}

func (ss FeeRulesSelectFields) FlatAmount() FeeRulesField {
	return FeeRulesField("flat_amount")
}

func (ss FeeRulesSelectFields) CapAmount() FeeRulesField {
	return FeeRulesField("cap_amount")
}

func (ss FeeRulesSelectFields) FloorAmount() FeeRulesField {
	return FeeRulesField("floor_amount")
}

func (ss FeeRulesSelectFields) TaxInclusive() FeeRulesField {
	return FeeRulesField("tax_inclusive")
}

func (ss FeeRulesSelectFields) SortOrder() FeeRulesField {
	return FeeRulesField("sort_order")
}

func (ss FeeRulesSelectFields) Metadata() FeeRulesField {
	return FeeRulesField("metadata")
}

func (ss FeeRulesSelectFields) MetaCreatedAt() FeeRulesField {
	return FeeRulesField("meta_created_at")
}

func (ss FeeRulesSelectFields) MetaCreatedBy() FeeRulesField {
	return FeeRulesField("meta_created_by")
}

func (ss FeeRulesSelectFields) MetaUpdatedAt() FeeRulesField {
	return FeeRulesField("meta_updated_at")
}

func (ss FeeRulesSelectFields) MetaUpdatedBy() FeeRulesField {
	return FeeRulesField("meta_updated_by")
}

func (ss FeeRulesSelectFields) MetaDeletedAt() FeeRulesField {
	return FeeRulesField("meta_deleted_at")
}

func (ss FeeRulesSelectFields) MetaDeletedBy() FeeRulesField {
	return FeeRulesField("meta_deleted_by")
}

func (ss FeeRulesSelectFields) All() FeeRulesFieldList {
	return []FeeRulesField{
		ss.Id(),
		ss.FeeRuleVersionId(),
		ss.RuleName(),
		ss.MinAmount(),
		ss.MaxAmount(),
		ss.PercentageRate(),
		ss.FlatAmount(),
		ss.CapAmount(),
		ss.FloorAmount(),
		ss.TaxInclusive(),
		ss.SortOrder(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFeeRulesSelectFields() FeeRulesSelectFields {
	return FeeRulesSelectFields{}
}

type FeeRulesUpdateFieldOption struct {
	useIncrement bool
}
type FeeRulesUpdateField struct {
	feeRulesField FeeRulesField
	opt           FeeRulesUpdateFieldOption
	value         interface{}
}
type FeeRulesUpdateFieldList []FeeRulesUpdateField

func defaultFeeRulesUpdateFieldOption() FeeRulesUpdateFieldOption {
	return FeeRulesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFeeRulesOption(useIncrement bool) func(*FeeRulesUpdateFieldOption) {
	return func(pcufo *FeeRulesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFeeRulesUpdateField(field FeeRulesField, val interface{}, opts ...func(*FeeRulesUpdateFieldOption)) FeeRulesUpdateField {
	defaultOpt := defaultFeeRulesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FeeRulesUpdateField{
		feeRulesField: field,
		value:         val,
		opt:           defaultOpt,
	}
}
func defaultFeeRulesUpdateFields(feeRules model.FeeRules) (feeRulesUpdateFieldList FeeRulesUpdateFieldList) {
	selectFields := NewFeeRulesSelectFields()
	feeRulesUpdateFieldList = append(feeRulesUpdateFieldList,
		NewFeeRulesUpdateField(selectFields.Id(), feeRules.Id),
		NewFeeRulesUpdateField(selectFields.FeeRuleVersionId(), feeRules.FeeRuleVersionId),
		NewFeeRulesUpdateField(selectFields.RuleName(), feeRules.RuleName),
		NewFeeRulesUpdateField(selectFields.MinAmount(), feeRules.MinAmount),
		NewFeeRulesUpdateField(selectFields.MaxAmount(), feeRules.MaxAmount),
		NewFeeRulesUpdateField(selectFields.PercentageRate(), feeRules.PercentageRate),
		NewFeeRulesUpdateField(selectFields.FlatAmount(), feeRules.FlatAmount),
		NewFeeRulesUpdateField(selectFields.CapAmount(), feeRules.CapAmount),
		NewFeeRulesUpdateField(selectFields.FloorAmount(), feeRules.FloorAmount),
		NewFeeRulesUpdateField(selectFields.TaxInclusive(), feeRules.TaxInclusive),
		NewFeeRulesUpdateField(selectFields.SortOrder(), feeRules.SortOrder),
		NewFeeRulesUpdateField(selectFields.Metadata(), feeRules.Metadata),
		NewFeeRulesUpdateField(selectFields.MetaCreatedAt(), feeRules.MetaCreatedAt),
		NewFeeRulesUpdateField(selectFields.MetaCreatedBy(), feeRules.MetaCreatedBy),
		NewFeeRulesUpdateField(selectFields.MetaUpdatedAt(), feeRules.MetaUpdatedAt),
		NewFeeRulesUpdateField(selectFields.MetaUpdatedBy(), feeRules.MetaUpdatedBy),
		NewFeeRulesUpdateField(selectFields.MetaDeletedAt(), feeRules.MetaDeletedAt),
		NewFeeRulesUpdateField(selectFields.MetaDeletedBy(), feeRules.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFeeRulesCommand(feeRulesUpdateFieldList FeeRulesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range feeRulesUpdateFieldList {
		field := string(updateField.feeRulesField)
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

func (repo *RepositoryImpl) BulkCreateFeeRules(ctx context.Context, feeRulesList []*model.FeeRules, fieldsInsert ...FeeRulesField) (err error) {
	var (
		fieldsStr         string
		valueListStr      []string
		argsList          []interface{}
		primaryIds        []model.FeeRulesPrimaryID
		feeRulesValueList []model.FeeRules
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFeeRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, feeRules := range feeRulesList {

		primaryIds = append(primaryIds, feeRules.ToFeeRulesPrimaryID())

		feeRulesValueList = append(feeRulesValueList, *feeRules)
	}

	_, notFoundIds, err := repo.IsExistFeeRulesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeRules] failed checking feeRules whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FeeRulesPrimaryID{}
		mapNotFoundIds := map[model.FeeRulesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "feeRules", fmt.Sprintf("feeRules with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFeeRules(feeRulesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(feeRulesQueries.insertFeeRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeRules] failed exec create feeRules query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFeeRulesByIDs(ctx context.Context, primaryIDs []model.FeeRulesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFeeRulesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRulesByIDs] failed checking feeRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRules with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_rules\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(feeRulesQueries.deleteFeeRules + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRulesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRulesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFeeRulesByIDs(ctx context.Context, ids []model.FeeRulesPrimaryID) (exists bool, notFoundIds []model.FeeRulesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_rules\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(feeRulesQueries.selectFeeRules, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRulesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FeeRulesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRulesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FeeRulesPrimaryID]bool{}
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

// BulkUpdateFeeRules is used to bulk update feeRules, by default it will update all field
// if want to update specific field, then fill feeRulessMapUpdateFieldsRequest else please fill feeRulessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFeeRules(ctx context.Context, feeRulessMap map[model.FeeRulesPrimaryID]*model.FeeRules, feeRulessMapUpdateFieldsRequest map[model.FeeRulesPrimaryID]FeeRulesUpdateFieldList) (err error) {
	if len(feeRulessMap) == 0 && len(feeRulessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		feeRulessMapUpdateField map[model.FeeRulesPrimaryID]FeeRulesUpdateFieldList = map[model.FeeRulesPrimaryID]FeeRulesUpdateFieldList{}
		asTableValues           string                                              = "myvalues"
	)

	if len(feeRulessMap) > 0 {
		for id, feeRules := range feeRulessMap {
			if feeRules == nil {
				log.Error().Err(err).Msg("[BulkUpdateFeeRules] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			feeRulessMapUpdateField[id] = defaultFeeRulesUpdateFields(*feeRules)
		}
	} else {
		feeRulessMapUpdateField = feeRulessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFeeRulesQuery(feeRulessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFeeRulesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeRules] failed checking feeRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRules with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFeeRulesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"fee_rules\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeRules] failed exec query")
	}
	return
}

type FeeRulesFieldParameter struct {
	param string
	args  []interface{}
}

func NewFeeRulesFieldParameter(param string, args ...interface{}) FeeRulesFieldParameter {
	return FeeRulesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFeeRulesQuery(mapFeeRuless map[model.FeeRulesPrimaryID]FeeRulesUpdateFieldList, asTableValues string) (primaryIDs []model.FeeRulesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FeeRulesPrimaryID]map[string]interface{}{}
	feeRulesSelectFields := NewFeeRulesSelectFields()
	for id, updateFields := range mapFeeRuless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.feeRulesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFeeRuless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFeeRulesFieldType(updateField.feeRulesField)))
			args = append(args, fields[string(updateField.feeRulesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.feeRulesField))
		if updateField.feeRulesField == feeRulesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.feeRulesField, asTableValues, updateField.feeRulesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.feeRulesField,
				"\"fee_rules\"", updateField.feeRulesField,
				asTableValues, updateField.feeRulesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFeeRulesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FeeRulesPrimaryID, asTableValue string) (whereQry string) {
	feeRulesSelectFields := NewFeeRulesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"fee_rules\".\"id\" = %s.\"id\"::"+GetFeeRulesFieldType(feeRulesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFeeRulesFieldType(feeRulesField FeeRulesField) string {
	selectFeeRulesFields := NewFeeRulesSelectFields()
	switch feeRulesField {

	case selectFeeRulesFields.Id():
		return "uuid"

	case selectFeeRulesFields.FeeRuleVersionId():
		return "uuid"

	case selectFeeRulesFields.RuleName():
		return "text"

	case selectFeeRulesFields.MinAmount():
		return "numeric"

	case selectFeeRulesFields.MaxAmount():
		return "numeric"

	case selectFeeRulesFields.PercentageRate():
		return "numeric"

	case selectFeeRulesFields.FlatAmount():
		return "numeric"

	case selectFeeRulesFields.CapAmount():
		return "numeric"

	case selectFeeRulesFields.FloorAmount():
		return "numeric"

	case selectFeeRulesFields.TaxInclusive():
		return "bool"

	case selectFeeRulesFields.SortOrder():
		return "int4"

	case selectFeeRulesFields.Metadata():
		return "jsonb"

	case selectFeeRulesFields.MetaCreatedAt():
		return "timestamptz"

	case selectFeeRulesFields.MetaCreatedBy():
		return "uuid"

	case selectFeeRulesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFeeRulesFields.MetaUpdatedBy():
		return "uuid"

	case selectFeeRulesFields.MetaDeletedAt():
		return "timestamptz"

	case selectFeeRulesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFeeRules(ctx context.Context, feeRules *model.FeeRules, fieldsInsert ...FeeRulesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFeeRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FeeRulesPrimaryID{
		Id: feeRules.Id,
	}
	exists, err := repo.IsExistFeeRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeRules] failed checking feeRules whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "feeRules", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFeeRules([]model.FeeRules{*feeRules}, fieldsInsert...)
	commandQuery := fmt.Sprintf(feeRulesQueries.insertFeeRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeRules] failed exec create feeRules query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFeeRulesByID(ctx context.Context, primaryID model.FeeRulesPrimaryID) (err error) {
	exists, err := repo.IsExistFeeRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeRulesByID] failed checking feeRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFeeRulesCompositePrimaryKeyWhere([]model.FeeRulesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(feeRulesQueries.deleteFeeRules + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeRulesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRulesByFilter(ctx context.Context, filter model.Filter) (result []model.FeeRulesFilterResult, err error) {
	query, args, err := composeFeeRulesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRulesByFilter] failed compose feeRules filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRulesByFilter] failed get feeRules by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFeeRulesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FeeRulesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFeeRulesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFeeRulesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFeeRulesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFeeRulesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["fee_rule_version_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_rule_version_id\"")
			selectedColumns["fee_rule_version_id"] = struct{}{}
		}
		if _, selected := selectedColumns["rule_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"rule_name\"")
			selectedColumns["rule_name"] = struct{}{}
		}
		if _, selected := selectedColumns["min_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"min_amount\"")
			selectedColumns["min_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["max_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"max_amount\"")
			selectedColumns["max_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["percentage_rate"]; !selected {
			selectColumns = append(selectColumns, "base.\"percentage_rate\"")
			selectedColumns["percentage_rate"] = struct{}{}
		}
		if _, selected := selectedColumns["flat_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"flat_amount\"")
			selectedColumns["flat_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["cap_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"cap_amount\"")
			selectedColumns["cap_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["floor_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"floor_amount\"")
			selectedColumns["floor_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_inclusive"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_inclusive\"")
			selectedColumns["tax_inclusive"] = struct{}{}
		}
		if _, selected := selectedColumns["sort_order"]; !selected {
			selectColumns = append(selectColumns, "base.\"sort_order\"")
			selectedColumns["sort_order"] = struct{}{}
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

type feeRulesFilterPlaceholder struct {
	index int
}

func (p *feeRulesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFeeRulesFilterPredicate(filterField model.FilterField, placeholders *feeRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFeeRulesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFeeRulesFilterSQLExpr(spec)
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

func composeFeeRulesFilterGroup(group model.FilterGroup, placeholders *feeRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFeeRulesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFeeRulesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFeeRulesFilterWhereQueries(filter model.Filter, placeholders *feeRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFeeRulesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFeeRulesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFeeRulesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFeeRulesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFeeRulesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFeeRulesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := feeRulesFilterPlaceholder{index: 1}
	whereQueries, err := composeFeeRulesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFeeRulesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFeeRulesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFeeRulesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"fee_rules\" base%s", strings.Join(selectColumns, ","), composeFeeRulesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFeeRulesByID(ctx context.Context, primaryID model.FeeRulesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFeeRulesCompositePrimaryKeyWhere([]model.FeeRulesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", feeRulesQueries.selectCountFeeRules, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRulesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRules(ctx context.Context, selectFields ...FeeRulesField) (feeRulesList model.FeeRulesList, err error) {
	var (
		defaultFeeRulesSelectFields = defaultFeeRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeRulesSelectFields = composeFeeRulesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(feeRulesQueries.selectFeeRules, defaultFeeRulesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &feeRulesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRules] failed get feeRules list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRulesByID(ctx context.Context, primaryID model.FeeRulesPrimaryID, selectFields ...FeeRulesField) (feeRules model.FeeRules, err error) {
	var (
		defaultFeeRulesSelectFields = defaultFeeRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeRulesSelectFields = composeFeeRulesSelectFields(selectFields...)
	}
	whereQry, params := composeFeeRulesCompositePrimaryKeyWhere([]model.FeeRulesPrimaryID{primaryID})
	query := fmt.Sprintf(feeRulesQueries.selectFeeRules+" WHERE "+whereQry, defaultFeeRulesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &feeRules, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("feeRules with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFeeRulesByID] failed get feeRules")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFeeRulesByID(ctx context.Context, primaryID model.FeeRulesPrimaryID, feeRules *model.FeeRules, feeRulesUpdateFields ...FeeRulesUpdateField) (err error) {
	exists, err := repo.IsExistFeeRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRules] failed checking feeRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if feeRules == nil {
		if len(feeRulesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFeeRulesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		feeRules = &model.FeeRules{}
	}
	var (
		defaultFeeRulesUpdateFields = defaultFeeRulesUpdateFields(*feeRules)
		tempUpdateField             FeeRulesUpdateFieldList
		selectFields                = NewFeeRulesSelectFields()
	)
	if len(feeRulesUpdateFields) > 0 {
		for _, updateField := range feeRulesUpdateFields {
			if updateField.feeRulesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFeeRulesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFeeRulesCompositePrimaryKeyWhere([]model.FeeRulesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFeeRulesCommand(defaultFeeRulesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(feeRulesQueries.updateFeeRules+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRules] error when try to update feeRules by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFeeRulesByFilter(ctx context.Context, filter model.Filter, feeRulesUpdateFields ...FeeRulesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(feeRulesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FeeRulesUpdateFieldList
		selectFields = NewFeeRulesSelectFields()
	)
	for _, updateField := range feeRulesUpdateFields {
		if updateField.feeRulesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFeeRulesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := feeRulesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFeeRulesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"fee_rules\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRulesByFilter] error when try to update feeRules by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRulesByFilter] failed get rows affected")
	}
	return
}

var (
	feeRulesQueries = struct {
		selectFeeRules      string
		selectCountFeeRules string
		deleteFeeRules      string
		updateFeeRules      string
		insertFeeRules      string
	}{
		selectFeeRules:      "SELECT %s FROM \"fee_rules\"",
		selectCountFeeRules: "SELECT COUNT(\"id\") FROM \"fee_rules\"",
		deleteFeeRules:      "DELETE FROM \"fee_rules\"",
		updateFeeRules:      "UPDATE \"fee_rules\" SET %s ",
		insertFeeRules:      "INSERT INTO \"fee_rules\" %s VALUES %s",
	}
)

type FeeRulesRepository interface {
	CreateFeeRules(ctx context.Context, feeRules *model.FeeRules, fieldsInsert ...FeeRulesField) error
	BulkCreateFeeRules(ctx context.Context, feeRulesList []*model.FeeRules, fieldsInsert ...FeeRulesField) error
	ResolveFeeRules(ctx context.Context, selectFields ...FeeRulesField) (model.FeeRulesList, error)
	ResolveFeeRulesByID(ctx context.Context, primaryID model.FeeRulesPrimaryID, selectFields ...FeeRulesField) (model.FeeRules, error)
	UpdateFeeRulesByID(ctx context.Context, id model.FeeRulesPrimaryID, feeRules *model.FeeRules, feeRulesUpdateFields ...FeeRulesUpdateField) error
	UpdateFeeRulesByFilter(ctx context.Context, filter model.Filter, feeRulesUpdateFields ...FeeRulesUpdateField) (rowsAffected int64, err error)
	BulkUpdateFeeRules(ctx context.Context, feeRulesListMap map[model.FeeRulesPrimaryID]*model.FeeRules, FeeRulessMapUpdateFieldsRequest map[model.FeeRulesPrimaryID]FeeRulesUpdateFieldList) (err error)
	DeleteFeeRulesByID(ctx context.Context, id model.FeeRulesPrimaryID) error
	BulkDeleteFeeRulesByIDs(ctx context.Context, ids []model.FeeRulesPrimaryID) error
	ResolveFeeRulesByFilter(ctx context.Context, filter model.Filter) (result []model.FeeRulesFilterResult, err error)
	IsExistFeeRulesByIDs(ctx context.Context, ids []model.FeeRulesPrimaryID) (exists bool, notFoundIds []model.FeeRulesPrimaryID, err error)
	IsExistFeeRulesByID(ctx context.Context, id model.FeeRulesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
