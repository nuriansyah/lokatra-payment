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

func composeInsertFieldsAndParamsLedgerJournalBatches(ledgerJournalBatchesList []model.LedgerJournalBatches, fieldsInsert ...LedgerJournalBatchesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerJournalBatchesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerJournalBatches := range ledgerJournalBatchesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerJournalBatches.Id)
			case selectField.BatchCode():
				args = append(args, ledgerJournalBatches.BatchCode)
			case selectField.BatchType():
				args = append(args, ledgerJournalBatches.BatchType)
			case selectField.SourceRef():
				args = append(args, ledgerJournalBatches.SourceRef)
			case selectField.BatchStatus():
				args = append(args, ledgerJournalBatches.BatchStatus)
			case selectField.BookedAt():
				args = append(args, ledgerJournalBatches.BookedAt)
			case selectField.Description():
				args = append(args, ledgerJournalBatches.Description)
			case selectField.Metadata():
				args = append(args, ledgerJournalBatches.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerJournalBatches.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerJournalBatches.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerJournalBatches.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerJournalBatches.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerJournalBatches.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerJournalBatches.MetaDeletedBy)

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

func composeLedgerJournalBatchesCompositePrimaryKeyWhere(primaryIDs []model.LedgerJournalBatchesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_journal_batches\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerJournalBatchesSelectFields() string {
	fields := NewLedgerJournalBatchesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerJournalBatchesSelectFields(selectFields ...LedgerJournalBatchesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerJournalBatchesField string
type LedgerJournalBatchesFieldList []LedgerJournalBatchesField

type LedgerJournalBatchesSelectFields struct {
}

func (ss LedgerJournalBatchesSelectFields) Id() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("id")
}

func (ss LedgerJournalBatchesSelectFields) BatchCode() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("batch_code")
}

func (ss LedgerJournalBatchesSelectFields) BatchType() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("batch_type")
}

func (ss LedgerJournalBatchesSelectFields) SourceRef() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("source_ref")
}

func (ss LedgerJournalBatchesSelectFields) BatchStatus() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("batch_status")
}

func (ss LedgerJournalBatchesSelectFields) BookedAt() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("booked_at")
}

func (ss LedgerJournalBatchesSelectFields) Description() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("description")
}

func (ss LedgerJournalBatchesSelectFields) Metadata() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("metadata")
}

func (ss LedgerJournalBatchesSelectFields) MetaCreatedAt() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("meta_created_at")
}

func (ss LedgerJournalBatchesSelectFields) MetaCreatedBy() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("meta_created_by")
}

func (ss LedgerJournalBatchesSelectFields) MetaUpdatedAt() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("meta_updated_at")
}

func (ss LedgerJournalBatchesSelectFields) MetaUpdatedBy() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("meta_updated_by")
}

func (ss LedgerJournalBatchesSelectFields) MetaDeletedAt() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("meta_deleted_at")
}

func (ss LedgerJournalBatchesSelectFields) MetaDeletedBy() LedgerJournalBatchesField {
	return LedgerJournalBatchesField("meta_deleted_by")
}

