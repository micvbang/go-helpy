package int32y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int32.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v int32) *int32 {
	return helpy.Pointer(v)
}

// Int32OrDefault returns int32 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Int32OrDefault(v *int32, defaultVal int32) *int32 {
	return helpy.DerefOrValue(v, defaultVal)
}
