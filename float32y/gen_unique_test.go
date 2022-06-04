package float32y_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/float32y"
	"github.com/micvbang/go-helpy/slicey"
)

// Code generated. DO NOT EDIT.

func TestUnique(t *testing.T) {
	const numTestCases = 50

	type test struct {
		input    []float32
		expected []float32
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = UniqueGenerateTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			got := float32y.Unique(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func UniqueGenerateTestCase(n int) (test []float32, expected []float32) {
	// This test generator is somewhat silly, in that it implements
	// Unique internally..

	test = make([]float32, n)
	expected = make([]float32, 0, n)
	seen := make(map[float32]struct{})

	for i := 0; i < n; i++ {
		v := float32y.RandomN(50)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}
