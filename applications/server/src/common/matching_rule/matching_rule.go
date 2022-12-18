package matching_rule

import (
	"lightup/src/common/http"
	"lightup/src/common/matching_rule/num_rule"
	"lightup/src/common/matching_rule/str_rule"
)

type RuleType string

const (
	Number RuleType = "number"
	String RuleType = "string"
	Ip     RuleType = "ip"
)

type Rule struct {
	Type     RuleType `json:"type" bson:"type" binding:"required"`
	Key      string   `json:"key" bson:"key" binding:"required"`
	Operator string   `json:"operator" bson:"operator" binding:"required"`
	Value    string   `json:"value" bson:"value" binding:"required"`
}

func (r *Rule) IsMatch(strVal string) (bool, error) {
	switch r.Type {
	case Number:
		return num_rule.IsMatch(r.Operator, r.Value, strVal)
	case String:
		return str_rule.IsMatch(r.Operator, r.Value, strVal)
	default:
		return false, http.Error{
			StatusCode: 400,
			Message:    "Invalid rule type " + string(r.Type),
		}
	}
}
