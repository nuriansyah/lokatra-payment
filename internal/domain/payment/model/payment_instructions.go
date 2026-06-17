package model

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type PaymentInstructionsDBFieldNameType string

type paymentInstructionsDBFieldName struct {
	Id                  PaymentInstructionsDBFieldNameType
	PaymentAttemptId    PaymentInstructionsDBFieldNameType
	InstructionType     PaymentInstructionsDBFieldNameType
	IsActive            PaymentInstructionsDBFieldNameType
	DisplayName         PaymentInstructionsDBFieldNameType
	AccountNumber       PaymentInstructionsDBFieldNameType
	AccountNumberMasked PaymentInstructionsDBFieldNameType
	AccountHolderName   PaymentInstructionsDBFieldNameType
	BankCode            PaymentInstructionsDBFieldNameType
	BillerCode          PaymentInstructionsDBFieldNameType
	PaymentCode         PaymentInstructionsDBFieldNameType
	QrString            PaymentInstructionsDBFieldNameType
	QrImageUrl          PaymentInstructionsDBFieldNameType
	CheckoutUrl         PaymentInstructionsDBFieldNameType
	DeeplinkUrl         PaymentInstructionsDBFieldNameType
	RetailOutletCode    PaymentInstructionsDBFieldNameType
	ExpiresAt           PaymentInstructionsDBFieldNameType
	Metadata            PaymentInstructionsDBFieldNameType
	MetaCreatedAt       PaymentInstructionsDBFieldNameType
	MetaCreatedBy       PaymentInstructionsDBFieldNameType
	MetaUpdatedAt       PaymentInstructionsDBFieldNameType
	MetaUpdatedBy       PaymentInstructionsDBFieldNameType
	MetaDeletedAt       PaymentInstructionsDBFieldNameType
	MetaDeletedBy       PaymentInstructionsDBFieldNameType
}

var PaymentInstructionsDBFieldName = paymentInstructionsDBFieldName{
	Id:                  "id",
	PaymentAttemptId:    "payment_attempt_id",
	InstructionType:     "instruction_type",
	IsActive:            "is_active",
	DisplayName:         "display_name",
	AccountNumber:       "account_number",
	AccountNumberMasked: "account_number_masked",
	AccountHolderName:   "account_holder_name",
	BankCode:            "bank_code",
	BillerCode:          "biller_code",
	PaymentCode:         "payment_code",
	QrString:            "qr_string",
	QrImageUrl:          "qr_image_url",
	CheckoutUrl:         "checkout_url",
	DeeplinkUrl:         "deeplink_url",
	RetailOutletCode:    "retail_outlet_code",
	ExpiresAt:           "expires_at",
	Metadata:            "metadata",
	MetaCreatedAt:       "meta_created_at",
	MetaCreatedBy:       "meta_created_by",
	MetaUpdatedAt:       "meta_updated_at",
	MetaUpdatedBy:       "meta_updated_by",
	MetaDeletedAt:       "meta_deleted_at",
	MetaDeletedBy:       "meta_deleted_by",
}

