package float32y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given float32.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v float32) *float32 {
	return helpy.Pointer(v)
}

// Float32OrDefault returns float32 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Float32OrDefault(v *float32, defaultVal float32) *float32 {
	return helpy.DerefOrValue(v, defaultVal)
}
