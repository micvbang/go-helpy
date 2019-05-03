package inty

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []int, v int) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
