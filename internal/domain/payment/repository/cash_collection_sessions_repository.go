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

func composeInsertFieldsAndParamsCashCollectionSessions(cashCollectionSessionsList []model.CashCollectionSessions, fieldsInsert ...CashCollectionSessionsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewCashCollectionSessionsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, cashCollectionSessions := range cashCollectionSessionsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, cashCollectionSessions.Id)
			case selectField.SessionCode():
				args = append(args, cashCollectionSessions.SessionCode)
			case selectField.MerchantId():
				args = append(args, cashCollectionSessions.MerchantId)
			case selectField.CollectorId():
				args = append(args, cashCollectionSessions.CollectorId)
			case selectField.LocationId():
				args = append(args, cashCollectionSessions.LocationId)
			case selectField.OpenedAt():
				args = append(args, cashCollectionSessions.OpenedAt)
			case selectField.ClosedAt():
				args = append(args, cashCollectionSessions.ClosedAt)
			case selectField.Status():
				args = append(args, cashCollectionSessions.Status)
			case selectField.OpeningFloatAmount():
				args = append(args, cashCollectionSessions.OpeningFloatAmount)
			case selectField.ExpectedAmount():
				args = append(args, cashCollectionSessions.ExpectedAmount)
			case selectField.CountedAmount():
				args = append(args, cashCollectionSessions.CountedAmount)
			case selectField.VarianceAmount():
				args = append(args, cashCollectionSessions.VarianceAmount)
			case selectField.Currency():
				args = append(args, cashCollectionSessions.Currency)
			case selectField.Notes():
				args = append(args, cashCollectionSessions.Notes)
			case selectField.Metadata():
				args = append(args, cashCollectionSessions.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, cashCollectionSessions.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, cashCollectionSessions.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, cashCollectionSessions.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, cashCollectionSessions.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, cashCollectionSessions.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, cashCollectionSessions.MetaDeletedBy)

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

