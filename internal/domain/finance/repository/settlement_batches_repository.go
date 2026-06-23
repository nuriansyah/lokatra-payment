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

func composeInsertFieldsAndParamsSettlementBatches(settlementBatchesList []model.SettlementBatches, fieldsInsert ...SettlementBatchesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewSettlementBatchesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, settlementBatches := range settlementBatchesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, settlementBatches.Id)
			case selectField.BatchCode():
				args = append(args, settlementBatches.BatchCode)
			case selectField.MerchantPartyId():
				args = append(args, settlementBatches.MerchantPartyId)
			case selectField.CurrencyCode():
				args = append(args, settlementBatches.CurrencyCode)
			case selectField.PeriodStart():
				args = append(args, settlementBatches.PeriodStart)
			case selectField.PeriodEnd():
				args = append(args, settlementBatches.PeriodEnd)
			case selectField.GrossAmount():
				args = append(args, settlementBatches.GrossAmount)
			case selectField.FeeAmount():
				args = append(args, settlementBatches.FeeAmount)
			case selectField.TaxAmount():
				args = append(args, settlementBatches.TaxAmount)
			case selectField.ReserveAmount():
				args = append(args, settlementBatches.ReserveAmount)
			case selectField.AdjustmentAmount():
				args = append(args, settlementBatches.AdjustmentAmount)
			case selectField.NetAmount():
				args = append(args, settlementBatches.NetAmount)
			case selectField.BatchStatus():
				args = append(args, settlementBatches.BatchStatus)
			case selectField.ApprovedAt():
				args = append(args, settlementBatches.ApprovedAt)
			case selectField.LockedAt():
				args = append(args, settlementBatches.LockedAt)
			case selectField.Metadata():
				args = append(args, settlementBatches.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, settlementBatches.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, settlementBatches.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, settlementBatches.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, settlementBatches.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, settlementBatches.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, settlementBatches.MetaDeletedBy)

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

