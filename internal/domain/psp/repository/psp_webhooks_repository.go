package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/psp/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsPspWebhooks(pspWebhooksList []model.PspWebhooks, fieldsInsert ...PspWebhooksField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewPspWebhooksSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, pspWebhooks := range pspWebhooksList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, pspWebhooks.Id)
			case selectField.PspAccountId():
				args = append(args, pspWebhooks.PspAccountId)
			case selectField.Psp():
				args = append(args, pspWebhooks.Psp)
			case selectField.PspEventId():
				args = append(args, pspWebhooks.PspEventId)
			case selectField.PspEventType():
				args = append(args, pspWebhooks.PspEventType)
			case selectField.ReceivedAt():
				args = append(args, pspWebhooks.ReceivedAt)
			case selectField.Headers():
				args = append(args, pspWebhooks.Headers)
			case selectField.RawPayload():
				args = append(args, pspWebhooks.RawPayload)
			case selectField.HmacValid():
				args = append(args, pspWebhooks.HmacValid)
			case selectField.Status():
				args = append(args, pspWebhooks.Status)
			case selectField.ProcessingAttempts():
				args = append(args, pspWebhooks.ProcessingAttempts)
			case selectField.LastError():
				args = append(args, pspWebhooks.LastError)
			case selectField.ProcessedAt():
				args = append(args, pspWebhooks.ProcessedAt)
			case selectField.ResolvedPaymentId():
				args = append(args, pspWebhooks.ResolvedPaymentId)
			case selectField.ResolvedIntentId():
				args = append(args, pspWebhooks.ResolvedIntentId)
			case selectField.MetaCreatedAt():
				args = append(args, pspWebhooks.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, pspWebhooks.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, pspWebhooks.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, pspWebhooks.MetaUpdatedBy)

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

