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

func composeInsertFieldsAndParamsReserveHolds(reserveHoldsList []model.ReserveHolds, fieldsInsert ...ReserveHoldsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewReserveHoldsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, reserveHolds := range reserveHoldsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, reserveHolds.Id)
			case selectField.MerchantPartyId():
				args = append(args, reserveHolds.MerchantPartyId)
			case selectField.CurrencyCode():
				args = append(args, reserveHolds.CurrencyCode)
			case selectField.SourceType():
				args = append(args, reserveHolds.SourceType)
			case selectField.SourceId():
				args = append(args, reserveHolds.SourceId)
			case selectField.HoldType():
				args = append(args, reserveHolds.HoldType)
			case selectField.HoldStatus():
				args = append(args, reserveHolds.HoldStatus)
			case selectField.HoldAmount():
				args = append(args, reserveHolds.HoldAmount)
			case selectField.ReleasedAmount():
				args = append(args, reserveHolds.ReleasedAmount)
			case selectField.EffectiveAt():
				args = append(args, reserveHolds.EffectiveAt)
			case selectField.ReleaseAt():
				args = append(args, reserveHolds.ReleaseAt)
			case selectField.ReleasedAt():
				args = append(args, reserveHolds.ReleasedAt)
			case selectField.ReasonCode():
				args = append(args, reserveHolds.ReasonCode)
			case selectField.ReasonDetail():
				args = append(args, reserveHolds.ReasonDetail)
			case selectField.Metadata():
				args = append(args, reserveHolds.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, reserveHolds.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, reserveHolds.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, reserveHolds.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, reserveHolds.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, reserveHolds.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, reserveHolds.MetaDeletedBy)

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

