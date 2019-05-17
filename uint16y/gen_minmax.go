package uint16y

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v uint16, vs ...uint16) uint16 {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value from v and vs.
func Max(v uint16, vs ...uint16) uint16 {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}
