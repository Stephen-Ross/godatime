package extensions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDivModQuotient(t *testing.T) {
	var expected int64 = 7
	actual, _ := DivMod(100, 13)

	assert.Equal(t, expected, actual)
}

func TestDivModRemainder(t *testing.T) {
	var expected int64 = 9
	_, actual := DivMod(100, 13)

	assert.Equal(t, expected, actual)
}