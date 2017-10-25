package godatime

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOffsetMaxReturnsTheGreaterOffset(t *testing.T) {
	offset := Offset{seconds:1}
	otherOffset := Offset{seconds:2}

	actual := OffsetMax(&offset, &otherOffset)

	assert.Equal(t, otherOffset, *actual)
}

func TestOffsetMinReturnsTheLowerOffset(t *testing.T) {
	offset := Offset{seconds:1}
	otherOffset := Offset{seconds:2}

	actual := OffsetMin(&offset, &otherOffset)

	assert.Equal(t, offset, *actual)
}

func TestOffsetSecondsReturnsTheSecondsOfTheOffset(t *testing.T) {
	offset := OffsetFromSeconds(10)

	actual := offset.Seconds()

	assert.Equal(t, 10, actual)
}

func TestOffsetMillisecondsReturnsTheMillisecondsOfTheOffset(t *testing.T) {
	offset := OffsetFromMilliseconds(5000)

	actual := offset.Milliseconds()

	assert.Equal(t, 5000, actual)
}

func TestOffsetTicksReturnsTheTicksOfTheOffset(t *testing.T) {
	offset := OffsetFromTicks(50000000)

	actual := offset.Ticks()

	assert.Equal(t, (int64)(50000000), actual)
}

func TestOffsetNanosecondsReturnsTheNanosecondsOfTheOffset(t *testing.T) {
	offset := OffsetFromNanoseconds(5000000000)

	actual := offset.Nanoseconds()

	assert.Equal(t, (int64)(5000000000), actual)
}

func TestOffsetNegateCorrectlyNegatesTheOffset(t *testing.T) {
	offset := Offset{seconds:5}

	negatedOffset := offset.Negate()

	assert.Equal(t, -5, negatedOffset.Seconds())
}

func TestOffsetAddCorrectlyAddsTheTwoOffsetsIntoANewOffset(t *testing.T) {
	offset := Offset{seconds:2}
	otherOffset := Offset{seconds:3}

	actual := offset.Add(&otherOffset)

	assert.Equal(t, 5, actual.Seconds())
}

func TestOffsetSubtractCorrectlySubtractsTheTwoOffsetsIntoANewOffset(t *testing.T) {
	offset := Offset{seconds:5}
	otherOffset := Offset{seconds:3}

	actual := offset.Subtract(&otherOffset)

	assert.Equal(t, 2, actual.Seconds())
}

func TestOffsetLessThanReturnsFalseWhenGreaterThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:5}
	otherOffset := Offset{seconds:1}

	actual := offset.LessThan(&otherOffset)

	assert.False(t, actual)
}

func TestOffsetLessThanReturnsFalseWhenEqualToTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:2}
	otherOffset := Offset{seconds:2}

	actual := offset.LessThan(&otherOffset)

	assert.False(t, actual)
}

func TestOffsetLessThanReturnsTrueWhenLessThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:1}
	otherOffset := Offset{seconds:5}

	actual := offset.LessThan(&otherOffset)

	assert.True(t, actual)
}

func TestOffsetLessThanOrEqualsReturnsFalseWhenGreaterThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:5}
	otherOffset := Offset{seconds:1}

	actual := offset.LessThanOrEquals(&otherOffset)

	assert.False(t, actual)
}

func TestOffsetLessThanOrEqualsReturnsTrueWhenEqualToTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:2}
	otherOffset := Offset{seconds:2}

	actual := offset.LessThanOrEquals(&otherOffset)

	assert.True(t, actual)
}

func TestOffsetLessThanOrEqualsReturnsTrueWhenLessThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:1}
	otherOffset := Offset{seconds:5}

	actual := offset.LessThanOrEquals(&otherOffset)

	assert.True(t, actual)
}

func TestOffsetGreaterThanReturnsFalseWhenLessThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:1}
	otherOffset := Offset{seconds:5}

	actual := offset.GreaterThan(&otherOffset)

	assert.False(t, actual)
}

func TestOffsetGreaterThanReturnsFalseWhenEqualToTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:2}
	otherOffset := Offset{seconds:2}

	actual := offset.GreaterThan(&otherOffset)

	assert.False(t, actual)
}

func TestOffsetGreaterThanReturnsTrueWhenGreaterThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:5}
	otherOffset := Offset{seconds:1}

	actual := offset.GreaterThan(&otherOffset)

	assert.True(t, actual)
}

func TestOffsetGreaterThanOrEqualsReturnsFalseWhenLessThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:1}
	otherOffset := Offset{seconds:5}

	actual := offset.GreaterThanOrEquals(&otherOffset)

	assert.False(t, actual)
}

func TestOffsetGreaterThanOrEqualsReturnsTrueWhenEqualToTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:2}
	otherOffset := Offset{seconds:2}

	actual := offset.GreaterThanOrEquals(&otherOffset)

	assert.True(t, actual)
}

func TestOffsetGreaterThanOrEqualsReturnsTrueWhenGreaterThanTheOtherOffset(t *testing.T) {
	offset := Offset{seconds:5}
	otherOffset := Offset{seconds:1}

	actual := offset.GreaterThanOrEquals(&otherOffset)

	assert.True(t, actual)
}

func TestOffsetEqualsReturnsFalseWhenTheOffsetsDiffer(t *testing.T) {
	offset := Offset{seconds:5}
	otherOffset := Offset{seconds:1}

	actual := offset.Equals(&otherOffset)

	assert.False(t, actual)
}