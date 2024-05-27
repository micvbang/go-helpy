package sizey_test

import (
	"fmt"
	"testing"

	"github.com/micvbang/go-helpy/sizey"
)

func TestByteSizes(t *testing.T) {
	tests := map[string]struct {
		expected uint64
		got      uint64
	}{
		"b":  {expected: pow(1024, 0), got: sizey.B},
		"kb": {expected: pow(1024, 1), got: sizey.KB},
		"mb": {expected: pow(1024, 2), got: sizey.MB},
		"gb": {expected: pow(1024, 3), got: sizey.GB},
		"tb": {expected: pow(1024, 4), got: sizey.TB},
		"pb": {expected: pow(1024, 5), got: sizey.PB},
		"eb": {expected: pow(1024, 6), got: sizey.EB},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expected != test.got {
				t.Fatalf("expected %d, got %d", test.expected, test.got)
			}
		})
	}
}

func pow(v uint64, e int) uint64 {
	if e == 0 {
		return 1
	}

	r := uint64(1)
	for i := 0; i < e; i++ {
		r *= v
	}
	return r
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		bytes    uint64
		expected string
	}{
		{bytes: 0, expected: "0B"},
		{bytes: 1, expected: "1B"},
		{bytes: 512, expected: "512B"},
		{bytes: 1024, expected: "1KiB"},
		{bytes: 1025, expected: "1.0KiB"},
		{bytes: 1536, expected: "1.5KiB"},
		{bytes: 2047, expected: "2.0KiB"},
		{bytes: 2048, expected: "2KiB"},
		{bytes: 43520, expected: "42.5KiB"},
		{bytes: pow(1024, 0), expected: "1B"},
		{bytes: pow(1024, 1), expected: "1KiB"},
		{bytes: pow(1024, 2), expected: "1MiB"},
		{bytes: pow(1024, 3), expected: "1GiB"},
		{bytes: pow(1024, 4), expected: "1TiB"},
		{bytes: pow(1024, 5), expected: "1PiB"},
		{bytes: pow(1024, 6), expected: "1EiB"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := sizey.FormatBytes(test.bytes)
			if test.expected != got {
				t.Fatalf("expected '%s', got '%s'", test.expected, got)
			}
		})
	}
}
