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

func composeInsertFieldsAndParamsDisputeEvidence(disputeEvidenceList []model.DisputeEvidence, fieldsInsert ...DisputeEvidenceField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewDisputeEvidenceSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, disputeEvidence := range disputeEvidenceList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, disputeEvidence.Id)
			case selectField.DisputeId():
				args = append(args, disputeEvidence.DisputeId)
			case selectField.EvidenceType():
				args = append(args, disputeEvidence.EvidenceType)
			case selectField.StorageUri():
				args = append(args, disputeEvidence.StorageUri)
			case selectField.ContentHash():
				args = append(args, disputeEvidence.ContentHash)
			case selectField.IdempotencyKey():
				args = append(args, disputeEvidence.IdempotencyKey)
			case selectField.SubmittedBy():
				args = append(args, disputeEvidence.SubmittedBy)
			case selectField.SubmittedAt():
				args = append(args, disputeEvidence.SubmittedAt)
			case selectField.Metadata():
				args = append(args, disputeEvidence.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, disputeEvidence.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, disputeEvidence.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, disputeEvidence.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, disputeEvidence.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, disputeEvidence.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, disputeEvidence.MetaDeletedBy)

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

func composeDisputeEvidenceCompositePrimaryKeyWhere(primaryIDs []model.DisputeEvidencePrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"dispute_evidence\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultDisputeEvidenceSelectFields() string {
	fields := NewDisputeEvidenceSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeDisputeEvidenceSelectFields(selectFields ...DisputeEvidenceField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type DisputeEvidenceField string
type DisputeEvidenceFieldList []DisputeEvidenceField

type DisputeEvidenceSelectFields struct {
}

func (ss DisputeEvidenceSelectFields) Id() DisputeEvidenceField {
	return DisputeEvidenceField("id")
}

func (ss DisputeEvidenceSelectFields) DisputeId() DisputeEvidenceField {
	return DisputeEvidenceField("dispute_id")
}

func (ss DisputeEvidenceSelectFields) EvidenceType() DisputeEvidenceField {
	return DisputeEvidenceField("evidence_type")
}

func (ss DisputeEvidenceSelectFields) StorageUri() DisputeEvidenceField {
	return DisputeEvidenceField("storage_uri")
}

func (ss DisputeEvidenceSelectFields) ContentHash() DisputeEvidenceField {
	return DisputeEvidenceField("content_hash")
}

func (ss DisputeEvidenceSelectFields) IdempotencyKey() DisputeEvidenceField {
	return DisputeEvidenceField("idempotency_key")
}

func (ss DisputeEvidenceSelectFields) SubmittedBy() DisputeEvidenceField {
	return DisputeEvidenceField("submitted_by")
}

func (ss DisputeEvidenceSelectFields) SubmittedAt() DisputeEvidenceField {
	return DisputeEvidenceField("submitted_at")
}

func (ss DisputeEvidenceSelectFields) Metadata() DisputeEvidenceField {
	return DisputeEvidenceField("metadata")
}

func (ss DisputeEvidenceSelectFields) MetaCreatedAt() DisputeEvidenceField {
	return DisputeEvidenceField("meta_created_at")
}

func (ss DisputeEvidenceSelectFields) MetaCreatedBy() DisputeEvidenceField {
	return DisputeEvidenceField("meta_created_by")
}

func (ss DisputeEvidenceSelectFields) MetaUpdatedAt() DisputeEvidenceField {
	return DisputeEvidenceField("meta_updated_at")
}

func (ss DisputeEvidenceSelectFields) MetaUpdatedBy() DisputeEvidenceField {
	return DisputeEvidenceField("meta_updated_by")
}

func (ss DisputeEvidenceSelectFields) MetaDeletedAt() DisputeEvidenceField {
	return DisputeEvidenceField("meta_deleted_at")
}

func (ss DisputeEvidenceSelectFields) MetaDeletedBy() DisputeEvidenceField {
	return DisputeEvidenceField("meta_deleted_by")
}

func (ss DisputeEvidenceSelectFields) All() DisputeEvidenceFieldList {
	return []DisputeEvidenceField{
		ss.Id(),
		ss.DisputeId(),
		ss.EvidenceType(),
		ss.StorageUri(),
		ss.ContentHash(),
		ss.IdempotencyKey(),
		ss.SubmittedBy(),
		ss.SubmittedAt(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewDisputeEvidenceSelectFields() DisputeEvidenceSelectFields {
	return DisputeEvidenceSelectFields{}
}

type DisputeEvidenceUpdateFieldOption struct {
	useIncrement bool
}
type DisputeEvidenceUpdateField struct {
	disputeEvidenceField DisputeEvidenceField
	opt                  DisputeEvidenceUpdateFieldOption
	value                interface{}
}
type DisputeEvidenceUpdateFieldList []DisputeEvidenceUpdateField

func defaultDisputeEvidenceUpdateFieldOption() DisputeEvidenceUpdateFieldOption {
	return DisputeEvidenceUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementDisputeEvidenceOption(useIncrement bool) func(*DisputeEvidenceUpdateFieldOption) {
	return func(pcufo *DisputeEvidenceUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewDisputeEvidenceUpdateField(field DisputeEvidenceField, val interface{}, opts ...func(*DisputeEvidenceUpdateFieldOption)) DisputeEvidenceUpdateField {
	defaultOpt := defaultDisputeEvidenceUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return DisputeEvidenceUpdateField{
		disputeEvidenceField: field,
		value:                val,
		opt:                  defaultOpt,
	}
}
func defaultDisputeEvidenceUpdateFields(disputeEvidence model.DisputeEvidence) (disputeEvidenceUpdateFieldList DisputeEvidenceUpdateFieldList) {
	selectFields := NewDisputeEvidenceSelectFields()
	disputeEvidenceUpdateFieldList = append(disputeEvidenceUpdateFieldList,
		NewDisputeEvidenceUpdateField(selectFields.Id(), disputeEvidence.Id),
		NewDisputeEvidenceUpdateField(selectFields.DisputeId(), disputeEvidence.DisputeId),
		NewDisputeEvidenceUpdateField(selectFields.EvidenceType(), disputeEvidence.EvidenceType),
		NewDisputeEvidenceUpdateField(selectFields.StorageUri(), disputeEvidence.StorageUri),
		NewDisputeEvidenceUpdateField(selectFields.ContentHash(), disputeEvidence.ContentHash),
		NewDisputeEvidenceUpdateField(selectFields.IdempotencyKey(), disputeEvidence.IdempotencyKey),
		NewDisputeEvidenceUpdateField(selectFields.SubmittedBy(), disputeEvidence.SubmittedBy),
		NewDisputeEvidenceUpdateField(selectFields.SubmittedAt(), disputeEvidence.SubmittedAt),
		NewDisputeEvidenceUpdateField(selectFields.Metadata(), disputeEvidence.Metadata),
		NewDisputeEvidenceUpdateField(selectFields.MetaCreatedAt(), disputeEvidence.MetaCreatedAt),
		NewDisputeEvidenceUpdateField(selectFields.MetaCreatedBy(), disputeEvidence.MetaCreatedBy),
		NewDisputeEvidenceUpdateField(selectFields.MetaUpdatedAt(), disputeEvidence.MetaUpdatedAt),
		NewDisputeEvidenceUpdateField(selectFields.MetaUpdatedBy(), disputeEvidence.MetaUpdatedBy),
		NewDisputeEvidenceUpdateField(selectFields.MetaDeletedAt(), disputeEvidence.MetaDeletedAt),
		NewDisputeEvidenceUpdateField(selectFields.MetaDeletedBy(), disputeEvidence.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsDisputeEvidenceCommand(disputeEvidenceUpdateFieldList DisputeEvidenceUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range disputeEvidenceUpdateFieldList {
		field := string(updateField.disputeEvidenceField)
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

func (repo *RepositoryImpl) BulkCreateDisputeEvidence(ctx context.Context, disputeEvidenceList []*model.DisputeEvidence, fieldsInsert ...DisputeEvidenceField) (err error) {
	var (
		fieldsStr                string
		valueListStr             []string
		argsList                 []interface{}
		primaryIds               []model.DisputeEvidencePrimaryID
		disputeEvidenceValueList []model.DisputeEvidence
	)

	if len(fieldsInsert) == 0 {
		selectField := NewDisputeEvidenceSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, disputeEvidence := range disputeEvidenceList {

		primaryIds = append(primaryIds, disputeEvidence.ToDisputeEvidencePrimaryID())

		disputeEvidenceValueList = append(disputeEvidenceValueList, *disputeEvidence)
	}

	_, notFoundIds, err := repo.IsExistDisputeEvidenceByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputeEvidence] failed checking disputeEvidence whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.DisputeEvidencePrimaryID{}
		mapNotFoundIds := map[model.DisputeEvidencePrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "disputeEvidence", fmt.Sprintf("disputeEvidence with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsDisputeEvidence(disputeEvidenceValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(disputeEvidenceQueries.insertDisputeEvidence, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputeEvidence] failed exec create disputeEvidence query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteDisputeEvidenceByIDs(ctx context.Context, primaryIDs []model.DisputeEvidencePrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistDisputeEvidenceByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeEvidenceByIDs] failed checking disputeEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeEvidence with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"dispute_evidence\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(disputeEvidenceQueries.deleteDisputeEvidence + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeEvidenceByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputeEvidenceByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistDisputeEvidenceByIDs(ctx context.Context, ids []model.DisputeEvidencePrimaryID) (exists bool, notFoundIds []model.DisputeEvidencePrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"dispute_evidence\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(disputeEvidenceQueries.selectDisputeEvidence, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeEvidenceByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.DisputeEvidencePrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeEvidenceByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.DisputeEvidencePrimaryID]bool{}
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

// BulkUpdateDisputeEvidence is used to bulk update disputeEvidence, by default it will update all field
// if want to update specific field, then fill disputeEvidencesMapUpdateFieldsRequest else please fill disputeEvidencesMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateDisputeEvidence(ctx context.Context, disputeEvidencesMap map[model.DisputeEvidencePrimaryID]*model.DisputeEvidence, disputeEvidencesMapUpdateFieldsRequest map[model.DisputeEvidencePrimaryID]DisputeEvidenceUpdateFieldList) (err error) {
	if len(disputeEvidencesMap) == 0 && len(disputeEvidencesMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		disputeEvidencesMapUpdateField map[model.DisputeEvidencePrimaryID]DisputeEvidenceUpdateFieldList = map[model.DisputeEvidencePrimaryID]DisputeEvidenceUpdateFieldList{}
		asTableValues                  string                                                            = "myvalues"
	)

	if len(disputeEvidencesMap) > 0 {
		for id, disputeEvidence := range disputeEvidencesMap {
			if disputeEvidence == nil {
				log.Error().Err(err).Msg("[BulkUpdateDisputeEvidence] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			disputeEvidencesMapUpdateField[id] = defaultDisputeEvidenceUpdateFields(*disputeEvidence)
		}
	} else {
		disputeEvidencesMapUpdateField = disputeEvidencesMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateDisputeEvidenceQuery(disputeEvidencesMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistDisputeEvidenceByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputeEvidence] failed checking disputeEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeEvidence with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeDisputeEvidenceCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"dispute_evidence\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputeEvidence] failed exec query")
	}
	return
}

type DisputeEvidenceFieldParameter struct {
	param string
	args  []interface{}
}

func NewDisputeEvidenceFieldParameter(param string, args ...interface{}) DisputeEvidenceFieldParameter {
	return DisputeEvidenceFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateDisputeEvidenceQuery(mapDisputeEvidences map[model.DisputeEvidencePrimaryID]DisputeEvidenceUpdateFieldList, asTableValues string) (primaryIDs []model.DisputeEvidencePrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.DisputeEvidencePrimaryID]map[string]interface{}{}
	disputeEvidenceSelectFields := NewDisputeEvidenceSelectFields()
	for id, updateFields := range mapDisputeEvidences {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.disputeEvidenceField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapDisputeEvidences[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetDisputeEvidenceFieldType(updateField.disputeEvidenceField)))
			args = append(args, fields[string(updateField.disputeEvidenceField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.disputeEvidenceField))
		if updateField.disputeEvidenceField == disputeEvidenceSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.disputeEvidenceField, asTableValues, updateField.disputeEvidenceField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.disputeEvidenceField,
				"\"dispute_evidence\"", updateField.disputeEvidenceField,
				asTableValues, updateField.disputeEvidenceField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeDisputeEvidenceCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.DisputeEvidencePrimaryID, asTableValue string) (whereQry string) {
	disputeEvidenceSelectFields := NewDisputeEvidenceSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"dispute_evidence\".\"id\" = %s.\"id\"::"+GetDisputeEvidenceFieldType(disputeEvidenceSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetDisputeEvidenceFieldType(disputeEvidenceField DisputeEvidenceField) string {
	selectDisputeEvidenceFields := NewDisputeEvidenceSelectFields()
	switch disputeEvidenceField {

	case selectDisputeEvidenceFields.Id():
		return "uuid"

	case selectDisputeEvidenceFields.DisputeId():
		return "uuid"

	case selectDisputeEvidenceFields.EvidenceType():
		return "text"

	case selectDisputeEvidenceFields.StorageUri():
		return "text"

	case selectDisputeEvidenceFields.ContentHash():
		return "text"

	case selectDisputeEvidenceFields.IdempotencyKey():
		return "text"

	case selectDisputeEvidenceFields.SubmittedBy():
		return "uuid"

	case selectDisputeEvidenceFields.SubmittedAt():
		return "timestamptz"

	case selectDisputeEvidenceFields.Metadata():
		return "jsonb"

	case selectDisputeEvidenceFields.MetaCreatedAt():
		return "timestamptz"

	case selectDisputeEvidenceFields.MetaCreatedBy():
		return "uuid"

	case selectDisputeEvidenceFields.MetaUpdatedAt():
		return "timestamptz"

	case selectDisputeEvidenceFields.MetaUpdatedBy():
		return "uuid"

	case selectDisputeEvidenceFields.MetaDeletedAt():
		return "timestamptz"

	case selectDisputeEvidenceFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateDisputeEvidence(ctx context.Context, disputeEvidence *model.DisputeEvidence, fieldsInsert ...DisputeEvidenceField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewDisputeEvidenceSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.DisputeEvidencePrimaryID{
		Id: disputeEvidence.Id,
	}
	exists, err := repo.IsExistDisputeEvidenceByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputeEvidence] failed checking disputeEvidence whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "disputeEvidence", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsDisputeEvidence([]model.DisputeEvidence{*disputeEvidence}, fieldsInsert...)
	commandQuery := fmt.Sprintf(disputeEvidenceQueries.insertDisputeEvidence, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputeEvidence] failed exec create disputeEvidence query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteDisputeEvidenceByID(ctx context.Context, primaryID model.DisputeEvidencePrimaryID) (err error) {
	exists, err := repo.IsExistDisputeEvidenceByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputeEvidenceByID] failed checking disputeEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeEvidence with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeDisputeEvidenceCompositePrimaryKeyWhere([]model.DisputeEvidencePrimaryID{primaryID})
	commandQuery := fmt.Sprintf(disputeEvidenceQueries.deleteDisputeEvidence + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputeEvidenceByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeEvidenceByFilter(ctx context.Context, filter model.Filter) (result []model.DisputeEvidenceFilterResult, err error) {
	query, args, err := composeDisputeEvidenceFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeEvidenceByFilter] failed compose disputeEvidence filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeEvidenceByFilter] failed get disputeEvidence by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeDisputeEvidenceFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.DisputeEvidenceFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeDisputeEvidenceFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeDisputeEvidenceSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeDisputeEvidenceFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 15 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewDisputeEvidenceFilterFieldSpecFromStr(sourceField)
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
		if _, selected := selectedColumns["dispute_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"dispute_id\"")
			selectedColumns["dispute_id"] = struct{}{}
		}
		if _, selected := selectedColumns["evidence_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"evidence_type\"")
			selectedColumns["evidence_type"] = struct{}{}
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
		if _, selected := selectedColumns["submitted_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"submitted_by\"")
			selectedColumns["submitted_by"] = struct{}{}
		}
		if _, selected := selectedColumns["submitted_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"submitted_at\"")
			selectedColumns["submitted_at"] = struct{}{}
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

type disputeEvidenceFilterPlaceholder struct {
	index int
}

func (p *disputeEvidenceFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeDisputeEvidenceFilterPredicate(filterField model.FilterField, placeholders *disputeEvidenceFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewDisputeEvidenceFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeDisputeEvidenceFilterSQLExpr(spec)
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

func composeDisputeEvidenceFilterGroup(group model.FilterGroup, placeholders *disputeEvidenceFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeDisputeEvidenceFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeDisputeEvidenceFilterGroup(child, placeholders, args, requiredJoins)
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

func composeDisputeEvidenceFilterWhereQueries(filter model.Filter, placeholders *disputeEvidenceFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeDisputeEvidenceFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeDisputeEvidenceFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeDisputeEvidenceFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateDisputeEvidenceFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeDisputeEvidenceSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeDisputeEvidenceFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := disputeEvidenceFilterPlaceholder{index: 1}
	whereQueries, err := composeDisputeEvidenceFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewDisputeEvidenceFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeDisputeEvidenceFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeDisputeEvidenceSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"dispute_evidence\" base%s", strings.Join(selectColumns, ","), composeDisputeEvidenceFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistDisputeEvidenceByID(ctx context.Context, primaryID model.DisputeEvidencePrimaryID) (exists bool, err error) {
	whereQuery, params := composeDisputeEvidenceCompositePrimaryKeyWhere([]model.DisputeEvidencePrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", disputeEvidenceQueries.selectCountDisputeEvidence, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputeEvidenceByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeEvidence(ctx context.Context, selectFields ...DisputeEvidenceField) (disputeEvidenceList model.DisputeEvidenceList, err error) {
	var (
		defaultDisputeEvidenceSelectFields = defaultDisputeEvidenceSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputeEvidenceSelectFields = composeDisputeEvidenceSelectFields(selectFields...)
	}
	query := fmt.Sprintf(disputeEvidenceQueries.selectDisputeEvidence, defaultDisputeEvidenceSelectFields)

	err = repo.db.Read.SelectContext(ctx, &disputeEvidenceList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputeEvidence] failed get disputeEvidence list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputeEvidenceByID(ctx context.Context, primaryID model.DisputeEvidencePrimaryID, selectFields ...DisputeEvidenceField) (disputeEvidence model.DisputeEvidence, err error) {
	var (
		defaultDisputeEvidenceSelectFields = defaultDisputeEvidenceSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputeEvidenceSelectFields = composeDisputeEvidenceSelectFields(selectFields...)
	}
	whereQry, params := composeDisputeEvidenceCompositePrimaryKeyWhere([]model.DisputeEvidencePrimaryID{primaryID})
	query := fmt.Sprintf(disputeEvidenceQueries.selectDisputeEvidence+" WHERE "+whereQry, defaultDisputeEvidenceSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &disputeEvidence, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("disputeEvidence with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveDisputeEvidenceByID] failed get disputeEvidence")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateDisputeEvidenceByID(ctx context.Context, primaryID model.DisputeEvidencePrimaryID, disputeEvidence *model.DisputeEvidence, disputeEvidenceUpdateFields ...DisputeEvidenceUpdateField) (err error) {
	exists, err := repo.IsExistDisputeEvidenceByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeEvidence] failed checking disputeEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputeEvidence with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if disputeEvidence == nil {
		if len(disputeEvidenceUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateDisputeEvidenceByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		disputeEvidence = &model.DisputeEvidence{}
	}
	var (
		defaultDisputeEvidenceUpdateFields = defaultDisputeEvidenceUpdateFields(*disputeEvidence)
		tempUpdateField                    DisputeEvidenceUpdateFieldList
		selectFields                       = NewDisputeEvidenceSelectFields()
	)
	if len(disputeEvidenceUpdateFields) > 0 {
		for _, updateField := range disputeEvidenceUpdateFields {
			if updateField.disputeEvidenceField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultDisputeEvidenceUpdateFields = tempUpdateField
	}
	whereQuery, params := composeDisputeEvidenceCompositePrimaryKeyWhere([]model.DisputeEvidencePrimaryID{primaryID})
	fields, args := composeUpdateFieldsDisputeEvidenceCommand(defaultDisputeEvidenceUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(disputeEvidenceQueries.updateDisputeEvidence+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeEvidence] error when try to update disputeEvidence by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateDisputeEvidenceByFilter(ctx context.Context, filter model.Filter, disputeEvidenceUpdateFields ...DisputeEvidenceUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(disputeEvidenceUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields DisputeEvidenceUpdateFieldList
		selectFields = NewDisputeEvidenceSelectFields()
	)
	for _, updateField := range disputeEvidenceUpdateFields {
		if updateField.disputeEvidenceField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsDisputeEvidenceCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := disputeEvidenceFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeDisputeEvidenceFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"dispute_evidence\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeEvidenceByFilter] error when try to update disputeEvidence by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputeEvidenceByFilter] failed get rows affected")
	}
	return
}

var (
	disputeEvidenceQueries = struct {
		selectDisputeEvidence      string
		selectCountDisputeEvidence string
		deleteDisputeEvidence      string
		updateDisputeEvidence      string
		insertDisputeEvidence      string
	}{
		selectDisputeEvidence:      "SELECT %s FROM \"dispute_evidence\"",
		selectCountDisputeEvidence: "SELECT COUNT(\"id\") FROM \"dispute_evidence\"",
		deleteDisputeEvidence:      "DELETE FROM \"dispute_evidence\"",
		updateDisputeEvidence:      "UPDATE \"dispute_evidence\" SET %s ",
		insertDisputeEvidence:      "INSERT INTO \"dispute_evidence\" %s VALUES %s",
	}
)

type DisputeEvidenceRepository interface {
	CreateDisputeEvidence(ctx context.Context, disputeEvidence *model.DisputeEvidence, fieldsInsert ...DisputeEvidenceField) error
	BulkCreateDisputeEvidence(ctx context.Context, disputeEvidenceList []*model.DisputeEvidence, fieldsInsert ...DisputeEvidenceField) error
	ResolveDisputeEvidence(ctx context.Context, selectFields ...DisputeEvidenceField) (model.DisputeEvidenceList, error)
	ResolveDisputeEvidenceByID(ctx context.Context, primaryID model.DisputeEvidencePrimaryID, selectFields ...DisputeEvidenceField) (model.DisputeEvidence, error)
	UpdateDisputeEvidenceByID(ctx context.Context, id model.DisputeEvidencePrimaryID, disputeEvidence *model.DisputeEvidence, disputeEvidenceUpdateFields ...DisputeEvidenceUpdateField) error
	UpdateDisputeEvidenceByFilter(ctx context.Context, filter model.Filter, disputeEvidenceUpdateFields ...DisputeEvidenceUpdateField) (rowsAffected int64, err error)
	BulkUpdateDisputeEvidence(ctx context.Context, disputeEvidenceListMap map[model.DisputeEvidencePrimaryID]*model.DisputeEvidence, DisputeEvidencesMapUpdateFieldsRequest map[model.DisputeEvidencePrimaryID]DisputeEvidenceUpdateFieldList) (err error)
	DeleteDisputeEvidenceByID(ctx context.Context, id model.DisputeEvidencePrimaryID) error
	BulkDeleteDisputeEvidenceByIDs(ctx context.Context, ids []model.DisputeEvidencePrimaryID) error
	ResolveDisputeEvidenceByFilter(ctx context.Context, filter model.Filter) (result []model.DisputeEvidenceFilterResult, err error)
	IsExistDisputeEvidenceByIDs(ctx context.Context, ids []model.DisputeEvidencePrimaryID) (exists bool, notFoundIds []model.DisputeEvidencePrimaryID, err error)
	IsExistDisputeEvidenceByID(ctx context.Context, id model.DisputeEvidencePrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
