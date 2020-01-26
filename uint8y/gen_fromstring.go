package uint8y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint8.
func FromString(s string) (uint8, error) {

	v, err := strconv.ParseUint(s, 10, 8)

	return uint8(v), err
}
