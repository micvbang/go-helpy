package int16y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique int16s from the given input
func Unique(vs []int16) []int16 {
	output := make([]int16, 0, len(vs))
	seen := make(map[int16]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
