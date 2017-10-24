package godatime

import (
	"fmt"
	"./extensions"
)

type Period struct {
	nanoseconds int64
	ticks int64
	milliseconds int64
	seconds int64
	minutes int64
	hours int64
	days int
	weeks int
	months int
	years int
}

func NewPeriod(builder *PeriodBuilder) *Period {
	return &Period{
		nanoseconds: builder.Nanoseconds,
		ticks: builder.Ticks,
		milliseconds: builder.Milliseconds,
		seconds: builder.Seconds,
		minutes: builder.Minutes,
		hours: builder.Hours,
		days: builder.Days,
		weeks: builder.Weeks,
		months: builder.Months,
		years: builder.Years}
}

func Between(start, end *LocalTime) *Period {
	remaining := end.NanosecondOfDay() - start.NanosecondOfDay()

	unitsBetween := func(nanosecondsPerUnit int64) (result int64) {
		quotient, remainder := extensions.DivMod(remaining, nanosecondsPerUnit)
		remaining = remainder
		return quotient
	}

	hours := unitsBetween(nanosecondsPerHour)
	minutes := unitsBetween(nanosecondsPerMinute)
	seconds := unitsBetween(nanosecondsPerSecond)
	milliseconds := unitsBetween(nanosecondsPerMillisecond)
	ticks := unitsBetween(nanosecondsPerTick)
	nanoseconds := unitsBetween(1)

	return &Period{nanoseconds:nanoseconds,
		ticks:ticks,
		milliseconds:milliseconds,
		seconds:seconds,
		minutes:minutes,
		hours:hours}
}

func (p *Period) Nanoseconds() int64 {
	return p.nanoseconds
}

func (p *Period) Ticks() int64 {
	return p.ticks
}

func (p *Period) Milliseconds() int64 {
	return p.milliseconds
}

func (p *Period) Seconds() int64 {
	return p.seconds
}

func (p *Period) Minutes() int64 {
	return p.minutes
}

func (p *Period) Hours() int64 {
	return p.hours
}

func (p *Period) Days() int {
	return p.days
}

func (p *Period) Weeks() int {
	return p.weeks
}

func (p *Period) Months() int {
	return p.months
}

func (p *Period) Years() int {
	return p.years
}

func (p *Period) Add(right *Period) *Period {
	return &Period{years:p.years + right.years,
		months:p.months + right.months,
		weeks:p.weeks + right.weeks,
		days:p.days + right.days,
		hours:p.hours + right.hours,
		minutes:p.minutes + right.minutes,
		seconds:p.seconds + right.seconds,
		milliseconds:p.milliseconds + right.milliseconds,
		ticks:p.ticks + right.ticks,
		nanoseconds:p.nanoseconds + right.nanoseconds}
}

func (p *Period) Subtract(right *Period) *Period {
	return &Period{years:p.years - right.years,
		months:p.months - right.months,
		weeks:p.weeks - right.weeks,
		days:p.days - right.days,
		hours:p.hours - right.hours,
		minutes:p.minutes - right.minutes,
		seconds:p.seconds - right.seconds,
		milliseconds:p.milliseconds - right.milliseconds,
		ticks:p.ticks - right.ticks,
		nanoseconds:p.nanoseconds - right.nanoseconds}
}

func (p *Period) HasTimeComponent() bool {
	return p.hours != 0 || p.minutes != 0 || p.seconds != 0 || p.milliseconds != 0 || p.ticks != 0 || p.nanoseconds != 0
}

func (p *Period) HasDateComponent() bool {
	return p.years != 0 || p.months != 0 || p.weeks != 0 || p.days != 0
}

func (p *Period) Normalize() *Period {
	totalNanoseconds := p.totalNanoseconds()

	return &Period{years:p.years,
		months:p.months,
		days:(int)(totalNanoseconds / nanosecondsPerDay),
		hours:(totalNanoseconds / nanosecondsPerHour) % hoursPerDay,
		minutes:(totalNanoseconds / nanosecondsPerMinute) % minutesPerHour,
		seconds:(totalNanoseconds / nanosecondsPerSecond) % secondsPerMinute,
		milliseconds:(totalNanoseconds / nanosecondsPerMillisecond) % millisecondsPerSecond,
		nanoseconds:totalNanoseconds % nanosecondsPerMillisecond}
}

func (p *Period) Equals(other *Period) bool {
	return other != nil && p.years == other.years &&
		p.months == other.months && p.weeks == other.weeks &&
			p.days == other.days && p.hours == other.hours &&
				p.minutes == other.minutes && p.seconds == other.seconds &&
					p.milliseconds == other.milliseconds && p.ticks == other.ticks &&
						p.nanoseconds == other.nanoseconds
}

func (p *Period) String() string {
	return fmt.Sprintf("%v.%v.%v.%v.%v:%v:%v.%v", p.years, p.months, p.weeks, p.days, p.hours, p.minutes, p.seconds, p.milliseconds)
}

func (p *Period) totalNanoseconds() int64 {
	return (int64)(p.nanoseconds + (int64)(p.ticks *nanosecondsPerTick) +
		(int64)(p.milliseconds *nanosecondsPerMillisecond) +
		(int64)(p.seconds *nanosecondsPerSecond) +
		(int64)(p.minutes *nanosecondsPerMinute) +
		(int64)(p.hours *nanosecondsPerHour) +
		(int64)((int64)(p.days) *nanosecondsPerDay) +
		(int64)((int64)(p.weeks) *nanosecondsPerWeek))
}