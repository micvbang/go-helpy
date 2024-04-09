package timey

import (
	"time"

	"github.com/micvbang/go-helpy"
)

// MockTime helps mock time for functions that take a func() time.Time argument
// to retrieve the time, instead of using the global time.Now().
// This is especially useful when testing functions that depend on time.
type MockTime struct {
	Time time.Time
}

// NewMockTime returns a MockTime with the given time.Time instance. If
// now is nil, it will default to time.Now().
func NewMockTime(now *time.Time) *MockTime {
	if now == nil {
		now = helpy.Pointer(time.Now())
	}

	return &MockTime{Time: *now}
}

// Now returns the current (mocked) time.
func (mt *MockTime) Now() time.Time {
	return mt.Time
}

// Add adds d to the current (mocked) time and returns the resulting time.
func (mt *MockTime) Add(d time.Duration) time.Time {
	mt.Time = mt.Time.Add(d)
	return mt.Time
}
