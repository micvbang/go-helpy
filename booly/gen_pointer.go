package booly

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given bool.
func Pointer(v bool) *bool {
	return &v
}

// Bool dereferences and returns bool. The bool default value is returned if v is nil.
func Bool(v *bool) bool {
	if v == nil {
		var dv bool
		return dv
	}
	return *v
}

// BoolOrDefault dereferences and returns bool. defaultVal is returned if v is nil.
func BoolOrDefault(v *bool, defaultVal bool) bool {
	if v == nil {
		var dv bool
		return dv
	}
	return *v
}
