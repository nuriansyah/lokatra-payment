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

func composeInsertFieldsAndParamsChargebacks(chargebacksList []model.Chargebacks, fieldsInsert ...ChargebacksField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewChargebacksSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, chargebacks := range chargebacksList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, chargebacks.Id)
			case selectField.ChargebackCode():
				args = append(args, chargebacks.ChargebackCode)
			case selectField.PaymentRefId():
				args = append(args, chargebacks.PaymentRefId)
			case selectField.MerchantPartyId():
				args = append(args, chargebacks.MerchantPartyId)
			case selectField.ProviderAccountId():
				args = append(args, chargebacks.ProviderAccountId)
			case selectField.CurrencyCode():
				args = append(args, chargebacks.CurrencyCode)
			case selectField.DisputedAmount():
				args = append(args, chargebacks.DisputedAmount)
			case selectField.ChargebackStatus():
				args = append(args, chargebacks.ChargebackStatus)
			case selectField.ReasonCode():
				args = append(args, chargebacks.ReasonCode)
			case selectField.OpenedAt():
				args = append(args, chargebacks.OpenedAt)
			case selectField.ClosedAt():
				args = append(args, chargebacks.ClosedAt)
			case selectField.DueAt():
				args = append(args, chargebacks.DueAt)
			case selectField.Metadata():
				args = append(args, chargebacks.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, chargebacks.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, chargebacks.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, chargebacks.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, chargebacks.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, chargebacks.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, chargebacks.MetaDeletedBy)

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

