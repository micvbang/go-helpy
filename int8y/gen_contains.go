package int8y

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []int8, v int8) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
