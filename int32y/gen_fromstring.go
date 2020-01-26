package int32y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int32.
func FromString(s string) (int32, error) {

	v, err := strconv.ParseInt(s, 10, 32)

	return int32(v), err
}
