package int64y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique int64s from the given input
func Unique(vs []int64) []int64 {
	output := make([]int64, 0, len(vs))
	seen := make(map[int64]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			output = append(output, v)
		}
	}

	return output
}
