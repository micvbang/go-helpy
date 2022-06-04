package helpy

// Max returns the largest value among v and vs.
func Max[T ComparableNumber](v T, vs ...T) T {
	max := v
	for _, v := range vs {
		if v > max {
			max = v
		}
	}

	return max
}
