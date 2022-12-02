package num_rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseRule = &NumericMatchingRule{
	Key:      "key",
	Operator: LessThan,
	Value:    1.0,
}

func Test_lessThen(t *testing.T) {
	rule := get_new_rule(LessThan)

	res, err := rule.IsMatch("0.5")
	assert.True(t, res)
	assert.Nil(t, err)

	resEqual, err := rule.IsMatch("1.0")
	assert.False(t, resEqual)
	assert.Nil(t, err)

	resGreater, err := rule.IsMatch("2.0")
	assert.False(t, resGreater)
	assert.Nil(t, err)
}

func Test_lessEqualThen(t *testing.T) {
	rule := get_new_rule(LessOrEqual)

	res, err := rule.IsMatch("0.5")
	assert.True(t, res)
	assert.Nil(t, err)

	resEqual, err := rule.IsMatch("1.0")
	assert.True(t, resEqual)
	assert.Nil(t, err)

	resGreater, err := rule.IsMatch("2.0")
	assert.False(t, resGreater)
	assert.Nil(t, err)
}

func Test_greaterThen(t *testing.T) {
	rule := get_new_rule(GreaterThan)

	res, err := rule.IsMatch("2")
	assert.True(t, res)
	assert.Nil(t, err)

	resEqual, err := rule.IsMatch("1")
	assert.False(t, resEqual)
	assert.Nil(t, err)

	resSamller, err := rule.IsMatch("0")
	assert.False(t, resSamller)
	assert.Nil(t, err)
}

func Test_greaterEqualThen(t *testing.T) {
	rule := get_new_rule(GreaterOrEqual)

	matchingRes, err := rule.IsMatch("2")
	assert.True(t, matchingRes)
	assert.Nil(t, err)

	matchingResEqual, err := rule.IsMatch("1")
	assert.True(t, matchingResEqual)
	assert.Nil(t, err)

	notMatchingResSamller, err := rule.IsMatch("0")
	assert.False(t, notMatchingResSamller)
	assert.Nil(t, err)
}

func Test_equal(t *testing.T) {
	rule := get_new_rule(Equal)

	res1, err := rule.IsMatch("1")
	assert.True(t, res1)
	assert.Nil(t, err)

	res2, err := rule.IsMatch("1.0")
	assert.True(t, res2)
	assert.Nil(t, err)

	res3, err := rule.IsMatch("1.1")
	assert.False(t, res3)
	assert.Nil(t, err)

	res4, err := rule.IsMatch("5")
	assert.False(t, res4)
	assert.Nil(t, err)
}

func Test_match_fail(t *testing.T) {
	invalideValue := "asd"
	res, err := baseRule.IsMatch(invalideValue)

	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), invalideValue)
}

func get_new_rule(operation NumericOperator) *NumericMatchingRule {
	gtRule := baseRule
	gtRule.Operator = operation

	return gtRule
}
