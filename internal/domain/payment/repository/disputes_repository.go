package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

func composeInsertFieldsAndParamsDisputes(disputesList []model.Disputes, fieldsInsert ...DisputesField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewDisputesSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, disputes := range disputesList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, disputes.Id)
			case selectField.DisputeCode():
				args = append(args, disputes.DisputeCode)
			case selectField.PaymentId():
				args = append(args, disputes.PaymentId)
			case selectField.PspDisputeId():
				args = append(args, disputes.PspDisputeId)
			case selectField.DisputeType():
				args = append(args, disputes.DisputeType)
			case selectField.ReasonCode():
				args = append(args, disputes.ReasonCode)
			case selectField.ReasonDescription():
				args = append(args, disputes.ReasonDescription)
			case selectField.Amount():
				args = append(args, disputes.Amount)
			case selectField.Currency():
				args = append(args, disputes.Currency)
			case selectField.Status():
				args = append(args, disputes.Status)
			case selectField.OpenedAt():
				args = append(args, disputes.OpenedAt)
			case selectField.RespondBy():
				args = append(args, disputes.RespondBy)
			case selectField.ResolvedAt():
				args = append(args, disputes.ResolvedAt)
			case selectField.Outcome():
				args = append(args, disputes.Outcome)
			case selectField.EvidenceDueAt():
				args = append(args, disputes.EvidenceDueAt)
			case selectField.EvidenceSubmittedAt():
				args = append(args, disputes.EvidenceSubmittedAt)
			case selectField.EvidenceFiles():
				args = append(args, disputes.EvidenceFiles)
			case selectField.Notes():
				args = append(args, disputes.Notes)
			case selectField.HandledBy():
				args = append(args, disputes.HandledBy)
			case selectField.MetaCreatedAt():
				args = append(args, disputes.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, disputes.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, disputes.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, disputes.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, disputes.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, disputes.MetaDeletedBy)

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

