package uint16y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint16.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v uint16) *uint16 {
	return helpy.Pointer(v)
}

// Uint16OrDefault returns uint16 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Uint16OrDefault(v *uint16, defaultVal uint16) *uint16 {
	return helpy.DerefOrValue(v, defaultVal)
}
