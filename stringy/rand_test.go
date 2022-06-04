package stringy_test

import (
	"testing"

	"github.com/micvbang/go-helpy/stringy"
)

func TestRandomN(t *testing.T) {
	tests := map[string]struct {
		expected int
	}{
		"empty": {expected: 0},
		"one":   {expected: 1},
		"many":  {expected: 100000},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := stringy.RandomN(test.expected)
			if len(got) != test.expected {
				t.Errorf("expected string of length %v, got %v", test.expected, got)
			}
		})
	}
}
