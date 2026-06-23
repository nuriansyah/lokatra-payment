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

func composeInsertFieldsAndParamsSettlementPolicyVersions(settlementPolicyVersionsList []model.SettlementPolicyVersions, fieldsInsert ...SettlementPolicyVersionsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewSettlementPolicyVersionsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, settlementPolicyVersions := range settlementPolicyVersionsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, settlementPolicyVersions.Id)
			case selectField.SettlementPolicyId():
				args = append(args, settlementPolicyVersions.SettlementPolicyId)
			case selectField.VersionNo():
				args = append(args, settlementPolicyVersions.VersionNo)
			case selectField.TriggerType():
				args = append(args, settlementPolicyVersions.TriggerType)
			case selectField.DelayDays():
				args = append(args, settlementPolicyVersions.DelayDays)
			case selectField.MinPayoutAmount():
				args = append(args, settlementPolicyVersions.MinPayoutAmount)
			case selectField.PayoutFrequency():
				args = append(args, settlementPolicyVersions.PayoutFrequency)
			case selectField.ReservePolicyId():
				args = append(args, settlementPolicyVersions.ReservePolicyId)
			case selectField.AutoRelease():
				args = append(args, settlementPolicyVersions.AutoRelease)
			case selectField.Conditions():
				args = append(args, settlementPolicyVersions.Conditions)
			case selectField.IsCurrent():
				args = append(args, settlementPolicyVersions.IsCurrent)
			case selectField.MetaCreatedAt():
				args = append(args, settlementPolicyVersions.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, settlementPolicyVersions.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, settlementPolicyVersions.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, settlementPolicyVersions.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, settlementPolicyVersions.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, settlementPolicyVersions.MetaDeletedBy)

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

func composeSettlementPolicyVersionsCompositePrimaryKeyWhere(primaryIDs []model.SettlementPolicyVersionsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"settlement_policy_versions\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultSettlementPolicyVersionsSelectFields() string {
	fields := NewSettlementPolicyVersionsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeSettlementPolicyVersionsSelectFields(selectFields ...SettlementPolicyVersionsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type SettlementPolicyVersionsField string
type SettlementPolicyVersionsFieldList []SettlementPolicyVersionsField

type SettlementPolicyVersionsSelectFields struct {
}

func (ss SettlementPolicyVersionsSelectFields) Id() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("id")
}

func (ss SettlementPolicyVersionsSelectFields) SettlementPolicyId() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("settlement_policy_id")
}

func (ss SettlementPolicyVersionsSelectFields) VersionNo() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("version_no")
}

func (ss SettlementPolicyVersionsSelectFields) TriggerType() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("trigger_type")
}

func (ss SettlementPolicyVersionsSelectFields) DelayDays() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("delay_days")
}

func (ss SettlementPolicyVersionsSelectFields) MinPayoutAmount() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("min_payout_amount")
}

func (ss SettlementPolicyVersionsSelectFields) PayoutFrequency() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("payout_frequency")
}

func (ss SettlementPolicyVersionsSelectFields) ReservePolicyId() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("reserve_policy_id")
}

func (ss SettlementPolicyVersionsSelectFields) AutoRelease() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("auto_release")
}

func (ss SettlementPolicyVersionsSelectFields) Conditions() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("conditions")
}

func (ss SettlementPolicyVersionsSelectFields) IsCurrent() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("is_current")
}

func (ss SettlementPolicyVersionsSelectFields) MetaCreatedAt() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("meta_created_at")
}

func (ss SettlementPolicyVersionsSelectFields) MetaCreatedBy() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("meta_created_by")
}

func (ss SettlementPolicyVersionsSelectFields) MetaUpdatedAt() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("meta_updated_at")
}

func (ss SettlementPolicyVersionsSelectFields) MetaUpdatedBy() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("meta_updated_by")
}

func (ss SettlementPolicyVersionsSelectFields) MetaDeletedAt() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("meta_deleted_at")
}

func (ss SettlementPolicyVersionsSelectFields) MetaDeletedBy() SettlementPolicyVersionsField {
	return SettlementPolicyVersionsField("meta_deleted_by")
}

