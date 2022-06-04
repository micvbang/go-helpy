package uint8y

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint8.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v uint8) *uint8 {
	return helpy.Pointer(v)
}

// Uint8OrDefault returns uint8 if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func Uint8OrDefault(v *uint8, defaultVal uint8) *uint8 {
	return helpy.DerefOrValue(v, defaultVal)
}
