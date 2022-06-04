package int64y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int64.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v int64) *int64 {
	return helpy.Pointer(v)
}

// Int64OrDefault returns int64 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Int64OrDefault(v *int64, defaultVal int64) *int64 {
	return helpy.DerefOrValue(v, defaultVal)
}
