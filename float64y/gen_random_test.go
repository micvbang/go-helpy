package float64y

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
	for i := 1; i < 10000; i++ {
		v := float64(i)
		got := RandomN(v)
		if got < 0 || got >= v {
			t.Errorf("random number outside expected range [0;%v[", v)
		}
	}
}