func composePspWebhooksCompositePrimaryKeyWhere(primaryIDs []model.PspWebhooksPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"psp_webhooks\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultPspWebhooksSelectFields() string {
	fields := NewPspWebhooksSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composePspWebhooksSelectFields(selectFields ...PspWebhooksField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type PspWebhooksField string
type PspWebhooksFieldList []PspWebhooksField

type PspWebhooksSelectFields struct {
}

func (ss PspWebhooksSelectFields) Id() PspWebhooksField {
	return PspWebhooksField("id")
}

func (ss PspWebhooksSelectFields) PspAccountId() PspWebhooksField {
	return PspWebhooksField("psp_account_id")
}

func (ss PspWebhooksSelectFields) Psp() PspWebhooksField {
	return PspWebhooksField("psp")
}

func (ss PspWebhooksSelectFields) PspEventId() PspWebhooksField {
	return PspWebhooksField("psp_event_id")
}

func (ss PspWebhooksSelectFields) PspEventType() PspWebhooksField {
	return PspWebhooksField("psp_event_type")
}

func (ss PspWebhooksSelectFields) ReceivedAt() PspWebhooksField {
	return PspWebhooksField("received_at")
}

func (ss PspWebhooksSelectFields) Headers() PspWebhooksField {
	return PspWebhooksField("headers")
}

func (ss PspWebhooksSelectFields) RawPayload() PspWebhooksField {
	return PspWebhooksField("raw_payload")
}

func (ss PspWebhooksSelectFields) HmacValid() PspWebhooksField {
	return PspWebhooksField("hmac_valid")
}

func (ss PspWebhooksSelectFields) Status() PspWebhooksField {
	return PspWebhooksField("status")
}

func (ss PspWebhooksSelectFields) ProcessingAttempts() PspWebhooksField {
	return PspWebhooksField("processing_attempts")
}

func (ss PspWebhooksSelectFields) LastError() PspWebhooksField {
	return PspWebhooksField("last_error")
}

func (ss PspWebhooksSelectFields) ProcessedAt() PspWebhooksField {
	return PspWebhooksField("processed_at")
}

func (ss PspWebhooksSelectFields) ResolvedPaymentId() PspWebhooksField {
	return PspWebhooksField("resolved_payment_id")
}

func (ss PspWebhooksSelectFields) ResolvedIntentId() PspWebhooksField {
	return PspWebhooksField("resolved_intent_id")
}

func (ss PspWebhooksSelectFields) MetaCreatedAt() PspWebhooksField {
	return PspWebhooksField("meta_created_at")
}

func (ss PspWebhooksSelectFields) MetaCreatedBy() PspWebhooksField {
	return PspWebhooksField("meta_created_by")
}

func (ss PspWebhooksSelectFields) MetaUpdatedAt() PspWebhooksField {
	return PspWebhooksField("meta_updated_at")
}

func (ss PspWebhooksSelectFields) MetaUpdatedBy() PspWebhooksField {
	return PspWebhooksField("meta_updated_by")
}

func (ss PspWebhooksSelectFields) All() PspWebhooksFieldList {
	return []PspWebhooksField{
		ss.Id(),
		ss.PspAccountId(),
		ss.Psp(),
		ss.PspEventId(),
		ss.PspEventType(),
		ss.ReceivedAt(),
		ss.Headers(),
		ss.RawPayload(),
		ss.HmacValid(),
		ss.Status(),
		ss.ProcessingAttempts(),
		ss.LastError(),
		ss.ProcessedAt(),
		ss.ResolvedPaymentId(),
		ss.ResolvedIntentId(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
	}
}

func NewPspWebhooksSelectFields() PspWebhooksSelectFields {
	return PspWebhooksSelectFields{}
}

type PspWebhooksUpdateFieldOption struct {
	useIncrement bool
}
type PspWebhooksUpdateField struct {
	pspWebhooksField PspWebhooksField
	opt              PspWebhooksUpdateFieldOption
	value            interface{}
}
type PspWebhooksUpdateFieldList []PspWebhooksUpdateField

func defaultPspWebhooksUpdateFieldOption() PspWebhooksUpdateFieldOption {
	return PspWebhooksUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementPspWebhooksOption(useIncrement bool) func(*PspWebhooksUpdateFieldOption) {
	return func(pcufo *PspWebhooksUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewPspWebhooksUpdateField(field PspWebhooksField, val interface{}, opts ...func(*PspWebhooksUpdateFieldOption)) PspWebhooksUpdateField {
	defaultOpt := defaultPspWebhooksUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return PspWebhooksUpdateField{
		pspWebhooksField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultPspWebhooksUpdateFields(pspWebhooks model.PspWebhooks) (pspWebhooksUpdateFieldList PspWebhooksUpdateFieldList) {
	selectFields := NewPspWebhooksSelectFields()
	pspWebhooksUpdateFieldList = append(pspWebhooksUpdateFieldList,
		NewPspWebhooksUpdateField(selectFields.Id(), pspWebhooks.Id),
		NewPspWebhooksUpdateField(selectFields.PspAccountId(), pspWebhooks.PspAccountId),
		NewPspWebhooksUpdateField(selectFields.Psp(), pspWebhooks.Psp),
		NewPspWebhooksUpdateField(selectFields.PspEventId(), pspWebhooks.PspEventId),
		NewPspWebhooksUpdateField(selectFields.PspEventType(), pspWebhooks.PspEventType),
		NewPspWebhooksUpdateField(selectFields.ReceivedAt(), pspWebhooks.ReceivedAt),
		NewPspWebhooksUpdateField(selectFields.Headers(), pspWebhooks.Headers),
		NewPspWebhooksUpdateField(selectFields.RawPayload(), pspWebhooks.RawPayload),
		NewPspWebhooksUpdateField(selectFields.HmacValid(), pspWebhooks.HmacValid),
		NewPspWebhooksUpdateField(selectFields.Status(), pspWebhooks.Status),
		NewPspWebhooksUpdateField(selectFields.ProcessingAttempts(), pspWebhooks.ProcessingAttempts),
		NewPspWebhooksUpdateField(selectFields.LastError(), pspWebhooks.LastError),
		NewPspWebhooksUpdateField(selectFields.ProcessedAt(), pspWebhooks.ProcessedAt),
		NewPspWebhooksUpdateField(selectFields.ResolvedPaymentId(), pspWebhooks.ResolvedPaymentId),
		NewPspWebhooksUpdateField(selectFields.ResolvedIntentId(), pspWebhooks.ResolvedIntentId),
		NewPspWebhooksUpdateField(selectFields.MetaCreatedAt(), pspWebhooks.MetaCreatedAt),
		NewPspWebhooksUpdateField(selectFields.MetaCreatedBy(), pspWebhooks.MetaCreatedBy),
		NewPspWebhooksUpdateField(selectFields.MetaUpdatedAt(), pspWebhooks.MetaUpdatedAt),
		NewPspWebhooksUpdateField(selectFields.MetaUpdatedBy(), pspWebhooks.MetaUpdatedBy),
	)
	return
}
func composeUpdateFieldsPspWebhooksCommand(pspWebhooksUpdateFieldList PspWebhooksUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range pspWebhooksUpdateFieldList {
		field := string(updateField.pspWebhooksField)
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

func (repo *RepositoryImpl) BulkCreatePspWebhooks(ctx context.Context, pspWebhooksList []*model.PspWebhooks, fieldsInsert ...PspWebhooksField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.PspWebhooksPrimaryID
		pspWebhooksValueList []model.PspWebhooks
	)

	if len(fieldsInsert) == 0 {
		selectField := NewPspWebhooksSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, pspWebhooks := range pspWebhooksList {

		primaryIds = append(primaryIds, pspWebhooks.ToPspWebhooksPrimaryID())

		pspWebhooksValueList = append(pspWebhooksValueList, *pspWebhooks)
	}

	_, notFoundIds, err := repo.IsExistPspWebhooksByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePspWebhooks] failed checking pspWebhooks whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.PspWebhooksPrimaryID{}
		mapNotFoundIds := map[model.PspWebhooksPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "pspWebhooks", fmt.Sprintf("pspWebhooks with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsPspWebhooks(pspWebhooksValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(pspWebhooksQueries.insertPspWebhooks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreatePspWebhooks] failed exec create pspWebhooks query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeletePspWebhooksByIDs(ctx context.Context, primaryIDs []model.PspWebhooksPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistPspWebhooksByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePspWebhooksByIDs] failed checking pspWebhooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("pspWebhooks with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"psp_webhooks\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(pspWebhooksQueries.deletePspWebhooks + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePspWebhooksByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeletePspWebhooksByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistPspWebhooksByIDs(ctx context.Context, ids []model.PspWebhooksPrimaryID) (exists bool, notFoundIds []model.PspWebhooksPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"psp_webhooks\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(pspWebhooksQueries.selectPspWebhooks, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPspWebhooksByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.PspWebhooksPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPspWebhooksByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.PspWebhooksPrimaryID]bool{}
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

// BulkUpdatePspWebhooks is used to bulk update pspWebhooks, by default it will update all field
// if want to update specific field, then fill pspWebhookssMapUpdateFieldsRequest else please fill pspWebhookssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdatePspWebhooks(ctx context.Context, pspWebhookssMap map[model.PspWebhooksPrimaryID]*model.PspWebhooks, pspWebhookssMapUpdateFieldsRequest map[model.PspWebhooksPrimaryID]PspWebhooksUpdateFieldList) (err error) {
	if len(pspWebhookssMap) == 0 && len(pspWebhookssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		pspWebhookssMapUpdateField map[model.PspWebhooksPrimaryID]PspWebhooksUpdateFieldList = map[model.PspWebhooksPrimaryID]PspWebhooksUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(pspWebhookssMap) > 0 {
		for id, pspWebhooks := range pspWebhookssMap {
			if pspWebhooks == nil {
				log.Error().Err(err).Msg("[BulkUpdatePspWebhooks] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			pspWebhookssMapUpdateField[id] = defaultPspWebhooksUpdateFields(*pspWebhooks)
		}
	} else {
		pspWebhookssMapUpdateField = pspWebhookssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdatePspWebhooksQuery(pspWebhookssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistPspWebhooksByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePspWebhooks] failed checking pspWebhooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("pspWebhooks with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composePspWebhooksCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"psp_webhooks\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdatePspWebhooks] failed exec query")
	}
	return
}

type PspWebhooksFieldParameter struct {
	param string
	args  []interface{}
}

func NewPspWebhooksFieldParameter(param string, args ...interface{}) PspWebhooksFieldParameter {
	return PspWebhooksFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdatePspWebhooksQuery(mapPspWebhookss map[model.PspWebhooksPrimaryID]PspWebhooksUpdateFieldList, asTableValues string) (primaryIDs []model.PspWebhooksPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.PspWebhooksPrimaryID]map[string]interface{}{}
	pspWebhooksSelectFields := NewPspWebhooksSelectFields()
	for id, updateFields := range mapPspWebhookss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.pspWebhooksField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapPspWebhookss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetPspWebhooksFieldType(updateField.pspWebhooksField)))
			args = append(args, fields[string(updateField.pspWebhooksField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.pspWebhooksField))
		if updateField.pspWebhooksField == pspWebhooksSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.pspWebhooksField, asTableValues, updateField.pspWebhooksField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.pspWebhooksField,
				"\"psp_webhooks\"", updateField.pspWebhooksField,
				asTableValues, updateField.pspWebhooksField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composePspWebhooksCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.PspWebhooksPrimaryID, asTableValue string) (whereQry string) {
	pspWebhooksSelectFields := NewPspWebhooksSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"psp_webhooks\".\"id\" = %s.\"id\"::"+GetPspWebhooksFieldType(pspWebhooksSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetPspWebhooksFieldType(pspWebhooksField PspWebhooksField) string {
	selectPspWebhooksFields := NewPspWebhooksSelectFields()
	switch pspWebhooksField {

	case selectPspWebhooksFields.Id():
		return "uuid"

	case selectPspWebhooksFields.PspAccountId():
		return "uuid"

	case selectPspWebhooksFields.Psp():
		return "text"

	case selectPspWebhooksFields.PspEventId():
		return "text"

	case selectPspWebhooksFields.PspEventType():
		return "text"

	case selectPspWebhooksFields.ReceivedAt():
		return "timestamptz"

	case selectPspWebhooksFields.Headers():
		return "jsonb"

	case selectPspWebhooksFields.RawPayload():
		return "jsonb"

	case selectPspWebhooksFields.HmacValid():
		return "bool"

	case selectPspWebhooksFields.Status():
		return "webhook_status_enum"

	case selectPspWebhooksFields.ProcessingAttempts():
		return "int4"

	case selectPspWebhooksFields.LastError():
		return "text"

	case selectPspWebhooksFields.ProcessedAt():
		return "timestamptz"

	case selectPspWebhooksFields.ResolvedPaymentId():
		return "uuid"

	case selectPspWebhooksFields.ResolvedIntentId():
		return "uuid"

	case selectPspWebhooksFields.MetaCreatedAt():
		return "timestamptz"

	case selectPspWebhooksFields.MetaCreatedBy():
		return "uuid"

	case selectPspWebhooksFields.MetaUpdatedAt():
		return "timestamptz"

	case selectPspWebhooksFields.MetaUpdatedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreatePspWebhooks(ctx context.Context, pspWebhooks *model.PspWebhooks, fieldsInsert ...PspWebhooksField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewPspWebhooksSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.PspWebhooksPrimaryID{
		Id: pspWebhooks.Id,
	}
	exists, err := repo.IsExistPspWebhooksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePspWebhooks] failed checking pspWebhooks whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "pspWebhooks", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsPspWebhooks([]model.PspWebhooks{*pspWebhooks}, fieldsInsert...)
	commandQuery := fmt.Sprintf(pspWebhooksQueries.insertPspWebhooks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreatePspWebhooks] failed exec create pspWebhooks query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeletePspWebhooksByID(ctx context.Context, primaryID model.PspWebhooksPrimaryID) (err error) {
	exists, err := repo.IsExistPspWebhooksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePspWebhooksByID] failed checking pspWebhooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("pspWebhooks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composePspWebhooksCompositePrimaryKeyWhere([]model.PspWebhooksPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(pspWebhooksQueries.deletePspWebhooks + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeletePspWebhooksByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolvePspWebhooksByFilter(ctx context.Context, filter model.Filter) (result []model.PspWebhooksFilterResult, err error) {
	query, args, err := composePspWebhooksFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePspWebhooksByFilter] failed compose pspWebhooks filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePspWebhooksByFilter] failed get pspWebhooks by filter")
		err = failure.InternalError(err)
	}
	return
}

func composePspWebhooksFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidatePspWebhooksFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultPspWebhooksSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := PspWebhooksFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, PspWebhooksField(filterSelectField))
		}
		selectFields = composePspWebhooksSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(pspWebhooksQueries.selectPspWebhooks, selectFields)

	if len(filter.FilterFields) > 0 {
		var (
			whereQueries []string
			whereArgs    []interface{}
		)
		for _, filterField := range filter.FilterFields {
			switch filterField.Operator {
			case model.OperatorEqual:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" = $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			case model.OperatorRange:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" BETWEEN $%d AND $%d", filterField.Field, index, index+1))
				valueArray, ok := filterField.Value.([]interface{})
				if !ok && len(valueArray) != 2 {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				whereArgs = append(whereArgs, valueArray...)
				index += 2
			case model.OperatorIn:
				valueArray, ok := filterField.Value.([]interface{})
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				var placeholder []string
				for range valueArray {
					placeholder = append(placeholder, fmt.Sprintf("$%d", index))
					index++
				}
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IN (%s)", filterField.Field, strings.Join(placeholder, ",")))
				whereArgs = append(whereArgs, valueArray...)
			case model.OperatorIsNull:
				value, ok := filterField.Value.(bool)
				if !ok {
					err = failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					return
				}
				if value {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NULL", filterField.Field))
				} else {
					whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" IS NOT NULL", filterField.Field))
				}
			case model.OperatorNot:
				whereQueries = append(whereQueries, fmt.Sprintf("\"%s\" != $%d", filterField.Field, index))
				whereArgs = append(whereArgs, filterField.Value)
				index++
			}
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(whereQueries, " AND "))
		args = append(args, whereArgs...)
	}

	sortQuery := []string{}
	for _, sort := range filter.Sorts {
		sortQuery = append(sortQuery, fmt.Sprintf("\"%s\" %s", sort.Field, sort.Order))
	}
	if len(sortQuery) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sortQuery, ","))
	}
	if filter.Pagination.PageSize > 0 {
		query += fmt.Sprintf(" LIMIT %d", filter.Pagination.PageSize)
		if filter.Pagination.Page > 0 {
			query += fmt.Sprintf(" OFFSET %d", (filter.Pagination.Page-1)*filter.Pagination.PageSize)
		}
	}

	return
}

func (repo *RepositoryImpl) IsExistPspWebhooksByID(ctx context.Context, primaryID model.PspWebhooksPrimaryID) (exists bool, err error) {
	whereQuery, params := composePspWebhooksCompositePrimaryKeyWhere([]model.PspWebhooksPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", pspWebhooksQueries.selectCountPspWebhooks, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistPspWebhooksByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePspWebhooks(ctx context.Context, selectFields ...PspWebhooksField) (pspWebhooksList model.PspWebhooksList, err error) {
	var (
		defaultPspWebhooksSelectFields = defaultPspWebhooksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPspWebhooksSelectFields = composePspWebhooksSelectFields(selectFields...)
	}
	query := fmt.Sprintf(pspWebhooksQueries.selectPspWebhooks, defaultPspWebhooksSelectFields)

	err = repo.db.Read.Select(&pspWebhooksList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolvePspWebhooks] failed get pspWebhooks list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolvePspWebhooksByID(ctx context.Context, primaryID model.PspWebhooksPrimaryID, selectFields ...PspWebhooksField) (pspWebhooks model.PspWebhooks, err error) {
	var (
		defaultPspWebhooksSelectFields = defaultPspWebhooksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultPspWebhooksSelectFields = composePspWebhooksSelectFields(selectFields...)
	}
	whereQry, params := composePspWebhooksCompositePrimaryKeyWhere([]model.PspWebhooksPrimaryID{primaryID})
	query := fmt.Sprintf(pspWebhooksQueries.selectPspWebhooks+" WHERE "+whereQry, defaultPspWebhooksSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&pspWebhooks, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("pspWebhooks with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolvePspWebhooksByID] failed get pspWebhooks")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdatePspWebhooksByID(ctx context.Context, primaryID model.PspWebhooksPrimaryID, pspWebhooks *model.PspWebhooks, pspWebhooksUpdateFields ...PspWebhooksUpdateField) (err error) {
	exists, err := repo.IsExistPspWebhooksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePspWebhooks] failed checking pspWebhooks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("pspWebhooks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if pspWebhooks == nil {
		if len(pspWebhooksUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdatePspWebhooksByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		pspWebhooks = &model.PspWebhooks{}
	}
	var (
		defaultPspWebhooksUpdateFields = defaultPspWebhooksUpdateFields(*pspWebhooks)
		tempUpdateField                PspWebhooksUpdateFieldList
		selectFields                   = NewPspWebhooksSelectFields()
	)
	if len(pspWebhooksUpdateFields) > 0 {
		for _, updateField := range pspWebhooksUpdateFields {
			if updateField.pspWebhooksField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultPspWebhooksUpdateFields = tempUpdateField
	}
	whereQuery, params := composePspWebhooksCompositePrimaryKeyWhere([]model.PspWebhooksPrimaryID{primaryID})
	fields, args := composeUpdateFieldsPspWebhooksCommand(defaultPspWebhooksUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(pspWebhooksQueries.updatePspWebhooks+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdatePspWebhooks] error when try to update pspWebhooks by id")
	}
	return err
}

var (
	pspWebhooksQueries = struct {
		selectPspWebhooks      string
		selectCountPspWebhooks string
		deletePspWebhooks      string
		updatePspWebhooks      string
		insertPspWebhooks      string
	}{
		selectPspWebhooks:      "SELECT %s FROM \"psp_webhooks\"",
		selectCountPspWebhooks: "SELECT COUNT(\"id\") FROM \"psp_webhooks\"",
		deletePspWebhooks:      "DELETE FROM \"psp_webhooks\"",
		updatePspWebhooks:      "UPDATE \"psp_webhooks\" SET %s ",
		insertPspWebhooks:      "INSERT INTO \"psp_webhooks\" %s VALUES %s",
	}
)

type PspWebhooksRepository interface {
	CreatePspWebhooks(ctx context.Context, pspWebhooks *model.PspWebhooks, fieldsInsert ...PspWebhooksField) error
	BulkCreatePspWebhooks(ctx context.Context, pspWebhooksList []*model.PspWebhooks, fieldsInsert ...PspWebhooksField) error
	ResolvePspWebhooks(ctx context.Context, selectFields ...PspWebhooksField) (model.PspWebhooksList, error)
	ResolvePspWebhooksByID(ctx context.Context, primaryID model.PspWebhooksPrimaryID, selectFields ...PspWebhooksField) (model.PspWebhooks, error)
	UpdatePspWebhooksByID(ctx context.Context, id model.PspWebhooksPrimaryID, pspWebhooks *model.PspWebhooks, pspWebhooksUpdateFields ...PspWebhooksUpdateField) error
	BulkUpdatePspWebhooks(ctx context.Context, pspWebhooksListMap map[model.PspWebhooksPrimaryID]*model.PspWebhooks, PspWebhookssMapUpdateFieldsRequest map[model.PspWebhooksPrimaryID]PspWebhooksUpdateFieldList) (err error)
	DeletePspWebhooksByID(ctx context.Context, id model.PspWebhooksPrimaryID) error
	BulkDeletePspWebhooksByIDs(ctx context.Context, ids []model.PspWebhooksPrimaryID) error
	ResolvePspWebhooksByFilter(ctx context.Context, filter model.Filter) (result []model.PspWebhooksFilterResult, err error)
	IsExistPspWebhooksByIDs(ctx context.Context, ids []model.PspWebhooksPrimaryID) (exists bool, notFoundIds []model.PspWebhooksPrimaryID, err error)
	IsExistPspWebhooksByID(ctx context.Context, id model.PspWebhooksPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
