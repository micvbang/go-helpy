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

// BoolOrDefault returns bool if it is not nil, and a pointer to defaultVal otherwise.
func BoolOrDefault(v *bool, defaultVal bool) *bool {
	if v == nil {
		return &defaultVal
	}
	return v
}
