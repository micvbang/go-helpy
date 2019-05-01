package inty

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int.
func Pointer(v int) *int {
	return &v
}

// Int dereferences and returns int. If v is nil, the default value is returned.
func Int(v *int) int {
	if v == nil {
		var dv int
		return dv
	}
	return *v
}
