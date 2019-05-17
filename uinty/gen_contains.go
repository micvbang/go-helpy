package uinty

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []uint, v uint) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
