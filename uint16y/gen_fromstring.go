package uint16y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint16.
func FromString(s string) (uint16, error) {

	v, err := strconv.ParseUint(s, 10, 16)

	return uint16(v), err
}
