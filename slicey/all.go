package slicey

// All returns true if and only if f returns true for all vs.
func All[T any](vs []T, f func(v T) bool) bool {
	if f == nil {
		return true
	}

	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}
