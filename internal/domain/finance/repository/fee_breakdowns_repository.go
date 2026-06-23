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

func composeInsertFieldsAndParamsFeeBreakdowns(feeBreakdownsList []model.FeeBreakdowns, fieldsInsert ...FeeBreakdownsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewFeeBreakdownsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, feeBreakdowns := range feeBreakdownsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, feeBreakdowns.Id)
			case selectField.SourceType():
				args = append(args, feeBreakdowns.SourceType)
			case selectField.SourceId():
				args = append(args, feeBreakdowns.SourceId)
			case selectField.FeeProfileId():
				args = append(args, feeBreakdowns.FeeProfileId)
			case selectField.FeeRuleId():
				args = append(args, feeBreakdowns.FeeRuleId)
			case selectField.CurrencyCode():
				args = append(args, feeBreakdowns.CurrencyCode)
			case selectField.BaseAmount():
				args = append(args, feeBreakdowns.BaseAmount)
			case selectField.FeeAmount():
				args = append(args, feeBreakdowns.FeeAmount)
			case selectField.TaxAmount():
				args = append(args, feeBreakdowns.TaxAmount)
			case selectField.RecipientPartyId():
				args = append(args, feeBreakdowns.RecipientPartyId)
			case selectField.BreakdownStatus():
				args = append(args, feeBreakdowns.BreakdownStatus)
			case selectField.Metadata():
				args = append(args, feeBreakdowns.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, feeBreakdowns.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, feeBreakdowns.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, feeBreakdowns.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, feeBreakdowns.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, feeBreakdowns.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, feeBreakdowns.MetaDeletedBy)

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

