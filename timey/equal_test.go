package timey

import (
	"testing"
	"time"
)

func TestTruncEqual(t *testing.T) {
	now := time.Now().Truncate(time.Hour)

	tests := map[string]struct {
		t0       time.Time
		t1       time.Time
		d        time.Duration
		expected bool
	}{
		"equal": {
			t0:       now,
			t1:       now,
			d:        time.Second,
			expected: true,
		},
		"add second, trunc second": {
			t0:       now,
			t1:       now.Add(time.Second),
			d:        time.Second,
			expected: false,
		},
		"add second, trunc minute": {
			t0:       now,
			t1:       now.Add(time.Second),
			d:        time.Minute,
			expected: true,
		},
		"add day, trunc minute": {
			t0:       now,
			t1:       AddDays(now, 1),
			d:        time.Minute,
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := TruncEqual(test.d, test.t0, test.t1)
			if test.expected != got {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestDiffEqual(t *testing.T) {
	now := time.Now().Truncate(time.Hour)

	tests := map[string]struct {
		t0       time.Time
		t1       time.Time
		d        time.Duration
		expected bool
	}{
		"equal": {
			t0:       now,
			t1:       now,
			d:        time.Second,
			expected: true,
		},
		"add 1 second, diff 1 second": {
			t0:       now,
			t1:       now.Add(time.Second),
			d:        time.Second,
			expected: true,
		},
		"add 1 second, diff 1 second, rev": {
			t1:       now.Add(time.Second),
			t0:       now,
			d:        time.Second,
			expected: true,
		},
		"sub 1 second, diff 1 second": {
			t0:       now,
			t1:       now.Add(-time.Second),
			d:        time.Second,
			expected: true,
		},
		"add 2 seconds, diff 1 second": {
			t0:       now,
			t1:       now.Add(2 * time.Second),
			d:        time.Second,
			expected: false,
		},
		"add 1 hour, diff 1 minute": {
			t0:       now,
			t1:       now.Add(time.Hour),
			d:        time.Minute,
			expected: false,
		},
		"add 1 minute, diff 1 hour": {
			t0:       now,
			t1:       now.Add(time.Hour),
			d:        time.Minute,
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := DiffEqual(test.d, test.t0, test.t1)
			if test.expected != got {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
