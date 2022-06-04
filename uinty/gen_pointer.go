package uinty

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given uint.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v uint) *uint {
	return helpy.Pointer(v)
}

// UintOrDefault returns uint if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func UintOrDefault(v *uint, defaultVal uint) *uint {
	return helpy.DerefOrValue(v, defaultVal)
}
