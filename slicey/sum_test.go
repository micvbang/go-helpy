package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

func TestSum(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"empty": {
			expected: 0,
		},
		"zeroes": {
			input:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: 0,
		},
		"single": {
			input:    []int{42},
			expected: 42,
		},
		"multiple": {
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 45,
		},
		"multiple, negative": {
			input:    []int{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Sum(test.input)

			if got != test.expected {
				t.Errorf("expected sum of %v to be %d, got %d", test.input, test.expected, got)
			}
		})
	}
}
