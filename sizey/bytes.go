package sizey

const (
	_      = iota // ignore first value by assigning to blank identifier
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)
