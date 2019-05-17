package uint16y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique uint16s from the given input
func Unique(vs []uint16) []uint16 {
	output := make([]uint16, 0, len(vs))
	seen := make(map[uint16]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
