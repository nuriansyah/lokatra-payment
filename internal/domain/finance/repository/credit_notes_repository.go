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

func composeInsertFieldsAndParamsCreditNotes(creditNotesList []model.CreditNotes, fieldsInsert ...CreditNotesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewCreditNotesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, creditNotes := range creditNotesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, creditNotes.Id)
			case selectField.CreditNoteNo():
				args = append(args, creditNotes.CreditNoteNo)
			case selectField.TaxInvoiceId():
				args = append(args, creditNotes.TaxInvoiceId)
			case selectField.SourceType():
				args = append(args, creditNotes.SourceType)
			case selectField.SourceId():
				args = append(args, creditNotes.SourceId)
			case selectField.CurrencyCode():
				args = append(args, creditNotes.CurrencyCode)
			case selectField.TaxableAmount():
				args = append(args, creditNotes.TaxableAmount)
			case selectField.TaxAmount():
				args = append(args, creditNotes.TaxAmount)
			case selectField.TotalAmount():
				args = append(args, creditNotes.TotalAmount)
			case selectField.ReasonCode():
				args = append(args, creditNotes.ReasonCode)
			case selectField.ReasonDetail():
				args = append(args, creditNotes.ReasonDetail)
			case selectField.IssuedAt():
				args = append(args, creditNotes.IssuedAt)
			case selectField.Metadata():
				args = append(args, creditNotes.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, creditNotes.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, creditNotes.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, creditNotes.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, creditNotes.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, creditNotes.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, creditNotes.MetaDeletedBy)

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

