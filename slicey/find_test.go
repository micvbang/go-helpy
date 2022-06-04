package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/inty"
	"github.com/micvbang/go-helpy/slicey"
)

// TestFindSimpleValue verifies that Find finds the expected values in a slice
// of simple values.
func TestFindSimpleValue(t *testing.T) {
	tests := map[string]struct {
		haystack       []int
		needle         int
		expectedExists bool
	}{
		"nil haystack": {
			haystack:       nil,
			expectedExists: false,
		},
		"empty haystack": {
			haystack:       []int{},
			expectedExists: false,
		},
		"one element exists": {
			haystack:       []int{42},
			needle:         42,
			expectedExists: true,
		},
		"one element, doesn't exist": {
			haystack:       []int{1337},
			needle:         42,
			expectedExists: false,
		},
		"multiple elements, exists first": {
			haystack:       []int{42, 1337, -1500, 17, 71, 18821390293},
			needle:         42,
			expectedExists: true,
		},
		"multiple elements, exists middle": {
			haystack:       []int{1337, -1500, 42, 17, 71, 18821390293},
			needle:         42,
			expectedExists: true,
		},
		"multiple elements, exists last": {
			haystack:       []int{1337, -1500, 17, 71, 18821390293, 42},
			needle:         42,
			expectedExists: true,
		},
		"multiple elements, doesn't exist": {
			haystack:       []int{-42, 1337, -1500, 17, 71, 18821390293},
			needle:         42,
			expectedExists: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotExists := slicey.Find(test.haystack, func(v int) bool {
				return test.needle == v
			})
			if gotExists != test.expectedExists {
				t.Errorf("expected exist %v, got %v", test.expectedExists, gotExists)
			}

			if test.expectedExists && got != test.needle {
				t.Errorf("expected to find %v, found %v", test.needle, got)
			}
		})
	}
}

// TestFindStruct verifies that Find finds the expected values in a slice of
// complex types.
func TestFindStruct(t *testing.T) {
	type complex struct {
		v       string
		sidecar int
	}

	n := func(v string) complex {
		return complex{
			v:       v,
			sidecar: inty.Random(),
		}
	}

	tests := map[string]struct {
		haystack       []complex
		needle         complex
		expectedExists bool
	}{
		"nil haystack": {
			haystack:       nil,
			expectedExists: false,
		},
		"empty haystack": {
			haystack:       []complex{},
			expectedExists: false,
		},
		"one element exists": {
			haystack:       []complex{n("yes")},
			needle:         n("yes"),
			expectedExists: true,
		},
		"one element, doesn't exist": {
			haystack:       []complex{n("1337")},
			needle:         n("42"),
			expectedExists: false,
		},
		"multiple elements, exists": {
			haystack:       []complex{n("me"), n("nah"), n("me")},
			needle:         n("me"),
			expectedExists: true,
		},
		"multiple elements, exists first": {
			haystack:       []complex{n("me"), n("nah"), n("sup")},
			needle:         n("me"),
			expectedExists: true,
		},
		"multiple elements, exists middle": {
			haystack:       []complex{n("1337"), n("-1500"), n("me"), n("17"), n("71"), n("18821390293")},
			needle:         n("me"),
			expectedExists: true,
		},
		"multiple elements, exists last": {
			haystack:       []complex{n("1337"), n("-1500"), n("17"), n("71"), n("18821390293"), n("me")},
			needle:         n("me"),
			expectedExists: true,
		},
		"multiple elements, doesn't exist": {
			haystack:       []complex{n("-42"), n("1337"), n("-1500"), n("17"), n("71"), n("18821390293")},
			needle:         n("42"),
			expectedExists: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotExists := slicey.Find(test.haystack, func(v complex) bool {
				return test.needle.v == v.v
			})
			if gotExists != test.expectedExists {
				t.Errorf("expected exist %v, got %v", test.expectedExists, gotExists)
			}

			if test.expectedExists && got.v != test.needle.v {
				t.Errorf("expected to find %v, found %v", test.needle, got)
			}
		})
	}
}

// TestFindFindsFirst verifies that Find stops the first time f returns true.
func TestFindFindsFirst(t *testing.T) {
	calls := 0
	got, gotExists := slicey.Find([]int{1, 2, 2, 2, 2, 1}, func(v int) bool {
		calls += 1
		return v == 1
	})
	if gotExists != true {
		t.Errorf("expected exists true, got %v", gotExists)
	}

	if got != 1 {
		t.Errorf("expected to find 1, got %v", got)
	}

	if calls != 1 {
		t.Errorf("expected 1 calls, got %v", calls)
	}
}
