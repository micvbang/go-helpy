package uint16y_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/slicey"
	"github.com/micvbang/go-helpy/uint16y"
)

// Code generated. DO NOT EDIT.

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []uint16
		expected []uint16
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			got := uint16y.Unique(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func UniqueGenerateTestCase(n int) (test []uint16, expected []uint16) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]uint16, n)
	expected = make([]uint16, 0, n)
	seen := make(map[uint16]struct{})

	for i := 0; i < n; i++ {
		v := uint16y.RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}
