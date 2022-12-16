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
