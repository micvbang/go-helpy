package float32y

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v float32, vs ...float32) float32 {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value from v and vs.
func Max(v float32, vs ...float32) float32 {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}
