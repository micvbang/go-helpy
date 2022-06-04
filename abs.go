package helpy

func Abs[T Number](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
