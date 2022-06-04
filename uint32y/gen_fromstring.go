package uint32y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint32.
func FromString(s string) (uint32, error) {

	v, err := strconv.ParseUint(s, 10, 32)

	return uint32(v), err
}

// FromStringOrDefault parses s and returns a uint32.
func FromStringOrDefault(s string, defaultVal uint32) uint32 {

	v, err := strconv.ParseUint(s, 10, 32)

	if err != nil {
		return defaultVal
	}
	return uint32(v)
}
