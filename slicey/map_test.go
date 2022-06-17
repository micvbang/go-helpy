package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

// TestMap verifies that Map uses f to map values from type T to U.
func TestMap(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int64
	}{
		"nil input": {
			input:    nil,
			expected: []int64{},
		},
		"empty input": {
			input:    []int{},
			expected: []int64{},
		},
		"one value": {
			input:    []int{500},
			expected: []int64{500},
		},
		"multiple values": {
			input:    []int{2, 4, 6, 8, 1, 3, 5, 7},
			expected: []int64{2, 4, 6, 8, 1, 3, 5, 7},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Map(test.input, func(v int) int64 {
				return int64(v)
			})

			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
