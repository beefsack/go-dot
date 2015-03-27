package dot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRune(t *testing.T) {
	assert.Equal(t, '⠀', Rune([4][2]bool{
		{false, false},
		{false, false},
		{false, false},
		{false, false},
	}))
	assert.Equal(t, '⣌', Rune([4][2]bool{
		{false, true},
		{false, false},
		{true, false},
		{true, true},
	}))
	assert.Equal(t, '⣿', Rune([4][2]bool{
		{true, true},
		{true, true},
		{true, true},
		{true, true},
	}))
}
