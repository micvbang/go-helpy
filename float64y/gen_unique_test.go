package float64y_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/float64y"
	"github.com/micvbang/go-helpy/slicey"
)

// Code generated. DO NOT EDIT.

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []float64
		expected []float64
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			got := float64y.Unique(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func UniqueGenerateTestCase(n int) (test []float64, expected []float64) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]float64, n)
	expected = make([]float64, 0, n)
	seen := make(map[float64]struct{})

	for i := 0; i < n; i++ {
		v := float64y.RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}
