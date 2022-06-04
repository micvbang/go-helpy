package helpy_test

import (
	"testing"

	"github.com/micvbang/go-helpy"
	"github.com/stretchr/testify/require"
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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := tc.input[0], tc.input[1:len(tc.input)]
			require.Equal(t, tc.expected, helpy.Max(v, vs...))
		})
	}
}
