package timey

import (
	"math/rand"
	"time"
)

// RandomDurationN returns a non-negative pseudo-random duration in [0,n). It
// panics if d <= 0.
func RandomDurationN(d time.Duration) time.Duration {
	return time.Duration(rand.Int63n(int64(d)))
}
