package int32y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int32.
func Pointer(v int32) *int32 {
	return &v
}

// Int32 dereferences and returns int32. The int32 default value is returned if v is nil.
func Int32(v *int32) int32 {
	if v == nil {
		var dv int32
		return dv
	}
	return *v
}

// Int32OrDefault returns int32 if it is not nil, and a pointer to defaultVal otherwise.
func Int32OrDefault(v *int32, defaultVal int32) *int32 {
	if v == nil {
		return &defaultVal
	}
	return v
}
