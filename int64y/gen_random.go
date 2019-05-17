package int64y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type int64
func Random() int64 {
	return int64(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type int64. It panics if n <= 0.
func RandomN(n int64) int64 {
	return int64(rand.Int63n(int64(n)))
}
