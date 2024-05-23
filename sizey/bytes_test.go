package sizey_test

import (
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
