package slicey

import (
	"math/rand"
)

// Random returns a random element from vs. If vs is empty or nil,
// the default value of T is returned.
func Random[T any](vs []T) T {
	if len(vs) == 0 {
		var v T
		return v
	}

	return vs[int(rand.Int31n(int32(len(vs))))]
}
