package int64y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int64.
func Pointer(v int64) *int64 {
	return &v
}

// Int64 dereferences and returns int64. If v is nil, the default value is returned.
func Int64(v *int64) int64 {
	if v == nil {
		var dv int64
		return dv
	}
	return *v
}

}
