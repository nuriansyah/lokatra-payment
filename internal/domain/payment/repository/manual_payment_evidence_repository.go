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

func composeInsertFieldsAndParamsManualPaymentEvidence(manualPaymentEvidenceList []model.ManualPaymentEvidence, fieldsInsert ...ManualPaymentEvidenceField) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string
		index  = 0
	)

	selectField := NewManualPaymentEvidenceSelectFields()
	for _, field := range fieldsInsert {

		fields = append(fields, fmt.Sprintf("\"%s\"", string(field)))
	}
	for _, manualPaymentEvidence := range manualPaymentEvidenceList {
		var values []string
		for _, field := range fieldsInsert {
			switch field {
			case selectField.Id():
				args = append(args, manualPaymentEvidence.Id)
			case selectField.PaymentIntentId():
				args = append(args, manualPaymentEvidence.PaymentIntentId)
			case selectField.PaymentAttemptId():
				args = append(args, manualPaymentEvidence.PaymentAttemptId)
			case selectField.SubmittedBy():
				args = append(args, manualPaymentEvidence.SubmittedBy)
			case selectField.EvidenceType():
				args = append(args, manualPaymentEvidence.EvidenceType)
			case selectField.EvidenceUrl():
				args = append(args, manualPaymentEvidence.EvidenceUrl)
			case selectField.Amount():
				args = append(args, manualPaymentEvidence.Amount)
			case selectField.ExpectedAmount():
				args = append(args, manualPaymentEvidence.ExpectedAmount)
			case selectField.VarianceAmount():
				args = append(args, manualPaymentEvidence.VarianceAmount)
			case selectField.VarianceStatus():
				args = append(args, manualPaymentEvidence.VarianceStatus)
			case selectField.Currency():
				args = append(args, manualPaymentEvidence.Currency)
			case selectField.BankCode():
				args = append(args, manualPaymentEvidence.BankCode)
			case selectField.BankName():
				args = append(args, manualPaymentEvidence.BankName)
			case selectField.SenderAccountName():
				args = append(args, manualPaymentEvidence.SenderAccountName)
			case selectField.SenderAccountNumberMasked():
				args = append(args, manualPaymentEvidence.SenderAccountNumberMasked)
			case selectField.Notes():
				args = append(args, manualPaymentEvidence.Notes)
			case selectField.Status():
				args = append(args, manualPaymentEvidence.Status)
			case selectField.ReviewedBy():
				args = append(args, manualPaymentEvidence.ReviewedBy)
			case selectField.ReviewedAt():
				args = append(args, manualPaymentEvidence.ReviewedAt)
			case selectField.RejectionReason():
				args = append(args, manualPaymentEvidence.RejectionReason)
			case selectField.PolicyDecision():
				args = append(args, manualPaymentEvidence.PolicyDecision)
			case selectField.Metadata():
				args = append(args, manualPaymentEvidence.Metadata)
			case selectField.MetaCreatedAt():
				args = append(args, manualPaymentEvidence.MetaCreatedAt)
			case selectField.MetaCreatedBy():
				args = append(args, manualPaymentEvidence.MetaCreatedBy)
			case selectField.MetaUpdatedAt():
				args = append(args, manualPaymentEvidence.MetaUpdatedAt)
			case selectField.MetaUpdatedBy():
				args = append(args, manualPaymentEvidence.MetaUpdatedBy)
			case selectField.MetaDeletedAt():
				args = append(args, manualPaymentEvidence.MetaDeletedAt)
			case selectField.MetaDeletedBy():
				args = append(args, manualPaymentEvidence.MetaDeletedBy)

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

func composeManualPaymentEvidenceCompositePrimaryKeyWhere(primaryIDs []model.ManualPaymentEvidencePrimaryID) (whereQry string, params []interface{}) {
	var primaryKeyQry []string
	for _, primaryID := range primaryIDs {
		var arrWhereQry []string
		id := "\"manual_payment_evidence\".\"id\" = ?"
		params = append(params, primaryID.Id)
		arrWhereQry = append(arrWhereQry, id)

		qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))
		primaryKeyQry = append(primaryKeyQry, qry)
	}

	return strings.Join(primaryKeyQry, " OR "), params
}

func defaultManualPaymentEvidenceSelectFields() string {
	fields := NewManualPaymentEvidenceSelectFields().All()
	fieldsStr := []string{}
	for _, field := range fields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

func composeManualPaymentEvidenceSelectFields(selectFields ...ManualPaymentEvidenceField) string {
	fieldsStr := []string{}
	for _, field := range selectFields {
		fieldWithBacktick := fmt.Sprintf("\"%s\"", string(field))
		fieldsStr = append(fieldsStr, fieldWithBacktick)
	}
	return strings.Join(fieldsStr, ",")
}

type ManualPaymentEvidenceField string
type ManualPaymentEvidenceFieldList []ManualPaymentEvidenceField

type ManualPaymentEvidenceSelectFields struct {
}

func (ss ManualPaymentEvidenceSelectFields) Id() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("id")
}

func (ss ManualPaymentEvidenceSelectFields) PaymentIntentId() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("payment_intent_id")
}

func (ss ManualPaymentEvidenceSelectFields) PaymentAttemptId() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("payment_attempt_id")
}

func (ss ManualPaymentEvidenceSelectFields) SubmittedBy() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("submitted_by")
}

func (ss ManualPaymentEvidenceSelectFields) EvidenceType() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("evidence_type")
}

func (ss ManualPaymentEvidenceSelectFields) EvidenceUrl() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("evidence_url")
}

func (ss ManualPaymentEvidenceSelectFields) Amount() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("amount")
}

func (ss ManualPaymentEvidenceSelectFields) ExpectedAmount() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("expected_amount")
}

func (ss ManualPaymentEvidenceSelectFields) VarianceAmount() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("variance_amount")
}

func (ss ManualPaymentEvidenceSelectFields) VarianceStatus() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("variance_status")
}

func (ss ManualPaymentEvidenceSelectFields) Currency() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("currency")
}

func (ss ManualPaymentEvidenceSelectFields) BankCode() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("bank_code")
}

func (ss ManualPaymentEvidenceSelectFields) BankName() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("bank_name")
}

func (ss ManualPaymentEvidenceSelectFields) SenderAccountName() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("sender_account_name")
}

func (ss ManualPaymentEvidenceSelectFields) SenderAccountNumberMasked() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("sender_account_number_masked")
}

