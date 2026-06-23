package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/finance/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsMerchantBalanceSummary(merchantBalanceSummaryList []model.MerchantBalanceSummary, fieldsInsert ...MerchantBalanceSummaryField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewMerchantBalanceSummarySelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, merchantBalanceSummary := range merchantBalanceSummaryList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.MerchantPartyId():
				args = append(args, merchantBalanceSummary.MerchantPartyId)
			case selectField.CurrencyCode():
				args = append(args, merchantBalanceSummary.CurrencyCode)
			case selectField.PendingAmount():
				args = append(args, merchantBalanceSummary.PendingAmount)
			case selectField.AvailableAmount():
				args = append(args, merchantBalanceSummary.AvailableAmount)
			case selectField.ReservedAmount():
				args = append(args, merchantBalanceSummary.ReservedAmount)
			case selectField.PayableAmount():
				args = append(args, merchantBalanceSummary.PayableAmount)
			case selectField.PaidOutAmount():
				args = append(args, merchantBalanceSummary.PaidOutAmount)
			case selectField.NegativeAmount():
				args = append(args, merchantBalanceSummary.NegativeAmount)
			case selectField.RefundableAmount():
				args = append(args, merchantBalanceSummary.RefundableAmount)
			case selectField.RefreshedAt():
				args = append(args, merchantBalanceSummary.RefreshedAt)

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

func composeMerchantBalanceSummaryCompositePrimaryKeyWhere(primaryIDs []model.MerchantBalanceSummaryPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		merchantPartyId := "\"merchant_balance_summary\".\"merchant_party_id\" = ?"
		params = append(params, primaryID.MerchantPartyId)
		arrWhereQry = append(arrWhereQry, merchantPartyId)
		currencyCode := "\"merchant_balance_summary\".\"currency_code\" = ?"
		params = append(params, primaryID.CurrencyCode)
		arrWhereQry = append(arrWhereQry, currencyCode)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultMerchantBalanceSummarySelectFields() string {
	fields := NewMerchantBalanceSummarySelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeMerchantBalanceSummarySelectFields(selectFields ...MerchantBalanceSummaryField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type MerchantBalanceSummaryField string
type MerchantBalanceSummaryFieldList []MerchantBalanceSummaryField

type MerchantBalanceSummarySelectFields struct {
}

func (ss MerchantBalanceSummarySelectFields) MerchantPartyId() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("merchant_party_id")
}

func (ss MerchantBalanceSummarySelectFields) CurrencyCode() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("currency_code")
}

func (ss MerchantBalanceSummarySelectFields) PendingAmount() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("pending_amount")
}

func (ss MerchantBalanceSummarySelectFields) AvailableAmount() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("available_amount")
}

func (ss MerchantBalanceSummarySelectFields) ReservedAmount() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("reserved_amount")
}

func (ss MerchantBalanceSummarySelectFields) PayableAmount() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("payable_amount")
}

func (ss MerchantBalanceSummarySelectFields) PaidOutAmount() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("paid_out_amount")
}

func (ss MerchantBalanceSummarySelectFields) NegativeAmount() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("negative_amount")
}

func (ss MerchantBalanceSummarySelectFields) RefundableAmount() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("refundable_amount")
}

func (ss MerchantBalanceSummarySelectFields) RefreshedAt() MerchantBalanceSummaryField {
	return MerchantBalanceSummaryField("refreshed_at")
}

func (ss MerchantBalanceSummarySelectFields) All() MerchantBalanceSummaryFieldList {
	return []MerchantBalanceSummaryField{
		ss.MerchantPartyId(),
		ss.CurrencyCode(),
		ss.PendingAmount(),
		ss.AvailableAmount(),
		ss.ReservedAmount(),
		ss.PayableAmount(),
		ss.PaidOutAmount(),
		ss.NegativeAmount(),
		ss.RefundableAmount(),
		ss.RefreshedAt(),
	}
}

func NewMerchantBalanceSummarySelectFields() MerchantBalanceSummarySelectFields {
	return MerchantBalanceSummarySelectFields{}
}

