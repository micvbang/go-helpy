package uint64y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint64.
func FromString(s string) (uint64, error) {

	v, err := strconv.ParseUint(s, 10, 64)

	return uint64(v), err
}
