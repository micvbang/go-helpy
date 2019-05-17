package int16y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type int16
func Random() int16 {
	return int16(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type int16. It panics if n <= 0.
func RandomN(n int16) int16 {
	return int16(rand.Int63n(int64(n)))
}
