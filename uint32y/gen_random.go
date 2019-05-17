package uint32y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type uint32
func Random() uint32 {
	return uint32(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type uint32. It panics if n <= 0.
func RandomN(n uint32) uint32 {
	return uint32(rand.Int63n(int64(n)))
}
