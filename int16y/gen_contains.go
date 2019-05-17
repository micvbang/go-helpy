package int16y

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []int16, v int16) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
