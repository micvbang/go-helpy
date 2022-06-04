package int8y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int8.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v int8) *int8 {
	return helpy.Pointer(v)
}

// Int8OrDefault returns int8 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Int8OrDefault(v *int8, defaultVal int8) *int8 {
	return helpy.DerefOrValue(v, defaultVal)
}
