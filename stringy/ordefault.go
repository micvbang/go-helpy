package stringy

// OrDefault returns s if its length is > 0, and defaultVal otherwise.
func OrDefault(s string, defaultVal string) string {
	if len(s) > 0 {
		return s
	}
	return defaultVal
}
