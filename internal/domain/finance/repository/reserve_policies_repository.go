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

func composeInsertFieldsAndParamsReservePolicies(reservePoliciesList []model.ReservePolicies, fieldsInsert ...ReservePoliciesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewReservePoliciesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, reservePolicies := range reservePoliciesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, reservePolicies.Id)
			case selectField.PolicyCode():
				args = append(args, reservePolicies.PolicyCode)
			case selectField.MerchantPartyId():
				args = append(args, reservePolicies.MerchantPartyId)
			case selectField.PolicyScope():
				args = append(args, reservePolicies.PolicyScope)
			case selectField.ReserveStatus():
				args = append(args, reservePolicies.ReserveStatus)
			case selectField.Description():
				args = append(args, reservePolicies.Description)
			case selectField.Metadata():
				args = append(args, reservePolicies.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, reservePolicies.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, reservePolicies.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, reservePolicies.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, reservePolicies.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, reservePolicies.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, reservePolicies.MetaDeletedBy)

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

func composeReservePoliciesCompositePrimaryKeyWhere(primaryIDs []model.ReservePoliciesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"reserve_policies\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultReservePoliciesSelectFields() string {
	fields := NewReservePoliciesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeReservePoliciesSelectFields(selectFields ...ReservePoliciesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ReservePoliciesField string
type ReservePoliciesFieldList []ReservePoliciesField

type ReservePoliciesSelectFields struct {
}

func (ss ReservePoliciesSelectFields) Id() ReservePoliciesField {
	return ReservePoliciesField("id")
}

func (ss ReservePoliciesSelectFields) PolicyCode() ReservePoliciesField {
	return ReservePoliciesField("policy_code")
}

func (ss ReservePoliciesSelectFields) MerchantPartyId() ReservePoliciesField {
	return ReservePoliciesField("merchant_party_id")
}

func (ss ReservePoliciesSelectFields) PolicyScope() ReservePoliciesField {
	return ReservePoliciesField("policy_scope")
}

func (ss ReservePoliciesSelectFields) ReserveStatus() ReservePoliciesField {
	return ReservePoliciesField("reserve_status")
}

func (ss ReservePoliciesSelectFields) Description() ReservePoliciesField {
	return ReservePoliciesField("description")
}

func (ss ReservePoliciesSelectFields) Metadata() ReservePoliciesField {
	return ReservePoliciesField("metadata")
}

func (ss ReservePoliciesSelectFields) MetaCreatedAt() ReservePoliciesField {
	return ReservePoliciesField("meta_created_at")
}

func (ss ReservePoliciesSelectFields) MetaCreatedBy() ReservePoliciesField {
	return ReservePoliciesField("meta_created_by")
}

func (ss ReservePoliciesSelectFields) MetaUpdatedAt() ReservePoliciesField {
	return ReservePoliciesField("meta_updated_at")
}

func (ss ReservePoliciesSelectFields) MetaUpdatedBy() ReservePoliciesField {
	return ReservePoliciesField("meta_updated_by")
}

func (ss ReservePoliciesSelectFields) MetaDeletedAt() ReservePoliciesField {
	return ReservePoliciesField("meta_deleted_at")
}

func (ss ReservePoliciesSelectFields) MetaDeletedBy() ReservePoliciesField {
	return ReservePoliciesField("meta_deleted_by")
}

func (ss ReservePoliciesSelectFields) All() ReservePoliciesFieldList {
	return []ReservePoliciesField{
		ss.Id(),
		ss.PolicyCode(),
		ss.MerchantPartyId(),
		ss.PolicyScope(),
		ss.ReserveStatus(),
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

func NewReservePoliciesSelectFields() ReservePoliciesSelectFields {
	return ReservePoliciesSelectFields{}
}

type ReservePoliciesUpdateFieldOption struct {
	useIncrement bool
}
type ReservePoliciesUpdateField struct {
	reservePoliciesField ReservePoliciesField
	opt                  ReservePoliciesUpdateFieldOption
	value                interface{}
}
type ReservePoliciesUpdateFieldList []ReservePoliciesUpdateField

func defaultReservePoliciesUpdateFieldOption() ReservePoliciesUpdateFieldOption {
	return ReservePoliciesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementReservePoliciesOption(useIncrement bool) func(*ReservePoliciesUpdateFieldOption) {
	return func(pcufo *ReservePoliciesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewReservePoliciesUpdateField(field ReservePoliciesField, val interface{}, opts ...func(*ReservePoliciesUpdateFieldOption)) ReservePoliciesUpdateField {
	defaultOpt := defaultReservePoliciesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ReservePoliciesUpdateField{
		reservePoliciesField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultReservePoliciesUpdateFields(reservePolicies model.ReservePolicies) (reservePoliciesUpdateFieldList ReservePoliciesUpdateFieldList) {
	selectFields := NewReservePoliciesSelectFields()
	reservePoliciesUpdateFieldList = append(reservePoliciesUpdateFieldList,
		NewReservePoliciesUpdateField(selectFields.Id(), reservePolicies.Id),
		NewReservePoliciesUpdateField(selectFields.PolicyCode(), reservePolicies.PolicyCode),
		NewReservePoliciesUpdateField(selectFields.MerchantPartyId(), reservePolicies.MerchantPartyId),
		NewReservePoliciesUpdateField(selectFields.PolicyScope(), reservePolicies.PolicyScope),
		NewReservePoliciesUpdateField(selectFields.ReserveStatus(), reservePolicies.ReserveStatus),
		NewReservePoliciesUpdateField(selectFields.Description(), reservePolicies.Description),
		NewReservePoliciesUpdateField(selectFields.Metadata(), reservePolicies.Metadata),
		NewReservePoliciesUpdateField(selectFields.MetaCreatedAt(), reservePolicies.MetaCreatedAt),
		NewReservePoliciesUpdateField(selectFields.MetaCreatedBy(), reservePolicies.MetaCreatedBy),
		NewReservePoliciesUpdateField(selectFields.MetaUpdatedAt(), reservePolicies.MetaUpdatedAt),
		NewReservePoliciesUpdateField(selectFields.MetaUpdatedBy(), reservePolicies.MetaUpdatedBy),
		NewReservePoliciesUpdateField(selectFields.MetaDeletedAt(), reservePolicies.MetaDeletedAt),
		NewReservePoliciesUpdateField(selectFields.MetaDeletedBy(), reservePolicies.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsReservePoliciesCommand(reservePoliciesUpdateFieldList ReservePoliciesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range reservePoliciesUpdateFieldList {
		field := string(updateField.reservePoliciesField)
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

func (repo *RepositoryImpl) BulkCreateReservePolicies(ctx context.Context, reservePoliciesList []*model.ReservePolicies, fieldsInsert ...ReservePoliciesField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.ReservePoliciesPrimaryID
		reservePoliciesValueList []model.ReservePolicies
	)

	if len(fieldsInsert) == 0 {
		selectField := NewReservePoliciesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, reservePolicies := range reservePoliciesList {

		primaryIds = append(primaryIds, reservePolicies.ToReservePoliciesPrimaryID())

		reservePoliciesValueList = append(reservePoliciesValueList, *reservePolicies)
	}

	_, notFoundIds, err := repo.IsExistReservePoliciesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReservePolicies] failed checking reservePolicies whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ReservePoliciesPrimaryID{}
		mapNotFoundIds := map[model.ReservePoliciesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "reservePolicies", fmt.Sprintf("reservePolicies with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsReservePolicies(reservePoliciesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(reservePoliciesQueries.insertReservePolicies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReservePolicies] failed exec create reservePolicies query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteReservePoliciesByIDs(ctx context.Context, primaryIDs []model.ReservePoliciesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistReservePoliciesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReservePoliciesByIDs] failed checking reservePolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reservePolicies with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reserve_policies\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(reservePoliciesQueries.deleteReservePolicies + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReservePoliciesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReservePoliciesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistReservePoliciesByIDs(ctx context.Context, ids []model.ReservePoliciesPrimaryID) (exists bool, notFoundIds []model.ReservePoliciesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reserve_policies\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(reservePoliciesQueries.selectReservePolicies, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReservePoliciesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ReservePoliciesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReservePoliciesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ReservePoliciesPrimaryID]bool{}
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

// BulkUpdateReservePolicies is used to bulk update reservePolicies, by default it will update all field
// if want to update specific field, then fill reservePoliciessMapUpdateFieldsRequest else please fill reservePoliciessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateReservePolicies(ctx context.Context, reservePoliciessMap map[model.ReservePoliciesPrimaryID]*model.ReservePolicies, reservePoliciessMapUpdateFieldsRequest map[model.ReservePoliciesPrimaryID]ReservePoliciesUpdateFieldList) (err error) {
	if len(reservePoliciessMap) == 0 && len(reservePoliciessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		reservePoliciessMapUpdateField map[model.ReservePoliciesPrimaryID]ReservePoliciesUpdateFieldList = map[model.ReservePoliciesPrimaryID]ReservePoliciesUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(reservePoliciessMap) > 0 {
		for id, reservePolicies := range reservePoliciessMap {
			if reservePolicies == nil {
				log.Error().Err(err).Msg("[BulkUpdateReservePolicies] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			reservePoliciessMapUpdateField[id] = defaultReservePoliciesUpdateFields(*reservePolicies)
		}
	} else {
		reservePoliciessMapUpdateField = reservePoliciessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateReservePoliciesQuery(reservePoliciessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistReservePoliciesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReservePolicies] failed checking reservePolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reservePolicies with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeReservePoliciesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"reserve_policies\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReservePolicies] failed exec query")
	}
	return
}

type ReservePoliciesFieldParameter struct {
	param string
	args  []interface{}
}

func NewReservePoliciesFieldParameter(param string, args ...interface{}) ReservePoliciesFieldParameter {
	return ReservePoliciesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateReservePoliciesQuery(mapReservePoliciess map[model.ReservePoliciesPrimaryID]ReservePoliciesUpdateFieldList, asTableValues string) (primaryIDs []model.ReservePoliciesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ReservePoliciesPrimaryID]map[string]interface{}{}
	reservePoliciesSelectFields := NewReservePoliciesSelectFields()
	for id, updateFields := range mapReservePoliciess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.reservePoliciesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapReservePoliciess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetReservePoliciesFieldType(updateField.reservePoliciesField)))
			args = append(args, fields[string(updateField.reservePoliciesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.reservePoliciesField))
		if updateField.reservePoliciesField == reservePoliciesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.reservePoliciesField, asTableValues, updateField.reservePoliciesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.reservePoliciesField,
				"\"reserve_policies\"", updateField.reservePoliciesField,
				asTableValues, updateField.reservePoliciesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeReservePoliciesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ReservePoliciesPrimaryID, asTableValue string) (whereQry string) {
	reservePoliciesSelectFields := NewReservePoliciesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"reserve_policies\".\"id\" = %s.\"id\"::"+GetReservePoliciesFieldType(reservePoliciesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetReservePoliciesFieldType(reservePoliciesField ReservePoliciesField) string {
	selectReservePoliciesFields := NewReservePoliciesSelectFields()
	switch reservePoliciesField {

	case selectReservePoliciesFields.Id():
		return "uuid"

	case selectReservePoliciesFields.PolicyCode():
		return "text"

	case selectReservePoliciesFields.MerchantPartyId():
		return "uuid"

	case selectReservePoliciesFields.PolicyScope():
		return "reserve_policies_policy_scope_enum"

	case selectReservePoliciesFields.ReserveStatus():
		return "reserve_status_enum"

	case selectReservePoliciesFields.Description():
		return "text"

	case selectReservePoliciesFields.Metadata():
		return "jsonb"

	case selectReservePoliciesFields.MetaCreatedAt():
		return "timestamptz"

	case selectReservePoliciesFields.MetaCreatedBy():
		return "uuid"

	case selectReservePoliciesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectReservePoliciesFields.MetaUpdatedBy():
		return "uuid"

	case selectReservePoliciesFields.MetaDeletedAt():
		return "timestamptz"

	case selectReservePoliciesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateReservePolicies(ctx context.Context, reservePolicies *model.ReservePolicies, fieldsInsert ...ReservePoliciesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewReservePoliciesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ReservePoliciesPrimaryID{
		Id: reservePolicies.Id,
	}
	exists, err := repo.IsExistReservePoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReservePolicies] failed checking reservePolicies whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "reservePolicies", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsReservePolicies([]model.ReservePolicies{*reservePolicies}, fieldsInsert...)
	commandQuery := fmt.Sprintf(reservePoliciesQueries.insertReservePolicies, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReservePolicies] failed exec create reservePolicies query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteReservePoliciesByID(ctx context.Context, primaryID model.ReservePoliciesPrimaryID) (err error) {
	exists, err := repo.IsExistReservePoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReservePoliciesByID] failed checking reservePolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reservePolicies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeReservePoliciesCompositePrimaryKeyWhere([]model.ReservePoliciesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(reservePoliciesQueries.deleteReservePolicies + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReservePoliciesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveReservePoliciesByFilter(ctx context.Context, filter model.Filter) (result []model.ReservePoliciesFilterResult, err error) {
	query, args, err := composeReservePoliciesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReservePoliciesByFilter] failed compose reservePolicies filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReservePoliciesByFilter] failed get reservePolicies by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeReservePoliciesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ReservePoliciesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeReservePoliciesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeReservePoliciesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeReservePoliciesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 13 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewReservePoliciesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["reserve_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"reserve_status\"")
			selectedColumns["reserve_status"] = struct{}{}
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

type reservePoliciesFilterPlaceholder struct {
	index int
}

func (p *reservePoliciesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeReservePoliciesFilterPredicate(filterField model.FilterField, placeholders *reservePoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewReservePoliciesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeReservePoliciesFilterSQLExpr(spec)
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

func composeReservePoliciesFilterGroup(group model.FilterGroup, placeholders *reservePoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeReservePoliciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeReservePoliciesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeReservePoliciesFilterWhereQueries(filter model.Filter, placeholders *reservePoliciesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeReservePoliciesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeReservePoliciesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeReservePoliciesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateReservePoliciesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeReservePoliciesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeReservePoliciesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := reservePoliciesFilterPlaceholder{index: 1}
	whereQueries, err := composeReservePoliciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewReservePoliciesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeReservePoliciesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeReservePoliciesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"reserve_policies\" base%s", strings.Join(selectColumns, ","), composeReservePoliciesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistReservePoliciesByID(ctx context.Context, primaryID model.ReservePoliciesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeReservePoliciesCompositePrimaryKeyWhere([]model.ReservePoliciesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", reservePoliciesQueries.selectCountReservePolicies, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReservePoliciesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReservePolicies(ctx context.Context, selectFields ...ReservePoliciesField) (reservePoliciesList model.ReservePoliciesList, err error) {
	var (
		defaultReservePoliciesSelectFields = defaultReservePoliciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReservePoliciesSelectFields = composeReservePoliciesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(reservePoliciesQueries.selectReservePolicies, defaultReservePoliciesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &reservePoliciesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReservePolicies] failed get reservePolicies list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReservePoliciesByID(ctx context.Context, primaryID model.ReservePoliciesPrimaryID, selectFields ...ReservePoliciesField) (reservePolicies model.ReservePolicies, err error) {
	var (
		defaultReservePoliciesSelectFields = defaultReservePoliciesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReservePoliciesSelectFields = composeReservePoliciesSelectFields(selectFields...)
	}
	whereQry, params := composeReservePoliciesCompositePrimaryKeyWhere([]model.ReservePoliciesPrimaryID{primaryID})
	query := fmt.Sprintf(reservePoliciesQueries.selectReservePolicies+" WHERE "+whereQry, defaultReservePoliciesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &reservePolicies, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("reservePolicies with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveReservePoliciesByID] failed get reservePolicies")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateReservePoliciesByID(ctx context.Context, primaryID model.ReservePoliciesPrimaryID, reservePolicies *model.ReservePolicies, reservePoliciesUpdateFields ...ReservePoliciesUpdateField) (err error) {
	exists, err := repo.IsExistReservePoliciesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReservePolicies] failed checking reservePolicies whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reservePolicies with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if reservePolicies == nil {
		if len(reservePoliciesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateReservePoliciesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		reservePolicies = &model.ReservePolicies{}
	}
	var (
		defaultReservePoliciesUpdateFields = defaultReservePoliciesUpdateFields(*reservePolicies)
		tempUpdateField                    ReservePoliciesUpdateFieldList
		selectFields                       = NewReservePoliciesSelectFields()
	)
	if len(reservePoliciesUpdateFields) > 0 {
		for _, updateField := range reservePoliciesUpdateFields {
			if updateField.reservePoliciesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultReservePoliciesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeReservePoliciesCompositePrimaryKeyWhere([]model.ReservePoliciesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsReservePoliciesCommand(defaultReservePoliciesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(reservePoliciesQueries.updateReservePolicies+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReservePolicies] error when try to update reservePolicies by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateReservePoliciesByFilter(ctx context.Context, filter model.Filter, reservePoliciesUpdateFields ...ReservePoliciesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(reservePoliciesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ReservePoliciesUpdateFieldList
		selectFields = NewReservePoliciesSelectFields()
	)
	for _, updateField := range reservePoliciesUpdateFields {
		if updateField.reservePoliciesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsReservePoliciesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := reservePoliciesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeReservePoliciesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"reserve_policies\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReservePoliciesByFilter] error when try to update reservePolicies by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReservePoliciesByFilter] failed get rows affected")
	}
	return
}

var (
	reservePoliciesQueries = struct {
		selectReservePolicies      string
		selectCountReservePolicies string
		deleteReservePolicies      string
		updateReservePolicies      string
		insertReservePolicies      string
	}{
		selectReservePolicies:      "SELECT %s FROM \"reserve_policies\"",
		selectCountReservePolicies: "SELECT COUNT(\"id\") FROM \"reserve_policies\"",
		deleteReservePolicies:      "DELETE FROM \"reserve_policies\"",
		updateReservePolicies:      "UPDATE \"reserve_policies\" SET %s ",
		insertReservePolicies:      "INSERT INTO \"reserve_policies\" %s VALUES %s",
	}
)

type ReservePoliciesRepository interface {
	CreateReservePolicies(ctx context.Context, reservePolicies *model.ReservePolicies, fieldsInsert ...ReservePoliciesField) error
	BulkCreateReservePolicies(ctx context.Context, reservePoliciesList []*model.ReservePolicies, fieldsInsert ...ReservePoliciesField) error
	ResolveReservePolicies(ctx context.Context, selectFields ...ReservePoliciesField) (model.ReservePoliciesList, error)
	ResolveReservePoliciesByID(ctx context.Context, primaryID model.ReservePoliciesPrimaryID, selectFields ...ReservePoliciesField) (model.ReservePolicies, error)
	UpdateReservePoliciesByID(ctx context.Context, id model.ReservePoliciesPrimaryID, reservePolicies *model.ReservePolicies, reservePoliciesUpdateFields ...ReservePoliciesUpdateField) error
	UpdateReservePoliciesByFilter(ctx context.Context, filter model.Filter, reservePoliciesUpdateFields ...ReservePoliciesUpdateField) (rowsAffected int64, err error)
	BulkUpdateReservePolicies(ctx context.Context, reservePoliciesListMap map[model.ReservePoliciesPrimaryID]*model.ReservePolicies, ReservePoliciessMapUpdateFieldsRequest map[model.ReservePoliciesPrimaryID]ReservePoliciesUpdateFieldList) (err error)
	DeleteReservePoliciesByID(ctx context.Context, id model.ReservePoliciesPrimaryID) error
	BulkDeleteReservePoliciesByIDs(ctx context.Context, ids []model.ReservePoliciesPrimaryID) error
	ResolveReservePoliciesByFilter(ctx context.Context, filter model.Filter) (result []model.ReservePoliciesFilterResult, err error)
	IsExistReservePoliciesByIDs(ctx context.Context, ids []model.ReservePoliciesPrimaryID) (exists bool, notFoundIds []model.ReservePoliciesPrimaryID, err error)
	IsExistReservePoliciesByID(ctx context.Context, id model.ReservePoliciesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
