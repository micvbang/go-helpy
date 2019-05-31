package booly

import (
	"strconv"
	"strings"
)

// ToString returns "true" or "false", matching the input bool.
func ToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// FromString parses the string to a bool.
// Valid input values are any casing of "true" and "false", and 1 and 0.
// If any other input is given, FromString returns false.
func FromString(s string) bool {
	b, _ := strconv.ParseBool(strings.ToLower(s))
	// Ignore errors and return false on errors
	return b
}
