package uint8y

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestMin(t *testing.T) {
	tests := map[string]struct {
		input    []uint8
		expected uint8
	}{
		"in order":      {input: []uint8{1, 2}, expected: 1},
		"reverse order": {input: []uint8{5, 2}, expected: 2},
		"3 arguments":   {input: []uint8{10, 5, 15}, expected: 5},
		"4 arguments":   {input: []uint8{1, 122, 100}, expected: 1},
		"10 arguments":  {input: []uint8{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := tc.input[0], tc.input[1:len(tc.input)]
			require.Equal(t, tc.expected, Min(v, vs...))
		})
	}
}
