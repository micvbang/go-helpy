package int8y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int8.
func FromString(s string) (int8, error) {

	v, err := strconv.ParseInt(s, 10, 8)

	return int8(v), err
}

// FromStringOrDefault parses s and returns a int8.
func FromStringOrDefault(s string, defaultVal int8) int8 {

	v, err := strconv.ParseInt(s, 10, 8)

	if err != nil {
		return defaultVal
	}
	return int8(v)
}
