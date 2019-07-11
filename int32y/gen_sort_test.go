package int32y

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestSort(t *testing.T) {
	tests := map[string]struct {
		input    []int32
		expected []int32
	}{
		"in order": {
			input:    []int32{1, 2, 3, 4, 5},
			expected: []int32{1, 2, 3, 4, 5},
		},
		"reverse order": {
			input:    []int32{5, 4, 3, 2, 1},
			expected: []int32{1, 2, 3, 4, 5},
		},
		"random order": {
			input:    []int32{4, 1, 2, 5, 3},
			expected: []int32{1, 2, 3, 4, 5},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, Sort(tc.input))
		})
	}
}
