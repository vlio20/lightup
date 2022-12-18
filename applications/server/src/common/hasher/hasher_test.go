package hasher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasher(t *testing.T) {
	hasher := New()
	hashedPassword, err := hasher.Hash("123")

	assert.Nil(t, err)
	assert.True(t, hasher.CheckHash("123", hashedPassword))
	assert.False(t, hasher.CheckHash("1234", hashedPassword))
}

func Test_HashStringToFloat(t *testing.T) {
	hasher := New()
	res1 := hasher.HashStringToFloat("123", "asda2d")
	res2 := hasher.HashStringToFloat("123", "asda2d")
	res3 := hasher.HashStringToFloat("124", "asda2d")
	res4 := hasher.HashStringToFloat("123", "asda3d")

	assert.Equal(t, res1, res2)
	assert.NotEqual(t, res1, res3)
	assert.NotEqual(t, res1, res4)
	assert.NotEqual(t, res2, res3)
}
