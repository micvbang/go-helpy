package inty

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique ints from the given input
func Unique(vs []int) []int {
	output := make([]int, 0, len(vs))
	seen := make(map[int]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
