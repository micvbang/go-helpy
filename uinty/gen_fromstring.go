package uinty

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint.
func FromString(s string) (uint, error) {

	v, err := strconv.ParseUint(s, 10, 32)

	return uint(v), err
}
