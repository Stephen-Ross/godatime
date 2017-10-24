package godatime

type Offset struct {
	seconds int
}

func OffsetFromSeconds(seconds int) *Offset {
	return &Offset{seconds:seconds}
}

func OffsetFromMilliseconds(milliseconds int) *Offset {
	return &Offset{seconds:(int)(milliseconds / millisecondsPerSecond)}
}

func OffsetFromTicks(ticks int64) *Offset {
	return &Offset{seconds:(int)(ticks / ticksPerSecond)}
}

func OffsetFromNanoseconds(nanoseconds int64) *Offset {
	return &Offset{seconds:(int)(nanoseconds / nanosecondsPerSecond)}
}

func OffsetFromHours(hours int) *Offset {
	return &Offset{seconds:hours * secondsPerHour}
}

func OffsetFromHoursAndMinutes(hours, minutes int) *Offset {
	return &Offset{seconds:hours *secondsPerHour + minutes *secondsPerMinute}
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

func (o *Offset) Add(other *Offset) *Offset {
	return OffsetFromSeconds(o.seconds + other.seconds)
}

func (o *Offset) Subtract(other *Offset) *Offset {
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
