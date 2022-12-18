package num_rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lessThen(t *testing.T) {
	res, err := IsMatch(LessThan, "1.0", "0.5")
	assert.True(t, res)
	assert.Nil(t, err)

	resEqual, err := IsMatch(LessThan, "1.0", "1.0")
	assert.False(t, resEqual)
	assert.Nil(t, err)

	resGreater, err := IsMatch(LessThan, "1.0", "2.0")
	assert.False(t, resGreater)
	assert.Nil(t, err)
}

func Test_lessEqualThen(t *testing.T) {
	less, err := IsMatch(LessOrEqual, "1.0", "0.5")
	assert.True(t, less)
	assert.Nil(t, err)

	resEqual, err := IsMatch(LessOrEqual, "1.0", "1.0")
	assert.True(t, resEqual)
	assert.Nil(t, err)

	resGreater, err := IsMatch(LessOrEqual, "1.0", "1.5")
	assert.False(t, resGreater)
	assert.Nil(t, err)
}

func Test_greaterThen(t *testing.T) {
	greater, err := IsMatch(GreaterThan, "1.0", "1.5")
	assert.True(t, greater)
	assert.Nil(t, err)

	resEqual, err := IsMatch(GreaterThan, "1.0", "1.0")
	assert.False(t, resEqual)
	assert.Nil(t, err)

	resSmaller, err := IsMatch(GreaterThan, "1.0", "0.5")
	assert.False(t, resSmaller)
	assert.Nil(t, err)
}

func Test_greaterEqualThen(t *testing.T) {
	greater, err := IsMatch(GreaterOrEqual, "1.0", "1.5")
	assert.True(t, greater)
	assert.Nil(t, err)

	equal, err := IsMatch(GreaterOrEqual, "1.0", "1.0")
	assert.True(t, equal)
	assert.Nil(t, err)

	smaller, err := IsMatch(GreaterOrEqual, "1.0", "0.5")
	assert.False(t, smaller)
	assert.Nil(t, err)
}

func Test_equal(t *testing.T) {
	equal, err := IsMatch(Equal, "1.0", "1.0")
	assert.True(t, equal)
	assert.Nil(t, err)

	greater, err := IsMatch(Equal, "1.0", "1.5")
	assert.False(t, greater)
	assert.Nil(t, err)

	smaller, err := IsMatch(Equal, "1.0", "0.5")
	assert.False(t, smaller)
	assert.Nil(t, err)
}

func Test_match_fail(t *testing.T) {
	invalidValue := "asd"
	res, err := IsMatch(LessThan, "1.0", invalidValue)

	assert.False(t, res)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), invalidValue)
}
