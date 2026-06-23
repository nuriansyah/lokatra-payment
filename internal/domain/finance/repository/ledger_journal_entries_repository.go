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

func composeInsertFieldsAndParamsLedgerJournalEntries(ledgerJournalEntriesList []model.LedgerJournalEntries, fieldsInsert ...LedgerJournalEntriesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerJournalEntriesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerJournalEntries := range ledgerJournalEntriesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerJournalEntries.Id)
			case selectField.BatchId():
				args = append(args, ledgerJournalEntries.BatchId)
			case selectField.JournalCode():
				args = append(args, ledgerJournalEntries.JournalCode)
			case selectField.BookId():
				args = append(args, ledgerJournalEntries.BookId)
			case selectField.JournalType():
				args = append(args, ledgerJournalEntries.JournalType)
			case selectField.SourceType():
				args = append(args, ledgerJournalEntries.SourceType)
			case selectField.SourceId():
				args = append(args, ledgerJournalEntries.SourceId)
			case selectField.IdempotencyKey():
				args = append(args, ledgerJournalEntries.IdempotencyKey)
			case selectField.JournalStatus():
				args = append(args, ledgerJournalEntries.JournalStatus)
			case selectField.EffectiveAt():
				args = append(args, ledgerJournalEntries.EffectiveAt)
			case selectField.BookedAt():
				args = append(args, ledgerJournalEntries.BookedAt)
			case selectField.Description():
				args = append(args, ledgerJournalEntries.Description)
			case selectField.ReversalOfJournalId():
				args = append(args, ledgerJournalEntries.ReversalOfJournalId)
			case selectField.Metadata():
				args = append(args, ledgerJournalEntries.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerJournalEntries.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerJournalEntries.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerJournalEntries.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerJournalEntries.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerJournalEntries.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerJournalEntries.MetaDeletedBy)

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

func composeLedgerJournalEntriesCompositePrimaryKeyWhere(primaryIDs []model.LedgerJournalEntriesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_journal_entries\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerJournalEntriesSelectFields() string {
	fields := NewLedgerJournalEntriesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerJournalEntriesSelectFields(selectFields ...LedgerJournalEntriesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerJournalEntriesField string
type LedgerJournalEntriesFieldList []LedgerJournalEntriesField

type LedgerJournalEntriesSelectFields struct {
}

func (ss LedgerJournalEntriesSelectFields) Id() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("id")
}

func (ss LedgerJournalEntriesSelectFields) BatchId() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("batch_id")
}

func (ss LedgerJournalEntriesSelectFields) JournalCode() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("journal_code")
}

func (ss LedgerJournalEntriesSelectFields) BookId() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("book_id")
}

func (ss LedgerJournalEntriesSelectFields) JournalType() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("journal_type")
}

func (ss LedgerJournalEntriesSelectFields) SourceType() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("source_type")
}

func (ss LedgerJournalEntriesSelectFields) SourceId() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("source_id")
}

func (ss LedgerJournalEntriesSelectFields) IdempotencyKey() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("idempotency_key")
}

func (ss LedgerJournalEntriesSelectFields) JournalStatus() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("journal_status")
}

func (ss LedgerJournalEntriesSelectFields) EffectiveAt() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("effective_at")
}

func (ss LedgerJournalEntriesSelectFields) BookedAt() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("booked_at")
}

func (ss LedgerJournalEntriesSelectFields) Description() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("description")
}

func (ss LedgerJournalEntriesSelectFields) ReversalOfJournalId() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("reversal_of_journal_id")
}

func (ss LedgerJournalEntriesSelectFields) Metadata() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("metadata")
}

func (ss LedgerJournalEntriesSelectFields) MetaCreatedAt() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("meta_created_at")
}

func (ss LedgerJournalEntriesSelectFields) MetaCreatedBy() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("meta_created_by")
}

func (ss LedgerJournalEntriesSelectFields) MetaUpdatedAt() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("meta_updated_at")
}

func (ss LedgerJournalEntriesSelectFields) MetaUpdatedBy() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("meta_updated_by")
}

func (ss LedgerJournalEntriesSelectFields) MetaDeletedAt() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("meta_deleted_at")
}

func (ss LedgerJournalEntriesSelectFields) MetaDeletedBy() LedgerJournalEntriesField {
	return LedgerJournalEntriesField("meta_deleted_by")
}

