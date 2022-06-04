package slicey

// Equal returns true if vs1 and vs2 are equal (contain the same elements in
// the same order.)
func Equal[T comparable](vs1 []T, vs2 []T) bool {
	if len(vs1) != len(vs2) {
		return false
	}

	for i := 0; i < len(vs1); i++ {
		if vs1[i] != vs2[i] {
			return false
		}
	}

	return true
}
