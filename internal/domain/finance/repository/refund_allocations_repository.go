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

func composeInsertFieldsAndParamsRefundAllocations(refundAllocationsList []model.RefundAllocations, fieldsInsert ...RefundAllocationsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewRefundAllocationsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, refundAllocations := range refundAllocationsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, refundAllocations.Id)
			case selectField.RefundId():
				args = append(args, refundAllocations.RefundId)
			case selectField.AllocationType():
				args = append(args, refundAllocations.AllocationType)
			case selectField.ResponsiblePartyId():
				args = append(args, refundAllocations.ResponsiblePartyId)
			case selectField.Amount():
				args = append(args, refundAllocations.Amount)
			case selectField.CurrencyCode():
				args = append(args, refundAllocations.CurrencyCode)
			case selectField.AllocationStatus():
				args = append(args, refundAllocations.AllocationStatus)
			case selectField.Metadata():
				args = append(args, refundAllocations.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, refundAllocations.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, refundAllocations.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, refundAllocations.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, refundAllocations.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, refundAllocations.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, refundAllocations.MetaDeletedBy)

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

func composeRefundAllocationsCompositePrimaryKeyWhere(primaryIDs []model.RefundAllocationsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"refund_allocations\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultRefundAllocationsSelectFields() string {
	fields := NewRefundAllocationsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeRefundAllocationsSelectFields(selectFields ...RefundAllocationsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type RefundAllocationsField string
type RefundAllocationsFieldList []RefundAllocationsField

type RefundAllocationsSelectFields struct {
}

func (ss RefundAllocationsSelectFields) Id() RefundAllocationsField {
	return RefundAllocationsField("id")
}

func (ss RefundAllocationsSelectFields) RefundId() RefundAllocationsField {
	return RefundAllocationsField("refund_id")
}

func (ss RefundAllocationsSelectFields) AllocationType() RefundAllocationsField {
	return RefundAllocationsField("allocation_type")
}

func (ss RefundAllocationsSelectFields) ResponsiblePartyId() RefundAllocationsField {
	return RefundAllocationsField("responsible_party_id")
}

func (ss RefundAllocationsSelectFields) Amount() RefundAllocationsField {
	return RefundAllocationsField("amount")
}

func (ss RefundAllocationsSelectFields) CurrencyCode() RefundAllocationsField {
	return RefundAllocationsField("currency_code")
}

func (ss RefundAllocationsSelectFields) AllocationStatus() RefundAllocationsField {
	return RefundAllocationsField("allocation_status")
}

func (ss RefundAllocationsSelectFields) Metadata() RefundAllocationsField {
	return RefundAllocationsField("metadata")
}

func (ss RefundAllocationsSelectFields) MetaCreatedAt() RefundAllocationsField {
	return RefundAllocationsField("meta_created_at")
}

func (ss RefundAllocationsSelectFields) MetaCreatedBy() RefundAllocationsField {
	return RefundAllocationsField("meta_created_by")
}

func (ss RefundAllocationsSelectFields) MetaUpdatedAt() RefundAllocationsField {
	return RefundAllocationsField("meta_updated_at")
}

func (ss RefundAllocationsSelectFields) MetaUpdatedBy() RefundAllocationsField {
	return RefundAllocationsField("meta_updated_by")
}

func (ss RefundAllocationsSelectFields) MetaDeletedAt() RefundAllocationsField {
	return RefundAllocationsField("meta_deleted_at")
}

func (ss RefundAllocationsSelectFields) MetaDeletedBy() RefundAllocationsField {
	return RefundAllocationsField("meta_deleted_by")
}

func (ss RefundAllocationsSelectFields) All() RefundAllocationsFieldList {
	return []RefundAllocationsField{
		ss.Id(),
		ss.RefundId(),
		ss.AllocationType(),
		ss.ResponsiblePartyId(),
		ss.Amount(),
		ss.CurrencyCode(),
		ss.AllocationStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewRefundAllocationsSelectFields() RefundAllocationsSelectFields {
	return RefundAllocationsSelectFields{}
}

type RefundAllocationsUpdateFieldOption struct {
	useIncrement bool
}
type RefundAllocationsUpdateField struct {
	refundAllocationsField RefundAllocationsField
	opt                    RefundAllocationsUpdateFieldOption
	value                  interface{}
}
type RefundAllocationsUpdateFieldList []RefundAllocationsUpdateField

func defaultRefundAllocationsUpdateFieldOption() RefundAllocationsUpdateFieldOption {
	return RefundAllocationsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementRefundAllocationsOption(useIncrement bool) func(*RefundAllocationsUpdateFieldOption) {
	return func(pcufo *RefundAllocationsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewRefundAllocationsUpdateField(field RefundAllocationsField, val interface{}, opts ...func(*RefundAllocationsUpdateFieldOption)) RefundAllocationsUpdateField {
	defaultOpt := defaultRefundAllocationsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return RefundAllocationsUpdateField{
		refundAllocationsField: field,
		value:                  val,
		opt:                    defaultOpt,
	}
}
func defaultRefundAllocationsUpdateFields(refundAllocations model.RefundAllocations) (refundAllocationsUpdateFieldList RefundAllocationsUpdateFieldList) {
	selectFields := NewRefundAllocationsSelectFields()
	refundAllocationsUpdateFieldList = append(refundAllocationsUpdateFieldList,
		NewRefundAllocationsUpdateField(selectFields.Id(), refundAllocations.Id),
		NewRefundAllocationsUpdateField(selectFields.RefundId(), refundAllocations.RefundId),
		NewRefundAllocationsUpdateField(selectFields.AllocationType(), refundAllocations.AllocationType),
		NewRefundAllocationsUpdateField(selectFields.ResponsiblePartyId(), refundAllocations.ResponsiblePartyId),
		NewRefundAllocationsUpdateField(selectFields.Amount(), refundAllocations.Amount),
		NewRefundAllocationsUpdateField(selectFields.CurrencyCode(), refundAllocations.CurrencyCode),
		NewRefundAllocationsUpdateField(selectFields.AllocationStatus(), refundAllocations.AllocationStatus),
		NewRefundAllocationsUpdateField(selectFields.Metadata(), refundAllocations.Metadata),
		NewRefundAllocationsUpdateField(selectFields.MetaCreatedAt(), refundAllocations.MetaCreatedAt),
		NewRefundAllocationsUpdateField(selectFields.MetaCreatedBy(), refundAllocations.MetaCreatedBy),
		NewRefundAllocationsUpdateField(selectFields.MetaUpdatedAt(), refundAllocations.MetaUpdatedAt),
		NewRefundAllocationsUpdateField(selectFields.MetaUpdatedBy(), refundAllocations.MetaUpdatedBy),
		NewRefundAllocationsUpdateField(selectFields.MetaDeletedAt(), refundAllocations.MetaDeletedAt),
		NewRefundAllocationsUpdateField(selectFields.MetaDeletedBy(), refundAllocations.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsRefundAllocationsCommand(refundAllocationsUpdateFieldList RefundAllocationsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range refundAllocationsUpdateFieldList {
		field := string(updateField.refundAllocationsField)
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

func (repo *RepositoryImpl) BulkCreateRefundAllocations(ctx context.Context, refundAllocationsList []*model.RefundAllocations, fieldsInsert ...RefundAllocationsField) (err error) {
	var (
		fieldsStr                  string
		valueListStr               []string
		argsList                   []interface{}
		primaryIds                 []model.RefundAllocationsPrimaryID
		refundAllocationsValueList []model.RefundAllocations
	)

	if len(fieldsInsert) == 0 {
		selectField := NewRefundAllocationsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, refundAllocations := range refundAllocationsList {

		primaryIds = append(primaryIds, refundAllocations.ToRefundAllocationsPrimaryID())

		refundAllocationsValueList = append(refundAllocationsValueList, *refundAllocations)
	}

	_, notFoundIds, err := repo.IsExistRefundAllocationsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundAllocations] failed checking refundAllocations whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.RefundAllocationsPrimaryID{}
		mapNotFoundIds := map[model.RefundAllocationsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "refundAllocations", fmt.Sprintf("refundAllocations with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsRefundAllocations(refundAllocationsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(refundAllocationsQueries.insertRefundAllocations, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateRefundAllocations] failed exec create refundAllocations query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteRefundAllocationsByIDs(ctx context.Context, primaryIDs []model.RefundAllocationsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistRefundAllocationsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundAllocationsByIDs] failed checking refundAllocations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAllocations with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_allocations\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(refundAllocationsQueries.deleteRefundAllocations + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundAllocationsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteRefundAllocationsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistRefundAllocationsByIDs(ctx context.Context, ids []model.RefundAllocationsPrimaryID) (exists bool, notFoundIds []model.RefundAllocationsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"refund_allocations\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(refundAllocationsQueries.selectRefundAllocations, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundAllocationsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.RefundAllocationsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundAllocationsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.RefundAllocationsPrimaryID]bool{}
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

// BulkUpdateRefundAllocations is used to bulk update refundAllocations, by default it will update all field
// if want to update specific field, then fill refundAllocationssMapUpdateFieldsRequest else please fill refundAllocationssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateRefundAllocations(ctx context.Context, refundAllocationssMap map[model.RefundAllocationsPrimaryID]*model.RefundAllocations, refundAllocationssMapUpdateFieldsRequest map[model.RefundAllocationsPrimaryID]RefundAllocationsUpdateFieldList) (err error) {
	if len(refundAllocationssMap) == 0 && len(refundAllocationssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		refundAllocationssMapUpdateField map[model.RefundAllocationsPrimaryID]RefundAllocationsUpdateFieldList = map[model.RefundAllocationsPrimaryID]RefundAllocationsUpdateFieldList{}
		asTableValues                    string                                                                = "myvalues"
	)

	if len(refundAllocationssMap) > 0 {
		for id, refundAllocations := range refundAllocationssMap {
			if refundAllocations == nil {
				log.Error().Err(err).Msg("[BulkUpdateRefundAllocations] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			refundAllocationssMapUpdateField[id] = defaultRefundAllocationsUpdateFields(*refundAllocations)
		}
	} else {
		refundAllocationssMapUpdateField = refundAllocationssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateRefundAllocationsQuery(refundAllocationssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistRefundAllocationsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundAllocations] failed checking refundAllocations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAllocations with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeRefundAllocationsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"refund_allocations\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateRefundAllocations] failed exec query")
	}
	return
}

type RefundAllocationsFieldParameter struct {
	param string
	args  []interface{}
}

func NewRefundAllocationsFieldParameter(param string, args ...interface{}) RefundAllocationsFieldParameter {
	return RefundAllocationsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateRefundAllocationsQuery(mapRefundAllocationss map[model.RefundAllocationsPrimaryID]RefundAllocationsUpdateFieldList, asTableValues string) (primaryIDs []model.RefundAllocationsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.RefundAllocationsPrimaryID]map[string]interface{}{}
	refundAllocationsSelectFields := NewRefundAllocationsSelectFields()
	for id, updateFields := range mapRefundAllocationss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.refundAllocationsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapRefundAllocationss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetRefundAllocationsFieldType(updateField.refundAllocationsField)))
			args = append(args, fields[string(updateField.refundAllocationsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.refundAllocationsField))
		if updateField.refundAllocationsField == refundAllocationsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.refundAllocationsField, asTableValues, updateField.refundAllocationsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.refundAllocationsField,
				"\"refund_allocations\"", updateField.refundAllocationsField,
				asTableValues, updateField.refundAllocationsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeRefundAllocationsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.RefundAllocationsPrimaryID, asTableValue string) (whereQry string) {
	refundAllocationsSelectFields := NewRefundAllocationsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"refund_allocations\".\"id\" = %s.\"id\"::"+GetRefundAllocationsFieldType(refundAllocationsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetRefundAllocationsFieldType(refundAllocationsField RefundAllocationsField) string {
	selectRefundAllocationsFields := NewRefundAllocationsSelectFields()
	switch refundAllocationsField {

	case selectRefundAllocationsFields.Id():
		return "uuid"

	case selectRefundAllocationsFields.RefundId():
		return "uuid"

	case selectRefundAllocationsFields.AllocationType():
		return "allocation_type_enum"

	case selectRefundAllocationsFields.ResponsiblePartyId():
		return "uuid"

	case selectRefundAllocationsFields.Amount():
		return "numeric"

	case selectRefundAllocationsFields.CurrencyCode():
		return "text"

	case selectRefundAllocationsFields.AllocationStatus():
		return "allocation_status_enum"

	case selectRefundAllocationsFields.Metadata():
		return "jsonb"

	case selectRefundAllocationsFields.MetaCreatedAt():
		return "timestamptz"

	case selectRefundAllocationsFields.MetaCreatedBy():
		return "uuid"

	case selectRefundAllocationsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectRefundAllocationsFields.MetaUpdatedBy():
		return "uuid"

	case selectRefundAllocationsFields.MetaDeletedAt():
		return "timestamptz"

	case selectRefundAllocationsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateRefundAllocations(ctx context.Context, refundAllocations *model.RefundAllocations, fieldsInsert ...RefundAllocationsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewRefundAllocationsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.RefundAllocationsPrimaryID{
		Id: refundAllocations.Id,
	}
	exists, err := repo.IsExistRefundAllocationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundAllocations] failed checking refundAllocations whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "refundAllocations", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsRefundAllocations([]model.RefundAllocations{*refundAllocations}, fieldsInsert...)
	commandQuery := fmt.Sprintf(refundAllocationsQueries.insertRefundAllocations, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateRefundAllocations] failed exec create refundAllocations query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteRefundAllocationsByID(ctx context.Context, primaryID model.RefundAllocationsPrimaryID) (err error) {
	exists, err := repo.IsExistRefundAllocationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundAllocationsByID] failed checking refundAllocations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAllocations with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeRefundAllocationsCompositePrimaryKeyWhere([]model.RefundAllocationsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(refundAllocationsQueries.deleteRefundAllocations + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteRefundAllocationsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundAllocationsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundAllocationsFilterResult, err error) {
	query, args, err := composeRefundAllocationsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundAllocationsByFilter] failed compose refundAllocations filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundAllocationsByFilter] failed get refundAllocations by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeRefundAllocationsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.RefundAllocationsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeRefundAllocationsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeRefundAllocationsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeRefundAllocationsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewRefundAllocationsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["refund_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"refund_id\"")
			selectedColumns["refund_id"] = struct{}{}
		}
		if _, selected := selectedColumns["allocation_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"allocation_type\"")
			selectedColumns["allocation_type"] = struct{}{}
		}
		if _, selected := selectedColumns["responsible_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"responsible_party_id\"")
			selectedColumns["responsible_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["allocation_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"allocation_status\"")
			selectedColumns["allocation_status"] = struct{}{}
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

type refundAllocationsFilterPlaceholder struct {
	index int
}

func (p *refundAllocationsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeRefundAllocationsFilterPredicate(filterField model.FilterField, placeholders *refundAllocationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewRefundAllocationsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeRefundAllocationsFilterSQLExpr(spec)
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

func composeRefundAllocationsFilterGroup(group model.FilterGroup, placeholders *refundAllocationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeRefundAllocationsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeRefundAllocationsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeRefundAllocationsFilterWhereQueries(filter model.Filter, placeholders *refundAllocationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeRefundAllocationsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeRefundAllocationsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeRefundAllocationsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateRefundAllocationsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeRefundAllocationsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeRefundAllocationsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := refundAllocationsFilterPlaceholder{index: 1}
	whereQueries, err := composeRefundAllocationsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewRefundAllocationsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeRefundAllocationsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeRefundAllocationsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"refund_allocations\" base%s", strings.Join(selectColumns, ","), composeRefundAllocationsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistRefundAllocationsByID(ctx context.Context, primaryID model.RefundAllocationsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeRefundAllocationsCompositePrimaryKeyWhere([]model.RefundAllocationsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", refundAllocationsQueries.selectCountRefundAllocations, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistRefundAllocationsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundAllocations(ctx context.Context, selectFields ...RefundAllocationsField) (refundAllocationsList model.RefundAllocationsList, err error) {
	var (
		defaultRefundAllocationsSelectFields = defaultRefundAllocationsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundAllocationsSelectFields = composeRefundAllocationsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(refundAllocationsQueries.selectRefundAllocations, defaultRefundAllocationsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &refundAllocationsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveRefundAllocations] failed get refundAllocations list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveRefundAllocationsByID(ctx context.Context, primaryID model.RefundAllocationsPrimaryID, selectFields ...RefundAllocationsField) (refundAllocations model.RefundAllocations, err error) {
	var (
		defaultRefundAllocationsSelectFields = defaultRefundAllocationsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultRefundAllocationsSelectFields = composeRefundAllocationsSelectFields(selectFields...)
	}
	whereQry, params := composeRefundAllocationsCompositePrimaryKeyWhere([]model.RefundAllocationsPrimaryID{primaryID})
	query := fmt.Sprintf(refundAllocationsQueries.selectRefundAllocations+" WHERE "+whereQry, defaultRefundAllocationsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &refundAllocations, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("refundAllocations with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveRefundAllocationsByID] failed get refundAllocations")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateRefundAllocationsByID(ctx context.Context, primaryID model.RefundAllocationsPrimaryID, refundAllocations *model.RefundAllocations, refundAllocationsUpdateFields ...RefundAllocationsUpdateField) (err error) {
	exists, err := repo.IsExistRefundAllocationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAllocations] failed checking refundAllocations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("refundAllocations with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if refundAllocations == nil {
		if len(refundAllocationsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateRefundAllocationsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		refundAllocations = &model.RefundAllocations{}
	}
	var (
		defaultRefundAllocationsUpdateFields = defaultRefundAllocationsUpdateFields(*refundAllocations)
		tempUpdateField                      RefundAllocationsUpdateFieldList
		selectFields                         = NewRefundAllocationsSelectFields()
	)
	if len(refundAllocationsUpdateFields) > 0 {
		for _, updateField := range refundAllocationsUpdateFields {
			if updateField.refundAllocationsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultRefundAllocationsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeRefundAllocationsCompositePrimaryKeyWhere([]model.RefundAllocationsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsRefundAllocationsCommand(defaultRefundAllocationsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(refundAllocationsQueries.updateRefundAllocations+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAllocations] error when try to update refundAllocations by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateRefundAllocationsByFilter(ctx context.Context, filter model.Filter, refundAllocationsUpdateFields ...RefundAllocationsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(refundAllocationsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields RefundAllocationsUpdateFieldList
		selectFields = NewRefundAllocationsSelectFields()
	)
	for _, updateField := range refundAllocationsUpdateFields {
		if updateField.refundAllocationsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsRefundAllocationsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := refundAllocationsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeRefundAllocationsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"refund_allocations\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAllocationsByFilter] error when try to update refundAllocations by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateRefundAllocationsByFilter] failed get rows affected")
	}
	return
}

var (
	refundAllocationsQueries = struct {
		selectRefundAllocations      string
		selectCountRefundAllocations string
		deleteRefundAllocations      string
		updateRefundAllocations      string
		insertRefundAllocations      string
	}{
		selectRefundAllocations:      "SELECT %s FROM \"refund_allocations\"",
		selectCountRefundAllocations: "SELECT COUNT(\"id\") FROM \"refund_allocations\"",
		deleteRefundAllocations:      "DELETE FROM \"refund_allocations\"",
		updateRefundAllocations:      "UPDATE \"refund_allocations\" SET %s ",
		insertRefundAllocations:      "INSERT INTO \"refund_allocations\" %s VALUES %s",
	}
)

type RefundAllocationsRepository interface {
	CreateRefundAllocations(ctx context.Context, refundAllocations *model.RefundAllocations, fieldsInsert ...RefundAllocationsField) error
	BulkCreateRefundAllocations(ctx context.Context, refundAllocationsList []*model.RefundAllocations, fieldsInsert ...RefundAllocationsField) error
	ResolveRefundAllocations(ctx context.Context, selectFields ...RefundAllocationsField) (model.RefundAllocationsList, error)
	ResolveRefundAllocationsByID(ctx context.Context, primaryID model.RefundAllocationsPrimaryID, selectFields ...RefundAllocationsField) (model.RefundAllocations, error)
	UpdateRefundAllocationsByID(ctx context.Context, id model.RefundAllocationsPrimaryID, refundAllocations *model.RefundAllocations, refundAllocationsUpdateFields ...RefundAllocationsUpdateField) error
	UpdateRefundAllocationsByFilter(ctx context.Context, filter model.Filter, refundAllocationsUpdateFields ...RefundAllocationsUpdateField) (rowsAffected int64, err error)
	BulkUpdateRefundAllocations(ctx context.Context, refundAllocationsListMap map[model.RefundAllocationsPrimaryID]*model.RefundAllocations, RefundAllocationssMapUpdateFieldsRequest map[model.RefundAllocationsPrimaryID]RefundAllocationsUpdateFieldList) (err error)
	DeleteRefundAllocationsByID(ctx context.Context, id model.RefundAllocationsPrimaryID) error
	BulkDeleteRefundAllocationsByIDs(ctx context.Context, ids []model.RefundAllocationsPrimaryID) error
	ResolveRefundAllocationsByFilter(ctx context.Context, filter model.Filter) (result []model.RefundAllocationsFilterResult, err error)
	IsExistRefundAllocationsByIDs(ctx context.Context, ids []model.RefundAllocationsPrimaryID) (exists bool, notFoundIds []model.RefundAllocationsPrimaryID, err error)
	IsExistRefundAllocationsByID(ctx context.Context, id model.RefundAllocationsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
