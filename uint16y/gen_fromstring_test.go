package uint16y

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestFromString(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected uint16
	}{
		"positive": {input: "42", expected: 42},
		"zero":     {input: "0", expected: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := FromString(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, v)
		})
	}
}
