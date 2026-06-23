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

func composeInsertFieldsAndParamsDisputeCases(disputeCasesList []model.DisputeCases, fieldsInsert ...DisputeCasesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewDisputeCasesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, disputeCases := range disputeCasesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, disputeCases.Id)
			case selectField.DisputeCode():
				args = append(args, disputeCases.DisputeCode)
			case selectField.BookId():
				args = append(args, disputeCases.BookId)
			case selectField.SourceType():
				args = append(args, disputeCases.SourceType)
			case selectField.SourceId():
				args = append(args, disputeCases.SourceId)
			case selectField.MerchantPartyId():
				args = append(args, disputeCases.MerchantPartyId)
			case selectField.CustomerPartyId():
				args = append(args, disputeCases.CustomerPartyId)
			case selectField.ProviderCode():
				args = append(args, disputeCases.ProviderCode)
			case selectField.ProviderDisputeRef():
				args = append(args, disputeCases.ProviderDisputeRef)
			case selectField.IdempotencyKey():
				args = append(args, disputeCases.IdempotencyKey)
			case selectField.CurrencyCode():
				args = append(args, disputeCases.CurrencyCode)
			case selectField.DisputedAmount():
				args = append(args, disputeCases.DisputedAmount)
			case selectField.DisputeReasonCode():
				args = append(args, disputeCases.DisputeReasonCode)
			case selectField.DisputeStatus():
				args = append(args, disputeCases.DisputeStatus)
			case selectField.OpenedAt():
				args = append(args, disputeCases.OpenedAt)
			case selectField.DueAt():
				args = append(args, disputeCases.DueAt)
			case selectField.ClosedAt():
				args = append(args, disputeCases.ClosedAt)
			case selectField.Metadata():
				args = append(args, disputeCases.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, disputeCases.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, disputeCases.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, disputeCases.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, disputeCases.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, disputeCases.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, disputeCases.MetaDeletedBy)

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

