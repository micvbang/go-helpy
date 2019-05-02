package uint64y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint64.
func Pointer(v uint64) *uint64 {
	return &v
}

// Uint64 dereferences and returns uint64. The uint64 default value is returned if v is nil.
func Uint64(v *uint64) uint64 {
	if v == nil {
		var dv uint64
		return dv
	}
	return *v
}

// Uint64OrDefault returns uint64 if it is not nil, and a pointer to defaultVal otherwise.
func Uint64OrDefault(v *uint64, defaultVal uint64) *uint64 {
	if v == nil {
		return &defaultVal
	}
	return v
}
