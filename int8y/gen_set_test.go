package int8y

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
		"none common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{5, 6, 7, 8}),
			expected: ToSet([]int8{}),
		},
		"one common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{2, 6, 7, 8}),
			expected: ToSet([]int8{2}),
		},
		"multiple common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{2, 1, 7, 4}),
			expected: ToSet([]int8{1, 2, 4}),
		},
		"all common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{4, 3, 2, 1}),
			expected: ToSet([]int8{1, 2, 3, 4}),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(tc.s1)
			s2Len := len(tc.s2)

			require.True(t, tc.s1.Intersect(tc.s2).Equal(tc.expected))

			// Ensure original sets weren't modified
			require.Equal(t, len(tc.s1), s1Len)
			require.Equal(t, len(tc.s2), s2Len)
		})
	}
}

func TestSetUnion(t *testing.T) {
	tests := map[string]struct {
		s1       Set
		s2       Set
		expected Set
	}{
		"none common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{5, 6, 7, 8}),
			expected: ToSet([]int8{1, 2, 3, 4, 5, 6, 7, 8}),
		},
		"one common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{2, 6, 7, 8}),
			expected: ToSet([]int8{1, 2, 3, 4, 6, 7, 8}),
		},
		"multiple common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{2, 1, 7, 4}),
			expected: ToSet([]int8{1, 2, 3, 4, 7}),
		},
		"all common": {
			s1:       ToSet([]int8{1, 2, 3, 4}),
			s2:       ToSet([]int8{4, 3, 2, 1}),
			expected: ToSet([]int8{1, 2, 3, 4}),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(tc.s1)
			s2Len := len(tc.s2)

			require.True(t, tc.s1.Union(tc.s2).Equal(tc.expected))

			// Ensure original sets weren't modified
			require.Equal(t, len(tc.s1), s1Len)
			require.Equal(t, len(tc.s2), s2Len)
		})
	}
}