func composeDisputeCasesCompositePrimaryKeyWhere(primaryIDs []model.DisputeCasesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"dispute_cases\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultDisputeCasesSelectFields() string {
	fields := NewDisputeCasesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeDisputeCasesSelectFields(selectFields ...DisputeCasesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type DisputeCasesField string
type DisputeCasesFieldList []DisputeCasesField

type DisputeCasesSelectFields struct {
}

func (ss DisputeCasesSelectFields) Id() DisputeCasesField {
	return DisputeCasesField("id")
}

func (ss DisputeCasesSelectFields) DisputeCode() DisputeCasesField {
	return DisputeCasesField("dispute_code")
}

func (ss DisputeCasesSelectFields) BookId() DisputeCasesField {
	return DisputeCasesField("book_id")
}

func (ss DisputeCasesSelectFields) SourceType() DisputeCasesField {
	return DisputeCasesField("source_type")
}

func (ss DisputeCasesSelectFields) SourceId() DisputeCasesField {
	return DisputeCasesField("source_id")
}

func (ss DisputeCasesSelectFields) MerchantPartyId() DisputeCasesField {
	return DisputeCasesField("merchant_party_id")
}

func (ss DisputeCasesSelectFields) CustomerPartyId() DisputeCasesField {
	return DisputeCasesField("customer_party_id")
}

func (ss DisputeCasesSelectFields) ProviderCode() DisputeCasesField {
	return DisputeCasesField("provider_code")
}

func (ss DisputeCasesSelectFields) ProviderDisputeRef() DisputeCasesField {
	return DisputeCasesField("provider_dispute_ref")
}

func (ss DisputeCasesSelectFields) IdempotencyKey() DisputeCasesField {
	return DisputeCasesField("idempotency_key")
}

func (ss DisputeCasesSelectFields) CurrencyCode() DisputeCasesField {
	return DisputeCasesField("currency_code")
}

func (ss DisputeCasesSelectFields) DisputedAmount() DisputeCasesField {
	return DisputeCasesField("disputed_amount")
}

func (ss DisputeCasesSelectFields) DisputeReasonCode() DisputeCasesField {
	return DisputeCasesField("dispute_reason_code")
}

func (ss DisputeCasesSelectFields) DisputeStatus() DisputeCasesField {
	return DisputeCasesField("dispute_status")
}

func (ss DisputeCasesSelectFields) OpenedAt() DisputeCasesField {
	return DisputeCasesField("opened_at")
}

func (ss DisputeCasesSelectFields) DueAt() DisputeCasesField {
	return DisputeCasesField("due_at")
}

func (ss DisputeCasesSelectFields) ClosedAt() DisputeCasesField {
	return DisputeCasesField("closed_at")
}

func (ss DisputeCasesSelectFields) Metadata() DisputeCasesField {
	return DisputeCasesField("metadata")
}

func (ss DisputeCasesSelectFields) MetaCreatedAt() DisputeCasesField {
	return DisputeCasesField("meta_created_at")
}

func (ss DisputeCasesSelectFields) MetaCreatedBy() DisputeCasesField {
	return DisputeCasesField("meta_created_by")
}

func (ss DisputeCasesSelectFields) MetaUpdatedAt() DisputeCasesField {
	return DisputeCasesField("meta_updated_at")
}

func (ss DisputeCasesSelectFields) MetaUpdatedBy() DisputeCasesField {
	return DisputeCasesField("meta_updated_by")
}

func (ss DisputeCasesSelectFields) MetaDeletedAt() DisputeCasesField {
	return DisputeCasesField("meta_deleted_at")
}

func (ss DisputeCasesSelectFields) MetaDeletedBy() DisputeCasesField {
	return DisputeCasesField("meta_deleted_by")
}

func (ss DisputeCasesSelectFields) All() DisputeCasesFieldList {
	return []DisputeCasesField{
		ss.Id(),
		ss.DisputeCode(),
		ss.BookId(),
		ss.SourceType(),
		ss.SourceId(),
		ss.MerchantPartyId(),
		ss.CustomerPartyId(),
		ss.ProviderCode(),
		ss.ProviderDisputeRef(),
		ss.IdempotencyKey(),
		ss.CurrencyCode(),
		ss.DisputedAmount(),
		ss.DisputeReasonCode(),
		ss.DisputeStatus(),
		ss.OpenedAt(),
		ss.DueAt(),
		ss.ClosedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewDisputeCasesSelectFields() DisputeCasesSelectFields {
	return DisputeCasesSelectFields{}
}

type DisputeCasesUpdateFieldOption struct {
	useIncrement bool
}
type DisputeCasesUpdateField struct {
	disputeCasesField DisputeCasesField
	opt               DisputeCasesUpdateFieldOption
	value             interface{}
}
type DisputeCasesUpdateFieldList []DisputeCasesUpdateField

func defaultDisputeCasesUpdateFieldOption() DisputeCasesUpdateFieldOption {
	return DisputeCasesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementDisputeCasesOption(useIncrement bool) func(*DisputeCasesUpdateFieldOption) {
	return func(pcufo *DisputeCasesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewDisputeCasesUpdateField(field DisputeCasesField, val interface{}, opts ...func(*DisputeCasesUpdateFieldOption)) DisputeCasesUpdateField {
	defaultOpt := defaultDisputeCasesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return DisputeCasesUpdateField{
		disputeCasesField: field,
		value:             val,
		opt:               defaultOpt,
	}
}
func defaultDisputeCasesUpdateFields(disputeCases model.DisputeCases) (disputeCasesUpdateFieldList DisputeCasesUpdateFieldList) {
	selectFields := NewDisputeCasesSelectFields()
	disputeCasesUpdateFieldList = append(disputeCasesUpdateFieldList,
		NewDisputeCasesUpdateField(selectFields.Id(), disputeCases.Id),
		NewDisputeCasesUpdateField(selectFields.DisputeCode(), disputeCases.DisputeCode),
		NewDisputeCasesUpdateField(selectFields.BookId(), disputeCases.BookId),
		NewDisputeCasesUpdateField(selectFields.SourceType(), disputeCases.SourceType),
		NewDisputeCasesUpdateField(selectFields.SourceId(), disputeCases.SourceId),
		NewDisputeCasesUpdateField(selectFields.MerchantPartyId(), disputeCases.MerchantPartyId),
		NewDisputeCasesUpdateField(selectFields.CustomerPartyId(), disputeCases.CustomerPartyId),
		NewDisputeCasesUpdateField(selectFields.ProviderCode(), disputeCases.ProviderCode),
		NewDisputeCasesUpdateField(selectFields.ProviderDisputeRef(), disputeCases.ProviderDisputeRef),
		NewDisputeCasesUpdateField(selectFields.IdempotencyKey(), disputeCases.IdempotencyKey),
		NewDisputeCasesUpdateField(selectFields.CurrencyCode(), disputeCases.CurrencyCode),
		NewDisputeCasesUpdateField(selectFields.DisputedAmount(), disputeCases.DisputedAmount),
		NewDisputeCasesUpdateField(selectFields.DisputeReasonCode(), disputeCases.DisputeReasonCode),
		NewDisputeCasesUpdateField(selectFields.DisputeStatus(), disputeCases.DisputeStatus),
		NewDisputeCasesUpdateField(selectFields.OpenedAt(), disputeCases.OpenedAt),
		NewDisputeCasesUpdateField(selectFields.DueAt(), disputeCases.DueAt),
		NewDisputeCasesUpdateField(selectFields.ClosedAt(), disputeCases.ClosedAt),
		NewDisputeCasesUpdateField(selectFields.Metadata(), disputeCases.Metadata),
		NewDisputeCasesUpdateField(selectFields.MetaCreatedAt(), disputeCases.MetaCreatedAt),
		NewDisputeCasesUpdateField(selectFields.MetaCreatedBy(), disputeCases.MetaCreatedBy),
		NewDisputeCasesUpdateField(selectFields.MetaUpdatedAt(), disputeCases.MetaUpdatedAt),
		NewDisputeCasesUpdateField(selectFields.MetaUpdatedBy(), disputeCases.MetaUpdatedBy),
		NewDisputeCasesUpdateField(selectFields.MetaDeletedAt(), disputeCases.MetaDeletedAt),
		NewDisputeCasesUpdateField(selectFields.MetaDeletedBy(), disputeCases.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsDisputeCasesCommand(disputeCasesUpdateFieldList DisputeCasesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range disputeCasesUpdateFieldList {
		field := string(updateField.disputeCasesField)
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

func (repo *RepositoryImpl) BulkCreateDisputeCases(ctx context.Context, disputeCasesList []*model.DisputeCases, fieldsInsert ...DisputeCasesField) (err error) {
	var (
		fieldsStr             string
		valueListStr          []string
		argsList              []interface{}
		primaryIds            []model.DisputeCasesPrimaryID
		disputeCasesValueList []model.DisputeCases
	)

	if len(fieldsInsert) == 0 {
		selectField := NewDisputeCasesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, disputeCases := range disputeCasesList {

		primaryIds = append(primaryIds, disputeCases.ToDisputeCasesPrimaryID())

		disputeCasesValueList = append(disputeCasesValueList, *disputeCases)
	}

	_, notFoundIds, err := repo.IsExistDisputeCasesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputeCases] failed checking disputeCases whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.DisputeCasesPrimaryID{}
		mapNotFoundIds := map[model.DisputeCasesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "disputeCases", fmt.Sprintf("disputeCases with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsDisputeCases(disputeCasesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(disputeCasesQueries.insertDisputeCases, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputeCases] failed exec create disputeCases query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteDisputeCasesByIDs(ctx context.Context, primaryIDs []model.DisputeCasesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistDisputeCasesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeCasesByIDs] failed checking disputeCases whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeCases with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"dispute_cases\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(disputeCasesQueries.deleteDisputeCases + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeCasesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeCasesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistDisputeCasesByIDs(ctx context.Context, ids []model.DisputeCasesPrimaryID) (exists bool, notFoundIds []model.DisputeCasesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"dispute_cases\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(disputeCasesQueries.selectDisputeCases, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeCasesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.DisputeCasesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeCasesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.DisputeCasesPrimaryID]bool{}
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

// BulkUpdateDisputeCases is used to bulk update disputeCases, by default it will update all field
// if want to update specific field, then fill disputeCasessMapUpdateFieldsRequest else please fill disputeCasessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateDisputeCases(ctx context.Context, disputeCasessMap map[model.DisputeCasesPrimaryID]*model.DisputeCases, disputeCasessMapUpdateFieldsRequest map[model.DisputeCasesPrimaryID]DisputeCasesUpdateFieldList) (err error) {
	if len(disputeCasessMap) == 0 && len(disputeCasessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		disputeCasessMapUpdateField map[model.DisputeCasesPrimaryID]DisputeCasesUpdateFieldList = map[model.DisputeCasesPrimaryID]DisputeCasesUpdateFieldList{}
		asTableValues               string                                                      = "myvalues"
	)

	if len(disputeCasessMap) > 0 {
		for id, disputeCases := range disputeCasessMap {
			if disputeCases == nil {
				log.Error().Err(err).Msg("[BulkUpdateDisputeCases] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			disputeCasessMapUpdateField[id] = defaultDisputeCasesUpdateFields(*disputeCases)
		}
	} else {
		disputeCasessMapUpdateField = disputeCasessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateDisputeCasesQuery(disputeCasessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistDisputeCasesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputeCases] failed checking disputeCases whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeCases with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeDisputeCasesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"dispute_cases\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputeCases] failed exec query")
	}
	return
}

type DisputeCasesFieldParameter struct {
	param string
	args  []interface{}
}

func NewDisputeCasesFieldParameter(param string, args ...interface{}) DisputeCasesFieldParameter {
	return DisputeCasesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateDisputeCasesQuery(mapDisputeCasess map[model.DisputeCasesPrimaryID]DisputeCasesUpdateFieldList, asTableValues string) (primaryIDs []model.DisputeCasesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.DisputeCasesPrimaryID]map[string]interface{}{}
	disputeCasesSelectFields := NewDisputeCasesSelectFields()
	for id, updateFields := range mapDisputeCasess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.disputeCasesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapDisputeCasess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetDisputeCasesFieldType(updateField.disputeCasesField)))
			args = append(args, fields[string(updateField.disputeCasesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.disputeCasesField))
		if updateField.disputeCasesField == disputeCasesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.disputeCasesField, asTableValues, updateField.disputeCasesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.disputeCasesField,
				"\"dispute_cases\"", updateField.disputeCasesField,
				asTableValues, updateField.disputeCasesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeDisputeCasesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.DisputeCasesPrimaryID, asTableValue string) (whereQry string) {
	disputeCasesSelectFields := NewDisputeCasesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"dispute_cases\".\"id\" = %s.\"id\"::"+GetDisputeCasesFieldType(disputeCasesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetDisputeCasesFieldType(disputeCasesField DisputeCasesField) string {
	selectDisputeCasesFields := NewDisputeCasesSelectFields()
	switch disputeCasesField {

	case selectDisputeCasesFields.Id():
		return "uuid"

	case selectDisputeCasesFields.DisputeCode():
		return "text"

	case selectDisputeCasesFields.BookId():
		return "uuid"

	case selectDisputeCasesFields.SourceType():
		return "text"

	case selectDisputeCasesFields.SourceId():
		return "uuid"

	case selectDisputeCasesFields.MerchantPartyId():
		return "uuid"

	case selectDisputeCasesFields.CustomerPartyId():
		return "uuid"

	case selectDisputeCasesFields.ProviderCode():
		return "text"

	case selectDisputeCasesFields.ProviderDisputeRef():
		return "text"

	case selectDisputeCasesFields.IdempotencyKey():
		return "text"

	case selectDisputeCasesFields.CurrencyCode():
		return "text"

	case selectDisputeCasesFields.DisputedAmount():
		return "numeric"

	case selectDisputeCasesFields.DisputeReasonCode():
		return "text"

	case selectDisputeCasesFields.DisputeStatus():
		return "dispute_status_enum"

	case selectDisputeCasesFields.OpenedAt():
		return "timestamptz"

	case selectDisputeCasesFields.DueAt():
		return "timestamptz"

	case selectDisputeCasesFields.ClosedAt():
		return "timestamptz"

	case selectDisputeCasesFields.Metadata():
		return "jsonb"

	case selectDisputeCasesFields.MetaCreatedAt():
		return "timestamptz"

	case selectDisputeCasesFields.MetaCreatedBy():
		return "uuid"

	case selectDisputeCasesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectDisputeCasesFields.MetaUpdatedBy():
		return "uuid"

	case selectDisputeCasesFields.MetaDeletedAt():
		return "timestamptz"

	case selectDisputeCasesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateDisputeCases(ctx context.Context, disputeCases *model.DisputeCases, fieldsInsert ...DisputeCasesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewDisputeCasesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.DisputeCasesPrimaryID{
		Id: disputeCases.Id,
	}
	exists, err := repo.IsExistDisputeCasesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputeCases] failed checking disputeCases whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "disputeCases", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsDisputeCases([]model.DisputeCases{*disputeCases}, fieldsInsert...)
	commandQuery := fmt.Sprintf(disputeCasesQueries.insertDisputeCases, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputeCases] failed exec create disputeCases query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteDisputeCasesByID(ctx context.Context, primaryID model.DisputeCasesPrimaryID) (err error) {
	exists, err := repo.IsExistDisputeCasesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputeCasesByID] failed checking disputeCases whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeCases with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeDisputeCasesCompositePrimaryKeyWhere([]model.DisputeCasesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(disputeCasesQueries.deleteDisputeCases + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputeCasesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeCasesByFilter(ctx context.Context, filter model.Filter) (result []model.DisputeCasesFilterResult, err error) {
	query, args, err := composeDisputeCasesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeCasesByFilter] failed compose disputeCases filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeCasesByFilter] failed get disputeCases by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeDisputeCasesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.DisputeCasesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeDisputeCasesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeDisputeCasesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeDisputeCasesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 24 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewDisputeCasesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 24+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["dispute_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"dispute_code\"")
			selectedColumns["dispute_code"] = struct{}{}
		}
		if _, selected := selectedColumns["book_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_id\"")
			selectedColumns["book_id"] = struct{}{}
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
		if _, selected := selectedColumns["customer_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"customer_party_id\"")
			selectedColumns["customer_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_code\"")
			selectedColumns["provider_code"] = struct{}{}
		}
		if _, selected := selectedColumns["provider_dispute_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"provider_dispute_ref\"")
			selectedColumns["provider_dispute_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["disputed_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"disputed_amount\"")
			selectedColumns["disputed_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["dispute_reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"dispute_reason_code\"")
			selectedColumns["dispute_reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["dispute_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"dispute_status\"")
			selectedColumns["dispute_status"] = struct{}{}
		}
		if _, selected := selectedColumns["opened_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"opened_at\"")
			selectedColumns["opened_at"] = struct{}{}
		}
		if _, selected := selectedColumns["due_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"due_at\"")
			selectedColumns["due_at"] = struct{}{}
		}
		if _, selected := selectedColumns["closed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"closed_at\"")
			selectedColumns["closed_at"] = struct{}{}
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

type disputeCasesFilterPlaceholder struct {
	index int
}

func (p *disputeCasesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeDisputeCasesFilterPredicate(filterField model.FilterField, placeholders *disputeCasesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewDisputeCasesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeDisputeCasesFilterSQLExpr(spec)
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

func composeDisputeCasesFilterGroup(group model.FilterGroup, placeholders *disputeCasesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeDisputeCasesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeDisputeCasesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeDisputeCasesFilterWhereQueries(filter model.Filter, placeholders *disputeCasesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeDisputeCasesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeDisputeCasesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeDisputeCasesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateDisputeCasesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeDisputeCasesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeDisputeCasesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := disputeCasesFilterPlaceholder{index: 1}
	whereQueries, err := composeDisputeCasesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewDisputeCasesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeDisputeCasesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeDisputeCasesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"dispute_cases\" base%s", strings.Join(selectColumns, ","), composeDisputeCasesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistDisputeCasesByID(ctx context.Context, primaryID model.DisputeCasesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeDisputeCasesCompositePrimaryKeyWhere([]model.DisputeCasesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", disputeCasesQueries.selectCountDisputeCases, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeCasesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeCases(ctx context.Context, selectFields ...DisputeCasesField) (disputeCasesList model.DisputeCasesList, err error) {
	var (
		defaultDisputeCasesSelectFields = defaultDisputeCasesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputeCasesSelectFields = composeDisputeCasesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(disputeCasesQueries.selectDisputeCases, defaultDisputeCasesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &disputeCasesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeCases] failed get disputeCases list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeCasesByID(ctx context.Context, primaryID model.DisputeCasesPrimaryID, selectFields ...DisputeCasesField) (disputeCases model.DisputeCases, err error) {
	var (
		defaultDisputeCasesSelectFields = defaultDisputeCasesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputeCasesSelectFields = composeDisputeCasesSelectFields(selectFields...)
	}
	whereQry, params := composeDisputeCasesCompositePrimaryKeyWhere([]model.DisputeCasesPrimaryID{primaryID})
	query := fmt.Sprintf(disputeCasesQueries.selectDisputeCases+" WHERE "+whereQry, defaultDisputeCasesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &disputeCases, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("disputeCases with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveDisputeCasesByID] failed get disputeCases")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateDisputeCasesByID(ctx context.Context, primaryID model.DisputeCasesPrimaryID, disputeCases *model.DisputeCases, disputeCasesUpdateFields ...DisputeCasesUpdateField) (err error) {
	exists, err := repo.IsExistDisputeCasesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeCases] failed checking disputeCases whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeCases with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if disputeCases == nil {
		if len(disputeCasesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateDisputeCasesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		disputeCases = &model.DisputeCases{}
	}
	var (
		defaultDisputeCasesUpdateFields = defaultDisputeCasesUpdateFields(*disputeCases)
		tempUpdateField                 DisputeCasesUpdateFieldList
		selectFields                    = NewDisputeCasesSelectFields()
	)
	if len(disputeCasesUpdateFields) > 0 {
		for _, updateField := range disputeCasesUpdateFields {
			if updateField.disputeCasesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultDisputeCasesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeDisputeCasesCompositePrimaryKeyWhere([]model.DisputeCasesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsDisputeCasesCommand(defaultDisputeCasesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(disputeCasesQueries.updateDisputeCases+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeCases] error when try to update disputeCases by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateDisputeCasesByFilter(ctx context.Context, filter model.Filter, disputeCasesUpdateFields ...DisputeCasesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(disputeCasesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields DisputeCasesUpdateFieldList
		selectFields = NewDisputeCasesSelectFields()
	)
	for _, updateField := range disputeCasesUpdateFields {
		if updateField.disputeCasesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsDisputeCasesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := disputeCasesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeDisputeCasesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"dispute_cases\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeCasesByFilter] error when try to update disputeCases by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeCasesByFilter] failed get rows affected")
	}
	return
}

var (
	disputeCasesQueries = struct {
		selectDisputeCases      string
		selectCountDisputeCases string
		deleteDisputeCases      string
		updateDisputeCases      string
		insertDisputeCases      string
	}{
		selectDisputeCases:      "SELECT %s FROM \"dispute_cases\"",
		selectCountDisputeCases: "SELECT COUNT(\"id\") FROM \"dispute_cases\"",
		deleteDisputeCases:      "DELETE FROM \"dispute_cases\"",
		updateDisputeCases:      "UPDATE \"dispute_cases\" SET %s ",
		insertDisputeCases:      "INSERT INTO \"dispute_cases\" %s VALUES %s",
	}
)

type DisputeCasesRepository interface {
	CreateDisputeCases(ctx context.Context, disputeCases *model.DisputeCases, fieldsInsert ...DisputeCasesField) error
	BulkCreateDisputeCases(ctx context.Context, disputeCasesList []*model.DisputeCases, fieldsInsert ...DisputeCasesField) error
	ResolveDisputeCases(ctx context.Context, selectFields ...DisputeCasesField) (model.DisputeCasesList, error)
	ResolveDisputeCasesByID(ctx context.Context, primaryID model.DisputeCasesPrimaryID, selectFields ...DisputeCasesField) (model.DisputeCases, error)
	UpdateDisputeCasesByID(ctx context.Context, id model.DisputeCasesPrimaryID, disputeCases *model.DisputeCases, disputeCasesUpdateFields ...DisputeCasesUpdateField) error
	UpdateDisputeCasesByFilter(ctx context.Context, filter model.Filter, disputeCasesUpdateFields ...DisputeCasesUpdateField) (rowsAffected int64, err error)
	BulkUpdateDisputeCases(ctx context.Context, disputeCasesListMap map[model.DisputeCasesPrimaryID]*model.DisputeCases, DisputeCasessMapUpdateFieldsRequest map[model.DisputeCasesPrimaryID]DisputeCasesUpdateFieldList) (err error)
	DeleteDisputeCasesByID(ctx context.Context, id model.DisputeCasesPrimaryID) error
	BulkDeleteDisputeCasesByIDs(ctx context.Context, ids []model.DisputeCasesPrimaryID) error
	ResolveDisputeCasesByFilter(ctx context.Context, filter model.Filter) (result []model.DisputeCasesFilterResult, err error)
	IsExistDisputeCasesByIDs(ctx context.Context, ids []model.DisputeCasesPrimaryID) (exists bool, notFoundIds []model.DisputeCasesPrimaryID, err error)
	IsExistDisputeCasesByID(ctx context.Context, id model.DisputeCasesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
