package uint16y

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
		v := uint16(i)
		got := RandomN(v)
		if got >= v {
			t.Errorf("expected value < %v, got %v", v, got)
		}

	}
}
