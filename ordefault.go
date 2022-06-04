package helpy

// ValueOrDefault returns v if it isn't the default value of T, and otherwise
// returns defaultVal.
func ValueOrDefault[T comparable](v T, defaultVal T) T {
	var d T
	if v == d {
		return defaultVal
	}
	return v
}
