package slicey_test

import (
	"testing"

	"github.com/micvbang/go-helpy/slicey"
)

func TestRandomNilEmptySlices(t *testing.T) {
	tests := map[string][]string{
		"nil":   nil,
		"empty": {},
	}

	for name, testSlice := range tests {
		t.Run(name, func(t *testing.T) {
			got := slicey.Random(testSlice)
			if got != "" {
				t.Errorf("expected empty string, got %s", got)
			}
		})
	}
}

func TestRandomOneElement(t *testing.T) {
	expected := 42
	got := slicey.Random([]int{expected})
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestRandomMultipleElements(t *testing.T) {
	numValues := 500
	values := make([]int, numValues)
	for i := 0; i < numValues; i++ {
		values[i] = i
	}

	for i := 0; i < 100000; i++ {
		got := slicey.Random(values)
		if got < 0 || got > numValues-1 {
			t.Errorf("expected value in range [0;%d[, got %d", numValues, got)
		}
	}
}
