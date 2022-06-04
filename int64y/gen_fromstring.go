package int64y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a int64.
func FromString(s string) (int64, error) {

	v, err := strconv.ParseInt(s, 10, 64)

	return int64(v), err
}

// FromStringOrDefault parses s and returns a int64.
func FromStringOrDefault(s string, defaultVal int64) int64 {

	v, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return defaultVal
	}
	return int64(v)
}
