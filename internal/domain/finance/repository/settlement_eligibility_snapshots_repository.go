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

func composeInsertFieldsAndParamsSettlementEligibilitySnapshots(settlementEligibilitySnapshotsList []model.SettlementEligibilitySnapshots, fieldsInsert ...SettlementEligibilitySnapshotsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewSettlementEligibilitySnapshotsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, settlementEligibilitySnapshots := range settlementEligibilitySnapshotsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, settlementEligibilitySnapshots.Id)
			case selectField.SourceType():
				args = append(args, settlementEligibilitySnapshots.SourceType)
			case selectField.SourceId():
				args = append(args, settlementEligibilitySnapshots.SourceId)
			case selectField.MerchantPartyId():
				args = append(args, settlementEligibilitySnapshots.MerchantPartyId)
			case selectField.SettlementPolicyVersionId():
				args = append(args, settlementEligibilitySnapshots.SettlementPolicyVersionId)
			case selectField.CurrencyCode():
				args = append(args, settlementEligibilitySnapshots.CurrencyCode)
			case selectField.GrossAmount():
				args = append(args, settlementEligibilitySnapshots.GrossAmount)
			case selectField.FeeAmount():
				args = append(args, settlementEligibilitySnapshots.FeeAmount)
			case selectField.TaxAmount():
				args = append(args, settlementEligibilitySnapshots.TaxAmount)
			case selectField.ReserveAmount():
				args = append(args, settlementEligibilitySnapshots.ReserveAmount)
			case selectField.NetSettleableAmount():
				args = append(args, settlementEligibilitySnapshots.NetSettleableAmount)
			case selectField.EligibilityStatus():
				args = append(args, settlementEligibilitySnapshots.EligibilityStatus)
			case selectField.EligibleAt():
				args = append(args, settlementEligibilitySnapshots.EligibleAt)
			case selectField.SnapshotPayload():
				args = append(args, settlementEligibilitySnapshots.SnapshotPayload)
			case selectField.MetaCreatedAt():
				args = append(args, settlementEligibilitySnapshots.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, settlementEligibilitySnapshots.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, settlementEligibilitySnapshots.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, settlementEligibilitySnapshots.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, settlementEligibilitySnapshots.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, settlementEligibilitySnapshots.MetaDeletedBy)

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

func composeSettlementEligibilitySnapshotsCompositePrimaryKeyWhere(primaryIDs []model.SettlementEligibilitySnapshotsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"settlement_eligibility_snapshots\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultSettlementEligibilitySnapshotsSelectFields() string {
	fields := NewSettlementEligibilitySnapshotsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeSettlementEligibilitySnapshotsSelectFields(selectFields ...SettlementEligibilitySnapshotsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type SettlementEligibilitySnapshotsField string
type SettlementEligibilitySnapshotsFieldList []SettlementEligibilitySnapshotsField

type SettlementEligibilitySnapshotsSelectFields struct {
}

func (ss SettlementEligibilitySnapshotsSelectFields) Id() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("id")
}

func (ss SettlementEligibilitySnapshotsSelectFields) SourceType() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("source_type")
}

func (ss SettlementEligibilitySnapshotsSelectFields) SourceId() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("source_id")
}

func (ss SettlementEligibilitySnapshotsSelectFields) MerchantPartyId() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("merchant_party_id")
}

func (ss SettlementEligibilitySnapshotsSelectFields) SettlementPolicyVersionId() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("settlement_policy_version_id")
}

func (ss SettlementEligibilitySnapshotsSelectFields) CurrencyCode() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("currency_code")
}

func (ss SettlementEligibilitySnapshotsSelectFields) GrossAmount() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("gross_amount")
}

func (ss SettlementEligibilitySnapshotsSelectFields) FeeAmount() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("fee_amount")
}

func (ss SettlementEligibilitySnapshotsSelectFields) TaxAmount() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("tax_amount")
}

func (ss SettlementEligibilitySnapshotsSelectFields) ReserveAmount() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("reserve_amount")
}

