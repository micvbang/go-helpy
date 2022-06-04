package slicey

func Find[T any](vs []T, f func(T) bool) (T, bool) {
	for _, v := range vs {
		if f(v) {
			return v, true
		}
	}
	var v T
	return v, false
}
