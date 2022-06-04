package uint32y

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

// Code generated. DO NOT EDIT.

func TestSort(t *testing.T) {
	tests := map[string]struct {
		input    []uint32
		expected []uint32
	}{
		"in order": {
			input:    []uint32{1, 2, 3, 4, 5},
			expected: []uint32{1, 2, 3, 4, 5},
		},
		"reverse order": {
			input:    []uint32{5, 4, 3, 2, 1},
			expected: []uint32{1, 2, 3, 4, 5},
		},
		"random order": {
			input:    []uint32{4, 1, 2, 5, 3},
			expected: []uint32{1, 2, 3, 4, 5},
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
