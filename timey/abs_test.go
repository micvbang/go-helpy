package timey

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDurationAbs(t *testing.T) {
	tests := map[string]struct {
		d           time.Duration
		expectedAbs time.Duration
	}{
		"500":      {d: time.Duration(500), expectedAbs: time.Duration(500)},
		"0":        {d: time.Duration(0), expectedAbs: time.Duration(0)},
		"-5":       {d: time.Duration(-5), expectedAbs: time.Duration(5)},
		"5 hours":  {d: 5 * time.Hour, expectedAbs: 5 * time.Hour},
		"-5 hours": {d: -5 * time.Hour, expectedAbs: 5 * time.Hour},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expectedAbs, DurationAbs(test.d))
		})
	}
}
