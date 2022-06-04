package slicey_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/inty"
	"github.com/micvbang/go-helpy/slicey"
)

// TestUnique verifies that Unique handles nil and empty inputs, and that
// it only returns unique elements.
func TestUniqueSimple(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int
	}{
		"nil input": {
			input:    nil,
			expected: []int{},
		},
		"empty input": {
			input:    []int{},
			expected: []int{},
		},
		"one element": {
			input:    []int{1},
			expected: []int{1},
		},
		"two elements -> one": {
			input:    []int{1, 1},
			expected: []int{1},
		},
		"two elements -> two": {
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if !slicey.Equal(test.expected, slicey.Unique(test.input)) {
				t.Errorf("expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

// TestUnique verifies that Unique only returns unique elements.
func TestUniqueGenerated(t *testing.T) {
	const numTestCases = 1000

	type test struct {
		input    []int
		expected []int
	}
	tests := make([]test, numTestCases)

	for i := 0; i < numTestCases; i++ {
		test := test{}
		test.input, test.expected = generateUniqueTestCase(i)
		tests[i] = test
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			if !slicey.Equal(test.expected, slicey.Unique(test.input)) {
				t.Errorf("expected %v, got %v", test.expected, test.input)
			}
		})
	}
}

func generateUniqueTestCase(elements int) (test []int, expected []int) {
	test = make([]int, elements)
	expected = make([]int, 0, elements)
	seen := make(map[int]struct{})

	for i := 0; i < elements; i++ {
		v := inty.RandomN(int(elements/2) + 10)
		test[i] = v

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			expected = append(expected, v)
		}
	}

	return test, expected
}
