package model

import (
	"fmt"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/shopspring/decimal"
	"time"
)

type FinanceOpsQueueSummaryDBFieldNameType string

type financeOpsQueueSummaryDBFieldName struct {
	QueueName    FinanceOpsQueueSummaryDBFieldNameType
	QueueStatus  FinanceOpsQueueSummaryDBFieldNameType
	ItemCount    FinanceOpsQueueSummaryDBFieldNameType
	OldestItemAt FinanceOpsQueueSummaryDBFieldNameType
	TotalAmount  FinanceOpsQueueSummaryDBFieldNameType
	CurrencyCode FinanceOpsQueueSummaryDBFieldNameType
	RefreshedAt  FinanceOpsQueueSummaryDBFieldNameType
}

var FinanceOpsQueueSummaryDBFieldName = financeOpsQueueSummaryDBFieldName{
	QueueName:    "queue_name",
	QueueStatus:  "queue_status",
	ItemCount:    "item_count",
	OldestItemAt: "oldest_item_at",
	TotalAmount:  "total_amount",
	CurrencyCode: "currency_code",
	RefreshedAt:  "refreshed_at",
}

func NewFinanceOpsQueueSummaryDBFieldNameFromStr(field string) (dbField FinanceOpsQueueSummaryDBFieldNameType, found bool) {
	switch field {

	case string(FinanceOpsQueueSummaryDBFieldName.QueueName):
		return FinanceOpsQueueSummaryDBFieldName.QueueName, true

	case string(FinanceOpsQueueSummaryDBFieldName.QueueStatus):
		return FinanceOpsQueueSummaryDBFieldName.QueueStatus, true

	case string(FinanceOpsQueueSummaryDBFieldName.ItemCount):
		return FinanceOpsQueueSummaryDBFieldName.ItemCount, true

	case string(FinanceOpsQueueSummaryDBFieldName.OldestItemAt):
		return FinanceOpsQueueSummaryDBFieldName.OldestItemAt, true

	case string(FinanceOpsQueueSummaryDBFieldName.TotalAmount):
		return FinanceOpsQueueSummaryDBFieldName.TotalAmount, true

	case string(FinanceOpsQueueSummaryDBFieldName.CurrencyCode):
		return FinanceOpsQueueSummaryDBFieldName.CurrencyCode, true

	case string(FinanceOpsQueueSummaryDBFieldName.RefreshedAt):
		return FinanceOpsQueueSummaryDBFieldName.RefreshedAt, true

	}
	return "unknown", false
}

var FinanceOpsQueueSummaryFilterJoins = map[string]JoinSpec{}

var FinanceOpsQueueSummaryFilterFields = map[string]FilterFieldSpec{
	"queue_name": {
		SourcePath:        "queue_name",
		DefaultOutputPath: "queueName",
		Column:            "queue_name",
		SQLAlias:          "queue_name",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"queue_status": {
		SourcePath:        "queue_status",
		DefaultOutputPath: "queueStatus",
		Column:            "queue_status",
		SQLAlias:          "queue_status",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"item_count": {
		SourcePath:        "item_count",
		DefaultOutputPath: "itemCount",
		Column:            "item_count",
		SQLAlias:          "item_count",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"oldest_item_at": {
		SourcePath:        "oldest_item_at",
		DefaultOutputPath: "oldestItemAt",
		Column:            "oldest_item_at",
		SQLAlias:          "oldest_item_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
	"total_amount": {
		SourcePath:        "total_amount",
		DefaultOutputPath: "totalAmount",
		Column:            "total_amount",
		SQLAlias:          "total_amount",
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
	"refreshed_at": {
		SourcePath:        "refreshed_at",
		DefaultOutputPath: "refreshedAt",
		Column:            "refreshed_at",
		SQLAlias:          "refreshed_at",
		Selectable:        true,
		Filterable:        true,
		Sortable:          true,
	},
}

func NewFinanceOpsQueueSummaryFilterFieldSpecFromStr(field string) (spec FilterFieldSpec, found bool) {
	spec, found = FinanceOpsQueueSummaryFilterFields[field]
	return
}

type FinanceOpsQueueSummaryFilterResult struct {
	FinanceOpsQueueSummary
	FilterCount int `db:"count"`
}

func ValidateFinanceOpsQueueSummaryFieldNameFilter(filter Filter) (err error) {
	for _, selectField := range filter.SelectFields {
		sourceField, _, _ := ParseProjection(selectField)
		spec, exist := NewFinanceOpsQueueSummaryFilterFieldSpecFromStr(sourceField)
		if !exist || !spec.Selectable || spec.Relation != "" {
			err = failure.BadRequest(fmt.Errorf("field %s is not selectable", sourceField))
			return
		}
	}
	for _, sort := range filter.Sorts {
		spec, exist := NewFinanceOpsQueueSummaryFilterFieldSpecFromStr(sort.Field)
		if !exist || !spec.Sortable {
			err = failure.BadRequest(fmt.Errorf("field %s is not sortable", sort.Field))
			return
		}
	}
	for _, field := range filter.FilterFields {
		spec, exist := NewFinanceOpsQueueSummaryFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	if filter.Where != nil {
		err = validateFinanceOpsQueueSummaryFilterGroupFieldNames(*filter.Where)
		if err != nil {
			return
		}
	}
	return
}

func validateFinanceOpsQueueSummaryFilterGroupFieldNames(group FilterGroup) (err error) {
	for _, field := range group.FilterFields {
		spec, exist := NewFinanceOpsQueueSummaryFilterFieldSpecFromStr(field.Field)
		if !exist || !spec.Filterable {
			err = failure.BadRequest(fmt.Errorf("field %s is not filterable", field.Field))
			return
		}
	}
	for _, child := range group.Groups {
		err = validateFinanceOpsQueueSummaryFilterGroupFieldNames(child)
		if err != nil {
			return
		}
	}
	return
}

type FinanceOpsQueueSummary struct {
	QueueName    string          `db:"queue_name"`
	QueueStatus  string          `db:"queue_status"`
	ItemCount    int             `db:"item_count"`
	OldestItemAt null.Time       `db:"oldest_item_at"`
	TotalAmount  decimal.Decimal `db:"total_amount"`
	CurrencyCode string          `db:"currency_code"`
	RefreshedAt  time.Time       `db:"refreshed_at"`
}
type FinanceOpsQueueSummaryPrimaryID struct {
	QueueName   string `db:"queue_name"`
	QueueStatus string `db:"queue_status"`

	CurrencyCode string `db:"currency_code"`
}

func (d FinanceOpsQueueSummary) ToFinanceOpsQueueSummaryPrimaryID() FinanceOpsQueueSummaryPrimaryID {
	return FinanceOpsQueueSummaryPrimaryID{
		QueueName:   d.QueueName,
		QueueStatus: d.QueueStatus,

		CurrencyCode: d.CurrencyCode,
	}
}

type FinanceOpsQueueSummaryList []*FinanceOpsQueueSummary

type FinanceOpsQueueSummaryFilterResultList []*FinanceOpsQueueSummaryFilterResult
