package slicey

// Unique returns a new list containing unique values from v.
func Unique[T comparable](vs []T) []T {
	output := make([]T, 0, len(vs))
	seen := make(map[T]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			output = append(output, v)
		}
	}

	return output
}
