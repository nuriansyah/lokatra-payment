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

func composeInsertFieldsAndParamsFinanceMetricSnapshots(financeMetricSnapshotsList []model.FinanceMetricSnapshots, fieldsInsert ...FinanceMetricSnapshotsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFinanceMetricSnapshotsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, financeMetricSnapshots := range financeMetricSnapshotsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, financeMetricSnapshots.Id)
			case selectField.MetricName():
				args = append(args, financeMetricSnapshots.MetricName)
			case selectField.MetricScope():
				args = append(args, financeMetricSnapshots.MetricScope)
			case selectField.ScopeRef():
				args = append(args, financeMetricSnapshots.ScopeRef)
			case selectField.PeriodStart():
				args = append(args, financeMetricSnapshots.PeriodStart)
			case selectField.PeriodEnd():
				args = append(args, financeMetricSnapshots.PeriodEnd)
			case selectField.MetricValue():
				args = append(args, financeMetricSnapshots.MetricValue)
			case selectField.MetricUnit():
				args = append(args, financeMetricSnapshots.MetricUnit)
			case selectField.Dimensions():
				args = append(args, financeMetricSnapshots.Dimensions)
			case selectField.Metadata():
				args = append(args, financeMetricSnapshots.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, financeMetricSnapshots.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, financeMetricSnapshots.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, financeMetricSnapshots.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, financeMetricSnapshots.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, financeMetricSnapshots.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, financeMetricSnapshots.MetaDeletedBy)

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

func composeFinanceMetricSnapshotsCompositePrimaryKeyWhere(primaryIDs []model.FinanceMetricSnapshotsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"finance_metric_snapshots\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFinanceMetricSnapshotsSelectFields() string {
	fields := NewFinanceMetricSnapshotsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFinanceMetricSnapshotsSelectFields(selectFields ...FinanceMetricSnapshotsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FinanceMetricSnapshotsField string
type FinanceMetricSnapshotsFieldList []FinanceMetricSnapshotsField

type FinanceMetricSnapshotsSelectFields struct {
}

func (ss FinanceMetricSnapshotsSelectFields) Id() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("id")
}

func (ss FinanceMetricSnapshotsSelectFields) MetricName() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("metric_name")
}

func (ss FinanceMetricSnapshotsSelectFields) MetricScope() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("metric_scope")
}

func (ss FinanceMetricSnapshotsSelectFields) ScopeRef() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("scope_ref")
}

func (ss FinanceMetricSnapshotsSelectFields) PeriodStart() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("period_start")
}

func (ss FinanceMetricSnapshotsSelectFields) PeriodEnd() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("period_end")
}

func (ss FinanceMetricSnapshotsSelectFields) MetricValue() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("metric_value")
}

func (ss FinanceMetricSnapshotsSelectFields) MetricUnit() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("metric_unit")
}

func (ss FinanceMetricSnapshotsSelectFields) Dimensions() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("dimensions")
}

func (ss FinanceMetricSnapshotsSelectFields) Metadata() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("metadata")
}

func (ss FinanceMetricSnapshotsSelectFields) MetaCreatedAt() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("meta_created_at")
}

func (ss FinanceMetricSnapshotsSelectFields) MetaCreatedBy() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("meta_created_by")
}

func (ss FinanceMetricSnapshotsSelectFields) MetaUpdatedAt() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("meta_updated_at")
}

func (ss FinanceMetricSnapshotsSelectFields) MetaUpdatedBy() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("meta_updated_by")
}

func (ss FinanceMetricSnapshotsSelectFields) MetaDeletedAt() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("meta_deleted_at")
}

func (ss FinanceMetricSnapshotsSelectFields) MetaDeletedBy() FinanceMetricSnapshotsField {
	return FinanceMetricSnapshotsField("meta_deleted_by")
}

