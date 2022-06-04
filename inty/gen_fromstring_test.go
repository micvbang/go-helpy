package inty_test

import (
	"testing"

	"github.com/micvbang/go-helpy/inty"
)

// Code generated. DO NOT EDIT.

func TestFromString(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"positive": {input: "42", expected: 42},
		"zero":     {input: "0", expected: 0},
		"negative": {input: "-123", expected: -123},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v, err := inty.FromString(test.input)
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
		defaultVal int
		expected   int
	}{
		"positive": {input: "42", expected: 42},
		"zero":     {input: "0", expected: 0},
		"default":  {input: "sdf", defaultVal: 42, expected: 42},
		"negative": {input: "-123", expected: -123},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			v := inty.FromStringOrDefault(test.input, test.defaultVal)
			if v != test.expected {
				t.Errorf("expected %v, got %v", test.expected, v)
			}
		})
	}
}
