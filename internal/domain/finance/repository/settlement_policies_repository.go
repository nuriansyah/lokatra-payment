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

func composeInsertFieldsAndParamsSettlementPolicies(settlementPoliciesList []model.SettlementPolicies, fieldsInsert ...SettlementPoliciesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewSettlementPoliciesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, settlementPolicies := range settlementPoliciesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, settlementPolicies.Id)
			case selectField.PolicyCode():
				args = append(args, settlementPolicies.PolicyCode)
			case selectField.MerchantPartyId():
				args = append(args, settlementPolicies.MerchantPartyId)
			case selectField.PolicyScope():
				args = append(args, settlementPolicies.PolicyScope)
			case selectField.PolicyStatus():
				args = append(args, settlementPolicies.PolicyStatus)
			case selectField.Description():
				args = append(args, settlementPolicies.Description)
			case selectField.Metadata():
				args = append(args, settlementPolicies.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, settlementPolicies.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, settlementPolicies.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, settlementPolicies.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, settlementPolicies.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, settlementPolicies.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, settlementPolicies.MetaDeletedBy)

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

func composeSettlementPoliciesCompositePrimaryKeyWhere(primaryIDs []model.SettlementPoliciesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"settlement_policies\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultSettlementPoliciesSelectFields() string {
	fields := NewSettlementPoliciesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeSettlementPoliciesSelectFields(selectFields ...SettlementPoliciesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type SettlementPoliciesField string
type SettlementPoliciesFieldList []SettlementPoliciesField

type SettlementPoliciesSelectFields struct {
}

func (ss SettlementPoliciesSelectFields) Id() SettlementPoliciesField {
	return SettlementPoliciesField("id")
}

func (ss SettlementPoliciesSelectFields) PolicyCode() SettlementPoliciesField {
	return SettlementPoliciesField("policy_code")
}

func (ss SettlementPoliciesSelectFields) MerchantPartyId() SettlementPoliciesField {
	return SettlementPoliciesField("merchant_party_id")
}

func (ss SettlementPoliciesSelectFields) PolicyScope() SettlementPoliciesField {
	return SettlementPoliciesField("policy_scope")
}

func (ss SettlementPoliciesSelectFields) PolicyStatus() SettlementPoliciesField {
	return SettlementPoliciesField("policy_status")
}

func (ss SettlementPoliciesSelectFields) Description() SettlementPoliciesField {
	return SettlementPoliciesField("description")
}

func (ss SettlementPoliciesSelectFields) Metadata() SettlementPoliciesField {
	return SettlementPoliciesField("metadata")
}

func (ss SettlementPoliciesSelectFields) MetaCreatedAt() SettlementPoliciesField {
	return SettlementPoliciesField("meta_created_at")
}

func (ss SettlementPoliciesSelectFields) MetaCreatedBy() SettlementPoliciesField {
	return SettlementPoliciesField("meta_created_by")
}

func (ss SettlementPoliciesSelectFields) MetaUpdatedAt() SettlementPoliciesField {
	return SettlementPoliciesField("meta_updated_at")
}

func (ss SettlementPoliciesSelectFields) MetaUpdatedBy() SettlementPoliciesField {
	return SettlementPoliciesField("meta_updated_by")
}

func (ss SettlementPoliciesSelectFields) MetaDeletedAt() SettlementPoliciesField {
	return SettlementPoliciesField("meta_deleted_at")
}

func (ss SettlementPoliciesSelectFields) MetaDeletedBy() SettlementPoliciesField {
	return SettlementPoliciesField("meta_deleted_by")
}

func (ss SettlementPoliciesSelectFields) All() SettlementPoliciesFieldList {
	return []SettlementPoliciesField{
		ss.Id(),
		ss.PolicyCode(),
		ss.MerchantPartyId(),
		ss.PolicyScope(),
		ss.PolicyStatus(),
		ss.Description(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewSettlementPoliciesSelectFields() SettlementPoliciesSelectFields {
	return SettlementPoliciesSelectFields{}
}

type SettlementPoliciesUpdateFieldOption struct {
	useIncrement bool
}
type SettlementPoliciesUpdateField struct {
	settlementPoliciesField SettlementPoliciesField
	opt                     SettlementPoliciesUpdateFieldOption
	value                   interface{}
}
type SettlementPoliciesUpdateFieldList []SettlementPoliciesUpdateField

func defaultSettlementPoliciesUpdateFieldOption() SettlementPoliciesUpdateFieldOption {
	return SettlementPoliciesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementSettlementPoliciesOption(useIncrement bool) func(*SettlementPoliciesUpdateFieldOption) {
	return func(pcufo *SettlementPoliciesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewSettlementPoliciesUpdateField(field SettlementPoliciesField, val interface{}, opts ...func(*SettlementPoliciesUpdateFieldOption)) SettlementPoliciesUpdateField {
	defaultOpt := defaultSettlementPoliciesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return SettlementPoliciesUpdateField{
		settlementPoliciesField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultSettlementPoliciesUpdateFields(settlementPolicies model.SettlementPolicies) (settlementPoliciesUpdateFieldList SettlementPoliciesUpdateFieldList) {
	selectFields := NewSettlementPoliciesSelectFields()
	settlementPoliciesUpdateFieldList = append(settlementPoliciesUpdateFieldList,
		NewSettlementPoliciesUpdateField(selectFields.Id(), settlementPolicies.Id),
		NewSettlementPoliciesUpdateField(selectFields.PolicyCode(), settlementPolicies.PolicyCode),
		NewSettlementPoliciesUpdateField(selectFields.MerchantPartyId(), settlementPolicies.MerchantPartyId),
		NewSettlementPoliciesUpdateField(selectFields.PolicyScope(), settlementPolicies.PolicyScope),
		NewSettlementPoliciesUpdateField(selectFields.PolicyStatus(), settlementPolicies.PolicyStatus),
		NewSettlementPoliciesUpdateField(selectFields.Description(), settlementPolicies.Description),
		NewSettlementPoliciesUpdateField(selectFields.Metadata(), settlementPolicies.Metadata),
		NewSettlementPoliciesUpdateField(selectFields.MetaCreatedAt(), settlementPolicies.MetaCreatedAt),
		NewSettlementPoliciesUpdateField(selectFields.MetaCreatedBy(), settlementPolicies.MetaCreatedBy),
		NewSettlementPoliciesUpdateField(selectFields.MetaUpdatedAt(), settlementPolicies.MetaUpdatedAt),
		NewSettlementPoliciesUpdateField(selectFields.MetaUpdatedBy(), settlementPolicies.MetaUpdatedBy),
		NewSettlementPoliciesUpdateField(selectFields.MetaDeletedAt(), settlementPolicies.MetaDeletedAt),
		NewSettlementPoliciesUpdateField(selectFields.MetaDeletedBy(), settlementPolicies.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsSettlementPoliciesCommand(settlementPoliciesUpdateFieldList SettlementPoliciesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range settlementPoliciesUpdateFieldList {
		field := string(updateField.settlementPoliciesField)
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

func (repo *RepositoryImpl) BulkCreateSettlementPolicies(ctx context.Context, settlementPoliciesList []*model.SettlementPolicies, fieldsInsert ...SettlementPoliciesField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.SettlementPoliciesPrimaryID
		settlementPoliciesValueList []model.SettlementPolicies
	)

	if len(fieldsInsert) == 0 {
		selectField := NewSettlementPoliciesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, settlementPolicies := range settlementPoliciesList {

		primaryIds = append(primaryIds, settlementPolicies.ToSettlementPoliciesPrimaryID())

		settlementPoliciesValueList = append(settlementPoliciesValueList, *settlementPolicies)
	}

	_, notFoundIds, err := repo.IsExistSettlementPoliciesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementPolicies] failed checking settlementPolicies whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.SettlementPoliciesPrimaryID{}
		mapNotFoundIds := map[model.SettlementPoliciesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "settlementPolicies", fmt.Sprintf("settlementPolicies with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsSettlementPolicies(settlementPoliciesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(settlementPoliciesQueries.insertSettlementPolicies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementPolicies] failed exec create settlementPolicies query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteSettlementPoliciesByIDs(ctx context.Context, primaryIDs []model.SettlementPoliciesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistSettlementPoliciesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementPoliciesByIDs] failed checking settlementPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicies with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_policies\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(settlementPoliciesQueries.deleteSettlementPolicies + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementPoliciesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementPoliciesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistSettlementPoliciesByIDs(ctx context.Context, ids []model.SettlementPoliciesPrimaryID) (exists bool, notFoundIds []model.SettlementPoliciesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_policies\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(settlementPoliciesQueries.selectSettlementPolicies, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementPoliciesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.SettlementPoliciesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementPoliciesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.SettlementPoliciesPrimaryID]bool{}
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

// BulkUpdateSettlementPolicies is used to bulk update settlementPolicies, by default it will update all field
// if want to update specific field, then fill settlementPoliciessMapUpdateFieldsRequest else please fill settlementPoliciessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateSettlementPolicies(ctx context.Context, settlementPoliciessMap map[model.SettlementPoliciesPrimaryID]*model.SettlementPolicies, settlementPoliciessMapUpdateFieldsRequest map[model.SettlementPoliciesPrimaryID]SettlementPoliciesUpdateFieldList) (err error) {
	if len(settlementPoliciessMap) == 0 && len(settlementPoliciessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		settlementPoliciessMapUpdateField map[model.SettlementPoliciesPrimaryID]SettlementPoliciesUpdateFieldList = map[model.SettlementPoliciesPrimaryID]SettlementPoliciesUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(settlementPoliciessMap) > 0 {
		for id, settlementPolicies := range settlementPoliciessMap {
			if settlementPolicies == nil {
				log.Error().Err(err).Msg("[BulkUpdateSettlementPolicies] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			settlementPoliciessMapUpdateField[id] = defaultSettlementPoliciesUpdateFields(*settlementPolicies)
		}
	} else {
		settlementPoliciessMapUpdateField = settlementPoliciessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateSettlementPoliciesQuery(settlementPoliciessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistSettlementPoliciesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementPolicies] failed checking settlementPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicies with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeSettlementPoliciesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"settlement_policies\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementPolicies] failed exec query")
	}
	return
}

type SettlementPoliciesFieldParameter struct {
	param string
	args  []interface{}
}

func NewSettlementPoliciesFieldParameter(param string, args ...interface{}) SettlementPoliciesFieldParameter {
	return SettlementPoliciesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateSettlementPoliciesQuery(mapSettlementPoliciess map[model.SettlementPoliciesPrimaryID]SettlementPoliciesUpdateFieldList, asTableValues string) (primaryIDs []model.SettlementPoliciesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.SettlementPoliciesPrimaryID]map[string]interface{}{}
	settlementPoliciesSelectFields := NewSettlementPoliciesSelectFields()
	for id, updateFields := range mapSettlementPoliciess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.settlementPoliciesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapSettlementPoliciess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetSettlementPoliciesFieldType(updateField.settlementPoliciesField)))
			args = append(args, fields[string(updateField.settlementPoliciesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.settlementPoliciesField))
		if updateField.settlementPoliciesField == settlementPoliciesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.settlementPoliciesField, asTableValues, updateField.settlementPoliciesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.settlementPoliciesField,
				"\"settlement_policies\"", updateField.settlementPoliciesField,
				asTableValues, updateField.settlementPoliciesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeSettlementPoliciesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.SettlementPoliciesPrimaryID, asTableValue string) (whereQry string) {
	settlementPoliciesSelectFields := NewSettlementPoliciesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"settlement_policies\".\"id\" = %s.\"id\"::"+GetSettlementPoliciesFieldType(settlementPoliciesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetSettlementPoliciesFieldType(settlementPoliciesField SettlementPoliciesField) string {
	selectSettlementPoliciesFields := NewSettlementPoliciesSelectFields()
	switch settlementPoliciesField {

	case selectSettlementPoliciesFields.Id():
		return "uuid"

	case selectSettlementPoliciesFields.PolicyCode():
		return "text"

	case selectSettlementPoliciesFields.MerchantPartyId():
		return "uuid"

	case selectSettlementPoliciesFields.PolicyScope():
		return "settlement_policies_policy_scope_enum"

	case selectSettlementPoliciesFields.PolicyStatus():
		return "settlement_policies_policy_status_enum"

	case selectSettlementPoliciesFields.Description():
		return "text"

	case selectSettlementPoliciesFields.Metadata():
		return "jsonb"

	case selectSettlementPoliciesFields.MetaCreatedAt():
		return "timestamptz"

	case selectSettlementPoliciesFields.MetaCreatedBy():
		return "uuid"

	case selectSettlementPoliciesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectSettlementPoliciesFields.MetaUpdatedBy():
		return "uuid"

	case selectSettlementPoliciesFields.MetaDeletedAt():
		return "timestamptz"

	case selectSettlementPoliciesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateSettlementPolicies(ctx context.Context, settlementPolicies *model.SettlementPolicies, fieldsInsert ...SettlementPoliciesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewSettlementPoliciesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.SettlementPoliciesPrimaryID{
		Id: settlementPolicies.Id,
	}
	exists, err := repo.IsExistSettlementPoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementPolicies] failed checking settlementPolicies whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "settlementPolicies", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsSettlementPolicies([]model.SettlementPolicies{*settlementPolicies}, fieldsInsert...)
	commandQuery := fmt.Sprintf(settlementPoliciesQueries.insertSettlementPolicies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementPolicies] failed exec create settlementPolicies query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteSettlementPoliciesByID(ctx context.Context, primaryID model.SettlementPoliciesPrimaryID) (err error) {
	exists, err := repo.IsExistSettlementPoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementPoliciesByID] failed checking settlementPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeSettlementPoliciesCompositePrimaryKeyWhere([]model.SettlementPoliciesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(settlementPoliciesQueries.deleteSettlementPolicies + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementPoliciesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementPoliciesByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementPoliciesFilterResult, err error) {
	query, args, err := composeSettlementPoliciesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementPoliciesByFilter] failed compose settlementPolicies filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementPoliciesByFilter] failed get settlementPolicies by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeSettlementPoliciesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.SettlementPoliciesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeSettlementPoliciesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeSettlementPoliciesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeSettlementPoliciesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 13 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewSettlementPoliciesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 13+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["policy_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"policy_code\"")
			selectedColumns["policy_code"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["policy_scope"]; !selected {
			selectColumns = append(selectColumns, "base.\"policy_scope\"")
			selectedColumns["policy_scope"] = struct{}{}
		}
		if _, selected := selectedColumns["policy_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"policy_status\"")
			selectedColumns["policy_status"] = struct{}{}
		}
		if _, selected := selectedColumns["description"]; !selected {
			selectColumns = append(selectColumns, "base.\"description\"")
			selectedColumns["description"] = struct{}{}
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

type settlementPoliciesFilterPlaceholder struct {
	index int
}

func (p *settlementPoliciesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeSettlementPoliciesFilterPredicate(filterField model.FilterField, placeholders *settlementPoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewSettlementPoliciesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeSettlementPoliciesFilterSQLExpr(spec)
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

func composeSettlementPoliciesFilterGroup(group model.FilterGroup, placeholders *settlementPoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeSettlementPoliciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeSettlementPoliciesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeSettlementPoliciesFilterWhereQueries(filter model.Filter, placeholders *settlementPoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeSettlementPoliciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeSettlementPoliciesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeSettlementPoliciesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateSettlementPoliciesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeSettlementPoliciesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeSettlementPoliciesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := settlementPoliciesFilterPlaceholder{index: 1}
	whereQueries, err := composeSettlementPoliciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewSettlementPoliciesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeSettlementPoliciesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeSettlementPoliciesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"settlement_policies\" base%s", strings.Join(selectColumns, ","), composeSettlementPoliciesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistSettlementPoliciesByID(ctx context.Context, primaryID model.SettlementPoliciesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeSettlementPoliciesCompositePrimaryKeyWhere([]model.SettlementPoliciesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", settlementPoliciesQueries.selectCountSettlementPolicies, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementPoliciesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementPolicies(ctx context.Context, selectFields ...SettlementPoliciesField) (settlementPoliciesList model.SettlementPoliciesList, err error) {
	var (
		defaultSettlementPoliciesSelectFields = defaultSettlementPoliciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementPoliciesSelectFields = composeSettlementPoliciesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(settlementPoliciesQueries.selectSettlementPolicies, defaultSettlementPoliciesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &settlementPoliciesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementPolicies] failed get settlementPolicies list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementPoliciesByID(ctx context.Context, primaryID model.SettlementPoliciesPrimaryID, selectFields ...SettlementPoliciesField) (settlementPolicies model.SettlementPolicies, err error) {
	var (
		defaultSettlementPoliciesSelectFields = defaultSettlementPoliciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementPoliciesSelectFields = composeSettlementPoliciesSelectFields(selectFields...)
	}
	whereQry, params := composeSettlementPoliciesCompositePrimaryKeyWhere([]model.SettlementPoliciesPrimaryID{primaryID})
	query := fmt.Sprintf(settlementPoliciesQueries.selectSettlementPolicies+" WHERE "+whereQry, defaultSettlementPoliciesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &settlementPolicies, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("settlementPolicies with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveSettlementPoliciesByID] failed get settlementPolicies")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateSettlementPoliciesByID(ctx context.Context, primaryID model.SettlementPoliciesPrimaryID, settlementPolicies *model.SettlementPolicies, settlementPoliciesUpdateFields ...SettlementPoliciesUpdateField) (err error) {
	exists, err := repo.IsExistSettlementPoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPolicies] failed checking settlementPolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if settlementPolicies == nil {
		if len(settlementPoliciesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateSettlementPoliciesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		settlementPolicies = &model.SettlementPolicies{}
	}
	var (
		defaultSettlementPoliciesUpdateFields = defaultSettlementPoliciesUpdateFields(*settlementPolicies)
		tempUpdateField                       SettlementPoliciesUpdateFieldList
		selectFields                          = NewSettlementPoliciesSelectFields()
	)
	if len(settlementPoliciesUpdateFields) > 0 {
		for _, updateField := range settlementPoliciesUpdateFields {
			if updateField.settlementPoliciesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultSettlementPoliciesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeSettlementPoliciesCompositePrimaryKeyWhere([]model.SettlementPoliciesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsSettlementPoliciesCommand(defaultSettlementPoliciesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(settlementPoliciesQueries.updateSettlementPolicies+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPolicies] error when try to update settlementPolicies by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateSettlementPoliciesByFilter(ctx context.Context, filter model.Filter, settlementPoliciesUpdateFields ...SettlementPoliciesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(settlementPoliciesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields SettlementPoliciesUpdateFieldList
		selectFields = NewSettlementPoliciesSelectFields()
	)
	for _, updateField := range settlementPoliciesUpdateFields {
		if updateField.settlementPoliciesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsSettlementPoliciesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := settlementPoliciesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeSettlementPoliciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"settlement_policies\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPoliciesByFilter] error when try to update settlementPolicies by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPoliciesByFilter] failed get rows affected")
	}
	return
}

var (
	settlementPoliciesQueries = struct {
		selectSettlementPolicies      string
		selectCountSettlementPolicies string
		deleteSettlementPolicies      string
		updateSettlementPolicies      string
		insertSettlementPolicies      string
	}{
		selectSettlementPolicies:      "SELECT %s FROM \"settlement_policies\"",
		selectCountSettlementPolicies: "SELECT COUNT(\"id\") FROM \"settlement_policies\"",
		deleteSettlementPolicies:      "DELETE FROM \"settlement_policies\"",
		updateSettlementPolicies:      "UPDATE \"settlement_policies\" SET %s ",
		insertSettlementPolicies:      "INSERT INTO \"settlement_policies\" %s VALUES %s",
	}
)

type SettlementPoliciesRepository interface {
	CreateSettlementPolicies(ctx context.Context, settlementPolicies *model.SettlementPolicies, fieldsInsert ...SettlementPoliciesField) error
	BulkCreateSettlementPolicies(ctx context.Context, settlementPoliciesList []*model.SettlementPolicies, fieldsInsert ...SettlementPoliciesField) error
	ResolveSettlementPolicies(ctx context.Context, selectFields ...SettlementPoliciesField) (model.SettlementPoliciesList, error)
	ResolveSettlementPoliciesByID(ctx context.Context, primaryID model.SettlementPoliciesPrimaryID, selectFields ...SettlementPoliciesField) (model.SettlementPolicies, error)
	UpdateSettlementPoliciesByID(ctx context.Context, id model.SettlementPoliciesPrimaryID, settlementPolicies *model.SettlementPolicies, settlementPoliciesUpdateFields ...SettlementPoliciesUpdateField) error
	UpdateSettlementPoliciesByFilter(ctx context.Context, filter model.Filter, settlementPoliciesUpdateFields ...SettlementPoliciesUpdateField) (rowsAffected int64, err error)
	BulkUpdateSettlementPolicies(ctx context.Context, settlementPoliciesListMap map[model.SettlementPoliciesPrimaryID]*model.SettlementPolicies, SettlementPoliciessMapUpdateFieldsRequest map[model.SettlementPoliciesPrimaryID]SettlementPoliciesUpdateFieldList) (err error)
	DeleteSettlementPoliciesByID(ctx context.Context, id model.SettlementPoliciesPrimaryID) error
	BulkDeleteSettlementPoliciesByIDs(ctx context.Context, ids []model.SettlementPoliciesPrimaryID) error
	ResolveSettlementPoliciesByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementPoliciesFilterResult, err error)
	IsExistSettlementPoliciesByIDs(ctx context.Context, ids []model.SettlementPoliciesPrimaryID) (exists bool, notFoundIds []model.SettlementPoliciesPrimaryID, err error)
	IsExistSettlementPoliciesByID(ctx context.Context, id model.SettlementPoliciesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
