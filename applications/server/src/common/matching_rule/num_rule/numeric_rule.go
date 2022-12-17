package num_rule

import (
	"errors"
	"strconv"
)

type NumericOperator string

const (
	LessThan       NumericOperator = "lt"
	LessOrEqual    NumericOperator = "le"
	Equal          NumericOperator = "eq"
	GreaterOrEqual NumericOperator = "ge"
	GreaterThan    NumericOperator = "gt"
)

type NumericMatchingRule struct {
	Key      string          `json:"key" bson:"key" binding:"required"`
	Operator NumericOperator `json:"operator" bson:"operator" binding:"required"`
	Value    float64         `json:"value" bson:"value" binding:"required"`
}

func (r *NumericMatchingRule) GetKey() string {
	return r.Key
}

func (r *NumericMatchingRule) IsMatch(strVal string) (bool, error) {
	val, err := strconv.ParseFloat(strVal, 64)

	if err != nil {
		return false, errors.New("Invalid value" + strVal)
	}

	switch r.Operator {
	case LessThan:
		return val < r.Value, nil
	case LessOrEqual:
		return val <= r.Value, nil
	case Equal:
		return val == r.Value, nil
	case GreaterOrEqual:
		return val >= r.Value, nil
	case GreaterThan:
		return val > r.Value, nil
	default:
		return false, errors.New("Invalid operator " + string(r.Operator))
	}
}
