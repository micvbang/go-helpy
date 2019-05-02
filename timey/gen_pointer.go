package timey

import (
	"time"
)

// Code generated. DO NOT EDIT.

// Pointer returns a pointer to the given time.Time.
func Pointer(v time.Time) *time.Time {
	return &v
}

// Time dereferences and returns time.Time. The time.Time default value is returned if v is nil.
func Time(v *time.Time) time.Time {
	if v == nil {
		var dv time.Time
		return dv
	}
	return *v
}

// TimeOrDefault returns time.Time if it is not nil, and a pointer to defaultVal otherwise.
func TimeOrDefault(v *time.Time, defaultVal time.Time) *time.Time {
	if v == nil {
		return &defaultVal
	}
	return v
}
