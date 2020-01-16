package int64y

// Code generated. DO NOT EDIT.

// Abs returns the absolute value of v.
func Abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
