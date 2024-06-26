package helpy_test

import (
	"testing"

	"github.com/micvbang/go-helpy"
)

func TestClamp(t *testing.T) {
	tests := map[string]struct {
		value    int
		min      int
		max      int
		expected int
	}{
		"min < value < max": {value: 42, min: 0, max: 100, expected: 42},
		"min < max < value": {value: 101, min: 0, max: 100, expected: 100},
		"value < min < max": {value: -1, min: 0, max: 100, expected: 0},

		// min > max => min always returned
		"value < max < min": {value: 90, min: 105, max: 100, expected: 105},
		"max < value < min": {value: 101, min: 105, max: 100, expected: 105},
		"max < min < value": {value: 106, min: 105, max: 100, expected: 105},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := helpy.Clamp(test.value, test.min, test.max)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
