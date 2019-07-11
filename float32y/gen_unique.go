package float32y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique float32s from the given input
func Unique(vs []float32) []float32 {
	output := make([]float32, 0, len(vs))
	seen := make(map[float32]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			output = append(output, v)
		}
	}

	return output
}
