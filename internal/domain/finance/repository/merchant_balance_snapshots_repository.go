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

func composeInsertFieldsAndParamsMerchantBalanceSnapshots(merchantBalanceSnapshotsList []model.MerchantBalanceSnapshots, fieldsInsert ...MerchantBalanceSnapshotsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewMerchantBalanceSnapshotsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, merchantBalanceSnapshots := range merchantBalanceSnapshotsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, merchantBalanceSnapshots.Id)
			case selectField.BalanceAccountId():
				args = append(args, merchantBalanceSnapshots.BalanceAccountId)
			case selectField.SnapshotAt():
				args = append(args, merchantBalanceSnapshots.SnapshotAt)
			case selectField.AvailableAmount():
				args = append(args, merchantBalanceSnapshots.AvailableAmount)
			case selectField.PendingAmount():
				args = append(args, merchantBalanceSnapshots.PendingAmount)
			case selectField.ReservedAmount():
				args = append(args, merchantBalanceSnapshots.ReservedAmount)
			case selectField.DisputedAmount():
				args = append(args, merchantBalanceSnapshots.DisputedAmount)
			case selectField.NegativeAmount():
				args = append(args, merchantBalanceSnapshots.NegativeAmount)
			case selectField.Metadata():
				args = append(args, merchantBalanceSnapshots.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, merchantBalanceSnapshots.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, merchantBalanceSnapshots.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, merchantBalanceSnapshots.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, merchantBalanceSnapshots.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, merchantBalanceSnapshots.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, merchantBalanceSnapshots.MetaDeletedBy)

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

func composeMerchantBalanceSnapshotsCompositePrimaryKeyWhere(primaryIDs []model.MerchantBalanceSnapshotsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"merchant_balance_snapshots\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultMerchantBalanceSnapshotsSelectFields() string {
	fields := NewMerchantBalanceSnapshotsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeMerchantBalanceSnapshotsSelectFields(selectFields ...MerchantBalanceSnapshotsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type MerchantBalanceSnapshotsField string
type MerchantBalanceSnapshotsFieldList []MerchantBalanceSnapshotsField

type MerchantBalanceSnapshotsSelectFields struct {
}

func (ss MerchantBalanceSnapshotsSelectFields) Id() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("id")
}

func (ss MerchantBalanceSnapshotsSelectFields) BalanceAccountId() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("balance_account_id")
}

func (ss MerchantBalanceSnapshotsSelectFields) SnapshotAt() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("snapshot_at")
}

func (ss MerchantBalanceSnapshotsSelectFields) AvailableAmount() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("available_amount")
}

func (ss MerchantBalanceSnapshotsSelectFields) PendingAmount() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("pending_amount")
}

func (ss MerchantBalanceSnapshotsSelectFields) ReservedAmount() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("reserved_amount")
}

func (ss MerchantBalanceSnapshotsSelectFields) DisputedAmount() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("disputed_amount")
}

func (ss MerchantBalanceSnapshotsSelectFields) NegativeAmount() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("negative_amount")
}

func (ss MerchantBalanceSnapshotsSelectFields) Metadata() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("metadata")
}

func (ss MerchantBalanceSnapshotsSelectFields) MetaCreatedAt() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("meta_created_at")
}

func (ss MerchantBalanceSnapshotsSelectFields) MetaCreatedBy() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("meta_created_by")
}

func (ss MerchantBalanceSnapshotsSelectFields) MetaUpdatedAt() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("meta_updated_at")
}

func (ss MerchantBalanceSnapshotsSelectFields) MetaUpdatedBy() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("meta_updated_by")
}

func (ss MerchantBalanceSnapshotsSelectFields) MetaDeletedAt() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("meta_deleted_at")
}

func (ss MerchantBalanceSnapshotsSelectFields) MetaDeletedBy() MerchantBalanceSnapshotsField {
	return MerchantBalanceSnapshotsField("meta_deleted_by")
}

