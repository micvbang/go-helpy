package uint8y

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []uint8, v uint8) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
