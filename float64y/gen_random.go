package float64y

import (
	"math/rand"
)

// Code generated. DO NOT EDIT.

// Random returns a random value of type float64
func Random() float64 {
	return float64(rand.Float64())
}

// RandomN returns a non-negative pseudo-random number in [0,n) of type float64.
func RandomN(n float64) float64 {
	return float64(rand.Float64()) * n
}
