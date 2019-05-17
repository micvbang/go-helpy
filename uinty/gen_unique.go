package uinty

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique uints from the given input
func Unique(vs []uint) []uint {
	output := make([]uint, 0, len(vs))
	seen := make(map[uint]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
