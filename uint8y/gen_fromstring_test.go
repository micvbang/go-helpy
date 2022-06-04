package uint8y_test

import (
	"testing"

	"github.com/micvbang/go-helpy/uint8y"
)

// Code generated. DO NOT EDIT.

func TestFromString(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected uint8
	}{
		"positive": {input: "42", expected: 42},
		"zero":     {input: "0", expected: 0},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := uint8y.FromString(test.input)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if v != test.expected {
				t.Errorf("expected %v, got %v", test.expected, v)
			}
		})
	}
}

func TestFromStringOrDefault(t *testing.T) {
	tests := map[string]struct {
		input      string
		defaultVal uint8
		expected   uint8
	}{
		"positive": {input: "42", expected: 42},
		"zero":     {input: "0", expected: 0},
		"default":  {input: "sdf", defaultVal: 42, expected: 42},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := uint8y.FromStringOrDefault(test.input, test.defaultVal)
			if v != test.expected {
				t.Errorf("expected %v, got %v", test.expected, v)
			}
		})
	}
}
