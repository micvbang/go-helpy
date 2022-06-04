package helpy_test

import (
	"testing"

	"github.com/micvbang/go-helpy"
)

// TestMax verifies that Max returns the largest value from the given input.
func TestMax(t *testing.T) {
	tests := map[string]struct {
		input    []int32
		expected int32
	}{
		"in order":      {input: []int32{1, 2}, expected: 2},
		"reverse order": {input: []int32{5, 2}, expected: 5},
		"3 arguments":   {input: []int32{10, 5, 15}, expected: 15},
		"4 arguments":   {input: []int32{1, 122, 100}, expected: 122},
		"10 arguments":  {input: []int32{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 125},

		"one negative":          {input: []int32{-5, 10}, expected: 10},
		"both negative":         {input: []int32{-5, -100}, expected: -5},
		"10 arguments negative": {input: []int32{17, 25, -1, 0, -101, 127, 42, -13, 37, 69}, expected: 127},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := helpy.Max(v, vs...)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
