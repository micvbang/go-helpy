package booly_test

import (
	"testing"

	"github.com/micvbang/go-helpy/booly"
)

func TestRandom(t *testing.T) {
	const (
		samples            = 100000
		unlikelySampleSkew = samples / 3
	)

	values := map[bool]int{}
	for i := 0; i < samples; i++ {
		values[booly.Random()] += 1
	}

	if values[true] == 0 || values[false] == 0 {
		t.Errorf("expected samples of both true and false")
	}

	if values[true] <= unlikelySampleSkew || values[false] <= unlikelySampleSkew {
		pctTrue := float32(values[true]) / float32(samples)
		pctFalse := float32(values[false]) / float32(samples)
		t.Errorf("expected samples to be less skewed: true: %.2f, false: %.2f", pctTrue, pctFalse)
	}
}
