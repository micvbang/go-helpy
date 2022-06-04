package float64y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given float64.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v float64) *float64 {
	return helpy.Pointer(v)
}

// Float64OrDefault returns float64 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Float64OrDefault(v *float64, defaultVal float64) *float64 {
	return helpy.DerefOrValue(v, defaultVal)
}
