package int8y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique int8s from the given input
func Unique(vs []int8) []int8 {
	output := make([]int8, 0, len(vs))
	seen := make(map[int8]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
