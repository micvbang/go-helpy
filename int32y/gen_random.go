package int32y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type int32
func Random() int32 {
	return int32(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type int32. It panics if n <= 0.
func RandomN(n int32) int32 {
	return int32(rand.Int63n(int64(n)))
}
