package timey_test

import (
	"testing"
	"time"

	"github.com/micvbang/go-helpy/timey"
)

// TestRandomDurationN runs RandomDurationN and ensures that output is within expected range.
func TestRandomDurationN(t *testing.T) {
	durations := []time.Duration{
		1 * time.Nanosecond,
		2 * time.Microsecond,
		3 * time.Millisecond,
		4 * time.Second,
		5 * time.Minute,
		6 * time.Hour,
		24 * time.Hour,
		30 * 24 * time.Hour,
	}
	for _, d := range durations {
		got := timey.RandomDurationN(d)
		if got >= d {
			t.Errorf("expected value < %v, got %v", d, got)
		}

		if got < 0 {
			t.Errorf("expected value > 0, got %v", got)
		}

	}
}
