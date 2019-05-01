package stringy

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given string.
func Pointer(v string) *string {
	return &v
}

// String dereferences and returns string. If v is nil, the default value is returned.
func String(v *string) string {
	if v == nil {
		var dv string
		return dv
	}
	return *v
}

}
