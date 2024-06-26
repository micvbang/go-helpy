package float32y_test

import (
	"testing"

	"github.com/micvbang/go-helpy/float32y"
)

// Code generated. DO NOT EDIT.

func TestFromString(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected float32
	}{
		"positive": {input: "42.1337", expected: 42.1337},
		"zero":     {input: "0", expected: 0},
		"negative": {input: "-123.321", expected: -123.321},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := float32y.FromString(test.input)
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
		defaultVal float32
		expected   float32
	}{
		"positive": {input: "42.1337", expected: 42.1337},
		"zero":     {input: "0", expected: 0},
		"default":  {input: "sdf", defaultVal: 42.5, expected: 42.5},
		"negative": {input: "-123.321", expected: -123.321},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := float32y.FromStringOrDefault(test.input, test.defaultVal)
			if v != test.expected {
				t.Errorf("expected %v, got %v", test.expected, v)
			}
		})
	}
}