func (ss LedgerJournalBatchesSelectFields) All() LedgerJournalBatchesFieldList {
	return []LedgerJournalBatchesField{
		ss.Id(),
		ss.BatchCode(),
		ss.BatchType(),
		ss.SourceRef(),
		ss.BatchStatus(),
		ss.BookedAt(),
		ss.Description(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerJournalBatchesSelectFields() LedgerJournalBatchesSelectFields {
	return LedgerJournalBatchesSelectFields{}
}

type LedgerJournalBatchesUpdateFieldOption struct {
	useIncrement bool
}
type LedgerJournalBatchesUpdateField struct {
	ledgerJournalBatchesField LedgerJournalBatchesField
	opt                       LedgerJournalBatchesUpdateFieldOption
	value                     interface{}
}
type LedgerJournalBatchesUpdateFieldList []LedgerJournalBatchesUpdateField

func defaultLedgerJournalBatchesUpdateFieldOption() LedgerJournalBatchesUpdateFieldOption {
	return LedgerJournalBatchesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerJournalBatchesOption(useIncrement bool) func(*LedgerJournalBatchesUpdateFieldOption) {
	return func(pcufo *LedgerJournalBatchesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerJournalBatchesUpdateField(field LedgerJournalBatchesField, val interface{}, opts ...func(*LedgerJournalBatchesUpdateFieldOption)) LedgerJournalBatchesUpdateField {
	defaultOpt := defaultLedgerJournalBatchesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerJournalBatchesUpdateField{
		ledgerJournalBatchesField: field,
		value:                     val,
		opt:                       defaultOpt,
	}
}
func defaultLedgerJournalBatchesUpdateFields(ledgerJournalBatches model.LedgerJournalBatches) (ledgerJournalBatchesUpdateFieldList LedgerJournalBatchesUpdateFieldList) {
	selectFields := NewLedgerJournalBatchesSelectFields()
	ledgerJournalBatchesUpdateFieldList = append(ledgerJournalBatchesUpdateFieldList,
		NewLedgerJournalBatchesUpdateField(selectFields.Id(), ledgerJournalBatches.Id),
		NewLedgerJournalBatchesUpdateField(selectFields.BatchCode(), ledgerJournalBatches.BatchCode),
		NewLedgerJournalBatchesUpdateField(selectFields.BatchType(), ledgerJournalBatches.BatchType),
		NewLedgerJournalBatchesUpdateField(selectFields.SourceRef(), ledgerJournalBatches.SourceRef),
		NewLedgerJournalBatchesUpdateField(selectFields.BatchStatus(), ledgerJournalBatches.BatchStatus),
		NewLedgerJournalBatchesUpdateField(selectFields.BookedAt(), ledgerJournalBatches.BookedAt),
		NewLedgerJournalBatchesUpdateField(selectFields.Description(), ledgerJournalBatches.Description),
		NewLedgerJournalBatchesUpdateField(selectFields.Metadata(), ledgerJournalBatches.Metadata),
		NewLedgerJournalBatchesUpdateField(selectFields.MetaCreatedAt(), ledgerJournalBatches.MetaCreatedAt),
		NewLedgerJournalBatchesUpdateField(selectFields.MetaCreatedBy(), ledgerJournalBatches.MetaCreatedBy),
		NewLedgerJournalBatchesUpdateField(selectFields.MetaUpdatedAt(), ledgerJournalBatches.MetaUpdatedAt),
		NewLedgerJournalBatchesUpdateField(selectFields.MetaUpdatedBy(), ledgerJournalBatches.MetaUpdatedBy),
		NewLedgerJournalBatchesUpdateField(selectFields.MetaDeletedAt(), ledgerJournalBatches.MetaDeletedAt),
		NewLedgerJournalBatchesUpdateField(selectFields.MetaDeletedBy(), ledgerJournalBatches.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerJournalBatchesCommand(ledgerJournalBatchesUpdateFieldList LedgerJournalBatchesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerJournalBatchesUpdateFieldList {
		field := string(updateField.ledgerJournalBatchesField)
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

func (repo *RepositoryImpl) BulkCreateLedgerJournalBatches(ctx context.Context, ledgerJournalBatchesList []*model.LedgerJournalBatches, fieldsInsert ...LedgerJournalBatchesField) (err error) {
	var (
		fieldsStr                     string
		valueListStr                  []string
		argsList                      []interface{}
		primaryIds                    []model.LedgerJournalBatchesPrimaryID
		ledgerJournalBatchesValueList []model.LedgerJournalBatches
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerJournalBatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerJournalBatches := range ledgerJournalBatchesList {

		primaryIds = append(primaryIds, ledgerJournalBatches.ToLedgerJournalBatchesPrimaryID())

		ledgerJournalBatchesValueList = append(ledgerJournalBatchesValueList, *ledgerJournalBatches)
	}

	_, notFoundIds, err := repo.IsExistLedgerJournalBatchesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerJournalBatches] failed checking ledgerJournalBatches whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerJournalBatchesPrimaryID{}
		mapNotFoundIds := map[model.LedgerJournalBatchesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerJournalBatches", fmt.Sprintf("ledgerJournalBatches with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerJournalBatches(ledgerJournalBatchesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerJournalBatchesQueries.insertLedgerJournalBatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerJournalBatches] failed exec create ledgerJournalBatches query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerJournalBatchesByIDs(ctx context.Context, primaryIDs []model.LedgerJournalBatchesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerJournalBatchesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalBatchesByIDs] failed checking ledgerJournalBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalBatches with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_journal_batches\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerJournalBatchesQueries.deleteLedgerJournalBatches + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalBatchesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalBatchesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerJournalBatchesByIDs(ctx context.Context, ids []model.LedgerJournalBatchesPrimaryID) (exists bool, notFoundIds []model.LedgerJournalBatchesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_journal_batches\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerJournalBatchesQueries.selectLedgerJournalBatches, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalBatchesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerJournalBatchesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalBatchesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerJournalBatchesPrimaryID]bool{}
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

// BulkUpdateLedgerJournalBatches is used to bulk update ledgerJournalBatches, by default it will update all field
// if want to update specific field, then fill ledgerJournalBatchessMapUpdateFieldsRequest else please fill ledgerJournalBatchessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerJournalBatches(ctx context.Context, ledgerJournalBatchessMap map[model.LedgerJournalBatchesPrimaryID]*model.LedgerJournalBatches, ledgerJournalBatchessMapUpdateFieldsRequest map[model.LedgerJournalBatchesPrimaryID]LedgerJournalBatchesUpdateFieldList) (err error) {
	if len(ledgerJournalBatchessMap) == 0 && len(ledgerJournalBatchessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerJournalBatchessMapUpdateField map[model.LedgerJournalBatchesPrimaryID]LedgerJournalBatchesUpdateFieldList = map[model.LedgerJournalBatchesPrimaryID]LedgerJournalBatchesUpdateFieldList{}
		asTableValues                       string                                                                      = "myvalues"
	)

	if len(ledgerJournalBatchessMap) > 0 {
		for id, ledgerJournalBatches := range ledgerJournalBatchessMap {
			if ledgerJournalBatches == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerJournalBatches] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerJournalBatchessMapUpdateField[id] = defaultLedgerJournalBatchesUpdateFields(*ledgerJournalBatches)
		}
	} else {
		ledgerJournalBatchessMapUpdateField = ledgerJournalBatchessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerJournalBatchesQuery(ledgerJournalBatchessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerJournalBatchesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerJournalBatches] failed checking ledgerJournalBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalBatches with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerJournalBatchesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_journal_batches\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerJournalBatches] failed exec query")
	}
	return
}

type LedgerJournalBatchesFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerJournalBatchesFieldParameter(param string, args ...interface{}) LedgerJournalBatchesFieldParameter {
	return LedgerJournalBatchesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerJournalBatchesQuery(mapLedgerJournalBatchess map[model.LedgerJournalBatchesPrimaryID]LedgerJournalBatchesUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerJournalBatchesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerJournalBatchesPrimaryID]map[string]interface{}{}
	ledgerJournalBatchesSelectFields := NewLedgerJournalBatchesSelectFields()
	for id, updateFields := range mapLedgerJournalBatchess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerJournalBatchesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerJournalBatchess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerJournalBatchesFieldType(updateField.ledgerJournalBatchesField)))
			args = append(args, fields[string(updateField.ledgerJournalBatchesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerJournalBatchesField))
		if updateField.ledgerJournalBatchesField == ledgerJournalBatchesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerJournalBatchesField, asTableValues, updateField.ledgerJournalBatchesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerJournalBatchesField,
				"\"ledger_journal_batches\"", updateField.ledgerJournalBatchesField,
				asTableValues, updateField.ledgerJournalBatchesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerJournalBatchesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerJournalBatchesPrimaryID, asTableValue string) (whereQry string) {
	ledgerJournalBatchesSelectFields := NewLedgerJournalBatchesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_journal_batches\".\"id\" = %s.\"id\"::"+GetLedgerJournalBatchesFieldType(ledgerJournalBatchesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerJournalBatchesFieldType(ledgerJournalBatchesField LedgerJournalBatchesField) string {
	selectLedgerJournalBatchesFields := NewLedgerJournalBatchesSelectFields()
	switch ledgerJournalBatchesField {

	case selectLedgerJournalBatchesFields.Id():
		return "uuid"

	case selectLedgerJournalBatchesFields.BatchCode():
		return "text"

	case selectLedgerJournalBatchesFields.BatchType():
		return "batch_type_enum"

	case selectLedgerJournalBatchesFields.SourceRef():
		return "text"

	case selectLedgerJournalBatchesFields.BatchStatus():
		return "ledger_journal_batches_batch_status_enum"

	case selectLedgerJournalBatchesFields.BookedAt():
		return "timestamptz"

	case selectLedgerJournalBatchesFields.Description():
		return "text"

	case selectLedgerJournalBatchesFields.Metadata():
		return "jsonb"

	case selectLedgerJournalBatchesFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerJournalBatchesFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerJournalBatchesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerJournalBatchesFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerJournalBatchesFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerJournalBatchesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerJournalBatches(ctx context.Context, ledgerJournalBatches *model.LedgerJournalBatches, fieldsInsert ...LedgerJournalBatchesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerJournalBatchesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerJournalBatchesPrimaryID{
		Id: ledgerJournalBatches.Id,
	}
	exists, err := repo.IsExistLedgerJournalBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerJournalBatches] failed checking ledgerJournalBatches whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerJournalBatches", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerJournalBatches([]model.LedgerJournalBatches{*ledgerJournalBatches}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerJournalBatchesQueries.insertLedgerJournalBatches, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerJournalBatches] failed exec create ledgerJournalBatches query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerJournalBatchesByID(ctx context.Context, primaryID model.LedgerJournalBatchesPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerJournalBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerJournalBatchesByID] failed checking ledgerJournalBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalBatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerJournalBatchesCompositePrimaryKeyWhere([]model.LedgerJournalBatchesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerJournalBatchesQueries.deleteLedgerJournalBatches + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerJournalBatchesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalBatchesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerJournalBatchesFilterResult, err error) {
	query, args, err := composeLedgerJournalBatchesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalBatchesByFilter] failed compose ledgerJournalBatches filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalBatchesByFilter] failed get ledgerJournalBatches by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerJournalBatchesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerJournalBatchesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerJournalBatchesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerJournalBatchesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerJournalBatchesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 14 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerJournalBatchesFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 14+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["batch_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"batch_code\"")
			selectedColumns["batch_code"] = struct{}{}
		}
		if _, selected := selectedColumns["batch_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"batch_type\"")
			selectedColumns["batch_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_ref"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_ref\"")
			selectedColumns["source_ref"] = struct{}{}
		}
		if _, selected := selectedColumns["batch_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"batch_status\"")
			selectedColumns["batch_status"] = struct{}{}
		}
		if _, selected := selectedColumns["booked_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"booked_at\"")
			selectedColumns["booked_at"] = struct{}{}
		}
		if _, selected := selectedColumns["description"]; !selected {
			selectColumns = append(selectColumns, "base.\"description\"")
			selectedColumns["description"] = struct{}{}
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

type ledgerJournalBatchesFilterPlaceholder struct {
	index int
}

func (p *ledgerJournalBatchesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerJournalBatchesFilterPredicate(filterField model.FilterField, placeholders *ledgerJournalBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerJournalBatchesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerJournalBatchesFilterSQLExpr(spec)
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

func composeLedgerJournalBatchesFilterGroup(group model.FilterGroup, placeholders *ledgerJournalBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerJournalBatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerJournalBatchesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerJournalBatchesFilterWhereQueries(filter model.Filter, placeholders *ledgerJournalBatchesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerJournalBatchesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerJournalBatchesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerJournalBatchesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerJournalBatchesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerJournalBatchesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerJournalBatchesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerJournalBatchesFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerJournalBatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerJournalBatchesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerJournalBatchesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerJournalBatchesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_journal_batches\" base%s", strings.Join(selectColumns, ","), composeLedgerJournalBatchesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerJournalBatchesByID(ctx context.Context, primaryID model.LedgerJournalBatchesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerJournalBatchesCompositePrimaryKeyWhere([]model.LedgerJournalBatchesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerJournalBatchesQueries.selectCountLedgerJournalBatches, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalBatchesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalBatches(ctx context.Context, selectFields ...LedgerJournalBatchesField) (ledgerJournalBatchesList model.LedgerJournalBatchesList, err error) {
	var (
		defaultLedgerJournalBatchesSelectFields = defaultLedgerJournalBatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerJournalBatchesSelectFields = composeLedgerJournalBatchesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerJournalBatchesQueries.selectLedgerJournalBatches, defaultLedgerJournalBatchesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerJournalBatchesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalBatches] failed get ledgerJournalBatches list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalBatchesByID(ctx context.Context, primaryID model.LedgerJournalBatchesPrimaryID, selectFields ...LedgerJournalBatchesField) (ledgerJournalBatches model.LedgerJournalBatches, err error) {
	var (
		defaultLedgerJournalBatchesSelectFields = defaultLedgerJournalBatchesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerJournalBatchesSelectFields = composeLedgerJournalBatchesSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerJournalBatchesCompositePrimaryKeyWhere([]model.LedgerJournalBatchesPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerJournalBatchesQueries.selectLedgerJournalBatches+" WHERE "+whereQry, defaultLedgerJournalBatchesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerJournalBatches, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerJournalBatches with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerJournalBatchesByID] failed get ledgerJournalBatches")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerJournalBatchesByID(ctx context.Context, primaryID model.LedgerJournalBatchesPrimaryID, ledgerJournalBatches *model.LedgerJournalBatches, ledgerJournalBatchesUpdateFields ...LedgerJournalBatchesUpdateField) (err error) {
	exists, err := repo.IsExistLedgerJournalBatchesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalBatches] failed checking ledgerJournalBatches whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalBatches with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerJournalBatches == nil {
		if len(ledgerJournalBatchesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerJournalBatchesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerJournalBatches = &model.LedgerJournalBatches{}
	}
	var (
		defaultLedgerJournalBatchesUpdateFields = defaultLedgerJournalBatchesUpdateFields(*ledgerJournalBatches)
		tempUpdateField                         LedgerJournalBatchesUpdateFieldList
		selectFields                            = NewLedgerJournalBatchesSelectFields()
	)
	if len(ledgerJournalBatchesUpdateFields) > 0 {
		for _, updateField := range ledgerJournalBatchesUpdateFields {
			if updateField.ledgerJournalBatchesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerJournalBatchesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerJournalBatchesCompositePrimaryKeyWhere([]model.LedgerJournalBatchesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerJournalBatchesCommand(defaultLedgerJournalBatchesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerJournalBatchesQueries.updateLedgerJournalBatches+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalBatches] error when try to update ledgerJournalBatches by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerJournalBatchesByFilter(ctx context.Context, filter model.Filter, ledgerJournalBatchesUpdateFields ...LedgerJournalBatchesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerJournalBatchesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerJournalBatchesUpdateFieldList
		selectFields = NewLedgerJournalBatchesSelectFields()
	)
	for _, updateField := range ledgerJournalBatchesUpdateFields {
		if updateField.ledgerJournalBatchesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerJournalBatchesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerJournalBatchesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerJournalBatchesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_journal_batches\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalBatchesByFilter] error when try to update ledgerJournalBatches by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalBatchesByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerJournalBatchesQueries = struct {
		selectLedgerJournalBatches      string
		selectCountLedgerJournalBatches string
		deleteLedgerJournalBatches      string
		updateLedgerJournalBatches      string
		insertLedgerJournalBatches      string
	}{
		selectLedgerJournalBatches:      "SELECT %s FROM \"ledger_journal_batches\"",
		selectCountLedgerJournalBatches: "SELECT COUNT(\"id\") FROM \"ledger_journal_batches\"",
		deleteLedgerJournalBatches:      "DELETE FROM \"ledger_journal_batches\"",
		updateLedgerJournalBatches:      "UPDATE \"ledger_journal_batches\" SET %s ",
		insertLedgerJournalBatches:      "INSERT INTO \"ledger_journal_batches\" %s VALUES %s",
	}
)

type LedgerJournalBatchesRepository interface {
	CreateLedgerJournalBatches(ctx context.Context, ledgerJournalBatches *model.LedgerJournalBatches, fieldsInsert ...LedgerJournalBatchesField) error
	BulkCreateLedgerJournalBatches(ctx context.Context, ledgerJournalBatchesList []*model.LedgerJournalBatches, fieldsInsert ...LedgerJournalBatchesField) error
	ResolveLedgerJournalBatches(ctx context.Context, selectFields ...LedgerJournalBatchesField) (model.LedgerJournalBatchesList, error)
	ResolveLedgerJournalBatchesByID(ctx context.Context, primaryID model.LedgerJournalBatchesPrimaryID, selectFields ...LedgerJournalBatchesField) (model.LedgerJournalBatches, error)
	UpdateLedgerJournalBatchesByID(ctx context.Context, id model.LedgerJournalBatchesPrimaryID, ledgerJournalBatches *model.LedgerJournalBatches, ledgerJournalBatchesUpdateFields ...LedgerJournalBatchesUpdateField) error
	UpdateLedgerJournalBatchesByFilter(ctx context.Context, filter model.Filter, ledgerJournalBatchesUpdateFields ...LedgerJournalBatchesUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerJournalBatches(ctx context.Context, ledgerJournalBatchesListMap map[model.LedgerJournalBatchesPrimaryID]*model.LedgerJournalBatches, LedgerJournalBatchessMapUpdateFieldsRequest map[model.LedgerJournalBatchesPrimaryID]LedgerJournalBatchesUpdateFieldList) (err error)
	DeleteLedgerJournalBatchesByID(ctx context.Context, id model.LedgerJournalBatchesPrimaryID) error
	BulkDeleteLedgerJournalBatchesByIDs(ctx context.Context, ids []model.LedgerJournalBatchesPrimaryID) error
	ResolveLedgerJournalBatchesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerJournalBatchesFilterResult, err error)
	IsExistLedgerJournalBatchesByIDs(ctx context.Context, ids []model.LedgerJournalBatchesPrimaryID) (exists bool, notFoundIds []model.LedgerJournalBatchesPrimaryID, err error)
	IsExistLedgerJournalBatchesByID(ctx context.Context, id model.LedgerJournalBatchesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
