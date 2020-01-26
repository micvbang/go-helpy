package uint32y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint32.
func FromString(s string) (uint32, error) {

	v, err := strconv.ParseUint(s, 10, 32)

	return uint32(v), err
}