func NewPaymentInstructionsDBFieldNameFromStr(field string) (dbField PaymentInstructionsDBFieldNameType, found bool) {
	switch field {

	case string(PaymentInstructionsDBFieldName.Id):
		return PaymentInstructionsDBFieldName.Id, true

	case string(PaymentInstructionsDBFieldName.PaymentAttemptId):
		return PaymentInstructionsDBFieldName.PaymentAttemptId, true

	case string(PaymentInstructionsDBFieldName.InstructionType):
		return PaymentInstructionsDBFieldName.InstructionType, true

	case string(PaymentInstructionsDBFieldName.IsActive):
		return PaymentInstructionsDBFieldName.IsActive, true

	case string(PaymentInstructionsDBFieldName.DisplayName):
		return PaymentInstructionsDBFieldName.DisplayName, true

	case string(PaymentInstructionsDBFieldName.AccountNumber):
		return PaymentInstructionsDBFieldName.AccountNumber, true

	case string(PaymentInstructionsDBFieldName.AccountNumberMasked):
		return PaymentInstructionsDBFieldName.AccountNumberMasked, true

	case string(PaymentInstructionsDBFieldName.AccountHolderName):
		return PaymentInstructionsDBFieldName.AccountHolderName, true

	case string(PaymentInstructionsDBFieldName.BankCode):
		return PaymentInstructionsDBFieldName.BankCode, true

	case string(PaymentInstructionsDBFieldName.BillerCode):
		return PaymentInstructionsDBFieldName.BillerCode, true

	case string(PaymentInstructionsDBFieldName.PaymentCode):
		return PaymentInstructionsDBFieldName.PaymentCode, true

	case string(PaymentInstructionsDBFieldName.QrString):
		return PaymentInstructionsDBFieldName.QrString, true

	case string(PaymentInstructionsDBFieldName.QrImageUrl):
		return PaymentInstructionsDBFieldName.QrImageUrl, true

	case string(PaymentInstructionsDBFieldName.CheckoutUrl):
		return PaymentInstructionsDBFieldName.CheckoutUrl, true

	case string(PaymentInstructionsDBFieldName.DeeplinkUrl):
		return PaymentInstructionsDBFieldName.DeeplinkUrl, true

	case string(PaymentInstructionsDBFieldName.RetailOutletCode):
		return PaymentInstructionsDBFieldName.RetailOutletCode, true

	case string(PaymentInstructionsDBFieldName.ExpiresAt):
		return PaymentInstructionsDBFieldName.ExpiresAt, true

	case string(PaymentInstructionsDBFieldName.Metadata):
		return PaymentInstructionsDBFieldName.Metadata, true

	case string(PaymentInstructionsDBFieldName.MetaCreatedAt):
		return PaymentInstructionsDBFieldName.MetaCreatedAt, true

	case string(PaymentInstructionsDBFieldName.MetaCreatedBy):
		return PaymentInstructionsDBFieldName.MetaCreatedBy, true

	case string(PaymentInstructionsDBFieldName.MetaUpdatedAt):
		return PaymentInstructionsDBFieldName.MetaUpdatedAt, true

	case string(PaymentInstructionsDBFieldName.MetaUpdatedBy):
		return PaymentInstructionsDBFieldName.MetaUpdatedBy, true

	case string(PaymentInstructionsDBFieldName.MetaDeletedAt):
		return PaymentInstructionsDBFieldName.MetaDeletedAt, true

	case string(PaymentInstructionsDBFieldName.MetaDeletedBy):
		return PaymentInstructionsDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var PaymentInstructionsFilterJoins = map[string]JoinSpec{}

var PaymentInstructionsFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_attempt_id": {
		SourcePath:        "payment_attempt_id",
		DefaultOutputPath: "paymentAttemptId",
		Column:            "payment_attempt_id",
		SQLAlias:          "payment_attempt_id",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"instruction_type": {
		SourcePath:        "instruction_type",
		DefaultOutputPath: "instructionType",
		Column:            "instruction_type",
		SQLAlias:          "instruction_type",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"is_active": {
		SourcePath:        "is_active",
		DefaultOutputPath: "isActive",
		Column:            "is_active",
		SQLAlias:          "is_active",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"display_name": {
		SourcePath:        "display_name",
		DefaultOutputPath: "displayName",
		Column:            "display_name",
		SQLAlias:          "display_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_number": {
		SourcePath:        "account_number",
		DefaultOutputPath: "accountNumber",
		Column:            "account_number",
		SQLAlias:          "account_number",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_number_masked": {
		SourcePath:        "account_number_masked",
		DefaultOutputPath: "accountNumberMasked",
		Column:            "account_number_masked",
		SQLAlias:          "account_number_masked",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"account_holder_name": {
		SourcePath:        "account_holder_name",
		DefaultOutputPath: "accountHolderName",
		Column:            "account_holder_name",
		SQLAlias:          "account_holder_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"bank_code": {
		SourcePath:        "bank_code",
		DefaultOutputPath: "bankCode",
		Column:            "bank_code",
		SQLAlias:          "bank_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"biller_code": {
		SourcePath:        "biller_code",
		DefaultOutputPath: "billerCode",
		Column:            "biller_code",
		SQLAlias:          "biller_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"payment_code": {
		SourcePath:        "payment_code",
		DefaultOutputPath: "paymentCode",
		Column:            "payment_code",
		SQLAlias:          "payment_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"qr_string": {
		SourcePath:        "qr_string",
		DefaultOutputPath: "qrString",
		Column:            "qr_string",
		SQLAlias:          "qr_string",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"qr_image_url": {
		SourcePath:        "qr_image_url",
		DefaultOutputPath: "qrImageUrl",
		Column:            "qr_image_url",
		SQLAlias:          "qr_image_url",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"checkout_url": {
		SourcePath:        "checkout_url",
		DefaultOutputPath: "checkoutUrl",
		Column:            "checkout_url",
		SQLAlias:          "checkout_url",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"deeplink_url": {
		SourcePath:        "deeplink_url",
		DefaultOutputPath: "deeplinkUrl",
		Column:            "deeplink_url",
		SQLAlias:          "deeplink_url",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"retail_outlet_code": {
		SourcePath:        "retail_outlet_code",
		DefaultOutputPath: "retailOutletCode",
		Column:            "retail_outlet_code",
		SQLAlias:          "retail_outlet_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"expires_at": {
		SourcePath:        "expires_at",
		DefaultOutputPath: "expiresAt",
		Column:            "expires_at",
		SQLAlias:          "expires_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"metadata": {
		SourcePath:        "metadata",
		DefaultOutputPath: "metadata",
		Column:            "metadata",
		SQLAlias:          "metadata",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_at": {
		SourcePath:        "meta_created_at",
		DefaultOutputPath: "metaCreatedAt",
		Column:            "meta_created_at",
		SQLAlias:          "meta_created_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_created_by": {
		SourcePath:        "meta_created_by",
		DefaultOutputPath: "metaCreatedBy",
		Column:            "meta_created_by",
		SQLAlias:          "meta_created_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_at": {
		SourcePath:        "meta_updated_at",
		DefaultOutputPath: "metaUpdatedAt",
		Column:            "meta_updated_at",
		SQLAlias:          "meta_updated_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_updated_by": {
		SourcePath:        "meta_updated_by",
		DefaultOutputPath: "metaUpdatedBy",
		Column:            "meta_updated_by",
		SQLAlias:          "meta_updated_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_at": {
		SourcePath:        "meta_deleted_at",
		DefaultOutputPath: "metaDeletedAt",
		Column:            "meta_deleted_at",
		SQLAlias:          "meta_deleted_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"meta_deleted_by": {
		SourcePath:        "meta_deleted_by",
		DefaultOutputPath: "metaDeletedBy",
		Column:            "meta_deleted_by",
		SQLAlias:          "meta_deleted_by",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
}

func NewPaymentInstructionsFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = PaymentInstructionsFilterFields[field]
	return
}

type PaymentInstructionsFilterResult struct {
	PaymentInstructions
	FilterCount int `db:"count"`
}

func ValidatePaymentInstructionsFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewPaymentInstructionsFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewPaymentInstructionsFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewPaymentInstructionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validatePaymentInstructionsFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validatePaymentInstructionsFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewPaymentInstructionsFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validatePaymentInstructionsFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type PaymentInstructions struct {
	Id                  uuid.UUID       `db:"id"`
	PaymentAttemptId    uuid.UUID       `db:"payment_attempt_id"`
	InstructionType     string          `db:"instruction_type"`
	IsActive            bool            `db:"is_active"`
	DisplayName         null.String     `db:"display_name"`
	AccountNumber       null.String     `db:"account_number"`
	AccountNumberMasked null.String     `db:"account_number_masked"`
	AccountHolderName   null.String     `db:"account_holder_name"`
	BankCode            null.String     `db:"bank_code"`
	BillerCode          null.String     `db:"biller_code"`
	PaymentCode         null.String     `db:"payment_code"`
	QrString            null.String     `db:"qr_string"`
	QrImageUrl          null.String     `db:"qr_image_url"`
	CheckoutUrl         null.String     `db:"checkout_url"`
	DeeplinkUrl         null.String     `db:"deeplink_url"`
	RetailOutletCode    null.String     `db:"retail_outlet_code"`
	ExpiresAt           null.Time       `db:"expires_at"`
	Metadata            json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type PaymentInstructionsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d PaymentInstructions) ToPaymentInstructionsPrimaryID() PaymentInstructionsPrimaryID {
	return PaymentInstructionsPrimaryID{
		Id: d.Id,
	}
}

type PaymentInstructionsList []*PaymentInstructions

type PaymentInstructionsFilterResultList []*PaymentInstructionsFilterResult
