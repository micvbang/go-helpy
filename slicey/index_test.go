package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

func TestLast(t *testing.T) {
	tests := map[string]struct {
		vs       []int
		expected int
	}{
		"nil":      {vs: nil, expected: 0},
		"empty":    {vs: []int{}, expected: 0},
		"one":      {vs: []int{42}, expected: 42},
		"multiple": {vs: []int{1, 2, 3}, expected: 3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Last(test.vs)
			if got != test.expected {
				t.Errorf("expected %d, got %d", test.expected, got)
			}
		})
	}
}

func TestLastOrDefault(t *testing.T) {
	tests := map[string]struct {
		vs       []int
		def      int
		expected int
	}{
		"nil":      {vs: nil, def: 42, expected: 42},
		"empty":    {vs: []int{}, def: 42, expected: 42},
		"one":      {vs: []int{1337}, def: 42, expected: 1337},
		"multiple": {vs: []int{1, 2, 3}, def: 42, expected: 3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.LastOrDefault(test.vs, test.def)
			if got != test.expected {
				t.Errorf("expected %d, got %d", test.expected, got)
			}
		})
	}
}
