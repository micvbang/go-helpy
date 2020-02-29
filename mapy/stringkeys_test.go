package mapy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringKeys(t *testing.T) {
	tests := map[string]struct {
		m    interface{}
		keys []string
		err  error
	}{
		"map[string]string": {
			m: map[string]string{
				"key1": "val1",
				"key2": "val2",
			},
			keys: []string{"key1", "key2"},
		},
		"map[string]int": {
			m: map[string]int{
				"key1": 1,
				"key2": 2,
			},
			keys: []string{"key1", "key2"},
		},
		"map[string]bool": {
			m: map[string]bool{
				"key1": true,
				"key2": false,
			},
			keys: []string{"key1", "key2"},
		},
		"map[int]string": {
			m: map[int]string{
				1: "one",
				2: "two",
			},
			err: ErrInvalidKeyType{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			keys, err := StringKeys(test.m)
			if test.err != nil {
				require.IsType(t, test.err, err)
			}

			require.Equal(t, test.keys, keys)
		})
	}
}
