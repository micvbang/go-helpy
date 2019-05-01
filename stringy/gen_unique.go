package stringy

// Code generated. DO NOT EDIT.

// Unique returns a new list containing unique strings from the given input
func Unique(vs []string) []string {
	output := make([]string, 0, len(vs))
	seen := make(map[string]struct{}, len(vs))

	for _, v := range vs {
		if _, exists := seen[v]; !exists {
			output = append(output, v)
		}
	}

	return output
}

