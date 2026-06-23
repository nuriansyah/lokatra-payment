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

func composeInsertFieldsAndParamsLedgerReversalLinks(ledgerReversalLinksList []model.LedgerReversalLinks, fieldsInsert ...LedgerReversalLinksField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewLedgerReversalLinksSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, ledgerReversalLinks := range ledgerReversalLinksList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, ledgerReversalLinks.Id)
			case selectField.OriginalJournalId():
				args = append(args, ledgerReversalLinks.OriginalJournalId)
			case selectField.ReversalJournalId():
				args = append(args, ledgerReversalLinks.ReversalJournalId)
			case selectField.ReversalReason():
				args = append(args, ledgerReversalLinks.ReversalReason)
			case selectField.ReversedAt():
				args = append(args, ledgerReversalLinks.ReversedAt)
			case selectField.Metadata():
				args = append(args, ledgerReversalLinks.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, ledgerReversalLinks.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, ledgerReversalLinks.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, ledgerReversalLinks.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, ledgerReversalLinks.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, ledgerReversalLinks.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, ledgerReversalLinks.MetaDeletedBy)

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

func composeLedgerReversalLinksCompositePrimaryKeyWhere(primaryIDs []model.LedgerReversalLinksPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"ledger_reversal_links\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultLedgerReversalLinksSelectFields() string {
	fields := NewLedgerReversalLinksSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeLedgerReversalLinksSelectFields(selectFields ...LedgerReversalLinksField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type LedgerReversalLinksField string
type LedgerReversalLinksFieldList []LedgerReversalLinksField

type LedgerReversalLinksSelectFields struct {
}

func (ss LedgerReversalLinksSelectFields) Id() LedgerReversalLinksField {
	return LedgerReversalLinksField("id")
}

func (ss LedgerReversalLinksSelectFields) OriginalJournalId() LedgerReversalLinksField {
	return LedgerReversalLinksField("original_journal_id")
}

func (ss LedgerReversalLinksSelectFields) ReversalJournalId() LedgerReversalLinksField {
	return LedgerReversalLinksField("reversal_journal_id")
}

func (ss LedgerReversalLinksSelectFields) ReversalReason() LedgerReversalLinksField {
	return LedgerReversalLinksField("reversal_reason")
}

func (ss LedgerReversalLinksSelectFields) ReversedAt() LedgerReversalLinksField {
	return LedgerReversalLinksField("reversed_at")
}

func (ss LedgerReversalLinksSelectFields) Metadata() LedgerReversalLinksField {
	return LedgerReversalLinksField("metadata")
}

func (ss LedgerReversalLinksSelectFields) MetaCreatedAt() LedgerReversalLinksField {
	return LedgerReversalLinksField("meta_created_at")
}

func (ss LedgerReversalLinksSelectFields) MetaCreatedBy() LedgerReversalLinksField {
	return LedgerReversalLinksField("meta_created_by")
}

func (ss LedgerReversalLinksSelectFields) MetaUpdatedAt() LedgerReversalLinksField {
	return LedgerReversalLinksField("meta_updated_at")
}

func (ss LedgerReversalLinksSelectFields) MetaUpdatedBy() LedgerReversalLinksField {
	return LedgerReversalLinksField("meta_updated_by")
}

func (ss LedgerReversalLinksSelectFields) MetaDeletedAt() LedgerReversalLinksField {
	return LedgerReversalLinksField("meta_deleted_at")
}

func (ss LedgerReversalLinksSelectFields) MetaDeletedBy() LedgerReversalLinksField {
	return LedgerReversalLinksField("meta_deleted_by")
}

func (ss LedgerReversalLinksSelectFields) All() LedgerReversalLinksFieldList {
	return []LedgerReversalLinksField{
		ss.Id(),
		ss.OriginalJournalId(),
		ss.ReversalJournalId(),
		ss.ReversalReason(),
		ss.ReversedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewLedgerReversalLinksSelectFields() LedgerReversalLinksSelectFields {
	return LedgerReversalLinksSelectFields{}
}

type LedgerReversalLinksUpdateFieldOption struct {
	useIncrement bool
}
type LedgerReversalLinksUpdateField struct {
	ledgerReversalLinksField LedgerReversalLinksField
	opt                      LedgerReversalLinksUpdateFieldOption
	value                    interface{}
}
type LedgerReversalLinksUpdateFieldList []LedgerReversalLinksUpdateField

func defaultLedgerReversalLinksUpdateFieldOption() LedgerReversalLinksUpdateFieldOption {
	return LedgerReversalLinksUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementLedgerReversalLinksOption(useIncrement bool) func(*LedgerReversalLinksUpdateFieldOption) {
	return func(pcufo *LedgerReversalLinksUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewLedgerReversalLinksUpdateField(field LedgerReversalLinksField, val interface{}, opts ...func(*LedgerReversalLinksUpdateFieldOption)) LedgerReversalLinksUpdateField {
	defaultOpt := defaultLedgerReversalLinksUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return LedgerReversalLinksUpdateField{
		ledgerReversalLinksField: field,
		value:                    val,
		opt:                      defaultOpt,
	}
}
func defaultLedgerReversalLinksUpdateFields(ledgerReversalLinks model.LedgerReversalLinks) (ledgerReversalLinksUpdateFieldList LedgerReversalLinksUpdateFieldList) {
	selectFields := NewLedgerReversalLinksSelectFields()
	ledgerReversalLinksUpdateFieldList = append(ledgerReversalLinksUpdateFieldList,
		NewLedgerReversalLinksUpdateField(selectFields.Id(), ledgerReversalLinks.Id),
		NewLedgerReversalLinksUpdateField(selectFields.OriginalJournalId(), ledgerReversalLinks.OriginalJournalId),
		NewLedgerReversalLinksUpdateField(selectFields.ReversalJournalId(), ledgerReversalLinks.ReversalJournalId),
		NewLedgerReversalLinksUpdateField(selectFields.ReversalReason(), ledgerReversalLinks.ReversalReason),
		NewLedgerReversalLinksUpdateField(selectFields.ReversedAt(), ledgerReversalLinks.ReversedAt),
		NewLedgerReversalLinksUpdateField(selectFields.Metadata(), ledgerReversalLinks.Metadata),
		NewLedgerReversalLinksUpdateField(selectFields.MetaCreatedAt(), ledgerReversalLinks.MetaCreatedAt),
		NewLedgerReversalLinksUpdateField(selectFields.MetaCreatedBy(), ledgerReversalLinks.MetaCreatedBy),
		NewLedgerReversalLinksUpdateField(selectFields.MetaUpdatedAt(), ledgerReversalLinks.MetaUpdatedAt),
		NewLedgerReversalLinksUpdateField(selectFields.MetaUpdatedBy(), ledgerReversalLinks.MetaUpdatedBy),
		NewLedgerReversalLinksUpdateField(selectFields.MetaDeletedAt(), ledgerReversalLinks.MetaDeletedAt),
		NewLedgerReversalLinksUpdateField(selectFields.MetaDeletedBy(), ledgerReversalLinks.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsLedgerReversalLinksCommand(ledgerReversalLinksUpdateFieldList LedgerReversalLinksUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range ledgerReversalLinksUpdateFieldList {
		field := string(updateField.ledgerReversalLinksField)
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

func (repo *RepositoryImpl) BulkCreateLedgerReversalLinks(ctx context.Context, ledgerReversalLinksList []*model.LedgerReversalLinks, fieldsInsert ...LedgerReversalLinksField) (err error) {
	var (
		fieldsStr                    string
		valueListStr                 []string
		argsList                     []interface{}
		primaryIds                   []model.LedgerReversalLinksPrimaryID
		ledgerReversalLinksValueList []model.LedgerReversalLinks
	)

	if len(fieldsInsert) == 0 {
		selectField := NewLedgerReversalLinksSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, ledgerReversalLinks := range ledgerReversalLinksList {

		primaryIds = append(primaryIds, ledgerReversalLinks.ToLedgerReversalLinksPrimaryID())

		ledgerReversalLinksValueList = append(ledgerReversalLinksValueList, *ledgerReversalLinks)
	}

	_, notFoundIds, err := repo.IsExistLedgerReversalLinksByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerReversalLinks] failed checking ledgerReversalLinks whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.LedgerReversalLinksPrimaryID{}
		mapNotFoundIds := map[model.LedgerReversalLinksPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "ledgerReversalLinks", fmt.Sprintf("ledgerReversalLinks with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsLedgerReversalLinks(ledgerReversalLinksValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(ledgerReversalLinksQueries.insertLedgerReversalLinks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateLedgerReversalLinks] failed exec create ledgerReversalLinks query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteLedgerReversalLinksByIDs(ctx context.Context, primaryIDs []model.LedgerReversalLinksPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistLedgerReversalLinksByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerReversalLinksByIDs] failed checking ledgerReversalLinks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerReversalLinks with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_reversal_links\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(ledgerReversalLinksQueries.deleteLedgerReversalLinks + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerReversalLinksByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteLedgerReversalLinksByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistLedgerReversalLinksByIDs(ctx context.Context, ids []model.LedgerReversalLinksPrimaryID) (exists bool, notFoundIds []model.LedgerReversalLinksPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"ledger_reversal_links\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(ledgerReversalLinksQueries.selectLedgerReversalLinks, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerReversalLinksByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.LedgerReversalLinksPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerReversalLinksByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.LedgerReversalLinksPrimaryID]bool{}
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

// BulkUpdateLedgerReversalLinks is used to bulk update ledgerReversalLinks, by default it will update all field
// if want to update specific field, then fill ledgerReversalLinkssMapUpdateFieldsRequest else please fill ledgerReversalLinkssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateLedgerReversalLinks(ctx context.Context, ledgerReversalLinkssMap map[model.LedgerReversalLinksPrimaryID]*model.LedgerReversalLinks, ledgerReversalLinkssMapUpdateFieldsRequest map[model.LedgerReversalLinksPrimaryID]LedgerReversalLinksUpdateFieldList) (err error) {
	if len(ledgerReversalLinkssMap) == 0 && len(ledgerReversalLinkssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		ledgerReversalLinkssMapUpdateField map[model.LedgerReversalLinksPrimaryID]LedgerReversalLinksUpdateFieldList = map[model.LedgerReversalLinksPrimaryID]LedgerReversalLinksUpdateFieldList{}
		asTableValues                      string                                                                    = "myvalues"
	)

	if len(ledgerReversalLinkssMap) > 0 {
		for id, ledgerReversalLinks := range ledgerReversalLinkssMap {
			if ledgerReversalLinks == nil {
				log.Error().Err(err).Msg("[BulkUpdateLedgerReversalLinks] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			ledgerReversalLinkssMapUpdateField[id] = defaultLedgerReversalLinksUpdateFields(*ledgerReversalLinks)
		}
	} else {
		ledgerReversalLinkssMapUpdateField = ledgerReversalLinkssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateLedgerReversalLinksQuery(ledgerReversalLinkssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistLedgerReversalLinksByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerReversalLinks] failed checking ledgerReversalLinks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerReversalLinks with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeLedgerReversalLinksCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"ledger_reversal_links\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateLedgerReversalLinks] failed exec query")
	}
	return
}

type LedgerReversalLinksFieldParameter struct {
	param string
	args  []interface{}
}

func NewLedgerReversalLinksFieldParameter(param string, args ...interface{}) LedgerReversalLinksFieldParameter {
	return LedgerReversalLinksFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateLedgerReversalLinksQuery(mapLedgerReversalLinkss map[model.LedgerReversalLinksPrimaryID]LedgerReversalLinksUpdateFieldList, asTableValues string) (primaryIDs []model.LedgerReversalLinksPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.LedgerReversalLinksPrimaryID]map[string]interface{}{}
	ledgerReversalLinksSelectFields := NewLedgerReversalLinksSelectFields()
	for id, updateFields := range mapLedgerReversalLinkss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.ledgerReversalLinksField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapLedgerReversalLinkss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetLedgerReversalLinksFieldType(updateField.ledgerReversalLinksField)))
			args = append(args, fields[string(updateField.ledgerReversalLinksField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.ledgerReversalLinksField))
		if updateField.ledgerReversalLinksField == ledgerReversalLinksSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.ledgerReversalLinksField, asTableValues, updateField.ledgerReversalLinksField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.ledgerReversalLinksField,
				"\"ledger_reversal_links\"", updateField.ledgerReversalLinksField,
				asTableValues, updateField.ledgerReversalLinksField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeLedgerReversalLinksCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.LedgerReversalLinksPrimaryID, asTableValue string) (whereQry string) {
	ledgerReversalLinksSelectFields := NewLedgerReversalLinksSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"ledger_reversal_links\".\"id\" = %s.\"id\"::"+GetLedgerReversalLinksFieldType(ledgerReversalLinksSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetLedgerReversalLinksFieldType(ledgerReversalLinksField LedgerReversalLinksField) string {
	selectLedgerReversalLinksFields := NewLedgerReversalLinksSelectFields()
	switch ledgerReversalLinksField {

	case selectLedgerReversalLinksFields.Id():
		return "uuid"

	case selectLedgerReversalLinksFields.OriginalJournalId():
		return "uuid"

	case selectLedgerReversalLinksFields.ReversalJournalId():
		return "uuid"

	case selectLedgerReversalLinksFields.ReversalReason():
		return "text"

	case selectLedgerReversalLinksFields.ReversedAt():
		return "timestamptz"

	case selectLedgerReversalLinksFields.Metadata():
		return "jsonb"

	case selectLedgerReversalLinksFields.MetaCreatedAt():
		return "timestamptz"

	case selectLedgerReversalLinksFields.MetaCreatedBy():
		return "uuid"

	case selectLedgerReversalLinksFields.MetaUpdatedAt():
		return "timestamptz"

	case selectLedgerReversalLinksFields.MetaUpdatedBy():
		return "uuid"

	case selectLedgerReversalLinksFields.MetaDeletedAt():
		return "timestamptz"

	case selectLedgerReversalLinksFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateLedgerReversalLinks(ctx context.Context, ledgerReversalLinks *model.LedgerReversalLinks, fieldsInsert ...LedgerReversalLinksField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewLedgerReversalLinksSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.LedgerReversalLinksPrimaryID{
		Id: ledgerReversalLinks.Id,
	}
	exists, err := repo.IsExistLedgerReversalLinksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerReversalLinks] failed checking ledgerReversalLinks whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "ledgerReversalLinks", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsLedgerReversalLinks([]model.LedgerReversalLinks{*ledgerReversalLinks}, fieldsInsert...)
	commandQuery := fmt.Sprintf(ledgerReversalLinksQueries.insertLedgerReversalLinks, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateLedgerReversalLinks] failed exec create ledgerReversalLinks query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteLedgerReversalLinksByID(ctx context.Context, primaryID model.LedgerReversalLinksPrimaryID) (err error) {
	exists, err := repo.IsExistLedgerReversalLinksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerReversalLinksByID] failed checking ledgerReversalLinks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerReversalLinks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeLedgerReversalLinksCompositePrimaryKeyWhere([]model.LedgerReversalLinksPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(ledgerReversalLinksQueries.deleteLedgerReversalLinks + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteLedgerReversalLinksByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerReversalLinksByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerReversalLinksFilterResult, err error) {
	query, args, err := composeLedgerReversalLinksFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerReversalLinksByFilter] failed compose ledgerReversalLinks filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerReversalLinksByFilter] failed get ledgerReversalLinks by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeLedgerReversalLinksFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.LedgerReversalLinksFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeLedgerReversalLinksFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeLedgerReversalLinksSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeLedgerReversalLinksFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 12 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewLedgerReversalLinksFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 12+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["original_journal_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"original_journal_id\"")
			selectedColumns["original_journal_id"] = struct{}{}
		}
		if _, selected := selectedColumns["reversal_journal_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"reversal_journal_id\"")
			selectedColumns["reversal_journal_id"] = struct{}{}
		}
		if _, selected := selectedColumns["reversal_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"reversal_reason\"")
			selectedColumns["reversal_reason"] = struct{}{}
		}
		if _, selected := selectedColumns["reversed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"reversed_at\"")
			selectedColumns["reversed_at"] = struct{}{}
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

type ledgerReversalLinksFilterPlaceholder struct {
	index int
}

func (p *ledgerReversalLinksFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeLedgerReversalLinksFilterPredicate(filterField model.FilterField, placeholders *ledgerReversalLinksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewLedgerReversalLinksFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeLedgerReversalLinksFilterSQLExpr(spec)
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

func composeLedgerReversalLinksFilterGroup(group model.FilterGroup, placeholders *ledgerReversalLinksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeLedgerReversalLinksFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeLedgerReversalLinksFilterGroup(child, placeholders, args, requiredJoins)
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

func composeLedgerReversalLinksFilterWhereQueries(filter model.Filter, placeholders *ledgerReversalLinksFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeLedgerReversalLinksFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeLedgerReversalLinksFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeLedgerReversalLinksFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateLedgerReversalLinksFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeLedgerReversalLinksSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeLedgerReversalLinksFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := ledgerReversalLinksFilterPlaceholder{index: 1}
	whereQueries, err := composeLedgerReversalLinksFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewLedgerReversalLinksFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeLedgerReversalLinksFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeLedgerReversalLinksSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"ledger_reversal_links\" base%s", strings.Join(selectColumns, ","), composeLedgerReversalLinksFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistLedgerReversalLinksByID(ctx context.Context, primaryID model.LedgerReversalLinksPrimaryID) (exists bool, err error) {
	whereQuery, params := composeLedgerReversalLinksCompositePrimaryKeyWhere([]model.LedgerReversalLinksPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", ledgerReversalLinksQueries.selectCountLedgerReversalLinks, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistLedgerReversalLinksByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerReversalLinks(ctx context.Context, selectFields ...LedgerReversalLinksField) (ledgerReversalLinksList model.LedgerReversalLinksList, err error) {
	var (
		defaultLedgerReversalLinksSelectFields = defaultLedgerReversalLinksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerReversalLinksSelectFields = composeLedgerReversalLinksSelectFields(selectFields...)
	}
	query := fmt.Sprintf(ledgerReversalLinksQueries.selectLedgerReversalLinks, defaultLedgerReversalLinksSelectFields)

	err = repo.db.Read.SelectContext(ctx, &ledgerReversalLinksList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveLedgerReversalLinks] failed get ledgerReversalLinks list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveLedgerReversalLinksByID(ctx context.Context, primaryID model.LedgerReversalLinksPrimaryID, selectFields ...LedgerReversalLinksField) (ledgerReversalLinks model.LedgerReversalLinks, err error) {
	var (
		defaultLedgerReversalLinksSelectFields = defaultLedgerReversalLinksSelectFields()
	)
	if len(selectFields) > 0 {
		defaultLedgerReversalLinksSelectFields = composeLedgerReversalLinksSelectFields(selectFields...)
	}
	whereQry, params := composeLedgerReversalLinksCompositePrimaryKeyWhere([]model.LedgerReversalLinksPrimaryID{primaryID})
	query := fmt.Sprintf(ledgerReversalLinksQueries.selectLedgerReversalLinks+" WHERE "+whereQry, defaultLedgerReversalLinksSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &ledgerReversalLinks, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("ledgerReversalLinks with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveLedgerReversalLinksByID] failed get ledgerReversalLinks")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateLedgerReversalLinksByID(ctx context.Context, primaryID model.LedgerReversalLinksPrimaryID, ledgerReversalLinks *model.LedgerReversalLinks, ledgerReversalLinksUpdateFields ...LedgerReversalLinksUpdateField) (err error) {
	exists, err := repo.IsExistLedgerReversalLinksByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerReversalLinks] failed checking ledgerReversalLinks whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("ledgerReversalLinks with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if ledgerReversalLinks == nil {
		if len(ledgerReversalLinksUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateLedgerReversalLinksByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		ledgerReversalLinks = &model.LedgerReversalLinks{}
	}
	var (
		defaultLedgerReversalLinksUpdateFields = defaultLedgerReversalLinksUpdateFields(*ledgerReversalLinks)
		tempUpdateField                        LedgerReversalLinksUpdateFieldList
		selectFields                           = NewLedgerReversalLinksSelectFields()
	)
	if len(ledgerReversalLinksUpdateFields) > 0 {
		for _, updateField := range ledgerReversalLinksUpdateFields {
			if updateField.ledgerReversalLinksField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultLedgerReversalLinksUpdateFields = tempUpdateField
	}
	whereQuery, params := composeLedgerReversalLinksCompositePrimaryKeyWhere([]model.LedgerReversalLinksPrimaryID{primaryID})
	fields, args := composeUpdateFieldsLedgerReversalLinksCommand(defaultLedgerReversalLinksUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(ledgerReversalLinksQueries.updateLedgerReversalLinks+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerReversalLinks] error when try to update ledgerReversalLinks by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateLedgerReversalLinksByFilter(ctx context.Context, filter model.Filter, ledgerReversalLinksUpdateFields ...LedgerReversalLinksUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(ledgerReversalLinksUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields LedgerReversalLinksUpdateFieldList
		selectFields = NewLedgerReversalLinksSelectFields()
	)
	for _, updateField := range ledgerReversalLinksUpdateFields {
		if updateField.ledgerReversalLinksField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsLedgerReversalLinksCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := ledgerReversalLinksFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeLedgerReversalLinksFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"ledger_reversal_links\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerReversalLinksByFilter] error when try to update ledgerReversalLinks by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateLedgerReversalLinksByFilter] failed get rows affected")
	}
	return
}

var (
	ledgerReversalLinksQueries = struct {
		selectLedgerReversalLinks      string
		selectCountLedgerReversalLinks string
		deleteLedgerReversalLinks      string
		updateLedgerReversalLinks      string
		insertLedgerReversalLinks      string
	}{
		selectLedgerReversalLinks:      "SELECT %s FROM \"ledger_reversal_links\"",
		selectCountLedgerReversalLinks: "SELECT COUNT(\"id\") FROM \"ledger_reversal_links\"",
		deleteLedgerReversalLinks:      "DELETE FROM \"ledger_reversal_links\"",
		updateLedgerReversalLinks:      "UPDATE \"ledger_reversal_links\" SET %s ",
		insertLedgerReversalLinks:      "INSERT INTO \"ledger_reversal_links\" %s VALUES %s",
	}
)

type LedgerReversalLinksRepository interface {
	CreateLedgerReversalLinks(ctx context.Context, ledgerReversalLinks *model.LedgerReversalLinks, fieldsInsert ...LedgerReversalLinksField) error
	BulkCreateLedgerReversalLinks(ctx context.Context, ledgerReversalLinksList []*model.LedgerReversalLinks, fieldsInsert ...LedgerReversalLinksField) error
	ResolveLedgerReversalLinks(ctx context.Context, selectFields ...LedgerReversalLinksField) (model.LedgerReversalLinksList, error)
	ResolveLedgerReversalLinksByID(ctx context.Context, primaryID model.LedgerReversalLinksPrimaryID, selectFields ...LedgerReversalLinksField) (model.LedgerReversalLinks, error)
	UpdateLedgerReversalLinksByID(ctx context.Context, id model.LedgerReversalLinksPrimaryID, ledgerReversalLinks *model.LedgerReversalLinks, ledgerReversalLinksUpdateFields ...LedgerReversalLinksUpdateField) error
	UpdateLedgerReversalLinksByFilter(ctx context.Context, filter model.Filter, ledgerReversalLinksUpdateFields ...LedgerReversalLinksUpdateField) (rowsAffected int64, err error)
	BulkUpdateLedgerReversalLinks(ctx context.Context, ledgerReversalLinksListMap map[model.LedgerReversalLinksPrimaryID]*model.LedgerReversalLinks, LedgerReversalLinkssMapUpdateFieldsRequest map[model.LedgerReversalLinksPrimaryID]LedgerReversalLinksUpdateFieldList) (err error)
	DeleteLedgerReversalLinksByID(ctx context.Context, id model.LedgerReversalLinksPrimaryID) error
	BulkDeleteLedgerReversalLinksByIDs(ctx context.Context, ids []model.LedgerReversalLinksPrimaryID) error
	ResolveLedgerReversalLinksByFilter(ctx context.Context, filter model.Filter) (result []model.LedgerReversalLinksFilterResult, err error)
	IsExistLedgerReversalLinksByIDs(ctx context.Context, ids []model.LedgerReversalLinksPrimaryID) (exists bool, notFoundIds []model.LedgerReversalLinksPrimaryID, err error)
	IsExistLedgerReversalLinksByID(ctx context.Context, id model.LedgerReversalLinksPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