func composeReserveHoldsCompositePrimaryKeyWhere(primaryIDs []model.ReserveHoldsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"reserve_holds\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultReserveHoldsSelectFields() string {
	fields := NewReserveHoldsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeReserveHoldsSelectFields(selectFields ...ReserveHoldsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ReserveHoldsField string
type ReserveHoldsFieldList []ReserveHoldsField

type ReserveHoldsSelectFields struct {
}

func (ss ReserveHoldsSelectFields) Id() ReserveHoldsField {
	return ReserveHoldsField("id")
}

func (ss ReserveHoldsSelectFields) MerchantPartyId() ReserveHoldsField {
	return ReserveHoldsField("merchant_party_id")
}

func (ss ReserveHoldsSelectFields) CurrencyCode() ReserveHoldsField {
	return ReserveHoldsField("currency_code")
}

func (ss ReserveHoldsSelectFields) SourceType() ReserveHoldsField {
	return ReserveHoldsField("source_type")
}

func (ss ReserveHoldsSelectFields) SourceId() ReserveHoldsField {
	return ReserveHoldsField("source_id")
}

func (ss ReserveHoldsSelectFields) HoldType() ReserveHoldsField {
	return ReserveHoldsField("hold_type")
}

func (ss ReserveHoldsSelectFields) HoldStatus() ReserveHoldsField {
	return ReserveHoldsField("hold_status")
}

func (ss ReserveHoldsSelectFields) HoldAmount() ReserveHoldsField {
	return ReserveHoldsField("hold_amount")
}

func (ss ReserveHoldsSelectFields) ReleasedAmount() ReserveHoldsField {
	return ReserveHoldsField("released_amount")
}

func (ss ReserveHoldsSelectFields) EffectiveAt() ReserveHoldsField {
	return ReserveHoldsField("effective_at")
}

func (ss ReserveHoldsSelectFields) ReleaseAt() ReserveHoldsField {
	return ReserveHoldsField("release_at")
}

func (ss ReserveHoldsSelectFields) ReleasedAt() ReserveHoldsField {
	return ReserveHoldsField("released_at")
}

func (ss ReserveHoldsSelectFields) ReasonCode() ReserveHoldsField {
	return ReserveHoldsField("reason_code")
}

func (ss ReserveHoldsSelectFields) ReasonDetail() ReserveHoldsField {
	return ReserveHoldsField("reason_detail")
}

func (ss ReserveHoldsSelectFields) Metadata() ReserveHoldsField {
	return ReserveHoldsField("metadata")
}

func (ss ReserveHoldsSelectFields) MetaCreatedAt() ReserveHoldsField {
	return ReserveHoldsField("meta_created_at")
}

func (ss ReserveHoldsSelectFields) MetaCreatedBy() ReserveHoldsField {
	return ReserveHoldsField("meta_created_by")
}

func (ss ReserveHoldsSelectFields) MetaUpdatedAt() ReserveHoldsField {
	return ReserveHoldsField("meta_updated_at")
}

func (ss ReserveHoldsSelectFields) MetaUpdatedBy() ReserveHoldsField {
	return ReserveHoldsField("meta_updated_by")
}

func (ss ReserveHoldsSelectFields) MetaDeletedAt() ReserveHoldsField {
	return ReserveHoldsField("meta_deleted_at")
}

func (ss ReserveHoldsSelectFields) MetaDeletedBy() ReserveHoldsField {
	return ReserveHoldsField("meta_deleted_by")
}

func (ss ReserveHoldsSelectFields) All() ReserveHoldsFieldList {
	return []ReserveHoldsField{
		ss.Id(),
		ss.MerchantPartyId(),
		ss.CurrencyCode(),
		ss.SourceType(),
		ss.SourceId(),
		ss.HoldType(),
		ss.HoldStatus(),
		ss.HoldAmount(),
		ss.ReleasedAmount(),
		ss.EffectiveAt(),
		ss.ReleaseAt(),
		ss.ReleasedAt(),
		ss.ReasonCode(),
		ss.ReasonDetail(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewReserveHoldsSelectFields() ReserveHoldsSelectFields {
	return ReserveHoldsSelectFields{}
}

type ReserveHoldsUpdateFieldOption struct {
	useIncrement bool
}
type ReserveHoldsUpdateField struct {
	reserveHoldsField ReserveHoldsField
	opt               ReserveHoldsUpdateFieldOption
	value             interface{}
}
type ReserveHoldsUpdateFieldList []ReserveHoldsUpdateField

func defaultReserveHoldsUpdateFieldOption() ReserveHoldsUpdateFieldOption {
	return ReserveHoldsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementReserveHoldsOption(useIncrement bool) func(*ReserveHoldsUpdateFieldOption) {
	return func(pcufo *ReserveHoldsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewReserveHoldsUpdateField(field ReserveHoldsField, val interface{}, opts ...func(*ReserveHoldsUpdateFieldOption)) ReserveHoldsUpdateField {
	defaultOpt := defaultReserveHoldsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ReserveHoldsUpdateField{
		reserveHoldsField: field,
		value:             val,
		opt:               defaultOpt,
	}
}
func defaultReserveHoldsUpdateFields(reserveHolds model.ReserveHolds) (reserveHoldsUpdateFieldList ReserveHoldsUpdateFieldList) {
	selectFields := NewReserveHoldsSelectFields()
	reserveHoldsUpdateFieldList = append(reserveHoldsUpdateFieldList,
		NewReserveHoldsUpdateField(selectFields.Id(), reserveHolds.Id),
		NewReserveHoldsUpdateField(selectFields.MerchantPartyId(), reserveHolds.MerchantPartyId),
		NewReserveHoldsUpdateField(selectFields.CurrencyCode(), reserveHolds.CurrencyCode),
		NewReserveHoldsUpdateField(selectFields.SourceType(), reserveHolds.SourceType),
		NewReserveHoldsUpdateField(selectFields.SourceId(), reserveHolds.SourceId),
		NewReserveHoldsUpdateField(selectFields.HoldType(), reserveHolds.HoldType),
		NewReserveHoldsUpdateField(selectFields.HoldStatus(), reserveHolds.HoldStatus),
		NewReserveHoldsUpdateField(selectFields.HoldAmount(), reserveHolds.HoldAmount),
		NewReserveHoldsUpdateField(selectFields.ReleasedAmount(), reserveHolds.ReleasedAmount),
		NewReserveHoldsUpdateField(selectFields.EffectiveAt(), reserveHolds.EffectiveAt),
		NewReserveHoldsUpdateField(selectFields.ReleaseAt(), reserveHolds.ReleaseAt),
		NewReserveHoldsUpdateField(selectFields.ReleasedAt(), reserveHolds.ReleasedAt),
		NewReserveHoldsUpdateField(selectFields.ReasonCode(), reserveHolds.ReasonCode),
		NewReserveHoldsUpdateField(selectFields.ReasonDetail(), reserveHolds.ReasonDetail),
		NewReserveHoldsUpdateField(selectFields.Metadata(), reserveHolds.Metadata),
		NewReserveHoldsUpdateField(selectFields.MetaCreatedAt(), reserveHolds.MetaCreatedAt),
		NewReserveHoldsUpdateField(selectFields.MetaCreatedBy(), reserveHolds.MetaCreatedBy),
		NewReserveHoldsUpdateField(selectFields.MetaUpdatedAt(), reserveHolds.MetaUpdatedAt),
		NewReserveHoldsUpdateField(selectFields.MetaUpdatedBy(), reserveHolds.MetaUpdatedBy),
		NewReserveHoldsUpdateField(selectFields.MetaDeletedAt(), reserveHolds.MetaDeletedAt),
		NewReserveHoldsUpdateField(selectFields.MetaDeletedBy(), reserveHolds.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsReserveHoldsCommand(reserveHoldsUpdateFieldList ReserveHoldsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range reserveHoldsUpdateFieldList {
		field := string(updateField.reserveHoldsField)
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

func (repo *RepositoryImpl) BulkCreateReserveHolds(ctx context.Context, reserveHoldsList []*model.ReserveHolds, fieldsInsert ...ReserveHoldsField) (err error) {
	var (
		fieldsStr             string
		valueListStr          []string
		argsList              []interface{}
		primaryIds            []model.ReserveHoldsPrimaryID
		reserveHoldsValueList []model.ReserveHolds
	)

	if len(fieldsInsert) == 0 {
		selectField := NewReserveHoldsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, reserveHolds := range reserveHoldsList {

		primaryIds = append(primaryIds, reserveHolds.ToReserveHoldsPrimaryID())

		reserveHoldsValueList = append(reserveHoldsValueList, *reserveHolds)
	}

	_, notFoundIds, err := repo.IsExistReserveHoldsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReserveHolds] failed checking reserveHolds whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ReserveHoldsPrimaryID{}
		mapNotFoundIds := map[model.ReserveHoldsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "reserveHolds", fmt.Sprintf("reserveHolds with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsReserveHolds(reserveHoldsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(reserveHoldsQueries.insertReserveHolds, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateReserveHolds] failed exec create reserveHolds query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteReserveHoldsByIDs(ctx context.Context, primaryIDs []model.ReserveHoldsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistReserveHoldsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReserveHoldsByIDs] failed checking reserveHolds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reserveHolds with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reserve_holds\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(reserveHoldsQueries.deleteReserveHolds + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReserveHoldsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteReserveHoldsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistReserveHoldsByIDs(ctx context.Context, ids []model.ReserveHoldsPrimaryID) (exists bool, notFoundIds []model.ReserveHoldsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"reserve_holds\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(reserveHoldsQueries.selectReserveHolds, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReserveHoldsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ReserveHoldsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReserveHoldsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ReserveHoldsPrimaryID]bool{}
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

// BulkUpdateReserveHolds is used to bulk update reserveHolds, by default it will update all field
// if want to update specific field, then fill reserveHoldssMapUpdateFieldsRequest else please fill reserveHoldssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateReserveHolds(ctx context.Context, reserveHoldssMap map[model.ReserveHoldsPrimaryID]*model.ReserveHolds, reserveHoldssMapUpdateFieldsRequest map[model.ReserveHoldsPrimaryID]ReserveHoldsUpdateFieldList) (err error) {
	if len(reserveHoldssMap) == 0 && len(reserveHoldssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		reserveHoldssMapUpdateField map[model.ReserveHoldsPrimaryID]ReserveHoldsUpdateFieldList = map[model.ReserveHoldsPrimaryID]ReserveHoldsUpdateFieldList{}
		asTableValues               string                                                      = "myvalues"
	)

	if len(reserveHoldssMap) > 0 {
		for id, reserveHolds := range reserveHoldssMap {
			if reserveHolds == nil {
				log.Error().Err(err).Msg("[BulkUpdateReserveHolds] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			reserveHoldssMapUpdateField[id] = defaultReserveHoldsUpdateFields(*reserveHolds)
		}
	} else {
		reserveHoldssMapUpdateField = reserveHoldssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateReserveHoldsQuery(reserveHoldssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistReserveHoldsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReserveHolds] failed checking reserveHolds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reserveHolds with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeReserveHoldsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"reserve_holds\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateReserveHolds] failed exec query")
	}
	return
}

type ReserveHoldsFieldParameter struct {
	param string
	args  []interface{}
}

func NewReserveHoldsFieldParameter(param string, args ...interface{}) ReserveHoldsFieldParameter {
	return ReserveHoldsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateReserveHoldsQuery(mapReserveHoldss map[model.ReserveHoldsPrimaryID]ReserveHoldsUpdateFieldList, asTableValues string) (primaryIDs []model.ReserveHoldsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ReserveHoldsPrimaryID]map[string]interface{}{}
	reserveHoldsSelectFields := NewReserveHoldsSelectFields()
	for id, updateFields := range mapReserveHoldss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.reserveHoldsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapReserveHoldss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetReserveHoldsFieldType(updateField.reserveHoldsField)))
			args = append(args, fields[string(updateField.reserveHoldsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.reserveHoldsField))
		if updateField.reserveHoldsField == reserveHoldsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.reserveHoldsField, asTableValues, updateField.reserveHoldsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.reserveHoldsField,
				"\"reserve_holds\"", updateField.reserveHoldsField,
				asTableValues, updateField.reserveHoldsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeReserveHoldsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ReserveHoldsPrimaryID, asTableValue string) (whereQry string) {
	reserveHoldsSelectFields := NewReserveHoldsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"reserve_holds\".\"id\" = %s.\"id\"::"+GetReserveHoldsFieldType(reserveHoldsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetReserveHoldsFieldType(reserveHoldsField ReserveHoldsField) string {
	selectReserveHoldsFields := NewReserveHoldsSelectFields()
	switch reserveHoldsField {

	case selectReserveHoldsFields.Id():
		return "uuid"

	case selectReserveHoldsFields.MerchantPartyId():
		return "uuid"

	case selectReserveHoldsFields.CurrencyCode():
		return "text"

	case selectReserveHoldsFields.SourceType():
		return "text"

	case selectReserveHoldsFields.SourceId():
		return "uuid"

	case selectReserveHoldsFields.HoldType():
		return "text"

	case selectReserveHoldsFields.HoldStatus():
		return "hold_status_enum"

	case selectReserveHoldsFields.HoldAmount():
		return "numeric"

	case selectReserveHoldsFields.ReleasedAmount():
		return "numeric"

	case selectReserveHoldsFields.EffectiveAt():
		return "timestamptz"

	case selectReserveHoldsFields.ReleaseAt():
		return "timestamptz"

	case selectReserveHoldsFields.ReleasedAt():
		return "timestamptz"

	case selectReserveHoldsFields.ReasonCode():
		return "text"

	case selectReserveHoldsFields.ReasonDetail():
		return "text"

	case selectReserveHoldsFields.Metadata():
		return "jsonb"

	case selectReserveHoldsFields.MetaCreatedAt():
		return "timestamptz"

	case selectReserveHoldsFields.MetaCreatedBy():
		return "uuid"

	case selectReserveHoldsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectReserveHoldsFields.MetaUpdatedBy():
		return "uuid"

	case selectReserveHoldsFields.MetaDeletedAt():
		return "timestamptz"

	case selectReserveHoldsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateReserveHolds(ctx context.Context, reserveHolds *model.ReserveHolds, fieldsInsert ...ReserveHoldsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewReserveHoldsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ReserveHoldsPrimaryID{
		Id: reserveHolds.Id,
	}
	exists, err := repo.IsExistReserveHoldsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReserveHolds] failed checking reserveHolds whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "reserveHolds", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsReserveHolds([]model.ReserveHolds{*reserveHolds}, fieldsInsert...)
	commandQuery := fmt.Sprintf(reserveHoldsQueries.insertReserveHolds, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateReserveHolds] failed exec create reserveHolds query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteReserveHoldsByID(ctx context.Context, primaryID model.ReserveHoldsPrimaryID) (err error) {
	exists, err := repo.IsExistReserveHoldsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReserveHoldsByID] failed checking reserveHolds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reserveHolds with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeReserveHoldsCompositePrimaryKeyWhere([]model.ReserveHoldsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(reserveHoldsQueries.deleteReserveHolds + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteReserveHoldsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveReserveHoldsByFilter(ctx context.Context, filter model.Filter) (result []model.ReserveHoldsFilterResult, err error) {
	query, args, err := composeReserveHoldsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReserveHoldsByFilter] failed compose reserveHolds filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReserveHoldsByFilter] failed get reserveHolds by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeReserveHoldsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ReserveHoldsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeReserveHoldsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeReserveHoldsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeReserveHoldsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 21 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewReserveHoldsFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["merchant_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"merchant_party_id\"")
			selectedColumns["merchant_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["hold_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"hold_type\"")
			selectedColumns["hold_type"] = struct{}{}
		}
		if _, selected := selectedColumns["hold_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"hold_status\"")
			selectedColumns["hold_status"] = struct{}{}
		}
		if _, selected := selectedColumns["hold_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"hold_amount\"")
			selectedColumns["hold_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["released_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"released_amount\"")
			selectedColumns["released_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["effective_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"effective_at\"")
			selectedColumns["effective_at"] = struct{}{}
		}
		if _, selected := selectedColumns["release_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"release_at\"")
			selectedColumns["release_at"] = struct{}{}
		}
		if _, selected := selectedColumns["released_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"released_at\"")
			selectedColumns["released_at"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_detail"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_detail\"")
			selectedColumns["reason_detail"] = struct{}{}
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

type reserveHoldsFilterPlaceholder struct {
	index int
}

func (p *reserveHoldsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeReserveHoldsFilterPredicate(filterField model.FilterField, placeholders *reserveHoldsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewReserveHoldsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeReserveHoldsFilterSQLExpr(spec)
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

func composeReserveHoldsFilterGroup(group model.FilterGroup, placeholders *reserveHoldsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeReserveHoldsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeReserveHoldsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeReserveHoldsFilterWhereQueries(filter model.Filter, placeholders *reserveHoldsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeReserveHoldsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeReserveHoldsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeReserveHoldsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateReserveHoldsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeReserveHoldsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeReserveHoldsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := reserveHoldsFilterPlaceholder{index: 1}
	whereQueries, err := composeReserveHoldsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewReserveHoldsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeReserveHoldsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeReserveHoldsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"reserve_holds\" base%s", strings.Join(selectColumns, ","), composeReserveHoldsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistReserveHoldsByID(ctx context.Context, primaryID model.ReserveHoldsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeReserveHoldsCompositePrimaryKeyWhere([]model.ReserveHoldsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", reserveHoldsQueries.selectCountReserveHolds, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistReserveHoldsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReserveHolds(ctx context.Context, selectFields ...ReserveHoldsField) (reserveHoldsList model.ReserveHoldsList, err error) {
	var (
		defaultReserveHoldsSelectFields = defaultReserveHoldsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReserveHoldsSelectFields = composeReserveHoldsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(reserveHoldsQueries.selectReserveHolds, defaultReserveHoldsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &reserveHoldsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveReserveHolds] failed get reserveHolds list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveReserveHoldsByID(ctx context.Context, primaryID model.ReserveHoldsPrimaryID, selectFields ...ReserveHoldsField) (reserveHolds model.ReserveHolds, err error) {
	var (
		defaultReserveHoldsSelectFields = defaultReserveHoldsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultReserveHoldsSelectFields = composeReserveHoldsSelectFields(selectFields...)
	}
	whereQry, params := composeReserveHoldsCompositePrimaryKeyWhere([]model.ReserveHoldsPrimaryID{primaryID})
	query := fmt.Sprintf(reserveHoldsQueries.selectReserveHolds+" WHERE "+whereQry, defaultReserveHoldsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &reserveHolds, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("reserveHolds with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveReserveHoldsByID] failed get reserveHolds")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateReserveHoldsByID(ctx context.Context, primaryID model.ReserveHoldsPrimaryID, reserveHolds *model.ReserveHolds, reserveHoldsUpdateFields ...ReserveHoldsUpdateField) (err error) {
	exists, err := repo.IsExistReserveHoldsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReserveHolds] failed checking reserveHolds whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("reserveHolds with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if reserveHolds == nil {
		if len(reserveHoldsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateReserveHoldsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		reserveHolds = &model.ReserveHolds{}
	}
	var (
		defaultReserveHoldsUpdateFields = defaultReserveHoldsUpdateFields(*reserveHolds)
		tempUpdateField                 ReserveHoldsUpdateFieldList
		selectFields                    = NewReserveHoldsSelectFields()
	)
	if len(reserveHoldsUpdateFields) > 0 {
		for _, updateField := range reserveHoldsUpdateFields {
			if updateField.reserveHoldsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultReserveHoldsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeReserveHoldsCompositePrimaryKeyWhere([]model.ReserveHoldsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsReserveHoldsCommand(defaultReserveHoldsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(reserveHoldsQueries.updateReserveHolds+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReserveHolds] error when try to update reserveHolds by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateReserveHoldsByFilter(ctx context.Context, filter model.Filter, reserveHoldsUpdateFields ...ReserveHoldsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(reserveHoldsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ReserveHoldsUpdateFieldList
		selectFields = NewReserveHoldsSelectFields()
	)
	for _, updateField := range reserveHoldsUpdateFields {
		if updateField.reserveHoldsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsReserveHoldsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := reserveHoldsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeReserveHoldsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"reserve_holds\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReserveHoldsByFilter] error when try to update reserveHolds by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateReserveHoldsByFilter] failed get rows affected")
	}
	return
}

var (
	reserveHoldsQueries = struct {
		selectReserveHolds      string
		selectCountReserveHolds string
		deleteReserveHolds      string
		updateReserveHolds      string
		insertReserveHolds      string
	}{
		selectReserveHolds:      "SELECT %s FROM \"reserve_holds\"",
		selectCountReserveHolds: "SELECT COUNT(\"id\") FROM \"reserve_holds\"",
		deleteReserveHolds:      "DELETE FROM \"reserve_holds\"",
		updateReserveHolds:      "UPDATE \"reserve_holds\" SET %s ",
		insertReserveHolds:      "INSERT INTO \"reserve_holds\" %s VALUES %s",
	}
)

type ReserveHoldsRepository interface {
	CreateReserveHolds(ctx context.Context, reserveHolds *model.ReserveHolds, fieldsInsert ...ReserveHoldsField) error
	BulkCreateReserveHolds(ctx context.Context, reserveHoldsList []*model.ReserveHolds, fieldsInsert ...ReserveHoldsField) error
	ResolveReserveHolds(ctx context.Context, selectFields ...ReserveHoldsField) (model.ReserveHoldsList, error)
	ResolveReserveHoldsByID(ctx context.Context, primaryID model.ReserveHoldsPrimaryID, selectFields ...ReserveHoldsField) (model.ReserveHolds, error)
	UpdateReserveHoldsByID(ctx context.Context, id model.ReserveHoldsPrimaryID, reserveHolds *model.ReserveHolds, reserveHoldsUpdateFields ...ReserveHoldsUpdateField) error
	UpdateReserveHoldsByFilter(ctx context.Context, filter model.Filter, reserveHoldsUpdateFields ...ReserveHoldsUpdateField) (rowsAffected int64, err error)
	BulkUpdateReserveHolds(ctx context.Context, reserveHoldsListMap map[model.ReserveHoldsPrimaryID]*model.ReserveHolds, ReserveHoldssMapUpdateFieldsRequest map[model.ReserveHoldsPrimaryID]ReserveHoldsUpdateFieldList) (err error)
	DeleteReserveHoldsByID(ctx context.Context, id model.ReserveHoldsPrimaryID) error
	BulkDeleteReserveHoldsByIDs(ctx context.Context, ids []model.ReserveHoldsPrimaryID) error
	ResolveReserveHoldsByFilter(ctx context.Context, filter model.Filter) (result []model.ReserveHoldsFilterResult, err error)
	IsExistReserveHoldsByIDs(ctx context.Context, ids []model.ReserveHoldsPrimaryID) (exists bool, notFoundIds []model.ReserveHoldsPrimaryID, err error)
	IsExistReserveHoldsByID(ctx context.Context, id model.ReserveHoldsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