func (ss SettlementPolicyVersionsSelectFields) All() SettlementPolicyVersionsFieldList {
	return []SettlementPolicyVersionsField{
		ss.Id(),
		ss.SettlementPolicyId(),
		ss.VersionNo(),
		ss.TriggerType(),
		ss.DelayDays(),
		ss.MinPayoutAmount(),
		ss.PayoutFrequency(),
		ss.ReservePolicyId(),
		ss.AutoRelease(),
		ss.Conditions(),
		ss.IsCurrent(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewSettlementPolicyVersionsSelectFields() SettlementPolicyVersionsSelectFields {
	return SettlementPolicyVersionsSelectFields{}
}

type SettlementPolicyVersionsUpdateFieldOption struct {
	useIncrement bool
}
type SettlementPolicyVersionsUpdateField struct {
	settlementPolicyVersionsField SettlementPolicyVersionsField
	opt                           SettlementPolicyVersionsUpdateFieldOption
	value                         interface{}
}
type SettlementPolicyVersionsUpdateFieldList []SettlementPolicyVersionsUpdateField

func defaultSettlementPolicyVersionsUpdateFieldOption() SettlementPolicyVersionsUpdateFieldOption {
	return SettlementPolicyVersionsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementSettlementPolicyVersionsOption(useIncrement bool) func(*SettlementPolicyVersionsUpdateFieldOption) {
	return func(pcufo *SettlementPolicyVersionsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewSettlementPolicyVersionsUpdateField(field SettlementPolicyVersionsField, val interface{}, opts ...func(*SettlementPolicyVersionsUpdateFieldOption)) SettlementPolicyVersionsUpdateField {
	defaultOpt := defaultSettlementPolicyVersionsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return SettlementPolicyVersionsUpdateField{
		settlementPolicyVersionsField: field,
		value:                         val,
		opt:                           defaultOpt,
	}
}
func defaultSettlementPolicyVersionsUpdateFields(settlementPolicyVersions model.SettlementPolicyVersions) (settlementPolicyVersionsUpdateFieldList SettlementPolicyVersionsUpdateFieldList) {
	selectFields := NewSettlementPolicyVersionsSelectFields()
	settlementPolicyVersionsUpdateFieldList = append(settlementPolicyVersionsUpdateFieldList,
		NewSettlementPolicyVersionsUpdateField(selectFields.Id(), settlementPolicyVersions.Id),
		NewSettlementPolicyVersionsUpdateField(selectFields.SettlementPolicyId(), settlementPolicyVersions.SettlementPolicyId),
		NewSettlementPolicyVersionsUpdateField(selectFields.VersionNo(), settlementPolicyVersions.VersionNo),
		NewSettlementPolicyVersionsUpdateField(selectFields.TriggerType(), settlementPolicyVersions.TriggerType),
		NewSettlementPolicyVersionsUpdateField(selectFields.DelayDays(), settlementPolicyVersions.DelayDays),
		NewSettlementPolicyVersionsUpdateField(selectFields.MinPayoutAmount(), settlementPolicyVersions.MinPayoutAmount),
		NewSettlementPolicyVersionsUpdateField(selectFields.PayoutFrequency(), settlementPolicyVersions.PayoutFrequency),
		NewSettlementPolicyVersionsUpdateField(selectFields.ReservePolicyId(), settlementPolicyVersions.ReservePolicyId),
		NewSettlementPolicyVersionsUpdateField(selectFields.AutoRelease(), settlementPolicyVersions.AutoRelease),
		NewSettlementPolicyVersionsUpdateField(selectFields.Conditions(), settlementPolicyVersions.Conditions),
		NewSettlementPolicyVersionsUpdateField(selectFields.IsCurrent(), settlementPolicyVersions.IsCurrent),
		NewSettlementPolicyVersionsUpdateField(selectFields.MetaCreatedAt(), settlementPolicyVersions.MetaCreatedAt),
		NewSettlementPolicyVersionsUpdateField(selectFields.MetaCreatedBy(), settlementPolicyVersions.MetaCreatedBy),
		NewSettlementPolicyVersionsUpdateField(selectFields.MetaUpdatedAt(), settlementPolicyVersions.MetaUpdatedAt),
		NewSettlementPolicyVersionsUpdateField(selectFields.MetaUpdatedBy(), settlementPolicyVersions.MetaUpdatedBy),
		NewSettlementPolicyVersionsUpdateField(selectFields.MetaDeletedAt(), settlementPolicyVersions.MetaDeletedAt),
		NewSettlementPolicyVersionsUpdateField(selectFields.MetaDeletedBy(), settlementPolicyVersions.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsSettlementPolicyVersionsCommand(settlementPolicyVersionsUpdateFieldList SettlementPolicyVersionsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range settlementPolicyVersionsUpdateFieldList {
		field := string(updateField.settlementPolicyVersionsField)
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

func (repo *RepositoryImpl) BulkCreateSettlementPolicyVersions(ctx context.Context, settlementPolicyVersionsList []*model.SettlementPolicyVersions, fieldsInsert ...SettlementPolicyVersionsField) (err error) {
	var (
		fieldsStr                         string
		valueListStr                      []string
		argsList                          []interface{}
		primaryIds                        []model.SettlementPolicyVersionsPrimaryID
		settlementPolicyVersionsValueList []model.SettlementPolicyVersions
	)

	if len(fieldsInsert) == 0 {
		selectField := NewSettlementPolicyVersionsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, settlementPolicyVersions := range settlementPolicyVersionsList {

		primaryIds = append(primaryIds, settlementPolicyVersions.ToSettlementPolicyVersionsPrimaryID())

		settlementPolicyVersionsValueList = append(settlementPolicyVersionsValueList, *settlementPolicyVersions)
	}

	_, notFoundIds, err := repo.IsExistSettlementPolicyVersionsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementPolicyVersions] failed checking settlementPolicyVersions whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.SettlementPolicyVersionsPrimaryID{}
		mapNotFoundIds := map[model.SettlementPolicyVersionsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "settlementPolicyVersions", fmt.Sprintf("settlementPolicyVersions with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsSettlementPolicyVersions(settlementPolicyVersionsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(settlementPolicyVersionsQueries.insertSettlementPolicyVersions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementPolicyVersions] failed exec create settlementPolicyVersions query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteSettlementPolicyVersionsByIDs(ctx context.Context, primaryIDs []model.SettlementPolicyVersionsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistSettlementPolicyVersionsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementPolicyVersionsByIDs] failed checking settlementPolicyVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicyVersions with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_policy_versions\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(settlementPolicyVersionsQueries.deleteSettlementPolicyVersions + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementPolicyVersionsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementPolicyVersionsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistSettlementPolicyVersionsByIDs(ctx context.Context, ids []model.SettlementPolicyVersionsPrimaryID) (exists bool, notFoundIds []model.SettlementPolicyVersionsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_policy_versions\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(settlementPolicyVersionsQueries.selectSettlementPolicyVersions, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementPolicyVersionsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.SettlementPolicyVersionsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementPolicyVersionsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.SettlementPolicyVersionsPrimaryID]bool{}
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

// BulkUpdateSettlementPolicyVersions is used to bulk update settlementPolicyVersions, by default it will update all field
// if want to update specific field, then fill settlementPolicyVersionssMapUpdateFieldsRequest else please fill settlementPolicyVersionssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateSettlementPolicyVersions(ctx context.Context, settlementPolicyVersionssMap map[model.SettlementPolicyVersionsPrimaryID]*model.SettlementPolicyVersions, settlementPolicyVersionssMapUpdateFieldsRequest map[model.SettlementPolicyVersionsPrimaryID]SettlementPolicyVersionsUpdateFieldList) (err error) {
	if len(settlementPolicyVersionssMap) == 0 && len(settlementPolicyVersionssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		settlementPolicyVersionssMapUpdateField map[model.SettlementPolicyVersionsPrimaryID]SettlementPolicyVersionsUpdateFieldList = map[model.SettlementPolicyVersionsPrimaryID]SettlementPolicyVersionsUpdateFieldList{}
		asTableValues                           string                                                                              = "myvalues"
	)

	if len(settlementPolicyVersionssMap) > 0 {
		for id, settlementPolicyVersions := range settlementPolicyVersionssMap {
			if settlementPolicyVersions == nil {
				log.Error().Err(err).Msg("[BulkUpdateSettlementPolicyVersions] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			settlementPolicyVersionssMapUpdateField[id] = defaultSettlementPolicyVersionsUpdateFields(*settlementPolicyVersions)
		}
	} else {
		settlementPolicyVersionssMapUpdateField = settlementPolicyVersionssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateSettlementPolicyVersionsQuery(settlementPolicyVersionssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistSettlementPolicyVersionsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementPolicyVersions] failed checking settlementPolicyVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicyVersions with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeSettlementPolicyVersionsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"settlement_policy_versions\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementPolicyVersions] failed exec query")
	}
	return
}

type SettlementPolicyVersionsFieldParameter struct {
	param string
	args  []interface{}
}

func NewSettlementPolicyVersionsFieldParameter(param string, args ...interface{}) SettlementPolicyVersionsFieldParameter {
	return SettlementPolicyVersionsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateSettlementPolicyVersionsQuery(mapSettlementPolicyVersionss map[model.SettlementPolicyVersionsPrimaryID]SettlementPolicyVersionsUpdateFieldList, asTableValues string) (primaryIDs []model.SettlementPolicyVersionsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.SettlementPolicyVersionsPrimaryID]map[string]interface{}{}
	settlementPolicyVersionsSelectFields := NewSettlementPolicyVersionsSelectFields()
	for id, updateFields := range mapSettlementPolicyVersionss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.settlementPolicyVersionsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapSettlementPolicyVersionss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetSettlementPolicyVersionsFieldType(updateField.settlementPolicyVersionsField)))
			args = append(args, fields[string(updateField.settlementPolicyVersionsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.settlementPolicyVersionsField))
		if updateField.settlementPolicyVersionsField == settlementPolicyVersionsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.settlementPolicyVersionsField, asTableValues, updateField.settlementPolicyVersionsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.settlementPolicyVersionsField,
				"\"settlement_policy_versions\"", updateField.settlementPolicyVersionsField,
				asTableValues, updateField.settlementPolicyVersionsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeSettlementPolicyVersionsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.SettlementPolicyVersionsPrimaryID, asTableValue string) (whereQry string) {
	settlementPolicyVersionsSelectFields := NewSettlementPolicyVersionsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"settlement_policy_versions\".\"id\" = %s.\"id\"::"+GetSettlementPolicyVersionsFieldType(settlementPolicyVersionsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetSettlementPolicyVersionsFieldType(settlementPolicyVersionsField SettlementPolicyVersionsField) string {
	selectSettlementPolicyVersionsFields := NewSettlementPolicyVersionsSelectFields()
	switch settlementPolicyVersionsField {

	case selectSettlementPolicyVersionsFields.Id():
		return "uuid"

	case selectSettlementPolicyVersionsFields.SettlementPolicyId():
		return "uuid"

	case selectSettlementPolicyVersionsFields.VersionNo():
		return "int4"

	case selectSettlementPolicyVersionsFields.TriggerType():
		return "trigger_type_enum"

	case selectSettlementPolicyVersionsFields.DelayDays():
		return "int4"

	case selectSettlementPolicyVersionsFields.MinPayoutAmount():
		return "numeric"

	case selectSettlementPolicyVersionsFields.PayoutFrequency():
		return "payout_frequency_enum"

	case selectSettlementPolicyVersionsFields.ReservePolicyId():
		return "uuid"

	case selectSettlementPolicyVersionsFields.AutoRelease():
		return "bool"

	case selectSettlementPolicyVersionsFields.Conditions():
		return "jsonb"

	case selectSettlementPolicyVersionsFields.IsCurrent():
		return "bool"

	case selectSettlementPolicyVersionsFields.MetaCreatedAt():
		return "timestamptz"

	case selectSettlementPolicyVersionsFields.MetaCreatedBy():
		return "uuid"

	case selectSettlementPolicyVersionsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectSettlementPolicyVersionsFields.MetaUpdatedBy():
		return "uuid"

	case selectSettlementPolicyVersionsFields.MetaDeletedAt():
		return "timestamptz"

	case selectSettlementPolicyVersionsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateSettlementPolicyVersions(ctx context.Context, settlementPolicyVersions *model.SettlementPolicyVersions, fieldsInsert ...SettlementPolicyVersionsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewSettlementPolicyVersionsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.SettlementPolicyVersionsPrimaryID{
		Id: settlementPolicyVersions.Id,
	}
	exists, err := repo.IsExistSettlementPolicyVersionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementPolicyVersions] failed checking settlementPolicyVersions whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "settlementPolicyVersions", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsSettlementPolicyVersions([]model.SettlementPolicyVersions{*settlementPolicyVersions}, fieldsInsert...)
	commandQuery := fmt.Sprintf(settlementPolicyVersionsQueries.insertSettlementPolicyVersions, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementPolicyVersions] failed exec create settlementPolicyVersions query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteSettlementPolicyVersionsByID(ctx context.Context, primaryID model.SettlementPolicyVersionsPrimaryID) (err error) {
	exists, err := repo.IsExistSettlementPolicyVersionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementPolicyVersionsByID] failed checking settlementPolicyVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicyVersions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeSettlementPolicyVersionsCompositePrimaryKeyWhere([]model.SettlementPolicyVersionsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(settlementPolicyVersionsQueries.deleteSettlementPolicyVersions + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementPolicyVersionsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementPolicyVersionsByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementPolicyVersionsFilterResult, err error) {
	query, args, err := composeSettlementPolicyVersionsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementPolicyVersionsByFilter] failed compose settlementPolicyVersions filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementPolicyVersionsByFilter] failed get settlementPolicyVersions by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeSettlementPolicyVersionsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.SettlementPolicyVersionsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeSettlementPolicyVersionsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeSettlementPolicyVersionsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeSettlementPolicyVersionsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 17 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewSettlementPolicyVersionsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 17+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["settlement_policy_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"settlement_policy_id\"")
			selectedColumns["settlement_policy_id"] = struct{}{}
		}
		if _, selected := selectedColumns["version_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"version_no\"")
			selectedColumns["version_no"] = struct{}{}
		}
		if _, selected := selectedColumns["trigger_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"trigger_type\"")
			selectedColumns["trigger_type"] = struct{}{}
		}
		if _, selected := selectedColumns["delay_days"]; !selected {
			selectColumns = append(selectColumns, "base.\"delay_days\"")
			selectedColumns["delay_days"] = struct{}{}
		}
		if _, selected := selectedColumns["min_payout_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"min_payout_amount\"")
			selectedColumns["min_payout_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["payout_frequency"]; !selected {
			selectColumns = append(selectColumns, "base.\"payout_frequency\"")
			selectedColumns["payout_frequency"] = struct{}{}
		}
		if _, selected := selectedColumns["reserve_policy_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"reserve_policy_id\"")
			selectedColumns["reserve_policy_id"] = struct{}{}
		}
		if _, selected := selectedColumns["auto_release"]; !selected {
			selectColumns = append(selectColumns, "base.\"auto_release\"")
			selectedColumns["auto_release"] = struct{}{}
		}
		if _, selected := selectedColumns["conditions"]; !selected {
			selectColumns = append(selectColumns, "base.\"conditions\"")
			selectedColumns["conditions"] = struct{}{}
		}
		if _, selected := selectedColumns["is_current"]; !selected {
			selectColumns = append(selectColumns, "base.\"is_current\"")
			selectedColumns["is_current"] = struct{}{}
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

type settlementPolicyVersionsFilterPlaceholder struct {
	index int
}

func (p *settlementPolicyVersionsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeSettlementPolicyVersionsFilterPredicate(filterField model.FilterField, placeholders *settlementPolicyVersionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewSettlementPolicyVersionsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeSettlementPolicyVersionsFilterSQLExpr(spec)
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

func composeSettlementPolicyVersionsFilterGroup(group model.FilterGroup, placeholders *settlementPolicyVersionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeSettlementPolicyVersionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeSettlementPolicyVersionsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeSettlementPolicyVersionsFilterWhereQueries(filter model.Filter, placeholders *settlementPolicyVersionsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeSettlementPolicyVersionsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeSettlementPolicyVersionsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeSettlementPolicyVersionsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateSettlementPolicyVersionsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeSettlementPolicyVersionsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeSettlementPolicyVersionsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := settlementPolicyVersionsFilterPlaceholder{index: 1}
	whereQueries, err := composeSettlementPolicyVersionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewSettlementPolicyVersionsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeSettlementPolicyVersionsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeSettlementPolicyVersionsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"settlement_policy_versions\" base%s", strings.Join(selectColumns, ","), composeSettlementPolicyVersionsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistSettlementPolicyVersionsByID(ctx context.Context, primaryID model.SettlementPolicyVersionsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeSettlementPolicyVersionsCompositePrimaryKeyWhere([]model.SettlementPolicyVersionsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", settlementPolicyVersionsQueries.selectCountSettlementPolicyVersions, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementPolicyVersionsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementPolicyVersions(ctx context.Context, selectFields ...SettlementPolicyVersionsField) (settlementPolicyVersionsList model.SettlementPolicyVersionsList, err error) {
	var (
		defaultSettlementPolicyVersionsSelectFields = defaultSettlementPolicyVersionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementPolicyVersionsSelectFields = composeSettlementPolicyVersionsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(settlementPolicyVersionsQueries.selectSettlementPolicyVersions, defaultSettlementPolicyVersionsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &settlementPolicyVersionsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementPolicyVersions] failed get settlementPolicyVersions list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementPolicyVersionsByID(ctx context.Context, primaryID model.SettlementPolicyVersionsPrimaryID, selectFields ...SettlementPolicyVersionsField) (settlementPolicyVersions model.SettlementPolicyVersions, err error) {
	var (
		defaultSettlementPolicyVersionsSelectFields = defaultSettlementPolicyVersionsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementPolicyVersionsSelectFields = composeSettlementPolicyVersionsSelectFields(selectFields...)
	}
	whereQry, params := composeSettlementPolicyVersionsCompositePrimaryKeyWhere([]model.SettlementPolicyVersionsPrimaryID{primaryID})
	query := fmt.Sprintf(settlementPolicyVersionsQueries.selectSettlementPolicyVersions+" WHERE "+whereQry, defaultSettlementPolicyVersionsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &settlementPolicyVersions, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("settlementPolicyVersions with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveSettlementPolicyVersionsByID] failed get settlementPolicyVersions")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateSettlementPolicyVersionsByID(ctx context.Context, primaryID model.SettlementPolicyVersionsPrimaryID, settlementPolicyVersions *model.SettlementPolicyVersions, settlementPolicyVersionsUpdateFields ...SettlementPolicyVersionsUpdateField) (err error) {
	exists, err := repo.IsExistSettlementPolicyVersionsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPolicyVersions] failed checking settlementPolicyVersions whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementPolicyVersions with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if settlementPolicyVersions == nil {
		if len(settlementPolicyVersionsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateSettlementPolicyVersionsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		settlementPolicyVersions = &model.SettlementPolicyVersions{}
	}
	var (
		defaultSettlementPolicyVersionsUpdateFields = defaultSettlementPolicyVersionsUpdateFields(*settlementPolicyVersions)
		tempUpdateField                             SettlementPolicyVersionsUpdateFieldList
		selectFields                                = NewSettlementPolicyVersionsSelectFields()
	)
	if len(settlementPolicyVersionsUpdateFields) > 0 {
		for _, updateField := range settlementPolicyVersionsUpdateFields {
			if updateField.settlementPolicyVersionsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultSettlementPolicyVersionsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeSettlementPolicyVersionsCompositePrimaryKeyWhere([]model.SettlementPolicyVersionsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsSettlementPolicyVersionsCommand(defaultSettlementPolicyVersionsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(settlementPolicyVersionsQueries.updateSettlementPolicyVersions+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPolicyVersions] error when try to update settlementPolicyVersions by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateSettlementPolicyVersionsByFilter(ctx context.Context, filter model.Filter, settlementPolicyVersionsUpdateFields ...SettlementPolicyVersionsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(settlementPolicyVersionsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields SettlementPolicyVersionsUpdateFieldList
		selectFields = NewSettlementPolicyVersionsSelectFields()
	)
	for _, updateField := range settlementPolicyVersionsUpdateFields {
		if updateField.settlementPolicyVersionsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsSettlementPolicyVersionsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := settlementPolicyVersionsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeSettlementPolicyVersionsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"settlement_policy_versions\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPolicyVersionsByFilter] error when try to update settlementPolicyVersions by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementPolicyVersionsByFilter] failed get rows affected")
	}
	return
}

var (
	settlementPolicyVersionsQueries = struct {
		selectSettlementPolicyVersions      string
		selectCountSettlementPolicyVersions string
		deleteSettlementPolicyVersions      string
		updateSettlementPolicyVersions      string
		insertSettlementPolicyVersions      string
	}{
		selectSettlementPolicyVersions:      "SELECT %s FROM \"settlement_policy_versions\"",
		selectCountSettlementPolicyVersions: "SELECT COUNT(\"id\") FROM \"settlement_policy_versions\"",
		deleteSettlementPolicyVersions:      "DELETE FROM \"settlement_policy_versions\"",
		updateSettlementPolicyVersions:      "UPDATE \"settlement_policy_versions\" SET %s ",
		insertSettlementPolicyVersions:      "INSERT INTO \"settlement_policy_versions\" %s VALUES %s",
	}
)

type SettlementPolicyVersionsRepository interface {
	CreateSettlementPolicyVersions(ctx context.Context, settlementPolicyVersions *model.SettlementPolicyVersions, fieldsInsert ...SettlementPolicyVersionsField) error
	BulkCreateSettlementPolicyVersions(ctx context.Context, settlementPolicyVersionsList []*model.SettlementPolicyVersions, fieldsInsert ...SettlementPolicyVersionsField) error
	ResolveSettlementPolicyVersions(ctx context.Context, selectFields ...SettlementPolicyVersionsField) (model.SettlementPolicyVersionsList, error)
	ResolveSettlementPolicyVersionsByID(ctx context.Context, primaryID model.SettlementPolicyVersionsPrimaryID, selectFields ...SettlementPolicyVersionsField) (model.SettlementPolicyVersions, error)
	UpdateSettlementPolicyVersionsByID(ctx context.Context, id model.SettlementPolicyVersionsPrimaryID, settlementPolicyVersions *model.SettlementPolicyVersions, settlementPolicyVersionsUpdateFields ...SettlementPolicyVersionsUpdateField) error
	UpdateSettlementPolicyVersionsByFilter(ctx context.Context, filter model.Filter, settlementPolicyVersionsUpdateFields ...SettlementPolicyVersionsUpdateField) (rowsAffected int64, err error)
	BulkUpdateSettlementPolicyVersions(ctx context.Context, settlementPolicyVersionsListMap map[model.SettlementPolicyVersionsPrimaryID]*model.SettlementPolicyVersions, SettlementPolicyVersionssMapUpdateFieldsRequest map[model.SettlementPolicyVersionsPrimaryID]SettlementPolicyVersionsUpdateFieldList) (err error)
	DeleteSettlementPolicyVersionsByID(ctx context.Context, id model.SettlementPolicyVersionsPrimaryID) error
	BulkDeleteSettlementPolicyVersionsByIDs(ctx context.Context, ids []model.SettlementPolicyVersionsPrimaryID) error
	ResolveSettlementPolicyVersionsByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementPolicyVersionsFilterResult, err error)
	IsExistSettlementPolicyVersionsByIDs(ctx context.Context, ids []model.SettlementPolicyVersionsPrimaryID) (exists bool, notFoundIds []model.SettlementPolicyVersionsPrimaryID, err error)
	IsExistSettlementPolicyVersionsByID(ctx context.Context, id model.SettlementPolicyVersionsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
