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

func composeInsertFieldsAndParamsFinanceManualOperations(financeManualOperationsList []model.FinanceManualOperations, fieldsInsert ...FinanceManualOperationsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceManualOperationsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeManualOperations := range financeManualOperationsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeManualOperations.Id)
			case selectField.OperationCode():
				args = append(args, financeManualOperations.OperationCode)
			case selectField.OperationType():
				args = append(args, financeManualOperations.OperationType)
			case selectField.TargetRefType():
				args = append(args, financeManualOperations.TargetRefType)
			case selectField.TargetRefId():
				args = append(args, financeManualOperations.TargetRefId)
			case selectField.RequestedBy():
				args = append(args, financeManualOperations.RequestedBy)
			case selectField.OperationStatus():
				args = append(args, financeManualOperations.OperationStatus)
			case selectField.ReasonCode():
				args = append(args, financeManualOperations.ReasonCode)
			case selectField.ReasonDetail():
				args = append(args, financeManualOperations.ReasonDetail)
			case selectField.Payload():
				args = append(args, financeManualOperations.Payload)
			case selectField.ExecutedAt():
				args = append(args, financeManualOperations.ExecutedAt)
			case selectField.Metadata():
				args = append(args, financeManualOperations.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeManualOperations.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeManualOperations.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeManualOperations.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeManualOperations.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeManualOperations.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeManualOperations.MetaDeletedBy)

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

