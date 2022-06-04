package stringy

import (
	"github.com/micvbang/go-helpy"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given string.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v string) *string {
	return helpy.Pointer(v)
}

// StringOrDefault returns string if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func StringOrDefault(v *string, defaultVal string) *string {
	return helpy.DerefOrValue(v, defaultVal)
}
