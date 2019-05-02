package int64y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int64.
func Pointer(v int64) *int64 {
	return &v
}

// Int64 dereferences and returns int64. The int64 default value is returned if v is nil.
func Int64(v *int64) int64 {
	if v == nil {
		var dv int64
		return dv
	}
	return *v
}

// Int64OrDefault returns int64 if it is not nil, and a pointer to defaultVal otherwise.
func Int64OrDefault(v *int64, defaultVal int64) *int64 {
	if v == nil {
		return &defaultVal
	}
	return v
}
