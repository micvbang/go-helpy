package mapy_test

import (
	"testing"

	"github.com/micvbang/go-helpy"
	"github.com/micvbang/go-helpy/mapy"
)

func TestKeys(t *testing.T) {
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
				"key1": "value",
			},
			expected: helpy.MakeSet("key1"),
		},
		"map[string]bool": {
			m: map[string]string{
				"key1": "value",
				"key2": "value",
				"key3": "value",
				"key4": "value",
			},
			expected: helpy.MakeSet("key1", "key2", "key3", "key4"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := mapy.Keys(test.m)
			gotSet := helpy.ToSet(got)
			if !gotSet.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
