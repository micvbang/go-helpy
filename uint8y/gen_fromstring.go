package uint8y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a uint8.
func FromString(s string) (uint8, error) {

	v, err := strconv.ParseUint(s, 10, 8)

	return uint8(v), err
}

// FromStringOrDefault parses s and returns a uint8.
func FromStringOrDefault(s string, defaultVal uint8) uint8 {

	v, err := strconv.ParseUint(s, 10, 8)

	if err != nil {
		return defaultVal
	}
	return uint8(v)
}
