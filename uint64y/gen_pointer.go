package uint64y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint64.
func Pointer(v uint64) *uint64 {
	return &v
}

// Uint64 dereferences and returns uint64. If v is nil, the default value is returned.
func Uint64(v *uint64) uint64 {
	if v == nil {
		var dv uint64
		return dv
	}
	return *v
}

}
