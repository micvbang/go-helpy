package helpy

// Pointer returns a pointer to v.
func Pointer[T any](v T) *T {
	return &v
}

// DerefOrValue returns v if it is not nil, otherwise returns a pointer to
// defaultVal.
func DerefOrValue[T any](v *T, defaultVal T) *T {
	if v != nil {
		return v
	}
	return &defaultVal
}
