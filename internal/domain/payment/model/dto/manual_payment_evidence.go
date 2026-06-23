package dto

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/guregu/null"

	"github.com/gofrs/uuid"

	"github.com/nuriansyah/lokatra-payment/shared/nuuid"

	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type ManualPaymentEvidenceDTOFieldNameType string

type manualPaymentEvidenceDTOFieldName struct {
	Id                        ManualPaymentEvidenceDTOFieldNameType
	PaymentIntentId           ManualPaymentEvidenceDTOFieldNameType
	PaymentAttemptId          ManualPaymentEvidenceDTOFieldNameType
	SubmittedBy               ManualPaymentEvidenceDTOFieldNameType
	EvidenceType              ManualPaymentEvidenceDTOFieldNameType
	EvidenceUrl               ManualPaymentEvidenceDTOFieldNameType
	Amount                    ManualPaymentEvidenceDTOFieldNameType
	ExpectedAmount            ManualPaymentEvidenceDTOFieldNameType
	VarianceAmount            ManualPaymentEvidenceDTOFieldNameType
	VarianceStatus            ManualPaymentEvidenceDTOFieldNameType
	Currency                  ManualPaymentEvidenceDTOFieldNameType
	BankCode                  ManualPaymentEvidenceDTOFieldNameType
	BankName                  ManualPaymentEvidenceDTOFieldNameType
	SenderAccountName         ManualPaymentEvidenceDTOFieldNameType
	SenderAccountNumberMasked ManualPaymentEvidenceDTOFieldNameType
	Notes                     ManualPaymentEvidenceDTOFieldNameType
	Status                    ManualPaymentEvidenceDTOFieldNameType
	ReviewedBy                ManualPaymentEvidenceDTOFieldNameType
	ReviewedAt                ManualPaymentEvidenceDTOFieldNameType
	RejectionReason           ManualPaymentEvidenceDTOFieldNameType
	PolicyDecision            ManualPaymentEvidenceDTOFieldNameType
	Metadata                  ManualPaymentEvidenceDTOFieldNameType
	MetaCreatedAt             ManualPaymentEvidenceDTOFieldNameType
	MetaCreatedBy             ManualPaymentEvidenceDTOFieldNameType
	MetaUpdatedAt             ManualPaymentEvidenceDTOFieldNameType
	MetaUpdatedBy             ManualPaymentEvidenceDTOFieldNameType
	MetaDeletedAt             ManualPaymentEvidenceDTOFieldNameType
	MetaDeletedBy             ManualPaymentEvidenceDTOFieldNameType
}

var ManualPaymentEvidenceDTOFieldName = manualPaymentEvidenceDTOFieldName{
	Id:                        "id",
	PaymentIntentId:           "paymentIntentId",
	PaymentAttemptId:          "paymentAttemptId",
	SubmittedBy:               "submittedBy",
	EvidenceType:              "evidenceType",
	EvidenceUrl:               "evidenceUrl",
	Amount:                    "amount",
	ExpectedAmount:            "expectedAmount",
	VarianceAmount:            "varianceAmount",
	VarianceStatus:            "varianceStatus",
	Currency:                  "currency",
	BankCode:                  "bankCode",
	BankName:                  "bankName",
	SenderAccountName:         "senderAccountName",
	SenderAccountNumberMasked: "senderAccountNumberMasked",
	Notes:                     "notes",
	Status:                    "status",
	ReviewedBy:                "reviewedBy",
	ReviewedAt:                "reviewedAt",
	RejectionReason:           "rejectionReason",
	PolicyDecision:            "policyDecision",
	Metadata:                  "metadata",
	MetaCreatedAt:             "metaCreatedAt",
	MetaCreatedBy:             "metaCreatedBy",
	MetaUpdatedAt:             "metaUpdatedAt",
	MetaUpdatedBy:             "metaUpdatedBy",
	MetaDeletedAt:             "metaDeletedAt",
	MetaDeletedBy:             "metaDeletedBy",
}