func composeCashCollectionSessionsCompositePrimaryKeyWhere(primaryIDs []model.CashCollectionSessionsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"cash_collection_sessions\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultCashCollectionSessionsSelectFields() string {
	fields := NewCashCollectionSessionsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeCashCollectionSessionsSelectFields(selectFields ...CashCollectionSessionsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type CashCollectionSessionsField string
type CashCollectionSessionsFieldList []CashCollectionSessionsField

type CashCollectionSessionsSelectFields struct {
}

func (ss CashCollectionSessionsSelectFields) Id() CashCollectionSessionsField {
	return CashCollectionSessionsField("id")
}

func (ss CashCollectionSessionsSelectFields) SessionCode() CashCollectionSessionsField {
	return CashCollectionSessionsField("session_code")
}

func (ss CashCollectionSessionsSelectFields) MerchantId() CashCollectionSessionsField {
	return CashCollectionSessionsField("merchant_id")
}

func (ss CashCollectionSessionsSelectFields) CollectorId() CashCollectionSessionsField {
	return CashCollectionSessionsField("collector_id")
}

func (ss CashCollectionSessionsSelectFields) LocationId() CashCollectionSessionsField {
	return CashCollectionSessionsField("location_id")
}

func (ss CashCollectionSessionsSelectFields) OpenedAt() CashCollectionSessionsField {
	return CashCollectionSessionsField("opened_at")
}

func (ss CashCollectionSessionsSelectFields) ClosedAt() CashCollectionSessionsField {
	return CashCollectionSessionsField("closed_at")
}

func (ss CashCollectionSessionsSelectFields) Status() CashCollectionSessionsField {
	return CashCollectionSessionsField("status")
}

func (ss CashCollectionSessionsSelectFields) OpeningFloatAmount() CashCollectionSessionsField {
	return CashCollectionSessionsField("opening_float_amount")
}

func (ss CashCollectionSessionsSelectFields) ExpectedAmount() CashCollectionSessionsField {
	return CashCollectionSessionsField("expected_amount")
}

func (ss CashCollectionSessionsSelectFields) CountedAmount() CashCollectionSessionsField {
	return CashCollectionSessionsField("counted_amount")
}

func (ss CashCollectionSessionsSelectFields) VarianceAmount() CashCollectionSessionsField {
	return CashCollectionSessionsField("variance_amount")
}

func (ss CashCollectionSessionsSelectFields) Currency() CashCollectionSessionsField {
	return CashCollectionSessionsField("currency")
}

func (ss CashCollectionSessionsSelectFields) Notes() CashCollectionSessionsField {
	return CashCollectionSessionsField("notes")
}

func (ss CashCollectionSessionsSelectFields) Metadata() CashCollectionSessionsField {
	return CashCollectionSessionsField("metadata")
}

func (ss CashCollectionSessionsSelectFields) MetaCreatedAt() CashCollectionSessionsField {
	return CashCollectionSessionsField("meta_created_at")
}

func (ss CashCollectionSessionsSelectFields) MetaCreatedBy() CashCollectionSessionsField {
	return CashCollectionSessionsField("meta_created_by")
}

func (ss CashCollectionSessionsSelectFields) MetaUpdatedAt() CashCollectionSessionsField {
	return CashCollectionSessionsField("meta_updated_at")
}

func (ss CashCollectionSessionsSelectFields) MetaUpdatedBy() CashCollectionSessionsField {
	return CashCollectionSessionsField("meta_updated_by")
}

func (ss CashCollectionSessionsSelectFields) MetaDeletedAt() CashCollectionSessionsField {
	return CashCollectionSessionsField("meta_deleted_at")
}

func (ss CashCollectionSessionsSelectFields) MetaDeletedBy() CashCollectionSessionsField {
	return CashCollectionSessionsField("meta_deleted_by")
}

func (ss CashCollectionSessionsSelectFields) All() CashCollectionSessionsFieldList {
	return []CashCollectionSessionsField{
		ss.Id(),
		ss.SessionCode(),
		ss.MerchantId(),
		ss.CollectorId(),
		ss.LocationId(),
		ss.OpenedAt(),
		ss.ClosedAt(),
		ss.Status(),
		ss.OpeningFloatAmount(),
		ss.ExpectedAmount(),
		ss.CountedAmount(),
		ss.VarianceAmount(),
		ss.Currency(),
		ss.Notes(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewCashCollectionSessionsSelectFields() CashCollectionSessionsSelectFields {
	return CashCollectionSessionsSelectFields{}
}

type CashCollectionSessionsUpdateFieldOption struct {
	useIncrement bool
}
type CashCollectionSessionsUpdateField struct {
	cashCollectionSessionsField CashCollectionSessionsField
	opt                         CashCollectionSessionsUpdateFieldOption
	value                       interface{}
}
type CashCollectionSessionsUpdateFieldList []CashCollectionSessionsUpdateField

func defaultCashCollectionSessionsUpdateFieldOption() CashCollectionSessionsUpdateFieldOption {
	return CashCollectionSessionsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementCashCollectionSessionsOption(useIncrement bool) func(*CashCollectionSessionsUpdateFieldOption) {
	return func(pcufo *CashCollectionSessionsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewCashCollectionSessionsUpdateField(field CashCollectionSessionsField, val interface{}, opts ...func(*CashCollectionSessionsUpdateFieldOption)) CashCollectionSessionsUpdateField {
	defaultOpt := defaultCashCollectionSessionsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return CashCollectionSessionsUpdateField{
		cashCollectionSessionsField: field,
		value:                       val,
		opt:                         defaultOpt,
	}
}
func defaultCashCollectionSessionsUpdateFields(cashCollectionSessions model.CashCollectionSessions) (cashCollectionSessionsUpdateFieldList CashCollectionSessionsUpdateFieldList) {
	selectFields := NewCashCollectionSessionsSelectFields()
	cashCollectionSessionsUpdateFieldList = append(cashCollectionSessionsUpdateFieldList,
		NewCashCollectionSessionsUpdateField(selectFields.Id(), cashCollectionSessions.Id),
		NewCashCollectionSessionsUpdateField(selectFields.SessionCode(), cashCollectionSessions.SessionCode),
		NewCashCollectionSessionsUpdateField(selectFields.MerchantId(), cashCollectionSessions.MerchantId),
		NewCashCollectionSessionsUpdateField(selectFields.CollectorId(), cashCollectionSessions.CollectorId),
		NewCashCollectionSessionsUpdateField(selectFields.LocationId(), cashCollectionSessions.LocationId),
		NewCashCollectionSessionsUpdateField(selectFields.OpenedAt(), cashCollectionSessions.OpenedAt),
		NewCashCollectionSessionsUpdateField(selectFields.ClosedAt(), cashCollectionSessions.ClosedAt),
		NewCashCollectionSessionsUpdateField(selectFields.Status(), cashCollectionSessions.Status),
		NewCashCollectionSessionsUpdateField(selectFields.OpeningFloatAmount(), cashCollectionSessions.OpeningFloatAmount),
		NewCashCollectionSessionsUpdateField(selectFields.ExpectedAmount(), cashCollectionSessions.ExpectedAmount),
		NewCashCollectionSessionsUpdateField(selectFields.CountedAmount(), cashCollectionSessions.CountedAmount),
		NewCashCollectionSessionsUpdateField(selectFields.VarianceAmount(), cashCollectionSessions.VarianceAmount),
		NewCashCollectionSessionsUpdateField(selectFields.Currency(), cashCollectionSessions.Currency),
		NewCashCollectionSessionsUpdateField(selectFields.Notes(), cashCollectionSessions.Notes),
		NewCashCollectionSessionsUpdateField(selectFields.Metadata(), cashCollectionSessions.Metadata),
		NewCashCollectionSessionsUpdateField(selectFields.MetaCreatedAt(), cashCollectionSessions.MetaCreatedAt),
		NewCashCollectionSessionsUpdateField(selectFields.MetaCreatedBy(), cashCollectionSessions.MetaCreatedBy),
		NewCashCollectionSessionsUpdateField(selectFields.MetaUpdatedAt(), cashCollectionSessions.MetaUpdatedAt),
		NewCashCollectionSessionsUpdateField(selectFields.MetaUpdatedBy(), cashCollectionSessions.MetaUpdatedBy),
		NewCashCollectionSessionsUpdateField(selectFields.MetaDeletedAt(), cashCollectionSessions.MetaDeletedAt),
		NewCashCollectionSessionsUpdateField(selectFields.MetaDeletedBy(), cashCollectionSessions.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsCashCollectionSessionsCommand(cashCollectionSessionsUpdateFieldList CashCollectionSessionsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range cashCollectionSessionsUpdateFieldList {
		field := string(updateField.cashCollectionSessionsField)
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

func (repo *RepositoryImpl) BulkCreateCashCollectionSessions(ctx context.Context, cashCollectionSessionsList []*model.CashCollectionSessions, fieldsInsert ...CashCollectionSessionsField) (err error) {
	var (
		fieldsStr                       string
		valueListStr                    []string
		argsList                        []interface{}
		primaryIds                      []model.CashCollectionSessionsPrimaryID
		cashCollectionSessionsValueList []model.CashCollectionSessions
	)

	if len(fieldsInsert) == 0 {
		selectField := NewCashCollectionSessionsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, cashCollectionSessions := range cashCollectionSessionsList {

		primaryIds = append(primaryIds, cashCollectionSessions.ToCashCollectionSessionsPrimaryID())

		cashCollectionSessionsValueList = append(cashCollectionSessionsValueList, *cashCollectionSessions)
	}

	_, notFoundIds, err := repo.IsExistCashCollectionSessionsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateCashCollectionSessions] failed checking cashCollectionSessions whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.CashCollectionSessionsPrimaryID{}
		mapNotFoundIds := map[model.CashCollectionSessionsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "cashCollectionSessions", fmt.Sprintf("cashCollectionSessions with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsCashCollectionSessions(cashCollectionSessionsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(cashCollectionSessionsQueries.insertCashCollectionSessions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateCashCollectionSessions] failed exec create cashCollectionSessions query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteCashCollectionSessionsByIDs(ctx context.Context, primaryIDs []model.CashCollectionSessionsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistCashCollectionSessionsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCashCollectionSessionsByIDs] failed checking cashCollectionSessions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionSessions with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"cash_collection_sessions\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := cashCollectionSessionsQueries.deleteCashCollectionSessions + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCashCollectionSessionsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCashCollectionSessionsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistCashCollectionSessionsByIDs(ctx context.Context, ids []model.CashCollectionSessionsPrimaryID) (exists bool, notFoundIds []model.CashCollectionSessionsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"cash_collection_sessions\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(cashCollectionSessionsQueries.selectCashCollectionSessions, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCashCollectionSessionsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.CashCollectionSessionsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCashCollectionSessionsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.CashCollectionSessionsPrimaryID]bool{}
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

// BulkUpdateCashCollectionSessions is used to bulk update cashCollectionSessions, by default it will update all field
// if want to update specific field, then fill cashCollectionSessionssMapUpdateFieldsRequest else please fill cashCollectionSessionssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateCashCollectionSessions(ctx context.Context, cashCollectionSessionssMap map[model.CashCollectionSessionsPrimaryID]*model.CashCollectionSessions, cashCollectionSessionssMapUpdateFieldsRequest map[model.CashCollectionSessionsPrimaryID]CashCollectionSessionsUpdateFieldList) (err error) {
	if len(cashCollectionSessionssMap) == 0 && len(cashCollectionSessionssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		cashCollectionSessionssMapUpdateField map[model.CashCollectionSessionsPrimaryID]CashCollectionSessionsUpdateFieldList = map[model.CashCollectionSessionsPrimaryID]CashCollectionSessionsUpdateFieldList{}
		asTableValues                         string                                                                          = "myvalues"
	)

	if len(cashCollectionSessionssMap) > 0 {
		for id, cashCollectionSessions := range cashCollectionSessionssMap {
			if cashCollectionSessions == nil {
				log.Error().Err(err).Msg("[BulkUpdateCashCollectionSessions] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			cashCollectionSessionssMapUpdateField[id] = defaultCashCollectionSessionsUpdateFields(*cashCollectionSessions)
		}
	} else {
		cashCollectionSessionssMapUpdateField = cashCollectionSessionssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateCashCollectionSessionsQuery(cashCollectionSessionssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistCashCollectionSessionsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateCashCollectionSessions] failed checking cashCollectionSessions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionSessions with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeCashCollectionSessionsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"cash_collection_sessions\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateCashCollectionSessions] failed exec query")
	}
	return
}

type CashCollectionSessionsFieldParameter struct {
	param string
	args  []interface{}
}

func NewCashCollectionSessionsFieldParameter(param string, args ...interface{}) CashCollectionSessionsFieldParameter {
	return CashCollectionSessionsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateCashCollectionSessionsQuery(mapCashCollectionSessionss map[model.CashCollectionSessionsPrimaryID]CashCollectionSessionsUpdateFieldList, asTableValues string) (primaryIDs []model.CashCollectionSessionsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.CashCollectionSessionsPrimaryID]map[string]interface{}{}
	cashCollectionSessionsSelectFields := NewCashCollectionSessionsSelectFields()
	for id, updateFields := range mapCashCollectionSessionss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.cashCollectionSessionsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapCashCollectionSessionss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetCashCollectionSessionsFieldType(updateField.cashCollectionSessionsField)))
			args = append(args, fields[string(updateField.cashCollectionSessionsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.cashCollectionSessionsField))
		if updateField.cashCollectionSessionsField == cashCollectionSessionsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.cashCollectionSessionsField, asTableValues, updateField.cashCollectionSessionsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.cashCollectionSessionsField,
				"\"cash_collection_sessions\"", updateField.cashCollectionSessionsField,
				asTableValues, updateField.cashCollectionSessionsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeCashCollectionSessionsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.CashCollectionSessionsPrimaryID, asTableValue string) (whereQry string) {
	cashCollectionSessionsSelectFields := NewCashCollectionSessionsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"cash_collection_sessions\".\"id\" = %s.\"id\"::"+GetCashCollectionSessionsFieldType(cashCollectionSessionsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetCashCollectionSessionsFieldType(cashCollectionSessionsField CashCollectionSessionsField) string {
	selectCashCollectionSessionsFields := NewCashCollectionSessionsSelectFields()
	switch cashCollectionSessionsField {

	case selectCashCollectionSessionsFields.Id():
		return "uuid"

	case selectCashCollectionSessionsFields.SessionCode():
		return "text"

	case selectCashCollectionSessionsFields.MerchantId():
		return "uuid"

	case selectCashCollectionSessionsFields.CollectorId():
		return "uuid"

	case selectCashCollectionSessionsFields.LocationId():
		return "uuid"

	case selectCashCollectionSessionsFields.OpenedAt():
		return "timestamptz"

	case selectCashCollectionSessionsFields.ClosedAt():
		return "timestamptz"

	case selectCashCollectionSessionsFields.Status():
		return "cash_session_status_enum"

	case selectCashCollectionSessionsFields.OpeningFloatAmount():
		return "numeric"

	case selectCashCollectionSessionsFields.ExpectedAmount():
		return "numeric"

	case selectCashCollectionSessionsFields.CountedAmount():
		return "numeric"

	case selectCashCollectionSessionsFields.VarianceAmount():
		return "numeric"

	case selectCashCollectionSessionsFields.Currency():
		return "text"

	case selectCashCollectionSessionsFields.Notes():
		return "text"

	case selectCashCollectionSessionsFields.Metadata():
		return "jsonb"

	case selectCashCollectionSessionsFields.MetaCreatedAt():
		return "timestamptz"

	case selectCashCollectionSessionsFields.MetaCreatedBy():
		return "uuid"

	case selectCashCollectionSessionsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectCashCollectionSessionsFields.MetaUpdatedBy():
		return "uuid"

	case selectCashCollectionSessionsFields.MetaDeletedAt():
		return "timestamptz"

	case selectCashCollectionSessionsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateCashCollectionSessions(ctx context.Context, cashCollectionSessions *model.CashCollectionSessions, fieldsInsert ...CashCollectionSessionsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewCashCollectionSessionsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.CashCollectionSessionsPrimaryID{
		Id: cashCollectionSessions.Id,
	}
	exists, err := repo.IsExistCashCollectionSessionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateCashCollectionSessions] failed checking cashCollectionSessions whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "cashCollectionSessions", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsCashCollectionSessions([]model.CashCollectionSessions{*cashCollectionSessions}, fieldsInsert...)
	commandQuery := fmt.Sprintf(cashCollectionSessionsQueries.insertCashCollectionSessions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateCashCollectionSessions] failed exec create cashCollectionSessions query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteCashCollectionSessionsByID(ctx context.Context, primaryID model.CashCollectionSessionsPrimaryID) (err error) {
	exists, err := repo.IsExistCashCollectionSessionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteCashCollectionSessionsByID] failed checking cashCollectionSessions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionSessions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeCashCollectionSessionsCompositePrimaryKeyWhere([]model.CashCollectionSessionsPrimaryID{primaryID})
	commandQuery := cashCollectionSessionsQueries.deleteCashCollectionSessions + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteCashCollectionSessionsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveCashCollectionSessionsByFilter(ctx context.Context, filter model.Filter) (result []model.CashCollectionSessionsFilterResult, err error) {
	query, args, err := composeCashCollectionSessionsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCashCollectionSessionsByFilter] failed compose cashCollectionSessions filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCashCollectionSessionsByFilter] failed get cashCollectionSessions by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeCashCollectionSessionsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.CashCollectionSessionsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeCashCollectionSessionsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeCashCollectionSessionsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeCashCollectionSessionsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 21 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewCashCollectionSessionsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 21+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["session_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"session_code\"")
			selectedColumns["session_code"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_id\"")
			selectedColumns["merchant_id"] = struct{}{}
		}
		if _, selected := selectedColumns["collector_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"collector_id\"")
			selectedColumns["collector_id"] = struct{}{}
		}
		if _, selected := selectedColumns["location_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"location_id\"")
			selectedColumns["location_id"] = struct{}{}
		}
		if _, selected := selectedColumns["opened_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"opened_at\"")
			selectedColumns["opened_at"] = struct{}{}
		}
		if _, selected := selectedColumns["closed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"closed_at\"")
			selectedColumns["closed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["opening_float_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"opening_float_amount\"")
			selectedColumns["opening_float_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["expected_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"expected_amount\"")
			selectedColumns["expected_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["counted_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"counted_amount\"")
			selectedColumns["counted_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["variance_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"variance_amount\"")
			selectedColumns["variance_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["notes"]; !selected {
			selectColumns = append(selectColumns, "base.\"notes\"")
			selectedColumns["notes"] = struct{}{}
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

type cashCollectionSessionsFilterPlaceholder struct {
	index int
}

func (p *cashCollectionSessionsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeCashCollectionSessionsFilterPredicate(filterField model.FilterField, placeholders *cashCollectionSessionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewCashCollectionSessionsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeCashCollectionSessionsFilterSQLExpr(spec)
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

func composeCashCollectionSessionsFilterGroup(group model.FilterGroup, placeholders *cashCollectionSessionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeCashCollectionSessionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeCashCollectionSessionsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeCashCollectionSessionsFilterWhereQueries(filter model.Filter, placeholders *cashCollectionSessionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeCashCollectionSessionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeCashCollectionSessionsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeCashCollectionSessionsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateCashCollectionSessionsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeCashCollectionSessionsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeCashCollectionSessionsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := cashCollectionSessionsFilterPlaceholder{index: 1}
	whereQueries, err := composeCashCollectionSessionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewCashCollectionSessionsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeCashCollectionSessionsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeCashCollectionSessionsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"cash_collection_sessions\" base%s", strings.Join(selectColumns, ","), composeCashCollectionSessionsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistCashCollectionSessionsByID(ctx context.Context, primaryID model.CashCollectionSessionsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeCashCollectionSessionsCompositePrimaryKeyWhere([]model.CashCollectionSessionsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", cashCollectionSessionsQueries.selectCountCashCollectionSessions, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCashCollectionSessionsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveCashCollectionSessions(ctx context.Context, selectFields ...CashCollectionSessionsField) (cashCollectionSessionsList model.CashCollectionSessionsList, err error) {
	var (
		defaultCashCollectionSessionsSelectFields = defaultCashCollectionSessionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultCashCollectionSessionsSelectFields = composeCashCollectionSessionsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(cashCollectionSessionsQueries.selectCashCollectionSessions, defaultCashCollectionSessionsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &cashCollectionSessionsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCashCollectionSessions] failed get cashCollectionSessions list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveCashCollectionSessionsByID(ctx context.Context, primaryID model.CashCollectionSessionsPrimaryID, selectFields ...CashCollectionSessionsField) (cashCollectionSessions model.CashCollectionSessions, err error) {
	var (
		defaultCashCollectionSessionsSelectFields = defaultCashCollectionSessionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultCashCollectionSessionsSelectFields = composeCashCollectionSessionsSelectFields(selectFields...)
	}
	whereQry, params := composeCashCollectionSessionsCompositePrimaryKeyWhere([]model.CashCollectionSessionsPrimaryID{primaryID})
	query := fmt.Sprintf(cashCollectionSessionsQueries.selectCashCollectionSessions+" WHERE "+whereQry, defaultCashCollectionSessionsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &cashCollectionSessions, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("cashCollectionSessions with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveCashCollectionSessionsByID] failed get cashCollectionSessions")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateCashCollectionSessionsByID(ctx context.Context, primaryID model.CashCollectionSessionsPrimaryID, cashCollectionSessions *model.CashCollectionSessions, cashCollectionSessionsUpdateFields ...CashCollectionSessionsUpdateField) (err error) {
	exists, err := repo.IsExistCashCollectionSessionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionSessions] failed checking cashCollectionSessions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("cashCollectionSessions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if cashCollectionSessions == nil {
		if len(cashCollectionSessionsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateCashCollectionSessionsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		cashCollectionSessions = &model.CashCollectionSessions{}
	}
	var (
		defaultCashCollectionSessionsUpdateFields = defaultCashCollectionSessionsUpdateFields(*cashCollectionSessions)
		tempUpdateField                           CashCollectionSessionsUpdateFieldList
		selectFields                              = NewCashCollectionSessionsSelectFields()
	)
	if len(cashCollectionSessionsUpdateFields) > 0 {
		for _, updateField := range cashCollectionSessionsUpdateFields {
			if updateField.cashCollectionSessionsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultCashCollectionSessionsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeCashCollectionSessionsCompositePrimaryKeyWhere([]model.CashCollectionSessionsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsCashCollectionSessionsCommand(defaultCashCollectionSessionsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(cashCollectionSessionsQueries.updateCashCollectionSessions+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionSessions] error when try to update cashCollectionSessions by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateCashCollectionSessionsByFilter(ctx context.Context, filter model.Filter, cashCollectionSessionsUpdateFields ...CashCollectionSessionsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(cashCollectionSessionsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields CashCollectionSessionsUpdateFieldList
		selectFields = NewCashCollectionSessionsSelectFields()
	)
	for _, updateField := range cashCollectionSessionsUpdateFields {
		if updateField.cashCollectionSessionsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsCashCollectionSessionsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := cashCollectionSessionsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeCashCollectionSessionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"cash_collection_sessions\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionSessionsByFilter] error when try to update cashCollectionSessions by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCashCollectionSessionsByFilter] failed get rows affected")
	}
	return
}

var (
	cashCollectionSessionsQueries = struct {
		selectCashCollectionSessions      string
		selectCountCashCollectionSessions string
		deleteCashCollectionSessions      string
		updateCashCollectionSessions      string
		insertCashCollectionSessions      string
	}{
		selectCashCollectionSessions:      "SELECT %s FROM \"cash_collection_sessions\"",
		selectCountCashCollectionSessions: "SELECT COUNT(\"id\") FROM \"cash_collection_sessions\"",
		deleteCashCollectionSessions:      "DELETE FROM \"cash_collection_sessions\"",
		updateCashCollectionSessions:      "UPDATE \"cash_collection_sessions\" SET %s ",
		insertCashCollectionSessions:      "INSERT INTO \"cash_collection_sessions\" %s VALUES %s",
	}
)

type CashCollectionSessionsRepository interface {
	CreateCashCollectionSessions(ctx context.Context, cashCollectionSessions *model.CashCollectionSessions, fieldsInsert ...CashCollectionSessionsField) error
	BulkCreateCashCollectionSessions(ctx context.Context, cashCollectionSessionsList []*model.CashCollectionSessions, fieldsInsert ...CashCollectionSessionsField) error
	ResolveCashCollectionSessions(ctx context.Context, selectFields ...CashCollectionSessionsField) (model.CashCollectionSessionsList, error)
	ResolveCashCollectionSessionsByID(ctx context.Context, primaryID model.CashCollectionSessionsPrimaryID, selectFields ...CashCollectionSessionsField) (model.CashCollectionSessions, error)
	UpdateCashCollectionSessionsByID(ctx context.Context, id model.CashCollectionSessionsPrimaryID, cashCollectionSessions *model.CashCollectionSessions, cashCollectionSessionsUpdateFields ...CashCollectionSessionsUpdateField) error
	UpdateCashCollectionSessionsByFilter(ctx context.Context, filter model.Filter, cashCollectionSessionsUpdateFields ...CashCollectionSessionsUpdateField) (rowsAffected int64, err error)
	BulkUpdateCashCollectionSessions(ctx context.Context, cashCollectionSessionsListMap map[model.CashCollectionSessionsPrimaryID]*model.CashCollectionSessions, CashCollectionSessionssMapUpdateFieldsRequest map[model.CashCollectionSessionsPrimaryID]CashCollectionSessionsUpdateFieldList) (err error)
	DeleteCashCollectionSessionsByID(ctx context.Context, id model.CashCollectionSessionsPrimaryID) error
	BulkDeleteCashCollectionSessionsByIDs(ctx context.Context, ids []model.CashCollectionSessionsPrimaryID) error
	ResolveCashCollectionSessionsByFilter(ctx context.Context, filter model.Filter) (result []model.CashCollectionSessionsFilterResult, err error)
	IsExistCashCollectionSessionsByIDs(ctx context.Context, ids []model.CashCollectionSessionsPrimaryID) (exists bool, notFoundIds []model.CashCollectionSessionsPrimaryID, err error)
	IsExistCashCollectionSessionsByID(ctx context.Context, id model.CashCollectionSessionsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