type MerchantBalanceSummaryUpdateFieldOption struct {
	useIncrement bool
}
type MerchantBalanceSummaryUpdateField struct {
	merchantBalanceSummaryField MerchantBalanceSummaryField
	opt                         MerchantBalanceSummaryUpdateFieldOption
	value                       interface{}
}
type MerchantBalanceSummaryUpdateFieldList []MerchantBalanceSummaryUpdateField

func defaultMerchantBalanceSummaryUpdateFieldOption() MerchantBalanceSummaryUpdateFieldOption {
	return MerchantBalanceSummaryUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementMerchantBalanceSummaryOption(useIncrement bool) func(*MerchantBalanceSummaryUpdateFieldOption) {
	return func(pcufo *MerchantBalanceSummaryUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewMerchantBalanceSummaryUpdateField(field MerchantBalanceSummaryField, val interface{}, opts ...func(*MerchantBalanceSummaryUpdateFieldOption)) MerchantBalanceSummaryUpdateField {
	defaultOpt := defaultMerchantBalanceSummaryUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return MerchantBalanceSummaryUpdateField{
		merchantBalanceSummaryField: field,
		value:                       val,
		opt:                         defaultOpt,
	}
}
func defaultMerchantBalanceSummaryUpdateFields(merchantBalanceSummary model.MerchantBalanceSummary) (merchantBalanceSummaryUpdateFieldList MerchantBalanceSummaryUpdateFieldList) {
	selectFields := NewMerchantBalanceSummarySelectFields()
	merchantBalanceSummaryUpdateFieldList = append(merchantBalanceSummaryUpdateFieldList,
		NewMerchantBalanceSummaryUpdateField(selectFields.MerchantPartyId(), merchantBalanceSummary.MerchantPartyId),
		NewMerchantBalanceSummaryUpdateField(selectFields.CurrencyCode(), merchantBalanceSummary.CurrencyCode),
		NewMerchantBalanceSummaryUpdateField(selectFields.PendingAmount(), merchantBalanceSummary.PendingAmount),
		NewMerchantBalanceSummaryUpdateField(selectFields.AvailableAmount(), merchantBalanceSummary.AvailableAmount),
		NewMerchantBalanceSummaryUpdateField(selectFields.ReservedAmount(), merchantBalanceSummary.ReservedAmount),
		NewMerchantBalanceSummaryUpdateField(selectFields.PayableAmount(), merchantBalanceSummary.PayableAmount),
		NewMerchantBalanceSummaryUpdateField(selectFields.PaidOutAmount(), merchantBalanceSummary.PaidOutAmount),
		NewMerchantBalanceSummaryUpdateField(selectFields.NegativeAmount(), merchantBalanceSummary.NegativeAmount),
		NewMerchantBalanceSummaryUpdateField(selectFields.RefundableAmount(), merchantBalanceSummary.RefundableAmount),
		NewMerchantBalanceSummaryUpdateField(selectFields.RefreshedAt(), merchantBalanceSummary.RefreshedAt),
	)
	return
}
func composeUpdateFieldsMerchantBalanceSummaryCommand(merchantBalanceSummaryUpdateFieldList MerchantBalanceSummaryUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range merchantBalanceSummaryUpdateFieldList {
		field := string(updateField.merchantBalanceSummaryField)
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

func (repo *RepositoryImpl) BulkCreateMerchantBalanceSummary(ctx context.Context, merchantBalanceSummaryList []*model.MerchantBalanceSummary, fieldsInsert ...MerchantBalanceSummaryField) (err error) {
	var (
		fieldsStr                       string
		valueListStr                    []string
		argsList                        []interface{}
		primaryIds                      []model.MerchantBalanceSummaryPrimaryID
		merchantBalanceSummaryValueList []model.MerchantBalanceSummary
	)

	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceSummarySelectFields()
		fieldsInsert = selectField.All()
	}
	for _, merchantBalanceSummary := range merchantBalanceSummaryList {

		primaryIds = append(primaryIds, merchantBalanceSummary.ToMerchantBalanceSummaryPrimaryID())

		merchantBalanceSummaryValueList = append(merchantBalanceSummaryValueList, *merchantBalanceSummary)
	}

	_, notFoundIds, err := repo.IsExistMerchantBalanceSummaryByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceSummary] failed checking merchantBalanceSummary whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.MerchantBalanceSummaryPrimaryID{}
		mapNotFoundIds := map[model.MerchantBalanceSummaryPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "merchantBalanceSummary", fmt.Sprintf("merchantBalanceSummary with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsMerchantBalanceSummary(merchantBalanceSummaryValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(merchantBalanceSummaryQueries.insertMerchantBalanceSummary, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateMerchantBalanceSummary] failed exec create merchantBalanceSummary query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteMerchantBalanceSummaryByIDs(ctx context.Context, primaryIDs []model.MerchantBalanceSummaryPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistMerchantBalanceSummaryByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceSummaryByIDs] failed checking merchantBalanceSummary whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSummary with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	whereQuery, params := composeMerchantBalanceSummaryCompositePrimaryKeyWhere(primaryIDs)

	commandQuery := fmt.Sprintf(merchantBalanceSummaryQueries.deleteMerchantBalanceSummary + " WHERE " + whereQuery)

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteMerchantBalanceSummaryByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistMerchantBalanceSummaryByIDs(ctx context.Context, ids []model.MerchantBalanceSummaryPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceSummaryPrimaryID, err error) {

	whereQuery, params := composeMerchantBalanceSummaryCompositePrimaryKeyWhere(ids)

	query := fmt.Sprintf(merchantBalanceSummaryQueries.selectMerchantBalanceSummary, " \"merchant_party_id\"    , \"currency_code\"  ") + " WHERE " + whereQuery

	query = repo.db.Read.Rebind(query)
	var resIds []model.MerchantBalanceSummaryPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceSummaryByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.MerchantBalanceSummaryPrimaryID]bool{}
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

// BulkUpdateMerchantBalanceSummary is used to bulk update merchantBalanceSummary, by default it will update all field
// if want to update specific field, then fill merchantBalanceSummarysMapUpdateFieldsRequest else please fill merchantBalanceSummarysMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateMerchantBalanceSummary(ctx context.Context, merchantBalanceSummarysMap map[model.MerchantBalanceSummaryPrimaryID]*model.MerchantBalanceSummary, merchantBalanceSummarysMapUpdateFieldsRequest map[model.MerchantBalanceSummaryPrimaryID]MerchantBalanceSummaryUpdateFieldList) (err error) {
	if len(merchantBalanceSummarysMap) == 0 && len(merchantBalanceSummarysMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		merchantBalanceSummarysMapUpdateField map[model.MerchantBalanceSummaryPrimaryID]MerchantBalanceSummaryUpdateFieldList = map[model.MerchantBalanceSummaryPrimaryID]MerchantBalanceSummaryUpdateFieldList{}
		asTableValues                         string                                                                          = "myvalues"
	)

	if len(merchantBalanceSummarysMap) > 0 {
		for id, merchantBalanceSummary := range merchantBalanceSummarysMap {
			if merchantBalanceSummary == nil {
				log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceSummary] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			merchantBalanceSummarysMapUpdateField[id] = defaultMerchantBalanceSummaryUpdateFields(*merchantBalanceSummary)
		}
	} else {
		merchantBalanceSummarysMapUpdateField = merchantBalanceSummarysMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateMerchantBalanceSummaryQuery(merchantBalanceSummarysMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistMerchantBalanceSummaryByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceSummary] failed checking merchantBalanceSummary whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSummary with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeMerchantBalanceSummaryCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"merchant_balance_summary\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateMerchantBalanceSummary] failed exec query")
	}
	return
}

type MerchantBalanceSummaryFieldParameter struct {
	param string
	args  []interface{}
}

func NewMerchantBalanceSummaryFieldParameter(param string, args ...interface{}) MerchantBalanceSummaryFieldParameter {
	return MerchantBalanceSummaryFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateMerchantBalanceSummaryQuery(mapMerchantBalanceSummarys map[model.MerchantBalanceSummaryPrimaryID]MerchantBalanceSummaryUpdateFieldList, asTableValues string) (primaryIDs []model.MerchantBalanceSummaryPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.MerchantBalanceSummaryPrimaryID]map[string]interface{}{}
	merchantBalanceSummarySelectFields := NewMerchantBalanceSummarySelectFields()
	for id, updateFields := range mapMerchantBalanceSummarys {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.merchantBalanceSummaryField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapMerchantBalanceSummarys[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetMerchantBalanceSummaryFieldType(updateField.merchantBalanceSummaryField)))
			args = append(args, fields[string(updateField.merchantBalanceSummaryField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.merchantBalanceSummaryField))
		if updateField.merchantBalanceSummaryField == merchantBalanceSummarySelectFields.MerchantPartyId() || updateField.merchantBalanceSummaryField == merchantBalanceSummarySelectFields.CurrencyCode() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.merchantBalanceSummaryField, asTableValues, updateField.merchantBalanceSummaryField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.merchantBalanceSummaryField,
				"\"merchant_balance_summary\"", updateField.merchantBalanceSummaryField,
				asTableValues, updateField.merchantBalanceSummaryField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeMerchantBalanceSummaryCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.MerchantBalanceSummaryPrimaryID, asTableValue string) (whereQry string) {
	merchantBalanceSummarySelectFields := NewMerchantBalanceSummarySelectFields()
	var arrWhereQry []string
	merchantPartyId := fmt.Sprintf("\"merchant_balance_summary\".\"merchant_party_id\" = %s.\"merchant_party_id\"::"+GetMerchantBalanceSummaryFieldType(merchantBalanceSummarySelectFields.MerchantPartyId()), asTableValue)
	arrWhereQry = append(arrWhereQry, merchantPartyId)
	currencyCode := fmt.Sprintf("\"merchant_balance_summary\".\"currency_code\" = %s.\"currency_code\"::"+GetMerchantBalanceSummaryFieldType(merchantBalanceSummarySelectFields.CurrencyCode()), asTableValue)
	arrWhereQry = append(arrWhereQry, currencyCode)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetMerchantBalanceSummaryFieldType(merchantBalanceSummaryField MerchantBalanceSummaryField) string {
	selectMerchantBalanceSummaryFields := NewMerchantBalanceSummarySelectFields()
	switch merchantBalanceSummaryField {

	case selectMerchantBalanceSummaryFields.MerchantPartyId():
		return "uuid"

	case selectMerchantBalanceSummaryFields.CurrencyCode():
		return "text"

	case selectMerchantBalanceSummaryFields.PendingAmount():
		return "numeric"

	case selectMerchantBalanceSummaryFields.AvailableAmount():
		return "numeric"

	case selectMerchantBalanceSummaryFields.ReservedAmount():
		return "numeric"

	case selectMerchantBalanceSummaryFields.PayableAmount():
		return "numeric"

	case selectMerchantBalanceSummaryFields.PaidOutAmount():
		return "numeric"

	case selectMerchantBalanceSummaryFields.NegativeAmount():
		return "numeric"

	case selectMerchantBalanceSummaryFields.RefundableAmount():
		return "numeric"

	case selectMerchantBalanceSummaryFields.RefreshedAt():
		return "timestamptz"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateMerchantBalanceSummary(ctx context.Context, merchantBalanceSummary *model.MerchantBalanceSummary, fieldsInsert ...MerchantBalanceSummaryField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewMerchantBalanceSummarySelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.MerchantBalanceSummaryPrimaryID{
		MerchantPartyId: merchantBalanceSummary.MerchantPartyId,
		CurrencyCode:    merchantBalanceSummary.CurrencyCode,
	}
	exists, err := repo.IsExistMerchantBalanceSummaryByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceSummary] failed checking merchantBalanceSummary whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "merchantBalanceSummary", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsMerchantBalanceSummary([]model.MerchantBalanceSummary{*merchantBalanceSummary}, fieldsInsert...)
	commandQuery := fmt.Sprintf(merchantBalanceSummaryQueries.insertMerchantBalanceSummary, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateMerchantBalanceSummary] failed exec create merchantBalanceSummary query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteMerchantBalanceSummaryByID(ctx context.Context, primaryID model.MerchantBalanceSummaryPrimaryID) (err error) {
	exists, err := repo.IsExistMerchantBalanceSummaryByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceSummaryByID] failed checking merchantBalanceSummary whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSummary with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeMerchantBalanceSummaryCompositePrimaryKeyWhere([]model.MerchantBalanceSummaryPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(merchantBalanceSummaryQueries.deleteMerchantBalanceSummary + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteMerchantBalanceSummaryByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceSummaryByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceSummaryFilterResult, err error) {
	query, args, err := composeMerchantBalanceSummaryFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSummaryByFilter] failed compose merchantBalanceSummary filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSummaryByFilter] failed get merchantBalanceSummary by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeMerchantBalanceSummaryFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.MerchantBalanceSummaryFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeMerchantBalanceSummaryFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeMerchantBalanceSummarySortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeMerchantBalanceSummaryFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 10 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewMerchantBalanceSummaryFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 10+1)
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["pending_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"pending_amount\"")
			selectedColumns["pending_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["available_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"available_amount\"")
			selectedColumns["available_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["reserved_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"reserved_amount\"")
			selectedColumns["reserved_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["payable_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"payable_amount\"")
			selectedColumns["payable_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["paid_out_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"paid_out_amount\"")
			selectedColumns["paid_out_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["negative_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"negative_amount\"")
			selectedColumns["negative_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["refundable_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"refundable_amount\"")
			selectedColumns["refundable_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["refreshed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"refreshed_at\"")
			selectedColumns["refreshed_at"] = struct{}{}
		}

	} else {
		selectColumns = make([]string, 0, len(filter.SelectFields)+1)
		for _, field := range filter.SelectFields {
			if err = addColumn(field); err != nil {
				return
			}
		}
	}

	return
}

type merchantBalanceSummaryFilterPlaceholder struct {
	index int
}

func (p *merchantBalanceSummaryFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeMerchantBalanceSummaryFilterPredicate(filterField model.FilterField, placeholders *merchantBalanceSummaryFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewMerchantBalanceSummaryFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeMerchantBalanceSummaryFilterSQLExpr(spec)
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

func composeMerchantBalanceSummaryFilterGroup(group model.FilterGroup, placeholders *merchantBalanceSummaryFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeMerchantBalanceSummaryFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeMerchantBalanceSummaryFilterGroup(child, placeholders, args, requiredJoins)
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

func composeMerchantBalanceSummaryFilterWhereQueries(filter model.Filter, placeholders *merchantBalanceSummaryFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeMerchantBalanceSummaryFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeMerchantBalanceSummaryFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeMerchantBalanceSummaryFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateMerchantBalanceSummaryFieldNameFilter(filter)
	if err != nil {
		return
	}
	isCursorMode := filter.Pagination.IsCursorMode()
	requiredJoins := map[string]bool{}
	if isCursorMode {
		err = failure.BadRequestFromString("cursor pagination requires a single primary key")
		return
	}

	selectColumns, err := composeMerchantBalanceSummaryFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := merchantBalanceSummaryFilterPlaceholder{index: 1}
	whereQueries, err := composeMerchantBalanceSummaryFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
	if err != nil {
		return
	}

	sortQuery := []string{}
	if isCursorMode {

	} else {
		for _, sort := range filter.Sorts {
			spec, found := model.NewMerchantBalanceSummaryFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeMerchantBalanceSummaryFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeMerchantBalanceSummarySortOrder(sort.Order)
			if sortErr != nil {
				err = sortErr
				return
			}
			sortQuery = append(sortQuery, fmt.Sprintf("%s %s", sqlExpr, sortOrder))
		}

	}

	query = fmt.Sprintf("SELECT %s FROM \"merchant_balance_summary\" base%s", strings.Join(selectColumns, ","), composeMerchantBalanceSummaryFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistMerchantBalanceSummaryByID(ctx context.Context, primaryID model.MerchantBalanceSummaryPrimaryID) (exists bool, err error) {
	whereQuery, params := composeMerchantBalanceSummaryCompositePrimaryKeyWhere([]model.MerchantBalanceSummaryPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", merchantBalanceSummaryQueries.selectCountMerchantBalanceSummary, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistMerchantBalanceSummaryByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceSummary(ctx context.Context, selectFields ...MerchantBalanceSummaryField) (merchantBalanceSummaryList model.MerchantBalanceSummaryList, err error) {
	var (
		defaultMerchantBalanceSummarySelectFields = defaultMerchantBalanceSummarySelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceSummarySelectFields = composeMerchantBalanceSummarySelectFields(selectFields...)
	}
	query := fmt.Sprintf(merchantBalanceSummaryQueries.selectMerchantBalanceSummary, defaultMerchantBalanceSummarySelectFields)

	err = repo.db.Read.SelectContext(ctx, &merchantBalanceSummaryList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSummary] failed get merchantBalanceSummary list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveMerchantBalanceSummaryByID(ctx context.Context, primaryID model.MerchantBalanceSummaryPrimaryID, selectFields ...MerchantBalanceSummaryField) (merchantBalanceSummary model.MerchantBalanceSummary, err error) {
	var (
		defaultMerchantBalanceSummarySelectFields = defaultMerchantBalanceSummarySelectFields()
	)
	if len(selectFields) > 0 {
		defaultMerchantBalanceSummarySelectFields = composeMerchantBalanceSummarySelectFields(selectFields...)
	}
	whereQry, params := composeMerchantBalanceSummaryCompositePrimaryKeyWhere([]model.MerchantBalanceSummaryPrimaryID{primaryID})
	query := fmt.Sprintf(merchantBalanceSummaryQueries.selectMerchantBalanceSummary+" WHERE "+whereQry, defaultMerchantBalanceSummarySelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &merchantBalanceSummary, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("merchantBalanceSummary with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveMerchantBalanceSummaryByID] failed get merchantBalanceSummary")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateMerchantBalanceSummaryByID(ctx context.Context, primaryID model.MerchantBalanceSummaryPrimaryID, merchantBalanceSummary *model.MerchantBalanceSummary, merchantBalanceSummaryUpdateFields ...MerchantBalanceSummaryUpdateField) (err error) {
	exists, err := repo.IsExistMerchantBalanceSummaryByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSummary] failed checking merchantBalanceSummary whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("merchantBalanceSummary with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if merchantBalanceSummary == nil {
		if len(merchantBalanceSummaryUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateMerchantBalanceSummaryByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		merchantBalanceSummary = &model.MerchantBalanceSummary{}
	}
	var (
		defaultMerchantBalanceSummaryUpdateFields = defaultMerchantBalanceSummaryUpdateFields(*merchantBalanceSummary)
		tempUpdateField                           MerchantBalanceSummaryUpdateFieldList
		selectFields                              = NewMerchantBalanceSummarySelectFields()
	)
	if len(merchantBalanceSummaryUpdateFields) > 0 {
		for _, updateField := range merchantBalanceSummaryUpdateFields {
			if updateField.merchantBalanceSummaryField == selectFields.MerchantPartyId() || updateField.merchantBalanceSummaryField == selectFields.CurrencyCode() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultMerchantBalanceSummaryUpdateFields = tempUpdateField
	}
	whereQuery, params := composeMerchantBalanceSummaryCompositePrimaryKeyWhere([]model.MerchantBalanceSummaryPrimaryID{primaryID})
	fields, args := composeUpdateFieldsMerchantBalanceSummaryCommand(defaultMerchantBalanceSummaryUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(merchantBalanceSummaryQueries.updateMerchantBalanceSummary+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSummary] error when try to update merchantBalanceSummary by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateMerchantBalanceSummaryByFilter(ctx context.Context, filter model.Filter, merchantBalanceSummaryUpdateFields ...MerchantBalanceSummaryUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(merchantBalanceSummaryUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields MerchantBalanceSummaryUpdateFieldList
		selectFields = NewMerchantBalanceSummarySelectFields()
	)
	for _, updateField := range merchantBalanceSummaryUpdateFields {
		if updateField.merchantBalanceSummaryField == selectFields.MerchantPartyId() || updateField.merchantBalanceSummaryField == selectFields.CurrencyCode() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsMerchantBalanceSummaryCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := merchantBalanceSummaryFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeMerchantBalanceSummaryFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"merchant_balance_summary\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSummaryByFilter] error when try to update merchantBalanceSummary by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateMerchantBalanceSummaryByFilter] failed get rows affected")
	}
	return
}

var (
	merchantBalanceSummaryQueries = struct {
		selectMerchantBalanceSummary      string
		selectCountMerchantBalanceSummary string
		deleteMerchantBalanceSummary      string
		updateMerchantBalanceSummary      string
		insertMerchantBalanceSummary      string
	}{
		selectMerchantBalanceSummary:      "SELECT %s FROM \"merchant_balance_summary\"",
		selectCountMerchantBalanceSummary: "SELECT COUNT(*) FROM \"merchant_balance_summary\"",
		deleteMerchantBalanceSummary:      "DELETE FROM \"merchant_balance_summary\"",
		updateMerchantBalanceSummary:      "UPDATE \"merchant_balance_summary\" SET %s ",
		insertMerchantBalanceSummary:      "INSERT INTO \"merchant_balance_summary\" %s VALUES %s",
	}
)

type MerchantBalanceSummaryRepository interface {
	CreateMerchantBalanceSummary(ctx context.Context, merchantBalanceSummary *model.MerchantBalanceSummary, fieldsInsert ...MerchantBalanceSummaryField) error
	BulkCreateMerchantBalanceSummary(ctx context.Context, merchantBalanceSummaryList []*model.MerchantBalanceSummary, fieldsInsert ...MerchantBalanceSummaryField) error
	ResolveMerchantBalanceSummary(ctx context.Context, selectFields ...MerchantBalanceSummaryField) (model.MerchantBalanceSummaryList, error)
	ResolveMerchantBalanceSummaryByID(ctx context.Context, primaryID model.MerchantBalanceSummaryPrimaryID, selectFields ...MerchantBalanceSummaryField) (model.MerchantBalanceSummary, error)
	UpdateMerchantBalanceSummaryByID(ctx context.Context, id model.MerchantBalanceSummaryPrimaryID, merchantBalanceSummary *model.MerchantBalanceSummary, merchantBalanceSummaryUpdateFields ...MerchantBalanceSummaryUpdateField) error
	UpdateMerchantBalanceSummaryByFilter(ctx context.Context, filter model.Filter, merchantBalanceSummaryUpdateFields ...MerchantBalanceSummaryUpdateField) (rowsAffected int64, err error)
	BulkUpdateMerchantBalanceSummary(ctx context.Context, merchantBalanceSummaryListMap map[model.MerchantBalanceSummaryPrimaryID]*model.MerchantBalanceSummary, MerchantBalanceSummarysMapUpdateFieldsRequest map[model.MerchantBalanceSummaryPrimaryID]MerchantBalanceSummaryUpdateFieldList) (err error)
	DeleteMerchantBalanceSummaryByID(ctx context.Context, id model.MerchantBalanceSummaryPrimaryID) error
	BulkDeleteMerchantBalanceSummaryByIDs(ctx context.Context, ids []model.MerchantBalanceSummaryPrimaryID) error
	ResolveMerchantBalanceSummaryByFilter(ctx context.Context, filter model.Filter) (result []model.MerchantBalanceSummaryFilterResult, err error)
	IsExistMerchantBalanceSummaryByIDs(ctx context.Context, ids []model.MerchantBalanceSummaryPrimaryID) (exists bool, notFoundIds []model.MerchantBalanceSummaryPrimaryID, err error)
	IsExistMerchantBalanceSummaryByID(ctx context.Context, id model.MerchantBalanceSummaryPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
