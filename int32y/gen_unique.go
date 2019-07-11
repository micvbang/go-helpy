package int32y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique int32s from the given input
func Unique(vs []int32) []int32 {
	output := make([]int32, 0, len(vs))
	seen := make(map[int32]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			output = append(output, v)
		}
	}

	return output
}
