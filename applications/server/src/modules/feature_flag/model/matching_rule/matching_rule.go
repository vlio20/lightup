package matching_rule

type FilteringRule interface {
	IsMatch(val string) (bool, error)
}
