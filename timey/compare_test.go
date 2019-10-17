package timey

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTruncCompare(t *testing.T) {
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
			require.Equal(t, test.expected, got)
		})
	}
}
