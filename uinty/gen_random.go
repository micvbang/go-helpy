package uinty

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type uint
func Random() uint {
	return uint(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type uint. It panics if n <= 0.
func RandomN(n uint) uint {
	return uint(rand.Int63n(int64(n)))
}