func (ss SettlementEligibilitySnapshotsSelectFields) NetSettleableAmount() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("net_settleable_amount")
}

func (ss SettlementEligibilitySnapshotsSelectFields) EligibilityStatus() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("eligibility_status")
}

func (ss SettlementEligibilitySnapshotsSelectFields) EligibleAt() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("eligible_at")
}

func (ss SettlementEligibilitySnapshotsSelectFields) SnapshotPayload() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("snapshot_payload")
}

func (ss SettlementEligibilitySnapshotsSelectFields) MetaCreatedAt() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("meta_created_at")
}

func (ss SettlementEligibilitySnapshotsSelectFields) MetaCreatedBy() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("meta_created_by")
}

func (ss SettlementEligibilitySnapshotsSelectFields) MetaUpdatedAt() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("meta_updated_at")
}

func (ss SettlementEligibilitySnapshotsSelectFields) MetaUpdatedBy() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("meta_updated_by")
}

func (ss SettlementEligibilitySnapshotsSelectFields) MetaDeletedAt() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("meta_deleted_at")
}

func (ss SettlementEligibilitySnapshotsSelectFields) MetaDeletedBy() SettlementEligibilitySnapshotsField {
	return SettlementEligibilitySnapshotsField("meta_deleted_by")
}

func (ss SettlementEligibilitySnapshotsSelectFields) All() SettlementEligibilitySnapshotsFieldList {
	return []SettlementEligibilitySnapshotsField{
		ss.Id(),
		ss.SourceType(),
		ss.SourceId(),
		ss.MerchantPartyId(),
		ss.SettlementPolicyVersionId(),
		ss.CurrencyCode(),
		ss.GrossAmount(),
		ss.FeeAmount(),
		ss.TaxAmount(),
		ss.ReserveAmount(),
		ss.NetSettleableAmount(),
		ss.EligibilityStatus(),
		ss.EligibleAt(),
		ss.SnapshotPayload(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewSettlementEligibilitySnapshotsSelectFields() SettlementEligibilitySnapshotsSelectFields {
	return SettlementEligibilitySnapshotsSelectFields{}
}

type SettlementEligibilitySnapshotsUpdateFieldOption struct {
	useIncrement bool
}
type SettlementEligibilitySnapshotsUpdateField struct {
	settlementEligibilitySnapshotsField SettlementEligibilitySnapshotsField
	opt                                 SettlementEligibilitySnapshotsUpdateFieldOption
	value                               interface{}
}
type SettlementEligibilitySnapshotsUpdateFieldList []SettlementEligibilitySnapshotsUpdateField

func defaultSettlementEligibilitySnapshotsUpdateFieldOption() SettlementEligibilitySnapshotsUpdateFieldOption {
	return SettlementEligibilitySnapshotsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementSettlementEligibilitySnapshotsOption(useIncrement bool) func(*SettlementEligibilitySnapshotsUpdateFieldOption) {
	return func(pcufo *SettlementEligibilitySnapshotsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewSettlementEligibilitySnapshotsUpdateField(field SettlementEligibilitySnapshotsField, val interface{}, opts ...func(*SettlementEligibilitySnapshotsUpdateFieldOption)) SettlementEligibilitySnapshotsUpdateField {
	defaultOpt := defaultSettlementEligibilitySnapshotsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return SettlementEligibilitySnapshotsUpdateField{
		settlementEligibilitySnapshotsField: field,
		value:                               val,
		opt:                                 defaultOpt,
	}
}
func defaultSettlementEligibilitySnapshotsUpdateFields(settlementEligibilitySnapshots model.SettlementEligibilitySnapshots) (settlementEligibilitySnapshotsUpdateFieldList SettlementEligibilitySnapshotsUpdateFieldList) {
	selectFields := NewSettlementEligibilitySnapshotsSelectFields()
	settlementEligibilitySnapshotsUpdateFieldList = append(settlementEligibilitySnapshotsUpdateFieldList,
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.Id(), settlementEligibilitySnapshots.Id),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.SourceType(), settlementEligibilitySnapshots.SourceType),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.SourceId(), settlementEligibilitySnapshots.SourceId),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.MerchantPartyId(), settlementEligibilitySnapshots.MerchantPartyId),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.SettlementPolicyVersionId(), settlementEligibilitySnapshots.SettlementPolicyVersionId),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.CurrencyCode(), settlementEligibilitySnapshots.CurrencyCode),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.GrossAmount(), settlementEligibilitySnapshots.GrossAmount),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.FeeAmount(), settlementEligibilitySnapshots.FeeAmount),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.TaxAmount(), settlementEligibilitySnapshots.TaxAmount),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.ReserveAmount(), settlementEligibilitySnapshots.ReserveAmount),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.NetSettleableAmount(), settlementEligibilitySnapshots.NetSettleableAmount),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.EligibilityStatus(), settlementEligibilitySnapshots.EligibilityStatus),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.EligibleAt(), settlementEligibilitySnapshots.EligibleAt),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.SnapshotPayload(), settlementEligibilitySnapshots.SnapshotPayload),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.MetaCreatedAt(), settlementEligibilitySnapshots.MetaCreatedAt),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.MetaCreatedBy(), settlementEligibilitySnapshots.MetaCreatedBy),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.MetaUpdatedAt(), settlementEligibilitySnapshots.MetaUpdatedAt),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.MetaUpdatedBy(), settlementEligibilitySnapshots.MetaUpdatedBy),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.MetaDeletedAt(), settlementEligibilitySnapshots.MetaDeletedAt),
		NewSettlementEligibilitySnapshotsUpdateField(selectFields.MetaDeletedBy(), settlementEligibilitySnapshots.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsSettlementEligibilitySnapshotsCommand(settlementEligibilitySnapshotsUpdateFieldList SettlementEligibilitySnapshotsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range settlementEligibilitySnapshotsUpdateFieldList {
		field := string(updateField.settlementEligibilitySnapshotsField)
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

func (repo *RepositoryImpl) BulkCreateSettlementEligibilitySnapshots(ctx context.Context, settlementEligibilitySnapshotsList []*model.SettlementEligibilitySnapshots, fieldsInsert ...SettlementEligibilitySnapshotsField) (err error) {
	var (
		fieldsStr                               string
		valueListStr                            []string
		argsList                                []interface{}
		primaryIds                              []model.SettlementEligibilitySnapshotsPrimaryID
		settlementEligibilitySnapshotsValueList []model.SettlementEligibilitySnapshots
	)

	if len(fieldsInsert) == 0 {
		selectField := NewSettlementEligibilitySnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, settlementEligibilitySnapshots := range settlementEligibilitySnapshotsList {

		primaryIds = append(primaryIds, settlementEligibilitySnapshots.ToSettlementEligibilitySnapshotsPrimaryID())

		settlementEligibilitySnapshotsValueList = append(settlementEligibilitySnapshotsValueList, *settlementEligibilitySnapshots)
	}

	_, notFoundIds, err := repo.IsExistSettlementEligibilitySnapshotsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementEligibilitySnapshots] failed checking settlementEligibilitySnapshots whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.SettlementEligibilitySnapshotsPrimaryID{}
		mapNotFoundIds := map[model.SettlementEligibilitySnapshotsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "settlementEligibilitySnapshots", fmt.Sprintf("settlementEligibilitySnapshots with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsSettlementEligibilitySnapshots(settlementEligibilitySnapshotsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(settlementEligibilitySnapshotsQueries.insertSettlementEligibilitySnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateSettlementEligibilitySnapshots] failed exec create settlementEligibilitySnapshots query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteSettlementEligibilitySnapshotsByIDs(ctx context.Context, primaryIDs []model.SettlementEligibilitySnapshotsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistSettlementEligibilitySnapshotsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementEligibilitySnapshotsByIDs] failed checking settlementEligibilitySnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementEligibilitySnapshots with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_eligibility_snapshots\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(settlementEligibilitySnapshotsQueries.deleteSettlementEligibilitySnapshots + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementEligibilitySnapshotsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteSettlementEligibilitySnapshotsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistSettlementEligibilitySnapshotsByIDs(ctx context.Context, ids []model.SettlementEligibilitySnapshotsPrimaryID) (exists bool, notFoundIds []model.SettlementEligibilitySnapshotsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"settlement_eligibility_snapshots\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(settlementEligibilitySnapshotsQueries.selectSettlementEligibilitySnapshots, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementEligibilitySnapshotsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.SettlementEligibilitySnapshotsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementEligibilitySnapshotsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.SettlementEligibilitySnapshotsPrimaryID]bool{}
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

// BulkUpdateSettlementEligibilitySnapshots is used to bulk update settlementEligibilitySnapshots, by default it will update all field
// if want to update specific field, then fill settlementEligibilitySnapshotssMapUpdateFieldsRequest else please fill settlementEligibilitySnapshotssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateSettlementEligibilitySnapshots(ctx context.Context, settlementEligibilitySnapshotssMap map[model.SettlementEligibilitySnapshotsPrimaryID]*model.SettlementEligibilitySnapshots, settlementEligibilitySnapshotssMapUpdateFieldsRequest map[model.SettlementEligibilitySnapshotsPrimaryID]SettlementEligibilitySnapshotsUpdateFieldList) (err error) {
	if len(settlementEligibilitySnapshotssMap) == 0 && len(settlementEligibilitySnapshotssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		settlementEligibilitySnapshotssMapUpdateField map[model.SettlementEligibilitySnapshotsPrimaryID]SettlementEligibilitySnapshotsUpdateFieldList = map[model.SettlementEligibilitySnapshotsPrimaryID]SettlementEligibilitySnapshotsUpdateFieldList{}
		asTableValues                                 string                                                                                          = "myvalues"
	)

	if len(settlementEligibilitySnapshotssMap) > 0 {
		for id, settlementEligibilitySnapshots := range settlementEligibilitySnapshotssMap {
			if settlementEligibilitySnapshots == nil {
				log.Error().Err(err).Msg("[BulkUpdateSettlementEligibilitySnapshots] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			settlementEligibilitySnapshotssMapUpdateField[id] = defaultSettlementEligibilitySnapshotsUpdateFields(*settlementEligibilitySnapshots)
		}
	} else {
		settlementEligibilitySnapshotssMapUpdateField = settlementEligibilitySnapshotssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateSettlementEligibilitySnapshotsQuery(settlementEligibilitySnapshotssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistSettlementEligibilitySnapshotsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementEligibilitySnapshots] failed checking settlementEligibilitySnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementEligibilitySnapshots with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeSettlementEligibilitySnapshotsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"settlement_eligibility_snapshots\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateSettlementEligibilitySnapshots] failed exec query")
	}
	return
}

type SettlementEligibilitySnapshotsFieldParameter struct {
	param string
	args  []interface{}
}

func NewSettlementEligibilitySnapshotsFieldParameter(param string, args ...interface{}) SettlementEligibilitySnapshotsFieldParameter {
	return SettlementEligibilitySnapshotsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateSettlementEligibilitySnapshotsQuery(mapSettlementEligibilitySnapshotss map[model.SettlementEligibilitySnapshotsPrimaryID]SettlementEligibilitySnapshotsUpdateFieldList, asTableValues string) (primaryIDs []model.SettlementEligibilitySnapshotsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.SettlementEligibilitySnapshotsPrimaryID]map[string]interface{}{}
	settlementEligibilitySnapshotsSelectFields := NewSettlementEligibilitySnapshotsSelectFields()
	for id, updateFields := range mapSettlementEligibilitySnapshotss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.settlementEligibilitySnapshotsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapSettlementEligibilitySnapshotss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetSettlementEligibilitySnapshotsFieldType(updateField.settlementEligibilitySnapshotsField)))
			args = append(args, fields[string(updateField.settlementEligibilitySnapshotsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.settlementEligibilitySnapshotsField))
		if updateField.settlementEligibilitySnapshotsField == settlementEligibilitySnapshotsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.settlementEligibilitySnapshotsField, asTableValues, updateField.settlementEligibilitySnapshotsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.settlementEligibilitySnapshotsField,
				"\"settlement_eligibility_snapshots\"", updateField.settlementEligibilitySnapshotsField,
				asTableValues, updateField.settlementEligibilitySnapshotsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeSettlementEligibilitySnapshotsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.SettlementEligibilitySnapshotsPrimaryID, asTableValue string) (whereQry string) {
	settlementEligibilitySnapshotsSelectFields := NewSettlementEligibilitySnapshotsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"settlement_eligibility_snapshots\".\"id\" = %s.\"id\"::"+GetSettlementEligibilitySnapshotsFieldType(settlementEligibilitySnapshotsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetSettlementEligibilitySnapshotsFieldType(settlementEligibilitySnapshotsField SettlementEligibilitySnapshotsField) string {
	selectSettlementEligibilitySnapshotsFields := NewSettlementEligibilitySnapshotsSelectFields()
	switch settlementEligibilitySnapshotsField {

	case selectSettlementEligibilitySnapshotsFields.Id():
		return "uuid"

	case selectSettlementEligibilitySnapshotsFields.SourceType():
		return "text"

	case selectSettlementEligibilitySnapshotsFields.SourceId():
		return "uuid"

	case selectSettlementEligibilitySnapshotsFields.MerchantPartyId():
		return "uuid"

	case selectSettlementEligibilitySnapshotsFields.SettlementPolicyVersionId():
		return "uuid"

	case selectSettlementEligibilitySnapshotsFields.CurrencyCode():
		return "text"

	case selectSettlementEligibilitySnapshotsFields.GrossAmount():
		return "numeric"

	case selectSettlementEligibilitySnapshotsFields.FeeAmount():
		return "numeric"

	case selectSettlementEligibilitySnapshotsFields.TaxAmount():
		return "numeric"

	case selectSettlementEligibilitySnapshotsFields.ReserveAmount():
		return "numeric"

	case selectSettlementEligibilitySnapshotsFields.NetSettleableAmount():
		return "numeric"

	case selectSettlementEligibilitySnapshotsFields.EligibilityStatus():
		return "eligibility_status_enum"

	case selectSettlementEligibilitySnapshotsFields.EligibleAt():
		return "timestamptz"

	case selectSettlementEligibilitySnapshotsFields.SnapshotPayload():
		return "jsonb"

	case selectSettlementEligibilitySnapshotsFields.MetaCreatedAt():
		return "timestamptz"

	case selectSettlementEligibilitySnapshotsFields.MetaCreatedBy():
		return "uuid"

	case selectSettlementEligibilitySnapshotsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectSettlementEligibilitySnapshotsFields.MetaUpdatedBy():
		return "uuid"

	case selectSettlementEligibilitySnapshotsFields.MetaDeletedAt():
		return "timestamptz"

	case selectSettlementEligibilitySnapshotsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateSettlementEligibilitySnapshots(ctx context.Context, settlementEligibilitySnapshots *model.SettlementEligibilitySnapshots, fieldsInsert ...SettlementEligibilitySnapshotsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewSettlementEligibilitySnapshotsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.SettlementEligibilitySnapshotsPrimaryID{
		Id: settlementEligibilitySnapshots.Id,
	}
	exists, err := repo.IsExistSettlementEligibilitySnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementEligibilitySnapshots] failed checking settlementEligibilitySnapshots whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "settlementEligibilitySnapshots", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsSettlementEligibilitySnapshots([]model.SettlementEligibilitySnapshots{*settlementEligibilitySnapshots}, fieldsInsert...)
	commandQuery := fmt.Sprintf(settlementEligibilitySnapshotsQueries.insertSettlementEligibilitySnapshots, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateSettlementEligibilitySnapshots] failed exec create settlementEligibilitySnapshots query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteSettlementEligibilitySnapshotsByID(ctx context.Context, primaryID model.SettlementEligibilitySnapshotsPrimaryID) (err error) {
	exists, err := repo.IsExistSettlementEligibilitySnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementEligibilitySnapshotsByID] failed checking settlementEligibilitySnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementEligibilitySnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeSettlementEligibilitySnapshotsCompositePrimaryKeyWhere([]model.SettlementEligibilitySnapshotsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(settlementEligibilitySnapshotsQueries.deleteSettlementEligibilitySnapshots + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteSettlementEligibilitySnapshotsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementEligibilitySnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementEligibilitySnapshotsFilterResult, err error) {
	query, args, err := composeSettlementEligibilitySnapshotsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementEligibilitySnapshotsByFilter] failed compose settlementEligibilitySnapshots filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementEligibilitySnapshotsByFilter] failed get settlementEligibilitySnapshots by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeSettlementEligibilitySnapshotsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.SettlementEligibilitySnapshotsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeSettlementEligibilitySnapshotsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeSettlementEligibilitySnapshotsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeSettlementEligibilitySnapshotsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 20+1)
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
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["settlement_policy_version_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"settlement_policy_version_id\"")
			selectedColumns["settlement_policy_version_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
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
		if _, selected := selectedColumns["net_settleable_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"net_settleable_amount\"")
			selectedColumns["net_settleable_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["eligibility_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"eligibility_status\"")
			selectedColumns["eligibility_status"] = struct{}{}
		}
		if _, selected := selectedColumns["eligible_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"eligible_at\"")
			selectedColumns["eligible_at"] = struct{}{}
		}
		if _, selected := selectedColumns["snapshot_payload"]; !selected {
			selectColumns = append(selectColumns, "base.\"snapshot_payload\"")
			selectedColumns["snapshot_payload"] = struct{}{}
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

type settlementEligibilitySnapshotsFilterPlaceholder struct {
	index int
}

func (p *settlementEligibilitySnapshotsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeSettlementEligibilitySnapshotsFilterPredicate(filterField model.FilterField, placeholders *settlementEligibilitySnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeSettlementEligibilitySnapshotsFilterSQLExpr(spec)
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

func composeSettlementEligibilitySnapshotsFilterGroup(group model.FilterGroup, placeholders *settlementEligibilitySnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeSettlementEligibilitySnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeSettlementEligibilitySnapshotsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeSettlementEligibilitySnapshotsFilterWhereQueries(filter model.Filter, placeholders *settlementEligibilitySnapshotsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeSettlementEligibilitySnapshotsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeSettlementEligibilitySnapshotsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeSettlementEligibilitySnapshotsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateSettlementEligibilitySnapshotsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeSettlementEligibilitySnapshotsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeSettlementEligibilitySnapshotsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := settlementEligibilitySnapshotsFilterPlaceholder{index: 1}
	whereQueries, err := composeSettlementEligibilitySnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewSettlementEligibilitySnapshotsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeSettlementEligibilitySnapshotsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeSettlementEligibilitySnapshotsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"settlement_eligibility_snapshots\" base%s", strings.Join(selectColumns, ","), composeSettlementEligibilitySnapshotsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistSettlementEligibilitySnapshotsByID(ctx context.Context, primaryID model.SettlementEligibilitySnapshotsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeSettlementEligibilitySnapshotsCompositePrimaryKeyWhere([]model.SettlementEligibilitySnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", settlementEligibilitySnapshotsQueries.selectCountSettlementEligibilitySnapshots, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistSettlementEligibilitySnapshotsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementEligibilitySnapshots(ctx context.Context, selectFields ...SettlementEligibilitySnapshotsField) (settlementEligibilitySnapshotsList model.SettlementEligibilitySnapshotsList, err error) {
	var (
		defaultSettlementEligibilitySnapshotsSelectFields = defaultSettlementEligibilitySnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementEligibilitySnapshotsSelectFields = composeSettlementEligibilitySnapshotsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(settlementEligibilitySnapshotsQueries.selectSettlementEligibilitySnapshots, defaultSettlementEligibilitySnapshotsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &settlementEligibilitySnapshotsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveSettlementEligibilitySnapshots] failed get settlementEligibilitySnapshots list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveSettlementEligibilitySnapshotsByID(ctx context.Context, primaryID model.SettlementEligibilitySnapshotsPrimaryID, selectFields ...SettlementEligibilitySnapshotsField) (settlementEligibilitySnapshots model.SettlementEligibilitySnapshots, err error) {
	var (
		defaultSettlementEligibilitySnapshotsSelectFields = defaultSettlementEligibilitySnapshotsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultSettlementEligibilitySnapshotsSelectFields = composeSettlementEligibilitySnapshotsSelectFields(selectFields...)
	}
	whereQry, params := composeSettlementEligibilitySnapshotsCompositePrimaryKeyWhere([]model.SettlementEligibilitySnapshotsPrimaryID{primaryID})
	query := fmt.Sprintf(settlementEligibilitySnapshotsQueries.selectSettlementEligibilitySnapshots+" WHERE "+whereQry, defaultSettlementEligibilitySnapshotsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &settlementEligibilitySnapshots, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("settlementEligibilitySnapshots with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveSettlementEligibilitySnapshotsByID] failed get settlementEligibilitySnapshots")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateSettlementEligibilitySnapshotsByID(ctx context.Context, primaryID model.SettlementEligibilitySnapshotsPrimaryID, settlementEligibilitySnapshots *model.SettlementEligibilitySnapshots, settlementEligibilitySnapshotsUpdateFields ...SettlementEligibilitySnapshotsUpdateField) (err error) {
	exists, err := repo.IsExistSettlementEligibilitySnapshotsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementEligibilitySnapshots] failed checking settlementEligibilitySnapshots whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("settlementEligibilitySnapshots with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if settlementEligibilitySnapshots == nil {
		if len(settlementEligibilitySnapshotsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateSettlementEligibilitySnapshotsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		settlementEligibilitySnapshots = &model.SettlementEligibilitySnapshots{}
	}
	var (
		defaultSettlementEligibilitySnapshotsUpdateFields = defaultSettlementEligibilitySnapshotsUpdateFields(*settlementEligibilitySnapshots)
		tempUpdateField                                   SettlementEligibilitySnapshotsUpdateFieldList
		selectFields                                      = NewSettlementEligibilitySnapshotsSelectFields()
	)
	if len(settlementEligibilitySnapshotsUpdateFields) > 0 {
		for _, updateField := range settlementEligibilitySnapshotsUpdateFields {
			if updateField.settlementEligibilitySnapshotsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultSettlementEligibilitySnapshotsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeSettlementEligibilitySnapshotsCompositePrimaryKeyWhere([]model.SettlementEligibilitySnapshotsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsSettlementEligibilitySnapshotsCommand(defaultSettlementEligibilitySnapshotsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(settlementEligibilitySnapshotsQueries.updateSettlementEligibilitySnapshots+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementEligibilitySnapshots] error when try to update settlementEligibilitySnapshots by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateSettlementEligibilitySnapshotsByFilter(ctx context.Context, filter model.Filter, settlementEligibilitySnapshotsUpdateFields ...SettlementEligibilitySnapshotsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(settlementEligibilitySnapshotsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields SettlementEligibilitySnapshotsUpdateFieldList
		selectFields = NewSettlementEligibilitySnapshotsSelectFields()
	)
	for _, updateField := range settlementEligibilitySnapshotsUpdateFields {
		if updateField.settlementEligibilitySnapshotsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsSettlementEligibilitySnapshotsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := settlementEligibilitySnapshotsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeSettlementEligibilitySnapshotsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"settlement_eligibility_snapshots\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementEligibilitySnapshotsByFilter] error when try to update settlementEligibilitySnapshots by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateSettlementEligibilitySnapshotsByFilter] failed get rows affected")
	}
	return
}

var (
	settlementEligibilitySnapshotsQueries = struct {
		selectSettlementEligibilitySnapshots      string
		selectCountSettlementEligibilitySnapshots string
		deleteSettlementEligibilitySnapshots      string
		updateSettlementEligibilitySnapshots      string
		insertSettlementEligibilitySnapshots      string
	}{
		selectSettlementEligibilitySnapshots:      "SELECT %s FROM \"settlement_eligibility_snapshots\"",
		selectCountSettlementEligibilitySnapshots: "SELECT COUNT(\"id\") FROM \"settlement_eligibility_snapshots\"",
		deleteSettlementEligibilitySnapshots:      "DELETE FROM \"settlement_eligibility_snapshots\"",
		updateSettlementEligibilitySnapshots:      "UPDATE \"settlement_eligibility_snapshots\" SET %s ",
		insertSettlementEligibilitySnapshots:      "INSERT INTO \"settlement_eligibility_snapshots\" %s VALUES %s",
	}
)

type SettlementEligibilitySnapshotsRepository interface {
	CreateSettlementEligibilitySnapshots(ctx context.Context, settlementEligibilitySnapshots *model.SettlementEligibilitySnapshots, fieldsInsert ...SettlementEligibilitySnapshotsField) error
	BulkCreateSettlementEligibilitySnapshots(ctx context.Context, settlementEligibilitySnapshotsList []*model.SettlementEligibilitySnapshots, fieldsInsert ...SettlementEligibilitySnapshotsField) error
	ResolveSettlementEligibilitySnapshots(ctx context.Context, selectFields ...SettlementEligibilitySnapshotsField) (model.SettlementEligibilitySnapshotsList, error)
	ResolveSettlementEligibilitySnapshotsByID(ctx context.Context, primaryID model.SettlementEligibilitySnapshotsPrimaryID, selectFields ...SettlementEligibilitySnapshotsField) (model.SettlementEligibilitySnapshots, error)
	UpdateSettlementEligibilitySnapshotsByID(ctx context.Context, id model.SettlementEligibilitySnapshotsPrimaryID, settlementEligibilitySnapshots *model.SettlementEligibilitySnapshots, settlementEligibilitySnapshotsUpdateFields ...SettlementEligibilitySnapshotsUpdateField) error
	UpdateSettlementEligibilitySnapshotsByFilter(ctx context.Context, filter model.Filter, settlementEligibilitySnapshotsUpdateFields ...SettlementEligibilitySnapshotsUpdateField) (rowsAffected int64, err error)
	BulkUpdateSettlementEligibilitySnapshots(ctx context.Context, settlementEligibilitySnapshotsListMap map[model.SettlementEligibilitySnapshotsPrimaryID]*model.SettlementEligibilitySnapshots, SettlementEligibilitySnapshotssMapUpdateFieldsRequest map[model.SettlementEligibilitySnapshotsPrimaryID]SettlementEligibilitySnapshotsUpdateFieldList) (err error)
	DeleteSettlementEligibilitySnapshotsByID(ctx context.Context, id model.SettlementEligibilitySnapshotsPrimaryID) error
	BulkDeleteSettlementEligibilitySnapshotsByIDs(ctx context.Context, ids []model.SettlementEligibilitySnapshotsPrimaryID) error
	ResolveSettlementEligibilitySnapshotsByFilter(ctx context.Context, filter model.Filter) (result []model.SettlementEligibilitySnapshotsFilterResult, err error)
	IsExistSettlementEligibilitySnapshotsByIDs(ctx context.Context, ids []model.SettlementEligibilitySnapshotsPrimaryID) (exists bool, notFoundIds []model.SettlementEligibilitySnapshotsPrimaryID, err error)
	IsExistSettlementEligibilitySnapshotsByID(ctx context.Context, id model.SettlementEligibilitySnapshotsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
