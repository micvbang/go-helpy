package helpy

// Clamp returns value if it's in the range [lower, upper], otherwise the closest
// boundary of lower and upper.
// NOTE: if lower > upper, then lower is always returned.
func Clamp[T Number](value T, lower T, upper T) T {
	return max(lower, min(value, upper))
}
