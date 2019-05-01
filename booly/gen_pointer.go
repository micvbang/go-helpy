package booly

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given bool.
func Pointer(v bool) *bool {
	return &v
}

// Bool dereferences and returns bool. If v is nil, the default value is returned.
func Bool(v *bool) bool {
	if v == nil {
		var dv bool
		return dv
	}
	return *v
}
