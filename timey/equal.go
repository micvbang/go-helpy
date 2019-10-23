package timey

import (
	"time"
)

// TruncEqual truncates t0 and ts inputs down to a multiple of d (since the
// zero time) and returns true if all ts are equal to t0.
func TruncEqual(d time.Duration, t0 time.Time, ts ...time.Time) bool {
	t0 = t0.Truncate(d)

	equal := true
	for _, t1 := range ts {
		t1 = t1.Truncate(d)
		equal = equal && t0.Equal(t1)
	}

	return equal
}

// DiffEqual returns true if all ts inputs are within duration d range of t0.
func DiffEqual(d time.Duration, t0 time.Time, ts ...time.Time) bool {
	dAbs := DurationAbs(d)

	equal := true
	for _, t1 := range ts {
		equal = equal && DurationAbs(t0.Sub(t1)) <= dAbs
	}

	return equal
}
