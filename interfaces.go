package helpy

type ComparableNumber interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64 | ~uintptr
}

type Number interface {
	ComparableNumber | float32 | float64
}