func (ss ManualPaymentEvidenceSelectFields) Notes() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("notes")
}

func (ss ManualPaymentEvidenceSelectFields) Status() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("status")
}

func (ss ManualPaymentEvidenceSelectFields) ReviewedBy() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("reviewed_by")
}

func (ss ManualPaymentEvidenceSelectFields) ReviewedAt() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("reviewed_at")
}

func (ss ManualPaymentEvidenceSelectFields) RejectionReason() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("rejection_reason")
}

func (ss ManualPaymentEvidenceSelectFields) PolicyDecision() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("policy_decision")
}

func (ss ManualPaymentEvidenceSelectFields) Metadata() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("metadata")
}

func (ss ManualPaymentEvidenceSelectFields) MetaCreatedAt() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("meta_created_at")
}

func (ss ManualPaymentEvidenceSelectFields) MetaCreatedBy() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("meta_created_by")
}

func (ss ManualPaymentEvidenceSelectFields) MetaUpdatedAt() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("meta_updated_at")
}

func (ss ManualPaymentEvidenceSelectFields) MetaUpdatedBy() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("meta_updated_by")
}

func (ss ManualPaymentEvidenceSelectFields) MetaDeletedAt() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("meta_deleted_at")
}

func (ss ManualPaymentEvidenceSelectFields) MetaDeletedBy() ManualPaymentEvidenceField {
	return ManualPaymentEvidenceField("meta_deleted_by")
}

func (ss ManualPaymentEvidenceSelectFields) All() ManualPaymentEvidenceFieldList {
	return []ManualPaymentEvidenceField{
		ss.Id(),
		ss.PaymentIntentId(),
		ss.PaymentAttemptId(),
		ss.SubmittedBy(),
		ss.EvidenceType(),
		ss.EvidenceUrl(),
		ss.Amount(),
		ss.ExpectedAmount(),
		ss.VarianceAmount(),
		ss.VarianceStatus(),
		ss.Currency(),
		ss.BankCode(),
		ss.BankName(),
		ss.SenderAccountName(),
		ss.SenderAccountNumberMasked(),
		ss.Notes(),
		ss.Status(),
		ss.ReviewedBy(),
		ss.ReviewedAt(),
		ss.RejectionReason(),
		ss.PolicyDecision(),
		ss.Metadata(),
		ss.MetaCreatedAt(),
		ss.MetaCreatedBy(),
		ss.MetaUpdatedAt(),
		ss.MetaUpdatedBy(),
		ss.MetaDeletedAt(),
		ss.MetaDeletedBy(),
	}
}

func NewManualPaymentEvidenceSelectFields() ManualPaymentEvidenceSelectFields {
	return ManualPaymentEvidenceSelectFields{}
}

type ManualPaymentEvidenceUpdateFieldOption struct {
	useIncrement bool
}
type ManualPaymentEvidenceUpdateField struct {
	manualPaymentEvidenceField ManualPaymentEvidenceField
	opt                        ManualPaymentEvidenceUpdateFieldOption
	value                      interface{}
}
type ManualPaymentEvidenceUpdateFieldList []ManualPaymentEvidenceUpdateField

func defaultManualPaymentEvidenceUpdateFieldOption() ManualPaymentEvidenceUpdateFieldOption {
	return ManualPaymentEvidenceUpdateFieldOption{
		useIncrement: false,
	}
}
func SetUseIncrementManualPaymentEvidenceOption(useIncrement bool) func(*ManualPaymentEvidenceUpdateFieldOption) {
	return func(pcufo *ManualPaymentEvidenceUpdateFieldOption) {
		pcufo.useIncrement = useIncrement
	}
}

