package timey

import "time"

// Max returns the largest value from t0 and ts.
func Max(t0 time.Time, ts ...time.Time) time.Time {
	for _, t := range ts {
		if t.After(t0) {
			t0 = t
		}
	}

	return t0
}

// Max returns the smallest value from t0 and ts.
func Min(t0 time.Time, ts ...time.Time) time.Time {
	for _, t := range ts {
		if t.Before(t0) {
			t0 = t
		}
	}

	return t0
}
