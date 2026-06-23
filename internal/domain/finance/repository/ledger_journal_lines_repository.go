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

func composeInsertFieldsAndParamsLedgerJournalLines(ledgerJournalLinesList []model.LedgerJournalLines, fieldsInsert ...LedgerJournalLinesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerJournalLinesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerJournalLines := range ledgerJournalLinesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerJournalLines.Id)
			case selectField.JournalEntryId():
				args = append(args, ledgerJournalLines.JournalEntryId)
			case selectField.LineNo():
				args = append(args, ledgerJournalLines.LineNo)
			case selectField.AccountId():
				args = append(args, ledgerJournalLines.AccountId)
			case selectField.LineSide():
				args = append(args, ledgerJournalLines.LineSide)
			case selectField.Amount():
				args = append(args, ledgerJournalLines.Amount)
			case selectField.CurrencyCode():
				args = append(args, ledgerJournalLines.CurrencyCode)
			case selectField.FxRateLockId():
				args = append(args, ledgerJournalLines.FxRateLockId)
			case selectField.AmountReporting():
				args = append(args, ledgerJournalLines.AmountReporting)
			case selectField.ReferencePartyId():
				args = append(args, ledgerJournalLines.ReferencePartyId)
			case selectField.Dimensions():
				args = append(args, ledgerJournalLines.Dimensions)
			case selectField.Metadata():
				args = append(args, ledgerJournalLines.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerJournalLines.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerJournalLines.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerJournalLines.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerJournalLines.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerJournalLines.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerJournalLines.MetaDeletedBy)

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

