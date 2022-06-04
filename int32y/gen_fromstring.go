package int32y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int32.
func FromString(s string) (int32, error) {

	v, err := strconv.ParseInt(s, 10, 32)

	return int32(v), err
}

// FromStringOrDefault parses s and returns a int32.
func FromStringOrDefault(s string, defaultVal int32) int32 {

	v, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return defaultVal
	}
	return int32(v)
}