func composeSettlementBatchesCompositePrimaryKeyWhere(primaryIDs []model.SettlementBatchesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"settlement_batches\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultSettlementBatchesSelectFields() string {
	fields := NewSettlementBatchesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeSettlementBatchesSelectFields(selectFields ...SettlementBatchesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type SettlementBatchesField string
type SettlementBatchesFieldList []SettlementBatchesField

type SettlementBatchesSelectFields struct {
}

func (ss SettlementBatchesSelectFields) Id() SettlementBatchesField {
	return SettlementBatchesField("id")
}

func (ss SettlementBatchesSelectFields) BatchCode() SettlementBatchesField {
	return SettlementBatchesField("batch_code")
}

func (ss SettlementBatchesSelectFields) MerchantPartyId() SettlementBatchesField {
	return SettlementBatchesField("merchant_party_id")
}

func (ss SettlementBatchesSelectFields) CurrencyCode() SettlementBatchesField {
	return SettlementBatchesField("currency_code")
}

func (ss SettlementBatchesSelectFields) PeriodStart() SettlementBatchesField {
	return SettlementBatchesField("period_start")
}

func (ss SettlementBatchesSelectFields) PeriodEnd() SettlementBatchesField {
	return SettlementBatchesField("period_end")
}

func (ss SettlementBatchesSelectFields) GrossAmount() SettlementBatchesField {
	return SettlementBatchesField("gross_amount")
}

func (ss SettlementBatchesSelectFields) FeeAmount() SettlementBatchesField {
	return SettlementBatchesField("fee_amount")
}

func (ss SettlementBatchesSelectFields) TaxAmount() SettlementBatchesField {
	return SettlementBatchesField("tax_amount")
}

func (ss SettlementBatchesSelectFields) ReserveAmount() SettlementBatchesField {
	return SettlementBatchesField("reserve_amount")
}

func (ss SettlementBatchesSelectFields) AdjustmentAmount() SettlementBatchesField {
	return SettlementBatchesField("adjustment_amount")
}

func (ss SettlementBatchesSelectFields) NetAmount() SettlementBatchesField {
	return SettlementBatchesField("net_amount")
}

func (ss SettlementBatchesSelectFields) BatchStatus() SettlementBatchesField {
	return SettlementBatchesField("batch_status")
}

func (ss SettlementBatchesSelectFields) ApprovedAt() SettlementBatchesField {
	return SettlementBatchesField("approved_at")
}

func (ss SettlementBatchesSelectFields) LockedAt() SettlementBatchesField {
	return SettlementBatchesField("locked_at")
}

func (ss SettlementBatchesSelectFields) Metadata() SettlementBatchesField {
	return SettlementBatchesField("metadata")
}

func (ss SettlementBatchesSelectFields) MetaCreatedAt() SettlementBatchesField {
	return SettlementBatchesField("meta_created_at")
}

func (ss SettlementBatchesSelectFields) MetaCreatedBy() SettlementBatchesField {
	return SettlementBatchesField("meta_created_by")
}

func (ss SettlementBatchesSelectFields) MetaUpdatedAt() SettlementBatchesField {
	return SettlementBatchesField("meta_updated_at")
}

func (ss SettlementBatchesSelectFields) MetaUpdatedBy() SettlementBatchesField {
	return SettlementBatchesField("meta_updated_by")
}

func (ss SettlementBatchesSelectFields) MetaDeletedAt() SettlementBatchesField {
	return SettlementBatchesField("meta_deleted_at")
}

func (ss SettlementBatchesSelectFields) MetaDeletedBy() SettlementBatchesField {
	return SettlementBatchesField("meta_deleted_by")
}

func (ss SettlementBatchesSelectFields) All() SettlementBatchesFieldList {
	return []SettlementBatchesField{
		ss.Id(),
		ss.BatchCode(),
		ss.MerchantPartyId(),
		ss.CurrencyCode(),
		ss.PeriodStart(),
		ss.PeriodEnd(),
		ss.GrossAmount(),
		ss.FeeAmount(),
		ss.TaxAmount(),
		ss.ReserveAmount(),
		ss.AdjustmentAmount(),
		ss.NetAmount(),
		ss.BatchStatus(),
		ss.ApprovedAt(),
		ss.LockedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewSettlementBatchesSelectFields() SettlementBatchesSelectFields {
	return SettlementBatchesSelectFields{}
}

type SettlementBatchesUpdateFieldOption struct {
	useIncrement bool
}
type SettlementBatchesUpdateField struct {
	settlementBatchesField SettlementBatchesField
	opt                    SettlementBatchesUpdateFieldOption
	value                  interface{}
}
type SettlementBatchesUpdateFieldList []SettlementBatchesUpdateField

func defaultSettlementBatchesUpdateFieldOption() SettlementBatchesUpdateFieldOption {
	return SettlementBatchesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementSettlementBatchesOption(useIncrement bool) func(*SettlementBatchesUpdateFieldOption) {
	return func(pcufo *SettlementBatchesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewSettlementBatchesUpdateField(field SettlementBatchesField, val interface{}, opts ...func(*SettlementBatchesUpdateFieldOption)) SettlementBatchesUpdateField {
	defaultOpt := defaultSettlementBatchesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return SettlementBatchesUpdateField{
		settlementBatchesField: field,
		value:                  val,
		opt:                    defaultOpt,
	}
}
func defaultSettlementBatchesUpdateFields(settlementBatches model.SettlementBatches) (settlementBatchesUpdateFieldList SettlementBatchesUpdateFieldList) {
	selectFields := NewSettlementBatchesSelectFields()
	settlementBatchesUpdateFieldList = append(settlementBatchesUpdateFieldList,
		NewSettlementBatchesUpdateField(selectFields.Id(), settlementBatches.Id),
		NewSettlementBatchesUpdateField(selectFields.BatchCode(), settlementBatches.BatchCode),
		NewSettlementBatchesUpdateField(selectFields.MerchantPartyId(), settlementBatches.MerchantPartyId),
		NewSettlementBatchesUpdateField(selectFields.CurrencyCode(), settlementBatches.CurrencyCode),
		NewSettlementBatchesUpdateField(selectFields.PeriodStart(), settlementBatches.PeriodStart),
		NewSettlementBatchesUpdateField(selectFields.PeriodEnd(), settlementBatches.PeriodEnd),
		NewSettlementBatchesUpdateField(selectFields.GrossAmount(), settlementBatches.GrossAmount),
		NewSettlementBatchesUpdateField(selectFields.FeeAmount(), settlementBatches.FeeAmount),
		NewSettlementBatchesUpdateField(selectFields.TaxAmount(), settlementBatches.TaxAmount),
		NewSettlementBatchesUpdateField(selectFields.ReserveAmount(), settlementBatches.ReserveAmount),
		NewSettlementBatchesUpdateField(selectFields.AdjustmentAmount(), settlementBatches.AdjustmentAmount),
		NewSettlementBatchesUpdateField(selectFields.NetAmount(), settlementBatches.NetAmount),
		NewSettlementBatchesUpdateField(selectFields.BatchStatus(), settlementBatches.BatchStatus),
		NewSettlementBatchesUpdateField(selectFields.ApprovedAt(), settlementBatches.ApprovedAt),
		NewSettlementBatchesUpdateField(selectFields.LockedAt(), settlementBatches.LockedAt),
		NewSettlementBatchesUpdateField(selectFields.Metadata(), settlementBatches.Metadata),
		NewSettlementBatchesUpdateField(selectFields.MetaCreatedAt(), settlementBatches.MetaCreatedAt),
		NewSettlementBatchesUpdateField(selectFields.MetaCreatedBy(), settlementBatches.MetaCreatedBy),
		NewSettlementBatchesUpdateField(selectFields.MetaUpdatedAt(), settlementBatches.MetaUpdatedAt),
		NewSettlementBatchesUpdateField(selectFields.MetaUpdatedBy(), settlementBatches.MetaUpdatedBy),
		NewSettlementBatchesUpdateField(selectFields.MetaDeletedAt(), settlementBatches.MetaDeletedAt),
		NewSettlementBatchesUpdateField(selectFields.MetaDeletedBy(), settlementBatches.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsSettlementBatchesCommand(settlementBatchesUpdateFieldList SettlementBatchesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range settlementBatchesUpdateFieldList {
		field := string(updateField.settlementBatchesField)
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

func (repo *RepositoryImpl) BulkCreateSettlementBatches(ctx context.Context, settlementBatchesList []*model.SettlementBatches, fieldsInsert ...SettlementBatchesField) (err error) {
	var (
		fieldsStr                  string
		valueListStr               []string
		argsList                   []interface{}
		primaryIds                 []model.SettlementBatchesPrimaryID
		settlementBatchesValueList []model.SettlementBatches
	)

	if len(fieldsInsert) == 0 {
		selectField := NewSettlementBatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, settlementBatches := range settlementBatchesList {

		primaryIds = append(primaryIds, settlementBatches.ToSettlementBatchesPrimaryID())

		settlementBatchesValueList = append(settlementBatchesValueList, *settlementBatches)
	}

	_, notFoundIds, err := repo.IsExistSettlementBatchesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementBatches] failed checking settlementBatches whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.SettlementBatchesPrimaryID{}
		mapNotFoundIds := map[model.SettlementBatchesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "settlementBatches", fmt.Sprintf("settlementBatches with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsSettlementBatches(settlementBatchesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(settlementBatchesQueries.insertSettlementBatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementBatches] failed exec create settlementBatches query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteSettlementBatchesByIDs(ctx context.Context, primaryIDs []model.SettlementBatchesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistSettlementBatchesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementBatchesByIDs] failed checking settlementBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatches with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_batches\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(settlementBatchesQueries.deleteSettlementBatches + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementBatchesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementBatchesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistSettlementBatchesByIDs(ctx context.Context, ids []model.SettlementBatchesPrimaryID) (exists bool, notFoundIds []model.SettlementBatchesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_batches\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(settlementBatchesQueries.selectSettlementBatches, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementBatchesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.SettlementBatchesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementBatchesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.SettlementBatchesPrimaryID]bool{}
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

// BulkUpdateSettlementBatches is used to bulk update settlementBatches, by default it will update all field
// if want to update specific field, then fill settlementBatchessMapUpdateFieldsRequest else please fill settlementBatchessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateSettlementBatches(ctx context.Context, settlementBatchessMap map[model.SettlementBatchesPrimaryID]*model.SettlementBatches, settlementBatchessMapUpdateFieldsRequest map[model.SettlementBatchesPrimaryID]SettlementBatchesUpdateFieldList) (err error) {
	if len(settlementBatchessMap) == 0 && len(settlementBatchessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		settlementBatchessMapUpdateField map[model.SettlementBatchesPrimaryID]SettlementBatchesUpdateFieldList = map[model.SettlementBatchesPrimaryID]SettlementBatchesUpdateFieldList{}
		asTableValues                    string                                                                = "myvalues"
	)

	if len(settlementBatchessMap) > 0 {
		for id, settlementBatches := range settlementBatchessMap {
			if settlementBatches == nil {
				log.Error().Err(err).Msg("[BulkUpdateSettlementBatches] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			settlementBatchessMapUpdateField[id] = defaultSettlementBatchesUpdateFields(*settlementBatches)
		}
	} else {
		settlementBatchessMapUpdateField = settlementBatchessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateSettlementBatchesQuery(settlementBatchessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistSettlementBatchesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementBatches] failed checking settlementBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatches with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeSettlementBatchesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"settlement_batches\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementBatches] failed exec query")
	}
	return
}

type SettlementBatchesFieldParameter struct {
	param string
	args  []interface{}
}

func NewSettlementBatchesFieldParameter(param string, args ...interface{}) SettlementBatchesFieldParameter {
	return SettlementBatchesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateSettlementBatchesQuery(mapSettlementBatchess map[model.SettlementBatchesPrimaryID]SettlementBatchesUpdateFieldList, asTableValues string) (primaryIDs []model.SettlementBatchesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.SettlementBatchesPrimaryID]map[string]interface{}{}
	settlementBatchesSelectFields := NewSettlementBatchesSelectFields()
	for id, updateFields := range mapSettlementBatchess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.settlementBatchesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapSettlementBatchess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetSettlementBatchesFieldType(updateField.settlementBatchesField)))
			args = append(args, fields[string(updateField.settlementBatchesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.settlementBatchesField))
		if updateField.settlementBatchesField == settlementBatchesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.settlementBatchesField, asTableValues, updateField.settlementBatchesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.settlementBatchesField,
				"\"settlement_batches\"", updateField.settlementBatchesField,
				asTableValues, updateField.settlementBatchesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeSettlementBatchesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.SettlementBatchesPrimaryID, asTableValue string) (whereQry string) {
	settlementBatchesSelectFields := NewSettlementBatchesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"settlement_batches\".\"id\" = %s.\"id\"::"+GetSettlementBatchesFieldType(settlementBatchesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetSettlementBatchesFieldType(settlementBatchesField SettlementBatchesField) string {
	selectSettlementBatchesFields := NewSettlementBatchesSelectFields()
	switch settlementBatchesField {

	case selectSettlementBatchesFields.Id():
		return "uuid"

	case selectSettlementBatchesFields.BatchCode():
		return "text"

	case selectSettlementBatchesFields.MerchantPartyId():
		return "uuid"

	case selectSettlementBatchesFields.CurrencyCode():
		return "text"

	case selectSettlementBatchesFields.PeriodStart():
		return "timestamptz"

	case selectSettlementBatchesFields.PeriodEnd():
		return "timestamptz"

	case selectSettlementBatchesFields.GrossAmount():
		return "numeric"

	case selectSettlementBatchesFields.FeeAmount():
		return "numeric"

	case selectSettlementBatchesFields.TaxAmount():
		return "numeric"

	case selectSettlementBatchesFields.ReserveAmount():
		return "numeric"

	case selectSettlementBatchesFields.AdjustmentAmount():
		return "numeric"

	case selectSettlementBatchesFields.NetAmount():
		return "numeric"

	case selectSettlementBatchesFields.BatchStatus():
		return "settlement_batches_batch_status_enum"

	case selectSettlementBatchesFields.ApprovedAt():
		return "timestamptz"

	case selectSettlementBatchesFields.LockedAt():
		return "timestamptz"

	case selectSettlementBatchesFields.Metadata():
		return "jsonb"

	case selectSettlementBatchesFields.MetaCreatedAt():
		return "timestamptz"

	case selectSettlementBatchesFields.MetaCreatedBy():
		return "uuid"

	case selectSettlementBatchesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectSettlementBatchesFields.MetaUpdatedBy():
		return "uuid"

	case selectSettlementBatchesFields.MetaDeletedAt():
		return "timestamptz"

	case selectSettlementBatchesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateSettlementBatches(ctx context.Context, settlementBatches *model.SettlementBatches, fieldsInsert ...SettlementBatchesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewSettlementBatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.SettlementBatchesPrimaryID{
		Id: settlementBatches.Id,
	}
	exists, err := repo.IsExistSettlementBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementBatches] failed checking settlementBatches whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "settlementBatches", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsSettlementBatches([]model.SettlementBatches{*settlementBatches}, fieldsInsert...)
	commandQuery := fmt.Sprintf(settlementBatchesQueries.insertSettlementBatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementBatches] failed exec create settlementBatches query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteSettlementBatchesByID(ctx context.Context, primaryID model.SettlementBatchesPrimaryID) (err error) {
	exists, err := repo.IsExistSettlementBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementBatchesByID] failed checking settlementBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeSettlementBatchesCompositePrimaryKeyWhere([]model.SettlementBatchesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(settlementBatchesQueries.deleteSettlementBatches + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementBatchesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementBatchesByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementBatchesFilterResult, err error) {
	query, args, err := composeSettlementBatchesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementBatchesByFilter] failed compose settlementBatches filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementBatchesByFilter] failed get settlementBatches by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeSettlementBatchesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.SettlementBatchesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeSettlementBatchesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeSettlementBatchesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeSettlementBatchesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 22 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewSettlementBatchesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 22+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["batch_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"batch_code\"")
			selectedColumns["batch_code"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["period_start"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_start\"")
			selectedColumns["period_start"] = struct{}{}
		}
		if _, selected := selectedColumns["period_end"]; !selected {
			selectColumns = append(selectColumns, "base.\"period_end\"")
			selectedColumns["period_end"] = struct{}{}
		}
		if _, selected := selectedColumns["gross_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"gross_amount\"")
			selectedColumns["gross_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_amount\"")
			selectedColumns["fee_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_amount\"")
			selectedColumns["tax_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["reserve_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"reserve_amount\"")
			selectedColumns["reserve_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["adjustment_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"adjustment_amount\"")
			selectedColumns["adjustment_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["net_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"net_amount\"")
			selectedColumns["net_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["batch_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"batch_status\"")
			selectedColumns["batch_status"] = struct{}{}
		}
		if _, selected := selectedColumns["approved_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"approved_at\"")
			selectedColumns["approved_at"] = struct{}{}
		}
		if _, selected := selectedColumns["locked_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"locked_at\"")
			selectedColumns["locked_at"] = struct{}{}
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

type settlementBatchesFilterPlaceholder struct {
	index int
}

func (p *settlementBatchesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeSettlementBatchesFilterPredicate(filterField model.FilterField, placeholders *settlementBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewSettlementBatchesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeSettlementBatchesFilterSQLExpr(spec)
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

func composeSettlementBatchesFilterGroup(group model.FilterGroup, placeholders *settlementBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeSettlementBatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeSettlementBatchesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeSettlementBatchesFilterWhereQueries(filter model.Filter, placeholders *settlementBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeSettlementBatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeSettlementBatchesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeSettlementBatchesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateSettlementBatchesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeSettlementBatchesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeSettlementBatchesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := settlementBatchesFilterPlaceholder{index: 1}
	whereQueries, err := composeSettlementBatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewSettlementBatchesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeSettlementBatchesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeSettlementBatchesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"settlement_batches\" base%s", strings.Join(selectColumns, ","), composeSettlementBatchesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistSettlementBatchesByID(ctx context.Context, primaryID model.SettlementBatchesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeSettlementBatchesCompositePrimaryKeyWhere([]model.SettlementBatchesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", settlementBatchesQueries.selectCountSettlementBatches, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementBatchesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementBatches(ctx context.Context, selectFields ...SettlementBatchesField) (settlementBatchesList model.SettlementBatchesList, err error) {
	var (
		defaultSettlementBatchesSelectFields = defaultSettlementBatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementBatchesSelectFields = composeSettlementBatchesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(settlementBatchesQueries.selectSettlementBatches, defaultSettlementBatchesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &settlementBatchesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementBatches] failed get settlementBatches list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementBatchesByID(ctx context.Context, primaryID model.SettlementBatchesPrimaryID, selectFields ...SettlementBatchesField) (settlementBatches model.SettlementBatches, err error) {
	var (
		defaultSettlementBatchesSelectFields = defaultSettlementBatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementBatchesSelectFields = composeSettlementBatchesSelectFields(selectFields...)
	}
	whereQry, params := composeSettlementBatchesCompositePrimaryKeyWhere([]model.SettlementBatchesPrimaryID{primaryID})
	query := fmt.Sprintf(settlementBatchesQueries.selectSettlementBatches+" WHERE "+whereQry, defaultSettlementBatchesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &settlementBatches, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("settlementBatches with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveSettlementBatchesByID] failed get settlementBatches")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateSettlementBatchesByID(ctx context.Context, primaryID model.SettlementBatchesPrimaryID, settlementBatches *model.SettlementBatches, settlementBatchesUpdateFields ...SettlementBatchesUpdateField) (err error) {
	exists, err := repo.IsExistSettlementBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatches] failed checking settlementBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementBatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if settlementBatches == nil {
		if len(settlementBatchesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateSettlementBatchesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		settlementBatches = &model.SettlementBatches{}
	}
	var (
		defaultSettlementBatchesUpdateFields = defaultSettlementBatchesUpdateFields(*settlementBatches)
		tempUpdateField                      SettlementBatchesUpdateFieldList
		selectFields                         = NewSettlementBatchesSelectFields()
	)
	if len(settlementBatchesUpdateFields) > 0 {
		for _, updateField := range settlementBatchesUpdateFields {
			if updateField.settlementBatchesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultSettlementBatchesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeSettlementBatchesCompositePrimaryKeyWhere([]model.SettlementBatchesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsSettlementBatchesCommand(defaultSettlementBatchesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(settlementBatchesQueries.updateSettlementBatches+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatches] error when try to update settlementBatches by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateSettlementBatchesByFilter(ctx context.Context, filter model.Filter, settlementBatchesUpdateFields ...SettlementBatchesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(settlementBatchesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields SettlementBatchesUpdateFieldList
		selectFields = NewSettlementBatchesSelectFields()
	)
	for _, updateField := range settlementBatchesUpdateFields {
		if updateField.settlementBatchesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsSettlementBatchesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := settlementBatchesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeSettlementBatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"settlement_batches\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatchesByFilter] error when try to update settlementBatches by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementBatchesByFilter] failed get rows affected")
	}
	return
}

var (
	settlementBatchesQueries = struct {
		selectSettlementBatches      string
		selectCountSettlementBatches string
		deleteSettlementBatches      string
		updateSettlementBatches      string
		insertSettlementBatches      string
	}{
		selectSettlementBatches:      "SELECT %s FROM \"settlement_batches\"",
		selectCountSettlementBatches: "SELECT COUNT(\"id\") FROM \"settlement_batches\"",
		deleteSettlementBatches:      "DELETE FROM \"settlement_batches\"",
		updateSettlementBatches:      "UPDATE \"settlement_batches\" SET %s ",
		insertSettlementBatches:      "INSERT INTO \"settlement_batches\" %s VALUES %s",
	}
)

type SettlementBatchesRepository interface {
	CreateSettlementBatches(ctx context.Context, settlementBatches *model.SettlementBatches, fieldsInsert ...SettlementBatchesField) error
	BulkCreateSettlementBatches(ctx context.Context, settlementBatchesList []*model.SettlementBatches, fieldsInsert ...SettlementBatchesField) error
	ResolveSettlementBatches(ctx context.Context, selectFields ...SettlementBatchesField) (model.SettlementBatchesList, error)
	ResolveSettlementBatchesByID(ctx context.Context, primaryID model.SettlementBatchesPrimaryID, selectFields ...SettlementBatchesField) (model.SettlementBatches, error)
	UpdateSettlementBatchesByID(ctx context.Context, id model.SettlementBatchesPrimaryID, settlementBatches *model.SettlementBatches, settlementBatchesUpdateFields ...SettlementBatchesUpdateField) error
	UpdateSettlementBatchesByFilter(ctx context.Context, filter model.Filter, settlementBatchesUpdateFields ...SettlementBatchesUpdateField) (rowsAffected int64, err error)
	BulkUpdateSettlementBatches(ctx context.Context, settlementBatchesListMap map[model.SettlementBatchesPrimaryID]*model.SettlementBatches, SettlementBatchessMapUpdateFieldsRequest map[model.SettlementBatchesPrimaryID]SettlementBatchesUpdateFieldList) (err error)
	DeleteSettlementBatchesByID(ctx context.Context, id model.SettlementBatchesPrimaryID) error
	BulkDeleteSettlementBatchesByIDs(ctx context.Context, ids []model.SettlementBatchesPrimaryID) error
	ResolveSettlementBatchesByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementBatchesFilterResult, err error)
	IsExistSettlementBatchesByIDs(ctx context.Context, ids []model.SettlementBatchesPrimaryID) (exists bool, notFoundIds []model.SettlementBatchesPrimaryID, err error)
	IsExistSettlementBatchesByID(ctx context.Context, id model.SettlementBatchesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
