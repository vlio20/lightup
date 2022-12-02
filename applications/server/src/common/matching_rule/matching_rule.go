package matching_rule

type RuleType string

const (
	Number RuleType = "numer"
	String RuleType = "string"
	Ip     RuleType = "ip"
)

type Rule struct {
	Type RuleType `json:"type" bson:"type" binding:"required"`
}

type MatchingRule interface {
	IsMatch(val string) (bool, error)
}
