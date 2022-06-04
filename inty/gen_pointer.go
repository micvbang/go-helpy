package inty

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given int.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v int) *int {
	return helpy.Pointer(v)
}

// IntOrDefault returns int if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func IntOrDefault(v *int, defaultVal int) *int {
	return helpy.DerefOrValue(v, defaultVal)
}
