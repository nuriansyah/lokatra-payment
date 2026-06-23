package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

type BankStatementFilesDBFieldNameType string

type bankStatementFilesDBFieldName struct {
	Id                   BankStatementFilesDBFieldNameType
	BankCode             BankStatementFilesDBFieldNameType
	AccountNumberMasked  BankStatementFilesDBFieldNameType
	StatementPeriodStart BankStatementFilesDBFieldNameType
	StatementPeriodEnd   BankStatementFilesDBFieldNameType
	CurrencyCode         BankStatementFilesDBFieldNameType
	StorageUri           BankStatementFilesDBFieldNameType
	FileHash             BankStatementFilesDBFieldNameType
	ImportStatus         BankStatementFilesDBFieldNameType
	ImportedAt           BankStatementFilesDBFieldNameType
	RowCount             BankStatementFilesDBFieldNameType
	Metadata             BankStatementFilesDBFieldNameType
	MetaCreatedAt        BankStatementFilesDBFieldNameType
	MetaCreatedBy        BankStatementFilesDBFieldNameType
	MetaUpdatedAt        BankStatementFilesDBFieldNameType
	MetaUpdatedBy        BankStatementFilesDBFieldNameType
	MetaDeletedAt        BankStatementFilesDBFieldNameType
	MetaDeletedBy        BankStatementFilesDBFieldNameType
}

var BankStatementFilesDBFieldName = bankStatementFilesDBFieldName{
	Id:                   "id",
	BankCode:             "bank_code",
	AccountNumberMasked:  "account_number_masked",
	StatementPeriodStart: "statement_period_start",
	StatementPeriodEnd:   "statement_period_end",
	CurrencyCode:         "currency_code",
	StorageUri:           "storage_uri",
	FileHash:             "file_hash",
	ImportStatus:         "import_status",
	ImportedAt:           "imported_at",
	RowCount:             "row_count",
	Metadata:             "metadata",
	MetaCreatedAt:        "meta_created_at",
	MetaCreatedBy:        "meta_created_by",
	MetaUpdatedAt:        "meta_updated_at",
	MetaUpdatedBy:        "meta_updated_by",
	MetaDeletedAt:        "meta_deleted_at",
	MetaDeletedBy:        "meta_deleted_by",
}

func NewBankStatementFilesDBFieldNameFromStr(field string) (dbField BankStatementFilesDBFieldNameType, found bool) {
	switch field {

	case string(BankStatementFilesDBFieldName.Id):
		return BankStatementFilesDBFieldName.Id, true

	case string(BankStatementFilesDBFieldName.BankCode):
		return BankStatementFilesDBFieldName.BankCode, true

	case string(BankStatementFilesDBFieldName.AccountNumberMasked):
		return BankStatementFilesDBFieldName.AccountNumberMasked, true

	case string(BankStatementFilesDBFieldName.StatementPeriodStart):
		return BankStatementFilesDBFieldName.StatementPeriodStart, true

	case string(BankStatementFilesDBFieldName.StatementPeriodEnd):
		return BankStatementFilesDBFieldName.StatementPeriodEnd, true

	case string(BankStatementFilesDBFieldName.CurrencyCode):
		return BankStatementFilesDBFieldName.CurrencyCode, true

	case string(BankStatementFilesDBFieldName.StorageUri):
		return BankStatementFilesDBFieldName.StorageUri, true

	case string(BankStatementFilesDBFieldName.FileHash):
		return BankStatementFilesDBFieldName.FileHash, true

	case string(BankStatementFilesDBFieldName.ImportStatus):
		return BankStatementFilesDBFieldName.ImportStatus, true

	case string(BankStatementFilesDBFieldName.ImportedAt):
		return BankStatementFilesDBFieldName.ImportedAt, true

	case string(BankStatementFilesDBFieldName.RowCount):
		return BankStatementFilesDBFieldName.RowCount, true

	case string(BankStatementFilesDBFieldName.Metadata):
		return BankStatementFilesDBFieldName.Metadata, true

	case string(BankStatementFilesDBFieldName.MetaCreatedAt):
		return BankStatementFilesDBFieldName.MetaCreatedAt, true

	case string(BankStatementFilesDBFieldName.MetaCreatedBy):
		return BankStatementFilesDBFieldName.MetaCreatedBy, true

	case string(BankStatementFilesDBFieldName.MetaUpdatedAt):
		return BankStatementFilesDBFieldName.MetaUpdatedAt, true

	case string(BankStatementFilesDBFieldName.MetaUpdatedBy):
		return BankStatementFilesDBFieldName.MetaUpdatedBy, true

	case string(BankStatementFilesDBFieldName.MetaDeletedAt):
		return BankStatementFilesDBFieldName.MetaDeletedAt, true

	case string(BankStatementFilesDBFieldName.MetaDeletedBy):
		return BankStatementFilesDBFieldName.MetaDeletedBy, true

	}
	return "unknown", false
}

