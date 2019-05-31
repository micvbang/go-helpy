package booly

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToString(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{input: true, expected: "true"},
		{input: false, expected: "false"},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, ToString(test.input))
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Truth values
		{input: "true", expected: true},
		{input: "TRUE", expected: true},
		{input: "True", expected: true},
		{input: "1", expected: true},
		{input: "TrUe", expected: true},

		// False values
		{input: "false", expected: false},
		{input: "FALSE", expected: false},
		{input: "False", expected: false},
		{input: "FaLsE", expected: false},
		{input: "0", expected: false},

		// Unhandled => false values
		{input: ".", expected: false},
		{input: "aksldjalsdkj", expected: false},
		{input: "trueish", expected: false},
		{input: "truth", expected: false},
		{input: "yes", expected: false},
		{input: "no", expected: false},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, FromString(test.input))
	}
}
