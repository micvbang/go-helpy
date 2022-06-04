package int32y_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/int32y"
	"github.com/micvbang/go-helpy/slicey"
)

// Code generated. DO NOT EDIT.

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []int32
		expected []int32
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			got := int32y.Unique(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func UniqueGenerateTestCase(n int) (test []int32, expected []int32) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]int32, n)
	expected = make([]int32, 0, n)
	seen := make(map[int32]struct{})

	for i := 0; i < n; i++ {
		v := int32y.RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}