func (ss LedgerJournalEntriesSelectFields) All() LedgerJournalEntriesFieldList {
	return []LedgerJournalEntriesField{
		ss.Id(),
		ss.BatchId(),
		ss.JournalCode(),
		ss.BookId(),
		ss.JournalType(),
		ss.SourceType(),
		ss.SourceId(),
		ss.IdempotencyKey(),
		ss.JournalStatus(),
		ss.EffectiveAt(),
		ss.BookedAt(),
		ss.Description(),
		ss.ReversalOfJournalId(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerJournalEntriesSelectFields() LedgerJournalEntriesSelectFields {
	return LedgerJournalEntriesSelectFields{}
}

type LedgerJournalEntriesUpdateFieldOption struct {
	useIncrement bool
}
type LedgerJournalEntriesUpdateField struct {
	ledgerJournalEntriesField LedgerJournalEntriesField
	opt                       LedgerJournalEntriesUpdateFieldOption
	value                     interface{}
}
type LedgerJournalEntriesUpdateFieldList []LedgerJournalEntriesUpdateField

func defaultLedgerJournalEntriesUpdateFieldOption() LedgerJournalEntriesUpdateFieldOption {
	return LedgerJournalEntriesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerJournalEntriesOption(useIncrement bool) func(*LedgerJournalEntriesUpdateFieldOption) {
	return func(pcufo *LedgerJournalEntriesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerJournalEntriesUpdateField(field LedgerJournalEntriesField, val interface{}, opts ...func(*LedgerJournalEntriesUpdateFieldOption)) LedgerJournalEntriesUpdateField {
	defaultOpt := defaultLedgerJournalEntriesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerJournalEntriesUpdateField{
		ledgerJournalEntriesField: field,
		value:                     val,
		opt:                       defaultOpt,
	}
}
func defaultLedgerJournalEntriesUpdateFields(ledgerJournalEntries model.LedgerJournalEntries) (ledgerJournalEntriesUpdateFieldList LedgerJournalEntriesUpdateFieldList) {
	selectFields := NewLedgerJournalEntriesSelectFields()
	ledgerJournalEntriesUpdateFieldList = append(ledgerJournalEntriesUpdateFieldList,
		NewLedgerJournalEntriesUpdateField(selectFields.Id(), ledgerJournalEntries.Id),
		NewLedgerJournalEntriesUpdateField(selectFields.BatchId(), ledgerJournalEntries.BatchId),
		NewLedgerJournalEntriesUpdateField(selectFields.JournalCode(), ledgerJournalEntries.JournalCode),
		NewLedgerJournalEntriesUpdateField(selectFields.BookId(), ledgerJournalEntries.BookId),
		NewLedgerJournalEntriesUpdateField(selectFields.JournalType(), ledgerJournalEntries.JournalType),
		NewLedgerJournalEntriesUpdateField(selectFields.SourceType(), ledgerJournalEntries.SourceType),
		NewLedgerJournalEntriesUpdateField(selectFields.SourceId(), ledgerJournalEntries.SourceId),
		NewLedgerJournalEntriesUpdateField(selectFields.IdempotencyKey(), ledgerJournalEntries.IdempotencyKey),
		NewLedgerJournalEntriesUpdateField(selectFields.JournalStatus(), ledgerJournalEntries.JournalStatus),
		NewLedgerJournalEntriesUpdateField(selectFields.EffectiveAt(), ledgerJournalEntries.EffectiveAt),
		NewLedgerJournalEntriesUpdateField(selectFields.BookedAt(), ledgerJournalEntries.BookedAt),
		NewLedgerJournalEntriesUpdateField(selectFields.Description(), ledgerJournalEntries.Description),
		NewLedgerJournalEntriesUpdateField(selectFields.ReversalOfJournalId(), ledgerJournalEntries.ReversalOfJournalId),
		NewLedgerJournalEntriesUpdateField(selectFields.Metadata(), ledgerJournalEntries.Metadata),
		NewLedgerJournalEntriesUpdateField(selectFields.MetaCreatedAt(), ledgerJournalEntries.MetaCreatedAt),
		NewLedgerJournalEntriesUpdateField(selectFields.MetaCreatedBy(), ledgerJournalEntries.MetaCreatedBy),
		NewLedgerJournalEntriesUpdateField(selectFields.MetaUpdatedAt(), ledgerJournalEntries.MetaUpdatedAt),
		NewLedgerJournalEntriesUpdateField(selectFields.MetaUpdatedBy(), ledgerJournalEntries.MetaUpdatedBy),
		NewLedgerJournalEntriesUpdateField(selectFields.MetaDeletedAt(), ledgerJournalEntries.MetaDeletedAt),
		NewLedgerJournalEntriesUpdateField(selectFields.MetaDeletedBy(), ledgerJournalEntries.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerJournalEntriesCommand(ledgerJournalEntriesUpdateFieldList LedgerJournalEntriesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerJournalEntriesUpdateFieldList {
		field := string(updateField.ledgerJournalEntriesField)
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

func (repo *RepositoryImpl) BulkCreateLedgerJournalEntries(ctx context.Context, ledgerJournalEntriesList []*model.LedgerJournalEntries, fieldsInsert ...LedgerJournalEntriesField) (err error) {
	var (
		fieldsStr                     string
		valueListStr                  []string
		argsList                      []interface{}
		primaryIds                    []model.LedgerJournalEntriesPrimaryID
		ledgerJournalEntriesValueList []model.LedgerJournalEntries
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerJournalEntriesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerJournalEntries := range ledgerJournalEntriesList {

		primaryIds = append(primaryIds, ledgerJournalEntries.ToLedgerJournalEntriesPrimaryID())

		ledgerJournalEntriesValueList = append(ledgerJournalEntriesValueList, *ledgerJournalEntries)
	}

	_, notFoundIds, err := repo.IsExistLedgerJournalEntriesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerJournalEntries] failed checking ledgerJournalEntries whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerJournalEntriesPrimaryID{}
		mapNotFoundIds := map[model.LedgerJournalEntriesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerJournalEntries", fmt.Sprintf("ledgerJournalEntries with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerJournalEntries(ledgerJournalEntriesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerJournalEntriesQueries.insertLedgerJournalEntries, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerJournalEntries] failed exec create ledgerJournalEntries query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerJournalEntriesByIDs(ctx context.Context, primaryIDs []model.LedgerJournalEntriesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerJournalEntriesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalEntriesByIDs] failed checking ledgerJournalEntries whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalEntries with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_journal_entries\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerJournalEntriesQueries.deleteLedgerJournalEntries + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalEntriesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalEntriesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerJournalEntriesByIDs(ctx context.Context, ids []model.LedgerJournalEntriesPrimaryID) (exists bool, notFoundIds []model.LedgerJournalEntriesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_journal_entries\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerJournalEntriesQueries.selectLedgerJournalEntries, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalEntriesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerJournalEntriesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalEntriesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerJournalEntriesPrimaryID]bool{}
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

// BulkUpdateLedgerJournalEntries is used to bulk update ledgerJournalEntries, by default it will update all field
// if want to update specific field, then fill ledgerJournalEntriessMapUpdateFieldsRequest else please fill ledgerJournalEntriessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerJournalEntries(ctx context.Context, ledgerJournalEntriessMap map[model.LedgerJournalEntriesPrimaryID]*model.LedgerJournalEntries, ledgerJournalEntriessMapUpdateFieldsRequest map[model.LedgerJournalEntriesPrimaryID]LedgerJournalEntriesUpdateFieldList) (err error) {
	if len(ledgerJournalEntriessMap) == 0 && len(ledgerJournalEntriessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerJournalEntriessMapUpdateField map[model.LedgerJournalEntriesPrimaryID]LedgerJournalEntriesUpdateFieldList = map[model.LedgerJournalEntriesPrimaryID]LedgerJournalEntriesUpdateFieldList{}
		asTableValues                       string                                                                      = "myvalues"
	)

	if len(ledgerJournalEntriessMap) > 0 {
		for id, ledgerJournalEntries := range ledgerJournalEntriessMap {
			if ledgerJournalEntries == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerJournalEntries] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerJournalEntriessMapUpdateField[id] = defaultLedgerJournalEntriesUpdateFields(*ledgerJournalEntries)
		}
	} else {
		ledgerJournalEntriessMapUpdateField = ledgerJournalEntriessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerJournalEntriesQuery(ledgerJournalEntriessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerJournalEntriesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerJournalEntries] failed checking ledgerJournalEntries whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalEntries with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerJournalEntriesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_journal_entries\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerJournalEntries] failed exec query")
	}
	return
}

type LedgerJournalEntriesFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerJournalEntriesFieldParameter(param string, args ...interface{}) LedgerJournalEntriesFieldParameter {
	return LedgerJournalEntriesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerJournalEntriesQuery(mapLedgerJournalEntriess map[model.LedgerJournalEntriesPrimaryID]LedgerJournalEntriesUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerJournalEntriesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerJournalEntriesPrimaryID]map[string]interface{}{}
	ledgerJournalEntriesSelectFields := NewLedgerJournalEntriesSelectFields()
	for id, updateFields := range mapLedgerJournalEntriess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerJournalEntriesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerJournalEntriess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerJournalEntriesFieldType(updateField.ledgerJournalEntriesField)))
			args = append(args, fields[string(updateField.ledgerJournalEntriesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerJournalEntriesField))
		if updateField.ledgerJournalEntriesField == ledgerJournalEntriesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerJournalEntriesField, asTableValues, updateField.ledgerJournalEntriesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerJournalEntriesField,
				"\"ledger_journal_entries\"", updateField.ledgerJournalEntriesField,
				asTableValues, updateField.ledgerJournalEntriesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerJournalEntriesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerJournalEntriesPrimaryID, asTableValue string) (whereQry string) {
	ledgerJournalEntriesSelectFields := NewLedgerJournalEntriesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_journal_entries\".\"id\" = %s.\"id\"::"+GetLedgerJournalEntriesFieldType(ledgerJournalEntriesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerJournalEntriesFieldType(ledgerJournalEntriesField LedgerJournalEntriesField) string {
	selectLedgerJournalEntriesFields := NewLedgerJournalEntriesSelectFields()
	switch ledgerJournalEntriesField {

	case selectLedgerJournalEntriesFields.Id():
		return "uuid"

	case selectLedgerJournalEntriesFields.BatchId():
		return "uuid"

	case selectLedgerJournalEntriesFields.JournalCode():
		return "text"

	case selectLedgerJournalEntriesFields.BookId():
		return "uuid"

	case selectLedgerJournalEntriesFields.JournalType():
		return "journal_type_enum"

	case selectLedgerJournalEntriesFields.SourceType():
		return "text"

	case selectLedgerJournalEntriesFields.SourceId():
		return "uuid"

	case selectLedgerJournalEntriesFields.IdempotencyKey():
		return "text"

	case selectLedgerJournalEntriesFields.JournalStatus():
		return "journal_status_enum"

	case selectLedgerJournalEntriesFields.EffectiveAt():
		return "timestamptz"

	case selectLedgerJournalEntriesFields.BookedAt():
		return "timestamptz"

	case selectLedgerJournalEntriesFields.Description():
		return "text"

	case selectLedgerJournalEntriesFields.ReversalOfJournalId():
		return "uuid"

	case selectLedgerJournalEntriesFields.Metadata():
		return "jsonb"

	case selectLedgerJournalEntriesFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerJournalEntriesFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerJournalEntriesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerJournalEntriesFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerJournalEntriesFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerJournalEntriesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerJournalEntries(ctx context.Context, ledgerJournalEntries *model.LedgerJournalEntries, fieldsInsert ...LedgerJournalEntriesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerJournalEntriesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerJournalEntriesPrimaryID{
		Id: ledgerJournalEntries.Id,
	}
	exists, err := repo.IsExistLedgerJournalEntriesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerJournalEntries] failed checking ledgerJournalEntries whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerJournalEntries", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerJournalEntries([]model.LedgerJournalEntries{*ledgerJournalEntries}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerJournalEntriesQueries.insertLedgerJournalEntries, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerJournalEntries] failed exec create ledgerJournalEntries query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerJournalEntriesByID(ctx context.Context, primaryID model.LedgerJournalEntriesPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerJournalEntriesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerJournalEntriesByID] failed checking ledgerJournalEntries whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalEntries with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerJournalEntriesCompositePrimaryKeyWhere([]model.LedgerJournalEntriesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerJournalEntriesQueries.deleteLedgerJournalEntries + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerJournalEntriesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalEntriesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerJournalEntriesFilterResult, err error) {
	query, args, err := composeLedgerJournalEntriesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalEntriesByFilter] failed compose ledgerJournalEntries filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalEntriesByFilter] failed get ledgerJournalEntries by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerJournalEntriesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerJournalEntriesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerJournalEntriesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerJournalEntriesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerJournalEntriesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 20 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerJournalEntriesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["batch_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"batch_id\"")
			selectedColumns["batch_id"] = struct{}{}
		}
		if _, selected := selectedColumns["journal_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"journal_code\"")
			selectedColumns["journal_code"] = struct{}{}
		}
		if _, selected := selectedColumns["book_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"book_id\"")
			selectedColumns["book_id"] = struct{}{}
		}
		if _, selected := selectedColumns["journal_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"journal_type\"")
			selectedColumns["journal_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["journal_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"journal_status\"")
			selectedColumns["journal_status"] = struct{}{}
		}
		if _, selected := selectedColumns["effective_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"effective_at\"")
			selectedColumns["effective_at"] = struct{}{}
		}
		if _, selected := selectedColumns["booked_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"booked_at\"")
			selectedColumns["booked_at"] = struct{}{}
		}
		if _, selected := selectedColumns["description"]; !selected {
			selectColumns = append(selectColumns, "base.\"description\"")
			selectedColumns["description"] = struct{}{}
		}
		if _, selected := selectedColumns["reversal_of_journal_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"reversal_of_journal_id\"")
			selectedColumns["reversal_of_journal_id"] = struct{}{}
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

type ledgerJournalEntriesFilterPlaceholder struct {
	index int
}

func (p *ledgerJournalEntriesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerJournalEntriesFilterPredicate(filterField model.FilterField, placeholders *ledgerJournalEntriesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerJournalEntriesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerJournalEntriesFilterSQLExpr(spec)
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

func composeLedgerJournalEntriesFilterGroup(group model.FilterGroup, placeholders *ledgerJournalEntriesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerJournalEntriesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerJournalEntriesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerJournalEntriesFilterWhereQueries(filter model.Filter, placeholders *ledgerJournalEntriesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerJournalEntriesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerJournalEntriesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerJournalEntriesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerJournalEntriesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerJournalEntriesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerJournalEntriesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerJournalEntriesFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerJournalEntriesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerJournalEntriesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerJournalEntriesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerJournalEntriesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_journal_entries\" base%s", strings.Join(selectColumns, ","), composeLedgerJournalEntriesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerJournalEntriesByID(ctx context.Context, primaryID model.LedgerJournalEntriesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerJournalEntriesCompositePrimaryKeyWhere([]model.LedgerJournalEntriesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerJournalEntriesQueries.selectCountLedgerJournalEntries, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalEntriesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalEntries(ctx context.Context, selectFields ...LedgerJournalEntriesField) (ledgerJournalEntriesList model.LedgerJournalEntriesList, err error) {
	var (
		defaultLedgerJournalEntriesSelectFields = defaultLedgerJournalEntriesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerJournalEntriesSelectFields = composeLedgerJournalEntriesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerJournalEntriesQueries.selectLedgerJournalEntries, defaultLedgerJournalEntriesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerJournalEntriesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalEntries] failed get ledgerJournalEntries list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalEntriesByID(ctx context.Context, primaryID model.LedgerJournalEntriesPrimaryID, selectFields ...LedgerJournalEntriesField) (ledgerJournalEntries model.LedgerJournalEntries, err error) {
	var (
		defaultLedgerJournalEntriesSelectFields = defaultLedgerJournalEntriesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerJournalEntriesSelectFields = composeLedgerJournalEntriesSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerJournalEntriesCompositePrimaryKeyWhere([]model.LedgerJournalEntriesPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerJournalEntriesQueries.selectLedgerJournalEntries+" WHERE "+whereQry, defaultLedgerJournalEntriesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerJournalEntries, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerJournalEntries with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerJournalEntriesByID] failed get ledgerJournalEntries")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerJournalEntriesByID(ctx context.Context, primaryID model.LedgerJournalEntriesPrimaryID, ledgerJournalEntries *model.LedgerJournalEntries, ledgerJournalEntriesUpdateFields ...LedgerJournalEntriesUpdateField) (err error) {
	exists, err := repo.IsExistLedgerJournalEntriesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalEntries] failed checking ledgerJournalEntries whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalEntries with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerJournalEntries == nil {
		if len(ledgerJournalEntriesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerJournalEntriesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerJournalEntries = &model.LedgerJournalEntries{}
	}
	var (
		defaultLedgerJournalEntriesUpdateFields = defaultLedgerJournalEntriesUpdateFields(*ledgerJournalEntries)
		tempUpdateField                         LedgerJournalEntriesUpdateFieldList
		selectFields                            = NewLedgerJournalEntriesSelectFields()
	)
	if len(ledgerJournalEntriesUpdateFields) > 0 {
		for _, updateField := range ledgerJournalEntriesUpdateFields {
			if updateField.ledgerJournalEntriesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerJournalEntriesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerJournalEntriesCompositePrimaryKeyWhere([]model.LedgerJournalEntriesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerJournalEntriesCommand(defaultLedgerJournalEntriesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerJournalEntriesQueries.updateLedgerJournalEntries+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalEntries] error when try to update ledgerJournalEntries by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerJournalEntriesByFilter(ctx context.Context, filter model.Filter, ledgerJournalEntriesUpdateFields ...LedgerJournalEntriesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerJournalEntriesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerJournalEntriesUpdateFieldList
		selectFields = NewLedgerJournalEntriesSelectFields()
	)
	for _, updateField := range ledgerJournalEntriesUpdateFields {
		if updateField.ledgerJournalEntriesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerJournalEntriesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerJournalEntriesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerJournalEntriesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_journal_entries\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalEntriesByFilter] error when try to update ledgerJournalEntries by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalEntriesByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerJournalEntriesQueries = struct {
		selectLedgerJournalEntries      string
		selectCountLedgerJournalEntries string
		deleteLedgerJournalEntries      string
		updateLedgerJournalEntries      string
		insertLedgerJournalEntries      string
	}{
		selectLedgerJournalEntries:      "SELECT %s FROM \"ledger_journal_entries\"",
		selectCountLedgerJournalEntries: "SELECT COUNT(\"id\") FROM \"ledger_journal_entries\"",
		deleteLedgerJournalEntries:      "DELETE FROM \"ledger_journal_entries\"",
		updateLedgerJournalEntries:      "UPDATE \"ledger_journal_entries\" SET %s ",
		insertLedgerJournalEntries:      "INSERT INTO \"ledger_journal_entries\" %s VALUES %s",
	}
)

type LedgerJournalEntriesRepository interface {
	CreateLedgerJournalEntries(ctx context.Context, ledgerJournalEntries *model.LedgerJournalEntries, fieldsInsert ...LedgerJournalEntriesField) error
	BulkCreateLedgerJournalEntries(ctx context.Context, ledgerJournalEntriesList []*model.LedgerJournalEntries, fieldsInsert ...LedgerJournalEntriesField) error
	ResolveLedgerJournalEntries(ctx context.Context, selectFields ...LedgerJournalEntriesField) (model.LedgerJournalEntriesList, error)
	ResolveLedgerJournalEntriesByID(ctx context.Context, primaryID model.LedgerJournalEntriesPrimaryID, selectFields ...LedgerJournalEntriesField) (model.LedgerJournalEntries, error)
	UpdateLedgerJournalEntriesByID(ctx context.Context, id model.LedgerJournalEntriesPrimaryID, ledgerJournalEntries *model.LedgerJournalEntries, ledgerJournalEntriesUpdateFields ...LedgerJournalEntriesUpdateField) error
	UpdateLedgerJournalEntriesByFilter(ctx context.Context, filter model.Filter, ledgerJournalEntriesUpdateFields ...LedgerJournalEntriesUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerJournalEntries(ctx context.Context, ledgerJournalEntriesListMap map[model.LedgerJournalEntriesPrimaryID]*model.LedgerJournalEntries, LedgerJournalEntriessMapUpdateFieldsRequest map[model.LedgerJournalEntriesPrimaryID]LedgerJournalEntriesUpdateFieldList) (err error)
	DeleteLedgerJournalEntriesByID(ctx context.Context, id model.LedgerJournalEntriesPrimaryID) error
	BulkDeleteLedgerJournalEntriesByIDs(ctx context.Context, ids []model.LedgerJournalEntriesPrimaryID) error
	ResolveLedgerJournalEntriesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerJournalEntriesFilterResult, err error)
	IsExistLedgerJournalEntriesByIDs(ctx context.Context, ids []model.LedgerJournalEntriesPrimaryID) (exists bool, notFoundIds []model.LedgerJournalEntriesPrimaryID, err error)
	IsExistLedgerJournalEntriesByID(ctx context.Context, id model.LedgerJournalEntriesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
