package uint32y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint32.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v uint32) *uint32 {
	return helpy.Pointer(v)
}

// Uint32OrDefault returns uint32 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Uint32OrDefault(v *uint32, defaultVal uint32) *uint32 {
	return helpy.DerefOrValue(v, defaultVal)
}
