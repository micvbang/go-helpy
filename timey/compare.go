package timey

import "time"

// TruncEqual truncates time.Time inputs down to a multiple of d (since the
// zero time) and returns true of all time.Time inputs are equal
func TruncEqual(d time.Duration, t0 time.Time, ts ...time.Time) bool {
	t0 = t0.Truncate(d)

	equal := true
	for _, t1 := range ts {
		t1 = t1.Truncate(d)
		equal = equal && t0.Equal(t1)
	}

	return equal
}
