package remind_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/R167/go-remind"
	"github.com/stretchr/testify/require"
)

func TestDurationParsing(t *testing.T) {
	durations := []struct {
		input string
		dur   time.Duration
	}{
		{"a 5 minutes", 5 * time.Minute},
		{"a 5 minutes 30 seconds", 5*time.Minute + 30*time.Second},
		{"a 2 days", 2 * 24 * time.Hour},
		{"a week", 7 * 24 * time.Hour},
		{"a half hour", 30 * time.Minute},
		{"a minute", 1 * time.Minute},
	}

	for _, tc := range durations {
		t.Run(fmt.Sprintf("duration %s", tc.input), func(t *testing.T) {
			dur, err := remind.ParseDuration(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.dur, dur)
		})
	}
}

func BenchmarkDurationLoading(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := remind.ParseDuration("a week")
		if err != nil {
			panic(err)
		}
	}
}
