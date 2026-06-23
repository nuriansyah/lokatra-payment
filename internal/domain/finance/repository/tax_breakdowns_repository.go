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

func composeInsertFieldsAndParamsTaxBreakdowns(taxBreakdownsList []model.TaxBreakdowns, fieldsInsert ...TaxBreakdownsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewTaxBreakdownsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, taxBreakdowns := range taxBreakdownsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, taxBreakdowns.Id)
			case selectField.SourceType():
				args = append(args, taxBreakdowns.SourceType)
			case selectField.SourceId():
				args = append(args, taxBreakdowns.SourceId)
			case selectField.TaxRuleId():
				args = append(args, taxBreakdowns.TaxRuleId)
			case selectField.CurrencyCode():
				args = append(args, taxBreakdowns.CurrencyCode)
			case selectField.TaxableAmount():
				args = append(args, taxBreakdowns.TaxableAmount)
			case selectField.TaxAmount():
				args = append(args, taxBreakdowns.TaxAmount)
			case selectField.LiabilityPartyId():
				args = append(args, taxBreakdowns.LiabilityPartyId)
			case selectField.BeneficiaryPartyId():
				args = append(args, taxBreakdowns.BeneficiaryPartyId)
			case selectField.BreakdownStatus():
				args = append(args, taxBreakdowns.BreakdownStatus)
			case selectField.Metadata():
				args = append(args, taxBreakdowns.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, taxBreakdowns.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, taxBreakdowns.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, taxBreakdowns.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, taxBreakdowns.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, taxBreakdowns.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, taxBreakdowns.MetaDeletedBy)

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

