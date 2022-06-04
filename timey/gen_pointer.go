package timey

import (
	"github.com/micvbang/go-helpy"

	"time"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given time.Time.
// NOTE: this method is deprecated. Use helpy.Pointer instead.
func Pointer(v time.Time) *time.Time {
	return helpy.Pointer(v)
}

// TimeOrDefault returns time.Time if it is not nil, and a pointer to defaultVal otherwise.
// NOTE: this method is deprecated. Use helpy.DerefOrValue instead.
func TimeOrDefault(v *time.Time, defaultVal time.Time) *time.Time {
	return helpy.DerefOrValue(v, defaultVal)
}
