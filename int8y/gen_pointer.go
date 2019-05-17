package int8y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int8.
func Pointer(v int8) *int8 {
	return &v
}

// Int8 dereferences and returns int8. The int8 default value is returned if v is nil.
func Int8(v *int8) int8 {
	if v == nil {
		var dv int8
		return dv
	}
	return *v
}

// Int8OrDefault returns int8 if it is not nil, and a pointer to defaultVal otherwise.
func Int8OrDefault(v *int8, defaultVal int8) *int8 {
	if v == nil {
		return &defaultVal
	}
	return v
}
