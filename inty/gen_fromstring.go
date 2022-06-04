package inty

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int.
func FromString(s string) (int, error) {

	v, err := strconv.ParseInt(s, 10, 32)

	return int(v), err
}

// FromStringOrDefault parses s and returns a int.
func FromStringOrDefault(s string, defaultVal int) int {

	v, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return defaultVal
	}
	return int(v)
}
