package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsProviderHealthSnapshots(providerHealthSnapshotsList []model.ProviderHealthSnapshots, fieldsInsert ...ProviderHealthSnapshotsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewProviderHealthSnapshotsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, providerHealthSnapshots := range providerHealthSnapshotsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, providerHealthSnapshots.Id)
			case selectField.ProviderAccountId():
				args = append(args, providerHealthSnapshots.ProviderAccountId)
			case selectField.MethodCode():
				args = append(args, providerHealthSnapshots.MethodCode)
			case selectField.ChannelCode():
				args = append(args, providerHealthSnapshots.ChannelCode)
			case selectField.HealthScore():
				args = append(args, providerHealthSnapshots.HealthScore)
			case selectField.SuccessRate():
				args = append(args, providerHealthSnapshots.SuccessRate)
			case selectField.TimeoutRate():
				args = append(args, providerHealthSnapshots.TimeoutRate)
			case selectField.ErrorRate():
				args = append(args, providerHealthSnapshots.ErrorRate)
			case selectField.P95LatencyMs():
				args = append(args, providerHealthSnapshots.P95LatencyMs)
			case selectField.SampleSize():
				args = append(args, providerHealthSnapshots.SampleSize)
			case selectField.WindowStartedAt():
				args = append(args, providerHealthSnapshots.WindowStartedAt)
			case selectField.WindowEndedAt():
				args = append(args, providerHealthSnapshots.WindowEndedAt)
			case selectField.Metadata():
				args = append(args, providerHealthSnapshots.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, providerHealthSnapshots.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, providerHealthSnapshots.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, providerHealthSnapshots.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, providerHealthSnapshots.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, providerHealthSnapshots.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, providerHealthSnapshots.MetaDeletedBy)

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

func composeProviderHealthSnapshotsCompositePrimaryKeyWhere(primaryIDs []model.ProviderHealthSnapshotsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"provider_health_snapshots\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultProviderHealthSnapshotsSelectFields() string {
	fields := NewProviderHealthSnapshotsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeProviderHealthSnapshotsSelectFields(selectFields ...ProviderHealthSnapshotsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ProviderHealthSnapshotsField string
type ProviderHealthSnapshotsFieldList []ProviderHealthSnapshotsField

type ProviderHealthSnapshotsSelectFields struct {
}

func (ss ProviderHealthSnapshotsSelectFields) Id() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("id")
}

func (ss ProviderHealthSnapshotsSelectFields) ProviderAccountId() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("provider_account_id")
}

func (ss ProviderHealthSnapshotsSelectFields) MethodCode() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("method_code")
}

func (ss ProviderHealthSnapshotsSelectFields) ChannelCode() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("channel_code")
}

func (ss ProviderHealthSnapshotsSelectFields) HealthScore() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("health_score")
}

func (ss ProviderHealthSnapshotsSelectFields) SuccessRate() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("success_rate")
}

func (ss ProviderHealthSnapshotsSelectFields) TimeoutRate() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("timeout_rate")
}

func (ss ProviderHealthSnapshotsSelectFields) ErrorRate() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("error_rate")
}

func (ss ProviderHealthSnapshotsSelectFields) P95LatencyMs() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("p95_latency_ms")
}

func (ss ProviderHealthSnapshotsSelectFields) SampleSize() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("sample_size")
}

func (ss ProviderHealthSnapshotsSelectFields) WindowStartedAt() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("window_started_at")
}

func (ss ProviderHealthSnapshotsSelectFields) WindowEndedAt() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("window_ended_at")
}

func (ss ProviderHealthSnapshotsSelectFields) Metadata() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("metadata")
}

func (ss ProviderHealthSnapshotsSelectFields) MetaCreatedAt() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("meta_created_at")
}

func (ss ProviderHealthSnapshotsSelectFields) MetaCreatedBy() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("meta_created_by")
}

func (ss ProviderHealthSnapshotsSelectFields) MetaUpdatedAt() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("meta_updated_at")
}

func (ss ProviderHealthSnapshotsSelectFields) MetaUpdatedBy() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("meta_updated_by")
}

func (ss ProviderHealthSnapshotsSelectFields) MetaDeletedAt() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("meta_deleted_at")
}

