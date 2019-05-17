package float32y

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []float32, v float32) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
