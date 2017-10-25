package godatime

import (
	"testing"
	"math/rand"
	"github.com/stretchr/testify/assert"
)

func TestPeriodNanoseconds(t *testing.T) {
	expected := rand.Int63()
	period := Period{nanoseconds:expected}
	actual := period.Nanoseconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodTicks(t *testing.T) {
	expected := rand.Int63()
	period := Period{ticks:expected}
	actual := period.Ticks()

	assert.Equal(t, expected, actual)
}

func TestPeriodMilliseconds(t *testing.T) {
	expected := rand.Int63()
	period := Period{milliseconds:expected}
	actual := period.Milliseconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodSeconds(t *testing.T) {
	expected := rand.Int63()
	period := Period{seconds:expected}
	actual := period.Seconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodMinutes(t *testing.T) {
	expected := rand.Int63()
	period := Period{minutes:expected}
	actual := period.Minutes()

	assert.Equal(t, expected, actual)
}

func TestPeriodHours(t *testing.T) {
	expected := rand.Int63()
	period := Period{hours:expected}
	actual := period.Hours()

	assert.Equal(t, expected, actual)
}

func TestPeriodDays(t *testing.T) {
	expected := rand.Int()
	period := Period{days:expected}
	actual := period.Days()

	assert.Equal(t, expected, actual)
}

func TestPeriodWeeks(t *testing.T) {
	expected := rand.Int()
	period := Period{weeks:expected}
	actual := period.Weeks()

	assert.Equal(t, expected, actual)
}

func TestPeriodMonths(t *testing.T) {
	expected := rand.Int()
	period := Period{months:expected}
	actual := period.Months()

	assert.Equal(t, expected, actual)
}

func TestPeriodYears(t *testing.T) {
	expected := rand.Int()
	period := Period{years:expected}
	actual := period.Years()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddYears(t *testing.T) {
	initialPeriod := Period{years:rand.Int()}
	addPeriod := Period{years:rand.Int()}
	expected := initialPeriod.years + addPeriod.years

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Years()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddMonths(t *testing.T) {
	initialPeriod := Period{months:rand.Int()}
	addPeriod := Period{months:rand.Int()}
	expected := initialPeriod.months + addPeriod.months

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Months()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddWeeks(t *testing.T) {
	initialPeriod := Period{weeks:rand.Int()}
	addPeriod := Period{weeks:rand.Int()}
	expected := initialPeriod.weeks + addPeriod.weeks

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Weeks()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddDays(t *testing.T) {
	initialPeriod := Period{days:rand.Int()}
	addPeriod := Period{days:rand.Int()}
	expected := initialPeriod.days + addPeriod.days

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Days()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddHours(t *testing.T) {
	initialPeriod := Period{hours:rand.Int63()}
	addPeriod := Period{hours: rand.Int63()}
	expected := initialPeriod.hours + addPeriod.hours

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Hours()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddMinutes(t *testing.T) {
	initialPeriod := Period{minutes:rand.Int63()}
	addPeriod := Period{minutes:rand.Int63()}
	expected := initialPeriod.minutes + addPeriod.minutes

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Minutes()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddSeconds(t *testing.T) {
	initialPeriod := Period{seconds:rand.Int63()}
	addPeriod := Period{seconds:rand.Int63()}
	expected := initialPeriod.seconds + addPeriod.seconds

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Seconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddMilliseconds(t *testing.T) {
	initialPeriod := Period{milliseconds:rand.Int63()}
	addPeriod := Period{milliseconds:rand.Int63()}
	expected := initialPeriod.milliseconds + addPeriod.milliseconds

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Milliseconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddTicks(t *testing.T) {
	initialPeriod := Period{ticks:rand.Int63()}
	addPeriod := Period{ticks:rand.Int63()}
	expected := initialPeriod.ticks + addPeriod.ticks

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Ticks()

	assert.Equal(t, expected, actual)
}

func TestPeriodAddNanoseconds(t *testing.T) {
	initialPeriod := Period{nanoseconds:rand.Int63()}
	addPeriod := Period{nanoseconds:rand.Int63()}
	expected := initialPeriod.nanoseconds + addPeriod.nanoseconds

	newPeriod := initialPeriod.Add(&addPeriod)
	actual := newPeriod.Nanoseconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractYears(t *testing.T) {
	initialPeriod := Period{years:rand.Int()}
	subtractPeriod := Period{years:rand.Int()}
	expected := initialPeriod.years - subtractPeriod.years

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Years()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractMonths(t *testing.T) {
	initialPeriod := Period{months:rand.Int()}
	subtractPeriod := Period{months:rand.Int()}
	expected := initialPeriod.months - subtractPeriod.months

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Months()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractWeeks(t *testing.T) {
	initialPeriod := Period{weeks:rand.Int()}
	subtractPeriod := Period{weeks:rand.Int()}
	expected := initialPeriod.weeks - subtractPeriod.weeks

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Weeks()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractDays(t *testing.T) {
	initialPeriod := Period{days:rand.Int()}
	subtractPeriod := Period{days:rand.Int()}
	expected := initialPeriod.days - subtractPeriod.days

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Days()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractHours(t *testing.T) {
	initialPeriod := Period{hours:rand.Int63()}
	subtractPeriod := Period{hours:rand.Int63()}
	expected := initialPeriod.hours - subtractPeriod.hours

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Hours()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractMinutes(t *testing.T) {
	initialPeriod := Period{minutes:rand.Int63()}
	subtractPeriod := Period{minutes:rand.Int63()}
	expected := initialPeriod.minutes - subtractPeriod.minutes

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Minutes()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractSeconds(t *testing.T) {
	initialPeriod := Period{seconds:rand.Int63()}
	subtractPeriod := Period{seconds:rand.Int63()}
	expected := initialPeriod.seconds - subtractPeriod.seconds

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Seconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractMilliseconds(t *testing.T) {
	initialPeriod := Period{milliseconds:rand.Int63()}
	subtractPeriod := Period{milliseconds:rand.Int63()}
	expected := initialPeriod.milliseconds - subtractPeriod.milliseconds

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Milliseconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractTicks(t *testing.T) {
	initialPeriod := Period{ticks:rand.Int63()}
	subtractPeriod := Period{ticks:rand.Int63()}
	expected := initialPeriod.ticks - subtractPeriod.ticks

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Ticks()

	assert.Equal(t, expected, actual)
}

func TestPeriodSubtractNanoseconds(t *testing.T) {
	initialPeriod := Period{nanoseconds:rand.Int63()}
	subtractPeriod := Period{nanoseconds:rand.Int63()}
	expected := initialPeriod.nanoseconds - subtractPeriod.nanoseconds

	newPeriod := initialPeriod.Subtract(&subtractPeriod)
	actual := newPeriod.Nanoseconds()

	assert.Equal(t, expected, actual)
}

func TestPeriodHasDateComponentTrueWhenYearIsSet(t *testing.T) {
	period := Period{years:rand.Int()}

	actual := period.HasDateComponent()

	assert.True(t, actual)
}

func TestPeriodHasDateComponentTrueWhenMonthIsSet(t *testing.T) {
	period := Period{months:rand.Int()}

	actual := period.HasDateComponent()

	assert.True(t, actual)
}

func TestPeriodHasDateComponentTrueWhenDayIsSet(t *testing.T) {
	period := Period{days:rand.Int()}

	actual := period.HasDateComponent()

	assert.True(t, actual)
}

func TestPeriodHasDateComponentFalseWhenOnlyHoursIsSet(t *testing.T) {
	period := Period{hours:rand.Int63()}

	actual := period.HasDateComponent()

	assert.False(t, actual)
}

func TestPeriodHasDateComponentFalseWhenOnlyMinutesIsSet(t *testing.T) {
	period := Period{minutes:rand.Int63()}

	actual := period.HasDateComponent()

	assert.False(t, actual)
}

func TestPeriodHasDateComponentFalseWhenOnlySecondsIsSet(t *testing.T) {
	period := Period{seconds:rand.Int63()}

	actual := period.HasDateComponent()

	assert.False(t, actual)
}

func TestPeriodHasDateComponentFalseWhenOnlyMillisecondsIsSet(t *testing.T) {
	period := Period{milliseconds:rand.Int63()}

	actual := period.HasDateComponent()

	assert.False(t, actual)
}

func TestPeriodHasDateComponentFalseWhenOnlyTicksIsSet(t *testing.T) {
	period := Period{ticks:rand.Int63()}

	actual := period.HasDateComponent()

	assert.False(t, actual)
}

func TestPeriodHasDateComponentFalseWhenOnlyNanosecondsIsSet(t *testing.T) {
	period := Period{nanoseconds:rand.Int63()}

	actual := period.HasDateComponent()

	assert.False(t, actual)
}

func TestPeriodHasTimeComponentFalseWhenOnlyYearsIsSet(t *testing.T) {
	period := Period{years:rand.Int()}

	actual := period.HasTimeComponent()

	assert.False(t, actual)
}

func TestPeriodHasTimeComponentFalseWhenOnlyMonthsIsSet(t *testing.T) {
	period := Period{months:rand.Int()}

	actual := period.HasTimeComponent()

	assert.False(t, actual)
}

func TestPeriodHasTimeComponentFalseWhenOnlyDaysIsSet(t *testing.T) {
	period := Period{days:rand.Int()}

	actual := period.HasTimeComponent()

	assert.False(t, actual)
}

func TestPeriodHasTimeComponentTrueWhenHoursIsSet(t *testing.T) {
	period := Period{hours:rand.Int63()}

	actual := period.HasTimeComponent()

	assert.True(t, actual)
}

func TestPeriodHasTimeComponentTrueWhenMinutesIsSet(t *testing.T) {
	period := Period{minutes:rand.Int63()}

	actual := period.HasTimeComponent()

	assert.True(t, actual)
}

func TestPeriodHasTimeComponentTrueWhenSecondsIsSet(t *testing.T) {
	period := Period{seconds:rand.Int63()}

	actual := period.HasTimeComponent()

	assert.True(t, actual)
}

func TestPeriodHasTimeComponentTrueWhenMillisecondsIsSet(t *testing.T) {
	period := Period{milliseconds:rand.Int63()}

	actual := period.HasTimeComponent()

	assert.True(t, actual)
}

func TestPeriodHasTimeComponentTrueWhenTicksIsSet(t *testing.T) {
	period := Period{ticks:rand.Int63()}

	actual := period.HasTimeComponent()

	assert.True(t, actual)
}

func TestPeriodHasTimeComponentTrueWhenNanosecondsIsSet(t *testing.T) {
	period := Period{nanoseconds:rand.Int63()}

	actual := period.HasTimeComponent()

	assert.True(t, actual)
}

func TestPeriodNormalizeReturnsTheCorrectYear(t *testing.T) {
	period := Period{years:rand.Int()}

	normalized := period.Normalize()

	assert.Equal(t, period.Years(), normalized.Years())
}

func TestPeriodNormalizeReturnsTheCorrectMonths(t *testing.T) {
	period := Period{months:rand.Int()}

	normalized := period.Normalize()

	assert.Equal(t, period.Months(), normalized.Months())
}

func TestPeriodNormalizeReturnsZeroForTheWeeks(t *testing.T) {
	period := Period{weeks:rand.Int()}

	normalized := period.Normalize()

	assert.Zero(t, normalized.Weeks())
}

func TestPeriodNormalizeCorrectlyNormalizesTheDays(t *testing.T) {
	period := Period{weeks:rand.Int(),
		days:rand.Int(),
		hours:rand.Int63(),
		minutes:rand.Int63(),
		seconds:rand.Int63(),
		milliseconds:rand.Int63(),
		ticks:rand.Int63(),
		nanoseconds:rand.Int63()}

	totalNanoseconds := (int64)(period.nanoseconds + (int64)(period.ticks *nanosecondsPerTick) +
		(int64)(period.milliseconds *nanosecondsPerMillisecond) +
		(int64)(period.seconds *nanosecondsPerSecond) +
		(int64)(period.minutes *nanosecondsPerMinute) +
		(int64)(period.hours *nanosecondsPerHour) +
		(int64)((int64)(period.days) *nanosecondsPerDay) +
		(int64)((int64)(period.weeks) *nanosecondsPerWeek))

	expected := (int)(totalNanoseconds / nanosecondsPerDay)

	normalized := period.Normalize()

	assert.Equal(t, expected, normalized.Days())
}

func TestPeriodNormalizeCorrectlyNormalizesTheHours(t *testing.T) {
	period := Period{weeks:rand.Int(),
		days:rand.Int(),
		hours:rand.Int63(),
		minutes:rand.Int63(),
		seconds:rand.Int63(),
		milliseconds:rand.Int63(),
		ticks:rand.Int63(),
		nanoseconds:rand.Int63()}

	totalNanoseconds := (int64)(period.nanoseconds + (int64)(period.ticks *nanosecondsPerTick) +
		(int64)(period.milliseconds *nanosecondsPerMillisecond) +
		(int64)(period.seconds *nanosecondsPerSecond) +
		(int64)(period.minutes *nanosecondsPerMinute) +
		(int64)(period.hours *nanosecondsPerHour) +
		(int64)((int64)(period.days) *nanosecondsPerDay) +
		(int64)((int64)(period.weeks) *nanosecondsPerWeek))

	expected := (totalNanoseconds / nanosecondsPerHour) % hoursPerDay

	normalized := period.Normalize()

	assert.Equal(t, expected, normalized.Hours())
}

func TestPeriodNormalizeCorrectlyNormalizesTheMinutes(t *testing.T) {
	period := Period{weeks:rand.Int(),
		days:rand.Int(),
		hours:rand.Int63(),
		minutes:rand.Int63(),
		seconds:rand.Int63(),
		milliseconds:rand.Int63(),
		ticks:rand.Int63(),
		nanoseconds:rand.Int63()}

	totalNanoseconds := (int64)(period.nanoseconds + (int64)(period.ticks *nanosecondsPerTick) +
		(int64)(period.milliseconds *nanosecondsPerMillisecond) +
		(int64)(period.seconds *nanosecondsPerSecond) +
		(int64)(period.minutes *nanosecondsPerMinute) +
		(int64)(period.hours *nanosecondsPerHour) +
		(int64)((int64)(period.days) *nanosecondsPerDay) +
		(int64)((int64)(period.weeks) *nanosecondsPerWeek))

	expected := (totalNanoseconds / nanosecondsPerMinute) %  minutesPerHour

	normalized := period.Normalize()

	assert.Equal(t, expected, normalized.Minutes())
}

func TestPeriodNormalizeCorrectlyNormalizesTheSeconds(t *testing.T) {
	period := Period{weeks:rand.Int(),
		days:rand.Int(),
		hours:rand.Int63(),
		minutes:rand.Int63(),
		seconds:rand.Int63(),
		milliseconds:rand.Int63(),
		ticks:rand.Int63(),
		nanoseconds:rand.Int63()}

	totalNanoseconds := (int64)(period.nanoseconds + (int64)(period.ticks *nanosecondsPerTick) +
		(int64)(period.milliseconds *nanosecondsPerMillisecond) +
		(int64)(period.seconds *nanosecondsPerSecond) +
		(int64)(period.minutes *nanosecondsPerMinute) +
		(int64)(period.hours *nanosecondsPerHour) +
		(int64)((int64)(period.days) *nanosecondsPerDay) +
		(int64)((int64)(period.weeks) *nanosecondsPerWeek))

	expected := (totalNanoseconds / nanosecondsPerSecond) % secondsPerMinute

	normalized := period.Normalize()

	assert.Equal(t, expected, normalized.Seconds())
}

func TestPeriodNormalizeCorrectlyNormalizesTheMilliseconds(t *testing.T) {
	period := Period{weeks:rand.Int(),
		days:rand.Int(),
		hours:rand.Int63(),
		minutes:rand.Int63(),
		seconds:rand.Int63(),
		milliseconds:rand.Int63(),
		ticks:rand.Int63(),
		nanoseconds:rand.Int63()}

	totalNanoseconds := (int64)(period.nanoseconds + (int64)(period.ticks *nanosecondsPerTick) +
		(int64)(period.milliseconds *nanosecondsPerMillisecond) +
		(int64)(period.seconds *nanosecondsPerSecond) +
		(int64)(period.minutes *nanosecondsPerMinute) +
		(int64)(period.hours *nanosecondsPerHour) +
		(int64)((int64)(period.days) *nanosecondsPerDay) +
		(int64)((int64)(period.weeks) *nanosecondsPerWeek))

	expected := (totalNanoseconds / nanosecondsPerMillisecond) % millisecondsPerSecond

	normalized := period.Normalize()

	assert.Equal(t, expected, normalized.Milliseconds())
}

func TestPeriodNormalizeReturnsZeroFroTheTicks(t *testing.T) {
	period := Period{ticks:rand.Int63()}

	normalized := period.Normalize()

	assert.Zero(t, normalized.Ticks())
}

func TestPeriodNormalizeCorrectlyNormalizesTheNanoseconds(t *testing.T) {
	period := Period{weeks:rand.Int(),
		days:rand.Int(),
		hours:rand.Int63(),
		minutes:rand.Int63(),
		seconds:rand.Int63(),
		milliseconds:rand.Int63(),
		ticks:rand.Int63(),
		nanoseconds:rand.Int63()}

	totalNanoseconds := (int64)(period.nanoseconds + (int64)(period.ticks *nanosecondsPerTick) +
		(int64)(period.milliseconds *nanosecondsPerMillisecond) +
		(int64)(period.seconds *nanosecondsPerSecond) +
		(int64)(period.minutes *nanosecondsPerMinute) +
		(int64)(period.hours *nanosecondsPerHour) +
		(int64)((int64)(period.days) *nanosecondsPerDay) +
		(int64)((int64)(period.weeks) *nanosecondsPerWeek))

	expected := totalNanoseconds % nanosecondsPerMillisecond

	normalized := period.Normalize()

	assert.Equal(t, expected, normalized.Nanoseconds())
}

func TestPeriodEqualsReturnsFalseIfTheOtherPeriodIsNil(t *testing.T) {
	period := Period{}

	actual := period.Equals(nil)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfYearsDiffer(t *testing.T) {
	period := Period{years:1}
	otherPeriod := Period{years:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfMonthsDiffer(t *testing.T) {
	period := Period{months:1}
	otherPeriod := Period{months:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfWeeksDiffer(t *testing.T) {
	period := Period{weeks:1}
	otherPeriod := Period{weeks:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfDaysDiffer(t *testing.T) {
	period := Period{days:1}
	otherPeriod := Period{days:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfHoursDiffer(t *testing.T) {
	period := Period{hours:1}
	otherPeriod := Period{hours:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfMinutesDiffer(t *testing.T) {
	period := Period{minutes:1}
	otherPeriod := Period{minutes:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfSecondsDiffer(t *testing.T) {
	period := Period{seconds:1}
	otherPeriod := Period{seconds:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfMillisecondsDiffer(t *testing.T) {
	period := Period{milliseconds:1}
	otherPeriod := Period{milliseconds:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfTicksDiffer(t *testing.T) {
	period := Period{ticks:1}
	otherPeriod := Period{ticks:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}

func TestPeriodEqualsIsFalseIfNanosecondsDiffer(t *testing.T) {
	period := Period{nanoseconds:1}
	otherPeriod := Period{nanoseconds:2}

	actual := period.Equals(&otherPeriod)

	assert.False(t, actual)
}