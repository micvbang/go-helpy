package uint8y_test

import (
	"testing"

	"github.com/micvbang/go-helpy/uint8y"
)

// Code generated. DO NOT EDIT.

func TestMin(t *testing.T) {
	tests := map[string]struct {
		input    []uint8
		expected uint8
	}{
		"in order":      {input: []uint8{1, 2}, expected: 1},
		"reverse order": {input: []uint8{5, 2}, expected: 2},
		"3 arguments":   {input: []uint8{10, 5, 15}, expected: 5},
		"4 arguments":   {input: []uint8{1, 122, 100}, expected: 1},
		"10 arguments":  {input: []uint8{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := uint8y.Min(v, vs...)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := map[string]struct {
		input    []uint8
		expected uint8
	}{
		"in order":      {input: []uint8{1, 2}, expected: 2},
		"reverse order": {input: []uint8{5, 2}, expected: 5},
		"3 arguments":   {input: []uint8{10, 5, 15}, expected: 15},
		"4 arguments":   {input: []uint8{1, 122, 100}, expected: 122},
		"10 arguments":  {input: []uint8{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 125},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := uint8y.Max(v, vs...)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
