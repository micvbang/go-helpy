package mapy_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy"
	"github.com/micvbang/go-helpy/mapy"
)

func TestMap(t *testing.T) {
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
			expected: helpy.MakeSet("key1-value"),
		},
		"map[string]bool": {
			m: map[string]string{
				"key1": "val1",
				"key2": "val2",
				"key3": "val3",
				"key4": "val4",
			},
			expected: helpy.MakeSet("key1-val1", "key2-val2", "key3-val3", "key4-val4"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := mapy.Map(test.m, func(k string, v string) string {
				return fmt.Sprintf("%s-%s", k, v)
			})
			gotSet := helpy.ToSet(got)
			if !gotSet.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
		})
	}

}