func composeChargebacksCompositePrimaryKeyWhere(primaryIDs []model.ChargebacksPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"chargebacks\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultChargebacksSelectFields() string {
	fields := NewChargebacksSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeChargebacksSelectFields(selectFields ...ChargebacksField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ChargebacksField string
type ChargebacksFieldList []ChargebacksField

type ChargebacksSelectFields struct {
}

func (ss ChargebacksSelectFields) Id() ChargebacksField {
	return ChargebacksField("id")
}

func (ss ChargebacksSelectFields) ChargebackCode() ChargebacksField {
	return ChargebacksField("chargeback_code")
}

func (ss ChargebacksSelectFields) PaymentRefId() ChargebacksField {
	return ChargebacksField("payment_ref_id")
}

func (ss ChargebacksSelectFields) MerchantPartyId() ChargebacksField {
	return ChargebacksField("merchant_party_id")
}

func (ss ChargebacksSelectFields) ProviderAccountId() ChargebacksField {
	return ChargebacksField("provider_account_id")
}

func (ss ChargebacksSelectFields) CurrencyCode() ChargebacksField {
	return ChargebacksField("currency_code")
}

func (ss ChargebacksSelectFields) DisputedAmount() ChargebacksField {
	return ChargebacksField("disputed_amount")
}

func (ss ChargebacksSelectFields) ChargebackStatus() ChargebacksField {
	return ChargebacksField("chargeback_status")
}

func (ss ChargebacksSelectFields) ReasonCode() ChargebacksField {
	return ChargebacksField("reason_code")
}

func (ss ChargebacksSelectFields) OpenedAt() ChargebacksField {
	return ChargebacksField("opened_at")
}

func (ss ChargebacksSelectFields) ClosedAt() ChargebacksField {
	return ChargebacksField("closed_at")
}

func (ss ChargebacksSelectFields) DueAt() ChargebacksField {
	return ChargebacksField("due_at")
}

func (ss ChargebacksSelectFields) Metadata() ChargebacksField {
	return ChargebacksField("metadata")
}

func (ss ChargebacksSelectFields) MetaCreatedAt() ChargebacksField {
	return ChargebacksField("meta_created_at")
}

func (ss ChargebacksSelectFields) MetaCreatedBy() ChargebacksField {
	return ChargebacksField("meta_created_by")
}

func (ss ChargebacksSelectFields) MetaUpdatedAt() ChargebacksField {
	return ChargebacksField("meta_updated_at")
}

func (ss ChargebacksSelectFields) MetaUpdatedBy() ChargebacksField {
	return ChargebacksField("meta_updated_by")
}

func (ss ChargebacksSelectFields) MetaDeletedAt() ChargebacksField {
	return ChargebacksField("meta_deleted_at")
}

func (ss ChargebacksSelectFields) MetaDeletedBy() ChargebacksField {
	return ChargebacksField("meta_deleted_by")
}

func (ss ChargebacksSelectFields) All() ChargebacksFieldList {
	return []ChargebacksField{
		ss.Id(),
		ss.ChargebackCode(),
		ss.PaymentRefId(),
		ss.MerchantPartyId(),
		ss.ProviderAccountId(),
		ss.CurrencyCode(),
		ss.DisputedAmount(),
		ss.ChargebackStatus(),
		ss.ReasonCode(),
		ss.OpenedAt(),
		ss.ClosedAt(),
		ss.DueAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewChargebacksSelectFields() ChargebacksSelectFields {
	return ChargebacksSelectFields{}
}

type ChargebacksUpdateFieldOption struct {
	useIncrement bool
}
type ChargebacksUpdateField struct {
	chargebacksField ChargebacksField
	opt              ChargebacksUpdateFieldOption
	value            interface{}
}
type ChargebacksUpdateFieldList []ChargebacksUpdateField

func defaultChargebacksUpdateFieldOption() ChargebacksUpdateFieldOption {
	return ChargebacksUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementChargebacksOption(useIncrement bool) func(*ChargebacksUpdateFieldOption) {
	return func(pcufo *ChargebacksUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewChargebacksUpdateField(field ChargebacksField, val interface{}, opts ...func(*ChargebacksUpdateFieldOption)) ChargebacksUpdateField {
	defaultOpt := defaultChargebacksUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ChargebacksUpdateField{
		chargebacksField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultChargebacksUpdateFields(chargebacks model.Chargebacks) (chargebacksUpdateFieldList ChargebacksUpdateFieldList) {
	selectFields := NewChargebacksSelectFields()
	chargebacksUpdateFieldList = append(chargebacksUpdateFieldList,
		NewChargebacksUpdateField(selectFields.Id(), chargebacks.Id),
		NewChargebacksUpdateField(selectFields.ChargebackCode(), chargebacks.ChargebackCode),
		NewChargebacksUpdateField(selectFields.PaymentRefId(), chargebacks.PaymentRefId),
		NewChargebacksUpdateField(selectFields.MerchantPartyId(), chargebacks.MerchantPartyId),
		NewChargebacksUpdateField(selectFields.ProviderAccountId(), chargebacks.ProviderAccountId),
		NewChargebacksUpdateField(selectFields.CurrencyCode(), chargebacks.CurrencyCode),
		NewChargebacksUpdateField(selectFields.DisputedAmount(), chargebacks.DisputedAmount),
		NewChargebacksUpdateField(selectFields.ChargebackStatus(), chargebacks.ChargebackStatus),
		NewChargebacksUpdateField(selectFields.ReasonCode(), chargebacks.ReasonCode),
		NewChargebacksUpdateField(selectFields.OpenedAt(), chargebacks.OpenedAt),
		NewChargebacksUpdateField(selectFields.ClosedAt(), chargebacks.ClosedAt),
		NewChargebacksUpdateField(selectFields.DueAt(), chargebacks.DueAt),
		NewChargebacksUpdateField(selectFields.Metadata(), chargebacks.Metadata),
		NewChargebacksUpdateField(selectFields.MetaCreatedAt(), chargebacks.MetaCreatedAt),
		NewChargebacksUpdateField(selectFields.MetaCreatedBy(), chargebacks.MetaCreatedBy),
		NewChargebacksUpdateField(selectFields.MetaUpdatedAt(), chargebacks.MetaUpdatedAt),
		NewChargebacksUpdateField(selectFields.MetaUpdatedBy(), chargebacks.MetaUpdatedBy),
		NewChargebacksUpdateField(selectFields.MetaDeletedAt(), chargebacks.MetaDeletedAt),
		NewChargebacksUpdateField(selectFields.MetaDeletedBy(), chargebacks.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsChargebacksCommand(chargebacksUpdateFieldList ChargebacksUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range chargebacksUpdateFieldList {
		field := string(updateField.chargebacksField)
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

func (repo *RepositoryImpl) BulkCreateChargebacks(ctx context.Context, chargebacksList []*model.Chargebacks, fieldsInsert ...ChargebacksField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.ChargebacksPrimaryID
		chargebacksValueList []model.Chargebacks
	)

	if len(fieldsInsert) == 0 {
		selectField := NewChargebacksSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, chargebacks := range chargebacksList {

		primaryIds = append(primaryIds, chargebacks.ToChargebacksPrimaryID())

		chargebacksValueList = append(chargebacksValueList, *chargebacks)
	}

	_, notFoundIds, err := repo.IsExistChargebacksByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateChargebacks] failed checking chargebacks whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ChargebacksPrimaryID{}
		mapNotFoundIds := map[model.ChargebacksPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "chargebacks", fmt.Sprintf("chargebacks with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsChargebacks(chargebacksValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(chargebacksQueries.insertChargebacks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateChargebacks] failed exec create chargebacks query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteChargebacksByIDs(ctx context.Context, primaryIDs []model.ChargebacksPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistChargebacksByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteChargebacksByIDs] failed checking chargebacks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebacks with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"chargebacks\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(chargebacksQueries.deleteChargebacks + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteChargebacksByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteChargebacksByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistChargebacksByIDs(ctx context.Context, ids []model.ChargebacksPrimaryID) (exists bool, notFoundIds []model.ChargebacksPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"chargebacks\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(chargebacksQueries.selectChargebacks, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistChargebacksByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ChargebacksPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistChargebacksByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ChargebacksPrimaryID]bool{}
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

// BulkUpdateChargebacks is used to bulk update chargebacks, by default it will update all field
// if want to update specific field, then fill chargebackssMapUpdateFieldsRequest else please fill chargebackssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateChargebacks(ctx context.Context, chargebackssMap map[model.ChargebacksPrimaryID]*model.Chargebacks, chargebackssMapUpdateFieldsRequest map[model.ChargebacksPrimaryID]ChargebacksUpdateFieldList) (err error) {
	if len(chargebackssMap) == 0 && len(chargebackssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		chargebackssMapUpdateField map[model.ChargebacksPrimaryID]ChargebacksUpdateFieldList = map[model.ChargebacksPrimaryID]ChargebacksUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(chargebackssMap) > 0 {
		for id, chargebacks := range chargebackssMap {
			if chargebacks == nil {
				log.Error().Err(err).Msg("[BulkUpdateChargebacks] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			chargebackssMapUpdateField[id] = defaultChargebacksUpdateFields(*chargebacks)
		}
	} else {
		chargebackssMapUpdateField = chargebackssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateChargebacksQuery(chargebackssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistChargebacksByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateChargebacks] failed checking chargebacks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebacks with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeChargebacksCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"chargebacks\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateChargebacks] failed exec query")
	}
	return
}

type ChargebacksFieldParameter struct {
	param string
	args  []interface{}
}

func NewChargebacksFieldParameter(param string, args ...interface{}) ChargebacksFieldParameter {
	return ChargebacksFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateChargebacksQuery(mapChargebackss map[model.ChargebacksPrimaryID]ChargebacksUpdateFieldList, asTableValues string) (primaryIDs []model.ChargebacksPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ChargebacksPrimaryID]map[string]interface{}{}
	chargebacksSelectFields := NewChargebacksSelectFields()
	for id, updateFields := range mapChargebackss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.chargebacksField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapChargebackss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetChargebacksFieldType(updateField.chargebacksField)))
			args = append(args, fields[string(updateField.chargebacksField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.chargebacksField))
		if updateField.chargebacksField == chargebacksSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.chargebacksField, asTableValues, updateField.chargebacksField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.chargebacksField,
				"\"chargebacks\"", updateField.chargebacksField,
				asTableValues, updateField.chargebacksField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeChargebacksCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ChargebacksPrimaryID, asTableValue string) (whereQry string) {
	chargebacksSelectFields := NewChargebacksSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"chargebacks\".\"id\" = %s.\"id\"::"+GetChargebacksFieldType(chargebacksSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetChargebacksFieldType(chargebacksField ChargebacksField) string {
	selectChargebacksFields := NewChargebacksSelectFields()
	switch chargebacksField {

	case selectChargebacksFields.Id():
		return "uuid"

	case selectChargebacksFields.ChargebackCode():
		return "text"

	case selectChargebacksFields.PaymentRefId():
		return "uuid"

	case selectChargebacksFields.MerchantPartyId():
		return "uuid"

	case selectChargebacksFields.ProviderAccountId():
		return "uuid"

	case selectChargebacksFields.CurrencyCode():
		return "text"

	case selectChargebacksFields.DisputedAmount():
		return "numeric"

	case selectChargebacksFields.ChargebackStatus():
		return "chargeback_status_enum"

	case selectChargebacksFields.ReasonCode():
		return "text"

	case selectChargebacksFields.OpenedAt():
		return "timestamptz"

	case selectChargebacksFields.ClosedAt():
		return "timestamptz"

	case selectChargebacksFields.DueAt():
		return "timestamptz"

	case selectChargebacksFields.Metadata():
		return "jsonb"

	case selectChargebacksFields.MetaCreatedAt():
		return "timestamptz"

	case selectChargebacksFields.MetaCreatedBy():
		return "uuid"

	case selectChargebacksFields.MetaUpdatedAt():
		return "timestamptz"

	case selectChargebacksFields.MetaUpdatedBy():
		return "uuid"

	case selectChargebacksFields.MetaDeletedAt():
		return "timestamptz"

	case selectChargebacksFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateChargebacks(ctx context.Context, chargebacks *model.Chargebacks, fieldsInsert ...ChargebacksField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewChargebacksSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ChargebacksPrimaryID{
		Id: chargebacks.Id,
	}
	exists, err := repo.IsExistChargebacksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateChargebacks] failed checking chargebacks whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "chargebacks", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsChargebacks([]model.Chargebacks{*chargebacks}, fieldsInsert...)
	commandQuery := fmt.Sprintf(chargebacksQueries.insertChargebacks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateChargebacks] failed exec create chargebacks query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteChargebacksByID(ctx context.Context, primaryID model.ChargebacksPrimaryID) (err error) {
	exists, err := repo.IsExistChargebacksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteChargebacksByID] failed checking chargebacks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebacks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeChargebacksCompositePrimaryKeyWhere([]model.ChargebacksPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(chargebacksQueries.deleteChargebacks + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteChargebacksByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveChargebacksByFilter(ctx context.Context, filter model.Filter) (result []model.ChargebacksFilterResult, err error) {
	query, args, err := composeChargebacksFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveChargebacksByFilter] failed compose chargebacks filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveChargebacksByFilter] failed get chargebacks by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeChargebacksFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ChargebacksFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeChargebacksFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeChargebacksSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeChargebacksFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewChargebacksFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["chargeback_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"chargeback_code\"")
			selectedColumns["chargeback_code"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_ref_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_ref_id\"")
			selectedColumns["payment_ref_id"] = struct{}{}
		}
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_account_id\"")
			selectedColumns["provider_account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["disputed_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"disputed_amount\"")
			selectedColumns["disputed_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["chargeback_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"chargeback_status\"")
			selectedColumns["chargeback_status"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["opened_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"opened_at\"")
			selectedColumns["opened_at"] = struct{}{}
		}
		if _, selected := selectedColumns["closed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"closed_at\"")
			selectedColumns["closed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["due_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"due_at\"")
			selectedColumns["due_at"] = struct{}{}
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

type chargebacksFilterPlaceholder struct {
	index int
}

func (p *chargebacksFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeChargebacksFilterPredicate(filterField model.FilterField, placeholders *chargebacksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewChargebacksFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeChargebacksFilterSQLExpr(spec)
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

func composeChargebacksFilterGroup(group model.FilterGroup, placeholders *chargebacksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeChargebacksFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeChargebacksFilterGroup(child, placeholders, args, requiredJoins)
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

func composeChargebacksFilterWhereQueries(filter model.Filter, placeholders *chargebacksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeChargebacksFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeChargebacksFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeChargebacksFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateChargebacksFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeChargebacksSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeChargebacksFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := chargebacksFilterPlaceholder{index: 1}
	whereQueries, err := composeChargebacksFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewChargebacksFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeChargebacksFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeChargebacksSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"chargebacks\" base%s", strings.Join(selectColumns, ","), composeChargebacksFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistChargebacksByID(ctx context.Context, primaryID model.ChargebacksPrimaryID) (exists bool, err error) {
	whereQuery, params := composeChargebacksCompositePrimaryKeyWhere([]model.ChargebacksPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", chargebacksQueries.selectCountChargebacks, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistChargebacksByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveChargebacks(ctx context.Context, selectFields ...ChargebacksField) (chargebacksList model.ChargebacksList, err error) {
	var (
		defaultChargebacksSelectFields = defaultChargebacksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultChargebacksSelectFields = composeChargebacksSelectFields(selectFields...)
	}
	query := fmt.Sprintf(chargebacksQueries.selectChargebacks, defaultChargebacksSelectFields)

	err = repo.db.Read.SelectContext(ctx, &chargebacksList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveChargebacks] failed get chargebacks list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveChargebacksByID(ctx context.Context, primaryID model.ChargebacksPrimaryID, selectFields ...ChargebacksField) (chargebacks model.Chargebacks, err error) {
	var (
		defaultChargebacksSelectFields = defaultChargebacksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultChargebacksSelectFields = composeChargebacksSelectFields(selectFields...)
	}
	whereQry, params := composeChargebacksCompositePrimaryKeyWhere([]model.ChargebacksPrimaryID{primaryID})
	query := fmt.Sprintf(chargebacksQueries.selectChargebacks+" WHERE "+whereQry, defaultChargebacksSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &chargebacks, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("chargebacks with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveChargebacksByID] failed get chargebacks")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateChargebacksByID(ctx context.Context, primaryID model.ChargebacksPrimaryID, chargebacks *model.Chargebacks, chargebacksUpdateFields ...ChargebacksUpdateField) (err error) {
	exists, err := repo.IsExistChargebacksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebacks] failed checking chargebacks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("chargebacks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if chargebacks == nil {
		if len(chargebacksUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateChargebacksByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		chargebacks = &model.Chargebacks{}
	}
	var (
		defaultChargebacksUpdateFields = defaultChargebacksUpdateFields(*chargebacks)
		tempUpdateField                ChargebacksUpdateFieldList
		selectFields                   = NewChargebacksSelectFields()
	)
	if len(chargebacksUpdateFields) > 0 {
		for _, updateField := range chargebacksUpdateFields {
			if updateField.chargebacksField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultChargebacksUpdateFields = tempUpdateField
	}
	whereQuery, params := composeChargebacksCompositePrimaryKeyWhere([]model.ChargebacksPrimaryID{primaryID})
	fields, args := composeUpdateFieldsChargebacksCommand(defaultChargebacksUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(chargebacksQueries.updateChargebacks+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebacks] error when try to update chargebacks by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateChargebacksByFilter(ctx context.Context, filter model.Filter, chargebacksUpdateFields ...ChargebacksUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(chargebacksUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ChargebacksUpdateFieldList
		selectFields = NewChargebacksSelectFields()
	)
	for _, updateField := range chargebacksUpdateFields {
		if updateField.chargebacksField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsChargebacksCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := chargebacksFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeChargebacksFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"chargebacks\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebacksByFilter] error when try to update chargebacks by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateChargebacksByFilter] failed get rows affected")
	}
	return
}

var (
	chargebacksQueries = struct {
		selectChargebacks      string
		selectCountChargebacks string
		deleteChargebacks      string
		updateChargebacks      string
		insertChargebacks      string
	}{
		selectChargebacks:      "SELECT %s FROM \"chargebacks\"",
		selectCountChargebacks: "SELECT COUNT(\"id\") FROM \"chargebacks\"",
		deleteChargebacks:      "DELETE FROM \"chargebacks\"",
		updateChargebacks:      "UPDATE \"chargebacks\" SET %s ",
		insertChargebacks:      "INSERT INTO \"chargebacks\" %s VALUES %s",
	}
)

type ChargebacksRepository interface {
	CreateChargebacks(ctx context.Context, chargebacks *model.Chargebacks, fieldsInsert ...ChargebacksField) error
	BulkCreateChargebacks(ctx context.Context, chargebacksList []*model.Chargebacks, fieldsInsert ...ChargebacksField) error
	ResolveChargebacks(ctx context.Context, selectFields ...ChargebacksField) (model.ChargebacksList, error)
	ResolveChargebacksByID(ctx context.Context, primaryID model.ChargebacksPrimaryID, selectFields ...ChargebacksField) (model.Chargebacks, error)
	UpdateChargebacksByID(ctx context.Context, id model.ChargebacksPrimaryID, chargebacks *model.Chargebacks, chargebacksUpdateFields ...ChargebacksUpdateField) error
	UpdateChargebacksByFilter(ctx context.Context, filter model.Filter, chargebacksUpdateFields ...ChargebacksUpdateField) (rowsAffected int64, err error)
	BulkUpdateChargebacks(ctx context.Context, chargebacksListMap map[model.ChargebacksPrimaryID]*model.Chargebacks, ChargebackssMapUpdateFieldsRequest map[model.ChargebacksPrimaryID]ChargebacksUpdateFieldList) (err error)
	DeleteChargebacksByID(ctx context.Context, id model.ChargebacksPrimaryID) error
	BulkDeleteChargebacksByIDs(ctx context.Context, ids []model.ChargebacksPrimaryID) error
	ResolveChargebacksByFilter(ctx context.Context, filter model.Filter) (result []model.ChargebacksFilterResult, err error)
	IsExistChargebacksByIDs(ctx context.Context, ids []model.ChargebacksPrimaryID) (exists bool, notFoundIds []model.ChargebacksPrimaryID, err error)
	IsExistChargebacksByID(ctx context.Context, id model.ChargebacksPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
