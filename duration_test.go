package godatime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
)

func TestMaxDurationReturnsTheMaximumDuration(t *testing.T) {
	duration := Duration{days:1, nanosecondsOfDay:nanosecondsPerHour * 4}
	otherDuration := Duration{nanosecondsOfDay:nanosecondsPerHour * 10}

	actual := MaxDuration(&duration, &otherDuration)

	assert.Equal(t, duration, *actual)
}

func TestMinDurationReturnsTheMinimumDuration(t *testing.T) {
	duration := Duration{days:1, nanosecondsOfDay:nanosecondsPerHour * 4}
	otherDuration := Duration{nanosecondsOfDay:nanosecondsPerHour * 10}

	actual := MinDuration(&duration, &otherDuration)

	assert.Equal(t, otherDuration, *actual)
}

func TestDurationDaysReturnsTheCorrectDaysWhenDaysIsSet(t *testing.T) {
	duration := Duration{days:1}

	actual := duration.Days()

	assert.Equal(t, 1, actual)
}

func TestDurationDaysReturnsZeroWhenTheDaysAndNanosecondsAreNotSet(t *testing.T) {
	duration := Duration{}

	actual := duration.Days()

	assert.Equal(t, 0, actual)
}

func TestDurationDaysReturnsClosesToZeroWhenDaysAreNegative(t *testing.T) {
	duration := Duration{days:-5, nanosecondsOfDay:nanosecondsPerHour * rand.Int63()}

	actual := duration.Days()

	assert.Equal(t, -4, actual)
}

func TestDurationNanosecondOfDayReturnsTheSetValueIfDaysIsNonNegative(t *testing.T) {
	duration := Duration{days:4, nanosecondsOfDay:nanosecondsPerHour * 5}

	actual := duration.NanosecondOfDay()

	assert.Equal(t, nanosecondsPerHour * 5, actual)
}

func TestDurationNanosecondOfDayReturnsZeroIfNanosecondsIsNotSet(t *testing.T) {
	duration := Duration{days:4}

	actual := duration.NanosecondOfDay()

	assert.Equal(t, int64(0), actual)
}

func TestDurationNanosecondsReturnsTheValueAwayFromTheStartOfTheDayWhenDaysIsNegative(t *testing.T) {
	duration := Duration{days:-1, nanosecondsOfDay:nanosecondsPerHour * 5}

	actual := duration.NanosecondOfDay()

	assert.Equal(t, nanosecondsPerHour * 5 - nanosecondsPerDay, actual)
}

func TestDurationHoursReturnsTheCorrectNumberOfHoursTruncatedTowardsZero(t *testing.T) {
	duration := Duration{days:5, nanosecondsOfDay:nanosecondsPerHour*3 + nanosecondsPerMinute*30}

	actual := duration.Hours()

	assert.Equal(t, 3, actual)
}

func TestDurationMinutesReturnsTheCorrectMinutesTruncatedTowardsZero(t *testing.T) {
	duration := Duration{days:5, nanosecondsOfDay:nanosecondsPerHour*3 + nanosecondsPerMinute*30}

	actual := duration.Minutes()

	assert.Equal(t, 30, actual)
}

func TestDurationSecondsReturnsTheCorrectSecondsTruncatedTowardsZero(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerMinute*30 + nanosecondsPerSecond*15 + nanosecondsPerMillisecond*100}

	actual := duration.Seconds()

	assert.Equal(t, 15, actual)
}

func TestDurationMillisecondsReturnsTheCorrectMillisecondsTruncatedTowardsZero(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerSecond*15 + nanosecondsPerMillisecond*100 + nanosecondsPerTick*4000}

	actual := duration.Milliseconds()

	assert.Equal(t, 100, actual)
}

func TestDurationSubsecondTicksReturnsTheCorrectTicksTruncatedTowardsZero(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerSecond*10 + nanosecondsPerTick*40000}

	actual := duration.SubsecondTicks()

	assert.Equal(t, 40000, actual)
}

func TestDurationSubsecondNanosecondsReturnsTheCorrectNanosecondsTruncatedTowardsZero(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerSecond*10 + 80000000}

	actual := duration.SubsecondNanoseconds()

	assert.Equal(t, 80000000, actual)
}

func TestDurationTotalDaysReturnsTheTotalDays(t *testing.T) {
	duration := Duration{days:2, nanosecondsOfDay:nanosecondsPerHour*12}

	actual := duration.TotalDays()

	assert.Equal(t, 2.5, actual)
}

func TestDurationTotalHoursReturnsTheTotalHours(t *testing.T) {
	duration := Duration{days:1, nanosecondsOfDay:nanosecondsPerHour*5 + nanosecondsPerMinute*15}

	actual := duration.TotalHours()

	assert.Equal(t, 29.25, actual)
}

func TestDurationTotalMinutesReturnsTheTotalMinutes(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerHour*3 + nanosecondsPerMinute*10 + nanosecondsPerSecond*30}

	actual := duration.TotalMinutes()

	assert.Equal(t, 190.5, actual)
}

func TestDurationTotalSecondsReturnsTheTotalSeconds(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerMinute*10 + nanosecondsPerSecond*15 + nanosecondsPerMillisecond*500}

	actual := duration.TotalSeconds()

	assert.Equal(t, 615.5, actual)
}

func TestDurationTotalMillisecondsReturnsTheTotalMilliseconds(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerSecond*5 + nanosecondsPerMillisecond*100 + nanosecondsPerTick*5000}

	actual := duration.TotalMilliseconds()

	assert.Equal(t, 5100.5, actual)
}

func TestDurationTotalTicksReturnsTheTotalTicks(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerMillisecond*5 + nanosecondsPerTick*3000}

	actual := duration.TotalTicks()

	assert.Equal(t, 53000.0, actual)
}

func TestDurationTotalNanosecondsReturnsTheTotalNanoseconds(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerTick*3000 + 10000}

	actual := duration.TotalNanoseconds()

	assert.Equal(t, 310000.0, actual)
}

func TestDurationPlusCorrectlyAddsTheDays(t *testing.T) {
	duration := Duration{days:3}
	otherDuration := Duration{days:2}

	newDuration := duration.Plus(&otherDuration)

	assert.Equal(t, 5, newDuration.Days())
}

func TestDurationPlusCorrectlyIncrementsTheDayIfTheTwoTimesIntoDayMakeAnotherDay(t *testing.T) {
	duration := Duration{nanosecondsOfDay:nanosecondsPerHour * 12}
	otherDuration := Duration{nanosecondsOfDay:nanosecondsPerHour * 12}

	newDuration := duration.Plus(&otherDuration)

	assert.Equal(t, 1, newDuration.Days())
	assert.Equal(t, int64(0), newDuration.NanosecondOfDay())
}

func TestDurationMinusCorrectlyDecrementsTheDays(t *testing.T) {
	duration := Duration{days:3}
	otherDuration := Duration{days:2}

	newDuration := duration.Minus(&otherDuration)

	assert.Equal(t, 1, newDuration.Days())
}

func TestDurationMinusCorrectlyDecrementsTheDayIfTheTwoTimesIntoDayLoseADay(t *testing.T) {
	duration := Duration{days:2, nanosecondsOfDay:nanosecondsPerHour * 12}
	otherDuration := Duration{nanosecondsOfDay:nanosecondsPerHour * 23}

	newDuration := duration.Minus(&otherDuration)

	assert.Equal(t, 1, newDuration.Days())
	assert.Equal(t, 13, newDuration.Hours())
}