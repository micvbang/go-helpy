package stringy_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/slicey"
	"github.com/micvbang/go-helpy/stringy"
)

// Code generated. DO NOT EDIT.

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []string
		expected []string
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			got := stringy.Unique(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func UniqueGenerateTestCase(n int) (test []string, expected []string) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]string, n)
	expected = make([]string, 0, n)
	seen := make(map[string]struct{})

	for i := 0; i < n; i++ {
		v := stringy.RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}
