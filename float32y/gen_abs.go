package float32y

// Code generated. DO NOT EDIT.

// Abs returns the absolute value of v.
func Abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
