package booly

import "math/rand"

// Random returns a pseudo-random boolean value
func Random() bool {
	return rand.Intn(2) == 0
}
