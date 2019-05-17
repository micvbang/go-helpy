package uinty

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v uint, vs ...uint) uint {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value from v and vs.
func Max(v uint, vs ...uint) uint {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}
