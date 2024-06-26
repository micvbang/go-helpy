package float64y

import "strconv"

// Code generated. DO NOT EDIT.

// FromString parses s and returns a float64.
func FromString(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	return float64(v), err
}

// FromStringOrDefault parses s and returns a float64.
func FromStringOrDefault(s string, defaultVal float64) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultVal
	}
	return float64(v)
}
