package mapy_test

import (
	"testing"

	"github.com/micvbang/go-helpy/mapy"
)

func TestFromSlice(t *testing.T) {
	type typ struct {
		key   string
		value int
	}
	tests := map[string]struct {
		slice    []typ
		expected map[string]int
	}{
		"nil slice": {
			slice:    nil,
			expected: map[string]int{},
		},
		"empty slice": {
			slice:    []typ{},
			expected: map[string]int{},
		},
		"one value": {
			slice: []typ{
				{key: "key", value: 1},
			},
			expected: map[string]int{
				"key": 1,
			},
		},
		"multiple values, same key": {
			slice: []typ{
				{key: "key", value: 1},
				{key: "key", value: 2},
				{key: "key", value: 3},
			},
			expected: map[string]int{
				"key": 3,
			},
		},
		"multiple different values": {
			slice: []typ{
				{key: "key1", value: 1},
				{key: "key2", value: 2},
				{key: "key3", value: 3},
				{key: "key4", value: 4},
			},
			expected: map[string]int{
				"key1": 1,
				"key2": 2,
				"key3": 3,
				"key4": 4,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := mapy.FromSlice(test.slice, func(v typ) (string, int) {
				return v.key, v.value
			})

			if len(test.expected) != len(got) {
				t.Errorf("expected %v, got %v", test.expected, got)
			}
			for key, value := range test.expected {
				gotValue := got[key]
				if gotValue != value {
					t.Errorf("expected %v, got %v", value, gotValue)
				}
			}
		})
	}
}
