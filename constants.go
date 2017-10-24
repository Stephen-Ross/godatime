package godatime

const (
	nanosecondsPerTick int64 = 100

	nanosecondsPerMillisecond int64 = 1000000

	nanosecondsPerSecond int64 = 1000000000

	nanosecondsPerMinute int64 = nanosecondsPerSecond * secondsPerMinute

	nanosecondsPerHour int64 = nanosecondsPerMinute * minutesPerHour

	nanosecondsPerDay int64 = nanosecondsPerHour * hoursPerDay

	nanosecondsPerWeek int64 = nanosecondsPerDay * daysPerWeek

	ticksPerMillisecond int64 = 10000

	ticksPerSecond int64 = ticksPerMillisecond * millisecondsPerSecond

	ticksPerMinute int64 = ticksPerSecond * secondsPerMinute

	ticksPerHour int64 = ticksPerMinute * minutesPerHour

	ticksPerDay int64 = ticksPerHour * hoursPerDay

	ticksPerWeek int64 = ticksPerDay * daysPerWeek

	millisecondsPerSecond = 1000

	millisecondsPerMinute = millisecondsPerSecond * secondsPerMinute

	millisecondsPerHour = millisecondsPerMinute * minutesPerHour

	millisecondsPerDay = millisecondsPerHour * hoursPerDay

	millisecondsPerWeek = millisecondsPerDay * daysPerWeek

	secondsPerMinute = 60

	secondsPerHour = secondsPerMinute * minutesPerHour

	secondsPerDay = secondsPerHour * hoursPerDay

	secondsPerWeek = secondsPerDay * daysPerWeek

	minutesPerHour = 60

	minutesPerDay = minutesPerHour * hoursPerDay

	minutesPerWeek = minutesPerDay * daysPerWeek

	hoursPerDay = 24

	hoursPerWeek = hoursPerDay * daysPerWeek

	daysPerWeek = 7
)