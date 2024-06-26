package float32y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a float32.
func FromString(s string) (float32, error) {
	v, err := strconv.ParseFloat(s, 32)
	return float32(v), err
}

// FromStringOrDefault parses s and returns a float32.
func FromStringOrDefault(s string, defaultVal float32) float32 {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return defaultVal
	}
	return float32(v)
}
