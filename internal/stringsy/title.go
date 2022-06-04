package stringsy

import "unicode"

func TitleCasing(s string, runeCaser func(rune) rune) string {
	if len(s) == 0 {
		return s
	}

	r := rune(s[0])
	if !unicode.IsUpper(r) && !unicode.IsLower(r) {
		return s
	}
	return string(runeCaser(r)) + s[1:]
}
