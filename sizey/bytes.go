package sizey

import (
	"fmt"
	"math"

	"github.com/micvbang/go-helpy"
)

const (
	B  = 1 + iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

// FormatBytes returns a human readable formatting of bytes, e.g. "4.2MiB".
func FormatBytes[T helpy.Number](n T) string {
	bf := float64(n)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei"} {
		if math.Abs(bf) < 1024.0 {
			if math.Trunc(bf) == bf {
				return fmt.Sprintf("%.0f%sB", bf, unit)
			}
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}

	return fmt.Sprintf("%.0fdB", bf)
}
