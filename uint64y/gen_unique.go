package uint64y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique uint64s from the given input
func Unique(vs []uint64) []uint64 {
	output := make([]uint64, 0, len(vs))
	seen := make(map[uint64]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
