package timey_test

import (
	"testing"
	"time"

	"github.com/micvbang/go-helpy/timey"
)

func TestMin(t *testing.T) {
	t0 := time.Now()

	at := func(hours int) time.Time {
		return t0.Add(time.Duration(hours) * time.Hour)
	}

	tests := map[string]struct {
		input    []time.Time
		expected time.Time
	}{
		"in order":      {input: []time.Time{at(0), at(1)}, expected: at(0)},
		"reverse order": {input: []time.Time{at(-5), at(2)}, expected: at(-5)},
		"3 arguments":   {input: []time.Time{at(10), at(5), at(15)}, expected: at(5)},
		"4 arguments":   {input: []time.Time{at(9), at(1), at(122), at(100)}, expected: at(1)},
		"10 arguments": {
			input:    []time.Time{at(17), at(25), at(1), at(0), at(-101), at(125), at(-42), at(13), at(37), at(69)},
			expected: at(-101),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := timey.Min(v, vs...)
			if !got.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}

func TestMax(t *testing.T) {
	t0 := time.Now()

	at := func(hours int) time.Time {
		return t0.Add(time.Duration(hours) * time.Hour)
	}

	tests := map[string]struct {
		input    []time.Time
		expected time.Time
	}{
		"in order":      {input: []time.Time{at(0), at(1)}, expected: at(1)},
		"reverse order": {input: []time.Time{at(-5), at(2)}, expected: at(2)},
		"3 arguments":   {input: []time.Time{at(10), at(5), at(15)}, expected: at(15)},
		"4 arguments":   {input: []time.Time{at(9), at(1), at(122), at(-100)}, expected: at(122)},
		"10 arguments": {
			input:    []time.Time{at(17), at(25), at(1), at(0), at(-101), at(125), at(42), at(13), at(37), at(69)},
			expected: at(125),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, vs := test.input[0], test.input[1:len(test.input)]
			got := timey.Max(v, vs...)
			if !got.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
