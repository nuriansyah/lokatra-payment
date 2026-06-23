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

func composeInsertFieldsAndParamsTaxInvoices(taxInvoicesList []model.TaxInvoices, fieldsInsert ...TaxInvoicesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewTaxInvoicesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, taxInvoices := range taxInvoicesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, taxInvoices.Id)
			case selectField.InvoiceNo():
				args = append(args, taxInvoices.InvoiceNo)
			case selectField.SourceType():
				args = append(args, taxInvoices.SourceType)
			case selectField.SourceId():
				args = append(args, taxInvoices.SourceId)
			case selectField.MerchantPartyId():
				args = append(args, taxInvoices.MerchantPartyId)
			case selectField.CustomerPartyId():
				args = append(args, taxInvoices.CustomerPartyId)
			case selectField.CurrencyCode():
				args = append(args, taxInvoices.CurrencyCode)
			case selectField.TaxableAmount():
				args = append(args, taxInvoices.TaxableAmount)
			case selectField.TaxAmount():
				args = append(args, taxInvoices.TaxAmount)
			case selectField.TotalAmount():
				args = append(args, taxInvoices.TotalAmount)
			case selectField.InvoiceStatus():
				args = append(args, taxInvoices.InvoiceStatus)
			case selectField.IssuedAt():
				args = append(args, taxInvoices.IssuedAt)
			case selectField.Metadata():
				args = append(args, taxInvoices.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, taxInvoices.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, taxInvoices.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, taxInvoices.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, taxInvoices.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, taxInvoices.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, taxInvoices.MetaDeletedBy)

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

