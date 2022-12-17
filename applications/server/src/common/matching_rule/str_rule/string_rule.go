package str_rule

import (
	"errors"
	"strings"
)

type StringOperator string

const (
	StartsWith StringOperator = "startsWith"
	EndsWith   StringOperator = "endsWith"
	Contains   StringOperator = "contains"
	Equals     StringOperator = "equals"
)

type StringFilteringRule struct {
	Key      string         `json:"key" bson:"key" binding:"required"`
	Operator StringOperator `json:"operator" bson:"operator" binding:"required"`
	Value    string         `json:"value" bson:"value" binding:"required"`
}

func (r *StringFilteringRule) GetKey() string {
	return r.Key
}

func (r *StringFilteringRule) IsMatch(val string) (bool, error) {
	switch r.Operator {
	case StartsWith:
		return strings.HasPrefix(val, r.Value), nil
	case EndsWith:
		return strings.HasSuffix(val, r.Value), nil
	case Contains:
		return strings.Contains(val, r.Value), nil
	case Equals:
		return val == r.Value, nil
	default:
		return false, errors.New("Invalid operator " + string(r.Operator))
	}
}
