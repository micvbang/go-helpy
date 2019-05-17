package uint8y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint8.
func Pointer(v uint8) *uint8 {
	return &v
}

// Uint8 dereferences and returns uint8. The uint8 default value is returned if v is nil.
func Uint8(v *uint8) uint8 {
	if v == nil {
		var dv uint8
		return dv
	}
	return *v
}

// Uint8OrDefault returns uint8 if it is not nil, and a pointer to defaultVal otherwise.
func Uint8OrDefault(v *uint8, defaultVal uint8) *uint8 {
	if v == nil {
		return &defaultVal
	}
	return v
}
