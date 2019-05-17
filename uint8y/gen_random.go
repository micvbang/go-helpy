package uint8y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type uint8
func Random() uint8 {
	return uint8(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type uint8. It panics if n <= 0.
func RandomN(n uint8) uint8 {
	return uint8(rand.Int63n(int64(n)))
}