func transformManualPaymentEvidenceDTOFieldNameFromStr(field string) (dbField string, found bool) {
	switch field {

	case string(ManualPaymentEvidenceDTOFieldName.Id):
		return string(model.ManualPaymentEvidenceDBFieldName.Id), true

	case string(ManualPaymentEvidenceDTOFieldName.PaymentIntentId):
		return string(model.ManualPaymentEvidenceDBFieldName.PaymentIntentId), true

	case string(ManualPaymentEvidenceDTOFieldName.PaymentAttemptId):
		return string(model.ManualPaymentEvidenceDBFieldName.PaymentAttemptId), true

	case string(ManualPaymentEvidenceDTOFieldName.SubmittedBy):
		return string(model.ManualPaymentEvidenceDBFieldName.SubmittedBy), true

	case string(ManualPaymentEvidenceDTOFieldName.EvidenceType):
		return string(model.ManualPaymentEvidenceDBFieldName.EvidenceType), true

	case string(ManualPaymentEvidenceDTOFieldName.EvidenceUrl):
		return string(model.ManualPaymentEvidenceDBFieldName.EvidenceUrl), true

	case string(ManualPaymentEvidenceDTOFieldName.Amount):
		return string(model.ManualPaymentEvidenceDBFieldName.Amount), true

	case string(ManualPaymentEvidenceDTOFieldName.ExpectedAmount):
		return string(model.ManualPaymentEvidenceDBFieldName.ExpectedAmount), true

	case string(ManualPaymentEvidenceDTOFieldName.VarianceAmount):
		return string(model.ManualPaymentEvidenceDBFieldName.VarianceAmount), true

	case string(ManualPaymentEvidenceDTOFieldName.VarianceStatus):
		return string(model.ManualPaymentEvidenceDBFieldName.VarianceStatus), true

	case string(ManualPaymentEvidenceDTOFieldName.Currency):
		return string(model.ManualPaymentEvidenceDBFieldName.Currency), true

	case string(ManualPaymentEvidenceDTOFieldName.BankCode):
		return string(model.ManualPaymentEvidenceDBFieldName.BankCode), true

	case string(ManualPaymentEvidenceDTOFieldName.BankName):
		return string(model.ManualPaymentEvidenceDBFieldName.BankName), true

	case string(ManualPaymentEvidenceDTOFieldName.SenderAccountName):
		return string(model.ManualPaymentEvidenceDBFieldName.SenderAccountName), true

	case string(ManualPaymentEvidenceDTOFieldName.SenderAccountNumberMasked):
		return string(model.ManualPaymentEvidenceDBFieldName.SenderAccountNumberMasked), true

	case string(ManualPaymentEvidenceDTOFieldName.Notes):
		return string(model.ManualPaymentEvidenceDBFieldName.Notes), true

	case string(ManualPaymentEvidenceDTOFieldName.Status):
		return string(model.ManualPaymentEvidenceDBFieldName.Status), true

	case string(ManualPaymentEvidenceDTOFieldName.ReviewedBy):
		return string(model.ManualPaymentEvidenceDBFieldName.ReviewedBy), true

	case string(ManualPaymentEvidenceDTOFieldName.ReviewedAt):
		return string(model.ManualPaymentEvidenceDBFieldName.ReviewedAt), true

	case string(ManualPaymentEvidenceDTOFieldName.RejectionReason):
		return string(model.ManualPaymentEvidenceDBFieldName.RejectionReason), true

	case string(ManualPaymentEvidenceDTOFieldName.PolicyDecision):
		return string(model.ManualPaymentEvidenceDBFieldName.PolicyDecision), true

	case string(ManualPaymentEvidenceDTOFieldName.Metadata):
		return string(model.ManualPaymentEvidenceDBFieldName.Metadata), true

	case string(ManualPaymentEvidenceDTOFieldName.MetaCreatedAt):
		return string(model.ManualPaymentEvidenceDBFieldName.MetaCreatedAt), true

	case string(ManualPaymentEvidenceDTOFieldName.MetaCreatedBy):
		return string(model.ManualPaymentEvidenceDBFieldName.MetaCreatedBy), true

	case string(ManualPaymentEvidenceDTOFieldName.MetaUpdatedAt):
		return string(model.ManualPaymentEvidenceDBFieldName.MetaUpdatedAt), true

	case string(ManualPaymentEvidenceDTOFieldName.MetaUpdatedBy):
		return string(model.ManualPaymentEvidenceDBFieldName.MetaUpdatedBy), true

	case string(ManualPaymentEvidenceDTOFieldName.MetaDeletedAt):
		return string(model.ManualPaymentEvidenceDBFieldName.MetaDeletedAt), true

	case string(ManualPaymentEvidenceDTOFieldName.MetaDeletedBy):
		return string(model.ManualPaymentEvidenceDBFieldName.MetaDeletedBy), true

	}
	if _, found := model.NewManualPaymentEvidenceFilterFieldSpecFromStr(field); found {
		return field, true
	}
	return "", false
}

func isManualPaymentEvidenceBaseFilterField(field string) bool {
	spec, found := model.NewManualPaymentEvidenceFilterFieldSpecFromStr(field)
	return found && spec.Relation == ""
}

func composeManualPaymentEvidenceProjection(source, output string, explicitAlias bool) string {
	if explicitAlias {
		return source + " as " + output
	}
	return source
}

func validateManualPaymentEvidenceProjectionOutputPath(path string) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return failure.BadRequest(fmt.Errorf("field alias cannot be empty"))
	}
	if !strings.Contains(path, ".") {
		return nil
	}
	for _, part := range strings.Split(path, ".") {
		if strings.TrimSpace(part) == "" {
			return failure.BadRequest(fmt.Errorf("field alias %s is invalid", path))
		}
	}
	return nil
}

