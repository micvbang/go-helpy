package uint64y

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v uint64, vs ...uint64) uint64 {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value from v and vs.
func Max(v uint64, vs ...uint64) uint64 {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}
