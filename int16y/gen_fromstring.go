package int16y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int16.
func FromString(s string) (int16, error) {

	v, err := strconv.ParseInt(s, 10, 16)

	return int16(v), err
}

// FromStringOrDefault parses s and returns a int16.
func FromStringOrDefault(s string, defaultVal int16) int16 {

	v, err := strconv.ParseInt(s, 10, 16)

	if err != nil {
		return defaultVal
	}
	return int16(v)
}
