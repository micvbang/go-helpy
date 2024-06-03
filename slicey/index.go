package slicey

// Last returns the last element of vs. If slice is empty, vs type's default
// value is returned.
func Last[T any](vs []T) T {
	if len(vs) == 0 {
		var t T
		return t
	}

	return vs[len(vs)-1]
}

// Last returns the last element of vs. If slice is empty, def is
// returned.
func LastOrDefault[T any](vs []T, def T) T {
	if len(vs) == 0 {
		return def
	}

	return vs[len(vs)-1]
}
