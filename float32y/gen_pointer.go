package float32y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given float32.
func Pointer(v float32) *float32 {
	return &v
}

// Float32 dereferences and returns float32. The float32 default value is returned if v is nil.
func Float32(v *float32) float32 {
	if v == nil {
		var dv float32
		return dv
	}
	return *v
}

// Float32OrDefault returns float32 if it is not nil, and a pointer to defaultVal otherwise.
func Float32OrDefault(v *float32, defaultVal float32) *float32 {
	if v == nil {
		return &defaultVal
	}
	return v
}
