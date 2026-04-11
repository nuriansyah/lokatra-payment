package model

// imports for filter utilities
import (
	"fmt"

	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"

	OperatorEqual  = "eq"
	OperatorRange  = "range"
	OperatorIn     = "in"
	OperatorIsNull = "is_null"
	OperatorNot    = "not"
)

type Filter struct {
	SelectFields []string      `json:"fields"`
	Sorts        []Sort        `json:"sort"`
	Pagination   Pagination    `json:"pagination"`
	FilterFields []FilterField `json:"filter"`
}

type Sort struct {
	Field string `json:"field" validate:"required"`
	Order string `json:"order" validate:"required,oneof=ASC DESC"`
}

type Pagination struct {
	Page     int `json:"page" validate:"required,min=1"`
	PageSize int `json:"pageSize" validate:"required,min=1"`
}

type FilterField struct {
	Field    string      `json:"field" validate:"required"`
	Operator string      `json:"operator" validate:"required,oneof=eq range in is_null not"`
	Value    interface{} `json:"value" validate:"required"`
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
	if len(f.FilterFields) > 0 {
		for _, filterField := range f.FilterFields {
			// TODO: validate field name and type to the respective model
			err = validator.Struct(filterField)
			if err != nil {
				return
			}
			switch filterField.Operator {
			case OperatorEqual, OperatorNot:
				switch filterField.Value.(type) {
				case []interface{}:
					return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
				}
			case OperatorRange, OperatorIn:
				switch filterField.Value.(type) {
				case []interface{}:
					if filterField.Operator == OperatorRange && len(filterField.Value.([]interface{})) != 2 {
						return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
					}
				default:
					return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
				}
			case OperatorIsNull:
				_, ok := filterField.Value.(bool)
				if !ok {
					return failure.BadRequestFromString(fmt.Sprintf("invalid value type for operator %s", filterField.Operator))
				}
			}
		}
	}
	return validator.Struct(f.Pagination)
}