func transformManualPaymentEvidenceFilterGroupFieldNames(group *model.FilterGroup) (err error) {
	for index, field := range group.FilterFields {
		dbField, exist := transformManualPaymentEvidenceDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		group.FilterFields[index].Field = dbField
	}
	for index := range group.Groups {
		err = transformManualPaymentEvidenceFilterGroupFieldNames(&group.Groups[index])
		if err != nil {
			return
		}
	}
	return
}

func ValidateAndTransformManualPaymentEvidenceFieldNameFilter(filter *model.Filter) (err error) {
	for index, selectField := range filter.SelectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		dbField, exist := transformManualPaymentEvidenceDTOFieldNameFromStr(sourceField)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sourceField))
			return
		}
		if !isManualPaymentEvidenceBaseFilterField(dbField) {
			err = failure.BadRequest(fmt.Errorf("field %s cannot be selected in typed filter response", sourceField))
			return
		}
		if !explicitAlias && dbField != sourceField {
			outputField = sourceField
			explicitAlias = true
		}
		if explicitAlias {
			if err = validateManualPaymentEvidenceProjectionOutputPath(outputField); err != nil {
				return
			}
		}
		filter.SelectFields[index] = composeManualPaymentEvidenceProjection(dbField, outputField, explicitAlias)
	}
	for index, sort := range filter.Sorts {
		dbField, exist := transformManualPaymentEvidenceDTOFieldNameFromStr(sort.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", sort.Field))
			return
		}
		filter.Sorts[index].Field = dbField
	}
	for index, field := range filter.FilterFields {
		dbField, exist := transformManualPaymentEvidenceDTOFieldNameFromStr(field.Field)
		if !exist {
			err = failure.BadRequest(fmt.Errorf("field %s is not found", field.Field))
			return
		}
		filter.FilterFields[index].Field = dbField
	}
	if filter.Where != nil {
		err = transformManualPaymentEvidenceFilterGroupFieldNames(filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func SetDefaultManualPaymentEvidenceFilter(filter *model.Filter) {
	if filter.Pagination.Strategy == "" {
		if filter.Pagination.Page > 0 {
			filter.Pagination.Strategy = model.PaginationStrategyOffset
		} else {
			filter.Pagination.Strategy = model.DefaultPaginationStrategy
		}
	}

	if filter.Pagination.PageSize <= 0 {
		filter.Pagination.PageSize = model.DefaultPageSize
	}

	if filter.Pagination.PageSize > model.MaxPageSize {
		filter.Pagination.PageSize = model.MaxPageSize
	}

	if filter.Pagination.Strategy == model.PaginationStrategyOffset && filter.Pagination.Page <= 0 {
		filter.Pagination.Page = 1
	}

	if filter.Pagination.Direction == "" {
		filter.Pagination.Direction = model.CursorDirectionNext
	}

	if len(filter.Sorts) == 0 {
		filter.Sorts = append(filter.Sorts, model.Sort{
			Field: string(ManualPaymentEvidenceDTOFieldName.Id),
			Order: model.SortAsc,
		})
	}
}

type ManualPaymentEvidenceSelectableResponse map[string]interface{}
type ManualPaymentEvidenceSelectableListResponse []*ManualPaymentEvidenceSelectableResponse

func assignManualPaymentEvidenceNestedValue(out map[string]interface{}, path string, value interface{}) {
	parts := strings.Split(path, ".")
	current := out
	for _, part := range parts[:len(parts)-1] {
		next, ok := current[part].(map[string]interface{})
		if !ok {
			next = map[string]interface{}{}
			current[part] = next
		}
		current = next
	}
	current[parts[len(parts)-1]] = value
}

func setManualPaymentEvidenceSelectableValue(out ManualPaymentEvidenceSelectableResponse, key string, value interface{}, explicitAlias bool) {
	if explicitAlias && strings.Contains(key, ".") {
		assignManualPaymentEvidenceNestedValue(out, key, value)
		return
	}
	out[key] = value
}

func NewManualPaymentEvidenceSelectableResponse(manualPaymentEvidence model.ManualPaymentEvidence, filter model.Filter) ManualPaymentEvidenceSelectableResponse {
	selectFields := filter.SelectFields
	if len(selectFields) == 0 {
		selectFields = append(selectFields,
			string(model.ManualPaymentEvidenceDBFieldName.Id),
			string(model.ManualPaymentEvidenceDBFieldName.PaymentIntentId),
			string(model.ManualPaymentEvidenceDBFieldName.PaymentAttemptId),
			string(model.ManualPaymentEvidenceDBFieldName.SubmittedBy),
			string(model.ManualPaymentEvidenceDBFieldName.EvidenceType),
			string(model.ManualPaymentEvidenceDBFieldName.EvidenceUrl),
			string(model.ManualPaymentEvidenceDBFieldName.Amount),
			string(model.ManualPaymentEvidenceDBFieldName.ExpectedAmount),
			string(model.ManualPaymentEvidenceDBFieldName.VarianceAmount),
			string(model.ManualPaymentEvidenceDBFieldName.VarianceStatus),
			string(model.ManualPaymentEvidenceDBFieldName.Currency),
			string(model.ManualPaymentEvidenceDBFieldName.BankCode),
			string(model.ManualPaymentEvidenceDBFieldName.BankName),
			string(model.ManualPaymentEvidenceDBFieldName.SenderAccountName),
			string(model.ManualPaymentEvidenceDBFieldName.SenderAccountNumberMasked),
			string(model.ManualPaymentEvidenceDBFieldName.Notes),
			string(model.ManualPaymentEvidenceDBFieldName.Status),
			string(model.ManualPaymentEvidenceDBFieldName.ReviewedBy),
			string(model.ManualPaymentEvidenceDBFieldName.ReviewedAt),
			string(model.ManualPaymentEvidenceDBFieldName.RejectionReason),
			string(model.ManualPaymentEvidenceDBFieldName.PolicyDecision),
			string(model.ManualPaymentEvidenceDBFieldName.Metadata),
			string(model.ManualPaymentEvidenceDBFieldName.MetaCreatedAt),
			string(model.ManualPaymentEvidenceDBFieldName.MetaCreatedBy),
			string(model.ManualPaymentEvidenceDBFieldName.MetaUpdatedAt),
			string(model.ManualPaymentEvidenceDBFieldName.MetaUpdatedBy),
			string(model.ManualPaymentEvidenceDBFieldName.MetaDeletedAt),
			string(model.ManualPaymentEvidenceDBFieldName.MetaDeletedBy),
		)
	}
	manualPaymentEvidenceSelectableResponse := ManualPaymentEvidenceSelectableResponse{}
	for _, selectField := range selectFields {
		sourceField, outputField, explicitAlias := model.ParseProjection(selectField)
		switch sourceField {

		case string(model.ManualPaymentEvidenceDBFieldName.Id):
			key := string(ManualPaymentEvidenceDTOFieldName.Id)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.Id, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.PaymentIntentId):
			key := string(ManualPaymentEvidenceDTOFieldName.PaymentIntentId)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.PaymentIntentId, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.PaymentAttemptId):
			key := string(ManualPaymentEvidenceDTOFieldName.PaymentAttemptId)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.PaymentAttemptId.UUID, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.SubmittedBy):
			key := string(ManualPaymentEvidenceDTOFieldName.SubmittedBy)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.SubmittedBy, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.EvidenceType):
			key := string(ManualPaymentEvidenceDTOFieldName.EvidenceType)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.EvidenceType, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.EvidenceUrl):
			key := string(ManualPaymentEvidenceDTOFieldName.EvidenceUrl)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.EvidenceUrl.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.Amount):
			key := string(ManualPaymentEvidenceDTOFieldName.Amount)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.Amount, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.ExpectedAmount):
			key := string(ManualPaymentEvidenceDTOFieldName.ExpectedAmount)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.ExpectedAmount, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.VarianceAmount):
			key := string(ManualPaymentEvidenceDTOFieldName.VarianceAmount)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.VarianceAmount, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.VarianceStatus):
			key := string(ManualPaymentEvidenceDTOFieldName.VarianceStatus)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.VarianceStatus, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.Currency):
			key := string(ManualPaymentEvidenceDTOFieldName.Currency)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.Currency, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.BankCode):
			key := string(ManualPaymentEvidenceDTOFieldName.BankCode)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.BankCode.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.BankName):
			key := string(ManualPaymentEvidenceDTOFieldName.BankName)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.BankName.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.SenderAccountName):
			key := string(ManualPaymentEvidenceDTOFieldName.SenderAccountName)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.SenderAccountName.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.SenderAccountNumberMasked):
			key := string(ManualPaymentEvidenceDTOFieldName.SenderAccountNumberMasked)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.SenderAccountNumberMasked.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.Notes):
			key := string(ManualPaymentEvidenceDTOFieldName.Notes)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.Notes.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.Status):
			key := string(ManualPaymentEvidenceDTOFieldName.Status)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, model.ManualEvidenceStatus(manualPaymentEvidence.Status), explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.ReviewedBy):
			key := string(ManualPaymentEvidenceDTOFieldName.ReviewedBy)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.ReviewedBy.UUID, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.ReviewedAt):
			key := string(ManualPaymentEvidenceDTOFieldName.ReviewedAt)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.ReviewedAt.Time, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.RejectionReason):
			key := string(ManualPaymentEvidenceDTOFieldName.RejectionReason)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.RejectionReason.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.PolicyDecision):
			key := string(ManualPaymentEvidenceDTOFieldName.PolicyDecision)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.PolicyDecision.String, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.Metadata):
			key := string(ManualPaymentEvidenceDTOFieldName.Metadata)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.Metadata, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.MetaCreatedAt):
			key := string(ManualPaymentEvidenceDTOFieldName.MetaCreatedAt)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.MetaCreatedAt, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.MetaCreatedBy):
			key := string(ManualPaymentEvidenceDTOFieldName.MetaCreatedBy)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.MetaCreatedBy, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.MetaUpdatedAt):
			key := string(ManualPaymentEvidenceDTOFieldName.MetaUpdatedAt)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.MetaUpdatedAt, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.MetaUpdatedBy):
			key := string(ManualPaymentEvidenceDTOFieldName.MetaUpdatedBy)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.MetaUpdatedBy, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.MetaDeletedAt):
			key := string(ManualPaymentEvidenceDTOFieldName.MetaDeletedAt)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.MetaDeletedAt.Time, explicitAlias)

		case string(model.ManualPaymentEvidenceDBFieldName.MetaDeletedBy):
			key := string(ManualPaymentEvidenceDTOFieldName.MetaDeletedBy)
			if explicitAlias {
				key = outputField
			}
			setManualPaymentEvidenceSelectableValue(manualPaymentEvidenceSelectableResponse, key, manualPaymentEvidence.MetaDeletedBy, explicitAlias)

		}
	}
	return manualPaymentEvidenceSelectableResponse
}

