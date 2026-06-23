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

func composeInsertFieldsAndParamsFeeProfiles(feeProfilesList []model.FeeProfiles, fieldsInsert ...FeeProfilesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFeeProfilesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, feeProfiles := range feeProfilesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, feeProfiles.Id)
			case selectField.ProfileCode():
				args = append(args, feeProfiles.ProfileCode)
			case selectField.OwnerPartyId():
				args = append(args, feeProfiles.OwnerPartyId)
			case selectField.ProfileScope():
				args = append(args, feeProfiles.ProfileScope)
			case selectField.EffectiveStatus():
				args = append(args, feeProfiles.EffectiveStatus)
			case selectField.Description():
				args = append(args, feeProfiles.Description)
			case selectField.Metadata():
				args = append(args, feeProfiles.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, feeProfiles.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, feeProfiles.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, feeProfiles.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, feeProfiles.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, feeProfiles.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, feeProfiles.MetaDeletedBy)

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

func composeFeeProfilesCompositePrimaryKeyWhere(primaryIDs []model.FeeProfilesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"fee_profiles\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFeeProfilesSelectFields() string {
	fields := NewFeeProfilesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFeeProfilesSelectFields(selectFields ...FeeProfilesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FeeProfilesField string
type FeeProfilesFieldList []FeeProfilesField

type FeeProfilesSelectFields struct {
}

func (ss FeeProfilesSelectFields) Id() FeeProfilesField {
	return FeeProfilesField("id")
}

func (ss FeeProfilesSelectFields) ProfileCode() FeeProfilesField {
	return FeeProfilesField("profile_code")
}

func (ss FeeProfilesSelectFields) OwnerPartyId() FeeProfilesField {
	return FeeProfilesField("owner_party_id")
}

func (ss FeeProfilesSelectFields) ProfileScope() FeeProfilesField {
	return FeeProfilesField("profile_scope")
}

func (ss FeeProfilesSelectFields) EffectiveStatus() FeeProfilesField {
	return FeeProfilesField("effective_status")
}

func (ss FeeProfilesSelectFields) Description() FeeProfilesField {
	return FeeProfilesField("description")
}

func (ss FeeProfilesSelectFields) Metadata() FeeProfilesField {
	return FeeProfilesField("metadata")
}

func (ss FeeProfilesSelectFields) MetaCreatedAt() FeeProfilesField {
	return FeeProfilesField("meta_created_at")
}

func (ss FeeProfilesSelectFields) MetaCreatedBy() FeeProfilesField {
	return FeeProfilesField("meta_created_by")
}

func (ss FeeProfilesSelectFields) MetaUpdatedAt() FeeProfilesField {
	return FeeProfilesField("meta_updated_at")
}

func (ss FeeProfilesSelectFields) MetaUpdatedBy() FeeProfilesField {
	return FeeProfilesField("meta_updated_by")
}

func (ss FeeProfilesSelectFields) MetaDeletedAt() FeeProfilesField {
	return FeeProfilesField("meta_deleted_at")
}

func (ss FeeProfilesSelectFields) MetaDeletedBy() FeeProfilesField {
	return FeeProfilesField("meta_deleted_by")
}

func (ss FeeProfilesSelectFields) All() FeeProfilesFieldList {
	return []FeeProfilesField{
		ss.Id(),
		ss.ProfileCode(),
		ss.OwnerPartyId(),
		ss.ProfileScope(),
		ss.EffectiveStatus(),
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

func NewFeeProfilesSelectFields() FeeProfilesSelectFields {
	return FeeProfilesSelectFields{}
}

type FeeProfilesUpdateFieldOption struct {
	useIncrement bool
}
type FeeProfilesUpdateField struct {
	feeProfilesField FeeProfilesField
	opt              FeeProfilesUpdateFieldOption
	value            interface{}
}
type FeeProfilesUpdateFieldList []FeeProfilesUpdateField

func defaultFeeProfilesUpdateFieldOption() FeeProfilesUpdateFieldOption {
	return FeeProfilesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFeeProfilesOption(useIncrement bool) func(*FeeProfilesUpdateFieldOption) {
	return func(pcufo *FeeProfilesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFeeProfilesUpdateField(field FeeProfilesField, val interface{}, opts ...func(*FeeProfilesUpdateFieldOption)) FeeProfilesUpdateField {
	defaultOpt := defaultFeeProfilesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FeeProfilesUpdateField{
		feeProfilesField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultFeeProfilesUpdateFields(feeProfiles model.FeeProfiles) (feeProfilesUpdateFieldList FeeProfilesUpdateFieldList) {
	selectFields := NewFeeProfilesSelectFields()
	feeProfilesUpdateFieldList = append(feeProfilesUpdateFieldList,
		NewFeeProfilesUpdateField(selectFields.Id(), feeProfiles.Id),
		NewFeeProfilesUpdateField(selectFields.ProfileCode(), feeProfiles.ProfileCode),
		NewFeeProfilesUpdateField(selectFields.OwnerPartyId(), feeProfiles.OwnerPartyId),
		NewFeeProfilesUpdateField(selectFields.ProfileScope(), feeProfiles.ProfileScope),
		NewFeeProfilesUpdateField(selectFields.EffectiveStatus(), feeProfiles.EffectiveStatus),
		NewFeeProfilesUpdateField(selectFields.Description(), feeProfiles.Description),
		NewFeeProfilesUpdateField(selectFields.Metadata(), feeProfiles.Metadata),
		NewFeeProfilesUpdateField(selectFields.MetaCreatedAt(), feeProfiles.MetaCreatedAt),
		NewFeeProfilesUpdateField(selectFields.MetaCreatedBy(), feeProfiles.MetaCreatedBy),
		NewFeeProfilesUpdateField(selectFields.MetaUpdatedAt(), feeProfiles.MetaUpdatedAt),
		NewFeeProfilesUpdateField(selectFields.MetaUpdatedBy(), feeProfiles.MetaUpdatedBy),
		NewFeeProfilesUpdateField(selectFields.MetaDeletedAt(), feeProfiles.MetaDeletedAt),
		NewFeeProfilesUpdateField(selectFields.MetaDeletedBy(), feeProfiles.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFeeProfilesCommand(feeProfilesUpdateFieldList FeeProfilesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range feeProfilesUpdateFieldList {
		field := string(updateField.feeProfilesField)
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

func (repo *RepositoryImpl) BulkCreateFeeProfiles(ctx context.Context, feeProfilesList []*model.FeeProfiles, fieldsInsert ...FeeProfilesField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.FeeProfilesPrimaryID
		feeProfilesValueList []model.FeeProfiles
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFeeProfilesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, feeProfiles := range feeProfilesList {

		primaryIds = append(primaryIds, feeProfiles.ToFeeProfilesPrimaryID())

		feeProfilesValueList = append(feeProfilesValueList, *feeProfiles)
	}

	_, notFoundIds, err := repo.IsExistFeeProfilesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeProfiles] failed checking feeProfiles whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FeeProfilesPrimaryID{}
		mapNotFoundIds := map[model.FeeProfilesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "feeProfiles", fmt.Sprintf("feeProfiles with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFeeProfiles(feeProfilesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(feeProfilesQueries.insertFeeProfiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeProfiles] failed exec create feeProfiles query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFeeProfilesByIDs(ctx context.Context, primaryIDs []model.FeeProfilesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFeeProfilesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeProfilesByIDs] failed checking feeProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeProfiles with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_profiles\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(feeProfilesQueries.deleteFeeProfiles + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeProfilesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeProfilesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFeeProfilesByIDs(ctx context.Context, ids []model.FeeProfilesPrimaryID) (exists bool, notFoundIds []model.FeeProfilesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_profiles\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(feeProfilesQueries.selectFeeProfiles, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeProfilesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FeeProfilesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeProfilesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FeeProfilesPrimaryID]bool{}
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

// BulkUpdateFeeProfiles is used to bulk update feeProfiles, by default it will update all field
// if want to update specific field, then fill feeProfilessMapUpdateFieldsRequest else please fill feeProfilessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFeeProfiles(ctx context.Context, feeProfilessMap map[model.FeeProfilesPrimaryID]*model.FeeProfiles, feeProfilessMapUpdateFieldsRequest map[model.FeeProfilesPrimaryID]FeeProfilesUpdateFieldList) (err error) {
	if len(feeProfilessMap) == 0 && len(feeProfilessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		feeProfilessMapUpdateField map[model.FeeProfilesPrimaryID]FeeProfilesUpdateFieldList = map[model.FeeProfilesPrimaryID]FeeProfilesUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(feeProfilessMap) > 0 {
		for id, feeProfiles := range feeProfilessMap {
			if feeProfiles == nil {
				log.Error().Err(err).Msg("[BulkUpdateFeeProfiles] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			feeProfilessMapUpdateField[id] = defaultFeeProfilesUpdateFields(*feeProfiles)
		}
	} else {
		feeProfilessMapUpdateField = feeProfilessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFeeProfilesQuery(feeProfilessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFeeProfilesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeProfiles] failed checking feeProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeProfiles with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFeeProfilesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"fee_profiles\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeProfiles] failed exec query")
	}
	return
}

type FeeProfilesFieldParameter struct {
	param string
	args  []interface{}
}

func NewFeeProfilesFieldParameter(param string, args ...interface{}) FeeProfilesFieldParameter {
	return FeeProfilesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFeeProfilesQuery(mapFeeProfiless map[model.FeeProfilesPrimaryID]FeeProfilesUpdateFieldList, asTableValues string) (primaryIDs []model.FeeProfilesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FeeProfilesPrimaryID]map[string]interface{}{}
	feeProfilesSelectFields := NewFeeProfilesSelectFields()
	for id, updateFields := range mapFeeProfiless {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.feeProfilesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFeeProfiless[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFeeProfilesFieldType(updateField.feeProfilesField)))
			args = append(args, fields[string(updateField.feeProfilesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.feeProfilesField))
		if updateField.feeProfilesField == feeProfilesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.feeProfilesField, asTableValues, updateField.feeProfilesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.feeProfilesField,
				"\"fee_profiles\"", updateField.feeProfilesField,
				asTableValues, updateField.feeProfilesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFeeProfilesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FeeProfilesPrimaryID, asTableValue string) (whereQry string) {
	feeProfilesSelectFields := NewFeeProfilesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"fee_profiles\".\"id\" = %s.\"id\"::"+GetFeeProfilesFieldType(feeProfilesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFeeProfilesFieldType(feeProfilesField FeeProfilesField) string {
	selectFeeProfilesFields := NewFeeProfilesSelectFields()
	switch feeProfilesField {

	case selectFeeProfilesFields.Id():
		return "uuid"

	case selectFeeProfilesFields.ProfileCode():
		return "text"

	case selectFeeProfilesFields.OwnerPartyId():
		return "uuid"

	case selectFeeProfilesFields.ProfileScope():
		return "profile_scope_enum"

	case selectFeeProfilesFields.EffectiveStatus():
		return "effective_status_enum"

	case selectFeeProfilesFields.Description():
		return "text"

	case selectFeeProfilesFields.Metadata():
		return "jsonb"

	case selectFeeProfilesFields.MetaCreatedAt():
		return "timestamptz"

	case selectFeeProfilesFields.MetaCreatedBy():
		return "uuid"

	case selectFeeProfilesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFeeProfilesFields.MetaUpdatedBy():
		return "uuid"

	case selectFeeProfilesFields.MetaDeletedAt():
		return "timestamptz"

	case selectFeeProfilesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFeeProfiles(ctx context.Context, feeProfiles *model.FeeProfiles, fieldsInsert ...FeeProfilesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFeeProfilesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FeeProfilesPrimaryID{
		Id: feeProfiles.Id,
	}
	exists, err := repo.IsExistFeeProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeProfiles] failed checking feeProfiles whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "feeProfiles", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFeeProfiles([]model.FeeProfiles{*feeProfiles}, fieldsInsert...)
	commandQuery := fmt.Sprintf(feeProfilesQueries.insertFeeProfiles, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeProfiles] failed exec create feeProfiles query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFeeProfilesByID(ctx context.Context, primaryID model.FeeProfilesPrimaryID) (err error) {
	exists, err := repo.IsExistFeeProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeProfilesByID] failed checking feeProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeProfiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFeeProfilesCompositePrimaryKeyWhere([]model.FeeProfilesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(feeProfilesQueries.deleteFeeProfiles + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeProfilesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeProfilesByFilter(ctx context.Context, filter model.Filter) (result []model.FeeProfilesFilterResult, err error) {
	query, args, err := composeFeeProfilesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeProfilesByFilter] failed compose feeProfiles filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeProfilesByFilter] failed get feeProfiles by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFeeProfilesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FeeProfilesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFeeProfilesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFeeProfilesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFeeProfilesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 13 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFeeProfilesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["profile_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"profile_code\"")
			selectedColumns["profile_code"] = struct{}{}
		}
		if _, selected := selectedColumns["owner_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"owner_party_id\"")
			selectedColumns["owner_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["profile_scope"]; !selected {
			selectColumns = append(selectColumns, "base.\"profile_scope\"")
			selectedColumns["profile_scope"] = struct{}{}
		}
		if _, selected := selectedColumns["effective_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"effective_status\"")
			selectedColumns["effective_status"] = struct{}{}
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

type feeProfilesFilterPlaceholder struct {
	index int
}

func (p *feeProfilesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFeeProfilesFilterPredicate(filterField model.FilterField, placeholders *feeProfilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFeeProfilesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFeeProfilesFilterSQLExpr(spec)
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

func composeFeeProfilesFilterGroup(group model.FilterGroup, placeholders *feeProfilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFeeProfilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFeeProfilesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFeeProfilesFilterWhereQueries(filter model.Filter, placeholders *feeProfilesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFeeProfilesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFeeProfilesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFeeProfilesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFeeProfilesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFeeProfilesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFeeProfilesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := feeProfilesFilterPlaceholder{index: 1}
	whereQueries, err := composeFeeProfilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFeeProfilesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFeeProfilesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFeeProfilesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"fee_profiles\" base%s", strings.Join(selectColumns, ","), composeFeeProfilesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFeeProfilesByID(ctx context.Context, primaryID model.FeeProfilesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFeeProfilesCompositePrimaryKeyWhere([]model.FeeProfilesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", feeProfilesQueries.selectCountFeeProfiles, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeProfilesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeProfiles(ctx context.Context, selectFields ...FeeProfilesField) (feeProfilesList model.FeeProfilesList, err error) {
	var (
		defaultFeeProfilesSelectFields = defaultFeeProfilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeProfilesSelectFields = composeFeeProfilesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(feeProfilesQueries.selectFeeProfiles, defaultFeeProfilesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &feeProfilesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeProfiles] failed get feeProfiles list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeProfilesByID(ctx context.Context, primaryID model.FeeProfilesPrimaryID, selectFields ...FeeProfilesField) (feeProfiles model.FeeProfiles, err error) {
	var (
		defaultFeeProfilesSelectFields = defaultFeeProfilesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeProfilesSelectFields = composeFeeProfilesSelectFields(selectFields...)
	}
	whereQry, params := composeFeeProfilesCompositePrimaryKeyWhere([]model.FeeProfilesPrimaryID{primaryID})
	query := fmt.Sprintf(feeProfilesQueries.selectFeeProfiles+" WHERE "+whereQry, defaultFeeProfilesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &feeProfiles, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("feeProfiles with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFeeProfilesByID] failed get feeProfiles")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFeeProfilesByID(ctx context.Context, primaryID model.FeeProfilesPrimaryID, feeProfiles *model.FeeProfiles, feeProfilesUpdateFields ...FeeProfilesUpdateField) (err error) {
	exists, err := repo.IsExistFeeProfilesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeProfiles] failed checking feeProfiles whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeProfiles with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if feeProfiles == nil {
		if len(feeProfilesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFeeProfilesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		feeProfiles = &model.FeeProfiles{}
	}
	var (
		defaultFeeProfilesUpdateFields = defaultFeeProfilesUpdateFields(*feeProfiles)
		tempUpdateField                FeeProfilesUpdateFieldList
		selectFields                   = NewFeeProfilesSelectFields()
	)
	if len(feeProfilesUpdateFields) > 0 {
		for _, updateField := range feeProfilesUpdateFields {
			if updateField.feeProfilesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFeeProfilesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFeeProfilesCompositePrimaryKeyWhere([]model.FeeProfilesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFeeProfilesCommand(defaultFeeProfilesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(feeProfilesQueries.updateFeeProfiles+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeProfiles] error when try to update feeProfiles by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFeeProfilesByFilter(ctx context.Context, filter model.Filter, feeProfilesUpdateFields ...FeeProfilesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(feeProfilesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FeeProfilesUpdateFieldList
		selectFields = NewFeeProfilesSelectFields()
	)
	for _, updateField := range feeProfilesUpdateFields {
		if updateField.feeProfilesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFeeProfilesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := feeProfilesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFeeProfilesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"fee_profiles\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeProfilesByFilter] error when try to update feeProfiles by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeProfilesByFilter] failed get rows affected")
	}
	return
}

var (
	feeProfilesQueries = struct {
		selectFeeProfiles      string
		selectCountFeeProfiles string
		deleteFeeProfiles      string
		updateFeeProfiles      string
		insertFeeProfiles      string
	}{
		selectFeeProfiles:      "SELECT %s FROM \"fee_profiles\"",
		selectCountFeeProfiles: "SELECT COUNT(\"id\") FROM \"fee_profiles\"",
		deleteFeeProfiles:      "DELETE FROM \"fee_profiles\"",
		updateFeeProfiles:      "UPDATE \"fee_profiles\" SET %s ",
		insertFeeProfiles:      "INSERT INTO \"fee_profiles\" %s VALUES %s",
	}
)

type FeeProfilesRepository interface {
	CreateFeeProfiles(ctx context.Context, feeProfiles *model.FeeProfiles, fieldsInsert ...FeeProfilesField) error
	BulkCreateFeeProfiles(ctx context.Context, feeProfilesList []*model.FeeProfiles, fieldsInsert ...FeeProfilesField) error
	ResolveFeeProfiles(ctx context.Context, selectFields ...FeeProfilesField) (model.FeeProfilesList, error)
	ResolveFeeProfilesByID(ctx context.Context, primaryID model.FeeProfilesPrimaryID, selectFields ...FeeProfilesField) (model.FeeProfiles, error)
	UpdateFeeProfilesByID(ctx context.Context, id model.FeeProfilesPrimaryID, feeProfiles *model.FeeProfiles, feeProfilesUpdateFields ...FeeProfilesUpdateField) error
	UpdateFeeProfilesByFilter(ctx context.Context, filter model.Filter, feeProfilesUpdateFields ...FeeProfilesUpdateField) (rowsAffected int64, err error)
	BulkUpdateFeeProfiles(ctx context.Context, feeProfilesListMap map[model.FeeProfilesPrimaryID]*model.FeeProfiles, FeeProfilessMapUpdateFieldsRequest map[model.FeeProfilesPrimaryID]FeeProfilesUpdateFieldList) (err error)
	DeleteFeeProfilesByID(ctx context.Context, id model.FeeProfilesPrimaryID) error
	BulkDeleteFeeProfilesByIDs(ctx context.Context, ids []model.FeeProfilesPrimaryID) error
	ResolveFeeProfilesByFilter(ctx context.Context, filter model.Filter) (result []model.FeeProfilesFilterResult, err error)
	IsExistFeeProfilesByIDs(ctx context.Context, ids []model.FeeProfilesPrimaryID) (exists bool, notFoundIds []model.FeeProfilesPrimaryID, err error)
	IsExistFeeProfilesByID(ctx context.Context, id model.FeeProfilesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
