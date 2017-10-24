package godatime

type timePeriodField struct {
	unitNanoseconds int64
}

var (
	timePeriodNanoseconds = timePeriodField{unitNanoseconds:1}

	timePeriodTicks = timePeriodField{unitNanoseconds:nanosecondsPerTick}

	timePeriodMilliseconds = timePeriodField{unitNanoseconds:nanosecondsPerMillisecond}

	timePeriodSeconds = timePeriodField{unitNanoseconds:nanosecondsPerSecond}

	timePeriodMinutes = timePeriodField{unitNanoseconds:nanosecondsPerMinute}

	timePeriodHours = timePeriodField{unitNanoseconds:nanosecondsPerHour}
)

func (t *timePeriodField) unitsPerDay() int64 {
	return nanosecondsPerDay / t.unitNanoseconds
}

func (t *timePeriodField) AddValue(nanosecondsOfDay, value int64) *LocalTime{
	if value >= 0 {
		if value >= t.unitsPerDay() {
			value = value % t.unitsPerDay()
		}

		nanosecondsToAdd := value * t.unitNanoseconds
		newNanoseconds := nanosecondsOfDay + nanosecondsToAdd
		if newNanoseconds >= nanosecondsPerDay {
			newNanoseconds -= nanosecondsPerDay
		}

		return &LocalTime{nanoseconds: newNanoseconds}
	}

	if value <= -t.unitsPerDay() {
		value = value % t.unitsPerDay()
	}

	nanosecondsToAdd := value * t.unitNanoseconds
	newNanoseconds := nanosecondsOfDay + nanosecondsToAdd
	if newNanoseconds < 0 {
		newNanoseconds += nanosecondsPerDay
	}

	return &LocalTime{nanoseconds: newNanoseconds}
}