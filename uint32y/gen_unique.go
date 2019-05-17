package uint32y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique uint32s from the given input
func Unique(vs []uint32) []uint32 {
	output := make([]uint32, 0, len(vs))
	seen := make(map[uint32]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
