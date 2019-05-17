package uint8y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique uint8s from the given input
func Unique(vs []uint8) []uint8 {
	output := make([]uint8, 0, len(vs))
	seen := make(map[uint8]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