func composeFinanceManualOperationsCompositePrimaryKeyWhere(primaryIDs []model.FinanceManualOperationsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_manual_operations\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceManualOperationsSelectFields() string {
	fields := NewFinanceManualOperationsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceManualOperationsSelectFields(selectFields ...FinanceManualOperationsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceManualOperationsField string
type FinanceManualOperationsFieldList []FinanceManualOperationsField

type FinanceManualOperationsSelectFields struct {
}

func (ss FinanceManualOperationsSelectFields) Id() FinanceManualOperationsField {
	return FinanceManualOperationsField("id")
}

func (ss FinanceManualOperationsSelectFields) OperationCode() FinanceManualOperationsField {
	return FinanceManualOperationsField("operation_code")
}

func (ss FinanceManualOperationsSelectFields) OperationType() FinanceManualOperationsField {
	return FinanceManualOperationsField("operation_type")
}

func (ss FinanceManualOperationsSelectFields) TargetRefType() FinanceManualOperationsField {
	return FinanceManualOperationsField("target_ref_type")
}

func (ss FinanceManualOperationsSelectFields) TargetRefId() FinanceManualOperationsField {
	return FinanceManualOperationsField("target_ref_id")
}

func (ss FinanceManualOperationsSelectFields) RequestedBy() FinanceManualOperationsField {
	return FinanceManualOperationsField("requested_by")
}

func (ss FinanceManualOperationsSelectFields) OperationStatus() FinanceManualOperationsField {
	return FinanceManualOperationsField("operation_status")
}

func (ss FinanceManualOperationsSelectFields) ReasonCode() FinanceManualOperationsField {
	return FinanceManualOperationsField("reason_code")
}

func (ss FinanceManualOperationsSelectFields) ReasonDetail() FinanceManualOperationsField {
	return FinanceManualOperationsField("reason_detail")
}

func (ss FinanceManualOperationsSelectFields) Payload() FinanceManualOperationsField {
	return FinanceManualOperationsField("payload")
}

func (ss FinanceManualOperationsSelectFields) ExecutedAt() FinanceManualOperationsField {
	return FinanceManualOperationsField("executed_at")
}

func (ss FinanceManualOperationsSelectFields) Metadata() FinanceManualOperationsField {
	return FinanceManualOperationsField("metadata")
}

func (ss FinanceManualOperationsSelectFields) MetaCreatedAt() FinanceManualOperationsField {
	return FinanceManualOperationsField("meta_created_at")
}

func (ss FinanceManualOperationsSelectFields) MetaCreatedBy() FinanceManualOperationsField {
	return FinanceManualOperationsField("meta_created_by")
}

func (ss FinanceManualOperationsSelectFields) MetaUpdatedAt() FinanceManualOperationsField {
	return FinanceManualOperationsField("meta_updated_at")
}

func (ss FinanceManualOperationsSelectFields) MetaUpdatedBy() FinanceManualOperationsField {
	return FinanceManualOperationsField("meta_updated_by")
}

func (ss FinanceManualOperationsSelectFields) MetaDeletedAt() FinanceManualOperationsField {
	return FinanceManualOperationsField("meta_deleted_at")
}

func (ss FinanceManualOperationsSelectFields) MetaDeletedBy() FinanceManualOperationsField {
	return FinanceManualOperationsField("meta_deleted_by")
}

func (ss FinanceManualOperationsSelectFields) All() FinanceManualOperationsFieldList {
	return []FinanceManualOperationsField{
		ss.Id(),
		ss.OperationCode(),
		ss.OperationType(),
		ss.TargetRefType(),
		ss.TargetRefId(),
		ss.RequestedBy(),
		ss.OperationStatus(),
		ss.ReasonCode(),
		ss.ReasonDetail(),
		ss.Payload(),
		ss.ExecutedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinanceManualOperationsSelectFields() FinanceManualOperationsSelectFields {
	return FinanceManualOperationsSelectFields{}
}

type FinanceManualOperationsUpdateFieldOption struct {
	useIncrement bool
}
type FinanceManualOperationsUpdateField struct {
	financeManualOperationsField FinanceManualOperationsField
	opt                          FinanceManualOperationsUpdateFieldOption
	value                        interface{}
}
type FinanceManualOperationsUpdateFieldList []FinanceManualOperationsUpdateField

func defaultFinanceManualOperationsUpdateFieldOption() FinanceManualOperationsUpdateFieldOption {
	return FinanceManualOperationsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceManualOperationsOption(useIncrement bool) func(*FinanceManualOperationsUpdateFieldOption) {
	return func(pcufo *FinanceManualOperationsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceManualOperationsUpdateField(field FinanceManualOperationsField, val interface{}, opts ...func(*FinanceManualOperationsUpdateFieldOption)) FinanceManualOperationsUpdateField {
	defaultOpt := defaultFinanceManualOperationsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceManualOperationsUpdateField{
		financeManualOperationsField: field,
		value:                        val,
		opt:                          defaultOpt,
	}
}
func defaultFinanceManualOperationsUpdateFields(financeManualOperations model.FinanceManualOperations) (financeManualOperationsUpdateFieldList FinanceManualOperationsUpdateFieldList) {
	selectFields := NewFinanceManualOperationsSelectFields()
	financeManualOperationsUpdateFieldList = append(financeManualOperationsUpdateFieldList,
		NewFinanceManualOperationsUpdateField(selectFields.Id(), financeManualOperations.Id),
		NewFinanceManualOperationsUpdateField(selectFields.OperationCode(), financeManualOperations.OperationCode),
		NewFinanceManualOperationsUpdateField(selectFields.OperationType(), financeManualOperations.OperationType),
		NewFinanceManualOperationsUpdateField(selectFields.TargetRefType(), financeManualOperations.TargetRefType),
		NewFinanceManualOperationsUpdateField(selectFields.TargetRefId(), financeManualOperations.TargetRefId),
		NewFinanceManualOperationsUpdateField(selectFields.RequestedBy(), financeManualOperations.RequestedBy),
		NewFinanceManualOperationsUpdateField(selectFields.OperationStatus(), financeManualOperations.OperationStatus),
		NewFinanceManualOperationsUpdateField(selectFields.ReasonCode(), financeManualOperations.ReasonCode),
		NewFinanceManualOperationsUpdateField(selectFields.ReasonDetail(), financeManualOperations.ReasonDetail),
		NewFinanceManualOperationsUpdateField(selectFields.Payload(), financeManualOperations.Payload),
		NewFinanceManualOperationsUpdateField(selectFields.ExecutedAt(), financeManualOperations.ExecutedAt),
		NewFinanceManualOperationsUpdateField(selectFields.Metadata(), financeManualOperations.Metadata),
		NewFinanceManualOperationsUpdateField(selectFields.MetaCreatedAt(), financeManualOperations.MetaCreatedAt),
		NewFinanceManualOperationsUpdateField(selectFields.MetaCreatedBy(), financeManualOperations.MetaCreatedBy),
		NewFinanceManualOperationsUpdateField(selectFields.MetaUpdatedAt(), financeManualOperations.MetaUpdatedAt),
		NewFinanceManualOperationsUpdateField(selectFields.MetaUpdatedBy(), financeManualOperations.MetaUpdatedBy),
		NewFinanceManualOperationsUpdateField(selectFields.MetaDeletedAt(), financeManualOperations.MetaDeletedAt),
		NewFinanceManualOperationsUpdateField(selectFields.MetaDeletedBy(), financeManualOperations.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceManualOperationsCommand(financeManualOperationsUpdateFieldList FinanceManualOperationsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeManualOperationsUpdateFieldList {
		field := string(updateField.financeManualOperationsField)
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

func (repo *RepositoryImpl) BulkCreateFinanceManualOperations(ctx context.Context, financeManualOperationsList []*model.FinanceManualOperations, fieldsInsert ...FinanceManualOperationsField) (err error) {
	var (
		fieldsStr                        string
		valueListStr                     []string
		argsList                         []interface{}
		primaryIds                       []model.FinanceManualOperationsPrimaryID
		financeManualOperationsValueList []model.FinanceManualOperations
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceManualOperationsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeManualOperations := range financeManualOperationsList {

		primaryIds = append(primaryIds, financeManualOperations.ToFinanceManualOperationsPrimaryID())

		financeManualOperationsValueList = append(financeManualOperationsValueList, *financeManualOperations)
	}

	_, notFoundIds, err := repo.IsExistFinanceManualOperationsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceManualOperations] failed checking financeManualOperations whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceManualOperationsPrimaryID{}
		mapNotFoundIds := map[model.FinanceManualOperationsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeManualOperations", fmt.Sprintf("financeManualOperations with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceManualOperations(financeManualOperationsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeManualOperationsQueries.insertFinanceManualOperations, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceManualOperations] failed exec create financeManualOperations query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceManualOperationsByIDs(ctx context.Context, primaryIDs []model.FinanceManualOperationsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceManualOperationsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceManualOperationsByIDs] failed checking financeManualOperations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeManualOperations with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_manual_operations\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeManualOperationsQueries.deleteFinanceManualOperations + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceManualOperationsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceManualOperationsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceManualOperationsByIDs(ctx context.Context, ids []model.FinanceManualOperationsPrimaryID) (exists bool, notFoundIds []model.FinanceManualOperationsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_manual_operations\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeManualOperationsQueries.selectFinanceManualOperations, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceManualOperationsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceManualOperationsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceManualOperationsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceManualOperationsPrimaryID]bool{}
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

// BulkUpdateFinanceManualOperations is used to bulk update financeManualOperations, by default it will update all field
// if want to update specific field, then fill financeManualOperationssMapUpdateFieldsRequest else please fill financeManualOperationssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceManualOperations(ctx context.Context, financeManualOperationssMap map[model.FinanceManualOperationsPrimaryID]*model.FinanceManualOperations, financeManualOperationssMapUpdateFieldsRequest map[model.FinanceManualOperationsPrimaryID]FinanceManualOperationsUpdateFieldList) (err error) {
	if len(financeManualOperationssMap) == 0 && len(financeManualOperationssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeManualOperationssMapUpdateField map[model.FinanceManualOperationsPrimaryID]FinanceManualOperationsUpdateFieldList = map[model.FinanceManualOperationsPrimaryID]FinanceManualOperationsUpdateFieldList{}
		asTableValues                          string                                                                            = "myvalues"
	)

	if len(financeManualOperationssMap) > 0 {
		for id, financeManualOperations := range financeManualOperationssMap {
			if financeManualOperations == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceManualOperations] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeManualOperationssMapUpdateField[id] = defaultFinanceManualOperationsUpdateFields(*financeManualOperations)
		}
	} else {
		financeManualOperationssMapUpdateField = financeManualOperationssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceManualOperationsQuery(financeManualOperationssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceManualOperationsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceManualOperations] failed checking financeManualOperations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeManualOperations with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceManualOperationsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_manual_operations\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceManualOperations] failed exec query")
	}
	return
}

type FinanceManualOperationsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceManualOperationsFieldParameter(param string, args ...interface{}) FinanceManualOperationsFieldParameter {
	return FinanceManualOperationsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceManualOperationsQuery(mapFinanceManualOperationss map[model.FinanceManualOperationsPrimaryID]FinanceManualOperationsUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceManualOperationsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceManualOperationsPrimaryID]map[string]interface{}{}
	financeManualOperationsSelectFields := NewFinanceManualOperationsSelectFields()
	for id, updateFields := range mapFinanceManualOperationss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeManualOperationsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceManualOperationss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceManualOperationsFieldType(updateField.financeManualOperationsField)))
			args = append(args, fields[string(updateField.financeManualOperationsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeManualOperationsField))
		if updateField.financeManualOperationsField == financeManualOperationsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeManualOperationsField, asTableValues, updateField.financeManualOperationsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeManualOperationsField,
				"\"finance_manual_operations\"", updateField.financeManualOperationsField,
				asTableValues, updateField.financeManualOperationsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceManualOperationsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceManualOperationsPrimaryID, asTableValue string) (whereQry string) {
	financeManualOperationsSelectFields := NewFinanceManualOperationsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_manual_operations\".\"id\" = %s.\"id\"::"+GetFinanceManualOperationsFieldType(financeManualOperationsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceManualOperationsFieldType(financeManualOperationsField FinanceManualOperationsField) string {
	selectFinanceManualOperationsFields := NewFinanceManualOperationsSelectFields()
	switch financeManualOperationsField {

	case selectFinanceManualOperationsFields.Id():
		return "uuid"

	case selectFinanceManualOperationsFields.OperationCode():
		return "text"

	case selectFinanceManualOperationsFields.OperationType():
		return "operation_type_enum"

	case selectFinanceManualOperationsFields.TargetRefType():
		return "text"

	case selectFinanceManualOperationsFields.TargetRefId():
		return "uuid"

	case selectFinanceManualOperationsFields.RequestedBy():
		return "uuid"

	case selectFinanceManualOperationsFields.OperationStatus():
		return "operation_status_enum"

	case selectFinanceManualOperationsFields.ReasonCode():
		return "text"

	case selectFinanceManualOperationsFields.ReasonDetail():
		return "text"

	case selectFinanceManualOperationsFields.Payload():
		return "jsonb"

	case selectFinanceManualOperationsFields.ExecutedAt():
		return "timestamptz"

	case selectFinanceManualOperationsFields.Metadata():
		return "jsonb"

	case selectFinanceManualOperationsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceManualOperationsFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceManualOperationsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceManualOperationsFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceManualOperationsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceManualOperationsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceManualOperations(ctx context.Context, financeManualOperations *model.FinanceManualOperations, fieldsInsert ...FinanceManualOperationsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceManualOperationsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceManualOperationsPrimaryID{
		Id: financeManualOperations.Id,
	}
	exists, err := repo.IsExistFinanceManualOperationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceManualOperations] failed checking financeManualOperations whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeManualOperations", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceManualOperations([]model.FinanceManualOperations{*financeManualOperations}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeManualOperationsQueries.insertFinanceManualOperations, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceManualOperations] failed exec create financeManualOperations query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceManualOperationsByID(ctx context.Context, primaryID model.FinanceManualOperationsPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceManualOperationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceManualOperationsByID] failed checking financeManualOperations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeManualOperations with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceManualOperationsCompositePrimaryKeyWhere([]model.FinanceManualOperationsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeManualOperationsQueries.deleteFinanceManualOperations + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceManualOperationsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceManualOperationsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceManualOperationsFilterResult, err error) {
	query, args, err := composeFinanceManualOperationsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceManualOperationsByFilter] failed compose financeManualOperations filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceManualOperationsByFilter] failed get financeManualOperations by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceManualOperationsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceManualOperationsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceManualOperationsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceManualOperationsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceManualOperationsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceManualOperationsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["operation_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"operation_code\"")
			selectedColumns["operation_code"] = struct{}{}
		}
		if _, selected := selectedColumns["operation_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"operation_type\"")
			selectedColumns["operation_type"] = struct{}{}
		}
		if _, selected := selectedColumns["target_ref_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"target_ref_type\"")
			selectedColumns["target_ref_type"] = struct{}{}
		}
		if _, selected := selectedColumns["target_ref_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"target_ref_id\"")
			selectedColumns["target_ref_id"] = struct{}{}
		}
		if _, selected := selectedColumns["requested_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"requested_by\"")
			selectedColumns["requested_by"] = struct{}{}
		}
		if _, selected := selectedColumns["operation_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"operation_status\"")
			selectedColumns["operation_status"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_detail"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_detail\"")
			selectedColumns["reason_detail"] = struct{}{}
		}
		if _, selected := selectedColumns["payload"]; !selected {
			selectColumns = append(selectColumns, "base.\"payload\"")
			selectedColumns["payload"] = struct{}{}
		}
		if _, selected := selectedColumns["executed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"executed_at\"")
			selectedColumns["executed_at"] = struct{}{}
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

type financeManualOperationsFilterPlaceholder struct {
	index int
}

func (p *financeManualOperationsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceManualOperationsFilterPredicate(filterField model.FilterField, placeholders *financeManualOperationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceManualOperationsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceManualOperationsFilterSQLExpr(spec)
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

func composeFinanceManualOperationsFilterGroup(group model.FilterGroup, placeholders *financeManualOperationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceManualOperationsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceManualOperationsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceManualOperationsFilterWhereQueries(filter model.Filter, placeholders *financeManualOperationsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceManualOperationsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceManualOperationsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceManualOperationsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceManualOperationsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceManualOperationsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceManualOperationsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeManualOperationsFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceManualOperationsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceManualOperationsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceManualOperationsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceManualOperationsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_manual_operations\" base%s", strings.Join(selectColumns, ","), composeFinanceManualOperationsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceManualOperationsByID(ctx context.Context, primaryID model.FinanceManualOperationsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceManualOperationsCompositePrimaryKeyWhere([]model.FinanceManualOperationsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeManualOperationsQueries.selectCountFinanceManualOperations, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceManualOperationsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceManualOperations(ctx context.Context, selectFields ...FinanceManualOperationsField) (financeManualOperationsList model.FinanceManualOperationsList, err error) {
	var (
		defaultFinanceManualOperationsSelectFields = defaultFinanceManualOperationsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceManualOperationsSelectFields = composeFinanceManualOperationsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeManualOperationsQueries.selectFinanceManualOperations, defaultFinanceManualOperationsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeManualOperationsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceManualOperations] failed get financeManualOperations list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceManualOperationsByID(ctx context.Context, primaryID model.FinanceManualOperationsPrimaryID, selectFields ...FinanceManualOperationsField) (financeManualOperations model.FinanceManualOperations, err error) {
	var (
		defaultFinanceManualOperationsSelectFields = defaultFinanceManualOperationsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceManualOperationsSelectFields = composeFinanceManualOperationsSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceManualOperationsCompositePrimaryKeyWhere([]model.FinanceManualOperationsPrimaryID{primaryID})
	query := fmt.Sprintf(financeManualOperationsQueries.selectFinanceManualOperations+" WHERE "+whereQry, defaultFinanceManualOperationsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeManualOperations, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeManualOperations with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceManualOperationsByID] failed get financeManualOperations")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceManualOperationsByID(ctx context.Context, primaryID model.FinanceManualOperationsPrimaryID, financeManualOperations *model.FinanceManualOperations, financeManualOperationsUpdateFields ...FinanceManualOperationsUpdateField) (err error) {
	exists, err := repo.IsExistFinanceManualOperationsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceManualOperations] failed checking financeManualOperations whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeManualOperations with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeManualOperations == nil {
		if len(financeManualOperationsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceManualOperationsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeManualOperations = &model.FinanceManualOperations{}
	}
	var (
		defaultFinanceManualOperationsUpdateFields = defaultFinanceManualOperationsUpdateFields(*financeManualOperations)
		tempUpdateField                            FinanceManualOperationsUpdateFieldList
		selectFields                               = NewFinanceManualOperationsSelectFields()
	)
	if len(financeManualOperationsUpdateFields) > 0 {
		for _, updateField := range financeManualOperationsUpdateFields {
			if updateField.financeManualOperationsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceManualOperationsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceManualOperationsCompositePrimaryKeyWhere([]model.FinanceManualOperationsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceManualOperationsCommand(defaultFinanceManualOperationsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeManualOperationsQueries.updateFinanceManualOperations+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceManualOperations] error when try to update financeManualOperations by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceManualOperationsByFilter(ctx context.Context, filter model.Filter, financeManualOperationsUpdateFields ...FinanceManualOperationsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeManualOperationsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceManualOperationsUpdateFieldList
		selectFields = NewFinanceManualOperationsSelectFields()
	)
	for _, updateField := range financeManualOperationsUpdateFields {
		if updateField.financeManualOperationsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceManualOperationsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeManualOperationsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceManualOperationsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_manual_operations\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceManualOperationsByFilter] error when try to update financeManualOperations by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceManualOperationsByFilter] failed get rows affected")
	}
	return
}

var (
	financeManualOperationsQueries = struct {
		selectFinanceManualOperations      string
		selectCountFinanceManualOperations string
		deleteFinanceManualOperations      string
		updateFinanceManualOperations      string
		insertFinanceManualOperations      string
	}{
		selectFinanceManualOperations:      "SELECT %s FROM \"finance_manual_operations\"",
		selectCountFinanceManualOperations: "SELECT COUNT(\"id\") FROM \"finance_manual_operations\"",
		deleteFinanceManualOperations:      "DELETE FROM \"finance_manual_operations\"",
		updateFinanceManualOperations:      "UPDATE \"finance_manual_operations\" SET %s ",
		insertFinanceManualOperations:      "INSERT INTO \"finance_manual_operations\" %s VALUES %s",
	}
)

type FinanceManualOperationsRepository interface {
	CreateFinanceManualOperations(ctx context.Context, financeManualOperations *model.FinanceManualOperations, fieldsInsert ...FinanceManualOperationsField) error
	BulkCreateFinanceManualOperations(ctx context.Context, financeManualOperationsList []*model.FinanceManualOperations, fieldsInsert ...FinanceManualOperationsField) error
	ResolveFinanceManualOperations(ctx context.Context, selectFields ...FinanceManualOperationsField) (model.FinanceManualOperationsList, error)
	ResolveFinanceManualOperationsByID(ctx context.Context, primaryID model.FinanceManualOperationsPrimaryID, selectFields ...FinanceManualOperationsField) (model.FinanceManualOperations, error)
	UpdateFinanceManualOperationsByID(ctx context.Context, id model.FinanceManualOperationsPrimaryID, financeManualOperations *model.FinanceManualOperations, financeManualOperationsUpdateFields ...FinanceManualOperationsUpdateField) error
	UpdateFinanceManualOperationsByFilter(ctx context.Context, filter model.Filter, financeManualOperationsUpdateFields ...FinanceManualOperationsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceManualOperations(ctx context.Context, financeManualOperationsListMap map[model.FinanceManualOperationsPrimaryID]*model.FinanceManualOperations, FinanceManualOperationssMapUpdateFieldsRequest map[model.FinanceManualOperationsPrimaryID]FinanceManualOperationsUpdateFieldList) (err error)
	DeleteFinanceManualOperationsByID(ctx context.Context, id model.FinanceManualOperationsPrimaryID) error
	BulkDeleteFinanceManualOperationsByIDs(ctx context.Context, ids []model.FinanceManualOperationsPrimaryID) error
	ResolveFinanceManualOperationsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceManualOperationsFilterResult, err error)
	IsExistFinanceManualOperationsByIDs(ctx context.Context, ids []model.FinanceManualOperationsPrimaryID) (exists bool, notFoundIds []model.FinanceManualOperationsPrimaryID, err error)
	IsExistFinanceManualOperationsByID(ctx context.Context, id model.FinanceManualOperationsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
