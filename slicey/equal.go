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

// EqualFunc returns true if vs1 and vs2 are of the same length and f returns
// true for all pairs from vs1 and vs2 to be compared.
func EqualFunc[T any](vs1 []T, vs2 []T, f func(v1 T, v2 T) bool) bool {
	if len(vs1) != len(vs2) {
		return false
	}

	for i := 0; i < len(vs1); i++ {
		if !f(vs1[i], vs2[i]) {
			return false
		}
	}

	return true
}
