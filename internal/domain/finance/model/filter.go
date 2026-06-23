package model

// imports for filter utilities
import (
	"fmt"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"strings"
)

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"

	PaginationStrategyCursor  = "cursor"
	PaginationStrategyOffset  = "offset"
	DefaultPaginationStrategy = "cursor"

	CursorDirectionNext = "next"
	CursorDirectionPrev = "prev"

	DefaultPageSize = 10
	MaxPageSize     = 100

	OperatorEqual  = "eq"
	OperatorRange  = "range"
	OperatorIn     = "in"
	OperatorNotIn  = "not_in"
	OperatorIsNull = "is_null"
	OperatorNot    = "not"
	OperatorGT     = "gt"
	OperatorGTE    = "gte"
	OperatorLT     = "lt"
	OperatorLTE    = "lte"
	OperatorLike   = "like"

	FilterLogicAnd = "and"
	FilterLogicOr  = "or"

	MaxFilterGroupDepth = 5
	MaxFilterPredicates = 50
)

type Filter struct {
	SelectFields []string      `json:"fields"`
	Sorts        []Sort        `json:"sort"`
	Pagination   Pagination    `json:"pagination"`
	FilterFields []FilterField `json:"filter"`
	Where        *FilterGroup  `json:"where"`
}

type Sort struct {
	Field string `json:"field" validate:"required"`
	Order string `json:"order" validate:"required,oneof=ASC DESC"`
}

type Pagination struct {
	Strategy  string      `json:"strategy" validate:"omitempty,oneof=cursor offset"`
	Page      int         `json:"page" validate:"omitempty,min=1"`
	PageSize  int         `json:"pageSize" validate:"required,min=1"`
	Cursor    interface{} `json:"cursor"`
	Direction string      `json:"direction" validate:"omitempty,oneof=next prev"`
}

func (p Pagination) IsCursorMode() bool {
	if p.Strategy == PaginationStrategyOffset {
		return false
	}
	return p.Page <= 0
}

func (p Pagination) IsOffsetMode() bool {
	return !p.IsCursorMode()
}

type FilterField struct {
	Field    string      `json:"field" validate:"required"`
	Operator string      `json:"operator" validate:"required,oneof=eq not gt gte lt lte range in not_in is_null like"`
	Value    interface{} `json:"value"`
}

type FilterGroup struct {
	Logic        string        `json:"logic" validate:"omitempty,oneof=and or AND OR"`
	FilterFields []FilterField `json:"filter"`
	Groups       []FilterGroup `json:"groups"`
}

type JoinSpec struct {
	Name          string
	Table         string
	Alias         string
	Type          string
	LocalColumn   string
	ForeignColumn string
}

type FilterFieldSpec struct {
	SourcePath        string
	DefaultOutputPath string
	Relation          string
	Column            string
	SQLAlias          string
	Selectable        bool
	Filterable        bool
	Sortable          bool
}

func ParseProjection(raw string) (sourcePath, outputPath string, explicitAlias bool) {
	raw = strings.TrimSpace(raw)
	lower := strings.ToLower(raw)
	index := strings.Index(lower, " as ")
	if index < 0 {
		return raw, raw, false
	}
	sourcePath = strings.TrimSpace(raw[:index])
	outputPath = strings.TrimSpace(raw[index+4:])
	if outputPath == "" {
		outputPath = sourcePath
		explicitAlias = false
		return
	}
	return sourcePath, outputPath, true
}

func validateFilterField(filterField FilterField) error {
	validator := shared.GetValidator()
	if err := validator.Struct(filterField); err != nil {
		return err
	}
	switch filterField.Operator {
	case OperatorEqual, OperatorNot, OperatorGT, OperatorGTE, OperatorLT, OperatorLTE:
		if filterField.Value == nil {
			return failure.BadRequestFromString(fmt.Sprintf("value is required for operator %s", filterField.Operator))
		}
		if _, ok := filterField.Value.([]interface{}); ok {
			return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
	case OperatorRange:
		valueArray, ok := filterField.Value.([]interface{})
		if !ok || len(valueArray) != 2 {
			return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
	case OperatorIn, OperatorNotIn:
		valueArray, ok := filterField.Value.([]interface{})
		if !ok || len(valueArray) == 0 {
			return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
	case OperatorIsNull:
		_, ok := filterField.Value.(bool)
		if !ok {
			return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
	case OperatorLike:
		value, ok := filterField.Value.(string)
		if !ok || strings.TrimSpace(value) == "" {
			return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
		}
	}
	return nil
}

func validateFilterGroup(group FilterGroup, depth int, predicateCount *int) error {
	if depth > MaxFilterGroupDepth {
		return failure.BadRequestFromString(fmt.Sprintf("filter group depth cannot be greater than %d", MaxFilterGroupDepth))
	}
	validator := shared.GetValidator()
	if err := validator.Struct(group); err != nil {
		return err
	}
	for _, filterField := range group.FilterFields {
		(*predicateCount)++
		if *predicateCount > MaxFilterPredicates {
			return failure.BadRequestFromString(fmt.Sprintf("filter predicates cannot be greater than %d", MaxFilterPredicates))
		}
		if err := validateFilterField(filterField); err != nil {
			return err
		}
	}
	for _, child := range group.Groups {
		if err := validateFilterGroup(child, depth+1, predicateCount); err != nil {
			return err
		}
	}
	return nil
}

func (f *Filter) Validate() (err error) {
	validator := shared.GetValidator()
	if len(f.Sorts) > 0 {
		for _, sort := range f.Sorts {
			// TODO: validate field name to the respective model
			err = validator.Struct(sort)
			if err != nil {
				return
			}
		}
	}
	predicateCount := 0
	if len(f.FilterFields) > 0 {
		for _, filterField := range f.FilterFields {
			predicateCount++
			if predicateCount > MaxFilterPredicates {
				return failure.BadRequestFromString(fmt.Sprintf("filter predicates cannot be greater than %d", MaxFilterPredicates))
			}
			if err = validateFilterField(filterField); err != nil {
				return
			}
		}
	}
	if f.Where != nil {
		if err = validateFilterGroup(*f.Where, 1, &predicateCount); err != nil {
			return
		}
	}
	err = validator.Struct(f.Pagination)
	if err != nil {
		return
	}
	hasCursor := false
	if f.Pagination.Cursor != nil {
		switch v := f.Pagination.Cursor.(type) {
		case string:
			hasCursor = strings.TrimSpace(v) != ""
		default:
			hasCursor = true
		}
	}
	if f.Pagination.PageSize > MaxPageSize {
		return failure.BadRequestFromString(fmt.Sprintf("pagination.pageSize cannot be greater than %d", MaxPageSize))
	}
	if f.Pagination.Strategy == PaginationStrategyCursor && f.Pagination.Page > 0 {
		return failure.BadRequestFromString("pagination.page cannot be used with cursor strategy")
	}
	if f.Pagination.Page > 0 && hasCursor {
		return failure.BadRequestFromString("pagination.page and pagination.cursor cannot be used together")
	}
	if f.Pagination.Strategy == PaginationStrategyOffset && hasCursor {
		return failure.BadRequestFromString("pagination.cursor cannot be used with offset strategy")
	}
	return
}