func NewManualPaymentEvidenceListResponseFromFilterResult(result []model.ManualPaymentEvidenceFilterResult, filter model.Filter) ManualPaymentEvidenceSelectableListResponse {
	dtoManualPaymentEvidenceListResponse := ManualPaymentEvidenceSelectableListResponse{}
	for _, row := range result {
		dtoManualPaymentEvidenceResponse := NewManualPaymentEvidenceSelectableResponse(row.ManualPaymentEvidence, filter)
		dtoManualPaymentEvidenceListResponse = append(dtoManualPaymentEvidenceListResponse, &dtoManualPaymentEvidenceResponse)
	}
	return dtoManualPaymentEvidenceListResponse
}

type ManualPaymentEvidenceFilterResponse struct {
	Metadata Metadata                                    `json:"metadata"`
	Data     ManualPaymentEvidenceSelectableListResponse `json:"data"`
}

func reverseManualPaymentEvidenceFilterResults(result []model.ManualPaymentEvidenceFilterResult) {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
}

func NewManualPaymentEvidenceFilterResponse(result []model.ManualPaymentEvidenceFilterResult, filter model.Filter) (resp ManualPaymentEvidenceFilterResponse) {
	resp.Metadata.Strategy = filter.Pagination.Strategy
	resp.Metadata.PageSize = filter.Pagination.PageSize
	resp.Metadata.Page = filter.Pagination.Page
	dataResult := result

	if filter.Pagination.IsCursorMode() && filter.Pagination.PageSize > 0 {
		if len(dataResult) > filter.Pagination.PageSize {
			resp.Metadata.HasMore = true
			if filter.Pagination.Direction == model.CursorDirectionPrev {
				resp.Metadata.HasPrev = true
			} else {
				resp.Metadata.HasNext = true
			}
			dataResult = dataResult[:filter.Pagination.PageSize]
		}
		if filter.Pagination.Direction == model.CursorDirectionPrev {
			reverseManualPaymentEvidenceFilterResults(dataResult)
			if filter.Pagination.Cursor != nil {
				resp.Metadata.HasNext = true
			}
		} else if filter.Pagination.Cursor != nil {
			resp.Metadata.HasPrev = true
		}
		if len(dataResult) > 0 {
			resp.Metadata.NextCursor = dataResult[len(dataResult)-1].Id
			resp.Metadata.PrevCursor = dataResult[0].Id
		}
		if resp.Metadata.Page <= 0 {
			resp.Metadata.Page = 1
		}
	} else {
		if len(dataResult) > 0 {
			resp.Metadata.TotalData = dataResult[0].FilterCount
			resp.Metadata.TotalPage = int(math.Ceil(float64(resp.Metadata.TotalData) / float64(filter.Pagination.PageSize)))
			resp.Metadata.HasPrev = filter.Pagination.Page > 1
			resp.Metadata.HasNext = filter.Pagination.Page < resp.Metadata.TotalPage
			resp.Metadata.HasMore = resp.Metadata.HasNext
		}
	}

	resp.Data = NewManualPaymentEvidenceListResponseFromFilterResult(dataResult, filter)
	return resp
}

