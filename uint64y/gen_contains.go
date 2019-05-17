package uint64y

// Code generated. DO NOT EDIT.

// Contains returns true if vs contains v.
func Contains(vs []uint64, v uint64) bool {
	for _, c := range vs {
		if c == v {
			return true
		}
	}

	return false
}