func composeTaxBreakdownsCompositePrimaryKeyWhere(primaryIDs []model.TaxBreakdownsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"tax_breakdowns\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultTaxBreakdownsSelectFields() string {
	fields := NewTaxBreakdownsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeTaxBreakdownsSelectFields(selectFields ...TaxBreakdownsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type TaxBreakdownsField string
type TaxBreakdownsFieldList []TaxBreakdownsField

type TaxBreakdownsSelectFields struct {
}

func (ss TaxBreakdownsSelectFields) Id() TaxBreakdownsField {
	return TaxBreakdownsField("id")
}

func (ss TaxBreakdownsSelectFields) SourceType() TaxBreakdownsField {
	return TaxBreakdownsField("source_type")
}

func (ss TaxBreakdownsSelectFields) SourceId() TaxBreakdownsField {
	return TaxBreakdownsField("source_id")
}

func (ss TaxBreakdownsSelectFields) TaxRuleId() TaxBreakdownsField {
	return TaxBreakdownsField("tax_rule_id")
}

func (ss TaxBreakdownsSelectFields) CurrencyCode() TaxBreakdownsField {
	return TaxBreakdownsField("currency_code")
}

func (ss TaxBreakdownsSelectFields) TaxableAmount() TaxBreakdownsField {
	return TaxBreakdownsField("taxable_amount")
}

func (ss TaxBreakdownsSelectFields) TaxAmount() TaxBreakdownsField {
	return TaxBreakdownsField("tax_amount")
}

func (ss TaxBreakdownsSelectFields) LiabilityPartyId() TaxBreakdownsField {
	return TaxBreakdownsField("liability_party_id")
}

func (ss TaxBreakdownsSelectFields) BeneficiaryPartyId() TaxBreakdownsField {
	return TaxBreakdownsField("beneficiary_party_id")
}

func (ss TaxBreakdownsSelectFields) BreakdownStatus() TaxBreakdownsField {
	return TaxBreakdownsField("breakdown_status")
}

func (ss TaxBreakdownsSelectFields) Metadata() TaxBreakdownsField {
	return TaxBreakdownsField("metadata")
}

func (ss TaxBreakdownsSelectFields) MetaCreatedAt() TaxBreakdownsField {
	return TaxBreakdownsField("meta_created_at")
}

func (ss TaxBreakdownsSelectFields) MetaCreatedBy() TaxBreakdownsField {
	return TaxBreakdownsField("meta_created_by")
}

func (ss TaxBreakdownsSelectFields) MetaUpdatedAt() TaxBreakdownsField {
	return TaxBreakdownsField("meta_updated_at")
}

func (ss TaxBreakdownsSelectFields) MetaUpdatedBy() TaxBreakdownsField {
	return TaxBreakdownsField("meta_updated_by")
}

func (ss TaxBreakdownsSelectFields) MetaDeletedAt() TaxBreakdownsField {
	return TaxBreakdownsField("meta_deleted_at")
}

func (ss TaxBreakdownsSelectFields) MetaDeletedBy() TaxBreakdownsField {
	return TaxBreakdownsField("meta_deleted_by")
}

func (ss TaxBreakdownsSelectFields) All() TaxBreakdownsFieldList {
	return []TaxBreakdownsField{
		ss.Id(),
		ss.SourceType(),
		ss.SourceId(),
		ss.TaxRuleId(),
		ss.CurrencyCode(),
		ss.TaxableAmount(),
		ss.TaxAmount(),
		ss.LiabilityPartyId(),
		ss.BeneficiaryPartyId(),
		ss.BreakdownStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewTaxBreakdownsSelectFields() TaxBreakdownsSelectFields {
	return TaxBreakdownsSelectFields{}
}

type TaxBreakdownsUpdateFieldOption struct {
	useIncrement bool
}
type TaxBreakdownsUpdateField struct {
	taxBreakdownsField TaxBreakdownsField
	opt                TaxBreakdownsUpdateFieldOption
	value              interface{}
}
type TaxBreakdownsUpdateFieldList []TaxBreakdownsUpdateField

func defaultTaxBreakdownsUpdateFieldOption() TaxBreakdownsUpdateFieldOption {
	return TaxBreakdownsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementTaxBreakdownsOption(useIncrement bool) func(*TaxBreakdownsUpdateFieldOption) {
	return func(pcufo *TaxBreakdownsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewTaxBreakdownsUpdateField(field TaxBreakdownsField, val interface{}, opts ...func(*TaxBreakdownsUpdateFieldOption)) TaxBreakdownsUpdateField {
	defaultOpt := defaultTaxBreakdownsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return TaxBreakdownsUpdateField{
		taxBreakdownsField: field,
		value:              val,
		opt:                defaultOpt,
	}
}
func defaultTaxBreakdownsUpdateFields(taxBreakdowns model.TaxBreakdowns) (taxBreakdownsUpdateFieldList TaxBreakdownsUpdateFieldList) {
	selectFields := NewTaxBreakdownsSelectFields()
	taxBreakdownsUpdateFieldList = append(taxBreakdownsUpdateFieldList,
		NewTaxBreakdownsUpdateField(selectFields.Id(), taxBreakdowns.Id),
		NewTaxBreakdownsUpdateField(selectFields.SourceType(), taxBreakdowns.SourceType),
		NewTaxBreakdownsUpdateField(selectFields.SourceId(), taxBreakdowns.SourceId),
		NewTaxBreakdownsUpdateField(selectFields.TaxRuleId(), taxBreakdowns.TaxRuleId),
		NewTaxBreakdownsUpdateField(selectFields.CurrencyCode(), taxBreakdowns.CurrencyCode),
		NewTaxBreakdownsUpdateField(selectFields.TaxableAmount(), taxBreakdowns.TaxableAmount),
		NewTaxBreakdownsUpdateField(selectFields.TaxAmount(), taxBreakdowns.TaxAmount),
		NewTaxBreakdownsUpdateField(selectFields.LiabilityPartyId(), taxBreakdowns.LiabilityPartyId),
		NewTaxBreakdownsUpdateField(selectFields.BeneficiaryPartyId(), taxBreakdowns.BeneficiaryPartyId),
		NewTaxBreakdownsUpdateField(selectFields.BreakdownStatus(), taxBreakdowns.BreakdownStatus),
		NewTaxBreakdownsUpdateField(selectFields.Metadata(), taxBreakdowns.Metadata),
		NewTaxBreakdownsUpdateField(selectFields.MetaCreatedAt(), taxBreakdowns.MetaCreatedAt),
		NewTaxBreakdownsUpdateField(selectFields.MetaCreatedBy(), taxBreakdowns.MetaCreatedBy),
		NewTaxBreakdownsUpdateField(selectFields.MetaUpdatedAt(), taxBreakdowns.MetaUpdatedAt),
		NewTaxBreakdownsUpdateField(selectFields.MetaUpdatedBy(), taxBreakdowns.MetaUpdatedBy),
		NewTaxBreakdownsUpdateField(selectFields.MetaDeletedAt(), taxBreakdowns.MetaDeletedAt),
		NewTaxBreakdownsUpdateField(selectFields.MetaDeletedBy(), taxBreakdowns.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsTaxBreakdownsCommand(taxBreakdownsUpdateFieldList TaxBreakdownsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range taxBreakdownsUpdateFieldList {
		field := string(updateField.taxBreakdownsField)
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

func (repo *RepositoryImpl) BulkCreateTaxBreakdowns(ctx context.Context, taxBreakdownsList []*model.TaxBreakdowns, fieldsInsert ...TaxBreakdownsField) (err error) {
	var (
		fieldsStr              string
		valueListStr           []string
		argsList               []interface{}
		primaryIds             []model.TaxBreakdownsPrimaryID
		taxBreakdownsValueList []model.TaxBreakdowns
	)

	if len(fieldsInsert) == 0 {
		selectField := NewTaxBreakdownsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, taxBreakdowns := range taxBreakdownsList {

		primaryIds = append(primaryIds, taxBreakdowns.ToTaxBreakdownsPrimaryID())

		taxBreakdownsValueList = append(taxBreakdownsValueList, *taxBreakdowns)
	}

	_, notFoundIds, err := repo.IsExistTaxBreakdownsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxBreakdowns] failed checking taxBreakdowns whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.TaxBreakdownsPrimaryID{}
		mapNotFoundIds := map[model.TaxBreakdownsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "taxBreakdowns", fmt.Sprintf("taxBreakdowns with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsTaxBreakdowns(taxBreakdownsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(taxBreakdownsQueries.insertTaxBreakdowns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxBreakdowns] failed exec create taxBreakdowns query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteTaxBreakdownsByIDs(ctx context.Context, primaryIDs []model.TaxBreakdownsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistTaxBreakdownsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxBreakdownsByIDs] failed checking taxBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxBreakdowns with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_breakdowns\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(taxBreakdownsQueries.deleteTaxBreakdowns + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxBreakdownsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxBreakdownsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistTaxBreakdownsByIDs(ctx context.Context, ids []model.TaxBreakdownsPrimaryID) (exists bool, notFoundIds []model.TaxBreakdownsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_breakdowns\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(taxBreakdownsQueries.selectTaxBreakdowns, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxBreakdownsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.TaxBreakdownsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxBreakdownsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.TaxBreakdownsPrimaryID]bool{}
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

// BulkUpdateTaxBreakdowns is used to bulk update taxBreakdowns, by default it will update all field
// if want to update specific field, then fill taxBreakdownssMapUpdateFieldsRequest else please fill taxBreakdownssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateTaxBreakdowns(ctx context.Context, taxBreakdownssMap map[model.TaxBreakdownsPrimaryID]*model.TaxBreakdowns, taxBreakdownssMapUpdateFieldsRequest map[model.TaxBreakdownsPrimaryID]TaxBreakdownsUpdateFieldList) (err error) {
	if len(taxBreakdownssMap) == 0 && len(taxBreakdownssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		taxBreakdownssMapUpdateField map[model.TaxBreakdownsPrimaryID]TaxBreakdownsUpdateFieldList = map[model.TaxBreakdownsPrimaryID]TaxBreakdownsUpdateFieldList{}
		asTableValues                string                                                        = "myvalues"
	)

	if len(taxBreakdownssMap) > 0 {
		for id, taxBreakdowns := range taxBreakdownssMap {
			if taxBreakdowns == nil {
				log.Error().Err(err).Msg("[BulkUpdateTaxBreakdowns] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			taxBreakdownssMapUpdateField[id] = defaultTaxBreakdownsUpdateFields(*taxBreakdowns)
		}
	} else {
		taxBreakdownssMapUpdateField = taxBreakdownssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateTaxBreakdownsQuery(taxBreakdownssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistTaxBreakdownsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxBreakdowns] failed checking taxBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxBreakdowns with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeTaxBreakdownsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"tax_breakdowns\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxBreakdowns] failed exec query")
	}
	return
}

type TaxBreakdownsFieldParameter struct {
	param string
	args  []interface{}
}

func NewTaxBreakdownsFieldParameter(param string, args ...interface{}) TaxBreakdownsFieldParameter {
	return TaxBreakdownsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateTaxBreakdownsQuery(mapTaxBreakdownss map[model.TaxBreakdownsPrimaryID]TaxBreakdownsUpdateFieldList, asTableValues string) (primaryIDs []model.TaxBreakdownsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.TaxBreakdownsPrimaryID]map[string]interface{}{}
	taxBreakdownsSelectFields := NewTaxBreakdownsSelectFields()
	for id, updateFields := range mapTaxBreakdownss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.taxBreakdownsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapTaxBreakdownss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetTaxBreakdownsFieldType(updateField.taxBreakdownsField)))
			args = append(args, fields[string(updateField.taxBreakdownsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.taxBreakdownsField))
		if updateField.taxBreakdownsField == taxBreakdownsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.taxBreakdownsField, asTableValues, updateField.taxBreakdownsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.taxBreakdownsField,
				"\"tax_breakdowns\"", updateField.taxBreakdownsField,
				asTableValues, updateField.taxBreakdownsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeTaxBreakdownsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.TaxBreakdownsPrimaryID, asTableValue string) (whereQry string) {
	taxBreakdownsSelectFields := NewTaxBreakdownsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"tax_breakdowns\".\"id\" = %s.\"id\"::"+GetTaxBreakdownsFieldType(taxBreakdownsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetTaxBreakdownsFieldType(taxBreakdownsField TaxBreakdownsField) string {
	selectTaxBreakdownsFields := NewTaxBreakdownsSelectFields()
	switch taxBreakdownsField {

	case selectTaxBreakdownsFields.Id():
		return "uuid"

	case selectTaxBreakdownsFields.SourceType():
		return "text"

	case selectTaxBreakdownsFields.SourceId():
		return "uuid"

	case selectTaxBreakdownsFields.TaxRuleId():
		return "uuid"

	case selectTaxBreakdownsFields.CurrencyCode():
		return "text"

	case selectTaxBreakdownsFields.TaxableAmount():
		return "numeric"

	case selectTaxBreakdownsFields.TaxAmount():
		return "numeric"

	case selectTaxBreakdownsFields.LiabilityPartyId():
		return "uuid"

	case selectTaxBreakdownsFields.BeneficiaryPartyId():
		return "uuid"

	case selectTaxBreakdownsFields.BreakdownStatus():
		return "tax_breakdowns_breakdown_status_enum"

	case selectTaxBreakdownsFields.Metadata():
		return "jsonb"

	case selectTaxBreakdownsFields.MetaCreatedAt():
		return "timestamptz"

	case selectTaxBreakdownsFields.MetaCreatedBy():
		return "uuid"

	case selectTaxBreakdownsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectTaxBreakdownsFields.MetaUpdatedBy():
		return "uuid"

	case selectTaxBreakdownsFields.MetaDeletedAt():
		return "timestamptz"

	case selectTaxBreakdownsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateTaxBreakdowns(ctx context.Context, taxBreakdowns *model.TaxBreakdowns, fieldsInsert ...TaxBreakdownsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewTaxBreakdownsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.TaxBreakdownsPrimaryID{
		Id: taxBreakdowns.Id,
	}
	exists, err := repo.IsExistTaxBreakdownsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxBreakdowns] failed checking taxBreakdowns whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "taxBreakdowns", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsTaxBreakdowns([]model.TaxBreakdowns{*taxBreakdowns}, fieldsInsert...)
	commandQuery := fmt.Sprintf(taxBreakdownsQueries.insertTaxBreakdowns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxBreakdowns] failed exec create taxBreakdowns query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteTaxBreakdownsByID(ctx context.Context, primaryID model.TaxBreakdownsPrimaryID) (err error) {
	exists, err := repo.IsExistTaxBreakdownsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxBreakdownsByID] failed checking taxBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxBreakdowns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeTaxBreakdownsCompositePrimaryKeyWhere([]model.TaxBreakdownsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(taxBreakdownsQueries.deleteTaxBreakdowns + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxBreakdownsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxBreakdownsByFilter(ctx context.Context, filter model.Filter) (result []model.TaxBreakdownsFilterResult, err error) {
	query, args, err := composeTaxBreakdownsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxBreakdownsByFilter] failed compose taxBreakdowns filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxBreakdownsByFilter] failed get taxBreakdowns by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeTaxBreakdownsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.TaxBreakdownsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeTaxBreakdownsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeTaxBreakdownsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeTaxBreakdownsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 17 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewTaxBreakdownsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_rule_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_rule_id\"")
			selectedColumns["tax_rule_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["taxable_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"taxable_amount\"")
			selectedColumns["taxable_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_amount\"")
			selectedColumns["tax_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["liability_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"liability_party_id\"")
			selectedColumns["liability_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["beneficiary_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"beneficiary_party_id\"")
			selectedColumns["beneficiary_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["breakdown_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"breakdown_status\"")
			selectedColumns["breakdown_status"] = struct{}{}
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

type taxBreakdownsFilterPlaceholder struct {
	index int
}

func (p *taxBreakdownsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeTaxBreakdownsFilterPredicate(filterField model.FilterField, placeholders *taxBreakdownsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewTaxBreakdownsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeTaxBreakdownsFilterSQLExpr(spec)
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

func composeTaxBreakdownsFilterGroup(group model.FilterGroup, placeholders *taxBreakdownsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeTaxBreakdownsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeTaxBreakdownsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeTaxBreakdownsFilterWhereQueries(filter model.Filter, placeholders *taxBreakdownsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeTaxBreakdownsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeTaxBreakdownsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeTaxBreakdownsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateTaxBreakdownsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeTaxBreakdownsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeTaxBreakdownsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := taxBreakdownsFilterPlaceholder{index: 1}
	whereQueries, err := composeTaxBreakdownsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewTaxBreakdownsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeTaxBreakdownsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeTaxBreakdownsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"tax_breakdowns\" base%s", strings.Join(selectColumns, ","), composeTaxBreakdownsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistTaxBreakdownsByID(ctx context.Context, primaryID model.TaxBreakdownsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeTaxBreakdownsCompositePrimaryKeyWhere([]model.TaxBreakdownsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", taxBreakdownsQueries.selectCountTaxBreakdowns, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxBreakdownsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxBreakdowns(ctx context.Context, selectFields ...TaxBreakdownsField) (taxBreakdownsList model.TaxBreakdownsList, err error) {
	var (
		defaultTaxBreakdownsSelectFields = defaultTaxBreakdownsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxBreakdownsSelectFields = composeTaxBreakdownsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(taxBreakdownsQueries.selectTaxBreakdowns, defaultTaxBreakdownsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &taxBreakdownsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxBreakdowns] failed get taxBreakdowns list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxBreakdownsByID(ctx context.Context, primaryID model.TaxBreakdownsPrimaryID, selectFields ...TaxBreakdownsField) (taxBreakdowns model.TaxBreakdowns, err error) {
	var (
		defaultTaxBreakdownsSelectFields = defaultTaxBreakdownsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxBreakdownsSelectFields = composeTaxBreakdownsSelectFields(selectFields...)
	}
	whereQry, params := composeTaxBreakdownsCompositePrimaryKeyWhere([]model.TaxBreakdownsPrimaryID{primaryID})
	query := fmt.Sprintf(taxBreakdownsQueries.selectTaxBreakdowns+" WHERE "+whereQry, defaultTaxBreakdownsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &taxBreakdowns, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("taxBreakdowns with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveTaxBreakdownsByID] failed get taxBreakdowns")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateTaxBreakdownsByID(ctx context.Context, primaryID model.TaxBreakdownsPrimaryID, taxBreakdowns *model.TaxBreakdowns, taxBreakdownsUpdateFields ...TaxBreakdownsUpdateField) (err error) {
	exists, err := repo.IsExistTaxBreakdownsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxBreakdowns] failed checking taxBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxBreakdowns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if taxBreakdowns == nil {
		if len(taxBreakdownsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateTaxBreakdownsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		taxBreakdowns = &model.TaxBreakdowns{}
	}
	var (
		defaultTaxBreakdownsUpdateFields = defaultTaxBreakdownsUpdateFields(*taxBreakdowns)
		tempUpdateField                  TaxBreakdownsUpdateFieldList
		selectFields                     = NewTaxBreakdownsSelectFields()
	)
	if len(taxBreakdownsUpdateFields) > 0 {
		for _, updateField := range taxBreakdownsUpdateFields {
			if updateField.taxBreakdownsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultTaxBreakdownsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeTaxBreakdownsCompositePrimaryKeyWhere([]model.TaxBreakdownsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsTaxBreakdownsCommand(defaultTaxBreakdownsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(taxBreakdownsQueries.updateTaxBreakdowns+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxBreakdowns] error when try to update taxBreakdowns by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateTaxBreakdownsByFilter(ctx context.Context, filter model.Filter, taxBreakdownsUpdateFields ...TaxBreakdownsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(taxBreakdownsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields TaxBreakdownsUpdateFieldList
		selectFields = NewTaxBreakdownsSelectFields()
	)
	for _, updateField := range taxBreakdownsUpdateFields {
		if updateField.taxBreakdownsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsTaxBreakdownsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := taxBreakdownsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeTaxBreakdownsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"tax_breakdowns\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxBreakdownsByFilter] error when try to update taxBreakdowns by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxBreakdownsByFilter] failed get rows affected")
	}
	return
}

var (
	taxBreakdownsQueries = struct {
		selectTaxBreakdowns      string
		selectCountTaxBreakdowns string
		deleteTaxBreakdowns      string
		updateTaxBreakdowns      string
		insertTaxBreakdowns      string
	}{
		selectTaxBreakdowns:      "SELECT %s FROM \"tax_breakdowns\"",
		selectCountTaxBreakdowns: "SELECT COUNT(\"id\") FROM \"tax_breakdowns\"",
		deleteTaxBreakdowns:      "DELETE FROM \"tax_breakdowns\"",
		updateTaxBreakdowns:      "UPDATE \"tax_breakdowns\" SET %s ",
		insertTaxBreakdowns:      "INSERT INTO \"tax_breakdowns\" %s VALUES %s",
	}
)

type TaxBreakdownsRepository interface {
	CreateTaxBreakdowns(ctx context.Context, taxBreakdowns *model.TaxBreakdowns, fieldsInsert ...TaxBreakdownsField) error
	BulkCreateTaxBreakdowns(ctx context.Context, taxBreakdownsList []*model.TaxBreakdowns, fieldsInsert ...TaxBreakdownsField) error
	ResolveTaxBreakdowns(ctx context.Context, selectFields ...TaxBreakdownsField) (model.TaxBreakdownsList, error)
	ResolveTaxBreakdownsByID(ctx context.Context, primaryID model.TaxBreakdownsPrimaryID, selectFields ...TaxBreakdownsField) (model.TaxBreakdowns, error)
	UpdateTaxBreakdownsByID(ctx context.Context, id model.TaxBreakdownsPrimaryID, taxBreakdowns *model.TaxBreakdowns, taxBreakdownsUpdateFields ...TaxBreakdownsUpdateField) error
	UpdateTaxBreakdownsByFilter(ctx context.Context, filter model.Filter, taxBreakdownsUpdateFields ...TaxBreakdownsUpdateField) (rowsAffected int64, err error)
	BulkUpdateTaxBreakdowns(ctx context.Context, taxBreakdownsListMap map[model.TaxBreakdownsPrimaryID]*model.TaxBreakdowns, TaxBreakdownssMapUpdateFieldsRequest map[model.TaxBreakdownsPrimaryID]TaxBreakdownsUpdateFieldList) (err error)
	DeleteTaxBreakdownsByID(ctx context.Context, id model.TaxBreakdownsPrimaryID) error
	BulkDeleteTaxBreakdownsByIDs(ctx context.Context, ids []model.TaxBreakdownsPrimaryID) error
	ResolveTaxBreakdownsByFilter(ctx context.Context, filter model.Filter) (result []model.TaxBreakdownsFilterResult, err error)
	IsExistTaxBreakdownsByIDs(ctx context.Context, ids []model.TaxBreakdownsPrimaryID) (exists bool, notFoundIds []model.TaxBreakdownsPrimaryID, err error)
	IsExistTaxBreakdownsByID(ctx context.Context, id model.TaxBreakdownsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
