package int64y

import (
	"testing"
)

// Code generated. DO NOT EDIT.

// TestRandom simply runs Random to ensure that it compiles during testing.
func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		Random()
	}
}

// TestRandomN runs RandomN and ensures that output is within expected range.
func TestRandomN(t *testing.T) {
	for i := 1; i < 127; i++ {
		v := int64(i)
		got := RandomN(v)
		if got >= v {
			t.Errorf("expected value < %v, got %v", v, got)
		}

		if got >= v {
			t.Errorf("expected value > 0, got %v", got)
		}

	}
}
