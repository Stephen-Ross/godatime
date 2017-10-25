package godatime

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLocalTimeMaxReturnsTheLargerTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:1}
	otherLocalTime := LocalTime{nanoseconds:2}

	actual := LocalTimeMax(&localTime, &otherLocalTime)

	assert.Equal(t, otherLocalTime, *actual)
}

func TestLocalTimeMinReturnsSmallerTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:1}
	otherLocalTime := LocalTime{nanoseconds:2}

	actual := LocalTimeMin(&localTime, &otherLocalTime)

	assert.Equal(t, localTime, *actual)
}

func TestLocalTimeHourReturnsTheHourOfTheTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour + nanosecondsPerMinute}

	actual := localTime.Hour()

	assert.Equal(t, 1, actual)
}

func TestLocalTimeClockHourOfHalfDayReturnsTheTwelveHourTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour * 23}

	actual := localTime.ClockHourOfHalfDay()

	assert.Equal(t, 11, actual)
}

func TestLocalTimeMinuteReturnsTheMinutesOfTheTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour + nanosecondsPerMinute * 5}

	actual := localTime.Minute()

	assert.Equal(t, 5, actual)
}

func TestLocalTimeSecondReturnsTheSecondsOfTheTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMinute + nanosecondsPerSecond * 4}

	actual := localTime.Second()

	assert.Equal(t, 4, actual)
}

func TestLocalTimeMillisecondReturnsTheMillisecondsOfTheTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerSecond + nanosecondsPerMillisecond * 100}

	actual := localTime.Millisecond()

	assert.Equal(t, 100, actual)
}

func TestLocalTimeTickOfDayReturnsTheTicksOfTheTime(t *testing.T) {
	localTime, _ := NewLocalTimeTicksSinceMidnight(ticksPerMinute * 5)

	actual := localTime.TickOfDay()

	assert.Equal(t, ticksPerMinute* 5, actual)
}

func TestLocalTimeTickOfSecondReturnsTheTicksOfTheCurrentSecond(t *testing.T) {
	localTime, _ := NewLocalTimeTicksSinceMidnight(ticksPerSecond / 2)

	actual := localTime.TickOfSecond()

	assert.Equal(t, (int)(ticksPerSecond / 2), actual)
}

func TestLocalTimeNanosecondOfDayReturnsTheNanosecondsOfTheTime(t *testing.T) {
	localTime, _ := NewLocalTimeNanoseconds(nanosecondsPerMinute * 5)

	actual := localTime.NanosecondOfDay()

	assert.Equal(t, nanosecondsPerMinute * 5, actual)
}

func TestLocalTimeNanosecondOfSecondReturnsTheNanosecondsOfTheCurrentSecond(t *testing.T) {
	localTime, _ := NewLocalTimeNanoseconds(nanosecondsPerSecond / 2)

	actual := localTime.NanosecondOfSecond()

	assert.Equal(t, (int)(nanosecondsPerSecond / 2), actual)
}

func TestLocalTimePlusHoursAddsToTheHours(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour * 3}

	newTime := localTime.PlusHours(3)

	actual := newTime.Hour()

	assert.Equal(t, 6, actual)
}

func TestLocalTimePlusMinutesAddsTheMinutes(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMinute * 30}

	newTime := localTime.PlusMinutes(10)

	actual := newTime.Minute()

	assert.Equal(t, 40, actual)
}

func TestLocalTimePlusSecondsAddTheSeconds(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerSecond * 30}

	newTime := localTime.PlusSeconds(10)

	actual := newTime.Second()

	assert.Equal(t, 40, actual)
}

func TestLocalTimePlusMillisecondsAddsTheMilliseconds(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMillisecond * 100}

	newTime := localTime.PlusMilliseconds(100)

	actual := newTime.Millisecond()

	assert.Equal(t, 200, actual)
}

func TestLocalTimePlusTicksAddsTheTicks(t *testing.T) {
	localTime, _ := NewLocalTimeTicksSinceMidnight(40)

	newTime := localTime.PlusTicks(100)

	actual := newTime.TickOfDay()

	assert.Equal(t, int64(140), actual)
}

func TestLocalTimePlusNanosecondsAddsTheNanoseconds(t *testing.T) {
	localTime, _ := NewLocalTimeNanoseconds(100)

	newTime := localTime.PlusNanoseconds(100)

	actual := newTime.NanosecondOfDay()

	assert.Equal(t, int64(200), actual)
}

func TestLocalTimeAddCorrectlyAddsThePeriodsHours(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	period := Period{hours:5}

	newTime := localTime.Add(&period)

	actual := newTime.Hour()

	assert.Equal(t, 6, actual)
}

func TestLocalTimeAddCorrectlyAddsThePeriodsMinutes(t *testing.T) {
	localTime := LocalTime{}
	period := Period{minutes:30}

	newTime := localTime.Add(&period)

	actual := newTime.Minute()

	assert.Equal(t, 30, actual)
}

func TestLocalTimeAddCorrectlyAddsThePeriodsSeconds(t *testing.T) {
	localTime := LocalTime{}
	period := Period{seconds:30}

	newTime := localTime.Add(&period)

	actual := newTime.Second()

	assert.Equal(t, 30, actual)
}