func NewManualPaymentEvidenceUpdateField(field ManualPaymentEvidenceField, val interface{}, opts ...func(*ManualPaymentEvidenceUpdateFieldOption)) ManualPaymentEvidenceUpdateField {
	defaultOpt := defaultManualPaymentEvidenceUpdateFieldOption()
	for _, opt := range opts {
		opt(&defaultOpt)
	}
	return ManualPaymentEvidenceUpdateField{
		manualPaymentEvidenceField: field,
		value:                      val,
		opt:                        defaultOpt,
	}
}
func defaultManualPaymentEvidenceUpdateFields(manualPaymentEvidence model.ManualPaymentEvidence) (manualPaymentEvidenceUpdateFieldList ManualPaymentEvidenceUpdateFieldList) {
	selectFields := NewManualPaymentEvidenceSelectFields()
	manualPaymentEvidenceUpdateFieldList = append(manualPaymentEvidenceUpdateFieldList,
		NewManualPaymentEvidenceUpdateField(selectFields.Id(), manualPaymentEvidence.Id),
		NewManualPaymentEvidenceUpdateField(selectFields.PaymentIntentId(), manualPaymentEvidence.PaymentIntentId),
		NewManualPaymentEvidenceUpdateField(selectFields.PaymentAttemptId(), manualPaymentEvidence.PaymentAttemptId),
		NewManualPaymentEvidenceUpdateField(selectFields.SubmittedBy(), manualPaymentEvidence.SubmittedBy),
		NewManualPaymentEvidenceUpdateField(selectFields.EvidenceType(), manualPaymentEvidence.EvidenceType),
		NewManualPaymentEvidenceUpdateField(selectFields.EvidenceUrl(), manualPaymentEvidence.EvidenceUrl),
		NewManualPaymentEvidenceUpdateField(selectFields.Amount(), manualPaymentEvidence.Amount),
		NewManualPaymentEvidenceUpdateField(selectFields.ExpectedAmount(), manualPaymentEvidence.ExpectedAmount),
		NewManualPaymentEvidenceUpdateField(selectFields.VarianceAmount(), manualPaymentEvidence.VarianceAmount),
		NewManualPaymentEvidenceUpdateField(selectFields.VarianceStatus(), manualPaymentEvidence.VarianceStatus),
		NewManualPaymentEvidenceUpdateField(selectFields.Currency(), manualPaymentEvidence.Currency),
		NewManualPaymentEvidenceUpdateField(selectFields.BankCode(), manualPaymentEvidence.BankCode),
		NewManualPaymentEvidenceUpdateField(selectFields.BankName(), manualPaymentEvidence.BankName),
		NewManualPaymentEvidenceUpdateField(selectFields.SenderAccountName(), manualPaymentEvidence.SenderAccountName),
		NewManualPaymentEvidenceUpdateField(selectFields.SenderAccountNumberMasked(), manualPaymentEvidence.SenderAccountNumberMasked),
		NewManualPaymentEvidenceUpdateField(selectFields.Notes(), manualPaymentEvidence.Notes),
		NewManualPaymentEvidenceUpdateField(selectFields.Status(), manualPaymentEvidence.Status),
		NewManualPaymentEvidenceUpdateField(selectFields.ReviewedBy(), manualPaymentEvidence.ReviewedBy),
		NewManualPaymentEvidenceUpdateField(selectFields.ReviewedAt(), manualPaymentEvidence.ReviewedAt),
		NewManualPaymentEvidenceUpdateField(selectFields.RejectionReason(), manualPaymentEvidence.RejectionReason),
		NewManualPaymentEvidenceUpdateField(selectFields.PolicyDecision(), manualPaymentEvidence.PolicyDecision),
		NewManualPaymentEvidenceUpdateField(selectFields.Metadata(), manualPaymentEvidence.Metadata),
		NewManualPaymentEvidenceUpdateField(selectFields.MetaCreatedAt(), manualPaymentEvidence.MetaCreatedAt),
		NewManualPaymentEvidenceUpdateField(selectFields.MetaCreatedBy(), manualPaymentEvidence.MetaCreatedBy),
		NewManualPaymentEvidenceUpdateField(selectFields.MetaUpdatedAt(), manualPaymentEvidence.MetaUpdatedAt),
		NewManualPaymentEvidenceUpdateField(selectFields.MetaUpdatedBy(), manualPaymentEvidence.MetaUpdatedBy),
		NewManualPaymentEvidenceUpdateField(selectFields.MetaDeletedAt(), manualPaymentEvidence.MetaDeletedAt),
		NewManualPaymentEvidenceUpdateField(selectFields.MetaDeletedBy(), manualPaymentEvidence.MetaDeletedBy),
	)
	return
}
func composeUpdateFieldsManualPaymentEvidenceCommand(manualPaymentEvidenceUpdateFieldList ManualPaymentEvidenceUpdateFieldList, startIndex int) ([]string, []interface{}) {
	var (
		updatedFieldsQuery []string
		args               []interface{}
		index              int = startIndex
	)

	for _, updateField := range manualPaymentEvidenceUpdateFieldList {
		field := string(updateField.manualPaymentEvidenceField)
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

func (repo *RepositoryImpl) BulkCreateManualPaymentEvidence(ctx context.Context, manualPaymentEvidenceList []*model.ManualPaymentEvidence, fieldsInsert ...ManualPaymentEvidenceField) (err error) {
	var (
		fieldsStr                      string
		valueListStr                   []string
		argsList                       []interface{}
		primaryIds                     []model.ManualPaymentEvidencePrimaryID
		manualPaymentEvidenceValueList []model.ManualPaymentEvidence
	)

	if len(fieldsInsert) == 0 {
		selectField := NewManualPaymentEvidenceSelectFields()
		fieldsInsert = selectField.All()
	}
	for _, manualPaymentEvidence := range manualPaymentEvidenceList {

		primaryIds = append(primaryIds, manualPaymentEvidence.ToManualPaymentEvidencePrimaryID())

		manualPaymentEvidenceValueList = append(manualPaymentEvidenceValueList, *manualPaymentEvidence)
	}

	_, notFoundIds, err := repo.IsExistManualPaymentEvidenceByIDs(ctx, primaryIds)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateManualPaymentEvidence] failed checking manualPaymentEvidence whether already exists or not")
		return err
	}
	// ensure that all id is not found
	if len(notFoundIds) != len(primaryIds) {
		existsIds := []model.ManualPaymentEvidencePrimaryID{}
		mapNotFoundIds := map[model.ManualPaymentEvidencePrimaryID]bool{}
		for _, id := range notFoundIds {
			mapNotFoundIds[id] = true
		}
		for _, reqId := range primaryIds {
			if exist := mapNotFoundIds[reqId]; !exist {
				existsIds = append(existsIds, reqId)
			}
		}
		err = failure.Conflict("create", "manualPaymentEvidence", fmt.Sprintf("manualPaymentEvidence with ids '%v' already", fmt.Sprint(existsIds)))
		return err
	}

	fields, values, args := composeInsertFieldsAndParamsManualPaymentEvidence(manualPaymentEvidenceValueList, fieldsInsert...)
	fieldsStr = fields
	valueListStr = append(valueListStr, values...)
	argsList = append(argsList, args...)
	commandQuery := fmt.Sprintf(manualPaymentEvidenceQueries.insertManualPaymentEvidence, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, argsList)
	if err != nil {
		log.Error().Err(err).Msg("[BulkCreateManualPaymentEvidence] failed exec create manualPaymentEvidence query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func (repo *RepositoryImpl) BulkDeleteManualPaymentEvidenceByIDs(ctx context.Context, primaryIDs []model.ManualPaymentEvidencePrimaryID) (err error) {
	exists, notFoundIds, err := repo.IsExistManualPaymentEvidenceByIDs(ctx, primaryIDs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteManualPaymentEvidenceByIDs] failed checking manualPaymentEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("manualPaymentEvidence with id '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"manual_payment_evidence\".\"id\" IN (?)"
	for _, id := range primaryIDs {
		params = append(params, id.Id.String())
	}

	commandQuery := manualPaymentEvidenceQueries.deleteManualPaymentEvidence + " WHERE " + whereQuery

	commandQuery, params, err = sqlx.In(commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteManualPaymentEvidenceByIDs] failed preparing query")
		return failure.InternalError(err)
	}

	query := repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, query, params)
	if err != nil {
		log.Error().Err(err).Msg("[BulkDeleteManualPaymentEvidenceByID] failed exec delete query")
		return failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) IsExistManualPaymentEvidenceByIDs(ctx context.Context, ids []model.ManualPaymentEvidencePrimaryID) (exists bool, notFoundIds []model.ManualPaymentEvidencePrimaryID, err error) {

	var (
		whereQuery string
		params     []interface{}
	)
	whereQuery = "\"manual_payment_evidence\".\"id\" IN (?)"
	for _, id := range ids {
		params = append(params, id.Id.String())
	}

	query := fmt.Sprintf(manualPaymentEvidenceQueries.selectManualPaymentEvidence, " \"id\"   ") + " WHERE " + whereQuery

	query, params, err = sqlx.In(query, params)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistManualPaymentEvidenceByIDs] failed preparing query")
		return false, nil, failure.InternalError(err)
	}

	query = repo.db.Read.Rebind(query)
	var resIds []model.ManualPaymentEvidencePrimaryID
	err = repo.db.Read.SelectContext(ctx, &resIds, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistManualPaymentEvidenceByIDs] failed get ids")
		return false, nil, failure.InternalError(err)
	}
	mapReqIds := map[model.ManualPaymentEvidencePrimaryID]bool{}
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

