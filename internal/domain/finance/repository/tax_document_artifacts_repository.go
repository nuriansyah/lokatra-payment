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

func composeInsertFieldsAndParamsTaxDocumentArtifacts(taxDocumentArtifactsList []model.TaxDocumentArtifacts, fieldsInsert ...TaxDocumentArtifactsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewTaxDocumentArtifactsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, taxDocumentArtifacts := range taxDocumentArtifactsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, taxDocumentArtifacts.Id)
			case selectField.DocumentType():
				args = append(args, taxDocumentArtifacts.DocumentType)
			case selectField.DocumentId():
				args = append(args, taxDocumentArtifacts.DocumentId)
			case selectField.ArtifactType():
				args = append(args, taxDocumentArtifacts.ArtifactType)
			case selectField.StorageUri():
				args = append(args, taxDocumentArtifacts.StorageUri)
			case selectField.ContentHash():
				args = append(args, taxDocumentArtifacts.ContentHash)
			case selectField.IdempotencyKey():
				args = append(args, taxDocumentArtifacts.IdempotencyKey)
			case selectField.GeneratedAt():
				args = append(args, taxDocumentArtifacts.GeneratedAt)
			case selectField.Metadata():
				args = append(args, taxDocumentArtifacts.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, taxDocumentArtifacts.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, taxDocumentArtifacts.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, taxDocumentArtifacts.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, taxDocumentArtifacts.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, taxDocumentArtifacts.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, taxDocumentArtifacts.MetaDeletedBy)

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

func composeTaxDocumentArtifactsCompositePrimaryKeyWhere(primaryIDs []model.TaxDocumentArtifactsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"tax_document_artifacts\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultTaxDocumentArtifactsSelectFields() string {
	fields := NewTaxDocumentArtifactsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeTaxDocumentArtifactsSelectFields(selectFields ...TaxDocumentArtifactsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type TaxDocumentArtifactsField string
type TaxDocumentArtifactsFieldList []TaxDocumentArtifactsField

type TaxDocumentArtifactsSelectFields struct {
}

func (ss TaxDocumentArtifactsSelectFields) Id() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("id")
}

func (ss TaxDocumentArtifactsSelectFields) DocumentType() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("document_type")
}

func (ss TaxDocumentArtifactsSelectFields) DocumentId() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("document_id")
}

func (ss TaxDocumentArtifactsSelectFields) ArtifactType() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("artifact_type")
}

func (ss TaxDocumentArtifactsSelectFields) StorageUri() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("storage_uri")
}

func (ss TaxDocumentArtifactsSelectFields) ContentHash() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("content_hash")
}

func (ss TaxDocumentArtifactsSelectFields) IdempotencyKey() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("idempotency_key")
}

func (ss TaxDocumentArtifactsSelectFields) GeneratedAt() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("generated_at")
}

func (ss TaxDocumentArtifactsSelectFields) Metadata() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("metadata")
}

func (ss TaxDocumentArtifactsSelectFields) MetaCreatedAt() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("meta_created_at")
}

func (ss TaxDocumentArtifactsSelectFields) MetaCreatedBy() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("meta_created_by")
}

func (ss TaxDocumentArtifactsSelectFields) MetaUpdatedAt() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("meta_updated_at")
}

func (ss TaxDocumentArtifactsSelectFields) MetaUpdatedBy() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("meta_updated_by")
}

func (ss TaxDocumentArtifactsSelectFields) MetaDeletedAt() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("meta_deleted_at")
}

func (ss TaxDocumentArtifactsSelectFields) MetaDeletedBy() TaxDocumentArtifactsField {
	return TaxDocumentArtifactsField("meta_deleted_by")
}