func TestLocalTimeAddCorrectlyAddsThePeriodsMilliseconds(t *testing.T) {
	localTime := LocalTime{}
	period := Period{milliseconds:100}

	newTime := localTime.Add(&period)

	actual := newTime.Millisecond()

	assert.Equal(t, 100, actual)
}

func TestLocalTimeAddCorrectlyAddsThePeriodsTicks(t *testing.T) {
	localTime := LocalTime{}
	period := Period{ticks:100}

	newTime := localTime.Add(&period)

	actual := newTime.TickOfDay()

	assert.Equal(t, (int64)(100), actual)
}

func TestLocalTimeAddCorrectlyAddsThePeriodsNanoseconds(t *testing.T) {
	localTime := LocalTime{}
	period := Period{nanoseconds:1000}

	newTime := localTime.Add(&period)

	actual := newTime.NanosecondOfDay()

	assert.Equal(t, (int64)(1000), actual)
}

func TestLocalTimeSubtractCorrectlyCalculatesTheHoursBetweenTheTimes(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour*3}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour*5}

	delta := localTime.Subtract(&otherTime)

	actual := delta.Hours()

	assert.Equal(t, int64(2), actual)
}

func TestLocalTimeSubtractCorrectlyCalculatesTheMinutesBetweenTheTimes(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMinute*20}
	otherTime := LocalTime{nanoseconds:nanosecondsPerMinute*30}

	delta := localTime.Subtract(&otherTime)

	actual := delta.Minutes()

	assert.Equal(t, (int64)(10), actual)
}

func TestLocalTimeSubtractCorrectlyCalculatesTheSecondsBetweenTheTimes(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerSecond*15}
	otherTime := LocalTime{nanoseconds:nanosecondsPerSecond*30}

	delta := localTime.Subtract(&otherTime)

	actual := delta.Seconds()

	assert.Equal(t, (int64)(15), actual)
}

func TestLocalTimeSubtractCorrectlyCalculatesTheMillisecondsBetweenTheTimes(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMillisecond*100}
	otherTime := LocalTime{nanoseconds:nanosecondsPerMillisecond*300}

	delta := localTime.Subtract(&otherTime)

	actual := delta.Milliseconds()

	assert.Equal(t, (int64)(200), actual)
}

func TestLocalTimeSubtractCorrectlyCalculatesTheTicksBetweenTheTimes(t *testing.T) {
	localTime, _ := NewLocalTimeTicksSinceMidnight(300)
	otherTime, _ := NewLocalTimeTicksSinceMidnight(700)

	delta := localTime.Subtract(otherTime)

	actual := delta.Ticks()

	assert.Equal(t, (int64)(400), actual)
}

func TestLocalTimeEqualsReturnsFalseIfTheNanosecondsDiffer(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerMinute}

	actual := localTime.Equals(&otherTime)

	assert.False(t, actual)
}

func TestLocalTimeLessThanReturnsFalseWhenGreaterThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerMinute}

	actual := localTime.LessThan(&otherTime)

	assert.False(t, actual)
}

func TestLocalTimeLessThanReturnsTrueWhenLessThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMinute}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.LessThan(&otherTime)

	assert.True(t, actual)
}

func TestLocalTimeLessThanReturnsFalseWhenEqualToTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.LessThan(&otherTime)

	assert.False(t, actual)
}

func TestLocalTimeLessThanOrEqualsReturnsFalseWhenGreaterThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerMinute}

	actual := localTime.LessThanOrEquals(&otherTime)

	assert.False(t, actual)
}

func TestLocalTimeLessThanOrEqualsReturnsTrueWhenLessThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMinute}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.LessThanOrEquals(&otherTime)

	assert.True(t, actual)
}

func TestLocalTimeLessThanOrEqualsReturnsTrueWhenEqualToTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.LessThanOrEquals(&otherTime)

	assert.True(t, actual)
}

func TestLocalTimeGreaterThanReturnsFalseWhenLessThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMinute}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.GreaterThan(&otherTime)

	assert.False(t, actual)
}

func TestLocalTimeGreaterThanReturnsFalseWhenEqualToTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.GreaterThan(&otherTime)

	assert.False(t, actual)
}

func TestLocalTimeGreaterThanReturnsTrueWhenGreaterThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerMinute}

	actual := localTime.GreaterThan(&otherTime)

	assert.True(t, actual)
}

func TestLocalTimeGreaterThanOrEqualsReturnsFalseWhenLessThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerMinute}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.GreaterThanOrEquals(&otherTime)

	assert.False(t, actual)
}

func TestLocalTimeGreaterThanOrEqualsReturnsTrueWhenEqualToTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerHour}

	actual := localTime.GreaterThanOrEquals(&otherTime)

	assert.True(t, actual)
}

func TestLocalTimeGreaterThanOrEqualsReturnsTrueWhenGreaterThanTheOtherTime(t *testing.T) {
	localTime := LocalTime{nanoseconds:nanosecondsPerHour}
	otherTime := LocalTime{nanoseconds:nanosecondsPerMinute}

	actual := localTime.GreaterThanOrEquals(&otherTime)

	assert.True(t, actual)
}