// BulkUpdateManualPaymentEvidence is used to bulk update manualPaymentEvidence, by default it will update all field
// if want to update specific field, then fill manualPaymentEvidencesMapUpdateFieldsRequest else please fill manualPaymentEvidencesMap
// with specific field you can also use increment
func (r *RepositoryImpl) BulkUpdateManualPaymentEvidence(ctx context.Context, manualPaymentEvidencesMap map[model.ManualPaymentEvidencePrimaryID]*model.ManualPaymentEvidence, manualPaymentEvidencesMapUpdateFieldsRequest map[model.ManualPaymentEvidencePrimaryID]ManualPaymentEvidenceUpdateFieldList) (err error) {
	if len(manualPaymentEvidencesMap) == 0 && len(manualPaymentEvidencesMapUpdateFieldsRequest) == 0 {
		return
	}
	var (
		manualPaymentEvidencesMapUpdateField map[model.ManualPaymentEvidencePrimaryID]ManualPaymentEvidenceUpdateFieldList = map[model.ManualPaymentEvidencePrimaryID]ManualPaymentEvidenceUpdateFieldList{}
		asTableValues                        string                                                                        = "myvalues"
	)

	if len(manualPaymentEvidencesMap) > 0 {
		for id, manualPaymentEvidence := range manualPaymentEvidencesMap {
			if manualPaymentEvidence == nil {
				log.Error().Err(err).Msg("[BulkUpdateManualPaymentEvidence] has nil value")
				return failure.InternalError(errors.New("has nil value"))
			}

			manualPaymentEvidencesMapUpdateField[id] = defaultManualPaymentEvidenceUpdateFields(*manualPaymentEvidence)
		}
	} else {
		manualPaymentEvidencesMapUpdateField = manualPaymentEvidencesMapUpdateFieldsRequest
	}

	ids, paramSetter, paramValues, paramFieldNames, bulkUpdateArgs := composeBulkUpdateManualPaymentEvidenceQuery(manualPaymentEvidencesMapUpdateField, asTableValues)
	exists, notFoundIds, err := r.IsExistManualPaymentEvidenceByIDs(ctx, ids)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateManualPaymentEvidence] failed checking manualPaymentEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("manualPaymentEvidence with ids '%v' not found", fmt.Sprint(notFoundIds)))
		return err
	}
	whereQry := composeManualPaymentEvidenceCompositePrimaryKeyBulkUpdateWhere(ids[0], asTableValues)
	command := fmt.Sprintf(`UPDATE %s 
		SET 
			%s 
		FROM (
			VALUES
			%s
		) AS %s(%s)
		WHERE `+whereQry,
		"\"manual_payment_evidence\"",
		strings.Join(paramSetter, ","),
		strings.Join(paramValues, ","),
		asTableValues,
		strings.Join(paramFieldNames, ","),
	)
	command = r.db.Read.Rebind(command)
	_, err = r.exec(ctx, command, bulkUpdateArgs)
	if err != nil {
		log.Error().Err(err).Msg("[BulkUpdateManualPaymentEvidence] failed exec query")
	}
	return
}

type ManualPaymentEvidenceFieldParameter struct {
	param string
	args  []interface{}
}

func NewManualPaymentEvidenceFieldParameter(param string, args ...interface{}) ManualPaymentEvidenceFieldParameter {
	return ManualPaymentEvidenceFieldParameter{
		param: param,
		args:  args,
	}
}

func composeBulkUpdateManualPaymentEvidenceQuery(mapManualPaymentEvidences map[model.ManualPaymentEvidencePrimaryID]ManualPaymentEvidenceUpdateFieldList, asTableValues string) (primaryIDs []model.ManualPaymentEvidencePrimaryID, paramSetter, paramValues, paramFieldNames []string, args []interface{}) {

	var (
		index int = 1
	)
	// id - fieldName - value
	mapFields := map[model.ManualPaymentEvidencePrimaryID]map[string]interface{}{}
	manualPaymentEvidenceSelectFields := NewManualPaymentEvidenceSelectFields()
	for id, updateFields := range mapManualPaymentEvidences {
		primaryIDs = append(primaryIDs, id)
		mapFields[id] = map[string]interface{}{}
		for _, updateField := range updateFields {
			fieldName := string(updateField.manualPaymentEvidenceField)
			mapFields[id][fieldName] = updateField.value
		}
	}
	if len(primaryIDs) == 0 {
		return
	}
	// compose with same sequence values
	updateFields := mapManualPaymentEvidences[primaryIDs[0]]
	for _, fields := range mapFields {
		indexs := []string{}
		for _, updateField := range updateFields {
			indexs = append(indexs, fmt.Sprintf("$%d::%s", index, GetManualPaymentEvidenceFieldType(updateField.manualPaymentEvidenceField)))
			args = append(args, fields[string(updateField.manualPaymentEvidenceField)])
			index++
		}
		paramValues = append(paramValues, fmt.Sprintf("(%s)", strings.Join(indexs, ",")))
	}

	for _, updateField := range updateFields {
		paramFieldNames = append(paramFieldNames, fmt.Sprintf("\"%s\"", updateField.manualPaymentEvidenceField))
		if updateField.manualPaymentEvidenceField == manualPaymentEvidenceSelectFields.Id() {
			continue
		}
		setter := fmt.Sprintf("\"%s\" = \"%s\".\"%s\"", updateField.manualPaymentEvidenceField, asTableValues, updateField.manualPaymentEvidenceField)
		if updateField.opt.useIncrement {
			setter = fmt.Sprintf("\"%s\" = \"%s\".\"%s\" + \"%s\".\"%s\"", updateField.manualPaymentEvidenceField,
				"\"manual_payment_evidence\"", updateField.manualPaymentEvidenceField,
				asTableValues, updateField.manualPaymentEvidenceField)
		}
		paramSetter = append(paramSetter, setter)
	}

	return
}
func composeManualPaymentEvidenceCompositePrimaryKeyBulkUpdateWhere(primaryIDs model.ManualPaymentEvidencePrimaryID, asTableValue string) (whereQry string) {
	manualPaymentEvidenceSelectFields := NewManualPaymentEvidenceSelectFields()
	var arrWhereQry []string
	id := fmt.Sprintf("\"manual_payment_evidence\".\"id\" = %s.\"id\"::"+GetManualPaymentEvidenceFieldType(manualPaymentEvidenceSelectFields.Id()), asTableValue)
	arrWhereQry = append(arrWhereQry, id)

	qry := fmt.Sprintf("(%s)", strings.Join(arrWhereQry, " AND "))

	return qry
}

