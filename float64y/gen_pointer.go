package float64y

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given float64.
func Pointer(v float64) *float64 {
	return &v
}

// Float64 dereferences and returns float64. The float64 default value is returned if v is nil.
func Float64(v *float64) float64 {
	if v == nil {
		var dv float64
		return dv
	}
	return *v
}

// Float64OrDefault returns float64 if it is not nil, and a pointer to defaultVal otherwise.
func Float64OrDefault(v *float64, defaultVal float64) *float64 {
	if v == nil {
		return &defaultVal
	}
	return v
}
