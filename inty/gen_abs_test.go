package inty

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected int
	}{
		"positive": {input: 42, expected: 42},
		"negative": {input: -123, expected: 123},
		"zero":     {input: 0, expected: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expected, Abs(tc.input))
		})
	}
}
