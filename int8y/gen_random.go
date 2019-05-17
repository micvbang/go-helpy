package int8y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type int8
func Random() int8 {
	return int8(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type int8. It panics if n <= 0.
func RandomN(n int8) int8 {
	return int8(rand.Int63n(int64(n)))
}
