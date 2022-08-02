package mapy

func Map[K comparable, V any, T any](m map[K]V, f func(K, V) T) []T {
	output := make([]T, 0, len(m))
	for k, v := range m {
		output = append(output, f(k, v))
	}
	return output
}
