package helpy_test

import (
	"testing"

	"github.com/micvbang/go-helpy"
)

func TestSetIntersect(t *testing.T) {
	tests := map[string]struct {
		s1       helpy.Set[int]
		s2       helpy.Set[int]
		expected helpy.Set[int]
	}{
		"none common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{5, 6, 7, 8}),
			expected: helpy.ToSet([]int{}),
		},
		"one common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{2, 6, 7, 8}),
			expected: helpy.ToSet([]int{2}),
		},
		"multiple common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{2, 1, 7, 4}),
			expected: helpy.ToSet([]int{1, 2, 4}),
		},
		"all common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{4, 3, 2, 1}),
			expected: helpy.ToSet([]int{1, 2, 3, 4}),
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
		s1       helpy.Set[int]
		s2       helpy.Set[int]
		expected helpy.Set[int]
	}{
		"none common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{5, 6, 7, 8}),
			expected: helpy.ToSet([]int{1, 2, 3, 4, 5, 6, 7, 8}),
		},
		"one common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{2, 6, 7, 8}),
			expected: helpy.ToSet([]int{1, 2, 3, 4, 6, 7, 8}),
		},
		"multiple common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{2, 1, 7, 4}),
			expected: helpy.ToSet([]int{1, 2, 3, 4, 7}),
		},
		"all common": {
			s1:       helpy.ToSet([]int{1, 2, 3, 4}),
			s2:       helpy.ToSet([]int{4, 3, 2, 1}),
			expected: helpy.ToSet([]int{1, 2, 3, 4}),
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

func TestSetContains(t *testing.T) {
	tests := map[string]struct {
		s        helpy.Set[int]
		v        int
		expected bool
	}{
		"nil set": {
			s:        helpy.ToSet[int](nil),
			v:        1,
			expected: false,
		},
		"empty set": {
			s:        helpy.ToSet([]int{}),
			v:        1,
			expected: false,
		},
		"one element contains": {
			s:        helpy.ToSet([]int{1}),
			v:        1,
			expected: true,
		},
		"multiple elements contains first": {
			s:        helpy.ToSet([]int{1, 2, 3, 4, 5, 6}),
			v:        1,
			expected: true,
		},
		"multiple elements contains middle": {
			s:        helpy.ToSet([]int{1, 2, 3, 4, 5, 6}),
			v:        3,
			expected: true,
		},
		"multiple elements contains last": {
			s:        helpy.ToSet([]int{1, 2, 3, 4, 5, 6}),
			v:        6,
			expected: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.s.Contains(test.v)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestSetEqual(t *testing.T) {
	tests := map[string]struct {
		s1       helpy.Set[int]
		s2       helpy.Set[int]
		expected bool
	}{
		"nil empty": {
			s1:       nil,
			s2:       helpy.Set[int]{},
			expected: true,
		},
		"empty nil": {
			s1:       helpy.Set[int]{},
			s2:       nil,
			expected: true,
		},
		"nil something": {
			s1:       nil,
			s2:       helpy.MakeSet(1, 2, 3),
			expected: false,
		},
		"something nil": {
			s1:       helpy.MakeSet(1, 2, 3),
			s2:       nil,
			expected: false,
		},
		"empty something": {
			s1:       helpy.Set[int]{},
			s2:       helpy.MakeSet(1, 2, 3),
			expected: false,
		},
		"something empty": {
			s1:       helpy.MakeSet(1, 2, 3),
			s2:       helpy.Set[int]{},
			expected: false,
		},
		"none common": {
			s1:       helpy.MakeSet(1, 2, 3, 4),
			s2:       helpy.MakeSet(5, 6, 7, 8),
			expected: false,
		},
		"one common": {
			s1:       helpy.MakeSet(1, 2, 3, 4),
			s2:       helpy.MakeSet(2, 6, 7, 8),
			expected: false,
		},
		"multiple common": {
			s1:       helpy.MakeSet(1, 2, 3, 4),
			s2:       helpy.MakeSet(2, 1, 7, 4),
			expected: false,
		},
		"all common in order": {
			s1:       helpy.MakeSet(1, 2, 3, 4),
			s2:       helpy.MakeSet(1, 2, 3, 4),
			expected: true,
		},
		"all common out of order": {
			s1:       helpy.MakeSet(1, 2, 3, 4),
			s2:       helpy.MakeSet(4, 3, 2, 1),
			expected: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s1Len := len(test.s1)
			s2Len := len(test.s2)

			got := test.s1.Equal(test.s2)
			if got != test.expected {
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

func TestSetAdd(t *testing.T) {
	tests := map[string]struct {
		s        helpy.Set[int]
		v        int
		expected helpy.Set[int]
	}{
		// TODO: how can this case be handled?
		// "nil set": {
		// 	s:        nil,
		// 	v:        1,
		// 	expected: helpy.MakeSet(1),
		// },
		"empty set": {
			s:        helpy.ToSet([]int{}),
			v:        1,
			expected: helpy.MakeSet(1),
		},
		"one element already exists": {
			s:        helpy.ToSet([]int{1}),
			v:        1,
			expected: helpy.MakeSet(1),
		},
		"one element doesn't exist": {
			s:        helpy.ToSet([]int{1}),
			v:        2,
			expected: helpy.MakeSet(1, 2),
		},
		"multiple elements already exists": {
			s:        helpy.ToSet([]int{1, 2, 3, 4, 5, 6}),
			v:        3,
			expected: helpy.MakeSet(1, 2, 3, 4, 5, 6),
		},
		"multiple elements doesn't exists": {
			s:        helpy.ToSet([]int{1, 2, 3, 4, 5, 6}),
			v:        100,
			expected: helpy.MakeSet(1, 2, 3, 4, 5, 6, 100),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.s.Add(test.v)

			if len(test.s) != len(test.expected) {
				t.Errorf("expected set to grow to size %v, got %v", len(test.expected), len(test.s))
			}

			if !test.s.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, test.s)
			}
		})
	}
}

func TestSetSubtract(t *testing.T) {
	tests := map[string]struct {
		s1       helpy.Set[int]
		s2       helpy.Set[int]
		expected helpy.Set[int]
	}{
		"empty s2": {
			s1:       helpy.MakeSet(1, 2, 3),
			s2:       helpy.MakeSet[int](),
			expected: helpy.MakeSet(1, 2, 3),
		},
		"empty s1": {
			s1:       helpy.MakeSet[int](),
			s2:       helpy.MakeSet(1, 2, 3),
			expected: helpy.MakeSet[int](),
		},
		"s1 == s2": {
			s1:       helpy.MakeSet(1, 2, 3),
			s2:       helpy.MakeSet(1, 2, 3),
			expected: helpy.MakeSet[int](),
		},
		"removes one": {
			s1:       helpy.MakeSet(1, 2, 3),
			s2:       helpy.MakeSet(2),
			expected: helpy.MakeSet(1, 3),
		},
		"removes multiple": {
			s1:       helpy.MakeSet(1, 2, 3),
			s2:       helpy.MakeSet(2, 3),
			expected: helpy.MakeSet(1),
		},
		"removes all": {
			s1:       helpy.MakeSet(1, 2, 3),
			s2:       helpy.MakeSet(1, 2, 3),
			expected: helpy.MakeSet[int](),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.s1.Subtract(test.s2)
			if !got.Equal(test.expected) {
				t.Errorf("expecetd %v, got %v", test.expected, got)
			}
		})
	}
}
