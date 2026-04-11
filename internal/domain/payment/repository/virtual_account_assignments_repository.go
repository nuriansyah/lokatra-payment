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

func composeInsertFieldsAndParamsVirtualAccountAssignments(virtualAccountAssignmentsList []model.VirtualAccountAssignments, fieldsInsert ...VirtualAccountAssignmentsField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewVirtualAccountAssignmentsSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, virtualAccountAssignments := range virtualAccountAssignmentsList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, virtualAccountAssignments.Id)
			case selectField.IntentId():
				args = append(args, virtualAccountAssignments.IntentId)
			case selectField.BankCode():
				args = append(args, virtualAccountAssignments.BankCode)
			case selectField.VaNumber():
				args = append(args, virtualAccountAssignments.VaNumber)
			case selectField.VaNumberMasked():
				args = append(args, virtualAccountAssignments.VaNumberMasked)
			case selectField.ExpiresAt():
				args = append(args, virtualAccountAssignments.ExpiresAt)
			case selectField.IsReusable():
				args = append(args, virtualAccountAssignments.IsReusable)
			case selectField.PaidAt():
				args = append(args, virtualAccountAssignments.PaidAt)
			case selectField.PspTransactionId():
				args = append(args, virtualAccountAssignments.PspTransactionId)
			case selectField.MetaCreatedAt():
				args = append(args, virtualAccountAssignments.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, virtualAccountAssignments.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, virtualAccountAssignments.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, virtualAccountAssignments.MetaUpdatedBy)

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

func composeVirtualAccountAssignmentsCompositePrimaryKeyWhere(primaryIDs []model.VirtualAccountAssignmentsPrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"virtual_account_assignments\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultVirtualAccountAssignmentsSelectFields() string {
	fields := NewVirtualAccountAssignmentsSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeVirtualAccountAssignmentsSelectFields(selectFields ...VirtualAccountAssignmentsField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type VirtualAccountAssignmentsField string
type VirtualAccountAssignmentsFieldList []VirtualAccountAssignmentsField

type VirtualAccountAssignmentsSelectFields struct {
}

func (ss VirtualAccountAssignmentsSelectFields) Id() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("id")
}

func (ss VirtualAccountAssignmentsSelectFields) IntentId() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("intent_id")
}

func (ss VirtualAccountAssignmentsSelectFields) BankCode() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("bank_code")
}

func (ss VirtualAccountAssignmentsSelectFields) VaNumber() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("va_number")
}

func (ss VirtualAccountAssignmentsSelectFields) VaNumberMasked() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("va_number_masked")
}

func (ss VirtualAccountAssignmentsSelectFields) ExpiresAt() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("expires_at")
}

func (ss VirtualAccountAssignmentsSelectFields) IsReusable() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("is_reusable")
}

func (ss VirtualAccountAssignmentsSelectFields) PaidAt() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("paid_at")
}

func (ss VirtualAccountAssignmentsSelectFields) PspTransactionId() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("psp_transaction_id")
}

func (ss VirtualAccountAssignmentsSelectFields) MetaCreatedAt() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("meta_created_at")
}

func (ss VirtualAccountAssignmentsSelectFields) MetaCreatedBy() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("meta_created_by")
}

func (ss VirtualAccountAssignmentsSelectFields) MetaUpdatedAt() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("meta_updated_at")
}

func (ss VirtualAccountAssignmentsSelectFields) MetaUpdatedBy() VirtualAccountAssignmentsField {
	return VirtualAccountAssignmentsField("meta_updated_by")
}

