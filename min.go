package helpy

// Min returns the smallest element among v and vs.
func Min[T ComparableNumber](v T, vs ...T) T {
	min := v
	for _, v := range vs {
		if v < min {
			min = v
		}
	}

	return min
}
