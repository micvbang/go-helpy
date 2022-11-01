package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		input      []int
		filterFunc func(int) bool
		expected   []int
	}{
		"filter matches all": {
			input: []int{1, 2, 3, 4, 5},
			filterFunc: func(int) bool {
				return true
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		"filter matches none": {
			input: []int{1, 2, 3, 4, 5},
			filterFunc: func(int) bool {
				return false
			},
			expected: []int{},
		},
		"filter matches some": {
			input: []int{1, 2, 3, 4, 5},
			filterFunc: func(v int) bool {
				return v == 2 || v == 4
			},
			expected: []int{2, 4},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Filter(test.input, test.filterFunc)

			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected to find %v, found %v", test.expected, got)
			}
		})
	}
}