func composeLedgerJournalLinesCompositePrimaryKeyWhere(primaryIDs []model.LedgerJournalLinesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_journal_lines\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerJournalLinesSelectFields() string {
	fields := NewLedgerJournalLinesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerJournalLinesSelectFields(selectFields ...LedgerJournalLinesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerJournalLinesField string
type LedgerJournalLinesFieldList []LedgerJournalLinesField

type LedgerJournalLinesSelectFields struct {
}

func (ss LedgerJournalLinesSelectFields) Id() LedgerJournalLinesField {
	return LedgerJournalLinesField("id")
}

func (ss LedgerJournalLinesSelectFields) JournalEntryId() LedgerJournalLinesField {
	return LedgerJournalLinesField("journal_entry_id")
}

func (ss LedgerJournalLinesSelectFields) LineNo() LedgerJournalLinesField {
	return LedgerJournalLinesField("line_no")
}

func (ss LedgerJournalLinesSelectFields) AccountId() LedgerJournalLinesField {
	return LedgerJournalLinesField("account_id")
}

func (ss LedgerJournalLinesSelectFields) LineSide() LedgerJournalLinesField {
	return LedgerJournalLinesField("line_side")
}

func (ss LedgerJournalLinesSelectFields) Amount() LedgerJournalLinesField {
	return LedgerJournalLinesField("amount")
}

func (ss LedgerJournalLinesSelectFields) CurrencyCode() LedgerJournalLinesField {
	return LedgerJournalLinesField("currency_code")
}

func (ss LedgerJournalLinesSelectFields) FxRateLockId() LedgerJournalLinesField {
	return LedgerJournalLinesField("fx_rate_lock_id")
}

func (ss LedgerJournalLinesSelectFields) AmountReporting() LedgerJournalLinesField {
	return LedgerJournalLinesField("amount_reporting")
}

func (ss LedgerJournalLinesSelectFields) ReferencePartyId() LedgerJournalLinesField {
	return LedgerJournalLinesField("reference_party_id")
}

func (ss LedgerJournalLinesSelectFields) Dimensions() LedgerJournalLinesField {
	return LedgerJournalLinesField("dimensions")
}

func (ss LedgerJournalLinesSelectFields) Metadata() LedgerJournalLinesField {
	return LedgerJournalLinesField("metadata")
}

func (ss LedgerJournalLinesSelectFields) MetaCreatedAt() LedgerJournalLinesField {
	return LedgerJournalLinesField("meta_created_at")
}

func (ss LedgerJournalLinesSelectFields) MetaCreatedBy() LedgerJournalLinesField {
	return LedgerJournalLinesField("meta_created_by")
}

func (ss LedgerJournalLinesSelectFields) MetaUpdatedAt() LedgerJournalLinesField {
	return LedgerJournalLinesField("meta_updated_at")
}

func (ss LedgerJournalLinesSelectFields) MetaUpdatedBy() LedgerJournalLinesField {
	return LedgerJournalLinesField("meta_updated_by")
}

func (ss LedgerJournalLinesSelectFields) MetaDeletedAt() LedgerJournalLinesField {
	return LedgerJournalLinesField("meta_deleted_at")
}

func (ss LedgerJournalLinesSelectFields) MetaDeletedBy() LedgerJournalLinesField {
	return LedgerJournalLinesField("meta_deleted_by")
}

func (ss LedgerJournalLinesSelectFields) All() LedgerJournalLinesFieldList {
	return []LedgerJournalLinesField{
		ss.Id(),
		ss.JournalEntryId(),
		ss.LineNo(),
		ss.AccountId(),
		ss.LineSide(),
		ss.Amount(),
		ss.CurrencyCode(),
		ss.FxRateLockId(),
		ss.AmountReporting(),
		ss.ReferencePartyId(),
		ss.Dimensions(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerJournalLinesSelectFields() LedgerJournalLinesSelectFields {
	return LedgerJournalLinesSelectFields{}
}

type LedgerJournalLinesUpdateFieldOption struct {
	useIncrement bool
}
type LedgerJournalLinesUpdateField struct {
	ledgerJournalLinesField LedgerJournalLinesField
	opt                     LedgerJournalLinesUpdateFieldOption
	value                   interface{}
}
type LedgerJournalLinesUpdateFieldList []LedgerJournalLinesUpdateField

func defaultLedgerJournalLinesUpdateFieldOption() LedgerJournalLinesUpdateFieldOption {
	return LedgerJournalLinesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerJournalLinesOption(useIncrement bool) func(*LedgerJournalLinesUpdateFieldOption) {
	return func(pcufo *LedgerJournalLinesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerJournalLinesUpdateField(field LedgerJournalLinesField, val interface{}, opts ...func(*LedgerJournalLinesUpdateFieldOption)) LedgerJournalLinesUpdateField {
	defaultOpt := defaultLedgerJournalLinesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerJournalLinesUpdateField{
		ledgerJournalLinesField: field,
		value:                   val,
		opt:                     defaultOpt,
	}
}
func defaultLedgerJournalLinesUpdateFields(ledgerJournalLines model.LedgerJournalLines) (ledgerJournalLinesUpdateFieldList LedgerJournalLinesUpdateFieldList) {
	selectFields := NewLedgerJournalLinesSelectFields()
	ledgerJournalLinesUpdateFieldList = append(ledgerJournalLinesUpdateFieldList,
		NewLedgerJournalLinesUpdateField(selectFields.Id(), ledgerJournalLines.Id),
		NewLedgerJournalLinesUpdateField(selectFields.JournalEntryId(), ledgerJournalLines.JournalEntryId),
		NewLedgerJournalLinesUpdateField(selectFields.LineNo(), ledgerJournalLines.LineNo),
		NewLedgerJournalLinesUpdateField(selectFields.AccountId(), ledgerJournalLines.AccountId),
		NewLedgerJournalLinesUpdateField(selectFields.LineSide(), ledgerJournalLines.LineSide),
		NewLedgerJournalLinesUpdateField(selectFields.Amount(), ledgerJournalLines.Amount),
		NewLedgerJournalLinesUpdateField(selectFields.CurrencyCode(), ledgerJournalLines.CurrencyCode),
		NewLedgerJournalLinesUpdateField(selectFields.FxRateLockId(), ledgerJournalLines.FxRateLockId),
		NewLedgerJournalLinesUpdateField(selectFields.AmountReporting(), ledgerJournalLines.AmountReporting),
		NewLedgerJournalLinesUpdateField(selectFields.ReferencePartyId(), ledgerJournalLines.ReferencePartyId),
		NewLedgerJournalLinesUpdateField(selectFields.Dimensions(), ledgerJournalLines.Dimensions),
		NewLedgerJournalLinesUpdateField(selectFields.Metadata(), ledgerJournalLines.Metadata),
		NewLedgerJournalLinesUpdateField(selectFields.MetaCreatedAt(), ledgerJournalLines.MetaCreatedAt),
		NewLedgerJournalLinesUpdateField(selectFields.MetaCreatedBy(), ledgerJournalLines.MetaCreatedBy),
		NewLedgerJournalLinesUpdateField(selectFields.MetaUpdatedAt(), ledgerJournalLines.MetaUpdatedAt),
		NewLedgerJournalLinesUpdateField(selectFields.MetaUpdatedBy(), ledgerJournalLines.MetaUpdatedBy),
		NewLedgerJournalLinesUpdateField(selectFields.MetaDeletedAt(), ledgerJournalLines.MetaDeletedAt),
		NewLedgerJournalLinesUpdateField(selectFields.MetaDeletedBy(), ledgerJournalLines.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerJournalLinesCommand(ledgerJournalLinesUpdateFieldList LedgerJournalLinesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerJournalLinesUpdateFieldList {
		field := string(updateField.ledgerJournalLinesField)
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

func (repo *RepositoryImpl) BulkCreateLedgerJournalLines(ctx context.Context, ledgerJournalLinesList []*model.LedgerJournalLines, fieldsInsert ...LedgerJournalLinesField) (err error) {
	var (
		fieldsStr                   string
		valueListStr                []string
		argsList                    []interface{}
		primaryIds                  []model.LedgerJournalLinesPrimaryID
		ledgerJournalLinesValueList []model.LedgerJournalLines
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerJournalLinesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerJournalLines := range ledgerJournalLinesList {

		primaryIds = append(primaryIds, ledgerJournalLines.ToLedgerJournalLinesPrimaryID())

		ledgerJournalLinesValueList = append(ledgerJournalLinesValueList, *ledgerJournalLines)
	}

	_, notFoundIds, err := repo.IsExistLedgerJournalLinesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerJournalLines] failed checking ledgerJournalLines whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerJournalLinesPrimaryID{}
		mapNotFoundIds := map[model.LedgerJournalLinesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerJournalLines", fmt.Sprintf("ledgerJournalLines with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerJournalLines(ledgerJournalLinesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerJournalLinesQueries.insertLedgerJournalLines, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerJournalLines] failed exec create ledgerJournalLines query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerJournalLinesByIDs(ctx context.Context, primaryIDs []model.LedgerJournalLinesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerJournalLinesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalLinesByIDs] failed checking ledgerJournalLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalLines with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_journal_lines\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerJournalLinesQueries.deleteLedgerJournalLines + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalLinesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerJournalLinesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerJournalLinesByIDs(ctx context.Context, ids []model.LedgerJournalLinesPrimaryID) (exists bool, notFoundIds []model.LedgerJournalLinesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_journal_lines\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerJournalLinesQueries.selectLedgerJournalLines, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalLinesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerJournalLinesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalLinesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerJournalLinesPrimaryID]bool{}
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

// BulkUpdateLedgerJournalLines is used to bulk update ledgerJournalLines, by default it will update all field
// if want to update specific field, then fill ledgerJournalLinessMapUpdateFieldsRequest else please fill ledgerJournalLinessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerJournalLines(ctx context.Context, ledgerJournalLinessMap map[model.LedgerJournalLinesPrimaryID]*model.LedgerJournalLines, ledgerJournalLinessMapUpdateFieldsRequest map[model.LedgerJournalLinesPrimaryID]LedgerJournalLinesUpdateFieldList) (err error) {
	if len(ledgerJournalLinessMap) == 0 && len(ledgerJournalLinessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerJournalLinessMapUpdateField map[model.LedgerJournalLinesPrimaryID]LedgerJournalLinesUpdateFieldList = map[model.LedgerJournalLinesPrimaryID]LedgerJournalLinesUpdateFieldList{}
		asTableValues                     string                                                                  = "myvalues"
	)

	if len(ledgerJournalLinessMap) > 0 {
		for id, ledgerJournalLines := range ledgerJournalLinessMap {
			if ledgerJournalLines == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerJournalLines] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerJournalLinessMapUpdateField[id] = defaultLedgerJournalLinesUpdateFields(*ledgerJournalLines)
		}
	} else {
		ledgerJournalLinessMapUpdateField = ledgerJournalLinessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerJournalLinesQuery(ledgerJournalLinessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerJournalLinesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerJournalLines] failed checking ledgerJournalLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalLines with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerJournalLinesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_journal_lines\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerJournalLines] failed exec query")
	}
	return
}

type LedgerJournalLinesFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerJournalLinesFieldParameter(param string, args ...interface{}) LedgerJournalLinesFieldParameter {
	return LedgerJournalLinesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerJournalLinesQuery(mapLedgerJournalLiness map[model.LedgerJournalLinesPrimaryID]LedgerJournalLinesUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerJournalLinesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerJournalLinesPrimaryID]map[string]interface{}{}
	ledgerJournalLinesSelectFields := NewLedgerJournalLinesSelectFields()
	for id, updateFields := range mapLedgerJournalLiness {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerJournalLinesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerJournalLiness[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerJournalLinesFieldType(updateField.ledgerJournalLinesField)))
			args = append(args, fields[string(updateField.ledgerJournalLinesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerJournalLinesField))
		if updateField.ledgerJournalLinesField == ledgerJournalLinesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerJournalLinesField, asTableValues, updateField.ledgerJournalLinesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerJournalLinesField,
				"\"ledger_journal_lines\"", updateField.ledgerJournalLinesField,
				asTableValues, updateField.ledgerJournalLinesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerJournalLinesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerJournalLinesPrimaryID, asTableValue string) (whereQry string) {
	ledgerJournalLinesSelectFields := NewLedgerJournalLinesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_journal_lines\".\"id\" = %s.\"id\"::"+GetLedgerJournalLinesFieldType(ledgerJournalLinesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerJournalLinesFieldType(ledgerJournalLinesField LedgerJournalLinesField) string {
	selectLedgerJournalLinesFields := NewLedgerJournalLinesSelectFields()
	switch ledgerJournalLinesField {

	case selectLedgerJournalLinesFields.Id():
		return "uuid"

	case selectLedgerJournalLinesFields.JournalEntryId():
		return "uuid"

	case selectLedgerJournalLinesFields.LineNo():
		return "int4"

	case selectLedgerJournalLinesFields.AccountId():
		return "uuid"

	case selectLedgerJournalLinesFields.LineSide():
		return "line_side_enum"

	case selectLedgerJournalLinesFields.Amount():
		return "numeric"

	case selectLedgerJournalLinesFields.CurrencyCode():
		return "text"

	case selectLedgerJournalLinesFields.FxRateLockId():
		return "uuid"

	case selectLedgerJournalLinesFields.AmountReporting():
		return "numeric"

	case selectLedgerJournalLinesFields.ReferencePartyId():
		return "uuid"

	case selectLedgerJournalLinesFields.Dimensions():
		return "jsonb"

	case selectLedgerJournalLinesFields.Metadata():
		return "jsonb"

	case selectLedgerJournalLinesFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerJournalLinesFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerJournalLinesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerJournalLinesFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerJournalLinesFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerJournalLinesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerJournalLines(ctx context.Context, ledgerJournalLines *model.LedgerJournalLines, fieldsInsert ...LedgerJournalLinesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerJournalLinesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerJournalLinesPrimaryID{
		Id: ledgerJournalLines.Id,
	}
	exists, err := repo.IsExistLedgerJournalLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerJournalLines] failed checking ledgerJournalLines whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerJournalLines", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerJournalLines([]model.LedgerJournalLines{*ledgerJournalLines}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerJournalLinesQueries.insertLedgerJournalLines, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerJournalLines] failed exec create ledgerJournalLines query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerJournalLinesByID(ctx context.Context, primaryID model.LedgerJournalLinesPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerJournalLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerJournalLinesByID] failed checking ledgerJournalLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalLines with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerJournalLinesCompositePrimaryKeyWhere([]model.LedgerJournalLinesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerJournalLinesQueries.deleteLedgerJournalLines + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerJournalLinesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalLinesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerJournalLinesFilterResult, err error) {
	query, args, err := composeLedgerJournalLinesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalLinesByFilter] failed compose ledgerJournalLines filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalLinesByFilter] failed get ledgerJournalLines by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerJournalLinesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerJournalLinesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerJournalLinesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerJournalLinesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerJournalLinesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 18 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerJournalLinesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["journal_entry_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"journal_entry_id\"")
			selectedColumns["journal_entry_id"] = struct{}{}
		}
		if _, selected := selectedColumns["line_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"line_no\"")
			selectedColumns["line_no"] = struct{}{}
		}
		if _, selected := selectedColumns["account_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"account_id\"")
			selectedColumns["account_id"] = struct{}{}
		}
		if _, selected := selectedColumns["line_side"]; !selected {
			selectColumns = append(selectColumns, "base.\"line_side\"")
			selectedColumns["line_side"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["fx_rate_lock_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"fx_rate_lock_id\"")
			selectedColumns["fx_rate_lock_id"] = struct{}{}
		}
		if _, selected := selectedColumns["amount_reporting"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount_reporting\"")
			selectedColumns["amount_reporting"] = struct{}{}
		}
		if _, selected := selectedColumns["reference_party_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"reference_party_id\"")
			selectedColumns["reference_party_id"] = struct{}{}
		}
		if _, selected := selectedColumns["dimensions"]; !selected {
			selectColumns = append(selectColumns, "base.\"dimensions\"")
			selectedColumns["dimensions"] = struct{}{}
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

type ledgerJournalLinesFilterPlaceholder struct {
	index int
}

func (p *ledgerJournalLinesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerJournalLinesFilterPredicate(filterField model.FilterField, placeholders *ledgerJournalLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerJournalLinesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerJournalLinesFilterSQLExpr(spec)
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

func composeLedgerJournalLinesFilterGroup(group model.FilterGroup, placeholders *ledgerJournalLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerJournalLinesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerJournalLinesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerJournalLinesFilterWhereQueries(filter model.Filter, placeholders *ledgerJournalLinesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerJournalLinesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerJournalLinesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerJournalLinesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerJournalLinesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerJournalLinesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerJournalLinesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerJournalLinesFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerJournalLinesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerJournalLinesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerJournalLinesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerJournalLinesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_journal_lines\" base%s", strings.Join(selectColumns, ","), composeLedgerJournalLinesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerJournalLinesByID(ctx context.Context, primaryID model.LedgerJournalLinesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerJournalLinesCompositePrimaryKeyWhere([]model.LedgerJournalLinesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerJournalLinesQueries.selectCountLedgerJournalLines, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerJournalLinesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalLines(ctx context.Context, selectFields ...LedgerJournalLinesField) (ledgerJournalLinesList model.LedgerJournalLinesList, err error) {
	var (
		defaultLedgerJournalLinesSelectFields = defaultLedgerJournalLinesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerJournalLinesSelectFields = composeLedgerJournalLinesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerJournalLinesQueries.selectLedgerJournalLines, defaultLedgerJournalLinesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerJournalLinesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerJournalLines] failed get ledgerJournalLines list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerJournalLinesByID(ctx context.Context, primaryID model.LedgerJournalLinesPrimaryID, selectFields ...LedgerJournalLinesField) (ledgerJournalLines model.LedgerJournalLines, err error) {
	var (
		defaultLedgerJournalLinesSelectFields = defaultLedgerJournalLinesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerJournalLinesSelectFields = composeLedgerJournalLinesSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerJournalLinesCompositePrimaryKeyWhere([]model.LedgerJournalLinesPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerJournalLinesQueries.selectLedgerJournalLines+" WHERE "+whereQry, defaultLedgerJournalLinesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerJournalLines, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerJournalLines with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerJournalLinesByID] failed get ledgerJournalLines")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerJournalLinesByID(ctx context.Context, primaryID model.LedgerJournalLinesPrimaryID, ledgerJournalLines *model.LedgerJournalLines, ledgerJournalLinesUpdateFields ...LedgerJournalLinesUpdateField) (err error) {
	exists, err := repo.IsExistLedgerJournalLinesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalLines] failed checking ledgerJournalLines whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerJournalLines with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerJournalLines == nil {
		if len(ledgerJournalLinesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerJournalLinesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerJournalLines = &model.LedgerJournalLines{}
	}
	var (
		defaultLedgerJournalLinesUpdateFields = defaultLedgerJournalLinesUpdateFields(*ledgerJournalLines)
		tempUpdateField                       LedgerJournalLinesUpdateFieldList
		selectFields                          = NewLedgerJournalLinesSelectFields()
	)
	if len(ledgerJournalLinesUpdateFields) > 0 {
		for _, updateField := range ledgerJournalLinesUpdateFields {
			if updateField.ledgerJournalLinesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerJournalLinesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerJournalLinesCompositePrimaryKeyWhere([]model.LedgerJournalLinesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerJournalLinesCommand(defaultLedgerJournalLinesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerJournalLinesQueries.updateLedgerJournalLines+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalLines] error when try to update ledgerJournalLines by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerJournalLinesByFilter(ctx context.Context, filter model.Filter, ledgerJournalLinesUpdateFields ...LedgerJournalLinesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerJournalLinesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerJournalLinesUpdateFieldList
		selectFields = NewLedgerJournalLinesSelectFields()
	)
	for _, updateField := range ledgerJournalLinesUpdateFields {
		if updateField.ledgerJournalLinesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerJournalLinesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerJournalLinesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerJournalLinesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_journal_lines\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalLinesByFilter] error when try to update ledgerJournalLines by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerJournalLinesByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerJournalLinesQueries = struct {
		selectLedgerJournalLines      string
		selectCountLedgerJournalLines string
		deleteLedgerJournalLines      string
		updateLedgerJournalLines      string
		insertLedgerJournalLines      string
	}{
		selectLedgerJournalLines:      "SELECT %s FROM \"ledger_journal_lines\"",
		selectCountLedgerJournalLines: "SELECT COUNT(\"id\") FROM \"ledger_journal_lines\"",
		deleteLedgerJournalLines:      "DELETE FROM \"ledger_journal_lines\"",
		updateLedgerJournalLines:      "UPDATE \"ledger_journal_lines\" SET %s ",
		insertLedgerJournalLines:      "INSERT INTO \"ledger_journal_lines\" %s VALUES %s",
	}
)

type LedgerJournalLinesRepository interface {
	CreateLedgerJournalLines(ctx context.Context, ledgerJournalLines *model.LedgerJournalLines, fieldsInsert ...LedgerJournalLinesField) error
	BulkCreateLedgerJournalLines(ctx context.Context, ledgerJournalLinesList []*model.LedgerJournalLines, fieldsInsert ...LedgerJournalLinesField) error
	ResolveLedgerJournalLines(ctx context.Context, selectFields ...LedgerJournalLinesField) (model.LedgerJournalLinesList, error)
	ResolveLedgerJournalLinesByID(ctx context.Context, primaryID model.LedgerJournalLinesPrimaryID, selectFields ...LedgerJournalLinesField) (model.LedgerJournalLines, error)
	UpdateLedgerJournalLinesByID(ctx context.Context, id model.LedgerJournalLinesPrimaryID, ledgerJournalLines *model.LedgerJournalLines, ledgerJournalLinesUpdateFields ...LedgerJournalLinesUpdateField) error
	UpdateLedgerJournalLinesByFilter(ctx context.Context, filter model.Filter, ledgerJournalLinesUpdateFields ...LedgerJournalLinesUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerJournalLines(ctx context.Context, ledgerJournalLinesListMap map[model.LedgerJournalLinesPrimaryID]*model.LedgerJournalLines, LedgerJournalLinessMapUpdateFieldsRequest map[model.LedgerJournalLinesPrimaryID]LedgerJournalLinesUpdateFieldList) (err error)
	DeleteLedgerJournalLinesByID(ctx context.Context, id model.LedgerJournalLinesPrimaryID) error
	BulkDeleteLedgerJournalLinesByIDs(ctx context.Context, ids []model.LedgerJournalLinesPrimaryID) error
	ResolveLedgerJournalLinesByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerJournalLinesFilterResult, err error)
	IsExistLedgerJournalLinesByIDs(ctx context.Context, ids []model.LedgerJournalLinesPrimaryID) (exists bool, notFoundIds []model.LedgerJournalLinesPrimaryID, err error)
	IsExistLedgerJournalLinesByID(ctx context.Context, id model.LedgerJournalLinesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
