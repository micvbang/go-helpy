package timey

import (
	"time"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given time.Time.
func Pointer(v time.Time) *time.Time {
	return &v
}

// Time dereferences and returns time.Time. If v is nil, the default value is returned.
func Time(v *time.Time) time.Time {
	if v == nil {
		var dv time.Time
		return dv
	}
	return *v
}
