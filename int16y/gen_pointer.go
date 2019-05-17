package int16y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int16.
func Pointer(v int16) *int16 {
	return &v
}

// Int16 dereferences and returns int16. The int16 default value is returned if v is nil.
func Int16(v *int16) int16 {
	if v == nil {
		var dv int16
		return dv
	}
	return *v
}

// Int16OrDefault returns int16 if it is not nil, and a pointer to defaultVal otherwise.
func Int16OrDefault(v *int16, defaultVal int16) *int16 {
	if v == nil {
		return &defaultVal
	}
	return v
}
