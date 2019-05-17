package int8y

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v int8, vs ...int8) int8 {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value from v and vs.
func Max(v int8, vs ...int8) int8 {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}