var BankStatementFilesFilterJoins = map[string]JoinSpec{}

var BankStatementFilesFilterFields = map[string]FilterFieldSpec{
	"id": {
		SourcePath:        "id",
		DefaultOutputPath: "id",
		Column:            "id",
		SQLAlias:          "id",
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
	"account_number_masked": {
		SourcePath:        "account_number_masked",
		DefaultOutputPath: "accountNumberMasked",
		Column:            "account_number_masked",
		SQLAlias:          "account_number_masked",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"statement_period_start": {
		SourcePath:        "statement_period_start",
		DefaultOutputPath: "statementPeriodStart",
		Column:            "statement_period_start",
		SQLAlias:          "statement_period_start",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"statement_period_end": {
		SourcePath:        "statement_period_end",
		DefaultOutputPath: "statementPeriodEnd",
		Column:            "statement_period_end",
		SQLAlias:          "statement_period_end",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"currency_code": {
		SourcePath:        "currency_code",
		DefaultOutputPath: "currencyCode",
		Column:            "currency_code",
		SQLAlias:          "currency_code",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"storage_uri": {
		SourcePath:        "storage_uri",
		DefaultOutputPath: "storageUri",
		Column:            "storage_uri",
		SQLAlias:          "storage_uri",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"file_hash": {
		SourcePath:        "file_hash",
		DefaultOutputPath: "fileHash",
		Column:            "file_hash",
		SQLAlias:          "file_hash",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"import_status": {
		SourcePath:        "import_status",
		DefaultOutputPath: "importStatus",
		Column:            "import_status",
		SQLAlias:          "import_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"imported_at": {
		SourcePath:        "imported_at",
		DefaultOutputPath: "importedAt",
		Column:            "imported_at",
		SQLAlias:          "imported_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"row_count": {
		SourcePath:        "row_count",
		DefaultOutputPath: "rowCount",
		Column:            "row_count",
		SQLAlias:          "row_count",
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

func NewBankStatementFilesFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = BankStatementFilesFilterFields[field]
	return
}

type BankStatementFilesFilterResult struct {
	BankStatementFiles
	FilterCount int `db:"count"`
}

func ValidateBankStatementFilesFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewBankStatementFilesFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewBankStatementFilesFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewBankStatementFilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateBankStatementFilesFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateBankStatementFilesFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewBankStatementFilesFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateBankStatementFilesFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type ImportStatus string

const (
	ImportStatusUploaded   ImportStatus = "uploaded"
	ImportStatusProcessing ImportStatus = "processing"
	ImportStatusCompleted  ImportStatus = "completed"
	ImportStatusFailed     ImportStatus = "failed"
	ImportStatusCancelled  ImportStatus = "cancelled"
)

type BankStatementFiles struct {
	Id                   uuid.UUID       `db:"id"`
	BankCode             string          `db:"bank_code"`
	AccountNumberMasked  null.String     `db:"account_number_masked"`
	StatementPeriodStart time.Time       `db:"statement_period_start"`
	StatementPeriodEnd   time.Time       `db:"statement_period_end"`
	CurrencyCode         string          `db:"currency_code"`
	StorageUri           null.String     `db:"storage_uri"`
	FileHash             string          `db:"file_hash"`
	ImportStatus         ImportStatus    `db:"import_status"`
	ImportedAt           null.Time       `db:"imported_at"`
	RowCount             int             `db:"row_count"`
	Metadata             json.RawMessage `db:"metadata"`

	shared.MetaSignature
}
type BankStatementFilesPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d BankStatementFiles) ToBankStatementFilesPrimaryID() BankStatementFilesPrimaryID {
	return BankStatementFilesPrimaryID{
		Id: d.Id,
	}
}

type BankStatementFilesList []*BankStatementFiles

type BankStatementFilesFilterResultList []*BankStatementFilesFilterResult