func composeCreditNotesCompositePrimaryKeyWhere(primaryIDs []model.CreditNotesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"credit_notes\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultCreditNotesSelectFields() string {
	fields := NewCreditNotesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeCreditNotesSelectFields(selectFields ...CreditNotesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type CreditNotesField string
type CreditNotesFieldList []CreditNotesField

type CreditNotesSelectFields struct {
}

func (ss CreditNotesSelectFields) Id() CreditNotesField {
	return CreditNotesField("id")
}

func (ss CreditNotesSelectFields) CreditNoteNo() CreditNotesField {
	return CreditNotesField("credit_note_no")
}

func (ss CreditNotesSelectFields) TaxInvoiceId() CreditNotesField {
	return CreditNotesField("tax_invoice_id")
}

func (ss CreditNotesSelectFields) SourceType() CreditNotesField {
	return CreditNotesField("source_type")
}

func (ss CreditNotesSelectFields) SourceId() CreditNotesField {
	return CreditNotesField("source_id")
}

func (ss CreditNotesSelectFields) CurrencyCode() CreditNotesField {
	return CreditNotesField("currency_code")
}

func (ss CreditNotesSelectFields) TaxableAmount() CreditNotesField {
	return CreditNotesField("taxable_amount")
}

func (ss CreditNotesSelectFields) TaxAmount() CreditNotesField {
	return CreditNotesField("tax_amount")
}

func (ss CreditNotesSelectFields) TotalAmount() CreditNotesField {
	return CreditNotesField("total_amount")
}

func (ss CreditNotesSelectFields) ReasonCode() CreditNotesField {
	return CreditNotesField("reason_code")
}

func (ss CreditNotesSelectFields) ReasonDetail() CreditNotesField {
	return CreditNotesField("reason_detail")
}

func (ss CreditNotesSelectFields) IssuedAt() CreditNotesField {
	return CreditNotesField("issued_at")
}

func (ss CreditNotesSelectFields) Metadata() CreditNotesField {
	return CreditNotesField("metadata")
}

func (ss CreditNotesSelectFields) MetaCreatedAt() CreditNotesField {
	return CreditNotesField("meta_created_at")
}

func (ss CreditNotesSelectFields) MetaCreatedBy() CreditNotesField {
	return CreditNotesField("meta_created_by")
}

func (ss CreditNotesSelectFields) MetaUpdatedAt() CreditNotesField {
	return CreditNotesField("meta_updated_at")
}

func (ss CreditNotesSelectFields) MetaUpdatedBy() CreditNotesField {
	return CreditNotesField("meta_updated_by")
}

func (ss CreditNotesSelectFields) MetaDeletedAt() CreditNotesField {
	return CreditNotesField("meta_deleted_at")
}

func (ss CreditNotesSelectFields) MetaDeletedBy() CreditNotesField {
	return CreditNotesField("meta_deleted_by")
}

func (ss CreditNotesSelectFields) All() CreditNotesFieldList {
	return []CreditNotesField{
		ss.Id(),
		ss.CreditNoteNo(),
		ss.TaxInvoiceId(),
		ss.SourceType(),
		ss.SourceId(),
		ss.CurrencyCode(),
		ss.TaxableAmount(),
		ss.TaxAmount(),
		ss.TotalAmount(),
		ss.ReasonCode(),
		ss.ReasonDetail(),
		ss.IssuedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewCreditNotesSelectFields() CreditNotesSelectFields {
	return CreditNotesSelectFields{}
}

type CreditNotesUpdateFieldOption struct {
	useIncrement bool
}
type CreditNotesUpdateField struct {
	creditNotesField CreditNotesField
	opt              CreditNotesUpdateFieldOption
	value            interface{}
}
type CreditNotesUpdateFieldList []CreditNotesUpdateField

func defaultCreditNotesUpdateFieldOption() CreditNotesUpdateFieldOption {
	return CreditNotesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementCreditNotesOption(useIncrement bool) func(*CreditNotesUpdateFieldOption) {
	return func(pcufo *CreditNotesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewCreditNotesUpdateField(field CreditNotesField, val interface{}, opts ...func(*CreditNotesUpdateFieldOption)) CreditNotesUpdateField {
	defaultOpt := defaultCreditNotesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return CreditNotesUpdateField{
		creditNotesField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultCreditNotesUpdateFields(creditNotes model.CreditNotes) (creditNotesUpdateFieldList CreditNotesUpdateFieldList) {
	selectFields := NewCreditNotesSelectFields()
	creditNotesUpdateFieldList = append(creditNotesUpdateFieldList,
		NewCreditNotesUpdateField(selectFields.Id(), creditNotes.Id),
		NewCreditNotesUpdateField(selectFields.CreditNoteNo(), creditNotes.CreditNoteNo),
		NewCreditNotesUpdateField(selectFields.TaxInvoiceId(), creditNotes.TaxInvoiceId),
		NewCreditNotesUpdateField(selectFields.SourceType(), creditNotes.SourceType),
		NewCreditNotesUpdateField(selectFields.SourceId(), creditNotes.SourceId),
		NewCreditNotesUpdateField(selectFields.CurrencyCode(), creditNotes.CurrencyCode),
		NewCreditNotesUpdateField(selectFields.TaxableAmount(), creditNotes.TaxableAmount),
		NewCreditNotesUpdateField(selectFields.TaxAmount(), creditNotes.TaxAmount),
		NewCreditNotesUpdateField(selectFields.TotalAmount(), creditNotes.TotalAmount),
		NewCreditNotesUpdateField(selectFields.ReasonCode(), creditNotes.ReasonCode),
		NewCreditNotesUpdateField(selectFields.ReasonDetail(), creditNotes.ReasonDetail),
		NewCreditNotesUpdateField(selectFields.IssuedAt(), creditNotes.IssuedAt),
		NewCreditNotesUpdateField(selectFields.Metadata(), creditNotes.Metadata),
		NewCreditNotesUpdateField(selectFields.MetaCreatedAt(), creditNotes.MetaCreatedAt),
		NewCreditNotesUpdateField(selectFields.MetaCreatedBy(), creditNotes.MetaCreatedBy),
		NewCreditNotesUpdateField(selectFields.MetaUpdatedAt(), creditNotes.MetaUpdatedAt),
		NewCreditNotesUpdateField(selectFields.MetaUpdatedBy(), creditNotes.MetaUpdatedBy),
		NewCreditNotesUpdateField(selectFields.MetaDeletedAt(), creditNotes.MetaDeletedAt),
		NewCreditNotesUpdateField(selectFields.MetaDeletedBy(), creditNotes.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsCreditNotesCommand(creditNotesUpdateFieldList CreditNotesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range creditNotesUpdateFieldList {
		field := string(updateField.creditNotesField)
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

func (repo *RepositoryImpl) BulkCreateCreditNotes(ctx context.Context, creditNotesList []*model.CreditNotes, fieldsInsert ...CreditNotesField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.CreditNotesPrimaryID
		creditNotesValueList []model.CreditNotes
	)

	if len(fieldsInsert) == 0 {
		selectField := NewCreditNotesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, creditNotes := range creditNotesList {

		primaryIds = append(primaryIds, creditNotes.ToCreditNotesPrimaryID())

		creditNotesValueList = append(creditNotesValueList, *creditNotes)
	}

	_, notFoundIds, err := repo.IsExistCreditNotesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateCreditNotes] failed checking creditNotes whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.CreditNotesPrimaryID{}
		mapNotFoundIds := map[model.CreditNotesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "creditNotes", fmt.Sprintf("creditNotes with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsCreditNotes(creditNotesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(creditNotesQueries.insertCreditNotes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateCreditNotes] failed exec create creditNotes query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteCreditNotesByIDs(ctx context.Context, primaryIDs []model.CreditNotesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistCreditNotesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCreditNotesByIDs] failed checking creditNotes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("creditNotes with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"credit_notes\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(creditNotesQueries.deleteCreditNotes + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCreditNotesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteCreditNotesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistCreditNotesByIDs(ctx context.Context, ids []model.CreditNotesPrimaryID) (exists bool, notFoundIds []model.CreditNotesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"credit_notes\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(creditNotesQueries.selectCreditNotes, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCreditNotesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.CreditNotesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCreditNotesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.CreditNotesPrimaryID]bool{}
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

// BulkUpdateCreditNotes is used to bulk update creditNotes, by default it will update all field
// if want to update specific field, then fill creditNotessMapUpdateFieldsRequest else please fill creditNotessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateCreditNotes(ctx context.Context, creditNotessMap map[model.CreditNotesPrimaryID]*model.CreditNotes, creditNotessMapUpdateFieldsRequest map[model.CreditNotesPrimaryID]CreditNotesUpdateFieldList) (err error) {
	if len(creditNotessMap) == 0 && len(creditNotessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		creditNotessMapUpdateField map[model.CreditNotesPrimaryID]CreditNotesUpdateFieldList = map[model.CreditNotesPrimaryID]CreditNotesUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(creditNotessMap) > 0 {
		for id, creditNotes := range creditNotessMap {
			if creditNotes == nil {
				log.Error().Err(err).Msg("[BulkUpdateCreditNotes] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			creditNotessMapUpdateField[id] = defaultCreditNotesUpdateFields(*creditNotes)
		}
	} else {
		creditNotessMapUpdateField = creditNotessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateCreditNotesQuery(creditNotessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistCreditNotesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateCreditNotes] failed checking creditNotes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("creditNotes with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeCreditNotesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"credit_notes\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateCreditNotes] failed exec query")
	}
	return
}

type CreditNotesFieldParameter struct {
	param string
	args  []interface{}
}

func NewCreditNotesFieldParameter(param string, args ...interface{}) CreditNotesFieldParameter {
	return CreditNotesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateCreditNotesQuery(mapCreditNotess map[model.CreditNotesPrimaryID]CreditNotesUpdateFieldList, asTableValues string) (primaryIDs []model.CreditNotesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.CreditNotesPrimaryID]map[string]interface{}{}
	creditNotesSelectFields := NewCreditNotesSelectFields()
	for id, updateFields := range mapCreditNotess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.creditNotesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapCreditNotess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetCreditNotesFieldType(updateField.creditNotesField)))
			args = append(args, fields[string(updateField.creditNotesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.creditNotesField))
		if updateField.creditNotesField == creditNotesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.creditNotesField, asTableValues, updateField.creditNotesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.creditNotesField,
				"\"credit_notes\"", updateField.creditNotesField,
				asTableValues, updateField.creditNotesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeCreditNotesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.CreditNotesPrimaryID, asTableValue string) (whereQry string) {
	creditNotesSelectFields := NewCreditNotesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"credit_notes\".\"id\" = %s.\"id\"::"+GetCreditNotesFieldType(creditNotesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetCreditNotesFieldType(creditNotesField CreditNotesField) string {
	selectCreditNotesFields := NewCreditNotesSelectFields()
	switch creditNotesField {

	case selectCreditNotesFields.Id():
		return "uuid"

	case selectCreditNotesFields.CreditNoteNo():
		return "text"

	case selectCreditNotesFields.TaxInvoiceId():
		return "uuid"

	case selectCreditNotesFields.SourceType():
		return "text"

	case selectCreditNotesFields.SourceId():
		return "uuid"

	case selectCreditNotesFields.CurrencyCode():
		return "text"

	case selectCreditNotesFields.TaxableAmount():
		return "numeric"

	case selectCreditNotesFields.TaxAmount():
		return "numeric"

	case selectCreditNotesFields.TotalAmount():
		return "numeric"

	case selectCreditNotesFields.ReasonCode():
		return "text"

	case selectCreditNotesFields.ReasonDetail():
		return "text"

	case selectCreditNotesFields.IssuedAt():
		return "timestamptz"

	case selectCreditNotesFields.Metadata():
		return "jsonb"

	case selectCreditNotesFields.MetaCreatedAt():
		return "timestamptz"

	case selectCreditNotesFields.MetaCreatedBy():
		return "uuid"

	case selectCreditNotesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectCreditNotesFields.MetaUpdatedBy():
		return "uuid"

	case selectCreditNotesFields.MetaDeletedAt():
		return "timestamptz"

	case selectCreditNotesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateCreditNotes(ctx context.Context, creditNotes *model.CreditNotes, fieldsInsert ...CreditNotesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewCreditNotesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.CreditNotesPrimaryID{
		Id: creditNotes.Id,
	}
	exists, err := repo.IsExistCreditNotesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateCreditNotes] failed checking creditNotes whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "creditNotes", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsCreditNotes([]model.CreditNotes{*creditNotes}, fieldsInsert...)
	commandQuery := fmt.Sprintf(creditNotesQueries.insertCreditNotes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateCreditNotes] failed exec create creditNotes query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteCreditNotesByID(ctx context.Context, primaryID model.CreditNotesPrimaryID) (err error) {
	exists, err := repo.IsExistCreditNotesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteCreditNotesByID] failed checking creditNotes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("creditNotes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeCreditNotesCompositePrimaryKeyWhere([]model.CreditNotesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(creditNotesQueries.deleteCreditNotes + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteCreditNotesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveCreditNotesByFilter(ctx context.Context, filter model.Filter) (result []model.CreditNotesFilterResult, err error) {
	query, args, err := composeCreditNotesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCreditNotesByFilter] failed compose creditNotes filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCreditNotesByFilter] failed get creditNotes by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeCreditNotesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.CreditNotesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeCreditNotesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeCreditNotesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeCreditNotesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewCreditNotesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["credit_note_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"credit_note_no\"")
			selectedColumns["credit_note_no"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_invoice_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_invoice_id\"")
			selectedColumns["tax_invoice_id"] = struct{}{}
		}
		if _, selected := selectedColumns["source_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_type\"")
			selectedColumns["source_type"] = struct{}{}
		}
		if _, selected := selectedColumns["source_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"source_id\"")
			selectedColumns["source_id"] = struct{}{}
		}
		if _, selected := selectedColumns["currency_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency_code\"")
			selectedColumns["currency_code"] = struct{}{}
		}
		if _, selected := selectedColumns["taxable_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"taxable_amount\"")
			selectedColumns["taxable_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["tax_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"tax_amount\"")
			selectedColumns["tax_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["total_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"total_amount\"")
			selectedColumns["total_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_code\"")
			selectedColumns["reason_code"] = struct{}{}
		}
		if _, selected := selectedColumns["reason_detail"]; !selected {
			selectColumns = append(selectColumns, "base.\"reason_detail\"")
			selectedColumns["reason_detail"] = struct{}{}
		}
		if _, selected := selectedColumns["issued_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"issued_at\"")
			selectedColumns["issued_at"] = struct{}{}
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

type creditNotesFilterPlaceholder struct {
	index int
}

func (p *creditNotesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeCreditNotesFilterPredicate(filterField model.FilterField, placeholders *creditNotesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewCreditNotesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeCreditNotesFilterSQLExpr(spec)
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

func composeCreditNotesFilterGroup(group model.FilterGroup, placeholders *creditNotesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeCreditNotesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeCreditNotesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeCreditNotesFilterWhereQueries(filter model.Filter, placeholders *creditNotesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeCreditNotesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeCreditNotesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeCreditNotesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateCreditNotesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeCreditNotesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeCreditNotesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := creditNotesFilterPlaceholder{index: 1}
	whereQueries, err := composeCreditNotesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewCreditNotesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeCreditNotesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeCreditNotesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"credit_notes\" base%s", strings.Join(selectColumns, ","), composeCreditNotesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistCreditNotesByID(ctx context.Context, primaryID model.CreditNotesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeCreditNotesCompositePrimaryKeyWhere([]model.CreditNotesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", creditNotesQueries.selectCountCreditNotes, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistCreditNotesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveCreditNotes(ctx context.Context, selectFields ...CreditNotesField) (creditNotesList model.CreditNotesList, err error) {
	var (
		defaultCreditNotesSelectFields = defaultCreditNotesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultCreditNotesSelectFields = composeCreditNotesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(creditNotesQueries.selectCreditNotes, defaultCreditNotesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &creditNotesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveCreditNotes] failed get creditNotes list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveCreditNotesByID(ctx context.Context, primaryID model.CreditNotesPrimaryID, selectFields ...CreditNotesField) (creditNotes model.CreditNotes, err error) {
	var (
		defaultCreditNotesSelectFields = defaultCreditNotesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultCreditNotesSelectFields = composeCreditNotesSelectFields(selectFields...)
	}
	whereQry, params := composeCreditNotesCompositePrimaryKeyWhere([]model.CreditNotesPrimaryID{primaryID})
	query := fmt.Sprintf(creditNotesQueries.selectCreditNotes+" WHERE "+whereQry, defaultCreditNotesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &creditNotes, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("creditNotes with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveCreditNotesByID] failed get creditNotes")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateCreditNotesByID(ctx context.Context, primaryID model.CreditNotesPrimaryID, creditNotes *model.CreditNotes, creditNotesUpdateFields ...CreditNotesUpdateField) (err error) {
	exists, err := repo.IsExistCreditNotesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCreditNotes] failed checking creditNotes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("creditNotes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if creditNotes == nil {
		if len(creditNotesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateCreditNotesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		creditNotes = &model.CreditNotes{}
	}
	var (
		defaultCreditNotesUpdateFields = defaultCreditNotesUpdateFields(*creditNotes)
		tempUpdateField                CreditNotesUpdateFieldList
		selectFields                   = NewCreditNotesSelectFields()
	)
	if len(creditNotesUpdateFields) > 0 {
		for _, updateField := range creditNotesUpdateFields {
			if updateField.creditNotesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultCreditNotesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeCreditNotesCompositePrimaryKeyWhere([]model.CreditNotesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsCreditNotesCommand(defaultCreditNotesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(creditNotesQueries.updateCreditNotes+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCreditNotes] error when try to update creditNotes by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateCreditNotesByFilter(ctx context.Context, filter model.Filter, creditNotesUpdateFields ...CreditNotesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(creditNotesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields CreditNotesUpdateFieldList
		selectFields = NewCreditNotesSelectFields()
	)
	for _, updateField := range creditNotesUpdateFields {
		if updateField.creditNotesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsCreditNotesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := creditNotesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeCreditNotesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"credit_notes\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCreditNotesByFilter] error when try to update creditNotes by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateCreditNotesByFilter] failed get rows affected")
	}
	return
}

var (
	creditNotesQueries = struct {
		selectCreditNotes      string
		selectCountCreditNotes string
		deleteCreditNotes      string
		updateCreditNotes      string
		insertCreditNotes      string
	}{
		selectCreditNotes:      "SELECT %s FROM \"credit_notes\"",
		selectCountCreditNotes: "SELECT COUNT(\"id\") FROM \"credit_notes\"",
		deleteCreditNotes:      "DELETE FROM \"credit_notes\"",
		updateCreditNotes:      "UPDATE \"credit_notes\" SET %s ",
		insertCreditNotes:      "INSERT INTO \"credit_notes\" %s VALUES %s",
	}
)

type CreditNotesRepository interface {
	CreateCreditNotes(ctx context.Context, creditNotes *model.CreditNotes, fieldsInsert ...CreditNotesField) error
	BulkCreateCreditNotes(ctx context.Context, creditNotesList []*model.CreditNotes, fieldsInsert ...CreditNotesField) error
	ResolveCreditNotes(ctx context.Context, selectFields ...CreditNotesField) (model.CreditNotesList, error)
	ResolveCreditNotesByID(ctx context.Context, primaryID model.CreditNotesPrimaryID, selectFields ...CreditNotesField) (model.CreditNotes, error)
	UpdateCreditNotesByID(ctx context.Context, id model.CreditNotesPrimaryID, creditNotes *model.CreditNotes, creditNotesUpdateFields ...CreditNotesUpdateField) error
	UpdateCreditNotesByFilter(ctx context.Context, filter model.Filter, creditNotesUpdateFields ...CreditNotesUpdateField) (rowsAffected int64, err error)
	BulkUpdateCreditNotes(ctx context.Context, creditNotesListMap map[model.CreditNotesPrimaryID]*model.CreditNotes, CreditNotessMapUpdateFieldsRequest map[model.CreditNotesPrimaryID]CreditNotesUpdateFieldList) (err error)
	DeleteCreditNotesByID(ctx context.Context, id model.CreditNotesPrimaryID) error
	BulkDeleteCreditNotesByIDs(ctx context.Context, ids []model.CreditNotesPrimaryID) error
	ResolveCreditNotesByFilter(ctx context.Context, filter model.Filter) (result []model.CreditNotesFilterResult, err error)
	IsExistCreditNotesByIDs(ctx context.Context, ids []model.CreditNotesPrimaryID) (exists bool, notFoundIds []model.CreditNotesPrimaryID, err error)
	IsExistCreditNotesByID(ctx context.Context, id model.CreditNotesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
