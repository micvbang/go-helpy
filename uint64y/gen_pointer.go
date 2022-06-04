package uint64y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint64.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v uint64) *uint64 {
	return helpy.Pointer(v)
}

// Uint64OrDefault returns uint64 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Uint64OrDefault(v *uint64, defaultVal uint64) *uint64 {
	return helpy.DerefOrValue(v, defaultVal)
}
