package uint16y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint16.
func FromString(s string) (uint16, error) {

	v, err := strconv.ParseUint(s, 10, 16)

	return uint16(v), err
}

// FromStringOrDefault parses s and returns a uint16.
func FromStringOrDefault(s string, defaultVal uint16) uint16 {

	v, err := strconv.ParseUint(s, 10, 16)

	if err != nil {
		return defaultVal
	}
	return uint16(v)
}
