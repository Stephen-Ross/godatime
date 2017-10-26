package godatime

import "math"

type Duration struct {
	days int
	nanosecondsOfDay int64
}

func NewDuration(builder *DurationBuilder) *Duration {
	daysNanoseconds := int64(builder.Days) * nanosecondsPerDay
	hoursNanoseconds := builder.Hours * nanosecondsPerHour
	minutesNanoseconds := builder.Minutes * nanosecondsPerMinute
	secondsNanoseconds := builder.Seconds * nanosecondsPerSecond
	millisecondsNanoseconds := builder.Milliseconds * nanosecondsPerMillisecond

	return DurationFromNanoseconds(daysNanoseconds + hoursNanoseconds + minutesNanoseconds + secondsNanoseconds + millisecondsNanoseconds)
}

func DurationFromTicks(ticks int64) *Duration {
	var days int64
	var ticksOfDay int64

	if ticks >= 0 {
		days = ticks / ticksPerDay
		ticksOfDay = ticks - days * ticksPerDay
	} else {
		days = ((ticks + 1) / ticksPerDay) - 1
		ticksOfDay = ticks - (days + 1) * ticksPerDay + ticksPerDay
	}

	return &Duration{days:int(days), nanosecondsOfDay:ticksOfDay*nanosecondsPerTick}
}

func DurationFromNanoseconds(nanoseconds int64) *Duration {
	var days int64
	var nanosecondsOfDay int64

	if nanoseconds >= 0 {
		days = nanoseconds / nanosecondsPerDay
	} else {
		days = ((nanoseconds + 1) / nanosecondsPerDay) + 1
	}

	if nanoseconds >= math.MinInt64 + nanosecondsPerDay {
		nanosecondsOfDay = nanoseconds - days * nanosecondsPerDay
	} else {
		nanosecondsOfDay = nanoseconds - (days + 1) * nanosecondsPerDay + nanosecondsPerDay
	}

	return &Duration{days:int(days), nanosecondsOfDay:nanosecondsOfDay}
}

func MaxDuration(x, y *Duration) *Duration {
	if x.GreaterThan(y) {
		return x
	}

	return y
}

func MinDuration(x, y *Duration ) *Duration {
	if x.LessThan(y) {
		return x
	}

	return y
}

func (d *Duration) Days() int {
	if d.days >= 0 || d.nanosecondsOfDay == 0 {
		return d.days
	}

	return d.days + 1
}

func (d *Duration) NanosecondOfDay() int64 {
	if d.days >= 0 {
		return d.nanosecondsOfDay
	}

	if d.nanosecondsOfDay == 0 {
		return 0
	}
	return d.nanosecondsOfDay - nanosecondsPerDay
}

func (d *Duration) Hours() int {
	return (int)(d.NanosecondOfDay() / nanosecondsPerHour)
}

func (d *Duration) Minutes() int {
	return (int)((d.NanosecondOfDay() / nanosecondsPerMinute) % minutesPerHour)
}

func (d *Duration) Seconds() int {
	return (int)((d.NanosecondOfDay() / nanosecondsPerSecond) % secondsPerMinute)
}

func (d *Duration) Milliseconds() int {
	return (int)((d.NanosecondOfDay() / nanosecondsPerMillisecond) % millisecondsPerSecond)
}

func (d *Duration) SubsecondTicks() int {
	return (int)((d.NanosecondOfDay() / nanosecondsPerTick) % ticksPerSecond)
}

func (d *Duration) SubsecondNanoseconds() int {
	return (int)(d.NanosecondOfDay() % nanosecondsPerSecond)
}

func (d *Duration) TotalDays() float64 {
	return d.TotalNanoseconds() / float64(nanosecondsPerDay)
}

func (d *Duration) TotalHours() float64 {
	return d.TotalNanoseconds() / float64(nanosecondsPerHour)
}

func (d *Duration) TotalMinutes() float64 {
	return d.TotalNanoseconds() / float64(nanosecondsPerMinute)
}

func (d *Duration) TotalSeconds() float64 {
	return d.TotalNanoseconds() / float64(nanosecondsPerSecond)
}

func (d *Duration) TotalMilliseconds() float64 {
	return d.TotalNanoseconds() / float64(nanosecondsPerMillisecond)
}

func (d *Duration) TotalTicks() float64 {
	return d.TotalNanoseconds() / float64(nanosecondsPerTick)
}

func (d *Duration) TotalNanoseconds() float64 {
	return float64((float64(d.days) * float64(nanosecondsPerDay)) + float64(d.nanosecondsOfDay))
}

func (d *Duration) Plus(right *Duration) *Duration {
	newDays := d.days + right.days
	newNanoseconds := d.nanosecondsOfDay + right.nanosecondsOfDay
	if newNanoseconds >= nanosecondsPerDay {
		newDays++
		newNanoseconds -= nanosecondsPerDay
	}

	return &Duration{days:newDays, nanosecondsOfDay:newNanoseconds}
}

func (d *Duration) Minus(right *Duration) *Duration {
	newDays := d.days - right.days
	newNanoseconds := d.nanosecondsOfDay - right.nanosecondsOfDay
	if newNanoseconds < 0 {
		newDays--
		newNanoseconds += nanosecondsPerDay
	}

	return &Duration{days:newDays, nanosecondsOfDay:newNanoseconds}
}

func (d *Duration) Negate() *Duration {
	if d.nanosecondsOfDay == 0 {
		return &Duration{days:-d.days}
	}

	return &Duration{days:-d.days - 1, nanosecondsOfDay:nanosecondsPerDay - d.nanosecondsOfDay}
}

func (d *Duration) Equals(right *Duration) bool {
	return d.TotalNanoseconds() == right.TotalNanoseconds()
}

func (d *Duration) LessThan(right *Duration) bool {
	return d.TotalNanoseconds() < right.TotalNanoseconds()
}

func (d *Duration) LessThanOrEquals(right *Duration) bool {
	return d.TotalNanoseconds() <= right.TotalNanoseconds()
}

func (d *Duration) GreaterThan(right *Duration) bool {
	return d.TotalNanoseconds() > right.TotalNanoseconds()
}

func (d *Duration) GreaterThanOrEquals(right *Duration) bool {
	return d.TotalNanoseconds() >= right.TotalNanoseconds()
}