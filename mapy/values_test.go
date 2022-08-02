package mapy_test

import (
	"testing"

	"github.com/micvbang/go-helpy"
	"github.com/micvbang/go-helpy/mapy"
)

func TestValues(t *testing.T) {
	tests := map[string]struct {
		m        map[string]string
		expected helpy.Set[string]
		err      error
	}{
		"empty": {
			m:        map[string]string{},
			expected: helpy.MakeSet[string](),
		},
		"one element": {
			m: map[string]string{
				"key": "value",
			},
			expected: helpy.MakeSet("value"),
		},
		"multiple elements": {
			m: map[string]string{
				"key1": "val1",
				"key2": "val2",
				"key3": "val3",
				"key4": "val4",
			},
			expected: helpy.MakeSet("val1", "val2", "val3", "val4"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := mapy.Values(test.m)
			gotSet := helpy.ToSet(got)
			if !gotSet.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
