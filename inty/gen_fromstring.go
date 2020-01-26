package inty

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int.
func FromString(s string) (int, error) {

	v, err := strconv.ParseInt(s, 10, 32)

	return int(v), err
}
