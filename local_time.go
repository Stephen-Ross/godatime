package godatime

import "godatime/preconditions"

type LocalTime struct {
	nanoseconds int64
}

func NewLocalTime(builder *LocalTimeBuilder) (*LocalTime, error) {
	if err := preconditions.CheckArgumentRangeInt64("Hour", builder.Hour, 0, hoursPerDay - 1); err != nil {
		return nil, err
	}

	if err := preconditions.CheckArgumentRangeInt64("Minute", builder.Minute, 0, minutesPerHour - 1); err != nil {
		return nil, err
	}

	if err := preconditions.CheckArgumentRangeInt64("Second", builder.Second, 0, secondsPerMinute - 1); err != nil {
		return nil, err
	}

	if err := preconditions.CheckArgumentRangeInt64("millisecond", builder.Millisecond, 0, millisecondsPerSecond - 1); err != nil {
		return nil, err
	}

	return &LocalTime{nanoseconds:builder.Hour * nanosecondsPerHour + builder.Minute *
		nanosecondsPerMinute + builder.Second * nanosecondsPerSecond + builder.Millisecond * nanosecondsPerMillisecond}, nil
}

func NewLocalTimeNanoseconds(nanoseconds int64) (*LocalTime, error) {
	if err := preconditions.CheckArgumentRangeInt64("nanoseconds", nanoseconds, 0, nanosecondsPerDay - 1); err != nil {
		return nil, err
	}

	return &LocalTime{nanoseconds:nanoseconds}, nil
}

func NewLocalTimeTicksSinceMidnight(ticks int64) (*LocalTime, error) {
	if err := preconditions.CheckArgumentRangeInt64("ticks", ticks, 0, ticksPerDay - 1); err != nil {
		return nil, err
	}

	return &LocalTime{nanoseconds:ticks * nanosecondsPerTick}, nil
}

func LocalTimeMax(x, y *LocalTime) *LocalTime {
	if x.GreaterThan(y) {
		return x
	}

	return y
}

func LocalTimeMin(x, y *LocalTime) *LocalTime {
	if x.LessThan(y) {
		return x
	}

	return y
}

func (l *LocalTime) Hour() int {
	return (int)((l.nanoseconds >> 13) / 439453125)
}

func (l *LocalTime) ClockHourOfHalfDay() int {
	hourOfHalfDay := l.Hour() % 12
	if hourOfHalfDay == 0 {
		return 12
	}

	return hourOfHalfDay
}

func (l *LocalTime) Minute() int {
	return (int)(((l.nanoseconds >> 11) / 29296875) % minutesPerHour)
}

func (l *LocalTime) Second() int {
	return (int)((l.nanoseconds / nanosecondsPerSecond) % secondsPerMinute)
}

func (l *LocalTime) Millisecond() int {
	return (int)((l.nanoseconds / nanosecondsPerMillisecond) % millisecondsPerSecond)
}

func (l *LocalTime) TickOfSecond() int {
	return (int)(l.TickOfDay() % ticksPerSecond)
}

func (l *LocalTime) TickOfDay() int64 {
	return l.nanoseconds / nanosecondsPerTick
}

func (l *LocalTime) NanosecondOfSecond() int {
	return (int)(l.nanoseconds % nanosecondsPerSecond)
}

func (l *LocalTime) NanosecondOfDay() int64 {
	return l.nanoseconds
}

func (l *LocalTime) PlusHours(hours int64) *LocalTime {
	return timePeriodHours.AddValue(l.nanoseconds, hours)
}

func (l *LocalTime) PlusMinutes(minutes int64) *LocalTime {
	return timePeriodMinutes.AddValue(l.nanoseconds, minutes)
}

func (l *LocalTime) PlusSeconds(seconds int64) *LocalTime {
	return timePeriodSeconds.AddValue(l.nanoseconds, seconds)
}

func (l *LocalTime) PlusMilliseconds(milliseconds int64) *LocalTime {
	return timePeriodMilliseconds.AddValue(l.nanoseconds, milliseconds)
}

func (l *LocalTime) PlusTicks(ticks int64) *LocalTime {
	return timePeriodTicks.AddValue(l.nanoseconds, ticks)
}

func (l *LocalTime) PlusNanoseconds(nanoseconds int64) *LocalTime {
	return timePeriodNanoseconds.AddValue(l.nanoseconds, nanoseconds)
}

func (l *LocalTime) Add(period *Period) *LocalTime {
	return l.PlusHours(period.hours).
		PlusMinutes(period.minutes).
		PlusSeconds(period.seconds).
		PlusMilliseconds(period.milliseconds).
		PlusTicks(period.ticks).
		PlusNanoseconds(period.nanoseconds)
}

func (l *LocalTime) Subtract(other *LocalTime) *Period {
	return Between(l, other)
}

func (l *LocalTime) Equals(other *LocalTime) bool {
	return l.nanoseconds == other.nanoseconds
}

func (l *LocalTime) LessThan(other *LocalTime) bool {
	return l.nanoseconds < other.nanoseconds
}

func (l *LocalTime) LessThanOrEquals(other *LocalTime) bool {
	return l.nanoseconds <= other.nanoseconds
}

func (l *LocalTime) GreaterThan(other *LocalTime) bool {
	return l.nanoseconds > other.nanoseconds
}

func (l *LocalTime) GreaterThanOrEquals(other *LocalTime) bool {
	return l.nanoseconds >= other.nanoseconds
}