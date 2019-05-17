package int32y

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestMin(t *testing.T) {
	tests := map[string]struct {
		input    []int32
		expected int32
	}{
		"in order":      {input: []int32{1, 2}, expected: 1},
		"reverse order": {input: []int32{5, 2}, expected: 2},
		"3 arguments":   {input: []int32{10, 5, 15}, expected: 5},
		"4 arguments":   {input: []int32{1, 122, 100}, expected: 1},
		"10 arguments":  {input: []int32{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 0},

		"one negative":          {input: []int32{-5, 10}, expected: -5},
		"both negative":         {input: []int32{-5, -100}, expected: -100},
		"10 arguments negative": {input: []int32{17, 25, -1, 0, -101, 127, 42, -13, 37, 69}, expected: -101},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := tc.input[0], tc.input[1:len(tc.input)]
			require.Equal(t, tc.expected, Min(v, vs...))
		})
	}
}