func (ss TaxDocumentArtifactsSelectFields) All() TaxDocumentArtifactsFieldList {
	return []TaxDocumentArtifactsField{
		ss.Id(),
		ss.DocumentType(),
		ss.DocumentId(),
		ss.ArtifactType(),
		ss.StorageUri(),
		ss.ContentHash(),
		ss.IdempotencyKey(),
		ss.GeneratedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewTaxDocumentArtifactsSelectFields() TaxDocumentArtifactsSelectFields {
	return TaxDocumentArtifactsSelectFields{}
}

type TaxDocumentArtifactsUpdateFieldOption struct {
	useIncrement bool
}
type TaxDocumentArtifactsUpdateField struct {
	taxDocumentArtifactsField TaxDocumentArtifactsField
	opt                       TaxDocumentArtifactsUpdateFieldOption
	value                     interface{}
}
type TaxDocumentArtifactsUpdateFieldList []TaxDocumentArtifactsUpdateField

func defaultTaxDocumentArtifactsUpdateFieldOption() TaxDocumentArtifactsUpdateFieldOption {
	return TaxDocumentArtifactsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementTaxDocumentArtifactsOption(useIncrement bool) func(*TaxDocumentArtifactsUpdateFieldOption) {
	return func(pcufo *TaxDocumentArtifactsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewTaxDocumentArtifactsUpdateField(field TaxDocumentArtifactsField, val interface{}, opts ...func(*TaxDocumentArtifactsUpdateFieldOption)) TaxDocumentArtifactsUpdateField {
	defaultOpt := defaultTaxDocumentArtifactsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return TaxDocumentArtifactsUpdateField{
		taxDocumentArtifactsField: field,
		value:                     val,
		opt:                       defaultOpt,
	}
}
func defaultTaxDocumentArtifactsUpdateFields(taxDocumentArtifacts model.TaxDocumentArtifacts) (taxDocumentArtifactsUpdateFieldList TaxDocumentArtifactsUpdateFieldList) {
	selectFields := NewTaxDocumentArtifactsSelectFields()
	taxDocumentArtifactsUpdateFieldList = append(taxDocumentArtifactsUpdateFieldList,
		NewTaxDocumentArtifactsUpdateField(selectFields.Id(), taxDocumentArtifacts.Id),
		NewTaxDocumentArtifactsUpdateField(selectFields.DocumentType(), taxDocumentArtifacts.DocumentType),
		NewTaxDocumentArtifactsUpdateField(selectFields.DocumentId(), taxDocumentArtifacts.DocumentId),
		NewTaxDocumentArtifactsUpdateField(selectFields.ArtifactType(), taxDocumentArtifacts.ArtifactType),
		NewTaxDocumentArtifactsUpdateField(selectFields.StorageUri(), taxDocumentArtifacts.StorageUri),
		NewTaxDocumentArtifactsUpdateField(selectFields.ContentHash(), taxDocumentArtifacts.ContentHash),
		NewTaxDocumentArtifactsUpdateField(selectFields.IdempotencyKey(), taxDocumentArtifacts.IdempotencyKey),
		NewTaxDocumentArtifactsUpdateField(selectFields.GeneratedAt(), taxDocumentArtifacts.GeneratedAt),
		NewTaxDocumentArtifactsUpdateField(selectFields.Metadata(), taxDocumentArtifacts.Metadata),
		NewTaxDocumentArtifactsUpdateField(selectFields.MetaCreatedAt(), taxDocumentArtifacts.MetaCreatedAt),
		NewTaxDocumentArtifactsUpdateField(selectFields.MetaCreatedBy(), taxDocumentArtifacts.MetaCreatedBy),
		NewTaxDocumentArtifactsUpdateField(selectFields.MetaUpdatedAt(), taxDocumentArtifacts.MetaUpdatedAt),
		NewTaxDocumentArtifactsUpdateField(selectFields.MetaUpdatedBy(), taxDocumentArtifacts.MetaUpdatedBy),
		NewTaxDocumentArtifactsUpdateField(selectFields.MetaDeletedAt(), taxDocumentArtifacts.MetaDeletedAt),
		NewTaxDocumentArtifactsUpdateField(selectFields.MetaDeletedBy(), taxDocumentArtifacts.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsTaxDocumentArtifactsCommand(taxDocumentArtifactsUpdateFieldList TaxDocumentArtifactsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range taxDocumentArtifactsUpdateFieldList {
		field := string(updateField.taxDocumentArtifactsField)
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

func (repo *RepositoryImpl) BulkCreateTaxDocumentArtifacts(ctx context.Context, taxDocumentArtifactsList []*model.TaxDocumentArtifacts, fieldsInsert ...TaxDocumentArtifactsField) (err error) {
	var (
		fieldsStr                     string
		valueListStr                  []string
		argsList                      []interface{}
		primaryIds                    []model.TaxDocumentArtifactsPrimaryID
		taxDocumentArtifactsValueList []model.TaxDocumentArtifacts
	)

	if len(fieldsInsert) == 0 {
		selectField := NewTaxDocumentArtifactsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, taxDocumentArtifacts := range taxDocumentArtifactsList {

		primaryIds = append(primaryIds, taxDocumentArtifacts.ToTaxDocumentArtifactsPrimaryID())

		taxDocumentArtifactsValueList = append(taxDocumentArtifactsValueList, *taxDocumentArtifacts)
	}

	_, notFoundIds, err := repo.IsExistTaxDocumentArtifactsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxDocumentArtifacts] failed checking taxDocumentArtifacts whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.TaxDocumentArtifactsPrimaryID{}
		mapNotFoundIds := map[model.TaxDocumentArtifactsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "taxDocumentArtifacts", fmt.Sprintf("taxDocumentArtifacts with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsTaxDocumentArtifacts(taxDocumentArtifactsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(taxDocumentArtifactsQueries.insertTaxDocumentArtifacts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateTaxDocumentArtifacts] failed exec create taxDocumentArtifacts query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteTaxDocumentArtifactsByIDs(ctx context.Context, primaryIDs []model.TaxDocumentArtifactsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistTaxDocumentArtifactsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxDocumentArtifactsByIDs] failed checking taxDocumentArtifacts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentArtifacts with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_document_artifacts\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(taxDocumentArtifactsQueries.deleteTaxDocumentArtifacts + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxDocumentArtifactsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteTaxDocumentArtifactsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistTaxDocumentArtifactsByIDs(ctx context.Context, ids []model.TaxDocumentArtifactsPrimaryID) (exists bool, notFoundIds []model.TaxDocumentArtifactsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"tax_document_artifacts\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(taxDocumentArtifactsQueries.selectTaxDocumentArtifacts, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxDocumentArtifactsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.TaxDocumentArtifactsPrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxDocumentArtifactsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.TaxDocumentArtifactsPrimaryID]bool{}
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

// BulkUpdateTaxDocumentArtifacts is used to bulk update taxDocumentArtifacts, by default it will update all field
// if want to update specific field, then fill taxDocumentArtifactssMapUpdateFieldsRequest else please fill taxDocumentArtifactssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateTaxDocumentArtifacts(ctx context.Context, taxDocumentArtifactssMap map[model.TaxDocumentArtifactsPrimaryID]*model.TaxDocumentArtifacts, taxDocumentArtifactssMapUpdateFieldsRequest map[model.TaxDocumentArtifactsPrimaryID]TaxDocumentArtifactsUpdateFieldList) (err error) {
	if len(taxDocumentArtifactssMap) == 0 && len(taxDocumentArtifactssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		taxDocumentArtifactssMapUpdateField map[model.TaxDocumentArtifactsPrimaryID]TaxDocumentArtifactsUpdateFieldList = map[model.TaxDocumentArtifactsPrimaryID]TaxDocumentArtifactsUpdateFieldList{}
		asTableValues                       string                                                                      = "myvalues"
	)

	if len(taxDocumentArtifactssMap) > 0 {
		for id, taxDocumentArtifacts := range taxDocumentArtifactssMap {
			if taxDocumentArtifacts == nil {
				log.Error().Err(err).Msg("[BulkUpdateTaxDocumentArtifacts] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			taxDocumentArtifactssMapUpdateField[id] = defaultTaxDocumentArtifactsUpdateFields(*taxDocumentArtifacts)
		}
	} else {
		taxDocumentArtifactssMapUpdateField = taxDocumentArtifactssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateTaxDocumentArtifactsQuery(taxDocumentArtifactssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistTaxDocumentArtifactsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxDocumentArtifacts] failed checking taxDocumentArtifacts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentArtifacts with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeTaxDocumentArtifactsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"tax_document_artifacts\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateTaxDocumentArtifacts] failed exec query")
	}
	return
}

type TaxDocumentArtifactsFieldParameter struct {
	param string
	args  []interface{}
}

func NewTaxDocumentArtifactsFieldParameter(param string, args ...interface{}) TaxDocumentArtifactsFieldParameter {
	return TaxDocumentArtifactsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateTaxDocumentArtifactsQuery(mapTaxDocumentArtifactss map[model.TaxDocumentArtifactsPrimaryID]TaxDocumentArtifactsUpdateFieldList, asTableValues string) (primaryIDs []model.TaxDocumentArtifactsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.TaxDocumentArtifactsPrimaryID]map[string]interface{}{}
	taxDocumentArtifactsSelectFields := NewTaxDocumentArtifactsSelectFields()
	for id, updateFields := range mapTaxDocumentArtifactss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.taxDocumentArtifactsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapTaxDocumentArtifactss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetTaxDocumentArtifactsFieldType(updateField.taxDocumentArtifactsField)))
			args = append(args, fields[string(updateField.taxDocumentArtifactsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.taxDocumentArtifactsField))
		if updateField.taxDocumentArtifactsField == taxDocumentArtifactsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.taxDocumentArtifactsField, asTableValues, updateField.taxDocumentArtifactsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.taxDocumentArtifactsField,
				"\"tax_document_artifacts\"", updateField.taxDocumentArtifactsField,
				asTableValues, updateField.taxDocumentArtifactsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeTaxDocumentArtifactsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.TaxDocumentArtifactsPrimaryID, asTableValue string) (whereQry string) {
	taxDocumentArtifactsSelectFields := NewTaxDocumentArtifactsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"tax_document_artifacts\".\"id\" = %s.\"id\"::"+GetTaxDocumentArtifactsFieldType(taxDocumentArtifactsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetTaxDocumentArtifactsFieldType(taxDocumentArtifactsField TaxDocumentArtifactsField) string {
	selectTaxDocumentArtifactsFields := NewTaxDocumentArtifactsSelectFields()
	switch taxDocumentArtifactsField {

	case selectTaxDocumentArtifactsFields.Id():
		return "uuid"

	case selectTaxDocumentArtifactsFields.DocumentType():
		return "document_type_enum"

	case selectTaxDocumentArtifactsFields.DocumentId():
		return "uuid"

	case selectTaxDocumentArtifactsFields.ArtifactType():
		return "text"

	case selectTaxDocumentArtifactsFields.StorageUri():
		return "text"

	case selectTaxDocumentArtifactsFields.ContentHash():
		return "text"

	case selectTaxDocumentArtifactsFields.IdempotencyKey():
		return "text"

	case selectTaxDocumentArtifactsFields.GeneratedAt():
		return "timestamptz"

	case selectTaxDocumentArtifactsFields.Metadata():
		return "jsonb"

	case selectTaxDocumentArtifactsFields.MetaCreatedAt():
		return "timestamptz"

	case selectTaxDocumentArtifactsFields.MetaCreatedBy():
		return "uuid"

	case selectTaxDocumentArtifactsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectTaxDocumentArtifactsFields.MetaUpdatedBy():
		return "uuid"

	case selectTaxDocumentArtifactsFields.MetaDeletedAt():
		return "timestamptz"

	case selectTaxDocumentArtifactsFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateTaxDocumentArtifacts(ctx context.Context, taxDocumentArtifacts *model.TaxDocumentArtifacts, fieldsInsert ...TaxDocumentArtifactsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewTaxDocumentArtifactsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.TaxDocumentArtifactsPrimaryID{
		Id: taxDocumentArtifacts.Id,
	}
	exists, err := repo.IsExistTaxDocumentArtifactsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxDocumentArtifacts] failed checking taxDocumentArtifacts whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "taxDocumentArtifacts", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsTaxDocumentArtifacts([]model.TaxDocumentArtifacts{*taxDocumentArtifacts}, fieldsInsert...)
	commandQuery := fmt.Sprintf(taxDocumentArtifactsQueries.insertTaxDocumentArtifacts, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateTaxDocumentArtifacts] failed exec create taxDocumentArtifacts query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteTaxDocumentArtifactsByID(ctx context.Context, primaryID model.TaxDocumentArtifactsPrimaryID) (err error) {
	exists, err := repo.IsExistTaxDocumentArtifactsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxDocumentArtifactsByID] failed checking taxDocumentArtifacts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentArtifacts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeTaxDocumentArtifactsCompositePrimaryKeyWhere([]model.TaxDocumentArtifactsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(taxDocumentArtifactsQueries.deleteTaxDocumentArtifacts + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteTaxDocumentArtifactsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxDocumentArtifactsByFilter(ctx context.Context, filter model.Filter) (result []model.TaxDocumentArtifactsFilterResult, err error) {
	query, args, err := composeTaxDocumentArtifactsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxDocumentArtifactsByFilter] failed compose taxDocumentArtifacts filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxDocumentArtifactsByFilter] failed get taxDocumentArtifacts by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeTaxDocumentArtifactsFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.TaxDocumentArtifactsFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeTaxDocumentArtifactsFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeTaxDocumentArtifactsSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeTaxDocumentArtifactsFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewTaxDocumentArtifactsFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 15+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["document_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"document_type\"")
			selectedColumns["document_type"] = struct{}{}
		}
		if _, selected := selectedColumns["document_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"document_id\"")
			selectedColumns["document_id"] = struct{}{}
		}
		if _, selected := selectedColumns["artifact_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"artifact_type\"")
			selectedColumns["artifact_type"] = struct{}{}
		}
		if _, selected := selectedColumns["storage_uri"]; !selected {
			selectColumns = append(selectColumns, "base.\"storage_uri\"")
			selectedColumns["storage_uri"] = struct{}{}
		}
		if _, selected := selectedColumns["content_hash"]; !selected {
			selectColumns = append(selectColumns, "base.\"content_hash\"")
			selectedColumns["content_hash"] = struct{}{}
		}
		if _, selected := selectedColumns["idempotency_key"]; !selected {
			selectColumns = append(selectColumns, "base.\"idempotency_key\"")
			selectedColumns["idempotency_key"] = struct{}{}
		}
		if _, selected := selectedColumns["generated_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"generated_at\"")
			selectedColumns["generated_at"] = struct{}{}
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

type taxDocumentArtifactsFilterPlaceholder struct {
	index int
}

func (p *taxDocumentArtifactsFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeTaxDocumentArtifactsFilterPredicate(filterField model.FilterField, placeholders *taxDocumentArtifactsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewTaxDocumentArtifactsFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeTaxDocumentArtifactsFilterSQLExpr(spec)
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

func composeTaxDocumentArtifactsFilterGroup(group model.FilterGroup, placeholders *taxDocumentArtifactsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeTaxDocumentArtifactsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeTaxDocumentArtifactsFilterGroup(child, placeholders, args, requiredJoins)
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

func composeTaxDocumentArtifactsFilterWhereQueries(filter model.Filter, placeholders *taxDocumentArtifactsFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeTaxDocumentArtifactsFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeTaxDocumentArtifactsFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeTaxDocumentArtifactsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateTaxDocumentArtifactsFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeTaxDocumentArtifactsSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeTaxDocumentArtifactsFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := taxDocumentArtifactsFilterPlaceholder{index: 1}
	whereQueries, err := composeTaxDocumentArtifactsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewTaxDocumentArtifactsFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeTaxDocumentArtifactsFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeTaxDocumentArtifactsSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"tax_document_artifacts\" base%s", strings.Join(selectColumns, ","), composeTaxDocumentArtifactsFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistTaxDocumentArtifactsByID(ctx context.Context, primaryID model.TaxDocumentArtifactsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeTaxDocumentArtifactsCompositePrimaryKeyWhere([]model.TaxDocumentArtifactsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", taxDocumentArtifactsQueries.selectCountTaxDocumentArtifacts, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistTaxDocumentArtifactsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxDocumentArtifacts(ctx context.Context, selectFields ...TaxDocumentArtifactsField) (taxDocumentArtifactsList model.TaxDocumentArtifactsList, err error) {
	var (
		defaultTaxDocumentArtifactsSelectFields = defaultTaxDocumentArtifactsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxDocumentArtifactsSelectFields = composeTaxDocumentArtifactsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(taxDocumentArtifactsQueries.selectTaxDocumentArtifacts, defaultTaxDocumentArtifactsSelectFields)

	err = repo.db.Read.SelectContext(ctx, &taxDocumentArtifactsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveTaxDocumentArtifacts] failed get taxDocumentArtifacts list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveTaxDocumentArtifactsByID(ctx context.Context, primaryID model.TaxDocumentArtifactsPrimaryID, selectFields ...TaxDocumentArtifactsField) (taxDocumentArtifacts model.TaxDocumentArtifacts, err error) {
	var (
		defaultTaxDocumentArtifactsSelectFields = defaultTaxDocumentArtifactsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultTaxDocumentArtifactsSelectFields = composeTaxDocumentArtifactsSelectFields(selectFields...)
	}
	whereQry, params := composeTaxDocumentArtifactsCompositePrimaryKeyWhere([]model.TaxDocumentArtifactsPrimaryID{primaryID})
	query := fmt.Sprintf(taxDocumentArtifactsQueries.selectTaxDocumentArtifacts+" WHERE "+whereQry, defaultTaxDocumentArtifactsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &taxDocumentArtifacts, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("taxDocumentArtifacts with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveTaxDocumentArtifactsByID] failed get taxDocumentArtifacts")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateTaxDocumentArtifactsByID(ctx context.Context, primaryID model.TaxDocumentArtifactsPrimaryID, taxDocumentArtifacts *model.TaxDocumentArtifacts, taxDocumentArtifactsUpdateFields ...TaxDocumentArtifactsUpdateField) (err error) {
	exists, err := repo.IsExistTaxDocumentArtifactsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentArtifacts] failed checking taxDocumentArtifacts whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("taxDocumentArtifacts with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if taxDocumentArtifacts == nil {
		if len(taxDocumentArtifactsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateTaxDocumentArtifactsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		taxDocumentArtifacts = &model.TaxDocumentArtifacts{}
	}
	var (
		defaultTaxDocumentArtifactsUpdateFields = defaultTaxDocumentArtifactsUpdateFields(*taxDocumentArtifacts)
		tempUpdateField                         TaxDocumentArtifactsUpdateFieldList
		selectFields                            = NewTaxDocumentArtifactsSelectFields()
	)
	if len(taxDocumentArtifactsUpdateFields) > 0 {
		for _, updateField := range taxDocumentArtifactsUpdateFields {
			if updateField.taxDocumentArtifactsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultTaxDocumentArtifactsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeTaxDocumentArtifactsCompositePrimaryKeyWhere([]model.TaxDocumentArtifactsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsTaxDocumentArtifactsCommand(defaultTaxDocumentArtifactsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(taxDocumentArtifactsQueries.updateTaxDocumentArtifacts+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentArtifacts] error when try to update taxDocumentArtifacts by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateTaxDocumentArtifactsByFilter(ctx context.Context, filter model.Filter, taxDocumentArtifactsUpdateFields ...TaxDocumentArtifactsUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(taxDocumentArtifactsUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields TaxDocumentArtifactsUpdateFieldList
		selectFields = NewTaxDocumentArtifactsSelectFields()
	)
	for _, updateField := range taxDocumentArtifactsUpdateFields {
		if updateField.taxDocumentArtifactsField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsTaxDocumentArtifactsCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := taxDocumentArtifactsFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeTaxDocumentArtifactsFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"tax_document_artifacts\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentArtifactsByFilter] error when try to update taxDocumentArtifacts by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateTaxDocumentArtifactsByFilter] failed get rows affected")
	}
	return
}

var (
	taxDocumentArtifactsQueries = struct {
		selectTaxDocumentArtifacts      string
		selectCountTaxDocumentArtifacts string
		deleteTaxDocumentArtifacts      string
		updateTaxDocumentArtifacts      string
		insertTaxDocumentArtifacts      string
	}{
		selectTaxDocumentArtifacts:      "SELECT %s FROM \"tax_document_artifacts\"",
		selectCountTaxDocumentArtifacts: "SELECT COUNT(\"id\") FROM \"tax_document_artifacts\"",
		deleteTaxDocumentArtifacts:      "DELETE FROM \"tax_document_artifacts\"",
		updateTaxDocumentArtifacts:      "UPDATE \"tax_document_artifacts\" SET %s ",
		insertTaxDocumentArtifacts:      "INSERT INTO \"tax_document_artifacts\" %s VALUES %s",
	}
)

type TaxDocumentArtifactsRepository interface {
	CreateTaxDocumentArtifacts(ctx context.Context, taxDocumentArtifacts *model.TaxDocumentArtifacts, fieldsInsert ...TaxDocumentArtifactsField) error
	BulkCreateTaxDocumentArtifacts(ctx context.Context, taxDocumentArtifactsList []*model.TaxDocumentArtifacts, fieldsInsert ...TaxDocumentArtifactsField) error
	ResolveTaxDocumentArtifacts(ctx context.Context, selectFields ...TaxDocumentArtifactsField) (model.TaxDocumentArtifactsList, error)
	ResolveTaxDocumentArtifactsByID(ctx context.Context, primaryID model.TaxDocumentArtifactsPrimaryID, selectFields ...TaxDocumentArtifactsField) (model.TaxDocumentArtifacts, error)
	UpdateTaxDocumentArtifactsByID(ctx context.Context, id model.TaxDocumentArtifactsPrimaryID, taxDocumentArtifacts *model.TaxDocumentArtifacts, taxDocumentArtifactsUpdateFields ...TaxDocumentArtifactsUpdateField) error
	UpdateTaxDocumentArtifactsByFilter(ctx context.Context, filter model.Filter, taxDocumentArtifactsUpdateFields ...TaxDocumentArtifactsUpdateField) (rowsAffected int64, err error)
	BulkUpdateTaxDocumentArtifacts(ctx context.Context, taxDocumentArtifactsListMap map[model.TaxDocumentArtifactsPrimaryID]*model.TaxDocumentArtifacts, TaxDocumentArtifactssMapUpdateFieldsRequest map[model.TaxDocumentArtifactsPrimaryID]TaxDocumentArtifactsUpdateFieldList) (err error)
	DeleteTaxDocumentArtifactsByID(ctx context.Context, id model.TaxDocumentArtifactsPrimaryID) error
	BulkDeleteTaxDocumentArtifactsByIDs(ctx context.Context, ids []model.TaxDocumentArtifactsPrimaryID) error
	ResolveTaxDocumentArtifactsByFilter(ctx context.Context, filter model.Filter) (result []model.TaxDocumentArtifactsFilterResult, err error)
	IsExistTaxDocumentArtifactsByIDs(ctx context.Context, ids []model.TaxDocumentArtifactsPrimaryID) (exists bool, notFoundIds []model.TaxDocumentArtifactsPrimaryID, err error)
	IsExistTaxDocumentArtifactsByID(ctx context.Context, id model.TaxDocumentArtifactsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