func composeDisputesCompositePrimaryKeyWhere(primaryIDs []model.DisputesPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"disputes\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultDisputesSelectFields() string {
	fields := NewDisputesSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeDisputesSelectFields(selectFields ...DisputesField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type DisputesField string
type DisputesFieldList []DisputesField

type DisputesSelectFields struct {
}

func (ss DisputesSelectFields) Id() DisputesField {
	return DisputesField("id")
}

func (ss DisputesSelectFields) DisputeCode() DisputesField {
	return DisputesField("dispute_code")
}

func (ss DisputesSelectFields) PaymentId() DisputesField {
	return DisputesField("payment_id")
}

func (ss DisputesSelectFields) PspDisputeId() DisputesField {
	return DisputesField("psp_dispute_id")
}

func (ss DisputesSelectFields) DisputeType() DisputesField {
	return DisputesField("dispute_type")
}

func (ss DisputesSelectFields) ReasonCode() DisputesField {
	return DisputesField("reason_code")
}

func (ss DisputesSelectFields) ReasonDescription() DisputesField {
	return DisputesField("reason_description")
}

func (ss DisputesSelectFields) Amount() DisputesField {
	return DisputesField("amount")
}

func (ss DisputesSelectFields) Currency() DisputesField {
	return DisputesField("currency")
}

func (ss DisputesSelectFields) Status() DisputesField {
	return DisputesField("status")
}

func (ss DisputesSelectFields) OpenedAt() DisputesField {
	return DisputesField("opened_at")
}

func (ss DisputesSelectFields) RespondBy() DisputesField {
	return DisputesField("respond_by")
}

func (ss DisputesSelectFields) ResolvedAt() DisputesField {
	return DisputesField("resolved_at")
}

func (ss DisputesSelectFields) Outcome() DisputesField {
	return DisputesField("outcome")
}

func (ss DisputesSelectFields) EvidenceDueAt() DisputesField {
	return DisputesField("evidence_due_at")
}

func (ss DisputesSelectFields) EvidenceSubmittedAt() DisputesField {
	return DisputesField("evidence_submitted_at")
}

func (ss DisputesSelectFields) EvidenceFiles() DisputesField {
	return DisputesField("evidence_files")
}

func (ss DisputesSelectFields) Notes() DisputesField {
	return DisputesField("notes")
}

func (ss DisputesSelectFields) HandledBy() DisputesField {
	return DisputesField("handled_by")
}

func (ss DisputesSelectFields) MetaCreatedAt() DisputesField {
	return DisputesField("meta_created_at")
}

func (ss DisputesSelectFields) MetaCreatedBy() DisputesField {
	return DisputesField("meta_created_by")
}

func (ss DisputesSelectFields) MetaUpdatedAt() DisputesField {
	return DisputesField("meta_updated_at")
}

func (ss DisputesSelectFields) MetaUpdatedBy() DisputesField {
	return DisputesField("meta_updated_by")
}

func (ss DisputesSelectFields) MetaDeletedAt() DisputesField {
	return DisputesField("meta_deleted_at")
}

func (ss DisputesSelectFields) MetaDeletedBy() DisputesField {
	return DisputesField("meta_deleted_by")
}

func (ss DisputesSelectFields) All() DisputesFieldList {
	return []DisputesField{
		ss.Id(),
		ss.DisputeCode(),
		ss.PaymentId(),
		ss.PspDisputeId(),
		ss.DisputeType(),
		ss.ReasonCode(),
		ss.ReasonDescription(),
		ss.Amount(),
		ss.Currency(),
		ss.Status(),
		ss.OpenedAt(),
		ss.RespondBy(),
		ss.ResolvedAt(),
		ss.Outcome(),
		ss.EvidenceDueAt(),
		ss.EvidenceSubmittedAt(),
		ss.EvidenceFiles(),
		ss.Notes(),
		ss.HandledBy(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewDisputesSelectFields() DisputesSelectFields {
	return DisputesSelectFields{}
}

type DisputesUpdateFieldOption struct {
	useIncrement bool
}
type DisputesUpdateField struct {
	disputesField DisputesField
	opt           DisputesUpdateFieldOption
	value         interface{}
}
type DisputesUpdateFieldList []DisputesUpdateField

func defaultDisputesUpdateFieldOption() DisputesUpdateFieldOption {
	return DisputesUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementDisputesOption(useIncrement bool) func(*DisputesUpdateFieldOption) {
	return func(pcufo *DisputesUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewDisputesUpdateField(field DisputesField, val interface{}, opts ...func(*DisputesUpdateFieldOption)) DisputesUpdateField {
	defaultOpt := defaultDisputesUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return DisputesUpdateField{
		disputesField: field,
		value:         val,
		opt:           defaultOpt,
	}
}
func defaultDisputesUpdateFields(disputes model.Disputes) (disputesUpdateFieldList DisputesUpdateFieldList) {
	selectFields := NewDisputesSelectFields()
	disputesUpdateFieldList = append(disputesUpdateFieldList,
		NewDisputesUpdateField(selectFields.Id(), disputes.Id),
		NewDisputesUpdateField(selectFields.DisputeCode(), disputes.DisputeCode),
		NewDisputesUpdateField(selectFields.PaymentId(), disputes.PaymentId),
		NewDisputesUpdateField(selectFields.PspDisputeId(), disputes.PspDisputeId),
		NewDisputesUpdateField(selectFields.DisputeType(), disputes.DisputeType),
		NewDisputesUpdateField(selectFields.ReasonCode(), disputes.ReasonCode),
		NewDisputesUpdateField(selectFields.ReasonDescription(), disputes.ReasonDescription),
		NewDisputesUpdateField(selectFields.Amount(), disputes.Amount),
		NewDisputesUpdateField(selectFields.Currency(), disputes.Currency),
		NewDisputesUpdateField(selectFields.Status(), disputes.Status),
		NewDisputesUpdateField(selectFields.OpenedAt(), disputes.OpenedAt),
		NewDisputesUpdateField(selectFields.RespondBy(), disputes.RespondBy),
		NewDisputesUpdateField(selectFields.ResolvedAt(), disputes.ResolvedAt),
		NewDisputesUpdateField(selectFields.Outcome(), disputes.Outcome),
		NewDisputesUpdateField(selectFields.EvidenceDueAt(), disputes.EvidenceDueAt),
		NewDisputesUpdateField(selectFields.EvidenceSubmittedAt(), disputes.EvidenceSubmittedAt),
		NewDisputesUpdateField(selectFields.EvidenceFiles(), disputes.EvidenceFiles),
		NewDisputesUpdateField(selectFields.Notes(), disputes.Notes),
		NewDisputesUpdateField(selectFields.HandledBy(), disputes.HandledBy),
		NewDisputesUpdateField(selectFields.MetaCreatedAt(), disputes.MetaCreatedAt),
		NewDisputesUpdateField(selectFields.MetaCreatedBy(), disputes.MetaCreatedBy),
		NewDisputesUpdateField(selectFields.MetaUpdatedAt(), disputes.MetaUpdatedAt),
		NewDisputesUpdateField(selectFields.MetaUpdatedBy(), disputes.MetaUpdatedBy),
		NewDisputesUpdateField(selectFields.MetaDeletedAt(), disputes.MetaDeletedAt),
		NewDisputesUpdateField(selectFields.MetaDeletedBy(), disputes.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsDisputesCommand(disputesUpdateFieldList DisputesUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range disputesUpdateFieldList {
		field := string(updateField.disputesField)
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

func (repo *RepositoryImpl) BulkCreateDisputes(ctx context.Context, disputesList []*model.Disputes, fieldsInsert ...DisputesField) (err error) {
	var (
		fieldsStr         string
		valueListStr      []string
		argsList          []interface{}
		primaryIds        []model.DisputesPrimaryID
		disputesValueList []model.Disputes
	)

	if len(fieldsInsert) == 0 {
		selectField := NewDisputesSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, disputes := range disputesList {

		primaryIds = append(primaryIds, disputes.ToDisputesPrimaryID())

		disputesValueList = append(disputesValueList, *disputes)
	}

	_, notFoundIds, err := repo.IsExistDisputesByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputes] failed checking disputes whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.DisputesPrimaryID{}
		mapNotFoundIds := map[model.DisputesPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "disputes", fmt.Sprintf("disputes with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsDisputes(disputesValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(disputesQueries.insertDisputes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateDisputes] failed exec create disputes query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteDisputesByIDs(ctx context.Context, primaryIDs []model.DisputesPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistDisputesByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputesByIDs] failed checking disputes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputes with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"disputes\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(disputesQueries.deleteDisputes + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputesByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteDisputesByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistDisputesByIDs(ctx context.Context, ids []model.DisputesPrimaryID) (exists bool, notFoundIds []model.DisputesPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"disputes\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(disputesQueries.selectDisputes, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputesByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.DisputesPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputesByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.DisputesPrimaryID]bool{}
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

// BulkUpdateDisputes is used to bulk update disputes, by default it will update all field
// if want to update specific field, then fill disputessMapUpdateFieldsRequest else please fill disputessMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateDisputes(ctx context.Context, disputessMap map[model.DisputesPrimaryID]*model.Disputes, disputessMapUpdateFieldsRequest map[model.DisputesPrimaryID]DisputesUpdateFieldList) (err error) {
	if len(disputessMap) == 0 && len(disputessMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		disputessMapUpdateField map[model.DisputesPrimaryID]DisputesUpdateFieldList = map[model.DisputesPrimaryID]DisputesUpdateFieldList{}
		asTableValues           string                                              = "myvalues"
	)

	if len(disputessMap) > 0 {
		for id, disputes := range disputessMap {
			if disputes == nil {
				log.Error().Err(err).Msg("[BulkUpdateDisputes] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			disputessMapUpdateField[id] = defaultDisputesUpdateFields(*disputes)
		}
	} else {
		disputessMapUpdateField = disputessMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateDisputesQuery(disputessMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistDisputesByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputes] failed checking disputes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputes with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeDisputesCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"disputes\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateDisputes] failed exec query")
	}
	return
}

type DisputesFieldParameter struct {
	param string
	args  []interface{}
}

func NewDisputesFieldParameter(param string, args ...interface{}) DisputesFieldParameter {
	return DisputesFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateDisputesQuery(mapDisputess map[model.DisputesPrimaryID]DisputesUpdateFieldList, asTableValues string) (primaryIDs []model.DisputesPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.DisputesPrimaryID]map[string]interface{}{}
	disputesSelectFields := NewDisputesSelectFields()
	for id, updateFields := range mapDisputess {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.disputesField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapDisputess[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetDisputesFieldType(updateField.disputesField)))
			args = append(args, fields[string(updateField.disputesField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.disputesField))
		if updateField.disputesField == disputesSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.disputesField, asTableValues, updateField.disputesField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.disputesField,
				"\"disputes\"", updateField.disputesField,
				asTableValues, updateField.disputesField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeDisputesCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.DisputesPrimaryID, asTableValue string) (whereQry string) {
	disputesSelectFields := NewDisputesSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"disputes\".\"id\" = %s.\"id\"::"+GetDisputesFieldType(disputesSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetDisputesFieldType(disputesField DisputesField) string {
	selectDisputesFields := NewDisputesSelectFields()
	switch disputesField {

	case selectDisputesFields.Id():
		return "uuid"

	case selectDisputesFields.DisputeCode():
		return "text"

	case selectDisputesFields.PaymentId():
		return "uuid"

	case selectDisputesFields.PspDisputeId():
		return "text"

	case selectDisputesFields.DisputeType():
		return "text"

	case selectDisputesFields.ReasonCode():
		return "text"

	case selectDisputesFields.ReasonDescription():
		return "text"

	case selectDisputesFields.Amount():
		return "numeric"

	case selectDisputesFields.Currency():
		return "payment_currency"

	case selectDisputesFields.Status():
		return "dispute_status_enum"

	case selectDisputesFields.OpenedAt():
		return "timestamptz"

	case selectDisputesFields.RespondBy():
		return "timestamptz"

	case selectDisputesFields.ResolvedAt():
		return "timestamptz"

	case selectDisputesFields.Outcome():
		return "text"

	case selectDisputesFields.EvidenceDueAt():
		return "timestamptz"

	case selectDisputesFields.EvidenceSubmittedAt():
		return "timestamptz"

	case selectDisputesFields.EvidenceFiles():
		return "jsonb"

	case selectDisputesFields.Notes():
		return "text"

	case selectDisputesFields.HandledBy():
		return "uuid"

	case selectDisputesFields.MetaCreatedAt():
		return "timestamptz"

	case selectDisputesFields.MetaCreatedBy():
		return "uuid"

	case selectDisputesFields.MetaUpdatedAt():
		return "timestamptz"

	case selectDisputesFields.MetaUpdatedBy():
		return "uuid"

	case selectDisputesFields.MetaDeletedAt():
		return "timestamptz"

	case selectDisputesFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateDisputes(ctx context.Context, disputes *model.Disputes, fieldsInsert ...DisputesField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewDisputesSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.DisputesPrimaryID{
		Id: disputes.Id,
	}
	exists, err := repo.IsExistDisputesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputes] failed checking disputes whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "disputes", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsDisputes([]model.Disputes{*disputes}, fieldsInsert...)
	commandQuery := fmt.Sprintf(disputesQueries.insertDisputes, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateDisputes] failed exec create disputes query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteDisputesByID(ctx context.Context, primaryID model.DisputesPrimaryID) (err error) {
	exists, err := repo.IsExistDisputesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputesByID] failed checking disputes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeDisputesCompositePrimaryKeyWhere([]model.DisputesPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(disputesQueries.deleteDisputes + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteDisputesByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputesByFilter(ctx context.Context, filter model.Filter) (result []model.DisputesFilterResult, err error) {
	query, args, err := composeDisputesFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputesByFilter] failed compose disputes filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputesByFilter] failed get disputes by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeDisputesFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateDisputesFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultDisputesSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := DisputesFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, DisputesField(filterSelectField))
		}
		selectFields = composeDisputesSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(disputesQueries.selectDisputes, selectFields)

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

func (repo *RepositoryImpl) IsExistDisputesByID(ctx context.Context, primaryID model.DisputesPrimaryID) (exists bool, err error) {
	whereQuery, params := composeDisputesCompositePrimaryKeyWhere([]model.DisputesPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", disputesQueries.selectCountDisputes, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistDisputesByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputes(ctx context.Context, selectFields ...DisputesField) (disputesList model.DisputesList, err error) {
	var (
		defaultDisputesSelectFields = defaultDisputesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputesSelectFields = composeDisputesSelectFields(selectFields...)
	}
	query := fmt.Sprintf(disputesQueries.selectDisputes, defaultDisputesSelectFields)

	err = repo.db.Read.Select(&disputesList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveDisputes] failed get disputes list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveDisputesByID(ctx context.Context, primaryID model.DisputesPrimaryID, selectFields ...DisputesField) (disputes model.Disputes, err error) {
	var (
		defaultDisputesSelectFields = defaultDisputesSelectFields()
	)
	if len(selectFields) > 0 {
		defaultDisputesSelectFields = composeDisputesSelectFields(selectFields...)
	}
	whereQry, params := composeDisputesCompositePrimaryKeyWhere([]model.DisputesPrimaryID{primaryID})
	query := fmt.Sprintf(disputesQueries.selectDisputes+" WHERE "+whereQry, defaultDisputesSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&disputes, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("disputes with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveDisputesByID] failed get disputes")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateDisputesByID(ctx context.Context, primaryID model.DisputesPrimaryID, disputes *model.Disputes, disputesUpdateFields ...DisputesUpdateField) (err error) {
	exists, err := repo.IsExistDisputesByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputes] failed checking disputes whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("disputes with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if disputes == nil {
		if len(disputesUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateDisputesByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		disputes = &model.Disputes{}
	}
	var (
		defaultDisputesUpdateFields = defaultDisputesUpdateFields(*disputes)
		tempUpdateField             DisputesUpdateFieldList
		selectFields                = NewDisputesSelectFields()
	)
	if len(disputesUpdateFields) > 0 {
		for _, updateField := range disputesUpdateFields {
			if updateField.disputesField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultDisputesUpdateFields = tempUpdateField
	}
	whereQuery, params := composeDisputesCompositePrimaryKeyWhere([]model.DisputesPrimaryID{primaryID})
	fields, args := composeUpdateFieldsDisputesCommand(defaultDisputesUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(disputesQueries.updateDisputes+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateDisputes] error when try to update disputes by id")
	}
	return err
}

var (
	disputesQueries = struct {
		selectDisputes      string
		selectCountDisputes string
		deleteDisputes      string
		updateDisputes      string
		insertDisputes      string
	}{
		selectDisputes:      "SELECT %s FROM \"disputes\"",
		selectCountDisputes: "SELECT COUNT(\"id\") FROM \"disputes\"",
		deleteDisputes:      "DELETE FROM \"disputes\"",
		updateDisputes:      "UPDATE \"disputes\" SET %s ",
		insertDisputes:      "INSERT INTO \"disputes\" %s VALUES %s",
	}
)

type DisputesRepository interface {
	CreateDisputes(ctx context.Context, disputes *model.Disputes, fieldsInsert ...DisputesField) error
	BulkCreateDisputes(ctx context.Context, disputesList []*model.Disputes, fieldsInsert ...DisputesField) error
	ResolveDisputes(ctx context.Context, selectFields ...DisputesField) (model.DisputesList, error)
	ResolveDisputesByID(ctx context.Context, primaryID model.DisputesPrimaryID, selectFields ...DisputesField) (model.Disputes, error)
	UpdateDisputesByID(ctx context.Context, id model.DisputesPrimaryID, disputes *model.Disputes, disputesUpdateFields ...DisputesUpdateField) error
	BulkUpdateDisputes(ctx context.Context, disputesListMap map[model.DisputesPrimaryID]*model.Disputes, DisputessMapUpdateFieldsRequest map[model.DisputesPrimaryID]DisputesUpdateFieldList) (err error)
	DeleteDisputesByID(ctx context.Context, id model.DisputesPrimaryID) error
	BulkDeleteDisputesByIDs(ctx context.Context, ids []model.DisputesPrimaryID) error
	ResolveDisputesByFilter(ctx context.Context, filter model.Filter) (result []model.DisputesFilterResult, err error)
	IsExistDisputesByIDs(ctx context.Context, ids []model.DisputesPrimaryID) (exists bool, notFoundIds []model.DisputesPrimaryID, err error)
	IsExistDisputesByID(ctx context.Context, id model.DisputesPrimaryID) (exists bool, err error)
}
