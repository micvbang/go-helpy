package stringy

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given string.
func Pointer(v string) *string {
	return &v
}

// String dereferences and returns string. The string default value is returned if v is nil.
func String(v *string) string {
	if v == nil {
		var dv string
		return dv
	}
	return *v
}

// StringOrDefault returns string if it is not nil, and a pointer to defaultVal otherwise.
func StringOrDefault(v *string, defaultVal string) *string {
	if v == nil {
		return &defaultVal
	}
	return v
}
