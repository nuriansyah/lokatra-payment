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

func composeInsertFieldsAndParamsTaxRules(taxRulesList []model.TaxRules, fieldsInsert ...TaxRulesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewTaxRulesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, taxRules := range taxRulesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, taxRules.Id)
			case selectField.RuleCode():
				args = append(args, taxRules.RuleCode)
			case selectField.CountryCode():
				args = append(args, taxRules.CountryCode)
			case selectField.TaxType():
				args = append(args, taxRules.TaxType)
			case selectField.Rate():
				args = append(args, taxRules.Rate)
			case selectField.TaxInclusive():
				args = append(args, taxRules.TaxInclusive)
			case selectField.LiabilityPartyId():
				args = append(args, taxRules.LiabilityPartyId)
			case selectField.BeneficiaryPartyId():
				args = append(args, taxRules.BeneficiaryPartyId)
			case selectField.Priority():
				args = append(args, taxRules.Priority)
			case selectField.IsActive():
				args = append(args, taxRules.IsActive)
			case selectField.ValidFrom():
				args = append(args, taxRules.ValidFrom)
			case selectField.ValidUntil():
				args = append(args, taxRules.ValidUntil)
			case selectField.Metadata():
				args = append(args, taxRules.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, taxRules.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, taxRules.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, taxRules.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, taxRules.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, taxRules.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, taxRules.MetaDeletedBy)

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

