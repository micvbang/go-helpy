package slicey

// Filter returns all elements from vs for which f returns true.
func Filter[T any](vs []T, f func(T) bool) []T {
	output := make([]T, 0, len(vs))
	for _, v := range vs {
		if f(v) {
			output = append(output, v)
		}
	}
	return output
}