func (ss MerchantBalanceSnapshotsSelectFields) All() MerchantBalanceSnapshotsFieldList {
	return []MerchantBalanceSnapshotsField{
		ss.Id(),
		ss.BalanceAccountId(),
		ss.SnapshotAt(),
		ss.AvailableAmount(),
		ss.PendingAmount(),
		ss.ReservedAmount(),
		ss.DisputedAmount(),
		ss.NegativeAmount(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewMerchantBalanceSnapshotsSelectFields() MerchantBalanceSnapshotsSelectFields {
	return MerchantBalanceSnapshotsSelectFields{}
}

type MerchantBalanceSnapshotsUpdateFieldOption struct {
	useIncrement bool
}
type MerchantBalanceSnapshotsUpdateField struct {
	merchantBalanceSnapshotsField MerchantBalanceSnapshotsField
	opt                           MerchantBalanceSnapshotsUpdateFieldOption
	value                         interface{}
}
type MerchantBalanceSnapshotsUpdateFieldList []MerchantBalanceSnapshotsUpdateField

func defaultMerchantBalanceSnapshotsUpdateFieldOption() MerchantBalanceSnapshotsUpdateFieldOption {
	return MerchantBalanceSnapshotsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementMerchantBalanceSnapshotsOption(useIncrement bool) func(*MerchantBalanceSnapshotsUpdateFieldOption) {
	return func(pcufo *MerchantBalanceSnapshotsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewMerchantBalanceSnapshotsUpdateField(field MerchantBalanceSnapshotsField, val interface{}, opts ...func(*MerchantBalanceSnapshotsUpdateFieldOption)) MerchantBalanceSnapshotsUpdateField {
	defaultOpt := defaultMerchantBalanceSnapshotsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return MerchantBalanceSnapshotsUpdateField{
		merchantBalanceSnapshotsField: field,
		value:                         val,
		opt:                           defaultOpt,
	}
}
func defaultMerchantBalanceSnapshotsUpdateFields(merchantBalanceSnapshots model.MerchantBalanceSnapshots) (merchantBalanceSnapshotsUpdateFieldList MerchantBalanceSnapshotsUpdateFieldList) {
	selectFields := NewMerchantBalanceSnapshotsSelectFields()
	merchantBalanceSnapshotsUpdateFieldList = append(merchantBalanceSnapshotsUpdateFieldList,
		NewMerchantBalanceSnapshotsUpdateField(selectFields.Id(), merchantBalanceSnapshots.Id),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.BalanceAccountId(), merchantBalanceSnapshots.BalanceAccountId),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.SnapshotAt(), merchantBalanceSnapshots.SnapshotAt),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.AvailableAmount(), merchantBalanceSnapshots.AvailableAmount),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.PendingAmount(), merchantBalanceSnapshots.PendingAmount),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.ReservedAmount(), merchantBalanceSnapshots.ReservedAmount),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.DisputedAmount(), merchantBalanceSnapshots.DisputedAmount),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.NegativeAmount(), merchantBalanceSnapshots.NegativeAmount),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.Metadata(), merchantBalanceSnapshots.Metadata),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.MetaCreatedAt(), merchantBalanceSnapshots.MetaCreatedAt),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.MetaCreatedBy(), merchantBalanceSnapshots.MetaCreatedBy),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.MetaUpdatedAt(), merchantBalanceSnapshots.MetaUpdatedAt),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.MetaUpdatedBy(), merchantBalanceSnapshots.MetaUpdatedBy),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.MetaDeletedAt(), merchantBalanceSnapshots.MetaDeletedAt),
		NewMerchantBalanceSnapshotsUpdateField(selectFields.MetaDeletedBy(), merchantBalanceSnapshots.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsMerchantBalanceSnapshotsCommand(merchantBalanceSnapshotsUpdateFieldList MerchantBalanceSnapshotsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range merchantBalanceSnapshotsUpdateFieldList {
		field := string(updateField.merchantBalanceSnapshotsField)
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

func (repo *RepositoryImpl) BulkCreateMerchantBalanceSnapshots(ctx context.Context, merchantBalanceSnapshotsList []*model.MerchantBalanceSnapshots, fieldsInsert ...MerchantBalanceSnapshotsField) (err error) {
	var (
		fieldsStr                         string
		valueListStr                      []string
		argsList                          []interface{}
		primaryIds                        []model.MerchantBalanceSnapshotsPrimaryID
		merchantBalanceSnapshotsValueList []model.MerchantBalanceSnapshots
	)

	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceSnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, merchantBalanceSnapshots := range merchantBalanceSnapshotsList {

		primaryIds = append(primaryIds, merchantBalanceSnapshots.ToMerchantBalanceSnapshotsPrimaryID())

		merchantBalanceSnapshotsValueList = append(merchantBalanceSnapshotsValueList, *merchantBalanceSnapshots)
	}

	_, notFoundIds, err := repo.IsExistMerchantBalanceSnapshotsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceSnapshots] failed checking merchantBalanceSnapshots whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.MerchantBalanceSnapshotsPrimaryID{}
		mapNotFoundIds := map[model.MerchantBalanceSnapshotsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "merchantBalanceSnapshots", fmt.Sprintf("merchantBalanceSnapshots with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsMerchantBalanceSnapshots(merchantBalanceSnapshotsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(merchantBalanceSnapshotsQueries.insertMerchantBalanceSnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceSnapshots] failed exec create merchantBalanceSnapshots query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteMerchantBalanceSnapshotsByIDs(ctx context.Context, primaryIDs []model.MerchantBalanceSnapshotsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistMerchantBalanceSnapshotsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceSnapshotsByIDs] failed checking merchantBalanceSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSnapshots with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"merchant_balance_snapshots\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(merchantBalanceSnapshotsQueries.deleteMerchantBalanceSnapshots + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceSnapshotsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceSnapshotsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistMerchantBalanceSnapshotsByIDs(ctx context.Context, ids []model.MerchantBalanceSnapshotsPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceSnapshotsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"merchant_balance_snapshots\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(merchantBalanceSnapshotsQueries.selectMerchantBalanceSnapshots, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceSnapshotsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.MerchantBalanceSnapshotsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceSnapshotsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.MerchantBalanceSnapshotsPrimaryID]bool{}
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

// BulkUpdateMerchantBalanceSnapshots is used to bulk update merchantBalanceSnapshots, by default it will update all field
// if want to update specific field, then fill merchantBalanceSnapshotssMapUpdateFieldsRequest else please fill merchantBalanceSnapshotssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateMerchantBalanceSnapshots(ctx context.Context, merchantBalanceSnapshotssMap map[model.MerchantBalanceSnapshotsPrimaryID]*model.MerchantBalanceSnapshots, merchantBalanceSnapshotssMapUpdateFieldsRequest map[model.MerchantBalanceSnapshotsPrimaryID]MerchantBalanceSnapshotsUpdateFieldList) (err error) {
	if len(merchantBalanceSnapshotssMap) == 0 && len(merchantBalanceSnapshotssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		merchantBalanceSnapshotssMapUpdateField map[model.MerchantBalanceSnapshotsPrimaryID]MerchantBalanceSnapshotsUpdateFieldList = map[model.MerchantBalanceSnapshotsPrimaryID]MerchantBalanceSnapshotsUpdateFieldList{}
		asTableValues                           string                                                                              = "myvalues"
	)

	if len(merchantBalanceSnapshotssMap) > 0 {
		for id, merchantBalanceSnapshots := range merchantBalanceSnapshotssMap {
			if merchantBalanceSnapshots == nil {
				log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceSnapshots] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			merchantBalanceSnapshotssMapUpdateField[id] = defaultMerchantBalanceSnapshotsUpdateFields(*merchantBalanceSnapshots)
		}
	} else {
		merchantBalanceSnapshotssMapUpdateField = merchantBalanceSnapshotssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateMerchantBalanceSnapshotsQuery(merchantBalanceSnapshotssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistMerchantBalanceSnapshotsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceSnapshots] failed checking merchantBalanceSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSnapshots with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeMerchantBalanceSnapshotsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"merchant_balance_snapshots\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceSnapshots] failed exec query")
	}
	return
}

type MerchantBalanceSnapshotsFieldParameter struct {
	param string
	args  []interface{}
}

func NewMerchantBalanceSnapshotsFieldParameter(param string, args ...interface{}) MerchantBalanceSnapshotsFieldParameter {
	return MerchantBalanceSnapshotsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateMerchantBalanceSnapshotsQuery(mapMerchantBalanceSnapshotss map[model.MerchantBalanceSnapshotsPrimaryID]MerchantBalanceSnapshotsUpdateFieldList, asTableValues string) (primaryIDs []model.MerchantBalanceSnapshotsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.MerchantBalanceSnapshotsPrimaryID]map[string]interface{}{}
	merchantBalanceSnapshotsSelectFields := NewMerchantBalanceSnapshotsSelectFields()
	for id, updateFields := range mapMerchantBalanceSnapshotss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.merchantBalanceSnapshotsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapMerchantBalanceSnapshotss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetMerchantBalanceSnapshotsFieldType(updateField.merchantBalanceSnapshotsField)))
			args = append(args, fields[string(updateField.merchantBalanceSnapshotsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.merchantBalanceSnapshotsField))
		if updateField.merchantBalanceSnapshotsField == merchantBalanceSnapshotsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.merchantBalanceSnapshotsField, asTableValues, updateField.merchantBalanceSnapshotsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.merchantBalanceSnapshotsField,
				"\"merchant_balance_snapshots\"", updateField.merchantBalanceSnapshotsField,
				asTableValues, updateField.merchantBalanceSnapshotsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeMerchantBalanceSnapshotsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.MerchantBalanceSnapshotsPrimaryID, asTableValue string) (whereQry string) {
	merchantBalanceSnapshotsSelectFields := NewMerchantBalanceSnapshotsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"merchant_balance_snapshots\".\"id\" = %s.\"id\"::"+GetMerchantBalanceSnapshotsFieldType(merchantBalanceSnapshotsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetMerchantBalanceSnapshotsFieldType(merchantBalanceSnapshotsField MerchantBalanceSnapshotsField) string {
	selectMerchantBalanceSnapshotsFields := NewMerchantBalanceSnapshotsSelectFields()
	switch merchantBalanceSnapshotsField {

	case selectMerchantBalanceSnapshotsFields.Id():
		return "uuid"

	case selectMerchantBalanceSnapshotsFields.BalanceAccountId():
		return "uuid"

	case selectMerchantBalanceSnapshotsFields.SnapshotAt():
		return "timestamptz"

	case selectMerchantBalanceSnapshotsFields.AvailableAmount():
		return "numeric"

	case selectMerchantBalanceSnapshotsFields.PendingAmount():
		return "numeric"

	case selectMerchantBalanceSnapshotsFields.ReservedAmount():
		return "numeric"

	case selectMerchantBalanceSnapshotsFields.DisputedAmount():
		return "numeric"

	case selectMerchantBalanceSnapshotsFields.NegativeAmount():
		return "numeric"

	case selectMerchantBalanceSnapshotsFields.Metadata():
		return "jsonb"

	case selectMerchantBalanceSnapshotsFields.MetaCreatedAt():
		return "timestamptz"

	case selectMerchantBalanceSnapshotsFields.MetaCreatedBy():
		return "uuid"

	case selectMerchantBalanceSnapshotsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectMerchantBalanceSnapshotsFields.MetaUpdatedBy():
		return "uuid"

	case selectMerchantBalanceSnapshotsFields.MetaDeletedAt():
		return "timestamptz"

	case selectMerchantBalanceSnapshotsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateMerchantBalanceSnapshots(ctx context.Context, merchantBalanceSnapshots *model.MerchantBalanceSnapshots, fieldsInsert ...MerchantBalanceSnapshotsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceSnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.MerchantBalanceSnapshotsPrimaryID{
		Id: merchantBalanceSnapshots.Id,
	}
	exists, err := repo.IsExistMerchantBalanceSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceSnapshots] failed checking merchantBalanceSnapshots whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "merchantBalanceSnapshots", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsMerchantBalanceSnapshots([]model.MerchantBalanceSnapshots{*merchantBalanceSnapshots}, fieldsInsert...)
	commandQuery := fmt.Sprintf(merchantBalanceSnapshotsQueries.insertMerchantBalanceSnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceSnapshots] failed exec create merchantBalanceSnapshots query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteMerchantBalanceSnapshotsByID(ctx context.Context, primaryID model.MerchantBalanceSnapshotsPrimaryID) (err error) {
	exists, err := repo.IsExistMerchantBalanceSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceSnapshotsByID] failed checking merchantBalanceSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeMerchantBalanceSnapshotsCompositePrimaryKeyWhere([]model.MerchantBalanceSnapshotsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(merchantBalanceSnapshotsQueries.deleteMerchantBalanceSnapshots + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceSnapshotsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceSnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceSnapshotsFilterResult, err error) {
	query, args, err := composeMerchantBalanceSnapshotsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSnapshotsByFilter] failed compose merchantBalanceSnapshots filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSnapshotsByFilter] failed get merchantBalanceSnapshots by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeMerchantBalanceSnapshotsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.MerchantBalanceSnapshotsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeMerchantBalanceSnapshotsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeMerchantBalanceSnapshotsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeMerchantBalanceSnapshotsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 15+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["balance_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"balance_account_id\"")
			selectedColumns["balance_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["snapshot_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"snapshot_at\"")
			selectedColumns["snapshot_at"] = struct{}{}
		}
		if _, selected := selectedColumns["available_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"available_amount\"")
			selectedColumns["available_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["pending_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"pending_amount\"")
			selectedColumns["pending_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["reserved_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"reserved_amount\"")
			selectedColumns["reserved_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["disputed_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"disputed_amount\"")
			selectedColumns["disputed_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["negative_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"negative_amount\"")
			selectedColumns["negative_amount"] = struct{}{}
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

type merchantBalanceSnapshotsFilterPlaceholder struct {
	index int
}

func (p *merchantBalanceSnapshotsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeMerchantBalanceSnapshotsFilterPredicate(filterField model.FilterField, placeholders *merchantBalanceSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeMerchantBalanceSnapshotsFilterSQLExpr(spec)
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

func composeMerchantBalanceSnapshotsFilterGroup(group model.FilterGroup, placeholders *merchantBalanceSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeMerchantBalanceSnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeMerchantBalanceSnapshotsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeMerchantBalanceSnapshotsFilterWhereQueries(filter model.Filter, placeholders *merchantBalanceSnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeMerchantBalanceSnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeMerchantBalanceSnapshotsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeMerchantBalanceSnapshotsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateMerchantBalanceSnapshotsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeMerchantBalanceSnapshotsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeMerchantBalanceSnapshotsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := merchantBalanceSnapshotsFilterPlaceholder{index: 1}
	whereQueries, err := composeMerchantBalanceSnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewMerchantBalanceSnapshotsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeMerchantBalanceSnapshotsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeMerchantBalanceSnapshotsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"merchant_balance_snapshots\" base%s", strings.Join(selectColumns, ","), composeMerchantBalanceSnapshotsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistMerchantBalanceSnapshotsByID(ctx context.Context, primaryID model.MerchantBalanceSnapshotsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeMerchantBalanceSnapshotsCompositePrimaryKeyWhere([]model.MerchantBalanceSnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", merchantBalanceSnapshotsQueries.selectCountMerchantBalanceSnapshots, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceSnapshotsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceSnapshots(ctx context.Context, selectFields ...MerchantBalanceSnapshotsField) (merchantBalanceSnapshotsList model.MerchantBalanceSnapshotsList, err error) {
	var (
		defaultMerchantBalanceSnapshotsSelectFields = defaultMerchantBalanceSnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceSnapshotsSelectFields = composeMerchantBalanceSnapshotsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(merchantBalanceSnapshotsQueries.selectMerchantBalanceSnapshots, defaultMerchantBalanceSnapshotsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &merchantBalanceSnapshotsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSnapshots] failed get merchantBalanceSnapshots list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceSnapshotsByID(ctx context.Context, primaryID model.MerchantBalanceSnapshotsPrimaryID, selectFields ...MerchantBalanceSnapshotsField) (merchantBalanceSnapshots model.MerchantBalanceSnapshots, err error) {
	var (
		defaultMerchantBalanceSnapshotsSelectFields = defaultMerchantBalanceSnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceSnapshotsSelectFields = composeMerchantBalanceSnapshotsSelectFields(selectFields...)
	}
	whereQry, params := composeMerchantBalanceSnapshotsCompositePrimaryKeyWhere([]model.MerchantBalanceSnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf(merchantBalanceSnapshotsQueries.selectMerchantBalanceSnapshots+" WHERE "+whereQry, defaultMerchantBalanceSnapshotsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &merchantBalanceSnapshots, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("merchantBalanceSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSnapshotsByID] failed get merchantBalanceSnapshots")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateMerchantBalanceSnapshotsByID(ctx context.Context, primaryID model.MerchantBalanceSnapshotsPrimaryID, merchantBalanceSnapshots *model.MerchantBalanceSnapshots, merchantBalanceSnapshotsUpdateFields ...MerchantBalanceSnapshotsUpdateField) (err error) {
	exists, err := repo.IsExistMerchantBalanceSnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSnapshots] failed checking merchantBalanceSnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if merchantBalanceSnapshots == nil {
		if len(merchantBalanceSnapshotsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateMerchantBalanceSnapshotsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		merchantBalanceSnapshots = &model.MerchantBalanceSnapshots{}
	}
	var (
		defaultMerchantBalanceSnapshotsUpdateFields = defaultMerchantBalanceSnapshotsUpdateFields(*merchantBalanceSnapshots)
		tempUpdateField                             MerchantBalanceSnapshotsUpdateFieldList
		selectFields                                = NewMerchantBalanceSnapshotsSelectFields()
	)
	if len(merchantBalanceSnapshotsUpdateFields) > 0 {
		for _, updateField := range merchantBalanceSnapshotsUpdateFields {
			if updateField.merchantBalanceSnapshotsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultMerchantBalanceSnapshotsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeMerchantBalanceSnapshotsCompositePrimaryKeyWhere([]model.MerchantBalanceSnapshotsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsMerchantBalanceSnapshotsCommand(defaultMerchantBalanceSnapshotsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(merchantBalanceSnapshotsQueries.updateMerchantBalanceSnapshots+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSnapshots] error when try to update merchantBalanceSnapshots by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateMerchantBalanceSnapshotsByFilter(ctx context.Context, filter model.Filter, merchantBalanceSnapshotsUpdateFields ...MerchantBalanceSnapshotsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(merchantBalanceSnapshotsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields MerchantBalanceSnapshotsUpdateFieldList
		selectFields = NewMerchantBalanceSnapshotsSelectFields()
	)
	for _, updateField := range merchantBalanceSnapshotsUpdateFields {
		if updateField.merchantBalanceSnapshotsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsMerchantBalanceSnapshotsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := merchantBalanceSnapshotsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeMerchantBalanceSnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"merchant_balance_snapshots\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSnapshotsByFilter] error when try to update merchantBalanceSnapshots by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSnapshotsByFilter] failed get rows affected")
	}
	return
}

var (
	merchantBalanceSnapshotsQueries = struct {
		selectMerchantBalanceSnapshots      string
		selectCountMerchantBalanceSnapshots string
		deleteMerchantBalanceSnapshots      string
		updateMerchantBalanceSnapshots      string
		insertMerchantBalanceSnapshots      string
	}{
		selectMerchantBalanceSnapshots:      "SELECT %s FROM \"merchant_balance_snapshots\"",
		selectCountMerchantBalanceSnapshots: "SELECT COUNT(\"id\") FROM \"merchant_balance_snapshots\"",
		deleteMerchantBalanceSnapshots:      "DELETE FROM \"merchant_balance_snapshots\"",
		updateMerchantBalanceSnapshots:      "UPDATE \"merchant_balance_snapshots\" SET %s ",
		insertMerchantBalanceSnapshots:      "INSERT INTO \"merchant_balance_snapshots\" %s VALUES %s",
	}
)

type MerchantBalanceSnapshotsRepository interface {
	CreateMerchantBalanceSnapshots(ctx context.Context, merchantBalanceSnapshots *model.MerchantBalanceSnapshots, fieldsInsert ...MerchantBalanceSnapshotsField) error
	BulkCreateMerchantBalanceSnapshots(ctx context.Context, merchantBalanceSnapshotsList []*model.MerchantBalanceSnapshots, fieldsInsert ...MerchantBalanceSnapshotsField) error
	ResolveMerchantBalanceSnapshots(ctx context.Context, selectFields ...MerchantBalanceSnapshotsField) (model.MerchantBalanceSnapshotsList, error)
	ResolveMerchantBalanceSnapshotsByID(ctx context.Context, primaryID model.MerchantBalanceSnapshotsPrimaryID, selectFields ...MerchantBalanceSnapshotsField) (model.MerchantBalanceSnapshots, error)
	UpdateMerchantBalanceSnapshotsByID(ctx context.Context, id model.MerchantBalanceSnapshotsPrimaryID, merchantBalanceSnapshots *model.MerchantBalanceSnapshots, merchantBalanceSnapshotsUpdateFields ...MerchantBalanceSnapshotsUpdateField) error
	UpdateMerchantBalanceSnapshotsByFilter(ctx context.Context, filter model.Filter, merchantBalanceSnapshotsUpdateFields ...MerchantBalanceSnapshotsUpdateField) (rowsAffected int64, err error)
	BulkUpdateMerchantBalanceSnapshots(ctx context.Context, merchantBalanceSnapshotsListMap map[model.MerchantBalanceSnapshotsPrimaryID]*model.MerchantBalanceSnapshots, MerchantBalanceSnapshotssMapUpdateFieldsRequest map[model.MerchantBalanceSnapshotsPrimaryID]MerchantBalanceSnapshotsUpdateFieldList) (err error)
	DeleteMerchantBalanceSnapshotsByID(ctx context.Context, id model.MerchantBalanceSnapshotsPrimaryID) error
	BulkDeleteMerchantBalanceSnapshotsByIDs(ctx context.Context, ids []model.MerchantBalanceSnapshotsPrimaryID) error
	ResolveMerchantBalanceSnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceSnapshotsFilterResult, err error)
	IsExistMerchantBalanceSnapshotsByIDs(ctx context.Context, ids []model.MerchantBalanceSnapshotsPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceSnapshotsPrimaryID, err error)
	IsExistMerchantBalanceSnapshotsByID(ctx context.Context, id model.MerchantBalanceSnapshotsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
