package int16y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int16.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v int16) *int16 {
	return helpy.Pointer(v)
}

// Int16OrDefault returns int16 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Int16OrDefault(v *int16, defaultVal int16) *int16 {
	return helpy.DerefOrValue(v, defaultVal)
}
