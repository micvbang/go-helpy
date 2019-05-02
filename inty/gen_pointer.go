package inty

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int.
func Pointer(v int) *int {
	return &v
}

// Int dereferences and returns int. The int default value is returned if v is nil.
func Int(v *int) int {
	if v == nil {
		var dv int
		return dv
	}
	return *v
}

// IntOrDefault dereferences and returns int. defaultVal is returned if v is nil.
func IntOrDefault(v *int, defaultVal int) int {
	if v == nil {
		var dv int
		return dv
	}
	return *v
}
