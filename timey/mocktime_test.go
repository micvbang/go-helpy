package timey

import (
	"testing"
	"time"
)

// TestNewMockTime verifies that NewMockTime defaults to time.Now() when
// given nil.
func TestNewMockTimeNil(t *testing.T) {
	mt := NewMockTime(nil)
	expected := time.Now()
	got := mt.Now()

	if !DiffEqual(time.Millisecond, got, expected) {
		t.Errorf("expected default value to be time.Now()")
	}
}

// TestNewMockTime verifies that NewMockTime uses the inputs its given as the
// starting point.
func TestNewMockTime(t *testing.T) {
	expected := time.Now().Add(24 * time.Hour)

	mt := NewMockTime(&expected)
	got := mt.Now()

	if !DiffEqual(time.Millisecond, got, expected) {
		t.Errorf("expected default value to be time.Now()")
	}
}

// TestMockTimeAdd verifies that calls to Add increments the time by the given
// durations, and that MockTime.Now() returns the incremented value.
func TestMockTimeAdd(t *testing.T) {
	now := time.Now()

	mt := NewMockTime(&now)

	diffs := []time.Duration{
		5 * time.Second,
		24 * time.Minute,
		1 * time.Hour,
	}

	expected := now
	for _, diff := range diffs {
		expected = expected.Add(diff)
		mt.Add(diff)
		got := mt.Now()

		if !DiffEqual(time.Millisecond, got, expected) {
			t.Errorf("expected Add() to increment time")
		}
	}
}
