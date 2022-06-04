package float32y_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

// TestContains verifies that Contains returns true only when f returns true.
func TestContains(t *testing.T) {
	tests := map[string]struct {
		haystack []int
		needle   int
		expected bool
	}{
		"nil haystack": {
			haystack: nil,
			expected: false,
		},
		"empty haystack": {
			haystack: []int{},
			expected: false,
		},
		"one element exists": {
			haystack: []int{42},
			needle:   42,
			expected: true,
		},
		"one element, doesn't exist": {
			haystack: []int{1337},
			needle:   42,
			expected: false,
		},
		"multiple elements, exists first": {
			haystack: []int{42, 1337, -1500, 17, 71, 18821390293},
			needle:   42,
			expected: true,
		},
		"multiple elements, exists middle": {
			haystack: []int{1337, -1500, 42, 17, 71, 18821390293},
			needle:   42,
			expected: true,
		},
		"multiple elements, exists last": {
			haystack: []int{1337, -1500, 17, 71, 18821390293, 42},
			needle:   42,
			expected: true,
		},
		"multiple elements, doesn't exist": {
			haystack: []int{-42, 1337, -1500, 17, 71, 18821390293},
			needle:   42,
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Contains(test.haystack, test.needle)

			if got != test.expected {
				t.Errorf("expected exist %v, got %v", test.expected, got)
			}
		})
	}
}
