package uint16y

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []uint16, v uint16) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