func GetManualPaymentEvidenceFieldType(manualPaymentEvidenceField ManualPaymentEvidenceField) string {
	selectManualPaymentEvidenceFields := NewManualPaymentEvidenceSelectFields()
	switch manualPaymentEvidenceField {

	case selectManualPaymentEvidenceFields.Id():
		return "uuid"

	case selectManualPaymentEvidenceFields.PaymentIntentId():
		return "uuid"

	case selectManualPaymentEvidenceFields.PaymentAttemptId():
		return "uuid"

	case selectManualPaymentEvidenceFields.SubmittedBy():
		return "uuid"

	case selectManualPaymentEvidenceFields.EvidenceType():
		return "text"

	case selectManualPaymentEvidenceFields.EvidenceUrl():
		return "text"

	case selectManualPaymentEvidenceFields.Amount():
		return "numeric"

	case selectManualPaymentEvidenceFields.ExpectedAmount():
		return "numeric"

	case selectManualPaymentEvidenceFields.VarianceAmount():
		return "numeric"

	case selectManualPaymentEvidenceFields.VarianceStatus():
		return "text"

	case selectManualPaymentEvidenceFields.Currency():
		return "text"

	case selectManualPaymentEvidenceFields.BankCode():
		return "text"

	case selectManualPaymentEvidenceFields.BankName():
		return "text"

	case selectManualPaymentEvidenceFields.SenderAccountName():
		return "text"

	case selectManualPaymentEvidenceFields.SenderAccountNumberMasked():
		return "text"

	case selectManualPaymentEvidenceFields.Notes():
		return "text"

	case selectManualPaymentEvidenceFields.Status():
		return "manual_evidence_status_enum"

	case selectManualPaymentEvidenceFields.ReviewedBy():
		return "uuid"

	case selectManualPaymentEvidenceFields.ReviewedAt():
		return "timestamptz"

	case selectManualPaymentEvidenceFields.RejectionReason():
		return "text"

	case selectManualPaymentEvidenceFields.PolicyDecision():
		return "text"

	case selectManualPaymentEvidenceFields.Metadata():
		return "jsonb"

	case selectManualPaymentEvidenceFields.MetaCreatedAt():
		return "timestamptz"

	case selectManualPaymentEvidenceFields.MetaCreatedBy():
		return "uuid"

	case selectManualPaymentEvidenceFields.MetaUpdatedAt():
		return "timestamptz"

	case selectManualPaymentEvidenceFields.MetaUpdatedBy():
		return "uuid"

	case selectManualPaymentEvidenceFields.MetaDeletedAt():
		return "timestamptz"

	case selectManualPaymentEvidenceFields.MetaDeletedBy():
		return "uuid"

	default:
		return "unknown"
	}
}

func (repo *RepositoryImpl) CreateManualPaymentEvidence(ctx context.Context, manualPaymentEvidence *model.ManualPaymentEvidence, fieldsInsert ...ManualPaymentEvidenceField) (err error) {
	if len(fieldsInsert) == 0 {
		selectField := NewManualPaymentEvidenceSelectFields()
		fieldsInsert = selectField.All()
	}
	primaryID := model.ManualPaymentEvidencePrimaryID{
		Id: manualPaymentEvidence.Id,
	}
	exists, err := repo.IsExistManualPaymentEvidenceByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[CreateManualPaymentEvidence] failed checking manualPaymentEvidence whether already exists or not")
		return err
	}
	if exists {
		err = failure.Conflict("create", "manualPaymentEvidence", "already exists")
		return
	}
	fieldsStr, valueListStr, args := composeInsertFieldsAndParamsManualPaymentEvidence([]model.ManualPaymentEvidence{*manualPaymentEvidence}, fieldsInsert...)
	commandQuery := fmt.Sprintf(manualPaymentEvidenceQueries.insertManualPaymentEvidence, fieldsStr, strings.Join(valueListStr, ","))

	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateManualPaymentEvidence] failed exec create manualPaymentEvidence query")
		return
	}

	return
}

func (repo *RepositoryImpl) DeleteManualPaymentEvidenceByID(ctx context.Context, primaryID model.ManualPaymentEvidencePrimaryID) (err error) {
	exists, err := repo.IsExistManualPaymentEvidenceByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteManualPaymentEvidenceByID] failed checking manualPaymentEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("manualPaymentEvidence with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	whereQuery, params := composeManualPaymentEvidenceCompositePrimaryKeyWhere([]model.ManualPaymentEvidencePrimaryID{primaryID})
	commandQuery := manualPaymentEvidenceQueries.deleteManualPaymentEvidence + " WHERE " + whereQuery
	commandQuery = repo.db.Read.Rebind(commandQuery)
	_, err = repo.exec(ctx, commandQuery, params)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteManualPaymentEvidenceByID] failed exec delete query")
	}
	return
}

func (repo *RepositoryImpl) ResolveManualPaymentEvidenceByFilter(ctx context.Context, filter model.Filter) (result []model.ManualPaymentEvidenceFilterResult, err error) {
	query, args, err := composeManualPaymentEvidenceFilterQuery(filter)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveManualPaymentEvidenceByFilter] failed compose manualPaymentEvidence filter")
		return
	}
	err = repo.db.Read.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveManualPaymentEvidenceByFilter] failed get manualPaymentEvidence by filter")
		err = failure.InternalError(err)
	}
	return
}

func composeManualPaymentEvidenceFilterSQLExpr(spec model.FilterFieldSpec) (string, error) {
	if spec.Relation == "" {
		return fmt.Sprintf("base.\"%s\"", spec.Column), nil
	}
	joinSpec, found := model.ManualPaymentEvidenceFilterJoins[spec.Relation]
	if !found {
		return "", failure.BadRequestFromString(fmt.Sprintf("join %s is not allowed", spec.Relation))
	}
	return fmt.Sprintf("%s.\"%s\"", joinSpec.Alias, spec.Column), nil
}