func composeTaxInvoicesCompositePrimaryKeyWhere(primaryIDs []model.TaxInvoicesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"tax_invoices\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultTaxInvoicesSelectFields() string {
	fields := NewTaxInvoicesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeTaxInvoicesSelectFields(selectFields ...TaxInvoicesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type TaxInvoicesField string
type TaxInvoicesFieldList []TaxInvoicesField

type TaxInvoicesSelectFields struct {
}

func (ss TaxInvoicesSelectFields) Id() TaxInvoicesField {
	return TaxInvoicesField("id")
}

func (ss TaxInvoicesSelectFields) InvoiceNo() TaxInvoicesField {
	return TaxInvoicesField("invoice_no")
}

func (ss TaxInvoicesSelectFields) SourceType() TaxInvoicesField {
	return TaxInvoicesField("source_type")
}

func (ss TaxInvoicesSelectFields) SourceId() TaxInvoicesField {
	return TaxInvoicesField("source_id")
}

func (ss TaxInvoicesSelectFields) MerchantPartyId() TaxInvoicesField {
	return TaxInvoicesField("merchant_party_id")
}

func (ss TaxInvoicesSelectFields) CustomerPartyId() TaxInvoicesField {
	return TaxInvoicesField("customer_party_id")
}

func (ss TaxInvoicesSelectFields) CurrencyCode() TaxInvoicesField {
	return TaxInvoicesField("currency_code")
}

func (ss TaxInvoicesSelectFields) TaxableAmount() TaxInvoicesField {
	return TaxInvoicesField("taxable_amount")
}

func (ss TaxInvoicesSelectFields) TaxAmount() TaxInvoicesField {
	return TaxInvoicesField("tax_amount")
}

func (ss TaxInvoicesSelectFields) TotalAmount() TaxInvoicesField {
	return TaxInvoicesField("total_amount")
}

func (ss TaxInvoicesSelectFields) InvoiceStatus() TaxInvoicesField {
	return TaxInvoicesField("invoice_status")
}

func (ss TaxInvoicesSelectFields) IssuedAt() TaxInvoicesField {
	return TaxInvoicesField("issued_at")
}

func (ss TaxInvoicesSelectFields) Metadata() TaxInvoicesField {
	return TaxInvoicesField("metadata")
}

func (ss TaxInvoicesSelectFields) MetaCreatedAt() TaxInvoicesField {
	return TaxInvoicesField("meta_created_at")
}

func (ss TaxInvoicesSelectFields) MetaCreatedBy() TaxInvoicesField {
	return TaxInvoicesField("meta_created_by")
}

func (ss TaxInvoicesSelectFields) MetaUpdatedAt() TaxInvoicesField {
	return TaxInvoicesField("meta_updated_at")
}

func (ss TaxInvoicesSelectFields) MetaUpdatedBy() TaxInvoicesField {
	return TaxInvoicesField("meta_updated_by")
}

func (ss TaxInvoicesSelectFields) MetaDeletedAt() TaxInvoicesField {
	return TaxInvoicesField("meta_deleted_at")
}

func (ss TaxInvoicesSelectFields) MetaDeletedBy() TaxInvoicesField {
	return TaxInvoicesField("meta_deleted_by")
}

func (ss TaxInvoicesSelectFields) All() TaxInvoicesFieldList {
	return []TaxInvoicesField{
		ss.Id(),
		ss.InvoiceNo(),
		ss.SourceType(),
		ss.SourceId(),
		ss.MerchantPartyId(),
		ss.CustomerPartyId(),
		ss.CurrencyCode(),
		ss.TaxableAmount(),
		ss.TaxAmount(),
		ss.TotalAmount(),
		ss.InvoiceStatus(),
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

func NewTaxInvoicesSelectFields() TaxInvoicesSelectFields {
	return TaxInvoicesSelectFields{}
}

type TaxInvoicesUpdateFieldOption struct {
	useIncrement bool
}
type TaxInvoicesUpdateField struct {
	taxInvoicesField TaxInvoicesField
	opt              TaxInvoicesUpdateFieldOption
	value            interface{}
}
type TaxInvoicesUpdateFieldList []TaxInvoicesUpdateField

func defaultTaxInvoicesUpdateFieldOption() TaxInvoicesUpdateFieldOption {
	return TaxInvoicesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementTaxInvoicesOption(useIncrement bool) func(*TaxInvoicesUpdateFieldOption) {
	return func(pcufo *TaxInvoicesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewTaxInvoicesUpdateField(field TaxInvoicesField, val interface{}, opts ...func(*TaxInvoicesUpdateFieldOption)) TaxInvoicesUpdateField {
	defaultOpt := defaultTaxInvoicesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return TaxInvoicesUpdateField{
		taxInvoicesField: field,
		value:            val,
		opt:              defaultOpt,
	}
}
func defaultTaxInvoicesUpdateFields(taxInvoices model.TaxInvoices) (taxInvoicesUpdateFieldList TaxInvoicesUpdateFieldList) {
	selectFields := NewTaxInvoicesSelectFields()
	taxInvoicesUpdateFieldList = append(taxInvoicesUpdateFieldList,
		NewTaxInvoicesUpdateField(selectFields.Id(), taxInvoices.Id),
		NewTaxInvoicesUpdateField(selectFields.InvoiceNo(), taxInvoices.InvoiceNo),
		NewTaxInvoicesUpdateField(selectFields.SourceType(), taxInvoices.SourceType),
		NewTaxInvoicesUpdateField(selectFields.SourceId(), taxInvoices.SourceId),
		NewTaxInvoicesUpdateField(selectFields.MerchantPartyId(), taxInvoices.MerchantPartyId),
		NewTaxInvoicesUpdateField(selectFields.CustomerPartyId(), taxInvoices.CustomerPartyId),
		NewTaxInvoicesUpdateField(selectFields.CurrencyCode(), taxInvoices.CurrencyCode),
		NewTaxInvoicesUpdateField(selectFields.TaxableAmount(), taxInvoices.TaxableAmount),
		NewTaxInvoicesUpdateField(selectFields.TaxAmount(), taxInvoices.TaxAmount),
		NewTaxInvoicesUpdateField(selectFields.TotalAmount(), taxInvoices.TotalAmount),
		NewTaxInvoicesUpdateField(selectFields.InvoiceStatus(), taxInvoices.InvoiceStatus),
		NewTaxInvoicesUpdateField(selectFields.IssuedAt(), taxInvoices.IssuedAt),
		NewTaxInvoicesUpdateField(selectFields.Metadata(), taxInvoices.Metadata),
		NewTaxInvoicesUpdateField(selectFields.MetaCreatedAt(), taxInvoices.MetaCreatedAt),
		NewTaxInvoicesUpdateField(selectFields.MetaCreatedBy(), taxInvoices.MetaCreatedBy),
		NewTaxInvoicesUpdateField(selectFields.MetaUpdatedAt(), taxInvoices.MetaUpdatedAt),
		NewTaxInvoicesUpdateField(selectFields.MetaUpdatedBy(), taxInvoices.MetaUpdatedBy),
		NewTaxInvoicesUpdateField(selectFields.MetaDeletedAt(), taxInvoices.MetaDeletedAt),
		NewTaxInvoicesUpdateField(selectFields.MetaDeletedBy(), taxInvoices.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsTaxInvoicesCommand(taxInvoicesUpdateFieldList TaxInvoicesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range taxInvoicesUpdateFieldList {
		field := string(updateField.taxInvoicesField)
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

func (repo *RepositoryImpl) BulkCreateTaxInvoices(ctx context.Context, taxInvoicesList []*model.TaxInvoices, fieldsInsert ...TaxInvoicesField) (err error) {
	var (
		fieldsStr            string
		valueListStr         []string
		argsList             []interface{}
		primaryIds           []model.TaxInvoicesPrimaryID
		taxInvoicesValueList []model.TaxInvoices
	)

	if len(fieldsInsert) == 0 {
		selectField := NewTaxInvoicesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, taxInvoices := range taxInvoicesList {

		primaryIds = append(primaryIds, taxInvoices.ToTaxInvoicesPrimaryID())

		taxInvoicesValueList = append(taxInvoicesValueList, *taxInvoices)
	}

	_, notFoundIds, err := repo.IsExistTaxInvoicesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxInvoices] failed checking taxInvoices whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.TaxInvoicesPrimaryID{}
		mapNotFoundIds := map[model.TaxInvoicesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "taxInvoices", fmt.Sprintf("taxInvoices with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsTaxInvoices(taxInvoicesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(taxInvoicesQueries.insertTaxInvoices, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxInvoices] failed exec create taxInvoices query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteTaxInvoicesByIDs(ctx context.Context, primaryIDs []model.TaxInvoicesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistTaxInvoicesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxInvoicesByIDs] failed checking taxInvoices whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxInvoices with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_invoices\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(taxInvoicesQueries.deleteTaxInvoices + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxInvoicesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxInvoicesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistTaxInvoicesByIDs(ctx context.Context, ids []model.TaxInvoicesPrimaryID) (exists bool, notFoundIds []model.TaxInvoicesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_invoices\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(taxInvoicesQueries.selectTaxInvoices, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxInvoicesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.TaxInvoicesPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxInvoicesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.TaxInvoicesPrimaryID]bool{}
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

// BulkUpdateTaxInvoices is used to bulk update taxInvoices, by default it will update all field
// if want to update specific field, then fill taxInvoicessMapUpdateFieldsRequest else please fill taxInvoicessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateTaxInvoices(ctx context.Context, taxInvoicessMap map[model.TaxInvoicesPrimaryID]*model.TaxInvoices, taxInvoicessMapUpdateFieldsRequest map[model.TaxInvoicesPrimaryID]TaxInvoicesUpdateFieldList) (err error) {
	if len(taxInvoicessMap) == 0 && len(taxInvoicessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		taxInvoicessMapUpdateField map[model.TaxInvoicesPrimaryID]TaxInvoicesUpdateFieldList = map[model.TaxInvoicesPrimaryID]TaxInvoicesUpdateFieldList{}
		asTableValues              string                                                    = "myvalues"
	)

	if len(taxInvoicessMap) > 0 {
		for id, taxInvoices := range taxInvoicessMap {
			if taxInvoices == nil {
				log.Error().Err(err).Msg("[BulkUpdateTaxInvoices] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			taxInvoicessMapUpdateField[id] = defaultTaxInvoicesUpdateFields(*taxInvoices)
		}
	} else {
		taxInvoicessMapUpdateField = taxInvoicessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateTaxInvoicesQuery(taxInvoicessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistTaxInvoicesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxInvoices] failed checking taxInvoices whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxInvoices with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeTaxInvoicesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"tax_invoices\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxInvoices] failed exec query")
	}
	return
}

type TaxInvoicesFieldParameter struct {
	param string
	args  []interface{}
}

func NewTaxInvoicesFieldParameter(param string, args ...interface{}) TaxInvoicesFieldParameter {
	return TaxInvoicesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateTaxInvoicesQuery(mapTaxInvoicess map[model.TaxInvoicesPrimaryID]TaxInvoicesUpdateFieldList, asTableValues string) (primaryIDs []model.TaxInvoicesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.TaxInvoicesPrimaryID]map[string]interface{}{}
	taxInvoicesSelectFields := NewTaxInvoicesSelectFields()
	for id, updateFields := range mapTaxInvoicess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.taxInvoicesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapTaxInvoicess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetTaxInvoicesFieldType(updateField.taxInvoicesField)))
			args = append(args, fields[string(updateField.taxInvoicesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.taxInvoicesField))
		if updateField.taxInvoicesField == taxInvoicesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.taxInvoicesField, asTableValues, updateField.taxInvoicesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.taxInvoicesField,
				"\"tax_invoices\"", updateField.taxInvoicesField,
				asTableValues, updateField.taxInvoicesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeTaxInvoicesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.TaxInvoicesPrimaryID, asTableValue string) (whereQry string) {
	taxInvoicesSelectFields := NewTaxInvoicesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"tax_invoices\".\"id\" = %s.\"id\"::"+GetTaxInvoicesFieldType(taxInvoicesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetTaxInvoicesFieldType(taxInvoicesField TaxInvoicesField) string {
	selectTaxInvoicesFields := NewTaxInvoicesSelectFields()
	switch taxInvoicesField {

	case selectTaxInvoicesFields.Id():
		return "uuid"

	case selectTaxInvoicesFields.InvoiceNo():
		return "text"

	case selectTaxInvoicesFields.SourceType():
		return "text"

	case selectTaxInvoicesFields.SourceId():
		return "uuid"

	case selectTaxInvoicesFields.MerchantPartyId():
		return "uuid"

	case selectTaxInvoicesFields.CustomerPartyId():
		return "uuid"

	case selectTaxInvoicesFields.CurrencyCode():
		return "text"

	case selectTaxInvoicesFields.TaxableAmount():
		return "numeric"

	case selectTaxInvoicesFields.TaxAmount():
		return "numeric"

	case selectTaxInvoicesFields.TotalAmount():
		return "numeric"

	case selectTaxInvoicesFields.InvoiceStatus():
		return "invoice_status_enum"

	case selectTaxInvoicesFields.IssuedAt():
		return "timestamptz"

	case selectTaxInvoicesFields.Metadata():
		return "jsonb"

	case selectTaxInvoicesFields.MetaCreatedAt():
		return "timestamptz"

	case selectTaxInvoicesFields.MetaCreatedBy():
		return "uuid"

	case selectTaxInvoicesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectTaxInvoicesFields.MetaUpdatedBy():
		return "uuid"

	case selectTaxInvoicesFields.MetaDeletedAt():
		return "timestamptz"

	case selectTaxInvoicesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateTaxInvoices(ctx context.Context, taxInvoices *model.TaxInvoices, fieldsInsert ...TaxInvoicesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewTaxInvoicesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.TaxInvoicesPrimaryID{
		Id: taxInvoices.Id,
	}
	exists, err := repo.IsExistTaxInvoicesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxInvoices] failed checking taxInvoices whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "taxInvoices", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsTaxInvoices([]model.TaxInvoices{*taxInvoices}, fieldsInsert...)
	commandQuery := fmt.Sprintf(taxInvoicesQueries.insertTaxInvoices, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxInvoices] failed exec create taxInvoices query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteTaxInvoicesByID(ctx context.Context, primaryID model.TaxInvoicesPrimaryID) (err error) {
	exists, err := repo.IsExistTaxInvoicesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxInvoicesByID] failed checking taxInvoices whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxInvoices with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeTaxInvoicesCompositePrimaryKeyWhere([]model.TaxInvoicesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(taxInvoicesQueries.deleteTaxInvoices + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxInvoicesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxInvoicesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxInvoicesFilterResult, err error) {
	query, args, err := composeTaxInvoicesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxInvoicesByFilter] failed compose taxInvoices filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxInvoicesByFilter] failed get taxInvoices by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeTaxInvoicesFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.TaxInvoicesFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeTaxInvoicesFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeTaxInvoicesSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeTaxInvoicesFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 19 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewTaxInvoicesFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["invoice_no"]; !selected {
			selectColumns = append(selectColumns, "base.\"invoice_no\"")
			selectedColumns["invoice_no"] = struct{}{}
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
		if _, selected := selectedColumns["invoice_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"invoice_status\"")
			selectedColumns["invoice_status"] = struct{}{}
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

type taxInvoicesFilterPlaceholder struct {
	index int
}

func (p *taxInvoicesFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeTaxInvoicesFilterPredicate(filterField model.FilterField, placeholders *taxInvoicesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewTaxInvoicesFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeTaxInvoicesFilterSQLExpr(spec)
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

func composeTaxInvoicesFilterGroup(group model.FilterGroup, placeholders *taxInvoicesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeTaxInvoicesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeTaxInvoicesFilterGroup(child, placeholders, args, requiredJoins)
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

func composeTaxInvoicesFilterWhereQueries(filter model.Filter, placeholders *taxInvoicesFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeTaxInvoicesFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeTaxInvoicesFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeTaxInvoicesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateTaxInvoicesFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeTaxInvoicesSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeTaxInvoicesFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := taxInvoicesFilterPlaceholder{index: 1}
	whereQueries, err := composeTaxInvoicesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewTaxInvoicesFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeTaxInvoicesFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeTaxInvoicesSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"tax_invoices\" base%s", strings.Join(selectColumns, ","), composeTaxInvoicesFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistTaxInvoicesByID(ctx context.Context, primaryID model.TaxInvoicesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeTaxInvoicesCompositePrimaryKeyWhere([]model.TaxInvoicesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", taxInvoicesQueries.selectCountTaxInvoices, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxInvoicesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxInvoices(ctx context.Context, selectFields ...TaxInvoicesField) (taxInvoicesList model.TaxInvoicesList, err error) {
	var (
		defaultTaxInvoicesSelectFields = defaultTaxInvoicesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxInvoicesSelectFields = composeTaxInvoicesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(taxInvoicesQueries.selectTaxInvoices, defaultTaxInvoicesSelectFields)

	err = repo.db.Read.SelectContext(ctx, &taxInvoicesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxInvoices] failed get taxInvoices list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxInvoicesByID(ctx context.Context, primaryID model.TaxInvoicesPrimaryID, selectFields ...TaxInvoicesField) (taxInvoices model.TaxInvoices, err error) {
	var (
		defaultTaxInvoicesSelectFields = defaultTaxInvoicesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxInvoicesSelectFields = composeTaxInvoicesSelectFields(selectFields...)
	}
	whereQry, params := composeTaxInvoicesCompositePrimaryKeyWhere([]model.TaxInvoicesPrimaryID{primaryID})
	query := fmt.Sprintf(taxInvoicesQueries.selectTaxInvoices+" WHERE "+whereQry, defaultTaxInvoicesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &taxInvoices, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("taxInvoices with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveTaxInvoicesByID] failed get taxInvoices")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateTaxInvoicesByID(ctx context.Context, primaryID model.TaxInvoicesPrimaryID, taxInvoices *model.TaxInvoices, taxInvoicesUpdateFields ...TaxInvoicesUpdateField) (err error) {
	exists, err := repo.IsExistTaxInvoicesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxInvoices] failed checking taxInvoices whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxInvoices with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if taxInvoices == nil {
		if len(taxInvoicesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateTaxInvoicesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		taxInvoices = &model.TaxInvoices{}
	}
	var (
		defaultTaxInvoicesUpdateFields = defaultTaxInvoicesUpdateFields(*taxInvoices)
		tempUpdateField                TaxInvoicesUpdateFieldList
		selectFields                   = NewTaxInvoicesSelectFields()
	)
	if len(taxInvoicesUpdateFields) > 0 {
		for _, updateField := range taxInvoicesUpdateFields {
			if updateField.taxInvoicesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultTaxInvoicesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeTaxInvoicesCompositePrimaryKeyWhere([]model.TaxInvoicesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsTaxInvoicesCommand(defaultTaxInvoicesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(taxInvoicesQueries.updateTaxInvoices+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxInvoices] error when try to update taxInvoices by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateTaxInvoicesByFilter(ctx context.Context, filter model.Filter, taxInvoicesUpdateFields ...TaxInvoicesUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(taxInvoicesUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields TaxInvoicesUpdateFieldList
		selectFields = NewTaxInvoicesSelectFields()
	)
	for _, updateField := range taxInvoicesUpdateFields {
		if updateField.taxInvoicesField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsTaxInvoicesCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := taxInvoicesFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeTaxInvoicesFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"tax_invoices\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxInvoicesByFilter] error when try to update taxInvoices by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxInvoicesByFilter] failed get rows affected")
	}
	return
}

var (
	taxInvoicesQueries = struct {
		selectTaxInvoices      string
		selectCountTaxInvoices string
		deleteTaxInvoices      string
		updateTaxInvoices      string
		insertTaxInvoices      string
	}{
		selectTaxInvoices:      "SELECT %s FROM \"tax_invoices\"",
		selectCountTaxInvoices: "SELECT COUNT(\"id\") FROM \"tax_invoices\"",
		deleteTaxInvoices:      "DELETE FROM \"tax_invoices\"",
		updateTaxInvoices:      "UPDATE \"tax_invoices\" SET %s ",
		insertTaxInvoices:      "INSERT INTO \"tax_invoices\" %s VALUES %s",
	}
)

type TaxInvoicesRepository interface {
	CreateTaxInvoices(ctx context.Context, taxInvoices *model.TaxInvoices, fieldsInsert ...TaxInvoicesField) error
	BulkCreateTaxInvoices(ctx context.Context, taxInvoicesList []*model.TaxInvoices, fieldsInsert ...TaxInvoicesField) error
	ResolveTaxInvoices(ctx context.Context, selectFields ...TaxInvoicesField) (model.TaxInvoicesList, error)
	ResolveTaxInvoicesByID(ctx context.Context, primaryID model.TaxInvoicesPrimaryID, selectFields ...TaxInvoicesField) (model.TaxInvoices, error)
	UpdateTaxInvoicesByID(ctx context.Context, id model.TaxInvoicesPrimaryID, taxInvoices *model.TaxInvoices, taxInvoicesUpdateFields ...TaxInvoicesUpdateField) error
	UpdateTaxInvoicesByFilter(ctx context.Context, filter model.Filter, taxInvoicesUpdateFields ...TaxInvoicesUpdateField) (rowsAffected int64, err error)
	BulkUpdateTaxInvoices(ctx context.Context, taxInvoicesListMap map[model.TaxInvoicesPrimaryID]*model.TaxInvoices, TaxInvoicessMapUpdateFieldsRequest map[model.TaxInvoicesPrimaryID]TaxInvoicesUpdateFieldList) (err error)
	DeleteTaxInvoicesByID(ctx context.Context, id model.TaxInvoicesPrimaryID) error
	BulkDeleteTaxInvoicesByIDs(ctx context.Context, ids []model.TaxInvoicesPrimaryID) error
	ResolveTaxInvoicesByFilter(ctx context.Context, filter model.Filter) (result []model.TaxInvoicesFilterResult, err error)
	IsExistTaxInvoicesByIDs(ctx context.Context, ids []model.TaxInvoicesPrimaryID) (exists bool, notFoundIds []model.TaxInvoicesPrimaryID, err error)
	IsExistTaxInvoicesByID(ctx context.Context, id model.TaxInvoicesPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
