package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		vs1      []int
		vs2      []int
		expected bool
	}{
		"nils": {
			vs1:      nil,
			vs2:      nil,
			expected: true,
		},
		"empty": {
			vs1:      []int{},
			vs2:      []int{},
			expected: true,
		},
		"not same length": {
			vs1:      []int{1},
			vs2:      []int{1, 2},
			expected: false,
		},
		"equal, one element": {
			vs1:      []int{1},
			vs2:      []int{1},
			expected: true,
		},
		"equal, multiple elements": {
			vs1:      []int{3, 2, 1},
			vs2:      []int{3, 2, 1},
			expected: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Equal(test.vs1, test.vs2)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestEqualFunc(t *testing.T) {
	tests := map[string]struct {
		vs1      []int
		vs2      []int
		expected bool
	}{
		"nils": {
			vs1:      nil,
			vs2:      nil,
			expected: true,
		},
		"empty": {
			vs1:      []int{},
			vs2:      []int{},
			expected: true,
		},
		"not same length": {
			vs1:      []int{1},
			vs2:      []int{1, 2},
			expected: false,
		},
		"equal, one element": {
			vs1:      []int{1},
			vs2:      []int{1},
			expected: true,
		},
		"equal, multiple elements": {
			vs1:      []int{3, 2, 1},
			vs2:      []int{3, 2, 1},
			expected: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.EqualFunc(test.vs1, test.vs2, func(v1 int, v2 int) bool {
				return v1 == v2
			})
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