func composeFeeBreakdownsCompositePrimaryKeyWhere(primaryIDs []model.FeeBreakdownsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"fee_breakdowns\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultFeeBreakdownsSelectFields() string {
	fields := NewFeeBreakdownsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeFeeBreakdownsSelectFields(selectFields ...FeeBreakdownsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type FeeBreakdownsField string
type FeeBreakdownsFieldList []FeeBreakdownsField

type FeeBreakdownsSelectFields struct {
}

func (ss FeeBreakdownsSelectFields) Id() FeeBreakdownsField {
	return FeeBreakdownsField("id")
}

func (ss FeeBreakdownsSelectFields) SourceType() FeeBreakdownsField {
	return FeeBreakdownsField("source_type")
}

func (ss FeeBreakdownsSelectFields) SourceId() FeeBreakdownsField {
	return FeeBreakdownsField("source_id")
}

func (ss FeeBreakdownsSelectFields) FeeProfileId() FeeBreakdownsField {
	return FeeBreakdownsField("fee_profile_id")
}

func (ss FeeBreakdownsSelectFields) FeeRuleId() FeeBreakdownsField {
	return FeeBreakdownsField("fee_rule_id")
}

func (ss FeeBreakdownsSelectFields) CurrencyCode() FeeBreakdownsField {
	return FeeBreakdownsField("currency_code")
}

func (ss FeeBreakdownsSelectFields) BaseAmount() FeeBreakdownsField {
	return FeeBreakdownsField("base_amount")
}

func (ss FeeBreakdownsSelectFields) FeeAmount() FeeBreakdownsField {
	return FeeBreakdownsField("fee_amount")
}

func (ss FeeBreakdownsSelectFields) TaxAmount() FeeBreakdownsField {
	return FeeBreakdownsField("tax_amount")
}

func (ss FeeBreakdownsSelectFields) RecipientPartyId() FeeBreakdownsField {
	return FeeBreakdownsField("recipient_party_id")
}

func (ss FeeBreakdownsSelectFields) BreakdownStatus() FeeBreakdownsField {
	return FeeBreakdownsField("breakdown_status")
}

func (ss FeeBreakdownsSelectFields) Metadata() FeeBreakdownsField {
	return FeeBreakdownsField("metadata")
}

func (ss FeeBreakdownsSelectFields) MetaCreatedAt() FeeBreakdownsField {
	return FeeBreakdownsField("meta_created_at")
}

func (ss FeeBreakdownsSelectFields) MetaCreatedBy() FeeBreakdownsField {
	return FeeBreakdownsField("meta_created_by")
}

func (ss FeeBreakdownsSelectFields) MetaUpdatedAt() FeeBreakdownsField {
	return FeeBreakdownsField("meta_updated_at")
}

func (ss FeeBreakdownsSelectFields) MetaUpdatedBy() FeeBreakdownsField {
	return FeeBreakdownsField("meta_updated_by")
}

func (ss FeeBreakdownsSelectFields) MetaDeletedAt() FeeBreakdownsField {
	return FeeBreakdownsField("meta_deleted_at")
}

func (ss FeeBreakdownsSelectFields) MetaDeletedBy() FeeBreakdownsField {
	return FeeBreakdownsField("meta_deleted_by")
}

func (ss FeeBreakdownsSelectFields) All() FeeBreakdownsFieldList {
	return []FeeBreakdownsField{
		ss.Id(),
		ss.SourceType(),
		ss.SourceId(),
		ss.FeeProfileId(),
		ss.FeeRuleId(),
		ss.CurrencyCode(),
		ss.BaseAmount(),
		ss.FeeAmount(),
		ss.TaxAmount(),
		ss.RecipientPartyId(),
		ss.BreakdownStatus(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewFeeBreakdownsSelectFields() FeeBreakdownsSelectFields {
	return FeeBreakdownsSelectFields{}
}

type FeeBreakdownsUpdateFieldOption struct {
	useIncrement bool
}
type FeeBreakdownsUpdateField struct {
	feeBreakdownsField FeeBreakdownsField
	opt                FeeBreakdownsUpdateFieldOption
	value              interface{}
}
type FeeBreakdownsUpdateFieldList []FeeBreakdownsUpdateField

func defaultFeeBreakdownsUpdateFieldOption() FeeBreakdownsUpdateFieldOption {
	return FeeBreakdownsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementFeeBreakdownsOption(useIncrement bool) func(*FeeBreakdownsUpdateFieldOption) {
	return func(pcufo *FeeBreakdownsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewFeeBreakdownsUpdateField(field FeeBreakdownsField, val interface{}, opts ...func(*FeeBreakdownsUpdateFieldOption)) FeeBreakdownsUpdateField {
	defaultOpt := defaultFeeBreakdownsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return FeeBreakdownsUpdateField{
		feeBreakdownsField: field,
		value:              val,
		opt:                defaultOpt,
	}
}
func defaultFeeBreakdownsUpdateFields(feeBreakdowns model.FeeBreakdowns) (feeBreakdownsUpdateFieldList FeeBreakdownsUpdateFieldList) {
	selectFields := NewFeeBreakdownsSelectFields()
	feeBreakdownsUpdateFieldList = append(feeBreakdownsUpdateFieldList,
		NewFeeBreakdownsUpdateField(selectFields.Id(), feeBreakdowns.Id),
		NewFeeBreakdownsUpdateField(selectFields.SourceType(), feeBreakdowns.SourceType),
		NewFeeBreakdownsUpdateField(selectFields.SourceId(), feeBreakdowns.SourceId),
		NewFeeBreakdownsUpdateField(selectFields.FeeProfileId(), feeBreakdowns.FeeProfileId),
		NewFeeBreakdownsUpdateField(selectFields.FeeRuleId(), feeBreakdowns.FeeRuleId),
		NewFeeBreakdownsUpdateField(selectFields.CurrencyCode(), feeBreakdowns.CurrencyCode),
		NewFeeBreakdownsUpdateField(selectFields.BaseAmount(), feeBreakdowns.BaseAmount),
		NewFeeBreakdownsUpdateField(selectFields.FeeAmount(), feeBreakdowns.FeeAmount),
		NewFeeBreakdownsUpdateField(selectFields.TaxAmount(), feeBreakdowns.TaxAmount),
		NewFeeBreakdownsUpdateField(selectFields.RecipientPartyId(), feeBreakdowns.RecipientPartyId),
		NewFeeBreakdownsUpdateField(selectFields.BreakdownStatus(), feeBreakdowns.BreakdownStatus),
		NewFeeBreakdownsUpdateField(selectFields.Metadata(), feeBreakdowns.Metadata),
		NewFeeBreakdownsUpdateField(selectFields.MetaCreatedAt(), feeBreakdowns.MetaCreatedAt),
		NewFeeBreakdownsUpdateField(selectFields.MetaCreatedBy(), feeBreakdowns.MetaCreatedBy),
		NewFeeBreakdownsUpdateField(selectFields.MetaUpdatedAt(), feeBreakdowns.MetaUpdatedAt),
		NewFeeBreakdownsUpdateField(selectFields.MetaUpdatedBy(), feeBreakdowns.MetaUpdatedBy),
		NewFeeBreakdownsUpdateField(selectFields.MetaDeletedAt(), feeBreakdowns.MetaDeletedAt),
		NewFeeBreakdownsUpdateField(selectFields.MetaDeletedBy(), feeBreakdowns.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsFeeBreakdownsCommand(feeBreakdownsUpdateFieldList FeeBreakdownsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range feeBreakdownsUpdateFieldList {
		field := string(updateField.feeBreakdownsField)
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

func (repo *RepositoryImpl) BulkCreateFeeBreakdowns(ctx context.Context, feeBreakdownsList []*model.FeeBreakdowns, fieldsInsert ...FeeBreakdownsField) (err error) {
	var (
		fieldsStr              string
		valueListStr           []string
		argsList               []interface{}
		primaryIds             []model.FeeBreakdownsPrimaryID
		feeBreakdownsValueList []model.FeeBreakdowns
	)

	if len(fieldsInsert) == 0 {
		selectField := NewFeeBreakdownsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, feeBreakdowns := range feeBreakdownsList {

		primaryIds = append(primaryIds, feeBreakdowns.ToFeeBreakdownsPrimaryID())

		feeBreakdownsValueList = append(feeBreakdownsValueList, *feeBreakdowns)
	}

	_, notFoundIds, err := repo.IsExistFeeBreakdownsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeBreakdowns] failed checking feeBreakdowns whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.FeeBreakdownsPrimaryID{}
		mapNotFoundIds := map[model.FeeBreakdownsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "feeBreakdowns", fmt.Sprintf("feeBreakdowns with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsFeeBreakdowns(feeBreakdownsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(feeBreakdownsQueries.insertFeeBreakdowns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateFeeBreakdowns] failed exec create feeBreakdowns query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteFeeBreakdownsByIDs(ctx context.Context, primaryIDs []model.FeeBreakdownsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistFeeBreakdownsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeBreakdownsByIDs] failed checking feeBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeBreakdowns with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_breakdowns\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(feeBreakdownsQueries.deleteFeeBreakdowns + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeBreakdownsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteFeeBreakdownsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistFeeBreakdownsByIDs(ctx context.Context, ids []model.FeeBreakdownsPrimaryID) (exists bool, notFoundIds []model.FeeBreakdownsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"fee_breakdowns\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(feeBreakdownsQueries.selectFeeBreakdowns, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeBreakdownsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.FeeBreakdownsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeBreakdownsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.FeeBreakdownsPrimaryID]bool{}
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

// BulkUpdateFeeBreakdowns is used to bulk update feeBreakdowns, by default it will update all field
// if want to update specific field, then fill feeBreakdownssMapUpdateFieldsRequest else please fill feeBreakdownssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateFeeBreakdowns(ctx context.Context, feeBreakdownssMap map[model.FeeBreakdownsPrimaryID]*model.FeeBreakdowns, feeBreakdownssMapUpdateFieldsRequest map[model.FeeBreakdownsPrimaryID]FeeBreakdownsUpdateFieldList) (err error) {
	if len(feeBreakdownssMap) == 0 && len(feeBreakdownssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		feeBreakdownssMapUpdateField map[model.FeeBreakdownsPrimaryID]FeeBreakdownsUpdateFieldList = map[model.FeeBreakdownsPrimaryID]FeeBreakdownsUpdateFieldList{}
		asTableValues                string                                                        = "myvalues"
	)

	if len(feeBreakdownssMap) > 0 {
		for id, feeBreakdowns := range feeBreakdownssMap {
			if feeBreakdowns == nil {
				log.Error().Err(err).Msg("[BulkUpdateFeeBreakdowns] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			feeBreakdownssMapUpdateField[id] = defaultFeeBreakdownsUpdateFields(*feeBreakdowns)
		}
	} else {
		feeBreakdownssMapUpdateField = feeBreakdownssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateFeeBreakdownsQuery(feeBreakdownssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistFeeBreakdownsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeBreakdowns] failed checking feeBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeBreakdowns with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeFeeBreakdownsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"fee_breakdowns\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateFeeBreakdowns] failed exec query")
	}
	return
}

type FeeBreakdownsFieldParameter struct {
	param string
	args  []interface{}
}

func NewFeeBreakdownsFieldParameter(param string, args ...interface{}) FeeBreakdownsFieldParameter {
	return FeeBreakdownsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateFeeBreakdownsQuery(mapFeeBreakdownss map[model.FeeBreakdownsPrimaryID]FeeBreakdownsUpdateFieldList, asTableValues string) (primaryIDs []model.FeeBreakdownsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.FeeBreakdownsPrimaryID]map[string]interface{}{}
	feeBreakdownsSelectFields := NewFeeBreakdownsSelectFields()
	for id, updateFields := range mapFeeBreakdownss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.feeBreakdownsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapFeeBreakdownss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetFeeBreakdownsFieldType(updateField.feeBreakdownsField)))
			args = append(args, fields[string(updateField.feeBreakdownsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.feeBreakdownsField))
		if updateField.feeBreakdownsField == feeBreakdownsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.feeBreakdownsField, asTableValues, updateField.feeBreakdownsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.feeBreakdownsField,
				"\"fee_breakdowns\"", updateField.feeBreakdownsField,
				asTableValues, updateField.feeBreakdownsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeFeeBreakdownsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.FeeBreakdownsPrimaryID, asTableValue string) (whereQry string) {
	feeBreakdownsSelectFields := NewFeeBreakdownsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"fee_breakdowns\".\"id\" = %s.\"id\"::"+GetFeeBreakdownsFieldType(feeBreakdownsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetFeeBreakdownsFieldType(feeBreakdownsField FeeBreakdownsField) string {
	selectFeeBreakdownsFields := NewFeeBreakdownsSelectFields()
	switch feeBreakdownsField {

	case selectFeeBreakdownsFields.Id():
		return "uuid"

	case selectFeeBreakdownsFields.SourceType():
		return "text"

	case selectFeeBreakdownsFields.SourceId():
		return "uuid"

	case selectFeeBreakdownsFields.FeeProfileId():
		return "uuid"

	case selectFeeBreakdownsFields.FeeRuleId():
		return "uuid"

	case selectFeeBreakdownsFields.CurrencyCode():
		return "text"

	case selectFeeBreakdownsFields.BaseAmount():
		return "numeric"

	case selectFeeBreakdownsFields.FeeAmount():
		return "numeric"

	case selectFeeBreakdownsFields.TaxAmount():
		return "numeric"

	case selectFeeBreakdownsFields.RecipientPartyId():
		return "uuid"

	case selectFeeBreakdownsFields.BreakdownStatus():
		return "fee_breakdowns_breakdown_status_enum"

	case selectFeeBreakdownsFields.Metadata():
		return "jsonb"

	case selectFeeBreakdownsFields.MetaCreatedAt():
		return "timestamptz"

	case selectFeeBreakdownsFields.MetaCreatedBy():
		return "uuid"

	case selectFeeBreakdownsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectFeeBreakdownsFields.MetaUpdatedBy():
		return "uuid"

	case selectFeeBreakdownsFields.MetaDeletedAt():
		return "timestamptz"

	case selectFeeBreakdownsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateFeeBreakdowns(ctx context.Context, feeBreakdowns *model.FeeBreakdowns, fieldsInsert ...FeeBreakdownsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewFeeBreakdownsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.FeeBreakdownsPrimaryID{
		Id: feeBreakdowns.Id,
	}
	exists, err := repo.IsExistFeeBreakdownsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeBreakdowns] failed checking feeBreakdowns whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "feeBreakdowns", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsFeeBreakdowns([]model.FeeBreakdowns{*feeBreakdowns}, fieldsInsert...)
	commandQuery := fmt.Sprintf(feeBreakdownsQueries.insertFeeBreakdowns, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateFeeBreakdowns] failed exec create feeBreakdowns query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteFeeBreakdownsByID(ctx context.Context, primaryID model.FeeBreakdownsPrimaryID) (err error) {
	exists, err := repo.IsExistFeeBreakdownsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeBreakdownsByID] failed checking feeBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeBreakdowns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeFeeBreakdownsCompositePrimaryKeyWhere([]model.FeeBreakdownsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(feeBreakdownsQueries.deleteFeeBreakdowns + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteFeeBreakdownsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeBreakdownsByFilter(ctx context.Context, filter model.Filter) (result []model.FeeBreakdownsFilterResult, err error) {
	query, args, err := composeFeeBreakdownsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeBreakdownsByFilter] failed compose feeBreakdowns filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeBreakdownsByFilter] failed get feeBreakdowns by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeFeeBreakdownsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.FeeBreakdownsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeFeeBreakdownsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeFeeBreakdownsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeFeeBreakdownsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewFeeBreakdownsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_profile_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_profile_id\"")
			selectedColumns["fee_profile_id"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_rule_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_rule_id\"")
			selectedColumns["fee_rule_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["base_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"base_amount\"")
			selectedColumns["base_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["fee_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"fee_amount\"")
			selectedColumns["fee_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_amount\"")
			selectedColumns["tax_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["recipient_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"recipient_party_id\"")
			selectedColumns["recipient_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["breakdown_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"breakdown_status\"")
			selectedColumns["breakdown_status"] = struct{}{}
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

type feeBreakdownsFilterPlaceholder struct {
	index int
}

func (p *feeBreakdownsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeFeeBreakdownsFilterPredicate(filterField model.FilterField, placeholders *feeBreakdownsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewFeeBreakdownsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeFeeBreakdownsFilterSQLExpr(spec)
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

func composeFeeBreakdownsFilterGroup(group model.FilterGroup, placeholders *feeBreakdownsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeFeeBreakdownsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeFeeBreakdownsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeFeeBreakdownsFilterWhereQueries(filter model.Filter, placeholders *feeBreakdownsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeFeeBreakdownsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeFeeBreakdownsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeFeeBreakdownsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateFeeBreakdownsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeFeeBreakdownsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeFeeBreakdownsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := feeBreakdownsFilterPlaceholder{index: 1}
	whereQueries, err := composeFeeBreakdownsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewFeeBreakdownsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeFeeBreakdownsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeFeeBreakdownsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"fee_breakdowns\" base%s", strings.Join(selectColumns, ","), composeFeeBreakdownsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistFeeBreakdownsByID(ctx context.Context, primaryID model.FeeBreakdownsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeFeeBreakdownsCompositePrimaryKeyWhere([]model.FeeBreakdownsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", feeBreakdownsQueries.selectCountFeeBreakdowns, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistFeeBreakdownsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeBreakdowns(ctx context.Context, selectFields ...FeeBreakdownsField) (feeBreakdownsList model.FeeBreakdownsList, err error) {
	var (
		defaultFeeBreakdownsSelectFields = defaultFeeBreakdownsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeBreakdownsSelectFields = composeFeeBreakdownsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(feeBreakdownsQueries.selectFeeBreakdowns, defaultFeeBreakdownsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &feeBreakdownsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveFeeBreakdowns] failed get feeBreakdowns list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveFeeBreakdownsByID(ctx context.Context, primaryID model.FeeBreakdownsPrimaryID, selectFields ...FeeBreakdownsField) (feeBreakdowns model.FeeBreakdowns, err error) {
	var (
		defaultFeeBreakdownsSelectFields = defaultFeeBreakdownsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultFeeBreakdownsSelectFields = composeFeeBreakdownsSelectFields(selectFields...)
	}
	whereQry, params := composeFeeBreakdownsCompositePrimaryKeyWhere([]model.FeeBreakdownsPrimaryID{primaryID})
	query := fmt.Sprintf(feeBreakdownsQueries.selectFeeBreakdowns+" WHERE "+whereQry, defaultFeeBreakdownsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &feeBreakdowns, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("feeBreakdowns with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveFeeBreakdownsByID] failed get feeBreakdowns")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateFeeBreakdownsByID(ctx context.Context, primaryID model.FeeBreakdownsPrimaryID, feeBreakdowns *model.FeeBreakdowns, feeBreakdownsUpdateFields ...FeeBreakdownsUpdateField) (err error) {
	exists, err := repo.IsExistFeeBreakdownsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeBreakdowns] failed checking feeBreakdowns whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("feeBreakdowns with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if feeBreakdowns == nil {
		if len(feeBreakdownsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateFeeBreakdownsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		feeBreakdowns = &model.FeeBreakdowns{}
	}
	var (
		defaultFeeBreakdownsUpdateFields = defaultFeeBreakdownsUpdateFields(*feeBreakdowns)
		tempUpdateField                  FeeBreakdownsUpdateFieldList
		selectFields                     = NewFeeBreakdownsSelectFields()
	)
	if len(feeBreakdownsUpdateFields) > 0 {
		for _, updateField := range feeBreakdownsUpdateFields {
			if updateField.feeBreakdownsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultFeeBreakdownsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeFeeBreakdownsCompositePrimaryKeyWhere([]model.FeeBreakdownsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsFeeBreakdownsCommand(defaultFeeBreakdownsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(feeBreakdownsQueries.updateFeeBreakdowns+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeBreakdowns] error when try to update feeBreakdowns by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateFeeBreakdownsByFilter(ctx context.Context, filter model.Filter, feeBreakdownsUpdateFields ...FeeBreakdownsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(feeBreakdownsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields FeeBreakdownsUpdateFieldList
		selectFields = NewFeeBreakdownsSelectFields()
	)
	for _, updateField := range feeBreakdownsUpdateFields {
		if updateField.feeBreakdownsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsFeeBreakdownsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := feeBreakdownsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeFeeBreakdownsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"fee_breakdowns\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeBreakdownsByFilter] error when try to update feeBreakdowns by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateFeeBreakdownsByFilter] failed get rows affected")
	}
	return
}

var (
	feeBreakdownsQueries = struct {
		selectFeeBreakdowns      string
		selectCountFeeBreakdowns string
		deleteFeeBreakdowns      string
		updateFeeBreakdowns      string
		insertFeeBreakdowns      string
	}{
		selectFeeBreakdowns:      "SELECT %s FROM \"fee_breakdowns\"",
		selectCountFeeBreakdowns: "SELECT COUNT(\"id\") FROM \"fee_breakdowns\"",
		deleteFeeBreakdowns:      "DELETE FROM \"fee_breakdowns\"",
		updateFeeBreakdowns:      "UPDATE \"fee_breakdowns\" SET %s ",
		insertFeeBreakdowns:      "INSERT INTO \"fee_breakdowns\" %s VALUES %s",
	}
)

type FeeBreakdownsRepository interface {
	CreateFeeBreakdowns(ctx context.Context, feeBreakdowns *model.FeeBreakdowns, fieldsInsert ...FeeBreakdownsField) error
	BulkCreateFeeBreakdowns(ctx context.Context, feeBreakdownsList []*model.FeeBreakdowns, fieldsInsert ...FeeBreakdownsField) error
	ResolveFeeBreakdowns(ctx context.Context, selectFields ...FeeBreakdownsField) (model.FeeBreakdownsList, error)
	ResolveFeeBreakdownsByID(ctx context.Context, primaryID model.FeeBreakdownsPrimaryID, selectFields ...FeeBreakdownsField) (model.FeeBreakdowns, error)
	UpdateFeeBreakdownsByID(ctx context.Context, id model.FeeBreakdownsPrimaryID, feeBreakdowns *model.FeeBreakdowns, feeBreakdownsUpdateFields ...FeeBreakdownsUpdateField) error
	UpdateFeeBreakdownsByFilter(ctx context.Context, filter model.Filter, feeBreakdownsUpdateFields ...FeeBreakdownsUpdateField) (rowsAffected int64, err error)
	BulkUpdateFeeBreakdowns(ctx context.Context, feeBreakdownsListMap map[model.FeeBreakdownsPrimaryID]*model.FeeBreakdowns, FeeBreakdownssMapUpdateFieldsRequest map[model.FeeBreakdownsPrimaryID]FeeBreakdownsUpdateFieldList) (err error)
	DeleteFeeBreakdownsByID(ctx context.Context, id model.FeeBreakdownsPrimaryID) error
	BulkDeleteFeeBreakdownsByIDs(ctx context.Context, ids []model.FeeBreakdownsPrimaryID) error
	ResolveFeeBreakdownsByFilter(ctx context.Context, filter model.Filter) (result []model.FeeBreakdownsFilterResult, err error)
	IsExistFeeBreakdownsByIDs(ctx context.Context, ids []model.FeeBreakdownsPrimaryID) (exists bool, notFoundIds []model.FeeBreakdownsPrimaryID, err error)
	IsExistFeeBreakdownsByID(ctx context.Context, id model.FeeBreakdownsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
