package uint64y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type uint64
func Random() uint64 {
	return uint64(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type uint64. It panics if n <= 0.
func RandomN(n uint64) uint64 {
	return uint64(rand.Int63n(int64(n)))
}
