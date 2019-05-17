package uint16y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a non-negative pseudo-random number of type uint16
func Random() uint16 {
	return uint16(rand.Int63())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type uint16. It panics if n <= 0.
func RandomN(n uint16) uint16 {
	return uint16(rand.Int63n(int64(n)))
}
