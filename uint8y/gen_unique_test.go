package uint8y_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/slicey"
	"github.com/micvbang/go-helpy/uint8y"
)

// Code generated. DO NOT EDIT.

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []uint8
		expected []uint8
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			got := uint8y.Unique(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func UniqueGenerateTestCase(n int) (test []uint8, expected []uint8) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]uint8, n)
	expected = make([]uint8, 0, n)
	seen := make(map[uint8]struct{})

	for i := 0; i < n; i++ {
		v := uint8y.RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}
