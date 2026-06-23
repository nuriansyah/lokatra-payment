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

func composeInsertFieldsAndParamsTaxProfiles(taxProfilesList []model.TaxProfiles, fieldsInsert ...TaxProfilesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewTaxProfilesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, taxProfiles := range taxProfilesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, taxProfiles.Id)
			case selectField.OwnerPartyId():
				args = append(args, taxProfiles.OwnerPartyId)
			case selectField.CountryCode():
				args = append(args, taxProfiles.CountryCode)
			case selectField.TaxResidencyCountry():
				args = append(args, taxProfiles.TaxResidencyCountry)
			case selectField.TaxIdMasked():
				args = append(args, taxProfiles.TaxIdMasked)
			case selectField.TaxEntityType():
				args = append(args, taxProfiles.TaxEntityType)
			case selectField.IsVatRegistered():
				args = append(args, taxProfiles.IsVatRegistered)
			case selectField.IsWithholdingSubject():
				args = append(args, taxProfiles.IsWithholdingSubject)
			case selectField.ProfileStatus():
				args = append(args, taxProfiles.ProfileStatus)
			case selectField.Metadata():
				args = append(args, taxProfiles.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, taxProfiles.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, taxProfiles.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, taxProfiles.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, taxProfiles.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, taxProfiles.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, taxProfiles.MetaDeletedBy)

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

