package inty

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v int, vs ...int) int {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return v
}

// Max returns the maximum value from v and vs.
func Max(v int, vs ...int) int {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return v
}
