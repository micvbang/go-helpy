package timey

import "time"

// DurationAbs returns the absolute (positive) value of a time.Duration.
func DurationAbs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
