package float64y

// Code generated. DO NOT EDIT.

// Min returns the minimum value from v and vs.
func Min(v float64, vs ...float64) float64 {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value from v and vs.
func Max(v float64, vs ...float64) float64 {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}
