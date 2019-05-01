package booly

// ToInt returns 1 if b is true, and 0 if b is false.
func ToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