func composeManualPaymentEvidenceFilterJoins(requiredJoins map[string]bool) string {
	if len(requiredJoins) == 0 {
		return ""
	}
	joinQueries := []string{}

	if len(joinQueries) == 0 {
		return ""
	}
	return " " + strings.Join(joinQueries, " ")
}

func normalizeManualPaymentEvidenceSortOrder(order string) (string, error) {
	order = strings.ToUpper(strings.TrimSpace(order))
	switch order {
	case model.SortAsc, model.SortDesc:
		return order, nil
	default:
		return "", failure.BadRequestFromString(fmt.Sprintf("sort order %s is not allowed", order))
	}
}

func composeManualPaymentEvidenceFilterSelectColumns(filter model.Filter, isCursorMode bool) (selectColumns []string, err error) {
	selectedColumnCapacity := len(filter.SelectFields) + 1
	if len(filter.SelectFields) == 0 {
		selectedColumnCapacity = 28 + 1
	}
	selectedColumns := make(map[string]struct{}, selectedColumnCapacity)
	addColumn := func(field string) error {
		sourceField, _, _ := model.ParseProjection(field)
		spec, found := model.NewManualPaymentEvidenceFilterFieldSpecFromStr(sourceField)
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
		selectColumns = make([]string, 0, 28+1)
		if _, selected := selectedColumns["id"]; !selected {
			selectColumns = append(selectColumns, "base.\"id\"")
			selectedColumns["id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_intent_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_intent_id\"")
			selectedColumns["payment_intent_id"] = struct{}{}
		}
		if _, selected := selectedColumns["payment_attempt_id"]; !selected {
			selectColumns = append(selectColumns, "base.\"payment_attempt_id\"")
			selectedColumns["payment_attempt_id"] = struct{}{}
		}
		if _, selected := selectedColumns["submitted_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"submitted_by\"")
			selectedColumns["submitted_by"] = struct{}{}
		}
		if _, selected := selectedColumns["evidence_type"]; !selected {
			selectColumns = append(selectColumns, "base.\"evidence_type\"")
			selectedColumns["evidence_type"] = struct{}{}
		}
		if _, selected := selectedColumns["evidence_url"]; !selected {
			selectColumns = append(selectColumns, "base.\"evidence_url\"")
			selectedColumns["evidence_url"] = struct{}{}
		}
		if _, selected := selectedColumns["amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"amount\"")
			selectedColumns["amount"] = struct{}{}
		}
		if _, selected := selectedColumns["expected_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"expected_amount\"")
			selectedColumns["expected_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["variance_amount"]; !selected {
			selectColumns = append(selectColumns, "base.\"variance_amount\"")
			selectedColumns["variance_amount"] = struct{}{}
		}
		if _, selected := selectedColumns["variance_status"]; !selected {
			selectColumns = append(selectColumns, "base.\"variance_status\"")
			selectedColumns["variance_status"] = struct{}{}
		}
		if _, selected := selectedColumns["currency"]; !selected {
			selectColumns = append(selectColumns, "base.\"currency\"")
			selectedColumns["currency"] = struct{}{}
		}
		if _, selected := selectedColumns["bank_code"]; !selected {
			selectColumns = append(selectColumns, "base.\"bank_code\"")
			selectedColumns["bank_code"] = struct{}{}
		}
		if _, selected := selectedColumns["bank_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"bank_name\"")
			selectedColumns["bank_name"] = struct{}{}
		}
		if _, selected := selectedColumns["sender_account_name"]; !selected {
			selectColumns = append(selectColumns, "base.\"sender_account_name\"")
			selectedColumns["sender_account_name"] = struct{}{}
		}
		if _, selected := selectedColumns["sender_account_number_masked"]; !selected {
			selectColumns = append(selectColumns, "base.\"sender_account_number_masked\"")
			selectedColumns["sender_account_number_masked"] = struct{}{}
		}
		if _, selected := selectedColumns["notes"]; !selected {
			selectColumns = append(selectColumns, "base.\"notes\"")
			selectedColumns["notes"] = struct{}{}
		}
		if _, selected := selectedColumns["status"]; !selected {
			selectColumns = append(selectColumns, "base.\"status\"")
			selectedColumns["status"] = struct{}{}
		}
		if _, selected := selectedColumns["reviewed_by"]; !selected {
			selectColumns = append(selectColumns, "base.\"reviewed_by\"")
			selectedColumns["reviewed_by"] = struct{}{}
		}
		if _, selected := selectedColumns["reviewed_at"]; !selected {
			selectColumns = append(selectColumns, "base.\"reviewed_at\"")
			selectedColumns["reviewed_at"] = struct{}{}
		}
		if _, selected := selectedColumns["rejection_reason"]; !selected {
			selectColumns = append(selectColumns, "base.\"rejection_reason\"")
			selectedColumns["rejection_reason"] = struct{}{}
		}
		if _, selected := selectedColumns["policy_decision"]; !selected {
			selectColumns = append(selectColumns, "base.\"policy_decision\"")
			selectedColumns["policy_decision"] = struct{}{}
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

type manualPaymentEvidenceFilterPlaceholder struct {
	index int
}

func (p *manualPaymentEvidenceFilterPlaceholder) Next() string {
	placeholder := fmt.Sprintf("$%d", p.index)
	p.index++
	return placeholder
}

func composeManualPaymentEvidenceFilterPredicate(filterField model.FilterField, placeholders *manualPaymentEvidenceFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
	spec, found := model.NewManualPaymentEvidenceFilterFieldSpecFromStr(filterField.Field)
	if !found || !spec.Filterable {
		return "", failure.BadRequestFromString(fmt.Sprintf("field %s is not filterable", filterField.Field))
	}
	sqlExpr, err := composeManualPaymentEvidenceFilterSQLExpr(spec)
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

func composeManualPaymentEvidenceFilterGroup(group model.FilterGroup, placeholders *manualPaymentEvidenceFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (string, error) {
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
		predicate, err := composeManualPaymentEvidenceFilterPredicate(filterField, placeholders, args, requiredJoins)
		if err != nil {
			return "", err
		}
		if predicate != "" {
			parts = append(parts, predicate)
		}
	}
	for _, child := range group.Groups {
		childQuery, err := composeManualPaymentEvidenceFilterGroup(child, placeholders, args, requiredJoins)
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

func composeManualPaymentEvidenceFilterWhereQueries(filter model.Filter, placeholders *manualPaymentEvidenceFilterPlaceholder, args *[]interface{}, requiredJoins map[string]bool) (whereQueries []string, err error) {
	for _, filterField := range filter.FilterFields {
		predicate, predicateErr := composeManualPaymentEvidenceFilterPredicate(filterField, placeholders, args, requiredJoins)
		if predicateErr != nil {
			err = predicateErr
			return
		}
		if predicate != "" {
			whereQueries = append(whereQueries, predicate)
		}
	}
	if filter.Where != nil {
		groupQuery, groupErr := composeManualPaymentEvidenceFilterGroup(*filter.Where, placeholders, args, requiredJoins)
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

func composeManualPaymentEvidenceFilterQuery(filter model.Filter) (query string, args []interface{}, err error) {
	err = model.ValidateManualPaymentEvidenceFieldNameFilter(filter)
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
			sortOrder, sortErr := normalizeManualPaymentEvidenceSortOrder(filter.Sorts[0].Order)
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

	selectColumns, err := composeManualPaymentEvidenceFilterSelectColumns(filter, isCursorMode)
	if err != nil {
		return
	}
	if isCursorMode {
		selectColumns = append(selectColumns, "0 AS count")
	} else {
		selectColumns = append(selectColumns, "COUNT(*) OVER() AS count")
	}

	placeholders := manualPaymentEvidenceFilterPlaceholder{index: 1}
	whereQueries, err := composeManualPaymentEvidenceFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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
			spec, found := model.NewManualPaymentEvidenceFilterFieldSpecFromStr(sort.Field)
			if !found || !spec.Sortable {
				err = failure.BadRequestFromString(fmt.Sprintf("field %s is not sortable", sort.Field))
				return
			}
			sqlExpr, exprErr := composeManualPaymentEvidenceFilterSQLExpr(spec)
			if exprErr != nil {
				err = exprErr
				return
			}
			if spec.Relation != "" {
				requiredJoins[spec.Relation] = true
			}
			sortOrder, sortErr := normalizeManualPaymentEvidenceSortOrder(sort.Order)
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

	query = fmt.Sprintf("SELECT %s FROM \"manual_payment_evidence\" base%s", strings.Join(selectColumns, ","), composeManualPaymentEvidenceFilterJoins(requiredJoins))
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

func (repo *RepositoryImpl) IsExistManualPaymentEvidenceByID(ctx context.Context, primaryID model.ManualPaymentEvidencePrimaryID) (exists bool, err error) {
	whereQuery, params := composeManualPaymentEvidenceCompositePrimaryKeyWhere([]model.ManualPaymentEvidencePrimaryID{primaryID})
	query := fmt.Sprintf("%s WHERE %s", manualPaymentEvidenceQueries.selectCountManualPaymentEvidence, whereQuery)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &exists, query, params...)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistManualPaymentEvidenceByID] failed get count")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveManualPaymentEvidence(ctx context.Context, selectFields ...ManualPaymentEvidenceField) (manualPaymentEvidenceList model.ManualPaymentEvidenceList, err error) {
	var (
		defaultManualPaymentEvidenceSelectFields = defaultManualPaymentEvidenceSelectFields()
	)
	if len(selectFields) > 0 {
		defaultManualPaymentEvidenceSelectFields = composeManualPaymentEvidenceSelectFields(selectFields...)
	}
	query := fmt.Sprintf(manualPaymentEvidenceQueries.selectManualPaymentEvidence, defaultManualPaymentEvidenceSelectFields)

	err = repo.db.Read.SelectContext(ctx, &manualPaymentEvidenceList, query)
	if err != nil {
		log.Error().Err(err).Msg("[ResolveManualPaymentEvidence] failed get manualPaymentEvidence list")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) ResolveManualPaymentEvidenceByID(ctx context.Context, primaryID model.ManualPaymentEvidencePrimaryID, selectFields ...ManualPaymentEvidenceField) (manualPaymentEvidence model.ManualPaymentEvidence, err error) {
	var (
		defaultManualPaymentEvidenceSelectFields = defaultManualPaymentEvidenceSelectFields()
	)
	if len(selectFields) > 0 {
		defaultManualPaymentEvidenceSelectFields = composeManualPaymentEvidenceSelectFields(selectFields...)
	}
	whereQry, params := composeManualPaymentEvidenceCompositePrimaryKeyWhere([]model.ManualPaymentEvidencePrimaryID{primaryID})
	query := fmt.Sprintf(manualPaymentEvidenceQueries.selectManualPaymentEvidence+" WHERE "+whereQry, defaultManualPaymentEvidenceSelectFields)
	query = repo.db.Read.Rebind(query)
	err = repo.db.Read.GetContext(ctx, &manualPaymentEvidence, query, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = failure.NotFound(fmt.Sprintf("manualPaymentEvidence with id '%s' not found", fmt.Sprint(primaryID)))
			return
		}
		log.Error().Err(err).Msg("[ResolveManualPaymentEvidenceByID] failed get manualPaymentEvidence")
		err = failure.InternalError(err)
	}
	return
}

func (repo *RepositoryImpl) UpdateManualPaymentEvidenceByID(ctx context.Context, primaryID model.ManualPaymentEvidencePrimaryID, manualPaymentEvidence *model.ManualPaymentEvidence, manualPaymentEvidenceUpdateFields ...ManualPaymentEvidenceUpdateField) (err error) {
	exists, err := repo.IsExistManualPaymentEvidenceByID(ctx, primaryID)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateManualPaymentEvidence] failed checking manualPaymentEvidence whether already exists or not")
		return err
	}
	if !exists {
		err = failure.NotFound(fmt.Sprintf("manualPaymentEvidence with id '%s' not found", fmt.Sprint(primaryID)))
		return
	}
	if manualPaymentEvidence == nil {
		if len(manualPaymentEvidenceUpdateFields) == 0 {
			log.Error().Err(err).Msg("[UpdateManualPaymentEvidenceByID] didn't define specific value but has nil value")
			return failure.InternalError(errors.New("didn't define specific value but has nil value"))
		}
		manualPaymentEvidence = &model.ManualPaymentEvidence{}
	}
	var (
		defaultManualPaymentEvidenceUpdateFields = defaultManualPaymentEvidenceUpdateFields(*manualPaymentEvidence)
		tempUpdateField                          ManualPaymentEvidenceUpdateFieldList
		selectFields                             = NewManualPaymentEvidenceSelectFields()
	)
	if len(manualPaymentEvidenceUpdateFields) > 0 {
		for _, updateField := range manualPaymentEvidenceUpdateFields {
			if updateField.manualPaymentEvidenceField == selectFields.Id() {
				continue
			}
			tempUpdateField = append(tempUpdateField, updateField)
		}
		defaultManualPaymentEvidenceUpdateFields = tempUpdateField
	}
	whereQuery, params := composeManualPaymentEvidenceCompositePrimaryKeyWhere([]model.ManualPaymentEvidencePrimaryID{primaryID})
	fields, args := composeUpdateFieldsManualPaymentEvidenceCommand(defaultManualPaymentEvidenceUpdateFields, len(params)+1)
	commandQuery := fmt.Sprintf(manualPaymentEvidenceQueries.updateManualPaymentEvidence+" WHERE "+whereQuery, strings.Join(fields, ","))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	args = append(params, args...)
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateManualPaymentEvidence] error when try to update manualPaymentEvidence by id")
	}
	return err
}

func (repo *RepositoryImpl) UpdateManualPaymentEvidenceByFilter(ctx context.Context, filter model.Filter, manualPaymentEvidenceUpdateFields ...ManualPaymentEvidenceUpdateField) (rowsAffected int64, err error) {
	if len(filter.FilterFields) == 0 && filter.Where == nil {
		err = failure.BadRequestFromString("update by filter requires at least one filter predicate")
		return
	}
	if len(manualPaymentEvidenceUpdateFields) == 0 {
		err = failure.BadRequestFromString("update fields are required")
		return
	}

	var (
		updateFields ManualPaymentEvidenceUpdateFieldList
		selectFields = NewManualPaymentEvidenceSelectFields()
	)
	for _, updateField := range manualPaymentEvidenceUpdateFields {
		if updateField.manualPaymentEvidenceField == selectFields.Id() {
			continue
		}
		updateFields = append(updateFields, updateField)
	}
	if len(updateFields) == 0 {
		err = failure.BadRequestFromString("no mutable update fields provided")
		return
	}

	fields, updateArgs := composeUpdateFieldsManualPaymentEvidenceCommand(updateFields, 1)
	args := append([]interface{}{}, updateArgs...)
	requiredJoins := map[string]bool{}
	placeholders := manualPaymentEvidenceFilterPlaceholder{index: len(updateArgs) + 1}
	whereQueries, err := composeManualPaymentEvidenceFilterWhereQueries(filter, &placeholders, &args, requiredJoins)
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

	commandQuery := fmt.Sprintf("UPDATE \"manual_payment_evidence\" AS base SET %s WHERE %s", strings.Join(fields, ","), strings.Join(whereQueries, " AND "))
	commandQuery = repo.db.Read.Rebind(commandQuery)
	result, err := repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateManualPaymentEvidenceByFilter] error when try to update manualPaymentEvidence by filter")
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("[UpdateManualPaymentEvidenceByFilter] failed get rows affected")
	}
	return
}

