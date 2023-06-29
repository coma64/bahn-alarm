package time_conversion

import "time"

// ToToday Creates a new time.Time where year, month and day are set to time.Now
// The time.Time needs to be first converted ToToday and then into the time.zone needed
func ToToday(t time.Time) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// TimeOnly Strips the year, month and day from a time.Time
func TimeOnly(t time.Time) time.Time {
	year, month, day := t.Date()
	return t.AddDate(-year, -int(month)+1, -day+1)
}
