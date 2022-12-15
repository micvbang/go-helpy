package mapy

// FromSlice constructs a map by using f to extract keys and values from the
// given slice.
// NOTE: if f returns the same key multiple times, the map will contain the
// value of the last element in the slice for which f returns that key.
func FromSlice[T any, K comparable, V any](vs []T, f func(v T) (K, V)) map[K]V {
	m := make(map[K]V, len(vs))
	for _, v := range vs {
		key, value := f(v)
		m[key] = value
	}

	return m
}
