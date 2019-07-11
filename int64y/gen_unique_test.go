package int64y

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func UniqueGenerateTestCase(n int) (test []int64, expected []int64) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]int64, n)
	expected = make([]int64, 0, n)
	seen := make(map[int64]struct{})

	for i := 0; i < n; i++ {
		v := RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []int64
		expected []int64
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			require.Equal(t, tc.expected, Unique(tc.input))
		})
	}
}
