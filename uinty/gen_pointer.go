package uinty

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint.
func Pointer(v uint) *uint {
	return &v
}

// Uint dereferences and returns uint. The uint default value is returned if v is nil.
func Uint(v *uint) uint {
	if v == nil {
		var dv uint
		return dv
	}
	return *v
}

// UintOrDefault returns uint if it is not nil, and a pointer to defaultVal otherwise.
func UintOrDefault(v *uint, defaultVal uint) *uint {
	if v == nil {
		return &defaultVal
	}
	return v
}