func composeTaxRulesCompositePrimaryKeyWhere(primaryIDs []model.TaxRulesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"tax_rules\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultTaxRulesSelectFields() string {
	fields := NewTaxRulesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeTaxRulesSelectFields(selectFields ...TaxRulesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type TaxRulesField string
type TaxRulesFieldList []TaxRulesField

type TaxRulesSelectFields struct {
}

func (ss TaxRulesSelectFields) Id() TaxRulesField {
	return TaxRulesField("id")
}

func (ss TaxRulesSelectFields) RuleCode() TaxRulesField {
	return TaxRulesField("rule_code")
}

func (ss TaxRulesSelectFields) CountryCode() TaxRulesField {
	return TaxRulesField("country_code")
}

func (ss TaxRulesSelectFields) TaxType() TaxRulesField {
	return TaxRulesField("tax_type")
}

func (ss TaxRulesSelectFields) Rate() TaxRulesField {
	return TaxRulesField("rate")
}

func (ss TaxRulesSelectFields) TaxInclusive() TaxRulesField {
	return TaxRulesField("tax_inclusive")
}

func (ss TaxRulesSelectFields) LiabilityPartyId() TaxRulesField {
	return TaxRulesField("liability_party_id")
}

func (ss TaxRulesSelectFields) BeneficiaryPartyId() TaxRulesField {
	return TaxRulesField("beneficiary_party_id")
}

func (ss TaxRulesSelectFields) Priority() TaxRulesField {
	return TaxRulesField("priority")
}

func (ss TaxRulesSelectFields) IsActive() TaxRulesField {
	return TaxRulesField("is_active")
}

func (ss TaxRulesSelectFields) ValidFrom() TaxRulesField {
	return TaxRulesField("valid_from")
}

func (ss TaxRulesSelectFields) ValidUntil() TaxRulesField {
	return TaxRulesField("valid_until")
}

func (ss TaxRulesSelectFields) Metadata() TaxRulesField {
	return TaxRulesField("metadata")
}

func (ss TaxRulesSelectFields) MetaCreatedAt() TaxRulesField {
	return TaxRulesField("meta_created_at")
}

func (ss TaxRulesSelectFields) MetaCreatedBy() TaxRulesField {
	return TaxRulesField("meta_created_by")
}

func (ss TaxRulesSelectFields) MetaUpdatedAt() TaxRulesField {
	return TaxRulesField("meta_updated_at")
}

func (ss TaxRulesSelectFields) MetaUpdatedBy() TaxRulesField {
	return TaxRulesField("meta_updated_by")
}

func (ss TaxRulesSelectFields) MetaDeletedAt() TaxRulesField {
	return TaxRulesField("meta_deleted_at")
}

func (ss TaxRulesSelectFields) MetaDeletedBy() TaxRulesField {
	return TaxRulesField("meta_deleted_by")
}

func (ss TaxRulesSelectFields) All() TaxRulesFieldList {
	return []TaxRulesField{
		ss.Id(),
		ss.RuleCode(),
		ss.CountryCode(),
		ss.TaxType(),
		ss.Rate(),
		ss.TaxInclusive(),
		ss.LiabilityPartyId(),
		ss.BeneficiaryPartyId(),
		ss.Priority(),
		ss.IsActive(),
		ss.ValidFrom(),
		ss.ValidUntil(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewTaxRulesSelectFields() TaxRulesSelectFields {
	return TaxRulesSelectFields{}
}

type TaxRulesUpdateFieldOption struct {
	useIncrement bool
}
type TaxRulesUpdateField struct {
	taxRulesField TaxRulesField
	opt           TaxRulesUpdateFieldOption
	value         interface{}
}
type TaxRulesUpdateFieldList []TaxRulesUpdateField

func defaultTaxRulesUpdateFieldOption() TaxRulesUpdateFieldOption {
	return TaxRulesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementTaxRulesOption(useIncrement bool) func(*TaxRulesUpdateFieldOption) {
	return func(pcufo *TaxRulesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewTaxRulesUpdateField(field TaxRulesField, val interface{}, opts ...func(*TaxRulesUpdateFieldOption)) TaxRulesUpdateField {
	defaultOpt := defaultTaxRulesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return TaxRulesUpdateField{
		taxRulesField: field,
		value:         val,
		opt:           defaultOpt,
	}
}
func defaultTaxRulesUpdateFields(taxRules model.TaxRules) (taxRulesUpdateFieldList TaxRulesUpdateFieldList) {
	selectFields := NewTaxRulesSelectFields()
	taxRulesUpdateFieldList = append(taxRulesUpdateFieldList,
		NewTaxRulesUpdateField(selectFields.Id(), taxRules.Id),
		NewTaxRulesUpdateField(selectFields.RuleCode(), taxRules.RuleCode),
		NewTaxRulesUpdateField(selectFields.CountryCode(), taxRules.CountryCode),
		NewTaxRulesUpdateField(selectFields.TaxType(), taxRules.TaxType),
		NewTaxRulesUpdateField(selectFields.Rate(), taxRules.Rate),
		NewTaxRulesUpdateField(selectFields.TaxInclusive(), taxRules.TaxInclusive),
		NewTaxRulesUpdateField(selectFields.LiabilityPartyId(), taxRules.LiabilityPartyId),
		NewTaxRulesUpdateField(selectFields.BeneficiaryPartyId(), taxRules.BeneficiaryPartyId),
		NewTaxRulesUpdateField(selectFields.Priority(), taxRules.Priority),
		NewTaxRulesUpdateField(selectFields.IsActive(), taxRules.IsActive),
		NewTaxRulesUpdateField(selectFields.ValidFrom(), taxRules.ValidFrom),
		NewTaxRulesUpdateField(selectFields.ValidUntil(), taxRules.ValidUntil),
		NewTaxRulesUpdateField(selectFields.Metadata(), taxRules.Metadata),
		NewTaxRulesUpdateField(selectFields.MetaCreatedAt(), taxRules.MetaCreatedAt),
		NewTaxRulesUpdateField(selectFields.MetaCreatedBy(), taxRules.MetaCreatedBy),
		NewTaxRulesUpdateField(selectFields.MetaUpdatedAt(), taxRules.MetaUpdatedAt),
		NewTaxRulesUpdateField(selectFields.MetaUpdatedBy(), taxRules.MetaUpdatedBy),
		NewTaxRulesUpdateField(selectFields.MetaDeletedAt(), taxRules.MetaDeletedAt),
		NewTaxRulesUpdateField(selectFields.MetaDeletedBy(), taxRules.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsTaxRulesCommand(taxRulesUpdateFieldList TaxRulesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range taxRulesUpdateFieldList {
		field := string(updateField.taxRulesField)
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

func (repo *RepositoryImpl) BulkCreateTaxRules(ctx context.Context, taxRulesList []*model.TaxRules, fieldsInsert ...TaxRulesField) (err error) {
	var (
		fieldsStr         string
		valueListStr      []string
		argsList          []interface{}
		primaryIds        []model.TaxRulesPrimaryID
		taxRulesValueList []model.TaxRules
	)

	if len(fieldsInsert) == 0 {
		selectField := NewTaxRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, taxRules := range taxRulesList {

		primaryIds = append(primaryIds, taxRules.ToTaxRulesPrimaryID())

		taxRulesValueList = append(taxRulesValueList, *taxRules)
	}

	_, notFoundIds, err := repo.IsExistTaxRulesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxRules] failed checking taxRules whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.TaxRulesPrimaryID{}
		mapNotFoundIds := map[model.TaxRulesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "taxRules", fmt.Sprintf("taxRules with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsTaxRules(taxRulesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(taxRulesQueries.insertTaxRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxRules] failed exec create taxRules query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteTaxRulesByIDs(ctx context.Context, primaryIDs []model.TaxRulesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistTaxRulesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxRulesByIDs] failed checking taxRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxRules with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_rules\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(taxRulesQueries.deleteTaxRules + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxRulesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxRulesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistTaxRulesByIDs(ctx context.Context, ids []model.TaxRulesPrimaryID) (exists bool, notFoundIds []model.TaxRulesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_rules\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(taxRulesQueries.selectTaxRules, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxRulesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.TaxRulesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxRulesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.TaxRulesPrimaryID]bool{}
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

// BulkUpdateTaxRules is used to bulk update taxRules, by default it will update all field
// if want to update specific field, then fill taxRulessMapUpdateFieldsRequest else please fill taxRulessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateTaxRules(ctx context.Context, taxRulessMap map[model.TaxRulesPrimaryID]*model.TaxRules, taxRulessMapUpdateFieldsRequest map[model.TaxRulesPrimaryID]TaxRulesUpdateFieldList) (err error) {
	if len(taxRulessMap) == 0 && len(taxRulessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		taxRulessMapUpdateField map[model.TaxRulesPrimaryID]TaxRulesUpdateFieldList = map[model.TaxRulesPrimaryID]TaxRulesUpdateFieldList{}
		asTableValues           string                                              = "myvalues"
	)

	if len(taxRulessMap) > 0 {
		for id, taxRules := range taxRulessMap {
			if taxRules == nil {
				log.Error().Err(err).Msg("[BulkUpdateTaxRules] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			taxRulessMapUpdateField[id] = defaultTaxRulesUpdateFields(*taxRules)
		}
	} else {
		taxRulessMapUpdateField = taxRulessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateTaxRulesQuery(taxRulessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistTaxRulesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxRules] failed checking taxRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxRules with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeTaxRulesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"tax_rules\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxRules] failed exec query")
	}
	return
}

type TaxRulesFieldParameter struct {
	param string
	args  []interface{}
}

func NewTaxRulesFieldParameter(param string, args ...interface{}) TaxRulesFieldParameter {
	return TaxRulesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateTaxRulesQuery(mapTaxRuless map[model.TaxRulesPrimaryID]TaxRulesUpdateFieldList, asTableValues string) (primaryIDs []model.TaxRulesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.TaxRulesPrimaryID]map[string]interface{}{}
	taxRulesSelectFields := NewTaxRulesSelectFields()
	for id, updateFields := range mapTaxRuless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.taxRulesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapTaxRuless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetTaxRulesFieldType(updateField.taxRulesField)))
			args = append(args, fields[string(updateField.taxRulesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.taxRulesField))
		if updateField.taxRulesField == taxRulesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.taxRulesField, asTableValues, updateField.taxRulesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.taxRulesField,
				"\"tax_rules\"", updateField.taxRulesField,
				asTableValues, updateField.taxRulesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeTaxRulesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.TaxRulesPrimaryID, asTableValue string) (whereQry string) {
	taxRulesSelectFields := NewTaxRulesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"tax_rules\".\"id\" = %s.\"id\"::"+GetTaxRulesFieldType(taxRulesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetTaxRulesFieldType(taxRulesField TaxRulesField) string {
	selectTaxRulesFields := NewTaxRulesSelectFields()
	switch taxRulesField {

	case selectTaxRulesFields.Id():
		return "uuid"

	case selectTaxRulesFields.RuleCode():
		return "text"

	case selectTaxRulesFields.CountryCode():
		return "text"

	case selectTaxRulesFields.TaxType():
		return "text"

	case selectTaxRulesFields.Rate():
		return "numeric"

	case selectTaxRulesFields.TaxInclusive():
		return "bool"

	case selectTaxRulesFields.LiabilityPartyId():
		return "uuid"

	case selectTaxRulesFields.BeneficiaryPartyId():
		return "uuid"

	case selectTaxRulesFields.Priority():
		return "int4"

	case selectTaxRulesFields.IsActive():
		return "bool"

	case selectTaxRulesFields.ValidFrom():
		return "timestamptz"

	case selectTaxRulesFields.ValidUntil():
		return "timestamptz"

	case selectTaxRulesFields.Metadata():
		return "jsonb"

	case selectTaxRulesFields.MetaCreatedAt():
		return "timestamptz"

	case selectTaxRulesFields.MetaCreatedBy():
		return "uuid"

	case selectTaxRulesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectTaxRulesFields.MetaUpdatedBy():
		return "uuid"

	case selectTaxRulesFields.MetaDeletedAt():
		return "timestamptz"

	case selectTaxRulesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateTaxRules(ctx context.Context, taxRules *model.TaxRules, fieldsInsert ...TaxRulesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewTaxRulesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.TaxRulesPrimaryID{
		Id: taxRules.Id,
	}
	exists, err := repo.IsExistTaxRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxRules] failed checking taxRules whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "taxRules", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsTaxRules([]model.TaxRules{*taxRules}, fieldsInsert...)
	commandQuery := fmt.Sprintf(taxRulesQueries.insertTaxRules, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxRules] failed exec create taxRules query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteTaxRulesByID(ctx context.Context, primaryID model.TaxRulesPrimaryID) (err error) {
	exists, err := repo.IsExistTaxRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxRulesByID] failed checking taxRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeTaxRulesCompositePrimaryKeyWhere([]model.TaxRulesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(taxRulesQueries.deleteTaxRules + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxRulesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxRulesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxRulesFilterResult, err error) {
	query, args, err := composeTaxRulesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxRulesByFilter] failed compose taxRules filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxRulesByFilter] failed get taxRules by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeTaxRulesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.TaxRulesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeTaxRulesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeTaxRulesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeTaxRulesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewTaxRulesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["rule_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"rule_code\"")
			selectedColumns["rule_code"] = struct{}{}
		}
		if _, selected := selectedColumns["country_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"country_code\"")
			selectedColumns["country_code"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_type\"")
			selectedColumns["tax_type"] = struct{}{}
		}
		if _, selected := selectedColumns["rate"]; !selected {
			selectColumns = append(selectColumns, "base.\"rate\"")
			selectedColumns["rate"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_inclusive"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_inclusive\"")
			selectedColumns["tax_inclusive"] = struct{}{}
		}
		if _, selected := selectedColumns["liability_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"liability_party_id\"")
			selectedColumns["liability_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["beneficiary_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"beneficiary_party_id\"")
			selectedColumns["beneficiary_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["priority"]; !selected {
			selectColumns = append(selectColumns, "base.\"priority\"")
			selectedColumns["priority"] = struct{}{}
		}
		if _, selected := selectedColumns["is_active"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_active\"")
			selectedColumns["is_active"] = struct{}{}
		}
		if _, selected := selectedColumns["valid_from"]; !selected {
			selectColumns = append(selectColumns, "base.\"valid_from\"")
			selectedColumns["valid_from"] = struct{}{}
		}
		if _, selected := selectedColumns["valid_until"]; !selected {
			selectColumns = append(selectColumns, "base.\"valid_until\"")
			selectedColumns["valid_until"] = struct{}{}
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

type taxRulesFilterPlaceholder struct {
	index int
}

func (p *taxRulesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeTaxRulesFilterPredicate(filterField model.FilterField, placeholders *taxRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewTaxRulesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeTaxRulesFilterSQLExpr(spec)
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

func composeTaxRulesFilterGroup(group model.FilterGroup, placeholders *taxRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeTaxRulesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeTaxRulesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeTaxRulesFilterWhereQueries(filter model.Filter, placeholders *taxRulesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeTaxRulesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeTaxRulesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeTaxRulesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateTaxRulesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeTaxRulesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeTaxRulesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := taxRulesFilterPlaceholder{index: 1}
	whereQueries, err := composeTaxRulesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewTaxRulesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeTaxRulesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeTaxRulesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"tax_rules\" base%s", strings.Join(selectColumns, ","), composeTaxRulesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistTaxRulesByID(ctx context.Context, primaryID model.TaxRulesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeTaxRulesCompositePrimaryKeyWhere([]model.TaxRulesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", taxRulesQueries.selectCountTaxRules, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxRulesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxRules(ctx context.Context, selectFields ...TaxRulesField) (taxRulesList model.TaxRulesList, err error) {
	var (
		defaultTaxRulesSelectFields = defaultTaxRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxRulesSelectFields = composeTaxRulesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(taxRulesQueries.selectTaxRules, defaultTaxRulesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &taxRulesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxRules] failed get taxRules list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxRulesByID(ctx context.Context, primaryID model.TaxRulesPrimaryID, selectFields ...TaxRulesField) (taxRules model.TaxRules, err error) {
	var (
		defaultTaxRulesSelectFields = defaultTaxRulesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxRulesSelectFields = composeTaxRulesSelectFields(selectFields...)
	}
	whereQry, params := composeTaxRulesCompositePrimaryKeyWhere([]model.TaxRulesPrimaryID{primaryID})
	query := fmt.Sprintf(taxRulesQueries.selectTaxRules+" WHERE "+whereQry, defaultTaxRulesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &taxRules, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("taxRules with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveTaxRulesByID] failed get taxRules")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateTaxRulesByID(ctx context.Context, primaryID model.TaxRulesPrimaryID, taxRules *model.TaxRules, taxRulesUpdateFields ...TaxRulesUpdateField) (err error) {
	exists, err := repo.IsExistTaxRulesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxRules] failed checking taxRules whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxRules with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if taxRules == nil {
		if len(taxRulesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateTaxRulesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		taxRules = &model.TaxRules{}
	}
	var (
		defaultTaxRulesUpdateFields = defaultTaxRulesUpdateFields(*taxRules)
		tempUpdateField             TaxRulesUpdateFieldList
		selectFields                = NewTaxRulesSelectFields()
	)
	if len(taxRulesUpdateFields) > 0 {
		for _, updateField := range taxRulesUpdateFields {
			if updateField.taxRulesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultTaxRulesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeTaxRulesCompositePrimaryKeyWhere([]model.TaxRulesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsTaxRulesCommand(defaultTaxRulesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(taxRulesQueries.updateTaxRules+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxRules] error when try to update taxRules by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateTaxRulesByFilter(ctx context.Context, filter model.Filter, taxRulesUpdateFields ...TaxRulesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(taxRulesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields TaxRulesUpdateFieldList
		selectFields = NewTaxRulesSelectFields()
	)
	for _, updateField := range taxRulesUpdateFields {
		if updateField.taxRulesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsTaxRulesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := taxRulesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeTaxRulesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"tax_rules\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxRulesByFilter] error when try to update taxRules by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxRulesByFilter] failed get rows affected")
	}
	return
}

var (
	taxRulesQueries = struct {
		selectTaxRules      string
		selectCountTaxRules string
		deleteTaxRules      string
		updateTaxRules      string
		insertTaxRules      string
	}{
		selectTaxRules:      "SELECT %s FROM \"tax_rules\"",
		selectCountTaxRules: "SELECT COUNT(\"id\") FROM \"tax_rules\"",
		deleteTaxRules:      "DELETE FROM \"tax_rules\"",
		updateTaxRules:      "UPDATE \"tax_rules\" SET %s ",
		insertTaxRules:      "INSERT INTO \"tax_rules\" %s VALUES %s",
	}
)

type TaxRulesRepository interface {
	CreateTaxRules(ctx context.Context, taxRules *model.TaxRules, fieldsInsert ...TaxRulesField) error
	BulkCreateTaxRules(ctx context.Context, taxRulesList []*model.TaxRules, fieldsInsert ...TaxRulesField) error
	ResolveTaxRules(ctx context.Context, selectFields ...TaxRulesField) (model.TaxRulesList, error)
	ResolveTaxRulesByID(ctx context.Context, primaryID model.TaxRulesPrimaryID, selectFields ...TaxRulesField) (model.TaxRules, error)
	UpdateTaxRulesByID(ctx context.Context, id model.TaxRulesPrimaryID, taxRules *model.TaxRules, taxRulesUpdateFields ...TaxRulesUpdateField) error
	UpdateTaxRulesByFilter(ctx context.Context, filter model.Filter, taxRulesUpdateFields ...TaxRulesUpdateField) (rowsAffected int64, err error)
	BulkUpdateTaxRules(ctx context.Context, taxRulesListMap map[model.TaxRulesPrimaryID]*model.TaxRules, TaxRulessMapUpdateFieldsRequest map[model.TaxRulesPrimaryID]TaxRulesUpdateFieldList) (err error)
	DeleteTaxRulesByID(ctx context.Context, id model.TaxRulesPrimaryID) error
	BulkDeleteTaxRulesByIDs(ctx context.Context, ids []model.TaxRulesPrimaryID) error
	ResolveTaxRulesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxRulesFilterResult, err error)
	IsExistTaxRulesByIDs(ctx context.Context, ids []model.TaxRulesPrimaryID) (exists bool, notFoundIds []model.TaxRulesPrimaryID, err error)
	IsExistTaxRulesByID(ctx context.Context, id model.TaxRulesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
