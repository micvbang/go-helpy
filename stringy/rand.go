package stringy

import "math/rand"

const letterBytes = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_!./#¤%&/()=[]{}'"`
const (
	letterIdxBits = 7                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandomN(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}
