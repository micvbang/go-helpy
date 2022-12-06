package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

func TestAll(t *testing.T) {
	tests := map[string]struct {
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		"empty input": {
			input:     []int{},
			predicate: func(v int) bool { return false }, // NOTE: function output isn't used
			expected:  true,
		},
		"nil input": {
			input:     nil,
			predicate: func(v int) bool { return false }, // NOTE: function output isn't used
			expected:  true,
		},
		"one matches": {
			input:     []int{1},
			predicate: func(v int) bool { return v == 1 },
			expected:  true,
		},
		"multiple matches": {
			input:     []int{1, 2, 3, 4, 5},
			predicate: func(v int) bool { return v > 0 },
			expected:  true,
		},
		"one does not match": {
			input:     []int{1},
			predicate: func(v int) bool { return v == 2 },
			expected:  false,
		},
		"no matches": {
			input:     []int{1, 2, 3, 4, 5},
			predicate: func(v int) bool { return v > 100 },
			expected:  false,
		},
		"some match, some do not match": {
			input:     []int{1, 2, 3, 4, 5},
			predicate: func(v int) bool { return v == 4 || v == 2 },
			expected:  false,
		},
		"predicate is nil": {
			input:     []int{1, 2, 3, 4, 5},
			predicate: nil, // NOTE: no predicate given
			expected:  true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.All(test.input, test.predicate)

			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
