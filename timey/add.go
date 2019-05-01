package timey

import "time"

// AddHours adds the given number of hours to the given time
func AddHours(t time.Time, n int) time.Time {
	return t.Add(time.Duration(n) * time.Hour)
}

// AddDays adds the given number of days to the given time
func AddDays(t time.Time, n int) time.Time {
	return t.AddDate(0, 0, n)
}

// AddMonths adds the given number of months to the given time
func AddMonths(t time.Time, n int) time.Time {
	return t.AddDate(0, n, 0)
}
