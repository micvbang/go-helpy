package int64y

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v int64, vs ...int64) int64 {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return v
}

// Max returns the maximum value from v and vs.
func Max(v int64, vs ...int64) int64 {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return v
}
