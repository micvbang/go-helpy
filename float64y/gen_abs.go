package float64y

// Code generated. DO NOT EDIT.

// Abs returns the absolute value of v.
func Abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
