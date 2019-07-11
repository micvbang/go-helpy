package uint8y

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestSort(t *testing.T) {
	tests := map[string]struct {
		input    []uint8
		expected []uint8
	}{
		"in order": {
			input:    []uint8{1, 2, 3, 4, 5},
			expected: []uint8{1, 2, 3, 4, 5},
		},
		"reverse order": {
			input:    []uint8{5, 4, 3, 2, 1},
			expected: []uint8{1, 2, 3, 4, 5},
		},
		"random order": {
			input:    []uint8{4, 1, 2, 5, 3},
			expected: []uint8{1, 2, 3, 4, 5},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, Sort(tc.input))
		})
	}
}