type ManualPaymentEvidenceCreateRequest struct {
	PaymentIntentId           uuid.UUID                  `json:"paymentIntentId"`
	PaymentAttemptId          uuid.UUID                  `json:"paymentAttemptId"`
	SubmittedBy               uuid.UUID                  `json:"submittedBy"`
	EvidenceType              string                     `json:"evidenceType"`
	EvidenceUrl               string                     `json:"evidenceUrl"`
	Amount                    decimal.Decimal            `json:"amount"`
	ExpectedAmount            decimal.Decimal            `json:"expectedAmount"`
	VarianceAmount            decimal.Decimal            `json:"varianceAmount"`
	VarianceStatus            string                     `json:"varianceStatus"`
	Currency                  string                     `json:"currency"`
	BankCode                  string                     `json:"bankCode"`
	BankName                  string                     `json:"bankName"`
	SenderAccountName         string                     `json:"senderAccountName"`
	SenderAccountNumberMasked string                     `json:"senderAccountNumberMasked"`
	Notes                     string                     `json:"notes"`
	Status                    model.ManualEvidenceStatus `json:"status" example:"submitted" enums:"submitted,under_review,approved,rejected"`
	ReviewedBy                uuid.UUID                  `json:"reviewedBy"`
	ReviewedAt                time.Time                  `json:"reviewedAt"`
	RejectionReason           string                     `json:"rejectionReason"`
	PolicyDecision            string                     `json:"policyDecision"`
	Metadata                  json.RawMessage            `json:"metadata"`
}

func (d *ManualPaymentEvidenceCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d *ManualPaymentEvidenceCreateRequest) ToModel() model.ManualPaymentEvidence {
	id, _ := uuid.NewV4()
	return model.ManualPaymentEvidence{
		Id:                        id,
		PaymentIntentId:           d.PaymentIntentId,
		PaymentAttemptId:          nuuid.From(d.PaymentAttemptId),
		SubmittedBy:               d.SubmittedBy,
		EvidenceType:              d.EvidenceType,
		EvidenceUrl:               null.StringFrom(d.EvidenceUrl),
		Amount:                    d.Amount,
		ExpectedAmount:            d.ExpectedAmount,
		VarianceAmount:            d.VarianceAmount,
		VarianceStatus:            d.VarianceStatus,
		Currency:                  d.Currency,
		BankCode:                  null.StringFrom(d.BankCode),
		BankName:                  null.StringFrom(d.BankName),
		SenderAccountName:         null.StringFrom(d.SenderAccountName),
		SenderAccountNumberMasked: null.StringFrom(d.SenderAccountNumberMasked),
		Notes:                     null.StringFrom(d.Notes),
		Status:                    d.Status,
		ReviewedBy:                nuuid.From(d.ReviewedBy),
		ReviewedAt:                null.TimeFrom(d.ReviewedAt),
		RejectionReason:           null.StringFrom(d.RejectionReason),
		PolicyDecision:            null.StringFrom(d.PolicyDecision),
		Metadata:                  d.Metadata,
	}
}

type ManualPaymentEvidenceListCreateRequest []*ManualPaymentEvidenceCreateRequest

func (d ManualPaymentEvidenceListCreateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, manualPaymentEvidence := range d {
		err = validator.Struct(manualPaymentEvidence)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ManualPaymentEvidenceListCreateRequest) ToModelList() []model.ManualPaymentEvidence {
	out := make([]model.ManualPaymentEvidence, len(d))
	for i, item := range d {
		out[i] = item.ToModel()
	}
	return out
}

type ManualPaymentEvidenceUpdateRequest struct {
	PaymentIntentId           uuid.UUID                  `json:"paymentIntentId"`
	PaymentAttemptId          uuid.UUID                  `json:"paymentAttemptId"`
	SubmittedBy               uuid.UUID                  `json:"submittedBy"`
	EvidenceType              string                     `json:"evidenceType"`
	EvidenceUrl               string                     `json:"evidenceUrl"`
	Amount                    decimal.Decimal            `json:"amount"`
	ExpectedAmount            decimal.Decimal            `json:"expectedAmount"`
	VarianceAmount            decimal.Decimal            `json:"varianceAmount"`
	VarianceStatus            string                     `json:"varianceStatus"`
	Currency                  string                     `json:"currency"`
	BankCode                  string                     `json:"bankCode"`
	BankName                  string                     `json:"bankName"`
	SenderAccountName         string                     `json:"senderAccountName"`
	SenderAccountNumberMasked string                     `json:"senderAccountNumberMasked"`
	Notes                     string                     `json:"notes"`
	Status                    model.ManualEvidenceStatus `json:"status" example:"submitted" enums:"submitted,under_review,approved,rejected"`
	ReviewedBy                uuid.UUID                  `json:"reviewedBy"`
	ReviewedAt                time.Time                  `json:"reviewedAt"`
	RejectionReason           string                     `json:"rejectionReason"`
	PolicyDecision            string                     `json:"policyDecision"`
	Metadata                  json.RawMessage            `json:"metadata"`
}

func (d *ManualPaymentEvidenceUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}

func (d ManualPaymentEvidenceUpdateRequest) ToModel() model.ManualPaymentEvidence {
	return model.ManualPaymentEvidence{
		PaymentIntentId:           d.PaymentIntentId,
		PaymentAttemptId:          nuuid.From(d.PaymentAttemptId),
		SubmittedBy:               d.SubmittedBy,
		EvidenceType:              d.EvidenceType,
		EvidenceUrl:               null.StringFrom(d.EvidenceUrl),
		Amount:                    d.Amount,
		ExpectedAmount:            d.ExpectedAmount,
		VarianceAmount:            d.VarianceAmount,
		VarianceStatus:            d.VarianceStatus,
		Currency:                  d.Currency,
		BankCode:                  null.StringFrom(d.BankCode),
		BankName:                  null.StringFrom(d.BankName),
		SenderAccountName:         null.StringFrom(d.SenderAccountName),
		SenderAccountNumberMasked: null.StringFrom(d.SenderAccountNumberMasked),
		Notes:                     null.StringFrom(d.Notes),
		Status:                    d.Status,
		ReviewedBy:                nuuid.From(d.ReviewedBy),
		ReviewedAt:                null.TimeFrom(d.ReviewedAt),
		RejectionReason:           null.StringFrom(d.RejectionReason),
		PolicyDecision:            null.StringFrom(d.PolicyDecision),
		Metadata:                  d.Metadata,
	}
}

type ManualPaymentEvidenceBulkUpdateRequest struct {
	Id                        uuid.UUID                  `json:"id"`
	PaymentIntentId           uuid.UUID                  `json:"paymentIntentId"`
	PaymentAttemptId          uuid.UUID                  `json:"paymentAttemptId"`
	SubmittedBy               uuid.UUID                  `json:"submittedBy"`
	EvidenceType              string                     `json:"evidenceType"`
	EvidenceUrl               string                     `json:"evidenceUrl"`
	Amount                    decimal.Decimal            `json:"amount"`
	ExpectedAmount            decimal.Decimal            `json:"expectedAmount"`
	VarianceAmount            decimal.Decimal            `json:"varianceAmount"`
	VarianceStatus            string                     `json:"varianceStatus"`
	Currency                  string                     `json:"currency"`
	BankCode                  string                     `json:"bankCode"`
	BankName                  string                     `json:"bankName"`
	SenderAccountName         string                     `json:"senderAccountName"`
	SenderAccountNumberMasked string                     `json:"senderAccountNumberMasked"`
	Notes                     string                     `json:"notes"`
	Status                    model.ManualEvidenceStatus `json:"status" example:"submitted" enums:"submitted,under_review,approved,rejected"`
	ReviewedBy                uuid.UUID                  `json:"reviewedBy"`
	ReviewedAt                time.Time                  `json:"reviewedAt"`
	RejectionReason           string                     `json:"rejectionReason"`
	PolicyDecision            string                     `json:"policyDecision"`
	Metadata                  json.RawMessage            `json:"metadata"`
}

func (d ManualPaymentEvidenceBulkUpdateRequest) PrimaryID() ManualPaymentEvidencePrimaryID {
	return ManualPaymentEvidencePrimaryID{
		Id: d.Id,
	}
}

type ManualPaymentEvidenceListBulkUpdateRequest []*ManualPaymentEvidenceBulkUpdateRequest

func (d ManualPaymentEvidenceListBulkUpdateRequest) Validate() (err error) {
	validator := shared.GetValidator()
	for _, manualPaymentEvidence := range d {
		err = validator.Struct(manualPaymentEvidence)
		if err != nil {
			return
		}
	}
	return nil
}

func (d ManualPaymentEvidenceBulkUpdateRequest) ToModel() model.ManualPaymentEvidence {
	return model.ManualPaymentEvidence{
		Id:                        d.Id,
		PaymentIntentId:           d.PaymentIntentId,
		PaymentAttemptId:          nuuid.From(d.PaymentAttemptId),
		SubmittedBy:               d.SubmittedBy,
		EvidenceType:              d.EvidenceType,
		EvidenceUrl:               null.StringFrom(d.EvidenceUrl),
		Amount:                    d.Amount,
		ExpectedAmount:            d.ExpectedAmount,
		VarianceAmount:            d.VarianceAmount,
		VarianceStatus:            d.VarianceStatus,
		Currency:                  d.Currency,
		BankCode:                  null.StringFrom(d.BankCode),
		BankName:                  null.StringFrom(d.BankName),
		SenderAccountName:         null.StringFrom(d.SenderAccountName),
		SenderAccountNumberMasked: null.StringFrom(d.SenderAccountNumberMasked),
		Notes:                     null.StringFrom(d.Notes),
		Status:                    d.Status,
		ReviewedBy:                nuuid.From(d.ReviewedBy),
		ReviewedAt:                null.TimeFrom(d.ReviewedAt),
		RejectionReason:           null.StringFrom(d.RejectionReason),
		PolicyDecision:            null.StringFrom(d.PolicyDecision),
		Metadata:                  d.Metadata,
	}
}

type ManualPaymentEvidenceResponse struct {
	Id                        uuid.UUID                  `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentIntentId           uuid.UUID                  `json:"paymentIntentId" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	PaymentAttemptId          uuid.UUID                  `json:"paymentAttemptId" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	SubmittedBy               uuid.UUID                  `json:"submittedBy" validate:"required,uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	EvidenceType              string                     `json:"evidenceType"`
	EvidenceUrl               string                     `json:"evidenceUrl" validate:"url"`
	Amount                    decimal.Decimal            `json:"amount" validate:"required" format:"decimal" example:"100.50"`
	ExpectedAmount            decimal.Decimal            `json:"expectedAmount" validate:"required" format:"decimal" example:"100.50"`
	VarianceAmount            decimal.Decimal            `json:"varianceAmount" format:"decimal" example:"100.50"`
	VarianceStatus            string                     `json:"varianceStatus"`
	Currency                  string                     `json:"currency"`
	BankCode                  string                     `json:"bankCode"`
	BankName                  string                     `json:"bankName"`
	SenderAccountName         string                     `json:"senderAccountName"`
	SenderAccountNumberMasked string                     `json:"senderAccountNumberMasked"`
	Notes                     string                     `json:"notes"`
	Status                    model.ManualEvidenceStatus `json:"status" validate:"oneof=submitted under_review approved rejected" enums:"submitted,under_review,approved,rejected"`
	ReviewedBy                uuid.UUID                  `json:"reviewedBy" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ReviewedAt                time.Time                  `json:"reviewedAt" format:"date-time" example:"2024-01-01T00:00:00Z"`
	RejectionReason           string                     `json:"rejectionReason"`
	PolicyDecision            string                     `json:"policyDecision"`
	Metadata                  json.RawMessage            `json:"metadata" swaggertype:"object"`
}

func NewManualPaymentEvidenceResponse(manualPaymentEvidence model.ManualPaymentEvidence) ManualPaymentEvidenceResponse {
	return ManualPaymentEvidenceResponse{
		Id:                        manualPaymentEvidence.Id,
		PaymentIntentId:           manualPaymentEvidence.PaymentIntentId,
		PaymentAttemptId:          manualPaymentEvidence.PaymentAttemptId.UUID,
		SubmittedBy:               manualPaymentEvidence.SubmittedBy,
		EvidenceType:              manualPaymentEvidence.EvidenceType,
		EvidenceUrl:               manualPaymentEvidence.EvidenceUrl.String,
		Amount:                    manualPaymentEvidence.Amount,
		ExpectedAmount:            manualPaymentEvidence.ExpectedAmount,
		VarianceAmount:            manualPaymentEvidence.VarianceAmount,
		VarianceStatus:            manualPaymentEvidence.VarianceStatus,
		Currency:                  manualPaymentEvidence.Currency,
		BankCode:                  manualPaymentEvidence.BankCode.String,
		BankName:                  manualPaymentEvidence.BankName.String,
		SenderAccountName:         manualPaymentEvidence.SenderAccountName.String,
		SenderAccountNumberMasked: manualPaymentEvidence.SenderAccountNumberMasked.String,
		Notes:                     manualPaymentEvidence.Notes.String,
		Status:                    model.ManualEvidenceStatus(manualPaymentEvidence.Status),
		ReviewedBy:                manualPaymentEvidence.ReviewedBy.UUID,
		ReviewedAt:                manualPaymentEvidence.ReviewedAt.Time,
		RejectionReason:           manualPaymentEvidence.RejectionReason.String,
		PolicyDecision:            manualPaymentEvidence.PolicyDecision.String,
		Metadata:                  manualPaymentEvidence.Metadata,
	}
}

type ManualPaymentEvidenceListResponse []*ManualPaymentEvidenceResponse

func NewManualPaymentEvidenceListResponse(manualPaymentEvidenceList model.ManualPaymentEvidenceList) ManualPaymentEvidenceListResponse {
	dtoManualPaymentEvidenceListResponse := ManualPaymentEvidenceListResponse{}
	for _, manualPaymentEvidence := range manualPaymentEvidenceList {
		dtoManualPaymentEvidenceResponse := NewManualPaymentEvidenceResponse(*manualPaymentEvidence)
		dtoManualPaymentEvidenceListResponse = append(dtoManualPaymentEvidenceListResponse, &dtoManualPaymentEvidenceResponse)
	}
	return dtoManualPaymentEvidenceListResponse
}

type ManualPaymentEvidencePrimaryIDList []ManualPaymentEvidencePrimaryID

func (d ManualPaymentEvidencePrimaryIDList) Validate() (err error) {
	validator := shared.GetValidator()
	for _, manualPaymentEvidence := range d {
		err = validator.Struct(manualPaymentEvidence)
		if err != nil {
			return
		}
	}
	return nil
}

type ManualPaymentEvidencePrimaryID struct {
	Id uuid.UUID `json:"id" validate:"uuid" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (d *ManualPaymentEvidencePrimaryID) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(d)
}
func (d ManualPaymentEvidencePrimaryID) ToModel() model.ManualPaymentEvidencePrimaryID {
	return model.ManualPaymentEvidencePrimaryID{
		Id: d.Id,
	}
}
