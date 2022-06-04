package booly

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given bool.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v bool) *bool {
	return helpy.Pointer(v)
}

// BoolOrDefault returns bool if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func BoolOrDefault(v *bool, defaultVal bool) *bool {
	return helpy.DerefOrValue(v, defaultVal)
}
