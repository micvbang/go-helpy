package slicey

import "github.com/micvbang/go-helpy"

// Sum returns the sum of all values in vs.
func Sum[T helpy.Number](vs []T) T {
	var sum T
	for _, v := range vs {
		sum += v
	}
	return sum
}
