package int8y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int8.
func FromString(s string) (int8, error) {

	v, err := strconv.ParseInt(s, 10, 8)

	return int8(v), err
}
