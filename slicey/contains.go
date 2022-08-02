package slicey

// Contains returns true if needle exists in vs.
func Contains[T comparable](vs []T, needle T) bool {
	for _, v := range vs {
		if v == needle {
			return true
		}
	}
	return false
}

// ContainsFunc returns true if f returns true for any element in vs.
func ContainsFunc[T any](vs []T, f func(v T) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
