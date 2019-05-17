package inty

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type int
func Random() int {
	return int(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type int. It panics if n <= 0.
func RandomN(n int) int {
	return int(rand.Int63n(int64(n)))
}
