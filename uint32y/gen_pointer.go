package uint32y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint32.
func Pointer(v uint32) *uint32 {
	return &v
}

// Uint32 dereferences and returns uint32. The uint32 default value is returned if v is nil.
func Uint32(v *uint32) uint32 {
	if v == nil {
		var dv uint32
		return dv
	}
	return *v
}

// Uint32OrDefault returns uint32 if it is not nil, and a pointer to defaultVal otherwise.
func Uint32OrDefault(v *uint32, defaultVal uint32) *uint32 {
	if v == nil {
		return &defaultVal
	}
	return v
}