func (ss ProviderHealthSnapshotsSelectFields) MetaDeletedBy() ProviderHealthSnapshotsField {
	return ProviderHealthSnapshotsField("meta_deleted_by")
}

func (ss ProviderHealthSnapshotsSelectFields) All() ProviderHealthSnapshotsFieldList {
	return []ProviderHealthSnapshotsField{
		ss.Id(),
		ss.ProviderAccountId(),
		ss.MethodCode(),
		ss.ChannelCode(),
		ss.HealthScore(),
		ss.SuccessRate(),
		ss.TimeoutRate(),
		ss.ErrorRate(),
		ss.P95LatencyMs(),
		ss.SampleSize(),
		ss.WindowStartedAt(),
		ss.WindowEndedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewProviderHealthSnapshotsSelectFields() ProviderHealthSnapshotsSelectFields {
	return ProviderHealthSnapshotsSelectFields{}
}

type ProviderHealthSnapshotsUpdateFieldOption struct {
	useIncrement bool
}
type ProviderHealthSnapshotsUpdateField struct {
	providerHealthSnapshotsField ProviderHealthSnapshotsField
	opt                          ProviderHealthSnapshotsUpdateFieldOption
	value                        interface{}
}
type ProviderHealthSnapshotsUpdateFieldList []ProviderHealthSnapshotsUpdateField

func defaultProviderHealthSnapshotsUpdateFieldOption() ProviderHealthSnapshotsUpdateFieldOption {
	return ProviderHealthSnapshotsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementProviderHealthSnapshotsOption(useIncrement bool) func(*ProviderHealthSnapshotsUpdateFieldOption) {
	return func(pcufo *ProviderHealthSnapshotsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewProviderHealthSnapshotsUpdateField(field ProviderHealthSnapshotsField, val interface{}, opts ...func(*ProviderHealthSnapshotsUpdateFieldOption)) ProviderHealthSnapshotsUpdateField {
	defaultOpt := defaultProviderHealthSnapshotsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ProviderHealthSnapshotsUpdateField{
		providerHealthSnapshotsField: field,
		value:                        val,
		opt:                          defaultOpt,
	}
}
func defaultProviderHealthSnapshotsUpdateFields(providerHealthSnapshots model.ProviderHealthSnapshots) (providerHealthSnapshotsUpdateFieldList ProviderHealthSnapshotsUpdateFieldList) {
	selectFields := NewProviderHealthSnapshotsSelectFields()
	providerHealthSnapshotsUpdateFieldList = append(providerHealthSnapshotsUpdateFieldList,
		NewProviderHealthSnapshotsUpdateField(selectFields.Id(), providerHealthSnapshots.Id),
		NewProviderHealthSnapshotsUpdateField(selectFields.ProviderAccountId(), providerHealthSnapshots.ProviderAccountId),
		NewProviderHealthSnapshotsUpdateField(selectFields.MethodCode(), providerHealthSnapshots.MethodCode),
		NewProviderHealthSnapshotsUpdateField(selectFields.ChannelCode(), providerHealthSnapshots.ChannelCode),
		NewProviderHealthSnapshotsUpdateField(selectFields.HealthScore(), providerHealthSnapshots.HealthScore),
		NewProviderHealthSnapshotsUpdateField(selectFields.SuccessRate(), providerHealthSnapshots.SuccessRate),
		NewProviderHealthSnapshotsUpdateField(selectFields.TimeoutRate(), providerHealthSnapshots.TimeoutRate),
		NewProviderHealthSnapshotsUpdateField(selectFields.ErrorRate(), providerHealthSnapshots.ErrorRate),
		NewProviderHealthSnapshotsUpdateField(selectFields.P95LatencyMs(), providerHealthSnapshots.P95LatencyMs),
		NewProviderHealthSnapshotsUpdateField(selectFields.SampleSize(), providerHealthSnapshots.SampleSize),
		NewProviderHealthSnapshotsUpdateField(selectFields.WindowStartedAt(), providerHealthSnapshots.WindowStartedAt),
		NewProviderHealthSnapshotsUpdateField(selectFields.WindowEndedAt(), providerHealthSnapshots.WindowEndedAt),
		NewProviderHealthSnapshotsUpdateField(selectFields.Metadata(), providerHealthSnapshots.Metadata),
		NewProviderHealthSnapshotsUpdateField(selectFields.MetaCreatedAt(), providerHealthSnapshots.MetaCreatedAt),
		NewProviderHealthSnapshotsUpdateField(selectFields.MetaCreatedBy(), providerHealthSnapshots.MetaCreatedBy),
		NewProviderHealthSnapshotsUpdateField(selectFields.MetaUpdatedAt(), providerHealthSnapshots.MetaUpdatedAt),
		NewProviderHealthSnapshotsUpdateField(selectFields.MetaUpdatedBy(), providerHealthSnapshots.MetaUpdatedBy),
		NewProviderHealthSnapshotsUpdateField(selectFields.MetaDeletedAt(), providerHealthSnapshots.MetaDeletedAt),
		NewProviderHealthSnapshotsUpdateField(selectFields.MetaDeletedBy(), providerHealthSnapshots.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsProviderHealthSnapshotsCommand(providerHealthSnapshotsUpdateFieldList ProviderHealthSnapshotsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range providerHealthSnapshotsUpdateFieldList {
		field := string(updateField.providerHealthSnapshotsField)
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

func (repo *RepositoryImpl) BulkCreateProviderHealthSnapshots(ctx context.Context, providerHealthSnapshotsList []*model.ProviderHealthSnapshots, fieldsInsert ...ProviderHealthSnapshotsField) (err error) {
	var (
		fieldsStr                        string
		valueListStr                     []string
		argsList                         []interface{}
		primaryIds                       []model.ProviderHealthSnapshotsPrimaryID
		providerHealthSnapshotsValueList []model.ProviderHealthSnapshots
	)

	if len(fieldsInsert) == 0 {
		selectField := NewProviderHealthSnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, providerHealthSnapshots := range providerHealthSnapshotsList {

		primaryIds = append(primaryIds, providerHealthSnapshots.ToProviderHealthSnapshotsPrimaryID())

		providerHealthSnapshotsValueList = append(providerHealthSnapshotsValueList, *providerHealthSnapshots)
	}

	_, notFoundIds, err := repo.IsExistProviderHealthSnapshotsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderHealthSnapshots] failed checking providerHealthSnapshots whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ProviderHealthSnapshotsPrimaryID{}
		mapNotFoundIds := map[model.ProviderHealthSnapshotsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "providerHealthSnapshots", fmt.Sprintf("providerHealthSnapshots with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsProviderHealthSnapshots(providerHealthSnapshotsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(providerHealthSnapshotsQueries.insertProviderHealthSnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateProviderHealthSnapshots] failed exec create providerHealthSnapshots query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteProviderHealthSnapshotsByIDs(ctx context.Context, primaryIDs []model.ProviderHealthSnapshotsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistProviderHealthSnapshotsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderHealthSnapshotsByIDs] failed checking providerHealthSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerHealthSnapshots with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_health_snapshots\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := providerHealthSnapshotsQueries.deleteProviderHealthSnapshots + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderHealthSnapshotsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteProviderHealthSnapshotsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistProviderHealthSnapshotsByIDs(ctx context.Context, ids []model.ProviderHealthSnapshotsPrimaryID) (exists bool, notFoundIds []model.ProviderHealthSnapshotsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"provider_health_snapshots\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(providerHealthSnapshotsQueries.selectProviderHealthSnapshots, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderHealthSnapshotsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ProviderHealthSnapshotsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderHealthSnapshotsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ProviderHealthSnapshotsPrimaryID]bool{}
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

// BulkUpdateProviderHealthSnapshots is used to bulk update providerHealthSnapshots, by default it will update all field
// if want to update specific field, then fill providerHealthSnapshotssMapUpdateFieldsRequest else please fill providerHealthSnapshotssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateProviderHealthSnapshots(ctx context.Context, providerHealthSnapshotssMap map[model.ProviderHealthSnapshotsPrimaryID]*model.ProviderHealthSnapshots, providerHealthSnapshotssMapUpdateFieldsRequest map[model.ProviderHealthSnapshotsPrimaryID]ProviderHealthSnapshotsUpdateFieldList) (err error) {
	if len(providerHealthSnapshotssMap) == 0 && len(providerHealthSnapshotssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		providerHealthSnapshotssMapUpdateField map[model.ProviderHealthSnapshotsPrimaryID]ProviderHealthSnapshotsUpdateFieldList = map[model.ProviderHealthSnapshotsPrimaryID]ProviderHealthSnapshotsUpdateFieldList{}
		asTableValues                          string                                                                            = "myvalues"
	)

	if len(providerHealthSnapshotssMap) > 0 {
		for id, providerHealthSnapshots := range providerHealthSnapshotssMap {
			if providerHealthSnapshots == nil {
				log.Error().Err(err).Msg("[BulkUpdateProviderHealthSnapshots] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			providerHealthSnapshotssMapUpdateField[id] = defaultProviderHealthSnapshotsUpdateFields(*providerHealthSnapshots)
		}
	} else {
		providerHealthSnapshotssMapUpdateField = providerHealthSnapshotssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateProviderHealthSnapshotsQuery(providerHealthSnapshotssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistProviderHealthSnapshotsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderHealthSnapshots] failed checking providerHealthSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerHealthSnapshots with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeProviderHealthSnapshotsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"provider_health_snapshots\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateProviderHealthSnapshots] failed exec query")
	}
	return
}

type ProviderHealthSnapshotsFieldParameter struct {
	param string
	args  []interface{}
}

func NewProviderHealthSnapshotsFieldParameter(param string, args ...interface{}) ProviderHealthSnapshotsFieldParameter {
	return ProviderHealthSnapshotsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateProviderHealthSnapshotsQuery(mapProviderHealthSnapshotss map[model.ProviderHealthSnapshotsPrimaryID]ProviderHealthSnapshotsUpdateFieldList, asTableValues string) (primaryIDs []model.ProviderHealthSnapshotsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ProviderHealthSnapshotsPrimaryID]map[string]interface{}{}
	providerHealthSnapshotsSelectFields := NewProviderHealthSnapshotsSelectFields()
	for id, updateFields := range mapProviderHealthSnapshotss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.providerHealthSnapshotsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapProviderHealthSnapshotss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetProviderHealthSnapshotsFieldType(updateField.providerHealthSnapshotsField)))
			args = append(args, fields[string(updateField.providerHealthSnapshotsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.providerHealthSnapshotsField))
		if updateField.providerHealthSnapshotsField == providerHealthSnapshotsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.providerHealthSnapshotsField, asTableValues, updateField.providerHealthSnapshotsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.providerHealthSnapshotsField,
				"\"provider_health_snapshots\"", updateField.providerHealthSnapshotsField,
				asTableValues, updateField.providerHealthSnapshotsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeProviderHealthSnapshotsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ProviderHealthSnapshotsPrimaryID, asTableValue string) (whereQry string) {
	providerHealthSnapshotsSelectFields := NewProviderHealthSnapshotsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"provider_health_snapshots\".\"id\" = %s.\"id\"::"+GetProviderHealthSnapshotsFieldType(providerHealthSnapshotsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetProviderHealthSnapshotsFieldType(providerHealthSnapshotsField ProviderHealthSnapshotsField) string {
	selectProviderHealthSnapshotsFields := NewProviderHealthSnapshotsSelectFields()
	switch providerHealthSnapshotsField {

	case selectProviderHealthSnapshotsFields.Id():
		return "uuid"

	case selectProviderHealthSnapshotsFields.ProviderAccountId():
		return "uuid"

	case selectProviderHealthSnapshotsFields.MethodCode():
		return "text"

	case selectProviderHealthSnapshotsFields.ChannelCode():
		return "text"

	case selectProviderHealthSnapshotsFields.HealthScore():
		return "int4"

	case selectProviderHealthSnapshotsFields.SuccessRate():
		return "numeric"

	case selectProviderHealthSnapshotsFields.TimeoutRate():
		return "numeric"

	case selectProviderHealthSnapshotsFields.ErrorRate():
		return "numeric"

	case selectProviderHealthSnapshotsFields.P95LatencyMs():
		return "int4"

	case selectProviderHealthSnapshotsFields.SampleSize():
		return "int4"

	case selectProviderHealthSnapshotsFields.WindowStartedAt():
		return "timestamptz"

	case selectProviderHealthSnapshotsFields.WindowEndedAt():
		return "timestamptz"

	case selectProviderHealthSnapshotsFields.Metadata():
		return "jsonb"

	case selectProviderHealthSnapshotsFields.MetaCreatedAt():
		return "timestamptz"

	case selectProviderHealthSnapshotsFields.MetaCreatedBy():
		return "uuid"

	case selectProviderHealthSnapshotsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectProviderHealthSnapshotsFields.MetaUpdatedBy():
		return "uuid"

	case selectProviderHealthSnapshotsFields.MetaDeletedAt():
		return "timestamptz"

	case selectProviderHealthSnapshotsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateProviderHealthSnapshots(ctx context.Context, providerHealthSnapshots *model.ProviderHealthSnapshots, fieldsInsert ...ProviderHealthSnapshotsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewProviderHealthSnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ProviderHealthSnapshotsPrimaryID{
		Id: providerHealthSnapshots.Id,
	}
	exists, err := repo.IsExistProviderHealthSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderHealthSnapshots] failed checking providerHealthSnapshots whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "providerHealthSnapshots", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsProviderHealthSnapshots([]model.ProviderHealthSnapshots{*providerHealthSnapshots}, fieldsInsert...)
	commandQuery := fmt.Sprintf(providerHealthSnapshotsQueries.insertProviderHealthSnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProviderHealthSnapshots] failed exec create providerHealthSnapshots query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteProviderHealthSnapshotsByID(ctx context.Context, primaryID model.ProviderHealthSnapshotsPrimaryID) (err error) {
	exists, err := repo.IsExistProviderHealthSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderHealthSnapshotsByID] failed checking providerHealthSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerHealthSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeProviderHealthSnapshotsCompositePrimaryKeyWhere([]model.ProviderHealthSnapshotsPrimaryID{primaryID})
	commandQuery := providerHealthSnapshotsQueries.deleteProviderHealthSnapshots + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteProviderHealthSnapshotsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderHealthSnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderHealthSnapshotsFilterResult, err error) {
	query, args, err := composeProviderHealthSnapshotsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderHealthSnapshotsByFilter] failed compose providerHealthSnapshots filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderHealthSnapshotsByFilter] failed get providerHealthSnapshots by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeProviderHealthSnapshotsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ProviderHealthSnapshotsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeProviderHealthSnapshotsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeProviderHealthSnapshotsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeProviderHealthSnapshotsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewProviderHealthSnapshotsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["method_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"method_code\"")
			selectedColumns["method_code"] = struct{}{}
		}
		if _, selected := selectedColumns["channel_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"channel_code\"")
			selectedColumns["channel_code"] = struct{}{}
		}
		if _, selected := selectedColumns["health_score"]; !selected {
			selectColumns = append(selectColumns, "base.\"health_score\"")
			selectedColumns["health_score"] = struct{}{}
		}
		if _, selected := selectedColumns["success_rate"]; !selected {
			selectColumns = append(selectColumns, "base.\"success_rate\"")
			selectedColumns["success_rate"] = struct{}{}
		}
		if _, selected := selectedColumns["timeout_rate"]; !selected {
			selectColumns = append(selectColumns, "base.\"timeout_rate\"")
			selectedColumns["timeout_rate"] = struct{}{}
		}
		if _, selected := selectedColumns["error_rate"]; !selected {
			selectColumns = append(selectColumns, "base.\"error_rate\"")
			selectedColumns["error_rate"] = struct{}{}
		}
		if _, selected := selectedColumns["p95_latency_ms"]; !selected {
			selectColumns = append(selectColumns, "base.\"p95_latency_ms\"")
			selectedColumns["p95_latency_ms"] = struct{}{}
		}
		if _, selected := selectedColumns["sample_size"]; !selected {
			selectColumns = append(selectColumns, "base.\"sample_size\"")
			selectedColumns["sample_size"] = struct{}{}
		}
		if _, selected := selectedColumns["window_started_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"window_started_at\"")
			selectedColumns["window_started_at"] = struct{}{}
		}
		if _, selected := selectedColumns["window_ended_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"window_ended_at\"")
			selectedColumns["window_ended_at"] = struct{}{}
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

type providerHealthSnapshotsFilterPlaceholder struct {
	index int
}

func (p *providerHealthSnapshotsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeProviderHealthSnapshotsFilterPredicate(filterField model.FilterField, placeholders *providerHealthSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewProviderHealthSnapshotsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeProviderHealthSnapshotsFilterSQLExpr(spec)
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

func composeProviderHealthSnapshotsFilterGroup(group model.FilterGroup, placeholders *providerHealthSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeProviderHealthSnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeProviderHealthSnapshotsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeProviderHealthSnapshotsFilterWhereQueries(filter model.Filter, placeholders *providerHealthSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeProviderHealthSnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeProviderHealthSnapshotsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeProviderHealthSnapshotsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateProviderHealthSnapshotsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeProviderHealthSnapshotsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeProviderHealthSnapshotsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := providerHealthSnapshotsFilterPlaceholder{index: 1}
	whereQueries, err := composeProviderHealthSnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewProviderHealthSnapshotsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeProviderHealthSnapshotsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeProviderHealthSnapshotsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"provider_health_snapshots\" base%s", strings.Join(selectColumns, ","), composeProviderHealthSnapshotsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistProviderHealthSnapshotsByID(ctx context.Context, primaryID model.ProviderHealthSnapshotsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeProviderHealthSnapshotsCompositePrimaryKeyWhere([]model.ProviderHealthSnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", providerHealthSnapshotsQueries.selectCountProviderHealthSnapshots, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProviderHealthSnapshotsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderHealthSnapshots(ctx context.Context, selectFields ...ProviderHealthSnapshotsField) (providerHealthSnapshotsList model.ProviderHealthSnapshotsList, err error) {
	var (
		defaultProviderHealthSnapshotsSelectFields = defaultProviderHealthSnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderHealthSnapshotsSelectFields = composeProviderHealthSnapshotsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(providerHealthSnapshotsQueries.selectProviderHealthSnapshots, defaultProviderHealthSnapshotsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &providerHealthSnapshotsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveProviderHealthSnapshots] failed get providerHealthSnapshots list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveProviderHealthSnapshotsByID(ctx context.Context, primaryID model.ProviderHealthSnapshotsPrimaryID, selectFields ...ProviderHealthSnapshotsField) (providerHealthSnapshots model.ProviderHealthSnapshots, err error) {
	var (
		defaultProviderHealthSnapshotsSelectFields = defaultProviderHealthSnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultProviderHealthSnapshotsSelectFields = composeProviderHealthSnapshotsSelectFields(selectFields...)
	}
	whereQry, params := composeProviderHealthSnapshotsCompositePrimaryKeyWhere([]model.ProviderHealthSnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf(providerHealthSnapshotsQueries.selectProviderHealthSnapshots+" WHERE "+whereQry, defaultProviderHealthSnapshotsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &providerHealthSnapshots, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("providerHealthSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveProviderHealthSnapshotsByID] failed get providerHealthSnapshots")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateProviderHealthSnapshotsByID(ctx context.Context, primaryID model.ProviderHealthSnapshotsPrimaryID, providerHealthSnapshots *model.ProviderHealthSnapshots, providerHealthSnapshotsUpdateFields ...ProviderHealthSnapshotsUpdateField) (err error) {
	exists, err := repo.IsExistProviderHealthSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderHealthSnapshots] failed checking providerHealthSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("providerHealthSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if providerHealthSnapshots == nil {
		if len(providerHealthSnapshotsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateProviderHealthSnapshotsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		providerHealthSnapshots = &model.ProviderHealthSnapshots{}
	}
	var (
		defaultProviderHealthSnapshotsUpdateFields = defaultProviderHealthSnapshotsUpdateFields(*providerHealthSnapshots)
		tempUpdateField                            ProviderHealthSnapshotsUpdateFieldList
		selectFields                               = NewProviderHealthSnapshotsSelectFields()
	)
	if len(providerHealthSnapshotsUpdateFields) > 0 {
		for _, updateField := range providerHealthSnapshotsUpdateFields {
			if updateField.providerHealthSnapshotsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultProviderHealthSnapshotsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeProviderHealthSnapshotsCompositePrimaryKeyWhere([]model.ProviderHealthSnapshotsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsProviderHealthSnapshotsCommand(defaultProviderHealthSnapshotsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(providerHealthSnapshotsQueries.updateProviderHealthSnapshots+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderHealthSnapshots] error when try to update providerHealthSnapshots by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateProviderHealthSnapshotsByFilter(ctx context.Context, filter model.Filter, providerHealthSnapshotsUpdateFields ...ProviderHealthSnapshotsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(providerHealthSnapshotsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ProviderHealthSnapshotsUpdateFieldList
		selectFields = NewProviderHealthSnapshotsSelectFields()
	)
	for _, updateField := range providerHealthSnapshotsUpdateFields {
		if updateField.providerHealthSnapshotsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsProviderHealthSnapshotsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := providerHealthSnapshotsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeProviderHealthSnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"provider_health_snapshots\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderHealthSnapshotsByFilter] error when try to update providerHealthSnapshots by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateProviderHealthSnapshotsByFilter] failed get rows affected")
	}
	return
}

var (
	providerHealthSnapshotsQueries = struct {
		selectProviderHealthSnapshots      string
		selectCountProviderHealthSnapshots string
		deleteProviderHealthSnapshots      string
		updateProviderHealthSnapshots      string
		insertProviderHealthSnapshots      string
	}{
		selectProviderHealthSnapshots:      "SELECT %s FROM \"provider_health_snapshots\"",
		selectCountProviderHealthSnapshots: "SELECT COUNT(\"id\") FROM \"provider_health_snapshots\"",
		deleteProviderHealthSnapshots:      "DELETE FROM \"provider_health_snapshots\"",
		updateProviderHealthSnapshots:      "UPDATE \"provider_health_snapshots\" SET %s ",
		insertProviderHealthSnapshots:      "INSERT INTO \"provider_health_snapshots\" %s VALUES %s",
	}
)

type ProviderHealthSnapshotsRepository interface {
	CreateProviderHealthSnapshots(ctx context.Context, providerHealthSnapshots *model.ProviderHealthSnapshots, fieldsInsert ...ProviderHealthSnapshotsField) error
	BulkCreateProviderHealthSnapshots(ctx context.Context, providerHealthSnapshotsList []*model.ProviderHealthSnapshots, fieldsInsert ...ProviderHealthSnapshotsField) error
	ResolveProviderHealthSnapshots(ctx context.Context, selectFields ...ProviderHealthSnapshotsField) (model.ProviderHealthSnapshotsList, error)
	ResolveProviderHealthSnapshotsByID(ctx context.Context, primaryID model.ProviderHealthSnapshotsPrimaryID, selectFields ...ProviderHealthSnapshotsField) (model.ProviderHealthSnapshots, error)
	UpdateProviderHealthSnapshotsByID(ctx context.Context, id model.ProviderHealthSnapshotsPrimaryID, providerHealthSnapshots *model.ProviderHealthSnapshots, providerHealthSnapshotsUpdateFields ...ProviderHealthSnapshotsUpdateField) error
	UpdateProviderHealthSnapshotsByFilter(ctx context.Context, filter model.Filter, providerHealthSnapshotsUpdateFields ...ProviderHealthSnapshotsUpdateField) (rowsAffected int64, err error)
	BulkUpdateProviderHealthSnapshots(ctx context.Context, providerHealthSnapshotsListMap map[model.ProviderHealthSnapshotsPrimaryID]*model.ProviderHealthSnapshots, ProviderHealthSnapshotssMapUpdateFieldsRequest map[model.ProviderHealthSnapshotsPrimaryID]ProviderHealthSnapshotsUpdateFieldList) (err error)
	DeleteProviderHealthSnapshotsByID(ctx context.Context, id model.ProviderHealthSnapshotsPrimaryID) error
	BulkDeleteProviderHealthSnapshotsByIDs(ctx context.Context, ids []model.ProviderHealthSnapshotsPrimaryID) error
	ResolveProviderHealthSnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.ProviderHealthSnapshotsFilterResult, err error)
	IsExistProviderHealthSnapshotsByIDs(ctx context.Context, ids []model.ProviderHealthSnapshotsPrimaryID) (exists bool, notFoundIds []model.ProviderHealthSnapshotsPrimaryID, err error)
	IsExistProviderHealthSnapshotsByID(ctx context.Context, id model.ProviderHealthSnapshotsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
