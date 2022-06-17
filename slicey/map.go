package slicey

// Map uses f to map values from type T to type V.
func Map[T any, U any](vs []T, f func(T) U) []U {
	output := make([]U, len(vs))
	for i, v := range vs {
		output[i] = f(v)
	}
	return output
}
