package godatime

import "godatime/preconditions"

type Offset struct {
	seconds int
}

const (
	minHours = -18
	maxHours = 18
	minSeconds = minHours * secondsPerHour
	maxSeconds = maxHours * secondsPerHour
	minMilliseconds = minHours * millisecondsPerHour
	maxMilliseconds = maxHours * millisecondsPerHour
	minTicks = minHours * ticksPerHour
	maxTicks = maxHours * ticksPerHour
	minNanoseconds = minHours * nanosecondsPerHour
	maxNanoseconds = maxHours * nanosecondsPerHour
)

func OffsetFromSeconds(seconds int) (*Offset, error) {
	if err := preconditions.CheckArgumentRange("seconds", seconds, minSeconds, maxSeconds); err != nil {
		return nil, err
	}

	return &Offset{seconds:seconds}, nil
}

func OffsetFromMilliseconds(milliseconds int) (*Offset, error) {
	if err := preconditions.CheckArgumentRange("milliseconds", milliseconds, minMilliseconds, maxMilliseconds); err != nil {
		return nil, err
	}

	return &Offset{seconds:(int)(milliseconds / millisecondsPerSecond)}, nil
}

func OffsetFromTicks(ticks int64) (*Offset, error) {
	if err := preconditions.CheckArgumentRangeInt64("ticks", ticks, minTicks, maxTicks); err != nil {
		return nil, err
	}

	return &Offset{seconds:(int)(ticks / ticksPerSecond)}, nil
}

func OffsetFromNanoseconds(nanoseconds int64) (*Offset, error) {
	if err := preconditions.CheckArgumentRangeInt64("nanoseconds", nanoseconds, minNanoseconds, maxNanoseconds); err != nil {
		return nil, err
	}

	return &Offset{seconds:(int)(nanoseconds / nanosecondsPerSecond)}, nil
}

func OffsetFromHours(hours int) (*Offset, error) {
	if err := preconditions.CheckArgumentRange("hours", hours, minHours, maxHours); err != nil {
		return nil, err
	}

	return &Offset{seconds:hours * secondsPerHour}, nil
}

func OffsetFromHoursAndMinutes(hours, minutes int) (*Offset, error) {
	return OffsetFromSeconds(hours *secondsPerHour + minutes *secondsPerMinute)
}

func OffsetMax(x, y *Offset) *Offset {
	if x.GreaterThan(y) {
		return x
	}

	return y
}

func OffsetMin(x, y *Offset) *Offset {
	if x.LessThan(y) {
		return x
	}

	return y
}

func (o *Offset) Seconds() int {
	return o.seconds
}

func (o *Offset) Milliseconds() int {
	return o.seconds * millisecondsPerSecond
}

func (o *Offset) Ticks() int64 {
	return (int64)(o.seconds) * ticksPerSecond
}

func (o *Offset) Nanoseconds() int64 {
	return (int64)(o.seconds) * nanosecondsPerSecond
}

func (o *Offset) Negate() *Offset {
	return &Offset{seconds:-o.seconds}
}

func (o *Offset) Add(other *Offset) (*Offset, error) {
	return OffsetFromSeconds(o.seconds + other.seconds)
}

func (o *Offset) Subtract(other *Offset) (*Offset, error) {
	return OffsetFromSeconds(o.seconds - other.seconds)
}

func (o *Offset) LessThan(other *Offset) bool {
	return o.seconds < other.seconds
}

func (o *Offset) LessThanOrEquals(other *Offset) bool {
	return o.seconds <= other.seconds
}

func (o *Offset) GreaterThan(other *Offset) bool {
	return o.seconds > other.seconds
}

func (o *Offset) GreaterThanOrEquals(other *Offset) bool {
	return o.seconds >= other.seconds
}

func (o *Offset) Equals(other *Offset) bool {
	return o.seconds == other.seconds
}
