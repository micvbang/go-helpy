package syncy_test

import (
	"testing"

	"github.com/micvbang/go-helpy/syncy"
)

func TestPoolSingleResource(t *testing.T) {
	p := syncy.NewPool(func() *int {
		var n = 0
		return &n
	})

	const expectedN = 10

	for i := 0; i < expectedN; i++ {
		n := p.Get()
		*n += 1
		p.Put(n)
	}

	n := p.Get()
	if *n != expectedN {
		t.Fatalf("expected n == %d, got n == %d", expectedN, *n)
	}
}
