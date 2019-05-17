package uint16y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint16.
func Pointer(v uint16) *uint16 {
	return &v
}

// Uint16 dereferences and returns uint16. The uint16 default value is returned if v is nil.
func Uint16(v *uint16) uint16 {
	if v == nil {
		var dv uint16
		return dv
	}
	return *v
}

// Uint16OrDefault returns uint16 if it is not nil, and a pointer to defaultVal otherwise.
func Uint16OrDefault(v *uint16, defaultVal uint16) *uint16 {
	if v == nil {
		return &defaultVal
	}
	return v
}
