package helpy_test

import (
	"testing"

	"github.com/micvbang/go-helpy"
)

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected int
	}{
		"positive": {input: 42, expected: 42},
		"negative": {input: -123, expected: 123},
		"zero":     {input: 0, expected: 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := helpy.Abs(test.input)
			if got != test.expected {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
