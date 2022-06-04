package timey

import (
	"testing"
	"time"
)

func TestDurationAbs(t *testing.T) {
	tests := map[string]struct {
		d        time.Duration
		expected time.Duration
	}{
		"500":      {d: time.Duration(500), expected: time.Duration(500)},
		"0":        {d: time.Duration(0), expected: time.Duration(0)},
		"-5":       {d: time.Duration(-5), expected: time.Duration(5)},
		"5 hours":  {d: 5 * time.Hour, expected: 5 * time.Hour},
		"-5 hours": {d: -5 * time.Hour, expected: 5 * time.Hour},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := DurationAbs(test.d)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
