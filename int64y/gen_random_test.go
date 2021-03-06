package int64y

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Code generated. DO NOT EDIT.

// TestRandom simply runs Random to ensure that it compiles during testing.
func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		Random()
	}
}

// TestRandomN runs RandomN and ensures that output is within expected range.
func TestRandomN(t *testing.T) {
	for i := 1; i < 127; i++ {
		n := RandomN(int64(i))
		require.True(t, 0 <= n && n < int64(i))
	}
}