func composeTaxProfilesCompositePrimaryKeyWhere(primaryIDs []model.TaxProfilesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"tax_profiles\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultTaxProfilesSelectFields() string {
	fields := NewTaxProfilesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeTaxProfilesSelectFields(selectFields ...TaxProfilesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type TaxProfilesField string
type TaxProfilesFieldList []TaxProfilesField

type TaxProfilesSelectFields struct {
}

func (ss TaxProfilesSelectFields) Id() TaxProfilesField {
	return TaxProfilesField("id")
}

func (ss TaxProfilesSelectFields) OwnerPartyId() TaxProfilesField {
	return TaxProfilesField("owner_party_id")
}

func (ss TaxProfilesSelectFields) CountryCode() TaxProfilesField {
	return TaxProfilesField("country_code")
}

func (ss TaxProfilesSelectFields) TaxResidencyCountry() TaxProfilesField {
	return TaxProfilesField("tax_residency_country")
}

func (ss TaxProfilesSelectFields) TaxIdMasked() TaxProfilesField {
	return TaxProfilesField("tax_id_masked")
}

func (ss TaxProfilesSelectFields) TaxEntityType() TaxProfilesField {
	return TaxProfilesField("tax_entity_type")
}

func (ss TaxProfilesSelectFields) IsVatRegistered() TaxProfilesField {
	return TaxProfilesField("is_vat_registered")
}

func (ss TaxProfilesSelectFields) IsWithholdingSubject() TaxProfilesField {
	return TaxProfilesField("is_withholding_subject")
}

func (ss TaxProfilesSelectFields) ProfileStatus() TaxProfilesField {
	return TaxProfilesField("profile_status")
}

func (ss TaxProfilesSelectFields) Metadata() TaxProfilesField {
	return TaxProfilesField("metadata")
}

func (ss TaxProfilesSelectFields) MetaCreatedAt() TaxProfilesField {
	return TaxProfilesField("meta_created_at")
}

func (ss TaxProfilesSelectFields) MetaCreatedBy() TaxProfilesField {
	return TaxProfilesField("meta_created_by")
}

func (ss TaxProfilesSelectFields) MetaUpdatedAt() TaxProfilesField {
	return TaxProfilesField("meta_updated_at")
}

func (ss TaxProfilesSelectFields) MetaUpdatedBy() TaxProfilesField {
	return TaxProfilesField("meta_updated_by")
}

func (ss TaxProfilesSelectFields) MetaDeletedAt() TaxProfilesField {
	return TaxProfilesField("meta_deleted_at")
}

func (ss TaxProfilesSelectFields) MetaDeletedBy() TaxProfilesField {
	return TaxProfilesField("meta_deleted_by")
}

func (ss TaxProfilesSelectFields) All() TaxProfilesFieldList {
	return []TaxProfilesField{
		ss.Id(),
		ss.OwnerPartyId(),
		ss.CountryCode(),
		ss.TaxResidencyCountry(),
		ss.TaxIdMasked(),
		ss.TaxEntityType(),
		ss.IsVatRegistered(),
		ss.IsWithholdingSubject(),
		ss.ProfileStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewTaxProfilesSelectFields() TaxProfilesSelectFields {
	return TaxProfilesSelectFields{}
}

type TaxProfilesUpdateFieldOption struct {
	useIncrement bool
}
type TaxProfilesUpdateField struct {
	taxProfilesField TaxProfilesField
	opt              TaxProfilesUpdateFieldOption
	value            interface{}
}
type TaxProfilesUpdateFieldList []TaxProfilesUpdateField

func defaultTaxProfilesUpdateFieldOption() TaxProfilesUpdateFieldOption {
	return TaxProfilesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementTaxProfilesOption(useIncrement bool) func(*TaxProfilesUpdateFieldOption) {
	return func(pcufo *TaxProfilesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewTaxProfilesUpdateField(field TaxProfilesField, val interface{}, opts ...func(*TaxProfilesUpdateFieldOption)) TaxProfilesUpdateField {
	defaultOpt := defaultTaxProfilesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return TaxProfilesUpdateField{
		taxProfilesField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultTaxProfilesUpdateFields(taxProfiles model.TaxProfiles) (taxProfilesUpdateFieldList TaxProfilesUpdateFieldList) {
	selectFields := NewTaxProfilesSelectFields()
	taxProfilesUpdateFieldList = append(taxProfilesUpdateFieldList,
		NewTaxProfilesUpdateField(selectFields.Id(), taxProfiles.Id),
		NewTaxProfilesUpdateField(selectFields.OwnerPartyId(), taxProfiles.OwnerPartyId),
		NewTaxProfilesUpdateField(selectFields.CountryCode(), taxProfiles.CountryCode),
		NewTaxProfilesUpdateField(selectFields.TaxResidencyCountry(), taxProfiles.TaxResidencyCountry),
		NewTaxProfilesUpdateField(selectFields.TaxIdMasked(), taxProfiles.TaxIdMasked),
		NewTaxProfilesUpdateField(selectFields.TaxEntityType(), taxProfiles.TaxEntityType),
		NewTaxProfilesUpdateField(selectFields.IsVatRegistered(), taxProfiles.IsVatRegistered),
		NewTaxProfilesUpdateField(selectFields.IsWithholdingSubject(), taxProfiles.IsWithholdingSubject),
		NewTaxProfilesUpdateField(selectFields.ProfileStatus(), taxProfiles.ProfileStatus),
		NewTaxProfilesUpdateField(selectFields.Metadata(), taxProfiles.Metadata),
		NewTaxProfilesUpdateField(selectFields.MetaCreatedAt(), taxProfiles.MetaCreatedAt),
		NewTaxProfilesUpdateField(selectFields.MetaCreatedBy(), taxProfiles.MetaCreatedBy),
		NewTaxProfilesUpdateField(selectFields.MetaUpdatedAt(), taxProfiles.MetaUpdatedAt),
		NewTaxProfilesUpdateField(selectFields.MetaUpdatedBy(), taxProfiles.MetaUpdatedBy),
		NewTaxProfilesUpdateField(selectFields.MetaDeletedAt(), taxProfiles.MetaDeletedAt),
		NewTaxProfilesUpdateField(selectFields.MetaDeletedBy(), taxProfiles.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsTaxProfilesCommand(taxProfilesUpdateFieldList TaxProfilesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range taxProfilesUpdateFieldList {
		field := string(updateField.taxProfilesField)
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

func (repo *RepositoryImpl) BulkCreateTaxProfiles(ctx context.Context, taxProfilesList []*model.TaxProfiles, fieldsInsert ...TaxProfilesField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.TaxProfilesPrimaryID
		taxProfilesValueList []model.TaxProfiles
	)

	if len(fieldsInsert) == 0 {
		selectField := NewTaxProfilesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, taxProfiles := range taxProfilesList {

		primaryIds = append(primaryIds, taxProfiles.ToTaxProfilesPrimaryID())

		taxProfilesValueList = append(taxProfilesValueList, *taxProfiles)
	}

	_, notFoundIds, err := repo.IsExistTaxProfilesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxProfiles] failed checking taxProfiles whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.TaxProfilesPrimaryID{}
		mapNotFoundIds := map[model.TaxProfilesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "taxProfiles", fmt.Sprintf("taxProfiles with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsTaxProfiles(taxProfilesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(taxProfilesQueries.insertTaxProfiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxProfiles] failed exec create taxProfiles query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteTaxProfilesByIDs(ctx context.Context, primaryIDs []model.TaxProfilesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistTaxProfilesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxProfilesByIDs] failed checking taxProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxProfiles with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_profiles\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(taxProfilesQueries.deleteTaxProfiles + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxProfilesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxProfilesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistTaxProfilesByIDs(ctx context.Context, ids []model.TaxProfilesPrimaryID) (exists bool, notFoundIds []model.TaxProfilesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_profiles\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(taxProfilesQueries.selectTaxProfiles, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxProfilesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.TaxProfilesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxProfilesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.TaxProfilesPrimaryID]bool{}
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

// BulkUpdateTaxProfiles is used to bulk update taxProfiles, by default it will update all field
// if want to update specific field, then fill taxProfilessMapUpdateFieldsRequest else please fill taxProfilessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateTaxProfiles(ctx context.Context, taxProfilessMap map[model.TaxProfilesPrimaryID]*model.TaxProfiles, taxProfilessMapUpdateFieldsRequest map[model.TaxProfilesPrimaryID]TaxProfilesUpdateFieldList) (err error) {
	if len(taxProfilessMap) == 0 && len(taxProfilessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		taxProfilessMapUpdateField map[model.TaxProfilesPrimaryID]TaxProfilesUpdateFieldList = map[model.TaxProfilesPrimaryID]TaxProfilesUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(taxProfilessMap) > 0 {
		for id, taxProfiles := range taxProfilessMap {
			if taxProfiles == nil {
				log.Error().Err(err).Msg("[BulkUpdateTaxProfiles] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			taxProfilessMapUpdateField[id] = defaultTaxProfilesUpdateFields(*taxProfiles)
		}
	} else {
		taxProfilessMapUpdateField = taxProfilessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateTaxProfilesQuery(taxProfilessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistTaxProfilesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxProfiles] failed checking taxProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxProfiles with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeTaxProfilesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"tax_profiles\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxProfiles] failed exec query")
	}
	return
}

type TaxProfilesFieldParameter struct {
	param string
	args  []interface{}
}

func NewTaxProfilesFieldParameter(param string, args ...interface{}) TaxProfilesFieldParameter {
	return TaxProfilesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateTaxProfilesQuery(mapTaxProfiless map[model.TaxProfilesPrimaryID]TaxProfilesUpdateFieldList, asTableValues string) (primaryIDs []model.TaxProfilesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.TaxProfilesPrimaryID]map[string]interface{}{}
	taxProfilesSelectFields := NewTaxProfilesSelectFields()
	for id, updateFields := range mapTaxProfiless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.taxProfilesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapTaxProfiless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetTaxProfilesFieldType(updateField.taxProfilesField)))
			args = append(args, fields[string(updateField.taxProfilesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.taxProfilesField))
		if updateField.taxProfilesField == taxProfilesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.taxProfilesField, asTableValues, updateField.taxProfilesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.taxProfilesField,
				"\"tax_profiles\"", updateField.taxProfilesField,
				asTableValues, updateField.taxProfilesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeTaxProfilesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.TaxProfilesPrimaryID, asTableValue string) (whereQry string) {
	taxProfilesSelectFields := NewTaxProfilesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"tax_profiles\".\"id\" = %s.\"id\"::"+GetTaxProfilesFieldType(taxProfilesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetTaxProfilesFieldType(taxProfilesField TaxProfilesField) string {
	selectTaxProfilesFields := NewTaxProfilesSelectFields()
	switch taxProfilesField {

	case selectTaxProfilesFields.Id():
		return "uuid"

	case selectTaxProfilesFields.OwnerPartyId():
		return "uuid"

	case selectTaxProfilesFields.CountryCode():
		return "text"

	case selectTaxProfilesFields.TaxResidencyCountry():
		return "text"

	case selectTaxProfilesFields.TaxIdMasked():
		return "text"

	case selectTaxProfilesFields.TaxEntityType():
		return "text"

	case selectTaxProfilesFields.IsVatRegistered():
		return "bool"

	case selectTaxProfilesFields.IsWithholdingSubject():
		return "bool"

	case selectTaxProfilesFields.ProfileStatus():
		return "profile_status_enum"

	case selectTaxProfilesFields.Metadata():
		return "jsonb"

	case selectTaxProfilesFields.MetaCreatedAt():
		return "timestamptz"

	case selectTaxProfilesFields.MetaCreatedBy():
		return "uuid"

	case selectTaxProfilesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectTaxProfilesFields.MetaUpdatedBy():
		return "uuid"

	case selectTaxProfilesFields.MetaDeletedAt():
		return "timestamptz"

	case selectTaxProfilesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateTaxProfiles(ctx context.Context, taxProfiles *model.TaxProfiles, fieldsInsert ...TaxProfilesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewTaxProfilesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.TaxProfilesPrimaryID{
		Id: taxProfiles.Id,
	}
	exists, err := repo.IsExistTaxProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxProfiles] failed checking taxProfiles whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "taxProfiles", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsTaxProfiles([]model.TaxProfiles{*taxProfiles}, fieldsInsert...)
	commandQuery := fmt.Sprintf(taxProfilesQueries.insertTaxProfiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxProfiles] failed exec create taxProfiles query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteTaxProfilesByID(ctx context.Context, primaryID model.TaxProfilesPrimaryID) (err error) {
	exists, err := repo.IsExistTaxProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxProfilesByID] failed checking taxProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxProfiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeTaxProfilesCompositePrimaryKeyWhere([]model.TaxProfilesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(taxProfilesQueries.deleteTaxProfiles + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxProfilesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxProfilesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxProfilesFilterResult, err error) {
	query, args, err := composeTaxProfilesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxProfilesByFilter] failed compose taxProfiles filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxProfilesByFilter] failed get taxProfiles by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeTaxProfilesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.TaxProfilesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeTaxProfilesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeTaxProfilesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeTaxProfilesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 16 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewTaxProfilesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 16+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["owner_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"owner_party_id\"")
			selectedColumns["owner_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["country_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"country_code\"")
			selectedColumns["country_code"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_residency_country"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_residency_country\"")
			selectedColumns["tax_residency_country"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_id_masked"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_id_masked\"")
			selectedColumns["tax_id_masked"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_entity_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_entity_type\"")
			selectedColumns["tax_entity_type"] = struct{}{}
		}
		if _, selected := selectedColumns["is_vat_registered"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_vat_registered\"")
			selectedColumns["is_vat_registered"] = struct{}{}
		}
		if _, selected := selectedColumns["is_withholding_subject"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_withholding_subject\"")
			selectedColumns["is_withholding_subject"] = struct{}{}
		}
		if _, selected := selectedColumns["profile_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"profile_status\"")
			selectedColumns["profile_status"] = struct{}{}
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

type taxProfilesFilterPlaceholder struct {
	index int
}

func (p *taxProfilesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeTaxProfilesFilterPredicate(filterField model.FilterField, placeholders *taxProfilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewTaxProfilesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeTaxProfilesFilterSQLExpr(spec)
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

func composeTaxProfilesFilterGroup(group model.FilterGroup, placeholders *taxProfilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeTaxProfilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeTaxProfilesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeTaxProfilesFilterWhereQueries(filter model.Filter, placeholders *taxProfilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeTaxProfilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeTaxProfilesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeTaxProfilesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateTaxProfilesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeTaxProfilesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeTaxProfilesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := taxProfilesFilterPlaceholder{index: 1}
	whereQueries, err := composeTaxProfilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewTaxProfilesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeTaxProfilesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeTaxProfilesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"tax_profiles\" base%s", strings.Join(selectColumns, ","), composeTaxProfilesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistTaxProfilesByID(ctx context.Context, primaryID model.TaxProfilesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeTaxProfilesCompositePrimaryKeyWhere([]model.TaxProfilesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", taxProfilesQueries.selectCountTaxProfiles, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxProfilesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxProfiles(ctx context.Context, selectFields ...TaxProfilesField) (taxProfilesList model.TaxProfilesList, err error) {
	var (
		defaultTaxProfilesSelectFields = defaultTaxProfilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxProfilesSelectFields = composeTaxProfilesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(taxProfilesQueries.selectTaxProfiles, defaultTaxProfilesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &taxProfilesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxProfiles] failed get taxProfiles list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxProfilesByID(ctx context.Context, primaryID model.TaxProfilesPrimaryID, selectFields ...TaxProfilesField) (taxProfiles model.TaxProfiles, err error) {
	var (
		defaultTaxProfilesSelectFields = defaultTaxProfilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxProfilesSelectFields = composeTaxProfilesSelectFields(selectFields...)
	}
	whereQry, params := composeTaxProfilesCompositePrimaryKeyWhere([]model.TaxProfilesPrimaryID{primaryID})
	query := fmt.Sprintf(taxProfilesQueries.selectTaxProfiles+" WHERE "+whereQry, defaultTaxProfilesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &taxProfiles, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("taxProfiles with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveTaxProfilesByID] failed get taxProfiles")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateTaxProfilesByID(ctx context.Context, primaryID model.TaxProfilesPrimaryID, taxProfiles *model.TaxProfiles, taxProfilesUpdateFields ...TaxProfilesUpdateField) (err error) {
	exists, err := repo.IsExistTaxProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxProfiles] failed checking taxProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxProfiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if taxProfiles == nil {
		if len(taxProfilesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateTaxProfilesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		taxProfiles = &model.TaxProfiles{}
	}
	var (
		defaultTaxProfilesUpdateFields = defaultTaxProfilesUpdateFields(*taxProfiles)
		tempUpdateField                TaxProfilesUpdateFieldList
		selectFields                   = NewTaxProfilesSelectFields()
	)
	if len(taxProfilesUpdateFields) > 0 {
		for _, updateField := range taxProfilesUpdateFields {
			if updateField.taxProfilesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultTaxProfilesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeTaxProfilesCompositePrimaryKeyWhere([]model.TaxProfilesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsTaxProfilesCommand(defaultTaxProfilesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(taxProfilesQueries.updateTaxProfiles+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxProfiles] error when try to update taxProfiles by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateTaxProfilesByFilter(ctx context.Context, filter model.Filter, taxProfilesUpdateFields ...TaxProfilesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(taxProfilesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields TaxProfilesUpdateFieldList
		selectFields = NewTaxProfilesSelectFields()
	)
	for _, updateField := range taxProfilesUpdateFields {
		if updateField.taxProfilesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsTaxProfilesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := taxProfilesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeTaxProfilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"tax_profiles\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxProfilesByFilter] error when try to update taxProfiles by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxProfilesByFilter] failed get rows affected")
	}
	return
}

var (
	taxProfilesQueries = struct {
		selectTaxProfiles      string
		selectCountTaxProfiles string
		deleteTaxProfiles      string
		updateTaxProfiles      string
		insertTaxProfiles      string
	}{
		selectTaxProfiles:      "SELECT %s FROM \"tax_profiles\"",
		selectCountTaxProfiles: "SELECT COUNT(\"id\") FROM \"tax_profiles\"",
		deleteTaxProfiles:      "DELETE FROM \"tax_profiles\"",
		updateTaxProfiles:      "UPDATE \"tax_profiles\" SET %s ",
		insertTaxProfiles:      "INSERT INTO \"tax_profiles\" %s VALUES %s",
	}
)

type TaxProfilesRepository interface {
	CreateTaxProfiles(ctx context.Context, taxProfiles *model.TaxProfiles, fieldsInsert ...TaxProfilesField) error
	BulkCreateTaxProfiles(ctx context.Context, taxProfilesList []*model.TaxProfiles, fieldsInsert ...TaxProfilesField) error
	ResolveTaxProfiles(ctx context.Context, selectFields ...TaxProfilesField) (model.TaxProfilesList, error)
	ResolveTaxProfilesByID(ctx context.Context, primaryID model.TaxProfilesPrimaryID, selectFields ...TaxProfilesField) (model.TaxProfiles, error)
	UpdateTaxProfilesByID(ctx context.Context, id model.TaxProfilesPrimaryID, taxProfiles *model.TaxProfiles, taxProfilesUpdateFields ...TaxProfilesUpdateField) error
	UpdateTaxProfilesByFilter(ctx context.Context, filter model.Filter, taxProfilesUpdateFields ...TaxProfilesUpdateField) (rowsAffected int64, err error)
	BulkUpdateTaxProfiles(ctx context.Context, taxProfilesListMap map[model.TaxProfilesPrimaryID]*model.TaxProfiles, TaxProfilessMapUpdateFieldsRequest map[model.TaxProfilesPrimaryID]TaxProfilesUpdateFieldList) (err error)
	DeleteTaxProfilesByID(ctx context.Context, id model.TaxProfilesPrimaryID) error
	BulkDeleteTaxProfilesByIDs(ctx context.Context, ids []model.TaxProfilesPrimaryID) error
	ResolveTaxProfilesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxProfilesFilterResult, err error)
	IsExistTaxProfilesByIDs(ctx context.Context, ids []model.TaxProfilesPrimaryID) (exists bool, notFoundIds []model.TaxProfilesPrimaryID, err error)
	IsExistTaxProfilesByID(ctx context.Context, id model.TaxProfilesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
