package mapy_test

import (
	"testing"

	"github.com/micvbang/go-helpy/mapy"
	"github.com/micvbang/go-helpy/stringy"
)

func TestKeys(t *testing.T) {
	tests := map[string]struct {
		m        map[string]interface{}
		expected stringy.Set
		err      error
	}{
		"map[string]string": {
			m: map[string]interface{}{
				"key1": "val1",
				"key2": "val2",
			},
			expected: stringy.MakeSet("key1", "key2"),
		},
		"map[string]int": {
			m: map[string]interface{}{
				"key1": 1,
				"key2": 2,
			},
			expected: stringy.MakeSet("key1", "key2"),
		},
		"map[string]bool": {
			m: map[string]interface{}{
				"key1": true,
				"key2": false,
			},
			expected: stringy.MakeSet("key1", "key2"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := mapy.Keys(test.m)
			gotSet := stringy.ToSet(got)
			if !gotSet.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}
}
