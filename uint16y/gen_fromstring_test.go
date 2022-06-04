package uint16y_test

import (
	"testing"

	"github.com/micvbang/go-helpy/uint16y"
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

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := uint16y.FromString(test.input)
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
		defaultVal uint16
		expected   uint16
	}{
		"positive": {input: "42", expected: 42},
		"zero":     {input: "0", expected: 0},
		"default":  {input: "sdf", defaultVal: 42, expected: 42},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := uint16y.FromStringOrDefault(test.input, test.defaultVal)
			if v != test.expected {
				t.Errorf("expected %v, got %v", test.expected, v)
			}
		})
	}
}
