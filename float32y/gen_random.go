package float32y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a random value of type float32
func Random() float32 {
	return float32(rand.Float64())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type float32.
func RandomN(n float32) float32 {
	return float32(rand.Float64()) * n
}
