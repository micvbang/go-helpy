package float64y

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique float64s from the given input
func Unique(vs []float64) []float64 {
	output := make([]float64, 0, len(vs))
	seen := make(map[float64]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}
