package uinty

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

// Code generated. DO NOT EDIT.

func TestSort(t *testing.T) {
	tests := map[string]struct {
		input    []uint
		expected []uint
	}{
		"in order": {
			input:    []uint{1, 2, 3, 4, 5},
			expected: []uint{1, 2, 3, 4, 5},
		},
		"reverse order": {
			input:    []uint{5, 4, 3, 2, 1},
			expected: []uint{1, 2, 3, 4, 5},
		},
		"random order": {
			input:    []uint{4, 1, 2, 5, 3},
			expected: []uint{1, 2, 3, 4, 5},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := Sort(test.input)
			if !slicey.Equal(test.expected, got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
