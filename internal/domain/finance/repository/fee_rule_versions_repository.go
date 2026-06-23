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

func composeInsertFieldsAndParamsFeeRuleVersions(feeRuleVersionsList []model.FeeRuleVersions, fieldsInsert ...FeeRuleVersionsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFeeRuleVersionsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, feeRuleVersions := range feeRuleVersionsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, feeRuleVersions.Id)
			case selectField.RuleSetId():
				args = append(args, feeRuleVersions.RuleSetId)
			case selectField.VersionNo():
				args = append(args, feeRuleVersions.VersionNo)
			case selectField.FormulaType():
				args = append(args, feeRuleVersions.FormulaType)
			case selectField.AppliesTo():
				args = append(args, feeRuleVersions.AppliesTo)
			case selectField.PayerType():
				args = append(args, feeRuleVersions.PayerType)
			case selectField.RecipientType():
				args = append(args, feeRuleVersions.RecipientType)
			case selectField.Conditions():
				args = append(args, feeRuleVersions.Conditions)
			case selectField.IsCurrent():
				args = append(args, feeRuleVersions.IsCurrent)
			case selectField.MetaCreatedAt():
				args = append(args, feeRuleVersions.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, feeRuleVersions.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, feeRuleVersions.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, feeRuleVersions.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, feeRuleVersions.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, feeRuleVersions.MetaDeletedBy)

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