func (ss VirtualAccountAssignmentsSelectFields) All() VirtualAccountAssignmentsFieldList {
	return []VirtualAccountAssignmentsField{
		ss.Id(),
		ss.IntentId(),
		ss.BankCode(),
		ss.VaNumber(),
		ss.VaNumberMasked(),
		ss.ExpiresAt(),
		ss.IsReusable(),
		ss.PaidAt(),
		ss.PspTransactionId(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
	}
}

func NewVirtualAccountAssignmentsSelectFields() VirtualAccountAssignmentsSelectFields {
	return VirtualAccountAssignmentsSelectFields{}
}

type VirtualAccountAssignmentsUpdateFieldOption struct {
	useIncrement bool
}
type VirtualAccountAssignmentsUpdateField struct {
	virtualAccountAssignmentsField VirtualAccountAssignmentsField
	opt                            VirtualAccountAssignmentsUpdateFieldOption
	value                          interface{}
}
type VirtualAccountAssignmentsUpdateFieldList []VirtualAccountAssignmentsUpdateField

func defaultVirtualAccountAssignmentsUpdateFieldOption() VirtualAccountAssignmentsUpdateFieldOption {
	return VirtualAccountAssignmentsUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementVirtualAccountAssignmentsOption(useIncrement bool) func(*VirtualAccountAssignmentsUpdateFieldOption) {
	return func(pcufo *VirtualAccountAssignmentsUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewVirtualAccountAssignmentsUpdateField(field VirtualAccountAssignmentsField, val interface{}, opts ...func(*VirtualAccountAssignmentsUpdateFieldOption)) VirtualAccountAssignmentsUpdateField {
	defaultOpt := defaultVirtualAccountAssignmentsUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return VirtualAccountAssignmentsUpdateField{
		virtualAccountAssignmentsField: field,
		value:                          val,
		opt:                            defaultOpt,
	}
}
func defaultVirtualAccountAssignmentsUpdateFields(virtualAccountAssignments model.VirtualAccountAssignments) (virtualAccountAssignmentsUpdateFieldList VirtualAccountAssignmentsUpdateFieldList) {
	selectFields := NewVirtualAccountAssignmentsSelectFields()
	virtualAccountAssignmentsUpdateFieldList = append(virtualAccountAssignmentsUpdateFieldList,
		NewVirtualAccountAssignmentsUpdateField(selectFields.Id(), virtualAccountAssignments.Id),
		NewVirtualAccountAssignmentsUpdateField(selectFields.IntentId(), virtualAccountAssignments.IntentId),
		NewVirtualAccountAssignmentsUpdateField(selectFields.BankCode(), virtualAccountAssignments.BankCode),
		NewVirtualAccountAssignmentsUpdateField(selectFields.VaNumber(), virtualAccountAssignments.VaNumber),
		NewVirtualAccountAssignmentsUpdateField(selectFields.VaNumberMasked(), virtualAccountAssignments.VaNumberMasked),
		NewVirtualAccountAssignmentsUpdateField(selectFields.ExpiresAt(), virtualAccountAssignments.ExpiresAt),
		NewVirtualAccountAssignmentsUpdateField(selectFields.IsReusable(), virtualAccountAssignments.IsReusable),
		NewVirtualAccountAssignmentsUpdateField(selectFields.PaidAt(), virtualAccountAssignments.PaidAt),
		NewVirtualAccountAssignmentsUpdateField(selectFields.PspTransactionId(), virtualAccountAssignments.PspTransactionId),
		NewVirtualAccountAssignmentsUpdateField(selectFields.MetaCreatedAt(), virtualAccountAssignments.MetaCreatedAt),
		NewVirtualAccountAssignmentsUpdateField(selectFields.MetaCreatedBy(), virtualAccountAssignments.MetaCreatedBy),
		NewVirtualAccountAssignmentsUpdateField(selectFields.MetaUpdatedAt(), virtualAccountAssignments.MetaUpdatedAt),
		NewVirtualAccountAssignmentsUpdateField(selectFields.MetaUpdatedBy(), virtualAccountAssignments.MetaUpdatedBy),
	)
	return
}
func composeUpdateFieldsVirtualAccountAssignmentsCommand(virtualAccountAssignmentsUpdateFieldList VirtualAccountAssignmentsUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range virtualAccountAssignmentsUpdateFieldList {
		field := string(updateField.virtualAccountAssignmentsField)
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

func (repo *RepositoryImpl) BulkCreateVirtualAccountAssignments(ctx context.Context, virtualAccountAssignmentsList []*model.VirtualAccountAssignments, fieldsInsert ...VirtualAccountAssignmentsField) (err error) {
	var (
		fieldsStr                          string
		valueListStr                       []string
		argsList                           []interface{}
		primaryIds                         []model.VirtualAccountAssignmentsPrimaryID
		virtualAccountAssignmentsValueList []model.VirtualAccountAssignments
	)

	if len(fieldsInsert) == 0 {
		selectField := NewVirtualAccountAssignmentsSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, virtualAccountAssignments := range virtualAccountAssignmentsList {

		primaryIds = append(primaryIds, virtualAccountAssignments.ToVirtualAccountAssignmentsPrimaryID())

		virtualAccountAssignmentsValueList = append(virtualAccountAssignmentsValueList, *virtualAccountAssignments)
	}

	_, notFoundIds, err := repo.IsExistVirtualAccountAssignmentsByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateVirtualAccountAssignments] failed checking virtualAccountAssignments whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.VirtualAccountAssignmentsPrimaryID{}
		mapNotFoundIds := map[model.VirtualAccountAssignmentsPrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "virtualAccountAssignments", fmt.Sprintf("virtualAccountAssignments with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsVirtualAccountAssignments(virtualAccountAssignmentsValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(virtualAccountAssignmentsQueries.insertVirtualAccountAssignments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateVirtualAccountAssignments] failed exec create virtualAccountAssignments query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteVirtualAccountAssignmentsByIDs(ctx context.Context, primaryIDs []model.VirtualAccountAssignmentsPrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistVirtualAccountAssignmentsByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteVirtualAccountAssignmentsByIDs] failed checking virtualAccountAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("virtualAccountAssignments with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"virtual_account_assignments\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := fmt.Sprintf(virtualAccountAssignmentsQueries.deleteVirtualAccountAssignments + " WHERE " + whereQuery)

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteVirtualAccountAssignmentsByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteVirtualAccountAssignmentsByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistVirtualAccountAssignmentsByIDs(ctx context.Context, ids []model.VirtualAccountAssignmentsPrimaryID) (exists bool, notFoundIds []model.VirtualAccountAssignmentsPrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"virtual_account_assignments\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(virtualAccountAssignmentsQueries.selectVirtualAccountAssignments, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistVirtualAccountAssignmentsByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.VirtualAccountAssignmentsPrimaryID
	err = repo.db.Read.Select(&resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistVirtualAccountAssignmentsByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.VirtualAccountAssignmentsPrimaryID]bool{}
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

// BulkUpdateVirtualAccountAssignments is used to bulk update virtualAccountAssignments, by default it will update all field
// if want to update specific field, then fill virtualAccountAssignmentssMapUpdateFieldsRequest else please fill virtualAccountAssignmentssMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateVirtualAccountAssignments(ctx context.Context, virtualAccountAssignmentssMap map[model.VirtualAccountAssignmentsPrimaryID]*model.VirtualAccountAssignments, virtualAccountAssignmentssMapUpdateFieldsRequest map[model.VirtualAccountAssignmentsPrimaryID]VirtualAccountAssignmentsUpdateFieldList) (err error) {
	if len(virtualAccountAssignmentssMap) == 0 && len(virtualAccountAssignmentssMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		virtualAccountAssignmentssMapUpdateField map[model.VirtualAccountAssignmentsPrimaryID]VirtualAccountAssignmentsUpdateFieldList = map[model.VirtualAccountAssignmentsPrimaryID]VirtualAccountAssignmentsUpdateFieldList{}
		asTableValues                            string                                                                                = "myvalues"
	)

	if len(virtualAccountAssignmentssMap) > 0 {
		for id, virtualAccountAssignments := range virtualAccountAssignmentssMap {
			if virtualAccountAssignments == nil {
				log.Error().Err(err).Msg("[BulkUpdateVirtualAccountAssignments] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			virtualAccountAssignmentssMapUpdateField[id] = defaultVirtualAccountAssignmentsUpdateFields(*virtualAccountAssignments)
		}
	} else {
		virtualAccountAssignmentssMapUpdateField = virtualAccountAssignmentssMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateVirtualAccountAssignmentsQuery(virtualAccountAssignmentssMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistVirtualAccountAssignmentsByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateVirtualAccountAssignments] failed checking virtualAccountAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("virtualAccountAssignments with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeVirtualAccountAssignmentsCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"virtual_account_assignments\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateVirtualAccountAssignments] failed exec query")
	}
	return
}

type VirtualAccountAssignmentsFieldParameter struct {
	param string
	args  []interface{}
}

func NewVirtualAccountAssignmentsFieldParameter(param string, args ...interface{}) VirtualAccountAssignmentsFieldParameter {
	return VirtualAccountAssignmentsFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateVirtualAccountAssignmentsQuery(mapVirtualAccountAssignmentss map[model.VirtualAccountAssignmentsPrimaryID]VirtualAccountAssignmentsUpdateFieldList, asTableValues string) (primaryIDs []model.VirtualAccountAssignmentsPrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.VirtualAccountAssignmentsPrimaryID]map[string]interface{}{}
	virtualAccountAssignmentsSelectFields := NewVirtualAccountAssignmentsSelectFields()
	for id, updateFields := range mapVirtualAccountAssignmentss {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.virtualAccountAssignmentsField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapVirtualAccountAssignmentss[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetVirtualAccountAssignmentsFieldType(updateField.virtualAccountAssignmentsField)))
			args = append(args, fields[string(updateField.virtualAccountAssignmentsField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.virtualAccountAssignmentsField))
		if updateField.virtualAccountAssignmentsField == virtualAccountAssignmentsSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.virtualAccountAssignmentsField, asTableValues, updateField.virtualAccountAssignmentsField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.virtualAccountAssignmentsField,
				"\"virtual_account_assignments\"", updateField.virtualAccountAssignmentsField,
				asTableValues, updateField.virtualAccountAssignmentsField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeVirtualAccountAssignmentsCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.VirtualAccountAssignmentsPrimaryID, asTableValue string) (whereQry string) {
	virtualAccountAssignmentsSelectFields := NewVirtualAccountAssignmentsSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"virtual_account_assignments\".\"id\" = %s.\"id\"::"+GetVirtualAccountAssignmentsFieldType(virtualAccountAssignmentsSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetVirtualAccountAssignmentsFieldType(virtualAccountAssignmentsField VirtualAccountAssignmentsField) string {
	selectVirtualAccountAssignmentsFields := NewVirtualAccountAssignmentsSelectFields()
	switch virtualAccountAssignmentsField {

	case selectVirtualAccountAssignmentsFields.Id():
		return "uuid"

	case selectVirtualAccountAssignmentsFields.IntentId():
		return "uuid"

	case selectVirtualAccountAssignmentsFields.BankCode():
		return "text"

	case selectVirtualAccountAssignmentsFields.VaNumber():
		return "text"

	case selectVirtualAccountAssignmentsFields.VaNumberMasked():
		return "text"

	case selectVirtualAccountAssignmentsFields.ExpiresAt():
		return "timestamptz"

	case selectVirtualAccountAssignmentsFields.IsReusable():
		return "bool"

	case selectVirtualAccountAssignmentsFields.PaidAt():
		return "timestamptz"

	case selectVirtualAccountAssignmentsFields.PspTransactionId():
		return "text"

	case selectVirtualAccountAssignmentsFields.MetaCreatedAt():
		return "timestamptz"

	case selectVirtualAccountAssignmentsFields.MetaCreatedBy():
		return "uuid"

	case selectVirtualAccountAssignmentsFields.MetaUpdatedAt():
		return "timestamptz"

	case selectVirtualAccountAssignmentsFields.MetaUpdatedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateVirtualAccountAssignments(ctx context.Context, virtualAccountAssignments *model.VirtualAccountAssignments, fieldsInsert ...VirtualAccountAssignmentsField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewVirtualAccountAssignmentsSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.VirtualAccountAssignmentsPrimaryID{
		Id: virtualAccountAssignments.Id,
	}
	exists, err := repo.IsExistVirtualAccountAssignmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateVirtualAccountAssignments] failed checking virtualAccountAssignments whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "virtualAccountAssignments", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsVirtualAccountAssignments([]model.VirtualAccountAssignments{*virtualAccountAssignments}, fieldsInsert...)
	commandQuery := fmt.Sprintf(virtualAccountAssignmentsQueries.insertVirtualAccountAssignments, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateVirtualAccountAssignments] failed exec create virtualAccountAssignments query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteVirtualAccountAssignmentsByID(ctx context.Context, primaryID model.VirtualAccountAssignmentsPrimaryID) (err error) {
	exists, err := repo.IsExistVirtualAccountAssignmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteVirtualAccountAssignmentsByID] failed checking virtualAccountAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("virtualAccountAssignments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeVirtualAccountAssignmentsCompositePrimaryKeyWhere([]model.VirtualAccountAssignmentsPrimaryID{primaryID})
	commandQuery := fmt.Sprintf(virtualAccountAssignmentsQueries.deleteVirtualAccountAssignments + " WHERE " + whereQuery)
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteVirtualAccountAssignmentsByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveVirtualAccountAssignmentsByFilter(ctx context.Context, filter model.Filter) (result []model.VirtualAccountAssignmentsFilterResult, err error) {
	query, args, err := composeVirtualAccountAssignmentsFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveVirtualAccountAssignmentsByFilter] failed compose virtualAccountAssignments filter")
		return
	}
	err = repo.db.Read.Select(&result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveVirtualAccountAssignmentsByFilter] failed get virtualAccountAssignments by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeVirtualAccountAssignmentsFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateVirtualAccountAssignmentsFieldNameFilter(filter)
	if err != nil {
		return
	}
	selectFields := defaultVirtualAccountAssignmentsSelectFields()
	index := 1
	if len(filter.SelectFields) > 0 {
		fields := VirtualAccountAssignmentsFieldList{}
		for _, filterSelectField := range filter.SelectFields {
			fields = append(fields, VirtualAccountAssignmentsField(filterSelectField))
		}
		selectFields = composeVirtualAccountAssignmentsSelectFields(fields...)
	}
	selectFields += ", COUNT(*) OVER() as count"
	query = fmt.Sprintf(virtualAccountAssignmentsQueries.selectVirtualAccountAssignments, selectFields)

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

func (repo *RepositoryImpl) IsExistVirtualAccountAssignmentsByID(ctx context.Context, primaryID model.VirtualAccountAssignmentsPrimaryID) (exists bool, err error) {
	whereQuery, params := composeVirtualAccountAssignmentsCompositePrimaryKeyWhere([]model.VirtualAccountAssignmentsPrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", virtualAccountAssignmentsQueries.selectCountVirtualAccountAssignments, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistVirtualAccountAssignmentsByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveVirtualAccountAssignments(ctx context.Context, selectFields ...VirtualAccountAssignmentsField) (virtualAccountAssignmentsList model.VirtualAccountAssignmentsList, err error) {
	var (
		defaultVirtualAccountAssignmentsSelectFields = defaultVirtualAccountAssignmentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultVirtualAccountAssignmentsSelectFields = composeVirtualAccountAssignmentsSelectFields(selectFields...)
	}
	query := fmt.Sprintf(virtualAccountAssignmentsQueries.selectVirtualAccountAssignments, defaultVirtualAccountAssignmentsSelectFields)

	err = repo.db.Read.Select(&virtualAccountAssignmentsList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveVirtualAccountAssignments] failed get virtualAccountAssignments list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveVirtualAccountAssignmentsByID(ctx context.Context, primaryID model.VirtualAccountAssignmentsPrimaryID, selectFields ...VirtualAccountAssignmentsField) (virtualAccountAssignments model.VirtualAccountAssignments, err error) {
	var (
		defaultVirtualAccountAssignmentsSelectFields = defaultVirtualAccountAssignmentsSelectFields()
	)
	if len(selectFields) > 0 {
		defaultVirtualAccountAssignmentsSelectFields = composeVirtualAccountAssignmentsSelectFields(selectFields...)
	}
	whereQry, params := composeVirtualAccountAssignmentsCompositePrimaryKeyWhere([]model.VirtualAccountAssignmentsPrimaryID{primaryID})
	query := fmt.Sprintf(virtualAccountAssignmentsQueries.selectVirtualAccountAssignments+" WHERE "+whereQry, defaultVirtualAccountAssignmentsSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.Get(&virtualAccountAssignments, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("virtualAccountAssignments with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveVirtualAccountAssignmentsByID] failed get virtualAccountAssignments")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateVirtualAccountAssignmentsByID(ctx context.Context, primaryID model.VirtualAccountAssignmentsPrimaryID, virtualAccountAssignments *model.VirtualAccountAssignments, virtualAccountAssignmentsUpdateFields ...VirtualAccountAssignmentsUpdateField) (err error) {
	exists, err := repo.IsExistVirtualAccountAssignmentsByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateVirtualAccountAssignments] failed checking virtualAccountAssignments whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("virtualAccountAssignments with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if virtualAccountAssignments == nil {
		if len(virtualAccountAssignmentsUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateVirtualAccountAssignmentsByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		virtualAccountAssignments = &model.VirtualAccountAssignments{}
	}
	var (
		defaultVirtualAccountAssignmentsUpdateFields = defaultVirtualAccountAssignmentsUpdateFields(*virtualAccountAssignments)
		tempUpdateField                              VirtualAccountAssignmentsUpdateFieldList
		selectFields                                 = NewVirtualAccountAssignmentsSelectFields()
	)
	if len(virtualAccountAssignmentsUpdateFields) > 0 {
		for _, updateField := range virtualAccountAssignmentsUpdateFields {
			if updateField.virtualAccountAssignmentsField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultVirtualAccountAssignmentsUpdateFields = tempUpdateField
	}
	whereQuery, params := composeVirtualAccountAssignmentsCompositePrimaryKeyWhere([]model.VirtualAccountAssignmentsPrimaryID{primaryID})
	fields, args := composeUpdateFieldsVirtualAccountAssignmentsCommand(defaultVirtualAccountAssignmentsUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(virtualAccountAssignmentsQueries.updateVirtualAccountAssignments+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateVirtualAccountAssignments] error when try to update virtualAccountAssignments by id")
	}
	return err
}

var (
	virtualAccountAssignmentsQueries = struct {
		selectVirtualAccountAssignments      string
		selectCountVirtualAccountAssignments string
		deleteVirtualAccountAssignments      string
		updateVirtualAccountAssignments      string
		insertVirtualAccountAssignments      string
	}{
		selectVirtualAccountAssignments:      "SELECT %s FROM \"virtual_account_assignments\"",
		selectCountVirtualAccountAssignments: "SELECT COUNT(\"id\") FROM \"virtual_account_assignments\"",
		deleteVirtualAccountAssignments:      "DELETE FROM \"virtual_account_assignments\"",
		updateVirtualAccountAssignments:      "UPDATE \"virtual_account_assignments\" SET %s ",
		insertVirtualAccountAssignments:      "INSERT INTO \"virtual_account_assignments\" %s VALUES %s",
	}
)

type VirtualAccountAssignmentsRepository interface {
	CreateVirtualAccountAssignments(ctx context.Context, virtualAccountAssignments *model.VirtualAccountAssignments, fieldsInsert ...VirtualAccountAssignmentsField) error
	BulkCreateVirtualAccountAssignments(ctx context.Context, virtualAccountAssignmentsList []*model.VirtualAccountAssignments, fieldsInsert ...VirtualAccountAssignmentsField) error
	ResolveVirtualAccountAssignments(ctx context.Context, selectFields ...VirtualAccountAssignmentsField) (model.VirtualAccountAssignmentsList, error)
	ResolveVirtualAccountAssignmentsByID(ctx context.Context, primaryID model.VirtualAccountAssignmentsPrimaryID, selectFields ...VirtualAccountAssignmentsField) (model.VirtualAccountAssignments, error)
	UpdateVirtualAccountAssignmentsByID(ctx context.Context, id model.VirtualAccountAssignmentsPrimaryID, virtualAccountAssignments *model.VirtualAccountAssignments, virtualAccountAssignmentsUpdateFields ...VirtualAccountAssignmentsUpdateField) error
	BulkUpdateVirtualAccountAssignments(ctx context.Context, virtualAccountAssignmentsListMap map[model.VirtualAccountAssignmentsPrimaryID]*model.VirtualAccountAssignments, VirtualAccountAssignmentssMapUpdateFieldsRequest map[model.VirtualAccountAssignmentsPrimaryID]VirtualAccountAssignmentsUpdateFieldList) (err error)
	DeleteVirtualAccountAssignmentsByID(ctx context.Context, id model.VirtualAccountAssignmentsPrimaryID) error
	BulkDeleteVirtualAccountAssignmentsByIDs(ctx context.Context, ids []model.VirtualAccountAssignmentsPrimaryID) error
	ResolveVirtualAccountAssignmentsByFilter(ctx context.Context, filter model.Filter) (result []model.VirtualAccountAssignmentsFilterResult, err error)
	IsExistVirtualAccountAssignmentsByIDs(ctx context.Context, ids []model.VirtualAccountAssignmentsPrimaryID) (exists bool, notFoundIds []model.VirtualAccountAssignmentsPrimaryID, err error)
	IsExistVirtualAccountAssignmentsByID(ctx context.Context, id model.VirtualAccountAssignmentsPrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
