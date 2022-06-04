package inty

import (
	"testing"
)

// Code generated. DO NOT EDIT.

func TestSetIntersect(t *testing.T) {
	tests := map[string]struct {
		s1       Set
		s2       Set
		expected Set
	}{
		"none common": {
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{5, 6, 7, 8}),
			expected: ToSet([]int{}),
		},
		"one common": {
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{2, 6, 7, 8}),
			expected: ToSet([]int{2}),
		},
		"multiple common": {
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{2, 1, 7, 4}),
			expected: ToSet([]int{1, 2, 4}),
		},
		"all common": {
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{4, 3, 2, 1}),
			expected: ToSet([]int{1, 2, 3, 4}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(test.s1)
			s2Len := len(test.s2)

			got := test.s1.Intersect(test.s2)
			if !got.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}

			// Ensure original sets weren't modified
			if len(test.s1) != s1Len {
				t.Errorf("expected set to be of size %v, got %v", s1Len, test.s1)
			}

			if len(test.s2) != s2Len {
				t.Errorf("expected set to be of size %v, got %v", s2Len, test.s2)
			}
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
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{5, 6, 7, 8}),
			expected: ToSet([]int{1, 2, 3, 4, 5, 6, 7, 8}),
		},
		"one common": {
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{2, 6, 7, 8}),
			expected: ToSet([]int{1, 2, 3, 4, 6, 7, 8}),
		},
		"multiple common": {
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{2, 1, 7, 4}),
			expected: ToSet([]int{1, 2, 3, 4, 7}),
		},
		"all common": {
			s1:       ToSet([]int{1, 2, 3, 4}),
			s2:       ToSet([]int{4, 3, 2, 1}),
			expected: ToSet([]int{1, 2, 3, 4}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(test.s1)
			s2Len := len(test.s2)

			got := test.s1.Union(test.s2)
			if !got.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}

			// Ensure original sets weren't modified
			if len(test.s1) != s1Len {
				t.Errorf("expected set to be of size %v, got %v", s1Len, test.s1)
			}

			if len(test.s2) != s2Len {
				t.Errorf("expected set to be of size %v, got %v", s2Len, test.s2)
			}
		})
	}
}