func (ss FinanceMetricSnapshotsSelectFields) All() FinanceMetricSnapshotsFieldList {
	return []FinanceMetricSnapshotsField{
		ss.Id(),
		ss.MetricName(),
		ss.MetricScope(),
		ss.ScopeRef(),
		ss.PeriodStart(),
		ss.PeriodEnd(),
		ss.MetricValue(),
		ss.MetricUnit(),
		ss.Dimensions(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFinanceMetricSnapshotsSelectFields() FinanceMetricSnapshotsSelectFields {
	return FinanceMetricSnapshotsSelectFields{}
}

type FinanceMetricSnapshotsUpdateFieldOption struct {
	useIncrement bool
}
type FinanceMetricSnapshotsUpdateField struct {
	financeMetricSnapshotsField FinanceMetricSnapshotsField
	opt                         FinanceMetricSnapshotsUpdateFieldOption
	value                       interface{}
}
type FinanceMetricSnapshotsUpdateFieldList []FinanceMetricSnapshotsUpdateField

func defaultFinanceMetricSnapshotsUpdateFieldOption() FinanceMetricSnapshotsUpdateFieldOption {
	return FinanceMetricSnapshotsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFinanceMetricSnapshotsOption(useIncrement bool) func(*FinanceMetricSnapshotsUpdateFieldOption) {
	return func(pcufo *FinanceMetricSnapshotsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFinanceMetricSnapshotsUpdateField(field FinanceMetricSnapshotsField, val interface{}, opts ...func(*FinanceMetricSnapshotsUpdateFieldOption)) FinanceMetricSnapshotsUpdateField {
	defaultOpt := defaultFinanceMetricSnapshotsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FinanceMetricSnapshotsUpdateField{
		financeMetricSnapshotsField: field,
		value:                       val,
		opt:                         defaultOpt,
	}
}
func defaultFinanceMetricSnapshotsUpdateFields(financeMetricSnapshots model.FinanceMetricSnapshots) (financeMetricSnapshotsUpdateFieldList FinanceMetricSnapshotsUpdateFieldList) {
	selectFields := NewFinanceMetricSnapshotsSelectFields()
	financeMetricSnapshotsUpdateFieldList = append(financeMetricSnapshotsUpdateFieldList,
		NewFinanceMetricSnapshotsUpdateField(selectFields.Id(), financeMetricSnapshots.Id),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetricName(), financeMetricSnapshots.MetricName),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetricScope(), financeMetricSnapshots.MetricScope),
		NewFinanceMetricSnapshotsUpdateField(selectFields.ScopeRef(), financeMetricSnapshots.ScopeRef),
		NewFinanceMetricSnapshotsUpdateField(selectFields.PeriodStart(), financeMetricSnapshots.PeriodStart),
		NewFinanceMetricSnapshotsUpdateField(selectFields.PeriodEnd(), financeMetricSnapshots.PeriodEnd),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetricValue(), financeMetricSnapshots.MetricValue),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetricUnit(), financeMetricSnapshots.MetricUnit),
		NewFinanceMetricSnapshotsUpdateField(selectFields.Dimensions(), financeMetricSnapshots.Dimensions),
		NewFinanceMetricSnapshotsUpdateField(selectFields.Metadata(), financeMetricSnapshots.Metadata),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetaCreatedAt(), financeMetricSnapshots.MetaCreatedAt),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetaCreatedBy(), financeMetricSnapshots.MetaCreatedBy),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetaUpdatedAt(), financeMetricSnapshots.MetaUpdatedAt),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetaUpdatedBy(), financeMetricSnapshots.MetaUpdatedBy),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetaDeletedAt(), financeMetricSnapshots.MetaDeletedAt),
		NewFinanceMetricSnapshotsUpdateField(selectFields.MetaDeletedBy(), financeMetricSnapshots.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFinanceMetricSnapshotsCommand(financeMetricSnapshotsUpdateFieldList FinanceMetricSnapshotsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range financeMetricSnapshotsUpdateFieldList {
		field := string(updateField.financeMetricSnapshotsField)
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

func (repo *RepositoryImpl) BulkCreateFinanceMetricSnapshots(ctx context.Context, financeMetricSnapshotsList []*model.FinanceMetricSnapshots, fieldsInsert ...FinanceMetricSnapshotsField) (err error) {
	var (
		fieldsStr                       string
		valueListStr                    []string
		argsList                        []interface{}
		primaryIds                      []model.FinanceMetricSnapshotsPrimaryID
		financeMetricSnapshotsValueList []model.FinanceMetricSnapshots
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFinanceMetricSnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, financeMetricSnapshots := range financeMetricSnapshotsList {

		primaryIds = append(primaryIds, financeMetricSnapshots.ToFinanceMetricSnapshotsPrimaryID())

		financeMetricSnapshotsValueList = append(financeMetricSnapshotsValueList, *financeMetricSnapshots)
	}

	_, notFoundIds, err := repo.IsExistFinanceMetricSnapshotsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceMetricSnapshots] failed checking financeMetricSnapshots whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FinanceMetricSnapshotsPrimaryID{}
		mapNotFoundIds := map[model.FinanceMetricSnapshotsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "financeMetricSnapshots", fmt.Sprintf("financeMetricSnapshots with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFinanceMetricSnapshots(financeMetricSnapshotsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(financeMetricSnapshotsQueries.insertFinanceMetricSnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFinanceMetricSnapshots] failed exec create financeMetricSnapshots query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFinanceMetricSnapshotsByIDs(ctx context.Context, primaryIDs []model.FinanceMetricSnapshotsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFinanceMetricSnapshotsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceMetricSnapshotsByIDs] failed checking financeMetricSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeMetricSnapshots with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_metric_snapshots\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(financeMetricSnapshotsQueries.deleteFinanceMetricSnapshots + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceMetricSnapshotsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFinanceMetricSnapshotsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFinanceMetricSnapshotsByIDs(ctx context.Context, ids []model.FinanceMetricSnapshotsPrimaryID) (exists bool, notFoundIds []model.FinanceMetricSnapshotsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"finance_metric_snapshots\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(financeMetricSnapshotsQueries.selectFinanceMetricSnapshots, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceMetricSnapshotsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FinanceMetricSnapshotsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceMetricSnapshotsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FinanceMetricSnapshotsPrimaryID]bool{}
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

// BulkUpdateFinanceMetricSnapshots is used to bulk update financeMetricSnapshots, by default it will update all field
// if want to update specific field, then fill financeMetricSnapshotssMapUpdateFieldsRequest else please fill financeMetricSnapshotssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFinanceMetricSnapshots(ctx context.Context, financeMetricSnapshotssMap map[model.FinanceMetricSnapshotsPrimaryID]*model.FinanceMetricSnapshots, financeMetricSnapshotssMapUpdateFieldsRequest map[model.FinanceMetricSnapshotsPrimaryID]FinanceMetricSnapshotsUpdateFieldList) (err error) {
	if len(financeMetricSnapshotssMap) == 0 && len(financeMetricSnapshotssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		financeMetricSnapshotssMapUpdateField map[model.FinanceMetricSnapshotsPrimaryID]FinanceMetricSnapshotsUpdateFieldList = map[model.FinanceMetricSnapshotsPrimaryID]FinanceMetricSnapshotsUpdateFieldList{}
		asTableValues                         string                                                                          = "myvalues"
	)

	if len(financeMetricSnapshotssMap) > 0 {
		for id, financeMetricSnapshots := range financeMetricSnapshotssMap {
			if financeMetricSnapshots == nil {
				log.Error().Err(err).Msg("[BulkUpdateFinanceMetricSnapshots] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			financeMetricSnapshotssMapUpdateField[id] = defaultFinanceMetricSnapshotsUpdateFields(*financeMetricSnapshots)
		}
	} else {
		financeMetricSnapshotssMapUpdateField = financeMetricSnapshotssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFinanceMetricSnapshotsQuery(financeMetricSnapshotssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFinanceMetricSnapshotsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceMetricSnapshots] failed checking financeMetricSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeMetricSnapshots with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFinanceMetricSnapshotsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"finance_metric_snapshots\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFinanceMetricSnapshots] failed exec query")
	}
	return
}

type FinanceMetricSnapshotsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFinanceMetricSnapshotsFieldParameter(param string, args ...interface{}) FinanceMetricSnapshotsFieldParameter {
	return FinanceMetricSnapshotsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFinanceMetricSnapshotsQuery(mapFinanceMetricSnapshotss map[model.FinanceMetricSnapshotsPrimaryID]FinanceMetricSnapshotsUpdateFieldList, asTableValues string) (primaryIDs []model.FinanceMetricSnapshotsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FinanceMetricSnapshotsPrimaryID]map[string]interface{}{}
	financeMetricSnapshotsSelectFields := NewFinanceMetricSnapshotsSelectFields()
	for id, updateFields := range mapFinanceMetricSnapshotss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.financeMetricSnapshotsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFinanceMetricSnapshotss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFinanceMetricSnapshotsFieldType(updateField.financeMetricSnapshotsField)))
			args = append(args, fields[string(updateField.financeMetricSnapshotsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.financeMetricSnapshotsField))
		if updateField.financeMetricSnapshotsField == financeMetricSnapshotsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.financeMetricSnapshotsField, asTableValues, updateField.financeMetricSnapshotsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.financeMetricSnapshotsField,
				"\"finance_metric_snapshots\"", updateField.financeMetricSnapshotsField,
				asTableValues, updateField.financeMetricSnapshotsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFinanceMetricSnapshotsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FinanceMetricSnapshotsPrimaryID, asTableValue string) (whereQry string) {
	financeMetricSnapshotsSelectFields := NewFinanceMetricSnapshotsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"finance_metric_snapshots\".\"id\" = %s.\"id\"::"+GetFinanceMetricSnapshotsFieldType(financeMetricSnapshotsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFinanceMetricSnapshotsFieldType(financeMetricSnapshotsField FinanceMetricSnapshotsField) string {
	selectFinanceMetricSnapshotsFields := NewFinanceMetricSnapshotsSelectFields()
	switch financeMetricSnapshotsField {

	case selectFinanceMetricSnapshotsFields.Id():
		return "uuid"

	case selectFinanceMetricSnapshotsFields.MetricName():
		return "text"

	case selectFinanceMetricSnapshotsFields.MetricScope():
		return "text"

	case selectFinanceMetricSnapshotsFields.ScopeRef():
		return "text"

	case selectFinanceMetricSnapshotsFields.PeriodStart():
		return "timestamptz"

	case selectFinanceMetricSnapshotsFields.PeriodEnd():
		return "timestamptz"

	case selectFinanceMetricSnapshotsFields.MetricValue():
		return "numeric"

	case selectFinanceMetricSnapshotsFields.MetricUnit():
		return "text"

	case selectFinanceMetricSnapshotsFields.Dimensions():
		return "jsonb"

	case selectFinanceMetricSnapshotsFields.Metadata():
		return "jsonb"

	case selectFinanceMetricSnapshotsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFinanceMetricSnapshotsFields.MetaCreatedBy():
		return "uuid"

	case selectFinanceMetricSnapshotsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFinanceMetricSnapshotsFields.MetaUpdatedBy():
		return "uuid"

	case selectFinanceMetricSnapshotsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFinanceMetricSnapshotsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFinanceMetricSnapshots(ctx context.Context, financeMetricSnapshots *model.FinanceMetricSnapshots, fieldsInsert ...FinanceMetricSnapshotsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFinanceMetricSnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FinanceMetricSnapshotsPrimaryID{
		Id: financeMetricSnapshots.Id,
	}
	exists, err := repo.IsExistFinanceMetricSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceMetricSnapshots] failed checking financeMetricSnapshots whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "financeMetricSnapshots", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFinanceMetricSnapshots([]model.FinanceMetricSnapshots{*financeMetricSnapshots}, fieldsInsert...)
	commandQuery := fmt.Sprintf(financeMetricSnapshotsQueries.insertFinanceMetricSnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFinanceMetricSnapshots] failed exec create financeMetricSnapshots query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFinanceMetricSnapshotsByID(ctx context.Context, primaryID model.FinanceMetricSnapshotsPrimaryID) (err error) {
	exists, err := repo.IsExistFinanceMetricSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceMetricSnapshotsByID] failed checking financeMetricSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeMetricSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFinanceMetricSnapshotsCompositePrimaryKeyWhere([]model.FinanceMetricSnapshotsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(financeMetricSnapshotsQueries.deleteFinanceMetricSnapshots + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFinanceMetricSnapshotsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceMetricSnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceMetricSnapshotsFilterResult, err error) {
	query, args, err := composeFinanceMetricSnapshotsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceMetricSnapshotsByFilter] failed compose financeMetricSnapshots filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceMetricSnapshotsByFilter] failed get financeMetricSnapshots by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFinanceMetricSnapshotsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FinanceMetricSnapshotsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFinanceMetricSnapshotsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFinanceMetricSnapshotsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFinanceMetricSnapshotsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 16 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFinanceMetricSnapshotsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["metric_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"metric_name\"")
			selectedColumns["metric_name"] = struct{}{}
		}
		if _, selected := selectedColumns["metric_scope"]; !selected {
			selectColumns = append(selectColumns, "base.\"metric_scope\"")
			selectedColumns["metric_scope"] = struct{}{}
		}
		if _, selected := selectedColumns["scope_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"scope_ref\"")
			selectedColumns["scope_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["period_start"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_start\"")
			selectedColumns["period_start"] = struct{}{}
		}
		if _, selected := selectedColumns["period_end"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_end\"")
			selectedColumns["period_end"] = struct{}{}
		}
		if _, selected := selectedColumns["metric_value"]; !selected {
			selectColumns = append(selectColumns, "base.\"metric_value\"")
			selectedColumns["metric_value"] = struct{}{}
		}
		if _, selected := selectedColumns["metric_unit"]; !selected {
			selectColumns = append(selectColumns, "base.\"metric_unit\"")
			selectedColumns["metric_unit"] = struct{}{}
		}
		if _, selected := selectedColumns["dimensions"]; !selected {
			selectColumns = append(selectColumns, "base.\"dimensions\"")
			selectedColumns["dimensions"] = struct{}{}
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

type financeMetricSnapshotsFilterPlaceholder struct {
	index int
}

func (p *financeMetricSnapshotsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFinanceMetricSnapshotsFilterPredicate(filterField model.FilterField, placeholders *financeMetricSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFinanceMetricSnapshotsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFinanceMetricSnapshotsFilterSQLExpr(spec)
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

func composeFinanceMetricSnapshotsFilterGroup(group model.FilterGroup, placeholders *financeMetricSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFinanceMetricSnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFinanceMetricSnapshotsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFinanceMetricSnapshotsFilterWhereQueries(filter model.Filter, placeholders *financeMetricSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFinanceMetricSnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFinanceMetricSnapshotsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFinanceMetricSnapshotsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFinanceMetricSnapshotsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFinanceMetricSnapshotsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFinanceMetricSnapshotsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := financeMetricSnapshotsFilterPlaceholder{index: 1}
	whereQueries, err := composeFinanceMetricSnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFinanceMetricSnapshotsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFinanceMetricSnapshotsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFinanceMetricSnapshotsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"finance_metric_snapshots\" base%s", strings.Join(selectColumns, ","), composeFinanceMetricSnapshotsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFinanceMetricSnapshotsByID(ctx context.Context, primaryID model.FinanceMetricSnapshotsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFinanceMetricSnapshotsCompositePrimaryKeyWhere([]model.FinanceMetricSnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", financeMetricSnapshotsQueries.selectCountFinanceMetricSnapshots, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFinanceMetricSnapshotsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceMetricSnapshots(ctx context.Context, selectFields ...FinanceMetricSnapshotsField) (financeMetricSnapshotsList model.FinanceMetricSnapshotsList, err error) {
	var (
		defaultFinanceMetricSnapshotsSelectFields = defaultFinanceMetricSnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceMetricSnapshotsSelectFields = composeFinanceMetricSnapshotsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(financeMetricSnapshotsQueries.selectFinanceMetricSnapshots, defaultFinanceMetricSnapshotsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &financeMetricSnapshotsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFinanceMetricSnapshots] failed get financeMetricSnapshots list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFinanceMetricSnapshotsByID(ctx context.Context, primaryID model.FinanceMetricSnapshotsPrimaryID, selectFields ...FinanceMetricSnapshotsField) (financeMetricSnapshots model.FinanceMetricSnapshots, err error) {
	var (
		defaultFinanceMetricSnapshotsSelectFields = defaultFinanceMetricSnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFinanceMetricSnapshotsSelectFields = composeFinanceMetricSnapshotsSelectFields(selectFields...)
	}
	whereQry, params := composeFinanceMetricSnapshotsCompositePrimaryKeyWhere([]model.FinanceMetricSnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf(financeMetricSnapshotsQueries.selectFinanceMetricSnapshots+" WHERE "+whereQry, defaultFinanceMetricSnapshotsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &financeMetricSnapshots, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("financeMetricSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFinanceMetricSnapshotsByID] failed get financeMetricSnapshots")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFinanceMetricSnapshotsByID(ctx context.Context, primaryID model.FinanceMetricSnapshotsPrimaryID, financeMetricSnapshots *model.FinanceMetricSnapshots, financeMetricSnapshotsUpdateFields ...FinanceMetricSnapshotsUpdateField) (err error) {
	exists, err := repo.IsExistFinanceMetricSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceMetricSnapshots] failed checking financeMetricSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("financeMetricSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if financeMetricSnapshots == nil {
		if len(financeMetricSnapshotsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFinanceMetricSnapshotsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		financeMetricSnapshots = &model.FinanceMetricSnapshots{}
	}
	var (
		defaultFinanceMetricSnapshotsUpdateFields = defaultFinanceMetricSnapshotsUpdateFields(*financeMetricSnapshots)
		tempUpdateField                           FinanceMetricSnapshotsUpdateFieldList
		selectFields                              = NewFinanceMetricSnapshotsSelectFields()
	)
	if len(financeMetricSnapshotsUpdateFields) > 0 {
		for _, updateField := range financeMetricSnapshotsUpdateFields {
			if updateField.financeMetricSnapshotsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFinanceMetricSnapshotsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFinanceMetricSnapshotsCompositePrimaryKeyWhere([]model.FinanceMetricSnapshotsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFinanceMetricSnapshotsCommand(defaultFinanceMetricSnapshotsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(financeMetricSnapshotsQueries.updateFinanceMetricSnapshots+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceMetricSnapshots] error when try to update financeMetricSnapshots by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFinanceMetricSnapshotsByFilter(ctx context.Context, filter model.Filter, financeMetricSnapshotsUpdateFields ...FinanceMetricSnapshotsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(financeMetricSnapshotsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FinanceMetricSnapshotsUpdateFieldList
		selectFields = NewFinanceMetricSnapshotsSelectFields()
	)
	for _, updateField := range financeMetricSnapshotsUpdateFields {
		if updateField.financeMetricSnapshotsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFinanceMetricSnapshotsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := financeMetricSnapshotsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFinanceMetricSnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"finance_metric_snapshots\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceMetricSnapshotsByFilter] error when try to update financeMetricSnapshots by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFinanceMetricSnapshotsByFilter] failed get rows affected")
	}
	return
}

var (
	financeMetricSnapshotsQueries = struct {
		selectFinanceMetricSnapshots      string
		selectCountFinanceMetricSnapshots string
		deleteFinanceMetricSnapshots      string
		updateFinanceMetricSnapshots      string
		insertFinanceMetricSnapshots      string
	}{
		selectFinanceMetricSnapshots:      "SELECT %s FROM \"finance_metric_snapshots\"",
		selectCountFinanceMetricSnapshots: "SELECT COUNT(\"id\") FROM \"finance_metric_snapshots\"",
		deleteFinanceMetricSnapshots:      "DELETE FROM \"finance_metric_snapshots\"",
		updateFinanceMetricSnapshots:      "UPDATE \"finance_metric_snapshots\" SET %s ",
		insertFinanceMetricSnapshots:      "INSERT INTO \"finance_metric_snapshots\" %s VALUES %s",
	}
)

type FinanceMetricSnapshotsRepository interface {
	CreateFinanceMetricSnapshots(ctx context.Context, financeMetricSnapshots *model.FinanceMetricSnapshots, fieldsInsert ...FinanceMetricSnapshotsField) error
	BulkCreateFinanceMetricSnapshots(ctx context.Context, financeMetricSnapshotsList []*model.FinanceMetricSnapshots, fieldsInsert ...FinanceMetricSnapshotsField) error
	ResolveFinanceMetricSnapshots(ctx context.Context, selectFields ...FinanceMetricSnapshotsField) (model.FinanceMetricSnapshotsList, error)
	ResolveFinanceMetricSnapshotsByID(ctx context.Context, primaryID model.FinanceMetricSnapshotsPrimaryID, selectFields ...FinanceMetricSnapshotsField) (model.FinanceMetricSnapshots, error)
	UpdateFinanceMetricSnapshotsByID(ctx context.Context, id model.FinanceMetricSnapshotsPrimaryID, financeMetricSnapshots *model.FinanceMetricSnapshots, financeMetricSnapshotsUpdateFields ...FinanceMetricSnapshotsUpdateField) error
	UpdateFinanceMetricSnapshotsByFilter(ctx context.Context, filter model.Filter, financeMetricSnapshotsUpdateFields ...FinanceMetricSnapshotsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFinanceMetricSnapshots(ctx context.Context, financeMetricSnapshotsListMap map[model.FinanceMetricSnapshotsPrimaryID]*model.FinanceMetricSnapshots, FinanceMetricSnapshotssMapUpdateFieldsRequest map[model.FinanceMetricSnapshotsPrimaryID]FinanceMetricSnapshotsUpdateFieldList) (err error)
	DeleteFinanceMetricSnapshotsByID(ctx context.Context, id model.FinanceMetricSnapshotsPrimaryID) error
	BulkDeleteFinanceMetricSnapshotsByIDs(ctx context.Context, ids []model.FinanceMetricSnapshotsPrimaryID) error
	ResolveFinanceMetricSnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.FinanceMetricSnapshotsFilterResult, err error)
	IsExistFinanceMetricSnapshotsByIDs(ctx context.Context, ids []model.FinanceMetricSnapshotsPrimaryID) (exists bool, notFoundIds []model.FinanceMetricSnapshotsPrimaryID, err error)
	IsExistFinanceMetricSnapshotsByID(ctx context.Context, id model.FinanceMetricSnapshotsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
