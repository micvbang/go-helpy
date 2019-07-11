package int16y

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestSetIntersect(t *testing.T) {
	tests := map[string]struct {
		s1       Set
		s2       Set
		expected Set
	}{
		"empty intersection": {
			s1:       ToSet([]int16{1, 2, 3, 4}),
			s2:       ToSet([]int16{5, 6, 7, 8}),
			expected: ToSet([]int16{}),
		},
		"one common": {
			s1:       ToSet([]int16{1, 2, 3, 4}),
			s2:       ToSet([]int16{2, 6, 7, 8}),
			expected: ToSet([]int16{2}),
		},
		"multiple common": {
			s1:       ToSet([]int16{1, 2, 3, 4}),
			s2:       ToSet([]int16{2, 1, 7, 4}),
			expected: ToSet([]int16{1, 2, 4}),
		},
		"all common": {
			s1:       ToSet([]int16{1, 2, 3, 4}),
			s2:       ToSet([]int16{4, 3, 2, 1}),
			expected: ToSet([]int16{1, 2, 3, 4}),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.True(t, tc.s1.Intersect(tc.s2).Equal(tc.expected))
		})
	}
}