var (
	manualPaymentEvidenceQueries = struct {
		selectManualPaymentEvidence      string
		selectCountManualPaymentEvidence string
		deleteManualPaymentEvidence      string
		updateManualPaymentEvidence      string
		insertManualPaymentEvidence      string
	}{
		selectManualPaymentEvidence:      "SELECT %s FROM \"manual_payment_evidence\"",
		selectCountManualPaymentEvidence: "SELECT COUNT(\"id\") FROM \"manual_payment_evidence\"",
		deleteManualPaymentEvidence:      "DELETE FROM \"manual_payment_evidence\"",
		updateManualPaymentEvidence:      "UPDATE \"manual_payment_evidence\" SET %s ",
		insertManualPaymentEvidence:      "INSERT INTO \"manual_payment_evidence\" %s VALUES %s",
	}
)

type ManualPaymentEvidenceRepository interface {
	CreateManualPaymentEvidence(ctx context.Context, manualPaymentEvidence *model.ManualPaymentEvidence, fieldsInsert ...ManualPaymentEvidenceField) error
	BulkCreateManualPaymentEvidence(ctx context.Context, manualPaymentEvidenceList []*model.ManualPaymentEvidence, fieldsInsert ...ManualPaymentEvidenceField) error
	ResolveManualPaymentEvidence(ctx context.Context, selectFields ...ManualPaymentEvidenceField) (model.ManualPaymentEvidenceList, error)
	ResolveManualPaymentEvidenceByID(ctx context.Context, primaryID model.ManualPaymentEvidencePrimaryID, selectFields ...ManualPaymentEvidenceField) (model.ManualPaymentEvidence, error)
	UpdateManualPaymentEvidenceByID(ctx context.Context, id model.ManualPaymentEvidencePrimaryID, manualPaymentEvidence *model.ManualPaymentEvidence, manualPaymentEvidenceUpdateFields ...ManualPaymentEvidenceUpdateField) error
	UpdateManualPaymentEvidenceByFilter(ctx context.Context, filter model.Filter, manualPaymentEvidenceUpdateFields ...ManualPaymentEvidenceUpdateField) (rowsAffected int64, err error)
	BulkUpdateManualPaymentEvidence(ctx context.Context, manualPaymentEvidenceListMap map[model.ManualPaymentEvidencePrimaryID]*model.ManualPaymentEvidence, ManualPaymentEvidencesMapUpdateFieldsRequest map[model.ManualPaymentEvidencePrimaryID]ManualPaymentEvidenceUpdateFieldList) (err error)
	DeleteManualPaymentEvidenceByID(ctx context.Context, id model.ManualPaymentEvidencePrimaryID) error
	BulkDeleteManualPaymentEvidenceByIDs(ctx context.Context, ids []model.ManualPaymentEvidencePrimaryID) error
	ResolveManualPaymentEvidenceByFilter(ctx context.Context, filter model.Filter) (result []model.ManualPaymentEvidenceFilterResult, err error)
	IsExistManualPaymentEvidenceByIDs(ctx context.Context, ids []model.ManualPaymentEvidencePrimaryID) (exists bool, notFoundIds []model.ManualPaymentEvidencePrimaryID, err error)
	IsExistManualPaymentEvidenceByID(ctx context.Context, id model.ManualPaymentEvidencePrimaryID) (exists bool, err error)
}

// intentionally empty: base repository.go defines exported interface and provider; impl adds methods
