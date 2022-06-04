package stringy

import "math/rand"

var runes []rune = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)

const (
	letterIdxBits = 7                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandomN(length int) string {
	b := make([]rune, length)
	for i := 0; i < length; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(runes) {
			b[i] = runes[idx]
			i++
		}
	}
	return string(b)
}