func composeFeeRuleVersionsCompositePrimaryKeyWhere(primaryIDs []model.FeeRuleVersionsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"fee_rule_versions\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFeeRuleVersionsSelectFields() string {
	fields := NewFeeRuleVersionsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFeeRuleVersionsSelectFields(selectFields ...FeeRuleVersionsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FeeRuleVersionsField string
type FeeRuleVersionsFieldList []FeeRuleVersionsField

type FeeRuleVersionsSelectFields struct {
}

func (ss FeeRuleVersionsSelectFields) Id() FeeRuleVersionsField {
	return FeeRuleVersionsField("id")
}

func (ss FeeRuleVersionsSelectFields) RuleSetId() FeeRuleVersionsField {
	return FeeRuleVersionsField("rule_set_id")
}

func (ss FeeRuleVersionsSelectFields) VersionNo() FeeRuleVersionsField {
	return FeeRuleVersionsField("version_no")
}

func (ss FeeRuleVersionsSelectFields) FormulaType() FeeRuleVersionsField {
	return FeeRuleVersionsField("formula_type")
}

func (ss FeeRuleVersionsSelectFields) AppliesTo() FeeRuleVersionsField {
	return FeeRuleVersionsField("applies_to")
}

func (ss FeeRuleVersionsSelectFields) PayerType() FeeRuleVersionsField {
	return FeeRuleVersionsField("payer_type")
}

func (ss FeeRuleVersionsSelectFields) RecipientType() FeeRuleVersionsField {
	return FeeRuleVersionsField("recipient_type")
}

func (ss FeeRuleVersionsSelectFields) Conditions() FeeRuleVersionsField {
	return FeeRuleVersionsField("conditions")
}

func (ss FeeRuleVersionsSelectFields) IsCurrent() FeeRuleVersionsField {
	return FeeRuleVersionsField("is_current")
}

func (ss FeeRuleVersionsSelectFields) MetaCreatedAt() FeeRuleVersionsField {
	return FeeRuleVersionsField("meta_created_at")
}

func (ss FeeRuleVersionsSelectFields) MetaCreatedBy() FeeRuleVersionsField {
	return FeeRuleVersionsField("meta_created_by")
}

func (ss FeeRuleVersionsSelectFields) MetaUpdatedAt() FeeRuleVersionsField {
	return FeeRuleVersionsField("meta_updated_at")
}

func (ss FeeRuleVersionsSelectFields) MetaUpdatedBy() FeeRuleVersionsField {
	return FeeRuleVersionsField("meta_updated_by")
}

func (ss FeeRuleVersionsSelectFields) MetaDeletedAt() FeeRuleVersionsField {
	return FeeRuleVersionsField("meta_deleted_at")
}

func (ss FeeRuleVersionsSelectFields) MetaDeletedBy() FeeRuleVersionsField {
	return FeeRuleVersionsField("meta_deleted_by")
}

func (ss FeeRuleVersionsSelectFields) All() FeeRuleVersionsFieldList {
	return []FeeRuleVersionsField{
		ss.Id(),
		ss.RuleSetId(),
		ss.VersionNo(),
		ss.FormulaType(),
		ss.AppliesTo(),
		ss.PayerType(),
		ss.RecipientType(),
		ss.Conditions(),
		ss.IsCurrent(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFeeRuleVersionsSelectFields() FeeRuleVersionsSelectFields {
	return FeeRuleVersionsSelectFields{}
}

type FeeRuleVersionsUpdateFieldOption struct {
	useIncrement bool
}
type FeeRuleVersionsUpdateField struct {
	feeRuleVersionsField FeeRuleVersionsField
	opt                  FeeRuleVersionsUpdateFieldOption
	value                interface{}
}
type FeeRuleVersionsUpdateFieldList []FeeRuleVersionsUpdateField

func defaultFeeRuleVersionsUpdateFieldOption() FeeRuleVersionsUpdateFieldOption {
	return FeeRuleVersionsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFeeRuleVersionsOption(useIncrement bool) func(*FeeRuleVersionsUpdateFieldOption) {
	return func(pcufo *FeeRuleVersionsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFeeRuleVersionsUpdateField(field FeeRuleVersionsField, val interface{}, opts ...func(*FeeRuleVersionsUpdateFieldOption)) FeeRuleVersionsUpdateField {
	defaultOpt := defaultFeeRuleVersionsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FeeRuleVersionsUpdateField{
		feeRuleVersionsField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultFeeRuleVersionsUpdateFields(feeRuleVersions model.FeeRuleVersions) (feeRuleVersionsUpdateFieldList FeeRuleVersionsUpdateFieldList) {
	selectFields := NewFeeRuleVersionsSelectFields()
	feeRuleVersionsUpdateFieldList = append(feeRuleVersionsUpdateFieldList,
		NewFeeRuleVersionsUpdateField(selectFields.Id(), feeRuleVersions.Id),
		NewFeeRuleVersionsUpdateField(selectFields.RuleSetId(), feeRuleVersions.RuleSetId),
		NewFeeRuleVersionsUpdateField(selectFields.VersionNo(), feeRuleVersions.VersionNo),
		NewFeeRuleVersionsUpdateField(selectFields.FormulaType(), feeRuleVersions.FormulaType),
		NewFeeRuleVersionsUpdateField(selectFields.AppliesTo(), feeRuleVersions.AppliesTo),
		NewFeeRuleVersionsUpdateField(selectFields.PayerType(), feeRuleVersions.PayerType),
		NewFeeRuleVersionsUpdateField(selectFields.RecipientType(), feeRuleVersions.RecipientType),
		NewFeeRuleVersionsUpdateField(selectFields.Conditions(), feeRuleVersions.Conditions),
		NewFeeRuleVersionsUpdateField(selectFields.IsCurrent(), feeRuleVersions.IsCurrent),
		NewFeeRuleVersionsUpdateField(selectFields.MetaCreatedAt(), feeRuleVersions.MetaCreatedAt),
		NewFeeRuleVersionsUpdateField(selectFields.MetaCreatedBy(), feeRuleVersions.MetaCreatedBy),
		NewFeeRuleVersionsUpdateField(selectFields.MetaUpdatedAt(), feeRuleVersions.MetaUpdatedAt),
		NewFeeRuleVersionsUpdateField(selectFields.MetaUpdatedBy(), feeRuleVersions.MetaUpdatedBy),
		NewFeeRuleVersionsUpdateField(selectFields.MetaDeletedAt(), feeRuleVersions.MetaDeletedAt),
		NewFeeRuleVersionsUpdateField(selectFields.MetaDeletedBy(), feeRuleVersions.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFeeRuleVersionsCommand(feeRuleVersionsUpdateFieldList FeeRuleVersionsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range feeRuleVersionsUpdateFieldList {
		field := string(updateField.feeRuleVersionsField)
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

func (repo *RepositoryImpl) BulkCreateFeeRuleVersions(ctx context.Context, feeRuleVersionsList []*model.FeeRuleVersions, fieldsInsert ...FeeRuleVersionsField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.FeeRuleVersionsPrimaryID
		feeRuleVersionsValueList []model.FeeRuleVersions
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFeeRuleVersionsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, feeRuleVersions := range feeRuleVersionsList {

		primaryIds = append(primaryIds, feeRuleVersions.ToFeeRuleVersionsPrimaryID())

		feeRuleVersionsValueList = append(feeRuleVersionsValueList, *feeRuleVersions)
	}

	_, notFoundIds, err := repo.IsExistFeeRuleVersionsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeRuleVersions] failed checking feeRuleVersions whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FeeRuleVersionsPrimaryID{}
		mapNotFoundIds := map[model.FeeRuleVersionsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "feeRuleVersions", fmt.Sprintf("feeRuleVersions with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFeeRuleVersions(feeRuleVersionsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(feeRuleVersionsQueries.insertFeeRuleVersions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeRuleVersions] failed exec create feeRuleVersions query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFeeRuleVersionsByIDs(ctx context.Context, primaryIDs []model.FeeRuleVersionsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFeeRuleVersionsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRuleVersionsByIDs] failed checking feeRuleVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleVersions with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_rule_versions\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(feeRuleVersionsQueries.deleteFeeRuleVersions + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRuleVersionsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeRuleVersionsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFeeRuleVersionsByIDs(ctx context.Context, ids []model.FeeRuleVersionsPrimaryID) (exists bool, notFoundIds []model.FeeRuleVersionsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_rule_versions\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(feeRuleVersionsQueries.selectFeeRuleVersions, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRuleVersionsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FeeRuleVersionsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRuleVersionsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FeeRuleVersionsPrimaryID]bool{}
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

// BulkUpdateFeeRuleVersions is used to bulk update feeRuleVersions, by default it will update all field
// if want to update specific field, then fill feeRuleVersionssMapUpdateFieldsRequest else please fill feeRuleVersionssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFeeRuleVersions(ctx context.Context, feeRuleVersionssMap map[model.FeeRuleVersionsPrimaryID]*model.FeeRuleVersions, feeRuleVersionssMapUpdateFieldsRequest map[model.FeeRuleVersionsPrimaryID]FeeRuleVersionsUpdateFieldList) (err error) {
	if len(feeRuleVersionssMap) == 0 && len(feeRuleVersionssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		feeRuleVersionssMapUpdateField map[model.FeeRuleVersionsPrimaryID]FeeRuleVersionsUpdateFieldList = map[model.FeeRuleVersionsPrimaryID]FeeRuleVersionsUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(feeRuleVersionssMap) > 0 {
		for id, feeRuleVersions := range feeRuleVersionssMap {
			if feeRuleVersions == nil {
				log.Error().Err(err).Msg("[BulkUpdateFeeRuleVersions] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			feeRuleVersionssMapUpdateField[id] = defaultFeeRuleVersionsUpdateFields(*feeRuleVersions)
		}
	} else {
		feeRuleVersionssMapUpdateField = feeRuleVersionssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFeeRuleVersionsQuery(feeRuleVersionssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFeeRuleVersionsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeRuleVersions] failed checking feeRuleVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleVersions with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFeeRuleVersionsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"fee_rule_versions\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeRuleVersions] failed exec query")
	}
	return
}

type FeeRuleVersionsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFeeRuleVersionsFieldParameter(param string, args ...interface{}) FeeRuleVersionsFieldParameter {
	return FeeRuleVersionsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFeeRuleVersionsQuery(mapFeeRuleVersionss map[model.FeeRuleVersionsPrimaryID]FeeRuleVersionsUpdateFieldList, asTableValues string) (primaryIDs []model.FeeRuleVersionsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FeeRuleVersionsPrimaryID]map[string]interface{}{}
	feeRuleVersionsSelectFields := NewFeeRuleVersionsSelectFields()
	for id, updateFields := range mapFeeRuleVersionss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.feeRuleVersionsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFeeRuleVersionss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFeeRuleVersionsFieldType(updateField.feeRuleVersionsField)))
			args = append(args, fields[string(updateField.feeRuleVersionsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.feeRuleVersionsField))
		if updateField.feeRuleVersionsField == feeRuleVersionsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.feeRuleVersionsField, asTableValues, updateField.feeRuleVersionsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.feeRuleVersionsField,
				"\"fee_rule_versions\"", updateField.feeRuleVersionsField,
				asTableValues, updateField.feeRuleVersionsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFeeRuleVersionsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FeeRuleVersionsPrimaryID, asTableValue string) (whereQry string) {
	feeRuleVersionsSelectFields := NewFeeRuleVersionsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"fee_rule_versions\".\"id\" = %s.\"id\"::"+GetFeeRuleVersionsFieldType(feeRuleVersionsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFeeRuleVersionsFieldType(feeRuleVersionsField FeeRuleVersionsField) string {
	selectFeeRuleVersionsFields := NewFeeRuleVersionsSelectFields()
	switch feeRuleVersionsField {

	case selectFeeRuleVersionsFields.Id():
		return "uuid"

	case selectFeeRuleVersionsFields.RuleSetId():
		return "uuid"

	case selectFeeRuleVersionsFields.VersionNo():
		return "int4"

	case selectFeeRuleVersionsFields.FormulaType():
		return "formula_type_enum"

	case selectFeeRuleVersionsFields.AppliesTo():
		return "applies_to_enum"

	case selectFeeRuleVersionsFields.PayerType():
		return "payer_type_enum"

	case selectFeeRuleVersionsFields.RecipientType():
		return "recipient_type_enum"

	case selectFeeRuleVersionsFields.Conditions():
		return "jsonb"

	case selectFeeRuleVersionsFields.IsCurrent():
		return "bool"

	case selectFeeRuleVersionsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFeeRuleVersionsFields.MetaCreatedBy():
		return "uuid"

	case selectFeeRuleVersionsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFeeRuleVersionsFields.MetaUpdatedBy():
		return "uuid"

	case selectFeeRuleVersionsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFeeRuleVersionsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFeeRuleVersions(ctx context.Context, feeRuleVersions *model.FeeRuleVersions, fieldsInsert ...FeeRuleVersionsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFeeRuleVersionsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FeeRuleVersionsPrimaryID{
		Id: feeRuleVersions.Id,
	}
	exists, err := repo.IsExistFeeRuleVersionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeRuleVersions] failed checking feeRuleVersions whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "feeRuleVersions", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFeeRuleVersions([]model.FeeRuleVersions{*feeRuleVersions}, fieldsInsert...)
	commandQuery := fmt.Sprintf(feeRuleVersionsQueries.insertFeeRuleVersions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeRuleVersions] failed exec create feeRuleVersions query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFeeRuleVersionsByID(ctx context.Context, primaryID model.FeeRuleVersionsPrimaryID) (err error) {
	exists, err := repo.IsExistFeeRuleVersionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeRuleVersionsByID] failed checking feeRuleVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleVersions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFeeRuleVersionsCompositePrimaryKeyWhere([]model.FeeRuleVersionsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(feeRuleVersionsQueries.deleteFeeRuleVersions + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeRuleVersionsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRuleVersionsByFilter(ctx context.Context, filter model.Filter) (result []model.FeeRuleVersionsFilterResult, err error) {
	query, args, err := composeFeeRuleVersionsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRuleVersionsByFilter] failed compose feeRuleVersions filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRuleVersionsByFilter] failed get feeRuleVersions by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFeeRuleVersionsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FeeRuleVersionsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFeeRuleVersionsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFeeRuleVersionsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFeeRuleVersionsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFeeRuleVersionsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["rule_set_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"rule_set_id\"")
			selectedColumns["rule_set_id"] = struct{}{}
		}
		if _, selected := selectedColumns["version_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"version_no\"")
			selectedColumns["version_no"] = struct{}{}
		}
		if _, selected := selectedColumns["formula_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"formula_type\"")
			selectedColumns["formula_type"] = struct{}{}
		}
		if _, selected := selectedColumns["applies_to"]; !selected {
			selectColumns = append(selectColumns, "base.\"applies_to\"")
			selectedColumns["applies_to"] = struct{}{}
		}
		if _, selected := selectedColumns["payer_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"payer_type\"")
			selectedColumns["payer_type"] = struct{}{}
		}
		if _, selected := selectedColumns["recipient_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"recipient_type\"")
			selectedColumns["recipient_type"] = struct{}{}
		}
		if _, selected := selectedColumns["conditions"]; !selected {
			selectColumns = append(selectColumns, "base.\"conditions\"")
			selectedColumns["conditions"] = struct{}{}
		}
		if _, selected := selectedColumns["is_current"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_current\"")
			selectedColumns["is_current"] = struct{}{}
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

type feeRuleVersionsFilterPlaceholder struct {
	index int
}

func (p *feeRuleVersionsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFeeRuleVersionsFilterPredicate(filterField model.FilterField, placeholders *feeRuleVersionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFeeRuleVersionsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFeeRuleVersionsFilterSQLExpr(spec)
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

func composeFeeRuleVersionsFilterGroup(group model.FilterGroup, placeholders *feeRuleVersionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFeeRuleVersionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFeeRuleVersionsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFeeRuleVersionsFilterWhereQueries(filter model.Filter, placeholders *feeRuleVersionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFeeRuleVersionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFeeRuleVersionsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFeeRuleVersionsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFeeRuleVersionsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFeeRuleVersionsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFeeRuleVersionsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := feeRuleVersionsFilterPlaceholder{index: 1}
	whereQueries, err := composeFeeRuleVersionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFeeRuleVersionsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFeeRuleVersionsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFeeRuleVersionsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"fee_rule_versions\" base%s", strings.Join(selectColumns, ","), composeFeeRuleVersionsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFeeRuleVersionsByID(ctx context.Context, primaryID model.FeeRuleVersionsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFeeRuleVersionsCompositePrimaryKeyWhere([]model.FeeRuleVersionsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", feeRuleVersionsQueries.selectCountFeeRuleVersions, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeRuleVersionsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRuleVersions(ctx context.Context, selectFields ...FeeRuleVersionsField) (feeRuleVersionsList model.FeeRuleVersionsList, err error) {
	var (
		defaultFeeRuleVersionsSelectFields = defaultFeeRuleVersionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeRuleVersionsSelectFields = composeFeeRuleVersionsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(feeRuleVersionsQueries.selectFeeRuleVersions, defaultFeeRuleVersionsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &feeRuleVersionsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeRuleVersions] failed get feeRuleVersions list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeRuleVersionsByID(ctx context.Context, primaryID model.FeeRuleVersionsPrimaryID, selectFields ...FeeRuleVersionsField) (feeRuleVersions model.FeeRuleVersions, err error) {
	var (
		defaultFeeRuleVersionsSelectFields = defaultFeeRuleVersionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeRuleVersionsSelectFields = composeFeeRuleVersionsSelectFields(selectFields...)
	}
	whereQry, params := composeFeeRuleVersionsCompositePrimaryKeyWhere([]model.FeeRuleVersionsPrimaryID{primaryID})
	query := fmt.Sprintf(feeRuleVersionsQueries.selectFeeRuleVersions+" WHERE "+whereQry, defaultFeeRuleVersionsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &feeRuleVersions, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("feeRuleVersions with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFeeRuleVersionsByID] failed get feeRuleVersions")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFeeRuleVersionsByID(ctx context.Context, primaryID model.FeeRuleVersionsPrimaryID, feeRuleVersions *model.FeeRuleVersions, feeRuleVersionsUpdateFields ...FeeRuleVersionsUpdateField) (err error) {
	exists, err := repo.IsExistFeeRuleVersionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleVersions] failed checking feeRuleVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeRuleVersions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if feeRuleVersions == nil {
		if len(feeRuleVersionsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFeeRuleVersionsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		feeRuleVersions = &model.FeeRuleVersions{}
	}
	var (
		defaultFeeRuleVersionsUpdateFields = defaultFeeRuleVersionsUpdateFields(*feeRuleVersions)
		tempUpdateField                    FeeRuleVersionsUpdateFieldList
		selectFields                       = NewFeeRuleVersionsSelectFields()
	)
	if len(feeRuleVersionsUpdateFields) > 0 {
		for _, updateField := range feeRuleVersionsUpdateFields {
			if updateField.feeRuleVersionsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFeeRuleVersionsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFeeRuleVersionsCompositePrimaryKeyWhere([]model.FeeRuleVersionsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFeeRuleVersionsCommand(defaultFeeRuleVersionsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(feeRuleVersionsQueries.updateFeeRuleVersions+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleVersions] error when try to update feeRuleVersions by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFeeRuleVersionsByFilter(ctx context.Context, filter model.Filter, feeRuleVersionsUpdateFields ...FeeRuleVersionsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(feeRuleVersionsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FeeRuleVersionsUpdateFieldList
		selectFields = NewFeeRuleVersionsSelectFields()
	)
	for _, updateField := range feeRuleVersionsUpdateFields {
		if updateField.feeRuleVersionsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFeeRuleVersionsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := feeRuleVersionsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFeeRuleVersionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"fee_rule_versions\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleVersionsByFilter] error when try to update feeRuleVersions by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeRuleVersionsByFilter] failed get rows affected")
	}
	return
}

var (
	feeRuleVersionsQueries = struct {
		selectFeeRuleVersions      string
		selectCountFeeRuleVersions string
		deleteFeeRuleVersions      string
		updateFeeRuleVersions      string
		insertFeeRuleVersions      string
	}{
		selectFeeRuleVersions:      "SELECT %s FROM \"fee_rule_versions\"",
		selectCountFeeRuleVersions: "SELECT COUNT(\"id\") FROM \"fee_rule_versions\"",
		deleteFeeRuleVersions:      "DELETE FROM \"fee_rule_versions\"",
		updateFeeRuleVersions:      "UPDATE \"fee_rule_versions\" SET %s ",
		insertFeeRuleVersions:      "INSERT INTO \"fee_rule_versions\" %s VALUES %s",
	}
)

type FeeRuleVersionsRepository interface {
	CreateFeeRuleVersions(ctx context.Context, feeRuleVersions *model.FeeRuleVersions, fieldsInsert ...FeeRuleVersionsField) error
	BulkCreateFeeRuleVersions(ctx context.Context, feeRuleVersionsList []*model.FeeRuleVersions, fieldsInsert ...FeeRuleVersionsField) error
	ResolveFeeRuleVersions(ctx context.Context, selectFields ...FeeRuleVersionsField) (model.FeeRuleVersionsList, error)
	ResolveFeeRuleVersionsByID(ctx context.Context, primaryID model.FeeRuleVersionsPrimaryID, selectFields ...FeeRuleVersionsField) (model.FeeRuleVersions, error)
	UpdateFeeRuleVersionsByID(ctx context.Context, id model.FeeRuleVersionsPrimaryID, feeRuleVersions *model.FeeRuleVersions, feeRuleVersionsUpdateFields ...FeeRuleVersionsUpdateField) error
	UpdateFeeRuleVersionsByFilter(ctx context.Context, filter model.Filter, feeRuleVersionsUpdateFields ...FeeRuleVersionsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFeeRuleVersions(ctx context.Context, feeRuleVersionsListMap map[model.FeeRuleVersionsPrimaryID]*model.FeeRuleVersions, FeeRuleVersionssMapUpdateFieldsRequest map[model.FeeRuleVersionsPrimaryID]FeeRuleVersionsUpdateFieldList) (err error)
	DeleteFeeRuleVersionsByID(ctx context.Context, id model.FeeRuleVersionsPrimaryID) error
	BulkDeleteFeeRuleVersionsByIDs(ctx context.Context, ids []model.FeeRuleVersionsPrimaryID) error
	ResolveFeeRuleVersionsByFilter(ctx context.Context, filter model.Filter) (result []model.FeeRuleVersionsFilterResult, err error)
	IsExistFeeRuleVersionsByIDs(ctx context.Context, ids []model.FeeRuleVersionsPrimaryID) (exists bool, notFoundIds []model.FeeRuleVersionsPrimaryID, err error)
	IsExistFeeRuleVersionsByID(ctx context.Context, id model.FeeRuleVersionsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
