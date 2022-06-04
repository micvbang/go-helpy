package inty_test

import (
	"testing"

	"github.com/micvbang/go-helpy/inty"
)

// Code generated. DO NOT EDIT.

func TestMin(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"in order":      {input: []int{1, 2}, expected: 1},
		"reverse order": {input: []int{5, 2}, expected: 2},
		"3 arguments":   {input: []int{10, 5, 15}, expected: 5},
		"4 arguments":   {input: []int{1, 122, 100}, expected: 1},
		"10 arguments":  {input: []int{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 0},

		"one negative":          {input: []int{-5, 10}, expected: -5},
		"both negative":         {input: []int{-5, -100}, expected: -100},
		"10 arguments negative": {input: []int{17, 25, -1, 0, -101, 127, 42, -13, 37, 69}, expected: -101},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := inty.Min(v, vs...)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"in order":      {input: []int{1, 2}, expected: 2},
		"reverse order": {input: []int{5, 2}, expected: 5},
		"3 arguments":   {input: []int{10, 5, 15}, expected: 15},
		"4 arguments":   {input: []int{1, 122, 100}, expected: 122},
		"10 arguments":  {input: []int{17, 25, 1, 0, 101, 125, 42, 13, 37, 69}, expected: 125},

		"one negative":          {input: []int{-5, 10}, expected: 10},
		"both negative":         {input: []int{-5, -100}, expected: -5},
		"10 arguments negative": {input: []int{17, 25, -1, 0, -101, 127, 42, -13, 37, 69}, expected: 127},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := inty.Max(v, vs...)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
